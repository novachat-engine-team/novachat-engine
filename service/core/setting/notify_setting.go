/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File :
 * @Desc :
 *
 */

package setting

import (
	"fmt"
	"novachat_engine/pkg/config"
	"novachat_engine/service/constants"
	"novachat_engine/service/mgo"
)

const (
	DBSetting              string = "db_settings"
	TableNameNotifySetting string = "tb_notify_settings"
)

func init() {
	mgo.DBE.Enable64ToString = false
}

func makeNotifySettingId(userId int64, peerId int64, peerNotifyType int32) string {
	if peerNotifyType == constants.PeerGlobalSetting.ToInt32() {
		return fmt.Sprintf("%d%d", 0, userId)
	}
	return fmt.Sprintf("%d%d", peerId, userId)
}

type Core struct {
}

func NewSettingCore(conf *config.MongodbConfig) *Core {
	mgo.InstallMongodb(conf)
	return &Core{}
}
