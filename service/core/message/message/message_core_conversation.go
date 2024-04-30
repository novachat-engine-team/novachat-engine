/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package message

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message"
	"novachat_engine/service/data/messages/message"
	"novachat_engine/service/mgo"
)

func (c *Core) GetConversationList(userId int64, excludePinned bool, peerType constants.PeerType, offsetDate int32, limit int64, foldId int32) ([]*data_message.Conversation, error) {

	log.Debugf("GetConversationList userId:%d excludePinned:%v peerType:%v offsetDate:%d limit:%d", userId, excludePinned, peerType, offsetDate, limit)
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	filter := bson.M{"date": bson.M{mgo.GTE: offsetDate}, "user_id": userId, "folder_id": foldId}
	if peerType != constants.PeerTypeEmpty {
		filter["peer_type"] = peerType.ToInt32()
	}
	if !excludePinned {
		filter["pinned"] = true
	}
	opf := options.Find()
	if limit > 0 {
		opf.SetLimit(limit)
	}
	cursor, err := mgo.GetMongoDB().Database(message.DBMessage).
		Collection(message.TableName(userId, message.TableConversation), op).
		Find(context.Background(), filter, opf)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("GetConversationList userId:%d excludePinned:%v peerType:%v offsetDate:%d limit:%d error:%s", userId, excludePinned, peerType, offsetDate, limit, err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNilDocument {
		log.Infof("GetConversationList userId:%d excludePinned:%v peerType:%v offsetDate:%d limit:%d empty", userId, excludePinned, peerType, offsetDate, limit)
		return nil, nil
	}

	var data []*data_message.Conversation
	err = cursor.All(context.Background(), &data)

	if err != nil {
		log.Warnf("GetConversationList userId:%d excludePinned:%v peerType:%v offsetDate:%d limit:%d error:%s", userId, excludePinned, peerType, offsetDate, limit, err.Error())
	}
	for idx := range data {
		data[idx].PeerId = message.MakePeerId(data[idx].PeerId, constants.PeerType(data[idx].PeerType))
	}

	log.Infof("GetConversationList userId:%d excludePinned:%v peerType:%v offsetDate:%d limit:%d data:%+v", userId, excludePinned, peerType, offsetDate, limit, data)
	return data, nil
}

func (c *Core) GetConversationByPeerList(userId int64, list []*data_message.Conversation) ([]*data_message.Conversation, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	idList := make([]int64, 0, len(list))
	for _, v := range list {
		v.PeerId = message.MakePeerId(v.PeerId, constants.PeerType(v.PeerType))
		//v.Id = message.MakeDialogId(userId, v.PeerId, constants.PeerType(v.PeerType))
		idList = append(idList, v.PeerId)
	}
	filter := bson.M{}
	filter["user_id"] = userId
	filter["peer_id"] = bson.M{mgo.IN: idList}
	filter["date"] = bson.M{mgo.GTE: 0}

	cursor, err := mgo.GetMongoDB().Database(message.DBMessage).
		Collection(message.TableName(userId, message.TableConversation), op).
		Find(context.Background(), filter)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("GetConversationByPeerList userId:%d error:%s", userId, err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNilDocument {
		log.Infof("GetConversationByPeerList userId:%d empty", userId)
		return nil, nil
	}

	var data []*data_message.Conversation
	err = cursor.All(context.Background(), &data)
	if err != nil {
		log.Warnf("GetConversationByPeerList userId:%d error:%s", userId, err.Error())
	}
	for idx := range data {
		data[idx].PeerId = message.MakePeerId(data[idx].PeerId, constants.PeerType(data[idx].PeerType))
	}

	log.Infof("GetConversationByPeerList userId:%d data:%+v", userId, data)
	return data, nil
}

func (c *Core) GetConversation(userId int64, conversation *data_message.Conversation) (*data_message.Conversation, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	conversation.UserId = userId
	conversation.PeerId = message.MakePeerId(conversation.PeerId, constants.PeerType(conversation.PeerType))
	//conversation.Id = message.MakeDialogId(userId, conversation.PeerId, constants.PeerType(conversation.PeerType))
	filter := mgo.DBE.MarshalCustomSpecMap(conversation, "UserId", "PeerId")
	sr := mgo.GetMongoDB().Database(message.DBMessage).
		Collection(message.TableName(userId, message.TableConversation), op).
		FindOne(context.Background(), filter)
	if sr.Err() != nil && sr.Err() != mongo.ErrNilDocument {
		log.Errorf("GetConversation userId:%d error:%s", userId, sr.Err().Error())
		return nil, sr.Err()
	}
	var data data_message.Conversation
	if sr.Err() == mongo.ErrNilDocument {
		log.Infof("GetConversation userId:%d empty", userId)
		return &data, nil
	}

	err := sr.Decode(&data)
	if err != nil {
		log.Warnf("GetConversation userId:%d error:%s", userId, err.Error())
	}

	data.PeerId = message.MakePeerId(conversation.PeerId, constants.PeerType(conversation.PeerType))
	log.Infof("GetConversation userId:%d data:%+v", userId, data)
	return &data, nil
}

func (c *Core) ConversationSaveDraft(userId int64, peerId int64, peerType constants.PeerType, draft string) error {
	conversation := data_message.Conversation{
		Id:       message.MakeDialogId(userId, peerId, peerType),
		UserId:   userId,
		PeerId:   message.MakePeerId(peerId, peerType),
		PeerType: peerType.ToInt32(),
		Draft:    draft,
	}

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())
	opu := options.Update()
	opu.SetUpsert(true)

	filter := mgo.DBE.MarshalCustomSpecMap(conversation, "UserId", "PeerId")
	update := bson.M{mgo.SET: mgo.DBE.MarshalCustomSpecMap(conversation, "Draft"),
		mgo.SETONINSERT: mgo.DBE.MarshalCustomSpecMap(conversation, "Id", "Pinned", "FolderId", "PinIds", "PeerType", "OutboxMaxId", "Date", "Pts", "Top", "InboxMaxId", "UnreadCount")}
	ur, err := mgo.GetMongoDB().Database(message.DBMessage).
		Collection(message.TableName(userId, message.TableConversation), op).
		UpdateOne(context.Background(), filter, update)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("ConversationSaveDraft userId:%d error:%s", userId, err)
		return err
	}
	_ = ur
	return nil
}

