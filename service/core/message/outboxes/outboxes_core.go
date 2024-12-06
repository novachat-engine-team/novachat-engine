/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package outboxes

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/snowflake"
	"novachat_engine/pkg/util"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message"
	messageCore "novachat_engine/service/core/message/message"
	data_message "novachat_engine/service/data/messages/message"
	"novachat_engine/service/mgo"
	"strings"
)

func init() {
	mgo.DBE.Enable64ToString = false
}

type Core struct {
	messageCore *messageCore.Core
}

func NewOutboxesCore(config *config.MongodbConfig, messageIdNode *snowflake.Node) *Core {
	mgo.InstallMongodb(config)

	return &Core{
		messageCore: messageCore.NewMessageCore(config, messageIdNode),
	}
}

func (c *Core) migratorTableMessage(nameList []string) {
	var err error
	for idx := 0; idx < int(math.Min(float64(message.GetTableLimit()), message.MaxTableLimit)); idx++ {

		if util.IndexStrings(&nameList, message.TableName(int64(idx), message.TableMessage)) >= 0 {
			continue
		}
		err = mgo.GetMongoDB().Database(message.DBMessage).CreateCollection(context.Background(), message.TableName(int64(idx), message.TableMessage))
		if err != nil {
			log.Fatalf("OutboxesCore %s Create tableName:%s error:%s", message.DBMessage, message.TableName(int64(idx), message.TableMessage), err.Error())
			panic(err.Error())
		}

		err = mgo.GetMongoDB().CreateManyIndexes(message.DBMessage, message.TableName(int64(idx), message.TableMessage),
			[][]*mgo.IndexKey{{
				{Key: "user_id", Order: 1, Unique: false},
				{Key: "peer_id", Order: 1, Unique: false},
			}, {
				{Key: "user_id", Order: 1, Unique: false},
				{Key: "message_id", Order: 1, Unique: false},
			},
			}, true)
		if err != nil {
			log.Fatalf("OutboxesCore Create Index DBName:%s TableName:%s", message.DBMessage, message.TableName(int64(idx), message.TableMessage), err.Error())
			panic(err.Error())
		}
	}
}

func (c *Core) migratorTableConversation(nameList []string) {
	var err error
	for idx := 0; idx < int(math.Min(float64(message.GetTableLimit()), message.MaxTableLimit)); idx++ {

		if util.IndexStrings(&nameList, message.TableName(int64(idx), message.TableConversation)) >= 0 {
			continue
		}
		err = mgo.GetMongoDB().Database(message.DBMessage).CreateCollection(context.Background(), message.TableName(int64(idx), message.TableConversation))
		if err != nil {
			log.Fatalf("OutboxesCore %s Create tableName:%s error:%s", message.DBMessage, message.TableName(int64(idx), message.TableConversation), err.Error())
			panic(err.Error())
		}

		err = mgo.GetMongoDB().CreateManyIndexes(message.DBMessage, message.TableName(int64(idx), message.TableConversation),
			[][]*mgo.IndexKey{{
				{Key: "user_id", Order: 1, Unique: false},
				{Key: "peer_id", Order: 1, Unique: false},
			},
			}, true)
		if err != nil {
			log.Fatalf("OutboxesCore Create Index DBName:%s TableName:%s", message.DBMessage, message.TableName(int64(idx), message.TableConversation), err.Error())
			panic(err.Error())
		}
	}
}

func (c *Core) Migrator() {
	allOpts := options.ListDatabases().
		SetNameOnly(true).
		SetAuthorizedDatabases(true)
	nameList, err := mgo.GetMongoDB().ListDatabaseNames(context.Background(), bson.M{}, allOpts)
	if err != nil {
		log.Fatalf("OutboxesCore error:%s", err.Error())
		panic(err)
	}

	found := false
	for _, name := range nameList {
		if strings.Compare(name, message.DBMessage) == 0 {
			found = true
			break
		}
	}
	if found { // collection
		found = false
		op := options.ListCollections()
		op.SetNameOnly(true)
		nameList, err = mgo.GetMongoDB().Database(message.DBMessage).ListCollectionNames(context.Background(), bson.M{}, op)
		if err != nil {
			log.Fatalf("OutboxesCore table error:%s", err.Error())
			panic(err)
		}
	}

	c.migratorTableMessage(nameList)
	c.migratorTableConversation(nameList)

	if util.IndexStrings(&nameList, message.TableMessageBox) < 0 {
		err = mgo.GetMongoDB().Database(message.DBMessage).CreateCollection(context.Background(), message.TableMessageBox)
		if err != nil {
			log.Fatalf("OutboxesCore %s Create tableName:%s error:%s", message.DBMessage, message.TableMessageBox, err.Error())
			panic(err.Error())
		}
	}
}

