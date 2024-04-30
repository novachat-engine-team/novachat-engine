/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.checkPassword_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//400	PASSWORD_HASH_INVALID	The provided password isn't valid
//400	SRP_ID_INVALID	Invalid SRP ID provided
//400	SRP_PASSWORD_CHANGED	Password has changed

//  auth.checkPassword#d18b4d16 password:InputCheckPasswordSRP = auth.Authorization;
//  auth.checkPassword#a63011e password_hash:bytes = auth.Authorization;
//
func (s *AuthServiceImpl) AuthCheckPassword(ctx context.Context, request *mtproto.TLAuthCheckPassword) (*mtproto.Auth_Authorization, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthCheckPassword %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AuthCheckPassword logic

	return nil, fmt.Errorf("%s", "Not impl AuthCheckPassword")
}
