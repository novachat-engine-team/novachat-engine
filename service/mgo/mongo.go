/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/18 18:32
 * @File : mongo.go
 */

package mgo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"strings"
	"time"
)

type MongoDB struct {
	conf *config.MongodbConfig
	*mongo.Client
}

func (m *MongoDB) pingLoop() {

	for {
		select {
		case <-time.NewTimer(3 * time.Second).C:
			err := m.Ping(context.TODO(), nil)
			if err != nil {
				log.Errorf("MongoDB ping error:%s", err.Error())
				//	return
			}
		}
	}
}

func formatHosts(hosts []string) string {
	return strings.Join(hosts, ",")
}

func NewMongoDB(conf *config.MongodbConfig) *MongoDB {
	m := &MongoDB{
		conf: conf,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	oc := options.Client()
	var client *mongo.Client
	var err error

	if conf.Direct {
		direct := ""
		direct = "direct"
		client, err = mongo.Connect(ctx, oc.ApplyURI(
			fmt.Sprintf("mongodb://%s/?connect=%s&maxpoolsize=%d&minpoolsize=%d&replicaset=%s&username=%s&password=%s&database=%s",
				formatHosts(conf.Addr),
				direct,
				conf.MaxPoolSize,
				conf.MinPoolSize,
				conf.ReplicaSet,
				conf.Username,
				conf.Password,
				conf.Database)))
		if err != nil {
			panic(err)
		}
	} else {
		if len(conf.Username) > 0 {
			credential := options.Credential{
				Username: conf.Username,
				Password: conf.Password,
			}
			co := oc.ApplyURI(fmt.Sprintf("mongodb://%s", formatHosts(conf.Addr))).
				SetAuth(credential).
				SetMaxPoolSize(uint64(conf.MaxPoolSize)).
				SetMinPoolSize(uint64(conf.MinPoolSize))
			client, err = mongo.Connect(ctx, co)
			if err != nil {
				panic(err)
			}
		} else {
			client, err = mongo.Connect(ctx, oc.ApplyURI(
				fmt.Sprintf("mongodb://%s/?maxpoolsize=%d&minpoolsize=%d&replicaset=%s&database=%s",
					formatHosts(conf.Addr),
					conf.MaxPoolSize,
					conf.MinPoolSize,
					conf.ReplicaSet,
					conf.Database)))
			if err != nil {
				panic(err)
			}
		}
	}

	ctx1, cancel1 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel1()
	err = client.Ping(ctx1, nil)
	if err != nil {
		panic(err)
	}

	m.Client = client
	go m.pingLoop()
	return m
}

func (m *MongoDB) InsertOne(dbName string, collectionName string, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	db := m.Database(dbName)
	col, _ := db.Collection(collectionName).Clone()
	insertOneResult, err := col.InsertOne(context.Background(), document, opts...)
	return insertOneResult, err
}

func (m *MongoDB) InsertMany(dbName string, collectionName string, documents []interface{},
	opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	db := m.Database(dbName)
	col, _ := db.Collection(collectionName).Clone()
	insertResult, err := col.InsertMany(context.Background(), documents, opts...)
	return insertResult, err
}

func (m *MongoDB) Find(dbName string, collectionName string, filter interface{}, result interface{}, opts ...*options.FindOptions) error {
	db := m.Database(dbName)
	col, _ := db.Collection(collectionName).Clone()
	cursor, err := col.Find(context.Background(), filter, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())
	err = cursor.All(context.Background(), result)
	if err != nil {
		return err
	}

	return nil
}

type IndexKey struct {
	Key                string
	Order              int32
	ExpireAfterSeconds int32
	Unique             bool
}

func (m *MongoDB) CreateIndexes(dbName string, collectionName string, keys []IndexKey, background bool) error {
	if len(keys) == 0 {
		return nil
	}

	indexMap := make(map[string]IndexKey, len(keys))
	for _, k := range keys {
		indexMap[fmt.Sprintf("%s_%d", k.Key, k.Order)] = k
	}
	find := func(key string, result []*mongo.IndexSpecification) bool {
		for _, r := range result {
			if r.Name == key {
				return true
			}
		}
		return false
	}
	db := m.Database(dbName)
	col, _ := db.Collection(collectionName).Clone()

	indexView := col.Indexes()
	var results []*mongo.IndexSpecification
	cur, err := indexView.List(context.Background())
	if err == nil {
		err = cur.All(context.Background(), &results)
		if err != nil {
			cur.Close(context.Background())
			return err
		}

		for k, v := range indexMap {
			if find(k, results) == false {
				io := &options.IndexOptions{Background: &background}
				if v.ExpireAfterSeconds > 0 {
					io.SetExpireAfterSeconds(v.ExpireAfterSeconds)
				}
				if v.Unique {
					io.SetUnique(v.Unique)
				}

				indexView.CreateOne(context.Background(), mongo.IndexModel{
					Keys:    bson.M{v.Key: v.Order},
					Options: io,
				})
			}
		}

		cur.Close(context.Background())
		return nil
	} else {
		return err
	}
}

func (m *MongoDB) CreateManyIndexes(dbName string, collectionName string, keys [][]*IndexKey, background bool) error {
	if len(keys) == 0 {
		return nil
	}

	indexMap := make(map[string][]*IndexKey, len(keys))
	builder := strings.Builder{}

	for _, k := range keys {
		builder.Reset()
		for idx, kk := range k {
			if idx == len(k)-1 {
				builder.WriteString(fmt.Sprintf("%s_%d", kk.Key, kk.Order))
			} else {
				builder.WriteString(fmt.Sprintf("%s_%d_", kk.Key, kk.Order))
			}
		}

		indexMap[builder.String()] = k
	}
	find := func(key string, result []*mongo.IndexSpecification) bool {
		for _, r := range result {
			if r.Name == key {
				return true
			}
		}
		return false
	}
	db := m.Database(dbName)
	col, _ := db.Collection(collectionName).Clone()

	indexView := col.Indexes()
	var results []*mongo.IndexSpecification
	cur, err := indexView.List(context.Background())
	if err == nil {
		err = cur.All(context.Background(), &results)
		if err != nil {
			cur.Close(context.Background())
			return err
		}

		for k, v := range indexMap {
			if find(k, results) == false {
				kM := bson.D{}
				expireAfterSeconds := int32(0)
				for _, vv := range v {
					kM = append(kM, bson.E{Key: vv.Key, Value: vv.Order})
					if expireAfterSeconds == 0 {
						expireAfterSeconds = vv.ExpireAfterSeconds
					}
				}
				io := &options.IndexOptions{Background: &background}
				if expireAfterSeconds > 0 {
					io.SetExpireAfterSeconds(expireAfterSeconds)
				}
				s, err := indexView.CreateOne(context.Background(), mongo.IndexModel{
					Keys:    kM,
					Options: io,
				})
				if err != nil {
					log.Errorf("CreateIndex dbName:%s collectionName:%s key:%s background:%v error: name(%s) %s",
						dbName, collectionName, k, background, s, err.Error())
				}
			}
		}

		cur.Close(context.Background())
		return nil
	} else {
		return err
	}
}
