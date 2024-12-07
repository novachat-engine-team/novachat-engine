/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package handler

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"math/rand"
	"novachat_engine/mtproto"
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/cmd/session/conf"
	"novachat_engine/pkg/cmd/session/handler/manual_msgs"
	"novachat_engine/pkg/cmd/session/handler/message_id"
	sessionService "novachat_engine/pkg/cmd/session/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/rpc/rpc"
	"novachat_engine/pkg/runtime"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/common/message"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	"novachat_engine/service/rpc_context"
	"reflect"
	"sync"
	"time"
)

const timeoutMsgAck = 60

type messageAck struct {
	MsgId     int64
	TimeStamp int64
}

type MessageHandler struct {
	sessionIdMap   sync.Map
	sessionHandler *SessionHandler
	conf           *conf.Conf
	initConnection *mtproto.TLInitConnection
	sessionContext *SessionContext
	reqMsg         sync.Map
	reqMsgPool     *sync.Pool
	ctx            context.Context
	cancel         context.CancelFunc
}

func NewMessageHandler(sessionContext *SessionContext, h *SessionHandler, conf *conf.Conf) *MessageHandler {
	ctx, cancel := context.WithCancel(context.Background())
	m := &MessageHandler{
		sessionContext: sessionContext,
		sessionHandler: h,
		conf:           conf,
		reqMsgPool: &sync.Pool{
			New: func() interface{} { return &messageAck{} },
		},
		ctx:    ctx,
		cancel: cancel,
	}
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-m.ctx.Done():
				return
			case <-ticker.C:
				now := time.Now().Unix()
				m.reqMsg.Range(func(key, value interface{}) bool {
					if value.(*messageAck).TimeStamp+timeoutMsgAck <= now {
						m.reqMsg.Delete(key)
						m.reqMsgPool.Put(value)
					}
					return true
				})
			}
		}
	}()
	return m
}

func (m *MessageHandler) CloseSession(closeEvent *sessionService.CloseEvent) error {
	m.sessionIdMap.Range(func(key, value interface{}) bool {
		m.sessionIdMap.Delete(key)
		return true
	})
	m.cancel()

	authClient.GetAuthClientByAuthKey(closeEvent.AuthKeyId).ReqClose(context.Background(), &authClient.Close{
		AuthKeyId: closeEvent.AuthKeyId,
		ConnId:    closeEvent.ConnId,
	})
	return nil
}

func (m *MessageHandler) CheckSaltInvalid(salt int64, futureSalt *authClient.AuthKeyInfo) bool {
	if futureSalt == nil {
		return false
	}

	if salt == futureSalt.GetSalt() {
		return true
	}

	return futureSalt.GetValidUntil() >= futureSalt.GetValidSince()+SaltUtilTime
}

func (m *MessageHandler) updateSaltInfo(msg *mtproto.TLMessageTransform, messageContext *MessageContext) bool {

	if m.CheckSaltInvalid(m.sessionContext.Salt, m.sessionContext.AuthKeyInfo) {
		m.sessionContext.Salt = m.sessionContext.AuthKeyInfo.Salt
		return true
	}
	ctx, cancel := rpc_context.Context(
		context.Background(),
		rpc_context.WithTimeout(constants.RpcDefaultTimeout),
		rpc_context.WithServerTrace(
			rpc_context.FormatServerTrace(rpc_context.MakeServerPeer(&m.conf.RpcDiscovery, &m.conf.RpcServer),
				rpc_context.RunFuncName()),
		),
	)
	authKeyInfo, err1 := authClient.GetAuthClientByAuthKey(m.sessionContext.AuthKeyId).ReqAuthKey(ctx, &authClient.AuthKey{AuthKeyId: m.sessionContext.AuthKeyId})
	cancel()
	if err1 != nil {
		log.Errorf("SessionHandler updateSaltInfo ReqAuthKey error:%v", err1.Error())
		return false
	}
	if authKeyInfo.AuthKeyId == 0 {
		_, err := m.sessionHandler.Send2Client(m.sessionContext, 0, false,
			errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_TEMP_AUTH_KEY_EMPTY),
			msg.MsgId, messageContext)
		if err != nil {
			log.Errorf("SessionHandler updateSaltInfo Send2Client error: %s", err.Error())
		}
		return false
	} else {
		authKeyInfo.AuthKey = nil
		m.sessionContext.AuthKeyInfo = authKeyInfo
		if authKeyInfo.PermAuthKeyId != 0 {
			m.sessionContext.PermAuthKeyId = authKeyInfo.PermAuthKeyId
		}
	}
	if m.CheckSaltInvalid(m.sessionContext.Salt, m.sessionContext.AuthKeyInfo) {
		m.sessionContext.Salt = m.sessionContext.AuthKeyInfo.Salt
		return true
	}

	var object mtproto.TLObject
	object = mtproto.NewTLBadServerSalt(&mtproto.BadMsgNotification{
		BadMsgId:      msg.MsgId,
		BadMsgSeqno:   msg.Seqno,
		ErrorCode:     constants.ServerSaltIncorrect,
		NewServerSalt: m.sessionContext.AuthKeyInfo.GetSalt(),
	}).To_BadMsgNotification()

	_, err := m.sessionHandler.Send2Client(m.sessionContext, 0, false, object, msg.MsgId, messageContext)
	if err != nil {
		log.Errorf("SessionHandler updateSaltInfo Send2Client error: %s", err.Error())
		return true
	}
	m.sessionContext.Salt = m.sessionContext.AuthKeyInfo.Salt
	m.sessionContext.AuthKeyInfo.ValidSince = int32(time.Now().Unix())
	m.sessionContext.AuthKeyInfo.ValidUntil = m.sessionContext.AuthKeyInfo.ValidSince + SaltUtilTime
	return true
}

