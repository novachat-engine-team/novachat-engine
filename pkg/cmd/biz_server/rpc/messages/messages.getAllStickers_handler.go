/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getAllStickers_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/hash"
)

//  messages.getAllStickers#1c9618b1 hash:int = messages.AllStickers;
//
func (s *MessagesServiceImpl) MessagesGetAllStickers(ctx context.Context, request *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetAllStickers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//TODO:(Coder) Impl MessagesGetAllStickers logic
	stickerSetAll, err := s.accountStickerSetCore.AllStickersByUser(md.UserId, md.Layer)
	stickers := mtproto.NewTLMessagesAllStickers(nil)
	stickers.SetHash(request.Hash)
	if err != nil {
		log.Fatalf("MessagesGetAllStickers %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	}

	stickers.SetHash(hash.SetHash(stickerSetAll))
	stickers.SetSets(stickerSetAll)
	log.Debugf("MessagesGetAllStickers %v, request: %v reply ok!!!!!!!!!", md, request)
	return stickers.To_Messages_AllStickers(), nil
}
