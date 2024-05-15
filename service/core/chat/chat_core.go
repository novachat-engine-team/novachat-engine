/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package chat

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
	data_chat "novachat_engine/service/data/chat"
	"novachat_engine/service/mgo"
	"strings"
)

const (
	DBChats               = "db_chats"
	TableChats            = "tb_chats"
	TableChatsConfig      = "tb_chats_config"
	TableChatsUsername    = "tb_chats_username"
	TableChatsParticipant = "tb_chats_participant"
	TableChatsConfigID    = "tb_chats_config_id"
)

func makeChatParticipantKey(chatId int64, userId int64) string {
	return fmt.Sprintf("%d%d", chatId, userId)
}

func init() {
	mgo.DBE.Enable64ToString = false
}

type Core struct {
}

func (c *Core) GetChat(chatId int64) (*data_chat.Chat, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	chat := &data_chat.Chat{
		ChatId: chatId,
	}
	rs := mgo.GetDatabase(DBChats).
		Collection(TableChats, op).
		FindOne(context.TODO(), mgo.DBE.MarshalCustomSpecMap(chat, "ChatId"))
	if rs.Err() != nil {
		log.Fatalf("GetChat chatId:%d error:%s", chatId, rs.Err().Error())
		return nil, rs.Err()
	}
	err := rs.Decode(chat)
	if err != nil {
		log.Fatalf("GetChat chatId:%d decode error:%s", chatId, err.Error())
		return nil, err
	}
	return chat, nil
}

func (c *Core) GetChatParticipants(chatId int64, userIdList []int64) ([]*data_chat.ChatParticipant, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	chatParticipant := &data_chat.ChatParticipant{
		ChatId: chatId,
	}
	filter := mgo.DBE.MarshalCustomSpecMap(chatParticipant, "ChatId")
	if len(userIdList) > 0 {
		filter["user_id"] = bson.M{mgo.IN: userIdList}
	}
	cursor, err := mgo.GetDatabase(DBChats).
		Collection(TableChatsParticipant, op).
		Find(context.TODO(), filter)
	if err != nil {
		log.Fatalf("GetChatParticipants chatId:%d error:%s", chatId, err.Error())
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var dataList []*data_chat.ChatParticipant
	err = cursor.All(context.TODO(), &dataList)
	if err != nil {
		log.Fatalf("GetChatParticipants chatId:%d decode error:%s", chatId, err.Error())
		return nil, err
	}
	return dataList, nil
}

func (c *Core) CreateChat(chatInfo *data_chat.Chat, chatParticipantList []*data_chat.ChatParticipant) (*data_chat.Chat, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())
	opfu := options.FindOneAndUpdate()
	opfu.SetReturnDocument(options.After)
	opfu.SetUpsert(true)

	chatConfig := data_chat.ChatsConfig{
		ChatIdKey: TableChatsConfigID,
	}
	chatIdInc := int64(1)
	db := mgo.GetDatabase(DBChats)
	sr := db.Collection(TableChatsConfig, op).FindOneAndUpdate(
		context.TODO(),
		mgo.DBE.MarshalCustomSpecMap(&chatConfig, "ChatIdKey"),
		bson.M{mgo.INC: bson.M{"chat_id": chatIdInc}}, opfu,
	)
	if sr.Err() != nil {
		log.Fatalf("CreateChat error:%s", sr.Err().Error())
		return nil, sr.Err()
	}

	err := sr.Decode(&chatConfig)
	if err != nil {
		log.Errorf("CreateChat decode chatConfig error:%s", err.Error())
		return nil, err
	}

	chatInfo.ChatId = chatConfig.ChatId

	var is *mongo.InsertOneResult
	var imr *mongo.InsertManyResult
	var documents = make([]interface{}, 0, len(chatParticipantList))
	for _, v := range chatParticipantList {
		v.ChatId = chatConfig.ChatId
		v.ChatParticipantKey = makeChatParticipantKey(v.ChatId, v.UserId)

		documents = append(documents, mgo.DBE.MarshalCustomSpecMap(v))
	}

	err = mgo.GetMongoDB().UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
		sessionContext.StartTransaction()

		is, err = db.Collection(TableChats).InsertOne(sessionContext, mgo.DBE.MarshalCustomSpecMap(chatInfo))
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)

			log.Errorf("CreateChat with participant InsertOne error:%s", err.Error())
			return err
		}
		_ = is

		if len(documents) > 0 {
			imr, err = db.Collection(TableChatsParticipant).InsertMany(sessionContext, documents)
			if err != nil {
				sessionContext.AbortTransaction(sessionContext)

				log.Errorf("CreateChat with participant InsertMany error:%s", err.Error())
				return err
			}
			_ = imr
		}

		sessionContext.CommitTransaction(sessionContext)
		return nil
	})
	if err != nil {
		log.Errorf("CreateChat with participant error:%s", err.Error())
		return nil, err
	}
	return chatInfo, nil
}

func NewChatCore(conf *config.MongodbConfig) *Core {
	mgo.InstallMongodb(conf)

	m := &Core{}
	m.Migrator()
	m.CreateIndex()
	return m
}

func (c *Core) Init() {

}

