/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/11/8 14:24
 * @File : sms_test.go
 */

package sms_module

import (
	"github.com/BurntSushi/toml"
	"novachat_engine/service/sms/config"
	"novachat_engine/service/sms/sms_service_impl/yunpian"
	"testing"
)

func TestSubmail(t *testing.T) {
	//fileName := "E:\\test\\novachat_engine\\server\\biz_server\\sms_config_submail.toml"
	fileName := "E:\\test\\novachat_engine\\server\\biz_server\\sms_config_yunpian.toml"
	InstallSms(config.SmsPlatformYunpian, fileName)

//	sms.SendSmsVerifyCode(1, "8615017910674", "11111", "11111")

	c := config.YunPianConfig{}
	_, err := toml.DecodeFile(fileName, &c)
	if err != nil {
		panic(err)
	}
	_smsService := yunpian.NewYunPianService(c)
	err = _smsService.Send("8615017910674", "11111")
	if err != nil {

	}
	_ = err
}
