/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Syncor: Coder (coderxw@gmail.com)
 * @Time : 2021/4/16 15:01
 * @File : rpc_sync_client.go
 */

package rpc_client

import (
	"fmt"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
)

var _syncClient *etcd.ManualClient
var _balance string

func InstallSyncClient(conf config.EtcdClientConfig) error {
	var err error
	_syncClient, err = etcd.NewManualClient(conf)

	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	_balance = conf.Balancer
	_syncClient.Start()
	return nil
}

func GetSyncClient(key string) SyncServiceClient {
	if _syncClient == nil {
		return nil
	}

	cli := _syncClient.GetClientConnByKey(key)
	if cli != nil {
		return NewSyncServiceClient(cli)
	}

	return nil
}

func GetSyncClientById(id int64) SyncServiceClient {
	key := fmt.Sprintf("%d", id)
	return GetSyncClient(key)
}
