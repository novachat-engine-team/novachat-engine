/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2023/06/12 18:03
 * @File : codec_schema_tl_rpc_encode.pb.go
 * @Desc :
 *
 */

package mtproto

import (
	"novachat_engine/pkg/log"
)

func (m *TLAccountAcceptAuthorization) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountAcceptAuthorization, layer)

	switch cmd {
	case Cmd_AccountAcceptAuthorization:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountAcceptAuthorization))
		x.Int(m.BotId)
		x.String(m.Scope)
		x.String(m.PublicKey)
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.ValueHashes)))
		for _, v := range m.ValueHashes {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.Credentials.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountAcceptAuthorization Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountCancelPasswordEmail) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountCancelPasswordEmail, layer)

	switch cmd {
	case Cmd_AccountCancelPasswordEmail:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountCancelPasswordEmail))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountCancelPasswordEmail Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountChangePhone) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountChangePhone, layer)

	switch cmd {
	case Cmd_AccountChangePhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountChangePhone))
		x.String(m.PhoneNumber)
		x.String(m.PhoneCodeHash)
		x.String(m.PhoneCode)

		return x.GetBuf()

	default:
		log.Errorf("AccountChangePhone Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountCheckUsername) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountCheckUsername, layer)

	switch cmd {
	case Cmd_AccountCheckUsername:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountCheckUsername))
		x.String(m.Username)

		return x.GetBuf()

	default:
		log.Errorf("AccountCheckUsername Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountConfirmPasswordEmail) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountConfirmPasswordEmail, layer)

	switch cmd {
	case Cmd_AccountConfirmPasswordEmail:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountConfirmPasswordEmail))
		x.String(m.Code)

		return x.GetBuf()

	default:
		log.Errorf("AccountConfirmPasswordEmail Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountConfirmPhone) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountConfirmPhone, layer)

	switch cmd {
	case Cmd_AccountConfirmPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountConfirmPhone))
		x.String(m.PhoneCodeHash)
		x.String(m.PhoneCode)

		return x.GetBuf()

	default:
		log.Errorf("AccountConfirmPhone Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountCreateTheme) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountCreateTheme, layer)

	switch cmd {
	case Cmd_AccountCreateTheme:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountCreateTheme))
		x.String(m.Slug)
		x.String(m.Title)
		x.Bytes(m.Document.Encode(layer))

		return x.GetBuf()
	case Cmd_AccountCreateTheme8432c21f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountCreateTheme8432c21f))
		// flags
		var flags int32 = 0
		if m.Document != nil {
			flags |= 1 << 2
		}
		if m.Settings != nil {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.String(m.Slug)
		x.String(m.Title)
		if m.Document != nil {
			x.Bytes(m.Document.Encode(layer))
		}
		if m.Settings != nil {
			x.Bytes(m.Settings.Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("AccountCreateTheme Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountDeleteAccount) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountDeleteAccount, layer)

	switch cmd {
	case Cmd_AccountDeleteAccount:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountDeleteAccount))
		x.String(m.Reason)

		return x.GetBuf()

	default:
		log.Errorf("AccountDeleteAccount Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountDeleteSecureValue) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountDeleteSecureValue, layer)

	switch cmd {
	case Cmd_AccountDeleteSecureValue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountDeleteSecureValue))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Types)))
		for _, v := range m.Types {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("AccountDeleteSecureValue Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountFinishTakeoutSession) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountFinishTakeoutSession, layer)

	switch cmd {
	case Cmd_AccountFinishTakeoutSession:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountFinishTakeoutSession))
		// flags
		var flags int32 = 0
		if m.Success == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("AccountFinishTakeoutSession Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetAccountTTL) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetAccountTTL, layer)

	switch cmd {
	case Cmd_AccountGetAccountTTL:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetAccountTTL))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountGetAccountTTL Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetAllSecureValues) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetAllSecureValues, layer)

	switch cmd {
	case Cmd_AccountGetAllSecureValues:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetAllSecureValues))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountGetAllSecureValues Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetAuthorizationForm) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetAuthorizationForm, layer)

	switch cmd {
	case Cmd_AccountGetAuthorizationForm:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetAuthorizationForm))
		x.Int(m.BotId)
		x.String(m.Scope)
		x.String(m.PublicKey)

		return x.GetBuf()

	default:
		log.Errorf("AccountGetAuthorizationForm Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetAuthorizations) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetAuthorizations, layer)

	switch cmd {
	case Cmd_AccountGetAuthorizations:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetAuthorizations))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountGetAuthorizations Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetAutoDownloadSettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetAutoDownloadSettings, layer)

	switch cmd {
	case Cmd_AccountGetAutoDownloadSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetAutoDownloadSettings))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountGetAutoDownloadSettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetContactSignUpNotification) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetContactSignUpNotification, layer)

	switch cmd {
	case Cmd_AccountGetContactSignUpNotification:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetContactSignUpNotification))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountGetContactSignUpNotification Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetContentSettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetContentSettings, layer)

	switch cmd {
	case Cmd_AccountGetContentSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetContentSettings))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountGetContentSettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetGlobalPrivacySettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetGlobalPrivacySettings, layer)

	switch cmd {
	case Cmd_AccountGetGlobalPrivacySettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetGlobalPrivacySettings))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountGetGlobalPrivacySettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetMultiWallPapers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetMultiWallPapers, layer)

	switch cmd {
	case Cmd_AccountGetMultiWallPapers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetMultiWallPapers))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Wallpapers)))
		for _, v := range m.Wallpapers {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("AccountGetMultiWallPapers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetNotifyExceptions) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetNotifyExceptions, layer)

	switch cmd {
	case Cmd_AccountGetNotifyExceptions:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetNotifyExceptions))
		// flags
		var flags int32 = 0
		if m.CompareSound == true {
			flags |= 1 << 1
		}
		if m.Peer != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.Peer != nil {
			x.Bytes(m.Peer.Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("AccountGetNotifyExceptions Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetNotifySettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetNotifySettings, layer)

	switch cmd {
	case Cmd_AccountGetNotifySettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetNotifySettings))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountGetNotifySettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetPassword) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetPassword, layer)

	switch cmd {
	case Cmd_AccountGetPassword:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetPassword))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountGetPassword Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetPasswordSettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetPasswordSettings, layer)

	switch cmd {
	case Cmd_AccountGetPasswordSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetPasswordSettings))
		x.StringBytes(m.CurrentPasswordHash)

		return x.GetBuf()
	case Cmd_AccountGetPasswordSettings9cd4eaf9:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetPasswordSettings9cd4eaf9))
		x.Bytes(m.Password.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountGetPasswordSettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetPrivacy) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetPrivacy, layer)

	switch cmd {
	case Cmd_AccountGetPrivacy:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetPrivacy))
		x.Bytes(m.Key.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountGetPrivacy Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetSecureValue) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetSecureValue, layer)

	switch cmd {
	case Cmd_AccountGetSecureValue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetSecureValue))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Types)))
		for _, v := range m.Types {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("AccountGetSecureValue Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetTheme) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetTheme, layer)

	switch cmd {
	case Cmd_AccountGetTheme:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetTheme))
		x.String(m.Format)
		x.Bytes(m.Theme.Encode(layer))
		x.Long(m.DocumentId)

		return x.GetBuf()

	default:
		log.Errorf("AccountGetTheme Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetThemes) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetThemes, layer)

	switch cmd {
	case Cmd_AccountGetThemes:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetThemes))
		x.String(m.Format)
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("AccountGetThemes Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetTmpPassword) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetTmpPassword, layer)

	switch cmd {
	case Cmd_AccountGetTmpPassword:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetTmpPassword))
		x.StringBytes(m.PasswordHash)
		x.Int(m.Period)

		return x.GetBuf()
	case Cmd_AccountGetTmpPassword449e0b51:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetTmpPassword449e0b51))
		x.Bytes(m.Password.Encode(layer))
		x.Int(m.Period)

		return x.GetBuf()

	default:
		log.Errorf("AccountGetTmpPassword Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetWallPaper) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetWallPaper, layer)

	switch cmd {
	case Cmd_AccountGetWallPaper:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetWallPaper))
		x.Bytes(m.Wallpaper.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountGetWallPaper Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetWallPapers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetWallPapers, layer)

	switch cmd {
	case Cmd_AccountGetWallPapers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetWallPapers))
		_ = m

		return x.GetBuf()
	case Cmd_AccountGetWallPapersaabb1763:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetWallPapersaabb1763))
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("AccountGetWallPapers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountGetWebAuthorizations) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountGetWebAuthorizations, layer)

	switch cmd {
	case Cmd_AccountGetWebAuthorizations:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountGetWebAuthorizations))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountGetWebAuthorizations Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountInitTakeoutSession) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountInitTakeoutSession, layer)

	switch cmd {
	case Cmd_AccountInitTakeoutSession:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountInitTakeoutSession))
		// flags
		var flags int32 = 0
		if m.Contacts == true {
			flags |= 1 << 0
		}
		if m.MessageUsers == true {
			flags |= 1 << 1
		}
		if m.MessageChats == true {
			flags |= 1 << 2
		}
		if m.MessageMegagroups == true {
			flags |= 1 << 3
		}
		if m.MessageChannels == true {
			flags |= 1 << 4
		}
		if m.Files == true {
			flags |= 1 << 5
		}
		if m.FileMaxSize != 0 {
			flags |= 1 << 5
		}
		x.Int(flags)

		if m.FileMaxSize != 0 {
			x.Int(m.FileMaxSize)
		}

		return x.GetBuf()

	default:
		log.Errorf("AccountInitTakeoutSession Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountInstallTheme) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountInstallTheme, layer)

	switch cmd {
	case Cmd_AccountInstallTheme:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountInstallTheme))
		x.String(m.Format)
		x.Bytes(m.Theme.Encode(layer))

		return x.GetBuf()
	case Cmd_AccountInstallTheme7ae43737:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountInstallTheme7ae43737))
		// flags
		var flags int32 = 0
		if m.Dark == true {
			flags |= 1 << 0
		}
		if len(m.Format) > 0 {
			flags |= 1 << 1
		}
		if m.Theme != nil {
			flags |= 1 << 1
		}
		x.Int(flags)

		if len(m.Format) > 0 {
			x.String(m.Format)
		}
		if m.Theme != nil {
			x.Bytes(m.Theme.Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("AccountInstallTheme Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountInstallWallPaper) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountInstallWallPaper, layer)

	switch cmd {
	case Cmd_AccountInstallWallPaper:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountInstallWallPaper))
		x.Bytes(m.Wallpaper.Encode(layer))

		return x.GetBuf()
	case Cmd_AccountInstallWallPaperfeed5769:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountInstallWallPaperfeed5769))
		x.Bytes(m.Wallpaper.Encode(layer))
		x.Bytes(m.Settings.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountInstallWallPaper Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountRegisterDevice) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountRegisterDevice, layer)

	switch cmd {
	case Cmd_AccountRegisterDevice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountRegisterDevice))
		x.Int(m.TokenType)
		x.String(m.Token)

		return x.GetBuf()
	case Cmd_AccountRegisterDevice446c712c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountRegisterDevice446c712c))
		x.Int(m.TokenType)
		x.String(m.Token)
		x.String(m.DeviceModel)
		x.String(m.SystemVersion)
		x.String(m.AppVersion)
		x.Bytes(m.AppSandbox.Encode(layer))
		x.String(m.LangCode)

		return x.GetBuf()
	case Cmd_AccountRegisterDevice5cbea590:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountRegisterDevice5cbea590))
		x.Int(m.TokenType)
		x.String(m.Token)
		x.Bytes(m.AppSandbox.Encode(layer))
		x.StringBytes(m.Secret)
		x.VectorInt(m.OtherUids)

		return x.GetBuf()
	case Cmd_AccountRegisterDevice68976c6f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountRegisterDevice68976c6f))
		// flags
		var flags int32 = 0
		if m.NoMuted == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.TokenType)
		x.String(m.Token)
		x.Bytes(m.AppSandbox.Encode(layer))
		x.StringBytes(m.Secret)
		x.VectorInt(m.OtherUids)

		return x.GetBuf()

	default:
		log.Errorf("AccountRegisterDevice Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountReportPeer) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountReportPeer, layer)

	switch cmd {
	case Cmd_AccountReportPeer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountReportPeer))
		x.Bytes(m.Peer.Encode(layer))
		x.Bytes(m.Reason.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountReportPeer Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountResendPasswordEmail) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountResendPasswordEmail, layer)

	switch cmd {
	case Cmd_AccountResendPasswordEmail:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountResendPasswordEmail))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountResendPasswordEmail Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountResetAuthorization) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountResetAuthorization, layer)

	switch cmd {
	case Cmd_AccountResetAuthorization:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountResetAuthorization))
		x.Long(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("AccountResetAuthorization Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountResetNotifySettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountResetNotifySettings, layer)

	switch cmd {
	case Cmd_AccountResetNotifySettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountResetNotifySettings))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountResetNotifySettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountResetWallPapers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountResetWallPapers, layer)

	switch cmd {
	case Cmd_AccountResetWallPapers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountResetWallPapers))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountResetWallPapers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountResetWebAuthorization) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountResetWebAuthorization, layer)

	switch cmd {
	case Cmd_AccountResetWebAuthorization:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountResetWebAuthorization))
		x.Long(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("AccountResetWebAuthorization Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountResetWebAuthorizations) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountResetWebAuthorizations, layer)

	switch cmd {
	case Cmd_AccountResetWebAuthorizations:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountResetWebAuthorizations))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AccountResetWebAuthorizations Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSaveAutoDownloadSettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSaveAutoDownloadSettings, layer)

	switch cmd {
	case Cmd_AccountSaveAutoDownloadSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSaveAutoDownloadSettings))
		// flags
		var flags int32 = 0
		if m.Low == true {
			flags |= 1 << 0
		}
		if m.High == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.Settings.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountSaveAutoDownloadSettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSaveSecureValue) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSaveSecureValue, layer)

	switch cmd {
	case Cmd_AccountSaveSecureValue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSaveSecureValue))
		x.Bytes(m.Value.Encode(layer))
		x.Long(m.SecureSecretId)

		return x.GetBuf()

	default:
		log.Errorf("AccountSaveSecureValue Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSaveTheme) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSaveTheme, layer)

	switch cmd {
	case Cmd_AccountSaveTheme:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSaveTheme))
		x.Bytes(m.Theme.Encode(layer))
		x.Bytes(m.Unsave.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountSaveTheme Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSaveWallPaper) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSaveWallPaper, layer)

	switch cmd {
	case Cmd_AccountSaveWallPaper:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSaveWallPaper))
		x.Bytes(m.Wallpaper.Encode(layer))
		x.Bytes(m.Unsave.Encode(layer))

		return x.GetBuf()
	case Cmd_AccountSaveWallPaper6c5a5b37:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSaveWallPaper6c5a5b37))
		x.Bytes(m.Wallpaper.Encode(layer))
		x.Bytes(m.Unsave.Encode(layer))
		x.Bytes(m.Settings.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountSaveWallPaper Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSendChangePhoneCode) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSendChangePhoneCode, layer)

	switch cmd {
	case Cmd_AccountSendChangePhoneCode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSendChangePhoneCode))
		// flags
		var flags int32 = 0
		if m.AllowFlashcall == true {
			flags |= 1 << 0
		}
		if m.CurrentNumber != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.PhoneNumber)
		if m.CurrentNumber != nil {
			x.Bytes(m.CurrentNumber.Encode(layer))
		}

		return x.GetBuf()
	case Cmd_AccountSendChangePhoneCode82574ae5:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSendChangePhoneCode82574ae5))
		x.String(m.PhoneNumber)
		x.Bytes(m.Settings.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountSendChangePhoneCode Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSendConfirmPhoneCode) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSendConfirmPhoneCode, layer)

	switch cmd {
	case Cmd_AccountSendConfirmPhoneCode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSendConfirmPhoneCode))
		// flags
		var flags int32 = 0
		if m.AllowFlashcall == true {
			flags |= 1 << 0
		}
		if m.CurrentNumber != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.Hash)
		if m.CurrentNumber != nil {
			x.Bytes(m.CurrentNumber.Encode(layer))
		}

		return x.GetBuf()
	case Cmd_AccountSendConfirmPhoneCode1b3faa88:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSendConfirmPhoneCode1b3faa88))
		x.String(m.Hash)
		x.Bytes(m.Settings.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountSendConfirmPhoneCode Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSendVerifyEmailCode) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSendVerifyEmailCode, layer)

	switch cmd {
	case Cmd_AccountSendVerifyEmailCode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSendVerifyEmailCode))
		x.String(m.Email)

		return x.GetBuf()

	default:
		log.Errorf("AccountSendVerifyEmailCode Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSendVerifyPhoneCode) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSendVerifyPhoneCode, layer)

	switch cmd {
	case Cmd_AccountSendVerifyPhoneCode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSendVerifyPhoneCode))
		// flags
		var flags int32 = 0
		if m.AllowFlashcall == true {
			flags |= 1 << 0
		}
		if m.CurrentNumber != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.PhoneNumber)
		if m.CurrentNumber != nil {
			x.Bytes(m.CurrentNumber.Encode(layer))
		}

		return x.GetBuf()
	case Cmd_AccountSendVerifyPhoneCodea5a356f9:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSendVerifyPhoneCodea5a356f9))
		x.String(m.PhoneNumber)
		x.Bytes(m.Settings.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountSendVerifyPhoneCode Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSetAccountTTL) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSetAccountTTL, layer)

	switch cmd {
	case Cmd_AccountSetAccountTTL:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSetAccountTTL))
		x.Bytes(m.Ttl.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountSetAccountTTL Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSetContactSignUpNotification) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSetContactSignUpNotification, layer)

	switch cmd {
	case Cmd_AccountSetContactSignUpNotification:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSetContactSignUpNotification))
		x.Bytes(m.Silent.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountSetContactSignUpNotification Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSetContentSettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSetContentSettings, layer)

	switch cmd {
	case Cmd_AccountSetContentSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSetContentSettings))
		// flags
		var flags int32 = 0
		if m.SensitiveEnabled == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("AccountSetContentSettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSetGlobalPrivacySettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSetGlobalPrivacySettings, layer)

	switch cmd {
	case Cmd_AccountSetGlobalPrivacySettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSetGlobalPrivacySettings))
		x.Bytes(m.Settings.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountSetGlobalPrivacySettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSetPrivacy) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountSetPrivacy, layer)

	switch cmd {
	case Cmd_AccountSetPrivacy:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSetPrivacy))
		x.Bytes(m.Key.Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Rules)))
		for _, v := range m.Rules {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("AccountSetPrivacy Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountUnregisterDevice) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountUnregisterDevice, layer)

	switch cmd {
	case Cmd_AccountUnregisterDevice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUnregisterDevice))
		x.Int(m.TokenType)
		x.String(m.Token)

		return x.GetBuf()
	case Cmd_AccountUnregisterDevice3076c4bf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUnregisterDevice3076c4bf))
		x.Int(m.TokenType)
		x.String(m.Token)
		x.VectorInt(m.OtherUids)

		return x.GetBuf()

	default:
		log.Errorf("AccountUnregisterDevice Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountUpdateDeviceLocked) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountUpdateDeviceLocked, layer)

	switch cmd {
	case Cmd_AccountUpdateDeviceLocked:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUpdateDeviceLocked))
		x.Int(m.Period)

		return x.GetBuf()

	default:
		log.Errorf("AccountUpdateDeviceLocked Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountUpdateNotifySettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountUpdateNotifySettings, layer)

	switch cmd {
	case Cmd_AccountUpdateNotifySettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUpdateNotifySettings))
		x.Bytes(m.Peer.Encode(layer))
		x.Bytes(m.Settings.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountUpdateNotifySettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountUpdatePasswordSettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountUpdatePasswordSettings, layer)

	switch cmd {
	case Cmd_AccountUpdatePasswordSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUpdatePasswordSettings))
		x.StringBytes(m.CurrentPasswordHash)
		x.Bytes(m.NewSettings.Encode(layer))

		return x.GetBuf()
	case Cmd_AccountUpdatePasswordSettingsa59b102f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUpdatePasswordSettingsa59b102f))
		x.Bytes(m.Password.Encode(layer))
		x.Bytes(m.NewSettings.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountUpdatePasswordSettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountUpdateProfile) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountUpdateProfile, layer)

	switch cmd {
	case Cmd_AccountUpdateProfile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUpdateProfile))
		// flags
		var flags int32 = 0
		if len(m.FirstName) > 0 {
			flags |= 1 << 0
		}
		if len(m.LastName) > 0 {
			flags |= 1 << 1
		}
		if len(m.About) > 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		if len(m.FirstName) > 0 {
			x.String(m.FirstName)
		}
		if len(m.LastName) > 0 {
			x.String(m.LastName)
		}
		if len(m.About) > 0 {
			x.String(m.About)
		}

		return x.GetBuf()

	default:
		log.Errorf("AccountUpdateProfile Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountUpdateStatus) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountUpdateStatus, layer)

	switch cmd {
	case Cmd_AccountUpdateStatus:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUpdateStatus))
		x.Bytes(m.Offline.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountUpdateStatus Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountUpdateTheme) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountUpdateTheme, layer)

	switch cmd {
	case Cmd_AccountUpdateTheme:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUpdateTheme))
		// flags
		var flags int32 = 0
		if len(m.Slug) > 0 {
			flags |= 1 << 0
		}
		if len(m.Title) > 0 {
			flags |= 1 << 1
		}
		if m.Document != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.Theme.Encode(layer))
		if len(m.Slug) > 0 {
			x.String(m.Slug)
		}
		if len(m.Title) > 0 {
			x.String(m.Title)
		}
		if m.Document != nil {
			x.Bytes(m.Document.Encode(layer))
		}

		return x.GetBuf()
	case Cmd_AccountUpdateTheme3b8ea202:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUpdateTheme3b8ea202))
		// flags
		var flags int32 = 0
		if len(m.Slug) > 0 {
			flags |= 1 << 0
		}
		if len(m.Title) > 0 {
			flags |= 1 << 1
		}
		if m.Document != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.Format)
		x.Bytes(m.Theme.Encode(layer))
		if len(m.Slug) > 0 {
			x.String(m.Slug)
		}
		if len(m.Title) > 0 {
			x.String(m.Title)
		}
		if m.Document != nil {
			x.Bytes(m.Document.Encode(layer))
		}

		return x.GetBuf()
	case Cmd_AccountUpdateTheme5cb367d5:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUpdateTheme5cb367d5))
		// flags
		var flags int32 = 0
		if len(m.Slug) > 0 {
			flags |= 1 << 0
		}
		if len(m.Title) > 0 {
			flags |= 1 << 1
		}
		if m.Document != nil {
			flags |= 1 << 2
		}
		if m.Settings != nil {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.String(m.Format)
		x.Bytes(m.Theme.Encode(layer))
		if len(m.Slug) > 0 {
			x.String(m.Slug)
		}
		if len(m.Title) > 0 {
			x.String(m.Title)
		}
		if m.Document != nil {
			x.Bytes(m.Document.Encode(layer))
		}
		if m.Settings != nil {
			x.Bytes(m.Settings.Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("AccountUpdateTheme Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountUpdateUsername) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountUpdateUsername, layer)

	switch cmd {
	case Cmd_AccountUpdateUsername:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUpdateUsername))
		x.String(m.Username)

		return x.GetBuf()

	default:
		log.Errorf("AccountUpdateUsername Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountUploadTheme) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountUploadTheme, layer)

	switch cmd {
	case Cmd_AccountUploadTheme:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUploadTheme))
		// flags
		var flags int32 = 0
		if m.Thumb != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.File.Encode(layer))
		if m.Thumb != nil {
			x.Bytes(m.Thumb.Encode(layer))
		}
		x.String(m.FileName)
		x.String(m.MimeType)

		return x.GetBuf()

	default:
		log.Errorf("AccountUploadTheme Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountUploadWallPaper) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountUploadWallPaper, layer)

	switch cmd {
	case Cmd_AccountUploadWallPaper:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUploadWallPaper))
		x.Bytes(m.File.Encode(layer))
		x.String(m.MimeType)

		return x.GetBuf()
	case Cmd_AccountUploadWallPaperdd853661:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountUploadWallPaperdd853661))
		x.Bytes(m.File.Encode(layer))
		x.String(m.MimeType)
		x.Bytes(m.Settings.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AccountUploadWallPaper Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountVerifyEmail) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountVerifyEmail, layer)

	switch cmd {
	case Cmd_AccountVerifyEmail:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountVerifyEmail))
		x.String(m.Email)
		x.String(m.Code)

		return x.GetBuf()

	default:
		log.Errorf("AccountVerifyEmail Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountVerifyPhone) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAccountVerifyPhone, layer)

	switch cmd {
	case Cmd_AccountVerifyPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountVerifyPhone))
		x.String(m.PhoneNumber)
		x.String(m.PhoneCodeHash)
		x.String(m.PhoneCode)

		return x.GetBuf()

	default:
		log.Errorf("AccountVerifyPhone Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthAcceptLoginToken) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthAcceptLoginToken, layer)

	switch cmd {
	case Cmd_AuthAcceptLoginToken:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthAcceptLoginToken))
		x.StringBytes(m.Token)

		return x.GetBuf()

	default:
		log.Errorf("AuthAcceptLoginToken Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthBindTempAuthKey) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthBindTempAuthKey, layer)

	switch cmd {
	case Cmd_AuthBindTempAuthKey:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthBindTempAuthKey))
		x.Long(m.PermAuthKeyId)
		x.Long(m.Nonce)
		x.Int(m.ExpiresAt)
		x.StringBytes(m.EncryptedMessage)

		return x.GetBuf()

	default:
		log.Errorf("AuthBindTempAuthKey Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthCancelCode) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthCancelCode, layer)

	switch cmd {
	case Cmd_AuthCancelCode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthCancelCode))
		x.String(m.PhoneNumber)
		x.String(m.PhoneCodeHash)

		return x.GetBuf()

	default:
		log.Errorf("AuthCancelCode Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthCheckPassword) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthCheckPassword, layer)

	switch cmd {
	case Cmd_AuthCheckPassword:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthCheckPassword))
		x.StringBytes(m.PasswordHash)

		return x.GetBuf()
	case Cmd_AuthCheckPasswordd18b4d16:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthCheckPasswordd18b4d16))
		x.Bytes(m.Password.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("AuthCheckPassword Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthCheckPhone) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthCheckPhone, layer)

	switch cmd {
	case Cmd_AuthCheckPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthCheckPhone))
		x.String(m.PhoneNumber)

		return x.GetBuf()

	default:
		log.Errorf("AuthCheckPhone Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthDropTempAuthKeys) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthDropTempAuthKeys, layer)

	switch cmd {
	case Cmd_AuthDropTempAuthKeys:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthDropTempAuthKeys))
		x.VectorLong(m.ExceptAuthKeys)

		return x.GetBuf()

	default:
		log.Errorf("AuthDropTempAuthKeys Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthExportAuthorization) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthExportAuthorization, layer)

	switch cmd {
	case Cmd_AuthExportAuthorization:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthExportAuthorization))
		x.Int(m.DcId)

		return x.GetBuf()

	default:
		log.Errorf("AuthExportAuthorization Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthExportLoginToken) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthExportLoginToken, layer)

	switch cmd {
	case Cmd_AuthExportLoginToken:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthExportLoginToken))
		x.Int(m.ApiId)
		x.String(m.ApiHash)
		x.VectorInt(m.ExceptIds)

		return x.GetBuf()

	default:
		log.Errorf("AuthExportLoginToken Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthImportAuthorization) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthImportAuthorization, layer)

	switch cmd {
	case Cmd_AuthImportAuthorization:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthImportAuthorization))
		x.Int(m.Id)
		x.StringBytes(m.Bytes)

		return x.GetBuf()

	default:
		log.Errorf("AuthImportAuthorization Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthImportBotAuthorization) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthImportBotAuthorization, layer)

	switch cmd {
	case Cmd_AuthImportBotAuthorization:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthImportBotAuthorization))
		x.Int(m.Flags)
		x.Int(m.ApiId)
		x.String(m.ApiHash)
		x.String(m.BotAuthToken)

		return x.GetBuf()

	default:
		log.Errorf("AuthImportBotAuthorization Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthImportLoginToken) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthImportLoginToken, layer)

	switch cmd {
	case Cmd_AuthImportLoginToken:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthImportLoginToken))
		x.StringBytes(m.Token)

		return x.GetBuf()

	default:
		log.Errorf("AuthImportLoginToken Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthLogOut) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthLogOut, layer)

	switch cmd {
	case Cmd_AuthLogOut:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthLogOut))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AuthLogOut Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthRecoverPassword) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthRecoverPassword, layer)

	switch cmd {
	case Cmd_AuthRecoverPassword:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthRecoverPassword))
		x.String(m.Code)

		return x.GetBuf()

	default:
		log.Errorf("AuthRecoverPassword Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthRequestPasswordRecovery) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthRequestPasswordRecovery, layer)

	switch cmd {
	case Cmd_AuthRequestPasswordRecovery:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthRequestPasswordRecovery))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AuthRequestPasswordRecovery Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthResendCode) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthResendCode, layer)

	switch cmd {
	case Cmd_AuthResendCode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthResendCode))
		x.String(m.PhoneNumber)
		x.String(m.PhoneCodeHash)

		return x.GetBuf()

	default:
		log.Errorf("AuthResendCode Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthResetAuthorizations) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthResetAuthorizations, layer)

	switch cmd {
	case Cmd_AuthResetAuthorizations:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthResetAuthorizations))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("AuthResetAuthorizations Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthSendCode) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthSendCode, layer)

	switch cmd {
	case Cmd_AuthSendCode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSendCode))
		// flags
		var flags int32 = 0
		if m.AllowFlashcall == true {
			flags |= 1 << 0
		}
		if m.CurrentNumber != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.PhoneNumber)
		if m.CurrentNumber != nil {
			x.Bytes(m.CurrentNumber.Encode(layer))
		}
		x.Int(m.ApiId)
		x.String(m.ApiHash)

		return x.GetBuf()
	case Cmd_AuthSendCodea677244f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSendCodea677244f))
		x.String(m.PhoneNumber)
		x.Int(m.ApiId)
		x.String(m.ApiHash)
		x.Bytes(m.Settings.Encode(layer))

		return x.GetBuf()
	case Cmd_AuthSendCodeccfd70cf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSendCodeccfd70cf))
		// flags
		var flags int32 = 0
		if m.AllowFlashcall == true {
			flags |= 1 << 0
		}
		if m.CurrentNumber != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.PhoneNumber)
		if m.CurrentNumber != nil {
			x.Bytes(m.CurrentNumber.Encode(layer))
		}
		x.Int(m.ApiId)
		x.String(m.ApiHash)
		x.String(m.LangCode)

		return x.GetBuf()

	default:
		log.Errorf("AuthSendCode Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthSendInvites) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthSendInvites, layer)

	switch cmd {
	case Cmd_AuthSendInvites:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSendInvites))
		x.VectorString(m.PhoneNumbers)

		x.String(m.Message)

		return x.GetBuf()

	default:
		log.Errorf("AuthSendInvites Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthSignIn) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthSignIn, layer)

	switch cmd {
	case Cmd_AuthSignIn:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSignIn))
		x.String(m.PhoneNumber)
		x.String(m.PhoneCodeHash)
		x.String(m.PhoneCode)

		return x.GetBuf()

	default:
		log.Errorf("AuthSignIn Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthSignUp) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassAuthSignUp, layer)

	switch cmd {
	case Cmd_AuthSignUp:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSignUp))
		x.String(m.PhoneNumber)
		x.String(m.PhoneCodeHash)
		x.String(m.PhoneCode)
		x.String(m.FirstName)
		x.String(m.LastName)

		return x.GetBuf()
	case Cmd_AuthSignUp80eee427:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSignUp80eee427))
		x.String(m.PhoneNumber)
		x.String(m.PhoneCodeHash)
		x.String(m.FirstName)
		x.String(m.LastName)

		return x.GetBuf()

	default:
		log.Errorf("AuthSignUp Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotsAnswerWebhookJSONQuery) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassBotsAnswerWebhookJSONQuery, layer)

	switch cmd {
	case Cmd_BotsAnswerWebhookJSONQuery:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotsAnswerWebhookJSONQuery))
		x.Long(m.QueryId)
		x.Bytes(m.Data.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("BotsAnswerWebhookJSONQuery Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotsSendCustomRequest) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassBotsSendCustomRequest, layer)

	switch cmd {
	case Cmd_BotsSendCustomRequest:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotsSendCustomRequest))
		x.String(m.CustomMethod)
		x.Bytes(m.Params.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("BotsSendCustomRequest Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotsSetBotCommands) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassBotsSetBotCommands, layer)

	switch cmd {
	case Cmd_BotsSetBotCommands:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotsSetBotCommands))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Commands)))
		for _, v := range m.Commands {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("BotsSetBotCommands Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsCheckUsername) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsCheckUsername, layer)

	switch cmd {
	case Cmd_ChannelsCheckUsername:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsCheckUsername))
		x.Bytes(m.Channel.Encode(layer))
		x.String(m.Username)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsCheckUsername Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsCreateChannel) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsCreateChannel, layer)

	switch cmd {
	case Cmd_ChannelsCreateChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsCreateChannel))
		// flags
		var flags int32 = 0
		if m.Broadcast == true {
			flags |= 1 << 0
		}
		if m.Megagroup == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.String(m.Title)
		x.String(m.About)

		return x.GetBuf()
	case Cmd_ChannelsCreateChannel3d5fb10f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsCreateChannel3d5fb10f))
		// flags
		var flags int32 = 0
		if m.Broadcast == true {
			flags |= 1 << 0
		}
		if m.Megagroup == true {
			flags |= 1 << 1
		}
		if m.GeoPoint != nil {
			flags |= 1 << 2
		}
		if len(m.Address) > 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.Title)
		x.String(m.About)
		if m.GeoPoint != nil {
			x.Bytes(m.GeoPoint.Encode(layer))
		}
		if len(m.Address) > 0 {
			x.String(m.Address)
		}

		return x.GetBuf()

	default:
		log.Errorf("ChannelsCreateChannel Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsDeleteChannel) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsDeleteChannel, layer)

	switch cmd {
	case Cmd_ChannelsDeleteChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsDeleteChannel))
		x.Bytes(m.Channel.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsDeleteChannel Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsDeleteHistory) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsDeleteHistory, layer)

	switch cmd {
	case Cmd_ChannelsDeleteHistory:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsDeleteHistory))
		x.Bytes(m.Channel.Encode(layer))
		x.Int(m.MaxId)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsDeleteHistory Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsDeleteMessages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsDeleteMessages, layer)

	switch cmd {
	case Cmd_ChannelsDeleteMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsDeleteMessages))
		x.Bytes(m.Channel.Encode(layer))
		x.VectorInt(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsDeleteMessages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsDeleteUserHistory) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsDeleteUserHistory, layer)

	switch cmd {
	case Cmd_ChannelsDeleteUserHistory:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsDeleteUserHistory))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsDeleteUserHistory Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsEditAbout) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsEditAbout, layer)

	switch cmd {
	case Cmd_ChannelsEditAbout:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsEditAbout))
		x.Bytes(m.Channel.Encode(layer))
		x.String(m.About)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsEditAbout Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsEditAdmin) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsEditAdmin, layer)

	switch cmd {
	case Cmd_ChannelsEditAdmin:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsEditAdmin))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))
		x.Bytes(m.AdminRights20B8821471.Encode(layer))

		return x.GetBuf()
	case Cmd_ChannelsEditAdmin70f893ba:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsEditAdmin70f893ba))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))
		x.Bytes(m.AdminRights70F893BA93.Encode(layer))

		return x.GetBuf()
	case Cmd_ChannelsEditAdmind33c8902:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsEditAdmind33c8902))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))
		x.Bytes(m.AdminRights70F893BA93.Encode(layer))
		x.String(m.Rank)

		return x.GetBuf()
	case Cmd_ChannelsEditAdmineb7611d0:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsEditAdmineb7611d0))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))
		x.Bytes(m.Role.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsEditAdmin Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsEditBanned) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsEditBanned, layer)

	switch cmd {
	case Cmd_ChannelsEditBanned:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsEditBanned))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))
		x.Bytes(m.BannedRightsBFD915CD71.Encode(layer))

		return x.GetBuf()
	case Cmd_ChannelsEditBanned72796912:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsEditBanned72796912))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))
		x.Bytes(m.BannedRights7279691293.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsEditBanned Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsEditCreator) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsEditCreator, layer)

	switch cmd {
	case Cmd_ChannelsEditCreator:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsEditCreator))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))
		x.Bytes(m.Password.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsEditCreator Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsEditLocation) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsEditLocation, layer)

	switch cmd {
	case Cmd_ChannelsEditLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsEditLocation))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.GeoPoint.Encode(layer))
		x.String(m.Address)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsEditLocation Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsEditPhoto) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsEditPhoto, layer)

	switch cmd {
	case Cmd_ChannelsEditPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsEditPhoto))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.Photo.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsEditPhoto Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsEditTitle) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsEditTitle, layer)

	switch cmd {
	case Cmd_ChannelsEditTitle:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsEditTitle))
		x.Bytes(m.Channel.Encode(layer))
		x.String(m.Title)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsEditTitle Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsExportInvite) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsExportInvite, layer)

	switch cmd {
	case Cmd_ChannelsExportInvite:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsExportInvite))
		x.Bytes(m.Channel.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsExportInvite Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsExportMessageLink) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsExportMessageLink, layer)

	switch cmd {
	case Cmd_ChannelsExportMessageLink:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsExportMessageLink))
		x.Bytes(m.Channel.Encode(layer))
		x.Int(m.Id)

		return x.GetBuf()
	case Cmd_ChannelsExportMessageLinkceb77163:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsExportMessageLinkceb77163))
		x.Bytes(m.Channel.Encode(layer))
		x.Int(m.Id)
		x.Bytes(m.GroupedCEB7716385.Encode(layer))

		return x.GetBuf()
	case Cmd_ChannelsExportMessageLinke63fadeb:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsExportMessageLinke63fadeb))
		// flags
		var flags int32 = 0
		if m.GroupedE63FADEB119 == true {
			flags |= 1 << 0
		}
		if m.Thread == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.Channel.Encode(layer))
		x.Int(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsExportMessageLink Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetAdminLog) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetAdminLog, layer)

	switch cmd {
	case Cmd_ChannelsGetAdminLog:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetAdminLog))
		// flags
		var flags int32 = 0
		if m.EventsFilter != nil {
			flags |= 1 << 0
		}
		if len(m.Admins) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.Channel.Encode(layer))
		x.String(m.Q)
		if m.EventsFilter != nil {
			x.Bytes(m.EventsFilter.Encode(layer))
		}
		if len(m.Admins) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Admins)))
			for _, v := range m.Admins {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		x.Long(m.MaxId)
		x.Long(m.MinId)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetAdminLog Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetAdminedPublicChannels) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetAdminedPublicChannels, layer)

	switch cmd {
	case Cmd_ChannelsGetAdminedPublicChannels:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetAdminedPublicChannels))
		_ = m

		return x.GetBuf()
	case Cmd_ChannelsGetAdminedPublicChannelsf8b036af:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetAdminedPublicChannelsf8b036af))
		// flags
		var flags int32 = 0
		if m.ByLocation == true {
			flags |= 1 << 0
		}
		if m.CheckLimit == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetAdminedPublicChannels Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetBroadcastsForDiscussion) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetBroadcastsForDiscussion, layer)

	switch cmd {
	case Cmd_ChannelsGetBroadcastsForDiscussion:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetBroadcastsForDiscussion))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetBroadcastsForDiscussion Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetChannels) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetChannels, layer)

	switch cmd {
	case Cmd_ChannelsGetChannels:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetChannels))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Id)))
		for _, v := range m.Id {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetChannels Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetDialogs) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetDialogs, layer)

	switch cmd {
	case Cmd_ChannelsGetDialogs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetDialogs))
		x.Int(m.Offset)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetDialogs Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetFullChannel) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetFullChannel, layer)

	switch cmd {
	case Cmd_ChannelsGetFullChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetFullChannel))
		x.Bytes(m.Channel.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetFullChannel Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetGroupsForDiscussion) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetGroupsForDiscussion, layer)

	switch cmd {
	case Cmd_ChannelsGetGroupsForDiscussion:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetGroupsForDiscussion))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetGroupsForDiscussion Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetImportantHistory) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetImportantHistory, layer)

	switch cmd {
	case Cmd_ChannelsGetImportantHistory:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetImportantHistory))
		x.Bytes(m.Channel.Encode(layer))
		x.Int(m.OffsetId)
		x.Int(m.OffsetDate)
		x.Int(m.AddOffset)
		x.Int(m.Limit)
		x.Int(m.MaxId)
		x.Int(m.MinId)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetImportantHistory Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetInactiveChannels) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetInactiveChannels, layer)

	switch cmd {
	case Cmd_ChannelsGetInactiveChannels:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetInactiveChannels))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetInactiveChannels Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetLeftChannels) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetLeftChannels, layer)

	switch cmd {
	case Cmd_ChannelsGetLeftChannels:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetLeftChannels))
		x.Int(m.Offset)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetLeftChannels Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetMessages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetMessages, layer)

	switch cmd {
	case Cmd_ChannelsGetMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetMessages))
		x.Bytes(m.Channel.Encode(layer))
		x.VectorInt(m.Id93D7B34771)

		return x.GetBuf()
	case Cmd_ChannelsGetMessagesad8c9a23:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetMessagesad8c9a23))
		x.Bytes(m.Channel.Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.IdAD8C9A2385)))
		for _, v := range m.IdAD8C9A2385 {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetMessages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetParticipant) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetParticipant, layer)

	switch cmd {
	case Cmd_ChannelsGetParticipant:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetParticipant))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetParticipant Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsGetParticipants) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsGetParticipants, layer)

	switch cmd {
	case Cmd_ChannelsGetParticipants:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetParticipants))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.Filter.Encode(layer))
		x.Int(m.Offset)
		x.Int(m.Limit)

		return x.GetBuf()
	case Cmd_ChannelsGetParticipants123e05e9:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsGetParticipants123e05e9))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.Filter.Encode(layer))
		x.Int(m.Offset)
		x.Int(m.Limit)
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsGetParticipants Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsInviteToChannel) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsInviteToChannel, layer)

	switch cmd {
	case Cmd_ChannelsInviteToChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsInviteToChannel))
		x.Bytes(m.Channel.Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Users)))
		for _, v := range m.Users {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("ChannelsInviteToChannel Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsJoinChannel) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsJoinChannel, layer)

	switch cmd {
	case Cmd_ChannelsJoinChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsJoinChannel))
		x.Bytes(m.Channel.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsJoinChannel Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsKickFromChannel) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsKickFromChannel, layer)

	switch cmd {
	case Cmd_ChannelsKickFromChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsKickFromChannel))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))
		x.Bytes(m.Kicked.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsKickFromChannel Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsLeaveChannel) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsLeaveChannel, layer)

	switch cmd {
	case Cmd_ChannelsLeaveChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsLeaveChannel))
		x.Bytes(m.Channel.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsLeaveChannel Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsReadHistory) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsReadHistory, layer)

	switch cmd {
	case Cmd_ChannelsReadHistory:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsReadHistory))
		x.Bytes(m.Channel.Encode(layer))
		x.Int(m.MaxId)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsReadHistory Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsReadMessageContents) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsReadMessageContents, layer)

	switch cmd {
	case Cmd_ChannelsReadMessageContents:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsReadMessageContents))
		x.Bytes(m.Channel.Encode(layer))
		x.VectorInt(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsReadMessageContents Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsReportSpam) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsReportSpam, layer)

	switch cmd {
	case Cmd_ChannelsReportSpam:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsReportSpam))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))
		x.VectorInt(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsReportSpam Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsSetDiscussionGroup) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsSetDiscussionGroup, layer)

	switch cmd {
	case Cmd_ChannelsSetDiscussionGroup:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsSetDiscussionGroup))
		x.Bytes(m.Broadcast.Encode(layer))
		x.Bytes(m.Group.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsSetDiscussionGroup Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsSetStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsSetStickers, layer)

	switch cmd {
	case Cmd_ChannelsSetStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsSetStickers))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.Stickerset.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsSetStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsToggleComments) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsToggleComments, layer)

	switch cmd {
	case Cmd_ChannelsToggleComments:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsToggleComments))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.Enabled.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsToggleComments Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsToggleInvites) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsToggleInvites, layer)

	switch cmd {
	case Cmd_ChannelsToggleInvites:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsToggleInvites))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.Enabled.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsToggleInvites Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsTogglePreHistoryHidden) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsTogglePreHistoryHidden, layer)

	switch cmd {
	case Cmd_ChannelsTogglePreHistoryHidden:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsTogglePreHistoryHidden))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.Enabled.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsTogglePreHistoryHidden Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsToggleSignatures) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsToggleSignatures, layer)

	switch cmd {
	case Cmd_ChannelsToggleSignatures:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsToggleSignatures))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.Enabled.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ChannelsToggleSignatures Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsToggleSlowMode) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsToggleSlowMode, layer)

	switch cmd {
	case Cmd_ChannelsToggleSlowMode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsToggleSlowMode))
		x.Bytes(m.Channel.Encode(layer))
		x.Int(m.Seconds)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsToggleSlowMode Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsUpdatePinnedMessage) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsUpdatePinnedMessage, layer)

	switch cmd {
	case Cmd_ChannelsUpdatePinnedMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsUpdatePinnedMessage))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Channel.Encode(layer))
		x.Int(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsUpdatePinnedMessage Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsUpdateUsername) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassChannelsUpdateUsername, layer)

	switch cmd {
	case Cmd_ChannelsUpdateUsername:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsUpdateUsername))
		x.Bytes(m.Channel.Encode(layer))
		x.String(m.Username)

		return x.GetBuf()

	default:
		log.Errorf("ChannelsUpdateUsername Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsAcceptContact) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsAcceptContact, layer)

	switch cmd {
	case Cmd_ContactsAcceptContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsAcceptContact))
		x.Bytes(m.Id.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ContactsAcceptContact Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsAddContact) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsAddContact, layer)

	switch cmd {
	case Cmd_ContactsAddContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsAddContact))
		// flags
		var flags int32 = 0
		if m.AddPhonePrivacyException == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Id.Encode(layer))
		x.String(m.FirstName)
		x.String(m.LastName)
		x.String(m.Phone)

		return x.GetBuf()

	default:
		log.Errorf("ContactsAddContact Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsBlock) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsBlock, layer)

	switch cmd {
	case Cmd_ContactsBlock:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsBlock))
		x.Bytes(m.Id332B49FC71.Encode(layer))

		return x.GetBuf()
	case Cmd_ContactsBlock68cc1411:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsBlock68cc1411))
		x.Bytes(m.Id68CC1411119.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ContactsBlock Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsBlockFromReplies) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsBlockFromReplies, layer)

	switch cmd {
	case Cmd_ContactsBlockFromReplies:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsBlockFromReplies))
		// flags
		var flags int32 = 0
		if m.DeleteMessage == true {
			flags |= 1 << 0
		}
		if m.DeleteHistory == true {
			flags |= 1 << 1
		}
		if m.ReportSpam == true {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.MsgId)

		return x.GetBuf()

	default:
		log.Errorf("ContactsBlockFromReplies Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsDeleteByPhones) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsDeleteByPhones, layer)

	switch cmd {
	case Cmd_ContactsDeleteByPhones:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsDeleteByPhones))
		x.VectorString(m.Phones)

		return x.GetBuf()

	default:
		log.Errorf("ContactsDeleteByPhones Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsDeleteContact) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsDeleteContact, layer)

	switch cmd {
	case Cmd_ContactsDeleteContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsDeleteContact))
		x.Bytes(m.Id.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ContactsDeleteContact Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsDeleteContacts) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsDeleteContacts, layer)

	switch cmd {
	case Cmd_ContactsDeleteContacts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsDeleteContacts))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Id)))
		for _, v := range m.Id {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_ContactsDeleteContacts96a0e00:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsDeleteContacts96a0e00))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Id)))
		for _, v := range m.Id {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("ContactsDeleteContacts Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsExportCard) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsExportCard, layer)

	switch cmd {
	case Cmd_ContactsExportCard:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsExportCard))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("ContactsExportCard Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsGetBlocked) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsGetBlocked, layer)

	switch cmd {
	case Cmd_ContactsGetBlocked:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsGetBlocked))
		x.Int(m.Offset)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("ContactsGetBlocked Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsGetContactIDs) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsGetContactIDs, layer)

	switch cmd {
	case Cmd_ContactsGetContactIDs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsGetContactIDs))
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("ContactsGetContactIDs Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsGetContacts) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsGetContacts, layer)

	switch cmd {
	case Cmd_ContactsGetContacts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsGetContacts))
		x.Int(m.HashC023849F71)

		return x.GetBuf()
	case Cmd_ContactsGetContacts22c6aa08:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsGetContacts22c6aa08))
		x.String(m.Hash22C6AA0851)

		return x.GetBuf()

	default:
		log.Errorf("ContactsGetContacts Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsGetLocated) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsGetLocated, layer)

	switch cmd {
	case Cmd_ContactsGetLocated:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsGetLocated))
		x.Bytes(m.GeoPoint.Encode(layer))

		return x.GetBuf()
	case Cmd_ContactsGetLocatedd348bc44:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsGetLocatedd348bc44))
		// flags
		var flags int32 = 0
		if m.Background == true {
			flags |= 1 << 1
		}
		if m.SelfExpires != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GeoPoint.Encode(layer))
		if m.SelfExpires != 0 {
			x.Int(m.SelfExpires)
		}

		return x.GetBuf()

	default:
		log.Errorf("ContactsGetLocated Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsGetSaved) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsGetSaved, layer)

	switch cmd {
	case Cmd_ContactsGetSaved:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsGetSaved))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("ContactsGetSaved Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsGetStatuses) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsGetStatuses, layer)

	switch cmd {
	case Cmd_ContactsGetStatuses:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsGetStatuses))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("ContactsGetStatuses Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsGetTopPeers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsGetTopPeers, layer)

	switch cmd {
	case Cmd_ContactsGetTopPeers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsGetTopPeers))
		// flags
		var flags int32 = 0
		if m.Correspondents == true {
			flags |= 1 << 0
		}
		if m.BotsPm == true {
			flags |= 1 << 1
		}
		if m.BotsInline == true {
			flags |= 1 << 2
		}
		if m.PhoneCalls == true {
			flags |= 1 << 3
		}
		if m.Groups == true {
			flags |= 1 << 10
		}
		if m.Channels == true {
			flags |= 1 << 15
		}
		x.Int(flags)

		x.Int(m.Offset)
		x.Int(m.Limit)
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("ContactsGetTopPeers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsImportCard) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsImportCard, layer)

	switch cmd {
	case Cmd_ContactsImportCard:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsImportCard))
		x.VectorInt(m.ExportCard)

		return x.GetBuf()

	default:
		log.Errorf("ContactsImportCard Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsImportContacts) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsImportContacts, layer)

	switch cmd {
	case Cmd_ContactsImportContacts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsImportContacts))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Contacts)))
		for _, v := range m.Contacts {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_ContactsImportContactsda30b32d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsImportContactsda30b32d))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Contacts)))
		for _, v := range m.Contacts {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.Replace.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ContactsImportContacts Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsResetSaved) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsResetSaved, layer)

	switch cmd {
	case Cmd_ContactsResetSaved:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsResetSaved))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("ContactsResetSaved Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsResetTopPeerRating) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsResetTopPeerRating, layer)

	switch cmd {
	case Cmd_ContactsResetTopPeerRating:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsResetTopPeerRating))
		x.Bytes(m.Category.Encode(layer))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ContactsResetTopPeerRating Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsResolveUsername) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsResolveUsername, layer)

	switch cmd {
	case Cmd_ContactsResolveUsername:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsResolveUsername))
		x.String(m.Username)

		return x.GetBuf()

	default:
		log.Errorf("ContactsResolveUsername Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsSearch) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsSearch, layer)

	switch cmd {
	case Cmd_ContactsSearch:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsSearch))
		x.String(m.Q)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("ContactsSearch Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsToggleTopPeers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsToggleTopPeers, layer)

	switch cmd {
	case Cmd_ContactsToggleTopPeers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsToggleTopPeers))
		x.Bytes(m.Enabled.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ContactsToggleTopPeers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsUnblock) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContactsUnblock, layer)

	switch cmd {
	case Cmd_ContactsUnblock:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsUnblock))
		x.Bytes(m.IdE54100BD71.Encode(layer))

		return x.GetBuf()
	case Cmd_ContactsUnblockbea65d50:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsUnblockbea65d50))
		x.Bytes(m.IdBEA65D50119.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("ContactsUnblock Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContestSaveDeveloperInfo) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassContestSaveDeveloperInfo, layer)

	switch cmd {
	case Cmd_ContestSaveDeveloperInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContestSaveDeveloperInfo))
		x.Int(m.VkId)
		x.String(m.Name)
		x.String(m.PhoneNumber)
		x.Int(m.Age)
		x.String(m.City)

		return x.GetBuf()

	default:
		log.Errorf("ContestSaveDeveloperInfo Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDestroyAuthKey) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassDestroyAuthKey, layer)

	switch cmd {
	case Cmd_DestroyAuthKey:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DestroyAuthKey))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("DestroyAuthKey Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDestroySession) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassDestroySession, layer)

	switch cmd {
	case Cmd_DestroySession:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DestroySession))
		x.Long(m.SessionId)

		return x.GetBuf()

	default:
		log.Errorf("DestroySession Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFoldersDeleteFolder) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassFoldersDeleteFolder, layer)

	switch cmd {
	case Cmd_FoldersDeleteFolder:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FoldersDeleteFolder))
		x.Int(m.FolderId)

		return x.GetBuf()

	default:
		log.Errorf("FoldersDeleteFolder Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFoldersEditPeerFolders) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassFoldersEditPeerFolders, layer)

	switch cmd {
	case Cmd_FoldersEditPeerFolders:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FoldersEditPeerFolders))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.FolderPeers)))
		for _, v := range m.FolderPeers {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("FoldersEditPeerFolders Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLGetFutureSalts) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassGetFutureSalts, layer)

	switch cmd {
	case Cmd_GetFutureSalts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_GetFutureSalts))
		x.Int(m.Num)

		return x.GetBuf()

	default:
		log.Errorf("GetFutureSalts Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpAcceptTermsOfService) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpAcceptTermsOfService, layer)

	switch cmd {
	case Cmd_HelpAcceptTermsOfService:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpAcceptTermsOfService))
		x.Bytes(m.Id.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("HelpAcceptTermsOfService Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpDismissSuggestion) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpDismissSuggestion, layer)

	switch cmd {
	case Cmd_HelpDismissSuggestion:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpDismissSuggestion))
		x.String(m.Suggestion)

		return x.GetBuf()

	default:
		log.Errorf("HelpDismissSuggestion Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpEditUserInfo) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpEditUserInfo, layer)

	switch cmd {
	case Cmd_HelpEditUserInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpEditUserInfo))
		x.Bytes(m.UserId.Encode(layer))
		x.String(m.Message)
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Entities)))
		for _, v := range m.Entities {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("HelpEditUserInfo Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetAppChangelog) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetAppChangelog, layer)

	switch cmd {
	case Cmd_HelpGetAppChangelog:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetAppChangelog))
		x.String(m.PrevAppVersion)

		return x.GetBuf()
	case Cmd_HelpGetAppChangelog5bab7fb2:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetAppChangelog5bab7fb2))
		x.String(m.DeviceModel)
		x.String(m.SystemVersion)
		x.String(m.AppVersion)
		x.String(m.LangCode)

		return x.GetBuf()

	default:
		log.Errorf("HelpGetAppChangelog Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetAppConfig) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetAppConfig, layer)

	switch cmd {
	case Cmd_HelpGetAppConfig:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetAppConfig))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("HelpGetAppConfig Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetAppUpdate) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetAppUpdate, layer)

	switch cmd {
	case Cmd_HelpGetAppUpdate:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetAppUpdate))
		_ = m

		return x.GetBuf()
	case Cmd_HelpGetAppUpdate522d5a7d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetAppUpdate522d5a7d))
		x.String(m.Source)

		return x.GetBuf()
	case Cmd_HelpGetAppUpdatec812ac7e:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetAppUpdatec812ac7e))
		x.String(m.DeviceModel)
		x.String(m.SystemVersion)
		x.String(m.AppVersion)
		x.String(m.LangCode)

		return x.GetBuf()

	default:
		log.Errorf("HelpGetAppUpdate Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetCdnConfig) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetCdnConfig, layer)

	switch cmd {
	case Cmd_HelpGetCdnConfig:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetCdnConfig))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("HelpGetCdnConfig Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetConfig) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetConfig, layer)

	switch cmd {
	case Cmd_HelpGetConfig:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetConfig))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("HelpGetConfig Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetCountriesList) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetCountriesList, layer)

	switch cmd {
	case Cmd_HelpGetCountriesList:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetCountriesList))
		x.String(m.LangCode)
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("HelpGetCountriesList Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetDeepLinkInfo) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetDeepLinkInfo, layer)

	switch cmd {
	case Cmd_HelpGetDeepLinkInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetDeepLinkInfo))
		x.String(m.Path)

		return x.GetBuf()

	default:
		log.Errorf("HelpGetDeepLinkInfo Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetInviteText) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetInviteText, layer)

	switch cmd {
	case Cmd_HelpGetInviteText:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetInviteText))
		_ = m

		return x.GetBuf()
	case Cmd_HelpGetInviteTexta4a95186:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetInviteTexta4a95186))
		x.String(m.LangCode)

		return x.GetBuf()

	default:
		log.Errorf("HelpGetInviteText Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetNearestDc) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetNearestDc, layer)

	switch cmd {
	case Cmd_HelpGetNearestDc:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetNearestDc))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("HelpGetNearestDc Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetPassportConfig) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetPassportConfig, layer)

	switch cmd {
	case Cmd_HelpGetPassportConfig:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetPassportConfig))
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("HelpGetPassportConfig Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetPromoData) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetPromoData, layer)

	switch cmd {
	case Cmd_HelpGetPromoData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetPromoData))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("HelpGetPromoData Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetProxyData) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetProxyData, layer)

	switch cmd {
	case Cmd_HelpGetProxyData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetProxyData))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("HelpGetProxyData Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetRecentMeUrls) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetRecentMeUrls, layer)

	switch cmd {
	case Cmd_HelpGetRecentMeUrls:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetRecentMeUrls))
		x.String(m.Referer)

		return x.GetBuf()

	default:
		log.Errorf("HelpGetRecentMeUrls Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetSupport) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetSupport, layer)

	switch cmd {
	case Cmd_HelpGetSupport:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetSupport))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("HelpGetSupport Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetSupportName) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetSupportName, layer)

	switch cmd {
	case Cmd_HelpGetSupportName:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetSupportName))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("HelpGetSupportName Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetTermsOfService) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetTermsOfService, layer)

	switch cmd {
	case Cmd_HelpGetTermsOfService:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetTermsOfService))
		_ = m

		return x.GetBuf()
	case Cmd_HelpGetTermsOfService37d78f83:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetTermsOfService37d78f83))
		x.String(m.LangCode)

		return x.GetBuf()

	default:
		log.Errorf("HelpGetTermsOfService Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetTermsOfServiceUpdate) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetTermsOfServiceUpdate, layer)

	switch cmd {
	case Cmd_HelpGetTermsOfServiceUpdate:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetTermsOfServiceUpdate))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("HelpGetTermsOfServiceUpdate Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpGetUserInfo) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpGetUserInfo, layer)

	switch cmd {
	case Cmd_HelpGetUserInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpGetUserInfo))
		x.Bytes(m.UserId.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("HelpGetUserInfo Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpHidePromoData) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpHidePromoData, layer)

	switch cmd {
	case Cmd_HelpHidePromoData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpHidePromoData))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("HelpHidePromoData Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpSaveAppLog) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpSaveAppLog, layer)

	switch cmd {
	case Cmd_HelpSaveAppLog:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpSaveAppLog))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Events)))
		for _, v := range m.Events {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("HelpSaveAppLog Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpSetBotUpdatesStatus) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpSetBotUpdatesStatus, layer)

	switch cmd {
	case Cmd_HelpSetBotUpdatesStatus:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpSetBotUpdatesStatus))
		x.Int(m.PendingUpdatesCount)
		x.String(m.Message)

		return x.GetBuf()

	default:
		log.Errorf("HelpSetBotUpdatesStatus Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpTest) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassHelpTest, layer)

	switch cmd {
	case Cmd_HelpTest:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpTest))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("HelpTest Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLLangpackGetDifference) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassLangpackGetDifference, layer)

	switch cmd {
	case Cmd_LangpackGetDifference:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangpackGetDifference))
		x.Int(m.FromVersion)

		return x.GetBuf()
	case Cmd_LangpackGetDifference9d51e814:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangpackGetDifference9d51e814))
		x.String(m.LangCode)
		x.Int(m.FromVersion)

		return x.GetBuf()
	case Cmd_LangpackGetDifferencecd984aa5:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangpackGetDifferencecd984aa5))
		x.String(m.LangPack)
		x.String(m.LangCode)
		x.Int(m.FromVersion)

		return x.GetBuf()

	default:
		log.Errorf("LangpackGetDifference Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLLangpackGetLangPack) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassLangpackGetLangPack, layer)

	switch cmd {
	case Cmd_LangpackGetLangPack:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangpackGetLangPack))
		x.String(m.LangCode)

		return x.GetBuf()
	case Cmd_LangpackGetLangPackf2f2330a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangpackGetLangPackf2f2330a))
		x.String(m.LangPack)
		x.String(m.LangCode)

		return x.GetBuf()

	default:
		log.Errorf("LangpackGetLangPack Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLLangpackGetLanguage) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassLangpackGetLanguage, layer)

	switch cmd {
	case Cmd_LangpackGetLanguage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangpackGetLanguage))
		x.String(m.LangPack)
		x.String(m.LangCode)

		return x.GetBuf()

	default:
		log.Errorf("LangpackGetLanguage Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLLangpackGetLanguages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassLangpackGetLanguages, layer)

	switch cmd {
	case Cmd_LangpackGetLanguages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangpackGetLanguages))
		_ = m

		return x.GetBuf()
	case Cmd_LangpackGetLanguages42c6978f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangpackGetLanguages42c6978f))
		x.String(m.LangPack)

		return x.GetBuf()

	default:
		log.Errorf("LangpackGetLanguages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLLangpackGetStrings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassLangpackGetStrings, layer)

	switch cmd {
	case Cmd_LangpackGetStrings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangpackGetStrings))
		x.String(m.LangCode)
		x.VectorString(m.Keys)

		return x.GetBuf()
	case Cmd_LangpackGetStringsefea3803:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangpackGetStringsefea3803))
		x.String(m.LangPack)
		x.String(m.LangCode)
		x.VectorString(m.Keys)

		return x.GetBuf()

	default:
		log.Errorf("LangpackGetStrings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesAcceptEncryption) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesAcceptEncryption, layer)

	switch cmd {
	case Cmd_MessagesAcceptEncryption:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesAcceptEncryption))
		x.Bytes(m.Peer.Encode(layer))
		x.StringBytes(m.GB)
		x.Long(m.KeyFingerprint)

		return x.GetBuf()

	default:
		log.Errorf("MessagesAcceptEncryption Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesAcceptUrlAuth) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesAcceptUrlAuth, layer)

	switch cmd {
	case Cmd_MessagesAcceptUrlAuth:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesAcceptUrlAuth))
		// flags
		var flags int32 = 0
		if m.WriteAllowed == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MsgId)
		x.Int(m.ButtonId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesAcceptUrlAuth Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesAddChatUser) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesAddChatUser, layer)

	switch cmd {
	case Cmd_MessagesAddChatUser:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesAddChatUser))
		x.Int(m.ChatId)
		x.Bytes(m.UserId.Encode(layer))
		x.Int(m.FwdLimit)

		return x.GetBuf()

	default:
		log.Errorf("MessagesAddChatUser Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesCheckChatInvite) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesCheckChatInvite, layer)

	switch cmd {
	case Cmd_MessagesCheckChatInvite:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesCheckChatInvite))
		x.String(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesCheckChatInvite Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesClearAllDrafts) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesClearAllDrafts, layer)

	switch cmd {
	case Cmd_MessagesClearAllDrafts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesClearAllDrafts))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("MessagesClearAllDrafts Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesClearRecentStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesClearRecentStickers, layer)

	switch cmd {
	case Cmd_MessagesClearRecentStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesClearRecentStickers))
		// flags
		var flags int32 = 0
		if m.Attached == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("MessagesClearRecentStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesCreateChat) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesCreateChat, layer)

	switch cmd {
	case Cmd_MessagesCreateChat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesCreateChat))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Users)))
		for _, v := range m.Users {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.String(m.Title)

		return x.GetBuf()

	default:
		log.Errorf("MessagesCreateChat Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesDeleteChatUser) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesDeleteChatUser, layer)

	switch cmd {
	case Cmd_MessagesDeleteChatUser:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDeleteChatUser))
		x.Int(m.ChatId)
		x.Bytes(m.UserId.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesDeleteChatUser Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesDeleteHistory) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesDeleteHistory, layer)

	switch cmd {
	case Cmd_MessagesDeleteHistory:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDeleteHistory))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MaxId)

		return x.GetBuf()
	case Cmd_MessagesDeleteHistory1c015b09:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDeleteHistory1c015b09))
		// flags
		var flags int32 = 0
		if m.JustClear == true {
			flags |= 1 << 0
		}
		if m.Revoke == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MaxId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesDeleteHistory Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesDeleteMessages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesDeleteMessages, layer)

	switch cmd {
	case Cmd_MessagesDeleteMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDeleteMessages))
		// flags
		var flags int32 = 0
		if m.Revoke == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.VectorInt(m.Id)

		return x.GetBuf()
	case Cmd_MessagesDeleteMessagesa5f18925:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDeleteMessagesa5f18925))
		x.VectorInt(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("MessagesDeleteMessages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesDeleteScheduledMessages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesDeleteScheduledMessages, layer)

	switch cmd {
	case Cmd_MessagesDeleteScheduledMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDeleteScheduledMessages))
		x.Bytes(m.Peer.Encode(layer))
		x.VectorInt(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("MessagesDeleteScheduledMessages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesDiscardEncryption) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesDiscardEncryption, layer)

	switch cmd {
	case Cmd_MessagesDiscardEncryption:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDiscardEncryption))
		x.Int(m.ChatId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesDiscardEncryption Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesEditChatAbout) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesEditChatAbout, layer)

	switch cmd {
	case Cmd_MessagesEditChatAbout:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditChatAbout))
		x.Bytes(m.Peer.Encode(layer))
		x.String(m.About)

		return x.GetBuf()

	default:
		log.Errorf("MessagesEditChatAbout Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesEditChatAdmin) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesEditChatAdmin, layer)

	switch cmd {
	case Cmd_MessagesEditChatAdmin:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditChatAdmin))
		x.Int(m.ChatId)
		x.Bytes(m.UserId.Encode(layer))
		x.Bytes(m.IsAdmin.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesEditChatAdmin Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesEditChatDefaultBannedRights) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesEditChatDefaultBannedRights, layer)

	switch cmd {
	case Cmd_MessagesEditChatDefaultBannedRights:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditChatDefaultBannedRights))
		x.Bytes(m.Peer.Encode(layer))
		x.Bytes(m.BannedRights.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesEditChatDefaultBannedRights Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesEditChatPhoto) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesEditChatPhoto, layer)

	switch cmd {
	case Cmd_MessagesEditChatPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditChatPhoto))
		x.Int(m.ChatId)
		x.Bytes(m.Photo.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesEditChatPhoto Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesEditChatTitle) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesEditChatTitle, layer)

	switch cmd {
	case Cmd_MessagesEditChatTitle:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditChatTitle))
		x.Int(m.ChatId)
		x.String(m.Title)

		return x.GetBuf()

	default:
		log.Errorf("MessagesEditChatTitle Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesEditInlineBotMessage) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesEditInlineBotMessage, layer)

	switch cmd {
	case Cmd_MessagesEditInlineBotMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditInlineBotMessage))
		// flags
		var flags int32 = 0
		if m.NoWebpage == true {
			flags |= 1 << 1
		}
		if len(m.Message) > 0 {
			flags |= 1 << 11
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Bytes(m.Id.Encode(layer))
		if len(m.Message) > 0 {
			x.String(m.Message)
		}
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_MessagesEditInlineBotMessage83557dba:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditInlineBotMessage83557dba))
		// flags
		var flags int32 = 0
		if m.NoWebpage == true {
			flags |= 1 << 1
		}
		if len(m.Message) > 0 {
			flags |= 1 << 11
		}
		if m.Media != nil {
			flags |= 1 << 14
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Bytes(m.Id.Encode(layer))
		if len(m.Message) > 0 {
			x.String(m.Message)
		}
		if m.Media != nil {
			x.Bytes(m.Media.Encode(layer))
		}
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_MessagesEditInlineBotMessageadc3e828:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditInlineBotMessageadc3e828))
		// flags
		var flags int32 = 0
		if m.NoWebpage == true {
			flags |= 1 << 1
		}
		if m.StopGeoLive == true {
			flags |= 1 << 12
		}
		if len(m.Message) > 0 {
			flags |= 1 << 11
		}
		if m.Media != nil {
			flags |= 1 << 14
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		if m.GeoPoint != nil {
			flags |= 1 << 13
		}
		x.Int(flags)

		x.Bytes(m.Id.Encode(layer))
		if len(m.Message) > 0 {
			x.String(m.Message)
		}
		if m.Media != nil {
			x.Bytes(m.Media.Encode(layer))
		}
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GeoPoint != nil {
			x.Bytes(m.GeoPoint.Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesEditInlineBotMessage Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesEditMessage) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesEditMessage, layer)

	switch cmd {
	case Cmd_MessagesEditMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditMessage))
		// flags
		var flags int32 = 0
		if m.NoWebpage == true {
			flags |= 1 << 1
		}
		if len(m.Message) > 0 {
			flags |= 1 << 11
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Id)
		if len(m.Message) > 0 {
			x.String(m.Message)
		}
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_MessagesEditMessage48f71778:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditMessage48f71778))
		// flags
		var flags int32 = 0
		if m.NoWebpage == true {
			flags |= 1 << 1
		}
		if len(m.Message) > 0 {
			flags |= 1 << 11
		}
		if m.Media != nil {
			flags |= 1 << 14
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		if m.ScheduleDate != 0 {
			flags |= 1 << 15
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Id)
		if len(m.Message) > 0 {
			x.String(m.Message)
		}
		if m.Media != nil {
			x.Bytes(m.Media.Encode(layer))
		}
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.ScheduleDate != 0 {
			x.Int(m.ScheduleDate)
		}

		return x.GetBuf()
	case Cmd_MessagesEditMessagec000e4c8:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditMessagec000e4c8))
		// flags
		var flags int32 = 0
		if m.NoWebpage == true {
			flags |= 1 << 1
		}
		if m.StopGeoLive == true {
			flags |= 1 << 12
		}
		if len(m.Message) > 0 {
			flags |= 1 << 11
		}
		if m.Media != nil {
			flags |= 1 << 14
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		if m.GeoPoint != nil {
			flags |= 1 << 13
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Id)
		if len(m.Message) > 0 {
			x.String(m.Message)
		}
		if m.Media != nil {
			x.Bytes(m.Media.Encode(layer))
		}
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GeoPoint != nil {
			x.Bytes(m.GeoPoint.Encode(layer))
		}

		return x.GetBuf()
	case Cmd_MessagesEditMessaged116f31e:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesEditMessaged116f31e))
		// flags
		var flags int32 = 0
		if m.NoWebpage == true {
			flags |= 1 << 1
		}
		if len(m.Message) > 0 {
			flags |= 1 << 11
		}
		if m.Media != nil {
			flags |= 1 << 14
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Id)
		if len(m.Message) > 0 {
			x.String(m.Message)
		}
		if m.Media != nil {
			x.Bytes(m.Media.Encode(layer))
		}
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesEditMessage Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesExportChatInvite) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesExportChatInvite, layer)

	switch cmd {
	case Cmd_MessagesExportChatInvite:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesExportChatInvite))
		x.Int(m.ChatId)

		return x.GetBuf()
	case Cmd_MessagesExportChatInvitedf7534c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesExportChatInvitedf7534c))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesExportChatInvite Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesFaveSticker) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesFaveSticker, layer)

	switch cmd {
	case Cmd_MessagesFaveSticker:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesFaveSticker))
		x.Bytes(m.Id.Encode(layer))
		x.Bytes(m.Unfave.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesFaveSticker Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesForwardMessage) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesForwardMessage, layer)

	switch cmd {
	case Cmd_MessagesForwardMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesForwardMessage))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Id)
		x.Long(m.RandomId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesForwardMessage Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesForwardMessages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesForwardMessages, layer)

	switch cmd {
	case Cmd_MessagesForwardMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesForwardMessages))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 5
		}
		if m.Background == true {
			flags |= 1 << 6
		}
		if m.WithMyScore == true {
			flags |= 1 << 8
		}
		x.Int(flags)

		x.Bytes(m.FromPeer.Encode(layer))
		x.VectorInt(m.Id)

		x.VectorLong(m.RandomId)

		x.Bytes(m.ToPeer.Encode(layer))

		return x.GetBuf()
	case Cmd_MessagesForwardMessagesd9fee60e:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesForwardMessagesd9fee60e))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 5
		}
		if m.Background == true {
			flags |= 1 << 6
		}
		if m.WithMyScore == true {
			flags |= 1 << 8
		}
		if m.ScheduleDate != 0 {
			flags |= 1 << 10
		}
		x.Int(flags)

		x.Bytes(m.FromPeer.Encode(layer))
		x.VectorInt(m.Id)

		x.VectorLong(m.RandomId)

		x.Bytes(m.ToPeer.Encode(layer))
		if m.ScheduleDate != 0 {
			x.Int(m.ScheduleDate)
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesForwardMessages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetAllChats) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetAllChats, layer)

	switch cmd {
	case Cmd_MessagesGetAllChats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetAllChats))
		x.VectorInt(m.ExceptIds)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetAllChats Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetAllDrafts) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetAllDrafts, layer)

	switch cmd {
	case Cmd_MessagesGetAllDrafts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetAllDrafts))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetAllDrafts Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetAllStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetAllStickers, layer)

	switch cmd {
	case Cmd_MessagesGetAllStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetAllStickers))
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetAllStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetArchivedStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetArchivedStickers, layer)

	switch cmd {
	case Cmd_MessagesGetArchivedStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetArchivedStickers))
		// flags
		var flags int32 = 0
		if m.Masks == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.OffsetId)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetArchivedStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetAttachedStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetAttachedStickers, layer)

	switch cmd {
	case Cmd_MessagesGetAttachedStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetAttachedStickers))
		x.Bytes(m.Media.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetAttachedStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetBotCallbackAnswer) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetBotCallbackAnswer, layer)

	switch cmd {
	case Cmd_MessagesGetBotCallbackAnswer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetBotCallbackAnswer))
		// flags
		var flags int32 = 0
		if m.Game == true {
			flags |= 1 << 1
		}
		if len(m.Data) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MsgId)
		if len(m.Data) > 0 {
			x.StringBytes(m.Data)
		}

		return x.GetBuf()
	case Cmd_MessagesGetBotCallbackAnswer9342ca07:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetBotCallbackAnswer9342ca07))
		// flags
		var flags int32 = 0
		if m.Game == true {
			flags |= 1 << 1
		}
		if len(m.Data) > 0 {
			flags |= 1 << 0
		}
		if m.Password != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MsgId)
		if len(m.Data) > 0 {
			x.StringBytes(m.Data)
		}
		if m.Password != nil {
			x.Bytes(m.Password.Encode(layer))
		}

		return x.GetBuf()
	case Cmd_MessagesGetBotCallbackAnswera6e94f04:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetBotCallbackAnswera6e94f04))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MsgId)
		x.StringBytes(m.Data)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetBotCallbackAnswer Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetChats) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetChats, layer)

	switch cmd {
	case Cmd_MessagesGetChats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetChats))
		x.VectorInt(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetChats Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetCommonChats) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetCommonChats, layer)

	switch cmd {
	case Cmd_MessagesGetCommonChats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetCommonChats))
		x.Bytes(m.UserId.Encode(layer))
		x.Int(m.MaxId)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetCommonChats Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetDhConfig) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetDhConfig, layer)

	switch cmd {
	case Cmd_MessagesGetDhConfig:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetDhConfig))
		x.Int(m.Version)
		x.Int(m.RandomLength)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetDhConfig Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetDialogFilters) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetDialogFilters, layer)

	switch cmd {
	case Cmd_MessagesGetDialogFilters:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetDialogFilters))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetDialogFilters Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetDialogUnreadMarks) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetDialogUnreadMarks, layer)

	switch cmd {
	case Cmd_MessagesGetDialogUnreadMarks:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetDialogUnreadMarks))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetDialogUnreadMarks Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetDialogs) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetDialogs, layer)

	switch cmd {
	case Cmd_MessagesGetDialogs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetDialogs))
		// flags
		var flags int32 = 0
		if m.ExcludePinned == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.OffsetDate)
		x.Int(m.OffsetId)
		x.Bytes(m.OffsetPeer.Encode(layer))
		x.Int(m.Limit)

		return x.GetBuf()
	case Cmd_MessagesGetDialogs6b47f94d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetDialogs6b47f94d))
		x.Int(m.OffsetDate)
		x.Int(m.OffsetId)
		x.Bytes(m.OffsetPeer.Encode(layer))
		x.Int(m.Limit)

		return x.GetBuf()
	case Cmd_MessagesGetDialogsa0ee3b73:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetDialogsa0ee3b73))
		// flags
		var flags int32 = 0
		if m.ExcludePinned == true {
			flags |= 1 << 0
		}
		if m.FolderId != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		if m.FolderId != 0 {
			x.Int(m.FolderId)
		}
		x.Int(m.OffsetDate)
		x.Int(m.OffsetId)
		x.Bytes(m.OffsetPeer.Encode(layer))
		x.Int(m.Limit)
		x.Int(m.Hash)

		return x.GetBuf()
	case Cmd_MessagesGetDialogsb098aee6:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetDialogsb098aee6))
		// flags
		var flags int32 = 0
		if m.ExcludePinned == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.OffsetDate)
		x.Int(m.OffsetId)
		x.Bytes(m.OffsetPeer.Encode(layer))
		x.Int(m.Limit)
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetDialogs Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetDiscussionMessage) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetDiscussionMessage, layer)

	switch cmd {
	case Cmd_MessagesGetDiscussionMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetDiscussionMessage))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MsgId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetDiscussionMessage Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetDocumentByHash) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetDocumentByHash, layer)

	switch cmd {
	case Cmd_MessagesGetDocumentByHash:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetDocumentByHash))
		x.StringBytes(m.Sha256)
		x.Int(m.Size_)
		x.String(m.MimeType)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetDocumentByHash Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetEmojiKeywords) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetEmojiKeywords, layer)

	switch cmd {
	case Cmd_MessagesGetEmojiKeywords:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetEmojiKeywords))
		x.String(m.LangCode)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetEmojiKeywords Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetEmojiKeywordsDifference) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetEmojiKeywordsDifference, layer)

	switch cmd {
	case Cmd_MessagesGetEmojiKeywordsDifference:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetEmojiKeywordsDifference))
		x.String(m.LangCode)
		x.Int(m.FromVersion)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetEmojiKeywordsDifference Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetEmojiKeywordsLanguages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetEmojiKeywordsLanguages, layer)

	switch cmd {
	case Cmd_MessagesGetEmojiKeywordsLanguages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetEmojiKeywordsLanguages))
		x.VectorString(m.LangCodes)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetEmojiKeywordsLanguages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetEmojiURL) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetEmojiURL, layer)

	switch cmd {
	case Cmd_MessagesGetEmojiURL:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetEmojiURL))
		x.String(m.LangCode)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetEmojiURL Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetFavedStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetFavedStickers, layer)

	switch cmd {
	case Cmd_MessagesGetFavedStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetFavedStickers))
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetFavedStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetFeaturedStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetFeaturedStickers, layer)

	switch cmd {
	case Cmd_MessagesGetFeaturedStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetFeaturedStickers))
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetFeaturedStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetFullChat) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetFullChat, layer)

	switch cmd {
	case Cmd_MessagesGetFullChat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetFullChat))
		x.Int(m.ChatId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetFullChat Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetGameHighScores) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetGameHighScores, layer)

	switch cmd {
	case Cmd_MessagesGetGameHighScores:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetGameHighScores))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Id)
		x.Bytes(m.UserId.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetGameHighScores Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetHistory) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetHistory, layer)

	switch cmd {
	case Cmd_MessagesGetHistory:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetHistory))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.OffsetId)
		x.Int(m.OffsetDate)
		x.Int(m.AddOffset)
		x.Int(m.Limit)
		x.Int(m.MaxId)
		x.Int(m.MinId)

		return x.GetBuf()
	case Cmd_MessagesGetHistorydcbb8260:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetHistorydcbb8260))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.OffsetId)
		x.Int(m.OffsetDate)
		x.Int(m.AddOffset)
		x.Int(m.Limit)
		x.Int(m.MaxId)
		x.Int(m.MinId)
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetHistory Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetInlineBotResults) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetInlineBotResults, layer)

	switch cmd {
	case Cmd_MessagesGetInlineBotResults:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetInlineBotResults))
		// flags
		var flags int32 = 0
		if m.GeoPoint != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Bot.Encode(layer))
		x.Bytes(m.Peer.Encode(layer))
		if m.GeoPoint != nil {
			x.Bytes(m.GeoPoint.Encode(layer))
		}
		x.String(m.Query)
		x.String(m.Offset)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetInlineBotResults Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetInlineGameHighScores) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetInlineGameHighScores, layer)

	switch cmd {
	case Cmd_MessagesGetInlineGameHighScores:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetInlineGameHighScores))
		x.Bytes(m.Id.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetInlineGameHighScores Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetMaskStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetMaskStickers, layer)

	switch cmd {
	case Cmd_MessagesGetMaskStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetMaskStickers))
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetMaskStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetMessageEditData) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetMessageEditData, layer)

	switch cmd {
	case Cmd_MessagesGetMessageEditData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetMessageEditData))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetMessageEditData Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetMessages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetMessages, layer)

	switch cmd {
	case Cmd_MessagesGetMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetMessages))
		x.VectorInt(m.Id4222FA7471)

		return x.GetBuf()
	case Cmd_MessagesGetMessages63c66506:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetMessages63c66506))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Id63C6650685)))
		for _, v := range m.Id63C6650685 {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetMessages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetMessagesViews) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetMessagesViews, layer)

	switch cmd {
	case Cmd_MessagesGetMessagesViews:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetMessagesViews))
		x.Bytes(m.Peer.Encode(layer))
		x.VectorInt(m.Id)

		x.Bytes(m.Increment.Encode(layer))

		return x.GetBuf()
	case Cmd_MessagesGetMessagesViews5784d3e1:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetMessagesViews5784d3e1))
		x.Bytes(m.Peer.Encode(layer))
		x.VectorInt(m.Id)

		x.Bytes(m.Increment.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetMessagesViews Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetOldFeaturedStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetOldFeaturedStickers, layer)

	switch cmd {
	case Cmd_MessagesGetOldFeaturedStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetOldFeaturedStickers))
		x.Int(m.Offset)
		x.Int(m.Limit)
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetOldFeaturedStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetOnlines) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetOnlines, layer)

	switch cmd {
	case Cmd_MessagesGetOnlines:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetOnlines))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetOnlines Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetPeerDialogs) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetPeerDialogs, layer)

	switch cmd {
	case Cmd_MessagesGetPeerDialogs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetPeerDialogs))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Peers2D9776B971)))
		for _, v := range m.Peers2D9776B971 {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesGetPeerDialogse470bcfd:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetPeerDialogse470bcfd))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.PeersE470BCFD85)))
		for _, v := range m.PeersE470BCFD85 {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetPeerDialogs Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetPeerSettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetPeerSettings, layer)

	switch cmd {
	case Cmd_MessagesGetPeerSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetPeerSettings))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetPeerSettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetPinnedDialogs) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetPinnedDialogs, layer)

	switch cmd {
	case Cmd_MessagesGetPinnedDialogs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetPinnedDialogs))
		_ = m

		return x.GetBuf()
	case Cmd_MessagesGetPinnedDialogsd6b94df2:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetPinnedDialogsd6b94df2))
		x.Int(m.FolderId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetPinnedDialogs Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetPollResults) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetPollResults, layer)

	switch cmd {
	case Cmd_MessagesGetPollResults:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetPollResults))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MsgId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetPollResults Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetPollVotes) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetPollVotes, layer)

	switch cmd {
	case Cmd_MessagesGetPollVotes:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetPollVotes))
		// flags
		var flags int32 = 0
		if len(m.Option) > 0 {
			flags |= 1 << 0
		}
		if len(m.Offset) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Id)
		if len(m.Option) > 0 {
			x.StringBytes(m.Option)
		}
		if len(m.Offset) > 0 {
			x.String(m.Offset)
		}
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetPollVotes Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetRecentLocations) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetRecentLocations, layer)

	switch cmd {
	case Cmd_MessagesGetRecentLocations:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetRecentLocations))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Limit)

		return x.GetBuf()
	case Cmd_MessagesGetRecentLocationsbbc45b09:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetRecentLocationsbbc45b09))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Limit)
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetRecentLocations Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetRecentStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetRecentStickers, layer)

	switch cmd {
	case Cmd_MessagesGetRecentStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetRecentStickers))
		// flags
		var flags int32 = 0
		if m.Attached == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetRecentStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetReplies) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetReplies, layer)

	switch cmd {
	case Cmd_MessagesGetReplies:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetReplies))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MsgId)
		x.Int(m.OffsetId)
		x.Int(m.OffsetDate)
		x.Int(m.AddOffset)
		x.Int(m.Limit)
		x.Int(m.MaxId)
		x.Int(m.MinId)
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetReplies Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetSavedGifs) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetSavedGifs, layer)

	switch cmd {
	case Cmd_MessagesGetSavedGifs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetSavedGifs))
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetSavedGifs Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetScheduledHistory) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetScheduledHistory, layer)

	switch cmd {
	case Cmd_MessagesGetScheduledHistory:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetScheduledHistory))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetScheduledHistory Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetScheduledMessages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetScheduledMessages, layer)

	switch cmd {
	case Cmd_MessagesGetScheduledMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetScheduledMessages))
		x.Bytes(m.Peer.Encode(layer))
		x.VectorInt(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetScheduledMessages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetSearchCounters) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetSearchCounters, layer)

	switch cmd {
	case Cmd_MessagesGetSearchCounters:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetSearchCounters))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Filters)))
		for _, v := range m.Filters {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetSearchCounters Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetSplitRanges) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetSplitRanges, layer)

	switch cmd {
	case Cmd_MessagesGetSplitRanges:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetSplitRanges))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetSplitRanges Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetStatsURL) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetStatsURL, layer)

	switch cmd {
	case Cmd_MessagesGetStatsURL:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetStatsURL))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()
	case Cmd_MessagesGetStatsURL812c2ae6:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetStatsURL812c2ae6))
		// flags
		var flags int32 = 0
		if m.Dark == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.String(m.Params)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetStatsURL Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetStickerSet) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetStickerSet, layer)

	switch cmd {
	case Cmd_MessagesGetStickerSet:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetStickerSet))
		x.Bytes(m.Stickerset.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetStickerSet Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetStickers, layer)

	switch cmd {
	case Cmd_MessagesGetStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetStickers))
		x.String(m.Emoticon)
		x.String(m.HashAE22E04551)

		return x.GetBuf()
	case Cmd_MessagesGetStickers43d4f2c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetStickers43d4f2c))
		x.String(m.Emoticon)
		x.Int(m.Hash43D4F2C85)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetSuggestedDialogFilters) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetSuggestedDialogFilters, layer)

	switch cmd {
	case Cmd_MessagesGetSuggestedDialogFilters:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetSuggestedDialogFilters))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetSuggestedDialogFilters Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetUnreadMentions) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetUnreadMentions, layer)

	switch cmd {
	case Cmd_MessagesGetUnreadMentions:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetUnreadMentions))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.OffsetId)
		x.Int(m.AddOffset)
		x.Int(m.Limit)
		x.Int(m.MaxId)
		x.Int(m.MinId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetUnreadMentions Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetWebPage) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetWebPage, layer)

	switch cmd {
	case Cmd_MessagesGetWebPage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetWebPage))
		x.String(m.Url)
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetWebPage Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesGetWebPagePreview) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesGetWebPagePreview, layer)

	switch cmd {
	case Cmd_MessagesGetWebPagePreview:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetWebPagePreview))
		x.String(m.Message)

		return x.GetBuf()
	case Cmd_MessagesGetWebPagePreview8b68b0cc:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesGetWebPagePreview8b68b0cc))
		// flags
		var flags int32 = 0
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.String(m.Message)
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesGetWebPagePreview Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesHidePeerSettingsBar) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesHidePeerSettingsBar, layer)

	switch cmd {
	case Cmd_MessagesHidePeerSettingsBar:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesHidePeerSettingsBar))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesHidePeerSettingsBar Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesHideReportSpam) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesHideReportSpam, layer)

	switch cmd {
	case Cmd_MessagesHideReportSpam:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesHideReportSpam))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesHideReportSpam Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesImportChatInvite) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesImportChatInvite, layer)

	switch cmd {
	case Cmd_MessagesImportChatInvite:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesImportChatInvite))
		x.String(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesImportChatInvite Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesInstallStickerSet) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesInstallStickerSet, layer)

	switch cmd {
	case Cmd_MessagesInstallStickerSet:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesInstallStickerSet))
		x.Bytes(m.Stickerset.Encode(layer))
		x.Bytes(m.Archived.Encode(layer))

		return x.GetBuf()
	case Cmd_MessagesInstallStickerSet7b30c3a6:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesInstallStickerSet7b30c3a6))
		x.Bytes(m.Stickerset.Encode(layer))
		x.Bytes(m.Disabled.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesInstallStickerSet Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesMarkDialogUnread) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesMarkDialogUnread, layer)

	switch cmd {
	case Cmd_MessagesMarkDialogUnread:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesMarkDialogUnread))
		// flags
		var flags int32 = 0
		if m.Unread == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesMarkDialogUnread Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesMigrateChat) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesMigrateChat, layer)

	switch cmd {
	case Cmd_MessagesMigrateChat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesMigrateChat))
		x.Int(m.ChatId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesMigrateChat Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReadDiscussion) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReadDiscussion, layer)

	switch cmd {
	case Cmd_MessagesReadDiscussion:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReadDiscussion))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MsgId)
		x.Int(m.ReadMaxId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesReadDiscussion Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReadEncryptedHistory) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReadEncryptedHistory, layer)

	switch cmd {
	case Cmd_MessagesReadEncryptedHistory:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReadEncryptedHistory))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MaxDate)

		return x.GetBuf()

	default:
		log.Errorf("MessagesReadEncryptedHistory Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReadFeaturedStickers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReadFeaturedStickers, layer)

	switch cmd {
	case Cmd_MessagesReadFeaturedStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReadFeaturedStickers))
		x.VectorLong(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("MessagesReadFeaturedStickers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReadHistory) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReadHistory, layer)

	switch cmd {
	case Cmd_MessagesReadHistory:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReadHistory))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MaxId)

		return x.GetBuf()
	case Cmd_MessagesReadHistoryb04f2510:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReadHistoryb04f2510))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MaxId)
		x.Int(m.Offset)

		return x.GetBuf()

	default:
		log.Errorf("MessagesReadHistory Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReadMentions) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReadMentions, layer)

	switch cmd {
	case Cmd_MessagesReadMentions:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReadMentions))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesReadMentions Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReadMessageContents) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReadMessageContents, layer)

	switch cmd {
	case Cmd_MessagesReadMessageContents:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReadMessageContents))
		x.VectorInt(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("MessagesReadMessageContents Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReceivedMessages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReceivedMessages, layer)

	switch cmd {
	case Cmd_MessagesReceivedMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReceivedMessages))
		x.Int(m.MaxId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesReceivedMessages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReceivedQueue) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReceivedQueue, layer)

	switch cmd {
	case Cmd_MessagesReceivedQueue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReceivedQueue))
		x.Int(m.MaxQts)

		return x.GetBuf()

	default:
		log.Errorf("MessagesReceivedQueue Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReorderPinnedDialogs) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReorderPinnedDialogs, layer)

	switch cmd {
	case Cmd_MessagesReorderPinnedDialogs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReorderPinnedDialogs))
		// flags
		var flags int32 = 0
		if m.Force == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Order959FF64471)))
		for _, v := range m.Order959FF64471 {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesReorderPinnedDialogs3b1adf37:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReorderPinnedDialogs3b1adf37))
		// flags
		var flags int32 = 0
		if m.Force == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.FolderId)
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Order5B51D63F85)))
		for _, v := range m.Order5B51D63F85 {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesReorderPinnedDialogs5b51d63f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReorderPinnedDialogs5b51d63f))
		// flags
		var flags int32 = 0
		if m.Force == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Order5B51D63F85)))
		for _, v := range m.Order5B51D63F85 {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesReorderPinnedDialogs Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReorderStickerSets) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReorderStickerSets, layer)

	switch cmd {
	case Cmd_MessagesReorderStickerSets:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReorderStickerSets))
		// flags
		var flags int32 = 0
		if m.Masks == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.VectorLong(m.Order)

		return x.GetBuf()
	case Cmd_MessagesReorderStickerSets9fcfbc30:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReorderStickerSets9fcfbc30))
		x.VectorLong(m.Order)

		return x.GetBuf()

	default:
		log.Errorf("MessagesReorderStickerSets Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReport) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReport, layer)

	switch cmd {
	case Cmd_MessagesReport:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReport))
		x.Bytes(m.Peer.Encode(layer))
		x.VectorInt(m.Id)

		x.Bytes(m.Reason.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesReport Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReportEncryptedSpam) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReportEncryptedSpam, layer)

	switch cmd {
	case Cmd_MessagesReportEncryptedSpam:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReportEncryptedSpam))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesReportEncryptedSpam Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesReportSpam) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesReportSpam, layer)

	switch cmd {
	case Cmd_MessagesReportSpam:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesReportSpam))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesReportSpam Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesRequestEncryption) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesRequestEncryption, layer)

	switch cmd {
	case Cmd_MessagesRequestEncryption:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesRequestEncryption))
		x.Bytes(m.UserId.Encode(layer))
		x.Int(m.RandomId)
		x.StringBytes(m.GA)

		return x.GetBuf()

	default:
		log.Errorf("MessagesRequestEncryption Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesRequestUrlAuth) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesRequestUrlAuth, layer)

	switch cmd {
	case Cmd_MessagesRequestUrlAuth:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesRequestUrlAuth))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MsgId)
		x.Int(m.ButtonId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesRequestUrlAuth Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSaveDraft) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSaveDraft, layer)

	switch cmd {
	case Cmd_MessagesSaveDraft:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSaveDraft))
		// flags
		var flags int32 = 0
		if m.NoWebpage == true {
			flags |= 1 << 1
		}
		if m.ReplyToMsgId != 0 {
			flags |= 1 << 0
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		if m.ReplyToMsgId != 0 {
			x.Int(m.ReplyToMsgId)
		}
		x.Bytes(m.Peer.Encode(layer))
		x.String(m.Message)
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesSaveDraft Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSaveGif) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSaveGif, layer)

	switch cmd {
	case Cmd_MessagesSaveGif:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSaveGif))
		x.Bytes(m.Id.Encode(layer))
		x.Bytes(m.Unsave.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesSaveGif Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSaveRecentSticker) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSaveRecentSticker, layer)

	switch cmd {
	case Cmd_MessagesSaveRecentSticker:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSaveRecentSticker))
		// flags
		var flags int32 = 0
		if m.Attached == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Id.Encode(layer))
		x.Bytes(m.Unsave.Encode(layer))

		return x.GetBuf()
	case Cmd_MessagesSaveRecentSticker348e39bf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSaveRecentSticker348e39bf))
		x.Bytes(m.Id.Encode(layer))
		x.Bytes(m.Unsave.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesSaveRecentSticker Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSearch) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSearch, layer)

	switch cmd {
	case Cmd_MessagesSearch:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearch))
		// flags
		var flags int32 = 0
		if m.FromId39E9EA071 != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.String(m.Q)
		if m.FromId39E9EA071 != nil {
			x.Bytes(m.FromId39E9EA071.Encode(layer))
		}
		x.Bytes(m.Filter.Encode(layer))
		x.Int(m.MinDate)
		x.Int(m.MaxDate)
		x.Int(m.OffsetId)
		x.Int(m.AddOffset)
		x.Int(m.Limit)
		x.Int(m.MaxId)
		x.Int(m.MinId)

		return x.GetBuf()
	case Cmd_MessagesSearch4e17810b:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearch4e17810b))
		// flags
		var flags int32 = 0
		if m.FromId39E9EA071 != nil {
			flags |= 1 << 0
		}
		if m.TopMsgId != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.String(m.Q)
		if m.FromId39E9EA071 != nil {
			x.Bytes(m.FromId39E9EA071.Encode(layer))
		}
		if m.TopMsgId != 0 {
			x.Int(m.TopMsgId)
		}
		x.Bytes(m.Filter.Encode(layer))
		x.Int(m.MinDate)
		x.Int(m.MaxDate)
		x.Int(m.OffsetId)
		x.Int(m.AddOffset)
		x.Int(m.Limit)
		x.Int(m.MaxId)
		x.Int(m.MinId)
		x.Int(m.Hash)

		return x.GetBuf()
	case Cmd_MessagesSearch8614ef68:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearch8614ef68))
		// flags
		var flags int32 = 0
		if m.FromId39E9EA071 != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.String(m.Q)
		if m.FromId39E9EA071 != nil {
			x.Bytes(m.FromId39E9EA071.Encode(layer))
		}
		x.Bytes(m.Filter.Encode(layer))
		x.Int(m.MinDate)
		x.Int(m.MaxDate)
		x.Int(m.OffsetId)
		x.Int(m.AddOffset)
		x.Int(m.Limit)
		x.Int(m.MaxId)
		x.Int(m.MinId)
		x.Int(m.Hash)

		return x.GetBuf()
	case Cmd_MessagesSearchc352eec:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearchc352eec))
		// flags
		var flags int32 = 0
		if m.FromIdC352EEC120 != nil {
			flags |= 1 << 0
		}
		if m.TopMsgId != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.String(m.Q)
		if m.FromIdC352EEC120 != nil {
			x.Bytes(m.FromIdC352EEC120.Encode(layer))
		}
		if m.TopMsgId != 0 {
			x.Int(m.TopMsgId)
		}
		x.Bytes(m.Filter.Encode(layer))
		x.Int(m.MinDate)
		x.Int(m.MaxDate)
		x.Int(m.OffsetId)
		x.Int(m.AddOffset)
		x.Int(m.Limit)
		x.Int(m.MaxId)
		x.Int(m.MinId)
		x.Int(m.Hash)

		return x.GetBuf()
	case Cmd_MessagesSearchd4569248:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearchd4569248))
		// flags
		var flags int32 = 0
		if m.ImportantOnly == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.String(m.Q)
		x.Bytes(m.Filter.Encode(layer))
		x.Int(m.MinDate)
		x.Int(m.MaxDate)
		x.Int(m.Offset)
		x.Int(m.MaxId)
		x.Int(m.Limit)

		return x.GetBuf()
	case Cmd_MessagesSearchf288a275:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearchf288a275))
		// flags
		var flags int32 = 0
		if m.FromId39E9EA071 != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.String(m.Q)
		if m.FromId39E9EA071 != nil {
			x.Bytes(m.FromId39E9EA071.Encode(layer))
		}
		x.Bytes(m.Filter.Encode(layer))
		x.Int(m.MinDate)
		x.Int(m.MaxDate)
		x.Int(m.Offset)
		x.Int(m.MaxId)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("MessagesSearch Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSearchGifs) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSearchGifs, layer)

	switch cmd {
	case Cmd_MessagesSearchGifs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearchGifs))
		x.String(m.Q)
		x.Int(m.Offset)

		return x.GetBuf()

	default:
		log.Errorf("MessagesSearchGifs Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSearchGlobal) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSearchGlobal, layer)

	switch cmd {
	case Cmd_MessagesSearchGlobal:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearchGlobal))
		x.String(m.Q)
		x.Int(m.OffsetDate)
		x.Bytes(m.OffsetPeer.Encode(layer))
		x.Int(m.OffsetId)
		x.Int(m.Limit)

		return x.GetBuf()
	case Cmd_MessagesSearchGlobal4bc6589a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearchGlobal4bc6589a))
		// flags
		var flags int32 = 0
		if m.FolderId != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.FolderId != 0 {
			x.Int(m.FolderId)
		}
		x.String(m.Q)
		x.Bytes(m.Filter.Encode(layer))
		x.Int(m.MinDate)
		x.Int(m.MaxDate)
		x.Int(m.OffsetRate)
		x.Bytes(m.OffsetPeer.Encode(layer))
		x.Int(m.OffsetId)
		x.Int(m.Limit)

		return x.GetBuf()
	case Cmd_MessagesSearchGlobalbf7225a4:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearchGlobalbf7225a4))
		// flags
		var flags int32 = 0
		if m.FolderId != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.FolderId != 0 {
			x.Int(m.FolderId)
		}
		x.String(m.Q)
		x.Int(m.OffsetRate)
		x.Bytes(m.OffsetPeer.Encode(layer))
		x.Int(m.OffsetId)
		x.Int(m.Limit)

		return x.GetBuf()
	case Cmd_MessagesSearchGlobalf79c611:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearchGlobalf79c611))
		x.String(m.Q)
		x.Int(m.OffsetRate)
		x.Bytes(m.OffsetPeer.Encode(layer))
		x.Int(m.OffsetId)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("MessagesSearchGlobal Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSearchStickerSets) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSearchStickerSets, layer)

	switch cmd {
	case Cmd_MessagesSearchStickerSets:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearchStickerSets))
		// flags
		var flags int32 = 0
		if m.ExcludeFeatured == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.Q)
		x.Int(m.Hash)

		return x.GetBuf()

	default:
		log.Errorf("MessagesSearchStickerSets Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSendBroadcast) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSendBroadcast, layer)

	switch cmd {
	case Cmd_MessagesSendBroadcast:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendBroadcast))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Contacts)))
		for _, v := range m.Contacts {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.VectorLong(m.RandomId)

		x.String(m.Message)
		x.Bytes(m.Media.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesSendBroadcast Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSendEncrypted) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSendEncrypted, layer)

	switch cmd {
	case Cmd_MessagesSendEncrypted:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendEncrypted))
		x.Bytes(m.Peer.Encode(layer))
		x.Long(m.RandomId)
		x.StringBytes(m.Data)

		return x.GetBuf()
	case Cmd_MessagesSendEncrypted44fa7a15:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendEncrypted44fa7a15))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Long(m.RandomId)
		x.StringBytes(m.Data)

		return x.GetBuf()

	default:
		log.Errorf("MessagesSendEncrypted Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSendEncryptedFile) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSendEncryptedFile, layer)

	switch cmd {
	case Cmd_MessagesSendEncryptedFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendEncryptedFile))
		x.Bytes(m.Peer.Encode(layer))
		x.Long(m.RandomId)
		x.StringBytes(m.Data)
		x.Bytes(m.File.Encode(layer))

		return x.GetBuf()
	case Cmd_MessagesSendEncryptedFile5559481d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendEncryptedFile5559481d))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Long(m.RandomId)
		x.StringBytes(m.Data)
		x.Bytes(m.File.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesSendEncryptedFile Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSendEncryptedService) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSendEncryptedService, layer)

	switch cmd {
	case Cmd_MessagesSendEncryptedService:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendEncryptedService))
		x.Bytes(m.Peer.Encode(layer))
		x.Long(m.RandomId)
		x.StringBytes(m.Data)

		return x.GetBuf()

	default:
		log.Errorf("MessagesSendEncryptedService Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSendInlineBotResult) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSendInlineBotResult, layer)

	switch cmd {
	case Cmd_MessagesSendInlineBotResult:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendInlineBotResult))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 5
		}
		if m.Background == true {
			flags |= 1 << 6
		}
		if m.ClearDraft == true {
			flags |= 1 << 7
		}
		if m.ReplyToMsgId != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		if m.ReplyToMsgId != 0 {
			x.Int(m.ReplyToMsgId)
		}
		x.Long(m.RandomId)
		x.Long(m.QueryId)
		x.String(m.Id)

		return x.GetBuf()
	case Cmd_MessagesSendInlineBotResult220815b0:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendInlineBotResult220815b0))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 5
		}
		if m.Background == true {
			flags |= 1 << 6
		}
		if m.ClearDraft == true {
			flags |= 1 << 7
		}
		if m.HideVia == true {
			flags |= 1 << 11
		}
		if m.ReplyToMsgId != 0 {
			flags |= 1 << 0
		}
		if m.ScheduleDate != 0 {
			flags |= 1 << 10
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		if m.ReplyToMsgId != 0 {
			x.Int(m.ReplyToMsgId)
		}
		x.Long(m.RandomId)
		x.Long(m.QueryId)
		x.String(m.Id)
		if m.ScheduleDate != 0 {
			x.Int(m.ScheduleDate)
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesSendInlineBotResult Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSendMedia) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSendMedia, layer)

	switch cmd {
	case Cmd_MessagesSendMedia:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendMedia))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 5
		}
		if m.Background == true {
			flags |= 1 << 6
		}
		if m.ClearDraft == true {
			flags |= 1 << 7
		}
		if m.ReplyToMsgId != 0 {
			flags |= 1 << 0
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		if m.ReplyToMsgId != 0 {
			x.Int(m.ReplyToMsgId)
		}
		x.Bytes(m.Media.Encode(layer))
		x.Long(m.RandomId)
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}

		return x.GetBuf()
	case Cmd_MessagesSendMedia3491eba9:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendMedia3491eba9))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 5
		}
		if m.Background == true {
			flags |= 1 << 6
		}
		if m.ClearDraft == true {
			flags |= 1 << 7
		}
		if m.ReplyToMsgId != 0 {
			flags |= 1 << 0
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		if m.ScheduleDate != 0 {
			flags |= 1 << 10
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		if m.ReplyToMsgId != 0 {
			x.Int(m.ReplyToMsgId)
		}
		x.Bytes(m.Media.Encode(layer))
		x.String(m.Message)
		x.Long(m.RandomId)
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.ScheduleDate != 0 {
			x.Int(m.ScheduleDate)
		}

		return x.GetBuf()
	case Cmd_MessagesSendMediab8d1262b:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendMediab8d1262b))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 5
		}
		if m.Background == true {
			flags |= 1 << 6
		}
		if m.ClearDraft == true {
			flags |= 1 << 7
		}
		if m.ReplyToMsgId != 0 {
			flags |= 1 << 0
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		if m.ReplyToMsgId != 0 {
			x.Int(m.ReplyToMsgId)
		}
		x.Bytes(m.Media.Encode(layer))
		x.String(m.Message)
		x.Long(m.RandomId)
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesSendMedia Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSendMessage) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSendMessage, layer)

	switch cmd {
	case Cmd_MessagesSendMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendMessage))
		// flags
		var flags int32 = 0
		if m.NoWebpage == true {
			flags |= 1 << 1
		}
		if m.Silent == true {
			flags |= 1 << 5
		}
		if m.Background == true {
			flags |= 1 << 6
		}
		if m.ClearDraft == true {
			flags |= 1 << 7
		}
		if m.ReplyToMsgId != 0 {
			flags |= 1 << 0
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		if m.ReplyToMsgId != 0 {
			x.Int(m.ReplyToMsgId)
		}
		x.String(m.Message)
		x.Long(m.RandomId)
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_MessagesSendMessage520c3870:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendMessage520c3870))
		// flags
		var flags int32 = 0
		if m.NoWebpage == true {
			flags |= 1 << 1
		}
		if m.Silent == true {
			flags |= 1 << 5
		}
		if m.Background == true {
			flags |= 1 << 6
		}
		if m.ClearDraft == true {
			flags |= 1 << 7
		}
		if m.ReplyToMsgId != 0 {
			flags |= 1 << 0
		}
		if m.ReplyMarkup != nil {
			flags |= 1 << 2
		}
		if len(m.Entities) > 0 {
			flags |= 1 << 3
		}
		if m.ScheduleDate != 0 {
			flags |= 1 << 10
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		if m.ReplyToMsgId != 0 {
			x.Int(m.ReplyToMsgId)
		}
		x.String(m.Message)
		x.Long(m.RandomId)
		if m.ReplyMarkup != nil {
			x.Bytes(m.ReplyMarkup.Encode(layer))
		}
		if len(m.Entities) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.Entities)))
			for _, v := range m.Entities {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.ScheduleDate != 0 {
			x.Int(m.ScheduleDate)
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesSendMessage Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSendMultiMedia) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSendMultiMedia, layer)

	switch cmd {
	case Cmd_MessagesSendMultiMedia:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendMultiMedia))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 5
		}
		if m.Background == true {
			flags |= 1 << 6
		}
		if m.ClearDraft == true {
			flags |= 1 << 7
		}
		if m.ReplyToMsgId != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		if m.ReplyToMsgId != 0 {
			x.Int(m.ReplyToMsgId)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.MultiMedia)))
		for _, v := range m.MultiMedia {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesSendMultiMediacc0110cb:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendMultiMediacc0110cb))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 5
		}
		if m.Background == true {
			flags |= 1 << 6
		}
		if m.ClearDraft == true {
			flags |= 1 << 7
		}
		if m.ReplyToMsgId != 0 {
			flags |= 1 << 0
		}
		if m.ScheduleDate != 0 {
			flags |= 1 << 10
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		if m.ReplyToMsgId != 0 {
			x.Int(m.ReplyToMsgId)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.MultiMedia)))
		for _, v := range m.MultiMedia {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.ScheduleDate != 0 {
			x.Int(m.ScheduleDate)
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesSendMultiMedia Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSendScheduledMessages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSendScheduledMessages, layer)

	switch cmd {
	case Cmd_MessagesSendScheduledMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendScheduledMessages))
		x.Bytes(m.Peer.Encode(layer))
		x.VectorInt(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("MessagesSendScheduledMessages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSendScreenshotNotification) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSendScreenshotNotification, layer)

	switch cmd {
	case Cmd_MessagesSendScreenshotNotification:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendScreenshotNotification))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.ReplyToMsgId)
		x.Long(m.RandomId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesSendScreenshotNotification Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSendVote) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSendVote, layer)

	switch cmd {
	case Cmd_MessagesSendVote:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSendVote))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.MsgId)

		return x.GetBuf()

	default:
		log.Errorf("MessagesSendVote Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSetBotCallbackAnswer) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSetBotCallbackAnswer, layer)

	switch cmd {
	case Cmd_MessagesSetBotCallbackAnswer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSetBotCallbackAnswer))
		// flags
		var flags int32 = 0
		if m.Alert == true {
			flags |= 1 << 1
		}
		if len(m.Message) > 0 {
			flags |= 1 << 0
		}
		if len(m.Url) > 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Long(m.QueryId)
		if len(m.Message) > 0 {
			x.String(m.Message)
		}
		if len(m.Url) > 0 {
			x.String(m.Url)
		}
		x.Int(m.CacheTime)

		return x.GetBuf()
	case Cmd_MessagesSetBotCallbackAnswer481c591a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSetBotCallbackAnswer481c591a))
		// flags
		var flags int32 = 0
		if m.Alert == true {
			flags |= 1 << 1
		}
		if len(m.Message) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.QueryId)
		if len(m.Message) > 0 {
			x.String(m.Message)
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesSetBotCallbackAnswer Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSetBotPrecheckoutResults) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSetBotPrecheckoutResults, layer)

	switch cmd {
	case Cmd_MessagesSetBotPrecheckoutResults:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSetBotPrecheckoutResults))
		// flags
		var flags int32 = 0
		if m.Success == true {
			flags |= 1 << 1
		}
		if len(m.Error) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.QueryId)
		if len(m.Error) > 0 {
			x.String(m.Error)
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesSetBotPrecheckoutResults Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSetBotShippingResults) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSetBotShippingResults, layer)

	switch cmd {
	case Cmd_MessagesSetBotShippingResults:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSetBotShippingResults))
		// flags
		var flags int32 = 0
		if len(m.Error) > 0 {
			flags |= 1 << 0
		}
		if len(m.ShippingOptions) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.QueryId)
		if len(m.Error) > 0 {
			x.String(m.Error)
		}
		if len(m.ShippingOptions) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.ShippingOptions)))
			for _, v := range m.ShippingOptions {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesSetBotShippingResults Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSetEncryptedTyping) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSetEncryptedTyping, layer)

	switch cmd {
	case Cmd_MessagesSetEncryptedTyping:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSetEncryptedTyping))
		x.Bytes(m.Peer.Encode(layer))
		x.Bytes(m.Typing.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesSetEncryptedTyping Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSetGameScore) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSetGameScore, layer)

	switch cmd {
	case Cmd_MessagesSetGameScore:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSetGameScore))
		// flags
		var flags int32 = 0
		if m.EditMessage == true {
			flags |= 1 << 0
		}
		if m.Force == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Id)
		x.Bytes(m.UserId.Encode(layer))
		x.Int(m.Score)

		return x.GetBuf()

	default:
		log.Errorf("MessagesSetGameScore Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSetInlineBotResults) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSetInlineBotResults, layer)

	switch cmd {
	case Cmd_MessagesSetInlineBotResults:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSetInlineBotResults))
		// flags
		var flags int32 = 0
		if m.Gallery == true {
			flags |= 1 << 0
		}
		if m.Private == true {
			flags |= 1 << 1
		}
		if len(m.NextOffset) > 0 {
			flags |= 1 << 2
		}
		if m.SwitchPm != nil {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Long(m.QueryId)
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Results)))
		for _, v := range m.Results {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.CacheTime)
		if len(m.NextOffset) > 0 {
			x.String(m.NextOffset)
		}
		if m.SwitchPm != nil {
			x.Bytes(m.SwitchPm.Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesSetInlineBotResults Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSetInlineGameScore) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSetInlineGameScore, layer)

	switch cmd {
	case Cmd_MessagesSetInlineGameScore:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSetInlineGameScore))
		// flags
		var flags int32 = 0
		if m.EditMessage == true {
			flags |= 1 << 0
		}
		if m.Force == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.Id.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))
		x.Int(m.Score)

		return x.GetBuf()

	default:
		log.Errorf("MessagesSetInlineGameScore Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSetTyping) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesSetTyping, layer)

	switch cmd {
	case Cmd_MessagesSetTyping:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSetTyping))
		x.Bytes(m.Peer.Encode(layer))
		x.Bytes(m.Action.Encode(layer))

		return x.GetBuf()
	case Cmd_MessagesSetTyping58943ee2:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSetTyping58943ee2))
		// flags
		var flags int32 = 0
		if m.TopMsgId != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		if m.TopMsgId != 0 {
			x.Int(m.TopMsgId)
		}
		x.Bytes(m.Action.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesSetTyping Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesStartBot) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesStartBot, layer)

	switch cmd {
	case Cmd_MessagesStartBot:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesStartBot))
		x.Bytes(m.Bot.Encode(layer))
		x.Bytes(m.Peer.Encode(layer))
		x.Long(m.RandomId)
		x.String(m.StartParam)

		return x.GetBuf()

	default:
		log.Errorf("MessagesStartBot Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesToggleChatAdmins) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesToggleChatAdmins, layer)

	switch cmd {
	case Cmd_MessagesToggleChatAdmins:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesToggleChatAdmins))
		x.Int(m.ChatId)
		x.Bytes(m.Enabled.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesToggleChatAdmins Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesToggleDialogPin) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesToggleDialogPin, layer)

	switch cmd {
	case Cmd_MessagesToggleDialogPin:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesToggleDialogPin))
		// flags
		var flags int32 = 0
		if m.Pinned == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer3289BE6A71.Encode(layer))

		return x.GetBuf()
	case Cmd_MessagesToggleDialogPina731e257:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesToggleDialogPina731e257))
		// flags
		var flags int32 = 0
		if m.Pinned == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.PeerA731E25785.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesToggleDialogPin Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesToggleStickerSets) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesToggleStickerSets, layer)

	switch cmd {
	case Cmd_MessagesToggleStickerSets:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesToggleStickerSets))
		// flags
		var flags int32 = 0
		if m.Uninstall == true {
			flags |= 1 << 0
		}
		if m.Archive == true {
			flags |= 1 << 1
		}
		if m.Unarchive == true {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Stickersets)))
		for _, v := range m.Stickersets {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesToggleStickerSets Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesUninstallStickerSet) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesUninstallStickerSet, layer)

	switch cmd {
	case Cmd_MessagesUninstallStickerSet:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesUninstallStickerSet))
		x.Bytes(m.Stickerset.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesUninstallStickerSet Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesUnpinAllMessages) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesUnpinAllMessages, layer)

	switch cmd {
	case Cmd_MessagesUnpinAllMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesUnpinAllMessages))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesUnpinAllMessages Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesUpdateDialogFilter) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesUpdateDialogFilter, layer)

	switch cmd {
	case Cmd_MessagesUpdateDialogFilter:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesUpdateDialogFilter))
		// flags
		var flags int32 = 0
		if m.Filter != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.Id)
		if m.Filter != nil {
			x.Bytes(m.Filter.Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("MessagesUpdateDialogFilter Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesUpdateDialogFiltersOrder) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesUpdateDialogFiltersOrder, layer)

	switch cmd {
	case Cmd_MessagesUpdateDialogFiltersOrder:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesUpdateDialogFiltersOrder))
		x.VectorInt(m.Order)

		return x.GetBuf()

	default:
		log.Errorf("MessagesUpdateDialogFiltersOrder Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesUpdatePinnedMessage) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesUpdatePinnedMessage, layer)

	switch cmd {
	case Cmd_MessagesUpdatePinnedMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesUpdatePinnedMessage))
		// flags
		var flags int32 = 0
		if m.Silent == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Id)

		return x.GetBuf()

	default:
		log.Errorf("MessagesUpdatePinnedMessage Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesUploadEncryptedFile) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesUploadEncryptedFile, layer)

	switch cmd {
	case Cmd_MessagesUploadEncryptedFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesUploadEncryptedFile))
		x.Bytes(m.Peer.Encode(layer))
		x.Bytes(m.File.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesUploadEncryptedFile Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesUploadMedia) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassMessagesUploadMedia, layer)

	switch cmd {
	case Cmd_MessagesUploadMedia:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesUploadMedia))
		x.Bytes(m.Peer.Encode(layer))
		x.Bytes(m.Media.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("MessagesUploadMedia Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsClearSavedInfo) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPaymentsClearSavedInfo, layer)

	switch cmd {
	case Cmd_PaymentsClearSavedInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsClearSavedInfo))
		// flags
		var flags int32 = 0
		if m.Credentials == true {
			flags |= 1 << 0
		}
		if m.Info == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("PaymentsClearSavedInfo Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsGetBankCardData) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPaymentsGetBankCardData, layer)

	switch cmd {
	case Cmd_PaymentsGetBankCardData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsGetBankCardData))
		x.String(m.Number)

		return x.GetBuf()

	default:
		log.Errorf("PaymentsGetBankCardData Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsGetPaymentForm) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPaymentsGetPaymentForm, layer)

	switch cmd {
	case Cmd_PaymentsGetPaymentForm:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsGetPaymentForm))
		x.Int(m.MsgId)

		return x.GetBuf()

	default:
		log.Errorf("PaymentsGetPaymentForm Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsGetPaymentReceipt) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPaymentsGetPaymentReceipt, layer)

	switch cmd {
	case Cmd_PaymentsGetPaymentReceipt:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsGetPaymentReceipt))
		x.Int(m.MsgId)

		return x.GetBuf()

	default:
		log.Errorf("PaymentsGetPaymentReceipt Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsGetSavedInfo) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPaymentsGetSavedInfo, layer)

	switch cmd {
	case Cmd_PaymentsGetSavedInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsGetSavedInfo))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("PaymentsGetSavedInfo Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsSendPaymentForm) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPaymentsSendPaymentForm, layer)

	switch cmd {
	case Cmd_PaymentsSendPaymentForm:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsSendPaymentForm))
		// flags
		var flags int32 = 0
		if len(m.RequestedInfoId) > 0 {
			flags |= 1 << 0
		}
		if len(m.ShippingOptionId) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.MsgId)
		if len(m.RequestedInfoId) > 0 {
			x.String(m.RequestedInfoId)
		}
		if len(m.ShippingOptionId) > 0 {
			x.String(m.ShippingOptionId)
		}
		x.Bytes(m.Credentials.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PaymentsSendPaymentForm Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsValidateRequestedInfo) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPaymentsValidateRequestedInfo, layer)

	switch cmd {
	case Cmd_PaymentsValidateRequestedInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsValidateRequestedInfo))
		// flags
		var flags int32 = 0
		if m.Save == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.MsgId)
		x.Bytes(m.Info.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PaymentsValidateRequestedInfo Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneAcceptCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneAcceptCall, layer)

	switch cmd {
	case Cmd_PhoneAcceptCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneAcceptCall))
		x.Bytes(m.Peer.Encode(layer))
		x.StringBytes(m.GB)
		x.Bytes(m.Protocol.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PhoneAcceptCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCheckGroupCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneCheckGroupCall, layer)

	switch cmd {
	case Cmd_PhoneCheckGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCheckGroupCall))
		x.Bytes(m.Call.Encode(layer))
		x.Int(m.Source)

		return x.GetBuf()

	default:
		log.Errorf("PhoneCheckGroupCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneConfirmCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneConfirmCall, layer)

	switch cmd {
	case Cmd_PhoneConfirmCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneConfirmCall))
		x.Bytes(m.Peer.Encode(layer))
		x.StringBytes(m.GA)
		x.Long(m.KeyFingerprint)
		x.Bytes(m.Protocol.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PhoneConfirmCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCreateGroupCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneCreateGroupCall, layer)

	switch cmd {
	case Cmd_PhoneCreateGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCreateGroupCall))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.RandomId)

		return x.GetBuf()

	default:
		log.Errorf("PhoneCreateGroupCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneDiscardCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneDiscardCall, layer)

	switch cmd {
	case Cmd_PhoneDiscardCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneDiscardCall))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Duration)
		x.Bytes(m.Reason.Encode(layer))
		x.Long(m.ConnectionId)

		return x.GetBuf()
	case Cmd_PhoneDiscardCallb2cbc1c0:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneDiscardCallb2cbc1c0))
		// flags
		var flags int32 = 0
		if m.Video == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Duration)
		x.Bytes(m.Reason.Encode(layer))
		x.Long(m.ConnectionId)

		return x.GetBuf()

	default:
		log.Errorf("PhoneDiscardCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneDiscardGroupCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneDiscardGroupCall, layer)

	switch cmd {
	case Cmd_PhoneDiscardGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneDiscardGroupCall))
		x.Bytes(m.Call.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PhoneDiscardGroupCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneEditGroupCallMember) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneEditGroupCallMember, layer)

	switch cmd {
	case Cmd_PhoneEditGroupCallMember:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneEditGroupCallMember))
		// flags
		var flags int32 = 0
		if m.Muted == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Call.Encode(layer))
		x.Bytes(m.UserId.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PhoneEditGroupCallMember Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneGetCallConfig) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneGetCallConfig, layer)

	switch cmd {
	case Cmd_PhoneGetCallConfig:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneGetCallConfig))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("PhoneGetCallConfig Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneGetGroupCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneGetGroupCall, layer)

	switch cmd {
	case Cmd_PhoneGetGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneGetGroupCall))
		x.Bytes(m.Call.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PhoneGetGroupCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneGetGroupParticipants) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneGetGroupParticipants, layer)

	switch cmd {
	case Cmd_PhoneGetGroupParticipants:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneGetGroupParticipants))
		x.Bytes(m.Call.Encode(layer))
		x.VectorInt(m.Ids)

		x.VectorInt(m.Sources)

		x.String(m.Offset)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("PhoneGetGroupParticipants Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneInviteToGroupCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneInviteToGroupCall, layer)

	switch cmd {
	case Cmd_PhoneInviteToGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneInviteToGroupCall))
		x.Bytes(m.Call.Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Users)))
		for _, v := range m.Users {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("PhoneInviteToGroupCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneJoinGroupCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneJoinGroupCall, layer)

	switch cmd {
	case Cmd_PhoneJoinGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneJoinGroupCall))
		// flags
		var flags int32 = 0
		if m.Muted == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Call.Encode(layer))
		x.Bytes(m.Params.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PhoneJoinGroupCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneLeaveGroupCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneLeaveGroupCall, layer)

	switch cmd {
	case Cmd_PhoneLeaveGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneLeaveGroupCall))
		x.Bytes(m.Call.Encode(layer))
		x.Int(m.Source)

		return x.GetBuf()

	default:
		log.Errorf("PhoneLeaveGroupCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneReceivedCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneReceivedCall, layer)

	switch cmd {
	case Cmd_PhoneReceivedCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneReceivedCall))
		x.Bytes(m.Peer.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PhoneReceivedCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneRequestCall) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneRequestCall, layer)

	switch cmd {
	case Cmd_PhoneRequestCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneRequestCall))
		x.Bytes(m.UserId.Encode(layer))
		x.Int(m.RandomId)
		x.StringBytes(m.GAHash)
		x.Bytes(m.Protocol.Encode(layer))

		return x.GetBuf()
	case Cmd_PhoneRequestCall42ff96ed:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneRequestCall42ff96ed))
		// flags
		var flags int32 = 0
		if m.Video == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.UserId.Encode(layer))
		x.Int(m.RandomId)
		x.StringBytes(m.GAHash)
		x.Bytes(m.Protocol.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PhoneRequestCall Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneSaveCallDebug) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneSaveCallDebug, layer)

	switch cmd {
	case Cmd_PhoneSaveCallDebug:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneSaveCallDebug))
		x.Bytes(m.Peer.Encode(layer))
		x.Bytes(m.Debug.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PhoneSaveCallDebug Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneSendSignalingData) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneSendSignalingData, layer)

	switch cmd {
	case Cmd_PhoneSendSignalingData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneSendSignalingData))
		x.Bytes(m.Peer.Encode(layer))
		x.StringBytes(m.Data)

		return x.GetBuf()

	default:
		log.Errorf("PhoneSendSignalingData Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneSetCallRating) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneSetCallRating, layer)

	switch cmd {
	case Cmd_PhoneSetCallRating:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneSetCallRating))
		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Rating)
		x.String(m.Comment)

		return x.GetBuf()
	case Cmd_PhoneSetCallRating59ead627:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneSetCallRating59ead627))
		// flags
		var flags int32 = 0
		if m.UserInitiative == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Peer.Encode(layer))
		x.Int(m.Rating)
		x.String(m.Comment)

		return x.GetBuf()

	default:
		log.Errorf("PhoneSetCallRating Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneToggleGroupCallSettings) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhoneToggleGroupCallSettings, layer)

	switch cmd {
	case Cmd_PhoneToggleGroupCallSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneToggleGroupCallSettings))
		// flags
		var flags int32 = 0
		if m.JoinMuted != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Call.Encode(layer))
		if m.JoinMuted != nil {
			x.Bytes(m.JoinMuted.Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("PhoneToggleGroupCallSettings Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotosDeletePhotos) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhotosDeletePhotos, layer)

	switch cmd {
	case Cmd_PhotosDeletePhotos:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotosDeletePhotos))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Id)))
		for _, v := range m.Id {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("PhotosDeletePhotos Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotosGetUserPhotos) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhotosGetUserPhotos, layer)

	switch cmd {
	case Cmd_PhotosGetUserPhotos:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotosGetUserPhotos))
		x.Bytes(m.UserId.Encode(layer))
		x.Int(m.Offset)
		x.Long(m.MaxId)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("PhotosGetUserPhotos Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotosUpdateProfilePhoto) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhotosUpdateProfilePhoto, layer)

	switch cmd {
	case Cmd_PhotosUpdateProfilePhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotosUpdateProfilePhoto))
		x.Bytes(m.Id.Encode(layer))

		return x.GetBuf()
	case Cmd_PhotosUpdateProfilePhoto72d4742c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotosUpdateProfilePhoto72d4742c))
		x.Bytes(m.Id.Encode(layer))

		return x.GetBuf()
	case Cmd_PhotosUpdateProfilePhotoeef579a0:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotosUpdateProfilePhotoeef579a0))
		x.Bytes(m.Id.Encode(layer))
		x.Bytes(m.Crop.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PhotosUpdateProfilePhoto Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotosUploadProfilePhoto) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPhotosUploadProfilePhoto, layer)

	switch cmd {
	case Cmd_PhotosUploadProfilePhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotosUploadProfilePhoto))
		x.Bytes(m.File.Encode(layer))

		return x.GetBuf()
	case Cmd_PhotosUploadProfilePhoto89f30f69:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotosUploadProfilePhoto89f30f69))
		// flags
		var flags int32 = 0
		if m.File != nil {
			flags |= 1 << 0
		}
		if m.Video != nil {
			flags |= 1 << 1
		}
		if doubleIsEqual(m.VideoStartTs, 0, 0.00000001) == true {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.File != nil {
			x.Bytes(m.File.Encode(layer))
		}
		if m.Video != nil {
			x.Bytes(m.Video.Encode(layer))
		}

		return x.GetBuf()
	case Cmd_PhotosUploadProfilePhotod50f9c88:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotosUploadProfilePhotod50f9c88))
		x.Bytes(m.File.Encode(layer))
		x.String(m.Caption)
		x.Bytes(m.GeoPoint.Encode(layer))
		x.Bytes(m.Crop.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("PhotosUploadProfilePhoto Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPing) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPing, layer)

	switch cmd {
	case Cmd_Ping:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Ping))
		x.Long(m.PingId)

		return x.GetBuf()

	default:
		log.Errorf("Ping Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPingDelayDisconnect) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassPingDelayDisconnect, layer)

	switch cmd {
	case Cmd_PingDelayDisconnect:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PingDelayDisconnect))
		x.Long(m.PingId)
		x.Int(m.DisconnectDelay)

		return x.GetBuf()

	default:
		log.Errorf("PingDelayDisconnect Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLReqPq) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassReqPq, layer)

	switch cmd {
	case Cmd_ReqPq:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ReqPq))
		x.Bytes(m.Nonce)

		return x.GetBuf()

	default:
		log.Errorf("ReqPq Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLReqPqMulti) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassReqPqMulti, layer)

	switch cmd {
	case Cmd_ReqPqMulti:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ReqPqMulti))
		x.Bytes(m.Nonce)

		return x.GetBuf()

	default:
		log.Errorf("ReqPqMulti Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLReq_DHParams) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassReq_DHParams, layer)

	switch cmd {
	case Cmd_Req_DHParams:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Req_DHParams))
		x.Bytes(m.Nonce)
		x.Bytes(m.ServerNonce)
		x.String(m.P)
		x.String(m.Q)
		x.Long(m.PublicKeyFingerprint)
		x.String(m.EncryptedData)

		return x.GetBuf()

	default:
		log.Errorf("Req_DHParams Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLRpcDropAnswer) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassRpcDropAnswer, layer)

	switch cmd {
	case Cmd_RpcDropAnswer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_RpcDropAnswer))
		x.Long(m.ReqMsgId)

		return x.GetBuf()

	default:
		log.Errorf("RpcDropAnswer Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSetClient_DHParams) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassSetClient_DHParams, layer)

	switch cmd {
	case Cmd_SetClient_DHParams:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SetClient_DHParams))
		x.Bytes(m.Nonce)
		x.Bytes(m.ServerNonce)
		x.String(m.EncryptedData)

		return x.GetBuf()

	default:
		log.Errorf("SetClient_DHParams Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsGetBroadcastStats) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassStatsGetBroadcastStats, layer)

	switch cmd {
	case Cmd_StatsGetBroadcastStats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsGetBroadcastStats))
		// flags
		var flags int32 = 0
		if m.Dark == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Channel.Encode(layer))

		return x.GetBuf()
	case Cmd_StatsGetBroadcastStatse6300dba:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsGetBroadcastStatse6300dba))
		// flags
		var flags int32 = 0
		if m.Dark == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Channel.Encode(layer))
		x.Int(m.TzOffset)

		return x.GetBuf()

	default:
		log.Errorf("StatsGetBroadcastStats Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsGetMegagroupStats) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassStatsGetMegagroupStats, layer)

	switch cmd {
	case Cmd_StatsGetMegagroupStats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsGetMegagroupStats))
		// flags
		var flags int32 = 0
		if m.Dark == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Channel.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("StatsGetMegagroupStats Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsGetMessagePublicForwards) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassStatsGetMessagePublicForwards, layer)

	switch cmd {
	case Cmd_StatsGetMessagePublicForwards:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsGetMessagePublicForwards))
		x.Bytes(m.Channel.Encode(layer))
		x.Int(m.MsgId)
		x.Int(m.OffsetRate)
		x.Bytes(m.OffsetPeer.Encode(layer))
		x.Int(m.OffsetId)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("StatsGetMessagePublicForwards Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsGetMessageStats) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassStatsGetMessageStats, layer)

	switch cmd {
	case Cmd_StatsGetMessageStats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsGetMessageStats))
		// flags
		var flags int32 = 0
		if m.Dark == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Channel.Encode(layer))
		x.Int(m.MsgId)

		return x.GetBuf()

	default:
		log.Errorf("StatsGetMessageStats Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsLoadAsyncGraph) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassStatsLoadAsyncGraph, layer)

	switch cmd {
	case Cmd_StatsLoadAsyncGraph:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsLoadAsyncGraph))
		// flags
		var flags int32 = 0
		if m.X != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.Token)
		if m.X != 0 {
			x.Long(m.X)
		}

		return x.GetBuf()

	default:
		log.Errorf("StatsLoadAsyncGraph Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStickersAddStickerToSet) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassStickersAddStickerToSet, layer)

	switch cmd {
	case Cmd_StickersAddStickerToSet:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickersAddStickerToSet))
		x.Bytes(m.Stickerset.Encode(layer))
		x.Bytes(m.Sticker.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("StickersAddStickerToSet Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStickersChangeStickerPosition) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassStickersChangeStickerPosition, layer)

	switch cmd {
	case Cmd_StickersChangeStickerPosition:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickersChangeStickerPosition))
		x.Bytes(m.Sticker.Encode(layer))
		x.Int(m.Position)

		return x.GetBuf()

	default:
		log.Errorf("StickersChangeStickerPosition Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStickersCreateStickerSet) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassStickersCreateStickerSet, layer)

	switch cmd {
	case Cmd_StickersCreateStickerSet:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickersCreateStickerSet))
		// flags
		var flags int32 = 0
		if m.Masks == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.UserId.Encode(layer))
		x.String(m.Title)
		x.String(m.ShortName)
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Stickers)))
		for _, v := range m.Stickers {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_StickersCreateStickerSetf1036780:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickersCreateStickerSetf1036780))
		// flags
		var flags int32 = 0
		if m.Masks == true {
			flags |= 1 << 0
		}
		if m.Animated == true {
			flags |= 1 << 1
		}
		if m.Thumb != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.UserId.Encode(layer))
		x.String(m.Title)
		x.String(m.ShortName)
		if m.Thumb != nil {
			x.Bytes(m.Thumb.Encode(layer))
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Stickers)))
		for _, v := range m.Stickers {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("StickersCreateStickerSet Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStickersRemoveStickerFromSet) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassStickersRemoveStickerFromSet, layer)

	switch cmd {
	case Cmd_StickersRemoveStickerFromSet:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickersRemoveStickerFromSet))
		x.Bytes(m.Sticker.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("StickersRemoveStickerFromSet Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStickersSetStickerSetThumb) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassStickersSetStickerSetThumb, layer)

	switch cmd {
	case Cmd_StickersSetStickerSetThumb:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickersSetStickerSetThumb))
		x.Bytes(m.Stickerset.Encode(layer))
		x.Bytes(m.Thumb.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("StickersSetStickerSetThumb Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesGetChannelDifference) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUpdatesGetChannelDifference, layer)

	switch cmd {
	case Cmd_UpdatesGetChannelDifference:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesGetChannelDifference))
		// flags
		var flags int32 = 0
		if m.Force == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.Filter.Encode(layer))
		x.Int(m.Pts)
		x.Int(m.Limit)

		return x.GetBuf()
	case Cmd_UpdatesGetChannelDifferencebb32d7c0:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesGetChannelDifferencebb32d7c0))
		x.Bytes(m.Channel.Encode(layer))
		x.Bytes(m.Filter.Encode(layer))
		x.Int(m.Pts)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("UpdatesGetChannelDifference Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesGetDifference) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUpdatesGetDifference, layer)

	switch cmd {
	case Cmd_UpdatesGetDifference:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesGetDifference))
		// flags
		var flags int32 = 0
		if m.PtsTotalLimit != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.Pts)
		if m.PtsTotalLimit != 0 {
			x.Int(m.PtsTotalLimit)
		}
		x.Int(m.Date)
		x.Int(m.Qts)

		return x.GetBuf()
	case Cmd_UpdatesGetDifferencea041495:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesGetDifferencea041495))
		x.Int(m.Pts)
		x.Int(m.Date)
		x.Int(m.Qts)

		return x.GetBuf()

	default:
		log.Errorf("UpdatesGetDifference Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesGetState) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUpdatesGetState, layer)

	switch cmd {
	case Cmd_UpdatesGetState:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesGetState))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("UpdatesGetState Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadGetCdnFile) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUploadGetCdnFile, layer)

	switch cmd {
	case Cmd_UploadGetCdnFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadGetCdnFile))
		x.StringBytes(m.FileToken)
		x.Int(m.Offset)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("UploadGetCdnFile Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadGetCdnFileHashes) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUploadGetCdnFileHashes, layer)

	switch cmd {
	case Cmd_UploadGetCdnFileHashes:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadGetCdnFileHashes))
		x.StringBytes(m.FileToken)
		x.Int(m.Offset)

		return x.GetBuf()
	case Cmd_UploadGetCdnFileHashes4da54231:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadGetCdnFileHashes4da54231))
		x.StringBytes(m.FileToken)
		x.Int(m.Offset)

		return x.GetBuf()

	default:
		log.Errorf("UploadGetCdnFileHashes Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadGetFile) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUploadGetFile, layer)

	switch cmd {
	case Cmd_UploadGetFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadGetFile))
		x.Bytes(m.Location.Encode(layer))
		x.Int(m.Offset)
		x.Int(m.Limit)

		return x.GetBuf()
	case Cmd_UploadGetFileb15a9afc:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadGetFileb15a9afc))
		// flags
		var flags int32 = 0
		if m.Precise == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.Location.Encode(layer))
		x.Int(m.Offset)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("UploadGetFile Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadGetFileHashes) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUploadGetFileHashes, layer)

	switch cmd {
	case Cmd_UploadGetFileHashes:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadGetFileHashes))
		x.Bytes(m.Location.Encode(layer))
		x.Int(m.Offset)

		return x.GetBuf()

	default:
		log.Errorf("UploadGetFileHashes Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadGetWebFile) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUploadGetWebFile, layer)

	switch cmd {
	case Cmd_UploadGetWebFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadGetWebFile))
		x.Bytes(m.Location.Encode(layer))
		x.Int(m.Offset)
		x.Int(m.Limit)

		return x.GetBuf()

	default:
		log.Errorf("UploadGetWebFile Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadReuploadCdnFile) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUploadReuploadCdnFile, layer)

	switch cmd {
	case Cmd_UploadReuploadCdnFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadReuploadCdnFile))
		x.StringBytes(m.FileToken)
		x.StringBytes(m.RequestToken)

		return x.GetBuf()
	case Cmd_UploadReuploadCdnFile9b2754a8:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadReuploadCdnFile9b2754a8))
		x.StringBytes(m.FileToken)
		x.StringBytes(m.RequestToken)

		return x.GetBuf()

	default:
		log.Errorf("UploadReuploadCdnFile Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadSaveBigFilePart) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUploadSaveBigFilePart, layer)

	switch cmd {
	case Cmd_UploadSaveBigFilePart:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadSaveBigFilePart))
		x.Long(m.FileId)
		x.Int(m.FilePart)
		x.Int(m.FileTotalParts)
		x.StringBytes(m.Bytes)

		return x.GetBuf()

	default:
		log.Errorf("UploadSaveBigFilePart Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadSaveFilePart) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUploadSaveFilePart, layer)

	switch cmd {
	case Cmd_UploadSaveFilePart:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadSaveFilePart))
		x.Long(m.FileId)
		x.Int(m.FilePart)
		x.StringBytes(m.Bytes)

		return x.GetBuf()

	default:
		log.Errorf("UploadSaveFilePart Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUsersGetFullUser) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUsersGetFullUser, layer)

	switch cmd {
	case Cmd_UsersGetFullUser:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UsersGetFullUser))
		x.Bytes(m.Id.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("UsersGetFullUser Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUsersGetUsers) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUsersGetUsers, layer)

	switch cmd {
	case Cmd_UsersGetUsers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UsersGetUsers))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Id)))
		for _, v := range m.Id {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("UsersGetUsers Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUsersSetSecureValueErrors) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassUsersSetSecureValueErrors, layer)

	switch cmd {
	case Cmd_UsersSetSecureValueErrors:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UsersSetSecureValueErrors))
		x.Bytes(m.Id.Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.Errors)))
		for _, v := range m.Errors {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("UsersSetSecureValueErrors Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWalletGetKeySecretSalt) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassWalletGetKeySecretSalt, layer)

	switch cmd {
	case Cmd_WalletGetKeySecretSalt:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WalletGetKeySecretSalt))
		x.Bytes(m.Revoke.Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("WalletGetKeySecretSalt Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWalletSendLiteRequest) Encode(layer int32) []byte {
	cmd := GetClassIDByLayer(ClassWalletSendLiteRequest, layer)

	switch cmd {
	case Cmd_WalletSendLiteRequest:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WalletSendLiteRequest))
		x.StringBytes(m.Body)

		return x.GetBuf()

	default:
		log.Errorf("WalletSendLiteRequest Encode Invalid layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}
	return nil
}
