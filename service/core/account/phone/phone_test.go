/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/28 14:04
 * @File : phone_test.go
 */

package phone

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"testing"
)

func TestVirtualPhone(t *testing.T) {
	t.Log(IsVirtualPhone("++2000000000000000"))
	t.Log(IsVirtualPhone("+2000000000000000"))
	t.Log(IsVirtualPhone("2000000000000000"))
	t.Log(IsVirtualPhone("+86 15017910600"))

	s, _ := Parse("+86 15017910600")

	aa := "wodeshi  %s ddfs"
	fmt.Println(fmt.Sprintf(aa, "tmd"))

	t.Log(s.GetRegionCode())
	t.Log(phonenumbers.GetSupportedRegions())
}