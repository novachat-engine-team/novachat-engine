/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.recoverPassword_handler.go
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

//  auth.recoverPassword#4ea56e92 code:string = auth.Authorization;
//
func (s *AuthServiceImpl) AuthRecoverPassword(ctx context.Context, request *mtproto.TLAuthRecoverPassword) (*mtproto.Auth_Authorization, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthRecoverPassword %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AuthRecoverPassword logic

	return nil, fmt.Errorf("%s", "Not impl AuthRecoverPassword")
}
