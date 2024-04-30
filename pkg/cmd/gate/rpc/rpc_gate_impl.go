/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/6 14:08
 * @File : rpc_gate_impl.go
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cmd/gate/data"
	sessionService "novachat_engine/pkg/cmd/session/rpc_client"
	"novachat_engine/pkg/log"
	net2 "novachat_engine/pkg/net"
	"novachat_engine/pkg/net/mtserver/codec"
	rpc2 "novachat_engine/pkg/rpc/rpc"
)

type Connections interface {
	GetConnection(connId uint64) *net2.TcpConnection
}

type Impl struct {
	connections Connections
}

func NewGateImpl(connections Connections) *Impl {
	impl := &Impl{connections: connections}
	return impl
}

func (s *Impl) ReceiveData(ctx context.Context, req *rpc2.RpcData) (*types.Any, error) {
	in := &sessionService.TransformData{}
	err := types.UnmarshalAny(req.Data, in)
	if err != nil {
		log.Warnf("ReceiveData req.Data:`%+v`", req.Data)
		return types.MarshalAny(&mtproto.TLInt64{
			Value: 0,
		})
	}

	conn := s.connections.GetConnection(in.ConnId)
	if conn == nil {
		log.Warnf("GetConnection authKeyId:%d, connId:%d is nil", in.AuthKeyId, in.ConnId)
		return types.MarshalAny(&mtproto.TLInt64{
			Value: 0,
		})
	}

	ctxConn, ok := conn.Context.(*data.Context)
	if ok == false {
		log.Errorf("conn.Context is not data.Context reqMsg:%d", in.ConnId)
		return types.MarshalAny(&mtproto.TLInt64{
			Value: 0,
		})
	}

	msgKey, msgRawData, err := ctxConn.CryptoAuthKey().AesIgeEncrypt(in.Buf)
	if err != nil {
		log.Errorf("CryptoAuthKey().AesIgeEncrypt error:%s reqMsg:%d", err.Error(), in.ConnId)
		return types.MarshalAny(&mtproto.TLInt64{
			Value: 0,
		})
	}

	if ctxConn.AuthKeyId() != in.AuthKeyId {
		log.Warnf("GetConnection authKeyId:%d, connId:%d AuthKeyId:%d", ctxConn.AuthKeyId(), in.ConnId, in.AuthKeyId)
		return types.MarshalAny(&mtproto.TLInt64{
			Value: 0,
		})
	}

	// len(authKeyId) + len(msgKey) + len(msgRawData)
	d := mtproto.NewEncodeBuf(8 + len(msgKey) + len(msgRawData))
	d.Long(in.AuthKeyId)
	d.Bytes(msgKey)
	d.Bytes(msgRawData)

	log.Debugf("send client msgId:%d connId:%d conn:%v", in.MsgId, in.ConnId, conn)
	err = conn.Send(&codec.MTPRawMessage{Payload: d.GetBuf()})
	if err != nil {
		log.Debugf("send client msgId:%d error:%s", in.MsgId, err.Error())
		return nil, err
	}

	return types.MarshalAny(&mtproto.TLInt64{
		Value: 1,
	})
}
