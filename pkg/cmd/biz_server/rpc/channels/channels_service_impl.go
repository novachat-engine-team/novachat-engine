/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels_service_impl.go
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
	accountUsers "novachat_engine/service/account/users"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/contact"
	"novachat_engine/service/core/account/users/sql_core"
	"novachat_engine/service/core/message/message"
	"novachat_engine/service/core/setting"
)

type ChannelsServiceImpl struct {
	conf               *conf.Config
	accountCore        *account.Core
	usersCore          *sql_core.Core
	contactCore        *contact.Core
	accountContactCore *accountContact.Core
	accountUsersCore   *accountUsers.Core
	accountSettingCore *accountSetting.Core
	messageCore        *message.Core
	accountMessageCore *accountMessage.Core
}

func NewChannelsServiceImpl(conf *conf.Config) *ChannelsServiceImpl {
	impl := &ChannelsServiceImpl{
		conf:        conf,
		accountCore: account.NewAccountCore(cache.GetRedisClient()),
		usersCore:   sql_core.NewUsersCore(&conf.MySQL),
		contactCore: contact.NewContactCore(&conf.MongoDB),
		messageCore: message.NewMessageCore(&conf.MongoDB, nil),
	}
	impl.accountContactCore = accountContact.NewContactCore(impl.contactCore)
	impl.accountUsersCore = accountUsers.NewUsersCore(impl.accountCore, impl.accountContactCore, impl.usersCore)
	impl.accountSettingCore = accountSetting.NewSettingCore(setting.NewSettingCore(&conf.MongoDB))
	impl.accountMessageCore = accountMessage.NewMessageCore(nil,
		impl.usersCore,
		impl.accountContactCore,
		nil,
		impl.accountCore,
		nil,
		impl.messageCore,
		nil)
	return impl
}
