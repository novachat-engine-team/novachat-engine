/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone_service_impl.go
 * @Desc :
 *
 */

package rpc

import (
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/cmd/biz_server/conf"
	accountUsers "novachat_engine/service/account/users"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/users/sql_core"
)

type PhoneServiceImpl struct {
	conf *conf.Config
	//phoneCallConfig *phone_call.PhoneCallConfig
	//phoneCallCore   *phone_call.PhoneCallCore
	accountCore      *account.Core
	userCore         *sql_core.Core
	accountUsersCore *accountUsers.Core
}

func NewPhoneServiceImpl(conf *conf.Config) *PhoneServiceImpl {
	impl := &PhoneServiceImpl{
		conf: conf,
		//phoneCallConfig: phone_call.NewPhoneCallConfig(conf.RelayConfig.PhoneCallWebRTC),
		//phoneCallCore:   phone_call.NewPhoneCallCore(conf.RelayConfig.PhoneCallWebRTC, conf.RelayConfig.P2P, conf.RelayConfig.LibVersion, conf.Server.ServerId, cache.GetRedisClient()),
		accountCore: account.NewAccountCore(cache.GetRedisClient()),
	}

	impl.userCore = sql_core.NewUsersCore(&conf.MySQL)
	impl.accountUsersCore = accountUsers.NewUsersCore(impl.accountCore, nil, impl.userCore)

	return impl
}
