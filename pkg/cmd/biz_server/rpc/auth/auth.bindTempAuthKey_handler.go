/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.bindTempAuthKey_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  auth.bindTempAuthKey#cdd42a05 perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
//
func (s *AuthServiceImpl) AuthBindTempAuthKey(ctx context.Context, request *mtproto.TLAuthBindTempAuthKey) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthBindTempAuthKey %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// AuthBindTempAuthKey on session

	return mtproto.ToMTBool(true), nil
}