func (m *MessageHandler) HandlerReceiveData(msg *mtproto.TLMessageTransform,
	messageContext *MessageContext) error {
	var err error
	if !m.updateSaltInfo(msg, messageContext) {
		log.Warnf("HandlerReceiveData updateSaltInfo salt info ")
		return nil
	}

	msgs, ok, err := m.parsedManually(msg)
	if err != nil {
		log.Errorf("SessionHandler HandlerReceiveData parsedManually error: %s", err.Error())
		return err
	}
	if !ok {
		log.Errorf("SessionHandler HandlerReceiveData parsedManually ok = false msg:%+v", msg)

		_, err = m.sessionHandler.Send2Client(m.sessionContext, 0, false,
			errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_TEMP_AUTH_KEY_EMPTY),
			msg.MsgId, messageContext)
		if err != nil {
			log.Errorf("SessionHandler updateSaltInfo Send2Client error: %s", err.Error())
		} else {
			err = fmt.Errorf("ReqAuthKeyWithLayer error")
		}
		return err
	}

	for idx := range msgs {
		oldConnId, ok := m.sessionIdMap.Load(messageContext.SessionId)
		if !ok || oldConnId.(uint64) != messageContext.ConnId {
			m.OnNewSessionCreated(msgs[idx].Message.MsgId, messageContext)
			m.sessionIdMap.Store(messageContext.SessionId, messageContext.ConnId)
		}

		ok, err = m.handleSessionMessage(msgs[idx], messageContext)
		if err != nil {
			return err
		}
		if ok == false {
			messageContext.RequestMessageList = append(messageContext.RequestMessageList, msgs[idx])
		}
	}

	m.doSessionResp(messageContext)
	m.doRpcRequest(messageContext)
	return err
}

