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

func (c *Core) GetMessageList(userId int64, idList []int32) ([]*data_message.Message, error) {

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())
	cursor, err := mgo.GetDatabase(message.DBMessage).
		Collection(message.TableName(userId, message.TableMessage), op).
		Find(context.Background(), bson.M{"user_id": userId, "id": bson.M{mgo.IN: idList}})
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("GetMessageList userId:%d GetMessageList error:%s", userId, err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNoDocuments {
		log.Warnf("GetMessageList userId:%d GetMessageList:%v empty", userId, idList)
		return []*data_message.Message{}, nil
	}

	var messageDatas []*data_message.Message
	err = cursor.All(context.Background(), &messageDatas)
	if err != nil {
		log.Warnf("GetMessageList userId:%d Decoder error:%s", userId, err.Error())
	}

	log.Infof("GetMessageList userId:%d messageIdList:%d len()=%d", userId, idList, len(messageDatas))
	return messageDatas, nil
}

func (c *Core) GetMessageByGlobalMessageIdList(userId int64, globalMessageIdList []int64, full bool) ([]*data_message.Message, error) {
	filter := bson.M{}
	if full {
		messageList, err := c.GetMessageByGlobalMessageIdList(userId, globalMessageIdList, false)
		if err != nil {
			log.Errorf("GetMessageByGlobalMessageIdList userId:%d full false error:%s", userId, err.Error())
			return nil, err
		}

		userIdList := make([]int64, 0, len(messageList))
		for _, v := range messageList {
			peerId, peerType := message.MakePeerType(v.PeerId)
			if peerType == constants.PeerTypeUser && peerId != userId {
				userIdList = append(userIdList, peerId)
			}
		}
		if len(userIdList) == 0 {
			log.Warnf("GetMessageByGlobalMessageIdList userId:%d self", userId)
			return messageList, nil
		}
		log.Debugf("GetMessageByGlobalMessageIdList userId:%d userIdList:%+v ", userId, userIdList)
		userIdList = append(userIdList, userId)
		filter["user_id"] = bson.M{mgo.IN: userIdList}
	} else {
		filter["user_id"] = userId
	}
	filter["global_message_id"] = bson.M{mgo.IN: globalMessageIdList}
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())
	cursor, err := mgo.GetDatabase(message.DBMessage).
		Collection(message.TableName(userId, message.TableMessage), op).
		Find(context.Background(), filter)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Errorf("GetMessageByGlobalMessageIdList userId:%d globalMessageIdList error:%s", userId, err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNoDocuments {
		log.Warnf("GetMessageByGlobalMessageIdList userId:%d globalMessageIdList:%v empty", userId, globalMessageIdList)
		return []*data_message.Message{}, nil
	}

	var messageDatas []*data_message.Message
	err = cursor.All(context.Background(), &messageDatas)
	if err != nil {
		log.Warnf("GetMessageByGlobalMessageIdList userId:%d Decoder error:%s", userId, err.Error())
	}

	log.Infof("GetMessageByGlobalMessageIdList userId:%d globalMessageIdList:%d len()=%d", userId, globalMessageIdList, len(messageDatas))
	return messageDatas, nil
}
