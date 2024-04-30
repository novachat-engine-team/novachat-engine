/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getStickers_handler.go
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
	"novachat_engine/service/core/message/message"
	"strconv"
)

//  messages.getStickers#43d4f2c emoticon:string hash:int = messages.Stickers;
//
func (s *MessagesServiceImpl) MessagesGetStickers(ctx context.Context, request *mtproto.TLMessagesGetStickers) (*mtproto.Messages_Stickers, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetStickers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.HashAE22E04551) == 0 {
		request.HashAE22E04551 = fmt.Sprintf("%d", request.Hash43D4F2C85)
	} else {
		hash, _ := strconv.ParseInt(request.HashAE22E04551, 10, 32)
		request.Hash43D4F2C85 = int32(hash)
	}
	log.Debugf("MessagesGetStickers %v, emoticon:%s", metadata.RpcMetaDataDebug(md), request.Emoticon)

	stickerSets, err := s.accountStickerSetCore.GetStickersByEmoticon(request.Emoticon)
	if err != nil {
		log.Fatalf("MessagesGetStickers %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	stickers := mtproto.NewTLMessagesStickers(nil)
	stickers.SetHash8A8ECD3271(request.HashAE22E04551)
	stickers.SetHashE4599BBD85(request.Hash43D4F2C85)
	if len(stickerSets) > 0 {
		documentsList := make([]*mtproto.Document, 0, len(stickerSets))
		for _, sticker := range stickerSets {
			for _, v := range sticker.Packs {
				for _, document := range v.Documents {
					documentsList = append(documentsList, message.ToMessageDocument(document, md.Layer))
				}
			}
		}
		stickers.SetStickers(documentsList)
	} else {
		stickers.SetStickers([]*mtproto.Document{})
	}

	log.Debugf("MessagesGetStickers %v, request: %v Emoticon:%s stickerSets:%+v reply ok!!!!!!!!", md, request, request.Emoticon, stickerSets)
	return stickers.To_Messages_Stickers(), nil
}
