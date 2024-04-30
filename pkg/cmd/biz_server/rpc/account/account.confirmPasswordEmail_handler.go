/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.confirmPasswordEmail_handler.go
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

//  account.confirmPasswordEmail#8fdf1920 code:string = Bool;
//
func (s *AccountServiceImpl) AccountConfirmPasswordEmail(ctx context.Context, request *mtproto.TLAccountConfirmPasswordEmail) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountConfirmPasswordEmail %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountConfirmPasswordEmail logic

	return nil, fmt.Errorf("%s", "Not impl AccountConfirmPasswordEmail")
}
