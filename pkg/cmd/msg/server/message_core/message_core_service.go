/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/10 13:38
 * @File : message_core_service.go
 */

package message_core

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/json-iterator/go"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/mq"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message/inboxes"
	"novachat_engine/service/core/message/outboxes"
	"runtime/debug"
)

type MessageCoreService struct {
	producer     mq.Producer
	chatProducer mq.Producer
	outboxesCore *outboxes.Core
	inboxesCore  *inboxes.Core
}

func NewMessageCoreService(
	producer mq.Producer,
	chatProducer mq.Producer,
	outboxesCore *outboxes.Core,
	inboxesCore *inboxes.Core) *MessageCoreService {
	outboxesCore.Migrator()
	return &MessageCoreService{
		producer:     producer,
		chatProducer: chatProducer,
		outboxesCore: outboxesCore,
		inboxesCore:  inboxesCore,
	}
}

func (m *MessageCoreService) OnMessageData(key string, value []byte) error {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("MessageCoreService::OnMessageData panic: %s", debug.Stack())
		}
	}()

	var err error

	switch key {
	case constants.InboxMessage:
		r := &msgService.SendMessages{}
		if err = proto.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.SendMessageData(r)

	case constants.InboxChatMessage:
		var r []*msgService.SendMessages
		if err = jsoniter.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.SendMessageDataList(r)

	case constants.InboxMessageReadHistory:
		r := &msgService.ReadHistory{}
		if err = proto.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.ReadHistoryMessageData(r)

	case constants.InboxChatMessageReadHistory:
		var r []*msgService.ReadHistory
		if err = jsoniter.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.ReadHistoryMessageDataList(r)

	case constants.InboxMessageRevokeMessages:
		r := &msgService.RevokeMessages{}
		if err = proto.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.RevokeMessageData(r)
	case constants.InboxChatMessageRevokeMessages:
		var r []*msgService.RevokeMessages
		if err = jsoniter.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.RevokeMessageDataList(r)
	case constants.OutboxMessagePinnedMessages:
		r := &msgService.PinnedMessage{}
		if err = proto.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.PinnedMessageData(r)
	case constants.InboxChatMessagePinnedMessages:
		var r []*msgService.PinnedMessage
		if err = jsoniter.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.PinnedMessageDataList(r)
	default:
		err = fmt.Errorf("OnMessageData invalid key: %s", key)
		log.Error(err.Error())
		return nil
	}
}

func (m *MessageCoreService) SendUserMessage(message *msgService.SendMessages) error {
	mData, _ := proto.Marshal(message)
	_, _, err := m.producer.SendMessage(constants.InboxMessage, mData)
	return err
}

func (m *MessageCoreService) SendChatMessage(message *msgService.SendMessages) error {
	mData, _ := proto.Marshal(message)
	_, _, err := m.chatProducer.SendMessage(constants.OutboxChatMessage, mData)
	return err
}
