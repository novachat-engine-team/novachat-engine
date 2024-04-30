/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getTmpPassword_handler.go
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

//  account.getTmpPassword#449e0b51 password:InputCheckPasswordSRP period:int = account.TmpPassword;
//  account.getTmpPassword#4a82327e password_hash:bytes period:int = account.TmpPassword;
//
func (s *AccountServiceImpl) AccountGetTmpPassword(ctx context.Context, request *mtproto.TLAccountGetTmpPassword) (*mtproto.Account_TmpPassword, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountGetTmpPassword %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountGetTmpPassword logic

	return nil, fmt.Errorf("%s", "Not impl AccountGetTmpPassword")
}
