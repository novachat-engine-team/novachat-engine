/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File :
 * @Desc :
 *
 */

package account

import (
	"github.com/json-iterator/go"
	"novachat_engine/service/data/users"
)

func CacheInfo2User(info string) (*data_users.Users, error) {
	var user = &data_users.Users{}
	err := jsoniter.Unmarshal([]byte(info), user)
	return user, err
}

func User2CacheInfo(user *data_users.Users) (string, error) {
	return jsoniter.MarshalToString(user)
}
