/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package users

import (
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	"novachat_engine/service/core/account/account"
	data_users "novachat_engine/service/data/users"
)

func (c *Core) UserDataUsername(username string) (*data_users.Users, error) {
	var userInfo *data_users.Users
	userCache, err := c.accountCore.UsernameUser(username)
	if err != nil || len(userCache) == 0 {
		userInfo, err = c.usersCore.Username(username)
		if err != nil {
			log.Errorf("UserDataUsername error:%s", err.Error())
			return nil, err
		}
	} else {
		userInfo, err = account.CacheInfo2User(userCache)
		if err != nil {
			log.Warnf("UserDataUsername error:%s", err.Error())
		}
	}
	return userInfo, err
}

func (c *Core) UserData(userId int64) (*data_users.Users, error) {
	var userInfo *data_users.Users
	userCache, err := c.accountCore.User(userId)
	if err != nil {
		log.Warnf("GetUserData accountCore User error:%s", err.Error())
	}
	if len(userCache) == 0 {
		userInfo, err = c.usersCore.User(userId)
		if err != nil {
			log.Errorf("GetUserData usersCore User error:%s", err.Error())
			return nil, err
		}
	} else {
		userInfo, err = account.CacheInfo2User(userCache)
		if err != nil {
			log.Warnf("GetUserData CacheInfo2User error:%s", err.Error())
		}
	}

	return userInfo, err
}

func (c *Core) UserDataList(userIdList []int64) ([]*data_users.Users, error) {

	userInfoList := make([]*data_users.Users, 0, len(userIdList))
	userCacheList, err := c.accountCore.UserList(userIdList)
	if err != nil {
		log.Warnf("UserDataList accountCore User error:%s", err.Error())
	}

	for _, v := range userCacheList {
		userInfo, err1 := account.CacheInfo2User(v)
		if err1 == nil {
			userInfoList = append(userInfoList, userInfo)
		} else {
			log.Warnf("UserDataList CacheInfo2User error:%s", err1.Error())
		}
	}

	var lessUserInfoList []*data_users.Users
	lessUserIdList := make([]int64, 0, len(userIdList))
	for _, v := range userIdList {
		if util.Index(userInfoList, func(idx int) bool {
			return userInfoList[idx].Id == v
		}) < 0 {
			lessUserIdList = append(lessUserIdList, v)
		}
	}

	if len(lessUserIdList) == 0 {
		lessUserInfoList, err = c.usersCore.UserIdList(userIdList)
		if err != nil {
			log.Errorf("UserDataList usersCore User error:%s", err.Error())
			return nil, err
		}
	}

	userInfoList = append(userInfoList, lessUserInfoList...)
	return userInfoList, err
}

func (c *Core) PhoneUser(phone string) (*data_users.Users, error) {
	userCache, err := c.accountCore.PhoneUser(phone)
	if err != nil {
		log.Warnf("PhoneUser error:%s", err.Error())
	}

	var userInfo *data_users.Users
	if len(userCache) == 0 {
		userInfo, err = c.usersCore.UserByPhoneNumber(phone)
		if err != nil {
			log.Errorf("PhoneUser UserByPhoneNumber error:%s", err.Error())
			return nil, err
		}
	} else {
		userInfo, err = account.CacheInfo2User(userCache)
		if err != nil {
			log.Warnf("PhoneUser CacheInfo2User error:%s", err.Error())
		}
	}
	return userInfo, err
}

func (c *Core) PhoneUserList(phoneList []string) ([]*data_users.Users, error) {
	userCacheList, err := c.accountCore.PhoneUserList(phoneList)
	if err != nil {
		log.Warnf("PhoneUser error:%s", err.Error())
	}

	var userInfoList = make([]*data_users.Users, 0, len(phoneList))
	var userInfo *data_users.Users
	for _, v := range userCacheList {
		if len(v) == 0 {
			continue
		}

		userInfo, err = account.CacheInfo2User(v)
		if err != nil {
			log.Warnf("PhoneUserList CacheInfo2User error:%s", err.Error())
		} else {
			userInfoList = append(userInfoList, userInfo)
		}
	}

	lessPhoneList := make([]string, 0, len(phoneList))
	for _, v := range phoneList {
		if util.Index(userInfoList, func(idx int) bool {
			return userInfoList[idx].Phone.Valid && (userInfoList[idx].Phone.String == v)
		}) < 0 {
			lessPhoneList = append(lessPhoneList, v)
		}
	}

	if len(lessPhoneList) > 0 {
		localUserList, err1 := c.usersCore.UsernameList(lessPhoneList)
		if err1 != nil {
			log.Warnf("PhoneUserList UsernameList error:%s", err1.Error())
		} else {
			userInfoList = append(userInfoList, localUserList...)
		}
	}

	return userInfoList, nil
}
