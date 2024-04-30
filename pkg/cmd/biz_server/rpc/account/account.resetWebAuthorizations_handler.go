/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.resetWebAuthorizations_handler.go
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

//  account.resetWebAuthorizations#682d2594 = Bool;
//
func (s *AccountServiceImpl) AccountResetWebAuthorizations(ctx context.Context, request *mtproto.TLAccountResetWebAuthorizations) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountResetWebAuthorizations %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountResetWebAuthorizations logic

	return nil, fmt.Errorf("%s", "Not impl AccountResetWebAuthorizations")
}
