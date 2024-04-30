/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.createTheme_handler.go
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

//  account.createTheme#8432c21f flags:# slug:string title:string document:flags.2?InputDocument settings:flags.3?InputThemeSettings = Theme;
//
func (s *AccountServiceImpl) AccountCreateTheme(ctx context.Context, request *mtproto.TLAccountCreateTheme) (*mtproto.Theme, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountCreateTheme %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountCreateTheme logic

	return nil, fmt.Errorf("%s", "Not impl AccountCreateTheme")
}
