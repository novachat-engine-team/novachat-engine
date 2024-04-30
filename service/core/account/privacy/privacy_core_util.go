/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/2/23 17:41
 * @File :
 *
 */

package privacy

import (
	"novachat_engine/service/constants"
	"novachat_engine/service/data/privacy"
)

func FillPrivacy(p *data_privacy.Privacy, key int32, op []*data_privacy.PrivacyOption) string {
	switch key {
	case constants.KeyPrivacyUserStatus.Int32():
		p.UserStatus = op
		return "UserStatus"
	case constants.KeyPrivacyChatInvite.Int32():
		p.ChatInvite = op
		return "ChatInvite"
	case constants.KeyPrivacyPhoneCall.Int32():
		p.PhoneCall = op
		return "PhoneCall"
	case constants.KeyPrivacyPhoneP2P.Int32():
		p.PhoneP2P = op
		return "PhoneP2P"
	case constants.KeyPrivacyForwards.Int32():
		p.Forwards = op
		return "Forwards"
	case constants.KeyPrivacyProfilePhoto.Int32():
		p.Profile = op
		return "Profile"
	case constants.KeyPrivacyPhoneNumber.Int32():
		p.PhoneNumber = op
		return "PhoneNumber"
	case constants.KeyPrivacyAddByPhone.Int32():
		p.AddByPhone = op
		return "AddByPhone"
	case constants.KeyPrivacyNone.Int32():
		fallthrough
	default:
		p.Global = op
		return "Global"
	}
}

func TokenPrivacy(p *data_privacy.Privacy, key int32) []*data_privacy.PrivacyOption {
	switch key {
	case constants.KeyPrivacyUserStatus.Int32():
		return p.UserStatus
	case constants.KeyPrivacyChatInvite.Int32():
		return p.ChatInvite
	case constants.KeyPrivacyPhoneCall.Int32():
		return p.PhoneCall
	case constants.KeyPrivacyPhoneP2P.Int32():
		return p.PhoneP2P
	case constants.KeyPrivacyForwards.Int32():
		return p.Forwards
	case constants.KeyPrivacyProfilePhoto.Int32():
		return p.Profile
	case constants.KeyPrivacyPhoneNumber.Int32():
		return p.PhoneNumber
	case constants.KeyPrivacyAddByPhone.Int32():
		return p.AddByPhone
	case constants.KeyPrivacyNone.Int32():
		fallthrough
	default:
		return p.Global
	}
}
