/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/5/12 16:36
 * @File : context.go
 */

package rpc_context

import (
	"context"
	"fmt"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/rpc/rpc_context"
	"novachat_engine/service/constants"
	"runtime"
)

type op struct {
    md          *metadata.RpcMetaData
    serverTrace []string
    timeout     int64
}

type OpOption func(option *op)

func WithBaseMD(md *metadata.RpcMetaData) OpOption {
    return func(op *op) {
        op.md = md
    }
}

func WithTimeout(s int64) OpOption {
    return func(option *op) {
        option.timeout = s
    }
}

func WithServerTrace(v string) OpOption {
    return func(option *op) {
        option.serverTrace = append(option.serverTrace, v)
    }
}

func RunFuncName() string {
    pc := make([]uintptr, 1)
    runtime.Callers(3, pc)
    f := runtime.FuncForPC(pc[0])
    return f.Name()
}

func MakeServerPeer(serverConfig *config.EtcdServerConfig, rpcServer *config.RpcServer) string {
    return serverConfig.ServerName + "/" + rpcServer.Addr
}

func FormatServerTrace(server string, funcName string) string {
    return fmt.Sprintf("%s/%s", server, funcName)
}

func Context(ctx context.Context, p ...OpOption) (context.Context, context.CancelFunc) {
    var cancelFunc context.CancelFunc
    var opV op
    opV.timeout = constants.RpcDefaultTimeout

    for _, o := range p {
        o(&opV)
    }

    ctx, cancelFunc = rpc_context.Context(ctx, opV.timeout)

    if opV.md == nil {
        opV.md = &metadata.RpcMetaData{
            ServerTrace: []string{},
        }
    }
    if len(opV.serverTrace) > 0 {
        opV.md.ServerTrace = append(opV.md.ServerTrace, opV.serverTrace...)
    }

    ctx = metadata.NewContextWithMetaData(ctx, opV.md)
    return ctx, cancelFunc
}
