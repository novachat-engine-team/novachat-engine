/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package service

import (
	"github.com/gomodule/redigo/redis"
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/log"
	chatCore "novachat_engine/service/core/chat"
	data_chat "novachat_engine/service/data/chat"
)

type ChatCommon struct {
	userId         int64
	chatCommonCore *chatCore.CommonChatCore
	redisClient    *cache.RedisClient
	chatCore       *chatCore.Core
}

func NewChatCommon(chatCommonCore *chatCore.CommonChatCore, chatCore *chatCore.Core) (*ChatCommon, error) {
	var err error
	cc := &ChatCommon{
		chatCommonCore: chatCommonCore,
		redisClient:    cache.GetRedisClient(),
		chatCore:       chatCore,
	}

	return cc, err
}

func (m *ChatCommon) getChatCommon(userId int64) ([]*data_chat.Chat, error) {
	//TODO:(Coderxw) chatCommonInfo values

	log.Debugf("getChatCommon userId:%d", userId)
	chatIdList, err := redis.Int64s(m.redisClient.Do("SMEMBERS", makeCommonChatKey(userId)))
	if err != nil && err != redis.ErrNil {
		log.Warnf("getChatCommon userId:%d error:%s", userId, err.Error())
	}
	if err == redis.ErrNil {
		chatIdList, err = m.chatCommonCore.GetChatCommon(m.userId)
	}

	if len(chatIdList) == 0 {
		return nil, nil
	}

	paramList := make([]interface{}, 0, len(chatIdList)+1)
	paramList = append(paramList, makeCommonChatKey(userId))
	for _, v := range chatIdList {
		paramList = append(paramList, v)
	}

	_, err = m.redisClient.Do("SADD", paramList)
	if err != nil {
		log.Warnf("getChatCommon SADD userId:%d error:%s", userId, err.Error())
	}

	chatList, err := m.chatCore.GetChatList(chatCore.GetChatListWithIds(chatIdList))
	if err != nil {
		log.Errorf("getChatCommon userId:%d error:%s", userId, err.Error())
		return nil, err
	}

	return chatList, nil
}
