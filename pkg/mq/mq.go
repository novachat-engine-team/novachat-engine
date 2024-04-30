/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/16 13:25
 * @File : mq.go
 */

package mq

type Producer interface {
	SendMessage(key string, value[]byte) (partition int32, offset int64, err error)
	SendMessages(key []string, value[][]byte) error
	Close() error
	Topic() string
}

type ConsumerCallback interface {
	OnMessageData(key string, value[]byte) error
}

type Consumer interface {
	Run()
	SetOnMessageData(cb ConsumerCallback)
	Topic() []string
	Group() string
	Close()
}
