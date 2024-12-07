/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/2/23 17:41
 * @File :
 *
 */

package users

import (
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
	data_contact "novachat_engine/service/data/contact"
	"novachat_engine/service/data/users"
)

func UserCoreContactUser(u *mtproto.User, contact bool, mutual bool, contactInfo *data_contact.Contact) *mtproto.User {
	u.Self = false

	if mutual == true {
		u.Contact = true
		u.MutualContact = true
	} else {
		u.Contact = contact
		u.MutualContact = false
	}

	if contact && contactInfo != nil && len(contactInfo.FirstName) > 0 && !contactInfo.Deleted {
		u.FirstName = contactInfo.FirstName
		u.LastName = contactInfo.LastName
	}
	return u
}

func UserCoreSelfUsers(u *mtproto.User) *mtproto.User {
	u.Self = true
	u.Contact = true
	u.MutualContact = true
	return u
}

func UserCore2Users(u *data_users.Users) *mtproto.User {
	if u == nil || u.Id == 0 {
		return mtproto.NewTLUserEmpty(nil).To_User()
	} else {
		//TODO:Coder
		//different layers

		var rr []*mtproto.RestrictionReason
		var r string
		if u.GetRestricted() {
			r = u.RestrictionReason
			rr = []*mtproto.RestrictionReason{mtproto.NewTLRestrictionReason(&mtproto.RestrictionReason{
				Platform: "",
				Reason:   r,
				Text:     "",
			}).To_RestrictionReason()}
		}
		return mtproto.NewTLUser(&mtproto.User{
			Id:                           constants.PeerTypeFromUserIDType(u.Id).ToInt32(),
			Self:                         false,
			Contact:                      false,
			MutualContact:                false,
			Deleted:                      u.GetDeleted(),
			Bot:                          u.GetBot(),
			BotChatHistory:               false,
			BotNochats:                   false,
			Verified:                     u.GetVerified(),
			Restricted:                   u.GetRestricted(),
			Min:                          false,
			BotInlineGeo:                 false,
			AccessHash:                   u.AccessHash,
			FirstName:                    u.FirstName,
			LastName:                     u.LastName,
			Username:                     u.GetUsername(),
			Phone:                        u.GetPhone(),
			Photo:                        nil,
			Status:                       nil,
			BotInfoVersion:               0,
			RestrictionReason2E13F4C371:  r,
			BotInlinePlaceholder:         "",
			LangCode:                     "",
			Support:                      u.GetSupport(),
			Scam:                         false,
			ApplyMinPhoto:                false,
			RestrictionReason938458C1117: rr,
		}).To_User()
	}
}
