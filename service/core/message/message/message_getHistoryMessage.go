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
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message"
	"novachat_engine/service/data/messages/message"
	"novachat_engine/service/mgo"
	"sort"
)

func (c *Core) getHistoryMessages(userId int64, peerId int64, peerType constants.PeerType, messageId int32, limit int32, minId int32, date int32, deleted bool) ([]*data_message.Message, error) {
	filter := bson.M{}
	sortOperator := bson.M{}
	if limit < 0 { // backward
		limit = -limit
		limit += 1

		if messageId != minId {
			if uint32(messageId) <= uint32(minId) {
				return []*data_message.Message{}, nil
			}
			// <
			filter["id"] = bson.M{mgo.LT: messageId, mgo.GT: minId}
		} else {
			if messageId == minId && minId == 0 {
				filter["id"] = bson.M{mgo.GT: minId}
				limit = -1
			} else {
				filter["id"] = messageId
			}
		}
		sortOperator["id"] = -1

	} else { // forward
		if messageId < minId {
			messageId = minId + 1
		}
		// >=
		filter["id"] = bson.M{mgo.GTE: messageId}
		sortOperator["id"] = 1
		limit += 1
	}
	if date > 0 {
		filter["date"] = bson.M{mgo.GTE: date}
	}

	filter["peer_id"] = message.MakePeerId(peerId, peerType)

	if !deleted {
		filter["deleted"] = false
	}

	opf := options.Find()
	if limit > 0 {
		opf.SetLimit(int64(limit))
	}
	opf.SetSort(sortOperator)
	var cursor *mongo.Cursor
	var err error
	switch peerType {
	case constants.PeerTypeChannel, constants.PeerTypeChat:
		cursor, err = mgo.GetMongoDB().Database(message.DBMessage).
			Collection(message.TableName(userId, message.TableChannelMessage)).
			Find(context.Background(), filter, opf)
	default:
		filter["user_id"] = userId
		cursor, err = mgo.GetMongoDB().Database(message.DBMessage).
			Collection(message.TableName(userId, message.TableMessage)).
			Find(context.Background(), filter, opf)
	}
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	defer cursor.Close(context.Background())
	var ret []*data_message.Message
	if err == mongo.ErrNoDocuments {
		cursor.Close(context.Background())
		return ret, nil
	}

	err = cursor.All(context.Background(), &ret)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(ret, func(i, j int) bool {
		if limit < 0 {
			return ret[i].Id > ret[j].Id
		} else {
			return ret[i].Id < ret[j].Id
		}
	})

	list := make([]*data_message.Message, 0, len(ret))
	for idx := range ret {
		if int32(len(ret)) > limit {
			if idx == len(ret)-1 {
				break
			} else {
				list = append(list, ret[idx])
			}
		} else {
			list = append(list, ret[idx])
		}
	}
	//	log.Debugf("getHistoryMessages currentId:%d list:%v", messageId, list)
	return list, nil
}

func (c *Core) GetHistoryMessages(
	userId int64,
	peerId int64,
	peerType constants.PeerType,
	messageId,
	limit,
	minId,
	date int32, deleted bool) ([]*data_message.Message, error) {

	var err error
	var messageDataList []*data_message.Message
	switch peerType {
	case constants.PeerTypeUser, constants.PeerTypeSelf, constants.PeerTypeChat, constants.PeerTypeChannel:
		messageDataList, err = c.getHistoryMessages(userId, peerId, peerType, messageId, limit, minId, date, deleted)
	default:
		panic(fmt.Sprintf("GetHistoryMessage peerType:%v error", peerType))
	}

	if err != nil {
		log.Errorf("GetHistoryMessages userId:%d peerId:%v, peerType:%v messageId:%d, limit:%d error:%s", userId, peerId, peerType, messageId, limit, err.Error())
		return nil, err
	}

	log.Infof("GetHistoryMessages userId:%d peerId:%v, peerType:%v messageId:%d, limit:%d len()=%d", userId, peerId, peerType, messageId, limit, len(messageDataList))
	//log.Debugf("messageList:%+v", messageDataList)
	return messageDataList, nil
}
