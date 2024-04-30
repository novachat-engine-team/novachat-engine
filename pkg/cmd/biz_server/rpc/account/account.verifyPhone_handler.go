/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.verifyPhone_handler.go
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
	"novachat_engine/service/constants"
	"novachat_engine/service/sms"
	"strings"
)

//  account.verifyPhone#4dd3a7f6 phone_number:string phone_code_hash:string phone_code:string = Bool;
//
func (s *AccountServiceImpl) AccountVerifyPhone(ctx context.Context, request *mtproto.TLAccountVerifyPhone) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountVerifyPhone %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if s.conf.EnvMode != constants.EnvModelTest {
		if strings.Compare(request.PhoneCode, sms.DefaultPhoneCode) != 0 {
			log.Errorf("AccountVerifyPhone - request: %v error", request)
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_INVALID)
		}
	} else {
		err := sms.VerifySmsCode(sms.CodeVerify, request.PhoneNumber, request.PhoneCode, request.PhoneCodeHash)
		if err != nil {
			log.Errorf("AccountVerifyPhone -  request: %v VerifySmsCode error:%s", request, err.Error())
			return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_BAD_REQUEST.ToInt32(), err.Error())
		}
	}

	return mtproto.ToMTBool(true), nil
}
