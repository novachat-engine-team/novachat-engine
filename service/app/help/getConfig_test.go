/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/15 14:25
 * @File : getConfig_test.go
 */

package help

import "testing"

func TestGetConfig(t *testing.T) {
	err := ParseConfig("../../../server/biz_server/")

	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Logf("%+v", GetConfig())
	}
}
