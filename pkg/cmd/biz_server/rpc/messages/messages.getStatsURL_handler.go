/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getStatsURL_handler.go
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

//  messages.getStatsURL#812c2ae6 flags:# dark:flags.0?true peer:InputPeer params:string = StatsURL;
//
func (s *MessagesServiceImpl) MessagesGetStatsURL(ctx context.Context, request *mtproto.TLMessagesGetStatsURL) (*mtproto.StatsURL, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetStatsURL %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesGetStatsURL logic

	return nil, fmt.Errorf("%s", "Not impl MessagesGetStatsURL")
}
