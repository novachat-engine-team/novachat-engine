/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/21 17:12
 * @File : banned.go
 */

package banned_memory

import (
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/log"
	"novachat_engine/service/core/account/banned_phone_number"
)

const (
	BannedNameMemory = "bannedNameMemory"
)

func init() {
	banned_phone_number.Register(BannedNameMemory, &bannedNameMemory{})
}

type bannedNameMemory struct {
	redisClient *cache.RedisClient
}

func (m *bannedNameMemory) Initialize() error {
	m.redisClient = cache.GetRedisClient()
	if m.redisClient == nil {
		log.Fatalf("cache.GetRedisClient == nil")
	}
	return nil
}

func (m *bannedNameMemory) BannedNumber(number string) error {
	return nil
}

func (m *bannedNameMemory) InvalidNumber(number string) bool {
	return false
}
