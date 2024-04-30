/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getSecureValue_handler.go
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

//  account.getSecureValue#73665bc2 types:Vector<SecureValueType> = Vector<SecureValue>;
//
func (s *AccountServiceImpl) AccountGetSecureValue(ctx context.Context, request *mtproto.TLAccountGetSecureValue) (*mtproto.Vector_SecureValue, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("AccountGetSecureValue %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl AccountGetSecureValue logic

    return nil, fmt.Errorf("%s", "Not impl AccountGetSecureValue")
}
