/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/23 18:02
 * @File :
 */

package auth

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	authService "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	dataAuth "novachat_engine/service/data/auth"
	"novachat_engine/service/mgo"
	"strings"
)

func init() {
	mgo.DBE.Enable64ToString = false
}

type Core struct {
	bsonEncoder *mgo.Encoder
}

func NewAuthCore(conf *config.MongodbConfig) *Core {
	mgo.InstallMongodb(conf)
	m := &Core{
		bsonEncoder: mgo.NewEncoder(mgo.WithEncodeTag(mgo.BsonTag), mgo.WithEmitDefaults(true), mgo.WithEnable64ToString(false)),
	}
	m.Migrator()
	return m
}

func (m *Core) Init() {

}

func (m *Core) getDataBase() *mongo.Database {
	op := options.Database()
	//op.SetRegistry(mgo.Registry())

	return mgo.GetMongoDB().Database(AuthDBName, op)
}

func (m *Core) Migrator() {
	allOpts := options.ListDatabases().
		SetNameOnly(true).
		SetAuthorizedDatabases(true)
	nameList, err := mgo.GetMongoDB().ListDatabaseNames(context.Background(), bson.M{}, allOpts)
	if err != nil {
		log.Warnf("NewAuthCore error:%s", err.Error())
		panic(err)
	}

	found := false
	for _, name := range nameList {
		if strings.Compare(name, AuthDBName) == 0 {
			found = true
			break
		}
	}
	if found { // collection
		found = false
		op := options.ListCollections()
		op.SetNameOnly(true)
		nameList, err = mgo.GetMongoDB().Database(AuthDBName).ListCollectionNames(context.Background(), bson.M{}, op)
		if err != nil {
			log.Warnf("NewAuthCore table error:%s", err.Error())
			panic(err)
		}
		for _, name := range nameList {
			if strings.Compare(name, AuthTableName) == 0 {
				found = true
				break
			}
		}
	}

	if !found {
		err = mgo.GetMongoDB().Database(AuthDBName).CreateCollection(context.Background(), AuthTableName)
		if err != nil {
			log.Fatalf("%s Create tableName:%s error:%s", AuthDBName, AuthTableName, err.Error())
			panic(err)
		}
	}
}

func (m *Core) UpdateAuthKey(auth *dataAuth.Auth) (bool, error) {
	updateOp := options.FindOneAndUpdate()
	updateOp.SetUpsert(true)
	updateOp.SetReturnDocument(options.Before)

	sr := m.getDataBase().
		Collection(AuthTableName).
		FindOneAndUpdate(context.Background(), m.bsonEncoder.MarshalCustomSpecMap(auth, "AuthKeyId"),
			bson.M{
				mgo.SETONINSERT: m.bsonEncoder.MarshalCustomSpecMap(auth, "AuthKey", "ExpiresIn", "ValidSince", "ValidUntil", "Salt", "TempAuthKey", "Layer", "Date", "PermAuthKeyId"),
			},
			updateOp)
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		log.Errorf("UpdateAuthKey auth:%+v error:%s", auth, sr.Err().Error())
		return false, sr.Err()
	}

	if sr.Err() == mongo.ErrNoDocuments {
		return true, nil
	}

	beforeAuth := dataAuth.Auth{}
	err := sr.Decode(&beforeAuth)
	if err != nil {
		log.Errorf("UpdateAuthKey Decode auth:%+v error:%s", auth, err.Error())
		return false, err
	}

	log.Debugf("UpdateAuthKey auth:%+v beforeAuth AuthKeyId:%v", auth, beforeAuth.AuthKeyId)
	return beforeAuth.AuthKeyId != auth.AuthKeyId, nil
}

func (m *Core) UpdateAuthKeyWithLayer(authKeyId int64, layer int32) (bool, error) {
	auth := &dataAuth.Auth{AuthKeyId: authKeyId, Layer: layer}

	sr := m.getDataBase().
		Collection(AuthTableName).
		FindOneAndUpdate(context.Background(),
			m.bsonEncoder.MarshalCustomSpecMap(auth, "AuthKeyId"),
			bson.M{mgo.SET: m.bsonEncoder.MarshalCustomSpecMap(auth, "Layer")})
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		log.Errorf("UpdateAuthKeyWithLayer auth:%+v error:%s", auth, sr.Err().Error())
		return false, sr.Err()
	}
	if sr.Err() == mongo.ErrNoDocuments {
		return false, nil
	}

	beforeAuth := dataAuth.Auth{}
	err := sr.Decode(&beforeAuth)
	if err != nil {
		log.Errorf("UpdateAuthKeyWithLayer Decode auth:%+v error:%s", auth, err.Error())
		return false, err
	}

	log.Debugf("UpdateAuthKeyWithLayer auth:%+v beforeAuth AuthKeyId:%v", auth, beforeAuth.AuthKeyId)
	return beforeAuth.AuthKeyId == authKeyId, nil
}

func (m *Core) AuthKey(authKeyId int64) (*dataAuth.Auth, error) {
	auth := &dataAuth.Auth{AuthKeyId: authKeyId}

	sr := m.getDataBase().
		Collection(AuthTableName).
		FindOne(context.Background(),
			m.bsonEncoder.MarshalCustomSpecMap(auth, "AuthKeyId"))
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		log.Errorf("AuthKey auth:%+v error:%s", auth, sr.Err().Error())
		return nil, sr.Err()
	}
	if sr.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}

	err := sr.Decode(auth)
	if err != nil {
		log.Errorf("AuthKey Decode auth:%+v error:%s", auth, err.Error())
		return nil, err
	}

	log.Debugf("AuthKey AuthKeyId:%+v PermAuthKeyId:%d UserId:%d", auth.AuthKeyId, auth.PermAuthKeyId, auth.UserId)
	return auth, nil
}

