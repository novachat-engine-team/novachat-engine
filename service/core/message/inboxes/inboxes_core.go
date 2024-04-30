/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package inboxes

import (
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/snowflake"
	"novachat_engine/service/core/message/message"
	"novachat_engine/service/mgo"
)

func init() {
	mgo.DBE.Enable64ToString = false
}

type Core struct {
	messageCore *message.Core
}

func NewInboxesCore(config *config.MongodbConfig, messageIdNode *snowflake.Node) *Core {
	mgo.InstallMongodb(config)

	return &Core{
		messageCore: message.NewMessageCore(config, messageIdNode),
	}
}
