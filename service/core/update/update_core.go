/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package update

import (
	"novachat_engine/pkg/config"
	"novachat_engine/service/mgo"
)

const (
	DBUpdate    = "db_updates"
	TableUpdate = "tb_updates"
)

func init() {
}

type Core struct {
	encoder *mgo.Encoder
}

func NewCore(config *config.MongodbConfig) *Core {
	mgo.InstallMongodb(config)

	c := &Core{
		encoder: mgo.NewEncoder(mgo.WithEncodeTag(mgo.BsonTag), mgo.WithEmitDefaults(false), mgo.WithEnable64ToString(false)),
	}
	return c
}
