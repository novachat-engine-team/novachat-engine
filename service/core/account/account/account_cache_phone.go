/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/16 17:15
 * @File :
 */

package account

func (m *Core) PushPhoneNumber(phoneNumber string) (bool, error) {
	return m.push(AccountCachePhoneNumberPrefix, phoneNumber)
}

func (m *Core) PushPhoneNumberList(phoneNumberList []string) ([]bool, error) {
	return m.pushList(AccountCachePhoneNumberPrefix, phoneNumberList)
}

func (m *Core) PhoneNumberExists(phoneNumber string) (bool, error) {
	return m.exists(AccountCachePhoneNumberPrefix, phoneNumber)
}

func (m *Core) PhoneNumberListExists(phoneNumberList []string) ([]bool, error) {
	return m.listExists(AccountCachePhoneNumberPrefix, phoneNumberList)
}
