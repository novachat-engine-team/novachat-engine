/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getWallPapers_handler.go
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

//  account.getWallPapers#aabb1763 hash:int = account.WallPapers;
//  account.getWallPapers#c04cfac2 = Vector<WallPaper>;
//
func (s *AccountServiceImpl) AccountGetWallPapers(ctx context.Context, request *mtproto.TLAccountGetWallPapers) (*mtproto.Response_AccountGetWallPapers, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountGetWallPapers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	ret := &mtproto.Response_AccountGetWallPapers{
		AccountGetWallPapersc04Cfac2: &mtproto.Vector_WallPaper{},
		AccountGetWallPapersaabb1763: mtproto.NewTLAccountWallPapers(nil).To_Account_WallPapers(),
	}

	if md.Layer <= 71 {
		ret.Cmd = mtproto.Cmd_AccountGetWallPapers.ToUInt32()
	} else {
		ret.Cmd = mtproto.Cmd_AccountGetWallPapersaabb1763.ToUInt32()
	}
	// Impl AccountGetWallPapers logic

	return ret, nil
}
