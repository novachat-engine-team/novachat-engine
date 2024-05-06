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
	"novachat_engine/service/core/id/message_id_pts"
	"novachat_engine/service/core/message"
	data_id "novachat_engine/service/data/messages/id/message_id_pts"
	data_message "novachat_engine/service/data/messages/message"
	"novachat_engine/service/mgo"
)

func (c *Core) DeleteMessages(userId int64, idList []int32) (int32, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	id, err := c.id.NextId(userId, message_id_pts.WithNextPts(1))
	if err != nil {
		log.Errorf("DeleteMessages NextId error:%s", err.Error())
		return 0, err
	}

	ur, err := mgo.GetDatabase(message.DBMessage).
		Collection(message.TableName(userId, message.TableMessage), op).
		UpdateMany(context.Background(), bson.M{"user_id": userId, "id": bson.M{mgo.IN: idList}}, bson.M{"deleted": true})
	if err != nil {
		log.Errorf("DeleteMessages error:%s", err.Error())
		return 0, err
	}

	log.Debugf("DeleteMessages ur:%+v", ur)
	return id.MessageId, nil
}

func (c *Core) DeleteMessagesByMaxId(userId, peerId int64, peerType constants.PeerType, maxId int32, now int32, pts int32) (int32, []int32, error) {
	conversation, err := c.GetConversation(userId, &data_message.Conversation{
		UserId:   userId,
		PeerId:   peerId,
		PeerType: peerType.ToInt32(),
	})
	if err != nil {
		log.Errorf("DeleteMessagesByMaxId Conversation error:%s", err.Error())
		return 0, nil, err
	}
	if conversation.ClearMaxId >= maxId {
		return conversation.Pts, nil, nil
	}

	var idData *data_id.Id
	if !(peerType == constants.PeerTypeChannel ||
		peerType == constants.PeerTypeChat) {
		idData, err = c.id.NextId(userId, message_id_pts.WithNextPts(1))
		if err != nil {
			log.Errorf("DeleteMessagesByMaxId NextId error:%s", err.Error())
			return 0, nil, err
		}
	} else {
		idData = &data_id.Id{Pts: pts}
	}

	peerId = message.MakePeerId(peerId, peerType)
	filter := bson.M{}
	filter["peer_id"] = peerId
	filter["user_id"] = userId
	filter["id"] = bson.M{mgo.LTE: maxId}
	filter["deleted"] = false

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	var messageIdList []int32
	if !(peerType == constants.PeerTypeChannel ||
		peerType == constants.PeerTypeChat) {
		cursor, err := mgo.GetDatabase(message.DBMessage).
			Collection(message.TableName(userId, message.TableMessage), op).
			Find(context.TODO(), filter)
		if err != nil {
			log.Errorf("DeleteMessagesByMaxId Find error:%s", err.Error())
			return 0, nil, err
		}

		var dataMessageList []*data_message.Message
		err = cursor.All(context.TODO(), &dataMessageList)
		cursor.Close(context.TODO())
		if err != nil {
			log.Errorf("DeleteMessagesByMaxId Decode Find error:%s", err.Error())
			return 0, nil, err
		}
		messageIdList = make([]int32, 0, len(dataMessageList))
		for _, v := range dataMessageList {
			messageIdList = append(messageIdList, v.Id)
		}
	}

	ots := options.Transaction()
	ots.SetReadConcern(readconcern.Majority())
	ots.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))

	err = mgo.GetMongoDB().UseSession(context.TODO(), func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction(ots)
		if err != nil {
			log.Errorf("DeleteMessagesByMaxId StartTransaction error:%s", err.Error())
			return err
		}

		if !(peerType == constants.PeerTypeChannel ||
			peerType == constants.PeerTypeChat) {
			var ur *mongo.UpdateResult
			ur, err = mgo.GetDatabase(message.DBMessage).
				Collection(message.TableName(userId, message.TableMessage), op).
				UpdateMany(sessionContext, filter, bson.M{"deleted": true})
			if err != nil {
				sessionContext.AbortTransaction(sessionContext)
				log.Errorf("DeleteMessagesByMaxId UpdateMany error:%s", err.Error())
				return err
			}
			_ = ur
		}

		conversation.UnreadCount = 0
		if conversation.InboxMaxId <= maxId {
			conversation.InboxMaxId = maxId
		}
		conversation.Date = now
		conversation.Draft = ""
		conversation.Pts = idData.Pts
		conversation.ClearMaxId = maxId

		_, err = mgo.GetMongoDB().Database(message.DBMessage).
			Collection(message.TableName(userId, message.TableConversation), op).
			UpdateOne(sessionContext, mgo.DBE.MarshalCustomSpecMap(conversation, "UserId", "PeerId"), mgo.DBE.MarshalCustomSpecMap(conversation, "UnreadCount", "InboxMaxId", "Date", "Draft", "Pts"))
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			log.Errorf("DeleteMessagesByMaxId TableConversation error:%s", err.Error())
			return err
		}

		sessionContext.CommitTransaction(sessionContext)
		return nil
	})

	if err != nil {
		log.Fatalf("DeleteMessagesByMaxId error:%s", err.Error())
		return 0, nil, err
	}

	return idData.Pts, messageIdList, nil
}