func (c *Core) CreateIndex() {
	indexView := mgo.GetMongoDB().Database(DBChats).Collection(TableChats).Indexes()
	nameList, err := indexView.CreateMany(context.TODO(), []mongo.IndexModel{{
		Keys: bson.D{
			{"chat_id", 1},
		},
	}})
	if err != nil {
		log.Warnf("Create Index error:%s", err.Error())
	} else {
		log.Debugf("Create Index nameList:%s", nameList)
	}

	indexView = mgo.GetMongoDB().Database(DBChats).Collection(TableChatsUsername).Indexes()
	nameList, err = indexView.CreateMany(context.TODO(), []mongo.IndexModel{{
		Keys: bson.D{
			{"chat_id", 1},
		},
	}})
	if err != nil {
		log.Warnf("Create Index error:%s", err.Error())
	} else {
		log.Debugf("Create Index nameList:%s", nameList)
	}
}

func (c *Core) Migrator() {
	allOpts := options.ListDatabases().
		SetNameOnly(true).
		SetAuthorizedDatabases(true)
	nameList, err := mgo.GetMongoDB().ListDatabaseNames(context.Background(), bson.M{}, allOpts)
	if err != nil {
		log.Warnf("NewChatCore error:%s", err.Error())
		panic(err)
	}

	found := false
	for _, name := range nameList {
		if strings.Compare(name, DBChats) == 0 {
			found = true
			break
		}
	}
	if found { // collection
		found = false
		op := options.ListCollections()
		op.SetNameOnly(true)
		nameList, err = mgo.GetMongoDB().Database(DBChats).ListCollectionNames(context.Background(), bson.M{}, op)
		if err != nil {
			log.Warnf("NewAuthCore table error:%s", err.Error())
			panic(err)
		}
		for _, name := range nameList {
			if strings.Compare(name, TableChats) == 0 {
				found = true
				break
			}
		}
	}

	if !found {
		err = mgo.GetMongoDB().Database(DBChats).CreateCollection(context.Background(), TableChats)
		if err != nil {
			log.Fatalf("%s Create tableName:%s error:%s", DBChats, TableChats, err.Error())
			panic(err)
		}
	}

	err = mgo.GetMongoDB().Database(DBChats).
		CreateCollection(context.Background(), TableChatsConfig)
	err = mgo.GetMongoDB().Database(DBChats).
		CreateCollection(context.Background(), TableChatsParticipant)
	err = mgo.GetMongoDB().Database(DBChats).
		CreateCollection(context.Background(), TableChatsUsername)
	_ = err
}

func (c *Core) EditProperty(chatId int64, chatInfo data_chat.Chat, property ...string) (bool, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())

	ur, err := mgo.GetDatabase(DBChats).
		Collection(TableChats, op).
		UpdateOne(context.TODO(), mgo.DBE.MarshalCustomSpecMap(chatInfo, "ChatId"),
			mgo.DBE.MarshalCustomSpecMap(chatInfo, property...))
	if err != nil {
		log.Fatalf("EditProperty chatId:%d error:%s", chatId, err.Error())
		return false, err
	}
	_ = ur
	log.Debugf("EditProperty chat:%v property:%+v", chatInfo, property)
	return true, nil
}

func (c *Core) ModifyParticipant(chatId int64, chatParticipantList []*data_chat.ChatParticipant, modifyChatParticipantList []*data_chat.ChatParticipant) (bool, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())

	models := make([]mongo.WriteModel, 0, len(chatParticipantList)+len(modifyChatParticipantList))
	for _, v := range chatParticipantList {
		v.ChatId = chatId
		v.ChatParticipantKey = makeChatParticipantKey(v.ChatId, v.UserId)

		insertOne := mongo.NewInsertOneModel()
		insertOne.SetDocument(mgo.DBE.MarshalCustomSpecMap(v))
		models = append(models, insertOne)
	}

	for _, v := range modifyChatParticipantList {
		v.ChatId = chatId
		v.ChatParticipantKey = makeChatParticipantKey(v.ChatId, v.UserId)

		updateOne := mongo.NewUpdateOneModel()
		updateOne.SetFilter(mgo.DBE.MarshalCustomSpecMap(v, "ChatParticipantKey"))
		update := mgo.DBE.MarshalCustomSpecMap(v)
		delete(update, "_id")
		updateOne.SetUpdate(update)
		models = append(models, updateOne)
	}

	bulkWriteResult, err := mgo.GetDatabase(DBChats).
		Collection(TableChats, op).
		BulkWrite(context.Background(), models)
	if err != nil {
		log.Fatalf("ModifyParticipant chatId:%d error:%s", chatId, err.Error())
		return false, err
	}
	_ = bulkWriteResult
	return true, nil
}

func (c *Core) GetChatByName(username string) (*data_chat.Chat, error) {

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())

	chat := &data_chat.Chat{Username: username}

	db := mgo.GetDatabase(DBChats)
	sr := db.Collection(TableChats, op).
		FindOne(context.TODO(), mgo.DBE.MarshalCustomSpecMap(chat, "Username"))
	if sr.Err() != nil {
		log.Errorf("GetChatByName username:%d error:%s", username, sr.Err().Error())
		return nil, sr.Err()
	}

	err := sr.Decode(chat)
	if err != nil {
		log.Fatalf("GetChatByName Decode username:%d error:%s", username, err.Error())
		return nil, sr.Err()
	}

	return chat, nil
}
