/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 23:51
 * @File : redis.go
 */

package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/letsfire/redigo"
	"github.com/letsfire/redigo/mode"
	"github.com/letsfire/redigo/mode/alone"
	"github.com/letsfire/redigo/mode/cluster"
	"github.com/letsfire/redigo/mode/sentinel"
	"novachat_engine/pkg/config"
	"time"
)

type RedisClient struct {
	redigo *redigo.Redigo
	mode   mode.IMode
}

func (r *RedisClient) GetModel() mode.IMode {
	return r.mode
}

func (r *RedisClient) Do(commandName string, args ...interface{}) (interface{}, error) {
	conn := r.mode.GetConn()

	reply, err := conn.Do(commandName, args...)
	if err != nil {
		if  _, ok := err.(redis.Error); ok {
			conn.Close()
			return nil, err
		}

		if nconn, nerr := r.mode.NewConn(); nerr != nil {
			return nil, nerr
		} else {
			reply, err = nconn.Do(commandName, args...)
			nconn.Close()
		}

	}
	conn.Close()

	return reply, err
}

type PipeListParam struct {
	Command string
	Args 	[]interface{}
}

func (r *RedisClient) RedisDoPipeline(params []PipeListParam) (interface{}, error) {
	conn := r.mode.GetConn()
	defer conn.Close()
	var err error

	for _, param := range params {
		if err != conn.Send(param.Command, param.Args...) {
			break
		}
	}

	if err = conn.Flush(); err != nil {
		return nil, err
	}

	ret := make([]interface{}, 0, len(params))
	for idx:=0;idx<len(params);idx++ {
		reply, err := conn.Receive()
		if err != nil {
			return nil, err
		} else {
			ret = append(ret, reply)
		}
	}
	return ret, nil
}


func NewRedisClient(conf config.RedisConfig) (*RedisClient, error) {
	client := &RedisClient{}
	var mode mode.IMode
	switch conf.Mode {
	case config.RedisModeAlone:
		mode = initAlone(&conf)
		break
	case config.RedisModeCluster:
		mode = initSentinel(&conf)
		break
	case config.RedisModeSentinel:
		mode = initCluster(&conf)
		break
	default:
		return nil, fmt.Errorf("redis.config.mode error")
	}

	client.redigo = redigo.New(mode)
	client.mode = mode

	return client, nil
}

func initAlone(conf *config.RedisConfig) mode.IMode {
	mode := alone.New(
		alone.Addr(conf.Addrs[0]),
		alone.PoolOpts(
			mode.MaxIdle(int(conf.Idle)),
			mode.MaxActive(int(conf.Active)),
			mode.IdleTimeout(0),
			mode.MaxConnLifetime(0),
			mode.TestOnBorrow(nil),
			mode.Wait(false),
		),
		alone.DialOpts(
			redis.DialReadTimeout(time.Second),
			redis.DialWriteTimeout(time.Second),
			redis.DialConnectTimeout(time.Second),
			redis.DialPassword(conf.Passwd),
			redis.DialDatabase(int(conf.DBNum)),
			redis.DialKeepAlive(time.Minute * 1),
			redis.DialNetDial(nil),
			redis.DialUseTLS(false),
			redis.DialTLSSkipVerify(false),
			redis.DialTLSConfig(nil),
		),
	)

	return mode
}

func initSentinel(conf *config.RedisConfig) mode.IMode {
	return sentinel.New(
		sentinel.Addrs(conf.Addrs),
		sentinel.PoolOpts(
			mode.MaxIdle(int(conf.Idle)),
			mode.MaxActive(int(conf.Active)),
			mode.IdleTimeout(0),
			mode.MaxConnLifetime(0),
			mode.TestOnBorrow(nil),
			mode.Wait(false),
		),
		sentinel.DialOpts(
			redis.DialReadTimeout(time.Second),
			redis.DialWriteTimeout(time.Second),
			redis.DialConnectTimeout(time.Second),
			redis.DialPassword(conf.Passwd),
			redis.DialDatabase(int(conf.DBNum)),
			redis.DialKeepAlive(time.Minute * 1),
			redis.DialNetDial(nil),
			redis.DialUseTLS(false),
			redis.DialTLSSkipVerify(false),
			redis.DialTLSConfig(nil),
		),
	)
}

func initCluster(conf *config.RedisConfig) mode.IMode {
	return cluster.New(
		cluster.Nodes(conf.Addrs),
		cluster.PoolOpts(
			mode.MaxIdle(int(conf.Idle)),
			mode.MaxActive(int(conf.Active)),
			mode.IdleTimeout(0),
			mode.MaxConnLifetime(0),
			mode.TestOnBorrow(nil),
			mode.Wait(false),
		),
		cluster.DialOpts(
			redis.DialReadTimeout(time.Second),
			redis.DialWriteTimeout(time.Second),
			redis.DialConnectTimeout(time.Second),
			redis.DialPassword(conf.Passwd),
			redis.DialDatabase(int(conf.DBNum)),
			redis.DialKeepAlive(time.Minute * 1),
			redis.DialNetDial(nil),
			redis.DialUseTLS(false),
			redis.DialTLSSkipVerify(false),
			redis.DialTLSConfig(nil),
		),
	)
}
