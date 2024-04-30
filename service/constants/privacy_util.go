/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/2/23 9:54
 * @File : privacy_util.go
 */

package constants

import (
	"fmt"
	"novachat_engine/mtproto"
)

func ToPrivacyKeyMTKey(p *mtproto.InputPrivacyKey) *mtproto.PrivacyKey {
	if p == nil {
		return nil
	}
	switch p.ClassName {
	case mtproto.ClassInputPrivacyKeyStatusTimestamp:
		return mtproto.NewTLPrivacyKeyStatusTimestamp(&mtproto.PrivacyKey{}).To_PrivacyKey()
	case mtproto.ClassInputPrivacyKeyChatInvite:
		return mtproto.NewTLPrivacyKeyChatInvite(&mtproto.PrivacyKey{}).To_PrivacyKey()
	case mtproto.ClassInputPrivacyKeyPhoneCall:
		return mtproto.NewTLPrivacyKeyPhoneCall(&mtproto.PrivacyKey{}).To_PrivacyKey()
	case mtproto.ClassInputPrivacyKeyPhoneP2P:
		return mtproto.NewTLPrivacyKeyPhoneNumber(&mtproto.PrivacyKey{}).To_PrivacyKey()
	case mtproto.ClassInputPrivacyKeyForwards:
		return mtproto.NewTLPrivacyKeyForwards(&mtproto.PrivacyKey{}).To_PrivacyKey()
	case mtproto.ClassInputPrivacyKeyProfilePhoto:
		return mtproto.NewTLPrivacyKeyProfilePhoto(&mtproto.PrivacyKey{}).To_PrivacyKey()
	case mtproto.ClassInputPrivacyKeyPhoneNumber:
		return mtproto.NewTLPrivacyKeyPhoneNumber(&mtproto.PrivacyKey{}).To_PrivacyKey()
	case mtproto.ClassInputPrivacyKeyAddedByPhone:
		return mtproto.NewTLPrivacyKeyAddedByPhone(&mtproto.PrivacyKey{}).To_PrivacyKey()
	default:
		panic(fmt.Sprintf("ToPrivacyKeyMTKey error:%s", p.ClassName))
	}
}

func ToPrivacyKey(p *mtproto.InputPrivacyKey) KeyPrivacy {
	if p == nil {
		return KeyPrivacyNone
	}
	switch p.ClassName {
	case mtproto.ClassInputPrivacyKeyStatusTimestamp:
		return KeyPrivacyUserStatus
	case mtproto.ClassInputPrivacyKeyChatInvite:
		return KeyPrivacyChatInvite
	case mtproto.ClassInputPrivacyKeyPhoneCall:
		return KeyPrivacyPhoneCall
	case mtproto.ClassInputPrivacyKeyPhoneP2P:
		return KeyPrivacyPhoneP2P
	case mtproto.ClassInputPrivacyKeyForwards:
		return KeyPrivacyForwards
	case mtproto.ClassInputPrivacyKeyProfilePhoto:
		return KeyPrivacyProfilePhoto
	case mtproto.ClassInputPrivacyKeyPhoneNumber:
		return KeyPrivacyPhoneNumber
	case mtproto.ClassInputPrivacyKeyAddedByPhone:
		return KeyPrivacyAddByPhone
	default:
		panic(fmt.Sprintf("ToPrivacyKey error:%s", p.ClassName))
	}
}
