/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/24 16:09
 * @File : session_handler.go
 */

package handler

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	metadata2 "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"novachat_engine/mtproto"
	gateClient "novachat_engine/pkg/cmd/gate/rpc_client"
	"novachat_engine/pkg/cmd/session/conf"
	"novachat_engine/pkg/cmd/session/handler/message_id"
	sessionService "novachat_engine/pkg/cmd/session/rpc_client"
	"novachat_engine/pkg/cmd/session/seq"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/mq"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/rpc/rpc"
	"novachat_engine/pkg/runtime"
	"novachat_engine/pkg/util"
	message2 "novachat_engine/service/common/message"
	"novachat_engine/service/constants"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const SaltUtilTime = int32(30 * 60)
const DefaultMQ = "default"

type SessionHandler struct {
	conf     *conf.Conf
	seqNoGen *seq.SeqNoGen
	enableMQ atomic.Value
	//mqMap          map[string]*config.MQKafkaProducerConfig
	bizRpc         *etcd.ManualClient
	bizRpcChan     sync.Map
	messageChannel *MessageChannel
	mqMap          sync.Map
}

func NewSessionHandler(conf *conf.Conf) *SessionHandler {
	s := &SessionHandler{
		conf:     conf,
		seqNoGen: seq.NewSeqNoGen(),
	}

	s.enableMQ.Store(conf.BizRpc.EnableMQ)
	if conf.BizRpc.EnableMQ && len(conf.BizRpc.BizMQList) > 0 {
		var produder mq.Producer
		for idx := range conf.BizRpc.BizMQList {
			p := mq.NewKafkaProducer(&conf.BizRpc.BizMQList[idx].KafkaProducer)
			s.mqMap.Store(conf.BizRpc.BizMQList[idx].Name, p)
			if idx == 0 {
				produder = p
			}
		}

		_, ok := s.mqMap.Load(DefaultMQ)
		if !ok {
			s.mqMap.Store(DefaultMQ, produder)
		}
	}
	var err error
	if !conf.BizRpc.EnableMQ {
		s.bizRpc, err = etcd.NewManualClient(conf.BizRpc.BizRpcClient)
		if err != nil {
			panic(err)
		}
		s.bizRpc.Start()
	}
	return s
}

func (s *SessionHandler) SetMessageChannel(q *MessageChannel) {
	s.messageChannel = q
}

