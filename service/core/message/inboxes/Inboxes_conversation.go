/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package inboxes

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message"
	data_message "novachat_engine/service/data/messages/message"
	"novachat_engine/service/mgo"
)

func (c *Core) updateInboxConversation(context context.Context, userId int64, peerId int64, peerType constants.PeerType, list []*msgService.SendMessageData, pts int32) error {
	peerId = message.MakePeerId(peerId, peerType)
	date := list[0].Message.Date
	inboxMaxId := list[0].Message.Id
	for _, v := range list {
		if v.Message.Id > inboxMaxId {
			inboxMaxId = v.Message.Id
			date = v.Message.Date
		}
	}
	if peerType == constants.PeerTypeUser {
		userId, peerId = peerId, userId
	}

	conversation := data_message.Conversation{
		Id:          message.MakeDialogId(userId, peerId, peerType),
		UserId:      userId,
		PeerId:      peerId,
		Date:        date,
		UnreadCount: 0,
		Top:         inboxMaxId,
		InboxMaxId:  inboxMaxId,
		PeerType:    peerType.ToInt32(),
		Pts:         pts,
	}

	filter := mgo.DBE.MarshalCustomSpecMap(conversation, "UserId", "PeerId")
	update := bson.M{mgo.SET: mgo.DBE.MarshalCustomSpecMap(conversation, "Date", "Pts", "Top", "InboxMaxId"),
		mgo.INC:         bson.M{"unread_count": 1},
		mgo.SETONINSERT: mgo.DBE.MarshalCustomSpecMap(conversation, "Id", "Pinned", "FolderId", "Draft", "PinIds", "PeerType", "OutboxMaxId")}

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())
	opu := options.Update()
	opu.SetUpsert(true)
	ups, err := mgo.GetMongoDB().Database(message.DBMessage).
		Collection(message.TableName(userId, message.TableConversation), op).
		UpdateOne(context, filter, update, opu)
	if err != nil {
		log.Fatalf("updateOutboxConversation UpdateOne error:%s", err.Error())
		return err
	}

	log.Debugf("updateOutboxConversation UpdateOne ups:%+v", ups)
	return nil
}
