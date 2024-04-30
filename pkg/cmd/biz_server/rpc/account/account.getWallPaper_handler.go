/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getWallPaper_handler.go
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

//  account.getWallPaper#fc8ddbea wallpaper:InputWallPaper = WallPaper;
//
func (s *AccountServiceImpl) AccountGetWallPaper(ctx context.Context, request *mtproto.TLAccountGetWallPaper) (*mtproto.WallPaper, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountGetWallPaper %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountGetWallPaper logic

	return nil, fmt.Errorf("%s", "Not impl AccountGetWallPaper")
}
