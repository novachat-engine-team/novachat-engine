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

func (m *Core) GetDocument(volumeId int64, localId int32) (*data_sfs.Document, error) {
	log.Debugf("GetDocument volumeId:%d", volumeId)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	document := data_sfs.Document{
		VolumeId: volumeId,
	}
	col := mgo.GetMongoDB().Database(DBSfs).Collection(TableDocument, op)
	sr := col.FindOne(context.Background(), m.encoder.MarshalCustomSpecMap(&document, "VolumeId"))
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		log.Errorf("GetDocument volumeId:%d error:%s", volumeId, sr.Err().Error())
		return nil, sr.Err()
	}

	if sr.Err() == mongo.ErrNoDocuments {
		document.VolumeId = 0
		log.Warnf("not found document volumeId:%d", volumeId)
		return &document, nil
	}

	err := sr.Decode(&document)
	if err != nil {
		document.VolumeId = 0
		log.Warnf("not found document Decode volumeId:%d error:%s", volumeId, err.Error())
	}

	return &document, nil
}

func (m *Core) GetDocumentList(volumeIdList []int64) ([]*data_sfs.Document, error) {
	log.Debugf("GetDocumentList volumeIdList:%+v", volumeIdList)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(DBSfs).Collection(TableDocument, op)
	cursor, err := col.Find(context.Background(), bson.M{"_id": bson.M{mgo.IN: volumeIdList}})
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("GetDocumentList volumeIdList:%+v error:%s", volumeIdList, err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err == mongo.ErrNoDocuments {
		log.Warnf("not found document volumeIdList:%+v", volumeIdList)
		return []*data_sfs.Document{}, nil
	}

	var documents []*data_sfs.Document
	err = cursor.All(context.Background(), &documents)
	if err != nil {
		log.Warnf("not found document Decode volumeIdList:%v error:%s", volumeIdList, err.Error())
	}

	return documents, nil
}

func (m *Core) SaveDocument(documentList []*data_sfs.Document) error {
	idList := make([]int64, 0, len(documentList))
	document := make([]interface{}, 0, len(documentList))
	for _, v := range documentList {
		document = append(document, m.encoder.MarshalCustomSpecMap(v))
		idList = append(idList, v.VolumeId)
	}
	log.Debugf("SaveDocument volumeIdList:%+v", idList)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(DBSfs).Collection(TableDocument, op)
	sr, err := col.InsertMany(context.Background(), document)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("SaveDocument volumeIdList:%v error:%s", idList, err.Error())
		return err
	}

	log.Debugf("%v", sr.InsertedIDs)
	return nil
}