func (m *MessageHandler) doRpcRequest(messageContext *MessageContext) {
	if len(messageContext.RequestMessageList) == 0 {
		return
	}

	enableMQ := m.sessionHandler.enableMQ.Load().(bool)
	var rpcStreamDataChan chan<- []*rpc.RpcStreamData
	if !enableMQ {
		var ok bool
		rpcStreamDataChan, ok = m.sessionHandler.GetBizRpcChan(m.sessionContext.AuthKeyId)
		if !ok {
			log.Fatalf("RpcRequest not found Biz authKeyId:%d", m.sessionContext.AuthKeyId)
			return
		}
	}

	var err error
	//var reqBinnedUser bool
	var userId int64
	requestData := make([]*rpc.RpcStreamData, 0, len(messageContext.RequestMessageList))
	for idx := range messageContext.RequestMessageList {
		typeName := reflect.TypeOf(messageContext.RequestMessageList[idx].Message.Object).String()
		uploadSession := conf.IsUploadSession(typeName)
		withoutLogin := conf.WithoutLoginConfig(typeName)

		if !uploadSession && !withoutLogin {
			if m.sessionContext.PermAuthKeyId == 0 {
				resp, err1 := authClient.GetAuthClientByAuthKey(m.sessionContext.AuthKeyId).ReqAuthKey(context.Background(), &authClient.AuthKey{AuthKeyId: m.sessionContext.AuthKeyId})
				if err1 != nil {
					log.Errorf("doRpcRequest GetAuthClientByAuthKey authKeyId:%d sessionId:%d error:%s", m.sessionContext.AuthKeyId, messageContext.SessionId, err.Error())
					return
				}
				if resp.PermAuthKeyId == 0 {
					m.rpcResponse(messageContext.RequestMessageList[idx].Message.MsgId, true, true, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_INVALID).To_RpcError(), messageContext, nil)
					log.Warnf("RpcRequest authKeyId:%d typeName:%v reqMsgId:%d", m.sessionContext.AuthKeyId, typeName, messageContext.RequestMessageList[idx].Message.MsgId)
					break
				}
				m.sessionContext.PermAuthKeyId = resp.PermAuthKeyId
			}
			resp, err1 := authClient.GetAuthClientByAuthKey(m.sessionContext.PermAuthKeyId).
				ReqBindedUser(context.Background(), &authClient.BindedUser{
					PermAuthKeyId: m.sessionContext.PermAuthKeyId,
					AuthKeyId:     m.sessionContext.AuthKeyId,
					SessionId:     messageContext.SessionId})
			if err1 != nil {
				log.Errorf("doRpcRequest ReqBindedUser PermAuthKeyId:%d authKeyId:%d sessionId:%d error:%s", m.sessionContext.PermAuthKeyId, m.sessionContext.AuthKeyId, messageContext.SessionId, err1.Error())
				return
			}
			var v mtproto.TLInt64
			err1 = types.UnmarshalAny(resp, &v)
			if err1 != nil {
				log.Errorf("doRpcRequest ReqBindedUser authKeyId:%d sessionId:%d UnmarshalAny error:%s", m.sessionContext.AuthKeyId, messageContext.SessionId, err1.Error())
				return
			}

			userId = v.Value
			if userId == 0 {
				m.rpcResponse(messageContext.RequestMessageList[idx].Message.MsgId, true, true, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_INVALID).To_RpcError(), messageContext, nil)
				log.Debugf("doRpcRequest ReqBindedUser authKeyId:%d permAuthKeyId:%d sessionId:%d RpcErrorCode_UNAUTHORIZED_AUTH_KEY_INVALID", m.sessionContext.AuthKeyId, m.sessionContext.PermAuthKeyId, messageContext.SessionId)
				log.Warnf("RpcRequest authKeyId:%d typeName:%v reqMsgId:%d", m.sessionContext.AuthKeyId, typeName, messageContext.RequestMessageList[idx].Message.MsgId)
				break
			}
			//m.sessionContext.UserId = v.Value
			//if m.sessionContext.UserId == 0 {
			//	m.rpcResponse(messageContext.RequestMessageList[idx].Message.MsgId, true, true, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_INVALID).To_RpcError(), messageContext, nil)
			//	log.Debugf("doRpcRequest ReqBindedUser authKeyId:%d permAuthKeyId:%d sessionId:%d RpcErrorCode_UNAUTHORIZED_AUTH_KEY_INVALID", m.sessionContext.AuthKeyId, m.sessionContext.PermAuthKeyId, messageContext.SessionId)
			//	log.Warnf("RpcRequest authKeyId:%d typeName:%v reqMsgId:%d", m.sessionContext.AuthKeyId, typeName, messageContext.RequestMessageList[idx].Message.MsgId)
			//	break
			//}

			//reqBinnedUser = true
		}
		//log.Debugf("RpcRequest authKeyId:%d typeName:%v reqMsgId:%d", m.sessionContext.AuthKeyId, typeName, messageContext.RequestMessageList[idx].Message.MsgId)

		any, _ := types.MarshalAny(messageContext.RequestMessageList[idx].Message.Object.(proto.Message))
		md := metadata.RpcMetaData{
			AuthKeyId:         m.sessionContext.AuthKeyId,
			SessionId:         messageContext.SessionId,
			UserId:            userId,
			ServerPeer:        m.sessionContext.SrcServer,
			Layer:             m.sessionContext.AuthKeyInfo.Layer,
			Ip:                m.sessionContext.Ip,
			SessionServerPeer: rpc_context.MakeServerPeer(&m.conf.RpcDiscovery, &m.conf.RpcServer),
			ServerTrace: []string{
				rpc_context.FormatServerTrace(rpc_context.MakeServerPeer(&m.conf.RpcDiscovery, &m.conf.RpcServer), rpc_context.RunFuncName()),
			},
			ReqMsgId:      messageContext.RequestMessageList[idx].Message.MsgId,
			PermAuthKeyId: m.sessionContext.PermAuthKeyId,
		}
		if m.initConnection != nil {
			md.ApiId = m.initConnection.ApiId
			md.DeviceModel = m.initConnection.DeviceModel
			md.SystemVersion = m.initConnection.SystemVersion
			md.AppVersion = m.initConnection.AppVersion
			md.SystemLangCode = m.initConnection.SystemLangCode
			md.LangPack = m.initConnection.LangPack
			md.LangCode = m.initConnection.LangCode
			if m.initConnection.Proxy != nil {
				md.ProxyAddress = m.initConnection.Proxy.GetAddress()
				md.ProxyPort = m.initConnection.Proxy.GetPort()
			}
		}

		if enableMQ {
			_, err = m.sessionHandler.Send2MQ(&rpc.RpcStreamData{
				Data: any,
				Md:   &md,
			})
			if err != nil {
				log.Fatalf("RpcRequest Send2MQ authKeyId:%d request len:%d error:%s", m.sessionContext.AuthKeyId, len(messageContext.RequestMessageList), err.Error())
				break
			}
		} else {
			requestData = append(requestData, &rpc.RpcStreamData{
				Data: any,
				Md:   &md,
			})
		}
	}
	if !enableMQ {
		runtime.GoSafeWithRecover(func(params []interface{}) {
			rpcStreamDataChan <- requestData
		}, nil, func(r interface{}) {
			log.Warnf("rpc close biz %v", m.sessionContext.AuthKeyId)
		})
	}
	messageContext.RequestMessageList = messageContext.RequestMessageList[:0]
}

