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

func (m *AuthImpl) ReqBindTempAuthKeyAuthKeyId(ctx context.Context, request *authService.BindTempAuthKeyAuthKeyId) (*types.Any, error) {
	ok, err := m.accountAuthCore.BindTempAuthKeyAuthKeyId(request.AuthKeyId, request.PermAuthKeyId, request.SessionId)
	if err != nil {
		log.Errorf("ReqBindTempAuthKeyAuthKeyId request:%+v error:%s", request, err.Error())
		return nil, err
	}

	log.Debugf("ReqBindTempAuthKeyAuthKeyId request:%+v ok:%v", request, ok)
	return types.MarshalAny(mtproto.ToMTBool(ok))
}
