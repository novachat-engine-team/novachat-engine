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
	"context"
	"errors"
	"fmt"
	"github.com/gogo/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"runtime/debug"
)

type MessageCoreService struct{}

func NewMessageCoreService() *MessageCoreService {
	return &MessageCoreService{}
}

func (m *MessageCoreService) OnMessageData(key string, value []byte) error {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("MessageCoreService::OnMessageData panic: %s", debug.Stack())
		}
	}()

	var err error

	switch key {
	case constants.OutboxChatMessage:
		r := &msgService.SendMessages{}
		if err = proto.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.SendMessageData(r)
	case constants.OutboxChatMessageReadHistory:
		r := &msgService.ReadHistory{}
		if err = jsoniter.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.ReadHistoryMessageData(r)

	case constants.OutboxChatMessagePinnedMessages:
		r := &msgService.PinnedMessage{}
		if err = proto.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.PinnedMessageData(r)
	case constants.DeleteChannelMessages:
		r := &chatService.DeleteMessagesUpdates{}
		if err = proto.Unmarshal(value, r); err != nil {
			log.Errorf(err.Error())
			return err
		}
		return m.DeleteChannelMessagesData(r)

	default:
		err = fmt.Errorf("OnMessageData invalid key: %s", key)
		log.Error(err.Error())
		return nil
	}
}

func (m *MessageCoreService) DeleteChannelMessagesData(r *chatService.DeleteMessagesUpdates) error {
	_, err := chatService.GetChatClientByKeyId(r.PeerId).ReqDeleteMessagesUpdates(context.Background(), r)
	if errors.Is(err, mtproto.DefaultRpcError) {
		log.Warnf("DeleteChannelMessagesData userId:%d chatId:%d error:%s", r.UserId, r.PeerId)
		err = nil
	}
	if err != nil {
		log.Errorf("DeleteChannelMessagesData userId:%d chatId:%d error:%s", r.UserId, r.PeerId)
	}
	return err
}
