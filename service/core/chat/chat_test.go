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
	"novachat_engine/pkg/config"
	"novachat_engine/service/mgo"
	"strings"
	"testing"
)

func TestAAAAA(t *testing.T) {
	mgo.InstallMongodb(&config.MongodbConfig{
		Direct:      true,
		Addr:        []string{"192.168.10.109:27017"},
		ReplicaSet:  "rs0",
		Database:    "test",
		Username:    "",
		Password:    "",
		MaxPoolSize: 0,
		MinPoolSize: 0,
	})

	var err error
	err = mgo.GetMongoDB().Database("test").
		CreateCollection(context.Background(), "test")
	if err != nil {
		if strings.Index(err.Error(), "Collection already exists") >= 0 {
			col := mgo.GetMongoDB().Database("test").Collection("test")
			col.Find(context.TODO(), bson.M{})
		} else {
			t.Fatal(err.Error())
		}
	}

	type Test struct {
		Id    string `bson:"_id" json:"_id"`
		Value int64  `bson:"value" json:"value"`
	}
	col := mgo.GetMongoDB().Database("test").Collection("test")
	rs := col.FindOne(context.TODO(), bson.M{"_id": "config"})
	v := &Test{}
	rs.Decode(v)
	t.Logf("%+v", v)

	_, err = col.UpdateOne(context.TODO(), bson.M{"_id": "config"}, bson.M{mgo.INC: bson.M{"value": 1}})
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	rs = col.FindOne(context.TODO(), bson.M{"_id": "config"})
	rs.Decode(v)

	t.Logf("%+v", v)
}
