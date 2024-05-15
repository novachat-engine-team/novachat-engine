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
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message"
	"novachat_engine/service/mgo"
)

func (c *Core) SendMessages(context context.Context, userId int64, peerId int64, peerType constants.PeerType, list []*msgService.SendMessageData, outboxType message.BoxType, globalMessageIdList *[]int64) error {
	if len(list) == 0 {
		log.Warnf("SendMessages userId:%df peerId:%d peerType:%v outboxType:%v len() == 0", userId, peerId, peerType, outboxType)
		return nil
	}
	log.Debugf("SendMessages userId:%df peerId:%d peerType:%v outboxType:%v len():%d", userId, peerId, peerType, outboxType, len(list))
	if peerType == constants.PeerTypeUser {
		if outboxType == message.InboxType {
			userId, peerId = peerId, userId
		}
	}
	document := make([]interface{}, 0, len(list))
	for idx, v := range list {

		dataMessage := ToMessageData(userId, constants.PeerTypeFromUserIDType32(v.Message.FromId90DDDC1171).ToInt(), peerId, peerType, v.Message)
		dataMessage.Did = c.messageIdNode.Generate().Int64()
		dataMessage.Out = outboxType == message.OutboxType

		//TODO: global message id
		if peerType == constants.PeerTypeUser {
			if outboxType == message.InboxType {
				dataMessage.GlobalMessageId = (*globalMessageIdList)[idx]
			} else {
				dataMessage.GlobalMessageId = dataMessage.Did
				*globalMessageIdList = append(*globalMessageIdList, dataMessage.GlobalMessageId)
			}
		}

		document = append(document, mgo.DBE.MarshalCustomSpecMap(dataMessage))
	}

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())
	imr, err := mgo.GetDatabase(message.DBMessage).
		Collection(message.TableName(userId, message.TableMessage), op).
		InsertMany(context, document)
	if err != nil {
		log.Errorf("SendMessages userId:%df peerId:%d peerType:%v outboxType:%v len():%d", userId, peerId, peerType, outboxType, len(list))
		return err
	}

	log.Infof("SendMessages userId:%df peerId:%d peerType:%v outboxType:%v len():%d InsertID:%+v", userId, peerId, peerType, outboxType, len(list), imr.InsertedIDs)
	return nil
}

func (c *Core) SendChannelMessages(context context.Context, userId int64, peerId int64, peerType constants.PeerType, list []*msgService.SendMessageData) error {
	if len(list) == 0 {
		log.Warnf("SendChannelMessages userId:%d peerId:%d peerType:%v len() == 0", userId, peerId, peerType)
		return nil
	}
	log.Debugf("SendChannelMessages userId:%d peerId:%d peerType:%v len():%d", userId, peerId, peerType, len(list))

	document := make([]interface{}, 0, len(list))
	for _, v := range list {
		dataMessage := ToMessageData(userId, constants.PeerTypeFromUserIDType32(v.Message.FromId90DDDC1171).ToInt(), peerId, peerType, v.Message)
		dataMessage.Did = c.messageIdNode.Generate().Int64()
		dataMessage.GlobalMessageId = dataMessage.Did
		document = append(document, mgo.DBE.MarshalCustomSpecMap(dataMessage))
	}

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())
	imr, err := mgo.GetDatabase(message.DBMessage).
		Collection(message.TableName(peerId, message.TableChannelMessage), op).
		InsertMany(context, document)
	if err != nil {
		log.Errorf("SendChannelMessages userId:%d peerId:%d peerType:%v len():%d", userId, peerId, peerType, len(list))
		return err
	}

	log.Infof("SendChannelMessages userId:%d peerId:%d peerType:%v len():%d InsertID:%+v", userId, peerId, peerType, len(list), imr.InsertedIDs)
	return nil
}
