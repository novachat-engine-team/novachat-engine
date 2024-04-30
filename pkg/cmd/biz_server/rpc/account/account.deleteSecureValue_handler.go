/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.deleteSecureValue_handler.go
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

//  account.deleteSecureValue#b880bc4b types:Vector<SecureValueType> = Bool;
//
func (s *AccountServiceImpl) AccountDeleteSecureValue(ctx context.Context, request *mtproto.TLAccountDeleteSecureValue) (*mtproto.Bool, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("AccountDeleteSecureValue %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl AccountDeleteSecureValue logic

    return nil, fmt.Errorf("%s", "Not impl AccountDeleteSecureValue")
}
