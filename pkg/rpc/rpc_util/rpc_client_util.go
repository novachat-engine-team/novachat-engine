/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/3/28 16:48
 * @File : rpc_client_util.go
 */

package rpc_util

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func WithUnaryClientInterceptor(op ...ClientOption) grpc.UnaryClientInterceptor {
	o := evaluateClientOptions(op)
	return func (ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = unaryClientHandler(ctx, r, metadata.MD{}, o.UnaryClientRecover)
			}
		}()

		var header, trailer metadata.MD
		opts = combine(opts, []grpc.CallOption{
			grpc.Header(&header), grpc.Trailer(&trailer),
		})

		err = invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			err = unaryClientHandler(ctx, err, trailer, o.UnaryClient)
		}
		return err
	}
}

func WithUnaryStreamClientInterceptor(op ...ClientOption) grpc.StreamClientInterceptor {
	o := evaluateClientOptions(op)
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (s grpc.ClientStream, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = unaryClientHandler(ctx, r, metadata.MD{}, o.UnaryClientRecover)
			}
		}()

		return streamer(ctx, desc, cc, method, opts...)
	}
}

func unaryClientHandler(ctx context.Context, p interface{}, md metadata.MD, f UnaryClientCallback) error {
	if f == nil {
		return status.Errorf(codes.Internal, "%+v", p)
	}
	return f(ctx, md, p)
}

func combine(o1 []grpc.CallOption, o2 []grpc.CallOption) []grpc.CallOption {
	if len(o1) == 0 {
		return o2
	} else if len(o2) == 0 {
		return o1
	}
	ret := make([]grpc.CallOption, len(o1)+len(o2))
	copy(ret, o1)
	copy(ret[len(o1):], o2)
	return ret
}