func (m *MessageHandler) doSessionResp(messageContext *MessageContext) {
	if len(messageContext.RespSessionMessageList) == 0 {
		return
	}

	m.sessionHandler.DoSessionResp(m.sessionContext, messageContext)
	messageContext.RespSessionMessageList = messageContext.RespSessionMessageList[:0]
}

func (m *MessageHandler) parsedManually(msg *mtproto.TLMessageTransform) ([]*mtproto.MessageTransform, bool, error) {
	//message msg_id:long seqno:int bytes:int body:Object = Message; // parsed manually
	//msg_container#73f1f8dc messages:vector<message> = MessageContainer; // parsed manually
	//msg_copy#e06046b2 orig_message:Message = MessageCopy; // parsed manually, not used - use msg_container
	//gzip_packed#3072cfa1 packed_data:string = Object; // parsed manually

	var err error
	msgs := make([]*mtproto.MessageTransform, 0, 1)
	switch msg.Object.(type) {
	case *mtproto.TLMsgContainer:
		for _, r := range msg.Object.(*mtproto.TLMsgContainer).Messages {
			ms, ok, e := m.parsedManually(r)
			if e != nil {
				log.Errorf("parsedManually TLMsgContainer error:%s", err.Error())
				return nil, ok, e
			}
			msgs = append(msgs, ms...)
		}
		break
	case *mtproto.TLMsgCopy:
		ms, ok, e := m.parsedManually(msg.Object.(*mtproto.TLMsgCopy).OrigMessage)
		if err != nil {
			log.Errorf("parsedManually TLMsgCopy error:%s", err.Error())
			return nil, ok, e
		}

		msgs = append(msgs, ms...)
		break
	case *mtproto.TLGZipPacked:
		decoder := mtproto.NewDecodeBuf(msg.Object.(*mtproto.TLGZipPacked).PackedData)
		ms, ok, e := m.parsedManually(&mtproto.TLMessageTransform{
			MsgId:  msg.MsgId,
			Seqno:  msg.Seqno,
			Object: decoder.Object(m.sessionContext.AuthKeyInfo.Layer),
		})
		if e != nil {
			err = e
			log.Errorf("parsedManually TLGZipPacked error:%s", err.Error())
			return nil, ok, err
		}
		msgs = append(msgs, ms...)
		break
	case *mtproto.TLInvokeWithLayer:
		invokeWithLayer := msg.Object.(*mtproto.TLInvokeWithLayer)
		ms, initConnection, e := manual_msgs.ManualTLInvokeWithLayer(msg.MsgId, msg.Seqno, invokeWithLayer)
		if e != nil {
			log.Errorf("parsedManually ManualTLInvokeWithLayer error:%s", err.Error())
			return nil, false, e
		}
		msgs = append(msgs, ms...)

		if m.sessionContext.AuthKeyInfo.Layer != invokeWithLayer.Layer {
			m.sessionContext.AuthKeyInfo.Layer = invokeWithLayer.Layer
			//TODO:Codexw log and retry
			ret, err1 := authClient.GetAuthClientByAuthKey(m.sessionContext.AuthKeyId).
				ReqAuthKeyWithLayer(context.Background(),
					&authClient.AuthKeyWithLayer{
						AuthKeyId: m.sessionContext.AuthKeyId,
						Layer:     invokeWithLayer.Layer,
					})
			if err1 != nil {
				log.Errorf("ReqAuthKeyWithLayer error:%s", err1.Error())
				return nil, false, err1
			}

			var ok mtproto.Bool
			types.UnmarshalAny(ret, &ok)
			if !mtproto.ToBool(&ok) {
				log.Errorf("ReqAuthKeyWithLayer ok = false authKeyId:%d layer:%d", m.sessionContext.AuthKeyInfo, invokeWithLayer.Layer)
				return nil, false, nil
			}
		}
		m.initConnection = initConnection
		break

	case *mtproto.TLInvokeAfterMsg: // unknown
		invokeAfterMsg, _ := msg.Object.(*mtproto.TLInvokeAfterMsg)
		msg.Object = invokeAfterMsg.Query
		return m.parsedManually(msg)
	case *mtproto.TLInvokeAfterMsgs:
		invokeAfterMsgs, _ := msg.Object.(*mtproto.TLInvokeAfterMsgs)
		msg.Object = invokeAfterMsgs.Query
		return m.parsedManually(msg)
	case *mtproto.TLInvokeWithoutUpdates: //ios
		invokeWithoutUpdates, _ := msg.Object.(*mtproto.TLInvokeWithoutUpdates)
		msg.Object = invokeWithoutUpdates.Query
		return m.parsedManually(msg)
	case *mtproto.TLInvokeWithMessagesRange: //ios
		invokeWithMessagesRange, _ := msg.Object.(*mtproto.TLInvokeWithMessagesRange)
		msg.Object = invokeWithMessagesRange.Query
		return m.parsedManually(msg)
	case *mtproto.TLInvokeWithTakeout: //ios
		invokeWithTakeout, _ := msg.Object.(*mtproto.TLInvokeWithTakeout)
		msg.Object = invokeWithTakeout.Query
		return m.parsedManually(msg)
	default:
		msgs = append(msgs, &mtproto.MessageTransform{
			Message: *msg,
		})
	}

	return msgs, true, nil
}

