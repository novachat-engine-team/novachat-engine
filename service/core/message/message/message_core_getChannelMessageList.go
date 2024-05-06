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
	data_message "novachat_engine/service/data/messages/message"
	"novachat_engine/service/mgo"
)

type ChannelMessageId struct {
	ChannelId int64
	IdList    []int32
}

func (c *Core) GetChannelMessageList(channelMsgId *ChannelMessageId) ([]*data_message.Message, error) {

	channelMsgId.ChannelId = message.MakePeerId(channelMsgId.ChannelId, constants.PeerTypeChannel)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())
	cursor, err := mgo.GetDatabase(message.DBMessage).
		Collection(message.TableName(channelMsgId.ChannelId, message.TableChannelMessage), op).
		Find(context.Background(), bson.M{"peer_id": channelMsgId.ChannelId, "id": bson.M{mgo.IN: channelMsgId.IdList}})
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("GetChannelMessageList channelId:%d GetMessageList error:%s", channelMsgId.ChannelId, err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNoDocuments {
		log.Warnf("GetChannelMessageList channelId:%d GetMessageList:%v empty", channelMsgId.ChannelId, channelMsgId.IdList)
		return []*data_message.Message{}, nil
	}

	var messageDatas []*data_message.Message
	err = cursor.All(context.Background(), &messageDatas)
	if err != nil {
		log.Warnf("GetChannelMessageList channelId:%d Decoder error:%s", channelMsgId.ChannelId, err.Error())
	}

	log.Infof("GetChannelMessageList channelId:%d messageIdList:%d len()=%d", channelMsgId.ChannelId, channelMsgId.IdList, len(messageDatas))
	return messageDatas, nil
}

func (c *Core) GetChannelMessageListByMaxId(peerId int64, maxId int32) ([]*data_message.Message, error) {

	peerId = message.MakePeerId(peerId, constants.PeerTypeChannel)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())
	cursor, err := mgo.GetDatabase(message.DBMessage).
		Collection(message.TableName(peerId, message.TableChannelMessage), op).
		Find(context.Background(), bson.M{"peer_id": peerId, "id": bson.M{mgo.LTE: maxId}, "deleted": false})
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("GetChannelMessageListByMaxId channelId:%d error:%s", peerId, err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNoDocuments {
		log.Warnf("GetChannelMessageListByMaxId channelId:%d maxId:%v empty", peerId, maxId)
		return []*data_message.Message{}, nil
	}

	var messageDatas []*data_message.Message
	err = cursor.All(context.Background(), &messageDatas)
	if err != nil {
		log.Warnf("GetChannelMessageListByMaxId channelId:%d Decoder error:%s", peerId, err.Error())
	}

	log.Infof("GetChannelMessageListByMaxId channelId:%d maxId:%d len()=%d", peerId, maxId, len(messageDatas))
	return messageDatas, nil
}

func (c *Core) GetChannelMessageListByUserId(peerId int64, userId int64) ([]*data_message.Message, error) {

	peerId = message.MakePeerId(peerId, constants.PeerTypeChannel)

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())
	cursor, err := mgo.GetDatabase(message.DBMessage).
		Collection(message.TableName(peerId, message.TableChannelMessage), op).
		Find(context.Background(), bson.M{"peer_id": peerId, "user_id": userId, "deleted": false})
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("GetChannelMessageListByUserId channelId:%d error:%s", peerId, err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNoDocuments {
		log.Warnf("GetChannelMessageListByUserId channelId:%d userId:%v empty", peerId, userId)
		return []*data_message.Message{}, nil
	}

	var messageDatas []*data_message.Message
	err = cursor.All(context.Background(), &messageDatas)
	if err != nil {
		log.Warnf("GetChannelMessageListByUserId channelId:%d Decoder error:%s", peerId, err.Error())
	}

	log.Infof("GetChannelMessageListByUserId channelId:%d userId:%d len()=%d", peerId, userId, len(messageDatas))
	return messageDatas, nil
}
