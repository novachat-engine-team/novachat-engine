/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.registerDevice_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"encoding/hex"
	"novachat_engine/mtproto"
	authService "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
)

//  account.registerDevice#68976c6f flags:# no_muted:flags.0?true token_type:int token:string app_sandbox:Bool secret:bytes other_uids:Vector<int> = Bool;
//  account.registerDevice#637ea878 token_type:int token:string = Bool;
//
func (s *AccountServiceImpl) AccountRegisterDevice(ctx context.Context, request *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	var err error
	_ = md

	if len(request.Token) == 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_TOKEN_INVALID)
		log.Errorf("AccountRegisterDevice - request: %v error:%s", request, err.Error())
		return nil, err
	}

	var ok bool
	tokenType := constants.RegisterDeviceTypeFromInt32(request.TokenType)
	if constants.RegisterDeviceTypeDefault == tokenType { //|| constants.RegisterDeviceTypeFCM == tokenType {
		ok = true
		log.Infof("AccountRegisterDevice - request: %v RegisterDeviceType:%v reply:%v", request, tokenType, true)
	} else {
		otherUids := make([]int64, 0, len(request.OtherUids))
		for _, v := range request.OtherUids {
			otherUids = append(otherUids, int64(v))
		}

		_, err = authService.GetAuthClientByAuthKey(md.UserId).ReqRegisterDevice(context.Background(), &authService.UserDevice{
			UserId:        md.UserId,
			AuthKeyId:     md.AuthKeyId,
			PermAuthKeyId: md.PermAuthKeyId,
			SessionId:     md.SessionId,
			Device: &authService.Device{
				TokenType:    request.TokenType,
				AppSandBox:   mtproto.ToBool(request.AppSandbox),
				Secret:       hex.EncodeToString(request.Secret),
				NoMuted:      false,
				DeviceModel:  md.DeviceModel,
				LangPack:     md.LangPack,
				OtherUidList: otherUids,
			},
		})
		if err != nil {
			log.Warnf("AccountRegisterDevice - request: %v RegisterDeviceType:%v reply:%v error:%s", request, tokenType, ok, err.Error())
		} else {
			log.Infof("AccountRegisterDevice - request: %v RegisterDeviceType:%v reply:%v", request, tokenType, ok)
		}
	}
	return mtproto.ToMTBool(ok), nil
}
