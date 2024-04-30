/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.toggleSlowMode_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  channels.toggleSlowMode#edd49ef0 channel:InputChannel seconds:int = Updates;
//
func (s *ChannelsServiceImpl) ChannelsToggleSlowMode(ctx context.Context, request *mtproto.TLChannelsToggleSlowMode) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsToggleSlowMode %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	return mtproto.NewTLUpdates(nil).To_Updates(), nil
}
