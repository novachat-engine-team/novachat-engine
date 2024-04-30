/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/16 17:15
 * @File :
 */

package account

import "fmt"

func MakeAccountCacheUsersPrefix(userId int64) string {
	return fmt.Sprintf("%s%d", AccountCacheUsersUserIdPrefix, userId)
}

func MakeAccountCacheUsersPhoneNumberPrefix(phoneNumber string) string {
	return fmt.Sprintf("%s%s", AccountCacheUsersPhoneNumberPrefix, phoneNumber)
}

func MakeAccountCacheUsersUsernamePrefix(username string) string {
	return fmt.Sprintf("%s%s", AccountCacheUsersUsernamePrefix, username)
}
