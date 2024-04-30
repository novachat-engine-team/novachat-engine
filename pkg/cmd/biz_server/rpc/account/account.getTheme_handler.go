/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getTheme_handler.go
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

//  account.getTheme#8d9d742b format:string theme:InputTheme document_id:long = Theme;
//
func (s *AccountServiceImpl) AccountGetTheme(ctx context.Context, request *mtproto.TLAccountGetTheme) (*mtproto.Theme, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountGetTheme %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountGetTheme logic

	return nil, fmt.Errorf("%s", "Not impl AccountGetTheme")
}
