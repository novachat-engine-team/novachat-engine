/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.updateUsername_handler.go
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
	"strings"
)

//  account.updateUsername#3e0bdd7c username:string = User;
//
func (s *AccountServiceImpl) AccountUpdateUsername(ctx context.Context, request *mtproto.TLAccountUpdateUsername) (*mtproto.User, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AccountUpdateUsername %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var err error
	if util.IsAlNumString(request.Username) == false {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USERNAME_INVALID)
		log.Warnf("AccountUpdateUsername - request: %v error:%s", request, err.Error())
		return nil, err
	}

	userInfo, err := s.accountUsersCore.UserData(md.UserId)
	if err != nil {
		log.Errorf("AccountUpdateUsername - request: %v GetByUserId error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}

	if strings.Compare(userInfo.GetUsername(), request.Username) == 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USERNAME_NOT_MODIFIED)
		log.Warnf("AccountUpdateUsername - request: %v error:%s", request, err.Error())
		return nil, err
	}

	if len(request.Username) > 0 {
		var exists bool
		exists, err = s.accountCore.UserNameExists(request.Username)
		if err != nil {
			log.Warnf("AccountUpdateUsername - request: %v UserNameExists error:%s", request, err.Error())
		} else {
			if exists {
				err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USERNAME_OCCUPIED)
				return nil, err
			}
		}
	}
	ok, err := s.usersCore.UpdateUsername(md.UserId, request.Username)
	if err != nil {
		log.Errorf("AccountUpdateUsername -request: %v UpdateUsername error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}

	if ok == false {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USERNAME_OCCUPIED)
		log.Warnf("AccountUpdateUsername - request: %v error:%s", request, err.Error())
		return nil, err
	}

	user := usersUtil.UserCoreSelfUsers(usersUtil.UserCore2Users(userInfo))
	user.Username = request.Username
	log.Debugf("AccountUpdateUsername - request: %v reply ok !!!!!!", request)
	return user, nil
}
