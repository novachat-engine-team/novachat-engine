/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/3/28 18:13
 * @File : rpc_client_handler.go
 */

package rpc_util

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"novachat_engine/pkg/rpc/rpc_error"
)

type UnaryClientCallback func(ctx context.Context, md metadata.MD, p interface{}) error

func UnaryClientRecoverHandler(ctx context.Context, md metadata.MD, p interface{}) error {
	err, ok := p.(error)
	if ok == true {
		err = rpc_error.RpcFromError(err, md)
	} else {
		err = status.Errorf(codes.Unknown, "panic unknown client recover triggered: %v", p)
	}
	return err
}

func UnaryClientHandler(ctx context.Context, md metadata.MD, p interface{}) error {
	err, ok := p.(error)
	if ok == true {
		err = rpc_error.RpcFromError(err, md)
	} else {
		err = status.Errorf(codes.Unknown, "panic unknown client triggered: %v", p)
	}
	return err
}

