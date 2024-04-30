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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/log"
	data_sfs "novachat_engine/service/data/fs"
	"novachat_engine/service/mgo"
)

func (m *Core) GetMd5sum(md5sum string) (*data_sfs.Md5Sum, error) {
	log.Debugf("GetMd5sum md5sum:%s", md5sum)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	value := data_sfs.Md5Sum{
		Md5Sum: md5sum,
	}
	col := mgo.GetMongoDB().Database(DBSfs).Collection(TableMD5, op)
	sr := col.FindOne(context.Background(), m.encoder.MarshalCustomSpecMap(&value, "Md5Sum"))
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		log.Errorf("GetMd5sum md5sum:%s error:%s", md5sum, sr.Err().Error())
		return nil, sr.Err()
	}

	if sr.Err() == mongo.ErrNoDocuments {
		value.Md5Sum = ""
		log.Warnf("not found Md5sum md5sum:%s", md5sum)
		return &value, nil
	}

	err := sr.Decode(&value)
	if err != nil {
		value.Md5Sum = ""
		log.Warnf("not found Md5sum decode md5sum:%s error:%s", md5sum, err.Error())
	}

	return &value, nil
}

func (m *Core) SaveMd5sum(md5Sum *data_sfs.Md5Sum) error {

	log.Debugf("SaveMd5sum md5Sum:%+v", md5Sum)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetMongoDB().Database(DBSfs).Collection(TableMD5, op)
	sr, err := col.InsertOne(context.Background(), md5Sum)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("SaveMd5sum volumeIdList:%v error:%s", md5Sum.Md5Sum, err.Error())
		return err
	}

	log.Debugf("%v", sr.InsertedID)
	return nil
}
