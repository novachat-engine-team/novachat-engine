/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : stats.getMessagePublicForwards_handler.go
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

//  stats.getMessagePublicForwards#5630281b channel:InputChannel msg_id:int offset_rate:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
//
func (s *StatsServiceImpl) StatsGetMessagePublicForwards(ctx context.Context, request *mtproto.TLStatsGetMessagePublicForwards) (*mtproto.Messages_Messages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("StatsGetMessagePublicForwards %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl StatsGetMessagePublicForwards logic

	return nil, fmt.Errorf("%s", "Not impl StatsGetMessagePublicForwards")
}
