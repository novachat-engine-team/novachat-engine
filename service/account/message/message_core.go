/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package message

import (
	"fmt"
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/config"
	"novachat_engine/service/account/contact"
	"novachat_engine/service/account/setting"
	"novachat_engine/service/account/users"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/privacy"
	"novachat_engine/service/core/account/users/sql_core"
	"novachat_engine/service/core/message/message"
	"novachat_engine/service/upload"
)

func makeMutexDeleteMessageName(userId int64, hash int32) string {
	return fmt.Sprintf("string:deletemessage:mutex:%d-%d", userId, hash)
}
func makeMutexMessageName(userId int64, randomId int64) string {
	return fmt.Sprintf("string:message:mutex:%d-%d", userId, randomId)
}

func makeMessageKey(userId int64, randomId int64) string {
	return fmt.Sprintf("string:message:%d-%d", userId, randomId)
}

type Core struct {
	userCore           *sql_core.Core
	accountCore        *account.Core
	accountContactCore *contact.Core
	privacyCore        *privacy.Core
	messageCore        *message.Core
	uploadCore         *upload.Core
	accountUsersCore   *users.Core
	settingCore        *setting.Core
}

func NewMessageCore(client *config.RedisConfig,
	userCore *sql_core.Core,
	contact *contact.Core,
	privacy *privacy.Core,
	accountCore *account.Core,
	uploadCore *upload.Core,
	messageCore *message.Core,
	settingCore *setting.Core) *Core {
	if client != nil {
		cache.InstallRedSync(*client)
	}

	return &Core{
		accountContactCore: contact,
		accountUsersCore:   users.NewUsersCore(accountCore, contact, userCore),
		accountCore:        accountCore,
		userCore:           userCore,
		privacyCore:        privacy,
		messageCore:        messageCore,
		uploadCore:         uploadCore,
		settingCore:        settingCore,
	}
}
