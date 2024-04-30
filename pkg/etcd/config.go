/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 23:11
 * @File : config.go
 */

package etcd

import "time"

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 28
	grpcInitialConnWindowSize = 1 << 28
	grpcKeepAliveTime    = time.Duration(10) * time.Second
	grpcKeepAliveTimeout = time.Duration(10) * time.Second
	grpcBackoffMaxDelay  = time.Duration(3) * time.Second
	grpcMaxSendMsgSize   = 1 << 28
	grpcMaxCallMsgSize   = 1 << 28
)
