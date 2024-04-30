/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2023/06/12 18:03
 * @File : codec_schema_decode.pb.go
 * @Desc :
 *
 */

package mtproto

import (
	"fmt"
)

func (m *AccessPointRule) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccessPointRule:
		m2 := NewTLAccessPointRule(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_AccessPointRule()

	default:
		return fmt.Errorf("AccessPointRule Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *AccountDaysTTL) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountDaysTTL:
		m2 := NewTLAccountDaysTTL(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_AccountDaysTTL()

	default:
		return fmt.Errorf("AccountDaysTTL Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_AuthorizationForm) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountAuthorizationForm:
		m2 := NewTLAccountAuthorizationForm(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_AuthorizationForm()

	default:
		return fmt.Errorf("Account_AuthorizationForm Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_Authorizations) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountAuthorizations:
		m2 := NewTLAccountAuthorizations(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_Authorizations()

	default:
		return fmt.Errorf("Account_Authorizations Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_AutoDownloadSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountAutoDownloadSettings:
		m2 := NewTLAccountAutoDownloadSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_AutoDownloadSettings()

	default:
		return fmt.Errorf("Account_AutoDownloadSettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_ContentSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountContentSettings:
		m2 := NewTLAccountContentSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_ContentSettings()

	default:
		return fmt.Errorf("Account_ContentSettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_Password) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountNoPassword:
		m2 := NewTLAccountNoPassword(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_Password()
	case Cmd_AccountPassword:
		m2 := NewTLAccountPassword(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_Password()

	case Cmd_AccountPasswordad2641f8:
		m2 := NewTLAccountPassword(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_Password()

	default:
		return fmt.Errorf("Account_Password Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_PasswordInputSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountPasswordInputSettings:
		m2 := NewTLAccountPasswordInputSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_PasswordInputSettings()

	case Cmd_AccountPasswordInputSettingsc23727c9:
		m2 := NewTLAccountPasswordInputSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_PasswordInputSettings()

	default:
		return fmt.Errorf("Account_PasswordInputSettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_PasswordSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountPasswordSettings:
		m2 := NewTLAccountPasswordSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_PasswordSettings()

	case Cmd_AccountPasswordSettings9a5c33e5:
		m2 := NewTLAccountPasswordSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_PasswordSettings()

	default:
		return fmt.Errorf("Account_PasswordSettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_PrivacyRules) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountPrivacyRules:
		m2 := NewTLAccountPrivacyRules(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_PrivacyRules()

	case Cmd_AccountPrivacyRules50a04e45:
		m2 := NewTLAccountPrivacyRules(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_PrivacyRules()

	default:
		return fmt.Errorf("Account_PrivacyRules Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_SentEmailCode) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountSentEmailCode:
		m2 := NewTLAccountSentEmailCode(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_SentEmailCode()

	default:
		return fmt.Errorf("Account_SentEmailCode Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_Takeout) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountTakeout:
		m2 := NewTLAccountTakeout(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_Takeout()

	default:
		return fmt.Errorf("Account_Takeout Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_Themes) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountThemesNotModified:
		m2 := NewTLAccountThemesNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_Themes()
	case Cmd_AccountThemes:
		m2 := NewTLAccountThemes(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_Themes()

	default:
		return fmt.Errorf("Account_Themes Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_TmpPassword) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountTmpPassword:
		m2 := NewTLAccountTmpPassword(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_TmpPassword()

	default:
		return fmt.Errorf("Account_TmpPassword Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_WallPapers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountWallPapersNotModified:
		m2 := NewTLAccountWallPapersNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_WallPapers()
	case Cmd_AccountWallPapers:
		m2 := NewTLAccountWallPapers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_WallPapers()

	default:
		return fmt.Errorf("Account_WallPapers Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Account_WebAuthorizations) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AccountWebAuthorizations:
		m2 := NewTLAccountWebAuthorizations(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Account_WebAuthorizations()

	default:
		return fmt.Errorf("Account_WebAuthorizations Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Auth_Authorization) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AuthAuthorization:
		m2 := NewTLAuthAuthorization(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_Authorization()
	case Cmd_AuthAuthorizationSignUpRequired:
		m2 := NewTLAuthAuthorizationSignUpRequired(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_Authorization()

	case Cmd_AuthAuthorizationff036af1:
		m2 := NewTLAuthAuthorization(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_Authorization()

	default:
		return fmt.Errorf("Auth_Authorization Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Auth_CheckedPhone) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AuthCheckedPhone:
		m2 := NewTLAuthCheckedPhone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_CheckedPhone()

	default:
		return fmt.Errorf("Auth_CheckedPhone Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Auth_CodeType) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AuthCodeTypeSms:
		m2 := NewTLAuthCodeTypeSms(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_CodeType()
	case Cmd_AuthCodeTypeCall:
		m2 := NewTLAuthCodeTypeCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_CodeType()
	case Cmd_AuthCodeTypeFlashCall:
		m2 := NewTLAuthCodeTypeFlashCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_CodeType()

	default:
		return fmt.Errorf("Auth_CodeType Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Auth_ExportedAuthorization) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AuthExportedAuthorization:
		m2 := NewTLAuthExportedAuthorization(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_ExportedAuthorization()

	default:
		return fmt.Errorf("Auth_ExportedAuthorization Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Auth_LoginToken) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AuthLoginToken:
		m2 := NewTLAuthLoginToken(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_LoginToken()
	case Cmd_AuthLoginTokenMigrateTo:
		m2 := NewTLAuthLoginTokenMigrateTo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_LoginToken()
	case Cmd_AuthLoginTokenSuccess:
		m2 := NewTLAuthLoginTokenSuccess(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_LoginToken()

	default:
		return fmt.Errorf("Auth_LoginToken Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Auth_PasswordRecovery) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AuthPasswordRecovery:
		m2 := NewTLAuthPasswordRecovery(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_PasswordRecovery()

	default:
		return fmt.Errorf("Auth_PasswordRecovery Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Auth_SentCode) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AuthSentCode:
		m2 := NewTLAuthSentCode(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_SentCode()

	case Cmd_AuthSentCode38faab5f:
		m2 := NewTLAuthSentCode(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_SentCode()

	default:
		return fmt.Errorf("Auth_SentCode Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Auth_SentCodeType) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AuthSentCodeTypeApp:
		m2 := NewTLAuthSentCodeTypeApp(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_SentCodeType()
	case Cmd_AuthSentCodeTypeSms:
		m2 := NewTLAuthSentCodeTypeSms(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_SentCodeType()
	case Cmd_AuthSentCodeTypeCall:
		m2 := NewTLAuthSentCodeTypeCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_SentCodeType()
	case Cmd_AuthSentCodeTypeFlashCall:
		m2 := NewTLAuthSentCodeTypeFlashCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Auth_SentCodeType()

	default:
		return fmt.Errorf("Auth_SentCodeType Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Authorization) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Authorization:
		m2 := NewTLAuthorization(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Authorization()

	case Cmd_Authorizationad01d61d:
		m2 := NewTLAuthorization(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Authorization()

	default:
		return fmt.Errorf("Authorization Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *AutoDownloadSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_AutoDownloadSettings:
		m2 := NewTLAutoDownloadSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_AutoDownloadSettings()

	case Cmd_AutoDownloadSettingse04232f3:
		m2 := NewTLAutoDownloadSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_AutoDownloadSettings()

	default:
		return fmt.Errorf("AutoDownloadSettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *BadMsgNotification) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_BadMsgNotification:
		m2 := NewTLBadMsgNotification(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BadMsgNotification()
	case Cmd_BadServerSalt:
		m2 := NewTLBadServerSalt(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BadMsgNotification()

	default:
		return fmt.Errorf("BadMsgNotification Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *BankCardOpenUrl) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_BankCardOpenUrl:
		m2 := NewTLBankCardOpenUrl(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BankCardOpenUrl()

	default:
		return fmt.Errorf("BankCardOpenUrl Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *BaseTheme) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_BaseThemeClassic:
		m2 := NewTLBaseThemeClassic(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BaseTheme()
	case Cmd_BaseThemeDay:
		m2 := NewTLBaseThemeDay(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BaseTheme()
	case Cmd_BaseThemeNight:
		m2 := NewTLBaseThemeNight(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BaseTheme()
	case Cmd_BaseThemeTinted:
		m2 := NewTLBaseThemeTinted(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BaseTheme()
	case Cmd_BaseThemeArctic:
		m2 := NewTLBaseThemeArctic(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BaseTheme()

	default:
		return fmt.Errorf("BaseTheme Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *BindAuthKeyInner) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_BindAuthKeyInner:
		m2 := NewTLBindAuthKeyInner(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BindAuthKeyInner()

	default:
		return fmt.Errorf("BindAuthKeyInner Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Bool) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_BoolFalse:
		m2 := NewTLBoolFalse(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Bool()
	case Cmd_BoolTrue:
		m2 := NewTLBoolTrue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Bool()

	default:
		return fmt.Errorf("Bool Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *BotCommand) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_BotCommand:
		m2 := NewTLBotCommand(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotCommand()

	default:
		return fmt.Errorf("BotCommand Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *BotInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_BotInfo:
		m2 := NewTLBotInfo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInfo()

	default:
		return fmt.Errorf("BotInfo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *BotInlineMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_BotInlineMessageMediaAuto:
		m2 := NewTLBotInlineMessageMediaAuto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineMessage()
	case Cmd_BotInlineMessageText:
		m2 := NewTLBotInlineMessageText(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineMessage()
	case Cmd_BotInlineMessageMediaGeo:
		m2 := NewTLBotInlineMessageMediaGeo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineMessage()
	case Cmd_BotInlineMessageMediaVenue:
		m2 := NewTLBotInlineMessageMediaVenue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineMessage()
	case Cmd_BotInlineMessageMediaContact:
		m2 := NewTLBotInlineMessageMediaContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineMessage()

	case Cmd_BotInlineMessageMediaAuto764cf810:
		m2 := NewTLBotInlineMessageMediaAuto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineMessage()
	case Cmd_BotInlineMessageMediaGeob722de65:
		m2 := NewTLBotInlineMessageMediaGeo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineMessage()
	case Cmd_BotInlineMessageMediaVenue8a86659c:
		m2 := NewTLBotInlineMessageMediaVenue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineMessage()
	case Cmd_BotInlineMessageMediaContact18d1cdc2:
		m2 := NewTLBotInlineMessageMediaContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineMessage()
	case Cmd_BotInlineMessageMediaGeo51846fd:
		m2 := NewTLBotInlineMessageMediaGeo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineMessage()

	default:
		return fmt.Errorf("BotInlineMessage Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *BotInlineResult) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_BotInlineResult:
		m2 := NewTLBotInlineResult(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineResult()
	case Cmd_BotInlineMediaResult:
		m2 := NewTLBotInlineMediaResult(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineResult()

	case Cmd_BotInlineResult11965f3a:
		m2 := NewTLBotInlineResult(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_BotInlineResult()

	default:
		return fmt.Errorf("BotInlineResult Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *CdnConfig) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_CdnConfig:
		m2 := NewTLCdnConfig(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_CdnConfig()

	default:
		return fmt.Errorf("CdnConfig Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *CdnFileHash) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_CdnFileHash:
		m2 := NewTLCdnFileHash(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_CdnFileHash()

	default:
		return fmt.Errorf("CdnFileHash Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *CdnPublicKey) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_CdnPublicKey:
		m2 := NewTLCdnPublicKey(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_CdnPublicKey()

	default:
		return fmt.Errorf("CdnPublicKey Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChannelAdminLogEvent) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelAdminLogEvent:
		m2 := NewTLChannelAdminLogEvent(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEvent()

	default:
		return fmt.Errorf("ChannelAdminLogEvent Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChannelAdminLogEventAction) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelAdminLogEventActionChangeTitle:
		m2 := NewTLChannelAdminLogEventActionChangeTitle(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionChangeAbout:
		m2 := NewTLChannelAdminLogEventActionChangeAbout(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionChangeUsername:
		m2 := NewTLChannelAdminLogEventActionChangeUsername(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionChangePhoto:
		m2 := NewTLChannelAdminLogEventActionChangePhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionToggleInvites:
		m2 := NewTLChannelAdminLogEventActionToggleInvites(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionToggleSignatures:
		m2 := NewTLChannelAdminLogEventActionToggleSignatures(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionUpdatePinned:
		m2 := NewTLChannelAdminLogEventActionUpdatePinned(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionEditMessage:
		m2 := NewTLChannelAdminLogEventActionEditMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionDeleteMessage:
		m2 := NewTLChannelAdminLogEventActionDeleteMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionParticipantJoin:
		m2 := NewTLChannelAdminLogEventActionParticipantJoin(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionParticipantLeave:
		m2 := NewTLChannelAdminLogEventActionParticipantLeave(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionParticipantInvite:
		m2 := NewTLChannelAdminLogEventActionParticipantInvite(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionParticipantToggleBan:
		m2 := NewTLChannelAdminLogEventActionParticipantToggleBan(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionParticipantToggleAdmin:
		m2 := NewTLChannelAdminLogEventActionParticipantToggleAdmin(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionChangeStickerSet:
		m2 := NewTLChannelAdminLogEventActionChangeStickerSet(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionTogglePreHistoryHidden:
		m2 := NewTLChannelAdminLogEventActionTogglePreHistoryHidden(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionDefaultBannedRights:
		m2 := NewTLChannelAdminLogEventActionDefaultBannedRights(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionStopPoll:
		m2 := NewTLChannelAdminLogEventActionStopPoll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionChangeLinkedChat:
		m2 := NewTLChannelAdminLogEventActionChangeLinkedChat(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionChangeLocation:
		m2 := NewTLChannelAdminLogEventActionChangeLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionToggleSlowMode:
		m2 := NewTLChannelAdminLogEventActionToggleSlowMode(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionStartGroupCall:
		m2 := NewTLChannelAdminLogEventActionStartGroupCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionDiscardGroupCall:
		m2 := NewTLChannelAdminLogEventActionDiscardGroupCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionParticipantMute:
		m2 := NewTLChannelAdminLogEventActionParticipantMute(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionParticipantUnmute:
		m2 := NewTLChannelAdminLogEventActionParticipantUnmute(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()
	case Cmd_ChannelAdminLogEventActionToggleGroupCallSetting:
		m2 := NewTLChannelAdminLogEventActionToggleGroupCallSetting(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()

	case Cmd_ChannelAdminLogEventActionChangePhoto434bd2af:
		m2 := NewTLChannelAdminLogEventActionChangePhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventAction()

	default:
		return fmt.Errorf("ChannelAdminLogEventAction Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChannelAdminLogEventsFilter) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelAdminLogEventsFilter:
		m2 := NewTLChannelAdminLogEventsFilter(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminLogEventsFilter()

	default:
		return fmt.Errorf("ChannelAdminLogEventsFilter Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChannelAdminRights) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelAdminRights:
		m2 := NewTLChannelAdminRights(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelAdminRights()

	default:
		return fmt.Errorf("ChannelAdminRights Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChannelBannedRights) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelBannedRights:
		m2 := NewTLChannelBannedRights(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelBannedRights()

	default:
		return fmt.Errorf("ChannelBannedRights Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChannelLocation) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelLocationEmpty:
		m2 := NewTLChannelLocationEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelLocation()
	case Cmd_ChannelLocation:
		m2 := NewTLChannelLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelLocation()

	default:
		return fmt.Errorf("ChannelLocation Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChannelMessagesFilter) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelMessagesFilterEmpty:
		m2 := NewTLChannelMessagesFilterEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelMessagesFilter()
	case Cmd_ChannelMessagesFilter:
		m2 := NewTLChannelMessagesFilter(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelMessagesFilter()
	case Cmd_ChannelMessagesFilterCollapsed:
		m2 := NewTLChannelMessagesFilterCollapsed(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelMessagesFilter()

	default:
		return fmt.Errorf("ChannelMessagesFilter Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChannelParticipant) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelParticipant:
		m2 := NewTLChannelParticipant(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantSelf:
		m2 := NewTLChannelParticipantSelf(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantCreator:
		m2 := NewTLChannelParticipantCreator(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantAdmin:
		m2 := NewTLChannelParticipantAdmin(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantBanned:
		m2 := NewTLChannelParticipantBanned(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantModerator:
		m2 := NewTLChannelParticipantModerator(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantEditor:
		m2 := NewTLChannelParticipantEditor(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantKicked:
		m2 := NewTLChannelParticipantKicked(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantLeft:
		m2 := NewTLChannelParticipantLeft(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()

	case Cmd_ChannelParticipantAdmin5daa6e23:
		m2 := NewTLChannelParticipantAdmin(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantBanned1c0facaf:
		m2 := NewTLChannelParticipantBanned(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantCreator808d15a4:
		m2 := NewTLChannelParticipantCreator(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantAdminccbebbaf:
		m2 := NewTLChannelParticipantAdmin(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()
	case Cmd_ChannelParticipantCreator447dca4b:
		m2 := NewTLChannelParticipantCreator(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipant()

	default:
		return fmt.Errorf("ChannelParticipant Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChannelParticipantRole) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelRoleEmpty:
		m2 := NewTLChannelRoleEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantRole()
	case Cmd_ChannelRoleModerator:
		m2 := NewTLChannelRoleModerator(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantRole()
	case Cmd_ChannelRoleEditor:
		m2 := NewTLChannelRoleEditor(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantRole()

	default:
		return fmt.Errorf("ChannelParticipantRole Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChannelParticipantsFilter) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelParticipantsRecent:
		m2 := NewTLChannelParticipantsRecent(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantsFilter()
	case Cmd_ChannelParticipantsAdmins:
		m2 := NewTLChannelParticipantsAdmins(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantsFilter()
	case Cmd_ChannelParticipantsKicked:
		m2 := NewTLChannelParticipantsKicked(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantsFilter()
	case Cmd_ChannelParticipantsBots:
		m2 := NewTLChannelParticipantsBots(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantsFilter()
	case Cmd_ChannelParticipantsBanned:
		m2 := NewTLChannelParticipantsBanned(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantsFilter()
	case Cmd_ChannelParticipantsSearch:
		m2 := NewTLChannelParticipantsSearch(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantsFilter()
	case Cmd_ChannelParticipantsContacts:
		m2 := NewTLChannelParticipantsContacts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantsFilter()
	case Cmd_ChannelParticipantsMentions:
		m2 := NewTLChannelParticipantsMentions(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantsFilter()

	case Cmd_ChannelParticipantsKicked3c37bb7a:
		m2 := NewTLChannelParticipantsKicked(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChannelParticipantsFilter()

	default:
		return fmt.Errorf("ChannelParticipantsFilter Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Channels_AdminLogResults) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelsAdminLogResults:
		m2 := NewTLChannelsAdminLogResults(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Channels_AdminLogResults()

	default:
		return fmt.Errorf("Channels_AdminLogResults Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Channels_ChannelParticipant) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelsChannelParticipant:
		m2 := NewTLChannelsChannelParticipant(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Channels_ChannelParticipant()

	default:
		return fmt.Errorf("Channels_ChannelParticipant Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Channels_ChannelParticipants) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChannelsChannelParticipants:
		m2 := NewTLChannelsChannelParticipants(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Channels_ChannelParticipants()
	case Cmd_ChannelsChannelParticipantsNotModified:
		m2 := NewTLChannelsChannelParticipantsNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Channels_ChannelParticipants()

	default:
		return fmt.Errorf("Channels_ChannelParticipants Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Chat) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChatEmpty:
		m2 := NewTLChatEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Chat()
	case Cmd_Chat:
		m2 := NewTLChat(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Chat()
	case Cmd_ChatForbidden:
		m2 := NewTLChatForbidden(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Chat()
	case Cmd_Channel:
		m2 := NewTLChannel(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Chat()
	case Cmd_ChannelForbidden:
		m2 := NewTLChannelForbidden(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Chat()

	case Cmd_Channela14dca52:
		m2 := NewTLChannel(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Chat()
	case Cmd_ChannelForbidden2d85832c:
		m2 := NewTLChannelForbidden(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Chat()
	case Cmd_Channelc88974ac:
		m2 := NewTLChannel(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Chat()
	case Cmd_Channel4df30834:
		m2 := NewTLChannel(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Chat()
	case Cmd_Chat3bda1bde:
		m2 := NewTLChat(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Chat()
	case Cmd_Channeld31a961e:
		m2 := NewTLChannel(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Chat()

	default:
		return fmt.Errorf("Chat Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChatAdminRights) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChatAdminRights:
		m2 := NewTLChatAdminRights(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatAdminRights()

	default:
		return fmt.Errorf("ChatAdminRights Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChatBannedRights) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChatBannedRights:
		m2 := NewTLChatBannedRights(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatBannedRights()

	default:
		return fmt.Errorf("ChatBannedRights Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChatFull) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChatFull:
		m2 := NewTLChatFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChannelFull:
		m2 := NewTLChannelFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()

	case Cmd_ChannelFull97bee562:
		m2 := NewTLChannelFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChannelFull76af5481:
		m2 := NewTLChannelFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChatFulledd2a791:
		m2 := NewTLChatFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChannelFull1c87a71a:
		m2 := NewTLChannelFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChatFull22a235da:
		m2 := NewTLChatFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChatFull1b7c9db3:
		m2 := NewTLChatFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChannelFull3648977:
		m2 := NewTLChannelFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChannelFull9882e516:
		m2 := NewTLChannelFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChannelFull10916653:
		m2 := NewTLChannelFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChannelFull2d895c74:
		m2 := NewTLChannelFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChannelFullf0e6672a:
		m2 := NewTLChannelFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChatFulldc8c181:
		m2 := NewTLChatFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()
	case Cmd_ChannelFullef3a6acd:
		m2 := NewTLChannelFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatFull()

	default:
		return fmt.Errorf("ChatFull Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChatInvite) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChatInviteAlready:
		m2 := NewTLChatInviteAlready(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatInvite()
	case Cmd_ChatInvite:
		m2 := NewTLChatInvite(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatInvite()
	case Cmd_ChatInvitePeek:
		m2 := NewTLChatInvitePeek(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatInvite()

	case Cmd_ChatInvite93e99b60:
		m2 := NewTLChatInvite(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatInvite()
	case Cmd_ChatInvitedfc2f58e:
		m2 := NewTLChatInvite(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatInvite()

	default:
		return fmt.Errorf("ChatInvite Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChatOnlines) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChatOnlines:
		m2 := NewTLChatOnlines(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatOnlines()

	default:
		return fmt.Errorf("ChatOnlines Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChatParticipant) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChatParticipant:
		m2 := NewTLChatParticipant(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatParticipant()
	case Cmd_ChatParticipantCreator:
		m2 := NewTLChatParticipantCreator(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatParticipant()
	case Cmd_ChatParticipantAdmin:
		m2 := NewTLChatParticipantAdmin(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatParticipant()

	default:
		return fmt.Errorf("ChatParticipant Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChatParticipants) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChatParticipantsForbidden:
		m2 := NewTLChatParticipantsForbidden(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatParticipants()
	case Cmd_ChatParticipants:
		m2 := NewTLChatParticipants(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatParticipants()

	default:
		return fmt.Errorf("ChatParticipants Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ChatPhoto) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChatPhotoEmpty:
		m2 := NewTLChatPhotoEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatPhoto()
	case Cmd_ChatPhoto:
		m2 := NewTLChatPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatPhoto()

	case Cmd_ChatPhoto475cdbd5:
		m2 := NewTLChatPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatPhoto()
	case Cmd_ChatPhotod20b9f3c:
		m2 := NewTLChatPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ChatPhoto()

	default:
		return fmt.Errorf("ChatPhoto Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Client_DH_Inner_Data) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Client_DHInnerData:
		m2 := NewTLClient_DHInnerData(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Client_DH_Inner_Data()

	default:
		return fmt.Errorf("Client_DH_Inner_Data Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *CodeSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_CodeSettings:
		m2 := NewTLCodeSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_CodeSettings()

	case Cmd_CodeSettingsdebebe83:
		m2 := NewTLCodeSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_CodeSettings()

	default:
		return fmt.Errorf("CodeSettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Config) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Config:
		m2 := NewTLConfig(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Config()

	case Cmd_Config317ceef4:
		m2 := NewTLConfig(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Config()
	case Cmd_Config3213dbba:
		m2 := NewTLConfig(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Config()
	case Cmd_Confige6ca25f6:
		m2 := NewTLConfig(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Config()
	case Cmd_Config330b4067:
		m2 := NewTLConfig(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Config()

	default:
		return fmt.Errorf("Config Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Contact) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Contact:
		m2 := NewTLContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contact()

	default:
		return fmt.Errorf("Contact Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ContactBlocked) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ContactBlocked:
		m2 := NewTLContactBlocked(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ContactBlocked()

	default:
		return fmt.Errorf("ContactBlocked Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ContactLink) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ContactLinkUnknown:
		m2 := NewTLContactLinkUnknown(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ContactLink()
	case Cmd_ContactLinkNone:
		m2 := NewTLContactLinkNone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ContactLink()
	case Cmd_ContactLinkHasPhone:
		m2 := NewTLContactLinkHasPhone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ContactLink()
	case Cmd_ContactLinkContact:
		m2 := NewTLContactLinkContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ContactLink()

	default:
		return fmt.Errorf("ContactLink Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ContactStatus) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ContactStatus:
		m2 := NewTLContactStatus(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ContactStatus()

	default:
		return fmt.Errorf("ContactStatus Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Contacts_Blocked) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ContactsBlocked:
		m2 := NewTLContactsBlocked(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_Blocked()
	case Cmd_ContactsBlockedSlice:
		m2 := NewTLContactsBlockedSlice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_Blocked()

	case Cmd_ContactsBlockedade1591:
		m2 := NewTLContactsBlocked(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_Blocked()
	case Cmd_ContactsBlockedSlicee1664194:
		m2 := NewTLContactsBlockedSlice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_Blocked()

	default:
		return fmt.Errorf("Contacts_Blocked Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Contacts_Contacts) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ContactsContactsNotModified:
		m2 := NewTLContactsContactsNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_Contacts()
	case Cmd_ContactsContacts:
		m2 := NewTLContactsContacts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_Contacts()

	case Cmd_ContactsContacts6f8b8cb2:
		m2 := NewTLContactsContacts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_Contacts()

	default:
		return fmt.Errorf("Contacts_Contacts Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Contacts_Found) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ContactsFound:
		m2 := NewTLContactsFound(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_Found()

	case Cmd_ContactsFoundb3134d9d:
		m2 := NewTLContactsFound(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_Found()

	default:
		return fmt.Errorf("Contacts_Found Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Contacts_ImportedContacts) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ContactsImportedContacts:
		m2 := NewTLContactsImportedContacts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_ImportedContacts()

	case Cmd_ContactsImportedContactsad524315:
		m2 := NewTLContactsImportedContacts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_ImportedContacts()

	default:
		return fmt.Errorf("Contacts_ImportedContacts Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Contacts_Link) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ContactsLink:
		m2 := NewTLContactsLink(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_Link()

	default:
		return fmt.Errorf("Contacts_Link Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Contacts_ResolvedPeer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ContactsResolvedPeer:
		m2 := NewTLContactsResolvedPeer(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_ResolvedPeer()

	default:
		return fmt.Errorf("Contacts_ResolvedPeer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Contacts_TopPeers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ContactsTopPeersNotModified:
		m2 := NewTLContactsTopPeersNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_TopPeers()
	case Cmd_ContactsTopPeers:
		m2 := NewTLContactsTopPeers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_TopPeers()
	case Cmd_ContactsTopPeersDisabled:
		m2 := NewTLContactsTopPeersDisabled(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Contacts_TopPeers()

	default:
		return fmt.Errorf("Contacts_TopPeers Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *DataJSON) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DataJSON:
		m2 := NewTLDataJSON(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DataJSON()

	default:
		return fmt.Errorf("DataJSON Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *DcOption) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DcOption:
		m2 := NewTLDcOption(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DcOption()

	case Cmd_DcOption18b7a10d:
		m2 := NewTLDcOption(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DcOption()

	default:
		return fmt.Errorf("DcOption Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *DestroyAuthKeyRes) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DestroyAuthKeyOk:
		m2 := NewTLDestroyAuthKeyOk(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DestroyAuthKeyRes()
	case Cmd_DestroyAuthKeyNone:
		m2 := NewTLDestroyAuthKeyNone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DestroyAuthKeyRes()
	case Cmd_DestroyAuthKeyFail:
		m2 := NewTLDestroyAuthKeyFail(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DestroyAuthKeyRes()

	default:
		return fmt.Errorf("DestroyAuthKeyRes Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *DestroySessionRes) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DestroySessionOk:
		m2 := NewTLDestroySessionOk(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DestroySessionRes()
	case Cmd_DestroySessionNone:
		m2 := NewTLDestroySessionNone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DestroySessionRes()

	default:
		return fmt.Errorf("DestroySessionRes Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Dialog) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Dialog:
		m2 := NewTLDialog(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Dialog()
	case Cmd_DialogChannel:
		m2 := NewTLDialogChannel(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Dialog()
	case Cmd_DialogFolder:
		m2 := NewTLDialogFolder(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Dialog()

	case Cmd_Dialogc1dd804a:
		m2 := NewTLDialog(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Dialog()
	case Cmd_Dialog2c171f72:
		m2 := NewTLDialog(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Dialog()

	default:
		return fmt.Errorf("Dialog Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *DialogFilter) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DialogFilter:
		m2 := NewTLDialogFilter(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DialogFilter()

	default:
		return fmt.Errorf("DialogFilter Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *DialogFilterSuggested) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DialogFilterSuggested:
		m2 := NewTLDialogFilterSuggested(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DialogFilterSuggested()

	default:
		return fmt.Errorf("DialogFilterSuggested Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *DialogPeer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DialogPeer:
		m2 := NewTLDialogPeer(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DialogPeer()
	case Cmd_DialogPeerFolder:
		m2 := NewTLDialogPeerFolder(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DialogPeer()

	default:
		return fmt.Errorf("DialogPeer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *DisabledFeature) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DisabledFeature:
		m2 := NewTLDisabledFeature(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DisabledFeature()

	default:
		return fmt.Errorf("DisabledFeature Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Document) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DocumentEmpty:
		m2 := NewTLDocumentEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Document()
	case Cmd_Document:
		m2 := NewTLDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Document()

	case Cmd_Documentf9a39f4f:
		m2 := NewTLDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Document()
	case Cmd_Document59534e4c:
		m2 := NewTLDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Document()
	case Cmd_Document9ba29cc1:
		m2 := NewTLDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Document()
	case Cmd_Document1e87342b:
		m2 := NewTLDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Document()

	default:
		return fmt.Errorf("Document Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *DocumentAttribute) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DocumentAttributeImageSize:
		m2 := NewTLDocumentAttributeImageSize(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DocumentAttribute()
	case Cmd_DocumentAttributeAnimated:
		m2 := NewTLDocumentAttributeAnimated(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DocumentAttribute()
	case Cmd_DocumentAttributeSticker:
		m2 := NewTLDocumentAttributeSticker(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DocumentAttribute()
	case Cmd_DocumentAttributeVideo:
		m2 := NewTLDocumentAttributeVideo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DocumentAttribute()
	case Cmd_DocumentAttributeAudio:
		m2 := NewTLDocumentAttributeAudio(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DocumentAttribute()
	case Cmd_DocumentAttributeFilename:
		m2 := NewTLDocumentAttributeFilename(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DocumentAttribute()
	case Cmd_DocumentAttributeHasStickers:
		m2 := NewTLDocumentAttributeHasStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DocumentAttribute()

	case Cmd_DocumentAttributeSticker3a556302:
		m2 := NewTLDocumentAttributeSticker(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DocumentAttribute()
	case Cmd_DocumentAttributeVideo5910cccb:
		m2 := NewTLDocumentAttributeVideo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DocumentAttribute()

	default:
		return fmt.Errorf("DocumentAttribute Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *DraftMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DraftMessageEmpty:
		m2 := NewTLDraftMessageEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DraftMessage()
	case Cmd_DraftMessage:
		m2 := NewTLDraftMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DraftMessage()

	case Cmd_DraftMessageEmpty1b0c841a:
		m2 := NewTLDraftMessageEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_DraftMessage()

	default:
		return fmt.Errorf("DraftMessage Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *EmojiKeyword) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_EmojiKeyword:
		m2 := NewTLEmojiKeyword(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EmojiKeyword()
	case Cmd_EmojiKeywordDeleted:
		m2 := NewTLEmojiKeywordDeleted(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EmojiKeyword()

	default:
		return fmt.Errorf("EmojiKeyword Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *EmojiKeywordsDifference) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_EmojiKeywordsDifference:
		m2 := NewTLEmojiKeywordsDifference(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EmojiKeywordsDifference()

	default:
		return fmt.Errorf("EmojiKeywordsDifference Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *EmojiLanguage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_EmojiLanguage:
		m2 := NewTLEmojiLanguage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EmojiLanguage()

	default:
		return fmt.Errorf("EmojiLanguage Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *EmojiURL) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_EmojiURL:
		m2 := NewTLEmojiURL(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EmojiURL()

	default:
		return fmt.Errorf("EmojiURL Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *EncryptedChat) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_EncryptedChatEmpty:
		m2 := NewTLEncryptedChatEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EncryptedChat()
	case Cmd_EncryptedChatWaiting:
		m2 := NewTLEncryptedChatWaiting(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EncryptedChat()
	case Cmd_EncryptedChatRequested:
		m2 := NewTLEncryptedChatRequested(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EncryptedChat()
	case Cmd_EncryptedChat:
		m2 := NewTLEncryptedChat(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EncryptedChat()
	case Cmd_EncryptedChatDiscarded:
		m2 := NewTLEncryptedChatDiscarded(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EncryptedChat()

	case Cmd_EncryptedChatRequested62718a82:
		m2 := NewTLEncryptedChatRequested(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EncryptedChat()

	default:
		return fmt.Errorf("EncryptedChat Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *EncryptedFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_EncryptedFileEmpty:
		m2 := NewTLEncryptedFileEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EncryptedFile()
	case Cmd_EncryptedFile:
		m2 := NewTLEncryptedFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EncryptedFile()

	default:
		return fmt.Errorf("EncryptedFile Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *EncryptedMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_EncryptedMessage:
		m2 := NewTLEncryptedMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EncryptedMessage()
	case Cmd_EncryptedMessageService:
		m2 := NewTLEncryptedMessageService(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_EncryptedMessage()

	default:
		return fmt.Errorf("EncryptedMessage Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Error) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Error:
		m2 := NewTLError(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Error()

	default:
		return fmt.Errorf("Error Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ExportedChatInvite) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ChatInviteEmpty:
		m2 := NewTLChatInviteEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ExportedChatInvite()
	case Cmd_ChatInviteExported:
		m2 := NewTLChatInviteExported(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ExportedChatInvite()

	default:
		return fmt.Errorf("ExportedChatInvite Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ExportedMessageLink) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ExportedMessageLink:
		m2 := NewTLExportedMessageLink(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ExportedMessageLink()

	case Cmd_ExportedMessageLink5dab1af4:
		m2 := NewTLExportedMessageLink(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ExportedMessageLink()

	default:
		return fmt.Errorf("ExportedMessageLink Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *FileHash) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_FileHash:
		m2 := NewTLFileHash(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_FileHash()

	default:
		return fmt.Errorf("FileHash Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *FileLocation) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_FileLocationUnavailable:
		m2 := NewTLFileLocationUnavailable(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_FileLocation()
	case Cmd_FileLocation:
		m2 := NewTLFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_FileLocation()
	case Cmd_FileLocationToBeDeprecated:
		m2 := NewTLFileLocationToBeDeprecated(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_FileLocation()

	case Cmd_FileLocation91d11eb:
		m2 := NewTLFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_FileLocation()

	default:
		return fmt.Errorf("FileLocation Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Folder) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Folder:
		m2 := NewTLFolder(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Folder()

	default:
		return fmt.Errorf("Folder Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *FolderPeer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_FolderPeer:
		m2 := NewTLFolderPeer(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_FolderPeer()

	default:
		return fmt.Errorf("FolderPeer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *FoundGif) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_FoundGif:
		m2 := NewTLFoundGif(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_FoundGif()
	case Cmd_FoundGifCached:
		m2 := NewTLFoundGifCached(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_FoundGif()

	default:
		return fmt.Errorf("FoundGif Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *FutureSalt) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_FutureSalt:
		m2 := NewTLFutureSalt(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_FutureSalt()

	default:
		return fmt.Errorf("FutureSalt Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *FutureSalts) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_FutureSalts:
		m2 := NewTLFutureSalts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_FutureSalts()

	default:
		return fmt.Errorf("FutureSalts Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Game) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Game:
		m2 := NewTLGame(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Game()

	default:
		return fmt.Errorf("Game Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *GeoPoint) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_GeoPointEmpty:
		m2 := NewTLGeoPointEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_GeoPoint()
	case Cmd_GeoPoint:
		m2 := NewTLGeoPoint(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_GeoPoint()

	case Cmd_GeoPoint296f104:
		m2 := NewTLGeoPoint(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_GeoPoint()
	case Cmd_GeoPointb2a2f663:
		m2 := NewTLGeoPoint(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_GeoPoint()

	default:
		return fmt.Errorf("GeoPoint Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *GlobalPrivacySettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_GlobalPrivacySettings:
		m2 := NewTLGlobalPrivacySettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_GlobalPrivacySettings()

	default:
		return fmt.Errorf("GlobalPrivacySettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *GroupCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_GroupCallDiscarded:
		m2 := NewTLGroupCallDiscarded(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_GroupCall()
	case Cmd_GroupCall:
		m2 := NewTLGroupCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_GroupCall()

	default:
		return fmt.Errorf("GroupCall Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *GroupCallParticipant) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_GroupCallParticipant:
		m2 := NewTLGroupCallParticipant(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_GroupCallParticipant()

	default:
		return fmt.Errorf("GroupCallParticipant Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_AppChangelog) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpAppChangelogEmpty:
		m2 := NewTLHelpAppChangelogEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_AppChangelog()
	case Cmd_HelpAppChangelog:
		m2 := NewTLHelpAppChangelog(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_AppChangelog()

	default:
		return fmt.Errorf("Help_AppChangelog Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_AppUpdate) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpAppUpdate:
		m2 := NewTLHelpAppUpdate(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_AppUpdate()
	case Cmd_HelpNoAppUpdate:
		m2 := NewTLHelpNoAppUpdate(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_AppUpdate()

	case Cmd_HelpAppUpdate1da7158f:
		m2 := NewTLHelpAppUpdate(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_AppUpdate()

	default:
		return fmt.Errorf("Help_AppUpdate Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_ConfigSimple) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpConfigSimple:
		m2 := NewTLHelpConfigSimple(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_ConfigSimple()

	case Cmd_HelpConfigSimple5a592a6c:
		m2 := NewTLHelpConfigSimple(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_ConfigSimple()

	default:
		return fmt.Errorf("Help_ConfigSimple Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_CountriesList) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpCountriesListNotModified:
		m2 := NewTLHelpCountriesListNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_CountriesList()
	case Cmd_HelpCountriesList:
		m2 := NewTLHelpCountriesList(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_CountriesList()

	default:
		return fmt.Errorf("Help_CountriesList Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_Country) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpCountry:
		m2 := NewTLHelpCountry(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_Country()

	default:
		return fmt.Errorf("Help_Country Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_CountryCode) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpCountryCode:
		m2 := NewTLHelpCountryCode(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_CountryCode()

	default:
		return fmt.Errorf("Help_CountryCode Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_DeepLinkInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpDeepLinkInfoEmpty:
		m2 := NewTLHelpDeepLinkInfoEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_DeepLinkInfo()
	case Cmd_HelpDeepLinkInfo:
		m2 := NewTLHelpDeepLinkInfo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_DeepLinkInfo()

	default:
		return fmt.Errorf("Help_DeepLinkInfo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_InviteText) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpInviteText:
		m2 := NewTLHelpInviteText(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_InviteText()

	default:
		return fmt.Errorf("Help_InviteText Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_PassportConfig) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpPassportConfigNotModified:
		m2 := NewTLHelpPassportConfigNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_PassportConfig()
	case Cmd_HelpPassportConfig:
		m2 := NewTLHelpPassportConfig(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_PassportConfig()

	default:
		return fmt.Errorf("Help_PassportConfig Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_PromoData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpPromoDataEmpty:
		m2 := NewTLHelpPromoDataEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_PromoData()
	case Cmd_HelpPromoData:
		m2 := NewTLHelpPromoData(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_PromoData()

	case Cmd_HelpPromoData8c39793f:
		m2 := NewTLHelpPromoData(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_PromoData()

	default:
		return fmt.Errorf("Help_PromoData Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_ProxyData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpProxyDataEmpty:
		m2 := NewTLHelpProxyDataEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_ProxyData()
	case Cmd_HelpProxyDataPromo:
		m2 := NewTLHelpProxyDataPromo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_ProxyData()

	default:
		return fmt.Errorf("Help_ProxyData Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_RecentMeUrls) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpRecentMeUrls:
		m2 := NewTLHelpRecentMeUrls(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_RecentMeUrls()

	default:
		return fmt.Errorf("Help_RecentMeUrls Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_Support) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpSupport:
		m2 := NewTLHelpSupport(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_Support()

	default:
		return fmt.Errorf("Help_Support Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_SupportName) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpSupportName:
		m2 := NewTLHelpSupportName(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_SupportName()

	default:
		return fmt.Errorf("Help_SupportName Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_TermsOfService) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpTermsOfService:
		m2 := NewTLHelpTermsOfService(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_TermsOfService()

	case Cmd_HelpTermsOfService780a0310:
		m2 := NewTLHelpTermsOfService(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_TermsOfService()

	default:
		return fmt.Errorf("Help_TermsOfService Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_TermsOfServiceUpdate) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpTermsOfServiceUpdateEmpty:
		m2 := NewTLHelpTermsOfServiceUpdateEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_TermsOfServiceUpdate()
	case Cmd_HelpTermsOfServiceUpdate:
		m2 := NewTLHelpTermsOfServiceUpdate(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_TermsOfServiceUpdate()

	default:
		return fmt.Errorf("Help_TermsOfServiceUpdate Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Help_UserInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HelpUserInfoEmpty:
		m2 := NewTLHelpUserInfoEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_UserInfo()
	case Cmd_HelpUserInfo:
		m2 := NewTLHelpUserInfo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Help_UserInfo()

	default:
		return fmt.Errorf("Help_UserInfo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *HighScore) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HighScore:
		m2 := NewTLHighScore(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_HighScore()

	default:
		return fmt.Errorf("HighScore Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *HttpWait) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_HttpWait:
		m2 := NewTLHttpWait(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_HttpWait()

	default:
		return fmt.Errorf("HttpWait Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ImportedContact) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ImportedContact:
		m2 := NewTLImportedContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ImportedContact()

	default:
		return fmt.Errorf("ImportedContact Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InlineBotSwitchPM) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InlineBotSwitchPM:
		m2 := NewTLInlineBotSwitchPM(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InlineBotSwitchPM()

	default:
		return fmt.Errorf("InlineBotSwitchPM Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InlineQueryPeerType) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InlineQueryPeerTypeSameBotPM:
		m2 := NewTLInlineQueryPeerTypeSameBotPM(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InlineQueryPeerType()
	case Cmd_InlineQueryPeerTypePM:
		m2 := NewTLInlineQueryPeerTypePM(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InlineQueryPeerType()
	case Cmd_InlineQueryPeerTypeChat:
		m2 := NewTLInlineQueryPeerTypeChat(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InlineQueryPeerType()
	case Cmd_InlineQueryPeerTypeMegagroup:
		m2 := NewTLInlineQueryPeerTypeMegagroup(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InlineQueryPeerType()
	case Cmd_InlineQueryPeerTypeBroadcast:
		m2 := NewTLInlineQueryPeerTypeBroadcast(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InlineQueryPeerType()

	default:
		return fmt.Errorf("InlineQueryPeerType Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputAppEvent) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputAppEvent:
		m2 := NewTLInputAppEvent(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputAppEvent()

	case Cmd_InputAppEvent1d1b1245:
		m2 := NewTLInputAppEvent(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputAppEvent()

	default:
		return fmt.Errorf("InputAppEvent Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputBotInlineMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputBotInlineMessageMediaAuto:
		m2 := NewTLInputBotInlineMessageMediaAuto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessage()
	case Cmd_InputBotInlineMessageText:
		m2 := NewTLInputBotInlineMessageText(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessage()
	case Cmd_InputBotInlineMessageMediaGeo:
		m2 := NewTLInputBotInlineMessageMediaGeo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessage()
	case Cmd_InputBotInlineMessageMediaVenue:
		m2 := NewTLInputBotInlineMessageMediaVenue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessage()
	case Cmd_InputBotInlineMessageMediaContact:
		m2 := NewTLInputBotInlineMessageMediaContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessage()
	case Cmd_InputBotInlineMessageGame:
		m2 := NewTLInputBotInlineMessageGame(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessage()

	case Cmd_InputBotInlineMessageMediaAuto3380c786:
		m2 := NewTLInputBotInlineMessageMediaAuto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessage()
	case Cmd_InputBotInlineMessageMediaGeoc1b15d65:
		m2 := NewTLInputBotInlineMessageMediaGeo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessage()
	case Cmd_InputBotInlineMessageMediaVenue417bbf11:
		m2 := NewTLInputBotInlineMessageMediaVenue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessage()
	case Cmd_InputBotInlineMessageMediaContacta6edbffd:
		m2 := NewTLInputBotInlineMessageMediaContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessage()
	case Cmd_InputBotInlineMessageMediaGeo96929a85:
		m2 := NewTLInputBotInlineMessageMediaGeo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessage()

	default:
		return fmt.Errorf("InputBotInlineMessage Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputBotInlineMessageID) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputBotInlineMessageID:
		m2 := NewTLInputBotInlineMessageID(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineMessageID()

	default:
		return fmt.Errorf("InputBotInlineMessageID Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputBotInlineResult) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputBotInlineResult:
		m2 := NewTLInputBotInlineResult(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineResult()
	case Cmd_InputBotInlineResultPhoto:
		m2 := NewTLInputBotInlineResultPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineResult()
	case Cmd_InputBotInlineResultDocument:
		m2 := NewTLInputBotInlineResultDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineResult()
	case Cmd_InputBotInlineResultGame:
		m2 := NewTLInputBotInlineResultGame(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineResult()

	case Cmd_InputBotInlineResult88bf9319:
		m2 := NewTLInputBotInlineResult(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputBotInlineResult()

	default:
		return fmt.Errorf("InputBotInlineResult Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputChannel) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputChannelEmpty:
		m2 := NewTLInputChannelEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputChannel()
	case Cmd_InputChannel:
		m2 := NewTLInputChannel(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputChannel()
	case Cmd_InputChannelFromMessage:
		m2 := NewTLInputChannelFromMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputChannel()

	default:
		return fmt.Errorf("InputChannel Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputChatPhoto) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputChatPhotoEmpty:
		m2 := NewTLInputChatPhotoEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputChatPhoto()
	case Cmd_InputChatUploadedPhoto:
		m2 := NewTLInputChatUploadedPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputChatPhoto()
	case Cmd_InputChatPhoto:
		m2 := NewTLInputChatPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputChatPhoto()

	case Cmd_InputChatUploadedPhoto94254732:
		m2 := NewTLInputChatUploadedPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputChatPhoto()
	case Cmd_InputChatPhotob2e1bf08:
		m2 := NewTLInputChatPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputChatPhoto()
	case Cmd_InputChatUploadedPhotoc642724e:
		m2 := NewTLInputChatUploadedPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputChatPhoto()

	default:
		return fmt.Errorf("InputChatPhoto Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputCheckPasswordSRP) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputCheckPasswordEmpty:
		m2 := NewTLInputCheckPasswordEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputCheckPasswordSRP()
	case Cmd_InputCheckPasswordSRP:
		m2 := NewTLInputCheckPasswordSRP(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputCheckPasswordSRP()

	default:
		return fmt.Errorf("InputCheckPasswordSRP Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputClientProxy) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputClientProxy:
		m2 := NewTLInputClientProxy(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputClientProxy()

	default:
		return fmt.Errorf("InputClientProxy Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputContact) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputPhoneContact:
		m2 := NewTLInputPhoneContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputContact()

	default:
		return fmt.Errorf("InputContact Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputDialogPeer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputDialogPeer:
		m2 := NewTLInputDialogPeer(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputDialogPeer()
	case Cmd_InputDialogPeerFolder:
		m2 := NewTLInputDialogPeerFolder(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputDialogPeer()

	default:
		return fmt.Errorf("InputDialogPeer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputDocument) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputDocumentEmpty:
		m2 := NewTLInputDocumentEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputDocument()
	case Cmd_InputDocument:
		m2 := NewTLInputDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputDocument()

	case Cmd_InputDocument1abfb575:
		m2 := NewTLInputDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputDocument()

	default:
		return fmt.Errorf("InputDocument Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputEncryptedChat) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputEncryptedChat:
		m2 := NewTLInputEncryptedChat(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputEncryptedChat()

	default:
		return fmt.Errorf("InputEncryptedChat Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputEncryptedFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputEncryptedFileEmpty:
		m2 := NewTLInputEncryptedFileEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputEncryptedFile()
	case Cmd_InputEncryptedFileUploaded:
		m2 := NewTLInputEncryptedFileUploaded(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputEncryptedFile()
	case Cmd_InputEncryptedFile:
		m2 := NewTLInputEncryptedFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputEncryptedFile()
	case Cmd_InputEncryptedFileBigUploaded:
		m2 := NewTLInputEncryptedFileBigUploaded(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputEncryptedFile()

	default:
		return fmt.Errorf("InputEncryptedFile Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputFile:
		m2 := NewTLInputFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFile()
	case Cmd_InputFileBig:
		m2 := NewTLInputFileBig(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFile()

	default:
		return fmt.Errorf("InputFile Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputFileLocation) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputFileLocation:
		m2 := NewTLInputFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()
	case Cmd_InputEncryptedFileLocation:
		m2 := NewTLInputEncryptedFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()
	case Cmd_InputDocumentFileLocation:
		m2 := NewTLInputDocumentFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()
	case Cmd_InputSecureFileLocation:
		m2 := NewTLInputSecureFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()
	case Cmd_InputTakeoutFileLocation:
		m2 := NewTLInputTakeoutFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()
	case Cmd_InputPhotoFileLocation:
		m2 := NewTLInputPhotoFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()
	case Cmd_InputPeerPhotoFileLocation:
		m2 := NewTLInputPeerPhotoFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()
	case Cmd_InputStickerSetThumb:
		m2 := NewTLInputStickerSetThumb(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()
	case Cmd_InputPhotoLegacyFileLocation:
		m2 := NewTLInputPhotoLegacyFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()

	case Cmd_InputDocumentFileLocation4e45abe9:
		m2 := NewTLInputDocumentFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()
	case Cmd_InputFileLocationdfdaabe1:
		m2 := NewTLInputFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()
	case Cmd_InputDocumentFileLocation196683d9:
		m2 := NewTLInputDocumentFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()
	case Cmd_InputDocumentFileLocationbad07584:
		m2 := NewTLInputDocumentFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFileLocation()

	default:
		return fmt.Errorf("InputFileLocation Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputFolderPeer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputFolderPeer:
		m2 := NewTLInputFolderPeer(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputFolderPeer()

	default:
		return fmt.Errorf("InputFolderPeer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputGame) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputGameID:
		m2 := NewTLInputGameID(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputGame()
	case Cmd_InputGameShortName:
		m2 := NewTLInputGameShortName(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputGame()

	default:
		return fmt.Errorf("InputGame Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputGeoPoint) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputGeoPointEmpty:
		m2 := NewTLInputGeoPointEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputGeoPoint()
	case Cmd_InputGeoPoint:
		m2 := NewTLInputGeoPoint(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputGeoPoint()

	case Cmd_InputGeoPoint48222faf:
		m2 := NewTLInputGeoPoint(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputGeoPoint()

	default:
		return fmt.Errorf("InputGeoPoint Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputGroupCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputGroupCall:
		m2 := NewTLInputGroupCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputGroupCall()

	default:
		return fmt.Errorf("InputGroupCall Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputMedia) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputMediaEmpty:
		m2 := NewTLInputMediaEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaUploadedPhoto:
		m2 := NewTLInputMediaUploadedPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaPhoto:
		m2 := NewTLInputMediaPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaGeoPoint:
		m2 := NewTLInputMediaGeoPoint(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaContact:
		m2 := NewTLInputMediaContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaUploadedDocument:
		m2 := NewTLInputMediaUploadedDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaDocument:
		m2 := NewTLInputMediaDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaVenue:
		m2 := NewTLInputMediaVenue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaGifExternal:
		m2 := NewTLInputMediaGifExternal(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaPhotoExternal:
		m2 := NewTLInputMediaPhotoExternal(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaDocumentExternal:
		m2 := NewTLInputMediaDocumentExternal(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaGame:
		m2 := NewTLInputMediaGame(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaInvoice:
		m2 := NewTLInputMediaInvoice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaUploadedThumbDocument:
		m2 := NewTLInputMediaUploadedThumbDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaGeoLive:
		m2 := NewTLInputMediaGeoLive(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaPoll:
		m2 := NewTLInputMediaPoll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaDice:
		m2 := NewTLInputMediaDice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()

	case Cmd_InputMediaPhotoe9bfb4f3:
		m2 := NewTLInputMediaPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaUploadedPhotof7aff1c0:
		m2 := NewTLInputMediaUploadedPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaUploadedDocument1d89306d:
		m2 := NewTLInputMediaUploadedDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaDocument1a77f29c:
		m2 := NewTLInputMediaDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaUploadedPhoto1e287d04:
		m2 := NewTLInputMediaUploadedPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaPhotob3ba0635:
		m2 := NewTLInputMediaPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaContactf8ab7dfb:
		m2 := NewTLInputMediaContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaDocument23ab23d2:
		m2 := NewTLInputMediaDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaVenuec13d1c11:
		m2 := NewTLInputMediaVenue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaPhotoExternale5bbfe1a:
		m2 := NewTLInputMediaPhotoExternal(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaDocumentExternalfb52dc99:
		m2 := NewTLInputMediaDocumentExternal(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaInvoicef4e096c3:
		m2 := NewTLInputMediaInvoice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaGeoLivece4e82fd:
		m2 := NewTLInputMediaGeoLive(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaPollabe9ca25:
		m2 := NewTLInputMediaPoll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaPollf94e5f1:
		m2 := NewTLInputMediaPoll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaDicee66fbf7b:
		m2 := NewTLInputMediaDice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaUploadedDocument5b38c6c1:
		m2 := NewTLInputMediaUploadedDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaGeoLive971fa843:
		m2 := NewTLInputMediaGeoLive(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()
	case Cmd_InputMediaDocument33473058:
		m2 := NewTLInputMediaDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMedia()

	default:
		return fmt.Errorf("InputMedia Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputMessageID:
		m2 := NewTLInputMessageID(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMessage()
	case Cmd_InputMessageReplyTo:
		m2 := NewTLInputMessageReplyTo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMessage()
	case Cmd_InputMessagePinned:
		m2 := NewTLInputMessagePinned(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMessage()
	case Cmd_InputMessageCallbackQuery:
		m2 := NewTLInputMessageCallbackQuery(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputMessage()

	default:
		return fmt.Errorf("InputMessage Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputNotifyPeer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputNotifyPeer:
		m2 := NewTLInputNotifyPeer(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputNotifyPeer()
	case Cmd_InputNotifyUsers:
		m2 := NewTLInputNotifyUsers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputNotifyPeer()
	case Cmd_InputNotifyChats:
		m2 := NewTLInputNotifyChats(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputNotifyPeer()
	case Cmd_InputNotifyAll:
		m2 := NewTLInputNotifyAll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputNotifyPeer()
	case Cmd_InputNotifyBroadcasts:
		m2 := NewTLInputNotifyBroadcasts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputNotifyPeer()

	default:
		return fmt.Errorf("InputNotifyPeer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputPaymentCredentials) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputPaymentCredentialsSaved:
		m2 := NewTLInputPaymentCredentialsSaved(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPaymentCredentials()
	case Cmd_InputPaymentCredentials:
		m2 := NewTLInputPaymentCredentials(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPaymentCredentials()
	case Cmd_InputPaymentCredentialsApplePay:
		m2 := NewTLInputPaymentCredentialsApplePay(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPaymentCredentials()
	case Cmd_InputPaymentCredentialsAndroidPay:
		m2 := NewTLInputPaymentCredentialsAndroidPay(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPaymentCredentials()

	default:
		return fmt.Errorf("InputPaymentCredentials Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputPeer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputPeerEmpty:
		m2 := NewTLInputPeerEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPeer()
	case Cmd_InputPeerSelf:
		m2 := NewTLInputPeerSelf(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPeer()
	case Cmd_InputPeerChat:
		m2 := NewTLInputPeerChat(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPeer()
	case Cmd_InputPeerUser:
		m2 := NewTLInputPeerUser(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPeer()
	case Cmd_InputPeerChannel:
		m2 := NewTLInputPeerChannel(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPeer()
	case Cmd_InputPeerUserFromMessage:
		m2 := NewTLInputPeerUserFromMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPeer()
	case Cmd_InputPeerChannelFromMessage:
		m2 := NewTLInputPeerChannelFromMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPeer()

	default:
		return fmt.Errorf("InputPeer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputPeerNotifyEvents) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputPeerNotifyEventsEmpty:
		m2 := NewTLInputPeerNotifyEventsEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPeerNotifyEvents()
	case Cmd_InputPeerNotifyEventsAll:
		m2 := NewTLInputPeerNotifyEventsAll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPeerNotifyEvents()

	default:
		return fmt.Errorf("InputPeerNotifyEvents Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputPeerNotifySettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputPeerNotifySettings:
		m2 := NewTLInputPeerNotifySettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPeerNotifySettings()

	case Cmd_InputPeerNotifySettings9c3d198e:
		m2 := NewTLInputPeerNotifySettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPeerNotifySettings()

	default:
		return fmt.Errorf("InputPeerNotifySettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputPhoneCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputPhoneCall:
		m2 := NewTLInputPhoneCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPhoneCall()

	default:
		return fmt.Errorf("InputPhoneCall Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputPhoto) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputPhotoEmpty:
		m2 := NewTLInputPhotoEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPhoto()
	case Cmd_InputPhoto:
		m2 := NewTLInputPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPhoto()

	case Cmd_InputPhoto3bb3b94a:
		m2 := NewTLInputPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPhoto()

	default:
		return fmt.Errorf("InputPhoto Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputPhotoCrop) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputPhotoCrop:
		m2 := NewTLInputPhotoCrop(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPhotoCrop()
	case Cmd_InputPhotoCropAuto:
		m2 := NewTLInputPhotoCropAuto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPhotoCrop()

	default:
		return fmt.Errorf("InputPhotoCrop Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputPrivacyKey) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputPrivacyKeyStatusTimestamp:
		m2 := NewTLInputPrivacyKeyStatusTimestamp(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyKey()
	case Cmd_InputPrivacyKeyChatInvite:
		m2 := NewTLInputPrivacyKeyChatInvite(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyKey()
	case Cmd_InputPrivacyKeyPhoneCall:
		m2 := NewTLInputPrivacyKeyPhoneCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyKey()
	case Cmd_InputPrivacyKeyPhoneP2P:
		m2 := NewTLInputPrivacyKeyPhoneP2P(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyKey()
	case Cmd_InputPrivacyKeyForwards:
		m2 := NewTLInputPrivacyKeyForwards(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyKey()
	case Cmd_InputPrivacyKeyProfilePhoto:
		m2 := NewTLInputPrivacyKeyProfilePhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyKey()
	case Cmd_InputPrivacyKeyPhoneNumber:
		m2 := NewTLInputPrivacyKeyPhoneNumber(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyKey()
	case Cmd_InputPrivacyKeyAddedByPhone:
		m2 := NewTLInputPrivacyKeyAddedByPhone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyKey()

	default:
		return fmt.Errorf("InputPrivacyKey Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputPrivacyRule) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputPrivacyValueAllowContacts:
		m2 := NewTLInputPrivacyValueAllowContacts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyRule()
	case Cmd_InputPrivacyValueAllowAll:
		m2 := NewTLInputPrivacyValueAllowAll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyRule()
	case Cmd_InputPrivacyValueAllowUsers:
		m2 := NewTLInputPrivacyValueAllowUsers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyRule()
	case Cmd_InputPrivacyValueDisallowContacts:
		m2 := NewTLInputPrivacyValueDisallowContacts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyRule()
	case Cmd_InputPrivacyValueDisallowAll:
		m2 := NewTLInputPrivacyValueDisallowAll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyRule()
	case Cmd_InputPrivacyValueDisallowUsers:
		m2 := NewTLInputPrivacyValueDisallowUsers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyRule()
	case Cmd_InputPrivacyValueAllowChatParticipants:
		m2 := NewTLInputPrivacyValueAllowChatParticipants(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyRule()
	case Cmd_InputPrivacyValueDisallowChatParticipants:
		m2 := NewTLInputPrivacyValueDisallowChatParticipants(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputPrivacyRule()

	default:
		return fmt.Errorf("InputPrivacyRule Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputSecureFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputSecureFileUploaded:
		m2 := NewTLInputSecureFileUploaded(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputSecureFile()
	case Cmd_InputSecureFile:
		m2 := NewTLInputSecureFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputSecureFile()

	default:
		return fmt.Errorf("InputSecureFile Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputSecureValue) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputSecureValue:
		m2 := NewTLInputSecureValue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputSecureValue()

	default:
		return fmt.Errorf("InputSecureValue Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputSingleMedia) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputSingleMedia:
		m2 := NewTLInputSingleMedia(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputSingleMedia()

	default:
		return fmt.Errorf("InputSingleMedia Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputStickerSet) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputStickerSetEmpty:
		m2 := NewTLInputStickerSetEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputStickerSet()
	case Cmd_InputStickerSetID:
		m2 := NewTLInputStickerSetID(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputStickerSet()
	case Cmd_InputStickerSetShortName:
		m2 := NewTLInputStickerSetShortName(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputStickerSet()
	case Cmd_InputStickerSetAnimatedEmoji:
		m2 := NewTLInputStickerSetAnimatedEmoji(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputStickerSet()
	case Cmd_InputStickerSetDice:
		m2 := NewTLInputStickerSetDice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputStickerSet()

	case Cmd_InputStickerSetDicee67f520e:
		m2 := NewTLInputStickerSetDice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputStickerSet()

	default:
		return fmt.Errorf("InputStickerSet Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputStickerSetItem) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputStickerSetItem:
		m2 := NewTLInputStickerSetItem(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputStickerSetItem()

	default:
		return fmt.Errorf("InputStickerSetItem Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputStickeredMedia) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputStickeredMediaPhoto:
		m2 := NewTLInputStickeredMediaPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputStickeredMedia()
	case Cmd_InputStickeredMediaDocument:
		m2 := NewTLInputStickeredMediaDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputStickeredMedia()

	default:
		return fmt.Errorf("InputStickeredMedia Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputTheme) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputTheme:
		m2 := NewTLInputTheme(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputTheme()
	case Cmd_InputThemeSlug:
		m2 := NewTLInputThemeSlug(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputTheme()

	default:
		return fmt.Errorf("InputTheme Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputThemeSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputThemeSettings:
		m2 := NewTLInputThemeSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputThemeSettings()

	default:
		return fmt.Errorf("InputThemeSettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputUser) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputUserEmpty:
		m2 := NewTLInputUserEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputUser()
	case Cmd_InputUserSelf:
		m2 := NewTLInputUserSelf(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputUser()
	case Cmd_InputUser:
		m2 := NewTLInputUser(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputUser()
	case Cmd_InputUserFromMessage:
		m2 := NewTLInputUserFromMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputUser()

	default:
		return fmt.Errorf("InputUser Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputWallPaper) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputWallPaper:
		m2 := NewTLInputWallPaper(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputWallPaper()
	case Cmd_InputWallPaperSlug:
		m2 := NewTLInputWallPaperSlug(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputWallPaper()
	case Cmd_InputWallPaperNoFile:
		m2 := NewTLInputWallPaperNoFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputWallPaper()

	default:
		return fmt.Errorf("InputWallPaper Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputWebDocument) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputWebDocument:
		m2 := NewTLInputWebDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputWebDocument()

	default:
		return fmt.Errorf("InputWebDocument Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *InputWebFileLocation) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputWebFileLocation:
		m2 := NewTLInputWebFileLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputWebFileLocation()
	case Cmd_InputWebFileGeoPointLocation:
		m2 := NewTLInputWebFileGeoPointLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_InputWebFileLocation()

	default:
		return fmt.Errorf("InputWebFileLocation Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Invoice) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Invoice:
		m2 := NewTLInvoice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Invoice()

	default:
		return fmt.Errorf("Invoice Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *IpPort) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_IpPort:
		m2 := NewTLIpPort(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_IpPort()
	case Cmd_IpPortSecret:
		m2 := NewTLIpPortSecret(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_IpPort()

	default:
		return fmt.Errorf("IpPort Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *JSONObjectValue) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_JsonObjectValue:
		m2 := NewTLJsonObjectValue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_JSONObjectValue()

	default:
		return fmt.Errorf("JSONObjectValue Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *JSONValue) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_JsonNull:
		m2 := NewTLJsonNull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_JSONValue()
	case Cmd_JsonBool:
		m2 := NewTLJsonBool(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_JSONValue()
	case Cmd_JsonNumber:
		m2 := NewTLJsonNumber(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_JSONValue()
	case Cmd_JsonString:
		m2 := NewTLJsonString(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_JSONValue()
	case Cmd_JsonArray:
		m2 := NewTLJsonArray(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_JSONValue()
	case Cmd_JsonObject:
		m2 := NewTLJsonObject(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_JSONValue()

	default:
		return fmt.Errorf("JSONValue Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *KeyboardButton) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_KeyboardButton:
		m2 := NewTLKeyboardButton(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()
	case Cmd_KeyboardButtonUrl:
		m2 := NewTLKeyboardButtonUrl(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()
	case Cmd_KeyboardButtonCallback:
		m2 := NewTLKeyboardButtonCallback(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()
	case Cmd_KeyboardButtonRequestPhone:
		m2 := NewTLKeyboardButtonRequestPhone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()
	case Cmd_KeyboardButtonRequestGeoLocation:
		m2 := NewTLKeyboardButtonRequestGeoLocation(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()
	case Cmd_KeyboardButtonSwitchInline:
		m2 := NewTLKeyboardButtonSwitchInline(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()
	case Cmd_KeyboardButtonGame:
		m2 := NewTLKeyboardButtonGame(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()
	case Cmd_KeyboardButtonBuy:
		m2 := NewTLKeyboardButtonBuy(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()
	case Cmd_KeyboardButtonUrlAuth:
		m2 := NewTLKeyboardButtonUrlAuth(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()
	case Cmd_InputKeyboardButtonUrlAuth:
		m2 := NewTLInputKeyboardButtonUrlAuth(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()
	case Cmd_KeyboardButtonRequestPoll:
		m2 := NewTLKeyboardButtonRequestPoll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()

	case Cmd_KeyboardButtonSwitchInlineea1b7a14:
		m2 := NewTLKeyboardButtonSwitchInline(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()
	case Cmd_KeyboardButtonCallback35bbdb6b:
		m2 := NewTLKeyboardButtonCallback(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButton()

	default:
		return fmt.Errorf("KeyboardButton Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *KeyboardButtonRow) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_KeyboardButtonRow:
		m2 := NewTLKeyboardButtonRow(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_KeyboardButtonRow()

	default:
		return fmt.Errorf("KeyboardButtonRow Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *LabeledPrice) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_LabeledPrice:
		m2 := NewTLLabeledPrice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_LabeledPrice()

	default:
		return fmt.Errorf("LabeledPrice Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *LangPackDifference) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_LangPackDifference:
		m2 := NewTLLangPackDifference(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_LangPackDifference()

	default:
		return fmt.Errorf("LangPackDifference Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *LangPackLanguage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_LangPackLanguage:
		m2 := NewTLLangPackLanguage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_LangPackLanguage()

	case Cmd_LangPackLanguage651b98d:
		m2 := NewTLLangPackLanguage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_LangPackLanguage()
	case Cmd_LangPackLanguageeeca5ce3:
		m2 := NewTLLangPackLanguage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_LangPackLanguage()

	default:
		return fmt.Errorf("LangPackLanguage Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *LangPackString) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_LangPackString:
		m2 := NewTLLangPackString(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_LangPackString()
	case Cmd_LangPackStringPluralized:
		m2 := NewTLLangPackStringPluralized(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_LangPackString()
	case Cmd_LangPackStringDeleted:
		m2 := NewTLLangPackStringDeleted(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_LangPackString()

	default:
		return fmt.Errorf("LangPackString Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MaskCoords) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MaskCoords:
		m2 := NewTLMaskCoords(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MaskCoords()

	default:
		return fmt.Errorf("MaskCoords Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Message) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageEmpty:
		m2 := NewTLMessageEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Message()
	case Cmd_Message:
		m2 := NewTLMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Message()
	case Cmd_MessageService:
		m2 := NewTLMessageService(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Message()

	case Cmd_Messagec09be45f:
		m2 := NewTLMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Message()
	case Cmd_Message44f9b43d:
		m2 := NewTLMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Message()
	case Cmd_Message452c0e65:
		m2 := NewTLMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Message()
	case Cmd_MessageService286fa604:
		m2 := NewTLMessageService(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Message()
	case Cmd_Message58ae39c9:
		m2 := NewTLMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Message()

	default:
		return fmt.Errorf("Message Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessageAction) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageActionEmpty:
		m2 := NewTLMessageActionEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionChatCreate:
		m2 := NewTLMessageActionChatCreate(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionChatEditTitle:
		m2 := NewTLMessageActionChatEditTitle(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionChatEditPhoto:
		m2 := NewTLMessageActionChatEditPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionChatDeletePhoto:
		m2 := NewTLMessageActionChatDeletePhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionChatAddUser:
		m2 := NewTLMessageActionChatAddUser(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionChatDeleteUser:
		m2 := NewTLMessageActionChatDeleteUser(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionChatJoinedByLink:
		m2 := NewTLMessageActionChatJoinedByLink(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionChannelCreate:
		m2 := NewTLMessageActionChannelCreate(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionChatMigrateTo:
		m2 := NewTLMessageActionChatMigrateTo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionChannelMigrateFrom:
		m2 := NewTLMessageActionChannelMigrateFrom(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionPinMessage:
		m2 := NewTLMessageActionPinMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionHistoryClear:
		m2 := NewTLMessageActionHistoryClear(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionGameScore:
		m2 := NewTLMessageActionGameScore(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionPaymentSentMe:
		m2 := NewTLMessageActionPaymentSentMe(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionPaymentSent:
		m2 := NewTLMessageActionPaymentSent(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionScreenshotTaken:
		m2 := NewTLMessageActionScreenshotTaken(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionCustomAction:
		m2 := NewTLMessageActionCustomAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionBotAllowed:
		m2 := NewTLMessageActionBotAllowed(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionSecureValuesSentMe:
		m2 := NewTLMessageActionSecureValuesSentMe(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionSecureValuesSent:
		m2 := NewTLMessageActionSecureValuesSent(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionContactSignUp:
		m2 := NewTLMessageActionContactSignUp(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionPhoneCall:
		m2 := NewTLMessageActionPhoneCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionGeoProximityReached:
		m2 := NewTLMessageActionGeoProximityReached(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionGroupCall:
		m2 := NewTLMessageActionGroupCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()
	case Cmd_MessageActionInviteToGroupCall:
		m2 := NewTLMessageActionInviteToGroupCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()

	case Cmd_MessageActionContactSignUpf3f25f76:
		m2 := NewTLMessageActionContactSignUp(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageAction()

	default:
		return fmt.Errorf("MessageAction Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessageEntity) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageEntityUnknown:
		m2 := NewTLMessageEntityUnknown(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityMention:
		m2 := NewTLMessageEntityMention(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityHashtag:
		m2 := NewTLMessageEntityHashtag(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityBotCommand:
		m2 := NewTLMessageEntityBotCommand(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityUrl:
		m2 := NewTLMessageEntityUrl(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityEmail:
		m2 := NewTLMessageEntityEmail(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityBold:
		m2 := NewTLMessageEntityBold(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityItalic:
		m2 := NewTLMessageEntityItalic(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityCode:
		m2 := NewTLMessageEntityCode(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityPre:
		m2 := NewTLMessageEntityPre(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityTextUrl:
		m2 := NewTLMessageEntityTextUrl(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityMentionName:
		m2 := NewTLMessageEntityMentionName(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_InputMessageEntityMentionName:
		m2 := NewTLInputMessageEntityMentionName(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityPhone:
		m2 := NewTLMessageEntityPhone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityCashtag:
		m2 := NewTLMessageEntityCashtag(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityUnderline:
		m2 := NewTLMessageEntityUnderline(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityStrike:
		m2 := NewTLMessageEntityStrike(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityBlockquote:
		m2 := NewTLMessageEntityBlockquote(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()
	case Cmd_MessageEntityBankCard:
		m2 := NewTLMessageEntityBankCard(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageEntity()

	default:
		return fmt.Errorf("MessageEntity Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessageFwdHeader) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageFwdHeader:
		m2 := NewTLMessageFwdHeader(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageFwdHeader()

	case Cmd_MessageFwdHeaderc786ddcb:
		m2 := NewTLMessageFwdHeader(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageFwdHeader()
	case Cmd_MessageFwdHeader559ebe6d:
		m2 := NewTLMessageFwdHeader(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageFwdHeader()
	case Cmd_MessageFwdHeaderec338270:
		m2 := NewTLMessageFwdHeader(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageFwdHeader()
	case Cmd_MessageFwdHeader353a686b:
		m2 := NewTLMessageFwdHeader(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageFwdHeader()
	case Cmd_MessageFwdHeader5f777dce:
		m2 := NewTLMessageFwdHeader(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageFwdHeader()

	default:
		return fmt.Errorf("MessageFwdHeader Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessageGroup) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageGroup:
		m2 := NewTLMessageGroup(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageGroup()

	default:
		return fmt.Errorf("MessageGroup Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessageInteractionCounters) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageInteractionCounters:
		m2 := NewTLMessageInteractionCounters(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageInteractionCounters()

	default:
		return fmt.Errorf("MessageInteractionCounters Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessageMedia) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageMediaEmpty:
		m2 := NewTLMessageMediaEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaPhoto:
		m2 := NewTLMessageMediaPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaGeo:
		m2 := NewTLMessageMediaGeo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaContact:
		m2 := NewTLMessageMediaContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaUnsupported:
		m2 := NewTLMessageMediaUnsupported(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaDocument:
		m2 := NewTLMessageMediaDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaWebPage:
		m2 := NewTLMessageMediaWebPage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaVenue:
		m2 := NewTLMessageMediaVenue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaGame:
		m2 := NewTLMessageMediaGame(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaInvoice:
		m2 := NewTLMessageMediaInvoice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaGeoLive:
		m2 := NewTLMessageMediaGeoLive(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaPoll:
		m2 := NewTLMessageMediaPoll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaDice:
		m2 := NewTLMessageMediaDice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()

	case Cmd_MessageMediaPhoto3d8ce53d:
		m2 := NewTLMessageMediaPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaDocumentf3e02ea8:
		m2 := NewTLMessageMediaDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaPhoto695150d7:
		m2 := NewTLMessageMediaPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaContactcbf24940:
		m2 := NewTLMessageMediaContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaDocument9cb070d7:
		m2 := NewTLMessageMediaDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaVenue2ec0533f:
		m2 := NewTLMessageMediaVenue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaDice3f7ee58b:
		m2 := NewTLMessageMediaDice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()
	case Cmd_MessageMediaGeoLiveb940c666:
		m2 := NewTLMessageMediaGeoLive(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageMedia()

	default:
		return fmt.Errorf("MessageMedia Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessageRange) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageRange:
		m2 := NewTLMessageRange(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageRange()

	default:
		return fmt.Errorf("MessageRange Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessageReplies) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageReplies:
		m2 := NewTLMessageReplies(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageReplies()

	default:
		return fmt.Errorf("MessageReplies Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessageReplyHeader) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageReplyHeader:
		m2 := NewTLMessageReplyHeader(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageReplyHeader()

	default:
		return fmt.Errorf("MessageReplyHeader Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessageUserVote) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageUserVote:
		m2 := NewTLMessageUserVote(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageUserVote()
	case Cmd_MessageUserVoteInputOption:
		m2 := NewTLMessageUserVoteInputOption(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageUserVote()
	case Cmd_MessageUserVoteMultiple:
		m2 := NewTLMessageUserVoteMultiple(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageUserVote()

	case Cmd_MessageUserVotea28e5559:
		m2 := NewTLMessageUserVote(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageUserVote()

	default:
		return fmt.Errorf("MessageUserVote Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessageViews) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessageViews:
		m2 := NewTLMessageViews(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessageViews()

	default:
		return fmt.Errorf("MessageViews Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MessagesFilter) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputMessagesFilterEmpty:
		m2 := NewTLInputMessagesFilterEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterPhotos:
		m2 := NewTLInputMessagesFilterPhotos(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterVideo:
		m2 := NewTLInputMessagesFilterVideo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterPhotoVideo:
		m2 := NewTLInputMessagesFilterPhotoVideo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterPhotoVideoDocuments:
		m2 := NewTLInputMessagesFilterPhotoVideoDocuments(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterDocument:
		m2 := NewTLInputMessagesFilterDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterUrl:
		m2 := NewTLInputMessagesFilterUrl(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterGif:
		m2 := NewTLInputMessagesFilterGif(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterVoice:
		m2 := NewTLInputMessagesFilterVoice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterMusic:
		m2 := NewTLInputMessagesFilterMusic(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterChatPhotos:
		m2 := NewTLInputMessagesFilterChatPhotos(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterPhoneCalls:
		m2 := NewTLInputMessagesFilterPhoneCalls(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterRoundVoice:
		m2 := NewTLInputMessagesFilterRoundVoice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterRoundVideo:
		m2 := NewTLInputMessagesFilterRoundVideo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterMyMentions:
		m2 := NewTLInputMessagesFilterMyMentions(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterGeo:
		m2 := NewTLInputMessagesFilterGeo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterContacts:
		m2 := NewTLInputMessagesFilterContacts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()
	case Cmd_InputMessagesFilterPinned:
		m2 := NewTLInputMessagesFilterPinned(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MessagesFilter()

	default:
		return fmt.Errorf("MessagesFilter Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_AffectedHistory) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesAffectedHistory:
		m2 := NewTLMessagesAffectedHistory(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_AffectedHistory()

	default:
		return fmt.Errorf("Messages_AffectedHistory Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_AffectedMessages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesAffectedMessages:
		m2 := NewTLMessagesAffectedMessages(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_AffectedMessages()

	default:
		return fmt.Errorf("Messages_AffectedMessages Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_AllStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesAllStickersNotModified:
		m2 := NewTLMessagesAllStickersNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_AllStickers()
	case Cmd_MessagesAllStickers:
		m2 := NewTLMessagesAllStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_AllStickers()

	default:
		return fmt.Errorf("Messages_AllStickers Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_ArchivedStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesArchivedStickers:
		m2 := NewTLMessagesArchivedStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_ArchivedStickers()

	default:
		return fmt.Errorf("Messages_ArchivedStickers Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_BotCallbackAnswer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesBotCallbackAnswer:
		m2 := NewTLMessagesBotCallbackAnswer(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_BotCallbackAnswer()

	case Cmd_MessagesBotCallbackAnswer1264f1c6:
		m2 := NewTLMessagesBotCallbackAnswer(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_BotCallbackAnswer()

	default:
		return fmt.Errorf("Messages_BotCallbackAnswer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_BotResults) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesBotResults:
		m2 := NewTLMessagesBotResults(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_BotResults()

	case Cmd_MessagesBotResults256709a6:
		m2 := NewTLMessagesBotResults(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_BotResults()
	case Cmd_MessagesBotResults947ca848:
		m2 := NewTLMessagesBotResults(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_BotResults()

	default:
		return fmt.Errorf("Messages_BotResults Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_ChatFull) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesChatFull:
		m2 := NewTLMessagesChatFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_ChatFull()

	default:
		return fmt.Errorf("Messages_ChatFull Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_Chats) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesChats:
		m2 := NewTLMessagesChats(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Chats()
	case Cmd_MessagesChatsSlice:
		m2 := NewTLMessagesChatsSlice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Chats()

	default:
		return fmt.Errorf("Messages_Chats Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_DhConfig) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesDhConfigNotModified:
		m2 := NewTLMessagesDhConfigNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_DhConfig()
	case Cmd_MessagesDhConfig:
		m2 := NewTLMessagesDhConfig(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_DhConfig()

	default:
		return fmt.Errorf("Messages_DhConfig Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_Dialogs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesDialogs:
		m2 := NewTLMessagesDialogs(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Dialogs()
	case Cmd_MessagesDialogsSlice:
		m2 := NewTLMessagesDialogsSlice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Dialogs()
	case Cmd_MessagesDialogsNotModified:
		m2 := NewTLMessagesDialogsNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Dialogs()

	default:
		return fmt.Errorf("Messages_Dialogs Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_DiscussionMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesDiscussionMessage:
		m2 := NewTLMessagesDiscussionMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_DiscussionMessage()

	default:
		return fmt.Errorf("Messages_DiscussionMessage Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_FavedStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesFavedStickersNotModified:
		m2 := NewTLMessagesFavedStickersNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_FavedStickers()
	case Cmd_MessagesFavedStickers:
		m2 := NewTLMessagesFavedStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_FavedStickers()

	default:
		return fmt.Errorf("Messages_FavedStickers Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_FeaturedStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesFeaturedStickersNotModified:
		m2 := NewTLMessagesFeaturedStickersNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_FeaturedStickers()
	case Cmd_MessagesFeaturedStickers:
		m2 := NewTLMessagesFeaturedStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_FeaturedStickers()

	case Cmd_MessagesFeaturedStickersNotModifiedc6dc0c66:
		m2 := NewTLMessagesFeaturedStickersNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_FeaturedStickers()
	case Cmd_MessagesFeaturedStickersb6abc341:
		m2 := NewTLMessagesFeaturedStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_FeaturedStickers()

	default:
		return fmt.Errorf("Messages_FeaturedStickers Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_FoundGifs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesFoundGifs:
		m2 := NewTLMessagesFoundGifs(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_FoundGifs()

	default:
		return fmt.Errorf("Messages_FoundGifs Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_FoundStickerSets) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesFoundStickerSetsNotModified:
		m2 := NewTLMessagesFoundStickerSetsNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_FoundStickerSets()
	case Cmd_MessagesFoundStickerSets:
		m2 := NewTLMessagesFoundStickerSets(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_FoundStickerSets()

	default:
		return fmt.Errorf("Messages_FoundStickerSets Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_HighScores) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesHighScores:
		m2 := NewTLMessagesHighScores(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_HighScores()

	default:
		return fmt.Errorf("Messages_HighScores Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_InactiveChats) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesInactiveChats:
		m2 := NewTLMessagesInactiveChats(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_InactiveChats()

	default:
		return fmt.Errorf("Messages_InactiveChats Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_MessageEditData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesMessageEditData:
		m2 := NewTLMessagesMessageEditData(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_MessageEditData()

	default:
		return fmt.Errorf("Messages_MessageEditData Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_MessageViews) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesMessageViews:
		m2 := NewTLMessagesMessageViews(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_MessageViews()

	default:
		return fmt.Errorf("Messages_MessageViews Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_Messages) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesMessages:
		m2 := NewTLMessagesMessages(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Messages()
	case Cmd_MessagesMessagesSlice:
		m2 := NewTLMessagesMessagesSlice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Messages()
	case Cmd_MessagesChannelMessages:
		m2 := NewTLMessagesChannelMessages(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Messages()
	case Cmd_MessagesMessagesNotModified:
		m2 := NewTLMessagesMessagesNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Messages()

	case Cmd_MessagesChannelMessagesbc0f17bc:
		m2 := NewTLMessagesChannelMessages(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Messages()
	case Cmd_MessagesMessagesSlicea6c47aaa:
		m2 := NewTLMessagesMessagesSlice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Messages()
	case Cmd_MessagesMessagesSlicec8edce1e:
		m2 := NewTLMessagesMessagesSlice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Messages()
	case Cmd_MessagesMessagesSlice3a54685e:
		m2 := NewTLMessagesMessagesSlice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Messages()
	case Cmd_MessagesChannelMessages64479808:
		m2 := NewTLMessagesChannelMessages(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Messages()

	default:
		return fmt.Errorf("Messages_Messages Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_PeerDialogs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesPeerDialogs:
		m2 := NewTLMessagesPeerDialogs(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_PeerDialogs()

	default:
		return fmt.Errorf("Messages_PeerDialogs Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_RecentStickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesRecentStickersNotModified:
		m2 := NewTLMessagesRecentStickersNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_RecentStickers()
	case Cmd_MessagesRecentStickers:
		m2 := NewTLMessagesRecentStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_RecentStickers()

	case Cmd_MessagesRecentStickers22f3afb3:
		m2 := NewTLMessagesRecentStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_RecentStickers()

	default:
		return fmt.Errorf("Messages_RecentStickers Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_SavedGifs) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesSavedGifsNotModified:
		m2 := NewTLMessagesSavedGifsNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_SavedGifs()
	case Cmd_MessagesSavedGifs:
		m2 := NewTLMessagesSavedGifs(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_SavedGifs()

	default:
		return fmt.Errorf("Messages_SavedGifs Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_SearchCounter) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesSearchCounter:
		m2 := NewTLMessagesSearchCounter(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_SearchCounter()

	default:
		return fmt.Errorf("Messages_SearchCounter Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_SentEncryptedMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesSentEncryptedMessage:
		m2 := NewTLMessagesSentEncryptedMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_SentEncryptedMessage()
	case Cmd_MessagesSentEncryptedFile:
		m2 := NewTLMessagesSentEncryptedFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_SentEncryptedMessage()

	default:
		return fmt.Errorf("Messages_SentEncryptedMessage Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_StickerSet) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesStickerSet:
		m2 := NewTLMessagesStickerSet(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_StickerSet()

	default:
		return fmt.Errorf("Messages_StickerSet Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_StickerSetInstallResult) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesStickerSetInstallResultSuccess:
		m2 := NewTLMessagesStickerSetInstallResultSuccess(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_StickerSetInstallResult()
	case Cmd_MessagesStickerSetInstallResultArchive:
		m2 := NewTLMessagesStickerSetInstallResultArchive(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_StickerSetInstallResult()

	default:
		return fmt.Errorf("Messages_StickerSetInstallResult Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_Stickers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesStickersNotModified:
		m2 := NewTLMessagesStickersNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Stickers()
	case Cmd_MessagesStickers:
		m2 := NewTLMessagesStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Stickers()

	case Cmd_MessagesStickerse4599bbd:
		m2 := NewTLMessagesStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_Stickers()

	default:
		return fmt.Errorf("Messages_Stickers Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Messages_VotesList) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MessagesVotesList:
		m2 := NewTLMessagesVotesList(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Messages_VotesList()

	default:
		return fmt.Errorf("Messages_VotesList Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MsgDetailedInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MsgDetailedInfo:
		m2 := NewTLMsgDetailedInfo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MsgDetailedInfo()
	case Cmd_MsgNewDetailedInfo:
		m2 := NewTLMsgNewDetailedInfo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MsgDetailedInfo()

	default:
		return fmt.Errorf("MsgDetailedInfo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MsgResendReq) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MsgResendReq:
		m2 := NewTLMsgResendReq(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MsgResendReq()

	default:
		return fmt.Errorf("MsgResendReq Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MsgsAck) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MsgsAck:
		m2 := NewTLMsgsAck(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MsgsAck()

	default:
		return fmt.Errorf("MsgsAck Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MsgsAllInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MsgsAllInfo:
		m2 := NewTLMsgsAllInfo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MsgsAllInfo()

	default:
		return fmt.Errorf("MsgsAllInfo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MsgsStateInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MsgsStateInfo:
		m2 := NewTLMsgsStateInfo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MsgsStateInfo()

	default:
		return fmt.Errorf("MsgsStateInfo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *MsgsStateReq) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_MsgsStateReq:
		m2 := NewTLMsgsStateReq(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_MsgsStateReq()

	default:
		return fmt.Errorf("MsgsStateReq Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *NearestDc) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_NearestDc:
		m2 := NewTLNearestDc(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_NearestDc()

	default:
		return fmt.Errorf("NearestDc Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *NewSession) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_NewSessionCreated:
		m2 := NewTLNewSessionCreated(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_NewSession()

	default:
		return fmt.Errorf("NewSession Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *NotifyPeer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_NotifyPeer:
		m2 := NewTLNotifyPeer(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_NotifyPeer()
	case Cmd_NotifyUsers:
		m2 := NewTLNotifyUsers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_NotifyPeer()
	case Cmd_NotifyChats:
		m2 := NewTLNotifyChats(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_NotifyPeer()
	case Cmd_NotifyAll:
		m2 := NewTLNotifyAll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_NotifyPeer()
	case Cmd_NotifyBroadcasts:
		m2 := NewTLNotifyBroadcasts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_NotifyPeer()

	default:
		return fmt.Errorf("NotifyPeer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Null) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Null:
		m2 := NewTLNull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Null()

	default:
		return fmt.Errorf("Null Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *P_QInnerData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PQInnerData:
		m2 := NewTLPQInnerData(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_P_QInnerData()
	case Cmd_PQInnerDataDc:
		m2 := NewTLPQInnerDataDc(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_P_QInnerData()
	case Cmd_PQInnerDataTemp:
		m2 := NewTLPQInnerDataTemp(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_P_QInnerData()
	case Cmd_PQInnerDataTempDc:
		m2 := NewTLPQInnerDataTempDc(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_P_QInnerData()

	default:
		return fmt.Errorf("P_QInnerData Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Page) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PagePart:
		m2 := NewTLPagePart(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Page()
	case Cmd_PageFull:
		m2 := NewTLPageFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Page()
	case Cmd_Page:
		m2 := NewTLPage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Page()

	case Cmd_Pageae891bec:
		m2 := NewTLPage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Page()
	case Cmd_Page98657f0d:
		m2 := NewTLPage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Page()

	default:
		return fmt.Errorf("Page Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PageBlock) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PageBlockUnsupported:
		m2 := NewTLPageBlockUnsupported(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockTitle:
		m2 := NewTLPageBlockTitle(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockSubtitle:
		m2 := NewTLPageBlockSubtitle(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockAuthorDate:
		m2 := NewTLPageBlockAuthorDate(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockHeader:
		m2 := NewTLPageBlockHeader(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockSubheader:
		m2 := NewTLPageBlockSubheader(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockParagraph:
		m2 := NewTLPageBlockParagraph(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockPreformatted:
		m2 := NewTLPageBlockPreformatted(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockFooter:
		m2 := NewTLPageBlockFooter(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockDivider:
		m2 := NewTLPageBlockDivider(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockAnchor:
		m2 := NewTLPageBlockAnchor(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockList:
		m2 := NewTLPageBlockList(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockBlockquote:
		m2 := NewTLPageBlockBlockquote(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockPullquote:
		m2 := NewTLPageBlockPullquote(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockPhoto:
		m2 := NewTLPageBlockPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockVideo:
		m2 := NewTLPageBlockVideo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockCover:
		m2 := NewTLPageBlockCover(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockEmbed:
		m2 := NewTLPageBlockEmbed(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockEmbedPost:
		m2 := NewTLPageBlockEmbedPost(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockCollage:
		m2 := NewTLPageBlockCollage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockSlideshow:
		m2 := NewTLPageBlockSlideshow(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockChannel:
		m2 := NewTLPageBlockChannel(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockAudio:
		m2 := NewTLPageBlockAudio(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockKicker:
		m2 := NewTLPageBlockKicker(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockTable:
		m2 := NewTLPageBlockTable(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockOrderedList:
		m2 := NewTLPageBlockOrderedList(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockDetails:
		m2 := NewTLPageBlockDetails(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockRelatedArticles:
		m2 := NewTLPageBlockRelatedArticles(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockMap:
		m2 := NewTLPageBlockMap(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()

	case Cmd_PageBlockListe4e88011:
		m2 := NewTLPageBlockList(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockPhoto1759c560:
		m2 := NewTLPageBlockPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockVideo7c8fe7b6:
		m2 := NewTLPageBlockVideo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockEmbeda8718dc5:
		m2 := NewTLPageBlockEmbed(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockEmbedPostf259a80b:
		m2 := NewTLPageBlockEmbedPost(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockCollage65a0fa4d:
		m2 := NewTLPageBlockCollage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockSlideshow31f9590:
		m2 := NewTLPageBlockSlideshow(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()
	case Cmd_PageBlockAudio804361ea:
		m2 := NewTLPageBlockAudio(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageBlock()

	default:
		return fmt.Errorf("PageBlock Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PageCaption) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PageCaption:
		m2 := NewTLPageCaption(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageCaption()

	default:
		return fmt.Errorf("PageCaption Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PageListItem) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PageListItemText:
		m2 := NewTLPageListItemText(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageListItem()
	case Cmd_PageListItemBlocks:
		m2 := NewTLPageListItemBlocks(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageListItem()

	default:
		return fmt.Errorf("PageListItem Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PageListOrderedItem) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PageListOrderedItemText:
		m2 := NewTLPageListOrderedItemText(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageListOrderedItem()
	case Cmd_PageListOrderedItemBlocks:
		m2 := NewTLPageListOrderedItemBlocks(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageListOrderedItem()

	default:
		return fmt.Errorf("PageListOrderedItem Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PageRelatedArticle) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PageRelatedArticle:
		m2 := NewTLPageRelatedArticle(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageRelatedArticle()

	case Cmd_PageRelatedArticleb390dc08:
		m2 := NewTLPageRelatedArticle(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageRelatedArticle()

	default:
		return fmt.Errorf("PageRelatedArticle Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PageTableCell) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PageTableCell:
		m2 := NewTLPageTableCell(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageTableCell()

	default:
		return fmt.Errorf("PageTableCell Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PageTableRow) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PageTableRow:
		m2 := NewTLPageTableRow(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PageTableRow()

	default:
		return fmt.Errorf("PageTableRow Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PasswordKdfAlgo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PasswordKdfAlgoUnknown:
		m2 := NewTLPasswordKdfAlgoUnknown(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PasswordKdfAlgo()
	case Cmd_PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow:
		m2 := NewTLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PasswordKdfAlgo()

	default:
		return fmt.Errorf("PasswordKdfAlgo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PaymentCharge) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PaymentCharge:
		m2 := NewTLPaymentCharge(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PaymentCharge()

	default:
		return fmt.Errorf("PaymentCharge Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PaymentRequestedInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PaymentRequestedInfo:
		m2 := NewTLPaymentRequestedInfo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PaymentRequestedInfo()

	default:
		return fmt.Errorf("PaymentRequestedInfo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PaymentSavedCredentials) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PaymentSavedCredentialsCard:
		m2 := NewTLPaymentSavedCredentialsCard(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PaymentSavedCredentials()

	default:
		return fmt.Errorf("PaymentSavedCredentials Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Payments_BankCardData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PaymentsBankCardData:
		m2 := NewTLPaymentsBankCardData(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Payments_BankCardData()

	default:
		return fmt.Errorf("Payments_BankCardData Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Payments_PaymentForm) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PaymentsPaymentForm:
		m2 := NewTLPaymentsPaymentForm(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Payments_PaymentForm()

	default:
		return fmt.Errorf("Payments_PaymentForm Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Payments_PaymentReceipt) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PaymentsPaymentReceipt:
		m2 := NewTLPaymentsPaymentReceipt(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Payments_PaymentReceipt()

	default:
		return fmt.Errorf("Payments_PaymentReceipt Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Payments_PaymentResult) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PaymentsPaymentResult:
		m2 := NewTLPaymentsPaymentResult(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Payments_PaymentResult()
	case Cmd_PaymentsPaymentVerficationNeeded:
		m2 := NewTLPaymentsPaymentVerficationNeeded(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Payments_PaymentResult()
	case Cmd_PaymentsPaymentVerificationNeeded:
		m2 := NewTLPaymentsPaymentVerificationNeeded(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Payments_PaymentResult()

	default:
		return fmt.Errorf("Payments_PaymentResult Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Payments_SavedInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PaymentsSavedInfo:
		m2 := NewTLPaymentsSavedInfo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Payments_SavedInfo()

	default:
		return fmt.Errorf("Payments_SavedInfo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Payments_ValidatedRequestedInfo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PaymentsValidatedRequestedInfo:
		m2 := NewTLPaymentsValidatedRequestedInfo(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Payments_ValidatedRequestedInfo()

	default:
		return fmt.Errorf("Payments_ValidatedRequestedInfo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Peer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PeerUser:
		m2 := NewTLPeerUser(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Peer()
	case Cmd_PeerChat:
		m2 := NewTLPeerChat(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Peer()
	case Cmd_PeerChannel:
		m2 := NewTLPeerChannel(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Peer()

	default:
		return fmt.Errorf("Peer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PeerBlocked) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PeerBlocked:
		m2 := NewTLPeerBlocked(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PeerBlocked()

	default:
		return fmt.Errorf("PeerBlocked Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PeerLocated) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PeerLocated:
		m2 := NewTLPeerLocated(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PeerLocated()
	case Cmd_PeerSelfLocated:
		m2 := NewTLPeerSelfLocated(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PeerLocated()

	default:
		return fmt.Errorf("PeerLocated Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PeerNotifyEvents) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PeerNotifyEventsEmpty:
		m2 := NewTLPeerNotifyEventsEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PeerNotifyEvents()
	case Cmd_PeerNotifyEventsAll:
		m2 := NewTLPeerNotifyEventsAll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PeerNotifyEvents()

	default:
		return fmt.Errorf("PeerNotifyEvents Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PeerNotifySettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PeerNotifySettingsEmpty:
		m2 := NewTLPeerNotifySettingsEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PeerNotifySettings()
	case Cmd_PeerNotifySettings:
		m2 := NewTLPeerNotifySettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PeerNotifySettings()

	case Cmd_PeerNotifySettingsaf509d20:
		m2 := NewTLPeerNotifySettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PeerNotifySettings()

	default:
		return fmt.Errorf("PeerNotifySettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PeerSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PeerSettings:
		m2 := NewTLPeerSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PeerSettings()

	case Cmd_PeerSettings733f2961:
		m2 := NewTLPeerSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PeerSettings()

	default:
		return fmt.Errorf("PeerSettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PhoneCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PhoneCallEmpty:
		m2 := NewTLPhoneCallEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCall()
	case Cmd_PhoneCallRequested:
		m2 := NewTLPhoneCallRequested(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCall()
	case Cmd_PhoneCallAccepted:
		m2 := NewTLPhoneCallAccepted(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCall()
	case Cmd_PhoneCall:
		m2 := NewTLPhoneCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCall()
	case Cmd_PhoneCallWaiting:
		m2 := NewTLPhoneCallWaiting(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCall()
	case Cmd_PhoneCallDiscarded:
		m2 := NewTLPhoneCallDiscarded(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCall()

	case Cmd_PhoneCalle6f9ddf3:
		m2 := NewTLPhoneCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCall()
	case Cmd_PhoneCallRequested87eabb53:
		m2 := NewTLPhoneCallRequested(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCall()
	case Cmd_PhoneCallAccepted997c454a:
		m2 := NewTLPhoneCallAccepted(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCall()
	case Cmd_PhoneCall8742ae7f:
		m2 := NewTLPhoneCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCall()

	default:
		return fmt.Errorf("PhoneCall Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PhoneCallDiscardReason) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PhoneCallDiscardReasonMissed:
		m2 := NewTLPhoneCallDiscardReasonMissed(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCallDiscardReason()
	case Cmd_PhoneCallDiscardReasonDisconnect:
		m2 := NewTLPhoneCallDiscardReasonDisconnect(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCallDiscardReason()
	case Cmd_PhoneCallDiscardReasonHangup:
		m2 := NewTLPhoneCallDiscardReasonHangup(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCallDiscardReason()
	case Cmd_PhoneCallDiscardReasonBusy:
		m2 := NewTLPhoneCallDiscardReasonBusy(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCallDiscardReason()

	default:
		return fmt.Errorf("PhoneCallDiscardReason Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PhoneCallProtocol) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PhoneCallProtocol:
		m2 := NewTLPhoneCallProtocol(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCallProtocol()

	case Cmd_PhoneCallProtocolfc878fc8:
		m2 := NewTLPhoneCallProtocol(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneCallProtocol()

	default:
		return fmt.Errorf("PhoneCallProtocol Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PhoneConnection) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PhoneConnection:
		m2 := NewTLPhoneConnection(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneConnection()
	case Cmd_PhoneConnectionWebrtc:
		m2 := NewTLPhoneConnectionWebrtc(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhoneConnection()

	default:
		return fmt.Errorf("PhoneConnection Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Phone_GroupCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PhoneGroupCall:
		m2 := NewTLPhoneGroupCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Phone_GroupCall()

	default:
		return fmt.Errorf("Phone_GroupCall Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Phone_GroupParticipants) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PhoneGroupParticipants:
		m2 := NewTLPhoneGroupParticipants(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Phone_GroupParticipants()

	default:
		return fmt.Errorf("Phone_GroupParticipants Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Phone_PhoneCall) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PhonePhoneCall:
		m2 := NewTLPhonePhoneCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Phone_PhoneCall()

	default:
		return fmt.Errorf("Phone_PhoneCall Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Photo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PhotoEmpty:
		m2 := NewTLPhotoEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Photo()
	case Cmd_Photo:
		m2 := NewTLPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Photo()

	case Cmd_Photocded42fe:
		m2 := NewTLPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Photo()
	case Cmd_Photo9c477dd8:
		m2 := NewTLPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Photo()
	case Cmd_Photod07504a5:
		m2 := NewTLPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Photo()
	case Cmd_Photofb197a65:
		m2 := NewTLPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Photo()

	default:
		return fmt.Errorf("Photo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PhotoSize) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PhotoSizeEmpty:
		m2 := NewTLPhotoSizeEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhotoSize()
	case Cmd_PhotoSize:
		m2 := NewTLPhotoSize(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhotoSize()
	case Cmd_PhotoCachedSize:
		m2 := NewTLPhotoCachedSize(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhotoSize()
	case Cmd_PhotoStrippedSize:
		m2 := NewTLPhotoStrippedSize(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhotoSize()
	case Cmd_PhotoSizeProgressive:
		m2 := NewTLPhotoSizeProgressive(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhotoSize()
	case Cmd_PhotoPathSize:
		m2 := NewTLPhotoPathSize(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PhotoSize()

	default:
		return fmt.Errorf("PhotoSize Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Photos_Photo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PhotosPhoto:
		m2 := NewTLPhotosPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Photos_Photo()

	default:
		return fmt.Errorf("Photos_Photo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Photos_Photos) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PhotosPhotos:
		m2 := NewTLPhotosPhotos(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Photos_Photos()
	case Cmd_PhotosPhotosSlice:
		m2 := NewTLPhotosPhotosSlice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Photos_Photos()

	default:
		return fmt.Errorf("Photos_Photos Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Poll) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Poll:
		m2 := NewTLPoll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Poll()

	case Cmd_Poll86e18161:
		m2 := NewTLPoll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Poll()

	default:
		return fmt.Errorf("Poll Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PollAnswer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PollAnswer:
		m2 := NewTLPollAnswer(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PollAnswer()

	default:
		return fmt.Errorf("PollAnswer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PollAnswerVoters) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PollAnswerVoters:
		m2 := NewTLPollAnswerVoters(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PollAnswerVoters()

	default:
		return fmt.Errorf("PollAnswerVoters Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PollResults) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PollResults:
		m2 := NewTLPollResults(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PollResults()

	case Cmd_PollResultsc87024a2:
		m2 := NewTLPollResults(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PollResults()
	case Cmd_PollResultsbadcc1a3:
		m2 := NewTLPollResults(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PollResults()

	default:
		return fmt.Errorf("PollResults Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Pong) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Pong:
		m2 := NewTLPong(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Pong()

	default:
		return fmt.Errorf("Pong Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PopularContact) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PopularContact:
		m2 := NewTLPopularContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PopularContact()

	default:
		return fmt.Errorf("PopularContact Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PostAddress) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PostAddress:
		m2 := NewTLPostAddress(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PostAddress()

	default:
		return fmt.Errorf("PostAddress Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PrivacyKey) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PrivacyKeyStatusTimestamp:
		m2 := NewTLPrivacyKeyStatusTimestamp(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyKey()
	case Cmd_PrivacyKeyChatInvite:
		m2 := NewTLPrivacyKeyChatInvite(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyKey()
	case Cmd_PrivacyKeyPhoneCall:
		m2 := NewTLPrivacyKeyPhoneCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyKey()
	case Cmd_PrivacyKeyPhoneP2P:
		m2 := NewTLPrivacyKeyPhoneP2P(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyKey()
	case Cmd_PrivacyKeyForwards:
		m2 := NewTLPrivacyKeyForwards(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyKey()
	case Cmd_PrivacyKeyProfilePhoto:
		m2 := NewTLPrivacyKeyProfilePhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyKey()
	case Cmd_PrivacyKeyPhoneNumber:
		m2 := NewTLPrivacyKeyPhoneNumber(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyKey()
	case Cmd_PrivacyKeyAddedByPhone:
		m2 := NewTLPrivacyKeyAddedByPhone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyKey()

	default:
		return fmt.Errorf("PrivacyKey Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *PrivacyRule) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_PrivacyValueAllowContacts:
		m2 := NewTLPrivacyValueAllowContacts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyRule()
	case Cmd_PrivacyValueAllowAll:
		m2 := NewTLPrivacyValueAllowAll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyRule()
	case Cmd_PrivacyValueAllowUsers:
		m2 := NewTLPrivacyValueAllowUsers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyRule()
	case Cmd_PrivacyValueDisallowContacts:
		m2 := NewTLPrivacyValueDisallowContacts(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyRule()
	case Cmd_PrivacyValueDisallowAll:
		m2 := NewTLPrivacyValueDisallowAll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyRule()
	case Cmd_PrivacyValueDisallowUsers:
		m2 := NewTLPrivacyValueDisallowUsers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyRule()
	case Cmd_PrivacyValueAllowChatParticipants:
		m2 := NewTLPrivacyValueAllowChatParticipants(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyRule()
	case Cmd_PrivacyValueDisallowChatParticipants:
		m2 := NewTLPrivacyValueDisallowChatParticipants(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_PrivacyRule()

	default:
		return fmt.Errorf("PrivacyRule Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ReceivedNotifyMessage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ReceivedNotifyMessage:
		m2 := NewTLReceivedNotifyMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReceivedNotifyMessage()

	default:
		return fmt.Errorf("ReceivedNotifyMessage Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *RecentMeUrl) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_RecentMeUrlUnknown:
		m2 := NewTLRecentMeUrlUnknown(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RecentMeUrl()
	case Cmd_RecentMeUrlUser:
		m2 := NewTLRecentMeUrlUser(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RecentMeUrl()
	case Cmd_RecentMeUrlChat:
		m2 := NewTLRecentMeUrlChat(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RecentMeUrl()
	case Cmd_RecentMeUrlChatInvite:
		m2 := NewTLRecentMeUrlChatInvite(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RecentMeUrl()
	case Cmd_RecentMeUrlStickerSet:
		m2 := NewTLRecentMeUrlStickerSet(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RecentMeUrl()

	default:
		return fmt.Errorf("RecentMeUrl Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ReplyMarkup) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ReplyKeyboardHide:
		m2 := NewTLReplyKeyboardHide(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReplyMarkup()
	case Cmd_ReplyKeyboardForceReply:
		m2 := NewTLReplyKeyboardForceReply(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReplyMarkup()
	case Cmd_ReplyKeyboardMarkup:
		m2 := NewTLReplyKeyboardMarkup(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReplyMarkup()
	case Cmd_ReplyInlineMarkup:
		m2 := NewTLReplyInlineMarkup(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReplyMarkup()

	default:
		return fmt.Errorf("ReplyMarkup Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ReportReason) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_InputReportReasonSpam:
		m2 := NewTLInputReportReasonSpam(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReportReason()
	case Cmd_InputReportReasonViolence:
		m2 := NewTLInputReportReasonViolence(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReportReason()
	case Cmd_InputReportReasonPornography:
		m2 := NewTLInputReportReasonPornography(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReportReason()
	case Cmd_InputReportReasonOther:
		m2 := NewTLInputReportReasonOther(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReportReason()
	case Cmd_InputReportReasonCopyright:
		m2 := NewTLInputReportReasonCopyright(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReportReason()
	case Cmd_InputReportReasonChildAbuse:
		m2 := NewTLInputReportReasonChildAbuse(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReportReason()
	case Cmd_InputReportReasonGeoIrrelevant:
		m2 := NewTLInputReportReasonGeoIrrelevant(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ReportReason()

	default:
		return fmt.Errorf("ReportReason Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ResPQ) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ResPQ:
		m2 := NewTLResPQ(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ResPQ()

	default:
		return fmt.Errorf("ResPQ Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *RestrictionReason) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_RestrictionReason:
		m2 := NewTLRestrictionReason(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RestrictionReason()

	default:
		return fmt.Errorf("RestrictionReason Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *RichText) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_TextEmpty:
		m2 := NewTLTextEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextPlain:
		m2 := NewTLTextPlain(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextBold:
		m2 := NewTLTextBold(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextItalic:
		m2 := NewTLTextItalic(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextUnderline:
		m2 := NewTLTextUnderline(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextStrike:
		m2 := NewTLTextStrike(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextFixed:
		m2 := NewTLTextFixed(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextUrl:
		m2 := NewTLTextUrl(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextEmail:
		m2 := NewTLTextEmail(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextConcat:
		m2 := NewTLTextConcat(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextSubscript:
		m2 := NewTLTextSubscript(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextSuperscript:
		m2 := NewTLTextSuperscript(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextMarked:
		m2 := NewTLTextMarked(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextPhone:
		m2 := NewTLTextPhone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextImage:
		m2 := NewTLTextImage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()
	case Cmd_TextAnchor:
		m2 := NewTLTextAnchor(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RichText()

	default:
		return fmt.Errorf("RichText Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *RpcDropAnswer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_RpcAnswerUnknown:
		m2 := NewTLRpcAnswerUnknown(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RpcDropAnswer()
	case Cmd_RpcAnswerDroppedRunning:
		m2 := NewTLRpcAnswerDroppedRunning(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RpcDropAnswer()
	case Cmd_RpcAnswerDropped:
		m2 := NewTLRpcAnswerDropped(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RpcDropAnswer()

	default:
		return fmt.Errorf("RpcDropAnswer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *RpcError) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_RpcError:
		m2 := NewTLRpcError(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_RpcError()

	default:
		return fmt.Errorf("RpcError Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SavedContact) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SavedPhoneContact:
		m2 := NewTLSavedPhoneContact(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SavedContact()

	default:
		return fmt.Errorf("SavedContact Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SecureCredentialsEncrypted) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SecureCredentialsEncrypted:
		m2 := NewTLSecureCredentialsEncrypted(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureCredentialsEncrypted()

	default:
		return fmt.Errorf("SecureCredentialsEncrypted Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SecureData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SecureData:
		m2 := NewTLSecureData(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureData()

	default:
		return fmt.Errorf("SecureData Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SecureFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SecureFileEmpty:
		m2 := NewTLSecureFileEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureFile()
	case Cmd_SecureFile:
		m2 := NewTLSecureFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureFile()

	default:
		return fmt.Errorf("SecureFile Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SecurePasswordKdfAlgo) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SecurePasswordKdfAlgoUnknown:
		m2 := NewTLSecurePasswordKdfAlgoUnknown(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecurePasswordKdfAlgo()
	case Cmd_SecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000:
		m2 := NewTLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecurePasswordKdfAlgo()
	case Cmd_SecurePasswordKdfAlgoSHA512:
		m2 := NewTLSecurePasswordKdfAlgoSHA512(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecurePasswordKdfAlgo()

	default:
		return fmt.Errorf("SecurePasswordKdfAlgo Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SecurePlainData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SecurePlainPhone:
		m2 := NewTLSecurePlainPhone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecurePlainData()
	case Cmd_SecurePlainEmail:
		m2 := NewTLSecurePlainEmail(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecurePlainData()

	default:
		return fmt.Errorf("SecurePlainData Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SecureRequiredType) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SecureRequiredType:
		m2 := NewTLSecureRequiredType(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureRequiredType()
	case Cmd_SecureRequiredTypeOneOf:
		m2 := NewTLSecureRequiredTypeOneOf(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureRequiredType()

	default:
		return fmt.Errorf("SecureRequiredType Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SecureSecretSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SecureSecretSettings:
		m2 := NewTLSecureSecretSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureSecretSettings()

	default:
		return fmt.Errorf("SecureSecretSettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SecureValue) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SecureValue:
		m2 := NewTLSecureValue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValue()

	default:
		return fmt.Errorf("SecureValue Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SecureValueError) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SecureValueErrorData:
		m2 := NewTLSecureValueErrorData(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueError()
	case Cmd_SecureValueErrorFrontSide:
		m2 := NewTLSecureValueErrorFrontSide(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueError()
	case Cmd_SecureValueErrorReverseSide:
		m2 := NewTLSecureValueErrorReverseSide(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueError()
	case Cmd_SecureValueErrorSelfie:
		m2 := NewTLSecureValueErrorSelfie(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueError()
	case Cmd_SecureValueErrorFile:
		m2 := NewTLSecureValueErrorFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueError()
	case Cmd_SecureValueErrorFiles:
		m2 := NewTLSecureValueErrorFiles(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueError()
	case Cmd_SecureValueError:
		m2 := NewTLSecureValueError(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueError()
	case Cmd_SecureValueErrorTranslationFile:
		m2 := NewTLSecureValueErrorTranslationFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueError()
	case Cmd_SecureValueErrorTranslationFiles:
		m2 := NewTLSecureValueErrorTranslationFiles(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueError()

	default:
		return fmt.Errorf("SecureValueError Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SecureValueHash) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SecureValueHash:
		m2 := NewTLSecureValueHash(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueHash()

	default:
		return fmt.Errorf("SecureValueHash Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SecureValueType) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SecureValueTypePersonalDetails:
		m2 := NewTLSecureValueTypePersonalDetails(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypePassport:
		m2 := NewTLSecureValueTypePassport(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypeDriverLicense:
		m2 := NewTLSecureValueTypeDriverLicense(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypeIdentityCard:
		m2 := NewTLSecureValueTypeIdentityCard(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypeInternalPassport:
		m2 := NewTLSecureValueTypeInternalPassport(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypeAddress:
		m2 := NewTLSecureValueTypeAddress(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypeUtilityBill:
		m2 := NewTLSecureValueTypeUtilityBill(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypeBankStatement:
		m2 := NewTLSecureValueTypeBankStatement(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypeRentalAgreement:
		m2 := NewTLSecureValueTypeRentalAgreement(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypePassportRegistration:
		m2 := NewTLSecureValueTypePassportRegistration(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypeTemporaryRegistration:
		m2 := NewTLSecureValueTypeTemporaryRegistration(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypePhone:
		m2 := NewTLSecureValueTypePhone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()
	case Cmd_SecureValueTypeEmail:
		m2 := NewTLSecureValueTypeEmail(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SecureValueType()

	default:
		return fmt.Errorf("SecureValueType Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SendMessageAction) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_SendMessageTypingAction:
		m2 := NewTLSendMessageTypingAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageCancelAction:
		m2 := NewTLSendMessageCancelAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageRecordVideoAction:
		m2 := NewTLSendMessageRecordVideoAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageUploadVideoAction:
		m2 := NewTLSendMessageUploadVideoAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageRecordAudioAction:
		m2 := NewTLSendMessageRecordAudioAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageUploadAudioAction:
		m2 := NewTLSendMessageUploadAudioAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageUploadPhotoAction:
		m2 := NewTLSendMessageUploadPhotoAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageUploadDocumentAction:
		m2 := NewTLSendMessageUploadDocumentAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageGeoLocationAction:
		m2 := NewTLSendMessageGeoLocationAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageChooseContactAction:
		m2 := NewTLSendMessageChooseContactAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageGamePlayAction:
		m2 := NewTLSendMessageGamePlayAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageRecordRoundAction:
		m2 := NewTLSendMessageRecordRoundAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SendMessageUploadRoundAction:
		m2 := NewTLSendMessageUploadRoundAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()
	case Cmd_SpeakingInGroupCallAction:
		m2 := NewTLSpeakingInGroupCallAction(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SendMessageAction()

	default:
		return fmt.Errorf("SendMessageAction Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Server_DHInnerData) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Server_DHInnerData:
		m2 := NewTLServer_DHInnerData(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Server_DHInnerData()

	default:
		return fmt.Errorf("Server_DHInnerData Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Server_DH_Params) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_Server_DHParamsFail:
		m2 := NewTLServer_DHParamsFail(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Server_DH_Params()
	case Cmd_Server_DHParamsOk:
		m2 := NewTLServer_DHParamsOk(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Server_DH_Params()

	default:
		return fmt.Errorf("Server_DH_Params Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *SetClient_DHParamsAnswer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_DhGenOk:
		m2 := NewTLDhGenOk(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SetClient_DHParamsAnswer()
	case Cmd_DhGenRetry:
		m2 := NewTLDhGenRetry(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SetClient_DHParamsAnswer()
	case Cmd_DhGenFail:
		m2 := NewTLDhGenFail(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_SetClient_DHParamsAnswer()

	default:
		return fmt.Errorf("SetClient_DHParamsAnswer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ShippingOption) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ShippingOption:
		m2 := NewTLShippingOption(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ShippingOption()

	default:
		return fmt.Errorf("ShippingOption Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *StatsAbsValueAndPrev) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StatsAbsValueAndPrev:
		m2 := NewTLStatsAbsValueAndPrev(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StatsAbsValueAndPrev()

	default:
		return fmt.Errorf("StatsAbsValueAndPrev Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *StatsDateRangeDays) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StatsDateRangeDays:
		m2 := NewTLStatsDateRangeDays(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StatsDateRangeDays()

	default:
		return fmt.Errorf("StatsDateRangeDays Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *StatsGraph) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StatsGraphAsync:
		m2 := NewTLStatsGraphAsync(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StatsGraph()
	case Cmd_StatsGraphError:
		m2 := NewTLStatsGraphError(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StatsGraph()
	case Cmd_StatsGraph:
		m2 := NewTLStatsGraph(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StatsGraph()

	default:
		return fmt.Errorf("StatsGraph Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *StatsGroupTopAdmin) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StatsGroupTopAdmin:
		m2 := NewTLStatsGroupTopAdmin(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StatsGroupTopAdmin()

	default:
		return fmt.Errorf("StatsGroupTopAdmin Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *StatsGroupTopInviter) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StatsGroupTopInviter:
		m2 := NewTLStatsGroupTopInviter(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StatsGroupTopInviter()

	default:
		return fmt.Errorf("StatsGroupTopInviter Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *StatsGroupTopPoster) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StatsGroupTopPoster:
		m2 := NewTLStatsGroupTopPoster(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StatsGroupTopPoster()

	default:
		return fmt.Errorf("StatsGroupTopPoster Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *StatsPercentValue) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StatsPercentValue:
		m2 := NewTLStatsPercentValue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StatsPercentValue()

	default:
		return fmt.Errorf("StatsPercentValue Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *StatsURL) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StatsURL:
		m2 := NewTLStatsURL(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StatsURL()

	default:
		return fmt.Errorf("StatsURL Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Stats_BroadcastStats) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StatsBroadcastStats:
		m2 := NewTLStatsBroadcastStats(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Stats_BroadcastStats()

	default:
		return fmt.Errorf("Stats_BroadcastStats Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Stats_MegagroupStats) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StatsMegagroupStats:
		m2 := NewTLStatsMegagroupStats(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Stats_MegagroupStats()

	default:
		return fmt.Errorf("Stats_MegagroupStats Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Stats_MessageStats) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StatsMessageStats:
		m2 := NewTLStatsMessageStats(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Stats_MessageStats()

	default:
		return fmt.Errorf("Stats_MessageStats Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *StickerPack) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StickerPack:
		m2 := NewTLStickerPack(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StickerPack()

	default:
		return fmt.Errorf("StickerPack Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *StickerSet) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StickerSet:
		m2 := NewTLStickerSet(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StickerSet()

	case Cmd_StickerSet5585a139:
		m2 := NewTLStickerSet(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StickerSet()
	case Cmd_StickerSet6a90bcb7:
		m2 := NewTLStickerSet(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StickerSet()
	case Cmd_StickerSeteeb46f27:
		m2 := NewTLStickerSet(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StickerSet()
	case Cmd_StickerSet40e237a8:
		m2 := NewTLStickerSet(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StickerSet()

	default:
		return fmt.Errorf("StickerSet Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *StickerSetCovered) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StickerSetCovered:
		m2 := NewTLStickerSetCovered(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StickerSetCovered()
	case Cmd_StickerSetMultiCovered:
		m2 := NewTLStickerSetMultiCovered(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_StickerSetCovered()

	default:
		return fmt.Errorf("StickerSetCovered Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Storage_FileType) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_StorageFileUnknown:
		m2 := NewTLStorageFileUnknown(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Storage_FileType()
	case Cmd_StorageFilePartial:
		m2 := NewTLStorageFilePartial(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Storage_FileType()
	case Cmd_StorageFileJpeg:
		m2 := NewTLStorageFileJpeg(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Storage_FileType()
	case Cmd_StorageFileGif:
		m2 := NewTLStorageFileGif(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Storage_FileType()
	case Cmd_StorageFilePng:
		m2 := NewTLStorageFilePng(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Storage_FileType()
	case Cmd_StorageFilePdf:
		m2 := NewTLStorageFilePdf(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Storage_FileType()
	case Cmd_StorageFileMp3:
		m2 := NewTLStorageFileMp3(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Storage_FileType()
	case Cmd_StorageFileMov:
		m2 := NewTLStorageFileMov(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Storage_FileType()
	case Cmd_StorageFileMp4:
		m2 := NewTLStorageFileMp4(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Storage_FileType()
	case Cmd_StorageFileWebp:
		m2 := NewTLStorageFileWebp(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Storage_FileType()

	default:
		return fmt.Errorf("Storage_FileType Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Theme) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ThemeDocumentNotModified:
		m2 := NewTLThemeDocumentNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Theme()
	case Cmd_Theme:
		m2 := NewTLTheme(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Theme()

	case Cmd_Themef7d90ce0:
		m2 := NewTLTheme(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Theme()
	case Cmd_Theme28f1114:
		m2 := NewTLTheme(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Theme()

	default:
		return fmt.Errorf("Theme Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *ThemeSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_ThemeSettings:
		m2 := NewTLThemeSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_ThemeSettings()

	default:
		return fmt.Errorf("ThemeSettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *TopPeer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_TopPeer:
		m2 := NewTLTopPeer(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_TopPeer()

	default:
		return fmt.Errorf("TopPeer Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *TopPeerCategory) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_TopPeerCategoryBotsPM:
		m2 := NewTLTopPeerCategoryBotsPM(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_TopPeerCategory()
	case Cmd_TopPeerCategoryBotsInline:
		m2 := NewTLTopPeerCategoryBotsInline(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_TopPeerCategory()
	case Cmd_TopPeerCategoryCorrespondents:
		m2 := NewTLTopPeerCategoryCorrespondents(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_TopPeerCategory()
	case Cmd_TopPeerCategoryGroups:
		m2 := NewTLTopPeerCategoryGroups(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_TopPeerCategory()
	case Cmd_TopPeerCategoryChannels:
		m2 := NewTLTopPeerCategoryChannels(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_TopPeerCategory()
	case Cmd_TopPeerCategoryPhoneCalls:
		m2 := NewTLTopPeerCategoryPhoneCalls(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_TopPeerCategory()
	case Cmd_TopPeerCategoryForwardUsers:
		m2 := NewTLTopPeerCategoryForwardUsers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_TopPeerCategory()
	case Cmd_TopPeerCategoryForwardChats:
		m2 := NewTLTopPeerCategoryForwardChats(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_TopPeerCategory()

	default:
		return fmt.Errorf("TopPeerCategory Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *TopPeerCategoryPeers) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_TopPeerCategoryPeers:
		m2 := NewTLTopPeerCategoryPeers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_TopPeerCategoryPeers()

	default:
		return fmt.Errorf("TopPeerCategoryPeers Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *True) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_True:
		m2 := NewTLTrue(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_True()

	default:
		return fmt.Errorf("True Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Update) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UpdateNewMessage:
		m2 := NewTLUpdateNewMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateMessageID:
		m2 := NewTLUpdateMessageID(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDeleteMessages:
		m2 := NewTLUpdateDeleteMessages(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateUserTyping:
		m2 := NewTLUpdateUserTyping(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChatUserTyping:
		m2 := NewTLUpdateChatUserTyping(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChatParticipants:
		m2 := NewTLUpdateChatParticipants(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateUserStatus:
		m2 := NewTLUpdateUserStatus(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateUserName:
		m2 := NewTLUpdateUserName(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateUserPhoto:
		m2 := NewTLUpdateUserPhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateContactRegistered:
		m2 := NewTLUpdateContactRegistered(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateContactLink:
		m2 := NewTLUpdateContactLink(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateNewEncryptedMessage:
		m2 := NewTLUpdateNewEncryptedMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateEncryptedChatTyping:
		m2 := NewTLUpdateEncryptedChatTyping(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateEncryption:
		m2 := NewTLUpdateEncryption(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateEncryptedMessagesRead:
		m2 := NewTLUpdateEncryptedMessagesRead(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChatParticipantAdd:
		m2 := NewTLUpdateChatParticipantAdd(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChatParticipantDelete:
		m2 := NewTLUpdateChatParticipantDelete(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDcOptions:
		m2 := NewTLUpdateDcOptions(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateUserBlocked:
		m2 := NewTLUpdateUserBlocked(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateNotifySettings:
		m2 := NewTLUpdateNotifySettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateServiceNotification:
		m2 := NewTLUpdateServiceNotification(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePrivacy:
		m2 := NewTLUpdatePrivacy(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateUserPhone:
		m2 := NewTLUpdateUserPhone(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateReadHistoryInbox:
		m2 := NewTLUpdateReadHistoryInbox(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateReadHistoryOutbox:
		m2 := NewTLUpdateReadHistoryOutbox(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateWebPage:
		m2 := NewTLUpdateWebPage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateReadMessagesContents:
		m2 := NewTLUpdateReadMessagesContents(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChannelTooLong:
		m2 := NewTLUpdateChannelTooLong(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChannel:
		m2 := NewTLUpdateChannel(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateNewChannelMessage:
		m2 := NewTLUpdateNewChannelMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateReadChannelInbox:
		m2 := NewTLUpdateReadChannelInbox(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDeleteChannelMessages:
		m2 := NewTLUpdateDeleteChannelMessages(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChannelMessageViews:
		m2 := NewTLUpdateChannelMessageViews(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChatAdmins:
		m2 := NewTLUpdateChatAdmins(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChatParticipantAdmin:
		m2 := NewTLUpdateChatParticipantAdmin(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateNewStickerSet:
		m2 := NewTLUpdateNewStickerSet(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateStickerSetsOrder:
		m2 := NewTLUpdateStickerSetsOrder(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateStickerSets:
		m2 := NewTLUpdateStickerSets(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateSavedGifs:
		m2 := NewTLUpdateSavedGifs(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateBotInlineQuery:
		m2 := NewTLUpdateBotInlineQuery(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateBotInlineSend:
		m2 := NewTLUpdateBotInlineSend(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateEditChannelMessage:
		m2 := NewTLUpdateEditChannelMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChannelPinnedMessage:
		m2 := NewTLUpdateChannelPinnedMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateBotCallbackQuery:
		m2 := NewTLUpdateBotCallbackQuery(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateEditMessage:
		m2 := NewTLUpdateEditMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateInlineBotCallbackQuery:
		m2 := NewTLUpdateInlineBotCallbackQuery(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateReadChannelOutbox:
		m2 := NewTLUpdateReadChannelOutbox(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDraftMessage:
		m2 := NewTLUpdateDraftMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateReadFeaturedStickers:
		m2 := NewTLUpdateReadFeaturedStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateRecentStickers:
		m2 := NewTLUpdateRecentStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateConfig:
		m2 := NewTLUpdateConfig(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePtsChanged:
		m2 := NewTLUpdatePtsChanged(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChannelWebPage:
		m2 := NewTLUpdateChannelWebPage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDialogPinned:
		m2 := NewTLUpdateDialogPinned(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePinnedDialogs:
		m2 := NewTLUpdatePinnedDialogs(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateBotWebhookJSON:
		m2 := NewTLUpdateBotWebhookJSON(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateBotWebhookJSONQuery:
		m2 := NewTLUpdateBotWebhookJSONQuery(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateBotShippingQuery:
		m2 := NewTLUpdateBotShippingQuery(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateBotPrecheckoutQuery:
		m2 := NewTLUpdateBotPrecheckoutQuery(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePhoneCall:
		m2 := NewTLUpdatePhoneCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateLangPackTooLong:
		m2 := NewTLUpdateLangPackTooLong(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateLangPack:
		m2 := NewTLUpdateLangPack(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateFavedStickers:
		m2 := NewTLUpdateFavedStickers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChannelReadMessagesContents:
		m2 := NewTLUpdateChannelReadMessagesContents(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateContactsReset:
		m2 := NewTLUpdateContactsReset(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateNewAuthorization:
		m2 := NewTLUpdateNewAuthorization(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChannelGroup:
		m2 := NewTLUpdateChannelGroup(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChannelAvailableMessages:
		m2 := NewTLUpdateChannelAvailableMessages(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDialogUnreadMark:
		m2 := NewTLUpdateDialogUnreadMark(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateUserPinnedMessage:
		m2 := NewTLUpdateUserPinnedMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChatPinnedMessage:
		m2 := NewTLUpdateChatPinnedMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateMessagePoll:
		m2 := NewTLUpdateMessagePoll(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChatDefaultBannedRights:
		m2 := NewTLUpdateChatDefaultBannedRights(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateFolderPeers:
		m2 := NewTLUpdateFolderPeers(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePeerSettings:
		m2 := NewTLUpdatePeerSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePeerLocated:
		m2 := NewTLUpdatePeerLocated(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateNewScheduledMessage:
		m2 := NewTLUpdateNewScheduledMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDeleteScheduledMessages:
		m2 := NewTLUpdateDeleteScheduledMessages(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateTheme:
		m2 := NewTLUpdateTheme(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateGeoLiveViewed:
		m2 := NewTLUpdateGeoLiveViewed(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateLoginToken:
		m2 := NewTLUpdateLoginToken(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateMessagePollVote:
		m2 := NewTLUpdateMessagePollVote(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDialogFilter:
		m2 := NewTLUpdateDialogFilter(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDialogFilterOrder:
		m2 := NewTLUpdateDialogFilterOrder(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDialogFilters:
		m2 := NewTLUpdateDialogFilters(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePhoneCallSignalingData:
		m2 := NewTLUpdatePhoneCallSignalingData(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChannelParticipant:
		m2 := NewTLUpdateChannelParticipant(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChannelMessageForwards:
		m2 := NewTLUpdateChannelMessageForwards(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateReadChannelDiscussionInbox:
		m2 := NewTLUpdateReadChannelDiscussionInbox(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateReadChannelDiscussionOutbox:
		m2 := NewTLUpdateReadChannelDiscussionOutbox(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePeerBlocked:
		m2 := NewTLUpdatePeerBlocked(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChannelUserTyping:
		m2 := NewTLUpdateChannelUserTyping(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePinnedMessages:
		m2 := NewTLUpdatePinnedMessages(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePinnedChannelMessages:
		m2 := NewTLUpdatePinnedChannelMessages(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChat:
		m2 := NewTLUpdateChat(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateGroupCallParticipants:
		m2 := NewTLUpdateGroupCallParticipants(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateGroupCall:
		m2 := NewTLUpdateGroupCall(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()

	case Cmd_UpdateServiceNotification382dd3e4:
		m2 := NewTLUpdateServiceNotification(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateStickerSetsOrderf0dfb451:
		m2 := NewTLUpdateStickerSetsOrder(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateBotCallbackQuerya68c688c:
		m2 := NewTLUpdateBotCallbackQuery(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateInlineBotCallbackQuery2cbd95af:
		m2 := NewTLUpdateInlineBotCallbackQuery(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDialogPinned19d27f3c:
		m2 := NewTLUpdateDialogPinned(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePinnedDialogsea4cb65b:
		m2 := NewTLUpdatePinnedDialogs(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateLangPackTooLong46560264:
		m2 := NewTLUpdateLangPackTooLong(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateChatPinnedMessagee10db349:
		m2 := NewTLUpdateChatPinnedMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateReadHistoryInbox9c974fdf:
		m2 := NewTLUpdateReadHistoryInbox(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateReadChannelInbox330b5424:
		m2 := NewTLUpdateReadChannelInbox(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateDialogPinned6e6fe51c:
		m2 := NewTLUpdateDialogPinned(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdatePinnedDialogsfa0f3ca2:
		m2 := NewTLUpdatePinnedDialogs(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()
	case Cmd_UpdateBotInlineQuery3f2038db:
		m2 := NewTLUpdateBotInlineQuery(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Update()

	default:
		return fmt.Errorf("Update Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Updates) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UpdatesTooLong:
		m2 := NewTLUpdatesTooLong(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates()
	case Cmd_UpdateShortMessage:
		m2 := NewTLUpdateShortMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates()
	case Cmd_UpdateShortChatMessage:
		m2 := NewTLUpdateShortChatMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates()
	case Cmd_UpdateShort:
		m2 := NewTLUpdateShort(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates()
	case Cmd_UpdatesCombined:
		m2 := NewTLUpdatesCombined(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates()
	case Cmd_Updates:
		m2 := NewTLUpdates(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates()
	case Cmd_UpdateShortSentMessage:
		m2 := NewTLUpdateShortSentMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates()

	case Cmd_UpdateShortMessage2296d2c8:
		m2 := NewTLUpdateShortMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates()
	case Cmd_UpdateShortChatMessage402d5dbb:
		m2 := NewTLUpdateShortChatMessage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates()

	default:
		return fmt.Errorf("Updates Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Updates_ChannelDifference) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UpdatesChannelDifferenceEmpty:
		m2 := NewTLUpdatesChannelDifferenceEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates_ChannelDifference()
	case Cmd_UpdatesChannelDifferenceTooLong:
		m2 := NewTLUpdatesChannelDifferenceTooLong(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates_ChannelDifference()
	case Cmd_UpdatesChannelDifference:
		m2 := NewTLUpdatesChannelDifference(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates_ChannelDifference()

	case Cmd_UpdatesChannelDifferenceTooLong5e167646:
		m2 := NewTLUpdatesChannelDifferenceTooLong(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates_ChannelDifference()
	case Cmd_UpdatesChannelDifferenceTooLonga4bcc6fe:
		m2 := NewTLUpdatesChannelDifferenceTooLong(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates_ChannelDifference()

	default:
		return fmt.Errorf("Updates_ChannelDifference Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Updates_Difference) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UpdatesDifferenceEmpty:
		m2 := NewTLUpdatesDifferenceEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates_Difference()
	case Cmd_UpdatesDifference:
		m2 := NewTLUpdatesDifference(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates_Difference()
	case Cmd_UpdatesDifferenceSlice:
		m2 := NewTLUpdatesDifferenceSlice(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates_Difference()
	case Cmd_UpdatesDifferenceTooLong:
		m2 := NewTLUpdatesDifferenceTooLong(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates_Difference()

	default:
		return fmt.Errorf("Updates_Difference Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Updates_State) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UpdatesState:
		m2 := NewTLUpdatesState(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Updates_State()

	default:
		return fmt.Errorf("Updates_State Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Upload_CdnFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UploadCdnFileReuploadNeeded:
		m2 := NewTLUploadCdnFileReuploadNeeded(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Upload_CdnFile()
	case Cmd_UploadCdnFile:
		m2 := NewTLUploadCdnFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Upload_CdnFile()

	default:
		return fmt.Errorf("Upload_CdnFile Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Upload_File) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UploadFile:
		m2 := NewTLUploadFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Upload_File()
	case Cmd_UploadFileCdnRedirect:
		m2 := NewTLUploadFileCdnRedirect(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Upload_File()

	case Cmd_UploadFileCdnRedirectf18cda44:
		m2 := NewTLUploadFileCdnRedirect(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Upload_File()

	default:
		return fmt.Errorf("Upload_File Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Upload_WebFile) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UploadWebFile:
		m2 := NewTLUploadWebFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Upload_WebFile()

	default:
		return fmt.Errorf("Upload_WebFile Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *UrlAuthResult) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UrlAuthResultRequest:
		m2 := NewTLUrlAuthResultRequest(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UrlAuthResult()
	case Cmd_UrlAuthResultAccepted:
		m2 := NewTLUrlAuthResultAccepted(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UrlAuthResult()
	case Cmd_UrlAuthResultDefault:
		m2 := NewTLUrlAuthResultDefault(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UrlAuthResult()

	default:
		return fmt.Errorf("UrlAuthResult Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *User) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UserEmpty:
		m2 := NewTLUserEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_User()
	case Cmd_User:
		m2 := NewTLUser(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_User()

	case Cmd_Userd10d979a:
		m2 := NewTLUser(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_User()
	case Cmd_User938458c1:
		m2 := NewTLUser(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_User()

	default:
		return fmt.Errorf("User Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *UserFull) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UserFull:
		m2 := NewTLUserFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserFull()

	case Cmd_UserFull5932fc03:
		m2 := NewTLUserFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserFull()
	case Cmd_UserFull8ea4a881:
		m2 := NewTLUserFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserFull()
	case Cmd_UserFull745559cc:
		m2 := NewTLUserFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserFull()
	case Cmd_UserFulledf17c12:
		m2 := NewTLUserFull(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserFull()

	default:
		return fmt.Errorf("UserFull Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *UserProfilePhoto) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UserProfilePhotoEmpty:
		m2 := NewTLUserProfilePhotoEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserProfilePhoto()
	case Cmd_UserProfilePhoto:
		m2 := NewTLUserProfilePhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserProfilePhoto()

	case Cmd_UserProfilePhotoecd75d8c:
		m2 := NewTLUserProfilePhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserProfilePhoto()
	case Cmd_UserProfilePhoto69d3ab26:
		m2 := NewTLUserProfilePhoto(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserProfilePhoto()

	default:
		return fmt.Errorf("UserProfilePhoto Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *UserStatus) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_UserStatusEmpty:
		m2 := NewTLUserStatusEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserStatus()
	case Cmd_UserStatusOnline:
		m2 := NewTLUserStatusOnline(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserStatus()
	case Cmd_UserStatusOffline:
		m2 := NewTLUserStatusOffline(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserStatus()
	case Cmd_UserStatusRecently:
		m2 := NewTLUserStatusRecently(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserStatus()
	case Cmd_UserStatusLastWeek:
		m2 := NewTLUserStatusLastWeek(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserStatus()
	case Cmd_UserStatusLastMonth:
		m2 := NewTLUserStatusLastMonth(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_UserStatus()

	default:
		return fmt.Errorf("UserStatus Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *VideoSize) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_VideoSize:
		m2 := NewTLVideoSize(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_VideoSize()

	case Cmd_VideoSizee831c556:
		m2 := NewTLVideoSize(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_VideoSize()

	default:
		return fmt.Errorf("VideoSize Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *WallPaper) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_WallPaper:
		m2 := NewTLWallPaper(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WallPaper()
	case Cmd_WallPaperSolid:
		m2 := NewTLWallPaperSolid(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WallPaper()
	case Cmd_WallPaperNoFile:
		m2 := NewTLWallPaperNoFile(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WallPaper()

	case Cmd_WallPaperf04f91ec:
		m2 := NewTLWallPaper(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WallPaper()
	case Cmd_WallPapera437c3ed:
		m2 := NewTLWallPaper(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WallPaper()

	default:
		return fmt.Errorf("WallPaper Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *WallPaperSettings) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_WallPaperSettings:
		m2 := NewTLWallPaperSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WallPaperSettings()

	case Cmd_WallPaperSettings5086cf8:
		m2 := NewTLWallPaperSettings(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WallPaperSettings()

	default:
		return fmt.Errorf("WallPaperSettings Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Wallet_KeySecretSalt) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_WalletSecretSalt:
		m2 := NewTLWalletSecretSalt(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Wallet_KeySecretSalt()

	default:
		return fmt.Errorf("Wallet_KeySecretSalt Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *Wallet_LiteResponse) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_WalletLiteResponse:
		m2 := NewTLWalletLiteResponse(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_Wallet_LiteResponse()

	default:
		return fmt.Errorf("Wallet_LiteResponse Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *WebAuthorization) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_WebAuthorization:
		m2 := NewTLWebAuthorization(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebAuthorization()

	default:
		return fmt.Errorf("WebAuthorization Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *WebDocument) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_WebDocument:
		m2 := NewTLWebDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebDocument()
	case Cmd_WebDocumentNoProxy:
		m2 := NewTLWebDocumentNoProxy(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebDocument()

	case Cmd_WebDocument1c570ed1:
		m2 := NewTLWebDocument(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebDocument()

	default:
		return fmt.Errorf("WebDocument Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *WebPage) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_WebPageEmpty:
		m2 := NewTLWebPageEmpty(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebPage()
	case Cmd_WebPagePending:
		m2 := NewTLWebPagePending(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebPage()
	case Cmd_WebPage:
		m2 := NewTLWebPage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebPage()
	case Cmd_WebPageNotModified:
		m2 := NewTLWebPageNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebPage()

	case Cmd_WebPageca820ed7:
		m2 := NewTLWebPage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebPage()
	case Cmd_WebPagefa64e172:
		m2 := NewTLWebPage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebPage()
	case Cmd_WebPagee89c45b2:
		m2 := NewTLWebPage(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebPage()
	case Cmd_WebPageNotModified7311ca11:
		m2 := NewTLWebPageNotModified(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebPage()

	default:
		return fmt.Errorf("WebPage Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}

func (m *WebPageAttribute) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	_ = err
	m.Cmd = TLCmd(dBuf.Int())
	switch m.Cmd {
	case Cmd_WebPageAttributeTheme:
		m2 := NewTLWebPageAttributeTheme(m)
		err = m2.Decode(dBuf, layer)
		if err != nil {
			return err
		}
		//m = m2.To_WebPageAttribute()

	default:
		return fmt.Errorf("WebPageAttribute Decode Invalid cmd: %x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}
