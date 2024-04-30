/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : stickers.removeStickerFromSet_handler.go
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

//  stickers.removeStickerFromSet#f7760f51 sticker:InputDocument = messages.StickerSet;
//
func (s *StickersServiceImpl) StickersRemoveStickerFromSet(ctx context.Context, request *mtproto.TLStickersRemoveStickerFromSet) (*mtproto.Messages_StickerSet, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("StickersRemoveStickerFromSet %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl StickersRemoveStickerFromSet logic

	return nil, fmt.Errorf("%s", "Not impl StickersRemoveStickerFromSet")
}
