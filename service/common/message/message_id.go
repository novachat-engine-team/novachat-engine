/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/25 14:20
 * @File : message_id.go
 */

package message

import (
	"go.uber.org/atomic"
	"time"
)

var msgIdSeq = atomic.NewInt64(0)

func MakeNextMessageId(isRpc bool) int64 {
	timeNano := time.Now().UnixNano()
	ts := timeNano / 1e9
	ms := (timeNano % 1e9) / 1e6
	sid := msgIdSeq.Add(1) & 0x1ffff
	msgIdSeq.CAS(0x1ffff, 0)
	last := 1
	if !isRpc {
		last = 3
	}
	msgId := ts<<32 | ms<<21 | sid<<3 | int64(last)
	return msgId
}
