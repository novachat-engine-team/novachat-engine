/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.sendVerifyEmailCode_handler.go
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

//  account.sendVerifyEmailCode#7011509f email:string = account.SentEmailCode;
//
func (s *AccountServiceImpl) AccountSendVerifyEmailCode(ctx context.Context, request *mtproto.TLAccountSendVerifyEmailCode) (*mtproto.Account_SentEmailCode, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountSendVerifyEmailCode %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountSendVerifyEmailCode logic

	return nil, fmt.Errorf("%s", "Not impl AccountSendVerifyEmailCode")
}
