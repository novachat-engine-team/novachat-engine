/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.getParticipant_handler.go
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
	data_chat "novachat_engine/service/data/chat"
)

//  channels.getParticipant#546dd7a6 channel:InputChannel user_id:InputUser = channels.ChannelParticipant;
//
func (s *ChannelsServiceImpl) ChannelsGetParticipant(ctx context.Context, request *mtproto.TLChannelsGetParticipant) (*mtproto.Channels_ChannelParticipant, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsGetParticipant %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if request.UserId.ClassName == mtproto.ClassInputUserSelf {
		request.UserId.UserId = constants.PeerTypeFromUserIDType(md.UserId).ToInt32()
	}
	userId := constants.PeerTypeFromUserIDType32(request.UserId.UserId).ToInt()
	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqGetParticipants(ctx, &chatService.GetParticipants{
		ChatId: chatId,
		UserId: userId,
	})
	if err != nil {
		log.Errorf("ChannelsGetParticipant %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var chatInfo chatService.Chat
	err = types.UnmarshalAny(resp, &chatInfo)
	if err != nil {
		log.Errorf("ChannelsGetParticipant %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var participant *data_chat.ChatParticipant
	if idx := util.Index(chatInfo.ParticipantList, func(idx int) bool {
		return chatInfo.ParticipantList[idx].UserId == userId
	}); idx >= 0 {
		participant = chatInfo.ParticipantList[idx]
	}

	users := make([]*mtproto.User, 0, 1)
	if participant != nil && md.UserId != participant.UserId {
		user, _ := s.accountUsersCore.GetUser(md.UserId, participant.UserId, md.Layer)
		if user != nil {
			users = append(users, user)
		}
	}

	return mtproto.NewTLChannelsChannelParticipant(&mtproto.Channels_ChannelParticipant{
		Participant: chat.ToChatParticipant(participant, chatInfo.ChatData.Creator, md.UserId),
		Users:       users,
	}).To_Channels_ChannelParticipant(), nil
}
