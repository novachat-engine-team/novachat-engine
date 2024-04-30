/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/11 14:16
 * @File : without_login_config.go
 */

package conf

import "strings"

var _updatesGetDifference = "mtproto.TLUpdatesGetDifference"
var _upload = "mtproto.TLUpload"
var _withoutLoginConfig = []string{
	"mtproto.TLAuth",
	"mtproto.TLPhotos",
	"mtproto.TLHelp",
	"mtproto.TLLangpack",
	"mtproto.TLUpload",
}

func WithoutLoginConfig(rpcName string) bool {
	rpcNameLen := len(rpcName)
	if rpcNameLen == 0 {
		return true
	}
	if rpcName[0] == '*' {
		rpcName = rpcName[1:]
		rpcNameLen = len(rpcName)
	}
	if rpcNameLen == 0 {
		return true
	}

	for _, v := range _withoutLoginConfig {
		if len(v) <= rpcNameLen && strings.Compare(rpcName[:len(v)], v) == 0 {
			return true
		}
	}

	return false
}

func IsUpdatesGetDifference(rpcName string) bool {
	rpcNameLen := len(rpcName)
	if rpcNameLen == 0 {
		return false
	}
	if rpcName[0] == '*' {
		rpcName = rpcName[1:]
		rpcNameLen = len(rpcName)
	}
	if rpcNameLen == 0 {
		return false
	}

	if len(_updatesGetDifference) <= rpcNameLen && strings.Compare(rpcName[:len(_updatesGetDifference)], _updatesGetDifference) == 0 {
		return true
	}

	return false
}

func IsUploadSession(rpcName string) bool {
	rpcNameLen := len(rpcName)
	if rpcNameLen == 0 {
		return false
	}
	if rpcName[0] == '*' {
		rpcName = rpcName[1:]
		rpcNameLen = len(rpcName)
	}
	if rpcNameLen == 0 {
		return false
	}

	if len(_upload) <= rpcNameLen && strings.Compare(rpcName[:len(_upload)], _upload) == 0 {
		return true
	}

	return false
}
