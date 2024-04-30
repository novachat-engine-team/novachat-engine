/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getAccountTTL_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
)

//  account.getAccountTTL#8fc711d = AccountDaysTTL;
//
func (s *AccountServiceImpl) AccountGetAccountTTL(ctx context.Context, request *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountGetAccountTTL %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	userInfo, err := s.accountUsersCore.UserData(md.UserId)
	if err != nil {
		log.Errorf("AccountGetAccountTTL - request: %v GetByUserId error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}

	return mtproto.NewTLAccountDaysTTL(&mtproto.AccountDaysTTL{
		Days: userInfo.DaysTtl,
	}).To_AccountDaysTTL(), nil
}
