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
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/id/message_id_pts"
	"novachat_engine/service/core/message"
	"novachat_engine/service/data/messages/id/message_id_pts"
	"novachat_engine/service/data/messages/message"
	"novachat_engine/service/mgo"
)

func (c *Core) readHistory(
	userId int64,
	peerId int64,
	peerType constants.PeerType,
	maxId int32,
	date int32,
	boxType message.BoxType) (int32, int32, error) {

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	st := options.Transaction()
	st.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	st.SetReadConcern(readconcern.Majority())

	var folderId int32
	var id *data_id.Id
	var pts int32
	var err error
	err = mgo.GetMongoDB().UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
		sessionContext.StartTransaction(st)
		peerId = message.MakePeerId(peerId, peerType)
		conversation := data_message.Conversation{
			Id:     message.MakeDialogId(userId, peerId, peerType),
			UserId: userId,
			PeerId: peerId,
			Date:   date,
		}
		sr := mgo.GetMongoDB().Database(message.DBMessage).
			Collection(message.TableName(userId, message.TableConversation), op).
			FindOne(sessionContext, mgo.DBE.MarshalCustomSpecMap(conversation, "UserId", "PeerId"))
		if sr.Err() != nil {
			sessionContext.AbortTransaction(sessionContext)
			err = sr.Err()
			log.Errorf("readHistory userId:%d peerId:%d peerType:%v error:%s", userId, peerId, peerType, sr.Err())
			return sr.Err()
		}

		err = sr.Decode(&conversation)
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			log.Errorf("readHistory userId:%d peerId:%d peerType:%v Decode error:%s", userId, peerId, peerType, err.Error())
			return err
		}

		log.Debugf("readHistory userId:%d peerId:%d maxId:%d boxType:%+v conversation:%+v", userId, peerId, maxId, boxType, conversation)
		var update bson.M
		folderId = conversation.FolderId
		if boxType == message.InboxType {
			if conversation.Top < maxId {
				sessionContext.AbortTransaction(sessionContext)
				return nil
			}
			if conversation.Top == maxId {
				if conversation.InboxMaxId <= maxId {
					sessionContext.AbortTransaction(sessionContext)
					return nil
				}
			}
			conversation.InboxMaxId = maxId
			filter := bson.M{"user_id": userId, "peer_id": peerId, "deleted": false, "id": bson.M{mgo.GT: maxId}}
			count, err1 := mgo.GetMongoDB().Database(message.DBMessage).
				Collection(message.TableName(userId, message.TableMessage), op).
				CountDocuments(sessionContext, filter)
			if err1 != nil && err1 != mongo.ErrNoDocuments {
				sessionContext.AbortTransaction(sessionContext)
				return err1
			}
			conversation.UnreadCount = int32(count)
		} else {
			if conversation.Top < maxId {
				sessionContext.AbortTransaction(sessionContext)
				return nil
			}
			if conversation.Top == maxId {
				if conversation.OutboxMaxId <= maxId {
					sessionContext.AbortTransaction(sessionContext)
					return nil
				}
			}
			conversation.OutboxMaxId = maxId
		}

		id, err = c.id.NextId(userId, message_id_pts.WithNextPts(1))
		if err != nil {
			log.Errorf("readHistory NextId true error:%s", err.Error())
			return err
		}
		pts = id.Pts
		conversation.Pts = pts
		if boxType == message.InboxType {
			update = mgo.DBE.MarshalCustomSpecMap(conversation, "Date", "Pts", "InboxMaxId", "UnreadCount")
		} else {
			update = mgo.DBE.MarshalCustomSpecMap(conversation, "Date", "Pts", "OutboxMaxId", "UnreadCount")
		}
		sr = mgo.GetMongoDB().Database(message.DBMessage).
			Collection(message.TableName(userId, message.TableConversation), op).
			FindOneAndUpdate(sessionContext, mgo.DBE.MarshalCustomSpecMap(conversation, "UserId", "PeerId"),
				bson.M{mgo.SET: update})
		if sr.Err() != nil {
			sessionContext.AbortTransaction(sessionContext)
			err = sr.Err()
			log.Errorf("readHistory userId:%d peerId:%d peerType:%v error:%s", userId, peerId, peerType, sr.Err())
			return sr.Err()
		}
		sessionContext.CommitTransaction(sessionContext)
		return nil
	})
	if err != nil {
		log.Errorf("readHistory userId:%d peerId:%d peerType:%v error:%s", userId, peerId, peerType, err)
		return 0, 0, err
	}

	log.Infof("readHistory userId:%d peerId:%d peerType:%v pts:%d maxId:%d date:%d", userId, peerId, peerType, pts, maxId, date)
	return pts, folderId, nil
}

func (c *Core) ReadHistory(
	userId int64,
	peerId int64,
	peerType constants.PeerType,
	maxId int32,
	date int32,
	boxType message.BoxType) (int32, int32, error) {
	switch peerType {
	case constants.PeerTypeUser, constants.PeerTypeSelf, constants.PeerTypeChat, constants.PeerTypeChannel:
		return c.readHistory(userId, peerId, peerType, maxId, date, boxType)
	default:
		panic(fmt.Sprintf("ReadHistory peerType:%v error", peerType))
	}
}
