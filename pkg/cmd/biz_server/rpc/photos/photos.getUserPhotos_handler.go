/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : photos.getUserPhotos_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
)

//  photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
//
func (s *PhotosServiceImpl) PhotosGetUserPhotos(ctx context.Context, request *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhotosGetUserPhotos %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PhotosGetUserPhotos logic
	inputUser := input.MakeInputUser(request.UserId)
	if inputUser.GetInputUserType() == constants.InputUserTypeSelf {
		inputUser = input.MakeInputUser(mtproto.NewTLInputUser(&mtproto.InputUser{
			UserId:     constants.PeerTypeFromUserIDType(md.UserId).ToInt32(),
			AccessHash: request.UserId.AccessHash,
		}).To_InputUser())
	}

	photos := mtproto.NewTLPhotosPhotos(&mtproto.Photos_Photos{
		Photos: []*mtproto.Photo{},
		Users:  []*mtproto.User{},
		Count:  0,
	})
	//userPhoto, err := s.userCore.UserPhotos(md.UserId)
	//if err != nil {
	//	log.Errorf("PhotosGetUserPhotos - request: %v error:%s", request, err.Error())
	//	return nil, err
	//}
	//if len(userPhoto) > 0 {
	//	l, err := sfsClient.GetSfsClient().GetPhotoFileList(context.Background(), &sfs_pb.GetPhotoFileList{
	//		PhotoId: userPhoto,
	//		Layer:   md.Layer,
	//	})
	//	if err != nil {
	//		log.Errorf("PhotosGetUserPhotos - request: %v GetPhotoFileList error:%s", request, err.Error())
	//		return nil, err
	//	}
	//
	//	photos.SetPhotos(l.Photos)
	//	photos.SetUsers([]*mtproto.User{})
	//
	//	log.Debugf("PhotosGetUserPhotos - request: %v ok:%v", request, photos.To_Photos_Photos())
	//} else {
	//	log.Debugf("PhotosGetUserPhotos - request: %v ok:%v", request, photos.To_Photos_Photos())
	//}
	return photos.To_Photos_Photos(), nil
}
