/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.installWallPaper_handler.go
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

//  account.installWallPaper#feed5769 wallpaper:InputWallPaper settings:WallPaperSettings = Bool;
//
func (s *AccountServiceImpl) AccountInstallWallPaper(ctx context.Context, request *mtproto.TLAccountInstallWallPaper) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountInstallWallPaper %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountInstallWallPaper logic

	return mtproto.ToMTBool(true), nil
}
