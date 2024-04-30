/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : folders.deleteFolder_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"time"
)

//  folders.deleteFolder#1c295881 folder_id:int = Updates;
//
func (s *FoldersServiceImpl) FoldersDeleteFolder(ctx context.Context, request *mtproto.TLFoldersDeleteFolder) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("FoldersDeleteFolder %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// unused
	return mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
