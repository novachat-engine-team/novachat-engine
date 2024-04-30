/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : photos.uploadProfilePhoto_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/utils"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/app/dc"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	data_fs "novachat_engine/service/data/fs"
	"novachat_engine/service/upload/photo"
	"time"
)

//  photos.uploadProfilePhoto#89f30f69 flags:# file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = photos.Photo;
//  photos.uploadProfilePhoto#4f32c098 file:InputFile = photos.Photo;
//
func (s *PhotosServiceImpl) PhotosUploadProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhotosUploadProfilePhoto %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if md.UserId == 0 {
		log.Fatalf("PhotosUploadProfilePhoto UploadPhoto user_id = 0 %+v", request)
		return mtproto.NewTLPhotosPhoto(&mtproto.Photos_Photo{
			Photo: mtproto.NewTLPhotoEmpty(nil).To_Photo(),
			Users: []*mtproto.User{},
		}).To_Photos_Photo(), nil
	}

	var photoValue *mtproto.Photo
	dataPhoto := &data_fs.PhotoProfile{}
	if request.GetVideo() != nil {
		documentInfo, err1 := sfsService.GetSfsClient(
			fmt.Sprintf("%d", request.GetFile().GetId())).ReqUploadedDocument(ctx,
			&sfsService.UploadedDocument{
				AuthKeyId: md.AuthKeyId,
				FileId:    request.GetFile().GetId(),
				FileName:  request.GetFile().GetName(),
				Parts:     request.GetFile().GetParts(),
				Thumb: &sfsService.UploadedDocument_Thumb{
					FileId:   request.GetVideo().GetId(),
					FileName: request.GetVideo().GetName(),
					Parts:    request.GetVideo().GetParts(),
				},
				VideoStartTs: request.GetVideoStartTs(),
				Video:        true,
			})
		if err1 != nil {
			log.Errorf("PhotosUploadProfilePhoto request:%+v ReqUploadedDocument error: %v", request, err1)
			return nil, err1
		}
		dataPhoto.Photo = utils.ToDataPhoto(documentInfo.Thumbs[0]).Detail[0]
		dataPhoto.Video = utils.DocumentToVideo(documentInfo).Detail[0]
	} else {
		photoInfo, err1 := sfsService.GetSfsClient(
			fmt.Sprintf("%d", request.GetFile().GetId())).ReqUploadPhoto(ctx,
			&sfsService.UploadPhoto{
				AuthKeyId: md.AuthKeyId,
				FileId:    request.GetFile().GetId(),
				FileName:  request.GetFile().GetName(),
				Parts:     request.GetFile().GetParts(),
			})
		if err1 != nil {
			log.Errorf("PhotosUploadProfilePhoto request:%+v ReqUploadPhoto error: %v", request, err1)
			return nil, err1
		}
		dataPhoto.Photo = utils.ToDataPhoto(photoInfo).Detail[0]
		photoValue = photo.PhotoInfo2Photo(photoInfo, md.Layer)
	}

	_, err := s.userCore.SetPhotoId(md.UserId, dataPhoto.Photo.VolumeId)
	if err != nil {
		log.Errorf("PhotosUploadProfilePhoto request:%+v UserSetPhotoId error: %v", request, err)
		return nil, err
	}

	userProfilePhoto := mtproto.NewTLUserProfilePhoto(&mtproto.UserProfilePhoto{
		PhotoId: dataPhoto.Photo.VolumeId,
		DcId:    dc.DefaultDc,
	}).To_UserProfilePhoto()
	userProfilePhoto.PhotoSmall = compat.NewFileLocationByLayer(
		dataPhoto.Photo.VolumeId,
		dataPhoto.Photo.LocalId,
		md.Layer,
		dc.DefaultDc)
	userProfilePhoto.HasVideo = dataPhoto.Video != nil
	if dataPhoto.Video != nil {
		userProfilePhoto.PhotoBig = compat.NewFileLocationByLayer(
			dataPhoto.Video.VolumeId,
			dataPhoto.Video.LocalId,
			md.Layer,
			dc.DefaultDc)
	}
	photosPhoto := mtproto.NewTLPhotosPhoto(&mtproto.Photos_Photo{
		Photo: photoValue,
		Users: []*mtproto.User{},
	})

	go func() {
		_, err = syncService.GetSyncClientById(md.UserId).ReqSyncUpdate(context.Background(), &syncService.SyncUpdate{
			UserId:          md.UserId,
			AuthKeyId:       0,
			IgnoreAuthKeyId: md.AuthKeyId,
			Updates: mtproto.NewTLUpdates(&mtproto.Updates{
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
				Updates: []*mtproto.Update{
					//  updateUserPhoto#95313b0c user_id:int date:int photo:UserProfilePhoto previous:Bool = Update;
					mtproto.NewTLUpdateUserPhoto(&mtproto.Update{
						UserId:   constants.PeerTypeFromUserIDType(md.UserId).ToInt32(),
						Date:     int32(time.Now().Unix()),
						Photo:    userProfilePhoto,
						Previous: mtproto.ToMTBool(false),
					}).To_Update(),
				},
				Date: int32(time.Now().Unix()),
				Seq:  0,
			}).To_Updates(),
			Push:     false,
			PeerType: constants.PeerTypeUser.ToInt32(),
		})
		if err != nil {
			log.Errorf("PhotosUploadProfilePhoto request:%+v error: %v", request, err)
		}
	}()

	return photosPhoto.To_Photos_Photo(), nil
}
