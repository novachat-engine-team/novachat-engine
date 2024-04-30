/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.uploadWallPaper_handler.go
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

//  account.uploadWallPaper#dd853661 file:InputFile mime_type:string settings:WallPaperSettings = WallPaper;
//
func (s *AccountServiceImpl) AccountUploadWallPaper(ctx context.Context, request *mtproto.TLAccountUploadWallPaper) (*mtproto.WallPaper, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountUploadWallPaper %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountUploadWallPaper logic

	return mtproto.NewTLWallPaperNoFile(nil).To_WallPaper(), nil
}
