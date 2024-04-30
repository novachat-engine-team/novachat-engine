/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/2/23 10:50
 * @File : privacy_sort.go
 */

package constants

//rules:<cmd:Cmd_InputPrivacyValueAllowUsers class_name:"TLInputPrivacyValueAllowUsers" >
//rules:<cmd:Cmd_InputPrivacyValueAllowChatParticipants class_name:"TLInputPrivacyValueAllowChatParticipants" chats:424 chats:364 chats:363 chats:359 >
//rules:<cmd:Cmd_InputPrivacyValueDisallowUsers class_name:"TLInputPrivacyValueDisallowUsers" users:<cmd:Cmd_InputUser class_name:"TLInputUser" user_id:100001022 access_hash:7206025280335552864 > >
//rules:<cmd:Cmd_InputPrivacyValueDisallowChatParticipants class_name:"TLInputPrivacyValueDisallowChatParticipants" >
//rules:<cmd:Cmd_InputPrivacyValueAllowContacts class_name:"TLInputPrivacyValueAllowContacts" >

func SortCompare(a, b OptionPrivacy) bool {
	return a <= b
}

func SortCompareInt32(a, b int32) bool {
	return a <= b
}
