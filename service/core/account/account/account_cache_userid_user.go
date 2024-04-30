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

func (m *Core) PushUser(userId int64, info string) error {
	_, err := m.redisClient.Do(SET, MakeAccountCacheUsersPrefix(userId), info)
	if err != nil {
		log.Errorf("PushUser userId:%d info:%s error:%s", userId, info, err.Error())
	}
	return err
}

func (m *Core) PushUserList(userIdList []int64, infoList []string) error {

	v := make([]interface{}, 0, len(userIdList)*2)
	for idx := range userIdList {
		v = append(v, MakeAccountCacheUsersPrefix(userIdList[idx]))
		v = append(v, infoList[idx])
	}

	_, err := m.redisClient.Do(MSET, v...)
	if err != nil {
		log.Errorf("PushUserList userIdList:%v info:%v error:%s", userIdList, infoList, err.Error())
	}
	return err
}

func (m *Core) User(userId int64) (string, error) {
	v, err := redis.String(m.redisClient.Do(GET, MakeAccountCacheUsersPrefix(userId)))
	if err != nil {
		log.Errorf("User userId:%v error:%s", userId, err.Error())
		return "", err
	}

	return v, nil
}

func (m *Core) UserList(userIdList []int64) ([]string, error) {
	v := make([]interface{}, 0, len(userIdList)*2)
	for idx := range userIdList {
		v = append(v, MakeAccountCacheUsersPrefix(userIdList[idx]))
	}
	vv, err := redis.Strings(m.redisClient.Do(MGET, v...))
	if err != nil {
		log.Errorf("UserList userIdList:%v error:%s", userIdList, err.Error())
		return nil, err
	}
	return vv, nil
}
