/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/20 19:51
 * @File : rpc_util.go
 */

package rpc_util

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RecoveryHandlerFunc is a function that recovers from the panic `p` by returning an `error`.
type UnaryRecoveryHandlerFunc func(ctx context.Context, p interface{}) (err error)
type StreamRecoveryHandlerFunc func(stream grpc.ServerStream, p interface{}) (err error)

// UnaryServerInterceptor returns a new unary server interceptor for panic recovery.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = unaryRecoverFrom(ctx, r, o.unaryRecoveryHandlerFunc)
			}
		}()

		r, err2 := handler(ctx, req)
		if err2 != nil {
			err = unaryRecoverFrom(ctx, err2, o.unaryRecoveryHandlerFunc2)
			return r, err
		} else {
			return r, nil
		}
	}
}

// StreamServerInterceptor returns a new streaming server interceptor for panic recovery.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateOptions(opts)
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = streamRecoverFrom(stream, r, o.streamRecoveryHandlerFunc)
			}
		}()

		return handler(srv, stream)
	}
}

func unaryRecoverFrom(ctx context.Context, p interface{}, f UnaryRecoveryHandlerFunc) error {
	if f == nil {
		return status.Errorf(codes.Internal, "%s", p)
	}
	return f(ctx, p)
}

func streamRecoverFrom(stream grpc.ServerStream, p interface{}, f StreamRecoveryHandlerFunc) error {
	if f == nil {
		return status.Errorf(codes.Internal, "%s", p)
	}
	return f(stream, p)
}

