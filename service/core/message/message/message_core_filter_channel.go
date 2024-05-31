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
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/common/hash"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/id/message_id_pts"
	"novachat_engine/service/core/message"
	"novachat_engine/service/data/messages/id/message_id_pts"
	"novachat_engine/service/data/messages/message"
	"novachat_engine/service/mgo"
)

func (c *Core) FillChannelMessageBox(ctx context.Context, userId int64, peerId int64, list []*msgService.SendMessageData) ([]int32, bool, error) {
	var id *data_id.Id
	var ptsList []int32
	var err error

	randIdList := make([]int64, 0, len(list))
	for _, v := range list {
		randIdList = append(randIdList, v.RandomId)
	}

	box, err1 := c.getChannelMessageBoxes(ctx, list[0].Message.GroupedId != 0, userId, peerId, randIdList)
	if err1 != nil {
		log.Fatalf("FillChannelMessageBox getChannelMessageBoxes error:%s", err1.Error())
		return nil, false, err1
	}

	ptsPeerId := message.MakePeerId(peerId, constants.PeerTypeChannel)
	if list[0].Message.GroupedId != 0 {
		if len(box) > 0 {
			ptsList = box[0].Pts
			return ptsList, false, nil
		}

		id, err = c.id.NextId(ptsPeerId, message_id_pts.WithNextMessageId(int32(len(randIdList))), message_id_pts.WithNextPts(int32(len(randIdList))))
		if err != nil {
			log.Errorf("FillChannelMessageBox NextId true error:%s", err.Error())
			return nil, false, err
		}
		var messageIdList = make([]int32, 0, len(list))
		ptsList = make([]int32, 0, len(list))
		for idx := range list {
			messageIdList = append(messageIdList, id.MessageId-int32(len(randIdList)-idx)+1)
			ptsList = append(ptsList, id.Pts-int32(len(randIdList)-idx)+1)
			list[idx].Message.Id = messageIdList[idx]
		}

		err = c.setChannelMessageBoxes(ctx, true, userId, peerId, randIdList, messageIdList, ptsList)
		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				return ptsList, true, nil
			}
			log.Errorf("FillChannelMessageBox setChannelMessageBoxes true error:%s", err.Error())
			return nil, false, err
		}
	} else {
		var messageIdList = make([]int32, len(list))
		ptsList = make([]int32, len(list))
		randomIdList := make([]int64, 0, len(list))
		messageId := int32(0)
		for idx, v := range list {
			messageId = 0
			for _, vv := range box {
				if vv.RandomId == v.RandomId {
					messageId = vv.MessageId[0]
					list[idx].Message.Id = messageId
					ptsList[idx] = vv.Pts[0]
					messageIdList[idx] = messageId
					break
				}
			}

			if messageId == 0 {
				randomIdList = append(randomIdList, v.RandomId)
			}
		}

		if len(randomIdList) > 0 {
			id, err = c.id.NextId(ptsPeerId, message_id_pts.WithNextMessageId(int32(len(randomIdList))), message_id_pts.WithNextPts(int32(len(randomIdList))))
			if err != nil {
				log.Errorf("FillChannelMessageBox NextId false error:%s", err.Error())
				return nil, false, err
			}

			log.Debugf("FillChannelMessageBox ptsPeerId:%d id:%+v", ptsPeerId, id)
			_idx := int32(1)
			tempMessageIdList := make([]int32, len(randomIdList))
			tempPtsList := make([]int32, len(randomIdList))
			for idx, v := range messageIdList {
				if v == 0 {
					tempPtsList[_idx-1] = id.Pts - int32(len(randomIdList)) + _idx
					tempMessageIdList[_idx-1] = id.MessageId - int32(len(randomIdList)) + _idx
					list[idx].Message.Id = tempMessageIdList[_idx-1]
					ptsList[idx] = tempPtsList[_idx-1]
					_idx += 1
				}
			}
			err = c.setChannelMessageBoxes(ctx, false, userId, peerId, randIdList, tempMessageIdList, tempPtsList)
			if err != nil {
				if mongo.IsDuplicateKeyError(err) {
					return tempPtsList, true, nil
				}

				log.Errorf("FillChannelMessageBox setChannelMessageBoxes false error:%s", err.Error())
				return nil, false, err
			}
			messageIdList = tempMessageIdList
		} else {
			//ok
			return nil, true, nil
		}
	}

	return ptsList, false, nil
}

