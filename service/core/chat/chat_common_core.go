/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package chat

import (
	"novachat_engine/pkg/config"
	"novachat_engine/service/mgo"
)

type CommonChatCore struct {
}

func (c CommonChatCore) GetChatCommon(userId int64) ([]int64, error) {
	panic("GetChatCommon")
}

func NewChatCommonCore(conf *config.MongodbConfig) *CommonChatCore {
	mgo.InstallMongodb(conf)
	return &CommonChatCore{}
}