func (s *SessionHandler) GetBizRpcChan(authKeyId int64) (chan<- []*rpc.RpcStreamData, bool) {
	if authKeyId == 0 {
		return nil, false
	}

	conn := s.bizRpc.GetClientConnByKey(fmt.Sprintf("%d", authKeyId))
	if conn == nil {
		return nil, false
	}

	queue, ok := s.bizRpcChan.LoadOrStore(conn, make(chan []*rpc.RpcStreamData))
	q := queue.(chan []*rpc.RpcStreamData)
	if !ok {
		runtime.GoWithRecover(func() {
			tryTime := 1
			for {
				if tryTime > 3 {
					log.Errorf("BizRpc ReceiveData connect authKeyId:%d tryTime >= 3", authKeyId)
					break
				}

				streamConn := rpc.NewRpcStreamServiceClient(conn)
				stream, err := streamConn.ReceiveData(metadata2.AppendToOutgoingContext(context.Background(), metadata.RpcMetadataAddressKey, s.conf.RpcServer.Addr))
				if err != nil {
					if strings.Contains(err.Error(), "EOF") {
						log.Warnf("BizRpc ReceiveData connect error:%s", err.Error())
						return
					}
					time.Sleep(time.Second * 3)
					log.Errorf("BizRpc ReceiveData connect error:%s", err.Error())
					tryTime++
					continue
				}

				queue, ok = s.bizRpcChan.LoadOrStore(conn, make(chan []*rpc.RpcStreamData))
				q = queue.(chan []*rpc.RpcStreamData)

				var object mtproto.TLObject
				var recvData *rpc.RpcStreamData
				var statusValue *status.Status
				done := make(chan bool)
				runtime.GoWithRecover(func() {
					for {
						select {
						case _, ok1 := <-done:
							if !ok1 {
								return
							}
						case sendObject, ok1 := <-q:
							if !ok1 {
								return
							}
							for idx := range sendObject {
								err = stream.Send(sendObject[idx])
								if err != nil {
									log.Warnf("ReceiveDataStream send reqMsgId:%d error:%s", sendObject[idx].Md.ReqMsgId, err.Error())
									return
								}
								//log.Debugf("ReceiveDataStream send reqMsgId:%d", sendObject[idx].Md.ReqMsgId)
							}
						}
					}
				}, nil)

				for {
					recvData, err = stream.Recv()
					if err != nil {
						statusValue, ok = status.FromError(err)
						if ok == true {
							err = fmt.Errorf(statusValue.Message())
							if statusValue.Code() == codes.Unavailable {

							}
						}
						log.Warnf("ReceiveDataStream recv shutdown, error:%s", err.Error())
						close(done)
						s.bizRpcChan.Delete(conn)
						close(q)
						break
					}

					messageName, _ := types.AnyMessageName(recvData.GetData())
					typ := proto.MessageType(messageName)
					for {
						if typ.Kind() == reflect.Ptr {
							typ = typ.Elem()
						} else {
							break
						}
					}
					value := reflect.New(typ)
					object, _ = value.Interface().(mtproto.TLObject)
					if object == nil {
						log.Errorf("ReceiveDataStream messageName:%s recv data:%+v", messageName, recvData.Data)
						continue
					}
					err = types.UnmarshalAny(recvData.Data, object.(proto.Message))
					if err != nil {
						log.Errorf("ReceiveDataStream recv data:%+v", recvData.Data)
						continue
					}

					session := s.messageChannel.GetChannel(recvData.Md.AuthKeyId)
					ok, _ = util.Contains(messageName, []string{"mtproto.auth_Authorization"})
					if ok {
						auth_Authorization := object.(*mtproto.Auth_Authorization)
						if auth_Authorization != nil && auth_Authorization.ClassName == mtproto.ClassAuthAuthorization {
							log.Debugf("ReceiveDataStream userId:%d auth_Authorization:%+v", auth_Authorization.User.Id, auth_Authorization)
							if !session.SessionContextUserId(recvData.Md.AuthKeyId, constants.PeerTypeFromUserIDType32(auth_Authorization.User.Id).ToInt()) {
								log.Warnf("SessionContextUserId authKeyId:%d userId:%d", recvData.Md.AuthKeyId, auth_Authorization.User.Id)
								return
							}
						}
					}

					ok, err = session.Resp2Client(recvData.Md.AuthKeyId, recvData.Md.SessionId, object, recvData.Md.ReqMsgId, true)
					if err != nil {
						log.Warnf("ReceiveDataStream Send2Client error:%s", err.Error())
					} else {
						if !ok {
							//TODO:coderxw
							log.Warnf("ReceiveDataStream Send2Client authKeyId:%d sessionId:%d reqMsgId:%d not arrived client", recvData.Md.AuthKeyId, recvData.Md.SessionId, recvData.Md.ReqMsgId)
						}
					}
				}
			}
		}, nil)
	}
	return q, true
}

func (s *SessionHandler) SetEnableMQ(enable bool) {
	s.enableMQ.Store(enable)
}

