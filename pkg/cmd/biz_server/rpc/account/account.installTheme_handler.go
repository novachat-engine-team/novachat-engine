/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.installTheme_handler.go
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

//  account.installTheme#7ae43737 flags:# dark:flags.0?true format:flags.1?string theme:flags.1?InputTheme = Bool;
//
func (s *AccountServiceImpl) AccountInstallTheme(ctx context.Context, request *mtproto.TLAccountInstallTheme) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountInstallTheme %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountInstallTheme logic

	return mtproto.ToMTBool(true), nil
}
