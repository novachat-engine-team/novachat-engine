/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/18 18:51
 * @File : mongo_install.go
 */

package mgo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"novachat_engine/pkg/config"
	"sync"
)

const (
	LT          string = "$lt"  // 小于
	LTE         string = "$lte" // 小于或等于
	GT          string = "$gt"  // 大于
	GTE         string = "$gte" // 大于或等于
	NE          string = "$ne"  // 不等于
	IN          string = "$in"  // 属于
	OR          string = "$or"
	MATCH       string = "$match"
	NOTIN       string = "$nin"
	LOOKUP      string = "$lookup"
	SET         string = "$set"
	ADDTOSET    string = "$addToSet"
	PUSH        string = "$push"
	POP         string = "$pop"  // $pop: {score: -1} 第一个元素  $pop: {score: 1} 最有一个元素
	PULL        string = "$pull" //$pull: { ListOfNames: { $in: [ "David"] }}
	INC         string = "$inc"
	SETONINSERT string = "$setOnInsert"
)

var once sync.Once
var _mgoClient *MongoDB

func InstallMongodb(conf *config.MongodbConfig) {
	once.Do(func() {
		_mgoClient = NewMongoDB(conf)
	})
}

func GetMongoDB() *MongoDB {
	return _mgoClient
}

func GetDatabase(dbName string) *mongo.Database {
	return GetMongoDB().Database(dbName)
}

func GetDatabasePrimaryPreferred(dbName string) *mongo.Database {
	do := options.Database().SetReadPreference(readpref.PrimaryPreferred())
	return GetMongoDB().Database(dbName, do)
}

func GetDatabaseSecondaryPreferred(dbName string) *mongo.Database {
	do := options.Database().SetReadPreference(readpref.SecondaryPreferred())
	return GetMongoDB().Database(dbName, do)
}

func GetDatabaseSecondary(dbName string) *mongo.Database {
	do := options.Database().SetReadPreference(readpref.Secondary())
	return GetMongoDB().Database(dbName, do)
}
