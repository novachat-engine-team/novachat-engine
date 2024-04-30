/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/09/22 10:25
 * @File : channels.getBroadcastsForDiscussion_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  channels.getBroadcastsForDiscussion#1a87f304 = messages.Chats;
//
func (s *ChannelsServiceImpl) ChannelsGetBroadcastsForDiscussion(ctx context.Context, request *mtproto.TLChannelsGetBroadcastsForDiscussion) (*mtproto.Messages_Chats, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsGetBroadcastsForDiscussion %v, request: %v", md, request)

	// Impl ChannelsGetBroadcastsForDiscussion logic

	return nil, fmt.Errorf("%s", "Not impl ChannelsGetBroadcastsForDiscussion")
}
