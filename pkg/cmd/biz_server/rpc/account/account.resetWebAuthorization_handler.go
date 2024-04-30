/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.resetWebAuthorization_handler.go
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

//  account.resetWebAuthorization#2d01b9ef hash:long = Bool;
//
func (s *AccountServiceImpl) AccountResetWebAuthorization(ctx context.Context, request *mtproto.TLAccountResetWebAuthorization) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountResetWebAuthorization %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountResetWebAuthorization logic

	return nil, fmt.Errorf("%s", "Not impl AccountResetWebAuthorization")
}
