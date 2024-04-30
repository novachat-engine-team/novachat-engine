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
	"go.mongodb.org/mongo-driver/mongo/options"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	"novachat_engine/service/mgo"
	"strings"
)

const (
	DBSfs               = "db_file"
	TablePhoto          = "tb_photo"
	TableDocument       = "tb_document"
	TableMD5            = "tb_md5sum"
	TableStickerSet     = "tb_stickerset"
	TableStickerInstall = "tb_stickerinstall"
	TableStickerRecent  = "tb_stickerrecent"
)

type Core struct {
	encoder *mgo.Encoder
}

func NewSfsCore(conf *config.MongodbConfig) *Core {
	mgo.InstallMongodb(conf)
	return &Core{
		encoder: mgo.NewEncoder(mgo.WithEncodeTag(mgo.BsonTag), mgo.WithEnable64ToString(false), mgo.WithEmitDefaults(false)),
	}
}

func (m *Core) Init() {}
func (m *Core) Migrator() {
	allOpts := options.ListDatabases().
		SetNameOnly(true).
		SetAuthorizedDatabases(true)
	nameList, err := mgo.GetMongoDB().ListDatabaseNames(context.Background(), bson.M{}, allOpts)
	if err != nil {
		log.Warnf("NewSfsCore error:%s", err.Error())
		panic(err)
	}

	found := false
	for _, name := range nameList {
		if strings.Compare(name, DBSfs) == 0 {
			found = true
			break
		}
	}
	if found { // collection
		found = false
		op := options.ListCollections()
		op.SetNameOnly(true)
		nameList, err = mgo.GetMongoDB().Database(DBSfs).ListCollectionNames(context.Background(), bson.M{}, op)
		if err != nil {
			log.Warnf("NewSfsCore table error:%s", err.Error())
			panic(err)
		}
	}

	for _, tableName := range []string{TableDocument, TablePhoto, TableMD5} {
		if util.IndexStrings(&nameList, tableName) < 0 {
			err = mgo.GetMongoDB().Database(DBSfs).CreateCollection(context.Background(), tableName)
			if err != nil {
				log.Fatalf("%s Create tableName:%s error:%s", DBSfs, tableName, err.Error())
				panic(err)
			}
		}
	}
}
