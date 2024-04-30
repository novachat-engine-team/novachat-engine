/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.updatePinnedMessage_handler.go
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

//  channels.updatePinnedMessage#a72ded52 flags:# silent:flags.0?true channel:InputChannel id:int = Updates;
//
func (s *ChannelsServiceImpl) ChannelsUpdatePinnedMessage(ctx context.Context, request *mtproto.TLChannelsUpdatePinnedMessage) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsUpdatePinnedMessage %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl ChannelsUpdatePinnedMessage logic

	return nil, fmt.Errorf("%s", "Not impl ChannelsUpdatePinnedMessage")
}
