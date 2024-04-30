/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : stickers.changeStickerPosition_handler.go
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

//  stickers.changeStickerPosition#ffb6d4ca sticker:InputDocument position:int = messages.StickerSet;
//
func (s *StickersServiceImpl) StickersChangeStickerPosition(ctx context.Context, request *mtproto.TLStickersChangeStickerPosition) (*mtproto.Messages_StickerSet, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("StickersChangeStickerPosition %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl StickersChangeStickerPosition logic

	return nil, fmt.Errorf("%s", "Not impl StickersChangeStickerPosition")
}
