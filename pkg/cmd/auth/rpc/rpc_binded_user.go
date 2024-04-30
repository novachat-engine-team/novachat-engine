/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	authService "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/log"
)

func (m *AuthImpl) ReqBindedUser(ctx context.Context, request *authService.BindedUser) (*types.Any, error) {

	log.Debugf("ReqBindedUser request:%+v", request)

	auth, err := m.accountAuthCore.AuthKey(request.PermAuthKeyId, request.SessionId, request.UserId)
	if err != nil {
		log.Errorf("ReqBindedUser request:%+v error:%s", request, err.Error())
		return nil, err
	}

	log.Debugf("ReqBindedUser request:%+v PermAuthKeyId:%d AuthKeyId:%d UserId:%d", request, auth.PermAuthKeyId, auth.AuthKeyId, auth.UserId)

	if request.UserId == 0 && auth.UserId != 0 {
		request.UserId = auth.UserId
		request.PermAuthKeyId = request.AuthKeyId
		resp, err1 := authService.GetAuthClientByAuthKey(auth.UserId).ReqBindedUser(ctx, request)
		if err1 != nil {
			log.Errorf("ReqBindedUser request:%+v userId:%d error:%s", request, auth.UserId, err.Error())
			return nil, err
		}
		_ = resp
		//return resp, nil
	}

	return types.MarshalAny(&mtproto.TLInt64{
		Value: auth.UserId,
	})
}
