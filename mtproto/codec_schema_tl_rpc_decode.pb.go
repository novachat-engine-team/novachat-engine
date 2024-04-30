/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2023/06/12 18:03
 * @File : codec_schema_tl_rpc_decode.pb.go
 * @Desc :
 *
 */

package mtproto

import (
	"novachat_engine/pkg/log"
)

func (m *TLAccountAcceptAuthorization) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountAcceptAuthorization:
		m.BotId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Scope = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PublicKey = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l4 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.ValueHashes = make([]*SecureValueHash, l4)
		for i := int32(0); i < l4; i++ {
			m.ValueHashes[i] = &SecureValueHash{}
			err = (*m.ValueHashes[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}
		m5 := &SecureCredentialsEncrypted{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Credentials = m5

	default:
		log.Errorf("AccountAcceptAuthorization Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountCancelPasswordEmail) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountCancelPasswordEmail:
		_ = m

	default:
		log.Errorf("AccountCancelPasswordEmail Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountChangePhone) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountChangePhone:
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCodeHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountChangePhone Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountCheckUsername) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountCheckUsername:
		m.Username = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountCheckUsername Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountConfirmPasswordEmail) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountConfirmPasswordEmail:
		m.Code = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountConfirmPasswordEmail Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountConfirmPhone) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountConfirmPhone:
		m.PhoneCodeHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountConfirmPhone Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountCreateTheme) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountCreateTheme:
		m.Slug = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Title = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &InputDocument{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Document = m3

	case Cmd_AccountCreateTheme8432c21f:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Slug = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Title = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 2)) != 0 {
			m4 := &InputDocument{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Document = m4
		}
		if (flags & (1 << 3)) != 0 {
			m5 := &InputThemeSettings{}
			err = m5.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Settings = m5
		}

	default:
		log.Errorf("AccountCreateTheme Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountDeleteAccount) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountDeleteAccount:
		m.Reason = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountDeleteAccount Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountDeleteSecureValue) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountDeleteSecureValue:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Types = make([]*SecureValueType, l1)
		for i := int32(0); i < l1; i++ {
			m.Types[i] = &SecureValueType{}
			err = (*m.Types[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("AccountDeleteSecureValue Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountFinishTakeoutSession) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountFinishTakeoutSession:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Success = flags&(1<<0) != 0

	default:
		log.Errorf("AccountFinishTakeoutSession Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetAccountTTL) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetAccountTTL:
		_ = m

	default:
		log.Errorf("AccountGetAccountTTL Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetAllSecureValues) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetAllSecureValues:
		_ = m

	default:
		log.Errorf("AccountGetAllSecureValues Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetAuthorizationForm) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetAuthorizationForm:
		m.BotId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Scope = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PublicKey = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountGetAuthorizationForm Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetAuthorizations) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetAuthorizations:
		_ = m

	default:
		log.Errorf("AccountGetAuthorizations Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetAutoDownloadSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetAutoDownloadSettings:
		_ = m

	default:
		log.Errorf("AccountGetAutoDownloadSettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetContactSignUpNotification) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetContactSignUpNotification:
		_ = m

	default:
		log.Errorf("AccountGetContactSignUpNotification Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetContentSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetContentSettings:
		_ = m

	default:
		log.Errorf("AccountGetContentSettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetGlobalPrivacySettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetGlobalPrivacySettings:
		_ = m

	default:
		log.Errorf("AccountGetGlobalPrivacySettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetMultiWallPapers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetMultiWallPapers:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Wallpapers = make([]*InputWallPaper, l1)
		for i := int32(0); i < l1; i++ {
			m.Wallpapers[i] = &InputWallPaper{}
			err = (*m.Wallpapers[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("AccountGetMultiWallPapers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetNotifyExceptions) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetNotifyExceptions:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.CompareSound = flags&(1<<1) != 0
		if (flags & (1 << 0)) != 0 {
			m3 := &InputNotifyPeer{}
			err = m3.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Peer = m3
		}

	default:
		log.Errorf("AccountGetNotifyExceptions Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetNotifySettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetNotifySettings:
		m1 := &InputNotifyPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("AccountGetNotifySettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetPassword) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetPassword:
		_ = m

	default:
		log.Errorf("AccountGetPassword Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetPasswordSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetPasswordSettings:
		m.CurrentPasswordHash = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_AccountGetPasswordSettings9cd4eaf9:
		m1 := &InputCheckPasswordSRP{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Password = m1

	default:
		log.Errorf("AccountGetPasswordSettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetPrivacy) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetPrivacy:
		m1 := &InputPrivacyKey{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Key = m1

	default:
		log.Errorf("AccountGetPrivacy Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetSecureValue) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetSecureValue:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Types = make([]*SecureValueType, l1)
		for i := int32(0); i < l1; i++ {
			m.Types[i] = &SecureValueType{}
			err = (*m.Types[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("AccountGetSecureValue Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetTheme) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetTheme:
		m.Format = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &InputTheme{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Theme = m2
		m.DocumentId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountGetTheme Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetThemes) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetThemes:
		m.Format = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountGetThemes Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetTmpPassword) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetTmpPassword:
		m.PasswordHash = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Period = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_AccountGetTmpPassword449e0b51:
		m1 := &InputCheckPasswordSRP{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Password = m1
		m.Period = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountGetTmpPassword Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetWallPaper) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetWallPaper:
		m1 := &InputWallPaper{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Wallpaper = m1

	default:
		log.Errorf("AccountGetWallPaper Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetWallPapers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetWallPapers:
		_ = m

	case Cmd_AccountGetWallPapersaabb1763:
		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountGetWallPapers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountGetWebAuthorizations) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountGetWebAuthorizations:
		_ = m

	default:
		log.Errorf("AccountGetWebAuthorizations Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountInitTakeoutSession) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountInitTakeoutSession:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Contacts = flags&(1<<0) != 0
		m.MessageUsers = flags&(1<<1) != 0
		m.MessageChats = flags&(1<<2) != 0
		m.MessageMegagroups = flags&(1<<3) != 0
		m.MessageChannels = flags&(1<<4) != 0
		m.Files = flags&(1<<5) != 0
		if (flags & (1 << 5)) != 0 {
			m.FileMaxSize = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	default:
		log.Errorf("AccountInitTakeoutSession Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountInstallTheme) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountInstallTheme:
		m.Format = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &InputTheme{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Theme = m2

	case Cmd_AccountInstallTheme7ae43737:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Dark = flags&(1<<0) != 0
		if (flags & (1 << 1)) != 0 {
			m.Format = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 1)) != 0 {
			m4 := &InputTheme{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Theme = m4
		}

	default:
		log.Errorf("AccountInstallTheme Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountInstallWallPaper) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountInstallWallPaper:
		m1 := &InputWallPaper{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Wallpaper = m1

	case Cmd_AccountInstallWallPaperfeed5769:
		m1 := &InputWallPaper{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Wallpaper = m1
		m2 := &WallPaperSettings{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Settings = m2

	default:
		log.Errorf("AccountInstallWallPaper Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountRegisterDevice) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountRegisterDevice:
		m.TokenType = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Token = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_AccountRegisterDevice446c712c:
		m.TokenType = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Token = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.DeviceModel = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.SystemVersion = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AppVersion = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m6 := &Bool{}
		err = m6.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.AppSandbox = m6
		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_AccountRegisterDevice5cbea590:
		m.TokenType = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Token = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &Bool{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.AppSandbox = m3
		m.Secret = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OtherUids = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_AccountRegisterDevice68976c6f:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.NoMuted = flags&(1<<0) != 0
		m.TokenType = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Token = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m5 := &Bool{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.AppSandbox = m5
		m.Secret = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OtherUids = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountRegisterDevice Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountReportPeer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountReportPeer:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m2 := &ReportReason{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Reason = m2

	default:
		log.Errorf("AccountReportPeer Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountResendPasswordEmail) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountResendPasswordEmail:
		_ = m

	default:
		log.Errorf("AccountResendPasswordEmail Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountResetAuthorization) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountResetAuthorization:
		m.Hash = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountResetAuthorization Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountResetNotifySettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountResetNotifySettings:
		_ = m

	default:
		log.Errorf("AccountResetNotifySettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountResetWallPapers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountResetWallPapers:
		_ = m

	default:
		log.Errorf("AccountResetWallPapers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountResetWebAuthorization) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountResetWebAuthorization:
		m.Hash = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountResetWebAuthorization Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountResetWebAuthorizations) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountResetWebAuthorizations:
		_ = m

	default:
		log.Errorf("AccountResetWebAuthorizations Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSaveAutoDownloadSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSaveAutoDownloadSettings:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Low = flags&(1<<0) != 0
		m.High = flags&(1<<1) != 0
		m4 := &AutoDownloadSettings{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Settings = m4

	default:
		log.Errorf("AccountSaveAutoDownloadSettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSaveSecureValue) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSaveSecureValue:
		m1 := &InputSecureValue{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Value = m1
		m.SecureSecretId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountSaveSecureValue Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSaveTheme) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSaveTheme:
		m1 := &InputTheme{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Theme = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Unsave = m2

	default:
		log.Errorf("AccountSaveTheme Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSaveWallPaper) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSaveWallPaper:
		m1 := &InputWallPaper{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Wallpaper = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Unsave = m2

	case Cmd_AccountSaveWallPaper6c5a5b37:
		m1 := &InputWallPaper{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Wallpaper = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Unsave = m2
		m3 := &WallPaperSettings{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Settings = m3

	default:
		log.Errorf("AccountSaveWallPaper Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSendChangePhoneCode) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSendChangePhoneCode:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.AllowFlashcall = flags&(1<<0) != 0
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m4 := &Bool{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.CurrentNumber = m4
		}

	case Cmd_AccountSendChangePhoneCode82574ae5:
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &CodeSettings{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Settings = m2

	default:
		log.Errorf("AccountSendChangePhoneCode Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSendConfirmPhoneCode) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSendConfirmPhoneCode:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.AllowFlashcall = flags&(1<<0) != 0
		m.Hash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m4 := &Bool{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.CurrentNumber = m4
		}

	case Cmd_AccountSendConfirmPhoneCode1b3faa88:
		m.Hash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &CodeSettings{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Settings = m2

	default:
		log.Errorf("AccountSendConfirmPhoneCode Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSendVerifyEmailCode) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSendVerifyEmailCode:
		m.Email = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountSendVerifyEmailCode Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSendVerifyPhoneCode) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSendVerifyPhoneCode:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.AllowFlashcall = flags&(1<<0) != 0
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m4 := &Bool{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.CurrentNumber = m4
		}

	case Cmd_AccountSendVerifyPhoneCodea5a356f9:
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &CodeSettings{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Settings = m2

	default:
		log.Errorf("AccountSendVerifyPhoneCode Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSetAccountTTL) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSetAccountTTL:
		m1 := &AccountDaysTTL{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Ttl = m1

	default:
		log.Errorf("AccountSetAccountTTL Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSetContactSignUpNotification) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSetContactSignUpNotification:
		m1 := &Bool{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Silent = m1

	default:
		log.Errorf("AccountSetContactSignUpNotification Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSetContentSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSetContentSettings:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.SensitiveEnabled = flags&(1<<0) != 0

	default:
		log.Errorf("AccountSetContentSettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSetGlobalPrivacySettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSetGlobalPrivacySettings:
		m1 := &GlobalPrivacySettings{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Settings = m1

	default:
		log.Errorf("AccountSetGlobalPrivacySettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountSetPrivacy) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountSetPrivacy:
		m1 := &InputPrivacyKey{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Key = m1
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l2 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Rules = make([]*InputPrivacyRule, l2)
		for i := int32(0); i < l2; i++ {
			m.Rules[i] = &InputPrivacyRule{}
			err = (*m.Rules[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("AccountSetPrivacy Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountUnregisterDevice) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountUnregisterDevice:
		m.TokenType = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Token = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_AccountUnregisterDevice3076c4bf:
		m.TokenType = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Token = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OtherUids = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountUnregisterDevice Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountUpdateDeviceLocked) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountUpdateDeviceLocked:
		m.Period = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountUpdateDeviceLocked Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountUpdateNotifySettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountUpdateNotifySettings:
		m1 := &InputNotifyPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m2 := &InputPeerNotifySettings{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Settings = m2

	default:
		log.Errorf("AccountUpdateNotifySettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountUpdatePasswordSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountUpdatePasswordSettings:
		m.CurrentPasswordHash = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &Account_PasswordInputSettings{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.NewSettings = m2

	case Cmd_AccountUpdatePasswordSettingsa59b102f:
		m1 := &InputCheckPasswordSRP{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Password = m1
		m2 := &Account_PasswordInputSettings{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.NewSettings = m2

	default:
		log.Errorf("AccountUpdatePasswordSettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountUpdateProfile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountUpdateProfile:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		if (flags & (1 << 0)) != 0 {
			m.FirstName = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.LastName = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 2)) != 0 {
			m.About = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	default:
		log.Errorf("AccountUpdateProfile Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountUpdateStatus) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountUpdateStatus:
		m1 := &Bool{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Offline = m1

	default:
		log.Errorf("AccountUpdateStatus Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountUpdateTheme) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountUpdateTheme:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputTheme{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Theme = m2
		if (flags & (1 << 0)) != 0 {
			m.Slug = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.Title = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 2)) != 0 {
			m5 := &InputDocument{}
			err = m5.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Document = m5
		}

	case Cmd_AccountUpdateTheme3b8ea202:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Format = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &InputTheme{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Theme = m3
		if (flags & (1 << 0)) != 0 {
			m.Slug = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.Title = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 2)) != 0 {
			m6 := &InputDocument{}
			err = m6.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Document = m6
		}

	case Cmd_AccountUpdateTheme5cb367d5:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Format = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &InputTheme{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Theme = m3
		if (flags & (1 << 0)) != 0 {
			m.Slug = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.Title = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 2)) != 0 {
			m6 := &InputDocument{}
			err = m6.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Document = m6
		}
		if (flags & (1 << 3)) != 0 {
			m7 := &InputThemeSettings{}
			err = m7.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Settings = m7
		}

	default:
		log.Errorf("AccountUpdateTheme Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountUpdateUsername) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountUpdateUsername:
		m.Username = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountUpdateUsername Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountUploadTheme) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountUploadTheme:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputFile{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.File = m2
		if (flags & (1 << 0)) != 0 {
			m3 := &InputFile{}
			err = m3.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Thumb = m3
		}
		m.FileName = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MimeType = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountUploadTheme Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountUploadWallPaper) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountUploadWallPaper:
		m1 := &InputFile{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.File = m1
		m.MimeType = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_AccountUploadWallPaperdd853661:
		m1 := &InputFile{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.File = m1
		m.MimeType = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &WallPaperSettings{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Settings = m3

	default:
		log.Errorf("AccountUploadWallPaper Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountVerifyEmail) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountVerifyEmail:
		m.Email = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Code = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountVerifyEmail Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAccountVerifyPhone) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AccountVerifyPhone:
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCodeHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AccountVerifyPhone Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthAcceptLoginToken) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthAcceptLoginToken:
		m.Token = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthAcceptLoginToken Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthBindTempAuthKey) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthBindTempAuthKey:
		m.PermAuthKeyId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Nonce = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ExpiresAt = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.EncryptedMessage = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthBindTempAuthKey Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthCancelCode) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthCancelCode:
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCodeHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthCancelCode Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthCheckPassword) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthCheckPassword:
		m.PasswordHash = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_AuthCheckPasswordd18b4d16:
		m1 := &InputCheckPasswordSRP{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Password = m1

	default:
		log.Errorf("AuthCheckPassword Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthCheckPhone) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthCheckPhone:
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthCheckPhone Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthDropTempAuthKeys) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthDropTempAuthKeys:
		m.ExceptAuthKeys = dBuf.VectorLong()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthDropTempAuthKeys Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthExportAuthorization) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthExportAuthorization:
		m.DcId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthExportAuthorization Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthExportLoginToken) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthExportLoginToken:
		m.ApiId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ApiHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ExceptIds = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthExportLoginToken Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthImportAuthorization) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthImportAuthorization:
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Bytes = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthImportAuthorization Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthImportBotAuthorization) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthImportBotAuthorization:
		m.Flags = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ApiId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ApiHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.BotAuthToken = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthImportBotAuthorization Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthImportLoginToken) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthImportLoginToken:
		m.Token = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthImportLoginToken Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthLogOut) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthLogOut:
		_ = m

	default:
		log.Errorf("AuthLogOut Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthRecoverPassword) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthRecoverPassword:
		m.Code = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthRecoverPassword Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthRequestPasswordRecovery) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthRequestPasswordRecovery:
		_ = m

	default:
		log.Errorf("AuthRequestPasswordRecovery Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthResendCode) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthResendCode:
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCodeHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthResendCode Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthResetAuthorizations) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthResetAuthorizations:
		_ = m

	default:
		log.Errorf("AuthResetAuthorizations Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthSendCode) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthSendCode:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.AllowFlashcall = flags&(1<<0) != 0
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m4 := &Bool{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.CurrentNumber = m4
		}
		m.ApiId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ApiHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_AuthSendCodea677244f:
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ApiId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ApiHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m4 := &CodeSettings{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Settings = m4

	case Cmd_AuthSendCodeccfd70cf:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.AllowFlashcall = flags&(1<<0) != 0
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m4 := &Bool{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.CurrentNumber = m4
		}
		m.ApiId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ApiHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthSendCode Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthSendInvites) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthSendInvites:
		m.PhoneNumbers = dBuf.VectorString()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Message = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthSendInvites Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthSignIn) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthSignIn:
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCodeHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthSignIn Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLAuthSignUp) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_AuthSignUp:
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCodeHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.FirstName = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.LastName = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_AuthSignUp80eee427:
		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneCodeHash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.FirstName = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.LastName = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("AuthSignUp Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLBotsAnswerWebhookJSONQuery) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_BotsAnswerWebhookJSONQuery:
		m.QueryId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &DataJSON{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Data = m2

	default:
		log.Errorf("BotsAnswerWebhookJSONQuery Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLBotsSendCustomRequest) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_BotsSendCustomRequest:
		m.CustomMethod = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &DataJSON{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Params = m2

	default:
		log.Errorf("BotsSendCustomRequest Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLBotsSetBotCommands) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_BotsSetBotCommands:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Commands = make([]*BotCommand, l1)
		for i := int32(0); i < l1; i++ {
			m.Commands[i] = &BotCommand{}
			err = (*m.Commands[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("BotsSetBotCommands Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsCheckUsername) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsCheckUsername:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.Username = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsCheckUsername Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsCreateChannel) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsCreateChannel:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Broadcast = flags&(1<<0) != 0
		m.Megagroup = flags&(1<<1) != 0
		m.Title = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.About = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_ChannelsCreateChannel3d5fb10f:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Broadcast = flags&(1<<0) != 0
		m.Megagroup = flags&(1<<1) != 0
		m.Title = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.About = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 2)) != 0 {
			m6 := &InputGeoPoint{}
			err = m6.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.GeoPoint = m6
		}
		if (flags & (1 << 2)) != 0 {
			m.Address = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	default:
		log.Errorf("ChannelsCreateChannel Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsDeleteChannel) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsDeleteChannel:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1

	default:
		log.Errorf("ChannelsDeleteChannel Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsDeleteHistory) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsDeleteHistory:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsDeleteHistory Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsDeleteMessages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsDeleteMessages:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsDeleteMessages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsDeleteUserHistory) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsDeleteUserHistory:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2

	default:
		log.Errorf("ChannelsDeleteUserHistory Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsEditAbout) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsEditAbout:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.About = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsEditAbout Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsEditAdmin) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsEditAdmin:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2
		m3 := &ChannelAdminRights{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.AdminRights20B8821471 = m3

	case Cmd_ChannelsEditAdmin70f893ba:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2
		m3 := &ChatAdminRights{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.AdminRights70F893BA93 = m3

	case Cmd_ChannelsEditAdmind33c8902:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2
		m3 := &ChatAdminRights{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.AdminRights70F893BA93 = m3
		m.Rank = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_ChannelsEditAdmineb7611d0:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2
		m3 := &ChannelParticipantRole{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Role = m3

	default:
		log.Errorf("ChannelsEditAdmin Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsEditBanned) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsEditBanned:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2
		m3 := &ChannelBannedRights{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.BannedRightsBFD915CD71 = m3

	case Cmd_ChannelsEditBanned72796912:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2
		m3 := &ChatBannedRights{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.BannedRights7279691293 = m3

	default:
		log.Errorf("ChannelsEditBanned Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsEditCreator) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsEditCreator:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2
		m3 := &InputCheckPasswordSRP{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Password = m3

	default:
		log.Errorf("ChannelsEditCreator Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsEditLocation) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsEditLocation:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputGeoPoint{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.GeoPoint = m2
		m.Address = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsEditLocation Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsEditPhoto) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsEditPhoto:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputChatPhoto{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Photo = m2

	default:
		log.Errorf("ChannelsEditPhoto Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsEditTitle) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsEditTitle:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.Title = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsEditTitle Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsExportInvite) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsExportInvite:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1

	default:
		log.Errorf("ChannelsExportInvite Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsExportMessageLink) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsExportMessageLink:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_ChannelsExportMessageLinkceb77163:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &Bool{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.GroupedCEB7716385 = m3

	case Cmd_ChannelsExportMessageLinke63fadeb:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.GroupedE63FADEB119 = flags&(1<<0) != 0
		m.Thread = flags&(1<<1) != 0
		m4 := &InputChannel{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m4
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsExportMessageLink Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetAdminLog) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetAdminLog:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputChannel{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m2
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m4 := &ChannelAdminLogEventsFilter{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.EventsFilter = m4
		}
		if (flags & (1 << 1)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l5 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Admins = make([]*InputUser, l5)
			for i := int32(0); i < l5; i++ {
				m.Admins[i] = &InputUser{}
				err = (*m.Admins[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}
		m.MaxId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MinId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsGetAdminLog Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetAdminedPublicChannels) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetAdminedPublicChannels:
		_ = m

	case Cmd_ChannelsGetAdminedPublicChannelsf8b036af:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.ByLocation = flags&(1<<0) != 0
		m.CheckLimit = flags&(1<<1) != 0

	default:
		log.Errorf("ChannelsGetAdminedPublicChannels Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetBroadcastsForDiscussion) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetBroadcastsForDiscussion:
		_ = m

	default:
		log.Errorf("ChannelsGetBroadcastsForDiscussion Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetChannels) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetChannels:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Id = make([]*InputChannel, l1)
		for i := int32(0); i < l1; i++ {
			m.Id[i] = &InputChannel{}
			err = (*m.Id[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("ChannelsGetChannels Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetDialogs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetDialogs:
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsGetDialogs Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetFullChannel) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetFullChannel:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1

	default:
		log.Errorf("ChannelsGetFullChannel Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetGroupsForDiscussion) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetGroupsForDiscussion:
		_ = m

	default:
		log.Errorf("ChannelsGetGroupsForDiscussion Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetImportantHistory) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetImportantHistory:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AddOffset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MinId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsGetImportantHistory Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetInactiveChannels) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetInactiveChannels:
		_ = m

	default:
		log.Errorf("ChannelsGetInactiveChannels Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetLeftChannels) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetLeftChannels:
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsGetLeftChannels Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetMessages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetMessages:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.Id93D7B34771 = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_ChannelsGetMessagesad8c9a23:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l2 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.IdAD8C9A2385 = make([]*InputMessage, l2)
		for i := int32(0); i < l2; i++ {
			m.IdAD8C9A2385[i] = &InputMessage{}
			err = (*m.IdAD8C9A2385[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("ChannelsGetMessages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetParticipant) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetParticipant:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2

	default:
		log.Errorf("ChannelsGetParticipant Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsGetParticipants) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsGetParticipants:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &ChannelParticipantsFilter{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Filter = m2
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_ChannelsGetParticipants123e05e9:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &ChannelParticipantsFilter{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Filter = m2
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsGetParticipants Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsInviteToChannel) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsInviteToChannel:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l2 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Users = make([]*InputUser, l2)
		for i := int32(0); i < l2; i++ {
			m.Users[i] = &InputUser{}
			err = (*m.Users[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("ChannelsInviteToChannel Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsJoinChannel) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsJoinChannel:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1

	default:
		log.Errorf("ChannelsJoinChannel Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsKickFromChannel) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsKickFromChannel:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2
		m3 := &Bool{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Kicked = m3

	default:
		log.Errorf("ChannelsKickFromChannel Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsLeaveChannel) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsLeaveChannel:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1

	default:
		log.Errorf("ChannelsLeaveChannel Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsReadHistory) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsReadHistory:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsReadHistory Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsReadMessageContents) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsReadMessageContents:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsReadMessageContents Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsReportSpam) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsReportSpam:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsReportSpam Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsSetDiscussionGroup) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsSetDiscussionGroup:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Broadcast = m1
		m2 := &InputChannel{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Group = m2

	default:
		log.Errorf("ChannelsSetDiscussionGroup Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsSetStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsSetStickers:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &InputStickerSet{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Stickerset = m2

	default:
		log.Errorf("ChannelsSetStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsToggleComments) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsToggleComments:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Enabled = m2

	default:
		log.Errorf("ChannelsToggleComments Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsToggleInvites) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsToggleInvites:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Enabled = m2

	default:
		log.Errorf("ChannelsToggleInvites Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsTogglePreHistoryHidden) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsTogglePreHistoryHidden:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Enabled = m2

	default:
		log.Errorf("ChannelsTogglePreHistoryHidden Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsToggleSignatures) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsToggleSignatures:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Enabled = m2

	default:
		log.Errorf("ChannelsToggleSignatures Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsToggleSlowMode) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsToggleSlowMode:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.Seconds = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsToggleSlowMode Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsUpdatePinnedMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsUpdatePinnedMessage:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<0) != 0
		m3 := &InputChannel{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m3
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsUpdatePinnedMessage Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLChannelsUpdateUsername) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ChannelsUpdateUsername:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.Username = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ChannelsUpdateUsername Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsAcceptContact) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsAcceptContact:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1

	default:
		log.Errorf("ContactsAcceptContact Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsAddContact) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsAddContact:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.AddPhonePrivacyException = flags&(1<<0) != 0
		m3 := &InputUser{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m3
		m.FirstName = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.LastName = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Phone = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ContactsAddContact Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsBlock) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsBlock:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id332B49FC71 = m1

	case Cmd_ContactsBlock68cc1411:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id68CC1411119 = m1

	default:
		log.Errorf("ContactsBlock Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsBlockFromReplies) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsBlockFromReplies:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.DeleteMessage = flags&(1<<0) != 0
		m.DeleteHistory = flags&(1<<1) != 0
		m.ReportSpam = flags&(1<<2) != 0
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ContactsBlockFromReplies Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsDeleteByPhones) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsDeleteByPhones:
		m.Phones = dBuf.VectorString()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ContactsDeleteByPhones Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsDeleteContact) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsDeleteContact:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1

	default:
		log.Errorf("ContactsDeleteContact Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsDeleteContacts) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsDeleteContacts:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Id = make([]*InputUser, l1)
		for i := int32(0); i < l1; i++ {
			m.Id[i] = &InputUser{}
			err = (*m.Id[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	case Cmd_ContactsDeleteContacts96a0e00:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Id = make([]*InputUser, l1)
		for i := int32(0); i < l1; i++ {
			m.Id[i] = &InputUser{}
			err = (*m.Id[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("ContactsDeleteContacts Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsExportCard) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsExportCard:
		_ = m

	default:
		log.Errorf("ContactsExportCard Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsGetBlocked) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsGetBlocked:
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ContactsGetBlocked Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsGetContactIDs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsGetContactIDs:
		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ContactsGetContactIDs Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsGetContacts) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsGetContacts:
		m.HashC023849F71 = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_ContactsGetContacts22c6aa08:
		m.Hash22C6AA0851 = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ContactsGetContacts Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsGetLocated) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsGetLocated:
		m1 := &InputGeoPoint{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.GeoPoint = m1

	case Cmd_ContactsGetLocatedd348bc44:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Background = flags&(1<<1) != 0
		m3 := &InputGeoPoint{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.GeoPoint = m3
		if (flags & (1 << 0)) != 0 {
			m.SelfExpires = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	default:
		log.Errorf("ContactsGetLocated Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsGetSaved) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsGetSaved:
		_ = m

	default:
		log.Errorf("ContactsGetSaved Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsGetStatuses) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsGetStatuses:
		_ = m

	default:
		log.Errorf("ContactsGetStatuses Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsGetTopPeers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsGetTopPeers:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Correspondents = flags&(1<<0) != 0
		m.BotsPm = flags&(1<<1) != 0
		m.BotsInline = flags&(1<<2) != 0
		m.PhoneCalls = flags&(1<<3) != 0
		m.Groups = flags&(1<<10) != 0
		m.Channels = flags&(1<<15) != 0
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ContactsGetTopPeers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsImportCard) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsImportCard:
		m.ExportCard = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ContactsImportCard Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsImportContacts) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsImportContacts:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Contacts = make([]*InputContact, l1)
		for i := int32(0); i < l1; i++ {
			m.Contacts[i] = &InputContact{}
			err = (*m.Contacts[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	case Cmd_ContactsImportContactsda30b32d:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Contacts = make([]*InputContact, l1)
		for i := int32(0); i < l1; i++ {
			m.Contacts[i] = &InputContact{}
			err = (*m.Contacts[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Replace = m2

	default:
		log.Errorf("ContactsImportContacts Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsResetSaved) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsResetSaved:
		_ = m

	default:
		log.Errorf("ContactsResetSaved Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsResetTopPeerRating) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsResetTopPeerRating:
		m1 := &TopPeerCategory{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Category = m1
		m2 := &InputPeer{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m2

	default:
		log.Errorf("ContactsResetTopPeerRating Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsResolveUsername) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsResolveUsername:
		m.Username = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ContactsResolveUsername Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsSearch) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsSearch:
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ContactsSearch Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsToggleTopPeers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsToggleTopPeers:
		m1 := &Bool{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Enabled = m1

	default:
		log.Errorf("ContactsToggleTopPeers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContactsUnblock) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContactsUnblock:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.IdE54100BD71 = m1

	case Cmd_ContactsUnblockbea65d50:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.IdBEA65D50119 = m1

	default:
		log.Errorf("ContactsUnblock Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLContestSaveDeveloperInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ContestSaveDeveloperInfo:
		m.VkId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Name = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PhoneNumber = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Age = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.City = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ContestSaveDeveloperInfo Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLDestroyAuthKey) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_DestroyAuthKey:
		_ = m

	default:
		log.Errorf("DestroyAuthKey Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLDestroySession) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_DestroySession:
		m.SessionId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("DestroySession Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLFoldersDeleteFolder) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_FoldersDeleteFolder:
		m.FolderId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("FoldersDeleteFolder Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLFoldersEditPeerFolders) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_FoldersEditPeerFolders:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.FolderPeers = make([]*InputFolderPeer, l1)
		for i := int32(0); i < l1; i++ {
			m.FolderPeers[i] = &InputFolderPeer{}
			err = (*m.FolderPeers[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("FoldersEditPeerFolders Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLGetFutureSalts) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_GetFutureSalts:
		m.Num = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("GetFutureSalts Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpAcceptTermsOfService) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpAcceptTermsOfService:
		m1 := &DataJSON{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1

	default:
		log.Errorf("HelpAcceptTermsOfService Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpDismissSuggestion) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpDismissSuggestion:
		m.Suggestion = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("HelpDismissSuggestion Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpEditUserInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpEditUserInfo:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m1
		m.Message = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l3 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Entities = make([]*MessageEntity, l3)
		for i := int32(0); i < l3; i++ {
			m.Entities[i] = &MessageEntity{}
			err = (*m.Entities[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("HelpEditUserInfo Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetAppChangelog) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetAppChangelog:
		m.PrevAppVersion = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_HelpGetAppChangelog5bab7fb2:
		m.DeviceModel = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.SystemVersion = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AppVersion = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("HelpGetAppChangelog Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetAppConfig) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetAppConfig:
		_ = m

	default:
		log.Errorf("HelpGetAppConfig Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetAppUpdate) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetAppUpdate:
		_ = m

	case Cmd_HelpGetAppUpdate522d5a7d:
		m.Source = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_HelpGetAppUpdatec812ac7e:
		m.DeviceModel = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.SystemVersion = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AppVersion = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("HelpGetAppUpdate Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetCdnConfig) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetCdnConfig:
		_ = m

	default:
		log.Errorf("HelpGetCdnConfig Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetConfig) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetConfig:
		_ = m

	default:
		log.Errorf("HelpGetConfig Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetCountriesList) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetCountriesList:
		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("HelpGetCountriesList Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetDeepLinkInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetDeepLinkInfo:
		m.Path = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("HelpGetDeepLinkInfo Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetInviteText) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetInviteText:
		_ = m

	case Cmd_HelpGetInviteTexta4a95186:
		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("HelpGetInviteText Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetNearestDc) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetNearestDc:
		_ = m

	default:
		log.Errorf("HelpGetNearestDc Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetPassportConfig) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetPassportConfig:
		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("HelpGetPassportConfig Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetPromoData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetPromoData:
		_ = m

	default:
		log.Errorf("HelpGetPromoData Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetProxyData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetProxyData:
		_ = m

	default:
		log.Errorf("HelpGetProxyData Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetRecentMeUrls) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetRecentMeUrls:
		m.Referer = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("HelpGetRecentMeUrls Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetSupport) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetSupport:
		_ = m

	default:
		log.Errorf("HelpGetSupport Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetSupportName) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetSupportName:
		_ = m

	default:
		log.Errorf("HelpGetSupportName Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetTermsOfService) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetTermsOfService:
		_ = m

	case Cmd_HelpGetTermsOfService37d78f83:
		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("HelpGetTermsOfService Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetTermsOfServiceUpdate) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetTermsOfServiceUpdate:
		_ = m

	default:
		log.Errorf("HelpGetTermsOfServiceUpdate Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpGetUserInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpGetUserInfo:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m1

	default:
		log.Errorf("HelpGetUserInfo Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpHidePromoData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpHidePromoData:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("HelpHidePromoData Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpSaveAppLog) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpSaveAppLog:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Events = make([]*InputAppEvent, l1)
		for i := int32(0); i < l1; i++ {
			m.Events[i] = &InputAppEvent{}
			err = (*m.Events[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("HelpSaveAppLog Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpSetBotUpdatesStatus) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpSetBotUpdatesStatus:
		m.PendingUpdatesCount = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Message = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("HelpSetBotUpdatesStatus Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLHelpTest) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_HelpTest:
		_ = m

	default:
		log.Errorf("HelpTest Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLLangpackGetDifference) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_LangpackGetDifference:
		m.FromVersion = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_LangpackGetDifference9d51e814:
		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.FromVersion = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_LangpackGetDifferencecd984aa5:
		m.LangPack = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.FromVersion = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("LangpackGetDifference Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLLangpackGetLangPack) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_LangpackGetLangPack:
		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_LangpackGetLangPackf2f2330a:
		m.LangPack = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("LangpackGetLangPack Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLLangpackGetLanguage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_LangpackGetLanguage:
		m.LangPack = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("LangpackGetLanguage Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLLangpackGetLanguages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_LangpackGetLanguages:
		_ = m

	case Cmd_LangpackGetLanguages42c6978f:
		m.LangPack = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("LangpackGetLanguages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLLangpackGetStrings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_LangpackGetStrings:
		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Keys = dBuf.VectorString()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_LangpackGetStringsefea3803:
		m.LangPack = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Keys = dBuf.VectorString()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("LangpackGetStrings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesAcceptEncryption) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesAcceptEncryption:
		m1 := &InputEncryptedChat{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.GB = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.KeyFingerprint = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesAcceptEncryption Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesAcceptUrlAuth) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesAcceptUrlAuth:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.WriteAllowed = flags&(1<<0) != 0
		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ButtonId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesAcceptUrlAuth Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesAddChatUser) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesAddChatUser:
		m.ChatId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2
		m.FwdLimit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesAddChatUser Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesCheckChatInvite) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesCheckChatInvite:
		m.Hash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesCheckChatInvite Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesClearAllDrafts) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesClearAllDrafts:
		_ = m

	default:
		log.Errorf("MessagesClearAllDrafts Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesClearRecentStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesClearRecentStickers:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Attached = flags&(1<<0) != 0

	default:
		log.Errorf("MessagesClearRecentStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesCreateChat) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesCreateChat:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Users = make([]*InputUser, l1)
		for i := int32(0); i < l1; i++ {
			m.Users[i] = &InputUser{}
			err = (*m.Users[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}
		m.Title = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesCreateChat Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesDeleteChatUser) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesDeleteChatUser:
		m.ChatId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2

	default:
		log.Errorf("MessagesDeleteChatUser Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesDeleteHistory) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesDeleteHistory:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesDeleteHistory1c015b09:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.JustClear = flags&(1<<0) != 0
		m.Revoke = flags&(1<<1) != 0
		m4 := &InputPeer{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m4
		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesDeleteHistory Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesDeleteMessages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesDeleteMessages:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Revoke = flags&(1<<0) != 0
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesDeleteMessagesa5f18925:
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesDeleteMessages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesDeleteScheduledMessages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesDeleteScheduledMessages:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesDeleteScheduledMessages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesDiscardEncryption) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesDiscardEncryption:
		m.ChatId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesDiscardEncryption Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesEditChatAbout) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesEditChatAbout:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.About = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesEditChatAbout Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesEditChatAdmin) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesEditChatAdmin:
		m.ChatId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2
		m3 := &Bool{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.IsAdmin = m3

	default:
		log.Errorf("MessagesEditChatAdmin Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesEditChatDefaultBannedRights) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesEditChatDefaultBannedRights:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m2 := &ChatBannedRights{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.BannedRights = m2

	default:
		log.Errorf("MessagesEditChatDefaultBannedRights Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesEditChatPhoto) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesEditChatPhoto:
		m.ChatId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &InputChatPhoto{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Photo = m2

	default:
		log.Errorf("MessagesEditChatPhoto Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesEditChatTitle) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesEditChatTitle:
		m.ChatId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Title = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesEditChatTitle Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesEditInlineBotMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesEditInlineBotMessage:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.NoWebpage = flags&(1<<1) != 0
		m3 := &InputBotInlineMessageID{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m3
		if (flags & (1 << 11)) != 0 {
			m.Message = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 2)) != 0 {
			m5 := &ReplyMarkup{}
			err = m5.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m5
		}
		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l6 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l6)
			for i := int32(0); i < l6; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}

	case Cmd_MessagesEditInlineBotMessage83557dba:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.NoWebpage = flags&(1<<1) != 0
		m3 := &InputBotInlineMessageID{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m3
		if (flags & (1 << 11)) != 0 {
			m.Message = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 14)) != 0 {
			m5 := &InputMedia{}
			err = m5.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Media = m5
		}
		if (flags & (1 << 2)) != 0 {
			m6 := &ReplyMarkup{}
			err = m6.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m6
		}
		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l7 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l7)
			for i := int32(0); i < l7; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}

	case Cmd_MessagesEditInlineBotMessageadc3e828:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.NoWebpage = flags&(1<<1) != 0
		m.StopGeoLive = flags&(1<<12) != 0
		m4 := &InputBotInlineMessageID{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m4
		if (flags & (1 << 11)) != 0 {
			m.Message = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 14)) != 0 {
			m6 := &InputMedia{}
			err = m6.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Media = m6
		}
		if (flags & (1 << 2)) != 0 {
			m7 := &ReplyMarkup{}
			err = m7.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m7
		}
		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l8 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l8)
			for i := int32(0); i < l8; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}
		if (flags & (1 << 13)) != 0 {
			m9 := &InputGeoPoint{}
			err = m9.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.GeoPoint = m9
		}

	default:
		log.Errorf("MessagesEditInlineBotMessage Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesEditMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesEditMessage:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.NoWebpage = flags&(1<<1) != 0
		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 11)) != 0 {
			m.Message = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 2)) != 0 {
			m6 := &ReplyMarkup{}
			err = m6.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m6
		}
		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l7 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l7)
			for i := int32(0); i < l7; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}

	case Cmd_MessagesEditMessage48f71778:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.NoWebpage = flags&(1<<1) != 0
		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 11)) != 0 {
			m.Message = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 14)) != 0 {
			m6 := &InputMedia{}
			err = m6.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Media = m6
		}
		if (flags & (1 << 2)) != 0 {
			m7 := &ReplyMarkup{}
			err = m7.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m7
		}
		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l8 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l8)
			for i := int32(0); i < l8; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}
		if (flags & (1 << 15)) != 0 {
			m.ScheduleDate = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	case Cmd_MessagesEditMessagec000e4c8:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.NoWebpage = flags&(1<<1) != 0
		m.StopGeoLive = flags&(1<<12) != 0
		m4 := &InputPeer{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m4
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 11)) != 0 {
			m.Message = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 14)) != 0 {
			m7 := &InputMedia{}
			err = m7.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Media = m7
		}
		if (flags & (1 << 2)) != 0 {
			m8 := &ReplyMarkup{}
			err = m8.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m8
		}
		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l9 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l9)
			for i := int32(0); i < l9; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}
		if (flags & (1 << 13)) != 0 {
			m10 := &InputGeoPoint{}
			err = m10.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.GeoPoint = m10
		}

	case Cmd_MessagesEditMessaged116f31e:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.NoWebpage = flags&(1<<1) != 0
		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 11)) != 0 {
			m.Message = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 14)) != 0 {
			m6 := &InputMedia{}
			err = m6.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Media = m6
		}
		if (flags & (1 << 2)) != 0 {
			m7 := &ReplyMarkup{}
			err = m7.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m7
		}
		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l8 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l8)
			for i := int32(0); i < l8; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}

	default:
		log.Errorf("MessagesEditMessage Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesExportChatInvite) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesExportChatInvite:
		m.ChatId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesExportChatInvitedf7534c:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("MessagesExportChatInvite Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesFaveSticker) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesFaveSticker:
		m1 := &InputDocument{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Unfave = m2

	default:
		log.Errorf("MessagesFaveSticker Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesForwardMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesForwardMessage:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesForwardMessage Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesForwardMessages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesForwardMessages:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<5) != 0
		m.Background = flags&(1<<6) != 0
		m.WithMyScore = flags&(1<<8) != 0
		m5 := &InputPeer{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.FromPeer = m5
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.RandomId = dBuf.VectorLong()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m8 := &InputPeer{}
		err = m8.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.ToPeer = m8

	case Cmd_MessagesForwardMessagesd9fee60e:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<5) != 0
		m.Background = flags&(1<<6) != 0
		m.WithMyScore = flags&(1<<8) != 0
		m5 := &InputPeer{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.FromPeer = m5
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.RandomId = dBuf.VectorLong()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m8 := &InputPeer{}
		err = m8.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.ToPeer = m8
		if (flags & (1 << 10)) != 0 {
			m.ScheduleDate = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	default:
		log.Errorf("MessagesForwardMessages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetAllChats) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetAllChats:
		m.ExceptIds = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetAllChats Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetAllDrafts) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetAllDrafts:
		_ = m

	default:
		log.Errorf("MessagesGetAllDrafts Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetAllStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetAllStickers:
		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetAllStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetArchivedStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetArchivedStickers:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Masks = flags&(1<<0) != 0
		m.OffsetId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetArchivedStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetAttachedStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetAttachedStickers:
		m1 := &InputStickeredMedia{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Media = m1

	default:
		log.Errorf("MessagesGetAttachedStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetBotCallbackAnswer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetBotCallbackAnswer:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Game = flags&(1<<1) != 0
		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m.Data = dBuf.StringBytes()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}

		}

	case Cmd_MessagesGetBotCallbackAnswer9342ca07:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Game = flags&(1<<1) != 0
		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m.Data = dBuf.StringBytes()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}

		}
		if (flags & (1 << 2)) != 0 {
			m6 := &InputCheckPasswordSRP{}
			err = m6.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Password = m6
		}

	case Cmd_MessagesGetBotCallbackAnswera6e94f04:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Data = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetBotCallbackAnswer Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetChats) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetChats:
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetChats Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetCommonChats) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetCommonChats:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m1
		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetCommonChats Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetDhConfig) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetDhConfig:
		m.Version = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.RandomLength = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetDhConfig Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetDialogFilters) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetDialogFilters:
		_ = m

	default:
		log.Errorf("MessagesGetDialogFilters Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetDialogUnreadMarks) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetDialogUnreadMarks:
		_ = m

	default:
		log.Errorf("MessagesGetDialogUnreadMarks Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetDialogs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetDialogs:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.ExcludePinned = flags&(1<<0) != 0
		m.OffsetDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m5 := &InputPeer{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.OffsetPeer = m5
		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesGetDialogs6b47f94d:
		m.OffsetDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.OffsetPeer = m3
		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesGetDialogsa0ee3b73:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.ExcludePinned = flags&(1<<0) != 0
		if (flags & (1 << 1)) != 0 {
			m.FolderId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m.OffsetDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m6 := &InputPeer{}
		err = m6.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.OffsetPeer = m6
		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesGetDialogsb098aee6:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.ExcludePinned = flags&(1<<0) != 0
		m.OffsetDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m5 := &InputPeer{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.OffsetPeer = m5
		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetDialogs Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetDiscussionMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetDiscussionMessage:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetDiscussionMessage Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetDocumentByHash) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetDocumentByHash:
		m.Sha256 = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Size_ = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MimeType = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetDocumentByHash Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetEmojiKeywords) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetEmojiKeywords:
		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetEmojiKeywords Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetEmojiKeywordsDifference) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetEmojiKeywordsDifference:
		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.FromVersion = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetEmojiKeywordsDifference Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetEmojiKeywordsLanguages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetEmojiKeywordsLanguages:
		m.LangCodes = dBuf.VectorString()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetEmojiKeywordsLanguages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetEmojiURL) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetEmojiURL:
		m.LangCode = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetEmojiURL Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetFavedStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetFavedStickers:
		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetFavedStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetFeaturedStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetFeaturedStickers:
		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetFeaturedStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetFullChat) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetFullChat:
		m.ChatId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetFullChat Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetGameHighScores) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetGameHighScores:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &InputUser{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m3

	default:
		log.Errorf("MessagesGetGameHighScores Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetHistory) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetHistory:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AddOffset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MinId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesGetHistorydcbb8260:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AddOffset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MinId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetHistory Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetInlineBotResults) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetInlineBotResults:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Bot = m2
		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		if (flags & (1 << 0)) != 0 {
			m4 := &InputGeoPoint{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.GeoPoint = m4
		}
		m.Query = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Offset = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetInlineBotResults Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetInlineGameHighScores) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetInlineGameHighScores:
		m1 := &InputBotInlineMessageID{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1
		m2 := &InputUser{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m2

	default:
		log.Errorf("MessagesGetInlineGameHighScores Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetMaskStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetMaskStickers:
		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetMaskStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetMessageEditData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetMessageEditData:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetMessageEditData Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetMessages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetMessages:
		m.Id4222FA7471 = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesGetMessages63c66506:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Id63C6650685 = make([]*InputMessage, l1)
		for i := int32(0); i < l1; i++ {
			m.Id63C6650685[i] = &InputMessage{}
			err = (*m.Id63C6650685[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("MessagesGetMessages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetMessagesViews) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetMessagesViews:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &Bool{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Increment = m3

	case Cmd_MessagesGetMessagesViews5784d3e1:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &Bool{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Increment = m3

	default:
		log.Errorf("MessagesGetMessagesViews Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetOldFeaturedStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetOldFeaturedStickers:
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetOldFeaturedStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetOnlines) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetOnlines:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("MessagesGetOnlines Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetPeerDialogs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetPeerDialogs:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Peers2D9776B971 = make([]*InputPeer, l1)
		for i := int32(0); i < l1; i++ {
			m.Peers2D9776B971[i] = &InputPeer{}
			err = (*m.Peers2D9776B971[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	case Cmd_MessagesGetPeerDialogse470bcfd:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.PeersE470BCFD85 = make([]*InputDialogPeer, l1)
		for i := int32(0); i < l1; i++ {
			m.PeersE470BCFD85[i] = &InputDialogPeer{}
			err = (*m.PeersE470BCFD85[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("MessagesGetPeerDialogs Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetPeerSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetPeerSettings:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("MessagesGetPeerSettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetPinnedDialogs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetPinnedDialogs:
		_ = m

	case Cmd_MessagesGetPinnedDialogsd6b94df2:
		m.FolderId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetPinnedDialogs Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetPollResults) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetPollResults:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetPollResults Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetPollVotes) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetPollVotes:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputPeer{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m2
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m.Option = dBuf.StringBytes()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}

		}
		if (flags & (1 << 1)) != 0 {
			m.Offset = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetPollVotes Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetRecentLocations) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetRecentLocations:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesGetRecentLocationsbbc45b09:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetRecentLocations Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetRecentStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetRecentStickers:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Attached = flags&(1<<0) != 0
		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetRecentStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetReplies) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetReplies:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AddOffset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MinId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetReplies Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetSavedGifs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetSavedGifs:
		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetSavedGifs Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetScheduledHistory) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetScheduledHistory:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetScheduledHistory Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetScheduledMessages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetScheduledMessages:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetScheduledMessages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetSearchCounters) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetSearchCounters:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l2 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Filters = make([]*MessagesFilter, l2)
		for i := int32(0); i < l2; i++ {
			m.Filters[i] = &MessagesFilter{}
			err = (*m.Filters[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("MessagesGetSearchCounters Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetSplitRanges) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetSplitRanges:
		_ = m

	default:
		log.Errorf("MessagesGetSplitRanges Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetStatsURL) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetStatsURL:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	case Cmd_MessagesGetStatsURL812c2ae6:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Dark = flags&(1<<0) != 0
		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.Params = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetStatsURL Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetStickerSet) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetStickerSet:
		m1 := &InputStickerSet{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Stickerset = m1

	default:
		log.Errorf("MessagesGetStickerSet Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetStickers:
		m.Emoticon = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.HashAE22E04551 = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesGetStickers43d4f2c:
		m.Emoticon = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash43D4F2C85 = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetSuggestedDialogFilters) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetSuggestedDialogFilters:
		_ = m

	default:
		log.Errorf("MessagesGetSuggestedDialogFilters Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetUnreadMentions) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetUnreadMentions:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AddOffset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MinId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetUnreadMentions Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetWebPage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetWebPage:
		m.Url = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesGetWebPage Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesGetWebPagePreview) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesGetWebPagePreview:
		m.Message = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesGetWebPagePreview8b68b0cc:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Message = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l3 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l3)
			for i := int32(0); i < l3; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}

	default:
		log.Errorf("MessagesGetWebPagePreview Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesHidePeerSettingsBar) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesHidePeerSettingsBar:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("MessagesHidePeerSettingsBar Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesHideReportSpam) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesHideReportSpam:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("MessagesHideReportSpam Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesImportChatInvite) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesImportChatInvite:
		m.Hash = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesImportChatInvite Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesInstallStickerSet) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesInstallStickerSet:
		m1 := &InputStickerSet{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Stickerset = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Archived = m2

	case Cmd_MessagesInstallStickerSet7b30c3a6:
		m1 := &InputStickerSet{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Stickerset = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Disabled = m2

	default:
		log.Errorf("MessagesInstallStickerSet Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesMarkDialogUnread) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesMarkDialogUnread:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Unread = flags&(1<<0) != 0
		m3 := &InputDialogPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3

	default:
		log.Errorf("MessagesMarkDialogUnread Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesMigrateChat) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesMigrateChat:
		m.ChatId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesMigrateChat Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReadDiscussion) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReadDiscussion:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ReadMaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesReadDiscussion Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReadEncryptedHistory) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReadEncryptedHistory:
		m1 := &InputEncryptedChat{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.MaxDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesReadEncryptedHistory Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReadFeaturedStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReadFeaturedStickers:
		m.Id = dBuf.VectorLong()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesReadFeaturedStickers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReadHistory) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReadHistory:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesReadHistoryb04f2510:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesReadHistory Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReadMentions) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReadMentions:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("MessagesReadMentions Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReadMessageContents) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReadMessageContents:
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesReadMessageContents Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReceivedMessages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReceivedMessages:
		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesReceivedMessages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReceivedQueue) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReceivedQueue:
		m.MaxQts = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesReceivedQueue Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReorderPinnedDialogs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReorderPinnedDialogs:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Force = flags&(1<<0) != 0
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l3 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Order959FF64471 = make([]*InputPeer, l3)
		for i := int32(0); i < l3; i++ {
			m.Order959FF64471[i] = &InputPeer{}
			err = (*m.Order959FF64471[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	case Cmd_MessagesReorderPinnedDialogs3b1adf37:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Force = flags&(1<<0) != 0
		m.FolderId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l4 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Order5B51D63F85 = make([]*InputDialogPeer, l4)
		for i := int32(0); i < l4; i++ {
			m.Order5B51D63F85[i] = &InputDialogPeer{}
			err = (*m.Order5B51D63F85[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	case Cmd_MessagesReorderPinnedDialogs5b51d63f:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Force = flags&(1<<0) != 0
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l3 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Order5B51D63F85 = make([]*InputDialogPeer, l3)
		for i := int32(0); i < l3; i++ {
			m.Order5B51D63F85[i] = &InputDialogPeer{}
			err = (*m.Order5B51D63F85[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("MessagesReorderPinnedDialogs Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReorderStickerSets) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReorderStickerSets:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Masks = flags&(1<<0) != 0
		m.Order = dBuf.VectorLong()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesReorderStickerSets9fcfbc30:
		m.Order = dBuf.VectorLong()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesReorderStickerSets Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReport) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReport:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &ReportReason{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Reason = m3

	default:
		log.Errorf("MessagesReport Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReportEncryptedSpam) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReportEncryptedSpam:
		m1 := &InputEncryptedChat{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("MessagesReportEncryptedSpam Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesReportSpam) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesReportSpam:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("MessagesReportSpam Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesRequestEncryption) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesRequestEncryption:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m1
		m.RandomId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.GA = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesRequestEncryption Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesRequestUrlAuth) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesRequestUrlAuth:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ButtonId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesRequestUrlAuth Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSaveDraft) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSaveDraft:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.NoWebpage = flags&(1<<1) != 0
		if (flags & (1 << 0)) != 0 {
			m.ReplyToMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m4 := &InputPeer{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m4
		m.Message = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l6 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l6)
			for i := int32(0); i < l6; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}

	default:
		log.Errorf("MessagesSaveDraft Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSaveGif) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSaveGif:
		m1 := &InputDocument{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Unsave = m2

	default:
		log.Errorf("MessagesSaveGif Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSaveRecentSticker) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSaveRecentSticker:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Attached = flags&(1<<0) != 0
		m3 := &InputDocument{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m3
		m4 := &Bool{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Unsave = m4

	case Cmd_MessagesSaveRecentSticker348e39bf:
		m1 := &InputDocument{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Unsave = m2

	default:
		log.Errorf("MessagesSaveRecentSticker Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSearch) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSearch:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputPeer{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m2
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m4 := &InputUser{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.FromId39E9EA071 = m4
		}
		m5 := &MessagesFilter{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Filter = m5
		m.MinDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AddOffset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MinId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesSearch4e17810b:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputPeer{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m2
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m4 := &InputUser{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.FromId39E9EA071 = m4
		}
		if (flags & (1 << 1)) != 0 {
			m.TopMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m6 := &MessagesFilter{}
		err = m6.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Filter = m6
		m.MinDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AddOffset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MinId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesSearch8614ef68:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputPeer{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m2
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m4 := &InputUser{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.FromId39E9EA071 = m4
		}
		m5 := &MessagesFilter{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Filter = m5
		m.MinDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AddOffset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MinId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesSearchc352eec:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputPeer{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m2
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m4 := &InputPeer{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.FromIdC352EEC120 = m4
		}
		if (flags & (1 << 1)) != 0 {
			m.TopMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m6 := &MessagesFilter{}
		err = m6.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Filter = m6
		m.MinDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.AddOffset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MinId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesSearchd4569248:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.ImportantOnly = flags&(1<<0) != 0
		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m5 := &MessagesFilter{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Filter = m5
		m.MinDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesSearchf288a275:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputPeer{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m2
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m4 := &InputUser{}
			err = m4.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.FromId39E9EA071 = m4
		}
		m5 := &MessagesFilter{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Filter = m5
		m.MinDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesSearch Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSearchGifs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSearchGifs:
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesSearchGifs Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSearchGlobal) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSearchGlobal:
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.OffsetPeer = m3
		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesSearchGlobal4bc6589a:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		if (flags & (1 << 0)) != 0 {
			m.FolderId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m4 := &MessagesFilter{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Filter = m4
		m.MinDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxDate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetRate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m8 := &InputPeer{}
		err = m8.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.OffsetPeer = m8
		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesSearchGlobalbf7225a4:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		if (flags & (1 << 0)) != 0 {
			m.FolderId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetRate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m5 := &InputPeer{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.OffsetPeer = m5
		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesSearchGlobalf79c611:
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetRate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.OffsetPeer = m3
		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesSearchGlobal Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSearchStickerSets) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSearchStickerSets:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.ExcludeFeatured = flags&(1<<0) != 0
		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Hash = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesSearchStickerSets Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSendBroadcast) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSendBroadcast:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Contacts = make([]*InputUser, l1)
		for i := int32(0); i < l1; i++ {
			m.Contacts[i] = &InputUser{}
			err = (*m.Contacts[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}
		m.RandomId = dBuf.VectorLong()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Message = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m4 := &InputMedia{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Media = m4

	default:
		log.Errorf("MessagesSendBroadcast Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSendEncrypted) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSendEncrypted:
		m1 := &InputEncryptedChat{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Data = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesSendEncrypted44fa7a15:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<0) != 0
		m3 := &InputEncryptedChat{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Data = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesSendEncrypted Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSendEncryptedFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSendEncryptedFile:
		m1 := &InputEncryptedChat{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Data = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m4 := &InputEncryptedFile{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.File = m4

	case Cmd_MessagesSendEncryptedFile5559481d:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<0) != 0
		m3 := &InputEncryptedChat{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Data = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m6 := &InputEncryptedFile{}
		err = m6.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.File = m6

	default:
		log.Errorf("MessagesSendEncryptedFile Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSendEncryptedService) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSendEncryptedService:
		m1 := &InputEncryptedChat{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Data = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesSendEncryptedService Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSendInlineBotResult) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSendInlineBotResult:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<5) != 0
		m.Background = flags&(1<<6) != 0
		m.ClearDraft = flags&(1<<7) != 0
		m5 := &InputPeer{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m5
		if (flags & (1 << 0)) != 0 {
			m.ReplyToMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.QueryId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Id = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesSendInlineBotResult220815b0:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<5) != 0
		m.Background = flags&(1<<6) != 0
		m.ClearDraft = flags&(1<<7) != 0
		m.HideVia = flags&(1<<11) != 0
		m6 := &InputPeer{}
		err = m6.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m6
		if (flags & (1 << 0)) != 0 {
			m.ReplyToMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.QueryId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Id = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 10)) != 0 {
			m.ScheduleDate = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	default:
		log.Errorf("MessagesSendInlineBotResult Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSendMedia) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSendMedia:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<5) != 0
		m.Background = flags&(1<<6) != 0
		m.ClearDraft = flags&(1<<7) != 0
		m5 := &InputPeer{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m5
		if (flags & (1 << 0)) != 0 {
			m.ReplyToMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m7 := &InputMedia{}
		err = m7.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Media = m7
		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 2)) != 0 {
			m9 := &ReplyMarkup{}
			err = m9.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m9
		}

	case Cmd_MessagesSendMedia3491eba9:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<5) != 0
		m.Background = flags&(1<<6) != 0
		m.ClearDraft = flags&(1<<7) != 0
		m5 := &InputPeer{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m5
		if (flags & (1 << 0)) != 0 {
			m.ReplyToMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m7 := &InputMedia{}
		err = m7.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Media = m7
		m.Message = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 2)) != 0 {
			m10 := &ReplyMarkup{}
			err = m10.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m10
		}
		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l11 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l11)
			for i := int32(0); i < l11; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}
		if (flags & (1 << 10)) != 0 {
			m.ScheduleDate = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	case Cmd_MessagesSendMediab8d1262b:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<5) != 0
		m.Background = flags&(1<<6) != 0
		m.ClearDraft = flags&(1<<7) != 0
		m5 := &InputPeer{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m5
		if (flags & (1 << 0)) != 0 {
			m.ReplyToMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m7 := &InputMedia{}
		err = m7.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Media = m7
		m.Message = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 2)) != 0 {
			m10 := &ReplyMarkup{}
			err = m10.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m10
		}
		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l11 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l11)
			for i := int32(0); i < l11; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}

	default:
		log.Errorf("MessagesSendMedia Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSendMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSendMessage:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.NoWebpage = flags&(1<<1) != 0
		m.Silent = flags&(1<<5) != 0
		m.Background = flags&(1<<6) != 0
		m.ClearDraft = flags&(1<<7) != 0
		m6 := &InputPeer{}
		err = m6.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m6
		if (flags & (1 << 0)) != 0 {
			m.ReplyToMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m.Message = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 2)) != 0 {
			m10 := &ReplyMarkup{}
			err = m10.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m10
		}
		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l11 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l11)
			for i := int32(0); i < l11; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}

	case Cmd_MessagesSendMessage520c3870:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.NoWebpage = flags&(1<<1) != 0
		m.Silent = flags&(1<<5) != 0
		m.Background = flags&(1<<6) != 0
		m.ClearDraft = flags&(1<<7) != 0
		m6 := &InputPeer{}
		err = m6.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m6
		if (flags & (1 << 0)) != 0 {
			m.ReplyToMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m.Message = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 2)) != 0 {
			m10 := &ReplyMarkup{}
			err = m10.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.ReplyMarkup = m10
		}
		if (flags & (1 << 3)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l11 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.Entities = make([]*MessageEntity, l11)
			for i := int32(0); i < l11; i++ {
				m.Entities[i] = &MessageEntity{}
				err = (*m.Entities[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}
		if (flags & (1 << 10)) != 0 {
			m.ScheduleDate = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	default:
		log.Errorf("MessagesSendMessage Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSendMultiMedia) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSendMultiMedia:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<5) != 0
		m.Background = flags&(1<<6) != 0
		m.ClearDraft = flags&(1<<7) != 0
		m5 := &InputPeer{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m5
		if (flags & (1 << 0)) != 0 {
			m.ReplyToMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l7 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.MultiMedia = make([]*InputSingleMedia, l7)
		for i := int32(0); i < l7; i++ {
			m.MultiMedia[i] = &InputSingleMedia{}
			err = (*m.MultiMedia[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	case Cmd_MessagesSendMultiMediacc0110cb:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<5) != 0
		m.Background = flags&(1<<6) != 0
		m.ClearDraft = flags&(1<<7) != 0
		m5 := &InputPeer{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m5
		if (flags & (1 << 0)) != 0 {
			m.ReplyToMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l7 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.MultiMedia = make([]*InputSingleMedia, l7)
		for i := int32(0); i < l7; i++ {
			m.MultiMedia[i] = &InputSingleMedia{}
			err = (*m.MultiMedia[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}
		if (flags & (1 << 10)) != 0 {
			m.ScheduleDate = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	default:
		log.Errorf("MessagesSendMultiMedia Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSendScheduledMessages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSendScheduledMessages:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Id = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesSendScheduledMessages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSendScreenshotNotification) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSendScreenshotNotification:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.ReplyToMsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesSendScreenshotNotification Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSendVote) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSendVote:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesSendVote Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSetBotCallbackAnswer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSetBotCallbackAnswer:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Alert = flags&(1<<1) != 0
		m.QueryId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m.Message = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 2)) != 0 {
			m.Url = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m.CacheTime = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_MessagesSetBotCallbackAnswer481c591a:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Alert = flags&(1<<1) != 0
		m.QueryId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m.Message = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	default:
		log.Errorf("MessagesSetBotCallbackAnswer Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSetBotPrecheckoutResults) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSetBotPrecheckoutResults:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Success = flags&(1<<1) != 0
		m.QueryId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m.Error = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	default:
		log.Errorf("MessagesSetBotPrecheckoutResults Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSetBotShippingResults) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSetBotShippingResults:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.QueryId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m.Error = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 1)) != 0 {
			dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			l4 := dBuf.Int()

			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
			m.ShippingOptions = make([]*ShippingOption, l4)
			for i := int32(0); i < l4; i++ {
				m.ShippingOptions[i] = &ShippingOption{}
				err = (*m.ShippingOptions[i]).Decode(dBuf, layer)
				if err != nil {
					return err
				}
			}
		}

	default:
		log.Errorf("MessagesSetBotShippingResults Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSetEncryptedTyping) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSetEncryptedTyping:
		m1 := &InputEncryptedChat{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Typing = m2

	default:
		log.Errorf("MessagesSetEncryptedTyping Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSetGameScore) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSetGameScore:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.EditMessage = flags&(1<<0) != 0
		m.Force = flags&(1<<1) != 0
		m4 := &InputPeer{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m4
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m6 := &InputUser{}
		err = m6.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m6
		m.Score = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesSetGameScore Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSetInlineBotResults) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSetInlineBotResults:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Gallery = flags&(1<<0) != 0
		m.Private = flags&(1<<1) != 0
		m.QueryId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l5 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Results = make([]*InputBotInlineResult, l5)
		for i := int32(0); i < l5; i++ {
			m.Results[i] = &InputBotInlineResult{}
			err = (*m.Results[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}
		m.CacheTime = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 2)) != 0 {
			m.NextOffset = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 3)) != 0 {
			m8 := &InlineBotSwitchPM{}
			err = m8.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.SwitchPm = m8
		}

	default:
		log.Errorf("MessagesSetInlineBotResults Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSetInlineGameScore) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSetInlineGameScore:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.EditMessage = flags&(1<<0) != 0
		m.Force = flags&(1<<1) != 0
		m4 := &InputBotInlineMessageID{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m4
		m5 := &InputUser{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m5
		m.Score = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesSetInlineGameScore Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesSetTyping) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesSetTyping:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m2 := &SendMessageAction{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Action = m2

	case Cmd_MessagesSetTyping58943ee2:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputPeer{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m2
		if (flags & (1 << 0)) != 0 {
			m.TopMsgId = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m4 := &SendMessageAction{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Action = m4

	default:
		log.Errorf("MessagesSetTyping Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesStartBot) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesStartBot:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Bot = m1
		m2 := &InputPeer{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m2
		m.RandomId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.StartParam = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesStartBot Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesToggleChatAdmins) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesToggleChatAdmins:
		m.ChatId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m2 := &Bool{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Enabled = m2

	default:
		log.Errorf("MessagesToggleChatAdmins Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesToggleDialogPin) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesToggleDialogPin:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Pinned = flags&(1<<0) != 0
		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer3289BE6A71 = m3

	case Cmd_MessagesToggleDialogPina731e257:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Pinned = flags&(1<<0) != 0
		m3 := &InputDialogPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.PeerA731E25785 = m3

	default:
		log.Errorf("MessagesToggleDialogPin Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesToggleStickerSets) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesToggleStickerSets:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Uninstall = flags&(1<<0) != 0
		m.Archive = flags&(1<<1) != 0
		m.Unarchive = flags&(1<<2) != 0
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l5 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Stickersets = make([]*InputStickerSet, l5)
		for i := int32(0); i < l5; i++ {
			m.Stickersets[i] = &InputStickerSet{}
			err = (*m.Stickersets[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("MessagesToggleStickerSets Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesUninstallStickerSet) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesUninstallStickerSet:
		m1 := &InputStickerSet{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Stickerset = m1

	default:
		log.Errorf("MessagesUninstallStickerSet Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesUnpinAllMessages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesUnpinAllMessages:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("MessagesUnpinAllMessages Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesUpdateDialogFilter) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesUpdateDialogFilter:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m3 := &DialogFilter{}
			err = m3.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Filter = m3
		}

	default:
		log.Errorf("MessagesUpdateDialogFilter Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesUpdateDialogFiltersOrder) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesUpdateDialogFiltersOrder:
		m.Order = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesUpdateDialogFiltersOrder Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesUpdatePinnedMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesUpdatePinnedMessage:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Silent = flags&(1<<0) != 0
		m3 := &InputPeer{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.Id = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("MessagesUpdatePinnedMessage Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesUploadEncryptedFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesUploadEncryptedFile:
		m1 := &InputEncryptedChat{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m2 := &InputEncryptedFile{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.File = m2

	default:
		log.Errorf("MessagesUploadEncryptedFile Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLMessagesUploadMedia) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_MessagesUploadMedia:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m2 := &InputMedia{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Media = m2

	default:
		log.Errorf("MessagesUploadMedia Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPaymentsClearSavedInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PaymentsClearSavedInfo:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Credentials = flags&(1<<0) != 0
		m.Info = flags&(1<<1) != 0

	default:
		log.Errorf("PaymentsClearSavedInfo Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPaymentsGetBankCardData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PaymentsGetBankCardData:
		m.Number = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PaymentsGetBankCardData Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPaymentsGetPaymentForm) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PaymentsGetPaymentForm:
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PaymentsGetPaymentForm Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPaymentsGetPaymentReceipt) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PaymentsGetPaymentReceipt:
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PaymentsGetPaymentReceipt Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPaymentsGetSavedInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PaymentsGetSavedInfo:
		_ = m

	default:
		log.Errorf("PaymentsGetSavedInfo Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPaymentsSendPaymentForm) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PaymentsSendPaymentForm:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m.RequestedInfoId = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		if (flags & (1 << 1)) != 0 {
			m.ShippingOptionId = dBuf.String()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m5 := &InputPaymentCredentials{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Credentials = m5

	default:
		log.Errorf("PaymentsSendPaymentForm Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPaymentsValidateRequestedInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PaymentsValidateRequestedInfo:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Save = flags&(1<<0) != 0
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m4 := &PaymentRequestedInfo{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Info = m4

	default:
		log.Errorf("PaymentsValidateRequestedInfo Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneAcceptCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneAcceptCall:
		m1 := &InputPhoneCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.GB = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &PhoneCallProtocol{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Protocol = m3

	default:
		log.Errorf("PhoneAcceptCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneCheckGroupCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneCheckGroupCall:
		m1 := &InputGroupCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Call = m1
		m.Source = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PhoneCheckGroupCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneConfirmCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneConfirmCall:
		m1 := &InputPhoneCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.GA = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.KeyFingerprint = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m4 := &PhoneCallProtocol{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Protocol = m4

	default:
		log.Errorf("PhoneConfirmCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneCreateGroupCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneCreateGroupCall:
		m1 := &InputPeer{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.RandomId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PhoneCreateGroupCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneDiscardCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneDiscardCall:
		m1 := &InputPhoneCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Duration = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &PhoneCallDiscardReason{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Reason = m3
		m.ConnectionId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_PhoneDiscardCallb2cbc1c0:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Video = flags&(1<<0) != 0
		m3 := &InputPhoneCall{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.Duration = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m5 := &PhoneCallDiscardReason{}
		err = m5.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Reason = m5
		m.ConnectionId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PhoneDiscardCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneDiscardGroupCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneDiscardGroupCall:
		m1 := &InputGroupCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Call = m1

	default:
		log.Errorf("PhoneDiscardGroupCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneEditGroupCallMember) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneEditGroupCallMember:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Muted = flags&(1<<0) != 0
		m3 := &InputGroupCall{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Call = m3
		m4 := &InputUser{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m4

	default:
		log.Errorf("PhoneEditGroupCallMember Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneGetCallConfig) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneGetCallConfig:
		_ = m

	default:
		log.Errorf("PhoneGetCallConfig Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneGetGroupCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneGetGroupCall:
		m1 := &InputGroupCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Call = m1

	default:
		log.Errorf("PhoneGetGroupCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneGetGroupParticipants) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneGetGroupParticipants:
		m1 := &InputGroupCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Call = m1
		m.Ids = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Sources = dBuf.VectorInt()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Offset = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PhoneGetGroupParticipants Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneInviteToGroupCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneInviteToGroupCall:
		m1 := &InputGroupCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Call = m1
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l2 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Users = make([]*InputUser, l2)
		for i := int32(0); i < l2; i++ {
			m.Users[i] = &InputUser{}
			err = (*m.Users[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("PhoneInviteToGroupCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneJoinGroupCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneJoinGroupCall:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Muted = flags&(1<<0) != 0
		m3 := &InputGroupCall{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Call = m3
		m4 := &DataJSON{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Params = m4

	default:
		log.Errorf("PhoneJoinGroupCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneLeaveGroupCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneLeaveGroupCall:
		m1 := &InputGroupCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Call = m1
		m.Source = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PhoneLeaveGroupCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneReceivedCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneReceivedCall:
		m1 := &InputPhoneCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1

	default:
		log.Errorf("PhoneReceivedCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneRequestCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneRequestCall:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m1
		m.RandomId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.GAHash = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m4 := &PhoneCallProtocol{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Protocol = m4

	case Cmd_PhoneRequestCall42ff96ed:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Video = flags&(1<<0) != 0
		m3 := &InputUser{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m3
		m.RandomId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.GAHash = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m6 := &PhoneCallProtocol{}
		err = m6.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Protocol = m6

	default:
		log.Errorf("PhoneRequestCall Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneSaveCallDebug) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneSaveCallDebug:
		m1 := &InputPhoneCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m2 := &DataJSON{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Debug = m2

	default:
		log.Errorf("PhoneSaveCallDebug Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneSendSignalingData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneSendSignalingData:
		m1 := &InputPhoneCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Data = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PhoneSendSignalingData Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneSetCallRating) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneSetCallRating:
		m1 := &InputPhoneCall{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m1
		m.Rating = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Comment = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_PhoneSetCallRating59ead627:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.UserInitiative = flags&(1<<0) != 0
		m3 := &InputPhoneCall{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Peer = m3
		m.Rating = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Comment = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PhoneSetCallRating Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhoneToggleGroupCallSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhoneToggleGroupCallSettings:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m2 := &InputGroupCall{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Call = m2
		if (flags & (1 << 0)) != 0 {
			m3 := &Bool{}
			err = m3.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.JoinMuted = m3
		}

	default:
		log.Errorf("PhoneToggleGroupCallSettings Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhotosDeletePhotos) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhotosDeletePhotos:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Id = make([]*InputPhoto, l1)
		for i := int32(0); i < l1; i++ {
			m.Id[i] = &InputPhoto{}
			err = (*m.Id[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("PhotosDeletePhotos Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhotosGetUserPhotos) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhotosGetUserPhotos:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m1
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.MaxId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PhotosGetUserPhotos Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhotosUpdateProfilePhoto) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhotosUpdateProfilePhoto:
		m1 := &InputPhoto{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1

	case Cmd_PhotosUpdateProfilePhoto72d4742c:
		m1 := &InputPhoto{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1

	case Cmd_PhotosUpdateProfilePhotoeef579a0:
		m1 := &InputPhoto{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1
		m2 := &InputPhotoCrop{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Crop = m2

	default:
		log.Errorf("PhotosUpdateProfilePhoto Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPhotosUploadProfilePhoto) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PhotosUploadProfilePhoto:
		m1 := &InputFile{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.File = m1

	case Cmd_PhotosUploadProfilePhoto89f30f69:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		if (flags & (1 << 0)) != 0 {
			m2 := &InputFile{}
			err = m2.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.File = m2
		}
		if (flags & (1 << 1)) != 0 {
			m3 := &InputFile{}
			err = m3.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Video = m3
		}

	case Cmd_PhotosUploadProfilePhotod50f9c88:
		m1 := &InputFile{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.File = m1
		m.Caption = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m3 := &InputGeoPoint{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.GeoPoint = m3
		m4 := &InputPhotoCrop{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Crop = m4

	default:
		log.Errorf("PhotosUploadProfilePhoto Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPing) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_Ping:
		m.PingId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("Ping Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLPingDelayDisconnect) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_PingDelayDisconnect:
		m.PingId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.DisconnectDelay = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("PingDelayDisconnect Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLReqPq) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ReqPq:
		m.Nonce = dBuf.Bytes(16)
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ReqPq Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLReqPqMulti) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_ReqPqMulti:
		m.Nonce = dBuf.Bytes(16)
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("ReqPqMulti Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLReq_DHParams) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_Req_DHParams:
		m.Nonce = dBuf.Bytes(16)
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ServerNonce = dBuf.Bytes(16)
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.P = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Q = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.PublicKeyFingerprint = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.EncryptedData = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("Req_DHParams Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLRpcDropAnswer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_RpcDropAnswer:
		m.ReqMsgId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("RpcDropAnswer Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLSetClient_DHParams) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_SetClient_DHParams:
		m.Nonce = dBuf.Bytes(16)
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ServerNonce = dBuf.Bytes(16)
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.EncryptedData = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("SetClient_DHParams Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLStatsGetBroadcastStats) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_StatsGetBroadcastStats:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Dark = flags&(1<<0) != 0
		m3 := &InputChannel{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m3

	case Cmd_StatsGetBroadcastStatse6300dba:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Dark = flags&(1<<0) != 0
		m3 := &InputChannel{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m3
		m.TzOffset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("StatsGetBroadcastStats Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLStatsGetMegagroupStats) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_StatsGetMegagroupStats:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Dark = flags&(1<<0) != 0
		m3 := &InputChannel{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m3

	default:
		log.Errorf("StatsGetMegagroupStats Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLStatsGetMessagePublicForwards) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_StatsGetMessagePublicForwards:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.OffsetRate = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m4 := &InputPeer{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.OffsetPeer = m4
		m.OffsetId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("StatsGetMessagePublicForwards Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLStatsGetMessageStats) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_StatsGetMessageStats:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Dark = flags&(1<<0) != 0
		m3 := &InputChannel{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m3
		m.MsgId = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("StatsGetMessageStats Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLStatsLoadAsyncGraph) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_StatsLoadAsyncGraph:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Token = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m.X = dBuf.Long()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}

	default:
		log.Errorf("StatsLoadAsyncGraph Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLStickersAddStickerToSet) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_StickersAddStickerToSet:
		m1 := &InputStickerSet{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Stickerset = m1
		m2 := &InputStickerSetItem{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Sticker = m2

	default:
		log.Errorf("StickersAddStickerToSet Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLStickersChangeStickerPosition) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_StickersChangeStickerPosition:
		m1 := &InputDocument{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Sticker = m1
		m.Position = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("StickersChangeStickerPosition Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLStickersCreateStickerSet) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_StickersCreateStickerSet:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Masks = flags&(1<<0) != 0
		m3 := &InputUser{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m3
		m.Title = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ShortName = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l6 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Stickers = make([]*InputStickerSetItem, l6)
		for i := int32(0); i < l6; i++ {
			m.Stickers[i] = &InputStickerSetItem{}
			err = (*m.Stickers[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	case Cmd_StickersCreateStickerSetf1036780:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Masks = flags&(1<<0) != 0
		m.Animated = flags&(1<<1) != 0
		m4 := &InputUser{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.UserId = m4
		m.Title = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.ShortName = dBuf.String()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 2)) != 0 {
			m7 := &InputDocument{}
			err = m7.Decode(dBuf, layer)
			if err != nil {
				return err
			}
			m.Thumb = m7
		}
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l8 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Stickers = make([]*InputStickerSetItem, l8)
		for i := int32(0); i < l8; i++ {
			m.Stickers[i] = &InputStickerSetItem{}
			err = (*m.Stickers[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("StickersCreateStickerSet Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLStickersRemoveStickerFromSet) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_StickersRemoveStickerFromSet:
		m1 := &InputDocument{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Sticker = m1

	default:
		log.Errorf("StickersRemoveStickerFromSet Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLStickersSetStickerSetThumb) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_StickersSetStickerSetThumb:
		m1 := &InputStickerSet{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Stickerset = m1
		m2 := &InputDocument{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Thumb = m2

	default:
		log.Errorf("StickersSetStickerSetThumb Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUpdatesGetChannelDifference) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UpdatesGetChannelDifference:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Force = flags&(1<<0) != 0
		m3 := &InputChannel{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m3
		m4 := &ChannelMessagesFilter{}
		err = m4.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Filter = m4
		m.Pts = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_UpdatesGetChannelDifferencebb32d7c0:
		m1 := &InputChannel{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Channel = m1
		m2 := &ChannelMessagesFilter{}
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Filter = m2
		m.Pts = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("UpdatesGetChannelDifference Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUpdatesGetDifference) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UpdatesGetDifference:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Pts = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		if (flags & (1 << 0)) != 0 {
			m.PtsTotalLimit = dBuf.Int()
			if dBuf.GetError() != nil {
				log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
				return dBuf.GetError()
			}
		}
		m.Date = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Qts = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_UpdatesGetDifferencea041495:
		m.Pts = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Date = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Qts = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("UpdatesGetDifference Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUpdatesGetState) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UpdatesGetState:
		_ = m

	default:
		log.Errorf("UpdatesGetState Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUploadGetCdnFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UploadGetCdnFile:
		m.FileToken = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("UploadGetCdnFile Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUploadGetCdnFileHashes) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UploadGetCdnFileHashes:
		m.FileToken = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_UploadGetCdnFileHashes4da54231:
		m.FileToken = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("UploadGetCdnFileHashes Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUploadGetFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UploadGetFile:
		m1 := &InputFileLocation{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Location = m1
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_UploadGetFileb15a9afc:
		flags := dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Flags = flags
		m.Precise = flags&(1<<0) != 0
		m3 := &InputFileLocation{}
		err = m3.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Location = m3
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("UploadGetFile Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUploadGetFileHashes) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UploadGetFileHashes:
		m1 := &InputFileLocation{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Location = m1
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("UploadGetFileHashes Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUploadGetWebFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UploadGetWebFile:
		m1 := &InputWebFileLocation{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Location = m1
		m.Offset = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Limit = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("UploadGetWebFile Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUploadReuploadCdnFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UploadReuploadCdnFile:
		m.FileToken = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.RequestToken = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	case Cmd_UploadReuploadCdnFile9b2754a8:
		m.FileToken = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.RequestToken = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("UploadReuploadCdnFile Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUploadSaveBigFilePart) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UploadSaveBigFilePart:
		m.FileId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.FilePart = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.FileTotalParts = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Bytes = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("UploadSaveBigFilePart Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUploadSaveFilePart) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UploadSaveFilePart:
		m.FileId = dBuf.Long()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.FilePart = dBuf.Int()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

		m.Bytes = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("UploadSaveFilePart Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUsersGetFullUser) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UsersGetFullUser:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1

	default:
		log.Errorf("UsersGetFullUser Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUsersGetUsers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UsersGetUsers:
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l1 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Id = make([]*InputUser, l1)
		for i := int32(0); i < l1; i++ {
			m.Id[i] = &InputUser{}
			err = (*m.Id[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("UsersGetUsers Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLUsersSetSecureValueErrors) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_UsersSetSecureValueErrors:
		m1 := &InputUser{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Id = m1
		dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		l2 := dBuf.Int()

		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}
		m.Errors = make([]*SecureValueError, l2)
		for i := int32(0); i < l2; i++ {
			m.Errors[i] = &SecureValueError{}
			err = (*m.Errors[i]).Decode(dBuf, layer)
			if err != nil {
				return err
			}
		}

	default:
		log.Errorf("UsersSetSecureValueErrors Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLWalletGetKeySecretSalt) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_WalletGetKeySecretSalt:
		m1 := &Bool{}
		err = m1.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		m.Revoke = m1

	default:
		log.Errorf("WalletGetKeySecretSalt Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}

func (m *TLWalletSendLiteRequest) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	switch m.Cmd {
	case Cmd_WalletSendLiteRequest:
		m.Body = dBuf.StringBytes()
		if dBuf.GetError() != nil {
			log.Errorf("Decode object:%+v error:%s", m, dBuf.GetError())
			return dBuf.GetError()
		}

	default:
		log.Errorf("WalletSendLiteRequest Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}

	return dBuf.GetError()
}
