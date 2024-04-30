/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getArchivedStickers_handler.go
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
	"novachat_engine/service/core/message/message"
	data_stickerset "novachat_engine/service/data/stickerset"
	"novachat_engine/service/upload/photo"
)

//  messages.getArchivedStickers#57f17692 flags:# masks:flags.0?true offset_id:long limit:int = messages.ArchivedStickers;
//
func (s *MessagesServiceImpl) MessagesGetArchivedStickers(ctx context.Context, request *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetArchivedStickers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	installList, err := s.accountStickerSetCore.GetInstalled(md.UserId)
	if err != nil {
		log.Errorf("MessagesGetArchivedStickers %v, request: %v GetInstalled error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	stickerSetIdList := make([]int64, 0, len(installList))
	for _, v := range installList {
		if v.Archived {
			stickerSetIdList = append(stickerSetIdList, v.StickerSetId)
		}
	}
	archivedStickers := mtproto.NewTLMessagesArchivedStickers(nil)
	if len(stickerSetIdList) == 0 {
		return archivedStickers.To_Messages_ArchivedStickers(), nil
	}

	stickerSetList, err := s.accountStickerSetCore.GetStickers(stickerSetIdList)
	if err != nil {
		log.Errorf("MessagesGetArchivedStickers %v, request: %v GetStickers error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	archivedStickers.SetCount(int32(len(stickerSetList)))
	sets := make([]*mtproto.StickerSetCovered, 0, len(stickerSetList))

	for _, stickerSet := range stickerSetList {

		var install *data_stickerset.StickerInstall
		if idx := util.Index(&installList, func(idx int) bool {
			return installList[idx].StickerSetId == stickerSet.Id
		}); idx >= 0 {
			install = installList[idx]
		}
		if install == nil {
			continue
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

	archivedStickers.SetSets(sets)
	log.Debugf("MessagesGetArchivedStickers %v, request: %v reply ok!!!!!!!!!!", md, request)
	return archivedStickers.To_Messages_ArchivedStickers(), nil
}
