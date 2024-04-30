/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 23:11
 * @File : codec.go
 */

package net

import "io"

type Codec interface {
	Init(rw io.ReadWriter)
	Receive() (interface{}, error)
	Send(interface{}) error
	Close() error
}

type ProtoCodec int32

const (
	MTProtoCodec ProtoCodec = 0
	CodecVersionMax ProtoCodec = 1
)

var _codecMakerLists []func()Codec

func init() {
	_codecMakerLists = make([]func()Codec, CodecVersionMax, CodecVersionMax)
}

func RegisterCodecMaker(ver ProtoCodec, maker func()Codec) {
	_codecMakerLists[ver] = maker
}

func GetCodecMaker(ver ProtoCodec) Codec {
	return _codecMakerLists[ver]()
}