/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.confirmPhone_handler.go
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

//  account.confirmPhone#5f2178c3 phone_code_hash:string phone_code:string = Bool;
//
func (s *AccountServiceImpl) AccountConfirmPhone(ctx context.Context, request *mtproto.TLAccountConfirmPhone) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountConfirmPhone %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountConfirmPhone logic

	return nil, fmt.Errorf("%s", "Not impl AccountConfirmPhone")
}
