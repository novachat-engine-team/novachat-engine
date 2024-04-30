/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/10/19 19:39
 * @File : support.go
 */

package support

var _cacheSupport = map[int64]int{}

func IsSupportUser(userId int64) bool {
	if _cacheSupport != nil {
		_, exists := _cacheSupport[userId]
		return exists
	}
	return false
}
