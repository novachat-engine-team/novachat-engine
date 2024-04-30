/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.setStickers_handler.go
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

//  channels.setStickers#ea8ca4f9 channel:InputChannel stickerset:InputStickerSet = Bool;
//
func (s *ChannelsServiceImpl) ChannelsSetStickers(ctx context.Context, request *mtproto.TLChannelsSetStickers) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ChannelsSetStickers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	return mtproto.ToMTBool(false), nil
}
