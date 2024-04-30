/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.checkPhone_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
)

//  auth.checkPhone#6fe51dfb phone_number:string = auth.CheckedPhone;
//
func (s *AuthServiceImpl) AuthCheckPhone(ctx context.Context, request *mtproto.TLAuthCheckPhone) (*mtproto.Auth_CheckedPhone, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AuthCheckPhone %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	ok, err := s.accountCore.PhoneNumberExists(request.PhoneNumber)
	if err != nil {
		log.Errorf("AuthCheckPhone - request: %v error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
	}

	checkPhone := mtproto.NewTLAuthCheckedPhone(nil)
	checkPhone.SetPhoneRegistered(mtproto.ToMTBool(ok))
	return checkPhone.To_Auth_CheckedPhone(), nil
}
