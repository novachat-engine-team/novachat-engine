/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.acceptAuthorization_handler.go
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

//  account.acceptAuthorization#e7027c94 bot_id:int scope:string public_key:string value_hashes:Vector<SecureValueHash> credentials:SecureCredentialsEncrypted = Bool;
//
func (s *AccountServiceImpl) AccountAcceptAuthorization(ctx context.Context, request *mtproto.TLAccountAcceptAuthorization) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountAcceptAuthorization %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountAcceptAuthorization logic

	return nil, fmt.Errorf("%s", "Not impl AccountAcceptAuthorization")
}
