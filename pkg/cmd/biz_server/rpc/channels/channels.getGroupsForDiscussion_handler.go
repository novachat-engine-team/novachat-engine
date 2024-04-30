/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.getGroupsForDiscussion_handler.go
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

//  channels.getGroupsForDiscussion#f5dad378 = messages.Chats;
//
func (s *ChannelsServiceImpl) ChannelsGetGroupsForDiscussion(ctx context.Context, request *mtproto.TLChannelsGetGroupsForDiscussion) (*mtproto.Messages_Chats, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ChannelsGetGroupsForDiscussion %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	messagesChats := mtproto.NewTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
		Count: 0,
	})
	return messagesChats.To_Messages_Chats(), nil
}
