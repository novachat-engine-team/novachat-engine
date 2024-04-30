/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : photos.deletePhotos_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"time"
)

//  photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
//
func (s *PhotosServiceImpl) PhotosDeletePhotos(ctx context.Context, request *mtproto.TLPhotosDeletePhotos) (*mtproto.VectorLong, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhotosDeletePhotos %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var err error
	if len(request.Id) == 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST)
		log.Errorf("PhotosDeletePhotos - request: %v error:%s", request, err.Error())
		return nil, err
	}

	photoIdList := make([]int64, 0, len(request.Id))
	for _, id := range request.Id {
		photoIdList = append(photoIdList, id.Id)
	}

	var photoId int64
	photoId, err = s.userCore.DeletePhoto(md.UserId, photoIdList)
	if err != nil {
		log.Errorf("PhotosDeletePhotos - request: %v UserDeletePhoto error:%s", request, err.Error())
		return nil, err
	}

	var userProfilePhoto *mtproto.UserProfilePhoto
	if photoId != 0 {
		photoInfo, err1 := sfsService.GetSfsClient(fmt.Sprintf("%d", photoId)).
			ReqGetPhoto(context.Background(), &sfsService.GetPhoto{
				PhotoId: photoId,
			})
		if err1 != nil {
			log.Errorf("PhotosDeletePhotos - request: %v GetPhotoFile error:%s", request, err1.Error())
			return nil, err1
		}
		_ = photoInfo
		//userProfilePhoto = mtproto.NewTLUserProfilePhoto(&mtproto.UserProfilePhoto{
		//	HasVideo:   false,
		//	PhotoId:    photo.Id,
		//	PhotoSmall: photo.Sizes[0].Location,
		//	PhotoBig:   photo.Sizes[len(photo.Sizes)-1].Location,
		//	DcId:       photo.DcId,
		//}).To_UserProfilePhoto()
	} else {
		userProfilePhoto = mtproto.NewTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	}

	_ = userProfilePhoto

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

	return &mtproto.VectorLong{
		Datas: photoIdList,
	}, nil
}
