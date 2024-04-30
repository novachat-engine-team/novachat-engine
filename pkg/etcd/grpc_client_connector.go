/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/26 16:08
 * @File : grpc_client_connector.go
 */

package etcd

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"novachat_engine/pkg/rpc/rpc_util"
)

var retryPolicy = `{
	"loadBalancingPolicy": "",
	"loadBalancingConfig": [{
	}],
	"methodConfig": [
	    // config per method or all methods under service
		"name": [{"service": "grpc.examples.echo.Echo"}],
        "waitForReady": true,
		"retryPolicy": {
			"MaxAttempts": 4,
			"InitialBackoff": ".01s",
			"MaxBackoff": ".01s",
			"BackoffMultiplier": 1.0,
			// this value is grpc code
			"RetryableStatusCodes": [ "UNAVAILABLE" ]
        }
	],
	"retryThrottling": {
		"maxTokens" : 1.0,
		"tokenRatio": 0.1
	}
	"healthCheckConfig": {
		"serviceName": ""
	}
}`

func NewGRPCClientConn(ctx context.Context, schema string, serverName string, balancer string, maxCallRecvMsgSize int32, maxCallSendMsgSize int32) (*grpc.ClientConn, error) {
	if maxCallRecvMsgSize <= 0 {
		maxCallRecvMsgSize = grpcMaxCallMsgSize
	}
	if maxCallSendMsgSize <= 0 {
		maxCallSendMsgSize = grpcMaxSendMsgSize
	}
	return grpc.DialContext(ctx, schema+":///"+serverName,
		//grpc.WithBalancerName(balancer),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy": "%s"}`, balancer)),
		grpc.WithInsecure(),
		grpc.WithInitialWindowSize(grpcInitialWindowSize),
		grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(int(maxCallRecvMsgSize))),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(int(maxCallSendMsgSize))),
		grpc.WithBackoffMaxDelay(grpcBackoffMaxDelay),
		//grpc.WithDefaultServiceConfig(retryPolicy),
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
	)
}

func NewGRPCClientConnAddr(ctx context.Context, addr string, maxCallRecvMsgSize int32, maxCallSendMsgSize int32) (*grpc.ClientConn, error) {
	if maxCallRecvMsgSize <= 0 {
		maxCallRecvMsgSize = grpcMaxCallMsgSize
	}
	if maxCallSendMsgSize <= 0 {
		maxCallSendMsgSize = grpcMaxSendMsgSize
	}
	return grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithInitialWindowSize(grpcInitialWindowSize),
		grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(int(maxCallRecvMsgSize))),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(int(maxCallSendMsgSize))),
		grpc.WithBackoffMaxDelay(grpcBackoffMaxDelay),
		//grpc.WithDefaultServiceConfig(retryPolicy),
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
	)
}