/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/21 11:56
 * @File : hashs.go
 */

package hash

import "sort"

// https://core.telegram.org/api/offsets#hash-generation
// hash = 0
// for id in ids:
// hash = (((hash * 0x4F25) & 0x7FFFFFFF) + id) & 0x7FFFFFFF

func HashIdList(ids []int32, withSort bool, desc bool) int32 {
	if withSort {
		sort.SliceStable(ids, func(i, j int) bool {
			if desc == true {
				return ids[i] <= ids[j]
			} else {
				return ids[i] >= ids[j]
			}
		})
	}

	hash := int32(0)
	b := int32(0x4F25)
	l := int32(0x7FFFFFFF)
	for _, id := range ids {
		hash = (((hash * b) & l) + id) & l
	}

	return hash
}

func HashId64ListNew(ids []int64) int64 {
	sort.SliceStable(ids, func(i, j int) bool {
		return ids[i] <= ids[j]
	})

	hash := int64(0)
	for _, v := range ids {
		hash ^= v >> 21
		hash ^= v << 35
		hash ^= v >> 4

		hash += v
	}

	return hash
}
