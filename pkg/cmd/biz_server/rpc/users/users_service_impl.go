/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : users_service_impl.go
 * @Desc :
 *
 */

package rpc

import (
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/cmd/biz_server/conf"
	accountBot "novachat_engine/service/account/bot"
	accountContact "novachat_engine/service/account/contact"
	accountPrivacy "novachat_engine/service/account/privacy"
	accountSetting "novachat_engine/service/account/setting"
	accountUsers "novachat_engine/service/account/users"
	"novachat_engine/service/app/phone_call"
	"novachat_engine/service/app/video_call"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/contact"
	"novachat_engine/service/core/account/privacy"
	"novachat_engine/service/core/account/users/sql_core"
	"novachat_engine/service/core/setting"
)

type UsersServiceImpl struct {
	conf               *conf.Config
	accountUserCore    *accountUsers.Core
	accountContactCore *accountContact.Core
	accountSettingCore *accountSetting.Core
	accountPrivacyCore *accountPrivacy.Core
	accountBotCore     *accountBot.Core
	phoneCallConfig    *phone_call.PhoneCallConfig
	videoCall          *video_call.VideoCall
	accountCore        *account.Core
	usersCore          *sql_core.Core
}

func NewUsersServiceImpl(conf *conf.Config) *UsersServiceImpl {
	impl := &UsersServiceImpl{
		conf:               conf,
		phoneCallConfig:    phone_call.NewPhoneCallConfig(conf.RelayConfig.PhoneCallWebRTC),
		videoCall:          video_call.NewVideoCall(),
		accountContactCore: accountContact.NewContactCore(contact.NewContactCore(&conf.MongoDB)),
		accountSettingCore: accountSetting.NewSettingCore(setting.NewSettingCore(&conf.MongoDB)),
	}
	impl.accountCore = account.NewAccountCore(cache.GetRedisClient())
	impl.usersCore = sql_core.NewUsersCore(&conf.MySQL)
	impl.accountBotCore = accountBot.NewBotCore()
	impl.accountPrivacyCore = accountPrivacy.NewPrivacyCore(privacy.NewPrivacyCore(&conf.MongoDB))
	impl.accountUserCore = accountUsers.NewUsersCore(impl.accountCore, impl.accountContactCore, impl.usersCore)
	return impl
}
