/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/19 18:15
 * @File : rpc_session_client.go
 */

package rpc_client

import (
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
)

var _sessionManualClient *etcd.ManualClient
var _balance string

func InstallSessionManualClient(conf config.EtcdClientConfig) error {
	var err error
	_sessionManualClient, err = etcd.NewManualClient(conf)

	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	_balance = conf.Balancer
	_sessionManualClient.Start()
	return nil
}

func StopSessionManualClient() {
	if _sessionManualClient != nil {
		_sessionManualClient.Stop()
	}
}

func GetSessionClientByServerPeer(serverPeer string) SessionServiceClient {
	if _sessionManualClient != nil {
		cli := _sessionManualClient.GetClientConn(serverPeer)
		if cli != nil {
			return NewSessionServiceClient(cli)
		}
	}
	return nil
}

func GetSessionClientByKey(key string) SessionServiceClient {
	if _sessionManualClient != nil {
		cli := _sessionManualClient.GetClientConnByKey(key)
		if cli != nil {
			return NewSessionServiceClient(cli)
		}
	}
	return nil
}
