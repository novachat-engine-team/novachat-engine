/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2023/06/12 18:03
 * @File : codec_schema_maker.pb.go
 * @Desc :
 *
 */

package mtproto

func NewTLAccessPointRule(v *AccessPointRule) *TLAccessPointRule {
	if v == nil {
		return &TLAccessPointRule{
			Data2: &AccessPointRule{
				Cmd:       Cmd_AccessPointRule,
				ClassName: ClassAccessPointRule,
			},
		}
	}
	v1 := &TLAccessPointRule{Data2: v}
	v1.Data2.ClassName = ClassAccessPointRule
	return v1
}

func NewTLAccountDaysTTL(v *AccountDaysTTL) *TLAccountDaysTTL {
	if v == nil {
		return &TLAccountDaysTTL{
			Data2: &AccountDaysTTL{
				Cmd:       Cmd_AccountDaysTTL,
				ClassName: ClassAccountDaysTTL,
			},
		}
	}
	v1 := &TLAccountDaysTTL{Data2: v}
	v1.Data2.ClassName = ClassAccountDaysTTL
	return v1
}

func NewTLAccountAuthorizationForm(v *Account_AuthorizationForm) *TLAccountAuthorizationForm {
	if v == nil {
		return &TLAccountAuthorizationForm{
			Data2: &Account_AuthorizationForm{
				Cmd:       Cmd_AccountAuthorizationForm,
				ClassName: ClassAccountAuthorizationForm,
			},
		}
	}
	v1 := &TLAccountAuthorizationForm{Data2: v}
	v1.Data2.ClassName = ClassAccountAuthorizationForm
	return v1
}

func NewTLAccountAuthorizations(v *Account_Authorizations) *TLAccountAuthorizations {
	if v == nil {
		return &TLAccountAuthorizations{
			Data2: &Account_Authorizations{
				Cmd:       Cmd_AccountAuthorizations,
				ClassName: ClassAccountAuthorizations,
			},
		}
	}
	v1 := &TLAccountAuthorizations{Data2: v}
	v1.Data2.ClassName = ClassAccountAuthorizations
	return v1
}

func NewTLAccountAutoDownloadSettings(v *Account_AutoDownloadSettings) *TLAccountAutoDownloadSettings {
	if v == nil {
		return &TLAccountAutoDownloadSettings{
			Data2: &Account_AutoDownloadSettings{
				Cmd:       Cmd_AccountAutoDownloadSettings,
				ClassName: ClassAccountAutoDownloadSettings,
			},
		}
	}
	v1 := &TLAccountAutoDownloadSettings{Data2: v}
	v1.Data2.ClassName = ClassAccountAutoDownloadSettings
	return v1
}

func NewTLAccountContentSettings(v *Account_ContentSettings) *TLAccountContentSettings {
	if v == nil {
		return &TLAccountContentSettings{
			Data2: &Account_ContentSettings{
				Cmd:       Cmd_AccountContentSettings,
				ClassName: ClassAccountContentSettings,
			},
		}
	}
	v1 := &TLAccountContentSettings{Data2: v}
	v1.Data2.ClassName = ClassAccountContentSettings
	return v1
}

func NewTLAccountNoPassword(v *Account_Password) *TLAccountNoPassword {
	if v == nil {
		return &TLAccountNoPassword{
			Data2: &Account_Password{
				Cmd:       Cmd_AccountNoPassword,
				ClassName: ClassAccountNoPassword,
			},
		}
	}
	v1 := &TLAccountNoPassword{Data2: v}
	v1.Data2.ClassName = ClassAccountNoPassword
	return v1
}

func NewTLAccountPassword(v *Account_Password) *TLAccountPassword {
	if v == nil {
		return &TLAccountPassword{
			Data2: &Account_Password{
				Cmd:       Cmd_AccountPassword,
				ClassName: ClassAccountPassword,
			},
		}
	}
	v1 := &TLAccountPassword{Data2: v}
	v1.Data2.ClassName = ClassAccountPassword
	return v1
}

func NewTLAccountPasswordInputSettings(v *Account_PasswordInputSettings) *TLAccountPasswordInputSettings {
	if v == nil {
		return &TLAccountPasswordInputSettings{
			Data2: &Account_PasswordInputSettings{
				Cmd:       Cmd_AccountPasswordInputSettings,
				ClassName: ClassAccountPasswordInputSettings,
			},
		}
	}
	v1 := &TLAccountPasswordInputSettings{Data2: v}
	v1.Data2.ClassName = ClassAccountPasswordInputSettings
	return v1
}

func NewTLAccountPasswordSettings(v *Account_PasswordSettings) *TLAccountPasswordSettings {
	if v == nil {
		return &TLAccountPasswordSettings{
			Data2: &Account_PasswordSettings{
				Cmd:       Cmd_AccountPasswordSettings,
				ClassName: ClassAccountPasswordSettings,
			},
		}
	}
	v1 := &TLAccountPasswordSettings{Data2: v}
	v1.Data2.ClassName = ClassAccountPasswordSettings
	return v1
}

func NewTLAccountPrivacyRules(v *Account_PrivacyRules) *TLAccountPrivacyRules {
	if v == nil {
		return &TLAccountPrivacyRules{
			Data2: &Account_PrivacyRules{
				Cmd:       Cmd_AccountPrivacyRules,
				ClassName: ClassAccountPrivacyRules,
			},
		}
	}
	v1 := &TLAccountPrivacyRules{Data2: v}
	v1.Data2.ClassName = ClassAccountPrivacyRules
	return v1
}

func NewTLAccountSentEmailCode(v *Account_SentEmailCode) *TLAccountSentEmailCode {
	if v == nil {
		return &TLAccountSentEmailCode{
			Data2: &Account_SentEmailCode{
				Cmd:       Cmd_AccountSentEmailCode,
				ClassName: ClassAccountSentEmailCode,
			},
		}
	}
	v1 := &TLAccountSentEmailCode{Data2: v}
	v1.Data2.ClassName = ClassAccountSentEmailCode
	return v1
}

func NewTLAccountTakeout(v *Account_Takeout) *TLAccountTakeout {
	if v == nil {
		return &TLAccountTakeout{
			Data2: &Account_Takeout{
				Cmd:       Cmd_AccountTakeout,
				ClassName: ClassAccountTakeout,
			},
		}
	}
	v1 := &TLAccountTakeout{Data2: v}
	v1.Data2.ClassName = ClassAccountTakeout
	return v1
}

func NewTLAccountThemesNotModified(v *Account_Themes) *TLAccountThemesNotModified {
	if v == nil {
		return &TLAccountThemesNotModified{
			Data2: &Account_Themes{
				Cmd:       Cmd_AccountThemesNotModified,
				ClassName: ClassAccountThemesNotModified,
			},
		}
	}
	v1 := &TLAccountThemesNotModified{Data2: v}
	v1.Data2.ClassName = ClassAccountThemesNotModified
	return v1
}

func NewTLAccountThemes(v *Account_Themes) *TLAccountThemes {
	if v == nil {
		return &TLAccountThemes{
			Data2: &Account_Themes{
				Cmd:       Cmd_AccountThemes,
				ClassName: ClassAccountThemes,
			},
		}
	}
	v1 := &TLAccountThemes{Data2: v}
	v1.Data2.ClassName = ClassAccountThemes
	return v1
}

func NewTLAccountTmpPassword(v *Account_TmpPassword) *TLAccountTmpPassword {
	if v == nil {
		return &TLAccountTmpPassword{
			Data2: &Account_TmpPassword{
				Cmd:       Cmd_AccountTmpPassword,
				ClassName: ClassAccountTmpPassword,
			},
		}
	}
	v1 := &TLAccountTmpPassword{Data2: v}
	v1.Data2.ClassName = ClassAccountTmpPassword
	return v1
}

func NewTLAccountWallPapersNotModified(v *Account_WallPapers) *TLAccountWallPapersNotModified {
	if v == nil {
		return &TLAccountWallPapersNotModified{
			Data2: &Account_WallPapers{
				Cmd:       Cmd_AccountWallPapersNotModified,
				ClassName: ClassAccountWallPapersNotModified,
			},
		}
	}
	v1 := &TLAccountWallPapersNotModified{Data2: v}
	v1.Data2.ClassName = ClassAccountWallPapersNotModified
	return v1
}

func NewTLAccountWallPapers(v *Account_WallPapers) *TLAccountWallPapers {
	if v == nil {
		return &TLAccountWallPapers{
			Data2: &Account_WallPapers{
				Cmd:       Cmd_AccountWallPapers,
				ClassName: ClassAccountWallPapers,
			},
		}
	}
	v1 := &TLAccountWallPapers{Data2: v}
	v1.Data2.ClassName = ClassAccountWallPapers
	return v1
}

func NewTLAccountWebAuthorizations(v *Account_WebAuthorizations) *TLAccountWebAuthorizations {
	if v == nil {
		return &TLAccountWebAuthorizations{
			Data2: &Account_WebAuthorizations{
				Cmd:       Cmd_AccountWebAuthorizations,
				ClassName: ClassAccountWebAuthorizations,
			},
		}
	}
	v1 := &TLAccountWebAuthorizations{Data2: v}
	v1.Data2.ClassName = ClassAccountWebAuthorizations
	return v1
}

func NewTLAuthAuthorization(v *Auth_Authorization) *TLAuthAuthorization {
	if v == nil {
		return &TLAuthAuthorization{
			Data2: &Auth_Authorization{
				Cmd:       Cmd_AuthAuthorization,
				ClassName: ClassAuthAuthorization,
			},
		}
	}
	v1 := &TLAuthAuthorization{Data2: v}
	v1.Data2.ClassName = ClassAuthAuthorization
	return v1
}

func NewTLAuthAuthorizationSignUpRequired(v *Auth_Authorization) *TLAuthAuthorizationSignUpRequired {
	if v == nil {
		return &TLAuthAuthorizationSignUpRequired{
			Data2: &Auth_Authorization{
				Cmd:       Cmd_AuthAuthorizationSignUpRequired,
				ClassName: ClassAuthAuthorizationSignUpRequired,
			},
		}
	}
	v1 := &TLAuthAuthorizationSignUpRequired{Data2: v}
	v1.Data2.ClassName = ClassAuthAuthorizationSignUpRequired
	return v1
}

func NewTLAuthCheckedPhone(v *Auth_CheckedPhone) *TLAuthCheckedPhone {
	if v == nil {
		return &TLAuthCheckedPhone{
			Data2: &Auth_CheckedPhone{
				Cmd:       Cmd_AuthCheckedPhone,
				ClassName: ClassAuthCheckedPhone,
			},
		}
	}
	v1 := &TLAuthCheckedPhone{Data2: v}
	v1.Data2.ClassName = ClassAuthCheckedPhone
	return v1
}

func NewTLAuthCodeTypeSms(v *Auth_CodeType) *TLAuthCodeTypeSms {
	if v == nil {
		return &TLAuthCodeTypeSms{
			Data2: &Auth_CodeType{
				Cmd:       Cmd_AuthCodeTypeSms,
				ClassName: ClassAuthCodeTypeSms,
			},
		}
	}
	v1 := &TLAuthCodeTypeSms{Data2: v}
	v1.Data2.ClassName = ClassAuthCodeTypeSms
	return v1
}

func NewTLAuthCodeTypeCall(v *Auth_CodeType) *TLAuthCodeTypeCall {
	if v == nil {
		return &TLAuthCodeTypeCall{
			Data2: &Auth_CodeType{
				Cmd:       Cmd_AuthCodeTypeCall,
				ClassName: ClassAuthCodeTypeCall,
			},
		}
	}
	v1 := &TLAuthCodeTypeCall{Data2: v}
	v1.Data2.ClassName = ClassAuthCodeTypeCall
	return v1
}

func NewTLAuthCodeTypeFlashCall(v *Auth_CodeType) *TLAuthCodeTypeFlashCall {
	if v == nil {
		return &TLAuthCodeTypeFlashCall{
			Data2: &Auth_CodeType{
				Cmd:       Cmd_AuthCodeTypeFlashCall,
				ClassName: ClassAuthCodeTypeFlashCall,
			},
		}
	}
	v1 := &TLAuthCodeTypeFlashCall{Data2: v}
	v1.Data2.ClassName = ClassAuthCodeTypeFlashCall
	return v1
}

func NewTLAuthExportedAuthorization(v *Auth_ExportedAuthorization) *TLAuthExportedAuthorization {
	if v == nil {
		return &TLAuthExportedAuthorization{
			Data2: &Auth_ExportedAuthorization{
				Cmd:       Cmd_AuthExportedAuthorization,
				ClassName: ClassAuthExportedAuthorization,
			},
		}
	}
	v1 := &TLAuthExportedAuthorization{Data2: v}
	v1.Data2.ClassName = ClassAuthExportedAuthorization
	return v1
}

func NewTLAuthLoginToken(v *Auth_LoginToken) *TLAuthLoginToken {
	if v == nil {
		return &TLAuthLoginToken{
			Data2: &Auth_LoginToken{
				Cmd:       Cmd_AuthLoginToken,
				ClassName: ClassAuthLoginToken,
			},
		}
	}
	v1 := &TLAuthLoginToken{Data2: v}
	v1.Data2.ClassName = ClassAuthLoginToken
	return v1
}

func NewTLAuthLoginTokenMigrateTo(v *Auth_LoginToken) *TLAuthLoginTokenMigrateTo {
	if v == nil {
		return &TLAuthLoginTokenMigrateTo{
			Data2: &Auth_LoginToken{
				Cmd:       Cmd_AuthLoginTokenMigrateTo,
				ClassName: ClassAuthLoginTokenMigrateTo,
			},
		}
	}
	v1 := &TLAuthLoginTokenMigrateTo{Data2: v}
	v1.Data2.ClassName = ClassAuthLoginTokenMigrateTo
	return v1
}

func NewTLAuthLoginTokenSuccess(v *Auth_LoginToken) *TLAuthLoginTokenSuccess {
	if v == nil {
		return &TLAuthLoginTokenSuccess{
			Data2: &Auth_LoginToken{
				Cmd:       Cmd_AuthLoginTokenSuccess,
				ClassName: ClassAuthLoginTokenSuccess,
			},
		}
	}
	v1 := &TLAuthLoginTokenSuccess{Data2: v}
	v1.Data2.ClassName = ClassAuthLoginTokenSuccess
	return v1
}

func NewTLAuthPasswordRecovery(v *Auth_PasswordRecovery) *TLAuthPasswordRecovery {
	if v == nil {
		return &TLAuthPasswordRecovery{
			Data2: &Auth_PasswordRecovery{
				Cmd:       Cmd_AuthPasswordRecovery,
				ClassName: ClassAuthPasswordRecovery,
			},
		}
	}
	v1 := &TLAuthPasswordRecovery{Data2: v}
	v1.Data2.ClassName = ClassAuthPasswordRecovery
	return v1
}

func NewTLAuthSentCode(v *Auth_SentCode) *TLAuthSentCode {
	if v == nil {
		return &TLAuthSentCode{
			Data2: &Auth_SentCode{
				Cmd:       Cmd_AuthSentCode,
				ClassName: ClassAuthSentCode,
			},
		}
	}
	v1 := &TLAuthSentCode{Data2: v}
	v1.Data2.ClassName = ClassAuthSentCode
	return v1
}

func NewTLAuthSentCodeTypeApp(v *Auth_SentCodeType) *TLAuthSentCodeTypeApp {
	if v == nil {
		return &TLAuthSentCodeTypeApp{
			Data2: &Auth_SentCodeType{
				Cmd:       Cmd_AuthSentCodeTypeApp,
				ClassName: ClassAuthSentCodeTypeApp,
			},
		}
	}
	v1 := &TLAuthSentCodeTypeApp{Data2: v}
	v1.Data2.ClassName = ClassAuthSentCodeTypeApp
	return v1
}

func NewTLAuthSentCodeTypeSms(v *Auth_SentCodeType) *TLAuthSentCodeTypeSms {
	if v == nil {
		return &TLAuthSentCodeTypeSms{
			Data2: &Auth_SentCodeType{
				Cmd:       Cmd_AuthSentCodeTypeSms,
				ClassName: ClassAuthSentCodeTypeSms,
			},
		}
	}
	v1 := &TLAuthSentCodeTypeSms{Data2: v}
	v1.Data2.ClassName = ClassAuthSentCodeTypeSms
	return v1
}

func NewTLAuthSentCodeTypeCall(v *Auth_SentCodeType) *TLAuthSentCodeTypeCall {
	if v == nil {
		return &TLAuthSentCodeTypeCall{
			Data2: &Auth_SentCodeType{
				Cmd:       Cmd_AuthSentCodeTypeCall,
				ClassName: ClassAuthSentCodeTypeCall,
			},
		}
	}
	v1 := &TLAuthSentCodeTypeCall{Data2: v}
	v1.Data2.ClassName = ClassAuthSentCodeTypeCall
	return v1
}

func NewTLAuthSentCodeTypeFlashCall(v *Auth_SentCodeType) *TLAuthSentCodeTypeFlashCall {
	if v == nil {
		return &TLAuthSentCodeTypeFlashCall{
			Data2: &Auth_SentCodeType{
				Cmd:       Cmd_AuthSentCodeTypeFlashCall,
				ClassName: ClassAuthSentCodeTypeFlashCall,
			},
		}
	}
	v1 := &TLAuthSentCodeTypeFlashCall{Data2: v}
	v1.Data2.ClassName = ClassAuthSentCodeTypeFlashCall
	return v1
}

func NewTLAuthorization(v *Authorization) *TLAuthorization {
	if v == nil {
		return &TLAuthorization{
			Data2: &Authorization{
				Cmd:       Cmd_Authorization,
				ClassName: ClassAuthorization,
			},
		}
	}
	v1 := &TLAuthorization{Data2: v}
	v1.Data2.ClassName = ClassAuthorization
	return v1
}

func NewTLAutoDownloadSettings(v *AutoDownloadSettings) *TLAutoDownloadSettings {
	if v == nil {
		return &TLAutoDownloadSettings{
			Data2: &AutoDownloadSettings{
				Cmd:       Cmd_AutoDownloadSettings,
				ClassName: ClassAutoDownloadSettings,
			},
		}
	}
	v1 := &TLAutoDownloadSettings{Data2: v}
	v1.Data2.ClassName = ClassAutoDownloadSettings
	return v1
}

func NewTLBadMsgNotification(v *BadMsgNotification) *TLBadMsgNotification {
	if v == nil {
		return &TLBadMsgNotification{
			Data2: &BadMsgNotification{
				Cmd:       Cmd_BadMsgNotification,
				ClassName: ClassBadMsgNotification,
			},
		}
	}
	v1 := &TLBadMsgNotification{Data2: v}
	v1.Data2.ClassName = ClassBadMsgNotification
	return v1
}

func NewTLBadServerSalt(v *BadMsgNotification) *TLBadServerSalt {
	if v == nil {
		return &TLBadServerSalt{
			Data2: &BadMsgNotification{
				Cmd:       Cmd_BadServerSalt,
				ClassName: ClassBadServerSalt,
			},
		}
	}
	v1 := &TLBadServerSalt{Data2: v}
	v1.Data2.ClassName = ClassBadServerSalt
	return v1
}

func NewTLBankCardOpenUrl(v *BankCardOpenUrl) *TLBankCardOpenUrl {
	if v == nil {
		return &TLBankCardOpenUrl{
			Data2: &BankCardOpenUrl{
				Cmd:       Cmd_BankCardOpenUrl,
				ClassName: ClassBankCardOpenUrl,
			},
		}
	}
	v1 := &TLBankCardOpenUrl{Data2: v}
	v1.Data2.ClassName = ClassBankCardOpenUrl
	return v1
}

func NewTLBaseThemeClassic(v *BaseTheme) *TLBaseThemeClassic {
	if v == nil {
		return &TLBaseThemeClassic{
			Data2: &BaseTheme{
				Cmd:       Cmd_BaseThemeClassic,
				ClassName: ClassBaseThemeClassic,
			},
		}
	}
	v1 := &TLBaseThemeClassic{Data2: v}
	v1.Data2.ClassName = ClassBaseThemeClassic
	return v1
}

func NewTLBaseThemeDay(v *BaseTheme) *TLBaseThemeDay {
	if v == nil {
		return &TLBaseThemeDay{
			Data2: &BaseTheme{
				Cmd:       Cmd_BaseThemeDay,
				ClassName: ClassBaseThemeDay,
			},
		}
	}
	v1 := &TLBaseThemeDay{Data2: v}
	v1.Data2.ClassName = ClassBaseThemeDay
	return v1
}

func NewTLBaseThemeNight(v *BaseTheme) *TLBaseThemeNight {
	if v == nil {
		return &TLBaseThemeNight{
			Data2: &BaseTheme{
				Cmd:       Cmd_BaseThemeNight,
				ClassName: ClassBaseThemeNight,
			},
		}
	}
	v1 := &TLBaseThemeNight{Data2: v}
	v1.Data2.ClassName = ClassBaseThemeNight
	return v1
}

func NewTLBaseThemeTinted(v *BaseTheme) *TLBaseThemeTinted {
	if v == nil {
		return &TLBaseThemeTinted{
			Data2: &BaseTheme{
				Cmd:       Cmd_BaseThemeTinted,
				ClassName: ClassBaseThemeTinted,
			},
		}
	}
	v1 := &TLBaseThemeTinted{Data2: v}
	v1.Data2.ClassName = ClassBaseThemeTinted
	return v1
}

func NewTLBaseThemeArctic(v *BaseTheme) *TLBaseThemeArctic {
	if v == nil {
		return &TLBaseThemeArctic{
			Data2: &BaseTheme{
				Cmd:       Cmd_BaseThemeArctic,
				ClassName: ClassBaseThemeArctic,
			},
		}
	}
	v1 := &TLBaseThemeArctic{Data2: v}
	v1.Data2.ClassName = ClassBaseThemeArctic
	return v1
}

func NewTLBindAuthKeyInner(v *BindAuthKeyInner) *TLBindAuthKeyInner {
	if v == nil {
		return &TLBindAuthKeyInner{
			Data2: &BindAuthKeyInner{
				Cmd:       Cmd_BindAuthKeyInner,
				ClassName: ClassBindAuthKeyInner,
			},
		}
	}
	v1 := &TLBindAuthKeyInner{Data2: v}
	v1.Data2.ClassName = ClassBindAuthKeyInner
	return v1
}

func NewTLBoolFalse(v *Bool) *TLBoolFalse {
	if v == nil {
		return &TLBoolFalse{
			Data2: &Bool{
				Cmd:       Cmd_BoolFalse,
				ClassName: ClassBoolFalse,
			},
		}
	}
	v1 := &TLBoolFalse{Data2: v}
	v1.Data2.ClassName = ClassBoolFalse
	return v1
}

func NewTLBoolTrue(v *Bool) *TLBoolTrue {
	if v == nil {
		return &TLBoolTrue{
			Data2: &Bool{
				Cmd:       Cmd_BoolTrue,
				ClassName: ClassBoolTrue,
			},
		}
	}
	v1 := &TLBoolTrue{Data2: v}
	v1.Data2.ClassName = ClassBoolTrue
	return v1
}

func NewTLBotCommand(v *BotCommand) *TLBotCommand {
	if v == nil {
		return &TLBotCommand{
			Data2: &BotCommand{
				Cmd:       Cmd_BotCommand,
				ClassName: ClassBotCommand,
			},
		}
	}
	v1 := &TLBotCommand{Data2: v}
	v1.Data2.ClassName = ClassBotCommand
	return v1
}

func NewTLBotInfo(v *BotInfo) *TLBotInfo {
	if v == nil {
		return &TLBotInfo{
			Data2: &BotInfo{
				Cmd:       Cmd_BotInfo,
				ClassName: ClassBotInfo,
			},
		}
	}
	v1 := &TLBotInfo{Data2: v}
	v1.Data2.ClassName = ClassBotInfo
	return v1
}

func NewTLBotInlineMessageMediaAuto(v *BotInlineMessage) *TLBotInlineMessageMediaAuto {
	if v == nil {
		return &TLBotInlineMessageMediaAuto{
			Data2: &BotInlineMessage{
				Cmd:       Cmd_BotInlineMessageMediaAuto,
				ClassName: ClassBotInlineMessageMediaAuto,
			},
		}
	}
	v1 := &TLBotInlineMessageMediaAuto{Data2: v}
	v1.Data2.ClassName = ClassBotInlineMessageMediaAuto
	return v1
}

func NewTLBotInlineMessageText(v *BotInlineMessage) *TLBotInlineMessageText {
	if v == nil {
		return &TLBotInlineMessageText{
			Data2: &BotInlineMessage{
				Cmd:       Cmd_BotInlineMessageText,
				ClassName: ClassBotInlineMessageText,
			},
		}
	}
	v1 := &TLBotInlineMessageText{Data2: v}
	v1.Data2.ClassName = ClassBotInlineMessageText
	return v1
}

func NewTLBotInlineMessageMediaGeo(v *BotInlineMessage) *TLBotInlineMessageMediaGeo {
	if v == nil {
		return &TLBotInlineMessageMediaGeo{
			Data2: &BotInlineMessage{
				Cmd:       Cmd_BotInlineMessageMediaGeo,
				ClassName: ClassBotInlineMessageMediaGeo,
			},
		}
	}
	v1 := &TLBotInlineMessageMediaGeo{Data2: v}
	v1.Data2.ClassName = ClassBotInlineMessageMediaGeo
	return v1
}

func NewTLBotInlineMessageMediaVenue(v *BotInlineMessage) *TLBotInlineMessageMediaVenue {
	if v == nil {
		return &TLBotInlineMessageMediaVenue{
			Data2: &BotInlineMessage{
				Cmd:       Cmd_BotInlineMessageMediaVenue,
				ClassName: ClassBotInlineMessageMediaVenue,
			},
		}
	}
	v1 := &TLBotInlineMessageMediaVenue{Data2: v}
	v1.Data2.ClassName = ClassBotInlineMessageMediaVenue
	return v1
}

func NewTLBotInlineMessageMediaContact(v *BotInlineMessage) *TLBotInlineMessageMediaContact {
	if v == nil {
		return &TLBotInlineMessageMediaContact{
			Data2: &BotInlineMessage{
				Cmd:       Cmd_BotInlineMessageMediaContact,
				ClassName: ClassBotInlineMessageMediaContact,
			},
		}
	}
	v1 := &TLBotInlineMessageMediaContact{Data2: v}
	v1.Data2.ClassName = ClassBotInlineMessageMediaContact
	return v1
}

func NewTLBotInlineResult(v *BotInlineResult) *TLBotInlineResult {
	if v == nil {
		return &TLBotInlineResult{
			Data2: &BotInlineResult{
				Cmd:       Cmd_BotInlineResult,
				ClassName: ClassBotInlineResult,
			},
		}
	}
	v1 := &TLBotInlineResult{Data2: v}
	v1.Data2.ClassName = ClassBotInlineResult
	return v1
}

func NewTLBotInlineMediaResult(v *BotInlineResult) *TLBotInlineMediaResult {
	if v == nil {
		return &TLBotInlineMediaResult{
			Data2: &BotInlineResult{
				Cmd:       Cmd_BotInlineMediaResult,
				ClassName: ClassBotInlineMediaResult,
			},
		}
	}
	v1 := &TLBotInlineMediaResult{Data2: v}
	v1.Data2.ClassName = ClassBotInlineMediaResult
	return v1
}

func NewTLCdnConfig(v *CdnConfig) *TLCdnConfig {
	if v == nil {
		return &TLCdnConfig{
			Data2: &CdnConfig{
				Cmd:       Cmd_CdnConfig,
				ClassName: ClassCdnConfig,
			},
		}
	}
	v1 := &TLCdnConfig{Data2: v}
	v1.Data2.ClassName = ClassCdnConfig
	return v1
}

func NewTLCdnFileHash(v *CdnFileHash) *TLCdnFileHash {
	if v == nil {
		return &TLCdnFileHash{
			Data2: &CdnFileHash{
				Cmd:       Cmd_CdnFileHash,
				ClassName: ClassCdnFileHash,
			},
		}
	}
	v1 := &TLCdnFileHash{Data2: v}
	v1.Data2.ClassName = ClassCdnFileHash
	return v1
}

func NewTLCdnPublicKey(v *CdnPublicKey) *TLCdnPublicKey {
	if v == nil {
		return &TLCdnPublicKey{
			Data2: &CdnPublicKey{
				Cmd:       Cmd_CdnPublicKey,
				ClassName: ClassCdnPublicKey,
			},
		}
	}
	v1 := &TLCdnPublicKey{Data2: v}
	v1.Data2.ClassName = ClassCdnPublicKey
	return v1
}

func NewTLChannelAdminLogEvent(v *ChannelAdminLogEvent) *TLChannelAdminLogEvent {
	if v == nil {
		return &TLChannelAdminLogEvent{
			Data2: &ChannelAdminLogEvent{
				Cmd:       Cmd_ChannelAdminLogEvent,
				ClassName: ClassChannelAdminLogEvent,
			},
		}
	}
	v1 := &TLChannelAdminLogEvent{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEvent
	return v1
}

func NewTLChannelAdminLogEventActionChangeTitle(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionChangeTitle {
	if v == nil {
		return &TLChannelAdminLogEventActionChangeTitle{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionChangeTitle,
				ClassName: ClassChannelAdminLogEventActionChangeTitle,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionChangeTitle{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionChangeTitle
	return v1
}

func NewTLChannelAdminLogEventActionChangeAbout(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionChangeAbout {
	if v == nil {
		return &TLChannelAdminLogEventActionChangeAbout{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionChangeAbout,
				ClassName: ClassChannelAdminLogEventActionChangeAbout,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionChangeAbout{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionChangeAbout
	return v1
}

func NewTLChannelAdminLogEventActionChangeUsername(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionChangeUsername {
	if v == nil {
		return &TLChannelAdminLogEventActionChangeUsername{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionChangeUsername,
				ClassName: ClassChannelAdminLogEventActionChangeUsername,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionChangeUsername{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionChangeUsername
	return v1
}

func NewTLChannelAdminLogEventActionChangePhoto(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionChangePhoto {
	if v == nil {
		return &TLChannelAdminLogEventActionChangePhoto{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionChangePhoto,
				ClassName: ClassChannelAdminLogEventActionChangePhoto,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionChangePhoto{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionChangePhoto
	return v1
}

func NewTLChannelAdminLogEventActionToggleInvites(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionToggleInvites {
	if v == nil {
		return &TLChannelAdminLogEventActionToggleInvites{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionToggleInvites,
				ClassName: ClassChannelAdminLogEventActionToggleInvites,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionToggleInvites{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionToggleInvites
	return v1
}

func NewTLChannelAdminLogEventActionToggleSignatures(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionToggleSignatures {
	if v == nil {
		return &TLChannelAdminLogEventActionToggleSignatures{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionToggleSignatures,
				ClassName: ClassChannelAdminLogEventActionToggleSignatures,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionToggleSignatures{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionToggleSignatures
	return v1
}

func NewTLChannelAdminLogEventActionUpdatePinned(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionUpdatePinned {
	if v == nil {
		return &TLChannelAdminLogEventActionUpdatePinned{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionUpdatePinned,
				ClassName: ClassChannelAdminLogEventActionUpdatePinned,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionUpdatePinned{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionUpdatePinned
	return v1
}

func NewTLChannelAdminLogEventActionEditMessage(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionEditMessage {
	if v == nil {
		return &TLChannelAdminLogEventActionEditMessage{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionEditMessage,
				ClassName: ClassChannelAdminLogEventActionEditMessage,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionEditMessage{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionEditMessage
	return v1
}

func NewTLChannelAdminLogEventActionDeleteMessage(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionDeleteMessage {
	if v == nil {
		return &TLChannelAdminLogEventActionDeleteMessage{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionDeleteMessage,
				ClassName: ClassChannelAdminLogEventActionDeleteMessage,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionDeleteMessage{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionDeleteMessage
	return v1
}

func NewTLChannelAdminLogEventActionParticipantJoin(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionParticipantJoin {
	if v == nil {
		return &TLChannelAdminLogEventActionParticipantJoin{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionParticipantJoin,
				ClassName: ClassChannelAdminLogEventActionParticipantJoin,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionParticipantJoin{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionParticipantJoin
	return v1
}

func NewTLChannelAdminLogEventActionParticipantLeave(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionParticipantLeave {
	if v == nil {
		return &TLChannelAdminLogEventActionParticipantLeave{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionParticipantLeave,
				ClassName: ClassChannelAdminLogEventActionParticipantLeave,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionParticipantLeave{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionParticipantLeave
	return v1
}

func NewTLChannelAdminLogEventActionParticipantInvite(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionParticipantInvite {
	if v == nil {
		return &TLChannelAdminLogEventActionParticipantInvite{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionParticipantInvite,
				ClassName: ClassChannelAdminLogEventActionParticipantInvite,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionParticipantInvite{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionParticipantInvite
	return v1
}

func NewTLChannelAdminLogEventActionParticipantToggleBan(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionParticipantToggleBan {
	if v == nil {
		return &TLChannelAdminLogEventActionParticipantToggleBan{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionParticipantToggleBan,
				ClassName: ClassChannelAdminLogEventActionParticipantToggleBan,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionParticipantToggleBan{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionParticipantToggleBan
	return v1
}

func NewTLChannelAdminLogEventActionParticipantToggleAdmin(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionParticipantToggleAdmin {
	if v == nil {
		return &TLChannelAdminLogEventActionParticipantToggleAdmin{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionParticipantToggleAdmin,
				ClassName: ClassChannelAdminLogEventActionParticipantToggleAdmin,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionParticipantToggleAdmin{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionParticipantToggleAdmin
	return v1
}

func NewTLChannelAdminLogEventActionChangeStickerSet(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionChangeStickerSet {
	if v == nil {
		return &TLChannelAdminLogEventActionChangeStickerSet{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionChangeStickerSet,
				ClassName: ClassChannelAdminLogEventActionChangeStickerSet,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionChangeStickerSet{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionChangeStickerSet
	return v1
}

func NewTLChannelAdminLogEventActionTogglePreHistoryHidden(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionTogglePreHistoryHidden {
	if v == nil {
		return &TLChannelAdminLogEventActionTogglePreHistoryHidden{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionTogglePreHistoryHidden,
				ClassName: ClassChannelAdminLogEventActionTogglePreHistoryHidden,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionTogglePreHistoryHidden{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionTogglePreHistoryHidden
	return v1
}

func NewTLChannelAdminLogEventActionDefaultBannedRights(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionDefaultBannedRights {
	if v == nil {
		return &TLChannelAdminLogEventActionDefaultBannedRights{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionDefaultBannedRights,
				ClassName: ClassChannelAdminLogEventActionDefaultBannedRights,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionDefaultBannedRights{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionDefaultBannedRights
	return v1
}

func NewTLChannelAdminLogEventActionStopPoll(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionStopPoll {
	if v == nil {
		return &TLChannelAdminLogEventActionStopPoll{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionStopPoll,
				ClassName: ClassChannelAdminLogEventActionStopPoll,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionStopPoll{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionStopPoll
	return v1
}

func NewTLChannelAdminLogEventActionChangeLinkedChat(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionChangeLinkedChat {
	if v == nil {
		return &TLChannelAdminLogEventActionChangeLinkedChat{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionChangeLinkedChat,
				ClassName: ClassChannelAdminLogEventActionChangeLinkedChat,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionChangeLinkedChat{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionChangeLinkedChat
	return v1
}

func NewTLChannelAdminLogEventActionChangeLocation(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionChangeLocation {
	if v == nil {
		return &TLChannelAdminLogEventActionChangeLocation{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionChangeLocation,
				ClassName: ClassChannelAdminLogEventActionChangeLocation,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionChangeLocation{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionChangeLocation
	return v1
}

func NewTLChannelAdminLogEventActionToggleSlowMode(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionToggleSlowMode {
	if v == nil {
		return &TLChannelAdminLogEventActionToggleSlowMode{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionToggleSlowMode,
				ClassName: ClassChannelAdminLogEventActionToggleSlowMode,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionToggleSlowMode{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionToggleSlowMode
	return v1
}

func NewTLChannelAdminLogEventActionStartGroupCall(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionStartGroupCall {
	if v == nil {
		return &TLChannelAdminLogEventActionStartGroupCall{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionStartGroupCall,
				ClassName: ClassChannelAdminLogEventActionStartGroupCall,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionStartGroupCall{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionStartGroupCall
	return v1
}

func NewTLChannelAdminLogEventActionDiscardGroupCall(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionDiscardGroupCall {
	if v == nil {
		return &TLChannelAdminLogEventActionDiscardGroupCall{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionDiscardGroupCall,
				ClassName: ClassChannelAdminLogEventActionDiscardGroupCall,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionDiscardGroupCall{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionDiscardGroupCall
	return v1
}

func NewTLChannelAdminLogEventActionParticipantMute(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionParticipantMute {
	if v == nil {
		return &TLChannelAdminLogEventActionParticipantMute{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionParticipantMute,
				ClassName: ClassChannelAdminLogEventActionParticipantMute,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionParticipantMute{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionParticipantMute
	return v1
}

func NewTLChannelAdminLogEventActionParticipantUnmute(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionParticipantUnmute {
	if v == nil {
		return &TLChannelAdminLogEventActionParticipantUnmute{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionParticipantUnmute,
				ClassName: ClassChannelAdminLogEventActionParticipantUnmute,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionParticipantUnmute{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionParticipantUnmute
	return v1
}

func NewTLChannelAdminLogEventActionToggleGroupCallSetting(v *ChannelAdminLogEventAction) *TLChannelAdminLogEventActionToggleGroupCallSetting {
	if v == nil {
		return &TLChannelAdminLogEventActionToggleGroupCallSetting{
			Data2: &ChannelAdminLogEventAction{
				Cmd:       Cmd_ChannelAdminLogEventActionToggleGroupCallSetting,
				ClassName: ClassChannelAdminLogEventActionToggleGroupCallSetting,
			},
		}
	}
	v1 := &TLChannelAdminLogEventActionToggleGroupCallSetting{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventActionToggleGroupCallSetting
	return v1
}

func NewTLChannelAdminLogEventsFilter(v *ChannelAdminLogEventsFilter) *TLChannelAdminLogEventsFilter {
	if v == nil {
		return &TLChannelAdminLogEventsFilter{
			Data2: &ChannelAdminLogEventsFilter{
				Cmd:       Cmd_ChannelAdminLogEventsFilter,
				ClassName: ClassChannelAdminLogEventsFilter,
			},
		}
	}
	v1 := &TLChannelAdminLogEventsFilter{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminLogEventsFilter
	return v1
}

func NewTLChannelAdminRights(v *ChannelAdminRights) *TLChannelAdminRights {
	if v == nil {
		return &TLChannelAdminRights{
			Data2: &ChannelAdminRights{
				Cmd:       Cmd_ChannelAdminRights,
				ClassName: ClassChannelAdminRights,
			},
		}
	}
	v1 := &TLChannelAdminRights{Data2: v}
	v1.Data2.ClassName = ClassChannelAdminRights
	return v1
}

func NewTLChannelBannedRights(v *ChannelBannedRights) *TLChannelBannedRights {
	if v == nil {
		return &TLChannelBannedRights{
			Data2: &ChannelBannedRights{
				Cmd:       Cmd_ChannelBannedRights,
				ClassName: ClassChannelBannedRights,
			},
		}
	}
	v1 := &TLChannelBannedRights{Data2: v}
	v1.Data2.ClassName = ClassChannelBannedRights
	return v1
}

func NewTLChannelLocationEmpty(v *ChannelLocation) *TLChannelLocationEmpty {
	if v == nil {
		return &TLChannelLocationEmpty{
			Data2: &ChannelLocation{
				Cmd:       Cmd_ChannelLocationEmpty,
				ClassName: ClassChannelLocationEmpty,
			},
		}
	}
	v1 := &TLChannelLocationEmpty{Data2: v}
	v1.Data2.ClassName = ClassChannelLocationEmpty
	return v1
}

func NewTLChannelLocation(v *ChannelLocation) *TLChannelLocation {
	if v == nil {
		return &TLChannelLocation{
			Data2: &ChannelLocation{
				Cmd:       Cmd_ChannelLocation,
				ClassName: ClassChannelLocation,
			},
		}
	}
	v1 := &TLChannelLocation{Data2: v}
	v1.Data2.ClassName = ClassChannelLocation
	return v1
}

func NewTLChannelMessagesFilterEmpty(v *ChannelMessagesFilter) *TLChannelMessagesFilterEmpty {
	if v == nil {
		return &TLChannelMessagesFilterEmpty{
			Data2: &ChannelMessagesFilter{
				Cmd:       Cmd_ChannelMessagesFilterEmpty,
				ClassName: ClassChannelMessagesFilterEmpty,
			},
		}
	}
	v1 := &TLChannelMessagesFilterEmpty{Data2: v}
	v1.Data2.ClassName = ClassChannelMessagesFilterEmpty
	return v1
}

func NewTLChannelMessagesFilter(v *ChannelMessagesFilter) *TLChannelMessagesFilter {
	if v == nil {
		return &TLChannelMessagesFilter{
			Data2: &ChannelMessagesFilter{
				Cmd:       Cmd_ChannelMessagesFilter,
				ClassName: ClassChannelMessagesFilter,
			},
		}
	}
	v1 := &TLChannelMessagesFilter{Data2: v}
	v1.Data2.ClassName = ClassChannelMessagesFilter
	return v1
}

func NewTLChannelMessagesFilterCollapsed(v *ChannelMessagesFilter) *TLChannelMessagesFilterCollapsed {
	if v == nil {
		return &TLChannelMessagesFilterCollapsed{
			Data2: &ChannelMessagesFilter{
				Cmd:       Cmd_ChannelMessagesFilterCollapsed,
				ClassName: ClassChannelMessagesFilterCollapsed,
			},
		}
	}
	v1 := &TLChannelMessagesFilterCollapsed{Data2: v}
	v1.Data2.ClassName = ClassChannelMessagesFilterCollapsed
	return v1
}

func NewTLChannelParticipant(v *ChannelParticipant) *TLChannelParticipant {
	if v == nil {
		return &TLChannelParticipant{
			Data2: &ChannelParticipant{
				Cmd:       Cmd_ChannelParticipant,
				ClassName: ClassChannelParticipant,
			},
		}
	}
	v1 := &TLChannelParticipant{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipant
	return v1
}

func NewTLChannelParticipantSelf(v *ChannelParticipant) *TLChannelParticipantSelf {
	if v == nil {
		return &TLChannelParticipantSelf{
			Data2: &ChannelParticipant{
				Cmd:       Cmd_ChannelParticipantSelf,
				ClassName: ClassChannelParticipantSelf,
			},
		}
	}
	v1 := &TLChannelParticipantSelf{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantSelf
	return v1
}

func NewTLChannelParticipantCreator(v *ChannelParticipant) *TLChannelParticipantCreator {
	if v == nil {
		return &TLChannelParticipantCreator{
			Data2: &ChannelParticipant{
				Cmd:       Cmd_ChannelParticipantCreator,
				ClassName: ClassChannelParticipantCreator,
			},
		}
	}
	v1 := &TLChannelParticipantCreator{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantCreator
	return v1
}

func NewTLChannelParticipantAdmin(v *ChannelParticipant) *TLChannelParticipantAdmin {
	if v == nil {
		return &TLChannelParticipantAdmin{
			Data2: &ChannelParticipant{
				Cmd:       Cmd_ChannelParticipantAdmin,
				ClassName: ClassChannelParticipantAdmin,
			},
		}
	}
	v1 := &TLChannelParticipantAdmin{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantAdmin
	return v1
}

func NewTLChannelParticipantBanned(v *ChannelParticipant) *TLChannelParticipantBanned {
	if v == nil {
		return &TLChannelParticipantBanned{
			Data2: &ChannelParticipant{
				Cmd:       Cmd_ChannelParticipantBanned,
				ClassName: ClassChannelParticipantBanned,
			},
		}
	}
	v1 := &TLChannelParticipantBanned{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantBanned
	return v1
}

func NewTLChannelParticipantModerator(v *ChannelParticipant) *TLChannelParticipantModerator {
	if v == nil {
		return &TLChannelParticipantModerator{
			Data2: &ChannelParticipant{
				Cmd:       Cmd_ChannelParticipantModerator,
				ClassName: ClassChannelParticipantModerator,
			},
		}
	}
	v1 := &TLChannelParticipantModerator{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantModerator
	return v1
}

func NewTLChannelParticipantEditor(v *ChannelParticipant) *TLChannelParticipantEditor {
	if v == nil {
		return &TLChannelParticipantEditor{
			Data2: &ChannelParticipant{
				Cmd:       Cmd_ChannelParticipantEditor,
				ClassName: ClassChannelParticipantEditor,
			},
		}
	}
	v1 := &TLChannelParticipantEditor{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantEditor
	return v1
}

func NewTLChannelParticipantKicked(v *ChannelParticipant) *TLChannelParticipantKicked {
	if v == nil {
		return &TLChannelParticipantKicked{
			Data2: &ChannelParticipant{
				Cmd:       Cmd_ChannelParticipantKicked,
				ClassName: ClassChannelParticipantKicked,
			},
		}
	}
	v1 := &TLChannelParticipantKicked{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantKicked
	return v1
}

func NewTLChannelParticipantLeft(v *ChannelParticipant) *TLChannelParticipantLeft {
	if v == nil {
		return &TLChannelParticipantLeft{
			Data2: &ChannelParticipant{
				Cmd:       Cmd_ChannelParticipantLeft,
				ClassName: ClassChannelParticipantLeft,
			},
		}
	}
	v1 := &TLChannelParticipantLeft{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantLeft
	return v1
}

func NewTLChannelRoleEmpty(v *ChannelParticipantRole) *TLChannelRoleEmpty {
	if v == nil {
		return &TLChannelRoleEmpty{
			Data2: &ChannelParticipantRole{
				Cmd:       Cmd_ChannelRoleEmpty,
				ClassName: ClassChannelRoleEmpty,
			},
		}
	}
	v1 := &TLChannelRoleEmpty{Data2: v}
	v1.Data2.ClassName = ClassChannelRoleEmpty
	return v1
}

func NewTLChannelRoleModerator(v *ChannelParticipantRole) *TLChannelRoleModerator {
	if v == nil {
		return &TLChannelRoleModerator{
			Data2: &ChannelParticipantRole{
				Cmd:       Cmd_ChannelRoleModerator,
				ClassName: ClassChannelRoleModerator,
			},
		}
	}
	v1 := &TLChannelRoleModerator{Data2: v}
	v1.Data2.ClassName = ClassChannelRoleModerator
	return v1
}

func NewTLChannelRoleEditor(v *ChannelParticipantRole) *TLChannelRoleEditor {
	if v == nil {
		return &TLChannelRoleEditor{
			Data2: &ChannelParticipantRole{
				Cmd:       Cmd_ChannelRoleEditor,
				ClassName: ClassChannelRoleEditor,
			},
		}
	}
	v1 := &TLChannelRoleEditor{Data2: v}
	v1.Data2.ClassName = ClassChannelRoleEditor
	return v1
}

func NewTLChannelParticipantsRecent(v *ChannelParticipantsFilter) *TLChannelParticipantsRecent {
	if v == nil {
		return &TLChannelParticipantsRecent{
			Data2: &ChannelParticipantsFilter{
				Cmd:       Cmd_ChannelParticipantsRecent,
				ClassName: ClassChannelParticipantsRecent,
			},
		}
	}
	v1 := &TLChannelParticipantsRecent{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantsRecent
	return v1
}

func NewTLChannelParticipantsAdmins(v *ChannelParticipantsFilter) *TLChannelParticipantsAdmins {
	if v == nil {
		return &TLChannelParticipantsAdmins{
			Data2: &ChannelParticipantsFilter{
				Cmd:       Cmd_ChannelParticipantsAdmins,
				ClassName: ClassChannelParticipantsAdmins,
			},
		}
	}
	v1 := &TLChannelParticipantsAdmins{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantsAdmins
	return v1
}

func NewTLChannelParticipantsKicked(v *ChannelParticipantsFilter) *TLChannelParticipantsKicked {
	if v == nil {
		return &TLChannelParticipantsKicked{
			Data2: &ChannelParticipantsFilter{
				Cmd:       Cmd_ChannelParticipantsKicked,
				ClassName: ClassChannelParticipantsKicked,
			},
		}
	}
	v1 := &TLChannelParticipantsKicked{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantsKicked
	return v1
}

func NewTLChannelParticipantsBots(v *ChannelParticipantsFilter) *TLChannelParticipantsBots {
	if v == nil {
		return &TLChannelParticipantsBots{
			Data2: &ChannelParticipantsFilter{
				Cmd:       Cmd_ChannelParticipantsBots,
				ClassName: ClassChannelParticipantsBots,
			},
		}
	}
	v1 := &TLChannelParticipantsBots{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantsBots
	return v1
}

func NewTLChannelParticipantsBanned(v *ChannelParticipantsFilter) *TLChannelParticipantsBanned {
	if v == nil {
		return &TLChannelParticipantsBanned{
			Data2: &ChannelParticipantsFilter{
				Cmd:       Cmd_ChannelParticipantsBanned,
				ClassName: ClassChannelParticipantsBanned,
			},
		}
	}
	v1 := &TLChannelParticipantsBanned{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantsBanned
	return v1
}

func NewTLChannelParticipantsSearch(v *ChannelParticipantsFilter) *TLChannelParticipantsSearch {
	if v == nil {
		return &TLChannelParticipantsSearch{
			Data2: &ChannelParticipantsFilter{
				Cmd:       Cmd_ChannelParticipantsSearch,
				ClassName: ClassChannelParticipantsSearch,
			},
		}
	}
	v1 := &TLChannelParticipantsSearch{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantsSearch
	return v1
}

func NewTLChannelParticipantsContacts(v *ChannelParticipantsFilter) *TLChannelParticipantsContacts {
	if v == nil {
		return &TLChannelParticipantsContacts{
			Data2: &ChannelParticipantsFilter{
				Cmd:       Cmd_ChannelParticipantsContacts,
				ClassName: ClassChannelParticipantsContacts,
			},
		}
	}
	v1 := &TLChannelParticipantsContacts{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantsContacts
	return v1
}

func NewTLChannelParticipantsMentions(v *ChannelParticipantsFilter) *TLChannelParticipantsMentions {
	if v == nil {
		return &TLChannelParticipantsMentions{
			Data2: &ChannelParticipantsFilter{
				Cmd:       Cmd_ChannelParticipantsMentions,
				ClassName: ClassChannelParticipantsMentions,
			},
		}
	}
	v1 := &TLChannelParticipantsMentions{Data2: v}
	v1.Data2.ClassName = ClassChannelParticipantsMentions
	return v1
}

func NewTLChannelsAdminLogResults(v *Channels_AdminLogResults) *TLChannelsAdminLogResults {
	if v == nil {
		return &TLChannelsAdminLogResults{
			Data2: &Channels_AdminLogResults{
				Cmd:       Cmd_ChannelsAdminLogResults,
				ClassName: ClassChannelsAdminLogResults,
			},
		}
	}
	v1 := &TLChannelsAdminLogResults{Data2: v}
	v1.Data2.ClassName = ClassChannelsAdminLogResults
	return v1
}

func NewTLChannelsChannelParticipant(v *Channels_ChannelParticipant) *TLChannelsChannelParticipant {
	if v == nil {
		return &TLChannelsChannelParticipant{
			Data2: &Channels_ChannelParticipant{
				Cmd:       Cmd_ChannelsChannelParticipant,
				ClassName: ClassChannelsChannelParticipant,
			},
		}
	}
	v1 := &TLChannelsChannelParticipant{Data2: v}
	v1.Data2.ClassName = ClassChannelsChannelParticipant
	return v1
}

func NewTLChannelsChannelParticipants(v *Channels_ChannelParticipants) *TLChannelsChannelParticipants {
	if v == nil {
		return &TLChannelsChannelParticipants{
			Data2: &Channels_ChannelParticipants{
				Cmd:       Cmd_ChannelsChannelParticipants,
				ClassName: ClassChannelsChannelParticipants,
			},
		}
	}
	v1 := &TLChannelsChannelParticipants{Data2: v}
	v1.Data2.ClassName = ClassChannelsChannelParticipants
	return v1
}

func NewTLChannelsChannelParticipantsNotModified(v *Channels_ChannelParticipants) *TLChannelsChannelParticipantsNotModified {
	if v == nil {
		return &TLChannelsChannelParticipantsNotModified{
			Data2: &Channels_ChannelParticipants{
				Cmd:       Cmd_ChannelsChannelParticipantsNotModified,
				ClassName: ClassChannelsChannelParticipantsNotModified,
			},
		}
	}
	v1 := &TLChannelsChannelParticipantsNotModified{Data2: v}
	v1.Data2.ClassName = ClassChannelsChannelParticipantsNotModified
	return v1
}

func NewTLChatEmpty(v *Chat) *TLChatEmpty {
	if v == nil {
		return &TLChatEmpty{
			Data2: &Chat{
				Cmd:       Cmd_ChatEmpty,
				ClassName: ClassChatEmpty,
			},
		}
	}
	v1 := &TLChatEmpty{Data2: v}
	v1.Data2.ClassName = ClassChatEmpty
	return v1
}

func NewTLChat(v *Chat) *TLChat {
	if v == nil {
		return &TLChat{
			Data2: &Chat{
				Cmd:       Cmd_Chat,
				ClassName: ClassChat,
			},
		}
	}
	v1 := &TLChat{Data2: v}
	v1.Data2.ClassName = ClassChat
	return v1
}

func NewTLChatForbidden(v *Chat) *TLChatForbidden {
	if v == nil {
		return &TLChatForbidden{
			Data2: &Chat{
				Cmd:       Cmd_ChatForbidden,
				ClassName: ClassChatForbidden,
			},
		}
	}
	v1 := &TLChatForbidden{Data2: v}
	v1.Data2.ClassName = ClassChatForbidden
	return v1
}

func NewTLChannel(v *Chat) *TLChannel {
	if v == nil {
		return &TLChannel{
			Data2: &Chat{
				Cmd:       Cmd_Channel,
				ClassName: ClassChannel,
			},
		}
	}
	v1 := &TLChannel{Data2: v}
	v1.Data2.ClassName = ClassChannel
	return v1
}

func NewTLChannelForbidden(v *Chat) *TLChannelForbidden {
	if v == nil {
		return &TLChannelForbidden{
			Data2: &Chat{
				Cmd:       Cmd_ChannelForbidden,
				ClassName: ClassChannelForbidden,
			},
		}
	}
	v1 := &TLChannelForbidden{Data2: v}
	v1.Data2.ClassName = ClassChannelForbidden
	return v1
}

func NewTLChatAdminRights(v *ChatAdminRights) *TLChatAdminRights {
	if v == nil {
		return &TLChatAdminRights{
			Data2: &ChatAdminRights{
				Cmd:       Cmd_ChatAdminRights,
				ClassName: ClassChatAdminRights,
			},
		}
	}
	v1 := &TLChatAdminRights{Data2: v}
	v1.Data2.ClassName = ClassChatAdminRights
	return v1
}

func NewTLChatBannedRights(v *ChatBannedRights) *TLChatBannedRights {
	if v == nil {
		return &TLChatBannedRights{
			Data2: &ChatBannedRights{
				Cmd:       Cmd_ChatBannedRights,
				ClassName: ClassChatBannedRights,
			},
		}
	}
	v1 := &TLChatBannedRights{Data2: v}
	v1.Data2.ClassName = ClassChatBannedRights
	return v1
}

func NewTLChatFull(v *ChatFull) *TLChatFull {
	if v == nil {
		return &TLChatFull{
			Data2: &ChatFull{
				Cmd:       Cmd_ChatFull,
				ClassName: ClassChatFull,
			},
		}
	}
	v1 := &TLChatFull{Data2: v}
	v1.Data2.ClassName = ClassChatFull
	return v1
}

func NewTLChannelFull(v *ChatFull) *TLChannelFull {
	if v == nil {
		return &TLChannelFull{
			Data2: &ChatFull{
				Cmd:       Cmd_ChannelFull,
				ClassName: ClassChannelFull,
			},
		}
	}
	v1 := &TLChannelFull{Data2: v}
	v1.Data2.ClassName = ClassChannelFull
	return v1
}

func NewTLChatInviteAlready(v *ChatInvite) *TLChatInviteAlready {
	if v == nil {
		return &TLChatInviteAlready{
			Data2: &ChatInvite{
				Cmd:       Cmd_ChatInviteAlready,
				ClassName: ClassChatInviteAlready,
			},
		}
	}
	v1 := &TLChatInviteAlready{Data2: v}
	v1.Data2.ClassName = ClassChatInviteAlready
	return v1
}

func NewTLChatInvite(v *ChatInvite) *TLChatInvite {
	if v == nil {
		return &TLChatInvite{
			Data2: &ChatInvite{
				Cmd:       Cmd_ChatInvite,
				ClassName: ClassChatInvite,
			},
		}
	}
	v1 := &TLChatInvite{Data2: v}
	v1.Data2.ClassName = ClassChatInvite
	return v1
}

func NewTLChatInvitePeek(v *ChatInvite) *TLChatInvitePeek {
	if v == nil {
		return &TLChatInvitePeek{
			Data2: &ChatInvite{
				Cmd:       Cmd_ChatInvitePeek,
				ClassName: ClassChatInvitePeek,
			},
		}
	}
	v1 := &TLChatInvitePeek{Data2: v}
	v1.Data2.ClassName = ClassChatInvitePeek
	return v1
}

func NewTLChatOnlines(v *ChatOnlines) *TLChatOnlines {
	if v == nil {
		return &TLChatOnlines{
			Data2: &ChatOnlines{
				Cmd:       Cmd_ChatOnlines,
				ClassName: ClassChatOnlines,
			},
		}
	}
	v1 := &TLChatOnlines{Data2: v}
	v1.Data2.ClassName = ClassChatOnlines
	return v1
}

func NewTLChatParticipant(v *ChatParticipant) *TLChatParticipant {
	if v == nil {
		return &TLChatParticipant{
			Data2: &ChatParticipant{
				Cmd:       Cmd_ChatParticipant,
				ClassName: ClassChatParticipant,
			},
		}
	}
	v1 := &TLChatParticipant{Data2: v}
	v1.Data2.ClassName = ClassChatParticipant
	return v1
}

func NewTLChatParticipantCreator(v *ChatParticipant) *TLChatParticipantCreator {
	if v == nil {
		return &TLChatParticipantCreator{
			Data2: &ChatParticipant{
				Cmd:       Cmd_ChatParticipantCreator,
				ClassName: ClassChatParticipantCreator,
			},
		}
	}
	v1 := &TLChatParticipantCreator{Data2: v}
	v1.Data2.ClassName = ClassChatParticipantCreator
	return v1
}

func NewTLChatParticipantAdmin(v *ChatParticipant) *TLChatParticipantAdmin {
	if v == nil {
		return &TLChatParticipantAdmin{
			Data2: &ChatParticipant{
				Cmd:       Cmd_ChatParticipantAdmin,
				ClassName: ClassChatParticipantAdmin,
			},
		}
	}
	v1 := &TLChatParticipantAdmin{Data2: v}
	v1.Data2.ClassName = ClassChatParticipantAdmin
	return v1
}

func NewTLChatParticipantsForbidden(v *ChatParticipants) *TLChatParticipantsForbidden {
	if v == nil {
		return &TLChatParticipantsForbidden{
			Data2: &ChatParticipants{
				Cmd:       Cmd_ChatParticipantsForbidden,
				ClassName: ClassChatParticipantsForbidden,
			},
		}
	}
	v1 := &TLChatParticipantsForbidden{Data2: v}
	v1.Data2.ClassName = ClassChatParticipantsForbidden
	return v1
}

func NewTLChatParticipants(v *ChatParticipants) *TLChatParticipants {
	if v == nil {
		return &TLChatParticipants{
			Data2: &ChatParticipants{
				Cmd:       Cmd_ChatParticipants,
				ClassName: ClassChatParticipants,
			},
		}
	}
	v1 := &TLChatParticipants{Data2: v}
	v1.Data2.ClassName = ClassChatParticipants
	return v1
}

func NewTLChatPhotoEmpty(v *ChatPhoto) *TLChatPhotoEmpty {
	if v == nil {
		return &TLChatPhotoEmpty{
			Data2: &ChatPhoto{
				Cmd:       Cmd_ChatPhotoEmpty,
				ClassName: ClassChatPhotoEmpty,
			},
		}
	}
	v1 := &TLChatPhotoEmpty{Data2: v}
	v1.Data2.ClassName = ClassChatPhotoEmpty
	return v1
}

func NewTLChatPhoto(v *ChatPhoto) *TLChatPhoto {
	if v == nil {
		return &TLChatPhoto{
			Data2: &ChatPhoto{
				Cmd:       Cmd_ChatPhoto,
				ClassName: ClassChatPhoto,
			},
		}
	}
	v1 := &TLChatPhoto{Data2: v}
	v1.Data2.ClassName = ClassChatPhoto
	return v1
}

func NewTLClient_DHInnerData(v *Client_DH_Inner_Data) *TLClient_DHInnerData {
	if v == nil {
		return &TLClient_DHInnerData{
			Data2: &Client_DH_Inner_Data{
				Cmd:       Cmd_Client_DHInnerData,
				ClassName: ClassClient_DHInnerData,
			},
		}
	}
	v1 := &TLClient_DHInnerData{Data2: v}
	v1.Data2.ClassName = ClassClient_DHInnerData
	return v1
}

func NewTLCodeSettings(v *CodeSettings) *TLCodeSettings {
	if v == nil {
		return &TLCodeSettings{
			Data2: &CodeSettings{
				Cmd:       Cmd_CodeSettings,
				ClassName: ClassCodeSettings,
			},
		}
	}
	v1 := &TLCodeSettings{Data2: v}
	v1.Data2.ClassName = ClassCodeSettings
	return v1
}

func NewTLConfig(v *Config) *TLConfig {
	if v == nil {
		return &TLConfig{
			Data2: &Config{
				Cmd:       Cmd_Config,
				ClassName: ClassConfig,
			},
		}
	}
	v1 := &TLConfig{Data2: v}
	v1.Data2.ClassName = ClassConfig
	return v1
}

func NewTLContact(v *Contact) *TLContact {
	if v == nil {
		return &TLContact{
			Data2: &Contact{
				Cmd:       Cmd_Contact,
				ClassName: ClassContact,
			},
		}
	}
	v1 := &TLContact{Data2: v}
	v1.Data2.ClassName = ClassContact
	return v1
}

func NewTLContactBlocked(v *ContactBlocked) *TLContactBlocked {
	if v == nil {
		return &TLContactBlocked{
			Data2: &ContactBlocked{
				Cmd:       Cmd_ContactBlocked,
				ClassName: ClassContactBlocked,
			},
		}
	}
	v1 := &TLContactBlocked{Data2: v}
	v1.Data2.ClassName = ClassContactBlocked
	return v1
}

func NewTLContactLinkUnknown(v *ContactLink) *TLContactLinkUnknown {
	if v == nil {
		return &TLContactLinkUnknown{
			Data2: &ContactLink{
				Cmd:       Cmd_ContactLinkUnknown,
				ClassName: ClassContactLinkUnknown,
			},
		}
	}
	v1 := &TLContactLinkUnknown{Data2: v}
	v1.Data2.ClassName = ClassContactLinkUnknown
	return v1
}

func NewTLContactLinkNone(v *ContactLink) *TLContactLinkNone {
	if v == nil {
		return &TLContactLinkNone{
			Data2: &ContactLink{
				Cmd:       Cmd_ContactLinkNone,
				ClassName: ClassContactLinkNone,
			},
		}
	}
	v1 := &TLContactLinkNone{Data2: v}
	v1.Data2.ClassName = ClassContactLinkNone
	return v1
}

func NewTLContactLinkHasPhone(v *ContactLink) *TLContactLinkHasPhone {
	if v == nil {
		return &TLContactLinkHasPhone{
			Data2: &ContactLink{
				Cmd:       Cmd_ContactLinkHasPhone,
				ClassName: ClassContactLinkHasPhone,
			},
		}
	}
	v1 := &TLContactLinkHasPhone{Data2: v}
	v1.Data2.ClassName = ClassContactLinkHasPhone
	return v1
}

func NewTLContactLinkContact(v *ContactLink) *TLContactLinkContact {
	if v == nil {
		return &TLContactLinkContact{
			Data2: &ContactLink{
				Cmd:       Cmd_ContactLinkContact,
				ClassName: ClassContactLinkContact,
			},
		}
	}
	v1 := &TLContactLinkContact{Data2: v}
	v1.Data2.ClassName = ClassContactLinkContact
	return v1
}

func NewTLContactStatus(v *ContactStatus) *TLContactStatus {
	if v == nil {
		return &TLContactStatus{
			Data2: &ContactStatus{
				Cmd:       Cmd_ContactStatus,
				ClassName: ClassContactStatus,
			},
		}
	}
	v1 := &TLContactStatus{Data2: v}
	v1.Data2.ClassName = ClassContactStatus
	return v1
}

func NewTLContactsBlocked(v *Contacts_Blocked) *TLContactsBlocked {
	if v == nil {
		return &TLContactsBlocked{
			Data2: &Contacts_Blocked{
				Cmd:       Cmd_ContactsBlocked,
				ClassName: ClassContactsBlocked,
			},
		}
	}
	v1 := &TLContactsBlocked{Data2: v}
	v1.Data2.ClassName = ClassContactsBlocked
	return v1
}

func NewTLContactsBlockedSlice(v *Contacts_Blocked) *TLContactsBlockedSlice {
	if v == nil {
		return &TLContactsBlockedSlice{
			Data2: &Contacts_Blocked{
				Cmd:       Cmd_ContactsBlockedSlice,
				ClassName: ClassContactsBlockedSlice,
			},
		}
	}
	v1 := &TLContactsBlockedSlice{Data2: v}
	v1.Data2.ClassName = ClassContactsBlockedSlice
	return v1
}

func NewTLContactsContactsNotModified(v *Contacts_Contacts) *TLContactsContactsNotModified {
	if v == nil {
		return &TLContactsContactsNotModified{
			Data2: &Contacts_Contacts{
				Cmd:       Cmd_ContactsContactsNotModified,
				ClassName: ClassContactsContactsNotModified,
			},
		}
	}
	v1 := &TLContactsContactsNotModified{Data2: v}
	v1.Data2.ClassName = ClassContactsContactsNotModified
	return v1
}

func NewTLContactsContacts(v *Contacts_Contacts) *TLContactsContacts {
	if v == nil {
		return &TLContactsContacts{
			Data2: &Contacts_Contacts{
				Cmd:       Cmd_ContactsContacts,
				ClassName: ClassContactsContacts,
			},
		}
	}
	v1 := &TLContactsContacts{Data2: v}
	v1.Data2.ClassName = ClassContactsContacts
	return v1
}

func NewTLContactsFound(v *Contacts_Found) *TLContactsFound {
	if v == nil {
		return &TLContactsFound{
			Data2: &Contacts_Found{
				Cmd:       Cmd_ContactsFound,
				ClassName: ClassContactsFound,
			},
		}
	}
	v1 := &TLContactsFound{Data2: v}
	v1.Data2.ClassName = ClassContactsFound
	return v1
}

func NewTLContactsImportedContacts(v *Contacts_ImportedContacts) *TLContactsImportedContacts {
	if v == nil {
		return &TLContactsImportedContacts{
			Data2: &Contacts_ImportedContacts{
				Cmd:       Cmd_ContactsImportedContacts,
				ClassName: ClassContactsImportedContacts,
			},
		}
	}
	v1 := &TLContactsImportedContacts{Data2: v}
	v1.Data2.ClassName = ClassContactsImportedContacts
	return v1
}

func NewTLContactsLink(v *Contacts_Link) *TLContactsLink {
	if v == nil {
		return &TLContactsLink{
			Data2: &Contacts_Link{
				Cmd:       Cmd_ContactsLink,
				ClassName: ClassContactsLink,
			},
		}
	}
	v1 := &TLContactsLink{Data2: v}
	v1.Data2.ClassName = ClassContactsLink
	return v1
}

func NewTLContactsResolvedPeer(v *Contacts_ResolvedPeer) *TLContactsResolvedPeer {
	if v == nil {
		return &TLContactsResolvedPeer{
			Data2: &Contacts_ResolvedPeer{
				Cmd:       Cmd_ContactsResolvedPeer,
				ClassName: ClassContactsResolvedPeer,
			},
		}
	}
	v1 := &TLContactsResolvedPeer{Data2: v}
	v1.Data2.ClassName = ClassContactsResolvedPeer
	return v1
}

func NewTLContactsTopPeersNotModified(v *Contacts_TopPeers) *TLContactsTopPeersNotModified {
	if v == nil {
		return &TLContactsTopPeersNotModified{
			Data2: &Contacts_TopPeers{
				Cmd:       Cmd_ContactsTopPeersNotModified,
				ClassName: ClassContactsTopPeersNotModified,
			},
		}
	}
	v1 := &TLContactsTopPeersNotModified{Data2: v}
	v1.Data2.ClassName = ClassContactsTopPeersNotModified
	return v1
}

func NewTLContactsTopPeers(v *Contacts_TopPeers) *TLContactsTopPeers {
	if v == nil {
		return &TLContactsTopPeers{
			Data2: &Contacts_TopPeers{
				Cmd:       Cmd_ContactsTopPeers,
				ClassName: ClassContactsTopPeers,
			},
		}
	}
	v1 := &TLContactsTopPeers{Data2: v}
	v1.Data2.ClassName = ClassContactsTopPeers
	return v1
}

func NewTLContactsTopPeersDisabled(v *Contacts_TopPeers) *TLContactsTopPeersDisabled {
	if v == nil {
		return &TLContactsTopPeersDisabled{
			Data2: &Contacts_TopPeers{
				Cmd:       Cmd_ContactsTopPeersDisabled,
				ClassName: ClassContactsTopPeersDisabled,
			},
		}
	}
	v1 := &TLContactsTopPeersDisabled{Data2: v}
	v1.Data2.ClassName = ClassContactsTopPeersDisabled
	return v1
}

func NewTLDataJSON(v *DataJSON) *TLDataJSON {
	if v == nil {
		return &TLDataJSON{
			Data2: &DataJSON{
				Cmd:       Cmd_DataJSON,
				ClassName: ClassDataJSON,
			},
		}
	}
	v1 := &TLDataJSON{Data2: v}
	v1.Data2.ClassName = ClassDataJSON
	return v1
}

func NewTLDcOption(v *DcOption) *TLDcOption {
	if v == nil {
		return &TLDcOption{
			Data2: &DcOption{
				Cmd:       Cmd_DcOption,
				ClassName: ClassDcOption,
			},
		}
	}
	v1 := &TLDcOption{Data2: v}
	v1.Data2.ClassName = ClassDcOption
	return v1
}

func NewTLDestroyAuthKeyOk(v *DestroyAuthKeyRes) *TLDestroyAuthKeyOk {
	if v == nil {
		return &TLDestroyAuthKeyOk{
			Data2: &DestroyAuthKeyRes{
				Cmd:       Cmd_DestroyAuthKeyOk,
				ClassName: ClassDestroyAuthKeyOk,
			},
		}
	}
	v1 := &TLDestroyAuthKeyOk{Data2: v}
	v1.Data2.ClassName = ClassDestroyAuthKeyOk
	return v1
}

func NewTLDestroyAuthKeyNone(v *DestroyAuthKeyRes) *TLDestroyAuthKeyNone {
	if v == nil {
		return &TLDestroyAuthKeyNone{
			Data2: &DestroyAuthKeyRes{
				Cmd:       Cmd_DestroyAuthKeyNone,
				ClassName: ClassDestroyAuthKeyNone,
			},
		}
	}
	v1 := &TLDestroyAuthKeyNone{Data2: v}
	v1.Data2.ClassName = ClassDestroyAuthKeyNone
	return v1
}

func NewTLDestroyAuthKeyFail(v *DestroyAuthKeyRes) *TLDestroyAuthKeyFail {
	if v == nil {
		return &TLDestroyAuthKeyFail{
			Data2: &DestroyAuthKeyRes{
				Cmd:       Cmd_DestroyAuthKeyFail,
				ClassName: ClassDestroyAuthKeyFail,
			},
		}
	}
	v1 := &TLDestroyAuthKeyFail{Data2: v}
	v1.Data2.ClassName = ClassDestroyAuthKeyFail
	return v1
}

func NewTLDestroySessionOk(v *DestroySessionRes) *TLDestroySessionOk {
	if v == nil {
		return &TLDestroySessionOk{
			Data2: &DestroySessionRes{
				Cmd:       Cmd_DestroySessionOk,
				ClassName: ClassDestroySessionOk,
			},
		}
	}
	v1 := &TLDestroySessionOk{Data2: v}
	v1.Data2.ClassName = ClassDestroySessionOk
	return v1
}

func NewTLDestroySessionNone(v *DestroySessionRes) *TLDestroySessionNone {
	if v == nil {
		return &TLDestroySessionNone{
			Data2: &DestroySessionRes{
				Cmd:       Cmd_DestroySessionNone,
				ClassName: ClassDestroySessionNone,
			},
		}
	}
	v1 := &TLDestroySessionNone{Data2: v}
	v1.Data2.ClassName = ClassDestroySessionNone
	return v1
}

func NewTLDialog(v *Dialog) *TLDialog {
	if v == nil {
		return &TLDialog{
			Data2: &Dialog{
				Cmd:       Cmd_Dialog,
				ClassName: ClassDialog,
			},
		}
	}
	v1 := &TLDialog{Data2: v}
	v1.Data2.ClassName = ClassDialog
	return v1
}

func NewTLDialogChannel(v *Dialog) *TLDialogChannel {
	if v == nil {
		return &TLDialogChannel{
			Data2: &Dialog{
				Cmd:       Cmd_DialogChannel,
				ClassName: ClassDialogChannel,
			},
		}
	}
	v1 := &TLDialogChannel{Data2: v}
	v1.Data2.ClassName = ClassDialogChannel
	return v1
}

func NewTLDialogFolder(v *Dialog) *TLDialogFolder {
	if v == nil {
		return &TLDialogFolder{
			Data2: &Dialog{
				Cmd:       Cmd_DialogFolder,
				ClassName: ClassDialogFolder,
			},
		}
	}
	v1 := &TLDialogFolder{Data2: v}
	v1.Data2.ClassName = ClassDialogFolder
	return v1
}

func NewTLDialogFilter(v *DialogFilter) *TLDialogFilter {
	if v == nil {
		return &TLDialogFilter{
			Data2: &DialogFilter{
				Cmd:       Cmd_DialogFilter,
				ClassName: ClassDialogFilter,
			},
		}
	}
	v1 := &TLDialogFilter{Data2: v}
	v1.Data2.ClassName = ClassDialogFilter
	return v1
}

func NewTLDialogFilterSuggested(v *DialogFilterSuggested) *TLDialogFilterSuggested {
	if v == nil {
		return &TLDialogFilterSuggested{
			Data2: &DialogFilterSuggested{
				Cmd:       Cmd_DialogFilterSuggested,
				ClassName: ClassDialogFilterSuggested,
			},
		}
	}
	v1 := &TLDialogFilterSuggested{Data2: v}
	v1.Data2.ClassName = ClassDialogFilterSuggested
	return v1
}

func NewTLDialogPeer(v *DialogPeer) *TLDialogPeer {
	if v == nil {
		return &TLDialogPeer{
			Data2: &DialogPeer{
				Cmd:       Cmd_DialogPeer,
				ClassName: ClassDialogPeer,
			},
		}
	}
	v1 := &TLDialogPeer{Data2: v}
	v1.Data2.ClassName = ClassDialogPeer
	return v1
}

func NewTLDialogPeerFolder(v *DialogPeer) *TLDialogPeerFolder {
	if v == nil {
		return &TLDialogPeerFolder{
			Data2: &DialogPeer{
				Cmd:       Cmd_DialogPeerFolder,
				ClassName: ClassDialogPeerFolder,
			},
		}
	}
	v1 := &TLDialogPeerFolder{Data2: v}
	v1.Data2.ClassName = ClassDialogPeerFolder
	return v1
}

func NewTLDisabledFeature(v *DisabledFeature) *TLDisabledFeature {
	if v == nil {
		return &TLDisabledFeature{
			Data2: &DisabledFeature{
				Cmd:       Cmd_DisabledFeature,
				ClassName: ClassDisabledFeature,
			},
		}
	}
	v1 := &TLDisabledFeature{Data2: v}
	v1.Data2.ClassName = ClassDisabledFeature
	return v1
}

func NewTLDocumentEmpty(v *Document) *TLDocumentEmpty {
	if v == nil {
		return &TLDocumentEmpty{
			Data2: &Document{
				Cmd:       Cmd_DocumentEmpty,
				ClassName: ClassDocumentEmpty,
			},
		}
	}
	v1 := &TLDocumentEmpty{Data2: v}
	v1.Data2.ClassName = ClassDocumentEmpty
	return v1
}

func NewTLDocument(v *Document) *TLDocument {
	if v == nil {
		return &TLDocument{
			Data2: &Document{
				Cmd:       Cmd_Document,
				ClassName: ClassDocument,
			},
		}
	}
	v1 := &TLDocument{Data2: v}
	v1.Data2.ClassName = ClassDocument
	return v1
}

func NewTLDocumentAttributeImageSize(v *DocumentAttribute) *TLDocumentAttributeImageSize {
	if v == nil {
		return &TLDocumentAttributeImageSize{
			Data2: &DocumentAttribute{
				Cmd:       Cmd_DocumentAttributeImageSize,
				ClassName: ClassDocumentAttributeImageSize,
			},
		}
	}
	v1 := &TLDocumentAttributeImageSize{Data2: v}
	v1.Data2.ClassName = ClassDocumentAttributeImageSize
	return v1
}

func NewTLDocumentAttributeAnimated(v *DocumentAttribute) *TLDocumentAttributeAnimated {
	if v == nil {
		return &TLDocumentAttributeAnimated{
			Data2: &DocumentAttribute{
				Cmd:       Cmd_DocumentAttributeAnimated,
				ClassName: ClassDocumentAttributeAnimated,
			},
		}
	}
	v1 := &TLDocumentAttributeAnimated{Data2: v}
	v1.Data2.ClassName = ClassDocumentAttributeAnimated
	return v1
}

func NewTLDocumentAttributeSticker(v *DocumentAttribute) *TLDocumentAttributeSticker {
	if v == nil {
		return &TLDocumentAttributeSticker{
			Data2: &DocumentAttribute{
				Cmd:       Cmd_DocumentAttributeSticker,
				ClassName: ClassDocumentAttributeSticker,
			},
		}
	}
	v1 := &TLDocumentAttributeSticker{Data2: v}
	v1.Data2.ClassName = ClassDocumentAttributeSticker
	return v1
}

func NewTLDocumentAttributeVideo(v *DocumentAttribute) *TLDocumentAttributeVideo {
	if v == nil {
		return &TLDocumentAttributeVideo{
			Data2: &DocumentAttribute{
				Cmd:       Cmd_DocumentAttributeVideo,
				ClassName: ClassDocumentAttributeVideo,
			},
		}
	}
	v1 := &TLDocumentAttributeVideo{Data2: v}
	v1.Data2.ClassName = ClassDocumentAttributeVideo
	return v1
}

func NewTLDocumentAttributeAudio(v *DocumentAttribute) *TLDocumentAttributeAudio {
	if v == nil {
		return &TLDocumentAttributeAudio{
			Data2: &DocumentAttribute{
				Cmd:       Cmd_DocumentAttributeAudio,
				ClassName: ClassDocumentAttributeAudio,
			},
		}
	}
	v1 := &TLDocumentAttributeAudio{Data2: v}
	v1.Data2.ClassName = ClassDocumentAttributeAudio
	return v1
}

func NewTLDocumentAttributeFilename(v *DocumentAttribute) *TLDocumentAttributeFilename {
	if v == nil {
		return &TLDocumentAttributeFilename{
			Data2: &DocumentAttribute{
				Cmd:       Cmd_DocumentAttributeFilename,
				ClassName: ClassDocumentAttributeFilename,
			},
		}
	}
	v1 := &TLDocumentAttributeFilename{Data2: v}
	v1.Data2.ClassName = ClassDocumentAttributeFilename
	return v1
}

func NewTLDocumentAttributeHasStickers(v *DocumentAttribute) *TLDocumentAttributeHasStickers {
	if v == nil {
		return &TLDocumentAttributeHasStickers{
			Data2: &DocumentAttribute{
				Cmd:       Cmd_DocumentAttributeHasStickers,
				ClassName: ClassDocumentAttributeHasStickers,
			},
		}
	}
	v1 := &TLDocumentAttributeHasStickers{Data2: v}
	v1.Data2.ClassName = ClassDocumentAttributeHasStickers
	return v1
}

func NewTLDraftMessageEmpty(v *DraftMessage) *TLDraftMessageEmpty {
	if v == nil {
		return &TLDraftMessageEmpty{
			Data2: &DraftMessage{
				Cmd:       Cmd_DraftMessageEmpty,
				ClassName: ClassDraftMessageEmpty,
			},
		}
	}
	v1 := &TLDraftMessageEmpty{Data2: v}
	v1.Data2.ClassName = ClassDraftMessageEmpty
	return v1
}

func NewTLDraftMessage(v *DraftMessage) *TLDraftMessage {
	if v == nil {
		return &TLDraftMessage{
			Data2: &DraftMessage{
				Cmd:       Cmd_DraftMessage,
				ClassName: ClassDraftMessage,
			},
		}
	}
	v1 := &TLDraftMessage{Data2: v}
	v1.Data2.ClassName = ClassDraftMessage
	return v1
}

func NewTLEmojiKeyword(v *EmojiKeyword) *TLEmojiKeyword {
	if v == nil {
		return &TLEmojiKeyword{
			Data2: &EmojiKeyword{
				Cmd:       Cmd_EmojiKeyword,
				ClassName: ClassEmojiKeyword,
			},
		}
	}
	v1 := &TLEmojiKeyword{Data2: v}
	v1.Data2.ClassName = ClassEmojiKeyword
	return v1
}

func NewTLEmojiKeywordDeleted(v *EmojiKeyword) *TLEmojiKeywordDeleted {
	if v == nil {
		return &TLEmojiKeywordDeleted{
			Data2: &EmojiKeyword{
				Cmd:       Cmd_EmojiKeywordDeleted,
				ClassName: ClassEmojiKeywordDeleted,
			},
		}
	}
	v1 := &TLEmojiKeywordDeleted{Data2: v}
	v1.Data2.ClassName = ClassEmojiKeywordDeleted
	return v1
}

func NewTLEmojiKeywordsDifference(v *EmojiKeywordsDifference) *TLEmojiKeywordsDifference {
	if v == nil {
		return &TLEmojiKeywordsDifference{
			Data2: &EmojiKeywordsDifference{
				Cmd:       Cmd_EmojiKeywordsDifference,
				ClassName: ClassEmojiKeywordsDifference,
			},
		}
	}
	v1 := &TLEmojiKeywordsDifference{Data2: v}
	v1.Data2.ClassName = ClassEmojiKeywordsDifference
	return v1
}

func NewTLEmojiLanguage(v *EmojiLanguage) *TLEmojiLanguage {
	if v == nil {
		return &TLEmojiLanguage{
			Data2: &EmojiLanguage{
				Cmd:       Cmd_EmojiLanguage,
				ClassName: ClassEmojiLanguage,
			},
		}
	}
	v1 := &TLEmojiLanguage{Data2: v}
	v1.Data2.ClassName = ClassEmojiLanguage
	return v1
}

func NewTLEmojiURL(v *EmojiURL) *TLEmojiURL {
	if v == nil {
		return &TLEmojiURL{
			Data2: &EmojiURL{
				Cmd:       Cmd_EmojiURL,
				ClassName: ClassEmojiURL,
			},
		}
	}
	v1 := &TLEmojiURL{Data2: v}
	v1.Data2.ClassName = ClassEmojiURL
	return v1
}

func NewTLEncryptedChatEmpty(v *EncryptedChat) *TLEncryptedChatEmpty {
	if v == nil {
		return &TLEncryptedChatEmpty{
			Data2: &EncryptedChat{
				Cmd:       Cmd_EncryptedChatEmpty,
				ClassName: ClassEncryptedChatEmpty,
			},
		}
	}
	v1 := &TLEncryptedChatEmpty{Data2: v}
	v1.Data2.ClassName = ClassEncryptedChatEmpty
	return v1
}

func NewTLEncryptedChatWaiting(v *EncryptedChat) *TLEncryptedChatWaiting {
	if v == nil {
		return &TLEncryptedChatWaiting{
			Data2: &EncryptedChat{
				Cmd:       Cmd_EncryptedChatWaiting,
				ClassName: ClassEncryptedChatWaiting,
			},
		}
	}
	v1 := &TLEncryptedChatWaiting{Data2: v}
	v1.Data2.ClassName = ClassEncryptedChatWaiting
	return v1
}

func NewTLEncryptedChatRequested(v *EncryptedChat) *TLEncryptedChatRequested {
	if v == nil {
		return &TLEncryptedChatRequested{
			Data2: &EncryptedChat{
				Cmd:       Cmd_EncryptedChatRequested,
				ClassName: ClassEncryptedChatRequested,
			},
		}
	}
	v1 := &TLEncryptedChatRequested{Data2: v}
	v1.Data2.ClassName = ClassEncryptedChatRequested
	return v1
}

func NewTLEncryptedChat(v *EncryptedChat) *TLEncryptedChat {
	if v == nil {
		return &TLEncryptedChat{
			Data2: &EncryptedChat{
				Cmd:       Cmd_EncryptedChat,
				ClassName: ClassEncryptedChat,
			},
		}
	}
	v1 := &TLEncryptedChat{Data2: v}
	v1.Data2.ClassName = ClassEncryptedChat
	return v1
}

func NewTLEncryptedChatDiscarded(v *EncryptedChat) *TLEncryptedChatDiscarded {
	if v == nil {
		return &TLEncryptedChatDiscarded{
			Data2: &EncryptedChat{
				Cmd:       Cmd_EncryptedChatDiscarded,
				ClassName: ClassEncryptedChatDiscarded,
			},
		}
	}
	v1 := &TLEncryptedChatDiscarded{Data2: v}
	v1.Data2.ClassName = ClassEncryptedChatDiscarded
	return v1
}

func NewTLEncryptedFileEmpty(v *EncryptedFile) *TLEncryptedFileEmpty {
	if v == nil {
		return &TLEncryptedFileEmpty{
			Data2: &EncryptedFile{
				Cmd:       Cmd_EncryptedFileEmpty,
				ClassName: ClassEncryptedFileEmpty,
			},
		}
	}
	v1 := &TLEncryptedFileEmpty{Data2: v}
	v1.Data2.ClassName = ClassEncryptedFileEmpty
	return v1
}

func NewTLEncryptedFile(v *EncryptedFile) *TLEncryptedFile {
	if v == nil {
		return &TLEncryptedFile{
			Data2: &EncryptedFile{
				Cmd:       Cmd_EncryptedFile,
				ClassName: ClassEncryptedFile,
			},
		}
	}
	v1 := &TLEncryptedFile{Data2: v}
	v1.Data2.ClassName = ClassEncryptedFile
	return v1
}

func NewTLEncryptedMessage(v *EncryptedMessage) *TLEncryptedMessage {
	if v == nil {
		return &TLEncryptedMessage{
			Data2: &EncryptedMessage{
				Cmd:       Cmd_EncryptedMessage,
				ClassName: ClassEncryptedMessage,
			},
		}
	}
	v1 := &TLEncryptedMessage{Data2: v}
	v1.Data2.ClassName = ClassEncryptedMessage
	return v1
}

func NewTLEncryptedMessageService(v *EncryptedMessage) *TLEncryptedMessageService {
	if v == nil {
		return &TLEncryptedMessageService{
			Data2: &EncryptedMessage{
				Cmd:       Cmd_EncryptedMessageService,
				ClassName: ClassEncryptedMessageService,
			},
		}
	}
	v1 := &TLEncryptedMessageService{Data2: v}
	v1.Data2.ClassName = ClassEncryptedMessageService
	return v1
}

func NewTLError(v *Error) *TLError {
	if v == nil {
		return &TLError{
			Data2: &Error{
				Cmd:       Cmd_Error,
				ClassName: ClassError,
			},
		}
	}
	v1 := &TLError{Data2: v}
	v1.Data2.ClassName = ClassError
	return v1
}

func NewTLChatInviteEmpty(v *ExportedChatInvite) *TLChatInviteEmpty {
	if v == nil {
		return &TLChatInviteEmpty{
			Data2: &ExportedChatInvite{
				Cmd:       Cmd_ChatInviteEmpty,
				ClassName: ClassChatInviteEmpty,
			},
		}
	}
	v1 := &TLChatInviteEmpty{Data2: v}
	v1.Data2.ClassName = ClassChatInviteEmpty
	return v1
}

func NewTLChatInviteExported(v *ExportedChatInvite) *TLChatInviteExported {
	if v == nil {
		return &TLChatInviteExported{
			Data2: &ExportedChatInvite{
				Cmd:       Cmd_ChatInviteExported,
				ClassName: ClassChatInviteExported,
			},
		}
	}
	v1 := &TLChatInviteExported{Data2: v}
	v1.Data2.ClassName = ClassChatInviteExported
	return v1
}

func NewTLExportedMessageLink(v *ExportedMessageLink) *TLExportedMessageLink {
	if v == nil {
		return &TLExportedMessageLink{
			Data2: &ExportedMessageLink{
				Cmd:       Cmd_ExportedMessageLink,
				ClassName: ClassExportedMessageLink,
			},
		}
	}
	v1 := &TLExportedMessageLink{Data2: v}
	v1.Data2.ClassName = ClassExportedMessageLink
	return v1
}

func NewTLFileHash(v *FileHash) *TLFileHash {
	if v == nil {
		return &TLFileHash{
			Data2: &FileHash{
				Cmd:       Cmd_FileHash,
				ClassName: ClassFileHash,
			},
		}
	}
	v1 := &TLFileHash{Data2: v}
	v1.Data2.ClassName = ClassFileHash
	return v1
}

func NewTLFileLocationUnavailable(v *FileLocation) *TLFileLocationUnavailable {
	if v == nil {
		return &TLFileLocationUnavailable{
			Data2: &FileLocation{
				Cmd:       Cmd_FileLocationUnavailable,
				ClassName: ClassFileLocationUnavailable,
			},
		}
	}
	v1 := &TLFileLocationUnavailable{Data2: v}
	v1.Data2.ClassName = ClassFileLocationUnavailable
	return v1
}

func NewTLFileLocation(v *FileLocation) *TLFileLocation {
	if v == nil {
		return &TLFileLocation{
			Data2: &FileLocation{
				Cmd:       Cmd_FileLocation,
				ClassName: ClassFileLocation,
			},
		}
	}
	v1 := &TLFileLocation{Data2: v}
	v1.Data2.ClassName = ClassFileLocation
	return v1
}

func NewTLFileLocationToBeDeprecated(v *FileLocation) *TLFileLocationToBeDeprecated {
	if v == nil {
		return &TLFileLocationToBeDeprecated{
			Data2: &FileLocation{
				Cmd:       Cmd_FileLocationToBeDeprecated,
				ClassName: ClassFileLocationToBeDeprecated,
			},
		}
	}
	v1 := &TLFileLocationToBeDeprecated{Data2: v}
	v1.Data2.ClassName = ClassFileLocationToBeDeprecated
	return v1
}

func NewTLFolder(v *Folder) *TLFolder {
	if v == nil {
		return &TLFolder{
			Data2: &Folder{
				Cmd:       Cmd_Folder,
				ClassName: ClassFolder,
			},
		}
	}
	v1 := &TLFolder{Data2: v}
	v1.Data2.ClassName = ClassFolder
	return v1
}

func NewTLFolderPeer(v *FolderPeer) *TLFolderPeer {
	if v == nil {
		return &TLFolderPeer{
			Data2: &FolderPeer{
				Cmd:       Cmd_FolderPeer,
				ClassName: ClassFolderPeer,
			},
		}
	}
	v1 := &TLFolderPeer{Data2: v}
	v1.Data2.ClassName = ClassFolderPeer
	return v1
}

func NewTLFoundGif(v *FoundGif) *TLFoundGif {
	if v == nil {
		return &TLFoundGif{
			Data2: &FoundGif{
				Cmd:       Cmd_FoundGif,
				ClassName: ClassFoundGif,
			},
		}
	}
	v1 := &TLFoundGif{Data2: v}
	v1.Data2.ClassName = ClassFoundGif
	return v1
}

func NewTLFoundGifCached(v *FoundGif) *TLFoundGifCached {
	if v == nil {
		return &TLFoundGifCached{
			Data2: &FoundGif{
				Cmd:       Cmd_FoundGifCached,
				ClassName: ClassFoundGifCached,
			},
		}
	}
	v1 := &TLFoundGifCached{Data2: v}
	v1.Data2.ClassName = ClassFoundGifCached
	return v1
}

func NewTLFutureSalt(v *FutureSalt) *TLFutureSalt {
	if v == nil {
		return &TLFutureSalt{
			Data2: &FutureSalt{
				Cmd:       Cmd_FutureSalt,
				ClassName: ClassFutureSalt,
			},
		}
	}
	v1 := &TLFutureSalt{Data2: v}
	v1.Data2.ClassName = ClassFutureSalt
	return v1
}

func NewTLFutureSalts(v *FutureSalts) *TLFutureSalts {
	if v == nil {
		return &TLFutureSalts{
			Data2: &FutureSalts{
				Cmd:       Cmd_FutureSalts,
				ClassName: ClassFutureSalts,
			},
		}
	}
	v1 := &TLFutureSalts{Data2: v}
	v1.Data2.ClassName = ClassFutureSalts
	return v1
}

func NewTLGame(v *Game) *TLGame {
	if v == nil {
		return &TLGame{
			Data2: &Game{
				Cmd:       Cmd_Game,
				ClassName: ClassGame,
			},
		}
	}
	v1 := &TLGame{Data2: v}
	v1.Data2.ClassName = ClassGame
	return v1
}

func NewTLGeoPointEmpty(v *GeoPoint) *TLGeoPointEmpty {
	if v == nil {
		return &TLGeoPointEmpty{
			Data2: &GeoPoint{
				Cmd:       Cmd_GeoPointEmpty,
				ClassName: ClassGeoPointEmpty,
			},
		}
	}
	v1 := &TLGeoPointEmpty{Data2: v}
	v1.Data2.ClassName = ClassGeoPointEmpty
	return v1
}

func NewTLGeoPoint(v *GeoPoint) *TLGeoPoint {
	if v == nil {
		return &TLGeoPoint{
			Data2: &GeoPoint{
				Cmd:       Cmd_GeoPoint,
				ClassName: ClassGeoPoint,
			},
		}
	}
	v1 := &TLGeoPoint{Data2: v}
	v1.Data2.ClassName = ClassGeoPoint
	return v1
}

func NewTLGlobalPrivacySettings(v *GlobalPrivacySettings) *TLGlobalPrivacySettings {
	if v == nil {
		return &TLGlobalPrivacySettings{
			Data2: &GlobalPrivacySettings{
				Cmd:       Cmd_GlobalPrivacySettings,
				ClassName: ClassGlobalPrivacySettings,
			},
		}
	}
	v1 := &TLGlobalPrivacySettings{Data2: v}
	v1.Data2.ClassName = ClassGlobalPrivacySettings
	return v1
}

func NewTLGroupCallDiscarded(v *GroupCall) *TLGroupCallDiscarded {
	if v == nil {
		return &TLGroupCallDiscarded{
			Data2: &GroupCall{
				Cmd:       Cmd_GroupCallDiscarded,
				ClassName: ClassGroupCallDiscarded,
			},
		}
	}
	v1 := &TLGroupCallDiscarded{Data2: v}
	v1.Data2.ClassName = ClassGroupCallDiscarded
	return v1
}

func NewTLGroupCall(v *GroupCall) *TLGroupCall {
	if v == nil {
		return &TLGroupCall{
			Data2: &GroupCall{
				Cmd:       Cmd_GroupCall,
				ClassName: ClassGroupCall,
			},
		}
	}
	v1 := &TLGroupCall{Data2: v}
	v1.Data2.ClassName = ClassGroupCall
	return v1
}

func NewTLGroupCallParticipant(v *GroupCallParticipant) *TLGroupCallParticipant {
	if v == nil {
		return &TLGroupCallParticipant{
			Data2: &GroupCallParticipant{
				Cmd:       Cmd_GroupCallParticipant,
				ClassName: ClassGroupCallParticipant,
			},
		}
	}
	v1 := &TLGroupCallParticipant{Data2: v}
	v1.Data2.ClassName = ClassGroupCallParticipant
	return v1
}

func NewTLHelpAppChangelogEmpty(v *Help_AppChangelog) *TLHelpAppChangelogEmpty {
	if v == nil {
		return &TLHelpAppChangelogEmpty{
			Data2: &Help_AppChangelog{
				Cmd:       Cmd_HelpAppChangelogEmpty,
				ClassName: ClassHelpAppChangelogEmpty,
			},
		}
	}
	v1 := &TLHelpAppChangelogEmpty{Data2: v}
	v1.Data2.ClassName = ClassHelpAppChangelogEmpty
	return v1
}

func NewTLHelpAppChangelog(v *Help_AppChangelog) *TLHelpAppChangelog {
	if v == nil {
		return &TLHelpAppChangelog{
			Data2: &Help_AppChangelog{
				Cmd:       Cmd_HelpAppChangelog,
				ClassName: ClassHelpAppChangelog,
			},
		}
	}
	v1 := &TLHelpAppChangelog{Data2: v}
	v1.Data2.ClassName = ClassHelpAppChangelog
	return v1
}

func NewTLHelpAppUpdate(v *Help_AppUpdate) *TLHelpAppUpdate {
	if v == nil {
		return &TLHelpAppUpdate{
			Data2: &Help_AppUpdate{
				Cmd:       Cmd_HelpAppUpdate,
				ClassName: ClassHelpAppUpdate,
			},
		}
	}
	v1 := &TLHelpAppUpdate{Data2: v}
	v1.Data2.ClassName = ClassHelpAppUpdate
	return v1
}

func NewTLHelpNoAppUpdate(v *Help_AppUpdate) *TLHelpNoAppUpdate {
	if v == nil {
		return &TLHelpNoAppUpdate{
			Data2: &Help_AppUpdate{
				Cmd:       Cmd_HelpNoAppUpdate,
				ClassName: ClassHelpNoAppUpdate,
			},
		}
	}
	v1 := &TLHelpNoAppUpdate{Data2: v}
	v1.Data2.ClassName = ClassHelpNoAppUpdate
	return v1
}

func NewTLHelpConfigSimple(v *Help_ConfigSimple) *TLHelpConfigSimple {
	if v == nil {
		return &TLHelpConfigSimple{
			Data2: &Help_ConfigSimple{
				Cmd:       Cmd_HelpConfigSimple,
				ClassName: ClassHelpConfigSimple,
			},
		}
	}
	v1 := &TLHelpConfigSimple{Data2: v}
	v1.Data2.ClassName = ClassHelpConfigSimple
	return v1
}

func NewTLHelpCountriesListNotModified(v *Help_CountriesList) *TLHelpCountriesListNotModified {
	if v == nil {
		return &TLHelpCountriesListNotModified{
			Data2: &Help_CountriesList{
				Cmd:       Cmd_HelpCountriesListNotModified,
				ClassName: ClassHelpCountriesListNotModified,
			},
		}
	}
	v1 := &TLHelpCountriesListNotModified{Data2: v}
	v1.Data2.ClassName = ClassHelpCountriesListNotModified
	return v1
}

func NewTLHelpCountriesList(v *Help_CountriesList) *TLHelpCountriesList {
	if v == nil {
		return &TLHelpCountriesList{
			Data2: &Help_CountriesList{
				Cmd:       Cmd_HelpCountriesList,
				ClassName: ClassHelpCountriesList,
			},
		}
	}
	v1 := &TLHelpCountriesList{Data2: v}
	v1.Data2.ClassName = ClassHelpCountriesList
	return v1
}

func NewTLHelpCountry(v *Help_Country) *TLHelpCountry {
	if v == nil {
		return &TLHelpCountry{
			Data2: &Help_Country{
				Cmd:       Cmd_HelpCountry,
				ClassName: ClassHelpCountry,
			},
		}
	}
	v1 := &TLHelpCountry{Data2: v}
	v1.Data2.ClassName = ClassHelpCountry
	return v1
}

func NewTLHelpCountryCode(v *Help_CountryCode) *TLHelpCountryCode {
	if v == nil {
		return &TLHelpCountryCode{
			Data2: &Help_CountryCode{
				Cmd:       Cmd_HelpCountryCode,
				ClassName: ClassHelpCountryCode,
			},
		}
	}
	v1 := &TLHelpCountryCode{Data2: v}
	v1.Data2.ClassName = ClassHelpCountryCode
	return v1
}

func NewTLHelpDeepLinkInfoEmpty(v *Help_DeepLinkInfo) *TLHelpDeepLinkInfoEmpty {
	if v == nil {
		return &TLHelpDeepLinkInfoEmpty{
			Data2: &Help_DeepLinkInfo{
				Cmd:       Cmd_HelpDeepLinkInfoEmpty,
				ClassName: ClassHelpDeepLinkInfoEmpty,
			},
		}
	}
	v1 := &TLHelpDeepLinkInfoEmpty{Data2: v}
	v1.Data2.ClassName = ClassHelpDeepLinkInfoEmpty
	return v1
}

func NewTLHelpDeepLinkInfo(v *Help_DeepLinkInfo) *TLHelpDeepLinkInfo {
	if v == nil {
		return &TLHelpDeepLinkInfo{
			Data2: &Help_DeepLinkInfo{
				Cmd:       Cmd_HelpDeepLinkInfo,
				ClassName: ClassHelpDeepLinkInfo,
			},
		}
	}
	v1 := &TLHelpDeepLinkInfo{Data2: v}
	v1.Data2.ClassName = ClassHelpDeepLinkInfo
	return v1
}

func NewTLHelpInviteText(v *Help_InviteText) *TLHelpInviteText {
	if v == nil {
		return &TLHelpInviteText{
			Data2: &Help_InviteText{
				Cmd:       Cmd_HelpInviteText,
				ClassName: ClassHelpInviteText,
			},
		}
	}
	v1 := &TLHelpInviteText{Data2: v}
	v1.Data2.ClassName = ClassHelpInviteText
	return v1
}

func NewTLHelpPassportConfigNotModified(v *Help_PassportConfig) *TLHelpPassportConfigNotModified {
	if v == nil {
		return &TLHelpPassportConfigNotModified{
			Data2: &Help_PassportConfig{
				Cmd:       Cmd_HelpPassportConfigNotModified,
				ClassName: ClassHelpPassportConfigNotModified,
			},
		}
	}
	v1 := &TLHelpPassportConfigNotModified{Data2: v}
	v1.Data2.ClassName = ClassHelpPassportConfigNotModified
	return v1
}

func NewTLHelpPassportConfig(v *Help_PassportConfig) *TLHelpPassportConfig {
	if v == nil {
		return &TLHelpPassportConfig{
			Data2: &Help_PassportConfig{
				Cmd:       Cmd_HelpPassportConfig,
				ClassName: ClassHelpPassportConfig,
			},
		}
	}
	v1 := &TLHelpPassportConfig{Data2: v}
	v1.Data2.ClassName = ClassHelpPassportConfig
	return v1
}

func NewTLHelpPromoDataEmpty(v *Help_PromoData) *TLHelpPromoDataEmpty {
	if v == nil {
		return &TLHelpPromoDataEmpty{
			Data2: &Help_PromoData{
				Cmd:       Cmd_HelpPromoDataEmpty,
				ClassName: ClassHelpPromoDataEmpty,
			},
		}
	}
	v1 := &TLHelpPromoDataEmpty{Data2: v}
	v1.Data2.ClassName = ClassHelpPromoDataEmpty
	return v1
}

func NewTLHelpPromoData(v *Help_PromoData) *TLHelpPromoData {
	if v == nil {
		return &TLHelpPromoData{
			Data2: &Help_PromoData{
				Cmd:       Cmd_HelpPromoData,
				ClassName: ClassHelpPromoData,
			},
		}
	}
	v1 := &TLHelpPromoData{Data2: v}
	v1.Data2.ClassName = ClassHelpPromoData
	return v1
}

func NewTLHelpProxyDataEmpty(v *Help_ProxyData) *TLHelpProxyDataEmpty {
	if v == nil {
		return &TLHelpProxyDataEmpty{
			Data2: &Help_ProxyData{
				Cmd:       Cmd_HelpProxyDataEmpty,
				ClassName: ClassHelpProxyDataEmpty,
			},
		}
	}
	v1 := &TLHelpProxyDataEmpty{Data2: v}
	v1.Data2.ClassName = ClassHelpProxyDataEmpty
	return v1
}

func NewTLHelpProxyDataPromo(v *Help_ProxyData) *TLHelpProxyDataPromo {
	if v == nil {
		return &TLHelpProxyDataPromo{
			Data2: &Help_ProxyData{
				Cmd:       Cmd_HelpProxyDataPromo,
				ClassName: ClassHelpProxyDataPromo,
			},
		}
	}
	v1 := &TLHelpProxyDataPromo{Data2: v}
	v1.Data2.ClassName = ClassHelpProxyDataPromo
	return v1
}

func NewTLHelpRecentMeUrls(v *Help_RecentMeUrls) *TLHelpRecentMeUrls {
	if v == nil {
		return &TLHelpRecentMeUrls{
			Data2: &Help_RecentMeUrls{
				Cmd:       Cmd_HelpRecentMeUrls,
				ClassName: ClassHelpRecentMeUrls,
			},
		}
	}
	v1 := &TLHelpRecentMeUrls{Data2: v}
	v1.Data2.ClassName = ClassHelpRecentMeUrls
	return v1
}

func NewTLHelpSupport(v *Help_Support) *TLHelpSupport {
	if v == nil {
		return &TLHelpSupport{
			Data2: &Help_Support{
				Cmd:       Cmd_HelpSupport,
				ClassName: ClassHelpSupport,
			},
		}
	}
	v1 := &TLHelpSupport{Data2: v}
	v1.Data2.ClassName = ClassHelpSupport
	return v1
}

func NewTLHelpSupportName(v *Help_SupportName) *TLHelpSupportName {
	if v == nil {
		return &TLHelpSupportName{
			Data2: &Help_SupportName{
				Cmd:       Cmd_HelpSupportName,
				ClassName: ClassHelpSupportName,
			},
		}
	}
	v1 := &TLHelpSupportName{Data2: v}
	v1.Data2.ClassName = ClassHelpSupportName
	return v1
}

func NewTLHelpTermsOfService(v *Help_TermsOfService) *TLHelpTermsOfService {
	if v == nil {
		return &TLHelpTermsOfService{
			Data2: &Help_TermsOfService{
				Cmd:       Cmd_HelpTermsOfService,
				ClassName: ClassHelpTermsOfService,
			},
		}
	}
	v1 := &TLHelpTermsOfService{Data2: v}
	v1.Data2.ClassName = ClassHelpTermsOfService
	return v1
}

func NewTLHelpTermsOfServiceUpdateEmpty(v *Help_TermsOfServiceUpdate) *TLHelpTermsOfServiceUpdateEmpty {
	if v == nil {
		return &TLHelpTermsOfServiceUpdateEmpty{
			Data2: &Help_TermsOfServiceUpdate{
				Cmd:       Cmd_HelpTermsOfServiceUpdateEmpty,
				ClassName: ClassHelpTermsOfServiceUpdateEmpty,
			},
		}
	}
	v1 := &TLHelpTermsOfServiceUpdateEmpty{Data2: v}
	v1.Data2.ClassName = ClassHelpTermsOfServiceUpdateEmpty
	return v1
}

func NewTLHelpTermsOfServiceUpdate(v *Help_TermsOfServiceUpdate) *TLHelpTermsOfServiceUpdate {
	if v == nil {
		return &TLHelpTermsOfServiceUpdate{
			Data2: &Help_TermsOfServiceUpdate{
				Cmd:       Cmd_HelpTermsOfServiceUpdate,
				ClassName: ClassHelpTermsOfServiceUpdate,
			},
		}
	}
	v1 := &TLHelpTermsOfServiceUpdate{Data2: v}
	v1.Data2.ClassName = ClassHelpTermsOfServiceUpdate
	return v1
}

func NewTLHelpUserInfoEmpty(v *Help_UserInfo) *TLHelpUserInfoEmpty {
	if v == nil {
		return &TLHelpUserInfoEmpty{
			Data2: &Help_UserInfo{
				Cmd:       Cmd_HelpUserInfoEmpty,
				ClassName: ClassHelpUserInfoEmpty,
			},
		}
	}
	v1 := &TLHelpUserInfoEmpty{Data2: v}
	v1.Data2.ClassName = ClassHelpUserInfoEmpty
	return v1
}

func NewTLHelpUserInfo(v *Help_UserInfo) *TLHelpUserInfo {
	if v == nil {
		return &TLHelpUserInfo{
			Data2: &Help_UserInfo{
				Cmd:       Cmd_HelpUserInfo,
				ClassName: ClassHelpUserInfo,
			},
		}
	}
	v1 := &TLHelpUserInfo{Data2: v}
	v1.Data2.ClassName = ClassHelpUserInfo
	return v1
}

func NewTLHighScore(v *HighScore) *TLHighScore {
	if v == nil {
		return &TLHighScore{
			Data2: &HighScore{
				Cmd:       Cmd_HighScore,
				ClassName: ClassHighScore,
			},
		}
	}
	v1 := &TLHighScore{Data2: v}
	v1.Data2.ClassName = ClassHighScore
	return v1
}

func NewTLHttpWait(v *HttpWait) *TLHttpWait {
	if v == nil {
		return &TLHttpWait{
			Data2: &HttpWait{
				Cmd:       Cmd_HttpWait,
				ClassName: ClassHttpWait,
			},
		}
	}
	v1 := &TLHttpWait{Data2: v}
	v1.Data2.ClassName = ClassHttpWait
	return v1
}

func NewTLImportedContact(v *ImportedContact) *TLImportedContact {
	if v == nil {
		return &TLImportedContact{
			Data2: &ImportedContact{
				Cmd:       Cmd_ImportedContact,
				ClassName: ClassImportedContact,
			},
		}
	}
	v1 := &TLImportedContact{Data2: v}
	v1.Data2.ClassName = ClassImportedContact
	return v1
}

func NewTLInlineBotSwitchPM(v *InlineBotSwitchPM) *TLInlineBotSwitchPM {
	if v == nil {
		return &TLInlineBotSwitchPM{
			Data2: &InlineBotSwitchPM{
				Cmd:       Cmd_InlineBotSwitchPM,
				ClassName: ClassInlineBotSwitchPM,
			},
		}
	}
	v1 := &TLInlineBotSwitchPM{Data2: v}
	v1.Data2.ClassName = ClassInlineBotSwitchPM
	return v1
}

func NewTLInlineQueryPeerTypeSameBotPM(v *InlineQueryPeerType) *TLInlineQueryPeerTypeSameBotPM {
	if v == nil {
		return &TLInlineQueryPeerTypeSameBotPM{
			Data2: &InlineQueryPeerType{
				Cmd:       Cmd_InlineQueryPeerTypeSameBotPM,
				ClassName: ClassInlineQueryPeerTypeSameBotPM,
			},
		}
	}
	v1 := &TLInlineQueryPeerTypeSameBotPM{Data2: v}
	v1.Data2.ClassName = ClassInlineQueryPeerTypeSameBotPM
	return v1
}

func NewTLInlineQueryPeerTypePM(v *InlineQueryPeerType) *TLInlineQueryPeerTypePM {
	if v == nil {
		return &TLInlineQueryPeerTypePM{
			Data2: &InlineQueryPeerType{
				Cmd:       Cmd_InlineQueryPeerTypePM,
				ClassName: ClassInlineQueryPeerTypePM,
			},
		}
	}
	v1 := &TLInlineQueryPeerTypePM{Data2: v}
	v1.Data2.ClassName = ClassInlineQueryPeerTypePM
	return v1
}

func NewTLInlineQueryPeerTypeChat(v *InlineQueryPeerType) *TLInlineQueryPeerTypeChat {
	if v == nil {
		return &TLInlineQueryPeerTypeChat{
			Data2: &InlineQueryPeerType{
				Cmd:       Cmd_InlineQueryPeerTypeChat,
				ClassName: ClassInlineQueryPeerTypeChat,
			},
		}
	}
	v1 := &TLInlineQueryPeerTypeChat{Data2: v}
	v1.Data2.ClassName = ClassInlineQueryPeerTypeChat
	return v1
}

func NewTLInlineQueryPeerTypeMegagroup(v *InlineQueryPeerType) *TLInlineQueryPeerTypeMegagroup {
	if v == nil {
		return &TLInlineQueryPeerTypeMegagroup{
			Data2: &InlineQueryPeerType{
				Cmd:       Cmd_InlineQueryPeerTypeMegagroup,
				ClassName: ClassInlineQueryPeerTypeMegagroup,
			},
		}
	}
	v1 := &TLInlineQueryPeerTypeMegagroup{Data2: v}
	v1.Data2.ClassName = ClassInlineQueryPeerTypeMegagroup
	return v1
}

func NewTLInlineQueryPeerTypeBroadcast(v *InlineQueryPeerType) *TLInlineQueryPeerTypeBroadcast {
	if v == nil {
		return &TLInlineQueryPeerTypeBroadcast{
			Data2: &InlineQueryPeerType{
				Cmd:       Cmd_InlineQueryPeerTypeBroadcast,
				ClassName: ClassInlineQueryPeerTypeBroadcast,
			},
		}
	}
	v1 := &TLInlineQueryPeerTypeBroadcast{Data2: v}
	v1.Data2.ClassName = ClassInlineQueryPeerTypeBroadcast
	return v1
}

func NewTLInputAppEvent(v *InputAppEvent) *TLInputAppEvent {
	if v == nil {
		return &TLInputAppEvent{
			Data2: &InputAppEvent{
				Cmd:       Cmd_InputAppEvent,
				ClassName: ClassInputAppEvent,
			},
		}
	}
	v1 := &TLInputAppEvent{Data2: v}
	v1.Data2.ClassName = ClassInputAppEvent
	return v1
}

func NewTLInputBotInlineMessageMediaAuto(v *InputBotInlineMessage) *TLInputBotInlineMessageMediaAuto {
	if v == nil {
		return &TLInputBotInlineMessageMediaAuto{
			Data2: &InputBotInlineMessage{
				Cmd:       Cmd_InputBotInlineMessageMediaAuto,
				ClassName: ClassInputBotInlineMessageMediaAuto,
			},
		}
	}
	v1 := &TLInputBotInlineMessageMediaAuto{Data2: v}
	v1.Data2.ClassName = ClassInputBotInlineMessageMediaAuto
	return v1
}

func NewTLInputBotInlineMessageText(v *InputBotInlineMessage) *TLInputBotInlineMessageText {
	if v == nil {
		return &TLInputBotInlineMessageText{
			Data2: &InputBotInlineMessage{
				Cmd:       Cmd_InputBotInlineMessageText,
				ClassName: ClassInputBotInlineMessageText,
			},
		}
	}
	v1 := &TLInputBotInlineMessageText{Data2: v}
	v1.Data2.ClassName = ClassInputBotInlineMessageText
	return v1
}

func NewTLInputBotInlineMessageMediaGeo(v *InputBotInlineMessage) *TLInputBotInlineMessageMediaGeo {
	if v == nil {
		return &TLInputBotInlineMessageMediaGeo{
			Data2: &InputBotInlineMessage{
				Cmd:       Cmd_InputBotInlineMessageMediaGeo,
				ClassName: ClassInputBotInlineMessageMediaGeo,
			},
		}
	}
	v1 := &TLInputBotInlineMessageMediaGeo{Data2: v}
	v1.Data2.ClassName = ClassInputBotInlineMessageMediaGeo
	return v1
}

func NewTLInputBotInlineMessageMediaVenue(v *InputBotInlineMessage) *TLInputBotInlineMessageMediaVenue {
	if v == nil {
		return &TLInputBotInlineMessageMediaVenue{
			Data2: &InputBotInlineMessage{
				Cmd:       Cmd_InputBotInlineMessageMediaVenue,
				ClassName: ClassInputBotInlineMessageMediaVenue,
			},
		}
	}
	v1 := &TLInputBotInlineMessageMediaVenue{Data2: v}
	v1.Data2.ClassName = ClassInputBotInlineMessageMediaVenue
	return v1
}

func NewTLInputBotInlineMessageMediaContact(v *InputBotInlineMessage) *TLInputBotInlineMessageMediaContact {
	if v == nil {
		return &TLInputBotInlineMessageMediaContact{
			Data2: &InputBotInlineMessage{
				Cmd:       Cmd_InputBotInlineMessageMediaContact,
				ClassName: ClassInputBotInlineMessageMediaContact,
			},
		}
	}
	v1 := &TLInputBotInlineMessageMediaContact{Data2: v}
	v1.Data2.ClassName = ClassInputBotInlineMessageMediaContact
	return v1
}

func NewTLInputBotInlineMessageGame(v *InputBotInlineMessage) *TLInputBotInlineMessageGame {
	if v == nil {
		return &TLInputBotInlineMessageGame{
			Data2: &InputBotInlineMessage{
				Cmd:       Cmd_InputBotInlineMessageGame,
				ClassName: ClassInputBotInlineMessageGame,
			},
		}
	}
	v1 := &TLInputBotInlineMessageGame{Data2: v}
	v1.Data2.ClassName = ClassInputBotInlineMessageGame
	return v1
}

func NewTLInputBotInlineMessageID(v *InputBotInlineMessageID) *TLInputBotInlineMessageID {
	if v == nil {
		return &TLInputBotInlineMessageID{
			Data2: &InputBotInlineMessageID{
				Cmd:       Cmd_InputBotInlineMessageID,
				ClassName: ClassInputBotInlineMessageID,
			},
		}
	}
	v1 := &TLInputBotInlineMessageID{Data2: v}
	v1.Data2.ClassName = ClassInputBotInlineMessageID
	return v1
}

func NewTLInputBotInlineResult(v *InputBotInlineResult) *TLInputBotInlineResult {
	if v == nil {
		return &TLInputBotInlineResult{
			Data2: &InputBotInlineResult{
				Cmd:       Cmd_InputBotInlineResult,
				ClassName: ClassInputBotInlineResult,
			},
		}
	}
	v1 := &TLInputBotInlineResult{Data2: v}
	v1.Data2.ClassName = ClassInputBotInlineResult
	return v1
}

func NewTLInputBotInlineResultPhoto(v *InputBotInlineResult) *TLInputBotInlineResultPhoto {
	if v == nil {
		return &TLInputBotInlineResultPhoto{
			Data2: &InputBotInlineResult{
				Cmd:       Cmd_InputBotInlineResultPhoto,
				ClassName: ClassInputBotInlineResultPhoto,
			},
		}
	}
	v1 := &TLInputBotInlineResultPhoto{Data2: v}
	v1.Data2.ClassName = ClassInputBotInlineResultPhoto
	return v1
}

func NewTLInputBotInlineResultDocument(v *InputBotInlineResult) *TLInputBotInlineResultDocument {
	if v == nil {
		return &TLInputBotInlineResultDocument{
			Data2: &InputBotInlineResult{
				Cmd:       Cmd_InputBotInlineResultDocument,
				ClassName: ClassInputBotInlineResultDocument,
			},
		}
	}
	v1 := &TLInputBotInlineResultDocument{Data2: v}
	v1.Data2.ClassName = ClassInputBotInlineResultDocument
	return v1
}

func NewTLInputBotInlineResultGame(v *InputBotInlineResult) *TLInputBotInlineResultGame {
	if v == nil {
		return &TLInputBotInlineResultGame{
			Data2: &InputBotInlineResult{
				Cmd:       Cmd_InputBotInlineResultGame,
				ClassName: ClassInputBotInlineResultGame,
			},
		}
	}
	v1 := &TLInputBotInlineResultGame{Data2: v}
	v1.Data2.ClassName = ClassInputBotInlineResultGame
	return v1
}

func NewTLInputChannelEmpty(v *InputChannel) *TLInputChannelEmpty {
	if v == nil {
		return &TLInputChannelEmpty{
			Data2: &InputChannel{
				Cmd:       Cmd_InputChannelEmpty,
				ClassName: ClassInputChannelEmpty,
			},
		}
	}
	v1 := &TLInputChannelEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputChannelEmpty
	return v1
}

func NewTLInputChannel(v *InputChannel) *TLInputChannel {
	if v == nil {
		return &TLInputChannel{
			Data2: &InputChannel{
				Cmd:       Cmd_InputChannel,
				ClassName: ClassInputChannel,
			},
		}
	}
	v1 := &TLInputChannel{Data2: v}
	v1.Data2.ClassName = ClassInputChannel
	return v1
}

func NewTLInputChannelFromMessage(v *InputChannel) *TLInputChannelFromMessage {
	if v == nil {
		return &TLInputChannelFromMessage{
			Data2: &InputChannel{
				Cmd:       Cmd_InputChannelFromMessage,
				ClassName: ClassInputChannelFromMessage,
			},
		}
	}
	v1 := &TLInputChannelFromMessage{Data2: v}
	v1.Data2.ClassName = ClassInputChannelFromMessage
	return v1
}

func NewTLInputChatPhotoEmpty(v *InputChatPhoto) *TLInputChatPhotoEmpty {
	if v == nil {
		return &TLInputChatPhotoEmpty{
			Data2: &InputChatPhoto{
				Cmd:       Cmd_InputChatPhotoEmpty,
				ClassName: ClassInputChatPhotoEmpty,
			},
		}
	}
	v1 := &TLInputChatPhotoEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputChatPhotoEmpty
	return v1
}

func NewTLInputChatUploadedPhoto(v *InputChatPhoto) *TLInputChatUploadedPhoto {
	if v == nil {
		return &TLInputChatUploadedPhoto{
			Data2: &InputChatPhoto{
				Cmd:       Cmd_InputChatUploadedPhoto,
				ClassName: ClassInputChatUploadedPhoto,
			},
		}
	}
	v1 := &TLInputChatUploadedPhoto{Data2: v}
	v1.Data2.ClassName = ClassInputChatUploadedPhoto
	return v1
}

func NewTLInputChatPhoto(v *InputChatPhoto) *TLInputChatPhoto {
	if v == nil {
		return &TLInputChatPhoto{
			Data2: &InputChatPhoto{
				Cmd:       Cmd_InputChatPhoto,
				ClassName: ClassInputChatPhoto,
			},
		}
	}
	v1 := &TLInputChatPhoto{Data2: v}
	v1.Data2.ClassName = ClassInputChatPhoto
	return v1
}

func NewTLInputCheckPasswordEmpty(v *InputCheckPasswordSRP) *TLInputCheckPasswordEmpty {
	if v == nil {
		return &TLInputCheckPasswordEmpty{
			Data2: &InputCheckPasswordSRP{
				Cmd:       Cmd_InputCheckPasswordEmpty,
				ClassName: ClassInputCheckPasswordEmpty,
			},
		}
	}
	v1 := &TLInputCheckPasswordEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputCheckPasswordEmpty
	return v1
}

func NewTLInputCheckPasswordSRP(v *InputCheckPasswordSRP) *TLInputCheckPasswordSRP {
	if v == nil {
		return &TLInputCheckPasswordSRP{
			Data2: &InputCheckPasswordSRP{
				Cmd:       Cmd_InputCheckPasswordSRP,
				ClassName: ClassInputCheckPasswordSRP,
			},
		}
	}
	v1 := &TLInputCheckPasswordSRP{Data2: v}
	v1.Data2.ClassName = ClassInputCheckPasswordSRP
	return v1
}

func NewTLInputClientProxy(v *InputClientProxy) *TLInputClientProxy {
	if v == nil {
		return &TLInputClientProxy{
			Data2: &InputClientProxy{
				Cmd:       Cmd_InputClientProxy,
				ClassName: ClassInputClientProxy,
			},
		}
	}
	v1 := &TLInputClientProxy{Data2: v}
	v1.Data2.ClassName = ClassInputClientProxy
	return v1
}

func NewTLInputPhoneContact(v *InputContact) *TLInputPhoneContact {
	if v == nil {
		return &TLInputPhoneContact{
			Data2: &InputContact{
				Cmd:       Cmd_InputPhoneContact,
				ClassName: ClassInputPhoneContact,
			},
		}
	}
	v1 := &TLInputPhoneContact{Data2: v}
	v1.Data2.ClassName = ClassInputPhoneContact
	return v1
}

func NewTLInputDialogPeer(v *InputDialogPeer) *TLInputDialogPeer {
	if v == nil {
		return &TLInputDialogPeer{
			Data2: &InputDialogPeer{
				Cmd:       Cmd_InputDialogPeer,
				ClassName: ClassInputDialogPeer,
			},
		}
	}
	v1 := &TLInputDialogPeer{Data2: v}
	v1.Data2.ClassName = ClassInputDialogPeer
	return v1
}

func NewTLInputDialogPeerFolder(v *InputDialogPeer) *TLInputDialogPeerFolder {
	if v == nil {
		return &TLInputDialogPeerFolder{
			Data2: &InputDialogPeer{
				Cmd:       Cmd_InputDialogPeerFolder,
				ClassName: ClassInputDialogPeerFolder,
			},
		}
	}
	v1 := &TLInputDialogPeerFolder{Data2: v}
	v1.Data2.ClassName = ClassInputDialogPeerFolder
	return v1
}

func NewTLInputDocumentEmpty(v *InputDocument) *TLInputDocumentEmpty {
	if v == nil {
		return &TLInputDocumentEmpty{
			Data2: &InputDocument{
				Cmd:       Cmd_InputDocumentEmpty,
				ClassName: ClassInputDocumentEmpty,
			},
		}
	}
	v1 := &TLInputDocumentEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputDocumentEmpty
	return v1
}

func NewTLInputDocument(v *InputDocument) *TLInputDocument {
	if v == nil {
		return &TLInputDocument{
			Data2: &InputDocument{
				Cmd:       Cmd_InputDocument,
				ClassName: ClassInputDocument,
			},
		}
	}
	v1 := &TLInputDocument{Data2: v}
	v1.Data2.ClassName = ClassInputDocument
	return v1
}

func NewTLInputEncryptedChat(v *InputEncryptedChat) *TLInputEncryptedChat {
	if v == nil {
		return &TLInputEncryptedChat{
			Data2: &InputEncryptedChat{
				Cmd:       Cmd_InputEncryptedChat,
				ClassName: ClassInputEncryptedChat,
			},
		}
	}
	v1 := &TLInputEncryptedChat{Data2: v}
	v1.Data2.ClassName = ClassInputEncryptedChat
	return v1
}

func NewTLInputEncryptedFileEmpty(v *InputEncryptedFile) *TLInputEncryptedFileEmpty {
	if v == nil {
		return &TLInputEncryptedFileEmpty{
			Data2: &InputEncryptedFile{
				Cmd:       Cmd_InputEncryptedFileEmpty,
				ClassName: ClassInputEncryptedFileEmpty,
			},
		}
	}
	v1 := &TLInputEncryptedFileEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputEncryptedFileEmpty
	return v1
}

func NewTLInputEncryptedFileUploaded(v *InputEncryptedFile) *TLInputEncryptedFileUploaded {
	if v == nil {
		return &TLInputEncryptedFileUploaded{
			Data2: &InputEncryptedFile{
				Cmd:       Cmd_InputEncryptedFileUploaded,
				ClassName: ClassInputEncryptedFileUploaded,
			},
		}
	}
	v1 := &TLInputEncryptedFileUploaded{Data2: v}
	v1.Data2.ClassName = ClassInputEncryptedFileUploaded
	return v1
}

func NewTLInputEncryptedFile(v *InputEncryptedFile) *TLInputEncryptedFile {
	if v == nil {
		return &TLInputEncryptedFile{
			Data2: &InputEncryptedFile{
				Cmd:       Cmd_InputEncryptedFile,
				ClassName: ClassInputEncryptedFile,
			},
		}
	}
	v1 := &TLInputEncryptedFile{Data2: v}
	v1.Data2.ClassName = ClassInputEncryptedFile
	return v1
}

func NewTLInputEncryptedFileBigUploaded(v *InputEncryptedFile) *TLInputEncryptedFileBigUploaded {
	if v == nil {
		return &TLInputEncryptedFileBigUploaded{
			Data2: &InputEncryptedFile{
				Cmd:       Cmd_InputEncryptedFileBigUploaded,
				ClassName: ClassInputEncryptedFileBigUploaded,
			},
		}
	}
	v1 := &TLInputEncryptedFileBigUploaded{Data2: v}
	v1.Data2.ClassName = ClassInputEncryptedFileBigUploaded
	return v1
}

func NewTLInputFile(v *InputFile) *TLInputFile {
	if v == nil {
		return &TLInputFile{
			Data2: &InputFile{
				Cmd:       Cmd_InputFile,
				ClassName: ClassInputFile,
			},
		}
	}
	v1 := &TLInputFile{Data2: v}
	v1.Data2.ClassName = ClassInputFile
	return v1
}

func NewTLInputFileBig(v *InputFile) *TLInputFileBig {
	if v == nil {
		return &TLInputFileBig{
			Data2: &InputFile{
				Cmd:       Cmd_InputFileBig,
				ClassName: ClassInputFileBig,
			},
		}
	}
	v1 := &TLInputFileBig{Data2: v}
	v1.Data2.ClassName = ClassInputFileBig
	return v1
}

func NewTLInputFileLocation(v *InputFileLocation) *TLInputFileLocation {
	if v == nil {
		return &TLInputFileLocation{
			Data2: &InputFileLocation{
				Cmd:       Cmd_InputFileLocation,
				ClassName: ClassInputFileLocation,
			},
		}
	}
	v1 := &TLInputFileLocation{Data2: v}
	v1.Data2.ClassName = ClassInputFileLocation
	return v1
}

func NewTLInputEncryptedFileLocation(v *InputFileLocation) *TLInputEncryptedFileLocation {
	if v == nil {
		return &TLInputEncryptedFileLocation{
			Data2: &InputFileLocation{
				Cmd:       Cmd_InputEncryptedFileLocation,
				ClassName: ClassInputEncryptedFileLocation,
			},
		}
	}
	v1 := &TLInputEncryptedFileLocation{Data2: v}
	v1.Data2.ClassName = ClassInputEncryptedFileLocation
	return v1
}

func NewTLInputDocumentFileLocation(v *InputFileLocation) *TLInputDocumentFileLocation {
	if v == nil {
		return &TLInputDocumentFileLocation{
			Data2: &InputFileLocation{
				Cmd:       Cmd_InputDocumentFileLocation,
				ClassName: ClassInputDocumentFileLocation,
			},
		}
	}
	v1 := &TLInputDocumentFileLocation{Data2: v}
	v1.Data2.ClassName = ClassInputDocumentFileLocation
	return v1
}

func NewTLInputSecureFileLocation(v *InputFileLocation) *TLInputSecureFileLocation {
	if v == nil {
		return &TLInputSecureFileLocation{
			Data2: &InputFileLocation{
				Cmd:       Cmd_InputSecureFileLocation,
				ClassName: ClassInputSecureFileLocation,
			},
		}
	}
	v1 := &TLInputSecureFileLocation{Data2: v}
	v1.Data2.ClassName = ClassInputSecureFileLocation
	return v1
}

func NewTLInputTakeoutFileLocation(v *InputFileLocation) *TLInputTakeoutFileLocation {
	if v == nil {
		return &TLInputTakeoutFileLocation{
			Data2: &InputFileLocation{
				Cmd:       Cmd_InputTakeoutFileLocation,
				ClassName: ClassInputTakeoutFileLocation,
			},
		}
	}
	v1 := &TLInputTakeoutFileLocation{Data2: v}
	v1.Data2.ClassName = ClassInputTakeoutFileLocation
	return v1
}

func NewTLInputPhotoFileLocation(v *InputFileLocation) *TLInputPhotoFileLocation {
	if v == nil {
		return &TLInputPhotoFileLocation{
			Data2: &InputFileLocation{
				Cmd:       Cmd_InputPhotoFileLocation,
				ClassName: ClassInputPhotoFileLocation,
			},
		}
	}
	v1 := &TLInputPhotoFileLocation{Data2: v}
	v1.Data2.ClassName = ClassInputPhotoFileLocation
	return v1
}

func NewTLInputPeerPhotoFileLocation(v *InputFileLocation) *TLInputPeerPhotoFileLocation {
	if v == nil {
		return &TLInputPeerPhotoFileLocation{
			Data2: &InputFileLocation{
				Cmd:       Cmd_InputPeerPhotoFileLocation,
				ClassName: ClassInputPeerPhotoFileLocation,
			},
		}
	}
	v1 := &TLInputPeerPhotoFileLocation{Data2: v}
	v1.Data2.ClassName = ClassInputPeerPhotoFileLocation
	return v1
}

func NewTLInputStickerSetThumb(v *InputFileLocation) *TLInputStickerSetThumb {
	if v == nil {
		return &TLInputStickerSetThumb{
			Data2: &InputFileLocation{
				Cmd:       Cmd_InputStickerSetThumb,
				ClassName: ClassInputStickerSetThumb,
			},
		}
	}
	v1 := &TLInputStickerSetThumb{Data2: v}
	v1.Data2.ClassName = ClassInputStickerSetThumb
	return v1
}

func NewTLInputPhotoLegacyFileLocation(v *InputFileLocation) *TLInputPhotoLegacyFileLocation {
	if v == nil {
		return &TLInputPhotoLegacyFileLocation{
			Data2: &InputFileLocation{
				Cmd:       Cmd_InputPhotoLegacyFileLocation,
				ClassName: ClassInputPhotoLegacyFileLocation,
			},
		}
	}
	v1 := &TLInputPhotoLegacyFileLocation{Data2: v}
	v1.Data2.ClassName = ClassInputPhotoLegacyFileLocation
	return v1
}

func NewTLInputFolderPeer(v *InputFolderPeer) *TLInputFolderPeer {
	if v == nil {
		return &TLInputFolderPeer{
			Data2: &InputFolderPeer{
				Cmd:       Cmd_InputFolderPeer,
				ClassName: ClassInputFolderPeer,
			},
		}
	}
	v1 := &TLInputFolderPeer{Data2: v}
	v1.Data2.ClassName = ClassInputFolderPeer
	return v1
}

func NewTLInputGameID(v *InputGame) *TLInputGameID {
	if v == nil {
		return &TLInputGameID{
			Data2: &InputGame{
				Cmd:       Cmd_InputGameID,
				ClassName: ClassInputGameID,
			},
		}
	}
	v1 := &TLInputGameID{Data2: v}
	v1.Data2.ClassName = ClassInputGameID
	return v1
}

func NewTLInputGameShortName(v *InputGame) *TLInputGameShortName {
	if v == nil {
		return &TLInputGameShortName{
			Data2: &InputGame{
				Cmd:       Cmd_InputGameShortName,
				ClassName: ClassInputGameShortName,
			},
		}
	}
	v1 := &TLInputGameShortName{Data2: v}
	v1.Data2.ClassName = ClassInputGameShortName
	return v1
}

func NewTLInputGeoPointEmpty(v *InputGeoPoint) *TLInputGeoPointEmpty {
	if v == nil {
		return &TLInputGeoPointEmpty{
			Data2: &InputGeoPoint{
				Cmd:       Cmd_InputGeoPointEmpty,
				ClassName: ClassInputGeoPointEmpty,
			},
		}
	}
	v1 := &TLInputGeoPointEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputGeoPointEmpty
	return v1
}

func NewTLInputGeoPoint(v *InputGeoPoint) *TLInputGeoPoint {
	if v == nil {
		return &TLInputGeoPoint{
			Data2: &InputGeoPoint{
				Cmd:       Cmd_InputGeoPoint,
				ClassName: ClassInputGeoPoint,
			},
		}
	}
	v1 := &TLInputGeoPoint{Data2: v}
	v1.Data2.ClassName = ClassInputGeoPoint
	return v1
}

func NewTLInputGroupCall(v *InputGroupCall) *TLInputGroupCall {
	if v == nil {
		return &TLInputGroupCall{
			Data2: &InputGroupCall{
				Cmd:       Cmd_InputGroupCall,
				ClassName: ClassInputGroupCall,
			},
		}
	}
	v1 := &TLInputGroupCall{Data2: v}
	v1.Data2.ClassName = ClassInputGroupCall
	return v1
}

func NewTLInputMediaEmpty(v *InputMedia) *TLInputMediaEmpty {
	if v == nil {
		return &TLInputMediaEmpty{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaEmpty,
				ClassName: ClassInputMediaEmpty,
			},
		}
	}
	v1 := &TLInputMediaEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputMediaEmpty
	return v1
}

func NewTLInputMediaUploadedPhoto(v *InputMedia) *TLInputMediaUploadedPhoto {
	if v == nil {
		return &TLInputMediaUploadedPhoto{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaUploadedPhoto,
				ClassName: ClassInputMediaUploadedPhoto,
			},
		}
	}
	v1 := &TLInputMediaUploadedPhoto{Data2: v}
	v1.Data2.ClassName = ClassInputMediaUploadedPhoto
	return v1
}

func NewTLInputMediaPhoto(v *InputMedia) *TLInputMediaPhoto {
	if v == nil {
		return &TLInputMediaPhoto{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaPhoto,
				ClassName: ClassInputMediaPhoto,
			},
		}
	}
	v1 := &TLInputMediaPhoto{Data2: v}
	v1.Data2.ClassName = ClassInputMediaPhoto
	return v1
}

func NewTLInputMediaGeoPoint(v *InputMedia) *TLInputMediaGeoPoint {
	if v == nil {
		return &TLInputMediaGeoPoint{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaGeoPoint,
				ClassName: ClassInputMediaGeoPoint,
			},
		}
	}
	v1 := &TLInputMediaGeoPoint{Data2: v}
	v1.Data2.ClassName = ClassInputMediaGeoPoint
	return v1
}

func NewTLInputMediaContact(v *InputMedia) *TLInputMediaContact {
	if v == nil {
		return &TLInputMediaContact{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaContact,
				ClassName: ClassInputMediaContact,
			},
		}
	}
	v1 := &TLInputMediaContact{Data2: v}
	v1.Data2.ClassName = ClassInputMediaContact
	return v1
}

func NewTLInputMediaUploadedDocument(v *InputMedia) *TLInputMediaUploadedDocument {
	if v == nil {
		return &TLInputMediaUploadedDocument{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaUploadedDocument,
				ClassName: ClassInputMediaUploadedDocument,
			},
		}
	}
	v1 := &TLInputMediaUploadedDocument{Data2: v}
	v1.Data2.ClassName = ClassInputMediaUploadedDocument
	return v1
}

func NewTLInputMediaDocument(v *InputMedia) *TLInputMediaDocument {
	if v == nil {
		return &TLInputMediaDocument{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaDocument,
				ClassName: ClassInputMediaDocument,
			},
		}
	}
	v1 := &TLInputMediaDocument{Data2: v}
	v1.Data2.ClassName = ClassInputMediaDocument
	return v1
}

func NewTLInputMediaVenue(v *InputMedia) *TLInputMediaVenue {
	if v == nil {
		return &TLInputMediaVenue{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaVenue,
				ClassName: ClassInputMediaVenue,
			},
		}
	}
	v1 := &TLInputMediaVenue{Data2: v}
	v1.Data2.ClassName = ClassInputMediaVenue
	return v1
}

func NewTLInputMediaGifExternal(v *InputMedia) *TLInputMediaGifExternal {
	if v == nil {
		return &TLInputMediaGifExternal{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaGifExternal,
				ClassName: ClassInputMediaGifExternal,
			},
		}
	}
	v1 := &TLInputMediaGifExternal{Data2: v}
	v1.Data2.ClassName = ClassInputMediaGifExternal
	return v1
}

func NewTLInputMediaPhotoExternal(v *InputMedia) *TLInputMediaPhotoExternal {
	if v == nil {
		return &TLInputMediaPhotoExternal{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaPhotoExternal,
				ClassName: ClassInputMediaPhotoExternal,
			},
		}
	}
	v1 := &TLInputMediaPhotoExternal{Data2: v}
	v1.Data2.ClassName = ClassInputMediaPhotoExternal
	return v1
}

func NewTLInputMediaDocumentExternal(v *InputMedia) *TLInputMediaDocumentExternal {
	if v == nil {
		return &TLInputMediaDocumentExternal{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaDocumentExternal,
				ClassName: ClassInputMediaDocumentExternal,
			},
		}
	}
	v1 := &TLInputMediaDocumentExternal{Data2: v}
	v1.Data2.ClassName = ClassInputMediaDocumentExternal
	return v1
}

func NewTLInputMediaGame(v *InputMedia) *TLInputMediaGame {
	if v == nil {
		return &TLInputMediaGame{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaGame,
				ClassName: ClassInputMediaGame,
			},
		}
	}
	v1 := &TLInputMediaGame{Data2: v}
	v1.Data2.ClassName = ClassInputMediaGame
	return v1
}

func NewTLInputMediaInvoice(v *InputMedia) *TLInputMediaInvoice {
	if v == nil {
		return &TLInputMediaInvoice{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaInvoice,
				ClassName: ClassInputMediaInvoice,
			},
		}
	}
	v1 := &TLInputMediaInvoice{Data2: v}
	v1.Data2.ClassName = ClassInputMediaInvoice
	return v1
}

func NewTLInputMediaUploadedThumbDocument(v *InputMedia) *TLInputMediaUploadedThumbDocument {
	if v == nil {
		return &TLInputMediaUploadedThumbDocument{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaUploadedThumbDocument,
				ClassName: ClassInputMediaUploadedThumbDocument,
			},
		}
	}
	v1 := &TLInputMediaUploadedThumbDocument{Data2: v}
	v1.Data2.ClassName = ClassInputMediaUploadedThumbDocument
	return v1
}

func NewTLInputMediaGeoLive(v *InputMedia) *TLInputMediaGeoLive {
	if v == nil {
		return &TLInputMediaGeoLive{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaGeoLive,
				ClassName: ClassInputMediaGeoLive,
			},
		}
	}
	v1 := &TLInputMediaGeoLive{Data2: v}
	v1.Data2.ClassName = ClassInputMediaGeoLive
	return v1
}

func NewTLInputMediaPoll(v *InputMedia) *TLInputMediaPoll {
	if v == nil {
		return &TLInputMediaPoll{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaPoll,
				ClassName: ClassInputMediaPoll,
			},
		}
	}
	v1 := &TLInputMediaPoll{Data2: v}
	v1.Data2.ClassName = ClassInputMediaPoll
	return v1
}

func NewTLInputMediaDice(v *InputMedia) *TLInputMediaDice {
	if v == nil {
		return &TLInputMediaDice{
			Data2: &InputMedia{
				Cmd:       Cmd_InputMediaDice,
				ClassName: ClassInputMediaDice,
			},
		}
	}
	v1 := &TLInputMediaDice{Data2: v}
	v1.Data2.ClassName = ClassInputMediaDice
	return v1
}

func NewTLInputMessageID(v *InputMessage) *TLInputMessageID {
	if v == nil {
		return &TLInputMessageID{
			Data2: &InputMessage{
				Cmd:       Cmd_InputMessageID,
				ClassName: ClassInputMessageID,
			},
		}
	}
	v1 := &TLInputMessageID{Data2: v}
	v1.Data2.ClassName = ClassInputMessageID
	return v1
}

func NewTLInputMessageReplyTo(v *InputMessage) *TLInputMessageReplyTo {
	if v == nil {
		return &TLInputMessageReplyTo{
			Data2: &InputMessage{
				Cmd:       Cmd_InputMessageReplyTo,
				ClassName: ClassInputMessageReplyTo,
			},
		}
	}
	v1 := &TLInputMessageReplyTo{Data2: v}
	v1.Data2.ClassName = ClassInputMessageReplyTo
	return v1
}

func NewTLInputMessagePinned(v *InputMessage) *TLInputMessagePinned {
	if v == nil {
		return &TLInputMessagePinned{
			Data2: &InputMessage{
				Cmd:       Cmd_InputMessagePinned,
				ClassName: ClassInputMessagePinned,
			},
		}
	}
	v1 := &TLInputMessagePinned{Data2: v}
	v1.Data2.ClassName = ClassInputMessagePinned
	return v1
}

func NewTLInputMessageCallbackQuery(v *InputMessage) *TLInputMessageCallbackQuery {
	if v == nil {
		return &TLInputMessageCallbackQuery{
			Data2: &InputMessage{
				Cmd:       Cmd_InputMessageCallbackQuery,
				ClassName: ClassInputMessageCallbackQuery,
			},
		}
	}
	v1 := &TLInputMessageCallbackQuery{Data2: v}
	v1.Data2.ClassName = ClassInputMessageCallbackQuery
	return v1
}

func NewTLInputNotifyPeer(v *InputNotifyPeer) *TLInputNotifyPeer {
	if v == nil {
		return &TLInputNotifyPeer{
			Data2: &InputNotifyPeer{
				Cmd:       Cmd_InputNotifyPeer,
				ClassName: ClassInputNotifyPeer,
			},
		}
	}
	v1 := &TLInputNotifyPeer{Data2: v}
	v1.Data2.ClassName = ClassInputNotifyPeer
	return v1
}

func NewTLInputNotifyUsers(v *InputNotifyPeer) *TLInputNotifyUsers {
	if v == nil {
		return &TLInputNotifyUsers{
			Data2: &InputNotifyPeer{
				Cmd:       Cmd_InputNotifyUsers,
				ClassName: ClassInputNotifyUsers,
			},
		}
	}
	v1 := &TLInputNotifyUsers{Data2: v}
	v1.Data2.ClassName = ClassInputNotifyUsers
	return v1
}

func NewTLInputNotifyChats(v *InputNotifyPeer) *TLInputNotifyChats {
	if v == nil {
		return &TLInputNotifyChats{
			Data2: &InputNotifyPeer{
				Cmd:       Cmd_InputNotifyChats,
				ClassName: ClassInputNotifyChats,
			},
		}
	}
	v1 := &TLInputNotifyChats{Data2: v}
	v1.Data2.ClassName = ClassInputNotifyChats
	return v1
}

func NewTLInputNotifyAll(v *InputNotifyPeer) *TLInputNotifyAll {
	if v == nil {
		return &TLInputNotifyAll{
			Data2: &InputNotifyPeer{
				Cmd:       Cmd_InputNotifyAll,
				ClassName: ClassInputNotifyAll,
			},
		}
	}
	v1 := &TLInputNotifyAll{Data2: v}
	v1.Data2.ClassName = ClassInputNotifyAll
	return v1
}

func NewTLInputNotifyBroadcasts(v *InputNotifyPeer) *TLInputNotifyBroadcasts {
	if v == nil {
		return &TLInputNotifyBroadcasts{
			Data2: &InputNotifyPeer{
				Cmd:       Cmd_InputNotifyBroadcasts,
				ClassName: ClassInputNotifyBroadcasts,
			},
		}
	}
	v1 := &TLInputNotifyBroadcasts{Data2: v}
	v1.Data2.ClassName = ClassInputNotifyBroadcasts
	return v1
}

func NewTLInputPaymentCredentialsSaved(v *InputPaymentCredentials) *TLInputPaymentCredentialsSaved {
	if v == nil {
		return &TLInputPaymentCredentialsSaved{
			Data2: &InputPaymentCredentials{
				Cmd:       Cmd_InputPaymentCredentialsSaved,
				ClassName: ClassInputPaymentCredentialsSaved,
			},
		}
	}
	v1 := &TLInputPaymentCredentialsSaved{Data2: v}
	v1.Data2.ClassName = ClassInputPaymentCredentialsSaved
	return v1
}

func NewTLInputPaymentCredentials(v *InputPaymentCredentials) *TLInputPaymentCredentials {
	if v == nil {
		return &TLInputPaymentCredentials{
			Data2: &InputPaymentCredentials{
				Cmd:       Cmd_InputPaymentCredentials,
				ClassName: ClassInputPaymentCredentials,
			},
		}
	}
	v1 := &TLInputPaymentCredentials{Data2: v}
	v1.Data2.ClassName = ClassInputPaymentCredentials
	return v1
}

func NewTLInputPaymentCredentialsApplePay(v *InputPaymentCredentials) *TLInputPaymentCredentialsApplePay {
	if v == nil {
		return &TLInputPaymentCredentialsApplePay{
			Data2: &InputPaymentCredentials{
				Cmd:       Cmd_InputPaymentCredentialsApplePay,
				ClassName: ClassInputPaymentCredentialsApplePay,
			},
		}
	}
	v1 := &TLInputPaymentCredentialsApplePay{Data2: v}
	v1.Data2.ClassName = ClassInputPaymentCredentialsApplePay
	return v1
}

func NewTLInputPaymentCredentialsAndroidPay(v *InputPaymentCredentials) *TLInputPaymentCredentialsAndroidPay {
	if v == nil {
		return &TLInputPaymentCredentialsAndroidPay{
			Data2: &InputPaymentCredentials{
				Cmd:       Cmd_InputPaymentCredentialsAndroidPay,
				ClassName: ClassInputPaymentCredentialsAndroidPay,
			},
		}
	}
	v1 := &TLInputPaymentCredentialsAndroidPay{Data2: v}
	v1.Data2.ClassName = ClassInputPaymentCredentialsAndroidPay
	return v1
}

func NewTLInputPeerEmpty(v *InputPeer) *TLInputPeerEmpty {
	if v == nil {
		return &TLInputPeerEmpty{
			Data2: &InputPeer{
				Cmd:       Cmd_InputPeerEmpty,
				ClassName: ClassInputPeerEmpty,
			},
		}
	}
	v1 := &TLInputPeerEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputPeerEmpty
	return v1
}

func NewTLInputPeerSelf(v *InputPeer) *TLInputPeerSelf {
	if v == nil {
		return &TLInputPeerSelf{
			Data2: &InputPeer{
				Cmd:       Cmd_InputPeerSelf,
				ClassName: ClassInputPeerSelf,
			},
		}
	}
	v1 := &TLInputPeerSelf{Data2: v}
	v1.Data2.ClassName = ClassInputPeerSelf
	return v1
}

func NewTLInputPeerChat(v *InputPeer) *TLInputPeerChat {
	if v == nil {
		return &TLInputPeerChat{
			Data2: &InputPeer{
				Cmd:       Cmd_InputPeerChat,
				ClassName: ClassInputPeerChat,
			},
		}
	}
	v1 := &TLInputPeerChat{Data2: v}
	v1.Data2.ClassName = ClassInputPeerChat
	return v1
}

func NewTLInputPeerUser(v *InputPeer) *TLInputPeerUser {
	if v == nil {
		return &TLInputPeerUser{
			Data2: &InputPeer{
				Cmd:       Cmd_InputPeerUser,
				ClassName: ClassInputPeerUser,
			},
		}
	}
	v1 := &TLInputPeerUser{Data2: v}
	v1.Data2.ClassName = ClassInputPeerUser
	return v1
}

func NewTLInputPeerChannel(v *InputPeer) *TLInputPeerChannel {
	if v == nil {
		return &TLInputPeerChannel{
			Data2: &InputPeer{
				Cmd:       Cmd_InputPeerChannel,
				ClassName: ClassInputPeerChannel,
			},
		}
	}
	v1 := &TLInputPeerChannel{Data2: v}
	v1.Data2.ClassName = ClassInputPeerChannel
	return v1
}

func NewTLInputPeerUserFromMessage(v *InputPeer) *TLInputPeerUserFromMessage {
	if v == nil {
		return &TLInputPeerUserFromMessage{
			Data2: &InputPeer{
				Cmd:       Cmd_InputPeerUserFromMessage,
				ClassName: ClassInputPeerUserFromMessage,
			},
		}
	}
	v1 := &TLInputPeerUserFromMessage{Data2: v}
	v1.Data2.ClassName = ClassInputPeerUserFromMessage
	return v1
}

func NewTLInputPeerChannelFromMessage(v *InputPeer) *TLInputPeerChannelFromMessage {
	if v == nil {
		return &TLInputPeerChannelFromMessage{
			Data2: &InputPeer{
				Cmd:       Cmd_InputPeerChannelFromMessage,
				ClassName: ClassInputPeerChannelFromMessage,
			},
		}
	}
	v1 := &TLInputPeerChannelFromMessage{Data2: v}
	v1.Data2.ClassName = ClassInputPeerChannelFromMessage
	return v1
}

func NewTLInputPeerNotifyEventsEmpty(v *InputPeerNotifyEvents) *TLInputPeerNotifyEventsEmpty {
	if v == nil {
		return &TLInputPeerNotifyEventsEmpty{
			Data2: &InputPeerNotifyEvents{
				Cmd:       Cmd_InputPeerNotifyEventsEmpty,
				ClassName: ClassInputPeerNotifyEventsEmpty,
			},
		}
	}
	v1 := &TLInputPeerNotifyEventsEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputPeerNotifyEventsEmpty
	return v1
}

func NewTLInputPeerNotifyEventsAll(v *InputPeerNotifyEvents) *TLInputPeerNotifyEventsAll {
	if v == nil {
		return &TLInputPeerNotifyEventsAll{
			Data2: &InputPeerNotifyEvents{
				Cmd:       Cmd_InputPeerNotifyEventsAll,
				ClassName: ClassInputPeerNotifyEventsAll,
			},
		}
	}
	v1 := &TLInputPeerNotifyEventsAll{Data2: v}
	v1.Data2.ClassName = ClassInputPeerNotifyEventsAll
	return v1
}

func NewTLInputPeerNotifySettings(v *InputPeerNotifySettings) *TLInputPeerNotifySettings {
	if v == nil {
		return &TLInputPeerNotifySettings{
			Data2: &InputPeerNotifySettings{
				Cmd:       Cmd_InputPeerNotifySettings,
				ClassName: ClassInputPeerNotifySettings,
			},
		}
	}
	v1 := &TLInputPeerNotifySettings{Data2: v}
	v1.Data2.ClassName = ClassInputPeerNotifySettings
	return v1
}

func NewTLInputPhoneCall(v *InputPhoneCall) *TLInputPhoneCall {
	if v == nil {
		return &TLInputPhoneCall{
			Data2: &InputPhoneCall{
				Cmd:       Cmd_InputPhoneCall,
				ClassName: ClassInputPhoneCall,
			},
		}
	}
	v1 := &TLInputPhoneCall{Data2: v}
	v1.Data2.ClassName = ClassInputPhoneCall
	return v1
}

func NewTLInputPhotoEmpty(v *InputPhoto) *TLInputPhotoEmpty {
	if v == nil {
		return &TLInputPhotoEmpty{
			Data2: &InputPhoto{
				Cmd:       Cmd_InputPhotoEmpty,
				ClassName: ClassInputPhotoEmpty,
			},
		}
	}
	v1 := &TLInputPhotoEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputPhotoEmpty
	return v1
}

func NewTLInputPhoto(v *InputPhoto) *TLInputPhoto {
	if v == nil {
		return &TLInputPhoto{
			Data2: &InputPhoto{
				Cmd:       Cmd_InputPhoto,
				ClassName: ClassInputPhoto,
			},
		}
	}
	v1 := &TLInputPhoto{Data2: v}
	v1.Data2.ClassName = ClassInputPhoto
	return v1
}

func NewTLInputPhotoCrop(v *InputPhotoCrop) *TLInputPhotoCrop {
	if v == nil {
		return &TLInputPhotoCrop{
			Data2: &InputPhotoCrop{
				Cmd:       Cmd_InputPhotoCrop,
				ClassName: ClassInputPhotoCrop,
			},
		}
	}
	v1 := &TLInputPhotoCrop{Data2: v}
	v1.Data2.ClassName = ClassInputPhotoCrop
	return v1
}

func NewTLInputPhotoCropAuto(v *InputPhotoCrop) *TLInputPhotoCropAuto {
	if v == nil {
		return &TLInputPhotoCropAuto{
			Data2: &InputPhotoCrop{
				Cmd:       Cmd_InputPhotoCropAuto,
				ClassName: ClassInputPhotoCropAuto,
			},
		}
	}
	v1 := &TLInputPhotoCropAuto{Data2: v}
	v1.Data2.ClassName = ClassInputPhotoCropAuto
	return v1
}

func NewTLInputPrivacyKeyStatusTimestamp(v *InputPrivacyKey) *TLInputPrivacyKeyStatusTimestamp {
	if v == nil {
		return &TLInputPrivacyKeyStatusTimestamp{
			Data2: &InputPrivacyKey{
				Cmd:       Cmd_InputPrivacyKeyStatusTimestamp,
				ClassName: ClassInputPrivacyKeyStatusTimestamp,
			},
		}
	}
	v1 := &TLInputPrivacyKeyStatusTimestamp{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyKeyStatusTimestamp
	return v1
}

func NewTLInputPrivacyKeyChatInvite(v *InputPrivacyKey) *TLInputPrivacyKeyChatInvite {
	if v == nil {
		return &TLInputPrivacyKeyChatInvite{
			Data2: &InputPrivacyKey{
				Cmd:       Cmd_InputPrivacyKeyChatInvite,
				ClassName: ClassInputPrivacyKeyChatInvite,
			},
		}
	}
	v1 := &TLInputPrivacyKeyChatInvite{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyKeyChatInvite
	return v1
}

func NewTLInputPrivacyKeyPhoneCall(v *InputPrivacyKey) *TLInputPrivacyKeyPhoneCall {
	if v == nil {
		return &TLInputPrivacyKeyPhoneCall{
			Data2: &InputPrivacyKey{
				Cmd:       Cmd_InputPrivacyKeyPhoneCall,
				ClassName: ClassInputPrivacyKeyPhoneCall,
			},
		}
	}
	v1 := &TLInputPrivacyKeyPhoneCall{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyKeyPhoneCall
	return v1
}

func NewTLInputPrivacyKeyPhoneP2P(v *InputPrivacyKey) *TLInputPrivacyKeyPhoneP2P {
	if v == nil {
		return &TLInputPrivacyKeyPhoneP2P{
			Data2: &InputPrivacyKey{
				Cmd:       Cmd_InputPrivacyKeyPhoneP2P,
				ClassName: ClassInputPrivacyKeyPhoneP2P,
			},
		}
	}
	v1 := &TLInputPrivacyKeyPhoneP2P{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyKeyPhoneP2P
	return v1
}

func NewTLInputPrivacyKeyForwards(v *InputPrivacyKey) *TLInputPrivacyKeyForwards {
	if v == nil {
		return &TLInputPrivacyKeyForwards{
			Data2: &InputPrivacyKey{
				Cmd:       Cmd_InputPrivacyKeyForwards,
				ClassName: ClassInputPrivacyKeyForwards,
			},
		}
	}
	v1 := &TLInputPrivacyKeyForwards{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyKeyForwards
	return v1
}

func NewTLInputPrivacyKeyProfilePhoto(v *InputPrivacyKey) *TLInputPrivacyKeyProfilePhoto {
	if v == nil {
		return &TLInputPrivacyKeyProfilePhoto{
			Data2: &InputPrivacyKey{
				Cmd:       Cmd_InputPrivacyKeyProfilePhoto,
				ClassName: ClassInputPrivacyKeyProfilePhoto,
			},
		}
	}
	v1 := &TLInputPrivacyKeyProfilePhoto{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyKeyProfilePhoto
	return v1
}

func NewTLInputPrivacyKeyPhoneNumber(v *InputPrivacyKey) *TLInputPrivacyKeyPhoneNumber {
	if v == nil {
		return &TLInputPrivacyKeyPhoneNumber{
			Data2: &InputPrivacyKey{
				Cmd:       Cmd_InputPrivacyKeyPhoneNumber,
				ClassName: ClassInputPrivacyKeyPhoneNumber,
			},
		}
	}
	v1 := &TLInputPrivacyKeyPhoneNumber{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyKeyPhoneNumber
	return v1
}

func NewTLInputPrivacyKeyAddedByPhone(v *InputPrivacyKey) *TLInputPrivacyKeyAddedByPhone {
	if v == nil {
		return &TLInputPrivacyKeyAddedByPhone{
			Data2: &InputPrivacyKey{
				Cmd:       Cmd_InputPrivacyKeyAddedByPhone,
				ClassName: ClassInputPrivacyKeyAddedByPhone,
			},
		}
	}
	v1 := &TLInputPrivacyKeyAddedByPhone{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyKeyAddedByPhone
	return v1
}

func NewTLInputPrivacyValueAllowContacts(v *InputPrivacyRule) *TLInputPrivacyValueAllowContacts {
	if v == nil {
		return &TLInputPrivacyValueAllowContacts{
			Data2: &InputPrivacyRule{
				Cmd:       Cmd_InputPrivacyValueAllowContacts,
				ClassName: ClassInputPrivacyValueAllowContacts,
			},
		}
	}
	v1 := &TLInputPrivacyValueAllowContacts{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyValueAllowContacts
	return v1
}

func NewTLInputPrivacyValueAllowAll(v *InputPrivacyRule) *TLInputPrivacyValueAllowAll {
	if v == nil {
		return &TLInputPrivacyValueAllowAll{
			Data2: &InputPrivacyRule{
				Cmd:       Cmd_InputPrivacyValueAllowAll,
				ClassName: ClassInputPrivacyValueAllowAll,
			},
		}
	}
	v1 := &TLInputPrivacyValueAllowAll{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyValueAllowAll
	return v1
}

func NewTLInputPrivacyValueAllowUsers(v *InputPrivacyRule) *TLInputPrivacyValueAllowUsers {
	if v == nil {
		return &TLInputPrivacyValueAllowUsers{
			Data2: &InputPrivacyRule{
				Cmd:       Cmd_InputPrivacyValueAllowUsers,
				ClassName: ClassInputPrivacyValueAllowUsers,
			},
		}
	}
	v1 := &TLInputPrivacyValueAllowUsers{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyValueAllowUsers
	return v1
}

func NewTLInputPrivacyValueDisallowContacts(v *InputPrivacyRule) *TLInputPrivacyValueDisallowContacts {
	if v == nil {
		return &TLInputPrivacyValueDisallowContacts{
			Data2: &InputPrivacyRule{
				Cmd:       Cmd_InputPrivacyValueDisallowContacts,
				ClassName: ClassInputPrivacyValueDisallowContacts,
			},
		}
	}
	v1 := &TLInputPrivacyValueDisallowContacts{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyValueDisallowContacts
	return v1
}

func NewTLInputPrivacyValueDisallowAll(v *InputPrivacyRule) *TLInputPrivacyValueDisallowAll {
	if v == nil {
		return &TLInputPrivacyValueDisallowAll{
			Data2: &InputPrivacyRule{
				Cmd:       Cmd_InputPrivacyValueDisallowAll,
				ClassName: ClassInputPrivacyValueDisallowAll,
			},
		}
	}
	v1 := &TLInputPrivacyValueDisallowAll{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyValueDisallowAll
	return v1
}

func NewTLInputPrivacyValueDisallowUsers(v *InputPrivacyRule) *TLInputPrivacyValueDisallowUsers {
	if v == nil {
		return &TLInputPrivacyValueDisallowUsers{
			Data2: &InputPrivacyRule{
				Cmd:       Cmd_InputPrivacyValueDisallowUsers,
				ClassName: ClassInputPrivacyValueDisallowUsers,
			},
		}
	}
	v1 := &TLInputPrivacyValueDisallowUsers{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyValueDisallowUsers
	return v1
}

func NewTLInputPrivacyValueAllowChatParticipants(v *InputPrivacyRule) *TLInputPrivacyValueAllowChatParticipants {
	if v == nil {
		return &TLInputPrivacyValueAllowChatParticipants{
			Data2: &InputPrivacyRule{
				Cmd:       Cmd_InputPrivacyValueAllowChatParticipants,
				ClassName: ClassInputPrivacyValueAllowChatParticipants,
			},
		}
	}
	v1 := &TLInputPrivacyValueAllowChatParticipants{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyValueAllowChatParticipants
	return v1
}

func NewTLInputPrivacyValueDisallowChatParticipants(v *InputPrivacyRule) *TLInputPrivacyValueDisallowChatParticipants {
	if v == nil {
		return &TLInputPrivacyValueDisallowChatParticipants{
			Data2: &InputPrivacyRule{
				Cmd:       Cmd_InputPrivacyValueDisallowChatParticipants,
				ClassName: ClassInputPrivacyValueDisallowChatParticipants,
			},
		}
	}
	v1 := &TLInputPrivacyValueDisallowChatParticipants{Data2: v}
	v1.Data2.ClassName = ClassInputPrivacyValueDisallowChatParticipants
	return v1
}

func NewTLInputSecureFileUploaded(v *InputSecureFile) *TLInputSecureFileUploaded {
	if v == nil {
		return &TLInputSecureFileUploaded{
			Data2: &InputSecureFile{
				Cmd:       Cmd_InputSecureFileUploaded,
				ClassName: ClassInputSecureFileUploaded,
			},
		}
	}
	v1 := &TLInputSecureFileUploaded{Data2: v}
	v1.Data2.ClassName = ClassInputSecureFileUploaded
	return v1
}

func NewTLInputSecureFile(v *InputSecureFile) *TLInputSecureFile {
	if v == nil {
		return &TLInputSecureFile{
			Data2: &InputSecureFile{
				Cmd:       Cmd_InputSecureFile,
				ClassName: ClassInputSecureFile,
			},
		}
	}
	v1 := &TLInputSecureFile{Data2: v}
	v1.Data2.ClassName = ClassInputSecureFile
	return v1
}

func NewTLInputSecureValue(v *InputSecureValue) *TLInputSecureValue {
	if v == nil {
		return &TLInputSecureValue{
			Data2: &InputSecureValue{
				Cmd:       Cmd_InputSecureValue,
				ClassName: ClassInputSecureValue,
			},
		}
	}
	v1 := &TLInputSecureValue{Data2: v}
	v1.Data2.ClassName = ClassInputSecureValue
	return v1
}

func NewTLInputSingleMedia(v *InputSingleMedia) *TLInputSingleMedia {
	if v == nil {
		return &TLInputSingleMedia{
			Data2: &InputSingleMedia{
				Cmd:       Cmd_InputSingleMedia,
				ClassName: ClassInputSingleMedia,
			},
		}
	}
	v1 := &TLInputSingleMedia{Data2: v}
	v1.Data2.ClassName = ClassInputSingleMedia
	return v1
}

func NewTLInputStickerSetEmpty(v *InputStickerSet) *TLInputStickerSetEmpty {
	if v == nil {
		return &TLInputStickerSetEmpty{
			Data2: &InputStickerSet{
				Cmd:       Cmd_InputStickerSetEmpty,
				ClassName: ClassInputStickerSetEmpty,
			},
		}
	}
	v1 := &TLInputStickerSetEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputStickerSetEmpty
	return v1
}

func NewTLInputStickerSetID(v *InputStickerSet) *TLInputStickerSetID {
	if v == nil {
		return &TLInputStickerSetID{
			Data2: &InputStickerSet{
				Cmd:       Cmd_InputStickerSetID,
				ClassName: ClassInputStickerSetID,
			},
		}
	}
	v1 := &TLInputStickerSetID{Data2: v}
	v1.Data2.ClassName = ClassInputStickerSetID
	return v1
}

func NewTLInputStickerSetShortName(v *InputStickerSet) *TLInputStickerSetShortName {
	if v == nil {
		return &TLInputStickerSetShortName{
			Data2: &InputStickerSet{
				Cmd:       Cmd_InputStickerSetShortName,
				ClassName: ClassInputStickerSetShortName,
			},
		}
	}
	v1 := &TLInputStickerSetShortName{Data2: v}
	v1.Data2.ClassName = ClassInputStickerSetShortName
	return v1
}

func NewTLInputStickerSetAnimatedEmoji(v *InputStickerSet) *TLInputStickerSetAnimatedEmoji {
	if v == nil {
		return &TLInputStickerSetAnimatedEmoji{
			Data2: &InputStickerSet{
				Cmd:       Cmd_InputStickerSetAnimatedEmoji,
				ClassName: ClassInputStickerSetAnimatedEmoji,
			},
		}
	}
	v1 := &TLInputStickerSetAnimatedEmoji{Data2: v}
	v1.Data2.ClassName = ClassInputStickerSetAnimatedEmoji
	return v1
}

func NewTLInputStickerSetDice(v *InputStickerSet) *TLInputStickerSetDice {
	if v == nil {
		return &TLInputStickerSetDice{
			Data2: &InputStickerSet{
				Cmd:       Cmd_InputStickerSetDice,
				ClassName: ClassInputStickerSetDice,
			},
		}
	}
	v1 := &TLInputStickerSetDice{Data2: v}
	v1.Data2.ClassName = ClassInputStickerSetDice
	return v1
}

func NewTLInputStickerSetItem(v *InputStickerSetItem) *TLInputStickerSetItem {
	if v == nil {
		return &TLInputStickerSetItem{
			Data2: &InputStickerSetItem{
				Cmd:       Cmd_InputStickerSetItem,
				ClassName: ClassInputStickerSetItem,
			},
		}
	}
	v1 := &TLInputStickerSetItem{Data2: v}
	v1.Data2.ClassName = ClassInputStickerSetItem
	return v1
}

func NewTLInputStickeredMediaPhoto(v *InputStickeredMedia) *TLInputStickeredMediaPhoto {
	if v == nil {
		return &TLInputStickeredMediaPhoto{
			Data2: &InputStickeredMedia{
				Cmd:       Cmd_InputStickeredMediaPhoto,
				ClassName: ClassInputStickeredMediaPhoto,
			},
		}
	}
	v1 := &TLInputStickeredMediaPhoto{Data2: v}
	v1.Data2.ClassName = ClassInputStickeredMediaPhoto
	return v1
}

func NewTLInputStickeredMediaDocument(v *InputStickeredMedia) *TLInputStickeredMediaDocument {
	if v == nil {
		return &TLInputStickeredMediaDocument{
			Data2: &InputStickeredMedia{
				Cmd:       Cmd_InputStickeredMediaDocument,
				ClassName: ClassInputStickeredMediaDocument,
			},
		}
	}
	v1 := &TLInputStickeredMediaDocument{Data2: v}
	v1.Data2.ClassName = ClassInputStickeredMediaDocument
	return v1
}

func NewTLInputTheme(v *InputTheme) *TLInputTheme {
	if v == nil {
		return &TLInputTheme{
			Data2: &InputTheme{
				Cmd:       Cmd_InputTheme,
				ClassName: ClassInputTheme,
			},
		}
	}
	v1 := &TLInputTheme{Data2: v}
	v1.Data2.ClassName = ClassInputTheme
	return v1
}

func NewTLInputThemeSlug(v *InputTheme) *TLInputThemeSlug {
	if v == nil {
		return &TLInputThemeSlug{
			Data2: &InputTheme{
				Cmd:       Cmd_InputThemeSlug,
				ClassName: ClassInputThemeSlug,
			},
		}
	}
	v1 := &TLInputThemeSlug{Data2: v}
	v1.Data2.ClassName = ClassInputThemeSlug
	return v1
}

func NewTLInputThemeSettings(v *InputThemeSettings) *TLInputThemeSettings {
	if v == nil {
		return &TLInputThemeSettings{
			Data2: &InputThemeSettings{
				Cmd:       Cmd_InputThemeSettings,
				ClassName: ClassInputThemeSettings,
			},
		}
	}
	v1 := &TLInputThemeSettings{Data2: v}
	v1.Data2.ClassName = ClassInputThemeSettings
	return v1
}

func NewTLInputUserEmpty(v *InputUser) *TLInputUserEmpty {
	if v == nil {
		return &TLInputUserEmpty{
			Data2: &InputUser{
				Cmd:       Cmd_InputUserEmpty,
				ClassName: ClassInputUserEmpty,
			},
		}
	}
	v1 := &TLInputUserEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputUserEmpty
	return v1
}

func NewTLInputUserSelf(v *InputUser) *TLInputUserSelf {
	if v == nil {
		return &TLInputUserSelf{
			Data2: &InputUser{
				Cmd:       Cmd_InputUserSelf,
				ClassName: ClassInputUserSelf,
			},
		}
	}
	v1 := &TLInputUserSelf{Data2: v}
	v1.Data2.ClassName = ClassInputUserSelf
	return v1
}

func NewTLInputUser(v *InputUser) *TLInputUser {
	if v == nil {
		return &TLInputUser{
			Data2: &InputUser{
				Cmd:       Cmd_InputUser,
				ClassName: ClassInputUser,
			},
		}
	}
	v1 := &TLInputUser{Data2: v}
	v1.Data2.ClassName = ClassInputUser
	return v1
}

func NewTLInputUserFromMessage(v *InputUser) *TLInputUserFromMessage {
	if v == nil {
		return &TLInputUserFromMessage{
			Data2: &InputUser{
				Cmd:       Cmd_InputUserFromMessage,
				ClassName: ClassInputUserFromMessage,
			},
		}
	}
	v1 := &TLInputUserFromMessage{Data2: v}
	v1.Data2.ClassName = ClassInputUserFromMessage
	return v1
}

func NewTLInputWallPaper(v *InputWallPaper) *TLInputWallPaper {
	if v == nil {
		return &TLInputWallPaper{
			Data2: &InputWallPaper{
				Cmd:       Cmd_InputWallPaper,
				ClassName: ClassInputWallPaper,
			},
		}
	}
	v1 := &TLInputWallPaper{Data2: v}
	v1.Data2.ClassName = ClassInputWallPaper
	return v1
}

func NewTLInputWallPaperSlug(v *InputWallPaper) *TLInputWallPaperSlug {
	if v == nil {
		return &TLInputWallPaperSlug{
			Data2: &InputWallPaper{
				Cmd:       Cmd_InputWallPaperSlug,
				ClassName: ClassInputWallPaperSlug,
			},
		}
	}
	v1 := &TLInputWallPaperSlug{Data2: v}
	v1.Data2.ClassName = ClassInputWallPaperSlug
	return v1
}

func NewTLInputWallPaperNoFile(v *InputWallPaper) *TLInputWallPaperNoFile {
	if v == nil {
		return &TLInputWallPaperNoFile{
			Data2: &InputWallPaper{
				Cmd:       Cmd_InputWallPaperNoFile,
				ClassName: ClassInputWallPaperNoFile,
			},
		}
	}
	v1 := &TLInputWallPaperNoFile{Data2: v}
	v1.Data2.ClassName = ClassInputWallPaperNoFile
	return v1
}

func NewTLInputWebDocument(v *InputWebDocument) *TLInputWebDocument {
	if v == nil {
		return &TLInputWebDocument{
			Data2: &InputWebDocument{
				Cmd:       Cmd_InputWebDocument,
				ClassName: ClassInputWebDocument,
			},
		}
	}
	v1 := &TLInputWebDocument{Data2: v}
	v1.Data2.ClassName = ClassInputWebDocument
	return v1
}

func NewTLInputWebFileLocation(v *InputWebFileLocation) *TLInputWebFileLocation {
	if v == nil {
		return &TLInputWebFileLocation{
			Data2: &InputWebFileLocation{
				Cmd:       Cmd_InputWebFileLocation,
				ClassName: ClassInputWebFileLocation,
			},
		}
	}
	v1 := &TLInputWebFileLocation{Data2: v}
	v1.Data2.ClassName = ClassInputWebFileLocation
	return v1
}

func NewTLInputWebFileGeoPointLocation(v *InputWebFileLocation) *TLInputWebFileGeoPointLocation {
	if v == nil {
		return &TLInputWebFileGeoPointLocation{
			Data2: &InputWebFileLocation{
				Cmd:       Cmd_InputWebFileGeoPointLocation,
				ClassName: ClassInputWebFileGeoPointLocation,
			},
		}
	}
	v1 := &TLInputWebFileGeoPointLocation{Data2: v}
	v1.Data2.ClassName = ClassInputWebFileGeoPointLocation
	return v1
}

func NewTLInvoice(v *Invoice) *TLInvoice {
	if v == nil {
		return &TLInvoice{
			Data2: &Invoice{
				Cmd:       Cmd_Invoice,
				ClassName: ClassInvoice,
			},
		}
	}
	v1 := &TLInvoice{Data2: v}
	v1.Data2.ClassName = ClassInvoice
	return v1
}

func NewTLIpPort(v *IpPort) *TLIpPort {
	if v == nil {
		return &TLIpPort{
			Data2: &IpPort{
				Cmd:       Cmd_IpPort,
				ClassName: ClassIpPort,
			},
		}
	}
	v1 := &TLIpPort{Data2: v}
	v1.Data2.ClassName = ClassIpPort
	return v1
}

func NewTLIpPortSecret(v *IpPort) *TLIpPortSecret {
	if v == nil {
		return &TLIpPortSecret{
			Data2: &IpPort{
				Cmd:       Cmd_IpPortSecret,
				ClassName: ClassIpPortSecret,
			},
		}
	}
	v1 := &TLIpPortSecret{Data2: v}
	v1.Data2.ClassName = ClassIpPortSecret
	return v1
}

func NewTLJsonObjectValue(v *JSONObjectValue) *TLJsonObjectValue {
	if v == nil {
		return &TLJsonObjectValue{
			Data2: &JSONObjectValue{
				Cmd:       Cmd_JsonObjectValue,
				ClassName: ClassJsonObjectValue,
			},
		}
	}
	v1 := &TLJsonObjectValue{Data2: v}
	v1.Data2.ClassName = ClassJsonObjectValue
	return v1
}

func NewTLJsonNull(v *JSONValue) *TLJsonNull {
	if v == nil {
		return &TLJsonNull{
			Data2: &JSONValue{
				Cmd:       Cmd_JsonNull,
				ClassName: ClassJsonNull,
			},
		}
	}
	v1 := &TLJsonNull{Data2: v}
	v1.Data2.ClassName = ClassJsonNull
	return v1
}

func NewTLJsonBool(v *JSONValue) *TLJsonBool {
	if v == nil {
		return &TLJsonBool{
			Data2: &JSONValue{
				Cmd:       Cmd_JsonBool,
				ClassName: ClassJsonBool,
			},
		}
	}
	v1 := &TLJsonBool{Data2: v}
	v1.Data2.ClassName = ClassJsonBool
	return v1
}

func NewTLJsonNumber(v *JSONValue) *TLJsonNumber {
	if v == nil {
		return &TLJsonNumber{
			Data2: &JSONValue{
				Cmd:       Cmd_JsonNumber,
				ClassName: ClassJsonNumber,
			},
		}
	}
	v1 := &TLJsonNumber{Data2: v}
	v1.Data2.ClassName = ClassJsonNumber
	return v1
}

func NewTLJsonString(v *JSONValue) *TLJsonString {
	if v == nil {
		return &TLJsonString{
			Data2: &JSONValue{
				Cmd:       Cmd_JsonString,
				ClassName: ClassJsonString,
			},
		}
	}
	v1 := &TLJsonString{Data2: v}
	v1.Data2.ClassName = ClassJsonString
	return v1
}

func NewTLJsonArray(v *JSONValue) *TLJsonArray {
	if v == nil {
		return &TLJsonArray{
			Data2: &JSONValue{
				Cmd:       Cmd_JsonArray,
				ClassName: ClassJsonArray,
			},
		}
	}
	v1 := &TLJsonArray{Data2: v}
	v1.Data2.ClassName = ClassJsonArray
	return v1
}

func NewTLJsonObject(v *JSONValue) *TLJsonObject {
	if v == nil {
		return &TLJsonObject{
			Data2: &JSONValue{
				Cmd:       Cmd_JsonObject,
				ClassName: ClassJsonObject,
			},
		}
	}
	v1 := &TLJsonObject{Data2: v}
	v1.Data2.ClassName = ClassJsonObject
	return v1
}

func NewTLKeyboardButton(v *KeyboardButton) *TLKeyboardButton {
	if v == nil {
		return &TLKeyboardButton{
			Data2: &KeyboardButton{
				Cmd:       Cmd_KeyboardButton,
				ClassName: ClassKeyboardButton,
			},
		}
	}
	v1 := &TLKeyboardButton{Data2: v}
	v1.Data2.ClassName = ClassKeyboardButton
	return v1
}

func NewTLKeyboardButtonUrl(v *KeyboardButton) *TLKeyboardButtonUrl {
	if v == nil {
		return &TLKeyboardButtonUrl{
			Data2: &KeyboardButton{
				Cmd:       Cmd_KeyboardButtonUrl,
				ClassName: ClassKeyboardButtonUrl,
			},
		}
	}
	v1 := &TLKeyboardButtonUrl{Data2: v}
	v1.Data2.ClassName = ClassKeyboardButtonUrl
	return v1
}

func NewTLKeyboardButtonCallback(v *KeyboardButton) *TLKeyboardButtonCallback {
	if v == nil {
		return &TLKeyboardButtonCallback{
			Data2: &KeyboardButton{
				Cmd:       Cmd_KeyboardButtonCallback,
				ClassName: ClassKeyboardButtonCallback,
			},
		}
	}
	v1 := &TLKeyboardButtonCallback{Data2: v}
	v1.Data2.ClassName = ClassKeyboardButtonCallback
	return v1
}

func NewTLKeyboardButtonRequestPhone(v *KeyboardButton) *TLKeyboardButtonRequestPhone {
	if v == nil {
		return &TLKeyboardButtonRequestPhone{
			Data2: &KeyboardButton{
				Cmd:       Cmd_KeyboardButtonRequestPhone,
				ClassName: ClassKeyboardButtonRequestPhone,
			},
		}
	}
	v1 := &TLKeyboardButtonRequestPhone{Data2: v}
	v1.Data2.ClassName = ClassKeyboardButtonRequestPhone
	return v1
}

func NewTLKeyboardButtonRequestGeoLocation(v *KeyboardButton) *TLKeyboardButtonRequestGeoLocation {
	if v == nil {
		return &TLKeyboardButtonRequestGeoLocation{
			Data2: &KeyboardButton{
				Cmd:       Cmd_KeyboardButtonRequestGeoLocation,
				ClassName: ClassKeyboardButtonRequestGeoLocation,
			},
		}
	}
	v1 := &TLKeyboardButtonRequestGeoLocation{Data2: v}
	v1.Data2.ClassName = ClassKeyboardButtonRequestGeoLocation
	return v1
}

func NewTLKeyboardButtonSwitchInline(v *KeyboardButton) *TLKeyboardButtonSwitchInline {
	if v == nil {
		return &TLKeyboardButtonSwitchInline{
			Data2: &KeyboardButton{
				Cmd:       Cmd_KeyboardButtonSwitchInline,
				ClassName: ClassKeyboardButtonSwitchInline,
			},
		}
	}
	v1 := &TLKeyboardButtonSwitchInline{Data2: v}
	v1.Data2.ClassName = ClassKeyboardButtonSwitchInline
	return v1
}

func NewTLKeyboardButtonGame(v *KeyboardButton) *TLKeyboardButtonGame {
	if v == nil {
		return &TLKeyboardButtonGame{
			Data2: &KeyboardButton{
				Cmd:       Cmd_KeyboardButtonGame,
				ClassName: ClassKeyboardButtonGame,
			},
		}
	}
	v1 := &TLKeyboardButtonGame{Data2: v}
	v1.Data2.ClassName = ClassKeyboardButtonGame
	return v1
}

func NewTLKeyboardButtonBuy(v *KeyboardButton) *TLKeyboardButtonBuy {
	if v == nil {
		return &TLKeyboardButtonBuy{
			Data2: &KeyboardButton{
				Cmd:       Cmd_KeyboardButtonBuy,
				ClassName: ClassKeyboardButtonBuy,
			},
		}
	}
	v1 := &TLKeyboardButtonBuy{Data2: v}
	v1.Data2.ClassName = ClassKeyboardButtonBuy
	return v1
}

func NewTLKeyboardButtonUrlAuth(v *KeyboardButton) *TLKeyboardButtonUrlAuth {
	if v == nil {
		return &TLKeyboardButtonUrlAuth{
			Data2: &KeyboardButton{
				Cmd:       Cmd_KeyboardButtonUrlAuth,
				ClassName: ClassKeyboardButtonUrlAuth,
			},
		}
	}
	v1 := &TLKeyboardButtonUrlAuth{Data2: v}
	v1.Data2.ClassName = ClassKeyboardButtonUrlAuth
	return v1
}

func NewTLInputKeyboardButtonUrlAuth(v *KeyboardButton) *TLInputKeyboardButtonUrlAuth {
	if v == nil {
		return &TLInputKeyboardButtonUrlAuth{
			Data2: &KeyboardButton{
				Cmd:       Cmd_InputKeyboardButtonUrlAuth,
				ClassName: ClassInputKeyboardButtonUrlAuth,
			},
		}
	}
	v1 := &TLInputKeyboardButtonUrlAuth{Data2: v}
	v1.Data2.ClassName = ClassInputKeyboardButtonUrlAuth
	return v1
}

func NewTLKeyboardButtonRequestPoll(v *KeyboardButton) *TLKeyboardButtonRequestPoll {
	if v == nil {
		return &TLKeyboardButtonRequestPoll{
			Data2: &KeyboardButton{
				Cmd:       Cmd_KeyboardButtonRequestPoll,
				ClassName: ClassKeyboardButtonRequestPoll,
			},
		}
	}
	v1 := &TLKeyboardButtonRequestPoll{Data2: v}
	v1.Data2.ClassName = ClassKeyboardButtonRequestPoll
	return v1
}

func NewTLKeyboardButtonRow(v *KeyboardButtonRow) *TLKeyboardButtonRow {
	if v == nil {
		return &TLKeyboardButtonRow{
			Data2: &KeyboardButtonRow{
				Cmd:       Cmd_KeyboardButtonRow,
				ClassName: ClassKeyboardButtonRow,
			},
		}
	}
	v1 := &TLKeyboardButtonRow{Data2: v}
	v1.Data2.ClassName = ClassKeyboardButtonRow
	return v1
}

func NewTLLabeledPrice(v *LabeledPrice) *TLLabeledPrice {
	if v == nil {
		return &TLLabeledPrice{
			Data2: &LabeledPrice{
				Cmd:       Cmd_LabeledPrice,
				ClassName: ClassLabeledPrice,
			},
		}
	}
	v1 := &TLLabeledPrice{Data2: v}
	v1.Data2.ClassName = ClassLabeledPrice
	return v1
}

func NewTLLangPackDifference(v *LangPackDifference) *TLLangPackDifference {
	if v == nil {
		return &TLLangPackDifference{
			Data2: &LangPackDifference{
				Cmd:       Cmd_LangPackDifference,
				ClassName: ClassLangPackDifference,
			},
		}
	}
	v1 := &TLLangPackDifference{Data2: v}
	v1.Data2.ClassName = ClassLangPackDifference
	return v1
}

func NewTLLangPackLanguage(v *LangPackLanguage) *TLLangPackLanguage {
	if v == nil {
		return &TLLangPackLanguage{
			Data2: &LangPackLanguage{
				Cmd:       Cmd_LangPackLanguage,
				ClassName: ClassLangPackLanguage,
			},
		}
	}
	v1 := &TLLangPackLanguage{Data2: v}
	v1.Data2.ClassName = ClassLangPackLanguage
	return v1
}

func NewTLLangPackString(v *LangPackString) *TLLangPackString {
	if v == nil {
		return &TLLangPackString{
			Data2: &LangPackString{
				Cmd:       Cmd_LangPackString,
				ClassName: ClassLangPackString,
			},
		}
	}
	v1 := &TLLangPackString{Data2: v}
	v1.Data2.ClassName = ClassLangPackString
	return v1
}

func NewTLLangPackStringPluralized(v *LangPackString) *TLLangPackStringPluralized {
	if v == nil {
		return &TLLangPackStringPluralized{
			Data2: &LangPackString{
				Cmd:       Cmd_LangPackStringPluralized,
				ClassName: ClassLangPackStringPluralized,
			},
		}
	}
	v1 := &TLLangPackStringPluralized{Data2: v}
	v1.Data2.ClassName = ClassLangPackStringPluralized
	return v1
}

func NewTLLangPackStringDeleted(v *LangPackString) *TLLangPackStringDeleted {
	if v == nil {
		return &TLLangPackStringDeleted{
			Data2: &LangPackString{
				Cmd:       Cmd_LangPackStringDeleted,
				ClassName: ClassLangPackStringDeleted,
			},
		}
	}
	v1 := &TLLangPackStringDeleted{Data2: v}
	v1.Data2.ClassName = ClassLangPackStringDeleted
	return v1
}

func NewTLMaskCoords(v *MaskCoords) *TLMaskCoords {
	if v == nil {
		return &TLMaskCoords{
			Data2: &MaskCoords{
				Cmd:       Cmd_MaskCoords,
				ClassName: ClassMaskCoords,
			},
		}
	}
	v1 := &TLMaskCoords{Data2: v}
	v1.Data2.ClassName = ClassMaskCoords
	return v1
}

func NewTLMessageEmpty(v *Message) *TLMessageEmpty {
	if v == nil {
		return &TLMessageEmpty{
			Data2: &Message{
				Cmd:       Cmd_MessageEmpty,
				ClassName: ClassMessageEmpty,
			},
		}
	}
	v1 := &TLMessageEmpty{Data2: v}
	v1.Data2.ClassName = ClassMessageEmpty
	return v1
}

func NewTLMessage(v *Message) *TLMessage {
	if v == nil {
		return &TLMessage{
			Data2: &Message{
				Cmd:       Cmd_Message,
				ClassName: ClassMessage,
			},
		}
	}
	v1 := &TLMessage{Data2: v}
	v1.Data2.ClassName = ClassMessage
	return v1
}

func NewTLMessageService(v *Message) *TLMessageService {
	if v == nil {
		return &TLMessageService{
			Data2: &Message{
				Cmd:       Cmd_MessageService,
				ClassName: ClassMessageService,
			},
		}
	}
	v1 := &TLMessageService{Data2: v}
	v1.Data2.ClassName = ClassMessageService
	return v1
}

func NewTLMessageActionEmpty(v *MessageAction) *TLMessageActionEmpty {
	if v == nil {
		return &TLMessageActionEmpty{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionEmpty,
				ClassName: ClassMessageActionEmpty,
			},
		}
	}
	v1 := &TLMessageActionEmpty{Data2: v}
	v1.Data2.ClassName = ClassMessageActionEmpty
	return v1
}

func NewTLMessageActionChatCreate(v *MessageAction) *TLMessageActionChatCreate {
	if v == nil {
		return &TLMessageActionChatCreate{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionChatCreate,
				ClassName: ClassMessageActionChatCreate,
			},
		}
	}
	v1 := &TLMessageActionChatCreate{Data2: v}
	v1.Data2.ClassName = ClassMessageActionChatCreate
	return v1
}

func NewTLMessageActionChatEditTitle(v *MessageAction) *TLMessageActionChatEditTitle {
	if v == nil {
		return &TLMessageActionChatEditTitle{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionChatEditTitle,
				ClassName: ClassMessageActionChatEditTitle,
			},
		}
	}
	v1 := &TLMessageActionChatEditTitle{Data2: v}
	v1.Data2.ClassName = ClassMessageActionChatEditTitle
	return v1
}

func NewTLMessageActionChatEditPhoto(v *MessageAction) *TLMessageActionChatEditPhoto {
	if v == nil {
		return &TLMessageActionChatEditPhoto{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionChatEditPhoto,
				ClassName: ClassMessageActionChatEditPhoto,
			},
		}
	}
	v1 := &TLMessageActionChatEditPhoto{Data2: v}
	v1.Data2.ClassName = ClassMessageActionChatEditPhoto
	return v1
}

func NewTLMessageActionChatDeletePhoto(v *MessageAction) *TLMessageActionChatDeletePhoto {
	if v == nil {
		return &TLMessageActionChatDeletePhoto{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionChatDeletePhoto,
				ClassName: ClassMessageActionChatDeletePhoto,
			},
		}
	}
	v1 := &TLMessageActionChatDeletePhoto{Data2: v}
	v1.Data2.ClassName = ClassMessageActionChatDeletePhoto
	return v1
}

func NewTLMessageActionChatAddUser(v *MessageAction) *TLMessageActionChatAddUser {
	if v == nil {
		return &TLMessageActionChatAddUser{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionChatAddUser,
				ClassName: ClassMessageActionChatAddUser,
			},
		}
	}
	v1 := &TLMessageActionChatAddUser{Data2: v}
	v1.Data2.ClassName = ClassMessageActionChatAddUser
	return v1
}

func NewTLMessageActionChatDeleteUser(v *MessageAction) *TLMessageActionChatDeleteUser {
	if v == nil {
		return &TLMessageActionChatDeleteUser{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionChatDeleteUser,
				ClassName: ClassMessageActionChatDeleteUser,
			},
		}
	}
	v1 := &TLMessageActionChatDeleteUser{Data2: v}
	v1.Data2.ClassName = ClassMessageActionChatDeleteUser
	return v1
}

func NewTLMessageActionChatJoinedByLink(v *MessageAction) *TLMessageActionChatJoinedByLink {
	if v == nil {
		return &TLMessageActionChatJoinedByLink{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionChatJoinedByLink,
				ClassName: ClassMessageActionChatJoinedByLink,
			},
		}
	}
	v1 := &TLMessageActionChatJoinedByLink{Data2: v}
	v1.Data2.ClassName = ClassMessageActionChatJoinedByLink
	return v1
}

func NewTLMessageActionChannelCreate(v *MessageAction) *TLMessageActionChannelCreate {
	if v == nil {
		return &TLMessageActionChannelCreate{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionChannelCreate,
				ClassName: ClassMessageActionChannelCreate,
			},
		}
	}
	v1 := &TLMessageActionChannelCreate{Data2: v}
	v1.Data2.ClassName = ClassMessageActionChannelCreate
	return v1
}

func NewTLMessageActionChatMigrateTo(v *MessageAction) *TLMessageActionChatMigrateTo {
	if v == nil {
		return &TLMessageActionChatMigrateTo{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionChatMigrateTo,
				ClassName: ClassMessageActionChatMigrateTo,
			},
		}
	}
	v1 := &TLMessageActionChatMigrateTo{Data2: v}
	v1.Data2.ClassName = ClassMessageActionChatMigrateTo
	return v1
}

func NewTLMessageActionChannelMigrateFrom(v *MessageAction) *TLMessageActionChannelMigrateFrom {
	if v == nil {
		return &TLMessageActionChannelMigrateFrom{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionChannelMigrateFrom,
				ClassName: ClassMessageActionChannelMigrateFrom,
			},
		}
	}
	v1 := &TLMessageActionChannelMigrateFrom{Data2: v}
	v1.Data2.ClassName = ClassMessageActionChannelMigrateFrom
	return v1
}

func NewTLMessageActionPinMessage(v *MessageAction) *TLMessageActionPinMessage {
	if v == nil {
		return &TLMessageActionPinMessage{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionPinMessage,
				ClassName: ClassMessageActionPinMessage,
			},
		}
	}
	v1 := &TLMessageActionPinMessage{Data2: v}
	v1.Data2.ClassName = ClassMessageActionPinMessage
	return v1
}

func NewTLMessageActionHistoryClear(v *MessageAction) *TLMessageActionHistoryClear {
	if v == nil {
		return &TLMessageActionHistoryClear{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionHistoryClear,
				ClassName: ClassMessageActionHistoryClear,
			},
		}
	}
	v1 := &TLMessageActionHistoryClear{Data2: v}
	v1.Data2.ClassName = ClassMessageActionHistoryClear
	return v1
}

func NewTLMessageActionGameScore(v *MessageAction) *TLMessageActionGameScore {
	if v == nil {
		return &TLMessageActionGameScore{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionGameScore,
				ClassName: ClassMessageActionGameScore,
			},
		}
	}
	v1 := &TLMessageActionGameScore{Data2: v}
	v1.Data2.ClassName = ClassMessageActionGameScore
	return v1
}

func NewTLMessageActionPaymentSentMe(v *MessageAction) *TLMessageActionPaymentSentMe {
	if v == nil {
		return &TLMessageActionPaymentSentMe{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionPaymentSentMe,
				ClassName: ClassMessageActionPaymentSentMe,
			},
		}
	}
	v1 := &TLMessageActionPaymentSentMe{Data2: v}
	v1.Data2.ClassName = ClassMessageActionPaymentSentMe
	return v1
}

func NewTLMessageActionPaymentSent(v *MessageAction) *TLMessageActionPaymentSent {
	if v == nil {
		return &TLMessageActionPaymentSent{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionPaymentSent,
				ClassName: ClassMessageActionPaymentSent,
			},
		}
	}
	v1 := &TLMessageActionPaymentSent{Data2: v}
	v1.Data2.ClassName = ClassMessageActionPaymentSent
	return v1
}

func NewTLMessageActionScreenshotTaken(v *MessageAction) *TLMessageActionScreenshotTaken {
	if v == nil {
		return &TLMessageActionScreenshotTaken{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionScreenshotTaken,
				ClassName: ClassMessageActionScreenshotTaken,
			},
		}
	}
	v1 := &TLMessageActionScreenshotTaken{Data2: v}
	v1.Data2.ClassName = ClassMessageActionScreenshotTaken
	return v1
}

func NewTLMessageActionCustomAction(v *MessageAction) *TLMessageActionCustomAction {
	if v == nil {
		return &TLMessageActionCustomAction{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionCustomAction,
				ClassName: ClassMessageActionCustomAction,
			},
		}
	}
	v1 := &TLMessageActionCustomAction{Data2: v}
	v1.Data2.ClassName = ClassMessageActionCustomAction
	return v1
}

func NewTLMessageActionBotAllowed(v *MessageAction) *TLMessageActionBotAllowed {
	if v == nil {
		return &TLMessageActionBotAllowed{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionBotAllowed,
				ClassName: ClassMessageActionBotAllowed,
			},
		}
	}
	v1 := &TLMessageActionBotAllowed{Data2: v}
	v1.Data2.ClassName = ClassMessageActionBotAllowed
	return v1
}

func NewTLMessageActionSecureValuesSentMe(v *MessageAction) *TLMessageActionSecureValuesSentMe {
	if v == nil {
		return &TLMessageActionSecureValuesSentMe{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionSecureValuesSentMe,
				ClassName: ClassMessageActionSecureValuesSentMe,
			},
		}
	}
	v1 := &TLMessageActionSecureValuesSentMe{Data2: v}
	v1.Data2.ClassName = ClassMessageActionSecureValuesSentMe
	return v1
}

func NewTLMessageActionSecureValuesSent(v *MessageAction) *TLMessageActionSecureValuesSent {
	if v == nil {
		return &TLMessageActionSecureValuesSent{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionSecureValuesSent,
				ClassName: ClassMessageActionSecureValuesSent,
			},
		}
	}
	v1 := &TLMessageActionSecureValuesSent{Data2: v}
	v1.Data2.ClassName = ClassMessageActionSecureValuesSent
	return v1
}

func NewTLMessageActionContactSignUp(v *MessageAction) *TLMessageActionContactSignUp {
	if v == nil {
		return &TLMessageActionContactSignUp{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionContactSignUp,
				ClassName: ClassMessageActionContactSignUp,
			},
		}
	}
	v1 := &TLMessageActionContactSignUp{Data2: v}
	v1.Data2.ClassName = ClassMessageActionContactSignUp
	return v1
}

func NewTLMessageActionPhoneCall(v *MessageAction) *TLMessageActionPhoneCall {
	if v == nil {
		return &TLMessageActionPhoneCall{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionPhoneCall,
				ClassName: ClassMessageActionPhoneCall,
			},
		}
	}
	v1 := &TLMessageActionPhoneCall{Data2: v}
	v1.Data2.ClassName = ClassMessageActionPhoneCall
	return v1
}

func NewTLMessageActionGeoProximityReached(v *MessageAction) *TLMessageActionGeoProximityReached {
	if v == nil {
		return &TLMessageActionGeoProximityReached{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionGeoProximityReached,
				ClassName: ClassMessageActionGeoProximityReached,
			},
		}
	}
	v1 := &TLMessageActionGeoProximityReached{Data2: v}
	v1.Data2.ClassName = ClassMessageActionGeoProximityReached
	return v1
}

func NewTLMessageActionGroupCall(v *MessageAction) *TLMessageActionGroupCall {
	if v == nil {
		return &TLMessageActionGroupCall{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionGroupCall,
				ClassName: ClassMessageActionGroupCall,
			},
		}
	}
	v1 := &TLMessageActionGroupCall{Data2: v}
	v1.Data2.ClassName = ClassMessageActionGroupCall
	return v1
}

func NewTLMessageActionInviteToGroupCall(v *MessageAction) *TLMessageActionInviteToGroupCall {
	if v == nil {
		return &TLMessageActionInviteToGroupCall{
			Data2: &MessageAction{
				Cmd:       Cmd_MessageActionInviteToGroupCall,
				ClassName: ClassMessageActionInviteToGroupCall,
			},
		}
	}
	v1 := &TLMessageActionInviteToGroupCall{Data2: v}
	v1.Data2.ClassName = ClassMessageActionInviteToGroupCall
	return v1
}

func NewTLMessageEntityUnknown(v *MessageEntity) *TLMessageEntityUnknown {
	if v == nil {
		return &TLMessageEntityUnknown{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityUnknown,
				ClassName: ClassMessageEntityUnknown,
			},
		}
	}
	v1 := &TLMessageEntityUnknown{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityUnknown
	return v1
}

func NewTLMessageEntityMention(v *MessageEntity) *TLMessageEntityMention {
	if v == nil {
		return &TLMessageEntityMention{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityMention,
				ClassName: ClassMessageEntityMention,
			},
		}
	}
	v1 := &TLMessageEntityMention{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityMention
	return v1
}

func NewTLMessageEntityHashtag(v *MessageEntity) *TLMessageEntityHashtag {
	if v == nil {
		return &TLMessageEntityHashtag{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityHashtag,
				ClassName: ClassMessageEntityHashtag,
			},
		}
	}
	v1 := &TLMessageEntityHashtag{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityHashtag
	return v1
}

func NewTLMessageEntityBotCommand(v *MessageEntity) *TLMessageEntityBotCommand {
	if v == nil {
		return &TLMessageEntityBotCommand{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityBotCommand,
				ClassName: ClassMessageEntityBotCommand,
			},
		}
	}
	v1 := &TLMessageEntityBotCommand{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityBotCommand
	return v1
}

func NewTLMessageEntityUrl(v *MessageEntity) *TLMessageEntityUrl {
	if v == nil {
		return &TLMessageEntityUrl{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityUrl,
				ClassName: ClassMessageEntityUrl,
			},
		}
	}
	v1 := &TLMessageEntityUrl{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityUrl
	return v1
}

func NewTLMessageEntityEmail(v *MessageEntity) *TLMessageEntityEmail {
	if v == nil {
		return &TLMessageEntityEmail{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityEmail,
				ClassName: ClassMessageEntityEmail,
			},
		}
	}
	v1 := &TLMessageEntityEmail{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityEmail
	return v1
}

func NewTLMessageEntityBold(v *MessageEntity) *TLMessageEntityBold {
	if v == nil {
		return &TLMessageEntityBold{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityBold,
				ClassName: ClassMessageEntityBold,
			},
		}
	}
	v1 := &TLMessageEntityBold{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityBold
	return v1
}

func NewTLMessageEntityItalic(v *MessageEntity) *TLMessageEntityItalic {
	if v == nil {
		return &TLMessageEntityItalic{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityItalic,
				ClassName: ClassMessageEntityItalic,
			},
		}
	}
	v1 := &TLMessageEntityItalic{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityItalic
	return v1
}

func NewTLMessageEntityCode(v *MessageEntity) *TLMessageEntityCode {
	if v == nil {
		return &TLMessageEntityCode{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityCode,
				ClassName: ClassMessageEntityCode,
			},
		}
	}
	v1 := &TLMessageEntityCode{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityCode
	return v1
}

func NewTLMessageEntityPre(v *MessageEntity) *TLMessageEntityPre {
	if v == nil {
		return &TLMessageEntityPre{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityPre,
				ClassName: ClassMessageEntityPre,
			},
		}
	}
	v1 := &TLMessageEntityPre{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityPre
	return v1
}

func NewTLMessageEntityTextUrl(v *MessageEntity) *TLMessageEntityTextUrl {
	if v == nil {
		return &TLMessageEntityTextUrl{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityTextUrl,
				ClassName: ClassMessageEntityTextUrl,
			},
		}
	}
	v1 := &TLMessageEntityTextUrl{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityTextUrl
	return v1
}

func NewTLMessageEntityMentionName(v *MessageEntity) *TLMessageEntityMentionName {
	if v == nil {
		return &TLMessageEntityMentionName{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityMentionName,
				ClassName: ClassMessageEntityMentionName,
			},
		}
	}
	v1 := &TLMessageEntityMentionName{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityMentionName
	return v1
}

func NewTLInputMessageEntityMentionName(v *MessageEntity) *TLInputMessageEntityMentionName {
	if v == nil {
		return &TLInputMessageEntityMentionName{
			Data2: &MessageEntity{
				Cmd:       Cmd_InputMessageEntityMentionName,
				ClassName: ClassInputMessageEntityMentionName,
			},
		}
	}
	v1 := &TLInputMessageEntityMentionName{Data2: v}
	v1.Data2.ClassName = ClassInputMessageEntityMentionName
	return v1
}

func NewTLMessageEntityPhone(v *MessageEntity) *TLMessageEntityPhone {
	if v == nil {
		return &TLMessageEntityPhone{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityPhone,
				ClassName: ClassMessageEntityPhone,
			},
		}
	}
	v1 := &TLMessageEntityPhone{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityPhone
	return v1
}

func NewTLMessageEntityCashtag(v *MessageEntity) *TLMessageEntityCashtag {
	if v == nil {
		return &TLMessageEntityCashtag{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityCashtag,
				ClassName: ClassMessageEntityCashtag,
			},
		}
	}
	v1 := &TLMessageEntityCashtag{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityCashtag
	return v1
}

func NewTLMessageEntityUnderline(v *MessageEntity) *TLMessageEntityUnderline {
	if v == nil {
		return &TLMessageEntityUnderline{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityUnderline,
				ClassName: ClassMessageEntityUnderline,
			},
		}
	}
	v1 := &TLMessageEntityUnderline{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityUnderline
	return v1
}

func NewTLMessageEntityStrike(v *MessageEntity) *TLMessageEntityStrike {
	if v == nil {
		return &TLMessageEntityStrike{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityStrike,
				ClassName: ClassMessageEntityStrike,
			},
		}
	}
	v1 := &TLMessageEntityStrike{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityStrike
	return v1
}

func NewTLMessageEntityBlockquote(v *MessageEntity) *TLMessageEntityBlockquote {
	if v == nil {
		return &TLMessageEntityBlockquote{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityBlockquote,
				ClassName: ClassMessageEntityBlockquote,
			},
		}
	}
	v1 := &TLMessageEntityBlockquote{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityBlockquote
	return v1
}

func NewTLMessageEntityBankCard(v *MessageEntity) *TLMessageEntityBankCard {
	if v == nil {
		return &TLMessageEntityBankCard{
			Data2: &MessageEntity{
				Cmd:       Cmd_MessageEntityBankCard,
				ClassName: ClassMessageEntityBankCard,
			},
		}
	}
	v1 := &TLMessageEntityBankCard{Data2: v}
	v1.Data2.ClassName = ClassMessageEntityBankCard
	return v1
}

func NewTLMessageFwdHeader(v *MessageFwdHeader) *TLMessageFwdHeader {
	if v == nil {
		return &TLMessageFwdHeader{
			Data2: &MessageFwdHeader{
				Cmd:       Cmd_MessageFwdHeader,
				ClassName: ClassMessageFwdHeader,
			},
		}
	}
	v1 := &TLMessageFwdHeader{Data2: v}
	v1.Data2.ClassName = ClassMessageFwdHeader
	return v1
}

func NewTLMessageGroup(v *MessageGroup) *TLMessageGroup {
	if v == nil {
		return &TLMessageGroup{
			Data2: &MessageGroup{
				Cmd:       Cmd_MessageGroup,
				ClassName: ClassMessageGroup,
			},
		}
	}
	v1 := &TLMessageGroup{Data2: v}
	v1.Data2.ClassName = ClassMessageGroup
	return v1
}

func NewTLMessageInteractionCounters(v *MessageInteractionCounters) *TLMessageInteractionCounters {
	if v == nil {
		return &TLMessageInteractionCounters{
			Data2: &MessageInteractionCounters{
				Cmd:       Cmd_MessageInteractionCounters,
				ClassName: ClassMessageInteractionCounters,
			},
		}
	}
	v1 := &TLMessageInteractionCounters{Data2: v}
	v1.Data2.ClassName = ClassMessageInteractionCounters
	return v1
}

func NewTLMessageMediaEmpty(v *MessageMedia) *TLMessageMediaEmpty {
	if v == nil {
		return &TLMessageMediaEmpty{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaEmpty,
				ClassName: ClassMessageMediaEmpty,
			},
		}
	}
	v1 := &TLMessageMediaEmpty{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaEmpty
	return v1
}

func NewTLMessageMediaPhoto(v *MessageMedia) *TLMessageMediaPhoto {
	if v == nil {
		return &TLMessageMediaPhoto{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaPhoto,
				ClassName: ClassMessageMediaPhoto,
			},
		}
	}
	v1 := &TLMessageMediaPhoto{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaPhoto
	return v1
}

func NewTLMessageMediaGeo(v *MessageMedia) *TLMessageMediaGeo {
	if v == nil {
		return &TLMessageMediaGeo{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaGeo,
				ClassName: ClassMessageMediaGeo,
			},
		}
	}
	v1 := &TLMessageMediaGeo{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaGeo
	return v1
}

func NewTLMessageMediaContact(v *MessageMedia) *TLMessageMediaContact {
	if v == nil {
		return &TLMessageMediaContact{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaContact,
				ClassName: ClassMessageMediaContact,
			},
		}
	}
	v1 := &TLMessageMediaContact{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaContact
	return v1
}

func NewTLMessageMediaUnsupported(v *MessageMedia) *TLMessageMediaUnsupported {
	if v == nil {
		return &TLMessageMediaUnsupported{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaUnsupported,
				ClassName: ClassMessageMediaUnsupported,
			},
		}
	}
	v1 := &TLMessageMediaUnsupported{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaUnsupported
	return v1
}

func NewTLMessageMediaDocument(v *MessageMedia) *TLMessageMediaDocument {
	if v == nil {
		return &TLMessageMediaDocument{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaDocument,
				ClassName: ClassMessageMediaDocument,
			},
		}
	}
	v1 := &TLMessageMediaDocument{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaDocument
	return v1
}

func NewTLMessageMediaWebPage(v *MessageMedia) *TLMessageMediaWebPage {
	if v == nil {
		return &TLMessageMediaWebPage{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaWebPage,
				ClassName: ClassMessageMediaWebPage,
			},
		}
	}
	v1 := &TLMessageMediaWebPage{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaWebPage
	return v1
}

func NewTLMessageMediaVenue(v *MessageMedia) *TLMessageMediaVenue {
	if v == nil {
		return &TLMessageMediaVenue{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaVenue,
				ClassName: ClassMessageMediaVenue,
			},
		}
	}
	v1 := &TLMessageMediaVenue{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaVenue
	return v1
}

func NewTLMessageMediaGame(v *MessageMedia) *TLMessageMediaGame {
	if v == nil {
		return &TLMessageMediaGame{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaGame,
				ClassName: ClassMessageMediaGame,
			},
		}
	}
	v1 := &TLMessageMediaGame{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaGame
	return v1
}

func NewTLMessageMediaInvoice(v *MessageMedia) *TLMessageMediaInvoice {
	if v == nil {
		return &TLMessageMediaInvoice{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaInvoice,
				ClassName: ClassMessageMediaInvoice,
			},
		}
	}
	v1 := &TLMessageMediaInvoice{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaInvoice
	return v1
}

func NewTLMessageMediaGeoLive(v *MessageMedia) *TLMessageMediaGeoLive {
	if v == nil {
		return &TLMessageMediaGeoLive{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaGeoLive,
				ClassName: ClassMessageMediaGeoLive,
			},
		}
	}
	v1 := &TLMessageMediaGeoLive{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaGeoLive
	return v1
}

func NewTLMessageMediaPoll(v *MessageMedia) *TLMessageMediaPoll {
	if v == nil {
		return &TLMessageMediaPoll{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaPoll,
				ClassName: ClassMessageMediaPoll,
			},
		}
	}
	v1 := &TLMessageMediaPoll{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaPoll
	return v1
}

func NewTLMessageMediaDice(v *MessageMedia) *TLMessageMediaDice {
	if v == nil {
		return &TLMessageMediaDice{
			Data2: &MessageMedia{
				Cmd:       Cmd_MessageMediaDice,
				ClassName: ClassMessageMediaDice,
			},
		}
	}
	v1 := &TLMessageMediaDice{Data2: v}
	v1.Data2.ClassName = ClassMessageMediaDice
	return v1
}

func NewTLMessageRange(v *MessageRange) *TLMessageRange {
	if v == nil {
		return &TLMessageRange{
			Data2: &MessageRange{
				Cmd:       Cmd_MessageRange,
				ClassName: ClassMessageRange,
			},
		}
	}
	v1 := &TLMessageRange{Data2: v}
	v1.Data2.ClassName = ClassMessageRange
	return v1
}

func NewTLMessageReplies(v *MessageReplies) *TLMessageReplies {
	if v == nil {
		return &TLMessageReplies{
			Data2: &MessageReplies{
				Cmd:       Cmd_MessageReplies,
				ClassName: ClassMessageReplies,
			},
		}
	}
	v1 := &TLMessageReplies{Data2: v}
	v1.Data2.ClassName = ClassMessageReplies
	return v1
}

func NewTLMessageReplyHeader(v *MessageReplyHeader) *TLMessageReplyHeader {
	if v == nil {
		return &TLMessageReplyHeader{
			Data2: &MessageReplyHeader{
				Cmd:       Cmd_MessageReplyHeader,
				ClassName: ClassMessageReplyHeader,
			},
		}
	}
	v1 := &TLMessageReplyHeader{Data2: v}
	v1.Data2.ClassName = ClassMessageReplyHeader
	return v1
}

func NewTLMessageUserVote(v *MessageUserVote) *TLMessageUserVote {
	if v == nil {
		return &TLMessageUserVote{
			Data2: &MessageUserVote{
				Cmd:       Cmd_MessageUserVote,
				ClassName: ClassMessageUserVote,
			},
		}
	}
	v1 := &TLMessageUserVote{Data2: v}
	v1.Data2.ClassName = ClassMessageUserVote
	return v1
}

func NewTLMessageUserVoteInputOption(v *MessageUserVote) *TLMessageUserVoteInputOption {
	if v == nil {
		return &TLMessageUserVoteInputOption{
			Data2: &MessageUserVote{
				Cmd:       Cmd_MessageUserVoteInputOption,
				ClassName: ClassMessageUserVoteInputOption,
			},
		}
	}
	v1 := &TLMessageUserVoteInputOption{Data2: v}
	v1.Data2.ClassName = ClassMessageUserVoteInputOption
	return v1
}

func NewTLMessageUserVoteMultiple(v *MessageUserVote) *TLMessageUserVoteMultiple {
	if v == nil {
		return &TLMessageUserVoteMultiple{
			Data2: &MessageUserVote{
				Cmd:       Cmd_MessageUserVoteMultiple,
				ClassName: ClassMessageUserVoteMultiple,
			},
		}
	}
	v1 := &TLMessageUserVoteMultiple{Data2: v}
	v1.Data2.ClassName = ClassMessageUserVoteMultiple
	return v1
}

func NewTLMessageViews(v *MessageViews) *TLMessageViews {
	if v == nil {
		return &TLMessageViews{
			Data2: &MessageViews{
				Cmd:       Cmd_MessageViews,
				ClassName: ClassMessageViews,
			},
		}
	}
	v1 := &TLMessageViews{Data2: v}
	v1.Data2.ClassName = ClassMessageViews
	return v1
}

func NewTLInputMessagesFilterEmpty(v *MessagesFilter) *TLInputMessagesFilterEmpty {
	if v == nil {
		return &TLInputMessagesFilterEmpty{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterEmpty,
				ClassName: ClassInputMessagesFilterEmpty,
			},
		}
	}
	v1 := &TLInputMessagesFilterEmpty{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterEmpty
	return v1
}

func NewTLInputMessagesFilterPhotos(v *MessagesFilter) *TLInputMessagesFilterPhotos {
	if v == nil {
		return &TLInputMessagesFilterPhotos{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterPhotos,
				ClassName: ClassInputMessagesFilterPhotos,
			},
		}
	}
	v1 := &TLInputMessagesFilterPhotos{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterPhotos
	return v1
}

func NewTLInputMessagesFilterVideo(v *MessagesFilter) *TLInputMessagesFilterVideo {
	if v == nil {
		return &TLInputMessagesFilterVideo{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterVideo,
				ClassName: ClassInputMessagesFilterVideo,
			},
		}
	}
	v1 := &TLInputMessagesFilterVideo{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterVideo
	return v1
}

func NewTLInputMessagesFilterPhotoVideo(v *MessagesFilter) *TLInputMessagesFilterPhotoVideo {
	if v == nil {
		return &TLInputMessagesFilterPhotoVideo{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterPhotoVideo,
				ClassName: ClassInputMessagesFilterPhotoVideo,
			},
		}
	}
	v1 := &TLInputMessagesFilterPhotoVideo{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterPhotoVideo
	return v1
}

func NewTLInputMessagesFilterPhotoVideoDocuments(v *MessagesFilter) *TLInputMessagesFilterPhotoVideoDocuments {
	if v == nil {
		return &TLInputMessagesFilterPhotoVideoDocuments{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterPhotoVideoDocuments,
				ClassName: ClassInputMessagesFilterPhotoVideoDocuments,
			},
		}
	}
	v1 := &TLInputMessagesFilterPhotoVideoDocuments{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterPhotoVideoDocuments
	return v1
}

func NewTLInputMessagesFilterDocument(v *MessagesFilter) *TLInputMessagesFilterDocument {
	if v == nil {
		return &TLInputMessagesFilterDocument{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterDocument,
				ClassName: ClassInputMessagesFilterDocument,
			},
		}
	}
	v1 := &TLInputMessagesFilterDocument{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterDocument
	return v1
}

func NewTLInputMessagesFilterUrl(v *MessagesFilter) *TLInputMessagesFilterUrl {
	if v == nil {
		return &TLInputMessagesFilterUrl{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterUrl,
				ClassName: ClassInputMessagesFilterUrl,
			},
		}
	}
	v1 := &TLInputMessagesFilterUrl{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterUrl
	return v1
}

func NewTLInputMessagesFilterGif(v *MessagesFilter) *TLInputMessagesFilterGif {
	if v == nil {
		return &TLInputMessagesFilterGif{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterGif,
				ClassName: ClassInputMessagesFilterGif,
			},
		}
	}
	v1 := &TLInputMessagesFilterGif{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterGif
	return v1
}

func NewTLInputMessagesFilterVoice(v *MessagesFilter) *TLInputMessagesFilterVoice {
	if v == nil {
		return &TLInputMessagesFilterVoice{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterVoice,
				ClassName: ClassInputMessagesFilterVoice,
			},
		}
	}
	v1 := &TLInputMessagesFilterVoice{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterVoice
	return v1
}

func NewTLInputMessagesFilterMusic(v *MessagesFilter) *TLInputMessagesFilterMusic {
	if v == nil {
		return &TLInputMessagesFilterMusic{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterMusic,
				ClassName: ClassInputMessagesFilterMusic,
			},
		}
	}
	v1 := &TLInputMessagesFilterMusic{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterMusic
	return v1
}

func NewTLInputMessagesFilterChatPhotos(v *MessagesFilter) *TLInputMessagesFilterChatPhotos {
	if v == nil {
		return &TLInputMessagesFilterChatPhotos{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterChatPhotos,
				ClassName: ClassInputMessagesFilterChatPhotos,
			},
		}
	}
	v1 := &TLInputMessagesFilterChatPhotos{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterChatPhotos
	return v1
}

func NewTLInputMessagesFilterPhoneCalls(v *MessagesFilter) *TLInputMessagesFilterPhoneCalls {
	if v == nil {
		return &TLInputMessagesFilterPhoneCalls{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterPhoneCalls,
				ClassName: ClassInputMessagesFilterPhoneCalls,
			},
		}
	}
	v1 := &TLInputMessagesFilterPhoneCalls{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterPhoneCalls
	return v1
}

func NewTLInputMessagesFilterRoundVoice(v *MessagesFilter) *TLInputMessagesFilterRoundVoice {
	if v == nil {
		return &TLInputMessagesFilterRoundVoice{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterRoundVoice,
				ClassName: ClassInputMessagesFilterRoundVoice,
			},
		}
	}
	v1 := &TLInputMessagesFilterRoundVoice{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterRoundVoice
	return v1
}

func NewTLInputMessagesFilterRoundVideo(v *MessagesFilter) *TLInputMessagesFilterRoundVideo {
	if v == nil {
		return &TLInputMessagesFilterRoundVideo{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterRoundVideo,
				ClassName: ClassInputMessagesFilterRoundVideo,
			},
		}
	}
	v1 := &TLInputMessagesFilterRoundVideo{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterRoundVideo
	return v1
}

func NewTLInputMessagesFilterMyMentions(v *MessagesFilter) *TLInputMessagesFilterMyMentions {
	if v == nil {
		return &TLInputMessagesFilterMyMentions{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterMyMentions,
				ClassName: ClassInputMessagesFilterMyMentions,
			},
		}
	}
	v1 := &TLInputMessagesFilterMyMentions{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterMyMentions
	return v1
}

func NewTLInputMessagesFilterGeo(v *MessagesFilter) *TLInputMessagesFilterGeo {
	if v == nil {
		return &TLInputMessagesFilterGeo{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterGeo,
				ClassName: ClassInputMessagesFilterGeo,
			},
		}
	}
	v1 := &TLInputMessagesFilterGeo{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterGeo
	return v1
}

func NewTLInputMessagesFilterContacts(v *MessagesFilter) *TLInputMessagesFilterContacts {
	if v == nil {
		return &TLInputMessagesFilterContacts{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterContacts,
				ClassName: ClassInputMessagesFilterContacts,
			},
		}
	}
	v1 := &TLInputMessagesFilterContacts{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterContacts
	return v1
}

func NewTLInputMessagesFilterPinned(v *MessagesFilter) *TLInputMessagesFilterPinned {
	if v == nil {
		return &TLInputMessagesFilterPinned{
			Data2: &MessagesFilter{
				Cmd:       Cmd_InputMessagesFilterPinned,
				ClassName: ClassInputMessagesFilterPinned,
			},
		}
	}
	v1 := &TLInputMessagesFilterPinned{Data2: v}
	v1.Data2.ClassName = ClassInputMessagesFilterPinned
	return v1
}

func NewTLMessagesAffectedHistory(v *Messages_AffectedHistory) *TLMessagesAffectedHistory {
	if v == nil {
		return &TLMessagesAffectedHistory{
			Data2: &Messages_AffectedHistory{
				Cmd:       Cmd_MessagesAffectedHistory,
				ClassName: ClassMessagesAffectedHistory,
			},
		}
	}
	v1 := &TLMessagesAffectedHistory{Data2: v}
	v1.Data2.ClassName = ClassMessagesAffectedHistory
	return v1
}

func NewTLMessagesAffectedMessages(v *Messages_AffectedMessages) *TLMessagesAffectedMessages {
	if v == nil {
		return &TLMessagesAffectedMessages{
			Data2: &Messages_AffectedMessages{
				Cmd:       Cmd_MessagesAffectedMessages,
				ClassName: ClassMessagesAffectedMessages,
			},
		}
	}
	v1 := &TLMessagesAffectedMessages{Data2: v}
	v1.Data2.ClassName = ClassMessagesAffectedMessages
	return v1
}

func NewTLMessagesAllStickersNotModified(v *Messages_AllStickers) *TLMessagesAllStickersNotModified {
	if v == nil {
		return &TLMessagesAllStickersNotModified{
			Data2: &Messages_AllStickers{
				Cmd:       Cmd_MessagesAllStickersNotModified,
				ClassName: ClassMessagesAllStickersNotModified,
			},
		}
	}
	v1 := &TLMessagesAllStickersNotModified{Data2: v}
	v1.Data2.ClassName = ClassMessagesAllStickersNotModified
	return v1
}

func NewTLMessagesAllStickers(v *Messages_AllStickers) *TLMessagesAllStickers {
	if v == nil {
		return &TLMessagesAllStickers{
			Data2: &Messages_AllStickers{
				Cmd:       Cmd_MessagesAllStickers,
				ClassName: ClassMessagesAllStickers,
			},
		}
	}
	v1 := &TLMessagesAllStickers{Data2: v}
	v1.Data2.ClassName = ClassMessagesAllStickers
	return v1
}

func NewTLMessagesArchivedStickers(v *Messages_ArchivedStickers) *TLMessagesArchivedStickers {
	if v == nil {
		return &TLMessagesArchivedStickers{
			Data2: &Messages_ArchivedStickers{
				Cmd:       Cmd_MessagesArchivedStickers,
				ClassName: ClassMessagesArchivedStickers,
			},
		}
	}
	v1 := &TLMessagesArchivedStickers{Data2: v}
	v1.Data2.ClassName = ClassMessagesArchivedStickers
	return v1
}

func NewTLMessagesBotCallbackAnswer(v *Messages_BotCallbackAnswer) *TLMessagesBotCallbackAnswer {
	if v == nil {
		return &TLMessagesBotCallbackAnswer{
			Data2: &Messages_BotCallbackAnswer{
				Cmd:       Cmd_MessagesBotCallbackAnswer,
				ClassName: ClassMessagesBotCallbackAnswer,
			},
		}
	}
	v1 := &TLMessagesBotCallbackAnswer{Data2: v}
	v1.Data2.ClassName = ClassMessagesBotCallbackAnswer
	return v1
}

func NewTLMessagesBotResults(v *Messages_BotResults) *TLMessagesBotResults {
	if v == nil {
		return &TLMessagesBotResults{
			Data2: &Messages_BotResults{
				Cmd:       Cmd_MessagesBotResults,
				ClassName: ClassMessagesBotResults,
			},
		}
	}
	v1 := &TLMessagesBotResults{Data2: v}
	v1.Data2.ClassName = ClassMessagesBotResults
	return v1
}

func NewTLMessagesChatFull(v *Messages_ChatFull) *TLMessagesChatFull {
	if v == nil {
		return &TLMessagesChatFull{
			Data2: &Messages_ChatFull{
				Cmd:       Cmd_MessagesChatFull,
				ClassName: ClassMessagesChatFull,
			},
		}
	}
	v1 := &TLMessagesChatFull{Data2: v}
	v1.Data2.ClassName = ClassMessagesChatFull
	return v1
}

func NewTLMessagesChats(v *Messages_Chats) *TLMessagesChats {
	if v == nil {
		return &TLMessagesChats{
			Data2: &Messages_Chats{
				Cmd:       Cmd_MessagesChats,
				ClassName: ClassMessagesChats,
			},
		}
	}
	v1 := &TLMessagesChats{Data2: v}
	v1.Data2.ClassName = ClassMessagesChats
	return v1
}

func NewTLMessagesChatsSlice(v *Messages_Chats) *TLMessagesChatsSlice {
	if v == nil {
		return &TLMessagesChatsSlice{
			Data2: &Messages_Chats{
				Cmd:       Cmd_MessagesChatsSlice,
				ClassName: ClassMessagesChatsSlice,
			},
		}
	}
	v1 := &TLMessagesChatsSlice{Data2: v}
	v1.Data2.ClassName = ClassMessagesChatsSlice
	return v1
}

func NewTLMessagesDhConfigNotModified(v *Messages_DhConfig) *TLMessagesDhConfigNotModified {
	if v == nil {
		return &TLMessagesDhConfigNotModified{
			Data2: &Messages_DhConfig{
				Cmd:       Cmd_MessagesDhConfigNotModified,
				ClassName: ClassMessagesDhConfigNotModified,
			},
		}
	}
	v1 := &TLMessagesDhConfigNotModified{Data2: v}
	v1.Data2.ClassName = ClassMessagesDhConfigNotModified
	return v1
}

func NewTLMessagesDhConfig(v *Messages_DhConfig) *TLMessagesDhConfig {
	if v == nil {
		return &TLMessagesDhConfig{
			Data2: &Messages_DhConfig{
				Cmd:       Cmd_MessagesDhConfig,
				ClassName: ClassMessagesDhConfig,
			},
		}
	}
	v1 := &TLMessagesDhConfig{Data2: v}
	v1.Data2.ClassName = ClassMessagesDhConfig
	return v1
}

func NewTLMessagesDialogs(v *Messages_Dialogs) *TLMessagesDialogs {
	if v == nil {
		return &TLMessagesDialogs{
			Data2: &Messages_Dialogs{
				Cmd:       Cmd_MessagesDialogs,
				ClassName: ClassMessagesDialogs,
			},
		}
	}
	v1 := &TLMessagesDialogs{Data2: v}
	v1.Data2.ClassName = ClassMessagesDialogs
	return v1
}

func NewTLMessagesDialogsSlice(v *Messages_Dialogs) *TLMessagesDialogsSlice {
	if v == nil {
		return &TLMessagesDialogsSlice{
			Data2: &Messages_Dialogs{
				Cmd:       Cmd_MessagesDialogsSlice,
				ClassName: ClassMessagesDialogsSlice,
			},
		}
	}
	v1 := &TLMessagesDialogsSlice{Data2: v}
	v1.Data2.ClassName = ClassMessagesDialogsSlice
	return v1
}

func NewTLMessagesDialogsNotModified(v *Messages_Dialogs) *TLMessagesDialogsNotModified {
	if v == nil {
		return &TLMessagesDialogsNotModified{
			Data2: &Messages_Dialogs{
				Cmd:       Cmd_MessagesDialogsNotModified,
				ClassName: ClassMessagesDialogsNotModified,
			},
		}
	}
	v1 := &TLMessagesDialogsNotModified{Data2: v}
	v1.Data2.ClassName = ClassMessagesDialogsNotModified
	return v1
}

func NewTLMessagesDiscussionMessage(v *Messages_DiscussionMessage) *TLMessagesDiscussionMessage {
	if v == nil {
		return &TLMessagesDiscussionMessage{
			Data2: &Messages_DiscussionMessage{
				Cmd:       Cmd_MessagesDiscussionMessage,
				ClassName: ClassMessagesDiscussionMessage,
			},
		}
	}
	v1 := &TLMessagesDiscussionMessage{Data2: v}
	v1.Data2.ClassName = ClassMessagesDiscussionMessage
	return v1
}

func NewTLMessagesFavedStickersNotModified(v *Messages_FavedStickers) *TLMessagesFavedStickersNotModified {
	if v == nil {
		return &TLMessagesFavedStickersNotModified{
			Data2: &Messages_FavedStickers{
				Cmd:       Cmd_MessagesFavedStickersNotModified,
				ClassName: ClassMessagesFavedStickersNotModified,
			},
		}
	}
	v1 := &TLMessagesFavedStickersNotModified{Data2: v}
	v1.Data2.ClassName = ClassMessagesFavedStickersNotModified
	return v1
}

func NewTLMessagesFavedStickers(v *Messages_FavedStickers) *TLMessagesFavedStickers {
	if v == nil {
		return &TLMessagesFavedStickers{
			Data2: &Messages_FavedStickers{
				Cmd:       Cmd_MessagesFavedStickers,
				ClassName: ClassMessagesFavedStickers,
			},
		}
	}
	v1 := &TLMessagesFavedStickers{Data2: v}
	v1.Data2.ClassName = ClassMessagesFavedStickers
	return v1
}

func NewTLMessagesFeaturedStickersNotModified(v *Messages_FeaturedStickers) *TLMessagesFeaturedStickersNotModified {
	if v == nil {
		return &TLMessagesFeaturedStickersNotModified{
			Data2: &Messages_FeaturedStickers{
				Cmd:       Cmd_MessagesFeaturedStickersNotModified,
				ClassName: ClassMessagesFeaturedStickersNotModified,
			},
		}
	}
	v1 := &TLMessagesFeaturedStickersNotModified{Data2: v}
	v1.Data2.ClassName = ClassMessagesFeaturedStickersNotModified
	return v1
}

func NewTLMessagesFeaturedStickers(v *Messages_FeaturedStickers) *TLMessagesFeaturedStickers {
	if v == nil {
		return &TLMessagesFeaturedStickers{
			Data2: &Messages_FeaturedStickers{
				Cmd:       Cmd_MessagesFeaturedStickers,
				ClassName: ClassMessagesFeaturedStickers,
			},
		}
	}
	v1 := &TLMessagesFeaturedStickers{Data2: v}
	v1.Data2.ClassName = ClassMessagesFeaturedStickers
	return v1
}

func NewTLMessagesFoundGifs(v *Messages_FoundGifs) *TLMessagesFoundGifs {
	if v == nil {
		return &TLMessagesFoundGifs{
			Data2: &Messages_FoundGifs{
				Cmd:       Cmd_MessagesFoundGifs,
				ClassName: ClassMessagesFoundGifs,
			},
		}
	}
	v1 := &TLMessagesFoundGifs{Data2: v}
	v1.Data2.ClassName = ClassMessagesFoundGifs
	return v1
}

func NewTLMessagesFoundStickerSetsNotModified(v *Messages_FoundStickerSets) *TLMessagesFoundStickerSetsNotModified {
	if v == nil {
		return &TLMessagesFoundStickerSetsNotModified{
			Data2: &Messages_FoundStickerSets{
				Cmd:       Cmd_MessagesFoundStickerSetsNotModified,
				ClassName: ClassMessagesFoundStickerSetsNotModified,
			},
		}
	}
	v1 := &TLMessagesFoundStickerSetsNotModified{Data2: v}
	v1.Data2.ClassName = ClassMessagesFoundStickerSetsNotModified
	return v1
}

func NewTLMessagesFoundStickerSets(v *Messages_FoundStickerSets) *TLMessagesFoundStickerSets {
	if v == nil {
		return &TLMessagesFoundStickerSets{
			Data2: &Messages_FoundStickerSets{
				Cmd:       Cmd_MessagesFoundStickerSets,
				ClassName: ClassMessagesFoundStickerSets,
			},
		}
	}
	v1 := &TLMessagesFoundStickerSets{Data2: v}
	v1.Data2.ClassName = ClassMessagesFoundStickerSets
	return v1
}

func NewTLMessagesHighScores(v *Messages_HighScores) *TLMessagesHighScores {
	if v == nil {
		return &TLMessagesHighScores{
			Data2: &Messages_HighScores{
				Cmd:       Cmd_MessagesHighScores,
				ClassName: ClassMessagesHighScores,
			},
		}
	}
	v1 := &TLMessagesHighScores{Data2: v}
	v1.Data2.ClassName = ClassMessagesHighScores
	return v1
}

func NewTLMessagesInactiveChats(v *Messages_InactiveChats) *TLMessagesInactiveChats {
	if v == nil {
		return &TLMessagesInactiveChats{
			Data2: &Messages_InactiveChats{
				Cmd:       Cmd_MessagesInactiveChats,
				ClassName: ClassMessagesInactiveChats,
			},
		}
	}
	v1 := &TLMessagesInactiveChats{Data2: v}
	v1.Data2.ClassName = ClassMessagesInactiveChats
	return v1
}

func NewTLMessagesMessageEditData(v *Messages_MessageEditData) *TLMessagesMessageEditData {
	if v == nil {
		return &TLMessagesMessageEditData{
			Data2: &Messages_MessageEditData{
				Cmd:       Cmd_MessagesMessageEditData,
				ClassName: ClassMessagesMessageEditData,
			},
		}
	}
	v1 := &TLMessagesMessageEditData{Data2: v}
	v1.Data2.ClassName = ClassMessagesMessageEditData
	return v1
}

func NewTLMessagesMessageViews(v *Messages_MessageViews) *TLMessagesMessageViews {
	if v == nil {
		return &TLMessagesMessageViews{
			Data2: &Messages_MessageViews{
				Cmd:       Cmd_MessagesMessageViews,
				ClassName: ClassMessagesMessageViews,
			},
		}
	}
	v1 := &TLMessagesMessageViews{Data2: v}
	v1.Data2.ClassName = ClassMessagesMessageViews
	return v1
}

func NewTLMessagesMessages(v *Messages_Messages) *TLMessagesMessages {
	if v == nil {
		return &TLMessagesMessages{
			Data2: &Messages_Messages{
				Cmd:       Cmd_MessagesMessages,
				ClassName: ClassMessagesMessages,
			},
		}
	}
	v1 := &TLMessagesMessages{Data2: v}
	v1.Data2.ClassName = ClassMessagesMessages
	return v1
}

func NewTLMessagesMessagesSlice(v *Messages_Messages) *TLMessagesMessagesSlice {
	if v == nil {
		return &TLMessagesMessagesSlice{
			Data2: &Messages_Messages{
				Cmd:       Cmd_MessagesMessagesSlice,
				ClassName: ClassMessagesMessagesSlice,
			},
		}
	}
	v1 := &TLMessagesMessagesSlice{Data2: v}
	v1.Data2.ClassName = ClassMessagesMessagesSlice
	return v1
}

func NewTLMessagesChannelMessages(v *Messages_Messages) *TLMessagesChannelMessages {
	if v == nil {
		return &TLMessagesChannelMessages{
			Data2: &Messages_Messages{
				Cmd:       Cmd_MessagesChannelMessages,
				ClassName: ClassMessagesChannelMessages,
			},
		}
	}
	v1 := &TLMessagesChannelMessages{Data2: v}
	v1.Data2.ClassName = ClassMessagesChannelMessages
	return v1
}

func NewTLMessagesMessagesNotModified(v *Messages_Messages) *TLMessagesMessagesNotModified {
	if v == nil {
		return &TLMessagesMessagesNotModified{
			Data2: &Messages_Messages{
				Cmd:       Cmd_MessagesMessagesNotModified,
				ClassName: ClassMessagesMessagesNotModified,
			},
		}
	}
	v1 := &TLMessagesMessagesNotModified{Data2: v}
	v1.Data2.ClassName = ClassMessagesMessagesNotModified
	return v1
}

func NewTLMessagesPeerDialogs(v *Messages_PeerDialogs) *TLMessagesPeerDialogs {
	if v == nil {
		return &TLMessagesPeerDialogs{
			Data2: &Messages_PeerDialogs{
				Cmd:       Cmd_MessagesPeerDialogs,
				ClassName: ClassMessagesPeerDialogs,
			},
		}
	}
	v1 := &TLMessagesPeerDialogs{Data2: v}
	v1.Data2.ClassName = ClassMessagesPeerDialogs
	return v1
}

func NewTLMessagesRecentStickersNotModified(v *Messages_RecentStickers) *TLMessagesRecentStickersNotModified {
	if v == nil {
		return &TLMessagesRecentStickersNotModified{
			Data2: &Messages_RecentStickers{
				Cmd:       Cmd_MessagesRecentStickersNotModified,
				ClassName: ClassMessagesRecentStickersNotModified,
			},
		}
	}
	v1 := &TLMessagesRecentStickersNotModified{Data2: v}
	v1.Data2.ClassName = ClassMessagesRecentStickersNotModified
	return v1
}

func NewTLMessagesRecentStickers(v *Messages_RecentStickers) *TLMessagesRecentStickers {
	if v == nil {
		return &TLMessagesRecentStickers{
			Data2: &Messages_RecentStickers{
				Cmd:       Cmd_MessagesRecentStickers,
				ClassName: ClassMessagesRecentStickers,
			},
		}
	}
	v1 := &TLMessagesRecentStickers{Data2: v}
	v1.Data2.ClassName = ClassMessagesRecentStickers
	return v1
}

func NewTLMessagesSavedGifsNotModified(v *Messages_SavedGifs) *TLMessagesSavedGifsNotModified {
	if v == nil {
		return &TLMessagesSavedGifsNotModified{
			Data2: &Messages_SavedGifs{
				Cmd:       Cmd_MessagesSavedGifsNotModified,
				ClassName: ClassMessagesSavedGifsNotModified,
			},
		}
	}
	v1 := &TLMessagesSavedGifsNotModified{Data2: v}
	v1.Data2.ClassName = ClassMessagesSavedGifsNotModified
	return v1
}

func NewTLMessagesSavedGifs(v *Messages_SavedGifs) *TLMessagesSavedGifs {
	if v == nil {
		return &TLMessagesSavedGifs{
			Data2: &Messages_SavedGifs{
				Cmd:       Cmd_MessagesSavedGifs,
				ClassName: ClassMessagesSavedGifs,
			},
		}
	}
	v1 := &TLMessagesSavedGifs{Data2: v}
	v1.Data2.ClassName = ClassMessagesSavedGifs
	return v1
}

func NewTLMessagesSearchCounter(v *Messages_SearchCounter) *TLMessagesSearchCounter {
	if v == nil {
		return &TLMessagesSearchCounter{
			Data2: &Messages_SearchCounter{
				Cmd:       Cmd_MessagesSearchCounter,
				ClassName: ClassMessagesSearchCounter,
			},
		}
	}
	v1 := &TLMessagesSearchCounter{Data2: v}
	v1.Data2.ClassName = ClassMessagesSearchCounter
	return v1
}

func NewTLMessagesSentEncryptedMessage(v *Messages_SentEncryptedMessage) *TLMessagesSentEncryptedMessage {
	if v == nil {
		return &TLMessagesSentEncryptedMessage{
			Data2: &Messages_SentEncryptedMessage{
				Cmd:       Cmd_MessagesSentEncryptedMessage,
				ClassName: ClassMessagesSentEncryptedMessage,
			},
		}
	}
	v1 := &TLMessagesSentEncryptedMessage{Data2: v}
	v1.Data2.ClassName = ClassMessagesSentEncryptedMessage
	return v1
}

func NewTLMessagesSentEncryptedFile(v *Messages_SentEncryptedMessage) *TLMessagesSentEncryptedFile {
	if v == nil {
		return &TLMessagesSentEncryptedFile{
			Data2: &Messages_SentEncryptedMessage{
				Cmd:       Cmd_MessagesSentEncryptedFile,
				ClassName: ClassMessagesSentEncryptedFile,
			},
		}
	}
	v1 := &TLMessagesSentEncryptedFile{Data2: v}
	v1.Data2.ClassName = ClassMessagesSentEncryptedFile
	return v1
}

func NewTLMessagesStickerSet(v *Messages_StickerSet) *TLMessagesStickerSet {
	if v == nil {
		return &TLMessagesStickerSet{
			Data2: &Messages_StickerSet{
				Cmd:       Cmd_MessagesStickerSet,
				ClassName: ClassMessagesStickerSet,
			},
		}
	}
	v1 := &TLMessagesStickerSet{Data2: v}
	v1.Data2.ClassName = ClassMessagesStickerSet
	return v1
}

func NewTLMessagesStickerSetInstallResultSuccess(v *Messages_StickerSetInstallResult) *TLMessagesStickerSetInstallResultSuccess {
	if v == nil {
		return &TLMessagesStickerSetInstallResultSuccess{
			Data2: &Messages_StickerSetInstallResult{
				Cmd:       Cmd_MessagesStickerSetInstallResultSuccess,
				ClassName: ClassMessagesStickerSetInstallResultSuccess,
			},
		}
	}
	v1 := &TLMessagesStickerSetInstallResultSuccess{Data2: v}
	v1.Data2.ClassName = ClassMessagesStickerSetInstallResultSuccess
	return v1
}

func NewTLMessagesStickerSetInstallResultArchive(v *Messages_StickerSetInstallResult) *TLMessagesStickerSetInstallResultArchive {
	if v == nil {
		return &TLMessagesStickerSetInstallResultArchive{
			Data2: &Messages_StickerSetInstallResult{
				Cmd:       Cmd_MessagesStickerSetInstallResultArchive,
				ClassName: ClassMessagesStickerSetInstallResultArchive,
			},
		}
	}
	v1 := &TLMessagesStickerSetInstallResultArchive{Data2: v}
	v1.Data2.ClassName = ClassMessagesStickerSetInstallResultArchive
	return v1
}

func NewTLMessagesStickersNotModified(v *Messages_Stickers) *TLMessagesStickersNotModified {
	if v == nil {
		return &TLMessagesStickersNotModified{
			Data2: &Messages_Stickers{
				Cmd:       Cmd_MessagesStickersNotModified,
				ClassName: ClassMessagesStickersNotModified,
			},
		}
	}
	v1 := &TLMessagesStickersNotModified{Data2: v}
	v1.Data2.ClassName = ClassMessagesStickersNotModified
	return v1
}

func NewTLMessagesStickers(v *Messages_Stickers) *TLMessagesStickers {
	if v == nil {
		return &TLMessagesStickers{
			Data2: &Messages_Stickers{
				Cmd:       Cmd_MessagesStickers,
				ClassName: ClassMessagesStickers,
			},
		}
	}
	v1 := &TLMessagesStickers{Data2: v}
	v1.Data2.ClassName = ClassMessagesStickers
	return v1
}

func NewTLMessagesVotesList(v *Messages_VotesList) *TLMessagesVotesList {
	if v == nil {
		return &TLMessagesVotesList{
			Data2: &Messages_VotesList{
				Cmd:       Cmd_MessagesVotesList,
				ClassName: ClassMessagesVotesList,
			},
		}
	}
	v1 := &TLMessagesVotesList{Data2: v}
	v1.Data2.ClassName = ClassMessagesVotesList
	return v1
}

func NewTLMsgDetailedInfo(v *MsgDetailedInfo) *TLMsgDetailedInfo {
	if v == nil {
		return &TLMsgDetailedInfo{
			Data2: &MsgDetailedInfo{
				Cmd:       Cmd_MsgDetailedInfo,
				ClassName: ClassMsgDetailedInfo,
			},
		}
	}
	v1 := &TLMsgDetailedInfo{Data2: v}
	v1.Data2.ClassName = ClassMsgDetailedInfo
	return v1
}

func NewTLMsgNewDetailedInfo(v *MsgDetailedInfo) *TLMsgNewDetailedInfo {
	if v == nil {
		return &TLMsgNewDetailedInfo{
			Data2: &MsgDetailedInfo{
				Cmd:       Cmd_MsgNewDetailedInfo,
				ClassName: ClassMsgNewDetailedInfo,
			},
		}
	}
	v1 := &TLMsgNewDetailedInfo{Data2: v}
	v1.Data2.ClassName = ClassMsgNewDetailedInfo
	return v1
}

func NewTLMsgResendReq(v *MsgResendReq) *TLMsgResendReq {
	if v == nil {
		return &TLMsgResendReq{
			Data2: &MsgResendReq{
				Cmd:       Cmd_MsgResendReq,
				ClassName: ClassMsgResendReq,
			},
		}
	}
	v1 := &TLMsgResendReq{Data2: v}
	v1.Data2.ClassName = ClassMsgResendReq
	return v1
}

func NewTLMsgsAck(v *MsgsAck) *TLMsgsAck {
	if v == nil {
		return &TLMsgsAck{
			Data2: &MsgsAck{
				Cmd:       Cmd_MsgsAck,
				ClassName: ClassMsgsAck,
			},
		}
	}
	v1 := &TLMsgsAck{Data2: v}
	v1.Data2.ClassName = ClassMsgsAck
	return v1
}

func NewTLMsgsAllInfo(v *MsgsAllInfo) *TLMsgsAllInfo {
	if v == nil {
		return &TLMsgsAllInfo{
			Data2: &MsgsAllInfo{
				Cmd:       Cmd_MsgsAllInfo,
				ClassName: ClassMsgsAllInfo,
			},
		}
	}
	v1 := &TLMsgsAllInfo{Data2: v}
	v1.Data2.ClassName = ClassMsgsAllInfo
	return v1
}

func NewTLMsgsStateInfo(v *MsgsStateInfo) *TLMsgsStateInfo {
	if v == nil {
		return &TLMsgsStateInfo{
			Data2: &MsgsStateInfo{
				Cmd:       Cmd_MsgsStateInfo,
				ClassName: ClassMsgsStateInfo,
			},
		}
	}
	v1 := &TLMsgsStateInfo{Data2: v}
	v1.Data2.ClassName = ClassMsgsStateInfo
	return v1
}

func NewTLMsgsStateReq(v *MsgsStateReq) *TLMsgsStateReq {
	if v == nil {
		return &TLMsgsStateReq{
			Data2: &MsgsStateReq{
				Cmd:       Cmd_MsgsStateReq,
				ClassName: ClassMsgsStateReq,
			},
		}
	}
	v1 := &TLMsgsStateReq{Data2: v}
	v1.Data2.ClassName = ClassMsgsStateReq
	return v1
}

func NewTLNearestDc(v *NearestDc) *TLNearestDc {
	if v == nil {
		return &TLNearestDc{
			Data2: &NearestDc{
				Cmd:       Cmd_NearestDc,
				ClassName: ClassNearestDc,
			},
		}
	}
	v1 := &TLNearestDc{Data2: v}
	v1.Data2.ClassName = ClassNearestDc
	return v1
}

func NewTLNewSessionCreated(v *NewSession) *TLNewSessionCreated {
	if v == nil {
		return &TLNewSessionCreated{
			Data2: &NewSession{
				Cmd:       Cmd_NewSessionCreated,
				ClassName: ClassNewSessionCreated,
			},
		}
	}
	v1 := &TLNewSessionCreated{Data2: v}
	v1.Data2.ClassName = ClassNewSessionCreated
	return v1
}

func NewTLNotifyPeer(v *NotifyPeer) *TLNotifyPeer {
	if v == nil {
		return &TLNotifyPeer{
			Data2: &NotifyPeer{
				Cmd:       Cmd_NotifyPeer,
				ClassName: ClassNotifyPeer,
			},
		}
	}
	v1 := &TLNotifyPeer{Data2: v}
	v1.Data2.ClassName = ClassNotifyPeer
	return v1
}

func NewTLNotifyUsers(v *NotifyPeer) *TLNotifyUsers {
	if v == nil {
		return &TLNotifyUsers{
			Data2: &NotifyPeer{
				Cmd:       Cmd_NotifyUsers,
				ClassName: ClassNotifyUsers,
			},
		}
	}
	v1 := &TLNotifyUsers{Data2: v}
	v1.Data2.ClassName = ClassNotifyUsers
	return v1
}

func NewTLNotifyChats(v *NotifyPeer) *TLNotifyChats {
	if v == nil {
		return &TLNotifyChats{
			Data2: &NotifyPeer{
				Cmd:       Cmd_NotifyChats,
				ClassName: ClassNotifyChats,
			},
		}
	}
	v1 := &TLNotifyChats{Data2: v}
	v1.Data2.ClassName = ClassNotifyChats
	return v1
}

func NewTLNotifyAll(v *NotifyPeer) *TLNotifyAll {
	if v == nil {
		return &TLNotifyAll{
			Data2: &NotifyPeer{
				Cmd:       Cmd_NotifyAll,
				ClassName: ClassNotifyAll,
			},
		}
	}
	v1 := &TLNotifyAll{Data2: v}
	v1.Data2.ClassName = ClassNotifyAll
	return v1
}

func NewTLNotifyBroadcasts(v *NotifyPeer) *TLNotifyBroadcasts {
	if v == nil {
		return &TLNotifyBroadcasts{
			Data2: &NotifyPeer{
				Cmd:       Cmd_NotifyBroadcasts,
				ClassName: ClassNotifyBroadcasts,
			},
		}
	}
	v1 := &TLNotifyBroadcasts{Data2: v}
	v1.Data2.ClassName = ClassNotifyBroadcasts
	return v1
}

func NewTLNull(v *Null) *TLNull {
	if v == nil {
		return &TLNull{
			Data2: &Null{
				Cmd:       Cmd_Null,
				ClassName: ClassNull,
			},
		}
	}
	v1 := &TLNull{Data2: v}
	v1.Data2.ClassName = ClassNull
	return v1
}

func NewTLPQInnerData(v *P_QInnerData) *TLPQInnerData {
	if v == nil {
		return &TLPQInnerData{
			Data2: &P_QInnerData{
				Cmd:       Cmd_PQInnerData,
				ClassName: ClassPQInnerData,
			},
		}
	}
	v1 := &TLPQInnerData{Data2: v}
	v1.Data2.ClassName = ClassPQInnerData
	return v1
}

func NewTLPQInnerDataDc(v *P_QInnerData) *TLPQInnerDataDc {
	if v == nil {
		return &TLPQInnerDataDc{
			Data2: &P_QInnerData{
				Cmd:       Cmd_PQInnerDataDc,
				ClassName: ClassPQInnerDataDc,
			},
		}
	}
	v1 := &TLPQInnerDataDc{Data2: v}
	v1.Data2.ClassName = ClassPQInnerDataDc
	return v1
}

func NewTLPQInnerDataTemp(v *P_QInnerData) *TLPQInnerDataTemp {
	if v == nil {
		return &TLPQInnerDataTemp{
			Data2: &P_QInnerData{
				Cmd:       Cmd_PQInnerDataTemp,
				ClassName: ClassPQInnerDataTemp,
			},
		}
	}
	v1 := &TLPQInnerDataTemp{Data2: v}
	v1.Data2.ClassName = ClassPQInnerDataTemp
	return v1
}

func NewTLPQInnerDataTempDc(v *P_QInnerData) *TLPQInnerDataTempDc {
	if v == nil {
		return &TLPQInnerDataTempDc{
			Data2: &P_QInnerData{
				Cmd:       Cmd_PQInnerDataTempDc,
				ClassName: ClassPQInnerDataTempDc,
			},
		}
	}
	v1 := &TLPQInnerDataTempDc{Data2: v}
	v1.Data2.ClassName = ClassPQInnerDataTempDc
	return v1
}

func NewTLPagePart(v *Page) *TLPagePart {
	if v == nil {
		return &TLPagePart{
			Data2: &Page{
				Cmd:       Cmd_PagePart,
				ClassName: ClassPagePart,
			},
		}
	}
	v1 := &TLPagePart{Data2: v}
	v1.Data2.ClassName = ClassPagePart
	return v1
}

func NewTLPageFull(v *Page) *TLPageFull {
	if v == nil {
		return &TLPageFull{
			Data2: &Page{
				Cmd:       Cmd_PageFull,
				ClassName: ClassPageFull,
			},
		}
	}
	v1 := &TLPageFull{Data2: v}
	v1.Data2.ClassName = ClassPageFull
	return v1
}

func NewTLPage(v *Page) *TLPage {
	if v == nil {
		return &TLPage{
			Data2: &Page{
				Cmd:       Cmd_Page,
				ClassName: ClassPage,
			},
		}
	}
	v1 := &TLPage{Data2: v}
	v1.Data2.ClassName = ClassPage
	return v1
}

func NewTLPageBlockUnsupported(v *PageBlock) *TLPageBlockUnsupported {
	if v == nil {
		return &TLPageBlockUnsupported{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockUnsupported,
				ClassName: ClassPageBlockUnsupported,
			},
		}
	}
	v1 := &TLPageBlockUnsupported{Data2: v}
	v1.Data2.ClassName = ClassPageBlockUnsupported
	return v1
}

func NewTLPageBlockTitle(v *PageBlock) *TLPageBlockTitle {
	if v == nil {
		return &TLPageBlockTitle{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockTitle,
				ClassName: ClassPageBlockTitle,
			},
		}
	}
	v1 := &TLPageBlockTitle{Data2: v}
	v1.Data2.ClassName = ClassPageBlockTitle
	return v1
}

func NewTLPageBlockSubtitle(v *PageBlock) *TLPageBlockSubtitle {
	if v == nil {
		return &TLPageBlockSubtitle{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockSubtitle,
				ClassName: ClassPageBlockSubtitle,
			},
		}
	}
	v1 := &TLPageBlockSubtitle{Data2: v}
	v1.Data2.ClassName = ClassPageBlockSubtitle
	return v1
}

func NewTLPageBlockAuthorDate(v *PageBlock) *TLPageBlockAuthorDate {
	if v == nil {
		return &TLPageBlockAuthorDate{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockAuthorDate,
				ClassName: ClassPageBlockAuthorDate,
			},
		}
	}
	v1 := &TLPageBlockAuthorDate{Data2: v}
	v1.Data2.ClassName = ClassPageBlockAuthorDate
	return v1
}

func NewTLPageBlockHeader(v *PageBlock) *TLPageBlockHeader {
	if v == nil {
		return &TLPageBlockHeader{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockHeader,
				ClassName: ClassPageBlockHeader,
			},
		}
	}
	v1 := &TLPageBlockHeader{Data2: v}
	v1.Data2.ClassName = ClassPageBlockHeader
	return v1
}

func NewTLPageBlockSubheader(v *PageBlock) *TLPageBlockSubheader {
	if v == nil {
		return &TLPageBlockSubheader{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockSubheader,
				ClassName: ClassPageBlockSubheader,
			},
		}
	}
	v1 := &TLPageBlockSubheader{Data2: v}
	v1.Data2.ClassName = ClassPageBlockSubheader
	return v1
}

func NewTLPageBlockParagraph(v *PageBlock) *TLPageBlockParagraph {
	if v == nil {
		return &TLPageBlockParagraph{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockParagraph,
				ClassName: ClassPageBlockParagraph,
			},
		}
	}
	v1 := &TLPageBlockParagraph{Data2: v}
	v1.Data2.ClassName = ClassPageBlockParagraph
	return v1
}

func NewTLPageBlockPreformatted(v *PageBlock) *TLPageBlockPreformatted {
	if v == nil {
		return &TLPageBlockPreformatted{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockPreformatted,
				ClassName: ClassPageBlockPreformatted,
			},
		}
	}
	v1 := &TLPageBlockPreformatted{Data2: v}
	v1.Data2.ClassName = ClassPageBlockPreformatted
	return v1
}

func NewTLPageBlockFooter(v *PageBlock) *TLPageBlockFooter {
	if v == nil {
		return &TLPageBlockFooter{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockFooter,
				ClassName: ClassPageBlockFooter,
			},
		}
	}
	v1 := &TLPageBlockFooter{Data2: v}
	v1.Data2.ClassName = ClassPageBlockFooter
	return v1
}

func NewTLPageBlockDivider(v *PageBlock) *TLPageBlockDivider {
	if v == nil {
		return &TLPageBlockDivider{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockDivider,
				ClassName: ClassPageBlockDivider,
			},
		}
	}
	v1 := &TLPageBlockDivider{Data2: v}
	v1.Data2.ClassName = ClassPageBlockDivider
	return v1
}

func NewTLPageBlockAnchor(v *PageBlock) *TLPageBlockAnchor {
	if v == nil {
		return &TLPageBlockAnchor{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockAnchor,
				ClassName: ClassPageBlockAnchor,
			},
		}
	}
	v1 := &TLPageBlockAnchor{Data2: v}
	v1.Data2.ClassName = ClassPageBlockAnchor
	return v1
}

func NewTLPageBlockList(v *PageBlock) *TLPageBlockList {
	if v == nil {
		return &TLPageBlockList{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockList,
				ClassName: ClassPageBlockList,
			},
		}
	}
	v1 := &TLPageBlockList{Data2: v}
	v1.Data2.ClassName = ClassPageBlockList
	return v1
}

func NewTLPageBlockBlockquote(v *PageBlock) *TLPageBlockBlockquote {
	if v == nil {
		return &TLPageBlockBlockquote{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockBlockquote,
				ClassName: ClassPageBlockBlockquote,
			},
		}
	}
	v1 := &TLPageBlockBlockquote{Data2: v}
	v1.Data2.ClassName = ClassPageBlockBlockquote
	return v1
}

func NewTLPageBlockPullquote(v *PageBlock) *TLPageBlockPullquote {
	if v == nil {
		return &TLPageBlockPullquote{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockPullquote,
				ClassName: ClassPageBlockPullquote,
			},
		}
	}
	v1 := &TLPageBlockPullquote{Data2: v}
	v1.Data2.ClassName = ClassPageBlockPullquote
	return v1
}

func NewTLPageBlockPhoto(v *PageBlock) *TLPageBlockPhoto {
	if v == nil {
		return &TLPageBlockPhoto{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockPhoto,
				ClassName: ClassPageBlockPhoto,
			},
		}
	}
	v1 := &TLPageBlockPhoto{Data2: v}
	v1.Data2.ClassName = ClassPageBlockPhoto
	return v1
}

func NewTLPageBlockVideo(v *PageBlock) *TLPageBlockVideo {
	if v == nil {
		return &TLPageBlockVideo{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockVideo,
				ClassName: ClassPageBlockVideo,
			},
		}
	}
	v1 := &TLPageBlockVideo{Data2: v}
	v1.Data2.ClassName = ClassPageBlockVideo
	return v1
}

func NewTLPageBlockCover(v *PageBlock) *TLPageBlockCover {
	if v == nil {
		return &TLPageBlockCover{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockCover,
				ClassName: ClassPageBlockCover,
			},
		}
	}
	v1 := &TLPageBlockCover{Data2: v}
	v1.Data2.ClassName = ClassPageBlockCover
	return v1
}

func NewTLPageBlockEmbed(v *PageBlock) *TLPageBlockEmbed {
	if v == nil {
		return &TLPageBlockEmbed{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockEmbed,
				ClassName: ClassPageBlockEmbed,
			},
		}
	}
	v1 := &TLPageBlockEmbed{Data2: v}
	v1.Data2.ClassName = ClassPageBlockEmbed
	return v1
}

func NewTLPageBlockEmbedPost(v *PageBlock) *TLPageBlockEmbedPost {
	if v == nil {
		return &TLPageBlockEmbedPost{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockEmbedPost,
				ClassName: ClassPageBlockEmbedPost,
			},
		}
	}
	v1 := &TLPageBlockEmbedPost{Data2: v}
	v1.Data2.ClassName = ClassPageBlockEmbedPost
	return v1
}

func NewTLPageBlockCollage(v *PageBlock) *TLPageBlockCollage {
	if v == nil {
		return &TLPageBlockCollage{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockCollage,
				ClassName: ClassPageBlockCollage,
			},
		}
	}
	v1 := &TLPageBlockCollage{Data2: v}
	v1.Data2.ClassName = ClassPageBlockCollage
	return v1
}

func NewTLPageBlockSlideshow(v *PageBlock) *TLPageBlockSlideshow {
	if v == nil {
		return &TLPageBlockSlideshow{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockSlideshow,
				ClassName: ClassPageBlockSlideshow,
			},
		}
	}
	v1 := &TLPageBlockSlideshow{Data2: v}
	v1.Data2.ClassName = ClassPageBlockSlideshow
	return v1
}

func NewTLPageBlockChannel(v *PageBlock) *TLPageBlockChannel {
	if v == nil {
		return &TLPageBlockChannel{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockChannel,
				ClassName: ClassPageBlockChannel,
			},
		}
	}
	v1 := &TLPageBlockChannel{Data2: v}
	v1.Data2.ClassName = ClassPageBlockChannel
	return v1
}

func NewTLPageBlockAudio(v *PageBlock) *TLPageBlockAudio {
	if v == nil {
		return &TLPageBlockAudio{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockAudio,
				ClassName: ClassPageBlockAudio,
			},
		}
	}
	v1 := &TLPageBlockAudio{Data2: v}
	v1.Data2.ClassName = ClassPageBlockAudio
	return v1
}

func NewTLPageBlockKicker(v *PageBlock) *TLPageBlockKicker {
	if v == nil {
		return &TLPageBlockKicker{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockKicker,
				ClassName: ClassPageBlockKicker,
			},
		}
	}
	v1 := &TLPageBlockKicker{Data2: v}
	v1.Data2.ClassName = ClassPageBlockKicker
	return v1
}

func NewTLPageBlockTable(v *PageBlock) *TLPageBlockTable {
	if v == nil {
		return &TLPageBlockTable{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockTable,
				ClassName: ClassPageBlockTable,
			},
		}
	}
	v1 := &TLPageBlockTable{Data2: v}
	v1.Data2.ClassName = ClassPageBlockTable
	return v1
}

func NewTLPageBlockOrderedList(v *PageBlock) *TLPageBlockOrderedList {
	if v == nil {
		return &TLPageBlockOrderedList{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockOrderedList,
				ClassName: ClassPageBlockOrderedList,
			},
		}
	}
	v1 := &TLPageBlockOrderedList{Data2: v}
	v1.Data2.ClassName = ClassPageBlockOrderedList
	return v1
}

func NewTLPageBlockDetails(v *PageBlock) *TLPageBlockDetails {
	if v == nil {
		return &TLPageBlockDetails{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockDetails,
				ClassName: ClassPageBlockDetails,
			},
		}
	}
	v1 := &TLPageBlockDetails{Data2: v}
	v1.Data2.ClassName = ClassPageBlockDetails
	return v1
}

func NewTLPageBlockRelatedArticles(v *PageBlock) *TLPageBlockRelatedArticles {
	if v == nil {
		return &TLPageBlockRelatedArticles{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockRelatedArticles,
				ClassName: ClassPageBlockRelatedArticles,
			},
		}
	}
	v1 := &TLPageBlockRelatedArticles{Data2: v}
	v1.Data2.ClassName = ClassPageBlockRelatedArticles
	return v1
}

func NewTLPageBlockMap(v *PageBlock) *TLPageBlockMap {
	if v == nil {
		return &TLPageBlockMap{
			Data2: &PageBlock{
				Cmd:       Cmd_PageBlockMap,
				ClassName: ClassPageBlockMap,
			},
		}
	}
	v1 := &TLPageBlockMap{Data2: v}
	v1.Data2.ClassName = ClassPageBlockMap
	return v1
}

func NewTLPageCaption(v *PageCaption) *TLPageCaption {
	if v == nil {
		return &TLPageCaption{
			Data2: &PageCaption{
				Cmd:       Cmd_PageCaption,
				ClassName: ClassPageCaption,
			},
		}
	}
	v1 := &TLPageCaption{Data2: v}
	v1.Data2.ClassName = ClassPageCaption
	return v1
}

func NewTLPageListItemText(v *PageListItem) *TLPageListItemText {
	if v == nil {
		return &TLPageListItemText{
			Data2: &PageListItem{
				Cmd:       Cmd_PageListItemText,
				ClassName: ClassPageListItemText,
			},
		}
	}
	v1 := &TLPageListItemText{Data2: v}
	v1.Data2.ClassName = ClassPageListItemText
	return v1
}

func NewTLPageListItemBlocks(v *PageListItem) *TLPageListItemBlocks {
	if v == nil {
		return &TLPageListItemBlocks{
			Data2: &PageListItem{
				Cmd:       Cmd_PageListItemBlocks,
				ClassName: ClassPageListItemBlocks,
			},
		}
	}
	v1 := &TLPageListItemBlocks{Data2: v}
	v1.Data2.ClassName = ClassPageListItemBlocks
	return v1
}

func NewTLPageListOrderedItemText(v *PageListOrderedItem) *TLPageListOrderedItemText {
	if v == nil {
		return &TLPageListOrderedItemText{
			Data2: &PageListOrderedItem{
				Cmd:       Cmd_PageListOrderedItemText,
				ClassName: ClassPageListOrderedItemText,
			},
		}
	}
	v1 := &TLPageListOrderedItemText{Data2: v}
	v1.Data2.ClassName = ClassPageListOrderedItemText
	return v1
}

func NewTLPageListOrderedItemBlocks(v *PageListOrderedItem) *TLPageListOrderedItemBlocks {
	if v == nil {
		return &TLPageListOrderedItemBlocks{
			Data2: &PageListOrderedItem{
				Cmd:       Cmd_PageListOrderedItemBlocks,
				ClassName: ClassPageListOrderedItemBlocks,
			},
		}
	}
	v1 := &TLPageListOrderedItemBlocks{Data2: v}
	v1.Data2.ClassName = ClassPageListOrderedItemBlocks
	return v1
}

func NewTLPageRelatedArticle(v *PageRelatedArticle) *TLPageRelatedArticle {
	if v == nil {
		return &TLPageRelatedArticle{
			Data2: &PageRelatedArticle{
				Cmd:       Cmd_PageRelatedArticle,
				ClassName: ClassPageRelatedArticle,
			},
		}
	}
	v1 := &TLPageRelatedArticle{Data2: v}
	v1.Data2.ClassName = ClassPageRelatedArticle
	return v1
}

func NewTLPageTableCell(v *PageTableCell) *TLPageTableCell {
	if v == nil {
		return &TLPageTableCell{
			Data2: &PageTableCell{
				Cmd:       Cmd_PageTableCell,
				ClassName: ClassPageTableCell,
			},
		}
	}
	v1 := &TLPageTableCell{Data2: v}
	v1.Data2.ClassName = ClassPageTableCell
	return v1
}

func NewTLPageTableRow(v *PageTableRow) *TLPageTableRow {
	if v == nil {
		return &TLPageTableRow{
			Data2: &PageTableRow{
				Cmd:       Cmd_PageTableRow,
				ClassName: ClassPageTableRow,
			},
		}
	}
	v1 := &TLPageTableRow{Data2: v}
	v1.Data2.ClassName = ClassPageTableRow
	return v1
}

func NewTLPasswordKdfAlgoUnknown(v *PasswordKdfAlgo) *TLPasswordKdfAlgoUnknown {
	if v == nil {
		return &TLPasswordKdfAlgoUnknown{
			Data2: &PasswordKdfAlgo{
				Cmd:       Cmd_PasswordKdfAlgoUnknown,
				ClassName: ClassPasswordKdfAlgoUnknown,
			},
		}
	}
	v1 := &TLPasswordKdfAlgoUnknown{Data2: v}
	v1.Data2.ClassName = ClassPasswordKdfAlgoUnknown
	return v1
}

func NewTLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow(v *PasswordKdfAlgo) *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow {
	if v == nil {
		return &TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow{
			Data2: &PasswordKdfAlgo{
				Cmd:       Cmd_PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow,
				ClassName: ClassPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow,
			},
		}
	}
	v1 := &TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow{Data2: v}
	v1.Data2.ClassName = ClassPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow
	return v1
}

func NewTLPaymentCharge(v *PaymentCharge) *TLPaymentCharge {
	if v == nil {
		return &TLPaymentCharge{
			Data2: &PaymentCharge{
				Cmd:       Cmd_PaymentCharge,
				ClassName: ClassPaymentCharge,
			},
		}
	}
	v1 := &TLPaymentCharge{Data2: v}
	v1.Data2.ClassName = ClassPaymentCharge
	return v1
}

func NewTLPaymentRequestedInfo(v *PaymentRequestedInfo) *TLPaymentRequestedInfo {
	if v == nil {
		return &TLPaymentRequestedInfo{
			Data2: &PaymentRequestedInfo{
				Cmd:       Cmd_PaymentRequestedInfo,
				ClassName: ClassPaymentRequestedInfo,
			},
		}
	}
	v1 := &TLPaymentRequestedInfo{Data2: v}
	v1.Data2.ClassName = ClassPaymentRequestedInfo
	return v1
}

func NewTLPaymentSavedCredentialsCard(v *PaymentSavedCredentials) *TLPaymentSavedCredentialsCard {
	if v == nil {
		return &TLPaymentSavedCredentialsCard{
			Data2: &PaymentSavedCredentials{
				Cmd:       Cmd_PaymentSavedCredentialsCard,
				ClassName: ClassPaymentSavedCredentialsCard,
			},
		}
	}
	v1 := &TLPaymentSavedCredentialsCard{Data2: v}
	v1.Data2.ClassName = ClassPaymentSavedCredentialsCard
	return v1
}

func NewTLPaymentsBankCardData(v *Payments_BankCardData) *TLPaymentsBankCardData {
	if v == nil {
		return &TLPaymentsBankCardData{
			Data2: &Payments_BankCardData{
				Cmd:       Cmd_PaymentsBankCardData,
				ClassName: ClassPaymentsBankCardData,
			},
		}
	}
	v1 := &TLPaymentsBankCardData{Data2: v}
	v1.Data2.ClassName = ClassPaymentsBankCardData
	return v1
}

func NewTLPaymentsPaymentForm(v *Payments_PaymentForm) *TLPaymentsPaymentForm {
	if v == nil {
		return &TLPaymentsPaymentForm{
			Data2: &Payments_PaymentForm{
				Cmd:       Cmd_PaymentsPaymentForm,
				ClassName: ClassPaymentsPaymentForm,
			},
		}
	}
	v1 := &TLPaymentsPaymentForm{Data2: v}
	v1.Data2.ClassName = ClassPaymentsPaymentForm
	return v1
}

func NewTLPaymentsPaymentReceipt(v *Payments_PaymentReceipt) *TLPaymentsPaymentReceipt {
	if v == nil {
		return &TLPaymentsPaymentReceipt{
			Data2: &Payments_PaymentReceipt{
				Cmd:       Cmd_PaymentsPaymentReceipt,
				ClassName: ClassPaymentsPaymentReceipt,
			},
		}
	}
	v1 := &TLPaymentsPaymentReceipt{Data2: v}
	v1.Data2.ClassName = ClassPaymentsPaymentReceipt
	return v1
}

func NewTLPaymentsPaymentResult(v *Payments_PaymentResult) *TLPaymentsPaymentResult {
	if v == nil {
		return &TLPaymentsPaymentResult{
			Data2: &Payments_PaymentResult{
				Cmd:       Cmd_PaymentsPaymentResult,
				ClassName: ClassPaymentsPaymentResult,
			},
		}
	}
	v1 := &TLPaymentsPaymentResult{Data2: v}
	v1.Data2.ClassName = ClassPaymentsPaymentResult
	return v1
}

func NewTLPaymentsPaymentVerficationNeeded(v *Payments_PaymentResult) *TLPaymentsPaymentVerficationNeeded {
	if v == nil {
		return &TLPaymentsPaymentVerficationNeeded{
			Data2: &Payments_PaymentResult{
				Cmd:       Cmd_PaymentsPaymentVerficationNeeded,
				ClassName: ClassPaymentsPaymentVerficationNeeded,
			},
		}
	}
	v1 := &TLPaymentsPaymentVerficationNeeded{Data2: v}
	v1.Data2.ClassName = ClassPaymentsPaymentVerficationNeeded
	return v1
}

func NewTLPaymentsPaymentVerificationNeeded(v *Payments_PaymentResult) *TLPaymentsPaymentVerificationNeeded {
	if v == nil {
		return &TLPaymentsPaymentVerificationNeeded{
			Data2: &Payments_PaymentResult{
				Cmd:       Cmd_PaymentsPaymentVerificationNeeded,
				ClassName: ClassPaymentsPaymentVerificationNeeded,
			},
		}
	}
	v1 := &TLPaymentsPaymentVerificationNeeded{Data2: v}
	v1.Data2.ClassName = ClassPaymentsPaymentVerificationNeeded
	return v1
}

func NewTLPaymentsSavedInfo(v *Payments_SavedInfo) *TLPaymentsSavedInfo {
	if v == nil {
		return &TLPaymentsSavedInfo{
			Data2: &Payments_SavedInfo{
				Cmd:       Cmd_PaymentsSavedInfo,
				ClassName: ClassPaymentsSavedInfo,
			},
		}
	}
	v1 := &TLPaymentsSavedInfo{Data2: v}
	v1.Data2.ClassName = ClassPaymentsSavedInfo
	return v1
}

func NewTLPaymentsValidatedRequestedInfo(v *Payments_ValidatedRequestedInfo) *TLPaymentsValidatedRequestedInfo {
	if v == nil {
		return &TLPaymentsValidatedRequestedInfo{
			Data2: &Payments_ValidatedRequestedInfo{
				Cmd:       Cmd_PaymentsValidatedRequestedInfo,
				ClassName: ClassPaymentsValidatedRequestedInfo,
			},
		}
	}
	v1 := &TLPaymentsValidatedRequestedInfo{Data2: v}
	v1.Data2.ClassName = ClassPaymentsValidatedRequestedInfo
	return v1
}

func NewTLPeerUser(v *Peer) *TLPeerUser {
	if v == nil {
		return &TLPeerUser{
			Data2: &Peer{
				Cmd:       Cmd_PeerUser,
				ClassName: ClassPeerUser,
			},
		}
	}
	v1 := &TLPeerUser{Data2: v}
	v1.Data2.ClassName = ClassPeerUser
	return v1
}

func NewTLPeerChat(v *Peer) *TLPeerChat {
	if v == nil {
		return &TLPeerChat{
			Data2: &Peer{
				Cmd:       Cmd_PeerChat,
				ClassName: ClassPeerChat,
			},
		}
	}
	v1 := &TLPeerChat{Data2: v}
	v1.Data2.ClassName = ClassPeerChat
	return v1
}

func NewTLPeerChannel(v *Peer) *TLPeerChannel {
	if v == nil {
		return &TLPeerChannel{
			Data2: &Peer{
				Cmd:       Cmd_PeerChannel,
				ClassName: ClassPeerChannel,
			},
		}
	}
	v1 := &TLPeerChannel{Data2: v}
	v1.Data2.ClassName = ClassPeerChannel
	return v1
}

func NewTLPeerBlocked(v *PeerBlocked) *TLPeerBlocked {
	if v == nil {
		return &TLPeerBlocked{
			Data2: &PeerBlocked{
				Cmd:       Cmd_PeerBlocked,
				ClassName: ClassPeerBlocked,
			},
		}
	}
	v1 := &TLPeerBlocked{Data2: v}
	v1.Data2.ClassName = ClassPeerBlocked
	return v1
}

func NewTLPeerLocated(v *PeerLocated) *TLPeerLocated {
	if v == nil {
		return &TLPeerLocated{
			Data2: &PeerLocated{
				Cmd:       Cmd_PeerLocated,
				ClassName: ClassPeerLocated,
			},
		}
	}
	v1 := &TLPeerLocated{Data2: v}
	v1.Data2.ClassName = ClassPeerLocated
	return v1
}

func NewTLPeerSelfLocated(v *PeerLocated) *TLPeerSelfLocated {
	if v == nil {
		return &TLPeerSelfLocated{
			Data2: &PeerLocated{
				Cmd:       Cmd_PeerSelfLocated,
				ClassName: ClassPeerSelfLocated,
			},
		}
	}
	v1 := &TLPeerSelfLocated{Data2: v}
	v1.Data2.ClassName = ClassPeerSelfLocated
	return v1
}

func NewTLPeerNotifyEventsEmpty(v *PeerNotifyEvents) *TLPeerNotifyEventsEmpty {
	if v == nil {
		return &TLPeerNotifyEventsEmpty{
			Data2: &PeerNotifyEvents{
				Cmd:       Cmd_PeerNotifyEventsEmpty,
				ClassName: ClassPeerNotifyEventsEmpty,
			},
		}
	}
	v1 := &TLPeerNotifyEventsEmpty{Data2: v}
	v1.Data2.ClassName = ClassPeerNotifyEventsEmpty
	return v1
}

func NewTLPeerNotifyEventsAll(v *PeerNotifyEvents) *TLPeerNotifyEventsAll {
	if v == nil {
		return &TLPeerNotifyEventsAll{
			Data2: &PeerNotifyEvents{
				Cmd:       Cmd_PeerNotifyEventsAll,
				ClassName: ClassPeerNotifyEventsAll,
			},
		}
	}
	v1 := &TLPeerNotifyEventsAll{Data2: v}
	v1.Data2.ClassName = ClassPeerNotifyEventsAll
	return v1
}

func NewTLPeerNotifySettingsEmpty(v *PeerNotifySettings) *TLPeerNotifySettingsEmpty {
	if v == nil {
		return &TLPeerNotifySettingsEmpty{
			Data2: &PeerNotifySettings{
				Cmd:       Cmd_PeerNotifySettingsEmpty,
				ClassName: ClassPeerNotifySettingsEmpty,
			},
		}
	}
	v1 := &TLPeerNotifySettingsEmpty{Data2: v}
	v1.Data2.ClassName = ClassPeerNotifySettingsEmpty
	return v1
}

func NewTLPeerNotifySettings(v *PeerNotifySettings) *TLPeerNotifySettings {
	if v == nil {
		return &TLPeerNotifySettings{
			Data2: &PeerNotifySettings{
				Cmd:       Cmd_PeerNotifySettings,
				ClassName: ClassPeerNotifySettings,
			},
		}
	}
	v1 := &TLPeerNotifySettings{Data2: v}
	v1.Data2.ClassName = ClassPeerNotifySettings
	return v1
}

func NewTLPeerSettings(v *PeerSettings) *TLPeerSettings {
	if v == nil {
		return &TLPeerSettings{
			Data2: &PeerSettings{
				Cmd:       Cmd_PeerSettings,
				ClassName: ClassPeerSettings,
			},
		}
	}
	v1 := &TLPeerSettings{Data2: v}
	v1.Data2.ClassName = ClassPeerSettings
	return v1
}

func NewTLPhoneCallEmpty(v *PhoneCall) *TLPhoneCallEmpty {
	if v == nil {
		return &TLPhoneCallEmpty{
			Data2: &PhoneCall{
				Cmd:       Cmd_PhoneCallEmpty,
				ClassName: ClassPhoneCallEmpty,
			},
		}
	}
	v1 := &TLPhoneCallEmpty{Data2: v}
	v1.Data2.ClassName = ClassPhoneCallEmpty
	return v1
}

func NewTLPhoneCallRequested(v *PhoneCall) *TLPhoneCallRequested {
	if v == nil {
		return &TLPhoneCallRequested{
			Data2: &PhoneCall{
				Cmd:       Cmd_PhoneCallRequested,
				ClassName: ClassPhoneCallRequested,
			},
		}
	}
	v1 := &TLPhoneCallRequested{Data2: v}
	v1.Data2.ClassName = ClassPhoneCallRequested
	return v1
}

func NewTLPhoneCallAccepted(v *PhoneCall) *TLPhoneCallAccepted {
	if v == nil {
		return &TLPhoneCallAccepted{
			Data2: &PhoneCall{
				Cmd:       Cmd_PhoneCallAccepted,
				ClassName: ClassPhoneCallAccepted,
			},
		}
	}
	v1 := &TLPhoneCallAccepted{Data2: v}
	v1.Data2.ClassName = ClassPhoneCallAccepted
	return v1
}

func NewTLPhoneCall(v *PhoneCall) *TLPhoneCall {
	if v == nil {
		return &TLPhoneCall{
			Data2: &PhoneCall{
				Cmd:       Cmd_PhoneCall,
				ClassName: ClassPhoneCall,
			},
		}
	}
	v1 := &TLPhoneCall{Data2: v}
	v1.Data2.ClassName = ClassPhoneCall
	return v1
}

func NewTLPhoneCallWaiting(v *PhoneCall) *TLPhoneCallWaiting {
	if v == nil {
		return &TLPhoneCallWaiting{
			Data2: &PhoneCall{
				Cmd:       Cmd_PhoneCallWaiting,
				ClassName: ClassPhoneCallWaiting,
			},
		}
	}
	v1 := &TLPhoneCallWaiting{Data2: v}
	v1.Data2.ClassName = ClassPhoneCallWaiting
	return v1
}

func NewTLPhoneCallDiscarded(v *PhoneCall) *TLPhoneCallDiscarded {
	if v == nil {
		return &TLPhoneCallDiscarded{
			Data2: &PhoneCall{
				Cmd:       Cmd_PhoneCallDiscarded,
				ClassName: ClassPhoneCallDiscarded,
			},
		}
	}
	v1 := &TLPhoneCallDiscarded{Data2: v}
	v1.Data2.ClassName = ClassPhoneCallDiscarded
	return v1
}

func NewTLPhoneCallDiscardReasonMissed(v *PhoneCallDiscardReason) *TLPhoneCallDiscardReasonMissed {
	if v == nil {
		return &TLPhoneCallDiscardReasonMissed{
			Data2: &PhoneCallDiscardReason{
				Cmd:       Cmd_PhoneCallDiscardReasonMissed,
				ClassName: ClassPhoneCallDiscardReasonMissed,
			},
		}
	}
	v1 := &TLPhoneCallDiscardReasonMissed{Data2: v}
	v1.Data2.ClassName = ClassPhoneCallDiscardReasonMissed
	return v1
}

func NewTLPhoneCallDiscardReasonDisconnect(v *PhoneCallDiscardReason) *TLPhoneCallDiscardReasonDisconnect {
	if v == nil {
		return &TLPhoneCallDiscardReasonDisconnect{
			Data2: &PhoneCallDiscardReason{
				Cmd:       Cmd_PhoneCallDiscardReasonDisconnect,
				ClassName: ClassPhoneCallDiscardReasonDisconnect,
			},
		}
	}
	v1 := &TLPhoneCallDiscardReasonDisconnect{Data2: v}
	v1.Data2.ClassName = ClassPhoneCallDiscardReasonDisconnect
	return v1
}

func NewTLPhoneCallDiscardReasonHangup(v *PhoneCallDiscardReason) *TLPhoneCallDiscardReasonHangup {
	if v == nil {
		return &TLPhoneCallDiscardReasonHangup{
			Data2: &PhoneCallDiscardReason{
				Cmd:       Cmd_PhoneCallDiscardReasonHangup,
				ClassName: ClassPhoneCallDiscardReasonHangup,
			},
		}
	}
	v1 := &TLPhoneCallDiscardReasonHangup{Data2: v}
	v1.Data2.ClassName = ClassPhoneCallDiscardReasonHangup
	return v1
}

func NewTLPhoneCallDiscardReasonBusy(v *PhoneCallDiscardReason) *TLPhoneCallDiscardReasonBusy {
	if v == nil {
		return &TLPhoneCallDiscardReasonBusy{
			Data2: &PhoneCallDiscardReason{
				Cmd:       Cmd_PhoneCallDiscardReasonBusy,
				ClassName: ClassPhoneCallDiscardReasonBusy,
			},
		}
	}
	v1 := &TLPhoneCallDiscardReasonBusy{Data2: v}
	v1.Data2.ClassName = ClassPhoneCallDiscardReasonBusy
	return v1
}

func NewTLPhoneCallProtocol(v *PhoneCallProtocol) *TLPhoneCallProtocol {
	if v == nil {
		return &TLPhoneCallProtocol{
			Data2: &PhoneCallProtocol{
				Cmd:       Cmd_PhoneCallProtocol,
				ClassName: ClassPhoneCallProtocol,
			},
		}
	}
	v1 := &TLPhoneCallProtocol{Data2: v}
	v1.Data2.ClassName = ClassPhoneCallProtocol
	return v1
}

func NewTLPhoneConnection(v *PhoneConnection) *TLPhoneConnection {
	if v == nil {
		return &TLPhoneConnection{
			Data2: &PhoneConnection{
				Cmd:       Cmd_PhoneConnection,
				ClassName: ClassPhoneConnection,
			},
		}
	}
	v1 := &TLPhoneConnection{Data2: v}
	v1.Data2.ClassName = ClassPhoneConnection
	return v1
}

func NewTLPhoneConnectionWebrtc(v *PhoneConnection) *TLPhoneConnectionWebrtc {
	if v == nil {
		return &TLPhoneConnectionWebrtc{
			Data2: &PhoneConnection{
				Cmd:       Cmd_PhoneConnectionWebrtc,
				ClassName: ClassPhoneConnectionWebrtc,
			},
		}
	}
	v1 := &TLPhoneConnectionWebrtc{Data2: v}
	v1.Data2.ClassName = ClassPhoneConnectionWebrtc
	return v1
}

func NewTLPhoneGroupCall(v *Phone_GroupCall) *TLPhoneGroupCall {
	if v == nil {
		return &TLPhoneGroupCall{
			Data2: &Phone_GroupCall{
				Cmd:       Cmd_PhoneGroupCall,
				ClassName: ClassPhoneGroupCall,
			},
		}
	}
	v1 := &TLPhoneGroupCall{Data2: v}
	v1.Data2.ClassName = ClassPhoneGroupCall
	return v1
}

func NewTLPhoneGroupParticipants(v *Phone_GroupParticipants) *TLPhoneGroupParticipants {
	if v == nil {
		return &TLPhoneGroupParticipants{
			Data2: &Phone_GroupParticipants{
				Cmd:       Cmd_PhoneGroupParticipants,
				ClassName: ClassPhoneGroupParticipants,
			},
		}
	}
	v1 := &TLPhoneGroupParticipants{Data2: v}
	v1.Data2.ClassName = ClassPhoneGroupParticipants
	return v1
}

func NewTLPhonePhoneCall(v *Phone_PhoneCall) *TLPhonePhoneCall {
	if v == nil {
		return &TLPhonePhoneCall{
			Data2: &Phone_PhoneCall{
				Cmd:       Cmd_PhonePhoneCall,
				ClassName: ClassPhonePhoneCall,
			},
		}
	}
	v1 := &TLPhonePhoneCall{Data2: v}
	v1.Data2.ClassName = ClassPhonePhoneCall
	return v1
}

func NewTLPhotoEmpty(v *Photo) *TLPhotoEmpty {
	if v == nil {
		return &TLPhotoEmpty{
			Data2: &Photo{
				Cmd:       Cmd_PhotoEmpty,
				ClassName: ClassPhotoEmpty,
			},
		}
	}
	v1 := &TLPhotoEmpty{Data2: v}
	v1.Data2.ClassName = ClassPhotoEmpty
	return v1
}

func NewTLPhoto(v *Photo) *TLPhoto {
	if v == nil {
		return &TLPhoto{
			Data2: &Photo{
				Cmd:       Cmd_Photo,
				ClassName: ClassPhoto,
			},
		}
	}
	v1 := &TLPhoto{Data2: v}
	v1.Data2.ClassName = ClassPhoto
	return v1
}

func NewTLPhotoSizeEmpty(v *PhotoSize) *TLPhotoSizeEmpty {
	if v == nil {
		return &TLPhotoSizeEmpty{
			Data2: &PhotoSize{
				Cmd:       Cmd_PhotoSizeEmpty,
				ClassName: ClassPhotoSizeEmpty,
			},
		}
	}
	v1 := &TLPhotoSizeEmpty{Data2: v}
	v1.Data2.ClassName = ClassPhotoSizeEmpty
	return v1
}

func NewTLPhotoSize(v *PhotoSize) *TLPhotoSize {
	if v == nil {
		return &TLPhotoSize{
			Data2: &PhotoSize{
				Cmd:       Cmd_PhotoSize,
				ClassName: ClassPhotoSize,
			},
		}
	}
	v1 := &TLPhotoSize{Data2: v}
	v1.Data2.ClassName = ClassPhotoSize
	return v1
}

func NewTLPhotoCachedSize(v *PhotoSize) *TLPhotoCachedSize {
	if v == nil {
		return &TLPhotoCachedSize{
			Data2: &PhotoSize{
				Cmd:       Cmd_PhotoCachedSize,
				ClassName: ClassPhotoCachedSize,
			},
		}
	}
	v1 := &TLPhotoCachedSize{Data2: v}
	v1.Data2.ClassName = ClassPhotoCachedSize
	return v1
}

func NewTLPhotoStrippedSize(v *PhotoSize) *TLPhotoStrippedSize {
	if v == nil {
		return &TLPhotoStrippedSize{
			Data2: &PhotoSize{
				Cmd:       Cmd_PhotoStrippedSize,
				ClassName: ClassPhotoStrippedSize,
			},
		}
	}
	v1 := &TLPhotoStrippedSize{Data2: v}
	v1.Data2.ClassName = ClassPhotoStrippedSize
	return v1
}

func NewTLPhotoSizeProgressive(v *PhotoSize) *TLPhotoSizeProgressive {
	if v == nil {
		return &TLPhotoSizeProgressive{
			Data2: &PhotoSize{
				Cmd:       Cmd_PhotoSizeProgressive,
				ClassName: ClassPhotoSizeProgressive,
			},
		}
	}
	v1 := &TLPhotoSizeProgressive{Data2: v}
	v1.Data2.ClassName = ClassPhotoSizeProgressive
	return v1
}

func NewTLPhotoPathSize(v *PhotoSize) *TLPhotoPathSize {
	if v == nil {
		return &TLPhotoPathSize{
			Data2: &PhotoSize{
				Cmd:       Cmd_PhotoPathSize,
				ClassName: ClassPhotoPathSize,
			},
		}
	}
	v1 := &TLPhotoPathSize{Data2: v}
	v1.Data2.ClassName = ClassPhotoPathSize
	return v1
}

func NewTLPhotosPhoto(v *Photos_Photo) *TLPhotosPhoto {
	if v == nil {
		return &TLPhotosPhoto{
			Data2: &Photos_Photo{
				Cmd:       Cmd_PhotosPhoto,
				ClassName: ClassPhotosPhoto,
			},
		}
	}
	v1 := &TLPhotosPhoto{Data2: v}
	v1.Data2.ClassName = ClassPhotosPhoto
	return v1
}

func NewTLPhotosPhotos(v *Photos_Photos) *TLPhotosPhotos {
	if v == nil {
		return &TLPhotosPhotos{
			Data2: &Photos_Photos{
				Cmd:       Cmd_PhotosPhotos,
				ClassName: ClassPhotosPhotos,
			},
		}
	}
	v1 := &TLPhotosPhotos{Data2: v}
	v1.Data2.ClassName = ClassPhotosPhotos
	return v1
}

func NewTLPhotosPhotosSlice(v *Photos_Photos) *TLPhotosPhotosSlice {
	if v == nil {
		return &TLPhotosPhotosSlice{
			Data2: &Photos_Photos{
				Cmd:       Cmd_PhotosPhotosSlice,
				ClassName: ClassPhotosPhotosSlice,
			},
		}
	}
	v1 := &TLPhotosPhotosSlice{Data2: v}
	v1.Data2.ClassName = ClassPhotosPhotosSlice
	return v1
}

func NewTLPoll(v *Poll) *TLPoll {
	if v == nil {
		return &TLPoll{
			Data2: &Poll{
				Cmd:       Cmd_Poll,
				ClassName: ClassPoll,
			},
		}
	}
	v1 := &TLPoll{Data2: v}
	v1.Data2.ClassName = ClassPoll
	return v1
}

func NewTLPollAnswer(v *PollAnswer) *TLPollAnswer {
	if v == nil {
		return &TLPollAnswer{
			Data2: &PollAnswer{
				Cmd:       Cmd_PollAnswer,
				ClassName: ClassPollAnswer,
			},
		}
	}
	v1 := &TLPollAnswer{Data2: v}
	v1.Data2.ClassName = ClassPollAnswer
	return v1
}

func NewTLPollAnswerVoters(v *PollAnswerVoters) *TLPollAnswerVoters {
	if v == nil {
		return &TLPollAnswerVoters{
			Data2: &PollAnswerVoters{
				Cmd:       Cmd_PollAnswerVoters,
				ClassName: ClassPollAnswerVoters,
			},
		}
	}
	v1 := &TLPollAnswerVoters{Data2: v}
	v1.Data2.ClassName = ClassPollAnswerVoters
	return v1
}

func NewTLPollResults(v *PollResults) *TLPollResults {
	if v == nil {
		return &TLPollResults{
			Data2: &PollResults{
				Cmd:       Cmd_PollResults,
				ClassName: ClassPollResults,
			},
		}
	}
	v1 := &TLPollResults{Data2: v}
	v1.Data2.ClassName = ClassPollResults
	return v1
}

func NewTLPong(v *Pong) *TLPong {
	if v == nil {
		return &TLPong{
			Data2: &Pong{
				Cmd:       Cmd_Pong,
				ClassName: ClassPong,
			},
		}
	}
	v1 := &TLPong{Data2: v}
	v1.Data2.ClassName = ClassPong
	return v1
}

func NewTLPopularContact(v *PopularContact) *TLPopularContact {
	if v == nil {
		return &TLPopularContact{
			Data2: &PopularContact{
				Cmd:       Cmd_PopularContact,
				ClassName: ClassPopularContact,
			},
		}
	}
	v1 := &TLPopularContact{Data2: v}
	v1.Data2.ClassName = ClassPopularContact
	return v1
}

func NewTLPostAddress(v *PostAddress) *TLPostAddress {
	if v == nil {
		return &TLPostAddress{
			Data2: &PostAddress{
				Cmd:       Cmd_PostAddress,
				ClassName: ClassPostAddress,
			},
		}
	}
	v1 := &TLPostAddress{Data2: v}
	v1.Data2.ClassName = ClassPostAddress
	return v1
}

func NewTLPrivacyKeyStatusTimestamp(v *PrivacyKey) *TLPrivacyKeyStatusTimestamp {
	if v == nil {
		return &TLPrivacyKeyStatusTimestamp{
			Data2: &PrivacyKey{
				Cmd:       Cmd_PrivacyKeyStatusTimestamp,
				ClassName: ClassPrivacyKeyStatusTimestamp,
			},
		}
	}
	v1 := &TLPrivacyKeyStatusTimestamp{Data2: v}
	v1.Data2.ClassName = ClassPrivacyKeyStatusTimestamp
	return v1
}

func NewTLPrivacyKeyChatInvite(v *PrivacyKey) *TLPrivacyKeyChatInvite {
	if v == nil {
		return &TLPrivacyKeyChatInvite{
			Data2: &PrivacyKey{
				Cmd:       Cmd_PrivacyKeyChatInvite,
				ClassName: ClassPrivacyKeyChatInvite,
			},
		}
	}
	v1 := &TLPrivacyKeyChatInvite{Data2: v}
	v1.Data2.ClassName = ClassPrivacyKeyChatInvite
	return v1
}

func NewTLPrivacyKeyPhoneCall(v *PrivacyKey) *TLPrivacyKeyPhoneCall {
	if v == nil {
		return &TLPrivacyKeyPhoneCall{
			Data2: &PrivacyKey{
				Cmd:       Cmd_PrivacyKeyPhoneCall,
				ClassName: ClassPrivacyKeyPhoneCall,
			},
		}
	}
	v1 := &TLPrivacyKeyPhoneCall{Data2: v}
	v1.Data2.ClassName = ClassPrivacyKeyPhoneCall
	return v1
}

func NewTLPrivacyKeyPhoneP2P(v *PrivacyKey) *TLPrivacyKeyPhoneP2P {
	if v == nil {
		return &TLPrivacyKeyPhoneP2P{
			Data2: &PrivacyKey{
				Cmd:       Cmd_PrivacyKeyPhoneP2P,
				ClassName: ClassPrivacyKeyPhoneP2P,
			},
		}
	}
	v1 := &TLPrivacyKeyPhoneP2P{Data2: v}
	v1.Data2.ClassName = ClassPrivacyKeyPhoneP2P
	return v1
}

func NewTLPrivacyKeyForwards(v *PrivacyKey) *TLPrivacyKeyForwards {
	if v == nil {
		return &TLPrivacyKeyForwards{
			Data2: &PrivacyKey{
				Cmd:       Cmd_PrivacyKeyForwards,
				ClassName: ClassPrivacyKeyForwards,
			},
		}
	}
	v1 := &TLPrivacyKeyForwards{Data2: v}
	v1.Data2.ClassName = ClassPrivacyKeyForwards
	return v1
}

func NewTLPrivacyKeyProfilePhoto(v *PrivacyKey) *TLPrivacyKeyProfilePhoto {
	if v == nil {
		return &TLPrivacyKeyProfilePhoto{
			Data2: &PrivacyKey{
				Cmd:       Cmd_PrivacyKeyProfilePhoto,
				ClassName: ClassPrivacyKeyProfilePhoto,
			},
		}
	}
	v1 := &TLPrivacyKeyProfilePhoto{Data2: v}
	v1.Data2.ClassName = ClassPrivacyKeyProfilePhoto
	return v1
}

func NewTLPrivacyKeyPhoneNumber(v *PrivacyKey) *TLPrivacyKeyPhoneNumber {
	if v == nil {
		return &TLPrivacyKeyPhoneNumber{
			Data2: &PrivacyKey{
				Cmd:       Cmd_PrivacyKeyPhoneNumber,
				ClassName: ClassPrivacyKeyPhoneNumber,
			},
		}
	}
	v1 := &TLPrivacyKeyPhoneNumber{Data2: v}
	v1.Data2.ClassName = ClassPrivacyKeyPhoneNumber
	return v1
}

func NewTLPrivacyKeyAddedByPhone(v *PrivacyKey) *TLPrivacyKeyAddedByPhone {
	if v == nil {
		return &TLPrivacyKeyAddedByPhone{
			Data2: &PrivacyKey{
				Cmd:       Cmd_PrivacyKeyAddedByPhone,
				ClassName: ClassPrivacyKeyAddedByPhone,
			},
		}
	}
	v1 := &TLPrivacyKeyAddedByPhone{Data2: v}
	v1.Data2.ClassName = ClassPrivacyKeyAddedByPhone
	return v1
}

func NewTLPrivacyValueAllowContacts(v *PrivacyRule) *TLPrivacyValueAllowContacts {
	if v == nil {
		return &TLPrivacyValueAllowContacts{
			Data2: &PrivacyRule{
				Cmd:       Cmd_PrivacyValueAllowContacts,
				ClassName: ClassPrivacyValueAllowContacts,
			},
		}
	}
	v1 := &TLPrivacyValueAllowContacts{Data2: v}
	v1.Data2.ClassName = ClassPrivacyValueAllowContacts
	return v1
}

func NewTLPrivacyValueAllowAll(v *PrivacyRule) *TLPrivacyValueAllowAll {
	if v == nil {
		return &TLPrivacyValueAllowAll{
			Data2: &PrivacyRule{
				Cmd:       Cmd_PrivacyValueAllowAll,
				ClassName: ClassPrivacyValueAllowAll,
			},
		}
	}
	v1 := &TLPrivacyValueAllowAll{Data2: v}
	v1.Data2.ClassName = ClassPrivacyValueAllowAll
	return v1
}

func NewTLPrivacyValueAllowUsers(v *PrivacyRule) *TLPrivacyValueAllowUsers {
	if v == nil {
		return &TLPrivacyValueAllowUsers{
			Data2: &PrivacyRule{
				Cmd:       Cmd_PrivacyValueAllowUsers,
				ClassName: ClassPrivacyValueAllowUsers,
			},
		}
	}
	v1 := &TLPrivacyValueAllowUsers{Data2: v}
	v1.Data2.ClassName = ClassPrivacyValueAllowUsers
	return v1
}

func NewTLPrivacyValueDisallowContacts(v *PrivacyRule) *TLPrivacyValueDisallowContacts {
	if v == nil {
		return &TLPrivacyValueDisallowContacts{
			Data2: &PrivacyRule{
				Cmd:       Cmd_PrivacyValueDisallowContacts,
				ClassName: ClassPrivacyValueDisallowContacts,
			},
		}
	}
	v1 := &TLPrivacyValueDisallowContacts{Data2: v}
	v1.Data2.ClassName = ClassPrivacyValueDisallowContacts
	return v1
}

func NewTLPrivacyValueDisallowAll(v *PrivacyRule) *TLPrivacyValueDisallowAll {
	if v == nil {
		return &TLPrivacyValueDisallowAll{
			Data2: &PrivacyRule{
				Cmd:       Cmd_PrivacyValueDisallowAll,
				ClassName: ClassPrivacyValueDisallowAll,
			},
		}
	}
	v1 := &TLPrivacyValueDisallowAll{Data2: v}
	v1.Data2.ClassName = ClassPrivacyValueDisallowAll
	return v1
}

func NewTLPrivacyValueDisallowUsers(v *PrivacyRule) *TLPrivacyValueDisallowUsers {
	if v == nil {
		return &TLPrivacyValueDisallowUsers{
			Data2: &PrivacyRule{
				Cmd:       Cmd_PrivacyValueDisallowUsers,
				ClassName: ClassPrivacyValueDisallowUsers,
			},
		}
	}
	v1 := &TLPrivacyValueDisallowUsers{Data2: v}
	v1.Data2.ClassName = ClassPrivacyValueDisallowUsers
	return v1
}

func NewTLPrivacyValueAllowChatParticipants(v *PrivacyRule) *TLPrivacyValueAllowChatParticipants {
	if v == nil {
		return &TLPrivacyValueAllowChatParticipants{
			Data2: &PrivacyRule{
				Cmd:       Cmd_PrivacyValueAllowChatParticipants,
				ClassName: ClassPrivacyValueAllowChatParticipants,
			},
		}
	}
	v1 := &TLPrivacyValueAllowChatParticipants{Data2: v}
	v1.Data2.ClassName = ClassPrivacyValueAllowChatParticipants
	return v1
}

func NewTLPrivacyValueDisallowChatParticipants(v *PrivacyRule) *TLPrivacyValueDisallowChatParticipants {
	if v == nil {
		return &TLPrivacyValueDisallowChatParticipants{
			Data2: &PrivacyRule{
				Cmd:       Cmd_PrivacyValueDisallowChatParticipants,
				ClassName: ClassPrivacyValueDisallowChatParticipants,
			},
		}
	}
	v1 := &TLPrivacyValueDisallowChatParticipants{Data2: v}
	v1.Data2.ClassName = ClassPrivacyValueDisallowChatParticipants
	return v1
}

func NewTLReceivedNotifyMessage(v *ReceivedNotifyMessage) *TLReceivedNotifyMessage {
	if v == nil {
		return &TLReceivedNotifyMessage{
			Data2: &ReceivedNotifyMessage{
				Cmd:       Cmd_ReceivedNotifyMessage,
				ClassName: ClassReceivedNotifyMessage,
			},
		}
	}
	v1 := &TLReceivedNotifyMessage{Data2: v}
	v1.Data2.ClassName = ClassReceivedNotifyMessage
	return v1
}

func NewTLRecentMeUrlUnknown(v *RecentMeUrl) *TLRecentMeUrlUnknown {
	if v == nil {
		return &TLRecentMeUrlUnknown{
			Data2: &RecentMeUrl{
				Cmd:       Cmd_RecentMeUrlUnknown,
				ClassName: ClassRecentMeUrlUnknown,
			},
		}
	}
	v1 := &TLRecentMeUrlUnknown{Data2: v}
	v1.Data2.ClassName = ClassRecentMeUrlUnknown
	return v1
}

func NewTLRecentMeUrlUser(v *RecentMeUrl) *TLRecentMeUrlUser {
	if v == nil {
		return &TLRecentMeUrlUser{
			Data2: &RecentMeUrl{
				Cmd:       Cmd_RecentMeUrlUser,
				ClassName: ClassRecentMeUrlUser,
			},
		}
	}
	v1 := &TLRecentMeUrlUser{Data2: v}
	v1.Data2.ClassName = ClassRecentMeUrlUser
	return v1
}

func NewTLRecentMeUrlChat(v *RecentMeUrl) *TLRecentMeUrlChat {
	if v == nil {
		return &TLRecentMeUrlChat{
			Data2: &RecentMeUrl{
				Cmd:       Cmd_RecentMeUrlChat,
				ClassName: ClassRecentMeUrlChat,
			},
		}
	}
	v1 := &TLRecentMeUrlChat{Data2: v}
	v1.Data2.ClassName = ClassRecentMeUrlChat
	return v1
}

func NewTLRecentMeUrlChatInvite(v *RecentMeUrl) *TLRecentMeUrlChatInvite {
	if v == nil {
		return &TLRecentMeUrlChatInvite{
			Data2: &RecentMeUrl{
				Cmd:       Cmd_RecentMeUrlChatInvite,
				ClassName: ClassRecentMeUrlChatInvite,
			},
		}
	}
	v1 := &TLRecentMeUrlChatInvite{Data2: v}
	v1.Data2.ClassName = ClassRecentMeUrlChatInvite
	return v1
}

func NewTLRecentMeUrlStickerSet(v *RecentMeUrl) *TLRecentMeUrlStickerSet {
	if v == nil {
		return &TLRecentMeUrlStickerSet{
			Data2: &RecentMeUrl{
				Cmd:       Cmd_RecentMeUrlStickerSet,
				ClassName: ClassRecentMeUrlStickerSet,
			},
		}
	}
	v1 := &TLRecentMeUrlStickerSet{Data2: v}
	v1.Data2.ClassName = ClassRecentMeUrlStickerSet
	return v1
}

func NewTLReplyKeyboardHide(v *ReplyMarkup) *TLReplyKeyboardHide {
	if v == nil {
		return &TLReplyKeyboardHide{
			Data2: &ReplyMarkup{
				Cmd:       Cmd_ReplyKeyboardHide,
				ClassName: ClassReplyKeyboardHide,
			},
		}
	}
	v1 := &TLReplyKeyboardHide{Data2: v}
	v1.Data2.ClassName = ClassReplyKeyboardHide
	return v1
}

func NewTLReplyKeyboardForceReply(v *ReplyMarkup) *TLReplyKeyboardForceReply {
	if v == nil {
		return &TLReplyKeyboardForceReply{
			Data2: &ReplyMarkup{
				Cmd:       Cmd_ReplyKeyboardForceReply,
				ClassName: ClassReplyKeyboardForceReply,
			},
		}
	}
	v1 := &TLReplyKeyboardForceReply{Data2: v}
	v1.Data2.ClassName = ClassReplyKeyboardForceReply
	return v1
}

func NewTLReplyKeyboardMarkup(v *ReplyMarkup) *TLReplyKeyboardMarkup {
	if v == nil {
		return &TLReplyKeyboardMarkup{
			Data2: &ReplyMarkup{
				Cmd:       Cmd_ReplyKeyboardMarkup,
				ClassName: ClassReplyKeyboardMarkup,
			},
		}
	}
	v1 := &TLReplyKeyboardMarkup{Data2: v}
	v1.Data2.ClassName = ClassReplyKeyboardMarkup
	return v1
}

func NewTLReplyInlineMarkup(v *ReplyMarkup) *TLReplyInlineMarkup {
	if v == nil {
		return &TLReplyInlineMarkup{
			Data2: &ReplyMarkup{
				Cmd:       Cmd_ReplyInlineMarkup,
				ClassName: ClassReplyInlineMarkup,
			},
		}
	}
	v1 := &TLReplyInlineMarkup{Data2: v}
	v1.Data2.ClassName = ClassReplyInlineMarkup
	return v1
}

func NewTLInputReportReasonSpam(v *ReportReason) *TLInputReportReasonSpam {
	if v == nil {
		return &TLInputReportReasonSpam{
			Data2: &ReportReason{
				Cmd:       Cmd_InputReportReasonSpam,
				ClassName: ClassInputReportReasonSpam,
			},
		}
	}
	v1 := &TLInputReportReasonSpam{Data2: v}
	v1.Data2.ClassName = ClassInputReportReasonSpam
	return v1
}

func NewTLInputReportReasonViolence(v *ReportReason) *TLInputReportReasonViolence {
	if v == nil {
		return &TLInputReportReasonViolence{
			Data2: &ReportReason{
				Cmd:       Cmd_InputReportReasonViolence,
				ClassName: ClassInputReportReasonViolence,
			},
		}
	}
	v1 := &TLInputReportReasonViolence{Data2: v}
	v1.Data2.ClassName = ClassInputReportReasonViolence
	return v1
}

func NewTLInputReportReasonPornography(v *ReportReason) *TLInputReportReasonPornography {
	if v == nil {
		return &TLInputReportReasonPornography{
			Data2: &ReportReason{
				Cmd:       Cmd_InputReportReasonPornography,
				ClassName: ClassInputReportReasonPornography,
			},
		}
	}
	v1 := &TLInputReportReasonPornography{Data2: v}
	v1.Data2.ClassName = ClassInputReportReasonPornography
	return v1
}

func NewTLInputReportReasonOther(v *ReportReason) *TLInputReportReasonOther {
	if v == nil {
		return &TLInputReportReasonOther{
			Data2: &ReportReason{
				Cmd:       Cmd_InputReportReasonOther,
				ClassName: ClassInputReportReasonOther,
			},
		}
	}
	v1 := &TLInputReportReasonOther{Data2: v}
	v1.Data2.ClassName = ClassInputReportReasonOther
	return v1
}

func NewTLInputReportReasonCopyright(v *ReportReason) *TLInputReportReasonCopyright {
	if v == nil {
		return &TLInputReportReasonCopyright{
			Data2: &ReportReason{
				Cmd:       Cmd_InputReportReasonCopyright,
				ClassName: ClassInputReportReasonCopyright,
			},
		}
	}
	v1 := &TLInputReportReasonCopyright{Data2: v}
	v1.Data2.ClassName = ClassInputReportReasonCopyright
	return v1
}

func NewTLInputReportReasonChildAbuse(v *ReportReason) *TLInputReportReasonChildAbuse {
	if v == nil {
		return &TLInputReportReasonChildAbuse{
			Data2: &ReportReason{
				Cmd:       Cmd_InputReportReasonChildAbuse,
				ClassName: ClassInputReportReasonChildAbuse,
			},
		}
	}
	v1 := &TLInputReportReasonChildAbuse{Data2: v}
	v1.Data2.ClassName = ClassInputReportReasonChildAbuse
	return v1
}

func NewTLInputReportReasonGeoIrrelevant(v *ReportReason) *TLInputReportReasonGeoIrrelevant {
	if v == nil {
		return &TLInputReportReasonGeoIrrelevant{
			Data2: &ReportReason{
				Cmd:       Cmd_InputReportReasonGeoIrrelevant,
				ClassName: ClassInputReportReasonGeoIrrelevant,
			},
		}
	}
	v1 := &TLInputReportReasonGeoIrrelevant{Data2: v}
	v1.Data2.ClassName = ClassInputReportReasonGeoIrrelevant
	return v1
}

func NewTLResPQ(v *ResPQ) *TLResPQ {
	if v == nil {
		return &TLResPQ{
			Data2: &ResPQ{
				Cmd:       Cmd_ResPQ,
				ClassName: ClassResPQ,
			},
		}
	}
	v1 := &TLResPQ{Data2: v}
	v1.Data2.ClassName = ClassResPQ
	return v1
}

func NewTLRestrictionReason(v *RestrictionReason) *TLRestrictionReason {
	if v == nil {
		return &TLRestrictionReason{
			Data2: &RestrictionReason{
				Cmd:       Cmd_RestrictionReason,
				ClassName: ClassRestrictionReason,
			},
		}
	}
	v1 := &TLRestrictionReason{Data2: v}
	v1.Data2.ClassName = ClassRestrictionReason
	return v1
}

func NewTLTextEmpty(v *RichText) *TLTextEmpty {
	if v == nil {
		return &TLTextEmpty{
			Data2: &RichText{
				Cmd:       Cmd_TextEmpty,
				ClassName: ClassTextEmpty,
			},
		}
	}
	v1 := &TLTextEmpty{Data2: v}
	v1.Data2.ClassName = ClassTextEmpty
	return v1
}

func NewTLTextPlain(v *RichText) *TLTextPlain {
	if v == nil {
		return &TLTextPlain{
			Data2: &RichText{
				Cmd:       Cmd_TextPlain,
				ClassName: ClassTextPlain,
			},
		}
	}
	v1 := &TLTextPlain{Data2: v}
	v1.Data2.ClassName = ClassTextPlain
	return v1
}

func NewTLTextBold(v *RichText) *TLTextBold {
	if v == nil {
		return &TLTextBold{
			Data2: &RichText{
				Cmd:       Cmd_TextBold,
				ClassName: ClassTextBold,
			},
		}
	}
	v1 := &TLTextBold{Data2: v}
	v1.Data2.ClassName = ClassTextBold
	return v1
}

func NewTLTextItalic(v *RichText) *TLTextItalic {
	if v == nil {
		return &TLTextItalic{
			Data2: &RichText{
				Cmd:       Cmd_TextItalic,
				ClassName: ClassTextItalic,
			},
		}
	}
	v1 := &TLTextItalic{Data2: v}
	v1.Data2.ClassName = ClassTextItalic
	return v1
}

func NewTLTextUnderline(v *RichText) *TLTextUnderline {
	if v == nil {
		return &TLTextUnderline{
			Data2: &RichText{
				Cmd:       Cmd_TextUnderline,
				ClassName: ClassTextUnderline,
			},
		}
	}
	v1 := &TLTextUnderline{Data2: v}
	v1.Data2.ClassName = ClassTextUnderline
	return v1
}

func NewTLTextStrike(v *RichText) *TLTextStrike {
	if v == nil {
		return &TLTextStrike{
			Data2: &RichText{
				Cmd:       Cmd_TextStrike,
				ClassName: ClassTextStrike,
			},
		}
	}
	v1 := &TLTextStrike{Data2: v}
	v1.Data2.ClassName = ClassTextStrike
	return v1
}

func NewTLTextFixed(v *RichText) *TLTextFixed {
	if v == nil {
		return &TLTextFixed{
			Data2: &RichText{
				Cmd:       Cmd_TextFixed,
				ClassName: ClassTextFixed,
			},
		}
	}
	v1 := &TLTextFixed{Data2: v}
	v1.Data2.ClassName = ClassTextFixed
	return v1
}

func NewTLTextUrl(v *RichText) *TLTextUrl {
	if v == nil {
		return &TLTextUrl{
			Data2: &RichText{
				Cmd:       Cmd_TextUrl,
				ClassName: ClassTextUrl,
			},
		}
	}
	v1 := &TLTextUrl{Data2: v}
	v1.Data2.ClassName = ClassTextUrl
	return v1
}

func NewTLTextEmail(v *RichText) *TLTextEmail {
	if v == nil {
		return &TLTextEmail{
			Data2: &RichText{
				Cmd:       Cmd_TextEmail,
				ClassName: ClassTextEmail,
			},
		}
	}
	v1 := &TLTextEmail{Data2: v}
	v1.Data2.ClassName = ClassTextEmail
	return v1
}

func NewTLTextConcat(v *RichText) *TLTextConcat {
	if v == nil {
		return &TLTextConcat{
			Data2: &RichText{
				Cmd:       Cmd_TextConcat,
				ClassName: ClassTextConcat,
			},
		}
	}
	v1 := &TLTextConcat{Data2: v}
	v1.Data2.ClassName = ClassTextConcat
	return v1
}

func NewTLTextSubscript(v *RichText) *TLTextSubscript {
	if v == nil {
		return &TLTextSubscript{
			Data2: &RichText{
				Cmd:       Cmd_TextSubscript,
				ClassName: ClassTextSubscript,
			},
		}
	}
	v1 := &TLTextSubscript{Data2: v}
	v1.Data2.ClassName = ClassTextSubscript
	return v1
}

func NewTLTextSuperscript(v *RichText) *TLTextSuperscript {
	if v == nil {
		return &TLTextSuperscript{
			Data2: &RichText{
				Cmd:       Cmd_TextSuperscript,
				ClassName: ClassTextSuperscript,
			},
		}
	}
	v1 := &TLTextSuperscript{Data2: v}
	v1.Data2.ClassName = ClassTextSuperscript
	return v1
}

func NewTLTextMarked(v *RichText) *TLTextMarked {
	if v == nil {
		return &TLTextMarked{
			Data2: &RichText{
				Cmd:       Cmd_TextMarked,
				ClassName: ClassTextMarked,
			},
		}
	}
	v1 := &TLTextMarked{Data2: v}
	v1.Data2.ClassName = ClassTextMarked
	return v1
}

func NewTLTextPhone(v *RichText) *TLTextPhone {
	if v == nil {
		return &TLTextPhone{
			Data2: &RichText{
				Cmd:       Cmd_TextPhone,
				ClassName: ClassTextPhone,
			},
		}
	}
	v1 := &TLTextPhone{Data2: v}
	v1.Data2.ClassName = ClassTextPhone
	return v1
}

func NewTLTextImage(v *RichText) *TLTextImage {
	if v == nil {
		return &TLTextImage{
			Data2: &RichText{
				Cmd:       Cmd_TextImage,
				ClassName: ClassTextImage,
			},
		}
	}
	v1 := &TLTextImage{Data2: v}
	v1.Data2.ClassName = ClassTextImage
	return v1
}

func NewTLTextAnchor(v *RichText) *TLTextAnchor {
	if v == nil {
		return &TLTextAnchor{
			Data2: &RichText{
				Cmd:       Cmd_TextAnchor,
				ClassName: ClassTextAnchor,
			},
		}
	}
	v1 := &TLTextAnchor{Data2: v}
	v1.Data2.ClassName = ClassTextAnchor
	return v1
}

func NewTLRpcAnswerUnknown(v *RpcDropAnswer) *TLRpcAnswerUnknown {
	if v == nil {
		return &TLRpcAnswerUnknown{
			Data2: &RpcDropAnswer{
				Cmd:       Cmd_RpcAnswerUnknown,
				ClassName: ClassRpcAnswerUnknown,
			},
		}
	}
	v1 := &TLRpcAnswerUnknown{Data2: v}
	v1.Data2.ClassName = ClassRpcAnswerUnknown
	return v1
}

func NewTLRpcAnswerDroppedRunning(v *RpcDropAnswer) *TLRpcAnswerDroppedRunning {
	if v == nil {
		return &TLRpcAnswerDroppedRunning{
			Data2: &RpcDropAnswer{
				Cmd:       Cmd_RpcAnswerDroppedRunning,
				ClassName: ClassRpcAnswerDroppedRunning,
			},
		}
	}
	v1 := &TLRpcAnswerDroppedRunning{Data2: v}
	v1.Data2.ClassName = ClassRpcAnswerDroppedRunning
	return v1
}

func NewTLRpcAnswerDropped(v *RpcDropAnswer) *TLRpcAnswerDropped {
	if v == nil {
		return &TLRpcAnswerDropped{
			Data2: &RpcDropAnswer{
				Cmd:       Cmd_RpcAnswerDropped,
				ClassName: ClassRpcAnswerDropped,
			},
		}
	}
	v1 := &TLRpcAnswerDropped{Data2: v}
	v1.Data2.ClassName = ClassRpcAnswerDropped
	return v1
}

func NewTLRpcError(v *RpcError) *TLRpcError {
	if v == nil {
		return &TLRpcError{
			Data2: &RpcError{
				Cmd:       Cmd_RpcError,
				ClassName: ClassRpcError,
			},
		}
	}
	v1 := &TLRpcError{Data2: v}
	v1.Data2.ClassName = ClassRpcError
	return v1
}

func NewTLSavedPhoneContact(v *SavedContact) *TLSavedPhoneContact {
	if v == nil {
		return &TLSavedPhoneContact{
			Data2: &SavedContact{
				Cmd:       Cmd_SavedPhoneContact,
				ClassName: ClassSavedPhoneContact,
			},
		}
	}
	v1 := &TLSavedPhoneContact{Data2: v}
	v1.Data2.ClassName = ClassSavedPhoneContact
	return v1
}

func NewTLSecureCredentialsEncrypted(v *SecureCredentialsEncrypted) *TLSecureCredentialsEncrypted {
	if v == nil {
		return &TLSecureCredentialsEncrypted{
			Data2: &SecureCredentialsEncrypted{
				Cmd:       Cmd_SecureCredentialsEncrypted,
				ClassName: ClassSecureCredentialsEncrypted,
			},
		}
	}
	v1 := &TLSecureCredentialsEncrypted{Data2: v}
	v1.Data2.ClassName = ClassSecureCredentialsEncrypted
	return v1
}

func NewTLSecureData(v *SecureData) *TLSecureData {
	if v == nil {
		return &TLSecureData{
			Data2: &SecureData{
				Cmd:       Cmd_SecureData,
				ClassName: ClassSecureData,
			},
		}
	}
	v1 := &TLSecureData{Data2: v}
	v1.Data2.ClassName = ClassSecureData
	return v1
}

func NewTLSecureFileEmpty(v *SecureFile) *TLSecureFileEmpty {
	if v == nil {
		return &TLSecureFileEmpty{
			Data2: &SecureFile{
				Cmd:       Cmd_SecureFileEmpty,
				ClassName: ClassSecureFileEmpty,
			},
		}
	}
	v1 := &TLSecureFileEmpty{Data2: v}
	v1.Data2.ClassName = ClassSecureFileEmpty
	return v1
}

func NewTLSecureFile(v *SecureFile) *TLSecureFile {
	if v == nil {
		return &TLSecureFile{
			Data2: &SecureFile{
				Cmd:       Cmd_SecureFile,
				ClassName: ClassSecureFile,
			},
		}
	}
	v1 := &TLSecureFile{Data2: v}
	v1.Data2.ClassName = ClassSecureFile
	return v1
}

func NewTLSecurePasswordKdfAlgoUnknown(v *SecurePasswordKdfAlgo) *TLSecurePasswordKdfAlgoUnknown {
	if v == nil {
		return &TLSecurePasswordKdfAlgoUnknown{
			Data2: &SecurePasswordKdfAlgo{
				Cmd:       Cmd_SecurePasswordKdfAlgoUnknown,
				ClassName: ClassSecurePasswordKdfAlgoUnknown,
			},
		}
	}
	v1 := &TLSecurePasswordKdfAlgoUnknown{Data2: v}
	v1.Data2.ClassName = ClassSecurePasswordKdfAlgoUnknown
	return v1
}

func NewTLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000(v *SecurePasswordKdfAlgo) *TLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000 {
	if v == nil {
		return &TLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000{
			Data2: &SecurePasswordKdfAlgo{
				Cmd:       Cmd_SecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000,
				ClassName: ClassSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000,
			},
		}
	}
	v1 := &TLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000{Data2: v}
	v1.Data2.ClassName = ClassSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000
	return v1
}

func NewTLSecurePasswordKdfAlgoSHA512(v *SecurePasswordKdfAlgo) *TLSecurePasswordKdfAlgoSHA512 {
	if v == nil {
		return &TLSecurePasswordKdfAlgoSHA512{
			Data2: &SecurePasswordKdfAlgo{
				Cmd:       Cmd_SecurePasswordKdfAlgoSHA512,
				ClassName: ClassSecurePasswordKdfAlgoSHA512,
			},
		}
	}
	v1 := &TLSecurePasswordKdfAlgoSHA512{Data2: v}
	v1.Data2.ClassName = ClassSecurePasswordKdfAlgoSHA512
	return v1
}

func NewTLSecurePlainPhone(v *SecurePlainData) *TLSecurePlainPhone {
	if v == nil {
		return &TLSecurePlainPhone{
			Data2: &SecurePlainData{
				Cmd:       Cmd_SecurePlainPhone,
				ClassName: ClassSecurePlainPhone,
			},
		}
	}
	v1 := &TLSecurePlainPhone{Data2: v}
	v1.Data2.ClassName = ClassSecurePlainPhone
	return v1
}

func NewTLSecurePlainEmail(v *SecurePlainData) *TLSecurePlainEmail {
	if v == nil {
		return &TLSecurePlainEmail{
			Data2: &SecurePlainData{
				Cmd:       Cmd_SecurePlainEmail,
				ClassName: ClassSecurePlainEmail,
			},
		}
	}
	v1 := &TLSecurePlainEmail{Data2: v}
	v1.Data2.ClassName = ClassSecurePlainEmail
	return v1
}

func NewTLSecureRequiredType(v *SecureRequiredType) *TLSecureRequiredType {
	if v == nil {
		return &TLSecureRequiredType{
			Data2: &SecureRequiredType{
				Cmd:       Cmd_SecureRequiredType,
				ClassName: ClassSecureRequiredType,
			},
		}
	}
	v1 := &TLSecureRequiredType{Data2: v}
	v1.Data2.ClassName = ClassSecureRequiredType
	return v1
}

func NewTLSecureRequiredTypeOneOf(v *SecureRequiredType) *TLSecureRequiredTypeOneOf {
	if v == nil {
		return &TLSecureRequiredTypeOneOf{
			Data2: &SecureRequiredType{
				Cmd:       Cmd_SecureRequiredTypeOneOf,
				ClassName: ClassSecureRequiredTypeOneOf,
			},
		}
	}
	v1 := &TLSecureRequiredTypeOneOf{Data2: v}
	v1.Data2.ClassName = ClassSecureRequiredTypeOneOf
	return v1
}

func NewTLSecureSecretSettings(v *SecureSecretSettings) *TLSecureSecretSettings {
	if v == nil {
		return &TLSecureSecretSettings{
			Data2: &SecureSecretSettings{
				Cmd:       Cmd_SecureSecretSettings,
				ClassName: ClassSecureSecretSettings,
			},
		}
	}
	v1 := &TLSecureSecretSettings{Data2: v}
	v1.Data2.ClassName = ClassSecureSecretSettings
	return v1
}

func NewTLSecureValue(v *SecureValue) *TLSecureValue {
	if v == nil {
		return &TLSecureValue{
			Data2: &SecureValue{
				Cmd:       Cmd_SecureValue,
				ClassName: ClassSecureValue,
			},
		}
	}
	v1 := &TLSecureValue{Data2: v}
	v1.Data2.ClassName = ClassSecureValue
	return v1
}

func NewTLSecureValueErrorData(v *SecureValueError) *TLSecureValueErrorData {
	if v == nil {
		return &TLSecureValueErrorData{
			Data2: &SecureValueError{
				Cmd:       Cmd_SecureValueErrorData,
				ClassName: ClassSecureValueErrorData,
			},
		}
	}
	v1 := &TLSecureValueErrorData{Data2: v}
	v1.Data2.ClassName = ClassSecureValueErrorData
	return v1
}

func NewTLSecureValueErrorFrontSide(v *SecureValueError) *TLSecureValueErrorFrontSide {
	if v == nil {
		return &TLSecureValueErrorFrontSide{
			Data2: &SecureValueError{
				Cmd:       Cmd_SecureValueErrorFrontSide,
				ClassName: ClassSecureValueErrorFrontSide,
			},
		}
	}
	v1 := &TLSecureValueErrorFrontSide{Data2: v}
	v1.Data2.ClassName = ClassSecureValueErrorFrontSide
	return v1
}

func NewTLSecureValueErrorReverseSide(v *SecureValueError) *TLSecureValueErrorReverseSide {
	if v == nil {
		return &TLSecureValueErrorReverseSide{
			Data2: &SecureValueError{
				Cmd:       Cmd_SecureValueErrorReverseSide,
				ClassName: ClassSecureValueErrorReverseSide,
			},
		}
	}
	v1 := &TLSecureValueErrorReverseSide{Data2: v}
	v1.Data2.ClassName = ClassSecureValueErrorReverseSide
	return v1
}

func NewTLSecureValueErrorSelfie(v *SecureValueError) *TLSecureValueErrorSelfie {
	if v == nil {
		return &TLSecureValueErrorSelfie{
			Data2: &SecureValueError{
				Cmd:       Cmd_SecureValueErrorSelfie,
				ClassName: ClassSecureValueErrorSelfie,
			},
		}
	}
	v1 := &TLSecureValueErrorSelfie{Data2: v}
	v1.Data2.ClassName = ClassSecureValueErrorSelfie
	return v1
}

func NewTLSecureValueErrorFile(v *SecureValueError) *TLSecureValueErrorFile {
	if v == nil {
		return &TLSecureValueErrorFile{
			Data2: &SecureValueError{
				Cmd:       Cmd_SecureValueErrorFile,
				ClassName: ClassSecureValueErrorFile,
			},
		}
	}
	v1 := &TLSecureValueErrorFile{Data2: v}
	v1.Data2.ClassName = ClassSecureValueErrorFile
	return v1
}

func NewTLSecureValueErrorFiles(v *SecureValueError) *TLSecureValueErrorFiles {
	if v == nil {
		return &TLSecureValueErrorFiles{
			Data2: &SecureValueError{
				Cmd:       Cmd_SecureValueErrorFiles,
				ClassName: ClassSecureValueErrorFiles,
			},
		}
	}
	v1 := &TLSecureValueErrorFiles{Data2: v}
	v1.Data2.ClassName = ClassSecureValueErrorFiles
	return v1
}

func NewTLSecureValueError(v *SecureValueError) *TLSecureValueError {
	if v == nil {
		return &TLSecureValueError{
			Data2: &SecureValueError{
				Cmd:       Cmd_SecureValueError,
				ClassName: ClassSecureValueError,
			},
		}
	}
	v1 := &TLSecureValueError{Data2: v}
	v1.Data2.ClassName = ClassSecureValueError
	return v1
}

func NewTLSecureValueErrorTranslationFile(v *SecureValueError) *TLSecureValueErrorTranslationFile {
	if v == nil {
		return &TLSecureValueErrorTranslationFile{
			Data2: &SecureValueError{
				Cmd:       Cmd_SecureValueErrorTranslationFile,
				ClassName: ClassSecureValueErrorTranslationFile,
			},
		}
	}
	v1 := &TLSecureValueErrorTranslationFile{Data2: v}
	v1.Data2.ClassName = ClassSecureValueErrorTranslationFile
	return v1
}

func NewTLSecureValueErrorTranslationFiles(v *SecureValueError) *TLSecureValueErrorTranslationFiles {
	if v == nil {
		return &TLSecureValueErrorTranslationFiles{
			Data2: &SecureValueError{
				Cmd:       Cmd_SecureValueErrorTranslationFiles,
				ClassName: ClassSecureValueErrorTranslationFiles,
			},
		}
	}
	v1 := &TLSecureValueErrorTranslationFiles{Data2: v}
	v1.Data2.ClassName = ClassSecureValueErrorTranslationFiles
	return v1
}

func NewTLSecureValueHash(v *SecureValueHash) *TLSecureValueHash {
	if v == nil {
		return &TLSecureValueHash{
			Data2: &SecureValueHash{
				Cmd:       Cmd_SecureValueHash,
				ClassName: ClassSecureValueHash,
			},
		}
	}
	v1 := &TLSecureValueHash{Data2: v}
	v1.Data2.ClassName = ClassSecureValueHash
	return v1
}

func NewTLSecureValueTypePersonalDetails(v *SecureValueType) *TLSecureValueTypePersonalDetails {
	if v == nil {
		return &TLSecureValueTypePersonalDetails{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypePersonalDetails,
				ClassName: ClassSecureValueTypePersonalDetails,
			},
		}
	}
	v1 := &TLSecureValueTypePersonalDetails{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypePersonalDetails
	return v1
}

func NewTLSecureValueTypePassport(v *SecureValueType) *TLSecureValueTypePassport {
	if v == nil {
		return &TLSecureValueTypePassport{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypePassport,
				ClassName: ClassSecureValueTypePassport,
			},
		}
	}
	v1 := &TLSecureValueTypePassport{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypePassport
	return v1
}

func NewTLSecureValueTypeDriverLicense(v *SecureValueType) *TLSecureValueTypeDriverLicense {
	if v == nil {
		return &TLSecureValueTypeDriverLicense{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypeDriverLicense,
				ClassName: ClassSecureValueTypeDriverLicense,
			},
		}
	}
	v1 := &TLSecureValueTypeDriverLicense{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypeDriverLicense
	return v1
}

func NewTLSecureValueTypeIdentityCard(v *SecureValueType) *TLSecureValueTypeIdentityCard {
	if v == nil {
		return &TLSecureValueTypeIdentityCard{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypeIdentityCard,
				ClassName: ClassSecureValueTypeIdentityCard,
			},
		}
	}
	v1 := &TLSecureValueTypeIdentityCard{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypeIdentityCard
	return v1
}

func NewTLSecureValueTypeInternalPassport(v *SecureValueType) *TLSecureValueTypeInternalPassport {
	if v == nil {
		return &TLSecureValueTypeInternalPassport{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypeInternalPassport,
				ClassName: ClassSecureValueTypeInternalPassport,
			},
		}
	}
	v1 := &TLSecureValueTypeInternalPassport{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypeInternalPassport
	return v1
}

func NewTLSecureValueTypeAddress(v *SecureValueType) *TLSecureValueTypeAddress {
	if v == nil {
		return &TLSecureValueTypeAddress{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypeAddress,
				ClassName: ClassSecureValueTypeAddress,
			},
		}
	}
	v1 := &TLSecureValueTypeAddress{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypeAddress
	return v1
}

func NewTLSecureValueTypeUtilityBill(v *SecureValueType) *TLSecureValueTypeUtilityBill {
	if v == nil {
		return &TLSecureValueTypeUtilityBill{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypeUtilityBill,
				ClassName: ClassSecureValueTypeUtilityBill,
			},
		}
	}
	v1 := &TLSecureValueTypeUtilityBill{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypeUtilityBill
	return v1
}

func NewTLSecureValueTypeBankStatement(v *SecureValueType) *TLSecureValueTypeBankStatement {
	if v == nil {
		return &TLSecureValueTypeBankStatement{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypeBankStatement,
				ClassName: ClassSecureValueTypeBankStatement,
			},
		}
	}
	v1 := &TLSecureValueTypeBankStatement{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypeBankStatement
	return v1
}

func NewTLSecureValueTypeRentalAgreement(v *SecureValueType) *TLSecureValueTypeRentalAgreement {
	if v == nil {
		return &TLSecureValueTypeRentalAgreement{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypeRentalAgreement,
				ClassName: ClassSecureValueTypeRentalAgreement,
			},
		}
	}
	v1 := &TLSecureValueTypeRentalAgreement{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypeRentalAgreement
	return v1
}

func NewTLSecureValueTypePassportRegistration(v *SecureValueType) *TLSecureValueTypePassportRegistration {
	if v == nil {
		return &TLSecureValueTypePassportRegistration{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypePassportRegistration,
				ClassName: ClassSecureValueTypePassportRegistration,
			},
		}
	}
	v1 := &TLSecureValueTypePassportRegistration{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypePassportRegistration
	return v1
}

func NewTLSecureValueTypeTemporaryRegistration(v *SecureValueType) *TLSecureValueTypeTemporaryRegistration {
	if v == nil {
		return &TLSecureValueTypeTemporaryRegistration{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypeTemporaryRegistration,
				ClassName: ClassSecureValueTypeTemporaryRegistration,
			},
		}
	}
	v1 := &TLSecureValueTypeTemporaryRegistration{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypeTemporaryRegistration
	return v1
}

func NewTLSecureValueTypePhone(v *SecureValueType) *TLSecureValueTypePhone {
	if v == nil {
		return &TLSecureValueTypePhone{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypePhone,
				ClassName: ClassSecureValueTypePhone,
			},
		}
	}
	v1 := &TLSecureValueTypePhone{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypePhone
	return v1
}

func NewTLSecureValueTypeEmail(v *SecureValueType) *TLSecureValueTypeEmail {
	if v == nil {
		return &TLSecureValueTypeEmail{
			Data2: &SecureValueType{
				Cmd:       Cmd_SecureValueTypeEmail,
				ClassName: ClassSecureValueTypeEmail,
			},
		}
	}
	v1 := &TLSecureValueTypeEmail{Data2: v}
	v1.Data2.ClassName = ClassSecureValueTypeEmail
	return v1
}

func NewTLSendMessageTypingAction(v *SendMessageAction) *TLSendMessageTypingAction {
	if v == nil {
		return &TLSendMessageTypingAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageTypingAction,
				ClassName: ClassSendMessageTypingAction,
			},
		}
	}
	v1 := &TLSendMessageTypingAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageTypingAction
	return v1
}

func NewTLSendMessageCancelAction(v *SendMessageAction) *TLSendMessageCancelAction {
	if v == nil {
		return &TLSendMessageCancelAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageCancelAction,
				ClassName: ClassSendMessageCancelAction,
			},
		}
	}
	v1 := &TLSendMessageCancelAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageCancelAction
	return v1
}

func NewTLSendMessageRecordVideoAction(v *SendMessageAction) *TLSendMessageRecordVideoAction {
	if v == nil {
		return &TLSendMessageRecordVideoAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageRecordVideoAction,
				ClassName: ClassSendMessageRecordVideoAction,
			},
		}
	}
	v1 := &TLSendMessageRecordVideoAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageRecordVideoAction
	return v1
}

func NewTLSendMessageUploadVideoAction(v *SendMessageAction) *TLSendMessageUploadVideoAction {
	if v == nil {
		return &TLSendMessageUploadVideoAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageUploadVideoAction,
				ClassName: ClassSendMessageUploadVideoAction,
			},
		}
	}
	v1 := &TLSendMessageUploadVideoAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageUploadVideoAction
	return v1
}

func NewTLSendMessageRecordAudioAction(v *SendMessageAction) *TLSendMessageRecordAudioAction {
	if v == nil {
		return &TLSendMessageRecordAudioAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageRecordAudioAction,
				ClassName: ClassSendMessageRecordAudioAction,
			},
		}
	}
	v1 := &TLSendMessageRecordAudioAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageRecordAudioAction
	return v1
}

func NewTLSendMessageUploadAudioAction(v *SendMessageAction) *TLSendMessageUploadAudioAction {
	if v == nil {
		return &TLSendMessageUploadAudioAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageUploadAudioAction,
				ClassName: ClassSendMessageUploadAudioAction,
			},
		}
	}
	v1 := &TLSendMessageUploadAudioAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageUploadAudioAction
	return v1
}

func NewTLSendMessageUploadPhotoAction(v *SendMessageAction) *TLSendMessageUploadPhotoAction {
	if v == nil {
		return &TLSendMessageUploadPhotoAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageUploadPhotoAction,
				ClassName: ClassSendMessageUploadPhotoAction,
			},
		}
	}
	v1 := &TLSendMessageUploadPhotoAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageUploadPhotoAction
	return v1
}

func NewTLSendMessageUploadDocumentAction(v *SendMessageAction) *TLSendMessageUploadDocumentAction {
	if v == nil {
		return &TLSendMessageUploadDocumentAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageUploadDocumentAction,
				ClassName: ClassSendMessageUploadDocumentAction,
			},
		}
	}
	v1 := &TLSendMessageUploadDocumentAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageUploadDocumentAction
	return v1
}

func NewTLSendMessageGeoLocationAction(v *SendMessageAction) *TLSendMessageGeoLocationAction {
	if v == nil {
		return &TLSendMessageGeoLocationAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageGeoLocationAction,
				ClassName: ClassSendMessageGeoLocationAction,
			},
		}
	}
	v1 := &TLSendMessageGeoLocationAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageGeoLocationAction
	return v1
}

func NewTLSendMessageChooseContactAction(v *SendMessageAction) *TLSendMessageChooseContactAction {
	if v == nil {
		return &TLSendMessageChooseContactAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageChooseContactAction,
				ClassName: ClassSendMessageChooseContactAction,
			},
		}
	}
	v1 := &TLSendMessageChooseContactAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageChooseContactAction
	return v1
}

func NewTLSendMessageGamePlayAction(v *SendMessageAction) *TLSendMessageGamePlayAction {
	if v == nil {
		return &TLSendMessageGamePlayAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageGamePlayAction,
				ClassName: ClassSendMessageGamePlayAction,
			},
		}
	}
	v1 := &TLSendMessageGamePlayAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageGamePlayAction
	return v1
}

func NewTLSendMessageRecordRoundAction(v *SendMessageAction) *TLSendMessageRecordRoundAction {
	if v == nil {
		return &TLSendMessageRecordRoundAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageRecordRoundAction,
				ClassName: ClassSendMessageRecordRoundAction,
			},
		}
	}
	v1 := &TLSendMessageRecordRoundAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageRecordRoundAction
	return v1
}

func NewTLSendMessageUploadRoundAction(v *SendMessageAction) *TLSendMessageUploadRoundAction {
	if v == nil {
		return &TLSendMessageUploadRoundAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SendMessageUploadRoundAction,
				ClassName: ClassSendMessageUploadRoundAction,
			},
		}
	}
	v1 := &TLSendMessageUploadRoundAction{Data2: v}
	v1.Data2.ClassName = ClassSendMessageUploadRoundAction
	return v1
}

func NewTLSpeakingInGroupCallAction(v *SendMessageAction) *TLSpeakingInGroupCallAction {
	if v == nil {
		return &TLSpeakingInGroupCallAction{
			Data2: &SendMessageAction{
				Cmd:       Cmd_SpeakingInGroupCallAction,
				ClassName: ClassSpeakingInGroupCallAction,
			},
		}
	}
	v1 := &TLSpeakingInGroupCallAction{Data2: v}
	v1.Data2.ClassName = ClassSpeakingInGroupCallAction
	return v1
}

func NewTLServer_DHInnerData(v *Server_DHInnerData) *TLServer_DHInnerData {
	if v == nil {
		return &TLServer_DHInnerData{
			Data2: &Server_DHInnerData{
				Cmd:       Cmd_Server_DHInnerData,
				ClassName: ClassServer_DHInnerData,
			},
		}
	}
	v1 := &TLServer_DHInnerData{Data2: v}
	v1.Data2.ClassName = ClassServer_DHInnerData
	return v1
}

func NewTLServer_DHParamsFail(v *Server_DH_Params) *TLServer_DHParamsFail {
	if v == nil {
		return &TLServer_DHParamsFail{
			Data2: &Server_DH_Params{
				Cmd:       Cmd_Server_DHParamsFail,
				ClassName: ClassServer_DHParamsFail,
			},
		}
	}
	v1 := &TLServer_DHParamsFail{Data2: v}
	v1.Data2.ClassName = ClassServer_DHParamsFail
	return v1
}

func NewTLServer_DHParamsOk(v *Server_DH_Params) *TLServer_DHParamsOk {
	if v == nil {
		return &TLServer_DHParamsOk{
			Data2: &Server_DH_Params{
				Cmd:       Cmd_Server_DHParamsOk,
				ClassName: ClassServer_DHParamsOk,
			},
		}
	}
	v1 := &TLServer_DHParamsOk{Data2: v}
	v1.Data2.ClassName = ClassServer_DHParamsOk
	return v1
}

func NewTLDhGenOk(v *SetClient_DHParamsAnswer) *TLDhGenOk {
	if v == nil {
		return &TLDhGenOk{
			Data2: &SetClient_DHParamsAnswer{
				Cmd:       Cmd_DhGenOk,
				ClassName: ClassDhGenOk,
			},
		}
	}
	v1 := &TLDhGenOk{Data2: v}
	v1.Data2.ClassName = ClassDhGenOk
	return v1
}

func NewTLDhGenRetry(v *SetClient_DHParamsAnswer) *TLDhGenRetry {
	if v == nil {
		return &TLDhGenRetry{
			Data2: &SetClient_DHParamsAnswer{
				Cmd:       Cmd_DhGenRetry,
				ClassName: ClassDhGenRetry,
			},
		}
	}
	v1 := &TLDhGenRetry{Data2: v}
	v1.Data2.ClassName = ClassDhGenRetry
	return v1
}

func NewTLDhGenFail(v *SetClient_DHParamsAnswer) *TLDhGenFail {
	if v == nil {
		return &TLDhGenFail{
			Data2: &SetClient_DHParamsAnswer{
				Cmd:       Cmd_DhGenFail,
				ClassName: ClassDhGenFail,
			},
		}
	}
	v1 := &TLDhGenFail{Data2: v}
	v1.Data2.ClassName = ClassDhGenFail
	return v1
}

func NewTLShippingOption(v *ShippingOption) *TLShippingOption {
	if v == nil {
		return &TLShippingOption{
			Data2: &ShippingOption{
				Cmd:       Cmd_ShippingOption,
				ClassName: ClassShippingOption,
			},
		}
	}
	v1 := &TLShippingOption{Data2: v}
	v1.Data2.ClassName = ClassShippingOption
	return v1
}

func NewTLStatsAbsValueAndPrev(v *StatsAbsValueAndPrev) *TLStatsAbsValueAndPrev {
	if v == nil {
		return &TLStatsAbsValueAndPrev{
			Data2: &StatsAbsValueAndPrev{
				Cmd:       Cmd_StatsAbsValueAndPrev,
				ClassName: ClassStatsAbsValueAndPrev,
			},
		}
	}
	v1 := &TLStatsAbsValueAndPrev{Data2: v}
	v1.Data2.ClassName = ClassStatsAbsValueAndPrev
	return v1
}

func NewTLStatsDateRangeDays(v *StatsDateRangeDays) *TLStatsDateRangeDays {
	if v == nil {
		return &TLStatsDateRangeDays{
			Data2: &StatsDateRangeDays{
				Cmd:       Cmd_StatsDateRangeDays,
				ClassName: ClassStatsDateRangeDays,
			},
		}
	}
	v1 := &TLStatsDateRangeDays{Data2: v}
	v1.Data2.ClassName = ClassStatsDateRangeDays
	return v1
}

func NewTLStatsGraphAsync(v *StatsGraph) *TLStatsGraphAsync {
	if v == nil {
		return &TLStatsGraphAsync{
			Data2: &StatsGraph{
				Cmd:       Cmd_StatsGraphAsync,
				ClassName: ClassStatsGraphAsync,
			},
		}
	}
	v1 := &TLStatsGraphAsync{Data2: v}
	v1.Data2.ClassName = ClassStatsGraphAsync
	return v1
}

func NewTLStatsGraphError(v *StatsGraph) *TLStatsGraphError {
	if v == nil {
		return &TLStatsGraphError{
			Data2: &StatsGraph{
				Cmd:       Cmd_StatsGraphError,
				ClassName: ClassStatsGraphError,
			},
		}
	}
	v1 := &TLStatsGraphError{Data2: v}
	v1.Data2.ClassName = ClassStatsGraphError
	return v1
}

func NewTLStatsGraph(v *StatsGraph) *TLStatsGraph {
	if v == nil {
		return &TLStatsGraph{
			Data2: &StatsGraph{
				Cmd:       Cmd_StatsGraph,
				ClassName: ClassStatsGraph,
			},
		}
	}
	v1 := &TLStatsGraph{Data2: v}
	v1.Data2.ClassName = ClassStatsGraph
	return v1
}

func NewTLStatsGroupTopAdmin(v *StatsGroupTopAdmin) *TLStatsGroupTopAdmin {
	if v == nil {
		return &TLStatsGroupTopAdmin{
			Data2: &StatsGroupTopAdmin{
				Cmd:       Cmd_StatsGroupTopAdmin,
				ClassName: ClassStatsGroupTopAdmin,
			},
		}
	}
	v1 := &TLStatsGroupTopAdmin{Data2: v}
	v1.Data2.ClassName = ClassStatsGroupTopAdmin
	return v1
}

func NewTLStatsGroupTopInviter(v *StatsGroupTopInviter) *TLStatsGroupTopInviter {
	if v == nil {
		return &TLStatsGroupTopInviter{
			Data2: &StatsGroupTopInviter{
				Cmd:       Cmd_StatsGroupTopInviter,
				ClassName: ClassStatsGroupTopInviter,
			},
		}
	}
	v1 := &TLStatsGroupTopInviter{Data2: v}
	v1.Data2.ClassName = ClassStatsGroupTopInviter
	return v1
}

func NewTLStatsGroupTopPoster(v *StatsGroupTopPoster) *TLStatsGroupTopPoster {
	if v == nil {
		return &TLStatsGroupTopPoster{
			Data2: &StatsGroupTopPoster{
				Cmd:       Cmd_StatsGroupTopPoster,
				ClassName: ClassStatsGroupTopPoster,
			},
		}
	}
	v1 := &TLStatsGroupTopPoster{Data2: v}
	v1.Data2.ClassName = ClassStatsGroupTopPoster
	return v1
}

func NewTLStatsPercentValue(v *StatsPercentValue) *TLStatsPercentValue {
	if v == nil {
		return &TLStatsPercentValue{
			Data2: &StatsPercentValue{
				Cmd:       Cmd_StatsPercentValue,
				ClassName: ClassStatsPercentValue,
			},
		}
	}
	v1 := &TLStatsPercentValue{Data2: v}
	v1.Data2.ClassName = ClassStatsPercentValue
	return v1
}

func NewTLStatsURL(v *StatsURL) *TLStatsURL {
	if v == nil {
		return &TLStatsURL{
			Data2: &StatsURL{
				Cmd:       Cmd_StatsURL,
				ClassName: ClassStatsURL,
			},
		}
	}
	v1 := &TLStatsURL{Data2: v}
	v1.Data2.ClassName = ClassStatsURL
	return v1
}

func NewTLStatsBroadcastStats(v *Stats_BroadcastStats) *TLStatsBroadcastStats {
	if v == nil {
		return &TLStatsBroadcastStats{
			Data2: &Stats_BroadcastStats{
				Cmd:       Cmd_StatsBroadcastStats,
				ClassName: ClassStatsBroadcastStats,
			},
		}
	}
	v1 := &TLStatsBroadcastStats{Data2: v}
	v1.Data2.ClassName = ClassStatsBroadcastStats
	return v1
}

func NewTLStatsMegagroupStats(v *Stats_MegagroupStats) *TLStatsMegagroupStats {
	if v == nil {
		return &TLStatsMegagroupStats{
			Data2: &Stats_MegagroupStats{
				Cmd:       Cmd_StatsMegagroupStats,
				ClassName: ClassStatsMegagroupStats,
			},
		}
	}
	v1 := &TLStatsMegagroupStats{Data2: v}
	v1.Data2.ClassName = ClassStatsMegagroupStats
	return v1
}

func NewTLStatsMessageStats(v *Stats_MessageStats) *TLStatsMessageStats {
	if v == nil {
		return &TLStatsMessageStats{
			Data2: &Stats_MessageStats{
				Cmd:       Cmd_StatsMessageStats,
				ClassName: ClassStatsMessageStats,
			},
		}
	}
	v1 := &TLStatsMessageStats{Data2: v}
	v1.Data2.ClassName = ClassStatsMessageStats
	return v1
}

func NewTLStickerPack(v *StickerPack) *TLStickerPack {
	if v == nil {
		return &TLStickerPack{
			Data2: &StickerPack{
				Cmd:       Cmd_StickerPack,
				ClassName: ClassStickerPack,
			},
		}
	}
	v1 := &TLStickerPack{Data2: v}
	v1.Data2.ClassName = ClassStickerPack
	return v1
}

func NewTLStickerSet(v *StickerSet) *TLStickerSet {
	if v == nil {
		return &TLStickerSet{
			Data2: &StickerSet{
				Cmd:       Cmd_StickerSet,
				ClassName: ClassStickerSet,
			},
		}
	}
	v1 := &TLStickerSet{Data2: v}
	v1.Data2.ClassName = ClassStickerSet
	return v1
}

func NewTLStickerSetCovered(v *StickerSetCovered) *TLStickerSetCovered {
	if v == nil {
		return &TLStickerSetCovered{
			Data2: &StickerSetCovered{
				Cmd:       Cmd_StickerSetCovered,
				ClassName: ClassStickerSetCovered,
			},
		}
	}
	v1 := &TLStickerSetCovered{Data2: v}
	v1.Data2.ClassName = ClassStickerSetCovered
	return v1
}

func NewTLStickerSetMultiCovered(v *StickerSetCovered) *TLStickerSetMultiCovered {
	if v == nil {
		return &TLStickerSetMultiCovered{
			Data2: &StickerSetCovered{
				Cmd:       Cmd_StickerSetMultiCovered,
				ClassName: ClassStickerSetMultiCovered,
			},
		}
	}
	v1 := &TLStickerSetMultiCovered{Data2: v}
	v1.Data2.ClassName = ClassStickerSetMultiCovered
	return v1
}

func NewTLStorageFileUnknown(v *Storage_FileType) *TLStorageFileUnknown {
	if v == nil {
		return &TLStorageFileUnknown{
			Data2: &Storage_FileType{
				Cmd:       Cmd_StorageFileUnknown,
				ClassName: ClassStorageFileUnknown,
			},
		}
	}
	v1 := &TLStorageFileUnknown{Data2: v}
	v1.Data2.ClassName = ClassStorageFileUnknown
	return v1
}

func NewTLStorageFilePartial(v *Storage_FileType) *TLStorageFilePartial {
	if v == nil {
		return &TLStorageFilePartial{
			Data2: &Storage_FileType{
				Cmd:       Cmd_StorageFilePartial,
				ClassName: ClassStorageFilePartial,
			},
		}
	}
	v1 := &TLStorageFilePartial{Data2: v}
	v1.Data2.ClassName = ClassStorageFilePartial
	return v1
}

func NewTLStorageFileJpeg(v *Storage_FileType) *TLStorageFileJpeg {
	if v == nil {
		return &TLStorageFileJpeg{
			Data2: &Storage_FileType{
				Cmd:       Cmd_StorageFileJpeg,
				ClassName: ClassStorageFileJpeg,
			},
		}
	}
	v1 := &TLStorageFileJpeg{Data2: v}
	v1.Data2.ClassName = ClassStorageFileJpeg
	return v1
}

func NewTLStorageFileGif(v *Storage_FileType) *TLStorageFileGif {
	if v == nil {
		return &TLStorageFileGif{
			Data2: &Storage_FileType{
				Cmd:       Cmd_StorageFileGif,
				ClassName: ClassStorageFileGif,
			},
		}
	}
	v1 := &TLStorageFileGif{Data2: v}
	v1.Data2.ClassName = ClassStorageFileGif
	return v1
}

func NewTLStorageFilePng(v *Storage_FileType) *TLStorageFilePng {
	if v == nil {
		return &TLStorageFilePng{
			Data2: &Storage_FileType{
				Cmd:       Cmd_StorageFilePng,
				ClassName: ClassStorageFilePng,
			},
		}
	}
	v1 := &TLStorageFilePng{Data2: v}
	v1.Data2.ClassName = ClassStorageFilePng
	return v1
}

func NewTLStorageFilePdf(v *Storage_FileType) *TLStorageFilePdf {
	if v == nil {
		return &TLStorageFilePdf{
			Data2: &Storage_FileType{
				Cmd:       Cmd_StorageFilePdf,
				ClassName: ClassStorageFilePdf,
			},
		}
	}
	v1 := &TLStorageFilePdf{Data2: v}
	v1.Data2.ClassName = ClassStorageFilePdf
	return v1
}

func NewTLStorageFileMp3(v *Storage_FileType) *TLStorageFileMp3 {
	if v == nil {
		return &TLStorageFileMp3{
			Data2: &Storage_FileType{
				Cmd:       Cmd_StorageFileMp3,
				ClassName: ClassStorageFileMp3,
			},
		}
	}
	v1 := &TLStorageFileMp3{Data2: v}
	v1.Data2.ClassName = ClassStorageFileMp3
	return v1
}

func NewTLStorageFileMov(v *Storage_FileType) *TLStorageFileMov {
	if v == nil {
		return &TLStorageFileMov{
			Data2: &Storage_FileType{
				Cmd:       Cmd_StorageFileMov,
				ClassName: ClassStorageFileMov,
			},
		}
	}
	v1 := &TLStorageFileMov{Data2: v}
	v1.Data2.ClassName = ClassStorageFileMov
	return v1
}

func NewTLStorageFileMp4(v *Storage_FileType) *TLStorageFileMp4 {
	if v == nil {
		return &TLStorageFileMp4{
			Data2: &Storage_FileType{
				Cmd:       Cmd_StorageFileMp4,
				ClassName: ClassStorageFileMp4,
			},
		}
	}
	v1 := &TLStorageFileMp4{Data2: v}
	v1.Data2.ClassName = ClassStorageFileMp4
	return v1
}

func NewTLStorageFileWebp(v *Storage_FileType) *TLStorageFileWebp {
	if v == nil {
		return &TLStorageFileWebp{
			Data2: &Storage_FileType{
				Cmd:       Cmd_StorageFileWebp,
				ClassName: ClassStorageFileWebp,
			},
		}
	}
	v1 := &TLStorageFileWebp{Data2: v}
	v1.Data2.ClassName = ClassStorageFileWebp
	return v1
}

func NewTLThemeDocumentNotModified(v *Theme) *TLThemeDocumentNotModified {
	if v == nil {
		return &TLThemeDocumentNotModified{
			Data2: &Theme{
				Cmd:       Cmd_ThemeDocumentNotModified,
				ClassName: ClassThemeDocumentNotModified,
			},
		}
	}
	v1 := &TLThemeDocumentNotModified{Data2: v}
	v1.Data2.ClassName = ClassThemeDocumentNotModified
	return v1
}

func NewTLTheme(v *Theme) *TLTheme {
	if v == nil {
		return &TLTheme{
			Data2: &Theme{
				Cmd:       Cmd_Theme,
				ClassName: ClassTheme,
			},
		}
	}
	v1 := &TLTheme{Data2: v}
	v1.Data2.ClassName = ClassTheme
	return v1
}

func NewTLThemeSettings(v *ThemeSettings) *TLThemeSettings {
	if v == nil {
		return &TLThemeSettings{
			Data2: &ThemeSettings{
				Cmd:       Cmd_ThemeSettings,
				ClassName: ClassThemeSettings,
			},
		}
	}
	v1 := &TLThemeSettings{Data2: v}
	v1.Data2.ClassName = ClassThemeSettings
	return v1
}

func NewTLTopPeer(v *TopPeer) *TLTopPeer {
	if v == nil {
		return &TLTopPeer{
			Data2: &TopPeer{
				Cmd:       Cmd_TopPeer,
				ClassName: ClassTopPeer,
			},
		}
	}
	v1 := &TLTopPeer{Data2: v}
	v1.Data2.ClassName = ClassTopPeer
	return v1
}

func NewTLTopPeerCategoryBotsPM(v *TopPeerCategory) *TLTopPeerCategoryBotsPM {
	if v == nil {
		return &TLTopPeerCategoryBotsPM{
			Data2: &TopPeerCategory{
				Cmd:       Cmd_TopPeerCategoryBotsPM,
				ClassName: ClassTopPeerCategoryBotsPM,
			},
		}
	}
	v1 := &TLTopPeerCategoryBotsPM{Data2: v}
	v1.Data2.ClassName = ClassTopPeerCategoryBotsPM
	return v1
}

func NewTLTopPeerCategoryBotsInline(v *TopPeerCategory) *TLTopPeerCategoryBotsInline {
	if v == nil {
		return &TLTopPeerCategoryBotsInline{
			Data2: &TopPeerCategory{
				Cmd:       Cmd_TopPeerCategoryBotsInline,
				ClassName: ClassTopPeerCategoryBotsInline,
			},
		}
	}
	v1 := &TLTopPeerCategoryBotsInline{Data2: v}
	v1.Data2.ClassName = ClassTopPeerCategoryBotsInline
	return v1
}

func NewTLTopPeerCategoryCorrespondents(v *TopPeerCategory) *TLTopPeerCategoryCorrespondents {
	if v == nil {
		return &TLTopPeerCategoryCorrespondents{
			Data2: &TopPeerCategory{
				Cmd:       Cmd_TopPeerCategoryCorrespondents,
				ClassName: ClassTopPeerCategoryCorrespondents,
			},
		}
	}
	v1 := &TLTopPeerCategoryCorrespondents{Data2: v}
	v1.Data2.ClassName = ClassTopPeerCategoryCorrespondents
	return v1
}

func NewTLTopPeerCategoryGroups(v *TopPeerCategory) *TLTopPeerCategoryGroups {
	if v == nil {
		return &TLTopPeerCategoryGroups{
			Data2: &TopPeerCategory{
				Cmd:       Cmd_TopPeerCategoryGroups,
				ClassName: ClassTopPeerCategoryGroups,
			},
		}
	}
	v1 := &TLTopPeerCategoryGroups{Data2: v}
	v1.Data2.ClassName = ClassTopPeerCategoryGroups
	return v1
}

func NewTLTopPeerCategoryChannels(v *TopPeerCategory) *TLTopPeerCategoryChannels {
	if v == nil {
		return &TLTopPeerCategoryChannels{
			Data2: &TopPeerCategory{
				Cmd:       Cmd_TopPeerCategoryChannels,
				ClassName: ClassTopPeerCategoryChannels,
			},
		}
	}
	v1 := &TLTopPeerCategoryChannels{Data2: v}
	v1.Data2.ClassName = ClassTopPeerCategoryChannels
	return v1
}

func NewTLTopPeerCategoryPhoneCalls(v *TopPeerCategory) *TLTopPeerCategoryPhoneCalls {
	if v == nil {
		return &TLTopPeerCategoryPhoneCalls{
			Data2: &TopPeerCategory{
				Cmd:       Cmd_TopPeerCategoryPhoneCalls,
				ClassName: ClassTopPeerCategoryPhoneCalls,
			},
		}
	}
	v1 := &TLTopPeerCategoryPhoneCalls{Data2: v}
	v1.Data2.ClassName = ClassTopPeerCategoryPhoneCalls
	return v1
}

func NewTLTopPeerCategoryForwardUsers(v *TopPeerCategory) *TLTopPeerCategoryForwardUsers {
	if v == nil {
		return &TLTopPeerCategoryForwardUsers{
			Data2: &TopPeerCategory{
				Cmd:       Cmd_TopPeerCategoryForwardUsers,
				ClassName: ClassTopPeerCategoryForwardUsers,
			},
		}
	}
	v1 := &TLTopPeerCategoryForwardUsers{Data2: v}
	v1.Data2.ClassName = ClassTopPeerCategoryForwardUsers
	return v1
}

func NewTLTopPeerCategoryForwardChats(v *TopPeerCategory) *TLTopPeerCategoryForwardChats {
	if v == nil {
		return &TLTopPeerCategoryForwardChats{
			Data2: &TopPeerCategory{
				Cmd:       Cmd_TopPeerCategoryForwardChats,
				ClassName: ClassTopPeerCategoryForwardChats,
			},
		}
	}
	v1 := &TLTopPeerCategoryForwardChats{Data2: v}
	v1.Data2.ClassName = ClassTopPeerCategoryForwardChats
	return v1
}

func NewTLTopPeerCategoryPeers(v *TopPeerCategoryPeers) *TLTopPeerCategoryPeers {
	if v == nil {
		return &TLTopPeerCategoryPeers{
			Data2: &TopPeerCategoryPeers{
				Cmd:       Cmd_TopPeerCategoryPeers,
				ClassName: ClassTopPeerCategoryPeers,
			},
		}
	}
	v1 := &TLTopPeerCategoryPeers{Data2: v}
	v1.Data2.ClassName = ClassTopPeerCategoryPeers
	return v1
}

func NewTLTrue(v *True) *TLTrue {
	if v == nil {
		return &TLTrue{
			Data2: &True{
				Cmd:       Cmd_True,
				ClassName: ClassTrue,
			},
		}
	}
	v1 := &TLTrue{Data2: v}
	v1.Data2.ClassName = ClassTrue
	return v1
}

func NewTLUpdateNewMessage(v *Update) *TLUpdateNewMessage {
	if v == nil {
		return &TLUpdateNewMessage{
			Data2: &Update{
				Cmd:       Cmd_UpdateNewMessage,
				ClassName: ClassUpdateNewMessage,
			},
		}
	}
	v1 := &TLUpdateNewMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateNewMessage
	return v1
}

func NewTLUpdateMessageID(v *Update) *TLUpdateMessageID {
	if v == nil {
		return &TLUpdateMessageID{
			Data2: &Update{
				Cmd:       Cmd_UpdateMessageID,
				ClassName: ClassUpdateMessageID,
			},
		}
	}
	v1 := &TLUpdateMessageID{Data2: v}
	v1.Data2.ClassName = ClassUpdateMessageID
	return v1
}

func NewTLUpdateDeleteMessages(v *Update) *TLUpdateDeleteMessages {
	if v == nil {
		return &TLUpdateDeleteMessages{
			Data2: &Update{
				Cmd:       Cmd_UpdateDeleteMessages,
				ClassName: ClassUpdateDeleteMessages,
			},
		}
	}
	v1 := &TLUpdateDeleteMessages{Data2: v}
	v1.Data2.ClassName = ClassUpdateDeleteMessages
	return v1
}

func NewTLUpdateUserTyping(v *Update) *TLUpdateUserTyping {
	if v == nil {
		return &TLUpdateUserTyping{
			Data2: &Update{
				Cmd:       Cmd_UpdateUserTyping,
				ClassName: ClassUpdateUserTyping,
			},
		}
	}
	v1 := &TLUpdateUserTyping{Data2: v}
	v1.Data2.ClassName = ClassUpdateUserTyping
	return v1
}

func NewTLUpdateChatUserTyping(v *Update) *TLUpdateChatUserTyping {
	if v == nil {
		return &TLUpdateChatUserTyping{
			Data2: &Update{
				Cmd:       Cmd_UpdateChatUserTyping,
				ClassName: ClassUpdateChatUserTyping,
			},
		}
	}
	v1 := &TLUpdateChatUserTyping{Data2: v}
	v1.Data2.ClassName = ClassUpdateChatUserTyping
	return v1
}

func NewTLUpdateChatParticipants(v *Update) *TLUpdateChatParticipants {
	if v == nil {
		return &TLUpdateChatParticipants{
			Data2: &Update{
				Cmd:       Cmd_UpdateChatParticipants,
				ClassName: ClassUpdateChatParticipants,
			},
		}
	}
	v1 := &TLUpdateChatParticipants{Data2: v}
	v1.Data2.ClassName = ClassUpdateChatParticipants
	return v1
}

func NewTLUpdateUserStatus(v *Update) *TLUpdateUserStatus {
	if v == nil {
		return &TLUpdateUserStatus{
			Data2: &Update{
				Cmd:       Cmd_UpdateUserStatus,
				ClassName: ClassUpdateUserStatus,
			},
		}
	}
	v1 := &TLUpdateUserStatus{Data2: v}
	v1.Data2.ClassName = ClassUpdateUserStatus
	return v1
}

func NewTLUpdateUserName(v *Update) *TLUpdateUserName {
	if v == nil {
		return &TLUpdateUserName{
			Data2: &Update{
				Cmd:       Cmd_UpdateUserName,
				ClassName: ClassUpdateUserName,
			},
		}
	}
	v1 := &TLUpdateUserName{Data2: v}
	v1.Data2.ClassName = ClassUpdateUserName
	return v1
}

func NewTLUpdateUserPhoto(v *Update) *TLUpdateUserPhoto {
	if v == nil {
		return &TLUpdateUserPhoto{
			Data2: &Update{
				Cmd:       Cmd_UpdateUserPhoto,
				ClassName: ClassUpdateUserPhoto,
			},
		}
	}
	v1 := &TLUpdateUserPhoto{Data2: v}
	v1.Data2.ClassName = ClassUpdateUserPhoto
	return v1
}

func NewTLUpdateContactRegistered(v *Update) *TLUpdateContactRegistered {
	if v == nil {
		return &TLUpdateContactRegistered{
			Data2: &Update{
				Cmd:       Cmd_UpdateContactRegistered,
				ClassName: ClassUpdateContactRegistered,
			},
		}
	}
	v1 := &TLUpdateContactRegistered{Data2: v}
	v1.Data2.ClassName = ClassUpdateContactRegistered
	return v1
}

func NewTLUpdateContactLink(v *Update) *TLUpdateContactLink {
	if v == nil {
		return &TLUpdateContactLink{
			Data2: &Update{
				Cmd:       Cmd_UpdateContactLink,
				ClassName: ClassUpdateContactLink,
			},
		}
	}
	v1 := &TLUpdateContactLink{Data2: v}
	v1.Data2.ClassName = ClassUpdateContactLink
	return v1
}

func NewTLUpdateNewEncryptedMessage(v *Update) *TLUpdateNewEncryptedMessage {
	if v == nil {
		return &TLUpdateNewEncryptedMessage{
			Data2: &Update{
				Cmd:       Cmd_UpdateNewEncryptedMessage,
				ClassName: ClassUpdateNewEncryptedMessage,
			},
		}
	}
	v1 := &TLUpdateNewEncryptedMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateNewEncryptedMessage
	return v1
}

func NewTLUpdateEncryptedChatTyping(v *Update) *TLUpdateEncryptedChatTyping {
	if v == nil {
		return &TLUpdateEncryptedChatTyping{
			Data2: &Update{
				Cmd:       Cmd_UpdateEncryptedChatTyping,
				ClassName: ClassUpdateEncryptedChatTyping,
			},
		}
	}
	v1 := &TLUpdateEncryptedChatTyping{Data2: v}
	v1.Data2.ClassName = ClassUpdateEncryptedChatTyping
	return v1
}

func NewTLUpdateEncryption(v *Update) *TLUpdateEncryption {
	if v == nil {
		return &TLUpdateEncryption{
			Data2: &Update{
				Cmd:       Cmd_UpdateEncryption,
				ClassName: ClassUpdateEncryption,
			},
		}
	}
	v1 := &TLUpdateEncryption{Data2: v}
	v1.Data2.ClassName = ClassUpdateEncryption
	return v1
}

func NewTLUpdateEncryptedMessagesRead(v *Update) *TLUpdateEncryptedMessagesRead {
	if v == nil {
		return &TLUpdateEncryptedMessagesRead{
			Data2: &Update{
				Cmd:       Cmd_UpdateEncryptedMessagesRead,
				ClassName: ClassUpdateEncryptedMessagesRead,
			},
		}
	}
	v1 := &TLUpdateEncryptedMessagesRead{Data2: v}
	v1.Data2.ClassName = ClassUpdateEncryptedMessagesRead
	return v1
}

func NewTLUpdateChatParticipantAdd(v *Update) *TLUpdateChatParticipantAdd {
	if v == nil {
		return &TLUpdateChatParticipantAdd{
			Data2: &Update{
				Cmd:       Cmd_UpdateChatParticipantAdd,
				ClassName: ClassUpdateChatParticipantAdd,
			},
		}
	}
	v1 := &TLUpdateChatParticipantAdd{Data2: v}
	v1.Data2.ClassName = ClassUpdateChatParticipantAdd
	return v1
}

func NewTLUpdateChatParticipantDelete(v *Update) *TLUpdateChatParticipantDelete {
	if v == nil {
		return &TLUpdateChatParticipantDelete{
			Data2: &Update{
				Cmd:       Cmd_UpdateChatParticipantDelete,
				ClassName: ClassUpdateChatParticipantDelete,
			},
		}
	}
	v1 := &TLUpdateChatParticipantDelete{Data2: v}
	v1.Data2.ClassName = ClassUpdateChatParticipantDelete
	return v1
}

func NewTLUpdateDcOptions(v *Update) *TLUpdateDcOptions {
	if v == nil {
		return &TLUpdateDcOptions{
			Data2: &Update{
				Cmd:       Cmd_UpdateDcOptions,
				ClassName: ClassUpdateDcOptions,
			},
		}
	}
	v1 := &TLUpdateDcOptions{Data2: v}
	v1.Data2.ClassName = ClassUpdateDcOptions
	return v1
}

func NewTLUpdateUserBlocked(v *Update) *TLUpdateUserBlocked {
	if v == nil {
		return &TLUpdateUserBlocked{
			Data2: &Update{
				Cmd:       Cmd_UpdateUserBlocked,
				ClassName: ClassUpdateUserBlocked,
			},
		}
	}
	v1 := &TLUpdateUserBlocked{Data2: v}
	v1.Data2.ClassName = ClassUpdateUserBlocked
	return v1
}

func NewTLUpdateNotifySettings(v *Update) *TLUpdateNotifySettings {
	if v == nil {
		return &TLUpdateNotifySettings{
			Data2: &Update{
				Cmd:       Cmd_UpdateNotifySettings,
				ClassName: ClassUpdateNotifySettings,
			},
		}
	}
	v1 := &TLUpdateNotifySettings{Data2: v}
	v1.Data2.ClassName = ClassUpdateNotifySettings
	return v1
}

func NewTLUpdateServiceNotification(v *Update) *TLUpdateServiceNotification {
	if v == nil {
		return &TLUpdateServiceNotification{
			Data2: &Update{
				Cmd:       Cmd_UpdateServiceNotification,
				ClassName: ClassUpdateServiceNotification,
			},
		}
	}
	v1 := &TLUpdateServiceNotification{Data2: v}
	v1.Data2.ClassName = ClassUpdateServiceNotification
	return v1
}

func NewTLUpdatePrivacy(v *Update) *TLUpdatePrivacy {
	if v == nil {
		return &TLUpdatePrivacy{
			Data2: &Update{
				Cmd:       Cmd_UpdatePrivacy,
				ClassName: ClassUpdatePrivacy,
			},
		}
	}
	v1 := &TLUpdatePrivacy{Data2: v}
	v1.Data2.ClassName = ClassUpdatePrivacy
	return v1
}

func NewTLUpdateUserPhone(v *Update) *TLUpdateUserPhone {
	if v == nil {
		return &TLUpdateUserPhone{
			Data2: &Update{
				Cmd:       Cmd_UpdateUserPhone,
				ClassName: ClassUpdateUserPhone,
			},
		}
	}
	v1 := &TLUpdateUserPhone{Data2: v}
	v1.Data2.ClassName = ClassUpdateUserPhone
	return v1
}

func NewTLUpdateReadHistoryInbox(v *Update) *TLUpdateReadHistoryInbox {
	if v == nil {
		return &TLUpdateReadHistoryInbox{
			Data2: &Update{
				Cmd:       Cmd_UpdateReadHistoryInbox,
				ClassName: ClassUpdateReadHistoryInbox,
			},
		}
	}
	v1 := &TLUpdateReadHistoryInbox{Data2: v}
	v1.Data2.ClassName = ClassUpdateReadHistoryInbox
	return v1
}

func NewTLUpdateReadHistoryOutbox(v *Update) *TLUpdateReadHistoryOutbox {
	if v == nil {
		return &TLUpdateReadHistoryOutbox{
			Data2: &Update{
				Cmd:       Cmd_UpdateReadHistoryOutbox,
				ClassName: ClassUpdateReadHistoryOutbox,
			},
		}
	}
	v1 := &TLUpdateReadHistoryOutbox{Data2: v}
	v1.Data2.ClassName = ClassUpdateReadHistoryOutbox
	return v1
}

func NewTLUpdateWebPage(v *Update) *TLUpdateWebPage {
	if v == nil {
		return &TLUpdateWebPage{
			Data2: &Update{
				Cmd:       Cmd_UpdateWebPage,
				ClassName: ClassUpdateWebPage,
			},
		}
	}
	v1 := &TLUpdateWebPage{Data2: v}
	v1.Data2.ClassName = ClassUpdateWebPage
	return v1
}

func NewTLUpdateReadMessagesContents(v *Update) *TLUpdateReadMessagesContents {
	if v == nil {
		return &TLUpdateReadMessagesContents{
			Data2: &Update{
				Cmd:       Cmd_UpdateReadMessagesContents,
				ClassName: ClassUpdateReadMessagesContents,
			},
		}
	}
	v1 := &TLUpdateReadMessagesContents{Data2: v}
	v1.Data2.ClassName = ClassUpdateReadMessagesContents
	return v1
}

func NewTLUpdateChannelTooLong(v *Update) *TLUpdateChannelTooLong {
	if v == nil {
		return &TLUpdateChannelTooLong{
			Data2: &Update{
				Cmd:       Cmd_UpdateChannelTooLong,
				ClassName: ClassUpdateChannelTooLong,
			},
		}
	}
	v1 := &TLUpdateChannelTooLong{Data2: v}
	v1.Data2.ClassName = ClassUpdateChannelTooLong
	return v1
}

func NewTLUpdateChannel(v *Update) *TLUpdateChannel {
	if v == nil {
		return &TLUpdateChannel{
			Data2: &Update{
				Cmd:       Cmd_UpdateChannel,
				ClassName: ClassUpdateChannel,
			},
		}
	}
	v1 := &TLUpdateChannel{Data2: v}
	v1.Data2.ClassName = ClassUpdateChannel
	return v1
}

func NewTLUpdateNewChannelMessage(v *Update) *TLUpdateNewChannelMessage {
	if v == nil {
		return &TLUpdateNewChannelMessage{
			Data2: &Update{
				Cmd:       Cmd_UpdateNewChannelMessage,
				ClassName: ClassUpdateNewChannelMessage,
			},
		}
	}
	v1 := &TLUpdateNewChannelMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateNewChannelMessage
	return v1
}

func NewTLUpdateReadChannelInbox(v *Update) *TLUpdateReadChannelInbox {
	if v == nil {
		return &TLUpdateReadChannelInbox{
			Data2: &Update{
				Cmd:       Cmd_UpdateReadChannelInbox,
				ClassName: ClassUpdateReadChannelInbox,
			},
		}
	}
	v1 := &TLUpdateReadChannelInbox{Data2: v}
	v1.Data2.ClassName = ClassUpdateReadChannelInbox
	return v1
}

func NewTLUpdateDeleteChannelMessages(v *Update) *TLUpdateDeleteChannelMessages {
	if v == nil {
		return &TLUpdateDeleteChannelMessages{
			Data2: &Update{
				Cmd:       Cmd_UpdateDeleteChannelMessages,
				ClassName: ClassUpdateDeleteChannelMessages,
			},
		}
	}
	v1 := &TLUpdateDeleteChannelMessages{Data2: v}
	v1.Data2.ClassName = ClassUpdateDeleteChannelMessages
	return v1
}

func NewTLUpdateChannelMessageViews(v *Update) *TLUpdateChannelMessageViews {
	if v == nil {
		return &TLUpdateChannelMessageViews{
			Data2: &Update{
				Cmd:       Cmd_UpdateChannelMessageViews,
				ClassName: ClassUpdateChannelMessageViews,
			},
		}
	}
	v1 := &TLUpdateChannelMessageViews{Data2: v}
	v1.Data2.ClassName = ClassUpdateChannelMessageViews
	return v1
}

func NewTLUpdateChatAdmins(v *Update) *TLUpdateChatAdmins {
	if v == nil {
		return &TLUpdateChatAdmins{
			Data2: &Update{
				Cmd:       Cmd_UpdateChatAdmins,
				ClassName: ClassUpdateChatAdmins,
			},
		}
	}
	v1 := &TLUpdateChatAdmins{Data2: v}
	v1.Data2.ClassName = ClassUpdateChatAdmins
	return v1
}

func NewTLUpdateChatParticipantAdmin(v *Update) *TLUpdateChatParticipantAdmin {
	if v == nil {
		return &TLUpdateChatParticipantAdmin{
			Data2: &Update{
				Cmd:       Cmd_UpdateChatParticipantAdmin,
				ClassName: ClassUpdateChatParticipantAdmin,
			},
		}
	}
	v1 := &TLUpdateChatParticipantAdmin{Data2: v}
	v1.Data2.ClassName = ClassUpdateChatParticipantAdmin
	return v1
}

func NewTLUpdateNewStickerSet(v *Update) *TLUpdateNewStickerSet {
	if v == nil {
		return &TLUpdateNewStickerSet{
			Data2: &Update{
				Cmd:       Cmd_UpdateNewStickerSet,
				ClassName: ClassUpdateNewStickerSet,
			},
		}
	}
	v1 := &TLUpdateNewStickerSet{Data2: v}
	v1.Data2.ClassName = ClassUpdateNewStickerSet
	return v1
}

func NewTLUpdateStickerSetsOrder(v *Update) *TLUpdateStickerSetsOrder {
	if v == nil {
		return &TLUpdateStickerSetsOrder{
			Data2: &Update{
				Cmd:       Cmd_UpdateStickerSetsOrder,
				ClassName: ClassUpdateStickerSetsOrder,
			},
		}
	}
	v1 := &TLUpdateStickerSetsOrder{Data2: v}
	v1.Data2.ClassName = ClassUpdateStickerSetsOrder
	return v1
}

func NewTLUpdateStickerSets(v *Update) *TLUpdateStickerSets {
	if v == nil {
		return &TLUpdateStickerSets{
			Data2: &Update{
				Cmd:       Cmd_UpdateStickerSets,
				ClassName: ClassUpdateStickerSets,
			},
		}
	}
	v1 := &TLUpdateStickerSets{Data2: v}
	v1.Data2.ClassName = ClassUpdateStickerSets
	return v1
}

func NewTLUpdateSavedGifs(v *Update) *TLUpdateSavedGifs {
	if v == nil {
		return &TLUpdateSavedGifs{
			Data2: &Update{
				Cmd:       Cmd_UpdateSavedGifs,
				ClassName: ClassUpdateSavedGifs,
			},
		}
	}
	v1 := &TLUpdateSavedGifs{Data2: v}
	v1.Data2.ClassName = ClassUpdateSavedGifs
	return v1
}

func NewTLUpdateBotInlineQuery(v *Update) *TLUpdateBotInlineQuery {
	if v == nil {
		return &TLUpdateBotInlineQuery{
			Data2: &Update{
				Cmd:       Cmd_UpdateBotInlineQuery,
				ClassName: ClassUpdateBotInlineQuery,
			},
		}
	}
	v1 := &TLUpdateBotInlineQuery{Data2: v}
	v1.Data2.ClassName = ClassUpdateBotInlineQuery
	return v1
}

func NewTLUpdateBotInlineSend(v *Update) *TLUpdateBotInlineSend {
	if v == nil {
		return &TLUpdateBotInlineSend{
			Data2: &Update{
				Cmd:       Cmd_UpdateBotInlineSend,
				ClassName: ClassUpdateBotInlineSend,
			},
		}
	}
	v1 := &TLUpdateBotInlineSend{Data2: v}
	v1.Data2.ClassName = ClassUpdateBotInlineSend
	return v1
}

func NewTLUpdateEditChannelMessage(v *Update) *TLUpdateEditChannelMessage {
	if v == nil {
		return &TLUpdateEditChannelMessage{
			Data2: &Update{
				Cmd:       Cmd_UpdateEditChannelMessage,
				ClassName: ClassUpdateEditChannelMessage,
			},
		}
	}
	v1 := &TLUpdateEditChannelMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateEditChannelMessage
	return v1
}

func NewTLUpdateChannelPinnedMessage(v *Update) *TLUpdateChannelPinnedMessage {
	if v == nil {
		return &TLUpdateChannelPinnedMessage{
			Data2: &Update{
				Cmd:       Cmd_UpdateChannelPinnedMessage,
				ClassName: ClassUpdateChannelPinnedMessage,
			},
		}
	}
	v1 := &TLUpdateChannelPinnedMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateChannelPinnedMessage
	return v1
}

func NewTLUpdateBotCallbackQuery(v *Update) *TLUpdateBotCallbackQuery {
	if v == nil {
		return &TLUpdateBotCallbackQuery{
			Data2: &Update{
				Cmd:       Cmd_UpdateBotCallbackQuery,
				ClassName: ClassUpdateBotCallbackQuery,
			},
		}
	}
	v1 := &TLUpdateBotCallbackQuery{Data2: v}
	v1.Data2.ClassName = ClassUpdateBotCallbackQuery
	return v1
}

func NewTLUpdateEditMessage(v *Update) *TLUpdateEditMessage {
	if v == nil {
		return &TLUpdateEditMessage{
			Data2: &Update{
				Cmd:       Cmd_UpdateEditMessage,
				ClassName: ClassUpdateEditMessage,
			},
		}
	}
	v1 := &TLUpdateEditMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateEditMessage
	return v1
}

func NewTLUpdateInlineBotCallbackQuery(v *Update) *TLUpdateInlineBotCallbackQuery {
	if v == nil {
		return &TLUpdateInlineBotCallbackQuery{
			Data2: &Update{
				Cmd:       Cmd_UpdateInlineBotCallbackQuery,
				ClassName: ClassUpdateInlineBotCallbackQuery,
			},
		}
	}
	v1 := &TLUpdateInlineBotCallbackQuery{Data2: v}
	v1.Data2.ClassName = ClassUpdateInlineBotCallbackQuery
	return v1
}

func NewTLUpdateReadChannelOutbox(v *Update) *TLUpdateReadChannelOutbox {
	if v == nil {
		return &TLUpdateReadChannelOutbox{
			Data2: &Update{
				Cmd:       Cmd_UpdateReadChannelOutbox,
				ClassName: ClassUpdateReadChannelOutbox,
			},
		}
	}
	v1 := &TLUpdateReadChannelOutbox{Data2: v}
	v1.Data2.ClassName = ClassUpdateReadChannelOutbox
	return v1
}

func NewTLUpdateDraftMessage(v *Update) *TLUpdateDraftMessage {
	if v == nil {
		return &TLUpdateDraftMessage{
			Data2: &Update{
				Cmd:       Cmd_UpdateDraftMessage,
				ClassName: ClassUpdateDraftMessage,
			},
		}
	}
	v1 := &TLUpdateDraftMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateDraftMessage
	return v1
}

func NewTLUpdateReadFeaturedStickers(v *Update) *TLUpdateReadFeaturedStickers {
	if v == nil {
		return &TLUpdateReadFeaturedStickers{
			Data2: &Update{
				Cmd:       Cmd_UpdateReadFeaturedStickers,
				ClassName: ClassUpdateReadFeaturedStickers,
			},
		}
	}
	v1 := &TLUpdateReadFeaturedStickers{Data2: v}
	v1.Data2.ClassName = ClassUpdateReadFeaturedStickers
	return v1
}

func NewTLUpdateRecentStickers(v *Update) *TLUpdateRecentStickers {
	if v == nil {
		return &TLUpdateRecentStickers{
			Data2: &Update{
				Cmd:       Cmd_UpdateRecentStickers,
				ClassName: ClassUpdateRecentStickers,
			},
		}
	}
	v1 := &TLUpdateRecentStickers{Data2: v}
	v1.Data2.ClassName = ClassUpdateRecentStickers
	return v1
}

func NewTLUpdateConfig(v *Update) *TLUpdateConfig {
	if v == nil {
		return &TLUpdateConfig{
			Data2: &Update{
				Cmd:       Cmd_UpdateConfig,
				ClassName: ClassUpdateConfig,
			},
		}
	}
	v1 := &TLUpdateConfig{Data2: v}
	v1.Data2.ClassName = ClassUpdateConfig
	return v1
}

func NewTLUpdatePtsChanged(v *Update) *TLUpdatePtsChanged {
	if v == nil {
		return &TLUpdatePtsChanged{
			Data2: &Update{
				Cmd:       Cmd_UpdatePtsChanged,
				ClassName: ClassUpdatePtsChanged,
			},
		}
	}
	v1 := &TLUpdatePtsChanged{Data2: v}
	v1.Data2.ClassName = ClassUpdatePtsChanged
	return v1
}

func NewTLUpdateChannelWebPage(v *Update) *TLUpdateChannelWebPage {
	if v == nil {
		return &TLUpdateChannelWebPage{
			Data2: &Update{
				Cmd:       Cmd_UpdateChannelWebPage,
				ClassName: ClassUpdateChannelWebPage,
			},
		}
	}
	v1 := &TLUpdateChannelWebPage{Data2: v}
	v1.Data2.ClassName = ClassUpdateChannelWebPage
	return v1
}

func NewTLUpdateDialogPinned(v *Update) *TLUpdateDialogPinned {
	if v == nil {
		return &TLUpdateDialogPinned{
			Data2: &Update{
				Cmd:       Cmd_UpdateDialogPinned,
				ClassName: ClassUpdateDialogPinned,
			},
		}
	}
	v1 := &TLUpdateDialogPinned{Data2: v}
	v1.Data2.ClassName = ClassUpdateDialogPinned
	return v1
}

func NewTLUpdatePinnedDialogs(v *Update) *TLUpdatePinnedDialogs {
	if v == nil {
		return &TLUpdatePinnedDialogs{
			Data2: &Update{
				Cmd:       Cmd_UpdatePinnedDialogs,
				ClassName: ClassUpdatePinnedDialogs,
			},
		}
	}
	v1 := &TLUpdatePinnedDialogs{Data2: v}
	v1.Data2.ClassName = ClassUpdatePinnedDialogs
	return v1
}

func NewTLUpdateBotWebhookJSON(v *Update) *TLUpdateBotWebhookJSON {
	if v == nil {
		return &TLUpdateBotWebhookJSON{
			Data2: &Update{
				Cmd:       Cmd_UpdateBotWebhookJSON,
				ClassName: ClassUpdateBotWebhookJSON,
			},
		}
	}
	v1 := &TLUpdateBotWebhookJSON{Data2: v}
	v1.Data2.ClassName = ClassUpdateBotWebhookJSON
	return v1
}

func NewTLUpdateBotWebhookJSONQuery(v *Update) *TLUpdateBotWebhookJSONQuery {
	if v == nil {
		return &TLUpdateBotWebhookJSONQuery{
			Data2: &Update{
				Cmd:       Cmd_UpdateBotWebhookJSONQuery,
				ClassName: ClassUpdateBotWebhookJSONQuery,
			},
		}
	}
	v1 := &TLUpdateBotWebhookJSONQuery{Data2: v}
	v1.Data2.ClassName = ClassUpdateBotWebhookJSONQuery
	return v1
}

func NewTLUpdateBotShippingQuery(v *Update) *TLUpdateBotShippingQuery {
	if v == nil {
		return &TLUpdateBotShippingQuery{
			Data2: &Update{
				Cmd:       Cmd_UpdateBotShippingQuery,
				ClassName: ClassUpdateBotShippingQuery,
			},
		}
	}
	v1 := &TLUpdateBotShippingQuery{Data2: v}
	v1.Data2.ClassName = ClassUpdateBotShippingQuery
	return v1
}

func NewTLUpdateBotPrecheckoutQuery(v *Update) *TLUpdateBotPrecheckoutQuery {
	if v == nil {
		return &TLUpdateBotPrecheckoutQuery{
			Data2: &Update{
				Cmd:       Cmd_UpdateBotPrecheckoutQuery,
				ClassName: ClassUpdateBotPrecheckoutQuery,
			},
		}
	}
	v1 := &TLUpdateBotPrecheckoutQuery{Data2: v}
	v1.Data2.ClassName = ClassUpdateBotPrecheckoutQuery
	return v1
}

func NewTLUpdatePhoneCall(v *Update) *TLUpdatePhoneCall {
	if v == nil {
		return &TLUpdatePhoneCall{
			Data2: &Update{
				Cmd:       Cmd_UpdatePhoneCall,
				ClassName: ClassUpdatePhoneCall,
			},
		}
	}
	v1 := &TLUpdatePhoneCall{Data2: v}
	v1.Data2.ClassName = ClassUpdatePhoneCall
	return v1
}

func NewTLUpdateLangPackTooLong(v *Update) *TLUpdateLangPackTooLong {
	if v == nil {
		return &TLUpdateLangPackTooLong{
			Data2: &Update{
				Cmd:       Cmd_UpdateLangPackTooLong,
				ClassName: ClassUpdateLangPackTooLong,
			},
		}
	}
	v1 := &TLUpdateLangPackTooLong{Data2: v}
	v1.Data2.ClassName = ClassUpdateLangPackTooLong
	return v1
}

func NewTLUpdateLangPack(v *Update) *TLUpdateLangPack {
	if v == nil {
		return &TLUpdateLangPack{
			Data2: &Update{
				Cmd:       Cmd_UpdateLangPack,
				ClassName: ClassUpdateLangPack,
			},
		}
	}
	v1 := &TLUpdateLangPack{Data2: v}
	v1.Data2.ClassName = ClassUpdateLangPack
	return v1
}

func NewTLUpdateFavedStickers(v *Update) *TLUpdateFavedStickers {
	if v == nil {
		return &TLUpdateFavedStickers{
			Data2: &Update{
				Cmd:       Cmd_UpdateFavedStickers,
				ClassName: ClassUpdateFavedStickers,
			},
		}
	}
	v1 := &TLUpdateFavedStickers{Data2: v}
	v1.Data2.ClassName = ClassUpdateFavedStickers
	return v1
}

func NewTLUpdateChannelReadMessagesContents(v *Update) *TLUpdateChannelReadMessagesContents {
	if v == nil {
		return &TLUpdateChannelReadMessagesContents{
			Data2: &Update{
				Cmd:       Cmd_UpdateChannelReadMessagesContents,
				ClassName: ClassUpdateChannelReadMessagesContents,
			},
		}
	}
	v1 := &TLUpdateChannelReadMessagesContents{Data2: v}
	v1.Data2.ClassName = ClassUpdateChannelReadMessagesContents
	return v1
}

func NewTLUpdateContactsReset(v *Update) *TLUpdateContactsReset {
	if v == nil {
		return &TLUpdateContactsReset{
			Data2: &Update{
				Cmd:       Cmd_UpdateContactsReset,
				ClassName: ClassUpdateContactsReset,
			},
		}
	}
	v1 := &TLUpdateContactsReset{Data2: v}
	v1.Data2.ClassName = ClassUpdateContactsReset
	return v1
}

func NewTLUpdateNewAuthorization(v *Update) *TLUpdateNewAuthorization {
	if v == nil {
		return &TLUpdateNewAuthorization{
			Data2: &Update{
				Cmd:       Cmd_UpdateNewAuthorization,
				ClassName: ClassUpdateNewAuthorization,
			},
		}
	}
	v1 := &TLUpdateNewAuthorization{Data2: v}
	v1.Data2.ClassName = ClassUpdateNewAuthorization
	return v1
}

func NewTLUpdateChannelGroup(v *Update) *TLUpdateChannelGroup {
	if v == nil {
		return &TLUpdateChannelGroup{
			Data2: &Update{
				Cmd:       Cmd_UpdateChannelGroup,
				ClassName: ClassUpdateChannelGroup,
			},
		}
	}
	v1 := &TLUpdateChannelGroup{Data2: v}
	v1.Data2.ClassName = ClassUpdateChannelGroup
	return v1
}

func NewTLUpdateChannelAvailableMessages(v *Update) *TLUpdateChannelAvailableMessages {
	if v == nil {
		return &TLUpdateChannelAvailableMessages{
			Data2: &Update{
				Cmd:       Cmd_UpdateChannelAvailableMessages,
				ClassName: ClassUpdateChannelAvailableMessages,
			},
		}
	}
	v1 := &TLUpdateChannelAvailableMessages{Data2: v}
	v1.Data2.ClassName = ClassUpdateChannelAvailableMessages
	return v1
}

func NewTLUpdateDialogUnreadMark(v *Update) *TLUpdateDialogUnreadMark {
	if v == nil {
		return &TLUpdateDialogUnreadMark{
			Data2: &Update{
				Cmd:       Cmd_UpdateDialogUnreadMark,
				ClassName: ClassUpdateDialogUnreadMark,
			},
		}
	}
	v1 := &TLUpdateDialogUnreadMark{Data2: v}
	v1.Data2.ClassName = ClassUpdateDialogUnreadMark
	return v1
}

func NewTLUpdateUserPinnedMessage(v *Update) *TLUpdateUserPinnedMessage {
	if v == nil {
		return &TLUpdateUserPinnedMessage{
			Data2: &Update{
				Cmd:       Cmd_UpdateUserPinnedMessage,
				ClassName: ClassUpdateUserPinnedMessage,
			},
		}
	}
	v1 := &TLUpdateUserPinnedMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateUserPinnedMessage
	return v1
}

func NewTLUpdateChatPinnedMessage(v *Update) *TLUpdateChatPinnedMessage {
	if v == nil {
		return &TLUpdateChatPinnedMessage{
			Data2: &Update{
				Cmd:       Cmd_UpdateChatPinnedMessage,
				ClassName: ClassUpdateChatPinnedMessage,
			},
		}
	}
	v1 := &TLUpdateChatPinnedMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateChatPinnedMessage
	return v1
}

func NewTLUpdateMessagePoll(v *Update) *TLUpdateMessagePoll {
	if v == nil {
		return &TLUpdateMessagePoll{
			Data2: &Update{
				Cmd:       Cmd_UpdateMessagePoll,
				ClassName: ClassUpdateMessagePoll,
			},
		}
	}
	v1 := &TLUpdateMessagePoll{Data2: v}
	v1.Data2.ClassName = ClassUpdateMessagePoll
	return v1
}

func NewTLUpdateChatDefaultBannedRights(v *Update) *TLUpdateChatDefaultBannedRights {
	if v == nil {
		return &TLUpdateChatDefaultBannedRights{
			Data2: &Update{
				Cmd:       Cmd_UpdateChatDefaultBannedRights,
				ClassName: ClassUpdateChatDefaultBannedRights,
			},
		}
	}
	v1 := &TLUpdateChatDefaultBannedRights{Data2: v}
	v1.Data2.ClassName = ClassUpdateChatDefaultBannedRights
	return v1
}

func NewTLUpdateFolderPeers(v *Update) *TLUpdateFolderPeers {
	if v == nil {
		return &TLUpdateFolderPeers{
			Data2: &Update{
				Cmd:       Cmd_UpdateFolderPeers,
				ClassName: ClassUpdateFolderPeers,
			},
		}
	}
	v1 := &TLUpdateFolderPeers{Data2: v}
	v1.Data2.ClassName = ClassUpdateFolderPeers
	return v1
}

func NewTLUpdatePeerSettings(v *Update) *TLUpdatePeerSettings {
	if v == nil {
		return &TLUpdatePeerSettings{
			Data2: &Update{
				Cmd:       Cmd_UpdatePeerSettings,
				ClassName: ClassUpdatePeerSettings,
			},
		}
	}
	v1 := &TLUpdatePeerSettings{Data2: v}
	v1.Data2.ClassName = ClassUpdatePeerSettings
	return v1
}

func NewTLUpdatePeerLocated(v *Update) *TLUpdatePeerLocated {
	if v == nil {
		return &TLUpdatePeerLocated{
			Data2: &Update{
				Cmd:       Cmd_UpdatePeerLocated,
				ClassName: ClassUpdatePeerLocated,
			},
		}
	}
	v1 := &TLUpdatePeerLocated{Data2: v}
	v1.Data2.ClassName = ClassUpdatePeerLocated
	return v1
}

func NewTLUpdateNewScheduledMessage(v *Update) *TLUpdateNewScheduledMessage {
	if v == nil {
		return &TLUpdateNewScheduledMessage{
			Data2: &Update{
				Cmd:       Cmd_UpdateNewScheduledMessage,
				ClassName: ClassUpdateNewScheduledMessage,
			},
		}
	}
	v1 := &TLUpdateNewScheduledMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateNewScheduledMessage
	return v1
}

func NewTLUpdateDeleteScheduledMessages(v *Update) *TLUpdateDeleteScheduledMessages {
	if v == nil {
		return &TLUpdateDeleteScheduledMessages{
			Data2: &Update{
				Cmd:       Cmd_UpdateDeleteScheduledMessages,
				ClassName: ClassUpdateDeleteScheduledMessages,
			},
		}
	}
	v1 := &TLUpdateDeleteScheduledMessages{Data2: v}
	v1.Data2.ClassName = ClassUpdateDeleteScheduledMessages
	return v1
}

func NewTLUpdateTheme(v *Update) *TLUpdateTheme {
	if v == nil {
		return &TLUpdateTheme{
			Data2: &Update{
				Cmd:       Cmd_UpdateTheme,
				ClassName: ClassUpdateTheme,
			},
		}
	}
	v1 := &TLUpdateTheme{Data2: v}
	v1.Data2.ClassName = ClassUpdateTheme
	return v1
}

func NewTLUpdateGeoLiveViewed(v *Update) *TLUpdateGeoLiveViewed {
	if v == nil {
		return &TLUpdateGeoLiveViewed{
			Data2: &Update{
				Cmd:       Cmd_UpdateGeoLiveViewed,
				ClassName: ClassUpdateGeoLiveViewed,
			},
		}
	}
	v1 := &TLUpdateGeoLiveViewed{Data2: v}
	v1.Data2.ClassName = ClassUpdateGeoLiveViewed
	return v1
}

func NewTLUpdateLoginToken(v *Update) *TLUpdateLoginToken {
	if v == nil {
		return &TLUpdateLoginToken{
			Data2: &Update{
				Cmd:       Cmd_UpdateLoginToken,
				ClassName: ClassUpdateLoginToken,
			},
		}
	}
	v1 := &TLUpdateLoginToken{Data2: v}
	v1.Data2.ClassName = ClassUpdateLoginToken
	return v1
}

func NewTLUpdateMessagePollVote(v *Update) *TLUpdateMessagePollVote {
	if v == nil {
		return &TLUpdateMessagePollVote{
			Data2: &Update{
				Cmd:       Cmd_UpdateMessagePollVote,
				ClassName: ClassUpdateMessagePollVote,
			},
		}
	}
	v1 := &TLUpdateMessagePollVote{Data2: v}
	v1.Data2.ClassName = ClassUpdateMessagePollVote
	return v1
}

func NewTLUpdateDialogFilter(v *Update) *TLUpdateDialogFilter {
	if v == nil {
		return &TLUpdateDialogFilter{
			Data2: &Update{
				Cmd:       Cmd_UpdateDialogFilter,
				ClassName: ClassUpdateDialogFilter,
			},
		}
	}
	v1 := &TLUpdateDialogFilter{Data2: v}
	v1.Data2.ClassName = ClassUpdateDialogFilter
	return v1
}

func NewTLUpdateDialogFilterOrder(v *Update) *TLUpdateDialogFilterOrder {
	if v == nil {
		return &TLUpdateDialogFilterOrder{
			Data2: &Update{
				Cmd:       Cmd_UpdateDialogFilterOrder,
				ClassName: ClassUpdateDialogFilterOrder,
			},
		}
	}
	v1 := &TLUpdateDialogFilterOrder{Data2: v}
	v1.Data2.ClassName = ClassUpdateDialogFilterOrder
	return v1
}

func NewTLUpdateDialogFilters(v *Update) *TLUpdateDialogFilters {
	if v == nil {
		return &TLUpdateDialogFilters{
			Data2: &Update{
				Cmd:       Cmd_UpdateDialogFilters,
				ClassName: ClassUpdateDialogFilters,
			},
		}
	}
	v1 := &TLUpdateDialogFilters{Data2: v}
	v1.Data2.ClassName = ClassUpdateDialogFilters
	return v1
}

func NewTLUpdatePhoneCallSignalingData(v *Update) *TLUpdatePhoneCallSignalingData {
	if v == nil {
		return &TLUpdatePhoneCallSignalingData{
			Data2: &Update{
				Cmd:       Cmd_UpdatePhoneCallSignalingData,
				ClassName: ClassUpdatePhoneCallSignalingData,
			},
		}
	}
	v1 := &TLUpdatePhoneCallSignalingData{Data2: v}
	v1.Data2.ClassName = ClassUpdatePhoneCallSignalingData
	return v1
}

func NewTLUpdateChannelParticipant(v *Update) *TLUpdateChannelParticipant {
	if v == nil {
		return &TLUpdateChannelParticipant{
			Data2: &Update{
				Cmd:       Cmd_UpdateChannelParticipant,
				ClassName: ClassUpdateChannelParticipant,
			},
		}
	}
	v1 := &TLUpdateChannelParticipant{Data2: v}
	v1.Data2.ClassName = ClassUpdateChannelParticipant
	return v1
}

func NewTLUpdateChannelMessageForwards(v *Update) *TLUpdateChannelMessageForwards {
	if v == nil {
		return &TLUpdateChannelMessageForwards{
			Data2: &Update{
				Cmd:       Cmd_UpdateChannelMessageForwards,
				ClassName: ClassUpdateChannelMessageForwards,
			},
		}
	}
	v1 := &TLUpdateChannelMessageForwards{Data2: v}
	v1.Data2.ClassName = ClassUpdateChannelMessageForwards
	return v1
}

func NewTLUpdateReadChannelDiscussionInbox(v *Update) *TLUpdateReadChannelDiscussionInbox {
	if v == nil {
		return &TLUpdateReadChannelDiscussionInbox{
			Data2: &Update{
				Cmd:       Cmd_UpdateReadChannelDiscussionInbox,
				ClassName: ClassUpdateReadChannelDiscussionInbox,
			},
		}
	}
	v1 := &TLUpdateReadChannelDiscussionInbox{Data2: v}
	v1.Data2.ClassName = ClassUpdateReadChannelDiscussionInbox
	return v1
}

func NewTLUpdateReadChannelDiscussionOutbox(v *Update) *TLUpdateReadChannelDiscussionOutbox {
	if v == nil {
		return &TLUpdateReadChannelDiscussionOutbox{
			Data2: &Update{
				Cmd:       Cmd_UpdateReadChannelDiscussionOutbox,
				ClassName: ClassUpdateReadChannelDiscussionOutbox,
			},
		}
	}
	v1 := &TLUpdateReadChannelDiscussionOutbox{Data2: v}
	v1.Data2.ClassName = ClassUpdateReadChannelDiscussionOutbox
	return v1
}

func NewTLUpdatePeerBlocked(v *Update) *TLUpdatePeerBlocked {
	if v == nil {
		return &TLUpdatePeerBlocked{
			Data2: &Update{
				Cmd:       Cmd_UpdatePeerBlocked,
				ClassName: ClassUpdatePeerBlocked,
			},
		}
	}
	v1 := &TLUpdatePeerBlocked{Data2: v}
	v1.Data2.ClassName = ClassUpdatePeerBlocked
	return v1
}

func NewTLUpdateChannelUserTyping(v *Update) *TLUpdateChannelUserTyping {
	if v == nil {
		return &TLUpdateChannelUserTyping{
			Data2: &Update{
				Cmd:       Cmd_UpdateChannelUserTyping,
				ClassName: ClassUpdateChannelUserTyping,
			},
		}
	}
	v1 := &TLUpdateChannelUserTyping{Data2: v}
	v1.Data2.ClassName = ClassUpdateChannelUserTyping
	return v1
}

func NewTLUpdatePinnedMessages(v *Update) *TLUpdatePinnedMessages {
	if v == nil {
		return &TLUpdatePinnedMessages{
			Data2: &Update{
				Cmd:       Cmd_UpdatePinnedMessages,
				ClassName: ClassUpdatePinnedMessages,
			},
		}
	}
	v1 := &TLUpdatePinnedMessages{Data2: v}
	v1.Data2.ClassName = ClassUpdatePinnedMessages
	return v1
}

func NewTLUpdatePinnedChannelMessages(v *Update) *TLUpdatePinnedChannelMessages {
	if v == nil {
		return &TLUpdatePinnedChannelMessages{
			Data2: &Update{
				Cmd:       Cmd_UpdatePinnedChannelMessages,
				ClassName: ClassUpdatePinnedChannelMessages,
			},
		}
	}
	v1 := &TLUpdatePinnedChannelMessages{Data2: v}
	v1.Data2.ClassName = ClassUpdatePinnedChannelMessages
	return v1
}

func NewTLUpdateChat(v *Update) *TLUpdateChat {
	if v == nil {
		return &TLUpdateChat{
			Data2: &Update{
				Cmd:       Cmd_UpdateChat,
				ClassName: ClassUpdateChat,
			},
		}
	}
	v1 := &TLUpdateChat{Data2: v}
	v1.Data2.ClassName = ClassUpdateChat
	return v1
}

func NewTLUpdateGroupCallParticipants(v *Update) *TLUpdateGroupCallParticipants {
	if v == nil {
		return &TLUpdateGroupCallParticipants{
			Data2: &Update{
				Cmd:       Cmd_UpdateGroupCallParticipants,
				ClassName: ClassUpdateGroupCallParticipants,
			},
		}
	}
	v1 := &TLUpdateGroupCallParticipants{Data2: v}
	v1.Data2.ClassName = ClassUpdateGroupCallParticipants
	return v1
}

func NewTLUpdateGroupCall(v *Update) *TLUpdateGroupCall {
	if v == nil {
		return &TLUpdateGroupCall{
			Data2: &Update{
				Cmd:       Cmd_UpdateGroupCall,
				ClassName: ClassUpdateGroupCall,
			},
		}
	}
	v1 := &TLUpdateGroupCall{Data2: v}
	v1.Data2.ClassName = ClassUpdateGroupCall
	return v1
}

func NewTLUpdatesTooLong(v *Updates) *TLUpdatesTooLong {
	if v == nil {
		return &TLUpdatesTooLong{
			Data2: &Updates{
				Cmd:       Cmd_UpdatesTooLong,
				ClassName: ClassUpdatesTooLong,
			},
		}
	}
	v1 := &TLUpdatesTooLong{Data2: v}
	v1.Data2.ClassName = ClassUpdatesTooLong
	return v1
}

func NewTLUpdateShortMessage(v *Updates) *TLUpdateShortMessage {
	if v == nil {
		return &TLUpdateShortMessage{
			Data2: &Updates{
				Cmd:       Cmd_UpdateShortMessage,
				ClassName: ClassUpdateShortMessage,
			},
		}
	}
	v1 := &TLUpdateShortMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateShortMessage
	return v1
}

func NewTLUpdateShortChatMessage(v *Updates) *TLUpdateShortChatMessage {
	if v == nil {
		return &TLUpdateShortChatMessage{
			Data2: &Updates{
				Cmd:       Cmd_UpdateShortChatMessage,
				ClassName: ClassUpdateShortChatMessage,
			},
		}
	}
	v1 := &TLUpdateShortChatMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateShortChatMessage
	return v1
}

func NewTLUpdateShort(v *Updates) *TLUpdateShort {
	if v == nil {
		return &TLUpdateShort{
			Data2: &Updates{
				Cmd:       Cmd_UpdateShort,
				ClassName: ClassUpdateShort,
			},
		}
	}
	v1 := &TLUpdateShort{Data2: v}
	v1.Data2.ClassName = ClassUpdateShort
	return v1
}

func NewTLUpdatesCombined(v *Updates) *TLUpdatesCombined {
	if v == nil {
		return &TLUpdatesCombined{
			Data2: &Updates{
				Cmd:       Cmd_UpdatesCombined,
				ClassName: ClassUpdatesCombined,
			},
		}
	}
	v1 := &TLUpdatesCombined{Data2: v}
	v1.Data2.ClassName = ClassUpdatesCombined
	return v1
}

func NewTLUpdates(v *Updates) *TLUpdates {
	if v == nil {
		return &TLUpdates{
			Data2: &Updates{
				Cmd:       Cmd_Updates,
				ClassName: ClassUpdates,
			},
		}
	}
	v1 := &TLUpdates{Data2: v}
	v1.Data2.ClassName = ClassUpdates
	return v1
}

func NewTLUpdateShortSentMessage(v *Updates) *TLUpdateShortSentMessage {
	if v == nil {
		return &TLUpdateShortSentMessage{
			Data2: &Updates{
				Cmd:       Cmd_UpdateShortSentMessage,
				ClassName: ClassUpdateShortSentMessage,
			},
		}
	}
	v1 := &TLUpdateShortSentMessage{Data2: v}
	v1.Data2.ClassName = ClassUpdateShortSentMessage
	return v1
}

func NewTLUpdatesChannelDifferenceEmpty(v *Updates_ChannelDifference) *TLUpdatesChannelDifferenceEmpty {
	if v == nil {
		return &TLUpdatesChannelDifferenceEmpty{
			Data2: &Updates_ChannelDifference{
				Cmd:       Cmd_UpdatesChannelDifferenceEmpty,
				ClassName: ClassUpdatesChannelDifferenceEmpty,
			},
		}
	}
	v1 := &TLUpdatesChannelDifferenceEmpty{Data2: v}
	v1.Data2.ClassName = ClassUpdatesChannelDifferenceEmpty
	return v1
}

func NewTLUpdatesChannelDifferenceTooLong(v *Updates_ChannelDifference) *TLUpdatesChannelDifferenceTooLong {
	if v == nil {
		return &TLUpdatesChannelDifferenceTooLong{
			Data2: &Updates_ChannelDifference{
				Cmd:       Cmd_UpdatesChannelDifferenceTooLong,
				ClassName: ClassUpdatesChannelDifferenceTooLong,
			},
		}
	}
	v1 := &TLUpdatesChannelDifferenceTooLong{Data2: v}
	v1.Data2.ClassName = ClassUpdatesChannelDifferenceTooLong
	return v1
}

func NewTLUpdatesChannelDifference(v *Updates_ChannelDifference) *TLUpdatesChannelDifference {
	if v == nil {
		return &TLUpdatesChannelDifference{
			Data2: &Updates_ChannelDifference{
				Cmd:       Cmd_UpdatesChannelDifference,
				ClassName: ClassUpdatesChannelDifference,
			},
		}
	}
	v1 := &TLUpdatesChannelDifference{Data2: v}
	v1.Data2.ClassName = ClassUpdatesChannelDifference
	return v1
}

func NewTLUpdatesDifferenceEmpty(v *Updates_Difference) *TLUpdatesDifferenceEmpty {
	if v == nil {
		return &TLUpdatesDifferenceEmpty{
			Data2: &Updates_Difference{
				Cmd:       Cmd_UpdatesDifferenceEmpty,
				ClassName: ClassUpdatesDifferenceEmpty,
			},
		}
	}
	v1 := &TLUpdatesDifferenceEmpty{Data2: v}
	v1.Data2.ClassName = ClassUpdatesDifferenceEmpty
	return v1
}

func NewTLUpdatesDifference(v *Updates_Difference) *TLUpdatesDifference {
	if v == nil {
		return &TLUpdatesDifference{
			Data2: &Updates_Difference{
				Cmd:       Cmd_UpdatesDifference,
				ClassName: ClassUpdatesDifference,
			},
		}
	}
	v1 := &TLUpdatesDifference{Data2: v}
	v1.Data2.ClassName = ClassUpdatesDifference
	return v1
}

func NewTLUpdatesDifferenceSlice(v *Updates_Difference) *TLUpdatesDifferenceSlice {
	if v == nil {
		return &TLUpdatesDifferenceSlice{
			Data2: &Updates_Difference{
				Cmd:       Cmd_UpdatesDifferenceSlice,
				ClassName: ClassUpdatesDifferenceSlice,
			},
		}
	}
	v1 := &TLUpdatesDifferenceSlice{Data2: v}
	v1.Data2.ClassName = ClassUpdatesDifferenceSlice
	return v1
}

func NewTLUpdatesDifferenceTooLong(v *Updates_Difference) *TLUpdatesDifferenceTooLong {
	if v == nil {
		return &TLUpdatesDifferenceTooLong{
			Data2: &Updates_Difference{
				Cmd:       Cmd_UpdatesDifferenceTooLong,
				ClassName: ClassUpdatesDifferenceTooLong,
			},
		}
	}
	v1 := &TLUpdatesDifferenceTooLong{Data2: v}
	v1.Data2.ClassName = ClassUpdatesDifferenceTooLong
	return v1
}

func NewTLUpdatesState(v *Updates_State) *TLUpdatesState {
	if v == nil {
		return &TLUpdatesState{
			Data2: &Updates_State{
				Cmd:       Cmd_UpdatesState,
				ClassName: ClassUpdatesState,
			},
		}
	}
	v1 := &TLUpdatesState{Data2: v}
	v1.Data2.ClassName = ClassUpdatesState
	return v1
}

func NewTLUploadCdnFileReuploadNeeded(v *Upload_CdnFile) *TLUploadCdnFileReuploadNeeded {
	if v == nil {
		return &TLUploadCdnFileReuploadNeeded{
			Data2: &Upload_CdnFile{
				Cmd:       Cmd_UploadCdnFileReuploadNeeded,
				ClassName: ClassUploadCdnFileReuploadNeeded,
			},
		}
	}
	v1 := &TLUploadCdnFileReuploadNeeded{Data2: v}
	v1.Data2.ClassName = ClassUploadCdnFileReuploadNeeded
	return v1
}

func NewTLUploadCdnFile(v *Upload_CdnFile) *TLUploadCdnFile {
	if v == nil {
		return &TLUploadCdnFile{
			Data2: &Upload_CdnFile{
				Cmd:       Cmd_UploadCdnFile,
				ClassName: ClassUploadCdnFile,
			},
		}
	}
	v1 := &TLUploadCdnFile{Data2: v}
	v1.Data2.ClassName = ClassUploadCdnFile
	return v1
}

func NewTLUploadFile(v *Upload_File) *TLUploadFile {
	if v == nil {
		return &TLUploadFile{
			Data2: &Upload_File{
				Cmd:       Cmd_UploadFile,
				ClassName: ClassUploadFile,
			},
		}
	}
	v1 := &TLUploadFile{Data2: v}
	v1.Data2.ClassName = ClassUploadFile
	return v1
}

func NewTLUploadFileCdnRedirect(v *Upload_File) *TLUploadFileCdnRedirect {
	if v == nil {
		return &TLUploadFileCdnRedirect{
			Data2: &Upload_File{
				Cmd:       Cmd_UploadFileCdnRedirect,
				ClassName: ClassUploadFileCdnRedirect,
			},
		}
	}
	v1 := &TLUploadFileCdnRedirect{Data2: v}
	v1.Data2.ClassName = ClassUploadFileCdnRedirect
	return v1
}

func NewTLUploadWebFile(v *Upload_WebFile) *TLUploadWebFile {
	if v == nil {
		return &TLUploadWebFile{
			Data2: &Upload_WebFile{
				Cmd:       Cmd_UploadWebFile,
				ClassName: ClassUploadWebFile,
			},
		}
	}
	v1 := &TLUploadWebFile{Data2: v}
	v1.Data2.ClassName = ClassUploadWebFile
	return v1
}

func NewTLUrlAuthResultRequest(v *UrlAuthResult) *TLUrlAuthResultRequest {
	if v == nil {
		return &TLUrlAuthResultRequest{
			Data2: &UrlAuthResult{
				Cmd:       Cmd_UrlAuthResultRequest,
				ClassName: ClassUrlAuthResultRequest,
			},
		}
	}
	v1 := &TLUrlAuthResultRequest{Data2: v}
	v1.Data2.ClassName = ClassUrlAuthResultRequest
	return v1
}

func NewTLUrlAuthResultAccepted(v *UrlAuthResult) *TLUrlAuthResultAccepted {
	if v == nil {
		return &TLUrlAuthResultAccepted{
			Data2: &UrlAuthResult{
				Cmd:       Cmd_UrlAuthResultAccepted,
				ClassName: ClassUrlAuthResultAccepted,
			},
		}
	}
	v1 := &TLUrlAuthResultAccepted{Data2: v}
	v1.Data2.ClassName = ClassUrlAuthResultAccepted
	return v1
}

func NewTLUrlAuthResultDefault(v *UrlAuthResult) *TLUrlAuthResultDefault {
	if v == nil {
		return &TLUrlAuthResultDefault{
			Data2: &UrlAuthResult{
				Cmd:       Cmd_UrlAuthResultDefault,
				ClassName: ClassUrlAuthResultDefault,
			},
		}
	}
	v1 := &TLUrlAuthResultDefault{Data2: v}
	v1.Data2.ClassName = ClassUrlAuthResultDefault
	return v1
}

func NewTLUserEmpty(v *User) *TLUserEmpty {
	if v == nil {
		return &TLUserEmpty{
			Data2: &User{
				Cmd:       Cmd_UserEmpty,
				ClassName: ClassUserEmpty,
			},
		}
	}
	v1 := &TLUserEmpty{Data2: v}
	v1.Data2.ClassName = ClassUserEmpty
	return v1
}

func NewTLUser(v *User) *TLUser {
	if v == nil {
		return &TLUser{
			Data2: &User{
				Cmd:       Cmd_User,
				ClassName: ClassUser,
			},
		}
	}
	v1 := &TLUser{Data2: v}
	v1.Data2.ClassName = ClassUser
	return v1
}

func NewTLUserFull(v *UserFull) *TLUserFull {
	if v == nil {
		return &TLUserFull{
			Data2: &UserFull{
				Cmd:       Cmd_UserFull,
				ClassName: ClassUserFull,
			},
		}
	}
	v1 := &TLUserFull{Data2: v}
	v1.Data2.ClassName = ClassUserFull
	return v1
}

func NewTLUserProfilePhotoEmpty(v *UserProfilePhoto) *TLUserProfilePhotoEmpty {
	if v == nil {
		return &TLUserProfilePhotoEmpty{
			Data2: &UserProfilePhoto{
				Cmd:       Cmd_UserProfilePhotoEmpty,
				ClassName: ClassUserProfilePhotoEmpty,
			},
		}
	}
	v1 := &TLUserProfilePhotoEmpty{Data2: v}
	v1.Data2.ClassName = ClassUserProfilePhotoEmpty
	return v1
}

func NewTLUserProfilePhoto(v *UserProfilePhoto) *TLUserProfilePhoto {
	if v == nil {
		return &TLUserProfilePhoto{
			Data2: &UserProfilePhoto{
				Cmd:       Cmd_UserProfilePhoto,
				ClassName: ClassUserProfilePhoto,
			},
		}
	}
	v1 := &TLUserProfilePhoto{Data2: v}
	v1.Data2.ClassName = ClassUserProfilePhoto
	return v1
}

func NewTLUserStatusEmpty(v *UserStatus) *TLUserStatusEmpty {
	if v == nil {
		return &TLUserStatusEmpty{
			Data2: &UserStatus{
				Cmd:       Cmd_UserStatusEmpty,
				ClassName: ClassUserStatusEmpty,
			},
		}
	}
	v1 := &TLUserStatusEmpty{Data2: v}
	v1.Data2.ClassName = ClassUserStatusEmpty
	return v1
}

func NewTLUserStatusOnline(v *UserStatus) *TLUserStatusOnline {
	if v == nil {
		return &TLUserStatusOnline{
			Data2: &UserStatus{
				Cmd:       Cmd_UserStatusOnline,
				ClassName: ClassUserStatusOnline,
			},
		}
	}
	v1 := &TLUserStatusOnline{Data2: v}
	v1.Data2.ClassName = ClassUserStatusOnline
	return v1
}

func NewTLUserStatusOffline(v *UserStatus) *TLUserStatusOffline {
	if v == nil {
		return &TLUserStatusOffline{
			Data2: &UserStatus{
				Cmd:       Cmd_UserStatusOffline,
				ClassName: ClassUserStatusOffline,
			},
		}
	}
	v1 := &TLUserStatusOffline{Data2: v}
	v1.Data2.ClassName = ClassUserStatusOffline
	return v1
}

func NewTLUserStatusRecently(v *UserStatus) *TLUserStatusRecently {
	if v == nil {
		return &TLUserStatusRecently{
			Data2: &UserStatus{
				Cmd:       Cmd_UserStatusRecently,
				ClassName: ClassUserStatusRecently,
			},
		}
	}
	v1 := &TLUserStatusRecently{Data2: v}
	v1.Data2.ClassName = ClassUserStatusRecently
	return v1
}

func NewTLUserStatusLastWeek(v *UserStatus) *TLUserStatusLastWeek {
	if v == nil {
		return &TLUserStatusLastWeek{
			Data2: &UserStatus{
				Cmd:       Cmd_UserStatusLastWeek,
				ClassName: ClassUserStatusLastWeek,
			},
		}
	}
	v1 := &TLUserStatusLastWeek{Data2: v}
	v1.Data2.ClassName = ClassUserStatusLastWeek
	return v1
}

func NewTLUserStatusLastMonth(v *UserStatus) *TLUserStatusLastMonth {
	if v == nil {
		return &TLUserStatusLastMonth{
			Data2: &UserStatus{
				Cmd:       Cmd_UserStatusLastMonth,
				ClassName: ClassUserStatusLastMonth,
			},
		}
	}
	v1 := &TLUserStatusLastMonth{Data2: v}
	v1.Data2.ClassName = ClassUserStatusLastMonth
	return v1
}

func NewTLVideoSize(v *VideoSize) *TLVideoSize {
	if v == nil {
		return &TLVideoSize{
			Data2: &VideoSize{
				Cmd:       Cmd_VideoSize,
				ClassName: ClassVideoSize,
			},
		}
	}
	v1 := &TLVideoSize{Data2: v}
	v1.Data2.ClassName = ClassVideoSize
	return v1
}

func NewTLWallPaper(v *WallPaper) *TLWallPaper {
	if v == nil {
		return &TLWallPaper{
			Data2: &WallPaper{
				Cmd:       Cmd_WallPaper,
				ClassName: ClassWallPaper,
			},
		}
	}
	v1 := &TLWallPaper{Data2: v}
	v1.Data2.ClassName = ClassWallPaper
	return v1
}

func NewTLWallPaperSolid(v *WallPaper) *TLWallPaperSolid {
	if v == nil {
		return &TLWallPaperSolid{
			Data2: &WallPaper{
				Cmd:       Cmd_WallPaperSolid,
				ClassName: ClassWallPaperSolid,
			},
		}
	}
	v1 := &TLWallPaperSolid{Data2: v}
	v1.Data2.ClassName = ClassWallPaperSolid
	return v1
}

func NewTLWallPaperNoFile(v *WallPaper) *TLWallPaperNoFile {
	if v == nil {
		return &TLWallPaperNoFile{
			Data2: &WallPaper{
				Cmd:       Cmd_WallPaperNoFile,
				ClassName: ClassWallPaperNoFile,
			},
		}
	}
	v1 := &TLWallPaperNoFile{Data2: v}
	v1.Data2.ClassName = ClassWallPaperNoFile
	return v1
}

func NewTLWallPaperSettings(v *WallPaperSettings) *TLWallPaperSettings {
	if v == nil {
		return &TLWallPaperSettings{
			Data2: &WallPaperSettings{
				Cmd:       Cmd_WallPaperSettings,
				ClassName: ClassWallPaperSettings,
			},
		}
	}
	v1 := &TLWallPaperSettings{Data2: v}
	v1.Data2.ClassName = ClassWallPaperSettings
	return v1
}

func NewTLWalletSecretSalt(v *Wallet_KeySecretSalt) *TLWalletSecretSalt {
	if v == nil {
		return &TLWalletSecretSalt{
			Data2: &Wallet_KeySecretSalt{
				Cmd:       Cmd_WalletSecretSalt,
				ClassName: ClassWalletSecretSalt,
			},
		}
	}
	v1 := &TLWalletSecretSalt{Data2: v}
	v1.Data2.ClassName = ClassWalletSecretSalt
	return v1
}

func NewTLWalletLiteResponse(v *Wallet_LiteResponse) *TLWalletLiteResponse {
	if v == nil {
		return &TLWalletLiteResponse{
			Data2: &Wallet_LiteResponse{
				Cmd:       Cmd_WalletLiteResponse,
				ClassName: ClassWalletLiteResponse,
			},
		}
	}
	v1 := &TLWalletLiteResponse{Data2: v}
	v1.Data2.ClassName = ClassWalletLiteResponse
	return v1
}

func NewTLWebAuthorization(v *WebAuthorization) *TLWebAuthorization {
	if v == nil {
		return &TLWebAuthorization{
			Data2: &WebAuthorization{
				Cmd:       Cmd_WebAuthorization,
				ClassName: ClassWebAuthorization,
			},
		}
	}
	v1 := &TLWebAuthorization{Data2: v}
	v1.Data2.ClassName = ClassWebAuthorization
	return v1
}

func NewTLWebDocument(v *WebDocument) *TLWebDocument {
	if v == nil {
		return &TLWebDocument{
			Data2: &WebDocument{
				Cmd:       Cmd_WebDocument,
				ClassName: ClassWebDocument,
			},
		}
	}
	v1 := &TLWebDocument{Data2: v}
	v1.Data2.ClassName = ClassWebDocument
	return v1
}

func NewTLWebDocumentNoProxy(v *WebDocument) *TLWebDocumentNoProxy {
	if v == nil {
		return &TLWebDocumentNoProxy{
			Data2: &WebDocument{
				Cmd:       Cmd_WebDocumentNoProxy,
				ClassName: ClassWebDocumentNoProxy,
			},
		}
	}
	v1 := &TLWebDocumentNoProxy{Data2: v}
	v1.Data2.ClassName = ClassWebDocumentNoProxy
	return v1
}

func NewTLWebPageEmpty(v *WebPage) *TLWebPageEmpty {
	if v == nil {
		return &TLWebPageEmpty{
			Data2: &WebPage{
				Cmd:       Cmd_WebPageEmpty,
				ClassName: ClassWebPageEmpty,
			},
		}
	}
	v1 := &TLWebPageEmpty{Data2: v}
	v1.Data2.ClassName = ClassWebPageEmpty
	return v1
}

func NewTLWebPagePending(v *WebPage) *TLWebPagePending {
	if v == nil {
		return &TLWebPagePending{
			Data2: &WebPage{
				Cmd:       Cmd_WebPagePending,
				ClassName: ClassWebPagePending,
			},
		}
	}
	v1 := &TLWebPagePending{Data2: v}
	v1.Data2.ClassName = ClassWebPagePending
	return v1
}

func NewTLWebPage(v *WebPage) *TLWebPage {
	if v == nil {
		return &TLWebPage{
			Data2: &WebPage{
				Cmd:       Cmd_WebPage,
				ClassName: ClassWebPage,
			},
		}
	}
	v1 := &TLWebPage{Data2: v}
	v1.Data2.ClassName = ClassWebPage
	return v1
}

func NewTLWebPageNotModified(v *WebPage) *TLWebPageNotModified {
	if v == nil {
		return &TLWebPageNotModified{
			Data2: &WebPage{
				Cmd:       Cmd_WebPageNotModified,
				ClassName: ClassWebPageNotModified,
			},
		}
	}
	v1 := &TLWebPageNotModified{Data2: v}
	v1.Data2.ClassName = ClassWebPageNotModified
	return v1
}

func NewTLWebPageAttributeTheme(v *WebPageAttribute) *TLWebPageAttributeTheme {
	if v == nil {
		return &TLWebPageAttributeTheme{
			Data2: &WebPageAttribute{
				Cmd:       Cmd_WebPageAttributeTheme,
				ClassName: ClassWebPageAttributeTheme,
			},
		}
	}
	v1 := &TLWebPageAttributeTheme{Data2: v}
	v1.Data2.ClassName = ClassWebPageAttributeTheme
	return v1
}

func NewTLAccountCreateTheme8432c21f() *TLAccountCreateTheme {
	return &TLAccountCreateTheme{
		Cmd:       Cmd_AccountCreateTheme8432c21f,
		ClassName: ClassAccountCreateTheme,
	}
	//return &TLAccountCreateTheme{}
}

func NewTLAccountGetPasswordSettings9cd4eaf9() *TLAccountGetPasswordSettings {
	return &TLAccountGetPasswordSettings{
		Cmd:       Cmd_AccountGetPasswordSettings9cd4eaf9,
		ClassName: ClassAccountGetPasswordSettings,
	}
	//return &TLAccountGetPasswordSettings{}
}

func NewTLAccountGetTmpPassword449e0b51() *TLAccountGetTmpPassword {
	return &TLAccountGetTmpPassword{
		Cmd:       Cmd_AccountGetTmpPassword449e0b51,
		ClassName: ClassAccountGetTmpPassword,
	}
	//return &TLAccountGetTmpPassword{}
}

func NewTLAccountGetWallPapersaabb1763() *TLAccountGetWallPapers {
	return &TLAccountGetWallPapers{
		Cmd:       Cmd_AccountGetWallPapersaabb1763,
		ClassName: ClassAccountGetWallPapers,
	}
	//return &TLAccountGetWallPapers{}
}

func NewTLAccountInstallTheme7ae43737() *TLAccountInstallTheme {
	return &TLAccountInstallTheme{
		Cmd:       Cmd_AccountInstallTheme7ae43737,
		ClassName: ClassAccountInstallTheme,
	}
	//return &TLAccountInstallTheme{}
}

func NewTLAccountInstallWallPaperfeed5769() *TLAccountInstallWallPaper {
	return &TLAccountInstallWallPaper{
		Cmd:       Cmd_AccountInstallWallPaperfeed5769,
		ClassName: ClassAccountInstallWallPaper,
	}
	//return &TLAccountInstallWallPaper{}
}

func NewTLAccountRegisterDevice68976c6f() *TLAccountRegisterDevice {
	return &TLAccountRegisterDevice{
		Cmd:       Cmd_AccountRegisterDevice68976c6f,
		ClassName: ClassAccountRegisterDevice,
	}
	//return &TLAccountRegisterDevice{}
}

func NewTLAccountRegisterDevice5cbea590() *TLAccountRegisterDevice {
	return &TLAccountRegisterDevice{
		Cmd:       Cmd_AccountRegisterDevice5cbea590,
		ClassName: ClassAccountRegisterDevice,
	}
	//return &TLAccountRegisterDevice{}
}

func NewTLAccountRegisterDevice446c712c() *TLAccountRegisterDevice {
	return &TLAccountRegisterDevice{
		Cmd:       Cmd_AccountRegisterDevice446c712c,
		ClassName: ClassAccountRegisterDevice,
	}
	//return &TLAccountRegisterDevice{}
}

func NewTLAccountSaveWallPaper6c5a5b37() *TLAccountSaveWallPaper {
	return &TLAccountSaveWallPaper{
		Cmd:       Cmd_AccountSaveWallPaper6c5a5b37,
		ClassName: ClassAccountSaveWallPaper,
	}
	//return &TLAccountSaveWallPaper{}
}

func NewTLAccountSendChangePhoneCode82574ae5() *TLAccountSendChangePhoneCode {
	return &TLAccountSendChangePhoneCode{
		Cmd:       Cmd_AccountSendChangePhoneCode82574ae5,
		ClassName: ClassAccountSendChangePhoneCode,
	}
	//return &TLAccountSendChangePhoneCode{}
}

func NewTLAccountSendConfirmPhoneCode1b3faa88() *TLAccountSendConfirmPhoneCode {
	return &TLAccountSendConfirmPhoneCode{
		Cmd:       Cmd_AccountSendConfirmPhoneCode1b3faa88,
		ClassName: ClassAccountSendConfirmPhoneCode,
	}
	//return &TLAccountSendConfirmPhoneCode{}
}

func NewTLAccountSendVerifyPhoneCodea5a356f9() *TLAccountSendVerifyPhoneCode {
	return &TLAccountSendVerifyPhoneCode{
		Cmd:       Cmd_AccountSendVerifyPhoneCodea5a356f9,
		ClassName: ClassAccountSendVerifyPhoneCode,
	}
	//return &TLAccountSendVerifyPhoneCode{}
}

func NewTLAccountUnregisterDevice3076c4bf() *TLAccountUnregisterDevice {
	return &TLAccountUnregisterDevice{
		Cmd:       Cmd_AccountUnregisterDevice3076c4bf,
		ClassName: ClassAccountUnregisterDevice,
	}
	//return &TLAccountUnregisterDevice{}
}

func NewTLAccountUpdatePasswordSettingsa59b102f() *TLAccountUpdatePasswordSettings {
	return &TLAccountUpdatePasswordSettings{
		Cmd:       Cmd_AccountUpdatePasswordSettingsa59b102f,
		ClassName: ClassAccountUpdatePasswordSettings,
	}
	//return &TLAccountUpdatePasswordSettings{}
}

func NewTLAccountUpdateTheme5cb367d5() *TLAccountUpdateTheme {
	return &TLAccountUpdateTheme{
		Cmd:       Cmd_AccountUpdateTheme5cb367d5,
		ClassName: ClassAccountUpdateTheme,
	}
	//return &TLAccountUpdateTheme{}
}

func NewTLAccountUpdateTheme3b8ea202() *TLAccountUpdateTheme {
	return &TLAccountUpdateTheme{
		Cmd:       Cmd_AccountUpdateTheme3b8ea202,
		ClassName: ClassAccountUpdateTheme,
	}
	//return &TLAccountUpdateTheme{}
}

func NewTLAccountUploadWallPaperdd853661() *TLAccountUploadWallPaper {
	return &TLAccountUploadWallPaper{
		Cmd:       Cmd_AccountUploadWallPaperdd853661,
		ClassName: ClassAccountUploadWallPaper,
	}
	//return &TLAccountUploadWallPaper{}
}

func NewTLAuthCheckPasswordd18b4d16() *TLAuthCheckPassword {
	return &TLAuthCheckPassword{
		Cmd:       Cmd_AuthCheckPasswordd18b4d16,
		ClassName: ClassAuthCheckPassword,
	}
	//return &TLAuthCheckPassword{}
}

func NewTLAuthSendCodea677244f() *TLAuthSendCode {
	return &TLAuthSendCode{
		Cmd:       Cmd_AuthSendCodea677244f,
		ClassName: ClassAuthSendCode,
	}
	//return &TLAuthSendCode{}
}

func NewTLAuthSendCodeccfd70cf() *TLAuthSendCode {
	return &TLAuthSendCode{
		Cmd:       Cmd_AuthSendCodeccfd70cf,
		ClassName: ClassAuthSendCode,
	}
	//return &TLAuthSendCode{}
}

func NewTLAuthSignUp80eee427() *TLAuthSignUp {
	return &TLAuthSignUp{
		Cmd:       Cmd_AuthSignUp80eee427,
		ClassName: ClassAuthSignUp,
	}
	//return &TLAuthSignUp{}
}

func NewTLChannelsCreateChannel3d5fb10f() *TLChannelsCreateChannel {
	return &TLChannelsCreateChannel{
		Cmd:       Cmd_ChannelsCreateChannel3d5fb10f,
		ClassName: ClassChannelsCreateChannel,
	}
	//return &TLChannelsCreateChannel{}
}

func NewTLChannelsEditAdmind33c8902() *TLChannelsEditAdmin {
	return &TLChannelsEditAdmin{
		Cmd:       Cmd_ChannelsEditAdmind33c8902,
		ClassName: ClassChannelsEditAdmin,
	}
	//return &TLChannelsEditAdmin{}
}

func NewTLChannelsEditAdmin70f893ba() *TLChannelsEditAdmin {
	return &TLChannelsEditAdmin{
		Cmd:       Cmd_ChannelsEditAdmin70f893ba,
		ClassName: ClassChannelsEditAdmin,
	}
	//return &TLChannelsEditAdmin{}
}

func NewTLChannelsEditAdmineb7611d0() *TLChannelsEditAdmin {
	return &TLChannelsEditAdmin{
		Cmd:       Cmd_ChannelsEditAdmineb7611d0,
		ClassName: ClassChannelsEditAdmin,
	}
	//return &TLChannelsEditAdmin{}
}

func NewTLChannelsEditBanned72796912() *TLChannelsEditBanned {
	return &TLChannelsEditBanned{
		Cmd:       Cmd_ChannelsEditBanned72796912,
		ClassName: ClassChannelsEditBanned,
	}
	//return &TLChannelsEditBanned{}
}

func NewTLChannelsExportMessageLinke63fadeb() *TLChannelsExportMessageLink {
	return &TLChannelsExportMessageLink{
		Cmd:       Cmd_ChannelsExportMessageLinke63fadeb,
		ClassName: ClassChannelsExportMessageLink,
	}
	//return &TLChannelsExportMessageLink{}
}

func NewTLChannelsExportMessageLinkceb77163() *TLChannelsExportMessageLink {
	return &TLChannelsExportMessageLink{
		Cmd:       Cmd_ChannelsExportMessageLinkceb77163,
		ClassName: ClassChannelsExportMessageLink,
	}
	//return &TLChannelsExportMessageLink{}
}

func NewTLChannelsGetAdminedPublicChannelsf8b036af() *TLChannelsGetAdminedPublicChannels {
	return &TLChannelsGetAdminedPublicChannels{
		Cmd:       Cmd_ChannelsGetAdminedPublicChannelsf8b036af,
		ClassName: ClassChannelsGetAdminedPublicChannels,
	}
	//return &TLChannelsGetAdminedPublicChannels{}
}

func NewTLChannelsGetMessagesad8c9a23() *TLChannelsGetMessages {
	return &TLChannelsGetMessages{
		Cmd:       Cmd_ChannelsGetMessagesad8c9a23,
		ClassName: ClassChannelsGetMessages,
	}
	//return &TLChannelsGetMessages{}
}

func NewTLChannelsGetParticipants123e05e9() *TLChannelsGetParticipants {
	return &TLChannelsGetParticipants{
		Cmd:       Cmd_ChannelsGetParticipants123e05e9,
		ClassName: ClassChannelsGetParticipants,
	}
	//return &TLChannelsGetParticipants{}
}

func NewTLContactsBlock68cc1411() *TLContactsBlock {
	return &TLContactsBlock{
		Cmd:       Cmd_ContactsBlock68cc1411,
		ClassName: ClassContactsBlock,
	}
	//return &TLContactsBlock{}
}

func NewTLContactsDeleteContacts96a0e00() *TLContactsDeleteContacts {
	return &TLContactsDeleteContacts{
		Cmd:       Cmd_ContactsDeleteContacts96a0e00,
		ClassName: ClassContactsDeleteContacts,
	}
	//return &TLContactsDeleteContacts{}
}

func NewTLContactsGetContacts22c6aa08() *TLContactsGetContacts {
	return &TLContactsGetContacts{
		Cmd:       Cmd_ContactsGetContacts22c6aa08,
		ClassName: ClassContactsGetContacts,
	}
	//return &TLContactsGetContacts{}
}

func NewTLContactsGetLocatedd348bc44() *TLContactsGetLocated {
	return &TLContactsGetLocated{
		Cmd:       Cmd_ContactsGetLocatedd348bc44,
		ClassName: ClassContactsGetLocated,
	}
	//return &TLContactsGetLocated{}
}

func NewTLContactsImportContactsda30b32d() *TLContactsImportContacts {
	return &TLContactsImportContacts{
		Cmd:       Cmd_ContactsImportContactsda30b32d,
		ClassName: ClassContactsImportContacts,
	}
	//return &TLContactsImportContacts{}
}

func NewTLContactsUnblockbea65d50() *TLContactsUnblock {
	return &TLContactsUnblock{
		Cmd:       Cmd_ContactsUnblockbea65d50,
		ClassName: ClassContactsUnblock,
	}
	//return &TLContactsUnblock{}
}

func NewTLHelpGetAppChangelog5bab7fb2() *TLHelpGetAppChangelog {
	return &TLHelpGetAppChangelog{
		Cmd:       Cmd_HelpGetAppChangelog5bab7fb2,
		ClassName: ClassHelpGetAppChangelog,
	}
	//return &TLHelpGetAppChangelog{}
}

func NewTLHelpGetAppUpdate522d5a7d() *TLHelpGetAppUpdate {
	return &TLHelpGetAppUpdate{
		Cmd:       Cmd_HelpGetAppUpdate522d5a7d,
		ClassName: ClassHelpGetAppUpdate,
	}
	//return &TLHelpGetAppUpdate{}
}

func NewTLHelpGetAppUpdatec812ac7e() *TLHelpGetAppUpdate {
	return &TLHelpGetAppUpdate{
		Cmd:       Cmd_HelpGetAppUpdatec812ac7e,
		ClassName: ClassHelpGetAppUpdate,
	}
	//return &TLHelpGetAppUpdate{}
}

func NewTLHelpGetInviteTexta4a95186() *TLHelpGetInviteText {
	return &TLHelpGetInviteText{
		Cmd:       Cmd_HelpGetInviteTexta4a95186,
		ClassName: ClassHelpGetInviteText,
	}
	//return &TLHelpGetInviteText{}
}

func NewTLHelpGetTermsOfService37d78f83() *TLHelpGetTermsOfService {
	return &TLHelpGetTermsOfService{
		Cmd:       Cmd_HelpGetTermsOfService37d78f83,
		ClassName: ClassHelpGetTermsOfService,
	}
	//return &TLHelpGetTermsOfService{}
}

func NewTLLangpackGetDifferencecd984aa5() *TLLangpackGetDifference {
	return &TLLangpackGetDifference{
		Cmd:       Cmd_LangpackGetDifferencecd984aa5,
		ClassName: ClassLangpackGetDifference,
	}
	//return &TLLangpackGetDifference{}
}

func NewTLLangpackGetDifference9d51e814() *TLLangpackGetDifference {
	return &TLLangpackGetDifference{
		Cmd:       Cmd_LangpackGetDifference9d51e814,
		ClassName: ClassLangpackGetDifference,
	}
	//return &TLLangpackGetDifference{}
}

func NewTLLangpackGetLangPackf2f2330a() *TLLangpackGetLangPack {
	return &TLLangpackGetLangPack{
		Cmd:       Cmd_LangpackGetLangPackf2f2330a,
		ClassName: ClassLangpackGetLangPack,
	}
	//return &TLLangpackGetLangPack{}
}

func NewTLLangpackGetLanguages42c6978f() *TLLangpackGetLanguages {
	return &TLLangpackGetLanguages{
		Cmd:       Cmd_LangpackGetLanguages42c6978f,
		ClassName: ClassLangpackGetLanguages,
	}
	//return &TLLangpackGetLanguages{}
}

func NewTLLangpackGetStringsefea3803() *TLLangpackGetStrings {
	return &TLLangpackGetStrings{
		Cmd:       Cmd_LangpackGetStringsefea3803,
		ClassName: ClassLangpackGetStrings,
	}
	//return &TLLangpackGetStrings{}
}

func NewTLMessagesDeleteHistory1c015b09() *TLMessagesDeleteHistory {
	return &TLMessagesDeleteHistory{
		Cmd:       Cmd_MessagesDeleteHistory1c015b09,
		ClassName: ClassMessagesDeleteHistory,
	}
	//return &TLMessagesDeleteHistory{}
}

func NewTLMessagesDeleteMessagesa5f18925() *TLMessagesDeleteMessages {
	return &TLMessagesDeleteMessages{
		Cmd:       Cmd_MessagesDeleteMessagesa5f18925,
		ClassName: ClassMessagesDeleteMessages,
	}
	//return &TLMessagesDeleteMessages{}
}

func NewTLMessagesEditInlineBotMessage83557dba() *TLMessagesEditInlineBotMessage {
	return &TLMessagesEditInlineBotMessage{
		Cmd:       Cmd_MessagesEditInlineBotMessage83557dba,
		ClassName: ClassMessagesEditInlineBotMessage,
	}
	//return &TLMessagesEditInlineBotMessage{}
}

func NewTLMessagesEditInlineBotMessageadc3e828() *TLMessagesEditInlineBotMessage {
	return &TLMessagesEditInlineBotMessage{
		Cmd:       Cmd_MessagesEditInlineBotMessageadc3e828,
		ClassName: ClassMessagesEditInlineBotMessage,
	}
	//return &TLMessagesEditInlineBotMessage{}
}

func NewTLMessagesEditMessage48f71778() *TLMessagesEditMessage {
	return &TLMessagesEditMessage{
		Cmd:       Cmd_MessagesEditMessage48f71778,
		ClassName: ClassMessagesEditMessage,
	}
	//return &TLMessagesEditMessage{}
}

func NewTLMessagesEditMessaged116f31e() *TLMessagesEditMessage {
	return &TLMessagesEditMessage{
		Cmd:       Cmd_MessagesEditMessaged116f31e,
		ClassName: ClassMessagesEditMessage,
	}
	//return &TLMessagesEditMessage{}
}

func NewTLMessagesEditMessagec000e4c8() *TLMessagesEditMessage {
	return &TLMessagesEditMessage{
		Cmd:       Cmd_MessagesEditMessagec000e4c8,
		ClassName: ClassMessagesEditMessage,
	}
	//return &TLMessagesEditMessage{}
}

func NewTLMessagesExportChatInvitedf7534c() *TLMessagesExportChatInvite {
	return &TLMessagesExportChatInvite{
		Cmd:       Cmd_MessagesExportChatInvitedf7534c,
		ClassName: ClassMessagesExportChatInvite,
	}
	//return &TLMessagesExportChatInvite{}
}

func NewTLMessagesForwardMessagesd9fee60e() *TLMessagesForwardMessages {
	return &TLMessagesForwardMessages{
		Cmd:       Cmd_MessagesForwardMessagesd9fee60e,
		ClassName: ClassMessagesForwardMessages,
	}
	//return &TLMessagesForwardMessages{}
}

func NewTLMessagesGetBotCallbackAnswer9342ca07() *TLMessagesGetBotCallbackAnswer {
	return &TLMessagesGetBotCallbackAnswer{
		Cmd:       Cmd_MessagesGetBotCallbackAnswer9342ca07,
		ClassName: ClassMessagesGetBotCallbackAnswer,
	}
	//return &TLMessagesGetBotCallbackAnswer{}
}

func NewTLMessagesGetBotCallbackAnswera6e94f04() *TLMessagesGetBotCallbackAnswer {
	return &TLMessagesGetBotCallbackAnswer{
		Cmd:       Cmd_MessagesGetBotCallbackAnswera6e94f04,
		ClassName: ClassMessagesGetBotCallbackAnswer,
	}
	//return &TLMessagesGetBotCallbackAnswer{}
}

func NewTLMessagesGetDialogsa0ee3b73() *TLMessagesGetDialogs {
	return &TLMessagesGetDialogs{
		Cmd:       Cmd_MessagesGetDialogsa0ee3b73,
		ClassName: ClassMessagesGetDialogs,
	}
	//return &TLMessagesGetDialogs{}
}

func NewTLMessagesGetDialogsb098aee6() *TLMessagesGetDialogs {
	return &TLMessagesGetDialogs{
		Cmd:       Cmd_MessagesGetDialogsb098aee6,
		ClassName: ClassMessagesGetDialogs,
	}
	//return &TLMessagesGetDialogs{}
}

func NewTLMessagesGetDialogs6b47f94d() *TLMessagesGetDialogs {
	return &TLMessagesGetDialogs{
		Cmd:       Cmd_MessagesGetDialogs6b47f94d,
		ClassName: ClassMessagesGetDialogs,
	}
	//return &TLMessagesGetDialogs{}
}

func NewTLMessagesGetHistorydcbb8260() *TLMessagesGetHistory {
	return &TLMessagesGetHistory{
		Cmd:       Cmd_MessagesGetHistorydcbb8260,
		ClassName: ClassMessagesGetHistory,
	}
	//return &TLMessagesGetHistory{}
}

func NewTLMessagesGetMessages63c66506() *TLMessagesGetMessages {
	return &TLMessagesGetMessages{
		Cmd:       Cmd_MessagesGetMessages63c66506,
		ClassName: ClassMessagesGetMessages,
	}
	//return &TLMessagesGetMessages{}
}

func NewTLMessagesGetMessagesViews5784d3e1() *TLMessagesGetMessagesViews {
	return &TLMessagesGetMessagesViews{
		Cmd:       Cmd_MessagesGetMessagesViews5784d3e1,
		ClassName: ClassMessagesGetMessagesViews,
	}
	//return &TLMessagesGetMessagesViews{}
}

func NewTLMessagesGetPeerDialogse470bcfd() *TLMessagesGetPeerDialogs {
	return &TLMessagesGetPeerDialogs{
		Cmd:       Cmd_MessagesGetPeerDialogse470bcfd,
		ClassName: ClassMessagesGetPeerDialogs,
	}
	//return &TLMessagesGetPeerDialogs{}
}

func NewTLMessagesGetPinnedDialogsd6b94df2() *TLMessagesGetPinnedDialogs {
	return &TLMessagesGetPinnedDialogs{
		Cmd:       Cmd_MessagesGetPinnedDialogsd6b94df2,
		ClassName: ClassMessagesGetPinnedDialogs,
	}
	//return &TLMessagesGetPinnedDialogs{}
}

func NewTLMessagesGetRecentLocationsbbc45b09() *TLMessagesGetRecentLocations {
	return &TLMessagesGetRecentLocations{
		Cmd:       Cmd_MessagesGetRecentLocationsbbc45b09,
		ClassName: ClassMessagesGetRecentLocations,
	}
	//return &TLMessagesGetRecentLocations{}
}

func NewTLMessagesGetStatsURL812c2ae6() *TLMessagesGetStatsURL {
	return &TLMessagesGetStatsURL{
		Cmd:       Cmd_MessagesGetStatsURL812c2ae6,
		ClassName: ClassMessagesGetStatsURL,
	}
	//return &TLMessagesGetStatsURL{}
}

func NewTLMessagesGetStickers43d4f2c() *TLMessagesGetStickers {
	return &TLMessagesGetStickers{
		Cmd:       Cmd_MessagesGetStickers43d4f2c,
		ClassName: ClassMessagesGetStickers,
	}
	//return &TLMessagesGetStickers{}
}

func NewTLMessagesGetWebPagePreview8b68b0cc() *TLMessagesGetWebPagePreview {
	return &TLMessagesGetWebPagePreview{
		Cmd:       Cmd_MessagesGetWebPagePreview8b68b0cc,
		ClassName: ClassMessagesGetWebPagePreview,
	}
	//return &TLMessagesGetWebPagePreview{}
}

func NewTLMessagesInstallStickerSet7b30c3a6() *TLMessagesInstallStickerSet {
	return &TLMessagesInstallStickerSet{
		Cmd:       Cmd_MessagesInstallStickerSet7b30c3a6,
		ClassName: ClassMessagesInstallStickerSet,
	}
	//return &TLMessagesInstallStickerSet{}
}

func NewTLMessagesReadHistoryb04f2510() *TLMessagesReadHistory {
	return &TLMessagesReadHistory{
		Cmd:       Cmd_MessagesReadHistoryb04f2510,
		ClassName: ClassMessagesReadHistory,
	}
	//return &TLMessagesReadHistory{}
}

func NewTLMessagesReorderPinnedDialogs3b1adf37() *TLMessagesReorderPinnedDialogs {
	return &TLMessagesReorderPinnedDialogs{
		Cmd:       Cmd_MessagesReorderPinnedDialogs3b1adf37,
		ClassName: ClassMessagesReorderPinnedDialogs,
	}
	//return &TLMessagesReorderPinnedDialogs{}
}

func NewTLMessagesReorderPinnedDialogs5b51d63f() *TLMessagesReorderPinnedDialogs {
	return &TLMessagesReorderPinnedDialogs{
		Cmd:       Cmd_MessagesReorderPinnedDialogs5b51d63f,
		ClassName: ClassMessagesReorderPinnedDialogs,
	}
	//return &TLMessagesReorderPinnedDialogs{}
}

func NewTLMessagesReorderStickerSets9fcfbc30() *TLMessagesReorderStickerSets {
	return &TLMessagesReorderStickerSets{
		Cmd:       Cmd_MessagesReorderStickerSets9fcfbc30,
		ClassName: ClassMessagesReorderStickerSets,
	}
	//return &TLMessagesReorderStickerSets{}
}

func NewTLMessagesSaveRecentSticker348e39bf() *TLMessagesSaveRecentSticker {
	return &TLMessagesSaveRecentSticker{
		Cmd:       Cmd_MessagesSaveRecentSticker348e39bf,
		ClassName: ClassMessagesSaveRecentSticker,
	}
	//return &TLMessagesSaveRecentSticker{}
}

func NewTLMessagesSearchc352eec() *TLMessagesSearch {
	return &TLMessagesSearch{
		Cmd:       Cmd_MessagesSearchc352eec,
		ClassName: ClassMessagesSearch,
	}
	//return &TLMessagesSearch{}
}

func NewTLMessagesSearch4e17810b() *TLMessagesSearch {
	return &TLMessagesSearch{
		Cmd:       Cmd_MessagesSearch4e17810b,
		ClassName: ClassMessagesSearch,
	}
	//return &TLMessagesSearch{}
}

func NewTLMessagesSearch8614ef68() *TLMessagesSearch {
	return &TLMessagesSearch{
		Cmd:       Cmd_MessagesSearch8614ef68,
		ClassName: ClassMessagesSearch,
	}
	//return &TLMessagesSearch{}
}

func NewTLMessagesSearchd4569248() *TLMessagesSearch {
	return &TLMessagesSearch{
		Cmd:       Cmd_MessagesSearchd4569248,
		ClassName: ClassMessagesSearch,
	}
	//return &TLMessagesSearch{}
}

func NewTLMessagesSearchf288a275() *TLMessagesSearch {
	return &TLMessagesSearch{
		Cmd:       Cmd_MessagesSearchf288a275,
		ClassName: ClassMessagesSearch,
	}
	//return &TLMessagesSearch{}
}

func NewTLMessagesSearchGlobal4bc6589a() *TLMessagesSearchGlobal {
	return &TLMessagesSearchGlobal{
		Cmd:       Cmd_MessagesSearchGlobal4bc6589a,
		ClassName: ClassMessagesSearchGlobal,
	}
	//return &TLMessagesSearchGlobal{}
}

func NewTLMessagesSearchGlobalbf7225a4() *TLMessagesSearchGlobal {
	return &TLMessagesSearchGlobal{
		Cmd:       Cmd_MessagesSearchGlobalbf7225a4,
		ClassName: ClassMessagesSearchGlobal,
	}
	//return &TLMessagesSearchGlobal{}
}

func NewTLMessagesSearchGlobalf79c611() *TLMessagesSearchGlobal {
	return &TLMessagesSearchGlobal{
		Cmd:       Cmd_MessagesSearchGlobalf79c611,
		ClassName: ClassMessagesSearchGlobal,
	}
	//return &TLMessagesSearchGlobal{}
}

func NewTLMessagesSendEncrypted44fa7a15() *TLMessagesSendEncrypted {
	return &TLMessagesSendEncrypted{
		Cmd:       Cmd_MessagesSendEncrypted44fa7a15,
		ClassName: ClassMessagesSendEncrypted,
	}
	//return &TLMessagesSendEncrypted{}
}

func NewTLMessagesSendEncryptedFile5559481d() *TLMessagesSendEncryptedFile {
	return &TLMessagesSendEncryptedFile{
		Cmd:       Cmd_MessagesSendEncryptedFile5559481d,
		ClassName: ClassMessagesSendEncryptedFile,
	}
	//return &TLMessagesSendEncryptedFile{}
}

func NewTLMessagesSendInlineBotResult220815b0() *TLMessagesSendInlineBotResult {
	return &TLMessagesSendInlineBotResult{
		Cmd:       Cmd_MessagesSendInlineBotResult220815b0,
		ClassName: ClassMessagesSendInlineBotResult,
	}
	//return &TLMessagesSendInlineBotResult{}
}

func NewTLMessagesSendMedia3491eba9() *TLMessagesSendMedia {
	return &TLMessagesSendMedia{
		Cmd:       Cmd_MessagesSendMedia3491eba9,
		ClassName: ClassMessagesSendMedia,
	}
	//return &TLMessagesSendMedia{}
}

func NewTLMessagesSendMediab8d1262b() *TLMessagesSendMedia {
	return &TLMessagesSendMedia{
		Cmd:       Cmd_MessagesSendMediab8d1262b,
		ClassName: ClassMessagesSendMedia,
	}
	//return &TLMessagesSendMedia{}
}

func NewTLMessagesSendMessage520c3870() *TLMessagesSendMessage {
	return &TLMessagesSendMessage{
		Cmd:       Cmd_MessagesSendMessage520c3870,
		ClassName: ClassMessagesSendMessage,
	}
	//return &TLMessagesSendMessage{}
}

func NewTLMessagesSendMultiMediacc0110cb() *TLMessagesSendMultiMedia {
	return &TLMessagesSendMultiMedia{
		Cmd:       Cmd_MessagesSendMultiMediacc0110cb,
		ClassName: ClassMessagesSendMultiMedia,
	}
	//return &TLMessagesSendMultiMedia{}
}

func NewTLMessagesSetBotCallbackAnswer481c591a() *TLMessagesSetBotCallbackAnswer {
	return &TLMessagesSetBotCallbackAnswer{
		Cmd:       Cmd_MessagesSetBotCallbackAnswer481c591a,
		ClassName: ClassMessagesSetBotCallbackAnswer,
	}
	//return &TLMessagesSetBotCallbackAnswer{}
}

func NewTLMessagesSetTyping58943ee2() *TLMessagesSetTyping {
	return &TLMessagesSetTyping{
		Cmd:       Cmd_MessagesSetTyping58943ee2,
		ClassName: ClassMessagesSetTyping,
	}
	//return &TLMessagesSetTyping{}
}

func NewTLMessagesToggleDialogPina731e257() *TLMessagesToggleDialogPin {
	return &TLMessagesToggleDialogPin{
		Cmd:       Cmd_MessagesToggleDialogPina731e257,
		ClassName: ClassMessagesToggleDialogPin,
	}
	//return &TLMessagesToggleDialogPin{}
}

func NewTLPhoneDiscardCallb2cbc1c0() *TLPhoneDiscardCall {
	return &TLPhoneDiscardCall{
		Cmd:       Cmd_PhoneDiscardCallb2cbc1c0,
		ClassName: ClassPhoneDiscardCall,
	}
	//return &TLPhoneDiscardCall{}
}

func NewTLPhoneRequestCall42ff96ed() *TLPhoneRequestCall {
	return &TLPhoneRequestCall{
		Cmd:       Cmd_PhoneRequestCall42ff96ed,
		ClassName: ClassPhoneRequestCall,
	}
	//return &TLPhoneRequestCall{}
}

func NewTLPhoneSetCallRating59ead627() *TLPhoneSetCallRating {
	return &TLPhoneSetCallRating{
		Cmd:       Cmd_PhoneSetCallRating59ead627,
		ClassName: ClassPhoneSetCallRating,
	}
	//return &TLPhoneSetCallRating{}
}

func NewTLPhotosUpdateProfilePhoto72d4742c() *TLPhotosUpdateProfilePhoto {
	return &TLPhotosUpdateProfilePhoto{
		Cmd:       Cmd_PhotosUpdateProfilePhoto72d4742c,
		ClassName: ClassPhotosUpdateProfilePhoto,
	}
	//return &TLPhotosUpdateProfilePhoto{}
}

func NewTLPhotosUpdateProfilePhotoeef579a0() *TLPhotosUpdateProfilePhoto {
	return &TLPhotosUpdateProfilePhoto{
		Cmd:       Cmd_PhotosUpdateProfilePhotoeef579a0,
		ClassName: ClassPhotosUpdateProfilePhoto,
	}
	//return &TLPhotosUpdateProfilePhoto{}
}

func NewTLPhotosUploadProfilePhoto89f30f69() *TLPhotosUploadProfilePhoto {
	return &TLPhotosUploadProfilePhoto{
		Cmd:       Cmd_PhotosUploadProfilePhoto89f30f69,
		ClassName: ClassPhotosUploadProfilePhoto,
	}
	//return &TLPhotosUploadProfilePhoto{}
}

func NewTLPhotosUploadProfilePhotod50f9c88() *TLPhotosUploadProfilePhoto {
	return &TLPhotosUploadProfilePhoto{
		Cmd:       Cmd_PhotosUploadProfilePhotod50f9c88,
		ClassName: ClassPhotosUploadProfilePhoto,
	}
	//return &TLPhotosUploadProfilePhoto{}
}

func NewTLStatsGetBroadcastStatse6300dba() *TLStatsGetBroadcastStats {
	return &TLStatsGetBroadcastStats{
		Cmd:       Cmd_StatsGetBroadcastStatse6300dba,
		ClassName: ClassStatsGetBroadcastStats,
	}
	//return &TLStatsGetBroadcastStats{}
}

func NewTLStickersCreateStickerSetf1036780() *TLStickersCreateStickerSet {
	return &TLStickersCreateStickerSet{
		Cmd:       Cmd_StickersCreateStickerSetf1036780,
		ClassName: ClassStickersCreateStickerSet,
	}
	//return &TLStickersCreateStickerSet{}
}

func NewTLUpdatesGetChannelDifferencebb32d7c0() *TLUpdatesGetChannelDifference {
	return &TLUpdatesGetChannelDifference{
		Cmd:       Cmd_UpdatesGetChannelDifferencebb32d7c0,
		ClassName: ClassUpdatesGetChannelDifference,
	}
	//return &TLUpdatesGetChannelDifference{}
}

func NewTLUpdatesGetDifferencea041495() *TLUpdatesGetDifference {
	return &TLUpdatesGetDifference{
		Cmd:       Cmd_UpdatesGetDifferencea041495,
		ClassName: ClassUpdatesGetDifference,
	}
	//return &TLUpdatesGetDifference{}
}

func NewTLUploadGetCdnFileHashes4da54231() *TLUploadGetCdnFileHashes {
	return &TLUploadGetCdnFileHashes{
		Cmd:       Cmd_UploadGetCdnFileHashes4da54231,
		ClassName: ClassUploadGetCdnFileHashes,
	}
	//return &TLUploadGetCdnFileHashes{}
}

func NewTLUploadGetFileb15a9afc() *TLUploadGetFile {
	return &TLUploadGetFile{
		Cmd:       Cmd_UploadGetFileb15a9afc,
		ClassName: ClassUploadGetFile,
	}
	//return &TLUploadGetFile{}
}

func NewTLUploadReuploadCdnFile9b2754a8() *TLUploadReuploadCdnFile {
	return &TLUploadReuploadCdnFile{
		Cmd:       Cmd_UploadReuploadCdnFile9b2754a8,
		ClassName: ClassUploadReuploadCdnFile,
	}
	//return &TLUploadReuploadCdnFile{}
}

func NewTLAccountAcceptAuthorization() *TLAccountAcceptAuthorization {
	return &TLAccountAcceptAuthorization{
		Cmd:       Cmd_AccountAcceptAuthorization,
		ClassName: ClassAccountAcceptAuthorization,
	}
	//return &TLAccountAcceptAuthorization{}
}

func NewTLAccountCancelPasswordEmail() *TLAccountCancelPasswordEmail {
	return &TLAccountCancelPasswordEmail{
		Cmd:       Cmd_AccountCancelPasswordEmail,
		ClassName: ClassAccountCancelPasswordEmail,
	}
	//return &TLAccountCancelPasswordEmail{}
}

func NewTLAccountChangePhone() *TLAccountChangePhone {
	return &TLAccountChangePhone{
		Cmd:       Cmd_AccountChangePhone,
		ClassName: ClassAccountChangePhone,
	}
	//return &TLAccountChangePhone{}
}

func NewTLAccountCheckUsername() *TLAccountCheckUsername {
	return &TLAccountCheckUsername{
		Cmd:       Cmd_AccountCheckUsername,
		ClassName: ClassAccountCheckUsername,
	}
	//return &TLAccountCheckUsername{}
}

func NewTLAccountConfirmPasswordEmail() *TLAccountConfirmPasswordEmail {
	return &TLAccountConfirmPasswordEmail{
		Cmd:       Cmd_AccountConfirmPasswordEmail,
		ClassName: ClassAccountConfirmPasswordEmail,
	}
	//return &TLAccountConfirmPasswordEmail{}
}

func NewTLAccountConfirmPhone() *TLAccountConfirmPhone {
	return &TLAccountConfirmPhone{
		Cmd:       Cmd_AccountConfirmPhone,
		ClassName: ClassAccountConfirmPhone,
	}
	//return &TLAccountConfirmPhone{}
}

func NewTLAccountCreateTheme() *TLAccountCreateTheme {
	return &TLAccountCreateTheme{
		Cmd:       Cmd_AccountCreateTheme,
		ClassName: ClassAccountCreateTheme,
	}
	//return &TLAccountCreateTheme{}
}

func NewTLAccountDeleteAccount() *TLAccountDeleteAccount {
	return &TLAccountDeleteAccount{
		Cmd:       Cmd_AccountDeleteAccount,
		ClassName: ClassAccountDeleteAccount,
	}
	//return &TLAccountDeleteAccount{}
}

func NewTLAccountDeleteSecureValue() *TLAccountDeleteSecureValue {
	return &TLAccountDeleteSecureValue{
		Cmd:       Cmd_AccountDeleteSecureValue,
		ClassName: ClassAccountDeleteSecureValue,
	}
	//return &TLAccountDeleteSecureValue{}
}

func NewTLAccountFinishTakeoutSession() *TLAccountFinishTakeoutSession {
	return &TLAccountFinishTakeoutSession{
		Cmd:       Cmd_AccountFinishTakeoutSession,
		ClassName: ClassAccountFinishTakeoutSession,
	}
	//return &TLAccountFinishTakeoutSession{}
}

func NewTLAccountGetAccountTTL() *TLAccountGetAccountTTL {
	return &TLAccountGetAccountTTL{
		Cmd:       Cmd_AccountGetAccountTTL,
		ClassName: ClassAccountGetAccountTTL,
	}
	//return &TLAccountGetAccountTTL{}
}

func NewTLAccountGetAllSecureValues() *TLAccountGetAllSecureValues {
	return &TLAccountGetAllSecureValues{
		Cmd:       Cmd_AccountGetAllSecureValues,
		ClassName: ClassAccountGetAllSecureValues,
	}
	//return &TLAccountGetAllSecureValues{}
}

func NewTLAccountGetAuthorizationForm() *TLAccountGetAuthorizationForm {
	return &TLAccountGetAuthorizationForm{
		Cmd:       Cmd_AccountGetAuthorizationForm,
		ClassName: ClassAccountGetAuthorizationForm,
	}
	//return &TLAccountGetAuthorizationForm{}
}

func NewTLAccountGetAuthorizations() *TLAccountGetAuthorizations {
	return &TLAccountGetAuthorizations{
		Cmd:       Cmd_AccountGetAuthorizations,
		ClassName: ClassAccountGetAuthorizations,
	}
	//return &TLAccountGetAuthorizations{}
}

func NewTLAccountGetAutoDownloadSettings() *TLAccountGetAutoDownloadSettings {
	return &TLAccountGetAutoDownloadSettings{
		Cmd:       Cmd_AccountGetAutoDownloadSettings,
		ClassName: ClassAccountGetAutoDownloadSettings,
	}
	//return &TLAccountGetAutoDownloadSettings{}
}

func NewTLAccountGetContactSignUpNotification() *TLAccountGetContactSignUpNotification {
	return &TLAccountGetContactSignUpNotification{
		Cmd:       Cmd_AccountGetContactSignUpNotification,
		ClassName: ClassAccountGetContactSignUpNotification,
	}
	//return &TLAccountGetContactSignUpNotification{}
}

func NewTLAccountGetContentSettings() *TLAccountGetContentSettings {
	return &TLAccountGetContentSettings{
		Cmd:       Cmd_AccountGetContentSettings,
		ClassName: ClassAccountGetContentSettings,
	}
	//return &TLAccountGetContentSettings{}
}

func NewTLAccountGetGlobalPrivacySettings() *TLAccountGetGlobalPrivacySettings {
	return &TLAccountGetGlobalPrivacySettings{
		Cmd:       Cmd_AccountGetGlobalPrivacySettings,
		ClassName: ClassAccountGetGlobalPrivacySettings,
	}
	//return &TLAccountGetGlobalPrivacySettings{}
}

func NewTLAccountGetMultiWallPapers() *TLAccountGetMultiWallPapers {
	return &TLAccountGetMultiWallPapers{
		Cmd:       Cmd_AccountGetMultiWallPapers,
		ClassName: ClassAccountGetMultiWallPapers,
	}
	//return &TLAccountGetMultiWallPapers{}
}

func NewTLAccountGetNotifyExceptions() *TLAccountGetNotifyExceptions {
	return &TLAccountGetNotifyExceptions{
		Cmd:       Cmd_AccountGetNotifyExceptions,
		ClassName: ClassAccountGetNotifyExceptions,
	}
	//return &TLAccountGetNotifyExceptions{}
}

func NewTLAccountGetNotifySettings() *TLAccountGetNotifySettings {
	return &TLAccountGetNotifySettings{
		Cmd:       Cmd_AccountGetNotifySettings,
		ClassName: ClassAccountGetNotifySettings,
	}
	//return &TLAccountGetNotifySettings{}
}

func NewTLAccountGetPassword() *TLAccountGetPassword {
	return &TLAccountGetPassword{
		Cmd:       Cmd_AccountGetPassword,
		ClassName: ClassAccountGetPassword,
	}
	//return &TLAccountGetPassword{}
}

func NewTLAccountGetPasswordSettings() *TLAccountGetPasswordSettings {
	return &TLAccountGetPasswordSettings{
		Cmd:       Cmd_AccountGetPasswordSettings,
		ClassName: ClassAccountGetPasswordSettings,
	}
	//return &TLAccountGetPasswordSettings{}
}

func NewTLAccountGetPrivacy() *TLAccountGetPrivacy {
	return &TLAccountGetPrivacy{
		Cmd:       Cmd_AccountGetPrivacy,
		ClassName: ClassAccountGetPrivacy,
	}
	//return &TLAccountGetPrivacy{}
}

func NewTLAccountGetSecureValue() *TLAccountGetSecureValue {
	return &TLAccountGetSecureValue{
		Cmd:       Cmd_AccountGetSecureValue,
		ClassName: ClassAccountGetSecureValue,
	}
	//return &TLAccountGetSecureValue{}
}

func NewTLAccountGetTheme() *TLAccountGetTheme {
	return &TLAccountGetTheme{
		Cmd:       Cmd_AccountGetTheme,
		ClassName: ClassAccountGetTheme,
	}
	//return &TLAccountGetTheme{}
}

func NewTLAccountGetThemes() *TLAccountGetThemes {
	return &TLAccountGetThemes{
		Cmd:       Cmd_AccountGetThemes,
		ClassName: ClassAccountGetThemes,
	}
	//return &TLAccountGetThemes{}
}

func NewTLAccountGetTmpPassword() *TLAccountGetTmpPassword {
	return &TLAccountGetTmpPassword{
		Cmd:       Cmd_AccountGetTmpPassword,
		ClassName: ClassAccountGetTmpPassword,
	}
	//return &TLAccountGetTmpPassword{}
}

func NewTLAccountGetWallPaper() *TLAccountGetWallPaper {
	return &TLAccountGetWallPaper{
		Cmd:       Cmd_AccountGetWallPaper,
		ClassName: ClassAccountGetWallPaper,
	}
	//return &TLAccountGetWallPaper{}
}

func NewTLAccountGetWallPapers() *TLAccountGetWallPapers {
	return &TLAccountGetWallPapers{
		Cmd:       Cmd_AccountGetWallPapers,
		ClassName: ClassAccountGetWallPapers,
	}
	//return &TLAccountGetWallPapers{}
}

func NewTLAccountGetWebAuthorizations() *TLAccountGetWebAuthorizations {
	return &TLAccountGetWebAuthorizations{
		Cmd:       Cmd_AccountGetWebAuthorizations,
		ClassName: ClassAccountGetWebAuthorizations,
	}
	//return &TLAccountGetWebAuthorizations{}
}

func NewTLAccountInitTakeoutSession() *TLAccountInitTakeoutSession {
	return &TLAccountInitTakeoutSession{
		Cmd:       Cmd_AccountInitTakeoutSession,
		ClassName: ClassAccountInitTakeoutSession,
	}
	//return &TLAccountInitTakeoutSession{}
}

func NewTLAccountInstallTheme() *TLAccountInstallTheme {
	return &TLAccountInstallTheme{
		Cmd:       Cmd_AccountInstallTheme,
		ClassName: ClassAccountInstallTheme,
	}
	//return &TLAccountInstallTheme{}
}

func NewTLAccountInstallWallPaper() *TLAccountInstallWallPaper {
	return &TLAccountInstallWallPaper{
		Cmd:       Cmd_AccountInstallWallPaper,
		ClassName: ClassAccountInstallWallPaper,
	}
	//return &TLAccountInstallWallPaper{}
}

func NewTLAccountRegisterDevice() *TLAccountRegisterDevice {
	return &TLAccountRegisterDevice{
		Cmd:       Cmd_AccountRegisterDevice,
		ClassName: ClassAccountRegisterDevice,
	}
	//return &TLAccountRegisterDevice{}
}

func NewTLAccountReportPeer() *TLAccountReportPeer {
	return &TLAccountReportPeer{
		Cmd:       Cmd_AccountReportPeer,
		ClassName: ClassAccountReportPeer,
	}
	//return &TLAccountReportPeer{}
}

func NewTLAccountResendPasswordEmail() *TLAccountResendPasswordEmail {
	return &TLAccountResendPasswordEmail{
		Cmd:       Cmd_AccountResendPasswordEmail,
		ClassName: ClassAccountResendPasswordEmail,
	}
	//return &TLAccountResendPasswordEmail{}
}

func NewTLAccountResetAuthorization() *TLAccountResetAuthorization {
	return &TLAccountResetAuthorization{
		Cmd:       Cmd_AccountResetAuthorization,
		ClassName: ClassAccountResetAuthorization,
	}
	//return &TLAccountResetAuthorization{}
}

func NewTLAccountResetNotifySettings() *TLAccountResetNotifySettings {
	return &TLAccountResetNotifySettings{
		Cmd:       Cmd_AccountResetNotifySettings,
		ClassName: ClassAccountResetNotifySettings,
	}
	//return &TLAccountResetNotifySettings{}
}

func NewTLAccountResetWallPapers() *TLAccountResetWallPapers {
	return &TLAccountResetWallPapers{
		Cmd:       Cmd_AccountResetWallPapers,
		ClassName: ClassAccountResetWallPapers,
	}
	//return &TLAccountResetWallPapers{}
}

func NewTLAccountResetWebAuthorization() *TLAccountResetWebAuthorization {
	return &TLAccountResetWebAuthorization{
		Cmd:       Cmd_AccountResetWebAuthorization,
		ClassName: ClassAccountResetWebAuthorization,
	}
	//return &TLAccountResetWebAuthorization{}
}

func NewTLAccountResetWebAuthorizations() *TLAccountResetWebAuthorizations {
	return &TLAccountResetWebAuthorizations{
		Cmd:       Cmd_AccountResetWebAuthorizations,
		ClassName: ClassAccountResetWebAuthorizations,
	}
	//return &TLAccountResetWebAuthorizations{}
}

func NewTLAccountSaveAutoDownloadSettings() *TLAccountSaveAutoDownloadSettings {
	return &TLAccountSaveAutoDownloadSettings{
		Cmd:       Cmd_AccountSaveAutoDownloadSettings,
		ClassName: ClassAccountSaveAutoDownloadSettings,
	}
	//return &TLAccountSaveAutoDownloadSettings{}
}

func NewTLAccountSaveSecureValue() *TLAccountSaveSecureValue {
	return &TLAccountSaveSecureValue{
		Cmd:       Cmd_AccountSaveSecureValue,
		ClassName: ClassAccountSaveSecureValue,
	}
	//return &TLAccountSaveSecureValue{}
}

func NewTLAccountSaveTheme() *TLAccountSaveTheme {
	return &TLAccountSaveTheme{
		Cmd:       Cmd_AccountSaveTheme,
		ClassName: ClassAccountSaveTheme,
	}
	//return &TLAccountSaveTheme{}
}

func NewTLAccountSaveWallPaper() *TLAccountSaveWallPaper {
	return &TLAccountSaveWallPaper{
		Cmd:       Cmd_AccountSaveWallPaper,
		ClassName: ClassAccountSaveWallPaper,
	}
	//return &TLAccountSaveWallPaper{}
}

func NewTLAccountSendChangePhoneCode() *TLAccountSendChangePhoneCode {
	return &TLAccountSendChangePhoneCode{
		Cmd:       Cmd_AccountSendChangePhoneCode,
		ClassName: ClassAccountSendChangePhoneCode,
	}
	//return &TLAccountSendChangePhoneCode{}
}

func NewTLAccountSendConfirmPhoneCode() *TLAccountSendConfirmPhoneCode {
	return &TLAccountSendConfirmPhoneCode{
		Cmd:       Cmd_AccountSendConfirmPhoneCode,
		ClassName: ClassAccountSendConfirmPhoneCode,
	}
	//return &TLAccountSendConfirmPhoneCode{}
}

func NewTLAccountSendVerifyEmailCode() *TLAccountSendVerifyEmailCode {
	return &TLAccountSendVerifyEmailCode{
		Cmd:       Cmd_AccountSendVerifyEmailCode,
		ClassName: ClassAccountSendVerifyEmailCode,
	}
	//return &TLAccountSendVerifyEmailCode{}
}

func NewTLAccountSendVerifyPhoneCode() *TLAccountSendVerifyPhoneCode {
	return &TLAccountSendVerifyPhoneCode{
		Cmd:       Cmd_AccountSendVerifyPhoneCode,
		ClassName: ClassAccountSendVerifyPhoneCode,
	}
	//return &TLAccountSendVerifyPhoneCode{}
}

func NewTLAccountSetAccountTTL() *TLAccountSetAccountTTL {
	return &TLAccountSetAccountTTL{
		Cmd:       Cmd_AccountSetAccountTTL,
		ClassName: ClassAccountSetAccountTTL,
	}
	//return &TLAccountSetAccountTTL{}
}

func NewTLAccountSetContactSignUpNotification() *TLAccountSetContactSignUpNotification {
	return &TLAccountSetContactSignUpNotification{
		Cmd:       Cmd_AccountSetContactSignUpNotification,
		ClassName: ClassAccountSetContactSignUpNotification,
	}
	//return &TLAccountSetContactSignUpNotification{}
}

func NewTLAccountSetContentSettings() *TLAccountSetContentSettings {
	return &TLAccountSetContentSettings{
		Cmd:       Cmd_AccountSetContentSettings,
		ClassName: ClassAccountSetContentSettings,
	}
	//return &TLAccountSetContentSettings{}
}

func NewTLAccountSetGlobalPrivacySettings() *TLAccountSetGlobalPrivacySettings {
	return &TLAccountSetGlobalPrivacySettings{
		Cmd:       Cmd_AccountSetGlobalPrivacySettings,
		ClassName: ClassAccountSetGlobalPrivacySettings,
	}
	//return &TLAccountSetGlobalPrivacySettings{}
}

func NewTLAccountSetPrivacy() *TLAccountSetPrivacy {
	return &TLAccountSetPrivacy{
		Cmd:       Cmd_AccountSetPrivacy,
		ClassName: ClassAccountSetPrivacy,
	}
	//return &TLAccountSetPrivacy{}
}

func NewTLAccountUnregisterDevice() *TLAccountUnregisterDevice {
	return &TLAccountUnregisterDevice{
		Cmd:       Cmd_AccountUnregisterDevice,
		ClassName: ClassAccountUnregisterDevice,
	}
	//return &TLAccountUnregisterDevice{}
}

func NewTLAccountUpdateDeviceLocked() *TLAccountUpdateDeviceLocked {
	return &TLAccountUpdateDeviceLocked{
		Cmd:       Cmd_AccountUpdateDeviceLocked,
		ClassName: ClassAccountUpdateDeviceLocked,
	}
	//return &TLAccountUpdateDeviceLocked{}
}

func NewTLAccountUpdateNotifySettings() *TLAccountUpdateNotifySettings {
	return &TLAccountUpdateNotifySettings{
		Cmd:       Cmd_AccountUpdateNotifySettings,
		ClassName: ClassAccountUpdateNotifySettings,
	}
	//return &TLAccountUpdateNotifySettings{}
}

func NewTLAccountUpdatePasswordSettings() *TLAccountUpdatePasswordSettings {
	return &TLAccountUpdatePasswordSettings{
		Cmd:       Cmd_AccountUpdatePasswordSettings,
		ClassName: ClassAccountUpdatePasswordSettings,
	}
	//return &TLAccountUpdatePasswordSettings{}
}

func NewTLAccountUpdateProfile() *TLAccountUpdateProfile {
	return &TLAccountUpdateProfile{
		Cmd:       Cmd_AccountUpdateProfile,
		ClassName: ClassAccountUpdateProfile,
	}
	//return &TLAccountUpdateProfile{}
}

func NewTLAccountUpdateStatus() *TLAccountUpdateStatus {
	return &TLAccountUpdateStatus{
		Cmd:       Cmd_AccountUpdateStatus,
		ClassName: ClassAccountUpdateStatus,
	}
	//return &TLAccountUpdateStatus{}
}

func NewTLAccountUpdateTheme() *TLAccountUpdateTheme {
	return &TLAccountUpdateTheme{
		Cmd:       Cmd_AccountUpdateTheme,
		ClassName: ClassAccountUpdateTheme,
	}
	//return &TLAccountUpdateTheme{}
}

func NewTLAccountUpdateUsername() *TLAccountUpdateUsername {
	return &TLAccountUpdateUsername{
		Cmd:       Cmd_AccountUpdateUsername,
		ClassName: ClassAccountUpdateUsername,
	}
	//return &TLAccountUpdateUsername{}
}

func NewTLAccountUploadTheme() *TLAccountUploadTheme {
	return &TLAccountUploadTheme{
		Cmd:       Cmd_AccountUploadTheme,
		ClassName: ClassAccountUploadTheme,
	}
	//return &TLAccountUploadTheme{}
}

func NewTLAccountUploadWallPaper() *TLAccountUploadWallPaper {
	return &TLAccountUploadWallPaper{
		Cmd:       Cmd_AccountUploadWallPaper,
		ClassName: ClassAccountUploadWallPaper,
	}
	//return &TLAccountUploadWallPaper{}
}

func NewTLAccountVerifyEmail() *TLAccountVerifyEmail {
	return &TLAccountVerifyEmail{
		Cmd:       Cmd_AccountVerifyEmail,
		ClassName: ClassAccountVerifyEmail,
	}
	//return &TLAccountVerifyEmail{}
}

func NewTLAccountVerifyPhone() *TLAccountVerifyPhone {
	return &TLAccountVerifyPhone{
		Cmd:       Cmd_AccountVerifyPhone,
		ClassName: ClassAccountVerifyPhone,
	}
	//return &TLAccountVerifyPhone{}
}

func NewTLAuthAcceptLoginToken() *TLAuthAcceptLoginToken {
	return &TLAuthAcceptLoginToken{
		Cmd:       Cmd_AuthAcceptLoginToken,
		ClassName: ClassAuthAcceptLoginToken,
	}
	//return &TLAuthAcceptLoginToken{}
}

func NewTLAuthBindTempAuthKey() *TLAuthBindTempAuthKey {
	return &TLAuthBindTempAuthKey{
		Cmd:       Cmd_AuthBindTempAuthKey,
		ClassName: ClassAuthBindTempAuthKey,
	}
	//return &TLAuthBindTempAuthKey{}
}

func NewTLAuthCancelCode() *TLAuthCancelCode {
	return &TLAuthCancelCode{
		Cmd:       Cmd_AuthCancelCode,
		ClassName: ClassAuthCancelCode,
	}
	//return &TLAuthCancelCode{}
}

func NewTLAuthCheckPassword() *TLAuthCheckPassword {
	return &TLAuthCheckPassword{
		Cmd:       Cmd_AuthCheckPassword,
		ClassName: ClassAuthCheckPassword,
	}
	//return &TLAuthCheckPassword{}
}

func NewTLAuthCheckPhone() *TLAuthCheckPhone {
	return &TLAuthCheckPhone{
		Cmd:       Cmd_AuthCheckPhone,
		ClassName: ClassAuthCheckPhone,
	}
	//return &TLAuthCheckPhone{}
}

func NewTLAuthDropTempAuthKeys() *TLAuthDropTempAuthKeys {
	return &TLAuthDropTempAuthKeys{
		Cmd:       Cmd_AuthDropTempAuthKeys,
		ClassName: ClassAuthDropTempAuthKeys,
	}
	//return &TLAuthDropTempAuthKeys{}
}

func NewTLAuthExportAuthorization() *TLAuthExportAuthorization {
	return &TLAuthExportAuthorization{
		Cmd:       Cmd_AuthExportAuthorization,
		ClassName: ClassAuthExportAuthorization,
	}
	//return &TLAuthExportAuthorization{}
}

func NewTLAuthExportLoginToken() *TLAuthExportLoginToken {
	return &TLAuthExportLoginToken{
		Cmd:       Cmd_AuthExportLoginToken,
		ClassName: ClassAuthExportLoginToken,
	}
	//return &TLAuthExportLoginToken{}
}

func NewTLAuthImportAuthorization() *TLAuthImportAuthorization {
	return &TLAuthImportAuthorization{
		Cmd:       Cmd_AuthImportAuthorization,
		ClassName: ClassAuthImportAuthorization,
	}
	//return &TLAuthImportAuthorization{}
}

func NewTLAuthImportBotAuthorization() *TLAuthImportBotAuthorization {
	return &TLAuthImportBotAuthorization{
		Cmd:       Cmd_AuthImportBotAuthorization,
		ClassName: ClassAuthImportBotAuthorization,
	}
	//return &TLAuthImportBotAuthorization{}
}

func NewTLAuthImportLoginToken() *TLAuthImportLoginToken {
	return &TLAuthImportLoginToken{
		Cmd:       Cmd_AuthImportLoginToken,
		ClassName: ClassAuthImportLoginToken,
	}
	//return &TLAuthImportLoginToken{}
}

func NewTLAuthLogOut() *TLAuthLogOut {
	return &TLAuthLogOut{
		Cmd:       Cmd_AuthLogOut,
		ClassName: ClassAuthLogOut,
	}
	//return &TLAuthLogOut{}
}

func NewTLAuthRecoverPassword() *TLAuthRecoverPassword {
	return &TLAuthRecoverPassword{
		Cmd:       Cmd_AuthRecoverPassword,
		ClassName: ClassAuthRecoverPassword,
	}
	//return &TLAuthRecoverPassword{}
}

func NewTLAuthRequestPasswordRecovery() *TLAuthRequestPasswordRecovery {
	return &TLAuthRequestPasswordRecovery{
		Cmd:       Cmd_AuthRequestPasswordRecovery,
		ClassName: ClassAuthRequestPasswordRecovery,
	}
	//return &TLAuthRequestPasswordRecovery{}
}

func NewTLAuthResendCode() *TLAuthResendCode {
	return &TLAuthResendCode{
		Cmd:       Cmd_AuthResendCode,
		ClassName: ClassAuthResendCode,
	}
	//return &TLAuthResendCode{}
}

func NewTLAuthResetAuthorizations() *TLAuthResetAuthorizations {
	return &TLAuthResetAuthorizations{
		Cmd:       Cmd_AuthResetAuthorizations,
		ClassName: ClassAuthResetAuthorizations,
	}
	//return &TLAuthResetAuthorizations{}
}

func NewTLAuthSendCode() *TLAuthSendCode {
	return &TLAuthSendCode{
		Cmd:       Cmd_AuthSendCode,
		ClassName: ClassAuthSendCode,
	}
	//return &TLAuthSendCode{}
}

func NewTLAuthSendInvites() *TLAuthSendInvites {
	return &TLAuthSendInvites{
		Cmd:       Cmd_AuthSendInvites,
		ClassName: ClassAuthSendInvites,
	}
	//return &TLAuthSendInvites{}
}

func NewTLAuthSignIn() *TLAuthSignIn {
	return &TLAuthSignIn{
		Cmd:       Cmd_AuthSignIn,
		ClassName: ClassAuthSignIn,
	}
	//return &TLAuthSignIn{}
}

func NewTLAuthSignUp() *TLAuthSignUp {
	return &TLAuthSignUp{
		Cmd:       Cmd_AuthSignUp,
		ClassName: ClassAuthSignUp,
	}
	//return &TLAuthSignUp{}
}

func NewTLBotsAnswerWebhookJSONQuery() *TLBotsAnswerWebhookJSONQuery {
	return &TLBotsAnswerWebhookJSONQuery{
		Cmd:       Cmd_BotsAnswerWebhookJSONQuery,
		ClassName: ClassBotsAnswerWebhookJSONQuery,
	}
	//return &TLBotsAnswerWebhookJSONQuery{}
}

func NewTLBotsSendCustomRequest() *TLBotsSendCustomRequest {
	return &TLBotsSendCustomRequest{
		Cmd:       Cmd_BotsSendCustomRequest,
		ClassName: ClassBotsSendCustomRequest,
	}
	//return &TLBotsSendCustomRequest{}
}

func NewTLBotsSetBotCommands() *TLBotsSetBotCommands {
	return &TLBotsSetBotCommands{
		Cmd:       Cmd_BotsSetBotCommands,
		ClassName: ClassBotsSetBotCommands,
	}
	//return &TLBotsSetBotCommands{}
}

func NewTLChannelsCheckUsername() *TLChannelsCheckUsername {
	return &TLChannelsCheckUsername{
		Cmd:       Cmd_ChannelsCheckUsername,
		ClassName: ClassChannelsCheckUsername,
	}
	//return &TLChannelsCheckUsername{}
}

func NewTLChannelsCreateChannel() *TLChannelsCreateChannel {
	return &TLChannelsCreateChannel{
		Cmd:       Cmd_ChannelsCreateChannel,
		ClassName: ClassChannelsCreateChannel,
	}
	//return &TLChannelsCreateChannel{}
}

func NewTLChannelsDeleteChannel() *TLChannelsDeleteChannel {
	return &TLChannelsDeleteChannel{
		Cmd:       Cmd_ChannelsDeleteChannel,
		ClassName: ClassChannelsDeleteChannel,
	}
	//return &TLChannelsDeleteChannel{}
}

func NewTLChannelsDeleteHistory() *TLChannelsDeleteHistory {
	return &TLChannelsDeleteHistory{
		Cmd:       Cmd_ChannelsDeleteHistory,
		ClassName: ClassChannelsDeleteHistory,
	}
	//return &TLChannelsDeleteHistory{}
}

func NewTLChannelsDeleteMessages() *TLChannelsDeleteMessages {
	return &TLChannelsDeleteMessages{
		Cmd:       Cmd_ChannelsDeleteMessages,
		ClassName: ClassChannelsDeleteMessages,
	}
	//return &TLChannelsDeleteMessages{}
}

func NewTLChannelsDeleteUserHistory() *TLChannelsDeleteUserHistory {
	return &TLChannelsDeleteUserHistory{
		Cmd:       Cmd_ChannelsDeleteUserHistory,
		ClassName: ClassChannelsDeleteUserHistory,
	}
	//return &TLChannelsDeleteUserHistory{}
}

func NewTLChannelsEditAbout() *TLChannelsEditAbout {
	return &TLChannelsEditAbout{
		Cmd:       Cmd_ChannelsEditAbout,
		ClassName: ClassChannelsEditAbout,
	}
	//return &TLChannelsEditAbout{}
}

func NewTLChannelsEditAdmin() *TLChannelsEditAdmin {
	return &TLChannelsEditAdmin{
		Cmd:       Cmd_ChannelsEditAdmin,
		ClassName: ClassChannelsEditAdmin,
	}
	//return &TLChannelsEditAdmin{}
}

func NewTLChannelsEditBanned() *TLChannelsEditBanned {
	return &TLChannelsEditBanned{
		Cmd:       Cmd_ChannelsEditBanned,
		ClassName: ClassChannelsEditBanned,
	}
	//return &TLChannelsEditBanned{}
}

func NewTLChannelsEditCreator() *TLChannelsEditCreator {
	return &TLChannelsEditCreator{
		Cmd:       Cmd_ChannelsEditCreator,
		ClassName: ClassChannelsEditCreator,
	}
	//return &TLChannelsEditCreator{}
}

func NewTLChannelsEditLocation() *TLChannelsEditLocation {
	return &TLChannelsEditLocation{
		Cmd:       Cmd_ChannelsEditLocation,
		ClassName: ClassChannelsEditLocation,
	}
	//return &TLChannelsEditLocation{}
}

func NewTLChannelsEditPhoto() *TLChannelsEditPhoto {
	return &TLChannelsEditPhoto{
		Cmd:       Cmd_ChannelsEditPhoto,
		ClassName: ClassChannelsEditPhoto,
	}
	//return &TLChannelsEditPhoto{}
}

func NewTLChannelsEditTitle() *TLChannelsEditTitle {
	return &TLChannelsEditTitle{
		Cmd:       Cmd_ChannelsEditTitle,
		ClassName: ClassChannelsEditTitle,
	}
	//return &TLChannelsEditTitle{}
}

func NewTLChannelsExportInvite() *TLChannelsExportInvite {
	return &TLChannelsExportInvite{
		Cmd:       Cmd_ChannelsExportInvite,
		ClassName: ClassChannelsExportInvite,
	}
	//return &TLChannelsExportInvite{}
}

func NewTLChannelsExportMessageLink() *TLChannelsExportMessageLink {
	return &TLChannelsExportMessageLink{
		Cmd:       Cmd_ChannelsExportMessageLink,
		ClassName: ClassChannelsExportMessageLink,
	}
	//return &TLChannelsExportMessageLink{}
}

func NewTLChannelsGetAdminLog() *TLChannelsGetAdminLog {
	return &TLChannelsGetAdminLog{
		Cmd:       Cmd_ChannelsGetAdminLog,
		ClassName: ClassChannelsGetAdminLog,
	}
	//return &TLChannelsGetAdminLog{}
}

func NewTLChannelsGetAdminedPublicChannels() *TLChannelsGetAdminedPublicChannels {
	return &TLChannelsGetAdminedPublicChannels{
		Cmd:       Cmd_ChannelsGetAdminedPublicChannels,
		ClassName: ClassChannelsGetAdminedPublicChannels,
	}
	//return &TLChannelsGetAdminedPublicChannels{}
}

func NewTLChannelsGetBroadcastsForDiscussion() *TLChannelsGetBroadcastsForDiscussion {
	return &TLChannelsGetBroadcastsForDiscussion{
		Cmd:       Cmd_ChannelsGetBroadcastsForDiscussion,
		ClassName: ClassChannelsGetBroadcastsForDiscussion,
	}
	//return &TLChannelsGetBroadcastsForDiscussion{}
}

func NewTLChannelsGetChannels() *TLChannelsGetChannels {
	return &TLChannelsGetChannels{
		Cmd:       Cmd_ChannelsGetChannels,
		ClassName: ClassChannelsGetChannels,
	}
	//return &TLChannelsGetChannels{}
}

func NewTLChannelsGetDialogs() *TLChannelsGetDialogs {
	return &TLChannelsGetDialogs{
		Cmd:       Cmd_ChannelsGetDialogs,
		ClassName: ClassChannelsGetDialogs,
	}
	//return &TLChannelsGetDialogs{}
}

func NewTLChannelsGetFullChannel() *TLChannelsGetFullChannel {
	return &TLChannelsGetFullChannel{
		Cmd:       Cmd_ChannelsGetFullChannel,
		ClassName: ClassChannelsGetFullChannel,
	}
	//return &TLChannelsGetFullChannel{}
}

func NewTLChannelsGetGroupsForDiscussion() *TLChannelsGetGroupsForDiscussion {
	return &TLChannelsGetGroupsForDiscussion{
		Cmd:       Cmd_ChannelsGetGroupsForDiscussion,
		ClassName: ClassChannelsGetGroupsForDiscussion,
	}
	//return &TLChannelsGetGroupsForDiscussion{}
}

func NewTLChannelsGetImportantHistory() *TLChannelsGetImportantHistory {
	return &TLChannelsGetImportantHistory{
		Cmd:       Cmd_ChannelsGetImportantHistory,
		ClassName: ClassChannelsGetImportantHistory,
	}
	//return &TLChannelsGetImportantHistory{}
}

func NewTLChannelsGetInactiveChannels() *TLChannelsGetInactiveChannels {
	return &TLChannelsGetInactiveChannels{
		Cmd:       Cmd_ChannelsGetInactiveChannels,
		ClassName: ClassChannelsGetInactiveChannels,
	}
	//return &TLChannelsGetInactiveChannels{}
}

func NewTLChannelsGetLeftChannels() *TLChannelsGetLeftChannels {
	return &TLChannelsGetLeftChannels{
		Cmd:       Cmd_ChannelsGetLeftChannels,
		ClassName: ClassChannelsGetLeftChannels,
	}
	//return &TLChannelsGetLeftChannels{}
}

func NewTLChannelsGetMessages() *TLChannelsGetMessages {
	return &TLChannelsGetMessages{
		Cmd:       Cmd_ChannelsGetMessages,
		ClassName: ClassChannelsGetMessages,
	}
	//return &TLChannelsGetMessages{}
}

func NewTLChannelsGetParticipant() *TLChannelsGetParticipant {
	return &TLChannelsGetParticipant{
		Cmd:       Cmd_ChannelsGetParticipant,
		ClassName: ClassChannelsGetParticipant,
	}
	//return &TLChannelsGetParticipant{}
}

func NewTLChannelsGetParticipants() *TLChannelsGetParticipants {
	return &TLChannelsGetParticipants{
		Cmd:       Cmd_ChannelsGetParticipants,
		ClassName: ClassChannelsGetParticipants,
	}
	//return &TLChannelsGetParticipants{}
}

func NewTLChannelsInviteToChannel() *TLChannelsInviteToChannel {
	return &TLChannelsInviteToChannel{
		Cmd:       Cmd_ChannelsInviteToChannel,
		ClassName: ClassChannelsInviteToChannel,
	}
	//return &TLChannelsInviteToChannel{}
}

func NewTLChannelsJoinChannel() *TLChannelsJoinChannel {
	return &TLChannelsJoinChannel{
		Cmd:       Cmd_ChannelsJoinChannel,
		ClassName: ClassChannelsJoinChannel,
	}
	//return &TLChannelsJoinChannel{}
}

func NewTLChannelsKickFromChannel() *TLChannelsKickFromChannel {
	return &TLChannelsKickFromChannel{
		Cmd:       Cmd_ChannelsKickFromChannel,
		ClassName: ClassChannelsKickFromChannel,
	}
	//return &TLChannelsKickFromChannel{}
}

func NewTLChannelsLeaveChannel() *TLChannelsLeaveChannel {
	return &TLChannelsLeaveChannel{
		Cmd:       Cmd_ChannelsLeaveChannel,
		ClassName: ClassChannelsLeaveChannel,
	}
	//return &TLChannelsLeaveChannel{}
}

func NewTLChannelsReadHistory() *TLChannelsReadHistory {
	return &TLChannelsReadHistory{
		Cmd:       Cmd_ChannelsReadHistory,
		ClassName: ClassChannelsReadHistory,
	}
	//return &TLChannelsReadHistory{}
}

func NewTLChannelsReadMessageContents() *TLChannelsReadMessageContents {
	return &TLChannelsReadMessageContents{
		Cmd:       Cmd_ChannelsReadMessageContents,
		ClassName: ClassChannelsReadMessageContents,
	}
	//return &TLChannelsReadMessageContents{}
}

func NewTLChannelsReportSpam() *TLChannelsReportSpam {
	return &TLChannelsReportSpam{
		Cmd:       Cmd_ChannelsReportSpam,
		ClassName: ClassChannelsReportSpam,
	}
	//return &TLChannelsReportSpam{}
}

func NewTLChannelsSetDiscussionGroup() *TLChannelsSetDiscussionGroup {
	return &TLChannelsSetDiscussionGroup{
		Cmd:       Cmd_ChannelsSetDiscussionGroup,
		ClassName: ClassChannelsSetDiscussionGroup,
	}
	//return &TLChannelsSetDiscussionGroup{}
}

func NewTLChannelsSetStickers() *TLChannelsSetStickers {
	return &TLChannelsSetStickers{
		Cmd:       Cmd_ChannelsSetStickers,
		ClassName: ClassChannelsSetStickers,
	}
	//return &TLChannelsSetStickers{}
}

func NewTLChannelsToggleComments() *TLChannelsToggleComments {
	return &TLChannelsToggleComments{
		Cmd:       Cmd_ChannelsToggleComments,
		ClassName: ClassChannelsToggleComments,
	}
	//return &TLChannelsToggleComments{}
}

func NewTLChannelsToggleInvites() *TLChannelsToggleInvites {
	return &TLChannelsToggleInvites{
		Cmd:       Cmd_ChannelsToggleInvites,
		ClassName: ClassChannelsToggleInvites,
	}
	//return &TLChannelsToggleInvites{}
}

func NewTLChannelsTogglePreHistoryHidden() *TLChannelsTogglePreHistoryHidden {
	return &TLChannelsTogglePreHistoryHidden{
		Cmd:       Cmd_ChannelsTogglePreHistoryHidden,
		ClassName: ClassChannelsTogglePreHistoryHidden,
	}
	//return &TLChannelsTogglePreHistoryHidden{}
}

func NewTLChannelsToggleSignatures() *TLChannelsToggleSignatures {
	return &TLChannelsToggleSignatures{
		Cmd:       Cmd_ChannelsToggleSignatures,
		ClassName: ClassChannelsToggleSignatures,
	}
	//return &TLChannelsToggleSignatures{}
}

func NewTLChannelsToggleSlowMode() *TLChannelsToggleSlowMode {
	return &TLChannelsToggleSlowMode{
		Cmd:       Cmd_ChannelsToggleSlowMode,
		ClassName: ClassChannelsToggleSlowMode,
	}
	//return &TLChannelsToggleSlowMode{}
}

func NewTLChannelsUpdatePinnedMessage() *TLChannelsUpdatePinnedMessage {
	return &TLChannelsUpdatePinnedMessage{
		Cmd:       Cmd_ChannelsUpdatePinnedMessage,
		ClassName: ClassChannelsUpdatePinnedMessage,
	}
	//return &TLChannelsUpdatePinnedMessage{}
}

func NewTLChannelsUpdateUsername() *TLChannelsUpdateUsername {
	return &TLChannelsUpdateUsername{
		Cmd:       Cmd_ChannelsUpdateUsername,
		ClassName: ClassChannelsUpdateUsername,
	}
	//return &TLChannelsUpdateUsername{}
}

func NewTLContactsAcceptContact() *TLContactsAcceptContact {
	return &TLContactsAcceptContact{
		Cmd:       Cmd_ContactsAcceptContact,
		ClassName: ClassContactsAcceptContact,
	}
	//return &TLContactsAcceptContact{}
}

func NewTLContactsAddContact() *TLContactsAddContact {
	return &TLContactsAddContact{
		Cmd:       Cmd_ContactsAddContact,
		ClassName: ClassContactsAddContact,
	}
	//return &TLContactsAddContact{}
}

func NewTLContactsBlock() *TLContactsBlock {
	return &TLContactsBlock{
		Cmd:       Cmd_ContactsBlock,
		ClassName: ClassContactsBlock,
	}
	//return &TLContactsBlock{}
}

func NewTLContactsBlockFromReplies() *TLContactsBlockFromReplies {
	return &TLContactsBlockFromReplies{
		Cmd:       Cmd_ContactsBlockFromReplies,
		ClassName: ClassContactsBlockFromReplies,
	}
	//return &TLContactsBlockFromReplies{}
}

func NewTLContactsDeleteByPhones() *TLContactsDeleteByPhones {
	return &TLContactsDeleteByPhones{
		Cmd:       Cmd_ContactsDeleteByPhones,
		ClassName: ClassContactsDeleteByPhones,
	}
	//return &TLContactsDeleteByPhones{}
}

func NewTLContactsDeleteContact() *TLContactsDeleteContact {
	return &TLContactsDeleteContact{
		Cmd:       Cmd_ContactsDeleteContact,
		ClassName: ClassContactsDeleteContact,
	}
	//return &TLContactsDeleteContact{}
}

func NewTLContactsDeleteContacts() *TLContactsDeleteContacts {
	return &TLContactsDeleteContacts{
		Cmd:       Cmd_ContactsDeleteContacts,
		ClassName: ClassContactsDeleteContacts,
	}
	//return &TLContactsDeleteContacts{}
}

func NewTLContactsExportCard() *TLContactsExportCard {
	return &TLContactsExportCard{
		Cmd:       Cmd_ContactsExportCard,
		ClassName: ClassContactsExportCard,
	}
	//return &TLContactsExportCard{}
}

func NewTLContactsGetBlocked() *TLContactsGetBlocked {
	return &TLContactsGetBlocked{
		Cmd:       Cmd_ContactsGetBlocked,
		ClassName: ClassContactsGetBlocked,
	}
	//return &TLContactsGetBlocked{}
}

func NewTLContactsGetContactIDs() *TLContactsGetContactIDs {
	return &TLContactsGetContactIDs{
		Cmd:       Cmd_ContactsGetContactIDs,
		ClassName: ClassContactsGetContactIDs,
	}
	//return &TLContactsGetContactIDs{}
}

func NewTLContactsGetContacts() *TLContactsGetContacts {
	return &TLContactsGetContacts{
		Cmd:       Cmd_ContactsGetContacts,
		ClassName: ClassContactsGetContacts,
	}
	//return &TLContactsGetContacts{}
}

func NewTLContactsGetLocated() *TLContactsGetLocated {
	return &TLContactsGetLocated{
		Cmd:       Cmd_ContactsGetLocated,
		ClassName: ClassContactsGetLocated,
	}
	//return &TLContactsGetLocated{}
}

func NewTLContactsGetSaved() *TLContactsGetSaved {
	return &TLContactsGetSaved{
		Cmd:       Cmd_ContactsGetSaved,
		ClassName: ClassContactsGetSaved,
	}
	//return &TLContactsGetSaved{}
}

func NewTLContactsGetStatuses() *TLContactsGetStatuses {
	return &TLContactsGetStatuses{
		Cmd:       Cmd_ContactsGetStatuses,
		ClassName: ClassContactsGetStatuses,
	}
	//return &TLContactsGetStatuses{}
}

func NewTLContactsGetTopPeers() *TLContactsGetTopPeers {
	return &TLContactsGetTopPeers{
		Cmd:       Cmd_ContactsGetTopPeers,
		ClassName: ClassContactsGetTopPeers,
	}
	//return &TLContactsGetTopPeers{}
}

func NewTLContactsImportCard() *TLContactsImportCard {
	return &TLContactsImportCard{
		Cmd:       Cmd_ContactsImportCard,
		ClassName: ClassContactsImportCard,
	}
	//return &TLContactsImportCard{}
}

func NewTLContactsImportContacts() *TLContactsImportContacts {
	return &TLContactsImportContacts{
		Cmd:       Cmd_ContactsImportContacts,
		ClassName: ClassContactsImportContacts,
	}
	//return &TLContactsImportContacts{}
}

func NewTLContactsResetSaved() *TLContactsResetSaved {
	return &TLContactsResetSaved{
		Cmd:       Cmd_ContactsResetSaved,
		ClassName: ClassContactsResetSaved,
	}
	//return &TLContactsResetSaved{}
}

func NewTLContactsResetTopPeerRating() *TLContactsResetTopPeerRating {
	return &TLContactsResetTopPeerRating{
		Cmd:       Cmd_ContactsResetTopPeerRating,
		ClassName: ClassContactsResetTopPeerRating,
	}
	//return &TLContactsResetTopPeerRating{}
}

func NewTLContactsResolveUsername() *TLContactsResolveUsername {
	return &TLContactsResolveUsername{
		Cmd:       Cmd_ContactsResolveUsername,
		ClassName: ClassContactsResolveUsername,
	}
	//return &TLContactsResolveUsername{}
}

func NewTLContactsSearch() *TLContactsSearch {
	return &TLContactsSearch{
		Cmd:       Cmd_ContactsSearch,
		ClassName: ClassContactsSearch,
	}
	//return &TLContactsSearch{}
}

func NewTLContactsToggleTopPeers() *TLContactsToggleTopPeers {
	return &TLContactsToggleTopPeers{
		Cmd:       Cmd_ContactsToggleTopPeers,
		ClassName: ClassContactsToggleTopPeers,
	}
	//return &TLContactsToggleTopPeers{}
}

func NewTLContactsUnblock() *TLContactsUnblock {
	return &TLContactsUnblock{
		Cmd:       Cmd_ContactsUnblock,
		ClassName: ClassContactsUnblock,
	}
	//return &TLContactsUnblock{}
}

func NewTLContestSaveDeveloperInfo() *TLContestSaveDeveloperInfo {
	return &TLContestSaveDeveloperInfo{
		Cmd:       Cmd_ContestSaveDeveloperInfo,
		ClassName: ClassContestSaveDeveloperInfo,
	}
	//return &TLContestSaveDeveloperInfo{}
}

func NewTLDestroyAuthKey() *TLDestroyAuthKey {
	return &TLDestroyAuthKey{
		Cmd:       Cmd_DestroyAuthKey,
		ClassName: ClassDestroyAuthKey,
	}
	//return &TLDestroyAuthKey{}
}

func NewTLDestroySession() *TLDestroySession {
	return &TLDestroySession{
		Cmd:       Cmd_DestroySession,
		ClassName: ClassDestroySession,
	}
	//return &TLDestroySession{}
}

func NewTLFoldersDeleteFolder() *TLFoldersDeleteFolder {
	return &TLFoldersDeleteFolder{
		Cmd:       Cmd_FoldersDeleteFolder,
		ClassName: ClassFoldersDeleteFolder,
	}
	//return &TLFoldersDeleteFolder{}
}

func NewTLFoldersEditPeerFolders() *TLFoldersEditPeerFolders {
	return &TLFoldersEditPeerFolders{
		Cmd:       Cmd_FoldersEditPeerFolders,
		ClassName: ClassFoldersEditPeerFolders,
	}
	//return &TLFoldersEditPeerFolders{}
}

func NewTLGetFutureSalts() *TLGetFutureSalts {
	return &TLGetFutureSalts{
		Cmd:       Cmd_GetFutureSalts,
		ClassName: ClassGetFutureSalts,
	}
	//return &TLGetFutureSalts{}
}

func NewTLHelpAcceptTermsOfService() *TLHelpAcceptTermsOfService {
	return &TLHelpAcceptTermsOfService{
		Cmd:       Cmd_HelpAcceptTermsOfService,
		ClassName: ClassHelpAcceptTermsOfService,
	}
	//return &TLHelpAcceptTermsOfService{}
}

func NewTLHelpDismissSuggestion() *TLHelpDismissSuggestion {
	return &TLHelpDismissSuggestion{
		Cmd:       Cmd_HelpDismissSuggestion,
		ClassName: ClassHelpDismissSuggestion,
	}
	//return &TLHelpDismissSuggestion{}
}

func NewTLHelpEditUserInfo() *TLHelpEditUserInfo {
	return &TLHelpEditUserInfo{
		Cmd:       Cmd_HelpEditUserInfo,
		ClassName: ClassHelpEditUserInfo,
	}
	//return &TLHelpEditUserInfo{}
}

func NewTLHelpGetAppChangelog() *TLHelpGetAppChangelog {
	return &TLHelpGetAppChangelog{
		Cmd:       Cmd_HelpGetAppChangelog,
		ClassName: ClassHelpGetAppChangelog,
	}
	//return &TLHelpGetAppChangelog{}
}

func NewTLHelpGetAppConfig() *TLHelpGetAppConfig {
	return &TLHelpGetAppConfig{
		Cmd:       Cmd_HelpGetAppConfig,
		ClassName: ClassHelpGetAppConfig,
	}
	//return &TLHelpGetAppConfig{}
}

func NewTLHelpGetAppUpdate() *TLHelpGetAppUpdate {
	return &TLHelpGetAppUpdate{
		Cmd:       Cmd_HelpGetAppUpdate,
		ClassName: ClassHelpGetAppUpdate,
	}
	//return &TLHelpGetAppUpdate{}
}

func NewTLHelpGetCdnConfig() *TLHelpGetCdnConfig {
	return &TLHelpGetCdnConfig{
		Cmd:       Cmd_HelpGetCdnConfig,
		ClassName: ClassHelpGetCdnConfig,
	}
	//return &TLHelpGetCdnConfig{}
}

func NewTLHelpGetConfig() *TLHelpGetConfig {
	return &TLHelpGetConfig{
		Cmd:       Cmd_HelpGetConfig,
		ClassName: ClassHelpGetConfig,
	}
	//return &TLHelpGetConfig{}
}

func NewTLHelpGetCountriesList() *TLHelpGetCountriesList {
	return &TLHelpGetCountriesList{
		Cmd:       Cmd_HelpGetCountriesList,
		ClassName: ClassHelpGetCountriesList,
	}
	//return &TLHelpGetCountriesList{}
}

func NewTLHelpGetDeepLinkInfo() *TLHelpGetDeepLinkInfo {
	return &TLHelpGetDeepLinkInfo{
		Cmd:       Cmd_HelpGetDeepLinkInfo,
		ClassName: ClassHelpGetDeepLinkInfo,
	}
	//return &TLHelpGetDeepLinkInfo{}
}

func NewTLHelpGetInviteText() *TLHelpGetInviteText {
	return &TLHelpGetInviteText{
		Cmd:       Cmd_HelpGetInviteText,
		ClassName: ClassHelpGetInviteText,
	}
	//return &TLHelpGetInviteText{}
}

func NewTLHelpGetNearestDc() *TLHelpGetNearestDc {
	return &TLHelpGetNearestDc{
		Cmd:       Cmd_HelpGetNearestDc,
		ClassName: ClassHelpGetNearestDc,
	}
	//return &TLHelpGetNearestDc{}
}

func NewTLHelpGetPassportConfig() *TLHelpGetPassportConfig {
	return &TLHelpGetPassportConfig{
		Cmd:       Cmd_HelpGetPassportConfig,
		ClassName: ClassHelpGetPassportConfig,
	}
	//return &TLHelpGetPassportConfig{}
}

func NewTLHelpGetPromoData() *TLHelpGetPromoData {
	return &TLHelpGetPromoData{
		Cmd:       Cmd_HelpGetPromoData,
		ClassName: ClassHelpGetPromoData,
	}
	//return &TLHelpGetPromoData{}
}

func NewTLHelpGetProxyData() *TLHelpGetProxyData {
	return &TLHelpGetProxyData{
		Cmd:       Cmd_HelpGetProxyData,
		ClassName: ClassHelpGetProxyData,
	}
	//return &TLHelpGetProxyData{}
}

func NewTLHelpGetRecentMeUrls() *TLHelpGetRecentMeUrls {
	return &TLHelpGetRecentMeUrls{
		Cmd:       Cmd_HelpGetRecentMeUrls,
		ClassName: ClassHelpGetRecentMeUrls,
	}
	//return &TLHelpGetRecentMeUrls{}
}

func NewTLHelpGetSupport() *TLHelpGetSupport {
	return &TLHelpGetSupport{
		Cmd:       Cmd_HelpGetSupport,
		ClassName: ClassHelpGetSupport,
	}
	//return &TLHelpGetSupport{}
}

func NewTLHelpGetSupportName() *TLHelpGetSupportName {
	return &TLHelpGetSupportName{
		Cmd:       Cmd_HelpGetSupportName,
		ClassName: ClassHelpGetSupportName,
	}
	//return &TLHelpGetSupportName{}
}

func NewTLHelpGetTermsOfService() *TLHelpGetTermsOfService {
	return &TLHelpGetTermsOfService{
		Cmd:       Cmd_HelpGetTermsOfService,
		ClassName: ClassHelpGetTermsOfService,
	}
	//return &TLHelpGetTermsOfService{}
}

func NewTLHelpGetTermsOfServiceUpdate() *TLHelpGetTermsOfServiceUpdate {
	return &TLHelpGetTermsOfServiceUpdate{
		Cmd:       Cmd_HelpGetTermsOfServiceUpdate,
		ClassName: ClassHelpGetTermsOfServiceUpdate,
	}
	//return &TLHelpGetTermsOfServiceUpdate{}
}

func NewTLHelpGetUserInfo() *TLHelpGetUserInfo {
	return &TLHelpGetUserInfo{
		Cmd:       Cmd_HelpGetUserInfo,
		ClassName: ClassHelpGetUserInfo,
	}
	//return &TLHelpGetUserInfo{}
}

func NewTLHelpHidePromoData() *TLHelpHidePromoData {
	return &TLHelpHidePromoData{
		Cmd:       Cmd_HelpHidePromoData,
		ClassName: ClassHelpHidePromoData,
	}
	//return &TLHelpHidePromoData{}
}

func NewTLHelpSaveAppLog() *TLHelpSaveAppLog {
	return &TLHelpSaveAppLog{
		Cmd:       Cmd_HelpSaveAppLog,
		ClassName: ClassHelpSaveAppLog,
	}
	//return &TLHelpSaveAppLog{}
}

func NewTLHelpSetBotUpdatesStatus() *TLHelpSetBotUpdatesStatus {
	return &TLHelpSetBotUpdatesStatus{
		Cmd:       Cmd_HelpSetBotUpdatesStatus,
		ClassName: ClassHelpSetBotUpdatesStatus,
	}
	//return &TLHelpSetBotUpdatesStatus{}
}

func NewTLHelpTest() *TLHelpTest {
	return &TLHelpTest{
		Cmd:       Cmd_HelpTest,
		ClassName: ClassHelpTest,
	}
	//return &TLHelpTest{}
}

func NewTLLangpackGetDifference() *TLLangpackGetDifference {
	return &TLLangpackGetDifference{
		Cmd:       Cmd_LangpackGetDifference,
		ClassName: ClassLangpackGetDifference,
	}
	//return &TLLangpackGetDifference{}
}

func NewTLLangpackGetLangPack() *TLLangpackGetLangPack {
	return &TLLangpackGetLangPack{
		Cmd:       Cmd_LangpackGetLangPack,
		ClassName: ClassLangpackGetLangPack,
	}
	//return &TLLangpackGetLangPack{}
}

func NewTLLangpackGetLanguage() *TLLangpackGetLanguage {
	return &TLLangpackGetLanguage{
		Cmd:       Cmd_LangpackGetLanguage,
		ClassName: ClassLangpackGetLanguage,
	}
	//return &TLLangpackGetLanguage{}
}

func NewTLLangpackGetLanguages() *TLLangpackGetLanguages {
	return &TLLangpackGetLanguages{
		Cmd:       Cmd_LangpackGetLanguages,
		ClassName: ClassLangpackGetLanguages,
	}
	//return &TLLangpackGetLanguages{}
}

func NewTLLangpackGetStrings() *TLLangpackGetStrings {
	return &TLLangpackGetStrings{
		Cmd:       Cmd_LangpackGetStrings,
		ClassName: ClassLangpackGetStrings,
	}
	//return &TLLangpackGetStrings{}
}

func NewTLMessagesAcceptEncryption() *TLMessagesAcceptEncryption {
	return &TLMessagesAcceptEncryption{
		Cmd:       Cmd_MessagesAcceptEncryption,
		ClassName: ClassMessagesAcceptEncryption,
	}
	//return &TLMessagesAcceptEncryption{}
}

func NewTLMessagesAcceptUrlAuth() *TLMessagesAcceptUrlAuth {
	return &TLMessagesAcceptUrlAuth{
		Cmd:       Cmd_MessagesAcceptUrlAuth,
		ClassName: ClassMessagesAcceptUrlAuth,
	}
	//return &TLMessagesAcceptUrlAuth{}
}

func NewTLMessagesAddChatUser() *TLMessagesAddChatUser {
	return &TLMessagesAddChatUser{
		Cmd:       Cmd_MessagesAddChatUser,
		ClassName: ClassMessagesAddChatUser,
	}
	//return &TLMessagesAddChatUser{}
}

func NewTLMessagesCheckChatInvite() *TLMessagesCheckChatInvite {
	return &TLMessagesCheckChatInvite{
		Cmd:       Cmd_MessagesCheckChatInvite,
		ClassName: ClassMessagesCheckChatInvite,
	}
	//return &TLMessagesCheckChatInvite{}
}

func NewTLMessagesClearAllDrafts() *TLMessagesClearAllDrafts {
	return &TLMessagesClearAllDrafts{
		Cmd:       Cmd_MessagesClearAllDrafts,
		ClassName: ClassMessagesClearAllDrafts,
	}
	//return &TLMessagesClearAllDrafts{}
}

func NewTLMessagesClearRecentStickers() *TLMessagesClearRecentStickers {
	return &TLMessagesClearRecentStickers{
		Cmd:       Cmd_MessagesClearRecentStickers,
		ClassName: ClassMessagesClearRecentStickers,
	}
	//return &TLMessagesClearRecentStickers{}
}

func NewTLMessagesCreateChat() *TLMessagesCreateChat {
	return &TLMessagesCreateChat{
		Cmd:       Cmd_MessagesCreateChat,
		ClassName: ClassMessagesCreateChat,
	}
	//return &TLMessagesCreateChat{}
}

func NewTLMessagesDeleteChatUser() *TLMessagesDeleteChatUser {
	return &TLMessagesDeleteChatUser{
		Cmd:       Cmd_MessagesDeleteChatUser,
		ClassName: ClassMessagesDeleteChatUser,
	}
	//return &TLMessagesDeleteChatUser{}
}

func NewTLMessagesDeleteHistory() *TLMessagesDeleteHistory {
	return &TLMessagesDeleteHistory{
		Cmd:       Cmd_MessagesDeleteHistory,
		ClassName: ClassMessagesDeleteHistory,
	}
	//return &TLMessagesDeleteHistory{}
}

func NewTLMessagesDeleteMessages() *TLMessagesDeleteMessages {
	return &TLMessagesDeleteMessages{
		Cmd:       Cmd_MessagesDeleteMessages,
		ClassName: ClassMessagesDeleteMessages,
	}
	//return &TLMessagesDeleteMessages{}
}

func NewTLMessagesDeleteScheduledMessages() *TLMessagesDeleteScheduledMessages {
	return &TLMessagesDeleteScheduledMessages{
		Cmd:       Cmd_MessagesDeleteScheduledMessages,
		ClassName: ClassMessagesDeleteScheduledMessages,
	}
	//return &TLMessagesDeleteScheduledMessages{}
}

func NewTLMessagesDiscardEncryption() *TLMessagesDiscardEncryption {
	return &TLMessagesDiscardEncryption{
		Cmd:       Cmd_MessagesDiscardEncryption,
		ClassName: ClassMessagesDiscardEncryption,
	}
	//return &TLMessagesDiscardEncryption{}
}

func NewTLMessagesEditChatAbout() *TLMessagesEditChatAbout {
	return &TLMessagesEditChatAbout{
		Cmd:       Cmd_MessagesEditChatAbout,
		ClassName: ClassMessagesEditChatAbout,
	}
	//return &TLMessagesEditChatAbout{}
}

func NewTLMessagesEditChatAdmin() *TLMessagesEditChatAdmin {
	return &TLMessagesEditChatAdmin{
		Cmd:       Cmd_MessagesEditChatAdmin,
		ClassName: ClassMessagesEditChatAdmin,
	}
	//return &TLMessagesEditChatAdmin{}
}

func NewTLMessagesEditChatDefaultBannedRights() *TLMessagesEditChatDefaultBannedRights {
	return &TLMessagesEditChatDefaultBannedRights{
		Cmd:       Cmd_MessagesEditChatDefaultBannedRights,
		ClassName: ClassMessagesEditChatDefaultBannedRights,
	}
	//return &TLMessagesEditChatDefaultBannedRights{}
}

func NewTLMessagesEditChatPhoto() *TLMessagesEditChatPhoto {
	return &TLMessagesEditChatPhoto{
		Cmd:       Cmd_MessagesEditChatPhoto,
		ClassName: ClassMessagesEditChatPhoto,
	}
	//return &TLMessagesEditChatPhoto{}
}

func NewTLMessagesEditChatTitle() *TLMessagesEditChatTitle {
	return &TLMessagesEditChatTitle{
		Cmd:       Cmd_MessagesEditChatTitle,
		ClassName: ClassMessagesEditChatTitle,
	}
	//return &TLMessagesEditChatTitle{}
}

func NewTLMessagesEditInlineBotMessage() *TLMessagesEditInlineBotMessage {
	return &TLMessagesEditInlineBotMessage{
		Cmd:       Cmd_MessagesEditInlineBotMessage,
		ClassName: ClassMessagesEditInlineBotMessage,
	}
	//return &TLMessagesEditInlineBotMessage{}
}

func NewTLMessagesEditMessage() *TLMessagesEditMessage {
	return &TLMessagesEditMessage{
		Cmd:       Cmd_MessagesEditMessage,
		ClassName: ClassMessagesEditMessage,
	}
	//return &TLMessagesEditMessage{}
}

func NewTLMessagesExportChatInvite() *TLMessagesExportChatInvite {
	return &TLMessagesExportChatInvite{
		Cmd:       Cmd_MessagesExportChatInvite,
		ClassName: ClassMessagesExportChatInvite,
	}
	//return &TLMessagesExportChatInvite{}
}

func NewTLMessagesFaveSticker() *TLMessagesFaveSticker {
	return &TLMessagesFaveSticker{
		Cmd:       Cmd_MessagesFaveSticker,
		ClassName: ClassMessagesFaveSticker,
	}
	//return &TLMessagesFaveSticker{}
}

func NewTLMessagesForwardMessage() *TLMessagesForwardMessage {
	return &TLMessagesForwardMessage{
		Cmd:       Cmd_MessagesForwardMessage,
		ClassName: ClassMessagesForwardMessage,
	}
	//return &TLMessagesForwardMessage{}
}

func NewTLMessagesForwardMessages() *TLMessagesForwardMessages {
	return &TLMessagesForwardMessages{
		Cmd:       Cmd_MessagesForwardMessages,
		ClassName: ClassMessagesForwardMessages,
	}
	//return &TLMessagesForwardMessages{}
}

func NewTLMessagesGetAllChats() *TLMessagesGetAllChats {
	return &TLMessagesGetAllChats{
		Cmd:       Cmd_MessagesGetAllChats,
		ClassName: ClassMessagesGetAllChats,
	}
	//return &TLMessagesGetAllChats{}
}

func NewTLMessagesGetAllDrafts() *TLMessagesGetAllDrafts {
	return &TLMessagesGetAllDrafts{
		Cmd:       Cmd_MessagesGetAllDrafts,
		ClassName: ClassMessagesGetAllDrafts,
	}
	//return &TLMessagesGetAllDrafts{}
}

func NewTLMessagesGetAllStickers() *TLMessagesGetAllStickers {
	return &TLMessagesGetAllStickers{
		Cmd:       Cmd_MessagesGetAllStickers,
		ClassName: ClassMessagesGetAllStickers,
	}
	//return &TLMessagesGetAllStickers{}
}

func NewTLMessagesGetArchivedStickers() *TLMessagesGetArchivedStickers {
	return &TLMessagesGetArchivedStickers{
		Cmd:       Cmd_MessagesGetArchivedStickers,
		ClassName: ClassMessagesGetArchivedStickers,
	}
	//return &TLMessagesGetArchivedStickers{}
}

func NewTLMessagesGetAttachedStickers() *TLMessagesGetAttachedStickers {
	return &TLMessagesGetAttachedStickers{
		Cmd:       Cmd_MessagesGetAttachedStickers,
		ClassName: ClassMessagesGetAttachedStickers,
	}
	//return &TLMessagesGetAttachedStickers{}
}

func NewTLMessagesGetBotCallbackAnswer() *TLMessagesGetBotCallbackAnswer {
	return &TLMessagesGetBotCallbackAnswer{
		Cmd:       Cmd_MessagesGetBotCallbackAnswer,
		ClassName: ClassMessagesGetBotCallbackAnswer,
	}
	//return &TLMessagesGetBotCallbackAnswer{}
}

func NewTLMessagesGetChats() *TLMessagesGetChats {
	return &TLMessagesGetChats{
		Cmd:       Cmd_MessagesGetChats,
		ClassName: ClassMessagesGetChats,
	}
	//return &TLMessagesGetChats{}
}

func NewTLMessagesGetCommonChats() *TLMessagesGetCommonChats {
	return &TLMessagesGetCommonChats{
		Cmd:       Cmd_MessagesGetCommonChats,
		ClassName: ClassMessagesGetCommonChats,
	}
	//return &TLMessagesGetCommonChats{}
}

func NewTLMessagesGetDhConfig() *TLMessagesGetDhConfig {
	return &TLMessagesGetDhConfig{
		Cmd:       Cmd_MessagesGetDhConfig,
		ClassName: ClassMessagesGetDhConfig,
	}
	//return &TLMessagesGetDhConfig{}
}

func NewTLMessagesGetDialogFilters() *TLMessagesGetDialogFilters {
	return &TLMessagesGetDialogFilters{
		Cmd:       Cmd_MessagesGetDialogFilters,
		ClassName: ClassMessagesGetDialogFilters,
	}
	//return &TLMessagesGetDialogFilters{}
}

func NewTLMessagesGetDialogUnreadMarks() *TLMessagesGetDialogUnreadMarks {
	return &TLMessagesGetDialogUnreadMarks{
		Cmd:       Cmd_MessagesGetDialogUnreadMarks,
		ClassName: ClassMessagesGetDialogUnreadMarks,
	}
	//return &TLMessagesGetDialogUnreadMarks{}
}

func NewTLMessagesGetDialogs() *TLMessagesGetDialogs {
	return &TLMessagesGetDialogs{
		Cmd:       Cmd_MessagesGetDialogs,
		ClassName: ClassMessagesGetDialogs,
	}
	//return &TLMessagesGetDialogs{}
}

func NewTLMessagesGetDiscussionMessage() *TLMessagesGetDiscussionMessage {
	return &TLMessagesGetDiscussionMessage{
		Cmd:       Cmd_MessagesGetDiscussionMessage,
		ClassName: ClassMessagesGetDiscussionMessage,
	}
	//return &TLMessagesGetDiscussionMessage{}
}

func NewTLMessagesGetDocumentByHash() *TLMessagesGetDocumentByHash {
	return &TLMessagesGetDocumentByHash{
		Cmd:       Cmd_MessagesGetDocumentByHash,
		ClassName: ClassMessagesGetDocumentByHash,
	}
	//return &TLMessagesGetDocumentByHash{}
}

func NewTLMessagesGetEmojiKeywords() *TLMessagesGetEmojiKeywords {
	return &TLMessagesGetEmojiKeywords{
		Cmd:       Cmd_MessagesGetEmojiKeywords,
		ClassName: ClassMessagesGetEmojiKeywords,
	}
	//return &TLMessagesGetEmojiKeywords{}
}

func NewTLMessagesGetEmojiKeywordsDifference() *TLMessagesGetEmojiKeywordsDifference {
	return &TLMessagesGetEmojiKeywordsDifference{
		Cmd:       Cmd_MessagesGetEmojiKeywordsDifference,
		ClassName: ClassMessagesGetEmojiKeywordsDifference,
	}
	//return &TLMessagesGetEmojiKeywordsDifference{}
}

func NewTLMessagesGetEmojiKeywordsLanguages() *TLMessagesGetEmojiKeywordsLanguages {
	return &TLMessagesGetEmojiKeywordsLanguages{
		Cmd:       Cmd_MessagesGetEmojiKeywordsLanguages,
		ClassName: ClassMessagesGetEmojiKeywordsLanguages,
	}
	//return &TLMessagesGetEmojiKeywordsLanguages{}
}

func NewTLMessagesGetEmojiURL() *TLMessagesGetEmojiURL {
	return &TLMessagesGetEmojiURL{
		Cmd:       Cmd_MessagesGetEmojiURL,
		ClassName: ClassMessagesGetEmojiURL,
	}
	//return &TLMessagesGetEmojiURL{}
}

func NewTLMessagesGetFavedStickers() *TLMessagesGetFavedStickers {
	return &TLMessagesGetFavedStickers{
		Cmd:       Cmd_MessagesGetFavedStickers,
		ClassName: ClassMessagesGetFavedStickers,
	}
	//return &TLMessagesGetFavedStickers{}
}

func NewTLMessagesGetFeaturedStickers() *TLMessagesGetFeaturedStickers {
	return &TLMessagesGetFeaturedStickers{
		Cmd:       Cmd_MessagesGetFeaturedStickers,
		ClassName: ClassMessagesGetFeaturedStickers,
	}
	//return &TLMessagesGetFeaturedStickers{}
}

func NewTLMessagesGetFullChat() *TLMessagesGetFullChat {
	return &TLMessagesGetFullChat{
		Cmd:       Cmd_MessagesGetFullChat,
		ClassName: ClassMessagesGetFullChat,
	}
	//return &TLMessagesGetFullChat{}
}

func NewTLMessagesGetGameHighScores() *TLMessagesGetGameHighScores {
	return &TLMessagesGetGameHighScores{
		Cmd:       Cmd_MessagesGetGameHighScores,
		ClassName: ClassMessagesGetGameHighScores,
	}
	//return &TLMessagesGetGameHighScores{}
}

func NewTLMessagesGetHistory() *TLMessagesGetHistory {
	return &TLMessagesGetHistory{
		Cmd:       Cmd_MessagesGetHistory,
		ClassName: ClassMessagesGetHistory,
	}
	//return &TLMessagesGetHistory{}
}

func NewTLMessagesGetInlineBotResults() *TLMessagesGetInlineBotResults {
	return &TLMessagesGetInlineBotResults{
		Cmd:       Cmd_MessagesGetInlineBotResults,
		ClassName: ClassMessagesGetInlineBotResults,
	}
	//return &TLMessagesGetInlineBotResults{}
}

func NewTLMessagesGetInlineGameHighScores() *TLMessagesGetInlineGameHighScores {
	return &TLMessagesGetInlineGameHighScores{
		Cmd:       Cmd_MessagesGetInlineGameHighScores,
		ClassName: ClassMessagesGetInlineGameHighScores,
	}
	//return &TLMessagesGetInlineGameHighScores{}
}

func NewTLMessagesGetMaskStickers() *TLMessagesGetMaskStickers {
	return &TLMessagesGetMaskStickers{
		Cmd:       Cmd_MessagesGetMaskStickers,
		ClassName: ClassMessagesGetMaskStickers,
	}
	//return &TLMessagesGetMaskStickers{}
}

func NewTLMessagesGetMessageEditData() *TLMessagesGetMessageEditData {
	return &TLMessagesGetMessageEditData{
		Cmd:       Cmd_MessagesGetMessageEditData,
		ClassName: ClassMessagesGetMessageEditData,
	}
	//return &TLMessagesGetMessageEditData{}
}

func NewTLMessagesGetMessages() *TLMessagesGetMessages {
	return &TLMessagesGetMessages{
		Cmd:       Cmd_MessagesGetMessages,
		ClassName: ClassMessagesGetMessages,
	}
	//return &TLMessagesGetMessages{}
}

func NewTLMessagesGetMessagesViews() *TLMessagesGetMessagesViews {
	return &TLMessagesGetMessagesViews{
		Cmd:       Cmd_MessagesGetMessagesViews,
		ClassName: ClassMessagesGetMessagesViews,
	}
	//return &TLMessagesGetMessagesViews{}
}

func NewTLMessagesGetOldFeaturedStickers() *TLMessagesGetOldFeaturedStickers {
	return &TLMessagesGetOldFeaturedStickers{
		Cmd:       Cmd_MessagesGetOldFeaturedStickers,
		ClassName: ClassMessagesGetOldFeaturedStickers,
	}
	//return &TLMessagesGetOldFeaturedStickers{}
}

func NewTLMessagesGetOnlines() *TLMessagesGetOnlines {
	return &TLMessagesGetOnlines{
		Cmd:       Cmd_MessagesGetOnlines,
		ClassName: ClassMessagesGetOnlines,
	}
	//return &TLMessagesGetOnlines{}
}

func NewTLMessagesGetPeerDialogs() *TLMessagesGetPeerDialogs {
	return &TLMessagesGetPeerDialogs{
		Cmd:       Cmd_MessagesGetPeerDialogs,
		ClassName: ClassMessagesGetPeerDialogs,
	}
	//return &TLMessagesGetPeerDialogs{}
}

func NewTLMessagesGetPeerSettings() *TLMessagesGetPeerSettings {
	return &TLMessagesGetPeerSettings{
		Cmd:       Cmd_MessagesGetPeerSettings,
		ClassName: ClassMessagesGetPeerSettings,
	}
	//return &TLMessagesGetPeerSettings{}
}

func NewTLMessagesGetPinnedDialogs() *TLMessagesGetPinnedDialogs {
	return &TLMessagesGetPinnedDialogs{
		Cmd:       Cmd_MessagesGetPinnedDialogs,
		ClassName: ClassMessagesGetPinnedDialogs,
	}
	//return &TLMessagesGetPinnedDialogs{}
}

func NewTLMessagesGetPollResults() *TLMessagesGetPollResults {
	return &TLMessagesGetPollResults{
		Cmd:       Cmd_MessagesGetPollResults,
		ClassName: ClassMessagesGetPollResults,
	}
	//return &TLMessagesGetPollResults{}
}

func NewTLMessagesGetPollVotes() *TLMessagesGetPollVotes {
	return &TLMessagesGetPollVotes{
		Cmd:       Cmd_MessagesGetPollVotes,
		ClassName: ClassMessagesGetPollVotes,
	}
	//return &TLMessagesGetPollVotes{}
}

func NewTLMessagesGetRecentLocations() *TLMessagesGetRecentLocations {
	return &TLMessagesGetRecentLocations{
		Cmd:       Cmd_MessagesGetRecentLocations,
		ClassName: ClassMessagesGetRecentLocations,
	}
	//return &TLMessagesGetRecentLocations{}
}

func NewTLMessagesGetRecentStickers() *TLMessagesGetRecentStickers {
	return &TLMessagesGetRecentStickers{
		Cmd:       Cmd_MessagesGetRecentStickers,
		ClassName: ClassMessagesGetRecentStickers,
	}
	//return &TLMessagesGetRecentStickers{}
}

func NewTLMessagesGetReplies() *TLMessagesGetReplies {
	return &TLMessagesGetReplies{
		Cmd:       Cmd_MessagesGetReplies,
		ClassName: ClassMessagesGetReplies,
	}
	//return &TLMessagesGetReplies{}
}

func NewTLMessagesGetSavedGifs() *TLMessagesGetSavedGifs {
	return &TLMessagesGetSavedGifs{
		Cmd:       Cmd_MessagesGetSavedGifs,
		ClassName: ClassMessagesGetSavedGifs,
	}
	//return &TLMessagesGetSavedGifs{}
}

func NewTLMessagesGetScheduledHistory() *TLMessagesGetScheduledHistory {
	return &TLMessagesGetScheduledHistory{
		Cmd:       Cmd_MessagesGetScheduledHistory,
		ClassName: ClassMessagesGetScheduledHistory,
	}
	//return &TLMessagesGetScheduledHistory{}
}

func NewTLMessagesGetScheduledMessages() *TLMessagesGetScheduledMessages {
	return &TLMessagesGetScheduledMessages{
		Cmd:       Cmd_MessagesGetScheduledMessages,
		ClassName: ClassMessagesGetScheduledMessages,
	}
	//return &TLMessagesGetScheduledMessages{}
}

func NewTLMessagesGetSearchCounters() *TLMessagesGetSearchCounters {
	return &TLMessagesGetSearchCounters{
		Cmd:       Cmd_MessagesGetSearchCounters,
		ClassName: ClassMessagesGetSearchCounters,
	}
	//return &TLMessagesGetSearchCounters{}
}

func NewTLMessagesGetSplitRanges() *TLMessagesGetSplitRanges {
	return &TLMessagesGetSplitRanges{
		Cmd:       Cmd_MessagesGetSplitRanges,
		ClassName: ClassMessagesGetSplitRanges,
	}
	//return &TLMessagesGetSplitRanges{}
}

func NewTLMessagesGetStatsURL() *TLMessagesGetStatsURL {
	return &TLMessagesGetStatsURL{
		Cmd:       Cmd_MessagesGetStatsURL,
		ClassName: ClassMessagesGetStatsURL,
	}
	//return &TLMessagesGetStatsURL{}
}

func NewTLMessagesGetStickerSet() *TLMessagesGetStickerSet {
	return &TLMessagesGetStickerSet{
		Cmd:       Cmd_MessagesGetStickerSet,
		ClassName: ClassMessagesGetStickerSet,
	}
	//return &TLMessagesGetStickerSet{}
}

func NewTLMessagesGetStickers() *TLMessagesGetStickers {
	return &TLMessagesGetStickers{
		Cmd:       Cmd_MessagesGetStickers,
		ClassName: ClassMessagesGetStickers,
	}
	//return &TLMessagesGetStickers{}
}

func NewTLMessagesGetSuggestedDialogFilters() *TLMessagesGetSuggestedDialogFilters {
	return &TLMessagesGetSuggestedDialogFilters{
		Cmd:       Cmd_MessagesGetSuggestedDialogFilters,
		ClassName: ClassMessagesGetSuggestedDialogFilters,
	}
	//return &TLMessagesGetSuggestedDialogFilters{}
}

func NewTLMessagesGetUnreadMentions() *TLMessagesGetUnreadMentions {
	return &TLMessagesGetUnreadMentions{
		Cmd:       Cmd_MessagesGetUnreadMentions,
		ClassName: ClassMessagesGetUnreadMentions,
	}
	//return &TLMessagesGetUnreadMentions{}
}

func NewTLMessagesGetWebPage() *TLMessagesGetWebPage {
	return &TLMessagesGetWebPage{
		Cmd:       Cmd_MessagesGetWebPage,
		ClassName: ClassMessagesGetWebPage,
	}
	//return &TLMessagesGetWebPage{}
}

func NewTLMessagesGetWebPagePreview() *TLMessagesGetWebPagePreview {
	return &TLMessagesGetWebPagePreview{
		Cmd:       Cmd_MessagesGetWebPagePreview,
		ClassName: ClassMessagesGetWebPagePreview,
	}
	//return &TLMessagesGetWebPagePreview{}
}

func NewTLMessagesHidePeerSettingsBar() *TLMessagesHidePeerSettingsBar {
	return &TLMessagesHidePeerSettingsBar{
		Cmd:       Cmd_MessagesHidePeerSettingsBar,
		ClassName: ClassMessagesHidePeerSettingsBar,
	}
	//return &TLMessagesHidePeerSettingsBar{}
}

func NewTLMessagesHideReportSpam() *TLMessagesHideReportSpam {
	return &TLMessagesHideReportSpam{
		Cmd:       Cmd_MessagesHideReportSpam,
		ClassName: ClassMessagesHideReportSpam,
	}
	//return &TLMessagesHideReportSpam{}
}

func NewTLMessagesImportChatInvite() *TLMessagesImportChatInvite {
	return &TLMessagesImportChatInvite{
		Cmd:       Cmd_MessagesImportChatInvite,
		ClassName: ClassMessagesImportChatInvite,
	}
	//return &TLMessagesImportChatInvite{}
}

func NewTLMessagesInstallStickerSet() *TLMessagesInstallStickerSet {
	return &TLMessagesInstallStickerSet{
		Cmd:       Cmd_MessagesInstallStickerSet,
		ClassName: ClassMessagesInstallStickerSet,
	}
	//return &TLMessagesInstallStickerSet{}
}

func NewTLMessagesMarkDialogUnread() *TLMessagesMarkDialogUnread {
	return &TLMessagesMarkDialogUnread{
		Cmd:       Cmd_MessagesMarkDialogUnread,
		ClassName: ClassMessagesMarkDialogUnread,
	}
	//return &TLMessagesMarkDialogUnread{}
}

func NewTLMessagesMigrateChat() *TLMessagesMigrateChat {
	return &TLMessagesMigrateChat{
		Cmd:       Cmd_MessagesMigrateChat,
		ClassName: ClassMessagesMigrateChat,
	}
	//return &TLMessagesMigrateChat{}
}

func NewTLMessagesReadDiscussion() *TLMessagesReadDiscussion {
	return &TLMessagesReadDiscussion{
		Cmd:       Cmd_MessagesReadDiscussion,
		ClassName: ClassMessagesReadDiscussion,
	}
	//return &TLMessagesReadDiscussion{}
}

func NewTLMessagesReadEncryptedHistory() *TLMessagesReadEncryptedHistory {
	return &TLMessagesReadEncryptedHistory{
		Cmd:       Cmd_MessagesReadEncryptedHistory,
		ClassName: ClassMessagesReadEncryptedHistory,
	}
	//return &TLMessagesReadEncryptedHistory{}
}

func NewTLMessagesReadFeaturedStickers() *TLMessagesReadFeaturedStickers {
	return &TLMessagesReadFeaturedStickers{
		Cmd:       Cmd_MessagesReadFeaturedStickers,
		ClassName: ClassMessagesReadFeaturedStickers,
	}
	//return &TLMessagesReadFeaturedStickers{}
}

func NewTLMessagesReadHistory() *TLMessagesReadHistory {
	return &TLMessagesReadHistory{
		Cmd:       Cmd_MessagesReadHistory,
		ClassName: ClassMessagesReadHistory,
	}
	//return &TLMessagesReadHistory{}
}

func NewTLMessagesReadMentions() *TLMessagesReadMentions {
	return &TLMessagesReadMentions{
		Cmd:       Cmd_MessagesReadMentions,
		ClassName: ClassMessagesReadMentions,
	}
	//return &TLMessagesReadMentions{}
}

func NewTLMessagesReadMessageContents() *TLMessagesReadMessageContents {
	return &TLMessagesReadMessageContents{
		Cmd:       Cmd_MessagesReadMessageContents,
		ClassName: ClassMessagesReadMessageContents,
	}
	//return &TLMessagesReadMessageContents{}
}

func NewTLMessagesReceivedMessages() *TLMessagesReceivedMessages {
	return &TLMessagesReceivedMessages{
		Cmd:       Cmd_MessagesReceivedMessages,
		ClassName: ClassMessagesReceivedMessages,
	}
	//return &TLMessagesReceivedMessages{}
}

func NewTLMessagesReceivedQueue() *TLMessagesReceivedQueue {
	return &TLMessagesReceivedQueue{
		Cmd:       Cmd_MessagesReceivedQueue,
		ClassName: ClassMessagesReceivedQueue,
	}
	//return &TLMessagesReceivedQueue{}
}

func NewTLMessagesReorderPinnedDialogs() *TLMessagesReorderPinnedDialogs {
	return &TLMessagesReorderPinnedDialogs{
		Cmd:       Cmd_MessagesReorderPinnedDialogs,
		ClassName: ClassMessagesReorderPinnedDialogs,
	}
	//return &TLMessagesReorderPinnedDialogs{}
}

func NewTLMessagesReorderStickerSets() *TLMessagesReorderStickerSets {
	return &TLMessagesReorderStickerSets{
		Cmd:       Cmd_MessagesReorderStickerSets,
		ClassName: ClassMessagesReorderStickerSets,
	}
	//return &TLMessagesReorderStickerSets{}
}

func NewTLMessagesReport() *TLMessagesReport {
	return &TLMessagesReport{
		Cmd:       Cmd_MessagesReport,
		ClassName: ClassMessagesReport,
	}
	//return &TLMessagesReport{}
}

func NewTLMessagesReportEncryptedSpam() *TLMessagesReportEncryptedSpam {
	return &TLMessagesReportEncryptedSpam{
		Cmd:       Cmd_MessagesReportEncryptedSpam,
		ClassName: ClassMessagesReportEncryptedSpam,
	}
	//return &TLMessagesReportEncryptedSpam{}
}

func NewTLMessagesReportSpam() *TLMessagesReportSpam {
	return &TLMessagesReportSpam{
		Cmd:       Cmd_MessagesReportSpam,
		ClassName: ClassMessagesReportSpam,
	}
	//return &TLMessagesReportSpam{}
}

func NewTLMessagesRequestEncryption() *TLMessagesRequestEncryption {
	return &TLMessagesRequestEncryption{
		Cmd:       Cmd_MessagesRequestEncryption,
		ClassName: ClassMessagesRequestEncryption,
	}
	//return &TLMessagesRequestEncryption{}
}

func NewTLMessagesRequestUrlAuth() *TLMessagesRequestUrlAuth {
	return &TLMessagesRequestUrlAuth{
		Cmd:       Cmd_MessagesRequestUrlAuth,
		ClassName: ClassMessagesRequestUrlAuth,
	}
	//return &TLMessagesRequestUrlAuth{}
}

func NewTLMessagesSaveDraft() *TLMessagesSaveDraft {
	return &TLMessagesSaveDraft{
		Cmd:       Cmd_MessagesSaveDraft,
		ClassName: ClassMessagesSaveDraft,
	}
	//return &TLMessagesSaveDraft{}
}

func NewTLMessagesSaveGif() *TLMessagesSaveGif {
	return &TLMessagesSaveGif{
		Cmd:       Cmd_MessagesSaveGif,
		ClassName: ClassMessagesSaveGif,
	}
	//return &TLMessagesSaveGif{}
}

func NewTLMessagesSaveRecentSticker() *TLMessagesSaveRecentSticker {
	return &TLMessagesSaveRecentSticker{
		Cmd:       Cmd_MessagesSaveRecentSticker,
		ClassName: ClassMessagesSaveRecentSticker,
	}
	//return &TLMessagesSaveRecentSticker{}
}

func NewTLMessagesSearch() *TLMessagesSearch {
	return &TLMessagesSearch{
		Cmd:       Cmd_MessagesSearch,
		ClassName: ClassMessagesSearch,
	}
	//return &TLMessagesSearch{}
}

func NewTLMessagesSearchGifs() *TLMessagesSearchGifs {
	return &TLMessagesSearchGifs{
		Cmd:       Cmd_MessagesSearchGifs,
		ClassName: ClassMessagesSearchGifs,
	}
	//return &TLMessagesSearchGifs{}
}

func NewTLMessagesSearchGlobal() *TLMessagesSearchGlobal {
	return &TLMessagesSearchGlobal{
		Cmd:       Cmd_MessagesSearchGlobal,
		ClassName: ClassMessagesSearchGlobal,
	}
	//return &TLMessagesSearchGlobal{}
}

func NewTLMessagesSearchStickerSets() *TLMessagesSearchStickerSets {
	return &TLMessagesSearchStickerSets{
		Cmd:       Cmd_MessagesSearchStickerSets,
		ClassName: ClassMessagesSearchStickerSets,
	}
	//return &TLMessagesSearchStickerSets{}
}

func NewTLMessagesSendBroadcast() *TLMessagesSendBroadcast {
	return &TLMessagesSendBroadcast{
		Cmd:       Cmd_MessagesSendBroadcast,
		ClassName: ClassMessagesSendBroadcast,
	}
	//return &TLMessagesSendBroadcast{}
}

func NewTLMessagesSendEncrypted() *TLMessagesSendEncrypted {
	return &TLMessagesSendEncrypted{
		Cmd:       Cmd_MessagesSendEncrypted,
		ClassName: ClassMessagesSendEncrypted,
	}
	//return &TLMessagesSendEncrypted{}
}

func NewTLMessagesSendEncryptedFile() *TLMessagesSendEncryptedFile {
	return &TLMessagesSendEncryptedFile{
		Cmd:       Cmd_MessagesSendEncryptedFile,
		ClassName: ClassMessagesSendEncryptedFile,
	}
	//return &TLMessagesSendEncryptedFile{}
}

func NewTLMessagesSendEncryptedService() *TLMessagesSendEncryptedService {
	return &TLMessagesSendEncryptedService{
		Cmd:       Cmd_MessagesSendEncryptedService,
		ClassName: ClassMessagesSendEncryptedService,
	}
	//return &TLMessagesSendEncryptedService{}
}

func NewTLMessagesSendInlineBotResult() *TLMessagesSendInlineBotResult {
	return &TLMessagesSendInlineBotResult{
		Cmd:       Cmd_MessagesSendInlineBotResult,
		ClassName: ClassMessagesSendInlineBotResult,
	}
	//return &TLMessagesSendInlineBotResult{}
}

func NewTLMessagesSendMedia() *TLMessagesSendMedia {
	return &TLMessagesSendMedia{
		Cmd:       Cmd_MessagesSendMedia,
		ClassName: ClassMessagesSendMedia,
	}
	//return &TLMessagesSendMedia{}
}

func NewTLMessagesSendMessage() *TLMessagesSendMessage {
	return &TLMessagesSendMessage{
		Cmd:       Cmd_MessagesSendMessage,
		ClassName: ClassMessagesSendMessage,
	}
	//return &TLMessagesSendMessage{}
}

func NewTLMessagesSendMultiMedia() *TLMessagesSendMultiMedia {
	return &TLMessagesSendMultiMedia{
		Cmd:       Cmd_MessagesSendMultiMedia,
		ClassName: ClassMessagesSendMultiMedia,
	}
	//return &TLMessagesSendMultiMedia{}
}

func NewTLMessagesSendScheduledMessages() *TLMessagesSendScheduledMessages {
	return &TLMessagesSendScheduledMessages{
		Cmd:       Cmd_MessagesSendScheduledMessages,
		ClassName: ClassMessagesSendScheduledMessages,
	}
	//return &TLMessagesSendScheduledMessages{}
}

func NewTLMessagesSendScreenshotNotification() *TLMessagesSendScreenshotNotification {
	return &TLMessagesSendScreenshotNotification{
		Cmd:       Cmd_MessagesSendScreenshotNotification,
		ClassName: ClassMessagesSendScreenshotNotification,
	}
	//return &TLMessagesSendScreenshotNotification{}
}

func NewTLMessagesSendVote() *TLMessagesSendVote {
	return &TLMessagesSendVote{
		Cmd:       Cmd_MessagesSendVote,
		ClassName: ClassMessagesSendVote,
	}
	//return &TLMessagesSendVote{}
}

func NewTLMessagesSetBotCallbackAnswer() *TLMessagesSetBotCallbackAnswer {
	return &TLMessagesSetBotCallbackAnswer{
		Cmd:       Cmd_MessagesSetBotCallbackAnswer,
		ClassName: ClassMessagesSetBotCallbackAnswer,
	}
	//return &TLMessagesSetBotCallbackAnswer{}
}

func NewTLMessagesSetBotPrecheckoutResults() *TLMessagesSetBotPrecheckoutResults {
	return &TLMessagesSetBotPrecheckoutResults{
		Cmd:       Cmd_MessagesSetBotPrecheckoutResults,
		ClassName: ClassMessagesSetBotPrecheckoutResults,
	}
	//return &TLMessagesSetBotPrecheckoutResults{}
}

func NewTLMessagesSetBotShippingResults() *TLMessagesSetBotShippingResults {
	return &TLMessagesSetBotShippingResults{
		Cmd:       Cmd_MessagesSetBotShippingResults,
		ClassName: ClassMessagesSetBotShippingResults,
	}
	//return &TLMessagesSetBotShippingResults{}
}

func NewTLMessagesSetEncryptedTyping() *TLMessagesSetEncryptedTyping {
	return &TLMessagesSetEncryptedTyping{
		Cmd:       Cmd_MessagesSetEncryptedTyping,
		ClassName: ClassMessagesSetEncryptedTyping,
	}
	//return &TLMessagesSetEncryptedTyping{}
}

func NewTLMessagesSetGameScore() *TLMessagesSetGameScore {
	return &TLMessagesSetGameScore{
		Cmd:       Cmd_MessagesSetGameScore,
		ClassName: ClassMessagesSetGameScore,
	}
	//return &TLMessagesSetGameScore{}
}

func NewTLMessagesSetInlineBotResults() *TLMessagesSetInlineBotResults {
	return &TLMessagesSetInlineBotResults{
		Cmd:       Cmd_MessagesSetInlineBotResults,
		ClassName: ClassMessagesSetInlineBotResults,
	}
	//return &TLMessagesSetInlineBotResults{}
}

func NewTLMessagesSetInlineGameScore() *TLMessagesSetInlineGameScore {
	return &TLMessagesSetInlineGameScore{
		Cmd:       Cmd_MessagesSetInlineGameScore,
		ClassName: ClassMessagesSetInlineGameScore,
	}
	//return &TLMessagesSetInlineGameScore{}
}

func NewTLMessagesSetTyping() *TLMessagesSetTyping {
	return &TLMessagesSetTyping{
		Cmd:       Cmd_MessagesSetTyping,
		ClassName: ClassMessagesSetTyping,
	}
	//return &TLMessagesSetTyping{}
}

func NewTLMessagesStartBot() *TLMessagesStartBot {
	return &TLMessagesStartBot{
		Cmd:       Cmd_MessagesStartBot,
		ClassName: ClassMessagesStartBot,
	}
	//return &TLMessagesStartBot{}
}

func NewTLMessagesToggleChatAdmins() *TLMessagesToggleChatAdmins {
	return &TLMessagesToggleChatAdmins{
		Cmd:       Cmd_MessagesToggleChatAdmins,
		ClassName: ClassMessagesToggleChatAdmins,
	}
	//return &TLMessagesToggleChatAdmins{}
}

func NewTLMessagesToggleDialogPin() *TLMessagesToggleDialogPin {
	return &TLMessagesToggleDialogPin{
		Cmd:       Cmd_MessagesToggleDialogPin,
		ClassName: ClassMessagesToggleDialogPin,
	}
	//return &TLMessagesToggleDialogPin{}
}

func NewTLMessagesToggleStickerSets() *TLMessagesToggleStickerSets {
	return &TLMessagesToggleStickerSets{
		Cmd:       Cmd_MessagesToggleStickerSets,
		ClassName: ClassMessagesToggleStickerSets,
	}
	//return &TLMessagesToggleStickerSets{}
}

func NewTLMessagesUninstallStickerSet() *TLMessagesUninstallStickerSet {
	return &TLMessagesUninstallStickerSet{
		Cmd:       Cmd_MessagesUninstallStickerSet,
		ClassName: ClassMessagesUninstallStickerSet,
	}
	//return &TLMessagesUninstallStickerSet{}
}

func NewTLMessagesUnpinAllMessages() *TLMessagesUnpinAllMessages {
	return &TLMessagesUnpinAllMessages{
		Cmd:       Cmd_MessagesUnpinAllMessages,
		ClassName: ClassMessagesUnpinAllMessages,
	}
	//return &TLMessagesUnpinAllMessages{}
}

func NewTLMessagesUpdateDialogFilter() *TLMessagesUpdateDialogFilter {
	return &TLMessagesUpdateDialogFilter{
		Cmd:       Cmd_MessagesUpdateDialogFilter,
		ClassName: ClassMessagesUpdateDialogFilter,
	}
	//return &TLMessagesUpdateDialogFilter{}
}

func NewTLMessagesUpdateDialogFiltersOrder() *TLMessagesUpdateDialogFiltersOrder {
	return &TLMessagesUpdateDialogFiltersOrder{
		Cmd:       Cmd_MessagesUpdateDialogFiltersOrder,
		ClassName: ClassMessagesUpdateDialogFiltersOrder,
	}
	//return &TLMessagesUpdateDialogFiltersOrder{}
}

func NewTLMessagesUpdatePinnedMessage() *TLMessagesUpdatePinnedMessage {
	return &TLMessagesUpdatePinnedMessage{
		Cmd:       Cmd_MessagesUpdatePinnedMessage,
		ClassName: ClassMessagesUpdatePinnedMessage,
	}
	//return &TLMessagesUpdatePinnedMessage{}
}

func NewTLMessagesUploadEncryptedFile() *TLMessagesUploadEncryptedFile {
	return &TLMessagesUploadEncryptedFile{
		Cmd:       Cmd_MessagesUploadEncryptedFile,
		ClassName: ClassMessagesUploadEncryptedFile,
	}
	//return &TLMessagesUploadEncryptedFile{}
}

func NewTLMessagesUploadMedia() *TLMessagesUploadMedia {
	return &TLMessagesUploadMedia{
		Cmd:       Cmd_MessagesUploadMedia,
		ClassName: ClassMessagesUploadMedia,
	}
	//return &TLMessagesUploadMedia{}
}

func NewTLPaymentsClearSavedInfo() *TLPaymentsClearSavedInfo {
	return &TLPaymentsClearSavedInfo{
		Cmd:       Cmd_PaymentsClearSavedInfo,
		ClassName: ClassPaymentsClearSavedInfo,
	}
	//return &TLPaymentsClearSavedInfo{}
}

func NewTLPaymentsGetBankCardData() *TLPaymentsGetBankCardData {
	return &TLPaymentsGetBankCardData{
		Cmd:       Cmd_PaymentsGetBankCardData,
		ClassName: ClassPaymentsGetBankCardData,
	}
	//return &TLPaymentsGetBankCardData{}
}

func NewTLPaymentsGetPaymentForm() *TLPaymentsGetPaymentForm {
	return &TLPaymentsGetPaymentForm{
		Cmd:       Cmd_PaymentsGetPaymentForm,
		ClassName: ClassPaymentsGetPaymentForm,
	}
	//return &TLPaymentsGetPaymentForm{}
}

func NewTLPaymentsGetPaymentReceipt() *TLPaymentsGetPaymentReceipt {
	return &TLPaymentsGetPaymentReceipt{
		Cmd:       Cmd_PaymentsGetPaymentReceipt,
		ClassName: ClassPaymentsGetPaymentReceipt,
	}
	//return &TLPaymentsGetPaymentReceipt{}
}

func NewTLPaymentsGetSavedInfo() *TLPaymentsGetSavedInfo {
	return &TLPaymentsGetSavedInfo{
		Cmd:       Cmd_PaymentsGetSavedInfo,
		ClassName: ClassPaymentsGetSavedInfo,
	}
	//return &TLPaymentsGetSavedInfo{}
}

func NewTLPaymentsSendPaymentForm() *TLPaymentsSendPaymentForm {
	return &TLPaymentsSendPaymentForm{
		Cmd:       Cmd_PaymentsSendPaymentForm,
		ClassName: ClassPaymentsSendPaymentForm,
	}
	//return &TLPaymentsSendPaymentForm{}
}

func NewTLPaymentsValidateRequestedInfo() *TLPaymentsValidateRequestedInfo {
	return &TLPaymentsValidateRequestedInfo{
		Cmd:       Cmd_PaymentsValidateRequestedInfo,
		ClassName: ClassPaymentsValidateRequestedInfo,
	}
	//return &TLPaymentsValidateRequestedInfo{}
}

func NewTLPhoneAcceptCall() *TLPhoneAcceptCall {
	return &TLPhoneAcceptCall{
		Cmd:       Cmd_PhoneAcceptCall,
		ClassName: ClassPhoneAcceptCall,
	}
	//return &TLPhoneAcceptCall{}
}

func NewTLPhoneCheckGroupCall() *TLPhoneCheckGroupCall {
	return &TLPhoneCheckGroupCall{
		Cmd:       Cmd_PhoneCheckGroupCall,
		ClassName: ClassPhoneCheckGroupCall,
	}
	//return &TLPhoneCheckGroupCall{}
}

func NewTLPhoneConfirmCall() *TLPhoneConfirmCall {
	return &TLPhoneConfirmCall{
		Cmd:       Cmd_PhoneConfirmCall,
		ClassName: ClassPhoneConfirmCall,
	}
	//return &TLPhoneConfirmCall{}
}

func NewTLPhoneCreateGroupCall() *TLPhoneCreateGroupCall {
	return &TLPhoneCreateGroupCall{
		Cmd:       Cmd_PhoneCreateGroupCall,
		ClassName: ClassPhoneCreateGroupCall,
	}
	//return &TLPhoneCreateGroupCall{}
}

func NewTLPhoneDiscardCall() *TLPhoneDiscardCall {
	return &TLPhoneDiscardCall{
		Cmd:       Cmd_PhoneDiscardCall,
		ClassName: ClassPhoneDiscardCall,
	}
	//return &TLPhoneDiscardCall{}
}

func NewTLPhoneDiscardGroupCall() *TLPhoneDiscardGroupCall {
	return &TLPhoneDiscardGroupCall{
		Cmd:       Cmd_PhoneDiscardGroupCall,
		ClassName: ClassPhoneDiscardGroupCall,
	}
	//return &TLPhoneDiscardGroupCall{}
}

func NewTLPhoneEditGroupCallMember() *TLPhoneEditGroupCallMember {
	return &TLPhoneEditGroupCallMember{
		Cmd:       Cmd_PhoneEditGroupCallMember,
		ClassName: ClassPhoneEditGroupCallMember,
	}
	//return &TLPhoneEditGroupCallMember{}
}

func NewTLPhoneGetCallConfig() *TLPhoneGetCallConfig {
	return &TLPhoneGetCallConfig{
		Cmd:       Cmd_PhoneGetCallConfig,
		ClassName: ClassPhoneGetCallConfig,
	}
	//return &TLPhoneGetCallConfig{}
}

func NewTLPhoneGetGroupCall() *TLPhoneGetGroupCall {
	return &TLPhoneGetGroupCall{
		Cmd:       Cmd_PhoneGetGroupCall,
		ClassName: ClassPhoneGetGroupCall,
	}
	//return &TLPhoneGetGroupCall{}
}

func NewTLPhoneGetGroupParticipants() *TLPhoneGetGroupParticipants {
	return &TLPhoneGetGroupParticipants{
		Cmd:       Cmd_PhoneGetGroupParticipants,
		ClassName: ClassPhoneGetGroupParticipants,
	}
	//return &TLPhoneGetGroupParticipants{}
}

func NewTLPhoneInviteToGroupCall() *TLPhoneInviteToGroupCall {
	return &TLPhoneInviteToGroupCall{
		Cmd:       Cmd_PhoneInviteToGroupCall,
		ClassName: ClassPhoneInviteToGroupCall,
	}
	//return &TLPhoneInviteToGroupCall{}
}

func NewTLPhoneJoinGroupCall() *TLPhoneJoinGroupCall {
	return &TLPhoneJoinGroupCall{
		Cmd:       Cmd_PhoneJoinGroupCall,
		ClassName: ClassPhoneJoinGroupCall,
	}
	//return &TLPhoneJoinGroupCall{}
}

func NewTLPhoneLeaveGroupCall() *TLPhoneLeaveGroupCall {
	return &TLPhoneLeaveGroupCall{
		Cmd:       Cmd_PhoneLeaveGroupCall,
		ClassName: ClassPhoneLeaveGroupCall,
	}
	//return &TLPhoneLeaveGroupCall{}
}

func NewTLPhoneReceivedCall() *TLPhoneReceivedCall {
	return &TLPhoneReceivedCall{
		Cmd:       Cmd_PhoneReceivedCall,
		ClassName: ClassPhoneReceivedCall,
	}
	//return &TLPhoneReceivedCall{}
}

func NewTLPhoneRequestCall() *TLPhoneRequestCall {
	return &TLPhoneRequestCall{
		Cmd:       Cmd_PhoneRequestCall,
		ClassName: ClassPhoneRequestCall,
	}
	//return &TLPhoneRequestCall{}
}

func NewTLPhoneSaveCallDebug() *TLPhoneSaveCallDebug {
	return &TLPhoneSaveCallDebug{
		Cmd:       Cmd_PhoneSaveCallDebug,
		ClassName: ClassPhoneSaveCallDebug,
	}
	//return &TLPhoneSaveCallDebug{}
}

func NewTLPhoneSendSignalingData() *TLPhoneSendSignalingData {
	return &TLPhoneSendSignalingData{
		Cmd:       Cmd_PhoneSendSignalingData,
		ClassName: ClassPhoneSendSignalingData,
	}
	//return &TLPhoneSendSignalingData{}
}

func NewTLPhoneSetCallRating() *TLPhoneSetCallRating {
	return &TLPhoneSetCallRating{
		Cmd:       Cmd_PhoneSetCallRating,
		ClassName: ClassPhoneSetCallRating,
	}
	//return &TLPhoneSetCallRating{}
}

func NewTLPhoneToggleGroupCallSettings() *TLPhoneToggleGroupCallSettings {
	return &TLPhoneToggleGroupCallSettings{
		Cmd:       Cmd_PhoneToggleGroupCallSettings,
		ClassName: ClassPhoneToggleGroupCallSettings,
	}
	//return &TLPhoneToggleGroupCallSettings{}
}

func NewTLPhotosDeletePhotos() *TLPhotosDeletePhotos {
	return &TLPhotosDeletePhotos{
		Cmd:       Cmd_PhotosDeletePhotos,
		ClassName: ClassPhotosDeletePhotos,
	}
	//return &TLPhotosDeletePhotos{}
}

func NewTLPhotosGetUserPhotos() *TLPhotosGetUserPhotos {
	return &TLPhotosGetUserPhotos{
		Cmd:       Cmd_PhotosGetUserPhotos,
		ClassName: ClassPhotosGetUserPhotos,
	}
	//return &TLPhotosGetUserPhotos{}
}

func NewTLPhotosUpdateProfilePhoto() *TLPhotosUpdateProfilePhoto {
	return &TLPhotosUpdateProfilePhoto{
		Cmd:       Cmd_PhotosUpdateProfilePhoto,
		ClassName: ClassPhotosUpdateProfilePhoto,
	}
	//return &TLPhotosUpdateProfilePhoto{}
}

func NewTLPhotosUploadProfilePhoto() *TLPhotosUploadProfilePhoto {
	return &TLPhotosUploadProfilePhoto{
		Cmd:       Cmd_PhotosUploadProfilePhoto,
		ClassName: ClassPhotosUploadProfilePhoto,
	}
	//return &TLPhotosUploadProfilePhoto{}
}

func NewTLPing() *TLPing {
	return &TLPing{
		Cmd:       Cmd_Ping,
		ClassName: ClassPing,
	}
	//return &TLPing{}
}

func NewTLPingDelayDisconnect() *TLPingDelayDisconnect {
	return &TLPingDelayDisconnect{
		Cmd:       Cmd_PingDelayDisconnect,
		ClassName: ClassPingDelayDisconnect,
	}
	//return &TLPingDelayDisconnect{}
}

func NewTLReqPq() *TLReqPq {
	return &TLReqPq{
		Cmd:       Cmd_ReqPq,
		ClassName: ClassReqPq,
	}
	//return &TLReqPq{}
}

func NewTLReqPqMulti() *TLReqPqMulti {
	return &TLReqPqMulti{
		Cmd:       Cmd_ReqPqMulti,
		ClassName: ClassReqPqMulti,
	}
	//return &TLReqPqMulti{}
}

func NewTLReq_DHParams() *TLReq_DHParams {
	return &TLReq_DHParams{
		Cmd:       Cmd_Req_DHParams,
		ClassName: ClassReq_DHParams,
	}
	//return &TLReq_DHParams{}
}

func NewTLRpcDropAnswer() *TLRpcDropAnswer {
	return &TLRpcDropAnswer{
		Cmd:       Cmd_RpcDropAnswer,
		ClassName: ClassRpcDropAnswer,
	}
	//return &TLRpcDropAnswer{}
}

func NewTLSetClient_DHParams() *TLSetClient_DHParams {
	return &TLSetClient_DHParams{
		Cmd:       Cmd_SetClient_DHParams,
		ClassName: ClassSetClient_DHParams,
	}
	//return &TLSetClient_DHParams{}
}

func NewTLStatsGetBroadcastStats() *TLStatsGetBroadcastStats {
	return &TLStatsGetBroadcastStats{
		Cmd:       Cmd_StatsGetBroadcastStats,
		ClassName: ClassStatsGetBroadcastStats,
	}
	//return &TLStatsGetBroadcastStats{}
}

func NewTLStatsGetMegagroupStats() *TLStatsGetMegagroupStats {
	return &TLStatsGetMegagroupStats{
		Cmd:       Cmd_StatsGetMegagroupStats,
		ClassName: ClassStatsGetMegagroupStats,
	}
	//return &TLStatsGetMegagroupStats{}
}

func NewTLStatsGetMessagePublicForwards() *TLStatsGetMessagePublicForwards {
	return &TLStatsGetMessagePublicForwards{
		Cmd:       Cmd_StatsGetMessagePublicForwards,
		ClassName: ClassStatsGetMessagePublicForwards,
	}
	//return &TLStatsGetMessagePublicForwards{}
}

func NewTLStatsGetMessageStats() *TLStatsGetMessageStats {
	return &TLStatsGetMessageStats{
		Cmd:       Cmd_StatsGetMessageStats,
		ClassName: ClassStatsGetMessageStats,
	}
	//return &TLStatsGetMessageStats{}
}

func NewTLStatsLoadAsyncGraph() *TLStatsLoadAsyncGraph {
	return &TLStatsLoadAsyncGraph{
		Cmd:       Cmd_StatsLoadAsyncGraph,
		ClassName: ClassStatsLoadAsyncGraph,
	}
	//return &TLStatsLoadAsyncGraph{}
}

func NewTLStickersAddStickerToSet() *TLStickersAddStickerToSet {
	return &TLStickersAddStickerToSet{
		Cmd:       Cmd_StickersAddStickerToSet,
		ClassName: ClassStickersAddStickerToSet,
	}
	//return &TLStickersAddStickerToSet{}
}

func NewTLStickersChangeStickerPosition() *TLStickersChangeStickerPosition {
	return &TLStickersChangeStickerPosition{
		Cmd:       Cmd_StickersChangeStickerPosition,
		ClassName: ClassStickersChangeStickerPosition,
	}
	//return &TLStickersChangeStickerPosition{}
}

func NewTLStickersCreateStickerSet() *TLStickersCreateStickerSet {
	return &TLStickersCreateStickerSet{
		Cmd:       Cmd_StickersCreateStickerSet,
		ClassName: ClassStickersCreateStickerSet,
	}
	//return &TLStickersCreateStickerSet{}
}

func NewTLStickersRemoveStickerFromSet() *TLStickersRemoveStickerFromSet {
	return &TLStickersRemoveStickerFromSet{
		Cmd:       Cmd_StickersRemoveStickerFromSet,
		ClassName: ClassStickersRemoveStickerFromSet,
	}
	//return &TLStickersRemoveStickerFromSet{}
}

func NewTLStickersSetStickerSetThumb() *TLStickersSetStickerSetThumb {
	return &TLStickersSetStickerSetThumb{
		Cmd:       Cmd_StickersSetStickerSetThumb,
		ClassName: ClassStickersSetStickerSetThumb,
	}
	//return &TLStickersSetStickerSetThumb{}
}

func NewTLUpdatesGetChannelDifference() *TLUpdatesGetChannelDifference {
	return &TLUpdatesGetChannelDifference{
		Cmd:       Cmd_UpdatesGetChannelDifference,
		ClassName: ClassUpdatesGetChannelDifference,
	}
	//return &TLUpdatesGetChannelDifference{}
}

func NewTLUpdatesGetDifference() *TLUpdatesGetDifference {
	return &TLUpdatesGetDifference{
		Cmd:       Cmd_UpdatesGetDifference,
		ClassName: ClassUpdatesGetDifference,
	}
	//return &TLUpdatesGetDifference{}
}

func NewTLUpdatesGetState() *TLUpdatesGetState {
	return &TLUpdatesGetState{
		Cmd:       Cmd_UpdatesGetState,
		ClassName: ClassUpdatesGetState,
	}
	//return &TLUpdatesGetState{}
}

func NewTLUploadGetCdnFile() *TLUploadGetCdnFile {
	return &TLUploadGetCdnFile{
		Cmd:       Cmd_UploadGetCdnFile,
		ClassName: ClassUploadGetCdnFile,
	}
	//return &TLUploadGetCdnFile{}
}

func NewTLUploadGetCdnFileHashes() *TLUploadGetCdnFileHashes {
	return &TLUploadGetCdnFileHashes{
		Cmd:       Cmd_UploadGetCdnFileHashes,
		ClassName: ClassUploadGetCdnFileHashes,
	}
	//return &TLUploadGetCdnFileHashes{}
}

func NewTLUploadGetFile() *TLUploadGetFile {
	return &TLUploadGetFile{
		Cmd:       Cmd_UploadGetFile,
		ClassName: ClassUploadGetFile,
	}
	//return &TLUploadGetFile{}
}

func NewTLUploadGetFileHashes() *TLUploadGetFileHashes {
	return &TLUploadGetFileHashes{
		Cmd:       Cmd_UploadGetFileHashes,
		ClassName: ClassUploadGetFileHashes,
	}
	//return &TLUploadGetFileHashes{}
}

func NewTLUploadGetWebFile() *TLUploadGetWebFile {
	return &TLUploadGetWebFile{
		Cmd:       Cmd_UploadGetWebFile,
		ClassName: ClassUploadGetWebFile,
	}
	//return &TLUploadGetWebFile{}
}

func NewTLUploadReuploadCdnFile() *TLUploadReuploadCdnFile {
	return &TLUploadReuploadCdnFile{
		Cmd:       Cmd_UploadReuploadCdnFile,
		ClassName: ClassUploadReuploadCdnFile,
	}
	//return &TLUploadReuploadCdnFile{}
}

func NewTLUploadSaveBigFilePart() *TLUploadSaveBigFilePart {
	return &TLUploadSaveBigFilePart{
		Cmd:       Cmd_UploadSaveBigFilePart,
		ClassName: ClassUploadSaveBigFilePart,
	}
	//return &TLUploadSaveBigFilePart{}
}

func NewTLUploadSaveFilePart() *TLUploadSaveFilePart {
	return &TLUploadSaveFilePart{
		Cmd:       Cmd_UploadSaveFilePart,
		ClassName: ClassUploadSaveFilePart,
	}
	//return &TLUploadSaveFilePart{}
}

func NewTLUsersGetFullUser() *TLUsersGetFullUser {
	return &TLUsersGetFullUser{
		Cmd:       Cmd_UsersGetFullUser,
		ClassName: ClassUsersGetFullUser,
	}
	//return &TLUsersGetFullUser{}
}

func NewTLUsersGetUsers() *TLUsersGetUsers {
	return &TLUsersGetUsers{
		Cmd:       Cmd_UsersGetUsers,
		ClassName: ClassUsersGetUsers,
	}
	//return &TLUsersGetUsers{}
}

func NewTLUsersSetSecureValueErrors() *TLUsersSetSecureValueErrors {
	return &TLUsersSetSecureValueErrors{
		Cmd:       Cmd_UsersSetSecureValueErrors,
		ClassName: ClassUsersSetSecureValueErrors,
	}
	//return &TLUsersSetSecureValueErrors{}
}

func NewTLWalletGetKeySecretSalt() *TLWalletGetKeySecretSalt {
	return &TLWalletGetKeySecretSalt{
		Cmd:       Cmd_WalletGetKeySecretSalt,
		ClassName: ClassWalletGetKeySecretSalt,
	}
	//return &TLWalletGetKeySecretSalt{}
}

func NewTLWalletSendLiteRequest() *TLWalletSendLiteRequest {
	return &TLWalletSendLiteRequest{
		Cmd:       Cmd_WalletSendLiteRequest,
		ClassName: ClassWalletSendLiteRequest,
	}
	//return &TLWalletSendLiteRequest{}
}
