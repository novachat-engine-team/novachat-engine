/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/10/13 10:31
 * @File : status.go
 */

package constants

import (
	"novachat_engine/mtproto"
	"time"
)

/*******************************************************************
// https://telegram.org/faq#q-can-i-hide-my-last-seen-time

Q: Can I hide my ‘last seen’ time?

You can choose who sees this info in [Privacy and Security](!https://telegram.org/blog/privacy-revolution) settings.

Remember that you won‘t see Last Seen timestamps for people with whom you don’t share your own.
You will, however, see an approximate last seen value.
This keeps stalkers away but makes it possible to understand whether a person is reachable over Telegram.
There are four possible approximate values:

- Last seen recently — covers anything between 1 second and 2-3 days
- Last seen within a week — between 2-3 and seven days
- Last seen within a month — between 6-7 days and a month
- Last seen a long time ago — more than a month (this is also always shown to blocked users)

Q: Who can see me ‘online’?

The last seen rules apply to your online status as well.
People can only see you online if you're sharing your last seen status with them.
There is one exception to this: people will be able to see you online for a brief period
when you send them a message in a one-on-one chat or in a group where you both are members.

********************************************************************/

func MakeUserStatus(lastSeenAt int64, showStatus bool, hide bool) *mtproto.UserStatus {
	now := time.Now().Unix()

	if hide == true {
		return mtproto.NewTLUserStatusEmpty(nil).To_UserStatus()
	}

	if showStatus {
		if now <= lastSeenAt+60 {
			status := mtproto.NewTLUserStatusOnline(&mtproto.UserStatus{
				Expires: int32(lastSeenAt + 60),
			})
			return status.To_UserStatus()
		} else {
			status := mtproto.NewTLUserStatusOffline(&mtproto.UserStatus{
				WasOnline: int32(lastSeenAt),
			})
			return status.To_UserStatus()
		}
	} else {
		if now-lastSeenAt >= 60*60*24*30 {
			return mtproto.NewTLUserStatusRecently(nil).To_UserStatus()
		} else if now-lastSeenAt >= 60*60*24*7 {
			return mtproto.NewTLUserStatusLastMonth(nil).To_UserStatus()
		} else if now-lastSeenAt >= 60*60*24*3 {
			return mtproto.NewTLUserStatusLastWeek(nil).To_UserStatus()
		} else {
			return mtproto.NewTLUserStatusRecently(nil).To_UserStatus()
		}
	}
}
