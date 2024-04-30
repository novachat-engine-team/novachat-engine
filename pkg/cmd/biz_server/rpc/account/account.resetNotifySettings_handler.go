/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.resetNotifySettings_handler.go
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

//  account.resetNotifySettings#db7e1747 = Bool;
//
func (s *AccountServiceImpl) AccountResetNotifySettings(ctx context.Context, request *mtproto.TLAccountResetNotifySettings) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountResetNotifySettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	ok, err := s.accountSettingCore.ResetNotifySetting(md.UserId)
	if err != nil {
		log.Errorf("AccountResetNotifySettings - request: %v error:%s", request, err.Error())
		return nil, err
	}

	//TODO:Coder
	//Sync to more devices
	return mtproto.ToMTBool(ok), nil
}
