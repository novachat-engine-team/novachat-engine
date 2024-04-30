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
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/common/hash"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
	fs "novachat_engine/service/data/fs"
)

func (impl *Impl) ReqCreateChat(ctx context.Context, request *chatService.CreateChat) (*types.Any, error) {
	log.Debugf("ReqCreateChat request:%+v", request)

	var geoPoint *fs.GeoPoint
	var address string
	if request.GeoPoint != nil {
		address = request.GeoPoint.Address
		geoPoint = &fs.GeoPoint{
			Long:           request.GeoPoint.Long,
			Lat:            request.GeoPoint.Lat,
			AccessHash:     0,
			AccuracyRadius: 0,
		}
	}
	chat, err := impl.chatManager.CreateChat(
		request.UserId,
		request.PeerId,
		request.Title,
		request.Date,
		geoPoint,
		address)
	if err != nil {
		log.Errorf("ReqCreateChat request:%+v", request, err.Error())
		return nil, err
	}

	chatInfo := chat.GetChatInfo()
	participantsList := make([]int32, 0, len(request.PeerId)+1)
	participantsInt64List := make([]int64, 0, len(request.PeerId)+1)
	chatInfo.Iteration(func(value *data_chat.ChatParticipant) {
		participantsList = append(participantsList, constants.PeerTypeFromUserIDType(value.UserId).ToInt32())
		participantsInt64List = append(participantsInt64List, value.UserId)
	})

	participantsInt64List = append(participantsInt64List, chatInfo.ChatData.ChatId)

	chatPeer := mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChatIDType(chatInfo.ChatData.ChatId).ToInt32()}).To_Peer()
	message := mtproto.NewTLMessageService(&mtproto.Message{
		Date:              request.Date,
		FromId286FA604119: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(request.UserId).ToInt32()}).To_Peer(),
		FromId90DDDC1171:  constants.PeerTypeFromUserIDType(request.UserId).ToInt32(),
		ToId:              chatPeer,
		PeerId:            chatPeer,
		Action: mtproto.NewTLMessageActionChannelCreate(&mtproto.MessageAction{
			Title: request.Title,
			Users: participantsList,
		}).To_MessageAction(),
	})

	updates, err := msgService.GetMsgClient().
		ReqSendMessages(ctx, &msgService.SendMessages{
			AuthKeyId:  request.AuthKeyId,
			FromUserId: request.UserId,
			PeerId:     chatInfo.ChatData.ChatId,
			PeerType:   constants.PeerTypeChat.ToInt32(),
			DataList: []*msgService.SendMessageData{{
				RandomId: hash.HashId64ListNew(participantsInt64List),
				Message:  message.To_Message(),
			}},
		})
	if err != nil {
		log.Errorf("ReqCreateChat send create message request:%+v", request, err.Error())
		return nil, err
	}

	log.Infof("ReqCreateChat request:%+v chatId:%d", request, chatInfo.ChatData.ChatId)
	return types.MarshalAny(updates)
}
