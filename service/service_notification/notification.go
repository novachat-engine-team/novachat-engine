/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/12/17 17:44
 * @File : notification.go
 */

package service_notification

import (
	"novachat_engine/mtproto"
	"time"
)

type ServiceNotificationType string

const (
	ServiceNotificationTypeAuthKeyDrop ServiceNotificationType = "AUTH_KEY_DROP_"
	ServiceNotificationTypeRecharge    ServiceNotificationType = "RECHARGE"
)

func StringToServiceNotificationType(m string) ServiceNotificationType {
	return ServiceNotificationType(m)
}

func (m ServiceNotificationType) ToString() string {
	return string(m)
}

func Notification(popup bool, message string, media *mtproto.MessageMedia, notifyType ServiceNotificationType, messageEntity []*mtproto.MessageEntity) *mtproto.Update {
	if media == nil {
		media = mtproto.NewTLMessageMediaEmpty(nil).To_MessageMedia()
	}
	return mtproto.NewTLUpdateServiceNotification(&mtproto.Update{
		PopupEBE4681971:   popup,
		InboxDate:         int32(time.Now().Unix()),
		Type:              notifyType.ToString(),
		Media:             media,
		Entities:          messageEntity,
		MessageEBE4681971: message,
	}).To_Update()
}

type ServiceNotification struct {
	PopUp         bool                     `json:"p"`
	NotifyType    ServiceNotificationType  `json:"t"`
	Media         *mtproto.MessageMedia    `json:"d"`
	MessageEntity []*mtproto.MessageEntity `json:"e"`
	Message       string                   `json:"m"`
	ValidTime     int32                    `json:"v"`
	OrderId       int32                    `json:"o"`
}
