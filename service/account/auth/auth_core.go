/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/23 18:02
 * @File :
 */

package auth

import (
	"github.com/allegro/bigcache"
	jsoniter "github.com/json-iterator/go"
	authService "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	"novachat_engine/service/core/account/auth"
	dataAuth "novachat_engine/service/data/auth"
	"sync"
	"time"
)

type Core struct {
	authCore     *auth.Core
	usersCache   *bigcache.BigCache
	expiredCache sync.Map
}

func NewAuthCore(authCore *auth.Core) *Core {
	m := &Core{
		authCore: authCore,
	}
	var err error
	m.usersCache, err = bigcache.NewBigCache(
		bigcache.Config{
			Shards:             1024,
			LifeWindow:         3 * time.Minute,
			MaxEntriesInWindow: 1000 * 10 * 60,
			MaxEntrySize:       500,
			Verbose:            true,
			HardMaxCacheSize:   8192,
			OnRemove:           nil,
			OnRemoveWithReason: func(key string, entry []byte, reason bigcache.RemoveReason) {
				log.Debugf("OnRemove key:%s reason:%v", key, reason)
				switch reason {
				case bigcache.Expired:
					m.expiredCache.Store(key, entry)
				case bigcache.NoSpace:
					log.Warnf("OnRemove NoSpace key:%s", key)
				default:
					break
				}
			},
		})
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-ticker.C:
				m.expiredCache.Range(func(key, value interface{}) bool {
					m.usersCache.Set(key.(string), value.([]byte))
					m.expiredCache.Delete(key)
					return true
				})
			}
		}
	}()

	if err != nil {
		panic(err.Error())
	}
	return m
}

func (m *Core) UpdateAuthKey(auth *dataAuth.Auth) (bool, error) {
	ok, err := m.authCore.UpdateAuthKey(auth)

	key := makeAuthKeyIdKey(auth.AuthKeyId)
	data, err := jsoniter.Marshal(SessionInfo{
		AuthInfo:      auth,
		SessionIdList: []int64{},
	})
	err = m.usersCache.Set(key, data)
	if err != nil {
		log.Warnf("UpdateAuthKey authKeyId:%d error:%s", auth.AuthKeyId, err.Error())
	}
	return ok, err
}

func (m *Core) UpdateAuthKeyWithLayer(authKeyId int64, layer int32) (bool, error) {
	ok, err := m.authCore.UpdateAuthKeyWithLayer(authKeyId, layer)
	if err != nil {
		return ok, err
	}

	key := makeAuthKeyIdKey(authKeyId)
	var info SessionInfo
	entities, err := m.usersCache.Get(key)
	if err == nil {
		err = jsoniter.Unmarshal(entities, &info)
		if err != nil {
			log.Warnf("UpdateAuthKeyWithLayer authKeyId:%d error:%s", authKeyId, err.Error())
		} else {
			info.AuthInfo.Layer = layer
		}
	}

	if err == nil {
		data, _ := jsoniter.Marshal(info)
		err = m.usersCache.Set(key, data)
		if err != nil {
			log.Warnf("UpdateAuthKeyWithLayer authKeyId:%d error:%s", authKeyId, err.Error())
		}
	}

	return ok, nil
}

