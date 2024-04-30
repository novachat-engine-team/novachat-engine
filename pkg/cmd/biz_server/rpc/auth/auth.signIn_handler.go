/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.signIn_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	usersUtil "novachat_engine/service/account/users"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/account/phone"
	"novachat_engine/service/sms"
	"strings"
)

//(auth.signIn "79991234567" "2dc02d2cda9e615c84" "44444")
//=
//(auth.authorization
//expires:1403938438
//user:(userSelf
//id:603177
//first_name:"John"
//last_name:"Doe"
//phone:"79991234567"
//photo:(userProfilePhotoEmpty)
//status:(userStatusEmpty)
//inactive:(boolTrue)
//)
//)
//Name	Type	Description
//phone_number	string	Phone number in the international format
//phone_code_hash	string	SMS-message ID, obtained from auth.sendCode
//phone_code	string	Valid numerical code from the SMS-message
//Result
//Returns an auth.Authorization object with information on the new authorization.
//
//Possible errors
//Code	Type	Description
//400	PHONE_CODE_EMPTY	phone_code from the SMS is empty
//400	PHONE_CODE_EXPIRED	SMS expired
//400	PHONE_CODE_INVALID	Invalid SMS code was sent
//400	PHONE_NUMBER_INVALID	Invalid phone number
//400	PHONE_NUMBER_UNOCCUPIED	The code is valid but no user with the given number is registered

//  auth.signIn#bcd51581 phone_number:string phone_code_hash:string phone_code:string = auth.Authorization;
//
func (s *AuthServiceImpl) AuthSignIn(ctx context.Context, request *mtproto.TLAuthSignIn) (*mtproto.Auth_Authorization, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AuthSignIn %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.PhoneCode) == 0 {
		log.Errorf("AuthSignIn - request: %v error", request)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_EMPTY)
	}

	phoneNumberParser, err := phone.Parse(request.PhoneNumber)
	if err != nil {
		log.Errorf("AuthSignIn - request: %s verror:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
	}

	if s.conf.EnvMode != constants.EnvModelTest {
		err = sms.VerifySmsCode(sms.CodeSignIn, phoneNumberParser.NormalizeDigitsOnly(), request.PhoneCode, request.PhoneCodeHash)
		if err != nil {
			log.Errorf("AuthSignIn -  request: %v VerifySmsCode error:%s", request, err.Error())
			return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_BAD_REQUEST.ToInt32(), err.Error())
		}
	} else {
		if len(request.PhoneCodeHash) == 0 {
			log.Errorf("AuthSignIn - request: %v error", request)
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_HASH_EMPTY)
		}

		if strings.Compare(request.PhoneCode, sms.DefaultPhoneCode) != 0 {
			log.Errorf("AuthSignIn - request: %v error", request)
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_INVALID)
		}
	}

	userInfo, err := s.accountUsersCore.PhoneUser(phoneNumberParser.NormalizeDigitsOnly())
	if err != nil {
		log.Errorf("AuthSignIn - request: %v error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
	}
	if userInfo == nil || userInfo.Id == 0 {
		log.Debugf("AuthSignIn - request: %v ok!!!!! AuthUp", request)
		return mtproto.NewTLAuthAuthorizationSignUpRequired(nil).To_Auth_Authorization(), nil
	}
	user := usersUtil.UserCoreSelfUsers(usersUtil.UserCore2Users(userInfo))

	// forbided user
	if user.GetRestricted() == true {
		log.Errorf("AuthSignIn - metadata phoneNumber:%s, UserId:%d RpcErrorCode_FORBIDDEN_USER_RESTRICTED", request.PhoneNumber, user.GetId())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_FORBIDDEN_USER_RESTRICTED)
	}

	resp, err := authClient.GetAuthClientByAuthKey(md.PermAuthKeyId).ReqBindUser(context.Background(), &authClient.BindUser{
		AuthKeyId:     md.AuthKeyId,
		UserId:        constants.PeerTypeFromUserIDType32(user.GetId()).ToInt(),
		PermAuthKeyId: md.PermAuthKeyId,
	})
	if err != nil {
		log.Errorf("AuthSignUp - request: %v BindAuthKeyUser error:%s", request, err.Error())
		return nil, err
	}
	var ok mtproto.Bool
	err = types.UnmarshalAny(resp, &ok)
	if !mtproto.ToBool(&ok) {
		log.Warnf("AuthSignUp - request: %v AuthBindAuthKeyUser ok false", request)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_UNREGISTERED)
	}

	go func() {
		// PHONE_NUMBER_UNOCCUPIED unused
		sms.ClearSmsCode(sms.CodeSignIn, phoneNumberParser.NormalizeDigitsOnly(), request.PhoneCode, request.PhoneCodeHash)
	}()

	//if s.conf.BizLogic.KickOut && len(ok.AuthKeyIdList) > 0 {
	//	//TODO:(coder) sync kick out
	//}

	authorization := mtproto.NewTLAuthAuthorization(nil)
	user.Self = true
	authorization.SetUser(user)

	log.Infof("AuthSignIn - request: %v reply:%v AuthIn ok!!!!!", request, user)
	return authorization.To_Auth_Authorization(), nil
}
