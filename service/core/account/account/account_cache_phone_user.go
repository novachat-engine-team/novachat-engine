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

func (m *Core) PushPhoneUser(phone string, info string) error {
	_, err := m.redisClient.Do(SET, MakeAccountCacheUsersPhoneNumberPrefix(phone), info)
	if err != nil {
		log.Errorf("PushPhoneUser phone:%s info:%s error:%s", phone, info, err.Error())
	}
	return err
}

func (m *Core) PushPhoneList(phoneList []string, infoList []string) error {

	v := make([]interface{}, 0, len(phoneList)*2)
	for idx := range phoneList {
		v = append(v, MakeAccountCacheUsersPhoneNumberPrefix(phoneList[idx]))
		v = append(v, infoList[idx])
	}

	_, err := m.redisClient.Do(MSET, v...)
	if err != nil {
		log.Errorf("PushPhoneList phoneList:%v info:%v error:%s", phoneList, infoList, err.Error())
	}
	return err
}

func (m *Core) PhoneUser(phone string) (string, error) {
	v, err := redis.String(m.redisClient.Do(GET, MakeAccountCacheUsersPhoneNumberPrefix(phone)))
	if err != nil && err != redis.ErrNil {
		log.Errorf("PhoneUser phone:%v error:%s", phone, err.Error())
		return "", err
	}

	return v, nil
}

func (m *Core) PhoneUserList(phoneList []string) ([]string, error) {
	v := make([]interface{}, 0, len(phoneList)*2)
	for idx := range phoneList {
		v = append(v, MakeAccountCacheUsersPhoneNumberPrefix(phoneList[idx]))
	}
	vv, err := redis.Strings(m.redisClient.Do(MGET, v...))
	if err != nil {
		log.Errorf("PhoneUserList phoneList:%v error:%s", phoneList, err.Error())
		return nil, err
	}
	return vv, nil
}
