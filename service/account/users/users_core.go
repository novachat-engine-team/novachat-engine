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
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	"novachat_engine/service/account/contact"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/users/sql_core"
	data_contact "novachat_engine/service/data/contact"
	data_users "novachat_engine/service/data/users"
)

type Core struct {
	contact     *contact.Core
	accountCore *account.Core
	usersCore   *sql_core.Core
}

func NewUsersCore(accountCore *account.Core, contact *contact.Core, usersCore *sql_core.Core) *Core {
	return &Core{
		accountCore: accountCore,
		contact:     contact,
		usersCore:   usersCore,
	}
}

func (c *Core) GetUserList(userId int64, userIdList []int64) ([]*mtproto.User, error) {
	if len(userIdList) == 0 {
		return nil, nil
	}

	contacts, err := c.contact.GetContactByIdList(userId, userIdList)
	if err != nil {
		log.Warnf("GetUserList - GetContactByIdList error:%s", err.Error())
	}

	contactsCacheMap := make(map[int64]*data_contact.Contact, len(contacts))
	for _, v := range contacts {
		contactsCacheMap[v.PeerId] = v
	}

	userInfoList := make([]*data_users.Users, 0, len(userIdList))
	userCacheList, err := c.accountCore.UserList(userIdList)
	if err != nil {
		log.Warnf("GetUserList UserList error:%s", err.Error())
	} else {
		var userInfo *data_users.Users
		for _, v := range userCacheList {
			if len(v) == 0 {
				continue
			}
			userInfo, _ = account.CacheInfo2User(v)
			if userInfo == nil {
				continue
			}
			userInfoList = append(userInfoList, userInfo)
		}
	}

	loseUserIdList := make([]int64, 0, len(userIdList))
	for _, v := range userIdList {
		if util.Index(userInfoList, func(idx int) bool {
			return userInfoList[idx].Id == v
		}) < 0 {
			loseUserIdList = append(loseUserIdList, v)
		}
	}
	if len(loseUserIdList) > 0 {
		localUserList, err := c.usersCore.UserIdList(loseUserIdList)
		if err != nil {
			log.Warnf("GetUserList UserIdList error:%s", err.Error())
		}
		userInfoList = append(userInfoList, localUserList...)
	}
	//
	var user *mtproto.User
	userList := make([]*mtproto.User, 0, len(userInfoList))
	for _, v := range userInfoList {
		vv, _ := contactsCacheMap[v.Id]
		user = UserCore2Users(v)
		if v.Id == userId {
			UserCoreSelfUsers(user)
		} else {
			user = UserCoreContactUser(user, vv != nil, vv != nil && vv.GetContact() > data_contact.MutualTypeMyContact)
		}
		userList = append(userList, user)
	}
	return userList, nil
}

func (c *Core) GetUser(userId int64, peerId int64) (*mtproto.User, error) {

	var userInfo *data_users.Users
	var err error
	userInfo, err = c.UserData(userId)
	if err != nil {
		log.Warnf("GetUser GetUserData error:%s", err.Error())
		return nil, err
	}

	var user *mtproto.User
	user = UserCore2Users(userInfo)
	if peerId == userId {
		UserCoreSelfUsers(user)
	} else {

		contact, err1 := c.contact.GetContactById(userId, peerId)
		if err1 != nil {
			log.Warnf("GetUser - GetContactById error:%s", err1.Error())
		}

		user = UserCoreContactUser(user, contact != nil, contact != nil && contact.GetContact() > data_contact.MutualTypeMyContact)
	}
	return user, nil
}

func (c *Core) SearchUserData(query string, offset int32, limit int32) ([]*data_users.Users, error) {
	return c.usersCore.Search(query, offset, limit)
}
