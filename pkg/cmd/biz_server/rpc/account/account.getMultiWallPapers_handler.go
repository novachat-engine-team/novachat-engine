/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getMultiWallPapers_handler.go
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

//  account.getMultiWallPapers#65ad71dc wallpapers:Vector<InputWallPaper> = Vector<WallPaper>;
//
func (s *AccountServiceImpl) AccountGetMultiWallPapers(ctx context.Context, request *mtproto.TLAccountGetMultiWallPapers) (*mtproto.Vector_WallPaper, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountGetMultiWallPapers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountGetMultiWallPapers logic

	return nil, fmt.Errorf("%s", "Not impl AccountGetMultiWallPapers")
}
