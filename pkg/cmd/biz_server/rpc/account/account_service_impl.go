/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account_service_impl.go
 * @Desc :
 *
 */

package rpc

import (
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/cmd/biz_server/conf"
	accountContact "novachat_engine/service/account/contact"
	accountPrivacy "novachat_engine/service/account/privacy"
	accountSetting "novachat_engine/service/account/setting"
	accountUsers "novachat_engine/service/account/users"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/contact"
	"novachat_engine/service/core/account/privacy"
	"novachat_engine/service/core/account/users/sql_core"
	"novachat_engine/service/core/setting"
)

type AccountServiceImpl struct {
	conf *conf.Config

	accountUsersCore   *accountUsers.Core
	accountContactCore *accountContact.Core
	accountSettingCore *accountSetting.Core

	usersCore   *sql_core.Core
	privacyCore *accountPrivacy.Core
	accountCore *account.Core
}

func NewAccountServiceImpl(conf *conf.Config) *AccountServiceImpl {
	impl := &AccountServiceImpl{
		conf: conf,
	}

	impl.usersCore = sql_core.NewUsersCore(&conf.MySQL)
	impl.accountCore = account.NewAccountCore(cache.GetRedisClient())
	impl.privacyCore = accountPrivacy.NewPrivacyCore(privacy.NewPrivacyCore(&conf.MongoDB))

	impl.accountContactCore = accountContact.NewContactCore(contact.NewContactCore(&conf.MongoDB))
	impl.accountUsersCore = accountUsers.NewUsersCore(impl.accountCore, impl.accountContactCore, impl.usersCore)
	impl.accountSettingCore = accountSetting.NewSettingCore(setting.NewSettingCore(&conf.MongoDB))

	return impl
}
