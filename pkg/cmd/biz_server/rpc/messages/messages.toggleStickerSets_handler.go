/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.toggleStickerSets_handler.go
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

//  messages.toggleStickerSets#b5052fea flags:# uninstall:flags.0?true archive:flags.1?true unarchive:flags.2?true stickersets:Vector<InputStickerSet> = Bool;
//
func (s *MessagesServiceImpl) MessagesToggleStickerSets(ctx context.Context, request *mtproto.TLMessagesToggleStickerSets) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesToggleStickerSets %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesToggleStickerSets logic

	return mtproto.ToMTBool(true), nil
}
