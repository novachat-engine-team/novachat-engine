/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.requestCall_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/app/support"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
)

//  phone.requestCall#42ff96ed flags:# video:flags.0?true user_id:InputUser random_id:int g_a_hash:bytes protocol:PhoneCallProtocol = phone.PhoneCall;
//  phone.requestCall#5b95b3d4 user_id:InputUser random_id:int g_a_hash:bytes protocol:PhoneCallProtocol = phone.PhoneCall;
//
func (s *PhoneServiceImpl) PhoneRequestCall(ctx context.Context, request *mtproto.TLPhoneRequestCall) (*mtproto.Phone_PhoneCall, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("PhoneRequestCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var err error
	//var phonePhoneCall *mtproto.TLPhonePhoneCall
	inputUser := input.MakeInputUser(request.UserId)
	if constants.InputUserTypeUser != inputUser.GetInputUserType() {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USER_ID_INVALID)
		log.Errorf("PhoneRequestCall %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	if support.IsSupportUser(inputUser.GetId()) {
		log.Warnf("PhoneRequestCall %v, request: %v error: support", metadata.RpcMetaDataDebug(md), request)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED)
	}

	//USER_IS_BLOCKED
	//USER_PRIVACY_RESTRICTED
	userInfo, err := s.accountUsersCore.UserData(md.UserId)
	if err != nil {
		log.Errorf("UpdatesGetDifference - request: %v GetByUserId error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}
	if userInfo.GetDeleted() {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USER_IS_BLOCKED)
		log.Errorf("PhoneRequestCall %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	panic("PhoneRequestCall")
	//if phonePhoneCall, err = s.phoneCallCore.RequestPhoneCall(md.AuthKeyId, md.UserId, inputUser.GetId(), request.Video, request.RandomId, request.GAHash, request.Protocol); err != nil {
	//	log.Errorf("PhoneRequestCall %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//	return nil, err
	//}
	//
	//log.Debugf("PhoneRequestCall %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	//return phonePhoneCall.To_Phone_PhoneCall(), nil
}