func (m *MessageHandler) response(resMsgId int64, isRpc bool, reply bool, object mtproto.TLObject, messageContext *MessageContext, err error) {
	msg := &mtproto.TLMessageTransform{
		MsgId: message.MakeNextMessageId(isRpc),
	}

	if err != nil {
		if err1, ok := err.(*mtproto.TLRpcError); ok == true {
			object = err1
		} else {
			object = errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.GetClientCode(), err.Error()).To_RpcError()
		}
	}

	msg.Object = object

	messageContext.RespSessionMessageList = append(messageContext.RespSessionMessageList, &SessionMessage{
		Reply:    reply,
		Object:   msg,
		ReqMsgId: resMsgId,
	})
}

func (m *MessageHandler) rpcResponse(resMsgId int64, isRpc bool, reply bool, object mtproto.TLObject, messageContext *MessageContext, err error) {
	msg := &mtproto.TLMessageTransform{
		MsgId: message.MakeNextMessageId(isRpc),
	}

	if err != nil {
		if err1, ok := err.(*mtproto.TLRpcError); ok == true {
			object = err1
		} else {
			object = errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.GetClientCode(), err.Error()).To_RpcError()
		}
	}
	if reply {
		m.reqMsg.Store(msg.MsgId, &messageAck{
			MsgId:     resMsgId,
			TimeStamp: time.Now().Unix(),
		})
	}
	msg.Object = &mtproto.TLRpcResult{ReqMsgId: resMsgId, Result: object}

	messageContext.RespSessionMessageList = append(messageContext.RespSessionMessageList, &SessionMessage{
		Reply:    reply,
		Object:   msg,
		ReqMsgId: resMsgId,
	})
}

