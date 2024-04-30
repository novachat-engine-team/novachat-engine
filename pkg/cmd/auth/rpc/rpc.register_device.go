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

func (m *AuthImpl) ReqRegisterDevice(ctx context.Context, request *authService.UserDevice) (*types.Any, error) {
	log.Debugf("ReqRegisterDevice device:%+v", request)

	_, err := m.accountAuthCore.RegisterDevice(request.UserId, request.AuthKeyId, request.Device)
	if err != nil {
		log.Errorf("ReqRegisterDevice device:%+v error:%s", request, err.Error())
	}

	_, err = m.accountAuthCore.AuthKey(request.AuthKeyId, request.SessionId, request.SessionId)
	if err != nil {
		log.Errorf("ReqRegisterDevice device:%+v error:%s", request, err.Error())
	}

	return types.MarshalAny(mtproto.ToMTBool(true))
}
