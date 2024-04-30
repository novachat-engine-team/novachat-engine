/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/20 19:40
 * @File : rpc_error.go
 */

package rpc_error

import (
	"encoding/base64"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"novachat_engine/mtproto"
	"novachat_engine/service/common/errors"
)

var (
	headerRpcError = "rpc_error"
)

func RpcErrorFromMD(md metadata.MD) (*mtproto.TLRpcError, bool) {
	var rpcErr *mtproto.TLRpcError
	val := metautils.NiceMD(md).Get(headerRpcError)
	if val == "" {
		//rpcErr = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, fmt.Sprintf("Unknown error"))
		//log.Errorf("%v", rpcErr)
		return nil, false
	}

	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		//rpcErr = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, fmt.Sprintf("Base64 decode error, rpc_error: %s, error: %v", val, err))
		//log.Errorf("%v", rpcErr)
		return nil, false
	}

	rpcErr = &mtproto.TLRpcError{}
	err = proto.Unmarshal(buf, rpcErr)
	if err != nil {
		rpcErr = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, fmt.Sprintf("RpcError unmarshal error, rpc_error: %s, error: %v", val, err))
		//log.Errorf("%v", rpcErr)
		//return rpcErr, false
		return nil, false
	}

	return rpcErr, true
}

func RpcErrorToMD(md *mtproto.TLRpcError) (metadata.MD, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		//log.Errorf("Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metadata.Pairs(headerRpcError, base64.StdEncoding.EncodeToString(buf)), nil
}

func RpcFromError(err error, trailer metadata.MD) error {

	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.Unknown:
			err1, ok := RpcErrorFromMD(trailer)
			if ok == true {
				return err1
			}
			break
		default:
			break
		}
	}
	return err
}