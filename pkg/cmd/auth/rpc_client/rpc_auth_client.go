/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/16 15:01
 * @File : rpc_auth_client.go
 */

package rpc_client

import (
	"fmt"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
)

var _authClient *etcd.ManualClient
var _balance string

func InstallAuthClient(conf config.EtcdClientConfig) error {
	var err error
	_authClient, err = etcd.NewManualClient(conf)

	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	_balance = conf.Balancer
	_authClient.Start()
	return nil
}

func GetAuthClient(key string) AuthServiceClient {
	if _authClient == nil {
		return nil
	}

	cli := _authClient.GetClientConnByKey(key)
	if cli != nil {
		return NewAuthServiceClient(cli)
	}

	return nil
}

func GetAuthClientByAuthKey(authKeyId int64) AuthServiceClient {
	key := fmt.Sprintf("%d", authKeyId)
	return GetAuthClient(key)
}
