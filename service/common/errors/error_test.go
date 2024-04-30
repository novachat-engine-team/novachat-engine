/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/3/30 19:22
 * @File : error_test.go
 */

package errors

import (
	"novachat_engine/mtproto"
	"testing"
)

func TestAAA(t *testing.T) {
	err := NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_FORBIDDEN_USER_NOT_MUTUAL_CONTACT)
	err.GetErrorMessage()
}
