/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getAttachedStickers_handler.go
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

//  messages.getAttachedStickers#cc5b67cc media:InputStickeredMedia = Vector<StickerSetCovered>;
//
func (s *MessagesServiceImpl) MessagesGetAttachedStickers(ctx context.Context, request *mtproto.TLMessagesGetAttachedStickers) (*mtproto.Vector_StickerSetCovered, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetAttachedStickers %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	panic("MessagesGetAttachedStickers")
}
