/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/5/6 16:49
 * @File : etcd_data_server.go
 */

package etcd

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"strings"
	"time"
)

const key = "key"

type EtcdDataServer struct {
	cli        *clientv3.Client
	etcdConfig config.EtcdServerConfig
}

func NewEtcdDataServer(config config.EtcdServerConfig) *EtcdDataServer {
	return &EtcdDataServer{
		cli:        nil,
		etcdConfig: config,
	}
}

func (s *EtcdDataServer) Update(name string, value []byte) error {
	var err error

	if s.cli == nil {
		if s.etcdConfig.MaxCallSendMsgSize <= 0 {
			s.etcdConfig.MaxCallSendMsgSize = grpcMaxSendMsgSize
		}
		if s.etcdConfig.MaxCallRecvMsgSize <= 0 {
			s.etcdConfig.MaxCallRecvMsgSize = grpcMaxCallMsgSize
		}

		if s.etcdConfig.DialTimeout <= 0 {
			s.etcdConfig.DialTimeout = 3
		}

		s.cli, err = clientv3.New(clientv3.Config{
			Endpoints:            strings.Split(s.etcdConfig.EtcdAddr, ";"),
			AutoSyncInterval:     0,
			DialTimeout:          time.Duration(s.etcdConfig.DialTimeout) * time.Second,
			DialKeepAliveTime:    time.Duration(s.etcdConfig.DialKeepAliveTime) * time.Second,
			DialKeepAliveTimeout: time.Duration(s.etcdConfig.DialKeepAliveTimeout) * time.Second,
			MaxCallSendMsgSize:   int(s.etcdConfig.MaxCallSendMsgSize),
			MaxCallRecvMsgSize:   int(s.etcdConfig.MaxCallRecvMsgSize),
			TLS:                  nil,
			Username:             s.etcdConfig.UserName,
			Password:             s.etcdConfig.Password,
			RejectOldCluster:     false,
			DialOptions: []grpc.DialOption{
				grpc.WithInsecure(),
				grpc.WithInitialWindowSize(grpcInitialWindowSize),
				grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
				grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(int(s.etcdConfig.MaxCallRecvMsgSize))),
				grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(int(s.etcdConfig.MaxCallSendMsgSize))),
				grpc.WithBackoffMaxDelay(grpcBackoffMaxDelay),
				grpc.WithKeepaliveParams(keepalive.ClientParameters{
					Time:                grpcKeepAliveTime,
					Timeout:             grpcKeepAliveTimeout,
					PermitWithoutStream: true,
				}),
			},
			LogConfig:           nil,
			Context:             nil,
			PermitWithoutStream: true,
		})

		if err != nil {
			return err
		}
	}

	if s.etcdConfig.TTL > 0 {
		go func() {
			var leaseKeepAliveResponse <-chan *clientv3.LeaseKeepAliveResponse
			for {
				leaseKeepAliveResponse, err = s.withAlive(name, key, value, s.etcdConfig.TTL)
				for _ = range leaseKeepAliveResponse {
					// do nothing
				}
				log.Errorf("withAlive leaseKeepAliveResponse close")
			}
		}()
		return nil
	} else {
		_, err = s.withAlive(name, key, value, 0)
		if err != nil {
			log.Error(err.Error())
		}
		return err
	}
}

func (s *EtcdDataServer) withAlive(name string, key string, value []byte, ttl int64) (<-chan *clientv3.LeaseKeepAliveResponse, error) {

	op := []clientv3.OpOption{}
	var err error
	var id clientv3.LeaseID
	if ttl > 0 {
		leaseResp, err := s.cli.Grant(context.Background(), ttl)
		if err != nil {
			return nil, err
		}
		op = append(op, clientv3.WithLease(leaseResp.ID))
		id = leaseResp.ID
	}

	_, err = s.cli.Put(context.Background(), MakeServerKey(s.etcdConfig.Schema, name, key), string(value), op...)
	if err != nil {
		return nil, err
	}

	if ttl > 0 {
		var leaseKeepAliveResponse <-chan *clientv3.LeaseKeepAliveResponse
		leaseKeepAliveResponse, err = s.cli.KeepAlive(context.Background(), id)
		if err != nil {
			return nil, err
		}
		return leaseKeepAliveResponse, nil
	}
	return nil, nil
}