func (s *SessionHandler) Send2Client(sessionContext *SessionContext,
	msgId int64,
	reply bool,
	object mtproto.TLObject,
	reqMsgId int64,
	messageContext *MessageContext) (bool, error) {
	serverPeer, authKeyId, salt := sessionContext.SrcServer, sessionContext.AuthKeyId, sessionContext.Salt
	sessionId := messageContext.SessionId
	layer := sessionContext.AuthKeyInfo.Layer
	seqNo := s.seqNoGen.Generator(reply)
	if msgId == 0 {
		msgId = message_id.MakeNextMessageId(false)
	}
	if layer == 0 {
		layer = mtproto.CurrentLayer
	}
	//log.Debugf("SessionHandler::Send2Client serverPeer:%s, authKeyId:%d, salt:%d, sessionId:%d, messageId:%d, reqMsgId:%d reply:%v", serverPeer, authKeyId, salt, sessionId, msgId, reqMsgId, reply)

	messageTransform := mtproto.TLMessageTransform{
		MsgId:  msgId,
		Seqno:  seqNo,
		Bytes:  0,
		Object: object,
	}
	transform := mtproto.Transform{
		Salt:      salt,
		SessionId: messageContext.SessionId,
		Buf:       messageTransform.Encode(layer),
	}
	conn := gateClient.GetGateClient(serverPeer)
	if conn != nil {
		data := rpc.RpcData{}
		data.Data, _ = types.MarshalAny(&sessionService.TransformData{
			AuthKeyId: authKeyId,
			ConnId:    messageContext.ConnId,
			Buf:       transform.Encode(layer),
			MsgId:     msgId,
		})
		resp, err := conn.ReceiveData(context.Background(), &data)
		if err != nil {
			log.Infof("SessionHandler::Send2Client serverPeer:%s, authKeyId:%d, sessionId:%d, msgId:%d, object:%v, reqMsgId:%d error:%v", serverPeer, authKeyId, sessionId, msgId, object, reqMsgId, err.Error())
			return false, err
		} else {
			value := &mtproto.TLInt64{}
			err = types.UnmarshalAny(resp, value)
			log.Infof("SessionHandler::Send2Client serverPeer:%s, authKeyId:%d, sessionId:%d, msgId:%d, object:%v, resp:%v reqMsgId:%d", serverPeer, authKeyId, sessionId, msgId, object, value, reqMsgId)
			return value.Value == 1, err
		}
	} else {
		log.Errorf("SessionHandler::Send2Client serverPeer:%s, authKeyId:%d, sessionId:%d, reqMsgId:%d, object:%v reqMsgId:%d", serverPeer, authKeyId, sessionId, msgId, object, reqMsgId)
		return false, nil
	}
}

func (s *SessionHandler) Send2MQ(data *rpc.RpcStreamData) (int64, error) {
	messageName, _ := types.AnyMessageName(data.GetData())
	key := DefaultMQ
	splitList := strings.Split(messageName, "_")
	if len(splitList) > 2 {
		key = splitList[1]
	}

	producer, _ := s.mqMap.Load(key)
	value, _ := proto.Marshal(data)
	_, offset, err := producer.(mq.Producer).SendMessage(messageName, value)
	if err != nil {
		log.Errorf("Send2MQ error:%s", err.Error())
		return 0, nil
	}
	return offset, nil
}

