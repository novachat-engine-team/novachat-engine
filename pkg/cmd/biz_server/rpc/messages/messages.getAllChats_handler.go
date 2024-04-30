/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getAllChats_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/chat"
	"novachat_engine/service/constants"
)

//  messages.getAllChats#eba80ff0 except_ids:Vector<int> = messages.Chats;
//
func (s *MessagesServiceImpl) MessagesGetAllChats(ctx context.Context, request *mtproto.TLMessagesGetAllChats) (*mtproto.Messages_Chats, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetAllChats %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	exceptIds := make([]int64, 0, len(request.ExceptIds))
	for _, v := range request.ExceptIds {
		exceptIds = append(exceptIds, constants.PeerTypeFromChatIDType32(v).ToInt())
	}
	resp, err := chatService.GetChatClientByKeyId(md.AuthKeyId).ReqAllChat(ctx, &chatService.AllChat{
		ExceptIds: exceptIds,
		Ids:       nil,
		UserId:    md.UserId,
	})
	if err != nil {
		log.Errorf("MessagesGetAllChats %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var chats = make([]*mtproto.Chat, 0, len(resp.Values))
	for _, v := range resp.Values {
		if v.ChatData.Deleted == true {
			chats = append(chats, mtproto.NewTLChatForbidden(&mtproto.Chat{
				Id:    constants.PeerTypeFromChatIDType(v.ChatData.ChatId).ToInt32(),
				Title: v.ChatData.Title,
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