func (m *Core) AuthKey(authKeyId int64, sessionId int64, userId int64) (*dataAuth.Auth, error) {
	if userId == 0 {
		key := makeAuthKeyIdKey(authKeyId)
		entities, err := m.usersCache.Get(key)
		if err != nil || len(entities) == 0 {
			log.Infof("AuthKey not found key:%d", authKeyId)
		} else {
			var info SessionInfo
			err = jsoniter.Unmarshal(entities, &info)
			if err != nil {
				log.Infof("AuthKey authKeyId:%d error:%s", authKeyId, err.Error())
			} else {
				return info.AuthInfo, nil
			}
		}

		authInfo, err := m.authCore.AuthKey(authKeyId)
		if err != nil {
			return nil, err
		}
		info := SessionInfo{
			AuthInfo:      authInfo,
			SessionIdList: make([]int64, 0, 1),
		}
		if sessionId != 0 {
			info.SessionIdList = append(info.SessionIdList, sessionId)
		}

		value, ok := m.expiredCache.LoadAndDelete(key)
		if ok {
			var localInfo SessionInfo
			err = jsoniter.Unmarshal(value.([]byte), &localInfo)
			if err == nil && len(localInfo.SessionIdList) > 0 {
				for _, v := range localInfo.SessionIdList {
					if util.IndexInt64s(&info.SessionIdList, v) < 0 {
						info.SessionIdList = append(info.SessionIdList, v)
					}
				}
			}
		}

		log.Debugf("AuthKey authKeyId:%d sessionId:%d", authKeyId, sessionId)
		data, _ := jsoniter.Marshal(info)
		err = m.usersCache.Set(key, data)
		if err != nil {
			log.Warnf("AuthKey authKeyId:%d error:%s", authKeyId, err.Error())
		}

		log.Debugf("AuthKey from db:%d", authKeyId)
		return authInfo, nil
	} else {
		key := makeUserIdKey(userId)
		entities, err := m.usersCache.Get(key)
		if err != nil || len(entities) > 0 {
			log.Infof("AuthKey not found key:%d", authKeyId)
		} else {
			var info SessionInfo
			err = jsoniter.Unmarshal(entities, &info)
			if err != nil {
				log.Infof("AuthKey authKeyId:%d error:%s", authKeyId, err.Error())
			} else {
				if sessionId != 0 && util.IndexInt64s(&info.SessionIdList, sessionId) < 0 {
					info.SessionIdList = append(info.SessionIdList, sessionId)
					data, _ := jsoniter.Marshal(info)
					err = m.usersCache.Set(key, data)
					if err == nil {
						return info.AuthInfo, nil
					}
				} else {
					return info.AuthInfo, nil
				}
			}
		}

		authInfo, err := m.authCore.AuthKey(authKeyId)
		if err != nil {
			return nil, err
		}
		info := SessionInfo{
			AuthInfo:      authInfo,
			SessionIdList: make([]int64, 0, 1),
		}
		if sessionId != 0 {
			info.SessionIdList = append(info.SessionIdList, sessionId)
		}

		value, ok := m.expiredCache.LoadAndDelete(key)
		if ok {
			var localInfo SessionInfo
			err = jsoniter.Unmarshal(value.([]byte), &localInfo)
			if err == nil && len(localInfo.SessionIdList) > 0 {
				for _, v := range localInfo.SessionIdList {
					if util.IndexInt64s(&info.SessionIdList, v) < 0 {
						info.SessionIdList = append(info.SessionIdList, v)
					}
				}
			}
		}
		log.Debugf("AuthKey authKeyId:%d sessionId:%d", authKeyId, sessionId)
		data, _ := jsoniter.Marshal(info)
		err = m.usersCache.Set(key, data)
		if err != nil {
			log.Warnf("AuthKey authKeyId:%d error:%s", authKeyId, err.Error())
		}

		log.Debugf("AuthKey from db:%d", authKeyId)
		return authInfo, nil
	}
}

func (m *Core) BindTempAuthKeyAuthKeyId(authKeyId int64, permAuthKeyId int64, sessionId int64) (bool, error) {
	authValue, err := m.authCore.BindTempAuthKeyAuthKeyId(authKeyId, permAuthKeyId, sessionId)
	if err != nil {
		return false, err
	}

	info := SessionInfo{
		SessionIdList: []int64{},
	}
	var authInfo = authValue
	key := makeAuthKeyIdKey(authKeyId)
	entities, err := m.usersCache.Get(key)
	if err != nil {
		info.AuthInfo = authValue
	} else {
		err = jsoniter.Unmarshal(entities, &info)
		if err != nil {
			log.Infof("AuthKey authKeyId:%d error:%s", authKeyId, err.Error())
		}
		info.AuthInfo = authInfo
	}

	//TODO:(Coderxw) need check IOS app
	info.SessionIdList = append(info.SessionIdList, sessionId)

	data, _ := jsoniter.Marshal(info)
	err = m.usersCache.Set(key, data)
	if err != nil {
		log.Warnf("BindTempAuthKeyAuthKeyId authKeyId:%d error:%s", authKeyId, err.Error())
	}
	return true, nil
}

