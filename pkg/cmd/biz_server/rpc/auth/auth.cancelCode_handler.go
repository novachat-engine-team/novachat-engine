/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.cancelCode_handler.go
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
	"novachat_engine/service/core/account/phone"
	"novachat_engine/service/sms"
)

//  auth.cancelCode#1f040578 phone_number:string phone_code_hash:string = Bool;
//
func (s *AuthServiceImpl) AuthCancelCode(ctx context.Context, request *mtproto.TLAuthCancelCode) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AuthCancelCode %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	phoneNumber, err := phone.Parse(request.PhoneNumber)
	if err != nil {
		log.Errorf("AuthCancelCode %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
	}

	if s.banned.InvalidNumber(phoneNumber.NormalizeDigitsOnly()) {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_BANNED)
	}

	if len(request.PhoneCodeHash) == 0 {
		log.Errorf("AuthCancelCode %v, request: %v error:len(request.PhoneCodeHash) == 0", metadata.RpcMetaDataDebug(md), request)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_HASH_EMPTY)
	}

	isVirtualPhone := phone.IsVirtualPhone(request.PhoneNumber)

	log.Infof("AuthCancelCode authKeyId:%d phoneNumber:%s", md.AuthKeyId, request.PhoneNumber)
	if !isVirtualPhone {
		sms.ClearSmsCode(sms.CodeSignIn, phoneNumber.NormalizeDigitsOnly(), "", request.PhoneCodeHash)
	}
	return mtproto.ToMTBool(true), nil
}
