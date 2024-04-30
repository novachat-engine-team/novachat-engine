/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	authService "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/cmd/session/handler"
	sessionService "novachat_engine/pkg/cmd/session/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/rpc"
)

func (m *SessionImpl) ReceiveData(ctx context.Context, rpcData *rpc.RpcData) (*types.Any, error) {
	transformData := &sessionService.TransformData{}
	err := types.UnmarshalAny(rpcData.Data, transformData)
	if err != nil {
		log.Errorf("SessionImpl::ReceiveData TransformData error:%s", err.Error())
		return nil, err
	}

	transform := mtproto.Transform{}
	err = transform.Decode(transformData.Buf, 0)
	if err != nil {
		log.Errorf("SessionImpl::ReceiveData Transform error:%s", err.Error())
		return nil, err
	}

	decoder := mtproto.NewDecodeBuf(transform.Buf)
	message := mtproto.NewTLMessageTransform()
	err = message.Decode(decoder, 0)
	if err != nil {
		log.Errorf("SessionImpl::ReceiveData Transform error:%s", err.Error())
		return nil, err
	}
	//log.Debugf("SessionImpl::ReceiveData reqMsgId:%d gate connId:%d sessionId:%d", message.MsgId, transformData.ConnId, transform.SessionId)

	session := m.messageChannel.GetChannel(transformData.AuthKeyId)
	err = session.HandlerReceiveData(handler.SessionContext{
		AuthKeyId: transformData.AuthKeyId,
		SrcServer: transformData.SrcServer,
		Ip:        transformData.Ip,
		Salt:      transform.Salt,
		AuthKeyInfo: &authService.AuthKeyInfo{
			AuthKeyId:   transformData.AuthKeyId,
			ExpiresIn:   transformData.ExpiresIn,
			ValidSince:  transformData.ValidSince,
			ValidUntil:  transformData.ValidUntil,
			Layer:       transformData.Layer,
			TempAuthKey: transformData.TempAuthKey,
			Salt:        transform.Salt,
		},
	}, message, transform.SessionId, transformData.ConnId)
	if err != nil {
		log.Errorf("SessionImpl::ReceiveData HandlerReceiveData message:%s error:%s", message.DebugString(), err.Error())
		return nil, err
	}

	log.Infof("ReceiveDataEvent message %s", message.DebugString())

	return types.MarshalAny(mtproto.ToMTBool(true))
}
