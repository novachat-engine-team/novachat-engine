/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 23:11
 * @File : etcd_server.go
 */

package etcd

import (
	"novachat_engine/pkg/config"
	_ "novachat_engine/pkg/etcd/balancer"
)

//conf := config.EtcdServerConfig{
//	EtdAddr:              "http://192.168.95.131:2379",
//	Schema:               "project",
//	ServerName: 		  svcName,
//	TTL:                  15,
//	DialTimeout:          15,
//	MaxCallSendMsgSize : 1024,
//	MaxCallRecvMsgze:  1024,
//}

type EtcdServer EtcdNamingServer

func NewEtcdServer(conf config.EtcdServerConfig, serverAddr string) (*EtcdNamingServer, error) {
	etcdNamingServer := NewEtcdNamingServer(conf)

	err := etcdNamingServer.Register(conf.ServerName, serverAddr)
	if err != nil {
		return nil, err
	}

	return etcdNamingServer, nil
}

func StopEtcdServer(s *EtcdNamingServer, serverName, addr string) {
	s.UnRegister(serverName, addr)
}

