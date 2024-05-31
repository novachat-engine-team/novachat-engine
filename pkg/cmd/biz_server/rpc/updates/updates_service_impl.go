/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : updates_service_impl.go
 * @Desc :
 *
 */

package rpc

import (
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/cmd/biz_server/conf"
	accountContact "novachat_engine/service/account/contact"
	accountMessage "novachat_engine/service/account/message"
	accountSetting "novachat_engine/service/account/setting"
	accountUpdate "novachat_engine/service/account/update"
	accountUsers "novachat_engine/service/account/users"
	"novachat_engine/service/app/phone_call"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/contact"
	"novachat_engine/service/core/account/users/sql_core"
	"novachat_engine/service/core/message/message"
	"novachat_engine/service/core/setting"
	"novachat_engine/service/core/update"
)

type UpdatesServiceImpl struct {
	conf               *conf.Config
	accountCore        *account.Core
	usersCore          *sql_core.Core
	accountUpdateCore  *accountUpdate.Core
	accountUsersCore   *accountUsers.Core
	accountContactCore *accountContact.Core
	phoneCallCore      *phone_call.PhoneCallCore
	contactCore        *contact.Core
	accountMessageCore *accountMessage.Core
	messageCore        *message.Core
	accountSettingCore *accountSetting.Core

	//*phone_call.PhoneCallConfig
}

func NewUpdatesServiceImpl(conf *conf.Config) *UpdatesServiceImpl {
	impl := &UpdatesServiceImpl{
		conf:          conf,
		accountCore:   account.NewAccountCore(cache.GetRedisClient()),
		usersCore:     sql_core.NewUsersCore(&conf.MySQL),
		contactCore:   contact.NewContactCore(&conf.MongoDB),
		phoneCallCore: phone_call.NewPhoneCallCore(conf.RelayConfig.PhoneCallWebRTC, conf.RelayConfig.P2P, conf.RelayConfig.LibVersion, conf.Server.ServerId, cache.GetRedisClient()),
		messageCore:   message.NewMessageCore(&conf.MongoDB, nil),
		//PhoneCallConfig: phone_call.NewPhoneCallConfig(conf.RelayConfig.PhoneCallWebRTC),
	}

	impl.accountSettingCore = accountSetting.NewSettingCore(setting.NewSettingCore(&conf.MongoDB))
	impl.accountContactCore = accountContact.NewContactCore(impl.contactCore)
	impl.accountUsersCore = accountUsers.NewUsersCore(impl.accountCore, impl.accountContactCore, impl.usersCore)
	impl.accountUpdateCore = accountUpdate.NewCore(update.NewCore(&conf.MongoDB), impl.accountUsersCore)
	impl.accountMessageCore = accountMessage.NewMessageCore(nil,
		impl.usersCore,
		impl.accountContactCore,
		nil,
		nil,
		nil,
		impl.messageCore,
		nil)
	return impl
}
