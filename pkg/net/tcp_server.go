/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 23:11
 * @File : tcp_server.go
 */

package net

import (
	"net"
	"novachat_engine/pkg/log"
	"runtime/debug"
	"time"
)

type Server struct {
	listener       net.Listener
	connectManager *Manager
	cb             Callback
	running        bool
}

func NewTcpServer(listener net.Listener, cb Callback) *Server {
	return &Server{
		listener:       listener,
		connectManager: NewManager(),
		cb:             cb,
		running:        false,
	}
}

func (s *Server) Serve() error {
	s.running = true
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatalf("TcpServer Accept error:%s", err.Error())
			return err
		}

		newConn := NewBufferedConn(conn)

		codec := GetCodecMaker(MTProtoCodec)
		codec.Init(newConn)

		go s.establishTcpConnection(NewTcpConnection(codec, conn, 1024, s))
	}
	return nil
}

func (s *Server) Stop() {
	if s.running {
		if s.listener != nil {
			s.listener.Close()
		}
		s.connectManager.Close()
	}
}

func (s *Server) GetConnection(connID uint64) *TcpConnection {
	conn := s.connectManager.connection(connID)
	if conn == nil {
		return nil
	}
	return conn.(*TcpConnection)
}

func (s *Server) establishTcpConnection(conn *TcpConnection) {
	log.Debugf("establishTcpConnection...")
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("tcp_server handle panic: %v\n%s", err, debug.Stack())
			conn.Close()
		}
	}()

	s.OnNewConnection(conn)
	for {
		//conn.conn.SetReadDeadline(time.Now().Add(time.Minute * 1))
		//conn.conn.SetWriteDeadline(time.Now().Add(time.Minute * 1))
		conn.conn.SetReadDeadline(time.Now().Add(time.Second * 45))
		conn.conn.SetWriteDeadline(time.Now().Add(time.Second * 45))
		msg, err := conn.Receive()
		if err != nil {
			log.Warnf("conn recv error: %v", err)
			return
		}

		if msg == nil {
			continue
		}

		if err = s.OnConnectionDataArrived(conn, msg); err != nil {
		}
	}
}

func (s *Server) OnNewConnection(conn *TcpConnection) {
	s.connectManager.addConnection(conn)

	if s.cb != nil {
		s.cb.OnNewConnection(conn)
	}
}

func (s *Server) OnConnectionDataArrived(c *TcpConnection, msg interface{}) error {
	if s.cb != nil {
		return s.cb.OnConnectionDataArrived(c, msg)
	}
	return nil
}

func (s *Server) OnConnectionClosed(c *TcpConnection) {
	s.connectManager.delConnection(c)

	if s.cb != nil {
		s.cb.OnConnectionClosed(c)
	}
}
