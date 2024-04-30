/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File :
 * @Desc :
 *
 */

package privacy

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"novachat_engine/service/data/privacy"
	"novachat_engine/service/mgo"
)

const (
	DBPrivacy    string = "db_privacy"
	TablePrivacy string = "tb_privacy"
)

func init() {
	mgo.DBE.Enable64ToString = false
}

type Core struct{}

func (c Core) GetPrivacy(userId int64, key int32) ([]*data_privacy.PrivacyOption, error) {

	log.Debugf("GetPrivacy userId:%d privacyKey:%v", userId, key)
	op := options.Collection()
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetDatabase(DBPrivacy).
		Collection(TablePrivacy, op)

	name := FillPrivacy(&data_privacy.Privacy{}, key, nil)
	of := options.FindOne()
	of.SetProjection(bson.M{name: 1})

	result := col.FindOne(context.Background(), bson.M{"user_id": userId, "privacy_key": key}, of)
	if result.Err() != nil && (result.Err() != mongo.ErrNoDocuments || result.Err() != mongo.ErrNilDocument) {
		log.Errorf("GetPrivacy userId:%d privacyKey:%v error:%s", userId, key, result.Err().Error())
		return nil, result.Err()
	}
	if result.Err() == mongo.ErrNoDocuments || result.Err() == mongo.ErrNilDocument {
		return []*data_privacy.PrivacyOption{}, nil
	}
	var ret *data_privacy.Privacy
	err := result.Decode(&ret)
	if err != nil {
		log.Errorf("GetPrivacy Decode userId:%d privacyKey:%v error:%s", userId, key, err.Error())
	}
	return TokenPrivacy(ret, key), nil
}

func (c Core) GetPrivacyKeys(userId int64, keys []int32) (*data_privacy.Privacy, error) {

	log.Debugf("GetPrivacyKeys userId:%d privacyKey:%v", userId, keys)
	op := options.Collection()
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetDatabase(DBPrivacy).
		Collection(TablePrivacy, op)

	of := options.FindOne()
	projection := bson.M{}
	for _, v := range keys {
		projection[FillPrivacy(&data_privacy.Privacy{}, v, nil)] = 1
	}
	of.SetProjection(projection)

	result := col.FindOne(context.Background(), bson.M{"user_id": userId, "privacy_key": bson.M{mgo.IN: keys}}, of)
	if result.Err() != nil && (result.Err() != mongo.ErrNoDocuments || result.Err() != mongo.ErrNilDocument) {
		log.Errorf("GetPrivacyKeys userId:%d privacyKey:%v error:%s", userId, keys, result.Err().Error())
		return nil, result.Err()
	}
	if result.Err() == mongo.ErrNoDocuments || result.Err() == mongo.ErrNilDocument {
		return &data_privacy.Privacy{}, nil
	}
	var ret *data_privacy.Privacy
	err := result.Decode(&ret)
	if err != nil {
		log.Errorf("GetPrivacyKeys Decode userId:%d privacyKey:%v error:%s", userId, keys, err.Error())
	}
	return ret, nil
}

func (c Core) UpdatePrivacy(userId int64, key int32, privacy []*data_privacy.PrivacyOption) (bool, error) {
	log.Debugf("UpdatePrivacy userId:%d privacyKey:%d", userId, key)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	//op.SetRegistry(mgo.Registry())
	updateOptions := options.Update()
	updateOptions.SetUpsert(true)

	col := mgo.GetDatabase(DBPrivacy).
		Collection(TablePrivacy, op)

	data := &data_privacy.Privacy{}
	memberName := FillPrivacy(data, key, privacy)

	updateResult, err := col.UpdateOne(
		context.Background(),
		bson.M{"user_id": userId},
		bson.M{mgo.SET: mgo.DBE.MarshalCustomSpecMap(data, memberName)}, updateOptions)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("UpdatePrivacy userId:%d privacyKey:%d values:%+v error:%s", userId, key, privacy, err.Error())
		return false, err
	}

	log.Infof("UpdatePrivacy userId:%d privacyKey:%d values:%+v updateResult:%+v", userId, key, privacy, updateResult)
	return true, nil
}

func NewPrivacyCore(conf *config.MongodbConfig) *Core {
	mgo.InstallMongodb(conf)

	return &Core{}
}
