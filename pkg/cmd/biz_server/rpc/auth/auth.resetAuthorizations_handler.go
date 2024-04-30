/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.resetAuthorizations_handler.go
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

//  auth.resetAuthorizations#9fab0d1a = Bool;
//
func (s *AuthServiceImpl) AuthResetAuthorizations(ctx context.Context, request *mtproto.TLAuthResetAuthorizations) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthResetAuthorizations %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AuthResetAuthorizations logic

	return nil, fmt.Errorf("%s", "Not impl AuthResetAuthorizations")
}
