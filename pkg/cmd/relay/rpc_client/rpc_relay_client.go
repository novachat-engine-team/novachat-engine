/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/18 15:30
 * @File : rpc_relay_client.go
 */

package rpc_client

import (
	"google.golang.org/grpc"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
)

var _relayClient *grpc.ClientConn
var _relayManualClient *etcd.ManualClient
var _balance string

func InstallRelayClient(conf config.EtcdClientConfig) error {
	var err error
	_relayClient, err = etcd.NewEtcdClient(conf)

	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	_balance = conf.Balancer
	return nil
}

func GetRelayClient() RelayServiceClient {
	if _relayClient == nil {
		return nil
	}

	return NewRelayServiceClient(_relayClient)
}

func InstallRelayManualClient(conf config.EtcdClientConfig) error {
	var err error
	_relayManualClient, err = etcd.NewManualClient(conf)

	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	_balance = conf.Balancer
	_relayManualClient.Start()
	return nil
}

func StopRelayManualClient() {
	if _relayManualClient != nil {
		_relayManualClient.Stop()
	}
}

func GetRelayClientByServerPeer(serverPeer string) RelayServiceClient {
	if _relayManualClient != nil {
		cli := _relayManualClient.GetClientConn(serverPeer)
		if cli != nil {
			return NewRelayServiceClient(cli)
		}
	}
	return nil
}

func GetRelayClientByKey(key string) RelayServiceClient {
	if _relayManualClient != nil {
		cli := _relayManualClient.GetClientConnByKey(key)
		if cli != nil {
			return NewRelayServiceClient(cli)
		}
	}
	return nil
}
