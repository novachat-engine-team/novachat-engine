/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File :
 * @Desc :
 *
 */

package users

import (
	"fmt"
	"io"
)

func ReadRandomInt64(reader io.Reader) int64 {
	var b [8]byte
	_, err := io.ReadFull(reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot initialize objectid package with crypto.rand.Reader: %v", err))
	}

	return (int64(b[0]) << 0) | (int64(b[1]) << 8) | (int64(b[2]) << 16) | (int64(b[3]) << 24) |
		(int64(b[4]) << 32) | (int64(b[5]) << 40) | (int64(b[6]) << 48) | (int64(b[7]) << 56)
}
