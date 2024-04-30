/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/5/24 14:07
 * @File : ints.go
 */

package duplicates

func UniqueInt(s []int) []int {
	inResult := make(map[int]bool)
	var result = make([]int, 0, len(s))
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func UniqueInt32(s []int32) []int32 {
	inResult := make(map[int32]bool)
	var result = make([]int32, 0, len(s))
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func UniqueInt64(s []int64) []int64 {
	inResult := make(map[int64]bool)
	var result = make([]int64, 0, len(s))
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}