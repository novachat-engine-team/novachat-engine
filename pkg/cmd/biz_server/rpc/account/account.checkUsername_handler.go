/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.checkUsername_handler.go
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
	"novachat_engine/service/common/errors"
)

//Name	Type	Description
//username	string	username
//Accepted characters: A-z (case-insensitive), 0-9 and underscores.
//Length: 5-32 characters.
//Result
//Return Bool result on whether the passed username can be used.
//
//Possible errors
//Code	Type	Description
//400	USERNAME_INVALID	Unacceptable username

//  account.checkUsername#2714d86c username:string = Bool;
//
func (s *AccountServiceImpl) AccountCheckUsername(ctx context.Context, request *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AccountCheckUsername %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var err error
	if util.IsAlNumString(request.Username) == false {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USERNAME_INVALID)
		log.Debugf("AccountCheckUsername - request: %v error:%s", request, err.Error())
		return nil, err
	}

	ok, err := s.accountCore.UserNameExists(request.Username)
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
		log.Debugf("AccountCheckUsername -request: %v error:%s", request, err.Error())
		return nil, err
	}

	//if ok == true {
	//	err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USERNAME_OCCUPIED)
	//	log.Debugf("AccountCheckUsername - request: %v error:%s", request, err.Error())
	//	return nil, err
	//}

	log.Debugf("AccountCheckUsername - request: %v reply ok!!!!!!!!!!!!!", request)
	return mtproto.ToMTBool(!ok), nil
}
