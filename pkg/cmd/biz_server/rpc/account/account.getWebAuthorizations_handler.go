/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getWebAuthorizations_handler.go
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

//  account.getWebAuthorizations#182e6d6f = account.WebAuthorizations;
//
func (s *AccountServiceImpl) AccountGetWebAuthorizations(ctx context.Context, request *mtproto.TLAccountGetWebAuthorizations) (*mtproto.Account_WebAuthorizations, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountGetWebAuthorizations %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	reply := mtproto.NewTLAccountWebAuthorizations(&mtproto.Account_WebAuthorizations{
		Authorizations: []*mtproto.WebAuthorization{},
		Users:          []*mtproto.User{},
	}).To_Account_WebAuthorizations()

	log.Debugf("AccountGetWebAuthorizations - reply:%v ok!!!!!!!!!!!!!!! ", reply)
	return reply, nil
}
