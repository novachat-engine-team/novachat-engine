/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.updateDeviceLocked_handler.go
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

//  account.updateDeviceLocked#38df3532 period:int = Bool;
//
func (s *AccountServiceImpl) AccountUpdateDeviceLocked(ctx context.Context, request *mtproto.TLAccountUpdateDeviceLocked) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountUpdateDeviceLocked %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountUpdateDeviceLocked logic

	return mtproto.ToMTBool(false), nil
}
