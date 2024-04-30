/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/5/6 17:39
 * @File : service_test.go
 */

package service_notification

import (
	"novachat_engine/mtproto"
	"novachat_engine/pkg/config"
	"novachat_engine/service/message/messages"
	"testing"
)

func TestService(t *testing.T) {
	InstallServerNotification(config.EtcdServerConfig{
		Schema:               "novachat_engine",
		ServerName:           "notification",
		EtcdAddr:             "192.168.1.80:2379",
		TTL:                  0,
		DialTimeout:          0,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		UserName:             "",
		Password:             "",
	}, true)

	msg := "Notification Value https://www.baidu.com"

	messageEnties := make([]*mtproto.MessageEntity, 0, 1)

	posList, ok := messages.ParseUrl(msg)
	if ok == true {
	   for _, v := range posList {
	       entityUrl := mtproto.NewTLMessageEntityTextUrl(nil)
	       entityUrl.SetOffset(v.Beg)
	       entityUrl.SetLength(v.End)
	       entityUrl.SetUrl(v.Text)
	       messageEnties = append(messageEnties, entityUrl.To_MessageEntity())
	   }
	}

	t.Logf("%v", SetNotification([]ServiceNotification{
		{
			PopUp:         false,
			NotifyType:    "",
			Media:         nil,
			MessageEntity: messageEnties,
			Message:       msg,
			ValidTime:     0,
		},
	}))
}
