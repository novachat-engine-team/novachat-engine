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

func (c *Core) GetUpdateChannelDifference(userId, chat int64, pts, limit int32) ([]*data_update.UserUpdate, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	findOpt := options.Find()
	findOpt.SetLimit(int64(limit))

	filter := bson.M{"peer_id": chat, "pts": bson.M{mgo.GT: pts}, "owner_user_id": userId}

	col := mgo.GetDatabase(DBUpdate).
		Collection(message.TableName(chat, TableChannelUpdate), op)
	cursor, err := col.Find(context.Background(), filter, findOpt)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("GetUpdateChannelDifference error:%s", err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNilDocument {
		log.Warnf("GetUpdateChannelDifference error:%s", err.Error())
		return nil, nil
	}

	var updateList []*data_update.UserUpdate
	err = cursor.All(context.TODO(), &updateList)
	if err != nil {
		log.Warnf("GetUpdateChannelDifference decode error:%s", err.Error())
	}

	log.Debugf("GetUpdateChannelDifference chatId:%d, pts:%d,  limit:%d len:%d", chat, pts, limit, len(updateList))
	return updateList, nil
}
