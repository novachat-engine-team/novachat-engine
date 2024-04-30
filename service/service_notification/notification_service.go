/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/5/6 16:11
 * @File : notification_service.go
 */

package service_notification

import (
	"fmt"
	"github.com/allegro/bigcache"
	jsoniter "github.com/json-iterator/go"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"sync"
)

var _notificationCache *bigcache.BigCache
var _chanData chan etcd.DataOperatorData
var _dataClient *etcd.DataClient
var _clientOC sync.Once

var _dataServer *etcd.EtcdDataServer
var _serverOC sync.Once

const (
	NotificationKey = "notification"
)

func InstallClientNotification(config config.EtcdClientConfig, enable bool) error {
	var err error
	_clientOC.Do(func() {
		_notificationCache, err = bigcache.NewBigCache(bigcache.Config{
			Shards: 1024,
		})
		if err != nil {
			panic(fmt.Sprintf("InstallNotification NewBigCache error:%s", err.Error()))
		}

		if enable {
			_chanData = make(chan etcd.DataOperatorData, 1024)
			_dataClient, err = etcd.NewDataClient(config, _chanData)
			if err != nil {
				close(_chanData)

				panic(fmt.Sprintf("InstallNotification NewDataClient error:%s", err.Error()))
			}

			_dataClient.Start()

			go func() {
				for v := range _chanData {
					if v.Operator == etcd.DataOperatorRemove {
						_notificationCache.Delete(NotificationKey)
					} else if v.Operator == etcd.DataOperatorAdd {
						_notificationCache.Set(NotificationKey, v.Value)
					}
				}
			}()
		}
	})
	return nil
}

func GetNotification()([]ServiceNotification, error) {
	entries, err := _notificationCache.Get(NotificationKey)
	if err != nil {
		return nil, err
	}
	var v []ServiceNotification
	err = jsoniter.Unmarshal(entries, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func InstallServerNotification(config config.EtcdServerConfig, enable bool) {
	_serverOC.Do(func() {
		if enable {
			_dataServer = etcd.NewEtcdDataServer(config)
		}
	})
}

func SetNotification(data []ServiceNotification) error {
	if _dataServer != nil {
		v, err := jsoniter.Marshal(data)
		if err != nil {
			return fmt.Errorf("jsoniter.Marshal error:%s", err.Error())
		}

		return _dataServer.Update(NotificationKey, v)
	}
	return nil
}