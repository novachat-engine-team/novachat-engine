/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getPasswordSettings_handler.go
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

//  account.getPasswordSettings#9cd4eaf9 password:InputCheckPasswordSRP = account.PasswordSettings;
//  account.getPasswordSettings#bc8d11bb current_password_hash:bytes = account.PasswordSettings;
//
func (s *AccountServiceImpl) AccountGetPasswordSettings(ctx context.Context, request *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("AccountGetPasswordSettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl AccountGetPasswordSettings logic

    return nil, fmt.Errorf("%s", "Not impl AccountGetPasswordSettings")
}
