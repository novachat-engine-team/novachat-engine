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
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/service/core/message"
	data_update "novachat_engine/service/data/update"
	"novachat_engine/service/mgo"
)

func (c *Core) SaveUpdateChannelDataList(list []*data_update.UserUpdate) error {

	peerId := list[0].PeerId
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetDatabase(DBUpdate).
		Collection(message.TableName(peerId, TableChannelUpdate), op)

	var imr *mongo.InsertManyResult
	var ior *mongo.InsertOneResult
	var err error
	if len(list) == 1 {
		ior, err = col.InsertOne(context.TODO(), c.encoder.MarshalCustomSpecMap(list[0]))
	} else {
		documents := make([]interface{}, 0, len(list))
		for _, v := range list {
			documents = append(documents, c.encoder.MarshalCustomSpecMap(v))
		}
		imr, err = col.InsertMany(context.TODO(), documents)
	}

	if err != nil && !mongo.IsDuplicateKeyError(err) {
		return err
	}

	_ = imr
	_ = ior
	return nil
}
