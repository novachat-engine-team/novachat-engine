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
	"novachat_engine/pkg/cmd/session/conf"
	"novachat_engine/pkg/cmd/session/handler"
)

type SessionImpl struct {
	conf           *conf.Conf
	handler        *handler.SessionHandler
	messageChannel *handler.MessageChannel
}

func NewSessionImpl(handler *handler.SessionHandler, channel *handler.MessageChannel, conf *conf.Conf) *SessionImpl {
	impl := &SessionImpl{
		handler:        handler,
		messageChannel: channel,
	}

	serverPeer := conf.RpcDiscovery.ServerName + "/" + conf.RpcServer.Addr
	_ = serverPeer
	return impl
}
