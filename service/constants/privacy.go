/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/2/14 16:57
 * @File : privacy.go
 */

package constants

// Key:
//  + TL_inputPrivacyKeyStatusTimestamp
//  + TL_inputPrivacyKeyChatInvite
//  + TL_inputPrivacyKeyPhoneCall
//  + TL_inputPrivacyKeyPhoneP2P
//  + TL_inputPrivacyKeyForwards
//  + TL_inputPrivacyKeyProfilePhoto
//  + TL_inputPrivacyKeyPhoneNumber
//  + TL_inputPrivacyKeyAddedByPhone
type KeyPrivacy int32

const (
	KeyPrivacyNone         KeyPrivacy = 0
	KeyPrivacyUserStatus   KeyPrivacy = 1
	KeyPrivacyChatInvite   KeyPrivacy = 2
	KeyPrivacyPhoneCall    KeyPrivacy = 3
	KeyPrivacyPhoneP2P     KeyPrivacy = 4
	KeyPrivacyForwards     KeyPrivacy = 5
	KeyPrivacyProfilePhoto KeyPrivacy = 6
	KeyPrivacyPhoneNumber  KeyPrivacy = 7
	KeyPrivacyAddByPhone   KeyPrivacy = 8
)

func (m KeyPrivacy) Int32() int32 {
	return int32(m)
}

// Option:
//  + TL_inputPrivacyValueAllowContacts
//  + TL_inputPrivacyValueAllowAll
//  + TL_inputPrivacyValueAllowUsers
//  + TL_inputPrivacyValueDisallowContacts
//  + TL_inputPrivacyValueDisallowAll
//  + TL_inputPrivacyValueDisallowUsers
//  + TL_inputPrivacyValueAllowChatParticipants
//  + TL_inputPrivacyValueDisallowChatParticipants
type OptionPrivacy int32

const (
	OptionPrivacyAllowNone                OptionPrivacy = 0
	OptionPrivacyAllowUsers               OptionPrivacy = 1
	OptionPrivacyAllowChatParticipants    OptionPrivacy = 2
	OptionPrivacyAllowAllContacts         OptionPrivacy = 3
	OptionPrivacyDisallowUsers            OptionPrivacy = 4
	OptionPrivacyDisallowChatParticipants OptionPrivacy = 5
	OptionPrivacyDisallowContacts         OptionPrivacy = 6
	OptionPrivacyDisallowAll              OptionPrivacy = 7
	OptionPrivacyAllowAll                 OptionPrivacy = 8
)

func (m OptionPrivacy) Int32() int32 {
	return int32(m)
}
