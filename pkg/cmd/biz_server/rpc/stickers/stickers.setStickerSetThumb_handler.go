/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : stickers.setStickerSetThumb_handler.go
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

//  stickers.setStickerSetThumb#9a364e30 stickerset:InputStickerSet thumb:InputDocument = messages.StickerSet;
//
func (s *StickersServiceImpl) StickersSetStickerSetThumb(ctx context.Context, request *mtproto.TLStickersSetStickerSetThumb) (*mtproto.Messages_StickerSet, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("StickersSetStickerSetThumb %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl StickersSetStickerSetThumb logic

	return nil, fmt.Errorf("%s", "Not impl StickersSetStickerSetThumb")
}
