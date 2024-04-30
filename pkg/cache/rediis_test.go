/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/4/22 18:26
 * @File : rediis_test.go
 */

package cache

import (
	"novachat_engine/pkg/config"
	"testing"
)

func TestRedisTest(t *testing.T) {
	InstallRedis(config.RedisConfig{
		Addrs:  []string{"192.168.1.80:6379"},
		Mode:   "alone",
		Test:   false,
		Idle:   10,
		Active: 3,
		Passwd: "",
		DBNum:  0,
	})

	GetRedisClient().Do("get", 1)

	GetRedisClient().Do("get", 1)
}
