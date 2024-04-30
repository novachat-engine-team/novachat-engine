/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.verifyEmail_handler.go
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

//  account.verifyEmail#ecba39db email:string code:string = Bool;
//
func (s *AccountServiceImpl) AccountVerifyEmail(ctx context.Context, request *mtproto.TLAccountVerifyEmail) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountVerifyEmail %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountVerifyEmail logic

	return nil, fmt.Errorf("%s", "Not impl AccountVerifyEmail")
}
