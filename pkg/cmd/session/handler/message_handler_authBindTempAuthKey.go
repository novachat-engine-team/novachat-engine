/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/11/9 18:42
 * @File : message_handler_authBindTempAuthKey.go
 */

package handler

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/rpc_context"
)

func (m *MessageHandler) handlerAuthBindTempAuthKey(mt *mtproto.MessageTransform, messageContext *MessageContext) (bool, error) {

	var reply bool
	var err error
	var mtpRwaData []byte
	var object mtproto.TLObject
	for {
		payload := mt.Message.Object.(*mtproto.TLAuthBindTempAuthKey).EncryptedMessage
		decoder := mtproto.NewDecodeBuf(payload)
		authKeyId := decoder.Long()
		authKeyInfo := m.sessionContext.AuthKeyInfo

		ctx, cancel := rpc_context.Context(
			context.Background(),
			rpc_context.WithServerTrace(
				rpc_context.FormatServerTrace(
					rpc_context.MakeServerPeer(&m.conf.RpcDiscovery, &m.conf.RpcServer),
					rpc_context.RunFuncName()),
			),
			rpc_context.WithTimeout(10),
		)
		authKeyInfo, err = authClient.GetAuthClientByAuthKey(authKeyId).ReqAuthKey(ctx,
			&authClient.AuthKey{AuthKeyId: authKeyId})
		cancel()
		if err != nil || authKeyInfo.AuthKey == nil {
			if err != nil {
				log.Errorf("handlerAuthBindTempAuthKey GetAuthClientByAuthKey authKeyId:%d error:%s", authKeyId, err.Error())
			} else {
				log.Errorf("handlerAuthBindTempAuthKey GetAuthClientByAuthKey authKeyId:%d not exists", authKeyId)
			}

			if err != nil {
				object = &mtproto.TLRpcResult{ReqMsgId: mt.Message.MsgId, Result: mtproto.ToMTBool(false).To_BoolFalse()}
			} else {
				object = &mtproto.TLRpcResult{ReqMsgId: mt.Message.MsgId, Result: errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_ENCRYPTED_MESSAGE_INVALID).To_RpcError()}
			}
			break
		}

		cryptoAuthKey := crypto.NewAuthKey1(authKeyId, authKeyInfo.AuthKey)
		mtpRwaData, err = cryptoAuthKey.AesIgeDecrypt1(payload[8:8+16], payload[24:])
		if err != nil {
			log.Errorf("handlerAuthBindTempAuthKey authKeyId:%d error:%s", authKeyId, err.Error())
			object = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_ENCRYPTED_MESSAGE_INVALID).To_RpcError()
			break
		}

		decoder = mtproto.NewDecodeBuf(mtpRwaData)
		salt := decoder.Long()
		_ = salt
		sessionId := decoder.Long()
		_ = sessionId
		messageId := decoder.Long()
		seqNo := decoder.Int()
		messageLength := decoder.Int()

		cmd := decoder.UInt()
		if mtproto.Cmd_BindAuthKeyInner.ToUInt32() != cmd {
			log.Errorf("auth.bindTempAuthKey#cdd42a05 error cmd:0x%x", cmd)
			object = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_INPUT_REQUEST_TOO_LONG).To_RpcError()
			break
		}
		_ = seqNo
		_ = messageId
		_ = messageLength

		authKeyInner := mtproto.NewTLBindAuthKeyInner(nil)
		err = authKeyInner.Decode(decoder, m.sessionContext.AuthKeyInfo.Layer)
		if err != nil {
			log.Errorf("auth.bindTempAuthKey#cdd42a05 error:%s", err.Error())
			object = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_ENCRYPTED_MESSAGE_INVALID).To_RpcError()
			break
		}

		log.Debugf("handleSessionMessage TLAuthBindTempAuthKey AuthKeyId:%d TempAuthKeyId:%d PermAuthKeyId:%d s.authKeyId:%d", authKeyId, authKeyInner.GetTempAuthKeyId(), authKeyInner.GetPermAuthKeyId(), m.sessionContext.AuthKeyId)

		var ok mtproto.Bool
		var resp *types.Any
		ctx, cancel = rpc_context.Context(nil)
		_ = ctx
		resp, err = authClient.GetAuthClientByAuthKey(authKeyInner.GetTempAuthKeyId()).
			ReqBindTempAuthKeyAuthKeyId(ctx, &authClient.BindTempAuthKeyAuthKeyId{
				PermAuthKeyId: authKeyInner.GetPermAuthKeyId(),
				AuthKeyId:     authKeyInner.GetTempAuthKeyId(),
				SessionId:     authKeyInner.GetTempSessionId(),
			})
		cancel()
		if err != nil {
			log.Errorf("handleSessionMessage TLAuthBindTempAuthKey AuthKeyId:%d TempAuthKeyId:%d PermAuthKeyId:%d s.authKeyId:%d error:%s", authKeyId, authKeyInner.GetTempAuthKeyId(), authKeyInner.GetPermAuthKeyId(), m.sessionContext.AuthKeyId, err.Error())
			object = &mtproto.TLRpcResult{ReqMsgId: mt.Message.MsgId, Result: mtproto.ToMTBool(false).To_BoolFalse()}
			break
		}
		types.UnmarshalAny(resp, &ok)
		if !mtproto.ToBool(&ok) {
			log.Errorf("handleSessionMessage TLAuthBindTempAuthKey AuthKeyId:%d TempAuthKeyId:%d PermAuthKeyId:%d s.authKeyId:%d ok = false", authKeyId, authKeyInner.GetTempAuthKeyId(), authKeyInner.GetPermAuthKeyId(), m.sessionContext.AuthKeyId)
			object = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_UNREGISTERED)
			break
		}
		reply = true
		object = mtproto.ToMTBool(true).To_BoolTrue()
		m.sessionContext.PermAuthKeyId = authKeyInner.GetPermAuthKeyId()
		break
	}

	m.rpcResponse(mt.Message.MsgId, true, reply, object, messageContext, nil)
	return true, nil
}
