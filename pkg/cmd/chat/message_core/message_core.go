/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/11 10:48
 * @File : message_core.go
 */

package message_core

import (
	"context"
	"errors"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
)

func (m *MessageCoreService) SendMessageData(r *msgService.SendMessages) error {

	switch constants.PeerTypeFromInt32(r.PeerType) {
	case constants.PeerTypeChat, constants.PeerTypeChannel:
		break
	default:
		log.Errorf("SendMessageData peerType:%+v", r.PeerType)
		return nil
	}

	_, err := chatService.GetChatClientByKeyId(r.PeerId).ReqSendInBoxesMessages(context.Background(),
		&chatService.SendInBoxesMessages{
			Message: r,
		})
	if err != nil {
		log.Errorf("ReqSendInBoxesMessages userId:%d chatId:%d error:%s", r.FromUserId, r.PeerId)
	}
	return err
}

func (m *MessageCoreService) ReadHistoryMessageData(r *msgService.ReadHistory) error {
	switch constants.PeerTypeFromInt32(r.PeerType) {
	case constants.PeerTypeChat:
		break
	default:
		log.Errorf("ReadHistoryMessageData peerType:%+v", r.PeerType)
		return nil
	}

	_, err := chatService.GetChatClientByKeyId(r.PeerId).ReqReadInMessages(context.Background(),
		&chatService.ReadInMessages{
			Value: r,
		})
	if errors.Is(err, mtproto.DefaultRpcError) {
		log.Warnf("ReadHistoryMessageData userId:%d chatId:%d error:%s", r.UserId, r.PeerId)
		err = nil
	}
	if err != nil {
		log.Errorf("ReadHistoryMessageData userId:%d chatId:%d error:%s", r.UserId, r.PeerId)
	}
	return err
}

func (m *MessageCoreService) PinnedMessageData(r *msgService.PinnedMessage) error {
	switch constants.PeerTypeFromInt32(r.PeerType) {
	case constants.PeerTypeChat, constants.PeerTypeChannel:
		break
	default:
		log.Errorf("PinnedMessageData peerType:%+v", r.PeerType)
		return nil
	}

	_, err := chatService.GetChatClientByKeyId(r.PeerId).ReqInPinnedMessage(context.Background(),
		&chatService.ChatInPinnedMessage{
			PinnedMessage: r,
		})
	if errors.Is(err, mtproto.DefaultRpcError) {
		log.Warnf("PinnedMessageData userId:%d chatId:%d error:%s", r.UserId, r.PeerId)
		err = nil
	}
	if err != nil {
		log.Errorf("PinnedMessageData userId:%d chatId:%d error:%s", r.UserId, r.PeerId)
	}
	return err
}
