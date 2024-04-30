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

func (m *AuthImpl) ReqBindUser(ctx context.Context, request *authService.BindUser) (*types.Any, error) {
	ok, err := m.accountAuthCore.BindUser(request.PermAuthKeyId, request.UserId)
	if err != nil {
		log.Errorf("ReqBindUser request:%+v error:%s", request, err.Error())
		return nil, err
	}

	log.Debugf("ReqBindUser request:%+v ok:%v", request, ok)
	return types.MarshalAny(mtproto.ToMTBool(ok))
}
