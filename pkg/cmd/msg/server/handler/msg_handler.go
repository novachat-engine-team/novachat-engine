/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/10 14:02
 * @File : msg_handler.go
 */

package handler

import (
	"fmt"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cmd/msg/conf"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/cmd/msg/server/message_core"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/mq"
	"novachat_engine/pkg/snowflake"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message/inboxes"
	"novachat_engine/service/core/message/outboxes"
)

type MsgHandler struct {
	conf            *conf.Config
	messageConsumer mq.Consumer
	chatInConsumer  mq.Consumer
	messageProducer mq.Producer
	chatProducer    mq.Producer
	coreService     *message_core.MessageCoreService
	messageIdNode   *snowflake.Node
}

func NewMsgHandler(conf *conf.Config) *MsgHandler {
	var err error
	m := &MsgHandler{}
	m.conf = conf
	m.messageIdNode, err = snowflake.NewNode(int64(conf.Server.ServerId))
	if err != nil {
		panic(err.Error())
	}

	m.chatProducer = mq.NewKafkaProducer(&m.conf.ChatProducer)
	m.messageProducer = mq.NewKafkaProducer(&m.conf.MessageProducer)
	m.coreService = message_core.NewMessageCoreService(
		m.messageProducer,
		m.chatProducer,
		outboxes.NewOutboxesCore(&conf.MongoDB, m.messageIdNode),
		inboxes.NewInboxesCore(&conf.MongoDB, m.messageIdNode))

	m.messageConsumer = mq.NewKafkaConsumer(&m.conf.MessageConsumer)
	m.messageConsumer.SetOnMessageData(m.coreService)
	go m.messageConsumer.Run()

	m.chatInConsumer = mq.NewKafkaConsumer(&m.conf.ChatInConsumer)
	m.chatInConsumer.SetOnMessageData(m.coreService)
	go m.chatInConsumer.Run()

	syncClient.InstallSyncClient(m.conf.SyncRpcClient)
	//chatClient.InstallChatClient(m.conf.ChatRpcClient)
	return m
}

func (m *MsgHandler) Close() {
	m.messageProducer.Close()
	m.messageConsumer.Close()
}

func (m *MsgHandler) SendMessageOutBoxes(messages *msgService.SendMessages) (*mtproto.Updates, error) {
	switch constants.PeerTypeFromInt32(messages.PeerType) {
	case constants.PeerTypeSelf, constants.PeerTypeUser:
		return m.coreService.SendOutboxUserMessages(messages.AuthKeyId, messages.FromUserId, messages.PeerId, messages.DataList, messages.ClearDraft)
	case constants.PeerTypeChat:
		return m.coreService.SendOutboxChatMessages(messages.AuthKeyId, messages.FromUserId, messages.PeerId, messages.DataList, messages.ClearDraft)
	//case constants.PeerTypeChannel:
	//	return m.coreService.SendOutboxChannelMessages(messages.AuthKeyId, messages.FromUserId, messages.PeerId, messages.DataList, messages.ClearDraft)
	default:
		log.Errorf("SendMessageData peerType:%+v", messages.PeerType)
		return nil, fmt.Errorf("SendMessageData peerType:%+v", messages.PeerType)
	}
}

func (m *MsgHandler) SendEditMessageOutBoxes(edit *msgService.EditMessage) (*mtproto.Updates, error) {
	panic("SendEditMessageOutBoxes")
}

func (m *MsgHandler) ReadHistory(readHistory *msgService.ReadHistory) (*types.Any, error) {
	switch constants.PeerTypeFromInt32(readHistory.PeerType) {
	case constants.PeerTypeSelf, constants.PeerTypeUser:
		err := m.coreService.ReadHistoryMessage(readHistory)
		if err != nil {
			log.Errorf("ReadHistory peerType:%+v error:%s", readHistory.PeerType, err.Error())
			return nil, err
		}
	case constants.PeerTypeChat:
		err := m.coreService.ReadHistoryMessage(readHistory)
		if err != nil {
			log.Errorf("ReadHistory peerType:%+v error:%s", readHistory.PeerType, err.Error())
			return nil, err
		}
	case constants.PeerTypeChannel:
		fallthrough
	default:
		log.Warnf("ReadHistory peerType:%+v", readHistory.PeerType)
		return types.MarshalAny(mtproto.ToMTBool(false))
	}
	return types.MarshalAny(mtproto.ToMTBool(true))
}

func (m *MsgHandler) RevokeMessage(revokeMessage *msgService.RevokeMessages) (*types.Any, error) {
	switch constants.PeerTypeFromInt32(revokeMessage.PeerType) {
	case constants.PeerTypeSelf, constants.PeerTypeUser, constants.PeerTypeChat, constants.PeerTypeChannel:
		err := m.coreService.RevokeMessage(revokeMessage)
		if err != nil {
			log.Errorf("RevokeMessage peerType:%+v error:%s", revokeMessage.PeerType, err.Error())
			return nil, err
		}
	default:
		log.Warnf("RevokeMessage peerType:%+v", revokeMessage.PeerType)
		return types.MarshalAny(mtproto.ToMTBool(false))
	}
	return types.MarshalAny(mtproto.ToMTBool(true))
}

func (m *MsgHandler) PinnedMessage(message *msgService.PinnedMessage) (*types.Any, error) {
	switch constants.PeerTypeFromInt32(message.PeerType) {
	case constants.PeerTypeSelf, constants.PeerTypeUser, constants.PeerTypeChat, constants.PeerTypeChannel:
		return m.coreService.PinnedMessage(message.AuthKeyId, message.UserId, message.PeerId, constants.PeerTypeFromInt32(message.PeerType), message.MessageId, message.Unpin, message.OneSide, 0)
	default:
		log.Errorf("PinnedMessage peerType:%+v", message.PeerType)
		return nil, fmt.Errorf("PinnedMessage peerType:%+v", message.PeerType)
	}
}
