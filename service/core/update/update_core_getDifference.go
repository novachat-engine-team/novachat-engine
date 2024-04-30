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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/log"
	"novachat_engine/service/core/message"
	data_update "novachat_engine/service/data/update"
	"novachat_engine/service/mgo"
)

func (c *Core) GetUpdateDifference(userId int64, pts int32, qts int32, date int32, limit int32, sync bool) ([]*data_update.UserUpdate, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	findOpt := options.Find()
	findOpt.SetLimit(int64(limit))

	filter := bson.M{"owner_user_id": userId, "date": bson.M{mgo.GTE: date}, "pts": bson.M{mgo.GTE: pts}}
	if sync {
		findOpt.Sort = bson.M{"pts": -1}
	}
	col := mgo.GetDatabase(DBUpdate).
		Collection(message.TableName(userId, TableUpdate), op)
	cursor, err := col.Find(context.Background(), filter, findOpt)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("GetUpdateDifference error:%s", err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNilDocument {
		log.Warnf("GetUpdateDifference error:%s", err.Error())
		return nil, nil
	}

	var updateList []*data_update.UserUpdate
	err = cursor.All(context.TODO(), &updateList)
	if err != nil {
		log.Warnf("GetUpdateDifference decode error:%s", err.Error())
	}

	log.Debugf("GetUpdateDifference userId:%d, pts:%d, qts:%d, date:%d, limit:%d", userId, pts, qts, date, limit, len(updateList))
	return updateList, nil
}
