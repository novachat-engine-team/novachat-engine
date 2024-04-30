/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/21 16:55
 * @File : banned.go
 */

package banned_phone_number

import (
	"fmt"
)

type BannedPhoneNumber interface {
	Initialize() error
	BannedNumber(number string) error
	InvalidNumber(number string) bool
}

var _instance *BannedPhoneNumber
var _register = make(map[string]BannedPhoneNumber)

func Register(bannedName string, ins BannedPhoneNumber) {
	_register[bannedName] = ins
}

func InstallBannedPhoneNumber(bannedName string) (BannedPhoneNumber, error) {
	ins, ok := _register[bannedName]
	if !ok {
		return nil, fmt.Errorf("InstallBannedPhoneNumber bannedName:%s not register", bannedName)
	}
	return ins, nil
}


