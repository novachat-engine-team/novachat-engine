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

package setting

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/log"
	"novachat_engine/service/data/setting"
	"novachat_engine/service/mgo"
)

type NotifySettingListType struct {
	UserId         int64
	PeerId         int64
	PeerType       int32
	PeerNotifyType int32
}

func (c *Core) GetNotifySettingList(v []*NotifySettingListType) ([]*data_setting.Setting, error) {
	idList := make([]string, 0, len(v))
	for _, vv := range v {
		idList = append(idList, makeNotifySettingId(vv.UserId, vv.PeerId, vv.PeerNotifyType))
	}

	log.Debugf("GetNotifySettingList data:%v", idList)
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(DBSetting).Collection(TableNameNotifySetting, op)
	cursor, err := col.Find(
		context.Background(),
		bson.M{"_id": bson.M{mgo.IN: idList}})
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("GetNotifySettingList request:%+v error:%s", idList, err)
		return nil, err
	}

	var data []*data_setting.Setting
	if err == nil {
		defer cursor.Close(context.Background())
		err = cursor.All(context.Background(), &data)
		if err != nil {
			log.Errorf("GetNotifySettingList Decode data:%+v error:%s", data, err.Error())
			return nil, err
		}
	}

	log.Infof("GetNotifySettingList data:%+v", data)
	return data, nil
}

func (c *Core) GetNotifySetting(userId int64, peerId int64, peerNotifyType int32) (*data_setting.NotifySetting, error) {
	data := &data_setting.Setting{
		Id:         makeNotifySettingId(userId, peerId, peerNotifyType),
		UserId:     userId,
		PeerId:     peerId,
		PeerType:   0,
		NotifyType: peerNotifyType,
	}

	log.Debugf("GetNotifySetting data:%v", data)
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(DBSetting).Collection(TableNameNotifySetting, op)
	singleResult := col.FindOne(
		context.Background(),
		mgo.DBE.MarshalCustomSpecMap(data, "Id"))
	if singleResult.Err() != nil && singleResult.Err() != mongo.ErrNoDocuments {
		log.Errorf("GetNotifySetting request:%+v error:%s", data, singleResult.Err().Error())
		return nil, singleResult.Err()
	}
	if singleResult.Err() == nil {
		err := singleResult.Decode(data)
		if err != nil {
			log.Errorf("GetNotifySetting Decode data:%+v error:%s", data, err.Error())
			return nil, singleResult.Err()
		}
	}

	log.Infof("GetNotifySetting data:%+v", data)
	return data.Setting, nil
}
