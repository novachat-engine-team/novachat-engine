/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.sendChangePhoneCode_handler.go
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
	"strings"
)

//  account.sendChangePhoneCode#82574ae5 phone_number:string settings:CodeSettings = auth.SentCode;
//  account.sendChangePhoneCode#8e57deb flags:# allow_flashcall:flags.0?true phone_number:string current_number:flags.0?Bool = auth.SentCode;
//
func (s *AccountServiceImpl) AccountSendChangePhoneCode(ctx context.Context, request *mtproto.TLAccountSendChangePhoneCode) (*mtproto.Auth_SentCode, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountSendChangePhoneCode %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	phoneNumber := strings.TrimSpace(request.PhoneNumber)
	var err error
	if len(phoneNumber) == 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_EMPTY)
		log.Errorf("AccountSendChangePhoneCode - request:%s SendSmsCode error:%s", request, err.Error())
		return nil, err
	}
	phoneNumberParser, err := phone.Parse(phoneNumber)
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
		log.Errorf("AccountSendChangePhoneCode - request:%s SendSmsCode error:%s", request, err.Error())
		return nil, err
	}

	phoneNumber = phoneNumberParser.NormalizeDigitsOnly()
	userInfo, err := s.accountUsersCore.PhoneUser(phoneNumber)
	if err != nil {
		log.Errorf("AccountSendChangePhoneCode - request: %v error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
	}

	if constants.PeerTypeFromUserIDType(userInfo.Id).ToInt() == md.UserId {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
		log.Errorf("AccountSendChangePhoneCode - request:%s error:%s", request, err.Error())
		return nil, err
	}

	if userInfo.Id != 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_OCCUPIED)
		log.Errorf("AccountSendChangePhoneCode - UserByPhoneNumber request:%s SendSmsCode error:%s", request, err.Error())
		return nil, err
	}

	phoneCode := sms.DefaultPhoneCode
	if s.conf.EnvMode == constants.EnvModelTest {
		phoneCode = sms.DefaultPhoneCode
	} else {
		//if config.SendPhoneCodeIgnore(s.conf.Ignore.Phones, phoneNumber) {
		//	phoneCode = sms.DefaultPhoneCode
		//} else {
		//	if phoneNumberParser.IsVirtualPhone() {
		//		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
		//		log.Errorf("AccountSendChangePhoneCode - UserByPhoneNumber request:%s SendSmsCode error:%s", request, err.Error())
		//		return nil, err
		//	}

		phoneCode = random2.RandomAlphaOrNumeric(sms.CodeLength, false, true)
		//}
	}

	phoneHashCode := random2.RandomAlphaOrNumeric(32, true, false)
	err = sms.SendSmsVerifyCode(sms.CodeChange, phoneNumber, phoneCode, phoneHashCode)
	if err != nil {
		log.Errorf("AccountSendChangePhoneCode - request:%s   SendSmsCode error:%s", request, err.Error())
		return nil, err
	}

	log.Infof("AccountSendChangePhoneCode - request: %v ok!!!!!!!!!!!!!!", request)

	return mtproto.NewTLAuthSentCode(&mtproto.Auth_SentCode{
		Type: mtproto.NewTLAuthSentCodeTypeSms(&mtproto.Auth_SentCodeType{
			Length: sms.CodeLength,
		}).To_Auth_SentCodeType(),
		PhoneCodeHash: phoneHashCode,
		NextType:      mtproto.NewTLAuthCodeTypeSms(nil).To_Auth_CodeType(),
		Timeout:       sms.TTL,
	}).To_Auth_SentCode(), nil
}
