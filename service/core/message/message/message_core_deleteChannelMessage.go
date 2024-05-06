package message

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/id/message_id_pts"
	"novachat_engine/service/core/message"
	"novachat_engine/service/mgo"
)

func (c *Core) DeleteChannelMessages(chatId int64, userId int64, messageIdList []int32) (int32, error) {

	ptsPeerId := message.MakePeerId(chatId, constants.PeerTypeChannel)
	id, err := c.id.NextId(ptsPeerId, message_id_pts.WithNextPts(1))
	if err != nil {
		log.Errorf("DeleteChannelMessages error:%s", err.Error())
		return 0, err
	}
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())

	ur, err := mgo.GetDatabase(message.DBMessage).
		Collection(message.TableName(chatId, message.TableChannelMessage), op).
		UpdateMany(context.Background(), bson.M{"peer_id": ptsPeerId, "id": bson.M{mgo.IN: messageIdList}}, bson.M{"deleted": true})
	if err != nil {
		log.Errorf("DeleteChannelMessages error:%s", err.Error())
		return 0, err
	}

	log.Debugf("DeleteChannelMessages ur:%+v", ur)
	return id.Pts, nil
}
