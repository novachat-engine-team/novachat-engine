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
	"context"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/cmd/chat/service"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	data_chat "novachat_engine/service/data/chat"
)

func (impl *Impl) ReqAllChat(ctx context.Context, request *chatService.AllChat) (*chatService.ChatList, error) {
	log.Debugf("ReqAllChat request:%+v ", request)
	var chatList []*service.Chat
	var err error
	if len(request.Ids) == 0 && len(request.ExceptIds) == 0 {
		// all chats
		chatList, err = impl.chatManager.GetChats(nil, nil)
	} else {
		idList := make([]int64, 0, len(request.Ids))
		for _, v := range request.Ids {
			if util.IndexInt64s(&request.ExceptIds, v) < 0 {
				idList = append(idList, v)
			}
		}
		chatList, err = impl.chatManager.GetChats(idList, request.ExceptIds)
	}

	if err != nil {
		log.Errorf("ReqAllChat error:%s", err.Error())
		return nil, err
	}

	serviceChatList := make([]*chatService.Chat, 0, len(chatList))
	for _, v := range chatList {
		var participantList []*data_chat.ChatParticipant
		if request.UserId == 0 {
			participantList = make([]*data_chat.ChatParticipant, 0, v.GetChatInfo().Count)
		} else {
			participantList = make([]*data_chat.ChatParticipant, 0, 1)
		}
		v.GetChatInfo().Iteration(func(participant *data_chat.ChatParticipant) {
			if request.UserId == 0 || request.UserId == participant.UserId {
				participantList = append(participantList, participant)
			}
		})
		serviceChatList = append(serviceChatList, &chatService.Chat{
			ChatData:        v.GetChatInfo().ChatData,
			ParticipantList: participantList,
			Count:           v.GetChatInfo().Count,
		})
	}
	log.Infof("ReqAllChat request:%+v serviceChatList:%+v", request, serviceChatList)
	return &chatService.ChatList{Values: serviceChatList}, nil
}
