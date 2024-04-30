/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2023/06/12 18:03
 * @File : codec_schema_encode.pb.go
 * @Desc :
 *
 */

package mtproto

import (
	"novachat_engine/pkg/log"
)

func (m *AccessPointRule) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccessPointRule:
		t := m.To_AccessPointRule()
		return t.Encode(layer)

	default:
		log.Errorf("AccessPointRule Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *AccountDaysTTL) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountDaysTTL:
		t := m.To_AccountDaysTTL()
		return t.Encode(layer)

	default:
		log.Errorf("AccountDaysTTL Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_AuthorizationForm) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountAuthorizationForm:
		t := m.To_AccountAuthorizationForm()
		return t.Encode(layer)

	default:
		log.Errorf("Account_AuthorizationForm Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_Authorizations) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountAuthorizations:
		t := m.To_AccountAuthorizations()
		return t.Encode(layer)

	default:
		log.Errorf("Account_Authorizations Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_AutoDownloadSettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountAutoDownloadSettings:
		t := m.To_AccountAutoDownloadSettings()
		return t.Encode(layer)

	default:
		log.Errorf("Account_AutoDownloadSettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_ContentSettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountContentSettings:
		t := m.To_AccountContentSettings()
		return t.Encode(layer)

	default:
		log.Errorf("Account_ContentSettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_Password) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountNoPassword:
		t := m.To_AccountNoPassword()
		return t.Encode(layer)
	case Cmd_AccountPassword:
		t := m.To_AccountPassword()
		return t.Encode(layer)

	case Cmd_AccountPasswordad2641f8:
		t := m.To_AccountPassword()
		return t.Encode(layer)

	default:
		log.Errorf("Account_Password Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_PasswordInputSettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountPasswordInputSettings:
		t := m.To_AccountPasswordInputSettings()
		return t.Encode(layer)

	case Cmd_AccountPasswordInputSettingsc23727c9:
		t := m.To_AccountPasswordInputSettings()
		return t.Encode(layer)

	default:
		log.Errorf("Account_PasswordInputSettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_PasswordSettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountPasswordSettings:
		t := m.To_AccountPasswordSettings()
		return t.Encode(layer)

	case Cmd_AccountPasswordSettings9a5c33e5:
		t := m.To_AccountPasswordSettings()
		return t.Encode(layer)

	default:
		log.Errorf("Account_PasswordSettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_PrivacyRules) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountPrivacyRules:
		t := m.To_AccountPrivacyRules()
		return t.Encode(layer)

	case Cmd_AccountPrivacyRules50a04e45:
		t := m.To_AccountPrivacyRules()
		return t.Encode(layer)

	default:
		log.Errorf("Account_PrivacyRules Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_SentEmailCode) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountSentEmailCode:
		t := m.To_AccountSentEmailCode()
		return t.Encode(layer)

	default:
		log.Errorf("Account_SentEmailCode Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_Takeout) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountTakeout:
		t := m.To_AccountTakeout()
		return t.Encode(layer)

	default:
		log.Errorf("Account_Takeout Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_Themes) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountThemesNotModified:
		t := m.To_AccountThemesNotModified()
		return t.Encode(layer)
	case Cmd_AccountThemes:
		t := m.To_AccountThemes()
		return t.Encode(layer)

	default:
		log.Errorf("Account_Themes Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_TmpPassword) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountTmpPassword:
		t := m.To_AccountTmpPassword()
		return t.Encode(layer)

	default:
		log.Errorf("Account_TmpPassword Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_WallPapers) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountWallPapersNotModified:
		t := m.To_AccountWallPapersNotModified()
		return t.Encode(layer)
	case Cmd_AccountWallPapers:
		t := m.To_AccountWallPapers()
		return t.Encode(layer)

	default:
		log.Errorf("Account_WallPapers Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Account_WebAuthorizations) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AccountWebAuthorizations:
		t := m.To_AccountWebAuthorizations()
		return t.Encode(layer)

	default:
		log.Errorf("Account_WebAuthorizations Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Auth_Authorization) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AuthAuthorization:
		t := m.To_AuthAuthorization()
		return t.Encode(layer)
	case Cmd_AuthAuthorizationSignUpRequired:
		t := m.To_AuthAuthorizationSignUpRequired()
		return t.Encode(layer)

	case Cmd_AuthAuthorizationff036af1:
		t := m.To_AuthAuthorization()
		return t.Encode(layer)

	default:
		log.Errorf("Auth_Authorization Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Auth_CheckedPhone) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AuthCheckedPhone:
		t := m.To_AuthCheckedPhone()
		return t.Encode(layer)

	default:
		log.Errorf("Auth_CheckedPhone Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Auth_CodeType) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AuthCodeTypeSms:
		t := m.To_AuthCodeTypeSms()
		return t.Encode(layer)
	case Cmd_AuthCodeTypeCall:
		t := m.To_AuthCodeTypeCall()
		return t.Encode(layer)
	case Cmd_AuthCodeTypeFlashCall:
		t := m.To_AuthCodeTypeFlashCall()
		return t.Encode(layer)

	default:
		log.Errorf("Auth_CodeType Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Auth_ExportedAuthorization) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AuthExportedAuthorization:
		t := m.To_AuthExportedAuthorization()
		return t.Encode(layer)

	default:
		log.Errorf("Auth_ExportedAuthorization Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Auth_LoginToken) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AuthLoginToken:
		t := m.To_AuthLoginToken()
		return t.Encode(layer)
	case Cmd_AuthLoginTokenMigrateTo:
		t := m.To_AuthLoginTokenMigrateTo()
		return t.Encode(layer)
	case Cmd_AuthLoginTokenSuccess:
		t := m.To_AuthLoginTokenSuccess()
		return t.Encode(layer)

	default:
		log.Errorf("Auth_LoginToken Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Auth_PasswordRecovery) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AuthPasswordRecovery:
		t := m.To_AuthPasswordRecovery()
		return t.Encode(layer)

	default:
		log.Errorf("Auth_PasswordRecovery Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Auth_SentCode) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AuthSentCode:
		t := m.To_AuthSentCode()
		return t.Encode(layer)

	case Cmd_AuthSentCode38faab5f:
		t := m.To_AuthSentCode()
		return t.Encode(layer)

	default:
		log.Errorf("Auth_SentCode Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Auth_SentCodeType) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AuthSentCodeTypeApp:
		t := m.To_AuthSentCodeTypeApp()
		return t.Encode(layer)
	case Cmd_AuthSentCodeTypeSms:
		t := m.To_AuthSentCodeTypeSms()
		return t.Encode(layer)
	case Cmd_AuthSentCodeTypeCall:
		t := m.To_AuthSentCodeTypeCall()
		return t.Encode(layer)
	case Cmd_AuthSentCodeTypeFlashCall:
		t := m.To_AuthSentCodeTypeFlashCall()
		return t.Encode(layer)

	default:
		log.Errorf("Auth_SentCodeType Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Authorization) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Authorization:
		t := m.To_Authorization()
		return t.Encode(layer)

	case Cmd_Authorizationad01d61d:
		t := m.To_Authorization()
		return t.Encode(layer)

	default:
		log.Errorf("Authorization Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *AutoDownloadSettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_AutoDownloadSettings:
		t := m.To_AutoDownloadSettings()
		return t.Encode(layer)

	case Cmd_AutoDownloadSettingse04232f3:
		t := m.To_AutoDownloadSettings()
		return t.Encode(layer)

	default:
		log.Errorf("AutoDownloadSettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *BadMsgNotification) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_BadMsgNotification:
		t := m.To_BadMsgNotification()
		return t.Encode(layer)
	case Cmd_BadServerSalt:
		t := m.To_BadServerSalt()
		return t.Encode(layer)

	default:
		log.Errorf("BadMsgNotification Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *BankCardOpenUrl) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_BankCardOpenUrl:
		t := m.To_BankCardOpenUrl()
		return t.Encode(layer)

	default:
		log.Errorf("BankCardOpenUrl Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *BaseTheme) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_BaseThemeClassic:
		t := m.To_BaseThemeClassic()
		return t.Encode(layer)
	case Cmd_BaseThemeDay:
		t := m.To_BaseThemeDay()
		return t.Encode(layer)
	case Cmd_BaseThemeNight:
		t := m.To_BaseThemeNight()
		return t.Encode(layer)
	case Cmd_BaseThemeTinted:
		t := m.To_BaseThemeTinted()
		return t.Encode(layer)
	case Cmd_BaseThemeArctic:
		t := m.To_BaseThemeArctic()
		return t.Encode(layer)

	default:
		log.Errorf("BaseTheme Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *BindAuthKeyInner) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_BindAuthKeyInner:
		t := m.To_BindAuthKeyInner()
		return t.Encode(layer)

	default:
		log.Errorf("BindAuthKeyInner Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Bool) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_BoolFalse:
		t := m.To_BoolFalse()
		return t.Encode(layer)
	case Cmd_BoolTrue:
		t := m.To_BoolTrue()
		return t.Encode(layer)

	default:
		log.Errorf("Bool Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *BotCommand) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_BotCommand:
		t := m.To_BotCommand()
		return t.Encode(layer)

	default:
		log.Errorf("BotCommand Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *BotInfo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_BotInfo:
		t := m.To_BotInfo()
		return t.Encode(layer)

	default:
		log.Errorf("BotInfo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *BotInlineMessage) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_BotInlineMessageMediaAuto:
		t := m.To_BotInlineMessageMediaAuto()
		return t.Encode(layer)
	case Cmd_BotInlineMessageText:
		t := m.To_BotInlineMessageText()
		return t.Encode(layer)
	case Cmd_BotInlineMessageMediaGeo:
		t := m.To_BotInlineMessageMediaGeo()
		return t.Encode(layer)
	case Cmd_BotInlineMessageMediaVenue:
		t := m.To_BotInlineMessageMediaVenue()
		return t.Encode(layer)
	case Cmd_BotInlineMessageMediaContact:
		t := m.To_BotInlineMessageMediaContact()
		return t.Encode(layer)

	case Cmd_BotInlineMessageMediaAuto764cf810:
		t := m.To_BotInlineMessageMediaAuto()
		return t.Encode(layer)
	case Cmd_BotInlineMessageMediaGeob722de65:
		t := m.To_BotInlineMessageMediaGeo()
		return t.Encode(layer)
	case Cmd_BotInlineMessageMediaVenue8a86659c:
		t := m.To_BotInlineMessageMediaVenue()
		return t.Encode(layer)
	case Cmd_BotInlineMessageMediaContact18d1cdc2:
		t := m.To_BotInlineMessageMediaContact()
		return t.Encode(layer)
	case Cmd_BotInlineMessageMediaGeo51846fd:
		t := m.To_BotInlineMessageMediaGeo()
		return t.Encode(layer)

	default:
		log.Errorf("BotInlineMessage Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *BotInlineResult) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_BotInlineResult:
		t := m.To_BotInlineResult()
		return t.Encode(layer)
	case Cmd_BotInlineMediaResult:
		t := m.To_BotInlineMediaResult()
		return t.Encode(layer)

	case Cmd_BotInlineResult11965f3a:
		t := m.To_BotInlineResult()
		return t.Encode(layer)

	default:
		log.Errorf("BotInlineResult Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *CdnConfig) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_CdnConfig:
		t := m.To_CdnConfig()
		return t.Encode(layer)

	default:
		log.Errorf("CdnConfig Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *CdnFileHash) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_CdnFileHash:
		t := m.To_CdnFileHash()
		return t.Encode(layer)

	default:
		log.Errorf("CdnFileHash Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *CdnPublicKey) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_CdnPublicKey:
		t := m.To_CdnPublicKey()
		return t.Encode(layer)

	default:
		log.Errorf("CdnPublicKey Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChannelAdminLogEvent) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelAdminLogEvent:
		t := m.To_ChannelAdminLogEvent()
		return t.Encode(layer)

	default:
		log.Errorf("ChannelAdminLogEvent Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChannelAdminLogEventAction) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelAdminLogEventActionChangeTitle:
		t := m.To_ChannelAdminLogEventActionChangeTitle()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionChangeAbout:
		t := m.To_ChannelAdminLogEventActionChangeAbout()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionChangeUsername:
		t := m.To_ChannelAdminLogEventActionChangeUsername()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionChangePhoto:
		t := m.To_ChannelAdminLogEventActionChangePhoto()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionToggleInvites:
		t := m.To_ChannelAdminLogEventActionToggleInvites()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionToggleSignatures:
		t := m.To_ChannelAdminLogEventActionToggleSignatures()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionUpdatePinned:
		t := m.To_ChannelAdminLogEventActionUpdatePinned()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionEditMessage:
		t := m.To_ChannelAdminLogEventActionEditMessage()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionDeleteMessage:
		t := m.To_ChannelAdminLogEventActionDeleteMessage()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionParticipantJoin:
		t := m.To_ChannelAdminLogEventActionParticipantJoin()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionParticipantLeave:
		t := m.To_ChannelAdminLogEventActionParticipantLeave()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionParticipantInvite:
		t := m.To_ChannelAdminLogEventActionParticipantInvite()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionParticipantToggleBan:
		t := m.To_ChannelAdminLogEventActionParticipantToggleBan()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionParticipantToggleAdmin:
		t := m.To_ChannelAdminLogEventActionParticipantToggleAdmin()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionChangeStickerSet:
		t := m.To_ChannelAdminLogEventActionChangeStickerSet()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionTogglePreHistoryHidden:
		t := m.To_ChannelAdminLogEventActionTogglePreHistoryHidden()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionDefaultBannedRights:
		t := m.To_ChannelAdminLogEventActionDefaultBannedRights()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionStopPoll:
		t := m.To_ChannelAdminLogEventActionStopPoll()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionChangeLinkedChat:
		t := m.To_ChannelAdminLogEventActionChangeLinkedChat()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionChangeLocation:
		t := m.To_ChannelAdminLogEventActionChangeLocation()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionToggleSlowMode:
		t := m.To_ChannelAdminLogEventActionToggleSlowMode()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionStartGroupCall:
		t := m.To_ChannelAdminLogEventActionStartGroupCall()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionDiscardGroupCall:
		t := m.To_ChannelAdminLogEventActionDiscardGroupCall()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionParticipantMute:
		t := m.To_ChannelAdminLogEventActionParticipantMute()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionParticipantUnmute:
		t := m.To_ChannelAdminLogEventActionParticipantUnmute()
		return t.Encode(layer)
	case Cmd_ChannelAdminLogEventActionToggleGroupCallSetting:
		t := m.To_ChannelAdminLogEventActionToggleGroupCallSetting()
		return t.Encode(layer)

	case Cmd_ChannelAdminLogEventActionChangePhoto434bd2af:
		t := m.To_ChannelAdminLogEventActionChangePhoto()
		return t.Encode(layer)

	default:
		log.Errorf("ChannelAdminLogEventAction Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChannelAdminLogEventsFilter) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelAdminLogEventsFilter:
		t := m.To_ChannelAdminLogEventsFilter()
		return t.Encode(layer)

	default:
		log.Errorf("ChannelAdminLogEventsFilter Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChannelAdminRights) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelAdminRights:
		t := m.To_ChannelAdminRights()
		return t.Encode(layer)

	default:
		log.Errorf("ChannelAdminRights Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChannelBannedRights) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelBannedRights:
		t := m.To_ChannelBannedRights()
		return t.Encode(layer)

	default:
		log.Errorf("ChannelBannedRights Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChannelLocation) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelLocationEmpty:
		t := m.To_ChannelLocationEmpty()
		return t.Encode(layer)
	case Cmd_ChannelLocation:
		t := m.To_ChannelLocation()
		return t.Encode(layer)

	default:
		log.Errorf("ChannelLocation Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChannelMessagesFilter) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelMessagesFilterEmpty:
		t := m.To_ChannelMessagesFilterEmpty()
		return t.Encode(layer)
	case Cmd_ChannelMessagesFilter:
		t := m.To_ChannelMessagesFilter()
		return t.Encode(layer)
	case Cmd_ChannelMessagesFilterCollapsed:
		t := m.To_ChannelMessagesFilterCollapsed()
		return t.Encode(layer)

	default:
		log.Errorf("ChannelMessagesFilter Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChannelParticipant) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelParticipant:
		t := m.To_ChannelParticipant()
		return t.Encode(layer)
	case Cmd_ChannelParticipantSelf:
		t := m.To_ChannelParticipantSelf()
		return t.Encode(layer)
	case Cmd_ChannelParticipantCreator:
		t := m.To_ChannelParticipantCreator()
		return t.Encode(layer)
	case Cmd_ChannelParticipantAdmin:
		t := m.To_ChannelParticipantAdmin()
		return t.Encode(layer)
	case Cmd_ChannelParticipantBanned:
		t := m.To_ChannelParticipantBanned()
		return t.Encode(layer)
	case Cmd_ChannelParticipantModerator:
		t := m.To_ChannelParticipantModerator()
		return t.Encode(layer)
	case Cmd_ChannelParticipantEditor:
		t := m.To_ChannelParticipantEditor()
		return t.Encode(layer)
	case Cmd_ChannelParticipantKicked:
		t := m.To_ChannelParticipantKicked()
		return t.Encode(layer)
	case Cmd_ChannelParticipantLeft:
		t := m.To_ChannelParticipantLeft()
		return t.Encode(layer)

	case Cmd_ChannelParticipantAdmin5daa6e23:
		t := m.To_ChannelParticipantAdmin()
		return t.Encode(layer)
	case Cmd_ChannelParticipantBanned1c0facaf:
		t := m.To_ChannelParticipantBanned()
		return t.Encode(layer)
	case Cmd_ChannelParticipantCreator808d15a4:
		t := m.To_ChannelParticipantCreator()
		return t.Encode(layer)
	case Cmd_ChannelParticipantAdminccbebbaf:
		t := m.To_ChannelParticipantAdmin()
		return t.Encode(layer)
	case Cmd_ChannelParticipantCreator447dca4b:
		t := m.To_ChannelParticipantCreator()
		return t.Encode(layer)

	default:
		log.Errorf("ChannelParticipant Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChannelParticipantRole) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelRoleEmpty:
		t := m.To_ChannelRoleEmpty()
		return t.Encode(layer)
	case Cmd_ChannelRoleModerator:
		t := m.To_ChannelRoleModerator()
		return t.Encode(layer)
	case Cmd_ChannelRoleEditor:
		t := m.To_ChannelRoleEditor()
		return t.Encode(layer)

	default:
		log.Errorf("ChannelParticipantRole Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChannelParticipantsFilter) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelParticipantsRecent:
		t := m.To_ChannelParticipantsRecent()
		return t.Encode(layer)
	case Cmd_ChannelParticipantsAdmins:
		t := m.To_ChannelParticipantsAdmins()
		return t.Encode(layer)
	case Cmd_ChannelParticipantsKicked:
		t := m.To_ChannelParticipantsKicked()
		return t.Encode(layer)
	case Cmd_ChannelParticipantsBots:
		t := m.To_ChannelParticipantsBots()
		return t.Encode(layer)
	case Cmd_ChannelParticipantsBanned:
		t := m.To_ChannelParticipantsBanned()
		return t.Encode(layer)
	case Cmd_ChannelParticipantsSearch:
		t := m.To_ChannelParticipantsSearch()
		return t.Encode(layer)
	case Cmd_ChannelParticipantsContacts:
		t := m.To_ChannelParticipantsContacts()
		return t.Encode(layer)
	case Cmd_ChannelParticipantsMentions:
		t := m.To_ChannelParticipantsMentions()
		return t.Encode(layer)

	case Cmd_ChannelParticipantsKicked3c37bb7a:
		t := m.To_ChannelParticipantsKicked()
		return t.Encode(layer)

	default:
		log.Errorf("ChannelParticipantsFilter Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Channels_AdminLogResults) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelsAdminLogResults:
		t := m.To_ChannelsAdminLogResults()
		return t.Encode(layer)

	default:
		log.Errorf("Channels_AdminLogResults Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Channels_ChannelParticipant) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelsChannelParticipant:
		t := m.To_ChannelsChannelParticipant()
		return t.Encode(layer)

	default:
		log.Errorf("Channels_ChannelParticipant Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Channels_ChannelParticipants) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChannelsChannelParticipants:
		t := m.To_ChannelsChannelParticipants()
		return t.Encode(layer)
	case Cmd_ChannelsChannelParticipantsNotModified:
		t := m.To_ChannelsChannelParticipantsNotModified()
		return t.Encode(layer)

	default:
		log.Errorf("Channels_ChannelParticipants Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Chat) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChatEmpty:
		t := m.To_ChatEmpty()
		return t.Encode(layer)
	case Cmd_Chat:
		t := m.To_Chat()
		return t.Encode(layer)
	case Cmd_ChatForbidden:
		t := m.To_ChatForbidden()
		return t.Encode(layer)
	case Cmd_Channel:
		t := m.To_Channel()
		return t.Encode(layer)
	case Cmd_ChannelForbidden:
		t := m.To_ChannelForbidden()
		return t.Encode(layer)

	case Cmd_Channela14dca52:
		t := m.To_Channel()
		return t.Encode(layer)
	case Cmd_ChannelForbidden2d85832c:
		t := m.To_ChannelForbidden()
		return t.Encode(layer)
	case Cmd_Channelc88974ac:
		t := m.To_Channel()
		return t.Encode(layer)
	case Cmd_Channel4df30834:
		t := m.To_Channel()
		return t.Encode(layer)
	case Cmd_Chat3bda1bde:
		t := m.To_Chat()
		return t.Encode(layer)
	case Cmd_Channeld31a961e:
		t := m.To_Channel()
		return t.Encode(layer)

	default:
		log.Errorf("Chat Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChatAdminRights) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChatAdminRights:
		t := m.To_ChatAdminRights()
		return t.Encode(layer)

	default:
		log.Errorf("ChatAdminRights Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChatBannedRights) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChatBannedRights:
		t := m.To_ChatBannedRights()
		return t.Encode(layer)

	default:
		log.Errorf("ChatBannedRights Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChatFull) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChatFull:
		t := m.To_ChatFull()
		return t.Encode(layer)
	case Cmd_ChannelFull:
		t := m.To_ChannelFull()
		return t.Encode(layer)

	case Cmd_ChannelFull97bee562:
		t := m.To_ChannelFull()
		return t.Encode(layer)
	case Cmd_ChannelFull76af5481:
		t := m.To_ChannelFull()
		return t.Encode(layer)
	case Cmd_ChatFulledd2a791:
		t := m.To_ChatFull()
		return t.Encode(layer)
	case Cmd_ChannelFull1c87a71a:
		t := m.To_ChannelFull()
		return t.Encode(layer)
	case Cmd_ChatFull22a235da:
		t := m.To_ChatFull()
		return t.Encode(layer)
	case Cmd_ChatFull1b7c9db3:
		t := m.To_ChatFull()
		return t.Encode(layer)
	case Cmd_ChannelFull3648977:
		t := m.To_ChannelFull()
		return t.Encode(layer)
	case Cmd_ChannelFull9882e516:
		t := m.To_ChannelFull()
		return t.Encode(layer)
	case Cmd_ChannelFull10916653:
		t := m.To_ChannelFull()
		return t.Encode(layer)
	case Cmd_ChannelFull2d895c74:
		t := m.To_ChannelFull()
		return t.Encode(layer)
	case Cmd_ChannelFullf0e6672a:
		t := m.To_ChannelFull()
		return t.Encode(layer)
	case Cmd_ChatFulldc8c181:
		t := m.To_ChatFull()
		return t.Encode(layer)
	case Cmd_ChannelFullef3a6acd:
		t := m.To_ChannelFull()
		return t.Encode(layer)

	default:
		log.Errorf("ChatFull Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChatInvite) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChatInviteAlready:
		t := m.To_ChatInviteAlready()
		return t.Encode(layer)
	case Cmd_ChatInvite:
		t := m.To_ChatInvite()
		return t.Encode(layer)
	case Cmd_ChatInvitePeek:
		t := m.To_ChatInvitePeek()
		return t.Encode(layer)

	case Cmd_ChatInvite93e99b60:
		t := m.To_ChatInvite()
		return t.Encode(layer)
	case Cmd_ChatInvitedfc2f58e:
		t := m.To_ChatInvite()
		return t.Encode(layer)

	default:
		log.Errorf("ChatInvite Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChatOnlines) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChatOnlines:
		t := m.To_ChatOnlines()
		return t.Encode(layer)

	default:
		log.Errorf("ChatOnlines Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChatParticipant) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChatParticipant:
		t := m.To_ChatParticipant()
		return t.Encode(layer)
	case Cmd_ChatParticipantCreator:
		t := m.To_ChatParticipantCreator()
		return t.Encode(layer)
	case Cmd_ChatParticipantAdmin:
		t := m.To_ChatParticipantAdmin()
		return t.Encode(layer)

	default:
		log.Errorf("ChatParticipant Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChatParticipants) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChatParticipantsForbidden:
		t := m.To_ChatParticipantsForbidden()
		return t.Encode(layer)
	case Cmd_ChatParticipants:
		t := m.To_ChatParticipants()
		return t.Encode(layer)

	default:
		log.Errorf("ChatParticipants Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ChatPhoto) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChatPhotoEmpty:
		t := m.To_ChatPhotoEmpty()
		return t.Encode(layer)
	case Cmd_ChatPhoto:
		t := m.To_ChatPhoto()
		return t.Encode(layer)

	case Cmd_ChatPhoto475cdbd5:
		t := m.To_ChatPhoto()
		return t.Encode(layer)
	case Cmd_ChatPhotod20b9f3c:
		t := m.To_ChatPhoto()
		return t.Encode(layer)

	default:
		log.Errorf("ChatPhoto Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Client_DH_Inner_Data) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Client_DHInnerData:
		t := m.To_Client_DHInnerData()
		return t.Encode(layer)

	default:
		log.Errorf("Client_DH_Inner_Data Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *CodeSettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_CodeSettings:
		t := m.To_CodeSettings()
		return t.Encode(layer)

	case Cmd_CodeSettingsdebebe83:
		t := m.To_CodeSettings()
		return t.Encode(layer)

	default:
		log.Errorf("CodeSettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Config) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Config:
		t := m.To_Config()
		return t.Encode(layer)

	case Cmd_Config317ceef4:
		t := m.To_Config()
		return t.Encode(layer)
	case Cmd_Config3213dbba:
		t := m.To_Config()
		return t.Encode(layer)
	case Cmd_Confige6ca25f6:
		t := m.To_Config()
		return t.Encode(layer)
	case Cmd_Config330b4067:
		t := m.To_Config()
		return t.Encode(layer)

	default:
		log.Errorf("Config Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Contact) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Contact:
		t := m.To_Contact()
		return t.Encode(layer)

	default:
		log.Errorf("Contact Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ContactBlocked) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ContactBlocked:
		t := m.To_ContactBlocked()
		return t.Encode(layer)

	default:
		log.Errorf("ContactBlocked Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ContactLink) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ContactLinkUnknown:
		t := m.To_ContactLinkUnknown()
		return t.Encode(layer)
	case Cmd_ContactLinkNone:
		t := m.To_ContactLinkNone()
		return t.Encode(layer)
	case Cmd_ContactLinkHasPhone:
		t := m.To_ContactLinkHasPhone()
		return t.Encode(layer)
	case Cmd_ContactLinkContact:
		t := m.To_ContactLinkContact()
		return t.Encode(layer)

	default:
		log.Errorf("ContactLink Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ContactStatus) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ContactStatus:
		t := m.To_ContactStatus()
		return t.Encode(layer)

	default:
		log.Errorf("ContactStatus Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Contacts_Blocked) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ContactsBlocked:
		t := m.To_ContactsBlocked()
		return t.Encode(layer)
	case Cmd_ContactsBlockedSlice:
		t := m.To_ContactsBlockedSlice()
		return t.Encode(layer)

	case Cmd_ContactsBlockedade1591:
		t := m.To_ContactsBlocked()
		return t.Encode(layer)
	case Cmd_ContactsBlockedSlicee1664194:
		t := m.To_ContactsBlockedSlice()
		return t.Encode(layer)

	default:
		log.Errorf("Contacts_Blocked Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Contacts_Contacts) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ContactsContactsNotModified:
		t := m.To_ContactsContactsNotModified()
		return t.Encode(layer)
	case Cmd_ContactsContacts:
		t := m.To_ContactsContacts()
		return t.Encode(layer)

	case Cmd_ContactsContacts6f8b8cb2:
		t := m.To_ContactsContacts()
		return t.Encode(layer)

	default:
		log.Errorf("Contacts_Contacts Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Contacts_Found) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ContactsFound:
		t := m.To_ContactsFound()
		return t.Encode(layer)

	case Cmd_ContactsFoundb3134d9d:
		t := m.To_ContactsFound()
		return t.Encode(layer)

	default:
		log.Errorf("Contacts_Found Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Contacts_ImportedContacts) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ContactsImportedContacts:
		t := m.To_ContactsImportedContacts()
		return t.Encode(layer)

	case Cmd_ContactsImportedContactsad524315:
		t := m.To_ContactsImportedContacts()
		return t.Encode(layer)

	default:
		log.Errorf("Contacts_ImportedContacts Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Contacts_Link) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ContactsLink:
		t := m.To_ContactsLink()
		return t.Encode(layer)

	default:
		log.Errorf("Contacts_Link Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Contacts_ResolvedPeer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ContactsResolvedPeer:
		t := m.To_ContactsResolvedPeer()
		return t.Encode(layer)

	default:
		log.Errorf("Contacts_ResolvedPeer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Contacts_TopPeers) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ContactsTopPeersNotModified:
		t := m.To_ContactsTopPeersNotModified()
		return t.Encode(layer)
	case Cmd_ContactsTopPeers:
		t := m.To_ContactsTopPeers()
		return t.Encode(layer)
	case Cmd_ContactsTopPeersDisabled:
		t := m.To_ContactsTopPeersDisabled()
		return t.Encode(layer)

	default:
		log.Errorf("Contacts_TopPeers Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *DataJSON) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DataJSON:
		t := m.To_DataJSON()
		return t.Encode(layer)

	default:
		log.Errorf("DataJSON Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *DcOption) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DcOption:
		t := m.To_DcOption()
		return t.Encode(layer)

	case Cmd_DcOption18b7a10d:
		t := m.To_DcOption()
		return t.Encode(layer)

	default:
		log.Errorf("DcOption Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *DestroyAuthKeyRes) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DestroyAuthKeyOk:
		t := m.To_DestroyAuthKeyOk()
		return t.Encode(layer)
	case Cmd_DestroyAuthKeyNone:
		t := m.To_DestroyAuthKeyNone()
		return t.Encode(layer)
	case Cmd_DestroyAuthKeyFail:
		t := m.To_DestroyAuthKeyFail()
		return t.Encode(layer)

	default:
		log.Errorf("DestroyAuthKeyRes Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *DestroySessionRes) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DestroySessionOk:
		t := m.To_DestroySessionOk()
		return t.Encode(layer)
	case Cmd_DestroySessionNone:
		t := m.To_DestroySessionNone()
		return t.Encode(layer)

	default:
		log.Errorf("DestroySessionRes Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Dialog) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Dialog:
		t := m.To_Dialog()
		return t.Encode(layer)
	case Cmd_DialogChannel:
		t := m.To_DialogChannel()
		return t.Encode(layer)
	case Cmd_DialogFolder:
		t := m.To_DialogFolder()
		return t.Encode(layer)

	case Cmd_Dialogc1dd804a:
		t := m.To_Dialog()
		return t.Encode(layer)
	case Cmd_Dialog2c171f72:
		t := m.To_Dialog()
		return t.Encode(layer)

	default:
		log.Errorf("Dialog Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *DialogFilter) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DialogFilter:
		t := m.To_DialogFilter()
		return t.Encode(layer)

	default:
		log.Errorf("DialogFilter Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *DialogFilterSuggested) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DialogFilterSuggested:
		t := m.To_DialogFilterSuggested()
		return t.Encode(layer)

	default:
		log.Errorf("DialogFilterSuggested Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *DialogPeer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DialogPeer:
		t := m.To_DialogPeer()
		return t.Encode(layer)
	case Cmd_DialogPeerFolder:
		t := m.To_DialogPeerFolder()
		return t.Encode(layer)

	default:
		log.Errorf("DialogPeer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *DisabledFeature) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DisabledFeature:
		t := m.To_DisabledFeature()
		return t.Encode(layer)

	default:
		log.Errorf("DisabledFeature Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Document) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DocumentEmpty:
		t := m.To_DocumentEmpty()
		return t.Encode(layer)
	case Cmd_Document:
		t := m.To_Document()
		return t.Encode(layer)

	case Cmd_Documentf9a39f4f:
		t := m.To_Document()
		return t.Encode(layer)
	case Cmd_Document59534e4c:
		t := m.To_Document()
		return t.Encode(layer)
	case Cmd_Document9ba29cc1:
		t := m.To_Document()
		return t.Encode(layer)
	case Cmd_Document1e87342b:
		t := m.To_Document()
		return t.Encode(layer)

	default:
		log.Errorf("Document Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *DocumentAttribute) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DocumentAttributeImageSize:
		t := m.To_DocumentAttributeImageSize()
		return t.Encode(layer)
	case Cmd_DocumentAttributeAnimated:
		t := m.To_DocumentAttributeAnimated()
		return t.Encode(layer)
	case Cmd_DocumentAttributeSticker:
		t := m.To_DocumentAttributeSticker()
		return t.Encode(layer)
	case Cmd_DocumentAttributeVideo:
		t := m.To_DocumentAttributeVideo()
		return t.Encode(layer)
	case Cmd_DocumentAttributeAudio:
		t := m.To_DocumentAttributeAudio()
		return t.Encode(layer)
	case Cmd_DocumentAttributeFilename:
		t := m.To_DocumentAttributeFilename()
		return t.Encode(layer)
	case Cmd_DocumentAttributeHasStickers:
		t := m.To_DocumentAttributeHasStickers()
		return t.Encode(layer)

	case Cmd_DocumentAttributeSticker3a556302:
		t := m.To_DocumentAttributeSticker()
		return t.Encode(layer)
	case Cmd_DocumentAttributeVideo5910cccb:
		t := m.To_DocumentAttributeVideo()
		return t.Encode(layer)

	default:
		log.Errorf("DocumentAttribute Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *DraftMessage) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DraftMessageEmpty:
		t := m.To_DraftMessageEmpty()
		return t.Encode(layer)
	case Cmd_DraftMessage:
		t := m.To_DraftMessage()
		return t.Encode(layer)

	case Cmd_DraftMessageEmpty1b0c841a:
		t := m.To_DraftMessageEmpty()
		return t.Encode(layer)

	default:
		log.Errorf("DraftMessage Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *EmojiKeyword) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_EmojiKeyword:
		t := m.To_EmojiKeyword()
		return t.Encode(layer)
	case Cmd_EmojiKeywordDeleted:
		t := m.To_EmojiKeywordDeleted()
		return t.Encode(layer)

	default:
		log.Errorf("EmojiKeyword Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *EmojiKeywordsDifference) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_EmojiKeywordsDifference:
		t := m.To_EmojiKeywordsDifference()
		return t.Encode(layer)

	default:
		log.Errorf("EmojiKeywordsDifference Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *EmojiLanguage) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_EmojiLanguage:
		t := m.To_EmojiLanguage()
		return t.Encode(layer)

	default:
		log.Errorf("EmojiLanguage Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *EmojiURL) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_EmojiURL:
		t := m.To_EmojiURL()
		return t.Encode(layer)

	default:
		log.Errorf("EmojiURL Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *EncryptedChat) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_EncryptedChatEmpty:
		t := m.To_EncryptedChatEmpty()
		return t.Encode(layer)
	case Cmd_EncryptedChatWaiting:
		t := m.To_EncryptedChatWaiting()
		return t.Encode(layer)
	case Cmd_EncryptedChatRequested:
		t := m.To_EncryptedChatRequested()
		return t.Encode(layer)
	case Cmd_EncryptedChat:
		t := m.To_EncryptedChat()
		return t.Encode(layer)
	case Cmd_EncryptedChatDiscarded:
		t := m.To_EncryptedChatDiscarded()
		return t.Encode(layer)

	case Cmd_EncryptedChatRequested62718a82:
		t := m.To_EncryptedChatRequested()
		return t.Encode(layer)

	default:
		log.Errorf("EncryptedChat Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *EncryptedFile) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_EncryptedFileEmpty:
		t := m.To_EncryptedFileEmpty()
		return t.Encode(layer)
	case Cmd_EncryptedFile:
		t := m.To_EncryptedFile()
		return t.Encode(layer)

	default:
		log.Errorf("EncryptedFile Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *EncryptedMessage) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_EncryptedMessage:
		t := m.To_EncryptedMessage()
		return t.Encode(layer)
	case Cmd_EncryptedMessageService:
		t := m.To_EncryptedMessageService()
		return t.Encode(layer)

	default:
		log.Errorf("EncryptedMessage Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Error) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Error:
		t := m.To_Error()
		return t.Encode(layer)

	default:
		log.Errorf("Error Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ExportedChatInvite) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ChatInviteEmpty:
		t := m.To_ChatInviteEmpty()
		return t.Encode(layer)
	case Cmd_ChatInviteExported:
		t := m.To_ChatInviteExported()
		return t.Encode(layer)

	default:
		log.Errorf("ExportedChatInvite Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ExportedMessageLink) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ExportedMessageLink:
		t := m.To_ExportedMessageLink()
		return t.Encode(layer)

	case Cmd_ExportedMessageLink5dab1af4:
		t := m.To_ExportedMessageLink()
		return t.Encode(layer)

	default:
		log.Errorf("ExportedMessageLink Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *FileHash) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_FileHash:
		t := m.To_FileHash()
		return t.Encode(layer)

	default:
		log.Errorf("FileHash Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *FileLocation) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_FileLocationUnavailable:
		t := m.To_FileLocationUnavailable()
		return t.Encode(layer)
	case Cmd_FileLocation:
		t := m.To_FileLocation()
		return t.Encode(layer)
	case Cmd_FileLocationToBeDeprecated:
		t := m.To_FileLocationToBeDeprecated()
		return t.Encode(layer)

	case Cmd_FileLocation91d11eb:
		t := m.To_FileLocation()
		return t.Encode(layer)

	default:
		log.Errorf("FileLocation Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Folder) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Folder:
		t := m.To_Folder()
		return t.Encode(layer)

	default:
		log.Errorf("Folder Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *FolderPeer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_FolderPeer:
		t := m.To_FolderPeer()
		return t.Encode(layer)

	default:
		log.Errorf("FolderPeer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *FoundGif) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_FoundGif:
		t := m.To_FoundGif()
		return t.Encode(layer)
	case Cmd_FoundGifCached:
		t := m.To_FoundGifCached()
		return t.Encode(layer)

	default:
		log.Errorf("FoundGif Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *FutureSalt) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_FutureSalt:
		t := m.To_FutureSalt()
		return t.Encode(layer)

	default:
		log.Errorf("FutureSalt Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *FutureSalts) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_FutureSalts:
		t := m.To_FutureSalts()
		return t.Encode(layer)

	default:
		log.Errorf("FutureSalts Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Game) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Game:
		t := m.To_Game()
		return t.Encode(layer)

	default:
		log.Errorf("Game Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *GeoPoint) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_GeoPointEmpty:
		t := m.To_GeoPointEmpty()
		return t.Encode(layer)
	case Cmd_GeoPoint:
		t := m.To_GeoPoint()
		return t.Encode(layer)

	case Cmd_GeoPoint296f104:
		t := m.To_GeoPoint()
		return t.Encode(layer)
	case Cmd_GeoPointb2a2f663:
		t := m.To_GeoPoint()
		return t.Encode(layer)

	default:
		log.Errorf("GeoPoint Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *GlobalPrivacySettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_GlobalPrivacySettings:
		t := m.To_GlobalPrivacySettings()
		return t.Encode(layer)

	default:
		log.Errorf("GlobalPrivacySettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *GroupCall) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_GroupCallDiscarded:
		t := m.To_GroupCallDiscarded()
		return t.Encode(layer)
	case Cmd_GroupCall:
		t := m.To_GroupCall()
		return t.Encode(layer)

	default:
		log.Errorf("GroupCall Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *GroupCallParticipant) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_GroupCallParticipant:
		t := m.To_GroupCallParticipant()
		return t.Encode(layer)

	default:
		log.Errorf("GroupCallParticipant Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_AppChangelog) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpAppChangelogEmpty:
		t := m.To_HelpAppChangelogEmpty()
		return t.Encode(layer)
	case Cmd_HelpAppChangelog:
		t := m.To_HelpAppChangelog()
		return t.Encode(layer)

	default:
		log.Errorf("Help_AppChangelog Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_AppUpdate) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpAppUpdate:
		t := m.To_HelpAppUpdate()
		return t.Encode(layer)
	case Cmd_HelpNoAppUpdate:
		t := m.To_HelpNoAppUpdate()
		return t.Encode(layer)

	case Cmd_HelpAppUpdate1da7158f:
		t := m.To_HelpAppUpdate()
		return t.Encode(layer)

	default:
		log.Errorf("Help_AppUpdate Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_ConfigSimple) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpConfigSimple:
		t := m.To_HelpConfigSimple()
		return t.Encode(layer)

	case Cmd_HelpConfigSimple5a592a6c:
		t := m.To_HelpConfigSimple()
		return t.Encode(layer)

	default:
		log.Errorf("Help_ConfigSimple Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_CountriesList) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpCountriesListNotModified:
		t := m.To_HelpCountriesListNotModified()
		return t.Encode(layer)
	case Cmd_HelpCountriesList:
		t := m.To_HelpCountriesList()
		return t.Encode(layer)

	default:
		log.Errorf("Help_CountriesList Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_Country) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpCountry:
		t := m.To_HelpCountry()
		return t.Encode(layer)

	default:
		log.Errorf("Help_Country Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_CountryCode) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpCountryCode:
		t := m.To_HelpCountryCode()
		return t.Encode(layer)

	default:
		log.Errorf("Help_CountryCode Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_DeepLinkInfo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpDeepLinkInfoEmpty:
		t := m.To_HelpDeepLinkInfoEmpty()
		return t.Encode(layer)
	case Cmd_HelpDeepLinkInfo:
		t := m.To_HelpDeepLinkInfo()
		return t.Encode(layer)

	default:
		log.Errorf("Help_DeepLinkInfo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_InviteText) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpInviteText:
		t := m.To_HelpInviteText()
		return t.Encode(layer)

	default:
		log.Errorf("Help_InviteText Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_PassportConfig) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpPassportConfigNotModified:
		t := m.To_HelpPassportConfigNotModified()
		return t.Encode(layer)
	case Cmd_HelpPassportConfig:
		t := m.To_HelpPassportConfig()
		return t.Encode(layer)

	default:
		log.Errorf("Help_PassportConfig Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_PromoData) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpPromoDataEmpty:
		t := m.To_HelpPromoDataEmpty()
		return t.Encode(layer)
	case Cmd_HelpPromoData:
		t := m.To_HelpPromoData()
		return t.Encode(layer)

	case Cmd_HelpPromoData8c39793f:
		t := m.To_HelpPromoData()
		return t.Encode(layer)

	default:
		log.Errorf("Help_PromoData Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_ProxyData) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpProxyDataEmpty:
		t := m.To_HelpProxyDataEmpty()
		return t.Encode(layer)
	case Cmd_HelpProxyDataPromo:
		t := m.To_HelpProxyDataPromo()
		return t.Encode(layer)

	default:
		log.Errorf("Help_ProxyData Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_RecentMeUrls) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpRecentMeUrls:
		t := m.To_HelpRecentMeUrls()
		return t.Encode(layer)

	default:
		log.Errorf("Help_RecentMeUrls Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_Support) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpSupport:
		t := m.To_HelpSupport()
		return t.Encode(layer)

	default:
		log.Errorf("Help_Support Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_SupportName) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpSupportName:
		t := m.To_HelpSupportName()
		return t.Encode(layer)

	default:
		log.Errorf("Help_SupportName Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_TermsOfService) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpTermsOfService:
		t := m.To_HelpTermsOfService()
		return t.Encode(layer)

	case Cmd_HelpTermsOfService780a0310:
		t := m.To_HelpTermsOfService()
		return t.Encode(layer)

	default:
		log.Errorf("Help_TermsOfService Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_TermsOfServiceUpdate) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpTermsOfServiceUpdateEmpty:
		t := m.To_HelpTermsOfServiceUpdateEmpty()
		return t.Encode(layer)
	case Cmd_HelpTermsOfServiceUpdate:
		t := m.To_HelpTermsOfServiceUpdate()
		return t.Encode(layer)

	default:
		log.Errorf("Help_TermsOfServiceUpdate Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Help_UserInfo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HelpUserInfoEmpty:
		t := m.To_HelpUserInfoEmpty()
		return t.Encode(layer)
	case Cmd_HelpUserInfo:
		t := m.To_HelpUserInfo()
		return t.Encode(layer)

	default:
		log.Errorf("Help_UserInfo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *HighScore) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HighScore:
		t := m.To_HighScore()
		return t.Encode(layer)

	default:
		log.Errorf("HighScore Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *HttpWait) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_HttpWait:
		t := m.To_HttpWait()
		return t.Encode(layer)

	default:
		log.Errorf("HttpWait Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ImportedContact) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ImportedContact:
		t := m.To_ImportedContact()
		return t.Encode(layer)

	default:
		log.Errorf("ImportedContact Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InlineBotSwitchPM) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InlineBotSwitchPM:
		t := m.To_InlineBotSwitchPM()
		return t.Encode(layer)

	default:
		log.Errorf("InlineBotSwitchPM Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InlineQueryPeerType) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InlineQueryPeerTypeSameBotPM:
		t := m.To_InlineQueryPeerTypeSameBotPM()
		return t.Encode(layer)
	case Cmd_InlineQueryPeerTypePM:
		t := m.To_InlineQueryPeerTypePM()
		return t.Encode(layer)
	case Cmd_InlineQueryPeerTypeChat:
		t := m.To_InlineQueryPeerTypeChat()
		return t.Encode(layer)
	case Cmd_InlineQueryPeerTypeMegagroup:
		t := m.To_InlineQueryPeerTypeMegagroup()
		return t.Encode(layer)
	case Cmd_InlineQueryPeerTypeBroadcast:
		t := m.To_InlineQueryPeerTypeBroadcast()
		return t.Encode(layer)

	default:
		log.Errorf("InlineQueryPeerType Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputAppEvent) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputAppEvent:
		t := m.To_InputAppEvent()
		return t.Encode(layer)

	case Cmd_InputAppEvent1d1b1245:
		t := m.To_InputAppEvent()
		return t.Encode(layer)

	default:
		log.Errorf("InputAppEvent Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputBotInlineMessage) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputBotInlineMessageMediaAuto:
		t := m.To_InputBotInlineMessageMediaAuto()
		return t.Encode(layer)
	case Cmd_InputBotInlineMessageText:
		t := m.To_InputBotInlineMessageText()
		return t.Encode(layer)
	case Cmd_InputBotInlineMessageMediaGeo:
		t := m.To_InputBotInlineMessageMediaGeo()
		return t.Encode(layer)
	case Cmd_InputBotInlineMessageMediaVenue:
		t := m.To_InputBotInlineMessageMediaVenue()
		return t.Encode(layer)
	case Cmd_InputBotInlineMessageMediaContact:
		t := m.To_InputBotInlineMessageMediaContact()
		return t.Encode(layer)
	case Cmd_InputBotInlineMessageGame:
		t := m.To_InputBotInlineMessageGame()
		return t.Encode(layer)

	case Cmd_InputBotInlineMessageMediaAuto3380c786:
		t := m.To_InputBotInlineMessageMediaAuto()
		return t.Encode(layer)
	case Cmd_InputBotInlineMessageMediaGeoc1b15d65:
		t := m.To_InputBotInlineMessageMediaGeo()
		return t.Encode(layer)
	case Cmd_InputBotInlineMessageMediaVenue417bbf11:
		t := m.To_InputBotInlineMessageMediaVenue()
		return t.Encode(layer)
	case Cmd_InputBotInlineMessageMediaContacta6edbffd:
		t := m.To_InputBotInlineMessageMediaContact()
		return t.Encode(layer)
	case Cmd_InputBotInlineMessageMediaGeo96929a85:
		t := m.To_InputBotInlineMessageMediaGeo()
		return t.Encode(layer)

	default:
		log.Errorf("InputBotInlineMessage Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputBotInlineMessageID) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputBotInlineMessageID:
		t := m.To_InputBotInlineMessageID()
		return t.Encode(layer)

	default:
		log.Errorf("InputBotInlineMessageID Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputBotInlineResult) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputBotInlineResult:
		t := m.To_InputBotInlineResult()
		return t.Encode(layer)
	case Cmd_InputBotInlineResultPhoto:
		t := m.To_InputBotInlineResultPhoto()
		return t.Encode(layer)
	case Cmd_InputBotInlineResultDocument:
		t := m.To_InputBotInlineResultDocument()
		return t.Encode(layer)
	case Cmd_InputBotInlineResultGame:
		t := m.To_InputBotInlineResultGame()
		return t.Encode(layer)

	case Cmd_InputBotInlineResult88bf9319:
		t := m.To_InputBotInlineResult()
		return t.Encode(layer)

	default:
		log.Errorf("InputBotInlineResult Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputChannel) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputChannelEmpty:
		t := m.To_InputChannelEmpty()
		return t.Encode(layer)
	case Cmd_InputChannel:
		t := m.To_InputChannel()
		return t.Encode(layer)
	case Cmd_InputChannelFromMessage:
		t := m.To_InputChannelFromMessage()
		return t.Encode(layer)

	default:
		log.Errorf("InputChannel Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputChatPhoto) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputChatPhotoEmpty:
		t := m.To_InputChatPhotoEmpty()
		return t.Encode(layer)
	case Cmd_InputChatUploadedPhoto:
		t := m.To_InputChatUploadedPhoto()
		return t.Encode(layer)
	case Cmd_InputChatPhoto:
		t := m.To_InputChatPhoto()
		return t.Encode(layer)

	case Cmd_InputChatUploadedPhoto94254732:
		t := m.To_InputChatUploadedPhoto()
		return t.Encode(layer)
	case Cmd_InputChatPhotob2e1bf08:
		t := m.To_InputChatPhoto()
		return t.Encode(layer)
	case Cmd_InputChatUploadedPhotoc642724e:
		t := m.To_InputChatUploadedPhoto()
		return t.Encode(layer)

	default:
		log.Errorf("InputChatPhoto Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputCheckPasswordSRP) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputCheckPasswordEmpty:
		t := m.To_InputCheckPasswordEmpty()
		return t.Encode(layer)
	case Cmd_InputCheckPasswordSRP:
		t := m.To_InputCheckPasswordSRP()
		return t.Encode(layer)

	default:
		log.Errorf("InputCheckPasswordSRP Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputClientProxy) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputClientProxy:
		t := m.To_InputClientProxy()
		return t.Encode(layer)

	default:
		log.Errorf("InputClientProxy Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputContact) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputPhoneContact:
		t := m.To_InputPhoneContact()
		return t.Encode(layer)

	default:
		log.Errorf("InputContact Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputDialogPeer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputDialogPeer:
		t := m.To_InputDialogPeer()
		return t.Encode(layer)
	case Cmd_InputDialogPeerFolder:
		t := m.To_InputDialogPeerFolder()
		return t.Encode(layer)

	default:
		log.Errorf("InputDialogPeer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputDocument) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputDocumentEmpty:
		t := m.To_InputDocumentEmpty()
		return t.Encode(layer)
	case Cmd_InputDocument:
		t := m.To_InputDocument()
		return t.Encode(layer)

	case Cmd_InputDocument1abfb575:
		t := m.To_InputDocument()
		return t.Encode(layer)

	default:
		log.Errorf("InputDocument Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputEncryptedChat) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputEncryptedChat:
		t := m.To_InputEncryptedChat()
		return t.Encode(layer)

	default:
		log.Errorf("InputEncryptedChat Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputEncryptedFile) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputEncryptedFileEmpty:
		t := m.To_InputEncryptedFileEmpty()
		return t.Encode(layer)
	case Cmd_InputEncryptedFileUploaded:
		t := m.To_InputEncryptedFileUploaded()
		return t.Encode(layer)
	case Cmd_InputEncryptedFile:
		t := m.To_InputEncryptedFile()
		return t.Encode(layer)
	case Cmd_InputEncryptedFileBigUploaded:
		t := m.To_InputEncryptedFileBigUploaded()
		return t.Encode(layer)

	default:
		log.Errorf("InputEncryptedFile Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputFile) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputFile:
		t := m.To_InputFile()
		return t.Encode(layer)
	case Cmd_InputFileBig:
		t := m.To_InputFileBig()
		return t.Encode(layer)

	default:
		log.Errorf("InputFile Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputFileLocation) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputFileLocation:
		t := m.To_InputFileLocation()
		return t.Encode(layer)
	case Cmd_InputEncryptedFileLocation:
		t := m.To_InputEncryptedFileLocation()
		return t.Encode(layer)
	case Cmd_InputDocumentFileLocation:
		t := m.To_InputDocumentFileLocation()
		return t.Encode(layer)
	case Cmd_InputSecureFileLocation:
		t := m.To_InputSecureFileLocation()
		return t.Encode(layer)
	case Cmd_InputTakeoutFileLocation:
		t := m.To_InputTakeoutFileLocation()
		return t.Encode(layer)
	case Cmd_InputPhotoFileLocation:
		t := m.To_InputPhotoFileLocation()
		return t.Encode(layer)
	case Cmd_InputPeerPhotoFileLocation:
		t := m.To_InputPeerPhotoFileLocation()
		return t.Encode(layer)
	case Cmd_InputStickerSetThumb:
		t := m.To_InputStickerSetThumb()
		return t.Encode(layer)
	case Cmd_InputPhotoLegacyFileLocation:
		t := m.To_InputPhotoLegacyFileLocation()
		return t.Encode(layer)

	case Cmd_InputDocumentFileLocation4e45abe9:
		t := m.To_InputDocumentFileLocation()
		return t.Encode(layer)
	case Cmd_InputFileLocationdfdaabe1:
		t := m.To_InputFileLocation()
		return t.Encode(layer)
	case Cmd_InputDocumentFileLocation196683d9:
		t := m.To_InputDocumentFileLocation()
		return t.Encode(layer)
	case Cmd_InputDocumentFileLocationbad07584:
		t := m.To_InputDocumentFileLocation()
		return t.Encode(layer)

	default:
		log.Errorf("InputFileLocation Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputFolderPeer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputFolderPeer:
		t := m.To_InputFolderPeer()
		return t.Encode(layer)

	default:
		log.Errorf("InputFolderPeer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputGame) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputGameID:
		t := m.To_InputGameID()
		return t.Encode(layer)
	case Cmd_InputGameShortName:
		t := m.To_InputGameShortName()
		return t.Encode(layer)

	default:
		log.Errorf("InputGame Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputGeoPoint) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputGeoPointEmpty:
		t := m.To_InputGeoPointEmpty()
		return t.Encode(layer)
	case Cmd_InputGeoPoint:
		t := m.To_InputGeoPoint()
		return t.Encode(layer)

	case Cmd_InputGeoPoint48222faf:
		t := m.To_InputGeoPoint()
		return t.Encode(layer)

	default:
		log.Errorf("InputGeoPoint Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputGroupCall) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputGroupCall:
		t := m.To_InputGroupCall()
		return t.Encode(layer)

	default:
		log.Errorf("InputGroupCall Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputMedia) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputMediaEmpty:
		t := m.To_InputMediaEmpty()
		return t.Encode(layer)
	case Cmd_InputMediaUploadedPhoto:
		t := m.To_InputMediaUploadedPhoto()
		return t.Encode(layer)
	case Cmd_InputMediaPhoto:
		t := m.To_InputMediaPhoto()
		return t.Encode(layer)
	case Cmd_InputMediaGeoPoint:
		t := m.To_InputMediaGeoPoint()
		return t.Encode(layer)
	case Cmd_InputMediaContact:
		t := m.To_InputMediaContact()
		return t.Encode(layer)
	case Cmd_InputMediaUploadedDocument:
		t := m.To_InputMediaUploadedDocument()
		return t.Encode(layer)
	case Cmd_InputMediaDocument:
		t := m.To_InputMediaDocument()
		return t.Encode(layer)
	case Cmd_InputMediaVenue:
		t := m.To_InputMediaVenue()
		return t.Encode(layer)
	case Cmd_InputMediaGifExternal:
		t := m.To_InputMediaGifExternal()
		return t.Encode(layer)
	case Cmd_InputMediaPhotoExternal:
		t := m.To_InputMediaPhotoExternal()
		return t.Encode(layer)
	case Cmd_InputMediaDocumentExternal:
		t := m.To_InputMediaDocumentExternal()
		return t.Encode(layer)
	case Cmd_InputMediaGame:
		t := m.To_InputMediaGame()
		return t.Encode(layer)
	case Cmd_InputMediaInvoice:
		t := m.To_InputMediaInvoice()
		return t.Encode(layer)
	case Cmd_InputMediaUploadedThumbDocument:
		t := m.To_InputMediaUploadedThumbDocument()
		return t.Encode(layer)
	case Cmd_InputMediaGeoLive:
		t := m.To_InputMediaGeoLive()
		return t.Encode(layer)
	case Cmd_InputMediaPoll:
		t := m.To_InputMediaPoll()
		return t.Encode(layer)
	case Cmd_InputMediaDice:
		t := m.To_InputMediaDice()
		return t.Encode(layer)

	case Cmd_InputMediaPhotoe9bfb4f3:
		t := m.To_InputMediaPhoto()
		return t.Encode(layer)
	case Cmd_InputMediaUploadedPhotof7aff1c0:
		t := m.To_InputMediaUploadedPhoto()
		return t.Encode(layer)
	case Cmd_InputMediaUploadedDocument1d89306d:
		t := m.To_InputMediaUploadedDocument()
		return t.Encode(layer)
	case Cmd_InputMediaDocument1a77f29c:
		t := m.To_InputMediaDocument()
		return t.Encode(layer)
	case Cmd_InputMediaUploadedPhoto1e287d04:
		t := m.To_InputMediaUploadedPhoto()
		return t.Encode(layer)
	case Cmd_InputMediaPhotob3ba0635:
		t := m.To_InputMediaPhoto()
		return t.Encode(layer)
	case Cmd_InputMediaContactf8ab7dfb:
		t := m.To_InputMediaContact()
		return t.Encode(layer)
	case Cmd_InputMediaDocument23ab23d2:
		t := m.To_InputMediaDocument()
		return t.Encode(layer)
	case Cmd_InputMediaVenuec13d1c11:
		t := m.To_InputMediaVenue()
		return t.Encode(layer)
	case Cmd_InputMediaPhotoExternale5bbfe1a:
		t := m.To_InputMediaPhotoExternal()
		return t.Encode(layer)
	case Cmd_InputMediaDocumentExternalfb52dc99:
		t := m.To_InputMediaDocumentExternal()
		return t.Encode(layer)
	case Cmd_InputMediaInvoicef4e096c3:
		t := m.To_InputMediaInvoice()
		return t.Encode(layer)
	case Cmd_InputMediaGeoLivece4e82fd:
		t := m.To_InputMediaGeoLive()
		return t.Encode(layer)
	case Cmd_InputMediaPollabe9ca25:
		t := m.To_InputMediaPoll()
		return t.Encode(layer)
	case Cmd_InputMediaPollf94e5f1:
		t := m.To_InputMediaPoll()
		return t.Encode(layer)
	case Cmd_InputMediaDicee66fbf7b:
		t := m.To_InputMediaDice()
		return t.Encode(layer)
	case Cmd_InputMediaUploadedDocument5b38c6c1:
		t := m.To_InputMediaUploadedDocument()
		return t.Encode(layer)
	case Cmd_InputMediaGeoLive971fa843:
		t := m.To_InputMediaGeoLive()
		return t.Encode(layer)
	case Cmd_InputMediaDocument33473058:
		t := m.To_InputMediaDocument()
		return t.Encode(layer)

	default:
		log.Errorf("InputMedia Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputMessage) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputMessageID:
		t := m.To_InputMessageID()
		return t.Encode(layer)
	case Cmd_InputMessageReplyTo:
		t := m.To_InputMessageReplyTo()
		return t.Encode(layer)
	case Cmd_InputMessagePinned:
		t := m.To_InputMessagePinned()
		return t.Encode(layer)
	case Cmd_InputMessageCallbackQuery:
		t := m.To_InputMessageCallbackQuery()
		return t.Encode(layer)

	default:
		log.Errorf("InputMessage Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputNotifyPeer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputNotifyPeer:
		t := m.To_InputNotifyPeer()
		return t.Encode(layer)
	case Cmd_InputNotifyUsers:
		t := m.To_InputNotifyUsers()
		return t.Encode(layer)
	case Cmd_InputNotifyChats:
		t := m.To_InputNotifyChats()
		return t.Encode(layer)
	case Cmd_InputNotifyAll:
		t := m.To_InputNotifyAll()
		return t.Encode(layer)
	case Cmd_InputNotifyBroadcasts:
		t := m.To_InputNotifyBroadcasts()
		return t.Encode(layer)

	default:
		log.Errorf("InputNotifyPeer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputPaymentCredentials) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputPaymentCredentialsSaved:
		t := m.To_InputPaymentCredentialsSaved()
		return t.Encode(layer)
	case Cmd_InputPaymentCredentials:
		t := m.To_InputPaymentCredentials()
		return t.Encode(layer)
	case Cmd_InputPaymentCredentialsApplePay:
		t := m.To_InputPaymentCredentialsApplePay()
		return t.Encode(layer)
	case Cmd_InputPaymentCredentialsAndroidPay:
		t := m.To_InputPaymentCredentialsAndroidPay()
		return t.Encode(layer)

	default:
		log.Errorf("InputPaymentCredentials Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputPeer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputPeerEmpty:
		t := m.To_InputPeerEmpty()
		return t.Encode(layer)
	case Cmd_InputPeerSelf:
		t := m.To_InputPeerSelf()
		return t.Encode(layer)
	case Cmd_InputPeerChat:
		t := m.To_InputPeerChat()
		return t.Encode(layer)
	case Cmd_InputPeerUser:
		t := m.To_InputPeerUser()
		return t.Encode(layer)
	case Cmd_InputPeerChannel:
		t := m.To_InputPeerChannel()
		return t.Encode(layer)
	case Cmd_InputPeerUserFromMessage:
		t := m.To_InputPeerUserFromMessage()
		return t.Encode(layer)
	case Cmd_InputPeerChannelFromMessage:
		t := m.To_InputPeerChannelFromMessage()
		return t.Encode(layer)

	default:
		log.Errorf("InputPeer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputPeerNotifyEvents) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputPeerNotifyEventsEmpty:
		t := m.To_InputPeerNotifyEventsEmpty()
		return t.Encode(layer)
	case Cmd_InputPeerNotifyEventsAll:
		t := m.To_InputPeerNotifyEventsAll()
		return t.Encode(layer)

	default:
		log.Errorf("InputPeerNotifyEvents Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputPeerNotifySettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputPeerNotifySettings:
		t := m.To_InputPeerNotifySettings()
		return t.Encode(layer)

	case Cmd_InputPeerNotifySettings9c3d198e:
		t := m.To_InputPeerNotifySettings()
		return t.Encode(layer)

	default:
		log.Errorf("InputPeerNotifySettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputPhoneCall) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputPhoneCall:
		t := m.To_InputPhoneCall()
		return t.Encode(layer)

	default:
		log.Errorf("InputPhoneCall Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputPhoto) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputPhotoEmpty:
		t := m.To_InputPhotoEmpty()
		return t.Encode(layer)
	case Cmd_InputPhoto:
		t := m.To_InputPhoto()
		return t.Encode(layer)

	case Cmd_InputPhoto3bb3b94a:
		t := m.To_InputPhoto()
		return t.Encode(layer)

	default:
		log.Errorf("InputPhoto Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputPhotoCrop) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputPhotoCrop:
		t := m.To_InputPhotoCrop()
		return t.Encode(layer)
	case Cmd_InputPhotoCropAuto:
		t := m.To_InputPhotoCropAuto()
		return t.Encode(layer)

	default:
		log.Errorf("InputPhotoCrop Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputPrivacyKey) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputPrivacyKeyStatusTimestamp:
		t := m.To_InputPrivacyKeyStatusTimestamp()
		return t.Encode(layer)
	case Cmd_InputPrivacyKeyChatInvite:
		t := m.To_InputPrivacyKeyChatInvite()
		return t.Encode(layer)
	case Cmd_InputPrivacyKeyPhoneCall:
		t := m.To_InputPrivacyKeyPhoneCall()
		return t.Encode(layer)
	case Cmd_InputPrivacyKeyPhoneP2P:
		t := m.To_InputPrivacyKeyPhoneP2P()
		return t.Encode(layer)
	case Cmd_InputPrivacyKeyForwards:
		t := m.To_InputPrivacyKeyForwards()
		return t.Encode(layer)
	case Cmd_InputPrivacyKeyProfilePhoto:
		t := m.To_InputPrivacyKeyProfilePhoto()
		return t.Encode(layer)
	case Cmd_InputPrivacyKeyPhoneNumber:
		t := m.To_InputPrivacyKeyPhoneNumber()
		return t.Encode(layer)
	case Cmd_InputPrivacyKeyAddedByPhone:
		t := m.To_InputPrivacyKeyAddedByPhone()
		return t.Encode(layer)

	default:
		log.Errorf("InputPrivacyKey Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputPrivacyRule) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputPrivacyValueAllowContacts:
		t := m.To_InputPrivacyValueAllowContacts()
		return t.Encode(layer)
	case Cmd_InputPrivacyValueAllowAll:
		t := m.To_InputPrivacyValueAllowAll()
		return t.Encode(layer)
	case Cmd_InputPrivacyValueAllowUsers:
		t := m.To_InputPrivacyValueAllowUsers()
		return t.Encode(layer)
	case Cmd_InputPrivacyValueDisallowContacts:
		t := m.To_InputPrivacyValueDisallowContacts()
		return t.Encode(layer)
	case Cmd_InputPrivacyValueDisallowAll:
		t := m.To_InputPrivacyValueDisallowAll()
		return t.Encode(layer)
	case Cmd_InputPrivacyValueDisallowUsers:
		t := m.To_InputPrivacyValueDisallowUsers()
		return t.Encode(layer)
	case Cmd_InputPrivacyValueAllowChatParticipants:
		t := m.To_InputPrivacyValueAllowChatParticipants()
		return t.Encode(layer)
	case Cmd_InputPrivacyValueDisallowChatParticipants:
		t := m.To_InputPrivacyValueDisallowChatParticipants()
		return t.Encode(layer)

	default:
		log.Errorf("InputPrivacyRule Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputSecureFile) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputSecureFileUploaded:
		t := m.To_InputSecureFileUploaded()
		return t.Encode(layer)
	case Cmd_InputSecureFile:
		t := m.To_InputSecureFile()
		return t.Encode(layer)

	default:
		log.Errorf("InputSecureFile Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputSecureValue) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputSecureValue:
		t := m.To_InputSecureValue()
		return t.Encode(layer)

	default:
		log.Errorf("InputSecureValue Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputSingleMedia) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputSingleMedia:
		t := m.To_InputSingleMedia()
		return t.Encode(layer)

	default:
		log.Errorf("InputSingleMedia Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputStickerSet) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputStickerSetEmpty:
		t := m.To_InputStickerSetEmpty()
		return t.Encode(layer)
	case Cmd_InputStickerSetID:
		t := m.To_InputStickerSetID()
		return t.Encode(layer)
	case Cmd_InputStickerSetShortName:
		t := m.To_InputStickerSetShortName()
		return t.Encode(layer)
	case Cmd_InputStickerSetAnimatedEmoji:
		t := m.To_InputStickerSetAnimatedEmoji()
		return t.Encode(layer)
	case Cmd_InputStickerSetDice:
		t := m.To_InputStickerSetDice()
		return t.Encode(layer)

	case Cmd_InputStickerSetDicee67f520e:
		t := m.To_InputStickerSetDice()
		return t.Encode(layer)

	default:
		log.Errorf("InputStickerSet Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputStickerSetItem) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputStickerSetItem:
		t := m.To_InputStickerSetItem()
		return t.Encode(layer)

	default:
		log.Errorf("InputStickerSetItem Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputStickeredMedia) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputStickeredMediaPhoto:
		t := m.To_InputStickeredMediaPhoto()
		return t.Encode(layer)
	case Cmd_InputStickeredMediaDocument:
		t := m.To_InputStickeredMediaDocument()
		return t.Encode(layer)

	default:
		log.Errorf("InputStickeredMedia Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputTheme) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputTheme:
		t := m.To_InputTheme()
		return t.Encode(layer)
	case Cmd_InputThemeSlug:
		t := m.To_InputThemeSlug()
		return t.Encode(layer)

	default:
		log.Errorf("InputTheme Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputThemeSettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputThemeSettings:
		t := m.To_InputThemeSettings()
		return t.Encode(layer)

	default:
		log.Errorf("InputThemeSettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputUser) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputUserEmpty:
		t := m.To_InputUserEmpty()
		return t.Encode(layer)
	case Cmd_InputUserSelf:
		t := m.To_InputUserSelf()
		return t.Encode(layer)
	case Cmd_InputUser:
		t := m.To_InputUser()
		return t.Encode(layer)
	case Cmd_InputUserFromMessage:
		t := m.To_InputUserFromMessage()
		return t.Encode(layer)

	default:
		log.Errorf("InputUser Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputWallPaper) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputWallPaper:
		t := m.To_InputWallPaper()
		return t.Encode(layer)
	case Cmd_InputWallPaperSlug:
		t := m.To_InputWallPaperSlug()
		return t.Encode(layer)
	case Cmd_InputWallPaperNoFile:
		t := m.To_InputWallPaperNoFile()
		return t.Encode(layer)

	default:
		log.Errorf("InputWallPaper Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputWebDocument) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputWebDocument:
		t := m.To_InputWebDocument()
		return t.Encode(layer)

	default:
		log.Errorf("InputWebDocument Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *InputWebFileLocation) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputWebFileLocation:
		t := m.To_InputWebFileLocation()
		return t.Encode(layer)
	case Cmd_InputWebFileGeoPointLocation:
		t := m.To_InputWebFileGeoPointLocation()
		return t.Encode(layer)

	default:
		log.Errorf("InputWebFileLocation Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Invoice) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Invoice:
		t := m.To_Invoice()
		return t.Encode(layer)

	default:
		log.Errorf("Invoice Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *IpPort) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_IpPort:
		t := m.To_IpPort()
		return t.Encode(layer)
	case Cmd_IpPortSecret:
		t := m.To_IpPortSecret()
		return t.Encode(layer)

	default:
		log.Errorf("IpPort Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *JSONObjectValue) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_JsonObjectValue:
		t := m.To_JsonObjectValue()
		return t.Encode(layer)

	default:
		log.Errorf("JSONObjectValue Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *JSONValue) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_JsonNull:
		t := m.To_JsonNull()
		return t.Encode(layer)
	case Cmd_JsonBool:
		t := m.To_JsonBool()
		return t.Encode(layer)
	case Cmd_JsonNumber:
		t := m.To_JsonNumber()
		return t.Encode(layer)
	case Cmd_JsonString:
		t := m.To_JsonString()
		return t.Encode(layer)
	case Cmd_JsonArray:
		t := m.To_JsonArray()
		return t.Encode(layer)
	case Cmd_JsonObject:
		t := m.To_JsonObject()
		return t.Encode(layer)

	default:
		log.Errorf("JSONValue Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *KeyboardButton) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_KeyboardButton:
		t := m.To_KeyboardButton()
		return t.Encode(layer)
	case Cmd_KeyboardButtonUrl:
		t := m.To_KeyboardButtonUrl()
		return t.Encode(layer)
	case Cmd_KeyboardButtonCallback:
		t := m.To_KeyboardButtonCallback()
		return t.Encode(layer)
	case Cmd_KeyboardButtonRequestPhone:
		t := m.To_KeyboardButtonRequestPhone()
		return t.Encode(layer)
	case Cmd_KeyboardButtonRequestGeoLocation:
		t := m.To_KeyboardButtonRequestGeoLocation()
		return t.Encode(layer)
	case Cmd_KeyboardButtonSwitchInline:
		t := m.To_KeyboardButtonSwitchInline()
		return t.Encode(layer)
	case Cmd_KeyboardButtonGame:
		t := m.To_KeyboardButtonGame()
		return t.Encode(layer)
	case Cmd_KeyboardButtonBuy:
		t := m.To_KeyboardButtonBuy()
		return t.Encode(layer)
	case Cmd_KeyboardButtonUrlAuth:
		t := m.To_KeyboardButtonUrlAuth()
		return t.Encode(layer)
	case Cmd_InputKeyboardButtonUrlAuth:
		t := m.To_InputKeyboardButtonUrlAuth()
		return t.Encode(layer)
	case Cmd_KeyboardButtonRequestPoll:
		t := m.To_KeyboardButtonRequestPoll()
		return t.Encode(layer)

	case Cmd_KeyboardButtonSwitchInlineea1b7a14:
		t := m.To_KeyboardButtonSwitchInline()
		return t.Encode(layer)
	case Cmd_KeyboardButtonCallback35bbdb6b:
		t := m.To_KeyboardButtonCallback()
		return t.Encode(layer)

	default:
		log.Errorf("KeyboardButton Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *KeyboardButtonRow) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_KeyboardButtonRow:
		t := m.To_KeyboardButtonRow()
		return t.Encode(layer)

	default:
		log.Errorf("KeyboardButtonRow Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *LabeledPrice) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_LabeledPrice:
		t := m.To_LabeledPrice()
		return t.Encode(layer)

	default:
		log.Errorf("LabeledPrice Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *LangPackDifference) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_LangPackDifference:
		t := m.To_LangPackDifference()
		return t.Encode(layer)

	default:
		log.Errorf("LangPackDifference Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *LangPackLanguage) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_LangPackLanguage:
		t := m.To_LangPackLanguage()
		return t.Encode(layer)

	case Cmd_LangPackLanguage651b98d:
		t := m.To_LangPackLanguage()
		return t.Encode(layer)
	case Cmd_LangPackLanguageeeca5ce3:
		t := m.To_LangPackLanguage()
		return t.Encode(layer)

	default:
		log.Errorf("LangPackLanguage Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *LangPackString) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_LangPackString:
		t := m.To_LangPackString()
		return t.Encode(layer)
	case Cmd_LangPackStringPluralized:
		t := m.To_LangPackStringPluralized()
		return t.Encode(layer)
	case Cmd_LangPackStringDeleted:
		t := m.To_LangPackStringDeleted()
		return t.Encode(layer)

	default:
		log.Errorf("LangPackString Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MaskCoords) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MaskCoords:
		t := m.To_MaskCoords()
		return t.Encode(layer)

	default:
		log.Errorf("MaskCoords Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Message) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageEmpty:
		t := m.To_MessageEmpty()
		return t.Encode(layer)
	case Cmd_Message:
		t := m.To_Message()
		return t.Encode(layer)
	case Cmd_MessageService:
		t := m.To_MessageService()
		return t.Encode(layer)

	case Cmd_Messagec09be45f:
		t := m.To_Message()
		return t.Encode(layer)
	case Cmd_Message44f9b43d:
		t := m.To_Message()
		return t.Encode(layer)
	case Cmd_Message452c0e65:
		t := m.To_Message()
		return t.Encode(layer)
	case Cmd_MessageService286fa604:
		t := m.To_MessageService()
		return t.Encode(layer)
	case Cmd_Message58ae39c9:
		t := m.To_Message()
		return t.Encode(layer)

	default:
		log.Errorf("Message Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessageAction) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageActionEmpty:
		t := m.To_MessageActionEmpty()
		return t.Encode(layer)
	case Cmd_MessageActionChatCreate:
		t := m.To_MessageActionChatCreate()
		return t.Encode(layer)
	case Cmd_MessageActionChatEditTitle:
		t := m.To_MessageActionChatEditTitle()
		return t.Encode(layer)
	case Cmd_MessageActionChatEditPhoto:
		t := m.To_MessageActionChatEditPhoto()
		return t.Encode(layer)
	case Cmd_MessageActionChatDeletePhoto:
		t := m.To_MessageActionChatDeletePhoto()
		return t.Encode(layer)
	case Cmd_MessageActionChatAddUser:
		t := m.To_MessageActionChatAddUser()
		return t.Encode(layer)
	case Cmd_MessageActionChatDeleteUser:
		t := m.To_MessageActionChatDeleteUser()
		return t.Encode(layer)
	case Cmd_MessageActionChatJoinedByLink:
		t := m.To_MessageActionChatJoinedByLink()
		return t.Encode(layer)
	case Cmd_MessageActionChannelCreate:
		t := m.To_MessageActionChannelCreate()
		return t.Encode(layer)
	case Cmd_MessageActionChatMigrateTo:
		t := m.To_MessageActionChatMigrateTo()
		return t.Encode(layer)
	case Cmd_MessageActionChannelMigrateFrom:
		t := m.To_MessageActionChannelMigrateFrom()
		return t.Encode(layer)
	case Cmd_MessageActionPinMessage:
		t := m.To_MessageActionPinMessage()
		return t.Encode(layer)
	case Cmd_MessageActionHistoryClear:
		t := m.To_MessageActionHistoryClear()
		return t.Encode(layer)
	case Cmd_MessageActionGameScore:
		t := m.To_MessageActionGameScore()
		return t.Encode(layer)
	case Cmd_MessageActionPaymentSentMe:
		t := m.To_MessageActionPaymentSentMe()
		return t.Encode(layer)
	case Cmd_MessageActionPaymentSent:
		t := m.To_MessageActionPaymentSent()
		return t.Encode(layer)
	case Cmd_MessageActionScreenshotTaken:
		t := m.To_MessageActionScreenshotTaken()
		return t.Encode(layer)
	case Cmd_MessageActionCustomAction:
		t := m.To_MessageActionCustomAction()
		return t.Encode(layer)
	case Cmd_MessageActionBotAllowed:
		t := m.To_MessageActionBotAllowed()
		return t.Encode(layer)
	case Cmd_MessageActionSecureValuesSentMe:
		t := m.To_MessageActionSecureValuesSentMe()
		return t.Encode(layer)
	case Cmd_MessageActionSecureValuesSent:
		t := m.To_MessageActionSecureValuesSent()
		return t.Encode(layer)
	case Cmd_MessageActionContactSignUp:
		t := m.To_MessageActionContactSignUp()
		return t.Encode(layer)
	case Cmd_MessageActionPhoneCall:
		t := m.To_MessageActionPhoneCall()
		return t.Encode(layer)
	case Cmd_MessageActionGeoProximityReached:
		t := m.To_MessageActionGeoProximityReached()
		return t.Encode(layer)
	case Cmd_MessageActionGroupCall:
		t := m.To_MessageActionGroupCall()
		return t.Encode(layer)
	case Cmd_MessageActionInviteToGroupCall:
		t := m.To_MessageActionInviteToGroupCall()
		return t.Encode(layer)

	case Cmd_MessageActionContactSignUpf3f25f76:
		t := m.To_MessageActionContactSignUp()
		return t.Encode(layer)

	default:
		log.Errorf("MessageAction Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessageEntity) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageEntityUnknown:
		t := m.To_MessageEntityUnknown()
		return t.Encode(layer)
	case Cmd_MessageEntityMention:
		t := m.To_MessageEntityMention()
		return t.Encode(layer)
	case Cmd_MessageEntityHashtag:
		t := m.To_MessageEntityHashtag()
		return t.Encode(layer)
	case Cmd_MessageEntityBotCommand:
		t := m.To_MessageEntityBotCommand()
		return t.Encode(layer)
	case Cmd_MessageEntityUrl:
		t := m.To_MessageEntityUrl()
		return t.Encode(layer)
	case Cmd_MessageEntityEmail:
		t := m.To_MessageEntityEmail()
		return t.Encode(layer)
	case Cmd_MessageEntityBold:
		t := m.To_MessageEntityBold()
		return t.Encode(layer)
	case Cmd_MessageEntityItalic:
		t := m.To_MessageEntityItalic()
		return t.Encode(layer)
	case Cmd_MessageEntityCode:
		t := m.To_MessageEntityCode()
		return t.Encode(layer)
	case Cmd_MessageEntityPre:
		t := m.To_MessageEntityPre()
		return t.Encode(layer)
	case Cmd_MessageEntityTextUrl:
		t := m.To_MessageEntityTextUrl()
		return t.Encode(layer)
	case Cmd_MessageEntityMentionName:
		t := m.To_MessageEntityMentionName()
		return t.Encode(layer)
	case Cmd_InputMessageEntityMentionName:
		t := m.To_InputMessageEntityMentionName()
		return t.Encode(layer)
	case Cmd_MessageEntityPhone:
		t := m.To_MessageEntityPhone()
		return t.Encode(layer)
	case Cmd_MessageEntityCashtag:
		t := m.To_MessageEntityCashtag()
		return t.Encode(layer)
	case Cmd_MessageEntityUnderline:
		t := m.To_MessageEntityUnderline()
		return t.Encode(layer)
	case Cmd_MessageEntityStrike:
		t := m.To_MessageEntityStrike()
		return t.Encode(layer)
	case Cmd_MessageEntityBlockquote:
		t := m.To_MessageEntityBlockquote()
		return t.Encode(layer)
	case Cmd_MessageEntityBankCard:
		t := m.To_MessageEntityBankCard()
		return t.Encode(layer)

	default:
		log.Errorf("MessageEntity Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessageFwdHeader) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageFwdHeader:
		t := m.To_MessageFwdHeader()
		return t.Encode(layer)

	case Cmd_MessageFwdHeaderc786ddcb:
		t := m.To_MessageFwdHeader()
		return t.Encode(layer)
	case Cmd_MessageFwdHeader559ebe6d:
		t := m.To_MessageFwdHeader()
		return t.Encode(layer)
	case Cmd_MessageFwdHeaderec338270:
		t := m.To_MessageFwdHeader()
		return t.Encode(layer)
	case Cmd_MessageFwdHeader353a686b:
		t := m.To_MessageFwdHeader()
		return t.Encode(layer)
	case Cmd_MessageFwdHeader5f777dce:
		t := m.To_MessageFwdHeader()
		return t.Encode(layer)

	default:
		log.Errorf("MessageFwdHeader Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessageGroup) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageGroup:
		t := m.To_MessageGroup()
		return t.Encode(layer)

	default:
		log.Errorf("MessageGroup Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessageInteractionCounters) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageInteractionCounters:
		t := m.To_MessageInteractionCounters()
		return t.Encode(layer)

	default:
		log.Errorf("MessageInteractionCounters Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessageMedia) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageMediaEmpty:
		t := m.To_MessageMediaEmpty()
		return t.Encode(layer)
	case Cmd_MessageMediaPhoto:
		t := m.To_MessageMediaPhoto()
		return t.Encode(layer)
	case Cmd_MessageMediaGeo:
		t := m.To_MessageMediaGeo()
		return t.Encode(layer)
	case Cmd_MessageMediaContact:
		t := m.To_MessageMediaContact()
		return t.Encode(layer)
	case Cmd_MessageMediaUnsupported:
		t := m.To_MessageMediaUnsupported()
		return t.Encode(layer)
	case Cmd_MessageMediaDocument:
		t := m.To_MessageMediaDocument()
		return t.Encode(layer)
	case Cmd_MessageMediaWebPage:
		t := m.To_MessageMediaWebPage()
		return t.Encode(layer)
	case Cmd_MessageMediaVenue:
		t := m.To_MessageMediaVenue()
		return t.Encode(layer)
	case Cmd_MessageMediaGame:
		t := m.To_MessageMediaGame()
		return t.Encode(layer)
	case Cmd_MessageMediaInvoice:
		t := m.To_MessageMediaInvoice()
		return t.Encode(layer)
	case Cmd_MessageMediaGeoLive:
		t := m.To_MessageMediaGeoLive()
		return t.Encode(layer)
	case Cmd_MessageMediaPoll:
		t := m.To_MessageMediaPoll()
		return t.Encode(layer)
	case Cmd_MessageMediaDice:
		t := m.To_MessageMediaDice()
		return t.Encode(layer)

	case Cmd_MessageMediaPhoto3d8ce53d:
		t := m.To_MessageMediaPhoto()
		return t.Encode(layer)
	case Cmd_MessageMediaDocumentf3e02ea8:
		t := m.To_MessageMediaDocument()
		return t.Encode(layer)
	case Cmd_MessageMediaPhoto695150d7:
		t := m.To_MessageMediaPhoto()
		return t.Encode(layer)
	case Cmd_MessageMediaContactcbf24940:
		t := m.To_MessageMediaContact()
		return t.Encode(layer)
	case Cmd_MessageMediaDocument9cb070d7:
		t := m.To_MessageMediaDocument()
		return t.Encode(layer)
	case Cmd_MessageMediaVenue2ec0533f:
		t := m.To_MessageMediaVenue()
		return t.Encode(layer)
	case Cmd_MessageMediaDice3f7ee58b:
		t := m.To_MessageMediaDice()
		return t.Encode(layer)
	case Cmd_MessageMediaGeoLiveb940c666:
		t := m.To_MessageMediaGeoLive()
		return t.Encode(layer)

	default:
		log.Errorf("MessageMedia Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessageRange) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageRange:
		t := m.To_MessageRange()
		return t.Encode(layer)

	default:
		log.Errorf("MessageRange Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessageReplies) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageReplies:
		t := m.To_MessageReplies()
		return t.Encode(layer)

	default:
		log.Errorf("MessageReplies Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessageReplyHeader) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageReplyHeader:
		t := m.To_MessageReplyHeader()
		return t.Encode(layer)

	default:
		log.Errorf("MessageReplyHeader Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessageUserVote) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageUserVote:
		t := m.To_MessageUserVote()
		return t.Encode(layer)
	case Cmd_MessageUserVoteInputOption:
		t := m.To_MessageUserVoteInputOption()
		return t.Encode(layer)
	case Cmd_MessageUserVoteMultiple:
		t := m.To_MessageUserVoteMultiple()
		return t.Encode(layer)

	case Cmd_MessageUserVotea28e5559:
		t := m.To_MessageUserVote()
		return t.Encode(layer)

	default:
		log.Errorf("MessageUserVote Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessageViews) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessageViews:
		t := m.To_MessageViews()
		return t.Encode(layer)

	default:
		log.Errorf("MessageViews Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MessagesFilter) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputMessagesFilterEmpty:
		t := m.To_InputMessagesFilterEmpty()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterPhotos:
		t := m.To_InputMessagesFilterPhotos()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterVideo:
		t := m.To_InputMessagesFilterVideo()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterPhotoVideo:
		t := m.To_InputMessagesFilterPhotoVideo()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterPhotoVideoDocuments:
		t := m.To_InputMessagesFilterPhotoVideoDocuments()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterDocument:
		t := m.To_InputMessagesFilterDocument()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterUrl:
		t := m.To_InputMessagesFilterUrl()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterGif:
		t := m.To_InputMessagesFilterGif()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterVoice:
		t := m.To_InputMessagesFilterVoice()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterMusic:
		t := m.To_InputMessagesFilterMusic()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterChatPhotos:
		t := m.To_InputMessagesFilterChatPhotos()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterPhoneCalls:
		t := m.To_InputMessagesFilterPhoneCalls()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterRoundVoice:
		t := m.To_InputMessagesFilterRoundVoice()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterRoundVideo:
		t := m.To_InputMessagesFilterRoundVideo()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterMyMentions:
		t := m.To_InputMessagesFilterMyMentions()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterGeo:
		t := m.To_InputMessagesFilterGeo()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterContacts:
		t := m.To_InputMessagesFilterContacts()
		return t.Encode(layer)
	case Cmd_InputMessagesFilterPinned:
		t := m.To_InputMessagesFilterPinned()
		return t.Encode(layer)

	default:
		log.Errorf("MessagesFilter Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_AffectedHistory) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesAffectedHistory:
		t := m.To_MessagesAffectedHistory()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_AffectedHistory Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_AffectedMessages) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesAffectedMessages:
		t := m.To_MessagesAffectedMessages()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_AffectedMessages Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_AllStickers) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesAllStickersNotModified:
		t := m.To_MessagesAllStickersNotModified()
		return t.Encode(layer)
	case Cmd_MessagesAllStickers:
		t := m.To_MessagesAllStickers()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_AllStickers Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_ArchivedStickers) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesArchivedStickers:
		t := m.To_MessagesArchivedStickers()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_ArchivedStickers Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_BotCallbackAnswer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesBotCallbackAnswer:
		t := m.To_MessagesBotCallbackAnswer()
		return t.Encode(layer)

	case Cmd_MessagesBotCallbackAnswer1264f1c6:
		t := m.To_MessagesBotCallbackAnswer()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_BotCallbackAnswer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_BotResults) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesBotResults:
		t := m.To_MessagesBotResults()
		return t.Encode(layer)

	case Cmd_MessagesBotResults256709a6:
		t := m.To_MessagesBotResults()
		return t.Encode(layer)
	case Cmd_MessagesBotResults947ca848:
		t := m.To_MessagesBotResults()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_BotResults Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_ChatFull) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesChatFull:
		t := m.To_MessagesChatFull()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_ChatFull Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_Chats) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesChats:
		t := m.To_MessagesChats()
		return t.Encode(layer)
	case Cmd_MessagesChatsSlice:
		t := m.To_MessagesChatsSlice()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_Chats Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_DhConfig) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesDhConfigNotModified:
		t := m.To_MessagesDhConfigNotModified()
		return t.Encode(layer)
	case Cmd_MessagesDhConfig:
		t := m.To_MessagesDhConfig()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_DhConfig Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_Dialogs) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesDialogs:
		t := m.To_MessagesDialogs()
		return t.Encode(layer)
	case Cmd_MessagesDialogsSlice:
		t := m.To_MessagesDialogsSlice()
		return t.Encode(layer)
	case Cmd_MessagesDialogsNotModified:
		t := m.To_MessagesDialogsNotModified()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_Dialogs Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_DiscussionMessage) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesDiscussionMessage:
		t := m.To_MessagesDiscussionMessage()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_DiscussionMessage Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_FavedStickers) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesFavedStickersNotModified:
		t := m.To_MessagesFavedStickersNotModified()
		return t.Encode(layer)
	case Cmd_MessagesFavedStickers:
		t := m.To_MessagesFavedStickers()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_FavedStickers Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_FeaturedStickers) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesFeaturedStickersNotModified:
		t := m.To_MessagesFeaturedStickersNotModified()
		return t.Encode(layer)
	case Cmd_MessagesFeaturedStickers:
		t := m.To_MessagesFeaturedStickers()
		return t.Encode(layer)

	case Cmd_MessagesFeaturedStickersNotModifiedc6dc0c66:
		t := m.To_MessagesFeaturedStickersNotModified()
		return t.Encode(layer)
	case Cmd_MessagesFeaturedStickersb6abc341:
		t := m.To_MessagesFeaturedStickers()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_FeaturedStickers Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_FoundGifs) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesFoundGifs:
		t := m.To_MessagesFoundGifs()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_FoundGifs Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_FoundStickerSets) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesFoundStickerSetsNotModified:
		t := m.To_MessagesFoundStickerSetsNotModified()
		return t.Encode(layer)
	case Cmd_MessagesFoundStickerSets:
		t := m.To_MessagesFoundStickerSets()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_FoundStickerSets Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_HighScores) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesHighScores:
		t := m.To_MessagesHighScores()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_HighScores Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_InactiveChats) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesInactiveChats:
		t := m.To_MessagesInactiveChats()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_InactiveChats Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_MessageEditData) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesMessageEditData:
		t := m.To_MessagesMessageEditData()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_MessageEditData Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_MessageViews) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesMessageViews:
		t := m.To_MessagesMessageViews()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_MessageViews Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_Messages) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesMessages:
		t := m.To_MessagesMessages()
		return t.Encode(layer)
	case Cmd_MessagesMessagesSlice:
		t := m.To_MessagesMessagesSlice()
		return t.Encode(layer)
	case Cmd_MessagesChannelMessages:
		t := m.To_MessagesChannelMessages()
		return t.Encode(layer)
	case Cmd_MessagesMessagesNotModified:
		t := m.To_MessagesMessagesNotModified()
		return t.Encode(layer)

	case Cmd_MessagesChannelMessagesbc0f17bc:
		t := m.To_MessagesChannelMessages()
		return t.Encode(layer)
	case Cmd_MessagesMessagesSlicea6c47aaa:
		t := m.To_MessagesMessagesSlice()
		return t.Encode(layer)
	case Cmd_MessagesMessagesSlicec8edce1e:
		t := m.To_MessagesMessagesSlice()
		return t.Encode(layer)
	case Cmd_MessagesMessagesSlice3a54685e:
		t := m.To_MessagesMessagesSlice()
		return t.Encode(layer)
	case Cmd_MessagesChannelMessages64479808:
		t := m.To_MessagesChannelMessages()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_Messages Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_PeerDialogs) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesPeerDialogs:
		t := m.To_MessagesPeerDialogs()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_PeerDialogs Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_RecentStickers) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesRecentStickersNotModified:
		t := m.To_MessagesRecentStickersNotModified()
		return t.Encode(layer)
	case Cmd_MessagesRecentStickers:
		t := m.To_MessagesRecentStickers()
		return t.Encode(layer)

	case Cmd_MessagesRecentStickers22f3afb3:
		t := m.To_MessagesRecentStickers()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_RecentStickers Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_SavedGifs) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesSavedGifsNotModified:
		t := m.To_MessagesSavedGifsNotModified()
		return t.Encode(layer)
	case Cmd_MessagesSavedGifs:
		t := m.To_MessagesSavedGifs()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_SavedGifs Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_SearchCounter) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesSearchCounter:
		t := m.To_MessagesSearchCounter()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_SearchCounter Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_SentEncryptedMessage) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesSentEncryptedMessage:
		t := m.To_MessagesSentEncryptedMessage()
		return t.Encode(layer)
	case Cmd_MessagesSentEncryptedFile:
		t := m.To_MessagesSentEncryptedFile()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_SentEncryptedMessage Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_StickerSet) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesStickerSet:
		t := m.To_MessagesStickerSet()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_StickerSet Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_StickerSetInstallResult) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesStickerSetInstallResultSuccess:
		t := m.To_MessagesStickerSetInstallResultSuccess()
		return t.Encode(layer)
	case Cmd_MessagesStickerSetInstallResultArchive:
		t := m.To_MessagesStickerSetInstallResultArchive()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_StickerSetInstallResult Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_Stickers) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesStickersNotModified:
		t := m.To_MessagesStickersNotModified()
		return t.Encode(layer)
	case Cmd_MessagesStickers:
		t := m.To_MessagesStickers()
		return t.Encode(layer)

	case Cmd_MessagesStickerse4599bbd:
		t := m.To_MessagesStickers()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_Stickers Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Messages_VotesList) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MessagesVotesList:
		t := m.To_MessagesVotesList()
		return t.Encode(layer)

	default:
		log.Errorf("Messages_VotesList Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MsgDetailedInfo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MsgDetailedInfo:
		t := m.To_MsgDetailedInfo()
		return t.Encode(layer)
	case Cmd_MsgNewDetailedInfo:
		t := m.To_MsgNewDetailedInfo()
		return t.Encode(layer)

	default:
		log.Errorf("MsgDetailedInfo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MsgResendReq) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MsgResendReq:
		t := m.To_MsgResendReq()
		return t.Encode(layer)

	default:
		log.Errorf("MsgResendReq Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MsgsAck) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MsgsAck:
		t := m.To_MsgsAck()
		return t.Encode(layer)

	default:
		log.Errorf("MsgsAck Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MsgsAllInfo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MsgsAllInfo:
		t := m.To_MsgsAllInfo()
		return t.Encode(layer)

	default:
		log.Errorf("MsgsAllInfo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MsgsStateInfo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MsgsStateInfo:
		t := m.To_MsgsStateInfo()
		return t.Encode(layer)

	default:
		log.Errorf("MsgsStateInfo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *MsgsStateReq) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_MsgsStateReq:
		t := m.To_MsgsStateReq()
		return t.Encode(layer)

	default:
		log.Errorf("MsgsStateReq Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *NearestDc) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_NearestDc:
		t := m.To_NearestDc()
		return t.Encode(layer)

	default:
		log.Errorf("NearestDc Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *NewSession) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_NewSessionCreated:
		t := m.To_NewSessionCreated()
		return t.Encode(layer)

	default:
		log.Errorf("NewSession Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *NotifyPeer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_NotifyPeer:
		t := m.To_NotifyPeer()
		return t.Encode(layer)
	case Cmd_NotifyUsers:
		t := m.To_NotifyUsers()
		return t.Encode(layer)
	case Cmd_NotifyChats:
		t := m.To_NotifyChats()
		return t.Encode(layer)
	case Cmd_NotifyAll:
		t := m.To_NotifyAll()
		return t.Encode(layer)
	case Cmd_NotifyBroadcasts:
		t := m.To_NotifyBroadcasts()
		return t.Encode(layer)

	default:
		log.Errorf("NotifyPeer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Null) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Null:
		t := m.To_Null()
		return t.Encode(layer)

	default:
		log.Errorf("Null Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *P_QInnerData) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PQInnerData:
		t := m.To_PQInnerData()
		return t.Encode(layer)
	case Cmd_PQInnerDataDc:
		t := m.To_PQInnerDataDc()
		return t.Encode(layer)
	case Cmd_PQInnerDataTemp:
		t := m.To_PQInnerDataTemp()
		return t.Encode(layer)
	case Cmd_PQInnerDataTempDc:
		t := m.To_PQInnerDataTempDc()
		return t.Encode(layer)

	default:
		log.Errorf("P_QInnerData Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Page) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PagePart:
		t := m.To_PagePart()
		return t.Encode(layer)
	case Cmd_PageFull:
		t := m.To_PageFull()
		return t.Encode(layer)
	case Cmd_Page:
		t := m.To_Page()
		return t.Encode(layer)

	case Cmd_Pageae891bec:
		t := m.To_Page()
		return t.Encode(layer)
	case Cmd_Page98657f0d:
		t := m.To_Page()
		return t.Encode(layer)

	default:
		log.Errorf("Page Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PageBlock) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PageBlockUnsupported:
		t := m.To_PageBlockUnsupported()
		return t.Encode(layer)
	case Cmd_PageBlockTitle:
		t := m.To_PageBlockTitle()
		return t.Encode(layer)
	case Cmd_PageBlockSubtitle:
		t := m.To_PageBlockSubtitle()
		return t.Encode(layer)
	case Cmd_PageBlockAuthorDate:
		t := m.To_PageBlockAuthorDate()
		return t.Encode(layer)
	case Cmd_PageBlockHeader:
		t := m.To_PageBlockHeader()
		return t.Encode(layer)
	case Cmd_PageBlockSubheader:
		t := m.To_PageBlockSubheader()
		return t.Encode(layer)
	case Cmd_PageBlockParagraph:
		t := m.To_PageBlockParagraph()
		return t.Encode(layer)
	case Cmd_PageBlockPreformatted:
		t := m.To_PageBlockPreformatted()
		return t.Encode(layer)
	case Cmd_PageBlockFooter:
		t := m.To_PageBlockFooter()
		return t.Encode(layer)
	case Cmd_PageBlockDivider:
		t := m.To_PageBlockDivider()
		return t.Encode(layer)
	case Cmd_PageBlockAnchor:
		t := m.To_PageBlockAnchor()
		return t.Encode(layer)
	case Cmd_PageBlockList:
		t := m.To_PageBlockList()
		return t.Encode(layer)
	case Cmd_PageBlockBlockquote:
		t := m.To_PageBlockBlockquote()
		return t.Encode(layer)
	case Cmd_PageBlockPullquote:
		t := m.To_PageBlockPullquote()
		return t.Encode(layer)
	case Cmd_PageBlockPhoto:
		t := m.To_PageBlockPhoto()
		return t.Encode(layer)
	case Cmd_PageBlockVideo:
		t := m.To_PageBlockVideo()
		return t.Encode(layer)
	case Cmd_PageBlockCover:
		t := m.To_PageBlockCover()
		return t.Encode(layer)
	case Cmd_PageBlockEmbed:
		t := m.To_PageBlockEmbed()
		return t.Encode(layer)
	case Cmd_PageBlockEmbedPost:
		t := m.To_PageBlockEmbedPost()
		return t.Encode(layer)
	case Cmd_PageBlockCollage:
		t := m.To_PageBlockCollage()
		return t.Encode(layer)
	case Cmd_PageBlockSlideshow:
		t := m.To_PageBlockSlideshow()
		return t.Encode(layer)
	case Cmd_PageBlockChannel:
		t := m.To_PageBlockChannel()
		return t.Encode(layer)
	case Cmd_PageBlockAudio:
		t := m.To_PageBlockAudio()
		return t.Encode(layer)
	case Cmd_PageBlockKicker:
		t := m.To_PageBlockKicker()
		return t.Encode(layer)
	case Cmd_PageBlockTable:
		t := m.To_PageBlockTable()
		return t.Encode(layer)
	case Cmd_PageBlockOrderedList:
		t := m.To_PageBlockOrderedList()
		return t.Encode(layer)
	case Cmd_PageBlockDetails:
		t := m.To_PageBlockDetails()
		return t.Encode(layer)
	case Cmd_PageBlockRelatedArticles:
		t := m.To_PageBlockRelatedArticles()
		return t.Encode(layer)
	case Cmd_PageBlockMap:
		t := m.To_PageBlockMap()
		return t.Encode(layer)

	case Cmd_PageBlockListe4e88011:
		t := m.To_PageBlockList()
		return t.Encode(layer)
	case Cmd_PageBlockPhoto1759c560:
		t := m.To_PageBlockPhoto()
		return t.Encode(layer)
	case Cmd_PageBlockVideo7c8fe7b6:
		t := m.To_PageBlockVideo()
		return t.Encode(layer)
	case Cmd_PageBlockEmbeda8718dc5:
		t := m.To_PageBlockEmbed()
		return t.Encode(layer)
	case Cmd_PageBlockEmbedPostf259a80b:
		t := m.To_PageBlockEmbedPost()
		return t.Encode(layer)
	case Cmd_PageBlockCollage65a0fa4d:
		t := m.To_PageBlockCollage()
		return t.Encode(layer)
	case Cmd_PageBlockSlideshow31f9590:
		t := m.To_PageBlockSlideshow()
		return t.Encode(layer)
	case Cmd_PageBlockAudio804361ea:
		t := m.To_PageBlockAudio()
		return t.Encode(layer)

	default:
		log.Errorf("PageBlock Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PageCaption) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PageCaption:
		t := m.To_PageCaption()
		return t.Encode(layer)

	default:
		log.Errorf("PageCaption Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PageListItem) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PageListItemText:
		t := m.To_PageListItemText()
		return t.Encode(layer)
	case Cmd_PageListItemBlocks:
		t := m.To_PageListItemBlocks()
		return t.Encode(layer)

	default:
		log.Errorf("PageListItem Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PageListOrderedItem) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PageListOrderedItemText:
		t := m.To_PageListOrderedItemText()
		return t.Encode(layer)
	case Cmd_PageListOrderedItemBlocks:
		t := m.To_PageListOrderedItemBlocks()
		return t.Encode(layer)

	default:
		log.Errorf("PageListOrderedItem Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PageRelatedArticle) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PageRelatedArticle:
		t := m.To_PageRelatedArticle()
		return t.Encode(layer)

	case Cmd_PageRelatedArticleb390dc08:
		t := m.To_PageRelatedArticle()
		return t.Encode(layer)

	default:
		log.Errorf("PageRelatedArticle Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PageTableCell) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PageTableCell:
		t := m.To_PageTableCell()
		return t.Encode(layer)

	default:
		log.Errorf("PageTableCell Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PageTableRow) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PageTableRow:
		t := m.To_PageTableRow()
		return t.Encode(layer)

	default:
		log.Errorf("PageTableRow Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PasswordKdfAlgo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PasswordKdfAlgoUnknown:
		t := m.To_PasswordKdfAlgoUnknown()
		return t.Encode(layer)
	case Cmd_PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow:
		t := m.To_PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow()
		return t.Encode(layer)

	default:
		log.Errorf("PasswordKdfAlgo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PaymentCharge) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PaymentCharge:
		t := m.To_PaymentCharge()
		return t.Encode(layer)

	default:
		log.Errorf("PaymentCharge Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PaymentRequestedInfo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PaymentRequestedInfo:
		t := m.To_PaymentRequestedInfo()
		return t.Encode(layer)

	default:
		log.Errorf("PaymentRequestedInfo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PaymentSavedCredentials) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PaymentSavedCredentialsCard:
		t := m.To_PaymentSavedCredentialsCard()
		return t.Encode(layer)

	default:
		log.Errorf("PaymentSavedCredentials Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Payments_BankCardData) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PaymentsBankCardData:
		t := m.To_PaymentsBankCardData()
		return t.Encode(layer)

	default:
		log.Errorf("Payments_BankCardData Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Payments_PaymentForm) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PaymentsPaymentForm:
		t := m.To_PaymentsPaymentForm()
		return t.Encode(layer)

	default:
		log.Errorf("Payments_PaymentForm Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Payments_PaymentReceipt) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PaymentsPaymentReceipt:
		t := m.To_PaymentsPaymentReceipt()
		return t.Encode(layer)

	default:
		log.Errorf("Payments_PaymentReceipt Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Payments_PaymentResult) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PaymentsPaymentResult:
		t := m.To_PaymentsPaymentResult()
		return t.Encode(layer)
	case Cmd_PaymentsPaymentVerficationNeeded:
		t := m.To_PaymentsPaymentVerficationNeeded()
		return t.Encode(layer)
	case Cmd_PaymentsPaymentVerificationNeeded:
		t := m.To_PaymentsPaymentVerificationNeeded()
		return t.Encode(layer)

	default:
		log.Errorf("Payments_PaymentResult Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Payments_SavedInfo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PaymentsSavedInfo:
		t := m.To_PaymentsSavedInfo()
		return t.Encode(layer)

	default:
		log.Errorf("Payments_SavedInfo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Payments_ValidatedRequestedInfo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PaymentsValidatedRequestedInfo:
		t := m.To_PaymentsValidatedRequestedInfo()
		return t.Encode(layer)

	default:
		log.Errorf("Payments_ValidatedRequestedInfo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Peer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PeerUser:
		t := m.To_PeerUser()
		return t.Encode(layer)
	case Cmd_PeerChat:
		t := m.To_PeerChat()
		return t.Encode(layer)
	case Cmd_PeerChannel:
		t := m.To_PeerChannel()
		return t.Encode(layer)

	default:
		log.Errorf("Peer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PeerBlocked) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PeerBlocked:
		t := m.To_PeerBlocked()
		return t.Encode(layer)

	default:
		log.Errorf("PeerBlocked Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PeerLocated) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PeerLocated:
		t := m.To_PeerLocated()
		return t.Encode(layer)
	case Cmd_PeerSelfLocated:
		t := m.To_PeerSelfLocated()
		return t.Encode(layer)

	default:
		log.Errorf("PeerLocated Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PeerNotifyEvents) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PeerNotifyEventsEmpty:
		t := m.To_PeerNotifyEventsEmpty()
		return t.Encode(layer)
	case Cmd_PeerNotifyEventsAll:
		t := m.To_PeerNotifyEventsAll()
		return t.Encode(layer)

	default:
		log.Errorf("PeerNotifyEvents Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PeerNotifySettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PeerNotifySettingsEmpty:
		t := m.To_PeerNotifySettingsEmpty()
		return t.Encode(layer)
	case Cmd_PeerNotifySettings:
		t := m.To_PeerNotifySettings()
		return t.Encode(layer)

	case Cmd_PeerNotifySettingsaf509d20:
		t := m.To_PeerNotifySettings()
		return t.Encode(layer)

	default:
		log.Errorf("PeerNotifySettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PeerSettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PeerSettings:
		t := m.To_PeerSettings()
		return t.Encode(layer)

	case Cmd_PeerSettings733f2961:
		t := m.To_PeerSettings()
		return t.Encode(layer)

	default:
		log.Errorf("PeerSettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PhoneCall) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PhoneCallEmpty:
		t := m.To_PhoneCallEmpty()
		return t.Encode(layer)
	case Cmd_PhoneCallRequested:
		t := m.To_PhoneCallRequested()
		return t.Encode(layer)
	case Cmd_PhoneCallAccepted:
		t := m.To_PhoneCallAccepted()
		return t.Encode(layer)
	case Cmd_PhoneCall:
		t := m.To_PhoneCall()
		return t.Encode(layer)
	case Cmd_PhoneCallWaiting:
		t := m.To_PhoneCallWaiting()
		return t.Encode(layer)
	case Cmd_PhoneCallDiscarded:
		t := m.To_PhoneCallDiscarded()
		return t.Encode(layer)

	case Cmd_PhoneCalle6f9ddf3:
		t := m.To_PhoneCall()
		return t.Encode(layer)
	case Cmd_PhoneCallRequested87eabb53:
		t := m.To_PhoneCallRequested()
		return t.Encode(layer)
	case Cmd_PhoneCallAccepted997c454a:
		t := m.To_PhoneCallAccepted()
		return t.Encode(layer)
	case Cmd_PhoneCall8742ae7f:
		t := m.To_PhoneCall()
		return t.Encode(layer)

	default:
		log.Errorf("PhoneCall Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PhoneCallDiscardReason) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PhoneCallDiscardReasonMissed:
		t := m.To_PhoneCallDiscardReasonMissed()
		return t.Encode(layer)
	case Cmd_PhoneCallDiscardReasonDisconnect:
		t := m.To_PhoneCallDiscardReasonDisconnect()
		return t.Encode(layer)
	case Cmd_PhoneCallDiscardReasonHangup:
		t := m.To_PhoneCallDiscardReasonHangup()
		return t.Encode(layer)
	case Cmd_PhoneCallDiscardReasonBusy:
		t := m.To_PhoneCallDiscardReasonBusy()
		return t.Encode(layer)

	default:
		log.Errorf("PhoneCallDiscardReason Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PhoneCallProtocol) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PhoneCallProtocol:
		t := m.To_PhoneCallProtocol()
		return t.Encode(layer)

	case Cmd_PhoneCallProtocolfc878fc8:
		t := m.To_PhoneCallProtocol()
		return t.Encode(layer)

	default:
		log.Errorf("PhoneCallProtocol Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PhoneConnection) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PhoneConnection:
		t := m.To_PhoneConnection()
		return t.Encode(layer)
	case Cmd_PhoneConnectionWebrtc:
		t := m.To_PhoneConnectionWebrtc()
		return t.Encode(layer)

	default:
		log.Errorf("PhoneConnection Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Phone_GroupCall) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PhoneGroupCall:
		t := m.To_PhoneGroupCall()
		return t.Encode(layer)

	default:
		log.Errorf("Phone_GroupCall Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Phone_GroupParticipants) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PhoneGroupParticipants:
		t := m.To_PhoneGroupParticipants()
		return t.Encode(layer)

	default:
		log.Errorf("Phone_GroupParticipants Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Phone_PhoneCall) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PhonePhoneCall:
		t := m.To_PhonePhoneCall()
		return t.Encode(layer)

	default:
		log.Errorf("Phone_PhoneCall Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Photo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PhotoEmpty:
		t := m.To_PhotoEmpty()
		return t.Encode(layer)
	case Cmd_Photo:
		t := m.To_Photo()
		return t.Encode(layer)

	case Cmd_Photocded42fe:
		t := m.To_Photo()
		return t.Encode(layer)
	case Cmd_Photo9c477dd8:
		t := m.To_Photo()
		return t.Encode(layer)
	case Cmd_Photod07504a5:
		t := m.To_Photo()
		return t.Encode(layer)
	case Cmd_Photofb197a65:
		t := m.To_Photo()
		return t.Encode(layer)

	default:
		log.Errorf("Photo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PhotoSize) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PhotoSizeEmpty:
		t := m.To_PhotoSizeEmpty()
		return t.Encode(layer)
	case Cmd_PhotoSize:
		t := m.To_PhotoSize()
		return t.Encode(layer)
	case Cmd_PhotoCachedSize:
		t := m.To_PhotoCachedSize()
		return t.Encode(layer)
	case Cmd_PhotoStrippedSize:
		t := m.To_PhotoStrippedSize()
		return t.Encode(layer)
	case Cmd_PhotoSizeProgressive:
		t := m.To_PhotoSizeProgressive()
		return t.Encode(layer)
	case Cmd_PhotoPathSize:
		t := m.To_PhotoPathSize()
		return t.Encode(layer)

	default:
		log.Errorf("PhotoSize Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Photos_Photo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PhotosPhoto:
		t := m.To_PhotosPhoto()
		return t.Encode(layer)

	default:
		log.Errorf("Photos_Photo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Photos_Photos) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PhotosPhotos:
		t := m.To_PhotosPhotos()
		return t.Encode(layer)
	case Cmd_PhotosPhotosSlice:
		t := m.To_PhotosPhotosSlice()
		return t.Encode(layer)

	default:
		log.Errorf("Photos_Photos Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Poll) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Poll:
		t := m.To_Poll()
		return t.Encode(layer)

	case Cmd_Poll86e18161:
		t := m.To_Poll()
		return t.Encode(layer)

	default:
		log.Errorf("Poll Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PollAnswer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PollAnswer:
		t := m.To_PollAnswer()
		return t.Encode(layer)

	default:
		log.Errorf("PollAnswer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PollAnswerVoters) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PollAnswerVoters:
		t := m.To_PollAnswerVoters()
		return t.Encode(layer)

	default:
		log.Errorf("PollAnswerVoters Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PollResults) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PollResults:
		t := m.To_PollResults()
		return t.Encode(layer)

	case Cmd_PollResultsc87024a2:
		t := m.To_PollResults()
		return t.Encode(layer)
	case Cmd_PollResultsbadcc1a3:
		t := m.To_PollResults()
		return t.Encode(layer)

	default:
		log.Errorf("PollResults Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Pong) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Pong:
		t := m.To_Pong()
		return t.Encode(layer)

	default:
		log.Errorf("Pong Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PopularContact) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PopularContact:
		t := m.To_PopularContact()
		return t.Encode(layer)

	default:
		log.Errorf("PopularContact Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PostAddress) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PostAddress:
		t := m.To_PostAddress()
		return t.Encode(layer)

	default:
		log.Errorf("PostAddress Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PrivacyKey) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PrivacyKeyStatusTimestamp:
		t := m.To_PrivacyKeyStatusTimestamp()
		return t.Encode(layer)
	case Cmd_PrivacyKeyChatInvite:
		t := m.To_PrivacyKeyChatInvite()
		return t.Encode(layer)
	case Cmd_PrivacyKeyPhoneCall:
		t := m.To_PrivacyKeyPhoneCall()
		return t.Encode(layer)
	case Cmd_PrivacyKeyPhoneP2P:
		t := m.To_PrivacyKeyPhoneP2P()
		return t.Encode(layer)
	case Cmd_PrivacyKeyForwards:
		t := m.To_PrivacyKeyForwards()
		return t.Encode(layer)
	case Cmd_PrivacyKeyProfilePhoto:
		t := m.To_PrivacyKeyProfilePhoto()
		return t.Encode(layer)
	case Cmd_PrivacyKeyPhoneNumber:
		t := m.To_PrivacyKeyPhoneNumber()
		return t.Encode(layer)
	case Cmd_PrivacyKeyAddedByPhone:
		t := m.To_PrivacyKeyAddedByPhone()
		return t.Encode(layer)

	default:
		log.Errorf("PrivacyKey Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *PrivacyRule) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_PrivacyValueAllowContacts:
		t := m.To_PrivacyValueAllowContacts()
		return t.Encode(layer)
	case Cmd_PrivacyValueAllowAll:
		t := m.To_PrivacyValueAllowAll()
		return t.Encode(layer)
	case Cmd_PrivacyValueAllowUsers:
		t := m.To_PrivacyValueAllowUsers()
		return t.Encode(layer)
	case Cmd_PrivacyValueDisallowContacts:
		t := m.To_PrivacyValueDisallowContacts()
		return t.Encode(layer)
	case Cmd_PrivacyValueDisallowAll:
		t := m.To_PrivacyValueDisallowAll()
		return t.Encode(layer)
	case Cmd_PrivacyValueDisallowUsers:
		t := m.To_PrivacyValueDisallowUsers()
		return t.Encode(layer)
	case Cmd_PrivacyValueAllowChatParticipants:
		t := m.To_PrivacyValueAllowChatParticipants()
		return t.Encode(layer)
	case Cmd_PrivacyValueDisallowChatParticipants:
		t := m.To_PrivacyValueDisallowChatParticipants()
		return t.Encode(layer)

	default:
		log.Errorf("PrivacyRule Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ReceivedNotifyMessage) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ReceivedNotifyMessage:
		t := m.To_ReceivedNotifyMessage()
		return t.Encode(layer)

	default:
		log.Errorf("ReceivedNotifyMessage Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *RecentMeUrl) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_RecentMeUrlUnknown:
		t := m.To_RecentMeUrlUnknown()
		return t.Encode(layer)
	case Cmd_RecentMeUrlUser:
		t := m.To_RecentMeUrlUser()
		return t.Encode(layer)
	case Cmd_RecentMeUrlChat:
		t := m.To_RecentMeUrlChat()
		return t.Encode(layer)
	case Cmd_RecentMeUrlChatInvite:
		t := m.To_RecentMeUrlChatInvite()
		return t.Encode(layer)
	case Cmd_RecentMeUrlStickerSet:
		t := m.To_RecentMeUrlStickerSet()
		return t.Encode(layer)

	default:
		log.Errorf("RecentMeUrl Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ReplyMarkup) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ReplyKeyboardHide:
		t := m.To_ReplyKeyboardHide()
		return t.Encode(layer)
	case Cmd_ReplyKeyboardForceReply:
		t := m.To_ReplyKeyboardForceReply()
		return t.Encode(layer)
	case Cmd_ReplyKeyboardMarkup:
		t := m.To_ReplyKeyboardMarkup()
		return t.Encode(layer)
	case Cmd_ReplyInlineMarkup:
		t := m.To_ReplyInlineMarkup()
		return t.Encode(layer)

	default:
		log.Errorf("ReplyMarkup Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ReportReason) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_InputReportReasonSpam:
		t := m.To_InputReportReasonSpam()
		return t.Encode(layer)
	case Cmd_InputReportReasonViolence:
		t := m.To_InputReportReasonViolence()
		return t.Encode(layer)
	case Cmd_InputReportReasonPornography:
		t := m.To_InputReportReasonPornography()
		return t.Encode(layer)
	case Cmd_InputReportReasonOther:
		t := m.To_InputReportReasonOther()
		return t.Encode(layer)
	case Cmd_InputReportReasonCopyright:
		t := m.To_InputReportReasonCopyright()
		return t.Encode(layer)
	case Cmd_InputReportReasonChildAbuse:
		t := m.To_InputReportReasonChildAbuse()
		return t.Encode(layer)
	case Cmd_InputReportReasonGeoIrrelevant:
		t := m.To_InputReportReasonGeoIrrelevant()
		return t.Encode(layer)

	default:
		log.Errorf("ReportReason Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ResPQ) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ResPQ:
		t := m.To_ResPQ()
		return t.Encode(layer)

	default:
		log.Errorf("ResPQ Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *RestrictionReason) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_RestrictionReason:
		t := m.To_RestrictionReason()
		return t.Encode(layer)

	default:
		log.Errorf("RestrictionReason Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *RichText) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_TextEmpty:
		t := m.To_TextEmpty()
		return t.Encode(layer)
	case Cmd_TextPlain:
		t := m.To_TextPlain()
		return t.Encode(layer)
	case Cmd_TextBold:
		t := m.To_TextBold()
		return t.Encode(layer)
	case Cmd_TextItalic:
		t := m.To_TextItalic()
		return t.Encode(layer)
	case Cmd_TextUnderline:
		t := m.To_TextUnderline()
		return t.Encode(layer)
	case Cmd_TextStrike:
		t := m.To_TextStrike()
		return t.Encode(layer)
	case Cmd_TextFixed:
		t := m.To_TextFixed()
		return t.Encode(layer)
	case Cmd_TextUrl:
		t := m.To_TextUrl()
		return t.Encode(layer)
	case Cmd_TextEmail:
		t := m.To_TextEmail()
		return t.Encode(layer)
	case Cmd_TextConcat:
		t := m.To_TextConcat()
		return t.Encode(layer)
	case Cmd_TextSubscript:
		t := m.To_TextSubscript()
		return t.Encode(layer)
	case Cmd_TextSuperscript:
		t := m.To_TextSuperscript()
		return t.Encode(layer)
	case Cmd_TextMarked:
		t := m.To_TextMarked()
		return t.Encode(layer)
	case Cmd_TextPhone:
		t := m.To_TextPhone()
		return t.Encode(layer)
	case Cmd_TextImage:
		t := m.To_TextImage()
		return t.Encode(layer)
	case Cmd_TextAnchor:
		t := m.To_TextAnchor()
		return t.Encode(layer)

	default:
		log.Errorf("RichText Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *RpcDropAnswer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_RpcAnswerUnknown:
		t := m.To_RpcAnswerUnknown()
		return t.Encode(layer)
	case Cmd_RpcAnswerDroppedRunning:
		t := m.To_RpcAnswerDroppedRunning()
		return t.Encode(layer)
	case Cmd_RpcAnswerDropped:
		t := m.To_RpcAnswerDropped()
		return t.Encode(layer)

	default:
		log.Errorf("RpcDropAnswer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *RpcError) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_RpcError:
		t := m.To_RpcError()
		return t.Encode(layer)

	default:
		log.Errorf("RpcError Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SavedContact) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SavedPhoneContact:
		t := m.To_SavedPhoneContact()
		return t.Encode(layer)

	default:
		log.Errorf("SavedContact Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SecureCredentialsEncrypted) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SecureCredentialsEncrypted:
		t := m.To_SecureCredentialsEncrypted()
		return t.Encode(layer)

	default:
		log.Errorf("SecureCredentialsEncrypted Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SecureData) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SecureData:
		t := m.To_SecureData()
		return t.Encode(layer)

	default:
		log.Errorf("SecureData Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SecureFile) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SecureFileEmpty:
		t := m.To_SecureFileEmpty()
		return t.Encode(layer)
	case Cmd_SecureFile:
		t := m.To_SecureFile()
		return t.Encode(layer)

	default:
		log.Errorf("SecureFile Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SecurePasswordKdfAlgo) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SecurePasswordKdfAlgoUnknown:
		t := m.To_SecurePasswordKdfAlgoUnknown()
		return t.Encode(layer)
	case Cmd_SecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000:
		t := m.To_SecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000()
		return t.Encode(layer)
	case Cmd_SecurePasswordKdfAlgoSHA512:
		t := m.To_SecurePasswordKdfAlgoSHA512()
		return t.Encode(layer)

	default:
		log.Errorf("SecurePasswordKdfAlgo Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SecurePlainData) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SecurePlainPhone:
		t := m.To_SecurePlainPhone()
		return t.Encode(layer)
	case Cmd_SecurePlainEmail:
		t := m.To_SecurePlainEmail()
		return t.Encode(layer)

	default:
		log.Errorf("SecurePlainData Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SecureRequiredType) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SecureRequiredType:
		t := m.To_SecureRequiredType()
		return t.Encode(layer)
	case Cmd_SecureRequiredTypeOneOf:
		t := m.To_SecureRequiredTypeOneOf()
		return t.Encode(layer)

	default:
		log.Errorf("SecureRequiredType Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SecureSecretSettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SecureSecretSettings:
		t := m.To_SecureSecretSettings()
		return t.Encode(layer)

	default:
		log.Errorf("SecureSecretSettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SecureValue) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SecureValue:
		t := m.To_SecureValue()
		return t.Encode(layer)

	default:
		log.Errorf("SecureValue Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SecureValueError) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SecureValueErrorData:
		t := m.To_SecureValueErrorData()
		return t.Encode(layer)
	case Cmd_SecureValueErrorFrontSide:
		t := m.To_SecureValueErrorFrontSide()
		return t.Encode(layer)
	case Cmd_SecureValueErrorReverseSide:
		t := m.To_SecureValueErrorReverseSide()
		return t.Encode(layer)
	case Cmd_SecureValueErrorSelfie:
		t := m.To_SecureValueErrorSelfie()
		return t.Encode(layer)
	case Cmd_SecureValueErrorFile:
		t := m.To_SecureValueErrorFile()
		return t.Encode(layer)
	case Cmd_SecureValueErrorFiles:
		t := m.To_SecureValueErrorFiles()
		return t.Encode(layer)
	case Cmd_SecureValueError:
		t := m.To_SecureValueError()
		return t.Encode(layer)
	case Cmd_SecureValueErrorTranslationFile:
		t := m.To_SecureValueErrorTranslationFile()
		return t.Encode(layer)
	case Cmd_SecureValueErrorTranslationFiles:
		t := m.To_SecureValueErrorTranslationFiles()
		return t.Encode(layer)

	default:
		log.Errorf("SecureValueError Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SecureValueHash) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SecureValueHash:
		t := m.To_SecureValueHash()
		return t.Encode(layer)

	default:
		log.Errorf("SecureValueHash Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SecureValueType) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SecureValueTypePersonalDetails:
		t := m.To_SecureValueTypePersonalDetails()
		return t.Encode(layer)
	case Cmd_SecureValueTypePassport:
		t := m.To_SecureValueTypePassport()
		return t.Encode(layer)
	case Cmd_SecureValueTypeDriverLicense:
		t := m.To_SecureValueTypeDriverLicense()
		return t.Encode(layer)
	case Cmd_SecureValueTypeIdentityCard:
		t := m.To_SecureValueTypeIdentityCard()
		return t.Encode(layer)
	case Cmd_SecureValueTypeInternalPassport:
		t := m.To_SecureValueTypeInternalPassport()
		return t.Encode(layer)
	case Cmd_SecureValueTypeAddress:
		t := m.To_SecureValueTypeAddress()
		return t.Encode(layer)
	case Cmd_SecureValueTypeUtilityBill:
		t := m.To_SecureValueTypeUtilityBill()
		return t.Encode(layer)
	case Cmd_SecureValueTypeBankStatement:
		t := m.To_SecureValueTypeBankStatement()
		return t.Encode(layer)
	case Cmd_SecureValueTypeRentalAgreement:
		t := m.To_SecureValueTypeRentalAgreement()
		return t.Encode(layer)
	case Cmd_SecureValueTypePassportRegistration:
		t := m.To_SecureValueTypePassportRegistration()
		return t.Encode(layer)
	case Cmd_SecureValueTypeTemporaryRegistration:
		t := m.To_SecureValueTypeTemporaryRegistration()
		return t.Encode(layer)
	case Cmd_SecureValueTypePhone:
		t := m.To_SecureValueTypePhone()
		return t.Encode(layer)
	case Cmd_SecureValueTypeEmail:
		t := m.To_SecureValueTypeEmail()
		return t.Encode(layer)

	default:
		log.Errorf("SecureValueType Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SendMessageAction) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_SendMessageTypingAction:
		t := m.To_SendMessageTypingAction()
		return t.Encode(layer)
	case Cmd_SendMessageCancelAction:
		t := m.To_SendMessageCancelAction()
		return t.Encode(layer)
	case Cmd_SendMessageRecordVideoAction:
		t := m.To_SendMessageRecordVideoAction()
		return t.Encode(layer)
	case Cmd_SendMessageUploadVideoAction:
		t := m.To_SendMessageUploadVideoAction()
		return t.Encode(layer)
	case Cmd_SendMessageRecordAudioAction:
		t := m.To_SendMessageRecordAudioAction()
		return t.Encode(layer)
	case Cmd_SendMessageUploadAudioAction:
		t := m.To_SendMessageUploadAudioAction()
		return t.Encode(layer)
	case Cmd_SendMessageUploadPhotoAction:
		t := m.To_SendMessageUploadPhotoAction()
		return t.Encode(layer)
	case Cmd_SendMessageUploadDocumentAction:
		t := m.To_SendMessageUploadDocumentAction()
		return t.Encode(layer)
	case Cmd_SendMessageGeoLocationAction:
		t := m.To_SendMessageGeoLocationAction()
		return t.Encode(layer)
	case Cmd_SendMessageChooseContactAction:
		t := m.To_SendMessageChooseContactAction()
		return t.Encode(layer)
	case Cmd_SendMessageGamePlayAction:
		t := m.To_SendMessageGamePlayAction()
		return t.Encode(layer)
	case Cmd_SendMessageRecordRoundAction:
		t := m.To_SendMessageRecordRoundAction()
		return t.Encode(layer)
	case Cmd_SendMessageUploadRoundAction:
		t := m.To_SendMessageUploadRoundAction()
		return t.Encode(layer)
	case Cmd_SpeakingInGroupCallAction:
		t := m.To_SpeakingInGroupCallAction()
		return t.Encode(layer)

	default:
		log.Errorf("SendMessageAction Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Server_DHInnerData) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Server_DHInnerData:
		t := m.To_Server_DHInnerData()
		return t.Encode(layer)

	default:
		log.Errorf("Server_DHInnerData Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Server_DH_Params) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_Server_DHParamsFail:
		t := m.To_Server_DHParamsFail()
		return t.Encode(layer)
	case Cmd_Server_DHParamsOk:
		t := m.To_Server_DHParamsOk()
		return t.Encode(layer)

	default:
		log.Errorf("Server_DH_Params Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *SetClient_DHParamsAnswer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_DhGenOk:
		t := m.To_DhGenOk()
		return t.Encode(layer)
	case Cmd_DhGenRetry:
		t := m.To_DhGenRetry()
		return t.Encode(layer)
	case Cmd_DhGenFail:
		t := m.To_DhGenFail()
		return t.Encode(layer)

	default:
		log.Errorf("SetClient_DHParamsAnswer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ShippingOption) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ShippingOption:
		t := m.To_ShippingOption()
		return t.Encode(layer)

	default:
		log.Errorf("ShippingOption Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *StatsAbsValueAndPrev) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StatsAbsValueAndPrev:
		t := m.To_StatsAbsValueAndPrev()
		return t.Encode(layer)

	default:
		log.Errorf("StatsAbsValueAndPrev Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *StatsDateRangeDays) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StatsDateRangeDays:
		t := m.To_StatsDateRangeDays()
		return t.Encode(layer)

	default:
		log.Errorf("StatsDateRangeDays Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *StatsGraph) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StatsGraphAsync:
		t := m.To_StatsGraphAsync()
		return t.Encode(layer)
	case Cmd_StatsGraphError:
		t := m.To_StatsGraphError()
		return t.Encode(layer)
	case Cmd_StatsGraph:
		t := m.To_StatsGraph()
		return t.Encode(layer)

	default:
		log.Errorf("StatsGraph Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *StatsGroupTopAdmin) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StatsGroupTopAdmin:
		t := m.To_StatsGroupTopAdmin()
		return t.Encode(layer)

	default:
		log.Errorf("StatsGroupTopAdmin Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *StatsGroupTopInviter) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StatsGroupTopInviter:
		t := m.To_StatsGroupTopInviter()
		return t.Encode(layer)

	default:
		log.Errorf("StatsGroupTopInviter Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *StatsGroupTopPoster) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StatsGroupTopPoster:
		t := m.To_StatsGroupTopPoster()
		return t.Encode(layer)

	default:
		log.Errorf("StatsGroupTopPoster Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *StatsPercentValue) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StatsPercentValue:
		t := m.To_StatsPercentValue()
		return t.Encode(layer)

	default:
		log.Errorf("StatsPercentValue Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *StatsURL) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StatsURL:
		t := m.To_StatsURL()
		return t.Encode(layer)

	default:
		log.Errorf("StatsURL Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Stats_BroadcastStats) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StatsBroadcastStats:
		t := m.To_StatsBroadcastStats()
		return t.Encode(layer)

	default:
		log.Errorf("Stats_BroadcastStats Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Stats_MegagroupStats) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StatsMegagroupStats:
		t := m.To_StatsMegagroupStats()
		return t.Encode(layer)

	default:
		log.Errorf("Stats_MegagroupStats Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Stats_MessageStats) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StatsMessageStats:
		t := m.To_StatsMessageStats()
		return t.Encode(layer)

	default:
		log.Errorf("Stats_MessageStats Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *StickerPack) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StickerPack:
		t := m.To_StickerPack()
		return t.Encode(layer)

	default:
		log.Errorf("StickerPack Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *StickerSet) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StickerSet:
		t := m.To_StickerSet()
		return t.Encode(layer)

	case Cmd_StickerSet5585a139:
		t := m.To_StickerSet()
		return t.Encode(layer)
	case Cmd_StickerSet6a90bcb7:
		t := m.To_StickerSet()
		return t.Encode(layer)
	case Cmd_StickerSeteeb46f27:
		t := m.To_StickerSet()
		return t.Encode(layer)
	case Cmd_StickerSet40e237a8:
		t := m.To_StickerSet()
		return t.Encode(layer)

	default:
		log.Errorf("StickerSet Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *StickerSetCovered) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StickerSetCovered:
		t := m.To_StickerSetCovered()
		return t.Encode(layer)
	case Cmd_StickerSetMultiCovered:
		t := m.To_StickerSetMultiCovered()
		return t.Encode(layer)

	default:
		log.Errorf("StickerSetCovered Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Storage_FileType) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_StorageFileUnknown:
		t := m.To_StorageFileUnknown()
		return t.Encode(layer)
	case Cmd_StorageFilePartial:
		t := m.To_StorageFilePartial()
		return t.Encode(layer)
	case Cmd_StorageFileJpeg:
		t := m.To_StorageFileJpeg()
		return t.Encode(layer)
	case Cmd_StorageFileGif:
		t := m.To_StorageFileGif()
		return t.Encode(layer)
	case Cmd_StorageFilePng:
		t := m.To_StorageFilePng()
		return t.Encode(layer)
	case Cmd_StorageFilePdf:
		t := m.To_StorageFilePdf()
		return t.Encode(layer)
	case Cmd_StorageFileMp3:
		t := m.To_StorageFileMp3()
		return t.Encode(layer)
	case Cmd_StorageFileMov:
		t := m.To_StorageFileMov()
		return t.Encode(layer)
	case Cmd_StorageFileMp4:
		t := m.To_StorageFileMp4()
		return t.Encode(layer)
	case Cmd_StorageFileWebp:
		t := m.To_StorageFileWebp()
		return t.Encode(layer)

	default:
		log.Errorf("Storage_FileType Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Theme) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ThemeDocumentNotModified:
		t := m.To_ThemeDocumentNotModified()
		return t.Encode(layer)
	case Cmd_Theme:
		t := m.To_Theme()
		return t.Encode(layer)

	case Cmd_Themef7d90ce0:
		t := m.To_Theme()
		return t.Encode(layer)
	case Cmd_Theme28f1114:
		t := m.To_Theme()
		return t.Encode(layer)

	default:
		log.Errorf("Theme Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *ThemeSettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_ThemeSettings:
		t := m.To_ThemeSettings()
		return t.Encode(layer)

	default:
		log.Errorf("ThemeSettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *TopPeer) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_TopPeer:
		t := m.To_TopPeer()
		return t.Encode(layer)

	default:
		log.Errorf("TopPeer Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *TopPeerCategory) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_TopPeerCategoryBotsPM:
		t := m.To_TopPeerCategoryBotsPM()
		return t.Encode(layer)
	case Cmd_TopPeerCategoryBotsInline:
		t := m.To_TopPeerCategoryBotsInline()
		return t.Encode(layer)
	case Cmd_TopPeerCategoryCorrespondents:
		t := m.To_TopPeerCategoryCorrespondents()
		return t.Encode(layer)
	case Cmd_TopPeerCategoryGroups:
		t := m.To_TopPeerCategoryGroups()
		return t.Encode(layer)
	case Cmd_TopPeerCategoryChannels:
		t := m.To_TopPeerCategoryChannels()
		return t.Encode(layer)
	case Cmd_TopPeerCategoryPhoneCalls:
		t := m.To_TopPeerCategoryPhoneCalls()
		return t.Encode(layer)
	case Cmd_TopPeerCategoryForwardUsers:
		t := m.To_TopPeerCategoryForwardUsers()
		return t.Encode(layer)
	case Cmd_TopPeerCategoryForwardChats:
		t := m.To_TopPeerCategoryForwardChats()
		return t.Encode(layer)

	default:
		log.Errorf("TopPeerCategory Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *TopPeerCategoryPeers) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_TopPeerCategoryPeers:
		t := m.To_TopPeerCategoryPeers()
		return t.Encode(layer)

	default:
		log.Errorf("TopPeerCategoryPeers Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *True) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_True:
		t := m.To_True()
		return t.Encode(layer)

	default:
		log.Errorf("True Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Update) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UpdateNewMessage:
		t := m.To_UpdateNewMessage()
		return t.Encode(layer)
	case Cmd_UpdateMessageID:
		t := m.To_UpdateMessageID()
		return t.Encode(layer)
	case Cmd_UpdateDeleteMessages:
		t := m.To_UpdateDeleteMessages()
		return t.Encode(layer)
	case Cmd_UpdateUserTyping:
		t := m.To_UpdateUserTyping()
		return t.Encode(layer)
	case Cmd_UpdateChatUserTyping:
		t := m.To_UpdateChatUserTyping()
		return t.Encode(layer)
	case Cmd_UpdateChatParticipants:
		t := m.To_UpdateChatParticipants()
		return t.Encode(layer)
	case Cmd_UpdateUserStatus:
		t := m.To_UpdateUserStatus()
		return t.Encode(layer)
	case Cmd_UpdateUserName:
		t := m.To_UpdateUserName()
		return t.Encode(layer)
	case Cmd_UpdateUserPhoto:
		t := m.To_UpdateUserPhoto()
		return t.Encode(layer)
	case Cmd_UpdateContactRegistered:
		t := m.To_UpdateContactRegistered()
		return t.Encode(layer)
	case Cmd_UpdateContactLink:
		t := m.To_UpdateContactLink()
		return t.Encode(layer)
	case Cmd_UpdateNewEncryptedMessage:
		t := m.To_UpdateNewEncryptedMessage()
		return t.Encode(layer)
	case Cmd_UpdateEncryptedChatTyping:
		t := m.To_UpdateEncryptedChatTyping()
		return t.Encode(layer)
	case Cmd_UpdateEncryption:
		t := m.To_UpdateEncryption()
		return t.Encode(layer)
	case Cmd_UpdateEncryptedMessagesRead:
		t := m.To_UpdateEncryptedMessagesRead()
		return t.Encode(layer)
	case Cmd_UpdateChatParticipantAdd:
		t := m.To_UpdateChatParticipantAdd()
		return t.Encode(layer)
	case Cmd_UpdateChatParticipantDelete:
		t := m.To_UpdateChatParticipantDelete()
		return t.Encode(layer)
	case Cmd_UpdateDcOptions:
		t := m.To_UpdateDcOptions()
		return t.Encode(layer)
	case Cmd_UpdateUserBlocked:
		t := m.To_UpdateUserBlocked()
		return t.Encode(layer)
	case Cmd_UpdateNotifySettings:
		t := m.To_UpdateNotifySettings()
		return t.Encode(layer)
	case Cmd_UpdateServiceNotification:
		t := m.To_UpdateServiceNotification()
		return t.Encode(layer)
	case Cmd_UpdatePrivacy:
		t := m.To_UpdatePrivacy()
		return t.Encode(layer)
	case Cmd_UpdateUserPhone:
		t := m.To_UpdateUserPhone()
		return t.Encode(layer)
	case Cmd_UpdateReadHistoryInbox:
		t := m.To_UpdateReadHistoryInbox()
		return t.Encode(layer)
	case Cmd_UpdateReadHistoryOutbox:
		t := m.To_UpdateReadHistoryOutbox()
		return t.Encode(layer)
	case Cmd_UpdateWebPage:
		t := m.To_UpdateWebPage()
		return t.Encode(layer)
	case Cmd_UpdateReadMessagesContents:
		t := m.To_UpdateReadMessagesContents()
		return t.Encode(layer)
	case Cmd_UpdateChannelTooLong:
		t := m.To_UpdateChannelTooLong()
		return t.Encode(layer)
	case Cmd_UpdateChannel:
		t := m.To_UpdateChannel()
		return t.Encode(layer)
	case Cmd_UpdateNewChannelMessage:
		t := m.To_UpdateNewChannelMessage()
		return t.Encode(layer)
	case Cmd_UpdateReadChannelInbox:
		t := m.To_UpdateReadChannelInbox()
		return t.Encode(layer)
	case Cmd_UpdateDeleteChannelMessages:
		t := m.To_UpdateDeleteChannelMessages()
		return t.Encode(layer)
	case Cmd_UpdateChannelMessageViews:
		t := m.To_UpdateChannelMessageViews()
		return t.Encode(layer)
	case Cmd_UpdateChatAdmins:
		t := m.To_UpdateChatAdmins()
		return t.Encode(layer)
	case Cmd_UpdateChatParticipantAdmin:
		t := m.To_UpdateChatParticipantAdmin()
		return t.Encode(layer)
	case Cmd_UpdateNewStickerSet:
		t := m.To_UpdateNewStickerSet()
		return t.Encode(layer)
	case Cmd_UpdateStickerSetsOrder:
		t := m.To_UpdateStickerSetsOrder()
		return t.Encode(layer)
	case Cmd_UpdateStickerSets:
		t := m.To_UpdateStickerSets()
		return t.Encode(layer)
	case Cmd_UpdateSavedGifs:
		t := m.To_UpdateSavedGifs()
		return t.Encode(layer)
	case Cmd_UpdateBotInlineQuery:
		t := m.To_UpdateBotInlineQuery()
		return t.Encode(layer)
	case Cmd_UpdateBotInlineSend:
		t := m.To_UpdateBotInlineSend()
		return t.Encode(layer)
	case Cmd_UpdateEditChannelMessage:
		t := m.To_UpdateEditChannelMessage()
		return t.Encode(layer)
	case Cmd_UpdateChannelPinnedMessage:
		t := m.To_UpdateChannelPinnedMessage()
		return t.Encode(layer)
	case Cmd_UpdateBotCallbackQuery:
		t := m.To_UpdateBotCallbackQuery()
		return t.Encode(layer)
	case Cmd_UpdateEditMessage:
		t := m.To_UpdateEditMessage()
		return t.Encode(layer)
	case Cmd_UpdateInlineBotCallbackQuery:
		t := m.To_UpdateInlineBotCallbackQuery()
		return t.Encode(layer)
	case Cmd_UpdateReadChannelOutbox:
		t := m.To_UpdateReadChannelOutbox()
		return t.Encode(layer)
	case Cmd_UpdateDraftMessage:
		t := m.To_UpdateDraftMessage()
		return t.Encode(layer)
	case Cmd_UpdateReadFeaturedStickers:
		t := m.To_UpdateReadFeaturedStickers()
		return t.Encode(layer)
	case Cmd_UpdateRecentStickers:
		t := m.To_UpdateRecentStickers()
		return t.Encode(layer)
	case Cmd_UpdateConfig:
		t := m.To_UpdateConfig()
		return t.Encode(layer)
	case Cmd_UpdatePtsChanged:
		t := m.To_UpdatePtsChanged()
		return t.Encode(layer)
	case Cmd_UpdateChannelWebPage:
		t := m.To_UpdateChannelWebPage()
		return t.Encode(layer)
	case Cmd_UpdateDialogPinned:
		t := m.To_UpdateDialogPinned()
		return t.Encode(layer)
	case Cmd_UpdatePinnedDialogs:
		t := m.To_UpdatePinnedDialogs()
		return t.Encode(layer)
	case Cmd_UpdateBotWebhookJSON:
		t := m.To_UpdateBotWebhookJSON()
		return t.Encode(layer)
	case Cmd_UpdateBotWebhookJSONQuery:
		t := m.To_UpdateBotWebhookJSONQuery()
		return t.Encode(layer)
	case Cmd_UpdateBotShippingQuery:
		t := m.To_UpdateBotShippingQuery()
		return t.Encode(layer)
	case Cmd_UpdateBotPrecheckoutQuery:
		t := m.To_UpdateBotPrecheckoutQuery()
		return t.Encode(layer)
	case Cmd_UpdatePhoneCall:
		t := m.To_UpdatePhoneCall()
		return t.Encode(layer)
	case Cmd_UpdateLangPackTooLong:
		t := m.To_UpdateLangPackTooLong()
		return t.Encode(layer)
	case Cmd_UpdateLangPack:
		t := m.To_UpdateLangPack()
		return t.Encode(layer)
	case Cmd_UpdateFavedStickers:
		t := m.To_UpdateFavedStickers()
		return t.Encode(layer)
	case Cmd_UpdateChannelReadMessagesContents:
		t := m.To_UpdateChannelReadMessagesContents()
		return t.Encode(layer)
	case Cmd_UpdateContactsReset:
		t := m.To_UpdateContactsReset()
		return t.Encode(layer)
	case Cmd_UpdateNewAuthorization:
		t := m.To_UpdateNewAuthorization()
		return t.Encode(layer)
	case Cmd_UpdateChannelGroup:
		t := m.To_UpdateChannelGroup()
		return t.Encode(layer)
	case Cmd_UpdateChannelAvailableMessages:
		t := m.To_UpdateChannelAvailableMessages()
		return t.Encode(layer)
	case Cmd_UpdateDialogUnreadMark:
		t := m.To_UpdateDialogUnreadMark()
		return t.Encode(layer)
	case Cmd_UpdateUserPinnedMessage:
		t := m.To_UpdateUserPinnedMessage()
		return t.Encode(layer)
	case Cmd_UpdateChatPinnedMessage:
		t := m.To_UpdateChatPinnedMessage()
		return t.Encode(layer)
	case Cmd_UpdateMessagePoll:
		t := m.To_UpdateMessagePoll()
		return t.Encode(layer)
	case Cmd_UpdateChatDefaultBannedRights:
		t := m.To_UpdateChatDefaultBannedRights()
		return t.Encode(layer)
	case Cmd_UpdateFolderPeers:
		t := m.To_UpdateFolderPeers()
		return t.Encode(layer)
	case Cmd_UpdatePeerSettings:
		t := m.To_UpdatePeerSettings()
		return t.Encode(layer)
	case Cmd_UpdatePeerLocated:
		t := m.To_UpdatePeerLocated()
		return t.Encode(layer)
	case Cmd_UpdateNewScheduledMessage:
		t := m.To_UpdateNewScheduledMessage()
		return t.Encode(layer)
	case Cmd_UpdateDeleteScheduledMessages:
		t := m.To_UpdateDeleteScheduledMessages()
		return t.Encode(layer)
	case Cmd_UpdateTheme:
		t := m.To_UpdateTheme()
		return t.Encode(layer)
	case Cmd_UpdateGeoLiveViewed:
		t := m.To_UpdateGeoLiveViewed()
		return t.Encode(layer)
	case Cmd_UpdateLoginToken:
		t := m.To_UpdateLoginToken()
		return t.Encode(layer)
	case Cmd_UpdateMessagePollVote:
		t := m.To_UpdateMessagePollVote()
		return t.Encode(layer)
	case Cmd_UpdateDialogFilter:
		t := m.To_UpdateDialogFilter()
		return t.Encode(layer)
	case Cmd_UpdateDialogFilterOrder:
		t := m.To_UpdateDialogFilterOrder()
		return t.Encode(layer)
	case Cmd_UpdateDialogFilters:
		t := m.To_UpdateDialogFilters()
		return t.Encode(layer)
	case Cmd_UpdatePhoneCallSignalingData:
		t := m.To_UpdatePhoneCallSignalingData()
		return t.Encode(layer)
	case Cmd_UpdateChannelParticipant:
		t := m.To_UpdateChannelParticipant()
		return t.Encode(layer)
	case Cmd_UpdateChannelMessageForwards:
		t := m.To_UpdateChannelMessageForwards()
		return t.Encode(layer)
	case Cmd_UpdateReadChannelDiscussionInbox:
		t := m.To_UpdateReadChannelDiscussionInbox()
		return t.Encode(layer)
	case Cmd_UpdateReadChannelDiscussionOutbox:
		t := m.To_UpdateReadChannelDiscussionOutbox()
		return t.Encode(layer)
	case Cmd_UpdatePeerBlocked:
		t := m.To_UpdatePeerBlocked()
		return t.Encode(layer)
	case Cmd_UpdateChannelUserTyping:
		t := m.To_UpdateChannelUserTyping()
		return t.Encode(layer)
	case Cmd_UpdatePinnedMessages:
		t := m.To_UpdatePinnedMessages()
		return t.Encode(layer)
	case Cmd_UpdatePinnedChannelMessages:
		t := m.To_UpdatePinnedChannelMessages()
		return t.Encode(layer)
	case Cmd_UpdateChat:
		t := m.To_UpdateChat()
		return t.Encode(layer)
	case Cmd_UpdateGroupCallParticipants:
		t := m.To_UpdateGroupCallParticipants()
		return t.Encode(layer)
	case Cmd_UpdateGroupCall:
		t := m.To_UpdateGroupCall()
		return t.Encode(layer)

	case Cmd_UpdateServiceNotification382dd3e4:
		t := m.To_UpdateServiceNotification()
		return t.Encode(layer)
	case Cmd_UpdateStickerSetsOrderf0dfb451:
		t := m.To_UpdateStickerSetsOrder()
		return t.Encode(layer)
	case Cmd_UpdateBotCallbackQuerya68c688c:
		t := m.To_UpdateBotCallbackQuery()
		return t.Encode(layer)
	case Cmd_UpdateInlineBotCallbackQuery2cbd95af:
		t := m.To_UpdateInlineBotCallbackQuery()
		return t.Encode(layer)
	case Cmd_UpdateDialogPinned19d27f3c:
		t := m.To_UpdateDialogPinned()
		return t.Encode(layer)
	case Cmd_UpdatePinnedDialogsea4cb65b:
		t := m.To_UpdatePinnedDialogs()
		return t.Encode(layer)
	case Cmd_UpdateLangPackTooLong46560264:
		t := m.To_UpdateLangPackTooLong()
		return t.Encode(layer)
	case Cmd_UpdateChatPinnedMessagee10db349:
		t := m.To_UpdateChatPinnedMessage()
		return t.Encode(layer)
	case Cmd_UpdateReadHistoryInbox9c974fdf:
		t := m.To_UpdateReadHistoryInbox()
		return t.Encode(layer)
	case Cmd_UpdateReadChannelInbox330b5424:
		t := m.To_UpdateReadChannelInbox()
		return t.Encode(layer)
	case Cmd_UpdateDialogPinned6e6fe51c:
		t := m.To_UpdateDialogPinned()
		return t.Encode(layer)
	case Cmd_UpdatePinnedDialogsfa0f3ca2:
		t := m.To_UpdatePinnedDialogs()
		return t.Encode(layer)
	case Cmd_UpdateBotInlineQuery3f2038db:
		t := m.To_UpdateBotInlineQuery()
		return t.Encode(layer)

	default:
		log.Errorf("Update Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Updates) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UpdatesTooLong:
		t := m.To_UpdatesTooLong()
		return t.Encode(layer)
	case Cmd_UpdateShortMessage:
		t := m.To_UpdateShortMessage()
		return t.Encode(layer)
	case Cmd_UpdateShortChatMessage:
		t := m.To_UpdateShortChatMessage()
		return t.Encode(layer)
	case Cmd_UpdateShort:
		t := m.To_UpdateShort()
		return t.Encode(layer)
	case Cmd_UpdatesCombined:
		t := m.To_UpdatesCombined()
		return t.Encode(layer)
	case Cmd_Updates:
		t := m.To_Updates()
		return t.Encode(layer)
	case Cmd_UpdateShortSentMessage:
		t := m.To_UpdateShortSentMessage()
		return t.Encode(layer)

	case Cmd_UpdateShortMessage2296d2c8:
		t := m.To_UpdateShortMessage()
		return t.Encode(layer)
	case Cmd_UpdateShortChatMessage402d5dbb:
		t := m.To_UpdateShortChatMessage()
		return t.Encode(layer)

	default:
		log.Errorf("Updates Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Updates_ChannelDifference) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UpdatesChannelDifferenceEmpty:
		t := m.To_UpdatesChannelDifferenceEmpty()
		return t.Encode(layer)
	case Cmd_UpdatesChannelDifferenceTooLong:
		t := m.To_UpdatesChannelDifferenceTooLong()
		return t.Encode(layer)
	case Cmd_UpdatesChannelDifference:
		t := m.To_UpdatesChannelDifference()
		return t.Encode(layer)

	case Cmd_UpdatesChannelDifferenceTooLong5e167646:
		t := m.To_UpdatesChannelDifferenceTooLong()
		return t.Encode(layer)
	case Cmd_UpdatesChannelDifferenceTooLonga4bcc6fe:
		t := m.To_UpdatesChannelDifferenceTooLong()
		return t.Encode(layer)

	default:
		log.Errorf("Updates_ChannelDifference Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Updates_Difference) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UpdatesDifferenceEmpty:
		t := m.To_UpdatesDifferenceEmpty()
		return t.Encode(layer)
	case Cmd_UpdatesDifference:
		t := m.To_UpdatesDifference()
		return t.Encode(layer)
	case Cmd_UpdatesDifferenceSlice:
		t := m.To_UpdatesDifferenceSlice()
		return t.Encode(layer)
	case Cmd_UpdatesDifferenceTooLong:
		t := m.To_UpdatesDifferenceTooLong()
		return t.Encode(layer)

	default:
		log.Errorf("Updates_Difference Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Updates_State) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UpdatesState:
		t := m.To_UpdatesState()
		return t.Encode(layer)

	default:
		log.Errorf("Updates_State Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Upload_CdnFile) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UploadCdnFileReuploadNeeded:
		t := m.To_UploadCdnFileReuploadNeeded()
		return t.Encode(layer)
	case Cmd_UploadCdnFile:
		t := m.To_UploadCdnFile()
		return t.Encode(layer)

	default:
		log.Errorf("Upload_CdnFile Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Upload_File) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UploadFile:
		t := m.To_UploadFile()
		return t.Encode(layer)
	case Cmd_UploadFileCdnRedirect:
		t := m.To_UploadFileCdnRedirect()
		return t.Encode(layer)

	case Cmd_UploadFileCdnRedirectf18cda44:
		t := m.To_UploadFileCdnRedirect()
		return t.Encode(layer)

	default:
		log.Errorf("Upload_File Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Upload_WebFile) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UploadWebFile:
		t := m.To_UploadWebFile()
		return t.Encode(layer)

	default:
		log.Errorf("Upload_WebFile Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *UrlAuthResult) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UrlAuthResultRequest:
		t := m.To_UrlAuthResultRequest()
		return t.Encode(layer)
	case Cmd_UrlAuthResultAccepted:
		t := m.To_UrlAuthResultAccepted()
		return t.Encode(layer)
	case Cmd_UrlAuthResultDefault:
		t := m.To_UrlAuthResultDefault()
		return t.Encode(layer)

	default:
		log.Errorf("UrlAuthResult Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *User) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UserEmpty:
		t := m.To_UserEmpty()
		return t.Encode(layer)
	case Cmd_User:
		t := m.To_User()
		return t.Encode(layer)

	case Cmd_Userd10d979a:
		t := m.To_User()
		return t.Encode(layer)
	case Cmd_User938458c1:
		t := m.To_User()
		return t.Encode(layer)

	default:
		log.Errorf("User Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *UserFull) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UserFull:
		t := m.To_UserFull()
		return t.Encode(layer)

	case Cmd_UserFull5932fc03:
		t := m.To_UserFull()
		return t.Encode(layer)
	case Cmd_UserFull8ea4a881:
		t := m.To_UserFull()
		return t.Encode(layer)
	case Cmd_UserFull745559cc:
		t := m.To_UserFull()
		return t.Encode(layer)
	case Cmd_UserFulledf17c12:
		t := m.To_UserFull()
		return t.Encode(layer)

	default:
		log.Errorf("UserFull Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *UserProfilePhoto) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UserProfilePhotoEmpty:
		t := m.To_UserProfilePhotoEmpty()
		return t.Encode(layer)
	case Cmd_UserProfilePhoto:
		t := m.To_UserProfilePhoto()
		return t.Encode(layer)

	case Cmd_UserProfilePhotoecd75d8c:
		t := m.To_UserProfilePhoto()
		return t.Encode(layer)
	case Cmd_UserProfilePhoto69d3ab26:
		t := m.To_UserProfilePhoto()
		return t.Encode(layer)

	default:
		log.Errorf("UserProfilePhoto Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *UserStatus) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_UserStatusEmpty:
		t := m.To_UserStatusEmpty()
		return t.Encode(layer)
	case Cmd_UserStatusOnline:
		t := m.To_UserStatusOnline()
		return t.Encode(layer)
	case Cmd_UserStatusOffline:
		t := m.To_UserStatusOffline()
		return t.Encode(layer)
	case Cmd_UserStatusRecently:
		t := m.To_UserStatusRecently()
		return t.Encode(layer)
	case Cmd_UserStatusLastWeek:
		t := m.To_UserStatusLastWeek()
		return t.Encode(layer)
	case Cmd_UserStatusLastMonth:
		t := m.To_UserStatusLastMonth()
		return t.Encode(layer)

	default:
		log.Errorf("UserStatus Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *VideoSize) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_VideoSize:
		t := m.To_VideoSize()
		return t.Encode(layer)

	case Cmd_VideoSizee831c556:
		t := m.To_VideoSize()
		return t.Encode(layer)

	default:
		log.Errorf("VideoSize Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *WallPaper) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_WallPaper:
		t := m.To_WallPaper()
		return t.Encode(layer)
	case Cmd_WallPaperSolid:
		t := m.To_WallPaperSolid()
		return t.Encode(layer)
	case Cmd_WallPaperNoFile:
		t := m.To_WallPaperNoFile()
		return t.Encode(layer)

	case Cmd_WallPaperf04f91ec:
		t := m.To_WallPaper()
		return t.Encode(layer)
	case Cmd_WallPapera437c3ed:
		t := m.To_WallPaper()
		return t.Encode(layer)

	default:
		log.Errorf("WallPaper Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *WallPaperSettings) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_WallPaperSettings:
		t := m.To_WallPaperSettings()
		return t.Encode(layer)

	case Cmd_WallPaperSettings5086cf8:
		t := m.To_WallPaperSettings()
		return t.Encode(layer)

	default:
		log.Errorf("WallPaperSettings Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Wallet_KeySecretSalt) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_WalletSecretSalt:
		t := m.To_WalletSecretSalt()
		return t.Encode(layer)

	default:
		log.Errorf("Wallet_KeySecretSalt Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *Wallet_LiteResponse) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_WalletLiteResponse:
		t := m.To_WalletLiteResponse()
		return t.Encode(layer)

	default:
		log.Errorf("Wallet_LiteResponse Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *WebAuthorization) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_WebAuthorization:
		t := m.To_WebAuthorization()
		return t.Encode(layer)

	default:
		log.Errorf("WebAuthorization Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *WebDocument) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_WebDocument:
		t := m.To_WebDocument()
		return t.Encode(layer)
	case Cmd_WebDocumentNoProxy:
		t := m.To_WebDocumentNoProxy()
		return t.Encode(layer)

	case Cmd_WebDocument1c570ed1:
		t := m.To_WebDocument()
		return t.Encode(layer)

	default:
		log.Errorf("WebDocument Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *WebPage) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_WebPageEmpty:
		t := m.To_WebPageEmpty()
		return t.Encode(layer)
	case Cmd_WebPagePending:
		t := m.To_WebPagePending()
		return t.Encode(layer)
	case Cmd_WebPage:
		t := m.To_WebPage()
		return t.Encode(layer)
	case Cmd_WebPageNotModified:
		t := m.To_WebPageNotModified()
		return t.Encode(layer)

	case Cmd_WebPageca820ed7:
		t := m.To_WebPage()
		return t.Encode(layer)
	case Cmd_WebPagefa64e172:
		t := m.To_WebPage()
		return t.Encode(layer)
	case Cmd_WebPagee89c45b2:
		t := m.To_WebPage()
		return t.Encode(layer)
	case Cmd_WebPageNotModified7311ca11:
		t := m.To_WebPageNotModified()
		return t.Encode(layer)

	default:
		log.Errorf("WebPage Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}

func (m *WebPageAttribute) Encode(layer int32) []byte {
	switch m.GetCmd() {
	case Cmd_WebPageAttributeTheme:
		t := m.To_WebPageAttributeTheme()
		return t.Encode(layer)

	default:
		log.Errorf("WebPageAttribute Encode Cmd error: %x", m.GetCmd().ToUInt32())
		return nil
	}
}