func (m *Core) BindUser(authKeyId int64, userId int64) (bool, error) {
	authValue, err := m.authCore.BindUser(authKeyId, userId)
	if err != nil {
		return false, err
	}

	info := SessionInfo{
		SessionIdList: []int64{},
	}
	key := makeAuthKeyIdKey(authKeyId)
	entities, err := m.usersCache.Get(key)
	if err != nil {
		info.AuthInfo = authValue
	} else {
		err = jsoniter.Unmarshal(entities, &info)
		if err != nil {
			log.Infof("AuthKey authKeyId:%d error:%s", authKeyId, err.Error())
		}
		info.AuthInfo = authValue
	}

	data, _ := jsoniter.Marshal(info)
	err = m.usersCache.Set(key, data)
	if err != nil {
		log.Warnf("BindUser authKeyId:%d error:%s", authKeyId, err.Error())
	}

	var authList []*dataAuth.Auth
	var usersInfo UsersInfo
	key = makeUserIdKey(userId)
	entities, err = m.usersCache.Get(key)
	if err != nil {
		authList, err = m.authCore.LoadAuthUser([]int64{userId})
	} else {
		err = jsoniter.Unmarshal(entities, &usersInfo)
		if err == nil {
			if util.IndexInt64s(&usersInfo.AuthKeyIdList, authKeyId) < 0 {
				usersInfo.AuthKeyIdList = append(usersInfo.AuthKeyIdList, authKeyId)
				data, _ = jsoniter.Marshal(&usersInfo)
				err = m.usersCache.Set(key, data)
				if err != nil {
					log.Warnf("BindUser authKeyId:%d error:%s", authKeyId, err.Error())
				}
			}
		} else {
			authList, err = m.authCore.LoadAuthUser([]int64{userId})
		}
	}

	if len(authList) > 0 {
		usersInfo.AuthKeyIdList = make([]int64, 0, len(authList))
		for _, v := range authList {
			usersInfo.AuthKeyIdList = append(usersInfo.AuthKeyIdList, v.AuthKeyId)
		}
		if util.IndexInt64s(&usersInfo.AuthKeyIdList, authKeyId) < 0 {
			usersInfo.AuthKeyIdList = append(usersInfo.AuthKeyIdList, authKeyId)
		}

		data, _ = jsoniter.Marshal(&usersInfo)
		err = m.usersCache.Set(key, data)
		if err != nil {
			log.Warnf("BindUser authKeyId:%d error:%s", authKeyId, err.Error())
		}
	}
	return true, nil
}

