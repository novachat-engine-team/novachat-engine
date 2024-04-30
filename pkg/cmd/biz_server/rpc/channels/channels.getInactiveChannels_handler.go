/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.getInactiveChannels_handler.go
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

//  channels.getInactiveChannels#11e831ee = messages.InactiveChats;
//
func (s *ChannelsServiceImpl) ChannelsGetInactiveChannels(ctx context.Context, request *mtproto.TLChannelsGetInactiveChannels) (*mtproto.Messages_InactiveChats, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ChannelsGetInactiveChannels %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	inactiveChats := mtproto.NewTLMessagesInactiveChats(&mtproto.Messages_InactiveChats{
		Dates: []int32{},
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	})
	return inactiveChats.To_Messages_InactiveChats(), nil
}
