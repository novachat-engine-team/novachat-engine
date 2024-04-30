/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/5/24 14:07
 * @File : strings.go
 */

package duplicates

func UniqueString(s []string) []string {
	inResult := make(map[string]bool)
	var result = make([]string, 0, len(s))
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
