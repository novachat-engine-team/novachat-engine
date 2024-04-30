/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts_service_impl.go
 * @Desc :
 *
 */

package rpc

import (
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/cmd/biz_server/conf"
	accountContact "novachat_engine/service/account/contact"
	accountUsers "novachat_engine/service/account/users"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/contact"
	"novachat_engine/service/core/account/users/sql_core"
)

type ContactsServiceImpl struct {
	conf               *conf.Config
	userCore           *sql_core.Core
	accountCore        *account.Core
	contactCore        *contact.Core
	accountContactCore *accountContact.Core
	accountUserCore    *accountUsers.Core
}

func NewContactsServiceImpl(conf *conf.Config) *ContactsServiceImpl {
	impl := &ContactsServiceImpl{
		conf:        conf,
		userCore:    sql_core.NewUsersCore(&conf.MySQL),
		accountCore: account.NewAccountCore(cache.GetRedisClient()),
		contactCore: contact.NewContactCore(&conf.MongoDB),
	}
	impl.accountContactCore = accountContact.NewContactCore(impl.contactCore)
	impl.accountUserCore = accountUsers.NewUsersCore(impl.accountCore, impl.accountContactCore, impl.userCore)

	return impl
}
