/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/2/23 10:54
 * @File : privacy_sort_test.go
 */

package constants

import (
	"sort"
	"testing"
)

func TestSortCompare(t *testing.T) {
	v := []OptionPrivacy{
		OptionPrivacyAllowAll,
		OptionPrivacyDisallowUsers,
		OptionPrivacyAllowUsers,
	}

	sort.SliceStable(v, func(i, j int) bool {
		return SortCompare(v[i], v[j])
	})

	t.Logf("%+v", v)
}