func (c *Core) getChannelMessageBoxes(ctx context.Context, group bool, userId int64, peerId int64, randomIdList []int64) ([]*data_message.MessageBox, error) {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	idList := make([]string, 0, len(randomIdList))
	if group {
		idList = append(idList, fmt.Sprintf("%d%d%d", userId, hash.Hash64(randomIdList, len(randomIdList)), peerId))
	} else {
		for _, v := range randomIdList {
			idList = append(idList, fmt.Sprintf("%d%d%d", userId, v, peerId))
		}
	}

	col := mgo.GetMongoDB().Database(message.DBMessage).Collection(message.TableChannelMessageBox, op)
	if len(idList) == 1 {
		filter := bson.M{"_id": idList[0]}
		sr := col.FindOne(ctx, filter)
		if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
			log.Fatalf("getChannelMessageBoxes error:%s", sr.Err().Error())
			return nil, sr.Err()
		}
		if sr.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		box := &data_message.MessageBox{}
		err := sr.Decode(&box)
		if err != nil {
			log.Errorf("getChannelMessageBoxes data_message.MessageBox decode error:%s", err.Error())
			return nil, err
		}
		if len(box.Id) == 0 {
			return nil, nil
		}
		return []*data_message.MessageBox{box}, nil
	} else {
		filter := bson.M{"_id": bson.M{mgo.IN: idList}}

		cursor, err := col.Find(ctx, filter)
		if err != nil && err != mongo.ErrNoDocuments {
			log.Fatalf("getChannelMessageBoxes error:%s", err.Error())
			return nil, err
		}
		defer cursor.Close(ctx)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		defer cursor.Close(ctx)
		var ret []*data_message.MessageBox
		err = cursor.All(ctx, &ret)
		if err != nil {
			log.Errorf("getChannelMessageBoxes cdata_message.MessageBox decode error:%s", err.Error())
			return nil, err
		}
		return ret, nil
	}
}

func (c *Core) setChannelMessageBoxes(ctx context.Context, group bool, userId int64, peerId int64, randomIdList []int64, messageList []int32, ptsList []int32) error {
	if ctx == nil {
		ctx = context.Background()
	}

	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	idList := make([]string, 0, len(randomIdList))
	if group {
		idList = append(idList, fmt.Sprintf("%d%d", peerId, hash.Hash64(randomIdList, len(randomIdList))))
	} else {
		for _, v := range randomIdList {
			idList = append(idList, fmt.Sprintf("%d%d", peerId, v))
		}
	}

	if group == true {
		result, err := mgo.GetMongoDB().Database(message.DBMessage).Collection(message.TableChannelMessageBox, op).
			InsertOne(ctx, mgo.DBE.MarshalCustomSpecMap(&data_message.MessageBox{
				Id:        fmt.Sprintf("%d%s", peerId, idList[0]),
				UserId:    userId,
				PeerId:    peerId,
				MessageId: messageList,
				Pts:       ptsList,
			}, "Id", "UserId", "PeerId", "MessageId", "Pts"))

		if err != nil {
			log.Warnf("setChannelMessageBoxes InsertOne error:%s", err.Error())
			return err
		}
		_ = result
	} else {
		update := make([]interface{}, 0, len(randomIdList))
		for idx := range randomIdList {
			update = append(update, mgo.DBE.MarshalCustomSpecMap(&data_message.MessageBox{
				Id:        fmt.Sprintf("%d%s", peerId, idList[idx]),
				UserId:    userId,
				PeerId:    peerId,
				RandomId:  randomIdList[idx],
				MessageId: []int32{messageList[idx]},
				Pts:       []int32{ptsList[idx]},
			}, "Id", "RandomId", "PeerId", "UserId", "MessageId", "Pts"))
		}

		col, err := mgo.GetMongoDB().Database(message.DBMessage).Collection(message.TableChannelMessageBox, op).InsertMany(ctx, update)
		if err != nil {
			log.Warnf("setChannelMessageBoxes InsertMany error:%s", err.Error())
			return err
		}
		_ = col
	}

	return nil
}
