/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/11/8 13:48
 * @File : sms_module.go
 */

package sms_module

import (
	"github.com/BurntSushi/toml"
	"novachat_engine/service/sms"
	"novachat_engine/service/sms/config"
	"novachat_engine/service/sms/sms_service"
)

type testSmsService struct {
}

func (*testSmsService) Send(phoneNumber, Code string) error {
	return nil
}

func (*testSmsService) PlatForm() config.SmsPlatform {
	return config.SmsPlatformNone
}

func InstallSms(confPlatform config.SmsPlatform, fileName string) {
	var _smsService sms.Service
	if confPlatform == config.SmsPlatformTest {
		c := config.SubmailConfig{}
		_, err := toml.DecodeFile(fileName, &c)
		if err != nil {
			panic(err)
		}
		_smsService = &testSmsService{}
	}

	sms.Register(sms_service.NewSmsService(_smsService))
}
