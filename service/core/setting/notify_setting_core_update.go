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
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/log"
	"novachat_engine/service/data/setting"
	"novachat_engine/service/mgo"
)

func (c *Core) UpdateNotifySetting(userId int64, peerId int64, peerType int32, peerNotifyType int32, previews bool, silent bool, until int32, sound string) (bool, error) {
	data := &data_setting.Setting{
		Id:       makeNotifySettingId(userId, peerId, peerNotifyType),
		UserId:   userId,
		PeerId:   peerId,
		PeerType: peerType,
		Setting: &data_setting.NotifySetting{
			ShowPreviews: previews,
			Silent:       silent,
			MuteUntil:    until,
			Sound:        sound,
		},
		NotifyType: peerNotifyType,
	}
	log.Debugf("UpdateNotifySetting data:%v", data)
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	//op.SetRegistry(mgo.Registry())
	updateOptions := options.Update()
	updateOptions.SetUpsert(true)

	col := mgo.GetMongoDB().Database(DBSetting).Collection(TableNameNotifySetting, op)
	result, err := col.UpdateOne(context.Background(),
		mgo.DBE.MarshalCustomSpecMap(data, "Id"),
		bson.M{
			mgo.SET: mgo.DBE.MarshalCustomSpecMap(data, "UserId", "PeerId", "NotifyType", "Setting"),
		},
		updateOptions)
	if err != nil {
		log.Errorf("UpdateNotifySetting data:%v error:%s", data, err.Error())
		return false, err
	}
	_ = result

	return true, nil
}

func (c *Core) ResetNotifySetting(userId int64, notifyType int32) (bool, error) {
	log.Debugf("ResetNotifySetting userId:%v", userId)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	//op.SetRegistry(mgo.Registry())
	updateOptions := options.Update()
	updateOptions.SetUpsert(true)

	col := mgo.GetMongoDB().Database(DBSetting).Collection(TableNameNotifySetting, op)
	result, err := col.DeleteMany(
		context.Background(),
		bson.M{"_id": bson.M{mgo.NE: ""}, "user_id": userId},
	)
	if err != nil {
		log.Errorf("ResetNotifySetting userId:%v error:%s", userId, err.Error())
		return false, err
	}
	_ = result

	return true, nil
}
