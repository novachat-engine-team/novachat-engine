/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.sendVerifyPhoneCode_handler.go
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
    "novachat_engine/service/sms"
)

//  account.sendVerifyPhoneCode#a5a356f9 phone_number:string settings:CodeSettings = auth.SentCode;
//
func (s *AccountServiceImpl) AccountSendVerifyPhoneCode(ctx context.Context, request *mtproto.TLAccountSendVerifyPhoneCode) (*mtproto.Auth_SentCode, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("AccountSendVerifyPhoneCode %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    if request.Settings.AllowFlashcall {
        request.Settings.AllowFlashcall = false
        log.Debugf("AccountSendVerifyPhoneCode - Verify Phone Coding")
    }

    if len(request.PhoneNumber) == 0 {
        return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
    }

    phoneHashCode := random2.RandomAlphaOrNumeric(32, true, false)
    phoneCode := random2.RandomAlphaOrNumeric(10, true, true)

    reply := mtproto.NewTLAuthSentCode(nil)
    reply.SetPhoneCodeHash(phoneHashCode)
    reply.SetTimeout(sms.TTL)

    if request.Settings != nil && request.Settings.AllowFlashcall {
        reply.SetNextType(constants.MakeAuthCode(constants.CodeTypeFlashCall))
        reply.SetType(constants.MakeSendCodeType(constants.CodeTypeFlashCall, sms.CodeLength, "*"))
    } else {
        reply.SetNextType(constants.MakeAuthCode(constants.CodeTypeApp))
        reply.SetType(constants.MakeSendCodeType(constants.CodeTypeSms, sms.CodeLength, "*"))
    }

    if s.conf.EnvMode != constants.EnvModelTest {
        sms.SendSmsVerifyCode(sms.CodeVerify, request.PhoneNumber, phoneCode, phoneHashCode)
    }

    return reply.To_Auth_SentCode(), nil
}
