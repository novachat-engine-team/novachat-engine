/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 23:11
 * @File : resolver.go
 */

package etcd

// etcd client

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/rpc_util"
	"strings"
	"time"
)

type etcdResolver struct {
	cc         resolver.ClientConn
	cli        *clientv3.Client
	etcdConfig config.EtcdClientConfig
}

func MakeWatchKey(schema string, name string) string {
	return fmt.Sprintf("/%s/%s/", schema, name)
}

func MakeServerKey(schema string, name string, addr string) string {
	return fmt.Sprintf("/%s/%s/%s", schema, name, addr)
}

func NewResolver(conf config.EtcdClientConfig) resolver.Builder {
	return &etcdResolver{
		etcdConfig: conf,
	}
}

func (r *etcdResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	log.Infof("etcdResolver::Build target:%v cc:%v, opts:%v", target, cc, opts)

	var err error

	if r.cli == nil {
		if r.etcdConfig.MaxCallSendMsgSize <= 0 {
			r.etcdConfig.MaxCallSendMsgSize = grpcMaxSendMsgSize
		}
		if r.etcdConfig.MaxCallRecvMsgSize <= 0 {
			r.etcdConfig.MaxCallRecvMsgSize = grpcMaxCallMsgSize
		}
		if r.etcdConfig.DialTimeout == 0 {
			r.etcdConfig.DialTimeout = 3
		}

		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(r.etcdConfig.DialTimeout)*time.Second)
		defer cancelFunc()
		r.cli, err = clientv3.New(clientv3.Config{
			Endpoints:            strings.Split(r.etcdConfig.EtcdAddr, ";"),
			AutoSyncInterval:     0,
			DialTimeout:          time.Duration(r.etcdConfig.DialTimeout) * time.Second,
			DialKeepAliveTime:    time.Duration(r.etcdConfig.DialKeepAliveTime) * time.Second,
			DialKeepAliveTimeout: time.Duration(r.etcdConfig.DialKeepAliveTimeout) * time.Second,
			MaxCallSendMsgSize:   int(r.etcdConfig.MaxCallSendMsgSize),
			MaxCallRecvMsgSize:   int(r.etcdConfig.MaxCallRecvMsgSize),
			TLS:                  nil,
			Username:             r.etcdConfig.UserName,
			Password:             r.etcdConfig.Password,
			RejectOldCluster:     false,
			DialOptions: []grpc.DialOption{
				grpc.WithInsecure(),
				grpc.WithInitialWindowSize(grpcInitialWindowSize),
				grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
				grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(int(r.etcdConfig.MaxCallRecvMsgSize))),
				grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(int(r.etcdConfig.MaxCallSendMsgSize))),
				grpc.WithBackoffMaxDelay(grpcBackoffMaxDelay),
				grpc.WithKeepaliveParams(keepalive.ClientParameters{
					Time:                grpcKeepAliveTime,
					Timeout:             grpcKeepAliveTimeout,
					PermitWithoutStream: true,
				}),
				grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
					rpc_util.WithUnaryClientInterceptor(
						rpc_util.WithUnaryClientRecoveryHandler(rpc_util.UnaryClientRecoverHandler),
						rpc_util.WithUnaryClientHandler(rpc_util.UnaryClientHandler)),
				),
				),
				grpc.WithChainStreamInterceptor(rpc_util.WithUnaryStreamClientInterceptor(
					rpc_util.WithUnaryClientStreamHandler(rpc_util.UnaryClientRecoverHandler)),
				),
			},
			LogConfig:           nil,
			Context:             ctx,
			PermitWithoutStream: true,
		})
		if err != nil {
			return nil, err
		}
	}

	r.cc = cc

	watchChan := make(chan error)
	go r.watch(MakeWatchKey(target.Scheme, target.Endpoint), watchChan)

	err = <-watchChan
	defer close(watchChan)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *etcdResolver) Scheme() string {
	return r.etcdConfig.Schema
}

func (r *etcdResolver) ResolveNow(rr resolver.ResolveNowOptions) {
	log.Infof("etcdResolver::ResolveNow rr:%v", rr)
}

func (r *etcdResolver) Close() {
	log.Infof("etcdResolver::Close")
}

func (r etcdResolver) watch(keyPrefix string, watchChan chan error) {
	var addrList []resolver.Address

	getResp, err := r.cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		watchChan <- err
		return
	}

	for i := range getResp.Kvs {
		addrList = append(addrList, resolver.Address{Addr: strings.TrimPrefix(string(getResp.Kvs[i].Key), keyPrefix)})
	}

	r.cc.UpdateState(resolver.State{
		Addresses:     addrList,
		ServiceConfig: nil,
	})

	rch := r.cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	watchChan <- err
	for n := range rch {
		for _, ev := range n.Events {
			addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
			switch ev.Type {
			case mvccpb.PUT:
				if !exist(addrList, addr) {
					addrList = append(addrList, resolver.Address{Addr: addr})

					r.cc.UpdateState(resolver.State{
						Addresses:     addrList,
						ServiceConfig: nil,
					})
				}
			case mvccpb.DELETE:
				if s, ok := remove(addrList, addr); ok {
					addrList = s
					r.cc.UpdateState(resolver.State{
						Addresses:     addrList,
						ServiceConfig: nil,
					})
				}
			}
		}
	}
}

func exist(l []resolver.Address, addr string) bool {
	for i := range l {
		if l[i].Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}
