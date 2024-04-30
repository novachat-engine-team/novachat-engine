/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.sendInvites_handler.go
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

//  auth.sendInvites#771c1d97 phone_numbers:Vector<string> message:string = Bool;
//
func (s *AuthServiceImpl) AuthSendInvites(ctx context.Context, request *mtproto.TLAuthSendInvites) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthSendInvites %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AuthSendInvites logic

	return mtproto.ToMTBool(true), nil
}
