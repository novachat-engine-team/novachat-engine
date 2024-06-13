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
	data_message "novachat_engine/service/data/messages/message"
	"novachat_engine/service/mgo"
)

func (c *Core) PinnedMessage(
	userId int64,
	peerId int64,
	peerType constants.PeerType,
	messages []*data_message.PinnedMessage,
	msgId int32,
	unpin bool) (int32, error) {
	id, err := c.id.NextId(userId, message_id_pts.WithNextPts(1))
	if err != nil {
		log.Errorf("PinnedMessage NextId error:%s", err.Error())
		return 0, err
	}

	pts := id.Pts
	err = mgo.GetMongoDB().UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
		sessionContext.StartTransaction()

		err = c.ConversationPinnedMessage(sessionContext, userId, peerId, peerType, messages)
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			log.Errorf("PinnedMessage ConversationPinnedMessage error:%s", err.Error())
			return err
		}

		op := options.Collection()
		op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
		op.SetReadConcern(readconcern.Majority())
		_, err = mgo.GetDatabase(message.DBMessage).
			Collection(message.TableName(userId, message.TableMessage), op).
			UpdateOne(sessionContext, bson.M{"user_id": userId, "id": msgId}, bson.M{mgo.SET: bson.M{"pinned": !unpin}})
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			log.Errorf("PinnedMessage UpdateOne error:%s", err.Error())
			return err
		}
		sessionContext.CommitTransaction(sessionContext)
		return nil
	})

	if err != nil {
		log.Errorf("PinnedMessage userId:%d peerId:%d peerType:%d error:%s", userId, peerId, peerType, err.Error())
		return 0, err
	}
	return pts, err
}
