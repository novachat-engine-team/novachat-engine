/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.resetWallPapers_handler.go
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

//  account.resetWallPapers#bb3b9804 = Bool;
//
func (s *AccountServiceImpl) AccountResetWallPapers(ctx context.Context, request *mtproto.TLAccountResetWallPapers) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountResetWallPapers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountResetWallPapers logic

	return mtproto.ToMTBool(true), nil
}
