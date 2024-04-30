/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/16 17:15
 * @File :
 */

package account

func (m *Core) PushUserName(userName string) (bool, error) {
	return m.push(AccountCacheUserNamePrefix, userName)
}

func (m *Core) PushUserNameList(userNameList []string) ([]bool, error) {
	return m.pushList(AccountCacheUserNamePrefix, userNameList)
}

func (m *Core) UserNameExists(userName string) (bool, error) {
	return m.exists(AccountCacheUserNamePrefix, userName)
}

func (m *Core) UserNameListExists(userNameList []string) ([]bool, error) {
	return m.listExists(AccountCacheUserNamePrefix, userNameList)
}
