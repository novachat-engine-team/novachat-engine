/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package message

import (
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/snowflake"
	"novachat_engine/service/core/id/message_id_pts"
	"novachat_engine/service/mgo"
)

func init() {
	mgo.DBE.Enable64ToString = false
}

type Core struct {
	id            *message_id_pts.Core
	messageIdNode *snowflake.Node
}

func NewMessageCore(config *config.MongodbConfig, messageIdNode *snowflake.Node) *Core {
	mgo.InstallMongodb(config)

	return &Core{
		id:            message_id_pts.NewMessageIdPtsCore(config),
		messageIdNode: messageIdNode,
	}
}
