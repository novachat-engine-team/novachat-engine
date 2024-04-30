/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.saveWallPaper_handler.go
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

//  account.saveWallPaper#6c5a5b37 wallpaper:InputWallPaper unsave:Bool settings:WallPaperSettings = Bool;
//
func (s *AccountServiceImpl) AccountSaveWallPaper(ctx context.Context, request *mtproto.TLAccountSaveWallPaper) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountSaveWallPaper %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountSaveWallPaper logic

	return mtproto.ToMTBool(true), nil
}
