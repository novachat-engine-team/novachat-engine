/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/9/22 15:34
 * @File : aes_test.go
 */

package util

import (
	"fmt"
	"testing"
)

var PwdKey = "3lU7WkUi5lJaKwKR1y5uDcnFfqNygVlh"

func EnPwdCode(pwd string) (string, error) {
	result, err := AesEncrypt(pwd, PwdKey)
	if err != nil {
		return "", err
	}
	return result, err
}

func DePwdCode(pwd string) (string, error) {
	return AesDecrypt(pwd, PwdKey)
}

func TestAesEncrypt(t *testing.T) {
	fmt.Println(len(PwdKey))
	str := "13000000000"
	pwd, _ := EnPwdCode(str)
	bytes, _ := DePwdCode(pwd)
	fmt.Println(len(pwd))
	fmt.Println(bytes)
}
