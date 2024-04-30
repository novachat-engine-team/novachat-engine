/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getContentSettings_handler.go
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

//  account.getContentSettings#8b9b4dae = account.ContentSettings;
//
func (s *AccountServiceImpl) AccountGetContentSettings(ctx context.Context, request *mtproto.TLAccountGetContentSettings) (*mtproto.Account_ContentSettings, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountGetContentSettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	return mtproto.NewTLAccountContentSettings(&mtproto.Account_ContentSettings{
		SensitiveEnabled:   false,
		SensitiveCanChange: false,
	}).To_Account_ContentSettings(), nil
}
