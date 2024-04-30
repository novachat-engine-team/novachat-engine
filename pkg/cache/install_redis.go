/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/24 15:23
 * @File : install_redis.go
 */

package cache

import (
	"go.uber.org/atomic"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
)

var _init *atomic.Bool
var _redisClient *RedisClient
func init() {
	_init = atomic.NewBool(false)
}

func InstallRedis(conf config.RedisConfig) error {
	if _init.Load() == true {
		return nil
	}

	var err error
	_redisClient, err = NewRedisClient(conf)
	if err != nil {
		log.Errorf("cache InstallRedis error:%v", err.Error())
		return err
	}

	_init.Store(true)
	return nil
}

func GetRedisClient() *RedisClient {
	return _redisClient
}
