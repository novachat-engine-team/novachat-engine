/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package auth

import (
	"fmt"
	dataAuth "novachat_engine/service/data/auth"
)

const (
	KeyAuthKeyId = "key:authKeyId:"
	KeyUserId    = "key:userId:"
)

func makeAuthKeyIdKey(authKeyId int64) string {
	return fmt.Sprintf("%s%d", KeyAuthKeyId, authKeyId)
}

func makeUserIdKey(userId int64) string {
	return fmt.Sprintf("%s%d", KeyUserId, userId)
}

type SessionInfo struct {
	AuthInfo      *dataAuth.Auth `json:"info"`
	SessionIdList []int64        `json:"session_id_list"`
}

type UsersInfo struct {
	AuthKeyIdList []int64 `json:"auth_key_id_list"`
}