func (m *MessageHandler) handleSessionMessage(mt *mtproto.MessageTransform, messageContext *MessageContext) (bool, error) {
	switch mt.Message.Object.(type) {
	case *mtproto.TLHelpTest: // only ios
		m.response(mt.Message.MsgId, false, true, mtproto.ToMTBool(true).To_BoolTrue(), messageContext, nil)
		log.Infof("TLHelpTest authKeyId:%d sessionId:%d msg:%d", m.sessionContext.AuthKeyId, messageContext.SessionId, mt.Message.MsgId)
		return true, nil

	case *mtproto.TLPing:
		pong := mtproto.NewTLPong(nil)
		pong.SetPingId(mt.Message.Object.(*mtproto.TLPing).PingId)
		pong.SetMsgId(mt.Message.MsgId)
		m.response(mt.Message.MsgId, false, true, pong.To_Pong(), messageContext, nil)

		log.Infof("TLPing authKeyId:%d sessionId:%d", m.sessionContext.AuthKeyId, messageContext.SessionId)
		return true, nil
	case *mtproto.TLPingDelayDisconnect:
		pingDelayDisconnect := mt.Message.Object.(*mtproto.TLPingDelayDisconnect)
		pong := mtproto.NewTLPong(nil)
		pong.SetPingId(pingDelayDisconnect.PingId)
		pong.SetMsgId(mt.Message.MsgId)
		m.response(mt.Message.MsgId, false, true, pong.To_Pong(), messageContext, nil)
		log.Infof("TLPingDelayDisconnect authKeyId:%d sessionId:%d", m.sessionContext.AuthKeyId, messageContext.SessionId)
		return true, nil
	case *mtproto.TLAuthBindTempAuthKey:
		return m.handlerAuthBindTempAuthKey(mt, messageContext)
	case *mtproto.TLMsgsAck:
		idList := mt.Message.Object.(*mtproto.TLMsgsAck).GetMsgIds()
		msgIdList := make([]int64, 0, len(idList))
		for _, msgId := range idList {
			vv, ok := m.reqMsg.Load(msgId)
			if ok {
				msgIdList = append(msgIdList, vv.(*messageAck).MsgId)
			}
		}

		m.response(mt.Message.MsgId, false, false, mtproto.NewTLMsgsAck(&mtproto.MsgsAck{
			MsgIds: msgIdList,
		}).To_MsgsAck(), messageContext, nil)
		return true, nil

	case *mtproto.TLDestroyAuthKey: // all
		m.response(mt.Message.MsgId, false, false, mtproto.NewTLDestroyAuthKeyOk(nil).To_DestroyAuthKeyRes(), messageContext, nil)
		return true, nil
	case *mtproto.TLRpcDropAnswer: // all
		dropAnswer, _ := mt.Message.Object.(*mtproto.TLRpcDropAnswer)
		_ = dropAnswer
		//v, ok := s.apiRequestMessage.Load(dropAnswer.ReqMsgId)
		//if ok == true {
		//	v.(*RpcRequestType).Status = RpcRequestDrop
		//}
		//
		//ack := mtproto.NewTLMsgsAck(nil)
		//ack.SetMsgIds([]int64{dropAnswer.ReqMsgId})
		//msg.Object = ack
		//s.respSessionMessageList = append(s.respSessionMessageList, &handler.SessionMessageType{
		//	Salt:      s.salt,
		//	SessionId: s.sessionId,
		//	Reply:     false,
		//	Object:    msg,
		//	ReqMsgId:  m.Message.MsgId,
		//})
		return true, nil
	case *mtproto.TLGetFutureSalts: // GENERIC
		getFutureSalts, _ := mt.Message.Object.(*mtproto.TLGetFutureSalts)
		m.onGetFutureSalts(mt.Message.MsgId, mt.Message.Seqno, getFutureSalts, messageContext)
		return true, nil
	case *mtproto.TLHttpWait: //
		//c.onHttpWait(id, message.MsgId, message.Seqno, message.Object)

	case *mtproto.TLDestroySession: // GENERIC
		destroySession, _ := mt.Message.Object.(*mtproto.TLDestroySession)
		m.onDestroySession(mt.Message.MsgId, mt.Message.Seqno, destroySession, messageContext)
		return true, nil

	case *mtproto.TLMsgsStateReq: // android unused
		idList := mt.Message.Object.(*mtproto.TLMsgsStateReq).GetMsgIds()
		msgIdList := make([]int64, 0, len(idList))
		for _, msgId := range idList {
			vv, ok := m.reqMsg.Load(msgId)
			if ok {
				msgIdList = append(msgIdList, vv.(*messageAck).MsgId)
			}
		}

		m.response(mt.Message.MsgId, false, false, mtproto.NewTLMsgsStateReq(&mtproto.MsgsStateReq{
			MsgIds: msgIdList,
		}).To_MsgsStateReq(), messageContext, nil)
		return true, nil
	case *mtproto.TLMsgsStateInfo: // android unused
		//c.onMsgsStateInfo(id, message.MsgId, message.Seqno, message.Object)
		return true, nil
	case *mtproto.TLMsgsAllInfo: // android unused
		//c.onMsgsAllInfo(id, message.MsgId, message.Seqno, message.Object)
		return true, nil
	case *mtproto.TLMsgResendReq: // all
		//c.onMsgResendReq(id, message.MsgId, message.Seqno, message.Object)
		return true, nil
	case *mtproto.TLMsgDetailedInfo: // all
		// log.Error("client side msg: ", object)
		return true, nil
	case *mtproto.TLMsgNewDetailedInfo: // all
		// log.Error("client side msg: ", object)
	case *mtproto.TLContestSaveDeveloperInfo: // known
		return true, nil
	case *mtproto.InitConnection:
		return true, nil
	case *mtproto.TLJsonObject:
		jsonObject := mt.Message.Object.(*mtproto.TLJsonObject)

		value := jsonObject.GetValue99C1D49D90()

		ret := "["
		for _, v := range value {
			ret += fmt.Sprintf("key:%v v:%v, ", v.Key, manual_msgs.JsonValue(v.Value))
		}
		ret += "]"

		log.Infof("*mtproto.TLJsonObject value:%v", ret)
		return true, nil

		//contestSaveDeveloperInfo, _ := message.Object.(*mtproto.TLContestSaveDeveloperInfo)
		//c.onContestSaveDeveloperInfo(id, message.MsgId, message.Seqno, contestSaveDeveloperInfo)
	default:
		break
	}
	return false, nil
}

