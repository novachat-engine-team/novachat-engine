/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.resendCode_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/random2"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/account/phone"
	"novachat_engine/service/sms"
)

//  auth.resendCode#3ef1a9bf phone_number:string phone_code_hash:string = auth.SentCode;
//
func (s *AuthServiceImpl) AuthResendCode(ctx context.Context, request *mtproto.TLAuthResendCode) (*mtproto.Auth_SentCode, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AuthResendCode %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	phoneNumber, err := phone.Parse(request.PhoneNumber)
	if err != nil {
		log.Errorf("AuthResendCode - request: %v error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
	}

	if s.banned.InvalidNumber(phoneNumber.NormalizeDigitsOnly()) {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_BANNED)
	}

	if len(request.PhoneCodeHash) == 0 {
		log.Errorf("AuthResendCode - request: %v error:len(request.PhoneCodeHash) == 0", request)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_HASH_EMPTY)
	}

	isVirtualPhone := phone.IsVirtualPhone(request.PhoneNumber)
	reply := mtproto.NewTLAuthSentCode(nil)
	reply.SetPhoneCodeHash(request.PhoneCodeHash)
	reply.SetNextType(constants.MakeAuthCode(constants.CodeTypeApp))
	reply.SetType(constants.MakeSendCodeType(constants.CodeTypeSms, sms.CodeLength, "*"))
	reply.SetTimeout(sms.TTL)
	//reply.SetPhoneRegistered(true)

	log.Infof("AuthResendCode - authKeyId:%d phoneNumber:%s", md.AuthKeyId, request.PhoneNumber)

	if !isVirtualPhone {
		var newPhoneCode string
		if s.conf.EnvMode == constants.EnvModelTest {
			newPhoneCode = sms.DefaultPhoneCode
		} else {
			//if config.SendPhoneCodeIgnore(s.conf.Ignore.Phones, phoneNumber.NormalizeDigitsOnly()) {
			//	newPhoneCode = sms.DefaultPhoneCode
			//} else {
			//	if phone.IsVirtualPhone(request.PhoneNumber) {
			//		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_BANNED)
			//	}
			newPhoneCode = random2.RandomAlphaOrNumeric(sms.CodeLength, false, true)
			//}
		}

		ttl, err := sms.ReSendSmsVerifyCode(sms.CodeSignIn, phoneNumber.NormalizeDigitsOnly(), request.PhoneCodeHash, newPhoneCode)
		if err != nil {
			log.Errorf("AuthResendCode - request: %v ReSendSmsVerifyCode error:%s", request, err.Error())
			return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_BAD_REQUEST.ToInt32(), err.Error())
		}
		reply.SetTimeout(ttl)
	}

	log.Debugf("AuthResendCode - request: %v ok!!!", request)
	return reply.To_Auth_SentCode(), nil

}
