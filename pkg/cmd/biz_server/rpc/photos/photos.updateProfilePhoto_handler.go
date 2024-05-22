/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : photos.updateProfilePhoto_handler.go
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
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/util"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/account/users"
	data_fs "novachat_engine/service/data/fs"
	data_users "novachat_engine/service/data/users"
	"novachat_engine/service/upload/photo"
	"time"
)

//  photos.updateProfilePhoto#72d4742c id:InputPhoto = photos.Photo;
//  photos.updateProfilePhoto#f0bb5152 id:InputPhoto = UserProfilePhoto;
//
func (s *PhotosServiceImpl) PhotosUpdateProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.Response_PhotosUpdateProfilePhoto, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhotosUpdateProfilePhoto %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var photoProfile *mtproto.UserProfilePhoto
	response := &mtproto.Response_PhotosUpdateProfilePhoto{
		Cmd:                              0,
		PhotosUpdateProfilePhoto72D4742C: nil,
		PhotosUpdateProfilePhotof0Bb5152: nil,
	}
	if request.Id.GetClassName() == mtproto.ClassInputPhotoEmpty {
		v, err := s.userCore.DeletePhoto(md.UserId, []int64{})
		if err != nil {
			log.Errorf("PhotosUpdateProfilePhoto - request: %v error:%s", request, err.Error())
			return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
		}
		var id int64
		if v != nil {
			id = v.PhotoId
		}
		response.PhotosUpdateProfilePhoto72D4742C = mtproto.NewTLPhotosPhoto(&mtproto.Photos_Photo{
			Photo: mtproto.NewTLPhotoEmpty(&mtproto.Photo{Id: id}).To_Photo(),
			Users: []*mtproto.User{},
		}).To_Photos_Photo()
		response.PhotosUpdateProfilePhotof0Bb5152 = mtproto.NewTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()

		photoProfile = mtproto.NewTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	} else {
		id := request.GetId().To_InputPhoto().GetId()
		userInfo, err := s.userCore.User(md.UserId)
		if err != nil {
			log.Errorf("PhotosUpdateProfilePhoto - request: %v User error:%s", request, err.Error())
			return nil, err
		}
		//sfsService.GetSfsClient(fmt.Sprintf("%d", id)).

		profilePhotoIds := users.MakeProfilePhotoData(userInfo.Photos)
		idx := util.Index(profilePhotoIds.IdList, func(idx int) bool {
			return profilePhotoIds.IdList[idx].PhotoId == id
		})
		var dataProfile *data_users.ProfilePhotoData
		if idx < 0 {
			dataProfile, err = s.userCore.SetPhotoId(md.UserId, 0, 0)
		} else {
			dataProfile, err = s.userCore.SetPhotoId(md.UserId, profilePhotoIds.IdList[idx].PhotoId, profilePhotoIds.IdList[idx].VideoId)
		}
		if err != nil {
			log.Errorf("PhotosUpdateProfilePhoto - request: %v UserSetPhotoId error:%s", request, err.Error())
			return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
		}

		var photoValue *mtproto.Photo
		if dataProfile.PhotoId != 0 || dataProfile.VideoId != 0 {
			sfsIns := sfsService.GetSfsClient("0")
			if sfsIns == nil {
				log.Warnf("PhotosUpdateProfilePhoto sfsService = nil")
			} else {
				if dataProfile.PhotoId != 0 {
					photoInfo, err := sfsIns.ReqGetPhoto(context.Background(), &sfsService.GetPhoto{
						PhotoId: dataProfile.PhotoId,
						LocalId: 0,
					})
					if err != nil {
						log.Infof("PhotosUpdateProfilePhoto ReqGetPhoto error:%s", err.Error())
					} else {
						photoValue = photo.PhotoInfo2Photo(photoInfo, md.Layer)
					}
				}
				if dataProfile.VideoId != 0 {
					videoInfo, err := sfsIns.ReqGetDocument(context.Background(), &sfsService.GetDocument{
						FileId:  dataProfile.VideoId,
						LocalId: 0,
					})
					if err != nil {
						log.Infof("PhotosUpdateProfilePhoto ReqGetDocument error:%s", err.Error())
					} else {
						photoProfile = photo.PhotoProfileUserProfilePhoto(&data_fs.PhotoProfile{
							Photo: utils.ToDataPhoto(videoInfo.Thumbs[0]).Detail[0],
							Video: utils.DocumentToVideo(videoInfo).Detail[0],
						}, md.Layer)
					}
				}
			}
		} else {
			photoValue = mtproto.NewTLPhotoEmpty(&mtproto.Photo{Id: 0}).To_Photo()
			photoProfile = mtproto.NewTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
		}

		response.PhotosUpdateProfilePhoto72D4742C = mtproto.NewTLPhotosPhoto(&mtproto.Photos_Photo{
			Photo: photoValue,
			Users: []*mtproto.User{},
		}).To_Photos_Photo()
		response.PhotosUpdateProfilePhotof0Bb5152 = photoProfile
	}

	go func() {
		_, err := syncClient.GetSyncClient(fmt.Sprintf("%d", md.UserId)).ReqSyncUpdate(context.Background(),
			&syncClient.SyncUpdate{
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
							Photo:    photoProfile,
							Previous: mtproto.ToMTBool(false),
						}).To_Update(),
					},
					Date: int32(time.Now().Unix()),
					Seq:  0,
				}).To_Updates(),
				Push:     false,
				PeerType: 0,
			})
		if err != nil {
			log.Errorf("PhotosUpdateProfilePhoto request:%+v error: %v", request, err)
		}
	}()

	if md.Layer >= mtproto.CurrentLayer {
		response.Cmd = mtproto.Cmd_PhotosUpdateProfilePhoto72d4742c.ToUInt32()
	} else {
		response.Cmd = mtproto.Cmd_PhotosUpdateProfilePhoto.ToUInt32()
	}
	return response, nil
}
