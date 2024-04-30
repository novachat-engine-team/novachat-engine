/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.signUp_handler.go
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

/*
(auth.signUp "79991234567" "2dc02d2cda9e615c84" "44444" "John" "Doe")
=
(auth.authorization
  expires:1403938438
  user:(userSelf
    id:603177
    first_name:"John"
    last_name:"Doe"
    phone:"79991234567"
    photo:(userProfilePhotoEmpty)
    status:(userStatusEmpty)
    inactive:(boolTrue)
  )
)
*/

//  auth.signUp#80eee427 phone_number:string phone_code_hash:string first_name:string last_name:string = auth.Authorization;
//  auth.signUp#1b067634 phone_number:string phone_code_hash:string phone_code:string first_name:string last_name:string = auth.Authorization;
//
func (s *AuthServiceImpl) AuthSignUp(ctx context.Context, request *mtproto.TLAuthSignUp) (*mtproto.Auth_Authorization, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AuthSignUp %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	phoneNumberParser, err := phone.Parse(request.PhoneNumber)
	if err != nil {
		log.Errorf("AuthSignUp - request: %v error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
	}

	if phoneNumberParser.IsVirtualPhone() {
		if s.conf.EnvMode != constants.EnvModelTest {
			log.Errorf("AuthSignUp - request: %v", request)
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
		}
	}

	if s.conf.EnvMode != constants.EnvModelTest {
		err = sms.VerifySmsCode(sms.CodeSignIn, phoneNumberParser.NormalizeDigitsOnly(), request.PhoneCode, request.PhoneCodeHash)
		if err != nil {
			log.Errorf("AuthSignUp - request: %v VerifySmsCode error:%s", request, err.Error())
			return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
		}
	} else {
		if len(request.PhoneCodeHash) == 0 {
			log.Errorf("AuthSignUp - request: %v error", request)
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_HASH_EMPTY)
		}

		if len(request.PhoneCode) > 0 && strings.Compare(request.PhoneCode, sms.DefaultPhoneCode) != 0 {
			log.Errorf("AuthSignUp - request: %v error", request)
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_INVALID)
		}
	}

	if len(request.FirstName) == 0 {
		log.Errorf("AuthSignUp - request: %v error", request)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_FIRSTNAME_INVALID)
	}

	userInfo, err := s.accountUsersCore.PhoneUser(phoneNumberParser.NormalizeDigitsOnly())
	if err != nil {
		log.Errorf("AuthSignUp - request: %v error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
	}
	if userInfo.Id == 0 {
		userInfo, err = s.usersCore.UserCreate(phoneNumberParser.NormalizeDigitsOnly(),
			request.FirstName,
			request.LastName,
		)
		if err != nil {
			log.Errorf("AuthSignUp - request: %v error:%s", request, err.Error())
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
		}
	}

	go func() {
		// PHONE_NUMBER_UNOCCUPIED unused
		err = sms.ClearSmsCode(sms.CodeSignIn, phoneNumberParser.NormalizeDigitsOnly(), request.PhoneCode, request.PhoneCodeHash)
		_ = err
	}()

	user := usersUtil.UserCoreSelfUsers(usersUtil.UserCore2Users(userInfo))
	// forbided user
	if user.GetRestricted() == true {
		log.Errorf("AuthSignUp - metadata phoneNumber:%s, UserId:%d RpcErrorCode_FORBIDDEN_USER_RESTRICTED", request.PhoneNumber, user.GetId())
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

	//if s.conf.BizLogic.KickOut && len(ok.AuthKeyIdList) > 0 {
	//	//TODO:(coder) sync kick out
	//}

	authorization := mtproto.NewTLAuthAuthorization(nil)
	user.Self = true
	authorization.SetUser(user)
	log.Infof("AuthSignUp %v, request: %v reply:%v ok!!!!!", md, request, user)
	return authorization.To_Auth_Authorization(), nil
}
