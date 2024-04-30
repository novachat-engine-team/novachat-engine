/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/21 17:12
 * @File : banned.go
 */

package banned_db

import (
	"novachat_engine/service/core/account/banned_phone_number"
)

const (
	BannedNameDB = "bannedNameDB"
)

func init() {
	banned_phone_number.Register(BannedNameDB, &bannedNameDB{})
}

type bannedNameDB struct {
}

func (m *bannedNameDB) Initialize() error {
	return nil
}

func (m *bannedNameDB) BannedNumber(number string) error {
	return nil
}

func (m *bannedNameDB) InvalidNumber(number string) bool {
	return false
}
