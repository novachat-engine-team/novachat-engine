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
	authService "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/log"
)

// ReqAuthKey
// return AuthKey and salt
// but salt not to next
func (m *AuthImpl) ReqAuthKey(ctx context.Context, request *authService.AuthKey) (*authService.AuthKeyInfo, error) {
	log.Debugf("ReqAuthKey request.AuthKeyId %d", request.AuthKeyId)
	auth, err := m.accountAuthCore.AuthKey(request.AuthKeyId, 0, 0)
	if err != nil {
		log.Errorf("ReqAuthKey request:%+v error:%s", request, err.Error())
		return nil, err
	}
	authKeyInfo := authService.AuthKeyInfo{}
	if auth == nil {
		log.Debugf("ReqAuthKey request:%+v auth is empty", request)
		return &authKeyInfo, nil
	}

	//log.Debugf("ReqAuthKey request authKeyId%d authKey:%s ", auth.AuthKeyId, hex.EncodeToString(auth.AuthKey))
	log.Debugf("ReqAuthKey request authKeyId%d", auth.AuthKeyId)

	authKeyInfo.AuthKeyId = auth.AuthKeyId
	authKeyInfo.ExpiresIn = auth.ExpiresIn
	authKeyInfo.ValidSince = auth.ValidSince
	authKeyInfo.ValidUntil = auth.ValidUntil
	authKeyInfo.Salt = auth.Salt
	authKeyInfo.TempAuthKey = auth.TempAuthKey
	authKeyInfo.AuthKey = auth.AuthKey
	authKeyInfo.Layer = auth.Layer
	authKeyInfo.PermAuthKeyId = auth.PermAuthKeyId
	return &authKeyInfo, nil
}
