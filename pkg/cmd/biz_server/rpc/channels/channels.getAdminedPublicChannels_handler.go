/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.getAdminedPublicChannels_handler.go
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

//  channels.getAdminedPublicChannels#f8b036af flags:# by_location:flags.0?true check_limit:flags.1?true = messages.Chats;
//  channels.getAdminedPublicChannels#8d8d82d7 = messages.Chats;
//
func (s *ChannelsServiceImpl) ChannelsGetAdminedPublicChannels(ctx context.Context, request *mtproto.TLChannelsGetAdminedPublicChannels) (*mtproto.Messages_Chats, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ChannelsGetAdminedPublicChannels %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	messagesChats := mtproto.NewTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
		Count: 0,
	})
	return messagesChats.To_Messages_Chats(), nil
}
