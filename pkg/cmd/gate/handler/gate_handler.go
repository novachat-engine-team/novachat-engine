/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/2/21 16:13
 * @File : gate_handler.go
 */

package handler

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/cmd/gate/conf"
	"novachat_engine/pkg/cmd/gate/data"
	"novachat_engine/pkg/cmd/gate/handshake"
	sessionClient "novachat_engine/pkg/cmd/session/rpc_client"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/log"
	net2 "novachat_engine/pkg/net"
	"novachat_engine/pkg/net/mtserver/codec"
	rpc2 "novachat_engine/pkg/rpc/rpc"
	"novachat_engine/service/rpc_context"
	"time"
)

const (
	cacheExpire = 5 * time.Minute
)

type GateHandler struct {
	conf *conf.Config
	hs   handshake.HandShake
	//sessionManager *session.SessionManager
}

func NewGateHandler(conf *conf.Config, hs handshake.HandShake) *GateHandler {
	s := &GateHandler{
		conf: conf,
		hs:   hs,
		//sessionManager: session.NewSessionManager(),
	}

	return s
}

func (s *GateHandler) OnNewConnection(conn *net2.TcpConnection) {
	if conn.Context == nil {
		c := &data.Context{}
		c.SetStatus(data.StatusReady)
		conn.Context = c
	}

	ctx, ok := conn.Context.(*data.Context)
	if ok == true {
		ctx.Reset()
	} else {
		log.Warnf("OnNewConnection - conn.Context:%+v", conn.Context)
		c := &data.Context{}
		c.SetStatus(data.StatusReady)
		conn.Context = c
	}

	log.Debugf("OnNewConnection - conn:%+v", conn)
}

func (s *GateHandler) OnConnectionDataArrived(conn *net2.TcpConnection, m interface{}) error {

	var err error
	msg, ok := m.(*codec.MTPRawMessage)
	if ok == false {
		err = fmt.Errorf("GateHandler::OnConnectionDataArrived conn msg not codec.MTPRawMessage m = %+v", m)
		fmt.Errorf(err.Error())
		return err
	}

	ctx := conn.Context.(*data.Context)
	if msg.AuthKeyId() == 0 {
		err = s.onHandShake(ctx, conn, msg)
		if err != nil {
			ctx.Reset()
			conn.Close()
		}
	} else {
		if ctx == nil {
			c := &data.Context{}
			c.SetStatus(data.StatusReady)
			conn.Context = c
		}
		if ctx.AuthKeyId() == 0 {
			// authKeyId ok
			st := ctx.Status() //
			if st < data.StatusDHParamOk {
				var authKeyInfo *authClient.AuthKeyInfo
				ctx1, cancel := rpc_context.Context(
					context.Background(),
					rpc_context.WithServerTrace(
						rpc_context.FormatServerTrace(s.conf.Server.Addr,
							rpc_context.RunFuncName()),
					),
				)
				authKeyInfo, err = authClient.GetAuthClientByAuthKey(msg.AuthKeyId()).
					ReqAuthKey(ctx1, &authClient.AuthKey{AuthKeyId: msg.AuthKeyId()})
				cancel()
				if err != nil {
					log.Errorf(err.Error())
					return err
				}
				ctx.SetAuthKeyId(authKeyInfo.AuthKeyId)
				ctx.SetAuthKey(authKeyInfo.AuthKey)
				ctx.SetValidUntil(authKeyInfo.ValidUntil)
				ctx.SetValidSince(authKeyInfo.ValidSince)
				ctx.SetExpiresIn(authKeyInfo.ExpiresIn)
				ctx.SetCryptoAuthKey(crypto.NewAuthKey(ctx.AuthKeyId(), ctx.AuthKey()))
				if ctx.AuthKeyId() == 0 || ctx.AuthKey() == nil || ctx.CryptoAuthKey() == nil {
					err = fmt.Errorf("GateHandler::OnConnectionDataArrived process msg error msg:%+v, ctx:%+v", msg, ctx)
					//conn.Close()
					encoder := mtproto.NewEncodeBuf(4)
					encoder.Int(mtproto.RPCErrorCode_HANDSHAKE_RESET_AUTH_KEY.ToInt32())
					return conn.Send(&codec.MTPRawMessage{Payload: encoder.GetBuf()})
				}
			} else {
				err = fmt.Errorf("GateHandler::OnConnectionDataArrived process msg error msg:%+v, ctx:%+v", msg, ctx)
				log.Error(err.Error())
				encoder := mtproto.NewEncodeBuf(4)
				encoder.Int(mtproto.RPCErrorCode_HANDSHAKE_RESET_AUTH_KEY.ToInt32())
				return conn.Send(&codec.MTPRawMessage{Payload: encoder.GetBuf()})
			}
		} else {
			if ctx.CryptoAuthKey() == nil {
				ctx.SetCryptoAuthKey(crypto.NewAuthKey(ctx.AuthKeyId(), ctx.AuthKey()))
			}
		}

		err = s.onMessage(ctx, conn, msg)
		if err != nil {
			log.Errorf("GateHandler::OnConnectionDataArrived onMessage error:%s", err.Error())
		}
	}
	return err
}

