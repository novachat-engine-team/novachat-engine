/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package update

import (
	accountUsers "novachat_engine/service/account/users"
	"novachat_engine/service/core/update"
)

type Core struct {
	accountUsersCore *accountUsers.Core
	updateCore       *update.Core
}

func NewCore(updateCore *update.Core, accountUsersCore *accountUsers.Core) *Core {
	c := &Core{
		accountUsersCore: accountUsersCore,
		updateCore:       updateCore,
	}
	return c
}
