/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.updateProfile_handler.go
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
	usersUtil "novachat_engine/service/account/users"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
)

//  account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
//
func (s *AccountServiceImpl) AccountUpdateProfile(ctx context.Context, request *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountUpdateProfile %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.About) >= constants.UserAboutLimit {
		log.Errorf("AccountUpdateProfile %v, request: %v About limit length:%d", md, request, len(request.About))
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_ABOUT_TOO_LONG)
	}

	var err error
	var ok bool
	if _, ok = util.FilterDirtyWord(request.FirstName+request.LastName+"    "+request.About, false); ok == true {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_FORBIDDEN_DIRTY_WORD)
		log.Errorf("AccountUpdateProfile %v, request: %v FilterDirtyWord error:%s", md, request, err.Error())
		return nil, err
	}

	if len(request.FirstName) == 0 && len(request.About) == 0 && len(request.LastName) == 0 {
		log.Errorf("AccountUpdateProfile %v, request: %v FirstName length = 0", md, request)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_FIRSTNAME_INVALID)
	}

	userInfo, err := s.accountUsersCore.UserData(md.UserId)
	if err != nil {
		log.Errorf("AccountUpdateProfile %v, request: %v UpdateProfile error:%s", md, request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
	}

	firstname := userInfo.FirstName
	if util.CheckBit(request.Flags, 0) {
		firstname = request.FirstName
		userInfo.FirstName = firstname
	}
	lastname := userInfo.LastName
	if util.CheckBit(request.Flags, 1) {
		lastname = request.LastName
		userInfo.LastName = lastname
	}
	about := userInfo.About
	if util.CheckBit(request.Flags, 2) {
		about = request.About
		userInfo.About = about
	}

	_, err = s.usersCore.UpdateProfile(md.UserId, firstname, lastname, about)
	if err != nil {
		log.Errorf("AccountUpdateProfile %v, request: %v UpdateProfile error:%s", md, request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
	}

	user, err := s.accountUsersCore.GetUser(md.UserId, md.UserId, md.Layer)
	if err != nil {
		user = usersUtil.UserCoreSelfUsers(usersUtil.UserCore2Users(userInfo))
	}
	return user, nil
}
