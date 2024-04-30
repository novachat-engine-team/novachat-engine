/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getStickerSet_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/util"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/core/message/message"
	data_stickerset "novachat_engine/service/data/stickerset"
	"novachat_engine/service/upload/photo"
)

//  messages.getStickerSet#2619a90e stickerset:InputStickerSet = messages.StickerSet;
//
func (s *MessagesServiceImpl) MessagesGetStickerSet(ctx context.Context, request *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetStickerSet %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	var sticker *data_stickerset.StickerSet
	reply := mtproto.NewTLMessagesStickerSet(&mtproto.Messages_StickerSet{})
	inputSticker := request.GetStickerset()
	switch inputSticker.ClassName {

	//  + TL_inputStickerSetEmpty
	//  + TL_inputStickerSetID
	//  + TL_inputStickerSetShortName
	//  + TL_inputStickerSetAnimatedEmoji
	//  + TL_inputStickerSetDice
	case mtproto.ClassInputStickerSetID:
		stickerSetList, err := s.accountStickerSetCore.GetStickers([]int64{request.Stickerset.Id})
		if err != nil {
			log.Errorf("MessagesGetStickerSet %v, request: %v error:%s", md, request, err.Error())
			return nil, err
		}
		if len(stickerSetList) == 0 {
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_STICKER_ID_INVALID)
		}
		sticker = stickerSetList[0]
	case mtproto.ClassInputStickerSetAnimatedEmoji:
		emojiId := int64(1258816259751996)
		stickerSetList, err := s.accountStickerSetCore.GetStickers([]int64{emojiId})
		if err != nil {
			log.Errorf("MessagesGetStickerSet %v, request: %v error:%s", md, request, err.Error())
			return nil, err
		}
		if len(stickerSetList) == 0 {
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_STICKER_ID_INVALID)
		}
		sticker = stickerSetList[0]

	case mtproto.ClassInputStickerSetDice:
		log.Debugf("MessagesGetStickerSet %v, emoticon:%v", metadata.RpcMetaDataDebug(md), request.Stickerset.Emoticon)
		reply.SetSet(mtproto.NewTLStickerSet(&mtproto.StickerSet{
			Archived:      false,
			Official:      false,
			Masks:         false,
			Animated:      false,
			InstalledDate: 0,
			Id:            0,
			AccessHash:    0,
			Title:         "",
			ShortName:     "",
			Thumbs:        nil,
			ThumbDcId:     0,
			Count:         0,
			Hash:          0,
			Installed:     false,
		}).To_StickerSet())
		log.Debugf("MessagesGetStickerSet %v, request: %v ret:%+v", metadata.RpcMetaDataDebug(md), request)
		return reply.To_Messages_StickerSet(), nil

	case mtproto.ClassInputStickerSetShortName:
		if len(inputSticker.GetShortName()) == 0 {
			err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_STICKER_ID_INVALID)
			return nil, err
		}
		stickerSet, err := s.accountStickerSetCore.GetStickersByShortName(inputSticker.GetShortName())
		if err != nil {
			log.Errorf("MessagesGetStickerSet %v, request: %v error:%s", md, request, err.Error())
			return nil, err
		}
		if stickerSet == nil {
			err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_STICKER_ID_INVALID)
			return nil, err
		}
		sticker = stickerSet
	default:
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_STICKERSET_INVALID)
		log.Errorf("MessagesGetStickerSet %v, request: %v error:%s", md, request, err.Error())
		return nil, err
	}

	installList, _ := s.accountStickerSetCore.GetInstalled(md.UserId)
	var install *data_stickerset.StickerInstall
	if idx := util.Index(&installList, func(idx int) bool {
		return installList[idx].StickerSetId == sticker.Id
	}); idx >= 0 {
		install = installList[idx]
	}

	var packs = make([]*mtproto.StickerPack, 0, len(sticker.Packs))
	var documents = make([]*mtproto.Document, 0, len(sticker.Packs))
	for _, pack := range sticker.Packs {
		documentIdList := make([]int64, 0, len(pack.Documents))
		for _, document := range pack.Documents {
			documents = append(documents, message.ToMessageDocument(document, md.Layer))
			documentIdList = append(documentIdList, document.VolumeId)
		}
		packs = append(packs, mtproto.NewTLStickerPack(&mtproto.StickerPack{
			Emoticon:  pack.Emoticon,
			Documents: documentIdList,
		}).To_StickerPack())
	}

	photoSizes := make([]*mtproto.PhotoSize, 0, len(sticker.Thumbs))
	for _, thumbs := range sticker.Thumbs {
		for _, thumb := range thumbs.Thumbs {
			photoSizes = append(photoSizes, photo.PhotoData2PhotoSize(thumb, md.Layer))
		}
	}

	var photoSize *mtproto.PhotoSize
	if len(photoSizes) > 0 {
		photoSize = photoSizes[0]
	}

	reply.SetSet(mtproto.NewTLStickerSet(&mtproto.StickerSet{
		Installed:     false,
		Archived:      false,
		Official:      sticker.Official,
		Masks:         sticker.Masks,
		Id:            sticker.Id,
		AccessHash:    sticker.AccessHash,
		Title:         sticker.Title,
		ShortName:     sticker.ShortName,
		Count:         sticker.Count,
		Hash:          sticker.Hash,
		InstalledDate: 0,
		Thumb:         photoSize,
		ThumbDcId:     sticker.ThumbDcId,
		Animated:      sticker.Animated,
		Thumbs:        photoSizes,
	}).To_StickerSet())

	if install != nil {
		reply.GetSet().Installed = install.Installed
		reply.GetSet().InstalledDate = install.Date
	}

	reply.SetPacks(packs)
	reply.SetDocuments(documents)
	return reply.To_Messages_StickerSet(), nil
}
