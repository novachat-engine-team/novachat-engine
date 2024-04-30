/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.getLeftChannels_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  channels.getLeftChannels#8341ecc0 offset:int = messages.Chats;
//
func (s *ChannelsServiceImpl) ChannelsGetLeftChannels(ctx context.Context, request *mtproto.TLChannelsGetLeftChannels) (*mtproto.Messages_Chats, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsGetLeftChannels %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	return mtproto.NewTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
		Count: 0,
	}).To_Messages_Chats(), nil
}
