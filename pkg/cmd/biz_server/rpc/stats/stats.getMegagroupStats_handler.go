/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : stats.getMegagroupStats_handler.go
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

//  stats.getMegagroupStats#dcdf8607 flags:# dark:flags.0?true channel:InputChannel = stats.MegagroupStats;
//
func (s *StatsServiceImpl) StatsGetMegagroupStats(ctx context.Context, request *mtproto.TLStatsGetMegagroupStats) (*mtproto.Stats_MegagroupStats, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("StatsGetMegagroupStats %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl StatsGetMegagroupStats logic

	return nil, fmt.Errorf("%s", "Not impl StatsGetMegagroupStats")
}
