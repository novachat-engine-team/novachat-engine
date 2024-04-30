/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/16 15:01
 * @File : rpc_chat_client.go
 */

package rpc_client

import (
	"fmt"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
)

var _chatClient *etcd.ManualClient
var _balance string

func InstallChatClient(conf config.EtcdClientConfig) error {
	var err error
	_chatClient, err = etcd.NewManualClient(conf)

	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	_balance = conf.Balancer
	_chatClient.Start()
	return nil
}

func getChatClient(key string) ChatServiceClient {
	if _chatClient == nil {
		return nil
	}

	cli := _chatClient.GetClientConnByKey(key)
	if cli != nil {
		return NewChatServiceClient(cli)
	}

	return nil
}

func GetChatClientByKeyId(keyId int64) ChatServiceClient {
	key := fmt.Sprintf("%d", keyId)
	return getChatClient(key)
}
