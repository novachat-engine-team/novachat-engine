/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/2 13:41
 * @File : mongo_test.go
 */

package mgo

import "testing"

type A struct {
	A string `bson:"a"`
	B []int32 `bson:"b"`
	C int64 `bson:"c"`
	D map[int32][]string `bson:"d_map"`
}

func TestToBsonM(t *testing.T) {
	a := A {
		A: "sssss",
		B: []int32{1, 2,3,45},
		C: 1000,
		D: map[int32][]string{
			1: []string{"1111","1111","1111","1111","1111","1111"},
			2: []string{"2222","2222","2222","2222","2222","2222"},
		},
	}

	v := ToBsonM(&a)
	t.Logf("%+v", v)

	v = ToBsonM(a)
	t.Logf("%+v", v)


	vList := ToBsonMList([]A{a})
	t.Logf("%v", vList)
}

