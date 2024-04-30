/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getAllSecureValues_handler.go
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

//  account.getAllSecureValues#b288bc7d = Vector<SecureValue>;
//
func (s *AccountServiceImpl) AccountGetAllSecureValues(ctx context.Context, request *mtproto.TLAccountGetAllSecureValues) (*mtproto.Vector_SecureValue, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("AccountGetAllSecureValues %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl AccountGetAllSecureValues logic

    return nil, fmt.Errorf("%s", "Not impl AccountGetAllSecureValues")
}
