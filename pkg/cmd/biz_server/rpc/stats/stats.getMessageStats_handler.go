/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : stats.getMessageStats_handler.go
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

//  stats.getMessageStats#b6e0a3f5 flags:# dark:flags.0?true channel:InputChannel msg_id:int = stats.MessageStats;
//
func (s *StatsServiceImpl) StatsGetMessageStats(ctx context.Context, request *mtproto.TLStatsGetMessageStats) (*mtproto.Stats_MessageStats, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("StatsGetMessageStats %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl StatsGetMessageStats logic

	return nil, fmt.Errorf("%s", "Not impl StatsGetMessageStats")
}
