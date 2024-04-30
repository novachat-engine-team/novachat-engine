/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/12/22 17:01
 * @File : context.go
 */

package rpc_context

import (
	"context"
	"time"
)

func Context(ctx context.Context, s int64) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.TODO()
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(s))
	return ctx, cancel
}
