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
	dataAuth "novachat_engine/service/data/auth"
	"time"
)

// ReqUpdateAuthKey
// update/save AuthKey
func (m *AuthImpl) ReqUpdateAuthKey(ctx context.Context, request *authService.UpdateAuthKey) (*types.Any, error) {
	ok, err := m.accountAuthCore.UpdateAuthKey(&dataAuth.Auth{
		AuthKeyId:   request.AuthKeyId,
		AuthKey:     request.AuthKey,
		ExpiresIn:   request.ExpiresIn,
		ValidSince:  request.ValidSince,
		ValidUntil:  request.ValidUntil,
		Salt:        request.Salt,
		TempAuthKey: request.TempAuthKey,
		Date:        int32(time.Now().Unix()),
	})
	if err != nil {
		log.Errorf("ReqUpdateAuthKey request:%+v error:%s", request, err.Error())
		return nil, err
	}

	//log.Debugf("ReqUpdateAuthKey request authKeyId:%d authKey:%s ok:%v", request.AuthKeyId, hex.EncodeToString(request.AuthKey), ok)
	log.Debugf("ReqUpdateAuthKey request authKeyId:%d ok:%v", request.AuthKeyId, ok)
	return types.MarshalAny(mtproto.ToMTBool(ok))
}