func (m *Core) BindTempAuthKeyAuthKeyId(authKeyId int64, permAuthKeyId int64, sessionId int64) (*dataAuth.Auth, error) {
	auth := &dataAuth.Auth{AuthKeyId: authKeyId, PermAuthKeyId: permAuthKeyId}

	sr := m.getDataBase().
		Collection(AuthTableName).
		FindOneAndUpdate(context.Background(),
			m.bsonEncoder.MarshalCustomSpecMap(auth, "AuthKeyId"),
			bson.M{mgo.SET: m.bsonEncoder.MarshalCustomSpecMap(auth, "PermAuthKeyId")})
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		log.Errorf("BindTempAuthKeyAuthKeyId auth:%+v error:%s", auth, sr.Err().Error())
		return nil, sr.Err()
	}
	if sr.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}

	auth.AuthKeyId = 0
	err := sr.Decode(&auth)
	if err != nil {
		log.Errorf("BindTempAuthKeyAuthKeyId Decode auth:%+v error:%s", auth, err.Error())
		return nil, err
	}
	permAuth := &dataAuth.Auth{AuthKeyId: permAuthKeyId}
	sr = m.getDataBase().Collection(AuthTableName).FindOne(
		context.Background(),
		m.bsonEncoder.MarshalCustomSpecMap(permAuth, "AuthKeyId"))
	if sr.Err() != nil {
		log.Errorf("BindTempAuthKeyAuthKeyId Find permAuthKeyId:%d error:%s", permAuthKeyId, sr.Err().Error())
		return nil, sr.Err()
	}
	err = sr.Decode(permAuth)
	if err != nil {
		log.Errorf("BindTempAuthKeyAuthKeyId Decode auth permAuthKeyId :%+v error:%s", permAuth, err.Error())
		return nil, err
	}

	auth.UserId = permAuth.UserId
	auth.PermAuthKeyId = permAuthKeyId
	log.Debugf("BindTempAuthKeyAuthKeyId auth userId:%d Layer:%d  beforeAuth AuthKeyId:%v", auth.UserId, auth.Layer, auth.AuthKeyId)
	return auth, nil
}

func (m *Core) BindUser(authKeyId int64, userId int64) (*dataAuth.Auth, error) {
	auth := &dataAuth.Auth{AuthKeyId: authKeyId, UserId: userId}

	sr := m.getDataBase().
		Collection(AuthTableName).
		FindOneAndUpdate(context.Background(),
			m.bsonEncoder.MarshalCustomSpecMap(auth, "AuthKeyId"),
			bson.M{mgo.SET: m.bsonEncoder.MarshalCustomSpecMap(auth, "UserId")})
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		log.Errorf("BindUser auth:%+v error:%s", auth, sr.Err().Error())
		return nil, sr.Err()
	}
	if sr.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}

	err := sr.Decode(auth)
	if err != nil {
		log.Errorf("BindUser Decode auth:%+v error:%s", auth, err.Error())
		return nil, err
	}

	auth.UserId = userId
	log.Debugf("BindUser auth:%+v beforeAuth ", auth)
	return auth, nil
}

func (m *Core) LoadAuthUser(userId []int64) ([]*dataAuth.Auth, error) {
	filter := bson.M{"user_id": bson.M{mgo.IN: userId}}
	cr, err := m.getDataBase().
		Collection(AuthTableName).
		Find(context.Background(), filter)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("LoadAuthUser error:%s", err.Error())
		return nil, err
	}
	defer cr.Close(context.Background())
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	var authList []*dataAuth.Auth
	err = cr.All(context.Background(), &authList)
	if err != nil {
		log.Errorf("LoadAuthUser Decode error:%s", err.Error())
		return nil, err
	}

	log.Debugf("LoadAuthUser len()=%d userId:%d", len(authList), userId)
	return authList, nil
}

func (m *Core) RegisterDevice(authKeyId int64, device *authService.Device) (*dataAuth.Auth, error) {
	devices := &dataAuth.Device{
		TokenType:    device.TokenType,
		AppSandBox:   device.AppSandBox,
		Secret:       device.Secret,
		NoMuted:      device.NoMuted,
		DeviceModel:  device.DeviceModel,
		LangPack:     device.LangPack,
		OtherUidList: device.OtherUidList,
	}
	auth := &dataAuth.Auth{AuthKeyId: authKeyId, Devices: devices}

	sr := m.getDataBase().
		Collection(AuthTableName).
		FindOneAndUpdate(context.Background(),
			m.bsonEncoder.MarshalCustomSpecMap(auth, "AuthKeyId"),
			bson.M{mgo.SET: m.bsonEncoder.MarshalCustomSpecMap(auth, "Devices")})
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		log.Errorf("RegisterDevice auth:%+v error:%s", auth, sr.Err().Error())
		return nil, sr.Err()
	}
	if sr.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}

	err := sr.Decode(auth)
	if err != nil {
		log.Errorf("RegisterDevice Decode auth:%+v error:%s", auth, err.Error())
		return nil, err
	}

	auth.Devices = devices
	log.Debugf("RegisterDevice auth:%+v ", auth)
	return auth, nil
}