func (m *Core) UnbindUser(authKeyId int64, userId int64) (bool, error) {
	authValue, err := m.authCore.BindUser(authKeyId, 0)
	if err != nil {
		return false, err
	}

	info := SessionInfo{
		SessionIdList: []int64{},
	}
	key := makeAuthKeyIdKey(authKeyId)
	entities, err := m.usersCache.Get(key)
	if err != nil {
		info.AuthInfo = authValue
	} else {
		err = jsoniter.Unmarshal(entities, &info)
		if err != nil {
			log.Infof("AuthKey authKeyId:%d error:%s", authKeyId, err.Error())
		}
		info.AuthInfo = authValue
	}

	data, _ := jsoniter.Marshal(info)
	err = m.usersCache.Set(key, data)
	if err != nil {
		log.Warnf("UnbindUser authKeyId:%d error:%s", authKeyId, err.Error())
	}

	var authList []*dataAuth.Auth
	var usersInfo UsersInfo
	key = makeUserIdKey(userId)
	entities, err = m.usersCache.Get(key)
	if err != nil {
		authList, err = m.authCore.LoadAuthUser([]int64{userId})
	} else {
		err = jsoniter.Unmarshal(entities, &usersInfo)
		if err == nil {
		} else {
			authList, err = m.authCore.LoadAuthUser([]int64{userId})
		}
	}

	if len(authList) > 0 {
		usersInfo.AuthKeyIdList = make([]int64, 0, len(authList))
		for _, v := range authList {
			usersInfo.AuthKeyIdList = append(usersInfo.AuthKeyIdList, v.AuthKeyId)
		}
		if idx := util.IndexInt64s(&usersInfo.AuthKeyIdList, authKeyId); idx >= 0 {
			copy(usersInfo.AuthKeyIdList[idx:], usersInfo.AuthKeyIdList[idx+1:])
			usersInfo.AuthKeyIdList = usersInfo.AuthKeyIdList[:len(usersInfo.AuthKeyIdList)-1]
		}

		data, _ = jsoniter.Marshal(&usersInfo)
		err = m.usersCache.Set(key, data)
		if err != nil {
			log.Warnf("UnbindUser authKeyId:%d error:%s", authKeyId, err.Error())
		}
	}
	return true, nil
}

func (m *Core) RegisterDevice(userId int64, authKeyId int64, device *authService.Device) (bool, error) {
	authValue, err := m.authCore.RegisterDevice(authKeyId, device)
	if err != nil {
		return false, err
	}

	var info SessionInfo
	key := makeAuthKeyIdKey(authKeyId)
	entities, err := m.usersCache.Get(key)
	if err != nil {
		info.AuthInfo = authValue
	} else {
		err = jsoniter.Unmarshal(entities, &info)
		if err != nil {
			log.Infof("RegisterDevice authKeyId:%d error:%s", authKeyId, err.Error())
		}
		info.AuthInfo = authValue
	}

	data, _ := jsoniter.Marshal(info)
	err = m.usersCache.Set(key, data)
	if err != nil {
		log.Warnf("RegisterDevice authKeyId:%d error:%s", authKeyId, err.Error())
	}
	return true, nil
}

func (m *Core) UserSession(userList []int64) ([]*SessionInfo, error) {
	tempUserList := make([]int64, 0, len(userList))
	sessionInfo := make([]*SessionInfo, 0, len(userList))

	for _, v := range userList {
		entities, err := m.usersCache.Get(makeUserIdKey(v))
		if err != nil {
			tempUserList = append(tempUserList, v)
		} else {
			var usersInfo SessionInfo
			err = jsoniter.Unmarshal(entities, &usersInfo)
			if err != nil {
				log.Warnf("UserSession userId:%d  value:`%s` error:%s", v, string(entities), err.Error())

				tempUserList = append(tempUserList, v)
			} else {
				sessionInfo = append(sessionInfo, &usersInfo)
			}
		}
	}

	if len(tempUserList) > 0 {
		authList, err := m.authCore.LoadAuthUser(tempUserList)
		if err != nil {
			log.Warnf("UserSession LoadAuthUser tempUserList:%+v error:%s", tempUserList, err.Error())
		}
		for _, v := range authList {
			sessionInfo = append(sessionInfo, &SessionInfo{
				AuthInfo: v,
			})
		}
	}

	log.Debugf("UserSession userList:%v sessionInfo:%+v", userList, sessionInfo)
	return sessionInfo, nil
}

func (m *Core) Close(authKeyId int64, connId uint64) {
	_ = connId

	key := makeAuthKeyIdKey(authKeyId)
	m.expiredCache.Delete(key)
	m.usersCache.Delete(key)
}
