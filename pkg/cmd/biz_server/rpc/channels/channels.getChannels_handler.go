/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.getChannels_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"math"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/chat"
	"novachat_engine/service/constants"
)

//  channels.getChannels#a7f6bbb id:Vector<InputChannel> = messages.Chats;
//
func (s *ChannelsServiceImpl) ChannelsGetChannels(ctx context.Context, request *mtproto.TLChannelsGetChannels) (*mtproto.Messages_Chats, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsGetChannels %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.Id) == 0 {
		return mtproto.NewTLMessagesChats(&mtproto.Messages_Chats{}).To_Messages_Chats(), nil
	}

	ids := make([]int64, 0, len(request.Id))
	for _, v := range request.Id {
		ids = append(ids, constants.PeerTypeFromChatIDType32(v.ChannelId).ToInt())
	}
	resp, err := chatService.GetChatClientByKeyId(md.AuthKeyId).ReqAllChat(ctx, &chatService.AllChat{
		ExceptIds: nil,
		Ids:       ids,
		UserId:    md.UserId,
	})
	if err != nil {
		log.Errorf("ChannelsGetChannels %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var chats = make([]*mtproto.Chat, 0, len(resp.Values))
	for _, v := range resp.Values {
		if v.ChatData.Deleted == true {
			chats = append(chats, mtproto.NewTLChatForbidden(&mtproto.Chat{
				Id:         constants.PeerTypeFromChatIDType(v.ChatData.ChatId).ToInt32(),
				Title:      v.ChatData.Title,
				Broadcast:  v.ChatData.Broadcast,
				Megagroup:  true,
				AccessHash: v.ChatData.AccessHash,
				UntilDate:  math.MaxInt32,
			}).To_Chat())
			continue
		}

		chats = append(chats, chat.ToChat(md.UserId, v, md.Layer))
	}

	return mtproto.NewTLMessagesChats(&mtproto.Messages_Chats{
		Chats: chats,
		Count: int32(len(chats)),
	}).To_Messages_Chats(), nil
}
