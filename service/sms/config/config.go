/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/11/8 11:50
 * @File : config.go
 */

package config

type SmsPlatform string

func (m *SmsPlatform) String() string {
	return string(*m)
}

func FromString(m string) SmsPlatform {
	return SmsPlatform(m)
}

const (
	SmsPlatformNone SmsPlatform = ""
	SmsPlatformTest SmsPlatform = "test"
)