func (c *Core) ConversationClearAllDraft(userId int64) error {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	filter := bson.M{"user_id": userId}
	update := bson.M{"draft": ""}
	ur, err := mgo.GetMongoDB().Database(message.DBMessage).
		Collection(message.TableName(userId, message.TableConversation), op).
		UpdateMany(context.Background(), filter, update)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("ConversationClearAllDraft userId:%d error:%s", userId, err)
		return err
	}
	_ = ur
	return nil
}

func (c *Core) ConversationPinnedMessage(ctx context.Context, userId int64, peerId int64, peerType constants.PeerType, messages []*data_message.PinnedMessage) error {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())
	opu := options.Update()
	opu.SetUpsert(true)

	conversation := data_message.Conversation{
		Id:       message.MakeDialogId(userId, peerId, peerType),
		UserId:   userId,
		PeerId:   message.MakePeerId(peerId, peerType),
		PeerType: peerType.ToInt32(),
		PinIds:   messages,
	}
	filter := mgo.DBE.MarshalCustomSpecMap(conversation, "UserId", "PeerId")
	update := bson.M{mgo.SET: mgo.DBE.MarshalCustomSpecMap(conversation, "PinIds"),
		mgo.SETONINSERT: mgo.DBE.MarshalCustomSpecMap(conversation, "Id", "Pinned", "FolderId", "Draft", "PeerType", "OutboxMaxId", "Date", "Pts", "Top", "InboxMaxId", "UnreadCount")}
	ur, err := mgo.GetMongoDB().Database(message.DBMessage).
		Collection(message.TableName(userId, message.TableConversation), op).
		UpdateOne(ctx, filter, update, opu)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("ConversationPinnedMessage userId:%d error:%s", userId, err)
		return err
	}
	_ = ur
	return nil
}

func (c *Core) ConversationUpdatePts(ctx context.Context, userId int64, peerId int64, peerType constants.PeerType, pts int32) error {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	conversation := data_message.Conversation{
		Id:       message.MakeDialogId(userId, peerId, peerType),
		UserId:   userId,
		PeerId:   message.MakePeerId(peerId, peerType),
		PeerType: peerType.ToInt32(),
		Pts:      pts,
	}
	filter := mgo.DBE.MarshalCustomSpecMap(conversation, "UserId", "PeerId")
	filter["pts"] = bson.M{mgo.GT: pts}
	update := bson.M{mgo.SET: mgo.DBE.MarshalCustomSpecMap(conversation, "Pts")}
	ur, err := mgo.GetMongoDB().Database(message.DBMessage).
		Collection(message.TableName(userId, message.TableConversation), op).
		UpdateOne(ctx, filter, update)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("ConversationPinnedMessage userId:%d error:%s", userId, err)
		return err
	}
	_ = ur
	return nil
}
