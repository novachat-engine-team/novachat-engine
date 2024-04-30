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

func (m *AuthImpl) ReqClose(ctx context.Context, request *authService.Close) (*types.Any, error) {
	log.Debugf("ReqClose authKeyId:%d connId:%d", request.AuthKeyId, request.ConnId)
	m.accountAuthCore.Close(request.AuthKeyId, request.ConnId)
	return types.MarshalAny(mtproto.ToMTBool(true))
}
