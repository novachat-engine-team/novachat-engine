/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.saveSecureValue_handler.go
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

//  account.saveSecureValue#899fe31d value:InputSecureValue secure_secret_id:long = SecureValue;
//
func (s *AccountServiceImpl) AccountSaveSecureValue(ctx context.Context, request *mtproto.TLAccountSaveSecureValue) (*mtproto.SecureValue, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountSaveSecureValue %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountSaveSecureValue logic

	return nil, fmt.Errorf("%s", "Not impl AccountSaveSecureValue")
}
