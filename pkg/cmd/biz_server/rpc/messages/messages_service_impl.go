/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages_service_impl.go
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
	accountStickerSet "novachat_engine/service/account/stickerset"
	accountUsers "novachat_engine/service/account/users"
	"novachat_engine/service/app/help"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/contact"
	"novachat_engine/service/core/account/privacy"
	"novachat_engine/service/core/account/users/sql_core"
	"novachat_engine/service/core/message/message"
	"novachat_engine/service/core/setting"
	"novachat_engine/service/core/stickerset"
	"novachat_engine/service/upload"
)

type MessagesServiceImpl struct {
	conf                  *conf.Config
	accountCore           *account.Core
	usersCore             *sql_core.Core
	accountMessageCore    *accountMessage.Core
	accountContactCore    *accountContact.Core
	accountUsersCore      *accountUsers.Core
	accountStickerSetCore *accountStickerSet.Core
	contactCore           *contact.Core
	privacyCore           *privacy.Core
	uploadCore            *upload.Core
	messageCore           *message.Core
	accountSettingCore    *accountSetting.Core
}

func NewMessagesServiceImpl(conf *conf.Config) *MessagesServiceImpl {
	impl := &MessagesServiceImpl{
		conf:        conf,
		accountCore: account.NewAccountCore(cache.GetRedisClient()),
		usersCore:   sql_core.NewUsersCore(&conf.MySQL),
		contactCore: contact.NewContactCore(&conf.MongoDB),
		privacyCore: privacy.NewPrivacyCore(&conf.MongoDB),
		uploadCore:  upload.NewUploadCore(&conf.SfsClient),
		messageCore: message.NewMessageCore(&conf.MongoDB, nil),
		accountStickerSetCore: accountStickerSet.NewStickerSetCore(
			stickerset.NewStickerSetCore(&conf.MongoDB),
			conf.Redis,
			help.GetConfig().StickersFavedLimit,
			help.GetConfig().StickersRecentLimit),
	}

	impl.accountSettingCore = accountSetting.NewSettingCore(setting.NewSettingCore(&conf.MongoDB))
	impl.accountContactCore = accountContact.NewContactCore(impl.contactCore)
	impl.accountUsersCore = accountUsers.NewUsersCore(impl.accountCore, impl.accountContactCore, impl.usersCore)
	impl.accountMessageCore = accountMessage.NewMessageCore(&conf.Redis,
		impl.usersCore,
		impl.accountContactCore,
		impl.privacyCore,
		impl.accountCore,
		impl.uploadCore,
		impl.messageCore,
		impl.accountSettingCore)
	return impl
}
