/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2023/06/12 18:03
 * @File : codec_schema_member.pb.go
 * @Desc :
 *
 */

package mtproto

func (m *TLAccessPointRule) SetPhonePrefixRules(v string) { m.Data2.PhonePrefixRules = v }
func (m *TLAccessPointRule) GetPhonePrefixRules() string  { return m.Data2.PhonePrefixRules }

func (m *TLAccessPointRule) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLAccessPointRule) GetDcId() int32  { return m.Data2.DcId }

func (m *TLAccessPointRule) SetIps(v []*IpPort) { m.Data2.Ips = v }
func (m *TLAccessPointRule) GetIps() []*IpPort  { return m.Data2.Ips }

func (m *TLAccountDaysTTL) SetDays(v int32) { m.Data2.Days = v }
func (m *TLAccountDaysTTL) GetDays() int32  { return m.Data2.Days }

func (m *TLAccountAuthorizationForm) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLAccountAuthorizationForm) GetFlags() int32  { return m.Data2.Flags }

func (m *TLAccountAuthorizationForm) SetRequiredTypes(v []*SecureRequiredType) {
	m.Data2.RequiredTypes = v
}
func (m *TLAccountAuthorizationForm) GetRequiredTypes() []*SecureRequiredType {
	return m.Data2.RequiredTypes
}

func (m *TLAccountAuthorizationForm) SetValues(v []*SecureValue) { m.Data2.Values = v }
func (m *TLAccountAuthorizationForm) GetValues() []*SecureValue  { return m.Data2.Values }

func (m *TLAccountAuthorizationForm) SetErrors(v []*SecureValueError) { m.Data2.Errors = v }
func (m *TLAccountAuthorizationForm) GetErrors() []*SecureValueError  { return m.Data2.Errors }

func (m *TLAccountAuthorizationForm) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLAccountAuthorizationForm) GetUsers() []*User  { return m.Data2.Users }

func (m *TLAccountAuthorizationForm) SetPrivacyPolicyUrl(v string) { m.Data2.PrivacyPolicyUrl = v }
func (m *TLAccountAuthorizationForm) GetPrivacyPolicyUrl() string  { return m.Data2.PrivacyPolicyUrl }

func (m *TLAccountAuthorizations) SetAuthorizations(v []*Authorization) { m.Data2.Authorizations = v }
func (m *TLAccountAuthorizations) GetAuthorizations() []*Authorization  { return m.Data2.Authorizations }

func (m *TLAccountAutoDownloadSettings) SetLow(v *AutoDownloadSettings) { m.Data2.Low = v }
func (m *TLAccountAutoDownloadSettings) GetLow() *AutoDownloadSettings  { return m.Data2.Low }

func (m *TLAccountAutoDownloadSettings) SetMedium(v *AutoDownloadSettings) { m.Data2.Medium = v }
func (m *TLAccountAutoDownloadSettings) GetMedium() *AutoDownloadSettings  { return m.Data2.Medium }

func (m *TLAccountAutoDownloadSettings) SetHigh(v *AutoDownloadSettings) { m.Data2.High = v }
func (m *TLAccountAutoDownloadSettings) GetHigh() *AutoDownloadSettings  { return m.Data2.High }

func (m *TLAccountContentSettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLAccountContentSettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLAccountContentSettings) SetSensitiveEnabled(v bool) { m.Data2.SensitiveEnabled = v }
func (m *TLAccountContentSettings) GetSensitiveEnabled() bool  { return m.Data2.SensitiveEnabled }

func (m *TLAccountContentSettings) SetSensitiveCanChange(v bool) { m.Data2.SensitiveCanChange = v }
func (m *TLAccountContentSettings) GetSensitiveCanChange() bool  { return m.Data2.SensitiveCanChange }

func (m *TLAccountNoPassword) SetNewSalt(v []byte) { m.Data2.NewSalt = v }
func (m *TLAccountNoPassword) GetNewSalt() []byte  { return m.Data2.NewSalt }

func (m *TLAccountNoPassword) SetEmailUnconfirmedPattern(v string) {
	m.Data2.EmailUnconfirmedPattern = v
}
func (m *TLAccountNoPassword) GetEmailUnconfirmedPattern() string {
	return m.Data2.EmailUnconfirmedPattern
}

func (m *TLAccountPassword) SetCurrentSalt(v []byte) { m.Data2.CurrentSalt = v }
func (m *TLAccountPassword) GetCurrentSalt() []byte  { return m.Data2.CurrentSalt }

func (m *TLAccountPassword) SetNewSalt(v []byte) { m.Data2.NewSalt = v }
func (m *TLAccountPassword) GetNewSalt() []byte  { return m.Data2.NewSalt }

func (m *TLAccountPassword) SetHint(v string) { m.Data2.Hint = v }
func (m *TLAccountPassword) GetHint() string  { return m.Data2.Hint }

func (m *TLAccountPassword) SetHasRecovery7C18141C71(v *Bool) { m.Data2.HasRecovery7C18141C71 = v }
func (m *TLAccountPassword) GetHasRecovery7C18141C71() *Bool  { return m.Data2.HasRecovery7C18141C71 }

func (m *TLAccountPassword) SetEmailUnconfirmedPattern(v string) { m.Data2.EmailUnconfirmedPattern = v }
func (m *TLAccountPassword) GetEmailUnconfirmedPattern() string {
	return m.Data2.EmailUnconfirmedPattern
}

func (m *TLAccountPassword) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLAccountPassword) GetFlags() int32  { return m.Data2.Flags }

func (m *TLAccountPassword) SetHasRecoveryAD2641F885(v bool) { m.Data2.HasRecoveryAD2641F885 = v }
func (m *TLAccountPassword) GetHasRecoveryAD2641F885() bool  { return m.Data2.HasRecoveryAD2641F885 }

func (m *TLAccountPassword) SetHasSecureValues(v bool) { m.Data2.HasSecureValues = v }
func (m *TLAccountPassword) GetHasSecureValues() bool  { return m.Data2.HasSecureValues }

func (m *TLAccountPassword) SetHasPassword(v bool) { m.Data2.HasPassword = v }
func (m *TLAccountPassword) GetHasPassword() bool  { return m.Data2.HasPassword }

func (m *TLAccountPassword) SetCurrentAlgo(v *PasswordKdfAlgo) { m.Data2.CurrentAlgo = v }
func (m *TLAccountPassword) GetCurrentAlgo() *PasswordKdfAlgo  { return m.Data2.CurrentAlgo }

func (m *TLAccountPassword) SetSrp_B(v []byte) { m.Data2.Srp_B = v }
func (m *TLAccountPassword) GetSrp_B() []byte  { return m.Data2.Srp_B }

func (m *TLAccountPassword) SetSrpId(v int64) { m.Data2.SrpId = v }
func (m *TLAccountPassword) GetSrpId() int64  { return m.Data2.SrpId }

func (m *TLAccountPassword) SetNewAlgo(v *PasswordKdfAlgo) { m.Data2.NewAlgo = v }
func (m *TLAccountPassword) GetNewAlgo() *PasswordKdfAlgo  { return m.Data2.NewAlgo }

func (m *TLAccountPassword) SetNewSecureAlgo(v *SecurePasswordKdfAlgo) { m.Data2.NewSecureAlgo = v }
func (m *TLAccountPassword) GetNewSecureAlgo() *SecurePasswordKdfAlgo  { return m.Data2.NewSecureAlgo }

func (m *TLAccountPassword) SetSecureRandom(v []byte) { m.Data2.SecureRandom = v }
func (m *TLAccountPassword) GetSecureRandom() []byte  { return m.Data2.SecureRandom }

func (m *TLAccountPasswordInputSettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLAccountPasswordInputSettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLAccountPasswordInputSettings) SetNewSalt(v []byte) { m.Data2.NewSalt = v }
func (m *TLAccountPasswordInputSettings) GetNewSalt() []byte  { return m.Data2.NewSalt }

func (m *TLAccountPasswordInputSettings) SetNewPasswordHash(v []byte) { m.Data2.NewPasswordHash = v }
func (m *TLAccountPasswordInputSettings) GetNewPasswordHash() []byte  { return m.Data2.NewPasswordHash }

func (m *TLAccountPasswordInputSettings) SetHint(v string) { m.Data2.Hint = v }
func (m *TLAccountPasswordInputSettings) GetHint() string  { return m.Data2.Hint }

func (m *TLAccountPasswordInputSettings) SetEmail(v string) { m.Data2.Email = v }
func (m *TLAccountPasswordInputSettings) GetEmail() string  { return m.Data2.Email }

func (m *TLAccountPasswordInputSettings) SetNewAlgo(v *PasswordKdfAlgo) { m.Data2.NewAlgo = v }
func (m *TLAccountPasswordInputSettings) GetNewAlgo() *PasswordKdfAlgo  { return m.Data2.NewAlgo }

func (m *TLAccountPasswordInputSettings) SetNewSecureSettings(v *SecureSecretSettings) {
	m.Data2.NewSecureSettings = v
}
func (m *TLAccountPasswordInputSettings) GetNewSecureSettings() *SecureSecretSettings {
	return m.Data2.NewSecureSettings
}

func (m *TLAccountPasswordSettings) SetEmail(v string) { m.Data2.Email = v }
func (m *TLAccountPasswordSettings) GetEmail() string  { return m.Data2.Email }

func (m *TLAccountPasswordSettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLAccountPasswordSettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLAccountPasswordSettings) SetSecureSettings(v *SecureSecretSettings) {
	m.Data2.SecureSettings = v
}
func (m *TLAccountPasswordSettings) GetSecureSettings() *SecureSecretSettings {
	return m.Data2.SecureSettings
}

func (m *TLAccountPrivacyRules) SetRules(v []*PrivacyRule) { m.Data2.Rules = v }
func (m *TLAccountPrivacyRules) GetRules() []*PrivacyRule  { return m.Data2.Rules }

func (m *TLAccountPrivacyRules) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLAccountPrivacyRules) GetUsers() []*User  { return m.Data2.Users }

func (m *TLAccountPrivacyRules) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLAccountPrivacyRules) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLAccountSentEmailCode) SetEmailPattern(v string) { m.Data2.EmailPattern = v }
func (m *TLAccountSentEmailCode) GetEmailPattern() string  { return m.Data2.EmailPattern }

func (m *TLAccountSentEmailCode) SetLength(v int32) { m.Data2.Length = v }
func (m *TLAccountSentEmailCode) GetLength() int32  { return m.Data2.Length }

func (m *TLAccountTakeout) SetId(v int64) { m.Data2.Id = v }
func (m *TLAccountTakeout) GetId() int64  { return m.Data2.Id }

func (m *TLAccountThemes) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLAccountThemes) GetHash() int32  { return m.Data2.Hash }

func (m *TLAccountThemes) SetThemes(v []*Theme) { m.Data2.Themes = v }
func (m *TLAccountThemes) GetThemes() []*Theme  { return m.Data2.Themes }

func (m *TLAccountTmpPassword) SetTmpPassword(v []byte) { m.Data2.TmpPassword = v }
func (m *TLAccountTmpPassword) GetTmpPassword() []byte  { return m.Data2.TmpPassword }

func (m *TLAccountTmpPassword) SetValidUntil(v int32) { m.Data2.ValidUntil = v }
func (m *TLAccountTmpPassword) GetValidUntil() int32  { return m.Data2.ValidUntil }

func (m *TLAccountWallPapers) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLAccountWallPapers) GetHash() int32  { return m.Data2.Hash }

func (m *TLAccountWallPapers) SetWallpapers(v []*WallPaper) { m.Data2.Wallpapers = v }
func (m *TLAccountWallPapers) GetWallpapers() []*WallPaper  { return m.Data2.Wallpapers }

func (m *TLAccountWebAuthorizations) SetAuthorizations(v []*WebAuthorization) {
	m.Data2.Authorizations = v
}
func (m *TLAccountWebAuthorizations) GetAuthorizations() []*WebAuthorization {
	return m.Data2.Authorizations
}

func (m *TLAccountWebAuthorizations) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLAccountWebAuthorizations) GetUsers() []*User  { return m.Data2.Users }

func (m *TLAuthAuthorization) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLAuthAuthorization) GetFlags() int32  { return m.Data2.Flags }

func (m *TLAuthAuthorization) SetTmpSessions(v int32) { m.Data2.TmpSessions = v }
func (m *TLAuthAuthorization) GetTmpSessions() int32  { return m.Data2.TmpSessions }

func (m *TLAuthAuthorization) SetUser(v *User) { m.Data2.User = v }
func (m *TLAuthAuthorization) GetUser() *User  { return m.Data2.User }

func (m *TLAuthAuthorizationSignUpRequired) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLAuthAuthorizationSignUpRequired) GetFlags() int32  { return m.Data2.Flags }

func (m *TLAuthAuthorizationSignUpRequired) SetTermsOfService(v *Help_TermsOfService) {
	m.Data2.TermsOfService = v
}
func (m *TLAuthAuthorizationSignUpRequired) GetTermsOfService() *Help_TermsOfService {
	return m.Data2.TermsOfService
}

func (m *TLAuthCheckedPhone) SetPhoneRegistered(v *Bool) { m.Data2.PhoneRegistered = v }
func (m *TLAuthCheckedPhone) GetPhoneRegistered() *Bool  { return m.Data2.PhoneRegistered }

func (m *TLAuthExportedAuthorization) SetId(v int32) { m.Data2.Id = v }
func (m *TLAuthExportedAuthorization) GetId() int32  { return m.Data2.Id }

func (m *TLAuthExportedAuthorization) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLAuthExportedAuthorization) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLAuthLoginToken) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLAuthLoginToken) GetExpires() int32  { return m.Data2.Expires }

func (m *TLAuthLoginToken) SetToken(v []byte) { m.Data2.Token = v }
func (m *TLAuthLoginToken) GetToken() []byte  { return m.Data2.Token }

func (m *TLAuthLoginTokenMigrateTo) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLAuthLoginTokenMigrateTo) GetDcId() int32  { return m.Data2.DcId }

func (m *TLAuthLoginTokenMigrateTo) SetToken(v []byte) { m.Data2.Token = v }
func (m *TLAuthLoginTokenMigrateTo) GetToken() []byte  { return m.Data2.Token }

func (m *TLAuthLoginTokenSuccess) SetAuthorization(v *Auth_Authorization) { m.Data2.Authorization = v }
func (m *TLAuthLoginTokenSuccess) GetAuthorization() *Auth_Authorization {
	return m.Data2.Authorization
}

func (m *TLAuthPasswordRecovery) SetEmailPattern(v string) { m.Data2.EmailPattern = v }
func (m *TLAuthPasswordRecovery) GetEmailPattern() string  { return m.Data2.EmailPattern }

func (m *TLAuthSentCode) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLAuthSentCode) GetFlags() int32  { return m.Data2.Flags }

func (m *TLAuthSentCode) SetPhoneRegistered(v bool) { m.Data2.PhoneRegistered = v }
func (m *TLAuthSentCode) GetPhoneRegistered() bool  { return m.Data2.PhoneRegistered }

func (m *TLAuthSentCode) SetType(v *Auth_SentCodeType) { m.Data2.Type = v }
func (m *TLAuthSentCode) GetType() *Auth_SentCodeType  { return m.Data2.Type }

func (m *TLAuthSentCode) SetPhoneCodeHash(v string) { m.Data2.PhoneCodeHash = v }
func (m *TLAuthSentCode) GetPhoneCodeHash() string  { return m.Data2.PhoneCodeHash }

func (m *TLAuthSentCode) SetNextType(v *Auth_CodeType) { m.Data2.NextType = v }
func (m *TLAuthSentCode) GetNextType() *Auth_CodeType  { return m.Data2.NextType }

func (m *TLAuthSentCode) SetTimeout(v int32) { m.Data2.Timeout = v }
func (m *TLAuthSentCode) GetTimeout() int32  { return m.Data2.Timeout }

func (m *TLAuthSentCode) SetTermsOfService(v *Help_TermsOfService) { m.Data2.TermsOfService = v }
func (m *TLAuthSentCode) GetTermsOfService() *Help_TermsOfService  { return m.Data2.TermsOfService }

func (m *TLAuthSentCodeTypeApp) SetLength(v int32) { m.Data2.Length = v }
func (m *TLAuthSentCodeTypeApp) GetLength() int32  { return m.Data2.Length }

func (m *TLAuthSentCodeTypeSms) SetLength(v int32) { m.Data2.Length = v }
func (m *TLAuthSentCodeTypeSms) GetLength() int32  { return m.Data2.Length }

func (m *TLAuthSentCodeTypeCall) SetLength(v int32) { m.Data2.Length = v }
func (m *TLAuthSentCodeTypeCall) GetLength() int32  { return m.Data2.Length }

func (m *TLAuthSentCodeTypeFlashCall) SetPattern(v string) { m.Data2.Pattern = v }
func (m *TLAuthSentCodeTypeFlashCall) GetPattern() string  { return m.Data2.Pattern }

func (m *TLAuthorization) SetHash(v int64) { m.Data2.Hash = v }
func (m *TLAuthorization) GetHash() int64  { return m.Data2.Hash }

func (m *TLAuthorization) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLAuthorization) GetFlags() int32  { return m.Data2.Flags }

func (m *TLAuthorization) SetDeviceModel(v string) { m.Data2.DeviceModel = v }
func (m *TLAuthorization) GetDeviceModel() string  { return m.Data2.DeviceModel }

func (m *TLAuthorization) SetPlatform(v string) { m.Data2.Platform = v }
func (m *TLAuthorization) GetPlatform() string  { return m.Data2.Platform }

func (m *TLAuthorization) SetSystemVersion(v string) { m.Data2.SystemVersion = v }
func (m *TLAuthorization) GetSystemVersion() string  { return m.Data2.SystemVersion }

func (m *TLAuthorization) SetApiId(v int32) { m.Data2.ApiId = v }
func (m *TLAuthorization) GetApiId() int32  { return m.Data2.ApiId }

func (m *TLAuthorization) SetAppName(v string) { m.Data2.AppName = v }
func (m *TLAuthorization) GetAppName() string  { return m.Data2.AppName }

func (m *TLAuthorization) SetAppVersion(v string) { m.Data2.AppVersion = v }
func (m *TLAuthorization) GetAppVersion() string  { return m.Data2.AppVersion }

func (m *TLAuthorization) SetDateCreated(v int32) { m.Data2.DateCreated = v }
func (m *TLAuthorization) GetDateCreated() int32  { return m.Data2.DateCreated }

func (m *TLAuthorization) SetDateActive(v int32) { m.Data2.DateActive = v }
func (m *TLAuthorization) GetDateActive() int32  { return m.Data2.DateActive }

func (m *TLAuthorization) SetIp(v string) { m.Data2.Ip = v }
func (m *TLAuthorization) GetIp() string  { return m.Data2.Ip }

func (m *TLAuthorization) SetCountry(v string) { m.Data2.Country = v }
func (m *TLAuthorization) GetCountry() string  { return m.Data2.Country }

func (m *TLAuthorization) SetRegion(v string) { m.Data2.Region = v }
func (m *TLAuthorization) GetRegion() string  { return m.Data2.Region }

func (m *TLAuthorization) SetCurrent(v bool) { m.Data2.Current = v }
func (m *TLAuthorization) GetCurrent() bool  { return m.Data2.Current }

func (m *TLAuthorization) SetOfficialApp(v bool) { m.Data2.OfficialApp = v }
func (m *TLAuthorization) GetOfficialApp() bool  { return m.Data2.OfficialApp }

func (m *TLAuthorization) SetPasswordPending(v bool) { m.Data2.PasswordPending = v }
func (m *TLAuthorization) GetPasswordPending() bool  { return m.Data2.PasswordPending }

func (m *TLAutoDownloadSettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLAutoDownloadSettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLAutoDownloadSettings) SetDisabled(v bool) { m.Data2.Disabled = v }
func (m *TLAutoDownloadSettings) GetDisabled() bool  { return m.Data2.Disabled }

func (m *TLAutoDownloadSettings) SetVideoPreloadLarge(v bool) { m.Data2.VideoPreloadLarge = v }
func (m *TLAutoDownloadSettings) GetVideoPreloadLarge() bool  { return m.Data2.VideoPreloadLarge }

func (m *TLAutoDownloadSettings) SetAudioPreloadNext(v bool) { m.Data2.AudioPreloadNext = v }
func (m *TLAutoDownloadSettings) GetAudioPreloadNext() bool  { return m.Data2.AudioPreloadNext }

func (m *TLAutoDownloadSettings) SetPhonecallsLessData(v bool) { m.Data2.PhonecallsLessData = v }
func (m *TLAutoDownloadSettings) GetPhonecallsLessData() bool  { return m.Data2.PhonecallsLessData }

func (m *TLAutoDownloadSettings) SetPhotoSizeMax(v int32) { m.Data2.PhotoSizeMax = v }
func (m *TLAutoDownloadSettings) GetPhotoSizeMax() int32  { return m.Data2.PhotoSizeMax }

func (m *TLAutoDownloadSettings) SetVideoSizeMax(v int32) { m.Data2.VideoSizeMax = v }
func (m *TLAutoDownloadSettings) GetVideoSizeMax() int32  { return m.Data2.VideoSizeMax }

func (m *TLAutoDownloadSettings) SetFileSizeMax(v int32) { m.Data2.FileSizeMax = v }
func (m *TLAutoDownloadSettings) GetFileSizeMax() int32  { return m.Data2.FileSizeMax }

func (m *TLAutoDownloadSettings) SetVideoUploadMaxbitrate(v int32) { m.Data2.VideoUploadMaxbitrate = v }
func (m *TLAutoDownloadSettings) GetVideoUploadMaxbitrate() int32 {
	return m.Data2.VideoUploadMaxbitrate
}

func (m *TLBadMsgNotification) SetBadMsgId(v int64) { m.Data2.BadMsgId = v }
func (m *TLBadMsgNotification) GetBadMsgId() int64  { return m.Data2.BadMsgId }

func (m *TLBadMsgNotification) SetBadMsgSeqno(v int32) { m.Data2.BadMsgSeqno = v }
func (m *TLBadMsgNotification) GetBadMsgSeqno() int32  { return m.Data2.BadMsgSeqno }

func (m *TLBadMsgNotification) SetErrorCode(v int32) { m.Data2.ErrorCode = v }
func (m *TLBadMsgNotification) GetErrorCode() int32  { return m.Data2.ErrorCode }

func (m *TLBadServerSalt) SetBadMsgId(v int64) { m.Data2.BadMsgId = v }
func (m *TLBadServerSalt) GetBadMsgId() int64  { return m.Data2.BadMsgId }

func (m *TLBadServerSalt) SetBadMsgSeqno(v int32) { m.Data2.BadMsgSeqno = v }
func (m *TLBadServerSalt) GetBadMsgSeqno() int32  { return m.Data2.BadMsgSeqno }

func (m *TLBadServerSalt) SetErrorCode(v int32) { m.Data2.ErrorCode = v }
func (m *TLBadServerSalt) GetErrorCode() int32  { return m.Data2.ErrorCode }

func (m *TLBadServerSalt) SetNewServerSalt(v int64) { m.Data2.NewServerSalt = v }
func (m *TLBadServerSalt) GetNewServerSalt() int64  { return m.Data2.NewServerSalt }

func (m *TLBankCardOpenUrl) SetUrl(v string) { m.Data2.Url = v }
func (m *TLBankCardOpenUrl) GetUrl() string  { return m.Data2.Url }

func (m *TLBankCardOpenUrl) SetName(v string) { m.Data2.Name = v }
func (m *TLBankCardOpenUrl) GetName() string  { return m.Data2.Name }

func (m *TLBindAuthKeyInner) SetNonce(v int64) { m.Data2.Nonce = v }
func (m *TLBindAuthKeyInner) GetNonce() int64  { return m.Data2.Nonce }

func (m *TLBindAuthKeyInner) SetTempAuthKeyId(v int64) { m.Data2.TempAuthKeyId = v }
func (m *TLBindAuthKeyInner) GetTempAuthKeyId() int64  { return m.Data2.TempAuthKeyId }

func (m *TLBindAuthKeyInner) SetPermAuthKeyId(v int64) { m.Data2.PermAuthKeyId = v }
func (m *TLBindAuthKeyInner) GetPermAuthKeyId() int64  { return m.Data2.PermAuthKeyId }

func (m *TLBindAuthKeyInner) SetTempSessionId(v int64) { m.Data2.TempSessionId = v }
func (m *TLBindAuthKeyInner) GetTempSessionId() int64  { return m.Data2.TempSessionId }

func (m *TLBindAuthKeyInner) SetExpiresAt(v int32) { m.Data2.ExpiresAt = v }
func (m *TLBindAuthKeyInner) GetExpiresAt() int32  { return m.Data2.ExpiresAt }

func (m *TLBotCommand) SetCommand(v string) { m.Data2.Command = v }
func (m *TLBotCommand) GetCommand() string  { return m.Data2.Command }

func (m *TLBotCommand) SetDescription(v string) { m.Data2.Description = v }
func (m *TLBotCommand) GetDescription() string  { return m.Data2.Description }

func (m *TLBotInfo) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLBotInfo) GetUserId() int32  { return m.Data2.UserId }

func (m *TLBotInfo) SetDescription(v string) { m.Data2.Description = v }
func (m *TLBotInfo) GetDescription() string  { return m.Data2.Description }

func (m *TLBotInfo) SetCommands(v []*BotCommand) { m.Data2.Commands = v }
func (m *TLBotInfo) GetCommands() []*BotCommand  { return m.Data2.Commands }

func (m *TLBotInlineMessageMediaAuto) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLBotInlineMessageMediaAuto) GetFlags() int32  { return m.Data2.Flags }

func (m *TLBotInlineMessageMediaAuto) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLBotInlineMessageMediaAuto) GetCaption() string  { return m.Data2.Caption }

func (m *TLBotInlineMessageMediaAuto) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLBotInlineMessageMediaAuto) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLBotInlineMessageMediaAuto) SetMessage(v string) { m.Data2.Message = v }
func (m *TLBotInlineMessageMediaAuto) GetMessage() string  { return m.Data2.Message }

func (m *TLBotInlineMessageMediaAuto) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLBotInlineMessageMediaAuto) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLBotInlineMessageText) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLBotInlineMessageText) GetFlags() int32  { return m.Data2.Flags }

func (m *TLBotInlineMessageText) SetNoWebpage(v bool) { m.Data2.NoWebpage = v }
func (m *TLBotInlineMessageText) GetNoWebpage() bool  { return m.Data2.NoWebpage }

func (m *TLBotInlineMessageText) SetMessage(v string) { m.Data2.Message = v }
func (m *TLBotInlineMessageText) GetMessage() string  { return m.Data2.Message }

func (m *TLBotInlineMessageText) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLBotInlineMessageText) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLBotInlineMessageText) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLBotInlineMessageText) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLBotInlineMessageMediaGeo) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLBotInlineMessageMediaGeo) GetFlags() int32  { return m.Data2.Flags }

func (m *TLBotInlineMessageMediaGeo) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLBotInlineMessageMediaGeo) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLBotInlineMessageMediaGeo) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLBotInlineMessageMediaGeo) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLBotInlineMessageMediaGeo) SetPeriod(v int32) { m.Data2.Period = v }
func (m *TLBotInlineMessageMediaGeo) GetPeriod() int32  { return m.Data2.Period }

func (m *TLBotInlineMessageMediaGeo) SetHeading(v int32) { m.Data2.Heading = v }
func (m *TLBotInlineMessageMediaGeo) GetHeading() int32  { return m.Data2.Heading }

func (m *TLBotInlineMessageMediaGeo) SetProximityNotificationRadius(v int32) {
	m.Data2.ProximityNotificationRadius = v
}
func (m *TLBotInlineMessageMediaGeo) GetProximityNotificationRadius() int32 {
	return m.Data2.ProximityNotificationRadius
}

func (m *TLBotInlineMessageMediaVenue) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLBotInlineMessageMediaVenue) GetFlags() int32  { return m.Data2.Flags }

func (m *TLBotInlineMessageMediaVenue) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLBotInlineMessageMediaVenue) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLBotInlineMessageMediaVenue) SetTitle(v string) { m.Data2.Title = v }
func (m *TLBotInlineMessageMediaVenue) GetTitle() string  { return m.Data2.Title }

func (m *TLBotInlineMessageMediaVenue) SetAddress(v string) { m.Data2.Address = v }
func (m *TLBotInlineMessageMediaVenue) GetAddress() string  { return m.Data2.Address }

func (m *TLBotInlineMessageMediaVenue) SetProvider(v string) { m.Data2.Provider = v }
func (m *TLBotInlineMessageMediaVenue) GetProvider() string  { return m.Data2.Provider }

func (m *TLBotInlineMessageMediaVenue) SetVenueId(v string) { m.Data2.VenueId = v }
func (m *TLBotInlineMessageMediaVenue) GetVenueId() string  { return m.Data2.VenueId }

func (m *TLBotInlineMessageMediaVenue) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLBotInlineMessageMediaVenue) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLBotInlineMessageMediaVenue) SetVenueType(v string) { m.Data2.VenueType = v }
func (m *TLBotInlineMessageMediaVenue) GetVenueType() string  { return m.Data2.VenueType }

func (m *TLBotInlineMessageMediaContact) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLBotInlineMessageMediaContact) GetFlags() int32  { return m.Data2.Flags }

func (m *TLBotInlineMessageMediaContact) SetPhoneNumber(v string) { m.Data2.PhoneNumber = v }
func (m *TLBotInlineMessageMediaContact) GetPhoneNumber() string  { return m.Data2.PhoneNumber }

func (m *TLBotInlineMessageMediaContact) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLBotInlineMessageMediaContact) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLBotInlineMessageMediaContact) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLBotInlineMessageMediaContact) GetLastName() string  { return m.Data2.LastName }

func (m *TLBotInlineMessageMediaContact) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLBotInlineMessageMediaContact) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLBotInlineMessageMediaContact) SetVcard(v string) { m.Data2.Vcard = v }
func (m *TLBotInlineMessageMediaContact) GetVcard() string  { return m.Data2.Vcard }

func (m *TLBotInlineResult) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLBotInlineResult) GetFlags() int32  { return m.Data2.Flags }

func (m *TLBotInlineResult) SetId(v string) { m.Data2.Id = v }
func (m *TLBotInlineResult) GetId() string  { return m.Data2.Id }

func (m *TLBotInlineResult) SetType(v string) { m.Data2.Type = v }
func (m *TLBotInlineResult) GetType() string  { return m.Data2.Type }

func (m *TLBotInlineResult) SetTitle(v string) { m.Data2.Title = v }
func (m *TLBotInlineResult) GetTitle() string  { return m.Data2.Title }

func (m *TLBotInlineResult) SetDescription(v string) { m.Data2.Description = v }
func (m *TLBotInlineResult) GetDescription() string  { return m.Data2.Description }

func (m *TLBotInlineResult) SetUrl(v string) { m.Data2.Url = v }
func (m *TLBotInlineResult) GetUrl() string  { return m.Data2.Url }

func (m *TLBotInlineResult) SetThumbUrl(v string) { m.Data2.ThumbUrl = v }
func (m *TLBotInlineResult) GetThumbUrl() string  { return m.Data2.ThumbUrl }

func (m *TLBotInlineResult) SetContentUrl(v string) { m.Data2.ContentUrl = v }
func (m *TLBotInlineResult) GetContentUrl() string  { return m.Data2.ContentUrl }

func (m *TLBotInlineResult) SetContentType(v string) { m.Data2.ContentType = v }
func (m *TLBotInlineResult) GetContentType() string  { return m.Data2.ContentType }

func (m *TLBotInlineResult) SetW(v int32) { m.Data2.W = v }
func (m *TLBotInlineResult) GetW() int32  { return m.Data2.W }

func (m *TLBotInlineResult) SetH(v int32) { m.Data2.H = v }
func (m *TLBotInlineResult) GetH() int32  { return m.Data2.H }

func (m *TLBotInlineResult) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLBotInlineResult) GetDuration() int32  { return m.Data2.Duration }

func (m *TLBotInlineResult) SetSendMessage(v *BotInlineMessage) { m.Data2.SendMessage = v }
func (m *TLBotInlineResult) GetSendMessage() *BotInlineMessage  { return m.Data2.SendMessage }

func (m *TLBotInlineResult) SetThumb(v *WebDocument) { m.Data2.Thumb = v }
func (m *TLBotInlineResult) GetThumb() *WebDocument  { return m.Data2.Thumb }

func (m *TLBotInlineResult) SetContent(v *WebDocument) { m.Data2.Content = v }
func (m *TLBotInlineResult) GetContent() *WebDocument  { return m.Data2.Content }

func (m *TLBotInlineMediaResult) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLBotInlineMediaResult) GetFlags() int32  { return m.Data2.Flags }

func (m *TLBotInlineMediaResult) SetId(v string) { m.Data2.Id = v }
func (m *TLBotInlineMediaResult) GetId() string  { return m.Data2.Id }

func (m *TLBotInlineMediaResult) SetType(v string) { m.Data2.Type = v }
func (m *TLBotInlineMediaResult) GetType() string  { return m.Data2.Type }

func (m *TLBotInlineMediaResult) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLBotInlineMediaResult) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLBotInlineMediaResult) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLBotInlineMediaResult) GetDocument() *Document  { return m.Data2.Document }

func (m *TLBotInlineMediaResult) SetTitle(v string) { m.Data2.Title = v }
func (m *TLBotInlineMediaResult) GetTitle() string  { return m.Data2.Title }

func (m *TLBotInlineMediaResult) SetDescription(v string) { m.Data2.Description = v }
func (m *TLBotInlineMediaResult) GetDescription() string  { return m.Data2.Description }

func (m *TLBotInlineMediaResult) SetSendMessage(v *BotInlineMessage) { m.Data2.SendMessage = v }
func (m *TLBotInlineMediaResult) GetSendMessage() *BotInlineMessage  { return m.Data2.SendMessage }

func (m *TLCdnConfig) SetPublicKeys(v []*CdnPublicKey) { m.Data2.PublicKeys = v }
func (m *TLCdnConfig) GetPublicKeys() []*CdnPublicKey  { return m.Data2.PublicKeys }

func (m *TLCdnFileHash) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLCdnFileHash) GetOffset() int32  { return m.Data2.Offset }

func (m *TLCdnFileHash) SetLimit(v int32) { m.Data2.Limit = v }
func (m *TLCdnFileHash) GetLimit() int32  { return m.Data2.Limit }

func (m *TLCdnFileHash) SetHash(v []byte) { m.Data2.Hash = v }
func (m *TLCdnFileHash) GetHash() []byte  { return m.Data2.Hash }

func (m *TLCdnPublicKey) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLCdnPublicKey) GetDcId() int32  { return m.Data2.DcId }

func (m *TLCdnPublicKey) SetPublicKey(v string) { m.Data2.PublicKey = v }
func (m *TLCdnPublicKey) GetPublicKey() string  { return m.Data2.PublicKey }

func (m *TLChannelAdminLogEvent) SetId(v int64) { m.Data2.Id = v }
func (m *TLChannelAdminLogEvent) GetId() int64  { return m.Data2.Id }

func (m *TLChannelAdminLogEvent) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelAdminLogEvent) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelAdminLogEvent) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelAdminLogEvent) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelAdminLogEvent) SetAction(v *ChannelAdminLogEventAction) { m.Data2.Action = v }
func (m *TLChannelAdminLogEvent) GetAction() *ChannelAdminLogEventAction  { return m.Data2.Action }

func (m *TLChannelAdminLogEventActionChangeTitle) SetPrevValueE6DFB82571(v string) {
	m.Data2.PrevValueE6DFB82571 = v
}
func (m *TLChannelAdminLogEventActionChangeTitle) GetPrevValueE6DFB82571() string {
	return m.Data2.PrevValueE6DFB82571
}

func (m *TLChannelAdminLogEventActionChangeTitle) SetNewValueE6DFB82571(v string) {
	m.Data2.NewValueE6DFB82571 = v
}
func (m *TLChannelAdminLogEventActionChangeTitle) GetNewValueE6DFB82571() string {
	return m.Data2.NewValueE6DFB82571
}

func (m *TLChannelAdminLogEventActionChangeAbout) SetPrevValueE6DFB82571(v string) {
	m.Data2.PrevValueE6DFB82571 = v
}
func (m *TLChannelAdminLogEventActionChangeAbout) GetPrevValueE6DFB82571() string {
	return m.Data2.PrevValueE6DFB82571
}

func (m *TLChannelAdminLogEventActionChangeAbout) SetNewValueE6DFB82571(v string) {
	m.Data2.NewValueE6DFB82571 = v
}
func (m *TLChannelAdminLogEventActionChangeAbout) GetNewValueE6DFB82571() string {
	return m.Data2.NewValueE6DFB82571
}

func (m *TLChannelAdminLogEventActionChangeUsername) SetPrevValueE6DFB82571(v string) {
	m.Data2.PrevValueE6DFB82571 = v
}
func (m *TLChannelAdminLogEventActionChangeUsername) GetPrevValueE6DFB82571() string {
	return m.Data2.PrevValueE6DFB82571
}

func (m *TLChannelAdminLogEventActionChangeUsername) SetNewValueE6DFB82571(v string) {
	m.Data2.NewValueE6DFB82571 = v
}
func (m *TLChannelAdminLogEventActionChangeUsername) GetNewValueE6DFB82571() string {
	return m.Data2.NewValueE6DFB82571
}

func (m *TLChannelAdminLogEventActionChangePhoto) SetPrevPhotoB82F55C371(v *ChatPhoto) {
	m.Data2.PrevPhotoB82F55C371 = v
}
func (m *TLChannelAdminLogEventActionChangePhoto) GetPrevPhotoB82F55C371() *ChatPhoto {
	return m.Data2.PrevPhotoB82F55C371
}

func (m *TLChannelAdminLogEventActionChangePhoto) SetNewPhotoB82F55C371(v *ChatPhoto) {
	m.Data2.NewPhotoB82F55C371 = v
}
func (m *TLChannelAdminLogEventActionChangePhoto) GetNewPhotoB82F55C371() *ChatPhoto {
	return m.Data2.NewPhotoB82F55C371
}

func (m *TLChannelAdminLogEventActionChangePhoto) SetPrevPhoto434BD2AF98(v *Photo) {
	m.Data2.PrevPhoto434BD2AF98 = v
}
func (m *TLChannelAdminLogEventActionChangePhoto) GetPrevPhoto434BD2AF98() *Photo {
	return m.Data2.PrevPhoto434BD2AF98
}

func (m *TLChannelAdminLogEventActionChangePhoto) SetNewPhoto434BD2AF98(v *Photo) {
	m.Data2.NewPhoto434BD2AF98 = v
}
func (m *TLChannelAdminLogEventActionChangePhoto) GetNewPhoto434BD2AF98() *Photo {
	return m.Data2.NewPhoto434BD2AF98
}

func (m *TLChannelAdminLogEventActionToggleInvites) SetNewValue1B7907AE71(v *Bool) {
	m.Data2.NewValue1B7907AE71 = v
}
func (m *TLChannelAdminLogEventActionToggleInvites) GetNewValue1B7907AE71() *Bool {
	return m.Data2.NewValue1B7907AE71
}

func (m *TLChannelAdminLogEventActionToggleSignatures) SetNewValue1B7907AE71(v *Bool) {
	m.Data2.NewValue1B7907AE71 = v
}
func (m *TLChannelAdminLogEventActionToggleSignatures) GetNewValue1B7907AE71() *Bool {
	return m.Data2.NewValue1B7907AE71
}

func (m *TLChannelAdminLogEventActionUpdatePinned) SetMessage(v *Message) { m.Data2.Message = v }
func (m *TLChannelAdminLogEventActionUpdatePinned) GetMessage() *Message  { return m.Data2.Message }

func (m *TLChannelAdminLogEventActionEditMessage) SetPrevMessage(v *Message) { m.Data2.PrevMessage = v }
func (m *TLChannelAdminLogEventActionEditMessage) GetPrevMessage() *Message {
	return m.Data2.PrevMessage
}

func (m *TLChannelAdminLogEventActionEditMessage) SetNewMessage(v *Message) { m.Data2.NewMessage = v }
func (m *TLChannelAdminLogEventActionEditMessage) GetNewMessage() *Message  { return m.Data2.NewMessage }

func (m *TLChannelAdminLogEventActionDeleteMessage) SetMessage(v *Message) { m.Data2.Message = v }
func (m *TLChannelAdminLogEventActionDeleteMessage) GetMessage() *Message  { return m.Data2.Message }

func (m *TLChannelAdminLogEventActionParticipantInvite) SetParticipantE31C34D871(v *ChannelParticipant) {
	m.Data2.ParticipantE31C34D871 = v
}
func (m *TLChannelAdminLogEventActionParticipantInvite) GetParticipantE31C34D871() *ChannelParticipant {
	return m.Data2.ParticipantE31C34D871
}

func (m *TLChannelAdminLogEventActionParticipantToggleBan) SetPrevParticipant(v *ChannelParticipant) {
	m.Data2.PrevParticipant = v
}
func (m *TLChannelAdminLogEventActionParticipantToggleBan) GetPrevParticipant() *ChannelParticipant {
	return m.Data2.PrevParticipant
}

func (m *TLChannelAdminLogEventActionParticipantToggleBan) SetNewParticipant(v *ChannelParticipant) {
	m.Data2.NewParticipant = v
}
func (m *TLChannelAdminLogEventActionParticipantToggleBan) GetNewParticipant() *ChannelParticipant {
	return m.Data2.NewParticipant
}

func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) SetPrevParticipant(v *ChannelParticipant) {
	m.Data2.PrevParticipant = v
}
func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) GetPrevParticipant() *ChannelParticipant {
	return m.Data2.PrevParticipant
}

func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) SetNewParticipant(v *ChannelParticipant) {
	m.Data2.NewParticipant = v
}
func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) GetNewParticipant() *ChannelParticipant {
	return m.Data2.NewParticipant
}

func (m *TLChannelAdminLogEventActionChangeStickerSet) SetPrevStickerset(v *InputStickerSet) {
	m.Data2.PrevStickerset = v
}
func (m *TLChannelAdminLogEventActionChangeStickerSet) GetPrevStickerset() *InputStickerSet {
	return m.Data2.PrevStickerset
}

func (m *TLChannelAdminLogEventActionChangeStickerSet) SetNewStickerset(v *InputStickerSet) {
	m.Data2.NewStickerset = v
}
func (m *TLChannelAdminLogEventActionChangeStickerSet) GetNewStickerset() *InputStickerSet {
	return m.Data2.NewStickerset
}

func (m *TLChannelAdminLogEventActionTogglePreHistoryHidden) SetNewValue1B7907AE71(v *Bool) {
	m.Data2.NewValue1B7907AE71 = v
}
func (m *TLChannelAdminLogEventActionTogglePreHistoryHidden) GetNewValue1B7907AE71() *Bool {
	return m.Data2.NewValue1B7907AE71
}

func (m *TLChannelAdminLogEventActionDefaultBannedRights) SetPrevBannedRights(v *ChatBannedRights) {
	m.Data2.PrevBannedRights = v
}
func (m *TLChannelAdminLogEventActionDefaultBannedRights) GetPrevBannedRights() *ChatBannedRights {
	return m.Data2.PrevBannedRights
}

func (m *TLChannelAdminLogEventActionDefaultBannedRights) SetNewBannedRights(v *ChatBannedRights) {
	m.Data2.NewBannedRights = v
}
func (m *TLChannelAdminLogEventActionDefaultBannedRights) GetNewBannedRights() *ChatBannedRights {
	return m.Data2.NewBannedRights
}

func (m *TLChannelAdminLogEventActionStopPoll) SetMessage(v *Message) { m.Data2.Message = v }
func (m *TLChannelAdminLogEventActionStopPoll) GetMessage() *Message  { return m.Data2.Message }

func (m *TLChannelAdminLogEventActionChangeLinkedChat) SetPrevValueA26F881B100(v int32) {
	m.Data2.PrevValueA26F881B100 = v
}
func (m *TLChannelAdminLogEventActionChangeLinkedChat) GetPrevValueA26F881B100() int32 {
	return m.Data2.PrevValueA26F881B100
}

func (m *TLChannelAdminLogEventActionChangeLinkedChat) SetNewValueA26F881B100(v int32) {
	m.Data2.NewValueA26F881B100 = v
}
func (m *TLChannelAdminLogEventActionChangeLinkedChat) GetNewValueA26F881B100() int32 {
	return m.Data2.NewValueA26F881B100
}

func (m *TLChannelAdminLogEventActionChangeLocation) SetPrevValueE6B76AE102(v *ChannelLocation) {
	m.Data2.PrevValueE6B76AE102 = v
}
func (m *TLChannelAdminLogEventActionChangeLocation) GetPrevValueE6B76AE102() *ChannelLocation {
	return m.Data2.PrevValueE6B76AE102
}

func (m *TLChannelAdminLogEventActionChangeLocation) SetNewValueE6B76AE102(v *ChannelLocation) {
	m.Data2.NewValueE6B76AE102 = v
}
func (m *TLChannelAdminLogEventActionChangeLocation) GetNewValueE6B76AE102() *ChannelLocation {
	return m.Data2.NewValueE6B76AE102
}

func (m *TLChannelAdminLogEventActionToggleSlowMode) SetPrevValueA26F881B100(v int32) {
	m.Data2.PrevValueA26F881B100 = v
}
func (m *TLChannelAdminLogEventActionToggleSlowMode) GetPrevValueA26F881B100() int32 {
	return m.Data2.PrevValueA26F881B100
}

func (m *TLChannelAdminLogEventActionToggleSlowMode) SetNewValueA26F881B100(v int32) {
	m.Data2.NewValueA26F881B100 = v
}
func (m *TLChannelAdminLogEventActionToggleSlowMode) GetNewValueA26F881B100() int32 {
	return m.Data2.NewValueA26F881B100
}

func (m *TLChannelAdminLogEventActionStartGroupCall) SetCall(v *InputGroupCall) { m.Data2.Call = v }
func (m *TLChannelAdminLogEventActionStartGroupCall) GetCall() *InputGroupCall  { return m.Data2.Call }

func (m *TLChannelAdminLogEventActionDiscardGroupCall) SetCall(v *InputGroupCall) { m.Data2.Call = v }
func (m *TLChannelAdminLogEventActionDiscardGroupCall) GetCall() *InputGroupCall  { return m.Data2.Call }

func (m *TLChannelAdminLogEventActionParticipantMute) SetParticipantF92424D2122(v *GroupCallParticipant) {
	m.Data2.ParticipantF92424D2122 = v
}
func (m *TLChannelAdminLogEventActionParticipantMute) GetParticipantF92424D2122() *GroupCallParticipant {
	return m.Data2.ParticipantF92424D2122
}

func (m *TLChannelAdminLogEventActionParticipantUnmute) SetParticipantF92424D2122(v *GroupCallParticipant) {
	m.Data2.ParticipantF92424D2122 = v
}
func (m *TLChannelAdminLogEventActionParticipantUnmute) GetParticipantF92424D2122() *GroupCallParticipant {
	return m.Data2.ParticipantF92424D2122
}

func (m *TLChannelAdminLogEventActionToggleGroupCallSetting) SetJoinMuted(v *Bool) {
	m.Data2.JoinMuted = v
}
func (m *TLChannelAdminLogEventActionToggleGroupCallSetting) GetJoinMuted() *Bool {
	return m.Data2.JoinMuted
}

func (m *TLChannelAdminLogEventsFilter) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChannelAdminLogEventsFilter) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChannelAdminLogEventsFilter) SetJoin(v bool) { m.Data2.Join = v }
func (m *TLChannelAdminLogEventsFilter) GetJoin() bool  { return m.Data2.Join }

func (m *TLChannelAdminLogEventsFilter) SetLeave(v bool) { m.Data2.Leave = v }
func (m *TLChannelAdminLogEventsFilter) GetLeave() bool  { return m.Data2.Leave }

func (m *TLChannelAdminLogEventsFilter) SetInvite(v bool) { m.Data2.Invite = v }
func (m *TLChannelAdminLogEventsFilter) GetInvite() bool  { return m.Data2.Invite }

func (m *TLChannelAdminLogEventsFilter) SetBan(v bool) { m.Data2.Ban = v }
func (m *TLChannelAdminLogEventsFilter) GetBan() bool  { return m.Data2.Ban }

func (m *TLChannelAdminLogEventsFilter) SetUnban(v bool) { m.Data2.Unban = v }
func (m *TLChannelAdminLogEventsFilter) GetUnban() bool  { return m.Data2.Unban }

func (m *TLChannelAdminLogEventsFilter) SetKick(v bool) { m.Data2.Kick = v }
func (m *TLChannelAdminLogEventsFilter) GetKick() bool  { return m.Data2.Kick }

func (m *TLChannelAdminLogEventsFilter) SetUnkick(v bool) { m.Data2.Unkick = v }
func (m *TLChannelAdminLogEventsFilter) GetUnkick() bool  { return m.Data2.Unkick }

func (m *TLChannelAdminLogEventsFilter) SetPromote(v bool) { m.Data2.Promote = v }
func (m *TLChannelAdminLogEventsFilter) GetPromote() bool  { return m.Data2.Promote }

func (m *TLChannelAdminLogEventsFilter) SetDemote(v bool) { m.Data2.Demote = v }
func (m *TLChannelAdminLogEventsFilter) GetDemote() bool  { return m.Data2.Demote }

func (m *TLChannelAdminLogEventsFilter) SetInfo(v bool) { m.Data2.Info = v }
func (m *TLChannelAdminLogEventsFilter) GetInfo() bool  { return m.Data2.Info }

func (m *TLChannelAdminLogEventsFilter) SetSettings(v bool) { m.Data2.Settings = v }
func (m *TLChannelAdminLogEventsFilter) GetSettings() bool  { return m.Data2.Settings }

func (m *TLChannelAdminLogEventsFilter) SetPinned(v bool) { m.Data2.Pinned = v }
func (m *TLChannelAdminLogEventsFilter) GetPinned() bool  { return m.Data2.Pinned }

func (m *TLChannelAdminLogEventsFilter) SetEdit(v bool) { m.Data2.Edit = v }
func (m *TLChannelAdminLogEventsFilter) GetEdit() bool  { return m.Data2.Edit }

func (m *TLChannelAdminLogEventsFilter) SetDelete(v bool) { m.Data2.Delete = v }
func (m *TLChannelAdminLogEventsFilter) GetDelete() bool  { return m.Data2.Delete }

func (m *TLChannelAdminLogEventsFilter) SetGroupCall(v bool) { m.Data2.GroupCall = v }
func (m *TLChannelAdminLogEventsFilter) GetGroupCall() bool  { return m.Data2.GroupCall }

func (m *TLChannelAdminRights) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChannelAdminRights) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChannelAdminRights) SetChangeInfo(v bool) { m.Data2.ChangeInfo = v }
func (m *TLChannelAdminRights) GetChangeInfo() bool  { return m.Data2.ChangeInfo }

func (m *TLChannelAdminRights) SetPostMessages(v bool) { m.Data2.PostMessages = v }
func (m *TLChannelAdminRights) GetPostMessages() bool  { return m.Data2.PostMessages }

func (m *TLChannelAdminRights) SetEditMessages(v bool) { m.Data2.EditMessages = v }
func (m *TLChannelAdminRights) GetEditMessages() bool  { return m.Data2.EditMessages }

func (m *TLChannelAdminRights) SetDeleteMessages(v bool) { m.Data2.DeleteMessages = v }
func (m *TLChannelAdminRights) GetDeleteMessages() bool  { return m.Data2.DeleteMessages }

func (m *TLChannelAdminRights) SetBanUsers(v bool) { m.Data2.BanUsers = v }
func (m *TLChannelAdminRights) GetBanUsers() bool  { return m.Data2.BanUsers }

func (m *TLChannelAdminRights) SetInviteUsers(v bool) { m.Data2.InviteUsers = v }
func (m *TLChannelAdminRights) GetInviteUsers() bool  { return m.Data2.InviteUsers }

func (m *TLChannelAdminRights) SetInviteLink(v bool) { m.Data2.InviteLink = v }
func (m *TLChannelAdminRights) GetInviteLink() bool  { return m.Data2.InviteLink }

func (m *TLChannelAdminRights) SetPinMessages(v bool) { m.Data2.PinMessages = v }
func (m *TLChannelAdminRights) GetPinMessages() bool  { return m.Data2.PinMessages }

func (m *TLChannelAdminRights) SetAddAdmins(v bool) { m.Data2.AddAdmins = v }
func (m *TLChannelAdminRights) GetAddAdmins() bool  { return m.Data2.AddAdmins }

func (m *TLChannelBannedRights) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChannelBannedRights) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChannelBannedRights) SetViewMessages(v bool) { m.Data2.ViewMessages = v }
func (m *TLChannelBannedRights) GetViewMessages() bool  { return m.Data2.ViewMessages }

func (m *TLChannelBannedRights) SetSendMessages(v bool) { m.Data2.SendMessages = v }
func (m *TLChannelBannedRights) GetSendMessages() bool  { return m.Data2.SendMessages }

func (m *TLChannelBannedRights) SetSendMedia(v bool) { m.Data2.SendMedia = v }
func (m *TLChannelBannedRights) GetSendMedia() bool  { return m.Data2.SendMedia }

func (m *TLChannelBannedRights) SetSendStickers(v bool) { m.Data2.SendStickers = v }
func (m *TLChannelBannedRights) GetSendStickers() bool  { return m.Data2.SendStickers }

func (m *TLChannelBannedRights) SetSendGifs(v bool) { m.Data2.SendGifs = v }
func (m *TLChannelBannedRights) GetSendGifs() bool  { return m.Data2.SendGifs }

func (m *TLChannelBannedRights) SetSendGames(v bool) { m.Data2.SendGames = v }
func (m *TLChannelBannedRights) GetSendGames() bool  { return m.Data2.SendGames }

func (m *TLChannelBannedRights) SetSendInline(v bool) { m.Data2.SendInline = v }
func (m *TLChannelBannedRights) GetSendInline() bool  { return m.Data2.SendInline }

func (m *TLChannelBannedRights) SetEmbedLinks(v bool) { m.Data2.EmbedLinks = v }
func (m *TLChannelBannedRights) GetEmbedLinks() bool  { return m.Data2.EmbedLinks }

func (m *TLChannelBannedRights) SetUntilDate(v int32) { m.Data2.UntilDate = v }
func (m *TLChannelBannedRights) GetUntilDate() int32  { return m.Data2.UntilDate }

func (m *TLChannelLocation) SetGeoPoint(v *GeoPoint) { m.Data2.GeoPoint = v }
func (m *TLChannelLocation) GetGeoPoint() *GeoPoint  { return m.Data2.GeoPoint }

func (m *TLChannelLocation) SetAddress(v string) { m.Data2.Address = v }
func (m *TLChannelLocation) GetAddress() string  { return m.Data2.Address }

func (m *TLChannelMessagesFilter) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChannelMessagesFilter) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChannelMessagesFilter) SetExcludeNewMessages(v bool) { m.Data2.ExcludeNewMessages = v }
func (m *TLChannelMessagesFilter) GetExcludeNewMessages() bool  { return m.Data2.ExcludeNewMessages }

func (m *TLChannelMessagesFilter) SetRanges(v []*MessageRange) { m.Data2.Ranges = v }
func (m *TLChannelMessagesFilter) GetRanges() []*MessageRange  { return m.Data2.Ranges }

func (m *TLChannelParticipant) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipant) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipant) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelParticipant) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelParticipantSelf) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantSelf) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantSelf) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLChannelParticipantSelf) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLChannelParticipantSelf) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelParticipantSelf) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelParticipantCreator) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantCreator) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantCreator) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChannelParticipantCreator) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChannelParticipantCreator) SetRank(v string) { m.Data2.Rank = v }
func (m *TLChannelParticipantCreator) GetRank() string  { return m.Data2.Rank }

func (m *TLChannelParticipantCreator) SetAdminRights5DAA6E2393(v *ChatAdminRights) {
	m.Data2.AdminRights5DAA6E2393 = v
}
func (m *TLChannelParticipantCreator) GetAdminRights5DAA6E2393() *ChatAdminRights {
	return m.Data2.AdminRights5DAA6E2393
}

func (m *TLChannelParticipantAdmin) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChannelParticipantAdmin) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChannelParticipantAdmin) SetCanEdit(v bool) { m.Data2.CanEdit = v }
func (m *TLChannelParticipantAdmin) GetCanEdit() bool  { return m.Data2.CanEdit }

func (m *TLChannelParticipantAdmin) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantAdmin) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantAdmin) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLChannelParticipantAdmin) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLChannelParticipantAdmin) SetPromotedBy(v int32) { m.Data2.PromotedBy = v }
func (m *TLChannelParticipantAdmin) GetPromotedBy() int32  { return m.Data2.PromotedBy }

func (m *TLChannelParticipantAdmin) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelParticipantAdmin) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelParticipantAdmin) SetAdminRightsA82FA89871(v *ChannelAdminRights) {
	m.Data2.AdminRightsA82FA89871 = v
}
func (m *TLChannelParticipantAdmin) GetAdminRightsA82FA89871() *ChannelAdminRights {
	return m.Data2.AdminRightsA82FA89871
}

func (m *TLChannelParticipantAdmin) SetSelf(v bool) { m.Data2.Self = v }
func (m *TLChannelParticipantAdmin) GetSelf() bool  { return m.Data2.Self }

func (m *TLChannelParticipantAdmin) SetAdminRights5DAA6E2393(v *ChatAdminRights) {
	m.Data2.AdminRights5DAA6E2393 = v
}
func (m *TLChannelParticipantAdmin) GetAdminRights5DAA6E2393() *ChatAdminRights {
	return m.Data2.AdminRights5DAA6E2393
}

func (m *TLChannelParticipantAdmin) SetRank(v string) { m.Data2.Rank = v }
func (m *TLChannelParticipantAdmin) GetRank() string  { return m.Data2.Rank }

func (m *TLChannelParticipantBanned) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChannelParticipantBanned) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChannelParticipantBanned) SetLeft(v bool) { m.Data2.Left = v }
func (m *TLChannelParticipantBanned) GetLeft() bool  { return m.Data2.Left }

func (m *TLChannelParticipantBanned) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantBanned) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantBanned) SetKickedBy(v int32) { m.Data2.KickedBy = v }
func (m *TLChannelParticipantBanned) GetKickedBy() int32  { return m.Data2.KickedBy }

func (m *TLChannelParticipantBanned) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelParticipantBanned) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelParticipantBanned) SetBannedRights222C188671(v *ChannelBannedRights) {
	m.Data2.BannedRights222C188671 = v
}
func (m *TLChannelParticipantBanned) GetBannedRights222C188671() *ChannelBannedRights {
	return m.Data2.BannedRights222C188671
}

func (m *TLChannelParticipantBanned) SetBannedRights1C0FACAF93(v *ChatBannedRights) {
	m.Data2.BannedRights1C0FACAF93 = v
}
func (m *TLChannelParticipantBanned) GetBannedRights1C0FACAF93() *ChatBannedRights {
	return m.Data2.BannedRights1C0FACAF93
}

func (m *TLChannelParticipantModerator) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantModerator) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantModerator) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLChannelParticipantModerator) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLChannelParticipantModerator) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelParticipantModerator) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelParticipantEditor) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantEditor) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantEditor) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLChannelParticipantEditor) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLChannelParticipantEditor) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelParticipantEditor) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelParticipantKicked) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantKicked) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantKicked) SetKickedBy(v int32) { m.Data2.KickedBy = v }
func (m *TLChannelParticipantKicked) GetKickedBy() int32  { return m.Data2.KickedBy }

func (m *TLChannelParticipantKicked) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelParticipantKicked) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelParticipantLeft) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantLeft) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantsKicked) SetQ(v string) { m.Data2.Q = v }
func (m *TLChannelParticipantsKicked) GetQ() string  { return m.Data2.Q }

func (m *TLChannelParticipantsBanned) SetQ(v string) { m.Data2.Q = v }
func (m *TLChannelParticipantsBanned) GetQ() string  { return m.Data2.Q }

func (m *TLChannelParticipantsSearch) SetQ(v string) { m.Data2.Q = v }
func (m *TLChannelParticipantsSearch) GetQ() string  { return m.Data2.Q }

func (m *TLChannelParticipantsContacts) SetQ(v string) { m.Data2.Q = v }
func (m *TLChannelParticipantsContacts) GetQ() string  { return m.Data2.Q }

func (m *TLChannelParticipantsMentions) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChannelParticipantsMentions) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChannelParticipantsMentions) SetQ(v string) { m.Data2.Q = v }
func (m *TLChannelParticipantsMentions) GetQ() string  { return m.Data2.Q }

func (m *TLChannelParticipantsMentions) SetTopMsgId(v int32) { m.Data2.TopMsgId = v }
func (m *TLChannelParticipantsMentions) GetTopMsgId() int32  { return m.Data2.TopMsgId }

func (m *TLChannelsAdminLogResults) SetEvents(v []*ChannelAdminLogEvent) { m.Data2.Events = v }
func (m *TLChannelsAdminLogResults) GetEvents() []*ChannelAdminLogEvent  { return m.Data2.Events }

func (m *TLChannelsAdminLogResults) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLChannelsAdminLogResults) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLChannelsAdminLogResults) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLChannelsAdminLogResults) GetUsers() []*User  { return m.Data2.Users }

func (m *TLChannelsChannelParticipant) SetParticipant(v *ChannelParticipant) { m.Data2.Participant = v }
func (m *TLChannelsChannelParticipant) GetParticipant() *ChannelParticipant {
	return m.Data2.Participant
}

func (m *TLChannelsChannelParticipant) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLChannelsChannelParticipant) GetUsers() []*User  { return m.Data2.Users }

func (m *TLChannelsChannelParticipants) SetCount(v int32) { m.Data2.Count = v }
func (m *TLChannelsChannelParticipants) GetCount() int32  { return m.Data2.Count }

func (m *TLChannelsChannelParticipants) SetParticipants(v []*ChannelParticipant) {
	m.Data2.Participants = v
}
func (m *TLChannelsChannelParticipants) GetParticipants() []*ChannelParticipant {
	return m.Data2.Participants
}

func (m *TLChannelsChannelParticipants) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLChannelsChannelParticipants) GetUsers() []*User  { return m.Data2.Users }

func (m *TLChatEmpty) SetId(v int32) { m.Data2.Id = v }
func (m *TLChatEmpty) GetId() int32  { return m.Data2.Id }

func (m *TLChat) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChat) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChat) SetCreator(v bool) { m.Data2.Creator = v }
func (m *TLChat) GetCreator() bool  { return m.Data2.Creator }

func (m *TLChat) SetKicked(v bool) { m.Data2.Kicked = v }
func (m *TLChat) GetKicked() bool  { return m.Data2.Kicked }

func (m *TLChat) SetLeft(v bool) { m.Data2.Left = v }
func (m *TLChat) GetLeft() bool  { return m.Data2.Left }

func (m *TLChat) SetAdminsEnabled(v bool) { m.Data2.AdminsEnabled = v }
func (m *TLChat) GetAdminsEnabled() bool  { return m.Data2.AdminsEnabled }

func (m *TLChat) SetAdmin(v bool) { m.Data2.Admin = v }
func (m *TLChat) GetAdmin() bool  { return m.Data2.Admin }

func (m *TLChat) SetDeactivated(v bool) { m.Data2.Deactivated = v }
func (m *TLChat) GetDeactivated() bool  { return m.Data2.Deactivated }

func (m *TLChat) SetId(v int32) { m.Data2.Id = v }
func (m *TLChat) GetId() int32  { return m.Data2.Id }

func (m *TLChat) SetTitle(v string) { m.Data2.Title = v }
func (m *TLChat) GetTitle() string  { return m.Data2.Title }

func (m *TLChat) SetPhoto(v *ChatPhoto) { m.Data2.Photo = v }
func (m *TLChat) GetPhoto() *ChatPhoto  { return m.Data2.Photo }

func (m *TLChat) SetParticipantsCount(v int32) { m.Data2.ParticipantsCount = v }
func (m *TLChat) GetParticipantsCount() int32  { return m.Data2.ParticipantsCount }

func (m *TLChat) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChat) GetDate() int32  { return m.Data2.Date }

func (m *TLChat) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLChat) GetVersion() int32  { return m.Data2.Version }

func (m *TLChat) SetMigratedTo(v *InputChannel) { m.Data2.MigratedTo = v }
func (m *TLChat) GetMigratedTo() *InputChannel  { return m.Data2.MigratedTo }

func (m *TLChat) SetCallActive(v bool) { m.Data2.CallActive = v }
func (m *TLChat) GetCallActive() bool  { return m.Data2.CallActive }

func (m *TLChat) SetCallNotEmpty(v bool) { m.Data2.CallNotEmpty = v }
func (m *TLChat) GetCallNotEmpty() bool  { return m.Data2.CallNotEmpty }

func (m *TLChat) SetAdminRights4DF3083493(v *ChatAdminRights) { m.Data2.AdminRights4DF3083493 = v }
func (m *TLChat) GetAdminRights4DF3083493() *ChatAdminRights  { return m.Data2.AdminRights4DF3083493 }

func (m *TLChat) SetDefaultBannedRights(v *ChatBannedRights) { m.Data2.DefaultBannedRights = v }
func (m *TLChat) GetDefaultBannedRights() *ChatBannedRights  { return m.Data2.DefaultBannedRights }

func (m *TLChatForbidden) SetId(v int32) { m.Data2.Id = v }
func (m *TLChatForbidden) GetId() int32  { return m.Data2.Id }

func (m *TLChatForbidden) SetTitle(v string) { m.Data2.Title = v }
func (m *TLChatForbidden) GetTitle() string  { return m.Data2.Title }

func (m *TLChannel) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChannel) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChannel) SetCreator(v bool) { m.Data2.Creator = v }
func (m *TLChannel) GetCreator() bool  { return m.Data2.Creator }

func (m *TLChannel) SetLeft(v bool) { m.Data2.Left = v }
func (m *TLChannel) GetLeft() bool  { return m.Data2.Left }

func (m *TLChannel) SetEditor(v bool) { m.Data2.Editor = v }
func (m *TLChannel) GetEditor() bool  { return m.Data2.Editor }

func (m *TLChannel) SetBroadcast(v bool) { m.Data2.Broadcast = v }
func (m *TLChannel) GetBroadcast() bool  { return m.Data2.Broadcast }

func (m *TLChannel) SetVerified(v bool) { m.Data2.Verified = v }
func (m *TLChannel) GetVerified() bool  { return m.Data2.Verified }

func (m *TLChannel) SetMegagroup(v bool) { m.Data2.Megagroup = v }
func (m *TLChannel) GetMegagroup() bool  { return m.Data2.Megagroup }

func (m *TLChannel) SetRestricted(v bool) { m.Data2.Restricted = v }
func (m *TLChannel) GetRestricted() bool  { return m.Data2.Restricted }

func (m *TLChannel) SetDemocracy(v bool) { m.Data2.Democracy = v }
func (m *TLChannel) GetDemocracy() bool  { return m.Data2.Democracy }

func (m *TLChannel) SetSignatures(v bool) { m.Data2.Signatures = v }
func (m *TLChannel) GetSignatures() bool  { return m.Data2.Signatures }

func (m *TLChannel) SetMin(v bool) { m.Data2.Min = v }
func (m *TLChannel) GetMin() bool  { return m.Data2.Min }

func (m *TLChannel) SetId(v int32) { m.Data2.Id = v }
func (m *TLChannel) GetId() int32  { return m.Data2.Id }

func (m *TLChannel) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLChannel) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLChannel) SetTitle(v string) { m.Data2.Title = v }
func (m *TLChannel) GetTitle() string  { return m.Data2.Title }

func (m *TLChannel) SetUsername(v string) { m.Data2.Username = v }
func (m *TLChannel) GetUsername() string  { return m.Data2.Username }

func (m *TLChannel) SetPhoto(v *ChatPhoto) { m.Data2.Photo = v }
func (m *TLChannel) GetPhoto() *ChatPhoto  { return m.Data2.Photo }

func (m *TLChannel) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannel) GetDate() int32  { return m.Data2.Date }

func (m *TLChannel) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLChannel) GetVersion() int32  { return m.Data2.Version }

func (m *TLChannel) SetRestrictionReasonCB44B1C71(v string) { m.Data2.RestrictionReasonCB44B1C71 = v }
func (m *TLChannel) GetRestrictionReasonCB44B1C71() string  { return m.Data2.RestrictionReasonCB44B1C71 }

func (m *TLChannel) SetAdminRightsCB44B1C71(v *ChannelAdminRights) { m.Data2.AdminRightsCB44B1C71 = v }
func (m *TLChannel) GetAdminRightsCB44B1C71() *ChannelAdminRights {
	return m.Data2.AdminRightsCB44B1C71
}

func (m *TLChannel) SetBannedRightsCB44B1C71(v *ChannelBannedRights) {
	m.Data2.BannedRightsCB44B1C71 = v
}
func (m *TLChannel) GetBannedRightsCB44B1C71() *ChannelBannedRights {
	return m.Data2.BannedRightsCB44B1C71
}

func (m *TLChannel) SetKicked(v bool) { m.Data2.Kicked = v }
func (m *TLChannel) GetKicked() bool  { return m.Data2.Kicked }

func (m *TLChannel) SetModerator(v bool) { m.Data2.Moderator = v }
func (m *TLChannel) GetModerator() bool  { return m.Data2.Moderator }

func (m *TLChannel) SetParticipantsCount(v int32) { m.Data2.ParticipantsCount = v }
func (m *TLChannel) GetParticipantsCount() int32  { return m.Data2.ParticipantsCount }

func (m *TLChannel) SetAdminRights4DF3083493(v *ChatAdminRights) { m.Data2.AdminRights4DF3083493 = v }
func (m *TLChannel) GetAdminRights4DF3083493() *ChatAdminRights  { return m.Data2.AdminRights4DF3083493 }

func (m *TLChannel) SetBannedRights4DF3083493(v *ChatBannedRights) {
	m.Data2.BannedRights4DF3083493 = v
}
func (m *TLChannel) GetBannedRights4DF3083493() *ChatBannedRights {
	return m.Data2.BannedRights4DF3083493
}

func (m *TLChannel) SetDefaultBannedRights(v *ChatBannedRights) { m.Data2.DefaultBannedRights = v }
func (m *TLChannel) GetDefaultBannedRights() *ChatBannedRights  { return m.Data2.DefaultBannedRights }

func (m *TLChannel) SetScam(v bool) { m.Data2.Scam = v }
func (m *TLChannel) GetScam() bool  { return m.Data2.Scam }

func (m *TLChannel) SetHasLink(v bool) { m.Data2.HasLink = v }
func (m *TLChannel) GetHasLink() bool  { return m.Data2.HasLink }

func (m *TLChannel) SetHasGeo(v bool) { m.Data2.HasGeo = v }
func (m *TLChannel) GetHasGeo() bool  { return m.Data2.HasGeo }

func (m *TLChannel) SetSlowmodeEnabled(v bool) { m.Data2.SlowmodeEnabled = v }
func (m *TLChannel) GetSlowmodeEnabled() bool  { return m.Data2.SlowmodeEnabled }

func (m *TLChannel) SetCallActive(v bool) { m.Data2.CallActive = v }
func (m *TLChannel) GetCallActive() bool  { return m.Data2.CallActive }

func (m *TLChannel) SetCallNotEmpty(v bool) { m.Data2.CallNotEmpty = v }
func (m *TLChannel) GetCallNotEmpty() bool  { return m.Data2.CallNotEmpty }

func (m *TLChannel) SetRestrictionReasonD31A961E122(v []*RestrictionReason) {
	m.Data2.RestrictionReasonD31A961E122 = v
}
func (m *TLChannel) GetRestrictionReasonD31A961E122() []*RestrictionReason {
	return m.Data2.RestrictionReasonD31A961E122
}

func (m *TLChannelForbidden) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChannelForbidden) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChannelForbidden) SetBroadcast(v bool) { m.Data2.Broadcast = v }
func (m *TLChannelForbidden) GetBroadcast() bool  { return m.Data2.Broadcast }

func (m *TLChannelForbidden) SetMegagroup(v bool) { m.Data2.Megagroup = v }
func (m *TLChannelForbidden) GetMegagroup() bool  { return m.Data2.Megagroup }

func (m *TLChannelForbidden) SetId(v int32) { m.Data2.Id = v }
func (m *TLChannelForbidden) GetId() int32  { return m.Data2.Id }

func (m *TLChannelForbidden) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLChannelForbidden) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLChannelForbidden) SetTitle(v string) { m.Data2.Title = v }
func (m *TLChannelForbidden) GetTitle() string  { return m.Data2.Title }

func (m *TLChannelForbidden) SetUntilDate(v int32) { m.Data2.UntilDate = v }
func (m *TLChannelForbidden) GetUntilDate() int32  { return m.Data2.UntilDate }

func (m *TLChatAdminRights) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChatAdminRights) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChatAdminRights) SetChangeInfo(v bool) { m.Data2.ChangeInfo = v }
func (m *TLChatAdminRights) GetChangeInfo() bool  { return m.Data2.ChangeInfo }

func (m *TLChatAdminRights) SetPostMessages(v bool) { m.Data2.PostMessages = v }
func (m *TLChatAdminRights) GetPostMessages() bool  { return m.Data2.PostMessages }

func (m *TLChatAdminRights) SetEditMessages(v bool) { m.Data2.EditMessages = v }
func (m *TLChatAdminRights) GetEditMessages() bool  { return m.Data2.EditMessages }

func (m *TLChatAdminRights) SetDeleteMessages(v bool) { m.Data2.DeleteMessages = v }
func (m *TLChatAdminRights) GetDeleteMessages() bool  { return m.Data2.DeleteMessages }

func (m *TLChatAdminRights) SetBanUsers(v bool) { m.Data2.BanUsers = v }
func (m *TLChatAdminRights) GetBanUsers() bool  { return m.Data2.BanUsers }

func (m *TLChatAdminRights) SetInviteUsers(v bool) { m.Data2.InviteUsers = v }
func (m *TLChatAdminRights) GetInviteUsers() bool  { return m.Data2.InviteUsers }

func (m *TLChatAdminRights) SetPinMessages(v bool) { m.Data2.PinMessages = v }
func (m *TLChatAdminRights) GetPinMessages() bool  { return m.Data2.PinMessages }

func (m *TLChatAdminRights) SetAddAdmins(v bool) { m.Data2.AddAdmins = v }
func (m *TLChatAdminRights) GetAddAdmins() bool  { return m.Data2.AddAdmins }

func (m *TLChatAdminRights) SetAnonymous(v bool) { m.Data2.Anonymous = v }
func (m *TLChatAdminRights) GetAnonymous() bool  { return m.Data2.Anonymous }

func (m *TLChatAdminRights) SetManageCall(v bool) { m.Data2.ManageCall = v }
func (m *TLChatAdminRights) GetManageCall() bool  { return m.Data2.ManageCall }

func (m *TLChatBannedRights) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChatBannedRights) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChatBannedRights) SetViewMessages(v bool) { m.Data2.ViewMessages = v }
func (m *TLChatBannedRights) GetViewMessages() bool  { return m.Data2.ViewMessages }

func (m *TLChatBannedRights) SetSendMessages(v bool) { m.Data2.SendMessages = v }
func (m *TLChatBannedRights) GetSendMessages() bool  { return m.Data2.SendMessages }

func (m *TLChatBannedRights) SetSendMedia(v bool) { m.Data2.SendMedia = v }
func (m *TLChatBannedRights) GetSendMedia() bool  { return m.Data2.SendMedia }

func (m *TLChatBannedRights) SetSendStickers(v bool) { m.Data2.SendStickers = v }
func (m *TLChatBannedRights) GetSendStickers() bool  { return m.Data2.SendStickers }

func (m *TLChatBannedRights) SetSendGifs(v bool) { m.Data2.SendGifs = v }
func (m *TLChatBannedRights) GetSendGifs() bool  { return m.Data2.SendGifs }

func (m *TLChatBannedRights) SetSendGames(v bool) { m.Data2.SendGames = v }
func (m *TLChatBannedRights) GetSendGames() bool  { return m.Data2.SendGames }

func (m *TLChatBannedRights) SetSendInline(v bool) { m.Data2.SendInline = v }
func (m *TLChatBannedRights) GetSendInline() bool  { return m.Data2.SendInline }

func (m *TLChatBannedRights) SetEmbedLinks(v bool) { m.Data2.EmbedLinks = v }
func (m *TLChatBannedRights) GetEmbedLinks() bool  { return m.Data2.EmbedLinks }

func (m *TLChatBannedRights) SetSendPolls(v bool) { m.Data2.SendPolls = v }
func (m *TLChatBannedRights) GetSendPolls() bool  { return m.Data2.SendPolls }

func (m *TLChatBannedRights) SetChangeInfo(v bool) { m.Data2.ChangeInfo = v }
func (m *TLChatBannedRights) GetChangeInfo() bool  { return m.Data2.ChangeInfo }

func (m *TLChatBannedRights) SetInviteUsers(v bool) { m.Data2.InviteUsers = v }
func (m *TLChatBannedRights) GetInviteUsers() bool  { return m.Data2.InviteUsers }

func (m *TLChatBannedRights) SetPinMessages(v bool) { m.Data2.PinMessages = v }
func (m *TLChatBannedRights) GetPinMessages() bool  { return m.Data2.PinMessages }

func (m *TLChatBannedRights) SetUntilDate(v int32) { m.Data2.UntilDate = v }
func (m *TLChatBannedRights) GetUntilDate() int32  { return m.Data2.UntilDate }

func (m *TLChatFull) SetId(v int32) { m.Data2.Id = v }
func (m *TLChatFull) GetId() int32  { return m.Data2.Id }

func (m *TLChatFull) SetParticipants(v *ChatParticipants) { m.Data2.Participants = v }
func (m *TLChatFull) GetParticipants() *ChatParticipants  { return m.Data2.Participants }

func (m *TLChatFull) SetChatPhoto(v *Photo) { m.Data2.ChatPhoto = v }
func (m *TLChatFull) GetChatPhoto() *Photo  { return m.Data2.ChatPhoto }

func (m *TLChatFull) SetNotifySettings(v *PeerNotifySettings) { m.Data2.NotifySettings = v }
func (m *TLChatFull) GetNotifySettings() *PeerNotifySettings  { return m.Data2.NotifySettings }

func (m *TLChatFull) SetExportedInvite(v *ExportedChatInvite) { m.Data2.ExportedInvite = v }
func (m *TLChatFull) GetExportedInvite() *ExportedChatInvite  { return m.Data2.ExportedInvite }

func (m *TLChatFull) SetBotInfo(v []*BotInfo) { m.Data2.BotInfo = v }
func (m *TLChatFull) GetBotInfo() []*BotInfo  { return m.Data2.BotInfo }

func (m *TLChatFull) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChatFull) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChatFull) SetPinnedMsgId(v int32) { m.Data2.PinnedMsgId = v }
func (m *TLChatFull) GetPinnedMsgId() int32  { return m.Data2.PinnedMsgId }

func (m *TLChatFull) SetCanSetUsername(v bool) { m.Data2.CanSetUsername = v }
func (m *TLChatFull) GetCanSetUsername() bool  { return m.Data2.CanSetUsername }

func (m *TLChatFull) SetAbout(v string) { m.Data2.About = v }
func (m *TLChatFull) GetAbout() string  { return m.Data2.About }

func (m *TLChatFull) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLChatFull) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLChatFull) SetHasScheduled(v bool) { m.Data2.HasScheduled = v }
func (m *TLChatFull) GetHasScheduled() bool  { return m.Data2.HasScheduled }

func (m *TLChatFull) SetCall(v *InputGroupCall) { m.Data2.Call = v }
func (m *TLChatFull) GetCall() *InputGroupCall  { return m.Data2.Call }

func (m *TLChannelFull) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChannelFull) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChannelFull) SetCanViewParticipants(v bool) { m.Data2.CanViewParticipants = v }
func (m *TLChannelFull) GetCanViewParticipants() bool  { return m.Data2.CanViewParticipants }

func (m *TLChannelFull) SetCanSetUsername(v bool) { m.Data2.CanSetUsername = v }
func (m *TLChannelFull) GetCanSetUsername() bool  { return m.Data2.CanSetUsername }

func (m *TLChannelFull) SetCanSetStickers(v bool) { m.Data2.CanSetStickers = v }
func (m *TLChannelFull) GetCanSetStickers() bool  { return m.Data2.CanSetStickers }

func (m *TLChannelFull) SetId(v int32) { m.Data2.Id = v }
func (m *TLChannelFull) GetId() int32  { return m.Data2.Id }

func (m *TLChannelFull) SetAbout(v string) { m.Data2.About = v }
func (m *TLChannelFull) GetAbout() string  { return m.Data2.About }

func (m *TLChannelFull) SetParticipantsCount(v int32) { m.Data2.ParticipantsCount = v }
func (m *TLChannelFull) GetParticipantsCount() int32  { return m.Data2.ParticipantsCount }

func (m *TLChannelFull) SetAdminsCount(v int32) { m.Data2.AdminsCount = v }
func (m *TLChannelFull) GetAdminsCount() int32  { return m.Data2.AdminsCount }

func (m *TLChannelFull) SetKickedCount(v int32) { m.Data2.KickedCount = v }
func (m *TLChannelFull) GetKickedCount() int32  { return m.Data2.KickedCount }

func (m *TLChannelFull) SetBannedCount(v int32) { m.Data2.BannedCount = v }
func (m *TLChannelFull) GetBannedCount() int32  { return m.Data2.BannedCount }

func (m *TLChannelFull) SetReadInboxMaxId(v int32) { m.Data2.ReadInboxMaxId = v }
func (m *TLChannelFull) GetReadInboxMaxId() int32  { return m.Data2.ReadInboxMaxId }

func (m *TLChannelFull) SetReadOutboxMaxId(v int32) { m.Data2.ReadOutboxMaxId = v }
func (m *TLChannelFull) GetReadOutboxMaxId() int32  { return m.Data2.ReadOutboxMaxId }

func (m *TLChannelFull) SetUnreadCount(v int32) { m.Data2.UnreadCount = v }
func (m *TLChannelFull) GetUnreadCount() int32  { return m.Data2.UnreadCount }

func (m *TLChannelFull) SetChatPhoto(v *Photo) { m.Data2.ChatPhoto = v }
func (m *TLChannelFull) GetChatPhoto() *Photo  { return m.Data2.ChatPhoto }

func (m *TLChannelFull) SetNotifySettings(v *PeerNotifySettings) { m.Data2.NotifySettings = v }
func (m *TLChannelFull) GetNotifySettings() *PeerNotifySettings  { return m.Data2.NotifySettings }

func (m *TLChannelFull) SetExportedInvite(v *ExportedChatInvite) { m.Data2.ExportedInvite = v }
func (m *TLChannelFull) GetExportedInvite() *ExportedChatInvite  { return m.Data2.ExportedInvite }

func (m *TLChannelFull) SetBotInfo(v []*BotInfo) { m.Data2.BotInfo = v }
func (m *TLChannelFull) GetBotInfo() []*BotInfo  { return m.Data2.BotInfo }

func (m *TLChannelFull) SetMigratedFromChatId(v int32) { m.Data2.MigratedFromChatId = v }
func (m *TLChannelFull) GetMigratedFromChatId() int32  { return m.Data2.MigratedFromChatId }

func (m *TLChannelFull) SetMigratedFromMaxId(v int32) { m.Data2.MigratedFromMaxId = v }
func (m *TLChannelFull) GetMigratedFromMaxId() int32  { return m.Data2.MigratedFromMaxId }

func (m *TLChannelFull) SetPinnedMsgId(v int32) { m.Data2.PinnedMsgId = v }
func (m *TLChannelFull) GetPinnedMsgId() int32  { return m.Data2.PinnedMsgId }

func (m *TLChannelFull) SetStickerset(v *StickerSet) { m.Data2.Stickerset = v }
func (m *TLChannelFull) GetStickerset() *StickerSet  { return m.Data2.Stickerset }

func (m *TLChannelFull) SetUnreadImportantCount(v int32) { m.Data2.UnreadImportantCount = v }
func (m *TLChannelFull) GetUnreadImportantCount() int32  { return m.Data2.UnreadImportantCount }

func (m *TLChannelFull) SetHiddenPrehistory(v bool) { m.Data2.HiddenPrehistory = v }
func (m *TLChannelFull) GetHiddenPrehistory() bool  { return m.Data2.HiddenPrehistory }

func (m *TLChannelFull) SetAvailableMinId(v int32) { m.Data2.AvailableMinId = v }
func (m *TLChannelFull) GetAvailableMinId() int32  { return m.Data2.AvailableMinId }

func (m *TLChannelFull) SetCanViewStats(v bool) { m.Data2.CanViewStats = v }
func (m *TLChannelFull) GetCanViewStats() bool  { return m.Data2.CanViewStats }

func (m *TLChannelFull) SetOnlineCount(v int32) { m.Data2.OnlineCount = v }
func (m *TLChannelFull) GetOnlineCount() int32  { return m.Data2.OnlineCount }

func (m *TLChannelFull) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLChannelFull) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLChannelFull) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLChannelFull) GetPts() int32  { return m.Data2.Pts }

func (m *TLChannelFull) SetLinkedChatId(v int32) { m.Data2.LinkedChatId = v }
func (m *TLChannelFull) GetLinkedChatId() int32  { return m.Data2.LinkedChatId }

func (m *TLChannelFull) SetCanSetLocation(v bool) { m.Data2.CanSetLocation = v }
func (m *TLChannelFull) GetCanSetLocation() bool  { return m.Data2.CanSetLocation }

func (m *TLChannelFull) SetLocation(v *ChannelLocation) { m.Data2.Location = v }
func (m *TLChannelFull) GetLocation() *ChannelLocation  { return m.Data2.Location }

func (m *TLChannelFull) SetSlowmodeSeconds(v int32) { m.Data2.SlowmodeSeconds = v }
func (m *TLChannelFull) GetSlowmodeSeconds() int32  { return m.Data2.SlowmodeSeconds }

func (m *TLChannelFull) SetSlowmodeNextSendDate(v int32) { m.Data2.SlowmodeNextSendDate = v }
func (m *TLChannelFull) GetSlowmodeNextSendDate() int32  { return m.Data2.SlowmodeNextSendDate }

func (m *TLChannelFull) SetHasScheduled(v bool) { m.Data2.HasScheduled = v }
func (m *TLChannelFull) GetHasScheduled() bool  { return m.Data2.HasScheduled }

func (m *TLChannelFull) SetBlocked(v bool) { m.Data2.Blocked = v }
func (m *TLChannelFull) GetBlocked() bool  { return m.Data2.Blocked }

func (m *TLChannelFull) SetStatsDc(v int32) { m.Data2.StatsDc = v }
func (m *TLChannelFull) GetStatsDc() int32  { return m.Data2.StatsDc }

func (m *TLChannelFull) SetCall(v *InputGroupCall) { m.Data2.Call = v }
func (m *TLChannelFull) GetCall() *InputGroupCall  { return m.Data2.Call }

func (m *TLChatInviteAlready) SetChat(v *Chat) { m.Data2.Chat = v }
func (m *TLChatInviteAlready) GetChat() *Chat  { return m.Data2.Chat }

func (m *TLChatInvite) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChatInvite) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChatInvite) SetChannel(v bool) { m.Data2.Channel = v }
func (m *TLChatInvite) GetChannel() bool  { return m.Data2.Channel }

func (m *TLChatInvite) SetBroadcast(v bool) { m.Data2.Broadcast = v }
func (m *TLChatInvite) GetBroadcast() bool  { return m.Data2.Broadcast }

func (m *TLChatInvite) SetPublic(v bool) { m.Data2.Public = v }
func (m *TLChatInvite) GetPublic() bool  { return m.Data2.Public }

func (m *TLChatInvite) SetMegagroup(v bool) { m.Data2.Megagroup = v }
func (m *TLChatInvite) GetMegagroup() bool  { return m.Data2.Megagroup }

func (m *TLChatInvite) SetTitle(v string) { m.Data2.Title = v }
func (m *TLChatInvite) GetTitle() string  { return m.Data2.Title }

func (m *TLChatInvite) SetPhotoDB74F55871(v *ChatPhoto) { m.Data2.PhotoDB74F55871 = v }
func (m *TLChatInvite) GetPhotoDB74F55871() *ChatPhoto  { return m.Data2.PhotoDB74F55871 }

func (m *TLChatInvite) SetParticipantsCount(v int32) { m.Data2.ParticipantsCount = v }
func (m *TLChatInvite) GetParticipantsCount() int32  { return m.Data2.ParticipantsCount }

func (m *TLChatInvite) SetParticipants(v []*User) { m.Data2.Participants = v }
func (m *TLChatInvite) GetParticipants() []*User  { return m.Data2.Participants }

func (m *TLChatInvite) SetPhotoDFC2F58E98(v *Photo) { m.Data2.PhotoDFC2F58E98 = v }
func (m *TLChatInvite) GetPhotoDFC2F58E98() *Photo  { return m.Data2.PhotoDFC2F58E98 }

func (m *TLChatInvitePeek) SetChat(v *Chat) { m.Data2.Chat = v }
func (m *TLChatInvitePeek) GetChat() *Chat  { return m.Data2.Chat }

func (m *TLChatInvitePeek) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLChatInvitePeek) GetExpires() int32  { return m.Data2.Expires }

func (m *TLChatOnlines) SetOnlines(v int32) { m.Data2.Onlines = v }
func (m *TLChatOnlines) GetOnlines() int32  { return m.Data2.Onlines }

func (m *TLChatParticipant) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChatParticipant) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChatParticipant) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLChatParticipant) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLChatParticipant) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChatParticipant) GetDate() int32  { return m.Data2.Date }

func (m *TLChatParticipantCreator) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChatParticipantCreator) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChatParticipantAdmin) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChatParticipantAdmin) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChatParticipantAdmin) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLChatParticipantAdmin) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLChatParticipantAdmin) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChatParticipantAdmin) GetDate() int32  { return m.Data2.Date }

func (m *TLChatParticipantsForbidden) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChatParticipantsForbidden) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChatParticipantsForbidden) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLChatParticipantsForbidden) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLChatParticipantsForbidden) SetSelfParticipant(v *ChatParticipant) {
	m.Data2.SelfParticipant = v
}
func (m *TLChatParticipantsForbidden) GetSelfParticipant() *ChatParticipant {
	return m.Data2.SelfParticipant
}

func (m *TLChatParticipants) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLChatParticipants) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLChatParticipants) SetParticipants(v []*ChatParticipant) { m.Data2.Participants = v }
func (m *TLChatParticipants) GetParticipants() []*ChatParticipant  { return m.Data2.Participants }

func (m *TLChatParticipants) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLChatParticipants) GetVersion() int32  { return m.Data2.Version }

func (m *TLChatPhoto) SetPhotoSmall(v *FileLocation) { m.Data2.PhotoSmall = v }
func (m *TLChatPhoto) GetPhotoSmall() *FileLocation  { return m.Data2.PhotoSmall }

func (m *TLChatPhoto) SetPhotoBig(v *FileLocation) { m.Data2.PhotoBig = v }
func (m *TLChatPhoto) GetPhotoBig() *FileLocation  { return m.Data2.PhotoBig }

func (m *TLChatPhoto) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLChatPhoto) GetDcId() int32  { return m.Data2.DcId }

func (m *TLChatPhoto) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLChatPhoto) GetFlags() int32  { return m.Data2.Flags }

func (m *TLChatPhoto) SetHasVideo(v bool) { m.Data2.HasVideo = v }
func (m *TLChatPhoto) GetHasVideo() bool  { return m.Data2.HasVideo }

func (m *TLClient_DHInnerData) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLClient_DHInnerData) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLClient_DHInnerData) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLClient_DHInnerData) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLClient_DHInnerData) SetRetryId(v int64) { m.Data2.RetryId = v }
func (m *TLClient_DHInnerData) GetRetryId() int64  { return m.Data2.RetryId }

func (m *TLClient_DHInnerData) SetGB(v string) { m.Data2.GB = v }
func (m *TLClient_DHInnerData) GetGB() string  { return m.Data2.GB }

func (m *TLCodeSettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLCodeSettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLCodeSettings) SetAllowFlashcall(v bool) { m.Data2.AllowFlashcall = v }
func (m *TLCodeSettings) GetAllowFlashcall() bool  { return m.Data2.AllowFlashcall }

func (m *TLCodeSettings) SetCurrentNumber(v bool) { m.Data2.CurrentNumber = v }
func (m *TLCodeSettings) GetCurrentNumber() bool  { return m.Data2.CurrentNumber }

func (m *TLCodeSettings) SetAppHashPersistent(v bool) { m.Data2.AppHashPersistent = v }
func (m *TLCodeSettings) GetAppHashPersistent() bool  { return m.Data2.AppHashPersistent }

func (m *TLCodeSettings) SetAppHash(v string) { m.Data2.AppHash = v }
func (m *TLCodeSettings) GetAppHash() string  { return m.Data2.AppHash }

func (m *TLCodeSettings) SetAllowAppHash(v bool) { m.Data2.AllowAppHash = v }
func (m *TLCodeSettings) GetAllowAppHash() bool  { return m.Data2.AllowAppHash }

func (m *TLConfig) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLConfig) GetFlags() int32  { return m.Data2.Flags }

func (m *TLConfig) SetPhonecallsEnabled(v bool) { m.Data2.PhonecallsEnabled = v }
func (m *TLConfig) GetPhonecallsEnabled() bool  { return m.Data2.PhonecallsEnabled }

func (m *TLConfig) SetDate(v int32) { m.Data2.Date = v }
func (m *TLConfig) GetDate() int32  { return m.Data2.Date }

func (m *TLConfig) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLConfig) GetExpires() int32  { return m.Data2.Expires }

func (m *TLConfig) SetTestMode(v *Bool) { m.Data2.TestMode = v }
func (m *TLConfig) GetTestMode() *Bool  { return m.Data2.TestMode }

func (m *TLConfig) SetThisDc(v int32) { m.Data2.ThisDc = v }
func (m *TLConfig) GetThisDc() int32  { return m.Data2.ThisDc }

func (m *TLConfig) SetDcOptions(v []*DcOption) { m.Data2.DcOptions = v }
func (m *TLConfig) GetDcOptions() []*DcOption  { return m.Data2.DcOptions }

func (m *TLConfig) SetChatSizeMax(v int32) { m.Data2.ChatSizeMax = v }
func (m *TLConfig) GetChatSizeMax() int32  { return m.Data2.ChatSizeMax }

func (m *TLConfig) SetMegagroupSizeMax(v int32) { m.Data2.MegagroupSizeMax = v }
func (m *TLConfig) GetMegagroupSizeMax() int32  { return m.Data2.MegagroupSizeMax }

func (m *TLConfig) SetForwardedCountMax(v int32) { m.Data2.ForwardedCountMax = v }
func (m *TLConfig) GetForwardedCountMax() int32  { return m.Data2.ForwardedCountMax }

func (m *TLConfig) SetOnlineUpdatePeriodMs(v int32) { m.Data2.OnlineUpdatePeriodMs = v }
func (m *TLConfig) GetOnlineUpdatePeriodMs() int32  { return m.Data2.OnlineUpdatePeriodMs }

func (m *TLConfig) SetOfflineBlurTimeoutMs(v int32) { m.Data2.OfflineBlurTimeoutMs = v }
func (m *TLConfig) GetOfflineBlurTimeoutMs() int32  { return m.Data2.OfflineBlurTimeoutMs }

func (m *TLConfig) SetOfflineIdleTimeoutMs(v int32) { m.Data2.OfflineIdleTimeoutMs = v }
func (m *TLConfig) GetOfflineIdleTimeoutMs() int32  { return m.Data2.OfflineIdleTimeoutMs }

func (m *TLConfig) SetOnlineCloudTimeoutMs(v int32) { m.Data2.OnlineCloudTimeoutMs = v }
func (m *TLConfig) GetOnlineCloudTimeoutMs() int32  { return m.Data2.OnlineCloudTimeoutMs }

func (m *TLConfig) SetNotifyCloudDelayMs(v int32) { m.Data2.NotifyCloudDelayMs = v }
func (m *TLConfig) GetNotifyCloudDelayMs() int32  { return m.Data2.NotifyCloudDelayMs }

func (m *TLConfig) SetNotifyDefaultDelayMs(v int32) { m.Data2.NotifyDefaultDelayMs = v }
func (m *TLConfig) GetNotifyDefaultDelayMs() int32  { return m.Data2.NotifyDefaultDelayMs }

func (m *TLConfig) SetChatBigSize(v int32) { m.Data2.ChatBigSize = v }
func (m *TLConfig) GetChatBigSize() int32  { return m.Data2.ChatBigSize }

func (m *TLConfig) SetPushChatPeriodMs(v int32) { m.Data2.PushChatPeriodMs = v }
func (m *TLConfig) GetPushChatPeriodMs() int32  { return m.Data2.PushChatPeriodMs }

func (m *TLConfig) SetPushChatLimit(v int32) { m.Data2.PushChatLimit = v }
func (m *TLConfig) GetPushChatLimit() int32  { return m.Data2.PushChatLimit }

func (m *TLConfig) SetSavedGifsLimit(v int32) { m.Data2.SavedGifsLimit = v }
func (m *TLConfig) GetSavedGifsLimit() int32  { return m.Data2.SavedGifsLimit }

func (m *TLConfig) SetEditTimeLimit(v int32) { m.Data2.EditTimeLimit = v }
func (m *TLConfig) GetEditTimeLimit() int32  { return m.Data2.EditTimeLimit }

func (m *TLConfig) SetRatingEDecay(v int32) { m.Data2.RatingEDecay = v }
func (m *TLConfig) GetRatingEDecay() int32  { return m.Data2.RatingEDecay }

func (m *TLConfig) SetStickersRecentLimit(v int32) { m.Data2.StickersRecentLimit = v }
func (m *TLConfig) GetStickersRecentLimit() int32  { return m.Data2.StickersRecentLimit }

func (m *TLConfig) SetStickersFavedLimit(v int32) { m.Data2.StickersFavedLimit = v }
func (m *TLConfig) GetStickersFavedLimit() int32  { return m.Data2.StickersFavedLimit }

func (m *TLConfig) SetTmpSessions(v int32) { m.Data2.TmpSessions = v }
func (m *TLConfig) GetTmpSessions() int32  { return m.Data2.TmpSessions }

func (m *TLConfig) SetPinnedDialogsCountMax(v int32) { m.Data2.PinnedDialogsCountMax = v }
func (m *TLConfig) GetPinnedDialogsCountMax() int32  { return m.Data2.PinnedDialogsCountMax }

func (m *TLConfig) SetCallReceiveTimeoutMs(v int32) { m.Data2.CallReceiveTimeoutMs = v }
func (m *TLConfig) GetCallReceiveTimeoutMs() int32  { return m.Data2.CallReceiveTimeoutMs }

func (m *TLConfig) SetCallRingTimeoutMs(v int32) { m.Data2.CallRingTimeoutMs = v }
func (m *TLConfig) GetCallRingTimeoutMs() int32  { return m.Data2.CallRingTimeoutMs }

func (m *TLConfig) SetCallConnectTimeoutMs(v int32) { m.Data2.CallConnectTimeoutMs = v }
func (m *TLConfig) GetCallConnectTimeoutMs() int32  { return m.Data2.CallConnectTimeoutMs }

func (m *TLConfig) SetCallPacketTimeoutMs(v int32) { m.Data2.CallPacketTimeoutMs = v }
func (m *TLConfig) GetCallPacketTimeoutMs() int32  { return m.Data2.CallPacketTimeoutMs }

func (m *TLConfig) SetMeUrlPrefix(v string) { m.Data2.MeUrlPrefix = v }
func (m *TLConfig) GetMeUrlPrefix() string  { return m.Data2.MeUrlPrefix }

func (m *TLConfig) SetSuggestedLangCode(v string) { m.Data2.SuggestedLangCode = v }
func (m *TLConfig) GetSuggestedLangCode() string  { return m.Data2.SuggestedLangCode }

func (m *TLConfig) SetLangPackVersion(v int32) { m.Data2.LangPackVersion = v }
func (m *TLConfig) GetLangPackVersion() int32  { return m.Data2.LangPackVersion }

func (m *TLConfig) SetDisabledFeatures(v []*DisabledFeature) { m.Data2.DisabledFeatures = v }
func (m *TLConfig) GetDisabledFeatures() []*DisabledFeature  { return m.Data2.DisabledFeatures }

func (m *TLConfig) SetDefaultP2PContacts(v bool) { m.Data2.DefaultP2PContacts = v }
func (m *TLConfig) GetDefaultP2PContacts() bool  { return m.Data2.DefaultP2PContacts }

func (m *TLConfig) SetPreloadFeaturedStickers(v bool) { m.Data2.PreloadFeaturedStickers = v }
func (m *TLConfig) GetPreloadFeaturedStickers() bool  { return m.Data2.PreloadFeaturedStickers }

func (m *TLConfig) SetIgnorePhoneEntities(v bool) { m.Data2.IgnorePhoneEntities = v }
func (m *TLConfig) GetIgnorePhoneEntities() bool  { return m.Data2.IgnorePhoneEntities }

func (m *TLConfig) SetRevokePmInbox(v bool) { m.Data2.RevokePmInbox = v }
func (m *TLConfig) GetRevokePmInbox() bool  { return m.Data2.RevokePmInbox }

func (m *TLConfig) SetBlockedMode(v bool) { m.Data2.BlockedMode = v }
func (m *TLConfig) GetBlockedMode() bool  { return m.Data2.BlockedMode }

func (m *TLConfig) SetDcTxtDomainName(v string) { m.Data2.DcTxtDomainName = v }
func (m *TLConfig) GetDcTxtDomainName() string  { return m.Data2.DcTxtDomainName }

func (m *TLConfig) SetRevokeTimeLimit(v int32) { m.Data2.RevokeTimeLimit = v }
func (m *TLConfig) GetRevokeTimeLimit() int32  { return m.Data2.RevokeTimeLimit }

func (m *TLConfig) SetRevokePmTimeLimit(v int32) { m.Data2.RevokePmTimeLimit = v }
func (m *TLConfig) GetRevokePmTimeLimit() int32  { return m.Data2.RevokePmTimeLimit }

func (m *TLConfig) SetChannelsReadMediaPeriod(v int32) { m.Data2.ChannelsReadMediaPeriod = v }
func (m *TLConfig) GetChannelsReadMediaPeriod() int32  { return m.Data2.ChannelsReadMediaPeriod }

func (m *TLConfig) SetAutoupdateUrlPrefix(v string) { m.Data2.AutoupdateUrlPrefix = v }
func (m *TLConfig) GetAutoupdateUrlPrefix() string  { return m.Data2.AutoupdateUrlPrefix }

func (m *TLConfig) SetGifSearchUsername(v string) { m.Data2.GifSearchUsername = v }
func (m *TLConfig) GetGifSearchUsername() string  { return m.Data2.GifSearchUsername }

func (m *TLConfig) SetVenueSearchUsername(v string) { m.Data2.VenueSearchUsername = v }
func (m *TLConfig) GetVenueSearchUsername() string  { return m.Data2.VenueSearchUsername }

func (m *TLConfig) SetImgSearchUsername(v string) { m.Data2.ImgSearchUsername = v }
func (m *TLConfig) GetImgSearchUsername() string  { return m.Data2.ImgSearchUsername }

func (m *TLConfig) SetStaticMapsProvider(v string) { m.Data2.StaticMapsProvider = v }
func (m *TLConfig) GetStaticMapsProvider() string  { return m.Data2.StaticMapsProvider }

func (m *TLConfig) SetCaptionLengthMax(v int32) { m.Data2.CaptionLengthMax = v }
func (m *TLConfig) GetCaptionLengthMax() int32  { return m.Data2.CaptionLengthMax }

func (m *TLConfig) SetMessageLengthMax(v int32) { m.Data2.MessageLengthMax = v }
func (m *TLConfig) GetMessageLengthMax() int32  { return m.Data2.MessageLengthMax }

func (m *TLConfig) SetWebfileDcId(v int32) { m.Data2.WebfileDcId = v }
func (m *TLConfig) GetWebfileDcId() int32  { return m.Data2.WebfileDcId }

func (m *TLConfig) SetBaseLangPackVersion(v int32) { m.Data2.BaseLangPackVersion = v }
func (m *TLConfig) GetBaseLangPackVersion() int32  { return m.Data2.BaseLangPackVersion }

func (m *TLConfig) SetPfsEnabled(v bool) { m.Data2.PfsEnabled = v }
func (m *TLConfig) GetPfsEnabled() bool  { return m.Data2.PfsEnabled }

func (m *TLConfig) SetPinnedInfolderCountMax(v int32) { m.Data2.PinnedInfolderCountMax = v }
func (m *TLConfig) GetPinnedInfolderCountMax() int32  { return m.Data2.PinnedInfolderCountMax }

func (m *TLContact) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLContact) GetUserId() int32  { return m.Data2.UserId }

func (m *TLContact) SetMutual(v *Bool) { m.Data2.Mutual = v }
func (m *TLContact) GetMutual() *Bool  { return m.Data2.Mutual }

func (m *TLContactBlocked) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLContactBlocked) GetUserId() int32  { return m.Data2.UserId }

func (m *TLContactBlocked) SetDate(v int32) { m.Data2.Date = v }
func (m *TLContactBlocked) GetDate() int32  { return m.Data2.Date }

func (m *TLContactStatus) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLContactStatus) GetUserId() int32  { return m.Data2.UserId }

func (m *TLContactStatus) SetStatus(v *UserStatus) { m.Data2.Status = v }
func (m *TLContactStatus) GetStatus() *UserStatus  { return m.Data2.Status }

func (m *TLContactsBlocked) SetBlocked1C138D1571(v []*ContactBlocked) { m.Data2.Blocked1C138D1571 = v }
func (m *TLContactsBlocked) GetBlocked1C138D1571() []*ContactBlocked {
	return m.Data2.Blocked1C138D1571
}

func (m *TLContactsBlocked) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsBlocked) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsBlocked) SetBlockedADE1591119(v []*PeerBlocked) { m.Data2.BlockedADE1591119 = v }
func (m *TLContactsBlocked) GetBlockedADE1591119() []*PeerBlocked  { return m.Data2.BlockedADE1591119 }

func (m *TLContactsBlocked) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLContactsBlocked) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLContactsBlockedSlice) SetCount(v int32) { m.Data2.Count = v }
func (m *TLContactsBlockedSlice) GetCount() int32  { return m.Data2.Count }

func (m *TLContactsBlockedSlice) SetBlocked1C138D1571(v []*ContactBlocked) {
	m.Data2.Blocked1C138D1571 = v
}
func (m *TLContactsBlockedSlice) GetBlocked1C138D1571() []*ContactBlocked {
	return m.Data2.Blocked1C138D1571
}

func (m *TLContactsBlockedSlice) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsBlockedSlice) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsBlockedSlice) SetBlockedADE1591119(v []*PeerBlocked) {
	m.Data2.BlockedADE1591119 = v
}
func (m *TLContactsBlockedSlice) GetBlockedADE1591119() []*PeerBlocked {
	return m.Data2.BlockedADE1591119
}

func (m *TLContactsBlockedSlice) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLContactsBlockedSlice) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLContactsContacts) SetContacts(v []*Contact) { m.Data2.Contacts = v }
func (m *TLContactsContacts) GetContacts() []*Contact  { return m.Data2.Contacts }

func (m *TLContactsContacts) SetSavedCount(v int32) { m.Data2.SavedCount = v }
func (m *TLContactsContacts) GetSavedCount() int32  { return m.Data2.SavedCount }

func (m *TLContactsContacts) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsContacts) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsFound) SetResults(v []*Peer) { m.Data2.Results = v }
func (m *TLContactsFound) GetResults() []*Peer  { return m.Data2.Results }

func (m *TLContactsFound) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLContactsFound) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLContactsFound) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsFound) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsFound) SetMyResults(v []*Peer) { m.Data2.MyResults = v }
func (m *TLContactsFound) GetMyResults() []*Peer  { return m.Data2.MyResults }

func (m *TLContactsImportedContacts) SetImported(v []*ImportedContact) { m.Data2.Imported = v }
func (m *TLContactsImportedContacts) GetImported() []*ImportedContact  { return m.Data2.Imported }

func (m *TLContactsImportedContacts) SetPopularInvites(v []*PopularContact) {
	m.Data2.PopularInvites = v
}
func (m *TLContactsImportedContacts) GetPopularInvites() []*PopularContact {
	return m.Data2.PopularInvites
}

func (m *TLContactsImportedContacts) SetRetryContacts(v []int64) { m.Data2.RetryContacts = v }
func (m *TLContactsImportedContacts) GetRetryContacts() []int64  { return m.Data2.RetryContacts }

func (m *TLContactsImportedContacts) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsImportedContacts) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsLink) SetMyLink(v *ContactLink) { m.Data2.MyLink = v }
func (m *TLContactsLink) GetMyLink() *ContactLink  { return m.Data2.MyLink }

func (m *TLContactsLink) SetForeignLink(v *ContactLink) { m.Data2.ForeignLink = v }
func (m *TLContactsLink) GetForeignLink() *ContactLink  { return m.Data2.ForeignLink }

func (m *TLContactsLink) SetUser(v *User) { m.Data2.User = v }
func (m *TLContactsLink) GetUser() *User  { return m.Data2.User }

func (m *TLContactsResolvedPeer) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLContactsResolvedPeer) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLContactsResolvedPeer) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLContactsResolvedPeer) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLContactsResolvedPeer) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsResolvedPeer) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsTopPeers) SetCategories(v []*TopPeerCategoryPeers) { m.Data2.Categories = v }
func (m *TLContactsTopPeers) GetCategories() []*TopPeerCategoryPeers  { return m.Data2.Categories }

func (m *TLContactsTopPeers) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLContactsTopPeers) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLContactsTopPeers) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsTopPeers) GetUsers() []*User  { return m.Data2.Users }

func (m *TLDataJSON) SetData(v string) { m.Data2.Data = v }
func (m *TLDataJSON) GetData() string  { return m.Data2.Data }

func (m *TLDcOption) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLDcOption) GetFlags() int32  { return m.Data2.Flags }

func (m *TLDcOption) SetIpv6(v bool) { m.Data2.Ipv6 = v }
func (m *TLDcOption) GetIpv6() bool  { return m.Data2.Ipv6 }

func (m *TLDcOption) SetMediaOnly(v bool) { m.Data2.MediaOnly = v }
func (m *TLDcOption) GetMediaOnly() bool  { return m.Data2.MediaOnly }

func (m *TLDcOption) SetTcpoOnly(v bool) { m.Data2.TcpoOnly = v }
func (m *TLDcOption) GetTcpoOnly() bool  { return m.Data2.TcpoOnly }

func (m *TLDcOption) SetCdn(v bool) { m.Data2.Cdn = v }
func (m *TLDcOption) GetCdn() bool  { return m.Data2.Cdn }

func (m *TLDcOption) SetStatic(v bool) { m.Data2.Static = v }
func (m *TLDcOption) GetStatic() bool  { return m.Data2.Static }

func (m *TLDcOption) SetId(v int32) { m.Data2.Id = v }
func (m *TLDcOption) GetId() int32  { return m.Data2.Id }

func (m *TLDcOption) SetIpAddress(v string) { m.Data2.IpAddress = v }
func (m *TLDcOption) GetIpAddress() string  { return m.Data2.IpAddress }

func (m *TLDcOption) SetPort(v int32) { m.Data2.Port = v }
func (m *TLDcOption) GetPort() int32  { return m.Data2.Port }

func (m *TLDcOption) SetSecret(v []byte) { m.Data2.Secret = v }
func (m *TLDcOption) GetSecret() []byte  { return m.Data2.Secret }

func (m *TLDestroySessionOk) SetSessionId(v int64) { m.Data2.SessionId = v }
func (m *TLDestroySessionOk) GetSessionId() int64  { return m.Data2.SessionId }

func (m *TLDestroySessionNone) SetSessionId(v int64) { m.Data2.SessionId = v }
func (m *TLDestroySessionNone) GetSessionId() int64  { return m.Data2.SessionId }

func (m *TLDialog) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLDialog) GetFlags() int32  { return m.Data2.Flags }

func (m *TLDialog) SetPinned(v bool) { m.Data2.Pinned = v }
func (m *TLDialog) GetPinned() bool  { return m.Data2.Pinned }

func (m *TLDialog) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLDialog) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLDialog) SetTopMessage(v int32) { m.Data2.TopMessage = v }
func (m *TLDialog) GetTopMessage() int32  { return m.Data2.TopMessage }

func (m *TLDialog) SetReadInboxMaxId(v int32) { m.Data2.ReadInboxMaxId = v }
func (m *TLDialog) GetReadInboxMaxId() int32  { return m.Data2.ReadInboxMaxId }

func (m *TLDialog) SetReadOutboxMaxId(v int32) { m.Data2.ReadOutboxMaxId = v }
func (m *TLDialog) GetReadOutboxMaxId() int32  { return m.Data2.ReadOutboxMaxId }

func (m *TLDialog) SetUnreadCount(v int32) { m.Data2.UnreadCount = v }
func (m *TLDialog) GetUnreadCount() int32  { return m.Data2.UnreadCount }

func (m *TLDialog) SetUnreadMentionsCount(v int32) { m.Data2.UnreadMentionsCount = v }
func (m *TLDialog) GetUnreadMentionsCount() int32  { return m.Data2.UnreadMentionsCount }

func (m *TLDialog) SetNotifySettings(v *PeerNotifySettings) { m.Data2.NotifySettings = v }
func (m *TLDialog) GetNotifySettings() *PeerNotifySettings  { return m.Data2.NotifySettings }

func (m *TLDialog) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLDialog) GetPts() int32  { return m.Data2.Pts }

func (m *TLDialog) SetDraft(v *DraftMessage) { m.Data2.Draft = v }
func (m *TLDialog) GetDraft() *DraftMessage  { return m.Data2.Draft }

func (m *TLDialog) SetUnreadMark(v bool) { m.Data2.UnreadMark = v }
func (m *TLDialog) GetUnreadMark() bool  { return m.Data2.UnreadMark }

func (m *TLDialog) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLDialog) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLDialogChannel) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLDialogChannel) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLDialogChannel) SetTopMessage(v int32) { m.Data2.TopMessage = v }
func (m *TLDialogChannel) GetTopMessage() int32  { return m.Data2.TopMessage }

func (m *TLDialogChannel) SetTopImportantMessage(v int32) { m.Data2.TopImportantMessage = v }
func (m *TLDialogChannel) GetTopImportantMessage() int32  { return m.Data2.TopImportantMessage }

func (m *TLDialogChannel) SetReadInboxMaxId(v int32) { m.Data2.ReadInboxMaxId = v }
func (m *TLDialogChannel) GetReadInboxMaxId() int32  { return m.Data2.ReadInboxMaxId }

func (m *TLDialogChannel) SetUnreadCount(v int32) { m.Data2.UnreadCount = v }
func (m *TLDialogChannel) GetUnreadCount() int32  { return m.Data2.UnreadCount }

func (m *TLDialogChannel) SetUnreadImportantCount(v int32) { m.Data2.UnreadImportantCount = v }
func (m *TLDialogChannel) GetUnreadImportantCount() int32  { return m.Data2.UnreadImportantCount }

func (m *TLDialogChannel) SetNotifySettings(v *PeerNotifySettings) { m.Data2.NotifySettings = v }
func (m *TLDialogChannel) GetNotifySettings() *PeerNotifySettings  { return m.Data2.NotifySettings }

func (m *TLDialogChannel) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLDialogChannel) GetPts() int32  { return m.Data2.Pts }

func (m *TLDialogFolder) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLDialogFolder) GetFlags() int32  { return m.Data2.Flags }

func (m *TLDialogFolder) SetPinned(v bool) { m.Data2.Pinned = v }
func (m *TLDialogFolder) GetPinned() bool  { return m.Data2.Pinned }

func (m *TLDialogFolder) SetFolder(v *Folder) { m.Data2.Folder = v }
func (m *TLDialogFolder) GetFolder() *Folder  { return m.Data2.Folder }

func (m *TLDialogFolder) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLDialogFolder) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLDialogFolder) SetTopMessage(v int32) { m.Data2.TopMessage = v }
func (m *TLDialogFolder) GetTopMessage() int32  { return m.Data2.TopMessage }

func (m *TLDialogFolder) SetUnreadMutedPeersCount(v int32) { m.Data2.UnreadMutedPeersCount = v }
func (m *TLDialogFolder) GetUnreadMutedPeersCount() int32  { return m.Data2.UnreadMutedPeersCount }

func (m *TLDialogFolder) SetUnreadUnmutedPeersCount(v int32) { m.Data2.UnreadUnmutedPeersCount = v }
func (m *TLDialogFolder) GetUnreadUnmutedPeersCount() int32  { return m.Data2.UnreadUnmutedPeersCount }

func (m *TLDialogFolder) SetUnreadMutedMessagesCount(v int32) { m.Data2.UnreadMutedMessagesCount = v }
func (m *TLDialogFolder) GetUnreadMutedMessagesCount() int32  { return m.Data2.UnreadMutedMessagesCount }

func (m *TLDialogFolder) SetUnreadUnmutedMessagesCount(v int32) {
	m.Data2.UnreadUnmutedMessagesCount = v
}
func (m *TLDialogFolder) GetUnreadUnmutedMessagesCount() int32 {
	return m.Data2.UnreadUnmutedMessagesCount
}

func (m *TLDialogFilter) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLDialogFilter) GetFlags() int32  { return m.Data2.Flags }

func (m *TLDialogFilter) SetContacts(v bool) { m.Data2.Contacts = v }
func (m *TLDialogFilter) GetContacts() bool  { return m.Data2.Contacts }

func (m *TLDialogFilter) SetNonContacts(v bool) { m.Data2.NonContacts = v }
func (m *TLDialogFilter) GetNonContacts() bool  { return m.Data2.NonContacts }

func (m *TLDialogFilter) SetGroups(v bool) { m.Data2.Groups = v }
func (m *TLDialogFilter) GetGroups() bool  { return m.Data2.Groups }

func (m *TLDialogFilter) SetBroadcasts(v bool) { m.Data2.Broadcasts = v }
func (m *TLDialogFilter) GetBroadcasts() bool  { return m.Data2.Broadcasts }

func (m *TLDialogFilter) SetBots(v bool) { m.Data2.Bots = v }
func (m *TLDialogFilter) GetBots() bool  { return m.Data2.Bots }

func (m *TLDialogFilter) SetExcludeMuted(v bool) { m.Data2.ExcludeMuted = v }
func (m *TLDialogFilter) GetExcludeMuted() bool  { return m.Data2.ExcludeMuted }

func (m *TLDialogFilter) SetExcludeRead(v bool) { m.Data2.ExcludeRead = v }
func (m *TLDialogFilter) GetExcludeRead() bool  { return m.Data2.ExcludeRead }

func (m *TLDialogFilter) SetExcludeArchived(v bool) { m.Data2.ExcludeArchived = v }
func (m *TLDialogFilter) GetExcludeArchived() bool  { return m.Data2.ExcludeArchived }

func (m *TLDialogFilter) SetId(v int32) { m.Data2.Id = v }
func (m *TLDialogFilter) GetId() int32  { return m.Data2.Id }

func (m *TLDialogFilter) SetTitle(v string) { m.Data2.Title = v }
func (m *TLDialogFilter) GetTitle() string  { return m.Data2.Title }

func (m *TLDialogFilter) SetEmoticon(v string) { m.Data2.Emoticon = v }
func (m *TLDialogFilter) GetEmoticon() string  { return m.Data2.Emoticon }

func (m *TLDialogFilter) SetPinnedPeers(v []*InputPeer) { m.Data2.PinnedPeers = v }
func (m *TLDialogFilter) GetPinnedPeers() []*InputPeer  { return m.Data2.PinnedPeers }

func (m *TLDialogFilter) SetIncludePeers(v []*InputPeer) { m.Data2.IncludePeers = v }
func (m *TLDialogFilter) GetIncludePeers() []*InputPeer  { return m.Data2.IncludePeers }

func (m *TLDialogFilter) SetExcludePeers(v []*InputPeer) { m.Data2.ExcludePeers = v }
func (m *TLDialogFilter) GetExcludePeers() []*InputPeer  { return m.Data2.ExcludePeers }

func (m *TLDialogFilterSuggested) SetFilter(v *DialogFilter) { m.Data2.Filter = v }
func (m *TLDialogFilterSuggested) GetFilter() *DialogFilter  { return m.Data2.Filter }

func (m *TLDialogFilterSuggested) SetDescription(v string) { m.Data2.Description = v }
func (m *TLDialogFilterSuggested) GetDescription() string  { return m.Data2.Description }

func (m *TLDialogPeer) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLDialogPeer) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLDialogPeerFolder) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLDialogPeerFolder) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLDisabledFeature) SetFeature(v string) { m.Data2.Feature = v }
func (m *TLDisabledFeature) GetFeature() string  { return m.Data2.Feature }

func (m *TLDisabledFeature) SetDescription(v string) { m.Data2.Description = v }
func (m *TLDisabledFeature) GetDescription() string  { return m.Data2.Description }

func (m *TLDocumentEmpty) SetId(v int64) { m.Data2.Id = v }
func (m *TLDocumentEmpty) GetId() int64  { return m.Data2.Id }

func (m *TLDocument) SetId(v int64) { m.Data2.Id = v }
func (m *TLDocument) GetId() int64  { return m.Data2.Id }

func (m *TLDocument) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLDocument) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLDocument) SetDate(v int32) { m.Data2.Date = v }
func (m *TLDocument) GetDate() int32  { return m.Data2.Date }

func (m *TLDocument) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLDocument) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLDocument) SetSize_(v int32) { m.Data2.Size_ = v }
func (m *TLDocument) GetSize_() int32  { return m.Data2.Size_ }

func (m *TLDocument) SetThumb(v *PhotoSize) { m.Data2.Thumb = v }
func (m *TLDocument) GetThumb() *PhotoSize  { return m.Data2.Thumb }

func (m *TLDocument) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLDocument) GetDcId() int32  { return m.Data2.DcId }

func (m *TLDocument) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLDocument) GetVersion() int32  { return m.Data2.Version }

func (m *TLDocument) SetAttributes(v []*DocumentAttribute) { m.Data2.Attributes = v }
func (m *TLDocument) GetAttributes() []*DocumentAttribute  { return m.Data2.Attributes }

func (m *TLDocument) SetFileReference(v []byte) { m.Data2.FileReference = v }
func (m *TLDocument) GetFileReference() []byte  { return m.Data2.FileReference }

func (m *TLDocument) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLDocument) GetFlags() int32  { return m.Data2.Flags }

func (m *TLDocument) SetThumbs(v []*PhotoSize) { m.Data2.Thumbs = v }
func (m *TLDocument) GetThumbs() []*PhotoSize  { return m.Data2.Thumbs }

func (m *TLDocument) SetVideoThumbs(v []*VideoSize) { m.Data2.VideoThumbs = v }
func (m *TLDocument) GetVideoThumbs() []*VideoSize  { return m.Data2.VideoThumbs }

func (m *TLDocumentAttributeImageSize) SetW(v int32) { m.Data2.W = v }
func (m *TLDocumentAttributeImageSize) GetW() int32  { return m.Data2.W }

func (m *TLDocumentAttributeImageSize) SetH(v int32) { m.Data2.H = v }
func (m *TLDocumentAttributeImageSize) GetH() int32  { return m.Data2.H }

func (m *TLDocumentAttributeSticker) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLDocumentAttributeSticker) GetFlags() int32  { return m.Data2.Flags }

func (m *TLDocumentAttributeSticker) SetMask(v bool) { m.Data2.Mask = v }
func (m *TLDocumentAttributeSticker) GetMask() bool  { return m.Data2.Mask }

func (m *TLDocumentAttributeSticker) SetAlt(v string) { m.Data2.Alt = v }
func (m *TLDocumentAttributeSticker) GetAlt() string  { return m.Data2.Alt }

func (m *TLDocumentAttributeSticker) SetStickerset(v *InputStickerSet) { m.Data2.Stickerset = v }
func (m *TLDocumentAttributeSticker) GetStickerset() *InputStickerSet  { return m.Data2.Stickerset }

func (m *TLDocumentAttributeSticker) SetMaskCoords(v *MaskCoords) { m.Data2.MaskCoords = v }
func (m *TLDocumentAttributeSticker) GetMaskCoords() *MaskCoords  { return m.Data2.MaskCoords }

func (m *TLDocumentAttributeVideo) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLDocumentAttributeVideo) GetFlags() int32  { return m.Data2.Flags }

func (m *TLDocumentAttributeVideo) SetRoundMessage(v bool) { m.Data2.RoundMessage = v }
func (m *TLDocumentAttributeVideo) GetRoundMessage() bool  { return m.Data2.RoundMessage }

func (m *TLDocumentAttributeVideo) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLDocumentAttributeVideo) GetDuration() int32  { return m.Data2.Duration }

func (m *TLDocumentAttributeVideo) SetW(v int32) { m.Data2.W = v }
func (m *TLDocumentAttributeVideo) GetW() int32  { return m.Data2.W }

func (m *TLDocumentAttributeVideo) SetH(v int32) { m.Data2.H = v }
func (m *TLDocumentAttributeVideo) GetH() int32  { return m.Data2.H }

func (m *TLDocumentAttributeAudio) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLDocumentAttributeAudio) GetFlags() int32  { return m.Data2.Flags }

func (m *TLDocumentAttributeAudio) SetVoice(v bool) { m.Data2.Voice = v }
func (m *TLDocumentAttributeAudio) GetVoice() bool  { return m.Data2.Voice }

func (m *TLDocumentAttributeAudio) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLDocumentAttributeAudio) GetDuration() int32  { return m.Data2.Duration }

func (m *TLDocumentAttributeAudio) SetTitle(v string) { m.Data2.Title = v }
func (m *TLDocumentAttributeAudio) GetTitle() string  { return m.Data2.Title }

func (m *TLDocumentAttributeAudio) SetPerformer(v string) { m.Data2.Performer = v }
func (m *TLDocumentAttributeAudio) GetPerformer() string  { return m.Data2.Performer }

func (m *TLDocumentAttributeAudio) SetWaveform(v []byte) { m.Data2.Waveform = v }
func (m *TLDocumentAttributeAudio) GetWaveform() []byte  { return m.Data2.Waveform }

func (m *TLDocumentAttributeFilename) SetFileName(v string) { m.Data2.FileName = v }
func (m *TLDocumentAttributeFilename) GetFileName() string  { return m.Data2.FileName }

func (m *TLDraftMessageEmpty) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLDraftMessageEmpty) GetFlags() int32  { return m.Data2.Flags }

func (m *TLDraftMessageEmpty) SetDate(v int32) { m.Data2.Date = v }
func (m *TLDraftMessageEmpty) GetDate() int32  { return m.Data2.Date }

func (m *TLDraftMessage) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLDraftMessage) GetFlags() int32  { return m.Data2.Flags }

func (m *TLDraftMessage) SetNoWebpage(v bool) { m.Data2.NoWebpage = v }
func (m *TLDraftMessage) GetNoWebpage() bool  { return m.Data2.NoWebpage }

func (m *TLDraftMessage) SetReplyToMsgId(v int32) { m.Data2.ReplyToMsgId = v }
func (m *TLDraftMessage) GetReplyToMsgId() int32  { return m.Data2.ReplyToMsgId }

func (m *TLDraftMessage) SetMessage(v string) { m.Data2.Message = v }
func (m *TLDraftMessage) GetMessage() string  { return m.Data2.Message }

func (m *TLDraftMessage) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLDraftMessage) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLDraftMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLDraftMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLEmojiKeyword) SetKeyword(v string) { m.Data2.Keyword = v }
func (m *TLEmojiKeyword) GetKeyword() string  { return m.Data2.Keyword }

func (m *TLEmojiKeyword) SetEmoticons(v []string) { m.Data2.Emoticons = v }
func (m *TLEmojiKeyword) GetEmoticons() []string  { return m.Data2.Emoticons }

func (m *TLEmojiKeywordDeleted) SetKeyword(v string) { m.Data2.Keyword = v }
func (m *TLEmojiKeywordDeleted) GetKeyword() string  { return m.Data2.Keyword }

func (m *TLEmojiKeywordDeleted) SetEmoticons(v []string) { m.Data2.Emoticons = v }
func (m *TLEmojiKeywordDeleted) GetEmoticons() []string  { return m.Data2.Emoticons }

func (m *TLEmojiKeywordsDifference) SetLangCode(v string) { m.Data2.LangCode = v }
func (m *TLEmojiKeywordsDifference) GetLangCode() string  { return m.Data2.LangCode }

func (m *TLEmojiKeywordsDifference) SetFromVersion(v int32) { m.Data2.FromVersion = v }
func (m *TLEmojiKeywordsDifference) GetFromVersion() int32  { return m.Data2.FromVersion }

func (m *TLEmojiKeywordsDifference) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLEmojiKeywordsDifference) GetVersion() int32  { return m.Data2.Version }

func (m *TLEmojiKeywordsDifference) SetKeywords(v []*EmojiKeyword) { m.Data2.Keywords = v }
func (m *TLEmojiKeywordsDifference) GetKeywords() []*EmojiKeyword  { return m.Data2.Keywords }

func (m *TLEmojiLanguage) SetLangCode(v string) { m.Data2.LangCode = v }
func (m *TLEmojiLanguage) GetLangCode() string  { return m.Data2.LangCode }

func (m *TLEmojiURL) SetUrl(v string) { m.Data2.Url = v }
func (m *TLEmojiURL) GetUrl() string  { return m.Data2.Url }

func (m *TLEncryptedChatEmpty) SetId(v int32) { m.Data2.Id = v }
func (m *TLEncryptedChatEmpty) GetId() int32  { return m.Data2.Id }

func (m *TLEncryptedChatWaiting) SetId(v int32) { m.Data2.Id = v }
func (m *TLEncryptedChatWaiting) GetId() int32  { return m.Data2.Id }

func (m *TLEncryptedChatWaiting) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLEncryptedChatWaiting) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLEncryptedChatWaiting) SetDate(v int32) { m.Data2.Date = v }
func (m *TLEncryptedChatWaiting) GetDate() int32  { return m.Data2.Date }

func (m *TLEncryptedChatWaiting) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLEncryptedChatWaiting) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLEncryptedChatWaiting) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLEncryptedChatWaiting) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLEncryptedChatRequested) SetId(v int32) { m.Data2.Id = v }
func (m *TLEncryptedChatRequested) GetId() int32  { return m.Data2.Id }

func (m *TLEncryptedChatRequested) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLEncryptedChatRequested) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLEncryptedChatRequested) SetDate(v int32) { m.Data2.Date = v }
func (m *TLEncryptedChatRequested) GetDate() int32  { return m.Data2.Date }

func (m *TLEncryptedChatRequested) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLEncryptedChatRequested) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLEncryptedChatRequested) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLEncryptedChatRequested) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLEncryptedChatRequested) SetGA(v []byte) { m.Data2.GA = v }
func (m *TLEncryptedChatRequested) GetGA() []byte  { return m.Data2.GA }

func (m *TLEncryptedChatRequested) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLEncryptedChatRequested) GetFlags() int32  { return m.Data2.Flags }

func (m *TLEncryptedChatRequested) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLEncryptedChatRequested) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLEncryptedChat) SetId(v int32) { m.Data2.Id = v }
func (m *TLEncryptedChat) GetId() int32  { return m.Data2.Id }

func (m *TLEncryptedChat) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLEncryptedChat) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLEncryptedChat) SetDate(v int32) { m.Data2.Date = v }
func (m *TLEncryptedChat) GetDate() int32  { return m.Data2.Date }

func (m *TLEncryptedChat) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLEncryptedChat) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLEncryptedChat) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLEncryptedChat) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLEncryptedChat) SetGAOrB(v []byte) { m.Data2.GAOrB = v }
func (m *TLEncryptedChat) GetGAOrB() []byte  { return m.Data2.GAOrB }

func (m *TLEncryptedChat) SetKeyFingerprint(v int64) { m.Data2.KeyFingerprint = v }
func (m *TLEncryptedChat) GetKeyFingerprint() int64  { return m.Data2.KeyFingerprint }

func (m *TLEncryptedChatDiscarded) SetId(v int32) { m.Data2.Id = v }
func (m *TLEncryptedChatDiscarded) GetId() int32  { return m.Data2.Id }

func (m *TLEncryptedFile) SetId(v int64) { m.Data2.Id = v }
func (m *TLEncryptedFile) GetId() int64  { return m.Data2.Id }

func (m *TLEncryptedFile) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLEncryptedFile) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLEncryptedFile) SetSize_(v int32) { m.Data2.Size_ = v }
func (m *TLEncryptedFile) GetSize_() int32  { return m.Data2.Size_ }

func (m *TLEncryptedFile) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLEncryptedFile) GetDcId() int32  { return m.Data2.DcId }

func (m *TLEncryptedFile) SetKeyFingerprint(v int32) { m.Data2.KeyFingerprint = v }
func (m *TLEncryptedFile) GetKeyFingerprint() int32  { return m.Data2.KeyFingerprint }

func (m *TLEncryptedMessage) SetRandomId(v int64) { m.Data2.RandomId = v }
func (m *TLEncryptedMessage) GetRandomId() int64  { return m.Data2.RandomId }

func (m *TLEncryptedMessage) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLEncryptedMessage) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLEncryptedMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLEncryptedMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLEncryptedMessage) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLEncryptedMessage) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLEncryptedMessage) SetFile(v *EncryptedFile) { m.Data2.File = v }
func (m *TLEncryptedMessage) GetFile() *EncryptedFile  { return m.Data2.File }

func (m *TLEncryptedMessageService) SetRandomId(v int64) { m.Data2.RandomId = v }
func (m *TLEncryptedMessageService) GetRandomId() int64  { return m.Data2.RandomId }

func (m *TLEncryptedMessageService) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLEncryptedMessageService) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLEncryptedMessageService) SetDate(v int32) { m.Data2.Date = v }
func (m *TLEncryptedMessageService) GetDate() int32  { return m.Data2.Date }

func (m *TLEncryptedMessageService) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLEncryptedMessageService) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLError) SetCode(v int32) { m.Data2.Code = v }
func (m *TLError) GetCode() int32  { return m.Data2.Code }

func (m *TLError) SetText(v string) { m.Data2.Text = v }
func (m *TLError) GetText() string  { return m.Data2.Text }

func (m *TLChatInviteExported) SetLink(v string) { m.Data2.Link = v }
func (m *TLChatInviteExported) GetLink() string  { return m.Data2.Link }

func (m *TLExportedMessageLink) SetLink(v string) { m.Data2.Link = v }
func (m *TLExportedMessageLink) GetLink() string  { return m.Data2.Link }

func (m *TLExportedMessageLink) SetHtml(v string) { m.Data2.Html = v }
func (m *TLExportedMessageLink) GetHtml() string  { return m.Data2.Html }

func (m *TLFileHash) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLFileHash) GetOffset() int32  { return m.Data2.Offset }

func (m *TLFileHash) SetLimit(v int32) { m.Data2.Limit = v }
func (m *TLFileHash) GetLimit() int32  { return m.Data2.Limit }

func (m *TLFileHash) SetHash(v []byte) { m.Data2.Hash = v }
func (m *TLFileHash) GetHash() []byte  { return m.Data2.Hash }

func (m *TLFileLocationUnavailable) SetVolumeId(v int64) { m.Data2.VolumeId = v }
func (m *TLFileLocationUnavailable) GetVolumeId() int64  { return m.Data2.VolumeId }

func (m *TLFileLocationUnavailable) SetLocalId(v int32) { m.Data2.LocalId = v }
func (m *TLFileLocationUnavailable) GetLocalId() int32  { return m.Data2.LocalId }

func (m *TLFileLocationUnavailable) SetSecret(v int64) { m.Data2.Secret = v }
func (m *TLFileLocationUnavailable) GetSecret() int64  { return m.Data2.Secret }

func (m *TLFileLocation) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLFileLocation) GetDcId() int32  { return m.Data2.DcId }

func (m *TLFileLocation) SetVolumeId(v int64) { m.Data2.VolumeId = v }
func (m *TLFileLocation) GetVolumeId() int64  { return m.Data2.VolumeId }

func (m *TLFileLocation) SetLocalId(v int32) { m.Data2.LocalId = v }
func (m *TLFileLocation) GetLocalId() int32  { return m.Data2.LocalId }

func (m *TLFileLocation) SetSecret(v int64) { m.Data2.Secret = v }
func (m *TLFileLocation) GetSecret() int64  { return m.Data2.Secret }

func (m *TLFileLocation) SetFileReference(v []byte) { m.Data2.FileReference = v }
func (m *TLFileLocation) GetFileReference() []byte  { return m.Data2.FileReference }

func (m *TLFileLocationToBeDeprecated) SetVolumeId(v int64) { m.Data2.VolumeId = v }
func (m *TLFileLocationToBeDeprecated) GetVolumeId() int64  { return m.Data2.VolumeId }

func (m *TLFileLocationToBeDeprecated) SetLocalId(v int32) { m.Data2.LocalId = v }
func (m *TLFileLocationToBeDeprecated) GetLocalId() int32  { return m.Data2.LocalId }

func (m *TLFolder) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLFolder) GetFlags() int32  { return m.Data2.Flags }

func (m *TLFolder) SetAutofillNewBroadcasts(v bool) { m.Data2.AutofillNewBroadcasts = v }
func (m *TLFolder) GetAutofillNewBroadcasts() bool  { return m.Data2.AutofillNewBroadcasts }

func (m *TLFolder) SetAutofillPublicGroups(v bool) { m.Data2.AutofillPublicGroups = v }
func (m *TLFolder) GetAutofillPublicGroups() bool  { return m.Data2.AutofillPublicGroups }

func (m *TLFolder) SetAutofillNewCorrespondents(v bool) { m.Data2.AutofillNewCorrespondents = v }
func (m *TLFolder) GetAutofillNewCorrespondents() bool  { return m.Data2.AutofillNewCorrespondents }

func (m *TLFolder) SetId(v int32) { m.Data2.Id = v }
func (m *TLFolder) GetId() int32  { return m.Data2.Id }

func (m *TLFolder) SetTitle(v string) { m.Data2.Title = v }
func (m *TLFolder) GetTitle() string  { return m.Data2.Title }

func (m *TLFolder) SetPhoto(v *ChatPhoto) { m.Data2.Photo = v }
func (m *TLFolder) GetPhoto() *ChatPhoto  { return m.Data2.Photo }

func (m *TLFolderPeer) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLFolderPeer) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLFolderPeer) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLFolderPeer) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLFoundGif) SetUrl(v string) { m.Data2.Url = v }
func (m *TLFoundGif) GetUrl() string  { return m.Data2.Url }

func (m *TLFoundGif) SetThumbUrl(v string) { m.Data2.ThumbUrl = v }
func (m *TLFoundGif) GetThumbUrl() string  { return m.Data2.ThumbUrl }

func (m *TLFoundGif) SetContentUrl(v string) { m.Data2.ContentUrl = v }
func (m *TLFoundGif) GetContentUrl() string  { return m.Data2.ContentUrl }

func (m *TLFoundGif) SetContentType(v string) { m.Data2.ContentType = v }
func (m *TLFoundGif) GetContentType() string  { return m.Data2.ContentType }

func (m *TLFoundGif) SetW(v int32) { m.Data2.W = v }
func (m *TLFoundGif) GetW() int32  { return m.Data2.W }

func (m *TLFoundGif) SetH(v int32) { m.Data2.H = v }
func (m *TLFoundGif) GetH() int32  { return m.Data2.H }

func (m *TLFoundGifCached) SetUrl(v string) { m.Data2.Url = v }
func (m *TLFoundGifCached) GetUrl() string  { return m.Data2.Url }

func (m *TLFoundGifCached) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLFoundGifCached) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLFoundGifCached) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLFoundGifCached) GetDocument() *Document  { return m.Data2.Document }

func (m *TLFutureSalt) SetValidSince(v int32) { m.Data2.ValidSince = v }
func (m *TLFutureSalt) GetValidSince() int32  { return m.Data2.ValidSince }

func (m *TLFutureSalt) SetValidUntil(v int32) { m.Data2.ValidUntil = v }
func (m *TLFutureSalt) GetValidUntil() int32  { return m.Data2.ValidUntil }

func (m *TLFutureSalt) SetSalt(v int64) { m.Data2.Salt = v }
func (m *TLFutureSalt) GetSalt() int64  { return m.Data2.Salt }

func (m *TLFutureSalts) SetReqMsgId(v int64) { m.Data2.ReqMsgId = v }
func (m *TLFutureSalts) GetReqMsgId() int64  { return m.Data2.ReqMsgId }

func (m *TLFutureSalts) SetNow(v int32) { m.Data2.Now = v }
func (m *TLFutureSalts) GetNow() int32  { return m.Data2.Now }

func (m *TLFutureSalts) SetSalts(v []*TLFutureSalt) { m.Data2.Salts = v }
func (m *TLFutureSalts) GetSalts() []*TLFutureSalt  { return m.Data2.Salts }

func (m *TLGame) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLGame) GetFlags() int32  { return m.Data2.Flags }

func (m *TLGame) SetId(v int64) { m.Data2.Id = v }
func (m *TLGame) GetId() int64  { return m.Data2.Id }

func (m *TLGame) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLGame) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLGame) SetShortName(v string) { m.Data2.ShortName = v }
func (m *TLGame) GetShortName() string  { return m.Data2.ShortName }

func (m *TLGame) SetTitle(v string) { m.Data2.Title = v }
func (m *TLGame) GetTitle() string  { return m.Data2.Title }

func (m *TLGame) SetDescription(v string) { m.Data2.Description = v }
func (m *TLGame) GetDescription() string  { return m.Data2.Description }

func (m *TLGame) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLGame) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLGame) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLGame) GetDocument() *Document  { return m.Data2.Document }

func (m *TLGeoPoint) SetLong(v float64) { m.Data2.Long = v }
func (m *TLGeoPoint) GetLong() float64  { return m.Data2.Long }

func (m *TLGeoPoint) SetLat(v float64) { m.Data2.Lat = v }
func (m *TLGeoPoint) GetLat() float64  { return m.Data2.Lat }

func (m *TLGeoPoint) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLGeoPoint) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLGeoPoint) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLGeoPoint) GetFlags() int32  { return m.Data2.Flags }

func (m *TLGeoPoint) SetAccuracyRadius(v int32) { m.Data2.AccuracyRadius = v }
func (m *TLGeoPoint) GetAccuracyRadius() int32  { return m.Data2.AccuracyRadius }

func (m *TLGlobalPrivacySettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLGlobalPrivacySettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLGlobalPrivacySettings) SetArchiveAndMuteNewNoncontactPeers(v *Bool) {
	m.Data2.ArchiveAndMuteNewNoncontactPeers = v
}
func (m *TLGlobalPrivacySettings) GetArchiveAndMuteNewNoncontactPeers() *Bool {
	return m.Data2.ArchiveAndMuteNewNoncontactPeers
}

func (m *TLGroupCallDiscarded) SetId(v int64) { m.Data2.Id = v }
func (m *TLGroupCallDiscarded) GetId() int64  { return m.Data2.Id }

func (m *TLGroupCallDiscarded) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLGroupCallDiscarded) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLGroupCallDiscarded) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLGroupCallDiscarded) GetDuration() int32  { return m.Data2.Duration }

func (m *TLGroupCall) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLGroupCall) GetFlags() int32  { return m.Data2.Flags }

func (m *TLGroupCall) SetJoinMuted(v bool) { m.Data2.JoinMuted = v }
func (m *TLGroupCall) GetJoinMuted() bool  { return m.Data2.JoinMuted }

func (m *TLGroupCall) SetCanChangeJoinMuted(v bool) { m.Data2.CanChangeJoinMuted = v }
func (m *TLGroupCall) GetCanChangeJoinMuted() bool  { return m.Data2.CanChangeJoinMuted }

func (m *TLGroupCall) SetId(v int64) { m.Data2.Id = v }
func (m *TLGroupCall) GetId() int64  { return m.Data2.Id }

func (m *TLGroupCall) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLGroupCall) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLGroupCall) SetParticipantsCount(v int32) { m.Data2.ParticipantsCount = v }
func (m *TLGroupCall) GetParticipantsCount() int32  { return m.Data2.ParticipantsCount }

func (m *TLGroupCall) SetParams(v *DataJSON) { m.Data2.Params = v }
func (m *TLGroupCall) GetParams() *DataJSON  { return m.Data2.Params }

func (m *TLGroupCall) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLGroupCall) GetVersion() int32  { return m.Data2.Version }

func (m *TLGroupCallParticipant) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLGroupCallParticipant) GetFlags() int32  { return m.Data2.Flags }

func (m *TLGroupCallParticipant) SetMuted(v bool) { m.Data2.Muted = v }
func (m *TLGroupCallParticipant) GetMuted() bool  { return m.Data2.Muted }

func (m *TLGroupCallParticipant) SetLeft(v bool) { m.Data2.Left = v }
func (m *TLGroupCallParticipant) GetLeft() bool  { return m.Data2.Left }

func (m *TLGroupCallParticipant) SetCanSelfUnmute(v bool) { m.Data2.CanSelfUnmute = v }
func (m *TLGroupCallParticipant) GetCanSelfUnmute() bool  { return m.Data2.CanSelfUnmute }

func (m *TLGroupCallParticipant) SetJustJoined(v bool) { m.Data2.JustJoined = v }
func (m *TLGroupCallParticipant) GetJustJoined() bool  { return m.Data2.JustJoined }

func (m *TLGroupCallParticipant) SetVersioned(v bool) { m.Data2.Versioned = v }
func (m *TLGroupCallParticipant) GetVersioned() bool  { return m.Data2.Versioned }

func (m *TLGroupCallParticipant) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLGroupCallParticipant) GetUserId() int32  { return m.Data2.UserId }

func (m *TLGroupCallParticipant) SetDate(v int32) { m.Data2.Date = v }
func (m *TLGroupCallParticipant) GetDate() int32  { return m.Data2.Date }

func (m *TLGroupCallParticipant) SetActiveDate(v int32) { m.Data2.ActiveDate = v }
func (m *TLGroupCallParticipant) GetActiveDate() int32  { return m.Data2.ActiveDate }

func (m *TLGroupCallParticipant) SetSource(v int32) { m.Data2.Source = v }
func (m *TLGroupCallParticipant) GetSource() int32  { return m.Data2.Source }

func (m *TLHelpAppChangelog) SetText(v string) { m.Data2.Text = v }
func (m *TLHelpAppChangelog) GetText() string  { return m.Data2.Text }

func (m *TLHelpAppUpdate) SetId(v int32) { m.Data2.Id = v }
func (m *TLHelpAppUpdate) GetId() int32  { return m.Data2.Id }

func (m *TLHelpAppUpdate) SetCritical(v *Bool) { m.Data2.Critical = v }
func (m *TLHelpAppUpdate) GetCritical() *Bool  { return m.Data2.Critical }

func (m *TLHelpAppUpdate) SetUrl(v string) { m.Data2.Url = v }
func (m *TLHelpAppUpdate) GetUrl() string  { return m.Data2.Url }

func (m *TLHelpAppUpdate) SetText(v string) { m.Data2.Text = v }
func (m *TLHelpAppUpdate) GetText() string  { return m.Data2.Text }

func (m *TLHelpAppUpdate) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLHelpAppUpdate) GetFlags() int32  { return m.Data2.Flags }

func (m *TLHelpAppUpdate) SetPopup(v bool) { m.Data2.Popup = v }
func (m *TLHelpAppUpdate) GetPopup() bool  { return m.Data2.Popup }

func (m *TLHelpAppUpdate) SetVersion(v string) { m.Data2.Version = v }
func (m *TLHelpAppUpdate) GetVersion() string  { return m.Data2.Version }

func (m *TLHelpAppUpdate) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLHelpAppUpdate) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLHelpAppUpdate) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLHelpAppUpdate) GetDocument() *Document  { return m.Data2.Document }

func (m *TLHelpConfigSimple) SetDate(v int32) { m.Data2.Date = v }
func (m *TLHelpConfigSimple) GetDate() int32  { return m.Data2.Date }

func (m *TLHelpConfigSimple) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLHelpConfigSimple) GetExpires() int32  { return m.Data2.Expires }

func (m *TLHelpConfigSimple) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLHelpConfigSimple) GetDcId() int32  { return m.Data2.DcId }

func (m *TLHelpConfigSimple) SetIpPortList(v []*IpPort) { m.Data2.IpPortList = v }
func (m *TLHelpConfigSimple) GetIpPortList() []*IpPort  { return m.Data2.IpPortList }

func (m *TLHelpConfigSimple) SetRules(v []*AccessPointRule) { m.Data2.Rules = v }
func (m *TLHelpConfigSimple) GetRules() []*AccessPointRule  { return m.Data2.Rules }

func (m *TLHelpCountriesList) SetCountries(v []*Help_Country) { m.Data2.Countries = v }
func (m *TLHelpCountriesList) GetCountries() []*Help_Country  { return m.Data2.Countries }

func (m *TLHelpCountriesList) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLHelpCountriesList) GetHash() int32  { return m.Data2.Hash }

func (m *TLHelpCountry) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLHelpCountry) GetFlags() int32  { return m.Data2.Flags }

func (m *TLHelpCountry) SetHidden(v bool) { m.Data2.Hidden = v }
func (m *TLHelpCountry) GetHidden() bool  { return m.Data2.Hidden }

func (m *TLHelpCountry) SetIso2(v string) { m.Data2.Iso2 = v }
func (m *TLHelpCountry) GetIso2() string  { return m.Data2.Iso2 }

func (m *TLHelpCountry) SetDefaultName(v string) { m.Data2.DefaultName = v }
func (m *TLHelpCountry) GetDefaultName() string  { return m.Data2.DefaultName }

func (m *TLHelpCountry) SetName(v string) { m.Data2.Name = v }
func (m *TLHelpCountry) GetName() string  { return m.Data2.Name }

func (m *TLHelpCountry) SetCountryCodes(v []*Help_CountryCode) { m.Data2.CountryCodes = v }
func (m *TLHelpCountry) GetCountryCodes() []*Help_CountryCode  { return m.Data2.CountryCodes }

func (m *TLHelpCountryCode) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLHelpCountryCode) GetFlags() int32  { return m.Data2.Flags }

func (m *TLHelpCountryCode) SetCountryCode(v string) { m.Data2.CountryCode = v }
func (m *TLHelpCountryCode) GetCountryCode() string  { return m.Data2.CountryCode }

func (m *TLHelpCountryCode) SetPrefixes(v []string) { m.Data2.Prefixes = v }
func (m *TLHelpCountryCode) GetPrefixes() []string  { return m.Data2.Prefixes }

func (m *TLHelpCountryCode) SetPatterns(v []string) { m.Data2.Patterns = v }
func (m *TLHelpCountryCode) GetPatterns() []string  { return m.Data2.Patterns }

func (m *TLHelpDeepLinkInfo) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLHelpDeepLinkInfo) GetFlags() int32  { return m.Data2.Flags }

func (m *TLHelpDeepLinkInfo) SetUpdateApp(v bool) { m.Data2.UpdateApp = v }
func (m *TLHelpDeepLinkInfo) GetUpdateApp() bool  { return m.Data2.UpdateApp }

func (m *TLHelpDeepLinkInfo) SetMessage(v string) { m.Data2.Message = v }
func (m *TLHelpDeepLinkInfo) GetMessage() string  { return m.Data2.Message }

func (m *TLHelpDeepLinkInfo) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLHelpDeepLinkInfo) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLHelpInviteText) SetMessage(v string) { m.Data2.Message = v }
func (m *TLHelpInviteText) GetMessage() string  { return m.Data2.Message }

func (m *TLHelpPassportConfig) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLHelpPassportConfig) GetHash() int32  { return m.Data2.Hash }

func (m *TLHelpPassportConfig) SetCountriesLangs(v *DataJSON) { m.Data2.CountriesLangs = v }
func (m *TLHelpPassportConfig) GetCountriesLangs() *DataJSON  { return m.Data2.CountriesLangs }

func (m *TLHelpPromoDataEmpty) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLHelpPromoDataEmpty) GetExpires() int32  { return m.Data2.Expires }

func (m *TLHelpPromoData) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLHelpPromoData) GetFlags() int32  { return m.Data2.Flags }

func (m *TLHelpPromoData) SetProxy(v bool) { m.Data2.Proxy = v }
func (m *TLHelpPromoData) GetProxy() bool  { return m.Data2.Proxy }

func (m *TLHelpPromoData) SetPsa(v bool) { m.Data2.Psa = v }
func (m *TLHelpPromoData) GetPsa() bool  { return m.Data2.Psa }

func (m *TLHelpPromoData) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLHelpPromoData) GetExpires() int32  { return m.Data2.Expires }

func (m *TLHelpPromoData) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLHelpPromoData) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLHelpPromoData) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLHelpPromoData) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLHelpPromoData) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLHelpPromoData) GetUsers() []*User  { return m.Data2.Users }

func (m *TLHelpPromoData) SetPsaMessage(v string) { m.Data2.PsaMessage = v }
func (m *TLHelpPromoData) GetPsaMessage() string  { return m.Data2.PsaMessage }

func (m *TLHelpPromoData) SetPsaType(v string) { m.Data2.PsaType = v }
func (m *TLHelpPromoData) GetPsaType() string  { return m.Data2.PsaType }

func (m *TLHelpProxyDataEmpty) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLHelpProxyDataEmpty) GetExpires() int32  { return m.Data2.Expires }

func (m *TLHelpProxyDataPromo) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLHelpProxyDataPromo) GetExpires() int32  { return m.Data2.Expires }

func (m *TLHelpProxyDataPromo) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLHelpProxyDataPromo) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLHelpProxyDataPromo) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLHelpProxyDataPromo) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLHelpProxyDataPromo) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLHelpProxyDataPromo) GetUsers() []*User  { return m.Data2.Users }

func (m *TLHelpRecentMeUrls) SetUrls(v []*RecentMeUrl) { m.Data2.Urls = v }
func (m *TLHelpRecentMeUrls) GetUrls() []*RecentMeUrl  { return m.Data2.Urls }

func (m *TLHelpRecentMeUrls) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLHelpRecentMeUrls) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLHelpRecentMeUrls) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLHelpRecentMeUrls) GetUsers() []*User  { return m.Data2.Users }

func (m *TLHelpSupport) SetPhoneNumber(v string) { m.Data2.PhoneNumber = v }
func (m *TLHelpSupport) GetPhoneNumber() string  { return m.Data2.PhoneNumber }

func (m *TLHelpSupport) SetUser(v *User) { m.Data2.User = v }
func (m *TLHelpSupport) GetUser() *User  { return m.Data2.User }

func (m *TLHelpSupportName) SetName(v string) { m.Data2.Name = v }
func (m *TLHelpSupportName) GetName() string  { return m.Data2.Name }

func (m *TLHelpTermsOfService) SetText(v string) { m.Data2.Text = v }
func (m *TLHelpTermsOfService) GetText() string  { return m.Data2.Text }

func (m *TLHelpTermsOfService) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLHelpTermsOfService) GetFlags() int32  { return m.Data2.Flags }

func (m *TLHelpTermsOfService) SetPopup(v bool) { m.Data2.Popup = v }
func (m *TLHelpTermsOfService) GetPopup() bool  { return m.Data2.Popup }

func (m *TLHelpTermsOfService) SetId(v *DataJSON) { m.Data2.Id = v }
func (m *TLHelpTermsOfService) GetId() *DataJSON  { return m.Data2.Id }

func (m *TLHelpTermsOfService) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLHelpTermsOfService) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLHelpTermsOfService) SetMinAgeConfirm(v int32) { m.Data2.MinAgeConfirm = v }
func (m *TLHelpTermsOfService) GetMinAgeConfirm() int32  { return m.Data2.MinAgeConfirm }

func (m *TLHelpTermsOfServiceUpdateEmpty) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLHelpTermsOfServiceUpdateEmpty) GetExpires() int32  { return m.Data2.Expires }

func (m *TLHelpTermsOfServiceUpdate) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLHelpTermsOfServiceUpdate) GetExpires() int32  { return m.Data2.Expires }

func (m *TLHelpTermsOfServiceUpdate) SetTermsOfService(v *Help_TermsOfService) {
	m.Data2.TermsOfService = v
}
func (m *TLHelpTermsOfServiceUpdate) GetTermsOfService() *Help_TermsOfService {
	return m.Data2.TermsOfService
}

func (m *TLHelpUserInfo) SetMessage(v string) { m.Data2.Message = v }
func (m *TLHelpUserInfo) GetMessage() string  { return m.Data2.Message }

func (m *TLHelpUserInfo) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLHelpUserInfo) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLHelpUserInfo) SetAuthor(v string) { m.Data2.Author = v }
func (m *TLHelpUserInfo) GetAuthor() string  { return m.Data2.Author }

func (m *TLHelpUserInfo) SetDate(v int32) { m.Data2.Date = v }
func (m *TLHelpUserInfo) GetDate() int32  { return m.Data2.Date }

func (m *TLHighScore) SetPos(v int32) { m.Data2.Pos = v }
func (m *TLHighScore) GetPos() int32  { return m.Data2.Pos }

func (m *TLHighScore) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLHighScore) GetUserId() int32  { return m.Data2.UserId }

func (m *TLHighScore) SetScore(v int32) { m.Data2.Score = v }
func (m *TLHighScore) GetScore() int32  { return m.Data2.Score }

func (m *TLHttpWait) SetMaxDelay(v int32) { m.Data2.MaxDelay = v }
func (m *TLHttpWait) GetMaxDelay() int32  { return m.Data2.MaxDelay }

func (m *TLHttpWait) SetWaitAfter(v int32) { m.Data2.WaitAfter = v }
func (m *TLHttpWait) GetWaitAfter() int32  { return m.Data2.WaitAfter }

func (m *TLHttpWait) SetMaxWait(v int32) { m.Data2.MaxWait = v }
func (m *TLHttpWait) GetMaxWait() int32  { return m.Data2.MaxWait }

func (m *TLImportedContact) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLImportedContact) GetUserId() int32  { return m.Data2.UserId }

func (m *TLImportedContact) SetClientId(v int64) { m.Data2.ClientId = v }
func (m *TLImportedContact) GetClientId() int64  { return m.Data2.ClientId }

func (m *TLInlineBotSwitchPM) SetText(v string) { m.Data2.Text = v }
func (m *TLInlineBotSwitchPM) GetText() string  { return m.Data2.Text }

func (m *TLInlineBotSwitchPM) SetStartParam(v string) { m.Data2.StartParam = v }
func (m *TLInlineBotSwitchPM) GetStartParam() string  { return m.Data2.StartParam }

func (m *TLInputAppEvent) SetTime(v float64) { m.Data2.Time = v }
func (m *TLInputAppEvent) GetTime() float64  { return m.Data2.Time }

func (m *TLInputAppEvent) SetType(v string) { m.Data2.Type = v }
func (m *TLInputAppEvent) GetType() string  { return m.Data2.Type }

func (m *TLInputAppEvent) SetPeer(v int64) { m.Data2.Peer = v }
func (m *TLInputAppEvent) GetPeer() int64  { return m.Data2.Peer }

func (m *TLInputAppEvent) SetData770656A871(v string) { m.Data2.Data770656A871 = v }
func (m *TLInputAppEvent) GetData770656A871() string  { return m.Data2.Data770656A871 }

func (m *TLInputAppEvent) SetData1D1B124590(v *JSONValue) { m.Data2.Data1D1B124590 = v }
func (m *TLInputAppEvent) GetData1D1B124590() *JSONValue  { return m.Data2.Data1D1B124590 }

func (m *TLInputBotInlineMessageMediaAuto) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputBotInlineMessageMediaAuto) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputBotInlineMessageMediaAuto) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputBotInlineMessageMediaAuto) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputBotInlineMessageMediaAuto) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageMediaAuto) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLInputBotInlineMessageMediaAuto) SetMessage(v string) { m.Data2.Message = v }
func (m *TLInputBotInlineMessageMediaAuto) GetMessage() string  { return m.Data2.Message }

func (m *TLInputBotInlineMessageMediaAuto) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLInputBotInlineMessageMediaAuto) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLInputBotInlineMessageText) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputBotInlineMessageText) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputBotInlineMessageText) SetNoWebpage(v bool) { m.Data2.NoWebpage = v }
func (m *TLInputBotInlineMessageText) GetNoWebpage() bool  { return m.Data2.NoWebpage }

func (m *TLInputBotInlineMessageText) SetMessage(v string) { m.Data2.Message = v }
func (m *TLInputBotInlineMessageText) GetMessage() string  { return m.Data2.Message }

func (m *TLInputBotInlineMessageText) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLInputBotInlineMessageText) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLInputBotInlineMessageText) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageText) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLInputBotInlineMessageMediaGeo) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputBotInlineMessageMediaGeo) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputBotInlineMessageMediaGeo) SetGeoPoint(v *InputGeoPoint) { m.Data2.GeoPoint = v }
func (m *TLInputBotInlineMessageMediaGeo) GetGeoPoint() *InputGeoPoint  { return m.Data2.GeoPoint }

func (m *TLInputBotInlineMessageMediaGeo) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageMediaGeo) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLInputBotInlineMessageMediaGeo) SetPeriod(v int32) { m.Data2.Period = v }
func (m *TLInputBotInlineMessageMediaGeo) GetPeriod() int32  { return m.Data2.Period }

func (m *TLInputBotInlineMessageMediaGeo) SetHeading(v int32) { m.Data2.Heading = v }
func (m *TLInputBotInlineMessageMediaGeo) GetHeading() int32  { return m.Data2.Heading }

func (m *TLInputBotInlineMessageMediaGeo) SetProximityNotificationRadius(v int32) {
	m.Data2.ProximityNotificationRadius = v
}
func (m *TLInputBotInlineMessageMediaGeo) GetProximityNotificationRadius() int32 {
	return m.Data2.ProximityNotificationRadius
}

func (m *TLInputBotInlineMessageMediaVenue) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputBotInlineMessageMediaVenue) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputBotInlineMessageMediaVenue) SetGeoPoint(v *InputGeoPoint) { m.Data2.GeoPoint = v }
func (m *TLInputBotInlineMessageMediaVenue) GetGeoPoint() *InputGeoPoint  { return m.Data2.GeoPoint }

func (m *TLInputBotInlineMessageMediaVenue) SetTitle(v string) { m.Data2.Title = v }
func (m *TLInputBotInlineMessageMediaVenue) GetTitle() string  { return m.Data2.Title }

func (m *TLInputBotInlineMessageMediaVenue) SetAddress(v string) { m.Data2.Address = v }
func (m *TLInputBotInlineMessageMediaVenue) GetAddress() string  { return m.Data2.Address }

func (m *TLInputBotInlineMessageMediaVenue) SetProvider(v string) { m.Data2.Provider = v }
func (m *TLInputBotInlineMessageMediaVenue) GetProvider() string  { return m.Data2.Provider }

func (m *TLInputBotInlineMessageMediaVenue) SetVenueId(v string) { m.Data2.VenueId = v }
func (m *TLInputBotInlineMessageMediaVenue) GetVenueId() string  { return m.Data2.VenueId }

func (m *TLInputBotInlineMessageMediaVenue) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageMediaVenue) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLInputBotInlineMessageMediaVenue) SetVenueType(v string) { m.Data2.VenueType = v }
func (m *TLInputBotInlineMessageMediaVenue) GetVenueType() string  { return m.Data2.VenueType }

func (m *TLInputBotInlineMessageMediaContact) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputBotInlineMessageMediaContact) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputBotInlineMessageMediaContact) SetPhoneNumber(v string) { m.Data2.PhoneNumber = v }
func (m *TLInputBotInlineMessageMediaContact) GetPhoneNumber() string  { return m.Data2.PhoneNumber }

func (m *TLInputBotInlineMessageMediaContact) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLInputBotInlineMessageMediaContact) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLInputBotInlineMessageMediaContact) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLInputBotInlineMessageMediaContact) GetLastName() string  { return m.Data2.LastName }

func (m *TLInputBotInlineMessageMediaContact) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageMediaContact) GetReplyMarkup() *ReplyMarkup {
	return m.Data2.ReplyMarkup
}

func (m *TLInputBotInlineMessageMediaContact) SetVcard(v string) { m.Data2.Vcard = v }
func (m *TLInputBotInlineMessageMediaContact) GetVcard() string  { return m.Data2.Vcard }

func (m *TLInputBotInlineMessageGame) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputBotInlineMessageGame) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputBotInlineMessageGame) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageGame) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLInputBotInlineMessageID) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLInputBotInlineMessageID) GetDcId() int32  { return m.Data2.DcId }

func (m *TLInputBotInlineMessageID) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputBotInlineMessageID) GetId() int64  { return m.Data2.Id }

func (m *TLInputBotInlineMessageID) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputBotInlineMessageID) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputBotInlineResult) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputBotInlineResult) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputBotInlineResult) SetId(v string) { m.Data2.Id = v }
func (m *TLInputBotInlineResult) GetId() string  { return m.Data2.Id }

func (m *TLInputBotInlineResult) SetType(v string) { m.Data2.Type = v }
func (m *TLInputBotInlineResult) GetType() string  { return m.Data2.Type }

func (m *TLInputBotInlineResult) SetTitle(v string) { m.Data2.Title = v }
func (m *TLInputBotInlineResult) GetTitle() string  { return m.Data2.Title }

func (m *TLInputBotInlineResult) SetDescription(v string) { m.Data2.Description = v }
func (m *TLInputBotInlineResult) GetDescription() string  { return m.Data2.Description }

func (m *TLInputBotInlineResult) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputBotInlineResult) GetUrl() string  { return m.Data2.Url }

func (m *TLInputBotInlineResult) SetThumbUrl(v string) { m.Data2.ThumbUrl = v }
func (m *TLInputBotInlineResult) GetThumbUrl() string  { return m.Data2.ThumbUrl }

func (m *TLInputBotInlineResult) SetContentUrl(v string) { m.Data2.ContentUrl = v }
func (m *TLInputBotInlineResult) GetContentUrl() string  { return m.Data2.ContentUrl }

func (m *TLInputBotInlineResult) SetContentType(v string) { m.Data2.ContentType = v }
func (m *TLInputBotInlineResult) GetContentType() string  { return m.Data2.ContentType }

func (m *TLInputBotInlineResult) SetW(v int32) { m.Data2.W = v }
func (m *TLInputBotInlineResult) GetW() int32  { return m.Data2.W }

func (m *TLInputBotInlineResult) SetH(v int32) { m.Data2.H = v }
func (m *TLInputBotInlineResult) GetH() int32  { return m.Data2.H }

func (m *TLInputBotInlineResult) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLInputBotInlineResult) GetDuration() int32  { return m.Data2.Duration }

func (m *TLInputBotInlineResult) SetSendMessage(v *InputBotInlineMessage) { m.Data2.SendMessage = v }
func (m *TLInputBotInlineResult) GetSendMessage() *InputBotInlineMessage  { return m.Data2.SendMessage }

func (m *TLInputBotInlineResult) SetThumb(v *InputWebDocument) { m.Data2.Thumb = v }
func (m *TLInputBotInlineResult) GetThumb() *InputWebDocument  { return m.Data2.Thumb }

func (m *TLInputBotInlineResult) SetContent(v *InputWebDocument) { m.Data2.Content = v }
func (m *TLInputBotInlineResult) GetContent() *InputWebDocument  { return m.Data2.Content }

func (m *TLInputBotInlineResultPhoto) SetId(v string) { m.Data2.Id = v }
func (m *TLInputBotInlineResultPhoto) GetId() string  { return m.Data2.Id }

func (m *TLInputBotInlineResultPhoto) SetType(v string) { m.Data2.Type = v }
func (m *TLInputBotInlineResultPhoto) GetType() string  { return m.Data2.Type }

func (m *TLInputBotInlineResultPhoto) SetPhoto(v *InputPhoto) { m.Data2.Photo = v }
func (m *TLInputBotInlineResultPhoto) GetPhoto() *InputPhoto  { return m.Data2.Photo }

func (m *TLInputBotInlineResultPhoto) SetSendMessage(v *InputBotInlineMessage) {
	m.Data2.SendMessage = v
}
func (m *TLInputBotInlineResultPhoto) GetSendMessage() *InputBotInlineMessage {
	return m.Data2.SendMessage
}

func (m *TLInputBotInlineResultDocument) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputBotInlineResultDocument) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputBotInlineResultDocument) SetId(v string) { m.Data2.Id = v }
func (m *TLInputBotInlineResultDocument) GetId() string  { return m.Data2.Id }

func (m *TLInputBotInlineResultDocument) SetType(v string) { m.Data2.Type = v }
func (m *TLInputBotInlineResultDocument) GetType() string  { return m.Data2.Type }

func (m *TLInputBotInlineResultDocument) SetTitle(v string) { m.Data2.Title = v }
func (m *TLInputBotInlineResultDocument) GetTitle() string  { return m.Data2.Title }

func (m *TLInputBotInlineResultDocument) SetDescription(v string) { m.Data2.Description = v }
func (m *TLInputBotInlineResultDocument) GetDescription() string  { return m.Data2.Description }

func (m *TLInputBotInlineResultDocument) SetDocument(v *InputDocument) { m.Data2.Document = v }
func (m *TLInputBotInlineResultDocument) GetDocument() *InputDocument  { return m.Data2.Document }

func (m *TLInputBotInlineResultDocument) SetSendMessage(v *InputBotInlineMessage) {
	m.Data2.SendMessage = v
}
func (m *TLInputBotInlineResultDocument) GetSendMessage() *InputBotInlineMessage {
	return m.Data2.SendMessage
}

func (m *TLInputBotInlineResultGame) SetId(v string) { m.Data2.Id = v }
func (m *TLInputBotInlineResultGame) GetId() string  { return m.Data2.Id }

func (m *TLInputBotInlineResultGame) SetShortName(v string) { m.Data2.ShortName = v }
func (m *TLInputBotInlineResultGame) GetShortName() string  { return m.Data2.ShortName }

func (m *TLInputBotInlineResultGame) SetSendMessage(v *InputBotInlineMessage) {
	m.Data2.SendMessage = v
}
func (m *TLInputBotInlineResultGame) GetSendMessage() *InputBotInlineMessage {
	return m.Data2.SendMessage
}

func (m *TLInputChannel) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLInputChannel) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLInputChannel) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputChannel) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputChannelFromMessage) SetPeer(v *InputPeer) { m.Data2.Peer = v }
func (m *TLInputChannelFromMessage) GetPeer() *InputPeer  { return m.Data2.Peer }

func (m *TLInputChannelFromMessage) SetMsgId(v int32) { m.Data2.MsgId = v }
func (m *TLInputChannelFromMessage) GetMsgId() int32  { return m.Data2.MsgId }

func (m *TLInputChannelFromMessage) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLInputChannelFromMessage) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLInputChatUploadedPhoto) SetFile(v *InputFile) { m.Data2.File = v }
func (m *TLInputChatUploadedPhoto) GetFile() *InputFile  { return m.Data2.File }

func (m *TLInputChatUploadedPhoto) SetCrop(v *InputPhotoCrop) { m.Data2.Crop = v }
func (m *TLInputChatUploadedPhoto) GetCrop() *InputPhotoCrop  { return m.Data2.Crop }

func (m *TLInputChatUploadedPhoto) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputChatUploadedPhoto) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputChatUploadedPhoto) SetVideo(v *InputFile) { m.Data2.Video = v }
func (m *TLInputChatUploadedPhoto) GetVideo() *InputFile  { return m.Data2.Video }

func (m *TLInputChatUploadedPhoto) SetVideoStartTs(v float64) { m.Data2.VideoStartTs = v }
func (m *TLInputChatUploadedPhoto) GetVideoStartTs() float64  { return m.Data2.VideoStartTs }

func (m *TLInputChatPhoto) SetId(v *InputPhoto) { m.Data2.Id = v }
func (m *TLInputChatPhoto) GetId() *InputPhoto  { return m.Data2.Id }

func (m *TLInputChatPhoto) SetCrop(v *InputPhotoCrop) { m.Data2.Crop = v }
func (m *TLInputChatPhoto) GetCrop() *InputPhotoCrop  { return m.Data2.Crop }

func (m *TLInputCheckPasswordSRP) SetSrpId(v int64) { m.Data2.SrpId = v }
func (m *TLInputCheckPasswordSRP) GetSrpId() int64  { return m.Data2.SrpId }

func (m *TLInputCheckPasswordSRP) SetA(v []byte) { m.Data2.A = v }
func (m *TLInputCheckPasswordSRP) GetA() []byte  { return m.Data2.A }

func (m *TLInputCheckPasswordSRP) SetM1(v []byte) { m.Data2.M1 = v }
func (m *TLInputCheckPasswordSRP) GetM1() []byte  { return m.Data2.M1 }

func (m *TLInputClientProxy) SetAddress(v string) { m.Data2.Address = v }
func (m *TLInputClientProxy) GetAddress() string  { return m.Data2.Address }

func (m *TLInputClientProxy) SetPort(v int32) { m.Data2.Port = v }
func (m *TLInputClientProxy) GetPort() int32  { return m.Data2.Port }

func (m *TLInputPhoneContact) SetClientId(v int64) { m.Data2.ClientId = v }
func (m *TLInputPhoneContact) GetClientId() int64  { return m.Data2.ClientId }

func (m *TLInputPhoneContact) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLInputPhoneContact) GetPhone() string  { return m.Data2.Phone }

func (m *TLInputPhoneContact) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLInputPhoneContact) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLInputPhoneContact) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLInputPhoneContact) GetLastName() string  { return m.Data2.LastName }

func (m *TLInputDialogPeer) SetPeer(v *InputPeer) { m.Data2.Peer = v }
func (m *TLInputDialogPeer) GetPeer() *InputPeer  { return m.Data2.Peer }

func (m *TLInputDialogPeerFolder) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLInputDialogPeerFolder) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLInputDocument) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputDocument) GetId() int64  { return m.Data2.Id }

func (m *TLInputDocument) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputDocument) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputDocument) SetFileReference(v []byte) { m.Data2.FileReference = v }
func (m *TLInputDocument) GetFileReference() []byte  { return m.Data2.FileReference }

func (m *TLInputEncryptedChat) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLInputEncryptedChat) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLInputEncryptedChat) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputEncryptedChat) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputEncryptedFileUploaded) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputEncryptedFileUploaded) GetId() int64  { return m.Data2.Id }

func (m *TLInputEncryptedFileUploaded) SetParts(v int32) { m.Data2.Parts = v }
func (m *TLInputEncryptedFileUploaded) GetParts() int32  { return m.Data2.Parts }

func (m *TLInputEncryptedFileUploaded) SetMd5Checksum(v string) { m.Data2.Md5Checksum = v }
func (m *TLInputEncryptedFileUploaded) GetMd5Checksum() string  { return m.Data2.Md5Checksum }

func (m *TLInputEncryptedFileUploaded) SetKeyFingerprint(v int32) { m.Data2.KeyFingerprint = v }
func (m *TLInputEncryptedFileUploaded) GetKeyFingerprint() int32  { return m.Data2.KeyFingerprint }

func (m *TLInputEncryptedFile) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputEncryptedFile) GetId() int64  { return m.Data2.Id }

func (m *TLInputEncryptedFile) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputEncryptedFile) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputEncryptedFileBigUploaded) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputEncryptedFileBigUploaded) GetId() int64  { return m.Data2.Id }

func (m *TLInputEncryptedFileBigUploaded) SetParts(v int32) { m.Data2.Parts = v }
func (m *TLInputEncryptedFileBigUploaded) GetParts() int32  { return m.Data2.Parts }

func (m *TLInputEncryptedFileBigUploaded) SetKeyFingerprint(v int32) { m.Data2.KeyFingerprint = v }
func (m *TLInputEncryptedFileBigUploaded) GetKeyFingerprint() int32  { return m.Data2.KeyFingerprint }

func (m *TLInputFile) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputFile) GetId() int64  { return m.Data2.Id }

func (m *TLInputFile) SetParts(v int32) { m.Data2.Parts = v }
func (m *TLInputFile) GetParts() int32  { return m.Data2.Parts }

func (m *TLInputFile) SetName(v string) { m.Data2.Name = v }
func (m *TLInputFile) GetName() string  { return m.Data2.Name }

func (m *TLInputFile) SetMd5Checksum(v string) { m.Data2.Md5Checksum = v }
func (m *TLInputFile) GetMd5Checksum() string  { return m.Data2.Md5Checksum }

func (m *TLInputFileBig) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputFileBig) GetId() int64  { return m.Data2.Id }

func (m *TLInputFileBig) SetParts(v int32) { m.Data2.Parts = v }
func (m *TLInputFileBig) GetParts() int32  { return m.Data2.Parts }

func (m *TLInputFileBig) SetName(v string) { m.Data2.Name = v }
func (m *TLInputFileBig) GetName() string  { return m.Data2.Name }

func (m *TLInputFileLocation) SetVolumeId(v int64) { m.Data2.VolumeId = v }
func (m *TLInputFileLocation) GetVolumeId() int64  { return m.Data2.VolumeId }

func (m *TLInputFileLocation) SetLocalId(v int32) { m.Data2.LocalId = v }
func (m *TLInputFileLocation) GetLocalId() int32  { return m.Data2.LocalId }

func (m *TLInputFileLocation) SetSecret(v int64) { m.Data2.Secret = v }
func (m *TLInputFileLocation) GetSecret() int64  { return m.Data2.Secret }

func (m *TLInputFileLocation) SetFileReference(v []byte) { m.Data2.FileReference = v }
func (m *TLInputFileLocation) GetFileReference() []byte  { return m.Data2.FileReference }

func (m *TLInputEncryptedFileLocation) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputEncryptedFileLocation) GetId() int64  { return m.Data2.Id }

func (m *TLInputEncryptedFileLocation) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputEncryptedFileLocation) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputDocumentFileLocation) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputDocumentFileLocation) GetId() int64  { return m.Data2.Id }

func (m *TLInputDocumentFileLocation) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputDocumentFileLocation) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputDocumentFileLocation) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLInputDocumentFileLocation) GetVersion() int32  { return m.Data2.Version }

func (m *TLInputDocumentFileLocation) SetFileReference(v []byte) { m.Data2.FileReference = v }
func (m *TLInputDocumentFileLocation) GetFileReference() []byte  { return m.Data2.FileReference }

func (m *TLInputDocumentFileLocation) SetThumbSize(v string) { m.Data2.ThumbSize = v }
func (m *TLInputDocumentFileLocation) GetThumbSize() string  { return m.Data2.ThumbSize }

func (m *TLInputSecureFileLocation) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputSecureFileLocation) GetId() int64  { return m.Data2.Id }

func (m *TLInputSecureFileLocation) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputSecureFileLocation) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputPhotoFileLocation) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputPhotoFileLocation) GetId() int64  { return m.Data2.Id }

func (m *TLInputPhotoFileLocation) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputPhotoFileLocation) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputPhotoFileLocation) SetFileReference(v []byte) { m.Data2.FileReference = v }
func (m *TLInputPhotoFileLocation) GetFileReference() []byte  { return m.Data2.FileReference }

func (m *TLInputPhotoFileLocation) SetThumbSize(v string) { m.Data2.ThumbSize = v }
func (m *TLInputPhotoFileLocation) GetThumbSize() string  { return m.Data2.ThumbSize }

func (m *TLInputPeerPhotoFileLocation) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputPeerPhotoFileLocation) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputPeerPhotoFileLocation) SetBig(v bool) { m.Data2.Big = v }
func (m *TLInputPeerPhotoFileLocation) GetBig() bool  { return m.Data2.Big }

func (m *TLInputPeerPhotoFileLocation) SetPeer(v *InputPeer) { m.Data2.Peer = v }
func (m *TLInputPeerPhotoFileLocation) GetPeer() *InputPeer  { return m.Data2.Peer }

func (m *TLInputPeerPhotoFileLocation) SetVolumeId(v int64) { m.Data2.VolumeId = v }
func (m *TLInputPeerPhotoFileLocation) GetVolumeId() int64  { return m.Data2.VolumeId }

func (m *TLInputPeerPhotoFileLocation) SetLocalId(v int32) { m.Data2.LocalId = v }
func (m *TLInputPeerPhotoFileLocation) GetLocalId() int32  { return m.Data2.LocalId }

func (m *TLInputStickerSetThumb) SetStickerset(v *InputStickerSet) { m.Data2.Stickerset = v }
func (m *TLInputStickerSetThumb) GetStickerset() *InputStickerSet  { return m.Data2.Stickerset }

func (m *TLInputStickerSetThumb) SetVolumeId(v int64) { m.Data2.VolumeId = v }
func (m *TLInputStickerSetThumb) GetVolumeId() int64  { return m.Data2.VolumeId }

func (m *TLInputStickerSetThumb) SetLocalId(v int32) { m.Data2.LocalId = v }
func (m *TLInputStickerSetThumb) GetLocalId() int32  { return m.Data2.LocalId }

func (m *TLInputPhotoLegacyFileLocation) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputPhotoLegacyFileLocation) GetId() int64  { return m.Data2.Id }

func (m *TLInputPhotoLegacyFileLocation) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputPhotoLegacyFileLocation) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputPhotoLegacyFileLocation) SetFileReference(v []byte) { m.Data2.FileReference = v }
func (m *TLInputPhotoLegacyFileLocation) GetFileReference() []byte  { return m.Data2.FileReference }

func (m *TLInputPhotoLegacyFileLocation) SetVolumeId(v int64) { m.Data2.VolumeId = v }
func (m *TLInputPhotoLegacyFileLocation) GetVolumeId() int64  { return m.Data2.VolumeId }

func (m *TLInputPhotoLegacyFileLocation) SetLocalId(v int32) { m.Data2.LocalId = v }
func (m *TLInputPhotoLegacyFileLocation) GetLocalId() int32  { return m.Data2.LocalId }

func (m *TLInputPhotoLegacyFileLocation) SetSecret(v int64) { m.Data2.Secret = v }
func (m *TLInputPhotoLegacyFileLocation) GetSecret() int64  { return m.Data2.Secret }

func (m *TLInputFolderPeer) SetPeer(v *InputPeer) { m.Data2.Peer = v }
func (m *TLInputFolderPeer) GetPeer() *InputPeer  { return m.Data2.Peer }

func (m *TLInputFolderPeer) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLInputFolderPeer) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLInputGameID) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputGameID) GetId() int64  { return m.Data2.Id }

func (m *TLInputGameID) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputGameID) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputGameShortName) SetBotId(v *InputUser) { m.Data2.BotId = v }
func (m *TLInputGameShortName) GetBotId() *InputUser  { return m.Data2.BotId }

func (m *TLInputGameShortName) SetShortName(v string) { m.Data2.ShortName = v }
func (m *TLInputGameShortName) GetShortName() string  { return m.Data2.ShortName }

func (m *TLInputGeoPoint) SetLat(v float64) { m.Data2.Lat = v }
func (m *TLInputGeoPoint) GetLat() float64  { return m.Data2.Lat }

func (m *TLInputGeoPoint) SetLong(v float64) { m.Data2.Long = v }
func (m *TLInputGeoPoint) GetLong() float64  { return m.Data2.Long }

func (m *TLInputGeoPoint) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputGeoPoint) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputGeoPoint) SetAccuracyRadius(v int32) { m.Data2.AccuracyRadius = v }
func (m *TLInputGeoPoint) GetAccuracyRadius() int32  { return m.Data2.AccuracyRadius }

func (m *TLInputGroupCall) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputGroupCall) GetId() int64  { return m.Data2.Id }

func (m *TLInputGroupCall) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputGroupCall) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputMediaUploadedPhoto) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputMediaUploadedPhoto) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputMediaUploadedPhoto) SetFile(v *InputFile) { m.Data2.File = v }
func (m *TLInputMediaUploadedPhoto) GetFile() *InputFile  { return m.Data2.File }

func (m *TLInputMediaUploadedPhoto) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaUploadedPhoto) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaUploadedPhoto) SetStickers(v []*InputDocument) { m.Data2.Stickers = v }
func (m *TLInputMediaUploadedPhoto) GetStickers() []*InputDocument  { return m.Data2.Stickers }

func (m *TLInputMediaUploadedPhoto) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaUploadedPhoto) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaPhoto) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputMediaPhoto) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputMediaPhoto) SetId81FA373A71(v *InputPhoto) { m.Data2.Id81FA373A71 = v }
func (m *TLInputMediaPhoto) GetId81FA373A71() *InputPhoto  { return m.Data2.Id81FA373A71 }

func (m *TLInputMediaPhoto) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaPhoto) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaPhoto) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaPhoto) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaGeoPoint) SetGeoPoint(v *InputGeoPoint) { m.Data2.GeoPoint = v }
func (m *TLInputMediaGeoPoint) GetGeoPoint() *InputGeoPoint  { return m.Data2.GeoPoint }

func (m *TLInputMediaContact) SetPhoneNumber(v string) { m.Data2.PhoneNumber = v }
func (m *TLInputMediaContact) GetPhoneNumber() string  { return m.Data2.PhoneNumber }

func (m *TLInputMediaContact) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLInputMediaContact) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLInputMediaContact) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLInputMediaContact) GetLastName() string  { return m.Data2.LastName }

func (m *TLInputMediaContact) SetVcard(v string) { m.Data2.Vcard = v }
func (m *TLInputMediaContact) GetVcard() string  { return m.Data2.Vcard }

func (m *TLInputMediaUploadedDocument) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputMediaUploadedDocument) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputMediaUploadedDocument) SetFile(v *InputFile) { m.Data2.File = v }
func (m *TLInputMediaUploadedDocument) GetFile() *InputFile  { return m.Data2.File }

func (m *TLInputMediaUploadedDocument) SetThumb(v *InputFile) { m.Data2.Thumb = v }
func (m *TLInputMediaUploadedDocument) GetThumb() *InputFile  { return m.Data2.Thumb }

func (m *TLInputMediaUploadedDocument) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLInputMediaUploadedDocument) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLInputMediaUploadedDocument) SetAttributes(v []*DocumentAttribute) { m.Data2.Attributes = v }
func (m *TLInputMediaUploadedDocument) GetAttributes() []*DocumentAttribute {
	return m.Data2.Attributes
}

func (m *TLInputMediaUploadedDocument) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaUploadedDocument) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaUploadedDocument) SetStickers(v []*InputDocument) { m.Data2.Stickers = v }
func (m *TLInputMediaUploadedDocument) GetStickers() []*InputDocument  { return m.Data2.Stickers }

func (m *TLInputMediaUploadedDocument) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaUploadedDocument) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaUploadedDocument) SetNosoundVideo(v bool) { m.Data2.NosoundVideo = v }
func (m *TLInputMediaUploadedDocument) GetNosoundVideo() bool  { return m.Data2.NosoundVideo }

func (m *TLInputMediaUploadedDocument) SetForceFile(v bool) { m.Data2.ForceFile = v }
func (m *TLInputMediaUploadedDocument) GetForceFile() bool  { return m.Data2.ForceFile }

func (m *TLInputMediaDocument) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputMediaDocument) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputMediaDocument) SetId5ACB668E71(v *InputDocument) { m.Data2.Id5ACB668E71 = v }
func (m *TLInputMediaDocument) GetId5ACB668E71() *InputDocument  { return m.Data2.Id5ACB668E71 }

func (m *TLInputMediaDocument) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaDocument) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaDocument) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaDocument) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaDocument) SetQuery(v string) { m.Data2.Query = v }
func (m *TLInputMediaDocument) GetQuery() string  { return m.Data2.Query }

func (m *TLInputMediaVenue) SetGeoPoint(v *InputGeoPoint) { m.Data2.GeoPoint = v }
func (m *TLInputMediaVenue) GetGeoPoint() *InputGeoPoint  { return m.Data2.GeoPoint }

func (m *TLInputMediaVenue) SetTitle(v string) { m.Data2.Title = v }
func (m *TLInputMediaVenue) GetTitle() string  { return m.Data2.Title }

func (m *TLInputMediaVenue) SetAddress(v string) { m.Data2.Address = v }
func (m *TLInputMediaVenue) GetAddress() string  { return m.Data2.Address }

func (m *TLInputMediaVenue) SetProvider(v string) { m.Data2.Provider = v }
func (m *TLInputMediaVenue) GetProvider() string  { return m.Data2.Provider }

func (m *TLInputMediaVenue) SetVenueId(v string) { m.Data2.VenueId = v }
func (m *TLInputMediaVenue) GetVenueId() string  { return m.Data2.VenueId }

func (m *TLInputMediaVenue) SetVenueType(v string) { m.Data2.VenueType = v }
func (m *TLInputMediaVenue) GetVenueType() string  { return m.Data2.VenueType }

func (m *TLInputMediaGifExternal) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputMediaGifExternal) GetUrl() string  { return m.Data2.Url }

func (m *TLInputMediaGifExternal) SetQ(v string) { m.Data2.Q = v }
func (m *TLInputMediaGifExternal) GetQ() string  { return m.Data2.Q }

func (m *TLInputMediaPhotoExternal) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputMediaPhotoExternal) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputMediaPhotoExternal) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputMediaPhotoExternal) GetUrl() string  { return m.Data2.Url }

func (m *TLInputMediaPhotoExternal) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaPhotoExternal) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaPhotoExternal) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaPhotoExternal) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaDocumentExternal) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputMediaDocumentExternal) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputMediaDocumentExternal) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputMediaDocumentExternal) GetUrl() string  { return m.Data2.Url }

func (m *TLInputMediaDocumentExternal) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaDocumentExternal) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaDocumentExternal) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaDocumentExternal) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaGame) SetIdD33F43F371(v *InputGame) { m.Data2.IdD33F43F371 = v }
func (m *TLInputMediaGame) GetIdD33F43F371() *InputGame  { return m.Data2.IdD33F43F371 }

func (m *TLInputMediaInvoice) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputMediaInvoice) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputMediaInvoice) SetTitle(v string) { m.Data2.Title = v }
func (m *TLInputMediaInvoice) GetTitle() string  { return m.Data2.Title }

func (m *TLInputMediaInvoice) SetDescription(v string) { m.Data2.Description = v }
func (m *TLInputMediaInvoice) GetDescription() string  { return m.Data2.Description }

func (m *TLInputMediaInvoice) SetPhoto(v *InputWebDocument) { m.Data2.Photo = v }
func (m *TLInputMediaInvoice) GetPhoto() *InputWebDocument  { return m.Data2.Photo }

func (m *TLInputMediaInvoice) SetInvoice(v *Invoice) { m.Data2.Invoice = v }
func (m *TLInputMediaInvoice) GetInvoice() *Invoice  { return m.Data2.Invoice }

func (m *TLInputMediaInvoice) SetPayload(v []byte) { m.Data2.Payload = v }
func (m *TLInputMediaInvoice) GetPayload() []byte  { return m.Data2.Payload }

func (m *TLInputMediaInvoice) SetProvider(v string) { m.Data2.Provider = v }
func (m *TLInputMediaInvoice) GetProvider() string  { return m.Data2.Provider }

func (m *TLInputMediaInvoice) SetStartParam(v string) { m.Data2.StartParam = v }
func (m *TLInputMediaInvoice) GetStartParam() string  { return m.Data2.StartParam }

func (m *TLInputMediaInvoice) SetProviderData(v *DataJSON) { m.Data2.ProviderData = v }
func (m *TLInputMediaInvoice) GetProviderData() *DataJSON  { return m.Data2.ProviderData }

func (m *TLInputMediaUploadedThumbDocument) SetFile(v *InputFile) { m.Data2.File = v }
func (m *TLInputMediaUploadedThumbDocument) GetFile() *InputFile  { return m.Data2.File }

func (m *TLInputMediaUploadedThumbDocument) SetThumb(v *InputFile) { m.Data2.Thumb = v }
func (m *TLInputMediaUploadedThumbDocument) GetThumb() *InputFile  { return m.Data2.Thumb }

func (m *TLInputMediaUploadedThumbDocument) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLInputMediaUploadedThumbDocument) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLInputMediaUploadedThumbDocument) SetAttributes(v []*DocumentAttribute) {
	m.Data2.Attributes = v
}
func (m *TLInputMediaUploadedThumbDocument) GetAttributes() []*DocumentAttribute {
	return m.Data2.Attributes
}

func (m *TLInputMediaUploadedThumbDocument) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaUploadedThumbDocument) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaGeoLive) SetGeoPoint(v *InputGeoPoint) { m.Data2.GeoPoint = v }
func (m *TLInputMediaGeoLive) GetGeoPoint() *InputGeoPoint  { return m.Data2.GeoPoint }

func (m *TLInputMediaGeoLive) SetPeriod(v int32) { m.Data2.Period = v }
func (m *TLInputMediaGeoLive) GetPeriod() int32  { return m.Data2.Period }

func (m *TLInputMediaGeoLive) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputMediaGeoLive) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputMediaGeoLive) SetStopped(v bool) { m.Data2.Stopped = v }
func (m *TLInputMediaGeoLive) GetStopped() bool  { return m.Data2.Stopped }

func (m *TLInputMediaGeoLive) SetHeading(v int32) { m.Data2.Heading = v }
func (m *TLInputMediaGeoLive) GetHeading() int32  { return m.Data2.Heading }

func (m *TLInputMediaGeoLive) SetProximityNotificationRadius(v int32) {
	m.Data2.ProximityNotificationRadius = v
}
func (m *TLInputMediaGeoLive) GetProximityNotificationRadius() int32 {
	return m.Data2.ProximityNotificationRadius
}

func (m *TLInputMediaPoll) SetPoll(v *Poll) { m.Data2.Poll = v }
func (m *TLInputMediaPoll) GetPoll() *Poll  { return m.Data2.Poll }

func (m *TLInputMediaPoll) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputMediaPoll) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputMediaPoll) SetCorrectAnswers(v [][]byte) { m.Data2.CorrectAnswers = v }
func (m *TLInputMediaPoll) GetCorrectAnswers() [][]byte  { return m.Data2.CorrectAnswers }

func (m *TLInputMediaPoll) SetSolution(v string) { m.Data2.Solution = v }
func (m *TLInputMediaPoll) GetSolution() string  { return m.Data2.Solution }

func (m *TLInputMediaPoll) SetSolutionEntities(v []*MessageEntity) { m.Data2.SolutionEntities = v }
func (m *TLInputMediaPoll) GetSolutionEntities() []*MessageEntity  { return m.Data2.SolutionEntities }

func (m *TLInputMediaDice) SetEmoticon(v string) { m.Data2.Emoticon = v }
func (m *TLInputMediaDice) GetEmoticon() string  { return m.Data2.Emoticon }

func (m *TLInputMessageID) SetId(v int32) { m.Data2.Id = v }
func (m *TLInputMessageID) GetId() int32  { return m.Data2.Id }

func (m *TLInputMessageReplyTo) SetId(v int32) { m.Data2.Id = v }
func (m *TLInputMessageReplyTo) GetId() int32  { return m.Data2.Id }

func (m *TLInputMessageCallbackQuery) SetId(v int32) { m.Data2.Id = v }
func (m *TLInputMessageCallbackQuery) GetId() int32  { return m.Data2.Id }

func (m *TLInputMessageCallbackQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLInputMessageCallbackQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLInputNotifyPeer) SetPeer(v *InputPeer) { m.Data2.Peer = v }
func (m *TLInputNotifyPeer) GetPeer() *InputPeer  { return m.Data2.Peer }

func (m *TLInputPaymentCredentialsSaved) SetId(v string) { m.Data2.Id = v }
func (m *TLInputPaymentCredentialsSaved) GetId() string  { return m.Data2.Id }

func (m *TLInputPaymentCredentialsSaved) SetTmpPassword(v []byte) { m.Data2.TmpPassword = v }
func (m *TLInputPaymentCredentialsSaved) GetTmpPassword() []byte  { return m.Data2.TmpPassword }

func (m *TLInputPaymentCredentials) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputPaymentCredentials) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputPaymentCredentials) SetSave(v bool) { m.Data2.Save = v }
func (m *TLInputPaymentCredentials) GetSave() bool  { return m.Data2.Save }

func (m *TLInputPaymentCredentials) SetData(v *DataJSON) { m.Data2.Data = v }
func (m *TLInputPaymentCredentials) GetData() *DataJSON  { return m.Data2.Data }

func (m *TLInputPaymentCredentialsApplePay) SetPaymentData(v *DataJSON) { m.Data2.PaymentData = v }
func (m *TLInputPaymentCredentialsApplePay) GetPaymentData() *DataJSON  { return m.Data2.PaymentData }

func (m *TLInputPaymentCredentialsAndroidPay) SetPaymentToken(v *DataJSON) { m.Data2.PaymentToken = v }
func (m *TLInputPaymentCredentialsAndroidPay) GetPaymentToken() *DataJSON {
	return m.Data2.PaymentToken
}

func (m *TLInputPaymentCredentialsAndroidPay) SetGoogleTransactionId(v string) {
	m.Data2.GoogleTransactionId = v
}
func (m *TLInputPaymentCredentialsAndroidPay) GetGoogleTransactionId() string {
	return m.Data2.GoogleTransactionId
}

func (m *TLInputPeerChat) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLInputPeerChat) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLInputPeerUser) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLInputPeerUser) GetUserId() int32  { return m.Data2.UserId }

func (m *TLInputPeerUser) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputPeerUser) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputPeerChannel) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLInputPeerChannel) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLInputPeerChannel) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputPeerChannel) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputPeerUserFromMessage) SetPeer(v *InputPeer) { m.Data2.Peer = v }
func (m *TLInputPeerUserFromMessage) GetPeer() *InputPeer  { return m.Data2.Peer }

func (m *TLInputPeerUserFromMessage) SetMsgId(v int32) { m.Data2.MsgId = v }
func (m *TLInputPeerUserFromMessage) GetMsgId() int32  { return m.Data2.MsgId }

func (m *TLInputPeerUserFromMessage) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLInputPeerUserFromMessage) GetUserId() int32  { return m.Data2.UserId }

func (m *TLInputPeerChannelFromMessage) SetPeer(v *InputPeer) { m.Data2.Peer = v }
func (m *TLInputPeerChannelFromMessage) GetPeer() *InputPeer  { return m.Data2.Peer }

func (m *TLInputPeerChannelFromMessage) SetMsgId(v int32) { m.Data2.MsgId = v }
func (m *TLInputPeerChannelFromMessage) GetMsgId() int32  { return m.Data2.MsgId }

func (m *TLInputPeerChannelFromMessage) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLInputPeerChannelFromMessage) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLInputPeerNotifySettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputPeerNotifySettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputPeerNotifySettings) SetShowPreviews38935EB271(v bool) {
	m.Data2.ShowPreviews38935EB271 = v
}
func (m *TLInputPeerNotifySettings) GetShowPreviews38935EB271() bool {
	return m.Data2.ShowPreviews38935EB271
}

func (m *TLInputPeerNotifySettings) SetSilent38935EB271(v bool) { m.Data2.Silent38935EB271 = v }
func (m *TLInputPeerNotifySettings) GetSilent38935EB271() bool  { return m.Data2.Silent38935EB271 }

func (m *TLInputPeerNotifySettings) SetMuteUntil(v int32) { m.Data2.MuteUntil = v }
func (m *TLInputPeerNotifySettings) GetMuteUntil() int32  { return m.Data2.MuteUntil }

func (m *TLInputPeerNotifySettings) SetSound(v string) { m.Data2.Sound = v }
func (m *TLInputPeerNotifySettings) GetSound() string  { return m.Data2.Sound }

func (m *TLInputPeerNotifySettings) SetShowPreviews9C3D198E85(v *Bool) {
	m.Data2.ShowPreviews9C3D198E85 = v
}
func (m *TLInputPeerNotifySettings) GetShowPreviews9C3D198E85() *Bool {
	return m.Data2.ShowPreviews9C3D198E85
}

func (m *TLInputPeerNotifySettings) SetSilent9C3D198E85(v *Bool) { m.Data2.Silent9C3D198E85 = v }
func (m *TLInputPeerNotifySettings) GetSilent9C3D198E85() *Bool  { return m.Data2.Silent9C3D198E85 }

func (m *TLInputPhoneCall) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputPhoneCall) GetId() int64  { return m.Data2.Id }

func (m *TLInputPhoneCall) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputPhoneCall) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputPhoto) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputPhoto) GetId() int64  { return m.Data2.Id }

func (m *TLInputPhoto) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputPhoto) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputPhoto) SetFileReference(v []byte) { m.Data2.FileReference = v }
func (m *TLInputPhoto) GetFileReference() []byte  { return m.Data2.FileReference }

func (m *TLInputPhotoCrop) SetCropLeft(v float64) { m.Data2.CropLeft = v }
func (m *TLInputPhotoCrop) GetCropLeft() float64  { return m.Data2.CropLeft }

func (m *TLInputPhotoCrop) SetCropTop(v float64) { m.Data2.CropTop = v }
func (m *TLInputPhotoCrop) GetCropTop() float64  { return m.Data2.CropTop }

func (m *TLInputPhotoCrop) SetCropWidth(v float64) { m.Data2.CropWidth = v }
func (m *TLInputPhotoCrop) GetCropWidth() float64  { return m.Data2.CropWidth }

func (m *TLInputPrivacyValueAllowUsers) SetUsers(v []*InputUser) { m.Data2.Users = v }
func (m *TLInputPrivacyValueAllowUsers) GetUsers() []*InputUser  { return m.Data2.Users }

func (m *TLInputPrivacyValueDisallowUsers) SetUsers(v []*InputUser) { m.Data2.Users = v }
func (m *TLInputPrivacyValueDisallowUsers) GetUsers() []*InputUser  { return m.Data2.Users }

func (m *TLInputPrivacyValueAllowChatParticipants) SetChats(v []int32) { m.Data2.Chats = v }
func (m *TLInputPrivacyValueAllowChatParticipants) GetChats() []int32  { return m.Data2.Chats }

func (m *TLInputPrivacyValueDisallowChatParticipants) SetChats(v []int32) { m.Data2.Chats = v }
func (m *TLInputPrivacyValueDisallowChatParticipants) GetChats() []int32  { return m.Data2.Chats }

func (m *TLInputSecureFileUploaded) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputSecureFileUploaded) GetId() int64  { return m.Data2.Id }

func (m *TLInputSecureFileUploaded) SetParts(v int32) { m.Data2.Parts = v }
func (m *TLInputSecureFileUploaded) GetParts() int32  { return m.Data2.Parts }

func (m *TLInputSecureFileUploaded) SetMd5Checksum(v string) { m.Data2.Md5Checksum = v }
func (m *TLInputSecureFileUploaded) GetMd5Checksum() string  { return m.Data2.Md5Checksum }

func (m *TLInputSecureFileUploaded) SetFileHash(v []byte) { m.Data2.FileHash = v }
func (m *TLInputSecureFileUploaded) GetFileHash() []byte  { return m.Data2.FileHash }

func (m *TLInputSecureFileUploaded) SetSecret(v []byte) { m.Data2.Secret = v }
func (m *TLInputSecureFileUploaded) GetSecret() []byte  { return m.Data2.Secret }

func (m *TLInputSecureFile) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputSecureFile) GetId() int64  { return m.Data2.Id }

func (m *TLInputSecureFile) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputSecureFile) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputSecureValue) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputSecureValue) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputSecureValue) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLInputSecureValue) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLInputSecureValue) SetData(v *SecureData) { m.Data2.Data = v }
func (m *TLInputSecureValue) GetData() *SecureData  { return m.Data2.Data }

func (m *TLInputSecureValue) SetFrontSide(v *InputSecureFile) { m.Data2.FrontSide = v }
func (m *TLInputSecureValue) GetFrontSide() *InputSecureFile  { return m.Data2.FrontSide }

func (m *TLInputSecureValue) SetReverseSide(v *InputSecureFile) { m.Data2.ReverseSide = v }
func (m *TLInputSecureValue) GetReverseSide() *InputSecureFile  { return m.Data2.ReverseSide }

func (m *TLInputSecureValue) SetSelfie(v *InputSecureFile) { m.Data2.Selfie = v }
func (m *TLInputSecureValue) GetSelfie() *InputSecureFile  { return m.Data2.Selfie }

func (m *TLInputSecureValue) SetTranslation(v []*InputSecureFile) { m.Data2.Translation = v }
func (m *TLInputSecureValue) GetTranslation() []*InputSecureFile  { return m.Data2.Translation }

func (m *TLInputSecureValue) SetFiles(v []*InputSecureFile) { m.Data2.Files = v }
func (m *TLInputSecureValue) GetFiles() []*InputSecureFile  { return m.Data2.Files }

func (m *TLInputSecureValue) SetPlainData(v *SecurePlainData) { m.Data2.PlainData = v }
func (m *TLInputSecureValue) GetPlainData() *SecurePlainData  { return m.Data2.PlainData }

func (m *TLInputSingleMedia) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputSingleMedia) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputSingleMedia) SetMedia(v *InputMedia) { m.Data2.Media = v }
func (m *TLInputSingleMedia) GetMedia() *InputMedia  { return m.Data2.Media }

func (m *TLInputSingleMedia) SetRandomId(v int64) { m.Data2.RandomId = v }
func (m *TLInputSingleMedia) GetRandomId() int64  { return m.Data2.RandomId }

func (m *TLInputSingleMedia) SetMessage(v string) { m.Data2.Message = v }
func (m *TLInputSingleMedia) GetMessage() string  { return m.Data2.Message }

func (m *TLInputSingleMedia) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLInputSingleMedia) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLInputStickerSetID) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputStickerSetID) GetId() int64  { return m.Data2.Id }

func (m *TLInputStickerSetID) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputStickerSetID) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputStickerSetShortName) SetShortName(v string) { m.Data2.ShortName = v }
func (m *TLInputStickerSetShortName) GetShortName() string  { return m.Data2.ShortName }

func (m *TLInputStickerSetDice) SetEmoticon(v string) { m.Data2.Emoticon = v }
func (m *TLInputStickerSetDice) GetEmoticon() string  { return m.Data2.Emoticon }

func (m *TLInputStickerSetItem) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputStickerSetItem) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputStickerSetItem) SetDocument(v *InputDocument) { m.Data2.Document = v }
func (m *TLInputStickerSetItem) GetDocument() *InputDocument  { return m.Data2.Document }

func (m *TLInputStickerSetItem) SetEmoji(v string) { m.Data2.Emoji = v }
func (m *TLInputStickerSetItem) GetEmoji() string  { return m.Data2.Emoji }

func (m *TLInputStickerSetItem) SetMaskCoords(v *MaskCoords) { m.Data2.MaskCoords = v }
func (m *TLInputStickerSetItem) GetMaskCoords() *MaskCoords  { return m.Data2.MaskCoords }

func (m *TLInputStickeredMediaPhoto) SetId4A99215771(v *InputPhoto) { m.Data2.Id4A99215771 = v }
func (m *TLInputStickeredMediaPhoto) GetId4A99215771() *InputPhoto  { return m.Data2.Id4A99215771 }

func (m *TLInputStickeredMediaDocument) SetId438865B71(v *InputDocument) { m.Data2.Id438865B71 = v }
func (m *TLInputStickeredMediaDocument) GetId438865B71() *InputDocument  { return m.Data2.Id438865B71 }

func (m *TLInputTheme) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputTheme) GetId() int64  { return m.Data2.Id }

func (m *TLInputTheme) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputTheme) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputThemeSlug) SetSlug(v string) { m.Data2.Slug = v }
func (m *TLInputThemeSlug) GetSlug() string  { return m.Data2.Slug }

func (m *TLInputThemeSettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputThemeSettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputThemeSettings) SetBaseTheme(v *BaseTheme) { m.Data2.BaseTheme = v }
func (m *TLInputThemeSettings) GetBaseTheme() *BaseTheme  { return m.Data2.BaseTheme }

func (m *TLInputThemeSettings) SetAccentColor(v int32) { m.Data2.AccentColor = v }
func (m *TLInputThemeSettings) GetAccentColor() int32  { return m.Data2.AccentColor }

func (m *TLInputThemeSettings) SetMessageTopColor(v int32) { m.Data2.MessageTopColor = v }
func (m *TLInputThemeSettings) GetMessageTopColor() int32  { return m.Data2.MessageTopColor }

func (m *TLInputThemeSettings) SetMessageBottomColor(v int32) { m.Data2.MessageBottomColor = v }
func (m *TLInputThemeSettings) GetMessageBottomColor() int32  { return m.Data2.MessageBottomColor }

func (m *TLInputThemeSettings) SetWallpaper(v *InputWallPaper) { m.Data2.Wallpaper = v }
func (m *TLInputThemeSettings) GetWallpaper() *InputWallPaper  { return m.Data2.Wallpaper }

func (m *TLInputThemeSettings) SetWallpaperSettings(v *WallPaperSettings) {
	m.Data2.WallpaperSettings = v
}
func (m *TLInputThemeSettings) GetWallpaperSettings() *WallPaperSettings {
	return m.Data2.WallpaperSettings
}

func (m *TLInputUser) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLInputUser) GetUserId() int32  { return m.Data2.UserId }

func (m *TLInputUser) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputUser) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputUserFromMessage) SetPeer(v *InputPeer) { m.Data2.Peer = v }
func (m *TLInputUserFromMessage) GetPeer() *InputPeer  { return m.Data2.Peer }

func (m *TLInputUserFromMessage) SetMsgId(v int32) { m.Data2.MsgId = v }
func (m *TLInputUserFromMessage) GetMsgId() int32  { return m.Data2.MsgId }

func (m *TLInputUserFromMessage) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLInputUserFromMessage) GetUserId() int32  { return m.Data2.UserId }

func (m *TLInputWallPaper) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputWallPaper) GetId() int64  { return m.Data2.Id }

func (m *TLInputWallPaper) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputWallPaper) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputWallPaperSlug) SetSlug(v string) { m.Data2.Slug = v }
func (m *TLInputWallPaperSlug) GetSlug() string  { return m.Data2.Slug }

func (m *TLInputWebDocument) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputWebDocument) GetUrl() string  { return m.Data2.Url }

func (m *TLInputWebDocument) SetSize_(v int32) { m.Data2.Size_ = v }
func (m *TLInputWebDocument) GetSize_() int32  { return m.Data2.Size_ }

func (m *TLInputWebDocument) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLInputWebDocument) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLInputWebDocument) SetAttributes(v []*DocumentAttribute) { m.Data2.Attributes = v }
func (m *TLInputWebDocument) GetAttributes() []*DocumentAttribute  { return m.Data2.Attributes }

func (m *TLInputWebFileLocation) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputWebFileLocation) GetUrl() string  { return m.Data2.Url }

func (m *TLInputWebFileLocation) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputWebFileLocation) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputWebFileGeoPointLocation) SetGeoPoint(v *InputGeoPoint) { m.Data2.GeoPoint = v }
func (m *TLInputWebFileGeoPointLocation) GetGeoPoint() *InputGeoPoint  { return m.Data2.GeoPoint }

func (m *TLInputWebFileGeoPointLocation) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputWebFileGeoPointLocation) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputWebFileGeoPointLocation) SetW(v int32) { m.Data2.W = v }
func (m *TLInputWebFileGeoPointLocation) GetW() int32  { return m.Data2.W }

func (m *TLInputWebFileGeoPointLocation) SetH(v int32) { m.Data2.H = v }
func (m *TLInputWebFileGeoPointLocation) GetH() int32  { return m.Data2.H }

func (m *TLInputWebFileGeoPointLocation) SetZoom(v int32) { m.Data2.Zoom = v }
func (m *TLInputWebFileGeoPointLocation) GetZoom() int32  { return m.Data2.Zoom }

func (m *TLInputWebFileGeoPointLocation) SetScale(v int32) { m.Data2.Scale = v }
func (m *TLInputWebFileGeoPointLocation) GetScale() int32  { return m.Data2.Scale }

func (m *TLInvoice) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInvoice) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInvoice) SetTest(v bool) { m.Data2.Test = v }
func (m *TLInvoice) GetTest() bool  { return m.Data2.Test }

func (m *TLInvoice) SetNameRequested(v bool) { m.Data2.NameRequested = v }
func (m *TLInvoice) GetNameRequested() bool  { return m.Data2.NameRequested }

func (m *TLInvoice) SetPhoneRequested(v bool) { m.Data2.PhoneRequested = v }
func (m *TLInvoice) GetPhoneRequested() bool  { return m.Data2.PhoneRequested }

func (m *TLInvoice) SetEmailRequested(v bool) { m.Data2.EmailRequested = v }
func (m *TLInvoice) GetEmailRequested() bool  { return m.Data2.EmailRequested }

func (m *TLInvoice) SetShippingAddressRequested(v bool) { m.Data2.ShippingAddressRequested = v }
func (m *TLInvoice) GetShippingAddressRequested() bool  { return m.Data2.ShippingAddressRequested }

func (m *TLInvoice) SetFlexible(v bool) { m.Data2.Flexible = v }
func (m *TLInvoice) GetFlexible() bool  { return m.Data2.Flexible }

func (m *TLInvoice) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLInvoice) GetCurrency() string  { return m.Data2.Currency }

func (m *TLInvoice) SetPrices(v []*LabeledPrice) { m.Data2.Prices = v }
func (m *TLInvoice) GetPrices() []*LabeledPrice  { return m.Data2.Prices }

func (m *TLIpPort) SetIpv4(v int32) { m.Data2.Ipv4 = v }
func (m *TLIpPort) GetIpv4() int32  { return m.Data2.Ipv4 }

func (m *TLIpPort) SetPort(v int32) { m.Data2.Port = v }
func (m *TLIpPort) GetPort() int32  { return m.Data2.Port }

func (m *TLIpPortSecret) SetIpv4(v int32) { m.Data2.Ipv4 = v }
func (m *TLIpPortSecret) GetIpv4() int32  { return m.Data2.Ipv4 }

func (m *TLIpPortSecret) SetPort(v int32) { m.Data2.Port = v }
func (m *TLIpPortSecret) GetPort() int32  { return m.Data2.Port }

func (m *TLIpPortSecret) SetSecret(v []byte) { m.Data2.Secret = v }
func (m *TLIpPortSecret) GetSecret() []byte  { return m.Data2.Secret }

func (m *TLJsonObjectValue) SetKey(v string) { m.Data2.Key = v }
func (m *TLJsonObjectValue) GetKey() string  { return m.Data2.Key }

func (m *TLJsonObjectValue) SetValue(v *JSONValue) { m.Data2.Value = v }
func (m *TLJsonObjectValue) GetValue() *JSONValue  { return m.Data2.Value }

func (m *TLJsonBool) SetValueC7345E6A90(v *Bool) { m.Data2.ValueC7345E6A90 = v }
func (m *TLJsonBool) GetValueC7345E6A90() *Bool  { return m.Data2.ValueC7345E6A90 }

func (m *TLJsonNumber) SetValue2BE0DFA490(v float64) { m.Data2.Value2BE0DFA490 = v }
func (m *TLJsonNumber) GetValue2BE0DFA490() float64  { return m.Data2.Value2BE0DFA490 }

func (m *TLJsonString) SetValueB71E767A90(v string) { m.Data2.ValueB71E767A90 = v }
func (m *TLJsonString) GetValueB71E767A90() string  { return m.Data2.ValueB71E767A90 }

func (m *TLJsonArray) SetValueF744476390(v []*JSONValue) { m.Data2.ValueF744476390 = v }
func (m *TLJsonArray) GetValueF744476390() []*JSONValue  { return m.Data2.ValueF744476390 }

func (m *TLJsonObject) SetValue99C1D49D90(v []*JSONObjectValue) { m.Data2.Value99C1D49D90 = v }
func (m *TLJsonObject) GetValue99C1D49D90() []*JSONObjectValue  { return m.Data2.Value99C1D49D90 }

func (m *TLKeyboardButton) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButton) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonUrl) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonUrl) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonUrl) SetUrl(v string) { m.Data2.Url = v }
func (m *TLKeyboardButtonUrl) GetUrl() string  { return m.Data2.Url }

func (m *TLKeyboardButtonCallback) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonCallback) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonCallback) SetData(v []byte) { m.Data2.Data = v }
func (m *TLKeyboardButtonCallback) GetData() []byte  { return m.Data2.Data }

func (m *TLKeyboardButtonCallback) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLKeyboardButtonCallback) GetFlags() int32  { return m.Data2.Flags }

func (m *TLKeyboardButtonCallback) SetRequiresPassword(v bool) { m.Data2.RequiresPassword = v }
func (m *TLKeyboardButtonCallback) GetRequiresPassword() bool  { return m.Data2.RequiresPassword }

func (m *TLKeyboardButtonRequestPhone) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonRequestPhone) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonRequestGeoLocation) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonRequestGeoLocation) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonSwitchInline) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLKeyboardButtonSwitchInline) GetFlags() int32  { return m.Data2.Flags }

func (m *TLKeyboardButtonSwitchInline) SetSamePeer(v bool) { m.Data2.SamePeer = v }
func (m *TLKeyboardButtonSwitchInline) GetSamePeer() bool  { return m.Data2.SamePeer }

func (m *TLKeyboardButtonSwitchInline) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonSwitchInline) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonSwitchInline) SetQuery(v string) { m.Data2.Query = v }
func (m *TLKeyboardButtonSwitchInline) GetQuery() string  { return m.Data2.Query }

func (m *TLKeyboardButtonGame) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonGame) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonBuy) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonBuy) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonUrlAuth) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLKeyboardButtonUrlAuth) GetFlags() int32  { return m.Data2.Flags }

func (m *TLKeyboardButtonUrlAuth) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonUrlAuth) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonUrlAuth) SetFwdText(v string) { m.Data2.FwdText = v }
func (m *TLKeyboardButtonUrlAuth) GetFwdText() string  { return m.Data2.FwdText }

func (m *TLKeyboardButtonUrlAuth) SetUrl(v string) { m.Data2.Url = v }
func (m *TLKeyboardButtonUrlAuth) GetUrl() string  { return m.Data2.Url }

func (m *TLKeyboardButtonUrlAuth) SetButtonId(v int32) { m.Data2.ButtonId = v }
func (m *TLKeyboardButtonUrlAuth) GetButtonId() int32  { return m.Data2.ButtonId }

func (m *TLInputKeyboardButtonUrlAuth) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputKeyboardButtonUrlAuth) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputKeyboardButtonUrlAuth) SetRequestWriteAccess(v bool) { m.Data2.RequestWriteAccess = v }
func (m *TLInputKeyboardButtonUrlAuth) GetRequestWriteAccess() bool {
	return m.Data2.RequestWriteAccess
}

func (m *TLInputKeyboardButtonUrlAuth) SetText(v string) { m.Data2.Text = v }
func (m *TLInputKeyboardButtonUrlAuth) GetText() string  { return m.Data2.Text }

func (m *TLInputKeyboardButtonUrlAuth) SetFwdText(v string) { m.Data2.FwdText = v }
func (m *TLInputKeyboardButtonUrlAuth) GetFwdText() string  { return m.Data2.FwdText }

func (m *TLInputKeyboardButtonUrlAuth) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputKeyboardButtonUrlAuth) GetUrl() string  { return m.Data2.Url }

func (m *TLInputKeyboardButtonUrlAuth) SetBot(v *InputUser) { m.Data2.Bot = v }
func (m *TLInputKeyboardButtonUrlAuth) GetBot() *InputUser  { return m.Data2.Bot }

func (m *TLKeyboardButtonRequestPoll) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLKeyboardButtonRequestPoll) GetFlags() int32  { return m.Data2.Flags }

func (m *TLKeyboardButtonRequestPoll) SetQuiz(v *Bool) { m.Data2.Quiz = v }
func (m *TLKeyboardButtonRequestPoll) GetQuiz() *Bool  { return m.Data2.Quiz }

func (m *TLKeyboardButtonRequestPoll) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonRequestPoll) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonRow) SetButtons(v []*KeyboardButton) { m.Data2.Buttons = v }
func (m *TLKeyboardButtonRow) GetButtons() []*KeyboardButton  { return m.Data2.Buttons }

func (m *TLLabeledPrice) SetLabel(v string) { m.Data2.Label = v }
func (m *TLLabeledPrice) GetLabel() string  { return m.Data2.Label }

func (m *TLLabeledPrice) SetAmount(v int64) { m.Data2.Amount = v }
func (m *TLLabeledPrice) GetAmount() int64  { return m.Data2.Amount }

func (m *TLLangPackDifference) SetLangCode(v string) { m.Data2.LangCode = v }
func (m *TLLangPackDifference) GetLangCode() string  { return m.Data2.LangCode }

func (m *TLLangPackDifference) SetFromVersion(v int32) { m.Data2.FromVersion = v }
func (m *TLLangPackDifference) GetFromVersion() int32  { return m.Data2.FromVersion }

func (m *TLLangPackDifference) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLLangPackDifference) GetVersion() int32  { return m.Data2.Version }

func (m *TLLangPackDifference) SetStrings(v []*LangPackString) { m.Data2.Strings = v }
func (m *TLLangPackDifference) GetStrings() []*LangPackString  { return m.Data2.Strings }

func (m *TLLangPackLanguage) SetName(v string) { m.Data2.Name = v }
func (m *TLLangPackLanguage) GetName() string  { return m.Data2.Name }

func (m *TLLangPackLanguage) SetNativeName(v string) { m.Data2.NativeName = v }
func (m *TLLangPackLanguage) GetNativeName() string  { return m.Data2.NativeName }

func (m *TLLangPackLanguage) SetLangCode(v string) { m.Data2.LangCode = v }
func (m *TLLangPackLanguage) GetLangCode() string  { return m.Data2.LangCode }

func (m *TLLangPackLanguage) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLLangPackLanguage) GetFlags() int32  { return m.Data2.Flags }

func (m *TLLangPackLanguage) SetOfficial(v bool) { m.Data2.Official = v }
func (m *TLLangPackLanguage) GetOfficial() bool  { return m.Data2.Official }

func (m *TLLangPackLanguage) SetRtl(v bool) { m.Data2.Rtl = v }
func (m *TLLangPackLanguage) GetRtl() bool  { return m.Data2.Rtl }

func (m *TLLangPackLanguage) SetBaseLangCode(v string) { m.Data2.BaseLangCode = v }
func (m *TLLangPackLanguage) GetBaseLangCode() string  { return m.Data2.BaseLangCode }

func (m *TLLangPackLanguage) SetPluralCode(v string) { m.Data2.PluralCode = v }
func (m *TLLangPackLanguage) GetPluralCode() string  { return m.Data2.PluralCode }

func (m *TLLangPackLanguage) SetStringsCount(v int32) { m.Data2.StringsCount = v }
func (m *TLLangPackLanguage) GetStringsCount() int32  { return m.Data2.StringsCount }

func (m *TLLangPackLanguage) SetTranslatedCount(v int32) { m.Data2.TranslatedCount = v }
func (m *TLLangPackLanguage) GetTranslatedCount() int32  { return m.Data2.TranslatedCount }

func (m *TLLangPackLanguage) SetTranslationsUrl(v string) { m.Data2.TranslationsUrl = v }
func (m *TLLangPackLanguage) GetTranslationsUrl() string  { return m.Data2.TranslationsUrl }

func (m *TLLangPackString) SetKey(v string) { m.Data2.Key = v }
func (m *TLLangPackString) GetKey() string  { return m.Data2.Key }

func (m *TLLangPackString) SetValue(v string) { m.Data2.Value = v }
func (m *TLLangPackString) GetValue() string  { return m.Data2.Value }

func (m *TLLangPackStringPluralized) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLLangPackStringPluralized) GetFlags() int32  { return m.Data2.Flags }

func (m *TLLangPackStringPluralized) SetKey(v string) { m.Data2.Key = v }
func (m *TLLangPackStringPluralized) GetKey() string  { return m.Data2.Key }

func (m *TLLangPackStringPluralized) SetZeroValue(v string) { m.Data2.ZeroValue = v }
func (m *TLLangPackStringPluralized) GetZeroValue() string  { return m.Data2.ZeroValue }

func (m *TLLangPackStringPluralized) SetOneValue(v string) { m.Data2.OneValue = v }
func (m *TLLangPackStringPluralized) GetOneValue() string  { return m.Data2.OneValue }

func (m *TLLangPackStringPluralized) SetTwoValue(v string) { m.Data2.TwoValue = v }
func (m *TLLangPackStringPluralized) GetTwoValue() string  { return m.Data2.TwoValue }

func (m *TLLangPackStringPluralized) SetFewValue(v string) { m.Data2.FewValue = v }
func (m *TLLangPackStringPluralized) GetFewValue() string  { return m.Data2.FewValue }

func (m *TLLangPackStringPluralized) SetManyValue(v string) { m.Data2.ManyValue = v }
func (m *TLLangPackStringPluralized) GetManyValue() string  { return m.Data2.ManyValue }

func (m *TLLangPackStringPluralized) SetOtherValue(v string) { m.Data2.OtherValue = v }
func (m *TLLangPackStringPluralized) GetOtherValue() string  { return m.Data2.OtherValue }

func (m *TLLangPackStringDeleted) SetKey(v string) { m.Data2.Key = v }
func (m *TLLangPackStringDeleted) GetKey() string  { return m.Data2.Key }

func (m *TLMaskCoords) SetN(v int32) { m.Data2.N = v }
func (m *TLMaskCoords) GetN() int32  { return m.Data2.N }

func (m *TLMaskCoords) SetX(v float64) { m.Data2.X = v }
func (m *TLMaskCoords) GetX() float64  { return m.Data2.X }

func (m *TLMaskCoords) SetY(v float64) { m.Data2.Y = v }
func (m *TLMaskCoords) GetY() float64  { return m.Data2.Y }

func (m *TLMaskCoords) SetZoom(v float64) { m.Data2.Zoom = v }
func (m *TLMaskCoords) GetZoom() float64  { return m.Data2.Zoom }

func (m *TLMessageEmpty) SetId(v int32) { m.Data2.Id = v }
func (m *TLMessageEmpty) GetId() int32  { return m.Data2.Id }

func (m *TLMessage) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessage) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessage) SetOut(v bool) { m.Data2.Out = v }
func (m *TLMessage) GetOut() bool  { return m.Data2.Out }

func (m *TLMessage) SetMentioned(v bool) { m.Data2.Mentioned = v }
func (m *TLMessage) GetMentioned() bool  { return m.Data2.Mentioned }

func (m *TLMessage) SetMediaUnread(v bool) { m.Data2.MediaUnread = v }
func (m *TLMessage) GetMediaUnread() bool  { return m.Data2.MediaUnread }

func (m *TLMessage) SetSilent(v bool) { m.Data2.Silent = v }
func (m *TLMessage) GetSilent() bool  { return m.Data2.Silent }

func (m *TLMessage) SetPost(v bool) { m.Data2.Post = v }
func (m *TLMessage) GetPost() bool  { return m.Data2.Post }

func (m *TLMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLMessage) GetId() int32  { return m.Data2.Id }

func (m *TLMessage) SetFromId90DDDC1171(v int32) { m.Data2.FromId90DDDC1171 = v }
func (m *TLMessage) GetFromId90DDDC1171() int32  { return m.Data2.FromId90DDDC1171 }

func (m *TLMessage) SetToId(v *Peer) { m.Data2.ToId = v }
func (m *TLMessage) GetToId() *Peer  { return m.Data2.ToId }

func (m *TLMessage) SetFwdFrom(v *MessageFwdHeader) { m.Data2.FwdFrom = v }
func (m *TLMessage) GetFwdFrom() *MessageFwdHeader  { return m.Data2.FwdFrom }

func (m *TLMessage) SetViaBotId(v int32) { m.Data2.ViaBotId = v }
func (m *TLMessage) GetViaBotId() int32  { return m.Data2.ViaBotId }

func (m *TLMessage) SetReplyToMsgId(v int32) { m.Data2.ReplyToMsgId = v }
func (m *TLMessage) GetReplyToMsgId() int32  { return m.Data2.ReplyToMsgId }

func (m *TLMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLMessage) SetMessage(v string) { m.Data2.Message = v }
func (m *TLMessage) GetMessage() string  { return m.Data2.Message }

func (m *TLMessage) SetMedia(v *MessageMedia) { m.Data2.Media = v }
func (m *TLMessage) GetMedia() *MessageMedia  { return m.Data2.Media }

func (m *TLMessage) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLMessage) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLMessage) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLMessage) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLMessage) SetViews(v int32) { m.Data2.Views = v }
func (m *TLMessage) GetViews() int32  { return m.Data2.Views }

func (m *TLMessage) SetEditDate(v int32) { m.Data2.EditDate = v }
func (m *TLMessage) GetEditDate() int32  { return m.Data2.EditDate }

func (m *TLMessage) SetPostAuthor(v string) { m.Data2.PostAuthor = v }
func (m *TLMessage) GetPostAuthor() string  { return m.Data2.PostAuthor }

func (m *TLMessage) SetUnread(v bool) { m.Data2.Unread = v }
func (m *TLMessage) GetUnread() bool  { return m.Data2.Unread }

func (m *TLMessage) SetGroupedId(v int64) { m.Data2.GroupedId = v }
func (m *TLMessage) GetGroupedId() int64  { return m.Data2.GroupedId }

func (m *TLMessage) SetFromScheduled(v bool) { m.Data2.FromScheduled = v }
func (m *TLMessage) GetFromScheduled() bool  { return m.Data2.FromScheduled }

func (m *TLMessage) SetLegacy(v bool) { m.Data2.Legacy = v }
func (m *TLMessage) GetLegacy() bool  { return m.Data2.Legacy }

func (m *TLMessage) SetEditHide(v bool) { m.Data2.EditHide = v }
func (m *TLMessage) GetEditHide() bool  { return m.Data2.EditHide }

func (m *TLMessage) SetRestrictionReason(v []*RestrictionReason) { m.Data2.RestrictionReason = v }
func (m *TLMessage) GetRestrictionReason() []*RestrictionReason  { return m.Data2.RestrictionReason }

func (m *TLMessage) SetPinned(v bool) { m.Data2.Pinned = v }
func (m *TLMessage) GetPinned() bool  { return m.Data2.Pinned }

func (m *TLMessage) SetFromId286FA604119(v *Peer) { m.Data2.FromId286FA604119 = v }
func (m *TLMessage) GetFromId286FA604119() *Peer  { return m.Data2.FromId286FA604119 }

func (m *TLMessage) SetPeerId(v *Peer) { m.Data2.PeerId = v }
func (m *TLMessage) GetPeerId() *Peer  { return m.Data2.PeerId }

func (m *TLMessage) SetReplyTo(v *MessageReplyHeader) { m.Data2.ReplyTo = v }
func (m *TLMessage) GetReplyTo() *MessageReplyHeader  { return m.Data2.ReplyTo }

func (m *TLMessage) SetForwards(v int32) { m.Data2.Forwards = v }
func (m *TLMessage) GetForwards() int32  { return m.Data2.Forwards }

func (m *TLMessage) SetReplies(v *MessageReplies) { m.Data2.Replies = v }
func (m *TLMessage) GetReplies() *MessageReplies  { return m.Data2.Replies }

func (m *TLMessageService) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageService) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageService) SetOut(v bool) { m.Data2.Out = v }
func (m *TLMessageService) GetOut() bool  { return m.Data2.Out }

func (m *TLMessageService) SetMentioned(v bool) { m.Data2.Mentioned = v }
func (m *TLMessageService) GetMentioned() bool  { return m.Data2.Mentioned }

func (m *TLMessageService) SetMediaUnread(v bool) { m.Data2.MediaUnread = v }
func (m *TLMessageService) GetMediaUnread() bool  { return m.Data2.MediaUnread }

func (m *TLMessageService) SetSilent(v bool) { m.Data2.Silent = v }
func (m *TLMessageService) GetSilent() bool  { return m.Data2.Silent }

func (m *TLMessageService) SetPost(v bool) { m.Data2.Post = v }
func (m *TLMessageService) GetPost() bool  { return m.Data2.Post }

func (m *TLMessageService) SetId(v int32) { m.Data2.Id = v }
func (m *TLMessageService) GetId() int32  { return m.Data2.Id }

func (m *TLMessageService) SetFromId90DDDC1171(v int32) { m.Data2.FromId90DDDC1171 = v }
func (m *TLMessageService) GetFromId90DDDC1171() int32  { return m.Data2.FromId90DDDC1171 }

func (m *TLMessageService) SetToId(v *Peer) { m.Data2.ToId = v }
func (m *TLMessageService) GetToId() *Peer  { return m.Data2.ToId }

func (m *TLMessageService) SetReplyToMsgId(v int32) { m.Data2.ReplyToMsgId = v }
func (m *TLMessageService) GetReplyToMsgId() int32  { return m.Data2.ReplyToMsgId }

func (m *TLMessageService) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessageService) GetDate() int32  { return m.Data2.Date }

func (m *TLMessageService) SetAction(v *MessageAction) { m.Data2.Action = v }
func (m *TLMessageService) GetAction() *MessageAction  { return m.Data2.Action }

func (m *TLMessageService) SetLegacy(v bool) { m.Data2.Legacy = v }
func (m *TLMessageService) GetLegacy() bool  { return m.Data2.Legacy }

func (m *TLMessageService) SetFromId286FA604119(v *Peer) { m.Data2.FromId286FA604119 = v }
func (m *TLMessageService) GetFromId286FA604119() *Peer  { return m.Data2.FromId286FA604119 }

func (m *TLMessageService) SetPeerId(v *Peer) { m.Data2.PeerId = v }
func (m *TLMessageService) GetPeerId() *Peer  { return m.Data2.PeerId }

func (m *TLMessageService) SetReplyTo(v *MessageReplyHeader) { m.Data2.ReplyTo = v }
func (m *TLMessageService) GetReplyTo() *MessageReplyHeader  { return m.Data2.ReplyTo }

func (m *TLMessageActionChatCreate) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageActionChatCreate) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageActionChatCreate) SetUsers(v []int32) { m.Data2.Users = v }
func (m *TLMessageActionChatCreate) GetUsers() []int32  { return m.Data2.Users }

func (m *TLMessageActionChatEditTitle) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageActionChatEditTitle) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageActionChatEditPhoto) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLMessageActionChatEditPhoto) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLMessageActionChatAddUser) SetUsers(v []int32) { m.Data2.Users = v }
func (m *TLMessageActionChatAddUser) GetUsers() []int32  { return m.Data2.Users }

func (m *TLMessageActionChatDeleteUser) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLMessageActionChatDeleteUser) GetUserId() int32  { return m.Data2.UserId }

func (m *TLMessageActionChatJoinedByLink) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLMessageActionChatJoinedByLink) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLMessageActionChannelCreate) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageActionChannelCreate) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageActionChatMigrateTo) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLMessageActionChatMigrateTo) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLMessageActionChannelMigrateFrom) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageActionChannelMigrateFrom) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageActionChannelMigrateFrom) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLMessageActionChannelMigrateFrom) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLMessageActionGameScore) SetGameId(v int64) { m.Data2.GameId = v }
func (m *TLMessageActionGameScore) GetGameId() int64  { return m.Data2.GameId }

func (m *TLMessageActionGameScore) SetScore(v int32) { m.Data2.Score = v }
func (m *TLMessageActionGameScore) GetScore() int32  { return m.Data2.Score }

func (m *TLMessageActionPaymentSentMe) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageActionPaymentSentMe) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageActionPaymentSentMe) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLMessageActionPaymentSentMe) GetCurrency() string  { return m.Data2.Currency }

func (m *TLMessageActionPaymentSentMe) SetTotalAmount(v int64) { m.Data2.TotalAmount = v }
func (m *TLMessageActionPaymentSentMe) GetTotalAmount() int64  { return m.Data2.TotalAmount }

func (m *TLMessageActionPaymentSentMe) SetPayload(v []byte) { m.Data2.Payload = v }
func (m *TLMessageActionPaymentSentMe) GetPayload() []byte  { return m.Data2.Payload }

func (m *TLMessageActionPaymentSentMe) SetInfo(v *PaymentRequestedInfo) { m.Data2.Info = v }
func (m *TLMessageActionPaymentSentMe) GetInfo() *PaymentRequestedInfo  { return m.Data2.Info }

func (m *TLMessageActionPaymentSentMe) SetShippingOptionId(v string) { m.Data2.ShippingOptionId = v }
func (m *TLMessageActionPaymentSentMe) GetShippingOptionId() string  { return m.Data2.ShippingOptionId }

func (m *TLMessageActionPaymentSentMe) SetCharge(v *PaymentCharge) { m.Data2.Charge = v }
func (m *TLMessageActionPaymentSentMe) GetCharge() *PaymentCharge  { return m.Data2.Charge }

func (m *TLMessageActionPaymentSent) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLMessageActionPaymentSent) GetCurrency() string  { return m.Data2.Currency }

func (m *TLMessageActionPaymentSent) SetTotalAmount(v int64) { m.Data2.TotalAmount = v }
func (m *TLMessageActionPaymentSent) GetTotalAmount() int64  { return m.Data2.TotalAmount }

func (m *TLMessageActionCustomAction) SetMessage(v string) { m.Data2.Message = v }
func (m *TLMessageActionCustomAction) GetMessage() string  { return m.Data2.Message }

func (m *TLMessageActionBotAllowed) SetDomain(v string) { m.Data2.Domain = v }
func (m *TLMessageActionBotAllowed) GetDomain() string  { return m.Data2.Domain }

func (m *TLMessageActionSecureValuesSentMe) SetValues(v []*SecureValue) { m.Data2.Values = v }
func (m *TLMessageActionSecureValuesSentMe) GetValues() []*SecureValue  { return m.Data2.Values }

func (m *TLMessageActionSecureValuesSentMe) SetCredentials(v *SecureCredentialsEncrypted) {
	m.Data2.Credentials = v
}
func (m *TLMessageActionSecureValuesSentMe) GetCredentials() *SecureCredentialsEncrypted {
	return m.Data2.Credentials
}

func (m *TLMessageActionSecureValuesSent) SetTypes(v []*SecureValueType) { m.Data2.Types = v }
func (m *TLMessageActionSecureValuesSent) GetTypes() []*SecureValueType  { return m.Data2.Types }

func (m *TLMessageActionContactSignUp) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageActionContactSignUp) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageActionContactSignUp) SetSilent(v bool) { m.Data2.Silent = v }
func (m *TLMessageActionContactSignUp) GetSilent() bool  { return m.Data2.Silent }

func (m *TLMessageActionPhoneCall) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageActionPhoneCall) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageActionPhoneCall) SetVideo(v bool) { m.Data2.Video = v }
func (m *TLMessageActionPhoneCall) GetVideo() bool  { return m.Data2.Video }

func (m *TLMessageActionPhoneCall) SetCallId(v int64) { m.Data2.CallId = v }
func (m *TLMessageActionPhoneCall) GetCallId() int64  { return m.Data2.CallId }

func (m *TLMessageActionPhoneCall) SetReason(v *PhoneCallDiscardReason) { m.Data2.Reason = v }
func (m *TLMessageActionPhoneCall) GetReason() *PhoneCallDiscardReason  { return m.Data2.Reason }

func (m *TLMessageActionPhoneCall) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLMessageActionPhoneCall) GetDuration() int32  { return m.Data2.Duration }

func (m *TLMessageActionGeoProximityReached) SetFromId(v *Peer) { m.Data2.FromId = v }
func (m *TLMessageActionGeoProximityReached) GetFromId() *Peer  { return m.Data2.FromId }

func (m *TLMessageActionGeoProximityReached) SetToId(v *Peer) { m.Data2.ToId = v }
func (m *TLMessageActionGeoProximityReached) GetToId() *Peer  { return m.Data2.ToId }

func (m *TLMessageActionGeoProximityReached) SetDistance(v int32) { m.Data2.Distance = v }
func (m *TLMessageActionGeoProximityReached) GetDistance() int32  { return m.Data2.Distance }

func (m *TLMessageActionGroupCall) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageActionGroupCall) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageActionGroupCall) SetCall(v *InputGroupCall) { m.Data2.Call = v }
func (m *TLMessageActionGroupCall) GetCall() *InputGroupCall  { return m.Data2.Call }

func (m *TLMessageActionGroupCall) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLMessageActionGroupCall) GetDuration() int32  { return m.Data2.Duration }

func (m *TLMessageActionInviteToGroupCall) SetCall(v *InputGroupCall) { m.Data2.Call = v }
func (m *TLMessageActionInviteToGroupCall) GetCall() *InputGroupCall  { return m.Data2.Call }

func (m *TLMessageActionInviteToGroupCall) SetUsers(v []int32) { m.Data2.Users = v }
func (m *TLMessageActionInviteToGroupCall) GetUsers() []int32  { return m.Data2.Users }

func (m *TLMessageEntityUnknown) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityUnknown) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityUnknown) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityUnknown) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityMention) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityMention) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityMention) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityMention) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityHashtag) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityHashtag) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityHashtag) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityHashtag) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityBotCommand) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityBotCommand) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityBotCommand) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityBotCommand) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityUrl) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityUrl) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityUrl) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityUrl) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityEmail) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityEmail) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityEmail) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityEmail) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityBold) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityBold) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityBold) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityBold) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityItalic) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityItalic) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityItalic) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityItalic) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityCode) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityCode) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityCode) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityCode) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityPre) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityPre) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityPre) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityPre) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityPre) SetLanguage(v string) { m.Data2.Language = v }
func (m *TLMessageEntityPre) GetLanguage() string  { return m.Data2.Language }

func (m *TLMessageEntityTextUrl) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityTextUrl) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityTextUrl) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityTextUrl) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityTextUrl) SetUrl(v string) { m.Data2.Url = v }
func (m *TLMessageEntityTextUrl) GetUrl() string  { return m.Data2.Url }

func (m *TLMessageEntityMentionName) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityMentionName) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityMentionName) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityMentionName) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityMentionName) SetUserId352DCA5871(v int32) { m.Data2.UserId352DCA5871 = v }
func (m *TLMessageEntityMentionName) GetUserId352DCA5871() int32  { return m.Data2.UserId352DCA5871 }

func (m *TLInputMessageEntityMentionName) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLInputMessageEntityMentionName) GetOffset() int32  { return m.Data2.Offset }

func (m *TLInputMessageEntityMentionName) SetLength(v int32) { m.Data2.Length = v }
func (m *TLInputMessageEntityMentionName) GetLength() int32  { return m.Data2.Length }

func (m *TLInputMessageEntityMentionName) SetUserId208E68C971(v *InputUser) {
	m.Data2.UserId208E68C971 = v
}
func (m *TLInputMessageEntityMentionName) GetUserId208E68C971() *InputUser {
	return m.Data2.UserId208E68C971
}

func (m *TLMessageEntityPhone) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityPhone) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityPhone) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityPhone) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityCashtag) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityCashtag) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityCashtag) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityCashtag) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityUnderline) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityUnderline) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityUnderline) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityUnderline) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityStrike) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityStrike) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityStrike) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityStrike) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityBlockquote) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityBlockquote) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityBlockquote) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityBlockquote) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityBankCard) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityBankCard) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityBankCard) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityBankCard) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageFwdHeader) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageFwdHeader) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageFwdHeader) SetFromIdFADFF4AC71(v int32) { m.Data2.FromIdFADFF4AC71 = v }
func (m *TLMessageFwdHeader) GetFromIdFADFF4AC71() int32  { return m.Data2.FromIdFADFF4AC71 }

func (m *TLMessageFwdHeader) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessageFwdHeader) GetDate() int32  { return m.Data2.Date }

func (m *TLMessageFwdHeader) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLMessageFwdHeader) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLMessageFwdHeader) SetChannelPost(v int32) { m.Data2.ChannelPost = v }
func (m *TLMessageFwdHeader) GetChannelPost() int32  { return m.Data2.ChannelPost }

func (m *TLMessageFwdHeader) SetPostAuthor(v string) { m.Data2.PostAuthor = v }
func (m *TLMessageFwdHeader) GetPostAuthor() string  { return m.Data2.PostAuthor }

func (m *TLMessageFwdHeader) SetSavedFromPeer(v *Peer) { m.Data2.SavedFromPeer = v }
func (m *TLMessageFwdHeader) GetSavedFromPeer() *Peer  { return m.Data2.SavedFromPeer }

func (m *TLMessageFwdHeader) SetSavedFromMsgId(v int32) { m.Data2.SavedFromMsgId = v }
func (m *TLMessageFwdHeader) GetSavedFromMsgId() int32  { return m.Data2.SavedFromMsgId }

func (m *TLMessageFwdHeader) SetFromName(v string) { m.Data2.FromName = v }
func (m *TLMessageFwdHeader) GetFromName() string  { return m.Data2.FromName }

func (m *TLMessageFwdHeader) SetPsaType(v string) { m.Data2.PsaType = v }
func (m *TLMessageFwdHeader) GetPsaType() string  { return m.Data2.PsaType }

func (m *TLMessageFwdHeader) SetFromId5F777DCE119(v *Peer) { m.Data2.FromId5F777DCE119 = v }
func (m *TLMessageFwdHeader) GetFromId5F777DCE119() *Peer  { return m.Data2.FromId5F777DCE119 }

func (m *TLMessageGroup) SetMinId(v int32) { m.Data2.MinId = v }
func (m *TLMessageGroup) GetMinId() int32  { return m.Data2.MinId }

func (m *TLMessageGroup) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLMessageGroup) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLMessageGroup) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessageGroup) GetCount() int32  { return m.Data2.Count }

func (m *TLMessageGroup) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessageGroup) GetDate() int32  { return m.Data2.Date }

func (m *TLMessageInteractionCounters) SetMsgId(v int32) { m.Data2.MsgId = v }
func (m *TLMessageInteractionCounters) GetMsgId() int32  { return m.Data2.MsgId }

func (m *TLMessageInteractionCounters) SetViews(v int32) { m.Data2.Views = v }
func (m *TLMessageInteractionCounters) GetViews() int32  { return m.Data2.Views }

func (m *TLMessageInteractionCounters) SetForwards(v int32) { m.Data2.Forwards = v }
func (m *TLMessageInteractionCounters) GetForwards() int32  { return m.Data2.Forwards }

func (m *TLMessageMediaPhoto) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageMediaPhoto) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageMediaPhoto) SetPhotoB5223B0F71(v *Photo) { m.Data2.PhotoB5223B0F71 = v }
func (m *TLMessageMediaPhoto) GetPhotoB5223B0F71() *Photo  { return m.Data2.PhotoB5223B0F71 }

func (m *TLMessageMediaPhoto) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLMessageMediaPhoto) GetCaption() string  { return m.Data2.Caption }

func (m *TLMessageMediaPhoto) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLMessageMediaPhoto) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLMessageMediaGeo) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLMessageMediaGeo) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLMessageMediaContact) SetPhoneNumber(v string) { m.Data2.PhoneNumber = v }
func (m *TLMessageMediaContact) GetPhoneNumber() string  { return m.Data2.PhoneNumber }

func (m *TLMessageMediaContact) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLMessageMediaContact) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLMessageMediaContact) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLMessageMediaContact) GetLastName() string  { return m.Data2.LastName }

func (m *TLMessageMediaContact) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLMessageMediaContact) GetUserId() int32  { return m.Data2.UserId }

func (m *TLMessageMediaContact) SetVcard(v string) { m.Data2.Vcard = v }
func (m *TLMessageMediaContact) GetVcard() string  { return m.Data2.Vcard }

func (m *TLMessageMediaDocument) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageMediaDocument) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageMediaDocument) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLMessageMediaDocument) GetDocument() *Document  { return m.Data2.Document }

func (m *TLMessageMediaDocument) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLMessageMediaDocument) GetCaption() string  { return m.Data2.Caption }

func (m *TLMessageMediaDocument) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLMessageMediaDocument) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLMessageMediaWebPage) SetWebpage(v *WebPage) { m.Data2.Webpage = v }
func (m *TLMessageMediaWebPage) GetWebpage() *WebPage  { return m.Data2.Webpage }

func (m *TLMessageMediaVenue) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLMessageMediaVenue) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLMessageMediaVenue) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageMediaVenue) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageMediaVenue) SetAddress(v string) { m.Data2.Address = v }
func (m *TLMessageMediaVenue) GetAddress() string  { return m.Data2.Address }

func (m *TLMessageMediaVenue) SetProvider(v string) { m.Data2.Provider = v }
func (m *TLMessageMediaVenue) GetProvider() string  { return m.Data2.Provider }

func (m *TLMessageMediaVenue) SetVenueId(v string) { m.Data2.VenueId = v }
func (m *TLMessageMediaVenue) GetVenueId() string  { return m.Data2.VenueId }

func (m *TLMessageMediaVenue) SetVenueType(v string) { m.Data2.VenueType = v }
func (m *TLMessageMediaVenue) GetVenueType() string  { return m.Data2.VenueType }

func (m *TLMessageMediaGame) SetGame(v *Game) { m.Data2.Game = v }
func (m *TLMessageMediaGame) GetGame() *Game  { return m.Data2.Game }

func (m *TLMessageMediaInvoice) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageMediaInvoice) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageMediaInvoice) SetShippingAddressRequested(v bool) {
	m.Data2.ShippingAddressRequested = v
}
func (m *TLMessageMediaInvoice) GetShippingAddressRequested() bool {
	return m.Data2.ShippingAddressRequested
}

func (m *TLMessageMediaInvoice) SetTest(v bool) { m.Data2.Test = v }
func (m *TLMessageMediaInvoice) GetTest() bool  { return m.Data2.Test }

func (m *TLMessageMediaInvoice) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageMediaInvoice) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageMediaInvoice) SetDescription(v string) { m.Data2.Description = v }
func (m *TLMessageMediaInvoice) GetDescription() string  { return m.Data2.Description }

func (m *TLMessageMediaInvoice) SetPhoto8455134771(v *WebDocument) { m.Data2.Photo8455134771 = v }
func (m *TLMessageMediaInvoice) GetPhoto8455134771() *WebDocument  { return m.Data2.Photo8455134771 }

func (m *TLMessageMediaInvoice) SetReceiptMsgId(v int32) { m.Data2.ReceiptMsgId = v }
func (m *TLMessageMediaInvoice) GetReceiptMsgId() int32  { return m.Data2.ReceiptMsgId }

func (m *TLMessageMediaInvoice) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLMessageMediaInvoice) GetCurrency() string  { return m.Data2.Currency }

func (m *TLMessageMediaInvoice) SetTotalAmount(v int64) { m.Data2.TotalAmount = v }
func (m *TLMessageMediaInvoice) GetTotalAmount() int64  { return m.Data2.TotalAmount }

func (m *TLMessageMediaInvoice) SetStartParam(v string) { m.Data2.StartParam = v }
func (m *TLMessageMediaInvoice) GetStartParam() string  { return m.Data2.StartParam }

func (m *TLMessageMediaGeoLive) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLMessageMediaGeoLive) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLMessageMediaGeoLive) SetPeriod(v int32) { m.Data2.Period = v }
func (m *TLMessageMediaGeoLive) GetPeriod() int32  { return m.Data2.Period }

func (m *TLMessageMediaGeoLive) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageMediaGeoLive) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageMediaGeoLive) SetHeading(v int32) { m.Data2.Heading = v }
func (m *TLMessageMediaGeoLive) GetHeading() int32  { return m.Data2.Heading }

func (m *TLMessageMediaGeoLive) SetProximityNotificationRadius(v int32) {
	m.Data2.ProximityNotificationRadius = v
}
func (m *TLMessageMediaGeoLive) GetProximityNotificationRadius() int32 {
	return m.Data2.ProximityNotificationRadius
}

func (m *TLMessageMediaPoll) SetPoll(v *Poll) { m.Data2.Poll = v }
func (m *TLMessageMediaPoll) GetPoll() *Poll  { return m.Data2.Poll }

func (m *TLMessageMediaPoll) SetResults(v *PollResults) { m.Data2.Results = v }
func (m *TLMessageMediaPoll) GetResults() *PollResults  { return m.Data2.Results }

func (m *TLMessageMediaDice) SetValue(v int32) { m.Data2.Value = v }
func (m *TLMessageMediaDice) GetValue() int32  { return m.Data2.Value }

func (m *TLMessageMediaDice) SetEmoticon(v string) { m.Data2.Emoticon = v }
func (m *TLMessageMediaDice) GetEmoticon() string  { return m.Data2.Emoticon }

func (m *TLMessageRange) SetMinId(v int32) { m.Data2.MinId = v }
func (m *TLMessageRange) GetMinId() int32  { return m.Data2.MinId }

func (m *TLMessageRange) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLMessageRange) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLMessageReplies) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageReplies) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageReplies) SetComments(v bool) { m.Data2.Comments = v }
func (m *TLMessageReplies) GetComments() bool  { return m.Data2.Comments }

func (m *TLMessageReplies) SetReplies(v int32) { m.Data2.Replies = v }
func (m *TLMessageReplies) GetReplies() int32  { return m.Data2.Replies }

func (m *TLMessageReplies) SetRepliesPts(v int32) { m.Data2.RepliesPts = v }
func (m *TLMessageReplies) GetRepliesPts() int32  { return m.Data2.RepliesPts }

func (m *TLMessageReplies) SetRecentRepliers(v []*Peer) { m.Data2.RecentRepliers = v }
func (m *TLMessageReplies) GetRecentRepliers() []*Peer  { return m.Data2.RecentRepliers }

func (m *TLMessageReplies) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLMessageReplies) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLMessageReplies) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLMessageReplies) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLMessageReplies) SetReadMaxId(v int32) { m.Data2.ReadMaxId = v }
func (m *TLMessageReplies) GetReadMaxId() int32  { return m.Data2.ReadMaxId }

func (m *TLMessageReplyHeader) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageReplyHeader) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageReplyHeader) SetReplyToMsgId(v int32) { m.Data2.ReplyToMsgId = v }
func (m *TLMessageReplyHeader) GetReplyToMsgId() int32  { return m.Data2.ReplyToMsgId }

func (m *TLMessageReplyHeader) SetReplyToPeerId(v *Peer) { m.Data2.ReplyToPeerId = v }
func (m *TLMessageReplyHeader) GetReplyToPeerId() *Peer  { return m.Data2.ReplyToPeerId }

func (m *TLMessageReplyHeader) SetReplyToTopId(v int32) { m.Data2.ReplyToTopId = v }
func (m *TLMessageReplyHeader) GetReplyToTopId() int32  { return m.Data2.ReplyToTopId }

func (m *TLMessageUserVote) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLMessageUserVote) GetUserId() int32  { return m.Data2.UserId }

func (m *TLMessageUserVote) SetOption(v []byte) { m.Data2.Option = v }
func (m *TLMessageUserVote) GetOption() []byte  { return m.Data2.Option }

func (m *TLMessageUserVote) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessageUserVote) GetDate() int32  { return m.Data2.Date }

func (m *TLMessageUserVoteInputOption) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLMessageUserVoteInputOption) GetUserId() int32  { return m.Data2.UserId }

func (m *TLMessageUserVoteInputOption) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessageUserVoteInputOption) GetDate() int32  { return m.Data2.Date }

func (m *TLMessageUserVoteMultiple) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLMessageUserVoteMultiple) GetUserId() int32  { return m.Data2.UserId }

func (m *TLMessageUserVoteMultiple) SetOptions(v [][]byte) { m.Data2.Options = v }
func (m *TLMessageUserVoteMultiple) GetOptions() [][]byte  { return m.Data2.Options }

func (m *TLMessageUserVoteMultiple) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessageUserVoteMultiple) GetDate() int32  { return m.Data2.Date }

func (m *TLMessageViews) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessageViews) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessageViews) SetViews(v int32) { m.Data2.Views = v }
func (m *TLMessageViews) GetViews() int32  { return m.Data2.Views }

func (m *TLMessageViews) SetForwards(v int32) { m.Data2.Forwards = v }
func (m *TLMessageViews) GetForwards() int32  { return m.Data2.Forwards }

func (m *TLMessageViews) SetReplies(v *MessageReplies) { m.Data2.Replies = v }
func (m *TLMessageViews) GetReplies() *MessageReplies  { return m.Data2.Replies }

func (m *TLInputMessagesFilterPhoneCalls) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLInputMessagesFilterPhoneCalls) GetFlags() int32  { return m.Data2.Flags }

func (m *TLInputMessagesFilterPhoneCalls) SetMissed(v bool) { m.Data2.Missed = v }
func (m *TLInputMessagesFilterPhoneCalls) GetMissed() bool  { return m.Data2.Missed }

func (m *TLMessagesAffectedHistory) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLMessagesAffectedHistory) GetPts() int32  { return m.Data2.Pts }

func (m *TLMessagesAffectedHistory) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLMessagesAffectedHistory) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLMessagesAffectedHistory) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessagesAffectedHistory) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessagesAffectedMessages) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLMessagesAffectedMessages) GetPts() int32  { return m.Data2.Pts }

func (m *TLMessagesAffectedMessages) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLMessagesAffectedMessages) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLMessagesAllStickers) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLMessagesAllStickers) GetHash() int32  { return m.Data2.Hash }

func (m *TLMessagesAllStickers) SetSets(v []*StickerSet) { m.Data2.Sets = v }
func (m *TLMessagesAllStickers) GetSets() []*StickerSet  { return m.Data2.Sets }

func (m *TLMessagesArchivedStickers) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesArchivedStickers) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesArchivedStickers) SetSets(v []*StickerSetCovered) { m.Data2.Sets = v }
func (m *TLMessagesArchivedStickers) GetSets() []*StickerSetCovered  { return m.Data2.Sets }

func (m *TLMessagesBotCallbackAnswer) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessagesBotCallbackAnswer) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessagesBotCallbackAnswer) SetAlert(v bool) { m.Data2.Alert = v }
func (m *TLMessagesBotCallbackAnswer) GetAlert() bool  { return m.Data2.Alert }

func (m *TLMessagesBotCallbackAnswer) SetHasUrl(v bool) { m.Data2.HasUrl = v }
func (m *TLMessagesBotCallbackAnswer) GetHasUrl() bool  { return m.Data2.HasUrl }

func (m *TLMessagesBotCallbackAnswer) SetMessage(v string) { m.Data2.Message = v }
func (m *TLMessagesBotCallbackAnswer) GetMessage() string  { return m.Data2.Message }

func (m *TLMessagesBotCallbackAnswer) SetUrl(v string) { m.Data2.Url = v }
func (m *TLMessagesBotCallbackAnswer) GetUrl() string  { return m.Data2.Url }

func (m *TLMessagesBotCallbackAnswer) SetCacheTime(v int32) { m.Data2.CacheTime = v }
func (m *TLMessagesBotCallbackAnswer) GetCacheTime() int32  { return m.Data2.CacheTime }

func (m *TLMessagesBotResults) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessagesBotResults) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessagesBotResults) SetGallery(v bool) { m.Data2.Gallery = v }
func (m *TLMessagesBotResults) GetGallery() bool  { return m.Data2.Gallery }

func (m *TLMessagesBotResults) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLMessagesBotResults) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLMessagesBotResults) SetNextOffset(v string) { m.Data2.NextOffset = v }
func (m *TLMessagesBotResults) GetNextOffset() string  { return m.Data2.NextOffset }

func (m *TLMessagesBotResults) SetSwitchPm(v *InlineBotSwitchPM) { m.Data2.SwitchPm = v }
func (m *TLMessagesBotResults) GetSwitchPm() *InlineBotSwitchPM  { return m.Data2.SwitchPm }

func (m *TLMessagesBotResults) SetResults(v []*BotInlineResult) { m.Data2.Results = v }
func (m *TLMessagesBotResults) GetResults() []*BotInlineResult  { return m.Data2.Results }

func (m *TLMessagesBotResults) SetCacheTime(v int32) { m.Data2.CacheTime = v }
func (m *TLMessagesBotResults) GetCacheTime() int32  { return m.Data2.CacheTime }

func (m *TLMessagesBotResults) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesBotResults) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesChatFull) SetFullChat(v *ChatFull) { m.Data2.FullChat = v }
func (m *TLMessagesChatFull) GetFullChat() *ChatFull  { return m.Data2.FullChat }

func (m *TLMessagesChatFull) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesChatFull) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesChatFull) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesChatFull) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesChats) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesChats) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesChatsSlice) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesChatsSlice) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesChatsSlice) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesChatsSlice) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesDhConfigNotModified) SetRandom(v []byte) { m.Data2.Random = v }
func (m *TLMessagesDhConfigNotModified) GetRandom() []byte  { return m.Data2.Random }

func (m *TLMessagesDhConfig) SetG(v int32) { m.Data2.G = v }
func (m *TLMessagesDhConfig) GetG() int32  { return m.Data2.G }

func (m *TLMessagesDhConfig) SetP(v []byte) { m.Data2.P = v }
func (m *TLMessagesDhConfig) GetP() []byte  { return m.Data2.P }

func (m *TLMessagesDhConfig) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLMessagesDhConfig) GetVersion() int32  { return m.Data2.Version }

func (m *TLMessagesDhConfig) SetRandom(v []byte) { m.Data2.Random = v }
func (m *TLMessagesDhConfig) GetRandom() []byte  { return m.Data2.Random }

func (m *TLMessagesDialogs) SetDialogs(v []*Dialog) { m.Data2.Dialogs = v }
func (m *TLMessagesDialogs) GetDialogs() []*Dialog  { return m.Data2.Dialogs }

func (m *TLMessagesDialogs) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesDialogs) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesDialogs) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesDialogs) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesDialogs) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesDialogs) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesDialogsSlice) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesDialogsSlice) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesDialogsSlice) SetDialogs(v []*Dialog) { m.Data2.Dialogs = v }
func (m *TLMessagesDialogsSlice) GetDialogs() []*Dialog  { return m.Data2.Dialogs }

func (m *TLMessagesDialogsSlice) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesDialogsSlice) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesDialogsSlice) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesDialogsSlice) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesDialogsSlice) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesDialogsSlice) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesDialogsNotModified) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesDialogsNotModified) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesDiscussionMessage) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessagesDiscussionMessage) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessagesDiscussionMessage) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesDiscussionMessage) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesDiscussionMessage) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLMessagesDiscussionMessage) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLMessagesDiscussionMessage) SetReadInboxMaxId(v int32) { m.Data2.ReadInboxMaxId = v }
func (m *TLMessagesDiscussionMessage) GetReadInboxMaxId() int32  { return m.Data2.ReadInboxMaxId }

func (m *TLMessagesDiscussionMessage) SetReadOutboxMaxId(v int32) { m.Data2.ReadOutboxMaxId = v }
func (m *TLMessagesDiscussionMessage) GetReadOutboxMaxId() int32  { return m.Data2.ReadOutboxMaxId }

func (m *TLMessagesDiscussionMessage) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesDiscussionMessage) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesDiscussionMessage) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesDiscussionMessage) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesFavedStickers) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLMessagesFavedStickers) GetHash() int32  { return m.Data2.Hash }

func (m *TLMessagesFavedStickers) SetPacks(v []*StickerPack) { m.Data2.Packs = v }
func (m *TLMessagesFavedStickers) GetPacks() []*StickerPack  { return m.Data2.Packs }

func (m *TLMessagesFavedStickers) SetStickers(v []*Document) { m.Data2.Stickers = v }
func (m *TLMessagesFavedStickers) GetStickers() []*Document  { return m.Data2.Stickers }

func (m *TLMessagesFeaturedStickersNotModified) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesFeaturedStickersNotModified) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesFeaturedStickers) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLMessagesFeaturedStickers) GetHash() int32  { return m.Data2.Hash }

func (m *TLMessagesFeaturedStickers) SetSets(v []*StickerSetCovered) { m.Data2.Sets = v }
func (m *TLMessagesFeaturedStickers) GetSets() []*StickerSetCovered  { return m.Data2.Sets }

func (m *TLMessagesFeaturedStickers) SetUnread(v []int64) { m.Data2.Unread = v }
func (m *TLMessagesFeaturedStickers) GetUnread() []int64  { return m.Data2.Unread }

func (m *TLMessagesFeaturedStickers) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesFeaturedStickers) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesFoundGifs) SetNextOffset(v int32) { m.Data2.NextOffset = v }
func (m *TLMessagesFoundGifs) GetNextOffset() int32  { return m.Data2.NextOffset }

func (m *TLMessagesFoundGifs) SetResults(v []*FoundGif) { m.Data2.Results = v }
func (m *TLMessagesFoundGifs) GetResults() []*FoundGif  { return m.Data2.Results }

func (m *TLMessagesFoundStickerSets) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLMessagesFoundStickerSets) GetHash() int32  { return m.Data2.Hash }

func (m *TLMessagesFoundStickerSets) SetSets(v []*StickerSetCovered) { m.Data2.Sets = v }
func (m *TLMessagesFoundStickerSets) GetSets() []*StickerSetCovered  { return m.Data2.Sets }

func (m *TLMessagesHighScores) SetScores(v []*HighScore) { m.Data2.Scores = v }
func (m *TLMessagesHighScores) GetScores() []*HighScore  { return m.Data2.Scores }

func (m *TLMessagesHighScores) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesHighScores) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesInactiveChats) SetDates(v []int32) { m.Data2.Dates = v }
func (m *TLMessagesInactiveChats) GetDates() []int32  { return m.Data2.Dates }

func (m *TLMessagesInactiveChats) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesInactiveChats) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesInactiveChats) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesInactiveChats) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesMessageEditData) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessagesMessageEditData) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessagesMessageEditData) SetCaption(v bool) { m.Data2.Caption = v }
func (m *TLMessagesMessageEditData) GetCaption() bool  { return m.Data2.Caption }

func (m *TLMessagesMessageViews) SetViews(v []*MessageViews) { m.Data2.Views = v }
func (m *TLMessagesMessageViews) GetViews() []*MessageViews  { return m.Data2.Views }

func (m *TLMessagesMessageViews) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesMessageViews) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesMessageViews) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesMessageViews) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesMessages) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesMessages) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesMessages) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesMessages) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesMessages) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesMessages) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesMessagesSlice) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesMessagesSlice) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesMessagesSlice) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesMessagesSlice) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesMessagesSlice) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesMessagesSlice) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesMessagesSlice) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesMessagesSlice) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesMessagesSlice) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessagesMessagesSlice) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessagesMessagesSlice) SetInexact(v bool) { m.Data2.Inexact = v }
func (m *TLMessagesMessagesSlice) GetInexact() bool  { return m.Data2.Inexact }

func (m *TLMessagesMessagesSlice) SetNextRate(v int32) { m.Data2.NextRate = v }
func (m *TLMessagesMessagesSlice) GetNextRate() int32  { return m.Data2.NextRate }

func (m *TLMessagesMessagesSlice) SetOffsetIdOffset(v int32) { m.Data2.OffsetIdOffset = v }
func (m *TLMessagesMessagesSlice) GetOffsetIdOffset() int32  { return m.Data2.OffsetIdOffset }

func (m *TLMessagesChannelMessages) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessagesChannelMessages) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessagesChannelMessages) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLMessagesChannelMessages) GetPts() int32  { return m.Data2.Pts }

func (m *TLMessagesChannelMessages) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesChannelMessages) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesChannelMessages) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesChannelMessages) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesChannelMessages) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesChannelMessages) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesChannelMessages) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesChannelMessages) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesChannelMessages) SetCollapsed(v []*MessageGroup) { m.Data2.Collapsed = v }
func (m *TLMessagesChannelMessages) GetCollapsed() []*MessageGroup  { return m.Data2.Collapsed }

func (m *TLMessagesChannelMessages) SetInexact(v bool) { m.Data2.Inexact = v }
func (m *TLMessagesChannelMessages) GetInexact() bool  { return m.Data2.Inexact }

func (m *TLMessagesChannelMessages) SetOffsetIdOffset(v int32) { m.Data2.OffsetIdOffset = v }
func (m *TLMessagesChannelMessages) GetOffsetIdOffset() int32  { return m.Data2.OffsetIdOffset }

func (m *TLMessagesMessagesNotModified) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesMessagesNotModified) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesPeerDialogs) SetDialogs(v []*Dialog) { m.Data2.Dialogs = v }
func (m *TLMessagesPeerDialogs) GetDialogs() []*Dialog  { return m.Data2.Dialogs }

func (m *TLMessagesPeerDialogs) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesPeerDialogs) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesPeerDialogs) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesPeerDialogs) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesPeerDialogs) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesPeerDialogs) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesPeerDialogs) SetState(v *Updates_State) { m.Data2.State = v }
func (m *TLMessagesPeerDialogs) GetState() *Updates_State  { return m.Data2.State }

func (m *TLMessagesRecentStickers) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLMessagesRecentStickers) GetHash() int32  { return m.Data2.Hash }

func (m *TLMessagesRecentStickers) SetStickers(v []*Document) { m.Data2.Stickers = v }
func (m *TLMessagesRecentStickers) GetStickers() []*Document  { return m.Data2.Stickers }

func (m *TLMessagesRecentStickers) SetPacks(v []*StickerPack) { m.Data2.Packs = v }
func (m *TLMessagesRecentStickers) GetPacks() []*StickerPack  { return m.Data2.Packs }

func (m *TLMessagesRecentStickers) SetDates(v []int32) { m.Data2.Dates = v }
func (m *TLMessagesRecentStickers) GetDates() []int32  { return m.Data2.Dates }

func (m *TLMessagesSavedGifs) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLMessagesSavedGifs) GetHash() int32  { return m.Data2.Hash }

func (m *TLMessagesSavedGifs) SetGifs(v []*Document) { m.Data2.Gifs = v }
func (m *TLMessagesSavedGifs) GetGifs() []*Document  { return m.Data2.Gifs }

func (m *TLMessagesSearchCounter) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessagesSearchCounter) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessagesSearchCounter) SetInexact(v bool) { m.Data2.Inexact = v }
func (m *TLMessagesSearchCounter) GetInexact() bool  { return m.Data2.Inexact }

func (m *TLMessagesSearchCounter) SetFilter(v *MessagesFilter) { m.Data2.Filter = v }
func (m *TLMessagesSearchCounter) GetFilter() *MessagesFilter  { return m.Data2.Filter }

func (m *TLMessagesSearchCounter) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesSearchCounter) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesSentEncryptedMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessagesSentEncryptedMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLMessagesSentEncryptedFile) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessagesSentEncryptedFile) GetDate() int32  { return m.Data2.Date }

func (m *TLMessagesSentEncryptedFile) SetFile(v *EncryptedFile) { m.Data2.File = v }
func (m *TLMessagesSentEncryptedFile) GetFile() *EncryptedFile  { return m.Data2.File }

func (m *TLMessagesStickerSet) SetSet(v *StickerSet) { m.Data2.Set = v }
func (m *TLMessagesStickerSet) GetSet() *StickerSet  { return m.Data2.Set }

func (m *TLMessagesStickerSet) SetPacks(v []*StickerPack) { m.Data2.Packs = v }
func (m *TLMessagesStickerSet) GetPacks() []*StickerPack  { return m.Data2.Packs }

func (m *TLMessagesStickerSet) SetDocuments(v []*Document) { m.Data2.Documents = v }
func (m *TLMessagesStickerSet) GetDocuments() []*Document  { return m.Data2.Documents }

func (m *TLMessagesStickerSetInstallResultArchive) SetSets(v []*StickerSetCovered) { m.Data2.Sets = v }
func (m *TLMessagesStickerSetInstallResultArchive) GetSets() []*StickerSetCovered {
	return m.Data2.Sets
}

func (m *TLMessagesStickers) SetHash8A8ECD3271(v string) { m.Data2.Hash8A8ECD3271 = v }
func (m *TLMessagesStickers) GetHash8A8ECD3271() string  { return m.Data2.Hash8A8ECD3271 }

func (m *TLMessagesStickers) SetStickers(v []*Document) { m.Data2.Stickers = v }
func (m *TLMessagesStickers) GetStickers() []*Document  { return m.Data2.Stickers }

func (m *TLMessagesStickers) SetHashE4599BBD85(v int32) { m.Data2.HashE4599BBD85 = v }
func (m *TLMessagesStickers) GetHashE4599BBD85() int32  { return m.Data2.HashE4599BBD85 }

func (m *TLMessagesVotesList) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLMessagesVotesList) GetFlags() int32  { return m.Data2.Flags }

func (m *TLMessagesVotesList) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesVotesList) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesVotesList) SetVotes(v []*MessageUserVote) { m.Data2.Votes = v }
func (m *TLMessagesVotesList) GetVotes() []*MessageUserVote  { return m.Data2.Votes }

func (m *TLMessagesVotesList) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesVotesList) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesVotesList) SetNextOffset(v string) { m.Data2.NextOffset = v }
func (m *TLMessagesVotesList) GetNextOffset() string  { return m.Data2.NextOffset }

func (m *TLMsgDetailedInfo) SetMsgId(v int64) { m.Data2.MsgId = v }
func (m *TLMsgDetailedInfo) GetMsgId() int64  { return m.Data2.MsgId }

func (m *TLMsgDetailedInfo) SetAnswerMsgId(v int64) { m.Data2.AnswerMsgId = v }
func (m *TLMsgDetailedInfo) GetAnswerMsgId() int64  { return m.Data2.AnswerMsgId }

func (m *TLMsgDetailedInfo) SetBytes(v int32) { m.Data2.Bytes = v }
func (m *TLMsgDetailedInfo) GetBytes() int32  { return m.Data2.Bytes }

func (m *TLMsgDetailedInfo) SetStatus(v int32) { m.Data2.Status = v }
func (m *TLMsgDetailedInfo) GetStatus() int32  { return m.Data2.Status }

func (m *TLMsgNewDetailedInfo) SetAnswerMsgId(v int64) { m.Data2.AnswerMsgId = v }
func (m *TLMsgNewDetailedInfo) GetAnswerMsgId() int64  { return m.Data2.AnswerMsgId }

func (m *TLMsgNewDetailedInfo) SetBytes(v int32) { m.Data2.Bytes = v }
func (m *TLMsgNewDetailedInfo) GetBytes() int32  { return m.Data2.Bytes }

func (m *TLMsgNewDetailedInfo) SetStatus(v int32) { m.Data2.Status = v }
func (m *TLMsgNewDetailedInfo) GetStatus() int32  { return m.Data2.Status }

func (m *TLMsgResendReq) SetMsgIds(v []int64) { m.Data2.MsgIds = v }
func (m *TLMsgResendReq) GetMsgIds() []int64  { return m.Data2.MsgIds }

func (m *TLMsgsAck) SetMsgIds(v []int64) { m.Data2.MsgIds = v }
func (m *TLMsgsAck) GetMsgIds() []int64  { return m.Data2.MsgIds }

func (m *TLMsgsAllInfo) SetMsgIds(v []int64) { m.Data2.MsgIds = v }
func (m *TLMsgsAllInfo) GetMsgIds() []int64  { return m.Data2.MsgIds }

func (m *TLMsgsAllInfo) SetInfo(v string) { m.Data2.Info = v }
func (m *TLMsgsAllInfo) GetInfo() string  { return m.Data2.Info }

func (m *TLMsgsStateInfo) SetReqMsgId(v int64) { m.Data2.ReqMsgId = v }
func (m *TLMsgsStateInfo) GetReqMsgId() int64  { return m.Data2.ReqMsgId }

func (m *TLMsgsStateInfo) SetInfo(v string) { m.Data2.Info = v }
func (m *TLMsgsStateInfo) GetInfo() string  { return m.Data2.Info }

func (m *TLMsgsStateReq) SetMsgIds(v []int64) { m.Data2.MsgIds = v }
func (m *TLMsgsStateReq) GetMsgIds() []int64  { return m.Data2.MsgIds }

func (m *TLNearestDc) SetCountry(v string) { m.Data2.Country = v }
func (m *TLNearestDc) GetCountry() string  { return m.Data2.Country }

func (m *TLNearestDc) SetThisDc(v int32) { m.Data2.ThisDc = v }
func (m *TLNearestDc) GetThisDc() int32  { return m.Data2.ThisDc }

func (m *TLNearestDc) SetNearestDc(v int32) { m.Data2.NearestDc = v }
func (m *TLNearestDc) GetNearestDc() int32  { return m.Data2.NearestDc }

func (m *TLNewSessionCreated) SetFirstMsgId(v int64) { m.Data2.FirstMsgId = v }
func (m *TLNewSessionCreated) GetFirstMsgId() int64  { return m.Data2.FirstMsgId }

func (m *TLNewSessionCreated) SetUniqueId(v int64) { m.Data2.UniqueId = v }
func (m *TLNewSessionCreated) GetUniqueId() int64  { return m.Data2.UniqueId }

func (m *TLNewSessionCreated) SetServerSalt(v int64) { m.Data2.ServerSalt = v }
func (m *TLNewSessionCreated) GetServerSalt() int64  { return m.Data2.ServerSalt }

func (m *TLNotifyPeer) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLNotifyPeer) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLPQInnerData) SetPq(v string) { m.Data2.Pq = v }
func (m *TLPQInnerData) GetPq() string  { return m.Data2.Pq }

func (m *TLPQInnerData) SetP(v string) { m.Data2.P = v }
func (m *TLPQInnerData) GetP() string  { return m.Data2.P }

func (m *TLPQInnerData) SetQ(v string) { m.Data2.Q = v }
func (m *TLPQInnerData) GetQ() string  { return m.Data2.Q }

func (m *TLPQInnerData) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLPQInnerData) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLPQInnerData) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLPQInnerData) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLPQInnerData) SetNewNonce(v []byte) { m.Data2.NewNonce = v }
func (m *TLPQInnerData) GetNewNonce() []byte  { return m.Data2.NewNonce }

func (m *TLPQInnerDataDc) SetPq(v string) { m.Data2.Pq = v }
func (m *TLPQInnerDataDc) GetPq() string  { return m.Data2.Pq }

func (m *TLPQInnerDataDc) SetP(v string) { m.Data2.P = v }
func (m *TLPQInnerDataDc) GetP() string  { return m.Data2.P }

func (m *TLPQInnerDataDc) SetQ(v string) { m.Data2.Q = v }
func (m *TLPQInnerDataDc) GetQ() string  { return m.Data2.Q }

func (m *TLPQInnerDataDc) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLPQInnerDataDc) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLPQInnerDataDc) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLPQInnerDataDc) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLPQInnerDataDc) SetNewNonce(v []byte) { m.Data2.NewNonce = v }
func (m *TLPQInnerDataDc) GetNewNonce() []byte  { return m.Data2.NewNonce }

func (m *TLPQInnerDataDc) SetDc(v int32) { m.Data2.Dc = v }
func (m *TLPQInnerDataDc) GetDc() int32  { return m.Data2.Dc }

func (m *TLPQInnerDataTemp) SetPq(v string) { m.Data2.Pq = v }
func (m *TLPQInnerDataTemp) GetPq() string  { return m.Data2.Pq }

func (m *TLPQInnerDataTemp) SetP(v string) { m.Data2.P = v }
func (m *TLPQInnerDataTemp) GetP() string  { return m.Data2.P }

func (m *TLPQInnerDataTemp) SetQ(v string) { m.Data2.Q = v }
func (m *TLPQInnerDataTemp) GetQ() string  { return m.Data2.Q }

func (m *TLPQInnerDataTemp) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLPQInnerDataTemp) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLPQInnerDataTemp) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLPQInnerDataTemp) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLPQInnerDataTemp) SetNewNonce(v []byte) { m.Data2.NewNonce = v }
func (m *TLPQInnerDataTemp) GetNewNonce() []byte  { return m.Data2.NewNonce }

func (m *TLPQInnerDataTemp) SetExpiresIn(v int32) { m.Data2.ExpiresIn = v }
func (m *TLPQInnerDataTemp) GetExpiresIn() int32  { return m.Data2.ExpiresIn }

func (m *TLPQInnerDataTempDc) SetPq(v string) { m.Data2.Pq = v }
func (m *TLPQInnerDataTempDc) GetPq() string  { return m.Data2.Pq }

func (m *TLPQInnerDataTempDc) SetP(v string) { m.Data2.P = v }
func (m *TLPQInnerDataTempDc) GetP() string  { return m.Data2.P }

func (m *TLPQInnerDataTempDc) SetQ(v string) { m.Data2.Q = v }
func (m *TLPQInnerDataTempDc) GetQ() string  { return m.Data2.Q }

func (m *TLPQInnerDataTempDc) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLPQInnerDataTempDc) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLPQInnerDataTempDc) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLPQInnerDataTempDc) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLPQInnerDataTempDc) SetNewNonce(v []byte) { m.Data2.NewNonce = v }
func (m *TLPQInnerDataTempDc) GetNewNonce() []byte  { return m.Data2.NewNonce }

func (m *TLPQInnerDataTempDc) SetDc(v int32) { m.Data2.Dc = v }
func (m *TLPQInnerDataTempDc) GetDc() int32  { return m.Data2.Dc }

func (m *TLPQInnerDataTempDc) SetExpiresIn(v int32) { m.Data2.ExpiresIn = v }
func (m *TLPQInnerDataTempDc) GetExpiresIn() int32  { return m.Data2.ExpiresIn }

func (m *TLPagePart) SetBlocks(v []*PageBlock) { m.Data2.Blocks = v }
func (m *TLPagePart) GetBlocks() []*PageBlock  { return m.Data2.Blocks }

func (m *TLPagePart) SetPhotos(v []*Photo) { m.Data2.Photos = v }
func (m *TLPagePart) GetPhotos() []*Photo  { return m.Data2.Photos }

func (m *TLPagePart) SetDocuments(v []*Document) { m.Data2.Documents = v }
func (m *TLPagePart) GetDocuments() []*Document  { return m.Data2.Documents }

func (m *TLPageFull) SetBlocks(v []*PageBlock) { m.Data2.Blocks = v }
func (m *TLPageFull) GetBlocks() []*PageBlock  { return m.Data2.Blocks }

func (m *TLPageFull) SetPhotos(v []*Photo) { m.Data2.Photos = v }
func (m *TLPageFull) GetPhotos() []*Photo  { return m.Data2.Photos }

func (m *TLPageFull) SetDocuments(v []*Document) { m.Data2.Documents = v }
func (m *TLPageFull) GetDocuments() []*Document  { return m.Data2.Documents }

func (m *TLPage) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPage) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPage) SetPart(v bool) { m.Data2.Part = v }
func (m *TLPage) GetPart() bool  { return m.Data2.Part }

func (m *TLPage) SetRtl(v bool) { m.Data2.Rtl = v }
func (m *TLPage) GetRtl() bool  { return m.Data2.Rtl }

func (m *TLPage) SetBlocks(v []*PageBlock) { m.Data2.Blocks = v }
func (m *TLPage) GetBlocks() []*PageBlock  { return m.Data2.Blocks }

func (m *TLPage) SetPhotos(v []*Photo) { m.Data2.Photos = v }
func (m *TLPage) GetPhotos() []*Photo  { return m.Data2.Photos }

func (m *TLPage) SetDocuments(v []*Document) { m.Data2.Documents = v }
func (m *TLPage) GetDocuments() []*Document  { return m.Data2.Documents }

func (m *TLPage) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPage) GetUrl() string  { return m.Data2.Url }

func (m *TLPage) SetV2(v bool) { m.Data2.V2 = v }
func (m *TLPage) GetV2() bool  { return m.Data2.V2 }

func (m *TLPage) SetViews(v int32) { m.Data2.Views = v }
func (m *TLPage) GetViews() int32  { return m.Data2.Views }

func (m *TLPageBlockTitle) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockTitle) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockSubtitle) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockSubtitle) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockAuthorDate) SetAuthorBAAFE5E071(v *RichText) { m.Data2.AuthorBAAFE5E071 = v }
func (m *TLPageBlockAuthorDate) GetAuthorBAAFE5E071() *RichText  { return m.Data2.AuthorBAAFE5E071 }

func (m *TLPageBlockAuthorDate) SetPublishedDate(v int32) { m.Data2.PublishedDate = v }
func (m *TLPageBlockAuthorDate) GetPublishedDate() int32  { return m.Data2.PublishedDate }

func (m *TLPageBlockHeader) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockHeader) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockSubheader) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockSubheader) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockParagraph) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockParagraph) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockPreformatted) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockPreformatted) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockPreformatted) SetLanguage(v string) { m.Data2.Language = v }
func (m *TLPageBlockPreformatted) GetLanguage() string  { return m.Data2.Language }

func (m *TLPageBlockFooter) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockFooter) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockAnchor) SetName(v string) { m.Data2.Name = v }
func (m *TLPageBlockAnchor) GetName() string  { return m.Data2.Name }

func (m *TLPageBlockList) SetOrdered(v *Bool) { m.Data2.Ordered = v }
func (m *TLPageBlockList) GetOrdered() *Bool  { return m.Data2.Ordered }

func (m *TLPageBlockList) SetItems3A58C7F471(v []*RichText) { m.Data2.Items3A58C7F471 = v }
func (m *TLPageBlockList) GetItems3A58C7F471() []*RichText  { return m.Data2.Items3A58C7F471 }

func (m *TLPageBlockList) SetItemsE4E8801188(v []*PageListItem) { m.Data2.ItemsE4E8801188 = v }
func (m *TLPageBlockList) GetItemsE4E8801188() []*PageListItem  { return m.Data2.ItemsE4E8801188 }

func (m *TLPageBlockBlockquote) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockBlockquote) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockBlockquote) SetCaption263D7C2671(v *RichText) { m.Data2.Caption263D7C2671 = v }
func (m *TLPageBlockBlockquote) GetCaption263D7C2671() *RichText  { return m.Data2.Caption263D7C2671 }

func (m *TLPageBlockPullquote) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockPullquote) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockPullquote) SetCaption263D7C2671(v *RichText) { m.Data2.Caption263D7C2671 = v }
func (m *TLPageBlockPullquote) GetCaption263D7C2671() *RichText  { return m.Data2.Caption263D7C2671 }

func (m *TLPageBlockPhoto) SetPhotoId(v int64) { m.Data2.PhotoId = v }
func (m *TLPageBlockPhoto) GetPhotoId() int64  { return m.Data2.PhotoId }

func (m *TLPageBlockPhoto) SetCaption263D7C2671(v *RichText) { m.Data2.Caption263D7C2671 = v }
func (m *TLPageBlockPhoto) GetCaption263D7C2671() *RichText  { return m.Data2.Caption263D7C2671 }

func (m *TLPageBlockPhoto) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPageBlockPhoto) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPageBlockPhoto) SetCaption1759C56088(v *PageCaption) { m.Data2.Caption1759C56088 = v }
func (m *TLPageBlockPhoto) GetCaption1759C56088() *PageCaption  { return m.Data2.Caption1759C56088 }

func (m *TLPageBlockPhoto) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPageBlockPhoto) GetUrl() string  { return m.Data2.Url }

func (m *TLPageBlockPhoto) SetWebpageId(v int64) { m.Data2.WebpageId = v }
func (m *TLPageBlockPhoto) GetWebpageId() int64  { return m.Data2.WebpageId }

func (m *TLPageBlockVideo) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPageBlockVideo) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPageBlockVideo) SetAutoplay(v bool) { m.Data2.Autoplay = v }
func (m *TLPageBlockVideo) GetAutoplay() bool  { return m.Data2.Autoplay }

func (m *TLPageBlockVideo) SetLoop(v bool) { m.Data2.Loop = v }
func (m *TLPageBlockVideo) GetLoop() bool  { return m.Data2.Loop }

func (m *TLPageBlockVideo) SetVideoId(v int64) { m.Data2.VideoId = v }
func (m *TLPageBlockVideo) GetVideoId() int64  { return m.Data2.VideoId }

func (m *TLPageBlockVideo) SetCaption263D7C2671(v *RichText) { m.Data2.Caption263D7C2671 = v }
func (m *TLPageBlockVideo) GetCaption263D7C2671() *RichText  { return m.Data2.Caption263D7C2671 }

func (m *TLPageBlockVideo) SetCaption1759C56088(v *PageCaption) { m.Data2.Caption1759C56088 = v }
func (m *TLPageBlockVideo) GetCaption1759C56088() *PageCaption  { return m.Data2.Caption1759C56088 }

func (m *TLPageBlockCover) SetCover(v *PageBlock) { m.Data2.Cover = v }
func (m *TLPageBlockCover) GetCover() *PageBlock  { return m.Data2.Cover }

func (m *TLPageBlockEmbed) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPageBlockEmbed) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPageBlockEmbed) SetFullWidth(v bool) { m.Data2.FullWidth = v }
func (m *TLPageBlockEmbed) GetFullWidth() bool  { return m.Data2.FullWidth }

func (m *TLPageBlockEmbed) SetAllowScrolling(v bool) { m.Data2.AllowScrolling = v }
func (m *TLPageBlockEmbed) GetAllowScrolling() bool  { return m.Data2.AllowScrolling }

func (m *TLPageBlockEmbed) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPageBlockEmbed) GetUrl() string  { return m.Data2.Url }

func (m *TLPageBlockEmbed) SetHtml(v string) { m.Data2.Html = v }
func (m *TLPageBlockEmbed) GetHtml() string  { return m.Data2.Html }

func (m *TLPageBlockEmbed) SetPosterPhotoId(v int64) { m.Data2.PosterPhotoId = v }
func (m *TLPageBlockEmbed) GetPosterPhotoId() int64  { return m.Data2.PosterPhotoId }

func (m *TLPageBlockEmbed) SetW(v int32) { m.Data2.W = v }
func (m *TLPageBlockEmbed) GetW() int32  { return m.Data2.W }

func (m *TLPageBlockEmbed) SetH(v int32) { m.Data2.H = v }
func (m *TLPageBlockEmbed) GetH() int32  { return m.Data2.H }

func (m *TLPageBlockEmbed) SetCaption263D7C2671(v *RichText) { m.Data2.Caption263D7C2671 = v }
func (m *TLPageBlockEmbed) GetCaption263D7C2671() *RichText  { return m.Data2.Caption263D7C2671 }

func (m *TLPageBlockEmbed) SetCaption1759C56088(v *PageCaption) { m.Data2.Caption1759C56088 = v }
func (m *TLPageBlockEmbed) GetCaption1759C56088() *PageCaption  { return m.Data2.Caption1759C56088 }

func (m *TLPageBlockEmbedPost) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPageBlockEmbedPost) GetUrl() string  { return m.Data2.Url }

func (m *TLPageBlockEmbedPost) SetWebpageId(v int64) { m.Data2.WebpageId = v }
func (m *TLPageBlockEmbedPost) GetWebpageId() int64  { return m.Data2.WebpageId }

func (m *TLPageBlockEmbedPost) SetAuthorPhotoId(v int64) { m.Data2.AuthorPhotoId = v }
func (m *TLPageBlockEmbedPost) GetAuthorPhotoId() int64  { return m.Data2.AuthorPhotoId }

func (m *TLPageBlockEmbedPost) SetAuthor292C7BE971(v string) { m.Data2.Author292C7BE971 = v }
func (m *TLPageBlockEmbedPost) GetAuthor292C7BE971() string  { return m.Data2.Author292C7BE971 }

func (m *TLPageBlockEmbedPost) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPageBlockEmbedPost) GetDate() int32  { return m.Data2.Date }

func (m *TLPageBlockEmbedPost) SetBlocks(v []*PageBlock) { m.Data2.Blocks = v }
func (m *TLPageBlockEmbedPost) GetBlocks() []*PageBlock  { return m.Data2.Blocks }

func (m *TLPageBlockEmbedPost) SetCaption263D7C2671(v *RichText) { m.Data2.Caption263D7C2671 = v }
func (m *TLPageBlockEmbedPost) GetCaption263D7C2671() *RichText  { return m.Data2.Caption263D7C2671 }

func (m *TLPageBlockEmbedPost) SetCaption1759C56088(v *PageCaption) { m.Data2.Caption1759C56088 = v }
func (m *TLPageBlockEmbedPost) GetCaption1759C56088() *PageCaption  { return m.Data2.Caption1759C56088 }

func (m *TLPageBlockCollage) SetItems8B31C4F71(v []*PageBlock) { m.Data2.Items8B31C4F71 = v }
func (m *TLPageBlockCollage) GetItems8B31C4F71() []*PageBlock  { return m.Data2.Items8B31C4F71 }

func (m *TLPageBlockCollage) SetCaption263D7C2671(v *RichText) { m.Data2.Caption263D7C2671 = v }
func (m *TLPageBlockCollage) GetCaption263D7C2671() *RichText  { return m.Data2.Caption263D7C2671 }

func (m *TLPageBlockCollage) SetCaption1759C56088(v *PageCaption) { m.Data2.Caption1759C56088 = v }
func (m *TLPageBlockCollage) GetCaption1759C56088() *PageCaption  { return m.Data2.Caption1759C56088 }

func (m *TLPageBlockSlideshow) SetItems8B31C4F71(v []*PageBlock) { m.Data2.Items8B31C4F71 = v }
func (m *TLPageBlockSlideshow) GetItems8B31C4F71() []*PageBlock  { return m.Data2.Items8B31C4F71 }

func (m *TLPageBlockSlideshow) SetCaption263D7C2671(v *RichText) { m.Data2.Caption263D7C2671 = v }
func (m *TLPageBlockSlideshow) GetCaption263D7C2671() *RichText  { return m.Data2.Caption263D7C2671 }

func (m *TLPageBlockSlideshow) SetCaption1759C56088(v *PageCaption) { m.Data2.Caption1759C56088 = v }
func (m *TLPageBlockSlideshow) GetCaption1759C56088() *PageCaption  { return m.Data2.Caption1759C56088 }

func (m *TLPageBlockChannel) SetChannel(v *Chat) { m.Data2.Channel = v }
func (m *TLPageBlockChannel) GetChannel() *Chat  { return m.Data2.Channel }

func (m *TLPageBlockAudio) SetAudioId(v int64) { m.Data2.AudioId = v }
func (m *TLPageBlockAudio) GetAudioId() int64  { return m.Data2.AudioId }

func (m *TLPageBlockAudio) SetCaption263D7C2671(v *RichText) { m.Data2.Caption263D7C2671 = v }
func (m *TLPageBlockAudio) GetCaption263D7C2671() *RichText  { return m.Data2.Caption263D7C2671 }

func (m *TLPageBlockAudio) SetCaption1759C56088(v *PageCaption) { m.Data2.Caption1759C56088 = v }
func (m *TLPageBlockAudio) GetCaption1759C56088() *PageCaption  { return m.Data2.Caption1759C56088 }

func (m *TLPageBlockKicker) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockKicker) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockTable) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPageBlockTable) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPageBlockTable) SetBordered(v bool) { m.Data2.Bordered = v }
func (m *TLPageBlockTable) GetBordered() bool  { return m.Data2.Bordered }

func (m *TLPageBlockTable) SetStriped(v bool) { m.Data2.Striped = v }
func (m *TLPageBlockTable) GetStriped() bool  { return m.Data2.Striped }

func (m *TLPageBlockTable) SetTitle(v *RichText) { m.Data2.Title = v }
func (m *TLPageBlockTable) GetTitle() *RichText  { return m.Data2.Title }

func (m *TLPageBlockTable) SetRows(v []*PageTableRow) { m.Data2.Rows = v }
func (m *TLPageBlockTable) GetRows() []*PageTableRow  { return m.Data2.Rows }

func (m *TLPageBlockOrderedList) SetItems9A8AE1E188(v []*PageListOrderedItem) {
	m.Data2.Items9A8AE1E188 = v
}
func (m *TLPageBlockOrderedList) GetItems9A8AE1E188() []*PageListOrderedItem {
	return m.Data2.Items9A8AE1E188
}

func (m *TLPageBlockDetails) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPageBlockDetails) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPageBlockDetails) SetOpen(v bool) { m.Data2.Open = v }
func (m *TLPageBlockDetails) GetOpen() bool  { return m.Data2.Open }

func (m *TLPageBlockDetails) SetBlocks(v []*PageBlock) { m.Data2.Blocks = v }
func (m *TLPageBlockDetails) GetBlocks() []*PageBlock  { return m.Data2.Blocks }

func (m *TLPageBlockDetails) SetTitle(v *RichText) { m.Data2.Title = v }
func (m *TLPageBlockDetails) GetTitle() *RichText  { return m.Data2.Title }

func (m *TLPageBlockRelatedArticles) SetTitle(v *RichText) { m.Data2.Title = v }
func (m *TLPageBlockRelatedArticles) GetTitle() *RichText  { return m.Data2.Title }

func (m *TLPageBlockRelatedArticles) SetArticles(v []*PageRelatedArticle) { m.Data2.Articles = v }
func (m *TLPageBlockRelatedArticles) GetArticles() []*PageRelatedArticle  { return m.Data2.Articles }

func (m *TLPageBlockMap) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLPageBlockMap) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLPageBlockMap) SetZoom(v int32) { m.Data2.Zoom = v }
func (m *TLPageBlockMap) GetZoom() int32  { return m.Data2.Zoom }

func (m *TLPageBlockMap) SetW(v int32) { m.Data2.W = v }
func (m *TLPageBlockMap) GetW() int32  { return m.Data2.W }

func (m *TLPageBlockMap) SetH(v int32) { m.Data2.H = v }
func (m *TLPageBlockMap) GetH() int32  { return m.Data2.H }

func (m *TLPageBlockMap) SetCaption1759C56088(v *PageCaption) { m.Data2.Caption1759C56088 = v }
func (m *TLPageBlockMap) GetCaption1759C56088() *PageCaption  { return m.Data2.Caption1759C56088 }

func (m *TLPageCaption) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageCaption) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageCaption) SetCredit(v *RichText) { m.Data2.Credit = v }
func (m *TLPageCaption) GetCredit() *RichText  { return m.Data2.Credit }

func (m *TLPageListItemText) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageListItemText) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageListItemBlocks) SetBlocks(v []*PageBlock) { m.Data2.Blocks = v }
func (m *TLPageListItemBlocks) GetBlocks() []*PageBlock  { return m.Data2.Blocks }

func (m *TLPageListOrderedItemText) SetNum(v string) { m.Data2.Num = v }
func (m *TLPageListOrderedItemText) GetNum() string  { return m.Data2.Num }

func (m *TLPageListOrderedItemText) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageListOrderedItemText) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageListOrderedItemBlocks) SetNum(v string) { m.Data2.Num = v }
func (m *TLPageListOrderedItemBlocks) GetNum() string  { return m.Data2.Num }

func (m *TLPageListOrderedItemBlocks) SetBlocks(v []*PageBlock) { m.Data2.Blocks = v }
func (m *TLPageListOrderedItemBlocks) GetBlocks() []*PageBlock  { return m.Data2.Blocks }

func (m *TLPageRelatedArticle) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPageRelatedArticle) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPageRelatedArticle) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPageRelatedArticle) GetUrl() string  { return m.Data2.Url }

func (m *TLPageRelatedArticle) SetWebpageId(v int64) { m.Data2.WebpageId = v }
func (m *TLPageRelatedArticle) GetWebpageId() int64  { return m.Data2.WebpageId }

func (m *TLPageRelatedArticle) SetTitle(v string) { m.Data2.Title = v }
func (m *TLPageRelatedArticle) GetTitle() string  { return m.Data2.Title }

func (m *TLPageRelatedArticle) SetDescription(v string) { m.Data2.Description = v }
func (m *TLPageRelatedArticle) GetDescription() string  { return m.Data2.Description }

func (m *TLPageRelatedArticle) SetPhotoId(v int64) { m.Data2.PhotoId = v }
func (m *TLPageRelatedArticle) GetPhotoId() int64  { return m.Data2.PhotoId }

func (m *TLPageRelatedArticle) SetAuthor(v string) { m.Data2.Author = v }
func (m *TLPageRelatedArticle) GetAuthor() string  { return m.Data2.Author }

func (m *TLPageRelatedArticle) SetPublishedDate(v int32) { m.Data2.PublishedDate = v }
func (m *TLPageRelatedArticle) GetPublishedDate() int32  { return m.Data2.PublishedDate }

func (m *TLPageTableCell) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPageTableCell) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPageTableCell) SetHeader(v bool) { m.Data2.Header = v }
func (m *TLPageTableCell) GetHeader() bool  { return m.Data2.Header }

func (m *TLPageTableCell) SetAlignCenter(v bool) { m.Data2.AlignCenter = v }
func (m *TLPageTableCell) GetAlignCenter() bool  { return m.Data2.AlignCenter }

func (m *TLPageTableCell) SetAlignRight(v bool) { m.Data2.AlignRight = v }
func (m *TLPageTableCell) GetAlignRight() bool  { return m.Data2.AlignRight }

func (m *TLPageTableCell) SetValignMiddle(v bool) { m.Data2.ValignMiddle = v }
func (m *TLPageTableCell) GetValignMiddle() bool  { return m.Data2.ValignMiddle }

func (m *TLPageTableCell) SetValignBottom(v bool) { m.Data2.ValignBottom = v }
func (m *TLPageTableCell) GetValignBottom() bool  { return m.Data2.ValignBottom }

func (m *TLPageTableCell) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageTableCell) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageTableCell) SetColspan(v int32) { m.Data2.Colspan = v }
func (m *TLPageTableCell) GetColspan() int32  { return m.Data2.Colspan }

func (m *TLPageTableCell) SetRowspan(v int32) { m.Data2.Rowspan = v }
func (m *TLPageTableCell) GetRowspan() int32  { return m.Data2.Rowspan }

func (m *TLPageTableRow) SetCells(v []*PageTableCell) { m.Data2.Cells = v }
func (m *TLPageTableRow) GetCells() []*PageTableCell  { return m.Data2.Cells }

func (m *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow) SetSalt1(v []byte) {
	m.Data2.Salt1 = v
}
func (m *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow) GetSalt1() []byte {
	return m.Data2.Salt1
}

func (m *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow) SetSalt2(v []byte) {
	m.Data2.Salt2 = v
}
func (m *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow) GetSalt2() []byte {
	return m.Data2.Salt2
}

func (m *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow) SetG(v int32) {
	m.Data2.G = v
}
func (m *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow) GetG() int32 {
	return m.Data2.G
}

func (m *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow) SetP(v []byte) {
	m.Data2.P = v
}
func (m *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow) GetP() []byte {
	return m.Data2.P
}

func (m *TLPaymentCharge) SetId(v string) { m.Data2.Id = v }
func (m *TLPaymentCharge) GetId() string  { return m.Data2.Id }

func (m *TLPaymentCharge) SetProviderChargeId(v string) { m.Data2.ProviderChargeId = v }
func (m *TLPaymentCharge) GetProviderChargeId() string  { return m.Data2.ProviderChargeId }

func (m *TLPaymentRequestedInfo) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPaymentRequestedInfo) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPaymentRequestedInfo) SetName(v string) { m.Data2.Name = v }
func (m *TLPaymentRequestedInfo) GetName() string  { return m.Data2.Name }

func (m *TLPaymentRequestedInfo) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLPaymentRequestedInfo) GetPhone() string  { return m.Data2.Phone }

func (m *TLPaymentRequestedInfo) SetEmail(v string) { m.Data2.Email = v }
func (m *TLPaymentRequestedInfo) GetEmail() string  { return m.Data2.Email }

func (m *TLPaymentRequestedInfo) SetShippingAddress(v *PostAddress) { m.Data2.ShippingAddress = v }
func (m *TLPaymentRequestedInfo) GetShippingAddress() *PostAddress  { return m.Data2.ShippingAddress }

func (m *TLPaymentSavedCredentialsCard) SetId(v string) { m.Data2.Id = v }
func (m *TLPaymentSavedCredentialsCard) GetId() string  { return m.Data2.Id }

func (m *TLPaymentSavedCredentialsCard) SetTitle(v string) { m.Data2.Title = v }
func (m *TLPaymentSavedCredentialsCard) GetTitle() string  { return m.Data2.Title }

func (m *TLPaymentsBankCardData) SetTitle(v string) { m.Data2.Title = v }
func (m *TLPaymentsBankCardData) GetTitle() string  { return m.Data2.Title }

func (m *TLPaymentsBankCardData) SetOpenUrls(v []*BankCardOpenUrl) { m.Data2.OpenUrls = v }
func (m *TLPaymentsBankCardData) GetOpenUrls() []*BankCardOpenUrl  { return m.Data2.OpenUrls }

func (m *TLPaymentsPaymentForm) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPaymentsPaymentForm) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPaymentsPaymentForm) SetCanSaveCredentials(v bool) { m.Data2.CanSaveCredentials = v }
func (m *TLPaymentsPaymentForm) GetCanSaveCredentials() bool  { return m.Data2.CanSaveCredentials }

func (m *TLPaymentsPaymentForm) SetPasswordMissing(v bool) { m.Data2.PasswordMissing = v }
func (m *TLPaymentsPaymentForm) GetPasswordMissing() bool  { return m.Data2.PasswordMissing }

func (m *TLPaymentsPaymentForm) SetBotId(v int32) { m.Data2.BotId = v }
func (m *TLPaymentsPaymentForm) GetBotId() int32  { return m.Data2.BotId }

func (m *TLPaymentsPaymentForm) SetInvoice(v *Invoice) { m.Data2.Invoice = v }
func (m *TLPaymentsPaymentForm) GetInvoice() *Invoice  { return m.Data2.Invoice }

func (m *TLPaymentsPaymentForm) SetProviderId(v int32) { m.Data2.ProviderId = v }
func (m *TLPaymentsPaymentForm) GetProviderId() int32  { return m.Data2.ProviderId }

func (m *TLPaymentsPaymentForm) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPaymentsPaymentForm) GetUrl() string  { return m.Data2.Url }

func (m *TLPaymentsPaymentForm) SetNativeProvider(v string) { m.Data2.NativeProvider = v }
func (m *TLPaymentsPaymentForm) GetNativeProvider() string  { return m.Data2.NativeProvider }

func (m *TLPaymentsPaymentForm) SetNativeParams(v *DataJSON) { m.Data2.NativeParams = v }
func (m *TLPaymentsPaymentForm) GetNativeParams() *DataJSON  { return m.Data2.NativeParams }

func (m *TLPaymentsPaymentForm) SetSavedInfo(v *PaymentRequestedInfo) { m.Data2.SavedInfo = v }
func (m *TLPaymentsPaymentForm) GetSavedInfo() *PaymentRequestedInfo  { return m.Data2.SavedInfo }

func (m *TLPaymentsPaymentForm) SetSavedCredentials(v *PaymentSavedCredentials) {
	m.Data2.SavedCredentials = v
}
func (m *TLPaymentsPaymentForm) GetSavedCredentials() *PaymentSavedCredentials {
	return m.Data2.SavedCredentials
}

func (m *TLPaymentsPaymentForm) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPaymentsPaymentForm) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPaymentsPaymentReceipt) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPaymentsPaymentReceipt) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPaymentsPaymentReceipt) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPaymentsPaymentReceipt) GetDate() int32  { return m.Data2.Date }

func (m *TLPaymentsPaymentReceipt) SetBotId(v int32) { m.Data2.BotId = v }
func (m *TLPaymentsPaymentReceipt) GetBotId() int32  { return m.Data2.BotId }

func (m *TLPaymentsPaymentReceipt) SetInvoice(v *Invoice) { m.Data2.Invoice = v }
func (m *TLPaymentsPaymentReceipt) GetInvoice() *Invoice  { return m.Data2.Invoice }

func (m *TLPaymentsPaymentReceipt) SetProviderId(v int32) { m.Data2.ProviderId = v }
func (m *TLPaymentsPaymentReceipt) GetProviderId() int32  { return m.Data2.ProviderId }

func (m *TLPaymentsPaymentReceipt) SetInfo(v *PaymentRequestedInfo) { m.Data2.Info = v }
func (m *TLPaymentsPaymentReceipt) GetInfo() *PaymentRequestedInfo  { return m.Data2.Info }

func (m *TLPaymentsPaymentReceipt) SetShipping(v *ShippingOption) { m.Data2.Shipping = v }
func (m *TLPaymentsPaymentReceipt) GetShipping() *ShippingOption  { return m.Data2.Shipping }

func (m *TLPaymentsPaymentReceipt) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLPaymentsPaymentReceipt) GetCurrency() string  { return m.Data2.Currency }

func (m *TLPaymentsPaymentReceipt) SetTotalAmount(v int64) { m.Data2.TotalAmount = v }
func (m *TLPaymentsPaymentReceipt) GetTotalAmount() int64  { return m.Data2.TotalAmount }

func (m *TLPaymentsPaymentReceipt) SetCredentialsTitle(v string) { m.Data2.CredentialsTitle = v }
func (m *TLPaymentsPaymentReceipt) GetCredentialsTitle() string  { return m.Data2.CredentialsTitle }

func (m *TLPaymentsPaymentReceipt) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPaymentsPaymentReceipt) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPaymentsPaymentResult) SetUpdates(v *Updates) { m.Data2.Updates = v }
func (m *TLPaymentsPaymentResult) GetUpdates() *Updates  { return m.Data2.Updates }

func (m *TLPaymentsPaymentVerficationNeeded) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPaymentsPaymentVerficationNeeded) GetUrl() string  { return m.Data2.Url }

func (m *TLPaymentsPaymentVerificationNeeded) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPaymentsPaymentVerificationNeeded) GetUrl() string  { return m.Data2.Url }

func (m *TLPaymentsSavedInfo) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPaymentsSavedInfo) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPaymentsSavedInfo) SetHasSavedCredentials(v bool) { m.Data2.HasSavedCredentials = v }
func (m *TLPaymentsSavedInfo) GetHasSavedCredentials() bool  { return m.Data2.HasSavedCredentials }

func (m *TLPaymentsSavedInfo) SetSavedInfo(v *PaymentRequestedInfo) { m.Data2.SavedInfo = v }
func (m *TLPaymentsSavedInfo) GetSavedInfo() *PaymentRequestedInfo  { return m.Data2.SavedInfo }

func (m *TLPaymentsValidatedRequestedInfo) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPaymentsValidatedRequestedInfo) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPaymentsValidatedRequestedInfo) SetId(v string) { m.Data2.Id = v }
func (m *TLPaymentsValidatedRequestedInfo) GetId() string  { return m.Data2.Id }

func (m *TLPaymentsValidatedRequestedInfo) SetShippingOptions(v []*ShippingOption) {
	m.Data2.ShippingOptions = v
}
func (m *TLPaymentsValidatedRequestedInfo) GetShippingOptions() []*ShippingOption {
	return m.Data2.ShippingOptions
}

func (m *TLPeerUser) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLPeerUser) GetUserId() int32  { return m.Data2.UserId }

func (m *TLPeerChat) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLPeerChat) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLPeerChannel) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLPeerChannel) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLPeerBlocked) SetPeerId(v *Peer) { m.Data2.PeerId = v }
func (m *TLPeerBlocked) GetPeerId() *Peer  { return m.Data2.PeerId }

func (m *TLPeerBlocked) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPeerBlocked) GetDate() int32  { return m.Data2.Date }

func (m *TLPeerLocated) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLPeerLocated) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLPeerLocated) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLPeerLocated) GetExpires() int32  { return m.Data2.Expires }

func (m *TLPeerLocated) SetDistance(v int32) { m.Data2.Distance = v }
func (m *TLPeerLocated) GetDistance() int32  { return m.Data2.Distance }

func (m *TLPeerSelfLocated) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLPeerSelfLocated) GetExpires() int32  { return m.Data2.Expires }

func (m *TLPeerNotifySettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPeerNotifySettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPeerNotifySettings) SetShowPreviews9ACDA4C071(v bool) { m.Data2.ShowPreviews9ACDA4C071 = v }
func (m *TLPeerNotifySettings) GetShowPreviews9ACDA4C071() bool {
	return m.Data2.ShowPreviews9ACDA4C071
}

func (m *TLPeerNotifySettings) SetSilent9ACDA4C071(v bool) { m.Data2.Silent9ACDA4C071 = v }
func (m *TLPeerNotifySettings) GetSilent9ACDA4C071() bool  { return m.Data2.Silent9ACDA4C071 }

func (m *TLPeerNotifySettings) SetMuteUntil(v int32) { m.Data2.MuteUntil = v }
func (m *TLPeerNotifySettings) GetMuteUntil() int32  { return m.Data2.MuteUntil }

func (m *TLPeerNotifySettings) SetSound(v string) { m.Data2.Sound = v }
func (m *TLPeerNotifySettings) GetSound() string  { return m.Data2.Sound }

func (m *TLPeerNotifySettings) SetShowPreviewsAF509D2085(v *Bool) { m.Data2.ShowPreviewsAF509D2085 = v }
func (m *TLPeerNotifySettings) GetShowPreviewsAF509D2085() *Bool {
	return m.Data2.ShowPreviewsAF509D2085
}

func (m *TLPeerNotifySettings) SetSilentAF509D2085(v *Bool) { m.Data2.SilentAF509D2085 = v }
func (m *TLPeerNotifySettings) GetSilentAF509D2085() *Bool  { return m.Data2.SilentAF509D2085 }

func (m *TLPeerSettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPeerSettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPeerSettings) SetReportSpam(v bool) { m.Data2.ReportSpam = v }
func (m *TLPeerSettings) GetReportSpam() bool  { return m.Data2.ReportSpam }

func (m *TLPeerSettings) SetAddContact(v bool) { m.Data2.AddContact = v }
func (m *TLPeerSettings) GetAddContact() bool  { return m.Data2.AddContact }

func (m *TLPeerSettings) SetBlockContact(v bool) { m.Data2.BlockContact = v }
func (m *TLPeerSettings) GetBlockContact() bool  { return m.Data2.BlockContact }

func (m *TLPeerSettings) SetShareContact(v bool) { m.Data2.ShareContact = v }
func (m *TLPeerSettings) GetShareContact() bool  { return m.Data2.ShareContact }

func (m *TLPeerSettings) SetNeedContactsException(v bool) { m.Data2.NeedContactsException = v }
func (m *TLPeerSettings) GetNeedContactsException() bool  { return m.Data2.NeedContactsException }

func (m *TLPeerSettings) SetReportGeo(v bool) { m.Data2.ReportGeo = v }
func (m *TLPeerSettings) GetReportGeo() bool  { return m.Data2.ReportGeo }

func (m *TLPeerSettings) SetAutoarchived(v bool) { m.Data2.Autoarchived = v }
func (m *TLPeerSettings) GetAutoarchived() bool  { return m.Data2.Autoarchived }

func (m *TLPeerSettings) SetGeoDistance(v int32) { m.Data2.GeoDistance = v }
func (m *TLPeerSettings) GetGeoDistance() int32  { return m.Data2.GeoDistance }

func (m *TLPhoneCallEmpty) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCallEmpty) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCallRequested) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCallRequested) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCallRequested) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLPhoneCallRequested) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLPhoneCallRequested) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPhoneCallRequested) GetDate() int32  { return m.Data2.Date }

func (m *TLPhoneCallRequested) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLPhoneCallRequested) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLPhoneCallRequested) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLPhoneCallRequested) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLPhoneCallRequested) SetGAHash(v []byte) { m.Data2.GAHash = v }
func (m *TLPhoneCallRequested) GetGAHash() []byte  { return m.Data2.GAHash }

func (m *TLPhoneCallRequested) SetProtocol(v *PhoneCallProtocol) { m.Data2.Protocol = v }
func (m *TLPhoneCallRequested) GetProtocol() *PhoneCallProtocol  { return m.Data2.Protocol }

func (m *TLPhoneCallRequested) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPhoneCallRequested) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPhoneCallRequested) SetVideo(v bool) { m.Data2.Video = v }
func (m *TLPhoneCallRequested) GetVideo() bool  { return m.Data2.Video }

func (m *TLPhoneCallAccepted) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCallAccepted) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCallAccepted) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLPhoneCallAccepted) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLPhoneCallAccepted) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPhoneCallAccepted) GetDate() int32  { return m.Data2.Date }

func (m *TLPhoneCallAccepted) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLPhoneCallAccepted) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLPhoneCallAccepted) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLPhoneCallAccepted) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLPhoneCallAccepted) SetGB(v []byte) { m.Data2.GB = v }
func (m *TLPhoneCallAccepted) GetGB() []byte  { return m.Data2.GB }

func (m *TLPhoneCallAccepted) SetProtocol(v *PhoneCallProtocol) { m.Data2.Protocol = v }
func (m *TLPhoneCallAccepted) GetProtocol() *PhoneCallProtocol  { return m.Data2.Protocol }

func (m *TLPhoneCallAccepted) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPhoneCallAccepted) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPhoneCallAccepted) SetVideo(v bool) { m.Data2.Video = v }
func (m *TLPhoneCallAccepted) GetVideo() bool  { return m.Data2.Video }

func (m *TLPhoneCall) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCall) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCall) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLPhoneCall) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLPhoneCall) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPhoneCall) GetDate() int32  { return m.Data2.Date }

func (m *TLPhoneCall) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLPhoneCall) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLPhoneCall) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLPhoneCall) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLPhoneCall) SetGAOrB(v []byte) { m.Data2.GAOrB = v }
func (m *TLPhoneCall) GetGAOrB() []byte  { return m.Data2.GAOrB }

func (m *TLPhoneCall) SetKeyFingerprint(v int64) { m.Data2.KeyFingerprint = v }
func (m *TLPhoneCall) GetKeyFingerprint() int64  { return m.Data2.KeyFingerprint }

func (m *TLPhoneCall) SetProtocol(v *PhoneCallProtocol) { m.Data2.Protocol = v }
func (m *TLPhoneCall) GetProtocol() *PhoneCallProtocol  { return m.Data2.Protocol }

func (m *TLPhoneCall) SetConnection(v *PhoneConnection) { m.Data2.Connection = v }
func (m *TLPhoneCall) GetConnection() *PhoneConnection  { return m.Data2.Connection }

func (m *TLPhoneCall) SetAlternativeConnections(v []*PhoneConnection) {
	m.Data2.AlternativeConnections = v
}
func (m *TLPhoneCall) GetAlternativeConnections() []*PhoneConnection {
	return m.Data2.AlternativeConnections
}

func (m *TLPhoneCall) SetStartDate(v int32) { m.Data2.StartDate = v }
func (m *TLPhoneCall) GetStartDate() int32  { return m.Data2.StartDate }

func (m *TLPhoneCall) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPhoneCall) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPhoneCall) SetP2PAllowed(v bool) { m.Data2.P2PAllowed = v }
func (m *TLPhoneCall) GetP2PAllowed() bool  { return m.Data2.P2PAllowed }

func (m *TLPhoneCall) SetVideo(v bool) { m.Data2.Video = v }
func (m *TLPhoneCall) GetVideo() bool  { return m.Data2.Video }

func (m *TLPhoneCall) SetConnections(v []*PhoneConnection) { m.Data2.Connections = v }
func (m *TLPhoneCall) GetConnections() []*PhoneConnection  { return m.Data2.Connections }

func (m *TLPhoneCallWaiting) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPhoneCallWaiting) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPhoneCallWaiting) SetVideo(v bool) { m.Data2.Video = v }
func (m *TLPhoneCallWaiting) GetVideo() bool  { return m.Data2.Video }

func (m *TLPhoneCallWaiting) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCallWaiting) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCallWaiting) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLPhoneCallWaiting) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLPhoneCallWaiting) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPhoneCallWaiting) GetDate() int32  { return m.Data2.Date }

func (m *TLPhoneCallWaiting) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLPhoneCallWaiting) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLPhoneCallWaiting) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLPhoneCallWaiting) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLPhoneCallWaiting) SetProtocol(v *PhoneCallProtocol) { m.Data2.Protocol = v }
func (m *TLPhoneCallWaiting) GetProtocol() *PhoneCallProtocol  { return m.Data2.Protocol }

func (m *TLPhoneCallWaiting) SetReceiveDate(v int32) { m.Data2.ReceiveDate = v }
func (m *TLPhoneCallWaiting) GetReceiveDate() int32  { return m.Data2.ReceiveDate }

func (m *TLPhoneCallDiscarded) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPhoneCallDiscarded) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPhoneCallDiscarded) SetNeedRating(v bool) { m.Data2.NeedRating = v }
func (m *TLPhoneCallDiscarded) GetNeedRating() bool  { return m.Data2.NeedRating }

func (m *TLPhoneCallDiscarded) SetNeedDebug(v bool) { m.Data2.NeedDebug = v }
func (m *TLPhoneCallDiscarded) GetNeedDebug() bool  { return m.Data2.NeedDebug }

func (m *TLPhoneCallDiscarded) SetVideo(v bool) { m.Data2.Video = v }
func (m *TLPhoneCallDiscarded) GetVideo() bool  { return m.Data2.Video }

func (m *TLPhoneCallDiscarded) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCallDiscarded) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCallDiscarded) SetReason(v *PhoneCallDiscardReason) { m.Data2.Reason = v }
func (m *TLPhoneCallDiscarded) GetReason() *PhoneCallDiscardReason  { return m.Data2.Reason }

func (m *TLPhoneCallDiscarded) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLPhoneCallDiscarded) GetDuration() int32  { return m.Data2.Duration }

func (m *TLPhoneCallProtocol) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPhoneCallProtocol) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPhoneCallProtocol) SetUdpP2P(v bool) { m.Data2.UdpP2P = v }
func (m *TLPhoneCallProtocol) GetUdpP2P() bool  { return m.Data2.UdpP2P }

func (m *TLPhoneCallProtocol) SetUdpReflector(v bool) { m.Data2.UdpReflector = v }
func (m *TLPhoneCallProtocol) GetUdpReflector() bool  { return m.Data2.UdpReflector }

func (m *TLPhoneCallProtocol) SetMinLayer(v int32) { m.Data2.MinLayer = v }
func (m *TLPhoneCallProtocol) GetMinLayer() int32  { return m.Data2.MinLayer }

func (m *TLPhoneCallProtocol) SetMaxLayer(v int32) { m.Data2.MaxLayer = v }
func (m *TLPhoneCallProtocol) GetMaxLayer() int32  { return m.Data2.MaxLayer }

func (m *TLPhoneCallProtocol) SetLibraryVersions(v []string) { m.Data2.LibraryVersions = v }
func (m *TLPhoneCallProtocol) GetLibraryVersions() []string  { return m.Data2.LibraryVersions }

func (m *TLPhoneConnection) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneConnection) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneConnection) SetIp(v string) { m.Data2.Ip = v }
func (m *TLPhoneConnection) GetIp() string  { return m.Data2.Ip }

func (m *TLPhoneConnection) SetIpv6(v string) { m.Data2.Ipv6 = v }
func (m *TLPhoneConnection) GetIpv6() string  { return m.Data2.Ipv6 }

func (m *TLPhoneConnection) SetPort(v int32) { m.Data2.Port = v }
func (m *TLPhoneConnection) GetPort() int32  { return m.Data2.Port }

func (m *TLPhoneConnection) SetPeerTag(v []byte) { m.Data2.PeerTag = v }
func (m *TLPhoneConnection) GetPeerTag() []byte  { return m.Data2.PeerTag }

func (m *TLPhoneConnectionWebrtc) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPhoneConnectionWebrtc) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPhoneConnectionWebrtc) SetTurn(v bool) { m.Data2.Turn = v }
func (m *TLPhoneConnectionWebrtc) GetTurn() bool  { return m.Data2.Turn }

func (m *TLPhoneConnectionWebrtc) SetStun(v bool) { m.Data2.Stun = v }
func (m *TLPhoneConnectionWebrtc) GetStun() bool  { return m.Data2.Stun }

func (m *TLPhoneConnectionWebrtc) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneConnectionWebrtc) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneConnectionWebrtc) SetIp(v string) { m.Data2.Ip = v }
func (m *TLPhoneConnectionWebrtc) GetIp() string  { return m.Data2.Ip }

func (m *TLPhoneConnectionWebrtc) SetIpv6(v string) { m.Data2.Ipv6 = v }
func (m *TLPhoneConnectionWebrtc) GetIpv6() string  { return m.Data2.Ipv6 }

func (m *TLPhoneConnectionWebrtc) SetPort(v int32) { m.Data2.Port = v }
func (m *TLPhoneConnectionWebrtc) GetPort() int32  { return m.Data2.Port }

func (m *TLPhoneConnectionWebrtc) SetUsername(v string) { m.Data2.Username = v }
func (m *TLPhoneConnectionWebrtc) GetUsername() string  { return m.Data2.Username }

func (m *TLPhoneConnectionWebrtc) SetPassword(v string) { m.Data2.Password = v }
func (m *TLPhoneConnectionWebrtc) GetPassword() string  { return m.Data2.Password }

func (m *TLPhoneGroupCall) SetCall(v *GroupCall) { m.Data2.Call = v }
func (m *TLPhoneGroupCall) GetCall() *GroupCall  { return m.Data2.Call }

func (m *TLPhoneGroupCall) SetParticipants(v []*GroupCallParticipant) { m.Data2.Participants = v }
func (m *TLPhoneGroupCall) GetParticipants() []*GroupCallParticipant  { return m.Data2.Participants }

func (m *TLPhoneGroupCall) SetParticipantsNextOffset(v string) { m.Data2.ParticipantsNextOffset = v }
func (m *TLPhoneGroupCall) GetParticipantsNextOffset() string  { return m.Data2.ParticipantsNextOffset }

func (m *TLPhoneGroupCall) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPhoneGroupCall) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPhoneGroupParticipants) SetCount(v int32) { m.Data2.Count = v }
func (m *TLPhoneGroupParticipants) GetCount() int32  { return m.Data2.Count }

func (m *TLPhoneGroupParticipants) SetParticipants(v []*GroupCallParticipant) {
	m.Data2.Participants = v
}
func (m *TLPhoneGroupParticipants) GetParticipants() []*GroupCallParticipant {
	return m.Data2.Participants
}

func (m *TLPhoneGroupParticipants) SetNextOffset(v string) { m.Data2.NextOffset = v }
func (m *TLPhoneGroupParticipants) GetNextOffset() string  { return m.Data2.NextOffset }

func (m *TLPhoneGroupParticipants) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPhoneGroupParticipants) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPhoneGroupParticipants) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLPhoneGroupParticipants) GetVersion() int32  { return m.Data2.Version }

func (m *TLPhonePhoneCall) SetPhoneCall(v *PhoneCall) { m.Data2.PhoneCall = v }
func (m *TLPhonePhoneCall) GetPhoneCall() *PhoneCall  { return m.Data2.PhoneCall }

func (m *TLPhonePhoneCall) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPhonePhoneCall) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPhotoEmpty) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhotoEmpty) GetId() int64  { return m.Data2.Id }

func (m *TLPhoto) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPhoto) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPhoto) SetHasStickers(v bool) { m.Data2.HasStickers = v }
func (m *TLPhoto) GetHasStickers() bool  { return m.Data2.HasStickers }

func (m *TLPhoto) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoto) GetId() int64  { return m.Data2.Id }

func (m *TLPhoto) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLPhoto) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLPhoto) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPhoto) GetDate() int32  { return m.Data2.Date }

func (m *TLPhoto) SetSizes(v []*PhotoSize) { m.Data2.Sizes = v }
func (m *TLPhoto) GetSizes() []*PhotoSize  { return m.Data2.Sizes }

func (m *TLPhoto) SetFileReference(v []byte) { m.Data2.FileReference = v }
func (m *TLPhoto) GetFileReference() []byte  { return m.Data2.FileReference }

func (m *TLPhoto) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLPhoto) GetDcId() int32  { return m.Data2.DcId }

func (m *TLPhoto) SetVideoSizes(v []*VideoSize) { m.Data2.VideoSizes = v }
func (m *TLPhoto) GetVideoSizes() []*VideoSize  { return m.Data2.VideoSizes }

func (m *TLPhotoSizeEmpty) SetType(v string) { m.Data2.Type = v }
func (m *TLPhotoSizeEmpty) GetType() string  { return m.Data2.Type }

func (m *TLPhotoSize) SetType(v string) { m.Data2.Type = v }
func (m *TLPhotoSize) GetType() string  { return m.Data2.Type }

func (m *TLPhotoSize) SetLocation(v *FileLocation) { m.Data2.Location = v }
func (m *TLPhotoSize) GetLocation() *FileLocation  { return m.Data2.Location }

func (m *TLPhotoSize) SetW(v int32) { m.Data2.W = v }
func (m *TLPhotoSize) GetW() int32  { return m.Data2.W }

func (m *TLPhotoSize) SetH(v int32) { m.Data2.H = v }
func (m *TLPhotoSize) GetH() int32  { return m.Data2.H }

func (m *TLPhotoSize) SetSize_(v int32) { m.Data2.Size_ = v }
func (m *TLPhotoSize) GetSize_() int32  { return m.Data2.Size_ }

func (m *TLPhotoCachedSize) SetType(v string) { m.Data2.Type = v }
func (m *TLPhotoCachedSize) GetType() string  { return m.Data2.Type }

func (m *TLPhotoCachedSize) SetLocation(v *FileLocation) { m.Data2.Location = v }
func (m *TLPhotoCachedSize) GetLocation() *FileLocation  { return m.Data2.Location }

func (m *TLPhotoCachedSize) SetW(v int32) { m.Data2.W = v }
func (m *TLPhotoCachedSize) GetW() int32  { return m.Data2.W }

func (m *TLPhotoCachedSize) SetH(v int32) { m.Data2.H = v }
func (m *TLPhotoCachedSize) GetH() int32  { return m.Data2.H }

func (m *TLPhotoCachedSize) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLPhotoCachedSize) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLPhotoStrippedSize) SetType(v string) { m.Data2.Type = v }
func (m *TLPhotoStrippedSize) GetType() string  { return m.Data2.Type }

func (m *TLPhotoStrippedSize) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLPhotoStrippedSize) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLPhotoSizeProgressive) SetType(v string) { m.Data2.Type = v }
func (m *TLPhotoSizeProgressive) GetType() string  { return m.Data2.Type }

func (m *TLPhotoSizeProgressive) SetLocation(v *FileLocation) { m.Data2.Location = v }
func (m *TLPhotoSizeProgressive) GetLocation() *FileLocation  { return m.Data2.Location }

func (m *TLPhotoSizeProgressive) SetW(v int32) { m.Data2.W = v }
func (m *TLPhotoSizeProgressive) GetW() int32  { return m.Data2.W }

func (m *TLPhotoSizeProgressive) SetH(v int32) { m.Data2.H = v }
func (m *TLPhotoSizeProgressive) GetH() int32  { return m.Data2.H }

func (m *TLPhotoSizeProgressive) SetSizes(v []int32) { m.Data2.Sizes = v }
func (m *TLPhotoSizeProgressive) GetSizes() []int32  { return m.Data2.Sizes }

func (m *TLPhotoPathSize) SetType(v string) { m.Data2.Type = v }
func (m *TLPhotoPathSize) GetType() string  { return m.Data2.Type }

func (m *TLPhotoPathSize) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLPhotoPathSize) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLPhotosPhoto) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLPhotosPhoto) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLPhotosPhoto) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPhotosPhoto) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPhotosPhotos) SetPhotos(v []*Photo) { m.Data2.Photos = v }
func (m *TLPhotosPhotos) GetPhotos() []*Photo  { return m.Data2.Photos }

func (m *TLPhotosPhotos) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPhotosPhotos) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPhotosPhotosSlice) SetCount(v int32) { m.Data2.Count = v }
func (m *TLPhotosPhotosSlice) GetCount() int32  { return m.Data2.Count }

func (m *TLPhotosPhotosSlice) SetPhotos(v []*Photo) { m.Data2.Photos = v }
func (m *TLPhotosPhotosSlice) GetPhotos() []*Photo  { return m.Data2.Photos }

func (m *TLPhotosPhotosSlice) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPhotosPhotosSlice) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPoll) SetId(v int64) { m.Data2.Id = v }
func (m *TLPoll) GetId() int64  { return m.Data2.Id }

func (m *TLPoll) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPoll) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPoll) SetClosed(v bool) { m.Data2.Closed = v }
func (m *TLPoll) GetClosed() bool  { return m.Data2.Closed }

func (m *TLPoll) SetQuestion(v string) { m.Data2.Question = v }
func (m *TLPoll) GetQuestion() string  { return m.Data2.Question }

func (m *TLPoll) SetAnswers(v []*PollAnswer) { m.Data2.Answers = v }
func (m *TLPoll) GetAnswers() []*PollAnswer  { return m.Data2.Answers }

func (m *TLPoll) SetPublicVoters(v bool) { m.Data2.PublicVoters = v }
func (m *TLPoll) GetPublicVoters() bool  { return m.Data2.PublicVoters }

func (m *TLPoll) SetMultipleChoice(v bool) { m.Data2.MultipleChoice = v }
func (m *TLPoll) GetMultipleChoice() bool  { return m.Data2.MultipleChoice }

func (m *TLPoll) SetQuiz(v bool) { m.Data2.Quiz = v }
func (m *TLPoll) GetQuiz() bool  { return m.Data2.Quiz }

func (m *TLPoll) SetClosePeriod(v int32) { m.Data2.ClosePeriod = v }
func (m *TLPoll) GetClosePeriod() int32  { return m.Data2.ClosePeriod }

func (m *TLPoll) SetCloseDate(v int32) { m.Data2.CloseDate = v }
func (m *TLPoll) GetCloseDate() int32  { return m.Data2.CloseDate }

func (m *TLPollAnswer) SetText(v string) { m.Data2.Text = v }
func (m *TLPollAnswer) GetText() string  { return m.Data2.Text }

func (m *TLPollAnswer) SetOption(v []byte) { m.Data2.Option = v }
func (m *TLPollAnswer) GetOption() []byte  { return m.Data2.Option }

func (m *TLPollAnswerVoters) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPollAnswerVoters) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPollAnswerVoters) SetChosen(v bool) { m.Data2.Chosen = v }
func (m *TLPollAnswerVoters) GetChosen() bool  { return m.Data2.Chosen }

func (m *TLPollAnswerVoters) SetOption(v []byte) { m.Data2.Option = v }
func (m *TLPollAnswerVoters) GetOption() []byte  { return m.Data2.Option }

func (m *TLPollAnswerVoters) SetVoters(v int32) { m.Data2.Voters = v }
func (m *TLPollAnswerVoters) GetVoters() int32  { return m.Data2.Voters }

func (m *TLPollResults) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLPollResults) GetFlags() int32  { return m.Data2.Flags }

func (m *TLPollResults) SetMin(v bool) { m.Data2.Min = v }
func (m *TLPollResults) GetMin() bool  { return m.Data2.Min }

func (m *TLPollResults) SetResults(v []*PollAnswerVoters) { m.Data2.Results = v }
func (m *TLPollResults) GetResults() []*PollAnswerVoters  { return m.Data2.Results }

func (m *TLPollResults) SetTotalVoters(v int32) { m.Data2.TotalVoters = v }
func (m *TLPollResults) GetTotalVoters() int32  { return m.Data2.TotalVoters }

func (m *TLPollResults) SetRecentVoters(v []int32) { m.Data2.RecentVoters = v }
func (m *TLPollResults) GetRecentVoters() []int32  { return m.Data2.RecentVoters }

func (m *TLPollResults) SetSolution(v string) { m.Data2.Solution = v }
func (m *TLPollResults) GetSolution() string  { return m.Data2.Solution }

func (m *TLPollResults) SetSolutionEntities(v []*MessageEntity) { m.Data2.SolutionEntities = v }
func (m *TLPollResults) GetSolutionEntities() []*MessageEntity  { return m.Data2.SolutionEntities }

func (m *TLPong) SetMsgId(v int64) { m.Data2.MsgId = v }
func (m *TLPong) GetMsgId() int64  { return m.Data2.MsgId }

func (m *TLPong) SetPingId(v int64) { m.Data2.PingId = v }
func (m *TLPong) GetPingId() int64  { return m.Data2.PingId }

func (m *TLPopularContact) SetClientId(v int64) { m.Data2.ClientId = v }
func (m *TLPopularContact) GetClientId() int64  { return m.Data2.ClientId }

func (m *TLPopularContact) SetImporters(v int32) { m.Data2.Importers = v }
func (m *TLPopularContact) GetImporters() int32  { return m.Data2.Importers }

func (m *TLPostAddress) SetStreetLine1(v string) { m.Data2.StreetLine1 = v }
func (m *TLPostAddress) GetStreetLine1() string  { return m.Data2.StreetLine1 }

func (m *TLPostAddress) SetStreetLine2(v string) { m.Data2.StreetLine2 = v }
func (m *TLPostAddress) GetStreetLine2() string  { return m.Data2.StreetLine2 }

func (m *TLPostAddress) SetCity(v string) { m.Data2.City = v }
func (m *TLPostAddress) GetCity() string  { return m.Data2.City }

func (m *TLPostAddress) SetState(v string) { m.Data2.State = v }
func (m *TLPostAddress) GetState() string  { return m.Data2.State }

func (m *TLPostAddress) SetCountryIso2(v string) { m.Data2.CountryIso2 = v }
func (m *TLPostAddress) GetCountryIso2() string  { return m.Data2.CountryIso2 }

func (m *TLPostAddress) SetPostCode(v string) { m.Data2.PostCode = v }
func (m *TLPostAddress) GetPostCode() string  { return m.Data2.PostCode }

func (m *TLPrivacyValueAllowUsers) SetUsers(v []int32) { m.Data2.Users = v }
func (m *TLPrivacyValueAllowUsers) GetUsers() []int32  { return m.Data2.Users }

func (m *TLPrivacyValueDisallowUsers) SetUsers(v []int32) { m.Data2.Users = v }
func (m *TLPrivacyValueDisallowUsers) GetUsers() []int32  { return m.Data2.Users }

func (m *TLPrivacyValueAllowChatParticipants) SetChats(v []int32) { m.Data2.Chats = v }
func (m *TLPrivacyValueAllowChatParticipants) GetChats() []int32  { return m.Data2.Chats }

func (m *TLPrivacyValueDisallowChatParticipants) SetChats(v []int32) { m.Data2.Chats = v }
func (m *TLPrivacyValueDisallowChatParticipants) GetChats() []int32  { return m.Data2.Chats }

func (m *TLReceivedNotifyMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLReceivedNotifyMessage) GetId() int32  { return m.Data2.Id }

func (m *TLReceivedNotifyMessage) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLReceivedNotifyMessage) GetFlags() int32  { return m.Data2.Flags }

func (m *TLRecentMeUrlUnknown) SetUrl(v string) { m.Data2.Url = v }
func (m *TLRecentMeUrlUnknown) GetUrl() string  { return m.Data2.Url }

func (m *TLRecentMeUrlUser) SetUrl(v string) { m.Data2.Url = v }
func (m *TLRecentMeUrlUser) GetUrl() string  { return m.Data2.Url }

func (m *TLRecentMeUrlUser) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLRecentMeUrlUser) GetUserId() int32  { return m.Data2.UserId }

func (m *TLRecentMeUrlChat) SetUrl(v string) { m.Data2.Url = v }
func (m *TLRecentMeUrlChat) GetUrl() string  { return m.Data2.Url }

func (m *TLRecentMeUrlChat) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLRecentMeUrlChat) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLRecentMeUrlChatInvite) SetUrl(v string) { m.Data2.Url = v }
func (m *TLRecentMeUrlChatInvite) GetUrl() string  { return m.Data2.Url }

func (m *TLRecentMeUrlChatInvite) SetChatInvite(v *ChatInvite) { m.Data2.ChatInvite = v }
func (m *TLRecentMeUrlChatInvite) GetChatInvite() *ChatInvite  { return m.Data2.ChatInvite }

func (m *TLRecentMeUrlStickerSet) SetUrl(v string) { m.Data2.Url = v }
func (m *TLRecentMeUrlStickerSet) GetUrl() string  { return m.Data2.Url }

func (m *TLRecentMeUrlStickerSet) SetSet(v *StickerSetCovered) { m.Data2.Set = v }
func (m *TLRecentMeUrlStickerSet) GetSet() *StickerSetCovered  { return m.Data2.Set }

func (m *TLReplyKeyboardHide) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLReplyKeyboardHide) GetFlags() int32  { return m.Data2.Flags }

func (m *TLReplyKeyboardHide) SetSelective(v bool) { m.Data2.Selective = v }
func (m *TLReplyKeyboardHide) GetSelective() bool  { return m.Data2.Selective }

func (m *TLReplyKeyboardForceReply) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLReplyKeyboardForceReply) GetFlags() int32  { return m.Data2.Flags }

func (m *TLReplyKeyboardForceReply) SetSingleUse(v bool) { m.Data2.SingleUse = v }
func (m *TLReplyKeyboardForceReply) GetSingleUse() bool  { return m.Data2.SingleUse }

func (m *TLReplyKeyboardForceReply) SetSelective(v bool) { m.Data2.Selective = v }
func (m *TLReplyKeyboardForceReply) GetSelective() bool  { return m.Data2.Selective }

func (m *TLReplyKeyboardMarkup) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLReplyKeyboardMarkup) GetFlags() int32  { return m.Data2.Flags }

func (m *TLReplyKeyboardMarkup) SetResize(v bool) { m.Data2.Resize = v }
func (m *TLReplyKeyboardMarkup) GetResize() bool  { return m.Data2.Resize }

func (m *TLReplyKeyboardMarkup) SetSingleUse(v bool) { m.Data2.SingleUse = v }
func (m *TLReplyKeyboardMarkup) GetSingleUse() bool  { return m.Data2.SingleUse }

func (m *TLReplyKeyboardMarkup) SetSelective(v bool) { m.Data2.Selective = v }
func (m *TLReplyKeyboardMarkup) GetSelective() bool  { return m.Data2.Selective }

func (m *TLReplyKeyboardMarkup) SetRows(v []*KeyboardButtonRow) { m.Data2.Rows = v }
func (m *TLReplyKeyboardMarkup) GetRows() []*KeyboardButtonRow  { return m.Data2.Rows }

func (m *TLReplyInlineMarkup) SetRows(v []*KeyboardButtonRow) { m.Data2.Rows = v }
func (m *TLReplyInlineMarkup) GetRows() []*KeyboardButtonRow  { return m.Data2.Rows }

func (m *TLInputReportReasonOther) SetText(v string) { m.Data2.Text = v }
func (m *TLInputReportReasonOther) GetText() string  { return m.Data2.Text }

func (m *TLResPQ) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLResPQ) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLResPQ) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLResPQ) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLResPQ) SetPq(v string) { m.Data2.Pq = v }
func (m *TLResPQ) GetPq() string  { return m.Data2.Pq }

func (m *TLResPQ) SetServerPublicKeyFingerprints(v []int64) { m.Data2.ServerPublicKeyFingerprints = v }
func (m *TLResPQ) GetServerPublicKeyFingerprints() []int64 {
	return m.Data2.ServerPublicKeyFingerprints
}

func (m *TLRestrictionReason) SetPlatform(v string) { m.Data2.Platform = v }
func (m *TLRestrictionReason) GetPlatform() string  { return m.Data2.Platform }

func (m *TLRestrictionReason) SetReason(v string) { m.Data2.Reason = v }
func (m *TLRestrictionReason) GetReason() string  { return m.Data2.Reason }

func (m *TLRestrictionReason) SetText(v string) { m.Data2.Text = v }
func (m *TLRestrictionReason) GetText() string  { return m.Data2.Text }

func (m *TLTextPlain) SetText744694E071(v string) { m.Data2.Text744694E071 = v }
func (m *TLTextPlain) GetText744694E071() string  { return m.Data2.Text744694E071 }

func (m *TLTextBold) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextBold) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextItalic) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextItalic) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextUnderline) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextUnderline) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextStrike) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextStrike) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextFixed) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextFixed) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextUrl) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextUrl) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextUrl) SetUrl(v string) { m.Data2.Url = v }
func (m *TLTextUrl) GetUrl() string  { return m.Data2.Url }

func (m *TLTextUrl) SetWebpageId(v int64) { m.Data2.WebpageId = v }
func (m *TLTextUrl) GetWebpageId() int64  { return m.Data2.WebpageId }

func (m *TLTextEmail) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextEmail) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextEmail) SetEmail(v string) { m.Data2.Email = v }
func (m *TLTextEmail) GetEmail() string  { return m.Data2.Email }

func (m *TLTextConcat) SetTexts(v []*RichText) { m.Data2.Texts = v }
func (m *TLTextConcat) GetTexts() []*RichText  { return m.Data2.Texts }

func (m *TLTextSubscript) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextSubscript) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextSuperscript) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextSuperscript) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextMarked) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextMarked) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextPhone) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextPhone) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextPhone) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLTextPhone) GetPhone() string  { return m.Data2.Phone }

func (m *TLTextImage) SetDocumentId(v int64) { m.Data2.DocumentId = v }
func (m *TLTextImage) GetDocumentId() int64  { return m.Data2.DocumentId }

func (m *TLTextImage) SetW(v int32) { m.Data2.W = v }
func (m *TLTextImage) GetW() int32  { return m.Data2.W }

func (m *TLTextImage) SetH(v int32) { m.Data2.H = v }
func (m *TLTextImage) GetH() int32  { return m.Data2.H }

func (m *TLTextAnchor) SetText6724ABC471(v *RichText) { m.Data2.Text6724ABC471 = v }
func (m *TLTextAnchor) GetText6724ABC471() *RichText  { return m.Data2.Text6724ABC471 }

func (m *TLTextAnchor) SetName(v string) { m.Data2.Name = v }
func (m *TLTextAnchor) GetName() string  { return m.Data2.Name }

func (m *TLRpcAnswerDropped) SetMsgId(v int64) { m.Data2.MsgId = v }
func (m *TLRpcAnswerDropped) GetMsgId() int64  { return m.Data2.MsgId }

func (m *TLRpcAnswerDropped) SetSeqNo(v int32) { m.Data2.SeqNo = v }
func (m *TLRpcAnswerDropped) GetSeqNo() int32  { return m.Data2.SeqNo }

func (m *TLRpcAnswerDropped) SetBytes(v int32) { m.Data2.Bytes = v }
func (m *TLRpcAnswerDropped) GetBytes() int32  { return m.Data2.Bytes }

func (m *TLRpcError) SetErrorCode(v int32) { m.Data2.ErrorCode = v }
func (m *TLRpcError) GetErrorCode() int32  { return m.Data2.ErrorCode }

func (m *TLRpcError) SetErrorMessage(v string) { m.Data2.ErrorMessage = v }
func (m *TLRpcError) GetErrorMessage() string  { return m.Data2.ErrorMessage }

func (m *TLSavedPhoneContact) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLSavedPhoneContact) GetPhone() string  { return m.Data2.Phone }

func (m *TLSavedPhoneContact) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLSavedPhoneContact) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLSavedPhoneContact) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLSavedPhoneContact) GetLastName() string  { return m.Data2.LastName }

func (m *TLSavedPhoneContact) SetDate(v int32) { m.Data2.Date = v }
func (m *TLSavedPhoneContact) GetDate() int32  { return m.Data2.Date }

func (m *TLSecureCredentialsEncrypted) SetData(v []byte) { m.Data2.Data = v }
func (m *TLSecureCredentialsEncrypted) GetData() []byte  { return m.Data2.Data }

func (m *TLSecureCredentialsEncrypted) SetHash(v []byte) { m.Data2.Hash = v }
func (m *TLSecureCredentialsEncrypted) GetHash() []byte  { return m.Data2.Hash }

func (m *TLSecureCredentialsEncrypted) SetSecret(v []byte) { m.Data2.Secret = v }
func (m *TLSecureCredentialsEncrypted) GetSecret() []byte  { return m.Data2.Secret }

func (m *TLSecureData) SetData(v []byte) { m.Data2.Data = v }
func (m *TLSecureData) GetData() []byte  { return m.Data2.Data }

func (m *TLSecureData) SetDataHash(v []byte) { m.Data2.DataHash = v }
func (m *TLSecureData) GetDataHash() []byte  { return m.Data2.DataHash }

func (m *TLSecureData) SetSecret(v []byte) { m.Data2.Secret = v }
func (m *TLSecureData) GetSecret() []byte  { return m.Data2.Secret }

func (m *TLSecureFile) SetId(v int64) { m.Data2.Id = v }
func (m *TLSecureFile) GetId() int64  { return m.Data2.Id }

func (m *TLSecureFile) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLSecureFile) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLSecureFile) SetSize_(v int32) { m.Data2.Size_ = v }
func (m *TLSecureFile) GetSize_() int32  { return m.Data2.Size_ }

func (m *TLSecureFile) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLSecureFile) GetDcId() int32  { return m.Data2.DcId }

func (m *TLSecureFile) SetDate(v int32) { m.Data2.Date = v }
func (m *TLSecureFile) GetDate() int32  { return m.Data2.Date }

func (m *TLSecureFile) SetFileHash(v []byte) { m.Data2.FileHash = v }
func (m *TLSecureFile) GetFileHash() []byte  { return m.Data2.FileHash }

func (m *TLSecureFile) SetSecret(v []byte) { m.Data2.Secret = v }
func (m *TLSecureFile) GetSecret() []byte  { return m.Data2.Secret }

func (m *TLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000) SetSalt(v []byte) { m.Data2.Salt = v }
func (m *TLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000) GetSalt() []byte  { return m.Data2.Salt }

func (m *TLSecurePasswordKdfAlgoSHA512) SetSalt(v []byte) { m.Data2.Salt = v }
func (m *TLSecurePasswordKdfAlgoSHA512) GetSalt() []byte  { return m.Data2.Salt }

func (m *TLSecurePlainPhone) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLSecurePlainPhone) GetPhone() string  { return m.Data2.Phone }

func (m *TLSecurePlainEmail) SetEmail(v string) { m.Data2.Email = v }
func (m *TLSecurePlainEmail) GetEmail() string  { return m.Data2.Email }

func (m *TLSecureRequiredType) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLSecureRequiredType) GetFlags() int32  { return m.Data2.Flags }

func (m *TLSecureRequiredType) SetNativeNames(v bool) { m.Data2.NativeNames = v }
func (m *TLSecureRequiredType) GetNativeNames() bool  { return m.Data2.NativeNames }

func (m *TLSecureRequiredType) SetSelfieRequired(v bool) { m.Data2.SelfieRequired = v }
func (m *TLSecureRequiredType) GetSelfieRequired() bool  { return m.Data2.SelfieRequired }

func (m *TLSecureRequiredType) SetTranslationRequired(v bool) { m.Data2.TranslationRequired = v }
func (m *TLSecureRequiredType) GetTranslationRequired() bool  { return m.Data2.TranslationRequired }

func (m *TLSecureRequiredType) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureRequiredType) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureRequiredTypeOneOf) SetTypes(v []*SecureRequiredType) { m.Data2.Types = v }
func (m *TLSecureRequiredTypeOneOf) GetTypes() []*SecureRequiredType  { return m.Data2.Types }

func (m *TLSecureSecretSettings) SetSecureAlgo(v *SecurePasswordKdfAlgo) { m.Data2.SecureAlgo = v }
func (m *TLSecureSecretSettings) GetSecureAlgo() *SecurePasswordKdfAlgo  { return m.Data2.SecureAlgo }

func (m *TLSecureSecretSettings) SetSecureSecret(v []byte) { m.Data2.SecureSecret = v }
func (m *TLSecureSecretSettings) GetSecureSecret() []byte  { return m.Data2.SecureSecret }

func (m *TLSecureSecretSettings) SetSecureSecretId(v int64) { m.Data2.SecureSecretId = v }
func (m *TLSecureSecretSettings) GetSecureSecretId() int64  { return m.Data2.SecureSecretId }

func (m *TLSecureValue) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLSecureValue) GetFlags() int32  { return m.Data2.Flags }

func (m *TLSecureValue) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureValue) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureValue) SetData(v *SecureData) { m.Data2.Data = v }
func (m *TLSecureValue) GetData() *SecureData  { return m.Data2.Data }

func (m *TLSecureValue) SetFrontSide(v *SecureFile) { m.Data2.FrontSide = v }
func (m *TLSecureValue) GetFrontSide() *SecureFile  { return m.Data2.FrontSide }

func (m *TLSecureValue) SetReverseSide(v *SecureFile) { m.Data2.ReverseSide = v }
func (m *TLSecureValue) GetReverseSide() *SecureFile  { return m.Data2.ReverseSide }

func (m *TLSecureValue) SetSelfie(v *SecureFile) { m.Data2.Selfie = v }
func (m *TLSecureValue) GetSelfie() *SecureFile  { return m.Data2.Selfie }

func (m *TLSecureValue) SetTranslation(v []*SecureFile) { m.Data2.Translation = v }
func (m *TLSecureValue) GetTranslation() []*SecureFile  { return m.Data2.Translation }

func (m *TLSecureValue) SetFiles(v []*SecureFile) { m.Data2.Files = v }
func (m *TLSecureValue) GetFiles() []*SecureFile  { return m.Data2.Files }

func (m *TLSecureValue) SetPlainData(v *SecurePlainData) { m.Data2.PlainData = v }
func (m *TLSecureValue) GetPlainData() *SecurePlainData  { return m.Data2.PlainData }

func (m *TLSecureValue) SetHash(v []byte) { m.Data2.Hash = v }
func (m *TLSecureValue) GetHash() []byte  { return m.Data2.Hash }

func (m *TLSecureValueErrorData) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureValueErrorData) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureValueErrorData) SetDataHash(v []byte) { m.Data2.DataHash = v }
func (m *TLSecureValueErrorData) GetDataHash() []byte  { return m.Data2.DataHash }

func (m *TLSecureValueErrorData) SetField(v string) { m.Data2.Field = v }
func (m *TLSecureValueErrorData) GetField() string  { return m.Data2.Field }

func (m *TLSecureValueErrorData) SetText(v string) { m.Data2.Text = v }
func (m *TLSecureValueErrorData) GetText() string  { return m.Data2.Text }

func (m *TLSecureValueErrorFrontSide) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureValueErrorFrontSide) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureValueErrorFrontSide) SetFileHashBE3DFA85(v []byte) { m.Data2.FileHashBE3DFA85 = v }
func (m *TLSecureValueErrorFrontSide) GetFileHashBE3DFA85() []byte  { return m.Data2.FileHashBE3DFA85 }

func (m *TLSecureValueErrorFrontSide) SetText(v string) { m.Data2.Text = v }
func (m *TLSecureValueErrorFrontSide) GetText() string  { return m.Data2.Text }

func (m *TLSecureValueErrorReverseSide) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureValueErrorReverseSide) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureValueErrorReverseSide) SetFileHashBE3DFA85(v []byte) { m.Data2.FileHashBE3DFA85 = v }
func (m *TLSecureValueErrorReverseSide) GetFileHashBE3DFA85() []byte  { return m.Data2.FileHashBE3DFA85 }

func (m *TLSecureValueErrorReverseSide) SetText(v string) { m.Data2.Text = v }
func (m *TLSecureValueErrorReverseSide) GetText() string  { return m.Data2.Text }

func (m *TLSecureValueErrorSelfie) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureValueErrorSelfie) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureValueErrorSelfie) SetFileHashBE3DFA85(v []byte) { m.Data2.FileHashBE3DFA85 = v }
func (m *TLSecureValueErrorSelfie) GetFileHashBE3DFA85() []byte  { return m.Data2.FileHashBE3DFA85 }

func (m *TLSecureValueErrorSelfie) SetText(v string) { m.Data2.Text = v }
func (m *TLSecureValueErrorSelfie) GetText() string  { return m.Data2.Text }

func (m *TLSecureValueErrorFile) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureValueErrorFile) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureValueErrorFile) SetFileHashBE3DFA85(v []byte) { m.Data2.FileHashBE3DFA85 = v }
func (m *TLSecureValueErrorFile) GetFileHashBE3DFA85() []byte  { return m.Data2.FileHashBE3DFA85 }

func (m *TLSecureValueErrorFile) SetText(v string) { m.Data2.Text = v }
func (m *TLSecureValueErrorFile) GetText() string  { return m.Data2.Text }

func (m *TLSecureValueErrorFiles) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureValueErrorFiles) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureValueErrorFiles) SetFileHash666220E985(v [][]byte) { m.Data2.FileHash666220E985 = v }
func (m *TLSecureValueErrorFiles) GetFileHash666220E985() [][]byte  { return m.Data2.FileHash666220E985 }

func (m *TLSecureValueErrorFiles) SetText(v string) { m.Data2.Text = v }
func (m *TLSecureValueErrorFiles) GetText() string  { return m.Data2.Text }

func (m *TLSecureValueError) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureValueError) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureValueError) SetHash(v []byte) { m.Data2.Hash = v }
func (m *TLSecureValueError) GetHash() []byte  { return m.Data2.Hash }

func (m *TLSecureValueError) SetText(v string) { m.Data2.Text = v }
func (m *TLSecureValueError) GetText() string  { return m.Data2.Text }

func (m *TLSecureValueErrorTranslationFile) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureValueErrorTranslationFile) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureValueErrorTranslationFile) SetFileHashBE3DFA85(v []byte) {
	m.Data2.FileHashBE3DFA85 = v
}
func (m *TLSecureValueErrorTranslationFile) GetFileHashBE3DFA85() []byte {
	return m.Data2.FileHashBE3DFA85
}

func (m *TLSecureValueErrorTranslationFile) SetText(v string) { m.Data2.Text = v }
func (m *TLSecureValueErrorTranslationFile) GetText() string  { return m.Data2.Text }

func (m *TLSecureValueErrorTranslationFiles) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureValueErrorTranslationFiles) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureValueErrorTranslationFiles) SetFileHash666220E985(v [][]byte) {
	m.Data2.FileHash666220E985 = v
}
func (m *TLSecureValueErrorTranslationFiles) GetFileHash666220E985() [][]byte {
	return m.Data2.FileHash666220E985
}

func (m *TLSecureValueErrorTranslationFiles) SetText(v string) { m.Data2.Text = v }
func (m *TLSecureValueErrorTranslationFiles) GetText() string  { return m.Data2.Text }

func (m *TLSecureValueHash) SetType(v *SecureValueType) { m.Data2.Type = v }
func (m *TLSecureValueHash) GetType() *SecureValueType  { return m.Data2.Type }

func (m *TLSecureValueHash) SetHash(v []byte) { m.Data2.Hash = v }
func (m *TLSecureValueHash) GetHash() []byte  { return m.Data2.Hash }

func (m *TLSendMessageUploadVideoAction) SetProgress(v int32) { m.Data2.Progress = v }
func (m *TLSendMessageUploadVideoAction) GetProgress() int32  { return m.Data2.Progress }

func (m *TLSendMessageUploadAudioAction) SetProgress(v int32) { m.Data2.Progress = v }
func (m *TLSendMessageUploadAudioAction) GetProgress() int32  { return m.Data2.Progress }

func (m *TLSendMessageUploadPhotoAction) SetProgress(v int32) { m.Data2.Progress = v }
func (m *TLSendMessageUploadPhotoAction) GetProgress() int32  { return m.Data2.Progress }

func (m *TLSendMessageUploadDocumentAction) SetProgress(v int32) { m.Data2.Progress = v }
func (m *TLSendMessageUploadDocumentAction) GetProgress() int32  { return m.Data2.Progress }

func (m *TLSendMessageUploadRoundAction) SetProgress(v int32) { m.Data2.Progress = v }
func (m *TLSendMessageUploadRoundAction) GetProgress() int32  { return m.Data2.Progress }

func (m *TLServer_DHInnerData) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLServer_DHInnerData) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLServer_DHInnerData) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLServer_DHInnerData) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLServer_DHInnerData) SetG(v int32) { m.Data2.G = v }
func (m *TLServer_DHInnerData) GetG() int32  { return m.Data2.G }

func (m *TLServer_DHInnerData) SetDhPrime(v string) { m.Data2.DhPrime = v }
func (m *TLServer_DHInnerData) GetDhPrime() string  { return m.Data2.DhPrime }

func (m *TLServer_DHInnerData) SetGA(v string) { m.Data2.GA = v }
func (m *TLServer_DHInnerData) GetGA() string  { return m.Data2.GA }

func (m *TLServer_DHInnerData) SetServerTime(v int32) { m.Data2.ServerTime = v }
func (m *TLServer_DHInnerData) GetServerTime() int32  { return m.Data2.ServerTime }

func (m *TLServer_DHParamsFail) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLServer_DHParamsFail) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLServer_DHParamsFail) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLServer_DHParamsFail) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLServer_DHParamsFail) SetNewNonceHash(v []byte) { m.Data2.NewNonceHash = v }
func (m *TLServer_DHParamsFail) GetNewNonceHash() []byte  { return m.Data2.NewNonceHash }

func (m *TLServer_DHParamsOk) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLServer_DHParamsOk) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLServer_DHParamsOk) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLServer_DHParamsOk) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLServer_DHParamsOk) SetEncryptedAnswer(v string) { m.Data2.EncryptedAnswer = v }
func (m *TLServer_DHParamsOk) GetEncryptedAnswer() string  { return m.Data2.EncryptedAnswer }

func (m *TLDhGenOk) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLDhGenOk) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLDhGenOk) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLDhGenOk) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLDhGenOk) SetNewNonceHash1(v []byte) { m.Data2.NewNonceHash1 = v }
func (m *TLDhGenOk) GetNewNonceHash1() []byte  { return m.Data2.NewNonceHash1 }

func (m *TLDhGenRetry) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLDhGenRetry) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLDhGenRetry) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLDhGenRetry) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLDhGenRetry) SetNewNonceHash2(v []byte) { m.Data2.NewNonceHash2 = v }
func (m *TLDhGenRetry) GetNewNonceHash2() []byte  { return m.Data2.NewNonceHash2 }

func (m *TLDhGenFail) SetNonce(v []byte) { m.Data2.Nonce = v }
func (m *TLDhGenFail) GetNonce() []byte  { return m.Data2.Nonce }

func (m *TLDhGenFail) SetServerNonce(v []byte) { m.Data2.ServerNonce = v }
func (m *TLDhGenFail) GetServerNonce() []byte  { return m.Data2.ServerNonce }

func (m *TLDhGenFail) SetNewNonceHash3(v []byte) { m.Data2.NewNonceHash3 = v }
func (m *TLDhGenFail) GetNewNonceHash3() []byte  { return m.Data2.NewNonceHash3 }

func (m *TLShippingOption) SetId(v string) { m.Data2.Id = v }
func (m *TLShippingOption) GetId() string  { return m.Data2.Id }

func (m *TLShippingOption) SetTitle(v string) { m.Data2.Title = v }
func (m *TLShippingOption) GetTitle() string  { return m.Data2.Title }

func (m *TLShippingOption) SetPrices(v []*LabeledPrice) { m.Data2.Prices = v }
func (m *TLShippingOption) GetPrices() []*LabeledPrice  { return m.Data2.Prices }

func (m *TLStatsAbsValueAndPrev) SetCurrent(v float64) { m.Data2.Current = v }
func (m *TLStatsAbsValueAndPrev) GetCurrent() float64  { return m.Data2.Current }

func (m *TLStatsAbsValueAndPrev) SetPrevious(v float64) { m.Data2.Previous = v }
func (m *TLStatsAbsValueAndPrev) GetPrevious() float64  { return m.Data2.Previous }

func (m *TLStatsDateRangeDays) SetMinDate(v int32) { m.Data2.MinDate = v }
func (m *TLStatsDateRangeDays) GetMinDate() int32  { return m.Data2.MinDate }

func (m *TLStatsDateRangeDays) SetMaxDate(v int32) { m.Data2.MaxDate = v }
func (m *TLStatsDateRangeDays) GetMaxDate() int32  { return m.Data2.MaxDate }

func (m *TLStatsGraphAsync) SetToken(v string) { m.Data2.Token = v }
func (m *TLStatsGraphAsync) GetToken() string  { return m.Data2.Token }

func (m *TLStatsGraphError) SetError(v string) { m.Data2.Error = v }
func (m *TLStatsGraphError) GetError() string  { return m.Data2.Error }

func (m *TLStatsGraph) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLStatsGraph) GetFlags() int32  { return m.Data2.Flags }

func (m *TLStatsGraph) SetJson(v *DataJSON) { m.Data2.Json = v }
func (m *TLStatsGraph) GetJson() *DataJSON  { return m.Data2.Json }

func (m *TLStatsGraph) SetZoomToken(v string) { m.Data2.ZoomToken = v }
func (m *TLStatsGraph) GetZoomToken() string  { return m.Data2.ZoomToken }

func (m *TLStatsGroupTopAdmin) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLStatsGroupTopAdmin) GetUserId() int32  { return m.Data2.UserId }

func (m *TLStatsGroupTopAdmin) SetDeleted(v int32) { m.Data2.Deleted = v }
func (m *TLStatsGroupTopAdmin) GetDeleted() int32  { return m.Data2.Deleted }

func (m *TLStatsGroupTopAdmin) SetKicked(v int32) { m.Data2.Kicked = v }
func (m *TLStatsGroupTopAdmin) GetKicked() int32  { return m.Data2.Kicked }

func (m *TLStatsGroupTopAdmin) SetBanned(v int32) { m.Data2.Banned = v }
func (m *TLStatsGroupTopAdmin) GetBanned() int32  { return m.Data2.Banned }

func (m *TLStatsGroupTopInviter) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLStatsGroupTopInviter) GetUserId() int32  { return m.Data2.UserId }

func (m *TLStatsGroupTopInviter) SetInvitations(v int32) { m.Data2.Invitations = v }
func (m *TLStatsGroupTopInviter) GetInvitations() int32  { return m.Data2.Invitations }

func (m *TLStatsGroupTopPoster) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLStatsGroupTopPoster) GetUserId() int32  { return m.Data2.UserId }

func (m *TLStatsGroupTopPoster) SetMessages(v int32) { m.Data2.Messages = v }
func (m *TLStatsGroupTopPoster) GetMessages() int32  { return m.Data2.Messages }

func (m *TLStatsGroupTopPoster) SetAvgChars(v int32) { m.Data2.AvgChars = v }
func (m *TLStatsGroupTopPoster) GetAvgChars() int32  { return m.Data2.AvgChars }

func (m *TLStatsPercentValue) SetPart(v float64) { m.Data2.Part = v }
func (m *TLStatsPercentValue) GetPart() float64  { return m.Data2.Part }

func (m *TLStatsPercentValue) SetTotal(v float64) { m.Data2.Total = v }
func (m *TLStatsPercentValue) GetTotal() float64  { return m.Data2.Total }

func (m *TLStatsURL) SetUrl(v string) { m.Data2.Url = v }
func (m *TLStatsURL) GetUrl() string  { return m.Data2.Url }

func (m *TLStatsBroadcastStats) SetPeriod(v *StatsDateRangeDays) { m.Data2.Period = v }
func (m *TLStatsBroadcastStats) GetPeriod() *StatsDateRangeDays  { return m.Data2.Period }

func (m *TLStatsBroadcastStats) SetFollowers(v *StatsAbsValueAndPrev) { m.Data2.Followers = v }
func (m *TLStatsBroadcastStats) GetFollowers() *StatsAbsValueAndPrev  { return m.Data2.Followers }

func (m *TLStatsBroadcastStats) SetViewsPerPost(v *StatsAbsValueAndPrev) { m.Data2.ViewsPerPost = v }
func (m *TLStatsBroadcastStats) GetViewsPerPost() *StatsAbsValueAndPrev  { return m.Data2.ViewsPerPost }

func (m *TLStatsBroadcastStats) SetSharesPerPost(v *StatsAbsValueAndPrev) { m.Data2.SharesPerPost = v }
func (m *TLStatsBroadcastStats) GetSharesPerPost() *StatsAbsValueAndPrev {
	return m.Data2.SharesPerPost
}

func (m *TLStatsBroadcastStats) SetEnabledNotifications(v *StatsPercentValue) {
	m.Data2.EnabledNotifications = v
}
func (m *TLStatsBroadcastStats) GetEnabledNotifications() *StatsPercentValue {
	return m.Data2.EnabledNotifications
}

func (m *TLStatsBroadcastStats) SetGrowthGraph(v *StatsGraph) { m.Data2.GrowthGraph = v }
func (m *TLStatsBroadcastStats) GetGrowthGraph() *StatsGraph  { return m.Data2.GrowthGraph }

func (m *TLStatsBroadcastStats) SetFollowersGraph(v *StatsGraph) { m.Data2.FollowersGraph = v }
func (m *TLStatsBroadcastStats) GetFollowersGraph() *StatsGraph  { return m.Data2.FollowersGraph }

func (m *TLStatsBroadcastStats) SetMuteGraph(v *StatsGraph) { m.Data2.MuteGraph = v }
func (m *TLStatsBroadcastStats) GetMuteGraph() *StatsGraph  { return m.Data2.MuteGraph }

func (m *TLStatsBroadcastStats) SetTopHoursGraph(v *StatsGraph) { m.Data2.TopHoursGraph = v }
func (m *TLStatsBroadcastStats) GetTopHoursGraph() *StatsGraph  { return m.Data2.TopHoursGraph }

func (m *TLStatsBroadcastStats) SetInteractionsGraph(v *StatsGraph) { m.Data2.InteractionsGraph = v }
func (m *TLStatsBroadcastStats) GetInteractionsGraph() *StatsGraph  { return m.Data2.InteractionsGraph }

func (m *TLStatsBroadcastStats) SetIvInteractionsGraph(v *StatsGraph) {
	m.Data2.IvInteractionsGraph = v
}
func (m *TLStatsBroadcastStats) GetIvInteractionsGraph() *StatsGraph {
	return m.Data2.IvInteractionsGraph
}

func (m *TLStatsBroadcastStats) SetViewsBySourceGraph(v *StatsGraph) { m.Data2.ViewsBySourceGraph = v }
func (m *TLStatsBroadcastStats) GetViewsBySourceGraph() *StatsGraph {
	return m.Data2.ViewsBySourceGraph
}

func (m *TLStatsBroadcastStats) SetNewFollowersBySourceGraph(v *StatsGraph) {
	m.Data2.NewFollowersBySourceGraph = v
}
func (m *TLStatsBroadcastStats) GetNewFollowersBySourceGraph() *StatsGraph {
	return m.Data2.NewFollowersBySourceGraph
}

func (m *TLStatsBroadcastStats) SetLanguagesGraph(v *StatsGraph) { m.Data2.LanguagesGraph = v }
func (m *TLStatsBroadcastStats) GetLanguagesGraph() *StatsGraph  { return m.Data2.LanguagesGraph }

func (m *TLStatsBroadcastStats) SetRecentMessageInteractions(v []*MessageInteractionCounters) {
	m.Data2.RecentMessageInteractions = v
}
func (m *TLStatsBroadcastStats) GetRecentMessageInteractions() []*MessageInteractionCounters {
	return m.Data2.RecentMessageInteractions
}

func (m *TLStatsMegagroupStats) SetPeriod(v *StatsDateRangeDays) { m.Data2.Period = v }
func (m *TLStatsMegagroupStats) GetPeriod() *StatsDateRangeDays  { return m.Data2.Period }

func (m *TLStatsMegagroupStats) SetMembers(v *StatsAbsValueAndPrev) { m.Data2.Members = v }
func (m *TLStatsMegagroupStats) GetMembers() *StatsAbsValueAndPrev  { return m.Data2.Members }

func (m *TLStatsMegagroupStats) SetMessages(v *StatsAbsValueAndPrev) { m.Data2.Messages = v }
func (m *TLStatsMegagroupStats) GetMessages() *StatsAbsValueAndPrev  { return m.Data2.Messages }

func (m *TLStatsMegagroupStats) SetViewers(v *StatsAbsValueAndPrev) { m.Data2.Viewers = v }
func (m *TLStatsMegagroupStats) GetViewers() *StatsAbsValueAndPrev  { return m.Data2.Viewers }

func (m *TLStatsMegagroupStats) SetPosters(v *StatsAbsValueAndPrev) { m.Data2.Posters = v }
func (m *TLStatsMegagroupStats) GetPosters() *StatsAbsValueAndPrev  { return m.Data2.Posters }

func (m *TLStatsMegagroupStats) SetGrowthGraph(v *StatsGraph) { m.Data2.GrowthGraph = v }
func (m *TLStatsMegagroupStats) GetGrowthGraph() *StatsGraph  { return m.Data2.GrowthGraph }

func (m *TLStatsMegagroupStats) SetMembersGraph(v *StatsGraph) { m.Data2.MembersGraph = v }
func (m *TLStatsMegagroupStats) GetMembersGraph() *StatsGraph  { return m.Data2.MembersGraph }

func (m *TLStatsMegagroupStats) SetNewMembersBySourceGraph(v *StatsGraph) {
	m.Data2.NewMembersBySourceGraph = v
}
func (m *TLStatsMegagroupStats) GetNewMembersBySourceGraph() *StatsGraph {
	return m.Data2.NewMembersBySourceGraph
}

func (m *TLStatsMegagroupStats) SetLanguagesGraph(v *StatsGraph) { m.Data2.LanguagesGraph = v }
func (m *TLStatsMegagroupStats) GetLanguagesGraph() *StatsGraph  { return m.Data2.LanguagesGraph }

func (m *TLStatsMegagroupStats) SetMessagesGraph(v *StatsGraph) { m.Data2.MessagesGraph = v }
func (m *TLStatsMegagroupStats) GetMessagesGraph() *StatsGraph  { return m.Data2.MessagesGraph }

func (m *TLStatsMegagroupStats) SetActionsGraph(v *StatsGraph) { m.Data2.ActionsGraph = v }
func (m *TLStatsMegagroupStats) GetActionsGraph() *StatsGraph  { return m.Data2.ActionsGraph }

func (m *TLStatsMegagroupStats) SetTopHoursGraph(v *StatsGraph) { m.Data2.TopHoursGraph = v }
func (m *TLStatsMegagroupStats) GetTopHoursGraph() *StatsGraph  { return m.Data2.TopHoursGraph }

func (m *TLStatsMegagroupStats) SetWeekdaysGraph(v *StatsGraph) { m.Data2.WeekdaysGraph = v }
func (m *TLStatsMegagroupStats) GetWeekdaysGraph() *StatsGraph  { return m.Data2.WeekdaysGraph }

func (m *TLStatsMegagroupStats) SetTopPosters(v []*StatsGroupTopPoster) { m.Data2.TopPosters = v }
func (m *TLStatsMegagroupStats) GetTopPosters() []*StatsGroupTopPoster  { return m.Data2.TopPosters }

func (m *TLStatsMegagroupStats) SetTopAdmins(v []*StatsGroupTopAdmin) { m.Data2.TopAdmins = v }
func (m *TLStatsMegagroupStats) GetTopAdmins() []*StatsGroupTopAdmin  { return m.Data2.TopAdmins }

func (m *TLStatsMegagroupStats) SetTopInviters(v []*StatsGroupTopInviter) { m.Data2.TopInviters = v }
func (m *TLStatsMegagroupStats) GetTopInviters() []*StatsGroupTopInviter  { return m.Data2.TopInviters }

func (m *TLStatsMegagroupStats) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLStatsMegagroupStats) GetUsers() []*User  { return m.Data2.Users }

func (m *TLStatsMessageStats) SetViewsGraph(v *StatsGraph) { m.Data2.ViewsGraph = v }
func (m *TLStatsMessageStats) GetViewsGraph() *StatsGraph  { return m.Data2.ViewsGraph }

func (m *TLStickerPack) SetEmoticon(v string) { m.Data2.Emoticon = v }
func (m *TLStickerPack) GetEmoticon() string  { return m.Data2.Emoticon }

func (m *TLStickerPack) SetDocuments(v []int64) { m.Data2.Documents = v }
func (m *TLStickerPack) GetDocuments() []int64  { return m.Data2.Documents }

func (m *TLStickerSet) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLStickerSet) GetFlags() int32  { return m.Data2.Flags }

func (m *TLStickerSet) SetInstalled(v bool) { m.Data2.Installed = v }
func (m *TLStickerSet) GetInstalled() bool  { return m.Data2.Installed }

func (m *TLStickerSet) SetArchived(v bool) { m.Data2.Archived = v }
func (m *TLStickerSet) GetArchived() bool  { return m.Data2.Archived }

func (m *TLStickerSet) SetOfficial(v bool) { m.Data2.Official = v }
func (m *TLStickerSet) GetOfficial() bool  { return m.Data2.Official }

func (m *TLStickerSet) SetMasks(v bool) { m.Data2.Masks = v }
func (m *TLStickerSet) GetMasks() bool  { return m.Data2.Masks }

func (m *TLStickerSet) SetId(v int64) { m.Data2.Id = v }
func (m *TLStickerSet) GetId() int64  { return m.Data2.Id }

func (m *TLStickerSet) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLStickerSet) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLStickerSet) SetTitle(v string) { m.Data2.Title = v }
func (m *TLStickerSet) GetTitle() string  { return m.Data2.Title }

func (m *TLStickerSet) SetShortName(v string) { m.Data2.ShortName = v }
func (m *TLStickerSet) GetShortName() string  { return m.Data2.ShortName }

func (m *TLStickerSet) SetCount(v int32) { m.Data2.Count = v }
func (m *TLStickerSet) GetCount() int32  { return m.Data2.Count }

func (m *TLStickerSet) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLStickerSet) GetHash() int32  { return m.Data2.Hash }

func (m *TLStickerSet) SetInstalledDate(v int32) { m.Data2.InstalledDate = v }
func (m *TLStickerSet) GetInstalledDate() int32  { return m.Data2.InstalledDate }

func (m *TLStickerSet) SetThumb(v *PhotoSize) { m.Data2.Thumb = v }
func (m *TLStickerSet) GetThumb() *PhotoSize  { return m.Data2.Thumb }

func (m *TLStickerSet) SetThumbDcId(v int32) { m.Data2.ThumbDcId = v }
func (m *TLStickerSet) GetThumbDcId() int32  { return m.Data2.ThumbDcId }

func (m *TLStickerSet) SetAnimated(v bool) { m.Data2.Animated = v }
func (m *TLStickerSet) GetAnimated() bool  { return m.Data2.Animated }

func (m *TLStickerSet) SetThumbs(v []*PhotoSize) { m.Data2.Thumbs = v }
func (m *TLStickerSet) GetThumbs() []*PhotoSize  { return m.Data2.Thumbs }

func (m *TLStickerSetCovered) SetSet(v *StickerSet) { m.Data2.Set = v }
func (m *TLStickerSetCovered) GetSet() *StickerSet  { return m.Data2.Set }

func (m *TLStickerSetCovered) SetCover(v *Document) { m.Data2.Cover = v }
func (m *TLStickerSetCovered) GetCover() *Document  { return m.Data2.Cover }

func (m *TLStickerSetMultiCovered) SetSet(v *StickerSet) { m.Data2.Set = v }
func (m *TLStickerSetMultiCovered) GetSet() *StickerSet  { return m.Data2.Set }

func (m *TLStickerSetMultiCovered) SetCovers(v []*Document) { m.Data2.Covers = v }
func (m *TLStickerSetMultiCovered) GetCovers() []*Document  { return m.Data2.Covers }

func (m *TLTheme) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLTheme) GetFlags() int32  { return m.Data2.Flags }

func (m *TLTheme) SetCreator(v bool) { m.Data2.Creator = v }
func (m *TLTheme) GetCreator() bool  { return m.Data2.Creator }

func (m *TLTheme) SetDefault(v bool) { m.Data2.Default = v }
func (m *TLTheme) GetDefault() bool  { return m.Data2.Default }

func (m *TLTheme) SetId(v int64) { m.Data2.Id = v }
func (m *TLTheme) GetId() int64  { return m.Data2.Id }

func (m *TLTheme) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLTheme) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLTheme) SetSlug(v string) { m.Data2.Slug = v }
func (m *TLTheme) GetSlug() string  { return m.Data2.Slug }

func (m *TLTheme) SetTitle(v string) { m.Data2.Title = v }
func (m *TLTheme) GetTitle() string  { return m.Data2.Title }

func (m *TLTheme) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLTheme) GetDocument() *Document  { return m.Data2.Document }

func (m *TLTheme) SetInstallsCount(v int32) { m.Data2.InstallsCount = v }
func (m *TLTheme) GetInstallsCount() int32  { return m.Data2.InstallsCount }

func (m *TLTheme) SetSettings(v *ThemeSettings) { m.Data2.Settings = v }
func (m *TLTheme) GetSettings() *ThemeSettings  { return m.Data2.Settings }

func (m *TLThemeSettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLThemeSettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLThemeSettings) SetBaseTheme(v *BaseTheme) { m.Data2.BaseTheme = v }
func (m *TLThemeSettings) GetBaseTheme() *BaseTheme  { return m.Data2.BaseTheme }

func (m *TLThemeSettings) SetAccentColor(v int32) { m.Data2.AccentColor = v }
func (m *TLThemeSettings) GetAccentColor() int32  { return m.Data2.AccentColor }

func (m *TLThemeSettings) SetMessageTopColor(v int32) { m.Data2.MessageTopColor = v }
func (m *TLThemeSettings) GetMessageTopColor() int32  { return m.Data2.MessageTopColor }

func (m *TLThemeSettings) SetMessageBottomColor(v int32) { m.Data2.MessageBottomColor = v }
func (m *TLThemeSettings) GetMessageBottomColor() int32  { return m.Data2.MessageBottomColor }

func (m *TLThemeSettings) SetWallpaper(v *WallPaper) { m.Data2.Wallpaper = v }
func (m *TLThemeSettings) GetWallpaper() *WallPaper  { return m.Data2.Wallpaper }

func (m *TLTopPeer) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLTopPeer) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLTopPeer) SetRating(v float64) { m.Data2.Rating = v }
func (m *TLTopPeer) GetRating() float64  { return m.Data2.Rating }

func (m *TLTopPeerCategoryPeers) SetCategory(v *TopPeerCategory) { m.Data2.Category = v }
func (m *TLTopPeerCategoryPeers) GetCategory() *TopPeerCategory  { return m.Data2.Category }

func (m *TLTopPeerCategoryPeers) SetCount(v int32) { m.Data2.Count = v }
func (m *TLTopPeerCategoryPeers) GetCount() int32  { return m.Data2.Count }

func (m *TLTopPeerCategoryPeers) SetPeers(v []*TopPeer) { m.Data2.Peers = v }
func (m *TLTopPeerCategoryPeers) GetPeers() []*TopPeer  { return m.Data2.Peers }

func (m *TLUpdateNewMessage) SetMessage1F2B0AFD71(v *Message) { m.Data2.Message1F2B0AFD71 = v }
func (m *TLUpdateNewMessage) GetMessage1F2B0AFD71() *Message  { return m.Data2.Message1F2B0AFD71 }

func (m *TLUpdateNewMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateNewMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateNewMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateNewMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateMessageID) SetId4E90BFD671(v int32) { m.Data2.Id4E90BFD671 = v }
func (m *TLUpdateMessageID) GetId4E90BFD671() int32  { return m.Data2.Id4E90BFD671 }

func (m *TLUpdateMessageID) SetRandomId(v int64) { m.Data2.RandomId = v }
func (m *TLUpdateMessageID) GetRandomId() int64  { return m.Data2.RandomId }

func (m *TLUpdateDeleteMessages) SetMessages(v []int32) { m.Data2.Messages = v }
func (m *TLUpdateDeleteMessages) GetMessages() []int32  { return m.Data2.Messages }

func (m *TLUpdateDeleteMessages) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateDeleteMessages) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateDeleteMessages) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateDeleteMessages) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateUserTyping) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserTyping) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserTyping) SetAction(v *SendMessageAction) { m.Data2.Action = v }
func (m *TLUpdateUserTyping) GetAction() *SendMessageAction  { return m.Data2.Action }

func (m *TLUpdateChatUserTyping) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChatUserTyping) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateChatUserTyping) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateChatUserTyping) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateChatUserTyping) SetAction(v *SendMessageAction) { m.Data2.Action = v }
func (m *TLUpdateChatUserTyping) GetAction() *SendMessageAction  { return m.Data2.Action }

func (m *TLUpdateChatParticipants) SetParticipants776119871(v *ChatParticipants) {
	m.Data2.Participants776119871 = v
}
func (m *TLUpdateChatParticipants) GetParticipants776119871() *ChatParticipants {
	return m.Data2.Participants776119871
}

func (m *TLUpdateUserStatus) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserStatus) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserStatus) SetStatus(v *UserStatus) { m.Data2.Status = v }
func (m *TLUpdateUserStatus) GetStatus() *UserStatus  { return m.Data2.Status }

func (m *TLUpdateUserName) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserName) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserName) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLUpdateUserName) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLUpdateUserName) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLUpdateUserName) GetLastName() string  { return m.Data2.LastName }

func (m *TLUpdateUserName) SetUsername(v string) { m.Data2.Username = v }
func (m *TLUpdateUserName) GetUsername() string  { return m.Data2.Username }

func (m *TLUpdateUserPhoto) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserPhoto) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserPhoto) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateUserPhoto) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateUserPhoto) SetPhoto(v *UserProfilePhoto) { m.Data2.Photo = v }
func (m *TLUpdateUserPhoto) GetPhoto() *UserProfilePhoto  { return m.Data2.Photo }

func (m *TLUpdateUserPhoto) SetPrevious(v *Bool) { m.Data2.Previous = v }
func (m *TLUpdateUserPhoto) GetPrevious() *Bool  { return m.Data2.Previous }

func (m *TLUpdateContactRegistered) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateContactRegistered) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateContactRegistered) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateContactRegistered) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateContactLink) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateContactLink) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateContactLink) SetMyLink(v *ContactLink) { m.Data2.MyLink = v }
func (m *TLUpdateContactLink) GetMyLink() *ContactLink  { return m.Data2.MyLink }

func (m *TLUpdateContactLink) SetForeignLink(v *ContactLink) { m.Data2.ForeignLink = v }
func (m *TLUpdateContactLink) GetForeignLink() *ContactLink  { return m.Data2.ForeignLink }

func (m *TLUpdateNewEncryptedMessage) SetMessage12BCBD9A71(v *EncryptedMessage) {
	m.Data2.Message12BCBD9A71 = v
}
func (m *TLUpdateNewEncryptedMessage) GetMessage12BCBD9A71() *EncryptedMessage {
	return m.Data2.Message12BCBD9A71
}

func (m *TLUpdateNewEncryptedMessage) SetQts(v int32) { m.Data2.Qts = v }
func (m *TLUpdateNewEncryptedMessage) GetQts() int32  { return m.Data2.Qts }

func (m *TLUpdateEncryptedChatTyping) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateEncryptedChatTyping) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateEncryption) SetChat(v *EncryptedChat) { m.Data2.Chat = v }
func (m *TLUpdateEncryption) GetChat() *EncryptedChat  { return m.Data2.Chat }

func (m *TLUpdateEncryption) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateEncryption) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateEncryptedMessagesRead) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateEncryptedMessagesRead) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateEncryptedMessagesRead) SetMaxDate(v int32) { m.Data2.MaxDate = v }
func (m *TLUpdateEncryptedMessagesRead) GetMaxDate() int32  { return m.Data2.MaxDate }

func (m *TLUpdateEncryptedMessagesRead) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateEncryptedMessagesRead) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateChatParticipantAdd) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChatParticipantAdd) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateChatParticipantAdd) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateChatParticipantAdd) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateChatParticipantAdd) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLUpdateChatParticipantAdd) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLUpdateChatParticipantAdd) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateChatParticipantAdd) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateChatParticipantAdd) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLUpdateChatParticipantAdd) GetVersion() int32  { return m.Data2.Version }

func (m *TLUpdateChatParticipantDelete) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChatParticipantDelete) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateChatParticipantDelete) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateChatParticipantDelete) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateChatParticipantDelete) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLUpdateChatParticipantDelete) GetVersion() int32  { return m.Data2.Version }

func (m *TLUpdateDcOptions) SetDcOptions(v []*DcOption) { m.Data2.DcOptions = v }
func (m *TLUpdateDcOptions) GetDcOptions() []*DcOption  { return m.Data2.DcOptions }

func (m *TLUpdateUserBlocked) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserBlocked) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserBlocked) SetBlocked(v *Bool) { m.Data2.Blocked = v }
func (m *TLUpdateUserBlocked) GetBlocked() *Bool  { return m.Data2.Blocked }

func (m *TLUpdateNotifySettings) SetPeerBEC268EF71(v *NotifyPeer) { m.Data2.PeerBEC268EF71 = v }
func (m *TLUpdateNotifySettings) GetPeerBEC268EF71() *NotifyPeer  { return m.Data2.PeerBEC268EF71 }

func (m *TLUpdateNotifySettings) SetNotifySettings(v *PeerNotifySettings) { m.Data2.NotifySettings = v }
func (m *TLUpdateNotifySettings) GetNotifySettings() *PeerNotifySettings {
	return m.Data2.NotifySettings
}

func (m *TLUpdateServiceNotification) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateServiceNotification) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateServiceNotification) SetPopupEBE4681971(v bool) { m.Data2.PopupEBE4681971 = v }
func (m *TLUpdateServiceNotification) GetPopupEBE4681971() bool  { return m.Data2.PopupEBE4681971 }

func (m *TLUpdateServiceNotification) SetInboxDate(v int32) { m.Data2.InboxDate = v }
func (m *TLUpdateServiceNotification) GetInboxDate() int32  { return m.Data2.InboxDate }

func (m *TLUpdateServiceNotification) SetType(v string) { m.Data2.Type = v }
func (m *TLUpdateServiceNotification) GetType() string  { return m.Data2.Type }

func (m *TLUpdateServiceNotification) SetMessageEBE4681971(v string) { m.Data2.MessageEBE4681971 = v }
func (m *TLUpdateServiceNotification) GetMessageEBE4681971() string  { return m.Data2.MessageEBE4681971 }

func (m *TLUpdateServiceNotification) SetMedia(v *MessageMedia) { m.Data2.Media = v }
func (m *TLUpdateServiceNotification) GetMedia() *MessageMedia  { return m.Data2.Media }

func (m *TLUpdateServiceNotification) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLUpdateServiceNotification) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLUpdateServiceNotification) SetPopup382DD3E451(v *Bool) { m.Data2.Popup382DD3E451 = v }
func (m *TLUpdateServiceNotification) GetPopup382DD3E451() *Bool  { return m.Data2.Popup382DD3E451 }

func (m *TLUpdatePrivacy) SetKey(v *PrivacyKey) { m.Data2.Key = v }
func (m *TLUpdatePrivacy) GetKey() *PrivacyKey  { return m.Data2.Key }

func (m *TLUpdatePrivacy) SetRules(v []*PrivacyRule) { m.Data2.Rules = v }
func (m *TLUpdatePrivacy) GetRules() []*PrivacyRule  { return m.Data2.Rules }

func (m *TLUpdateUserPhone) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserPhone) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserPhone) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLUpdateUserPhone) GetPhone() string  { return m.Data2.Phone }

func (m *TLUpdateReadHistoryInbox) SetPeer9961FD5C71(v *Peer) { m.Data2.Peer9961FD5C71 = v }
func (m *TLUpdateReadHistoryInbox) GetPeer9961FD5C71() *Peer  { return m.Data2.Peer9961FD5C71 }

func (m *TLUpdateReadHistoryInbox) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLUpdateReadHistoryInbox) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLUpdateReadHistoryInbox) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateReadHistoryInbox) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateReadHistoryInbox) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateReadHistoryInbox) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateReadHistoryInbox) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateReadHistoryInbox) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateReadHistoryInbox) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLUpdateReadHistoryInbox) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLUpdateReadHistoryInbox) SetStillUnreadCount(v int32) { m.Data2.StillUnreadCount = v }
func (m *TLUpdateReadHistoryInbox) GetStillUnreadCount() int32  { return m.Data2.StillUnreadCount }

func (m *TLUpdateReadHistoryOutbox) SetPeer9961FD5C71(v *Peer) { m.Data2.Peer9961FD5C71 = v }
func (m *TLUpdateReadHistoryOutbox) GetPeer9961FD5C71() *Peer  { return m.Data2.Peer9961FD5C71 }

func (m *TLUpdateReadHistoryOutbox) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLUpdateReadHistoryOutbox) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLUpdateReadHistoryOutbox) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateReadHistoryOutbox) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateReadHistoryOutbox) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateReadHistoryOutbox) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateWebPage) SetWebpage(v *WebPage) { m.Data2.Webpage = v }
func (m *TLUpdateWebPage) GetWebpage() *WebPage  { return m.Data2.Webpage }

func (m *TLUpdateWebPage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateWebPage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateWebPage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateWebPage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateReadMessagesContents) SetMessages(v []int32) { m.Data2.Messages = v }
func (m *TLUpdateReadMessagesContents) GetMessages() []int32  { return m.Data2.Messages }

func (m *TLUpdateReadMessagesContents) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateReadMessagesContents) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateReadMessagesContents) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateReadMessagesContents) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateChannelTooLong) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateChannelTooLong) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateChannelTooLong) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelTooLong) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelTooLong) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateChannelTooLong) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateChannel) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannel) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateNewChannelMessage) SetMessage1F2B0AFD71(v *Message) { m.Data2.Message1F2B0AFD71 = v }
func (m *TLUpdateNewChannelMessage) GetMessage1F2B0AFD71() *Message  { return m.Data2.Message1F2B0AFD71 }

func (m *TLUpdateNewChannelMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateNewChannelMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateNewChannelMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateNewChannelMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateReadChannelInbox) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateReadChannelInbox) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateReadChannelInbox) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLUpdateReadChannelInbox) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLUpdateReadChannelInbox) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateReadChannelInbox) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateReadChannelInbox) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLUpdateReadChannelInbox) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLUpdateReadChannelInbox) SetStillUnreadCount(v int32) { m.Data2.StillUnreadCount = v }
func (m *TLUpdateReadChannelInbox) GetStillUnreadCount() int32  { return m.Data2.StillUnreadCount }

func (m *TLUpdateReadChannelInbox) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateReadChannelInbox) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateDeleteChannelMessages) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateDeleteChannelMessages) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateDeleteChannelMessages) SetMessages(v []int32) { m.Data2.Messages = v }
func (m *TLUpdateDeleteChannelMessages) GetMessages() []int32  { return m.Data2.Messages }

func (m *TLUpdateDeleteChannelMessages) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateDeleteChannelMessages) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateDeleteChannelMessages) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateDeleteChannelMessages) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateChannelMessageViews) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelMessageViews) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelMessageViews) SetId4E90BFD671(v int32) { m.Data2.Id4E90BFD671 = v }
func (m *TLUpdateChannelMessageViews) GetId4E90BFD671() int32  { return m.Data2.Id4E90BFD671 }

func (m *TLUpdateChannelMessageViews) SetViews(v int32) { m.Data2.Views = v }
func (m *TLUpdateChannelMessageViews) GetViews() int32  { return m.Data2.Views }

func (m *TLUpdateChatAdmins) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChatAdmins) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateChatAdmins) SetEnabled(v *Bool) { m.Data2.Enabled = v }
func (m *TLUpdateChatAdmins) GetEnabled() *Bool  { return m.Data2.Enabled }

func (m *TLUpdateChatAdmins) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLUpdateChatAdmins) GetVersion() int32  { return m.Data2.Version }

func (m *TLUpdateChatParticipantAdmin) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChatParticipantAdmin) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateChatParticipantAdmin) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateChatParticipantAdmin) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateChatParticipantAdmin) SetIsAdmin(v *Bool) { m.Data2.IsAdmin = v }
func (m *TLUpdateChatParticipantAdmin) GetIsAdmin() *Bool  { return m.Data2.IsAdmin }

func (m *TLUpdateChatParticipantAdmin) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLUpdateChatParticipantAdmin) GetVersion() int32  { return m.Data2.Version }

func (m *TLUpdateNewStickerSet) SetStickerset(v *Messages_StickerSet) { m.Data2.Stickerset = v }
func (m *TLUpdateNewStickerSet) GetStickerset() *Messages_StickerSet  { return m.Data2.Stickerset }

func (m *TLUpdateStickerSetsOrder) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateStickerSetsOrder) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateStickerSetsOrder) SetMasks(v bool) { m.Data2.Masks = v }
func (m *TLUpdateStickerSetsOrder) GetMasks() bool  { return m.Data2.Masks }

func (m *TLUpdateStickerSetsOrder) SetOrderBB2D20171(v []int64) { m.Data2.OrderBB2D20171 = v }
func (m *TLUpdateStickerSetsOrder) GetOrderBB2D20171() []int64  { return m.Data2.OrderBB2D20171 }

func (m *TLUpdateBotInlineQuery) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateBotInlineQuery) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateBotInlineQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateBotInlineQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateBotInlineQuery) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateBotInlineQuery) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateBotInlineQuery) SetQuery(v string) { m.Data2.Query = v }
func (m *TLUpdateBotInlineQuery) GetQuery() string  { return m.Data2.Query }

func (m *TLUpdateBotInlineQuery) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLUpdateBotInlineQuery) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLUpdateBotInlineQuery) SetOffset(v string) { m.Data2.Offset = v }
func (m *TLUpdateBotInlineQuery) GetOffset() string  { return m.Data2.Offset }

func (m *TLUpdateBotInlineQuery) SetPeerType(v *InlineQueryPeerType) { m.Data2.PeerType = v }
func (m *TLUpdateBotInlineQuery) GetPeerType() *InlineQueryPeerType  { return m.Data2.PeerType }

func (m *TLUpdateBotInlineSend) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateBotInlineSend) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateBotInlineSend) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateBotInlineSend) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateBotInlineSend) SetQuery(v string) { m.Data2.Query = v }
func (m *TLUpdateBotInlineSend) GetQuery() string  { return m.Data2.Query }

func (m *TLUpdateBotInlineSend) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLUpdateBotInlineSend) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLUpdateBotInlineSend) SetIdE48F96471(v string) { m.Data2.IdE48F96471 = v }
func (m *TLUpdateBotInlineSend) GetIdE48F96471() string  { return m.Data2.IdE48F96471 }

func (m *TLUpdateBotInlineSend) SetMsgIdE48F96471(v *InputBotInlineMessageID) {
	m.Data2.MsgIdE48F96471 = v
}
func (m *TLUpdateBotInlineSend) GetMsgIdE48F96471() *InputBotInlineMessageID {
	return m.Data2.MsgIdE48F96471
}

func (m *TLUpdateEditChannelMessage) SetMessage1F2B0AFD71(v *Message) { m.Data2.Message1F2B0AFD71 = v }
func (m *TLUpdateEditChannelMessage) GetMessage1F2B0AFD71() *Message {
	return m.Data2.Message1F2B0AFD71
}

func (m *TLUpdateEditChannelMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateEditChannelMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateEditChannelMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateEditChannelMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateChannelPinnedMessage) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelPinnedMessage) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelPinnedMessage) SetId4E90BFD671(v int32) { m.Data2.Id4E90BFD671 = v }
func (m *TLUpdateChannelPinnedMessage) GetId4E90BFD671() int32  { return m.Data2.Id4E90BFD671 }

func (m *TLUpdateBotCallbackQuery) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateBotCallbackQuery) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateBotCallbackQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateBotCallbackQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateBotCallbackQuery) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateBotCallbackQuery) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateBotCallbackQuery) SetPeer9961FD5C71(v *Peer) { m.Data2.Peer9961FD5C71 = v }
func (m *TLUpdateBotCallbackQuery) GetPeer9961FD5C71() *Peer  { return m.Data2.Peer9961FD5C71 }

func (m *TLUpdateBotCallbackQuery) SetMsgIdE73547E171(v int32) { m.Data2.MsgIdE73547E171 = v }
func (m *TLUpdateBotCallbackQuery) GetMsgIdE73547E171() int32  { return m.Data2.MsgIdE73547E171 }

func (m *TLUpdateBotCallbackQuery) SetChatInstance(v int64) { m.Data2.ChatInstance = v }
func (m *TLUpdateBotCallbackQuery) GetChatInstance() int64  { return m.Data2.ChatInstance }

func (m *TLUpdateBotCallbackQuery) SetDataE73547E171(v []byte) { m.Data2.DataE73547E171 = v }
func (m *TLUpdateBotCallbackQuery) GetDataE73547E171() []byte  { return m.Data2.DataE73547E171 }

func (m *TLUpdateBotCallbackQuery) SetGameShortName(v string) { m.Data2.GameShortName = v }
func (m *TLUpdateBotCallbackQuery) GetGameShortName() string  { return m.Data2.GameShortName }

func (m *TLUpdateEditMessage) SetMessage1F2B0AFD71(v *Message) { m.Data2.Message1F2B0AFD71 = v }
func (m *TLUpdateEditMessage) GetMessage1F2B0AFD71() *Message  { return m.Data2.Message1F2B0AFD71 }

func (m *TLUpdateEditMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateEditMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateEditMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateEditMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateInlineBotCallbackQuery) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateInlineBotCallbackQuery) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateInlineBotCallbackQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateInlineBotCallbackQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateInlineBotCallbackQuery) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateInlineBotCallbackQuery) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateInlineBotCallbackQuery) SetMsgIdE48F96471(v *InputBotInlineMessageID) {
	m.Data2.MsgIdE48F96471 = v
}
func (m *TLUpdateInlineBotCallbackQuery) GetMsgIdE48F96471() *InputBotInlineMessageID {
	return m.Data2.MsgIdE48F96471
}

func (m *TLUpdateInlineBotCallbackQuery) SetChatInstance(v int64) { m.Data2.ChatInstance = v }
func (m *TLUpdateInlineBotCallbackQuery) GetChatInstance() int64  { return m.Data2.ChatInstance }

func (m *TLUpdateInlineBotCallbackQuery) SetDataE73547E171(v []byte) { m.Data2.DataE73547E171 = v }
func (m *TLUpdateInlineBotCallbackQuery) GetDataE73547E171() []byte  { return m.Data2.DataE73547E171 }

func (m *TLUpdateInlineBotCallbackQuery) SetGameShortName(v string) { m.Data2.GameShortName = v }
func (m *TLUpdateInlineBotCallbackQuery) GetGameShortName() string  { return m.Data2.GameShortName }

func (m *TLUpdateReadChannelOutbox) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateReadChannelOutbox) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateReadChannelOutbox) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLUpdateReadChannelOutbox) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLUpdateDraftMessage) SetPeer9961FD5C71(v *Peer) { m.Data2.Peer9961FD5C71 = v }
func (m *TLUpdateDraftMessage) GetPeer9961FD5C71() *Peer  { return m.Data2.Peer9961FD5C71 }

func (m *TLUpdateDraftMessage) SetDraft(v *DraftMessage) { m.Data2.Draft = v }
func (m *TLUpdateDraftMessage) GetDraft() *DraftMessage  { return m.Data2.Draft }

func (m *TLUpdateChannelWebPage) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelWebPage) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelWebPage) SetWebpage(v *WebPage) { m.Data2.Webpage = v }
func (m *TLUpdateChannelWebPage) GetWebpage() *WebPage  { return m.Data2.Webpage }

func (m *TLUpdateChannelWebPage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateChannelWebPage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateChannelWebPage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateChannelWebPage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateDialogPinned) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateDialogPinned) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateDialogPinned) SetPinned(v bool) { m.Data2.Pinned = v }
func (m *TLUpdateDialogPinned) GetPinned() bool  { return m.Data2.Pinned }

func (m *TLUpdateDialogPinned) SetPeer9961FD5C71(v *Peer) { m.Data2.Peer9961FD5C71 = v }
func (m *TLUpdateDialogPinned) GetPeer9961FD5C71() *Peer  { return m.Data2.Peer9961FD5C71 }

func (m *TLUpdateDialogPinned) SetPeer19D27F3C85(v *DialogPeer) { m.Data2.Peer19D27F3C85 = v }
func (m *TLUpdateDialogPinned) GetPeer19D27F3C85() *DialogPeer  { return m.Data2.Peer19D27F3C85 }

func (m *TLUpdateDialogPinned) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLUpdateDialogPinned) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLUpdatePinnedDialogs) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdatePinnedDialogs) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdatePinnedDialogs) SetOrderD8CAF68D71(v []*Peer) { m.Data2.OrderD8CAF68D71 = v }
func (m *TLUpdatePinnedDialogs) GetOrderD8CAF68D71() []*Peer  { return m.Data2.OrderD8CAF68D71 }

func (m *TLUpdatePinnedDialogs) SetOrderEA4CB65B85(v []*DialogPeer) { m.Data2.OrderEA4CB65B85 = v }
func (m *TLUpdatePinnedDialogs) GetOrderEA4CB65B85() []*DialogPeer  { return m.Data2.OrderEA4CB65B85 }

func (m *TLUpdatePinnedDialogs) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLUpdatePinnedDialogs) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLUpdateBotWebhookJSON) SetData8317C0C371(v *DataJSON) { m.Data2.Data8317C0C371 = v }
func (m *TLUpdateBotWebhookJSON) GetData8317C0C371() *DataJSON  { return m.Data2.Data8317C0C371 }

func (m *TLUpdateBotWebhookJSONQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateBotWebhookJSONQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateBotWebhookJSONQuery) SetData8317C0C371(v *DataJSON) { m.Data2.Data8317C0C371 = v }
func (m *TLUpdateBotWebhookJSONQuery) GetData8317C0C371() *DataJSON  { return m.Data2.Data8317C0C371 }

func (m *TLUpdateBotWebhookJSONQuery) SetTimeout(v int32) { m.Data2.Timeout = v }
func (m *TLUpdateBotWebhookJSONQuery) GetTimeout() int32  { return m.Data2.Timeout }

func (m *TLUpdateBotShippingQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateBotShippingQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateBotShippingQuery) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateBotShippingQuery) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateBotShippingQuery) SetPayload(v []byte) { m.Data2.Payload = v }
func (m *TLUpdateBotShippingQuery) GetPayload() []byte  { return m.Data2.Payload }

func (m *TLUpdateBotShippingQuery) SetShippingAddress(v *PostAddress) { m.Data2.ShippingAddress = v }
func (m *TLUpdateBotShippingQuery) GetShippingAddress() *PostAddress  { return m.Data2.ShippingAddress }

func (m *TLUpdateBotPrecheckoutQuery) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateBotPrecheckoutQuery) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateBotPrecheckoutQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateBotPrecheckoutQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateBotPrecheckoutQuery) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateBotPrecheckoutQuery) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateBotPrecheckoutQuery) SetPayload(v []byte) { m.Data2.Payload = v }
func (m *TLUpdateBotPrecheckoutQuery) GetPayload() []byte  { return m.Data2.Payload }

func (m *TLUpdateBotPrecheckoutQuery) SetInfo(v *PaymentRequestedInfo) { m.Data2.Info = v }
func (m *TLUpdateBotPrecheckoutQuery) GetInfo() *PaymentRequestedInfo  { return m.Data2.Info }

func (m *TLUpdateBotPrecheckoutQuery) SetShippingOptionId(v string) { m.Data2.ShippingOptionId = v }
func (m *TLUpdateBotPrecheckoutQuery) GetShippingOptionId() string  { return m.Data2.ShippingOptionId }

func (m *TLUpdateBotPrecheckoutQuery) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLUpdateBotPrecheckoutQuery) GetCurrency() string  { return m.Data2.Currency }

func (m *TLUpdateBotPrecheckoutQuery) SetTotalAmount(v int64) { m.Data2.TotalAmount = v }
func (m *TLUpdateBotPrecheckoutQuery) GetTotalAmount() int64  { return m.Data2.TotalAmount }

func (m *TLUpdatePhoneCall) SetPhoneCall(v *PhoneCall) { m.Data2.PhoneCall = v }
func (m *TLUpdatePhoneCall) GetPhoneCall() *PhoneCall  { return m.Data2.PhoneCall }

func (m *TLUpdateLangPackTooLong) SetLangCode(v string) { m.Data2.LangCode = v }
func (m *TLUpdateLangPackTooLong) GetLangCode() string  { return m.Data2.LangCode }

func (m *TLUpdateLangPack) SetDifference(v *LangPackDifference) { m.Data2.Difference = v }
func (m *TLUpdateLangPack) GetDifference() *LangPackDifference  { return m.Data2.Difference }

func (m *TLUpdateChannelReadMessagesContents) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelReadMessagesContents) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelReadMessagesContents) SetMessages(v []int32) { m.Data2.Messages = v }
func (m *TLUpdateChannelReadMessagesContents) GetMessages() []int32  { return m.Data2.Messages }

func (m *TLUpdateNewAuthorization) SetAuthKeyId(v int64) { m.Data2.AuthKeyId = v }
func (m *TLUpdateNewAuthorization) GetAuthKeyId() int64  { return m.Data2.AuthKeyId }

func (m *TLUpdateNewAuthorization) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateNewAuthorization) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateNewAuthorization) SetDevice(v string) { m.Data2.Device = v }
func (m *TLUpdateNewAuthorization) GetDevice() string  { return m.Data2.Device }

func (m *TLUpdateNewAuthorization) SetLocation(v string) { m.Data2.Location = v }
func (m *TLUpdateNewAuthorization) GetLocation() string  { return m.Data2.Location }

func (m *TLUpdateChannelGroup) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelGroup) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelGroup) SetGroup(v *MessageGroup) { m.Data2.Group = v }
func (m *TLUpdateChannelGroup) GetGroup() *MessageGroup  { return m.Data2.Group }

func (m *TLUpdateChannelAvailableMessages) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelAvailableMessages) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelAvailableMessages) SetAvailableMinId(v int32) { m.Data2.AvailableMinId = v }
func (m *TLUpdateChannelAvailableMessages) GetAvailableMinId() int32  { return m.Data2.AvailableMinId }

func (m *TLUpdateDialogUnreadMark) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateDialogUnreadMark) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateDialogUnreadMark) SetUnread(v bool) { m.Data2.Unread = v }
func (m *TLUpdateDialogUnreadMark) GetUnread() bool  { return m.Data2.Unread }

func (m *TLUpdateDialogUnreadMark) SetPeer19D27F3C85(v *DialogPeer) { m.Data2.Peer19D27F3C85 = v }
func (m *TLUpdateDialogUnreadMark) GetPeer19D27F3C85() *DialogPeer  { return m.Data2.Peer19D27F3C85 }

func (m *TLUpdateUserPinnedMessage) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserPinnedMessage) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserPinnedMessage) SetId4E90BFD671(v int32) { m.Data2.Id4E90BFD671 = v }
func (m *TLUpdateUserPinnedMessage) GetId4E90BFD671() int32  { return m.Data2.Id4E90BFD671 }

func (m *TLUpdateChatPinnedMessage) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChatPinnedMessage) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateChatPinnedMessage) SetId4E90BFD671(v int32) { m.Data2.Id4E90BFD671 = v }
func (m *TLUpdateChatPinnedMessage) GetId4E90BFD671() int32  { return m.Data2.Id4E90BFD671 }

func (m *TLUpdateChatPinnedMessage) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLUpdateChatPinnedMessage) GetVersion() int32  { return m.Data2.Version }

func (m *TLUpdateMessagePoll) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateMessagePoll) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateMessagePoll) SetPollId(v int64) { m.Data2.PollId = v }
func (m *TLUpdateMessagePoll) GetPollId() int64  { return m.Data2.PollId }

func (m *TLUpdateMessagePoll) SetPoll(v *Poll) { m.Data2.Poll = v }
func (m *TLUpdateMessagePoll) GetPoll() *Poll  { return m.Data2.Poll }

func (m *TLUpdateMessagePoll) SetResults(v *PollResults) { m.Data2.Results = v }
func (m *TLUpdateMessagePoll) GetResults() *PollResults  { return m.Data2.Results }

func (m *TLUpdateChatDefaultBannedRights) SetPeer9961FD5C71(v *Peer) { m.Data2.Peer9961FD5C71 = v }
func (m *TLUpdateChatDefaultBannedRights) GetPeer9961FD5C71() *Peer  { return m.Data2.Peer9961FD5C71 }

func (m *TLUpdateChatDefaultBannedRights) SetDefaultBannedRights(v *ChatBannedRights) {
	m.Data2.DefaultBannedRights = v
}
func (m *TLUpdateChatDefaultBannedRights) GetDefaultBannedRights() *ChatBannedRights {
	return m.Data2.DefaultBannedRights
}

func (m *TLUpdateChatDefaultBannedRights) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLUpdateChatDefaultBannedRights) GetVersion() int32  { return m.Data2.Version }

func (m *TLUpdateFolderPeers) SetFolderPeers(v []*FolderPeer) { m.Data2.FolderPeers = v }
func (m *TLUpdateFolderPeers) GetFolderPeers() []*FolderPeer  { return m.Data2.FolderPeers }

func (m *TLUpdateFolderPeers) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateFolderPeers) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateFolderPeers) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateFolderPeers) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdatePeerSettings) SetPeer9961FD5C71(v *Peer) { m.Data2.Peer9961FD5C71 = v }
func (m *TLUpdatePeerSettings) GetPeer9961FD5C71() *Peer  { return m.Data2.Peer9961FD5C71 }

func (m *TLUpdatePeerSettings) SetSettings(v *PeerSettings) { m.Data2.Settings = v }
func (m *TLUpdatePeerSettings) GetSettings() *PeerSettings  { return m.Data2.Settings }

func (m *TLUpdatePeerLocated) SetPeers(v []*PeerLocated) { m.Data2.Peers = v }
func (m *TLUpdatePeerLocated) GetPeers() []*PeerLocated  { return m.Data2.Peers }

func (m *TLUpdateNewScheduledMessage) SetMessage1F2B0AFD71(v *Message) { m.Data2.Message1F2B0AFD71 = v }
func (m *TLUpdateNewScheduledMessage) GetMessage1F2B0AFD71() *Message {
	return m.Data2.Message1F2B0AFD71
}

func (m *TLUpdateDeleteScheduledMessages) SetPeer9961FD5C71(v *Peer) { m.Data2.Peer9961FD5C71 = v }
func (m *TLUpdateDeleteScheduledMessages) GetPeer9961FD5C71() *Peer  { return m.Data2.Peer9961FD5C71 }

func (m *TLUpdateDeleteScheduledMessages) SetMessages(v []int32) { m.Data2.Messages = v }
func (m *TLUpdateDeleteScheduledMessages) GetMessages() []int32  { return m.Data2.Messages }

func (m *TLUpdateTheme) SetTheme(v *Theme) { m.Data2.Theme = v }
func (m *TLUpdateTheme) GetTheme() *Theme  { return m.Data2.Theme }

func (m *TLUpdateGeoLiveViewed) SetPeer9961FD5C71(v *Peer) { m.Data2.Peer9961FD5C71 = v }
func (m *TLUpdateGeoLiveViewed) GetPeer9961FD5C71() *Peer  { return m.Data2.Peer9961FD5C71 }

func (m *TLUpdateGeoLiveViewed) SetMsgIdE73547E171(v int32) { m.Data2.MsgIdE73547E171 = v }
func (m *TLUpdateGeoLiveViewed) GetMsgIdE73547E171() int32  { return m.Data2.MsgIdE73547E171 }

func (m *TLUpdateMessagePollVote) SetPollId(v int64) { m.Data2.PollId = v }
func (m *TLUpdateMessagePollVote) GetPollId() int64  { return m.Data2.PollId }

func (m *TLUpdateMessagePollVote) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateMessagePollVote) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateMessagePollVote) SetOptions(v [][]byte) { m.Data2.Options = v }
func (m *TLUpdateMessagePollVote) GetOptions() [][]byte  { return m.Data2.Options }

func (m *TLUpdateDialogFilter) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateDialogFilter) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateDialogFilter) SetId4E90BFD671(v int32) { m.Data2.Id4E90BFD671 = v }
func (m *TLUpdateDialogFilter) GetId4E90BFD671() int32  { return m.Data2.Id4E90BFD671 }

func (m *TLUpdateDialogFilter) SetFilter(v *DialogFilter) { m.Data2.Filter = v }
func (m *TLUpdateDialogFilter) GetFilter() *DialogFilter  { return m.Data2.Filter }

func (m *TLUpdateDialogFilterOrder) SetOrderA5D72105111(v []int32) { m.Data2.OrderA5D72105111 = v }
func (m *TLUpdateDialogFilterOrder) GetOrderA5D72105111() []int32  { return m.Data2.OrderA5D72105111 }

func (m *TLUpdatePhoneCallSignalingData) SetPhoneCallId(v int64) { m.Data2.PhoneCallId = v }
func (m *TLUpdatePhoneCallSignalingData) GetPhoneCallId() int64  { return m.Data2.PhoneCallId }

func (m *TLUpdatePhoneCallSignalingData) SetDataE73547E171(v []byte) { m.Data2.DataE73547E171 = v }
func (m *TLUpdatePhoneCallSignalingData) GetDataE73547E171() []byte  { return m.Data2.DataE73547E171 }

func (m *TLUpdateChannelParticipant) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateChannelParticipant) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateChannelParticipant) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelParticipant) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelParticipant) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateChannelParticipant) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateChannelParticipant) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateChannelParticipant) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateChannelParticipant) SetPrevParticipant(v *ChannelParticipant) {
	m.Data2.PrevParticipant = v
}
func (m *TLUpdateChannelParticipant) GetPrevParticipant() *ChannelParticipant {
	return m.Data2.PrevParticipant
}

func (m *TLUpdateChannelParticipant) SetNewParticipant(v *ChannelParticipant) {
	m.Data2.NewParticipant = v
}
func (m *TLUpdateChannelParticipant) GetNewParticipant() *ChannelParticipant {
	return m.Data2.NewParticipant
}

func (m *TLUpdateChannelParticipant) SetQts(v int32) { m.Data2.Qts = v }
func (m *TLUpdateChannelParticipant) GetQts() int32  { return m.Data2.Qts }

func (m *TLUpdateChannelMessageForwards) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelMessageForwards) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelMessageForwards) SetId4E90BFD671(v int32) { m.Data2.Id4E90BFD671 = v }
func (m *TLUpdateChannelMessageForwards) GetId4E90BFD671() int32  { return m.Data2.Id4E90BFD671 }

func (m *TLUpdateChannelMessageForwards) SetForwards(v int32) { m.Data2.Forwards = v }
func (m *TLUpdateChannelMessageForwards) GetForwards() int32  { return m.Data2.Forwards }

func (m *TLUpdateReadChannelDiscussionInbox) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateReadChannelDiscussionInbox) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateReadChannelDiscussionInbox) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateReadChannelDiscussionInbox) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateReadChannelDiscussionInbox) SetTopMsgId(v int32) { m.Data2.TopMsgId = v }
func (m *TLUpdateReadChannelDiscussionInbox) GetTopMsgId() int32  { return m.Data2.TopMsgId }

func (m *TLUpdateReadChannelDiscussionInbox) SetReadMaxId(v int32) { m.Data2.ReadMaxId = v }
func (m *TLUpdateReadChannelDiscussionInbox) GetReadMaxId() int32  { return m.Data2.ReadMaxId }

func (m *TLUpdateReadChannelDiscussionInbox) SetBroadcastId(v int32) { m.Data2.BroadcastId = v }
func (m *TLUpdateReadChannelDiscussionInbox) GetBroadcastId() int32  { return m.Data2.BroadcastId }

func (m *TLUpdateReadChannelDiscussionInbox) SetBroadcastPost(v int32) { m.Data2.BroadcastPost = v }
func (m *TLUpdateReadChannelDiscussionInbox) GetBroadcastPost() int32  { return m.Data2.BroadcastPost }

func (m *TLUpdateReadChannelDiscussionOutbox) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateReadChannelDiscussionOutbox) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateReadChannelDiscussionOutbox) SetTopMsgId(v int32) { m.Data2.TopMsgId = v }
func (m *TLUpdateReadChannelDiscussionOutbox) GetTopMsgId() int32  { return m.Data2.TopMsgId }

func (m *TLUpdateReadChannelDiscussionOutbox) SetReadMaxId(v int32) { m.Data2.ReadMaxId = v }
func (m *TLUpdateReadChannelDiscussionOutbox) GetReadMaxId() int32  { return m.Data2.ReadMaxId }

func (m *TLUpdatePeerBlocked) SetPeerId(v *Peer) { m.Data2.PeerId = v }
func (m *TLUpdatePeerBlocked) GetPeerId() *Peer  { return m.Data2.PeerId }

func (m *TLUpdatePeerBlocked) SetBlocked(v *Bool) { m.Data2.Blocked = v }
func (m *TLUpdatePeerBlocked) GetBlocked() *Bool  { return m.Data2.Blocked }

func (m *TLUpdateChannelUserTyping) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateChannelUserTyping) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateChannelUserTyping) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelUserTyping) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelUserTyping) SetTopMsgId(v int32) { m.Data2.TopMsgId = v }
func (m *TLUpdateChannelUserTyping) GetTopMsgId() int32  { return m.Data2.TopMsgId }

func (m *TLUpdateChannelUserTyping) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateChannelUserTyping) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateChannelUserTyping) SetAction(v *SendMessageAction) { m.Data2.Action = v }
func (m *TLUpdateChannelUserTyping) GetAction() *SendMessageAction  { return m.Data2.Action }

func (m *TLUpdatePinnedMessages) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdatePinnedMessages) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdatePinnedMessages) SetPinned(v bool) { m.Data2.Pinned = v }
func (m *TLUpdatePinnedMessages) GetPinned() bool  { return m.Data2.Pinned }

func (m *TLUpdatePinnedMessages) SetPeer9961FD5C71(v *Peer) { m.Data2.Peer9961FD5C71 = v }
func (m *TLUpdatePinnedMessages) GetPeer9961FD5C71() *Peer  { return m.Data2.Peer9961FD5C71 }

func (m *TLUpdatePinnedMessages) SetMessages(v []int32) { m.Data2.Messages = v }
func (m *TLUpdatePinnedMessages) GetMessages() []int32  { return m.Data2.Messages }

func (m *TLUpdatePinnedMessages) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatePinnedMessages) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatePinnedMessages) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdatePinnedMessages) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdatePinnedChannelMessages) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdatePinnedChannelMessages) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdatePinnedChannelMessages) SetPinned(v bool) { m.Data2.Pinned = v }
func (m *TLUpdatePinnedChannelMessages) GetPinned() bool  { return m.Data2.Pinned }

func (m *TLUpdatePinnedChannelMessages) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdatePinnedChannelMessages) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdatePinnedChannelMessages) SetMessages(v []int32) { m.Data2.Messages = v }
func (m *TLUpdatePinnedChannelMessages) GetMessages() []int32  { return m.Data2.Messages }

func (m *TLUpdatePinnedChannelMessages) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatePinnedChannelMessages) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatePinnedChannelMessages) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdatePinnedChannelMessages) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateChat) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChat) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateGroupCallParticipants) SetCallF2EBDB4E122(v *InputGroupCall) {
	m.Data2.CallF2EBDB4E122 = v
}
func (m *TLUpdateGroupCallParticipants) GetCallF2EBDB4E122() *InputGroupCall {
	return m.Data2.CallF2EBDB4E122
}

func (m *TLUpdateGroupCallParticipants) SetParticipantsF2EBDB4E122(v []*GroupCallParticipant) {
	m.Data2.ParticipantsF2EBDB4E122 = v
}
func (m *TLUpdateGroupCallParticipants) GetParticipantsF2EBDB4E122() []*GroupCallParticipant {
	return m.Data2.ParticipantsF2EBDB4E122
}

func (m *TLUpdateGroupCallParticipants) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLUpdateGroupCallParticipants) GetVersion() int32  { return m.Data2.Version }

func (m *TLUpdateGroupCall) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateGroupCall) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateGroupCall) SetCallA45EB99B122(v *GroupCall) { m.Data2.CallA45EB99B122 = v }
func (m *TLUpdateGroupCall) GetCallA45EB99B122() *GroupCall  { return m.Data2.CallA45EB99B122 }

func (m *TLUpdateShortMessage) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateShortMessage) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateShortMessage) SetOut(v bool) { m.Data2.Out = v }
func (m *TLUpdateShortMessage) GetOut() bool  { return m.Data2.Out }

func (m *TLUpdateShortMessage) SetMentioned(v bool) { m.Data2.Mentioned = v }
func (m *TLUpdateShortMessage) GetMentioned() bool  { return m.Data2.Mentioned }

func (m *TLUpdateShortMessage) SetMediaUnread(v bool) { m.Data2.MediaUnread = v }
func (m *TLUpdateShortMessage) GetMediaUnread() bool  { return m.Data2.MediaUnread }

func (m *TLUpdateShortMessage) SetSilent(v bool) { m.Data2.Silent = v }
func (m *TLUpdateShortMessage) GetSilent() bool  { return m.Data2.Silent }

func (m *TLUpdateShortMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLUpdateShortMessage) GetId() int32  { return m.Data2.Id }

func (m *TLUpdateShortMessage) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateShortMessage) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateShortMessage) SetMessage(v string) { m.Data2.Message = v }
func (m *TLUpdateShortMessage) GetMessage() string  { return m.Data2.Message }

func (m *TLUpdateShortMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateShortMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateShortMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateShortMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateShortMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateShortMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateShortMessage) SetFwdFrom(v *MessageFwdHeader) { m.Data2.FwdFrom = v }
func (m *TLUpdateShortMessage) GetFwdFrom() *MessageFwdHeader  { return m.Data2.FwdFrom }

func (m *TLUpdateShortMessage) SetViaBotId(v int32) { m.Data2.ViaBotId = v }
func (m *TLUpdateShortMessage) GetViaBotId() int32  { return m.Data2.ViaBotId }

func (m *TLUpdateShortMessage) SetReplyToMsgId(v int32) { m.Data2.ReplyToMsgId = v }
func (m *TLUpdateShortMessage) GetReplyToMsgId() int32  { return m.Data2.ReplyToMsgId }

func (m *TLUpdateShortMessage) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLUpdateShortMessage) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLUpdateShortMessage) SetReplyTo(v *MessageReplyHeader) { m.Data2.ReplyTo = v }
func (m *TLUpdateShortMessage) GetReplyTo() *MessageReplyHeader  { return m.Data2.ReplyTo }

func (m *TLUpdateShortChatMessage) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateShortChatMessage) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateShortChatMessage) SetOut(v bool) { m.Data2.Out = v }
func (m *TLUpdateShortChatMessage) GetOut() bool  { return m.Data2.Out }

func (m *TLUpdateShortChatMessage) SetMentioned(v bool) { m.Data2.Mentioned = v }
func (m *TLUpdateShortChatMessage) GetMentioned() bool  { return m.Data2.Mentioned }

func (m *TLUpdateShortChatMessage) SetMediaUnread(v bool) { m.Data2.MediaUnread = v }
func (m *TLUpdateShortChatMessage) GetMediaUnread() bool  { return m.Data2.MediaUnread }

func (m *TLUpdateShortChatMessage) SetSilent(v bool) { m.Data2.Silent = v }
func (m *TLUpdateShortChatMessage) GetSilent() bool  { return m.Data2.Silent }

func (m *TLUpdateShortChatMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLUpdateShortChatMessage) GetId() int32  { return m.Data2.Id }

func (m *TLUpdateShortChatMessage) SetFromId(v int32) { m.Data2.FromId = v }
func (m *TLUpdateShortChatMessage) GetFromId() int32  { return m.Data2.FromId }

func (m *TLUpdateShortChatMessage) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateShortChatMessage) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateShortChatMessage) SetMessage(v string) { m.Data2.Message = v }
func (m *TLUpdateShortChatMessage) GetMessage() string  { return m.Data2.Message }

func (m *TLUpdateShortChatMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateShortChatMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateShortChatMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateShortChatMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateShortChatMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateShortChatMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateShortChatMessage) SetFwdFrom(v *MessageFwdHeader) { m.Data2.FwdFrom = v }
func (m *TLUpdateShortChatMessage) GetFwdFrom() *MessageFwdHeader  { return m.Data2.FwdFrom }

func (m *TLUpdateShortChatMessage) SetViaBotId(v int32) { m.Data2.ViaBotId = v }
func (m *TLUpdateShortChatMessage) GetViaBotId() int32  { return m.Data2.ViaBotId }

func (m *TLUpdateShortChatMessage) SetReplyToMsgId(v int32) { m.Data2.ReplyToMsgId = v }
func (m *TLUpdateShortChatMessage) GetReplyToMsgId() int32  { return m.Data2.ReplyToMsgId }

func (m *TLUpdateShortChatMessage) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLUpdateShortChatMessage) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLUpdateShortChatMessage) SetReplyTo(v *MessageReplyHeader) { m.Data2.ReplyTo = v }
func (m *TLUpdateShortChatMessage) GetReplyTo() *MessageReplyHeader  { return m.Data2.ReplyTo }

func (m *TLUpdateShort) SetUpdate(v *Update) { m.Data2.Update = v }
func (m *TLUpdateShort) GetUpdate() *Update  { return m.Data2.Update }

func (m *TLUpdateShort) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateShort) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdatesCombined) SetUpdates(v []*Update) { m.Data2.Updates = v }
func (m *TLUpdatesCombined) GetUpdates() []*Update  { return m.Data2.Updates }

func (m *TLUpdatesCombined) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdatesCombined) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdatesCombined) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdatesCombined) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdatesCombined) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdatesCombined) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdatesCombined) SetSeqStart(v int32) { m.Data2.SeqStart = v }
func (m *TLUpdatesCombined) GetSeqStart() int32  { return m.Data2.SeqStart }

func (m *TLUpdatesCombined) SetSeq(v int32) { m.Data2.Seq = v }
func (m *TLUpdatesCombined) GetSeq() int32  { return m.Data2.Seq }

func (m *TLUpdates) SetUpdates(v []*Update) { m.Data2.Updates = v }
func (m *TLUpdates) GetUpdates() []*Update  { return m.Data2.Updates }

func (m *TLUpdates) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdates) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdates) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdates) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdates) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdates) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdates) SetSeq(v int32) { m.Data2.Seq = v }
func (m *TLUpdates) GetSeq() int32  { return m.Data2.Seq }

func (m *TLUpdateShortSentMessage) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdateShortSentMessage) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdateShortSentMessage) SetOut(v bool) { m.Data2.Out = v }
func (m *TLUpdateShortSentMessage) GetOut() bool  { return m.Data2.Out }

func (m *TLUpdateShortSentMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLUpdateShortSentMessage) GetId() int32  { return m.Data2.Id }

func (m *TLUpdateShortSentMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateShortSentMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateShortSentMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateShortSentMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateShortSentMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateShortSentMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateShortSentMessage) SetMedia(v *MessageMedia) { m.Data2.Media = v }
func (m *TLUpdateShortSentMessage) GetMedia() *MessageMedia  { return m.Data2.Media }

func (m *TLUpdateShortSentMessage) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLUpdateShortSentMessage) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLUpdatesChannelDifferenceEmpty) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdatesChannelDifferenceEmpty) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdatesChannelDifferenceEmpty) SetFinal(v bool) { m.Data2.Final = v }
func (m *TLUpdatesChannelDifferenceEmpty) GetFinal() bool  { return m.Data2.Final }

func (m *TLUpdatesChannelDifferenceEmpty) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatesChannelDifferenceEmpty) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatesChannelDifferenceEmpty) SetTimeout(v int32) { m.Data2.Timeout = v }
func (m *TLUpdatesChannelDifferenceEmpty) GetTimeout() int32  { return m.Data2.Timeout }

func (m *TLUpdatesChannelDifferenceTooLong) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdatesChannelDifferenceTooLong) SetFinal(v bool) { m.Data2.Final = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetFinal() bool  { return m.Data2.Final }

func (m *TLUpdatesChannelDifferenceTooLong) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatesChannelDifferenceTooLong) SetTimeout(v int32) { m.Data2.Timeout = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetTimeout() int32  { return m.Data2.Timeout }

func (m *TLUpdatesChannelDifferenceTooLong) SetTopMessage(v int32) { m.Data2.TopMessage = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetTopMessage() int32  { return m.Data2.TopMessage }

func (m *TLUpdatesChannelDifferenceTooLong) SetReadInboxMaxId(v int32) { m.Data2.ReadInboxMaxId = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetReadInboxMaxId() int32  { return m.Data2.ReadInboxMaxId }

func (m *TLUpdatesChannelDifferenceTooLong) SetReadOutboxMaxId(v int32) { m.Data2.ReadOutboxMaxId = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetReadOutboxMaxId() int32 {
	return m.Data2.ReadOutboxMaxId
}

func (m *TLUpdatesChannelDifferenceTooLong) SetUnreadCount(v int32) { m.Data2.UnreadCount = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetUnreadCount() int32  { return m.Data2.UnreadCount }

func (m *TLUpdatesChannelDifferenceTooLong) SetUnreadMentionsCount(v int32) {
	m.Data2.UnreadMentionsCount = v
}
func (m *TLUpdatesChannelDifferenceTooLong) GetUnreadMentionsCount() int32 {
	return m.Data2.UnreadMentionsCount
}

func (m *TLUpdatesChannelDifferenceTooLong) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLUpdatesChannelDifferenceTooLong) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdatesChannelDifferenceTooLong) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdatesChannelDifferenceTooLong) SetTopImportantMessage(v int32) {
	m.Data2.TopImportantMessage = v
}
func (m *TLUpdatesChannelDifferenceTooLong) GetTopImportantMessage() int32 {
	return m.Data2.TopImportantMessage
}

func (m *TLUpdatesChannelDifferenceTooLong) SetUnreadImportantCount(v int32) {
	m.Data2.UnreadImportantCount = v
}
func (m *TLUpdatesChannelDifferenceTooLong) GetUnreadImportantCount() int32 {
	return m.Data2.UnreadImportantCount
}

func (m *TLUpdatesChannelDifferenceTooLong) SetDialog(v *Dialog) { m.Data2.Dialog = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetDialog() *Dialog  { return m.Data2.Dialog }

func (m *TLUpdatesChannelDifference) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUpdatesChannelDifference) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUpdatesChannelDifference) SetFinal(v bool) { m.Data2.Final = v }
func (m *TLUpdatesChannelDifference) GetFinal() bool  { return m.Data2.Final }

func (m *TLUpdatesChannelDifference) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatesChannelDifference) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatesChannelDifference) SetTimeout(v int32) { m.Data2.Timeout = v }
func (m *TLUpdatesChannelDifference) GetTimeout() int32  { return m.Data2.Timeout }

func (m *TLUpdatesChannelDifference) SetNewMessages(v []*Message) { m.Data2.NewMessages = v }
func (m *TLUpdatesChannelDifference) GetNewMessages() []*Message  { return m.Data2.NewMessages }

func (m *TLUpdatesChannelDifference) SetOtherUpdates(v []*Update) { m.Data2.OtherUpdates = v }
func (m *TLUpdatesChannelDifference) GetOtherUpdates() []*Update  { return m.Data2.OtherUpdates }

func (m *TLUpdatesChannelDifference) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdatesChannelDifference) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdatesChannelDifference) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdatesChannelDifference) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdatesDifferenceEmpty) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdatesDifferenceEmpty) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdatesDifferenceEmpty) SetSeq(v int32) { m.Data2.Seq = v }
func (m *TLUpdatesDifferenceEmpty) GetSeq() int32  { return m.Data2.Seq }

func (m *TLUpdatesDifference) SetNewMessages(v []*Message) { m.Data2.NewMessages = v }
func (m *TLUpdatesDifference) GetNewMessages() []*Message  { return m.Data2.NewMessages }

func (m *TLUpdatesDifference) SetNewEncryptedMessages(v []*EncryptedMessage) {
	m.Data2.NewEncryptedMessages = v
}
func (m *TLUpdatesDifference) GetNewEncryptedMessages() []*EncryptedMessage {
	return m.Data2.NewEncryptedMessages
}

func (m *TLUpdatesDifference) SetOtherUpdates(v []*Update) { m.Data2.OtherUpdates = v }
func (m *TLUpdatesDifference) GetOtherUpdates() []*Update  { return m.Data2.OtherUpdates }

func (m *TLUpdatesDifference) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdatesDifference) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdatesDifference) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdatesDifference) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdatesDifference) SetState(v *Updates_State) { m.Data2.State = v }
func (m *TLUpdatesDifference) GetState() *Updates_State  { return m.Data2.State }

func (m *TLUpdatesDifferenceSlice) SetNewMessages(v []*Message) { m.Data2.NewMessages = v }
func (m *TLUpdatesDifferenceSlice) GetNewMessages() []*Message  { return m.Data2.NewMessages }

func (m *TLUpdatesDifferenceSlice) SetNewEncryptedMessages(v []*EncryptedMessage) {
	m.Data2.NewEncryptedMessages = v
}
func (m *TLUpdatesDifferenceSlice) GetNewEncryptedMessages() []*EncryptedMessage {
	return m.Data2.NewEncryptedMessages
}

func (m *TLUpdatesDifferenceSlice) SetOtherUpdates(v []*Update) { m.Data2.OtherUpdates = v }
func (m *TLUpdatesDifferenceSlice) GetOtherUpdates() []*Update  { return m.Data2.OtherUpdates }

func (m *TLUpdatesDifferenceSlice) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdatesDifferenceSlice) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdatesDifferenceSlice) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdatesDifferenceSlice) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdatesDifferenceSlice) SetIntermediateState(v *Updates_State) {
	m.Data2.IntermediateState = v
}
func (m *TLUpdatesDifferenceSlice) GetIntermediateState() *Updates_State {
	return m.Data2.IntermediateState
}

func (m *TLUpdatesDifferenceTooLong) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatesDifferenceTooLong) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatesState) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatesState) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatesState) SetQts(v int32) { m.Data2.Qts = v }
func (m *TLUpdatesState) GetQts() int32  { return m.Data2.Qts }

func (m *TLUpdatesState) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdatesState) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdatesState) SetSeq(v int32) { m.Data2.Seq = v }
func (m *TLUpdatesState) GetSeq() int32  { return m.Data2.Seq }

func (m *TLUpdatesState) SetUnreadCount(v int32) { m.Data2.UnreadCount = v }
func (m *TLUpdatesState) GetUnreadCount() int32  { return m.Data2.UnreadCount }

func (m *TLUploadCdnFileReuploadNeeded) SetRequestToken(v []byte) { m.Data2.RequestToken = v }
func (m *TLUploadCdnFileReuploadNeeded) GetRequestToken() []byte  { return m.Data2.RequestToken }

func (m *TLUploadCdnFile) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLUploadCdnFile) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLUploadFile) SetType(v *Storage_FileType) { m.Data2.Type = v }
func (m *TLUploadFile) GetType() *Storage_FileType  { return m.Data2.Type }

func (m *TLUploadFile) SetMtime(v int32) { m.Data2.Mtime = v }
func (m *TLUploadFile) GetMtime() int32  { return m.Data2.Mtime }

func (m *TLUploadFile) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLUploadFile) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLUploadFileCdnRedirect) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLUploadFileCdnRedirect) GetDcId() int32  { return m.Data2.DcId }

func (m *TLUploadFileCdnRedirect) SetFileToken(v []byte) { m.Data2.FileToken = v }
func (m *TLUploadFileCdnRedirect) GetFileToken() []byte  { return m.Data2.FileToken }

func (m *TLUploadFileCdnRedirect) SetEncryptionKey(v []byte) { m.Data2.EncryptionKey = v }
func (m *TLUploadFileCdnRedirect) GetEncryptionKey() []byte  { return m.Data2.EncryptionKey }

func (m *TLUploadFileCdnRedirect) SetEncryptionIv(v []byte) { m.Data2.EncryptionIv = v }
func (m *TLUploadFileCdnRedirect) GetEncryptionIv() []byte  { return m.Data2.EncryptionIv }

func (m *TLUploadFileCdnRedirect) SetCdnFileHashes(v []*CdnFileHash) { m.Data2.CdnFileHashes = v }
func (m *TLUploadFileCdnRedirect) GetCdnFileHashes() []*CdnFileHash  { return m.Data2.CdnFileHashes }

func (m *TLUploadFileCdnRedirect) SetFileHashes(v []*FileHash) { m.Data2.FileHashes = v }
func (m *TLUploadFileCdnRedirect) GetFileHashes() []*FileHash  { return m.Data2.FileHashes }

func (m *TLUploadWebFile) SetSize_(v int32) { m.Data2.Size_ = v }
func (m *TLUploadWebFile) GetSize_() int32  { return m.Data2.Size_ }

func (m *TLUploadWebFile) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLUploadWebFile) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLUploadWebFile) SetFileType(v *Storage_FileType) { m.Data2.FileType = v }
func (m *TLUploadWebFile) GetFileType() *Storage_FileType  { return m.Data2.FileType }

func (m *TLUploadWebFile) SetMtime(v int32) { m.Data2.Mtime = v }
func (m *TLUploadWebFile) GetMtime() int32  { return m.Data2.Mtime }

func (m *TLUploadWebFile) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLUploadWebFile) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLUrlAuthResultRequest) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUrlAuthResultRequest) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUrlAuthResultRequest) SetRequestWriteAccess(v bool) { m.Data2.RequestWriteAccess = v }
func (m *TLUrlAuthResultRequest) GetRequestWriteAccess() bool  { return m.Data2.RequestWriteAccess }

func (m *TLUrlAuthResultRequest) SetBot(v *User) { m.Data2.Bot = v }
func (m *TLUrlAuthResultRequest) GetBot() *User  { return m.Data2.Bot }

func (m *TLUrlAuthResultRequest) SetDomain(v string) { m.Data2.Domain = v }
func (m *TLUrlAuthResultRequest) GetDomain() string  { return m.Data2.Domain }

func (m *TLUrlAuthResultAccepted) SetUrl(v string) { m.Data2.Url = v }
func (m *TLUrlAuthResultAccepted) GetUrl() string  { return m.Data2.Url }

func (m *TLUserEmpty) SetId(v int32) { m.Data2.Id = v }
func (m *TLUserEmpty) GetId() int32  { return m.Data2.Id }

func (m *TLUser) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUser) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUser) SetSelf(v bool) { m.Data2.Self = v }
func (m *TLUser) GetSelf() bool  { return m.Data2.Self }

func (m *TLUser) SetContact(v bool) { m.Data2.Contact = v }
func (m *TLUser) GetContact() bool  { return m.Data2.Contact }

func (m *TLUser) SetMutualContact(v bool) { m.Data2.MutualContact = v }
func (m *TLUser) GetMutualContact() bool  { return m.Data2.MutualContact }

func (m *TLUser) SetDeleted(v bool) { m.Data2.Deleted = v }
func (m *TLUser) GetDeleted() bool  { return m.Data2.Deleted }

func (m *TLUser) SetBot(v bool) { m.Data2.Bot = v }
func (m *TLUser) GetBot() bool  { return m.Data2.Bot }

func (m *TLUser) SetBotChatHistory(v bool) { m.Data2.BotChatHistory = v }
func (m *TLUser) GetBotChatHistory() bool  { return m.Data2.BotChatHistory }

func (m *TLUser) SetBotNochats(v bool) { m.Data2.BotNochats = v }
func (m *TLUser) GetBotNochats() bool  { return m.Data2.BotNochats }

func (m *TLUser) SetVerified(v bool) { m.Data2.Verified = v }
func (m *TLUser) GetVerified() bool  { return m.Data2.Verified }

func (m *TLUser) SetRestricted(v bool) { m.Data2.Restricted = v }
func (m *TLUser) GetRestricted() bool  { return m.Data2.Restricted }

func (m *TLUser) SetMin(v bool) { m.Data2.Min = v }
func (m *TLUser) GetMin() bool  { return m.Data2.Min }

func (m *TLUser) SetBotInlineGeo(v bool) { m.Data2.BotInlineGeo = v }
func (m *TLUser) GetBotInlineGeo() bool  { return m.Data2.BotInlineGeo }

func (m *TLUser) SetId(v int32) { m.Data2.Id = v }
func (m *TLUser) GetId() int32  { return m.Data2.Id }

func (m *TLUser) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLUser) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLUser) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLUser) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLUser) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLUser) GetLastName() string  { return m.Data2.LastName }

func (m *TLUser) SetUsername(v string) { m.Data2.Username = v }
func (m *TLUser) GetUsername() string  { return m.Data2.Username }

func (m *TLUser) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLUser) GetPhone() string  { return m.Data2.Phone }

func (m *TLUser) SetPhoto(v *UserProfilePhoto) { m.Data2.Photo = v }
func (m *TLUser) GetPhoto() *UserProfilePhoto  { return m.Data2.Photo }

func (m *TLUser) SetStatus(v *UserStatus) { m.Data2.Status = v }
func (m *TLUser) GetStatus() *UserStatus  { return m.Data2.Status }

func (m *TLUser) SetBotInfoVersion(v int32) { m.Data2.BotInfoVersion = v }
func (m *TLUser) GetBotInfoVersion() int32  { return m.Data2.BotInfoVersion }

func (m *TLUser) SetRestrictionReason2E13F4C371(v string) { m.Data2.RestrictionReason2E13F4C371 = v }
func (m *TLUser) GetRestrictionReason2E13F4C371() string  { return m.Data2.RestrictionReason2E13F4C371 }

func (m *TLUser) SetBotInlinePlaceholder(v string) { m.Data2.BotInlinePlaceholder = v }
func (m *TLUser) GetBotInlinePlaceholder() string  { return m.Data2.BotInlinePlaceholder }

func (m *TLUser) SetLangCode(v string) { m.Data2.LangCode = v }
func (m *TLUser) GetLangCode() string  { return m.Data2.LangCode }

func (m *TLUser) SetSupport(v bool) { m.Data2.Support = v }
func (m *TLUser) GetSupport() bool  { return m.Data2.Support }

func (m *TLUser) SetScam(v bool) { m.Data2.Scam = v }
func (m *TLUser) GetScam() bool  { return m.Data2.Scam }

func (m *TLUser) SetApplyMinPhoto(v bool) { m.Data2.ApplyMinPhoto = v }
func (m *TLUser) GetApplyMinPhoto() bool  { return m.Data2.ApplyMinPhoto }

func (m *TLUser) SetRestrictionReason938458C1117(v []*RestrictionReason) {
	m.Data2.RestrictionReason938458C1117 = v
}
func (m *TLUser) GetRestrictionReason938458C1117() []*RestrictionReason {
	return m.Data2.RestrictionReason938458C1117
}

func (m *TLUserFull) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUserFull) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUserFull) SetBlocked(v bool) { m.Data2.Blocked = v }
func (m *TLUserFull) GetBlocked() bool  { return m.Data2.Blocked }

func (m *TLUserFull) SetPhoneCallsAvailable(v bool) { m.Data2.PhoneCallsAvailable = v }
func (m *TLUserFull) GetPhoneCallsAvailable() bool  { return m.Data2.PhoneCallsAvailable }

func (m *TLUserFull) SetPhoneCallsPrivate(v bool) { m.Data2.PhoneCallsPrivate = v }
func (m *TLUserFull) GetPhoneCallsPrivate() bool  { return m.Data2.PhoneCallsPrivate }

func (m *TLUserFull) SetUser(v *User) { m.Data2.User = v }
func (m *TLUserFull) GetUser() *User  { return m.Data2.User }

func (m *TLUserFull) SetAbout(v string) { m.Data2.About = v }
func (m *TLUserFull) GetAbout() string  { return m.Data2.About }

func (m *TLUserFull) SetLink(v *Contacts_Link) { m.Data2.Link = v }
func (m *TLUserFull) GetLink() *Contacts_Link  { return m.Data2.Link }

func (m *TLUserFull) SetProfilePhoto(v *Photo) { m.Data2.ProfilePhoto = v }
func (m *TLUserFull) GetProfilePhoto() *Photo  { return m.Data2.ProfilePhoto }

func (m *TLUserFull) SetNotifySettings(v *PeerNotifySettings) { m.Data2.NotifySettings = v }
func (m *TLUserFull) GetNotifySettings() *PeerNotifySettings  { return m.Data2.NotifySettings }

func (m *TLUserFull) SetBotInfo(v *BotInfo) { m.Data2.BotInfo = v }
func (m *TLUserFull) GetBotInfo() *BotInfo  { return m.Data2.BotInfo }

func (m *TLUserFull) SetCommonChatsCount(v int32) { m.Data2.CommonChatsCount = v }
func (m *TLUserFull) GetCommonChatsCount() int32  { return m.Data2.CommonChatsCount }

func (m *TLUserFull) SetCanPinMessage(v bool) { m.Data2.CanPinMessage = v }
func (m *TLUserFull) GetCanPinMessage() bool  { return m.Data2.CanPinMessage }

func (m *TLUserFull) SetPinnedMsgId(v int32) { m.Data2.PinnedMsgId = v }
func (m *TLUserFull) GetPinnedMsgId() int32  { return m.Data2.PinnedMsgId }

func (m *TLUserFull) SetFolderId(v int32) { m.Data2.FolderId = v }
func (m *TLUserFull) GetFolderId() int32  { return m.Data2.FolderId }

func (m *TLUserFull) SetHasScheduled(v bool) { m.Data2.HasScheduled = v }
func (m *TLUserFull) GetHasScheduled() bool  { return m.Data2.HasScheduled }

func (m *TLUserFull) SetVideoCallsAvailable(v bool) { m.Data2.VideoCallsAvailable = v }
func (m *TLUserFull) GetVideoCallsAvailable() bool  { return m.Data2.VideoCallsAvailable }

func (m *TLUserFull) SetSettings(v *PeerSettings) { m.Data2.Settings = v }
func (m *TLUserFull) GetSettings() *PeerSettings  { return m.Data2.Settings }

func (m *TLUserProfilePhoto) SetPhotoId(v int64) { m.Data2.PhotoId = v }
func (m *TLUserProfilePhoto) GetPhotoId() int64  { return m.Data2.PhotoId }

func (m *TLUserProfilePhoto) SetPhotoSmall(v *FileLocation) { m.Data2.PhotoSmall = v }
func (m *TLUserProfilePhoto) GetPhotoSmall() *FileLocation  { return m.Data2.PhotoSmall }

func (m *TLUserProfilePhoto) SetPhotoBig(v *FileLocation) { m.Data2.PhotoBig = v }
func (m *TLUserProfilePhoto) GetPhotoBig() *FileLocation  { return m.Data2.PhotoBig }

func (m *TLUserProfilePhoto) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLUserProfilePhoto) GetDcId() int32  { return m.Data2.DcId }

func (m *TLUserProfilePhoto) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLUserProfilePhoto) GetFlags() int32  { return m.Data2.Flags }

func (m *TLUserProfilePhoto) SetHasVideo(v bool) { m.Data2.HasVideo = v }
func (m *TLUserProfilePhoto) GetHasVideo() bool  { return m.Data2.HasVideo }

func (m *TLUserStatusOnline) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLUserStatusOnline) GetExpires() int32  { return m.Data2.Expires }

func (m *TLUserStatusOffline) SetWasOnline(v int32) { m.Data2.WasOnline = v }
func (m *TLUserStatusOffline) GetWasOnline() int32  { return m.Data2.WasOnline }

func (m *TLVideoSize) SetType(v string) { m.Data2.Type = v }
func (m *TLVideoSize) GetType() string  { return m.Data2.Type }

func (m *TLVideoSize) SetLocation(v *FileLocation) { m.Data2.Location = v }
func (m *TLVideoSize) GetLocation() *FileLocation  { return m.Data2.Location }

func (m *TLVideoSize) SetW(v int32) { m.Data2.W = v }
func (m *TLVideoSize) GetW() int32  { return m.Data2.W }

func (m *TLVideoSize) SetH(v int32) { m.Data2.H = v }
func (m *TLVideoSize) GetH() int32  { return m.Data2.H }

func (m *TLVideoSize) SetSize_(v int32) { m.Data2.Size_ = v }
func (m *TLVideoSize) GetSize_() int32  { return m.Data2.Size_ }

func (m *TLVideoSize) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLVideoSize) GetFlags() int32  { return m.Data2.Flags }

func (m *TLVideoSize) SetVideoStartTs(v float64) { m.Data2.VideoStartTs = v }
func (m *TLVideoSize) GetVideoStartTs() float64  { return m.Data2.VideoStartTs }

func (m *TLWallPaper) SetIdCCB0365771(v int32) { m.Data2.IdCCB0365771 = v }
func (m *TLWallPaper) GetIdCCB0365771() int32  { return m.Data2.IdCCB0365771 }

func (m *TLWallPaper) SetTitle(v string) { m.Data2.Title = v }
func (m *TLWallPaper) GetTitle() string  { return m.Data2.Title }

func (m *TLWallPaper) SetSizes(v []*PhotoSize) { m.Data2.Sizes = v }
func (m *TLWallPaper) GetSizes() []*PhotoSize  { return m.Data2.Sizes }

func (m *TLWallPaper) SetColor(v int32) { m.Data2.Color = v }
func (m *TLWallPaper) GetColor() int32  { return m.Data2.Color }

func (m *TLWallPaper) SetIdF04F91EC93(v int64) { m.Data2.IdF04F91EC93 = v }
func (m *TLWallPaper) GetIdF04F91EC93() int64  { return m.Data2.IdF04F91EC93 }

func (m *TLWallPaper) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLWallPaper) GetFlags() int32  { return m.Data2.Flags }

func (m *TLWallPaper) SetCreator(v bool) { m.Data2.Creator = v }
func (m *TLWallPaper) GetCreator() bool  { return m.Data2.Creator }

func (m *TLWallPaper) SetDefault(v bool) { m.Data2.Default = v }
func (m *TLWallPaper) GetDefault() bool  { return m.Data2.Default }

func (m *TLWallPaper) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLWallPaper) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLWallPaper) SetSlug(v string) { m.Data2.Slug = v }
func (m *TLWallPaper) GetSlug() string  { return m.Data2.Slug }

func (m *TLWallPaper) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLWallPaper) GetDocument() *Document  { return m.Data2.Document }

func (m *TLWallPaper) SetPattern(v bool) { m.Data2.Pattern = v }
func (m *TLWallPaper) GetPattern() bool  { return m.Data2.Pattern }

func (m *TLWallPaper) SetDark(v bool) { m.Data2.Dark = v }
func (m *TLWallPaper) GetDark() bool  { return m.Data2.Dark }

func (m *TLWallPaper) SetSettings(v *WallPaperSettings) { m.Data2.Settings = v }
func (m *TLWallPaper) GetSettings() *WallPaperSettings  { return m.Data2.Settings }

func (m *TLWallPaperSolid) SetIdCCB0365771(v int32) { m.Data2.IdCCB0365771 = v }
func (m *TLWallPaperSolid) GetIdCCB0365771() int32  { return m.Data2.IdCCB0365771 }

func (m *TLWallPaperSolid) SetTitle(v string) { m.Data2.Title = v }
func (m *TLWallPaperSolid) GetTitle() string  { return m.Data2.Title }

func (m *TLWallPaperSolid) SetBgColor(v int32) { m.Data2.BgColor = v }
func (m *TLWallPaperSolid) GetBgColor() int32  { return m.Data2.BgColor }

func (m *TLWallPaperSolid) SetColor(v int32) { m.Data2.Color = v }
func (m *TLWallPaperSolid) GetColor() int32  { return m.Data2.Color }

func (m *TLWallPaperNoFile) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLWallPaperNoFile) GetFlags() int32  { return m.Data2.Flags }

func (m *TLWallPaperNoFile) SetDefault(v bool) { m.Data2.Default = v }
func (m *TLWallPaperNoFile) GetDefault() bool  { return m.Data2.Default }

func (m *TLWallPaperNoFile) SetDark(v bool) { m.Data2.Dark = v }
func (m *TLWallPaperNoFile) GetDark() bool  { return m.Data2.Dark }

func (m *TLWallPaperNoFile) SetSettings(v *WallPaperSettings) { m.Data2.Settings = v }
func (m *TLWallPaperNoFile) GetSettings() *WallPaperSettings  { return m.Data2.Settings }

func (m *TLWallPaperSettings) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLWallPaperSettings) GetFlags() int32  { return m.Data2.Flags }

func (m *TLWallPaperSettings) SetBlur(v bool) { m.Data2.Blur = v }
func (m *TLWallPaperSettings) GetBlur() bool  { return m.Data2.Blur }

func (m *TLWallPaperSettings) SetMotion(v bool) { m.Data2.Motion = v }
func (m *TLWallPaperSettings) GetMotion() bool  { return m.Data2.Motion }

func (m *TLWallPaperSettings) SetBackgroundColor(v int32) { m.Data2.BackgroundColor = v }
func (m *TLWallPaperSettings) GetBackgroundColor() int32  { return m.Data2.BackgroundColor }

func (m *TLWallPaperSettings) SetIntensity(v int32) { m.Data2.Intensity = v }
func (m *TLWallPaperSettings) GetIntensity() int32  { return m.Data2.Intensity }

func (m *TLWallPaperSettings) SetSecondBackgroundColor(v int32) { m.Data2.SecondBackgroundColor = v }
func (m *TLWallPaperSettings) GetSecondBackgroundColor() int32  { return m.Data2.SecondBackgroundColor }

func (m *TLWallPaperSettings) SetRotation(v int32) { m.Data2.Rotation = v }
func (m *TLWallPaperSettings) GetRotation() int32  { return m.Data2.Rotation }

func (m *TLWalletSecretSalt) SetSalt(v []byte) { m.Data2.Salt = v }
func (m *TLWalletSecretSalt) GetSalt() []byte  { return m.Data2.Salt }

func (m *TLWalletLiteResponse) SetResponse(v []byte) { m.Data2.Response = v }
func (m *TLWalletLiteResponse) GetResponse() []byte  { return m.Data2.Response }

func (m *TLWebAuthorization) SetHash(v int64) { m.Data2.Hash = v }
func (m *TLWebAuthorization) GetHash() int64  { return m.Data2.Hash }

func (m *TLWebAuthorization) SetBotId(v int32) { m.Data2.BotId = v }
func (m *TLWebAuthorization) GetBotId() int32  { return m.Data2.BotId }

func (m *TLWebAuthorization) SetDomain(v string) { m.Data2.Domain = v }
func (m *TLWebAuthorization) GetDomain() string  { return m.Data2.Domain }

func (m *TLWebAuthorization) SetBrowser(v string) { m.Data2.Browser = v }
func (m *TLWebAuthorization) GetBrowser() string  { return m.Data2.Browser }

func (m *TLWebAuthorization) SetPlatform(v string) { m.Data2.Platform = v }
func (m *TLWebAuthorization) GetPlatform() string  { return m.Data2.Platform }

func (m *TLWebAuthorization) SetDateCreated(v int32) { m.Data2.DateCreated = v }
func (m *TLWebAuthorization) GetDateCreated() int32  { return m.Data2.DateCreated }

func (m *TLWebAuthorization) SetDateActive(v int32) { m.Data2.DateActive = v }
func (m *TLWebAuthorization) GetDateActive() int32  { return m.Data2.DateActive }

func (m *TLWebAuthorization) SetIp(v string) { m.Data2.Ip = v }
func (m *TLWebAuthorization) GetIp() string  { return m.Data2.Ip }

func (m *TLWebAuthorization) SetRegion(v string) { m.Data2.Region = v }
func (m *TLWebAuthorization) GetRegion() string  { return m.Data2.Region }

func (m *TLWebDocument) SetUrl(v string) { m.Data2.Url = v }
func (m *TLWebDocument) GetUrl() string  { return m.Data2.Url }

func (m *TLWebDocument) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLWebDocument) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLWebDocument) SetSize_(v int32) { m.Data2.Size_ = v }
func (m *TLWebDocument) GetSize_() int32  { return m.Data2.Size_ }

func (m *TLWebDocument) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLWebDocument) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLWebDocument) SetAttributes(v []*DocumentAttribute) { m.Data2.Attributes = v }
func (m *TLWebDocument) GetAttributes() []*DocumentAttribute  { return m.Data2.Attributes }

func (m *TLWebDocument) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLWebDocument) GetDcId() int32  { return m.Data2.DcId }

func (m *TLWebDocumentNoProxy) SetUrl(v string) { m.Data2.Url = v }
func (m *TLWebDocumentNoProxy) GetUrl() string  { return m.Data2.Url }

func (m *TLWebDocumentNoProxy) SetSize_(v int32) { m.Data2.Size_ = v }
func (m *TLWebDocumentNoProxy) GetSize_() int32  { return m.Data2.Size_ }

func (m *TLWebDocumentNoProxy) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLWebDocumentNoProxy) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLWebDocumentNoProxy) SetAttributes(v []*DocumentAttribute) { m.Data2.Attributes = v }
func (m *TLWebDocumentNoProxy) GetAttributes() []*DocumentAttribute  { return m.Data2.Attributes }

func (m *TLWebPageEmpty) SetId(v int64) { m.Data2.Id = v }
func (m *TLWebPageEmpty) GetId() int64  { return m.Data2.Id }

func (m *TLWebPagePending) SetId(v int64) { m.Data2.Id = v }
func (m *TLWebPagePending) GetId() int64  { return m.Data2.Id }

func (m *TLWebPagePending) SetDate(v int32) { m.Data2.Date = v }
func (m *TLWebPagePending) GetDate() int32  { return m.Data2.Date }

func (m *TLWebPage) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLWebPage) GetFlags() int32  { return m.Data2.Flags }

func (m *TLWebPage) SetId(v int64) { m.Data2.Id = v }
func (m *TLWebPage) GetId() int64  { return m.Data2.Id }

func (m *TLWebPage) SetUrl(v string) { m.Data2.Url = v }
func (m *TLWebPage) GetUrl() string  { return m.Data2.Url }

func (m *TLWebPage) SetDisplayUrl(v string) { m.Data2.DisplayUrl = v }
func (m *TLWebPage) GetDisplayUrl() string  { return m.Data2.DisplayUrl }

func (m *TLWebPage) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLWebPage) GetHash() int32  { return m.Data2.Hash }

func (m *TLWebPage) SetType(v string) { m.Data2.Type = v }
func (m *TLWebPage) GetType() string  { return m.Data2.Type }

func (m *TLWebPage) SetSiteName(v string) { m.Data2.SiteName = v }
func (m *TLWebPage) GetSiteName() string  { return m.Data2.SiteName }

func (m *TLWebPage) SetTitle(v string) { m.Data2.Title = v }
func (m *TLWebPage) GetTitle() string  { return m.Data2.Title }

func (m *TLWebPage) SetDescription(v string) { m.Data2.Description = v }
func (m *TLWebPage) GetDescription() string  { return m.Data2.Description }

func (m *TLWebPage) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLWebPage) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLWebPage) SetEmbedUrl(v string) { m.Data2.EmbedUrl = v }
func (m *TLWebPage) GetEmbedUrl() string  { return m.Data2.EmbedUrl }

func (m *TLWebPage) SetEmbedType(v string) { m.Data2.EmbedType = v }
func (m *TLWebPage) GetEmbedType() string  { return m.Data2.EmbedType }

func (m *TLWebPage) SetEmbedWidth(v int32) { m.Data2.EmbedWidth = v }
func (m *TLWebPage) GetEmbedWidth() int32  { return m.Data2.EmbedWidth }

func (m *TLWebPage) SetEmbedHeight(v int32) { m.Data2.EmbedHeight = v }
func (m *TLWebPage) GetEmbedHeight() int32  { return m.Data2.EmbedHeight }

func (m *TLWebPage) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLWebPage) GetDuration() int32  { return m.Data2.Duration }

func (m *TLWebPage) SetAuthor(v string) { m.Data2.Author = v }
func (m *TLWebPage) GetAuthor() string  { return m.Data2.Author }

func (m *TLWebPage) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLWebPage) GetDocument() *Document  { return m.Data2.Document }

func (m *TLWebPage) SetCachedPage(v *Page) { m.Data2.CachedPage = v }
func (m *TLWebPage) GetCachedPage() *Page  { return m.Data2.CachedPage }

func (m *TLWebPage) SetDocuments(v []*Document) { m.Data2.Documents = v }
func (m *TLWebPage) GetDocuments() []*Document  { return m.Data2.Documents }

func (m *TLWebPage) SetAttributes(v []*WebPageAttribute) { m.Data2.Attributes = v }
func (m *TLWebPage) GetAttributes() []*WebPageAttribute  { return m.Data2.Attributes }

func (m *TLWebPageNotModified) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLWebPageNotModified) GetFlags() int32  { return m.Data2.Flags }

func (m *TLWebPageNotModified) SetCachedPageViews(v int32) { m.Data2.CachedPageViews = v }
func (m *TLWebPageNotModified) GetCachedPageViews() int32  { return m.Data2.CachedPageViews }

func (m *TLWebPageAttributeTheme) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLWebPageAttributeTheme) GetFlags() int32  { return m.Data2.Flags }

func (m *TLWebPageAttributeTheme) SetDocuments(v []*Document) { m.Data2.Documents = v }
func (m *TLWebPageAttributeTheme) GetDocuments() []*Document  { return m.Data2.Documents }

func (m *TLWebPageAttributeTheme) SetSettings(v *ThemeSettings) { m.Data2.Settings = v }
func (m *TLWebPageAttributeTheme) GetSettings() *ThemeSettings  { return m.Data2.Settings }
