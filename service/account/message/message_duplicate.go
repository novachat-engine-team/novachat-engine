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
	"github.com/gomodule/redigo/redis"
	"github.com/json-iterator/go"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/log"
)

type DuplicateMessage struct {
	MessageId int32 `json:"mid"`
	Pts       int32 `json:"pts"`
	PtsCount  int32 `json:"pc"`
}

func (c *Core) makeDuplicateMessage(userId int64, randomId int64) (*mtproto.Updates, error) {
	s, err := redis.String(cache.GetRedisClient().Do("GET", makeMessageKey(userId, randomId)))
	if err != nil {
		log.Warnf("makeDuplicateMessage MessageId Key:%s error:%s", makeMessageKey(userId, randomId), err.Error())
		return nil, nil
	}

	if len(s) == 0 {
		return nil, nil
	}

	var cm DuplicateMessage
	err = jsoniter.UnmarshalFromString(s, &cm)
	if err != nil {
		log.Warnf("makeDuplicateMessage MessageId Key s:`%s` error:%s", makeMessageKey(userId, randomId), s, err.Error())
		return nil, nil
	}

	if cm.MessageId == 0 {
		log.Warnf("makeDuplicateMessage MessageId = 0 Key s:`%s`", makeMessageKey(userId, randomId), s)
		return nil, nil
	}

	messages, _ := c.GetMessages(userId, []int32{cm.MessageId}, false)
	if len(messages) == 0 {
		log.Warnf("makeDuplicateMessage Message Not Found Key userId:%d MessageId:%d s:`%s`", makeMessageKey(userId, randomId), userId, cm.MessageId, s)
		return nil, nil
	}

	return mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{
			mtproto.NewTLUpdateNewMessage(&mtproto.Update{Message1F2B0AFD71: messages[0], Pts: cm.Pts, PtsCount: cm.PtsCount}).To_Update(),
			mtproto.NewTLUpdateMessageID(&mtproto.Update{Id4E90BFD671: cm.MessageId, RandomId: randomId}).To_Update(),
		},
	}).To_Updates(), nil
}

func (c *Core) putDuplicateMessage(userId int64, randomId int64, messageId, pts, ptsCount int32) error {
	s, _ := jsoniter.MarshalToString(&DuplicateMessage{
		MessageId: messageId,
		Pts:       pts,
		PtsCount:  ptsCount,
	})
	_, err := cache.GetRedisClient().Do("SET", makeMessageKey(userId, randomId), s, "EX", 15)
	return err
}
