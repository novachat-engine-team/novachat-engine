/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.resendPasswordEmail_handler.go
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

//  account.resendPasswordEmail#7a7f2a15 = Bool;
//
func (s *AccountServiceImpl) AccountResendPasswordEmail(ctx context.Context, request *mtproto.TLAccountResendPasswordEmail) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountResendPasswordEmail %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountResendPasswordEmail logic

	return nil, fmt.Errorf("%s", "Not impl AccountResendPasswordEmail")
}
