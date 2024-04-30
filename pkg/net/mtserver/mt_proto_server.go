/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 23:11
 * @File : mt_proto_server.go
 */

package mtserver

import (
	"fmt"
	"net"
	"novachat_engine/pkg/log"
	net2 "novachat_engine/pkg/net"
	"novachat_engine/pkg/net/mtserver/codec"
)

type MTProtoServer struct {
	*net2.Server
	callback net2.Callback
}

func NewMTProtoServer(addr string, cb net2.Callback) *MTProtoServer {
	lsn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	server := &MTProtoServer{
		callback: cb,
	}
	server.Server = net2.NewTcpServer(lsn, server)
	return server
}

func (s *MTProtoServer) GetConnection(connID uint64) *net2.TcpConnection {
	return s.Server.GetConnection(connID)
}

func (s *MTProtoServer) Serve() error {
	return s.Server.Serve()
}

func (s *MTProtoServer) Stop() {
	s.Server.Stop()
}

func (s *MTProtoServer) OnNewConnection(conn *net2.TcpConnection) {
	log.Debugf("onNewConnection %v", conn.RemoteAddr())

	if s.callback != nil {
		s.callback.OnNewConnection(conn)
	}
}

func (s *MTProtoServer) OnConnectionDataArrived(conn *net2.TcpConnection, msg interface{}) error {
	log.Debugf("onConnectionDataArrived %v", conn.RemoteAddr())

	message, ok := msg.(*codec.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("recv invalid MTPRawMessage: {%v}", msg)
		log.Error(err.Error())
		return err
	}

	if s.callback != nil {
		s.callback.OnConnectionDataArrived(conn, message)
	}

	return nil
}

func (s *MTProtoServer) OnConnectionClosed(conn *net2.TcpConnection) {
	log.Debugf("onConnectionClosed - %v", conn.RemoteAddr())

	if s.callback != nil {
		s.callback.OnConnectionClosed(conn)
	}
}
