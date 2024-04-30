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
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  photos.updateProfilePhoto#72d4742c id:InputPhoto = photos.Photo;
//  photos.updateProfilePhoto#f0bb5152 id:InputPhoto = UserProfilePhoto;
//
func (s *PhotosServiceImpl) PhotosUpdateProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUpdateProfilePhoto) (*mtproto.Response_PhotosUpdateProfilePhoto, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhotosUpdateProfilePhoto %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var userProfilePhoto *mtproto.UserProfilePhoto
	response := &mtproto.Response_PhotosUpdateProfilePhoto{
		Cmd:                              0,
		PhotosUpdateProfilePhoto72D4742C: nil,
		PhotosUpdateProfilePhotof0Bb5152: nil,
	}
	//if request.Id.GetClassName() == mtproto.ClassInputPhotoEmpty {
	//	v, err := s.userCore.DeletePhoto(md.UserId, []int64{})
	//	if err != nil {
	//		log.Errorf("PhotosUpdateProfilePhoto - request: %v error:%s", request, err.Error())
	//		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	//	}
	//	id := v
	//	response.PhotosUpdateProfilePhoto72D4742C = mtproto.NewTLPhotosPhoto(&mtproto.Photos_Photo{
	//		Photo: mtproto.NewTLPhotoEmpty(&mtproto.Photo{Id: id}).To_Photo(),
	//		Users: []*mtproto.User{},
	//	}).To_Photos_Photo()
	//	response.PhotosUpdateProfilePhotof0Bb5152 = mtproto.NewTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	//
	//	userProfilePhoto = mtproto.NewTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	//} else {
	//	id := request.GetId().To_InputPhoto()
	//
	//	if s.conf.BizLogic.FilterCloud.FilterCloud && s.conf.BizLogic.FilterCloud.Image {
	//		fileData, err := sfsClient.GetSfsClient().UploadGetFile(context.Background(), &sfs_pb.UploadGetFile{
	//			Id:        id.GetId(),
	//			VolumeId:  id.GetId(),
	//			Limit:     math.MaxInt32,
	//			ThumbSize: "m",
	//		})
	//		if err == nil && fileData != nil && fileData.Bytes != nil {
	//			ok := filter.FilterImage(fileData.Bytes, s.conf.BizLogic.FilterCloud.FilterCloud && s.conf.BizLogic.FilterCloud.Image)
	//			if ok == false {
	//				err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_FORBIDDEN_BLOCK_IMAGE)
	//				log.Errorf("PhotosUpdateProfilePhoto request:%+v error: %v", request, err)
	//				return nil, err
	//			}
	//		}
	//	}
	//
	//	photo, err := sfsClient.GetSfsClient().GetPhotoFile(context.Background(), &sfs_pb.GetPhotoFile{
	//		PhotoId: id.GetId(),
	//		Layer:   md.Layer,
	//	})
	//	if err != nil {
	//		log.Errorf("PhotosUpdateProfilePhoto - request: %v error:%s", request, err.Error())
	//		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	//	}
	//
	//	if len(photo.Sizes) == 0 {
	//		response.PhotosUpdateProfilePhotof0Bb5152 = mtproto.NewTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	//		userProfilePhoto = response.PhotosUpdateProfilePhotof0Bb5152
	//		idList, err := s.userCore.UserPhotos(md.UserId)
	//		if err != nil {
	//			log.Errorf("PhotosUpdateProfilePhoto - request: %v error:%s", request, err.Error())
	//			return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	//		}
	//		idPhoto := int64(0)
	//		if len(idList) > 0 {
	//			idPhoto = idList[0]
	//		}
	//		response.PhotosUpdateProfilePhoto72D4742C = mtproto.NewTLPhotosPhoto(&mtproto.Photos_Photo{
	//			Photo: mtproto.NewTLPhotoEmpty(&mtproto.Photo{Id: idPhoto}).To_Photo(),
	//			Users: []*mtproto.User{},
	//		}).To_Photos_Photo()
	//	} else {
	//		_, err = s.userCore.SetPhotoId(md.UserId, photo.Id)
	//		if err != nil {
	//			log.Errorf("PhotosUpdateProfilePhoto - request: %v UserSetPhotoId error:%s", request, err.Error())
	//			return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	//		}
	//
	//		response.PhotosUpdateProfilePhoto72D4742C = mtproto.NewTLPhotosPhoto(&mtproto.Photos_Photo{
	//			Photo: photo,
	//			Users: []*mtproto.User{},
	//		}).To_Photos_Photo()
	//
	//		response.PhotosUpdateProfilePhotof0Bb5152 = mtproto.NewTLUserProfilePhoto(&mtproto.UserProfilePhoto{
	//			HasVideo:   false,
	//			PhotoId:    photo.Id,
	//			PhotoSmall: photo.Sizes[0].Location,
	//			PhotoBig:   photo.Sizes[len(photo.Sizes)-1].Location,
	//			DcId:       photo.DcId,
	//		}).To_UserProfilePhoto()
	//
	//		userProfilePhoto = response.PhotosUpdateProfilePhotof0Bb5152
	//	}
	//}

	//go func() {
	//	_, err := syncClient.GetSyncClient().PushUpdates(context.Background(), &sync_pb.PushUpdates{
	//		UserId: md.UserId,
	//		Updates: mtproto.NewTLUpdates(&mtproto.Updates{
	//			Users: []*mtproto.User{},
	//			Chats: []*mtproto.Chat{},
	//			Updates: []*mtproto.Update{
	//				//  updateUserPhoto#95313b0c user_id:int date:int photo:UserProfilePhoto previous:Bool = Update;
	//				mtproto.NewTLUpdateUserPhoto(&mtproto.Update{
	//					UserId:   md.UserId,
	//					Date:     int32(time.Now().Unix()),
	//					Photo:    userProfilePhoto,
	//					Previous: mtproto.ToMTBool(false),
	//				}).To_Update(),
	//			},
	//			Date: int32(time.Now().Unix()),
	//			Seq:  0,
	//		}).To_Updates(),
	//	})
	//	if err != nil {
	//		log.Errorf("PhotosUpdateProfilePhoto request:%+v error: %v", request, err)
	//	}
	//}()

	_ = userProfilePhoto
	if md.Layer >= mtproto.CurrentLayer {
		response.Cmd = mtproto.Cmd_PhotosUpdateProfilePhoto72d4742c.ToUInt32()
	} else {
		response.Cmd = mtproto.Cmd_PhotosUpdateProfilePhoto.ToUInt32()
	}
	return response, nil
}
