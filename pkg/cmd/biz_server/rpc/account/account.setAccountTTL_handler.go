/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.setAccountTTL_handler.go
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

//  account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
//
func (s *AccountServiceImpl) AccountSetAccountTTL(ctx context.Context, request *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountSetAccountTTL %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var err error
	var ok bool
	ok, err = s.usersCore.UpdateAccountTTL(md.UserId, request.Ttl.Days)
	if err != nil {
		log.Errorf("AccountSetAccountTTL - request: %v error:%s", request, err.Error())
		return nil, err
	}

	return mtproto.ToMTBool(ok), nil
}
