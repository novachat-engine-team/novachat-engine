/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getMaskStickers_handler.go
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
	data_stickerset "novachat_engine/service/data/stickerset"
	"novachat_engine/service/upload/photo"
)

//  messages.getMaskStickers#65b8c79f hash:int = messages.AllStickers;
//
func (s *MessagesServiceImpl) MessagesGetMaskStickers(ctx context.Context, request *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetMaskStickers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	stickers := mtproto.NewTLMessagesAllStickers(nil)
	stickers.SetHash(request.Hash)

	stickerSetAll, err := s.accountStickerSetCore.GetAllStickers()
	if err != nil {
		log.Fatalf("MessagesGetMaskStickers %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	}
	installList, err := s.accountStickerSetCore.GetInstalled(md.UserId)
	if err != nil {
		log.Warnf("MessagesGetMaskStickers request: %v GetInstalled error:%s", request, err.Error())
	}

	defaultInstall := data_stickerset.StickerInstall{}
	sets := make([]*mtproto.StickerSet, 0, len(stickerSetAll))
	for _, v := range stickerSetAll {
		if v.Masks == false {
			continue
		}
		var install *data_stickerset.StickerInstall
		if idx := util.Index(&installList, func(idx int) bool {
			return installList[idx].StickerSetId == v.Id
		}); idx >= 0 {
			install = installList[idx]
		}
		if install == nil {
			install = &defaultInstall
		}
		photoSizes := make([]*mtproto.PhotoSize, 0, len(v.Thumbs))
		for _, thumbs := range v.Thumbs {
			for _, thumb := range thumbs.Thumbs {
				photoSizes = append(photoSizes, photo.PhotoData2PhotoSize(thumb, md.Layer))
			}
		}
		var photoSize *mtproto.PhotoSize
		if len(photoSizes) > 0 {
			photoSize = photoSizes[0]
		}
		sets = append(sets, mtproto.NewTLStickerSet(&mtproto.StickerSet{
			Installed:     install.Installed,
			Archived:      install.Archived,
			Official:      v.Official,
			Masks:         v.Masks,
			Id:            v.Id,
			AccessHash:    v.AccessHash,
			Title:         v.Title,
			ShortName:     v.ShortName,
			Count:         v.Count,
			Hash:          v.Hash,
			InstalledDate: install.Date,
			Thumb:         photoSize,
			ThumbDcId:     v.ThumbDcId,
			Animated:      v.Animated,
			Thumbs:        photoSizes,
		}).To_StickerSet())
	}
	stickers.SetSets(sets)
	return stickers.To_Messages_AllStickers(), nil
}
