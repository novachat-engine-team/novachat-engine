/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : stats.loadAsyncGraph_handler.go
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

//  stats.loadAsyncGraph#621d5fa0 flags:# token:string x:flags.0?long = StatsGraph;
//
func (s *StatsServiceImpl) StatsLoadAsyncGraph(ctx context.Context, request *mtproto.TLStatsLoadAsyncGraph) (*mtproto.StatsGraph, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("StatsLoadAsyncGraph %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl StatsLoadAsyncGraph logic

	return nil, fmt.Errorf("%s", "Not impl StatsLoadAsyncGraph")
}
