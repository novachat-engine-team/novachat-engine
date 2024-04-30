/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/19 18:15
 * @File : rpc_sfs_client.go
 */

package rpc_client

import (
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
)

var _sfsClient *etcd.ManualClient

func InstallSfsClient(conf config.EtcdClientConfig) error {
	var err error
	_sfsClient, err = etcd.NewManualClient(conf)

	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	_sfsClient.Start()
	return nil
}

func GetSfsClient(key string) SfsServiceClient {
	if _sfsClient == nil {
		return nil
	}

	cli := _sfsClient.GetClientConnByKey(key)
	if cli != nil {
		return NewSfsServiceClient(cli)
	}

	return nil
}
