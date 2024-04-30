/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getFavedStickers_handler.go
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
	"novachat_engine/service/core/message/message"
)

//  messages.getFavedStickers#21ce0b0e hash:int = messages.FavedStickers;
//
func (s *MessagesServiceImpl) MessagesGetFavedStickers(ctx context.Context, request *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetFavedStickers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	val, err := s.accountStickerSetCore.GetFaved(md.UserId)
	if err != nil {
		log.Errorf("MessagesGetFavedStickers %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	result := mtproto.NewTLMessagesFavedStickers(&mtproto.Messages_FavedStickers{
		Hash:     request.Hash,
		Packs:    make([]*mtproto.StickerPack, 0, len(val)),
		Stickers: nil,
	})

	documentList := make([]*mtproto.Document, 0, len(val)*10)
	for _, v := range val {
		for _, pack := range v.Packs {
			documentIdList := make([]int64, 0, len(pack.Documents))
			for _, document := range pack.Documents {
				documentList = append(documentList, message.ToMessageDocument(document, md.Layer))
				documentIdList = append(documentIdList, document.VolumeId)
			}

			result.Data2.Packs = append(result.Data2.Packs, mtproto.NewTLStickerPack(&mtproto.StickerPack{
				Emoticon:  pack.Emoticon,
				Documents: documentIdList,
			}).To_StickerPack())
		}
	}

	var hashValue int32
	hashValue = hash.DocumentsHash(documentList, hash.MaxCountHash)
	if request.Hash != 0 && hashValue == request.Hash {
		log.Infof("MessagesGetFavedStickers %v, request: %v NewTLMessagesFavedStickersNotModified ok", metadata.RpcMetaDataDebug(md), request)
		return mtproto.NewTLMessagesFavedStickersNotModified(&mtproto.Messages_FavedStickers{}).To_Messages_FavedStickers(), nil
	}

	result.SetHash(hashValue)
	result.SetStickers(documentList)
	log.Infof("MessagesGetRecentStickers %v, request: %v ok", metadata.RpcMetaDataDebug(md), request)

	return result.To_Messages_FavedStickers(), nil
}
