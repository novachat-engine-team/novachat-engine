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
