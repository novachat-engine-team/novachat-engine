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

func (m *Core) PushUsernameUser(username string, info string) error {
	_, err := m.redisClient.Do(SET, MakeAccountCacheUsersUsernamePrefix(username), info)
	if err != nil {
		log.Errorf("PushUsernameUser username:%s info:%s error:%s", username, info, err.Error())
	}
	return err
}

func (m *Core) PushUsernameList(usernameList []string, infoList []string) error {

	v := make([]interface{}, 0, len(usernameList)*2)
	for idx := range usernameList {
		v = append(v, MakeAccountCacheUsersUsernamePrefix(usernameList[idx]))
		v = append(v, infoList[idx])
	}

	_, err := m.redisClient.Do(MSET, v...)
	if err != nil {
		log.Errorf("PushUsernameList usernameList:%v info:%v error:%s", usernameList, infoList, err.Error())
	}
	return err
}

func (m *Core) UsernameUser(username string) (string, error) {
	v, err := redis.String(m.redisClient.Do(GET, MakeAccountCacheUsersUsernamePrefix(username)))
	if err != nil {
		log.Errorf("UsernameUser username:%v error:%s", username, err.Error())
		return "", err
	}

	return v, nil
}

func (m *Core) UsernameUserList(usernameList []string) ([]string, error) {
	if len(usernameList) == 0 {
		return nil, nil
	}
	
	v := make([]interface{}, 0, len(usernameList)*2)
	for idx := range usernameList {
		v = append(v, MakeAccountCacheUsersUsernamePrefix(usernameList[idx]))
	}
	vv, err := redis.Strings(m.redisClient.Do(MGET, v...))
	if err != nil {
		log.Errorf("UsernameUserList usernameList:%v error:%s", usernameList, err.Error())
		return nil, err
	}
	return vv, nil
}
