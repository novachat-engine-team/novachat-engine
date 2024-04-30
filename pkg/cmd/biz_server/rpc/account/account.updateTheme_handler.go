/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.updateTheme_handler.go
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

//  account.updateTheme#5cb367d5 flags:# format:string theme:InputTheme slug:flags.0?string title:flags.1?string document:flags.2?InputDocument settings:flags.3?InputThemeSettings = Theme;
//
func (s *AccountServiceImpl) AccountUpdateTheme(ctx context.Context, request *mtproto.TLAccountUpdateTheme) (*mtproto.Theme, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("AccountUpdateTheme %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl AccountUpdateTheme logic

    return nil, fmt.Errorf("%s", "Not impl AccountUpdateTheme")
}
