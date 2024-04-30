/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/2/23 16:04
 * @File : privacy_check.go
 */

package privacy

import (
	"novachat_engine/service/constants"
	data_privacy "novachat_engine/service/data/privacy"
)

func check(id int64, l []int64) bool {
	if id == 0 {
		return true
	}

	for _, v := range l {
		if id == v {
			return false
		}
	}
	return true
}

type checkType int

const (
	checkTypeNone   checkType = 0
	checkTypeActive checkType = 1
	checkTypeNative checkType = 2
)

func CheckPrivacy(userId int64, chatId int64, v []*data_privacy.PrivacyOption, bContacts bool) bool {
	if userId == 0 && chatId == 0 {
		return true
	}

	allowAll := checkTypeNone
	contactsAll := checkTypeNone
	allowUser := checkTypeNone
	allowChat := checkTypeNone

	for _, vv := range v {
		switch constants.OptionPrivacy(vv.PrivacyKey) {
		case constants.OptionPrivacyAllowAll:
			allowAll = checkTypeActive
			break
		case constants.OptionPrivacyDisallowAll:
			allowAll = checkTypeNative
			break
		case constants.OptionPrivacyAllowAllContacts:
			contactsAll = checkTypeActive
			break
		case constants.OptionPrivacyDisallowContacts:
			contactsAll = checkTypeNative
			break
		case constants.OptionPrivacyAllowUsers:
			if userId != 0 {
				if !check(userId, vv.AllowList) {
					allowUser = checkTypeActive
				} else {
					allowUser = checkTypeNative
				}
			}
			break
		case constants.OptionPrivacyAllowChatParticipants:
			if chatId != 0 {
				if !check(chatId, vv.AllowList) {
					allowChat = checkTypeActive
				} else {
					allowChat = checkTypeNative
				}
			}
			break
		case constants.OptionPrivacyDisallowUsers:
			if userId != 0 {
				if !check(userId, vv.DisallowList) {
					allowUser = checkTypeNative
				} else {
					allowUser = checkTypeActive
				}
			}
			break
		case constants.OptionPrivacyDisallowChatParticipants:
			if chatId != 0 {
				if !check(chatId, vv.DisallowList) {
					allowChat = checkTypeNative
				} else {
					allowChat = checkTypeActive
				}
			}
			break
		}
	}

	// false 不允许, true 允许
	if contactsAll == checkTypeActive {
		if bContacts == true {
			return true
		}
		return false
	} else if contactsAll == checkTypeNative {
		if bContacts == false {
			return false
		}
		return true
	}
	if allowChat == checkTypeActive {
		return true
	} else if allowChat == checkTypeNative {
		return false
	}
	if allowUser == checkTypeActive {
		return true
	} else if allowUser == checkTypeNative {
		return false
	}
	if allowAll == checkTypeActive {
		return true
	} else if allowAll == checkTypeNative {
		return false
	}
	return true
}
