/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : stats.getBroadcastStats_handler.go
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

//  stats.getBroadcastStats#ab42441a flags:# dark:flags.0?true channel:InputChannel = stats.BroadcastStats;
//
func (s *StatsServiceImpl) StatsGetBroadcastStats(ctx context.Context, request *mtproto.TLStatsGetBroadcastStats) (*mtproto.Stats_BroadcastStats, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("StatsGetBroadcastStats %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl StatsGetBroadcastStats logic

	return nil, fmt.Errorf("%s", "Not impl StatsGetBroadcastStats")
}
