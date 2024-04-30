/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : users.setSecureValueErrors_handler.go
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

//  users.setSecureValueErrors#90c894b5 id:InputUser errors:Vector<SecureValueError> = Bool;
//
func (s *UsersServiceImpl) UsersSetSecureValueErrors(ctx context.Context, request *mtproto.TLUsersSetSecureValueErrors) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("UsersSetSecureValueErrors %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl UsersSetSecureValueErrors logic

	return nil, fmt.Errorf("%s", "Not impl UsersSetSecureValueErrors")
}
