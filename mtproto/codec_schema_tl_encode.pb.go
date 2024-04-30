/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2023/06/12 18:03
 * @File : codec_schema_tl_encode.pb.go
 * @Desc :
 *
 */

package mtproto

import (
	"novachat_engine/pkg/log"
)

func (m *TLAccessPointRule) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccessPointRule, layer)

	switch cmd {
	case Cmd_AccessPointRule:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccessPointRule))
		x.String(m.GetPhonePrefixRules())
		x.Int(m.GetDcId())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetIps())))
		for _, v := range m.GetIps() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid AccessPointRule encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountDaysTTL) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountDaysTTL, layer)

	switch cmd {
	case Cmd_AccountDaysTTL:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountDaysTTL))
		x.Int(m.GetDays())

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountDaysTTL encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountAuthorizationForm) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountAuthorizationForm, layer)

	switch cmd {
	case Cmd_AccountAuthorizationForm:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountAuthorizationForm))
		// flags
		var flags int32 = 0
		if len(m.GetPrivacyPolicyUrl()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetRequiredTypes())))
		for _, v := range m.GetRequiredTypes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetValues())))
		for _, v := range m.GetValues() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetErrors())))
		for _, v := range m.GetErrors() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if len(m.GetPrivacyPolicyUrl()) > 0 {
			x.String(m.GetPrivacyPolicyUrl())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountAuthorizationForm encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountAuthorizations) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountAuthorizations, layer)

	switch cmd {
	case Cmd_AccountAuthorizations:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountAuthorizations))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAuthorizations())))
		for _, v := range m.GetAuthorizations() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountAuthorizations encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountAutoDownloadSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountAutoDownloadSettings, layer)

	switch cmd {
	case Cmd_AccountAutoDownloadSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountAutoDownloadSettings))
		x.Bytes(m.GetLow().Encode(layer))
		x.Bytes(m.GetMedium().Encode(layer))
		x.Bytes(m.GetHigh().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountAutoDownloadSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountContentSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountContentSettings, layer)

	switch cmd {
	case Cmd_AccountContentSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountContentSettings))
		// flags
		var flags int32 = 0
		if m.GetSensitiveEnabled() == true {
			flags |= 1 << 0
		}
		if m.GetSensitiveCanChange() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountContentSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountNoPassword) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountNoPassword, layer)

	switch cmd {
	case Cmd_AccountNoPassword:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountNoPassword))
		x.StringBytes(m.GetNewSalt())
		x.String(m.GetEmailUnconfirmedPattern())

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountNoPassword encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountPassword) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountPassword, layer)

	switch cmd {
	case Cmd_AccountPassword:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountPassword))
		x.StringBytes(m.GetCurrentSalt())
		x.StringBytes(m.GetNewSalt())
		x.String(m.GetHint())
		x.Bytes(m.GetHasRecovery7C18141C71().Encode(layer))
		x.String(m.GetEmailUnconfirmedPattern())

		return x.GetBuf()
	case Cmd_AccountPasswordad2641f8:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountPasswordad2641f8))
		// flags
		var flags int32 = 0
		if m.GetHasRecoveryAD2641F885() == true {
			flags |= 1 << 0
		}
		if m.GetHasSecureValues() == true {
			flags |= 1 << 1
		}
		if m.GetHasPassword() == true {
			flags |= 1 << 2
		}
		if m.GetCurrentAlgo() != nil {
			flags |= 1 << 2
		}
		if len(m.GetSrp_B()) > 0 {
			flags |= 1 << 2
		}
		if m.GetSrpId() != 0 {
			flags |= 1 << 2
		}
		if len(m.GetHint()) > 0 {
			flags |= 1 << 3
		}
		if len(m.GetEmailUnconfirmedPattern()) > 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		if m.GetCurrentAlgo() != nil {
			x.Bytes(m.GetCurrentAlgo().Encode(layer))
		}
		if len(m.GetSrp_B()) > 0 {
			x.StringBytes(m.GetSrp_B())
		}
		if m.GetSrpId() != 0 {
			x.Long(m.GetSrpId())
		}
		if len(m.GetHint()) > 0 {
			x.String(m.GetHint())
		}
		if len(m.GetEmailUnconfirmedPattern()) > 0 {
			x.String(m.GetEmailUnconfirmedPattern())
		}
		x.Bytes(m.GetNewAlgo().Encode(layer))
		x.Bytes(m.GetNewSecureAlgo().Encode(layer))
		x.StringBytes(m.GetSecureRandom())

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountPassword encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountPasswordInputSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountPasswordInputSettings, layer)

	switch cmd {
	case Cmd_AccountPasswordInputSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountPasswordInputSettings))
		// flags
		var flags int32 = 0
		if len(m.GetNewSalt()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetNewPasswordHash()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetHint()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetEmail()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		if len(m.GetNewSalt()) > 0 {
			x.StringBytes(m.GetNewSalt())
		}
		if len(m.GetNewPasswordHash()) > 0 {
			x.StringBytes(m.GetNewPasswordHash())
		}
		if len(m.GetHint()) > 0 {
			x.String(m.GetHint())
		}
		if len(m.GetEmail()) > 0 {
			x.String(m.GetEmail())
		}

		return x.GetBuf()
	case Cmd_AccountPasswordInputSettingsc23727c9:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountPasswordInputSettingsc23727c9))
		// flags
		var flags int32 = 0
		if m.GetNewAlgo() != nil {
			flags |= 1 << 0
		}
		if len(m.GetNewPasswordHash()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetHint()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetEmail()) > 0 {
			flags |= 1 << 1
		}
		if m.GetNewSecureSettings() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.GetNewAlgo() != nil {
			x.Bytes(m.GetNewAlgo().Encode(layer))
		}
		if len(m.GetNewPasswordHash()) > 0 {
			x.StringBytes(m.GetNewPasswordHash())
		}
		if len(m.GetHint()) > 0 {
			x.String(m.GetHint())
		}
		if len(m.GetEmail()) > 0 {
			x.String(m.GetEmail())
		}
		if m.GetNewSecureSettings() != nil {
			x.Bytes(m.GetNewSecureSettings().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountPasswordInputSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountPasswordSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountPasswordSettings, layer)

	switch cmd {
	case Cmd_AccountPasswordSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountPasswordSettings))
		x.String(m.GetEmail())

		return x.GetBuf()
	case Cmd_AccountPasswordSettings9a5c33e5:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountPasswordSettings9a5c33e5))
		// flags
		var flags int32 = 0
		if len(m.GetEmail()) > 0 {
			flags |= 1 << 0
		}
		if m.GetSecureSettings() != nil {
			flags |= 1 << 1
		}
		x.Int(flags)

		if len(m.GetEmail()) > 0 {
			x.String(m.GetEmail())
		}
		if m.GetSecureSettings() != nil {
			x.Bytes(m.GetSecureSettings().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountPasswordSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountPrivacyRules) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountPrivacyRules, layer)

	switch cmd {
	case Cmd_AccountPrivacyRules:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountPrivacyRules))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetRules())))
		for _, v := range m.GetRules() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_AccountPrivacyRules50a04e45:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountPrivacyRules50a04e45))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetRules())))
		for _, v := range m.GetRules() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountPrivacyRules encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountSentEmailCode) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountSentEmailCode, layer)

	switch cmd {
	case Cmd_AccountSentEmailCode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountSentEmailCode))
		x.String(m.GetEmailPattern())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountSentEmailCode encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountTakeout) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountTakeout, layer)

	switch cmd {
	case Cmd_AccountTakeout:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountTakeout))
		x.Long(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountTakeout encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountThemesNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountThemesNotModified, layer)

	switch cmd {
	case Cmd_AccountThemesNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountThemesNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountThemesNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountThemes) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountThemes, layer)

	switch cmd {
	case Cmd_AccountThemes:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountThemes))
		x.Int(m.GetHash())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetThemes())))
		for _, v := range m.GetThemes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountThemes encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountTmpPassword) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountTmpPassword, layer)

	switch cmd {
	case Cmd_AccountTmpPassword:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountTmpPassword))
		x.StringBytes(m.GetTmpPassword())
		x.Int(m.GetValidUntil())

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountTmpPassword encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountWallPapersNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountWallPapersNotModified, layer)

	switch cmd {
	case Cmd_AccountWallPapersNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountWallPapersNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountWallPapersNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountWallPapers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountWallPapers, layer)

	switch cmd {
	case Cmd_AccountWallPapers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountWallPapers))
		x.Int(m.GetHash())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetWallpapers())))
		for _, v := range m.GetWallpapers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountWallPapers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAccountWebAuthorizations) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAccountWebAuthorizations, layer)

	switch cmd {
	case Cmd_AccountWebAuthorizations:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AccountWebAuthorizations))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAuthorizations())))
		for _, v := range m.GetAuthorizations() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid AccountWebAuthorizations encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthAuthorization) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthAuthorization, layer)

	switch cmd {
	case Cmd_AuthAuthorization:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthAuthorization))
		// flags
		var flags int32 = 0
		if m.GetTmpSessions() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetTmpSessions() != 0 {
			x.Int(m.GetTmpSessions())
		}
		x.Bytes(m.GetUser().Encode(layer))

		return x.GetBuf()
	case Cmd_AuthAuthorizationff036af1:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthAuthorizationff036af1))
		x.Bytes(m.GetUser().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthAuthorization encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthAuthorizationSignUpRequired) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthAuthorizationSignUpRequired, layer)

	switch cmd {
	case Cmd_AuthAuthorizationSignUpRequired:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthAuthorizationSignUpRequired))
		// flags
		var flags int32 = 0
		if m.GetTermsOfService() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetTermsOfService() != nil {
			x.Bytes(m.GetTermsOfService().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthAuthorizationSignUpRequired encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthCheckedPhone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthCheckedPhone, layer)

	switch cmd {
	case Cmd_AuthCheckedPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthCheckedPhone))
		x.Bytes(m.GetPhoneRegistered().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthCheckedPhone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthCodeTypeSms) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthCodeTypeSms, layer)

	switch cmd {
	case Cmd_AuthCodeTypeSms:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthCodeTypeSms))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthCodeTypeSms encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthCodeTypeCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthCodeTypeCall, layer)

	switch cmd {
	case Cmd_AuthCodeTypeCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthCodeTypeCall))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthCodeTypeCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthCodeTypeFlashCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthCodeTypeFlashCall, layer)

	switch cmd {
	case Cmd_AuthCodeTypeFlashCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthCodeTypeFlashCall))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthCodeTypeFlashCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthExportedAuthorization) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthExportedAuthorization, layer)

	switch cmd {
	case Cmd_AuthExportedAuthorization:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthExportedAuthorization))
		x.Int(m.GetId())
		x.StringBytes(m.GetBytes())

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthExportedAuthorization encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthLoginToken) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthLoginToken, layer)

	switch cmd {
	case Cmd_AuthLoginToken:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthLoginToken))
		x.Int(m.GetExpires())
		x.StringBytes(m.GetToken())

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthLoginToken encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthLoginTokenMigrateTo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthLoginTokenMigrateTo, layer)

	switch cmd {
	case Cmd_AuthLoginTokenMigrateTo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthLoginTokenMigrateTo))
		x.Int(m.GetDcId())
		x.StringBytes(m.GetToken())

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthLoginTokenMigrateTo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthLoginTokenSuccess) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthLoginTokenSuccess, layer)

	switch cmd {
	case Cmd_AuthLoginTokenSuccess:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthLoginTokenSuccess))
		x.Bytes(m.GetAuthorization().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthLoginTokenSuccess encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthPasswordRecovery) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthPasswordRecovery, layer)

	switch cmd {
	case Cmd_AuthPasswordRecovery:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthPasswordRecovery))
		x.String(m.GetEmailPattern())

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthPasswordRecovery encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthSentCode) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthSentCode, layer)

	switch cmd {
	case Cmd_AuthSentCode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSentCode))
		// flags
		var flags int32 = 0
		if m.GetPhoneRegistered() == true {
			flags |= 1 << 0
		}
		if m.GetNextType() != nil {
			flags |= 1 << 1
		}
		if m.GetTimeout() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetType().Encode(layer))
		x.String(m.GetPhoneCodeHash())
		if m.GetNextType() != nil {
			x.Bytes(m.GetNextType().Encode(layer))
		}
		if m.GetTimeout() != 0 {
			x.Int(m.GetTimeout())
		}

		return x.GetBuf()
	case Cmd_AuthSentCode38faab5f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSentCode38faab5f))
		// flags
		var flags int32 = 0
		if m.GetPhoneRegistered() == true {
			flags |= 1 << 0
		}
		if m.GetNextType() != nil {
			flags |= 1 << 1
		}
		if m.GetTimeout() != 0 {
			flags |= 1 << 2
		}
		if m.GetTermsOfService() != nil {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Bytes(m.GetType().Encode(layer))
		x.String(m.GetPhoneCodeHash())
		if m.GetNextType() != nil {
			x.Bytes(m.GetNextType().Encode(layer))
		}
		if m.GetTimeout() != 0 {
			x.Int(m.GetTimeout())
		}
		if m.GetTermsOfService() != nil {
			x.Bytes(m.GetTermsOfService().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthSentCode encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthSentCodeTypeApp) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthSentCodeTypeApp, layer)

	switch cmd {
	case Cmd_AuthSentCodeTypeApp:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSentCodeTypeApp))
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthSentCodeTypeApp encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthSentCodeTypeSms) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthSentCodeTypeSms, layer)

	switch cmd {
	case Cmd_AuthSentCodeTypeSms:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSentCodeTypeSms))
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthSentCodeTypeSms encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthSentCodeTypeCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthSentCodeTypeCall, layer)

	switch cmd {
	case Cmd_AuthSentCodeTypeCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSentCodeTypeCall))
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthSentCodeTypeCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthSentCodeTypeFlashCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthSentCodeTypeFlashCall, layer)

	switch cmd {
	case Cmd_AuthSentCodeTypeFlashCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AuthSentCodeTypeFlashCall))
		x.String(m.GetPattern())

		return x.GetBuf()

	default:
		log.Errorf("invalid AuthSentCodeTypeFlashCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAuthorization) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAuthorization, layer)

	switch cmd {
	case Cmd_Authorization:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Authorization))
		x.Long(m.GetHash())
		x.Int(m.GetFlags())
		x.String(m.GetDeviceModel())
		x.String(m.GetPlatform())
		x.String(m.GetSystemVersion())
		x.Int(m.GetApiId())
		x.String(m.GetAppName())
		x.String(m.GetAppVersion())
		x.Int(m.GetDateCreated())
		x.Int(m.GetDateActive())
		x.String(m.GetIp())
		x.String(m.GetCountry())
		x.String(m.GetRegion())

		return x.GetBuf()
	case Cmd_Authorizationad01d61d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Authorizationad01d61d))
		// flags
		var flags int32 = 0
		if m.GetCurrent() == true {
			flags |= 1 << 0
		}
		if m.GetOfficialApp() == true {
			flags |= 1 << 1
		}
		if m.GetPasswordPending() == true {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Long(m.GetHash())
		x.String(m.GetDeviceModel())
		x.String(m.GetPlatform())
		x.String(m.GetSystemVersion())
		x.Int(m.GetApiId())
		x.String(m.GetAppName())
		x.String(m.GetAppVersion())
		x.Int(m.GetDateCreated())
		x.Int(m.GetDateActive())
		x.String(m.GetIp())
		x.String(m.GetCountry())
		x.String(m.GetRegion())

		return x.GetBuf()

	default:
		log.Errorf("invalid Authorization encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLAutoDownloadSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassAutoDownloadSettings, layer)

	switch cmd {
	case Cmd_AutoDownloadSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AutoDownloadSettings))
		// flags
		var flags int32 = 0
		if m.GetDisabled() == true {
			flags |= 1 << 0
		}
		if m.GetVideoPreloadLarge() == true {
			flags |= 1 << 1
		}
		if m.GetAudioPreloadNext() == true {
			flags |= 1 << 2
		}
		if m.GetPhonecallsLessData() == true {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Int(m.GetPhotoSizeMax())
		x.Int(m.GetVideoSizeMax())
		x.Int(m.GetFileSizeMax())

		return x.GetBuf()
	case Cmd_AutoDownloadSettingse04232f3:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_AutoDownloadSettingse04232f3))
		// flags
		var flags int32 = 0
		if m.GetDisabled() == true {
			flags |= 1 << 0
		}
		if m.GetVideoPreloadLarge() == true {
			flags |= 1 << 1
		}
		if m.GetAudioPreloadNext() == true {
			flags |= 1 << 2
		}
		if m.GetPhonecallsLessData() == true {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Int(m.GetPhotoSizeMax())
		x.Int(m.GetVideoSizeMax())
		x.Int(m.GetFileSizeMax())
		x.Int(m.GetVideoUploadMaxbitrate())

		return x.GetBuf()

	default:
		log.Errorf("invalid AutoDownloadSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBadMsgNotification) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBadMsgNotification, layer)

	switch cmd {
	case Cmd_BadMsgNotification:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BadMsgNotification))
		x.Long(m.GetBadMsgId())
		x.Int(m.GetBadMsgSeqno())
		x.Int(m.GetErrorCode())

		return x.GetBuf()

	default:
		log.Errorf("invalid BadMsgNotification encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBadServerSalt) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBadServerSalt, layer)

	switch cmd {
	case Cmd_BadServerSalt:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BadServerSalt))
		x.Long(m.GetBadMsgId())
		x.Int(m.GetBadMsgSeqno())
		x.Int(m.GetErrorCode())
		x.Long(m.GetNewServerSalt())

		return x.GetBuf()

	default:
		log.Errorf("invalid BadServerSalt encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBankCardOpenUrl) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBankCardOpenUrl, layer)

	switch cmd {
	case Cmd_BankCardOpenUrl:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BankCardOpenUrl))
		x.String(m.GetUrl())
		x.String(m.GetName())

		return x.GetBuf()

	default:
		log.Errorf("invalid BankCardOpenUrl encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBaseThemeClassic) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBaseThemeClassic, layer)

	switch cmd {
	case Cmd_BaseThemeClassic:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BaseThemeClassic))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid BaseThemeClassic encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBaseThemeDay) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBaseThemeDay, layer)

	switch cmd {
	case Cmd_BaseThemeDay:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BaseThemeDay))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid BaseThemeDay encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBaseThemeNight) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBaseThemeNight, layer)

	switch cmd {
	case Cmd_BaseThemeNight:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BaseThemeNight))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid BaseThemeNight encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBaseThemeTinted) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBaseThemeTinted, layer)

	switch cmd {
	case Cmd_BaseThemeTinted:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BaseThemeTinted))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid BaseThemeTinted encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBaseThemeArctic) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBaseThemeArctic, layer)

	switch cmd {
	case Cmd_BaseThemeArctic:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BaseThemeArctic))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid BaseThemeArctic encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBindAuthKeyInner) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBindAuthKeyInner, layer)

	switch cmd {
	case Cmd_BindAuthKeyInner:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BindAuthKeyInner))
		x.Long(m.GetNonce())
		x.Long(m.GetTempAuthKeyId())
		x.Long(m.GetPermAuthKeyId())
		x.Long(m.GetTempSessionId())
		x.Int(m.GetExpiresAt())

		return x.GetBuf()

	default:
		log.Errorf("invalid BindAuthKeyInner encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBoolFalse) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBoolFalse, layer)

	switch cmd {
	case Cmd_BoolFalse:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BoolFalse))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid BoolFalse encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBoolTrue) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBoolTrue, layer)

	switch cmd {
	case Cmd_BoolTrue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BoolTrue))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid BoolTrue encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotCommand) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBotCommand, layer)

	switch cmd {
	case Cmd_BotCommand:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotCommand))
		x.String(m.GetCommand())
		x.String(m.GetDescription())

		return x.GetBuf()

	default:
		log.Errorf("invalid BotCommand encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotInfo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBotInfo, layer)

	switch cmd {
	case Cmd_BotInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInfo))
		x.Int(m.GetUserId())
		x.String(m.GetDescription())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetCommands())))
		for _, v := range m.GetCommands() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid BotInfo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotInlineMessageMediaAuto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBotInlineMessageMediaAuto, layer)

	switch cmd {
	case Cmd_BotInlineMessageMediaAuto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineMessageMediaAuto))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetCaption())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_BotInlineMessageMediaAuto764cf810:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineMessageMediaAuto764cf810))
		// flags
		var flags int32 = 0
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 1
		}
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetMessage())
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid BotInlineMessageMediaAuto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotInlineMessageText) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBotInlineMessageText, layer)

	switch cmd {
	case Cmd_BotInlineMessageText:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineMessageText))
		// flags
		var flags int32 = 0
		if m.GetNoWebpage() == true {
			flags |= 1 << 0
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 1
		}
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetMessage())
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid BotInlineMessageText encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotInlineMessageMediaGeo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBotInlineMessageMediaGeo, layer)

	switch cmd {
	case Cmd_BotInlineMessageMediaGeo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineMessageMediaGeo))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetGeo().Encode(layer))
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_BotInlineMessageMediaGeo51846fd:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineMessageMediaGeo51846fd))
		// flags
		var flags int32 = 0
		if m.GetHeading() != 0 {
			flags |= 1 << 0
		}
		if m.GetPeriod() != 0 {
			flags |= 1 << 1
		}
		if m.GetProximityNotificationRadius() != 0 {
			flags |= 1 << 3
		}
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetGeo().Encode(layer))
		if m.GetHeading() != 0 {
			x.Int(m.GetHeading())
		}
		if m.GetPeriod() != 0 {
			x.Int(m.GetPeriod())
		}
		if m.GetProximityNotificationRadius() != 0 {
			x.Int(m.GetProximityNotificationRadius())
		}
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_BotInlineMessageMediaGeob722de65:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineMessageMediaGeob722de65))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetGeo().Encode(layer))
		x.Int(m.GetPeriod())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid BotInlineMessageMediaGeo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotInlineMessageMediaVenue) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBotInlineMessageMediaVenue, layer)

	switch cmd {
	case Cmd_BotInlineMessageMediaVenue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineMessageMediaVenue))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetGeo().Encode(layer))
		x.String(m.GetTitle())
		x.String(m.GetAddress())
		x.String(m.GetProvider())
		x.String(m.GetVenueId())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_BotInlineMessageMediaVenue8a86659c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineMessageMediaVenue8a86659c))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetGeo().Encode(layer))
		x.String(m.GetTitle())
		x.String(m.GetAddress())
		x.String(m.GetProvider())
		x.String(m.GetVenueId())
		x.String(m.GetVenueType())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid BotInlineMessageMediaVenue encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotInlineMessageMediaContact) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBotInlineMessageMediaContact, layer)

	switch cmd {
	case Cmd_BotInlineMessageMediaContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineMessageMediaContact))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetPhoneNumber())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_BotInlineMessageMediaContact18d1cdc2:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineMessageMediaContact18d1cdc2))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetPhoneNumber())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())
		x.String(m.GetVcard())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid BotInlineMessageMediaContact encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotInlineResult) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBotInlineResult, layer)

	switch cmd {
	case Cmd_BotInlineResult:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineResult))
		// flags
		var flags int32 = 0
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetUrl()) > 0 {
			flags |= 1 << 3
		}
		if len(m.GetThumbUrl()) > 0 {
			flags |= 1 << 4
		}
		if len(m.GetContentUrl()) > 0 {
			flags |= 1 << 5
		}
		if len(m.GetContentType()) > 0 {
			flags |= 1 << 5
		}
		if m.GetW() != 0 {
			flags |= 1 << 6
		}
		if m.GetH() != 0 {
			flags |= 1 << 6
		}
		if m.GetDuration() != 0 {
			flags |= 1 << 7
		}
		x.Int(flags)

		x.String(m.GetId())
		x.String(m.GetType())
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		if len(m.GetUrl()) > 0 {
			x.String(m.GetUrl())
		}
		if len(m.GetThumbUrl()) > 0 {
			x.String(m.GetThumbUrl())
		}
		if len(m.GetContentUrl()) > 0 {
			x.String(m.GetContentUrl())
		}
		if len(m.GetContentType()) > 0 {
			x.String(m.GetContentType())
		}
		if m.GetW() != 0 {
			x.Int(m.GetW())
		}
		if m.GetH() != 0 {
			x.Int(m.GetH())
		}
		if m.GetDuration() != 0 {
			x.Int(m.GetDuration())
		}
		x.Bytes(m.GetSendMessage().Encode(layer))

		return x.GetBuf()
	case Cmd_BotInlineResult11965f3a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineResult11965f3a))
		// flags
		var flags int32 = 0
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetUrl()) > 0 {
			flags |= 1 << 3
		}
		if m.GetThumb() != nil {
			flags |= 1 << 4
		}
		if m.GetContent() != nil {
			flags |= 1 << 5
		}
		x.Int(flags)

		x.String(m.GetId())
		x.String(m.GetType())
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		if len(m.GetUrl()) > 0 {
			x.String(m.GetUrl())
		}
		if m.GetThumb() != nil {
			x.Bytes(m.GetThumb().Encode(layer))
		}
		if m.GetContent() != nil {
			x.Bytes(m.GetContent().Encode(layer))
		}
		x.Bytes(m.GetSendMessage().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid BotInlineResult encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLBotInlineMediaResult) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassBotInlineMediaResult, layer)

	switch cmd {
	case Cmd_BotInlineMediaResult:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_BotInlineMediaResult))
		// flags
		var flags int32 = 0
		if m.GetPhoto() != nil {
			flags |= 1 << 0
		}
		if m.GetDocument() != nil {
			flags |= 1 << 1
		}
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.String(m.GetId())
		x.String(m.GetType())
		if m.GetPhoto() != nil {
			x.Bytes(m.GetPhoto().Encode(layer))
		}
		if m.GetDocument() != nil {
			x.Bytes(m.GetDocument().Encode(layer))
		}
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		x.Bytes(m.GetSendMessage().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid BotInlineMediaResult encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLCdnConfig) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassCdnConfig, layer)

	switch cmd {
	case Cmd_CdnConfig:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_CdnConfig))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPublicKeys())))
		for _, v := range m.GetPublicKeys() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid CdnConfig encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLCdnFileHash) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassCdnFileHash, layer)

	switch cmd {
	case Cmd_CdnFileHash:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_CdnFileHash))
		x.Int(m.GetOffset())
		x.Int(m.GetLimit())
		x.StringBytes(m.GetHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid CdnFileHash encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLCdnPublicKey) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassCdnPublicKey, layer)

	switch cmd {
	case Cmd_CdnPublicKey:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_CdnPublicKey))
		x.Int(m.GetDcId())
		x.String(m.GetPublicKey())

		return x.GetBuf()

	default:
		log.Errorf("invalid CdnPublicKey encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEvent) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEvent, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEvent:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEvent))
		x.Long(m.GetId())
		x.Int(m.GetDate())
		x.Int(m.GetUserId())
		x.Bytes(m.GetAction().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEvent encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionChangeTitle) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionChangeTitle, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionChangeTitle:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionChangeTitle))
		x.String(m.GetPrevValueE6DFB82571())
		x.String(m.GetNewValueE6DFB82571())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionChangeTitle encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionChangeAbout) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionChangeAbout, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionChangeAbout:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionChangeAbout))
		x.String(m.GetPrevValueE6DFB82571())
		x.String(m.GetNewValueE6DFB82571())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionChangeAbout encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionChangeUsername) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionChangeUsername, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionChangeUsername:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionChangeUsername))
		x.String(m.GetPrevValueE6DFB82571())
		x.String(m.GetNewValueE6DFB82571())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionChangeUsername encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionChangePhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionChangePhoto, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionChangePhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionChangePhoto))
		x.Bytes(m.GetPrevPhotoB82F55C371().Encode(layer))
		x.Bytes(m.GetNewPhotoB82F55C371().Encode(layer))

		return x.GetBuf()
	case Cmd_ChannelAdminLogEventActionChangePhoto434bd2af:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionChangePhoto434bd2af))
		x.Bytes(m.GetPrevPhoto434BD2AF98().Encode(layer))
		x.Bytes(m.GetNewPhoto434BD2AF98().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionChangePhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionToggleInvites) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionToggleInvites, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionToggleInvites:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionToggleInvites))
		x.Bytes(m.GetNewValue1B7907AE71().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionToggleInvites encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionToggleSignatures) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionToggleSignatures, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionToggleSignatures:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionToggleSignatures))
		x.Bytes(m.GetNewValue1B7907AE71().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionToggleSignatures encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionUpdatePinned) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionUpdatePinned, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionUpdatePinned:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionUpdatePinned))
		x.Bytes(m.GetMessage().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionUpdatePinned encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionEditMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionEditMessage, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionEditMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionEditMessage))
		x.Bytes(m.GetPrevMessage().Encode(layer))
		x.Bytes(m.GetNewMessage().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionEditMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionDeleteMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionDeleteMessage, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionDeleteMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionDeleteMessage))
		x.Bytes(m.GetMessage().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionDeleteMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionParticipantJoin) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionParticipantJoin, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionParticipantJoin:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionParticipantJoin))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionParticipantJoin encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionParticipantLeave) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionParticipantLeave, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionParticipantLeave:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionParticipantLeave))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionParticipantLeave encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionParticipantInvite) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionParticipantInvite, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionParticipantInvite:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionParticipantInvite))
		x.Bytes(m.GetParticipantE31C34D871().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionParticipantInvite encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionParticipantToggleBan) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionParticipantToggleBan, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionParticipantToggleBan:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionParticipantToggleBan))
		x.Bytes(m.GetPrevParticipant().Encode(layer))
		x.Bytes(m.GetNewParticipant().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionParticipantToggleBan encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionParticipantToggleAdmin, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionParticipantToggleAdmin:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionParticipantToggleAdmin))
		x.Bytes(m.GetPrevParticipant().Encode(layer))
		x.Bytes(m.GetNewParticipant().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionParticipantToggleAdmin encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionChangeStickerSet) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionChangeStickerSet, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionChangeStickerSet:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionChangeStickerSet))
		x.Bytes(m.GetPrevStickerset().Encode(layer))
		x.Bytes(m.GetNewStickerset().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionChangeStickerSet encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionTogglePreHistoryHidden) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionTogglePreHistoryHidden, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionTogglePreHistoryHidden:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionTogglePreHistoryHidden))
		x.Bytes(m.GetNewValue1B7907AE71().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionTogglePreHistoryHidden encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionDefaultBannedRights) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionDefaultBannedRights, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionDefaultBannedRights:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionDefaultBannedRights))
		x.Bytes(m.GetPrevBannedRights().Encode(layer))
		x.Bytes(m.GetNewBannedRights().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionDefaultBannedRights encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionStopPoll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionStopPoll, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionStopPoll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionStopPoll))
		x.Bytes(m.GetMessage().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionStopPoll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionChangeLinkedChat) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionChangeLinkedChat, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionChangeLinkedChat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionChangeLinkedChat))
		x.Int(m.GetPrevValueA26F881B100())
		x.Int(m.GetNewValueA26F881B100())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionChangeLinkedChat encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionChangeLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionChangeLocation, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionChangeLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionChangeLocation))
		x.Bytes(m.GetPrevValueE6B76AE102().Encode(layer))
		x.Bytes(m.GetNewValueE6B76AE102().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionChangeLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionToggleSlowMode) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionToggleSlowMode, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionToggleSlowMode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionToggleSlowMode))
		x.Int(m.GetPrevValueA26F881B100())
		x.Int(m.GetNewValueA26F881B100())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionToggleSlowMode encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionStartGroupCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionStartGroupCall, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionStartGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionStartGroupCall))
		x.Bytes(m.GetCall().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionStartGroupCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionDiscardGroupCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionDiscardGroupCall, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionDiscardGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionDiscardGroupCall))
		x.Bytes(m.GetCall().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionDiscardGroupCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionParticipantMute) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionParticipantMute, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionParticipantMute:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionParticipantMute))
		x.Bytes(m.GetParticipantF92424D2122().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionParticipantMute encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionParticipantUnmute) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionParticipantUnmute, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionParticipantUnmute:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionParticipantUnmute))
		x.Bytes(m.GetParticipantF92424D2122().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionParticipantUnmute encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventActionToggleGroupCallSetting) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventActionToggleGroupCallSetting, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventActionToggleGroupCallSetting:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventActionToggleGroupCallSetting))
		x.Bytes(m.GetJoinMuted().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventActionToggleGroupCallSetting encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminLogEventsFilter) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminLogEventsFilter, layer)

	switch cmd {
	case Cmd_ChannelAdminLogEventsFilter:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminLogEventsFilter))
		// flags
		var flags int32 = 0
		if m.GetJoin() == true {
			flags |= 1 << 0
		}
		if m.GetLeave() == true {
			flags |= 1 << 1
		}
		if m.GetInvite() == true {
			flags |= 1 << 2
		}
		if m.GetBan() == true {
			flags |= 1 << 3
		}
		if m.GetUnban() == true {
			flags |= 1 << 4
		}
		if m.GetKick() == true {
			flags |= 1 << 5
		}
		if m.GetUnkick() == true {
			flags |= 1 << 6
		}
		if m.GetPromote() == true {
			flags |= 1 << 7
		}
		if m.GetDemote() == true {
			flags |= 1 << 8
		}
		if m.GetInfo() == true {
			flags |= 1 << 9
		}
		if m.GetSettings() == true {
			flags |= 1 << 10
		}
		if m.GetPinned() == true {
			flags |= 1 << 11
		}
		if m.GetEdit() == true {
			flags |= 1 << 12
		}
		if m.GetDelete() == true {
			flags |= 1 << 13
		}
		if m.GetGroupCall() == true {
			flags |= 1 << 14
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminLogEventsFilter encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelAdminRights) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelAdminRights, layer)

	switch cmd {
	case Cmd_ChannelAdminRights:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelAdminRights))
		// flags
		var flags int32 = 0
		if m.GetChangeInfo() == true {
			flags |= 1 << 0
		}
		if m.GetPostMessages() == true {
			flags |= 1 << 1
		}
		if m.GetEditMessages() == true {
			flags |= 1 << 2
		}
		if m.GetDeleteMessages() == true {
			flags |= 1 << 3
		}
		if m.GetBanUsers() == true {
			flags |= 1 << 4
		}
		if m.GetInviteUsers() == true {
			flags |= 1 << 5
		}
		if m.GetInviteLink() == true {
			flags |= 1 << 6
		}
		if m.GetPinMessages() == true {
			flags |= 1 << 7
		}
		if m.GetAddAdmins() == true {
			flags |= 1 << 9
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelAdminRights encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelBannedRights) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelBannedRights, layer)

	switch cmd {
	case Cmd_ChannelBannedRights:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelBannedRights))
		// flags
		var flags int32 = 0
		if m.GetViewMessages() == true {
			flags |= 1 << 0
		}
		if m.GetSendMessages() == true {
			flags |= 1 << 1
		}
		if m.GetSendMedia() == true {
			flags |= 1 << 2
		}
		if m.GetSendStickers() == true {
			flags |= 1 << 3
		}
		if m.GetSendGifs() == true {
			flags |= 1 << 4
		}
		if m.GetSendGames() == true {
			flags |= 1 << 5
		}
		if m.GetSendInline() == true {
			flags |= 1 << 6
		}
		if m.GetEmbedLinks() == true {
			flags |= 1 << 7
		}
		x.Int(flags)

		x.Int(m.GetUntilDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelBannedRights encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelLocationEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelLocationEmpty, layer)

	switch cmd {
	case Cmd_ChannelLocationEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelLocationEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelLocationEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelLocation, layer)

	switch cmd {
	case Cmd_ChannelLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelLocation))
		x.Bytes(m.GetGeoPoint().Encode(layer))
		x.String(m.GetAddress())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelMessagesFilterEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelMessagesFilterEmpty, layer)

	switch cmd {
	case Cmd_ChannelMessagesFilterEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelMessagesFilterEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelMessagesFilterEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelMessagesFilter) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelMessagesFilter, layer)

	switch cmd {
	case Cmd_ChannelMessagesFilter:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelMessagesFilter))
		// flags
		var flags int32 = 0
		if m.GetExcludeNewMessages() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetRanges())))
		for _, v := range m.GetRanges() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelMessagesFilter encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelMessagesFilterCollapsed) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelMessagesFilterCollapsed, layer)

	switch cmd {
	case Cmd_ChannelMessagesFilterCollapsed:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelMessagesFilterCollapsed))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelMessagesFilterCollapsed encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipant) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipant, layer)

	switch cmd {
	case Cmd_ChannelParticipant:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipant))
		x.Int(m.GetUserId())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipant encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantSelf) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantSelf, layer)

	switch cmd {
	case Cmd_ChannelParticipantSelf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantSelf))
		x.Int(m.GetUserId())
		x.Int(m.GetInviterId())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantSelf encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantCreator) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantCreator, layer)

	switch cmd {
	case Cmd_ChannelParticipantCreator:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantCreator))
		x.Int(m.GetUserId())

		return x.GetBuf()
	case Cmd_ChannelParticipantCreator447dca4b:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantCreator447dca4b))
		// flags
		var flags int32 = 0
		if len(m.GetRank()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetUserId())
		x.Bytes(m.GetAdminRights5DAA6E2393().Encode(layer))
		if len(m.GetRank()) > 0 {
			x.String(m.GetRank())
		}

		return x.GetBuf()
	case Cmd_ChannelParticipantCreator808d15a4:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantCreator808d15a4))
		// flags
		var flags int32 = 0
		if len(m.GetRank()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetUserId())
		if len(m.GetRank()) > 0 {
			x.String(m.GetRank())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantCreator encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantAdmin) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantAdmin, layer)

	switch cmd {
	case Cmd_ChannelParticipantAdmin:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantAdmin))
		// flags
		var flags int32 = 0
		if m.GetCanEdit() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetUserId())
		x.Int(m.GetInviterId())
		x.Int(m.GetPromotedBy())
		x.Int(m.GetDate())
		x.Bytes(m.GetAdminRightsA82FA89871().Encode(layer))

		return x.GetBuf()
	case Cmd_ChannelParticipantAdmin5daa6e23:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantAdmin5daa6e23))
		// flags
		var flags int32 = 0
		if m.GetCanEdit() == true {
			flags |= 1 << 0
		}
		if m.GetSelf() == true {
			flags |= 1 << 1
		}
		if m.GetInviterId() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetUserId())
		if m.GetInviterId() != 0 {
			x.Int(m.GetInviterId())
		}
		x.Int(m.GetPromotedBy())
		x.Int(m.GetDate())
		x.Bytes(m.GetAdminRights5DAA6E2393().Encode(layer))

		return x.GetBuf()
	case Cmd_ChannelParticipantAdminccbebbaf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantAdminccbebbaf))
		// flags
		var flags int32 = 0
		if m.GetCanEdit() == true {
			flags |= 1 << 0
		}
		if m.GetSelf() == true {
			flags |= 1 << 1
		}
		if m.GetInviterId() != 0 {
			flags |= 1 << 1
		}
		if len(m.GetRank()) > 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.GetUserId())
		if m.GetInviterId() != 0 {
			x.Int(m.GetInviterId())
		}
		x.Int(m.GetPromotedBy())
		x.Int(m.GetDate())
		x.Bytes(m.GetAdminRights5DAA6E2393().Encode(layer))
		if len(m.GetRank()) > 0 {
			x.String(m.GetRank())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantAdmin encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantBanned) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantBanned, layer)

	switch cmd {
	case Cmd_ChannelParticipantBanned:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantBanned))
		// flags
		var flags int32 = 0
		if m.GetLeft() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetUserId())
		x.Int(m.GetKickedBy())
		x.Int(m.GetDate())
		x.Bytes(m.GetBannedRights222C188671().Encode(layer))

		return x.GetBuf()
	case Cmd_ChannelParticipantBanned1c0facaf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantBanned1c0facaf))
		// flags
		var flags int32 = 0
		if m.GetLeft() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetUserId())
		x.Int(m.GetKickedBy())
		x.Int(m.GetDate())
		x.Bytes(m.GetBannedRights1C0FACAF93().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantBanned encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantModerator) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantModerator, layer)

	switch cmd {
	case Cmd_ChannelParticipantModerator:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantModerator))
		x.Int(m.GetUserId())
		x.Int(m.GetInviterId())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantModerator encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantEditor) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantEditor, layer)

	switch cmd {
	case Cmd_ChannelParticipantEditor:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantEditor))
		x.Int(m.GetUserId())
		x.Int(m.GetInviterId())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantEditor encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantKicked) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantKicked, layer)

	switch cmd {
	case Cmd_ChannelParticipantKicked:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantKicked))
		x.Int(m.GetUserId())
		x.Int(m.GetKickedBy())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantKicked encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantLeft) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantLeft, layer)

	switch cmd {
	case Cmd_ChannelParticipantLeft:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantLeft))
		x.Int(m.GetUserId())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantLeft encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelRoleEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelRoleEmpty, layer)

	switch cmd {
	case Cmd_ChannelRoleEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelRoleEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelRoleEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelRoleModerator) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelRoleModerator, layer)

	switch cmd {
	case Cmd_ChannelRoleModerator:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelRoleModerator))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelRoleModerator encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelRoleEditor) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelRoleEditor, layer)

	switch cmd {
	case Cmd_ChannelRoleEditor:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelRoleEditor))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelRoleEditor encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantsRecent) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantsRecent, layer)

	switch cmd {
	case Cmd_ChannelParticipantsRecent:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantsRecent))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantsRecent encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantsAdmins) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantsAdmins, layer)

	switch cmd {
	case Cmd_ChannelParticipantsAdmins:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantsAdmins))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantsAdmins encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantsKicked) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantsKicked, layer)

	switch cmd {
	case Cmd_ChannelParticipantsKicked:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantsKicked))
		x.String(m.GetQ())

		return x.GetBuf()
	case Cmd_ChannelParticipantsKicked3c37bb7a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantsKicked3c37bb7a))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantsKicked encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantsBots) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantsBots, layer)

	switch cmd {
	case Cmd_ChannelParticipantsBots:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantsBots))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantsBots encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantsBanned) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantsBanned, layer)

	switch cmd {
	case Cmd_ChannelParticipantsBanned:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantsBanned))
		x.String(m.GetQ())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantsBanned encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantsSearch) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantsSearch, layer)

	switch cmd {
	case Cmd_ChannelParticipantsSearch:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantsSearch))
		x.String(m.GetQ())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantsSearch encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantsContacts) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantsContacts, layer)

	switch cmd {
	case Cmd_ChannelParticipantsContacts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantsContacts))
		x.String(m.GetQ())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantsContacts encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelParticipantsMentions) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelParticipantsMentions, layer)

	switch cmd {
	case Cmd_ChannelParticipantsMentions:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelParticipantsMentions))
		// flags
		var flags int32 = 0
		if len(m.GetQ()) > 0 {
			flags |= 1 << 0
		}
		if m.GetTopMsgId() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		if len(m.GetQ()) > 0 {
			x.String(m.GetQ())
		}
		if m.GetTopMsgId() != 0 {
			x.Int(m.GetTopMsgId())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelParticipantsMentions encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsAdminLogResults) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelsAdminLogResults, layer)

	switch cmd {
	case Cmd_ChannelsAdminLogResults:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsAdminLogResults))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetEvents())))
		for _, v := range m.GetEvents() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelsAdminLogResults encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsChannelParticipant) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelsChannelParticipant, layer)

	switch cmd {
	case Cmd_ChannelsChannelParticipant:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsChannelParticipant))
		x.Bytes(m.GetParticipant().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelsChannelParticipant encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsChannelParticipants) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelsChannelParticipants, layer)

	switch cmd {
	case Cmd_ChannelsChannelParticipants:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsChannelParticipants))
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetParticipants())))
		for _, v := range m.GetParticipants() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelsChannelParticipants encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelsChannelParticipantsNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelsChannelParticipantsNotModified, layer)

	switch cmd {
	case Cmd_ChannelsChannelParticipantsNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelsChannelParticipantsNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelsChannelParticipantsNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatEmpty, layer)

	switch cmd {
	case Cmd_ChatEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatEmpty))
		x.Int(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChat) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChat, layer)

	switch cmd {
	case Cmd_Chat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Chat))
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetKicked() == true {
			flags |= 1 << 1
		}
		if m.GetLeft() == true {
			flags |= 1 << 2
		}
		if m.GetAdminsEnabled() == true {
			flags |= 1 << 3
		}
		if m.GetAdmin() == true {
			flags |= 1 << 4
		}
		if m.GetDeactivated() == true {
			flags |= 1 << 5
		}
		if m.GetMigratedTo() != nil {
			flags |= 1 << 6
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetTitle())
		x.Bytes(m.GetPhoto().Encode(layer))
		x.Int(m.GetParticipantsCount())
		x.Int(m.GetDate())
		x.Int(m.GetVersion())
		if m.GetMigratedTo() != nil {
			x.Bytes(m.GetMigratedTo().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_Chat3bda1bde:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Chat3bda1bde))
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetKicked() == true {
			flags |= 1 << 1
		}
		if m.GetLeft() == true {
			flags |= 1 << 2
		}
		if m.GetDeactivated() == true {
			flags |= 1 << 5
		}
		if m.GetCallActive() == true {
			flags |= 1 << 23
		}
		if m.GetCallNotEmpty() == true {
			flags |= 1 << 24
		}
		if m.GetMigratedTo() != nil {
			flags |= 1 << 6
		}
		if m.GetAdminRights4DF3083493() != nil {
			flags |= 1 << 14
		}
		if m.GetDefaultBannedRights() != nil {
			flags |= 1 << 18
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetTitle())
		x.Bytes(m.GetPhoto().Encode(layer))
		x.Int(m.GetParticipantsCount())
		x.Int(m.GetDate())
		x.Int(m.GetVersion())
		if m.GetMigratedTo() != nil {
			x.Bytes(m.GetMigratedTo().Encode(layer))
		}
		if m.GetAdminRights4DF3083493() != nil {
			x.Bytes(m.GetAdminRights4DF3083493().Encode(layer))
		}
		if m.GetDefaultBannedRights() != nil {
			x.Bytes(m.GetDefaultBannedRights().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid Chat encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatForbidden) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatForbidden, layer)

	switch cmd {
	case Cmd_ChatForbidden:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatForbidden))
		x.Int(m.GetId())
		x.String(m.GetTitle())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatForbidden encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannel) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannel, layer)

	switch cmd {
	case Cmd_Channel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Channel))
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetLeft() == true {
			flags |= 1 << 2
		}
		if m.GetEditor() == true {
			flags |= 1 << 3
		}
		if m.GetBroadcast() == true {
			flags |= 1 << 5
		}
		if m.GetVerified() == true {
			flags |= 1 << 7
		}
		if m.GetMegagroup() == true {
			flags |= 1 << 8
		}
		if m.GetRestricted() == true {
			flags |= 1 << 9
		}
		if m.GetDemocracy() == true {
			flags |= 1 << 10
		}
		if m.GetSignatures() == true {
			flags |= 1 << 11
		}
		if m.GetMin() == true {
			flags |= 1 << 12
		}
		if m.GetAccessHash() != 0 {
			flags |= 1 << 13
		}
		if len(m.GetUsername()) > 0 {
			flags |= 1 << 6
		}
		if len(m.GetRestrictionReasonCB44B1C71()) > 0 {
			flags |= 1 << 9
		}
		if m.GetAdminRightsCB44B1C71() != nil {
			flags |= 1 << 14
		}
		if m.GetBannedRightsCB44B1C71() != nil {
			flags |= 1 << 15
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetAccessHash() != 0 {
			x.Long(m.GetAccessHash())
		}
		x.String(m.GetTitle())
		if len(m.GetUsername()) > 0 {
			x.String(m.GetUsername())
		}
		x.Bytes(m.GetPhoto().Encode(layer))
		x.Int(m.GetDate())
		x.Int(m.GetVersion())
		if len(m.GetRestrictionReasonCB44B1C71()) > 0 {
			x.String(m.GetRestrictionReasonCB44B1C71())
		}
		if m.GetAdminRightsCB44B1C71() != nil {
			x.Bytes(m.GetAdminRightsCB44B1C71().Encode(layer))
		}
		if m.GetBannedRightsCB44B1C71() != nil {
			x.Bytes(m.GetBannedRightsCB44B1C71().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_Channel4df30834:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Channel4df30834))
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetLeft() == true {
			flags |= 1 << 2
		}
		if m.GetBroadcast() == true {
			flags |= 1 << 5
		}
		if m.GetVerified() == true {
			flags |= 1 << 7
		}
		if m.GetMegagroup() == true {
			flags |= 1 << 8
		}
		if m.GetRestricted() == true {
			flags |= 1 << 9
		}
		if m.GetSignatures() == true {
			flags |= 1 << 11
		}
		if m.GetMin() == true {
			flags |= 1 << 12
		}
		if m.GetAccessHash() != 0 {
			flags |= 1 << 13
		}
		if len(m.GetUsername()) > 0 {
			flags |= 1 << 6
		}
		if len(m.GetRestrictionReasonCB44B1C71()) > 0 {
			flags |= 1 << 9
		}
		if m.GetAdminRights4DF3083493() != nil {
			flags |= 1 << 14
		}
		if m.GetBannedRights4DF3083493() != nil {
			flags |= 1 << 15
		}
		if m.GetDefaultBannedRights() != nil {
			flags |= 1 << 18
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 17
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetAccessHash() != 0 {
			x.Long(m.GetAccessHash())
		}
		x.String(m.GetTitle())
		if len(m.GetUsername()) > 0 {
			x.String(m.GetUsername())
		}
		x.Bytes(m.GetPhoto().Encode(layer))
		x.Int(m.GetDate())
		x.Int(m.GetVersion())
		if len(m.GetRestrictionReasonCB44B1C71()) > 0 {
			x.String(m.GetRestrictionReasonCB44B1C71())
		}
		if m.GetAdminRights4DF3083493() != nil {
			x.Bytes(m.GetAdminRights4DF3083493().Encode(layer))
		}
		if m.GetBannedRights4DF3083493() != nil {
			x.Bytes(m.GetBannedRights4DF3083493().Encode(layer))
		}
		if m.GetDefaultBannedRights() != nil {
			x.Bytes(m.GetDefaultBannedRights().Encode(layer))
		}
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}

		return x.GetBuf()
	case Cmd_Channela14dca52:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Channela14dca52))
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetKicked() == true {
			flags |= 1 << 1
		}
		if m.GetLeft() == true {
			flags |= 1 << 2
		}
		if m.GetEditor() == true {
			flags |= 1 << 3
		}
		if m.GetModerator() == true {
			flags |= 1 << 4
		}
		if m.GetBroadcast() == true {
			flags |= 1 << 5
		}
		if m.GetVerified() == true {
			flags |= 1 << 7
		}
		if m.GetMegagroup() == true {
			flags |= 1 << 8
		}
		if m.GetRestricted() == true {
			flags |= 1 << 9
		}
		if m.GetDemocracy() == true {
			flags |= 1 << 10
		}
		if m.GetSignatures() == true {
			flags |= 1 << 11
		}
		if m.GetMin() == true {
			flags |= 1 << 12
		}
		if m.GetAccessHash() != 0 {
			flags |= 1 << 13
		}
		if len(m.GetUsername()) > 0 {
			flags |= 1 << 6
		}
		if len(m.GetRestrictionReasonCB44B1C71()) > 0 {
			flags |= 1 << 9
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetAccessHash() != 0 {
			x.Long(m.GetAccessHash())
		}
		x.String(m.GetTitle())
		if len(m.GetUsername()) > 0 {
			x.String(m.GetUsername())
		}
		x.Bytes(m.GetPhoto().Encode(layer))
		x.Int(m.GetDate())
		x.Int(m.GetVersion())
		if len(m.GetRestrictionReasonCB44B1C71()) > 0 {
			x.String(m.GetRestrictionReasonCB44B1C71())
		}

		return x.GetBuf()
	case Cmd_Channelc88974ac:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Channelc88974ac))
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetLeft() == true {
			flags |= 1 << 2
		}
		if m.GetEditor() == true {
			flags |= 1 << 3
		}
		if m.GetBroadcast() == true {
			flags |= 1 << 5
		}
		if m.GetVerified() == true {
			flags |= 1 << 7
		}
		if m.GetMegagroup() == true {
			flags |= 1 << 8
		}
		if m.GetRestricted() == true {
			flags |= 1 << 9
		}
		if m.GetDemocracy() == true {
			flags |= 1 << 10
		}
		if m.GetSignatures() == true {
			flags |= 1 << 11
		}
		if m.GetMin() == true {
			flags |= 1 << 12
		}
		if m.GetAccessHash() != 0 {
			flags |= 1 << 13
		}
		if len(m.GetUsername()) > 0 {
			flags |= 1 << 6
		}
		if len(m.GetRestrictionReasonCB44B1C71()) > 0 {
			flags |= 1 << 9
		}
		if m.GetAdminRightsCB44B1C71() != nil {
			flags |= 1 << 14
		}
		if m.GetBannedRightsCB44B1C71() != nil {
			flags |= 1 << 15
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 17
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetAccessHash() != 0 {
			x.Long(m.GetAccessHash())
		}
		x.String(m.GetTitle())
		if len(m.GetUsername()) > 0 {
			x.String(m.GetUsername())
		}
		x.Bytes(m.GetPhoto().Encode(layer))
		x.Int(m.GetDate())
		x.Int(m.GetVersion())
		if len(m.GetRestrictionReasonCB44B1C71()) > 0 {
			x.String(m.GetRestrictionReasonCB44B1C71())
		}
		if m.GetAdminRightsCB44B1C71() != nil {
			x.Bytes(m.GetAdminRightsCB44B1C71().Encode(layer))
		}
		if m.GetBannedRightsCB44B1C71() != nil {
			x.Bytes(m.GetBannedRightsCB44B1C71().Encode(layer))
		}
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}

		return x.GetBuf()
	case Cmd_Channeld31a961e:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Channeld31a961e))
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetLeft() == true {
			flags |= 1 << 2
		}
		if m.GetBroadcast() == true {
			flags |= 1 << 5
		}
		if m.GetVerified() == true {
			flags |= 1 << 7
		}
		if m.GetMegagroup() == true {
			flags |= 1 << 8
		}
		if m.GetRestricted() == true {
			flags |= 1 << 9
		}
		if m.GetSignatures() == true {
			flags |= 1 << 11
		}
		if m.GetMin() == true {
			flags |= 1 << 12
		}
		if m.GetScam() == true {
			flags |= 1 << 19
		}
		if m.GetHasLink() == true {
			flags |= 1 << 20
		}
		if m.GetHasGeo() == true {
			flags |= 1 << 21
		}
		if m.GetSlowmodeEnabled() == true {
			flags |= 1 << 22
		}
		if m.GetCallActive() == true {
			flags |= 1 << 23
		}
		if m.GetCallNotEmpty() == true {
			flags |= 1 << 24
		}
		if m.GetAccessHash() != 0 {
			flags |= 1 << 13
		}
		if len(m.GetUsername()) > 0 {
			flags |= 1 << 6
		}
		if len(m.GetRestrictionReasonD31A961E122()) > 0 {
			flags |= 1 << 9
		}
		if m.GetAdminRights4DF3083493() != nil {
			flags |= 1 << 14
		}
		if m.GetBannedRights4DF3083493() != nil {
			flags |= 1 << 15
		}
		if m.GetDefaultBannedRights() != nil {
			flags |= 1 << 18
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 17
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetAccessHash() != 0 {
			x.Long(m.GetAccessHash())
		}
		x.String(m.GetTitle())
		if len(m.GetUsername()) > 0 {
			x.String(m.GetUsername())
		}
		x.Bytes(m.GetPhoto().Encode(layer))
		x.Int(m.GetDate())
		x.Int(m.GetVersion())
		if len(m.GetRestrictionReasonD31A961E122()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetRestrictionReasonD31A961E122())))
			for _, v := range m.GetRestrictionReasonD31A961E122() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetAdminRights4DF3083493() != nil {
			x.Bytes(m.GetAdminRights4DF3083493().Encode(layer))
		}
		if m.GetBannedRights4DF3083493() != nil {
			x.Bytes(m.GetBannedRights4DF3083493().Encode(layer))
		}
		if m.GetDefaultBannedRights() != nil {
			x.Bytes(m.GetDefaultBannedRights().Encode(layer))
		}
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid Channel encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelForbidden) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelForbidden, layer)

	switch cmd {
	case Cmd_ChannelForbidden:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelForbidden))
		// flags
		var flags int32 = 0
		if m.GetBroadcast() == true {
			flags |= 1 << 5
		}
		if m.GetMegagroup() == true {
			flags |= 1 << 8
		}
		if m.GetUntilDate() != 0 {
			flags |= 1 << 16
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.Long(m.GetAccessHash())
		x.String(m.GetTitle())
		if m.GetUntilDate() != 0 {
			x.Int(m.GetUntilDate())
		}

		return x.GetBuf()
	case Cmd_ChannelForbidden2d85832c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelForbidden2d85832c))
		x.Int(m.GetId())
		x.Long(m.GetAccessHash())
		x.String(m.GetTitle())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelForbidden encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatAdminRights) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatAdminRights, layer)

	switch cmd {
	case Cmd_ChatAdminRights:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatAdminRights))
		// flags
		var flags int32 = 0
		if m.GetChangeInfo() == true {
			flags |= 1 << 0
		}
		if m.GetPostMessages() == true {
			flags |= 1 << 1
		}
		if m.GetEditMessages() == true {
			flags |= 1 << 2
		}
		if m.GetDeleteMessages() == true {
			flags |= 1 << 3
		}
		if m.GetBanUsers() == true {
			flags |= 1 << 4
		}
		if m.GetInviteUsers() == true {
			flags |= 1 << 5
		}
		if m.GetPinMessages() == true {
			flags |= 1 << 7
		}
		if m.GetAddAdmins() == true {
			flags |= 1 << 9
		}
		if m.GetAnonymous() == true {
			flags |= 1 << 10
		}
		if m.GetManageCall() == true {
			flags |= 1 << 11
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatAdminRights encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatBannedRights) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatBannedRights, layer)

	switch cmd {
	case Cmd_ChatBannedRights:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatBannedRights))
		// flags
		var flags int32 = 0
		if m.GetViewMessages() == true {
			flags |= 1 << 0
		}
		if m.GetSendMessages() == true {
			flags |= 1 << 1
		}
		if m.GetSendMedia() == true {
			flags |= 1 << 2
		}
		if m.GetSendStickers() == true {
			flags |= 1 << 3
		}
		if m.GetSendGifs() == true {
			flags |= 1 << 4
		}
		if m.GetSendGames() == true {
			flags |= 1 << 5
		}
		if m.GetSendInline() == true {
			flags |= 1 << 6
		}
		if m.GetEmbedLinks() == true {
			flags |= 1 << 7
		}
		if m.GetSendPolls() == true {
			flags |= 1 << 8
		}
		if m.GetChangeInfo() == true {
			flags |= 1 << 10
		}
		if m.GetInviteUsers() == true {
			flags |= 1 << 15
		}
		if m.GetPinMessages() == true {
			flags |= 1 << 17
		}
		x.Int(flags)

		x.Int(m.GetUntilDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatBannedRights encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatFull) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatFull, layer)

	switch cmd {
	case Cmd_ChatFull:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatFull))
		x.Int(m.GetId())
		x.Bytes(m.GetParticipants().Encode(layer))
		x.Bytes(m.GetChatPhoto().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBotInfo())))
		for _, v := range m.GetBotInfo() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_ChatFull1b7c9db3:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatFull1b7c9db3))
		// flags
		var flags int32 = 0
		if m.GetCanSetUsername() == true {
			flags |= 1 << 7
		}
		if m.GetChatPhoto() != nil {
			flags |= 1 << 2
		}
		if len(m.GetBotInfo()) > 0 {
			flags |= 1 << 3
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 6
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 11
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		x.Bytes(m.GetParticipants().Encode(layer))
		if m.GetChatPhoto() != nil {
			x.Bytes(m.GetChatPhoto().Encode(layer))
		}
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		if len(m.GetBotInfo()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetBotInfo())))
			for _, v := range m.GetBotInfo() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}

		return x.GetBuf()
	case Cmd_ChatFull22a235da:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatFull22a235da))
		// flags
		var flags int32 = 0
		if m.GetCanSetUsername() == true {
			flags |= 1 << 7
		}
		if m.GetChatPhoto() != nil {
			flags |= 1 << 2
		}
		if len(m.GetBotInfo()) > 0 {
			flags |= 1 << 3
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 6
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		x.Bytes(m.GetParticipants().Encode(layer))
		if m.GetChatPhoto() != nil {
			x.Bytes(m.GetChatPhoto().Encode(layer))
		}
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		if len(m.GetBotInfo()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetBotInfo())))
			for _, v := range m.GetBotInfo() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}

		return x.GetBuf()
	case Cmd_ChatFulldc8c181:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatFulldc8c181))
		// flags
		var flags int32 = 0
		if m.GetCanSetUsername() == true {
			flags |= 1 << 7
		}
		if m.GetHasScheduled() == true {
			flags |= 1 << 8
		}
		if m.GetChatPhoto() != nil {
			flags |= 1 << 2
		}
		if len(m.GetBotInfo()) > 0 {
			flags |= 1 << 3
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 6
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 11
		}
		if m.GetCall() != nil {
			flags |= 1 << 12
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		x.Bytes(m.GetParticipants().Encode(layer))
		if m.GetChatPhoto() != nil {
			x.Bytes(m.GetChatPhoto().Encode(layer))
		}
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		if len(m.GetBotInfo()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetBotInfo())))
			for _, v := range m.GetBotInfo() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		if m.GetCall() != nil {
			x.Bytes(m.GetCall().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_ChatFulledd2a791:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatFulledd2a791))
		// flags
		var flags int32 = 0
		if m.GetChatPhoto() != nil {
			flags |= 1 << 2
		}
		if len(m.GetBotInfo()) > 0 {
			flags |= 1 << 3
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 6
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.Bytes(m.GetParticipants().Encode(layer))
		if m.GetChatPhoto() != nil {
			x.Bytes(m.GetChatPhoto().Encode(layer))
		}
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		if len(m.GetBotInfo()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetBotInfo())))
			for _, v := range m.GetBotInfo() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatFull encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChannelFull) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChannelFull, layer)

	switch cmd {
	case Cmd_ChannelFull:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelFull))
		// flags
		var flags int32 = 0
		if m.GetCanViewParticipants() == true {
			flags |= 1 << 3
		}
		if m.GetCanSetUsername() == true {
			flags |= 1 << 6
		}
		if m.GetCanSetStickers() == true {
			flags |= 1 << 7
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 0
		}
		if m.GetAdminsCount() != 0 {
			flags |= 1 << 1
		}
		if m.GetKickedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetBannedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetMigratedFromChatId() != 0 {
			flags |= 1 << 4
		}
		if m.GetMigratedFromMaxId() != 0 {
			flags |= 1 << 4
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 5
		}
		if m.GetStickerset() != nil {
			flags |= 1 << 8
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}
		if m.GetAdminsCount() != 0 {
			x.Int(m.GetAdminsCount())
		}
		if m.GetKickedCount() != 0 {
			x.Int(m.GetKickedCount())
		}
		if m.GetBannedCount() != 0 {
			x.Int(m.GetBannedCount())
		}
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Bytes(m.GetChatPhoto().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBotInfo())))
		for _, v := range m.GetBotInfo() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMigratedFromChatId() != 0 {
			x.Int(m.GetMigratedFromChatId())
		}
		if m.GetMigratedFromMaxId() != 0 {
			x.Int(m.GetMigratedFromMaxId())
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		if m.GetStickerset() != nil {
			x.Bytes(m.GetStickerset().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_ChannelFull10916653:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelFull10916653))
		// flags
		var flags int32 = 0
		if m.GetCanViewParticipants() == true {
			flags |= 1 << 3
		}
		if m.GetCanSetUsername() == true {
			flags |= 1 << 6
		}
		if m.GetCanSetStickers() == true {
			flags |= 1 << 7
		}
		if m.GetHiddenPrehistory() == true {
			flags |= 1 << 10
		}
		if m.GetCanViewStats() == true {
			flags |= 1 << 12
		}
		if m.GetCanSetLocation() == true {
			flags |= 1 << 16
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 0
		}
		if m.GetAdminsCount() != 0 {
			flags |= 1 << 1
		}
		if m.GetKickedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetBannedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetOnlineCount() != 0 {
			flags |= 1 << 13
		}
		if m.GetMigratedFromChatId() != 0 {
			flags |= 1 << 4
		}
		if m.GetMigratedFromMaxId() != 0 {
			flags |= 1 << 4
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 5
		}
		if m.GetStickerset() != nil {
			flags |= 1 << 8
		}
		if m.GetAvailableMinId() != 0 {
			flags |= 1 << 9
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 11
		}
		if m.GetLinkedChatId() != 0 {
			flags |= 1 << 14
		}
		if m.GetLocation() != nil {
			flags |= 1 << 15
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}
		if m.GetAdminsCount() != 0 {
			x.Int(m.GetAdminsCount())
		}
		if m.GetKickedCount() != 0 {
			x.Int(m.GetKickedCount())
		}
		if m.GetBannedCount() != 0 {
			x.Int(m.GetBannedCount())
		}
		if m.GetOnlineCount() != 0 {
			x.Int(m.GetOnlineCount())
		}
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Bytes(m.GetChatPhoto().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBotInfo())))
		for _, v := range m.GetBotInfo() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMigratedFromChatId() != 0 {
			x.Int(m.GetMigratedFromChatId())
		}
		if m.GetMigratedFromMaxId() != 0 {
			x.Int(m.GetMigratedFromMaxId())
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		if m.GetStickerset() != nil {
			x.Bytes(m.GetStickerset().Encode(layer))
		}
		if m.GetAvailableMinId() != 0 {
			x.Int(m.GetAvailableMinId())
		}
		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		if m.GetLinkedChatId() != 0 {
			x.Int(m.GetLinkedChatId())
		}
		if m.GetLocation() != nil {
			x.Bytes(m.GetLocation().Encode(layer))
		}
		x.Int(m.GetPts())

		return x.GetBuf()
	case Cmd_ChannelFull1c87a71a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelFull1c87a71a))
		// flags
		var flags int32 = 0
		if m.GetCanViewParticipants() == true {
			flags |= 1 << 3
		}
		if m.GetCanSetUsername() == true {
			flags |= 1 << 6
		}
		if m.GetCanSetStickers() == true {
			flags |= 1 << 7
		}
		if m.GetHiddenPrehistory() == true {
			flags |= 1 << 10
		}
		if m.GetCanViewStats() == true {
			flags |= 1 << 12
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 0
		}
		if m.GetAdminsCount() != 0 {
			flags |= 1 << 1
		}
		if m.GetKickedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetBannedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetOnlineCount() != 0 {
			flags |= 1 << 13
		}
		if m.GetMigratedFromChatId() != 0 {
			flags |= 1 << 4
		}
		if m.GetMigratedFromMaxId() != 0 {
			flags |= 1 << 4
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 5
		}
		if m.GetStickerset() != nil {
			flags |= 1 << 8
		}
		if m.GetAvailableMinId() != 0 {
			flags |= 1 << 9
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}
		if m.GetAdminsCount() != 0 {
			x.Int(m.GetAdminsCount())
		}
		if m.GetKickedCount() != 0 {
			x.Int(m.GetKickedCount())
		}
		if m.GetBannedCount() != 0 {
			x.Int(m.GetBannedCount())
		}
		if m.GetOnlineCount() != 0 {
			x.Int(m.GetOnlineCount())
		}
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Bytes(m.GetChatPhoto().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBotInfo())))
		for _, v := range m.GetBotInfo() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMigratedFromChatId() != 0 {
			x.Int(m.GetMigratedFromChatId())
		}
		if m.GetMigratedFromMaxId() != 0 {
			x.Int(m.GetMigratedFromMaxId())
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		if m.GetStickerset() != nil {
			x.Bytes(m.GetStickerset().Encode(layer))
		}
		if m.GetAvailableMinId() != 0 {
			x.Int(m.GetAvailableMinId())
		}

		return x.GetBuf()
	case Cmd_ChannelFull2d895c74:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelFull2d895c74))
		// flags
		var flags int32 = 0
		if m.GetCanViewParticipants() == true {
			flags |= 1 << 3
		}
		if m.GetCanSetUsername() == true {
			flags |= 1 << 6
		}
		if m.GetCanSetStickers() == true {
			flags |= 1 << 7
		}
		if m.GetHiddenPrehistory() == true {
			flags |= 1 << 10
		}
		if m.GetCanViewStats() == true {
			flags |= 1 << 12
		}
		if m.GetCanSetLocation() == true {
			flags |= 1 << 16
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 0
		}
		if m.GetAdminsCount() != 0 {
			flags |= 1 << 1
		}
		if m.GetKickedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetBannedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetOnlineCount() != 0 {
			flags |= 1 << 13
		}
		if m.GetMigratedFromChatId() != 0 {
			flags |= 1 << 4
		}
		if m.GetMigratedFromMaxId() != 0 {
			flags |= 1 << 4
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 5
		}
		if m.GetStickerset() != nil {
			flags |= 1 << 8
		}
		if m.GetAvailableMinId() != 0 {
			flags |= 1 << 9
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 11
		}
		if m.GetLinkedChatId() != 0 {
			flags |= 1 << 14
		}
		if m.GetLocation() != nil {
			flags |= 1 << 15
		}
		if m.GetSlowmodeSeconds() != 0 {
			flags |= 1 << 17
		}
		if m.GetSlowmodeNextSendDate() != 0 {
			flags |= 1 << 18
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}
		if m.GetAdminsCount() != 0 {
			x.Int(m.GetAdminsCount())
		}
		if m.GetKickedCount() != 0 {
			x.Int(m.GetKickedCount())
		}
		if m.GetBannedCount() != 0 {
			x.Int(m.GetBannedCount())
		}
		if m.GetOnlineCount() != 0 {
			x.Int(m.GetOnlineCount())
		}
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Bytes(m.GetChatPhoto().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBotInfo())))
		for _, v := range m.GetBotInfo() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMigratedFromChatId() != 0 {
			x.Int(m.GetMigratedFromChatId())
		}
		if m.GetMigratedFromMaxId() != 0 {
			x.Int(m.GetMigratedFromMaxId())
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		if m.GetStickerset() != nil {
			x.Bytes(m.GetStickerset().Encode(layer))
		}
		if m.GetAvailableMinId() != 0 {
			x.Int(m.GetAvailableMinId())
		}
		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		if m.GetLinkedChatId() != 0 {
			x.Int(m.GetLinkedChatId())
		}
		if m.GetLocation() != nil {
			x.Bytes(m.GetLocation().Encode(layer))
		}
		if m.GetSlowmodeSeconds() != 0 {
			x.Int(m.GetSlowmodeSeconds())
		}
		if m.GetSlowmodeNextSendDate() != 0 {
			x.Int(m.GetSlowmodeNextSendDate())
		}
		x.Int(m.GetPts())

		return x.GetBuf()
	case Cmd_ChannelFull3648977:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelFull3648977))
		// flags
		var flags int32 = 0
		if m.GetCanViewParticipants() == true {
			flags |= 1 << 3
		}
		if m.GetCanSetUsername() == true {
			flags |= 1 << 6
		}
		if m.GetCanSetStickers() == true {
			flags |= 1 << 7
		}
		if m.GetHiddenPrehistory() == true {
			flags |= 1 << 10
		}
		if m.GetCanViewStats() == true {
			flags |= 1 << 12
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 0
		}
		if m.GetAdminsCount() != 0 {
			flags |= 1 << 1
		}
		if m.GetKickedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetBannedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetOnlineCount() != 0 {
			flags |= 1 << 13
		}
		if m.GetMigratedFromChatId() != 0 {
			flags |= 1 << 4
		}
		if m.GetMigratedFromMaxId() != 0 {
			flags |= 1 << 4
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 5
		}
		if m.GetStickerset() != nil {
			flags |= 1 << 8
		}
		if m.GetAvailableMinId() != 0 {
			flags |= 1 << 9
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 11
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}
		if m.GetAdminsCount() != 0 {
			x.Int(m.GetAdminsCount())
		}
		if m.GetKickedCount() != 0 {
			x.Int(m.GetKickedCount())
		}
		if m.GetBannedCount() != 0 {
			x.Int(m.GetBannedCount())
		}
		if m.GetOnlineCount() != 0 {
			x.Int(m.GetOnlineCount())
		}
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Bytes(m.GetChatPhoto().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBotInfo())))
		for _, v := range m.GetBotInfo() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMigratedFromChatId() != 0 {
			x.Int(m.GetMigratedFromChatId())
		}
		if m.GetMigratedFromMaxId() != 0 {
			x.Int(m.GetMigratedFromMaxId())
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		if m.GetStickerset() != nil {
			x.Bytes(m.GetStickerset().Encode(layer))
		}
		if m.GetAvailableMinId() != 0 {
			x.Int(m.GetAvailableMinId())
		}
		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		x.Int(m.GetPts())

		return x.GetBuf()
	case Cmd_ChannelFull76af5481:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelFull76af5481))
		// flags
		var flags int32 = 0
		if m.GetCanViewParticipants() == true {
			flags |= 1 << 3
		}
		if m.GetCanSetUsername() == true {
			flags |= 1 << 6
		}
		if m.GetCanSetStickers() == true {
			flags |= 1 << 7
		}
		if m.GetHiddenPrehistory() == true {
			flags |= 1 << 10
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 0
		}
		if m.GetAdminsCount() != 0 {
			flags |= 1 << 1
		}
		if m.GetKickedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetBannedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetMigratedFromChatId() != 0 {
			flags |= 1 << 4
		}
		if m.GetMigratedFromMaxId() != 0 {
			flags |= 1 << 4
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 5
		}
		if m.GetStickerset() != nil {
			flags |= 1 << 8
		}
		if m.GetAvailableMinId() != 0 {
			flags |= 1 << 9
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}
		if m.GetAdminsCount() != 0 {
			x.Int(m.GetAdminsCount())
		}
		if m.GetKickedCount() != 0 {
			x.Int(m.GetKickedCount())
		}
		if m.GetBannedCount() != 0 {
			x.Int(m.GetBannedCount())
		}
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Bytes(m.GetChatPhoto().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBotInfo())))
		for _, v := range m.GetBotInfo() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMigratedFromChatId() != 0 {
			x.Int(m.GetMigratedFromChatId())
		}
		if m.GetMigratedFromMaxId() != 0 {
			x.Int(m.GetMigratedFromMaxId())
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		if m.GetStickerset() != nil {
			x.Bytes(m.GetStickerset().Encode(layer))
		}
		if m.GetAvailableMinId() != 0 {
			x.Int(m.GetAvailableMinId())
		}

		return x.GetBuf()
	case Cmd_ChannelFull97bee562:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelFull97bee562))
		// flags
		var flags int32 = 0
		if m.GetCanViewParticipants() == true {
			flags |= 1 << 3
		}
		if m.GetCanSetUsername() == true {
			flags |= 1 << 6
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 0
		}
		if m.GetAdminsCount() != 0 {
			flags |= 1 << 1
		}
		if m.GetKickedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetMigratedFromChatId() != 0 {
			flags |= 1 << 4
		}
		if m.GetMigratedFromMaxId() != 0 {
			flags |= 1 << 4
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 5
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}
		if m.GetAdminsCount() != 0 {
			x.Int(m.GetAdminsCount())
		}
		if m.GetKickedCount() != 0 {
			x.Int(m.GetKickedCount())
		}
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Int(m.GetUnreadImportantCount())
		x.Bytes(m.GetChatPhoto().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBotInfo())))
		for _, v := range m.GetBotInfo() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMigratedFromChatId() != 0 {
			x.Int(m.GetMigratedFromChatId())
		}
		if m.GetMigratedFromMaxId() != 0 {
			x.Int(m.GetMigratedFromMaxId())
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}

		return x.GetBuf()
	case Cmd_ChannelFull9882e516:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelFull9882e516))
		// flags
		var flags int32 = 0
		if m.GetCanViewParticipants() == true {
			flags |= 1 << 3
		}
		if m.GetCanSetUsername() == true {
			flags |= 1 << 6
		}
		if m.GetCanSetStickers() == true {
			flags |= 1 << 7
		}
		if m.GetHiddenPrehistory() == true {
			flags |= 1 << 10
		}
		if m.GetCanViewStats() == true {
			flags |= 1 << 12
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 0
		}
		if m.GetAdminsCount() != 0 {
			flags |= 1 << 1
		}
		if m.GetKickedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetBannedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetOnlineCount() != 0 {
			flags |= 1 << 13
		}
		if m.GetMigratedFromChatId() != 0 {
			flags |= 1 << 4
		}
		if m.GetMigratedFromMaxId() != 0 {
			flags |= 1 << 4
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 5
		}
		if m.GetStickerset() != nil {
			flags |= 1 << 8
		}
		if m.GetAvailableMinId() != 0 {
			flags |= 1 << 9
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 11
		}
		if m.GetLinkedChatId() != 0 {
			flags |= 1 << 13
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}
		if m.GetAdminsCount() != 0 {
			x.Int(m.GetAdminsCount())
		}
		if m.GetKickedCount() != 0 {
			x.Int(m.GetKickedCount())
		}
		if m.GetBannedCount() != 0 {
			x.Int(m.GetBannedCount())
		}
		if m.GetOnlineCount() != 0 {
			x.Int(m.GetOnlineCount())
		}
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Bytes(m.GetChatPhoto().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBotInfo())))
		for _, v := range m.GetBotInfo() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMigratedFromChatId() != 0 {
			x.Int(m.GetMigratedFromChatId())
		}
		if m.GetMigratedFromMaxId() != 0 {
			x.Int(m.GetMigratedFromMaxId())
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		if m.GetStickerset() != nil {
			x.Bytes(m.GetStickerset().Encode(layer))
		}
		if m.GetAvailableMinId() != 0 {
			x.Int(m.GetAvailableMinId())
		}
		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		if m.GetLinkedChatId() != 0 {
			x.Int(m.GetLinkedChatId())
		}
		x.Int(m.GetPts())

		return x.GetBuf()
	case Cmd_ChannelFullef3a6acd:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelFullef3a6acd))
		// flags
		var flags int32 = 0
		if m.GetCanViewParticipants() == true {
			flags |= 1 << 3
		}
		if m.GetCanSetUsername() == true {
			flags |= 1 << 6
		}
		if m.GetCanSetStickers() == true {
			flags |= 1 << 7
		}
		if m.GetHiddenPrehistory() == true {
			flags |= 1 << 10
		}
		if m.GetCanSetLocation() == true {
			flags |= 1 << 16
		}
		if m.GetHasScheduled() == true {
			flags |= 1 << 19
		}
		if m.GetCanViewStats() == true {
			flags |= 1 << 20
		}
		if m.GetBlocked() == true {
			flags |= 1 << 22
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 0
		}
		if m.GetAdminsCount() != 0 {
			flags |= 1 << 1
		}
		if m.GetKickedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetBannedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetOnlineCount() != 0 {
			flags |= 1 << 13
		}
		if m.GetMigratedFromChatId() != 0 {
			flags |= 1 << 4
		}
		if m.GetMigratedFromMaxId() != 0 {
			flags |= 1 << 4
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 5
		}
		if m.GetStickerset() != nil {
			flags |= 1 << 8
		}
		if m.GetAvailableMinId() != 0 {
			flags |= 1 << 9
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 11
		}
		if m.GetLinkedChatId() != 0 {
			flags |= 1 << 14
		}
		if m.GetLocation() != nil {
			flags |= 1 << 15
		}
		if m.GetSlowmodeSeconds() != 0 {
			flags |= 1 << 17
		}
		if m.GetSlowmodeNextSendDate() != 0 {
			flags |= 1 << 18
		}
		if m.GetStatsDc() != 0 {
			flags |= 1 << 12
		}
		if m.GetCall() != nil {
			flags |= 1 << 21
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}
		if m.GetAdminsCount() != 0 {
			x.Int(m.GetAdminsCount())
		}
		if m.GetKickedCount() != 0 {
			x.Int(m.GetKickedCount())
		}
		if m.GetBannedCount() != 0 {
			x.Int(m.GetBannedCount())
		}
		if m.GetOnlineCount() != 0 {
			x.Int(m.GetOnlineCount())
		}
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Bytes(m.GetChatPhoto().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBotInfo())))
		for _, v := range m.GetBotInfo() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMigratedFromChatId() != 0 {
			x.Int(m.GetMigratedFromChatId())
		}
		if m.GetMigratedFromMaxId() != 0 {
			x.Int(m.GetMigratedFromMaxId())
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		if m.GetStickerset() != nil {
			x.Bytes(m.GetStickerset().Encode(layer))
		}
		if m.GetAvailableMinId() != 0 {
			x.Int(m.GetAvailableMinId())
		}
		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		if m.GetLinkedChatId() != 0 {
			x.Int(m.GetLinkedChatId())
		}
		if m.GetLocation() != nil {
			x.Bytes(m.GetLocation().Encode(layer))
		}
		if m.GetSlowmodeSeconds() != 0 {
			x.Int(m.GetSlowmodeSeconds())
		}
		if m.GetSlowmodeNextSendDate() != 0 {
			x.Int(m.GetSlowmodeNextSendDate())
		}
		if m.GetStatsDc() != 0 {
			x.Int(m.GetStatsDc())
		}
		x.Int(m.GetPts())
		if m.GetCall() != nil {
			x.Bytes(m.GetCall().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_ChannelFullf0e6672a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChannelFullf0e6672a))
		// flags
		var flags int32 = 0
		if m.GetCanViewParticipants() == true {
			flags |= 1 << 3
		}
		if m.GetCanSetUsername() == true {
			flags |= 1 << 6
		}
		if m.GetCanSetStickers() == true {
			flags |= 1 << 7
		}
		if m.GetHiddenPrehistory() == true {
			flags |= 1 << 10
		}
		if m.GetCanSetLocation() == true {
			flags |= 1 << 16
		}
		if m.GetHasScheduled() == true {
			flags |= 1 << 19
		}
		if m.GetCanViewStats() == true {
			flags |= 1 << 20
		}
		if m.GetBlocked() == true {
			flags |= 1 << 22
		}
		if m.GetParticipantsCount() != 0 {
			flags |= 1 << 0
		}
		if m.GetAdminsCount() != 0 {
			flags |= 1 << 1
		}
		if m.GetKickedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetBannedCount() != 0 {
			flags |= 1 << 2
		}
		if m.GetOnlineCount() != 0 {
			flags |= 1 << 13
		}
		if m.GetMigratedFromChatId() != 0 {
			flags |= 1 << 4
		}
		if m.GetMigratedFromMaxId() != 0 {
			flags |= 1 << 4
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 5
		}
		if m.GetStickerset() != nil {
			flags |= 1 << 8
		}
		if m.GetAvailableMinId() != 0 {
			flags |= 1 << 9
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 11
		}
		if m.GetLinkedChatId() != 0 {
			flags |= 1 << 14
		}
		if m.GetLocation() != nil {
			flags |= 1 << 15
		}
		if m.GetSlowmodeSeconds() != 0 {
			flags |= 1 << 17
		}
		if m.GetSlowmodeNextSendDate() != 0 {
			flags |= 1 << 18
		}
		if m.GetStatsDc() != 0 {
			flags |= 1 << 12
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetAbout())
		if m.GetParticipantsCount() != 0 {
			x.Int(m.GetParticipantsCount())
		}
		if m.GetAdminsCount() != 0 {
			x.Int(m.GetAdminsCount())
		}
		if m.GetKickedCount() != 0 {
			x.Int(m.GetKickedCount())
		}
		if m.GetBannedCount() != 0 {
			x.Int(m.GetBannedCount())
		}
		if m.GetOnlineCount() != 0 {
			x.Int(m.GetOnlineCount())
		}
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Bytes(m.GetChatPhoto().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Bytes(m.GetExportedInvite().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBotInfo())))
		for _, v := range m.GetBotInfo() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMigratedFromChatId() != 0 {
			x.Int(m.GetMigratedFromChatId())
		}
		if m.GetMigratedFromMaxId() != 0 {
			x.Int(m.GetMigratedFromMaxId())
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		if m.GetStickerset() != nil {
			x.Bytes(m.GetStickerset().Encode(layer))
		}
		if m.GetAvailableMinId() != 0 {
			x.Int(m.GetAvailableMinId())
		}
		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		if m.GetLinkedChatId() != 0 {
			x.Int(m.GetLinkedChatId())
		}
		if m.GetLocation() != nil {
			x.Bytes(m.GetLocation().Encode(layer))
		}
		if m.GetSlowmodeSeconds() != 0 {
			x.Int(m.GetSlowmodeSeconds())
		}
		if m.GetSlowmodeNextSendDate() != 0 {
			x.Int(m.GetSlowmodeNextSendDate())
		}
		if m.GetStatsDc() != 0 {
			x.Int(m.GetStatsDc())
		}
		x.Int(m.GetPts())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChannelFull encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatInviteAlready) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatInviteAlready, layer)

	switch cmd {
	case Cmd_ChatInviteAlready:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatInviteAlready))
		x.Bytes(m.GetChat().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatInviteAlready encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatInvite) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatInvite, layer)

	switch cmd {
	case Cmd_ChatInvite:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatInvite))
		// flags
		var flags int32 = 0
		if m.GetChannel() == true {
			flags |= 1 << 0
		}
		if m.GetBroadcast() == true {
			flags |= 1 << 1
		}
		if m.GetPublic() == true {
			flags |= 1 << 2
		}
		if m.GetMegagroup() == true {
			flags |= 1 << 3
		}
		if len(m.GetParticipants()) > 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		x.String(m.GetTitle())
		x.Bytes(m.GetPhotoDB74F55871().Encode(layer))
		x.Int(m.GetParticipantsCount())
		if len(m.GetParticipants()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetParticipants())))
			for _, v := range m.GetParticipants() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_ChatInvite93e99b60:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatInvite93e99b60))
		// flags
		var flags int32 = 0
		if m.GetChannel() == true {
			flags |= 1 << 0
		}
		if m.GetBroadcast() == true {
			flags |= 1 << 1
		}
		if m.GetPublic() == true {
			flags |= 1 << 2
		}
		if m.GetMegagroup() == true {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.String(m.GetTitle())

		return x.GetBuf()
	case Cmd_ChatInvitedfc2f58e:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatInvitedfc2f58e))
		// flags
		var flags int32 = 0
		if m.GetChannel() == true {
			flags |= 1 << 0
		}
		if m.GetBroadcast() == true {
			flags |= 1 << 1
		}
		if m.GetPublic() == true {
			flags |= 1 << 2
		}
		if m.GetMegagroup() == true {
			flags |= 1 << 3
		}
		if len(m.GetParticipants()) > 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		x.String(m.GetTitle())
		x.Bytes(m.GetPhotoDFC2F58E98().Encode(layer))
		x.Int(m.GetParticipantsCount())
		if len(m.GetParticipants()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetParticipants())))
			for _, v := range m.GetParticipants() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatInvite encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatInvitePeek) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatInvitePeek, layer)

	switch cmd {
	case Cmd_ChatInvitePeek:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatInvitePeek))
		x.Bytes(m.GetChat().Encode(layer))
		x.Int(m.GetExpires())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatInvitePeek encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatOnlines) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatOnlines, layer)

	switch cmd {
	case Cmd_ChatOnlines:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatOnlines))
		x.Int(m.GetOnlines())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatOnlines encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatParticipant) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatParticipant, layer)

	switch cmd {
	case Cmd_ChatParticipant:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatParticipant))
		x.Int(m.GetUserId())
		x.Int(m.GetInviterId())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatParticipant encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatParticipantCreator) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatParticipantCreator, layer)

	switch cmd {
	case Cmd_ChatParticipantCreator:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatParticipantCreator))
		x.Int(m.GetUserId())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatParticipantCreator encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatParticipantAdmin) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatParticipantAdmin, layer)

	switch cmd {
	case Cmd_ChatParticipantAdmin:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatParticipantAdmin))
		x.Int(m.GetUserId())
		x.Int(m.GetInviterId())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatParticipantAdmin encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatParticipantsForbidden) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatParticipantsForbidden, layer)

	switch cmd {
	case Cmd_ChatParticipantsForbidden:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatParticipantsForbidden))
		// flags
		var flags int32 = 0
		if m.GetSelfParticipant() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetChatId())
		if m.GetSelfParticipant() != nil {
			x.Bytes(m.GetSelfParticipant().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatParticipantsForbidden encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatParticipants) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatParticipants, layer)

	switch cmd {
	case Cmd_ChatParticipants:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatParticipants))
		x.Int(m.GetChatId())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetParticipants())))
		for _, v := range m.GetParticipants() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetVersion())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatParticipants encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatPhotoEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatPhotoEmpty, layer)

	switch cmd {
	case Cmd_ChatPhotoEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatPhotoEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatPhotoEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatPhoto, layer)

	switch cmd {
	case Cmd_ChatPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatPhoto))
		x.Bytes(m.GetPhotoSmall().Encode(layer))
		x.Bytes(m.GetPhotoBig().Encode(layer))

		return x.GetBuf()
	case Cmd_ChatPhoto475cdbd5:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatPhoto475cdbd5))
		x.Bytes(m.GetPhotoSmall().Encode(layer))
		x.Bytes(m.GetPhotoBig().Encode(layer))
		x.Int(m.GetDcId())

		return x.GetBuf()
	case Cmd_ChatPhotod20b9f3c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatPhotod20b9f3c))
		// flags
		var flags int32 = 0
		if m.GetHasVideo() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetPhotoSmall().Encode(layer))
		x.Bytes(m.GetPhotoBig().Encode(layer))
		x.Int(m.GetDcId())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLClient_DHInnerData) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassClient_DHInnerData, layer)

	switch cmd {
	case Cmd_Client_DHInnerData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Client_DHInnerData))
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.Long(m.GetRetryId())
		x.String(m.GetGB())

		return x.GetBuf()

	default:
		log.Errorf("invalid Client_DHInnerData encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLCodeSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassCodeSettings, layer)

	switch cmd {
	case Cmd_CodeSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_CodeSettings))
		// flags
		var flags int32 = 0
		if m.GetAllowFlashcall() == true {
			flags |= 1 << 0
		}
		if m.GetCurrentNumber() == true {
			flags |= 1 << 1
		}
		if m.GetAppHashPersistent() == true {
			flags |= 1 << 2
		}
		if len(m.GetAppHash()) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		if len(m.GetAppHash()) > 0 {
			x.String(m.GetAppHash())
		}

		return x.GetBuf()
	case Cmd_CodeSettingsdebebe83:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_CodeSettingsdebebe83))
		// flags
		var flags int32 = 0
		if m.GetAllowFlashcall() == true {
			flags |= 1 << 0
		}
		if m.GetCurrentNumber() == true {
			flags |= 1 << 1
		}
		if m.GetAllowAppHash() == true {
			flags |= 1 << 4
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("invalid CodeSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLConfig) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassConfig, layer)

	switch cmd {
	case Cmd_Config:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Config))
		// flags
		var flags int32 = 0
		if m.GetPhonecallsEnabled() == true {
			flags |= 1 << 1
		}
		if m.GetTmpSessions() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetSuggestedLangCode()) > 0 {
			flags |= 1 << 2
		}
		if m.GetLangPackVersion() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.GetDate())
		x.Int(m.GetExpires())
		x.Bytes(m.GetTestMode().Encode(layer))
		x.Int(m.GetThisDc())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDcOptions())))
		for _, v := range m.GetDcOptions() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetChatSizeMax())
		x.Int(m.GetMegagroupSizeMax())
		x.Int(m.GetForwardedCountMax())
		x.Int(m.GetOnlineUpdatePeriodMs())
		x.Int(m.GetOfflineBlurTimeoutMs())
		x.Int(m.GetOfflineIdleTimeoutMs())
		x.Int(m.GetOnlineCloudTimeoutMs())
		x.Int(m.GetNotifyCloudDelayMs())
		x.Int(m.GetNotifyDefaultDelayMs())
		x.Int(m.GetChatBigSize())
		x.Int(m.GetPushChatPeriodMs())
		x.Int(m.GetPushChatLimit())
		x.Int(m.GetSavedGifsLimit())
		x.Int(m.GetEditTimeLimit())
		x.Int(m.GetRatingEDecay())
		x.Int(m.GetStickersRecentLimit())
		x.Int(m.GetStickersFavedLimit())
		if m.GetTmpSessions() != 0 {
			x.Int(m.GetTmpSessions())
		}
		x.Int(m.GetPinnedDialogsCountMax())
		x.Int(m.GetCallReceiveTimeoutMs())
		x.Int(m.GetCallRingTimeoutMs())
		x.Int(m.GetCallConnectTimeoutMs())
		x.Int(m.GetCallPacketTimeoutMs())
		x.String(m.GetMeUrlPrefix())
		if len(m.GetSuggestedLangCode()) > 0 {
			x.String(m.GetSuggestedLangCode())
		}
		if m.GetLangPackVersion() != 0 {
			x.Int(m.GetLangPackVersion())
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDisabledFeatures())))
		for _, v := range m.GetDisabledFeatures() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_Config317ceef4:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Config317ceef4))
		x.Int(m.GetDate())
		x.Int(m.GetExpires())
		x.Bytes(m.GetTestMode().Encode(layer))
		x.Int(m.GetThisDc())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDcOptions())))
		for _, v := range m.GetDcOptions() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetChatSizeMax())
		x.Int(m.GetMegagroupSizeMax())
		x.Int(m.GetForwardedCountMax())
		x.Int(m.GetOnlineUpdatePeriodMs())
		x.Int(m.GetOfflineBlurTimeoutMs())
		x.Int(m.GetOfflineIdleTimeoutMs())
		x.Int(m.GetOnlineCloudTimeoutMs())
		x.Int(m.GetNotifyCloudDelayMs())
		x.Int(m.GetNotifyDefaultDelayMs())
		x.Int(m.GetChatBigSize())
		x.Int(m.GetPushChatPeriodMs())
		x.Int(m.GetPushChatLimit())
		x.Int(m.GetSavedGifsLimit())
		x.Int(m.GetEditTimeLimit())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDisabledFeatures())))
		for _, v := range m.GetDisabledFeatures() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_Config3213dbba:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Config3213dbba))
		// flags
		var flags int32 = 0
		if m.GetPhonecallsEnabled() == true {
			flags |= 1 << 1
		}
		if m.GetDefaultP2PContacts() == true {
			flags |= 1 << 3
		}
		if m.GetPreloadFeaturedStickers() == true {
			flags |= 1 << 4
		}
		if m.GetIgnorePhoneEntities() == true {
			flags |= 1 << 5
		}
		if m.GetRevokePmInbox() == true {
			flags |= 1 << 6
		}
		if m.GetBlockedMode() == true {
			flags |= 1 << 8
		}
		if m.GetTmpSessions() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetAutoupdateUrlPrefix()) > 0 {
			flags |= 1 << 7
		}
		if len(m.GetGifSearchUsername()) > 0 {
			flags |= 1 << 9
		}
		if len(m.GetVenueSearchUsername()) > 0 {
			flags |= 1 << 10
		}
		if len(m.GetImgSearchUsername()) > 0 {
			flags |= 1 << 11
		}
		if len(m.GetStaticMapsProvider()) > 0 {
			flags |= 1 << 12
		}
		if len(m.GetSuggestedLangCode()) > 0 {
			flags |= 1 << 2
		}
		if m.GetLangPackVersion() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.GetDate())
		x.Int(m.GetExpires())
		x.Bytes(m.GetTestMode().Encode(layer))
		x.Int(m.GetThisDc())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDcOptions())))
		for _, v := range m.GetDcOptions() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.String(m.GetDcTxtDomainName())
		x.Int(m.GetChatSizeMax())
		x.Int(m.GetMegagroupSizeMax())
		x.Int(m.GetForwardedCountMax())
		x.Int(m.GetOnlineUpdatePeriodMs())
		x.Int(m.GetOfflineBlurTimeoutMs())
		x.Int(m.GetOfflineIdleTimeoutMs())
		x.Int(m.GetOnlineCloudTimeoutMs())
		x.Int(m.GetNotifyCloudDelayMs())
		x.Int(m.GetNotifyDefaultDelayMs())
		x.Int(m.GetPushChatPeriodMs())
		x.Int(m.GetPushChatLimit())
		x.Int(m.GetSavedGifsLimit())
		x.Int(m.GetEditTimeLimit())
		x.Int(m.GetRevokeTimeLimit())
		x.Int(m.GetRevokePmTimeLimit())
		x.Int(m.GetRatingEDecay())
		x.Int(m.GetStickersRecentLimit())
		x.Int(m.GetStickersFavedLimit())
		x.Int(m.GetChannelsReadMediaPeriod())
		if m.GetTmpSessions() != 0 {
			x.Int(m.GetTmpSessions())
		}
		x.Int(m.GetPinnedDialogsCountMax())
		x.Int(m.GetCallReceiveTimeoutMs())
		x.Int(m.GetCallRingTimeoutMs())
		x.Int(m.GetCallConnectTimeoutMs())
		x.Int(m.GetCallPacketTimeoutMs())
		x.String(m.GetMeUrlPrefix())
		if len(m.GetAutoupdateUrlPrefix()) > 0 {
			x.String(m.GetAutoupdateUrlPrefix())
		}
		if len(m.GetGifSearchUsername()) > 0 {
			x.String(m.GetGifSearchUsername())
		}
		if len(m.GetVenueSearchUsername()) > 0 {
			x.String(m.GetVenueSearchUsername())
		}
		if len(m.GetImgSearchUsername()) > 0 {
			x.String(m.GetImgSearchUsername())
		}
		if len(m.GetStaticMapsProvider()) > 0 {
			x.String(m.GetStaticMapsProvider())
		}
		x.Int(m.GetCaptionLengthMax())
		x.Int(m.GetMessageLengthMax())
		x.Int(m.GetWebfileDcId())
		if len(m.GetSuggestedLangCode()) > 0 {
			x.String(m.GetSuggestedLangCode())
		}
		if m.GetLangPackVersion() != 0 {
			x.Int(m.GetLangPackVersion())
		}

		return x.GetBuf()
	case Cmd_Config330b4067:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Config330b4067))
		// flags
		var flags int32 = 0
		if m.GetPhonecallsEnabled() == true {
			flags |= 1 << 1
		}
		if m.GetDefaultP2PContacts() == true {
			flags |= 1 << 3
		}
		if m.GetPreloadFeaturedStickers() == true {
			flags |= 1 << 4
		}
		if m.GetIgnorePhoneEntities() == true {
			flags |= 1 << 5
		}
		if m.GetRevokePmInbox() == true {
			flags |= 1 << 6
		}
		if m.GetBlockedMode() == true {
			flags |= 1 << 8
		}
		if m.GetPfsEnabled() == true {
			flags |= 1 << 13
		}
		if m.GetTmpSessions() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetAutoupdateUrlPrefix()) > 0 {
			flags |= 1 << 7
		}
		if len(m.GetGifSearchUsername()) > 0 {
			flags |= 1 << 9
		}
		if len(m.GetVenueSearchUsername()) > 0 {
			flags |= 1 << 10
		}
		if len(m.GetImgSearchUsername()) > 0 {
			flags |= 1 << 11
		}
		if len(m.GetStaticMapsProvider()) > 0 {
			flags |= 1 << 12
		}
		if len(m.GetSuggestedLangCode()) > 0 {
			flags |= 1 << 2
		}
		if m.GetLangPackVersion() != 0 {
			flags |= 1 << 2
		}
		if m.GetBaseLangPackVersion() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.GetDate())
		x.Int(m.GetExpires())
		x.Bytes(m.GetTestMode().Encode(layer))
		x.Int(m.GetThisDc())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDcOptions())))
		for _, v := range m.GetDcOptions() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.String(m.GetDcTxtDomainName())
		x.Int(m.GetChatSizeMax())
		x.Int(m.GetMegagroupSizeMax())
		x.Int(m.GetForwardedCountMax())
		x.Int(m.GetOnlineUpdatePeriodMs())
		x.Int(m.GetOfflineBlurTimeoutMs())
		x.Int(m.GetOfflineIdleTimeoutMs())
		x.Int(m.GetOnlineCloudTimeoutMs())
		x.Int(m.GetNotifyCloudDelayMs())
		x.Int(m.GetNotifyDefaultDelayMs())
		x.Int(m.GetPushChatPeriodMs())
		x.Int(m.GetPushChatLimit())
		x.Int(m.GetSavedGifsLimit())
		x.Int(m.GetEditTimeLimit())
		x.Int(m.GetRevokeTimeLimit())
		x.Int(m.GetRevokePmTimeLimit())
		x.Int(m.GetRatingEDecay())
		x.Int(m.GetStickersRecentLimit())
		x.Int(m.GetStickersFavedLimit())
		x.Int(m.GetChannelsReadMediaPeriod())
		if m.GetTmpSessions() != 0 {
			x.Int(m.GetTmpSessions())
		}
		x.Int(m.GetPinnedDialogsCountMax())
		x.Int(m.GetPinnedInfolderCountMax())
		x.Int(m.GetCallReceiveTimeoutMs())
		x.Int(m.GetCallRingTimeoutMs())
		x.Int(m.GetCallConnectTimeoutMs())
		x.Int(m.GetCallPacketTimeoutMs())
		x.String(m.GetMeUrlPrefix())
		if len(m.GetAutoupdateUrlPrefix()) > 0 {
			x.String(m.GetAutoupdateUrlPrefix())
		}
		if len(m.GetGifSearchUsername()) > 0 {
			x.String(m.GetGifSearchUsername())
		}
		if len(m.GetVenueSearchUsername()) > 0 {
			x.String(m.GetVenueSearchUsername())
		}
		if len(m.GetImgSearchUsername()) > 0 {
			x.String(m.GetImgSearchUsername())
		}
		if len(m.GetStaticMapsProvider()) > 0 {
			x.String(m.GetStaticMapsProvider())
		}
		x.Int(m.GetCaptionLengthMax())
		x.Int(m.GetMessageLengthMax())
		x.Int(m.GetWebfileDcId())
		if len(m.GetSuggestedLangCode()) > 0 {
			x.String(m.GetSuggestedLangCode())
		}
		if m.GetLangPackVersion() != 0 {
			x.Int(m.GetLangPackVersion())
		}
		if m.GetBaseLangPackVersion() != 0 {
			x.Int(m.GetBaseLangPackVersion())
		}

		return x.GetBuf()
	case Cmd_Confige6ca25f6:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Confige6ca25f6))
		// flags
		var flags int32 = 0
		if m.GetPhonecallsEnabled() == true {
			flags |= 1 << 1
		}
		if m.GetDefaultP2PContacts() == true {
			flags |= 1 << 3
		}
		if m.GetPreloadFeaturedStickers() == true {
			flags |= 1 << 4
		}
		if m.GetIgnorePhoneEntities() == true {
			flags |= 1 << 5
		}
		if m.GetRevokePmInbox() == true {
			flags |= 1 << 6
		}
		if m.GetBlockedMode() == true {
			flags |= 1 << 8
		}
		if m.GetTmpSessions() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetAutoupdateUrlPrefix()) > 0 {
			flags |= 1 << 7
		}
		if len(m.GetGifSearchUsername()) > 0 {
			flags |= 1 << 9
		}
		if len(m.GetVenueSearchUsername()) > 0 {
			flags |= 1 << 10
		}
		if len(m.GetImgSearchUsername()) > 0 {
			flags |= 1 << 11
		}
		if len(m.GetStaticMapsProvider()) > 0 {
			flags |= 1 << 12
		}
		if len(m.GetSuggestedLangCode()) > 0 {
			flags |= 1 << 2
		}
		if m.GetLangPackVersion() != 0 {
			flags |= 1 << 2
		}
		if m.GetBaseLangPackVersion() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.GetDate())
		x.Int(m.GetExpires())
		x.Bytes(m.GetTestMode().Encode(layer))
		x.Int(m.GetThisDc())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDcOptions())))
		for _, v := range m.GetDcOptions() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.String(m.GetDcTxtDomainName())
		x.Int(m.GetChatSizeMax())
		x.Int(m.GetMegagroupSizeMax())
		x.Int(m.GetForwardedCountMax())
		x.Int(m.GetOnlineUpdatePeriodMs())
		x.Int(m.GetOfflineBlurTimeoutMs())
		x.Int(m.GetOfflineIdleTimeoutMs())
		x.Int(m.GetOnlineCloudTimeoutMs())
		x.Int(m.GetNotifyCloudDelayMs())
		x.Int(m.GetNotifyDefaultDelayMs())
		x.Int(m.GetPushChatPeriodMs())
		x.Int(m.GetPushChatLimit())
		x.Int(m.GetSavedGifsLimit())
		x.Int(m.GetEditTimeLimit())
		x.Int(m.GetRevokeTimeLimit())
		x.Int(m.GetRevokePmTimeLimit())
		x.Int(m.GetRatingEDecay())
		x.Int(m.GetStickersRecentLimit())
		x.Int(m.GetStickersFavedLimit())
		x.Int(m.GetChannelsReadMediaPeriod())
		if m.GetTmpSessions() != 0 {
			x.Int(m.GetTmpSessions())
		}
		x.Int(m.GetPinnedDialogsCountMax())
		x.Int(m.GetCallReceiveTimeoutMs())
		x.Int(m.GetCallRingTimeoutMs())
		x.Int(m.GetCallConnectTimeoutMs())
		x.Int(m.GetCallPacketTimeoutMs())
		x.String(m.GetMeUrlPrefix())
		if len(m.GetAutoupdateUrlPrefix()) > 0 {
			x.String(m.GetAutoupdateUrlPrefix())
		}
		if len(m.GetGifSearchUsername()) > 0 {
			x.String(m.GetGifSearchUsername())
		}
		if len(m.GetVenueSearchUsername()) > 0 {
			x.String(m.GetVenueSearchUsername())
		}
		if len(m.GetImgSearchUsername()) > 0 {
			x.String(m.GetImgSearchUsername())
		}
		if len(m.GetStaticMapsProvider()) > 0 {
			x.String(m.GetStaticMapsProvider())
		}
		x.Int(m.GetCaptionLengthMax())
		x.Int(m.GetMessageLengthMax())
		x.Int(m.GetWebfileDcId())
		if len(m.GetSuggestedLangCode()) > 0 {
			x.String(m.GetSuggestedLangCode())
		}
		if m.GetLangPackVersion() != 0 {
			x.Int(m.GetLangPackVersion())
		}
		if m.GetBaseLangPackVersion() != 0 {
			x.Int(m.GetBaseLangPackVersion())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid Config encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContact) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContact, layer)

	switch cmd {
	case Cmd_Contact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Contact))
		x.Int(m.GetUserId())
		x.Bytes(m.GetMutual().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid Contact encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactBlocked) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactBlocked, layer)

	switch cmd {
	case Cmd_ContactBlocked:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactBlocked))
		x.Int(m.GetUserId())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactBlocked encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactLinkUnknown) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactLinkUnknown, layer)

	switch cmd {
	case Cmd_ContactLinkUnknown:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactLinkUnknown))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactLinkUnknown encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactLinkNone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactLinkNone, layer)

	switch cmd {
	case Cmd_ContactLinkNone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactLinkNone))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactLinkNone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactLinkHasPhone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactLinkHasPhone, layer)

	switch cmd {
	case Cmd_ContactLinkHasPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactLinkHasPhone))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactLinkHasPhone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactLinkContact) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactLinkContact, layer)

	switch cmd {
	case Cmd_ContactLinkContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactLinkContact))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactLinkContact encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactStatus) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactStatus, layer)

	switch cmd {
	case Cmd_ContactStatus:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactStatus))
		x.Int(m.GetUserId())
		x.Bytes(m.GetStatus().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactStatus encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsBlocked) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactsBlocked, layer)

	switch cmd {
	case Cmd_ContactsBlocked:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsBlocked))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocked1C138D1571())))
		for _, v := range m.GetBlocked1C138D1571() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_ContactsBlockedade1591:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsBlockedade1591))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlockedADE1591119())))
		for _, v := range m.GetBlockedADE1591119() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactsBlocked encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsBlockedSlice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactsBlockedSlice, layer)

	switch cmd {
	case Cmd_ContactsBlockedSlice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsBlockedSlice))
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocked1C138D1571())))
		for _, v := range m.GetBlocked1C138D1571() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_ContactsBlockedSlicee1664194:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsBlockedSlicee1664194))
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlockedADE1591119())))
		for _, v := range m.GetBlockedADE1591119() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactsBlockedSlice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsContactsNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactsContactsNotModified, layer)

	switch cmd {
	case Cmd_ContactsContactsNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsContactsNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactsContactsNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsContacts) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactsContacts, layer)

	switch cmd {
	case Cmd_ContactsContacts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsContacts))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetContacts())))
		for _, v := range m.GetContacts() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetSavedCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_ContactsContacts6f8b8cb2:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsContacts6f8b8cb2))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetContacts())))
		for _, v := range m.GetContacts() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactsContacts encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsFound) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactsFound, layer)

	switch cmd {
	case Cmd_ContactsFound:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsFound))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetResults())))
		for _, v := range m.GetResults() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_ContactsFoundb3134d9d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsFoundb3134d9d))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMyResults())))
		for _, v := range m.GetMyResults() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetResults())))
		for _, v := range m.GetResults() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactsFound encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsImportedContacts) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactsImportedContacts, layer)

	switch cmd {
	case Cmd_ContactsImportedContacts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsImportedContacts))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetImported())))
		for _, v := range m.GetImported() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPopularInvites())))
		for _, v := range m.GetPopularInvites() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.VectorLong(m.GetRetryContacts())

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_ContactsImportedContactsad524315:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsImportedContactsad524315))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetImported())))
		for _, v := range m.GetImported() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.VectorLong(m.GetRetryContacts())

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactsImportedContacts encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsLink) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactsLink, layer)

	switch cmd {
	case Cmd_ContactsLink:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsLink))
		x.Bytes(m.GetMyLink().Encode(layer))
		x.Bytes(m.GetForeignLink().Encode(layer))
		x.Bytes(m.GetUser().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactsLink encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsResolvedPeer) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactsResolvedPeer, layer)

	switch cmd {
	case Cmd_ContactsResolvedPeer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsResolvedPeer))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactsResolvedPeer encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsTopPeersNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactsTopPeersNotModified, layer)

	switch cmd {
	case Cmd_ContactsTopPeersNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsTopPeersNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactsTopPeersNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsTopPeers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactsTopPeers, layer)

	switch cmd {
	case Cmd_ContactsTopPeers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsTopPeers))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetCategories())))
		for _, v := range m.GetCategories() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactsTopPeers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLContactsTopPeersDisabled) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassContactsTopPeersDisabled, layer)

	switch cmd {
	case Cmd_ContactsTopPeersDisabled:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ContactsTopPeersDisabled))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ContactsTopPeersDisabled encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDataJSON) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDataJSON, layer)

	switch cmd {
	case Cmd_DataJSON:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DataJSON))
		x.String(m.GetData())

		return x.GetBuf()

	default:
		log.Errorf("invalid DataJSON encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDcOption) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDcOption, layer)

	switch cmd {
	case Cmd_DcOption:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DcOption))
		// flags
		var flags int32 = 0
		if m.GetIpv6() == true {
			flags |= 1 << 0
		}
		if m.GetMediaOnly() == true {
			flags |= 1 << 1
		}
		if m.GetTcpoOnly() == true {
			flags |= 1 << 2
		}
		if m.GetCdn() == true {
			flags |= 1 << 3
		}
		if m.GetStatic() == true {
			flags |= 1 << 4
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetIpAddress())
		x.Int(m.GetPort())

		return x.GetBuf()
	case Cmd_DcOption18b7a10d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DcOption18b7a10d))
		// flags
		var flags int32 = 0
		if m.GetIpv6() == true {
			flags |= 1 << 0
		}
		if m.GetMediaOnly() == true {
			flags |= 1 << 1
		}
		if m.GetTcpoOnly() == true {
			flags |= 1 << 2
		}
		if m.GetCdn() == true {
			flags |= 1 << 3
		}
		if m.GetStatic() == true {
			flags |= 1 << 4
		}
		if len(m.GetSecret()) > 0 {
			flags |= 1 << 10
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetIpAddress())
		x.Int(m.GetPort())
		if len(m.GetSecret()) > 0 {
			x.StringBytes(m.GetSecret())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid DcOption encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDestroyAuthKeyOk) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDestroyAuthKeyOk, layer)

	switch cmd {
	case Cmd_DestroyAuthKeyOk:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DestroyAuthKeyOk))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid DestroyAuthKeyOk encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDestroyAuthKeyNone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDestroyAuthKeyNone, layer)

	switch cmd {
	case Cmd_DestroyAuthKeyNone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DestroyAuthKeyNone))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid DestroyAuthKeyNone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDestroyAuthKeyFail) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDestroyAuthKeyFail, layer)

	switch cmd {
	case Cmd_DestroyAuthKeyFail:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DestroyAuthKeyFail))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid DestroyAuthKeyFail encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDestroySessionOk) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDestroySessionOk, layer)

	switch cmd {
	case Cmd_DestroySessionOk:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DestroySessionOk))
		x.Long(m.GetSessionId())

		return x.GetBuf()

	default:
		log.Errorf("invalid DestroySessionOk encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDestroySessionNone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDestroySessionNone, layer)

	switch cmd {
	case Cmd_DestroySessionNone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DestroySessionNone))
		x.Long(m.GetSessionId())

		return x.GetBuf()

	default:
		log.Errorf("invalid DestroySessionNone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDialog) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDialog, layer)

	switch cmd {
	case Cmd_Dialog:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Dialog))
		// flags
		var flags int32 = 0
		if m.GetPinned() == true {
			flags |= 1 << 2
		}
		if m.GetPts() != 0 {
			flags |= 1 << 0
		}
		if m.GetDraft() != nil {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetTopMessage())
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Int(m.GetUnreadMentionsCount())
		x.Bytes(m.GetNotifySettings().Encode(layer))
		if m.GetPts() != 0 {
			x.Int(m.GetPts())
		}
		if m.GetDraft() != nil {
			x.Bytes(m.GetDraft().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_Dialog2c171f72:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Dialog2c171f72))
		// flags
		var flags int32 = 0
		if m.GetPinned() == true {
			flags |= 1 << 2
		}
		if m.GetUnreadMark() == true {
			flags |= 1 << 3
		}
		if m.GetPts() != 0 {
			flags |= 1 << 0
		}
		if m.GetDraft() != nil {
			flags |= 1 << 1
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetTopMessage())
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Int(m.GetUnreadMentionsCount())
		x.Bytes(m.GetNotifySettings().Encode(layer))
		if m.GetPts() != 0 {
			x.Int(m.GetPts())
		}
		if m.GetDraft() != nil {
			x.Bytes(m.GetDraft().Encode(layer))
		}
		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}

		return x.GetBuf()
	case Cmd_Dialogc1dd804a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Dialogc1dd804a))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetTopMessage())
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Bytes(m.GetNotifySettings().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid Dialog encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDialogChannel) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDialogChannel, layer)

	switch cmd {
	case Cmd_DialogChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DialogChannel))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetTopMessage())
		x.Int(m.GetTopImportantMessage())
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Int(m.GetUnreadImportantCount())
		x.Bytes(m.GetNotifySettings().Encode(layer))
		x.Int(m.GetPts())

		return x.GetBuf()

	default:
		log.Errorf("invalid DialogChannel encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDialogFolder) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDialogFolder, layer)

	switch cmd {
	case Cmd_DialogFolder:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DialogFolder))
		// flags
		var flags int32 = 0
		if m.GetPinned() == true {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetFolder().Encode(layer))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetTopMessage())
		x.Int(m.GetUnreadMutedPeersCount())
		x.Int(m.GetUnreadUnmutedPeersCount())
		x.Int(m.GetUnreadMutedMessagesCount())
		x.Int(m.GetUnreadUnmutedMessagesCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid DialogFolder encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDialogFilter) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDialogFilter, layer)

	switch cmd {
	case Cmd_DialogFilter:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DialogFilter))
		// flags
		var flags int32 = 0
		if m.GetContacts() == true {
			flags |= 1 << 0
		}
		if m.GetNonContacts() == true {
			flags |= 1 << 1
		}
		if m.GetGroups() == true {
			flags |= 1 << 2
		}
		if m.GetBroadcasts() == true {
			flags |= 1 << 3
		}
		if m.GetBots() == true {
			flags |= 1 << 4
		}
		if m.GetExcludeMuted() == true {
			flags |= 1 << 11
		}
		if m.GetExcludeRead() == true {
			flags |= 1 << 12
		}
		if m.GetExcludeArchived() == true {
			flags |= 1 << 13
		}
		if len(m.GetEmoticon()) > 0 {
			flags |= 1 << 25
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetTitle())
		if len(m.GetEmoticon()) > 0 {
			x.String(m.GetEmoticon())
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPinnedPeers())))
		for _, v := range m.GetPinnedPeers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetIncludePeers())))
		for _, v := range m.GetIncludePeers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetExcludePeers())))
		for _, v := range m.GetExcludePeers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid DialogFilter encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDialogFilterSuggested) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDialogFilterSuggested, layer)

	switch cmd {
	case Cmd_DialogFilterSuggested:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DialogFilterSuggested))
		x.Bytes(m.GetFilter().Encode(layer))
		x.String(m.GetDescription())

		return x.GetBuf()

	default:
		log.Errorf("invalid DialogFilterSuggested encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDialogPeer) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDialogPeer, layer)

	switch cmd {
	case Cmd_DialogPeer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DialogPeer))
		x.Bytes(m.GetPeer().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid DialogPeer encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDialogPeerFolder) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDialogPeerFolder, layer)

	switch cmd {
	case Cmd_DialogPeerFolder:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DialogPeerFolder))
		x.Int(m.GetFolderId())

		return x.GetBuf()

	default:
		log.Errorf("invalid DialogPeerFolder encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDisabledFeature) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDisabledFeature, layer)

	switch cmd {
	case Cmd_DisabledFeature:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DisabledFeature))
		x.String(m.GetFeature())
		x.String(m.GetDescription())

		return x.GetBuf()

	default:
		log.Errorf("invalid DisabledFeature encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDocumentEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDocumentEmpty, layer)

	switch cmd {
	case Cmd_DocumentEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DocumentEmpty))
		x.Long(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid DocumentEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDocument) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDocument, layer)

	switch cmd {
	case Cmd_Document:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Document))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.String(m.GetMimeType())
		x.Int(m.GetSize_())
		x.Bytes(m.GetThumb().Encode(layer))
		x.Int(m.GetDcId())
		x.Int(m.GetVersion())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_Document1e87342b:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Document1e87342b))
		// flags
		var flags int32 = 0
		if len(m.GetThumbs()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetVideoThumbs()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())
		x.Int(m.GetDate())
		x.String(m.GetMimeType())
		x.Int(m.GetSize_())
		if len(m.GetThumbs()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetThumbs())))
			for _, v := range m.GetThumbs() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if len(m.GetVideoThumbs()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetVideoThumbs())))
			for _, v := range m.GetVideoThumbs() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		x.Int(m.GetDcId())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_Document59534e4c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Document59534e4c))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())
		x.Int(m.GetDate())
		x.String(m.GetMimeType())
		x.Int(m.GetSize_())
		x.Bytes(m.GetThumb().Encode(layer))
		x.Int(m.GetDcId())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_Document9ba29cc1:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Document9ba29cc1))
		// flags
		var flags int32 = 0
		if len(m.GetThumbs()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())
		x.Int(m.GetDate())
		x.String(m.GetMimeType())
		x.Int(m.GetSize_())
		if len(m.GetThumbs()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetThumbs())))
			for _, v := range m.GetThumbs() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		x.Int(m.GetDcId())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_Documentf9a39f4f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Documentf9a39f4f))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.String(m.GetMimeType())
		x.Int(m.GetSize_())
		x.Bytes(m.GetThumb().Encode(layer))
		x.Int(m.GetDcId())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid Document encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDocumentAttributeImageSize) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDocumentAttributeImageSize, layer)

	switch cmd {
	case Cmd_DocumentAttributeImageSize:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DocumentAttributeImageSize))
		x.Int(m.GetW())
		x.Int(m.GetH())

		return x.GetBuf()

	default:
		log.Errorf("invalid DocumentAttributeImageSize encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDocumentAttributeAnimated) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDocumentAttributeAnimated, layer)

	switch cmd {
	case Cmd_DocumentAttributeAnimated:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DocumentAttributeAnimated))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid DocumentAttributeAnimated encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDocumentAttributeSticker) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDocumentAttributeSticker, layer)

	switch cmd {
	case Cmd_DocumentAttributeSticker:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DocumentAttributeSticker))
		// flags
		var flags int32 = 0
		if m.GetMask() == true {
			flags |= 1 << 1
		}
		if m.GetMaskCoords() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetAlt())
		x.Bytes(m.GetStickerset().Encode(layer))
		if m.GetMaskCoords() != nil {
			x.Bytes(m.GetMaskCoords().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_DocumentAttributeSticker3a556302:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DocumentAttributeSticker3a556302))
		x.String(m.GetAlt())
		x.Bytes(m.GetStickerset().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid DocumentAttributeSticker encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDocumentAttributeVideo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDocumentAttributeVideo, layer)

	switch cmd {
	case Cmd_DocumentAttributeVideo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DocumentAttributeVideo))
		// flags
		var flags int32 = 0
		if m.GetRoundMessage() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetDuration())
		x.Int(m.GetW())
		x.Int(m.GetH())

		return x.GetBuf()
	case Cmd_DocumentAttributeVideo5910cccb:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DocumentAttributeVideo5910cccb))
		x.Int(m.GetDuration())
		x.Int(m.GetW())
		x.Int(m.GetH())

		return x.GetBuf()

	default:
		log.Errorf("invalid DocumentAttributeVideo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDocumentAttributeAudio) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDocumentAttributeAudio, layer)

	switch cmd {
	case Cmd_DocumentAttributeAudio:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DocumentAttributeAudio))
		// flags
		var flags int32 = 0
		if m.GetVoice() == true {
			flags |= 1 << 10
		}
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetPerformer()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetWaveform()) > 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.GetDuration())
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetPerformer()) > 0 {
			x.String(m.GetPerformer())
		}
		if len(m.GetWaveform()) > 0 {
			x.StringBytes(m.GetWaveform())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid DocumentAttributeAudio encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDocumentAttributeFilename) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDocumentAttributeFilename, layer)

	switch cmd {
	case Cmd_DocumentAttributeFilename:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DocumentAttributeFilename))
		x.String(m.GetFileName())

		return x.GetBuf()

	default:
		log.Errorf("invalid DocumentAttributeFilename encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDocumentAttributeHasStickers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDocumentAttributeHasStickers, layer)

	switch cmd {
	case Cmd_DocumentAttributeHasStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DocumentAttributeHasStickers))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid DocumentAttributeHasStickers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDraftMessageEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDraftMessageEmpty, layer)

	switch cmd {
	case Cmd_DraftMessageEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DraftMessageEmpty))
		_ = m

		return x.GetBuf()
	case Cmd_DraftMessageEmpty1b0c841a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DraftMessageEmpty1b0c841a))
		// flags
		var flags int32 = 0
		if m.GetDate() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetDate() != 0 {
			x.Int(m.GetDate())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid DraftMessageEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDraftMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDraftMessage, layer)

	switch cmd {
	case Cmd_DraftMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DraftMessage))
		// flags
		var flags int32 = 0
		if m.GetNoWebpage() == true {
			flags |= 1 << 1
		}
		if m.GetReplyToMsgId() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		if m.GetReplyToMsgId() != 0 {
			x.Int(m.GetReplyToMsgId())
		}
		x.String(m.GetMessage())
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid DraftMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEmojiKeyword) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEmojiKeyword, layer)

	switch cmd {
	case Cmd_EmojiKeyword:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EmojiKeyword))
		x.String(m.GetKeyword())
		x.VectorString(m.GetEmoticons())

		return x.GetBuf()

	default:
		log.Errorf("invalid EmojiKeyword encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEmojiKeywordDeleted) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEmojiKeywordDeleted, layer)

	switch cmd {
	case Cmd_EmojiKeywordDeleted:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EmojiKeywordDeleted))
		x.String(m.GetKeyword())
		x.VectorString(m.GetEmoticons())

		return x.GetBuf()

	default:
		log.Errorf("invalid EmojiKeywordDeleted encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEmojiKeywordsDifference) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEmojiKeywordsDifference, layer)

	switch cmd {
	case Cmd_EmojiKeywordsDifference:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EmojiKeywordsDifference))
		x.String(m.GetLangCode())
		x.Int(m.GetFromVersion())
		x.Int(m.GetVersion())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetKeywords())))
		for _, v := range m.GetKeywords() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid EmojiKeywordsDifference encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEmojiLanguage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEmojiLanguage, layer)

	switch cmd {
	case Cmd_EmojiLanguage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EmojiLanguage))
		x.String(m.GetLangCode())

		return x.GetBuf()

	default:
		log.Errorf("invalid EmojiLanguage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEmojiURL) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEmojiURL, layer)

	switch cmd {
	case Cmd_EmojiURL:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EmojiURL))
		x.String(m.GetUrl())

		return x.GetBuf()

	default:
		log.Errorf("invalid EmojiURL encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEncryptedChatEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEncryptedChatEmpty, layer)

	switch cmd {
	case Cmd_EncryptedChatEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EncryptedChatEmpty))
		x.Int(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid EncryptedChatEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEncryptedChatWaiting) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEncryptedChatWaiting, layer)

	switch cmd {
	case Cmd_EncryptedChatWaiting:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EncryptedChatWaiting))
		x.Int(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())

		return x.GetBuf()

	default:
		log.Errorf("invalid EncryptedChatWaiting encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEncryptedChatRequested) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEncryptedChatRequested, layer)

	switch cmd {
	case Cmd_EncryptedChatRequested:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EncryptedChatRequested))
		x.Int(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())
		x.StringBytes(m.GetGA())

		return x.GetBuf()
	case Cmd_EncryptedChatRequested62718a82:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EncryptedChatRequested62718a82))
		// flags
		var flags int32 = 0
		if m.GetFolderId() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		x.Int(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())
		x.StringBytes(m.GetGA())

		return x.GetBuf()

	default:
		log.Errorf("invalid EncryptedChatRequested encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEncryptedChat) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEncryptedChat, layer)

	switch cmd {
	case Cmd_EncryptedChat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EncryptedChat))
		x.Int(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())
		x.StringBytes(m.GetGAOrB())
		x.Long(m.GetKeyFingerprint())

		return x.GetBuf()

	default:
		log.Errorf("invalid EncryptedChat encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEncryptedChatDiscarded) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEncryptedChatDiscarded, layer)

	switch cmd {
	case Cmd_EncryptedChatDiscarded:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EncryptedChatDiscarded))
		x.Int(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid EncryptedChatDiscarded encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEncryptedFileEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEncryptedFileEmpty, layer)

	switch cmd {
	case Cmd_EncryptedFileEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EncryptedFileEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid EncryptedFileEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEncryptedFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEncryptedFile, layer)

	switch cmd {
	case Cmd_EncryptedFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EncryptedFile))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetSize_())
		x.Int(m.GetDcId())
		x.Int(m.GetKeyFingerprint())

		return x.GetBuf()

	default:
		log.Errorf("invalid EncryptedFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEncryptedMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEncryptedMessage, layer)

	switch cmd {
	case Cmd_EncryptedMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EncryptedMessage))
		x.Long(m.GetRandomId())
		x.Int(m.GetChatId())
		x.Int(m.GetDate())
		x.StringBytes(m.GetBytes())
		x.Bytes(m.GetFile().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid EncryptedMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLEncryptedMessageService) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassEncryptedMessageService, layer)

	switch cmd {
	case Cmd_EncryptedMessageService:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_EncryptedMessageService))
		x.Long(m.GetRandomId())
		x.Int(m.GetChatId())
		x.Int(m.GetDate())
		x.StringBytes(m.GetBytes())

		return x.GetBuf()

	default:
		log.Errorf("invalid EncryptedMessageService encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLError) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassError, layer)

	switch cmd {
	case Cmd_Error:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Error))
		x.Int(m.GetCode())
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid Error encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatInviteEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatInviteEmpty, layer)

	switch cmd {
	case Cmd_ChatInviteEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatInviteEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatInviteEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLChatInviteExported) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassChatInviteExported, layer)

	switch cmd {
	case Cmd_ChatInviteExported:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ChatInviteExported))
		x.String(m.GetLink())

		return x.GetBuf()

	default:
		log.Errorf("invalid ChatInviteExported encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLExportedMessageLink) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassExportedMessageLink, layer)

	switch cmd {
	case Cmd_ExportedMessageLink:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ExportedMessageLink))
		x.String(m.GetLink())

		return x.GetBuf()
	case Cmd_ExportedMessageLink5dab1af4:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ExportedMessageLink5dab1af4))
		x.String(m.GetLink())
		x.String(m.GetHtml())

		return x.GetBuf()

	default:
		log.Errorf("invalid ExportedMessageLink encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFileHash) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassFileHash, layer)

	switch cmd {
	case Cmd_FileHash:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FileHash))
		x.Int(m.GetOffset())
		x.Int(m.GetLimit())
		x.StringBytes(m.GetHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid FileHash encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFileLocationUnavailable) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassFileLocationUnavailable, layer)

	switch cmd {
	case Cmd_FileLocationUnavailable:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FileLocationUnavailable))
		x.Long(m.GetVolumeId())
		x.Int(m.GetLocalId())
		x.Long(m.GetSecret())

		return x.GetBuf()

	default:
		log.Errorf("invalid FileLocationUnavailable encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFileLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassFileLocation, layer)

	switch cmd {
	case Cmd_FileLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FileLocation))
		x.Int(m.GetDcId())
		x.Long(m.GetVolumeId())
		x.Int(m.GetLocalId())
		x.Long(m.GetSecret())

		return x.GetBuf()
	case Cmd_FileLocation91d11eb:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FileLocation91d11eb))
		x.Int(m.GetDcId())
		x.Long(m.GetVolumeId())
		x.Int(m.GetLocalId())
		x.Long(m.GetSecret())
		x.StringBytes(m.GetFileReference())

		return x.GetBuf()

	default:
		log.Errorf("invalid FileLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFileLocationToBeDeprecated) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassFileLocationToBeDeprecated, layer)

	switch cmd {
	case Cmd_FileLocationToBeDeprecated:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FileLocationToBeDeprecated))
		x.Long(m.GetVolumeId())
		x.Int(m.GetLocalId())

		return x.GetBuf()

	default:
		log.Errorf("invalid FileLocationToBeDeprecated encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFolder) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassFolder, layer)

	switch cmd {
	case Cmd_Folder:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Folder))
		// flags
		var flags int32 = 0
		if m.GetAutofillNewBroadcasts() == true {
			flags |= 1 << 0
		}
		if m.GetAutofillPublicGroups() == true {
			flags |= 1 << 1
		}
		if m.GetAutofillNewCorrespondents() == true {
			flags |= 1 << 2
		}
		if m.GetPhoto() != nil {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetTitle())
		if m.GetPhoto() != nil {
			x.Bytes(m.GetPhoto().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid Folder encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFolderPeer) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassFolderPeer, layer)

	switch cmd {
	case Cmd_FolderPeer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FolderPeer))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetFolderId())

		return x.GetBuf()

	default:
		log.Errorf("invalid FolderPeer encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFoundGif) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassFoundGif, layer)

	switch cmd {
	case Cmd_FoundGif:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FoundGif))
		x.String(m.GetUrl())
		x.String(m.GetThumbUrl())
		x.String(m.GetContentUrl())
		x.String(m.GetContentType())
		x.Int(m.GetW())
		x.Int(m.GetH())

		return x.GetBuf()

	default:
		log.Errorf("invalid FoundGif encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFoundGifCached) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassFoundGifCached, layer)

	switch cmd {
	case Cmd_FoundGifCached:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FoundGifCached))
		x.String(m.GetUrl())
		x.Bytes(m.GetPhoto().Encode(layer))
		x.Bytes(m.GetDocument().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid FoundGifCached encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFutureSalt) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassFutureSalt, layer)

	switch cmd {
	case Cmd_FutureSalt:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FutureSalt))
		x.Int(m.GetValidSince())
		x.Int(m.GetValidUntil())
		x.Long(m.GetSalt())

		return x.GetBuf()

	default:
		log.Errorf("invalid FutureSalt encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLFutureSalts) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassFutureSalts, layer)

	switch cmd {
	case Cmd_FutureSalts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_FutureSalts))
		x.Long(m.GetReqMsgId())
		x.Int(m.GetNow())
		x.Int(int32(len(m.GetSalts())))
		for _, v := range m.GetSalts() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid FutureSalts encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLGame) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassGame, layer)

	switch cmd {
	case Cmd_Game:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Game))
		// flags
		var flags int32 = 0
		if m.GetDocument() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.String(m.GetShortName())
		x.String(m.GetTitle())
		x.String(m.GetDescription())
		x.Bytes(m.GetPhoto().Encode(layer))
		if m.GetDocument() != nil {
			x.Bytes(m.GetDocument().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid Game encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLGeoPointEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassGeoPointEmpty, layer)

	switch cmd {
	case Cmd_GeoPointEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_GeoPointEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid GeoPointEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLGeoPoint) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassGeoPoint, layer)

	switch cmd {
	case Cmd_GeoPoint:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_GeoPoint))
		x.Double(m.GetLong())
		x.Double(m.GetLat())

		return x.GetBuf()
	case Cmd_GeoPoint296f104:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_GeoPoint296f104))
		x.Double(m.GetLong())
		x.Double(m.GetLat())
		x.Long(m.GetAccessHash())

		return x.GetBuf()
	case Cmd_GeoPointb2a2f663:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_GeoPointb2a2f663))
		// flags
		var flags int32 = 0
		if m.GetAccuracyRadius() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Double(m.GetLong())
		x.Double(m.GetLat())
		x.Long(m.GetAccessHash())
		if m.GetAccuracyRadius() != 0 {
			x.Int(m.GetAccuracyRadius())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid GeoPoint encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLGlobalPrivacySettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassGlobalPrivacySettings, layer)

	switch cmd {
	case Cmd_GlobalPrivacySettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_GlobalPrivacySettings))
		// flags
		var flags int32 = 0
		if m.GetArchiveAndMuteNewNoncontactPeers() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetArchiveAndMuteNewNoncontactPeers() != nil {
			x.Bytes(m.GetArchiveAndMuteNewNoncontactPeers().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid GlobalPrivacySettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLGroupCallDiscarded) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassGroupCallDiscarded, layer)

	switch cmd {
	case Cmd_GroupCallDiscarded:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_GroupCallDiscarded))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDuration())

		return x.GetBuf()

	default:
		log.Errorf("invalid GroupCallDiscarded encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLGroupCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassGroupCall, layer)

	switch cmd {
	case Cmd_GroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_GroupCall))
		// flags
		var flags int32 = 0
		if m.GetJoinMuted() == true {
			flags |= 1 << 1
		}
		if m.GetCanChangeJoinMuted() == true {
			flags |= 1 << 2
		}
		if m.GetParams() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetParticipantsCount())
		if m.GetParams() != nil {
			x.Bytes(m.GetParams().Encode(layer))
		}
		x.Int(m.GetVersion())

		return x.GetBuf()

	default:
		log.Errorf("invalid GroupCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLGroupCallParticipant) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassGroupCallParticipant, layer)

	switch cmd {
	case Cmd_GroupCallParticipant:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_GroupCallParticipant))
		// flags
		var flags int32 = 0
		if m.GetMuted() == true {
			flags |= 1 << 0
		}
		if m.GetLeft() == true {
			flags |= 1 << 1
		}
		if m.GetCanSelfUnmute() == true {
			flags |= 1 << 2
		}
		if m.GetJustJoined() == true {
			flags |= 1 << 4
		}
		if m.GetVersioned() == true {
			flags |= 1 << 5
		}
		if m.GetActiveDate() != 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Int(m.GetUserId())
		x.Int(m.GetDate())
		if m.GetActiveDate() != 0 {
			x.Int(m.GetActiveDate())
		}
		x.Int(m.GetSource())

		return x.GetBuf()

	default:
		log.Errorf("invalid GroupCallParticipant encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpAppChangelogEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpAppChangelogEmpty, layer)

	switch cmd {
	case Cmd_HelpAppChangelogEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpAppChangelogEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpAppChangelogEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpAppChangelog) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpAppChangelog, layer)

	switch cmd {
	case Cmd_HelpAppChangelog:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpAppChangelog))
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpAppChangelog encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpAppUpdate) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpAppUpdate, layer)

	switch cmd {
	case Cmd_HelpAppUpdate:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpAppUpdate))
		x.Int(m.GetId())
		x.Bytes(m.GetCritical().Encode(layer))
		x.String(m.GetUrl())
		x.String(m.GetText())

		return x.GetBuf()
	case Cmd_HelpAppUpdate1da7158f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpAppUpdate1da7158f))
		// flags
		var flags int32 = 0
		if m.GetPopup() == true {
			flags |= 1 << 0
		}
		if m.GetDocument() != nil {
			flags |= 1 << 1
		}
		if len(m.GetUrl()) > 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.String(m.GetVersion())
		x.String(m.GetText())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetEntities())))
		for _, v := range m.GetEntities() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetDocument() != nil {
			x.Bytes(m.GetDocument().Encode(layer))
		}
		if len(m.GetUrl()) > 0 {
			x.String(m.GetUrl())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpAppUpdate encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpNoAppUpdate) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpNoAppUpdate, layer)

	switch cmd {
	case Cmd_HelpNoAppUpdate:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpNoAppUpdate))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpNoAppUpdate encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpConfigSimple) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpConfigSimple, layer)

	switch cmd {
	case Cmd_HelpConfigSimple:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpConfigSimple))
		x.Int(m.GetDate())
		x.Int(m.GetExpires())
		x.Int(m.GetDcId())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetIpPortList())))
		for _, v := range m.GetIpPortList() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_HelpConfigSimple5a592a6c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpConfigSimple5a592a6c))
		x.Int(m.GetDate())
		x.Int(m.GetExpires())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetRules())))
		for _, v := range m.GetRules() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpConfigSimple encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpCountriesListNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpCountriesListNotModified, layer)

	switch cmd {
	case Cmd_HelpCountriesListNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpCountriesListNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpCountriesListNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpCountriesList) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpCountriesList, layer)

	switch cmd {
	case Cmd_HelpCountriesList:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpCountriesList))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetCountries())))
		for _, v := range m.GetCountries() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpCountriesList encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpCountry) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpCountry, layer)

	switch cmd {
	case Cmd_HelpCountry:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpCountry))
		// flags
		var flags int32 = 0
		if m.GetHidden() == true {
			flags |= 1 << 0
		}
		if len(m.GetName()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.String(m.GetIso2())
		x.String(m.GetDefaultName())
		if len(m.GetName()) > 0 {
			x.String(m.GetName())
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetCountryCodes())))
		for _, v := range m.GetCountryCodes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpCountry encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpCountryCode) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpCountryCode, layer)

	switch cmd {
	case Cmd_HelpCountryCode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpCountryCode))
		// flags
		var flags int32 = 0
		if len(m.GetPrefixes()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetPatterns()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.String(m.GetCountryCode())
		if len(m.GetPrefixes()) > 0 {
			x.VectorString(m.GetPrefixes())

		}
		if len(m.GetPatterns()) > 0 {
			x.VectorString(m.GetPatterns())

		}

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpCountryCode encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpDeepLinkInfoEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpDeepLinkInfoEmpty, layer)

	switch cmd {
	case Cmd_HelpDeepLinkInfoEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpDeepLinkInfoEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpDeepLinkInfoEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpDeepLinkInfo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpDeepLinkInfo, layer)

	switch cmd {
	case Cmd_HelpDeepLinkInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpDeepLinkInfo))
		// flags
		var flags int32 = 0
		if m.GetUpdateApp() == true {
			flags |= 1 << 0
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.String(m.GetMessage())
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpDeepLinkInfo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpInviteText) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpInviteText, layer)

	switch cmd {
	case Cmd_HelpInviteText:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpInviteText))
		x.String(m.GetMessage())

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpInviteText encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpPassportConfigNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpPassportConfigNotModified, layer)

	switch cmd {
	case Cmd_HelpPassportConfigNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpPassportConfigNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpPassportConfigNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpPassportConfig) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpPassportConfig, layer)

	switch cmd {
	case Cmd_HelpPassportConfig:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpPassportConfig))
		x.Int(m.GetHash())
		x.Bytes(m.GetCountriesLangs().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpPassportConfig encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpPromoDataEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpPromoDataEmpty, layer)

	switch cmd {
	case Cmd_HelpPromoDataEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpPromoDataEmpty))
		x.Int(m.GetExpires())

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpPromoDataEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpPromoData) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpPromoData, layer)

	switch cmd {
	case Cmd_HelpPromoData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpPromoData))
		// flags
		var flags int32 = 0
		if m.GetProxy() == true {
			flags |= 1 << 0
		}
		if m.GetPsa() == true {
			flags |= 1 << 1
		}
		if len(m.GetPsaMessage()) > 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.GetExpires())
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if len(m.GetPsaMessage()) > 0 {
			x.String(m.GetPsaMessage())
		}

		return x.GetBuf()
	case Cmd_HelpPromoData8c39793f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpPromoData8c39793f))
		// flags
		var flags int32 = 0
		if m.GetProxy() == true {
			flags |= 1 << 0
		}
		if len(m.GetPsaType()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetPsaMessage()) > 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.GetExpires())
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if len(m.GetPsaType()) > 0 {
			x.String(m.GetPsaType())
		}
		if len(m.GetPsaMessage()) > 0 {
			x.String(m.GetPsaMessage())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpPromoData encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpProxyDataEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpProxyDataEmpty, layer)

	switch cmd {
	case Cmd_HelpProxyDataEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpProxyDataEmpty))
		x.Int(m.GetExpires())

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpProxyDataEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpProxyDataPromo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpProxyDataPromo, layer)

	switch cmd {
	case Cmd_HelpProxyDataPromo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpProxyDataPromo))
		x.Int(m.GetExpires())
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpProxyDataPromo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpRecentMeUrls) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpRecentMeUrls, layer)

	switch cmd {
	case Cmd_HelpRecentMeUrls:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpRecentMeUrls))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUrls())))
		for _, v := range m.GetUrls() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpRecentMeUrls encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpSupport) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpSupport, layer)

	switch cmd {
	case Cmd_HelpSupport:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpSupport))
		x.String(m.GetPhoneNumber())
		x.Bytes(m.GetUser().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpSupport encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpSupportName) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpSupportName, layer)

	switch cmd {
	case Cmd_HelpSupportName:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpSupportName))
		x.String(m.GetName())

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpSupportName encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpTermsOfService) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpTermsOfService, layer)

	switch cmd {
	case Cmd_HelpTermsOfService:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpTermsOfService))
		x.String(m.GetText())

		return x.GetBuf()
	case Cmd_HelpTermsOfService780a0310:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpTermsOfService780a0310))
		// flags
		var flags int32 = 0
		if m.GetPopup() == true {
			flags |= 1 << 0
		}
		if m.GetMinAgeConfirm() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetId().Encode(layer))
		x.String(m.GetText())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetEntities())))
		for _, v := range m.GetEntities() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMinAgeConfirm() != 0 {
			x.Int(m.GetMinAgeConfirm())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpTermsOfService encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpTermsOfServiceUpdateEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpTermsOfServiceUpdateEmpty, layer)

	switch cmd {
	case Cmd_HelpTermsOfServiceUpdateEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpTermsOfServiceUpdateEmpty))
		x.Int(m.GetExpires())

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpTermsOfServiceUpdateEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpTermsOfServiceUpdate) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpTermsOfServiceUpdate, layer)

	switch cmd {
	case Cmd_HelpTermsOfServiceUpdate:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpTermsOfServiceUpdate))
		x.Int(m.GetExpires())
		x.Bytes(m.GetTermsOfService().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpTermsOfServiceUpdate encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpUserInfoEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpUserInfoEmpty, layer)

	switch cmd {
	case Cmd_HelpUserInfoEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpUserInfoEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpUserInfoEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHelpUserInfo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHelpUserInfo, layer)

	switch cmd {
	case Cmd_HelpUserInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HelpUserInfo))
		x.String(m.GetMessage())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetEntities())))
		for _, v := range m.GetEntities() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.String(m.GetAuthor())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid HelpUserInfo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHighScore) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHighScore, layer)

	switch cmd {
	case Cmd_HighScore:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HighScore))
		x.Int(m.GetPos())
		x.Int(m.GetUserId())
		x.Int(m.GetScore())

		return x.GetBuf()

	default:
		log.Errorf("invalid HighScore encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLHttpWait) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassHttpWait, layer)

	switch cmd {
	case Cmd_HttpWait:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_HttpWait))
		x.Int(m.GetMaxDelay())
		x.Int(m.GetWaitAfter())
		x.Int(m.GetMaxWait())

		return x.GetBuf()

	default:
		log.Errorf("invalid HttpWait encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLImportedContact) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassImportedContact, layer)

	switch cmd {
	case Cmd_ImportedContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ImportedContact))
		x.Int(m.GetUserId())
		x.Long(m.GetClientId())

		return x.GetBuf()

	default:
		log.Errorf("invalid ImportedContact encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInlineBotSwitchPM) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInlineBotSwitchPM, layer)

	switch cmd {
	case Cmd_InlineBotSwitchPM:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InlineBotSwitchPM))
		x.String(m.GetText())
		x.String(m.GetStartParam())

		return x.GetBuf()

	default:
		log.Errorf("invalid InlineBotSwitchPM encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInlineQueryPeerTypeSameBotPM) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInlineQueryPeerTypeSameBotPM, layer)

	switch cmd {
	case Cmd_InlineQueryPeerTypeSameBotPM:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InlineQueryPeerTypeSameBotPM))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InlineQueryPeerTypeSameBotPM encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInlineQueryPeerTypePM) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInlineQueryPeerTypePM, layer)

	switch cmd {
	case Cmd_InlineQueryPeerTypePM:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InlineQueryPeerTypePM))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InlineQueryPeerTypePM encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInlineQueryPeerTypeChat) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInlineQueryPeerTypeChat, layer)

	switch cmd {
	case Cmd_InlineQueryPeerTypeChat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InlineQueryPeerTypeChat))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InlineQueryPeerTypeChat encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInlineQueryPeerTypeMegagroup) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInlineQueryPeerTypeMegagroup, layer)

	switch cmd {
	case Cmd_InlineQueryPeerTypeMegagroup:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InlineQueryPeerTypeMegagroup))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InlineQueryPeerTypeMegagroup encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInlineQueryPeerTypeBroadcast) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInlineQueryPeerTypeBroadcast, layer)

	switch cmd {
	case Cmd_InlineQueryPeerTypeBroadcast:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InlineQueryPeerTypeBroadcast))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InlineQueryPeerTypeBroadcast encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputAppEvent) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputAppEvent, layer)

	switch cmd {
	case Cmd_InputAppEvent:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputAppEvent))
		x.Double(m.GetTime())
		x.String(m.GetType())
		x.Long(m.GetPeer())
		x.String(m.GetData770656A871())

		return x.GetBuf()
	case Cmd_InputAppEvent1d1b1245:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputAppEvent1d1b1245))
		x.Double(m.GetTime())
		x.String(m.GetType())
		x.Long(m.GetPeer())
		x.Bytes(m.GetData1D1B124590().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputAppEvent encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputBotInlineMessageMediaAuto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputBotInlineMessageMediaAuto, layer)

	switch cmd {
	case Cmd_InputBotInlineMessageMediaAuto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageMediaAuto))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetCaption())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_InputBotInlineMessageMediaAuto3380c786:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageMediaAuto3380c786))
		// flags
		var flags int32 = 0
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 1
		}
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetMessage())
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputBotInlineMessageMediaAuto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputBotInlineMessageText) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputBotInlineMessageText, layer)

	switch cmd {
	case Cmd_InputBotInlineMessageText:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageText))
		// flags
		var flags int32 = 0
		if m.GetNoWebpage() == true {
			flags |= 1 << 0
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 1
		}
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetMessage())
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputBotInlineMessageText encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputBotInlineMessageMediaGeo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputBotInlineMessageMediaGeo, layer)

	switch cmd {
	case Cmd_InputBotInlineMessageMediaGeo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageMediaGeo))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetGeoPoint().Encode(layer))
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_InputBotInlineMessageMediaGeo96929a85:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageMediaGeo96929a85))
		// flags
		var flags int32 = 0
		if m.GetHeading() != 0 {
			flags |= 1 << 0
		}
		if m.GetPeriod() != 0 {
			flags |= 1 << 1
		}
		if m.GetProximityNotificationRadius() != 0 {
			flags |= 1 << 3
		}
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetGeoPoint().Encode(layer))
		if m.GetHeading() != 0 {
			x.Int(m.GetHeading())
		}
		if m.GetPeriod() != 0 {
			x.Int(m.GetPeriod())
		}
		if m.GetProximityNotificationRadius() != 0 {
			x.Int(m.GetProximityNotificationRadius())
		}
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_InputBotInlineMessageMediaGeoc1b15d65:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageMediaGeoc1b15d65))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetGeoPoint().Encode(layer))
		x.Int(m.GetPeriod())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputBotInlineMessageMediaGeo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputBotInlineMessageMediaVenue) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputBotInlineMessageMediaVenue, layer)

	switch cmd {
	case Cmd_InputBotInlineMessageMediaVenue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageMediaVenue))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetGeoPoint().Encode(layer))
		x.String(m.GetTitle())
		x.String(m.GetAddress())
		x.String(m.GetProvider())
		x.String(m.GetVenueId())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_InputBotInlineMessageMediaVenue417bbf11:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageMediaVenue417bbf11))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetGeoPoint().Encode(layer))
		x.String(m.GetTitle())
		x.String(m.GetAddress())
		x.String(m.GetProvider())
		x.String(m.GetVenueId())
		x.String(m.GetVenueType())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputBotInlineMessageMediaVenue encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputBotInlineMessageMediaContact) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputBotInlineMessageMediaContact, layer)

	switch cmd {
	case Cmd_InputBotInlineMessageMediaContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageMediaContact))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetPhoneNumber())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_InputBotInlineMessageMediaContacta6edbffd:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageMediaContacta6edbffd))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetPhoneNumber())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())
		x.String(m.GetVcard())
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputBotInlineMessageMediaContact encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputBotInlineMessageGame) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputBotInlineMessageGame, layer)

	switch cmd {
	case Cmd_InputBotInlineMessageGame:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageGame))
		// flags
		var flags int32 = 0
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputBotInlineMessageGame encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputBotInlineMessageID) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputBotInlineMessageID, layer)

	switch cmd {
	case Cmd_InputBotInlineMessageID:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineMessageID))
		x.Int(m.GetDcId())
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputBotInlineMessageID encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputBotInlineResult) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputBotInlineResult, layer)

	switch cmd {
	case Cmd_InputBotInlineResult:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineResult))
		// flags
		var flags int32 = 0
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetUrl()) > 0 {
			flags |= 1 << 3
		}
		if len(m.GetThumbUrl()) > 0 {
			flags |= 1 << 4
		}
		if len(m.GetContentUrl()) > 0 {
			flags |= 1 << 5
		}
		if len(m.GetContentType()) > 0 {
			flags |= 1 << 5
		}
		if m.GetW() != 0 {
			flags |= 1 << 6
		}
		if m.GetH() != 0 {
			flags |= 1 << 6
		}
		if m.GetDuration() != 0 {
			flags |= 1 << 7
		}
		x.Int(flags)

		x.String(m.GetId())
		x.String(m.GetType())
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		if len(m.GetUrl()) > 0 {
			x.String(m.GetUrl())
		}
		if len(m.GetThumbUrl()) > 0 {
			x.String(m.GetThumbUrl())
		}
		if len(m.GetContentUrl()) > 0 {
			x.String(m.GetContentUrl())
		}
		if len(m.GetContentType()) > 0 {
			x.String(m.GetContentType())
		}
		if m.GetW() != 0 {
			x.Int(m.GetW())
		}
		if m.GetH() != 0 {
			x.Int(m.GetH())
		}
		if m.GetDuration() != 0 {
			x.Int(m.GetDuration())
		}
		x.Bytes(m.GetSendMessage().Encode(layer))

		return x.GetBuf()
	case Cmd_InputBotInlineResult88bf9319:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineResult88bf9319))
		// flags
		var flags int32 = 0
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetUrl()) > 0 {
			flags |= 1 << 3
		}
		if m.GetThumb() != nil {
			flags |= 1 << 4
		}
		if m.GetContent() != nil {
			flags |= 1 << 5
		}
		x.Int(flags)

		x.String(m.GetId())
		x.String(m.GetType())
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		if len(m.GetUrl()) > 0 {
			x.String(m.GetUrl())
		}
		if m.GetThumb() != nil {
			x.Bytes(m.GetThumb().Encode(layer))
		}
		if m.GetContent() != nil {
			x.Bytes(m.GetContent().Encode(layer))
		}
		x.Bytes(m.GetSendMessage().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputBotInlineResult encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputBotInlineResultPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputBotInlineResultPhoto, layer)

	switch cmd {
	case Cmd_InputBotInlineResultPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineResultPhoto))
		x.String(m.GetId())
		x.String(m.GetType())
		x.Bytes(m.GetPhoto().Encode(layer))
		x.Bytes(m.GetSendMessage().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputBotInlineResultPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputBotInlineResultDocument) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputBotInlineResultDocument, layer)

	switch cmd {
	case Cmd_InputBotInlineResultDocument:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineResultDocument))
		// flags
		var flags int32 = 0
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetId())
		x.String(m.GetType())
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		x.Bytes(m.GetDocument().Encode(layer))
		x.Bytes(m.GetSendMessage().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputBotInlineResultDocument encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputBotInlineResultGame) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputBotInlineResultGame, layer)

	switch cmd {
	case Cmd_InputBotInlineResultGame:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputBotInlineResultGame))
		x.String(m.GetId())
		x.String(m.GetShortName())
		x.Bytes(m.GetSendMessage().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputBotInlineResultGame encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputChannelEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputChannelEmpty, layer)

	switch cmd {
	case Cmd_InputChannelEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputChannelEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputChannelEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputChannel) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputChannel, layer)

	switch cmd {
	case Cmd_InputChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputChannel))
		x.Int(m.GetChannelId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputChannel encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputChannelFromMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputChannelFromMessage, layer)

	switch cmd {
	case Cmd_InputChannelFromMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputChannelFromMessage))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetMsgId())
		x.Int(m.GetChannelId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputChannelFromMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputChatPhotoEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputChatPhotoEmpty, layer)

	switch cmd {
	case Cmd_InputChatPhotoEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputChatPhotoEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputChatPhotoEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputChatUploadedPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputChatUploadedPhoto, layer)

	switch cmd {
	case Cmd_InputChatUploadedPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputChatUploadedPhoto))
		x.Bytes(m.GetFile().Encode(layer))

		return x.GetBuf()
	case Cmd_InputChatUploadedPhoto94254732:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputChatUploadedPhoto94254732))
		x.Bytes(m.GetFile().Encode(layer))
		x.Bytes(m.GetCrop().Encode(layer))

		return x.GetBuf()
	case Cmd_InputChatUploadedPhotoc642724e:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputChatUploadedPhotoc642724e))
		// flags
		var flags int32 = 0
		if m.GetFile() != nil {
			flags |= 1 << 0
		}
		if m.GetVideo() != nil {
			flags |= 1 << 1
		}
		if doubleIsEqual(m.GetVideoStartTs(), 0, 0.00000001) == true {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.GetFile() != nil {
			x.Bytes(m.GetFile().Encode(layer))
		}
		if m.GetVideo() != nil {
			x.Bytes(m.GetVideo().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputChatUploadedPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputChatPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputChatPhoto, layer)

	switch cmd {
	case Cmd_InputChatPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputChatPhoto))
		x.Bytes(m.GetId().Encode(layer))

		return x.GetBuf()
	case Cmd_InputChatPhotob2e1bf08:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputChatPhotob2e1bf08))
		x.Bytes(m.GetId().Encode(layer))
		x.Bytes(m.GetCrop().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputChatPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputCheckPasswordEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputCheckPasswordEmpty, layer)

	switch cmd {
	case Cmd_InputCheckPasswordEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputCheckPasswordEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputCheckPasswordEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputCheckPasswordSRP) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputCheckPasswordSRP, layer)

	switch cmd {
	case Cmd_InputCheckPasswordSRP:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputCheckPasswordSRP))
		x.Long(m.GetSrpId())
		x.StringBytes(m.GetA())
		x.StringBytes(m.GetM1())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputCheckPasswordSRP encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputClientProxy) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputClientProxy, layer)

	switch cmd {
	case Cmd_InputClientProxy:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputClientProxy))
		x.String(m.GetAddress())
		x.Int(m.GetPort())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputClientProxy encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPhoneContact) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPhoneContact, layer)

	switch cmd {
	case Cmd_InputPhoneContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPhoneContact))
		x.Long(m.GetClientId())
		x.String(m.GetPhone())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPhoneContact encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputDialogPeer) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputDialogPeer, layer)

	switch cmd {
	case Cmd_InputDialogPeer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputDialogPeer))
		x.Bytes(m.GetPeer().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputDialogPeer encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputDialogPeerFolder) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputDialogPeerFolder, layer)

	switch cmd {
	case Cmd_InputDialogPeerFolder:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputDialogPeerFolder))
		x.Int(m.GetFolderId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputDialogPeerFolder encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputDocumentEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputDocumentEmpty, layer)

	switch cmd {
	case Cmd_InputDocumentEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputDocumentEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputDocumentEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputDocument) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputDocument, layer)

	switch cmd {
	case Cmd_InputDocument:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputDocument))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()
	case Cmd_InputDocument1abfb575:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputDocument1abfb575))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputDocument encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputEncryptedChat) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputEncryptedChat, layer)

	switch cmd {
	case Cmd_InputEncryptedChat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputEncryptedChat))
		x.Int(m.GetChatId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputEncryptedChat encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputEncryptedFileEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputEncryptedFileEmpty, layer)

	switch cmd {
	case Cmd_InputEncryptedFileEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputEncryptedFileEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputEncryptedFileEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputEncryptedFileUploaded) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputEncryptedFileUploaded, layer)

	switch cmd {
	case Cmd_InputEncryptedFileUploaded:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputEncryptedFileUploaded))
		x.Long(m.GetId())
		x.Int(m.GetParts())
		x.String(m.GetMd5Checksum())
		x.Int(m.GetKeyFingerprint())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputEncryptedFileUploaded encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputEncryptedFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputEncryptedFile, layer)

	switch cmd {
	case Cmd_InputEncryptedFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputEncryptedFile))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputEncryptedFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputEncryptedFileBigUploaded) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputEncryptedFileBigUploaded, layer)

	switch cmd {
	case Cmd_InputEncryptedFileBigUploaded:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputEncryptedFileBigUploaded))
		x.Long(m.GetId())
		x.Int(m.GetParts())
		x.Int(m.GetKeyFingerprint())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputEncryptedFileBigUploaded encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputFile, layer)

	switch cmd {
	case Cmd_InputFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputFile))
		x.Long(m.GetId())
		x.Int(m.GetParts())
		x.String(m.GetName())
		x.String(m.GetMd5Checksum())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputFileBig) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputFileBig, layer)

	switch cmd {
	case Cmd_InputFileBig:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputFileBig))
		x.Long(m.GetId())
		x.Int(m.GetParts())
		x.String(m.GetName())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputFileBig encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputFileLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputFileLocation, layer)

	switch cmd {
	case Cmd_InputFileLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputFileLocation))
		x.Long(m.GetVolumeId())
		x.Int(m.GetLocalId())
		x.Long(m.GetSecret())

		return x.GetBuf()
	case Cmd_InputFileLocationdfdaabe1:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputFileLocationdfdaabe1))
		x.Long(m.GetVolumeId())
		x.Int(m.GetLocalId())
		x.Long(m.GetSecret())
		x.StringBytes(m.GetFileReference())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputFileLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputEncryptedFileLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputEncryptedFileLocation, layer)

	switch cmd {
	case Cmd_InputEncryptedFileLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputEncryptedFileLocation))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputEncryptedFileLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputDocumentFileLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputDocumentFileLocation, layer)

	switch cmd {
	case Cmd_InputDocumentFileLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputDocumentFileLocation))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetVersion())

		return x.GetBuf()
	case Cmd_InputDocumentFileLocation196683d9:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputDocumentFileLocation196683d9))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())

		return x.GetBuf()
	case Cmd_InputDocumentFileLocation4e45abe9:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputDocumentFileLocation4e45abe9))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()
	case Cmd_InputDocumentFileLocationbad07584:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputDocumentFileLocationbad07584))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())
		x.String(m.GetThumbSize())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputDocumentFileLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputSecureFileLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputSecureFileLocation, layer)

	switch cmd {
	case Cmd_InputSecureFileLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputSecureFileLocation))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputSecureFileLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputTakeoutFileLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputTakeoutFileLocation, layer)

	switch cmd {
	case Cmd_InputTakeoutFileLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputTakeoutFileLocation))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputTakeoutFileLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPhotoFileLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPhotoFileLocation, layer)

	switch cmd {
	case Cmd_InputPhotoFileLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPhotoFileLocation))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())
		x.String(m.GetThumbSize())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPhotoFileLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPeerPhotoFileLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPeerPhotoFileLocation, layer)

	switch cmd {
	case Cmd_InputPeerPhotoFileLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerPhotoFileLocation))
		// flags
		var flags int32 = 0
		if m.GetBig() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetPeer().Encode(layer))
		x.Long(m.GetVolumeId())
		x.Int(m.GetLocalId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPeerPhotoFileLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputStickerSetThumb) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputStickerSetThumb, layer)

	switch cmd {
	case Cmd_InputStickerSetThumb:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputStickerSetThumb))
		x.Bytes(m.GetStickerset().Encode(layer))
		x.Long(m.GetVolumeId())
		x.Int(m.GetLocalId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputStickerSetThumb encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPhotoLegacyFileLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPhotoLegacyFileLocation, layer)

	switch cmd {
	case Cmd_InputPhotoLegacyFileLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPhotoLegacyFileLocation))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())
		x.Long(m.GetVolumeId())
		x.Int(m.GetLocalId())
		x.Long(m.GetSecret())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPhotoLegacyFileLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputFolderPeer) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputFolderPeer, layer)

	switch cmd {
	case Cmd_InputFolderPeer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputFolderPeer))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetFolderId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputFolderPeer encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputGameID) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputGameID, layer)

	switch cmd {
	case Cmd_InputGameID:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputGameID))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputGameID encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputGameShortName) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputGameShortName, layer)

	switch cmd {
	case Cmd_InputGameShortName:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputGameShortName))
		x.Bytes(m.GetBotId().Encode(layer))
		x.String(m.GetShortName())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputGameShortName encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputGeoPointEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputGeoPointEmpty, layer)

	switch cmd {
	case Cmd_InputGeoPointEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputGeoPointEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputGeoPointEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputGeoPoint) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputGeoPoint, layer)

	switch cmd {
	case Cmd_InputGeoPoint:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputGeoPoint))
		x.Double(m.GetLat())
		x.Double(m.GetLong())

		return x.GetBuf()
	case Cmd_InputGeoPoint48222faf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputGeoPoint48222faf))
		// flags
		var flags int32 = 0
		if m.GetAccuracyRadius() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Double(m.GetLat())
		x.Double(m.GetLong())
		if m.GetAccuracyRadius() != 0 {
			x.Int(m.GetAccuracyRadius())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputGeoPoint encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputGroupCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputGroupCall, layer)

	switch cmd {
	case Cmd_InputGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputGroupCall))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputGroupCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaEmpty, layer)

	switch cmd {
	case Cmd_InputMediaEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaUploadedPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaUploadedPhoto, layer)

	switch cmd {
	case Cmd_InputMediaUploadedPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaUploadedPhoto))
		// flags
		var flags int32 = 0
		if len(m.GetStickers()) > 0 {
			flags |= 1 << 0
		}
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetFile().Encode(layer))
		x.String(m.GetCaption())
		if len(m.GetStickers()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetStickers())))
			for _, v := range m.GetStickers() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_InputMediaUploadedPhoto1e287d04:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaUploadedPhoto1e287d04))
		// flags
		var flags int32 = 0
		if len(m.GetStickers()) > 0 {
			flags |= 1 << 0
		}
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetFile().Encode(layer))
		if len(m.GetStickers()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetStickers())))
			for _, v := range m.GetStickers() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_InputMediaUploadedPhotof7aff1c0:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaUploadedPhotof7aff1c0))
		x.Bytes(m.GetFile().Encode(layer))
		x.String(m.GetCaption())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaUploadedPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaPhoto, layer)

	switch cmd {
	case Cmd_InputMediaPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaPhoto))
		// flags
		var flags int32 = 0
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetId81FA373A71().Encode(layer))
		x.String(m.GetCaption())
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_InputMediaPhotob3ba0635:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaPhotob3ba0635))
		// flags
		var flags int32 = 0
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetId81FA373A71().Encode(layer))
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_InputMediaPhotoe9bfb4f3:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaPhotoe9bfb4f3))
		x.Bytes(m.GetId81FA373A71().Encode(layer))
		x.String(m.GetCaption())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaGeoPoint) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaGeoPoint, layer)

	switch cmd {
	case Cmd_InputMediaGeoPoint:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaGeoPoint))
		x.Bytes(m.GetGeoPoint().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaGeoPoint encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaContact) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaContact, layer)

	switch cmd {
	case Cmd_InputMediaContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaContact))
		x.String(m.GetPhoneNumber())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())

		return x.GetBuf()
	case Cmd_InputMediaContactf8ab7dfb:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaContactf8ab7dfb))
		x.String(m.GetPhoneNumber())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())
		x.String(m.GetVcard())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaContact encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaUploadedDocument) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaUploadedDocument, layer)

	switch cmd {
	case Cmd_InputMediaUploadedDocument:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaUploadedDocument))
		// flags
		var flags int32 = 0
		if m.GetThumb() != nil {
			flags |= 1 << 2
		}
		if len(m.GetStickers()) > 0 {
			flags |= 1 << 0
		}
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetFile().Encode(layer))
		if m.GetThumb() != nil {
			x.Bytes(m.GetThumb().Encode(layer))
		}
		x.String(m.GetMimeType())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.String(m.GetCaption())
		if len(m.GetStickers()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetStickers())))
			for _, v := range m.GetStickers() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_InputMediaUploadedDocument1d89306d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaUploadedDocument1d89306d))
		x.Bytes(m.GetFile().Encode(layer))
		x.String(m.GetMimeType())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.String(m.GetCaption())

		return x.GetBuf()
	case Cmd_InputMediaUploadedDocument5b38c6c1:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaUploadedDocument5b38c6c1))
		// flags
		var flags int32 = 0
		if m.GetNosoundVideo() == true {
			flags |= 1 << 3
		}
		if m.GetForceFile() == true {
			flags |= 1 << 4
		}
		if m.GetThumb() != nil {
			flags |= 1 << 2
		}
		if len(m.GetStickers()) > 0 {
			flags |= 1 << 0
		}
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetFile().Encode(layer))
		if m.GetThumb() != nil {
			x.Bytes(m.GetThumb().Encode(layer))
		}
		x.String(m.GetMimeType())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if len(m.GetStickers()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetStickers())))
			for _, v := range m.GetStickers() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaUploadedDocument encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaDocument) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaDocument, layer)

	switch cmd {
	case Cmd_InputMediaDocument:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaDocument))
		// flags
		var flags int32 = 0
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetId5ACB668E71().Encode(layer))
		x.String(m.GetCaption())
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_InputMediaDocument1a77f29c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaDocument1a77f29c))
		x.Bytes(m.GetId5ACB668E71().Encode(layer))
		x.String(m.GetCaption())

		return x.GetBuf()
	case Cmd_InputMediaDocument23ab23d2:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaDocument23ab23d2))
		// flags
		var flags int32 = 0
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetId5ACB668E71().Encode(layer))
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_InputMediaDocument33473058:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaDocument33473058))
		// flags
		var flags int32 = 0
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetQuery()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetId5ACB668E71().Encode(layer))
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}
		if len(m.GetQuery()) > 0 {
			x.String(m.GetQuery())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaDocument encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaVenue) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaVenue, layer)

	switch cmd {
	case Cmd_InputMediaVenue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaVenue))
		x.Bytes(m.GetGeoPoint().Encode(layer))
		x.String(m.GetTitle())
		x.String(m.GetAddress())
		x.String(m.GetProvider())
		x.String(m.GetVenueId())

		return x.GetBuf()
	case Cmd_InputMediaVenuec13d1c11:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaVenuec13d1c11))
		x.Bytes(m.GetGeoPoint().Encode(layer))
		x.String(m.GetTitle())
		x.String(m.GetAddress())
		x.String(m.GetProvider())
		x.String(m.GetVenueId())
		x.String(m.GetVenueType())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaVenue encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaGifExternal) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaGifExternal, layer)

	switch cmd {
	case Cmd_InputMediaGifExternal:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaGifExternal))
		x.String(m.GetUrl())
		x.String(m.GetQ())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaGifExternal encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaPhotoExternal) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaPhotoExternal, layer)

	switch cmd {
	case Cmd_InputMediaPhotoExternal:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaPhotoExternal))
		// flags
		var flags int32 = 0
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetUrl())
		x.String(m.GetCaption())
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_InputMediaPhotoExternale5bbfe1a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaPhotoExternale5bbfe1a))
		// flags
		var flags int32 = 0
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetUrl())
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaPhotoExternal encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaDocumentExternal) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaDocumentExternal, layer)

	switch cmd {
	case Cmd_InputMediaDocumentExternal:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaDocumentExternal))
		// flags
		var flags int32 = 0
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetUrl())
		x.String(m.GetCaption())
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_InputMediaDocumentExternalfb52dc99:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaDocumentExternalfb52dc99))
		// flags
		var flags int32 = 0
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetUrl())
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaDocumentExternal encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaGame) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaGame, layer)

	switch cmd {
	case Cmd_InputMediaGame:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaGame))
		x.Bytes(m.GetIdD33F43F371().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaGame encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaInvoice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaInvoice, layer)

	switch cmd {
	case Cmd_InputMediaInvoice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaInvoice))
		// flags
		var flags int32 = 0
		if m.GetPhoto() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetTitle())
		x.String(m.GetDescription())
		if m.GetPhoto() != nil {
			x.Bytes(m.GetPhoto().Encode(layer))
		}
		x.Bytes(m.GetInvoice().Encode(layer))
		x.StringBytes(m.GetPayload())
		x.String(m.GetProvider())
		x.String(m.GetStartParam())

		return x.GetBuf()
	case Cmd_InputMediaInvoicef4e096c3:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaInvoicef4e096c3))
		// flags
		var flags int32 = 0
		if m.GetPhoto() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetTitle())
		x.String(m.GetDescription())
		if m.GetPhoto() != nil {
			x.Bytes(m.GetPhoto().Encode(layer))
		}
		x.Bytes(m.GetInvoice().Encode(layer))
		x.StringBytes(m.GetPayload())
		x.String(m.GetProvider())
		x.Bytes(m.GetProviderData().Encode(layer))
		x.String(m.GetStartParam())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaInvoice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaUploadedThumbDocument) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaUploadedThumbDocument, layer)

	switch cmd {
	case Cmd_InputMediaUploadedThumbDocument:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaUploadedThumbDocument))
		x.Bytes(m.GetFile().Encode(layer))
		x.Bytes(m.GetThumb().Encode(layer))
		x.String(m.GetMimeType())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.String(m.GetCaption())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaUploadedThumbDocument encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaGeoLive) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaGeoLive, layer)

	switch cmd {
	case Cmd_InputMediaGeoLive:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaGeoLive))
		x.Bytes(m.GetGeoPoint().Encode(layer))
		x.Int(m.GetPeriod())

		return x.GetBuf()
	case Cmd_InputMediaGeoLive971fa843:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaGeoLive971fa843))
		// flags
		var flags int32 = 0
		if m.GetStopped() == true {
			flags |= 1 << 0
		}
		if m.GetHeading() != 0 {
			flags |= 1 << 2
		}
		if m.GetPeriod() != 0 {
			flags |= 1 << 1
		}
		if m.GetProximityNotificationRadius() != 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Bytes(m.GetGeoPoint().Encode(layer))
		if m.GetHeading() != 0 {
			x.Int(m.GetHeading())
		}
		if m.GetPeriod() != 0 {
			x.Int(m.GetPeriod())
		}
		if m.GetProximityNotificationRadius() != 0 {
			x.Int(m.GetProximityNotificationRadius())
		}

		return x.GetBuf()
	case Cmd_InputMediaGeoLivece4e82fd:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaGeoLivece4e82fd))
		// flags
		var flags int32 = 0
		if m.GetStopped() == true {
			flags |= 1 << 0
		}
		if m.GetPeriod() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetGeoPoint().Encode(layer))
		if m.GetPeriod() != 0 {
			x.Int(m.GetPeriod())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaGeoLive encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaPoll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaPoll, layer)

	switch cmd {
	case Cmd_InputMediaPoll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaPoll))
		x.Bytes(m.GetPoll().Encode(layer))

		return x.GetBuf()
	case Cmd_InputMediaPollabe9ca25:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaPollabe9ca25))
		// flags
		var flags int32 = 0
		if len(m.GetCorrectAnswers()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetPoll().Encode(layer))
		if len(m.GetCorrectAnswers()) > 0 {

		}

		return x.GetBuf()
	case Cmd_InputMediaPollf94e5f1:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaPollf94e5f1))
		// flags
		var flags int32 = 0
		if len(m.GetCorrectAnswers()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetSolution()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetSolutionEntities()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetPoll().Encode(layer))
		if len(m.GetCorrectAnswers()) > 0 {

		}
		if len(m.GetSolution()) > 0 {
			x.String(m.GetSolution())
		}
		if len(m.GetSolutionEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetSolutionEntities())))
			for _, v := range m.GetSolutionEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaPoll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMediaDice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMediaDice, layer)

	switch cmd {
	case Cmd_InputMediaDice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaDice))
		_ = m

		return x.GetBuf()
	case Cmd_InputMediaDicee66fbf7b:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMediaDicee66fbf7b))
		x.String(m.GetEmoticon())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMediaDice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessageID) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessageID, layer)

	switch cmd {
	case Cmd_InputMessageID:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessageID))
		x.Int(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessageID encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessageReplyTo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessageReplyTo, layer)

	switch cmd {
	case Cmd_InputMessageReplyTo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessageReplyTo))
		x.Int(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessageReplyTo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagePinned) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagePinned, layer)

	switch cmd {
	case Cmd_InputMessagePinned:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagePinned))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagePinned encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessageCallbackQuery) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessageCallbackQuery, layer)

	switch cmd {
	case Cmd_InputMessageCallbackQuery:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessageCallbackQuery))
		x.Int(m.GetId())
		x.Long(m.GetQueryId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessageCallbackQuery encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputNotifyPeer) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputNotifyPeer, layer)

	switch cmd {
	case Cmd_InputNotifyPeer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputNotifyPeer))
		x.Bytes(m.GetPeer().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputNotifyPeer encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputNotifyUsers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputNotifyUsers, layer)

	switch cmd {
	case Cmd_InputNotifyUsers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputNotifyUsers))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputNotifyUsers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputNotifyChats) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputNotifyChats, layer)

	switch cmd {
	case Cmd_InputNotifyChats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputNotifyChats))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputNotifyChats encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputNotifyAll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputNotifyAll, layer)

	switch cmd {
	case Cmd_InputNotifyAll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputNotifyAll))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputNotifyAll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputNotifyBroadcasts) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputNotifyBroadcasts, layer)

	switch cmd {
	case Cmd_InputNotifyBroadcasts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputNotifyBroadcasts))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputNotifyBroadcasts encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPaymentCredentialsSaved) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPaymentCredentialsSaved, layer)

	switch cmd {
	case Cmd_InputPaymentCredentialsSaved:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPaymentCredentialsSaved))
		x.String(m.GetId())
		x.StringBytes(m.GetTmpPassword())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPaymentCredentialsSaved encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPaymentCredentials) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPaymentCredentials, layer)

	switch cmd {
	case Cmd_InputPaymentCredentials:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPaymentCredentials))
		// flags
		var flags int32 = 0
		if m.GetSave() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetData().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPaymentCredentials encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPaymentCredentialsApplePay) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPaymentCredentialsApplePay, layer)

	switch cmd {
	case Cmd_InputPaymentCredentialsApplePay:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPaymentCredentialsApplePay))
		x.Bytes(m.GetPaymentData().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPaymentCredentialsApplePay encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPaymentCredentialsAndroidPay) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPaymentCredentialsAndroidPay, layer)

	switch cmd {
	case Cmd_InputPaymentCredentialsAndroidPay:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPaymentCredentialsAndroidPay))
		x.Bytes(m.GetPaymentToken().Encode(layer))
		x.String(m.GetGoogleTransactionId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPaymentCredentialsAndroidPay encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPeerEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPeerEmpty, layer)

	switch cmd {
	case Cmd_InputPeerEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPeerEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPeerSelf) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPeerSelf, layer)

	switch cmd {
	case Cmd_InputPeerSelf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerSelf))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPeerSelf encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPeerChat) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPeerChat, layer)

	switch cmd {
	case Cmd_InputPeerChat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerChat))
		x.Int(m.GetChatId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPeerChat encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPeerUser) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPeerUser, layer)

	switch cmd {
	case Cmd_InputPeerUser:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerUser))
		x.Int(m.GetUserId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPeerUser encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPeerChannel) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPeerChannel, layer)

	switch cmd {
	case Cmd_InputPeerChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerChannel))
		x.Int(m.GetChannelId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPeerChannel encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPeerUserFromMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPeerUserFromMessage, layer)

	switch cmd {
	case Cmd_InputPeerUserFromMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerUserFromMessage))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetMsgId())
		x.Int(m.GetUserId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPeerUserFromMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPeerChannelFromMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPeerChannelFromMessage, layer)

	switch cmd {
	case Cmd_InputPeerChannelFromMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerChannelFromMessage))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetMsgId())
		x.Int(m.GetChannelId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPeerChannelFromMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPeerNotifyEventsEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPeerNotifyEventsEmpty, layer)

	switch cmd {
	case Cmd_InputPeerNotifyEventsEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerNotifyEventsEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPeerNotifyEventsEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPeerNotifyEventsAll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPeerNotifyEventsAll, layer)

	switch cmd {
	case Cmd_InputPeerNotifyEventsAll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerNotifyEventsAll))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPeerNotifyEventsAll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPeerNotifySettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPeerNotifySettings, layer)

	switch cmd {
	case Cmd_InputPeerNotifySettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerNotifySettings))
		// flags
		var flags int32 = 0
		if m.GetShowPreviews38935EB271() == true {
			flags |= 1 << 0
		}
		if m.GetSilent38935EB271() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetMuteUntil())
		x.String(m.GetSound())

		return x.GetBuf()
	case Cmd_InputPeerNotifySettings9c3d198e:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPeerNotifySettings9c3d198e))
		// flags
		var flags int32 = 0
		if m.GetShowPreviews9C3D198E85() != nil {
			flags |= 1 << 0
		}
		if m.GetSilent9C3D198E85() != nil {
			flags |= 1 << 1
		}
		if m.GetMuteUntil() != 0 {
			flags |= 1 << 2
		}
		if len(m.GetSound()) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		if m.GetShowPreviews9C3D198E85() != nil {
			x.Bytes(m.GetShowPreviews9C3D198E85().Encode(layer))
		}
		if m.GetSilent9C3D198E85() != nil {
			x.Bytes(m.GetSilent9C3D198E85().Encode(layer))
		}
		if m.GetMuteUntil() != 0 {
			x.Int(m.GetMuteUntil())
		}
		if len(m.GetSound()) > 0 {
			x.String(m.GetSound())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPeerNotifySettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPhoneCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPhoneCall, layer)

	switch cmd {
	case Cmd_InputPhoneCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPhoneCall))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPhoneCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPhotoEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPhotoEmpty, layer)

	switch cmd {
	case Cmd_InputPhotoEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPhotoEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPhotoEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPhoto, layer)

	switch cmd {
	case Cmd_InputPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPhoto))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()
	case Cmd_InputPhoto3bb3b94a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPhoto3bb3b94a))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPhotoCrop) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPhotoCrop, layer)

	switch cmd {
	case Cmd_InputPhotoCrop:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPhotoCrop))
		x.Double(m.GetCropLeft())
		x.Double(m.GetCropTop())
		x.Double(m.GetCropWidth())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPhotoCrop encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPhotoCropAuto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPhotoCropAuto, layer)

	switch cmd {
	case Cmd_InputPhotoCropAuto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPhotoCropAuto))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPhotoCropAuto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyKeyStatusTimestamp) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyKeyStatusTimestamp, layer)

	switch cmd {
	case Cmd_InputPrivacyKeyStatusTimestamp:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyKeyStatusTimestamp))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyKeyStatusTimestamp encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyKeyChatInvite) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyKeyChatInvite, layer)

	switch cmd {
	case Cmd_InputPrivacyKeyChatInvite:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyKeyChatInvite))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyKeyChatInvite encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyKeyPhoneCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyKeyPhoneCall, layer)

	switch cmd {
	case Cmd_InputPrivacyKeyPhoneCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyKeyPhoneCall))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyKeyPhoneCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyKeyPhoneP2P) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyKeyPhoneP2P, layer)

	switch cmd {
	case Cmd_InputPrivacyKeyPhoneP2P:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyKeyPhoneP2P))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyKeyPhoneP2P encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyKeyForwards) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyKeyForwards, layer)

	switch cmd {
	case Cmd_InputPrivacyKeyForwards:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyKeyForwards))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyKeyForwards encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyKeyProfilePhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyKeyProfilePhoto, layer)

	switch cmd {
	case Cmd_InputPrivacyKeyProfilePhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyKeyProfilePhoto))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyKeyProfilePhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyKeyPhoneNumber) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyKeyPhoneNumber, layer)

	switch cmd {
	case Cmd_InputPrivacyKeyPhoneNumber:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyKeyPhoneNumber))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyKeyPhoneNumber encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyKeyAddedByPhone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyKeyAddedByPhone, layer)

	switch cmd {
	case Cmd_InputPrivacyKeyAddedByPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyKeyAddedByPhone))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyKeyAddedByPhone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyValueAllowContacts) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyValueAllowContacts, layer)

	switch cmd {
	case Cmd_InputPrivacyValueAllowContacts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyValueAllowContacts))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyValueAllowContacts encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyValueAllowAll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyValueAllowAll, layer)

	switch cmd {
	case Cmd_InputPrivacyValueAllowAll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyValueAllowAll))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyValueAllowAll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyValueAllowUsers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyValueAllowUsers, layer)

	switch cmd {
	case Cmd_InputPrivacyValueAllowUsers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyValueAllowUsers))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyValueAllowUsers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyValueDisallowContacts) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyValueDisallowContacts, layer)

	switch cmd {
	case Cmd_InputPrivacyValueDisallowContacts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyValueDisallowContacts))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyValueDisallowContacts encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyValueDisallowAll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyValueDisallowAll, layer)

	switch cmd {
	case Cmd_InputPrivacyValueDisallowAll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyValueDisallowAll))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyValueDisallowAll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyValueDisallowUsers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyValueDisallowUsers, layer)

	switch cmd {
	case Cmd_InputPrivacyValueDisallowUsers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyValueDisallowUsers))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyValueDisallowUsers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyValueAllowChatParticipants) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyValueAllowChatParticipants, layer)

	switch cmd {
	case Cmd_InputPrivacyValueAllowChatParticipants:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyValueAllowChatParticipants))
		x.VectorInt(m.GetChats())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyValueAllowChatParticipants encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputPrivacyValueDisallowChatParticipants) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputPrivacyValueDisallowChatParticipants, layer)

	switch cmd {
	case Cmd_InputPrivacyValueDisallowChatParticipants:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputPrivacyValueDisallowChatParticipants))
		x.VectorInt(m.GetChats())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputPrivacyValueDisallowChatParticipants encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputSecureFileUploaded) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputSecureFileUploaded, layer)

	switch cmd {
	case Cmd_InputSecureFileUploaded:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputSecureFileUploaded))
		x.Long(m.GetId())
		x.Int(m.GetParts())
		x.String(m.GetMd5Checksum())
		x.StringBytes(m.GetFileHash())
		x.StringBytes(m.GetSecret())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputSecureFileUploaded encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputSecureFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputSecureFile, layer)

	switch cmd {
	case Cmd_InputSecureFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputSecureFile))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputSecureFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputSecureValue) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputSecureValue, layer)

	switch cmd {
	case Cmd_InputSecureValue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputSecureValue))
		// flags
		var flags int32 = 0
		if m.GetData() != nil {
			flags |= 1 << 0
		}
		if m.GetFrontSide() != nil {
			flags |= 1 << 1
		}
		if m.GetReverseSide() != nil {
			flags |= 1 << 2
		}
		if m.GetSelfie() != nil {
			flags |= 1 << 3
		}
		if len(m.GetTranslation()) > 0 {
			flags |= 1 << 6
		}
		if len(m.GetFiles()) > 0 {
			flags |= 1 << 4
		}
		if m.GetPlainData() != nil {
			flags |= 1 << 5
		}
		x.Int(flags)

		x.Bytes(m.GetType().Encode(layer))
		if m.GetData() != nil {
			x.Bytes(m.GetData().Encode(layer))
		}
		if m.GetFrontSide() != nil {
			x.Bytes(m.GetFrontSide().Encode(layer))
		}
		if m.GetReverseSide() != nil {
			x.Bytes(m.GetReverseSide().Encode(layer))
		}
		if m.GetSelfie() != nil {
			x.Bytes(m.GetSelfie().Encode(layer))
		}
		if len(m.GetTranslation()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetTranslation())))
			for _, v := range m.GetTranslation() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if len(m.GetFiles()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetFiles())))
			for _, v := range m.GetFiles() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetPlainData() != nil {
			x.Bytes(m.GetPlainData().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputSecureValue encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputSingleMedia) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputSingleMedia, layer)

	switch cmd {
	case Cmd_InputSingleMedia:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputSingleMedia))
		// flags
		var flags int32 = 0
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetMedia().Encode(layer))
		x.Long(m.GetRandomId())
		x.String(m.GetMessage())
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputSingleMedia encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputStickerSetEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputStickerSetEmpty, layer)

	switch cmd {
	case Cmd_InputStickerSetEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputStickerSetEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputStickerSetEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputStickerSetID) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputStickerSetID, layer)

	switch cmd {
	case Cmd_InputStickerSetID:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputStickerSetID))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputStickerSetID encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputStickerSetShortName) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputStickerSetShortName, layer)

	switch cmd {
	case Cmd_InputStickerSetShortName:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputStickerSetShortName))
		x.String(m.GetShortName())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputStickerSetShortName encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputStickerSetAnimatedEmoji) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputStickerSetAnimatedEmoji, layer)

	switch cmd {
	case Cmd_InputStickerSetAnimatedEmoji:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputStickerSetAnimatedEmoji))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputStickerSetAnimatedEmoji encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputStickerSetDice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputStickerSetDice, layer)

	switch cmd {
	case Cmd_InputStickerSetDice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputStickerSetDice))
		_ = m

		return x.GetBuf()
	case Cmd_InputStickerSetDicee67f520e:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputStickerSetDicee67f520e))
		x.String(m.GetEmoticon())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputStickerSetDice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputStickerSetItem) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputStickerSetItem, layer)

	switch cmd {
	case Cmd_InputStickerSetItem:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputStickerSetItem))
		// flags
		var flags int32 = 0
		if m.GetMaskCoords() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetDocument().Encode(layer))
		x.String(m.GetEmoji())
		if m.GetMaskCoords() != nil {
			x.Bytes(m.GetMaskCoords().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputStickerSetItem encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputStickeredMediaPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputStickeredMediaPhoto, layer)

	switch cmd {
	case Cmd_InputStickeredMediaPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputStickeredMediaPhoto))
		x.Bytes(m.GetId4A99215771().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputStickeredMediaPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputStickeredMediaDocument) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputStickeredMediaDocument, layer)

	switch cmd {
	case Cmd_InputStickeredMediaDocument:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputStickeredMediaDocument))
		x.Bytes(m.GetId438865B71().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputStickeredMediaDocument encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputTheme) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputTheme, layer)

	switch cmd {
	case Cmd_InputTheme:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputTheme))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputTheme encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputThemeSlug) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputThemeSlug, layer)

	switch cmd {
	case Cmd_InputThemeSlug:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputThemeSlug))
		x.String(m.GetSlug())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputThemeSlug encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputThemeSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputThemeSettings, layer)

	switch cmd {
	case Cmd_InputThemeSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputThemeSettings))
		// flags
		var flags int32 = 0
		if m.GetMessageTopColor() != 0 {
			flags |= 1 << 0
		}
		if m.GetMessageBottomColor() != 0 {
			flags |= 1 << 0
		}
		if m.GetWallpaper() != nil {
			flags |= 1 << 1
		}
		if m.GetWallpaperSettings() != nil {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetBaseTheme().Encode(layer))
		x.Int(m.GetAccentColor())
		if m.GetMessageTopColor() != 0 {
			x.Int(m.GetMessageTopColor())
		}
		if m.GetMessageBottomColor() != 0 {
			x.Int(m.GetMessageBottomColor())
		}
		if m.GetWallpaper() != nil {
			x.Bytes(m.GetWallpaper().Encode(layer))
		}
		if m.GetWallpaperSettings() != nil {
			x.Bytes(m.GetWallpaperSettings().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputThemeSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputUserEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputUserEmpty, layer)

	switch cmd {
	case Cmd_InputUserEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputUserEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputUserEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputUserSelf) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputUserSelf, layer)

	switch cmd {
	case Cmd_InputUserSelf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputUserSelf))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputUserSelf encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputUser) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputUser, layer)

	switch cmd {
	case Cmd_InputUser:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputUser))
		x.Int(m.GetUserId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputUser encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputUserFromMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputUserFromMessage, layer)

	switch cmd {
	case Cmd_InputUserFromMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputUserFromMessage))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetMsgId())
		x.Int(m.GetUserId())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputUserFromMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputWallPaper) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputWallPaper, layer)

	switch cmd {
	case Cmd_InputWallPaper:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputWallPaper))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputWallPaper encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputWallPaperSlug) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputWallPaperSlug, layer)

	switch cmd {
	case Cmd_InputWallPaperSlug:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputWallPaperSlug))
		x.String(m.GetSlug())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputWallPaperSlug encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputWallPaperNoFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputWallPaperNoFile, layer)

	switch cmd {
	case Cmd_InputWallPaperNoFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputWallPaperNoFile))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputWallPaperNoFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputWebDocument) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputWebDocument, layer)

	switch cmd {
	case Cmd_InputWebDocument:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputWebDocument))
		x.String(m.GetUrl())
		x.Int(m.GetSize_())
		x.String(m.GetMimeType())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid InputWebDocument encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputWebFileLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputWebFileLocation, layer)

	switch cmd {
	case Cmd_InputWebFileLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputWebFileLocation))
		x.String(m.GetUrl())
		x.Long(m.GetAccessHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputWebFileLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputWebFileGeoPointLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputWebFileGeoPointLocation, layer)

	switch cmd {
	case Cmd_InputWebFileGeoPointLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputWebFileGeoPointLocation))
		x.Bytes(m.GetGeoPoint().Encode(layer))
		x.Long(m.GetAccessHash())
		x.Int(m.GetW())
		x.Int(m.GetH())
		x.Int(m.GetZoom())
		x.Int(m.GetScale())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputWebFileGeoPointLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInvoice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInvoice, layer)

	switch cmd {
	case Cmd_Invoice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Invoice))
		// flags
		var flags int32 = 0
		if m.GetTest() == true {
			flags |= 1 << 0
		}
		if m.GetNameRequested() == true {
			flags |= 1 << 1
		}
		if m.GetPhoneRequested() == true {
			flags |= 1 << 2
		}
		if m.GetEmailRequested() == true {
			flags |= 1 << 3
		}
		if m.GetShippingAddressRequested() == true {
			flags |= 1 << 4
		}
		if m.GetFlexible() == true {
			flags |= 1 << 5
		}
		x.Int(flags)

		x.String(m.GetCurrency())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPrices())))
		for _, v := range m.GetPrices() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid Invoice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLIpPort) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassIpPort, layer)

	switch cmd {
	case Cmd_IpPort:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_IpPort))
		x.Int(m.GetIpv4())
		x.Int(m.GetPort())

		return x.GetBuf()

	default:
		log.Errorf("invalid IpPort encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLIpPortSecret) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassIpPortSecret, layer)

	switch cmd {
	case Cmd_IpPortSecret:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_IpPortSecret))
		x.Int(m.GetIpv4())
		x.Int(m.GetPort())
		x.StringBytes(m.GetSecret())

		return x.GetBuf()

	default:
		log.Errorf("invalid IpPortSecret encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLJsonObjectValue) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassJsonObjectValue, layer)

	switch cmd {
	case Cmd_JsonObjectValue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_JsonObjectValue))
		x.String(m.GetKey())
		x.Bytes(m.GetValue().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid JsonObjectValue encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLJsonNull) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassJsonNull, layer)

	switch cmd {
	case Cmd_JsonNull:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_JsonNull))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid JsonNull encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLJsonBool) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassJsonBool, layer)

	switch cmd {
	case Cmd_JsonBool:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_JsonBool))
		x.Bytes(m.GetValueC7345E6A90().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid JsonBool encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLJsonNumber) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassJsonNumber, layer)

	switch cmd {
	case Cmd_JsonNumber:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_JsonNumber))
		x.Double(m.GetValue2BE0DFA490())

		return x.GetBuf()

	default:
		log.Errorf("invalid JsonNumber encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLJsonString) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassJsonString, layer)

	switch cmd {
	case Cmd_JsonString:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_JsonString))
		x.String(m.GetValueB71E767A90())

		return x.GetBuf()

	default:
		log.Errorf("invalid JsonString encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLJsonArray) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassJsonArray, layer)

	switch cmd {
	case Cmd_JsonArray:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_JsonArray))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetValueF744476390())))
		for _, v := range m.GetValueF744476390() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid JsonArray encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLJsonObject) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassJsonObject, layer)

	switch cmd {
	case Cmd_JsonObject:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_JsonObject))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetValue99C1D49D90())))
		for _, v := range m.GetValue99C1D49D90() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid JsonObject encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLKeyboardButton) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassKeyboardButton, layer)

	switch cmd {
	case Cmd_KeyboardButton:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButton))
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid KeyboardButton encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLKeyboardButtonUrl) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassKeyboardButtonUrl, layer)

	switch cmd {
	case Cmd_KeyboardButtonUrl:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonUrl))
		x.String(m.GetText())
		x.String(m.GetUrl())

		return x.GetBuf()

	default:
		log.Errorf("invalid KeyboardButtonUrl encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLKeyboardButtonCallback) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassKeyboardButtonCallback, layer)

	switch cmd {
	case Cmd_KeyboardButtonCallback:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonCallback))
		x.String(m.GetText())
		x.StringBytes(m.GetData())

		return x.GetBuf()
	case Cmd_KeyboardButtonCallback35bbdb6b:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonCallback35bbdb6b))
		// flags
		var flags int32 = 0
		if m.GetRequiresPassword() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetText())
		x.StringBytes(m.GetData())

		return x.GetBuf()

	default:
		log.Errorf("invalid KeyboardButtonCallback encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLKeyboardButtonRequestPhone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassKeyboardButtonRequestPhone, layer)

	switch cmd {
	case Cmd_KeyboardButtonRequestPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonRequestPhone))
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid KeyboardButtonRequestPhone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLKeyboardButtonRequestGeoLocation) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassKeyboardButtonRequestGeoLocation, layer)

	switch cmd {
	case Cmd_KeyboardButtonRequestGeoLocation:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonRequestGeoLocation))
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid KeyboardButtonRequestGeoLocation encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLKeyboardButtonSwitchInline) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassKeyboardButtonSwitchInline, layer)

	switch cmd {
	case Cmd_KeyboardButtonSwitchInline:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonSwitchInline))
		// flags
		var flags int32 = 0
		if m.GetSamePeer() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetText())
		x.String(m.GetQuery())

		return x.GetBuf()
	case Cmd_KeyboardButtonSwitchInlineea1b7a14:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonSwitchInlineea1b7a14))
		x.String(m.GetText())
		x.String(m.GetQuery())

		return x.GetBuf()

	default:
		log.Errorf("invalid KeyboardButtonSwitchInline encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLKeyboardButtonGame) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassKeyboardButtonGame, layer)

	switch cmd {
	case Cmd_KeyboardButtonGame:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonGame))
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid KeyboardButtonGame encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLKeyboardButtonBuy) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassKeyboardButtonBuy, layer)

	switch cmd {
	case Cmd_KeyboardButtonBuy:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonBuy))
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid KeyboardButtonBuy encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLKeyboardButtonUrlAuth) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassKeyboardButtonUrlAuth, layer)

	switch cmd {
	case Cmd_KeyboardButtonUrlAuth:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonUrlAuth))
		// flags
		var flags int32 = 0
		if len(m.GetFwdText()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetText())
		if len(m.GetFwdText()) > 0 {
			x.String(m.GetFwdText())
		}
		x.String(m.GetUrl())
		x.Int(m.GetButtonId())

		return x.GetBuf()

	default:
		log.Errorf("invalid KeyboardButtonUrlAuth encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputKeyboardButtonUrlAuth) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputKeyboardButtonUrlAuth, layer)

	switch cmd {
	case Cmd_InputKeyboardButtonUrlAuth:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputKeyboardButtonUrlAuth))
		// flags
		var flags int32 = 0
		if m.GetRequestWriteAccess() == true {
			flags |= 1 << 0
		}
		if len(m.GetFwdText()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.String(m.GetText())
		if len(m.GetFwdText()) > 0 {
			x.String(m.GetFwdText())
		}
		x.String(m.GetUrl())
		x.Bytes(m.GetBot().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputKeyboardButtonUrlAuth encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLKeyboardButtonRequestPoll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassKeyboardButtonRequestPoll, layer)

	switch cmd {
	case Cmd_KeyboardButtonRequestPoll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonRequestPoll))
		// flags
		var flags int32 = 0
		if m.GetQuiz() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetQuiz() != nil {
			x.Bytes(m.GetQuiz().Encode(layer))
		}
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid KeyboardButtonRequestPoll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLKeyboardButtonRow) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassKeyboardButtonRow, layer)

	switch cmd {
	case Cmd_KeyboardButtonRow:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_KeyboardButtonRow))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetButtons())))
		for _, v := range m.GetButtons() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid KeyboardButtonRow encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLLabeledPrice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassLabeledPrice, layer)

	switch cmd {
	case Cmd_LabeledPrice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LabeledPrice))
		x.String(m.GetLabel())
		x.Long(m.GetAmount())

		return x.GetBuf()

	default:
		log.Errorf("invalid LabeledPrice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLLangPackDifference) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassLangPackDifference, layer)

	switch cmd {
	case Cmd_LangPackDifference:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangPackDifference))
		x.String(m.GetLangCode())
		x.Int(m.GetFromVersion())
		x.Int(m.GetVersion())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetStrings())))
		for _, v := range m.GetStrings() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid LangPackDifference encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLLangPackLanguage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassLangPackLanguage, layer)

	switch cmd {
	case Cmd_LangPackLanguage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangPackLanguage))
		x.String(m.GetName())
		x.String(m.GetNativeName())
		x.String(m.GetLangCode())

		return x.GetBuf()
	case Cmd_LangPackLanguage651b98d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangPackLanguage651b98d))
		// flags
		var flags int32 = 0
		if m.GetOfficial() == true {
			flags |= 1 << 0
		}
		if m.GetRtl() == true {
			flags |= 1 << 2
		}
		if len(m.GetBaseLangCode()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.String(m.GetName())
		x.String(m.GetNativeName())
		x.String(m.GetLangCode())
		if len(m.GetBaseLangCode()) > 0 {
			x.String(m.GetBaseLangCode())
		}
		x.String(m.GetPluralCode())
		x.Int(m.GetStringsCount())
		x.Int(m.GetTranslatedCount())

		return x.GetBuf()
	case Cmd_LangPackLanguageeeca5ce3:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangPackLanguageeeca5ce3))
		// flags
		var flags int32 = 0
		if m.GetOfficial() == true {
			flags |= 1 << 0
		}
		if m.GetRtl() == true {
			flags |= 1 << 2
		}
		if len(m.GetBaseLangCode()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.String(m.GetName())
		x.String(m.GetNativeName())
		x.String(m.GetLangCode())
		if len(m.GetBaseLangCode()) > 0 {
			x.String(m.GetBaseLangCode())
		}
		x.String(m.GetPluralCode())
		x.Int(m.GetStringsCount())
		x.Int(m.GetTranslatedCount())
		x.String(m.GetTranslationsUrl())

		return x.GetBuf()

	default:
		log.Errorf("invalid LangPackLanguage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLLangPackString) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassLangPackString, layer)

	switch cmd {
	case Cmd_LangPackString:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangPackString))
		x.String(m.GetKey())
		x.String(m.GetValue())

		return x.GetBuf()

	default:
		log.Errorf("invalid LangPackString encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLLangPackStringPluralized) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassLangPackStringPluralized, layer)

	switch cmd {
	case Cmd_LangPackStringPluralized:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangPackStringPluralized))
		// flags
		var flags int32 = 0
		if len(m.GetZeroValue()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetOneValue()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetTwoValue()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetFewValue()) > 0 {
			flags |= 1 << 3
		}
		if len(m.GetManyValue()) > 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		x.String(m.GetKey())
		if len(m.GetZeroValue()) > 0 {
			x.String(m.GetZeroValue())
		}
		if len(m.GetOneValue()) > 0 {
			x.String(m.GetOneValue())
		}
		if len(m.GetTwoValue()) > 0 {
			x.String(m.GetTwoValue())
		}
		if len(m.GetFewValue()) > 0 {
			x.String(m.GetFewValue())
		}
		if len(m.GetManyValue()) > 0 {
			x.String(m.GetManyValue())
		}
		x.String(m.GetOtherValue())

		return x.GetBuf()

	default:
		log.Errorf("invalid LangPackStringPluralized encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLLangPackStringDeleted) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassLangPackStringDeleted, layer)

	switch cmd {
	case Cmd_LangPackStringDeleted:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_LangPackStringDeleted))
		x.String(m.GetKey())

		return x.GetBuf()

	default:
		log.Errorf("invalid LangPackStringDeleted encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMaskCoords) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMaskCoords, layer)

	switch cmd {
	case Cmd_MaskCoords:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MaskCoords))
		x.Int(m.GetN())
		x.Double(m.GetX())
		x.Double(m.GetY())
		x.Double(m.GetZoom())

		return x.GetBuf()

	default:
		log.Errorf("invalid MaskCoords encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEmpty, layer)

	switch cmd {
	case Cmd_MessageEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEmpty))
		x.Int(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessage, layer)

	switch cmd {
	case Cmd_Message:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Message))
		// flags
		var flags int32 = 0
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMentioned() == true {
			flags |= 1 << 4
		}
		if m.GetMediaUnread() == true {
			flags |= 1 << 5
		}
		if m.GetSilent() == true {
			flags |= 1 << 13
		}
		if m.GetPost() == true {
			flags |= 1 << 14
		}
		if m.GetFromId90DDDC1171() != 0 {
			flags |= 1 << 8
		}
		if m.GetFwdFrom() != nil {
			flags |= 1 << 2
		}
		if m.GetViaBotId() != 0 {
			flags |= 1 << 11
		}
		if m.GetReplyToMsgId() != 0 {
			flags |= 1 << 3
		}
		if m.GetMedia() != nil {
			flags |= 1 << 9
		}
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 6
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 7
		}
		if m.GetViews() != 0 {
			flags |= 1 << 10
		}
		if m.GetEditDate() != 0 {
			flags |= 1 << 15
		}
		if len(m.GetPostAuthor()) > 0 {
			flags |= 1 << 16
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetFromId90DDDC1171() != 0 {
			x.Int(m.GetFromId90DDDC1171())
		}
		x.Bytes(m.GetToId().Encode(layer))
		if m.GetFwdFrom() != nil {
			x.Bytes(m.GetFwdFrom().Encode(layer))
		}
		if m.GetViaBotId() != 0 {
			x.Int(m.GetViaBotId())
		}
		if m.GetReplyToMsgId() != 0 {
			x.Int(m.GetReplyToMsgId())
		}
		x.Int(m.GetDate())
		x.String(m.GetMessage())
		if m.GetMedia() != nil {
			x.Bytes(m.GetMedia().Encode(layer))
		}
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetViews() != 0 {
			x.Int(m.GetViews())
		}
		if m.GetEditDate() != 0 {
			x.Int(m.GetEditDate())
		}
		if len(m.GetPostAuthor()) > 0 {
			x.String(m.GetPostAuthor())
		}

		return x.GetBuf()
	case Cmd_Message44f9b43d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Message44f9b43d))
		// flags
		var flags int32 = 0
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMentioned() == true {
			flags |= 1 << 4
		}
		if m.GetMediaUnread() == true {
			flags |= 1 << 5
		}
		if m.GetSilent() == true {
			flags |= 1 << 13
		}
		if m.GetPost() == true {
			flags |= 1 << 14
		}
		if m.GetFromId90DDDC1171() != 0 {
			flags |= 1 << 8
		}
		if m.GetFwdFrom() != nil {
			flags |= 1 << 2
		}
		if m.GetViaBotId() != 0 {
			flags |= 1 << 11
		}
		if m.GetReplyToMsgId() != 0 {
			flags |= 1 << 3
		}
		if m.GetMedia() != nil {
			flags |= 1 << 9
		}
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 6
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 7
		}
		if m.GetViews() != 0 {
			flags |= 1 << 10
		}
		if m.GetEditDate() != 0 {
			flags |= 1 << 15
		}
		if len(m.GetPostAuthor()) > 0 {
			flags |= 1 << 16
		}
		if m.GetGroupedId() != 0 {
			flags |= 1 << 17
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetFromId90DDDC1171() != 0 {
			x.Int(m.GetFromId90DDDC1171())
		}
		x.Bytes(m.GetToId().Encode(layer))
		if m.GetFwdFrom() != nil {
			x.Bytes(m.GetFwdFrom().Encode(layer))
		}
		if m.GetViaBotId() != 0 {
			x.Int(m.GetViaBotId())
		}
		if m.GetReplyToMsgId() != 0 {
			x.Int(m.GetReplyToMsgId())
		}
		x.Int(m.GetDate())
		x.String(m.GetMessage())
		if m.GetMedia() != nil {
			x.Bytes(m.GetMedia().Encode(layer))
		}
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetViews() != 0 {
			x.Int(m.GetViews())
		}
		if m.GetEditDate() != 0 {
			x.Int(m.GetEditDate())
		}
		if len(m.GetPostAuthor()) > 0 {
			x.String(m.GetPostAuthor())
		}
		if m.GetGroupedId() != 0 {
			x.Long(m.GetGroupedId())
		}

		return x.GetBuf()
	case Cmd_Message452c0e65:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Message452c0e65))
		// flags
		var flags int32 = 0
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMentioned() == true {
			flags |= 1 << 4
		}
		if m.GetMediaUnread() == true {
			flags |= 1 << 5
		}
		if m.GetSilent() == true {
			flags |= 1 << 13
		}
		if m.GetPost() == true {
			flags |= 1 << 14
		}
		if m.GetFromScheduled() == true {
			flags |= 1 << 18
		}
		if m.GetLegacy() == true {
			flags |= 1 << 19
		}
		if m.GetEditHide() == true {
			flags |= 1 << 21
		}
		if m.GetFromId90DDDC1171() != 0 {
			flags |= 1 << 8
		}
		if m.GetFwdFrom() != nil {
			flags |= 1 << 2
		}
		if m.GetViaBotId() != 0 {
			flags |= 1 << 11
		}
		if m.GetReplyToMsgId() != 0 {
			flags |= 1 << 3
		}
		if m.GetMedia() != nil {
			flags |= 1 << 9
		}
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 6
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 7
		}
		if m.GetViews() != 0 {
			flags |= 1 << 10
		}
		if m.GetEditDate() != 0 {
			flags |= 1 << 15
		}
		if len(m.GetPostAuthor()) > 0 {
			flags |= 1 << 16
		}
		if m.GetGroupedId() != 0 {
			flags |= 1 << 17
		}
		if len(m.GetRestrictionReason()) > 0 {
			flags |= 1 << 22
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetFromId90DDDC1171() != 0 {
			x.Int(m.GetFromId90DDDC1171())
		}
		x.Bytes(m.GetToId().Encode(layer))
		if m.GetFwdFrom() != nil {
			x.Bytes(m.GetFwdFrom().Encode(layer))
		}
		if m.GetViaBotId() != 0 {
			x.Int(m.GetViaBotId())
		}
		if m.GetReplyToMsgId() != 0 {
			x.Int(m.GetReplyToMsgId())
		}
		x.Int(m.GetDate())
		x.String(m.GetMessage())
		if m.GetMedia() != nil {
			x.Bytes(m.GetMedia().Encode(layer))
		}
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetViews() != 0 {
			x.Int(m.GetViews())
		}
		if m.GetEditDate() != 0 {
			x.Int(m.GetEditDate())
		}
		if len(m.GetPostAuthor()) > 0 {
			x.String(m.GetPostAuthor())
		}
		if m.GetGroupedId() != 0 {
			x.Long(m.GetGroupedId())
		}
		if len(m.GetRestrictionReason()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetRestrictionReason())))
			for _, v := range m.GetRestrictionReason() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_Message58ae39c9:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Message58ae39c9))
		// flags
		var flags int32 = 0
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMentioned() == true {
			flags |= 1 << 4
		}
		if m.GetMediaUnread() == true {
			flags |= 1 << 5
		}
		if m.GetSilent() == true {
			flags |= 1 << 13
		}
		if m.GetPost() == true {
			flags |= 1 << 14
		}
		if m.GetFromScheduled() == true {
			flags |= 1 << 18
		}
		if m.GetLegacy() == true {
			flags |= 1 << 19
		}
		if m.GetEditHide() == true {
			flags |= 1 << 21
		}
		if m.GetPinned() == true {
			flags |= 1 << 24
		}
		if m.GetFromId286FA604119() != nil {
			flags |= 1 << 8
		}
		if m.GetFwdFrom() != nil {
			flags |= 1 << 2
		}
		if m.GetViaBotId() != 0 {
			flags |= 1 << 11
		}
		if m.GetReplyTo() != nil {
			flags |= 1 << 3
		}
		if m.GetMedia() != nil {
			flags |= 1 << 9
		}
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 6
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 7
		}
		if m.GetViews() != 0 {
			flags |= 1 << 10
		}
		if m.GetForwards() != 0 {
			flags |= 1 << 10
		}
		if m.GetReplies() != nil {
			flags |= 1 << 23
		}
		if m.GetEditDate() != 0 {
			flags |= 1 << 15
		}
		if len(m.GetPostAuthor()) > 0 {
			flags |= 1 << 16
		}
		if m.GetGroupedId() != 0 {
			flags |= 1 << 17
		}
		if len(m.GetRestrictionReason()) > 0 {
			flags |= 1 << 22
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetFromId286FA604119() != nil {
			x.Bytes(m.GetFromId286FA604119().Encode(layer))
		}
		x.Bytes(m.GetPeerId().Encode(layer))
		if m.GetFwdFrom() != nil {
			x.Bytes(m.GetFwdFrom().Encode(layer))
		}
		if m.GetViaBotId() != 0 {
			x.Int(m.GetViaBotId())
		}
		if m.GetReplyTo() != nil {
			x.Bytes(m.GetReplyTo().Encode(layer))
		}
		x.Int(m.GetDate())
		x.String(m.GetMessage())
		if m.GetMedia() != nil {
			x.Bytes(m.GetMedia().Encode(layer))
		}
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetViews() != 0 {
			x.Int(m.GetViews())
		}
		if m.GetForwards() != 0 {
			x.Int(m.GetForwards())
		}
		if m.GetReplies() != nil {
			x.Bytes(m.GetReplies().Encode(layer))
		}
		if m.GetEditDate() != 0 {
			x.Int(m.GetEditDate())
		}
		if len(m.GetPostAuthor()) > 0 {
			x.String(m.GetPostAuthor())
		}
		if m.GetGroupedId() != 0 {
			x.Long(m.GetGroupedId())
		}
		if len(m.GetRestrictionReason()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetRestrictionReason())))
			for _, v := range m.GetRestrictionReason() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_Messagec09be45f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Messagec09be45f))
		// flags
		var flags int32 = 0
		if m.GetUnread() == true {
			flags |= 1 << 0
		}
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMentioned() == true {
			flags |= 1 << 4
		}
		if m.GetMediaUnread() == true {
			flags |= 1 << 5
		}
		if m.GetSilent() == true {
			flags |= 1 << 13
		}
		if m.GetPost() == true {
			flags |= 1 << 14
		}
		if m.GetFromId90DDDC1171() != 0 {
			flags |= 1 << 8
		}
		if m.GetFwdFrom() != nil {
			flags |= 1 << 2
		}
		if m.GetViaBotId() != 0 {
			flags |= 1 << 11
		}
		if m.GetReplyToMsgId() != 0 {
			flags |= 1 << 3
		}
		if m.GetMedia() != nil {
			flags |= 1 << 9
		}
		if m.GetReplyMarkup() != nil {
			flags |= 1 << 6
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 7
		}
		if m.GetViews() != 0 {
			flags |= 1 << 10
		}
		if m.GetEditDate() != 0 {
			flags |= 1 << 15
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetFromId90DDDC1171() != 0 {
			x.Int(m.GetFromId90DDDC1171())
		}
		x.Bytes(m.GetToId().Encode(layer))
		if m.GetFwdFrom() != nil {
			x.Bytes(m.GetFwdFrom().Encode(layer))
		}
		if m.GetViaBotId() != 0 {
			x.Int(m.GetViaBotId())
		}
		if m.GetReplyToMsgId() != 0 {
			x.Int(m.GetReplyToMsgId())
		}
		x.Int(m.GetDate())
		x.String(m.GetMessage())
		if m.GetMedia() != nil {
			x.Bytes(m.GetMedia().Encode(layer))
		}
		if m.GetReplyMarkup() != nil {
			x.Bytes(m.GetReplyMarkup().Encode(layer))
		}
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetViews() != 0 {
			x.Int(m.GetViews())
		}
		if m.GetEditDate() != 0 {
			x.Int(m.GetEditDate())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid Message encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageService) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageService, layer)

	switch cmd {
	case Cmd_MessageService:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageService))
		// flags
		var flags int32 = 0
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMentioned() == true {
			flags |= 1 << 4
		}
		if m.GetMediaUnread() == true {
			flags |= 1 << 5
		}
		if m.GetSilent() == true {
			flags |= 1 << 13
		}
		if m.GetPost() == true {
			flags |= 1 << 14
		}
		if m.GetFromId90DDDC1171() != 0 {
			flags |= 1 << 8
		}
		if m.GetReplyToMsgId() != 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetFromId90DDDC1171() != 0 {
			x.Int(m.GetFromId90DDDC1171())
		}
		x.Bytes(m.GetToId().Encode(layer))
		if m.GetReplyToMsgId() != 0 {
			x.Int(m.GetReplyToMsgId())
		}
		x.Int(m.GetDate())
		x.Bytes(m.GetAction().Encode(layer))

		return x.GetBuf()
	case Cmd_MessageService286fa604:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageService286fa604))
		// flags
		var flags int32 = 0
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMentioned() == true {
			flags |= 1 << 4
		}
		if m.GetMediaUnread() == true {
			flags |= 1 << 5
		}
		if m.GetSilent() == true {
			flags |= 1 << 13
		}
		if m.GetPost() == true {
			flags |= 1 << 14
		}
		if m.GetLegacy() == true {
			flags |= 1 << 19
		}
		if m.GetFromId286FA604119() != nil {
			flags |= 1 << 8
		}
		if m.GetReplyTo() != nil {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetFromId286FA604119() != nil {
			x.Bytes(m.GetFromId286FA604119().Encode(layer))
		}
		x.Bytes(m.GetPeerId().Encode(layer))
		if m.GetReplyTo() != nil {
			x.Bytes(m.GetReplyTo().Encode(layer))
		}
		x.Int(m.GetDate())
		x.Bytes(m.GetAction().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageService encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionEmpty, layer)

	switch cmd {
	case Cmd_MessageActionEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionChatCreate) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionChatCreate, layer)

	switch cmd {
	case Cmd_MessageActionChatCreate:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionChatCreate))
		x.String(m.GetTitle())
		x.VectorInt(m.GetUsers())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionChatCreate encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionChatEditTitle) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionChatEditTitle, layer)

	switch cmd {
	case Cmd_MessageActionChatEditTitle:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionChatEditTitle))
		x.String(m.GetTitle())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionChatEditTitle encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionChatEditPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionChatEditPhoto, layer)

	switch cmd {
	case Cmd_MessageActionChatEditPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionChatEditPhoto))
		x.Bytes(m.GetPhoto().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionChatEditPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionChatDeletePhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionChatDeletePhoto, layer)

	switch cmd {
	case Cmd_MessageActionChatDeletePhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionChatDeletePhoto))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionChatDeletePhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionChatAddUser) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionChatAddUser, layer)

	switch cmd {
	case Cmd_MessageActionChatAddUser:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionChatAddUser))
		x.VectorInt(m.GetUsers())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionChatAddUser encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionChatDeleteUser) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionChatDeleteUser, layer)

	switch cmd {
	case Cmd_MessageActionChatDeleteUser:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionChatDeleteUser))
		x.Int(m.GetUserId())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionChatDeleteUser encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionChatJoinedByLink) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionChatJoinedByLink, layer)

	switch cmd {
	case Cmd_MessageActionChatJoinedByLink:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionChatJoinedByLink))
		x.Int(m.GetInviterId())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionChatJoinedByLink encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionChannelCreate) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionChannelCreate, layer)

	switch cmd {
	case Cmd_MessageActionChannelCreate:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionChannelCreate))
		x.String(m.GetTitle())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionChannelCreate encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionChatMigrateTo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionChatMigrateTo, layer)

	switch cmd {
	case Cmd_MessageActionChatMigrateTo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionChatMigrateTo))
		x.Int(m.GetChannelId())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionChatMigrateTo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionChannelMigrateFrom) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionChannelMigrateFrom, layer)

	switch cmd {
	case Cmd_MessageActionChannelMigrateFrom:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionChannelMigrateFrom))
		x.String(m.GetTitle())
		x.Int(m.GetChatId())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionChannelMigrateFrom encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionPinMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionPinMessage, layer)

	switch cmd {
	case Cmd_MessageActionPinMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionPinMessage))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionPinMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionHistoryClear) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionHistoryClear, layer)

	switch cmd {
	case Cmd_MessageActionHistoryClear:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionHistoryClear))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionHistoryClear encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionGameScore) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionGameScore, layer)

	switch cmd {
	case Cmd_MessageActionGameScore:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionGameScore))
		x.Long(m.GetGameId())
		x.Int(m.GetScore())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionGameScore encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionPaymentSentMe) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionPaymentSentMe, layer)

	switch cmd {
	case Cmd_MessageActionPaymentSentMe:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionPaymentSentMe))
		// flags
		var flags int32 = 0
		if m.GetInfo() != nil {
			flags |= 1 << 0
		}
		if len(m.GetShippingOptionId()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.String(m.GetCurrency())
		x.Long(m.GetTotalAmount())
		x.StringBytes(m.GetPayload())
		if m.GetInfo() != nil {
			x.Bytes(m.GetInfo().Encode(layer))
		}
		if len(m.GetShippingOptionId()) > 0 {
			x.String(m.GetShippingOptionId())
		}
		x.Bytes(m.GetCharge().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionPaymentSentMe encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionPaymentSent) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionPaymentSent, layer)

	switch cmd {
	case Cmd_MessageActionPaymentSent:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionPaymentSent))
		x.String(m.GetCurrency())
		x.Long(m.GetTotalAmount())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionPaymentSent encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionScreenshotTaken) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionScreenshotTaken, layer)

	switch cmd {
	case Cmd_MessageActionScreenshotTaken:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionScreenshotTaken))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionScreenshotTaken encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionCustomAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionCustomAction, layer)

	switch cmd {
	case Cmd_MessageActionCustomAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionCustomAction))
		x.String(m.GetMessage())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionCustomAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionBotAllowed) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionBotAllowed, layer)

	switch cmd {
	case Cmd_MessageActionBotAllowed:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionBotAllowed))
		x.String(m.GetDomain())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionBotAllowed encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionSecureValuesSentMe) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionSecureValuesSentMe, layer)

	switch cmd {
	case Cmd_MessageActionSecureValuesSentMe:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionSecureValuesSentMe))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetValues())))
		for _, v := range m.GetValues() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.GetCredentials().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionSecureValuesSentMe encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionSecureValuesSent) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionSecureValuesSent, layer)

	switch cmd {
	case Cmd_MessageActionSecureValuesSent:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionSecureValuesSent))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetTypes())))
		for _, v := range m.GetTypes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionSecureValuesSent encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionContactSignUp) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionContactSignUp, layer)

	switch cmd {
	case Cmd_MessageActionContactSignUp:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionContactSignUp))
		// flags
		var flags int32 = 0
		if m.GetSilent() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		return x.GetBuf()
	case Cmd_MessageActionContactSignUpf3f25f76:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionContactSignUpf3f25f76))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionContactSignUp encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionPhoneCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionPhoneCall, layer)

	switch cmd {
	case Cmd_MessageActionPhoneCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionPhoneCall))
		// flags
		var flags int32 = 0
		if m.GetVideo() == true {
			flags |= 1 << 2
		}
		if m.GetReason() != nil {
			flags |= 1 << 0
		}
		if m.GetDuration() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetCallId())
		if m.GetReason() != nil {
			x.Bytes(m.GetReason().Encode(layer))
		}
		if m.GetDuration() != 0 {
			x.Int(m.GetDuration())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionPhoneCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionGeoProximityReached) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionGeoProximityReached, layer)

	switch cmd {
	case Cmd_MessageActionGeoProximityReached:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionGeoProximityReached))
		x.Bytes(m.GetFromId().Encode(layer))
		x.Bytes(m.GetToId().Encode(layer))
		x.Int(m.GetDistance())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionGeoProximityReached encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionGroupCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionGroupCall, layer)

	switch cmd {
	case Cmd_MessageActionGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionGroupCall))
		// flags
		var flags int32 = 0
		if m.GetDuration() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetCall().Encode(layer))
		if m.GetDuration() != 0 {
			x.Int(m.GetDuration())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionGroupCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageActionInviteToGroupCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageActionInviteToGroupCall, layer)

	switch cmd {
	case Cmd_MessageActionInviteToGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageActionInviteToGroupCall))
		x.Bytes(m.GetCall().Encode(layer))
		x.VectorInt(m.GetUsers())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageActionInviteToGroupCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityUnknown) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityUnknown, layer)

	switch cmd {
	case Cmd_MessageEntityUnknown:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityUnknown))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityUnknown encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityMention) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityMention, layer)

	switch cmd {
	case Cmd_MessageEntityMention:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityMention))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityMention encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityHashtag) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityHashtag, layer)

	switch cmd {
	case Cmd_MessageEntityHashtag:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityHashtag))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityHashtag encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityBotCommand) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityBotCommand, layer)

	switch cmd {
	case Cmd_MessageEntityBotCommand:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityBotCommand))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityBotCommand encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityUrl) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityUrl, layer)

	switch cmd {
	case Cmd_MessageEntityUrl:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityUrl))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityUrl encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityEmail) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityEmail, layer)

	switch cmd {
	case Cmd_MessageEntityEmail:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityEmail))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityEmail encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityBold) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityBold, layer)

	switch cmd {
	case Cmd_MessageEntityBold:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityBold))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityBold encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityItalic) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityItalic, layer)

	switch cmd {
	case Cmd_MessageEntityItalic:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityItalic))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityItalic encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityCode) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityCode, layer)

	switch cmd {
	case Cmd_MessageEntityCode:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityCode))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityCode encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityPre) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityPre, layer)

	switch cmd {
	case Cmd_MessageEntityPre:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityPre))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())
		x.String(m.GetLanguage())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityPre encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityTextUrl) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityTextUrl, layer)

	switch cmd {
	case Cmd_MessageEntityTextUrl:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityTextUrl))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())
		x.String(m.GetUrl())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityTextUrl encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityMentionName) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityMentionName, layer)

	switch cmd {
	case Cmd_MessageEntityMentionName:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityMentionName))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())
		x.Int(m.GetUserId352DCA5871())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityMentionName encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessageEntityMentionName) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessageEntityMentionName, layer)

	switch cmd {
	case Cmd_InputMessageEntityMentionName:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessageEntityMentionName))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())
		x.Bytes(m.GetUserId208E68C971().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessageEntityMentionName encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityPhone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityPhone, layer)

	switch cmd {
	case Cmd_MessageEntityPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityPhone))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityPhone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityCashtag) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityCashtag, layer)

	switch cmd {
	case Cmd_MessageEntityCashtag:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityCashtag))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityCashtag encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityUnderline) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityUnderline, layer)

	switch cmd {
	case Cmd_MessageEntityUnderline:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityUnderline))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityUnderline encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityStrike) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityStrike, layer)

	switch cmd {
	case Cmd_MessageEntityStrike:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityStrike))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityStrike encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityBlockquote) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityBlockquote, layer)

	switch cmd {
	case Cmd_MessageEntityBlockquote:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityBlockquote))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityBlockquote encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageEntityBankCard) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageEntityBankCard, layer)

	switch cmd {
	case Cmd_MessageEntityBankCard:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageEntityBankCard))
		x.Int(m.GetOffset())
		x.Int(m.GetLength())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageEntityBankCard encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageFwdHeader) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageFwdHeader, layer)

	switch cmd {
	case Cmd_MessageFwdHeader:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageFwdHeader))
		// flags
		var flags int32 = 0
		if m.GetFromIdFADFF4AC71() != 0 {
			flags |= 1 << 0
		}
		if m.GetChannelId() != 0 {
			flags |= 1 << 1
		}
		if m.GetChannelPost() != 0 {
			flags |= 1 << 2
		}
		if len(m.GetPostAuthor()) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		if m.GetFromIdFADFF4AC71() != 0 {
			x.Int(m.GetFromIdFADFF4AC71())
		}
		x.Int(m.GetDate())
		if m.GetChannelId() != 0 {
			x.Int(m.GetChannelId())
		}
		if m.GetChannelPost() != 0 {
			x.Int(m.GetChannelPost())
		}
		if len(m.GetPostAuthor()) > 0 {
			x.String(m.GetPostAuthor())
		}

		return x.GetBuf()
	case Cmd_MessageFwdHeader353a686b:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageFwdHeader353a686b))
		// flags
		var flags int32 = 0
		if m.GetFromIdFADFF4AC71() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetFromName()) > 0 {
			flags |= 1 << 5
		}
		if m.GetChannelId() != 0 {
			flags |= 1 << 1
		}
		if m.GetChannelPost() != 0 {
			flags |= 1 << 2
		}
		if len(m.GetPostAuthor()) > 0 {
			flags |= 1 << 3
		}
		if m.GetSavedFromPeer() != nil {
			flags |= 1 << 4
		}
		if m.GetSavedFromMsgId() != 0 {
			flags |= 1 << 4
		}
		if len(m.GetPsaType()) > 0 {
			flags |= 1 << 6
		}
		x.Int(flags)

		if m.GetFromIdFADFF4AC71() != 0 {
			x.Int(m.GetFromIdFADFF4AC71())
		}
		if len(m.GetFromName()) > 0 {
			x.String(m.GetFromName())
		}
		x.Int(m.GetDate())
		if m.GetChannelId() != 0 {
			x.Int(m.GetChannelId())
		}
		if m.GetChannelPost() != 0 {
			x.Int(m.GetChannelPost())
		}
		if len(m.GetPostAuthor()) > 0 {
			x.String(m.GetPostAuthor())
		}
		if m.GetSavedFromPeer() != nil {
			x.Bytes(m.GetSavedFromPeer().Encode(layer))
		}
		if m.GetSavedFromMsgId() != 0 {
			x.Int(m.GetSavedFromMsgId())
		}
		if len(m.GetPsaType()) > 0 {
			x.String(m.GetPsaType())
		}

		return x.GetBuf()
	case Cmd_MessageFwdHeader559ebe6d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageFwdHeader559ebe6d))
		// flags
		var flags int32 = 0
		if m.GetFromIdFADFF4AC71() != 0 {
			flags |= 1 << 0
		}
		if m.GetChannelId() != 0 {
			flags |= 1 << 1
		}
		if m.GetChannelPost() != 0 {
			flags |= 1 << 2
		}
		if len(m.GetPostAuthor()) > 0 {
			flags |= 1 << 3
		}
		if m.GetSavedFromPeer() != nil {
			flags |= 1 << 4
		}
		if m.GetSavedFromMsgId() != 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		if m.GetFromIdFADFF4AC71() != 0 {
			x.Int(m.GetFromIdFADFF4AC71())
		}
		x.Int(m.GetDate())
		if m.GetChannelId() != 0 {
			x.Int(m.GetChannelId())
		}
		if m.GetChannelPost() != 0 {
			x.Int(m.GetChannelPost())
		}
		if len(m.GetPostAuthor()) > 0 {
			x.String(m.GetPostAuthor())
		}
		if m.GetSavedFromPeer() != nil {
			x.Bytes(m.GetSavedFromPeer().Encode(layer))
		}
		if m.GetSavedFromMsgId() != 0 {
			x.Int(m.GetSavedFromMsgId())
		}

		return x.GetBuf()
	case Cmd_MessageFwdHeader5f777dce:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageFwdHeader5f777dce))
		// flags
		var flags int32 = 0
		if m.GetFromId5F777DCE119() != nil {
			flags |= 1 << 0
		}
		if len(m.GetFromName()) > 0 {
			flags |= 1 << 5
		}
		if m.GetChannelPost() != 0 {
			flags |= 1 << 2
		}
		if len(m.GetPostAuthor()) > 0 {
			flags |= 1 << 3
		}
		if m.GetSavedFromPeer() != nil {
			flags |= 1 << 4
		}
		if m.GetSavedFromMsgId() != 0 {
			flags |= 1 << 4
		}
		if len(m.GetPsaType()) > 0 {
			flags |= 1 << 6
		}
		x.Int(flags)

		if m.GetFromId5F777DCE119() != nil {
			x.Bytes(m.GetFromId5F777DCE119().Encode(layer))
		}
		if len(m.GetFromName()) > 0 {
			x.String(m.GetFromName())
		}
		x.Int(m.GetDate())
		if m.GetChannelPost() != 0 {
			x.Int(m.GetChannelPost())
		}
		if len(m.GetPostAuthor()) > 0 {
			x.String(m.GetPostAuthor())
		}
		if m.GetSavedFromPeer() != nil {
			x.Bytes(m.GetSavedFromPeer().Encode(layer))
		}
		if m.GetSavedFromMsgId() != 0 {
			x.Int(m.GetSavedFromMsgId())
		}
		if len(m.GetPsaType()) > 0 {
			x.String(m.GetPsaType())
		}

		return x.GetBuf()
	case Cmd_MessageFwdHeaderc786ddcb:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageFwdHeaderc786ddcb))
		// flags
		var flags int32 = 0
		if m.GetFromIdFADFF4AC71() != 0 {
			flags |= 1 << 0
		}
		if m.GetChannelId() != 0 {
			flags |= 1 << 1
		}
		if m.GetChannelPost() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.GetFromIdFADFF4AC71() != 0 {
			x.Int(m.GetFromIdFADFF4AC71())
		}
		x.Int(m.GetDate())
		if m.GetChannelId() != 0 {
			x.Int(m.GetChannelId())
		}
		if m.GetChannelPost() != 0 {
			x.Int(m.GetChannelPost())
		}

		return x.GetBuf()
	case Cmd_MessageFwdHeaderec338270:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageFwdHeaderec338270))
		// flags
		var flags int32 = 0
		if m.GetFromIdFADFF4AC71() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetFromName()) > 0 {
			flags |= 1 << 5
		}
		if m.GetChannelId() != 0 {
			flags |= 1 << 1
		}
		if m.GetChannelPost() != 0 {
			flags |= 1 << 2
		}
		if len(m.GetPostAuthor()) > 0 {
			flags |= 1 << 3
		}
		if m.GetSavedFromPeer() != nil {
			flags |= 1 << 4
		}
		if m.GetSavedFromMsgId() != 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		if m.GetFromIdFADFF4AC71() != 0 {
			x.Int(m.GetFromIdFADFF4AC71())
		}
		if len(m.GetFromName()) > 0 {
			x.String(m.GetFromName())
		}
		x.Int(m.GetDate())
		if m.GetChannelId() != 0 {
			x.Int(m.GetChannelId())
		}
		if m.GetChannelPost() != 0 {
			x.Int(m.GetChannelPost())
		}
		if len(m.GetPostAuthor()) > 0 {
			x.String(m.GetPostAuthor())
		}
		if m.GetSavedFromPeer() != nil {
			x.Bytes(m.GetSavedFromPeer().Encode(layer))
		}
		if m.GetSavedFromMsgId() != 0 {
			x.Int(m.GetSavedFromMsgId())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageFwdHeader encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageGroup) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageGroup, layer)

	switch cmd {
	case Cmd_MessageGroup:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageGroup))
		x.Int(m.GetMinId())
		x.Int(m.GetMaxId())
		x.Int(m.GetCount())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageGroup encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageInteractionCounters) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageInteractionCounters, layer)

	switch cmd {
	case Cmd_MessageInteractionCounters:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageInteractionCounters))
		x.Int(m.GetMsgId())
		x.Int(m.GetViews())
		x.Int(m.GetForwards())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageInteractionCounters encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaEmpty, layer)

	switch cmd {
	case Cmd_MessageMediaEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaPhoto, layer)

	switch cmd {
	case Cmd_MessageMediaPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaPhoto))
		// flags
		var flags int32 = 0
		if m.GetPhotoB5223B0F71() != nil {
			flags |= 1 << 0
		}
		if len(m.GetCaption()) > 0 {
			flags |= 1 << 1
		}
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.GetPhotoB5223B0F71() != nil {
			x.Bytes(m.GetPhotoB5223B0F71().Encode(layer))
		}
		if len(m.GetCaption()) > 0 {
			x.String(m.GetCaption())
		}
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_MessageMediaPhoto3d8ce53d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaPhoto3d8ce53d))
		x.Bytes(m.GetPhotoB5223B0F71().Encode(layer))
		x.String(m.GetCaption())

		return x.GetBuf()
	case Cmd_MessageMediaPhoto695150d7:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaPhoto695150d7))
		// flags
		var flags int32 = 0
		if m.GetPhotoB5223B0F71() != nil {
			flags |= 1 << 0
		}
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.GetPhotoB5223B0F71() != nil {
			x.Bytes(m.GetPhotoB5223B0F71().Encode(layer))
		}
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaGeo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaGeo, layer)

	switch cmd {
	case Cmd_MessageMediaGeo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaGeo))
		x.Bytes(m.GetGeo().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaGeo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaContact) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaContact, layer)

	switch cmd {
	case Cmd_MessageMediaContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaContact))
		x.String(m.GetPhoneNumber())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())
		x.Int(m.GetUserId())

		return x.GetBuf()
	case Cmd_MessageMediaContactcbf24940:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaContactcbf24940))
		x.String(m.GetPhoneNumber())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())
		x.String(m.GetVcard())
		x.Int(m.GetUserId())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaContact encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaUnsupported) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaUnsupported, layer)

	switch cmd {
	case Cmd_MessageMediaUnsupported:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaUnsupported))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaUnsupported encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaDocument) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaDocument, layer)

	switch cmd {
	case Cmd_MessageMediaDocument:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaDocument))
		// flags
		var flags int32 = 0
		if m.GetDocument() != nil {
			flags |= 1 << 0
		}
		if len(m.GetCaption()) > 0 {
			flags |= 1 << 1
		}
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.GetDocument() != nil {
			x.Bytes(m.GetDocument().Encode(layer))
		}
		if len(m.GetCaption()) > 0 {
			x.String(m.GetCaption())
		}
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_MessageMediaDocument9cb070d7:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaDocument9cb070d7))
		// flags
		var flags int32 = 0
		if m.GetDocument() != nil {
			flags |= 1 << 0
		}
		if m.GetTtlSeconds() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.GetDocument() != nil {
			x.Bytes(m.GetDocument().Encode(layer))
		}
		if m.GetTtlSeconds() != 0 {
			x.Int(m.GetTtlSeconds())
		}

		return x.GetBuf()
	case Cmd_MessageMediaDocumentf3e02ea8:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaDocumentf3e02ea8))
		x.Bytes(m.GetDocument().Encode(layer))
		x.String(m.GetCaption())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaDocument encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaWebPage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaWebPage, layer)

	switch cmd {
	case Cmd_MessageMediaWebPage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaWebPage))
		x.Bytes(m.GetWebpage().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaWebPage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaVenue) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaVenue, layer)

	switch cmd {
	case Cmd_MessageMediaVenue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaVenue))
		x.Bytes(m.GetGeo().Encode(layer))
		x.String(m.GetTitle())
		x.String(m.GetAddress())
		x.String(m.GetProvider())
		x.String(m.GetVenueId())

		return x.GetBuf()
	case Cmd_MessageMediaVenue2ec0533f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaVenue2ec0533f))
		x.Bytes(m.GetGeo().Encode(layer))
		x.String(m.GetTitle())
		x.String(m.GetAddress())
		x.String(m.GetProvider())
		x.String(m.GetVenueId())
		x.String(m.GetVenueType())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaVenue encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaGame) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaGame, layer)

	switch cmd {
	case Cmd_MessageMediaGame:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaGame))
		x.Bytes(m.GetGame().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaGame encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaInvoice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaInvoice, layer)

	switch cmd {
	case Cmd_MessageMediaInvoice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaInvoice))
		// flags
		var flags int32 = 0
		if m.GetShippingAddressRequested() == true {
			flags |= 1 << 1
		}
		if m.GetTest() == true {
			flags |= 1 << 3
		}
		if m.GetPhoto8455134771() != nil {
			flags |= 1 << 0
		}
		if m.GetReceiptMsgId() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetTitle())
		x.String(m.GetDescription())
		if m.GetPhoto8455134771() != nil {
			x.Bytes(m.GetPhoto8455134771().Encode(layer))
		}
		if m.GetReceiptMsgId() != 0 {
			x.Int(m.GetReceiptMsgId())
		}
		x.String(m.GetCurrency())
		x.Long(m.GetTotalAmount())
		x.String(m.GetStartParam())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaInvoice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaGeoLive) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaGeoLive, layer)

	switch cmd {
	case Cmd_MessageMediaGeoLive:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaGeoLive))
		x.Bytes(m.GetGeo().Encode(layer))
		x.Int(m.GetPeriod())

		return x.GetBuf()
	case Cmd_MessageMediaGeoLiveb940c666:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaGeoLiveb940c666))
		// flags
		var flags int32 = 0
		if m.GetHeading() != 0 {
			flags |= 1 << 0
		}
		if m.GetProximityNotificationRadius() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetGeo().Encode(layer))
		if m.GetHeading() != 0 {
			x.Int(m.GetHeading())
		}
		x.Int(m.GetPeriod())
		if m.GetProximityNotificationRadius() != 0 {
			x.Int(m.GetProximityNotificationRadius())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaGeoLive encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaPoll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaPoll, layer)

	switch cmd {
	case Cmd_MessageMediaPoll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaPoll))
		x.Bytes(m.GetPoll().Encode(layer))
		x.Bytes(m.GetResults().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaPoll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageMediaDice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageMediaDice, layer)

	switch cmd {
	case Cmd_MessageMediaDice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaDice))
		x.Int(m.GetValue())

		return x.GetBuf()
	case Cmd_MessageMediaDice3f7ee58b:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageMediaDice3f7ee58b))
		x.Int(m.GetValue())
		x.String(m.GetEmoticon())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageMediaDice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageRange) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageRange, layer)

	switch cmd {
	case Cmd_MessageRange:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageRange))
		x.Int(m.GetMinId())
		x.Int(m.GetMaxId())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageRange encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageReplies) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageReplies, layer)

	switch cmd {
	case Cmd_MessageReplies:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageReplies))
		// flags
		var flags int32 = 0
		if m.GetComments() == true {
			flags |= 1 << 0
		}
		if len(m.GetRecentRepliers()) > 0 {
			flags |= 1 << 1
		}
		if m.GetChannelId() != 0 {
			flags |= 1 << 0
		}
		if m.GetMaxId() != 0 {
			flags |= 1 << 2
		}
		if m.GetReadMaxId() != 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Int(m.GetReplies())
		x.Int(m.GetRepliesPts())
		if len(m.GetRecentRepliers()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetRecentRepliers())))
			for _, v := range m.GetRecentRepliers() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetChannelId() != 0 {
			x.Int(m.GetChannelId())
		}
		if m.GetMaxId() != 0 {
			x.Int(m.GetMaxId())
		}
		if m.GetReadMaxId() != 0 {
			x.Int(m.GetReadMaxId())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageReplies encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageReplyHeader) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageReplyHeader, layer)

	switch cmd {
	case Cmd_MessageReplyHeader:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageReplyHeader))
		// flags
		var flags int32 = 0
		if m.GetReplyToPeerId() != nil {
			flags |= 1 << 0
		}
		if m.GetReplyToTopId() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetReplyToMsgId())
		if m.GetReplyToPeerId() != nil {
			x.Bytes(m.GetReplyToPeerId().Encode(layer))
		}
		if m.GetReplyToTopId() != 0 {
			x.Int(m.GetReplyToTopId())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageReplyHeader encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageUserVote) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageUserVote, layer)

	switch cmd {
	case Cmd_MessageUserVote:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageUserVote))
		x.Int(m.GetUserId())
		x.StringBytes(m.GetOption())

		return x.GetBuf()
	case Cmd_MessageUserVotea28e5559:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageUserVotea28e5559))
		x.Int(m.GetUserId())
		x.StringBytes(m.GetOption())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageUserVote encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageUserVoteInputOption) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageUserVoteInputOption, layer)

	switch cmd {
	case Cmd_MessageUserVoteInputOption:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageUserVoteInputOption))
		x.Int(m.GetUserId())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageUserVoteInputOption encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageUserVoteMultiple) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageUserVoteMultiple, layer)

	switch cmd {
	case Cmd_MessageUserVoteMultiple:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageUserVoteMultiple))
		x.Int(m.GetUserId())

		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageUserVoteMultiple encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessageViews) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessageViews, layer)

	switch cmd {
	case Cmd_MessageViews:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessageViews))
		// flags
		var flags int32 = 0
		if m.GetViews() != 0 {
			flags |= 1 << 0
		}
		if m.GetForwards() != 0 {
			flags |= 1 << 1
		}
		if m.GetReplies() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.GetViews() != 0 {
			x.Int(m.GetViews())
		}
		if m.GetForwards() != 0 {
			x.Int(m.GetForwards())
		}
		if m.GetReplies() != nil {
			x.Bytes(m.GetReplies().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessageViews encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterEmpty, layer)

	switch cmd {
	case Cmd_InputMessagesFilterEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterPhotos) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterPhotos, layer)

	switch cmd {
	case Cmd_InputMessagesFilterPhotos:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterPhotos))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterPhotos encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterVideo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterVideo, layer)

	switch cmd {
	case Cmd_InputMessagesFilterVideo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterVideo))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterVideo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterPhotoVideo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterPhotoVideo, layer)

	switch cmd {
	case Cmd_InputMessagesFilterPhotoVideo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterPhotoVideo))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterPhotoVideo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterPhotoVideoDocuments) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterPhotoVideoDocuments, layer)

	switch cmd {
	case Cmd_InputMessagesFilterPhotoVideoDocuments:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterPhotoVideoDocuments))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterPhotoVideoDocuments encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterDocument) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterDocument, layer)

	switch cmd {
	case Cmd_InputMessagesFilterDocument:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterDocument))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterDocument encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterUrl) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterUrl, layer)

	switch cmd {
	case Cmd_InputMessagesFilterUrl:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterUrl))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterUrl encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterGif) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterGif, layer)

	switch cmd {
	case Cmd_InputMessagesFilterGif:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterGif))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterGif encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterVoice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterVoice, layer)

	switch cmd {
	case Cmd_InputMessagesFilterVoice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterVoice))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterVoice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterMusic) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterMusic, layer)

	switch cmd {
	case Cmd_InputMessagesFilterMusic:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterMusic))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterMusic encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterChatPhotos) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterChatPhotos, layer)

	switch cmd {
	case Cmd_InputMessagesFilterChatPhotos:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterChatPhotos))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterChatPhotos encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterPhoneCalls) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterPhoneCalls, layer)

	switch cmd {
	case Cmd_InputMessagesFilterPhoneCalls:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterPhoneCalls))
		// flags
		var flags int32 = 0
		if m.GetMissed() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterPhoneCalls encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterRoundVoice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterRoundVoice, layer)

	switch cmd {
	case Cmd_InputMessagesFilterRoundVoice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterRoundVoice))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterRoundVoice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterRoundVideo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterRoundVideo, layer)

	switch cmd {
	case Cmd_InputMessagesFilterRoundVideo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterRoundVideo))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterRoundVideo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterMyMentions) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterMyMentions, layer)

	switch cmd {
	case Cmd_InputMessagesFilterMyMentions:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterMyMentions))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterMyMentions encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterGeo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterGeo, layer)

	switch cmd {
	case Cmd_InputMessagesFilterGeo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterGeo))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterGeo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterContacts) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterContacts, layer)

	switch cmd {
	case Cmd_InputMessagesFilterContacts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterContacts))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterContacts encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputMessagesFilterPinned) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputMessagesFilterPinned, layer)

	switch cmd {
	case Cmd_InputMessagesFilterPinned:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputMessagesFilterPinned))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputMessagesFilterPinned encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesAffectedHistory) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesAffectedHistory, layer)

	switch cmd {
	case Cmd_MessagesAffectedHistory:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesAffectedHistory))
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())
		x.Int(m.GetOffset())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesAffectedHistory encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesAffectedMessages) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesAffectedMessages, layer)

	switch cmd {
	case Cmd_MessagesAffectedMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesAffectedMessages))
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesAffectedMessages encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesAllStickersNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesAllStickersNotModified, layer)

	switch cmd {
	case Cmd_MessagesAllStickersNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesAllStickersNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesAllStickersNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesAllStickers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesAllStickers, layer)

	switch cmd {
	case Cmd_MessagesAllStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesAllStickers))
		x.Int(m.GetHash())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSets())))
		for _, v := range m.GetSets() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesAllStickers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesArchivedStickers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesArchivedStickers, layer)

	switch cmd {
	case Cmd_MessagesArchivedStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesArchivedStickers))
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSets())))
		for _, v := range m.GetSets() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesArchivedStickers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesBotCallbackAnswer) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesBotCallbackAnswer, layer)

	switch cmd {
	case Cmd_MessagesBotCallbackAnswer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesBotCallbackAnswer))
		// flags
		var flags int32 = 0
		if m.GetAlert() == true {
			flags |= 1 << 1
		}
		if m.GetHasUrl() == true {
			flags |= 1 << 3
		}
		if len(m.GetMessage()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetUrl()) > 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		if len(m.GetMessage()) > 0 {
			x.String(m.GetMessage())
		}
		if len(m.GetUrl()) > 0 {
			x.String(m.GetUrl())
		}
		x.Int(m.GetCacheTime())

		return x.GetBuf()
	case Cmd_MessagesBotCallbackAnswer1264f1c6:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesBotCallbackAnswer1264f1c6))
		// flags
		var flags int32 = 0
		if m.GetAlert() == true {
			flags |= 1 << 1
		}
		if len(m.GetMessage()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if len(m.GetMessage()) > 0 {
			x.String(m.GetMessage())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesBotCallbackAnswer encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesBotResults) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesBotResults, layer)

	switch cmd {
	case Cmd_MessagesBotResults:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesBotResults))
		// flags
		var flags int32 = 0
		if m.GetGallery() == true {
			flags |= 1 << 0
		}
		if len(m.GetNextOffset()) > 0 {
			flags |= 1 << 1
		}
		if m.GetSwitchPm() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Long(m.GetQueryId())
		if len(m.GetNextOffset()) > 0 {
			x.String(m.GetNextOffset())
		}
		if m.GetSwitchPm() != nil {
			x.Bytes(m.GetSwitchPm().Encode(layer))
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetResults())))
		for _, v := range m.GetResults() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetCacheTime())

		return x.GetBuf()
	case Cmd_MessagesBotResults256709a6:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesBotResults256709a6))
		// flags
		var flags int32 = 0
		if m.GetGallery() == true {
			flags |= 1 << 0
		}
		if len(m.GetNextOffset()) > 0 {
			flags |= 1 << 1
		}
		if m.GetSwitchPm() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Long(m.GetQueryId())
		if len(m.GetNextOffset()) > 0 {
			x.String(m.GetNextOffset())
		}
		if m.GetSwitchPm() != nil {
			x.Bytes(m.GetSwitchPm().Encode(layer))
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetResults())))
		for _, v := range m.GetResults() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesBotResults947ca848:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesBotResults947ca848))
		// flags
		var flags int32 = 0
		if m.GetGallery() == true {
			flags |= 1 << 0
		}
		if len(m.GetNextOffset()) > 0 {
			flags |= 1 << 1
		}
		if m.GetSwitchPm() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Long(m.GetQueryId())
		if len(m.GetNextOffset()) > 0 {
			x.String(m.GetNextOffset())
		}
		if m.GetSwitchPm() != nil {
			x.Bytes(m.GetSwitchPm().Encode(layer))
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetResults())))
		for _, v := range m.GetResults() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetCacheTime())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesBotResults encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesChatFull) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesChatFull, layer)

	switch cmd {
	case Cmd_MessagesChatFull:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesChatFull))
		x.Bytes(m.GetFullChat().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesChatFull encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesChats) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesChats, layer)

	switch cmd {
	case Cmd_MessagesChats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesChats))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesChats encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesChatsSlice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesChatsSlice, layer)

	switch cmd {
	case Cmd_MessagesChatsSlice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesChatsSlice))
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesChatsSlice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesDhConfigNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesDhConfigNotModified, layer)

	switch cmd {
	case Cmd_MessagesDhConfigNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDhConfigNotModified))
		x.StringBytes(m.GetRandom())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesDhConfigNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesDhConfig) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesDhConfig, layer)

	switch cmd {
	case Cmd_MessagesDhConfig:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDhConfig))
		x.Int(m.GetG())
		x.StringBytes(m.GetP())
		x.Int(m.GetVersion())
		x.StringBytes(m.GetRandom())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesDhConfig encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesDialogs) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesDialogs, layer)

	switch cmd {
	case Cmd_MessagesDialogs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDialogs))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDialogs())))
		for _, v := range m.GetDialogs() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesDialogs encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesDialogsSlice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesDialogsSlice, layer)

	switch cmd {
	case Cmd_MessagesDialogsSlice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDialogsSlice))
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDialogs())))
		for _, v := range m.GetDialogs() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesDialogsSlice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesDialogsNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesDialogsNotModified, layer)

	switch cmd {
	case Cmd_MessagesDialogsNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDialogsNotModified))
		x.Int(m.GetCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesDialogsNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesDiscussionMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesDiscussionMessage, layer)

	switch cmd {
	case Cmd_MessagesDiscussionMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesDiscussionMessage))
		// flags
		var flags int32 = 0
		if m.GetMaxId() != 0 {
			flags |= 1 << 0
		}
		if m.GetReadInboxMaxId() != 0 {
			flags |= 1 << 1
		}
		if m.GetReadOutboxMaxId() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetMaxId() != 0 {
			x.Int(m.GetMaxId())
		}
		if m.GetReadInboxMaxId() != 0 {
			x.Int(m.GetReadInboxMaxId())
		}
		if m.GetReadOutboxMaxId() != 0 {
			x.Int(m.GetReadOutboxMaxId())
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesDiscussionMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesFavedStickersNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesFavedStickersNotModified, layer)

	switch cmd {
	case Cmd_MessagesFavedStickersNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesFavedStickersNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesFavedStickersNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesFavedStickers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesFavedStickers, layer)

	switch cmd {
	case Cmd_MessagesFavedStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesFavedStickers))
		x.Int(m.GetHash())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPacks())))
		for _, v := range m.GetPacks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetStickers())))
		for _, v := range m.GetStickers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesFavedStickers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesFeaturedStickersNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesFeaturedStickersNotModified, layer)

	switch cmd {
	case Cmd_MessagesFeaturedStickersNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesFeaturedStickersNotModified))
		_ = m

		return x.GetBuf()
	case Cmd_MessagesFeaturedStickersNotModifiedc6dc0c66:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesFeaturedStickersNotModifiedc6dc0c66))
		x.Int(m.GetCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesFeaturedStickersNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesFeaturedStickers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesFeaturedStickers, layer)

	switch cmd {
	case Cmd_MessagesFeaturedStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesFeaturedStickers))
		x.Int(m.GetHash())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSets())))
		for _, v := range m.GetSets() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.VectorLong(m.GetUnread())

		return x.GetBuf()
	case Cmd_MessagesFeaturedStickersb6abc341:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesFeaturedStickersb6abc341))
		x.Int(m.GetHash())
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSets())))
		for _, v := range m.GetSets() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.VectorLong(m.GetUnread())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesFeaturedStickers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesFoundGifs) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesFoundGifs, layer)

	switch cmd {
	case Cmd_MessagesFoundGifs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesFoundGifs))
		x.Int(m.GetNextOffset())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetResults())))
		for _, v := range m.GetResults() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesFoundGifs encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesFoundStickerSetsNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesFoundStickerSetsNotModified, layer)

	switch cmd {
	case Cmd_MessagesFoundStickerSetsNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesFoundStickerSetsNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesFoundStickerSetsNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesFoundStickerSets) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesFoundStickerSets, layer)

	switch cmd {
	case Cmd_MessagesFoundStickerSets:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesFoundStickerSets))
		x.Int(m.GetHash())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSets())))
		for _, v := range m.GetSets() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesFoundStickerSets encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesHighScores) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesHighScores, layer)

	switch cmd {
	case Cmd_MessagesHighScores:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesHighScores))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetScores())))
		for _, v := range m.GetScores() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesHighScores encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesInactiveChats) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesInactiveChats, layer)

	switch cmd {
	case Cmd_MessagesInactiveChats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesInactiveChats))
		x.VectorInt(m.GetDates())

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesInactiveChats encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesMessageEditData) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesMessageEditData, layer)

	switch cmd {
	case Cmd_MessagesMessageEditData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesMessageEditData))
		// flags
		var flags int32 = 0
		if m.GetCaption() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesMessageEditData encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesMessageViews) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesMessageViews, layer)

	switch cmd {
	case Cmd_MessagesMessageViews:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesMessageViews))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetViews())))
		for _, v := range m.GetViews() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesMessageViews encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesMessages) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesMessages, layer)

	switch cmd {
	case Cmd_MessagesMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesMessages))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesMessages encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesMessagesSlice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesMessagesSlice, layer)

	switch cmd {
	case Cmd_MessagesMessagesSlice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesMessagesSlice))
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesMessagesSlice3a54685e:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesMessagesSlice3a54685e))
		// flags
		var flags int32 = 0
		if m.GetInexact() == true {
			flags |= 1 << 1
		}
		if m.GetNextRate() != 0 {
			flags |= 1 << 0
		}
		if m.GetOffsetIdOffset() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.GetCount())
		if m.GetNextRate() != 0 {
			x.Int(m.GetNextRate())
		}
		if m.GetOffsetIdOffset() != 0 {
			x.Int(m.GetOffsetIdOffset())
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesMessagesSlicea6c47aaa:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesMessagesSlicea6c47aaa))
		// flags
		var flags int32 = 0
		if m.GetInexact() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesMessagesSlicec8edce1e:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesMessagesSlicec8edce1e))
		// flags
		var flags int32 = 0
		if m.GetInexact() == true {
			flags |= 1 << 1
		}
		if m.GetNextRate() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetCount())
		if m.GetNextRate() != 0 {
			x.Int(m.GetNextRate())
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesMessagesSlice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesChannelMessages) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesChannelMessages, layer)

	switch cmd {
	case Cmd_MessagesChannelMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesChannelMessages))
		// flags
		var flags int32 = 0
		x.Int(flags)

		x.Int(m.GetPts())
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesChannelMessages64479808:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesChannelMessages64479808))
		// flags
		var flags int32 = 0
		if m.GetInexact() == true {
			flags |= 1 << 1
		}
		if m.GetOffsetIdOffset() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(m.GetPts())
		x.Int(m.GetCount())
		if m.GetOffsetIdOffset() != 0 {
			x.Int(m.GetOffsetIdOffset())
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesChannelMessagesbc0f17bc:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesChannelMessagesbc0f17bc))
		// flags
		var flags int32 = 0
		if len(m.GetCollapsed()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetPts())
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if len(m.GetCollapsed()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetCollapsed())))
			for _, v := range m.GetCollapsed() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesChannelMessages encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesMessagesNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesMessagesNotModified, layer)

	switch cmd {
	case Cmd_MessagesMessagesNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesMessagesNotModified))
		x.Int(m.GetCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesMessagesNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesPeerDialogs) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesPeerDialogs, layer)

	switch cmd {
	case Cmd_MessagesPeerDialogs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesPeerDialogs))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDialogs())))
		for _, v := range m.GetDialogs() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.GetState().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesPeerDialogs encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesRecentStickersNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesRecentStickersNotModified, layer)

	switch cmd {
	case Cmd_MessagesRecentStickersNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesRecentStickersNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesRecentStickersNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesRecentStickers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesRecentStickers, layer)

	switch cmd {
	case Cmd_MessagesRecentStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesRecentStickers))
		x.Int(m.GetHash())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetStickers())))
		for _, v := range m.GetStickers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesRecentStickers22f3afb3:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesRecentStickers22f3afb3))
		x.Int(m.GetHash())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPacks())))
		for _, v := range m.GetPacks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetStickers())))
		for _, v := range m.GetStickers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.VectorInt(m.GetDates())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesRecentStickers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSavedGifsNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesSavedGifsNotModified, layer)

	switch cmd {
	case Cmd_MessagesSavedGifsNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSavedGifsNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesSavedGifsNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSavedGifs) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesSavedGifs, layer)

	switch cmd {
	case Cmd_MessagesSavedGifs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSavedGifs))
		x.Int(m.GetHash())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetGifs())))
		for _, v := range m.GetGifs() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesSavedGifs encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSearchCounter) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesSearchCounter, layer)

	switch cmd {
	case Cmd_MessagesSearchCounter:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSearchCounter))
		// flags
		var flags int32 = 0
		if m.GetInexact() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetFilter().Encode(layer))
		x.Int(m.GetCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesSearchCounter encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSentEncryptedMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesSentEncryptedMessage, layer)

	switch cmd {
	case Cmd_MessagesSentEncryptedMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSentEncryptedMessage))
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesSentEncryptedMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesSentEncryptedFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesSentEncryptedFile, layer)

	switch cmd {
	case Cmd_MessagesSentEncryptedFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesSentEncryptedFile))
		x.Int(m.GetDate())
		x.Bytes(m.GetFile().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesSentEncryptedFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesStickerSet) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesStickerSet, layer)

	switch cmd {
	case Cmd_MessagesStickerSet:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesStickerSet))
		x.Bytes(m.GetSet().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPacks())))
		for _, v := range m.GetPacks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDocuments())))
		for _, v := range m.GetDocuments() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesStickerSet encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesStickerSetInstallResultSuccess) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesStickerSetInstallResultSuccess, layer)

	switch cmd {
	case Cmd_MessagesStickerSetInstallResultSuccess:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesStickerSetInstallResultSuccess))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesStickerSetInstallResultSuccess encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesStickerSetInstallResultArchive) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesStickerSetInstallResultArchive, layer)

	switch cmd {
	case Cmd_MessagesStickerSetInstallResultArchive:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesStickerSetInstallResultArchive))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSets())))
		for _, v := range m.GetSets() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesStickerSetInstallResultArchive encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesStickersNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesStickersNotModified, layer)

	switch cmd {
	case Cmd_MessagesStickersNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesStickersNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesStickersNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesStickers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesStickers, layer)

	switch cmd {
	case Cmd_MessagesStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesStickers))
		x.String(m.GetHash8A8ECD3271())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetStickers())))
		for _, v := range m.GetStickers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_MessagesStickerse4599bbd:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesStickerse4599bbd))
		x.Int(m.GetHashE4599BBD85())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetStickers())))
		for _, v := range m.GetStickers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesStickers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMessagesVotesList) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMessagesVotesList, layer)

	switch cmd {
	case Cmd_MessagesVotesList:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MessagesVotesList))
		// flags
		var flags int32 = 0
		if len(m.GetNextOffset()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetVotes())))
		for _, v := range m.GetVotes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if len(m.GetNextOffset()) > 0 {
			x.String(m.GetNextOffset())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid MessagesVotesList encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMsgDetailedInfo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMsgDetailedInfo, layer)

	switch cmd {
	case Cmd_MsgDetailedInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MsgDetailedInfo))
		x.Long(m.GetMsgId())
		x.Long(m.GetAnswerMsgId())
		x.Int(m.GetBytes())
		x.Int(m.GetStatus())

		return x.GetBuf()

	default:
		log.Errorf("invalid MsgDetailedInfo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMsgNewDetailedInfo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMsgNewDetailedInfo, layer)

	switch cmd {
	case Cmd_MsgNewDetailedInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MsgNewDetailedInfo))
		x.Long(m.GetAnswerMsgId())
		x.Int(m.GetBytes())
		x.Int(m.GetStatus())

		return x.GetBuf()

	default:
		log.Errorf("invalid MsgNewDetailedInfo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMsgResendReq) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMsgResendReq, layer)

	switch cmd {
	case Cmd_MsgResendReq:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MsgResendReq))
		x.VectorLong(m.GetMsgIds())

		return x.GetBuf()

	default:
		log.Errorf("invalid MsgResendReq encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMsgsAck) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMsgsAck, layer)

	switch cmd {
	case Cmd_MsgsAck:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MsgsAck))
		x.VectorLong(m.GetMsgIds())

		return x.GetBuf()

	default:
		log.Errorf("invalid MsgsAck encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMsgsAllInfo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMsgsAllInfo, layer)

	switch cmd {
	case Cmd_MsgsAllInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MsgsAllInfo))
		x.VectorLong(m.GetMsgIds())

		x.String(m.GetInfo())

		return x.GetBuf()

	default:
		log.Errorf("invalid MsgsAllInfo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMsgsStateInfo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMsgsStateInfo, layer)

	switch cmd {
	case Cmd_MsgsStateInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MsgsStateInfo))
		x.Long(m.GetReqMsgId())
		x.String(m.GetInfo())

		return x.GetBuf()

	default:
		log.Errorf("invalid MsgsStateInfo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLMsgsStateReq) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassMsgsStateReq, layer)

	switch cmd {
	case Cmd_MsgsStateReq:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_MsgsStateReq))
		x.VectorLong(m.GetMsgIds())

		return x.GetBuf()

	default:
		log.Errorf("invalid MsgsStateReq encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLNearestDc) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassNearestDc, layer)

	switch cmd {
	case Cmd_NearestDc:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_NearestDc))
		x.String(m.GetCountry())
		x.Int(m.GetThisDc())
		x.Int(m.GetNearestDc())

		return x.GetBuf()

	default:
		log.Errorf("invalid NearestDc encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLNewSessionCreated) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassNewSessionCreated, layer)

	switch cmd {
	case Cmd_NewSessionCreated:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_NewSessionCreated))
		x.Long(m.GetFirstMsgId())
		x.Long(m.GetUniqueId())
		x.Long(m.GetServerSalt())

		return x.GetBuf()

	default:
		log.Errorf("invalid NewSessionCreated encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLNotifyPeer) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassNotifyPeer, layer)

	switch cmd {
	case Cmd_NotifyPeer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_NotifyPeer))
		x.Bytes(m.GetPeer().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid NotifyPeer encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLNotifyUsers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassNotifyUsers, layer)

	switch cmd {
	case Cmd_NotifyUsers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_NotifyUsers))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid NotifyUsers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLNotifyChats) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassNotifyChats, layer)

	switch cmd {
	case Cmd_NotifyChats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_NotifyChats))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid NotifyChats encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLNotifyAll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassNotifyAll, layer)

	switch cmd {
	case Cmd_NotifyAll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_NotifyAll))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid NotifyAll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLNotifyBroadcasts) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassNotifyBroadcasts, layer)

	switch cmd {
	case Cmd_NotifyBroadcasts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_NotifyBroadcasts))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid NotifyBroadcasts encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLNull) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassNull, layer)

	switch cmd {
	case Cmd_Null:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Null))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid Null encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPQInnerData) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPQInnerData, layer)

	switch cmd {
	case Cmd_PQInnerData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PQInnerData))
		x.String(m.GetPq())
		x.String(m.GetP())
		x.String(m.GetQ())
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.Bytes(m.GetNewNonce())

		return x.GetBuf()

	default:
		log.Errorf("invalid PQInnerData encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPQInnerDataDc) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPQInnerDataDc, layer)

	switch cmd {
	case Cmd_PQInnerDataDc:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PQInnerDataDc))
		x.String(m.GetPq())
		x.String(m.GetP())
		x.String(m.GetQ())
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.Bytes(m.GetNewNonce())
		x.Int(m.GetDc())

		return x.GetBuf()

	default:
		log.Errorf("invalid PQInnerDataDc encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPQInnerDataTemp) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPQInnerDataTemp, layer)

	switch cmd {
	case Cmd_PQInnerDataTemp:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PQInnerDataTemp))
		x.String(m.GetPq())
		x.String(m.GetP())
		x.String(m.GetQ())
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.Bytes(m.GetNewNonce())
		x.Int(m.GetExpiresIn())

		return x.GetBuf()

	default:
		log.Errorf("invalid PQInnerDataTemp encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPQInnerDataTempDc) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPQInnerDataTempDc, layer)

	switch cmd {
	case Cmd_PQInnerDataTempDc:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PQInnerDataTempDc))
		x.String(m.GetPq())
		x.String(m.GetP())
		x.String(m.GetQ())
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.Bytes(m.GetNewNonce())
		x.Int(m.GetDc())
		x.Int(m.GetExpiresIn())

		return x.GetBuf()

	default:
		log.Errorf("invalid PQInnerDataTempDc encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPagePart) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPagePart, layer)

	switch cmd {
	case Cmd_PagePart:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PagePart))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocks())))
		for _, v := range m.GetBlocks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPhotos())))
		for _, v := range m.GetPhotos() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDocuments())))
		for _, v := range m.GetDocuments() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PagePart encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageFull) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageFull, layer)

	switch cmd {
	case Cmd_PageFull:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageFull))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocks())))
		for _, v := range m.GetBlocks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPhotos())))
		for _, v := range m.GetPhotos() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDocuments())))
		for _, v := range m.GetDocuments() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PageFull encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPage, layer)

	switch cmd {
	case Cmd_Page:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Page))
		// flags
		var flags int32 = 0
		if m.GetPart() == true {
			flags |= 1 << 0
		}
		if m.GetRtl() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocks())))
		for _, v := range m.GetBlocks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPhotos())))
		for _, v := range m.GetPhotos() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDocuments())))
		for _, v := range m.GetDocuments() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_Page98657f0d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Page98657f0d))
		// flags
		var flags int32 = 0
		if m.GetPart() == true {
			flags |= 1 << 0
		}
		if m.GetRtl() == true {
			flags |= 1 << 1
		}
		if m.GetV2() == true {
			flags |= 1 << 2
		}
		if m.GetViews() != 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.String(m.GetUrl())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocks())))
		for _, v := range m.GetBlocks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPhotos())))
		for _, v := range m.GetPhotos() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDocuments())))
		for _, v := range m.GetDocuments() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetViews() != 0 {
			x.Int(m.GetViews())
		}

		return x.GetBuf()
	case Cmd_Pageae891bec:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Pageae891bec))
		// flags
		var flags int32 = 0
		if m.GetPart() == true {
			flags |= 1 << 0
		}
		if m.GetRtl() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.String(m.GetUrl())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocks())))
		for _, v := range m.GetBlocks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPhotos())))
		for _, v := range m.GetPhotos() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDocuments())))
		for _, v := range m.GetDocuments() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid Page encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockUnsupported) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockUnsupported, layer)

	switch cmd {
	case Cmd_PageBlockUnsupported:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockUnsupported))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockUnsupported encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockTitle) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockTitle, layer)

	switch cmd {
	case Cmd_PageBlockTitle:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockTitle))
		x.Bytes(m.GetText().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockTitle encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockSubtitle) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockSubtitle, layer)

	switch cmd {
	case Cmd_PageBlockSubtitle:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockSubtitle))
		x.Bytes(m.GetText().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockSubtitle encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockAuthorDate) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockAuthorDate, layer)

	switch cmd {
	case Cmd_PageBlockAuthorDate:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockAuthorDate))
		x.Bytes(m.GetAuthorBAAFE5E071().Encode(layer))
		x.Int(m.GetPublishedDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockAuthorDate encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockHeader) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockHeader, layer)

	switch cmd {
	case Cmd_PageBlockHeader:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockHeader))
		x.Bytes(m.GetText().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockHeader encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockSubheader) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockSubheader, layer)

	switch cmd {
	case Cmd_PageBlockSubheader:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockSubheader))
		x.Bytes(m.GetText().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockSubheader encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockParagraph) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockParagraph, layer)

	switch cmd {
	case Cmd_PageBlockParagraph:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockParagraph))
		x.Bytes(m.GetText().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockParagraph encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockPreformatted) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockPreformatted, layer)

	switch cmd {
	case Cmd_PageBlockPreformatted:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockPreformatted))
		x.Bytes(m.GetText().Encode(layer))
		x.String(m.GetLanguage())

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockPreformatted encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockFooter) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockFooter, layer)

	switch cmd {
	case Cmd_PageBlockFooter:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockFooter))
		x.Bytes(m.GetText().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockFooter encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockDivider) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockDivider, layer)

	switch cmd {
	case Cmd_PageBlockDivider:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockDivider))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockDivider encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockAnchor) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockAnchor, layer)

	switch cmd {
	case Cmd_PageBlockAnchor:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockAnchor))
		x.String(m.GetName())

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockAnchor encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockList) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockList, layer)

	switch cmd {
	case Cmd_PageBlockList:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockList))
		x.Bytes(m.GetOrdered().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetItems3A58C7F471())))
		for _, v := range m.GetItems3A58C7F471() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_PageBlockListe4e88011:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockListe4e88011))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetItemsE4E8801188())))
		for _, v := range m.GetItemsE4E8801188() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockList encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockBlockquote) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockBlockquote, layer)

	switch cmd {
	case Cmd_PageBlockBlockquote:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockBlockquote))
		x.Bytes(m.GetText().Encode(layer))
		x.Bytes(m.GetCaption263D7C2671().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockBlockquote encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockPullquote) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockPullquote, layer)

	switch cmd {
	case Cmd_PageBlockPullquote:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockPullquote))
		x.Bytes(m.GetText().Encode(layer))
		x.Bytes(m.GetCaption263D7C2671().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockPullquote encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockPhoto, layer)

	switch cmd {
	case Cmd_PageBlockPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockPhoto))
		x.Long(m.GetPhotoId())
		x.Bytes(m.GetCaption263D7C2671().Encode(layer))

		return x.GetBuf()
	case Cmd_PageBlockPhoto1759c560:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockPhoto1759c560))
		// flags
		var flags int32 = 0
		if len(m.GetUrl()) > 0 {
			flags |= 1 << 0
		}
		if m.GetWebpageId() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.GetPhotoId())
		x.Bytes(m.GetCaption1759C56088().Encode(layer))
		if len(m.GetUrl()) > 0 {
			x.String(m.GetUrl())
		}
		if m.GetWebpageId() != 0 {
			x.Long(m.GetWebpageId())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockVideo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockVideo, layer)

	switch cmd {
	case Cmd_PageBlockVideo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockVideo))
		// flags
		var flags int32 = 0
		if m.GetAutoplay() == true {
			flags |= 1 << 0
		}
		if m.GetLoop() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetVideoId())
		x.Bytes(m.GetCaption263D7C2671().Encode(layer))

		return x.GetBuf()
	case Cmd_PageBlockVideo7c8fe7b6:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockVideo7c8fe7b6))
		// flags
		var flags int32 = 0
		if m.GetAutoplay() == true {
			flags |= 1 << 0
		}
		if m.GetLoop() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetVideoId())
		x.Bytes(m.GetCaption1759C56088().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockVideo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockCover) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockCover, layer)

	switch cmd {
	case Cmd_PageBlockCover:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockCover))
		x.Bytes(m.GetCover().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockCover encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockEmbed) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockEmbed, layer)

	switch cmd {
	case Cmd_PageBlockEmbed:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockEmbed))
		// flags
		var flags int32 = 0
		if m.GetFullWidth() == true {
			flags |= 1 << 0
		}
		if m.GetAllowScrolling() == true {
			flags |= 1 << 3
		}
		if len(m.GetUrl()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetHtml()) > 0 {
			flags |= 1 << 2
		}
		if m.GetPosterPhotoId() != 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		if len(m.GetUrl()) > 0 {
			x.String(m.GetUrl())
		}
		if len(m.GetHtml()) > 0 {
			x.String(m.GetHtml())
		}
		if m.GetPosterPhotoId() != 0 {
			x.Long(m.GetPosterPhotoId())
		}
		x.Int(m.GetW())
		x.Int(m.GetH())
		x.Bytes(m.GetCaption263D7C2671().Encode(layer))

		return x.GetBuf()
	case Cmd_PageBlockEmbeda8718dc5:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockEmbeda8718dc5))
		// flags
		var flags int32 = 0
		if m.GetFullWidth() == true {
			flags |= 1 << 0
		}
		if m.GetAllowScrolling() == true {
			flags |= 1 << 3
		}
		if len(m.GetUrl()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetHtml()) > 0 {
			flags |= 1 << 2
		}
		if m.GetPosterPhotoId() != 0 {
			flags |= 1 << 4
		}
		if m.GetW() != 0 {
			flags |= 1 << 5
		}
		if m.GetH() != 0 {
			flags |= 1 << 5
		}
		x.Int(flags)

		if len(m.GetUrl()) > 0 {
			x.String(m.GetUrl())
		}
		if len(m.GetHtml()) > 0 {
			x.String(m.GetHtml())
		}
		if m.GetPosterPhotoId() != 0 {
			x.Long(m.GetPosterPhotoId())
		}
		if m.GetW() != 0 {
			x.Int(m.GetW())
		}
		if m.GetH() != 0 {
			x.Int(m.GetH())
		}
		x.Bytes(m.GetCaption1759C56088().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockEmbed encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockEmbedPost) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockEmbedPost, layer)

	switch cmd {
	case Cmd_PageBlockEmbedPost:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockEmbedPost))
		x.String(m.GetUrl())
		x.Long(m.GetWebpageId())
		x.Long(m.GetAuthorPhotoId())
		x.String(m.GetAuthor292C7BE971())
		x.Int(m.GetDate())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocks())))
		for _, v := range m.GetBlocks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.GetCaption263D7C2671().Encode(layer))

		return x.GetBuf()
	case Cmd_PageBlockEmbedPostf259a80b:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockEmbedPostf259a80b))
		x.String(m.GetUrl())
		x.Long(m.GetWebpageId())
		x.Long(m.GetAuthorPhotoId())
		x.String(m.GetAuthor292C7BE971())
		x.Int(m.GetDate())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocks())))
		for _, v := range m.GetBlocks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.GetCaption1759C56088().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockEmbedPost encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockCollage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockCollage, layer)

	switch cmd {
	case Cmd_PageBlockCollage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockCollage))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetItems8B31C4F71())))
		for _, v := range m.GetItems8B31C4F71() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.GetCaption263D7C2671().Encode(layer))

		return x.GetBuf()
	case Cmd_PageBlockCollage65a0fa4d:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockCollage65a0fa4d))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetItems8B31C4F71())))
		for _, v := range m.GetItems8B31C4F71() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.GetCaption1759C56088().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockCollage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockSlideshow) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockSlideshow, layer)

	switch cmd {
	case Cmd_PageBlockSlideshow:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockSlideshow))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetItems8B31C4F71())))
		for _, v := range m.GetItems8B31C4F71() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.GetCaption263D7C2671().Encode(layer))

		return x.GetBuf()
	case Cmd_PageBlockSlideshow31f9590:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockSlideshow31f9590))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetItems8B31C4F71())))
		for _, v := range m.GetItems8B31C4F71() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.GetCaption1759C56088().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockSlideshow encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockChannel) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockChannel, layer)

	switch cmd {
	case Cmd_PageBlockChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockChannel))
		x.Bytes(m.GetChannel().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockChannel encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockAudio) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockAudio, layer)

	switch cmd {
	case Cmd_PageBlockAudio:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockAudio))
		x.Long(m.GetAudioId())
		x.Bytes(m.GetCaption263D7C2671().Encode(layer))

		return x.GetBuf()
	case Cmd_PageBlockAudio804361ea:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockAudio804361ea))
		x.Long(m.GetAudioId())
		x.Bytes(m.GetCaption1759C56088().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockAudio encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockKicker) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockKicker, layer)

	switch cmd {
	case Cmd_PageBlockKicker:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockKicker))
		x.Bytes(m.GetText().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockKicker encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockTable) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockTable, layer)

	switch cmd {
	case Cmd_PageBlockTable:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockTable))
		// flags
		var flags int32 = 0
		if m.GetBordered() == true {
			flags |= 1 << 0
		}
		if m.GetStriped() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetTitle().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetRows())))
		for _, v := range m.GetRows() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockTable encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockOrderedList) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockOrderedList, layer)

	switch cmd {
	case Cmd_PageBlockOrderedList:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockOrderedList))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetItems9A8AE1E188())))
		for _, v := range m.GetItems9A8AE1E188() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockOrderedList encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockDetails) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockDetails, layer)

	switch cmd {
	case Cmd_PageBlockDetails:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockDetails))
		// flags
		var flags int32 = 0
		if m.GetOpen() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocks())))
		for _, v := range m.GetBlocks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.GetTitle().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockDetails encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockRelatedArticles) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockRelatedArticles, layer)

	switch cmd {
	case Cmd_PageBlockRelatedArticles:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockRelatedArticles))
		x.Bytes(m.GetTitle().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetArticles())))
		for _, v := range m.GetArticles() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockRelatedArticles encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageBlockMap) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageBlockMap, layer)

	switch cmd {
	case Cmd_PageBlockMap:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageBlockMap))
		x.Bytes(m.GetGeo().Encode(layer))
		x.Int(m.GetZoom())
		x.Int(m.GetW())
		x.Int(m.GetH())
		x.Bytes(m.GetCaption1759C56088().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageBlockMap encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageCaption) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageCaption, layer)

	switch cmd {
	case Cmd_PageCaption:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageCaption))
		x.Bytes(m.GetText().Encode(layer))
		x.Bytes(m.GetCredit().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageCaption encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageListItemText) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageListItemText, layer)

	switch cmd {
	case Cmd_PageListItemText:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageListItemText))
		x.Bytes(m.GetText().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageListItemText encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageListItemBlocks) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageListItemBlocks, layer)

	switch cmd {
	case Cmd_PageListItemBlocks:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageListItemBlocks))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocks())))
		for _, v := range m.GetBlocks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PageListItemBlocks encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageListOrderedItemText) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageListOrderedItemText, layer)

	switch cmd {
	case Cmd_PageListOrderedItemText:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageListOrderedItemText))
		x.String(m.GetNum())
		x.Bytes(m.GetText().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PageListOrderedItemText encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageListOrderedItemBlocks) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageListOrderedItemBlocks, layer)

	switch cmd {
	case Cmd_PageListOrderedItemBlocks:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageListOrderedItemBlocks))
		x.String(m.GetNum())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetBlocks())))
		for _, v := range m.GetBlocks() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PageListOrderedItemBlocks encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageRelatedArticle) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageRelatedArticle, layer)

	switch cmd {
	case Cmd_PageRelatedArticle:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageRelatedArticle))
		// flags
		var flags int32 = 0
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 1
		}
		if m.GetPhotoId() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.String(m.GetUrl())
		x.Long(m.GetWebpageId())
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		if m.GetPhotoId() != 0 {
			x.Long(m.GetPhotoId())
		}

		return x.GetBuf()
	case Cmd_PageRelatedArticleb390dc08:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageRelatedArticleb390dc08))
		// flags
		var flags int32 = 0
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 1
		}
		if m.GetPhotoId() != 0 {
			flags |= 1 << 2
		}
		if len(m.GetAuthor()) > 0 {
			flags |= 1 << 3
		}
		if m.GetPublishedDate() != 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		x.String(m.GetUrl())
		x.Long(m.GetWebpageId())
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		if m.GetPhotoId() != 0 {
			x.Long(m.GetPhotoId())
		}
		if len(m.GetAuthor()) > 0 {
			x.String(m.GetAuthor())
		}
		if m.GetPublishedDate() != 0 {
			x.Int(m.GetPublishedDate())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PageRelatedArticle encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageTableCell) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageTableCell, layer)

	switch cmd {
	case Cmd_PageTableCell:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageTableCell))
		// flags
		var flags int32 = 0
		if m.GetHeader() == true {
			flags |= 1 << 0
		}
		if m.GetAlignCenter() == true {
			flags |= 1 << 3
		}
		if m.GetAlignRight() == true {
			flags |= 1 << 4
		}
		if m.GetValignMiddle() == true {
			flags |= 1 << 5
		}
		if m.GetValignBottom() == true {
			flags |= 1 << 6
		}
		if m.GetText() != nil {
			flags |= 1 << 7
		}
		if m.GetColspan() != 0 {
			flags |= 1 << 1
		}
		if m.GetRowspan() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.GetText() != nil {
			x.Bytes(m.GetText().Encode(layer))
		}
		if m.GetColspan() != 0 {
			x.Int(m.GetColspan())
		}
		if m.GetRowspan() != 0 {
			x.Int(m.GetRowspan())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PageTableCell encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPageTableRow) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPageTableRow, layer)

	switch cmd {
	case Cmd_PageTableRow:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PageTableRow))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetCells())))
		for _, v := range m.GetCells() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PageTableRow encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPasswordKdfAlgoUnknown) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPasswordKdfAlgoUnknown, layer)

	switch cmd {
	case Cmd_PasswordKdfAlgoUnknown:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PasswordKdfAlgoUnknown))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PasswordKdfAlgoUnknown encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow, layer)

	switch cmd {
	case Cmd_PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow))
		x.StringBytes(m.GetSalt1())
		x.StringBytes(m.GetSalt2())
		x.Int(m.GetG())
		x.StringBytes(m.GetP())

		return x.GetBuf()

	default:
		log.Errorf("invalid PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentCharge) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPaymentCharge, layer)

	switch cmd {
	case Cmd_PaymentCharge:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentCharge))
		x.String(m.GetId())
		x.String(m.GetProviderChargeId())

		return x.GetBuf()

	default:
		log.Errorf("invalid PaymentCharge encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentRequestedInfo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPaymentRequestedInfo, layer)

	switch cmd {
	case Cmd_PaymentRequestedInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentRequestedInfo))
		// flags
		var flags int32 = 0
		if len(m.GetName()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetPhone()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetEmail()) > 0 {
			flags |= 1 << 2
		}
		if m.GetShippingAddress() != nil {
			flags |= 1 << 3
		}
		x.Int(flags)

		if len(m.GetName()) > 0 {
			x.String(m.GetName())
		}
		if len(m.GetPhone()) > 0 {
			x.String(m.GetPhone())
		}
		if len(m.GetEmail()) > 0 {
			x.String(m.GetEmail())
		}
		if m.GetShippingAddress() != nil {
			x.Bytes(m.GetShippingAddress().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PaymentRequestedInfo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentSavedCredentialsCard) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPaymentSavedCredentialsCard, layer)

	switch cmd {
	case Cmd_PaymentSavedCredentialsCard:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentSavedCredentialsCard))
		x.String(m.GetId())
		x.String(m.GetTitle())

		return x.GetBuf()

	default:
		log.Errorf("invalid PaymentSavedCredentialsCard encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsBankCardData) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPaymentsBankCardData, layer)

	switch cmd {
	case Cmd_PaymentsBankCardData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsBankCardData))
		x.String(m.GetTitle())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetOpenUrls())))
		for _, v := range m.GetOpenUrls() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PaymentsBankCardData encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsPaymentForm) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPaymentsPaymentForm, layer)

	switch cmd {
	case Cmd_PaymentsPaymentForm:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsPaymentForm))
		// flags
		var flags int32 = 0
		if m.GetCanSaveCredentials() == true {
			flags |= 1 << 2
		}
		if m.GetPasswordMissing() == true {
			flags |= 1 << 3
		}
		if len(m.GetNativeProvider()) > 0 {
			flags |= 1 << 4
		}
		if m.GetNativeParams() != nil {
			flags |= 1 << 4
		}
		if m.GetSavedInfo() != nil {
			flags |= 1 << 0
		}
		if m.GetSavedCredentials() != nil {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetBotId())
		x.Bytes(m.GetInvoice().Encode(layer))
		x.Int(m.GetProviderId())
		x.String(m.GetUrl())
		if len(m.GetNativeProvider()) > 0 {
			x.String(m.GetNativeProvider())
		}
		if m.GetNativeParams() != nil {
			x.Bytes(m.GetNativeParams().Encode(layer))
		}
		if m.GetSavedInfo() != nil {
			x.Bytes(m.GetSavedInfo().Encode(layer))
		}
		if m.GetSavedCredentials() != nil {
			x.Bytes(m.GetSavedCredentials().Encode(layer))
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PaymentsPaymentForm encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsPaymentReceipt) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPaymentsPaymentReceipt, layer)

	switch cmd {
	case Cmd_PaymentsPaymentReceipt:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsPaymentReceipt))
		// flags
		var flags int32 = 0
		if m.GetInfo() != nil {
			flags |= 1 << 0
		}
		if m.GetShipping() != nil {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetDate())
		x.Int(m.GetBotId())
		x.Bytes(m.GetInvoice().Encode(layer))
		x.Int(m.GetProviderId())
		if m.GetInfo() != nil {
			x.Bytes(m.GetInfo().Encode(layer))
		}
		if m.GetShipping() != nil {
			x.Bytes(m.GetShipping().Encode(layer))
		}
		x.String(m.GetCurrency())
		x.Long(m.GetTotalAmount())
		x.String(m.GetCredentialsTitle())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PaymentsPaymentReceipt encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsPaymentResult) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPaymentsPaymentResult, layer)

	switch cmd {
	case Cmd_PaymentsPaymentResult:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsPaymentResult))
		x.Bytes(m.GetUpdates().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PaymentsPaymentResult encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsPaymentVerficationNeeded) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPaymentsPaymentVerficationNeeded, layer)

	switch cmd {
	case Cmd_PaymentsPaymentVerficationNeeded:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsPaymentVerficationNeeded))
		x.String(m.GetUrl())

		return x.GetBuf()

	default:
		log.Errorf("invalid PaymentsPaymentVerficationNeeded encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsPaymentVerificationNeeded) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPaymentsPaymentVerificationNeeded, layer)

	switch cmd {
	case Cmd_PaymentsPaymentVerificationNeeded:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsPaymentVerificationNeeded))
		x.String(m.GetUrl())

		return x.GetBuf()

	default:
		log.Errorf("invalid PaymentsPaymentVerificationNeeded encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsSavedInfo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPaymentsSavedInfo, layer)

	switch cmd {
	case Cmd_PaymentsSavedInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsSavedInfo))
		// flags
		var flags int32 = 0
		if m.GetHasSavedCredentials() == true {
			flags |= 1 << 1
		}
		if m.GetSavedInfo() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetSavedInfo() != nil {
			x.Bytes(m.GetSavedInfo().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PaymentsSavedInfo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPaymentsValidatedRequestedInfo) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPaymentsValidatedRequestedInfo, layer)

	switch cmd {
	case Cmd_PaymentsValidatedRequestedInfo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PaymentsValidatedRequestedInfo))
		// flags
		var flags int32 = 0
		if len(m.GetId()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetShippingOptions()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		if len(m.GetId()) > 0 {
			x.String(m.GetId())
		}
		if len(m.GetShippingOptions()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetShippingOptions())))
			for _, v := range m.GetShippingOptions() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PaymentsValidatedRequestedInfo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPeerUser) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPeerUser, layer)

	switch cmd {
	case Cmd_PeerUser:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerUser))
		x.Int(m.GetUserId())

		return x.GetBuf()

	default:
		log.Errorf("invalid PeerUser encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPeerChat) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPeerChat, layer)

	switch cmd {
	case Cmd_PeerChat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerChat))
		x.Int(m.GetChatId())

		return x.GetBuf()

	default:
		log.Errorf("invalid PeerChat encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPeerChannel) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPeerChannel, layer)

	switch cmd {
	case Cmd_PeerChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerChannel))
		x.Int(m.GetChannelId())

		return x.GetBuf()

	default:
		log.Errorf("invalid PeerChannel encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPeerBlocked) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPeerBlocked, layer)

	switch cmd {
	case Cmd_PeerBlocked:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerBlocked))
		x.Bytes(m.GetPeerId().Encode(layer))
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid PeerBlocked encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPeerLocated) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPeerLocated, layer)

	switch cmd {
	case Cmd_PeerLocated:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerLocated))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Int(m.GetExpires())
		x.Int(m.GetDistance())

		return x.GetBuf()

	default:
		log.Errorf("invalid PeerLocated encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPeerSelfLocated) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPeerSelfLocated, layer)

	switch cmd {
	case Cmd_PeerSelfLocated:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerSelfLocated))
		x.Int(m.GetExpires())

		return x.GetBuf()

	default:
		log.Errorf("invalid PeerSelfLocated encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPeerNotifyEventsEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPeerNotifyEventsEmpty, layer)

	switch cmd {
	case Cmd_PeerNotifyEventsEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerNotifyEventsEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PeerNotifyEventsEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPeerNotifyEventsAll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPeerNotifyEventsAll, layer)

	switch cmd {
	case Cmd_PeerNotifyEventsAll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerNotifyEventsAll))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PeerNotifyEventsAll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPeerNotifySettingsEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPeerNotifySettingsEmpty, layer)

	switch cmd {
	case Cmd_PeerNotifySettingsEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerNotifySettingsEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PeerNotifySettingsEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPeerNotifySettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPeerNotifySettings, layer)

	switch cmd {
	case Cmd_PeerNotifySettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerNotifySettings))
		// flags
		var flags int32 = 0
		if m.GetShowPreviews9ACDA4C071() == true {
			flags |= 1 << 0
		}
		if m.GetSilent9ACDA4C071() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetMuteUntil())
		x.String(m.GetSound())

		return x.GetBuf()
	case Cmd_PeerNotifySettingsaf509d20:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerNotifySettingsaf509d20))
		// flags
		var flags int32 = 0
		if m.GetShowPreviewsAF509D2085() != nil {
			flags |= 1 << 0
		}
		if m.GetSilentAF509D2085() != nil {
			flags |= 1 << 1
		}
		if m.GetMuteUntil() != 0 {
			flags |= 1 << 2
		}
		if len(m.GetSound()) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		if m.GetShowPreviewsAF509D2085() != nil {
			x.Bytes(m.GetShowPreviewsAF509D2085().Encode(layer))
		}
		if m.GetSilentAF509D2085() != nil {
			x.Bytes(m.GetSilentAF509D2085().Encode(layer))
		}
		if m.GetMuteUntil() != 0 {
			x.Int(m.GetMuteUntil())
		}
		if len(m.GetSound()) > 0 {
			x.String(m.GetSound())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PeerNotifySettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPeerSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPeerSettings, layer)

	switch cmd {
	case Cmd_PeerSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerSettings))
		// flags
		var flags int32 = 0
		if m.GetReportSpam() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		return x.GetBuf()
	case Cmd_PeerSettings733f2961:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PeerSettings733f2961))
		// flags
		var flags int32 = 0
		if m.GetReportSpam() == true {
			flags |= 1 << 0
		}
		if m.GetAddContact() == true {
			flags |= 1 << 1
		}
		if m.GetBlockContact() == true {
			flags |= 1 << 2
		}
		if m.GetShareContact() == true {
			flags |= 1 << 3
		}
		if m.GetNeedContactsException() == true {
			flags |= 1 << 4
		}
		if m.GetReportGeo() == true {
			flags |= 1 << 5
		}
		if m.GetAutoarchived() == true {
			flags |= 1 << 7
		}
		if m.GetGeoDistance() != 0 {
			flags |= 1 << 6
		}
		x.Int(flags)

		if m.GetGeoDistance() != 0 {
			x.Int(m.GetGeoDistance())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PeerSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCallEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneCallEmpty, layer)

	switch cmd {
	case Cmd_PhoneCallEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallEmpty))
		x.Long(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneCallEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCallRequested) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneCallRequested, layer)

	switch cmd {
	case Cmd_PhoneCallRequested:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallRequested))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())
		x.StringBytes(m.GetGAHash())
		x.Bytes(m.GetProtocol().Encode(layer))

		return x.GetBuf()
	case Cmd_PhoneCallRequested87eabb53:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallRequested87eabb53))
		// flags
		var flags int32 = 0
		if m.GetVideo() == true {
			flags |= 1 << 6
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())
		x.StringBytes(m.GetGAHash())
		x.Bytes(m.GetProtocol().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneCallRequested encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCallAccepted) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneCallAccepted, layer)

	switch cmd {
	case Cmd_PhoneCallAccepted:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallAccepted))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())
		x.StringBytes(m.GetGB())
		x.Bytes(m.GetProtocol().Encode(layer))

		return x.GetBuf()
	case Cmd_PhoneCallAccepted997c454a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallAccepted997c454a))
		// flags
		var flags int32 = 0
		if m.GetVideo() == true {
			flags |= 1 << 6
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())
		x.StringBytes(m.GetGB())
		x.Bytes(m.GetProtocol().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneCallAccepted encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneCall, layer)

	switch cmd {
	case Cmd_PhoneCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCall))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())
		x.StringBytes(m.GetGAOrB())
		x.Long(m.GetKeyFingerprint())
		x.Bytes(m.GetProtocol().Encode(layer))
		x.Bytes(m.GetConnection().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAlternativeConnections())))
		for _, v := range m.GetAlternativeConnections() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetStartDate())

		return x.GetBuf()
	case Cmd_PhoneCall8742ae7f:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCall8742ae7f))
		// flags
		var flags int32 = 0
		if m.GetP2PAllowed() == true {
			flags |= 1 << 5
		}
		if m.GetVideo() == true {
			flags |= 1 << 6
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())
		x.StringBytes(m.GetGAOrB())
		x.Long(m.GetKeyFingerprint())
		x.Bytes(m.GetProtocol().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetConnections())))
		for _, v := range m.GetConnections() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetStartDate())

		return x.GetBuf()
	case Cmd_PhoneCalle6f9ddf3:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCalle6f9ddf3))
		// flags
		var flags int32 = 0
		if m.GetP2PAllowed() == true {
			flags |= 1 << 5
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())
		x.StringBytes(m.GetGAOrB())
		x.Long(m.GetKeyFingerprint())
		x.Bytes(m.GetProtocol().Encode(layer))
		x.Bytes(m.GetConnection().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAlternativeConnections())))
		for _, v := range m.GetAlternativeConnections() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetStartDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCallWaiting) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneCallWaiting, layer)

	switch cmd {
	case Cmd_PhoneCallWaiting:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallWaiting))
		// flags
		var flags int32 = 0
		if m.GetVideo() == true {
			flags |= 1 << 6
		}
		if m.GetReceiveDate() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(m.GetAdminId())
		x.Int(m.GetParticipantId())
		x.Bytes(m.GetProtocol().Encode(layer))
		if m.GetReceiveDate() != 0 {
			x.Int(m.GetReceiveDate())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneCallWaiting encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCallDiscarded) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneCallDiscarded, layer)

	switch cmd {
	case Cmd_PhoneCallDiscarded:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallDiscarded))
		// flags
		var flags int32 = 0
		if m.GetNeedRating() == true {
			flags |= 1 << 2
		}
		if m.GetNeedDebug() == true {
			flags |= 1 << 3
		}
		if m.GetVideo() == true {
			flags |= 1 << 6
		}
		if m.GetReason() != nil {
			flags |= 1 << 0
		}
		if m.GetDuration() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetId())
		if m.GetReason() != nil {
			x.Bytes(m.GetReason().Encode(layer))
		}
		if m.GetDuration() != 0 {
			x.Int(m.GetDuration())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneCallDiscarded encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCallDiscardReasonMissed) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneCallDiscardReasonMissed, layer)

	switch cmd {
	case Cmd_PhoneCallDiscardReasonMissed:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallDiscardReasonMissed))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneCallDiscardReasonMissed encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCallDiscardReasonDisconnect) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneCallDiscardReasonDisconnect, layer)

	switch cmd {
	case Cmd_PhoneCallDiscardReasonDisconnect:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallDiscardReasonDisconnect))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneCallDiscardReasonDisconnect encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCallDiscardReasonHangup) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneCallDiscardReasonHangup, layer)

	switch cmd {
	case Cmd_PhoneCallDiscardReasonHangup:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallDiscardReasonHangup))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneCallDiscardReasonHangup encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCallDiscardReasonBusy) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneCallDiscardReasonBusy, layer)

	switch cmd {
	case Cmd_PhoneCallDiscardReasonBusy:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallDiscardReasonBusy))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneCallDiscardReasonBusy encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneCallProtocol) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneCallProtocol, layer)

	switch cmd {
	case Cmd_PhoneCallProtocol:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallProtocol))
		// flags
		var flags int32 = 0
		if m.GetUdpP2P() == true {
			flags |= 1 << 0
		}
		if m.GetUdpReflector() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetMinLayer())
		x.Int(m.GetMaxLayer())

		return x.GetBuf()
	case Cmd_PhoneCallProtocolfc878fc8:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneCallProtocolfc878fc8))
		// flags
		var flags int32 = 0
		if m.GetUdpP2P() == true {
			flags |= 1 << 0
		}
		if m.GetUdpReflector() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetMinLayer())
		x.Int(m.GetMaxLayer())
		x.VectorString(m.GetLibraryVersions())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneCallProtocol encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneConnection) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneConnection, layer)

	switch cmd {
	case Cmd_PhoneConnection:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneConnection))
		x.Long(m.GetId())
		x.String(m.GetIp())
		x.String(m.GetIpv6())
		x.Int(m.GetPort())
		x.StringBytes(m.GetPeerTag())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneConnection encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneConnectionWebrtc) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneConnectionWebrtc, layer)

	switch cmd {
	case Cmd_PhoneConnectionWebrtc:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneConnectionWebrtc))
		// flags
		var flags int32 = 0
		if m.GetTurn() == true {
			flags |= 1 << 0
		}
		if m.GetStun() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.String(m.GetIp())
		x.String(m.GetIpv6())
		x.Int(m.GetPort())
		x.String(m.GetUsername())
		x.String(m.GetPassword())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneConnectionWebrtc encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneGroupCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneGroupCall, layer)

	switch cmd {
	case Cmd_PhoneGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneGroupCall))
		x.Bytes(m.GetCall().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetParticipants())))
		for _, v := range m.GetParticipants() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.String(m.GetParticipantsNextOffset())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneGroupCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoneGroupParticipants) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoneGroupParticipants, layer)

	switch cmd {
	case Cmd_PhoneGroupParticipants:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhoneGroupParticipants))
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetParticipants())))
		for _, v := range m.GetParticipants() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.String(m.GetNextOffset())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetVersion())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhoneGroupParticipants encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhonePhoneCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhonePhoneCall, layer)

	switch cmd {
	case Cmd_PhonePhoneCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhonePhoneCall))
		x.Bytes(m.GetPhoneCall().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PhonePhoneCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotoEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhotoEmpty, layer)

	switch cmd {
	case Cmd_PhotoEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotoEmpty))
		x.Long(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhotoEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhoto, layer)

	switch cmd {
	case Cmd_Photo:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Photo))
		// flags
		var flags int32 = 0
		if m.GetHasStickers() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSizes())))
		for _, v := range m.GetSizes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_Photo9c477dd8:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Photo9c477dd8))
		// flags
		var flags int32 = 0
		if m.GetHasStickers() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())
		x.Int(m.GetDate())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSizes())))
		for _, v := range m.GetSizes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_Photocded42fe:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Photocded42fe))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetDate())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSizes())))
		for _, v := range m.GetSizes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_Photod07504a5:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Photod07504a5))
		// flags
		var flags int32 = 0
		if m.GetHasStickers() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())
		x.Int(m.GetDate())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSizes())))
		for _, v := range m.GetSizes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetDcId())

		return x.GetBuf()
	case Cmd_Photofb197a65:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Photofb197a65))
		// flags
		var flags int32 = 0
		if m.GetHasStickers() == true {
			flags |= 1 << 0
		}
		if len(m.GetVideoSizes()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.StringBytes(m.GetFileReference())
		x.Int(m.GetDate())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSizes())))
		for _, v := range m.GetSizes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if len(m.GetVideoSizes()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetVideoSizes())))
			for _, v := range m.GetVideoSizes() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		x.Int(m.GetDcId())

		return x.GetBuf()

	default:
		log.Errorf("invalid Photo encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotoSizeEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhotoSizeEmpty, layer)

	switch cmd {
	case Cmd_PhotoSizeEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotoSizeEmpty))
		x.String(m.GetType())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhotoSizeEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotoSize) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhotoSize, layer)

	switch cmd {
	case Cmd_PhotoSize:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotoSize))
		x.String(m.GetType())
		x.Bytes(m.GetLocation().Encode(layer))
		x.Int(m.GetW())
		x.Int(m.GetH())
		x.Int(m.GetSize_())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhotoSize encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotoCachedSize) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhotoCachedSize, layer)

	switch cmd {
	case Cmd_PhotoCachedSize:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotoCachedSize))
		x.String(m.GetType())
		x.Bytes(m.GetLocation().Encode(layer))
		x.Int(m.GetW())
		x.Int(m.GetH())
		x.StringBytes(m.GetBytes())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhotoCachedSize encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotoStrippedSize) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhotoStrippedSize, layer)

	switch cmd {
	case Cmd_PhotoStrippedSize:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotoStrippedSize))
		x.String(m.GetType())
		x.StringBytes(m.GetBytes())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhotoStrippedSize encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotoSizeProgressive) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhotoSizeProgressive, layer)

	switch cmd {
	case Cmd_PhotoSizeProgressive:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotoSizeProgressive))
		x.String(m.GetType())
		x.Bytes(m.GetLocation().Encode(layer))
		x.Int(m.GetW())
		x.Int(m.GetH())
		x.VectorInt(m.GetSizes())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhotoSizeProgressive encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotoPathSize) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhotoPathSize, layer)

	switch cmd {
	case Cmd_PhotoPathSize:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotoPathSize))
		x.String(m.GetType())
		x.StringBytes(m.GetBytes())

		return x.GetBuf()

	default:
		log.Errorf("invalid PhotoPathSize encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotosPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhotosPhoto, layer)

	switch cmd {
	case Cmd_PhotosPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotosPhoto))
		x.Bytes(m.GetPhoto().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PhotosPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotosPhotos) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhotosPhotos, layer)

	switch cmd {
	case Cmd_PhotosPhotos:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotosPhotos))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPhotos())))
		for _, v := range m.GetPhotos() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PhotosPhotos encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPhotosPhotosSlice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPhotosPhotosSlice, layer)

	switch cmd {
	case Cmd_PhotosPhotosSlice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PhotosPhotosSlice))
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPhotos())))
		for _, v := range m.GetPhotos() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PhotosPhotosSlice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPoll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPoll, layer)

	switch cmd {
	case Cmd_Poll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Poll))
		x.Long(m.GetId())
		// flags
		var flags int32 = 0
		if m.GetClosed() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetQuestion())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAnswers())))
		for _, v := range m.GetAnswers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_Poll86e18161:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Poll86e18161))
		x.Long(m.GetId())
		// flags
		var flags int32 = 0
		if m.GetClosed() == true {
			flags |= 1 << 0
		}
		if m.GetPublicVoters() == true {
			flags |= 1 << 1
		}
		if m.GetMultipleChoice() == true {
			flags |= 1 << 2
		}
		if m.GetQuiz() == true {
			flags |= 1 << 3
		}
		if m.GetClosePeriod() != 0 {
			flags |= 1 << 4
		}
		if m.GetCloseDate() != 0 {
			flags |= 1 << 5
		}
		x.Int(flags)

		x.String(m.GetQuestion())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAnswers())))
		for _, v := range m.GetAnswers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		if m.GetClosePeriod() != 0 {
			x.Int(m.GetClosePeriod())
		}
		if m.GetCloseDate() != 0 {
			x.Int(m.GetCloseDate())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid Poll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPollAnswer) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPollAnswer, layer)

	switch cmd {
	case Cmd_PollAnswer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PollAnswer))
		x.String(m.GetText())
		x.StringBytes(m.GetOption())

		return x.GetBuf()

	default:
		log.Errorf("invalid PollAnswer encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPollAnswerVoters) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPollAnswerVoters, layer)

	switch cmd {
	case Cmd_PollAnswerVoters:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PollAnswerVoters))
		// flags
		var flags int32 = 0
		if m.GetChosen() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.StringBytes(m.GetOption())
		x.Int(m.GetVoters())

		return x.GetBuf()

	default:
		log.Errorf("invalid PollAnswerVoters encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPollResults) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPollResults, layer)

	switch cmd {
	case Cmd_PollResults:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PollResults))
		// flags
		var flags int32 = 0
		if m.GetMin() == true {
			flags |= 1 << 0
		}
		if len(m.GetResults()) > 0 {
			flags |= 1 << 1
		}
		if m.GetTotalVoters() != 0 {
			flags |= 1 << 2
		}
		x.Int(flags)

		if len(m.GetResults()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetResults())))
			for _, v := range m.GetResults() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetTotalVoters() != 0 {
			x.Int(m.GetTotalVoters())
		}

		return x.GetBuf()
	case Cmd_PollResultsbadcc1a3:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PollResultsbadcc1a3))
		// flags
		var flags int32 = 0
		if m.GetMin() == true {
			flags |= 1 << 0
		}
		if len(m.GetResults()) > 0 {
			flags |= 1 << 1
		}
		if m.GetTotalVoters() != 0 {
			flags |= 1 << 2
		}
		if len(m.GetRecentVoters()) > 0 {
			flags |= 1 << 3
		}
		if len(m.GetSolution()) > 0 {
			flags |= 1 << 4
		}
		if len(m.GetSolutionEntities()) > 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		if len(m.GetResults()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetResults())))
			for _, v := range m.GetResults() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetTotalVoters() != 0 {
			x.Int(m.GetTotalVoters())
		}
		if len(m.GetRecentVoters()) > 0 {
			x.VectorInt(m.GetRecentVoters())

		}
		if len(m.GetSolution()) > 0 {
			x.String(m.GetSolution())
		}
		if len(m.GetSolutionEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetSolutionEntities())))
			for _, v := range m.GetSolutionEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_PollResultsc87024a2:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PollResultsc87024a2))
		// flags
		var flags int32 = 0
		if m.GetMin() == true {
			flags |= 1 << 0
		}
		if len(m.GetResults()) > 0 {
			flags |= 1 << 1
		}
		if m.GetTotalVoters() != 0 {
			flags |= 1 << 2
		}
		if len(m.GetRecentVoters()) > 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		if len(m.GetResults()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetResults())))
			for _, v := range m.GetResults() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetTotalVoters() != 0 {
			x.Int(m.GetTotalVoters())
		}
		if len(m.GetRecentVoters()) > 0 {
			x.VectorInt(m.GetRecentVoters())

		}

		return x.GetBuf()

	default:
		log.Errorf("invalid PollResults encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPong) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPong, layer)

	switch cmd {
	case Cmd_Pong:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Pong))
		x.Long(m.GetMsgId())
		x.Long(m.GetPingId())

		return x.GetBuf()

	default:
		log.Errorf("invalid Pong encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPopularContact) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPopularContact, layer)

	switch cmd {
	case Cmd_PopularContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PopularContact))
		x.Long(m.GetClientId())
		x.Int(m.GetImporters())

		return x.GetBuf()

	default:
		log.Errorf("invalid PopularContact encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPostAddress) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPostAddress, layer)

	switch cmd {
	case Cmd_PostAddress:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PostAddress))
		x.String(m.GetStreetLine1())
		x.String(m.GetStreetLine2())
		x.String(m.GetCity())
		x.String(m.GetState())
		x.String(m.GetCountryIso2())
		x.String(m.GetPostCode())

		return x.GetBuf()

	default:
		log.Errorf("invalid PostAddress encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyKeyStatusTimestamp) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyKeyStatusTimestamp, layer)

	switch cmd {
	case Cmd_PrivacyKeyStatusTimestamp:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyKeyStatusTimestamp))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyKeyStatusTimestamp encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyKeyChatInvite) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyKeyChatInvite, layer)

	switch cmd {
	case Cmd_PrivacyKeyChatInvite:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyKeyChatInvite))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyKeyChatInvite encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyKeyPhoneCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyKeyPhoneCall, layer)

	switch cmd {
	case Cmd_PrivacyKeyPhoneCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyKeyPhoneCall))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyKeyPhoneCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyKeyPhoneP2P) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyKeyPhoneP2P, layer)

	switch cmd {
	case Cmd_PrivacyKeyPhoneP2P:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyKeyPhoneP2P))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyKeyPhoneP2P encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyKeyForwards) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyKeyForwards, layer)

	switch cmd {
	case Cmd_PrivacyKeyForwards:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyKeyForwards))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyKeyForwards encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyKeyProfilePhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyKeyProfilePhoto, layer)

	switch cmd {
	case Cmd_PrivacyKeyProfilePhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyKeyProfilePhoto))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyKeyProfilePhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyKeyPhoneNumber) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyKeyPhoneNumber, layer)

	switch cmd {
	case Cmd_PrivacyKeyPhoneNumber:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyKeyPhoneNumber))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyKeyPhoneNumber encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyKeyAddedByPhone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyKeyAddedByPhone, layer)

	switch cmd {
	case Cmd_PrivacyKeyAddedByPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyKeyAddedByPhone))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyKeyAddedByPhone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyValueAllowContacts) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyValueAllowContacts, layer)

	switch cmd {
	case Cmd_PrivacyValueAllowContacts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyValueAllowContacts))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyValueAllowContacts encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyValueAllowAll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyValueAllowAll, layer)

	switch cmd {
	case Cmd_PrivacyValueAllowAll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyValueAllowAll))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyValueAllowAll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyValueAllowUsers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyValueAllowUsers, layer)

	switch cmd {
	case Cmd_PrivacyValueAllowUsers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyValueAllowUsers))
		x.VectorInt(m.GetUsers())

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyValueAllowUsers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyValueDisallowContacts) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyValueDisallowContacts, layer)

	switch cmd {
	case Cmd_PrivacyValueDisallowContacts:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyValueDisallowContacts))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyValueDisallowContacts encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyValueDisallowAll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyValueDisallowAll, layer)

	switch cmd {
	case Cmd_PrivacyValueDisallowAll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyValueDisallowAll))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyValueDisallowAll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyValueDisallowUsers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyValueDisallowUsers, layer)

	switch cmd {
	case Cmd_PrivacyValueDisallowUsers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyValueDisallowUsers))
		x.VectorInt(m.GetUsers())

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyValueDisallowUsers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyValueAllowChatParticipants) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyValueAllowChatParticipants, layer)

	switch cmd {
	case Cmd_PrivacyValueAllowChatParticipants:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyValueAllowChatParticipants))
		x.VectorInt(m.GetChats())

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyValueAllowChatParticipants encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLPrivacyValueDisallowChatParticipants) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassPrivacyValueDisallowChatParticipants, layer)

	switch cmd {
	case Cmd_PrivacyValueDisallowChatParticipants:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_PrivacyValueDisallowChatParticipants))
		x.VectorInt(m.GetChats())

		return x.GetBuf()

	default:
		log.Errorf("invalid PrivacyValueDisallowChatParticipants encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLReceivedNotifyMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassReceivedNotifyMessage, layer)

	switch cmd {
	case Cmd_ReceivedNotifyMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ReceivedNotifyMessage))
		x.Int(m.GetId())
		x.Int(m.GetFlags())

		return x.GetBuf()

	default:
		log.Errorf("invalid ReceivedNotifyMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLRecentMeUrlUnknown) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassRecentMeUrlUnknown, layer)

	switch cmd {
	case Cmd_RecentMeUrlUnknown:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_RecentMeUrlUnknown))
		x.String(m.GetUrl())

		return x.GetBuf()

	default:
		log.Errorf("invalid RecentMeUrlUnknown encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLRecentMeUrlUser) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassRecentMeUrlUser, layer)

	switch cmd {
	case Cmd_RecentMeUrlUser:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_RecentMeUrlUser))
		x.String(m.GetUrl())
		x.Int(m.GetUserId())

		return x.GetBuf()

	default:
		log.Errorf("invalid RecentMeUrlUser encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLRecentMeUrlChat) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassRecentMeUrlChat, layer)

	switch cmd {
	case Cmd_RecentMeUrlChat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_RecentMeUrlChat))
		x.String(m.GetUrl())
		x.Int(m.GetChatId())

		return x.GetBuf()

	default:
		log.Errorf("invalid RecentMeUrlChat encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLRecentMeUrlChatInvite) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassRecentMeUrlChatInvite, layer)

	switch cmd {
	case Cmd_RecentMeUrlChatInvite:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_RecentMeUrlChatInvite))
		x.String(m.GetUrl())
		x.Bytes(m.GetChatInvite().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid RecentMeUrlChatInvite encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLRecentMeUrlStickerSet) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassRecentMeUrlStickerSet, layer)

	switch cmd {
	case Cmd_RecentMeUrlStickerSet:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_RecentMeUrlStickerSet))
		x.String(m.GetUrl())
		x.Bytes(m.GetSet().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid RecentMeUrlStickerSet encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLReplyKeyboardHide) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassReplyKeyboardHide, layer)

	switch cmd {
	case Cmd_ReplyKeyboardHide:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ReplyKeyboardHide))
		// flags
		var flags int32 = 0
		if m.GetSelective() == true {
			flags |= 1 << 2
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("invalid ReplyKeyboardHide encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLReplyKeyboardForceReply) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassReplyKeyboardForceReply, layer)

	switch cmd {
	case Cmd_ReplyKeyboardForceReply:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ReplyKeyboardForceReply))
		// flags
		var flags int32 = 0
		if m.GetSingleUse() == true {
			flags |= 1 << 1
		}
		if m.GetSelective() == true {
			flags |= 1 << 2
		}
		x.Int(flags)

		return x.GetBuf()

	default:
		log.Errorf("invalid ReplyKeyboardForceReply encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLReplyKeyboardMarkup) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassReplyKeyboardMarkup, layer)

	switch cmd {
	case Cmd_ReplyKeyboardMarkup:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ReplyKeyboardMarkup))
		// flags
		var flags int32 = 0
		if m.GetResize() == true {
			flags |= 1 << 0
		}
		if m.GetSingleUse() == true {
			flags |= 1 << 1
		}
		if m.GetSelective() == true {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetRows())))
		for _, v := range m.GetRows() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ReplyKeyboardMarkup encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLReplyInlineMarkup) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassReplyInlineMarkup, layer)

	switch cmd {
	case Cmd_ReplyInlineMarkup:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ReplyInlineMarkup))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetRows())))
		for _, v := range m.GetRows() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ReplyInlineMarkup encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputReportReasonSpam) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputReportReasonSpam, layer)

	switch cmd {
	case Cmd_InputReportReasonSpam:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputReportReasonSpam))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputReportReasonSpam encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputReportReasonViolence) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputReportReasonViolence, layer)

	switch cmd {
	case Cmd_InputReportReasonViolence:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputReportReasonViolence))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputReportReasonViolence encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputReportReasonPornography) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputReportReasonPornography, layer)

	switch cmd {
	case Cmd_InputReportReasonPornography:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputReportReasonPornography))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputReportReasonPornography encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputReportReasonOther) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputReportReasonOther, layer)

	switch cmd {
	case Cmd_InputReportReasonOther:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputReportReasonOther))
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid InputReportReasonOther encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputReportReasonCopyright) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputReportReasonCopyright, layer)

	switch cmd {
	case Cmd_InputReportReasonCopyright:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputReportReasonCopyright))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputReportReasonCopyright encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputReportReasonChildAbuse) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputReportReasonChildAbuse, layer)

	switch cmd {
	case Cmd_InputReportReasonChildAbuse:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputReportReasonChildAbuse))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputReportReasonChildAbuse encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLInputReportReasonGeoIrrelevant) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassInputReportReasonGeoIrrelevant, layer)

	switch cmd {
	case Cmd_InputReportReasonGeoIrrelevant:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_InputReportReasonGeoIrrelevant))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid InputReportReasonGeoIrrelevant encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLResPQ) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassResPQ, layer)

	switch cmd {
	case Cmd_ResPQ:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ResPQ))
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.String(m.GetPq())
		x.VectorLong(m.GetServerPublicKeyFingerprints())

		return x.GetBuf()

	default:
		log.Errorf("invalid ResPQ encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLRestrictionReason) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassRestrictionReason, layer)

	switch cmd {
	case Cmd_RestrictionReason:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_RestrictionReason))
		x.String(m.GetPlatform())
		x.String(m.GetReason())
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid RestrictionReason encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextEmpty, layer)

	switch cmd {
	case Cmd_TextEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid TextEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextPlain) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextPlain, layer)

	switch cmd {
	case Cmd_TextPlain:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextPlain))
		x.String(m.GetText744694E071())

		return x.GetBuf()

	default:
		log.Errorf("invalid TextPlain encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextBold) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextBold, layer)

	switch cmd {
	case Cmd_TextBold:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextBold))
		x.Bytes(m.GetText6724ABC471().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid TextBold encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextItalic) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextItalic, layer)

	switch cmd {
	case Cmd_TextItalic:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextItalic))
		x.Bytes(m.GetText6724ABC471().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid TextItalic encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextUnderline) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextUnderline, layer)

	switch cmd {
	case Cmd_TextUnderline:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextUnderline))
		x.Bytes(m.GetText6724ABC471().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid TextUnderline encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextStrike) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextStrike, layer)

	switch cmd {
	case Cmd_TextStrike:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextStrike))
		x.Bytes(m.GetText6724ABC471().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid TextStrike encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextFixed) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextFixed, layer)

	switch cmd {
	case Cmd_TextFixed:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextFixed))
		x.Bytes(m.GetText6724ABC471().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid TextFixed encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextUrl) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextUrl, layer)

	switch cmd {
	case Cmd_TextUrl:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextUrl))
		x.Bytes(m.GetText6724ABC471().Encode(layer))
		x.String(m.GetUrl())
		x.Long(m.GetWebpageId())

		return x.GetBuf()

	default:
		log.Errorf("invalid TextUrl encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextEmail) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextEmail, layer)

	switch cmd {
	case Cmd_TextEmail:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextEmail))
		x.Bytes(m.GetText6724ABC471().Encode(layer))
		x.String(m.GetEmail())

		return x.GetBuf()

	default:
		log.Errorf("invalid TextEmail encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextConcat) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextConcat, layer)

	switch cmd {
	case Cmd_TextConcat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextConcat))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetTexts())))
		for _, v := range m.GetTexts() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid TextConcat encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextSubscript) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextSubscript, layer)

	switch cmd {
	case Cmd_TextSubscript:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextSubscript))
		x.Bytes(m.GetText6724ABC471().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid TextSubscript encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextSuperscript) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextSuperscript, layer)

	switch cmd {
	case Cmd_TextSuperscript:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextSuperscript))
		x.Bytes(m.GetText6724ABC471().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid TextSuperscript encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextMarked) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextMarked, layer)

	switch cmd {
	case Cmd_TextMarked:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextMarked))
		x.Bytes(m.GetText6724ABC471().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid TextMarked encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextPhone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextPhone, layer)

	switch cmd {
	case Cmd_TextPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextPhone))
		x.Bytes(m.GetText6724ABC471().Encode(layer))
		x.String(m.GetPhone())

		return x.GetBuf()

	default:
		log.Errorf("invalid TextPhone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextImage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextImage, layer)

	switch cmd {
	case Cmd_TextImage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextImage))
		x.Long(m.GetDocumentId())
		x.Int(m.GetW())
		x.Int(m.GetH())

		return x.GetBuf()

	default:
		log.Errorf("invalid TextImage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTextAnchor) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTextAnchor, layer)

	switch cmd {
	case Cmd_TextAnchor:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TextAnchor))
		x.Bytes(m.GetText6724ABC471().Encode(layer))
		x.String(m.GetName())

		return x.GetBuf()

	default:
		log.Errorf("invalid TextAnchor encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLRpcAnswerUnknown) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassRpcAnswerUnknown, layer)

	switch cmd {
	case Cmd_RpcAnswerUnknown:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_RpcAnswerUnknown))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid RpcAnswerUnknown encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLRpcAnswerDroppedRunning) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassRpcAnswerDroppedRunning, layer)

	switch cmd {
	case Cmd_RpcAnswerDroppedRunning:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_RpcAnswerDroppedRunning))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid RpcAnswerDroppedRunning encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLRpcAnswerDropped) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassRpcAnswerDropped, layer)

	switch cmd {
	case Cmd_RpcAnswerDropped:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_RpcAnswerDropped))
		x.Long(m.GetMsgId())
		x.Int(m.GetSeqNo())
		x.Int(m.GetBytes())

		return x.GetBuf()

	default:
		log.Errorf("invalid RpcAnswerDropped encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLRpcError) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassRpcError, layer)

	switch cmd {
	case Cmd_RpcError:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_RpcError))
		x.Int(m.GetErrorCode())
		x.String(m.GetErrorMessage())

		return x.GetBuf()

	default:
		log.Errorf("invalid RpcError encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSavedPhoneContact) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSavedPhoneContact, layer)

	switch cmd {
	case Cmd_SavedPhoneContact:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SavedPhoneContact))
		x.String(m.GetPhone())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid SavedPhoneContact encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureCredentialsEncrypted) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureCredentialsEncrypted, layer)

	switch cmd {
	case Cmd_SecureCredentialsEncrypted:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureCredentialsEncrypted))
		x.StringBytes(m.GetData())
		x.StringBytes(m.GetHash())
		x.StringBytes(m.GetSecret())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureCredentialsEncrypted encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureData) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureData, layer)

	switch cmd {
	case Cmd_SecureData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureData))
		x.StringBytes(m.GetData())
		x.StringBytes(m.GetDataHash())
		x.StringBytes(m.GetSecret())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureData encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureFileEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureFileEmpty, layer)

	switch cmd {
	case Cmd_SecureFileEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureFileEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureFileEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureFile, layer)

	switch cmd {
	case Cmd_SecureFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureFile))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.Int(m.GetSize_())
		x.Int(m.GetDcId())
		x.Int(m.GetDate())
		x.StringBytes(m.GetFileHash())
		x.StringBytes(m.GetSecret())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecurePasswordKdfAlgoUnknown) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecurePasswordKdfAlgoUnknown, layer)

	switch cmd {
	case Cmd_SecurePasswordKdfAlgoUnknown:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecurePasswordKdfAlgoUnknown))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecurePasswordKdfAlgoUnknown encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000, layer)

	switch cmd {
	case Cmd_SecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000))
		x.StringBytes(m.GetSalt())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000 encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecurePasswordKdfAlgoSHA512) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecurePasswordKdfAlgoSHA512, layer)

	switch cmd {
	case Cmd_SecurePasswordKdfAlgoSHA512:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecurePasswordKdfAlgoSHA512))
		x.StringBytes(m.GetSalt())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecurePasswordKdfAlgoSHA512 encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecurePlainPhone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecurePlainPhone, layer)

	switch cmd {
	case Cmd_SecurePlainPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecurePlainPhone))
		x.String(m.GetPhone())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecurePlainPhone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecurePlainEmail) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecurePlainEmail, layer)

	switch cmd {
	case Cmd_SecurePlainEmail:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecurePlainEmail))
		x.String(m.GetEmail())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecurePlainEmail encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureRequiredType) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureRequiredType, layer)

	switch cmd {
	case Cmd_SecureRequiredType:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureRequiredType))
		// flags
		var flags int32 = 0
		if m.GetNativeNames() == true {
			flags |= 1 << 0
		}
		if m.GetSelfieRequired() == true {
			flags |= 1 << 1
		}
		if m.GetTranslationRequired() == true {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Bytes(m.GetType().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureRequiredType encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureRequiredTypeOneOf) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureRequiredTypeOneOf, layer)

	switch cmd {
	case Cmd_SecureRequiredTypeOneOf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureRequiredTypeOneOf))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetTypes())))
		for _, v := range m.GetTypes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureRequiredTypeOneOf encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureSecretSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureSecretSettings, layer)

	switch cmd {
	case Cmd_SecureSecretSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureSecretSettings))
		x.Bytes(m.GetSecureAlgo().Encode(layer))
		x.StringBytes(m.GetSecureSecret())
		x.Long(m.GetSecureSecretId())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureSecretSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValue) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValue, layer)

	switch cmd {
	case Cmd_SecureValue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValue))
		// flags
		var flags int32 = 0
		if m.GetData() != nil {
			flags |= 1 << 0
		}
		if m.GetFrontSide() != nil {
			flags |= 1 << 1
		}
		if m.GetReverseSide() != nil {
			flags |= 1 << 2
		}
		if m.GetSelfie() != nil {
			flags |= 1 << 3
		}
		if len(m.GetTranslation()) > 0 {
			flags |= 1 << 6
		}
		if len(m.GetFiles()) > 0 {
			flags |= 1 << 4
		}
		if m.GetPlainData() != nil {
			flags |= 1 << 5
		}
		x.Int(flags)

		x.Bytes(m.GetType().Encode(layer))
		if m.GetData() != nil {
			x.Bytes(m.GetData().Encode(layer))
		}
		if m.GetFrontSide() != nil {
			x.Bytes(m.GetFrontSide().Encode(layer))
		}
		if m.GetReverseSide() != nil {
			x.Bytes(m.GetReverseSide().Encode(layer))
		}
		if m.GetSelfie() != nil {
			x.Bytes(m.GetSelfie().Encode(layer))
		}
		if len(m.GetTranslation()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetTranslation())))
			for _, v := range m.GetTranslation() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if len(m.GetFiles()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetFiles())))
			for _, v := range m.GetFiles() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetPlainData() != nil {
			x.Bytes(m.GetPlainData().Encode(layer))
		}
		x.StringBytes(m.GetHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValue encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueErrorData) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueErrorData, layer)

	switch cmd {
	case Cmd_SecureValueErrorData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueErrorData))
		x.Bytes(m.GetType().Encode(layer))
		x.StringBytes(m.GetDataHash())
		x.String(m.GetField())
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueErrorData encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueErrorFrontSide) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueErrorFrontSide, layer)

	switch cmd {
	case Cmd_SecureValueErrorFrontSide:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueErrorFrontSide))
		x.Bytes(m.GetType().Encode(layer))
		x.StringBytes(m.GetFileHashBE3DFA85())
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueErrorFrontSide encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueErrorReverseSide) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueErrorReverseSide, layer)

	switch cmd {
	case Cmd_SecureValueErrorReverseSide:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueErrorReverseSide))
		x.Bytes(m.GetType().Encode(layer))
		x.StringBytes(m.GetFileHashBE3DFA85())
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueErrorReverseSide encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueErrorSelfie) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueErrorSelfie, layer)

	switch cmd {
	case Cmd_SecureValueErrorSelfie:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueErrorSelfie))
		x.Bytes(m.GetType().Encode(layer))
		x.StringBytes(m.GetFileHashBE3DFA85())
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueErrorSelfie encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueErrorFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueErrorFile, layer)

	switch cmd {
	case Cmd_SecureValueErrorFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueErrorFile))
		x.Bytes(m.GetType().Encode(layer))
		x.StringBytes(m.GetFileHashBE3DFA85())
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueErrorFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueErrorFiles) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueErrorFiles, layer)

	switch cmd {
	case Cmd_SecureValueErrorFiles:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueErrorFiles))
		x.Bytes(m.GetType().Encode(layer))

		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueErrorFiles encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueError) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueError, layer)

	switch cmd {
	case Cmd_SecureValueError:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueError))
		x.Bytes(m.GetType().Encode(layer))
		x.StringBytes(m.GetHash())
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueError encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueErrorTranslationFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueErrorTranslationFile, layer)

	switch cmd {
	case Cmd_SecureValueErrorTranslationFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueErrorTranslationFile))
		x.Bytes(m.GetType().Encode(layer))
		x.StringBytes(m.GetFileHashBE3DFA85())
		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueErrorTranslationFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueErrorTranslationFiles) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueErrorTranslationFiles, layer)

	switch cmd {
	case Cmd_SecureValueErrorTranslationFiles:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueErrorTranslationFiles))
		x.Bytes(m.GetType().Encode(layer))

		x.String(m.GetText())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueErrorTranslationFiles encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueHash) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueHash, layer)

	switch cmd {
	case Cmd_SecureValueHash:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueHash))
		x.Bytes(m.GetType().Encode(layer))
		x.StringBytes(m.GetHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueHash encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypePersonalDetails) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypePersonalDetails, layer)

	switch cmd {
	case Cmd_SecureValueTypePersonalDetails:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypePersonalDetails))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypePersonalDetails encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypePassport) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypePassport, layer)

	switch cmd {
	case Cmd_SecureValueTypePassport:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypePassport))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypePassport encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypeDriverLicense) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypeDriverLicense, layer)

	switch cmd {
	case Cmd_SecureValueTypeDriverLicense:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypeDriverLicense))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypeDriverLicense encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypeIdentityCard) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypeIdentityCard, layer)

	switch cmd {
	case Cmd_SecureValueTypeIdentityCard:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypeIdentityCard))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypeIdentityCard encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypeInternalPassport) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypeInternalPassport, layer)

	switch cmd {
	case Cmd_SecureValueTypeInternalPassport:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypeInternalPassport))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypeInternalPassport encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypeAddress) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypeAddress, layer)

	switch cmd {
	case Cmd_SecureValueTypeAddress:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypeAddress))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypeAddress encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypeUtilityBill) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypeUtilityBill, layer)

	switch cmd {
	case Cmd_SecureValueTypeUtilityBill:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypeUtilityBill))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypeUtilityBill encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypeBankStatement) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypeBankStatement, layer)

	switch cmd {
	case Cmd_SecureValueTypeBankStatement:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypeBankStatement))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypeBankStatement encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypeRentalAgreement) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypeRentalAgreement, layer)

	switch cmd {
	case Cmd_SecureValueTypeRentalAgreement:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypeRentalAgreement))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypeRentalAgreement encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypePassportRegistration) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypePassportRegistration, layer)

	switch cmd {
	case Cmd_SecureValueTypePassportRegistration:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypePassportRegistration))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypePassportRegistration encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypeTemporaryRegistration) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypeTemporaryRegistration, layer)

	switch cmd {
	case Cmd_SecureValueTypeTemporaryRegistration:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypeTemporaryRegistration))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypeTemporaryRegistration encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypePhone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypePhone, layer)

	switch cmd {
	case Cmd_SecureValueTypePhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypePhone))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypePhone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSecureValueTypeEmail) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSecureValueTypeEmail, layer)

	switch cmd {
	case Cmd_SecureValueTypeEmail:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SecureValueTypeEmail))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SecureValueTypeEmail encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageTypingAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageTypingAction, layer)

	switch cmd {
	case Cmd_SendMessageTypingAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageTypingAction))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageTypingAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageCancelAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageCancelAction, layer)

	switch cmd {
	case Cmd_SendMessageCancelAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageCancelAction))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageCancelAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageRecordVideoAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageRecordVideoAction, layer)

	switch cmd {
	case Cmd_SendMessageRecordVideoAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageRecordVideoAction))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageRecordVideoAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageUploadVideoAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageUploadVideoAction, layer)

	switch cmd {
	case Cmd_SendMessageUploadVideoAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageUploadVideoAction))
		x.Int(m.GetProgress())

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageUploadVideoAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageRecordAudioAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageRecordAudioAction, layer)

	switch cmd {
	case Cmd_SendMessageRecordAudioAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageRecordAudioAction))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageRecordAudioAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageUploadAudioAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageUploadAudioAction, layer)

	switch cmd {
	case Cmd_SendMessageUploadAudioAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageUploadAudioAction))
		x.Int(m.GetProgress())

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageUploadAudioAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageUploadPhotoAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageUploadPhotoAction, layer)

	switch cmd {
	case Cmd_SendMessageUploadPhotoAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageUploadPhotoAction))
		x.Int(m.GetProgress())

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageUploadPhotoAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageUploadDocumentAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageUploadDocumentAction, layer)

	switch cmd {
	case Cmd_SendMessageUploadDocumentAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageUploadDocumentAction))
		x.Int(m.GetProgress())

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageUploadDocumentAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageGeoLocationAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageGeoLocationAction, layer)

	switch cmd {
	case Cmd_SendMessageGeoLocationAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageGeoLocationAction))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageGeoLocationAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageChooseContactAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageChooseContactAction, layer)

	switch cmd {
	case Cmd_SendMessageChooseContactAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageChooseContactAction))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageChooseContactAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageGamePlayAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageGamePlayAction, layer)

	switch cmd {
	case Cmd_SendMessageGamePlayAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageGamePlayAction))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageGamePlayAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageRecordRoundAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageRecordRoundAction, layer)

	switch cmd {
	case Cmd_SendMessageRecordRoundAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageRecordRoundAction))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageRecordRoundAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSendMessageUploadRoundAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSendMessageUploadRoundAction, layer)

	switch cmd {
	case Cmd_SendMessageUploadRoundAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SendMessageUploadRoundAction))
		x.Int(m.GetProgress())

		return x.GetBuf()

	default:
		log.Errorf("invalid SendMessageUploadRoundAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLSpeakingInGroupCallAction) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassSpeakingInGroupCallAction, layer)

	switch cmd {
	case Cmd_SpeakingInGroupCallAction:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_SpeakingInGroupCallAction))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid SpeakingInGroupCallAction encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLServer_DHInnerData) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassServer_DHInnerData, layer)

	switch cmd {
	case Cmd_Server_DHInnerData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Server_DHInnerData))
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.Int(m.GetG())
		x.String(m.GetDhPrime())
		x.String(m.GetGA())
		x.Int(m.GetServerTime())

		return x.GetBuf()

	default:
		log.Errorf("invalid Server_DHInnerData encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLServer_DHParamsFail) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassServer_DHParamsFail, layer)

	switch cmd {
	case Cmd_Server_DHParamsFail:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Server_DHParamsFail))
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.Bytes(m.GetNewNonceHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid Server_DHParamsFail encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLServer_DHParamsOk) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassServer_DHParamsOk, layer)

	switch cmd {
	case Cmd_Server_DHParamsOk:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Server_DHParamsOk))
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.String(m.GetEncryptedAnswer())

		return x.GetBuf()

	default:
		log.Errorf("invalid Server_DHParamsOk encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDhGenOk) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDhGenOk, layer)

	switch cmd {
	case Cmd_DhGenOk:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DhGenOk))
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.Bytes(m.GetNewNonceHash1())

		return x.GetBuf()

	default:
		log.Errorf("invalid DhGenOk encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDhGenRetry) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDhGenRetry, layer)

	switch cmd {
	case Cmd_DhGenRetry:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DhGenRetry))
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.Bytes(m.GetNewNonceHash2())

		return x.GetBuf()

	default:
		log.Errorf("invalid DhGenRetry encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLDhGenFail) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassDhGenFail, layer)

	switch cmd {
	case Cmd_DhGenFail:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_DhGenFail))
		x.Bytes(m.GetNonce())
		x.Bytes(m.GetServerNonce())
		x.Bytes(m.GetNewNonceHash3())

		return x.GetBuf()

	default:
		log.Errorf("invalid DhGenFail encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLShippingOption) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassShippingOption, layer)

	switch cmd {
	case Cmd_ShippingOption:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ShippingOption))
		x.String(m.GetId())
		x.String(m.GetTitle())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPrices())))
		for _, v := range m.GetPrices() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ShippingOption encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsAbsValueAndPrev) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsAbsValueAndPrev, layer)

	switch cmd {
	case Cmd_StatsAbsValueAndPrev:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsAbsValueAndPrev))
		x.Double(m.GetCurrent())
		x.Double(m.GetPrevious())

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsAbsValueAndPrev encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsDateRangeDays) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsDateRangeDays, layer)

	switch cmd {
	case Cmd_StatsDateRangeDays:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsDateRangeDays))
		x.Int(m.GetMinDate())
		x.Int(m.GetMaxDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsDateRangeDays encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsGraphAsync) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsGraphAsync, layer)

	switch cmd {
	case Cmd_StatsGraphAsync:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsGraphAsync))
		x.String(m.GetToken())

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsGraphAsync encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsGraphError) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsGraphError, layer)

	switch cmd {
	case Cmd_StatsGraphError:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsGraphError))
		x.String(m.GetError())

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsGraphError encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsGraph) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsGraph, layer)

	switch cmd {
	case Cmd_StatsGraph:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsGraph))
		// flags
		var flags int32 = 0
		if len(m.GetZoomToken()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetJson().Encode(layer))
		if len(m.GetZoomToken()) > 0 {
			x.String(m.GetZoomToken())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsGraph encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsGroupTopAdmin) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsGroupTopAdmin, layer)

	switch cmd {
	case Cmd_StatsGroupTopAdmin:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsGroupTopAdmin))
		x.Int(m.GetUserId())
		x.Int(m.GetDeleted())
		x.Int(m.GetKicked())
		x.Int(m.GetBanned())

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsGroupTopAdmin encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsGroupTopInviter) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsGroupTopInviter, layer)

	switch cmd {
	case Cmd_StatsGroupTopInviter:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsGroupTopInviter))
		x.Int(m.GetUserId())
		x.Int(m.GetInvitations())

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsGroupTopInviter encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsGroupTopPoster) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsGroupTopPoster, layer)

	switch cmd {
	case Cmd_StatsGroupTopPoster:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsGroupTopPoster))
		x.Int(m.GetUserId())
		x.Int(m.GetMessages())
		x.Int(m.GetAvgChars())

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsGroupTopPoster encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsPercentValue) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsPercentValue, layer)

	switch cmd {
	case Cmd_StatsPercentValue:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsPercentValue))
		x.Double(m.GetPart())
		x.Double(m.GetTotal())

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsPercentValue encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsURL) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsURL, layer)

	switch cmd {
	case Cmd_StatsURL:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsURL))
		x.String(m.GetUrl())

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsURL encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsBroadcastStats) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsBroadcastStats, layer)

	switch cmd {
	case Cmd_StatsBroadcastStats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsBroadcastStats))
		x.Bytes(m.GetPeriod().Encode(layer))
		x.Bytes(m.GetFollowers().Encode(layer))
		x.Bytes(m.GetViewsPerPost().Encode(layer))
		x.Bytes(m.GetSharesPerPost().Encode(layer))
		x.Bytes(m.GetEnabledNotifications().Encode(layer))
		x.Bytes(m.GetGrowthGraph().Encode(layer))
		x.Bytes(m.GetFollowersGraph().Encode(layer))
		x.Bytes(m.GetMuteGraph().Encode(layer))
		x.Bytes(m.GetTopHoursGraph().Encode(layer))
		x.Bytes(m.GetInteractionsGraph().Encode(layer))
		x.Bytes(m.GetIvInteractionsGraph().Encode(layer))
		x.Bytes(m.GetViewsBySourceGraph().Encode(layer))
		x.Bytes(m.GetNewFollowersBySourceGraph().Encode(layer))
		x.Bytes(m.GetLanguagesGraph().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetRecentMessageInteractions())))
		for _, v := range m.GetRecentMessageInteractions() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsBroadcastStats encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsMegagroupStats) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsMegagroupStats, layer)

	switch cmd {
	case Cmd_StatsMegagroupStats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsMegagroupStats))
		x.Bytes(m.GetPeriod().Encode(layer))
		x.Bytes(m.GetMembers().Encode(layer))
		x.Bytes(m.GetMessages().Encode(layer))
		x.Bytes(m.GetViewers().Encode(layer))
		x.Bytes(m.GetPosters().Encode(layer))
		x.Bytes(m.GetGrowthGraph().Encode(layer))
		x.Bytes(m.GetMembersGraph().Encode(layer))
		x.Bytes(m.GetNewMembersBySourceGraph().Encode(layer))
		x.Bytes(m.GetLanguagesGraph().Encode(layer))
		x.Bytes(m.GetMessagesGraph().Encode(layer))
		x.Bytes(m.GetActionsGraph().Encode(layer))
		x.Bytes(m.GetTopHoursGraph().Encode(layer))
		x.Bytes(m.GetWeekdaysGraph().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetTopPosters())))
		for _, v := range m.GetTopPosters() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetTopAdmins())))
		for _, v := range m.GetTopAdmins() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetTopInviters())))
		for _, v := range m.GetTopInviters() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsMegagroupStats encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStatsMessageStats) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStatsMessageStats, layer)

	switch cmd {
	case Cmd_StatsMessageStats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StatsMessageStats))
		x.Bytes(m.GetViewsGraph().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid StatsMessageStats encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStickerPack) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStickerPack, layer)

	switch cmd {
	case Cmd_StickerPack:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickerPack))
		x.String(m.GetEmoticon())
		x.VectorLong(m.GetDocuments())

		return x.GetBuf()

	default:
		log.Errorf("invalid StickerPack encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStickerSet) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStickerSet, layer)

	switch cmd {
	case Cmd_StickerSet:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickerSet))
		// flags
		var flags int32 = 0
		if m.GetInstalled() == true {
			flags |= 1 << 0
		}
		if m.GetArchived() == true {
			flags |= 1 << 1
		}
		if m.GetOfficial() == true {
			flags |= 1 << 2
		}
		if m.GetMasks() == true {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.String(m.GetTitle())
		x.String(m.GetShortName())
		x.Int(m.GetCount())
		x.Int(m.GetHash())

		return x.GetBuf()
	case Cmd_StickerSet40e237a8:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickerSet40e237a8))
		// flags
		var flags int32 = 0
		if m.GetArchived() == true {
			flags |= 1 << 1
		}
		if m.GetOfficial() == true {
			flags |= 1 << 2
		}
		if m.GetMasks() == true {
			flags |= 1 << 3
		}
		if m.GetAnimated() == true {
			flags |= 1 << 5
		}
		if m.GetInstalledDate() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetThumbs()) > 0 {
			flags |= 1 << 4
		}
		if m.GetThumbDcId() != 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		if m.GetInstalledDate() != 0 {
			x.Int(m.GetInstalledDate())
		}
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.String(m.GetTitle())
		x.String(m.GetShortName())
		if len(m.GetThumbs()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetThumbs())))
			for _, v := range m.GetThumbs() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetThumbDcId() != 0 {
			x.Int(m.GetThumbDcId())
		}
		x.Int(m.GetCount())
		x.Int(m.GetHash())

		return x.GetBuf()
	case Cmd_StickerSet5585a139:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickerSet5585a139))
		// flags
		var flags int32 = 0
		if m.GetArchived() == true {
			flags |= 1 << 1
		}
		if m.GetOfficial() == true {
			flags |= 1 << 2
		}
		if m.GetMasks() == true {
			flags |= 1 << 3
		}
		if m.GetInstalledDate() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetInstalledDate() != 0 {
			x.Int(m.GetInstalledDate())
		}
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.String(m.GetTitle())
		x.String(m.GetShortName())
		x.Int(m.GetCount())
		x.Int(m.GetHash())

		return x.GetBuf()
	case Cmd_StickerSet6a90bcb7:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickerSet6a90bcb7))
		// flags
		var flags int32 = 0
		if m.GetArchived() == true {
			flags |= 1 << 1
		}
		if m.GetOfficial() == true {
			flags |= 1 << 2
		}
		if m.GetMasks() == true {
			flags |= 1 << 3
		}
		if m.GetInstalledDate() != 0 {
			flags |= 1 << 0
		}
		if m.GetThumb() != nil {
			flags |= 1 << 4
		}
		x.Int(flags)

		if m.GetInstalledDate() != 0 {
			x.Int(m.GetInstalledDate())
		}
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.String(m.GetTitle())
		x.String(m.GetShortName())
		if m.GetThumb() != nil {
			x.Bytes(m.GetThumb().Encode(layer))
		}
		x.Int(m.GetCount())
		x.Int(m.GetHash())

		return x.GetBuf()
	case Cmd_StickerSeteeb46f27:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickerSeteeb46f27))
		// flags
		var flags int32 = 0
		if m.GetArchived() == true {
			flags |= 1 << 1
		}
		if m.GetOfficial() == true {
			flags |= 1 << 2
		}
		if m.GetMasks() == true {
			flags |= 1 << 3
		}
		if m.GetInstalledDate() != 0 {
			flags |= 1 << 0
		}
		if m.GetThumb() != nil {
			flags |= 1 << 4
		}
		if m.GetThumbDcId() != 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		if m.GetInstalledDate() != 0 {
			x.Int(m.GetInstalledDate())
		}
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.String(m.GetTitle())
		x.String(m.GetShortName())
		if m.GetThumb() != nil {
			x.Bytes(m.GetThumb().Encode(layer))
		}
		if m.GetThumbDcId() != 0 {
			x.Int(m.GetThumbDcId())
		}
		x.Int(m.GetCount())
		x.Int(m.GetHash())

		return x.GetBuf()

	default:
		log.Errorf("invalid StickerSet encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStickerSetCovered) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStickerSetCovered, layer)

	switch cmd {
	case Cmd_StickerSetCovered:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickerSetCovered))
		x.Bytes(m.GetSet().Encode(layer))
		x.Bytes(m.GetCover().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid StickerSetCovered encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStickerSetMultiCovered) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStickerSetMultiCovered, layer)

	switch cmd {
	case Cmd_StickerSetMultiCovered:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StickerSetMultiCovered))
		x.Bytes(m.GetSet().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetCovers())))
		for _, v := range m.GetCovers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid StickerSetMultiCovered encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStorageFileUnknown) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStorageFileUnknown, layer)

	switch cmd {
	case Cmd_StorageFileUnknown:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StorageFileUnknown))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid StorageFileUnknown encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStorageFilePartial) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStorageFilePartial, layer)

	switch cmd {
	case Cmd_StorageFilePartial:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StorageFilePartial))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid StorageFilePartial encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStorageFileJpeg) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStorageFileJpeg, layer)

	switch cmd {
	case Cmd_StorageFileJpeg:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StorageFileJpeg))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid StorageFileJpeg encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStorageFileGif) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStorageFileGif, layer)

	switch cmd {
	case Cmd_StorageFileGif:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StorageFileGif))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid StorageFileGif encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStorageFilePng) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStorageFilePng, layer)

	switch cmd {
	case Cmd_StorageFilePng:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StorageFilePng))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid StorageFilePng encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStorageFilePdf) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStorageFilePdf, layer)

	switch cmd {
	case Cmd_StorageFilePdf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StorageFilePdf))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid StorageFilePdf encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStorageFileMp3) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStorageFileMp3, layer)

	switch cmd {
	case Cmd_StorageFileMp3:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StorageFileMp3))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid StorageFileMp3 encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStorageFileMov) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStorageFileMov, layer)

	switch cmd {
	case Cmd_StorageFileMov:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StorageFileMov))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid StorageFileMov encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStorageFileMp4) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStorageFileMp4, layer)

	switch cmd {
	case Cmd_StorageFileMp4:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StorageFileMp4))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid StorageFileMp4 encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLStorageFileWebp) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassStorageFileWebp, layer)

	switch cmd {
	case Cmd_StorageFileWebp:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_StorageFileWebp))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid StorageFileWebp encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLThemeDocumentNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassThemeDocumentNotModified, layer)

	switch cmd {
	case Cmd_ThemeDocumentNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ThemeDocumentNotModified))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid ThemeDocumentNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTheme) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTheme, layer)

	switch cmd {
	case Cmd_Theme:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Theme))
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetDefault() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.String(m.GetSlug())
		x.String(m.GetTitle())
		x.Bytes(m.GetDocument().Encode(layer))

		return x.GetBuf()
	case Cmd_Theme28f1114:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Theme28f1114))
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetDefault() == true {
			flags |= 1 << 1
		}
		if m.GetDocument() != nil {
			flags |= 1 << 2
		}
		if m.GetSettings() != nil {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.String(m.GetSlug())
		x.String(m.GetTitle())
		if m.GetDocument() != nil {
			x.Bytes(m.GetDocument().Encode(layer))
		}
		if m.GetSettings() != nil {
			x.Bytes(m.GetSettings().Encode(layer))
		}
		x.Int(m.GetInstallsCount())

		return x.GetBuf()
	case Cmd_Themef7d90ce0:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Themef7d90ce0))
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetDefault() == true {
			flags |= 1 << 1
		}
		if m.GetDocument() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.Long(m.GetAccessHash())
		x.String(m.GetSlug())
		x.String(m.GetTitle())
		if m.GetDocument() != nil {
			x.Bytes(m.GetDocument().Encode(layer))
		}
		x.Int(m.GetInstallsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid Theme encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLThemeSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassThemeSettings, layer)

	switch cmd {
	case Cmd_ThemeSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_ThemeSettings))
		// flags
		var flags int32 = 0
		if m.GetMessageTopColor() != 0 {
			flags |= 1 << 0
		}
		if m.GetMessageBottomColor() != 0 {
			flags |= 1 << 0
		}
		if m.GetWallpaper() != nil {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Bytes(m.GetBaseTheme().Encode(layer))
		x.Int(m.GetAccentColor())
		if m.GetMessageTopColor() != 0 {
			x.Int(m.GetMessageTopColor())
		}
		if m.GetMessageBottomColor() != 0 {
			x.Int(m.GetMessageBottomColor())
		}
		if m.GetWallpaper() != nil {
			x.Bytes(m.GetWallpaper().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid ThemeSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTopPeer) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTopPeer, layer)

	switch cmd {
	case Cmd_TopPeer:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TopPeer))
		x.Bytes(m.GetPeer().Encode(layer))
		x.Double(m.GetRating())

		return x.GetBuf()

	default:
		log.Errorf("invalid TopPeer encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTopPeerCategoryBotsPM) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTopPeerCategoryBotsPM, layer)

	switch cmd {
	case Cmd_TopPeerCategoryBotsPM:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TopPeerCategoryBotsPM))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid TopPeerCategoryBotsPM encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTopPeerCategoryBotsInline) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTopPeerCategoryBotsInline, layer)

	switch cmd {
	case Cmd_TopPeerCategoryBotsInline:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TopPeerCategoryBotsInline))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid TopPeerCategoryBotsInline encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTopPeerCategoryCorrespondents) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTopPeerCategoryCorrespondents, layer)

	switch cmd {
	case Cmd_TopPeerCategoryCorrespondents:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TopPeerCategoryCorrespondents))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid TopPeerCategoryCorrespondents encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTopPeerCategoryGroups) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTopPeerCategoryGroups, layer)

	switch cmd {
	case Cmd_TopPeerCategoryGroups:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TopPeerCategoryGroups))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid TopPeerCategoryGroups encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTopPeerCategoryChannels) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTopPeerCategoryChannels, layer)

	switch cmd {
	case Cmd_TopPeerCategoryChannels:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TopPeerCategoryChannels))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid TopPeerCategoryChannels encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTopPeerCategoryPhoneCalls) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTopPeerCategoryPhoneCalls, layer)

	switch cmd {
	case Cmd_TopPeerCategoryPhoneCalls:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TopPeerCategoryPhoneCalls))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid TopPeerCategoryPhoneCalls encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTopPeerCategoryForwardUsers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTopPeerCategoryForwardUsers, layer)

	switch cmd {
	case Cmd_TopPeerCategoryForwardUsers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TopPeerCategoryForwardUsers))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid TopPeerCategoryForwardUsers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTopPeerCategoryForwardChats) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTopPeerCategoryForwardChats, layer)

	switch cmd {
	case Cmd_TopPeerCategoryForwardChats:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TopPeerCategoryForwardChats))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid TopPeerCategoryForwardChats encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTopPeerCategoryPeers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTopPeerCategoryPeers, layer)

	switch cmd {
	case Cmd_TopPeerCategoryPeers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_TopPeerCategoryPeers))
		x.Bytes(m.GetCategory().Encode(layer))
		x.Int(m.GetCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPeers())))
		for _, v := range m.GetPeers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid TopPeerCategoryPeers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLTrue) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassTrue, layer)

	switch cmd {
	case Cmd_True:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_True))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid True encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateNewMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateNewMessage, layer)

	switch cmd {
	case Cmd_UpdateNewMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateNewMessage))
		x.Bytes(m.GetMessage1F2B0AFD71().Encode(layer))
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateNewMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateMessageID) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateMessageID, layer)

	switch cmd {
	case Cmd_UpdateMessageID:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateMessageID))
		x.Int(m.GetId4E90BFD671())
		x.Long(m.GetRandomId())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateMessageID encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateDeleteMessages) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateDeleteMessages, layer)

	switch cmd {
	case Cmd_UpdateDeleteMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDeleteMessages))
		x.VectorInt(m.GetMessages())

		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateDeleteMessages encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateUserTyping) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateUserTyping, layer)

	switch cmd {
	case Cmd_UpdateUserTyping:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateUserTyping))
		x.Int(m.GetUserId())
		x.Bytes(m.GetAction().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateUserTyping encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChatUserTyping) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChatUserTyping, layer)

	switch cmd {
	case Cmd_UpdateChatUserTyping:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChatUserTyping))
		x.Int(m.GetChatId())
		x.Int(m.GetUserId())
		x.Bytes(m.GetAction().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChatUserTyping encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChatParticipants) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChatParticipants, layer)

	switch cmd {
	case Cmd_UpdateChatParticipants:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChatParticipants))
		x.Bytes(m.GetParticipants776119871().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChatParticipants encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateUserStatus) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateUserStatus, layer)

	switch cmd {
	case Cmd_UpdateUserStatus:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateUserStatus))
		x.Int(m.GetUserId())
		x.Bytes(m.GetStatus().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateUserStatus encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateUserName) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateUserName, layer)

	switch cmd {
	case Cmd_UpdateUserName:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateUserName))
		x.Int(m.GetUserId())
		x.String(m.GetFirstName())
		x.String(m.GetLastName())
		x.String(m.GetUsername())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateUserName encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateUserPhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateUserPhoto, layer)

	switch cmd {
	case Cmd_UpdateUserPhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateUserPhoto))
		x.Int(m.GetUserId())
		x.Int(m.GetDate())
		x.Bytes(m.GetPhoto().Encode(layer))
		x.Bytes(m.GetPrevious().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateUserPhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateContactRegistered) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateContactRegistered, layer)

	switch cmd {
	case Cmd_UpdateContactRegistered:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateContactRegistered))
		x.Int(m.GetUserId())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateContactRegistered encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateContactLink) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateContactLink, layer)

	switch cmd {
	case Cmd_UpdateContactLink:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateContactLink))
		x.Int(m.GetUserId())
		x.Bytes(m.GetMyLink().Encode(layer))
		x.Bytes(m.GetForeignLink().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateContactLink encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateNewEncryptedMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateNewEncryptedMessage, layer)

	switch cmd {
	case Cmd_UpdateNewEncryptedMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateNewEncryptedMessage))
		x.Bytes(m.GetMessage12BCBD9A71().Encode(layer))
		x.Int(m.GetQts())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateNewEncryptedMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateEncryptedChatTyping) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateEncryptedChatTyping, layer)

	switch cmd {
	case Cmd_UpdateEncryptedChatTyping:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateEncryptedChatTyping))
		x.Int(m.GetChatId())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateEncryptedChatTyping encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateEncryption) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateEncryption, layer)

	switch cmd {
	case Cmd_UpdateEncryption:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateEncryption))
		x.Bytes(m.GetChat().Encode(layer))
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateEncryption encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateEncryptedMessagesRead) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateEncryptedMessagesRead, layer)

	switch cmd {
	case Cmd_UpdateEncryptedMessagesRead:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateEncryptedMessagesRead))
		x.Int(m.GetChatId())
		x.Int(m.GetMaxDate())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateEncryptedMessagesRead encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChatParticipantAdd) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChatParticipantAdd, layer)

	switch cmd {
	case Cmd_UpdateChatParticipantAdd:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChatParticipantAdd))
		x.Int(m.GetChatId())
		x.Int(m.GetUserId())
		x.Int(m.GetInviterId())
		x.Int(m.GetDate())
		x.Int(m.GetVersion())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChatParticipantAdd encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChatParticipantDelete) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChatParticipantDelete, layer)

	switch cmd {
	case Cmd_UpdateChatParticipantDelete:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChatParticipantDelete))
		x.Int(m.GetChatId())
		x.Int(m.GetUserId())
		x.Int(m.GetVersion())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChatParticipantDelete encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateDcOptions) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateDcOptions, layer)

	switch cmd {
	case Cmd_UpdateDcOptions:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDcOptions))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetDcOptions())))
		for _, v := range m.GetDcOptions() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateDcOptions encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateUserBlocked) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateUserBlocked, layer)

	switch cmd {
	case Cmd_UpdateUserBlocked:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateUserBlocked))
		x.Int(m.GetUserId())
		x.Bytes(m.GetBlocked().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateUserBlocked encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateNotifySettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateNotifySettings, layer)

	switch cmd {
	case Cmd_UpdateNotifySettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateNotifySettings))
		x.Bytes(m.GetPeerBEC268EF71().Encode(layer))
		x.Bytes(m.GetNotifySettings().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateNotifySettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateServiceNotification) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateServiceNotification, layer)

	switch cmd {
	case Cmd_UpdateServiceNotification:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateServiceNotification))
		// flags
		var flags int32 = 0
		if m.GetPopupEBE4681971() == true {
			flags |= 1 << 0
		}
		if m.GetInboxDate() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		if m.GetInboxDate() != 0 {
			x.Int(m.GetInboxDate())
		}
		x.String(m.GetType())
		x.String(m.GetMessageEBE4681971())
		x.Bytes(m.GetMedia().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetEntities())))
		for _, v := range m.GetEntities() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_UpdateServiceNotification382dd3e4:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateServiceNotification382dd3e4))
		x.String(m.GetType())
		x.String(m.GetMessageEBE4681971())
		x.Bytes(m.GetMedia().Encode(layer))
		x.Bytes(m.GetPopup382DD3E451().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateServiceNotification encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatePrivacy) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatePrivacy, layer)

	switch cmd {
	case Cmd_UpdatePrivacy:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePrivacy))
		x.Bytes(m.GetKey().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetRules())))
		for _, v := range m.GetRules() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatePrivacy encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateUserPhone) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateUserPhone, layer)

	switch cmd {
	case Cmd_UpdateUserPhone:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateUserPhone))
		x.Int(m.GetUserId())
		x.String(m.GetPhone())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateUserPhone encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateReadHistoryInbox) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateReadHistoryInbox, layer)

	switch cmd {
	case Cmd_UpdateReadHistoryInbox:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateReadHistoryInbox))
		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))
		x.Int(m.GetMaxId())
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()
	case Cmd_UpdateReadHistoryInbox9c974fdf:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateReadHistoryInbox9c974fdf))
		// flags
		var flags int32 = 0
		if m.GetFolderId() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))
		x.Int(m.GetMaxId())
		x.Int(m.GetStillUnreadCount())
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateReadHistoryInbox encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateReadHistoryOutbox) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateReadHistoryOutbox, layer)

	switch cmd {
	case Cmd_UpdateReadHistoryOutbox:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateReadHistoryOutbox))
		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))
		x.Int(m.GetMaxId())
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateReadHistoryOutbox encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateWebPage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateWebPage, layer)

	switch cmd {
	case Cmd_UpdateWebPage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateWebPage))
		x.Bytes(m.GetWebpage().Encode(layer))
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateWebPage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateReadMessagesContents) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateReadMessagesContents, layer)

	switch cmd {
	case Cmd_UpdateReadMessagesContents:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateReadMessagesContents))
		x.VectorInt(m.GetMessages())

		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateReadMessagesContents encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChannelTooLong) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChannelTooLong, layer)

	switch cmd {
	case Cmd_UpdateChannelTooLong:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChannelTooLong))
		// flags
		var flags int32 = 0
		if m.GetPts() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetChannelId())
		if m.GetPts() != 0 {
			x.Int(m.GetPts())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChannelTooLong encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChannel) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChannel, layer)

	switch cmd {
	case Cmd_UpdateChannel:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChannel))
		x.Int(m.GetChannelId())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChannel encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateNewChannelMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateNewChannelMessage, layer)

	switch cmd {
	case Cmd_UpdateNewChannelMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateNewChannelMessage))
		x.Bytes(m.GetMessage1F2B0AFD71().Encode(layer))
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateNewChannelMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateReadChannelInbox) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateReadChannelInbox, layer)

	switch cmd {
	case Cmd_UpdateReadChannelInbox:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateReadChannelInbox))
		x.Int(m.GetChannelId())
		x.Int(m.GetMaxId())

		return x.GetBuf()
	case Cmd_UpdateReadChannelInbox330b5424:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateReadChannelInbox330b5424))
		// flags
		var flags int32 = 0
		if m.GetFolderId() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		x.Int(m.GetChannelId())
		x.Int(m.GetMaxId())
		x.Int(m.GetStillUnreadCount())
		x.Int(m.GetPts())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateReadChannelInbox encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateDeleteChannelMessages) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateDeleteChannelMessages, layer)

	switch cmd {
	case Cmd_UpdateDeleteChannelMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDeleteChannelMessages))
		x.Int(m.GetChannelId())
		x.VectorInt(m.GetMessages())

		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateDeleteChannelMessages encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChannelMessageViews) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChannelMessageViews, layer)

	switch cmd {
	case Cmd_UpdateChannelMessageViews:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChannelMessageViews))
		x.Int(m.GetChannelId())
		x.Int(m.GetId4E90BFD671())
		x.Int(m.GetViews())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChannelMessageViews encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChatAdmins) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChatAdmins, layer)

	switch cmd {
	case Cmd_UpdateChatAdmins:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChatAdmins))
		x.Int(m.GetChatId())
		x.Bytes(m.GetEnabled().Encode(layer))
		x.Int(m.GetVersion())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChatAdmins encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChatParticipantAdmin) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChatParticipantAdmin, layer)

	switch cmd {
	case Cmd_UpdateChatParticipantAdmin:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChatParticipantAdmin))
		x.Int(m.GetChatId())
		x.Int(m.GetUserId())
		x.Bytes(m.GetIsAdmin().Encode(layer))
		x.Int(m.GetVersion())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChatParticipantAdmin encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateNewStickerSet) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateNewStickerSet, layer)

	switch cmd {
	case Cmd_UpdateNewStickerSet:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateNewStickerSet))
		x.Bytes(m.GetStickerset().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateNewStickerSet encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateStickerSetsOrder) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateStickerSetsOrder, layer)

	switch cmd {
	case Cmd_UpdateStickerSetsOrder:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateStickerSetsOrder))
		// flags
		var flags int32 = 0
		if m.GetMasks() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.VectorLong(m.GetOrderBB2D20171())

		return x.GetBuf()
	case Cmd_UpdateStickerSetsOrderf0dfb451:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateStickerSetsOrderf0dfb451))
		x.VectorLong(m.GetOrderBB2D20171())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateStickerSetsOrder encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateStickerSets) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateStickerSets, layer)

	switch cmd {
	case Cmd_UpdateStickerSets:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateStickerSets))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateStickerSets encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateSavedGifs) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateSavedGifs, layer)

	switch cmd {
	case Cmd_UpdateSavedGifs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateSavedGifs))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateSavedGifs encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateBotInlineQuery) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateBotInlineQuery, layer)

	switch cmd {
	case Cmd_UpdateBotInlineQuery:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateBotInlineQuery))
		// flags
		var flags int32 = 0
		if m.GetGeo() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.GetQueryId())
		x.Int(m.GetUserId())
		x.String(m.GetQuery())
		if m.GetGeo() != nil {
			x.Bytes(m.GetGeo().Encode(layer))
		}
		x.String(m.GetOffset())

		return x.GetBuf()
	case Cmd_UpdateBotInlineQuery3f2038db:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateBotInlineQuery3f2038db))
		// flags
		var flags int32 = 0
		if m.GetGeo() != nil {
			flags |= 1 << 0
		}
		if m.GetPeerType() != nil {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetQueryId())
		x.Int(m.GetUserId())
		x.String(m.GetQuery())
		if m.GetGeo() != nil {
			x.Bytes(m.GetGeo().Encode(layer))
		}
		if m.GetPeerType() != nil {
			x.Bytes(m.GetPeerType().Encode(layer))
		}
		x.String(m.GetOffset())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateBotInlineQuery encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateBotInlineSend) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateBotInlineSend, layer)

	switch cmd {
	case Cmd_UpdateBotInlineSend:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateBotInlineSend))
		// flags
		var flags int32 = 0
		if m.GetGeo() != nil {
			flags |= 1 << 0
		}
		if m.GetMsgIdE48F96471() != nil {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetUserId())
		x.String(m.GetQuery())
		if m.GetGeo() != nil {
			x.Bytes(m.GetGeo().Encode(layer))
		}
		x.String(m.GetIdE48F96471())
		if m.GetMsgIdE48F96471() != nil {
			x.Bytes(m.GetMsgIdE48F96471().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateBotInlineSend encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateEditChannelMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateEditChannelMessage, layer)

	switch cmd {
	case Cmd_UpdateEditChannelMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateEditChannelMessage))
		x.Bytes(m.GetMessage1F2B0AFD71().Encode(layer))
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateEditChannelMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChannelPinnedMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChannelPinnedMessage, layer)

	switch cmd {
	case Cmd_UpdateChannelPinnedMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChannelPinnedMessage))
		x.Int(m.GetChannelId())
		x.Int(m.GetId4E90BFD671())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChannelPinnedMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateBotCallbackQuery) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateBotCallbackQuery, layer)

	switch cmd {
	case Cmd_UpdateBotCallbackQuery:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateBotCallbackQuery))
		// flags
		var flags int32 = 0
		if len(m.GetDataE73547E171()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetGameShortName()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetQueryId())
		x.Int(m.GetUserId())
		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))
		x.Int(m.GetMsgIdE73547E171())
		x.Long(m.GetChatInstance())
		if len(m.GetDataE73547E171()) > 0 {
			x.StringBytes(m.GetDataE73547E171())
		}
		if len(m.GetGameShortName()) > 0 {
			x.String(m.GetGameShortName())
		}

		return x.GetBuf()
	case Cmd_UpdateBotCallbackQuerya68c688c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateBotCallbackQuerya68c688c))
		x.Long(m.GetQueryId())
		x.Int(m.GetUserId())
		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))
		x.Int(m.GetMsgIdE73547E171())
		x.StringBytes(m.GetDataE73547E171())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateBotCallbackQuery encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateEditMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateEditMessage, layer)

	switch cmd {
	case Cmd_UpdateEditMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateEditMessage))
		x.Bytes(m.GetMessage1F2B0AFD71().Encode(layer))
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateEditMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateInlineBotCallbackQuery) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateInlineBotCallbackQuery, layer)

	switch cmd {
	case Cmd_UpdateInlineBotCallbackQuery:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateInlineBotCallbackQuery))
		// flags
		var flags int32 = 0
		if len(m.GetDataE73547E171()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetGameShortName()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetQueryId())
		x.Int(m.GetUserId())
		x.Bytes(m.GetMsgIdE48F96471().Encode(layer))
		x.Long(m.GetChatInstance())
		if len(m.GetDataE73547E171()) > 0 {
			x.StringBytes(m.GetDataE73547E171())
		}
		if len(m.GetGameShortName()) > 0 {
			x.String(m.GetGameShortName())
		}

		return x.GetBuf()
	case Cmd_UpdateInlineBotCallbackQuery2cbd95af:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateInlineBotCallbackQuery2cbd95af))
		x.Long(m.GetQueryId())
		x.Int(m.GetUserId())
		x.Bytes(m.GetMsgIdE48F96471().Encode(layer))
		x.StringBytes(m.GetDataE73547E171())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateInlineBotCallbackQuery encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateReadChannelOutbox) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateReadChannelOutbox, layer)

	switch cmd {
	case Cmd_UpdateReadChannelOutbox:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateReadChannelOutbox))
		x.Int(m.GetChannelId())
		x.Int(m.GetMaxId())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateReadChannelOutbox encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateDraftMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateDraftMessage, layer)

	switch cmd {
	case Cmd_UpdateDraftMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDraftMessage))
		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))
		x.Bytes(m.GetDraft().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateDraftMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateReadFeaturedStickers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateReadFeaturedStickers, layer)

	switch cmd {
	case Cmd_UpdateReadFeaturedStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateReadFeaturedStickers))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateReadFeaturedStickers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateRecentStickers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateRecentStickers, layer)

	switch cmd {
	case Cmd_UpdateRecentStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateRecentStickers))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateRecentStickers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateConfig) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateConfig, layer)

	switch cmd {
	case Cmd_UpdateConfig:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateConfig))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateConfig encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatePtsChanged) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatePtsChanged, layer)

	switch cmd {
	case Cmd_UpdatePtsChanged:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePtsChanged))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatePtsChanged encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChannelWebPage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChannelWebPage, layer)

	switch cmd {
	case Cmd_UpdateChannelWebPage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChannelWebPage))
		x.Int(m.GetChannelId())
		x.Bytes(m.GetWebpage().Encode(layer))
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChannelWebPage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateDialogPinned) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateDialogPinned, layer)

	switch cmd {
	case Cmd_UpdateDialogPinned:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDialogPinned))
		// flags
		var flags int32 = 0
		if m.GetPinned() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))

		return x.GetBuf()
	case Cmd_UpdateDialogPinned19d27f3c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDialogPinned19d27f3c))
		// flags
		var flags int32 = 0
		if m.GetPinned() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetPeer19D27F3C85().Encode(layer))

		return x.GetBuf()
	case Cmd_UpdateDialogPinned6e6fe51c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDialogPinned6e6fe51c))
		// flags
		var flags int32 = 0
		if m.GetPinned() == true {
			flags |= 1 << 0
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		x.Bytes(m.GetPeer19D27F3C85().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateDialogPinned encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatePinnedDialogs) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatePinnedDialogs, layer)

	switch cmd {
	case Cmd_UpdatePinnedDialogs:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePinnedDialogs))
		// flags
		var flags int32 = 0
		if len(m.GetOrderD8CAF68D71()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if len(m.GetOrderD8CAF68D71()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetOrderD8CAF68D71())))
			for _, v := range m.GetOrderD8CAF68D71() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_UpdatePinnedDialogsea4cb65b:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePinnedDialogsea4cb65b))
		// flags
		var flags int32 = 0
		if len(m.GetOrderEA4CB65B85()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if len(m.GetOrderEA4CB65B85()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetOrderEA4CB65B85())))
			for _, v := range m.GetOrderEA4CB65B85() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_UpdatePinnedDialogsfa0f3ca2:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePinnedDialogsfa0f3ca2))
		// flags
		var flags int32 = 0
		if m.GetFolderId() != 0 {
			flags |= 1 << 1
		}
		if len(m.GetOrderEA4CB65B85()) > 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}
		if len(m.GetOrderEA4CB65B85()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetOrderEA4CB65B85())))
			for _, v := range m.GetOrderEA4CB65B85() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatePinnedDialogs encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateBotWebhookJSON) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateBotWebhookJSON, layer)

	switch cmd {
	case Cmd_UpdateBotWebhookJSON:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateBotWebhookJSON))
		x.Bytes(m.GetData8317C0C371().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateBotWebhookJSON encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateBotWebhookJSONQuery) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateBotWebhookJSONQuery, layer)

	switch cmd {
	case Cmd_UpdateBotWebhookJSONQuery:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateBotWebhookJSONQuery))
		x.Long(m.GetQueryId())
		x.Bytes(m.GetData8317C0C371().Encode(layer))
		x.Int(m.GetTimeout())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateBotWebhookJSONQuery encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateBotShippingQuery) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateBotShippingQuery, layer)

	switch cmd {
	case Cmd_UpdateBotShippingQuery:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateBotShippingQuery))
		x.Long(m.GetQueryId())
		x.Int(m.GetUserId())
		x.StringBytes(m.GetPayload())
		x.Bytes(m.GetShippingAddress().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateBotShippingQuery encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateBotPrecheckoutQuery) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateBotPrecheckoutQuery, layer)

	switch cmd {
	case Cmd_UpdateBotPrecheckoutQuery:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateBotPrecheckoutQuery))
		// flags
		var flags int32 = 0
		if m.GetInfo() != nil {
			flags |= 1 << 0
		}
		if len(m.GetShippingOptionId()) > 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetQueryId())
		x.Int(m.GetUserId())
		x.StringBytes(m.GetPayload())
		if m.GetInfo() != nil {
			x.Bytes(m.GetInfo().Encode(layer))
		}
		if len(m.GetShippingOptionId()) > 0 {
			x.String(m.GetShippingOptionId())
		}
		x.String(m.GetCurrency())
		x.Long(m.GetTotalAmount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateBotPrecheckoutQuery encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatePhoneCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatePhoneCall, layer)

	switch cmd {
	case Cmd_UpdatePhoneCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePhoneCall))
		x.Bytes(m.GetPhoneCall().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatePhoneCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateLangPackTooLong) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateLangPackTooLong, layer)

	switch cmd {
	case Cmd_UpdateLangPackTooLong:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateLangPackTooLong))
		_ = m

		return x.GetBuf()
	case Cmd_UpdateLangPackTooLong46560264:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateLangPackTooLong46560264))
		x.String(m.GetLangCode())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateLangPackTooLong encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateLangPack) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateLangPack, layer)

	switch cmd {
	case Cmd_UpdateLangPack:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateLangPack))
		x.Bytes(m.GetDifference().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateLangPack encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateFavedStickers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateFavedStickers, layer)

	switch cmd {
	case Cmd_UpdateFavedStickers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateFavedStickers))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateFavedStickers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChannelReadMessagesContents) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChannelReadMessagesContents, layer)

	switch cmd {
	case Cmd_UpdateChannelReadMessagesContents:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChannelReadMessagesContents))
		x.Int(m.GetChannelId())
		x.VectorInt(m.GetMessages())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChannelReadMessagesContents encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateContactsReset) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateContactsReset, layer)

	switch cmd {
	case Cmd_UpdateContactsReset:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateContactsReset))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateContactsReset encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateNewAuthorization) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateNewAuthorization, layer)

	switch cmd {
	case Cmd_UpdateNewAuthorization:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateNewAuthorization))
		x.Long(m.GetAuthKeyId())
		x.Int(m.GetDate())
		x.String(m.GetDevice())
		x.String(m.GetLocation())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateNewAuthorization encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChannelGroup) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChannelGroup, layer)

	switch cmd {
	case Cmd_UpdateChannelGroup:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChannelGroup))
		x.Int(m.GetChannelId())
		x.Bytes(m.GetGroup().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChannelGroup encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChannelAvailableMessages) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChannelAvailableMessages, layer)

	switch cmd {
	case Cmd_UpdateChannelAvailableMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChannelAvailableMessages))
		x.Int(m.GetChannelId())
		x.Int(m.GetAvailableMinId())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChannelAvailableMessages encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateDialogUnreadMark) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateDialogUnreadMark, layer)

	switch cmd {
	case Cmd_UpdateDialogUnreadMark:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDialogUnreadMark))
		// flags
		var flags int32 = 0
		if m.GetUnread() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetPeer19D27F3C85().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateDialogUnreadMark encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateUserPinnedMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateUserPinnedMessage, layer)

	switch cmd {
	case Cmd_UpdateUserPinnedMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateUserPinnedMessage))
		x.Int(m.GetUserId())
		x.Int(m.GetId4E90BFD671())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateUserPinnedMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChatPinnedMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChatPinnedMessage, layer)

	switch cmd {
	case Cmd_UpdateChatPinnedMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChatPinnedMessage))
		x.Int(m.GetChatId())
		x.Int(m.GetId4E90BFD671())

		return x.GetBuf()
	case Cmd_UpdateChatPinnedMessagee10db349:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChatPinnedMessagee10db349))
		x.Int(m.GetChatId())
		x.Int(m.GetId4E90BFD671())
		x.Int(m.GetVersion())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChatPinnedMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateMessagePoll) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateMessagePoll, layer)

	switch cmd {
	case Cmd_UpdateMessagePoll:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateMessagePoll))
		// flags
		var flags int32 = 0
		if m.GetPoll() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.GetPollId())
		if m.GetPoll() != nil {
			x.Bytes(m.GetPoll().Encode(layer))
		}
		x.Bytes(m.GetResults().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateMessagePoll encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChatDefaultBannedRights) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChatDefaultBannedRights, layer)

	switch cmd {
	case Cmd_UpdateChatDefaultBannedRights:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChatDefaultBannedRights))
		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))
		x.Bytes(m.GetDefaultBannedRights().Encode(layer))
		x.Int(m.GetVersion())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChatDefaultBannedRights encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateFolderPeers) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateFolderPeers, layer)

	switch cmd {
	case Cmd_UpdateFolderPeers:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateFolderPeers))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetFolderPeers())))
		for _, v := range m.GetFolderPeers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateFolderPeers encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatePeerSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatePeerSettings, layer)

	switch cmd {
	case Cmd_UpdatePeerSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePeerSettings))
		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))
		x.Bytes(m.GetSettings().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatePeerSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatePeerLocated) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatePeerLocated, layer)

	switch cmd {
	case Cmd_UpdatePeerLocated:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePeerLocated))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetPeers())))
		for _, v := range m.GetPeers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatePeerLocated encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateNewScheduledMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateNewScheduledMessage, layer)

	switch cmd {
	case Cmd_UpdateNewScheduledMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateNewScheduledMessage))
		x.Bytes(m.GetMessage1F2B0AFD71().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateNewScheduledMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateDeleteScheduledMessages) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateDeleteScheduledMessages, layer)

	switch cmd {
	case Cmd_UpdateDeleteScheduledMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDeleteScheduledMessages))
		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))
		x.VectorInt(m.GetMessages())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateDeleteScheduledMessages encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateTheme) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateTheme, layer)

	switch cmd {
	case Cmd_UpdateTheme:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateTheme))
		x.Bytes(m.GetTheme().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateTheme encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateGeoLiveViewed) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateGeoLiveViewed, layer)

	switch cmd {
	case Cmd_UpdateGeoLiveViewed:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateGeoLiveViewed))
		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))
		x.Int(m.GetMsgIdE73547E171())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateGeoLiveViewed encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateLoginToken) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateLoginToken, layer)

	switch cmd {
	case Cmd_UpdateLoginToken:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateLoginToken))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateLoginToken encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateMessagePollVote) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateMessagePollVote, layer)

	switch cmd {
	case Cmd_UpdateMessagePollVote:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateMessagePollVote))
		x.Long(m.GetPollId())
		x.Int(m.GetUserId())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateMessagePollVote encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateDialogFilter) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateDialogFilter, layer)

	switch cmd {
	case Cmd_UpdateDialogFilter:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDialogFilter))
		// flags
		var flags int32 = 0
		if m.GetFilter() != nil {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetId4E90BFD671())
		if m.GetFilter() != nil {
			x.Bytes(m.GetFilter().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateDialogFilter encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateDialogFilterOrder) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateDialogFilterOrder, layer)

	switch cmd {
	case Cmd_UpdateDialogFilterOrder:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDialogFilterOrder))
		x.VectorInt(m.GetOrderA5D72105111())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateDialogFilterOrder encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateDialogFilters) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateDialogFilters, layer)

	switch cmd {
	case Cmd_UpdateDialogFilters:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateDialogFilters))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateDialogFilters encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatePhoneCallSignalingData) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatePhoneCallSignalingData, layer)

	switch cmd {
	case Cmd_UpdatePhoneCallSignalingData:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePhoneCallSignalingData))
		x.Long(m.GetPhoneCallId())
		x.StringBytes(m.GetDataE73547E171())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatePhoneCallSignalingData encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChannelParticipant) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChannelParticipant, layer)

	switch cmd {
	case Cmd_UpdateChannelParticipant:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChannelParticipant))
		// flags
		var flags int32 = 0
		if m.GetPrevParticipant() != nil {
			flags |= 1 << 0
		}
		if m.GetNewParticipant() != nil {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetChannelId())
		x.Int(m.GetDate())
		x.Int(m.GetUserId())
		if m.GetPrevParticipant() != nil {
			x.Bytes(m.GetPrevParticipant().Encode(layer))
		}
		if m.GetNewParticipant() != nil {
			x.Bytes(m.GetNewParticipant().Encode(layer))
		}
		x.Int(m.GetQts())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChannelParticipant encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChannelMessageForwards) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChannelMessageForwards, layer)

	switch cmd {
	case Cmd_UpdateChannelMessageForwards:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChannelMessageForwards))
		x.Int(m.GetChannelId())
		x.Int(m.GetId4E90BFD671())
		x.Int(m.GetForwards())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChannelMessageForwards encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateReadChannelDiscussionInbox) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateReadChannelDiscussionInbox, layer)

	switch cmd {
	case Cmd_UpdateReadChannelDiscussionInbox:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateReadChannelDiscussionInbox))
		// flags
		var flags int32 = 0
		if m.GetBroadcastId() != 0 {
			flags |= 1 << 0
		}
		if m.GetBroadcastPost() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetChannelId())
		x.Int(m.GetTopMsgId())
		x.Int(m.GetReadMaxId())
		if m.GetBroadcastId() != 0 {
			x.Int(m.GetBroadcastId())
		}
		if m.GetBroadcastPost() != 0 {
			x.Int(m.GetBroadcastPost())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateReadChannelDiscussionInbox encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateReadChannelDiscussionOutbox) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateReadChannelDiscussionOutbox, layer)

	switch cmd {
	case Cmd_UpdateReadChannelDiscussionOutbox:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateReadChannelDiscussionOutbox))
		x.Int(m.GetChannelId())
		x.Int(m.GetTopMsgId())
		x.Int(m.GetReadMaxId())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateReadChannelDiscussionOutbox encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatePeerBlocked) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatePeerBlocked, layer)

	switch cmd {
	case Cmd_UpdatePeerBlocked:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePeerBlocked))
		x.Bytes(m.GetPeerId().Encode(layer))
		x.Bytes(m.GetBlocked().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatePeerBlocked encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChannelUserTyping) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChannelUserTyping, layer)

	switch cmd {
	case Cmd_UpdateChannelUserTyping:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChannelUserTyping))
		// flags
		var flags int32 = 0
		if m.GetTopMsgId() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetChannelId())
		if m.GetTopMsgId() != 0 {
			x.Int(m.GetTopMsgId())
		}
		x.Int(m.GetUserId())
		x.Bytes(m.GetAction().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChannelUserTyping encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatePinnedMessages) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatePinnedMessages, layer)

	switch cmd {
	case Cmd_UpdatePinnedMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePinnedMessages))
		// flags
		var flags int32 = 0
		if m.GetPinned() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetPeer9961FD5C71().Encode(layer))
		x.VectorInt(m.GetMessages())

		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatePinnedMessages encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatePinnedChannelMessages) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatePinnedChannelMessages, layer)

	switch cmd {
	case Cmd_UpdatePinnedChannelMessages:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatePinnedChannelMessages))
		// flags
		var flags int32 = 0
		if m.GetPinned() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Int(m.GetChannelId())
		x.VectorInt(m.GetMessages())

		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatePinnedChannelMessages encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateChat) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateChat, layer)

	switch cmd {
	case Cmd_UpdateChat:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateChat))
		x.Int(m.GetChatId())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateChat encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateGroupCallParticipants) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateGroupCallParticipants, layer)

	switch cmd {
	case Cmd_UpdateGroupCallParticipants:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateGroupCallParticipants))
		x.Bytes(m.GetCallF2EBDB4E122().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetParticipantsF2EBDB4E122())))
		for _, v := range m.GetParticipantsF2EBDB4E122() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetVersion())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateGroupCallParticipants encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateGroupCall) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateGroupCall, layer)

	switch cmd {
	case Cmd_UpdateGroupCall:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateGroupCall))
		x.Int(m.GetChatId())
		x.Bytes(m.GetCallA45EB99B122().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateGroupCall encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesTooLong) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatesTooLong, layer)

	switch cmd {
	case Cmd_UpdatesTooLong:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesTooLong))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatesTooLong encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateShortMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateShortMessage, layer)

	switch cmd {
	case Cmd_UpdateShortMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateShortMessage))
		// flags
		var flags int32 = 0
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMentioned() == true {
			flags |= 1 << 4
		}
		if m.GetMediaUnread() == true {
			flags |= 1 << 5
		}
		if m.GetSilent() == true {
			flags |= 1 << 13
		}
		if m.GetFwdFrom() != nil {
			flags |= 1 << 2
		}
		if m.GetViaBotId() != 0 {
			flags |= 1 << 11
		}
		if m.GetReplyToMsgId() != 0 {
			flags |= 1 << 3
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 7
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.Int(m.GetUserId())
		x.String(m.GetMessage())
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())
		x.Int(m.GetDate())
		if m.GetFwdFrom() != nil {
			x.Bytes(m.GetFwdFrom().Encode(layer))
		}
		if m.GetViaBotId() != 0 {
			x.Int(m.GetViaBotId())
		}
		if m.GetReplyToMsgId() != 0 {
			x.Int(m.GetReplyToMsgId())
		}
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_UpdateShortMessage2296d2c8:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateShortMessage2296d2c8))
		// flags
		var flags int32 = 0
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMentioned() == true {
			flags |= 1 << 4
		}
		if m.GetMediaUnread() == true {
			flags |= 1 << 5
		}
		if m.GetSilent() == true {
			flags |= 1 << 13
		}
		if m.GetFwdFrom() != nil {
			flags |= 1 << 2
		}
		if m.GetViaBotId() != 0 {
			flags |= 1 << 11
		}
		if m.GetReplyTo() != nil {
			flags |= 1 << 3
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 7
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.Int(m.GetUserId())
		x.String(m.GetMessage())
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())
		x.Int(m.GetDate())
		if m.GetFwdFrom() != nil {
			x.Bytes(m.GetFwdFrom().Encode(layer))
		}
		if m.GetViaBotId() != 0 {
			x.Int(m.GetViaBotId())
		}
		if m.GetReplyTo() != nil {
			x.Bytes(m.GetReplyTo().Encode(layer))
		}
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateShortMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateShortChatMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateShortChatMessage, layer)

	switch cmd {
	case Cmd_UpdateShortChatMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateShortChatMessage))
		// flags
		var flags int32 = 0
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMentioned() == true {
			flags |= 1 << 4
		}
		if m.GetMediaUnread() == true {
			flags |= 1 << 5
		}
		if m.GetSilent() == true {
			flags |= 1 << 13
		}
		if m.GetFwdFrom() != nil {
			flags |= 1 << 2
		}
		if m.GetViaBotId() != 0 {
			flags |= 1 << 11
		}
		if m.GetReplyToMsgId() != 0 {
			flags |= 1 << 3
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 7
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.Int(m.GetFromId())
		x.Int(m.GetChatId())
		x.String(m.GetMessage())
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())
		x.Int(m.GetDate())
		if m.GetFwdFrom() != nil {
			x.Bytes(m.GetFwdFrom().Encode(layer))
		}
		if m.GetViaBotId() != 0 {
			x.Int(m.GetViaBotId())
		}
		if m.GetReplyToMsgId() != 0 {
			x.Int(m.GetReplyToMsgId())
		}
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_UpdateShortChatMessage402d5dbb:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateShortChatMessage402d5dbb))
		// flags
		var flags int32 = 0
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMentioned() == true {
			flags |= 1 << 4
		}
		if m.GetMediaUnread() == true {
			flags |= 1 << 5
		}
		if m.GetSilent() == true {
			flags |= 1 << 13
		}
		if m.GetFwdFrom() != nil {
			flags |= 1 << 2
		}
		if m.GetViaBotId() != 0 {
			flags |= 1 << 11
		}
		if m.GetReplyTo() != nil {
			flags |= 1 << 3
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 7
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.Int(m.GetFromId())
		x.Int(m.GetChatId())
		x.String(m.GetMessage())
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())
		x.Int(m.GetDate())
		if m.GetFwdFrom() != nil {
			x.Bytes(m.GetFwdFrom().Encode(layer))
		}
		if m.GetViaBotId() != 0 {
			x.Int(m.GetViaBotId())
		}
		if m.GetReplyTo() != nil {
			x.Bytes(m.GetReplyTo().Encode(layer))
		}
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateShortChatMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateShort) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateShort, layer)

	switch cmd {
	case Cmd_UpdateShort:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateShort))
		x.Bytes(m.GetUpdate().Encode(layer))
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateShort encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesCombined) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatesCombined, layer)

	switch cmd {
	case Cmd_UpdatesCombined:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesCombined))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUpdates())))
		for _, v := range m.GetUpdates() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetDate())
		x.Int(m.GetSeqStart())
		x.Int(m.GetSeq())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatesCombined encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdates) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdates, layer)

	switch cmd {
	case Cmd_Updates:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Updates))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUpdates())))
		for _, v := range m.GetUpdates() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetDate())
		x.Int(m.GetSeq())

		return x.GetBuf()

	default:
		log.Errorf("invalid Updates encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdateShortSentMessage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdateShortSentMessage, layer)

	switch cmd {
	case Cmd_UpdateShortSentMessage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdateShortSentMessage))
		// flags
		var flags int32 = 0
		if m.GetOut() == true {
			flags |= 1 << 1
		}
		if m.GetMedia() != nil {
			flags |= 1 << 9
		}
		if len(m.GetEntities()) > 0 {
			flags |= 1 << 7
		}
		x.Int(flags)

		x.Int(m.GetId())
		x.Int(m.GetPts())
		x.Int(m.GetPtsCount())
		x.Int(m.GetDate())
		if m.GetMedia() != nil {
			x.Bytes(m.GetMedia().Encode(layer))
		}
		if len(m.GetEntities()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetEntities())))
			for _, v := range m.GetEntities() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdateShortSentMessage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesChannelDifferenceEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatesChannelDifferenceEmpty, layer)

	switch cmd {
	case Cmd_UpdatesChannelDifferenceEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesChannelDifferenceEmpty))
		// flags
		var flags int32 = 0
		if m.GetFinal() == true {
			flags |= 1 << 0
		}
		if m.GetTimeout() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetPts())
		if m.GetTimeout() != 0 {
			x.Int(m.GetTimeout())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatesChannelDifferenceEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesChannelDifferenceTooLong) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatesChannelDifferenceTooLong, layer)

	switch cmd {
	case Cmd_UpdatesChannelDifferenceTooLong:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesChannelDifferenceTooLong))
		// flags
		var flags int32 = 0
		if m.GetFinal() == true {
			flags |= 1 << 0
		}
		if m.GetTimeout() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetPts())
		if m.GetTimeout() != 0 {
			x.Int(m.GetTimeout())
		}
		x.Int(m.GetTopMessage())
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetReadOutboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Int(m.GetUnreadMentionsCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_UpdatesChannelDifferenceTooLong5e167646:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesChannelDifferenceTooLong5e167646))
		// flags
		var flags int32 = 0
		if m.GetFinal() == true {
			flags |= 1 << 0
		}
		if m.GetTimeout() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetPts())
		if m.GetTimeout() != 0 {
			x.Int(m.GetTimeout())
		}
		x.Int(m.GetTopMessage())
		x.Int(m.GetTopImportantMessage())
		x.Int(m.GetReadInboxMaxId())
		x.Int(m.GetUnreadCount())
		x.Int(m.GetUnreadImportantCount())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_UpdatesChannelDifferenceTooLonga4bcc6fe:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesChannelDifferenceTooLonga4bcc6fe))
		// flags
		var flags int32 = 0
		if m.GetFinal() == true {
			flags |= 1 << 0
		}
		if m.GetTimeout() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		if m.GetTimeout() != 0 {
			x.Int(m.GetTimeout())
		}
		x.Bytes(m.GetDialog().Encode(layer))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetMessages())))
		for _, v := range m.GetMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatesChannelDifferenceTooLong encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesChannelDifference) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatesChannelDifference, layer)

	switch cmd {
	case Cmd_UpdatesChannelDifference:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesChannelDifference))
		// flags
		var flags int32 = 0
		if m.GetFinal() == true {
			flags |= 1 << 0
		}
		if m.GetTimeout() != 0 {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Int(m.GetPts())
		if m.GetTimeout() != 0 {
			x.Int(m.GetTimeout())
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetNewMessages())))
		for _, v := range m.GetNewMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetOtherUpdates())))
		for _, v := range m.GetOtherUpdates() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatesChannelDifference encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesDifferenceEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatesDifferenceEmpty, layer)

	switch cmd {
	case Cmd_UpdatesDifferenceEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesDifferenceEmpty))
		x.Int(m.GetDate())
		x.Int(m.GetSeq())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatesDifferenceEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesDifference) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatesDifference, layer)

	switch cmd {
	case Cmd_UpdatesDifference:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesDifference))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetNewMessages())))
		for _, v := range m.GetNewMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetNewEncryptedMessages())))
		for _, v := range m.GetNewEncryptedMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetOtherUpdates())))
		for _, v := range m.GetOtherUpdates() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.GetState().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatesDifference encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesDifferenceSlice) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatesDifferenceSlice, layer)

	switch cmd {
	case Cmd_UpdatesDifferenceSlice:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesDifferenceSlice))
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetNewMessages())))
		for _, v := range m.GetNewMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetNewEncryptedMessages())))
		for _, v := range m.GetNewEncryptedMessages() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetOtherUpdates())))
		for _, v := range m.GetOtherUpdates() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetChats())))
		for _, v := range m.GetChats() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetUsers())))
		for _, v := range m.GetUsers() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Bytes(m.GetIntermediateState().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatesDifferenceSlice encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesDifferenceTooLong) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatesDifferenceTooLong, layer)

	switch cmd {
	case Cmd_UpdatesDifferenceTooLong:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesDifferenceTooLong))
		x.Int(m.GetPts())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatesDifferenceTooLong encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUpdatesState) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUpdatesState, layer)

	switch cmd {
	case Cmd_UpdatesState:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UpdatesState))
		x.Int(m.GetPts())
		x.Int(m.GetQts())
		x.Int(m.GetDate())
		x.Int(m.GetSeq())
		x.Int(m.GetUnreadCount())

		return x.GetBuf()

	default:
		log.Errorf("invalid UpdatesState encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadCdnFileReuploadNeeded) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUploadCdnFileReuploadNeeded, layer)

	switch cmd {
	case Cmd_UploadCdnFileReuploadNeeded:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadCdnFileReuploadNeeded))
		x.StringBytes(m.GetRequestToken())

		return x.GetBuf()

	default:
		log.Errorf("invalid UploadCdnFileReuploadNeeded encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadCdnFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUploadCdnFile, layer)

	switch cmd {
	case Cmd_UploadCdnFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadCdnFile))
		x.StringBytes(m.GetBytes())

		return x.GetBuf()

	default:
		log.Errorf("invalid UploadCdnFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUploadFile, layer)

	switch cmd {
	case Cmd_UploadFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadFile))
		x.Bytes(m.GetType().Encode(layer))
		x.Int(m.GetMtime())
		x.StringBytes(m.GetBytes())

		return x.GetBuf()

	default:
		log.Errorf("invalid UploadFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadFileCdnRedirect) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUploadFileCdnRedirect, layer)

	switch cmd {
	case Cmd_UploadFileCdnRedirect:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadFileCdnRedirect))
		x.Int(m.GetDcId())
		x.StringBytes(m.GetFileToken())
		x.StringBytes(m.GetEncryptionKey())
		x.StringBytes(m.GetEncryptionIv())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetCdnFileHashes())))
		for _, v := range m.GetCdnFileHashes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()
	case Cmd_UploadFileCdnRedirectf18cda44:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadFileCdnRedirectf18cda44))
		x.Int(m.GetDcId())
		x.StringBytes(m.GetFileToken())
		x.StringBytes(m.GetEncryptionKey())
		x.StringBytes(m.GetEncryptionIv())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetFileHashes())))
		for _, v := range m.GetFileHashes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UploadFileCdnRedirect encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUploadWebFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUploadWebFile, layer)

	switch cmd {
	case Cmd_UploadWebFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UploadWebFile))
		x.Int(m.GetSize_())
		x.String(m.GetMimeType())
		x.Bytes(m.GetFileType().Encode(layer))
		x.Int(m.GetMtime())
		x.StringBytes(m.GetBytes())

		return x.GetBuf()

	default:
		log.Errorf("invalid UploadWebFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUrlAuthResultRequest) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUrlAuthResultRequest, layer)

	switch cmd {
	case Cmd_UrlAuthResultRequest:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UrlAuthResultRequest))
		// flags
		var flags int32 = 0
		if m.GetRequestWriteAccess() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Bytes(m.GetBot().Encode(layer))
		x.String(m.GetDomain())

		return x.GetBuf()

	default:
		log.Errorf("invalid UrlAuthResultRequest encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUrlAuthResultAccepted) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUrlAuthResultAccepted, layer)

	switch cmd {
	case Cmd_UrlAuthResultAccepted:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UrlAuthResultAccepted))
		x.String(m.GetUrl())

		return x.GetBuf()

	default:
		log.Errorf("invalid UrlAuthResultAccepted encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUrlAuthResultDefault) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUrlAuthResultDefault, layer)

	switch cmd {
	case Cmd_UrlAuthResultDefault:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UrlAuthResultDefault))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UrlAuthResultDefault encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUserEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUserEmpty, layer)

	switch cmd {
	case Cmd_UserEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserEmpty))
		x.Int(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid UserEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUser) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUser, layer)

	switch cmd {
	case Cmd_User:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_User))
		// flags
		var flags int32 = 0
		if m.GetSelf() == true {
			flags |= 1 << 10
		}
		if m.GetContact() == true {
			flags |= 1 << 11
		}
		if m.GetMutualContact() == true {
			flags |= 1 << 12
		}
		if m.GetDeleted() == true {
			flags |= 1 << 13
		}
		if m.GetBot() == true {
			flags |= 1 << 14
		}
		if m.GetBotChatHistory() == true {
			flags |= 1 << 15
		}
		if m.GetBotNochats() == true {
			flags |= 1 << 16
		}
		if m.GetVerified() == true {
			flags |= 1 << 17
		}
		if m.GetRestricted() == true {
			flags |= 1 << 18
		}
		if m.GetMin() == true {
			flags |= 1 << 20
		}
		if m.GetBotInlineGeo() == true {
			flags |= 1 << 21
		}
		if m.GetAccessHash() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetFirstName()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetLastName()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetUsername()) > 0 {
			flags |= 1 << 3
		}
		if len(m.GetPhone()) > 0 {
			flags |= 1 << 4
		}
		if m.GetPhoto() != nil {
			flags |= 1 << 5
		}
		if m.GetStatus() != nil {
			flags |= 1 << 6
		}
		if m.GetBotInfoVersion() != 0 {
			flags |= 1 << 14
		}
		if len(m.GetRestrictionReason2E13F4C371()) > 0 {
			flags |= 1 << 18
		}
		if len(m.GetBotInlinePlaceholder()) > 0 {
			flags |= 1 << 19
		}
		if len(m.GetLangCode()) > 0 {
			flags |= 1 << 22
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetAccessHash() != 0 {
			x.Long(m.GetAccessHash())
		}
		if len(m.GetFirstName()) > 0 {
			x.String(m.GetFirstName())
		}
		if len(m.GetLastName()) > 0 {
			x.String(m.GetLastName())
		}
		if len(m.GetUsername()) > 0 {
			x.String(m.GetUsername())
		}
		if len(m.GetPhone()) > 0 {
			x.String(m.GetPhone())
		}
		if m.GetPhoto() != nil {
			x.Bytes(m.GetPhoto().Encode(layer))
		}
		if m.GetStatus() != nil {
			x.Bytes(m.GetStatus().Encode(layer))
		}
		if m.GetBotInfoVersion() != 0 {
			x.Int(m.GetBotInfoVersion())
		}
		if len(m.GetRestrictionReason2E13F4C371()) > 0 {
			x.String(m.GetRestrictionReason2E13F4C371())
		}
		if len(m.GetBotInlinePlaceholder()) > 0 {
			x.String(m.GetBotInlinePlaceholder())
		}
		if len(m.GetLangCode()) > 0 {
			x.String(m.GetLangCode())
		}

		return x.GetBuf()
	case Cmd_User938458c1:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_User938458c1))
		// flags
		var flags int32 = 0
		if m.GetSelf() == true {
			flags |= 1 << 10
		}
		if m.GetContact() == true {
			flags |= 1 << 11
		}
		if m.GetMutualContact() == true {
			flags |= 1 << 12
		}
		if m.GetDeleted() == true {
			flags |= 1 << 13
		}
		if m.GetBot() == true {
			flags |= 1 << 14
		}
		if m.GetBotChatHistory() == true {
			flags |= 1 << 15
		}
		if m.GetBotNochats() == true {
			flags |= 1 << 16
		}
		if m.GetVerified() == true {
			flags |= 1 << 17
		}
		if m.GetRestricted() == true {
			flags |= 1 << 18
		}
		if m.GetMin() == true {
			flags |= 1 << 20
		}
		if m.GetBotInlineGeo() == true {
			flags |= 1 << 21
		}
		if m.GetSupport() == true {
			flags |= 1 << 23
		}
		if m.GetScam() == true {
			flags |= 1 << 24
		}
		if m.GetApplyMinPhoto() == true {
			flags |= 1 << 25
		}
		if m.GetAccessHash() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetFirstName()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetLastName()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetUsername()) > 0 {
			flags |= 1 << 3
		}
		if len(m.GetPhone()) > 0 {
			flags |= 1 << 4
		}
		if m.GetPhoto() != nil {
			flags |= 1 << 5
		}
		if m.GetStatus() != nil {
			flags |= 1 << 6
		}
		if m.GetBotInfoVersion() != 0 {
			flags |= 1 << 14
		}
		if len(m.GetRestrictionReason938458C1117()) > 0 {
			flags |= 1 << 18
		}
		if len(m.GetBotInlinePlaceholder()) > 0 {
			flags |= 1 << 19
		}
		if len(m.GetLangCode()) > 0 {
			flags |= 1 << 22
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetAccessHash() != 0 {
			x.Long(m.GetAccessHash())
		}
		if len(m.GetFirstName()) > 0 {
			x.String(m.GetFirstName())
		}
		if len(m.GetLastName()) > 0 {
			x.String(m.GetLastName())
		}
		if len(m.GetUsername()) > 0 {
			x.String(m.GetUsername())
		}
		if len(m.GetPhone()) > 0 {
			x.String(m.GetPhone())
		}
		if m.GetPhoto() != nil {
			x.Bytes(m.GetPhoto().Encode(layer))
		}
		if m.GetStatus() != nil {
			x.Bytes(m.GetStatus().Encode(layer))
		}
		if m.GetBotInfoVersion() != 0 {
			x.Int(m.GetBotInfoVersion())
		}
		if len(m.GetRestrictionReason938458C1117()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetRestrictionReason938458C1117())))
			for _, v := range m.GetRestrictionReason938458C1117() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if len(m.GetBotInlinePlaceholder()) > 0 {
			x.String(m.GetBotInlinePlaceholder())
		}
		if len(m.GetLangCode()) > 0 {
			x.String(m.GetLangCode())
		}

		return x.GetBuf()
	case Cmd_Userd10d979a:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_Userd10d979a))
		// flags
		var flags int32 = 0
		if m.GetSelf() == true {
			flags |= 1 << 10
		}
		if m.GetContact() == true {
			flags |= 1 << 11
		}
		if m.GetMutualContact() == true {
			flags |= 1 << 12
		}
		if m.GetDeleted() == true {
			flags |= 1 << 13
		}
		if m.GetBot() == true {
			flags |= 1 << 14
		}
		if m.GetBotChatHistory() == true {
			flags |= 1 << 15
		}
		if m.GetBotNochats() == true {
			flags |= 1 << 16
		}
		if m.GetVerified() == true {
			flags |= 1 << 17
		}
		if m.GetRestricted() == true {
			flags |= 1 << 18
		}
		if m.GetMin() == true {
			flags |= 1 << 20
		}
		if m.GetBotInlineGeo() == true {
			flags |= 1 << 21
		}
		if m.GetAccessHash() != 0 {
			flags |= 1 << 0
		}
		if len(m.GetFirstName()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetLastName()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetUsername()) > 0 {
			flags |= 1 << 3
		}
		if len(m.GetPhone()) > 0 {
			flags |= 1 << 4
		}
		if m.GetPhoto() != nil {
			flags |= 1 << 5
		}
		if m.GetStatus() != nil {
			flags |= 1 << 6
		}
		if m.GetBotInfoVersion() != 0 {
			flags |= 1 << 14
		}
		if len(m.GetRestrictionReason2E13F4C371()) > 0 {
			flags |= 1 << 18
		}
		if len(m.GetBotInlinePlaceholder()) > 0 {
			flags |= 1 << 19
		}
		x.Int(flags)

		x.Int(m.GetId())
		if m.GetAccessHash() != 0 {
			x.Long(m.GetAccessHash())
		}
		if len(m.GetFirstName()) > 0 {
			x.String(m.GetFirstName())
		}
		if len(m.GetLastName()) > 0 {
			x.String(m.GetLastName())
		}
		if len(m.GetUsername()) > 0 {
			x.String(m.GetUsername())
		}
		if len(m.GetPhone()) > 0 {
			x.String(m.GetPhone())
		}
		if m.GetPhoto() != nil {
			x.Bytes(m.GetPhoto().Encode(layer))
		}
		if m.GetStatus() != nil {
			x.Bytes(m.GetStatus().Encode(layer))
		}
		if m.GetBotInfoVersion() != 0 {
			x.Int(m.GetBotInfoVersion())
		}
		if len(m.GetRestrictionReason2E13F4C371()) > 0 {
			x.String(m.GetRestrictionReason2E13F4C371())
		}
		if len(m.GetBotInlinePlaceholder()) > 0 {
			x.String(m.GetBotInlinePlaceholder())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid User encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUserFull) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUserFull, layer)

	switch cmd {
	case Cmd_UserFull:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserFull))
		// flags
		var flags int32 = 0
		if m.GetBlocked() == true {
			flags |= 1 << 0
		}
		if m.GetPhoneCallsAvailable() == true {
			flags |= 1 << 4
		}
		if m.GetPhoneCallsPrivate() == true {
			flags |= 1 << 5
		}
		if len(m.GetAbout()) > 0 {
			flags |= 1 << 1
		}
		if m.GetProfilePhoto() != nil {
			flags |= 1 << 2
		}
		if m.GetBotInfo() != nil {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Bytes(m.GetUser().Encode(layer))
		if len(m.GetAbout()) > 0 {
			x.String(m.GetAbout())
		}
		x.Bytes(m.GetLink().Encode(layer))
		if m.GetProfilePhoto() != nil {
			x.Bytes(m.GetProfilePhoto().Encode(layer))
		}
		x.Bytes(m.GetNotifySettings().Encode(layer))
		if m.GetBotInfo() != nil {
			x.Bytes(m.GetBotInfo().Encode(layer))
		}
		x.Int(m.GetCommonChatsCount())

		return x.GetBuf()
	case Cmd_UserFull5932fc03:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserFull5932fc03))
		// flags
		var flags int32 = 0
		if m.GetBlocked() == true {
			flags |= 1 << 0
		}
		if len(m.GetAbout()) > 0 {
			flags |= 1 << 1
		}
		if m.GetProfilePhoto() != nil {
			flags |= 1 << 2
		}
		if m.GetBotInfo() != nil {
			flags |= 1 << 3
		}
		x.Int(flags)

		x.Bytes(m.GetUser().Encode(layer))
		if len(m.GetAbout()) > 0 {
			x.String(m.GetAbout())
		}
		x.Bytes(m.GetLink().Encode(layer))
		if m.GetProfilePhoto() != nil {
			x.Bytes(m.GetProfilePhoto().Encode(layer))
		}
		x.Bytes(m.GetNotifySettings().Encode(layer))
		if m.GetBotInfo() != nil {
			x.Bytes(m.GetBotInfo().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_UserFull745559cc:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserFull745559cc))
		// flags
		var flags int32 = 0
		if m.GetBlocked() == true {
			flags |= 1 << 0
		}
		if m.GetPhoneCallsAvailable() == true {
			flags |= 1 << 4
		}
		if m.GetPhoneCallsPrivate() == true {
			flags |= 1 << 5
		}
		if m.GetCanPinMessage() == true {
			flags |= 1 << 7
		}
		if len(m.GetAbout()) > 0 {
			flags |= 1 << 1
		}
		if m.GetProfilePhoto() != nil {
			flags |= 1 << 2
		}
		if m.GetBotInfo() != nil {
			flags |= 1 << 3
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 6
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 11
		}
		x.Int(flags)

		x.Bytes(m.GetUser().Encode(layer))
		if len(m.GetAbout()) > 0 {
			x.String(m.GetAbout())
		}
		x.Bytes(m.GetLink().Encode(layer))
		if m.GetProfilePhoto() != nil {
			x.Bytes(m.GetProfilePhoto().Encode(layer))
		}
		x.Bytes(m.GetNotifySettings().Encode(layer))
		if m.GetBotInfo() != nil {
			x.Bytes(m.GetBotInfo().Encode(layer))
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		x.Int(m.GetCommonChatsCount())
		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}

		return x.GetBuf()
	case Cmd_UserFull8ea4a881:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserFull8ea4a881))
		// flags
		var flags int32 = 0
		if m.GetBlocked() == true {
			flags |= 1 << 0
		}
		if m.GetPhoneCallsAvailable() == true {
			flags |= 1 << 4
		}
		if m.GetPhoneCallsPrivate() == true {
			flags |= 1 << 5
		}
		if m.GetCanPinMessage() == true {
			flags |= 1 << 7
		}
		if len(m.GetAbout()) > 0 {
			flags |= 1 << 1
		}
		if m.GetProfilePhoto() != nil {
			flags |= 1 << 2
		}
		if m.GetBotInfo() != nil {
			flags |= 1 << 3
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 6
		}
		x.Int(flags)

		x.Bytes(m.GetUser().Encode(layer))
		if len(m.GetAbout()) > 0 {
			x.String(m.GetAbout())
		}
		x.Bytes(m.GetLink().Encode(layer))
		if m.GetProfilePhoto() != nil {
			x.Bytes(m.GetProfilePhoto().Encode(layer))
		}
		x.Bytes(m.GetNotifySettings().Encode(layer))
		if m.GetBotInfo() != nil {
			x.Bytes(m.GetBotInfo().Encode(layer))
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		x.Int(m.GetCommonChatsCount())

		return x.GetBuf()
	case Cmd_UserFulledf17c12:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserFulledf17c12))
		// flags
		var flags int32 = 0
		if m.GetBlocked() == true {
			flags |= 1 << 0
		}
		if m.GetPhoneCallsAvailable() == true {
			flags |= 1 << 4
		}
		if m.GetPhoneCallsPrivate() == true {
			flags |= 1 << 5
		}
		if m.GetCanPinMessage() == true {
			flags |= 1 << 7
		}
		if m.GetHasScheduled() == true {
			flags |= 1 << 12
		}
		if m.GetVideoCallsAvailable() == true {
			flags |= 1 << 13
		}
		if len(m.GetAbout()) > 0 {
			flags |= 1 << 1
		}
		if m.GetProfilePhoto() != nil {
			flags |= 1 << 2
		}
		if m.GetBotInfo() != nil {
			flags |= 1 << 3
		}
		if m.GetPinnedMsgId() != 0 {
			flags |= 1 << 6
		}
		if m.GetFolderId() != 0 {
			flags |= 1 << 11
		}
		x.Int(flags)

		x.Bytes(m.GetUser().Encode(layer))
		if len(m.GetAbout()) > 0 {
			x.String(m.GetAbout())
		}
		x.Bytes(m.GetSettings().Encode(layer))
		if m.GetProfilePhoto() != nil {
			x.Bytes(m.GetProfilePhoto().Encode(layer))
		}
		x.Bytes(m.GetNotifySettings().Encode(layer))
		if m.GetBotInfo() != nil {
			x.Bytes(m.GetBotInfo().Encode(layer))
		}
		if m.GetPinnedMsgId() != 0 {
			x.Int(m.GetPinnedMsgId())
		}
		x.Int(m.GetCommonChatsCount())
		if m.GetFolderId() != 0 {
			x.Int(m.GetFolderId())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid UserFull encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUserProfilePhotoEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUserProfilePhotoEmpty, layer)

	switch cmd {
	case Cmd_UserProfilePhotoEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserProfilePhotoEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UserProfilePhotoEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUserProfilePhoto) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUserProfilePhoto, layer)

	switch cmd {
	case Cmd_UserProfilePhoto:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserProfilePhoto))
		x.Long(m.GetPhotoId())
		x.Bytes(m.GetPhotoSmall().Encode(layer))
		x.Bytes(m.GetPhotoBig().Encode(layer))

		return x.GetBuf()
	case Cmd_UserProfilePhoto69d3ab26:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserProfilePhoto69d3ab26))
		// flags
		var flags int32 = 0
		if m.GetHasVideo() == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.Long(m.GetPhotoId())
		x.Bytes(m.GetPhotoSmall().Encode(layer))
		x.Bytes(m.GetPhotoBig().Encode(layer))
		x.Int(m.GetDcId())

		return x.GetBuf()
	case Cmd_UserProfilePhotoecd75d8c:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserProfilePhotoecd75d8c))
		x.Long(m.GetPhotoId())
		x.Bytes(m.GetPhotoSmall().Encode(layer))
		x.Bytes(m.GetPhotoBig().Encode(layer))
		x.Int(m.GetDcId())

		return x.GetBuf()

	default:
		log.Errorf("invalid UserProfilePhoto encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUserStatusEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUserStatusEmpty, layer)

	switch cmd {
	case Cmd_UserStatusEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserStatusEmpty))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UserStatusEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUserStatusOnline) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUserStatusOnline, layer)

	switch cmd {
	case Cmd_UserStatusOnline:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserStatusOnline))
		x.Int(m.GetExpires())

		return x.GetBuf()

	default:
		log.Errorf("invalid UserStatusOnline encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUserStatusOffline) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUserStatusOffline, layer)

	switch cmd {
	case Cmd_UserStatusOffline:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserStatusOffline))
		x.Int(m.GetWasOnline())

		return x.GetBuf()

	default:
		log.Errorf("invalid UserStatusOffline encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUserStatusRecently) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUserStatusRecently, layer)

	switch cmd {
	case Cmd_UserStatusRecently:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserStatusRecently))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UserStatusRecently encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUserStatusLastWeek) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUserStatusLastWeek, layer)

	switch cmd {
	case Cmd_UserStatusLastWeek:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserStatusLastWeek))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UserStatusLastWeek encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLUserStatusLastMonth) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassUserStatusLastMonth, layer)

	switch cmd {
	case Cmd_UserStatusLastMonth:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_UserStatusLastMonth))
		_ = m

		return x.GetBuf()

	default:
		log.Errorf("invalid UserStatusLastMonth encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLVideoSize) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassVideoSize, layer)

	switch cmd {
	case Cmd_VideoSize:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_VideoSize))
		x.String(m.GetType())
		x.Bytes(m.GetLocation().Encode(layer))
		x.Int(m.GetW())
		x.Int(m.GetH())
		x.Int(m.GetSize_())

		return x.GetBuf()
	case Cmd_VideoSizee831c556:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_VideoSizee831c556))
		// flags
		var flags int32 = 0
		if doubleIsEqual(m.GetVideoStartTs(), 0, 0.00000001) == true {
			flags |= 1 << 0
		}
		x.Int(flags)

		x.String(m.GetType())
		x.Bytes(m.GetLocation().Encode(layer))
		x.Int(m.GetW())
		x.Int(m.GetH())
		x.Int(m.GetSize_())

		return x.GetBuf()

	default:
		log.Errorf("invalid VideoSize encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWallPaper) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWallPaper, layer)

	switch cmd {
	case Cmd_WallPaper:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WallPaper))
		x.Int(m.GetIdCCB0365771())
		x.String(m.GetTitle())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetSizes())))
		for _, v := range m.GetSizes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetColor())

		return x.GetBuf()
	case Cmd_WallPapera437c3ed:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WallPapera437c3ed))
		x.Long(m.GetIdF04F91EC93())
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetDefault() == true {
			flags |= 1 << 1
		}
		if m.GetPattern() == true {
			flags |= 1 << 3
		}
		if m.GetDark() == true {
			flags |= 1 << 4
		}
		if m.GetSettings() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		x.Long(m.GetAccessHash())
		x.String(m.GetSlug())
		x.Bytes(m.GetDocument().Encode(layer))
		if m.GetSettings() != nil {
			x.Bytes(m.GetSettings().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_WallPaperf04f91ec:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WallPaperf04f91ec))
		x.Long(m.GetIdF04F91EC93())
		// flags
		var flags int32 = 0
		if m.GetCreator() == true {
			flags |= 1 << 0
		}
		if m.GetDefault() == true {
			flags |= 1 << 1
		}
		x.Int(flags)

		x.Long(m.GetAccessHash())
		x.String(m.GetSlug())
		x.Bytes(m.GetDocument().Encode(layer))

		return x.GetBuf()

	default:
		log.Errorf("invalid WallPaper encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWallPaperSolid) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWallPaperSolid, layer)

	switch cmd {
	case Cmd_WallPaperSolid:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WallPaperSolid))
		x.Int(m.GetIdCCB0365771())
		x.String(m.GetTitle())
		x.Int(m.GetBgColor())
		x.Int(m.GetColor())

		return x.GetBuf()

	default:
		log.Errorf("invalid WallPaperSolid encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWallPaperNoFile) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWallPaperNoFile, layer)

	switch cmd {
	case Cmd_WallPaperNoFile:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WallPaperNoFile))
		// flags
		var flags int32 = 0
		if m.GetDefault() == true {
			flags |= 1 << 1
		}
		if m.GetDark() == true {
			flags |= 1 << 4
		}
		if m.GetSettings() != nil {
			flags |= 1 << 2
		}
		x.Int(flags)

		if m.GetSettings() != nil {
			x.Bytes(m.GetSettings().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid WallPaperNoFile encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWallPaperSettings) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWallPaperSettings, layer)

	switch cmd {
	case Cmd_WallPaperSettings:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WallPaperSettings))
		// flags
		var flags int32 = 0
		if m.GetBlur() == true {
			flags |= 1 << 1
		}
		if m.GetMotion() == true {
			flags |= 1 << 2
		}
		if m.GetBackgroundColor() != 0 {
			flags |= 1 << 0
		}
		if m.GetIntensity() != 0 {
			flags |= 1 << 3
		}
		x.Int(flags)

		if m.GetBackgroundColor() != 0 {
			x.Int(m.GetBackgroundColor())
		}
		if m.GetIntensity() != 0 {
			x.Int(m.GetIntensity())
		}

		return x.GetBuf()
	case Cmd_WallPaperSettings5086cf8:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WallPaperSettings5086cf8))
		// flags
		var flags int32 = 0
		if m.GetBlur() == true {
			flags |= 1 << 1
		}
		if m.GetMotion() == true {
			flags |= 1 << 2
		}
		if m.GetBackgroundColor() != 0 {
			flags |= 1 << 0
		}
		if m.GetSecondBackgroundColor() != 0 {
			flags |= 1 << 4
		}
		if m.GetIntensity() != 0 {
			flags |= 1 << 3
		}
		if m.GetRotation() != 0 {
			flags |= 1 << 4
		}
		x.Int(flags)

		if m.GetBackgroundColor() != 0 {
			x.Int(m.GetBackgroundColor())
		}
		if m.GetSecondBackgroundColor() != 0 {
			x.Int(m.GetSecondBackgroundColor())
		}
		if m.GetIntensity() != 0 {
			x.Int(m.GetIntensity())
		}
		if m.GetRotation() != 0 {
			x.Int(m.GetRotation())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid WallPaperSettings encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWalletSecretSalt) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWalletSecretSalt, layer)

	switch cmd {
	case Cmd_WalletSecretSalt:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WalletSecretSalt))
		x.StringBytes(m.GetSalt())

		return x.GetBuf()

	default:
		log.Errorf("invalid WalletSecretSalt encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWalletLiteResponse) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWalletLiteResponse, layer)

	switch cmd {
	case Cmd_WalletLiteResponse:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WalletLiteResponse))
		x.StringBytes(m.GetResponse())

		return x.GetBuf()

	default:
		log.Errorf("invalid WalletLiteResponse encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWebAuthorization) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWebAuthorization, layer)

	switch cmd {
	case Cmd_WebAuthorization:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebAuthorization))
		x.Long(m.GetHash())
		x.Int(m.GetBotId())
		x.String(m.GetDomain())
		x.String(m.GetBrowser())
		x.String(m.GetPlatform())
		x.Int(m.GetDateCreated())
		x.Int(m.GetDateActive())
		x.String(m.GetIp())
		x.String(m.GetRegion())

		return x.GetBuf()

	default:
		log.Errorf("invalid WebAuthorization encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWebDocument) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWebDocument, layer)

	switch cmd {
	case Cmd_WebDocument:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebDocument))
		x.String(m.GetUrl())
		x.Long(m.GetAccessHash())
		x.Int(m.GetSize_())
		x.String(m.GetMimeType())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}
		x.Int(m.GetDcId())

		return x.GetBuf()
	case Cmd_WebDocument1c570ed1:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebDocument1c570ed1))
		x.String(m.GetUrl())
		x.Long(m.GetAccessHash())
		x.Int(m.GetSize_())
		x.String(m.GetMimeType())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid WebDocument encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWebDocumentNoProxy) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWebDocumentNoProxy, layer)

	switch cmd {
	case Cmd_WebDocumentNoProxy:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebDocumentNoProxy))
		x.String(m.GetUrl())
		x.Int(m.GetSize_())
		x.String(m.GetMimeType())
		x.Int(int32(Cmd_vector))
		x.Int(int32(len(m.GetAttributes())))
		for _, v := range m.GetAttributes() {
			x.buf = append(x.buf, (*v).Encode(layer)...)
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid WebDocumentNoProxy encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWebPageEmpty) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWebPageEmpty, layer)

	switch cmd {
	case Cmd_WebPageEmpty:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebPageEmpty))
		x.Long(m.GetId())

		return x.GetBuf()

	default:
		log.Errorf("invalid WebPageEmpty encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWebPagePending) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWebPagePending, layer)

	switch cmd {
	case Cmd_WebPagePending:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebPagePending))
		x.Long(m.GetId())
		x.Int(m.GetDate())

		return x.GetBuf()

	default:
		log.Errorf("invalid WebPagePending encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWebPage) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWebPage, layer)

	switch cmd {
	case Cmd_WebPage:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebPage))
		// flags
		var flags int32 = 0
		if len(m.GetType()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetSiteName()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 3
		}
		if m.GetPhoto() != nil {
			flags |= 1 << 4
		}
		if len(m.GetEmbedUrl()) > 0 {
			flags |= 1 << 5
		}
		if len(m.GetEmbedType()) > 0 {
			flags |= 1 << 5
		}
		if m.GetEmbedWidth() != 0 {
			flags |= 1 << 6
		}
		if m.GetEmbedHeight() != 0 {
			flags |= 1 << 6
		}
		if m.GetDuration() != 0 {
			flags |= 1 << 7
		}
		if len(m.GetAuthor()) > 0 {
			flags |= 1 << 8
		}
		if m.GetDocument() != nil {
			flags |= 1 << 9
		}
		if m.GetCachedPage() != nil {
			flags |= 1 << 10
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.String(m.GetUrl())
		x.String(m.GetDisplayUrl())
		x.Int(m.GetHash())
		if len(m.GetType()) > 0 {
			x.String(m.GetType())
		}
		if len(m.GetSiteName()) > 0 {
			x.String(m.GetSiteName())
		}
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		if m.GetPhoto() != nil {
			x.Bytes(m.GetPhoto().Encode(layer))
		}
		if len(m.GetEmbedUrl()) > 0 {
			x.String(m.GetEmbedUrl())
		}
		if len(m.GetEmbedType()) > 0 {
			x.String(m.GetEmbedType())
		}
		if m.GetEmbedWidth() != 0 {
			x.Int(m.GetEmbedWidth())
		}
		if m.GetEmbedHeight() != 0 {
			x.Int(m.GetEmbedHeight())
		}
		if m.GetDuration() != 0 {
			x.Int(m.GetDuration())
		}
		if len(m.GetAuthor()) > 0 {
			x.String(m.GetAuthor())
		}
		if m.GetDocument() != nil {
			x.Bytes(m.GetDocument().Encode(layer))
		}
		if m.GetCachedPage() != nil {
			x.Bytes(m.GetCachedPage().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_WebPageca820ed7:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebPageca820ed7))
		// flags
		var flags int32 = 0
		if len(m.GetType()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetSiteName()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 3
		}
		if m.GetPhoto() != nil {
			flags |= 1 << 4
		}
		if len(m.GetEmbedUrl()) > 0 {
			flags |= 1 << 5
		}
		if len(m.GetEmbedType()) > 0 {
			flags |= 1 << 5
		}
		if m.GetEmbedWidth() != 0 {
			flags |= 1 << 6
		}
		if m.GetEmbedHeight() != 0 {
			flags |= 1 << 6
		}
		if m.GetDuration() != 0 {
			flags |= 1 << 7
		}
		if len(m.GetAuthor()) > 0 {
			flags |= 1 << 8
		}
		if m.GetDocument() != nil {
			flags |= 1 << 9
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.String(m.GetUrl())
		x.String(m.GetDisplayUrl())
		if len(m.GetType()) > 0 {
			x.String(m.GetType())
		}
		if len(m.GetSiteName()) > 0 {
			x.String(m.GetSiteName())
		}
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		if m.GetPhoto() != nil {
			x.Bytes(m.GetPhoto().Encode(layer))
		}
		if len(m.GetEmbedUrl()) > 0 {
			x.String(m.GetEmbedUrl())
		}
		if len(m.GetEmbedType()) > 0 {
			x.String(m.GetEmbedType())
		}
		if m.GetEmbedWidth() != 0 {
			x.Int(m.GetEmbedWidth())
		}
		if m.GetEmbedHeight() != 0 {
			x.Int(m.GetEmbedHeight())
		}
		if m.GetDuration() != 0 {
			x.Int(m.GetDuration())
		}
		if len(m.GetAuthor()) > 0 {
			x.String(m.GetAuthor())
		}
		if m.GetDocument() != nil {
			x.Bytes(m.GetDocument().Encode(layer))
		}

		return x.GetBuf()
	case Cmd_WebPagee89c45b2:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebPagee89c45b2))
		// flags
		var flags int32 = 0
		if len(m.GetType()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetSiteName()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 3
		}
		if m.GetPhoto() != nil {
			flags |= 1 << 4
		}
		if len(m.GetEmbedUrl()) > 0 {
			flags |= 1 << 5
		}
		if len(m.GetEmbedType()) > 0 {
			flags |= 1 << 5
		}
		if m.GetEmbedWidth() != 0 {
			flags |= 1 << 6
		}
		if m.GetEmbedHeight() != 0 {
			flags |= 1 << 6
		}
		if m.GetDuration() != 0 {
			flags |= 1 << 7
		}
		if len(m.GetAuthor()) > 0 {
			flags |= 1 << 8
		}
		if m.GetDocument() != nil {
			flags |= 1 << 9
		}
		if m.GetCachedPage() != nil {
			flags |= 1 << 10
		}
		if len(m.GetAttributes()) > 0 {
			flags |= 1 << 12
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.String(m.GetUrl())
		x.String(m.GetDisplayUrl())
		x.Int(m.GetHash())
		if len(m.GetType()) > 0 {
			x.String(m.GetType())
		}
		if len(m.GetSiteName()) > 0 {
			x.String(m.GetSiteName())
		}
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		if m.GetPhoto() != nil {
			x.Bytes(m.GetPhoto().Encode(layer))
		}
		if len(m.GetEmbedUrl()) > 0 {
			x.String(m.GetEmbedUrl())
		}
		if len(m.GetEmbedType()) > 0 {
			x.String(m.GetEmbedType())
		}
		if m.GetEmbedWidth() != 0 {
			x.Int(m.GetEmbedWidth())
		}
		if m.GetEmbedHeight() != 0 {
			x.Int(m.GetEmbedHeight())
		}
		if m.GetDuration() != 0 {
			x.Int(m.GetDuration())
		}
		if len(m.GetAuthor()) > 0 {
			x.String(m.GetAuthor())
		}
		if m.GetDocument() != nil {
			x.Bytes(m.GetDocument().Encode(layer))
		}
		if m.GetCachedPage() != nil {
			x.Bytes(m.GetCachedPage().Encode(layer))
		}
		if len(m.GetAttributes()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetAttributes())))
			for _, v := range m.GetAttributes() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}

		return x.GetBuf()
	case Cmd_WebPagefa64e172:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebPagefa64e172))
		// flags
		var flags int32 = 0
		if len(m.GetType()) > 0 {
			flags |= 1 << 0
		}
		if len(m.GetSiteName()) > 0 {
			flags |= 1 << 1
		}
		if len(m.GetTitle()) > 0 {
			flags |= 1 << 2
		}
		if len(m.GetDescription()) > 0 {
			flags |= 1 << 3
		}
		if m.GetPhoto() != nil {
			flags |= 1 << 4
		}
		if len(m.GetEmbedUrl()) > 0 {
			flags |= 1 << 5
		}
		if len(m.GetEmbedType()) > 0 {
			flags |= 1 << 5
		}
		if m.GetEmbedWidth() != 0 {
			flags |= 1 << 6
		}
		if m.GetEmbedHeight() != 0 {
			flags |= 1 << 6
		}
		if m.GetDuration() != 0 {
			flags |= 1 << 7
		}
		if len(m.GetAuthor()) > 0 {
			flags |= 1 << 8
		}
		if m.GetDocument() != nil {
			flags |= 1 << 9
		}
		if len(m.GetDocuments()) > 0 {
			flags |= 1 << 11
		}
		if m.GetCachedPage() != nil {
			flags |= 1 << 10
		}
		x.Int(flags)

		x.Long(m.GetId())
		x.String(m.GetUrl())
		x.String(m.GetDisplayUrl())
		x.Int(m.GetHash())
		if len(m.GetType()) > 0 {
			x.String(m.GetType())
		}
		if len(m.GetSiteName()) > 0 {
			x.String(m.GetSiteName())
		}
		if len(m.GetTitle()) > 0 {
			x.String(m.GetTitle())
		}
		if len(m.GetDescription()) > 0 {
			x.String(m.GetDescription())
		}
		if m.GetPhoto() != nil {
			x.Bytes(m.GetPhoto().Encode(layer))
		}
		if len(m.GetEmbedUrl()) > 0 {
			x.String(m.GetEmbedUrl())
		}
		if len(m.GetEmbedType()) > 0 {
			x.String(m.GetEmbedType())
		}
		if m.GetEmbedWidth() != 0 {
			x.Int(m.GetEmbedWidth())
		}
		if m.GetEmbedHeight() != 0 {
			x.Int(m.GetEmbedHeight())
		}
		if m.GetDuration() != 0 {
			x.Int(m.GetDuration())
		}
		if len(m.GetAuthor()) > 0 {
			x.String(m.GetAuthor())
		}
		if m.GetDocument() != nil {
			x.Bytes(m.GetDocument().Encode(layer))
		}
		if len(m.GetDocuments()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetDocuments())))
			for _, v := range m.GetDocuments() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetCachedPage() != nil {
			x.Bytes(m.GetCachedPage().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid WebPage encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWebPageNotModified) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWebPageNotModified, layer)

	switch cmd {
	case Cmd_WebPageNotModified:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebPageNotModified))
		_ = m

		return x.GetBuf()
	case Cmd_WebPageNotModified7311ca11:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebPageNotModified7311ca11))
		// flags
		var flags int32 = 0
		if m.GetCachedPageViews() != 0 {
			flags |= 1 << 0
		}
		x.Int(flags)

		if m.GetCachedPageViews() != 0 {
			x.Int(m.GetCachedPageViews())
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid WebPageNotModified encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}

func (m *TLWebPageAttributeTheme) Encode(layer int32) []byte {

	cmd := GetClassIDByLayer(ClassWebPageAttributeTheme, layer)

	switch cmd {
	case Cmd_WebPageAttributeTheme:
		x := NewEncodeBuf(512)
		x.Int(int32(Cmd_WebPageAttributeTheme))
		// flags
		var flags int32 = 0
		if len(m.GetDocuments()) > 0 {
			flags |= 1 << 0
		}
		if m.GetSettings() != nil {
			flags |= 1 << 1
		}
		x.Int(flags)

		if len(m.GetDocuments()) > 0 {
			x.Int(int32(Cmd_vector))
			x.Int(int32(len(m.GetDocuments())))
			for _, v := range m.GetDocuments() {
				x.buf = append(x.buf, (*v).Encode(layer)...)
			}
		}
		if m.GetSettings() != nil {
			x.Bytes(m.GetSettings().Encode(layer))
		}

		return x.GetBuf()

	default:
		log.Errorf("invalid WebPageAttributeTheme encode layer:%d cmd: %x", layer, cmd.ToUInt32())
		return nil
	}
	return nil
}
