/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package message_id_pts

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"novachat_engine/service/core/id"
	data_id "novachat_engine/service/data/messages/id/message_id_pts"
	"novachat_engine/service/mgo"
)

type Core struct{}

type Option func(id *data_id.Id)

func init() {
	mgo.DBE.Enable64ToString = false
}

func WithNextMessageId(next int32) Option {
	return func(id *data_id.Id) {
		id.MessageId = next
	}
}

func WithNextPts(next int32) Option {
	return func(id *data_id.Id) {
		id.Pts = next
	}
}

func NewMessageIdPtsCore(conf *config.MongodbConfig) *Core {
	mgo.InstallMongodb(conf)
	return &Core{}
}

func (c *Core) NextId(userId int64, opts ...Option) (*data_id.Id, error) {
	if len(opts) == 0 {
		panic("NextId opt empty")
	}

	dataId := &data_id.Id{
		UserId: userId,
	}
	for _, opt := range opts {
		opt(dataId)
	}

	var err error
	op := options.Collection()
	ofu := options.FindOneAndUpdate()
	ofu.SetUpsert(true)
	//op.SetRegistry(mgo.Registry())

	sigResult := mgo.GetDatabase(id.DBID).Collection(id.TableID, op).
		FindOneAndUpdate(context.Background(), mgo.DBE.MarshalCustomSpecMap(dataId, "UserId"),
			bson.M{mgo.INC: mgo.DBE.MarshalCustomSpecMap(dataId, "MessageId", "Pts")},
			ofu)
	if sigResult.Err() != nil && sigResult.Err() != mongo.ErrNoDocuments {
		log.Fatalf("NextId userId:%v error:%s", userId, err.Error())
		return nil, err
	}

	var r = &data_id.Id{
		UserId: userId,
	}
	if sigResult.Err() == mongo.ErrNoDocuments {
		for _, opt := range opts {
			opt(r)
		}
	} else {
		err = sigResult.Decode(r)
		if err != nil {
			log.Errorf("NextId userId:%v decode error:%s", userId, err.Error())
			return nil, err
		}
	}

	return r, nil
}
