/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/25 18:34
 * @File : sms.go
 */

package sms

import (
	"fmt"
	"go.uber.org/atomic"
	"novachat_engine/service/sms/config"
)

const TTL = 60 * 5
const CodeLength = 5
const DefaultPhoneCode = "12345"

const (
	CodeSignIn int64 = 0
	CodeVerify int64 = 1
	CodeChange int64 = 2
)

type Code struct {
	PhoneCode     string
	PhoneCodeHash string
}

type Sms interface {
	SendSmsVerifyCode(authKeyId int64, phoneNumber, phoneCode, phoneCodeHash string) error
	ReSendSmsVerifyCode(authKeyId int64, phoneNumber, phoneCodeHash string, newPhoneCode string) (int32, error)
	VerifySmsCode(authKeyId int64, phoneNumber, phoneCode, phoneCodeHash string) error
	ClearSmsCode(authKeyId int64, phoneNumber, phoneCode, phoneCodeHash string) error
	SetTestMode(bool)
}

type Service interface {
	Send(phoneNumber, Code string) error
	PlatForm() config.SmsPlatform
}

var _sms atomic.Value

func Register(s Sms) {
	ins := _sms.Load()
	if ins == nil {
		_sms.Store(s)
	} else {
		v, ok := ins.(Sms)
		if ok == false || v != s {
			_sms.Store(s)
		}
	}
}

func SetTestMode(testMode bool) {
	ins := _sms.Load()
	if ins != nil {
		v, ok := ins.(Sms)
		if ok == true {
			v.SetTestMode(testMode)
		}
	}
}

func SendSmsVerifyCode(authKeyId int64, phoneNumber, phoneCode, phoneCodeHash string) error {
	ins := _sms.Load()
	if ins != nil {
		v, ok := ins.(Sms)
		if ok == true {
			return v.SendSmsVerifyCode(authKeyId, phoneNumber, phoneCode, phoneCodeHash)
		}
	}
	return fmt.Errorf("sms need Register")
}

func ReSendSmsVerifyCode(authKeyId int64, phoneNumber, phoneCodeHash string, newPhoneCode string) (int32, error) {
	ins := _sms.Load()
	if ins != nil {
		v, ok := ins.(Sms)
		if ok == true {
			return v.ReSendSmsVerifyCode(authKeyId, phoneNumber, phoneCodeHash, newPhoneCode)
		}
	}
	return 0, fmt.Errorf("sms need Register")
}

func VerifySmsCode(authKeyId int64, phoneNumber, phoneCode, phoneCodeHash string) error {
	ins := _sms.Load()
	if ins != nil {
		v, ok := ins.(Sms)
		if ok == true {
			return v.VerifySmsCode(authKeyId, phoneNumber, phoneCode, phoneCodeHash)
		}
	}
	return fmt.Errorf("sms need Register")
}

func ClearSmsCode(authKeyId int64, phoneNumber, phoneCode, phoneCodeHash string) error {
	ins := _sms.Load()
	if ins != nil {
		v, ok := ins.(Sms)
		if ok == true {
			return v.ClearSmsCode(authKeyId, phoneNumber, phoneCode, phoneCodeHash)
		}
	}
	return fmt.Errorf("sms need Register")
}
