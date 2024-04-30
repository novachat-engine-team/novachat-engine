/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.saveTheme_handler.go
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

//  account.saveTheme#f257106c theme:InputTheme unsave:Bool = Bool;
//
func (s *AccountServiceImpl) AccountSaveTheme(ctx context.Context, request *mtproto.TLAccountSaveTheme) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountSaveTheme %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountSaveTheme logic

	return mtproto.ToMTBool(true), nil
}
