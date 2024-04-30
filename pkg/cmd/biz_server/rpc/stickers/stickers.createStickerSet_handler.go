/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : stickers.createStickerSet_handler.go
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

//  stickers.createStickerSet#f1036780 flags:# masks:flags.0?true animated:flags.1?true user_id:InputUser title:string short_name:string thumb:flags.2?InputDocument stickers:Vector<InputStickerSetItem> = messages.StickerSet;
//  stickers.createStickerSet#9bd86e6a flags:# masks:flags.0?true user_id:InputUser title:string short_name:string stickers:Vector<InputStickerSetItem> = messages.StickerSet;
//
func (s *StickersServiceImpl) StickersCreateStickerSet(ctx context.Context, request *mtproto.TLStickersCreateStickerSet) (*mtproto.Messages_StickerSet, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("StickersCreateStickerSet %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl StickersCreateStickerSet logic

	return nil, fmt.Errorf("%s", "Not impl StickersCreateStickerSet")
}
