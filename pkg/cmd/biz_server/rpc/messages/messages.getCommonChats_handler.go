/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getCommonChats_handler.go
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

//  messages.getCommonChats#d0a48c4 user_id:InputUser max_id:int limit:int = messages.Chats;
//
func (s *MessagesServiceImpl) MessagesGetCommonChats(ctx context.Context, request *mtproto.TLMessagesGetCommonChats) (*mtproto.Messages_Chats, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetCommonChats %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatList, err := chatService.GetChatClientByKeyId(md.UserId).ReqChatCommon(ctx, &chatService.ChatCommon{
		UserId: constants.PeerTypeFromUserIDType32(request.UserId.UserId).ToInt(),
		MaxId:  request.MaxId,
		Limit:  request.Limit,
	})
	if err != nil {
		log.Errorf("MessagesGetCommonChats %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	if err != nil {
		log.Fatala("MessagesGetCommonChats %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	}

	chats := make([]*mtproto.Chat, 0, len(chatList.Values))
	for _, v := range chatList.Values {
		chats = append(chats, chat.ToChat(md.UserId, v, md.Layer))
	}

	return mtproto.NewTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
		Count: int32(len(chatList.Values)),
	}).To_Messages_Chats(), nil
}
