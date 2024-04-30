/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/2/23 10:04
 * @File : privacy_value_util.go
 */

package constants

import (
	"fmt"
	"novachat_engine/mtproto"
)

//  + TL_inputPrivacyValueAllowContacts
//  + TL_inputPrivacyValueAllowAll
//  + TL_inputPrivacyValueAllowUsers
//  + TL_inputPrivacyValueDisallowContacts
//  + TL_inputPrivacyValueDisallowAll
//  + TL_inputPrivacyValueDisallowUsers
//  + TL_inputPrivacyValueAllowChatParticipants
//  + TL_inputPrivacyValueDisallowChatParticipants

func ToPrivacyValueMTValue(selfUserId int64, p *mtproto.InputPrivacyRule) *mtproto.PrivacyRule {
	if p == nil {
		return nil
	}

	switch p.ClassName {
	case mtproto.ClassInputPrivacyValueAllowAll:
		return mtproto.NewTLPrivacyValueAllowAll(nil).To_PrivacyRule()
	case mtproto.ClassInputPrivacyValueDisallowAll:
		return mtproto.NewTLPrivacyValueDisallowAll(nil).To_PrivacyRule()
	case mtproto.ClassInputPrivacyValueAllowUsers:
		users := make([]int64, 0, len(p.Users))
		for _, v := range p.Users {
			if v.Peer != nil {
				users = append(users, PeerTypeFromUserIDType32(v.Peer.UserId).ToInt())
			} else {
				if v.ClassName == mtproto.ClassInputUserSelf {
					if selfUserId != 0 {
						users = append(users, selfUserId)
					}
				} else {
					if v.UserId != 0 {
						users = append(users, PeerTypeFromUserIDType32(v.UserId).ToInt())
					}
				}
			}
		}
		return mtproto.NewTLPrivacyValueAllowUsers(&mtproto.PrivacyRule{Users: ArrayInt64To32(users)}).To_PrivacyRule()
	case mtproto.ClassInputPrivacyValueDisallowUsers:
		users := make([]int32, 0, len(p.Users))
		for _, v := range p.Users {
			if v.Peer != nil {
				users = append(users, v.Peer.UserId)
			} else {
				if v.UserId != 0 {
					users = append(users, v.UserId)
				}
			}
		}
		return mtproto.NewTLPrivacyValueDisallowUsers(&mtproto.PrivacyRule{Chats: users}).To_PrivacyRule()
	case mtproto.ClassInputPrivacyValueAllowChatParticipants:
		return mtproto.NewTLPrivacyValueAllowChatParticipants(&mtproto.PrivacyRule{Chats: p.Chats}).To_PrivacyRule()
	case mtproto.ClassInputPrivacyValueDisallowChatParticipants:
		return mtproto.NewTLPrivacyValueDisallowChatParticipants(&mtproto.PrivacyRule{Chats: p.Chats}).To_PrivacyRule()
	case mtproto.ClassInputPrivacyValueAllowContacts:
		return mtproto.NewTLPrivacyValueAllowContacts(nil).To_PrivacyRule()
	case mtproto.ClassInputPrivacyValueDisallowContacts:
		return mtproto.NewTLPrivacyValueDisallowContacts(nil).To_PrivacyRule()
	default:
		panic(fmt.Sprintf("ToPrivacyValueMTValue error:%s", p.ClassName))
	}
}

func ToPrivacyValue(p *mtproto.InputPrivacyRule) OptionPrivacy {
	if p == nil {
		return OptionPrivacyAllowAll
	}

	switch p.ClassName {
	case mtproto.ClassInputPrivacyValueAllowAll:
		return OptionPrivacyAllowAll
	case mtproto.ClassInputPrivacyValueDisallowAll:
		return OptionPrivacyDisallowAll
	case mtproto.ClassInputPrivacyValueAllowUsers:
		return OptionPrivacyAllowUsers
	case mtproto.ClassInputPrivacyValueDisallowUsers:
		return OptionPrivacyDisallowUsers
	case mtproto.ClassInputPrivacyValueAllowChatParticipants:
		return OptionPrivacyAllowChatParticipants
	case mtproto.ClassInputPrivacyValueDisallowChatParticipants:
		return OptionPrivacyDisallowChatParticipants
	case mtproto.ClassInputPrivacyValueAllowContacts:
		return OptionPrivacyAllowAllContacts
	case mtproto.ClassInputPrivacyValueDisallowContacts:
		return OptionPrivacyDisallowContacts

	default:
		panic(fmt.Sprintf("ToPrivacyValue error:%s", p.ClassName))
	}
}

//func PrivacyToValue(p OptionPrivacy) *mtproto.PrivacyRule {
//if p == nil {
//	return OptionPrivacyAllowAll
//}
//
//switch p.ClassName {
//case mtproto.ClassInputPrivacyValueAllowAll:
//	return OptionPrivacyAllowAll
//case mtproto.ClassInputPrivacyValueDisallowAll:
//	return OptionPrivacyDisallowAll
//case mtproto.ClassInputPrivacyValueAllowUsers:
//	return OptionPrivacyAllowUsers
//case mtproto.ClassInputPrivacyValueDisallowUsers:
//	return OptionPrivacyDisallowUsers
//case mtproto.ClassInputPrivacyValueAllowChatParticipants:
//	return OptionPrivacyAllowChatParticipants
//case mtproto.ClassInputPrivacyValueDisallowChatParticipants:
//	return OptionPrivacyDisallowChatParticipants
//case mtproto.ClassInputPrivacyValueAllowContacts:
//	return OptionPrivacyAllowAllContacts
//case mtproto.ClassInputPrivacyValueDisallowContacts:
//	return OptionPrivacyDisallowContacts
//
//default:
//	panic(fmt.Sprintf("ToPrivacyValue error:%s", p.ClassName))
//}
//}
