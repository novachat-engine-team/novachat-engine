/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.editPhoto_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/utils"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	data_fs "novachat_engine/service/data/fs"
	"time"
)

//  channels.editPhoto#f12e57c9 channel:InputChannel photo:InputChatPhoto = Updates;
//
func (s *ChannelsServiceImpl) ChannelsEditPhoto(ctx context.Context, request *mtproto.TLChannelsEditPhoto) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsEditPhoto %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var err error
	//var action *mtproto.MessageAction
	//var photoId int64
	var photoInfo *sfsService.PhotoInfo
	dataChatPhoto := &data_fs.PhotoProfile{}

	chatPhoto := request.GetPhoto()
	switch chatPhoto.ClassName {
	case mtproto.ClassInputChatPhotoEmpty:
		//action = mtproto.NewTLMessageActionChatDeletePhoto(nil).To_MessageAction()
	case mtproto.ClassInputChatUploadedPhoto:
		var documentInfo *sfsService.DocumentInfo
		if request.Photo.Video != nil {
			documentInfo, err = sfsService.GetSfsClient(fmt.Sprintf("%d", chatPhoto.GetVideo().Id)).
				ReqUploadedDocument(ctx,
					&sfsService.UploadedDocument{
						AuthKeyId: md.AuthKeyId,
						FileId:    chatPhoto.GetVideo().Id,
						FileName:  chatPhoto.GetVideo().Name,
						Parts:     chatPhoto.GetVideo().Parts,
						Thumb: &sfsService.UploadedDocument_Thumb{
							FileId:   chatPhoto.GetFile().Id,
							FileName: chatPhoto.GetFile().Name,
							Parts:    chatPhoto.GetFile().Parts,
							Sticker:  nil,
						},
						VideoStartTs: chatPhoto.GetVideoStartTs(),
						Video:        true,
					})
		} else {
			photoInfo, err = sfsService.GetSfsClient(fmt.Sprintf("%d", chatPhoto.GetFile().Id)).
				ReqUploadPhoto(ctx,
					&sfsService.UploadPhoto{
						AuthKeyId: md.AuthKeyId,
						FileId:    chatPhoto.GetFile().Id,
						FileName:  chatPhoto.GetFile().Name,
						Parts:     chatPhoto.GetFile().Parts,
						Stickers:  nil,
					})
		}
		if err != nil {
			log.Errorf("ChannelsEditPhoto error: %v", err)
			return nil, err
		}
		if documentInfo == nil {
			dataChatPhoto.Photo = utils.ToDataPhoto(photoInfo).Detail[0]
		} else {
			dataChatPhoto.Photo = utils.ToDataPhoto(documentInfo.Thumbs[0]).Detail[0]
			dataChatPhoto.Video = utils.DocumentToVideo(documentInfo).Detail[0]
		}
	case mtproto.ClassInputChatPhoto:
		photoInfo, err = sfsService.GetSfsClient(fmt.Sprintf("%d", chatPhoto.GetId().Id)).
			ReqGetPhoto(context.Background(),
				&sfsService.GetPhoto{
					PhotoId: chatPhoto.GetId().Id,
				})
		if err != nil {
			log.Errorf("ChannelsEditPhoto InputChatPhoto error: %v", err)
			return nil, err
		}
		dataChatPhoto.Photo = utils.ToDataPhoto(photoInfo).Detail[0]
	}

	chatId := constants.PeerTypeFromChatIDType32(request.Channel.ChannelId).ToInt()
	if chatId == 0 {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
		log.Errorf("ChannelsEditPhoto %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	editPhotoResp, err := chatService.GetChatClientByKeyId(chatId).
		ReqEditPhoto(ctx, &chatService.EditPhoto{
			ChatId:    chatId,
			Photo:     dataChatPhoto,
			UserId:    md.UserId,
			AuthKeyId: md.AuthKeyId,
			Date:      int32(time.Now().Unix()),
		})
	if err != nil {
		log.Errorf("MessagesEditChatTitle %v, request: %v ReqEditTitle error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	var updates mtproto.Updates
	types.UnmarshalAny(editPhotoResp, &updates)
	return compat.UpdatesCompat(&updates, md.Layer), nil
}
