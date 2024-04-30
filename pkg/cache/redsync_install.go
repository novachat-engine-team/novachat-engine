/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package cache

import (
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/atomic"
	"novachat_engine/pkg/config"
	"time"
)

var _initRedSync *atomic.Bool
var _redisClientRS *redsync.Redsync

func init() {
	_initRedSync = atomic.NewBool(false)
}

func InstallRedSync(conf config.RedisConfig) error {
	if _initRedSync.Load() == true {
		return nil
	}

	pool := make([]redsync.Pool, 0, len(conf.Addrs))
	for _, v := range conf.Addrs {
		p := &redis.Pool{
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", v)
				if err != nil {
					return nil, err
				}
				if len(conf.Passwd) > 0 {
					if _, err := c.Do("AUTH", conf.Passwd); err != nil {
						c.Close()
						return nil, err
					}
				}
				if _, err := c.Do("SELECT", conf.DBNum); err != nil {
					c.Close()
					return nil, err
				}
				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
			MaxIdle:         int(conf.Idle),
			MaxActive:       int(conf.Active),
			IdleTimeout:     0,
			Wait:            false,
			MaxConnLifetime: 0,
		}
		pool = append(pool, p)
	}

	_redisClientRS = redsync.New(pool)
	_initRedSync.Store(true)
	return nil
}

func GetRedSyncClient() *redsync.Redsync {
	return _redisClientRS
}
