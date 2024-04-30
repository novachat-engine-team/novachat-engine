/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/09/22 10:25
 * @File : channels.getDialogs_handler.go
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

//  channels.getDialogs#a9d3d249 offset:int limit:int = messages.Dialogs;
//
func (s *ChannelsServiceImpl) ChannelsGetDialogs(ctx context.Context, request *mtproto.TLChannelsGetDialogs) (*mtproto.Messages_Dialogs, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsGetDialogs %v, request: %v", md, request)

	// Impl ChannelsGetDialogs logic

	return nil, fmt.Errorf("%s", "Not impl ChannelsGetDialogs")
}