func (c *Core) PinnedMessage(userId int64, peerId int64, peerType constants.PeerType, msgId int32, unpin bool, side bool, globalMessageId int64) (int32, int64, error) {
	if globalMessageId != 0 {
		msgList, err := c.messageCore.GetMessageByGlobalMessageIdList(userId, []int64{globalMessageId}, false)
		if err != nil {
			log.Errorf("PinnedMessage GetMessageByGlobalMessageIdList error:%s", err.Error())
			return -1, 0, err
		}
		if len(msgList) == 0 {
			log.Warnf("PinnedMessage GetMessageByGlobalMessageIdList empty userId:%d globalMessageId:%d", userId, globalMessageId)
			return 0, 0, nil
		}
		msgId = msgList[0].Id
	} else {
		if msgId != 0 {
			msgList, err := c.messageCore.GetMessageList(userId, []int32{msgId})
			if err != nil {
				log.Errorf("PinnedMessage GetMessageList error:%s", err.Error())
				return -1, 0, err
			}
			if len(msgList) == 0 {
				log.Warnf("PinnedMessage GetMessageList empty userId:%d msgId:%d", userId, msgId)
				return 0, 0, nil
			}
			globalMessageId = msgList[0].GlobalMessageId
		}
	}

	if msgId == 0 {
		log.Warnf("PinnedMessage msgId == 0", userId)
		return 0, 0, nil
	}

	conversation, err := c.messageCore.GetConversation(userId, &data_message.Conversation{
		PeerId:   peerId,
		PeerType: peerType.ToInt32()})
	if err != nil {
		log.Errorf("PinnedMessage error:%s", err.Error())
		return -1, 0, err
	}

	var pts int32
	if idx := util.Index(conversation.PinIds, func(idx int) bool {
		return conversation.PinIds[idx].Id == msgId
	}); idx >= 0 {
		if unpin {
			copy(conversation.PinIds, conversation.PinIds[:idx])
			copy(conversation.PinIds[idx:], conversation.PinIds[idx+1:])
			conversation.PinIds = conversation.PinIds[:len(conversation.PinIds)-1]
			pts, err = c.messageCore.PinnedMessage(userId, peerId, peerType, conversation.PinIds, msgId, unpin)
		} else {
			if side == false && !conversation.PinIds[idx].Full && conversation.PinIds[idx].OwnerUserId == userId {
				conversation.PinIds[idx].Full = side == false
				pts, err = c.messageCore.PinnedMessage(userId, peerId, peerType, conversation.PinIds, msgId, unpin)
			}
		}
	} else {
		if unpin {
			return 0, 0, nil
		}
		pts, err = c.messageCore.PinnedMessage(userId, peerId, peerType, []*data_message.PinnedMessage{{
			Id:          msgId,
			OwnerUserId: userId,
			Full:        side == false,
		}}, msgId, unpin)
	}
	if err != nil {
		log.Errorf("PinnedMessage Core PinnedMessage error:%s", err.Error())
		return -1, 0, err
	}

	log.Debugf("PinnedMessage userId:%d msgId:%d pts:%d", userId, msgId, pts)
	return pts, globalMessageId, nil
}

func (c *Core) PinnedChannelMessage(userId int64, peerId int64, peerType constants.PeerType, msgId int32, unpin bool, side bool) error {
	if msgId == 0 {
		log.Warnf("PinnedChannelMessage msgId == 0", userId)
		return nil
	}

	conversation, err := c.messageCore.GetConversation(userId, &data_message.Conversation{
		PeerId:   peerId,
		PeerType: peerType.ToInt32()})
	if err != nil {
		log.Errorf("PinnedChannelMessage error:%s", err.Error())
		return err
	}

	var pts int32
	if idx := util.Index(conversation.PinIds, func(idx int) bool {
		return conversation.PinIds[idx].Id == msgId
	}); idx >= 0 {
		if unpin {
			copy(conversation.PinIds, conversation.PinIds[:idx])
			copy(conversation.PinIds[idx:], conversation.PinIds[idx+1:])
			conversation.PinIds = conversation.PinIds[:len(conversation.PinIds)-1]
			err = c.messageCore.ConversationPinnedMessage(context.Background(), userId, peerId, peerType, conversation.PinIds)
		} else {
			if side == false && !conversation.PinIds[idx].Full && conversation.PinIds[idx].OwnerUserId == userId {
				conversation.PinIds[idx].Full = side == false
				err = c.messageCore.ConversationPinnedMessage(context.Background(), userId, peerId, peerType, conversation.PinIds)
			}
		}
	} else {
		if unpin {
			return nil
		}
		err = c.messageCore.ConversationPinnedMessage(context.Background(), userId, peerId, peerType, []*data_message.PinnedMessage{{
			Id:          msgId,
			OwnerUserId: userId,
			Full:        side == false,
		}})
	}
	if err != nil {
		log.Errorf("PinnedChannelMessage Core PinnedMessage error:%s", err.Error())
		return err
	}

	log.Debugf("PinnedChannelMessage userId:%d msgId:%d pts:%d:%s", userId, msgId, pts)
	return nil
}
