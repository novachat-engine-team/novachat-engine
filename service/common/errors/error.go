/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/7 16:44
 * @File : error.go
 */

package errors

import (
	"fmt"
	"novachat_engine/mtproto"
)

func NewRpcErrorWithCode(code int32) *mtproto.TLRpcError {
	err := mtproto.NewTLRpcError(nil)
	err.SetErrorCode(code)
	return err
}

func NewRpcErrorWithCodeString(code int32, msg string) *mtproto.TLRpcError {
	err := mtproto.NewTLRpcError(nil)
	err.SetErrorCode(code)
	err.SetErrorMessage(msg)
	return err
}

func NewRpcErrorWithRpcErrorCodeString(code mtproto.RpcErrorCode, msg string) *mtproto.TLRpcError {
	err := mtproto.NewTLRpcError(nil)
	err.SetErrorCode(code.GetClientCode())
	err.SetErrorMessage(msg)
	return err
}

func NewRpcErrorWithRpcErrorCode(code mtproto.RpcErrorCode) *mtproto.TLRpcError {
	err := mtproto.NewTLRpcError(nil)
	err.SetErrorCode(code.GetClientCode())
	err.SetErrorMessage(code.GetClientErrorString())
	return err
}

func NewRpcErrorWithRpcErrorCodeX(code mtproto.RpcErrorCode, x int32) *mtproto.TLRpcError {
	err := mtproto.NewTLRpcError(nil)
	err.SetErrorCode(code.GetClientCode())
	if code >= mtproto.RpcErrorCode_SEE_OTHER && code < mtproto.RpcErrorCode_SEE_OTHER_X_MAX {
		err.SetErrorMessage(fmt.Sprintf("%s_%d", code.GetClientErrorString(), x))
	} else {
		err.SetErrorMessage(code.GetClientErrorString())
	}
	return err
}