func (s *GateHandler) onHandShake(ctx *data.Context, conn *net2.TcpConnection, msg *codec.MTPRawMessage) error {
	log.Debugf("GateHandler::onHandShake: conn:%+v, ctx:%+v, msg:%+v", conn, ctx, msg)
	data1, err := s.hs.OnHandshake(ctx, msg.Payload, conn.GetConnID())
	if err != nil {
		log.Warnf("GateHandler::onHandShake: conn:%+v, ctx:%+v, msg:%+v error:%s", conn, ctx, msg, err.Error())
		return err
	}

	if ctx.Status() == data.StatusDHParamOk {
		ctx.SetStatus(data.StatusAuthKey)

		ctx.SetCryptoAuthKey(crypto.NewAuthKey(ctx.AuthKeyId(), ctx.AuthKey()))
	}

	if data1 == nil {
		log.Warnf("GateHandler::onHandShake: noop conn:%+v, ctx:%+v, msg:%+v noop", conn, ctx, msg)
		return nil
	}

	return conn.Send(&codec.MTPRawMessage{Payload: data1})
}

func (s *GateHandler) onMessage(ctx *data.Context, conn *net2.TcpConnection, mmsg *codec.MTPRawMessage) error {
	mtpRwaData, err := ctx.CryptoAuthKey().AesIgeDecrypt(mmsg.Payload[8:8+16], mmsg.Payload[24:])
	if err != nil {
		log.Errorf("GateHandler::onMessage error:%v", err.Error())
		//conn.Close()
		return err
	}

	data := &rpc2.RpcData{
		Data: nil,
	}
	data.Data, _ = types.MarshalAny(&sessionClient.TransformData{
		AuthKeyId:   ctx.AuthKeyId(),
		SrcServer:   rpc_context.MakeServerPeer(&s.conf.RpcDiscovery, &s.conf.RpcServer),
		Ip:          conn.RemoteAddr().String(),
		ConnId:      conn.GetConnID(),
		Buf:         mtpRwaData,
		ValidUntil:  ctx.ValidUntil(),
		ExpiresIn:   ctx.ExpiresIn(),
		ValidSince:  ctx.ValidSince(),
		Layer:       ctx.Layer(),
		TempAuthKey: ctx.AuthKeyTemp().ToInt32(),
	})
	cli := sessionClient.GetSessionClientByKey(fmt.Sprintf("%d", ctx.AuthKeyId()))
	if cli == nil {
		log.Errorf("GateHandler::onMessage GetSessionClientByKey %s cli = nil", fmt.Sprintf("%d", ctx.AuthKeyId()))
		return nil
	}
	_, err = cli.ReceiveData(context.Background(), data)
	if err != nil {
		log.Errorf("GateHandler::onMessage ReceiveData error:%v", err.Error())
	}
	return nil
}

func (s *GateHandler) OnConnectionClosed(conn *net2.TcpConnection) {
	log.Debugf("GateHandler::OnConnectionClosed - conn RemoteAddr:%v ConnId:%d", conn.RemoteAddr(), conn.GetConnID())

	ctx, _ := conn.Context.(*data.Context)
	if ctx.AuthKeyId() != 0 { // clear authKeyId
		sessionIdList := ctx.SessionIdList()
		if len(sessionIdList) == 0 {
			return
		}
		sessionClient.GetSessionClientByKey(fmt.Sprintf("%d", ctx.AuthKeyId())).
			Close(
				context.Background(),
				&sessionClient.CloseEvent{
					AuthKeyId: ctx.AuthKeyId(),
					ConnId:    conn.GetConnID(),
				})
	}
}
