/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/23 22:40
 * @File : session.go
 */

package handler

import (
	"novachat_engine/mtproto"
	authService "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/cmd/session/conf"
	sessionService "novachat_engine/pkg/cmd/session/rpc_client"
	"sync"
)

type SessionContext struct {
	AuthKeyId     int64
	SrcServer     string
	Ip            string
	Salt          int64
	AuthKeyInfo   *authService.AuthKeyInfo
	UserId        int64
	PermAuthKeyId int64
}

//
//type none struct{}
//
//const (
//	RpcRequestNormal = 0
//	RpcRequestDrop   = 1
//	//RpcRequestAck = 2
//)
//
//type RpcRequestType struct {
//	RpcResult  mtproto.TLObject
//	RpcRequest *mtproto.MessageTransform
//	QuickAckId int32
//	Date       int64
//	Status     int32
//}
//
//type RpcChannel struct {
//	QuickAck    int32
//	RpcRequest  []*RpcRequestType
//	ServerPeer  string
//	SessionId   int64
//	AuthKeyInfo *auth_pb.AuthKeyInfo
//	Ip          string
//}
//

type SessionMessage struct {
	Reply    bool
	Object   *mtproto.TLMessageTransform
	ReqMsgId int64
}

type MessageContext struct {
	RespSessionMessageList []*SessionMessage
	RequestMessageList     []*mtproto.MessageTransform
	SessionId              int64
	ConnId                 uint64
	SessionCreated         bool
}

type Session struct {
	rwLock             sync.RWMutex
	messageHandlerMap  map[int64]*MessageHandler
	sessionHandler     *SessionHandler
	conf               *conf.Conf
	messageContextPool *sync.Pool
	sessionContextPool *sync.Pool
}

func NewSession(sessionHandler *SessionHandler, conf *conf.Conf) *Session {
	return &Session{
		sessionHandler:    sessionHandler,
		conf:              conf,
		messageHandlerMap: make(map[int64]*MessageHandler),
		messageContextPool: &sync.Pool{
			New: func() interface{} {
				return &MessageContext{
					RespSessionMessageList: make([]*SessionMessage, 0, 32),
					RequestMessageList:     make([]*mtproto.MessageTransform, 0, 8),
				}
			},
		},
		sessionContextPool: &sync.Pool{
			New: func() interface{} {
				return &SessionContext{}
			},
		},
	}
}

func (m *Session) CloseSession(closeEvent *sessionService.CloseEvent) error {
	m.rwLock.Lock()
	s, ok := m.messageHandlerMap[closeEvent.AuthKeyId]
	if !ok {
		m.rwLock.Unlock()
		return nil
	}
	delete(m.messageHandlerMap, closeEvent.AuthKeyId)
	m.rwLock.Unlock()
	m.sessionContextPool.Put(s.sessionContext)
	return s.CloseSession(closeEvent)
}

func (m *Session) HandlerReceiveData(sessionContext SessionContext, message *mtproto.TLMessageTransform, sessionId int64, connId uint64) error {
	m.rwLock.RLock()
	s, ok := m.messageHandlerMap[sessionContext.AuthKeyId]
	if !ok {
		m.rwLock.RUnlock()
		sc := m.sessionContextPool.Get().(*SessionContext)
		*sc = sessionContext
		s = NewMessageHandler(sc, m.sessionHandler, m.conf)
		m.rwLock.Lock()
		m.messageHandlerMap[sessionContext.AuthKeyId] = s
		m.rwLock.Unlock()
	} else {
		m.rwLock.RUnlock()
	}
	messageContext := m.messageContextPool.Get().(*MessageContext)
	messageContext.SessionId = sessionId
	messageContext.ConnId = connId
	err := s.HandlerReceiveData(message, messageContext)
	m.messageContextPool.Put(messageContext)
	return err
}

func (m *Session) PushUpdates(authKeyId int64, sessionId []int64, updates *mtproto.Updates) (bool, error) {
	m.rwLock.RLock()
	s, ok := m.messageHandlerMap[authKeyId]
	if !ok {
		m.rwLock.RUnlock()
		return false, nil
	}
	m.rwLock.RUnlock()

	return s.PushUpdates(sessionId, updates)
}

func (m *Session) Resp2Client(
	authKeyId int64,
	sessionId int64,
	object mtproto.TLObject,
	reqMsgId int64,
	reply bool) (bool, error) {
	m.rwLock.RLock()
	s, ok := m.messageHandlerMap[authKeyId]
	if !ok {
		m.rwLock.RUnlock()
		return false, nil
	}
	m.rwLock.RUnlock()

	ctx := m.messageContextPool.Get().(*MessageContext)
	ok, err := s.Resp2Client(sessionId, object, reqMsgId, reply, ctx)
	ctx.RespSessionMessageList = ctx.RespSessionMessageList[:0]
	m.messageContextPool.Put(ctx)
	return ok, err
}

func (m *Session) SessionContextUserId(authKeyId int64, userId int64) bool {
	m.rwLock.RLock()
	s, ok := m.messageHandlerMap[authKeyId]
	if !ok {
		m.rwLock.RUnlock()
		return false
	}
	m.rwLock.RUnlock()
	s.sessionContext.UserId = userId
	return true
}
