/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.importLoginToken_handler.go
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

//  auth.importLoginToken#95ac5ce4 token:bytes = auth.LoginToken;
//
func (s *AuthServiceImpl) AuthImportLoginToken(ctx context.Context, request *mtproto.TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthImportLoginToken %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AuthImportLoginToken logic

	return nil, fmt.Errorf("%s", "Not impl AuthImportLoginToken")
}
