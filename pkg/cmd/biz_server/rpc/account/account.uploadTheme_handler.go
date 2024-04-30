/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.uploadTheme_handler.go
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

//  account.uploadTheme#1c3db333 flags:# file:InputFile thumb:flags.0?InputFile file_name:string mime_type:string = Document;
//
func (s *AccountServiceImpl) AccountUploadTheme(ctx context.Context, request *mtproto.TLAccountUploadTheme) (*mtproto.Document, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountUploadTheme %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountUploadTheme logic

	return nil, fmt.Errorf("%s", "Not impl AccountUploadTheme")
}
