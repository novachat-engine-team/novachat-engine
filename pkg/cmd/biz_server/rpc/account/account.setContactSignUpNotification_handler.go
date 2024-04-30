/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.setContactSignUpNotification_handler.go
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

//  account.setContactSignUpNotification#cff43f61 silent:Bool = Bool;
//
func (s *AccountServiceImpl) AccountSetContactSignUpNotification(ctx context.Context, request *mtproto.TLAccountSetContactSignUpNotification) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AccountSetContactSignUpNotification %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	silent := int32(0)
	if mtproto.ToBool(request.Silent) == true {
		silent = 1
	}
	_ = silent
	return mtproto.ToMTBool(true), nil
}