func (c *Core) DeleteMessagesIds(userId, peerId int64, peerType constants.PeerType, idList []int32, now int32, pts int32) (int32, error) {
	conversation, err := c.GetConversation(userId, &data_message.Conversation{
		UserId:   userId,
		PeerId:   peerId,
		PeerType: peerType.ToInt32(),
	})
	if err != nil {
		log.Errorf("DeleteMessagesIds Conversation error:%s", err.Error())
		return 0, err
	}

	var idData *data_id.Id
	if !(peerType == constants.PeerTypeChannel ||
		peerType == constants.PeerTypeChat) {
		idData, err = c.id.NextId(userId, message_id_pts.WithNextPts(1))
		if err != nil {
			log.Errorf("DeleteMessagesIds NextId error:%s", err.Error())
			return 0, err
		}
	} else {
		idData = &data_id.Id{Pts: pts}
	}

	peerId = message.MakePeerId(peerId, peerType)
	filter := bson.M{}
	filter["peer_id"] = peerId
	filter["user_id"] = userId
	filter["id"] = bson.M{mgo.IN: idList}
	filter["deleted"] = false

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	ots := options.Transaction()
	ots.SetReadConcern(readconcern.Majority())
	ots.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))

	err = mgo.GetMongoDB().UseSession(context.TODO(), func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction(ots)
		if err != nil {
			log.Errorf("DeleteMessagesIds StartTransaction error:%s", err.Error())
			return err
		}

		if !(peerType == constants.PeerTypeChannel ||
			peerType == constants.PeerTypeChat) {
			var ur *mongo.UpdateResult
			ur, err = mgo.GetDatabase(message.DBMessage).
				Collection(message.TableName(userId, message.TableMessage), op).
				UpdateMany(sessionContext, filter, bson.M{"deleted": true})
			if err != nil {
				sessionContext.AbortTransaction(sessionContext)
				log.Errorf("DeleteMessagesIds UpdateMany error:%s", err.Error())
				return err
			}
			_ = ur
		}

		maxId := idList[0]
		for _, vv := range idList {
			if conversation.InboxMaxId < vv {
				conversation.UnreadCount--
			}
			if maxId < vv {
				maxId = vv
			}
		}
		conversation.InboxMaxId = maxId
		conversation.Date = now
		conversation.Pts = idData.Pts

		_, err = mgo.GetMongoDB().Database(message.DBMessage).
			Collection(message.TableName(userId, message.TableConversation), op).
			UpdateOne(sessionContext, mgo.DBE.MarshalCustomSpecMap(conversation, "UserId", "PeerId"), mgo.DBE.MarshalCustomSpecMap(conversation, "UnreadCount", "InboxMaxId", "Date", "Draft", "Pts"))
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			log.Errorf("DeleteMessagesIds TableConversation error:%s", err.Error())
			return err
		}

		sessionContext.CommitTransaction(sessionContext)
		return nil
	})

	if err != nil {
		log.Fatalf("DeleteMessagesIds error:%s", err.Error())
		return 0, err
	}

	return idData.Pts, nil
}
