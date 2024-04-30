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
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/log"
	data_chat "novachat_engine/service/data/chat"
	"novachat_engine/service/mgo"
)

func (c *Core) GetChatList(opts ...GetChatListOptions) ([]*data_chat.Chat, error) {
	opt := getChatListOptions{}
	for _, v := range opts {
		v(&opt)
	}

	filter := bson.M{}
	if len(opt.Ids) == 0 && len(opt.ExceptIds) == 0 {
		filter["_id"] = bson.M{mgo.GT: 0}
	} else {
		if len(opt.Ids) > 0 {
			filter["_id"] = bson.M{mgo.IN: opt.Ids}
		} else {
			filter["_id"] = bson.M{mgo.NOTIN: opt.ExceptIds}

		}
	}

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	cursor, err := mgo.GetDatabase(DBChats).
		Collection(TableChats, op).
		Find(context.TODO(), filter)
	if err != nil {
		log.Fatalf("GetChatList error:%s", err.Error())
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var chatsList []*data_chat.Chat
	err = cursor.All(context.TODO(), &chatsList)
	if err != nil {
		log.Fatalf("GetChatList decode error:%s", err.Error())
		return nil, err
	}
	return chatsList, nil
}
