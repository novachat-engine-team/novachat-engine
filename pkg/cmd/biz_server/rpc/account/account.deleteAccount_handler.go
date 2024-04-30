/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.deleteAccount_handler.go
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

//  account.deleteAccount#418d4e0b reason:string = Bool;
//
func (s *AccountServiceImpl) AccountDeleteAccount(ctx context.Context, request *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountDeleteAccount %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// check session is longest
	// return 2FA_RECENT_CONFIRM
	ok, err := s.usersCore.DeleteAccount(md.UserId, request.Reason)
	if err != nil {
		return nil, err
	}

	return mtproto.ToMTBool(ok), err
}
