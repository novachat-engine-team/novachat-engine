/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/18 15:32
 * @File : rpc_webrtc_client.go
 */

package rpc_client

import (
	"google.golang.org/grpc"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
)

var _webRTCClient *grpc.ClientConn

func InstallWebRTCClient(conf config.EtcdClientConfig) error {
	var err error
	_webRTCClient, err = etcd.NewEtcdClient(conf)

	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	return nil
}

func GetWebRTCClient() RelayServiceClient {
	if _relayClient == nil {
		return nil
	}

	return NewRelayServiceClient(_webRTCClient)
}
