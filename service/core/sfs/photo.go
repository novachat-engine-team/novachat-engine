/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package sfs

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/log"
	data_sfs "novachat_engine/service/data/fs"
	"novachat_engine/service/mgo"
)

func (m *Core) GetPhoto(volumeId int64, localId int32) (*data_sfs.Photo, error) {
	log.Debugf("GetPhoto volumeId:%d", volumeId)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	photo := data_sfs.Photo{
		VolumeId: volumeId,
	}
	col := mgo.GetMongoDB().Database(DBSfs).Collection(TablePhoto, op)
	sr := col.FindOne(context.Background(), m.encoder.MarshalCustomSpecMap(&photo, "VolumeId"))
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		log.Errorf("GetPhoto volumeId:%d error:%s", volumeId, sr.Err().Error())
		return nil, sr.Err()
	}

	if sr.Err() == mongo.ErrNoDocuments {
		photo.VolumeId = 0
		log.Warnf("not found photo volumeId:%d", volumeId)
		return &photo, nil
	}

	err := sr.Decode(&photo)
	if err != nil {
		photo.VolumeId = 0
		log.Warnf("not found photo Decode volumeId:%d error:%s", volumeId, err.Error())
	}

	return &photo, nil
}

func (m *Core) GetPhotoList(volumeIdList []int64) ([]*data_sfs.Photo, error) {
	log.Debugf("GetPhotoList volumeIdList:%+v", volumeIdList)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(DBSfs).Collection(TablePhoto, op)
	cursor, err := col.Find(context.Background(), bson.M{"_id": bson.M{mgo.IN: volumeIdList}})
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("GetPhotoList volumeIdList:%+v error:%s", volumeIdList, err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNoDocuments {
		log.Warnf("not found photo volumeIdList:%+v", volumeIdList)
		return []*data_sfs.Photo{}, nil
	}

	var photos []*data_sfs.Photo
	err = cursor.All(context.Background(), &photos)
	if err != nil {
		log.Warnf("not found photo Decode volumeIdList:%v error:%s", volumeIdList, err.Error())
	}

	return photos, nil
}

func (m *Core) SavePhoto(photoList []*data_sfs.Photo) error {
	idList := make([]int64, 0, len(photoList))
	document := make([]interface{}, 0, len(photoList))
	for _, v := range photoList {
		document = append(document, m.encoder.MarshalCustomSpecMap(v))
		idList = append(idList, v.VolumeId)
	}
	log.Debugf("SavePhoto volumeIdList:%+v", idList)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(DBSfs).Collection(TablePhoto, op)
	sr, err := col.InsertMany(context.Background(), document)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("SavePhoto volumeIdList:%v error:%s", idList, err.Error())
		return err
	}

	log.Debugf("%v", sr.InsertedIDs)
	return nil
}