func (s *SessionHandler) DoSessionResp(sessionContext *SessionContext, messageContext *MessageContext) (bool, error) {
	if len(messageContext.RespSessionMessageList) == 0 {
		log.Debugf("SessionHandler::DoSessionResp, authKeyId:%d, sessionId:%d ok!!!", sessionContext.AuthKeyId, messageContext.SessionId)
		return true, nil
	}

	//if len(list) == 1 && s.unmergeResponseRpc(list[0].Object) {
	if len(messageContext.RespSessionMessageList) == 1 {
		if messageContext.RespSessionMessageList[0].Object != nil {
			if messageContext.RespSessionMessageList[0].Object.Object != nil {
				if messageContext.RespSessionMessageList[0].Object.MsgId == 0 {
					messageContext.RespSessionMessageList[0].Object.MsgId = message2.MakeNextMessageId(false)
				}
				messageContext.RespSessionMessageList[0].Object.Seqno = s.seqNoGen.Generator(messageContext.RespSessionMessageList[0].Reply)

				ok, err := s.Send2Client(sessionContext, messageContext.RespSessionMessageList[0].Object.MsgId,
					messageContext.RespSessionMessageList[0].Reply,
					messageContext.RespSessionMessageList[0].Object.Object,
					messageContext.RespSessionMessageList[0].ReqMsgId,
					messageContext)
				if err != nil {
					log.Warnf("SessionHandler::DoSessionResp authKeyId:%d, sessionId:%d, error:%s", sessionContext.AuthKeyId, messageContext.SessionId, err.Error())
				}
				return ok, err
			} else {
				log.Warnf("SessionHandler::DoSessionResp authKeyId:%d, sessionId:%d object is nil", sessionContext.AuthKeyId, messageContext.SessionId)
			}

			log.Debugf("SessionHandler::DoSessionResp authKeyId:%d, sessionId:%d reqMsgId:%d ok!!!", sessionContext.AuthKeyId, messageContext.SessionId, messageContext.RespSessionMessageList[0].ReqMsgId)
		}
		return true, nil
	}

	if len(messageContext.RespSessionMessageList) > 0 {
		container := mtproto.NewTLMsgContainer()
		container.Messages = make([]*mtproto.TLMessageTransform, 0, len(messageContext.RespSessionMessageList))
		for idx := range messageContext.RespSessionMessageList {
			if messageContext.RespSessionMessageList[idx].Object.Object != nil {
				if messageContext.RespSessionMessageList[idx].Object.MsgId == 0 {
					messageContext.RespSessionMessageList[idx].Object.MsgId = message2.MakeNextMessageId(false)
				}
				messageContext.RespSessionMessageList[idx].Object.Seqno = s.seqNoGen.Generator(messageContext.RespSessionMessageList[idx].Reply)

				container.Messages = append(container.Messages, messageContext.RespSessionMessageList[idx].Object)

				log.Debugf("SessionHandler::DoSessionResp authKeyId:%d, sessionId:%d reqMsgId:%d ", sessionContext.AuthKeyId, messageContext.SessionId, messageContext.RespSessionMessageList[idx].ReqMsgId)
			}
		}
		ok, err := s.Send2Client(
			sessionContext,
			message2.MakeNextMessageId(false),
			messageContext.RespSessionMessageList[0].Reply,
			container,
			messageContext.RespSessionMessageList[0].ReqMsgId,
			messageContext)
		if err != nil {
			log.Warnf("SessionHandler::DoSessionResp authKeyId:%d, sessionId:%d, error:%s", sessionContext.AuthKeyId, messageContext.SessionId, err.Error())
		} else {
			log.Debugf("SessionHandler::DoSessionResp authKeyId:%d, sessionId:%d MsgContainer ok!!!", sessionContext.AuthKeyId, messageContext.SessionId)
		}
		_ = ok
		return ok, err
	} else {
		log.Warnf("SessionHandler::DoSessionResp authKeyId:%d, sessionId:%d, noop", sessionContext.AuthKeyId, messageContext.SessionId)
	}
	return true, nil
}

func (s *SessionHandler) unmergeResponseRpc(object *mtproto.TLMessageTransform) bool {
	if object == nil {
		log.Error("unmergeResponseRpc object is nil")
		return true
	}
	switch object.Object.(type) {
	case *mtproto.TLMsgsAck,
		*mtproto.TLMsgContainer,
		*mtproto.TLPong,
		*mtproto.TLNewSessionCreated,
		*mtproto.MsgsStateInfo,
		*mtproto.TLRpcResult,
		*mtproto.TLBadMsgNotification,
		*mtproto.TLBadServerSalt,
		*mtproto.TLMsgDetailedInfo,
		*mtproto.TLMsgNewDetailedInfo,
		*mtproto.TLGZipPacked,
		*mtproto.TLError,
		*mtproto.TLRpcError,
		*mtproto.TLFutureSalts,
		*mtproto.TLDestroyAuthKeyOk,
		*mtproto.TLDestroySessionNone,
		*mtproto.TLUpdatesTooLong:
		return true
	default:
		return false
	}
}
