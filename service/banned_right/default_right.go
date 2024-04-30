/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/5 15:31
 * @File : default_right.go
 */

package banned_right

//const (
//	DefaultRight int32 = 0
//	DefaultAdminRight int32 = 0
//	DefaultCreatorRight int32 = 0
//)

func DefaultRight() int32 {
	return 0
}

func DefaultAdminRight() int32 {
	return 0
}

func DefaultCreatorRight() int32 {
	return 0
}

func DisableRight() int32 {
	return -1
}

func DisableAdminRight() int32 {
	return -1
}

