/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : stickers.addStickerToSet_handler.go
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

//  stickers.addStickerToSet#8653febe stickerset:InputStickerSet sticker:InputStickerSetItem = messages.StickerSet;
//
func (s *StickersServiceImpl) StickersAddStickerToSet(ctx context.Context, request *mtproto.TLStickersAddStickerToSet) (*mtproto.Messages_StickerSet, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("StickersAddStickerToSet %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl StickersAddStickerToSet logic

	return nil, fmt.Errorf("%s", "Not impl StickersAddStickerToSet")
}
