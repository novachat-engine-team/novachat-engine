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
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/cmd/chat/conf"
	"novachat_engine/pkg/cmd/chat/message_core"
	"novachat_engine/pkg/cmd/chat/service"
	"novachat_engine/pkg/mq"
	"novachat_engine/service/constants"
)

type Impl struct {
	conf           *conf.Config
	chatManager    *service.ChatManager
	chatConsumer   mq.Consumer
	coreService    *message_core.MessageCoreService
	chatInProducer mq.Producer
}

func NewImpl(conf *conf.Config) *Impl {

	cache.InstallRedis(conf.Redis)
	m := &Impl{
		conf:        conf,
		chatManager: service.NewChatManager(conf.Mongo),
	}

	m.coreService = message_core.NewMessageCoreService()

	m.chatInProducer = mq.NewKafkaProducer(&m.conf.ChatInProducer)
	m.chatConsumer = mq.NewKafkaConsumer(&m.conf.ChatConsumer, mq.WithConsumerCB(m.coreService))
	go m.chatConsumer.Run()

	return m
}

func (impl *Impl) getPiece() int32 {
	inboxPiece := impl.conf.InboxPiece
	if inboxPiece <= 0 {
		inboxPiece = constants.InboxPiece
	}

	return inboxPiece
}
