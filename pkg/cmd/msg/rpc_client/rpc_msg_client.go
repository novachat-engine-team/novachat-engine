/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/28 18:25
 * @File : rpc_msg_client.go
 */

package rpc_client

import (
	"fmt"
	"google.golang.org/grpc"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/etcd"
	"novachat_engine/pkg/log"
)

var _msgClient *grpc.ClientConn

func InstallMsgClient(conf config.EtcdClientConfig) error {
	var err error
	_msgClient, err = etcd.NewEtcdClient(conf)

	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	return nil
}

func GetMsgClient() MsgServiceClient {
	if _msgClient == nil {
		return nil
	}

	return NewMsgServiceClient(_msgClient)
}

// GetMsgByUserIdKey
// TODO:(Coderxw)
func GetMsgByUserIdKey(userId int64) string {
	key := fmt.Sprintf("%d", userId)
	return key
}
