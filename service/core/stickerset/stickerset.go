/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package stickerset

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"novachat_engine/service/core/sfs"
	data_stickerset "novachat_engine/service/data/stickerset"
	"novachat_engine/service/mgo"
)

type Core struct {
	encoder *mgo.Encoder
}

func NewStickerSetCore(conf *config.MongodbConfig) *Core {
	mgo.InstallMongodb(conf)

	return &Core{
		encoder: mgo.NewEncoder(mgo.WithEncodeTag(mgo.BsonTag), mgo.WithEnable64ToString(false), mgo.WithEmitDefaults(false)),
	}
}

func (s *Core) GetAllStickers() ([]*data_stickerset.StickerSet, error) {

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(sfs.DBSfs).Collection(sfs.TableStickerSet, op)
	cursor, err := col.Find(context.Background(), bson.M{})
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("GetAllStickers error:%s", err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNilDocument {
		return nil, nil
	}

	var l []*data_stickerset.StickerSet
	err = cursor.All(context.Background(), &l)
	if err != nil {
		log.Errorf("GetAllStickers decode error:%s", err.Error())
	}
	return l, nil
}

func (s *Core) GetStickers(idList []int64) ([]*data_stickerset.StickerSet, error) {
	if len(idList) == 0 {
		return nil, nil
	}

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(sfs.DBSfs).Collection(sfs.TableStickerSet, op)
	cursor, err := col.Find(context.Background(), bson.M{"_id": bson.M{mgo.IN: idList}})
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("GetStickers error:%s", err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNilDocument {
		return nil, nil
	}

	var l []*data_stickerset.StickerSet
	err = cursor.All(context.Background(), &l)
	if err != nil {
		log.Errorf("GetStickers decode error:%s", err.Error())
	}
	return l, nil
}

func (s *Core) Installed(userId int64) ([]*data_stickerset.StickerInstall, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(sfs.DBSfs).Collection(sfs.TableStickerInstall, op)
	cursor, err := col.Find(context.Background(), s.encoder.MarshalCustomSpecMap(data_stickerset.StickerInstall{UserId: userId}, "UserId"))
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("Installed error:%s", err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNilDocument {
		return nil, nil
	}

	var l []*data_stickerset.StickerInstall
	err = cursor.All(context.Background(), &l)
	if err != nil {
		log.Errorf("Installed decode error:%s", err.Error())
	}
	return l, nil
}

func (s *Core) Install(userId int64, stickerSetId int64, bInstall bool, date int32) error {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	updateOne := options.Update()
	updateOne.SetUpsert(bInstall)
	col := mgo.GetMongoDB().Database(sfs.DBSfs).Collection(sfs.TableStickerInstall, op)
	ur, err := col.UpdateOne(context.Background(),
		s.encoder.MarshalCustomSpecMap(data_stickerset.StickerInstall{Id: fmt.Sprintf("%d%d", userId, stickerSetId), UserId: userId}, "Id", "UserId"),
		bson.M{mgo.SET: s.encoder.MarshalCustomSpecMap(data_stickerset.StickerInstall{StickerSetId: stickerSetId, Installed: bInstall, Date: date}, "StickerSetId", "Installed", "Date")},
		updateOne)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("Install error:%s", err.Error())
		return err
	}

	_ = ur
	return nil
}

func (s *Core) Faved(userId int64) ([]int64, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(sfs.DBSfs).Collection(sfs.TableStickerRecent, op)
	sr := col.FindOne(context.Background(), s.encoder.MarshalCustomSpecMap(data_stickerset.StickerSetRecent{UserId: userId}, "UserId"))
	if sr.Err() != nil && sr.Err() != mongo.ErrNilDocument {
		log.Errorf("Faved error:%s", sr.Err().Error())
		return nil, sr.Err()
	}
	if sr.Err() == mongo.ErrNilDocument {
		return nil, nil
	}

	var l data_stickerset.StickerSetRecent
	err := sr.Decode(&l)
	if err != nil {
		log.Errorf("Faved decode error:%s", err.Error())
	}
	return l.Faved, nil
}

func (s *Core) UpdateFaved(userId int64, favedList []int64) error {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	uoo := options.Update()
	uoo.SetUpsert(true)

	col := mgo.GetMongoDB().Database(sfs.DBSfs).Collection(sfs.TableStickerRecent, op)
	ur, err := col.UpdateOne(context.Background(),
		s.encoder.MarshalCustomSpecMap(data_stickerset.StickerSetRecent{UserId: userId}, "UserId"),
		s.encoder.MarshalCustomSpecMap(data_stickerset.StickerSetRecent{Faved: favedList}, "Faved"), uoo)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("UpdateFaved error:%s", err.Error())
		return err
	}

	_ = ur
	return nil
}

func (s *Core) Recent(userId int64) ([]int64, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(sfs.DBSfs).Collection(sfs.TableStickerRecent, op)
	sr := col.FindOne(context.Background(), s.encoder.MarshalCustomSpecMap(data_stickerset.StickerSetRecent{UserId: userId}, "UserId"))
	if sr.Err() != nil && sr.Err() != mongo.ErrNilDocument {
		log.Errorf("Recent error:%s", sr.Err().Error())
		return nil, sr.Err()
	}
	if sr.Err() == mongo.ErrNilDocument {
		return nil, nil
	}

	var l data_stickerset.StickerSetRecent
	err := sr.Decode(&l)
	if err != nil {
		log.Errorf("Recent decode error:%s", err.Error())
	}
	return l.Recent, nil
}

func (s *Core) UpdateRecent(userId int64, recentList []int64) error {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	uoo := options.Update()
	uoo.SetUpsert(true)

	col := mgo.GetMongoDB().Database(sfs.DBSfs).Collection(sfs.TableStickerRecent, op)
	ur, err := col.UpdateOne(context.Background(),
		s.encoder.MarshalCustomSpecMap(data_stickerset.StickerSetRecent{UserId: userId}, "UserId"),
		s.encoder.MarshalCustomSpecMap(data_stickerset.StickerSetRecent{Recent: recentList}, "Recent"), uoo)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("UpdateRecent error:%s", err.Error())
		return err
	}

	_ = ur
	return nil
}