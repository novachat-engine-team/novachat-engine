/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.setBotUpdatesStatus_handler.go
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

//  help.setBotUpdatesStatus#ec22cfcd pending_updates_count:int message:string = Bool;
//
func (s *HelpServiceImpl) HelpSetBotUpdatesStatus(ctx context.Context, request *mtproto.TLHelpSetBotUpdatesStatus) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpSetBotUpdatesStatus %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl HelpSetBotUpdatesStatus logic

	return nil, fmt.Errorf("%s", "Not impl HelpSetBotUpdatesStatus")
}
