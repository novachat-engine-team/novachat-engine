/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/9/9 16:55
 * @File : banned_right_test.go
 */

package banned_right

import "testing"

func TestChannelAdminRightToChatAdminRight(t *testing.T) {
	right := int32(0b01)
	t.Logf("%b", right)
	r := RightsToChatBannedRight(right, 0)
	t.Log(r.GetViewMessages())


	t.Logf("%v", checkRight(right, BannedRight_ViewMessages))


	var right1 int32
	setRight(&right1, r.GetViewMessages(), BannedRight_ViewMessages)
	t.Logf("%b", right1)
}