func (m *MessageHandler) OnNewSessionCreated(msgId int64, messageContext *MessageContext) {
	newSessionCreated := mtproto.NewTLNewSessionCreated(nil)
	newSessionCreated.SetServerSalt(m.sessionContext.Salt)
	newSessionCreated.SetFirstMsgId(msgId)
	newSessionCreated.SetUniqueId(rand.Int63())

	log.Debugf("Session - OnNewSessionCreated:%v", newSessionCreated)

	messageContext.SessionCreated = true
	m.response(msgId, false, false, newSessionCreated.To_NewSession(), messageContext, nil)
	//m.response(msgId, false, false, mtproto.NewTLMsgsAck(&mtproto.MsgsAck{
	//	MsgIds: []int64{msgId},
	//}).To_MsgsAck(), messageContext, nil)
}

func (m *MessageHandler) onGetFutureSalts(msgId int64, seqNo int32, request *mtproto.TLGetFutureSalts, messageContext *MessageContext) {
	log.Debugf("onGetFutureSalts - request data: {msg_id: %d, seq_no: %d, request: {%s}}",
		msgId,
		seqNo,
		request)
	ctx, cancel := rpc_context.Context(
		context.Background(),
		rpc_context.WithServerTrace(
			rpc_context.FormatServerTrace(rpc_context.MakeServerPeer(&m.conf.RpcDiscovery, &m.conf.RpcServer),
				rpc_context.RunFuncName()),
		),
	)
	authKeyInfo, _ := authClient.GetAuthClientByAuthKey(m.sessionContext.AuthKeyId).ReqAuthKey(ctx,
		&authClient.AuthKey{
			AuthKeyId: m.sessionContext.AuthKeyId,
		})
	cancel()
	if authKeyInfo != nil && authKeyInfo.Salt != 0 {
		futureSalts := mtproto.NewTLFutureSalts(&mtproto.FutureSalts{
			ReqMsgId: msgId,
			Now:      int32(time.Now().Unix()),
			Salts: []*mtproto.TLFutureSalt{
				mtproto.NewTLFutureSalt(&mtproto.FutureSalt{
					ValidSince: authKeyInfo.ValidSince,
					ValidUntil: authKeyInfo.ValidUntil,
					Salt:       authKeyInfo.Salt,
				}),
			},
		})
		m.response(msgId, false, false, futureSalts.To_FutureSalts(), messageContext, nil)
	} else {
		m.response(msgId, false, false, mtproto.NewTLBadMsgNotification(&mtproto.BadMsgNotification{
			BadMsgId:    msgId,
			BadMsgSeqno: seqNo,
			ErrorCode:   constants.InvalidContainer,
		}).To_BadMsgNotification(), messageContext, nil)
	}
}

