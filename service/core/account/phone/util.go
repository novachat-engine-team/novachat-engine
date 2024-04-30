/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/25 17:26
 * @File : util.go
 */

package phone

import (
	"fmt"
	"strconv"
)

const defaultNiceNumber = 10000
const virtualPhoneNumberBase = 2000000000000000

var virtualPhoneNumberLen = 16 // default len

var ErrPhoneNumberCache = fmt.Errorf("error InstallRedis first to InstallPhoneNumber")

func init() {
	s := strconv.FormatInt(virtualPhoneNumberBase, 10)
	virtualPhoneNumberLen = len(s)
}

//var _virtual *cache1.VirtualPhoneNumber
//var _nice *cache1.NicePhoneNumber
//
//func InstallPhoneNumber() error {
//	r := cache.GetRedisClient()
//	if r == nil {
//		return ErrVirtualPhoneNumberFormat
//	}
//	_virtual = cache1.NewVirtualPhoneNumber(r)
//	_nice = cache1.NewNicePhoneNumber(r)
//
//	return nil
//}
//
//func GenerateNiceNumber() (string, error) {
//	return _nice.GetNicePhoneNumber(strconv.Itoa(defaultNiceNumber))
//}
//
//func GenerateVirtualNumber() (string, error) {
//	return _virtual.GetVirtualPhoneNumber(strconv.Itoa(virtualPhoneNumberBase))
//}
