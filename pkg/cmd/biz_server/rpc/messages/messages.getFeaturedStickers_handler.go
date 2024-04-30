/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getFeaturedStickers_handler.go
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
	"novachat_engine/service/common/hash"
	"novachat_engine/service/core/message/message"
	data_stickerset "novachat_engine/service/data/stickerset"
	"novachat_engine/service/upload/photo"
)

//  messages.getFeaturedStickers#2dacca4f hash:int = messages.FeaturedStickers;
//
func (s *MessagesServiceImpl) MessagesGetFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetFeaturedStickers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	stickerSetList, err := s.accountStickerSetCore.GetAllStickers()
	if err != nil {
		log.Errorf("MessagesGetAttachedStickers request: %v GetAllStickers error:%s", request, err.Error())
		return nil, err
	}

	installList, err := s.accountStickerSetCore.GetInstalled(md.UserId)
	if err != nil {
		log.Warnf("MessagesGetAttachedStickers request: %v GetInstalled error:%s", request, err.Error())
	}

	defaultInstall := data_stickerset.StickerInstall{}
	idList := make([]int64, 0, len(stickerSetList))
	sets := make([]*mtproto.StickerSetCovered, 0, len(stickerSetList))
	for _, stickerSet := range stickerSetList {
		idList = append(idList, stickerSet.Id)
		var install *data_stickerset.StickerInstall
		if idx := util.Index(&installList, func(idx int) bool {
			return installList[idx].StickerSetId == stickerSet.Id
		}); idx >= 0 {
			install = installList[idx]
		}
		if install == nil {
			install = &defaultInstall
		}

		photoSizes := make([]*mtproto.PhotoSize, 0, len(stickerSet.Thumbs))
		for _, thumbs := range stickerSet.Thumbs {
			for _, thumb := range thumbs.Thumbs {
				photoSizes = append(photoSizes, photo.PhotoData2PhotoSize(thumb, md.Layer))
			}
		}
		var photoSize *mtproto.PhotoSize
		if len(photoSizes) > 0 {
			photoSize = photoSizes[0]
		}

		sets = append(sets, mtproto.NewTLStickerSetCovered(&mtproto.StickerSetCovered{
			Set: mtproto.NewTLStickerSet(&mtproto.StickerSet{
				Installed:     install.Installed,
				Archived:      install.Archived,
				Official:      stickerSet.Official,
				Masks:         stickerSet.Masks,
				Id:            stickerSet.Id,
				AccessHash:    stickerSet.AccessHash,
				Title:         stickerSet.Title,
				ShortName:     stickerSet.ShortName,
				Count:         stickerSet.Count,
				Hash:          stickerSet.Hash,
				InstalledDate: install.Date,
				Thumb:         photoSize,
				ThumbDcId:     stickerSet.ThumbDcId,
				Animated:      stickerSet.Animated,
				Thumbs:        photoSizes,
			}).To_StickerSet(),
			Cover: message.ToMessageDocument(stickerSet.Packs[0].Documents[0], md.Layer),
		}).To_StickerSetCovered())
	}

	result := mtproto.NewTLMessagesFeaturedStickers(&mtproto.Messages_FeaturedStickers{})
	result.SetCount(int32(len(stickerSetList)))
	result.SetHash(hash.Hash(idList, hash.UnLimitMaxCountHash))
	log.Infof("MessagesGetFeaturedStickers %v, request: %v result:%+v", metadata.RpcMetaDataDebug(md), request, result.Data2)
	return result.To_Messages_FeaturedStickers(), nil
}