func (m *MessageHandler) onDestroySession(id int64, seqno int32, session *mtproto.TLDestroySession, messageContext *MessageContext) {
	destroySessionOk := &mtproto.TLDestroySessionOk{Data2: &mtproto.DestroySessionRes{
		SessionId: session.SessionId,
	}}
	m.response(id, false, false, destroySessionOk.To_DestroySessionRes(), messageContext, nil)
}

func (m *MessageHandler) PushUpdates(sessionIdList []int64, updates *mtproto.Updates) (bool, error) {
	if len(sessionIdList) > 0 {
		return m.pushUpdatesSessionIdList(sessionIdList, updates)
	} else {
		var connId uint64
		var ok bool
		var err error
		var okResult bool
		connIdMap := make(map[uint64]int64)
		m.sessionIdMap.Range(func(key, value interface{}) bool {
			connId, _ = value.(uint64)
			if connId > 0 {
				connIdMap[connId] = key.(int64)
			}
			return true
		})
		for k, v := range connIdMap {
			ok, err = m.sessionHandler.Send2Client(m.sessionContext, message_id.MakeNextMessageId(false), false,
				compat.UpdatesCompat(updates, m.sessionContext.AuthKeyInfo.Layer), 0,
				&MessageContext{SessionId: v, ConnId: k})
			if err != nil {
				log.Warnf("PushUpdates Send2Client sessionId:%d error:%s", v, err.Error())
			}
			okResult = ok || okResult
		}
		return okResult, nil
	}
}

func (m *MessageHandler) pushUpdatesSessionIdList(sessionIdList []int64, updates *mtproto.Updates) (bool, error) {
	var err error
	var okResult bool
	for _, sessionId := range sessionIdList {
		connId, ok := m.sessionIdMap.Load(sessionId)
		if !ok {
			continue
		}
		connIdValue, ok := connId.(uint64)
		if connIdValue == 0 || !ok {
			if !ok {
				log.Fatalf("PushUpdates Send2Client sessionId:%d connId:%+v", sessionId, connId)
			}
			continue
		}

		ok, err = m.sessionHandler.Send2Client(m.sessionContext, message_id.MakeNextMessageId(false), false,
			compat.UpdatesCompat(updates, m.sessionContext.AuthKeyInfo.Layer), 0, &MessageContext{SessionId: sessionId, ConnId: connIdValue})
		if err != nil {
			log.Warnf("PushUpdates Send2Client sessionId:%d error:%s", sessionId, err.Error())
		}
		okResult = ok || okResult
	}
	return okResult, nil
}

func (m *MessageHandler) Resp2Client(sessionId int64, object mtproto.TLObject, reqMsgId int64, reply bool, ctx *MessageContext) (bool, error) {
	connId, ok := m.sessionIdMap.Load(sessionId)
	if !ok {
		return false, nil
	}

	m.rpcResponse(reqMsgId, true, reply, object, ctx, nil)
	ctx.SessionId = sessionId
	ctx.ConnId = connId.(uint64)

	return m.sessionHandler.DoSessionResp(m.sessionContext, ctx)
}
