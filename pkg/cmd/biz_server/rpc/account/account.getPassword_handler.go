/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getPassword_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/account/password"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/rpc_context"
)

//  account.getPassword#548a30f5 = account.Password;
//
func (s *AccountServiceImpl) AccountGetPassword(ctx context.Context, request *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountGetPassword %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	ctx, cancel := rpc_context.Context(
		context.Background(),
		rpc_context.WithTimeout(constants.RpcDefaultTimeout),
		rpc_context.WithBaseMD(md),
		rpc_context.WithServerTrace(
			rpc_context.FormatServerTrace(s.conf.Server.Addr,
				rpc_context.RunFuncName()),
		),
	)
	authKeyInfo, err := authClient.GetAuthClientByAuthKey(md.AuthKeyId).ReqAuthKey(ctx, &authClient.AuthKey{AuthKeyId: md.AuthKeyId})
	cancel()
	if err != nil {
		log.Errorf("AccountGetPassword - , request: %v error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
	}
	if md.Layer <= 74 {
		return mtproto.NewTLAccountNoPassword(&mtproto.Account_Password{
			NewSalt:                 authKeyInfo.AuthKey,
			EmailUnconfirmedPattern: "",
		}).To_Account_Password(), nil
	} else {
		return password.ToNoPassword(), nil
	}
}
