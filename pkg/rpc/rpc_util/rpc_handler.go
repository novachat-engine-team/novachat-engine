/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/20 19:55
 * @File : rpc_handler.go
 */

package rpc_util

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/rpc_error"
	"runtime/debug"
)

func BizUnaryRecoveryHandler(ctx context.Context, p interface{}) (err error) {
	log.Errorf("BizUnaryRecoveryHandler - %s", debug.Stack())

	switch p.(type) {
	case *mtproto.TLRpcError:
		code, _ := p.(*mtproto.TLRpcError)
		md, _ := rpc_error.RpcErrorToMD(code)
		grpc.SetTrailer(ctx, md)
		err = status.Errorf(codes.Unknown, "panic triggered rpc_error: {%v}", p)
	default:
		err = status.Errorf(codes.Unknown, "panic unknown triggered: %v", p)
	}
	log.Errorf("Panic: %v", err.Error())
	return
}

func BizUnaryRecoveryHandler2(ctx context.Context, p interface{}) (err error) {
	switch p.(type) {
	case *mtproto.TLRpcError:
		code, _ := p.(*mtproto.TLRpcError)
		md, _ := rpc_error.RpcErrorToMD(code)
		grpc.SetTrailer(ctx, md)
		err = p.(*mtproto.TLRpcError)
		// err = status.Errorf(codes.Unknown, "panic triggered rpc_error: {%v}", p)
	default:
		err = status.Errorf(codes.Unknown, "panic unknown triggered: %v", p)
	}
	return
}

func BizStreamRecoveryHandler(stream grpc.ServerStream, p interface{}) (err error) {
	return
}
