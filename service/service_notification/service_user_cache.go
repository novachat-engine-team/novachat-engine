/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/5/23 16:20
 * @File : service_user_cache.go
 */

package service_notification

import (
	"fmt"
	"github.com/allegro/bigcache"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"sync"
)

type userCacheType struct {
	TM int32 	`json:"t"`
	Order int32 `json:"o"`
}
func (m *userCacheType) GetTime() int32 {
	return m.TM
}
func (m *userCacheType) GetOrderId() int32 {
	return m.Order
}
func (m *userCacheType) String() string {
	v, _ := jsoniter.MarshalToString(m)
	return v
}

type IUserCache interface {
	GetTime() 		int32
	GetOrderId() 	int32
	String() 		string
}

type defaultUserCache struct {

}
func (m *defaultUserCache) GetTime() int32 {
	return -1
}
func (m *defaultUserCache) GetOrderId() int32 {
	return UserCacheNone
}
func (m *defaultUserCache) String() string {
	return ""
}

var _userNotificationCache *bigcache.BigCache
var _userCacheOC sync.Once
var _defaultUserCache = &defaultUserCache{}
const UserCacheNone int32 = -1


func init() {
	var err error
	_userCacheOC.Do(func() {
		_userNotificationCache, err = bigcache.NewBigCache(bigcache.Config{
			Shards: 2048,
		})

		if err != nil {
			panic(fmt.Sprintf("service_notification init _userNotificationCache NewBigCache error:%s", err.Error()))
		}
	})
}

func GetUser(userId int32) IUserCache {
	key := strconv.FormatInt(int64(userId), 10)
	entries, err := _userNotificationCache.Get(key)
	if err != nil || len(entries) == 0 {
		return _defaultUserCache
	}

	v := &userCacheType{}
	err = jsoniter.Unmarshal(entries, v)
	if err != nil {
		fmt.Printf("GetUser IUserCache error:%s\n", err.Error())
	}

	return v
}

func SetUser(userId int32, tm int32, orderId int32) error {
	key := strconv.FormatInt(int64(userId), 10)

	v := &userCacheType{
		TM:    tm,
		Order: orderId,
	}

	entries, err := jsoniter.Marshal(v)
	if err != nil {
		return fmt.Errorf("userId:%d tm:%d orderId:%d Marshal error:%s", userId, tm, orderId, err.Error())
	}

	err = _userNotificationCache.Set(key, entries)
	if err != nil {
		return fmt.Errorf("userId:%d tm:%d orderId:%d error:%s", userId, tm, orderId, err.Error())
	}

	return nil
}

func SetUserCacheString(userId int32, s string) error {
	key := strconv.FormatInt(int64(userId), 10)

	v := &userCacheType{}
	err := jsoniter.UnmarshalFromString(s, v)
	if err != nil {
		return fmt.Errorf("SetUserCacheString userId:%d s:`%s` UnmarshalFromString error:%s", userId, s, err.Error())
	}

	err = _userNotificationCache.Set(key, []byte(s))
	if err != nil {
		return fmt.Errorf("SetUserCacheString userId:%d s:`%s` error:%s", userId, s, err.Error())
	}

	return nil
}
