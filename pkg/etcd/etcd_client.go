/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 23:11
 * @File : etcd_client.go
 */

package etcd

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"novachat_engine/pkg/config"
	_ "novachat_engine/pkg/etcd/balancer"
	"novachat_engine/pkg/log"
	"time"
)

//conf := config.EtcdClientConfig{
//	Schema:               "project",
//	ServerName: 		  "server",
//	DialTimeout:          0,
//	EtcdAddr:             "192.168.95.131:2379",
//	DialKeepAliveTime:    0,
//	DialKeepAliveTimeout: 0,
//	MaxCallSendMsgSize:   1024,
//	MaxCallRecvMsgSize:   1024,
//	UserName:             "",
//	Password:             "",
//}

func NewEtcdClient(conf config.EtcdClientConfig) (*grpc.ClientConn, error) {
	r := NewResolver(conf)
	resolver.Register(r)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)

	//conn, err := grpc.Dial(r.Scheme()+":///" + conf.ServerName, grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	conn, err := NewGRPCClientConn(ctx, r.Scheme(), conf.ServerName, conf.Balancer, conf.MaxCallRecvMsgSize, conf.MaxCallSendMsgSize)
	if err != nil {
		cancel()
		log.Errorf("NewEtcdClient grpc.Dial error:%v", err.Error())
	}
	return conn, err
}
