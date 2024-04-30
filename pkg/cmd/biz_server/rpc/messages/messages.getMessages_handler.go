/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getMessages_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
)

//  messages.getMessages#63c66506 id:Vector<InputMessage> = messages.Messages;
//  messages.getMessages#4222fa74 id:Vector<int> = messages.Messages;
//
func (s *MessagesServiceImpl) MessagesGetMessages(ctx context.Context, request *mtproto.TLMessagesGetMessages) (*mtproto.Messages_Messages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetMessages %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	inputMessageType := constants.InputMessageID
	var messageIdList []int32
	if len(request.Id4222FA7471) > 0 {
		inputMessageType = constants.InputMessageID
		messageIdList = request.Id4222FA7471

	} else if len(request.Id4222FA7471) == 0 && len(request.Id63C6650685) == 0 {
		return mtproto.NewTLMessagesMessages(nil).To_Messages_Messages(), nil
	} else {
		messageIdList = make([]int32, 0, len(request.Id63C6650685))
		for _, v := range request.Id63C6650685 {
			switch v.ClassName {
			case mtproto.ClassInputMessageID:
				inputMessageType = constants.InputMessageID
				messageIdList = append(messageIdList, v.Id)
			case mtproto.ClassInputMessagePinned:
				inputMessageType = constants.InputMessagePinned
			case mtproto.ClassInputMessageReplyTo:
				inputMessageType = constants.InputMessageReplyTo
				messageIdList = append(messageIdList, v.Id)
			case mtproto.ClassInputMessageCallbackQuery:
				inputMessageType = constants.InputMessageIDCallbackQuery
			}
		}
	}

	if inputMessageType != constants.InputMessageIDCallbackQuery {
		return mtproto.NewTLMessagesMessages(nil).To_Messages_Messages(), nil
	}

	messages, err := s.accountMessageCore.GetMessages(md.UserId, messageIdList, inputMessageType == constants.InputMessageReplyTo)
	if err != nil {
		log.Errorf("MessagesGetMessages - request: %v GetMessages error:%s", request, err.Error())
		return nil, err
	}
	var inputPeer *input.InputPeer
	userIdMap := make(map[int64]struct{}, len(messages))
	for _, v := range messages {
		_ = v.FromId90DDDC1171
		_ = v.PeerId
		if !(v.FromId90DDDC1171 == 0 ||
			constants.PeerTypeFromUserIDType32(v.FromId90DDDC1171).ToInt() == md.UserId) {
			userIdMap[constants.PeerTypeFromUserIDType32(v.FromId90DDDC1171).ToInt()] = struct{}{}
		}

		if v.FromId286FA604119 != nil && !(v.FromId286FA604119.UserId == 0 ||
			constants.PeerTypeFromUserIDType32(v.FromId286FA604119.UserId).ToInt() == md.UserId) {
			userIdMap[constants.PeerTypeFromUserIDType32(v.FromId286FA604119.UserId).ToInt()] = struct{}{}
		}

		inputPeer = input.MakePeer(v.PeerId)
		if inputPeer.GetPeerType() == constants.PeerTypeUser &&
			(inputPeer.GetPeerId() != md.UserId && inputPeer.GetPeerId() != 0) {
			userIdMap[inputPeer.GetPeerId()] = struct{}{}
		}
	}

	userIdList := make([]int64, 0, len(userIdMap))
	for k := range userIdMap {
		userIdList = append(userIdList, k)
	}
	userList, err := s.accountUsersCore.GetUserList(md.UserId, userIdList)
	if err != nil {
		log.Warnf("MessagesGetMessages - GetUserList error:%s", err.Error())
	}

	////  messages.messages#8c718e87 messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
	messagesMessages := mtproto.NewTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: messages,
		Chats:    []*mtproto.Chat{},
		Users:    userList,
	})
	return messagesMessages.To_Messages_Messages(), nil
}
