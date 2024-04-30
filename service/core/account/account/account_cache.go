/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/16 17:15
 * @File :
 */

package account

import (
	"github.com/gomodule/redigo/redis"
	"novachat_engine/pkg/log"
)

// Core redisbloom
//BF.RESERVE {key} {error_rate} {capacity} [EXPANSION expansion] [NONSCALING]
//BF.ADD	BF.ADD {key} {item}
//BF.MADD	{key} {item} [item…]
//BF.INSERT	BF.INSERT {key} [CAPACITY {cap}] [ERROR {error}] [EXPANSION expansion] [NOCREATE] [NONSCALING] ITEMS {item…}
//BF.EXISTS	 {key} {item}
//BF.MEXISTS	{key} {item} [item…]
//BF.SCANDUMP	{key} {iter}
//BF.LOADCHUNK	{key} {iter} {data}
//BF.INFO	{key}
//BF.DEBUG	BF.DEBUG {key}

func (m *Core) push(key string, l string) (bool, error) {
	v, err := redis.Int(m.redisClient.Do(BFAdd, key, l))
	if err != nil && err != redis.ErrNil {
		log.Errorf("push l:%s error:%s", l, err.Error())
		return false, err
	}

	if err == redis.ErrNil {
		return true, nil
	}
	return v == 1, nil
}

func (m *Core) pushList(key string, l []string) ([]bool, error) {
	ui := make([]interface{}, 0, len(l)+1)
	ui = append(ui, key)
	for _, v := range l {
		ui = append(ui, v)
	}
	v, err := redis.Ints(m.redisClient.Do(BFMAdd, ui...))
	if err != nil && err != redis.ErrNil {
		log.Errorf("pushList l:%+v error:%s", l, err.Error())
		return nil, err
	}

	if err == redis.ErrNil {
		return nil, nil
	}

	ret := make([]bool, 0, len(l))
	for _, vv := range v {
		ret = append(ret, vv == 1)
	}
	return ret, nil
}

func (m *Core) exists(key string, l string) (bool, error) {
	v, err := redis.Int(m.redisClient.Do(BFExists, key, l))
	if err != nil && err != redis.ErrNil {
		log.Errorf("exists l:%s error:%s", l, err.Error())
		return false, err
	}

	if err == redis.ErrNil {
		return true, nil
	}
	return v == 1, nil
}

func (m *Core) listExists(key string, l []string) ([]bool, error) {
	ui := make([]interface{}, 0, len(l)+1)
	ui = append(ui, key)
	for _, v := range l {
		ui = append(ui, v)
	}
	v, err := redis.Ints(m.redisClient.Do(BFMExists, ui...))
	if err != nil && err != redis.ErrNil {
		log.Errorf("listExists l:%+v error:%s", l, err.Error())
		return nil, err
	}

	if err == redis.ErrNil {
		return nil, nil
	}

	ret := make([]bool, 0, len(l))
	for _, vv := range v {
		ret = append(ret, vv == 1)
	}
	return ret, nil
}
