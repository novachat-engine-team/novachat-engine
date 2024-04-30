/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.installStickerSet_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/core/message/message"
	"novachat_engine/service/upload/photo"
	"time"
)

//  messages.installStickerSet#c78fe460 stickerset:InputStickerSet archived:Bool = messages.StickerSetInstallResult;
//  messages.installStickerSet#7b30c3a6 stickerset:InputStickerSet disabled:Bool = Bool;
//
func (s *MessagesServiceImpl) MessagesInstallStickerSet(ctx context.Context, request *mtproto.TLMessagesInstallStickerSet) (*mtproto.Response_MessagesInstallStickerSet, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesInstallStickerSet %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	date := int32(time.Now().Unix())
	//Response_MessagesInstallStickerSet
	var result *mtproto.Messages_StickerSetInstallResult
	stickerSet, err := s.accountStickerSetCore.Install(md.UserId, request.Stickerset.Id, date, mtproto.ToBool(request.Archived))
	if err != nil {
		log.Errorf("MessagesInstallStickerSet %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	if mtproto.ToBool(request.Archived) == true {

		if len(stickerSet.Packs) == 0 || len(stickerSet.Packs[0].Documents) == 0 {
			result = mtproto.NewTLMessagesStickerSetInstallResultSuccess(nil).To_Messages_StickerSetInstallResult()
		} else {
			coverd := make([]*mtproto.StickerSetCovered, 0, 1)
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
			coverd = append(coverd, mtproto.NewTLStickerSetCovered(&mtproto.StickerSetCovered{
				Set: mtproto.NewTLStickerSet(&mtproto.StickerSet{
					Installed:     true,
					Archived:      true,
					Official:      stickerSet.Official,
					Masks:         stickerSet.Masks,
					Id:            stickerSet.Id,
					AccessHash:    stickerSet.AccessHash,
					Title:         stickerSet.Title,
					ShortName:     stickerSet.ShortName,
					Count:         stickerSet.Count,
					Hash:          stickerSet.Hash,
					InstalledDate: date,
					Thumb:         photoSize,
					ThumbDcId:     stickerSet.ThumbDcId,
					Animated:      stickerSet.Animated,
					Thumbs:        photoSizes,
				}).To_StickerSet(),
				Cover: message.ToMessageDocument(stickerSet.Packs[0].Documents[0], md.Layer),
			}).To_StickerSetCovered())

			result = mtproto.NewTLMessagesStickerSetInstallResultArchive(nil).To_Messages_StickerSetInstallResult()
			result.Sets = coverd
		}
	} else {
		result = mtproto.NewTLMessagesStickerSetInstallResultSuccess(nil).To_Messages_StickerSetInstallResult()
	}

	value := &mtproto.Response_MessagesInstallStickerSet{}
	if md.Layer >= 51 {
		value.Cmd = mtproto.Cmd_MessagesInstallStickerSet7b30c3a6.ToUInt32()
		value.MessagesInstallStickerSet7B30C3A6 = mtproto.ToMTBool(true)
	} else {
		value.Cmd = mtproto.Cmd_MessagesInstallStickerSet7b30c3a6.ToUInt32()
		value.MessagesInstallStickerSetc78Fe460 = result
	}
	log.Infof("MessagesInstallStickerSet %v, request: %v ok", metadata.RpcMetaDataDebug(md), request)
	return value, nil
}
