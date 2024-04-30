/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getContactSignUpNotification_handler.go
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

//  account.getContactSignUpNotification#9f07c728 = Bool;
//
func (s *AccountServiceImpl) AccountGetContactSignUpNotification(ctx context.Context, request *mtproto.TLAccountGetContactSignUpNotification) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AccountGetContactSignUpNotification %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	ok := mtproto.ToMTBool(true)
	log.Debugf("AccountGetContactSignUpNotification -request: %v ok:%v reply", request, mtproto.ToBool(ok))
	return ok, nil
}
