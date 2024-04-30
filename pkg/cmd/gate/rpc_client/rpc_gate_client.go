/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/19 18:12
 * @File : rpc_gate_client.go
 */

package rpc_client

import (
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/rpc"
)

var _gateClient *etcd.ManualClient

func InstallGateClient(conf config.EtcdClientConfig) error {
	var err error
	_gateClient, err = etcd.NewManualClient(conf)

	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	_gateClient.Start()
	return nil
}

func GetGateClient(key string) rpc.RpcServiceClient {
	if _gateClient == nil {
		return nil
	}

	cli := _gateClient.GetClientConn(key)
	if cli != nil {
		return rpc.NewRpcServiceClient(cli)
	}

	return nil
}
