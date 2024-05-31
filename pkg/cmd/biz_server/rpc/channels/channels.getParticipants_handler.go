/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.getParticipants_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/util"
	"novachat_engine/service/chat"
	"novachat_engine/service/constants"
)

//  channels.getParticipants#123e05e9 channel:InputChannel filter:ChannelParticipantsFilter offset:int limit:int hash:int = channels.ChannelParticipants;
//  channels.getParticipants#24d98f92 channel:InputChannel filter:ChannelParticipantsFilter offset:int limit:int = channels.ChannelParticipants;
//
func (s *ChannelsServiceImpl) ChannelsGetParticipants(ctx context.Context, request *mtproto.TLChannelsGetParticipants) (*mtproto.Channels_ChannelParticipants, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsGetParticipants %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	filter := constants.MTToChannelParticipantFilterType(request.Filter).ToInt32()
	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqGetParticipants(ctx, &chatService.GetParticipants{
		ChatId: chatId,
		Filter: filter,
		Q:      request.Filter.Q,
		Offset: request.Offset,
		Limit:  request.Limit,
	})
	if err != nil {
		log.Errorf("ChannelsGetParticipants %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var chatInfo chatService.Chat
	err = types.UnmarshalAny(resp, &chatInfo)
	if err != nil {
		log.Errorf("ChannelsGetParticipants %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	userIdList := make([]int64, 0, len(chatInfo.ParticipantList))
	p := mtproto.NewTLChannelsChannelParticipants(&mtproto.Channels_ChannelParticipants{
		Count:        int32(len(chatInfo.ParticipantList)),
		Participants: make([]*mtproto.ChannelParticipant, 0, len(chatInfo.ParticipantList)),
		Users:        nil,
	}).To_Channels_ChannelParticipants()
	util.Foreach(chatInfo.ParticipantList, func(first interface{}, second interface{}) {
		participant := chatInfo.ParticipantList[first.(int)]
		p.Participants = append(p.Participants, chat.ToChatParticipant(participant, chatInfo.ChatData.Creator, md.UserId))

		if participant != nil {
			userIdList = append(userIdList, participant.UserId)
		}
	})

	users, _ := s.accountUsersCore.GetUserList(md.UserId, userIdList, md.Layer)
	p.Users = users
	return p, nil
}
