/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2023/06/12 18:03
 * @File : codec_schema_member_to.pb.go
 * @Desc :
 *
 */

package mtproto

//  + TL_AccessPointRule
//

//  accessPointRule#4679b65f phone_prefix_rules:string dc_id:int ips:vector<IpPort> = AccessPointRule;
//
func (m *AccessPointRule) To_AccessPointRule() *TLAccessPointRule {
	return &TLAccessPointRule{
		Data2: m,
	}
}

//  accessPointRule#4679b65f phone_prefix_rules:string dc_id:int ips:vector<IpPort> = AccessPointRule;
//
func (m *TLAccessPointRule) To_AccessPointRule() *AccessPointRule {
	m.Data2.Cmd = Cmd_AccessPointRule
	return m.Data2
}

//  + TL_AccountDaysTTL
//

//  accountDaysTTL#b8d0afdf days:int = AccountDaysTTL;
//
func (m *AccountDaysTTL) To_AccountDaysTTL() *TLAccountDaysTTL {
	return &TLAccountDaysTTL{
		Data2: m,
	}
}

//  accountDaysTTL#b8d0afdf days:int = AccountDaysTTL;
//
func (m *TLAccountDaysTTL) To_AccountDaysTTL() *AccountDaysTTL {
	m.Data2.Cmd = Cmd_AccountDaysTTL
	return m.Data2
}

//  + TL_AccountAuthorizationForm
//

//  account.authorizationForm#ad2e1cd8 flags:# required_types:Vector<SecureRequiredType> values:Vector<SecureValue> errors:Vector<SecureValueError> users:Vector<User> privacy_policy_url:flags.0?string = account.AuthorizationForm;
//
func (m *Account_AuthorizationForm) To_AccountAuthorizationForm() *TLAccountAuthorizationForm {
	return &TLAccountAuthorizationForm{
		Data2: m,
	}
}

//  account.authorizationForm#ad2e1cd8 flags:# required_types:Vector<SecureRequiredType> values:Vector<SecureValue> errors:Vector<SecureValueError> users:Vector<User> privacy_policy_url:flags.0?string = account.AuthorizationForm;
//
func (m *TLAccountAuthorizationForm) To_Account_AuthorizationForm() *Account_AuthorizationForm {
	m.Data2.Cmd = Cmd_AccountAuthorizationForm
	return m.Data2
}

//  + TL_AccountAuthorizations
//

//  account.authorizations#1250abde authorizations:Vector<Authorization> = account.Authorizations;
//
func (m *Account_Authorizations) To_AccountAuthorizations() *TLAccountAuthorizations {
	return &TLAccountAuthorizations{
		Data2: m,
	}
}

//  account.authorizations#1250abde authorizations:Vector<Authorization> = account.Authorizations;
//
func (m *TLAccountAuthorizations) To_Account_Authorizations() *Account_Authorizations {
	m.Data2.Cmd = Cmd_AccountAuthorizations
	return m.Data2
}

//  + TL_AccountAutoDownloadSettings
//

//  account.autoDownloadSettings#63cacf26 low:AutoDownloadSettings medium:AutoDownloadSettings high:AutoDownloadSettings = account.AutoDownloadSettings;
//
func (m *Account_AutoDownloadSettings) To_AccountAutoDownloadSettings() *TLAccountAutoDownloadSettings {
	return &TLAccountAutoDownloadSettings{
		Data2: m,
	}
}

//  account.autoDownloadSettings#63cacf26 low:AutoDownloadSettings medium:AutoDownloadSettings high:AutoDownloadSettings = account.AutoDownloadSettings;
//
func (m *TLAccountAutoDownloadSettings) To_Account_AutoDownloadSettings() *Account_AutoDownloadSettings {
	m.Data2.Cmd = Cmd_AccountAutoDownloadSettings
	return m.Data2
}

//  + TL_AccountContentSettings
//

//  account.contentSettings#57e28221 flags:# sensitive_enabled:flags.0?true sensitive_can_change:flags.1?true = account.ContentSettings;
//
func (m *Account_ContentSettings) To_AccountContentSettings() *TLAccountContentSettings {
	return &TLAccountContentSettings{
		Data2: m,
	}
}

//  account.contentSettings#57e28221 flags:# sensitive_enabled:flags.0?true sensitive_can_change:flags.1?true = account.ContentSettings;
//
func (m *TLAccountContentSettings) To_Account_ContentSettings() *Account_ContentSettings {
	m.Data2.Cmd = Cmd_AccountContentSettings
	return m.Data2
}

//  + TL_AccountNoPassword
//  + TL_AccountPassword
//

//  account.noPassword#96dabc18 new_salt:bytes email_unconfirmed_pattern:string = account.Password;
//
func (m *Account_Password) To_AccountNoPassword() *TLAccountNoPassword {
	return &TLAccountNoPassword{
		Data2: m,
	}
}

//  account.password#7c18141c current_salt:bytes new_salt:bytes hint:string has_recovery:Bool email_unconfirmed_pattern:string = account.Password;
//  account.password#ad2641f8 flags:# has_recovery:flags.0?true has_secure_values:flags.1?true has_password:flags.2?true current_algo:flags.2?PasswordKdfAlgo srp_B:flags.2?bytes srp_id:flags.2?long hint:flags.3?string email_unconfirmed_pattern:flags.4?string new_algo:PasswordKdfAlgo new_secure_algo:SecurePasswordKdfAlgo secure_random:bytes = account.Password;
//
func (m *Account_Password) To_AccountPassword() *TLAccountPassword {
	return &TLAccountPassword{
		Data2: m,
	}
}

//  account.noPassword#96dabc18 new_salt:bytes email_unconfirmed_pattern:string = account.Password;
//
func (m *TLAccountNoPassword) To_Account_Password() *Account_Password {
	m.Data2.Cmd = Cmd_AccountNoPassword
	return m.Data2
}

//  account.password#7c18141c current_salt:bytes new_salt:bytes hint:string has_recovery:Bool email_unconfirmed_pattern:string = account.Password;
//  account.password#ad2641f8 flags:# has_recovery:flags.0?true has_secure_values:flags.1?true has_password:flags.2?true current_algo:flags.2?PasswordKdfAlgo srp_B:flags.2?bytes srp_id:flags.2?long hint:flags.3?string email_unconfirmed_pattern:flags.4?string new_algo:PasswordKdfAlgo new_secure_algo:SecurePasswordKdfAlgo secure_random:bytes = account.Password;
//
func (m *TLAccountPassword) To_Account_Password() *Account_Password {
	m.Data2.Cmd = Cmd_AccountPassword
	return m.Data2
}

//  + TL_AccountPasswordInputSettings
//

//  account.passwordInputSettings#86916deb flags:# new_salt:flags.0?bytes new_password_hash:flags.0?bytes hint:flags.0?string email:flags.1?string = account.PasswordInputSettings;
//  account.passwordInputSettings#c23727c9 flags:# new_algo:flags.0?PasswordKdfAlgo new_password_hash:flags.0?bytes hint:flags.0?string email:flags.1?string new_secure_settings:flags.2?SecureSecretSettings = account.PasswordInputSettings;
//
func (m *Account_PasswordInputSettings) To_AccountPasswordInputSettings() *TLAccountPasswordInputSettings {
	return &TLAccountPasswordInputSettings{
		Data2: m,
	}
}

//  account.passwordInputSettings#86916deb flags:# new_salt:flags.0?bytes new_password_hash:flags.0?bytes hint:flags.0?string email:flags.1?string = account.PasswordInputSettings;
//  account.passwordInputSettings#c23727c9 flags:# new_algo:flags.0?PasswordKdfAlgo new_password_hash:flags.0?bytes hint:flags.0?string email:flags.1?string new_secure_settings:flags.2?SecureSecretSettings = account.PasswordInputSettings;
//
func (m *TLAccountPasswordInputSettings) To_Account_PasswordInputSettings() *Account_PasswordInputSettings {
	m.Data2.Cmd = Cmd_AccountPasswordInputSettings
	return m.Data2
}

//  + TL_AccountPasswordSettings
//

//  account.passwordSettings#b7b72ab3 email:string = account.PasswordSettings;
//  account.passwordSettings#9a5c33e5 flags:# email:flags.0?string secure_settings:flags.1?SecureSecretSettings = account.PasswordSettings;
//
func (m *Account_PasswordSettings) To_AccountPasswordSettings() *TLAccountPasswordSettings {
	return &TLAccountPasswordSettings{
		Data2: m,
	}
}

//  account.passwordSettings#b7b72ab3 email:string = account.PasswordSettings;
//  account.passwordSettings#9a5c33e5 flags:# email:flags.0?string secure_settings:flags.1?SecureSecretSettings = account.PasswordSettings;
//
func (m *TLAccountPasswordSettings) To_Account_PasswordSettings() *Account_PasswordSettings {
	m.Data2.Cmd = Cmd_AccountPasswordSettings
	return m.Data2
}

//  + TL_AccountPrivacyRules
//

//  account.privacyRules#554abb6f rules:Vector<PrivacyRule> users:Vector<User> = account.PrivacyRules;
//  account.privacyRules#50a04e45 rules:Vector<PrivacyRule> chats:Vector<Chat> users:Vector<User> = account.PrivacyRules;
//
func (m *Account_PrivacyRules) To_AccountPrivacyRules() *TLAccountPrivacyRules {
	return &TLAccountPrivacyRules{
		Data2: m,
	}
}

//  account.privacyRules#554abb6f rules:Vector<PrivacyRule> users:Vector<User> = account.PrivacyRules;
//  account.privacyRules#50a04e45 rules:Vector<PrivacyRule> chats:Vector<Chat> users:Vector<User> = account.PrivacyRules;
//
func (m *TLAccountPrivacyRules) To_Account_PrivacyRules() *Account_PrivacyRules {
	m.Data2.Cmd = Cmd_AccountPrivacyRules
	return m.Data2
}

//  + TL_AccountSentEmailCode
//

//  account.sentEmailCode#811f854f email_pattern:string length:int = account.SentEmailCode;
//
func (m *Account_SentEmailCode) To_AccountSentEmailCode() *TLAccountSentEmailCode {
	return &TLAccountSentEmailCode{
		Data2: m,
	}
}

//  account.sentEmailCode#811f854f email_pattern:string length:int = account.SentEmailCode;
//
func (m *TLAccountSentEmailCode) To_Account_SentEmailCode() *Account_SentEmailCode {
	m.Data2.Cmd = Cmd_AccountSentEmailCode
	return m.Data2
}

//  + TL_AccountTakeout
//

//  account.takeout#4dba4501 id:long = account.Takeout;
//
func (m *Account_Takeout) To_AccountTakeout() *TLAccountTakeout {
	return &TLAccountTakeout{
		Data2: m,
	}
}

//  account.takeout#4dba4501 id:long = account.Takeout;
//
func (m *TLAccountTakeout) To_Account_Takeout() *Account_Takeout {
	m.Data2.Cmd = Cmd_AccountTakeout
	return m.Data2
}

//  + TL_AccountThemesNotModified
//  + TL_AccountThemes
//

//  account.themesNotModified#f41eb622 = account.Themes;
//
func (m *Account_Themes) To_AccountThemesNotModified() *TLAccountThemesNotModified {
	return &TLAccountThemesNotModified{
		Data2: m,
	}
}

//  account.themes#7f676421 hash:int themes:Vector<Theme> = account.Themes;
//
func (m *Account_Themes) To_AccountThemes() *TLAccountThemes {
	return &TLAccountThemes{
		Data2: m,
	}
}

//  account.themesNotModified#f41eb622 = account.Themes;
//
func (m *TLAccountThemesNotModified) To_Account_Themes() *Account_Themes {
	m.Data2.Cmd = Cmd_AccountThemesNotModified
	return m.Data2
}

//  account.themes#7f676421 hash:int themes:Vector<Theme> = account.Themes;
//
func (m *TLAccountThemes) To_Account_Themes() *Account_Themes {
	m.Data2.Cmd = Cmd_AccountThemes
	return m.Data2
}

//  + TL_AccountTmpPassword
//

//  account.tmpPassword#db64fd34 tmp_password:bytes valid_until:int = account.TmpPassword;
//
func (m *Account_TmpPassword) To_AccountTmpPassword() *TLAccountTmpPassword {
	return &TLAccountTmpPassword{
		Data2: m,
	}
}

//  account.tmpPassword#db64fd34 tmp_password:bytes valid_until:int = account.TmpPassword;
//
func (m *TLAccountTmpPassword) To_Account_TmpPassword() *Account_TmpPassword {
	m.Data2.Cmd = Cmd_AccountTmpPassword
	return m.Data2
}

//  + TL_AccountWallPapersNotModified
//  + TL_AccountWallPapers
//

//  account.wallPapersNotModified#1c199183 = account.WallPapers;
//
func (m *Account_WallPapers) To_AccountWallPapersNotModified() *TLAccountWallPapersNotModified {
	return &TLAccountWallPapersNotModified{
		Data2: m,
	}
}

//  account.wallPapers#702b65a9 hash:int wallpapers:Vector<WallPaper> = account.WallPapers;
//
func (m *Account_WallPapers) To_AccountWallPapers() *TLAccountWallPapers {
	return &TLAccountWallPapers{
		Data2: m,
	}
}

//  account.wallPapersNotModified#1c199183 = account.WallPapers;
//
func (m *TLAccountWallPapersNotModified) To_Account_WallPapers() *Account_WallPapers {
	m.Data2.Cmd = Cmd_AccountWallPapersNotModified
	return m.Data2
}

//  account.wallPapers#702b65a9 hash:int wallpapers:Vector<WallPaper> = account.WallPapers;
//
func (m *TLAccountWallPapers) To_Account_WallPapers() *Account_WallPapers {
	m.Data2.Cmd = Cmd_AccountWallPapers
	return m.Data2
}

//  + TL_AccountWebAuthorizations
//

//  account.webAuthorizations#ed56c9fc authorizations:Vector<WebAuthorization> users:Vector<User> = account.WebAuthorizations;
//
func (m *Account_WebAuthorizations) To_AccountWebAuthorizations() *TLAccountWebAuthorizations {
	return &TLAccountWebAuthorizations{
		Data2: m,
	}
}

//  account.webAuthorizations#ed56c9fc authorizations:Vector<WebAuthorization> users:Vector<User> = account.WebAuthorizations;
//
func (m *TLAccountWebAuthorizations) To_Account_WebAuthorizations() *Account_WebAuthorizations {
	m.Data2.Cmd = Cmd_AccountWebAuthorizations
	return m.Data2
}

//  + TL_AuthAuthorization
//  + TL_AuthAuthorizationSignUpRequired
//

//  auth.authorization#cd050916 flags:# tmp_sessions:flags.0?int user:User = auth.Authorization;
//  auth.authorization#ff036af1 user:User = auth.Authorization;
//
func (m *Auth_Authorization) To_AuthAuthorization() *TLAuthAuthorization {
	return &TLAuthAuthorization{
		Data2: m,
	}
}

//  auth.authorizationSignUpRequired#44747e9a flags:# terms_of_service:flags.0?help.TermsOfService = auth.Authorization;
//
func (m *Auth_Authorization) To_AuthAuthorizationSignUpRequired() *TLAuthAuthorizationSignUpRequired {
	return &TLAuthAuthorizationSignUpRequired{
		Data2: m,
	}
}

//  auth.authorization#cd050916 flags:# tmp_sessions:flags.0?int user:User = auth.Authorization;
//  auth.authorization#ff036af1 user:User = auth.Authorization;
//
func (m *TLAuthAuthorization) To_Auth_Authorization() *Auth_Authorization {
	m.Data2.Cmd = Cmd_AuthAuthorization
	return m.Data2
}

//  auth.authorizationSignUpRequired#44747e9a flags:# terms_of_service:flags.0?help.TermsOfService = auth.Authorization;
//
func (m *TLAuthAuthorizationSignUpRequired) To_Auth_Authorization() *Auth_Authorization {
	m.Data2.Cmd = Cmd_AuthAuthorizationSignUpRequired
	return m.Data2
}

//  + TL_AuthCheckedPhone
//

//  auth.checkedPhone#811ea28e phone_registered:Bool = auth.CheckedPhone;
//
func (m *Auth_CheckedPhone) To_AuthCheckedPhone() *TLAuthCheckedPhone {
	return &TLAuthCheckedPhone{
		Data2: m,
	}
}

//  auth.checkedPhone#811ea28e phone_registered:Bool = auth.CheckedPhone;
//
func (m *TLAuthCheckedPhone) To_Auth_CheckedPhone() *Auth_CheckedPhone {
	m.Data2.Cmd = Cmd_AuthCheckedPhone
	return m.Data2
}

//  + TL_AuthCodeTypeSms
//  + TL_AuthCodeTypeCall
//  + TL_AuthCodeTypeFlashCall
//

//  auth.codeTypeSms#72a3158c = auth.CodeType;
//
func (m *Auth_CodeType) To_AuthCodeTypeSms() *TLAuthCodeTypeSms {
	return &TLAuthCodeTypeSms{
		Data2: m,
	}
}

//  auth.codeTypeCall#741cd3e3 = auth.CodeType;
//
func (m *Auth_CodeType) To_AuthCodeTypeCall() *TLAuthCodeTypeCall {
	return &TLAuthCodeTypeCall{
		Data2: m,
	}
}

//  auth.codeTypeFlashCall#226ccefb = auth.CodeType;
//
func (m *Auth_CodeType) To_AuthCodeTypeFlashCall() *TLAuthCodeTypeFlashCall {
	return &TLAuthCodeTypeFlashCall{
		Data2: m,
	}
}

//  auth.codeTypeSms#72a3158c = auth.CodeType;
//
func (m *TLAuthCodeTypeSms) To_Auth_CodeType() *Auth_CodeType {
	m.Data2.Cmd = Cmd_AuthCodeTypeSms
	return m.Data2
}

//  auth.codeTypeCall#741cd3e3 = auth.CodeType;
//
func (m *TLAuthCodeTypeCall) To_Auth_CodeType() *Auth_CodeType {
	m.Data2.Cmd = Cmd_AuthCodeTypeCall
	return m.Data2
}

//  auth.codeTypeFlashCall#226ccefb = auth.CodeType;
//
func (m *TLAuthCodeTypeFlashCall) To_Auth_CodeType() *Auth_CodeType {
	m.Data2.Cmd = Cmd_AuthCodeTypeFlashCall
	return m.Data2
}

//  + TL_AuthExportedAuthorization
//

//  auth.exportedAuthorization#df969c2d id:int bytes:bytes = auth.ExportedAuthorization;
//
func (m *Auth_ExportedAuthorization) To_AuthExportedAuthorization() *TLAuthExportedAuthorization {
	return &TLAuthExportedAuthorization{
		Data2: m,
	}
}

//  auth.exportedAuthorization#df969c2d id:int bytes:bytes = auth.ExportedAuthorization;
//
func (m *TLAuthExportedAuthorization) To_Auth_ExportedAuthorization() *Auth_ExportedAuthorization {
	m.Data2.Cmd = Cmd_AuthExportedAuthorization
	return m.Data2
}

//  + TL_AuthLoginToken
//  + TL_AuthLoginTokenMigrateTo
//  + TL_AuthLoginTokenSuccess
//

//  auth.loginToken#629f1980 expires:int token:bytes = auth.LoginToken;
//
func (m *Auth_LoginToken) To_AuthLoginToken() *TLAuthLoginToken {
	return &TLAuthLoginToken{
		Data2: m,
	}
}

//  auth.loginTokenMigrateTo#68e9916 dc_id:int token:bytes = auth.LoginToken;
//
func (m *Auth_LoginToken) To_AuthLoginTokenMigrateTo() *TLAuthLoginTokenMigrateTo {
	return &TLAuthLoginTokenMigrateTo{
		Data2: m,
	}
}

//  auth.loginTokenSuccess#390d5c5e authorization:auth.Authorization = auth.LoginToken;
//
func (m *Auth_LoginToken) To_AuthLoginTokenSuccess() *TLAuthLoginTokenSuccess {
	return &TLAuthLoginTokenSuccess{
		Data2: m,
	}
}

//  auth.loginToken#629f1980 expires:int token:bytes = auth.LoginToken;
//
func (m *TLAuthLoginToken) To_Auth_LoginToken() *Auth_LoginToken {
	m.Data2.Cmd = Cmd_AuthLoginToken
	return m.Data2
}

//  auth.loginTokenMigrateTo#68e9916 dc_id:int token:bytes = auth.LoginToken;
//
func (m *TLAuthLoginTokenMigrateTo) To_Auth_LoginToken() *Auth_LoginToken {
	m.Data2.Cmd = Cmd_AuthLoginTokenMigrateTo
	return m.Data2
}

//  auth.loginTokenSuccess#390d5c5e authorization:auth.Authorization = auth.LoginToken;
//
func (m *TLAuthLoginTokenSuccess) To_Auth_LoginToken() *Auth_LoginToken {
	m.Data2.Cmd = Cmd_AuthLoginTokenSuccess
	return m.Data2
}

//  + TL_AuthPasswordRecovery
//

//  auth.passwordRecovery#137948a5 email_pattern:string = auth.PasswordRecovery;
//
func (m *Auth_PasswordRecovery) To_AuthPasswordRecovery() *TLAuthPasswordRecovery {
	return &TLAuthPasswordRecovery{
		Data2: m,
	}
}

//  auth.passwordRecovery#137948a5 email_pattern:string = auth.PasswordRecovery;
//
func (m *TLAuthPasswordRecovery) To_Auth_PasswordRecovery() *Auth_PasswordRecovery {
	m.Data2.Cmd = Cmd_AuthPasswordRecovery
	return m.Data2
}

//  + TL_AuthSentCode
//

//  auth.sentCode#5e002502 flags:# phone_registered:flags.0?true type:auth.SentCodeType phone_code_hash:string next_type:flags.1?auth.CodeType timeout:flags.2?int = auth.SentCode;
//  auth.sentCode#38faab5f flags:# phone_registered:flags.0?true type:auth.SentCodeType phone_code_hash:string next_type:flags.1?auth.CodeType timeout:flags.2?int terms_of_service:flags.3?help.TermsOfService = auth.SentCode;
//
func (m *Auth_SentCode) To_AuthSentCode() *TLAuthSentCode {
	return &TLAuthSentCode{
		Data2: m,
	}
}

//  auth.sentCode#5e002502 flags:# phone_registered:flags.0?true type:auth.SentCodeType phone_code_hash:string next_type:flags.1?auth.CodeType timeout:flags.2?int = auth.SentCode;
//  auth.sentCode#38faab5f flags:# phone_registered:flags.0?true type:auth.SentCodeType phone_code_hash:string next_type:flags.1?auth.CodeType timeout:flags.2?int terms_of_service:flags.3?help.TermsOfService = auth.SentCode;
//
func (m *TLAuthSentCode) To_Auth_SentCode() *Auth_SentCode {
	m.Data2.Cmd = Cmd_AuthSentCode
	return m.Data2
}

//  + TL_AuthSentCodeTypeApp
//  + TL_AuthSentCodeTypeSms
//  + TL_AuthSentCodeTypeCall
//  + TL_AuthSentCodeTypeFlashCall
//

//  auth.sentCodeTypeApp#3dbb5986 length:int = auth.SentCodeType;
//
func (m *Auth_SentCodeType) To_AuthSentCodeTypeApp() *TLAuthSentCodeTypeApp {
	return &TLAuthSentCodeTypeApp{
		Data2: m,
	}
}

//  auth.sentCodeTypeSms#c000bba2 length:int = auth.SentCodeType;
//
func (m *Auth_SentCodeType) To_AuthSentCodeTypeSms() *TLAuthSentCodeTypeSms {
	return &TLAuthSentCodeTypeSms{
		Data2: m,
	}
}

//  auth.sentCodeTypeCall#5353e5a7 length:int = auth.SentCodeType;
//
func (m *Auth_SentCodeType) To_AuthSentCodeTypeCall() *TLAuthSentCodeTypeCall {
	return &TLAuthSentCodeTypeCall{
		Data2: m,
	}
}

//  auth.sentCodeTypeFlashCall#ab03c6d9 pattern:string = auth.SentCodeType;
//
func (m *Auth_SentCodeType) To_AuthSentCodeTypeFlashCall() *TLAuthSentCodeTypeFlashCall {
	return &TLAuthSentCodeTypeFlashCall{
		Data2: m,
	}
}

//  auth.sentCodeTypeApp#3dbb5986 length:int = auth.SentCodeType;
//
func (m *TLAuthSentCodeTypeApp) To_Auth_SentCodeType() *Auth_SentCodeType {
	m.Data2.Cmd = Cmd_AuthSentCodeTypeApp
	return m.Data2
}

//  auth.sentCodeTypeSms#c000bba2 length:int = auth.SentCodeType;
//
func (m *TLAuthSentCodeTypeSms) To_Auth_SentCodeType() *Auth_SentCodeType {
	m.Data2.Cmd = Cmd_AuthSentCodeTypeSms
	return m.Data2
}

//  auth.sentCodeTypeCall#5353e5a7 length:int = auth.SentCodeType;
//
func (m *TLAuthSentCodeTypeCall) To_Auth_SentCodeType() *Auth_SentCodeType {
	m.Data2.Cmd = Cmd_AuthSentCodeTypeCall
	return m.Data2
}

//  auth.sentCodeTypeFlashCall#ab03c6d9 pattern:string = auth.SentCodeType;
//
func (m *TLAuthSentCodeTypeFlashCall) To_Auth_SentCodeType() *Auth_SentCodeType {
	m.Data2.Cmd = Cmd_AuthSentCodeTypeFlashCall
	return m.Data2
}

//  + TL_Authorization
//

//  authorization#7bf2e6f6 hash:long flags:int device_model:string platform:string system_version:string api_id:int app_name:string app_version:string date_created:int date_active:int ip:string country:string region:string = Authorization;
//  authorization#ad01d61d flags:# current:flags.0?true official_app:flags.1?true password_pending:flags.2?true hash:long device_model:string platform:string system_version:string api_id:int app_name:string app_version:string date_created:int date_active:int ip:string country:string region:string = Authorization;
//
func (m *Authorization) To_Authorization() *TLAuthorization {
	return &TLAuthorization{
		Data2: m,
	}
}

//  authorization#7bf2e6f6 hash:long flags:int device_model:string platform:string system_version:string api_id:int app_name:string app_version:string date_created:int date_active:int ip:string country:string region:string = Authorization;
//  authorization#ad01d61d flags:# current:flags.0?true official_app:flags.1?true password_pending:flags.2?true hash:long device_model:string platform:string system_version:string api_id:int app_name:string app_version:string date_created:int date_active:int ip:string country:string region:string = Authorization;
//
func (m *TLAuthorization) To_Authorization() *Authorization {
	m.Data2.Cmd = Cmd_Authorization
	return m.Data2
}

//  + TL_AutoDownloadSettings
//

//  autoDownloadSettings#d246fd47 flags:# disabled:flags.0?true video_preload_large:flags.1?true audio_preload_next:flags.2?true phonecalls_less_data:flags.3?true photo_size_max:int video_size_max:int file_size_max:int = AutoDownloadSettings;
//  autoDownloadSettings#e04232f3 flags:# disabled:flags.0?true video_preload_large:flags.1?true audio_preload_next:flags.2?true phonecalls_less_data:flags.3?true photo_size_max:int video_size_max:int file_size_max:int video_upload_maxbitrate:int = AutoDownloadSettings;
//
func (m *AutoDownloadSettings) To_AutoDownloadSettings() *TLAutoDownloadSettings {
	return &TLAutoDownloadSettings{
		Data2: m,
	}
}

//  autoDownloadSettings#d246fd47 flags:# disabled:flags.0?true video_preload_large:flags.1?true audio_preload_next:flags.2?true phonecalls_less_data:flags.3?true photo_size_max:int video_size_max:int file_size_max:int = AutoDownloadSettings;
//  autoDownloadSettings#e04232f3 flags:# disabled:flags.0?true video_preload_large:flags.1?true audio_preload_next:flags.2?true phonecalls_less_data:flags.3?true photo_size_max:int video_size_max:int file_size_max:int video_upload_maxbitrate:int = AutoDownloadSettings;
//
func (m *TLAutoDownloadSettings) To_AutoDownloadSettings() *AutoDownloadSettings {
	m.Data2.Cmd = Cmd_AutoDownloadSettings
	return m.Data2
}

//  + TL_BadMsgNotification
//  + TL_BadServerSalt
//

//  bad_msg_notification#a7eff811 bad_msg_id:long bad_msg_seqno:int error_code:int = BadMsgNotification;
//
func (m *BadMsgNotification) To_BadMsgNotification() *TLBadMsgNotification {
	return &TLBadMsgNotification{
		Data2: m,
	}
}

//  bad_server_salt#edab447b bad_msg_id:long bad_msg_seqno:int error_code:int new_server_salt:long = BadMsgNotification;
//
func (m *BadMsgNotification) To_BadServerSalt() *TLBadServerSalt {
	return &TLBadServerSalt{
		Data2: m,
	}
}

//  bad_msg_notification#a7eff811 bad_msg_id:long bad_msg_seqno:int error_code:int = BadMsgNotification;
//
func (m *TLBadMsgNotification) To_BadMsgNotification() *BadMsgNotification {
	m.Data2.Cmd = Cmd_BadMsgNotification
	return m.Data2
}

//  bad_server_salt#edab447b bad_msg_id:long bad_msg_seqno:int error_code:int new_server_salt:long = BadMsgNotification;
//
func (m *TLBadServerSalt) To_BadMsgNotification() *BadMsgNotification {
	m.Data2.Cmd = Cmd_BadServerSalt
	return m.Data2
}

//  + TL_BankCardOpenUrl
//

//  bankCardOpenUrl#f568028a url:string name:string = BankCardOpenUrl;
//
func (m *BankCardOpenUrl) To_BankCardOpenUrl() *TLBankCardOpenUrl {
	return &TLBankCardOpenUrl{
		Data2: m,
	}
}

//  bankCardOpenUrl#f568028a url:string name:string = BankCardOpenUrl;
//
func (m *TLBankCardOpenUrl) To_BankCardOpenUrl() *BankCardOpenUrl {
	m.Data2.Cmd = Cmd_BankCardOpenUrl
	return m.Data2
}

//  + TL_BaseThemeClassic
//  + TL_BaseThemeDay
//  + TL_BaseThemeNight
//  + TL_BaseThemeTinted
//  + TL_BaseThemeArctic
//

//  baseThemeClassic#c3a12462 = BaseTheme;
//
func (m *BaseTheme) To_BaseThemeClassic() *TLBaseThemeClassic {
	return &TLBaseThemeClassic{
		Data2: m,
	}
}

//  baseThemeDay#fbd81688 = BaseTheme;
//
func (m *BaseTheme) To_BaseThemeDay() *TLBaseThemeDay {
	return &TLBaseThemeDay{
		Data2: m,
	}
}

//  baseThemeNight#b7b31ea8 = BaseTheme;
//
func (m *BaseTheme) To_BaseThemeNight() *TLBaseThemeNight {
	return &TLBaseThemeNight{
		Data2: m,
	}
}

//  baseThemeTinted#6d5f77ee = BaseTheme;
//
func (m *BaseTheme) To_BaseThemeTinted() *TLBaseThemeTinted {
	return &TLBaseThemeTinted{
		Data2: m,
	}
}

//  baseThemeArctic#5b11125a = BaseTheme;
//
func (m *BaseTheme) To_BaseThemeArctic() *TLBaseThemeArctic {
	return &TLBaseThemeArctic{
		Data2: m,
	}
}

//  baseThemeClassic#c3a12462 = BaseTheme;
//
func (m *TLBaseThemeClassic) To_BaseTheme() *BaseTheme {
	m.Data2.Cmd = Cmd_BaseThemeClassic
	return m.Data2
}

//  baseThemeDay#fbd81688 = BaseTheme;
//
func (m *TLBaseThemeDay) To_BaseTheme() *BaseTheme {
	m.Data2.Cmd = Cmd_BaseThemeDay
	return m.Data2
}

//  baseThemeNight#b7b31ea8 = BaseTheme;
//
func (m *TLBaseThemeNight) To_BaseTheme() *BaseTheme {
	m.Data2.Cmd = Cmd_BaseThemeNight
	return m.Data2
}

//  baseThemeTinted#6d5f77ee = BaseTheme;
//
func (m *TLBaseThemeTinted) To_BaseTheme() *BaseTheme {
	m.Data2.Cmd = Cmd_BaseThemeTinted
	return m.Data2
}

//  baseThemeArctic#5b11125a = BaseTheme;
//
func (m *TLBaseThemeArctic) To_BaseTheme() *BaseTheme {
	m.Data2.Cmd = Cmd_BaseThemeArctic
	return m.Data2
}

//  + TL_BindAuthKeyInner
//

//  bind_auth_key_inner#75a3f765 nonce:long temp_auth_key_id:long perm_auth_key_id:long temp_session_id:long expires_at:int = BindAuthKeyInner;
//
func (m *BindAuthKeyInner) To_BindAuthKeyInner() *TLBindAuthKeyInner {
	return &TLBindAuthKeyInner{
		Data2: m,
	}
}

//  bind_auth_key_inner#75a3f765 nonce:long temp_auth_key_id:long perm_auth_key_id:long temp_session_id:long expires_at:int = BindAuthKeyInner;
//
func (m *TLBindAuthKeyInner) To_BindAuthKeyInner() *BindAuthKeyInner {
	m.Data2.Cmd = Cmd_BindAuthKeyInner
	return m.Data2
}

//  + TL_BoolFalse
//  + TL_BoolTrue
//

//  boolFalse#bc799737 = Bool;
//
func (m *Bool) To_BoolFalse() *TLBoolFalse {
	return &TLBoolFalse{
		Data2: m,
	}
}

//  boolTrue#997275b5 = Bool;
//
func (m *Bool) To_BoolTrue() *TLBoolTrue {
	return &TLBoolTrue{
		Data2: m,
	}
}

//  boolFalse#bc799737 = Bool;
//
func (m *TLBoolFalse) To_Bool() *Bool {
	m.Data2.Cmd = Cmd_BoolFalse
	return m.Data2
}

//  boolTrue#997275b5 = Bool;
//
func (m *TLBoolTrue) To_Bool() *Bool {
	m.Data2.Cmd = Cmd_BoolTrue
	return m.Data2
}

//  + TL_BotCommand
//

//  botCommand#c27ac8c7 command:string description:string = BotCommand;
//
func (m *BotCommand) To_BotCommand() *TLBotCommand {
	return &TLBotCommand{
		Data2: m,
	}
}

//  botCommand#c27ac8c7 command:string description:string = BotCommand;
//
func (m *TLBotCommand) To_BotCommand() *BotCommand {
	m.Data2.Cmd = Cmd_BotCommand
	return m.Data2
}

//  + TL_BotInfo
//

//  botInfo#98e81d3a user_id:int description:string commands:Vector<BotCommand> = BotInfo;
//
func (m *BotInfo) To_BotInfo() *TLBotInfo {
	return &TLBotInfo{
		Data2: m,
	}
}

//  botInfo#98e81d3a user_id:int description:string commands:Vector<BotCommand> = BotInfo;
//
func (m *TLBotInfo) To_BotInfo() *BotInfo {
	m.Data2.Cmd = Cmd_BotInfo
	return m.Data2
}

//  + TL_BotInlineMessageMediaAuto
//  + TL_BotInlineMessageText
//  + TL_BotInlineMessageMediaGeo
//  + TL_BotInlineMessageMediaVenue
//  + TL_BotInlineMessageMediaContact
//

//  botInlineMessageMediaAuto#a74b15b flags:# caption:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//  botInlineMessageMediaAuto#764cf810 flags:# message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//
func (m *BotInlineMessage) To_BotInlineMessageMediaAuto() *TLBotInlineMessageMediaAuto {
	return &TLBotInlineMessageMediaAuto{
		Data2: m,
	}
}

//  botInlineMessageText#8c7f65e2 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//
func (m *BotInlineMessage) To_BotInlineMessageText() *TLBotInlineMessageText {
	return &TLBotInlineMessageText{
		Data2: m,
	}
}

//  botInlineMessageMediaGeo#3a8fd8b8 flags:# geo:GeoPoint reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//  botInlineMessageMediaGeo#b722de65 flags:# geo:GeoPoint period:int reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//  botInlineMessageMediaGeo#51846fd flags:# geo:GeoPoint heading:flags.0?int period:flags.1?int proximity_notification_radius:flags.3?int reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//
func (m *BotInlineMessage) To_BotInlineMessageMediaGeo() *TLBotInlineMessageMediaGeo {
	return &TLBotInlineMessageMediaGeo{
		Data2: m,
	}
}

//  botInlineMessageMediaVenue#4366232e flags:# geo:GeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//  botInlineMessageMediaVenue#8a86659c flags:# geo:GeoPoint title:string address:string provider:string venue_id:string venue_type:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//
func (m *BotInlineMessage) To_BotInlineMessageMediaVenue() *TLBotInlineMessageMediaVenue {
	return &TLBotInlineMessageMediaVenue{
		Data2: m,
	}
}

//  botInlineMessageMediaContact#35edb4d4 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//  botInlineMessageMediaContact#18d1cdc2 flags:# phone_number:string first_name:string last_name:string vcard:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//
func (m *BotInlineMessage) To_BotInlineMessageMediaContact() *TLBotInlineMessageMediaContact {
	return &TLBotInlineMessageMediaContact{
		Data2: m,
	}
}

//  botInlineMessageMediaAuto#a74b15b flags:# caption:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//  botInlineMessageMediaAuto#764cf810 flags:# message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//
func (m *TLBotInlineMessageMediaAuto) To_BotInlineMessage() *BotInlineMessage {
	m.Data2.Cmd = Cmd_BotInlineMessageMediaAuto
	return m.Data2
}

//  botInlineMessageText#8c7f65e2 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//
func (m *TLBotInlineMessageText) To_BotInlineMessage() *BotInlineMessage {
	m.Data2.Cmd = Cmd_BotInlineMessageText
	return m.Data2
}

//  botInlineMessageMediaGeo#3a8fd8b8 flags:# geo:GeoPoint reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//  botInlineMessageMediaGeo#b722de65 flags:# geo:GeoPoint period:int reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//  botInlineMessageMediaGeo#51846fd flags:# geo:GeoPoint heading:flags.0?int period:flags.1?int proximity_notification_radius:flags.3?int reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//
func (m *TLBotInlineMessageMediaGeo) To_BotInlineMessage() *BotInlineMessage {
	m.Data2.Cmd = Cmd_BotInlineMessageMediaGeo
	return m.Data2
}

//  botInlineMessageMediaVenue#4366232e flags:# geo:GeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//  botInlineMessageMediaVenue#8a86659c flags:# geo:GeoPoint title:string address:string provider:string venue_id:string venue_type:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//
func (m *TLBotInlineMessageMediaVenue) To_BotInlineMessage() *BotInlineMessage {
	m.Data2.Cmd = Cmd_BotInlineMessageMediaVenue
	return m.Data2
}

//  botInlineMessageMediaContact#35edb4d4 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//  botInlineMessageMediaContact#18d1cdc2 flags:# phone_number:string first_name:string last_name:string vcard:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
//
func (m *TLBotInlineMessageMediaContact) To_BotInlineMessage() *BotInlineMessage {
	m.Data2.Cmd = Cmd_BotInlineMessageMediaContact
	return m.Data2
}

//  + TL_BotInlineResult
//  + TL_BotInlineMediaResult
//

//  botInlineResult#9bebaeb9 flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:BotInlineMessage = BotInlineResult;
//  botInlineResult#11965f3a flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb:flags.4?WebDocument content:flags.5?WebDocument send_message:BotInlineMessage = BotInlineResult;
//
func (m *BotInlineResult) To_BotInlineResult() *TLBotInlineResult {
	return &TLBotInlineResult{
		Data2: m,
	}
}

//  botInlineMediaResult#17db940b flags:# id:string type:string photo:flags.0?Photo document:flags.1?Document title:flags.2?string description:flags.3?string send_message:BotInlineMessage = BotInlineResult;
//
func (m *BotInlineResult) To_BotInlineMediaResult() *TLBotInlineMediaResult {
	return &TLBotInlineMediaResult{
		Data2: m,
	}
}

//  botInlineResult#9bebaeb9 flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:BotInlineMessage = BotInlineResult;
//  botInlineResult#11965f3a flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb:flags.4?WebDocument content:flags.5?WebDocument send_message:BotInlineMessage = BotInlineResult;
//
func (m *TLBotInlineResult) To_BotInlineResult() *BotInlineResult {
	m.Data2.Cmd = Cmd_BotInlineResult
	return m.Data2
}

//  botInlineMediaResult#17db940b flags:# id:string type:string photo:flags.0?Photo document:flags.1?Document title:flags.2?string description:flags.3?string send_message:BotInlineMessage = BotInlineResult;
//
func (m *TLBotInlineMediaResult) To_BotInlineResult() *BotInlineResult {
	m.Data2.Cmd = Cmd_BotInlineMediaResult
	return m.Data2
}

//  + TL_CdnConfig
//

//  cdnConfig#5725e40a public_keys:Vector<CdnPublicKey> = CdnConfig;
//
func (m *CdnConfig) To_CdnConfig() *TLCdnConfig {
	return &TLCdnConfig{
		Data2: m,
	}
}

//  cdnConfig#5725e40a public_keys:Vector<CdnPublicKey> = CdnConfig;
//
func (m *TLCdnConfig) To_CdnConfig() *CdnConfig {
	m.Data2.Cmd = Cmd_CdnConfig
	return m.Data2
}

//  + TL_CdnFileHash
//

//  cdnFileHash#77eec38f offset:int limit:int hash:bytes = CdnFileHash;
//
func (m *CdnFileHash) To_CdnFileHash() *TLCdnFileHash {
	return &TLCdnFileHash{
		Data2: m,
	}
}

//  cdnFileHash#77eec38f offset:int limit:int hash:bytes = CdnFileHash;
//
func (m *TLCdnFileHash) To_CdnFileHash() *CdnFileHash {
	m.Data2.Cmd = Cmd_CdnFileHash
	return m.Data2
}

//  + TL_CdnPublicKey
//

//  cdnPublicKey#c982eaba dc_id:int public_key:string = CdnPublicKey;
//
func (m *CdnPublicKey) To_CdnPublicKey() *TLCdnPublicKey {
	return &TLCdnPublicKey{
		Data2: m,
	}
}

//  cdnPublicKey#c982eaba dc_id:int public_key:string = CdnPublicKey;
//
func (m *TLCdnPublicKey) To_CdnPublicKey() *CdnPublicKey {
	m.Data2.Cmd = Cmd_CdnPublicKey
	return m.Data2
}

//  + TL_ChannelAdminLogEvent
//

//  channelAdminLogEvent#3b5a3e40 id:long date:int user_id:int action:ChannelAdminLogEventAction = ChannelAdminLogEvent;
//
func (m *ChannelAdminLogEvent) To_ChannelAdminLogEvent() *TLChannelAdminLogEvent {
	return &TLChannelAdminLogEvent{
		Data2: m,
	}
}

//  channelAdminLogEvent#3b5a3e40 id:long date:int user_id:int action:ChannelAdminLogEventAction = ChannelAdminLogEvent;
//
func (m *TLChannelAdminLogEvent) To_ChannelAdminLogEvent() *ChannelAdminLogEvent {
	m.Data2.Cmd = Cmd_ChannelAdminLogEvent
	return m.Data2
}

//  + TL_ChannelAdminLogEventActionChangeTitle
//  + TL_ChannelAdminLogEventActionChangeAbout
//  + TL_ChannelAdminLogEventActionChangeUsername
//  + TL_ChannelAdminLogEventActionChangePhoto
//  + TL_ChannelAdminLogEventActionToggleInvites
//  + TL_ChannelAdminLogEventActionToggleSignatures
//  + TL_ChannelAdminLogEventActionUpdatePinned
//  + TL_ChannelAdminLogEventActionEditMessage
//  + TL_ChannelAdminLogEventActionDeleteMessage
//  + TL_ChannelAdminLogEventActionParticipantJoin
//  + TL_ChannelAdminLogEventActionParticipantLeave
//  + TL_ChannelAdminLogEventActionParticipantInvite
//  + TL_ChannelAdminLogEventActionParticipantToggleBan
//  + TL_ChannelAdminLogEventActionParticipantToggleAdmin
//  + TL_ChannelAdminLogEventActionChangeStickerSet
//  + TL_ChannelAdminLogEventActionTogglePreHistoryHidden
//  + TL_ChannelAdminLogEventActionDefaultBannedRights
//  + TL_ChannelAdminLogEventActionStopPoll
//  + TL_ChannelAdminLogEventActionChangeLinkedChat
//  + TL_ChannelAdminLogEventActionChangeLocation
//  + TL_ChannelAdminLogEventActionToggleSlowMode
//  + TL_ChannelAdminLogEventActionStartGroupCall
//  + TL_ChannelAdminLogEventActionDiscardGroupCall
//  + TL_ChannelAdminLogEventActionParticipantMute
//  + TL_ChannelAdminLogEventActionParticipantUnmute
//  + TL_ChannelAdminLogEventActionToggleGroupCallSetting
//

//  channelAdminLogEventActionChangeTitle#e6dfb825 prev_value:string new_value:string = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangeTitle() *TLChannelAdminLogEventActionChangeTitle {
	return &TLChannelAdminLogEventActionChangeTitle{
		Data2: m,
	}
}

//  channelAdminLogEventActionChangeAbout#55188a2e prev_value:string new_value:string = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangeAbout() *TLChannelAdminLogEventActionChangeAbout {
	return &TLChannelAdminLogEventActionChangeAbout{
		Data2: m,
	}
}

//  channelAdminLogEventActionChangeUsername#6a4afc38 prev_value:string new_value:string = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangeUsername() *TLChannelAdminLogEventActionChangeUsername {
	return &TLChannelAdminLogEventActionChangeUsername{
		Data2: m,
	}
}

//  channelAdminLogEventActionChangePhoto#b82f55c3 prev_photo:ChatPhoto new_photo:ChatPhoto = ChannelAdminLogEventAction;
//  channelAdminLogEventActionChangePhoto#434bd2af prev_photo:Photo new_photo:Photo = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangePhoto() *TLChannelAdminLogEventActionChangePhoto {
	return &TLChannelAdminLogEventActionChangePhoto{
		Data2: m,
	}
}

//  channelAdminLogEventActionToggleInvites#1b7907ae new_value:Bool = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionToggleInvites() *TLChannelAdminLogEventActionToggleInvites {
	return &TLChannelAdminLogEventActionToggleInvites{
		Data2: m,
	}
}

//  channelAdminLogEventActionToggleSignatures#26ae0971 new_value:Bool = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionToggleSignatures() *TLChannelAdminLogEventActionToggleSignatures {
	return &TLChannelAdminLogEventActionToggleSignatures{
		Data2: m,
	}
}

//  channelAdminLogEventActionUpdatePinned#e9e82c18 message:Message = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionUpdatePinned() *TLChannelAdminLogEventActionUpdatePinned {
	return &TLChannelAdminLogEventActionUpdatePinned{
		Data2: m,
	}
}

//  channelAdminLogEventActionEditMessage#709b2405 prev_message:Message new_message:Message = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionEditMessage() *TLChannelAdminLogEventActionEditMessage {
	return &TLChannelAdminLogEventActionEditMessage{
		Data2: m,
	}
}

//  channelAdminLogEventActionDeleteMessage#42e047bb message:Message = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionDeleteMessage() *TLChannelAdminLogEventActionDeleteMessage {
	return &TLChannelAdminLogEventActionDeleteMessage{
		Data2: m,
	}
}

//  channelAdminLogEventActionParticipantJoin#183040d3 = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantJoin() *TLChannelAdminLogEventActionParticipantJoin {
	return &TLChannelAdminLogEventActionParticipantJoin{
		Data2: m,
	}
}

//  channelAdminLogEventActionParticipantLeave#f89777f2 = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantLeave() *TLChannelAdminLogEventActionParticipantLeave {
	return &TLChannelAdminLogEventActionParticipantLeave{
		Data2: m,
	}
}

//  channelAdminLogEventActionParticipantInvite#e31c34d8 participant:ChannelParticipant = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantInvite() *TLChannelAdminLogEventActionParticipantInvite {
	return &TLChannelAdminLogEventActionParticipantInvite{
		Data2: m,
	}
}

//  channelAdminLogEventActionParticipantToggleBan#e6d83d7e prev_participant:ChannelParticipant new_participant:ChannelParticipant = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantToggleBan() *TLChannelAdminLogEventActionParticipantToggleBan {
	return &TLChannelAdminLogEventActionParticipantToggleBan{
		Data2: m,
	}
}

//  channelAdminLogEventActionParticipantToggleAdmin#d5676710 prev_participant:ChannelParticipant new_participant:ChannelParticipant = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantToggleAdmin() *TLChannelAdminLogEventActionParticipantToggleAdmin {
	return &TLChannelAdminLogEventActionParticipantToggleAdmin{
		Data2: m,
	}
}

//  channelAdminLogEventActionChangeStickerSet#b1c3caa7 prev_stickerset:InputStickerSet new_stickerset:InputStickerSet = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangeStickerSet() *TLChannelAdminLogEventActionChangeStickerSet {
	return &TLChannelAdminLogEventActionChangeStickerSet{
		Data2: m,
	}
}

//  channelAdminLogEventActionTogglePreHistoryHidden#5f5c95f1 new_value:Bool = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionTogglePreHistoryHidden() *TLChannelAdminLogEventActionTogglePreHistoryHidden {
	return &TLChannelAdminLogEventActionTogglePreHistoryHidden{
		Data2: m,
	}
}

//  channelAdminLogEventActionDefaultBannedRights#2df5fc0a prev_banned_rights:ChatBannedRights new_banned_rights:ChatBannedRights = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionDefaultBannedRights() *TLChannelAdminLogEventActionDefaultBannedRights {
	return &TLChannelAdminLogEventActionDefaultBannedRights{
		Data2: m,
	}
}

//  channelAdminLogEventActionStopPoll#8f079643 message:Message = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionStopPoll() *TLChannelAdminLogEventActionStopPoll {
	return &TLChannelAdminLogEventActionStopPoll{
		Data2: m,
	}
}

//  channelAdminLogEventActionChangeLinkedChat#a26f881b prev_value:int new_value:int = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangeLinkedChat() *TLChannelAdminLogEventActionChangeLinkedChat {
	return &TLChannelAdminLogEventActionChangeLinkedChat{
		Data2: m,
	}
}

//  channelAdminLogEventActionChangeLocation#e6b76ae prev_value:ChannelLocation new_value:ChannelLocation = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangeLocation() *TLChannelAdminLogEventActionChangeLocation {
	return &TLChannelAdminLogEventActionChangeLocation{
		Data2: m,
	}
}

//  channelAdminLogEventActionToggleSlowMode#53909779 prev_value:int new_value:int = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionToggleSlowMode() *TLChannelAdminLogEventActionToggleSlowMode {
	return &TLChannelAdminLogEventActionToggleSlowMode{
		Data2: m,
	}
}

//  channelAdminLogEventActionStartGroupCall#23209745 call:InputGroupCall = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionStartGroupCall() *TLChannelAdminLogEventActionStartGroupCall {
	return &TLChannelAdminLogEventActionStartGroupCall{
		Data2: m,
	}
}

//  channelAdminLogEventActionDiscardGroupCall#db9f9140 call:InputGroupCall = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionDiscardGroupCall() *TLChannelAdminLogEventActionDiscardGroupCall {
	return &TLChannelAdminLogEventActionDiscardGroupCall{
		Data2: m,
	}
}

//  channelAdminLogEventActionParticipantMute#f92424d2 participant:GroupCallParticipant = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantMute() *TLChannelAdminLogEventActionParticipantMute {
	return &TLChannelAdminLogEventActionParticipantMute{
		Data2: m,
	}
}

//  channelAdminLogEventActionParticipantUnmute#e64429c0 participant:GroupCallParticipant = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantUnmute() *TLChannelAdminLogEventActionParticipantUnmute {
	return &TLChannelAdminLogEventActionParticipantUnmute{
		Data2: m,
	}
}

//  channelAdminLogEventActionToggleGroupCallSetting#56d6a247 join_muted:Bool = ChannelAdminLogEventAction;
//
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionToggleGroupCallSetting() *TLChannelAdminLogEventActionToggleGroupCallSetting {
	return &TLChannelAdminLogEventActionToggleGroupCallSetting{
		Data2: m,
	}
}

//  channelAdminLogEventActionChangeTitle#e6dfb825 prev_value:string new_value:string = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionChangeTitle) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionChangeTitle
	return m.Data2
}

//  channelAdminLogEventActionChangeAbout#55188a2e prev_value:string new_value:string = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionChangeAbout) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionChangeAbout
	return m.Data2
}

//  channelAdminLogEventActionChangeUsername#6a4afc38 prev_value:string new_value:string = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionChangeUsername) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionChangeUsername
	return m.Data2
}

//  channelAdminLogEventActionChangePhoto#b82f55c3 prev_photo:ChatPhoto new_photo:ChatPhoto = ChannelAdminLogEventAction;
//  channelAdminLogEventActionChangePhoto#434bd2af prev_photo:Photo new_photo:Photo = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionChangePhoto) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionChangePhoto
	return m.Data2
}

//  channelAdminLogEventActionToggleInvites#1b7907ae new_value:Bool = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionToggleInvites) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionToggleInvites
	return m.Data2
}

//  channelAdminLogEventActionToggleSignatures#26ae0971 new_value:Bool = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionToggleSignatures) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionToggleSignatures
	return m.Data2
}

//  channelAdminLogEventActionUpdatePinned#e9e82c18 message:Message = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionUpdatePinned) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionUpdatePinned
	return m.Data2
}

//  channelAdminLogEventActionEditMessage#709b2405 prev_message:Message new_message:Message = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionEditMessage) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionEditMessage
	return m.Data2
}

//  channelAdminLogEventActionDeleteMessage#42e047bb message:Message = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionDeleteMessage) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionDeleteMessage
	return m.Data2
}

//  channelAdminLogEventActionParticipantJoin#183040d3 = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionParticipantJoin) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionParticipantJoin
	return m.Data2
}

//  channelAdminLogEventActionParticipantLeave#f89777f2 = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionParticipantLeave) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionParticipantLeave
	return m.Data2
}

//  channelAdminLogEventActionParticipantInvite#e31c34d8 participant:ChannelParticipant = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionParticipantInvite) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionParticipantInvite
	return m.Data2
}

//  channelAdminLogEventActionParticipantToggleBan#e6d83d7e prev_participant:ChannelParticipant new_participant:ChannelParticipant = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionParticipantToggleBan) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionParticipantToggleBan
	return m.Data2
}

//  channelAdminLogEventActionParticipantToggleAdmin#d5676710 prev_participant:ChannelParticipant new_participant:ChannelParticipant = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionParticipantToggleAdmin
	return m.Data2
}

//  channelAdminLogEventActionChangeStickerSet#b1c3caa7 prev_stickerset:InputStickerSet new_stickerset:InputStickerSet = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionChangeStickerSet) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionChangeStickerSet
	return m.Data2
}

//  channelAdminLogEventActionTogglePreHistoryHidden#5f5c95f1 new_value:Bool = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionTogglePreHistoryHidden) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionTogglePreHistoryHidden
	return m.Data2
}

//  channelAdminLogEventActionDefaultBannedRights#2df5fc0a prev_banned_rights:ChatBannedRights new_banned_rights:ChatBannedRights = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionDefaultBannedRights) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionDefaultBannedRights
	return m.Data2
}

//  channelAdminLogEventActionStopPoll#8f079643 message:Message = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionStopPoll) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionStopPoll
	return m.Data2
}

//  channelAdminLogEventActionChangeLinkedChat#a26f881b prev_value:int new_value:int = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionChangeLinkedChat) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionChangeLinkedChat
	return m.Data2
}

//  channelAdminLogEventActionChangeLocation#e6b76ae prev_value:ChannelLocation new_value:ChannelLocation = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionChangeLocation) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionChangeLocation
	return m.Data2
}

//  channelAdminLogEventActionToggleSlowMode#53909779 prev_value:int new_value:int = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionToggleSlowMode) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionToggleSlowMode
	return m.Data2
}

//  channelAdminLogEventActionStartGroupCall#23209745 call:InputGroupCall = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionStartGroupCall) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionStartGroupCall
	return m.Data2
}

//  channelAdminLogEventActionDiscardGroupCall#db9f9140 call:InputGroupCall = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionDiscardGroupCall) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionDiscardGroupCall
	return m.Data2
}

//  channelAdminLogEventActionParticipantMute#f92424d2 participant:GroupCallParticipant = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionParticipantMute) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionParticipantMute
	return m.Data2
}

//  channelAdminLogEventActionParticipantUnmute#e64429c0 participant:GroupCallParticipant = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionParticipantUnmute) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionParticipantUnmute
	return m.Data2
}

//  channelAdminLogEventActionToggleGroupCallSetting#56d6a247 join_muted:Bool = ChannelAdminLogEventAction;
//
func (m *TLChannelAdminLogEventActionToggleGroupCallSetting) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventActionToggleGroupCallSetting
	return m.Data2
}

//  + TL_ChannelAdminLogEventsFilter
//

//  channelAdminLogEventsFilter#ea107ae4 flags:# join:flags.0?true leave:flags.1?true invite:flags.2?true ban:flags.3?true unban:flags.4?true kick:flags.5?true unkick:flags.6?true promote:flags.7?true demote:flags.8?true info:flags.9?true settings:flags.10?true pinned:flags.11?true edit:flags.12?true delete:flags.13?true group_call:flags.14?true = ChannelAdminLogEventsFilter;
//
func (m *ChannelAdminLogEventsFilter) To_ChannelAdminLogEventsFilter() *TLChannelAdminLogEventsFilter {
	return &TLChannelAdminLogEventsFilter{
		Data2: m,
	}
}

//  channelAdminLogEventsFilter#ea107ae4 flags:# join:flags.0?true leave:flags.1?true invite:flags.2?true ban:flags.3?true unban:flags.4?true kick:flags.5?true unkick:flags.6?true promote:flags.7?true demote:flags.8?true info:flags.9?true settings:flags.10?true pinned:flags.11?true edit:flags.12?true delete:flags.13?true group_call:flags.14?true = ChannelAdminLogEventsFilter;
//
func (m *TLChannelAdminLogEventsFilter) To_ChannelAdminLogEventsFilter() *ChannelAdminLogEventsFilter {
	m.Data2.Cmd = Cmd_ChannelAdminLogEventsFilter
	return m.Data2
}

//  + TL_ChannelAdminRights
//

//  channelAdminRights#5d7ceba5 flags:# change_info:flags.0?true post_messages:flags.1?true edit_messages:flags.2?true delete_messages:flags.3?true ban_users:flags.4?true invite_users:flags.5?true invite_link:flags.6?true pin_messages:flags.7?true add_admins:flags.9?true = ChannelAdminRights;
//
func (m *ChannelAdminRights) To_ChannelAdminRights() *TLChannelAdminRights {
	return &TLChannelAdminRights{
		Data2: m,
	}
}

//  channelAdminRights#5d7ceba5 flags:# change_info:flags.0?true post_messages:flags.1?true edit_messages:flags.2?true delete_messages:flags.3?true ban_users:flags.4?true invite_users:flags.5?true invite_link:flags.6?true pin_messages:flags.7?true add_admins:flags.9?true = ChannelAdminRights;
//
func (m *TLChannelAdminRights) To_ChannelAdminRights() *ChannelAdminRights {
	m.Data2.Cmd = Cmd_ChannelAdminRights
	return m.Data2
}

//  + TL_ChannelBannedRights
//

//  channelBannedRights#58cf4249 flags:# view_messages:flags.0?true send_messages:flags.1?true send_media:flags.2?true send_stickers:flags.3?true send_gifs:flags.4?true send_games:flags.5?true send_inline:flags.6?true embed_links:flags.7?true until_date:int = ChannelBannedRights;
//
func (m *ChannelBannedRights) To_ChannelBannedRights() *TLChannelBannedRights {
	return &TLChannelBannedRights{
		Data2: m,
	}
}

//  channelBannedRights#58cf4249 flags:# view_messages:flags.0?true send_messages:flags.1?true send_media:flags.2?true send_stickers:flags.3?true send_gifs:flags.4?true send_games:flags.5?true send_inline:flags.6?true embed_links:flags.7?true until_date:int = ChannelBannedRights;
//
func (m *TLChannelBannedRights) To_ChannelBannedRights() *ChannelBannedRights {
	m.Data2.Cmd = Cmd_ChannelBannedRights
	return m.Data2
}

//  + TL_ChannelLocationEmpty
//  + TL_ChannelLocation
//

//  channelLocationEmpty#bfb5ad8b = ChannelLocation;
//
func (m *ChannelLocation) To_ChannelLocationEmpty() *TLChannelLocationEmpty {
	return &TLChannelLocationEmpty{
		Data2: m,
	}
}

//  channelLocation#209b82db geo_point:GeoPoint address:string = ChannelLocation;
//
func (m *ChannelLocation) To_ChannelLocation() *TLChannelLocation {
	return &TLChannelLocation{
		Data2: m,
	}
}

//  channelLocationEmpty#bfb5ad8b = ChannelLocation;
//
func (m *TLChannelLocationEmpty) To_ChannelLocation() *ChannelLocation {
	m.Data2.Cmd = Cmd_ChannelLocationEmpty
	return m.Data2
}

//  channelLocation#209b82db geo_point:GeoPoint address:string = ChannelLocation;
//
func (m *TLChannelLocation) To_ChannelLocation() *ChannelLocation {
	m.Data2.Cmd = Cmd_ChannelLocation
	return m.Data2
}

//  + TL_ChannelMessagesFilterEmpty
//  + TL_ChannelMessagesFilter
//  + TL_ChannelMessagesFilterCollapsed
//

//  channelMessagesFilterEmpty#94d42ee7 = ChannelMessagesFilter;
//
func (m *ChannelMessagesFilter) To_ChannelMessagesFilterEmpty() *TLChannelMessagesFilterEmpty {
	return &TLChannelMessagesFilterEmpty{
		Data2: m,
	}
}

//  channelMessagesFilter#cd77d957 flags:# exclude_new_messages:flags.1?true ranges:Vector<MessageRange> = ChannelMessagesFilter;
//
func (m *ChannelMessagesFilter) To_ChannelMessagesFilter() *TLChannelMessagesFilter {
	return &TLChannelMessagesFilter{
		Data2: m,
	}
}

//  channelMessagesFilterCollapsed#fa01232e = ChannelMessagesFilter;
//
func (m *ChannelMessagesFilter) To_ChannelMessagesFilterCollapsed() *TLChannelMessagesFilterCollapsed {
	return &TLChannelMessagesFilterCollapsed{
		Data2: m,
	}
}

//  channelMessagesFilterEmpty#94d42ee7 = ChannelMessagesFilter;
//
func (m *TLChannelMessagesFilterEmpty) To_ChannelMessagesFilter() *ChannelMessagesFilter {
	m.Data2.Cmd = Cmd_ChannelMessagesFilterEmpty
	return m.Data2
}

//  channelMessagesFilter#cd77d957 flags:# exclude_new_messages:flags.1?true ranges:Vector<MessageRange> = ChannelMessagesFilter;
//
func (m *TLChannelMessagesFilter) To_ChannelMessagesFilter() *ChannelMessagesFilter {
	m.Data2.Cmd = Cmd_ChannelMessagesFilter
	return m.Data2
}

//  channelMessagesFilterCollapsed#fa01232e = ChannelMessagesFilter;
//
func (m *TLChannelMessagesFilterCollapsed) To_ChannelMessagesFilter() *ChannelMessagesFilter {
	m.Data2.Cmd = Cmd_ChannelMessagesFilterCollapsed
	return m.Data2
}

//  + TL_ChannelParticipant
//  + TL_ChannelParticipantSelf
//  + TL_ChannelParticipantCreator
//  + TL_ChannelParticipantAdmin
//  + TL_ChannelParticipantBanned
//  + TL_ChannelParticipantModerator
//  + TL_ChannelParticipantEditor
//  + TL_ChannelParticipantKicked
//  + TL_ChannelParticipantLeft
//

//  channelParticipant#15ebac1d user_id:int date:int = ChannelParticipant;
//
func (m *ChannelParticipant) To_ChannelParticipant() *TLChannelParticipant {
	return &TLChannelParticipant{
		Data2: m,
	}
}

//  channelParticipantSelf#a3289a6d user_id:int inviter_id:int date:int = ChannelParticipant;
//
func (m *ChannelParticipant) To_ChannelParticipantSelf() *TLChannelParticipantSelf {
	return &TLChannelParticipantSelf{
		Data2: m,
	}
}

//  channelParticipantCreator#e3e2e1f9 user_id:int = ChannelParticipant;
//  channelParticipantCreator#808d15a4 flags:# user_id:int rank:flags.0?string = ChannelParticipant;
//  channelParticipantCreator#447dca4b flags:# user_id:int admin_rights:ChatAdminRights rank:flags.0?string = ChannelParticipant;
//
func (m *ChannelParticipant) To_ChannelParticipantCreator() *TLChannelParticipantCreator {
	return &TLChannelParticipantCreator{
		Data2: m,
	}
}

//  channelParticipantAdmin#a82fa898 flags:# can_edit:flags.0?true user_id:int inviter_id:int promoted_by:int date:int admin_rights:ChannelAdminRights = ChannelParticipant;
//  channelParticipantAdmin#5daa6e23 flags:# can_edit:flags.0?true self:flags.1?true user_id:int inviter_id:flags.1?int promoted_by:int date:int admin_rights:ChatAdminRights = ChannelParticipant;
//  channelParticipantAdmin#ccbebbaf flags:# can_edit:flags.0?true self:flags.1?true user_id:int inviter_id:flags.1?int promoted_by:int date:int admin_rights:ChatAdminRights rank:flags.2?string = ChannelParticipant;
//
func (m *ChannelParticipant) To_ChannelParticipantAdmin() *TLChannelParticipantAdmin {
	return &TLChannelParticipantAdmin{
		Data2: m,
	}
}

//  channelParticipantBanned#222c1886 flags:# left:flags.0?true user_id:int kicked_by:int date:int banned_rights:ChannelBannedRights = ChannelParticipant;
//  channelParticipantBanned#1c0facaf flags:# left:flags.0?true user_id:int kicked_by:int date:int banned_rights:ChatBannedRights = ChannelParticipant;
//
func (m *ChannelParticipant) To_ChannelParticipantBanned() *TLChannelParticipantBanned {
	return &TLChannelParticipantBanned{
		Data2: m,
	}
}

//  channelParticipantModerator#91057fef user_id:int inviter_id:int date:int = ChannelParticipant;
//
func (m *ChannelParticipant) To_ChannelParticipantModerator() *TLChannelParticipantModerator {
	return &TLChannelParticipantModerator{
		Data2: m,
	}
}

//  channelParticipantEditor#98192d61 user_id:int inviter_id:int date:int = ChannelParticipant;
//
func (m *ChannelParticipant) To_ChannelParticipantEditor() *TLChannelParticipantEditor {
	return &TLChannelParticipantEditor{
		Data2: m,
	}
}

//  channelParticipantKicked#8cc5e69a user_id:int kicked_by:int date:int = ChannelParticipant;
//
func (m *ChannelParticipant) To_ChannelParticipantKicked() *TLChannelParticipantKicked {
	return &TLChannelParticipantKicked{
		Data2: m,
	}
}

//  channelParticipantLeft#c3c6796b user_id:int = ChannelParticipant;
//
func (m *ChannelParticipant) To_ChannelParticipantLeft() *TLChannelParticipantLeft {
	return &TLChannelParticipantLeft{
		Data2: m,
	}
}

//  channelParticipant#15ebac1d user_id:int date:int = ChannelParticipant;
//
func (m *TLChannelParticipant) To_ChannelParticipant() *ChannelParticipant {
	m.Data2.Cmd = Cmd_ChannelParticipant
	return m.Data2
}

//  channelParticipantSelf#a3289a6d user_id:int inviter_id:int date:int = ChannelParticipant;
//
func (m *TLChannelParticipantSelf) To_ChannelParticipant() *ChannelParticipant {
	m.Data2.Cmd = Cmd_ChannelParticipantSelf
	return m.Data2
}

//  channelParticipantCreator#e3e2e1f9 user_id:int = ChannelParticipant;
//  channelParticipantCreator#808d15a4 flags:# user_id:int rank:flags.0?string = ChannelParticipant;
//  channelParticipantCreator#447dca4b flags:# user_id:int admin_rights:ChatAdminRights rank:flags.0?string = ChannelParticipant;
//
func (m *TLChannelParticipantCreator) To_ChannelParticipant() *ChannelParticipant {
	m.Data2.Cmd = Cmd_ChannelParticipantCreator
	return m.Data2
}

//  channelParticipantAdmin#a82fa898 flags:# can_edit:flags.0?true user_id:int inviter_id:int promoted_by:int date:int admin_rights:ChannelAdminRights = ChannelParticipant;
//  channelParticipantAdmin#5daa6e23 flags:# can_edit:flags.0?true self:flags.1?true user_id:int inviter_id:flags.1?int promoted_by:int date:int admin_rights:ChatAdminRights = ChannelParticipant;
//  channelParticipantAdmin#ccbebbaf flags:# can_edit:flags.0?true self:flags.1?true user_id:int inviter_id:flags.1?int promoted_by:int date:int admin_rights:ChatAdminRights rank:flags.2?string = ChannelParticipant;
//
func (m *TLChannelParticipantAdmin) To_ChannelParticipant() *ChannelParticipant {
	m.Data2.Cmd = Cmd_ChannelParticipantAdmin
	return m.Data2
}

//  channelParticipantBanned#222c1886 flags:# left:flags.0?true user_id:int kicked_by:int date:int banned_rights:ChannelBannedRights = ChannelParticipant;
//  channelParticipantBanned#1c0facaf flags:# left:flags.0?true user_id:int kicked_by:int date:int banned_rights:ChatBannedRights = ChannelParticipant;
//
func (m *TLChannelParticipantBanned) To_ChannelParticipant() *ChannelParticipant {
	m.Data2.Cmd = Cmd_ChannelParticipantBanned
	return m.Data2
}

//  channelParticipantModerator#91057fef user_id:int inviter_id:int date:int = ChannelParticipant;
//
func (m *TLChannelParticipantModerator) To_ChannelParticipant() *ChannelParticipant {
	m.Data2.Cmd = Cmd_ChannelParticipantModerator
	return m.Data2
}

//  channelParticipantEditor#98192d61 user_id:int inviter_id:int date:int = ChannelParticipant;
//
func (m *TLChannelParticipantEditor) To_ChannelParticipant() *ChannelParticipant {
	m.Data2.Cmd = Cmd_ChannelParticipantEditor
	return m.Data2
}

//  channelParticipantKicked#8cc5e69a user_id:int kicked_by:int date:int = ChannelParticipant;
//
func (m *TLChannelParticipantKicked) To_ChannelParticipant() *ChannelParticipant {
	m.Data2.Cmd = Cmd_ChannelParticipantKicked
	return m.Data2
}

//  channelParticipantLeft#c3c6796b user_id:int = ChannelParticipant;
//
func (m *TLChannelParticipantLeft) To_ChannelParticipant() *ChannelParticipant {
	m.Data2.Cmd = Cmd_ChannelParticipantLeft
	return m.Data2
}

//  + TL_ChannelRoleEmpty
//  + TL_ChannelRoleModerator
//  + TL_ChannelRoleEditor
//

//  channelRoleEmpty#b285a0c6 = ChannelParticipantRole;
//
func (m *ChannelParticipantRole) To_ChannelRoleEmpty() *TLChannelRoleEmpty {
	return &TLChannelRoleEmpty{
		Data2: m,
	}
}

//  channelRoleModerator#9618d975 = ChannelParticipantRole;
//
func (m *ChannelParticipantRole) To_ChannelRoleModerator() *TLChannelRoleModerator {
	return &TLChannelRoleModerator{
		Data2: m,
	}
}

//  channelRoleEditor#820bfe8c = ChannelParticipantRole;
//
func (m *ChannelParticipantRole) To_ChannelRoleEditor() *TLChannelRoleEditor {
	return &TLChannelRoleEditor{
		Data2: m,
	}
}

//  channelRoleEmpty#b285a0c6 = ChannelParticipantRole;
//
func (m *TLChannelRoleEmpty) To_ChannelParticipantRole() *ChannelParticipantRole {
	m.Data2.Cmd = Cmd_ChannelRoleEmpty
	return m.Data2
}

//  channelRoleModerator#9618d975 = ChannelParticipantRole;
//
func (m *TLChannelRoleModerator) To_ChannelParticipantRole() *ChannelParticipantRole {
	m.Data2.Cmd = Cmd_ChannelRoleModerator
	return m.Data2
}

//  channelRoleEditor#820bfe8c = ChannelParticipantRole;
//
func (m *TLChannelRoleEditor) To_ChannelParticipantRole() *ChannelParticipantRole {
	m.Data2.Cmd = Cmd_ChannelRoleEditor
	return m.Data2
}

//  + TL_ChannelParticipantsRecent
//  + TL_ChannelParticipantsAdmins
//  + TL_ChannelParticipantsKicked
//  + TL_ChannelParticipantsBots
//  + TL_ChannelParticipantsBanned
//  + TL_ChannelParticipantsSearch
//  + TL_ChannelParticipantsContacts
//  + TL_ChannelParticipantsMentions
//

//  channelParticipantsRecent#de3f3c79 = ChannelParticipantsFilter;
//
func (m *ChannelParticipantsFilter) To_ChannelParticipantsRecent() *TLChannelParticipantsRecent {
	return &TLChannelParticipantsRecent{
		Data2: m,
	}
}

//  channelParticipantsAdmins#b4608969 = ChannelParticipantsFilter;
//
func (m *ChannelParticipantsFilter) To_ChannelParticipantsAdmins() *TLChannelParticipantsAdmins {
	return &TLChannelParticipantsAdmins{
		Data2: m,
	}
}

//  channelParticipantsKicked#a3b54985 q:string = ChannelParticipantsFilter;
//  channelParticipantsKicked#3c37bb7a = ChannelParticipantsFilter;
//
func (m *ChannelParticipantsFilter) To_ChannelParticipantsKicked() *TLChannelParticipantsKicked {
	return &TLChannelParticipantsKicked{
		Data2: m,
	}
}

//  channelParticipantsBots#b0d1865b = ChannelParticipantsFilter;
//
func (m *ChannelParticipantsFilter) To_ChannelParticipantsBots() *TLChannelParticipantsBots {
	return &TLChannelParticipantsBots{
		Data2: m,
	}
}

//  channelParticipantsBanned#1427a5e1 q:string = ChannelParticipantsFilter;
//
func (m *ChannelParticipantsFilter) To_ChannelParticipantsBanned() *TLChannelParticipantsBanned {
	return &TLChannelParticipantsBanned{
		Data2: m,
	}
}

//  channelParticipantsSearch#656ac4b q:string = ChannelParticipantsFilter;
//
func (m *ChannelParticipantsFilter) To_ChannelParticipantsSearch() *TLChannelParticipantsSearch {
	return &TLChannelParticipantsSearch{
		Data2: m,
	}
}

//  channelParticipantsContacts#bb6ae88d q:string = ChannelParticipantsFilter;
//
func (m *ChannelParticipantsFilter) To_ChannelParticipantsContacts() *TLChannelParticipantsContacts {
	return &TLChannelParticipantsContacts{
		Data2: m,
	}
}

//  channelParticipantsMentions#e04b5ceb flags:# q:flags.0?string top_msg_id:flags.1?int = ChannelParticipantsFilter;
//
func (m *ChannelParticipantsFilter) To_ChannelParticipantsMentions() *TLChannelParticipantsMentions {
	return &TLChannelParticipantsMentions{
		Data2: m,
	}
}

//  channelParticipantsRecent#de3f3c79 = ChannelParticipantsFilter;
//
func (m *TLChannelParticipantsRecent) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	m.Data2.Cmd = Cmd_ChannelParticipantsRecent
	return m.Data2
}

//  channelParticipantsAdmins#b4608969 = ChannelParticipantsFilter;
//
func (m *TLChannelParticipantsAdmins) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	m.Data2.Cmd = Cmd_ChannelParticipantsAdmins
	return m.Data2
}

//  channelParticipantsKicked#a3b54985 q:string = ChannelParticipantsFilter;
//  channelParticipantsKicked#3c37bb7a = ChannelParticipantsFilter;
//
func (m *TLChannelParticipantsKicked) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	m.Data2.Cmd = Cmd_ChannelParticipantsKicked
	return m.Data2
}

//  channelParticipantsBots#b0d1865b = ChannelParticipantsFilter;
//
func (m *TLChannelParticipantsBots) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	m.Data2.Cmd = Cmd_ChannelParticipantsBots
	return m.Data2
}

//  channelParticipantsBanned#1427a5e1 q:string = ChannelParticipantsFilter;
//
func (m *TLChannelParticipantsBanned) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	m.Data2.Cmd = Cmd_ChannelParticipantsBanned
	return m.Data2
}

//  channelParticipantsSearch#656ac4b q:string = ChannelParticipantsFilter;
//
func (m *TLChannelParticipantsSearch) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	m.Data2.Cmd = Cmd_ChannelParticipantsSearch
	return m.Data2
}

//  channelParticipantsContacts#bb6ae88d q:string = ChannelParticipantsFilter;
//
func (m *TLChannelParticipantsContacts) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	m.Data2.Cmd = Cmd_ChannelParticipantsContacts
	return m.Data2
}

//  channelParticipantsMentions#e04b5ceb flags:# q:flags.0?string top_msg_id:flags.1?int = ChannelParticipantsFilter;
//
func (m *TLChannelParticipantsMentions) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	m.Data2.Cmd = Cmd_ChannelParticipantsMentions
	return m.Data2
}

//  + TL_ChannelsAdminLogResults
//

//  channels.adminLogResults#ed8af74d events:Vector<ChannelAdminLogEvent> chats:Vector<Chat> users:Vector<User> = channels.AdminLogResults;
//
func (m *Channels_AdminLogResults) To_ChannelsAdminLogResults() *TLChannelsAdminLogResults {
	return &TLChannelsAdminLogResults{
		Data2: m,
	}
}

//  channels.adminLogResults#ed8af74d events:Vector<ChannelAdminLogEvent> chats:Vector<Chat> users:Vector<User> = channels.AdminLogResults;
//
func (m *TLChannelsAdminLogResults) To_Channels_AdminLogResults() *Channels_AdminLogResults {
	m.Data2.Cmd = Cmd_ChannelsAdminLogResults
	return m.Data2
}

//  + TL_ChannelsChannelParticipant
//

//  channels.channelParticipant#d0d9b163 participant:ChannelParticipant users:Vector<User> = channels.ChannelParticipant;
//
func (m *Channels_ChannelParticipant) To_ChannelsChannelParticipant() *TLChannelsChannelParticipant {
	return &TLChannelsChannelParticipant{
		Data2: m,
	}
}

//  channels.channelParticipant#d0d9b163 participant:ChannelParticipant users:Vector<User> = channels.ChannelParticipant;
//
func (m *TLChannelsChannelParticipant) To_Channels_ChannelParticipant() *Channels_ChannelParticipant {
	m.Data2.Cmd = Cmd_ChannelsChannelParticipant
	return m.Data2
}

//  + TL_ChannelsChannelParticipants
//  + TL_ChannelsChannelParticipantsNotModified
//

//  channels.channelParticipants#f56ee2a8 count:int participants:Vector<ChannelParticipant> users:Vector<User> = channels.ChannelParticipants;
//
func (m *Channels_ChannelParticipants) To_ChannelsChannelParticipants() *TLChannelsChannelParticipants {
	return &TLChannelsChannelParticipants{
		Data2: m,
	}
}

//  channels.channelParticipantsNotModified#f0173fe9 = channels.ChannelParticipants;
//
func (m *Channels_ChannelParticipants) To_ChannelsChannelParticipantsNotModified() *TLChannelsChannelParticipantsNotModified {
	return &TLChannelsChannelParticipantsNotModified{
		Data2: m,
	}
}

//  channels.channelParticipants#f56ee2a8 count:int participants:Vector<ChannelParticipant> users:Vector<User> = channels.ChannelParticipants;
//
func (m *TLChannelsChannelParticipants) To_Channels_ChannelParticipants() *Channels_ChannelParticipants {
	m.Data2.Cmd = Cmd_ChannelsChannelParticipants
	return m.Data2
}

//  channels.channelParticipantsNotModified#f0173fe9 = channels.ChannelParticipants;
//
func (m *TLChannelsChannelParticipantsNotModified) To_Channels_ChannelParticipants() *Channels_ChannelParticipants {
	m.Data2.Cmd = Cmd_ChannelsChannelParticipantsNotModified
	return m.Data2
}

//  + TL_ChatEmpty
//  + TL_Chat
//  + TL_ChatForbidden
//  + TL_Channel
//  + TL_ChannelForbidden
//

//  chatEmpty#9ba2d800 id:int = Chat;
//
func (m *Chat) To_ChatEmpty() *TLChatEmpty {
	return &TLChatEmpty{
		Data2: m,
	}
}

//  chat#d91cdd54 flags:# creator:flags.0?true kicked:flags.1?true left:flags.2?true admins_enabled:flags.3?true admin:flags.4?true deactivated:flags.5?true id:int title:string photo:ChatPhoto participants_count:int date:int version:int migrated_to:flags.6?InputChannel = Chat;
//  chat#3bda1bde flags:# creator:flags.0?true kicked:flags.1?true left:flags.2?true deactivated:flags.5?true call_active:flags.23?true call_not_empty:flags.24?true id:int title:string photo:ChatPhoto participants_count:int date:int version:int migrated_to:flags.6?InputChannel admin_rights:flags.14?ChatAdminRights default_banned_rights:flags.18?ChatBannedRights = Chat;
//
func (m *Chat) To_Chat() *TLChat {
	return &TLChat{
		Data2: m,
	}
}

//  chatForbidden#7328bdb id:int title:string = Chat;
//
func (m *Chat) To_ChatForbidden() *TLChatForbidden {
	return &TLChatForbidden{
		Data2: m,
	}
}

//  channel#cb44b1c flags:# creator:flags.0?true left:flags.2?true editor:flags.3?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true democracy:flags.10?true signatures:flags.11?true min:flags.12?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?string admin_rights:flags.14?ChannelAdminRights banned_rights:flags.15?ChannelBannedRights = Chat;
//  channel#a14dca52 flags:# creator:flags.0?true kicked:flags.1?true left:flags.2?true editor:flags.3?true moderator:flags.4?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true democracy:flags.10?true signatures:flags.11?true min:flags.12?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?string = Chat;
//  channel#c88974ac flags:# creator:flags.0?true left:flags.2?true editor:flags.3?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true democracy:flags.10?true signatures:flags.11?true min:flags.12?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?string admin_rights:flags.14?ChannelAdminRights banned_rights:flags.15?ChannelBannedRights participants_count:flags.17?int = Chat;
//  channel#4df30834 flags:# creator:flags.0?true left:flags.2?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true signatures:flags.11?true min:flags.12?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?string admin_rights:flags.14?ChatAdminRights banned_rights:flags.15?ChatBannedRights default_banned_rights:flags.18?ChatBannedRights participants_count:flags.17?int = Chat;
//  channel#d31a961e flags:# creator:flags.0?true left:flags.2?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true signatures:flags.11?true min:flags.12?true scam:flags.19?true has_link:flags.20?true has_geo:flags.21?true slowmode_enabled:flags.22?true call_active:flags.23?true call_not_empty:flags.24?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?Vector<RestrictionReason> admin_rights:flags.14?ChatAdminRights banned_rights:flags.15?ChatBannedRights default_banned_rights:flags.18?ChatBannedRights participants_count:flags.17?int = Chat;
//
func (m *Chat) To_Channel() *TLChannel {
	return &TLChannel{
		Data2: m,
	}
}

//  channelForbidden#289da732 flags:# broadcast:flags.5?true megagroup:flags.8?true id:int access_hash:long title:string until_date:flags.16?int = Chat;
//  channelForbidden#2d85832c id:int access_hash:long title:string = Chat;
//
func (m *Chat) To_ChannelForbidden() *TLChannelForbidden {
	return &TLChannelForbidden{
		Data2: m,
	}
}

//  chatEmpty#9ba2d800 id:int = Chat;
//
func (m *TLChatEmpty) To_Chat() *Chat {
	m.Data2.Cmd = Cmd_ChatEmpty
	return m.Data2
}

//  chat#d91cdd54 flags:# creator:flags.0?true kicked:flags.1?true left:flags.2?true admins_enabled:flags.3?true admin:flags.4?true deactivated:flags.5?true id:int title:string photo:ChatPhoto participants_count:int date:int version:int migrated_to:flags.6?InputChannel = Chat;
//  chat#3bda1bde flags:# creator:flags.0?true kicked:flags.1?true left:flags.2?true deactivated:flags.5?true call_active:flags.23?true call_not_empty:flags.24?true id:int title:string photo:ChatPhoto participants_count:int date:int version:int migrated_to:flags.6?InputChannel admin_rights:flags.14?ChatAdminRights default_banned_rights:flags.18?ChatBannedRights = Chat;
//
func (m *TLChat) To_Chat() *Chat {
	m.Data2.Cmd = Cmd_Chat
	return m.Data2
}

//  chatForbidden#7328bdb id:int title:string = Chat;
//
func (m *TLChatForbidden) To_Chat() *Chat {
	m.Data2.Cmd = Cmd_ChatForbidden
	return m.Data2
}

//  channel#cb44b1c flags:# creator:flags.0?true left:flags.2?true editor:flags.3?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true democracy:flags.10?true signatures:flags.11?true min:flags.12?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?string admin_rights:flags.14?ChannelAdminRights banned_rights:flags.15?ChannelBannedRights = Chat;
//  channel#a14dca52 flags:# creator:flags.0?true kicked:flags.1?true left:flags.2?true editor:flags.3?true moderator:flags.4?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true democracy:flags.10?true signatures:flags.11?true min:flags.12?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?string = Chat;
//  channel#c88974ac flags:# creator:flags.0?true left:flags.2?true editor:flags.3?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true democracy:flags.10?true signatures:flags.11?true min:flags.12?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?string admin_rights:flags.14?ChannelAdminRights banned_rights:flags.15?ChannelBannedRights participants_count:flags.17?int = Chat;
//  channel#4df30834 flags:# creator:flags.0?true left:flags.2?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true signatures:flags.11?true min:flags.12?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?string admin_rights:flags.14?ChatAdminRights banned_rights:flags.15?ChatBannedRights default_banned_rights:flags.18?ChatBannedRights participants_count:flags.17?int = Chat;
//  channel#d31a961e flags:# creator:flags.0?true left:flags.2?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true signatures:flags.11?true min:flags.12?true scam:flags.19?true has_link:flags.20?true has_geo:flags.21?true slowmode_enabled:flags.22?true call_active:flags.23?true call_not_empty:flags.24?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?Vector<RestrictionReason> admin_rights:flags.14?ChatAdminRights banned_rights:flags.15?ChatBannedRights default_banned_rights:flags.18?ChatBannedRights participants_count:flags.17?int = Chat;
//
func (m *TLChannel) To_Chat() *Chat {
	m.Data2.Cmd = Cmd_Channel
	return m.Data2
}

//  channelForbidden#289da732 flags:# broadcast:flags.5?true megagroup:flags.8?true id:int access_hash:long title:string until_date:flags.16?int = Chat;
//  channelForbidden#2d85832c id:int access_hash:long title:string = Chat;
//
func (m *TLChannelForbidden) To_Chat() *Chat {
	m.Data2.Cmd = Cmd_ChannelForbidden
	return m.Data2
}

//  + TL_ChatAdminRights
//

//  chatAdminRights#5fb224d5 flags:# change_info:flags.0?true post_messages:flags.1?true edit_messages:flags.2?true delete_messages:flags.3?true ban_users:flags.4?true invite_users:flags.5?true pin_messages:flags.7?true add_admins:flags.9?true anonymous:flags.10?true manage_call:flags.11?true = ChatAdminRights;
//
func (m *ChatAdminRights) To_ChatAdminRights() *TLChatAdminRights {
	return &TLChatAdminRights{
		Data2: m,
	}
}

//  chatAdminRights#5fb224d5 flags:# change_info:flags.0?true post_messages:flags.1?true edit_messages:flags.2?true delete_messages:flags.3?true ban_users:flags.4?true invite_users:flags.5?true pin_messages:flags.7?true add_admins:flags.9?true anonymous:flags.10?true manage_call:flags.11?true = ChatAdminRights;
//
func (m *TLChatAdminRights) To_ChatAdminRights() *ChatAdminRights {
	m.Data2.Cmd = Cmd_ChatAdminRights
	return m.Data2
}

//  + TL_ChatBannedRights
//

//  chatBannedRights#9f120418 flags:# view_messages:flags.0?true send_messages:flags.1?true send_media:flags.2?true send_stickers:flags.3?true send_gifs:flags.4?true send_games:flags.5?true send_inline:flags.6?true embed_links:flags.7?true send_polls:flags.8?true change_info:flags.10?true invite_users:flags.15?true pin_messages:flags.17?true until_date:int = ChatBannedRights;
//
func (m *ChatBannedRights) To_ChatBannedRights() *TLChatBannedRights {
	return &TLChatBannedRights{
		Data2: m,
	}
}

//  chatBannedRights#9f120418 flags:# view_messages:flags.0?true send_messages:flags.1?true send_media:flags.2?true send_stickers:flags.3?true send_gifs:flags.4?true send_games:flags.5?true send_inline:flags.6?true embed_links:flags.7?true send_polls:flags.8?true change_info:flags.10?true invite_users:flags.15?true pin_messages:flags.17?true until_date:int = ChatBannedRights;
//
func (m *TLChatBannedRights) To_ChatBannedRights() *ChatBannedRights {
	m.Data2.Cmd = Cmd_ChatBannedRights
	return m.Data2
}

//  + TL_ChatFull
//  + TL_ChannelFull
//

//  chatFull#2e02a614 id:int participants:ChatParticipants chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> = ChatFull;
//  chatFull#edd2a791 flags:# id:int participants:ChatParticipants chat_photo:flags.2?Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:flags.3?Vector<BotInfo> pinned_msg_id:flags.6?int = ChatFull;
//  chatFull#22a235da flags:# can_set_username:flags.7?true id:int about:string participants:ChatParticipants chat_photo:flags.2?Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:flags.3?Vector<BotInfo> pinned_msg_id:flags.6?int = ChatFull;
//  chatFull#1b7c9db3 flags:# can_set_username:flags.7?true id:int about:string participants:ChatParticipants chat_photo:flags.2?Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:flags.3?Vector<BotInfo> pinned_msg_id:flags.6?int folder_id:flags.11?int = ChatFull;
//  chatFull#dc8c181 flags:# can_set_username:flags.7?true has_scheduled:flags.8?true id:int about:string participants:ChatParticipants chat_photo:flags.2?Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:flags.3?Vector<BotInfo> pinned_msg_id:flags.6?int folder_id:flags.11?int call:flags.12?InputGroupCall = ChatFull;
//
func (m *ChatFull) To_ChatFull() *TLChatFull {
	return &TLChatFull{
		Data2: m,
	}
}

//  channelFull#17f45fcf flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet = ChatFull;
//  channelFull#97bee562 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int read_inbox_max_id:int unread_count:int unread_important_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int = ChatFull;
//  channelFull#76af5481 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int = ChatFull;
//  channelFull#1c87a71a flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_view_stats:flags.12?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int = ChatFull;
//  channelFull#3648977 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_view_stats:flags.12?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int pts:int = ChatFull;
//  channelFull#9882e516 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_view_stats:flags.12?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int linked_chat_id:flags.13?int pts:int = ChatFull;
//  channelFull#10916653 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_view_stats:flags.12?true can_set_location:flags.16?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int linked_chat_id:flags.14?int location:flags.15?ChannelLocation pts:int = ChatFull;
//  channelFull#2d895c74 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_view_stats:flags.12?true can_set_location:flags.16?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int linked_chat_id:flags.14?int location:flags.15?ChannelLocation slowmode_seconds:flags.17?int slowmode_next_send_date:flags.18?int pts:int = ChatFull;
//  channelFull#f0e6672a flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_set_location:flags.16?true has_scheduled:flags.19?true can_view_stats:flags.20?true blocked:flags.22?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int linked_chat_id:flags.14?int location:flags.15?ChannelLocation slowmode_seconds:flags.17?int slowmode_next_send_date:flags.18?int stats_dc:flags.12?int pts:int = ChatFull;
//  channelFull#ef3a6acd flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_set_location:flags.16?true has_scheduled:flags.19?true can_view_stats:flags.20?true blocked:flags.22?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int linked_chat_id:flags.14?int location:flags.15?ChannelLocation slowmode_seconds:flags.17?int slowmode_next_send_date:flags.18?int stats_dc:flags.12?int pts:int call:flags.21?InputGroupCall = ChatFull;
//
func (m *ChatFull) To_ChannelFull() *TLChannelFull {
	return &TLChannelFull{
		Data2: m,
	}
}

//  chatFull#2e02a614 id:int participants:ChatParticipants chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> = ChatFull;
//  chatFull#edd2a791 flags:# id:int participants:ChatParticipants chat_photo:flags.2?Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:flags.3?Vector<BotInfo> pinned_msg_id:flags.6?int = ChatFull;
//  chatFull#22a235da flags:# can_set_username:flags.7?true id:int about:string participants:ChatParticipants chat_photo:flags.2?Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:flags.3?Vector<BotInfo> pinned_msg_id:flags.6?int = ChatFull;
//  chatFull#1b7c9db3 flags:# can_set_username:flags.7?true id:int about:string participants:ChatParticipants chat_photo:flags.2?Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:flags.3?Vector<BotInfo> pinned_msg_id:flags.6?int folder_id:flags.11?int = ChatFull;
//  chatFull#dc8c181 flags:# can_set_username:flags.7?true has_scheduled:flags.8?true id:int about:string participants:ChatParticipants chat_photo:flags.2?Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:flags.3?Vector<BotInfo> pinned_msg_id:flags.6?int folder_id:flags.11?int call:flags.12?InputGroupCall = ChatFull;
//
func (m *TLChatFull) To_ChatFull() *ChatFull {
	m.Data2.Cmd = Cmd_ChatFull
	return m.Data2
}

//  channelFull#17f45fcf flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet = ChatFull;
//  channelFull#97bee562 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int read_inbox_max_id:int unread_count:int unread_important_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int = ChatFull;
//  channelFull#76af5481 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int = ChatFull;
//  channelFull#1c87a71a flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_view_stats:flags.12?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int = ChatFull;
//  channelFull#3648977 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_view_stats:flags.12?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int pts:int = ChatFull;
//  channelFull#9882e516 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_view_stats:flags.12?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int linked_chat_id:flags.13?int pts:int = ChatFull;
//  channelFull#10916653 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_view_stats:flags.12?true can_set_location:flags.16?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int linked_chat_id:flags.14?int location:flags.15?ChannelLocation pts:int = ChatFull;
//  channelFull#2d895c74 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_view_stats:flags.12?true can_set_location:flags.16?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int linked_chat_id:flags.14?int location:flags.15?ChannelLocation slowmode_seconds:flags.17?int slowmode_next_send_date:flags.18?int pts:int = ChatFull;
//  channelFull#f0e6672a flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_set_location:flags.16?true has_scheduled:flags.19?true can_view_stats:flags.20?true blocked:flags.22?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int linked_chat_id:flags.14?int location:flags.15?ChannelLocation slowmode_seconds:flags.17?int slowmode_next_send_date:flags.18?int stats_dc:flags.12?int pts:int = ChatFull;
//  channelFull#ef3a6acd flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_set_location:flags.16?true has_scheduled:flags.19?true can_view_stats:flags.20?true blocked:flags.22?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int linked_chat_id:flags.14?int location:flags.15?ChannelLocation slowmode_seconds:flags.17?int slowmode_next_send_date:flags.18?int stats_dc:flags.12?int pts:int call:flags.21?InputGroupCall = ChatFull;
//
func (m *TLChannelFull) To_ChatFull() *ChatFull {
	m.Data2.Cmd = Cmd_ChannelFull
	return m.Data2
}

//  + TL_ChatInviteAlready
//  + TL_ChatInvite
//  + TL_ChatInvitePeek
//

//  chatInviteAlready#5a686d7c chat:Chat = ChatInvite;
//
func (m *ChatInvite) To_ChatInviteAlready() *TLChatInviteAlready {
	return &TLChatInviteAlready{
		Data2: m,
	}
}

//  chatInvite#db74f558 flags:# channel:flags.0?true broadcast:flags.1?true public:flags.2?true megagroup:flags.3?true title:string photo:ChatPhoto participants_count:int participants:flags.4?Vector<User> = ChatInvite;
//  chatInvite#93e99b60 flags:# channel:flags.0?true broadcast:flags.1?true public:flags.2?true megagroup:flags.3?true title:string = ChatInvite;
//  chatInvite#dfc2f58e flags:# channel:flags.0?true broadcast:flags.1?true public:flags.2?true megagroup:flags.3?true title:string photo:Photo participants_count:int participants:flags.4?Vector<User> = ChatInvite;
//
func (m *ChatInvite) To_ChatInvite() *TLChatInvite {
	return &TLChatInvite{
		Data2: m,
	}
}

//  chatInvitePeek#61695cb0 chat:Chat expires:int = ChatInvite;
//
func (m *ChatInvite) To_ChatInvitePeek() *TLChatInvitePeek {
	return &TLChatInvitePeek{
		Data2: m,
	}
}

//  chatInviteAlready#5a686d7c chat:Chat = ChatInvite;
//
func (m *TLChatInviteAlready) To_ChatInvite() *ChatInvite {
	m.Data2.Cmd = Cmd_ChatInviteAlready
	return m.Data2
}

//  chatInvite#db74f558 flags:# channel:flags.0?true broadcast:flags.1?true public:flags.2?true megagroup:flags.3?true title:string photo:ChatPhoto participants_count:int participants:flags.4?Vector<User> = ChatInvite;
//  chatInvite#93e99b60 flags:# channel:flags.0?true broadcast:flags.1?true public:flags.2?true megagroup:flags.3?true title:string = ChatInvite;
//  chatInvite#dfc2f58e flags:# channel:flags.0?true broadcast:flags.1?true public:flags.2?true megagroup:flags.3?true title:string photo:Photo participants_count:int participants:flags.4?Vector<User> = ChatInvite;
//
func (m *TLChatInvite) To_ChatInvite() *ChatInvite {
	m.Data2.Cmd = Cmd_ChatInvite
	return m.Data2
}

//  chatInvitePeek#61695cb0 chat:Chat expires:int = ChatInvite;
//
func (m *TLChatInvitePeek) To_ChatInvite() *ChatInvite {
	m.Data2.Cmd = Cmd_ChatInvitePeek
	return m.Data2
}

//  + TL_ChatOnlines
//

//  chatOnlines#f041e250 onlines:int = ChatOnlines;
//
func (m *ChatOnlines) To_ChatOnlines() *TLChatOnlines {
	return &TLChatOnlines{
		Data2: m,
	}
}

//  chatOnlines#f041e250 onlines:int = ChatOnlines;
//
func (m *TLChatOnlines) To_ChatOnlines() *ChatOnlines {
	m.Data2.Cmd = Cmd_ChatOnlines
	return m.Data2
}

//  + TL_ChatParticipant
//  + TL_ChatParticipantCreator
//  + TL_ChatParticipantAdmin
//

//  chatParticipant#c8d7493e user_id:int inviter_id:int date:int = ChatParticipant;
//
func (m *ChatParticipant) To_ChatParticipant() *TLChatParticipant {
	return &TLChatParticipant{
		Data2: m,
	}
}

//  chatParticipantCreator#da13538a user_id:int = ChatParticipant;
//
func (m *ChatParticipant) To_ChatParticipantCreator() *TLChatParticipantCreator {
	return &TLChatParticipantCreator{
		Data2: m,
	}
}

//  chatParticipantAdmin#e2d6e436 user_id:int inviter_id:int date:int = ChatParticipant;
//
func (m *ChatParticipant) To_ChatParticipantAdmin() *TLChatParticipantAdmin {
	return &TLChatParticipantAdmin{
		Data2: m,
	}
}

//  chatParticipant#c8d7493e user_id:int inviter_id:int date:int = ChatParticipant;
//
func (m *TLChatParticipant) To_ChatParticipant() *ChatParticipant {
	m.Data2.Cmd = Cmd_ChatParticipant
	return m.Data2
}

//  chatParticipantCreator#da13538a user_id:int = ChatParticipant;
//
func (m *TLChatParticipantCreator) To_ChatParticipant() *ChatParticipant {
	m.Data2.Cmd = Cmd_ChatParticipantCreator
	return m.Data2
}

//  chatParticipantAdmin#e2d6e436 user_id:int inviter_id:int date:int = ChatParticipant;
//
func (m *TLChatParticipantAdmin) To_ChatParticipant() *ChatParticipant {
	m.Data2.Cmd = Cmd_ChatParticipantAdmin
	return m.Data2
}

//  + TL_ChatParticipantsForbidden
//  + TL_ChatParticipants
//

//  chatParticipantsForbidden#fc900c2b flags:# chat_id:int self_participant:flags.0?ChatParticipant = ChatParticipants;
//
func (m *ChatParticipants) To_ChatParticipantsForbidden() *TLChatParticipantsForbidden {
	return &TLChatParticipantsForbidden{
		Data2: m,
	}
}

//  chatParticipants#3f460fed chat_id:int participants:Vector<ChatParticipant> version:int = ChatParticipants;
//
func (m *ChatParticipants) To_ChatParticipants() *TLChatParticipants {
	return &TLChatParticipants{
		Data2: m,
	}
}

//  chatParticipantsForbidden#fc900c2b flags:# chat_id:int self_participant:flags.0?ChatParticipant = ChatParticipants;
//
func (m *TLChatParticipantsForbidden) To_ChatParticipants() *ChatParticipants {
	m.Data2.Cmd = Cmd_ChatParticipantsForbidden
	return m.Data2
}

//  chatParticipants#3f460fed chat_id:int participants:Vector<ChatParticipant> version:int = ChatParticipants;
//
func (m *TLChatParticipants) To_ChatParticipants() *ChatParticipants {
	m.Data2.Cmd = Cmd_ChatParticipants
	return m.Data2
}

//  + TL_ChatPhotoEmpty
//  + TL_ChatPhoto
//

//  chatPhotoEmpty#37c1011c = ChatPhoto;
//
func (m *ChatPhoto) To_ChatPhotoEmpty() *TLChatPhotoEmpty {
	return &TLChatPhotoEmpty{
		Data2: m,
	}
}

//  chatPhoto#6153276a photo_small:FileLocation photo_big:FileLocation = ChatPhoto;
//  chatPhoto#475cdbd5 photo_small:FileLocation photo_big:FileLocation dc_id:int = ChatPhoto;
//  chatPhoto#d20b9f3c flags:# has_video:flags.0?true photo_small:FileLocation photo_big:FileLocation dc_id:int = ChatPhoto;
//
func (m *ChatPhoto) To_ChatPhoto() *TLChatPhoto {
	return &TLChatPhoto{
		Data2: m,
	}
}

//  chatPhotoEmpty#37c1011c = ChatPhoto;
//
func (m *TLChatPhotoEmpty) To_ChatPhoto() *ChatPhoto {
	m.Data2.Cmd = Cmd_ChatPhotoEmpty
	return m.Data2
}

//  chatPhoto#6153276a photo_small:FileLocation photo_big:FileLocation = ChatPhoto;
//  chatPhoto#475cdbd5 photo_small:FileLocation photo_big:FileLocation dc_id:int = ChatPhoto;
//  chatPhoto#d20b9f3c flags:# has_video:flags.0?true photo_small:FileLocation photo_big:FileLocation dc_id:int = ChatPhoto;
//
func (m *TLChatPhoto) To_ChatPhoto() *ChatPhoto {
	m.Data2.Cmd = Cmd_ChatPhoto
	return m.Data2
}

//  + TL_Client_DHInnerData
//

//  client_DH_inner_data#6643b654 nonce:int128 server_nonce:int128 retry_id:long g_b:string = Client_DH_Inner_Data;
//
func (m *Client_DH_Inner_Data) To_Client_DHInnerData() *TLClient_DHInnerData {
	return &TLClient_DHInnerData{
		Data2: m,
	}
}

//  client_DH_inner_data#6643b654 nonce:int128 server_nonce:int128 retry_id:long g_b:string = Client_DH_Inner_Data;
//
func (m *TLClient_DHInnerData) To_Client_DH_Inner_Data() *Client_DH_Inner_Data {
	m.Data2.Cmd = Cmd_Client_DHInnerData
	return m.Data2
}

//  + TL_CodeSettings
//

//  codeSettings#302f59f3 flags:# allow_flashcall:flags.0?true current_number:flags.1?true app_hash_persistent:flags.2?true app_hash:flags.3?string = CodeSettings;
//  codeSettings#debebe83 flags:# allow_flashcall:flags.0?true current_number:flags.1?true allow_app_hash:flags.4?true = CodeSettings;
//
func (m *CodeSettings) To_CodeSettings() *TLCodeSettings {
	return &TLCodeSettings{
		Data2: m,
	}
}

//  codeSettings#302f59f3 flags:# allow_flashcall:flags.0?true current_number:flags.1?true app_hash_persistent:flags.2?true app_hash:flags.3?string = CodeSettings;
//  codeSettings#debebe83 flags:# allow_flashcall:flags.0?true current_number:flags.1?true allow_app_hash:flags.4?true = CodeSettings;
//
func (m *TLCodeSettings) To_CodeSettings() *CodeSettings {
	m.Data2.Cmd = Cmd_CodeSettings
	return m.Data2
}

//  + TL_Config
//

//  config#8df376a4 flags:# phonecalls_enabled:flags.1?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int chat_big_size:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string suggested_lang_code:flags.2?string lang_pack_version:flags.2?int disabled_features:Vector<DisabledFeature> = Config;
//  config#317ceef4 date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int chat_big_size:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int disabled_features:Vector<DisabledFeature> = Config;
//  config#3213dbba flags:# phonecalls_enabled:flags.1?true default_p2p_contacts:flags.3?true preload_featured_stickers:flags.4?true ignore_phone_entities:flags.5?true revoke_pm_inbox:flags.6?true blocked_mode:flags.8?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> dc_txt_domain_name:string chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int revoke_time_limit:int revoke_pm_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int channels_read_media_period:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string autoupdate_url_prefix:flags.7?string gif_search_username:flags.9?string venue_search_username:flags.10?string img_search_username:flags.11?string static_maps_provider:flags.12?string caption_length_max:int message_length_max:int webfile_dc_id:int suggested_lang_code:flags.2?string lang_pack_version:flags.2?int = Config;
//  config#e6ca25f6 flags:# phonecalls_enabled:flags.1?true default_p2p_contacts:flags.3?true preload_featured_stickers:flags.4?true ignore_phone_entities:flags.5?true revoke_pm_inbox:flags.6?true blocked_mode:flags.8?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> dc_txt_domain_name:string chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int revoke_time_limit:int revoke_pm_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int channels_read_media_period:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string autoupdate_url_prefix:flags.7?string gif_search_username:flags.9?string venue_search_username:flags.10?string img_search_username:flags.11?string static_maps_provider:flags.12?string caption_length_max:int message_length_max:int webfile_dc_id:int suggested_lang_code:flags.2?string lang_pack_version:flags.2?int base_lang_pack_version:flags.2?int = Config;
//  config#330b4067 flags:# phonecalls_enabled:flags.1?true default_p2p_contacts:flags.3?true preload_featured_stickers:flags.4?true ignore_phone_entities:flags.5?true revoke_pm_inbox:flags.6?true blocked_mode:flags.8?true pfs_enabled:flags.13?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> dc_txt_domain_name:string chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int revoke_time_limit:int revoke_pm_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int channels_read_media_period:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int pinned_infolder_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string autoupdate_url_prefix:flags.7?string gif_search_username:flags.9?string venue_search_username:flags.10?string img_search_username:flags.11?string static_maps_provider:flags.12?string caption_length_max:int message_length_max:int webfile_dc_id:int suggested_lang_code:flags.2?string lang_pack_version:flags.2?int base_lang_pack_version:flags.2?int = Config;
//
func (m *Config) To_Config() *TLConfig {
	return &TLConfig{
		Data2: m,
	}
}

//  config#8df376a4 flags:# phonecalls_enabled:flags.1?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int chat_big_size:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string suggested_lang_code:flags.2?string lang_pack_version:flags.2?int disabled_features:Vector<DisabledFeature> = Config;
//  config#317ceef4 date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int chat_big_size:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int disabled_features:Vector<DisabledFeature> = Config;
//  config#3213dbba flags:# phonecalls_enabled:flags.1?true default_p2p_contacts:flags.3?true preload_featured_stickers:flags.4?true ignore_phone_entities:flags.5?true revoke_pm_inbox:flags.6?true blocked_mode:flags.8?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> dc_txt_domain_name:string chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int revoke_time_limit:int revoke_pm_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int channels_read_media_period:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string autoupdate_url_prefix:flags.7?string gif_search_username:flags.9?string venue_search_username:flags.10?string img_search_username:flags.11?string static_maps_provider:flags.12?string caption_length_max:int message_length_max:int webfile_dc_id:int suggested_lang_code:flags.2?string lang_pack_version:flags.2?int = Config;
//  config#e6ca25f6 flags:# phonecalls_enabled:flags.1?true default_p2p_contacts:flags.3?true preload_featured_stickers:flags.4?true ignore_phone_entities:flags.5?true revoke_pm_inbox:flags.6?true blocked_mode:flags.8?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> dc_txt_domain_name:string chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int revoke_time_limit:int revoke_pm_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int channels_read_media_period:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string autoupdate_url_prefix:flags.7?string gif_search_username:flags.9?string venue_search_username:flags.10?string img_search_username:flags.11?string static_maps_provider:flags.12?string caption_length_max:int message_length_max:int webfile_dc_id:int suggested_lang_code:flags.2?string lang_pack_version:flags.2?int base_lang_pack_version:flags.2?int = Config;
//  config#330b4067 flags:# phonecalls_enabled:flags.1?true default_p2p_contacts:flags.3?true preload_featured_stickers:flags.4?true ignore_phone_entities:flags.5?true revoke_pm_inbox:flags.6?true blocked_mode:flags.8?true pfs_enabled:flags.13?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> dc_txt_domain_name:string chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int revoke_time_limit:int revoke_pm_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int channels_read_media_period:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int pinned_infolder_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string autoupdate_url_prefix:flags.7?string gif_search_username:flags.9?string venue_search_username:flags.10?string img_search_username:flags.11?string static_maps_provider:flags.12?string caption_length_max:int message_length_max:int webfile_dc_id:int suggested_lang_code:flags.2?string lang_pack_version:flags.2?int base_lang_pack_version:flags.2?int = Config;
//
func (m *TLConfig) To_Config() *Config {
	m.Data2.Cmd = Cmd_Config
	return m.Data2
}

//  + TL_Contact
//

//  contact#f911c994 user_id:int mutual:Bool = Contact;
//
func (m *Contact) To_Contact() *TLContact {
	return &TLContact{
		Data2: m,
	}
}

//  contact#f911c994 user_id:int mutual:Bool = Contact;
//
func (m *TLContact) To_Contact() *Contact {
	m.Data2.Cmd = Cmd_Contact
	return m.Data2
}

//  + TL_ContactBlocked
//

//  contactBlocked#561bc879 user_id:int date:int = ContactBlocked;
//
func (m *ContactBlocked) To_ContactBlocked() *TLContactBlocked {
	return &TLContactBlocked{
		Data2: m,
	}
}

//  contactBlocked#561bc879 user_id:int date:int = ContactBlocked;
//
func (m *TLContactBlocked) To_ContactBlocked() *ContactBlocked {
	m.Data2.Cmd = Cmd_ContactBlocked
	return m.Data2
}

//  + TL_ContactLinkUnknown
//  + TL_ContactLinkNone
//  + TL_ContactLinkHasPhone
//  + TL_ContactLinkContact
//

//  contactLinkUnknown#5f4f9247 = ContactLink;
//
func (m *ContactLink) To_ContactLinkUnknown() *TLContactLinkUnknown {
	return &TLContactLinkUnknown{
		Data2: m,
	}
}

//  contactLinkNone#feedd3ad = ContactLink;
//
func (m *ContactLink) To_ContactLinkNone() *TLContactLinkNone {
	return &TLContactLinkNone{
		Data2: m,
	}
}

//  contactLinkHasPhone#268f3f59 = ContactLink;
//
func (m *ContactLink) To_ContactLinkHasPhone() *TLContactLinkHasPhone {
	return &TLContactLinkHasPhone{
		Data2: m,
	}
}

//  contactLinkContact#d502c2d0 = ContactLink;
//
func (m *ContactLink) To_ContactLinkContact() *TLContactLinkContact {
	return &TLContactLinkContact{
		Data2: m,
	}
}

//  contactLinkUnknown#5f4f9247 = ContactLink;
//
func (m *TLContactLinkUnknown) To_ContactLink() *ContactLink {
	m.Data2.Cmd = Cmd_ContactLinkUnknown
	return m.Data2
}

//  contactLinkNone#feedd3ad = ContactLink;
//
func (m *TLContactLinkNone) To_ContactLink() *ContactLink {
	m.Data2.Cmd = Cmd_ContactLinkNone
	return m.Data2
}

//  contactLinkHasPhone#268f3f59 = ContactLink;
//
func (m *TLContactLinkHasPhone) To_ContactLink() *ContactLink {
	m.Data2.Cmd = Cmd_ContactLinkHasPhone
	return m.Data2
}

//  contactLinkContact#d502c2d0 = ContactLink;
//
func (m *TLContactLinkContact) To_ContactLink() *ContactLink {
	m.Data2.Cmd = Cmd_ContactLinkContact
	return m.Data2
}

//  + TL_ContactStatus
//

//  contactStatus#d3680c61 user_id:int status:UserStatus = ContactStatus;
//
func (m *ContactStatus) To_ContactStatus() *TLContactStatus {
	return &TLContactStatus{
		Data2: m,
	}
}

//  contactStatus#d3680c61 user_id:int status:UserStatus = ContactStatus;
//
func (m *TLContactStatus) To_ContactStatus() *ContactStatus {
	m.Data2.Cmd = Cmd_ContactStatus
	return m.Data2
}

//  + TL_ContactsBlocked
//  + TL_ContactsBlockedSlice
//

//  contacts.blocked#1c138d15 blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
//  contacts.blocked#ade1591 blocked:Vector<PeerBlocked> chats:Vector<Chat> users:Vector<User> = contacts.Blocked;
//
func (m *Contacts_Blocked) To_ContactsBlocked() *TLContactsBlocked {
	return &TLContactsBlocked{
		Data2: m,
	}
}

//  contacts.blockedSlice#900802a1 count:int blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
//  contacts.blockedSlice#e1664194 count:int blocked:Vector<PeerBlocked> chats:Vector<Chat> users:Vector<User> = contacts.Blocked;
//
func (m *Contacts_Blocked) To_ContactsBlockedSlice() *TLContactsBlockedSlice {
	return &TLContactsBlockedSlice{
		Data2: m,
	}
}

//  contacts.blocked#1c138d15 blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
//  contacts.blocked#ade1591 blocked:Vector<PeerBlocked> chats:Vector<Chat> users:Vector<User> = contacts.Blocked;
//
func (m *TLContactsBlocked) To_Contacts_Blocked() *Contacts_Blocked {
	m.Data2.Cmd = Cmd_ContactsBlocked
	return m.Data2
}

//  contacts.blockedSlice#900802a1 count:int blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
//  contacts.blockedSlice#e1664194 count:int blocked:Vector<PeerBlocked> chats:Vector<Chat> users:Vector<User> = contacts.Blocked;
//
func (m *TLContactsBlockedSlice) To_Contacts_Blocked() *Contacts_Blocked {
	m.Data2.Cmd = Cmd_ContactsBlockedSlice
	return m.Data2
}

//  + TL_ContactsContactsNotModified
//  + TL_ContactsContacts
//

//  contacts.contactsNotModified#b74ba9d2 = contacts.Contacts;
//
func (m *Contacts_Contacts) To_ContactsContactsNotModified() *TLContactsContactsNotModified {
	return &TLContactsContactsNotModified{
		Data2: m,
	}
}

//  contacts.contacts#eae87e42 contacts:Vector<Contact> saved_count:int users:Vector<User> = contacts.Contacts;
//  contacts.contacts#6f8b8cb2 contacts:Vector<Contact> users:Vector<User> = contacts.Contacts;
//
func (m *Contacts_Contacts) To_ContactsContacts() *TLContactsContacts {
	return &TLContactsContacts{
		Data2: m,
	}
}

//  contacts.contactsNotModified#b74ba9d2 = contacts.Contacts;
//
func (m *TLContactsContactsNotModified) To_Contacts_Contacts() *Contacts_Contacts {
	m.Data2.Cmd = Cmd_ContactsContactsNotModified
	return m.Data2
}

//  contacts.contacts#eae87e42 contacts:Vector<Contact> saved_count:int users:Vector<User> = contacts.Contacts;
//  contacts.contacts#6f8b8cb2 contacts:Vector<Contact> users:Vector<User> = contacts.Contacts;
//
func (m *TLContactsContacts) To_Contacts_Contacts() *Contacts_Contacts {
	m.Data2.Cmd = Cmd_ContactsContacts
	return m.Data2
}

//  + TL_ContactsFound
//

//  contacts.found#1aa1f784 results:Vector<Peer> chats:Vector<Chat> users:Vector<User> = contacts.Found;
//  contacts.found#b3134d9d my_results:Vector<Peer> results:Vector<Peer> chats:Vector<Chat> users:Vector<User> = contacts.Found;
//
func (m *Contacts_Found) To_ContactsFound() *TLContactsFound {
	return &TLContactsFound{
		Data2: m,
	}
}

//  contacts.found#1aa1f784 results:Vector<Peer> chats:Vector<Chat> users:Vector<User> = contacts.Found;
//  contacts.found#b3134d9d my_results:Vector<Peer> results:Vector<Peer> chats:Vector<Chat> users:Vector<User> = contacts.Found;
//
func (m *TLContactsFound) To_Contacts_Found() *Contacts_Found {
	m.Data2.Cmd = Cmd_ContactsFound
	return m.Data2
}

//  + TL_ContactsImportedContacts
//

//  contacts.importedContacts#77d01c3b imported:Vector<ImportedContact> popular_invites:Vector<PopularContact> retry_contacts:Vector<long> users:Vector<User> = contacts.ImportedContacts;
//  contacts.importedContacts#ad524315 imported:Vector<ImportedContact> retry_contacts:Vector<long> users:Vector<User> = contacts.ImportedContacts;
//
func (m *Contacts_ImportedContacts) To_ContactsImportedContacts() *TLContactsImportedContacts {
	return &TLContactsImportedContacts{
		Data2: m,
	}
}

//  contacts.importedContacts#77d01c3b imported:Vector<ImportedContact> popular_invites:Vector<PopularContact> retry_contacts:Vector<long> users:Vector<User> = contacts.ImportedContacts;
//  contacts.importedContacts#ad524315 imported:Vector<ImportedContact> retry_contacts:Vector<long> users:Vector<User> = contacts.ImportedContacts;
//
func (m *TLContactsImportedContacts) To_Contacts_ImportedContacts() *Contacts_ImportedContacts {
	m.Data2.Cmd = Cmd_ContactsImportedContacts
	return m.Data2
}

//  + TL_ContactsLink
//

//  contacts.link#3ace484c my_link:ContactLink foreign_link:ContactLink user:User = contacts.Link;
//
func (m *Contacts_Link) To_ContactsLink() *TLContactsLink {
	return &TLContactsLink{
		Data2: m,
	}
}

//  contacts.link#3ace484c my_link:ContactLink foreign_link:ContactLink user:User = contacts.Link;
//
func (m *TLContactsLink) To_Contacts_Link() *Contacts_Link {
	m.Data2.Cmd = Cmd_ContactsLink
	return m.Data2
}

//  + TL_ContactsResolvedPeer
//

//  contacts.resolvedPeer#7f077ad9 peer:Peer chats:Vector<Chat> users:Vector<User> = contacts.ResolvedPeer;
//
func (m *Contacts_ResolvedPeer) To_ContactsResolvedPeer() *TLContactsResolvedPeer {
	return &TLContactsResolvedPeer{
		Data2: m,
	}
}

//  contacts.resolvedPeer#7f077ad9 peer:Peer chats:Vector<Chat> users:Vector<User> = contacts.ResolvedPeer;
//
func (m *TLContactsResolvedPeer) To_Contacts_ResolvedPeer() *Contacts_ResolvedPeer {
	m.Data2.Cmd = Cmd_ContactsResolvedPeer
	return m.Data2
}

//  + TL_ContactsTopPeersNotModified
//  + TL_ContactsTopPeers
//  + TL_ContactsTopPeersDisabled
//

//  contacts.topPeersNotModified#de266ef5 = contacts.TopPeers;
//
func (m *Contacts_TopPeers) To_ContactsTopPeersNotModified() *TLContactsTopPeersNotModified {
	return &TLContactsTopPeersNotModified{
		Data2: m,
	}
}

//  contacts.topPeers#70b772a8 categories:Vector<TopPeerCategoryPeers> chats:Vector<Chat> users:Vector<User> = contacts.TopPeers;
//
func (m *Contacts_TopPeers) To_ContactsTopPeers() *TLContactsTopPeers {
	return &TLContactsTopPeers{
		Data2: m,
	}
}

//  contacts.topPeersDisabled#b52c939d = contacts.TopPeers;
//
func (m *Contacts_TopPeers) To_ContactsTopPeersDisabled() *TLContactsTopPeersDisabled {
	return &TLContactsTopPeersDisabled{
		Data2: m,
	}
}

//  contacts.topPeersNotModified#de266ef5 = contacts.TopPeers;
//
func (m *TLContactsTopPeersNotModified) To_Contacts_TopPeers() *Contacts_TopPeers {
	m.Data2.Cmd = Cmd_ContactsTopPeersNotModified
	return m.Data2
}

//  contacts.topPeers#70b772a8 categories:Vector<TopPeerCategoryPeers> chats:Vector<Chat> users:Vector<User> = contacts.TopPeers;
//
func (m *TLContactsTopPeers) To_Contacts_TopPeers() *Contacts_TopPeers {
	m.Data2.Cmd = Cmd_ContactsTopPeers
	return m.Data2
}

//  contacts.topPeersDisabled#b52c939d = contacts.TopPeers;
//
func (m *TLContactsTopPeersDisabled) To_Contacts_TopPeers() *Contacts_TopPeers {
	m.Data2.Cmd = Cmd_ContactsTopPeersDisabled
	return m.Data2
}

//  + TL_DataJSON
//

//  dataJSON#7d748d04 data:string = DataJSON;
//
func (m *DataJSON) To_DataJSON() *TLDataJSON {
	return &TLDataJSON{
		Data2: m,
	}
}

//  dataJSON#7d748d04 data:string = DataJSON;
//
func (m *TLDataJSON) To_DataJSON() *DataJSON {
	m.Data2.Cmd = Cmd_DataJSON
	return m.Data2
}

//  + TL_DcOption
//

//  dcOption#5d8c6cc flags:# ipv6:flags.0?true media_only:flags.1?true tcpo_only:flags.2?true cdn:flags.3?true static:flags.4?true id:int ip_address:string port:int = DcOption;
//  dcOption#18b7a10d flags:# ipv6:flags.0?true media_only:flags.1?true tcpo_only:flags.2?true cdn:flags.3?true static:flags.4?true id:int ip_address:string port:int secret:flags.10?bytes = DcOption;
//
func (m *DcOption) To_DcOption() *TLDcOption {
	return &TLDcOption{
		Data2: m,
	}
}

//  dcOption#5d8c6cc flags:# ipv6:flags.0?true media_only:flags.1?true tcpo_only:flags.2?true cdn:flags.3?true static:flags.4?true id:int ip_address:string port:int = DcOption;
//  dcOption#18b7a10d flags:# ipv6:flags.0?true media_only:flags.1?true tcpo_only:flags.2?true cdn:flags.3?true static:flags.4?true id:int ip_address:string port:int secret:flags.10?bytes = DcOption;
//
func (m *TLDcOption) To_DcOption() *DcOption {
	m.Data2.Cmd = Cmd_DcOption
	return m.Data2
}

//  + TL_DestroyAuthKeyOk
//  + TL_DestroyAuthKeyNone
//  + TL_DestroyAuthKeyFail
//

//  destroy_auth_key_ok#f660e1d4 = DestroyAuthKeyRes;
//
func (m *DestroyAuthKeyRes) To_DestroyAuthKeyOk() *TLDestroyAuthKeyOk {
	return &TLDestroyAuthKeyOk{
		Data2: m,
	}
}

//  destroy_auth_key_none#0a9f2259 = DestroyAuthKeyRes;
//
func (m *DestroyAuthKeyRes) To_DestroyAuthKeyNone() *TLDestroyAuthKeyNone {
	return &TLDestroyAuthKeyNone{
		Data2: m,
	}
}

//  destroy_auth_key_fail#ea109b13 = DestroyAuthKeyRes;
//
func (m *DestroyAuthKeyRes) To_DestroyAuthKeyFail() *TLDestroyAuthKeyFail {
	return &TLDestroyAuthKeyFail{
		Data2: m,
	}
}

//  destroy_auth_key_ok#f660e1d4 = DestroyAuthKeyRes;
//
func (m *TLDestroyAuthKeyOk) To_DestroyAuthKeyRes() *DestroyAuthKeyRes {
	m.Data2.Cmd = Cmd_DestroyAuthKeyOk
	return m.Data2
}

//  destroy_auth_key_none#0a9f2259 = DestroyAuthKeyRes;
//
func (m *TLDestroyAuthKeyNone) To_DestroyAuthKeyRes() *DestroyAuthKeyRes {
	m.Data2.Cmd = Cmd_DestroyAuthKeyNone
	return m.Data2
}

//  destroy_auth_key_fail#ea109b13 = DestroyAuthKeyRes;
//
func (m *TLDestroyAuthKeyFail) To_DestroyAuthKeyRes() *DestroyAuthKeyRes {
	m.Data2.Cmd = Cmd_DestroyAuthKeyFail
	return m.Data2
}

//  + TL_DestroySessionOk
//  + TL_DestroySessionNone
//

//  destroy_session_ok#e22045fc session_id:long = DestroySessionRes;
//
func (m *DestroySessionRes) To_DestroySessionOk() *TLDestroySessionOk {
	return &TLDestroySessionOk{
		Data2: m,
	}
}

//  destroy_session_none#62d350c9 session_id:long = DestroySessionRes;
//
func (m *DestroySessionRes) To_DestroySessionNone() *TLDestroySessionNone {
	return &TLDestroySessionNone{
		Data2: m,
	}
}

//  destroy_session_ok#e22045fc session_id:long = DestroySessionRes;
//
func (m *TLDestroySessionOk) To_DestroySessionRes() *DestroySessionRes {
	m.Data2.Cmd = Cmd_DestroySessionOk
	return m.Data2
}

//  destroy_session_none#62d350c9 session_id:long = DestroySessionRes;
//
func (m *TLDestroySessionNone) To_DestroySessionRes() *DestroySessionRes {
	m.Data2.Cmd = Cmd_DestroySessionNone
	return m.Data2
}

//  + TL_Dialog
//  + TL_DialogChannel
//  + TL_DialogFolder
//

//  dialog#e4def5db flags:# pinned:flags.2?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage = Dialog;
//  dialog#c1dd804a peer:Peer top_message:int read_inbox_max_id:int unread_count:int notify_settings:PeerNotifySettings = Dialog;
//  dialog#2c171f72 flags:# pinned:flags.2?true unread_mark:flags.3?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage folder_id:flags.4?int = Dialog;
//
func (m *Dialog) To_Dialog() *TLDialog {
	return &TLDialog{
		Data2: m,
	}
}

//  dialogChannel#5b8496b2 peer:Peer top_message:int top_important_message:int read_inbox_max_id:int unread_count:int unread_important_count:int notify_settings:PeerNotifySettings pts:int = Dialog;
//
func (m *Dialog) To_DialogChannel() *TLDialogChannel {
	return &TLDialogChannel{
		Data2: m,
	}
}

//  dialogFolder#71bd134c flags:# pinned:flags.2?true folder:Folder peer:Peer top_message:int unread_muted_peers_count:int unread_unmuted_peers_count:int unread_muted_messages_count:int unread_unmuted_messages_count:int = Dialog;
//
func (m *Dialog) To_DialogFolder() *TLDialogFolder {
	return &TLDialogFolder{
		Data2: m,
	}
}

//  dialog#e4def5db flags:# pinned:flags.2?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage = Dialog;
//  dialog#c1dd804a peer:Peer top_message:int read_inbox_max_id:int unread_count:int notify_settings:PeerNotifySettings = Dialog;
//  dialog#2c171f72 flags:# pinned:flags.2?true unread_mark:flags.3?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage folder_id:flags.4?int = Dialog;
//
func (m *TLDialog) To_Dialog() *Dialog {
	m.Data2.Cmd = Cmd_Dialog
	return m.Data2
}

//  dialogChannel#5b8496b2 peer:Peer top_message:int top_important_message:int read_inbox_max_id:int unread_count:int unread_important_count:int notify_settings:PeerNotifySettings pts:int = Dialog;
//
func (m *TLDialogChannel) To_Dialog() *Dialog {
	m.Data2.Cmd = Cmd_DialogChannel
	return m.Data2
}

//  dialogFolder#71bd134c flags:# pinned:flags.2?true folder:Folder peer:Peer top_message:int unread_muted_peers_count:int unread_unmuted_peers_count:int unread_muted_messages_count:int unread_unmuted_messages_count:int = Dialog;
//
func (m *TLDialogFolder) To_Dialog() *Dialog {
	m.Data2.Cmd = Cmd_DialogFolder
	return m.Data2
}

//  + TL_DialogFilter
//

//  dialogFilter#7438f7e8 flags:# contacts:flags.0?true non_contacts:flags.1?true groups:flags.2?true broadcasts:flags.3?true bots:flags.4?true exclude_muted:flags.11?true exclude_read:flags.12?true exclude_archived:flags.13?true id:int title:string emoticon:flags.25?string pinned_peers:Vector<InputPeer> include_peers:Vector<InputPeer> exclude_peers:Vector<InputPeer> = DialogFilter;
//
func (m *DialogFilter) To_DialogFilter() *TLDialogFilter {
	return &TLDialogFilter{
		Data2: m,
	}
}

//  dialogFilter#7438f7e8 flags:# contacts:flags.0?true non_contacts:flags.1?true groups:flags.2?true broadcasts:flags.3?true bots:flags.4?true exclude_muted:flags.11?true exclude_read:flags.12?true exclude_archived:flags.13?true id:int title:string emoticon:flags.25?string pinned_peers:Vector<InputPeer> include_peers:Vector<InputPeer> exclude_peers:Vector<InputPeer> = DialogFilter;
//
func (m *TLDialogFilter) To_DialogFilter() *DialogFilter {
	m.Data2.Cmd = Cmd_DialogFilter
	return m.Data2
}

//  + TL_DialogFilterSuggested
//

//  dialogFilterSuggested#77744d4a filter:DialogFilter description:string = DialogFilterSuggested;
//
func (m *DialogFilterSuggested) To_DialogFilterSuggested() *TLDialogFilterSuggested {
	return &TLDialogFilterSuggested{
		Data2: m,
	}
}

//  dialogFilterSuggested#77744d4a filter:DialogFilter description:string = DialogFilterSuggested;
//
func (m *TLDialogFilterSuggested) To_DialogFilterSuggested() *DialogFilterSuggested {
	m.Data2.Cmd = Cmd_DialogFilterSuggested
	return m.Data2
}

//  + TL_DialogPeer
//  + TL_DialogPeerFolder
//

//  dialogPeer#e56dbf05 peer:Peer = DialogPeer;
//
func (m *DialogPeer) To_DialogPeer() *TLDialogPeer {
	return &TLDialogPeer{
		Data2: m,
	}
}

//  dialogPeerFolder#514519e2 folder_id:int = DialogPeer;
//
func (m *DialogPeer) To_DialogPeerFolder() *TLDialogPeerFolder {
	return &TLDialogPeerFolder{
		Data2: m,
	}
}

//  dialogPeer#e56dbf05 peer:Peer = DialogPeer;
//
func (m *TLDialogPeer) To_DialogPeer() *DialogPeer {
	m.Data2.Cmd = Cmd_DialogPeer
	return m.Data2
}

//  dialogPeerFolder#514519e2 folder_id:int = DialogPeer;
//
func (m *TLDialogPeerFolder) To_DialogPeer() *DialogPeer {
	m.Data2.Cmd = Cmd_DialogPeerFolder
	return m.Data2
}

//  + TL_DisabledFeature
//

//  disabledFeature#ae636f24 feature:string description:string = DisabledFeature;
//
func (m *DisabledFeature) To_DisabledFeature() *TLDisabledFeature {
	return &TLDisabledFeature{
		Data2: m,
	}
}

//  disabledFeature#ae636f24 feature:string description:string = DisabledFeature;
//
func (m *TLDisabledFeature) To_DisabledFeature() *DisabledFeature {
	m.Data2.Cmd = Cmd_DisabledFeature
	return m.Data2
}

//  + TL_DocumentEmpty
//  + TL_Document
//

//  documentEmpty#36f8c871 id:long = Document;
//
func (m *Document) To_DocumentEmpty() *TLDocumentEmpty {
	return &TLDocumentEmpty{
		Data2: m,
	}
}

//  document#87232bc7 id:long access_hash:long date:int mime_type:string size:int thumb:PhotoSize dc_id:int version:int attributes:Vector<DocumentAttribute> = Document;
//  document#f9a39f4f id:long access_hash:long date:int mime_type:string size:int thumb:PhotoSize dc_id:int attributes:Vector<DocumentAttribute> = Document;
//  document#59534e4c id:long access_hash:long file_reference:bytes date:int mime_type:string size:int thumb:PhotoSize dc_id:int attributes:Vector<DocumentAttribute> = Document;
//  document#9ba29cc1 flags:# id:long access_hash:long file_reference:bytes date:int mime_type:string size:int thumbs:flags.0?Vector<PhotoSize> dc_id:int attributes:Vector<DocumentAttribute> = Document;
//  document#1e87342b flags:# id:long access_hash:long file_reference:bytes date:int mime_type:string size:int thumbs:flags.0?Vector<PhotoSize> video_thumbs:flags.1?Vector<VideoSize> dc_id:int attributes:Vector<DocumentAttribute> = Document;
//
func (m *Document) To_Document() *TLDocument {
	return &TLDocument{
		Data2: m,
	}
}

//  documentEmpty#36f8c871 id:long = Document;
//
func (m *TLDocumentEmpty) To_Document() *Document {
	m.Data2.Cmd = Cmd_DocumentEmpty
	return m.Data2
}

//  document#87232bc7 id:long access_hash:long date:int mime_type:string size:int thumb:PhotoSize dc_id:int version:int attributes:Vector<DocumentAttribute> = Document;
//  document#f9a39f4f id:long access_hash:long date:int mime_type:string size:int thumb:PhotoSize dc_id:int attributes:Vector<DocumentAttribute> = Document;
//  document#59534e4c id:long access_hash:long file_reference:bytes date:int mime_type:string size:int thumb:PhotoSize dc_id:int attributes:Vector<DocumentAttribute> = Document;
//  document#9ba29cc1 flags:# id:long access_hash:long file_reference:bytes date:int mime_type:string size:int thumbs:flags.0?Vector<PhotoSize> dc_id:int attributes:Vector<DocumentAttribute> = Document;
//  document#1e87342b flags:# id:long access_hash:long file_reference:bytes date:int mime_type:string size:int thumbs:flags.0?Vector<PhotoSize> video_thumbs:flags.1?Vector<VideoSize> dc_id:int attributes:Vector<DocumentAttribute> = Document;
//
func (m *TLDocument) To_Document() *Document {
	m.Data2.Cmd = Cmd_Document
	return m.Data2
}

//  + TL_DocumentAttributeImageSize
//  + TL_DocumentAttributeAnimated
//  + TL_DocumentAttributeSticker
//  + TL_DocumentAttributeVideo
//  + TL_DocumentAttributeAudio
//  + TL_DocumentAttributeFilename
//  + TL_DocumentAttributeHasStickers
//

//  documentAttributeImageSize#6c37c15c w:int h:int = DocumentAttribute;
//
func (m *DocumentAttribute) To_DocumentAttributeImageSize() *TLDocumentAttributeImageSize {
	return &TLDocumentAttributeImageSize{
		Data2: m,
	}
}

//  documentAttributeAnimated#11b58939 = DocumentAttribute;
//
func (m *DocumentAttribute) To_DocumentAttributeAnimated() *TLDocumentAttributeAnimated {
	return &TLDocumentAttributeAnimated{
		Data2: m,
	}
}

//  documentAttributeSticker#6319d612 flags:# mask:flags.1?true alt:string stickerset:InputStickerSet mask_coords:flags.0?MaskCoords = DocumentAttribute;
//  documentAttributeSticker#3a556302 alt:string stickerset:InputStickerSet = DocumentAttribute;
//
func (m *DocumentAttribute) To_DocumentAttributeSticker() *TLDocumentAttributeSticker {
	return &TLDocumentAttributeSticker{
		Data2: m,
	}
}

//  documentAttributeVideo#ef02ce6 flags:# round_message:flags.0?true duration:int w:int h:int = DocumentAttribute;
//  documentAttributeVideo#5910cccb duration:int w:int h:int = DocumentAttribute;
//
func (m *DocumentAttribute) To_DocumentAttributeVideo() *TLDocumentAttributeVideo {
	return &TLDocumentAttributeVideo{
		Data2: m,
	}
}

//  documentAttributeAudio#9852f9c6 flags:# voice:flags.10?true duration:int title:flags.0?string performer:flags.1?string waveform:flags.2?bytes = DocumentAttribute;
//
func (m *DocumentAttribute) To_DocumentAttributeAudio() *TLDocumentAttributeAudio {
	return &TLDocumentAttributeAudio{
		Data2: m,
	}
}

//  documentAttributeFilename#15590068 file_name:string = DocumentAttribute;
//
func (m *DocumentAttribute) To_DocumentAttributeFilename() *TLDocumentAttributeFilename {
	return &TLDocumentAttributeFilename{
		Data2: m,
	}
}

//  documentAttributeHasStickers#9801d2f7 = DocumentAttribute;
//
func (m *DocumentAttribute) To_DocumentAttributeHasStickers() *TLDocumentAttributeHasStickers {
	return &TLDocumentAttributeHasStickers{
		Data2: m,
	}
}

//  documentAttributeImageSize#6c37c15c w:int h:int = DocumentAttribute;
//
func (m *TLDocumentAttributeImageSize) To_DocumentAttribute() *DocumentAttribute {
	m.Data2.Cmd = Cmd_DocumentAttributeImageSize
	return m.Data2
}

//  documentAttributeAnimated#11b58939 = DocumentAttribute;
//
func (m *TLDocumentAttributeAnimated) To_DocumentAttribute() *DocumentAttribute {
	m.Data2.Cmd = Cmd_DocumentAttributeAnimated
	return m.Data2
}

//  documentAttributeSticker#6319d612 flags:# mask:flags.1?true alt:string stickerset:InputStickerSet mask_coords:flags.0?MaskCoords = DocumentAttribute;
//  documentAttributeSticker#3a556302 alt:string stickerset:InputStickerSet = DocumentAttribute;
//
func (m *TLDocumentAttributeSticker) To_DocumentAttribute() *DocumentAttribute {
	m.Data2.Cmd = Cmd_DocumentAttributeSticker
	return m.Data2
}

//  documentAttributeVideo#ef02ce6 flags:# round_message:flags.0?true duration:int w:int h:int = DocumentAttribute;
//  documentAttributeVideo#5910cccb duration:int w:int h:int = DocumentAttribute;
//
func (m *TLDocumentAttributeVideo) To_DocumentAttribute() *DocumentAttribute {
	m.Data2.Cmd = Cmd_DocumentAttributeVideo
	return m.Data2
}

//  documentAttributeAudio#9852f9c6 flags:# voice:flags.10?true duration:int title:flags.0?string performer:flags.1?string waveform:flags.2?bytes = DocumentAttribute;
//
func (m *TLDocumentAttributeAudio) To_DocumentAttribute() *DocumentAttribute {
	m.Data2.Cmd = Cmd_DocumentAttributeAudio
	return m.Data2
}

//  documentAttributeFilename#15590068 file_name:string = DocumentAttribute;
//
func (m *TLDocumentAttributeFilename) To_DocumentAttribute() *DocumentAttribute {
	m.Data2.Cmd = Cmd_DocumentAttributeFilename
	return m.Data2
}

//  documentAttributeHasStickers#9801d2f7 = DocumentAttribute;
//
func (m *TLDocumentAttributeHasStickers) To_DocumentAttribute() *DocumentAttribute {
	m.Data2.Cmd = Cmd_DocumentAttributeHasStickers
	return m.Data2
}

//  + TL_DraftMessageEmpty
//  + TL_DraftMessage
//

//  draftMessageEmpty#ba4baec5 = DraftMessage;
//  draftMessageEmpty#1b0c841a flags:# date:flags.0?int = DraftMessage;
//
func (m *DraftMessage) To_DraftMessageEmpty() *TLDraftMessageEmpty {
	return &TLDraftMessageEmpty{
		Data2: m,
	}
}

//  draftMessage#fd8e711f flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int message:string entities:flags.3?Vector<MessageEntity> date:int = DraftMessage;
//
func (m *DraftMessage) To_DraftMessage() *TLDraftMessage {
	return &TLDraftMessage{
		Data2: m,
	}
}

//  draftMessageEmpty#ba4baec5 = DraftMessage;
//  draftMessageEmpty#1b0c841a flags:# date:flags.0?int = DraftMessage;
//
func (m *TLDraftMessageEmpty) To_DraftMessage() *DraftMessage {
	m.Data2.Cmd = Cmd_DraftMessageEmpty
	return m.Data2
}

//  draftMessage#fd8e711f flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int message:string entities:flags.3?Vector<MessageEntity> date:int = DraftMessage;
//
func (m *TLDraftMessage) To_DraftMessage() *DraftMessage {
	m.Data2.Cmd = Cmd_DraftMessage
	return m.Data2
}

//  + TL_EmojiKeyword
//  + TL_EmojiKeywordDeleted
//

//  emojiKeyword#d5b3b9f9 keyword:string emoticons:Vector<string> = EmojiKeyword;
//
func (m *EmojiKeyword) To_EmojiKeyword() *TLEmojiKeyword {
	return &TLEmojiKeyword{
		Data2: m,
	}
}

//  emojiKeywordDeleted#236df622 keyword:string emoticons:Vector<string> = EmojiKeyword;
//
func (m *EmojiKeyword) To_EmojiKeywordDeleted() *TLEmojiKeywordDeleted {
	return &TLEmojiKeywordDeleted{
		Data2: m,
	}
}

//  emojiKeyword#d5b3b9f9 keyword:string emoticons:Vector<string> = EmojiKeyword;
//
func (m *TLEmojiKeyword) To_EmojiKeyword() *EmojiKeyword {
	m.Data2.Cmd = Cmd_EmojiKeyword
	return m.Data2
}

//  emojiKeywordDeleted#236df622 keyword:string emoticons:Vector<string> = EmojiKeyword;
//
func (m *TLEmojiKeywordDeleted) To_EmojiKeyword() *EmojiKeyword {
	m.Data2.Cmd = Cmd_EmojiKeywordDeleted
	return m.Data2
}

//  + TL_EmojiKeywordsDifference
//

//  emojiKeywordsDifference#5cc761bd lang_code:string from_version:int version:int keywords:Vector<EmojiKeyword> = EmojiKeywordsDifference;
//
func (m *EmojiKeywordsDifference) To_EmojiKeywordsDifference() *TLEmojiKeywordsDifference {
	return &TLEmojiKeywordsDifference{
		Data2: m,
	}
}

//  emojiKeywordsDifference#5cc761bd lang_code:string from_version:int version:int keywords:Vector<EmojiKeyword> = EmojiKeywordsDifference;
//
func (m *TLEmojiKeywordsDifference) To_EmojiKeywordsDifference() *EmojiKeywordsDifference {
	m.Data2.Cmd = Cmd_EmojiKeywordsDifference
	return m.Data2
}

//  + TL_EmojiLanguage
//

//  emojiLanguage#b3fb5361 lang_code:string = EmojiLanguage;
//
func (m *EmojiLanguage) To_EmojiLanguage() *TLEmojiLanguage {
	return &TLEmojiLanguage{
		Data2: m,
	}
}

//  emojiLanguage#b3fb5361 lang_code:string = EmojiLanguage;
//
func (m *TLEmojiLanguage) To_EmojiLanguage() *EmojiLanguage {
	m.Data2.Cmd = Cmd_EmojiLanguage
	return m.Data2
}

//  + TL_EmojiURL
//

//  emojiURL#a575739d url:string = EmojiURL;
//
func (m *EmojiURL) To_EmojiURL() *TLEmojiURL {
	return &TLEmojiURL{
		Data2: m,
	}
}

//  emojiURL#a575739d url:string = EmojiURL;
//
func (m *TLEmojiURL) To_EmojiURL() *EmojiURL {
	m.Data2.Cmd = Cmd_EmojiURL
	return m.Data2
}

//  + TL_EncryptedChatEmpty
//  + TL_EncryptedChatWaiting
//  + TL_EncryptedChatRequested
//  + TL_EncryptedChat
//  + TL_EncryptedChatDiscarded
//

//  encryptedChatEmpty#ab7ec0a0 id:int = EncryptedChat;
//
func (m *EncryptedChat) To_EncryptedChatEmpty() *TLEncryptedChatEmpty {
	return &TLEncryptedChatEmpty{
		Data2: m,
	}
}

//  encryptedChatWaiting#3bf703dc id:int access_hash:long date:int admin_id:int participant_id:int = EncryptedChat;
//
func (m *EncryptedChat) To_EncryptedChatWaiting() *TLEncryptedChatWaiting {
	return &TLEncryptedChatWaiting{
		Data2: m,
	}
}

//  encryptedChatRequested#c878527e id:int access_hash:long date:int admin_id:int participant_id:int g_a:bytes = EncryptedChat;
//  encryptedChatRequested#62718a82 flags:# folder_id:flags.0?int id:int access_hash:long date:int admin_id:int participant_id:int g_a:bytes = EncryptedChat;
//
func (m *EncryptedChat) To_EncryptedChatRequested() *TLEncryptedChatRequested {
	return &TLEncryptedChatRequested{
		Data2: m,
	}
}

//  encryptedChat#fa56ce36 id:int access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long = EncryptedChat;
//
func (m *EncryptedChat) To_EncryptedChat() *TLEncryptedChat {
	return &TLEncryptedChat{
		Data2: m,
	}
}

//  encryptedChatDiscarded#13d6dd27 id:int = EncryptedChat;
//
func (m *EncryptedChat) To_EncryptedChatDiscarded() *TLEncryptedChatDiscarded {
	return &TLEncryptedChatDiscarded{
		Data2: m,
	}
}

//  encryptedChatEmpty#ab7ec0a0 id:int = EncryptedChat;
//
func (m *TLEncryptedChatEmpty) To_EncryptedChat() *EncryptedChat {
	m.Data2.Cmd = Cmd_EncryptedChatEmpty
	return m.Data2
}

//  encryptedChatWaiting#3bf703dc id:int access_hash:long date:int admin_id:int participant_id:int = EncryptedChat;
//
func (m *TLEncryptedChatWaiting) To_EncryptedChat() *EncryptedChat {
	m.Data2.Cmd = Cmd_EncryptedChatWaiting
	return m.Data2
}

//  encryptedChatRequested#c878527e id:int access_hash:long date:int admin_id:int participant_id:int g_a:bytes = EncryptedChat;
//  encryptedChatRequested#62718a82 flags:# folder_id:flags.0?int id:int access_hash:long date:int admin_id:int participant_id:int g_a:bytes = EncryptedChat;
//
func (m *TLEncryptedChatRequested) To_EncryptedChat() *EncryptedChat {
	m.Data2.Cmd = Cmd_EncryptedChatRequested
	return m.Data2
}

//  encryptedChat#fa56ce36 id:int access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long = EncryptedChat;
//
func (m *TLEncryptedChat) To_EncryptedChat() *EncryptedChat {
	m.Data2.Cmd = Cmd_EncryptedChat
	return m.Data2
}

//  encryptedChatDiscarded#13d6dd27 id:int = EncryptedChat;
//
func (m *TLEncryptedChatDiscarded) To_EncryptedChat() *EncryptedChat {
	m.Data2.Cmd = Cmd_EncryptedChatDiscarded
	return m.Data2
}

//  + TL_EncryptedFileEmpty
//  + TL_EncryptedFile
//

//  encryptedFileEmpty#c21f497e = EncryptedFile;
//
func (m *EncryptedFile) To_EncryptedFileEmpty() *TLEncryptedFileEmpty {
	return &TLEncryptedFileEmpty{
		Data2: m,
	}
}

//  encryptedFile#4a70994c id:long access_hash:long size:int dc_id:int key_fingerprint:int = EncryptedFile;
//
func (m *EncryptedFile) To_EncryptedFile() *TLEncryptedFile {
	return &TLEncryptedFile{
		Data2: m,
	}
}

//  encryptedFileEmpty#c21f497e = EncryptedFile;
//
func (m *TLEncryptedFileEmpty) To_EncryptedFile() *EncryptedFile {
	m.Data2.Cmd = Cmd_EncryptedFileEmpty
	return m.Data2
}

//  encryptedFile#4a70994c id:long access_hash:long size:int dc_id:int key_fingerprint:int = EncryptedFile;
//
func (m *TLEncryptedFile) To_EncryptedFile() *EncryptedFile {
	m.Data2.Cmd = Cmd_EncryptedFile
	return m.Data2
}

//  + TL_EncryptedMessage
//  + TL_EncryptedMessageService
//

//  encryptedMessage#ed18c118 random_id:long chat_id:int date:int bytes:bytes file:EncryptedFile = EncryptedMessage;
//
func (m *EncryptedMessage) To_EncryptedMessage() *TLEncryptedMessage {
	return &TLEncryptedMessage{
		Data2: m,
	}
}

//  encryptedMessageService#23734b06 random_id:long chat_id:int date:int bytes:bytes = EncryptedMessage;
//
func (m *EncryptedMessage) To_EncryptedMessageService() *TLEncryptedMessageService {
	return &TLEncryptedMessageService{
		Data2: m,
	}
}

//  encryptedMessage#ed18c118 random_id:long chat_id:int date:int bytes:bytes file:EncryptedFile = EncryptedMessage;
//
func (m *TLEncryptedMessage) To_EncryptedMessage() *EncryptedMessage {
	m.Data2.Cmd = Cmd_EncryptedMessage
	return m.Data2
}

//  encryptedMessageService#23734b06 random_id:long chat_id:int date:int bytes:bytes = EncryptedMessage;
//
func (m *TLEncryptedMessageService) To_EncryptedMessage() *EncryptedMessage {
	m.Data2.Cmd = Cmd_EncryptedMessageService
	return m.Data2
}

//  + TL_Error
//

//  error#c4b9f9bb code:int text:string = Error;
//
func (m *Error) To_Error() *TLError {
	return &TLError{
		Data2: m,
	}
}

//  error#c4b9f9bb code:int text:string = Error;
//
func (m *TLError) To_Error() *Error {
	m.Data2.Cmd = Cmd_Error
	return m.Data2
}

//  + TL_ChatInviteEmpty
//  + TL_ChatInviteExported
//

//  chatInviteEmpty#69df3769 = ExportedChatInvite;
//
func (m *ExportedChatInvite) To_ChatInviteEmpty() *TLChatInviteEmpty {
	return &TLChatInviteEmpty{
		Data2: m,
	}
}

//  chatInviteExported#fc2e05bc link:string = ExportedChatInvite;
//
func (m *ExportedChatInvite) To_ChatInviteExported() *TLChatInviteExported {
	return &TLChatInviteExported{
		Data2: m,
	}
}

//  chatInviteEmpty#69df3769 = ExportedChatInvite;
//
func (m *TLChatInviteEmpty) To_ExportedChatInvite() *ExportedChatInvite {
	m.Data2.Cmd = Cmd_ChatInviteEmpty
	return m.Data2
}

//  chatInviteExported#fc2e05bc link:string = ExportedChatInvite;
//
func (m *TLChatInviteExported) To_ExportedChatInvite() *ExportedChatInvite {
	m.Data2.Cmd = Cmd_ChatInviteExported
	return m.Data2
}

//  + TL_ExportedMessageLink
//

//  exportedMessageLink#1f486803 link:string = ExportedMessageLink;
//  exportedMessageLink#5dab1af4 link:string html:string = ExportedMessageLink;
//
func (m *ExportedMessageLink) To_ExportedMessageLink() *TLExportedMessageLink {
	return &TLExportedMessageLink{
		Data2: m,
	}
}

//  exportedMessageLink#1f486803 link:string = ExportedMessageLink;
//  exportedMessageLink#5dab1af4 link:string html:string = ExportedMessageLink;
//
func (m *TLExportedMessageLink) To_ExportedMessageLink() *ExportedMessageLink {
	m.Data2.Cmd = Cmd_ExportedMessageLink
	return m.Data2
}

//  + TL_FileHash
//

//  fileHash#6242c773 offset:int limit:int hash:bytes = FileHash;
//
func (m *FileHash) To_FileHash() *TLFileHash {
	return &TLFileHash{
		Data2: m,
	}
}

//  fileHash#6242c773 offset:int limit:int hash:bytes = FileHash;
//
func (m *TLFileHash) To_FileHash() *FileHash {
	m.Data2.Cmd = Cmd_FileHash
	return m.Data2
}

//  + TL_FileLocationUnavailable
//  + TL_FileLocation
//  + TL_FileLocationToBeDeprecated
//

//  fileLocationUnavailable#7c596b46 volume_id:long local_id:int secret:long = FileLocation;
//
func (m *FileLocation) To_FileLocationUnavailable() *TLFileLocationUnavailable {
	return &TLFileLocationUnavailable{
		Data2: m,
	}
}

//  fileLocation#53d69076 dc_id:int volume_id:long local_id:int secret:long = FileLocation;
//  fileLocation#91d11eb dc_id:int volume_id:long local_id:int secret:long file_reference:bytes = FileLocation;
//
func (m *FileLocation) To_FileLocation() *TLFileLocation {
	return &TLFileLocation{
		Data2: m,
	}
}

//  fileLocationToBeDeprecated#bc7fc6cd volume_id:long local_id:int = FileLocation;
//
func (m *FileLocation) To_FileLocationToBeDeprecated() *TLFileLocationToBeDeprecated {
	return &TLFileLocationToBeDeprecated{
		Data2: m,
	}
}

//  fileLocationUnavailable#7c596b46 volume_id:long local_id:int secret:long = FileLocation;
//
func (m *TLFileLocationUnavailable) To_FileLocation() *FileLocation {
	m.Data2.Cmd = Cmd_FileLocationUnavailable
	return m.Data2
}

//  fileLocation#53d69076 dc_id:int volume_id:long local_id:int secret:long = FileLocation;
//  fileLocation#91d11eb dc_id:int volume_id:long local_id:int secret:long file_reference:bytes = FileLocation;
//
func (m *TLFileLocation) To_FileLocation() *FileLocation {
	m.Data2.Cmd = Cmd_FileLocation
	return m.Data2
}

//  fileLocationToBeDeprecated#bc7fc6cd volume_id:long local_id:int = FileLocation;
//
func (m *TLFileLocationToBeDeprecated) To_FileLocation() *FileLocation {
	m.Data2.Cmd = Cmd_FileLocationToBeDeprecated
	return m.Data2
}

//  + TL_Folder
//

//  folder#ff544e65 flags:# autofill_new_broadcasts:flags.0?true autofill_public_groups:flags.1?true autofill_new_correspondents:flags.2?true id:int title:string photo:flags.3?ChatPhoto = Folder;
//
func (m *Folder) To_Folder() *TLFolder {
	return &TLFolder{
		Data2: m,
	}
}

//  folder#ff544e65 flags:# autofill_new_broadcasts:flags.0?true autofill_public_groups:flags.1?true autofill_new_correspondents:flags.2?true id:int title:string photo:flags.3?ChatPhoto = Folder;
//
func (m *TLFolder) To_Folder() *Folder {
	m.Data2.Cmd = Cmd_Folder
	return m.Data2
}

//  + TL_FolderPeer
//

//  folderPeer#e9baa668 peer:Peer folder_id:int = FolderPeer;
//
func (m *FolderPeer) To_FolderPeer() *TLFolderPeer {
	return &TLFolderPeer{
		Data2: m,
	}
}

//  folderPeer#e9baa668 peer:Peer folder_id:int = FolderPeer;
//
func (m *TLFolderPeer) To_FolderPeer() *FolderPeer {
	m.Data2.Cmd = Cmd_FolderPeer
	return m.Data2
}

//  + TL_FoundGif
//  + TL_FoundGifCached
//

//  foundGif#162ecc1f url:string thumb_url:string content_url:string content_type:string w:int h:int = FoundGif;
//
func (m *FoundGif) To_FoundGif() *TLFoundGif {
	return &TLFoundGif{
		Data2: m,
	}
}

//  foundGifCached#9c750409 url:string photo:Photo document:Document = FoundGif;
//
func (m *FoundGif) To_FoundGifCached() *TLFoundGifCached {
	return &TLFoundGifCached{
		Data2: m,
	}
}

//  foundGif#162ecc1f url:string thumb_url:string content_url:string content_type:string w:int h:int = FoundGif;
//
func (m *TLFoundGif) To_FoundGif() *FoundGif {
	m.Data2.Cmd = Cmd_FoundGif
	return m.Data2
}

//  foundGifCached#9c750409 url:string photo:Photo document:Document = FoundGif;
//
func (m *TLFoundGifCached) To_FoundGif() *FoundGif {
	m.Data2.Cmd = Cmd_FoundGifCached
	return m.Data2
}

//  + TL_FutureSalt
//

//  future_salt#0949d9dc valid_since:int valid_until:int salt:long = FutureSalt;
//
func (m *FutureSalt) To_FutureSalt() *TLFutureSalt {
	return &TLFutureSalt{
		Data2: m,
	}
}

//  future_salt#0949d9dc valid_since:int valid_until:int salt:long = FutureSalt;
//
func (m *TLFutureSalt) To_FutureSalt() *FutureSalt {
	m.Data2.Cmd = Cmd_FutureSalt
	return m.Data2
}

//  + TL_FutureSalts
//

//  future_salts#ae500895 req_msg_id:long now:int salts:vector<future_salt> = FutureSalts;
//
func (m *FutureSalts) To_FutureSalts() *TLFutureSalts {
	return &TLFutureSalts{
		Data2: m,
	}
}

//  future_salts#ae500895 req_msg_id:long now:int salts:vector<future_salt> = FutureSalts;
//
func (m *TLFutureSalts) To_FutureSalts() *FutureSalts {
	m.Data2.Cmd = Cmd_FutureSalts
	return m.Data2
}

//  + TL_Game
//

//  game#bdf9653b flags:# id:long access_hash:long short_name:string title:string description:string photo:Photo document:flags.0?Document = Game;
//
func (m *Game) To_Game() *TLGame {
	return &TLGame{
		Data2: m,
	}
}

//  game#bdf9653b flags:# id:long access_hash:long short_name:string title:string description:string photo:Photo document:flags.0?Document = Game;
//
func (m *TLGame) To_Game() *Game {
	m.Data2.Cmd = Cmd_Game
	return m.Data2
}

//  + TL_GeoPointEmpty
//  + TL_GeoPoint
//

//  geoPointEmpty#1117dd5f = GeoPoint;
//
func (m *GeoPoint) To_GeoPointEmpty() *TLGeoPointEmpty {
	return &TLGeoPointEmpty{
		Data2: m,
	}
}

//  geoPoint#2049d70c long:double lat:double = GeoPoint;
//  geoPoint#296f104 long:double lat:double access_hash:long = GeoPoint;
//  geoPoint#b2a2f663 flags:# long:double lat:double access_hash:long accuracy_radius:flags.0?int = GeoPoint;
//
func (m *GeoPoint) To_GeoPoint() *TLGeoPoint {
	return &TLGeoPoint{
		Data2: m,
	}
}

//  geoPointEmpty#1117dd5f = GeoPoint;
//
func (m *TLGeoPointEmpty) To_GeoPoint() *GeoPoint {
	m.Data2.Cmd = Cmd_GeoPointEmpty
	return m.Data2
}

//  geoPoint#2049d70c long:double lat:double = GeoPoint;
//  geoPoint#296f104 long:double lat:double access_hash:long = GeoPoint;
//  geoPoint#b2a2f663 flags:# long:double lat:double access_hash:long accuracy_radius:flags.0?int = GeoPoint;
//
func (m *TLGeoPoint) To_GeoPoint() *GeoPoint {
	m.Data2.Cmd = Cmd_GeoPoint
	return m.Data2
}

//  + TL_GlobalPrivacySettings
//

//  globalPrivacySettings#bea2f424 flags:# archive_and_mute_new_noncontact_peers:flags.0?Bool = GlobalPrivacySettings;
//
func (m *GlobalPrivacySettings) To_GlobalPrivacySettings() *TLGlobalPrivacySettings {
	return &TLGlobalPrivacySettings{
		Data2: m,
	}
}

//  globalPrivacySettings#bea2f424 flags:# archive_and_mute_new_noncontact_peers:flags.0?Bool = GlobalPrivacySettings;
//
func (m *TLGlobalPrivacySettings) To_GlobalPrivacySettings() *GlobalPrivacySettings {
	m.Data2.Cmd = Cmd_GlobalPrivacySettings
	return m.Data2
}

//  + TL_GroupCallDiscarded
//  + TL_GroupCall
//

//  groupCallDiscarded#7780bcb4 id:long access_hash:long duration:int = GroupCall;
//
func (m *GroupCall) To_GroupCallDiscarded() *TLGroupCallDiscarded {
	return &TLGroupCallDiscarded{
		Data2: m,
	}
}

//  groupCall#55903081 flags:# join_muted:flags.1?true can_change_join_muted:flags.2?true id:long access_hash:long participants_count:int params:flags.0?DataJSON version:int = GroupCall;
//
func (m *GroupCall) To_GroupCall() *TLGroupCall {
	return &TLGroupCall{
		Data2: m,
	}
}

//  groupCallDiscarded#7780bcb4 id:long access_hash:long duration:int = GroupCall;
//
func (m *TLGroupCallDiscarded) To_GroupCall() *GroupCall {
	m.Data2.Cmd = Cmd_GroupCallDiscarded
	return m.Data2
}

//  groupCall#55903081 flags:# join_muted:flags.1?true can_change_join_muted:flags.2?true id:long access_hash:long participants_count:int params:flags.0?DataJSON version:int = GroupCall;
//
func (m *TLGroupCall) To_GroupCall() *GroupCall {
	m.Data2.Cmd = Cmd_GroupCall
	return m.Data2
}

//  + TL_GroupCallParticipant
//

//  groupCallParticipant#56b087c9 flags:# muted:flags.0?true left:flags.1?true can_self_unmute:flags.2?true just_joined:flags.4?true versioned:flags.5?true user_id:int date:int active_date:flags.3?int source:int = GroupCallParticipant;
//
func (m *GroupCallParticipant) To_GroupCallParticipant() *TLGroupCallParticipant {
	return &TLGroupCallParticipant{
		Data2: m,
	}
}

//  groupCallParticipant#56b087c9 flags:# muted:flags.0?true left:flags.1?true can_self_unmute:flags.2?true just_joined:flags.4?true versioned:flags.5?true user_id:int date:int active_date:flags.3?int source:int = GroupCallParticipant;
//
func (m *TLGroupCallParticipant) To_GroupCallParticipant() *GroupCallParticipant {
	m.Data2.Cmd = Cmd_GroupCallParticipant
	return m.Data2
}

//  + TL_HelpAppChangelogEmpty
//  + TL_HelpAppChangelog
//

//  help.appChangelogEmpty#af7e0394 = help.AppChangelog;
//
func (m *Help_AppChangelog) To_HelpAppChangelogEmpty() *TLHelpAppChangelogEmpty {
	return &TLHelpAppChangelogEmpty{
		Data2: m,
	}
}

//  help.appChangelog#4668e6bd text:string = help.AppChangelog;
//
func (m *Help_AppChangelog) To_HelpAppChangelog() *TLHelpAppChangelog {
	return &TLHelpAppChangelog{
		Data2: m,
	}
}

//  help.appChangelogEmpty#af7e0394 = help.AppChangelog;
//
func (m *TLHelpAppChangelogEmpty) To_Help_AppChangelog() *Help_AppChangelog {
	m.Data2.Cmd = Cmd_HelpAppChangelogEmpty
	return m.Data2
}

//  help.appChangelog#4668e6bd text:string = help.AppChangelog;
//
func (m *TLHelpAppChangelog) To_Help_AppChangelog() *Help_AppChangelog {
	m.Data2.Cmd = Cmd_HelpAppChangelog
	return m.Data2
}

//  + TL_HelpAppUpdate
//  + TL_HelpNoAppUpdate
//

//  help.appUpdate#8987f311 id:int critical:Bool url:string text:string = help.AppUpdate;
//  help.appUpdate#1da7158f flags:# popup:flags.0?true id:int version:string text:string entities:Vector<MessageEntity> document:flags.1?Document url:flags.2?string = help.AppUpdate;
//
func (m *Help_AppUpdate) To_HelpAppUpdate() *TLHelpAppUpdate {
	return &TLHelpAppUpdate{
		Data2: m,
	}
}

//  help.noAppUpdate#c45a6536 = help.AppUpdate;
//
func (m *Help_AppUpdate) To_HelpNoAppUpdate() *TLHelpNoAppUpdate {
	return &TLHelpNoAppUpdate{
		Data2: m,
	}
}

//  help.appUpdate#8987f311 id:int critical:Bool url:string text:string = help.AppUpdate;
//  help.appUpdate#1da7158f flags:# popup:flags.0?true id:int version:string text:string entities:Vector<MessageEntity> document:flags.1?Document url:flags.2?string = help.AppUpdate;
//
func (m *TLHelpAppUpdate) To_Help_AppUpdate() *Help_AppUpdate {
	m.Data2.Cmd = Cmd_HelpAppUpdate
	return m.Data2
}

//  help.noAppUpdate#c45a6536 = help.AppUpdate;
//
func (m *TLHelpNoAppUpdate) To_Help_AppUpdate() *Help_AppUpdate {
	m.Data2.Cmd = Cmd_HelpNoAppUpdate
	return m.Data2
}

//  + TL_HelpConfigSimple
//

//  help.configSimple#d997c3c5 date:int expires:int dc_id:int ip_port_list:Vector<IpPort> = help.ConfigSimple;
//  help.configSimple#5a592a6c date:int expires:int rules:vector<AccessPointRule> = help.ConfigSimple;
//
func (m *Help_ConfigSimple) To_HelpConfigSimple() *TLHelpConfigSimple {
	return &TLHelpConfigSimple{
		Data2: m,
	}
}

//  help.configSimple#d997c3c5 date:int expires:int dc_id:int ip_port_list:Vector<IpPort> = help.ConfigSimple;
//  help.configSimple#5a592a6c date:int expires:int rules:vector<AccessPointRule> = help.ConfigSimple;
//
func (m *TLHelpConfigSimple) To_Help_ConfigSimple() *Help_ConfigSimple {
	m.Data2.Cmd = Cmd_HelpConfigSimple
	return m.Data2
}

//  + TL_HelpCountriesListNotModified
//  + TL_HelpCountriesList
//

//  help.countriesListNotModified#93cc1f32 = help.CountriesList;
//
func (m *Help_CountriesList) To_HelpCountriesListNotModified() *TLHelpCountriesListNotModified {
	return &TLHelpCountriesListNotModified{
		Data2: m,
	}
}

//  help.countriesList#87d0759e countries:Vector<help.Country> hash:int = help.CountriesList;
//
func (m *Help_CountriesList) To_HelpCountriesList() *TLHelpCountriesList {
	return &TLHelpCountriesList{
		Data2: m,
	}
}

//  help.countriesListNotModified#93cc1f32 = help.CountriesList;
//
func (m *TLHelpCountriesListNotModified) To_Help_CountriesList() *Help_CountriesList {
	m.Data2.Cmd = Cmd_HelpCountriesListNotModified
	return m.Data2
}

//  help.countriesList#87d0759e countries:Vector<help.Country> hash:int = help.CountriesList;
//
func (m *TLHelpCountriesList) To_Help_CountriesList() *Help_CountriesList {
	m.Data2.Cmd = Cmd_HelpCountriesList
	return m.Data2
}

//  + TL_HelpCountry
//

//  help.country#c3878e23 flags:# hidden:flags.0?true iso2:string default_name:string name:flags.1?string country_codes:Vector<help.CountryCode> = help.Country;
//
func (m *Help_Country) To_HelpCountry() *TLHelpCountry {
	return &TLHelpCountry{
		Data2: m,
	}
}

//  help.country#c3878e23 flags:# hidden:flags.0?true iso2:string default_name:string name:flags.1?string country_codes:Vector<help.CountryCode> = help.Country;
//
func (m *TLHelpCountry) To_Help_Country() *Help_Country {
	m.Data2.Cmd = Cmd_HelpCountry
	return m.Data2
}

//  + TL_HelpCountryCode
//

//  help.countryCode#4203c5ef flags:# country_code:string prefixes:flags.0?Vector<string> patterns:flags.1?Vector<string> = help.CountryCode;
//
func (m *Help_CountryCode) To_HelpCountryCode() *TLHelpCountryCode {
	return &TLHelpCountryCode{
		Data2: m,
	}
}

//  help.countryCode#4203c5ef flags:# country_code:string prefixes:flags.0?Vector<string> patterns:flags.1?Vector<string> = help.CountryCode;
//
func (m *TLHelpCountryCode) To_Help_CountryCode() *Help_CountryCode {
	m.Data2.Cmd = Cmd_HelpCountryCode
	return m.Data2
}

//  + TL_HelpDeepLinkInfoEmpty
//  + TL_HelpDeepLinkInfo
//

//  help.deepLinkInfoEmpty#66afa166 = help.DeepLinkInfo;
//
func (m *Help_DeepLinkInfo) To_HelpDeepLinkInfoEmpty() *TLHelpDeepLinkInfoEmpty {
	return &TLHelpDeepLinkInfoEmpty{
		Data2: m,
	}
}

//  help.deepLinkInfo#6a4ee832 flags:# update_app:flags.0?true message:string entities:flags.1?Vector<MessageEntity> = help.DeepLinkInfo;
//
func (m *Help_DeepLinkInfo) To_HelpDeepLinkInfo() *TLHelpDeepLinkInfo {
	return &TLHelpDeepLinkInfo{
		Data2: m,
	}
}

//  help.deepLinkInfoEmpty#66afa166 = help.DeepLinkInfo;
//
func (m *TLHelpDeepLinkInfoEmpty) To_Help_DeepLinkInfo() *Help_DeepLinkInfo {
	m.Data2.Cmd = Cmd_HelpDeepLinkInfoEmpty
	return m.Data2
}

//  help.deepLinkInfo#6a4ee832 flags:# update_app:flags.0?true message:string entities:flags.1?Vector<MessageEntity> = help.DeepLinkInfo;
//
func (m *TLHelpDeepLinkInfo) To_Help_DeepLinkInfo() *Help_DeepLinkInfo {
	m.Data2.Cmd = Cmd_HelpDeepLinkInfo
	return m.Data2
}

//  + TL_HelpInviteText
//

//  help.inviteText#18cb9f78 message:string = help.InviteText;
//
func (m *Help_InviteText) To_HelpInviteText() *TLHelpInviteText {
	return &TLHelpInviteText{
		Data2: m,
	}
}

//  help.inviteText#18cb9f78 message:string = help.InviteText;
//
func (m *TLHelpInviteText) To_Help_InviteText() *Help_InviteText {
	m.Data2.Cmd = Cmd_HelpInviteText
	return m.Data2
}

//  + TL_HelpPassportConfigNotModified
//  + TL_HelpPassportConfig
//

//  help.passportConfigNotModified#bfb9f457 = help.PassportConfig;
//
func (m *Help_PassportConfig) To_HelpPassportConfigNotModified() *TLHelpPassportConfigNotModified {
	return &TLHelpPassportConfigNotModified{
		Data2: m,
	}
}

//  help.passportConfig#a098d6af hash:int countries_langs:DataJSON = help.PassportConfig;
//
func (m *Help_PassportConfig) To_HelpPassportConfig() *TLHelpPassportConfig {
	return &TLHelpPassportConfig{
		Data2: m,
	}
}

//  help.passportConfigNotModified#bfb9f457 = help.PassportConfig;
//
func (m *TLHelpPassportConfigNotModified) To_Help_PassportConfig() *Help_PassportConfig {
	m.Data2.Cmd = Cmd_HelpPassportConfigNotModified
	return m.Data2
}

//  help.passportConfig#a098d6af hash:int countries_langs:DataJSON = help.PassportConfig;
//
func (m *TLHelpPassportConfig) To_Help_PassportConfig() *Help_PassportConfig {
	m.Data2.Cmd = Cmd_HelpPassportConfig
	return m.Data2
}

//  + TL_HelpPromoDataEmpty
//  + TL_HelpPromoData
//

//  help.promoDataEmpty#98f6ac75 expires:int = help.PromoData;
//
func (m *Help_PromoData) To_HelpPromoDataEmpty() *TLHelpPromoDataEmpty {
	return &TLHelpPromoDataEmpty{
		Data2: m,
	}
}

//  help.promoData#8f6adccf flags:# proxy:flags.0?true psa:flags.1?true expires:int peer:Peer chats:Vector<Chat> users:Vector<User> psa_message:flags.2?string = help.PromoData;
//  help.promoData#8c39793f flags:# proxy:flags.0?true expires:int peer:Peer chats:Vector<Chat> users:Vector<User> psa_type:flags.1?string psa_message:flags.2?string = help.PromoData;
//
func (m *Help_PromoData) To_HelpPromoData() *TLHelpPromoData {
	return &TLHelpPromoData{
		Data2: m,
	}
}

//  help.promoDataEmpty#98f6ac75 expires:int = help.PromoData;
//
func (m *TLHelpPromoDataEmpty) To_Help_PromoData() *Help_PromoData {
	m.Data2.Cmd = Cmd_HelpPromoDataEmpty
	return m.Data2
}

//  help.promoData#8f6adccf flags:# proxy:flags.0?true psa:flags.1?true expires:int peer:Peer chats:Vector<Chat> users:Vector<User> psa_message:flags.2?string = help.PromoData;
//  help.promoData#8c39793f flags:# proxy:flags.0?true expires:int peer:Peer chats:Vector<Chat> users:Vector<User> psa_type:flags.1?string psa_message:flags.2?string = help.PromoData;
//
func (m *TLHelpPromoData) To_Help_PromoData() *Help_PromoData {
	m.Data2.Cmd = Cmd_HelpPromoData
	return m.Data2
}

//  + TL_HelpProxyDataEmpty
//  + TL_HelpProxyDataPromo
//

//  help.proxyDataEmpty#e09e1fb8 expires:int = help.ProxyData;
//
func (m *Help_ProxyData) To_HelpProxyDataEmpty() *TLHelpProxyDataEmpty {
	return &TLHelpProxyDataEmpty{
		Data2: m,
	}
}

//  help.proxyDataPromo#2bf7ee23 expires:int peer:Peer chats:Vector<Chat> users:Vector<User> = help.ProxyData;
//
func (m *Help_ProxyData) To_HelpProxyDataPromo() *TLHelpProxyDataPromo {
	return &TLHelpProxyDataPromo{
		Data2: m,
	}
}

//  help.proxyDataEmpty#e09e1fb8 expires:int = help.ProxyData;
//
func (m *TLHelpProxyDataEmpty) To_Help_ProxyData() *Help_ProxyData {
	m.Data2.Cmd = Cmd_HelpProxyDataEmpty
	return m.Data2
}

//  help.proxyDataPromo#2bf7ee23 expires:int peer:Peer chats:Vector<Chat> users:Vector<User> = help.ProxyData;
//
func (m *TLHelpProxyDataPromo) To_Help_ProxyData() *Help_ProxyData {
	m.Data2.Cmd = Cmd_HelpProxyDataPromo
	return m.Data2
}

//  + TL_HelpRecentMeUrls
//

//  help.recentMeUrls#e0310d7 urls:Vector<RecentMeUrl> chats:Vector<Chat> users:Vector<User> = help.RecentMeUrls;
//
func (m *Help_RecentMeUrls) To_HelpRecentMeUrls() *TLHelpRecentMeUrls {
	return &TLHelpRecentMeUrls{
		Data2: m,
	}
}

//  help.recentMeUrls#e0310d7 urls:Vector<RecentMeUrl> chats:Vector<Chat> users:Vector<User> = help.RecentMeUrls;
//
func (m *TLHelpRecentMeUrls) To_Help_RecentMeUrls() *Help_RecentMeUrls {
	m.Data2.Cmd = Cmd_HelpRecentMeUrls
	return m.Data2
}

//  + TL_HelpSupport
//

//  help.support#17c6b5f6 phone_number:string user:User = help.Support;
//
func (m *Help_Support) To_HelpSupport() *TLHelpSupport {
	return &TLHelpSupport{
		Data2: m,
	}
}

//  help.support#17c6b5f6 phone_number:string user:User = help.Support;
//
func (m *TLHelpSupport) To_Help_Support() *Help_Support {
	m.Data2.Cmd = Cmd_HelpSupport
	return m.Data2
}

//  + TL_HelpSupportName
//

//  help.supportName#8c05f1c9 name:string = help.SupportName;
//
func (m *Help_SupportName) To_HelpSupportName() *TLHelpSupportName {
	return &TLHelpSupportName{
		Data2: m,
	}
}

//  help.supportName#8c05f1c9 name:string = help.SupportName;
//
func (m *TLHelpSupportName) To_Help_SupportName() *Help_SupportName {
	m.Data2.Cmd = Cmd_HelpSupportName
	return m.Data2
}

//  + TL_HelpTermsOfService
//

//  help.termsOfService#f1ee3e90 text:string = help.TermsOfService;
//  help.termsOfService#780a0310 flags:# popup:flags.0?true id:DataJSON text:string entities:Vector<MessageEntity> min_age_confirm:flags.1?int = help.TermsOfService;
//
func (m *Help_TermsOfService) To_HelpTermsOfService() *TLHelpTermsOfService {
	return &TLHelpTermsOfService{
		Data2: m,
	}
}

//  help.termsOfService#f1ee3e90 text:string = help.TermsOfService;
//  help.termsOfService#780a0310 flags:# popup:flags.0?true id:DataJSON text:string entities:Vector<MessageEntity> min_age_confirm:flags.1?int = help.TermsOfService;
//
func (m *TLHelpTermsOfService) To_Help_TermsOfService() *Help_TermsOfService {
	m.Data2.Cmd = Cmd_HelpTermsOfService
	return m.Data2
}

//  + TL_HelpTermsOfServiceUpdateEmpty
//  + TL_HelpTermsOfServiceUpdate
//

//  help.termsOfServiceUpdateEmpty#e3309f7f expires:int = help.TermsOfServiceUpdate;
//
func (m *Help_TermsOfServiceUpdate) To_HelpTermsOfServiceUpdateEmpty() *TLHelpTermsOfServiceUpdateEmpty {
	return &TLHelpTermsOfServiceUpdateEmpty{
		Data2: m,
	}
}

//  help.termsOfServiceUpdate#28ecf961 expires:int terms_of_service:help.TermsOfService = help.TermsOfServiceUpdate;
//
func (m *Help_TermsOfServiceUpdate) To_HelpTermsOfServiceUpdate() *TLHelpTermsOfServiceUpdate {
	return &TLHelpTermsOfServiceUpdate{
		Data2: m,
	}
}

//  help.termsOfServiceUpdateEmpty#e3309f7f expires:int = help.TermsOfServiceUpdate;
//
func (m *TLHelpTermsOfServiceUpdateEmpty) To_Help_TermsOfServiceUpdate() *Help_TermsOfServiceUpdate {
	m.Data2.Cmd = Cmd_HelpTermsOfServiceUpdateEmpty
	return m.Data2
}

//  help.termsOfServiceUpdate#28ecf961 expires:int terms_of_service:help.TermsOfService = help.TermsOfServiceUpdate;
//
func (m *TLHelpTermsOfServiceUpdate) To_Help_TermsOfServiceUpdate() *Help_TermsOfServiceUpdate {
	m.Data2.Cmd = Cmd_HelpTermsOfServiceUpdate
	return m.Data2
}

//  + TL_HelpUserInfoEmpty
//  + TL_HelpUserInfo
//

//  help.userInfoEmpty#f3ae2eed = help.UserInfo;
//
func (m *Help_UserInfo) To_HelpUserInfoEmpty() *TLHelpUserInfoEmpty {
	return &TLHelpUserInfoEmpty{
		Data2: m,
	}
}

//  help.userInfo#1eb3758 message:string entities:Vector<MessageEntity> author:string date:int = help.UserInfo;
//
func (m *Help_UserInfo) To_HelpUserInfo() *TLHelpUserInfo {
	return &TLHelpUserInfo{
		Data2: m,
	}
}

//  help.userInfoEmpty#f3ae2eed = help.UserInfo;
//
func (m *TLHelpUserInfoEmpty) To_Help_UserInfo() *Help_UserInfo {
	m.Data2.Cmd = Cmd_HelpUserInfoEmpty
	return m.Data2
}

//  help.userInfo#1eb3758 message:string entities:Vector<MessageEntity> author:string date:int = help.UserInfo;
//
func (m *TLHelpUserInfo) To_Help_UserInfo() *Help_UserInfo {
	m.Data2.Cmd = Cmd_HelpUserInfo
	return m.Data2
}

//  + TL_HighScore
//

//  highScore#58fffcd0 pos:int user_id:int score:int = HighScore;
//
func (m *HighScore) To_HighScore() *TLHighScore {
	return &TLHighScore{
		Data2: m,
	}
}

//  highScore#58fffcd0 pos:int user_id:int score:int = HighScore;
//
func (m *TLHighScore) To_HighScore() *HighScore {
	m.Data2.Cmd = Cmd_HighScore
	return m.Data2
}

//  + TL_HttpWait
//

//  http_wait#9299359f max_delay:int wait_after:int max_wait:int = HttpWait;
//
func (m *HttpWait) To_HttpWait() *TLHttpWait {
	return &TLHttpWait{
		Data2: m,
	}
}

//  http_wait#9299359f max_delay:int wait_after:int max_wait:int = HttpWait;
//
func (m *TLHttpWait) To_HttpWait() *HttpWait {
	m.Data2.Cmd = Cmd_HttpWait
	return m.Data2
}

//  + TL_ImportedContact
//

//  importedContact#d0028438 user_id:int client_id:long = ImportedContact;
//
func (m *ImportedContact) To_ImportedContact() *TLImportedContact {
	return &TLImportedContact{
		Data2: m,
	}
}

//  importedContact#d0028438 user_id:int client_id:long = ImportedContact;
//
func (m *TLImportedContact) To_ImportedContact() *ImportedContact {
	m.Data2.Cmd = Cmd_ImportedContact
	return m.Data2
}

//  + TL_InlineBotSwitchPM
//

//  inlineBotSwitchPM#3c20629f text:string start_param:string = InlineBotSwitchPM;
//
func (m *InlineBotSwitchPM) To_InlineBotSwitchPM() *TLInlineBotSwitchPM {
	return &TLInlineBotSwitchPM{
		Data2: m,
	}
}

//  inlineBotSwitchPM#3c20629f text:string start_param:string = InlineBotSwitchPM;
//
func (m *TLInlineBotSwitchPM) To_InlineBotSwitchPM() *InlineBotSwitchPM {
	m.Data2.Cmd = Cmd_InlineBotSwitchPM
	return m.Data2
}

//  + TL_InlineQueryPeerTypeSameBotPM
//  + TL_InlineQueryPeerTypePM
//  + TL_InlineQueryPeerTypeChat
//  + TL_InlineQueryPeerTypeMegagroup
//  + TL_InlineQueryPeerTypeBroadcast
//

//  inlineQueryPeerTypeSameBotPM#3081ed9d = InlineQueryPeerType;
//
func (m *InlineQueryPeerType) To_InlineQueryPeerTypeSameBotPM() *TLInlineQueryPeerTypeSameBotPM {
	return &TLInlineQueryPeerTypeSameBotPM{
		Data2: m,
	}
}

//  inlineQueryPeerTypePM#833c0fac = InlineQueryPeerType;
//
func (m *InlineQueryPeerType) To_InlineQueryPeerTypePM() *TLInlineQueryPeerTypePM {
	return &TLInlineQueryPeerTypePM{
		Data2: m,
	}
}

//  inlineQueryPeerTypeChat#d766c50a = InlineQueryPeerType;
//
func (m *InlineQueryPeerType) To_InlineQueryPeerTypeChat() *TLInlineQueryPeerTypeChat {
	return &TLInlineQueryPeerTypeChat{
		Data2: m,
	}
}

//  inlineQueryPeerTypeMegagroup#5ec4be43 = InlineQueryPeerType;
//
func (m *InlineQueryPeerType) To_InlineQueryPeerTypeMegagroup() *TLInlineQueryPeerTypeMegagroup {
	return &TLInlineQueryPeerTypeMegagroup{
		Data2: m,
	}
}

//  inlineQueryPeerTypeBroadcast#6334ee9a = InlineQueryPeerType;
//
func (m *InlineQueryPeerType) To_InlineQueryPeerTypeBroadcast() *TLInlineQueryPeerTypeBroadcast {
	return &TLInlineQueryPeerTypeBroadcast{
		Data2: m,
	}
}

//  inlineQueryPeerTypeSameBotPM#3081ed9d = InlineQueryPeerType;
//
func (m *TLInlineQueryPeerTypeSameBotPM) To_InlineQueryPeerType() *InlineQueryPeerType {
	m.Data2.Cmd = Cmd_InlineQueryPeerTypeSameBotPM
	return m.Data2
}

//  inlineQueryPeerTypePM#833c0fac = InlineQueryPeerType;
//
func (m *TLInlineQueryPeerTypePM) To_InlineQueryPeerType() *InlineQueryPeerType {
	m.Data2.Cmd = Cmd_InlineQueryPeerTypePM
	return m.Data2
}

//  inlineQueryPeerTypeChat#d766c50a = InlineQueryPeerType;
//
func (m *TLInlineQueryPeerTypeChat) To_InlineQueryPeerType() *InlineQueryPeerType {
	m.Data2.Cmd = Cmd_InlineQueryPeerTypeChat
	return m.Data2
}

//  inlineQueryPeerTypeMegagroup#5ec4be43 = InlineQueryPeerType;
//
func (m *TLInlineQueryPeerTypeMegagroup) To_InlineQueryPeerType() *InlineQueryPeerType {
	m.Data2.Cmd = Cmd_InlineQueryPeerTypeMegagroup
	return m.Data2
}

//  inlineQueryPeerTypeBroadcast#6334ee9a = InlineQueryPeerType;
//
func (m *TLInlineQueryPeerTypeBroadcast) To_InlineQueryPeerType() *InlineQueryPeerType {
	m.Data2.Cmd = Cmd_InlineQueryPeerTypeBroadcast
	return m.Data2
}

//  + TL_InputAppEvent
//

//  inputAppEvent#770656a8 time:double type:string peer:long data:string = InputAppEvent;
//  inputAppEvent#1d1b1245 time:double type:string peer:long data:JSONValue = InputAppEvent;
//
func (m *InputAppEvent) To_InputAppEvent() *TLInputAppEvent {
	return &TLInputAppEvent{
		Data2: m,
	}
}

//  inputAppEvent#770656a8 time:double type:string peer:long data:string = InputAppEvent;
//  inputAppEvent#1d1b1245 time:double type:string peer:long data:JSONValue = InputAppEvent;
//
func (m *TLInputAppEvent) To_InputAppEvent() *InputAppEvent {
	m.Data2.Cmd = Cmd_InputAppEvent
	return m.Data2
}

//  + TL_InputBotInlineMessageMediaAuto
//  + TL_InputBotInlineMessageText
//  + TL_InputBotInlineMessageMediaGeo
//  + TL_InputBotInlineMessageMediaVenue
//  + TL_InputBotInlineMessageMediaContact
//  + TL_InputBotInlineMessageGame
//

//  inputBotInlineMessageMediaAuto#292fed13 flags:# caption:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//  inputBotInlineMessageMediaAuto#3380c786 flags:# message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *InputBotInlineMessage) To_InputBotInlineMessageMediaAuto() *TLInputBotInlineMessageMediaAuto {
	return &TLInputBotInlineMessageMediaAuto{
		Data2: m,
	}
}

//  inputBotInlineMessageText#3dcd7a87 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *InputBotInlineMessage) To_InputBotInlineMessageText() *TLInputBotInlineMessageText {
	return &TLInputBotInlineMessageText{
		Data2: m,
	}
}

//  inputBotInlineMessageMediaGeo#f4a59de1 flags:# geo_point:InputGeoPoint reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//  inputBotInlineMessageMediaGeo#c1b15d65 flags:# geo_point:InputGeoPoint period:int reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//  inputBotInlineMessageMediaGeo#96929a85 flags:# geo_point:InputGeoPoint heading:flags.0?int period:flags.1?int proximity_notification_radius:flags.3?int reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *InputBotInlineMessage) To_InputBotInlineMessageMediaGeo() *TLInputBotInlineMessageMediaGeo {
	return &TLInputBotInlineMessageMediaGeo{
		Data2: m,
	}
}

//  inputBotInlineMessageMediaVenue#aaafadc8 flags:# geo_point:InputGeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//  inputBotInlineMessageMediaVenue#417bbf11 flags:# geo_point:InputGeoPoint title:string address:string provider:string venue_id:string venue_type:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *InputBotInlineMessage) To_InputBotInlineMessageMediaVenue() *TLInputBotInlineMessageMediaVenue {
	return &TLInputBotInlineMessageMediaVenue{
		Data2: m,
	}
}

//  inputBotInlineMessageMediaContact#2daf01a7 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//  inputBotInlineMessageMediaContact#a6edbffd flags:# phone_number:string first_name:string last_name:string vcard:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *InputBotInlineMessage) To_InputBotInlineMessageMediaContact() *TLInputBotInlineMessageMediaContact {
	return &TLInputBotInlineMessageMediaContact{
		Data2: m,
	}
}

//  inputBotInlineMessageGame#4b425864 flags:# reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *InputBotInlineMessage) To_InputBotInlineMessageGame() *TLInputBotInlineMessageGame {
	return &TLInputBotInlineMessageGame{
		Data2: m,
	}
}

//  inputBotInlineMessageMediaAuto#292fed13 flags:# caption:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//  inputBotInlineMessageMediaAuto#3380c786 flags:# message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *TLInputBotInlineMessageMediaAuto) To_InputBotInlineMessage() *InputBotInlineMessage {
	m.Data2.Cmd = Cmd_InputBotInlineMessageMediaAuto
	return m.Data2
}

//  inputBotInlineMessageText#3dcd7a87 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *TLInputBotInlineMessageText) To_InputBotInlineMessage() *InputBotInlineMessage {
	m.Data2.Cmd = Cmd_InputBotInlineMessageText
	return m.Data2
}

//  inputBotInlineMessageMediaGeo#f4a59de1 flags:# geo_point:InputGeoPoint reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//  inputBotInlineMessageMediaGeo#c1b15d65 flags:# geo_point:InputGeoPoint period:int reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//  inputBotInlineMessageMediaGeo#96929a85 flags:# geo_point:InputGeoPoint heading:flags.0?int period:flags.1?int proximity_notification_radius:flags.3?int reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *TLInputBotInlineMessageMediaGeo) To_InputBotInlineMessage() *InputBotInlineMessage {
	m.Data2.Cmd = Cmd_InputBotInlineMessageMediaGeo
	return m.Data2
}

//  inputBotInlineMessageMediaVenue#aaafadc8 flags:# geo_point:InputGeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//  inputBotInlineMessageMediaVenue#417bbf11 flags:# geo_point:InputGeoPoint title:string address:string provider:string venue_id:string venue_type:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *TLInputBotInlineMessageMediaVenue) To_InputBotInlineMessage() *InputBotInlineMessage {
	m.Data2.Cmd = Cmd_InputBotInlineMessageMediaVenue
	return m.Data2
}

//  inputBotInlineMessageMediaContact#2daf01a7 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//  inputBotInlineMessageMediaContact#a6edbffd flags:# phone_number:string first_name:string last_name:string vcard:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *TLInputBotInlineMessageMediaContact) To_InputBotInlineMessage() *InputBotInlineMessage {
	m.Data2.Cmd = Cmd_InputBotInlineMessageMediaContact
	return m.Data2
}

//  inputBotInlineMessageGame#4b425864 flags:# reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
//
func (m *TLInputBotInlineMessageGame) To_InputBotInlineMessage() *InputBotInlineMessage {
	m.Data2.Cmd = Cmd_InputBotInlineMessageGame
	return m.Data2
}

//  + TL_InputBotInlineMessageID
//

//  inputBotInlineMessageID#890c3d89 dc_id:int id:long access_hash:long = InputBotInlineMessageID;
//
func (m *InputBotInlineMessageID) To_InputBotInlineMessageID() *TLInputBotInlineMessageID {
	return &TLInputBotInlineMessageID{
		Data2: m,
	}
}

//  inputBotInlineMessageID#890c3d89 dc_id:int id:long access_hash:long = InputBotInlineMessageID;
//
func (m *TLInputBotInlineMessageID) To_InputBotInlineMessageID() *InputBotInlineMessageID {
	m.Data2.Cmd = Cmd_InputBotInlineMessageID
	return m.Data2
}

//  + TL_InputBotInlineResult
//  + TL_InputBotInlineResultPhoto
//  + TL_InputBotInlineResultDocument
//  + TL_InputBotInlineResultGame
//

//  inputBotInlineResult#2cbbe15a flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:InputBotInlineMessage = InputBotInlineResult;
//  inputBotInlineResult#88bf9319 flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb:flags.4?InputWebDocument content:flags.5?InputWebDocument send_message:InputBotInlineMessage = InputBotInlineResult;
//
func (m *InputBotInlineResult) To_InputBotInlineResult() *TLInputBotInlineResult {
	return &TLInputBotInlineResult{
		Data2: m,
	}
}

//  inputBotInlineResultPhoto#a8d864a7 id:string type:string photo:InputPhoto send_message:InputBotInlineMessage = InputBotInlineResult;
//
func (m *InputBotInlineResult) To_InputBotInlineResultPhoto() *TLInputBotInlineResultPhoto {
	return &TLInputBotInlineResultPhoto{
		Data2: m,
	}
}

//  inputBotInlineResultDocument#fff8fdc4 flags:# id:string type:string title:flags.1?string description:flags.2?string document:InputDocument send_message:InputBotInlineMessage = InputBotInlineResult;
//
func (m *InputBotInlineResult) To_InputBotInlineResultDocument() *TLInputBotInlineResultDocument {
	return &TLInputBotInlineResultDocument{
		Data2: m,
	}
}

//  inputBotInlineResultGame#4fa417f2 id:string short_name:string send_message:InputBotInlineMessage = InputBotInlineResult;
//
func (m *InputBotInlineResult) To_InputBotInlineResultGame() *TLInputBotInlineResultGame {
	return &TLInputBotInlineResultGame{
		Data2: m,
	}
}

//  inputBotInlineResult#2cbbe15a flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:InputBotInlineMessage = InputBotInlineResult;
//  inputBotInlineResult#88bf9319 flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb:flags.4?InputWebDocument content:flags.5?InputWebDocument send_message:InputBotInlineMessage = InputBotInlineResult;
//
func (m *TLInputBotInlineResult) To_InputBotInlineResult() *InputBotInlineResult {
	m.Data2.Cmd = Cmd_InputBotInlineResult
	return m.Data2
}

//  inputBotInlineResultPhoto#a8d864a7 id:string type:string photo:InputPhoto send_message:InputBotInlineMessage = InputBotInlineResult;
//
func (m *TLInputBotInlineResultPhoto) To_InputBotInlineResult() *InputBotInlineResult {
	m.Data2.Cmd = Cmd_InputBotInlineResultPhoto
	return m.Data2
}

//  inputBotInlineResultDocument#fff8fdc4 flags:# id:string type:string title:flags.1?string description:flags.2?string document:InputDocument send_message:InputBotInlineMessage = InputBotInlineResult;
//
func (m *TLInputBotInlineResultDocument) To_InputBotInlineResult() *InputBotInlineResult {
	m.Data2.Cmd = Cmd_InputBotInlineResultDocument
	return m.Data2
}

//  inputBotInlineResultGame#4fa417f2 id:string short_name:string send_message:InputBotInlineMessage = InputBotInlineResult;
//
func (m *TLInputBotInlineResultGame) To_InputBotInlineResult() *InputBotInlineResult {
	m.Data2.Cmd = Cmd_InputBotInlineResultGame
	return m.Data2
}

//  + TL_InputChannelEmpty
//  + TL_InputChannel
//  + TL_InputChannelFromMessage
//

//  inputChannelEmpty#ee8c1e86 = InputChannel;
//
func (m *InputChannel) To_InputChannelEmpty() *TLInputChannelEmpty {
	return &TLInputChannelEmpty{
		Data2: m,
	}
}

//  inputChannel#afeb712e channel_id:int access_hash:long = InputChannel;
//
func (m *InputChannel) To_InputChannel() *TLInputChannel {
	return &TLInputChannel{
		Data2: m,
	}
}

//  inputChannelFromMessage#2a286531 peer:InputPeer msg_id:int channel_id:int = InputChannel;
//
func (m *InputChannel) To_InputChannelFromMessage() *TLInputChannelFromMessage {
	return &TLInputChannelFromMessage{
		Data2: m,
	}
}

//  inputChannelEmpty#ee8c1e86 = InputChannel;
//
func (m *TLInputChannelEmpty) To_InputChannel() *InputChannel {
	m.Data2.Cmd = Cmd_InputChannelEmpty
	return m.Data2
}

//  inputChannel#afeb712e channel_id:int access_hash:long = InputChannel;
//
func (m *TLInputChannel) To_InputChannel() *InputChannel {
	m.Data2.Cmd = Cmd_InputChannel
	return m.Data2
}

//  inputChannelFromMessage#2a286531 peer:InputPeer msg_id:int channel_id:int = InputChannel;
//
func (m *TLInputChannelFromMessage) To_InputChannel() *InputChannel {
	m.Data2.Cmd = Cmd_InputChannelFromMessage
	return m.Data2
}

//  + TL_InputChatPhotoEmpty
//  + TL_InputChatUploadedPhoto
//  + TL_InputChatPhoto
//

//  inputChatPhotoEmpty#1ca48f57 = InputChatPhoto;
//
func (m *InputChatPhoto) To_InputChatPhotoEmpty() *TLInputChatPhotoEmpty {
	return &TLInputChatPhotoEmpty{
		Data2: m,
	}
}

//  inputChatUploadedPhoto#927c55b4 file:InputFile = InputChatPhoto;
//  inputChatUploadedPhoto#94254732 file:InputFile crop:InputPhotoCrop = InputChatPhoto;
//  inputChatUploadedPhoto#c642724e flags:# file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = InputChatPhoto;
//
func (m *InputChatPhoto) To_InputChatUploadedPhoto() *TLInputChatUploadedPhoto {
	return &TLInputChatUploadedPhoto{
		Data2: m,
	}
}

//  inputChatPhoto#8953ad37 id:InputPhoto = InputChatPhoto;
//  inputChatPhoto#b2e1bf08 id:InputPhoto crop:InputPhotoCrop = InputChatPhoto;
//
func (m *InputChatPhoto) To_InputChatPhoto() *TLInputChatPhoto {
	return &TLInputChatPhoto{
		Data2: m,
	}
}

//  inputChatPhotoEmpty#1ca48f57 = InputChatPhoto;
//
func (m *TLInputChatPhotoEmpty) To_InputChatPhoto() *InputChatPhoto {
	m.Data2.Cmd = Cmd_InputChatPhotoEmpty
	return m.Data2
}

//  inputChatUploadedPhoto#927c55b4 file:InputFile = InputChatPhoto;
//  inputChatUploadedPhoto#94254732 file:InputFile crop:InputPhotoCrop = InputChatPhoto;
//  inputChatUploadedPhoto#c642724e flags:# file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = InputChatPhoto;
//
func (m *TLInputChatUploadedPhoto) To_InputChatPhoto() *InputChatPhoto {
	m.Data2.Cmd = Cmd_InputChatUploadedPhoto
	return m.Data2
}

//  inputChatPhoto#8953ad37 id:InputPhoto = InputChatPhoto;
//  inputChatPhoto#b2e1bf08 id:InputPhoto crop:InputPhotoCrop = InputChatPhoto;
//
func (m *TLInputChatPhoto) To_InputChatPhoto() *InputChatPhoto {
	m.Data2.Cmd = Cmd_InputChatPhoto
	return m.Data2
}

//  + TL_InputCheckPasswordEmpty
//  + TL_InputCheckPasswordSRP
//

//  inputCheckPasswordEmpty#9880f658 = InputCheckPasswordSRP;
//
func (m *InputCheckPasswordSRP) To_InputCheckPasswordEmpty() *TLInputCheckPasswordEmpty {
	return &TLInputCheckPasswordEmpty{
		Data2: m,
	}
}

//  inputCheckPasswordSRP#d27ff082 srp_id:long A:bytes M1:bytes = InputCheckPasswordSRP;
//
func (m *InputCheckPasswordSRP) To_InputCheckPasswordSRP() *TLInputCheckPasswordSRP {
	return &TLInputCheckPasswordSRP{
		Data2: m,
	}
}

//  inputCheckPasswordEmpty#9880f658 = InputCheckPasswordSRP;
//
func (m *TLInputCheckPasswordEmpty) To_InputCheckPasswordSRP() *InputCheckPasswordSRP {
	m.Data2.Cmd = Cmd_InputCheckPasswordEmpty
	return m.Data2
}

//  inputCheckPasswordSRP#d27ff082 srp_id:long A:bytes M1:bytes = InputCheckPasswordSRP;
//
func (m *TLInputCheckPasswordSRP) To_InputCheckPasswordSRP() *InputCheckPasswordSRP {
	m.Data2.Cmd = Cmd_InputCheckPasswordSRP
	return m.Data2
}

//  + TL_InputClientProxy
//

//  inputClientProxy#75588b3f address:string port:int = InputClientProxy;
//
func (m *InputClientProxy) To_InputClientProxy() *TLInputClientProxy {
	return &TLInputClientProxy{
		Data2: m,
	}
}

//  inputClientProxy#75588b3f address:string port:int = InputClientProxy;
//
func (m *TLInputClientProxy) To_InputClientProxy() *InputClientProxy {
	m.Data2.Cmd = Cmd_InputClientProxy
	return m.Data2
}

//  + TL_InputPhoneContact
//

//  inputPhoneContact#f392b7f4 client_id:long phone:string first_name:string last_name:string = InputContact;
//
func (m *InputContact) To_InputPhoneContact() *TLInputPhoneContact {
	return &TLInputPhoneContact{
		Data2: m,
	}
}

//  inputPhoneContact#f392b7f4 client_id:long phone:string first_name:string last_name:string = InputContact;
//
func (m *TLInputPhoneContact) To_InputContact() *InputContact {
	m.Data2.Cmd = Cmd_InputPhoneContact
	return m.Data2
}

//  + TL_InputDialogPeer
//  + TL_InputDialogPeerFolder
//

//  inputDialogPeer#fcaafeb7 peer:InputPeer = InputDialogPeer;
//
func (m *InputDialogPeer) To_InputDialogPeer() *TLInputDialogPeer {
	return &TLInputDialogPeer{
		Data2: m,
	}
}

//  inputDialogPeerFolder#64600527 folder_id:int = InputDialogPeer;
//
func (m *InputDialogPeer) To_InputDialogPeerFolder() *TLInputDialogPeerFolder {
	return &TLInputDialogPeerFolder{
		Data2: m,
	}
}

//  inputDialogPeer#fcaafeb7 peer:InputPeer = InputDialogPeer;
//
func (m *TLInputDialogPeer) To_InputDialogPeer() *InputDialogPeer {
	m.Data2.Cmd = Cmd_InputDialogPeer
	return m.Data2
}

//  inputDialogPeerFolder#64600527 folder_id:int = InputDialogPeer;
//
func (m *TLInputDialogPeerFolder) To_InputDialogPeer() *InputDialogPeer {
	m.Data2.Cmd = Cmd_InputDialogPeerFolder
	return m.Data2
}

//  + TL_InputDocumentEmpty
//  + TL_InputDocument
//

//  inputDocumentEmpty#72f0eaae = InputDocument;
//
func (m *InputDocument) To_InputDocumentEmpty() *TLInputDocumentEmpty {
	return &TLInputDocumentEmpty{
		Data2: m,
	}
}

//  inputDocument#18798952 id:long access_hash:long = InputDocument;
//  inputDocument#1abfb575 id:long access_hash:long file_reference:bytes = InputDocument;
//
func (m *InputDocument) To_InputDocument() *TLInputDocument {
	return &TLInputDocument{
		Data2: m,
	}
}

//  inputDocumentEmpty#72f0eaae = InputDocument;
//
func (m *TLInputDocumentEmpty) To_InputDocument() *InputDocument {
	m.Data2.Cmd = Cmd_InputDocumentEmpty
	return m.Data2
}

//  inputDocument#18798952 id:long access_hash:long = InputDocument;
//  inputDocument#1abfb575 id:long access_hash:long file_reference:bytes = InputDocument;
//
func (m *TLInputDocument) To_InputDocument() *InputDocument {
	m.Data2.Cmd = Cmd_InputDocument
	return m.Data2
}

//  + TL_InputEncryptedChat
//

//  inputEncryptedChat#f141b5e1 chat_id:int access_hash:long = InputEncryptedChat;
//
func (m *InputEncryptedChat) To_InputEncryptedChat() *TLInputEncryptedChat {
	return &TLInputEncryptedChat{
		Data2: m,
	}
}

//  inputEncryptedChat#f141b5e1 chat_id:int access_hash:long = InputEncryptedChat;
//
func (m *TLInputEncryptedChat) To_InputEncryptedChat() *InputEncryptedChat {
	m.Data2.Cmd = Cmd_InputEncryptedChat
	return m.Data2
}

//  + TL_InputEncryptedFileEmpty
//  + TL_InputEncryptedFileUploaded
//  + TL_InputEncryptedFile
//  + TL_InputEncryptedFileBigUploaded
//

//  inputEncryptedFileEmpty#1837c364 = InputEncryptedFile;
//
func (m *InputEncryptedFile) To_InputEncryptedFileEmpty() *TLInputEncryptedFileEmpty {
	return &TLInputEncryptedFileEmpty{
		Data2: m,
	}
}

//  inputEncryptedFileUploaded#64bd0306 id:long parts:int md5_checksum:string key_fingerprint:int = InputEncryptedFile;
//
func (m *InputEncryptedFile) To_InputEncryptedFileUploaded() *TLInputEncryptedFileUploaded {
	return &TLInputEncryptedFileUploaded{
		Data2: m,
	}
}

//  inputEncryptedFile#5a17b5e5 id:long access_hash:long = InputEncryptedFile;
//
func (m *InputEncryptedFile) To_InputEncryptedFile() *TLInputEncryptedFile {
	return &TLInputEncryptedFile{
		Data2: m,
	}
}

//  inputEncryptedFileBigUploaded#2dc173c8 id:long parts:int key_fingerprint:int = InputEncryptedFile;
//
func (m *InputEncryptedFile) To_InputEncryptedFileBigUploaded() *TLInputEncryptedFileBigUploaded {
	return &TLInputEncryptedFileBigUploaded{
		Data2: m,
	}
}

//  inputEncryptedFileEmpty#1837c364 = InputEncryptedFile;
//
func (m *TLInputEncryptedFileEmpty) To_InputEncryptedFile() *InputEncryptedFile {
	m.Data2.Cmd = Cmd_InputEncryptedFileEmpty
	return m.Data2
}

//  inputEncryptedFileUploaded#64bd0306 id:long parts:int md5_checksum:string key_fingerprint:int = InputEncryptedFile;
//
func (m *TLInputEncryptedFileUploaded) To_InputEncryptedFile() *InputEncryptedFile {
	m.Data2.Cmd = Cmd_InputEncryptedFileUploaded
	return m.Data2
}

//  inputEncryptedFile#5a17b5e5 id:long access_hash:long = InputEncryptedFile;
//
func (m *TLInputEncryptedFile) To_InputEncryptedFile() *InputEncryptedFile {
	m.Data2.Cmd = Cmd_InputEncryptedFile
	return m.Data2
}

//  inputEncryptedFileBigUploaded#2dc173c8 id:long parts:int key_fingerprint:int = InputEncryptedFile;
//
func (m *TLInputEncryptedFileBigUploaded) To_InputEncryptedFile() *InputEncryptedFile {
	m.Data2.Cmd = Cmd_InputEncryptedFileBigUploaded
	return m.Data2
}

//  + TL_InputFile
//  + TL_InputFileBig
//

//  inputFile#f52ff27f id:long parts:int name:string md5_checksum:string = InputFile;
//
func (m *InputFile) To_InputFile() *TLInputFile {
	return &TLInputFile{
		Data2: m,
	}
}

//  inputFileBig#fa4f0bb5 id:long parts:int name:string = InputFile;
//
func (m *InputFile) To_InputFileBig() *TLInputFileBig {
	return &TLInputFileBig{
		Data2: m,
	}
}

//  inputFile#f52ff27f id:long parts:int name:string md5_checksum:string = InputFile;
//
func (m *TLInputFile) To_InputFile() *InputFile {
	m.Data2.Cmd = Cmd_InputFile
	return m.Data2
}

//  inputFileBig#fa4f0bb5 id:long parts:int name:string = InputFile;
//
func (m *TLInputFileBig) To_InputFile() *InputFile {
	m.Data2.Cmd = Cmd_InputFileBig
	return m.Data2
}

//  + TL_InputFileLocation
//  + TL_InputEncryptedFileLocation
//  + TL_InputDocumentFileLocation
//  + TL_InputSecureFileLocation
//  + TL_InputTakeoutFileLocation
//  + TL_InputPhotoFileLocation
//  + TL_InputPeerPhotoFileLocation
//  + TL_InputStickerSetThumb
//  + TL_InputPhotoLegacyFileLocation
//

//  inputFileLocation#14637196 volume_id:long local_id:int secret:long = InputFileLocation;
//  inputFileLocation#dfdaabe1 volume_id:long local_id:int secret:long file_reference:bytes = InputFileLocation;
//
func (m *InputFileLocation) To_InputFileLocation() *TLInputFileLocation {
	return &TLInputFileLocation{
		Data2: m,
	}
}

//  inputEncryptedFileLocation#f5235d55 id:long access_hash:long = InputFileLocation;
//
func (m *InputFileLocation) To_InputEncryptedFileLocation() *TLInputEncryptedFileLocation {
	return &TLInputEncryptedFileLocation{
		Data2: m,
	}
}

//  inputDocumentFileLocation#430f0724 id:long access_hash:long version:int = InputFileLocation;
//  inputDocumentFileLocation#4e45abe9 id:long access_hash:long = InputFileLocation;
//  inputDocumentFileLocation#196683d9 id:long access_hash:long file_reference:bytes = InputFileLocation;
//  inputDocumentFileLocation#bad07584 id:long access_hash:long file_reference:bytes thumb_size:string = InputFileLocation;
//
func (m *InputFileLocation) To_InputDocumentFileLocation() *TLInputDocumentFileLocation {
	return &TLInputDocumentFileLocation{
		Data2: m,
	}
}

//  inputSecureFileLocation#cbc7ee28 id:long access_hash:long = InputFileLocation;
//
func (m *InputFileLocation) To_InputSecureFileLocation() *TLInputSecureFileLocation {
	return &TLInputSecureFileLocation{
		Data2: m,
	}
}

//  inputTakeoutFileLocation#29be5899 = InputFileLocation;
//
func (m *InputFileLocation) To_InputTakeoutFileLocation() *TLInputTakeoutFileLocation {
	return &TLInputTakeoutFileLocation{
		Data2: m,
	}
}

//  inputPhotoFileLocation#40181ffe id:long access_hash:long file_reference:bytes thumb_size:string = InputFileLocation;
//
func (m *InputFileLocation) To_InputPhotoFileLocation() *TLInputPhotoFileLocation {
	return &TLInputPhotoFileLocation{
		Data2: m,
	}
}

//  inputPeerPhotoFileLocation#27d69997 flags:# big:flags.0?true peer:InputPeer volume_id:long local_id:int = InputFileLocation;
//
func (m *InputFileLocation) To_InputPeerPhotoFileLocation() *TLInputPeerPhotoFileLocation {
	return &TLInputPeerPhotoFileLocation{
		Data2: m,
	}
}

//  inputStickerSetThumb#dbaeae9 stickerset:InputStickerSet volume_id:long local_id:int = InputFileLocation;
//
func (m *InputFileLocation) To_InputStickerSetThumb() *TLInputStickerSetThumb {
	return &TLInputStickerSetThumb{
		Data2: m,
	}
}

//  inputPhotoLegacyFileLocation#d83466f3 id:long access_hash:long file_reference:bytes volume_id:long local_id:int secret:long = InputFileLocation;
//
func (m *InputFileLocation) To_InputPhotoLegacyFileLocation() *TLInputPhotoLegacyFileLocation {
	return &TLInputPhotoLegacyFileLocation{
		Data2: m,
	}
}

//  inputFileLocation#14637196 volume_id:long local_id:int secret:long = InputFileLocation;
//  inputFileLocation#dfdaabe1 volume_id:long local_id:int secret:long file_reference:bytes = InputFileLocation;
//
func (m *TLInputFileLocation) To_InputFileLocation() *InputFileLocation {
	m.Data2.Cmd = Cmd_InputFileLocation
	return m.Data2
}

//  inputEncryptedFileLocation#f5235d55 id:long access_hash:long = InputFileLocation;
//
func (m *TLInputEncryptedFileLocation) To_InputFileLocation() *InputFileLocation {
	m.Data2.Cmd = Cmd_InputEncryptedFileLocation
	return m.Data2
}

//  inputDocumentFileLocation#430f0724 id:long access_hash:long version:int = InputFileLocation;
//  inputDocumentFileLocation#4e45abe9 id:long access_hash:long = InputFileLocation;
//  inputDocumentFileLocation#196683d9 id:long access_hash:long file_reference:bytes = InputFileLocation;
//  inputDocumentFileLocation#bad07584 id:long access_hash:long file_reference:bytes thumb_size:string = InputFileLocation;
//
func (m *TLInputDocumentFileLocation) To_InputFileLocation() *InputFileLocation {
	m.Data2.Cmd = Cmd_InputDocumentFileLocation
	return m.Data2
}

//  inputSecureFileLocation#cbc7ee28 id:long access_hash:long = InputFileLocation;
//
func (m *TLInputSecureFileLocation) To_InputFileLocation() *InputFileLocation {
	m.Data2.Cmd = Cmd_InputSecureFileLocation
	return m.Data2
}

//  inputTakeoutFileLocation#29be5899 = InputFileLocation;
//
func (m *TLInputTakeoutFileLocation) To_InputFileLocation() *InputFileLocation {
	m.Data2.Cmd = Cmd_InputTakeoutFileLocation
	return m.Data2
}

//  inputPhotoFileLocation#40181ffe id:long access_hash:long file_reference:bytes thumb_size:string = InputFileLocation;
//
func (m *TLInputPhotoFileLocation) To_InputFileLocation() *InputFileLocation {
	m.Data2.Cmd = Cmd_InputPhotoFileLocation
	return m.Data2
}

//  inputPeerPhotoFileLocation#27d69997 flags:# big:flags.0?true peer:InputPeer volume_id:long local_id:int = InputFileLocation;
//
func (m *TLInputPeerPhotoFileLocation) To_InputFileLocation() *InputFileLocation {
	m.Data2.Cmd = Cmd_InputPeerPhotoFileLocation
	return m.Data2
}

//  inputStickerSetThumb#dbaeae9 stickerset:InputStickerSet volume_id:long local_id:int = InputFileLocation;
//
func (m *TLInputStickerSetThumb) To_InputFileLocation() *InputFileLocation {
	m.Data2.Cmd = Cmd_InputStickerSetThumb
	return m.Data2
}

//  inputPhotoLegacyFileLocation#d83466f3 id:long access_hash:long file_reference:bytes volume_id:long local_id:int secret:long = InputFileLocation;
//
func (m *TLInputPhotoLegacyFileLocation) To_InputFileLocation() *InputFileLocation {
	m.Data2.Cmd = Cmd_InputPhotoLegacyFileLocation
	return m.Data2
}

//  + TL_InputFolderPeer
//

//  inputFolderPeer#fbd2c296 peer:InputPeer folder_id:int = InputFolderPeer;
//
func (m *InputFolderPeer) To_InputFolderPeer() *TLInputFolderPeer {
	return &TLInputFolderPeer{
		Data2: m,
	}
}

//  inputFolderPeer#fbd2c296 peer:InputPeer folder_id:int = InputFolderPeer;
//
func (m *TLInputFolderPeer) To_InputFolderPeer() *InputFolderPeer {
	m.Data2.Cmd = Cmd_InputFolderPeer
	return m.Data2
}

//  + TL_InputGameID
//  + TL_InputGameShortName
//

//  inputGameID#32c3e77 id:long access_hash:long = InputGame;
//
func (m *InputGame) To_InputGameID() *TLInputGameID {
	return &TLInputGameID{
		Data2: m,
	}
}

//  inputGameShortName#c331e80a bot_id:InputUser short_name:string = InputGame;
//
func (m *InputGame) To_InputGameShortName() *TLInputGameShortName {
	return &TLInputGameShortName{
		Data2: m,
	}
}

//  inputGameID#32c3e77 id:long access_hash:long = InputGame;
//
func (m *TLInputGameID) To_InputGame() *InputGame {
	m.Data2.Cmd = Cmd_InputGameID
	return m.Data2
}

//  inputGameShortName#c331e80a bot_id:InputUser short_name:string = InputGame;
//
func (m *TLInputGameShortName) To_InputGame() *InputGame {
	m.Data2.Cmd = Cmd_InputGameShortName
	return m.Data2
}

//  + TL_InputGeoPointEmpty
//  + TL_InputGeoPoint
//

//  inputGeoPointEmpty#e4c123d6 = InputGeoPoint;
//
func (m *InputGeoPoint) To_InputGeoPointEmpty() *TLInputGeoPointEmpty {
	return &TLInputGeoPointEmpty{
		Data2: m,
	}
}

//  inputGeoPoint#f3b7acc9 lat:double long:double = InputGeoPoint;
//  inputGeoPoint#48222faf flags:# lat:double long:double accuracy_radius:flags.0?int = InputGeoPoint;
//
func (m *InputGeoPoint) To_InputGeoPoint() *TLInputGeoPoint {
	return &TLInputGeoPoint{
		Data2: m,
	}
}

//  inputGeoPointEmpty#e4c123d6 = InputGeoPoint;
//
func (m *TLInputGeoPointEmpty) To_InputGeoPoint() *InputGeoPoint {
	m.Data2.Cmd = Cmd_InputGeoPointEmpty
	return m.Data2
}

//  inputGeoPoint#f3b7acc9 lat:double long:double = InputGeoPoint;
//  inputGeoPoint#48222faf flags:# lat:double long:double accuracy_radius:flags.0?int = InputGeoPoint;
//
func (m *TLInputGeoPoint) To_InputGeoPoint() *InputGeoPoint {
	m.Data2.Cmd = Cmd_InputGeoPoint
	return m.Data2
}

//  + TL_InputGroupCall
//

//  inputGroupCall#d8aa840f id:long access_hash:long = InputGroupCall;
//
func (m *InputGroupCall) To_InputGroupCall() *TLInputGroupCall {
	return &TLInputGroupCall{
		Data2: m,
	}
}

//  inputGroupCall#d8aa840f id:long access_hash:long = InputGroupCall;
//
func (m *TLInputGroupCall) To_InputGroupCall() *InputGroupCall {
	m.Data2.Cmd = Cmd_InputGroupCall
	return m.Data2
}

//  + TL_InputMediaEmpty
//  + TL_InputMediaUploadedPhoto
//  + TL_InputMediaPhoto
//  + TL_InputMediaGeoPoint
//  + TL_InputMediaContact
//  + TL_InputMediaUploadedDocument
//  + TL_InputMediaDocument
//  + TL_InputMediaVenue
//  + TL_InputMediaGifExternal
//  + TL_InputMediaPhotoExternal
//  + TL_InputMediaDocumentExternal
//  + TL_InputMediaGame
//  + TL_InputMediaInvoice
//  + TL_InputMediaUploadedThumbDocument
//  + TL_InputMediaGeoLive
//  + TL_InputMediaPoll
//  + TL_InputMediaDice
//

//  inputMediaEmpty#9664f57f = InputMedia;
//
func (m *InputMedia) To_InputMediaEmpty() *TLInputMediaEmpty {
	return &TLInputMediaEmpty{
		Data2: m,
	}
}

//  inputMediaUploadedPhoto#2f37e231 flags:# file:InputFile caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
//  inputMediaUploadedPhoto#f7aff1c0 file:InputFile caption:string = InputMedia;
//  inputMediaUploadedPhoto#1e287d04 flags:# file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
//
func (m *InputMedia) To_InputMediaUploadedPhoto() *TLInputMediaUploadedPhoto {
	return &TLInputMediaUploadedPhoto{
		Data2: m,
	}
}

//  inputMediaPhoto#81fa373a flags:# id:InputPhoto caption:string ttl_seconds:flags.0?int = InputMedia;
//  inputMediaPhoto#e9bfb4f3 id:InputPhoto caption:string = InputMedia;
//  inputMediaPhoto#b3ba0635 flags:# id:InputPhoto ttl_seconds:flags.0?int = InputMedia;
//
func (m *InputMedia) To_InputMediaPhoto() *TLInputMediaPhoto {
	return &TLInputMediaPhoto{
		Data2: m,
	}
}

//  inputMediaGeoPoint#f9c44144 geo_point:InputGeoPoint = InputMedia;
//
func (m *InputMedia) To_InputMediaGeoPoint() *TLInputMediaGeoPoint {
	return &TLInputMediaGeoPoint{
		Data2: m,
	}
}

//  inputMediaContact#a6e45987 phone_number:string first_name:string last_name:string = InputMedia;
//  inputMediaContact#f8ab7dfb phone_number:string first_name:string last_name:string vcard:string = InputMedia;
//
func (m *InputMedia) To_InputMediaContact() *TLInputMediaContact {
	return &TLInputMediaContact{
		Data2: m,
	}
}

//  inputMediaUploadedDocument#e39621fd flags:# file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
//  inputMediaUploadedDocument#1d89306d file:InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string = InputMedia;
//  inputMediaUploadedDocument#5b38c6c1 flags:# nosound_video:flags.3?true force_file:flags.4?true file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
//
func (m *InputMedia) To_InputMediaUploadedDocument() *TLInputMediaUploadedDocument {
	return &TLInputMediaUploadedDocument{
		Data2: m,
	}
}

//  inputMediaDocument#5acb668e flags:# id:InputDocument caption:string ttl_seconds:flags.0?int = InputMedia;
//  inputMediaDocument#1a77f29c id:InputDocument caption:string = InputMedia;
//  inputMediaDocument#23ab23d2 flags:# id:InputDocument ttl_seconds:flags.0?int = InputMedia;
//  inputMediaDocument#33473058 flags:# id:InputDocument ttl_seconds:flags.0?int query:flags.1?string = InputMedia;
//
func (m *InputMedia) To_InputMediaDocument() *TLInputMediaDocument {
	return &TLInputMediaDocument{
		Data2: m,
	}
}

//  inputMediaVenue#2827a81a geo_point:InputGeoPoint title:string address:string provider:string venue_id:string = InputMedia;
//  inputMediaVenue#c13d1c11 geo_point:InputGeoPoint title:string address:string provider:string venue_id:string venue_type:string = InputMedia;
//
func (m *InputMedia) To_InputMediaVenue() *TLInputMediaVenue {
	return &TLInputMediaVenue{
		Data2: m,
	}
}

//  inputMediaGifExternal#4843b0fd url:string q:string = InputMedia;
//
func (m *InputMedia) To_InputMediaGifExternal() *TLInputMediaGifExternal {
	return &TLInputMediaGifExternal{
		Data2: m,
	}
}

//  inputMediaPhotoExternal#922aec1 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
//  inputMediaPhotoExternal#e5bbfe1a flags:# url:string ttl_seconds:flags.0?int = InputMedia;
//
func (m *InputMedia) To_InputMediaPhotoExternal() *TLInputMediaPhotoExternal {
	return &TLInputMediaPhotoExternal{
		Data2: m,
	}
}

//  inputMediaDocumentExternal#b6f74335 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
//  inputMediaDocumentExternal#fb52dc99 flags:# url:string ttl_seconds:flags.0?int = InputMedia;
//
func (m *InputMedia) To_InputMediaDocumentExternal() *TLInputMediaDocumentExternal {
	return &TLInputMediaDocumentExternal{
		Data2: m,
	}
}

//  inputMediaGame#d33f43f3 id:InputGame = InputMedia;
//
func (m *InputMedia) To_InputMediaGame() *TLInputMediaGame {
	return &TLInputMediaGame{
		Data2: m,
	}
}

//  inputMediaInvoice#92153685 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string start_param:string = InputMedia;
//  inputMediaInvoice#f4e096c3 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string provider_data:DataJSON start_param:string = InputMedia;
//
func (m *InputMedia) To_InputMediaInvoice() *TLInputMediaInvoice {
	return &TLInputMediaInvoice{
		Data2: m,
	}
}

//  inputMediaUploadedThumbDocument#ad613491 file:InputFile thumb:InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string = InputMedia;
//
func (m *InputMedia) To_InputMediaUploadedThumbDocument() *TLInputMediaUploadedThumbDocument {
	return &TLInputMediaUploadedThumbDocument{
		Data2: m,
	}
}

//  inputMediaGeoLive#7b1a118f geo_point:InputGeoPoint period:int = InputMedia;
//  inputMediaGeoLive#ce4e82fd flags:# stopped:flags.0?true geo_point:InputGeoPoint period:flags.1?int = InputMedia;
//  inputMediaGeoLive#971fa843 flags:# stopped:flags.0?true geo_point:InputGeoPoint heading:flags.2?int period:flags.1?int proximity_notification_radius:flags.3?int = InputMedia;
//
func (m *InputMedia) To_InputMediaGeoLive() *TLInputMediaGeoLive {
	return &TLInputMediaGeoLive{
		Data2: m,
	}
}

//  inputMediaPoll#6b3765b poll:Poll = InputMedia;
//  inputMediaPoll#abe9ca25 flags:# poll:Poll correct_answers:flags.0?Vector<bytes> = InputMedia;
//  inputMediaPoll#f94e5f1 flags:# poll:Poll correct_answers:flags.0?Vector<bytes> solution:flags.1?string solution_entities:flags.1?Vector<MessageEntity> = InputMedia;
//
func (m *InputMedia) To_InputMediaPoll() *TLInputMediaPoll {
	return &TLInputMediaPoll{
		Data2: m,
	}
}

//  inputMediaDice#aeffa807 = InputMedia;
//  inputMediaDice#e66fbf7b emoticon:string = InputMedia;
//
func (m *InputMedia) To_InputMediaDice() *TLInputMediaDice {
	return &TLInputMediaDice{
		Data2: m,
	}
}

//  inputMediaEmpty#9664f57f = InputMedia;
//
func (m *TLInputMediaEmpty) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaEmpty
	return m.Data2
}

//  inputMediaUploadedPhoto#2f37e231 flags:# file:InputFile caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
//  inputMediaUploadedPhoto#f7aff1c0 file:InputFile caption:string = InputMedia;
//  inputMediaUploadedPhoto#1e287d04 flags:# file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
//
func (m *TLInputMediaUploadedPhoto) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaUploadedPhoto
	return m.Data2
}

//  inputMediaPhoto#81fa373a flags:# id:InputPhoto caption:string ttl_seconds:flags.0?int = InputMedia;
//  inputMediaPhoto#e9bfb4f3 id:InputPhoto caption:string = InputMedia;
//  inputMediaPhoto#b3ba0635 flags:# id:InputPhoto ttl_seconds:flags.0?int = InputMedia;
//
func (m *TLInputMediaPhoto) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaPhoto
	return m.Data2
}

//  inputMediaGeoPoint#f9c44144 geo_point:InputGeoPoint = InputMedia;
//
func (m *TLInputMediaGeoPoint) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaGeoPoint
	return m.Data2
}

//  inputMediaContact#a6e45987 phone_number:string first_name:string last_name:string = InputMedia;
//  inputMediaContact#f8ab7dfb phone_number:string first_name:string last_name:string vcard:string = InputMedia;
//
func (m *TLInputMediaContact) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaContact
	return m.Data2
}

//  inputMediaUploadedDocument#e39621fd flags:# file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
//  inputMediaUploadedDocument#1d89306d file:InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string = InputMedia;
//  inputMediaUploadedDocument#5b38c6c1 flags:# nosound_video:flags.3?true force_file:flags.4?true file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
//
func (m *TLInputMediaUploadedDocument) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaUploadedDocument
	return m.Data2
}

//  inputMediaDocument#5acb668e flags:# id:InputDocument caption:string ttl_seconds:flags.0?int = InputMedia;
//  inputMediaDocument#1a77f29c id:InputDocument caption:string = InputMedia;
//  inputMediaDocument#23ab23d2 flags:# id:InputDocument ttl_seconds:flags.0?int = InputMedia;
//  inputMediaDocument#33473058 flags:# id:InputDocument ttl_seconds:flags.0?int query:flags.1?string = InputMedia;
//
func (m *TLInputMediaDocument) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaDocument
	return m.Data2
}

//  inputMediaVenue#2827a81a geo_point:InputGeoPoint title:string address:string provider:string venue_id:string = InputMedia;
//  inputMediaVenue#c13d1c11 geo_point:InputGeoPoint title:string address:string provider:string venue_id:string venue_type:string = InputMedia;
//
func (m *TLInputMediaVenue) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaVenue
	return m.Data2
}

//  inputMediaGifExternal#4843b0fd url:string q:string = InputMedia;
//
func (m *TLInputMediaGifExternal) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaGifExternal
	return m.Data2
}

//  inputMediaPhotoExternal#922aec1 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
//  inputMediaPhotoExternal#e5bbfe1a flags:# url:string ttl_seconds:flags.0?int = InputMedia;
//
func (m *TLInputMediaPhotoExternal) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaPhotoExternal
	return m.Data2
}

//  inputMediaDocumentExternal#b6f74335 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
//  inputMediaDocumentExternal#fb52dc99 flags:# url:string ttl_seconds:flags.0?int = InputMedia;
//
func (m *TLInputMediaDocumentExternal) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaDocumentExternal
	return m.Data2
}

//  inputMediaGame#d33f43f3 id:InputGame = InputMedia;
//
func (m *TLInputMediaGame) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaGame
	return m.Data2
}

//  inputMediaInvoice#92153685 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string start_param:string = InputMedia;
//  inputMediaInvoice#f4e096c3 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string provider_data:DataJSON start_param:string = InputMedia;
//
func (m *TLInputMediaInvoice) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaInvoice
	return m.Data2
}

//  inputMediaUploadedThumbDocument#ad613491 file:InputFile thumb:InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string = InputMedia;
//
func (m *TLInputMediaUploadedThumbDocument) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaUploadedThumbDocument
	return m.Data2
}

//  inputMediaGeoLive#7b1a118f geo_point:InputGeoPoint period:int = InputMedia;
//  inputMediaGeoLive#ce4e82fd flags:# stopped:flags.0?true geo_point:InputGeoPoint period:flags.1?int = InputMedia;
//  inputMediaGeoLive#971fa843 flags:# stopped:flags.0?true geo_point:InputGeoPoint heading:flags.2?int period:flags.1?int proximity_notification_radius:flags.3?int = InputMedia;
//
func (m *TLInputMediaGeoLive) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaGeoLive
	return m.Data2
}

//  inputMediaPoll#6b3765b poll:Poll = InputMedia;
//  inputMediaPoll#abe9ca25 flags:# poll:Poll correct_answers:flags.0?Vector<bytes> = InputMedia;
//  inputMediaPoll#f94e5f1 flags:# poll:Poll correct_answers:flags.0?Vector<bytes> solution:flags.1?string solution_entities:flags.1?Vector<MessageEntity> = InputMedia;
//
func (m *TLInputMediaPoll) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaPoll
	return m.Data2
}

//  inputMediaDice#aeffa807 = InputMedia;
//  inputMediaDice#e66fbf7b emoticon:string = InputMedia;
//
func (m *TLInputMediaDice) To_InputMedia() *InputMedia {
	m.Data2.Cmd = Cmd_InputMediaDice
	return m.Data2
}

//  + TL_InputMessageID
//  + TL_InputMessageReplyTo
//  + TL_InputMessagePinned
//  + TL_InputMessageCallbackQuery
//

//  inputMessageID#a676a322 id:int = InputMessage;
//
func (m *InputMessage) To_InputMessageID() *TLInputMessageID {
	return &TLInputMessageID{
		Data2: m,
	}
}

//  inputMessageReplyTo#bad88395 id:int = InputMessage;
//
func (m *InputMessage) To_InputMessageReplyTo() *TLInputMessageReplyTo {
	return &TLInputMessageReplyTo{
		Data2: m,
	}
}

//  inputMessagePinned#86872538 = InputMessage;
//
func (m *InputMessage) To_InputMessagePinned() *TLInputMessagePinned {
	return &TLInputMessagePinned{
		Data2: m,
	}
}

//  inputMessageCallbackQuery#acfa1a7e id:int query_id:long = InputMessage;
//
func (m *InputMessage) To_InputMessageCallbackQuery() *TLInputMessageCallbackQuery {
	return &TLInputMessageCallbackQuery{
		Data2: m,
	}
}

//  inputMessageID#a676a322 id:int = InputMessage;
//
func (m *TLInputMessageID) To_InputMessage() *InputMessage {
	m.Data2.Cmd = Cmd_InputMessageID
	return m.Data2
}

//  inputMessageReplyTo#bad88395 id:int = InputMessage;
//
func (m *TLInputMessageReplyTo) To_InputMessage() *InputMessage {
	m.Data2.Cmd = Cmd_InputMessageReplyTo
	return m.Data2
}

//  inputMessagePinned#86872538 = InputMessage;
//
func (m *TLInputMessagePinned) To_InputMessage() *InputMessage {
	m.Data2.Cmd = Cmd_InputMessagePinned
	return m.Data2
}

//  inputMessageCallbackQuery#acfa1a7e id:int query_id:long = InputMessage;
//
func (m *TLInputMessageCallbackQuery) To_InputMessage() *InputMessage {
	m.Data2.Cmd = Cmd_InputMessageCallbackQuery
	return m.Data2
}

//  + TL_InputNotifyPeer
//  + TL_InputNotifyUsers
//  + TL_InputNotifyChats
//  + TL_InputNotifyAll
//  + TL_InputNotifyBroadcasts
//

//  inputNotifyPeer#b8bc5b0c peer:InputPeer = InputNotifyPeer;
//
func (m *InputNotifyPeer) To_InputNotifyPeer() *TLInputNotifyPeer {
	return &TLInputNotifyPeer{
		Data2: m,
	}
}

//  inputNotifyUsers#193b4417 = InputNotifyPeer;
//
func (m *InputNotifyPeer) To_InputNotifyUsers() *TLInputNotifyUsers {
	return &TLInputNotifyUsers{
		Data2: m,
	}
}

//  inputNotifyChats#4a95e84e = InputNotifyPeer;
//
func (m *InputNotifyPeer) To_InputNotifyChats() *TLInputNotifyChats {
	return &TLInputNotifyChats{
		Data2: m,
	}
}

//  inputNotifyAll#a429b886 = InputNotifyPeer;
//
func (m *InputNotifyPeer) To_InputNotifyAll() *TLInputNotifyAll {
	return &TLInputNotifyAll{
		Data2: m,
	}
}

//  inputNotifyBroadcasts#b1db7c7e = InputNotifyPeer;
//
func (m *InputNotifyPeer) To_InputNotifyBroadcasts() *TLInputNotifyBroadcasts {
	return &TLInputNotifyBroadcasts{
		Data2: m,
	}
}

//  inputNotifyPeer#b8bc5b0c peer:InputPeer = InputNotifyPeer;
//
func (m *TLInputNotifyPeer) To_InputNotifyPeer() *InputNotifyPeer {
	m.Data2.Cmd = Cmd_InputNotifyPeer
	return m.Data2
}

//  inputNotifyUsers#193b4417 = InputNotifyPeer;
//
func (m *TLInputNotifyUsers) To_InputNotifyPeer() *InputNotifyPeer {
	m.Data2.Cmd = Cmd_InputNotifyUsers
	return m.Data2
}

//  inputNotifyChats#4a95e84e = InputNotifyPeer;
//
func (m *TLInputNotifyChats) To_InputNotifyPeer() *InputNotifyPeer {
	m.Data2.Cmd = Cmd_InputNotifyChats
	return m.Data2
}

//  inputNotifyAll#a429b886 = InputNotifyPeer;
//
func (m *TLInputNotifyAll) To_InputNotifyPeer() *InputNotifyPeer {
	m.Data2.Cmd = Cmd_InputNotifyAll
	return m.Data2
}

//  inputNotifyBroadcasts#b1db7c7e = InputNotifyPeer;
//
func (m *TLInputNotifyBroadcasts) To_InputNotifyPeer() *InputNotifyPeer {
	m.Data2.Cmd = Cmd_InputNotifyBroadcasts
	return m.Data2
}

//  + TL_InputPaymentCredentialsSaved
//  + TL_InputPaymentCredentials
//  + TL_InputPaymentCredentialsApplePay
//  + TL_InputPaymentCredentialsAndroidPay
//

//  inputPaymentCredentialsSaved#c10eb2cf id:string tmp_password:bytes = InputPaymentCredentials;
//
func (m *InputPaymentCredentials) To_InputPaymentCredentialsSaved() *TLInputPaymentCredentialsSaved {
	return &TLInputPaymentCredentialsSaved{
		Data2: m,
	}
}

//  inputPaymentCredentials#3417d728 flags:# save:flags.0?true data:DataJSON = InputPaymentCredentials;
//
func (m *InputPaymentCredentials) To_InputPaymentCredentials() *TLInputPaymentCredentials {
	return &TLInputPaymentCredentials{
		Data2: m,
	}
}

//  inputPaymentCredentialsApplePay#aa1c39f payment_data:DataJSON = InputPaymentCredentials;
//
func (m *InputPaymentCredentials) To_InputPaymentCredentialsApplePay() *TLInputPaymentCredentialsApplePay {
	return &TLInputPaymentCredentialsApplePay{
		Data2: m,
	}
}

//  inputPaymentCredentialsAndroidPay#ca05d50e payment_token:DataJSON google_transaction_id:string = InputPaymentCredentials;
//
func (m *InputPaymentCredentials) To_InputPaymentCredentialsAndroidPay() *TLInputPaymentCredentialsAndroidPay {
	return &TLInputPaymentCredentialsAndroidPay{
		Data2: m,
	}
}

//  inputPaymentCredentialsSaved#c10eb2cf id:string tmp_password:bytes = InputPaymentCredentials;
//
func (m *TLInputPaymentCredentialsSaved) To_InputPaymentCredentials() *InputPaymentCredentials {
	m.Data2.Cmd = Cmd_InputPaymentCredentialsSaved
	return m.Data2
}

//  inputPaymentCredentials#3417d728 flags:# save:flags.0?true data:DataJSON = InputPaymentCredentials;
//
func (m *TLInputPaymentCredentials) To_InputPaymentCredentials() *InputPaymentCredentials {
	m.Data2.Cmd = Cmd_InputPaymentCredentials
	return m.Data2
}

//  inputPaymentCredentialsApplePay#aa1c39f payment_data:DataJSON = InputPaymentCredentials;
//
func (m *TLInputPaymentCredentialsApplePay) To_InputPaymentCredentials() *InputPaymentCredentials {
	m.Data2.Cmd = Cmd_InputPaymentCredentialsApplePay
	return m.Data2
}

//  inputPaymentCredentialsAndroidPay#ca05d50e payment_token:DataJSON google_transaction_id:string = InputPaymentCredentials;
//
func (m *TLInputPaymentCredentialsAndroidPay) To_InputPaymentCredentials() *InputPaymentCredentials {
	m.Data2.Cmd = Cmd_InputPaymentCredentialsAndroidPay
	return m.Data2
}

//  + TL_InputPeerEmpty
//  + TL_InputPeerSelf
//  + TL_InputPeerChat
//  + TL_InputPeerUser
//  + TL_InputPeerChannel
//  + TL_InputPeerUserFromMessage
//  + TL_InputPeerChannelFromMessage
//

//  inputPeerEmpty#7f3b18ea = InputPeer;
//
func (m *InputPeer) To_InputPeerEmpty() *TLInputPeerEmpty {
	return &TLInputPeerEmpty{
		Data2: m,
	}
}

//  inputPeerSelf#7da07ec9 = InputPeer;
//
func (m *InputPeer) To_InputPeerSelf() *TLInputPeerSelf {
	return &TLInputPeerSelf{
		Data2: m,
	}
}

//  inputPeerChat#179be863 chat_id:int = InputPeer;
//
func (m *InputPeer) To_InputPeerChat() *TLInputPeerChat {
	return &TLInputPeerChat{
		Data2: m,
	}
}

//  inputPeerUser#7b8e7de6 user_id:int access_hash:long = InputPeer;
//
func (m *InputPeer) To_InputPeerUser() *TLInputPeerUser {
	return &TLInputPeerUser{
		Data2: m,
	}
}

//  inputPeerChannel#20adaef8 channel_id:int access_hash:long = InputPeer;
//
func (m *InputPeer) To_InputPeerChannel() *TLInputPeerChannel {
	return &TLInputPeerChannel{
		Data2: m,
	}
}

//  inputPeerUserFromMessage#17bae2e6 peer:InputPeer msg_id:int user_id:int = InputPeer;
//
func (m *InputPeer) To_InputPeerUserFromMessage() *TLInputPeerUserFromMessage {
	return &TLInputPeerUserFromMessage{
		Data2: m,
	}
}

//  inputPeerChannelFromMessage#9c95f7bb peer:InputPeer msg_id:int channel_id:int = InputPeer;
//
func (m *InputPeer) To_InputPeerChannelFromMessage() *TLInputPeerChannelFromMessage {
	return &TLInputPeerChannelFromMessage{
		Data2: m,
	}
}

//  inputPeerEmpty#7f3b18ea = InputPeer;
//
func (m *TLInputPeerEmpty) To_InputPeer() *InputPeer {
	m.Data2.Cmd = Cmd_InputPeerEmpty
	return m.Data2
}

//  inputPeerSelf#7da07ec9 = InputPeer;
//
func (m *TLInputPeerSelf) To_InputPeer() *InputPeer {
	m.Data2.Cmd = Cmd_InputPeerSelf
	return m.Data2
}

//  inputPeerChat#179be863 chat_id:int = InputPeer;
//
func (m *TLInputPeerChat) To_InputPeer() *InputPeer {
	m.Data2.Cmd = Cmd_InputPeerChat
	return m.Data2
}

//  inputPeerUser#7b8e7de6 user_id:int access_hash:long = InputPeer;
//
func (m *TLInputPeerUser) To_InputPeer() *InputPeer {
	m.Data2.Cmd = Cmd_InputPeerUser
	return m.Data2
}

//  inputPeerChannel#20adaef8 channel_id:int access_hash:long = InputPeer;
//
func (m *TLInputPeerChannel) To_InputPeer() *InputPeer {
	m.Data2.Cmd = Cmd_InputPeerChannel
	return m.Data2
}

//  inputPeerUserFromMessage#17bae2e6 peer:InputPeer msg_id:int user_id:int = InputPeer;
//
func (m *TLInputPeerUserFromMessage) To_InputPeer() *InputPeer {
	m.Data2.Cmd = Cmd_InputPeerUserFromMessage
	return m.Data2
}

//  inputPeerChannelFromMessage#9c95f7bb peer:InputPeer msg_id:int channel_id:int = InputPeer;
//
func (m *TLInputPeerChannelFromMessage) To_InputPeer() *InputPeer {
	m.Data2.Cmd = Cmd_InputPeerChannelFromMessage
	return m.Data2
}

//  + TL_InputPeerNotifyEventsEmpty
//  + TL_InputPeerNotifyEventsAll
//

//  inputPeerNotifyEventsEmpty#f03064d8 = InputPeerNotifyEvents;
//
func (m *InputPeerNotifyEvents) To_InputPeerNotifyEventsEmpty() *TLInputPeerNotifyEventsEmpty {
	return &TLInputPeerNotifyEventsEmpty{
		Data2: m,
	}
}

//  inputPeerNotifyEventsAll#e86a2c74 = InputPeerNotifyEvents;
//
func (m *InputPeerNotifyEvents) To_InputPeerNotifyEventsAll() *TLInputPeerNotifyEventsAll {
	return &TLInputPeerNotifyEventsAll{
		Data2: m,
	}
}

//  inputPeerNotifyEventsEmpty#f03064d8 = InputPeerNotifyEvents;
//
func (m *TLInputPeerNotifyEventsEmpty) To_InputPeerNotifyEvents() *InputPeerNotifyEvents {
	m.Data2.Cmd = Cmd_InputPeerNotifyEventsEmpty
	return m.Data2
}

//  inputPeerNotifyEventsAll#e86a2c74 = InputPeerNotifyEvents;
//
func (m *TLInputPeerNotifyEventsAll) To_InputPeerNotifyEvents() *InputPeerNotifyEvents {
	m.Data2.Cmd = Cmd_InputPeerNotifyEventsAll
	return m.Data2
}

//  + TL_InputPeerNotifySettings
//

//  inputPeerNotifySettings#38935eb2 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = InputPeerNotifySettings;
//  inputPeerNotifySettings#9c3d198e flags:# show_previews:flags.0?Bool silent:flags.1?Bool mute_until:flags.2?int sound:flags.3?string = InputPeerNotifySettings;
//
func (m *InputPeerNotifySettings) To_InputPeerNotifySettings() *TLInputPeerNotifySettings {
	return &TLInputPeerNotifySettings{
		Data2: m,
	}
}

//  inputPeerNotifySettings#38935eb2 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = InputPeerNotifySettings;
//  inputPeerNotifySettings#9c3d198e flags:# show_previews:flags.0?Bool silent:flags.1?Bool mute_until:flags.2?int sound:flags.3?string = InputPeerNotifySettings;
//
func (m *TLInputPeerNotifySettings) To_InputPeerNotifySettings() *InputPeerNotifySettings {
	m.Data2.Cmd = Cmd_InputPeerNotifySettings
	return m.Data2
}

//  + TL_InputPhoneCall
//

//  inputPhoneCall#1e36fded id:long access_hash:long = InputPhoneCall;
//
func (m *InputPhoneCall) To_InputPhoneCall() *TLInputPhoneCall {
	return &TLInputPhoneCall{
		Data2: m,
	}
}

//  inputPhoneCall#1e36fded id:long access_hash:long = InputPhoneCall;
//
func (m *TLInputPhoneCall) To_InputPhoneCall() *InputPhoneCall {
	m.Data2.Cmd = Cmd_InputPhoneCall
	return m.Data2
}

//  + TL_InputPhotoEmpty
//  + TL_InputPhoto
//

//  inputPhotoEmpty#1cd7bf0d = InputPhoto;
//
func (m *InputPhoto) To_InputPhotoEmpty() *TLInputPhotoEmpty {
	return &TLInputPhotoEmpty{
		Data2: m,
	}
}

//  inputPhoto#fb95c6c4 id:long access_hash:long = InputPhoto;
//  inputPhoto#3bb3b94a id:long access_hash:long file_reference:bytes = InputPhoto;
//
func (m *InputPhoto) To_InputPhoto() *TLInputPhoto {
	return &TLInputPhoto{
		Data2: m,
	}
}

//  inputPhotoEmpty#1cd7bf0d = InputPhoto;
//
func (m *TLInputPhotoEmpty) To_InputPhoto() *InputPhoto {
	m.Data2.Cmd = Cmd_InputPhotoEmpty
	return m.Data2
}

//  inputPhoto#fb95c6c4 id:long access_hash:long = InputPhoto;
//  inputPhoto#3bb3b94a id:long access_hash:long file_reference:bytes = InputPhoto;
//
func (m *TLInputPhoto) To_InputPhoto() *InputPhoto {
	m.Data2.Cmd = Cmd_InputPhoto
	return m.Data2
}

//  + TL_InputPhotoCrop
//  + TL_InputPhotoCropAuto
//

//  inputPhotoCrop#d9915325 crop_left:double crop_top:double crop_width:double = InputPhotoCrop;
//
func (m *InputPhotoCrop) To_InputPhotoCrop() *TLInputPhotoCrop {
	return &TLInputPhotoCrop{
		Data2: m,
	}
}

//  inputPhotoCropAuto#ade6b004 = InputPhotoCrop;
//
func (m *InputPhotoCrop) To_InputPhotoCropAuto() *TLInputPhotoCropAuto {
	return &TLInputPhotoCropAuto{
		Data2: m,
	}
}

//  inputPhotoCrop#d9915325 crop_left:double crop_top:double crop_width:double = InputPhotoCrop;
//
func (m *TLInputPhotoCrop) To_InputPhotoCrop() *InputPhotoCrop {
	m.Data2.Cmd = Cmd_InputPhotoCrop
	return m.Data2
}

//  inputPhotoCropAuto#ade6b004 = InputPhotoCrop;
//
func (m *TLInputPhotoCropAuto) To_InputPhotoCrop() *InputPhotoCrop {
	m.Data2.Cmd = Cmd_InputPhotoCropAuto
	return m.Data2
}

//  + TL_InputPrivacyKeyStatusTimestamp
//  + TL_InputPrivacyKeyChatInvite
//  + TL_InputPrivacyKeyPhoneCall
//  + TL_InputPrivacyKeyPhoneP2P
//  + TL_InputPrivacyKeyForwards
//  + TL_InputPrivacyKeyProfilePhoto
//  + TL_InputPrivacyKeyPhoneNumber
//  + TL_InputPrivacyKeyAddedByPhone
//

//  inputPrivacyKeyStatusTimestamp#4f96cb18 = InputPrivacyKey;
//
func (m *InputPrivacyKey) To_InputPrivacyKeyStatusTimestamp() *TLInputPrivacyKeyStatusTimestamp {
	return &TLInputPrivacyKeyStatusTimestamp{
		Data2: m,
	}
}

//  inputPrivacyKeyChatInvite#bdfb0426 = InputPrivacyKey;
//
func (m *InputPrivacyKey) To_InputPrivacyKeyChatInvite() *TLInputPrivacyKeyChatInvite {
	return &TLInputPrivacyKeyChatInvite{
		Data2: m,
	}
}

//  inputPrivacyKeyPhoneCall#fabadc5f = InputPrivacyKey;
//
func (m *InputPrivacyKey) To_InputPrivacyKeyPhoneCall() *TLInputPrivacyKeyPhoneCall {
	return &TLInputPrivacyKeyPhoneCall{
		Data2: m,
	}
}

//  inputPrivacyKeyPhoneP2P#db9e70d2 = InputPrivacyKey;
//
func (m *InputPrivacyKey) To_InputPrivacyKeyPhoneP2P() *TLInputPrivacyKeyPhoneP2P {
	return &TLInputPrivacyKeyPhoneP2P{
		Data2: m,
	}
}

//  inputPrivacyKeyForwards#a4dd4c08 = InputPrivacyKey;
//
func (m *InputPrivacyKey) To_InputPrivacyKeyForwards() *TLInputPrivacyKeyForwards {
	return &TLInputPrivacyKeyForwards{
		Data2: m,
	}
}

//  inputPrivacyKeyProfilePhoto#5719bacc = InputPrivacyKey;
//
func (m *InputPrivacyKey) To_InputPrivacyKeyProfilePhoto() *TLInputPrivacyKeyProfilePhoto {
	return &TLInputPrivacyKeyProfilePhoto{
		Data2: m,
	}
}

//  inputPrivacyKeyPhoneNumber#352dafa = InputPrivacyKey;
//
func (m *InputPrivacyKey) To_InputPrivacyKeyPhoneNumber() *TLInputPrivacyKeyPhoneNumber {
	return &TLInputPrivacyKeyPhoneNumber{
		Data2: m,
	}
}

//  inputPrivacyKeyAddedByPhone#d1219bdd = InputPrivacyKey;
//
func (m *InputPrivacyKey) To_InputPrivacyKeyAddedByPhone() *TLInputPrivacyKeyAddedByPhone {
	return &TLInputPrivacyKeyAddedByPhone{
		Data2: m,
	}
}

//  inputPrivacyKeyStatusTimestamp#4f96cb18 = InputPrivacyKey;
//
func (m *TLInputPrivacyKeyStatusTimestamp) To_InputPrivacyKey() *InputPrivacyKey {
	m.Data2.Cmd = Cmd_InputPrivacyKeyStatusTimestamp
	return m.Data2
}

//  inputPrivacyKeyChatInvite#bdfb0426 = InputPrivacyKey;
//
func (m *TLInputPrivacyKeyChatInvite) To_InputPrivacyKey() *InputPrivacyKey {
	m.Data2.Cmd = Cmd_InputPrivacyKeyChatInvite
	return m.Data2
}

//  inputPrivacyKeyPhoneCall#fabadc5f = InputPrivacyKey;
//
func (m *TLInputPrivacyKeyPhoneCall) To_InputPrivacyKey() *InputPrivacyKey {
	m.Data2.Cmd = Cmd_InputPrivacyKeyPhoneCall
	return m.Data2
}

//  inputPrivacyKeyPhoneP2P#db9e70d2 = InputPrivacyKey;
//
func (m *TLInputPrivacyKeyPhoneP2P) To_InputPrivacyKey() *InputPrivacyKey {
	m.Data2.Cmd = Cmd_InputPrivacyKeyPhoneP2P
	return m.Data2
}

//  inputPrivacyKeyForwards#a4dd4c08 = InputPrivacyKey;
//
func (m *TLInputPrivacyKeyForwards) To_InputPrivacyKey() *InputPrivacyKey {
	m.Data2.Cmd = Cmd_InputPrivacyKeyForwards
	return m.Data2
}

//  inputPrivacyKeyProfilePhoto#5719bacc = InputPrivacyKey;
//
func (m *TLInputPrivacyKeyProfilePhoto) To_InputPrivacyKey() *InputPrivacyKey {
	m.Data2.Cmd = Cmd_InputPrivacyKeyProfilePhoto
	return m.Data2
}

//  inputPrivacyKeyPhoneNumber#352dafa = InputPrivacyKey;
//
func (m *TLInputPrivacyKeyPhoneNumber) To_InputPrivacyKey() *InputPrivacyKey {
	m.Data2.Cmd = Cmd_InputPrivacyKeyPhoneNumber
	return m.Data2
}

//  inputPrivacyKeyAddedByPhone#d1219bdd = InputPrivacyKey;
//
func (m *TLInputPrivacyKeyAddedByPhone) To_InputPrivacyKey() *InputPrivacyKey {
	m.Data2.Cmd = Cmd_InputPrivacyKeyAddedByPhone
	return m.Data2
}

//  + TL_InputPrivacyValueAllowContacts
//  + TL_InputPrivacyValueAllowAll
//  + TL_InputPrivacyValueAllowUsers
//  + TL_InputPrivacyValueDisallowContacts
//  + TL_InputPrivacyValueDisallowAll
//  + TL_InputPrivacyValueDisallowUsers
//  + TL_InputPrivacyValueAllowChatParticipants
//  + TL_InputPrivacyValueDisallowChatParticipants
//

//  inputPrivacyValueAllowContacts#d09e07b = InputPrivacyRule;
//
func (m *InputPrivacyRule) To_InputPrivacyValueAllowContacts() *TLInputPrivacyValueAllowContacts {
	return &TLInputPrivacyValueAllowContacts{
		Data2: m,
	}
}

//  inputPrivacyValueAllowAll#184b35ce = InputPrivacyRule;
//
func (m *InputPrivacyRule) To_InputPrivacyValueAllowAll() *TLInputPrivacyValueAllowAll {
	return &TLInputPrivacyValueAllowAll{
		Data2: m,
	}
}

//  inputPrivacyValueAllowUsers#131cc67f users:Vector<InputUser> = InputPrivacyRule;
//
func (m *InputPrivacyRule) To_InputPrivacyValueAllowUsers() *TLInputPrivacyValueAllowUsers {
	return &TLInputPrivacyValueAllowUsers{
		Data2: m,
	}
}

//  inputPrivacyValueDisallowContacts#ba52007 = InputPrivacyRule;
//
func (m *InputPrivacyRule) To_InputPrivacyValueDisallowContacts() *TLInputPrivacyValueDisallowContacts {
	return &TLInputPrivacyValueDisallowContacts{
		Data2: m,
	}
}

//  inputPrivacyValueDisallowAll#d66b66c9 = InputPrivacyRule;
//
func (m *InputPrivacyRule) To_InputPrivacyValueDisallowAll() *TLInputPrivacyValueDisallowAll {
	return &TLInputPrivacyValueDisallowAll{
		Data2: m,
	}
}

//  inputPrivacyValueDisallowUsers#90110467 users:Vector<InputUser> = InputPrivacyRule;
//
func (m *InputPrivacyRule) To_InputPrivacyValueDisallowUsers() *TLInputPrivacyValueDisallowUsers {
	return &TLInputPrivacyValueDisallowUsers{
		Data2: m,
	}
}

//  inputPrivacyValueAllowChatParticipants#4c81c1ba chats:Vector<int> = InputPrivacyRule;
//
func (m *InputPrivacyRule) To_InputPrivacyValueAllowChatParticipants() *TLInputPrivacyValueAllowChatParticipants {
	return &TLInputPrivacyValueAllowChatParticipants{
		Data2: m,
	}
}

//  inputPrivacyValueDisallowChatParticipants#d82363af chats:Vector<int> = InputPrivacyRule;
//
func (m *InputPrivacyRule) To_InputPrivacyValueDisallowChatParticipants() *TLInputPrivacyValueDisallowChatParticipants {
	return &TLInputPrivacyValueDisallowChatParticipants{
		Data2: m,
	}
}

//  inputPrivacyValueAllowContacts#d09e07b = InputPrivacyRule;
//
func (m *TLInputPrivacyValueAllowContacts) To_InputPrivacyRule() *InputPrivacyRule {
	m.Data2.Cmd = Cmd_InputPrivacyValueAllowContacts
	return m.Data2
}

//  inputPrivacyValueAllowAll#184b35ce = InputPrivacyRule;
//
func (m *TLInputPrivacyValueAllowAll) To_InputPrivacyRule() *InputPrivacyRule {
	m.Data2.Cmd = Cmd_InputPrivacyValueAllowAll
	return m.Data2
}

//  inputPrivacyValueAllowUsers#131cc67f users:Vector<InputUser> = InputPrivacyRule;
//
func (m *TLInputPrivacyValueAllowUsers) To_InputPrivacyRule() *InputPrivacyRule {
	m.Data2.Cmd = Cmd_InputPrivacyValueAllowUsers
	return m.Data2
}

//  inputPrivacyValueDisallowContacts#ba52007 = InputPrivacyRule;
//
func (m *TLInputPrivacyValueDisallowContacts) To_InputPrivacyRule() *InputPrivacyRule {
	m.Data2.Cmd = Cmd_InputPrivacyValueDisallowContacts
	return m.Data2
}

//  inputPrivacyValueDisallowAll#d66b66c9 = InputPrivacyRule;
//
func (m *TLInputPrivacyValueDisallowAll) To_InputPrivacyRule() *InputPrivacyRule {
	m.Data2.Cmd = Cmd_InputPrivacyValueDisallowAll
	return m.Data2
}

//  inputPrivacyValueDisallowUsers#90110467 users:Vector<InputUser> = InputPrivacyRule;
//
func (m *TLInputPrivacyValueDisallowUsers) To_InputPrivacyRule() *InputPrivacyRule {
	m.Data2.Cmd = Cmd_InputPrivacyValueDisallowUsers
	return m.Data2
}

//  inputPrivacyValueAllowChatParticipants#4c81c1ba chats:Vector<int> = InputPrivacyRule;
//
func (m *TLInputPrivacyValueAllowChatParticipants) To_InputPrivacyRule() *InputPrivacyRule {
	m.Data2.Cmd = Cmd_InputPrivacyValueAllowChatParticipants
	return m.Data2
}

//  inputPrivacyValueDisallowChatParticipants#d82363af chats:Vector<int> = InputPrivacyRule;
//
func (m *TLInputPrivacyValueDisallowChatParticipants) To_InputPrivacyRule() *InputPrivacyRule {
	m.Data2.Cmd = Cmd_InputPrivacyValueDisallowChatParticipants
	return m.Data2
}

//  + TL_InputSecureFileUploaded
//  + TL_InputSecureFile
//

//  inputSecureFileUploaded#3334b0f0 id:long parts:int md5_checksum:string file_hash:bytes secret:bytes = InputSecureFile;
//
func (m *InputSecureFile) To_InputSecureFileUploaded() *TLInputSecureFileUploaded {
	return &TLInputSecureFileUploaded{
		Data2: m,
	}
}

//  inputSecureFile#5367e5be id:long access_hash:long = InputSecureFile;
//
func (m *InputSecureFile) To_InputSecureFile() *TLInputSecureFile {
	return &TLInputSecureFile{
		Data2: m,
	}
}

//  inputSecureFileUploaded#3334b0f0 id:long parts:int md5_checksum:string file_hash:bytes secret:bytes = InputSecureFile;
//
func (m *TLInputSecureFileUploaded) To_InputSecureFile() *InputSecureFile {
	m.Data2.Cmd = Cmd_InputSecureFileUploaded
	return m.Data2
}

//  inputSecureFile#5367e5be id:long access_hash:long = InputSecureFile;
//
func (m *TLInputSecureFile) To_InputSecureFile() *InputSecureFile {
	m.Data2.Cmd = Cmd_InputSecureFile
	return m.Data2
}

//  + TL_InputSecureValue
//

//  inputSecureValue#db21d0a7 flags:# type:SecureValueType data:flags.0?SecureData front_side:flags.1?InputSecureFile reverse_side:flags.2?InputSecureFile selfie:flags.3?InputSecureFile translation:flags.6?Vector<InputSecureFile> files:flags.4?Vector<InputSecureFile> plain_data:flags.5?SecurePlainData = InputSecureValue;
//
func (m *InputSecureValue) To_InputSecureValue() *TLInputSecureValue {
	return &TLInputSecureValue{
		Data2: m,
	}
}

//  inputSecureValue#db21d0a7 flags:# type:SecureValueType data:flags.0?SecureData front_side:flags.1?InputSecureFile reverse_side:flags.2?InputSecureFile selfie:flags.3?InputSecureFile translation:flags.6?Vector<InputSecureFile> files:flags.4?Vector<InputSecureFile> plain_data:flags.5?SecurePlainData = InputSecureValue;
//
func (m *TLInputSecureValue) To_InputSecureValue() *InputSecureValue {
	m.Data2.Cmd = Cmd_InputSecureValue
	return m.Data2
}

//  + TL_InputSingleMedia
//

//  inputSingleMedia#1cc6e91f flags:# media:InputMedia random_id:long message:string entities:flags.0?Vector<MessageEntity> = InputSingleMedia;
//
func (m *InputSingleMedia) To_InputSingleMedia() *TLInputSingleMedia {
	return &TLInputSingleMedia{
		Data2: m,
	}
}

//  inputSingleMedia#1cc6e91f flags:# media:InputMedia random_id:long message:string entities:flags.0?Vector<MessageEntity> = InputSingleMedia;
//
func (m *TLInputSingleMedia) To_InputSingleMedia() *InputSingleMedia {
	m.Data2.Cmd = Cmd_InputSingleMedia
	return m.Data2
}

//  + TL_InputStickerSetEmpty
//  + TL_InputStickerSetID
//  + TL_InputStickerSetShortName
//  + TL_InputStickerSetAnimatedEmoji
//  + TL_InputStickerSetDice
//

//  inputStickerSetEmpty#ffb62b95 = InputStickerSet;
//
func (m *InputStickerSet) To_InputStickerSetEmpty() *TLInputStickerSetEmpty {
	return &TLInputStickerSetEmpty{
		Data2: m,
	}
}

//  inputStickerSetID#9de7a269 id:long access_hash:long = InputStickerSet;
//
func (m *InputStickerSet) To_InputStickerSetID() *TLInputStickerSetID {
	return &TLInputStickerSetID{
		Data2: m,
	}
}

//  inputStickerSetShortName#861cc8a0 short_name:string = InputStickerSet;
//
func (m *InputStickerSet) To_InputStickerSetShortName() *TLInputStickerSetShortName {
	return &TLInputStickerSetShortName{
		Data2: m,
	}
}

//  inputStickerSetAnimatedEmoji#28703c8 = InputStickerSet;
//
func (m *InputStickerSet) To_InputStickerSetAnimatedEmoji() *TLInputStickerSetAnimatedEmoji {
	return &TLInputStickerSetAnimatedEmoji{
		Data2: m,
	}
}

//  inputStickerSetDice#79e21a53 = InputStickerSet;
//  inputStickerSetDice#e67f520e emoticon:string = InputStickerSet;
//
func (m *InputStickerSet) To_InputStickerSetDice() *TLInputStickerSetDice {
	return &TLInputStickerSetDice{
		Data2: m,
	}
}

//  inputStickerSetEmpty#ffb62b95 = InputStickerSet;
//
func (m *TLInputStickerSetEmpty) To_InputStickerSet() *InputStickerSet {
	m.Data2.Cmd = Cmd_InputStickerSetEmpty
	return m.Data2
}

//  inputStickerSetID#9de7a269 id:long access_hash:long = InputStickerSet;
//
func (m *TLInputStickerSetID) To_InputStickerSet() *InputStickerSet {
	m.Data2.Cmd = Cmd_InputStickerSetID
	return m.Data2
}

//  inputStickerSetShortName#861cc8a0 short_name:string = InputStickerSet;
//
func (m *TLInputStickerSetShortName) To_InputStickerSet() *InputStickerSet {
	m.Data2.Cmd = Cmd_InputStickerSetShortName
	return m.Data2
}

//  inputStickerSetAnimatedEmoji#28703c8 = InputStickerSet;
//
func (m *TLInputStickerSetAnimatedEmoji) To_InputStickerSet() *InputStickerSet {
	m.Data2.Cmd = Cmd_InputStickerSetAnimatedEmoji
	return m.Data2
}

//  inputStickerSetDice#79e21a53 = InputStickerSet;
//  inputStickerSetDice#e67f520e emoticon:string = InputStickerSet;
//
func (m *TLInputStickerSetDice) To_InputStickerSet() *InputStickerSet {
	m.Data2.Cmd = Cmd_InputStickerSetDice
	return m.Data2
}

//  + TL_InputStickerSetItem
//

//  inputStickerSetItem#ffa0a496 flags:# document:InputDocument emoji:string mask_coords:flags.0?MaskCoords = InputStickerSetItem;
//
func (m *InputStickerSetItem) To_InputStickerSetItem() *TLInputStickerSetItem {
	return &TLInputStickerSetItem{
		Data2: m,
	}
}

//  inputStickerSetItem#ffa0a496 flags:# document:InputDocument emoji:string mask_coords:flags.0?MaskCoords = InputStickerSetItem;
//
func (m *TLInputStickerSetItem) To_InputStickerSetItem() *InputStickerSetItem {
	m.Data2.Cmd = Cmd_InputStickerSetItem
	return m.Data2
}

//  + TL_InputStickeredMediaPhoto
//  + TL_InputStickeredMediaDocument
//

//  inputStickeredMediaPhoto#4a992157 id:InputPhoto = InputStickeredMedia;
//
func (m *InputStickeredMedia) To_InputStickeredMediaPhoto() *TLInputStickeredMediaPhoto {
	return &TLInputStickeredMediaPhoto{
		Data2: m,
	}
}

//  inputStickeredMediaDocument#438865b id:InputDocument = InputStickeredMedia;
//
func (m *InputStickeredMedia) To_InputStickeredMediaDocument() *TLInputStickeredMediaDocument {
	return &TLInputStickeredMediaDocument{
		Data2: m,
	}
}

//  inputStickeredMediaPhoto#4a992157 id:InputPhoto = InputStickeredMedia;
//
func (m *TLInputStickeredMediaPhoto) To_InputStickeredMedia() *InputStickeredMedia {
	m.Data2.Cmd = Cmd_InputStickeredMediaPhoto
	return m.Data2
}

//  inputStickeredMediaDocument#438865b id:InputDocument = InputStickeredMedia;
//
func (m *TLInputStickeredMediaDocument) To_InputStickeredMedia() *InputStickeredMedia {
	m.Data2.Cmd = Cmd_InputStickeredMediaDocument
	return m.Data2
}

//  + TL_InputTheme
//  + TL_InputThemeSlug
//

//  inputTheme#3c5693e9 id:long access_hash:long = InputTheme;
//
func (m *InputTheme) To_InputTheme() *TLInputTheme {
	return &TLInputTheme{
		Data2: m,
	}
}

//  inputThemeSlug#f5890df1 slug:string = InputTheme;
//
func (m *InputTheme) To_InputThemeSlug() *TLInputThemeSlug {
	return &TLInputThemeSlug{
		Data2: m,
	}
}

//  inputTheme#3c5693e9 id:long access_hash:long = InputTheme;
//
func (m *TLInputTheme) To_InputTheme() *InputTheme {
	m.Data2.Cmd = Cmd_InputTheme
	return m.Data2
}

//  inputThemeSlug#f5890df1 slug:string = InputTheme;
//
func (m *TLInputThemeSlug) To_InputTheme() *InputTheme {
	m.Data2.Cmd = Cmd_InputThemeSlug
	return m.Data2
}

//  + TL_InputThemeSettings
//

//  inputThemeSettings#bd507cd1 flags:# base_theme:BaseTheme accent_color:int message_top_color:flags.0?int message_bottom_color:flags.0?int wallpaper:flags.1?InputWallPaper wallpaper_settings:flags.1?WallPaperSettings = InputThemeSettings;
//
func (m *InputThemeSettings) To_InputThemeSettings() *TLInputThemeSettings {
	return &TLInputThemeSettings{
		Data2: m,
	}
}

//  inputThemeSettings#bd507cd1 flags:# base_theme:BaseTheme accent_color:int message_top_color:flags.0?int message_bottom_color:flags.0?int wallpaper:flags.1?InputWallPaper wallpaper_settings:flags.1?WallPaperSettings = InputThemeSettings;
//
func (m *TLInputThemeSettings) To_InputThemeSettings() *InputThemeSettings {
	m.Data2.Cmd = Cmd_InputThemeSettings
	return m.Data2
}

//  + TL_InputUserEmpty
//  + TL_InputUserSelf
//  + TL_InputUser
//  + TL_InputUserFromMessage
//

//  inputUserEmpty#b98886cf = InputUser;
//
func (m *InputUser) To_InputUserEmpty() *TLInputUserEmpty {
	return &TLInputUserEmpty{
		Data2: m,
	}
}

//  inputUserSelf#f7c1b13f = InputUser;
//
func (m *InputUser) To_InputUserSelf() *TLInputUserSelf {
	return &TLInputUserSelf{
		Data2: m,
	}
}

//  inputUser#d8292816 user_id:int access_hash:long = InputUser;
//
func (m *InputUser) To_InputUser() *TLInputUser {
	return &TLInputUser{
		Data2: m,
	}
}

//  inputUserFromMessage#2d117597 peer:InputPeer msg_id:int user_id:int = InputUser;
//
func (m *InputUser) To_InputUserFromMessage() *TLInputUserFromMessage {
	return &TLInputUserFromMessage{
		Data2: m,
	}
}

//  inputUserEmpty#b98886cf = InputUser;
//
func (m *TLInputUserEmpty) To_InputUser() *InputUser {
	m.Data2.Cmd = Cmd_InputUserEmpty
	return m.Data2
}

//  inputUserSelf#f7c1b13f = InputUser;
//
func (m *TLInputUserSelf) To_InputUser() *InputUser {
	m.Data2.Cmd = Cmd_InputUserSelf
	return m.Data2
}

//  inputUser#d8292816 user_id:int access_hash:long = InputUser;
//
func (m *TLInputUser) To_InputUser() *InputUser {
	m.Data2.Cmd = Cmd_InputUser
	return m.Data2
}

//  inputUserFromMessage#2d117597 peer:InputPeer msg_id:int user_id:int = InputUser;
//
func (m *TLInputUserFromMessage) To_InputUser() *InputUser {
	m.Data2.Cmd = Cmd_InputUserFromMessage
	return m.Data2
}

//  + TL_InputWallPaper
//  + TL_InputWallPaperSlug
//  + TL_InputWallPaperNoFile
//

//  inputWallPaper#e630b979 id:long access_hash:long = InputWallPaper;
//
func (m *InputWallPaper) To_InputWallPaper() *TLInputWallPaper {
	return &TLInputWallPaper{
		Data2: m,
	}
}

//  inputWallPaperSlug#72091c80 slug:string = InputWallPaper;
//
func (m *InputWallPaper) To_InputWallPaperSlug() *TLInputWallPaperSlug {
	return &TLInputWallPaperSlug{
		Data2: m,
	}
}

//  inputWallPaperNoFile#8427bbac = InputWallPaper;
//
func (m *InputWallPaper) To_InputWallPaperNoFile() *TLInputWallPaperNoFile {
	return &TLInputWallPaperNoFile{
		Data2: m,
	}
}

//  inputWallPaper#e630b979 id:long access_hash:long = InputWallPaper;
//
func (m *TLInputWallPaper) To_InputWallPaper() *InputWallPaper {
	m.Data2.Cmd = Cmd_InputWallPaper
	return m.Data2
}

//  inputWallPaperSlug#72091c80 slug:string = InputWallPaper;
//
func (m *TLInputWallPaperSlug) To_InputWallPaper() *InputWallPaper {
	m.Data2.Cmd = Cmd_InputWallPaperSlug
	return m.Data2
}

//  inputWallPaperNoFile#8427bbac = InputWallPaper;
//
func (m *TLInputWallPaperNoFile) To_InputWallPaper() *InputWallPaper {
	m.Data2.Cmd = Cmd_InputWallPaperNoFile
	return m.Data2
}

//  + TL_InputWebDocument
//

//  inputWebDocument#9bed434d url:string size:int mime_type:string attributes:Vector<DocumentAttribute> = InputWebDocument;
//
func (m *InputWebDocument) To_InputWebDocument() *TLInputWebDocument {
	return &TLInputWebDocument{
		Data2: m,
	}
}

//  inputWebDocument#9bed434d url:string size:int mime_type:string attributes:Vector<DocumentAttribute> = InputWebDocument;
//
func (m *TLInputWebDocument) To_InputWebDocument() *InputWebDocument {
	m.Data2.Cmd = Cmd_InputWebDocument
	return m.Data2
}

//  + TL_InputWebFileLocation
//  + TL_InputWebFileGeoPointLocation
//

//  inputWebFileLocation#c239d686 url:string access_hash:long = InputWebFileLocation;
//
func (m *InputWebFileLocation) To_InputWebFileLocation() *TLInputWebFileLocation {
	return &TLInputWebFileLocation{
		Data2: m,
	}
}

//  inputWebFileGeoPointLocation#9f2221c9 geo_point:InputGeoPoint access_hash:long w:int h:int zoom:int scale:int = InputWebFileLocation;
//
func (m *InputWebFileLocation) To_InputWebFileGeoPointLocation() *TLInputWebFileGeoPointLocation {
	return &TLInputWebFileGeoPointLocation{
		Data2: m,
	}
}

//  inputWebFileLocation#c239d686 url:string access_hash:long = InputWebFileLocation;
//
func (m *TLInputWebFileLocation) To_InputWebFileLocation() *InputWebFileLocation {
	m.Data2.Cmd = Cmd_InputWebFileLocation
	return m.Data2
}

//  inputWebFileGeoPointLocation#9f2221c9 geo_point:InputGeoPoint access_hash:long w:int h:int zoom:int scale:int = InputWebFileLocation;
//
func (m *TLInputWebFileGeoPointLocation) To_InputWebFileLocation() *InputWebFileLocation {
	m.Data2.Cmd = Cmd_InputWebFileGeoPointLocation
	return m.Data2
}

//  + TL_Invoice
//

//  invoice#c30aa358 flags:# test:flags.0?true name_requested:flags.1?true phone_requested:flags.2?true email_requested:flags.3?true shipping_address_requested:flags.4?true flexible:flags.5?true currency:string prices:Vector<LabeledPrice> = Invoice;
//
func (m *Invoice) To_Invoice() *TLInvoice {
	return &TLInvoice{
		Data2: m,
	}
}

//  invoice#c30aa358 flags:# test:flags.0?true name_requested:flags.1?true phone_requested:flags.2?true email_requested:flags.3?true shipping_address_requested:flags.4?true flexible:flags.5?true currency:string prices:Vector<LabeledPrice> = Invoice;
//
func (m *TLInvoice) To_Invoice() *Invoice {
	m.Data2.Cmd = Cmd_Invoice
	return m.Data2
}

//  + TL_IpPort
//  + TL_IpPortSecret
//

//  ipPort#d433ad73 ipv4:int port:int = IpPort;
//
func (m *IpPort) To_IpPort() *TLIpPort {
	return &TLIpPort{
		Data2: m,
	}
}

//  ipPortSecret#37982646 ipv4:int port:int secret:bytes = IpPort;
//
func (m *IpPort) To_IpPortSecret() *TLIpPortSecret {
	return &TLIpPortSecret{
		Data2: m,
	}
}

//  ipPort#d433ad73 ipv4:int port:int = IpPort;
//
func (m *TLIpPort) To_IpPort() *IpPort {
	m.Data2.Cmd = Cmd_IpPort
	return m.Data2
}

//  ipPortSecret#37982646 ipv4:int port:int secret:bytes = IpPort;
//
func (m *TLIpPortSecret) To_IpPort() *IpPort {
	m.Data2.Cmd = Cmd_IpPortSecret
	return m.Data2
}

//  + TL_JsonObjectValue
//

//  jsonObjectValue#c0de1bd9 key:string value:JSONValue = JSONObjectValue;
//
func (m *JSONObjectValue) To_JsonObjectValue() *TLJsonObjectValue {
	return &TLJsonObjectValue{
		Data2: m,
	}
}

//  jsonObjectValue#c0de1bd9 key:string value:JSONValue = JSONObjectValue;
//
func (m *TLJsonObjectValue) To_JSONObjectValue() *JSONObjectValue {
	m.Data2.Cmd = Cmd_JsonObjectValue
	return m.Data2
}

//  + TL_JsonNull
//  + TL_JsonBool
//  + TL_JsonNumber
//  + TL_JsonString
//  + TL_JsonArray
//  + TL_JsonObject
//

//  jsonNull#3f6d7b68 = JSONValue;
//
func (m *JSONValue) To_JsonNull() *TLJsonNull {
	return &TLJsonNull{
		Data2: m,
	}
}

//  jsonBool#c7345e6a value:Bool = JSONValue;
//
func (m *JSONValue) To_JsonBool() *TLJsonBool {
	return &TLJsonBool{
		Data2: m,
	}
}

//  jsonNumber#2be0dfa4 value:double = JSONValue;
//
func (m *JSONValue) To_JsonNumber() *TLJsonNumber {
	return &TLJsonNumber{
		Data2: m,
	}
}

//  jsonString#b71e767a value:string = JSONValue;
//
func (m *JSONValue) To_JsonString() *TLJsonString {
	return &TLJsonString{
		Data2: m,
	}
}

//  jsonArray#f7444763 value:Vector<JSONValue> = JSONValue;
//
func (m *JSONValue) To_JsonArray() *TLJsonArray {
	return &TLJsonArray{
		Data2: m,
	}
}

//  jsonObject#99c1d49d value:Vector<JSONObjectValue> = JSONValue;
//
func (m *JSONValue) To_JsonObject() *TLJsonObject {
	return &TLJsonObject{
		Data2: m,
	}
}

//  jsonNull#3f6d7b68 = JSONValue;
//
func (m *TLJsonNull) To_JSONValue() *JSONValue {
	m.Data2.Cmd = Cmd_JsonNull
	return m.Data2
}

//  jsonBool#c7345e6a value:Bool = JSONValue;
//
func (m *TLJsonBool) To_JSONValue() *JSONValue {
	m.Data2.Cmd = Cmd_JsonBool
	return m.Data2
}

//  jsonNumber#2be0dfa4 value:double = JSONValue;
//
func (m *TLJsonNumber) To_JSONValue() *JSONValue {
	m.Data2.Cmd = Cmd_JsonNumber
	return m.Data2
}

//  jsonString#b71e767a value:string = JSONValue;
//
func (m *TLJsonString) To_JSONValue() *JSONValue {
	m.Data2.Cmd = Cmd_JsonString
	return m.Data2
}

//  jsonArray#f7444763 value:Vector<JSONValue> = JSONValue;
//
func (m *TLJsonArray) To_JSONValue() *JSONValue {
	m.Data2.Cmd = Cmd_JsonArray
	return m.Data2
}

//  jsonObject#99c1d49d value:Vector<JSONObjectValue> = JSONValue;
//
func (m *TLJsonObject) To_JSONValue() *JSONValue {
	m.Data2.Cmd = Cmd_JsonObject
	return m.Data2
}

//  + TL_KeyboardButton
//  + TL_KeyboardButtonUrl
//  + TL_KeyboardButtonCallback
//  + TL_KeyboardButtonRequestPhone
//  + TL_KeyboardButtonRequestGeoLocation
//  + TL_KeyboardButtonSwitchInline
//  + TL_KeyboardButtonGame
//  + TL_KeyboardButtonBuy
//  + TL_KeyboardButtonUrlAuth
//  + TL_InputKeyboardButtonUrlAuth
//  + TL_KeyboardButtonRequestPoll
//

//  keyboardButton#a2fa4880 text:string = KeyboardButton;
//
func (m *KeyboardButton) To_KeyboardButton() *TLKeyboardButton {
	return &TLKeyboardButton{
		Data2: m,
	}
}

//  keyboardButtonUrl#258aff05 text:string url:string = KeyboardButton;
//
func (m *KeyboardButton) To_KeyboardButtonUrl() *TLKeyboardButtonUrl {
	return &TLKeyboardButtonUrl{
		Data2: m,
	}
}

//  keyboardButtonCallback#683a5e46 text:string data:bytes = KeyboardButton;
//  keyboardButtonCallback#35bbdb6b flags:# requires_password:flags.0?true text:string data:bytes = KeyboardButton;
//
func (m *KeyboardButton) To_KeyboardButtonCallback() *TLKeyboardButtonCallback {
	return &TLKeyboardButtonCallback{
		Data2: m,
	}
}

//  keyboardButtonRequestPhone#b16a6c29 text:string = KeyboardButton;
//
func (m *KeyboardButton) To_KeyboardButtonRequestPhone() *TLKeyboardButtonRequestPhone {
	return &TLKeyboardButtonRequestPhone{
		Data2: m,
	}
}

//  keyboardButtonRequestGeoLocation#fc796b3f text:string = KeyboardButton;
//
func (m *KeyboardButton) To_KeyboardButtonRequestGeoLocation() *TLKeyboardButtonRequestGeoLocation {
	return &TLKeyboardButtonRequestGeoLocation{
		Data2: m,
	}
}

//  keyboardButtonSwitchInline#568a748 flags:# same_peer:flags.0?true text:string query:string = KeyboardButton;
//  keyboardButtonSwitchInline#ea1b7a14 text:string query:string = KeyboardButton;
//
func (m *KeyboardButton) To_KeyboardButtonSwitchInline() *TLKeyboardButtonSwitchInline {
	return &TLKeyboardButtonSwitchInline{
		Data2: m,
	}
}

//  keyboardButtonGame#50f41ccf text:string = KeyboardButton;
//
func (m *KeyboardButton) To_KeyboardButtonGame() *TLKeyboardButtonGame {
	return &TLKeyboardButtonGame{
		Data2: m,
	}
}

//  keyboardButtonBuy#afd93fbb text:string = KeyboardButton;
//
func (m *KeyboardButton) To_KeyboardButtonBuy() *TLKeyboardButtonBuy {
	return &TLKeyboardButtonBuy{
		Data2: m,
	}
}

//  keyboardButtonUrlAuth#10b78d29 flags:# text:string fwd_text:flags.0?string url:string button_id:int = KeyboardButton;
//
func (m *KeyboardButton) To_KeyboardButtonUrlAuth() *TLKeyboardButtonUrlAuth {
	return &TLKeyboardButtonUrlAuth{
		Data2: m,
	}
}

//  inputKeyboardButtonUrlAuth#d02e7fd4 flags:# request_write_access:flags.0?true text:string fwd_text:flags.1?string url:string bot:InputUser = KeyboardButton;
//
func (m *KeyboardButton) To_InputKeyboardButtonUrlAuth() *TLInputKeyboardButtonUrlAuth {
	return &TLInputKeyboardButtonUrlAuth{
		Data2: m,
	}
}

//  keyboardButtonRequestPoll#bbc7515d flags:# quiz:flags.0?Bool text:string = KeyboardButton;
//
func (m *KeyboardButton) To_KeyboardButtonRequestPoll() *TLKeyboardButtonRequestPoll {
	return &TLKeyboardButtonRequestPoll{
		Data2: m,
	}
}

//  keyboardButton#a2fa4880 text:string = KeyboardButton;
//
func (m *TLKeyboardButton) To_KeyboardButton() *KeyboardButton {
	m.Data2.Cmd = Cmd_KeyboardButton
	return m.Data2
}

//  keyboardButtonUrl#258aff05 text:string url:string = KeyboardButton;
//
func (m *TLKeyboardButtonUrl) To_KeyboardButton() *KeyboardButton {
	m.Data2.Cmd = Cmd_KeyboardButtonUrl
	return m.Data2
}

//  keyboardButtonCallback#683a5e46 text:string data:bytes = KeyboardButton;
//  keyboardButtonCallback#35bbdb6b flags:# requires_password:flags.0?true text:string data:bytes = KeyboardButton;
//
func (m *TLKeyboardButtonCallback) To_KeyboardButton() *KeyboardButton {
	m.Data2.Cmd = Cmd_KeyboardButtonCallback
	return m.Data2
}

//  keyboardButtonRequestPhone#b16a6c29 text:string = KeyboardButton;
//
func (m *TLKeyboardButtonRequestPhone) To_KeyboardButton() *KeyboardButton {
	m.Data2.Cmd = Cmd_KeyboardButtonRequestPhone
	return m.Data2
}

//  keyboardButtonRequestGeoLocation#fc796b3f text:string = KeyboardButton;
//
func (m *TLKeyboardButtonRequestGeoLocation) To_KeyboardButton() *KeyboardButton {
	m.Data2.Cmd = Cmd_KeyboardButtonRequestGeoLocation
	return m.Data2
}

//  keyboardButtonSwitchInline#568a748 flags:# same_peer:flags.0?true text:string query:string = KeyboardButton;
//  keyboardButtonSwitchInline#ea1b7a14 text:string query:string = KeyboardButton;
//
func (m *TLKeyboardButtonSwitchInline) To_KeyboardButton() *KeyboardButton {
	m.Data2.Cmd = Cmd_KeyboardButtonSwitchInline
	return m.Data2
}

//  keyboardButtonGame#50f41ccf text:string = KeyboardButton;
//
func (m *TLKeyboardButtonGame) To_KeyboardButton() *KeyboardButton {
	m.Data2.Cmd = Cmd_KeyboardButtonGame
	return m.Data2
}

//  keyboardButtonBuy#afd93fbb text:string = KeyboardButton;
//
func (m *TLKeyboardButtonBuy) To_KeyboardButton() *KeyboardButton {
	m.Data2.Cmd = Cmd_KeyboardButtonBuy
	return m.Data2
}

//  keyboardButtonUrlAuth#10b78d29 flags:# text:string fwd_text:flags.0?string url:string button_id:int = KeyboardButton;
//
func (m *TLKeyboardButtonUrlAuth) To_KeyboardButton() *KeyboardButton {
	m.Data2.Cmd = Cmd_KeyboardButtonUrlAuth
	return m.Data2
}

//  inputKeyboardButtonUrlAuth#d02e7fd4 flags:# request_write_access:flags.0?true text:string fwd_text:flags.1?string url:string bot:InputUser = KeyboardButton;
//
func (m *TLInputKeyboardButtonUrlAuth) To_KeyboardButton() *KeyboardButton {
	m.Data2.Cmd = Cmd_InputKeyboardButtonUrlAuth
	return m.Data2
}

//  keyboardButtonRequestPoll#bbc7515d flags:# quiz:flags.0?Bool text:string = KeyboardButton;
//
func (m *TLKeyboardButtonRequestPoll) To_KeyboardButton() *KeyboardButton {
	m.Data2.Cmd = Cmd_KeyboardButtonRequestPoll
	return m.Data2
}

//  + TL_KeyboardButtonRow
//

//  keyboardButtonRow#77608b83 buttons:Vector<KeyboardButton> = KeyboardButtonRow;
//
func (m *KeyboardButtonRow) To_KeyboardButtonRow() *TLKeyboardButtonRow {
	return &TLKeyboardButtonRow{
		Data2: m,
	}
}

//  keyboardButtonRow#77608b83 buttons:Vector<KeyboardButton> = KeyboardButtonRow;
//
func (m *TLKeyboardButtonRow) To_KeyboardButtonRow() *KeyboardButtonRow {
	m.Data2.Cmd = Cmd_KeyboardButtonRow
	return m.Data2
}

//  + TL_LabeledPrice
//

//  labeledPrice#cb296bf8 label:string amount:long = LabeledPrice;
//
func (m *LabeledPrice) To_LabeledPrice() *TLLabeledPrice {
	return &TLLabeledPrice{
		Data2: m,
	}
}

//  labeledPrice#cb296bf8 label:string amount:long = LabeledPrice;
//
func (m *TLLabeledPrice) To_LabeledPrice() *LabeledPrice {
	m.Data2.Cmd = Cmd_LabeledPrice
	return m.Data2
}

//  + TL_LangPackDifference
//

//  langPackDifference#f385c1f6 lang_code:string from_version:int version:int strings:Vector<LangPackString> = LangPackDifference;
//
func (m *LangPackDifference) To_LangPackDifference() *TLLangPackDifference {
	return &TLLangPackDifference{
		Data2: m,
	}
}

//  langPackDifference#f385c1f6 lang_code:string from_version:int version:int strings:Vector<LangPackString> = LangPackDifference;
//
func (m *TLLangPackDifference) To_LangPackDifference() *LangPackDifference {
	m.Data2.Cmd = Cmd_LangPackDifference
	return m.Data2
}

//  + TL_LangPackLanguage
//

//  langPackLanguage#117698f1 name:string native_name:string lang_code:string = LangPackLanguage;
//  langPackLanguage#651b98d flags:# official:flags.0?true rtl:flags.2?true name:string native_name:string lang_code:string base_lang_code:flags.1?string plural_code:string strings_count:int translated_count:int = LangPackLanguage;
//  langPackLanguage#eeca5ce3 flags:# official:flags.0?true rtl:flags.2?true name:string native_name:string lang_code:string base_lang_code:flags.1?string plural_code:string strings_count:int translated_count:int translations_url:string = LangPackLanguage;
//
func (m *LangPackLanguage) To_LangPackLanguage() *TLLangPackLanguage {
	return &TLLangPackLanguage{
		Data2: m,
	}
}

//  langPackLanguage#117698f1 name:string native_name:string lang_code:string = LangPackLanguage;
//  langPackLanguage#651b98d flags:# official:flags.0?true rtl:flags.2?true name:string native_name:string lang_code:string base_lang_code:flags.1?string plural_code:string strings_count:int translated_count:int = LangPackLanguage;
//  langPackLanguage#eeca5ce3 flags:# official:flags.0?true rtl:flags.2?true name:string native_name:string lang_code:string base_lang_code:flags.1?string plural_code:string strings_count:int translated_count:int translations_url:string = LangPackLanguage;
//
func (m *TLLangPackLanguage) To_LangPackLanguage() *LangPackLanguage {
	m.Data2.Cmd = Cmd_LangPackLanguage
	return m.Data2
}

//  + TL_LangPackString
//  + TL_LangPackStringPluralized
//  + TL_LangPackStringDeleted
//

//  langPackString#cad181f6 key:string value:string = LangPackString;
//
func (m *LangPackString) To_LangPackString() *TLLangPackString {
	return &TLLangPackString{
		Data2: m,
	}
}

//  langPackStringPluralized#6c47ac9f flags:# key:string zero_value:flags.0?string one_value:flags.1?string two_value:flags.2?string few_value:flags.3?string many_value:flags.4?string other_value:string = LangPackString;
//
func (m *LangPackString) To_LangPackStringPluralized() *TLLangPackStringPluralized {
	return &TLLangPackStringPluralized{
		Data2: m,
	}
}

//  langPackStringDeleted#2979eeb2 key:string = LangPackString;
//
func (m *LangPackString) To_LangPackStringDeleted() *TLLangPackStringDeleted {
	return &TLLangPackStringDeleted{
		Data2: m,
	}
}

//  langPackString#cad181f6 key:string value:string = LangPackString;
//
func (m *TLLangPackString) To_LangPackString() *LangPackString {
	m.Data2.Cmd = Cmd_LangPackString
	return m.Data2
}

//  langPackStringPluralized#6c47ac9f flags:# key:string zero_value:flags.0?string one_value:flags.1?string two_value:flags.2?string few_value:flags.3?string many_value:flags.4?string other_value:string = LangPackString;
//
func (m *TLLangPackStringPluralized) To_LangPackString() *LangPackString {
	m.Data2.Cmd = Cmd_LangPackStringPluralized
	return m.Data2
}

//  langPackStringDeleted#2979eeb2 key:string = LangPackString;
//
func (m *TLLangPackStringDeleted) To_LangPackString() *LangPackString {
	m.Data2.Cmd = Cmd_LangPackStringDeleted
	return m.Data2
}

//  + TL_MaskCoords
//

//  maskCoords#aed6dbb2 n:int x:double y:double zoom:double = MaskCoords;
//
func (m *MaskCoords) To_MaskCoords() *TLMaskCoords {
	return &TLMaskCoords{
		Data2: m,
	}
}

//  maskCoords#aed6dbb2 n:int x:double y:double zoom:double = MaskCoords;
//
func (m *TLMaskCoords) To_MaskCoords() *MaskCoords {
	m.Data2.Cmd = Cmd_MaskCoords
	return m.Data2
}

//  + TL_MessageEmpty
//  + TL_Message
//  + TL_MessageService
//

//  messageEmpty#83e5de54 id:int = Message;
//
func (m *Message) To_MessageEmpty() *TLMessageEmpty {
	return &TLMessageEmpty{
		Data2: m,
	}
}

//  message#90dddc11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string = Message;
//  message#c09be45f flags:# unread:flags.0?true out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int = Message;
//  message#44f9b43d flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string grouped_id:flags.17?long = Message;
//  message#452c0e65 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true from_scheduled:flags.18?true legacy:flags.19?true edit_hide:flags.21?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string grouped_id:flags.17?long restriction_reason:flags.22?Vector<RestrictionReason> = Message;
//  message#58ae39c9 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true from_scheduled:flags.18?true legacy:flags.19?true edit_hide:flags.21?true pinned:flags.24?true id:int from_id:flags.8?Peer peer_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to:flags.3?MessageReplyHeader date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int forwards:flags.10?int replies:flags.23?MessageReplies edit_date:flags.15?int post_author:flags.16?string grouped_id:flags.17?long restriction_reason:flags.22?Vector<RestrictionReason> = Message;
//
func (m *Message) To_Message() *TLMessage {
	return &TLMessage{
		Data2: m,
	}
}

//  messageService#9e19a1f6 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer reply_to_msg_id:flags.3?int date:int action:MessageAction = Message;
//  messageService#286fa604 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true legacy:flags.19?true id:int from_id:flags.8?Peer peer_id:Peer reply_to:flags.3?MessageReplyHeader date:int action:MessageAction = Message;
//
func (m *Message) To_MessageService() *TLMessageService {
	return &TLMessageService{
		Data2: m,
	}
}

//  messageEmpty#83e5de54 id:int = Message;
//
func (m *TLMessageEmpty) To_Message() *Message {
	m.Data2.Cmd = Cmd_MessageEmpty
	return m.Data2
}

//  message#90dddc11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string = Message;
//  message#c09be45f flags:# unread:flags.0?true out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int = Message;
//  message#44f9b43d flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string grouped_id:flags.17?long = Message;
//  message#452c0e65 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true from_scheduled:flags.18?true legacy:flags.19?true edit_hide:flags.21?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string grouped_id:flags.17?long restriction_reason:flags.22?Vector<RestrictionReason> = Message;
//  message#58ae39c9 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true from_scheduled:flags.18?true legacy:flags.19?true edit_hide:flags.21?true pinned:flags.24?true id:int from_id:flags.8?Peer peer_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to:flags.3?MessageReplyHeader date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int forwards:flags.10?int replies:flags.23?MessageReplies edit_date:flags.15?int post_author:flags.16?string grouped_id:flags.17?long restriction_reason:flags.22?Vector<RestrictionReason> = Message;
//
func (m *TLMessage) To_Message() *Message {
	m.Data2.Cmd = Cmd_Message
	return m.Data2
}

//  messageService#9e19a1f6 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer reply_to_msg_id:flags.3?int date:int action:MessageAction = Message;
//  messageService#286fa604 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true legacy:flags.19?true id:int from_id:flags.8?Peer peer_id:Peer reply_to:flags.3?MessageReplyHeader date:int action:MessageAction = Message;
//
func (m *TLMessageService) To_Message() *Message {
	m.Data2.Cmd = Cmd_MessageService
	return m.Data2
}

//  + TL_MessageActionEmpty
//  + TL_MessageActionChatCreate
//  + TL_MessageActionChatEditTitle
//  + TL_MessageActionChatEditPhoto
//  + TL_MessageActionChatDeletePhoto
//  + TL_MessageActionChatAddUser
//  + TL_MessageActionChatDeleteUser
//  + TL_MessageActionChatJoinedByLink
//  + TL_MessageActionChannelCreate
//  + TL_MessageActionChatMigrateTo
//  + TL_MessageActionChannelMigrateFrom
//  + TL_MessageActionPinMessage
//  + TL_MessageActionHistoryClear
//  + TL_MessageActionGameScore
//  + TL_MessageActionPaymentSentMe
//  + TL_MessageActionPaymentSent
//  + TL_MessageActionScreenshotTaken
//  + TL_MessageActionCustomAction
//  + TL_MessageActionBotAllowed
//  + TL_MessageActionSecureValuesSentMe
//  + TL_MessageActionSecureValuesSent
//  + TL_MessageActionContactSignUp
//  + TL_MessageActionPhoneCall
//  + TL_MessageActionGeoProximityReached
//  + TL_MessageActionGroupCall
//  + TL_MessageActionInviteToGroupCall
//

//  messageActionEmpty#b6aef7b0 = MessageAction;
//
func (m *MessageAction) To_MessageActionEmpty() *TLMessageActionEmpty {
	return &TLMessageActionEmpty{
		Data2: m,
	}
}

//  messageActionChatCreate#a6638b9a title:string users:Vector<int> = MessageAction;
//
func (m *MessageAction) To_MessageActionChatCreate() *TLMessageActionChatCreate {
	return &TLMessageActionChatCreate{
		Data2: m,
	}
}

//  messageActionChatEditTitle#b5a1ce5a title:string = MessageAction;
//
func (m *MessageAction) To_MessageActionChatEditTitle() *TLMessageActionChatEditTitle {
	return &TLMessageActionChatEditTitle{
		Data2: m,
	}
}

//  messageActionChatEditPhoto#7fcb13a8 photo:Photo = MessageAction;
//
func (m *MessageAction) To_MessageActionChatEditPhoto() *TLMessageActionChatEditPhoto {
	return &TLMessageActionChatEditPhoto{
		Data2: m,
	}
}

//  messageActionChatDeletePhoto#95e3fbef = MessageAction;
//
func (m *MessageAction) To_MessageActionChatDeletePhoto() *TLMessageActionChatDeletePhoto {
	return &TLMessageActionChatDeletePhoto{
		Data2: m,
	}
}

//  messageActionChatAddUser#488a7337 users:Vector<int> = MessageAction;
//
func (m *MessageAction) To_MessageActionChatAddUser() *TLMessageActionChatAddUser {
	return &TLMessageActionChatAddUser{
		Data2: m,
	}
}

//  messageActionChatDeleteUser#b2ae9b0c user_id:int = MessageAction;
//
func (m *MessageAction) To_MessageActionChatDeleteUser() *TLMessageActionChatDeleteUser {
	return &TLMessageActionChatDeleteUser{
		Data2: m,
	}
}

//  messageActionChatJoinedByLink#f89cf5e8 inviter_id:int = MessageAction;
//
func (m *MessageAction) To_MessageActionChatJoinedByLink() *TLMessageActionChatJoinedByLink {
	return &TLMessageActionChatJoinedByLink{
		Data2: m,
	}
}

//  messageActionChannelCreate#95d2ac92 title:string = MessageAction;
//
func (m *MessageAction) To_MessageActionChannelCreate() *TLMessageActionChannelCreate {
	return &TLMessageActionChannelCreate{
		Data2: m,
	}
}

//  messageActionChatMigrateTo#51bdb021 channel_id:int = MessageAction;
//
func (m *MessageAction) To_MessageActionChatMigrateTo() *TLMessageActionChatMigrateTo {
	return &TLMessageActionChatMigrateTo{
		Data2: m,
	}
}

//  messageActionChannelMigrateFrom#b055eaee title:string chat_id:int = MessageAction;
//
func (m *MessageAction) To_MessageActionChannelMigrateFrom() *TLMessageActionChannelMigrateFrom {
	return &TLMessageActionChannelMigrateFrom{
		Data2: m,
	}
}

//  messageActionPinMessage#94bd38ed = MessageAction;
//
func (m *MessageAction) To_MessageActionPinMessage() *TLMessageActionPinMessage {
	return &TLMessageActionPinMessage{
		Data2: m,
	}
}

//  messageActionHistoryClear#9fbab604 = MessageAction;
//
func (m *MessageAction) To_MessageActionHistoryClear() *TLMessageActionHistoryClear {
	return &TLMessageActionHistoryClear{
		Data2: m,
	}
}

//  messageActionGameScore#92a72876 game_id:long score:int = MessageAction;
//
func (m *MessageAction) To_MessageActionGameScore() *TLMessageActionGameScore {
	return &TLMessageActionGameScore{
		Data2: m,
	}
}

//  messageActionPaymentSentMe#8f31b327 flags:# currency:string total_amount:long payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string charge:PaymentCharge = MessageAction;
//
func (m *MessageAction) To_MessageActionPaymentSentMe() *TLMessageActionPaymentSentMe {
	return &TLMessageActionPaymentSentMe{
		Data2: m,
	}
}

//  messageActionPaymentSent#40699cd0 currency:string total_amount:long = MessageAction;
//
func (m *MessageAction) To_MessageActionPaymentSent() *TLMessageActionPaymentSent {
	return &TLMessageActionPaymentSent{
		Data2: m,
	}
}

//  messageActionScreenshotTaken#4792929b = MessageAction;
//
func (m *MessageAction) To_MessageActionScreenshotTaken() *TLMessageActionScreenshotTaken {
	return &TLMessageActionScreenshotTaken{
		Data2: m,
	}
}

//  messageActionCustomAction#fae69f56 message:string = MessageAction;
//
func (m *MessageAction) To_MessageActionCustomAction() *TLMessageActionCustomAction {
	return &TLMessageActionCustomAction{
		Data2: m,
	}
}

//  messageActionBotAllowed#abe9affe domain:string = MessageAction;
//
func (m *MessageAction) To_MessageActionBotAllowed() *TLMessageActionBotAllowed {
	return &TLMessageActionBotAllowed{
		Data2: m,
	}
}

//  messageActionSecureValuesSentMe#1b287353 values:Vector<SecureValue> credentials:SecureCredentialsEncrypted = MessageAction;
//
func (m *MessageAction) To_MessageActionSecureValuesSentMe() *TLMessageActionSecureValuesSentMe {
	return &TLMessageActionSecureValuesSentMe{
		Data2: m,
	}
}

//  messageActionSecureValuesSent#d95c6154 types:Vector<SecureValueType> = MessageAction;
//
func (m *MessageAction) To_MessageActionSecureValuesSent() *TLMessageActionSecureValuesSent {
	return &TLMessageActionSecureValuesSent{
		Data2: m,
	}
}

//  messageActionContactSignUp#70ef8294 flags:# silent:flags.0?true = MessageAction;
//  messageActionContactSignUp#f3f25f76 = MessageAction;
//
func (m *MessageAction) To_MessageActionContactSignUp() *TLMessageActionContactSignUp {
	return &TLMessageActionContactSignUp{
		Data2: m,
	}
}

//  messageActionPhoneCall#80e11a7f flags:# video:flags.2?true call_id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = MessageAction;
//
func (m *MessageAction) To_MessageActionPhoneCall() *TLMessageActionPhoneCall {
	return &TLMessageActionPhoneCall{
		Data2: m,
	}
}

//  messageActionGeoProximityReached#98e0d697 from_id:Peer to_id:Peer distance:int = MessageAction;
//
func (m *MessageAction) To_MessageActionGeoProximityReached() *TLMessageActionGeoProximityReached {
	return &TLMessageActionGeoProximityReached{
		Data2: m,
	}
}

//  messageActionGroupCall#7a0d7f42 flags:# call:InputGroupCall duration:flags.0?int = MessageAction;
//
func (m *MessageAction) To_MessageActionGroupCall() *TLMessageActionGroupCall {
	return &TLMessageActionGroupCall{
		Data2: m,
	}
}

//  messageActionInviteToGroupCall#76b9f11a call:InputGroupCall users:Vector<int> = MessageAction;
//
func (m *MessageAction) To_MessageActionInviteToGroupCall() *TLMessageActionInviteToGroupCall {
	return &TLMessageActionInviteToGroupCall{
		Data2: m,
	}
}

//  messageActionEmpty#b6aef7b0 = MessageAction;
//
func (m *TLMessageActionEmpty) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionEmpty
	return m.Data2
}

//  messageActionChatCreate#a6638b9a title:string users:Vector<int> = MessageAction;
//
func (m *TLMessageActionChatCreate) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionChatCreate
	return m.Data2
}

//  messageActionChatEditTitle#b5a1ce5a title:string = MessageAction;
//
func (m *TLMessageActionChatEditTitle) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionChatEditTitle
	return m.Data2
}

//  messageActionChatEditPhoto#7fcb13a8 photo:Photo = MessageAction;
//
func (m *TLMessageActionChatEditPhoto) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionChatEditPhoto
	return m.Data2
}

//  messageActionChatDeletePhoto#95e3fbef = MessageAction;
//
func (m *TLMessageActionChatDeletePhoto) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionChatDeletePhoto
	return m.Data2
}

//  messageActionChatAddUser#488a7337 users:Vector<int> = MessageAction;
//
func (m *TLMessageActionChatAddUser) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionChatAddUser
	return m.Data2
}

//  messageActionChatDeleteUser#b2ae9b0c user_id:int = MessageAction;
//
func (m *TLMessageActionChatDeleteUser) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionChatDeleteUser
	return m.Data2
}

//  messageActionChatJoinedByLink#f89cf5e8 inviter_id:int = MessageAction;
//
func (m *TLMessageActionChatJoinedByLink) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionChatJoinedByLink
	return m.Data2
}

//  messageActionChannelCreate#95d2ac92 title:string = MessageAction;
//
func (m *TLMessageActionChannelCreate) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionChannelCreate
	return m.Data2
}

//  messageActionChatMigrateTo#51bdb021 channel_id:int = MessageAction;
//
func (m *TLMessageActionChatMigrateTo) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionChatMigrateTo
	return m.Data2
}

//  messageActionChannelMigrateFrom#b055eaee title:string chat_id:int = MessageAction;
//
func (m *TLMessageActionChannelMigrateFrom) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionChannelMigrateFrom
	return m.Data2
}

//  messageActionPinMessage#94bd38ed = MessageAction;
//
func (m *TLMessageActionPinMessage) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionPinMessage
	return m.Data2
}

//  messageActionHistoryClear#9fbab604 = MessageAction;
//
func (m *TLMessageActionHistoryClear) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionHistoryClear
	return m.Data2
}

//  messageActionGameScore#92a72876 game_id:long score:int = MessageAction;
//
func (m *TLMessageActionGameScore) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionGameScore
	return m.Data2
}

//  messageActionPaymentSentMe#8f31b327 flags:# currency:string total_amount:long payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string charge:PaymentCharge = MessageAction;
//
func (m *TLMessageActionPaymentSentMe) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionPaymentSentMe
	return m.Data2
}

//  messageActionPaymentSent#40699cd0 currency:string total_amount:long = MessageAction;
//
func (m *TLMessageActionPaymentSent) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionPaymentSent
	return m.Data2
}

//  messageActionScreenshotTaken#4792929b = MessageAction;
//
func (m *TLMessageActionScreenshotTaken) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionScreenshotTaken
	return m.Data2
}

//  messageActionCustomAction#fae69f56 message:string = MessageAction;
//
func (m *TLMessageActionCustomAction) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionCustomAction
	return m.Data2
}

//  messageActionBotAllowed#abe9affe domain:string = MessageAction;
//
func (m *TLMessageActionBotAllowed) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionBotAllowed
	return m.Data2
}

//  messageActionSecureValuesSentMe#1b287353 values:Vector<SecureValue> credentials:SecureCredentialsEncrypted = MessageAction;
//
func (m *TLMessageActionSecureValuesSentMe) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionSecureValuesSentMe
	return m.Data2
}

//  messageActionSecureValuesSent#d95c6154 types:Vector<SecureValueType> = MessageAction;
//
func (m *TLMessageActionSecureValuesSent) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionSecureValuesSent
	return m.Data2
}

//  messageActionContactSignUp#70ef8294 flags:# silent:flags.0?true = MessageAction;
//  messageActionContactSignUp#f3f25f76 = MessageAction;
//
func (m *TLMessageActionContactSignUp) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionContactSignUp
	return m.Data2
}

//  messageActionPhoneCall#80e11a7f flags:# video:flags.2?true call_id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = MessageAction;
//
func (m *TLMessageActionPhoneCall) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionPhoneCall
	return m.Data2
}

//  messageActionGeoProximityReached#98e0d697 from_id:Peer to_id:Peer distance:int = MessageAction;
//
func (m *TLMessageActionGeoProximityReached) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionGeoProximityReached
	return m.Data2
}

//  messageActionGroupCall#7a0d7f42 flags:# call:InputGroupCall duration:flags.0?int = MessageAction;
//
func (m *TLMessageActionGroupCall) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionGroupCall
	return m.Data2
}

//  messageActionInviteToGroupCall#76b9f11a call:InputGroupCall users:Vector<int> = MessageAction;
//
func (m *TLMessageActionInviteToGroupCall) To_MessageAction() *MessageAction {
	m.Data2.Cmd = Cmd_MessageActionInviteToGroupCall
	return m.Data2
}

//  + TL_MessageEntityUnknown
//  + TL_MessageEntityMention
//  + TL_MessageEntityHashtag
//  + TL_MessageEntityBotCommand
//  + TL_MessageEntityUrl
//  + TL_MessageEntityEmail
//  + TL_MessageEntityBold
//  + TL_MessageEntityItalic
//  + TL_MessageEntityCode
//  + TL_MessageEntityPre
//  + TL_MessageEntityTextUrl
//  + TL_MessageEntityMentionName
//  + TL_InputMessageEntityMentionName
//  + TL_MessageEntityPhone
//  + TL_MessageEntityCashtag
//  + TL_MessageEntityUnderline
//  + TL_MessageEntityStrike
//  + TL_MessageEntityBlockquote
//  + TL_MessageEntityBankCard
//

//  messageEntityUnknown#bb92ba95 offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityUnknown() *TLMessageEntityUnknown {
	return &TLMessageEntityUnknown{
		Data2: m,
	}
}

//  messageEntityMention#fa04579d offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityMention() *TLMessageEntityMention {
	return &TLMessageEntityMention{
		Data2: m,
	}
}

//  messageEntityHashtag#6f635b0d offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityHashtag() *TLMessageEntityHashtag {
	return &TLMessageEntityHashtag{
		Data2: m,
	}
}

//  messageEntityBotCommand#6cef8ac7 offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityBotCommand() *TLMessageEntityBotCommand {
	return &TLMessageEntityBotCommand{
		Data2: m,
	}
}

//  messageEntityUrl#6ed02538 offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityUrl() *TLMessageEntityUrl {
	return &TLMessageEntityUrl{
		Data2: m,
	}
}

//  messageEntityEmail#64e475c2 offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityEmail() *TLMessageEntityEmail {
	return &TLMessageEntityEmail{
		Data2: m,
	}
}

//  messageEntityBold#bd610bc9 offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityBold() *TLMessageEntityBold {
	return &TLMessageEntityBold{
		Data2: m,
	}
}

//  messageEntityItalic#826f8b60 offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityItalic() *TLMessageEntityItalic {
	return &TLMessageEntityItalic{
		Data2: m,
	}
}

//  messageEntityCode#28a20571 offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityCode() *TLMessageEntityCode {
	return &TLMessageEntityCode{
		Data2: m,
	}
}

//  messageEntityPre#73924be0 offset:int length:int language:string = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityPre() *TLMessageEntityPre {
	return &TLMessageEntityPre{
		Data2: m,
	}
}

//  messageEntityTextUrl#76a6d327 offset:int length:int url:string = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityTextUrl() *TLMessageEntityTextUrl {
	return &TLMessageEntityTextUrl{
		Data2: m,
	}
}

//  messageEntityMentionName#352dca58 offset:int length:int user_id:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityMentionName() *TLMessageEntityMentionName {
	return &TLMessageEntityMentionName{
		Data2: m,
	}
}

//  inputMessageEntityMentionName#208e68c9 offset:int length:int user_id:InputUser = MessageEntity;
//
func (m *MessageEntity) To_InputMessageEntityMentionName() *TLInputMessageEntityMentionName {
	return &TLInputMessageEntityMentionName{
		Data2: m,
	}
}

//  messageEntityPhone#9b69e34b offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityPhone() *TLMessageEntityPhone {
	return &TLMessageEntityPhone{
		Data2: m,
	}
}

//  messageEntityCashtag#4c4e743f offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityCashtag() *TLMessageEntityCashtag {
	return &TLMessageEntityCashtag{
		Data2: m,
	}
}

//  messageEntityUnderline#9c4e7e8b offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityUnderline() *TLMessageEntityUnderline {
	return &TLMessageEntityUnderline{
		Data2: m,
	}
}

//  messageEntityStrike#bf0693d4 offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityStrike() *TLMessageEntityStrike {
	return &TLMessageEntityStrike{
		Data2: m,
	}
}

//  messageEntityBlockquote#20df5d0 offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityBlockquote() *TLMessageEntityBlockquote {
	return &TLMessageEntityBlockquote{
		Data2: m,
	}
}

//  messageEntityBankCard#761e6af4 offset:int length:int = MessageEntity;
//
func (m *MessageEntity) To_MessageEntityBankCard() *TLMessageEntityBankCard {
	return &TLMessageEntityBankCard{
		Data2: m,
	}
}

//  messageEntityUnknown#bb92ba95 offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityUnknown) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityUnknown
	return m.Data2
}

//  messageEntityMention#fa04579d offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityMention) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityMention
	return m.Data2
}

//  messageEntityHashtag#6f635b0d offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityHashtag) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityHashtag
	return m.Data2
}

//  messageEntityBotCommand#6cef8ac7 offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityBotCommand) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityBotCommand
	return m.Data2
}

//  messageEntityUrl#6ed02538 offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityUrl) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityUrl
	return m.Data2
}

//  messageEntityEmail#64e475c2 offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityEmail) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityEmail
	return m.Data2
}

//  messageEntityBold#bd610bc9 offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityBold) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityBold
	return m.Data2
}

//  messageEntityItalic#826f8b60 offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityItalic) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityItalic
	return m.Data2
}

//  messageEntityCode#28a20571 offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityCode) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityCode
	return m.Data2
}

//  messageEntityPre#73924be0 offset:int length:int language:string = MessageEntity;
//
func (m *TLMessageEntityPre) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityPre
	return m.Data2
}

//  messageEntityTextUrl#76a6d327 offset:int length:int url:string = MessageEntity;
//
func (m *TLMessageEntityTextUrl) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityTextUrl
	return m.Data2
}

//  messageEntityMentionName#352dca58 offset:int length:int user_id:int = MessageEntity;
//
func (m *TLMessageEntityMentionName) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityMentionName
	return m.Data2
}

//  inputMessageEntityMentionName#208e68c9 offset:int length:int user_id:InputUser = MessageEntity;
//
func (m *TLInputMessageEntityMentionName) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_InputMessageEntityMentionName
	return m.Data2
}

//  messageEntityPhone#9b69e34b offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityPhone) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityPhone
	return m.Data2
}

//  messageEntityCashtag#4c4e743f offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityCashtag) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityCashtag
	return m.Data2
}

//  messageEntityUnderline#9c4e7e8b offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityUnderline) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityUnderline
	return m.Data2
}

//  messageEntityStrike#bf0693d4 offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityStrike) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityStrike
	return m.Data2
}

//  messageEntityBlockquote#20df5d0 offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityBlockquote) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityBlockquote
	return m.Data2
}

//  messageEntityBankCard#761e6af4 offset:int length:int = MessageEntity;
//
func (m *TLMessageEntityBankCard) To_MessageEntity() *MessageEntity {
	m.Data2.Cmd = Cmd_MessageEntityBankCard
	return m.Data2
}

//  + TL_MessageFwdHeader
//

//  messageFwdHeader#fadff4ac flags:# from_id:flags.0?int date:int channel_id:flags.1?int channel_post:flags.2?int post_author:flags.3?string = MessageFwdHeader;
//  messageFwdHeader#c786ddcb flags:# from_id:flags.0?int date:int channel_id:flags.1?int channel_post:flags.2?int = MessageFwdHeader;
//  messageFwdHeader#559ebe6d flags:# from_id:flags.0?int date:int channel_id:flags.1?int channel_post:flags.2?int post_author:flags.3?string saved_from_peer:flags.4?Peer saved_from_msg_id:flags.4?int = MessageFwdHeader;
//  messageFwdHeader#ec338270 flags:# from_id:flags.0?int from_name:flags.5?string date:int channel_id:flags.1?int channel_post:flags.2?int post_author:flags.3?string saved_from_peer:flags.4?Peer saved_from_msg_id:flags.4?int = MessageFwdHeader;
//  messageFwdHeader#353a686b flags:# from_id:flags.0?int from_name:flags.5?string date:int channel_id:flags.1?int channel_post:flags.2?int post_author:flags.3?string saved_from_peer:flags.4?Peer saved_from_msg_id:flags.4?int psa_type:flags.6?string = MessageFwdHeader;
//  messageFwdHeader#5f777dce flags:# from_id:flags.0?Peer from_name:flags.5?string date:int channel_post:flags.2?int post_author:flags.3?string saved_from_peer:flags.4?Peer saved_from_msg_id:flags.4?int psa_type:flags.6?string = MessageFwdHeader;
//
func (m *MessageFwdHeader) To_MessageFwdHeader() *TLMessageFwdHeader {
	return &TLMessageFwdHeader{
		Data2: m,
	}
}

//  messageFwdHeader#fadff4ac flags:# from_id:flags.0?int date:int channel_id:flags.1?int channel_post:flags.2?int post_author:flags.3?string = MessageFwdHeader;
//  messageFwdHeader#c786ddcb flags:# from_id:flags.0?int date:int channel_id:flags.1?int channel_post:flags.2?int = MessageFwdHeader;
//  messageFwdHeader#559ebe6d flags:# from_id:flags.0?int date:int channel_id:flags.1?int channel_post:flags.2?int post_author:flags.3?string saved_from_peer:flags.4?Peer saved_from_msg_id:flags.4?int = MessageFwdHeader;
//  messageFwdHeader#ec338270 flags:# from_id:flags.0?int from_name:flags.5?string date:int channel_id:flags.1?int channel_post:flags.2?int post_author:flags.3?string saved_from_peer:flags.4?Peer saved_from_msg_id:flags.4?int = MessageFwdHeader;
//  messageFwdHeader#353a686b flags:# from_id:flags.0?int from_name:flags.5?string date:int channel_id:flags.1?int channel_post:flags.2?int post_author:flags.3?string saved_from_peer:flags.4?Peer saved_from_msg_id:flags.4?int psa_type:flags.6?string = MessageFwdHeader;
//  messageFwdHeader#5f777dce flags:# from_id:flags.0?Peer from_name:flags.5?string date:int channel_post:flags.2?int post_author:flags.3?string saved_from_peer:flags.4?Peer saved_from_msg_id:flags.4?int psa_type:flags.6?string = MessageFwdHeader;
//
func (m *TLMessageFwdHeader) To_MessageFwdHeader() *MessageFwdHeader {
	m.Data2.Cmd = Cmd_MessageFwdHeader
	return m.Data2
}

//  + TL_MessageGroup
//

//  messageGroup#e8346f53 min_id:int max_id:int count:int date:int = MessageGroup;
//
func (m *MessageGroup) To_MessageGroup() *TLMessageGroup {
	return &TLMessageGroup{
		Data2: m,
	}
}

//  messageGroup#e8346f53 min_id:int max_id:int count:int date:int = MessageGroup;
//
func (m *TLMessageGroup) To_MessageGroup() *MessageGroup {
	m.Data2.Cmd = Cmd_MessageGroup
	return m.Data2
}

//  + TL_MessageInteractionCounters
//

//  messageInteractionCounters#ad4fc9bd msg_id:int views:int forwards:int = MessageInteractionCounters;
//
func (m *MessageInteractionCounters) To_MessageInteractionCounters() *TLMessageInteractionCounters {
	return &TLMessageInteractionCounters{
		Data2: m,
	}
}

//  messageInteractionCounters#ad4fc9bd msg_id:int views:int forwards:int = MessageInteractionCounters;
//
func (m *TLMessageInteractionCounters) To_MessageInteractionCounters() *MessageInteractionCounters {
	m.Data2.Cmd = Cmd_MessageInteractionCounters
	return m.Data2
}

//  + TL_MessageMediaEmpty
//  + TL_MessageMediaPhoto
//  + TL_MessageMediaGeo
//  + TL_MessageMediaContact
//  + TL_MessageMediaUnsupported
//  + TL_MessageMediaDocument
//  + TL_MessageMediaWebPage
//  + TL_MessageMediaVenue
//  + TL_MessageMediaGame
//  + TL_MessageMediaInvoice
//  + TL_MessageMediaGeoLive
//  + TL_MessageMediaPoll
//  + TL_MessageMediaDice
//

//  messageMediaEmpty#3ded6320 = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaEmpty() *TLMessageMediaEmpty {
	return &TLMessageMediaEmpty{
		Data2: m,
	}
}

//  messageMediaPhoto#b5223b0f flags:# photo:flags.0?Photo caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
//  messageMediaPhoto#3d8ce53d photo:Photo caption:string = MessageMedia;
//  messageMediaPhoto#695150d7 flags:# photo:flags.0?Photo ttl_seconds:flags.2?int = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaPhoto() *TLMessageMediaPhoto {
	return &TLMessageMediaPhoto{
		Data2: m,
	}
}

//  messageMediaGeo#56e0d474 geo:GeoPoint = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaGeo() *TLMessageMediaGeo {
	return &TLMessageMediaGeo{
		Data2: m,
	}
}

//  messageMediaContact#5e7d2f39 phone_number:string first_name:string last_name:string user_id:int = MessageMedia;
//  messageMediaContact#cbf24940 phone_number:string first_name:string last_name:string vcard:string user_id:int = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaContact() *TLMessageMediaContact {
	return &TLMessageMediaContact{
		Data2: m,
	}
}

//  messageMediaUnsupported#9f84f49e = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaUnsupported() *TLMessageMediaUnsupported {
	return &TLMessageMediaUnsupported{
		Data2: m,
	}
}

//  messageMediaDocument#7c4414d3 flags:# document:flags.0?Document caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
//  messageMediaDocument#f3e02ea8 document:Document caption:string = MessageMedia;
//  messageMediaDocument#9cb070d7 flags:# document:flags.0?Document ttl_seconds:flags.2?int = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaDocument() *TLMessageMediaDocument {
	return &TLMessageMediaDocument{
		Data2: m,
	}
}

//  messageMediaWebPage#a32dd600 webpage:WebPage = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaWebPage() *TLMessageMediaWebPage {
	return &TLMessageMediaWebPage{
		Data2: m,
	}
}

//  messageMediaVenue#7912b71f geo:GeoPoint title:string address:string provider:string venue_id:string = MessageMedia;
//  messageMediaVenue#2ec0533f geo:GeoPoint title:string address:string provider:string venue_id:string venue_type:string = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaVenue() *TLMessageMediaVenue {
	return &TLMessageMediaVenue{
		Data2: m,
	}
}

//  messageMediaGame#fdb19008 game:Game = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaGame() *TLMessageMediaGame {
	return &TLMessageMediaGame{
		Data2: m,
	}
}

//  messageMediaInvoice#84551347 flags:# shipping_address_requested:flags.1?true test:flags.3?true title:string description:string photo:flags.0?WebDocument receipt_msg_id:flags.2?int currency:string total_amount:long start_param:string = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaInvoice() *TLMessageMediaInvoice {
	return &TLMessageMediaInvoice{
		Data2: m,
	}
}

//  messageMediaGeoLive#7c3c2609 geo:GeoPoint period:int = MessageMedia;
//  messageMediaGeoLive#b940c666 flags:# geo:GeoPoint heading:flags.0?int period:int proximity_notification_radius:flags.1?int = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaGeoLive() *TLMessageMediaGeoLive {
	return &TLMessageMediaGeoLive{
		Data2: m,
	}
}

//  messageMediaPoll#4bd6e798 poll:Poll results:PollResults = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaPoll() *TLMessageMediaPoll {
	return &TLMessageMediaPoll{
		Data2: m,
	}
}

//  messageMediaDice#638fe46b value:int = MessageMedia;
//  messageMediaDice#3f7ee58b value:int emoticon:string = MessageMedia;
//
func (m *MessageMedia) To_MessageMediaDice() *TLMessageMediaDice {
	return &TLMessageMediaDice{
		Data2: m,
	}
}

//  messageMediaEmpty#3ded6320 = MessageMedia;
//
func (m *TLMessageMediaEmpty) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaEmpty
	return m.Data2
}

//  messageMediaPhoto#b5223b0f flags:# photo:flags.0?Photo caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
//  messageMediaPhoto#3d8ce53d photo:Photo caption:string = MessageMedia;
//  messageMediaPhoto#695150d7 flags:# photo:flags.0?Photo ttl_seconds:flags.2?int = MessageMedia;
//
func (m *TLMessageMediaPhoto) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaPhoto
	return m.Data2
}

//  messageMediaGeo#56e0d474 geo:GeoPoint = MessageMedia;
//
func (m *TLMessageMediaGeo) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaGeo
	return m.Data2
}

//  messageMediaContact#5e7d2f39 phone_number:string first_name:string last_name:string user_id:int = MessageMedia;
//  messageMediaContact#cbf24940 phone_number:string first_name:string last_name:string vcard:string user_id:int = MessageMedia;
//
func (m *TLMessageMediaContact) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaContact
	return m.Data2
}

//  messageMediaUnsupported#9f84f49e = MessageMedia;
//
func (m *TLMessageMediaUnsupported) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaUnsupported
	return m.Data2
}

//  messageMediaDocument#7c4414d3 flags:# document:flags.0?Document caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
//  messageMediaDocument#f3e02ea8 document:Document caption:string = MessageMedia;
//  messageMediaDocument#9cb070d7 flags:# document:flags.0?Document ttl_seconds:flags.2?int = MessageMedia;
//
func (m *TLMessageMediaDocument) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaDocument
	return m.Data2
}

//  messageMediaWebPage#a32dd600 webpage:WebPage = MessageMedia;
//
func (m *TLMessageMediaWebPage) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaWebPage
	return m.Data2
}

//  messageMediaVenue#7912b71f geo:GeoPoint title:string address:string provider:string venue_id:string = MessageMedia;
//  messageMediaVenue#2ec0533f geo:GeoPoint title:string address:string provider:string venue_id:string venue_type:string = MessageMedia;
//
func (m *TLMessageMediaVenue) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaVenue
	return m.Data2
}

//  messageMediaGame#fdb19008 game:Game = MessageMedia;
//
func (m *TLMessageMediaGame) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaGame
	return m.Data2
}

//  messageMediaInvoice#84551347 flags:# shipping_address_requested:flags.1?true test:flags.3?true title:string description:string photo:flags.0?WebDocument receipt_msg_id:flags.2?int currency:string total_amount:long start_param:string = MessageMedia;
//
func (m *TLMessageMediaInvoice) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaInvoice
	return m.Data2
}

//  messageMediaGeoLive#7c3c2609 geo:GeoPoint period:int = MessageMedia;
//  messageMediaGeoLive#b940c666 flags:# geo:GeoPoint heading:flags.0?int period:int proximity_notification_radius:flags.1?int = MessageMedia;
//
func (m *TLMessageMediaGeoLive) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaGeoLive
	return m.Data2
}

//  messageMediaPoll#4bd6e798 poll:Poll results:PollResults = MessageMedia;
//
func (m *TLMessageMediaPoll) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaPoll
	return m.Data2
}

//  messageMediaDice#638fe46b value:int = MessageMedia;
//  messageMediaDice#3f7ee58b value:int emoticon:string = MessageMedia;
//
func (m *TLMessageMediaDice) To_MessageMedia() *MessageMedia {
	m.Data2.Cmd = Cmd_MessageMediaDice
	return m.Data2
}

//  + TL_MessageRange
//

//  messageRange#ae30253 min_id:int max_id:int = MessageRange;
//
func (m *MessageRange) To_MessageRange() *TLMessageRange {
	return &TLMessageRange{
		Data2: m,
	}
}

//  messageRange#ae30253 min_id:int max_id:int = MessageRange;
//
func (m *TLMessageRange) To_MessageRange() *MessageRange {
	m.Data2.Cmd = Cmd_MessageRange
	return m.Data2
}

//  + TL_MessageReplies
//

//  messageReplies#4128faac flags:# comments:flags.0?true replies:int replies_pts:int recent_repliers:flags.1?Vector<Peer> channel_id:flags.0?int max_id:flags.2?int read_max_id:flags.3?int = MessageReplies;
//
func (m *MessageReplies) To_MessageReplies() *TLMessageReplies {
	return &TLMessageReplies{
		Data2: m,
	}
}

//  messageReplies#4128faac flags:# comments:flags.0?true replies:int replies_pts:int recent_repliers:flags.1?Vector<Peer> channel_id:flags.0?int max_id:flags.2?int read_max_id:flags.3?int = MessageReplies;
//
func (m *TLMessageReplies) To_MessageReplies() *MessageReplies {
	m.Data2.Cmd = Cmd_MessageReplies
	return m.Data2
}

//  + TL_MessageReplyHeader
//

//  messageReplyHeader#a6d57763 flags:# reply_to_msg_id:int reply_to_peer_id:flags.0?Peer reply_to_top_id:flags.1?int = MessageReplyHeader;
//
func (m *MessageReplyHeader) To_MessageReplyHeader() *TLMessageReplyHeader {
	return &TLMessageReplyHeader{
		Data2: m,
	}
}

//  messageReplyHeader#a6d57763 flags:# reply_to_msg_id:int reply_to_peer_id:flags.0?Peer reply_to_top_id:flags.1?int = MessageReplyHeader;
//
func (m *TLMessageReplyHeader) To_MessageReplyHeader() *MessageReplyHeader {
	m.Data2.Cmd = Cmd_MessageReplyHeader
	return m.Data2
}

//  + TL_MessageUserVote
//  + TL_MessageUserVoteInputOption
//  + TL_MessageUserVoteMultiple
//

//  messageUserVote#f212f56d user_id:int option:bytes = MessageUserVote;
//  messageUserVote#a28e5559 user_id:int option:bytes date:int = MessageUserVote;
//
func (m *MessageUserVote) To_MessageUserVote() *TLMessageUserVote {
	return &TLMessageUserVote{
		Data2: m,
	}
}

//  messageUserVoteInputOption#36377430 user_id:int date:int = MessageUserVote;
//
func (m *MessageUserVote) To_MessageUserVoteInputOption() *TLMessageUserVoteInputOption {
	return &TLMessageUserVoteInputOption{
		Data2: m,
	}
}

//  messageUserVoteMultiple#e8fe0de user_id:int options:Vector<bytes> date:int = MessageUserVote;
//
func (m *MessageUserVote) To_MessageUserVoteMultiple() *TLMessageUserVoteMultiple {
	return &TLMessageUserVoteMultiple{
		Data2: m,
	}
}

//  messageUserVote#f212f56d user_id:int option:bytes = MessageUserVote;
//  messageUserVote#a28e5559 user_id:int option:bytes date:int = MessageUserVote;
//
func (m *TLMessageUserVote) To_MessageUserVote() *MessageUserVote {
	m.Data2.Cmd = Cmd_MessageUserVote
	return m.Data2
}

//  messageUserVoteInputOption#36377430 user_id:int date:int = MessageUserVote;
//
func (m *TLMessageUserVoteInputOption) To_MessageUserVote() *MessageUserVote {
	m.Data2.Cmd = Cmd_MessageUserVoteInputOption
	return m.Data2
}

//  messageUserVoteMultiple#e8fe0de user_id:int options:Vector<bytes> date:int = MessageUserVote;
//
func (m *TLMessageUserVoteMultiple) To_MessageUserVote() *MessageUserVote {
	m.Data2.Cmd = Cmd_MessageUserVoteMultiple
	return m.Data2
}

//  + TL_MessageViews
//

//  messageViews#455b853d flags:# views:flags.0?int forwards:flags.1?int replies:flags.2?MessageReplies = MessageViews;
//
func (m *MessageViews) To_MessageViews() *TLMessageViews {
	return &TLMessageViews{
		Data2: m,
	}
}

//  messageViews#455b853d flags:# views:flags.0?int forwards:flags.1?int replies:flags.2?MessageReplies = MessageViews;
//
func (m *TLMessageViews) To_MessageViews() *MessageViews {
	m.Data2.Cmd = Cmd_MessageViews
	return m.Data2
}

//  + TL_InputMessagesFilterEmpty
//  + TL_InputMessagesFilterPhotos
//  + TL_InputMessagesFilterVideo
//  + TL_InputMessagesFilterPhotoVideo
//  + TL_InputMessagesFilterPhotoVideoDocuments
//  + TL_InputMessagesFilterDocument
//  + TL_InputMessagesFilterUrl
//  + TL_InputMessagesFilterGif
//  + TL_InputMessagesFilterVoice
//  + TL_InputMessagesFilterMusic
//  + TL_InputMessagesFilterChatPhotos
//  + TL_InputMessagesFilterPhoneCalls
//  + TL_InputMessagesFilterRoundVoice
//  + TL_InputMessagesFilterRoundVideo
//  + TL_InputMessagesFilterMyMentions
//  + TL_InputMessagesFilterGeo
//  + TL_InputMessagesFilterContacts
//  + TL_InputMessagesFilterPinned
//

//  inputMessagesFilterEmpty#57e2f66c = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterEmpty() *TLInputMessagesFilterEmpty {
	return &TLInputMessagesFilterEmpty{
		Data2: m,
	}
}

//  inputMessagesFilterPhotos#9609a51c = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterPhotos() *TLInputMessagesFilterPhotos {
	return &TLInputMessagesFilterPhotos{
		Data2: m,
	}
}

//  inputMessagesFilterVideo#9fc00e65 = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterVideo() *TLInputMessagesFilterVideo {
	return &TLInputMessagesFilterVideo{
		Data2: m,
	}
}

//  inputMessagesFilterPhotoVideo#56e9f0e4 = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterPhotoVideo() *TLInputMessagesFilterPhotoVideo {
	return &TLInputMessagesFilterPhotoVideo{
		Data2: m,
	}
}

//  inputMessagesFilterPhotoVideoDocuments#d95e73bb = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterPhotoVideoDocuments() *TLInputMessagesFilterPhotoVideoDocuments {
	return &TLInputMessagesFilterPhotoVideoDocuments{
		Data2: m,
	}
}

//  inputMessagesFilterDocument#9eddf188 = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterDocument() *TLInputMessagesFilterDocument {
	return &TLInputMessagesFilterDocument{
		Data2: m,
	}
}

//  inputMessagesFilterUrl#7ef0dd87 = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterUrl() *TLInputMessagesFilterUrl {
	return &TLInputMessagesFilterUrl{
		Data2: m,
	}
}

//  inputMessagesFilterGif#ffc86587 = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterGif() *TLInputMessagesFilterGif {
	return &TLInputMessagesFilterGif{
		Data2: m,
	}
}

//  inputMessagesFilterVoice#50f5c392 = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterVoice() *TLInputMessagesFilterVoice {
	return &TLInputMessagesFilterVoice{
		Data2: m,
	}
}

//  inputMessagesFilterMusic#3751b49e = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterMusic() *TLInputMessagesFilterMusic {
	return &TLInputMessagesFilterMusic{
		Data2: m,
	}
}

//  inputMessagesFilterChatPhotos#3a20ecb8 = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterChatPhotos() *TLInputMessagesFilterChatPhotos {
	return &TLInputMessagesFilterChatPhotos{
		Data2: m,
	}
}

//  inputMessagesFilterPhoneCalls#80c99768 flags:# missed:flags.0?true = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterPhoneCalls() *TLInputMessagesFilterPhoneCalls {
	return &TLInputMessagesFilterPhoneCalls{
		Data2: m,
	}
}

//  inputMessagesFilterRoundVoice#7a7c17a4 = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterRoundVoice() *TLInputMessagesFilterRoundVoice {
	return &TLInputMessagesFilterRoundVoice{
		Data2: m,
	}
}

//  inputMessagesFilterRoundVideo#b549da53 = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterRoundVideo() *TLInputMessagesFilterRoundVideo {
	return &TLInputMessagesFilterRoundVideo{
		Data2: m,
	}
}

//  inputMessagesFilterMyMentions#c1f8e69a = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterMyMentions() *TLInputMessagesFilterMyMentions {
	return &TLInputMessagesFilterMyMentions{
		Data2: m,
	}
}

//  inputMessagesFilterGeo#e7026d0d = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterGeo() *TLInputMessagesFilterGeo {
	return &TLInputMessagesFilterGeo{
		Data2: m,
	}
}

//  inputMessagesFilterContacts#e062db83 = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterContacts() *TLInputMessagesFilterContacts {
	return &TLInputMessagesFilterContacts{
		Data2: m,
	}
}

//  inputMessagesFilterPinned#1bb00451 = MessagesFilter;
//
func (m *MessagesFilter) To_InputMessagesFilterPinned() *TLInputMessagesFilterPinned {
	return &TLInputMessagesFilterPinned{
		Data2: m,
	}
}

//  inputMessagesFilterEmpty#57e2f66c = MessagesFilter;
//
func (m *TLInputMessagesFilterEmpty) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterEmpty
	return m.Data2
}

//  inputMessagesFilterPhotos#9609a51c = MessagesFilter;
//
func (m *TLInputMessagesFilterPhotos) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterPhotos
	return m.Data2
}

//  inputMessagesFilterVideo#9fc00e65 = MessagesFilter;
//
func (m *TLInputMessagesFilterVideo) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterVideo
	return m.Data2
}

//  inputMessagesFilterPhotoVideo#56e9f0e4 = MessagesFilter;
//
func (m *TLInputMessagesFilterPhotoVideo) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterPhotoVideo
	return m.Data2
}

//  inputMessagesFilterPhotoVideoDocuments#d95e73bb = MessagesFilter;
//
func (m *TLInputMessagesFilterPhotoVideoDocuments) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterPhotoVideoDocuments
	return m.Data2
}

//  inputMessagesFilterDocument#9eddf188 = MessagesFilter;
//
func (m *TLInputMessagesFilterDocument) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterDocument
	return m.Data2
}

//  inputMessagesFilterUrl#7ef0dd87 = MessagesFilter;
//
func (m *TLInputMessagesFilterUrl) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterUrl
	return m.Data2
}

//  inputMessagesFilterGif#ffc86587 = MessagesFilter;
//
func (m *TLInputMessagesFilterGif) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterGif
	return m.Data2
}

//  inputMessagesFilterVoice#50f5c392 = MessagesFilter;
//
func (m *TLInputMessagesFilterVoice) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterVoice
	return m.Data2
}

//  inputMessagesFilterMusic#3751b49e = MessagesFilter;
//
func (m *TLInputMessagesFilterMusic) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterMusic
	return m.Data2
}

//  inputMessagesFilterChatPhotos#3a20ecb8 = MessagesFilter;
//
func (m *TLInputMessagesFilterChatPhotos) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterChatPhotos
	return m.Data2
}

//  inputMessagesFilterPhoneCalls#80c99768 flags:# missed:flags.0?true = MessagesFilter;
//
func (m *TLInputMessagesFilterPhoneCalls) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterPhoneCalls
	return m.Data2
}

//  inputMessagesFilterRoundVoice#7a7c17a4 = MessagesFilter;
//
func (m *TLInputMessagesFilterRoundVoice) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterRoundVoice
	return m.Data2
}

//  inputMessagesFilterRoundVideo#b549da53 = MessagesFilter;
//
func (m *TLInputMessagesFilterRoundVideo) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterRoundVideo
	return m.Data2
}

//  inputMessagesFilterMyMentions#c1f8e69a = MessagesFilter;
//
func (m *TLInputMessagesFilterMyMentions) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterMyMentions
	return m.Data2
}

//  inputMessagesFilterGeo#e7026d0d = MessagesFilter;
//
func (m *TLInputMessagesFilterGeo) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterGeo
	return m.Data2
}

//  inputMessagesFilterContacts#e062db83 = MessagesFilter;
//
func (m *TLInputMessagesFilterContacts) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterContacts
	return m.Data2
}

//  inputMessagesFilterPinned#1bb00451 = MessagesFilter;
//
func (m *TLInputMessagesFilterPinned) To_MessagesFilter() *MessagesFilter {
	m.Data2.Cmd = Cmd_InputMessagesFilterPinned
	return m.Data2
}

//  + TL_MessagesAffectedHistory
//

//  messages.affectedHistory#b45c69d1 pts:int pts_count:int offset:int = messages.AffectedHistory;
//
func (m *Messages_AffectedHistory) To_MessagesAffectedHistory() *TLMessagesAffectedHistory {
	return &TLMessagesAffectedHistory{
		Data2: m,
	}
}

//  messages.affectedHistory#b45c69d1 pts:int pts_count:int offset:int = messages.AffectedHistory;
//
func (m *TLMessagesAffectedHistory) To_Messages_AffectedHistory() *Messages_AffectedHistory {
	m.Data2.Cmd = Cmd_MessagesAffectedHistory
	return m.Data2
}

//  + TL_MessagesAffectedMessages
//

//  messages.affectedMessages#84d19185 pts:int pts_count:int = messages.AffectedMessages;
//
func (m *Messages_AffectedMessages) To_MessagesAffectedMessages() *TLMessagesAffectedMessages {
	return &TLMessagesAffectedMessages{
		Data2: m,
	}
}

//  messages.affectedMessages#84d19185 pts:int pts_count:int = messages.AffectedMessages;
//
func (m *TLMessagesAffectedMessages) To_Messages_AffectedMessages() *Messages_AffectedMessages {
	m.Data2.Cmd = Cmd_MessagesAffectedMessages
	return m.Data2
}

//  + TL_MessagesAllStickersNotModified
//  + TL_MessagesAllStickers
//

//  messages.allStickersNotModified#e86602c3 = messages.AllStickers;
//
func (m *Messages_AllStickers) To_MessagesAllStickersNotModified() *TLMessagesAllStickersNotModified {
	return &TLMessagesAllStickersNotModified{
		Data2: m,
	}
}

//  messages.allStickers#edfd405f hash:int sets:Vector<StickerSet> = messages.AllStickers;
//
func (m *Messages_AllStickers) To_MessagesAllStickers() *TLMessagesAllStickers {
	return &TLMessagesAllStickers{
		Data2: m,
	}
}

//  messages.allStickersNotModified#e86602c3 = messages.AllStickers;
//
func (m *TLMessagesAllStickersNotModified) To_Messages_AllStickers() *Messages_AllStickers {
	m.Data2.Cmd = Cmd_MessagesAllStickersNotModified
	return m.Data2
}

//  messages.allStickers#edfd405f hash:int sets:Vector<StickerSet> = messages.AllStickers;
//
func (m *TLMessagesAllStickers) To_Messages_AllStickers() *Messages_AllStickers {
	m.Data2.Cmd = Cmd_MessagesAllStickers
	return m.Data2
}

//  + TL_MessagesArchivedStickers
//

//  messages.archivedStickers#4fcba9c8 count:int sets:Vector<StickerSetCovered> = messages.ArchivedStickers;
//
func (m *Messages_ArchivedStickers) To_MessagesArchivedStickers() *TLMessagesArchivedStickers {
	return &TLMessagesArchivedStickers{
		Data2: m,
	}
}

//  messages.archivedStickers#4fcba9c8 count:int sets:Vector<StickerSetCovered> = messages.ArchivedStickers;
//
func (m *TLMessagesArchivedStickers) To_Messages_ArchivedStickers() *Messages_ArchivedStickers {
	m.Data2.Cmd = Cmd_MessagesArchivedStickers
	return m.Data2
}

//  + TL_MessagesBotCallbackAnswer
//

//  messages.botCallbackAnswer#36585ea4 flags:# alert:flags.1?true has_url:flags.3?true message:flags.0?string url:flags.2?string cache_time:int = messages.BotCallbackAnswer;
//  messages.botCallbackAnswer#1264f1c6 flags:# alert:flags.1?true message:flags.0?string = messages.BotCallbackAnswer;
//
func (m *Messages_BotCallbackAnswer) To_MessagesBotCallbackAnswer() *TLMessagesBotCallbackAnswer {
	return &TLMessagesBotCallbackAnswer{
		Data2: m,
	}
}

//  messages.botCallbackAnswer#36585ea4 flags:# alert:flags.1?true has_url:flags.3?true message:flags.0?string url:flags.2?string cache_time:int = messages.BotCallbackAnswer;
//  messages.botCallbackAnswer#1264f1c6 flags:# alert:flags.1?true message:flags.0?string = messages.BotCallbackAnswer;
//
func (m *TLMessagesBotCallbackAnswer) To_Messages_BotCallbackAnswer() *Messages_BotCallbackAnswer {
	m.Data2.Cmd = Cmd_MessagesBotCallbackAnswer
	return m.Data2
}

//  + TL_MessagesBotResults
//

//  messages.botResults#ccd3563d flags:# gallery:flags.0?true query_id:long next_offset:flags.1?string switch_pm:flags.2?InlineBotSwitchPM results:Vector<BotInlineResult> cache_time:int = messages.BotResults;
//  messages.botResults#256709a6 flags:# gallery:flags.0?true query_id:long next_offset:flags.1?string switch_pm:flags.2?InlineBotSwitchPM results:Vector<BotInlineResult> = messages.BotResults;
//  messages.botResults#947ca848 flags:# gallery:flags.0?true query_id:long next_offset:flags.1?string switch_pm:flags.2?InlineBotSwitchPM results:Vector<BotInlineResult> cache_time:int users:Vector<User> = messages.BotResults;
//
func (m *Messages_BotResults) To_MessagesBotResults() *TLMessagesBotResults {
	return &TLMessagesBotResults{
		Data2: m,
	}
}

//  messages.botResults#ccd3563d flags:# gallery:flags.0?true query_id:long next_offset:flags.1?string switch_pm:flags.2?InlineBotSwitchPM results:Vector<BotInlineResult> cache_time:int = messages.BotResults;
//  messages.botResults#256709a6 flags:# gallery:flags.0?true query_id:long next_offset:flags.1?string switch_pm:flags.2?InlineBotSwitchPM results:Vector<BotInlineResult> = messages.BotResults;
//  messages.botResults#947ca848 flags:# gallery:flags.0?true query_id:long next_offset:flags.1?string switch_pm:flags.2?InlineBotSwitchPM results:Vector<BotInlineResult> cache_time:int users:Vector<User> = messages.BotResults;
//
func (m *TLMessagesBotResults) To_Messages_BotResults() *Messages_BotResults {
	m.Data2.Cmd = Cmd_MessagesBotResults
	return m.Data2
}

//  + TL_MessagesChatFull
//

//  messages.chatFull#e5d7d19c full_chat:ChatFull chats:Vector<Chat> users:Vector<User> = messages.ChatFull;
//
func (m *Messages_ChatFull) To_MessagesChatFull() *TLMessagesChatFull {
	return &TLMessagesChatFull{
		Data2: m,
	}
}

//  messages.chatFull#e5d7d19c full_chat:ChatFull chats:Vector<Chat> users:Vector<User> = messages.ChatFull;
//
func (m *TLMessagesChatFull) To_Messages_ChatFull() *Messages_ChatFull {
	m.Data2.Cmd = Cmd_MessagesChatFull
	return m.Data2
}

//  + TL_MessagesChats
//  + TL_MessagesChatsSlice
//

//  messages.chats#64ff9fd5 chats:Vector<Chat> = messages.Chats;
//
func (m *Messages_Chats) To_MessagesChats() *TLMessagesChats {
	return &TLMessagesChats{
		Data2: m,
	}
}

//  messages.chatsSlice#9cd81144 count:int chats:Vector<Chat> = messages.Chats;
//
func (m *Messages_Chats) To_MessagesChatsSlice() *TLMessagesChatsSlice {
	return &TLMessagesChatsSlice{
		Data2: m,
	}
}

//  messages.chats#64ff9fd5 chats:Vector<Chat> = messages.Chats;
//
func (m *TLMessagesChats) To_Messages_Chats() *Messages_Chats {
	m.Data2.Cmd = Cmd_MessagesChats
	return m.Data2
}

//  messages.chatsSlice#9cd81144 count:int chats:Vector<Chat> = messages.Chats;
//
func (m *TLMessagesChatsSlice) To_Messages_Chats() *Messages_Chats {
	m.Data2.Cmd = Cmd_MessagesChatsSlice
	return m.Data2
}

//  + TL_MessagesDhConfigNotModified
//  + TL_MessagesDhConfig
//

//  messages.dhConfigNotModified#c0e24635 random:bytes = messages.DhConfig;
//
func (m *Messages_DhConfig) To_MessagesDhConfigNotModified() *TLMessagesDhConfigNotModified {
	return &TLMessagesDhConfigNotModified{
		Data2: m,
	}
}

//  messages.dhConfig#2c221edd g:int p:bytes version:int random:bytes = messages.DhConfig;
//
func (m *Messages_DhConfig) To_MessagesDhConfig() *TLMessagesDhConfig {
	return &TLMessagesDhConfig{
		Data2: m,
	}
}

//  messages.dhConfigNotModified#c0e24635 random:bytes = messages.DhConfig;
//
func (m *TLMessagesDhConfigNotModified) To_Messages_DhConfig() *Messages_DhConfig {
	m.Data2.Cmd = Cmd_MessagesDhConfigNotModified
	return m.Data2
}

//  messages.dhConfig#2c221edd g:int p:bytes version:int random:bytes = messages.DhConfig;
//
func (m *TLMessagesDhConfig) To_Messages_DhConfig() *Messages_DhConfig {
	m.Data2.Cmd = Cmd_MessagesDhConfig
	return m.Data2
}

//  + TL_MessagesDialogs
//  + TL_MessagesDialogsSlice
//  + TL_MessagesDialogsNotModified
//

//  messages.dialogs#15ba6c40 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
//
func (m *Messages_Dialogs) To_MessagesDialogs() *TLMessagesDialogs {
	return &TLMessagesDialogs{
		Data2: m,
	}
}

//  messages.dialogsSlice#71e094f3 count:int dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
//
func (m *Messages_Dialogs) To_MessagesDialogsSlice() *TLMessagesDialogsSlice {
	return &TLMessagesDialogsSlice{
		Data2: m,
	}
}

//  messages.dialogsNotModified#f0e3e596 count:int = messages.Dialogs;
//
func (m *Messages_Dialogs) To_MessagesDialogsNotModified() *TLMessagesDialogsNotModified {
	return &TLMessagesDialogsNotModified{
		Data2: m,
	}
}

//  messages.dialogs#15ba6c40 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
//
func (m *TLMessagesDialogs) To_Messages_Dialogs() *Messages_Dialogs {
	m.Data2.Cmd = Cmd_MessagesDialogs
	return m.Data2
}

//  messages.dialogsSlice#71e094f3 count:int dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
//
func (m *TLMessagesDialogsSlice) To_Messages_Dialogs() *Messages_Dialogs {
	m.Data2.Cmd = Cmd_MessagesDialogsSlice
	return m.Data2
}

//  messages.dialogsNotModified#f0e3e596 count:int = messages.Dialogs;
//
func (m *TLMessagesDialogsNotModified) To_Messages_Dialogs() *Messages_Dialogs {
	m.Data2.Cmd = Cmd_MessagesDialogsNotModified
	return m.Data2
}

//  + TL_MessagesDiscussionMessage
//

//  messages.discussionMessage#f5dd8f9d flags:# messages:Vector<Message> max_id:flags.0?int read_inbox_max_id:flags.1?int read_outbox_max_id:flags.2?int chats:Vector<Chat> users:Vector<User> = messages.DiscussionMessage;
//
func (m *Messages_DiscussionMessage) To_MessagesDiscussionMessage() *TLMessagesDiscussionMessage {
	return &TLMessagesDiscussionMessage{
		Data2: m,
	}
}

//  messages.discussionMessage#f5dd8f9d flags:# messages:Vector<Message> max_id:flags.0?int read_inbox_max_id:flags.1?int read_outbox_max_id:flags.2?int chats:Vector<Chat> users:Vector<User> = messages.DiscussionMessage;
//
func (m *TLMessagesDiscussionMessage) To_Messages_DiscussionMessage() *Messages_DiscussionMessage {
	m.Data2.Cmd = Cmd_MessagesDiscussionMessage
	return m.Data2
}

//  + TL_MessagesFavedStickersNotModified
//  + TL_MessagesFavedStickers
//

//  messages.favedStickersNotModified#9e8fa6d3 = messages.FavedStickers;
//
func (m *Messages_FavedStickers) To_MessagesFavedStickersNotModified() *TLMessagesFavedStickersNotModified {
	return &TLMessagesFavedStickersNotModified{
		Data2: m,
	}
}

//  messages.favedStickers#f37f2f16 hash:int packs:Vector<StickerPack> stickers:Vector<Document> = messages.FavedStickers;
//
func (m *Messages_FavedStickers) To_MessagesFavedStickers() *TLMessagesFavedStickers {
	return &TLMessagesFavedStickers{
		Data2: m,
	}
}

//  messages.favedStickersNotModified#9e8fa6d3 = messages.FavedStickers;
//
func (m *TLMessagesFavedStickersNotModified) To_Messages_FavedStickers() *Messages_FavedStickers {
	m.Data2.Cmd = Cmd_MessagesFavedStickersNotModified
	return m.Data2
}

//  messages.favedStickers#f37f2f16 hash:int packs:Vector<StickerPack> stickers:Vector<Document> = messages.FavedStickers;
//
func (m *TLMessagesFavedStickers) To_Messages_FavedStickers() *Messages_FavedStickers {
	m.Data2.Cmd = Cmd_MessagesFavedStickers
	return m.Data2
}

//  + TL_MessagesFeaturedStickersNotModified
//  + TL_MessagesFeaturedStickers
//

//  messages.featuredStickersNotModified#4ede3cf = messages.FeaturedStickers;
//  messages.featuredStickersNotModified#c6dc0c66 count:int = messages.FeaturedStickers;
//
func (m *Messages_FeaturedStickers) To_MessagesFeaturedStickersNotModified() *TLMessagesFeaturedStickersNotModified {
	return &TLMessagesFeaturedStickersNotModified{
		Data2: m,
	}
}

//  messages.featuredStickers#f89d88e5 hash:int sets:Vector<StickerSetCovered> unread:Vector<long> = messages.FeaturedStickers;
//  messages.featuredStickers#b6abc341 hash:int count:int sets:Vector<StickerSetCovered> unread:Vector<long> = messages.FeaturedStickers;
//
func (m *Messages_FeaturedStickers) To_MessagesFeaturedStickers() *TLMessagesFeaturedStickers {
	return &TLMessagesFeaturedStickers{
		Data2: m,
	}
}

//  messages.featuredStickersNotModified#4ede3cf = messages.FeaturedStickers;
//  messages.featuredStickersNotModified#c6dc0c66 count:int = messages.FeaturedStickers;
//
func (m *TLMessagesFeaturedStickersNotModified) To_Messages_FeaturedStickers() *Messages_FeaturedStickers {
	m.Data2.Cmd = Cmd_MessagesFeaturedStickersNotModified
	return m.Data2
}

//  messages.featuredStickers#f89d88e5 hash:int sets:Vector<StickerSetCovered> unread:Vector<long> = messages.FeaturedStickers;
//  messages.featuredStickers#b6abc341 hash:int count:int sets:Vector<StickerSetCovered> unread:Vector<long> = messages.FeaturedStickers;
//
func (m *TLMessagesFeaturedStickers) To_Messages_FeaturedStickers() *Messages_FeaturedStickers {
	m.Data2.Cmd = Cmd_MessagesFeaturedStickers
	return m.Data2
}

//  + TL_MessagesFoundGifs
//

//  messages.foundGifs#450a1c0a next_offset:int results:Vector<FoundGif> = messages.FoundGifs;
//
func (m *Messages_FoundGifs) To_MessagesFoundGifs() *TLMessagesFoundGifs {
	return &TLMessagesFoundGifs{
		Data2: m,
	}
}

//  messages.foundGifs#450a1c0a next_offset:int results:Vector<FoundGif> = messages.FoundGifs;
//
func (m *TLMessagesFoundGifs) To_Messages_FoundGifs() *Messages_FoundGifs {
	m.Data2.Cmd = Cmd_MessagesFoundGifs
	return m.Data2
}

//  + TL_MessagesFoundStickerSetsNotModified
//  + TL_MessagesFoundStickerSets
//

//  messages.foundStickerSetsNotModified#d54b65d = messages.FoundStickerSets;
//
func (m *Messages_FoundStickerSets) To_MessagesFoundStickerSetsNotModified() *TLMessagesFoundStickerSetsNotModified {
	return &TLMessagesFoundStickerSetsNotModified{
		Data2: m,
	}
}

//  messages.foundStickerSets#5108d648 hash:int sets:Vector<StickerSetCovered> = messages.FoundStickerSets;
//
func (m *Messages_FoundStickerSets) To_MessagesFoundStickerSets() *TLMessagesFoundStickerSets {
	return &TLMessagesFoundStickerSets{
		Data2: m,
	}
}

//  messages.foundStickerSetsNotModified#d54b65d = messages.FoundStickerSets;
//
func (m *TLMessagesFoundStickerSetsNotModified) To_Messages_FoundStickerSets() *Messages_FoundStickerSets {
	m.Data2.Cmd = Cmd_MessagesFoundStickerSetsNotModified
	return m.Data2
}

//  messages.foundStickerSets#5108d648 hash:int sets:Vector<StickerSetCovered> = messages.FoundStickerSets;
//
func (m *TLMessagesFoundStickerSets) To_Messages_FoundStickerSets() *Messages_FoundStickerSets {
	m.Data2.Cmd = Cmd_MessagesFoundStickerSets
	return m.Data2
}

//  + TL_MessagesHighScores
//

//  messages.highScores#9a3bfd99 scores:Vector<HighScore> users:Vector<User> = messages.HighScores;
//
func (m *Messages_HighScores) To_MessagesHighScores() *TLMessagesHighScores {
	return &TLMessagesHighScores{
		Data2: m,
	}
}

//  messages.highScores#9a3bfd99 scores:Vector<HighScore> users:Vector<User> = messages.HighScores;
//
func (m *TLMessagesHighScores) To_Messages_HighScores() *Messages_HighScores {
	m.Data2.Cmd = Cmd_MessagesHighScores
	return m.Data2
}

//  + TL_MessagesInactiveChats
//

//  messages.inactiveChats#a927fec5 dates:Vector<int> chats:Vector<Chat> users:Vector<User> = messages.InactiveChats;
//
func (m *Messages_InactiveChats) To_MessagesInactiveChats() *TLMessagesInactiveChats {
	return &TLMessagesInactiveChats{
		Data2: m,
	}
}

//  messages.inactiveChats#a927fec5 dates:Vector<int> chats:Vector<Chat> users:Vector<User> = messages.InactiveChats;
//
func (m *TLMessagesInactiveChats) To_Messages_InactiveChats() *Messages_InactiveChats {
	m.Data2.Cmd = Cmd_MessagesInactiveChats
	return m.Data2
}

//  + TL_MessagesMessageEditData
//

//  messages.messageEditData#26b5dde6 flags:# caption:flags.0?true = messages.MessageEditData;
//
func (m *Messages_MessageEditData) To_MessagesMessageEditData() *TLMessagesMessageEditData {
	return &TLMessagesMessageEditData{
		Data2: m,
	}
}

//  messages.messageEditData#26b5dde6 flags:# caption:flags.0?true = messages.MessageEditData;
//
func (m *TLMessagesMessageEditData) To_Messages_MessageEditData() *Messages_MessageEditData {
	m.Data2.Cmd = Cmd_MessagesMessageEditData
	return m.Data2
}

//  + TL_MessagesMessageViews
//

//  messages.messageViews#b6c4f543 views:Vector<MessageViews> chats:Vector<Chat> users:Vector<User> = messages.MessageViews;
//
func (m *Messages_MessageViews) To_MessagesMessageViews() *TLMessagesMessageViews {
	return &TLMessagesMessageViews{
		Data2: m,
	}
}

//  messages.messageViews#b6c4f543 views:Vector<MessageViews> chats:Vector<Chat> users:Vector<User> = messages.MessageViews;
//
func (m *TLMessagesMessageViews) To_Messages_MessageViews() *Messages_MessageViews {
	m.Data2.Cmd = Cmd_MessagesMessageViews
	return m.Data2
}

//  + TL_MessagesMessages
//  + TL_MessagesMessagesSlice
//  + TL_MessagesChannelMessages
//  + TL_MessagesMessagesNotModified
//

//  messages.messages#8c718e87 messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//
func (m *Messages_Messages) To_MessagesMessages() *TLMessagesMessages {
	return &TLMessagesMessages{
		Data2: m,
	}
}

//  messages.messagesSlice#b446ae3 count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//  messages.messagesSlice#a6c47aaa flags:# inexact:flags.1?true count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//  messages.messagesSlice#c8edce1e flags:# inexact:flags.1?true count:int next_rate:flags.0?int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//  messages.messagesSlice#3a54685e flags:# inexact:flags.1?true count:int next_rate:flags.0?int offset_id_offset:flags.2?int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//
func (m *Messages_Messages) To_MessagesMessagesSlice() *TLMessagesMessagesSlice {
	return &TLMessagesMessagesSlice{
		Data2: m,
	}
}

//  messages.channelMessages#99262e37 flags:# pts:int count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//  messages.channelMessages#bc0f17bc flags:# pts:int count:int messages:Vector<Message> collapsed:flags.0?Vector<MessageGroup> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//  messages.channelMessages#64479808 flags:# inexact:flags.1?true pts:int count:int offset_id_offset:flags.2?int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//
func (m *Messages_Messages) To_MessagesChannelMessages() *TLMessagesChannelMessages {
	return &TLMessagesChannelMessages{
		Data2: m,
	}
}

//  messages.messagesNotModified#74535f21 count:int = messages.Messages;
//
func (m *Messages_Messages) To_MessagesMessagesNotModified() *TLMessagesMessagesNotModified {
	return &TLMessagesMessagesNotModified{
		Data2: m,
	}
}

//  messages.messages#8c718e87 messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//
func (m *TLMessagesMessages) To_Messages_Messages() *Messages_Messages {
	m.Data2.Cmd = Cmd_MessagesMessages
	return m.Data2
}

//  messages.messagesSlice#b446ae3 count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//  messages.messagesSlice#a6c47aaa flags:# inexact:flags.1?true count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//  messages.messagesSlice#c8edce1e flags:# inexact:flags.1?true count:int next_rate:flags.0?int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//  messages.messagesSlice#3a54685e flags:# inexact:flags.1?true count:int next_rate:flags.0?int offset_id_offset:flags.2?int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//
func (m *TLMessagesMessagesSlice) To_Messages_Messages() *Messages_Messages {
	m.Data2.Cmd = Cmd_MessagesMessagesSlice
	return m.Data2
}

//  messages.channelMessages#99262e37 flags:# pts:int count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//  messages.channelMessages#bc0f17bc flags:# pts:int count:int messages:Vector<Message> collapsed:flags.0?Vector<MessageGroup> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//  messages.channelMessages#64479808 flags:# inexact:flags.1?true pts:int count:int offset_id_offset:flags.2?int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
//
func (m *TLMessagesChannelMessages) To_Messages_Messages() *Messages_Messages {
	m.Data2.Cmd = Cmd_MessagesChannelMessages
	return m.Data2
}

//  messages.messagesNotModified#74535f21 count:int = messages.Messages;
//
func (m *TLMessagesMessagesNotModified) To_Messages_Messages() *Messages_Messages {
	m.Data2.Cmd = Cmd_MessagesMessagesNotModified
	return m.Data2
}

//  + TL_MessagesPeerDialogs
//

//  messages.peerDialogs#3371c354 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> state:updates.State = messages.PeerDialogs;
//
func (m *Messages_PeerDialogs) To_MessagesPeerDialogs() *TLMessagesPeerDialogs {
	return &TLMessagesPeerDialogs{
		Data2: m,
	}
}

//  messages.peerDialogs#3371c354 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> state:updates.State = messages.PeerDialogs;
//
func (m *TLMessagesPeerDialogs) To_Messages_PeerDialogs() *Messages_PeerDialogs {
	m.Data2.Cmd = Cmd_MessagesPeerDialogs
	return m.Data2
}

//  + TL_MessagesRecentStickersNotModified
//  + TL_MessagesRecentStickers
//

//  messages.recentStickersNotModified#b17f890 = messages.RecentStickers;
//
func (m *Messages_RecentStickers) To_MessagesRecentStickersNotModified() *TLMessagesRecentStickersNotModified {
	return &TLMessagesRecentStickersNotModified{
		Data2: m,
	}
}

//  messages.recentStickers#5ce20970 hash:int stickers:Vector<Document> = messages.RecentStickers;
//  messages.recentStickers#22f3afb3 hash:int packs:Vector<StickerPack> stickers:Vector<Document> dates:Vector<int> = messages.RecentStickers;
//
func (m *Messages_RecentStickers) To_MessagesRecentStickers() *TLMessagesRecentStickers {
	return &TLMessagesRecentStickers{
		Data2: m,
	}
}

//  messages.recentStickersNotModified#b17f890 = messages.RecentStickers;
//
func (m *TLMessagesRecentStickersNotModified) To_Messages_RecentStickers() *Messages_RecentStickers {
	m.Data2.Cmd = Cmd_MessagesRecentStickersNotModified
	return m.Data2
}

//  messages.recentStickers#5ce20970 hash:int stickers:Vector<Document> = messages.RecentStickers;
//  messages.recentStickers#22f3afb3 hash:int packs:Vector<StickerPack> stickers:Vector<Document> dates:Vector<int> = messages.RecentStickers;
//
func (m *TLMessagesRecentStickers) To_Messages_RecentStickers() *Messages_RecentStickers {
	m.Data2.Cmd = Cmd_MessagesRecentStickers
	return m.Data2
}

//  + TL_MessagesSavedGifsNotModified
//  + TL_MessagesSavedGifs
//

//  messages.savedGifsNotModified#e8025ca2 = messages.SavedGifs;
//
func (m *Messages_SavedGifs) To_MessagesSavedGifsNotModified() *TLMessagesSavedGifsNotModified {
	return &TLMessagesSavedGifsNotModified{
		Data2: m,
	}
}

//  messages.savedGifs#2e0709a5 hash:int gifs:Vector<Document> = messages.SavedGifs;
//
func (m *Messages_SavedGifs) To_MessagesSavedGifs() *TLMessagesSavedGifs {
	return &TLMessagesSavedGifs{
		Data2: m,
	}
}

//  messages.savedGifsNotModified#e8025ca2 = messages.SavedGifs;
//
func (m *TLMessagesSavedGifsNotModified) To_Messages_SavedGifs() *Messages_SavedGifs {
	m.Data2.Cmd = Cmd_MessagesSavedGifsNotModified
	return m.Data2
}

//  messages.savedGifs#2e0709a5 hash:int gifs:Vector<Document> = messages.SavedGifs;
//
func (m *TLMessagesSavedGifs) To_Messages_SavedGifs() *Messages_SavedGifs {
	m.Data2.Cmd = Cmd_MessagesSavedGifs
	return m.Data2
}

//  + TL_MessagesSearchCounter
//

//  messages.searchCounter#e844ebff flags:# inexact:flags.1?true filter:MessagesFilter count:int = messages.SearchCounter;
//
func (m *Messages_SearchCounter) To_MessagesSearchCounter() *TLMessagesSearchCounter {
	return &TLMessagesSearchCounter{
		Data2: m,
	}
}

//  messages.searchCounter#e844ebff flags:# inexact:flags.1?true filter:MessagesFilter count:int = messages.SearchCounter;
//
func (m *TLMessagesSearchCounter) To_Messages_SearchCounter() *Messages_SearchCounter {
	m.Data2.Cmd = Cmd_MessagesSearchCounter
	return m.Data2
}

//  + TL_MessagesSentEncryptedMessage
//  + TL_MessagesSentEncryptedFile
//

//  messages.sentEncryptedMessage#560f8935 date:int = messages.SentEncryptedMessage;
//
func (m *Messages_SentEncryptedMessage) To_MessagesSentEncryptedMessage() *TLMessagesSentEncryptedMessage {
	return &TLMessagesSentEncryptedMessage{
		Data2: m,
	}
}

//  messages.sentEncryptedFile#9493ff32 date:int file:EncryptedFile = messages.SentEncryptedMessage;
//
func (m *Messages_SentEncryptedMessage) To_MessagesSentEncryptedFile() *TLMessagesSentEncryptedFile {
	return &TLMessagesSentEncryptedFile{
		Data2: m,
	}
}

//  messages.sentEncryptedMessage#560f8935 date:int = messages.SentEncryptedMessage;
//
func (m *TLMessagesSentEncryptedMessage) To_Messages_SentEncryptedMessage() *Messages_SentEncryptedMessage {
	m.Data2.Cmd = Cmd_MessagesSentEncryptedMessage
	return m.Data2
}

//  messages.sentEncryptedFile#9493ff32 date:int file:EncryptedFile = messages.SentEncryptedMessage;
//
func (m *TLMessagesSentEncryptedFile) To_Messages_SentEncryptedMessage() *Messages_SentEncryptedMessage {
	m.Data2.Cmd = Cmd_MessagesSentEncryptedFile
	return m.Data2
}

//  + TL_MessagesStickerSet
//

//  messages.stickerSet#b60a24a6 set:StickerSet packs:Vector<StickerPack> documents:Vector<Document> = messages.StickerSet;
//
func (m *Messages_StickerSet) To_MessagesStickerSet() *TLMessagesStickerSet {
	return &TLMessagesStickerSet{
		Data2: m,
	}
}

//  messages.stickerSet#b60a24a6 set:StickerSet packs:Vector<StickerPack> documents:Vector<Document> = messages.StickerSet;
//
func (m *TLMessagesStickerSet) To_Messages_StickerSet() *Messages_StickerSet {
	m.Data2.Cmd = Cmd_MessagesStickerSet
	return m.Data2
}

//  + TL_MessagesStickerSetInstallResultSuccess
//  + TL_MessagesStickerSetInstallResultArchive
//

//  messages.stickerSetInstallResultSuccess#38641628 = messages.StickerSetInstallResult;
//
func (m *Messages_StickerSetInstallResult) To_MessagesStickerSetInstallResultSuccess() *TLMessagesStickerSetInstallResultSuccess {
	return &TLMessagesStickerSetInstallResultSuccess{
		Data2: m,
	}
}

//  messages.stickerSetInstallResultArchive#35e410a8 sets:Vector<StickerSetCovered> = messages.StickerSetInstallResult;
//
func (m *Messages_StickerSetInstallResult) To_MessagesStickerSetInstallResultArchive() *TLMessagesStickerSetInstallResultArchive {
	return &TLMessagesStickerSetInstallResultArchive{
		Data2: m,
	}
}

//  messages.stickerSetInstallResultSuccess#38641628 = messages.StickerSetInstallResult;
//
func (m *TLMessagesStickerSetInstallResultSuccess) To_Messages_StickerSetInstallResult() *Messages_StickerSetInstallResult {
	m.Data2.Cmd = Cmd_MessagesStickerSetInstallResultSuccess
	return m.Data2
}

//  messages.stickerSetInstallResultArchive#35e410a8 sets:Vector<StickerSetCovered> = messages.StickerSetInstallResult;
//
func (m *TLMessagesStickerSetInstallResultArchive) To_Messages_StickerSetInstallResult() *Messages_StickerSetInstallResult {
	m.Data2.Cmd = Cmd_MessagesStickerSetInstallResultArchive
	return m.Data2
}

//  + TL_MessagesStickersNotModified
//  + TL_MessagesStickers
//

//  messages.stickersNotModified#f1749a22 = messages.Stickers;
//
func (m *Messages_Stickers) To_MessagesStickersNotModified() *TLMessagesStickersNotModified {
	return &TLMessagesStickersNotModified{
		Data2: m,
	}
}

//  messages.stickers#8a8ecd32 hash:string stickers:Vector<Document> = messages.Stickers;
//  messages.stickers#e4599bbd hash:int stickers:Vector<Document> = messages.Stickers;
//
func (m *Messages_Stickers) To_MessagesStickers() *TLMessagesStickers {
	return &TLMessagesStickers{
		Data2: m,
	}
}

//  messages.stickersNotModified#f1749a22 = messages.Stickers;
//
func (m *TLMessagesStickersNotModified) To_Messages_Stickers() *Messages_Stickers {
	m.Data2.Cmd = Cmd_MessagesStickersNotModified
	return m.Data2
}

//  messages.stickers#8a8ecd32 hash:string stickers:Vector<Document> = messages.Stickers;
//  messages.stickers#e4599bbd hash:int stickers:Vector<Document> = messages.Stickers;
//
func (m *TLMessagesStickers) To_Messages_Stickers() *Messages_Stickers {
	m.Data2.Cmd = Cmd_MessagesStickers
	return m.Data2
}

//  + TL_MessagesVotesList
//

//  messages.votesList#823f649 flags:# count:int votes:Vector<MessageUserVote> users:Vector<User> next_offset:flags.0?string = messages.VotesList;
//
func (m *Messages_VotesList) To_MessagesVotesList() *TLMessagesVotesList {
	return &TLMessagesVotesList{
		Data2: m,
	}
}

//  messages.votesList#823f649 flags:# count:int votes:Vector<MessageUserVote> users:Vector<User> next_offset:flags.0?string = messages.VotesList;
//
func (m *TLMessagesVotesList) To_Messages_VotesList() *Messages_VotesList {
	m.Data2.Cmd = Cmd_MessagesVotesList
	return m.Data2
}

//  + TL_MsgDetailedInfo
//  + TL_MsgNewDetailedInfo
//

//  msg_detailed_info#276d3ec6 msg_id:long answer_msg_id:long bytes:int status:int = MsgDetailedInfo;
//
func (m *MsgDetailedInfo) To_MsgDetailedInfo() *TLMsgDetailedInfo {
	return &TLMsgDetailedInfo{
		Data2: m,
	}
}

//  msg_new_detailed_info#809db6df answer_msg_id:long bytes:int status:int = MsgDetailedInfo;
//
func (m *MsgDetailedInfo) To_MsgNewDetailedInfo() *TLMsgNewDetailedInfo {
	return &TLMsgNewDetailedInfo{
		Data2: m,
	}
}

//  msg_detailed_info#276d3ec6 msg_id:long answer_msg_id:long bytes:int status:int = MsgDetailedInfo;
//
func (m *TLMsgDetailedInfo) To_MsgDetailedInfo() *MsgDetailedInfo {
	m.Data2.Cmd = Cmd_MsgDetailedInfo
	return m.Data2
}

//  msg_new_detailed_info#809db6df answer_msg_id:long bytes:int status:int = MsgDetailedInfo;
//
func (m *TLMsgNewDetailedInfo) To_MsgDetailedInfo() *MsgDetailedInfo {
	m.Data2.Cmd = Cmd_MsgNewDetailedInfo
	return m.Data2
}

//  + TL_MsgResendReq
//

//  msg_resend_req#7d861a08 msg_ids:Vector<long> = MsgResendReq;
//
func (m *MsgResendReq) To_MsgResendReq() *TLMsgResendReq {
	return &TLMsgResendReq{
		Data2: m,
	}
}

//  msg_resend_req#7d861a08 msg_ids:Vector<long> = MsgResendReq;
//
func (m *TLMsgResendReq) To_MsgResendReq() *MsgResendReq {
	m.Data2.Cmd = Cmd_MsgResendReq
	return m.Data2
}

//  + TL_MsgsAck
//

//  msgs_ack#62d6b459 msg_ids:Vector<long> = MsgsAck;
//
func (m *MsgsAck) To_MsgsAck() *TLMsgsAck {
	return &TLMsgsAck{
		Data2: m,
	}
}

//  msgs_ack#62d6b459 msg_ids:Vector<long> = MsgsAck;
//
func (m *TLMsgsAck) To_MsgsAck() *MsgsAck {
	m.Data2.Cmd = Cmd_MsgsAck
	return m.Data2
}

//  + TL_MsgsAllInfo
//

//  msgs_all_info#8cc0d131 msg_ids:Vector<long> info:string = MsgsAllInfo;
//
func (m *MsgsAllInfo) To_MsgsAllInfo() *TLMsgsAllInfo {
	return &TLMsgsAllInfo{
		Data2: m,
	}
}

//  msgs_all_info#8cc0d131 msg_ids:Vector<long> info:string = MsgsAllInfo;
//
func (m *TLMsgsAllInfo) To_MsgsAllInfo() *MsgsAllInfo {
	m.Data2.Cmd = Cmd_MsgsAllInfo
	return m.Data2
}

//  + TL_MsgsStateInfo
//

//  msgs_state_info#04deb57d req_msg_id:long info:string = MsgsStateInfo;
//
func (m *MsgsStateInfo) To_MsgsStateInfo() *TLMsgsStateInfo {
	return &TLMsgsStateInfo{
		Data2: m,
	}
}

//  msgs_state_info#04deb57d req_msg_id:long info:string = MsgsStateInfo;
//
func (m *TLMsgsStateInfo) To_MsgsStateInfo() *MsgsStateInfo {
	m.Data2.Cmd = Cmd_MsgsStateInfo
	return m.Data2
}

//  + TL_MsgsStateReq
//

//  msgs_state_req#da69fb52 msg_ids:Vector<long> = MsgsStateReq;
//
func (m *MsgsStateReq) To_MsgsStateReq() *TLMsgsStateReq {
	return &TLMsgsStateReq{
		Data2: m,
	}
}

//  msgs_state_req#da69fb52 msg_ids:Vector<long> = MsgsStateReq;
//
func (m *TLMsgsStateReq) To_MsgsStateReq() *MsgsStateReq {
	m.Data2.Cmd = Cmd_MsgsStateReq
	return m.Data2
}

//  + TL_NearestDc
//

//  nearestDc#8e1a1775 country:string this_dc:int nearest_dc:int = NearestDc;
//
func (m *NearestDc) To_NearestDc() *TLNearestDc {
	return &TLNearestDc{
		Data2: m,
	}
}

//  nearestDc#8e1a1775 country:string this_dc:int nearest_dc:int = NearestDc;
//
func (m *TLNearestDc) To_NearestDc() *NearestDc {
	m.Data2.Cmd = Cmd_NearestDc
	return m.Data2
}

//  + TL_NewSessionCreated
//

//  new_session_created#9ec20908 first_msg_id:long unique_id:long server_salt:long = NewSession;
//
func (m *NewSession) To_NewSessionCreated() *TLNewSessionCreated {
	return &TLNewSessionCreated{
		Data2: m,
	}
}

//  new_session_created#9ec20908 first_msg_id:long unique_id:long server_salt:long = NewSession;
//
func (m *TLNewSessionCreated) To_NewSession() *NewSession {
	m.Data2.Cmd = Cmd_NewSessionCreated
	return m.Data2
}

//  + TL_NotifyPeer
//  + TL_NotifyUsers
//  + TL_NotifyChats
//  + TL_NotifyAll
//  + TL_NotifyBroadcasts
//

//  notifyPeer#9fd40bd8 peer:Peer = NotifyPeer;
//
func (m *NotifyPeer) To_NotifyPeer() *TLNotifyPeer {
	return &TLNotifyPeer{
		Data2: m,
	}
}

//  notifyUsers#b4c83b4c = NotifyPeer;
//
func (m *NotifyPeer) To_NotifyUsers() *TLNotifyUsers {
	return &TLNotifyUsers{
		Data2: m,
	}
}

//  notifyChats#c007cec3 = NotifyPeer;
//
func (m *NotifyPeer) To_NotifyChats() *TLNotifyChats {
	return &TLNotifyChats{
		Data2: m,
	}
}

//  notifyAll#74d07c60 = NotifyPeer;
//
func (m *NotifyPeer) To_NotifyAll() *TLNotifyAll {
	return &TLNotifyAll{
		Data2: m,
	}
}

//  notifyBroadcasts#d612e8ef = NotifyPeer;
//
func (m *NotifyPeer) To_NotifyBroadcasts() *TLNotifyBroadcasts {
	return &TLNotifyBroadcasts{
		Data2: m,
	}
}

//  notifyPeer#9fd40bd8 peer:Peer = NotifyPeer;
//
func (m *TLNotifyPeer) To_NotifyPeer() *NotifyPeer {
	m.Data2.Cmd = Cmd_NotifyPeer
	return m.Data2
}

//  notifyUsers#b4c83b4c = NotifyPeer;
//
func (m *TLNotifyUsers) To_NotifyPeer() *NotifyPeer {
	m.Data2.Cmd = Cmd_NotifyUsers
	return m.Data2
}

//  notifyChats#c007cec3 = NotifyPeer;
//
func (m *TLNotifyChats) To_NotifyPeer() *NotifyPeer {
	m.Data2.Cmd = Cmd_NotifyChats
	return m.Data2
}

//  notifyAll#74d07c60 = NotifyPeer;
//
func (m *TLNotifyAll) To_NotifyPeer() *NotifyPeer {
	m.Data2.Cmd = Cmd_NotifyAll
	return m.Data2
}

//  notifyBroadcasts#d612e8ef = NotifyPeer;
//
func (m *TLNotifyBroadcasts) To_NotifyPeer() *NotifyPeer {
	m.Data2.Cmd = Cmd_NotifyBroadcasts
	return m.Data2
}

//  + TL_Null
//

//  null#56730bcc = Null;
//
func (m *Null) To_Null() *TLNull {
	return &TLNull{
		Data2: m,
	}
}

//  null#56730bcc = Null;
//
func (m *TLNull) To_Null() *Null {
	m.Data2.Cmd = Cmd_Null
	return m.Data2
}

//  + TL_PQInnerData
//  + TL_PQInnerDataDc
//  + TL_PQInnerDataTemp
//  + TL_PQInnerDataTempDc
//

//  p_q_inner_data#83c95aec pq:string p:string q:string nonce:int128 server_nonce:int128 new_nonce:int256 = P_Q_inner_data;
//
func (m *P_QInnerData) To_PQInnerData() *TLPQInnerData {
	return &TLPQInnerData{
		Data2: m,
	}
}

//  p_q_inner_data_dc#a9f55f95 pq:string p:string q:string nonce:int128 server_nonce:int128 new_nonce:int256 dc:int = P_Q_inner_data;
//
func (m *P_QInnerData) To_PQInnerDataDc() *TLPQInnerDataDc {
	return &TLPQInnerDataDc{
		Data2: m,
	}
}

//  p_q_inner_data_temp#3c6a84d4 pq:string p:string q:string nonce:int128 server_nonce:int128 new_nonce:int256 expires_in:int = P_Q_inner_data;
//
func (m *P_QInnerData) To_PQInnerDataTemp() *TLPQInnerDataTemp {
	return &TLPQInnerDataTemp{
		Data2: m,
	}
}

//  p_q_inner_data_temp_dc#56fddf88 pq:string p:string q:string nonce:int128 server_nonce:int128 new_nonce:int256 dc:int expires_in:int = P_Q_inner_data;
//
func (m *P_QInnerData) To_PQInnerDataTempDc() *TLPQInnerDataTempDc {
	return &TLPQInnerDataTempDc{
		Data2: m,
	}
}

//  p_q_inner_data#83c95aec pq:string p:string q:string nonce:int128 server_nonce:int128 new_nonce:int256 = P_Q_inner_data;
//
func (m *TLPQInnerData) To_P_QInnerData() *P_QInnerData {
	m.Data2.Cmd = Cmd_PQInnerData
	return m.Data2
}

//  p_q_inner_data_dc#a9f55f95 pq:string p:string q:string nonce:int128 server_nonce:int128 new_nonce:int256 dc:int = P_Q_inner_data;
//
func (m *TLPQInnerDataDc) To_P_QInnerData() *P_QInnerData {
	m.Data2.Cmd = Cmd_PQInnerDataDc
	return m.Data2
}

//  p_q_inner_data_temp#3c6a84d4 pq:string p:string q:string nonce:int128 server_nonce:int128 new_nonce:int256 expires_in:int = P_Q_inner_data;
//
func (m *TLPQInnerDataTemp) To_P_QInnerData() *P_QInnerData {
	m.Data2.Cmd = Cmd_PQInnerDataTemp
	return m.Data2
}

//  p_q_inner_data_temp_dc#56fddf88 pq:string p:string q:string nonce:int128 server_nonce:int128 new_nonce:int256 dc:int expires_in:int = P_Q_inner_data;
//
func (m *TLPQInnerDataTempDc) To_P_QInnerData() *P_QInnerData {
	m.Data2.Cmd = Cmd_PQInnerDataTempDc
	return m.Data2
}

//  + TL_PagePart
//  + TL_PageFull
//  + TL_Page
//

//  pagePart#8e3f9ebe blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
//
func (m *Page) To_PagePart() *TLPagePart {
	return &TLPagePart{
		Data2: m,
	}
}

//  pageFull#556ec7aa blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
//
func (m *Page) To_PageFull() *TLPageFull {
	return &TLPageFull{
		Data2: m,
	}
}

//  page#f199a0a8 flags:# part:flags.0?true rtl:flags.1?true blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
//  page#ae891bec flags:# part:flags.0?true rtl:flags.1?true url:string blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
//  page#98657f0d flags:# part:flags.0?true rtl:flags.1?true v2:flags.2?true url:string blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> views:flags.3?int = Page;
//
func (m *Page) To_Page() *TLPage {
	return &TLPage{
		Data2: m,
	}
}

//  pagePart#8e3f9ebe blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
//
func (m *TLPagePart) To_Page() *Page {
	m.Data2.Cmd = Cmd_PagePart
	return m.Data2
}

//  pageFull#556ec7aa blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
//
func (m *TLPageFull) To_Page() *Page {
	m.Data2.Cmd = Cmd_PageFull
	return m.Data2
}

//  page#f199a0a8 flags:# part:flags.0?true rtl:flags.1?true blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
//  page#ae891bec flags:# part:flags.0?true rtl:flags.1?true url:string blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
//  page#98657f0d flags:# part:flags.0?true rtl:flags.1?true v2:flags.2?true url:string blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> views:flags.3?int = Page;
//
func (m *TLPage) To_Page() *Page {
	m.Data2.Cmd = Cmd_Page
	return m.Data2
}

//  + TL_PageBlockUnsupported
//  + TL_PageBlockTitle
//  + TL_PageBlockSubtitle
//  + TL_PageBlockAuthorDate
//  + TL_PageBlockHeader
//  + TL_PageBlockSubheader
//  + TL_PageBlockParagraph
//  + TL_PageBlockPreformatted
//  + TL_PageBlockFooter
//  + TL_PageBlockDivider
//  + TL_PageBlockAnchor
//  + TL_PageBlockList
//  + TL_PageBlockBlockquote
//  + TL_PageBlockPullquote
//  + TL_PageBlockPhoto
//  + TL_PageBlockVideo
//  + TL_PageBlockCover
//  + TL_PageBlockEmbed
//  + TL_PageBlockEmbedPost
//  + TL_PageBlockCollage
//  + TL_PageBlockSlideshow
//  + TL_PageBlockChannel
//  + TL_PageBlockAudio
//  + TL_PageBlockKicker
//  + TL_PageBlockTable
//  + TL_PageBlockOrderedList
//  + TL_PageBlockDetails
//  + TL_PageBlockRelatedArticles
//  + TL_PageBlockMap
//

//  pageBlockUnsupported#13567e8a = PageBlock;
//
func (m *PageBlock) To_PageBlockUnsupported() *TLPageBlockUnsupported {
	return &TLPageBlockUnsupported{
		Data2: m,
	}
}

//  pageBlockTitle#70abc3fd text:RichText = PageBlock;
//
func (m *PageBlock) To_PageBlockTitle() *TLPageBlockTitle {
	return &TLPageBlockTitle{
		Data2: m,
	}
}

//  pageBlockSubtitle#8ffa9a1f text:RichText = PageBlock;
//
func (m *PageBlock) To_PageBlockSubtitle() *TLPageBlockSubtitle {
	return &TLPageBlockSubtitle{
		Data2: m,
	}
}

//  pageBlockAuthorDate#baafe5e0 author:RichText published_date:int = PageBlock;
//
func (m *PageBlock) To_PageBlockAuthorDate() *TLPageBlockAuthorDate {
	return &TLPageBlockAuthorDate{
		Data2: m,
	}
}

//  pageBlockHeader#bfd064ec text:RichText = PageBlock;
//
func (m *PageBlock) To_PageBlockHeader() *TLPageBlockHeader {
	return &TLPageBlockHeader{
		Data2: m,
	}
}

//  pageBlockSubheader#f12bb6e1 text:RichText = PageBlock;
//
func (m *PageBlock) To_PageBlockSubheader() *TLPageBlockSubheader {
	return &TLPageBlockSubheader{
		Data2: m,
	}
}

//  pageBlockParagraph#467a0766 text:RichText = PageBlock;
//
func (m *PageBlock) To_PageBlockParagraph() *TLPageBlockParagraph {
	return &TLPageBlockParagraph{
		Data2: m,
	}
}

//  pageBlockPreformatted#c070d93e text:RichText language:string = PageBlock;
//
func (m *PageBlock) To_PageBlockPreformatted() *TLPageBlockPreformatted {
	return &TLPageBlockPreformatted{
		Data2: m,
	}
}

//  pageBlockFooter#48870999 text:RichText = PageBlock;
//
func (m *PageBlock) To_PageBlockFooter() *TLPageBlockFooter {
	return &TLPageBlockFooter{
		Data2: m,
	}
}

//  pageBlockDivider#db20b188 = PageBlock;
//
func (m *PageBlock) To_PageBlockDivider() *TLPageBlockDivider {
	return &TLPageBlockDivider{
		Data2: m,
	}
}

//  pageBlockAnchor#ce0d37b0 name:string = PageBlock;
//
func (m *PageBlock) To_PageBlockAnchor() *TLPageBlockAnchor {
	return &TLPageBlockAnchor{
		Data2: m,
	}
}

//  pageBlockList#3a58c7f4 ordered:Bool items:Vector<RichText> = PageBlock;
//  pageBlockList#e4e88011 items:Vector<PageListItem> = PageBlock;
//
func (m *PageBlock) To_PageBlockList() *TLPageBlockList {
	return &TLPageBlockList{
		Data2: m,
	}
}

//  pageBlockBlockquote#263d7c26 text:RichText caption:RichText = PageBlock;
//
func (m *PageBlock) To_PageBlockBlockquote() *TLPageBlockBlockquote {
	return &TLPageBlockBlockquote{
		Data2: m,
	}
}

//  pageBlockPullquote#4f4456d3 text:RichText caption:RichText = PageBlock;
//
func (m *PageBlock) To_PageBlockPullquote() *TLPageBlockPullquote {
	return &TLPageBlockPullquote{
		Data2: m,
	}
}

//  pageBlockPhoto#e9c69982 photo_id:long caption:RichText = PageBlock;
//  pageBlockPhoto#1759c560 flags:# photo_id:long caption:PageCaption url:flags.0?string webpage_id:flags.0?long = PageBlock;
//
func (m *PageBlock) To_PageBlockPhoto() *TLPageBlockPhoto {
	return &TLPageBlockPhoto{
		Data2: m,
	}
}

//  pageBlockVideo#d9d71866 flags:# autoplay:flags.0?true loop:flags.1?true video_id:long caption:RichText = PageBlock;
//  pageBlockVideo#7c8fe7b6 flags:# autoplay:flags.0?true loop:flags.1?true video_id:long caption:PageCaption = PageBlock;
//
func (m *PageBlock) To_PageBlockVideo() *TLPageBlockVideo {
	return &TLPageBlockVideo{
		Data2: m,
	}
}

//  pageBlockCover#39f23300 cover:PageBlock = PageBlock;
//
func (m *PageBlock) To_PageBlockCover() *TLPageBlockCover {
	return &TLPageBlockCover{
		Data2: m,
	}
}

//  pageBlockEmbed#cde200d1 flags:# full_width:flags.0?true allow_scrolling:flags.3?true url:flags.1?string html:flags.2?string poster_photo_id:flags.4?long w:int h:int caption:RichText = PageBlock;
//  pageBlockEmbed#a8718dc5 flags:# full_width:flags.0?true allow_scrolling:flags.3?true url:flags.1?string html:flags.2?string poster_photo_id:flags.4?long w:flags.5?int h:flags.5?int caption:PageCaption = PageBlock;
//
func (m *PageBlock) To_PageBlockEmbed() *TLPageBlockEmbed {
	return &TLPageBlockEmbed{
		Data2: m,
	}
}

//  pageBlockEmbedPost#292c7be9 url:string webpage_id:long author_photo_id:long author:string date:int blocks:Vector<PageBlock> caption:RichText = PageBlock;
//  pageBlockEmbedPost#f259a80b url:string webpage_id:long author_photo_id:long author:string date:int blocks:Vector<PageBlock> caption:PageCaption = PageBlock;
//
func (m *PageBlock) To_PageBlockEmbedPost() *TLPageBlockEmbedPost {
	return &TLPageBlockEmbedPost{
		Data2: m,
	}
}

//  pageBlockCollage#8b31c4f items:Vector<PageBlock> caption:RichText = PageBlock;
//  pageBlockCollage#65a0fa4d items:Vector<PageBlock> caption:PageCaption = PageBlock;
//
func (m *PageBlock) To_PageBlockCollage() *TLPageBlockCollage {
	return &TLPageBlockCollage{
		Data2: m,
	}
}

//  pageBlockSlideshow#130c8963 items:Vector<PageBlock> caption:RichText = PageBlock;
//  pageBlockSlideshow#31f9590 items:Vector<PageBlock> caption:PageCaption = PageBlock;
//
func (m *PageBlock) To_PageBlockSlideshow() *TLPageBlockSlideshow {
	return &TLPageBlockSlideshow{
		Data2: m,
	}
}

//  pageBlockChannel#ef1751b5 channel:Chat = PageBlock;
//
func (m *PageBlock) To_PageBlockChannel() *TLPageBlockChannel {
	return &TLPageBlockChannel{
		Data2: m,
	}
}

//  pageBlockAudio#31b81a7f audio_id:long caption:RichText = PageBlock;
//  pageBlockAudio#804361ea audio_id:long caption:PageCaption = PageBlock;
//
func (m *PageBlock) To_PageBlockAudio() *TLPageBlockAudio {
	return &TLPageBlockAudio{
		Data2: m,
	}
}

//  pageBlockKicker#1e148390 text:RichText = PageBlock;
//
func (m *PageBlock) To_PageBlockKicker() *TLPageBlockKicker {
	return &TLPageBlockKicker{
		Data2: m,
	}
}

//  pageBlockTable#bf4dea82 flags:# bordered:flags.0?true striped:flags.1?true title:RichText rows:Vector<PageTableRow> = PageBlock;
//
func (m *PageBlock) To_PageBlockTable() *TLPageBlockTable {
	return &TLPageBlockTable{
		Data2: m,
	}
}

//  pageBlockOrderedList#9a8ae1e1 items:Vector<PageListOrderedItem> = PageBlock;
//
func (m *PageBlock) To_PageBlockOrderedList() *TLPageBlockOrderedList {
	return &TLPageBlockOrderedList{
		Data2: m,
	}
}

//  pageBlockDetails#76768bed flags:# open:flags.0?true blocks:Vector<PageBlock> title:RichText = PageBlock;
//
func (m *PageBlock) To_PageBlockDetails() *TLPageBlockDetails {
	return &TLPageBlockDetails{
		Data2: m,
	}
}

//  pageBlockRelatedArticles#16115a96 title:RichText articles:Vector<PageRelatedArticle> = PageBlock;
//
func (m *PageBlock) To_PageBlockRelatedArticles() *TLPageBlockRelatedArticles {
	return &TLPageBlockRelatedArticles{
		Data2: m,
	}
}

//  pageBlockMap#a44f3ef6 geo:GeoPoint zoom:int w:int h:int caption:PageCaption = PageBlock;
//
func (m *PageBlock) To_PageBlockMap() *TLPageBlockMap {
	return &TLPageBlockMap{
		Data2: m,
	}
}

//  pageBlockUnsupported#13567e8a = PageBlock;
//
func (m *TLPageBlockUnsupported) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockUnsupported
	return m.Data2
}

//  pageBlockTitle#70abc3fd text:RichText = PageBlock;
//
func (m *TLPageBlockTitle) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockTitle
	return m.Data2
}

//  pageBlockSubtitle#8ffa9a1f text:RichText = PageBlock;
//
func (m *TLPageBlockSubtitle) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockSubtitle
	return m.Data2
}

//  pageBlockAuthorDate#baafe5e0 author:RichText published_date:int = PageBlock;
//
func (m *TLPageBlockAuthorDate) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockAuthorDate
	return m.Data2
}

//  pageBlockHeader#bfd064ec text:RichText = PageBlock;
//
func (m *TLPageBlockHeader) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockHeader
	return m.Data2
}

//  pageBlockSubheader#f12bb6e1 text:RichText = PageBlock;
//
func (m *TLPageBlockSubheader) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockSubheader
	return m.Data2
}

//  pageBlockParagraph#467a0766 text:RichText = PageBlock;
//
func (m *TLPageBlockParagraph) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockParagraph
	return m.Data2
}

//  pageBlockPreformatted#c070d93e text:RichText language:string = PageBlock;
//
func (m *TLPageBlockPreformatted) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockPreformatted
	return m.Data2
}

//  pageBlockFooter#48870999 text:RichText = PageBlock;
//
func (m *TLPageBlockFooter) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockFooter
	return m.Data2
}

//  pageBlockDivider#db20b188 = PageBlock;
//
func (m *TLPageBlockDivider) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockDivider
	return m.Data2
}

//  pageBlockAnchor#ce0d37b0 name:string = PageBlock;
//
func (m *TLPageBlockAnchor) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockAnchor
	return m.Data2
}

//  pageBlockList#3a58c7f4 ordered:Bool items:Vector<RichText> = PageBlock;
//  pageBlockList#e4e88011 items:Vector<PageListItem> = PageBlock;
//
func (m *TLPageBlockList) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockList
	return m.Data2
}

//  pageBlockBlockquote#263d7c26 text:RichText caption:RichText = PageBlock;
//
func (m *TLPageBlockBlockquote) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockBlockquote
	return m.Data2
}

//  pageBlockPullquote#4f4456d3 text:RichText caption:RichText = PageBlock;
//
func (m *TLPageBlockPullquote) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockPullquote
	return m.Data2
}

//  pageBlockPhoto#e9c69982 photo_id:long caption:RichText = PageBlock;
//  pageBlockPhoto#1759c560 flags:# photo_id:long caption:PageCaption url:flags.0?string webpage_id:flags.0?long = PageBlock;
//
func (m *TLPageBlockPhoto) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockPhoto
	return m.Data2
}

//  pageBlockVideo#d9d71866 flags:# autoplay:flags.0?true loop:flags.1?true video_id:long caption:RichText = PageBlock;
//  pageBlockVideo#7c8fe7b6 flags:# autoplay:flags.0?true loop:flags.1?true video_id:long caption:PageCaption = PageBlock;
//
func (m *TLPageBlockVideo) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockVideo
	return m.Data2
}

//  pageBlockCover#39f23300 cover:PageBlock = PageBlock;
//
func (m *TLPageBlockCover) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockCover
	return m.Data2
}

//  pageBlockEmbed#cde200d1 flags:# full_width:flags.0?true allow_scrolling:flags.3?true url:flags.1?string html:flags.2?string poster_photo_id:flags.4?long w:int h:int caption:RichText = PageBlock;
//  pageBlockEmbed#a8718dc5 flags:# full_width:flags.0?true allow_scrolling:flags.3?true url:flags.1?string html:flags.2?string poster_photo_id:flags.4?long w:flags.5?int h:flags.5?int caption:PageCaption = PageBlock;
//
func (m *TLPageBlockEmbed) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockEmbed
	return m.Data2
}

//  pageBlockEmbedPost#292c7be9 url:string webpage_id:long author_photo_id:long author:string date:int blocks:Vector<PageBlock> caption:RichText = PageBlock;
//  pageBlockEmbedPost#f259a80b url:string webpage_id:long author_photo_id:long author:string date:int blocks:Vector<PageBlock> caption:PageCaption = PageBlock;
//
func (m *TLPageBlockEmbedPost) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockEmbedPost
	return m.Data2
}

//  pageBlockCollage#8b31c4f items:Vector<PageBlock> caption:RichText = PageBlock;
//  pageBlockCollage#65a0fa4d items:Vector<PageBlock> caption:PageCaption = PageBlock;
//
func (m *TLPageBlockCollage) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockCollage
	return m.Data2
}

//  pageBlockSlideshow#130c8963 items:Vector<PageBlock> caption:RichText = PageBlock;
//  pageBlockSlideshow#31f9590 items:Vector<PageBlock> caption:PageCaption = PageBlock;
//
func (m *TLPageBlockSlideshow) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockSlideshow
	return m.Data2
}

//  pageBlockChannel#ef1751b5 channel:Chat = PageBlock;
//
func (m *TLPageBlockChannel) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockChannel
	return m.Data2
}

//  pageBlockAudio#31b81a7f audio_id:long caption:RichText = PageBlock;
//  pageBlockAudio#804361ea audio_id:long caption:PageCaption = PageBlock;
//
func (m *TLPageBlockAudio) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockAudio
	return m.Data2
}

//  pageBlockKicker#1e148390 text:RichText = PageBlock;
//
func (m *TLPageBlockKicker) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockKicker
	return m.Data2
}

//  pageBlockTable#bf4dea82 flags:# bordered:flags.0?true striped:flags.1?true title:RichText rows:Vector<PageTableRow> = PageBlock;
//
func (m *TLPageBlockTable) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockTable
	return m.Data2
}

//  pageBlockOrderedList#9a8ae1e1 items:Vector<PageListOrderedItem> = PageBlock;
//
func (m *TLPageBlockOrderedList) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockOrderedList
	return m.Data2
}

//  pageBlockDetails#76768bed flags:# open:flags.0?true blocks:Vector<PageBlock> title:RichText = PageBlock;
//
func (m *TLPageBlockDetails) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockDetails
	return m.Data2
}

//  pageBlockRelatedArticles#16115a96 title:RichText articles:Vector<PageRelatedArticle> = PageBlock;
//
func (m *TLPageBlockRelatedArticles) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockRelatedArticles
	return m.Data2
}

//  pageBlockMap#a44f3ef6 geo:GeoPoint zoom:int w:int h:int caption:PageCaption = PageBlock;
//
func (m *TLPageBlockMap) To_PageBlock() *PageBlock {
	m.Data2.Cmd = Cmd_PageBlockMap
	return m.Data2
}

//  + TL_PageCaption
//

//  pageCaption#6f747657 text:RichText credit:RichText = PageCaption;
//
func (m *PageCaption) To_PageCaption() *TLPageCaption {
	return &TLPageCaption{
		Data2: m,
	}
}

//  pageCaption#6f747657 text:RichText credit:RichText = PageCaption;
//
func (m *TLPageCaption) To_PageCaption() *PageCaption {
	m.Data2.Cmd = Cmd_PageCaption
	return m.Data2
}

//  + TL_PageListItemText
//  + TL_PageListItemBlocks
//

//  pageListItemText#b92fb6cd text:RichText = PageListItem;
//
func (m *PageListItem) To_PageListItemText() *TLPageListItemText {
	return &TLPageListItemText{
		Data2: m,
	}
}

//  pageListItemBlocks#25e073fc blocks:Vector<PageBlock> = PageListItem;
//
func (m *PageListItem) To_PageListItemBlocks() *TLPageListItemBlocks {
	return &TLPageListItemBlocks{
		Data2: m,
	}
}

//  pageListItemText#b92fb6cd text:RichText = PageListItem;
//
func (m *TLPageListItemText) To_PageListItem() *PageListItem {
	m.Data2.Cmd = Cmd_PageListItemText
	return m.Data2
}

//  pageListItemBlocks#25e073fc blocks:Vector<PageBlock> = PageListItem;
//
func (m *TLPageListItemBlocks) To_PageListItem() *PageListItem {
	m.Data2.Cmd = Cmd_PageListItemBlocks
	return m.Data2
}

//  + TL_PageListOrderedItemText
//  + TL_PageListOrderedItemBlocks
//

//  pageListOrderedItemText#5e068047 num:string text:RichText = PageListOrderedItem;
//
func (m *PageListOrderedItem) To_PageListOrderedItemText() *TLPageListOrderedItemText {
	return &TLPageListOrderedItemText{
		Data2: m,
	}
}

//  pageListOrderedItemBlocks#98dd8936 num:string blocks:Vector<PageBlock> = PageListOrderedItem;
//
func (m *PageListOrderedItem) To_PageListOrderedItemBlocks() *TLPageListOrderedItemBlocks {
	return &TLPageListOrderedItemBlocks{
		Data2: m,
	}
}

//  pageListOrderedItemText#5e068047 num:string text:RichText = PageListOrderedItem;
//
func (m *TLPageListOrderedItemText) To_PageListOrderedItem() *PageListOrderedItem {
	m.Data2.Cmd = Cmd_PageListOrderedItemText
	return m.Data2
}

//  pageListOrderedItemBlocks#98dd8936 num:string blocks:Vector<PageBlock> = PageListOrderedItem;
//
func (m *TLPageListOrderedItemBlocks) To_PageListOrderedItem() *PageListOrderedItem {
	m.Data2.Cmd = Cmd_PageListOrderedItemBlocks
	return m.Data2
}

//  + TL_PageRelatedArticle
//

//  pageRelatedArticle#f186f93c flags:# url:string webpage_id:long title:flags.0?string description:flags.1?string photo_id:flags.2?long = PageRelatedArticle;
//  pageRelatedArticle#b390dc08 flags:# url:string webpage_id:long title:flags.0?string description:flags.1?string photo_id:flags.2?long author:flags.3?string published_date:flags.4?int = PageRelatedArticle;
//
func (m *PageRelatedArticle) To_PageRelatedArticle() *TLPageRelatedArticle {
	return &TLPageRelatedArticle{
		Data2: m,
	}
}

//  pageRelatedArticle#f186f93c flags:# url:string webpage_id:long title:flags.0?string description:flags.1?string photo_id:flags.2?long = PageRelatedArticle;
//  pageRelatedArticle#b390dc08 flags:# url:string webpage_id:long title:flags.0?string description:flags.1?string photo_id:flags.2?long author:flags.3?string published_date:flags.4?int = PageRelatedArticle;
//
func (m *TLPageRelatedArticle) To_PageRelatedArticle() *PageRelatedArticle {
	m.Data2.Cmd = Cmd_PageRelatedArticle
	return m.Data2
}

//  + TL_PageTableCell
//

//  pageTableCell#34566b6a flags:# header:flags.0?true align_center:flags.3?true align_right:flags.4?true valign_middle:flags.5?true valign_bottom:flags.6?true text:flags.7?RichText colspan:flags.1?int rowspan:flags.2?int = PageTableCell;
//
func (m *PageTableCell) To_PageTableCell() *TLPageTableCell {
	return &TLPageTableCell{
		Data2: m,
	}
}

//  pageTableCell#34566b6a flags:# header:flags.0?true align_center:flags.3?true align_right:flags.4?true valign_middle:flags.5?true valign_bottom:flags.6?true text:flags.7?RichText colspan:flags.1?int rowspan:flags.2?int = PageTableCell;
//
func (m *TLPageTableCell) To_PageTableCell() *PageTableCell {
	m.Data2.Cmd = Cmd_PageTableCell
	return m.Data2
}

//  + TL_PageTableRow
//

//  pageTableRow#e0c0c5e5 cells:Vector<PageTableCell> = PageTableRow;
//
func (m *PageTableRow) To_PageTableRow() *TLPageTableRow {
	return &TLPageTableRow{
		Data2: m,
	}
}

//  pageTableRow#e0c0c5e5 cells:Vector<PageTableCell> = PageTableRow;
//
func (m *TLPageTableRow) To_PageTableRow() *PageTableRow {
	m.Data2.Cmd = Cmd_PageTableRow
	return m.Data2
}

//  + TL_PasswordKdfAlgoUnknown
//  + TL_PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow
//

//  passwordKdfAlgoUnknown#d45ab096 = PasswordKdfAlgo;
//
func (m *PasswordKdfAlgo) To_PasswordKdfAlgoUnknown() *TLPasswordKdfAlgoUnknown {
	return &TLPasswordKdfAlgoUnknown{
		Data2: m,
	}
}

//  passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a salt1:bytes salt2:bytes g:int p:bytes = PasswordKdfAlgo;
//
func (m *PasswordKdfAlgo) To_PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow() *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow {
	return &TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow{
		Data2: m,
	}
}

//  passwordKdfAlgoUnknown#d45ab096 = PasswordKdfAlgo;
//
func (m *TLPasswordKdfAlgoUnknown) To_PasswordKdfAlgo() *PasswordKdfAlgo {
	m.Data2.Cmd = Cmd_PasswordKdfAlgoUnknown
	return m.Data2
}

//  passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow#3a912d4a salt1:bytes salt2:bytes g:int p:bytes = PasswordKdfAlgo;
//
func (m *TLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow) To_PasswordKdfAlgo() *PasswordKdfAlgo {
	m.Data2.Cmd = Cmd_PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow
	return m.Data2
}

//  + TL_PaymentCharge
//

//  paymentCharge#ea02c27e id:string provider_charge_id:string = PaymentCharge;
//
func (m *PaymentCharge) To_PaymentCharge() *TLPaymentCharge {
	return &TLPaymentCharge{
		Data2: m,
	}
}

//  paymentCharge#ea02c27e id:string provider_charge_id:string = PaymentCharge;
//
func (m *TLPaymentCharge) To_PaymentCharge() *PaymentCharge {
	m.Data2.Cmd = Cmd_PaymentCharge
	return m.Data2
}

//  + TL_PaymentRequestedInfo
//

//  paymentRequestedInfo#909c3f94 flags:# name:flags.0?string phone:flags.1?string email:flags.2?string shipping_address:flags.3?PostAddress = PaymentRequestedInfo;
//
func (m *PaymentRequestedInfo) To_PaymentRequestedInfo() *TLPaymentRequestedInfo {
	return &TLPaymentRequestedInfo{
		Data2: m,
	}
}

//  paymentRequestedInfo#909c3f94 flags:# name:flags.0?string phone:flags.1?string email:flags.2?string shipping_address:flags.3?PostAddress = PaymentRequestedInfo;
//
func (m *TLPaymentRequestedInfo) To_PaymentRequestedInfo() *PaymentRequestedInfo {
	m.Data2.Cmd = Cmd_PaymentRequestedInfo
	return m.Data2
}

//  + TL_PaymentSavedCredentialsCard
//

//  paymentSavedCredentialsCard#cdc27a1f id:string title:string = PaymentSavedCredentials;
//
func (m *PaymentSavedCredentials) To_PaymentSavedCredentialsCard() *TLPaymentSavedCredentialsCard {
	return &TLPaymentSavedCredentialsCard{
		Data2: m,
	}
}

//  paymentSavedCredentialsCard#cdc27a1f id:string title:string = PaymentSavedCredentials;
//
func (m *TLPaymentSavedCredentialsCard) To_PaymentSavedCredentials() *PaymentSavedCredentials {
	m.Data2.Cmd = Cmd_PaymentSavedCredentialsCard
	return m.Data2
}

//  + TL_PaymentsBankCardData
//

//  payments.bankCardData#3e24e573 title:string open_urls:Vector<BankCardOpenUrl> = payments.BankCardData;
//
func (m *Payments_BankCardData) To_PaymentsBankCardData() *TLPaymentsBankCardData {
	return &TLPaymentsBankCardData{
		Data2: m,
	}
}

//  payments.bankCardData#3e24e573 title:string open_urls:Vector<BankCardOpenUrl> = payments.BankCardData;
//
func (m *TLPaymentsBankCardData) To_Payments_BankCardData() *Payments_BankCardData {
	m.Data2.Cmd = Cmd_PaymentsBankCardData
	return m.Data2
}

//  + TL_PaymentsPaymentForm
//

//  payments.paymentForm#3f56aea3 flags:# can_save_credentials:flags.2?true password_missing:flags.3?true bot_id:int invoice:Invoice provider_id:int url:string native_provider:flags.4?string native_params:flags.4?DataJSON saved_info:flags.0?PaymentRequestedInfo saved_credentials:flags.1?PaymentSavedCredentials users:Vector<User> = payments.PaymentForm;
//
func (m *Payments_PaymentForm) To_PaymentsPaymentForm() *TLPaymentsPaymentForm {
	return &TLPaymentsPaymentForm{
		Data2: m,
	}
}

//  payments.paymentForm#3f56aea3 flags:# can_save_credentials:flags.2?true password_missing:flags.3?true bot_id:int invoice:Invoice provider_id:int url:string native_provider:flags.4?string native_params:flags.4?DataJSON saved_info:flags.0?PaymentRequestedInfo saved_credentials:flags.1?PaymentSavedCredentials users:Vector<User> = payments.PaymentForm;
//
func (m *TLPaymentsPaymentForm) To_Payments_PaymentForm() *Payments_PaymentForm {
	m.Data2.Cmd = Cmd_PaymentsPaymentForm
	return m.Data2
}

//  + TL_PaymentsPaymentReceipt
//

//  payments.paymentReceipt#500911e1 flags:# date:int bot_id:int invoice:Invoice provider_id:int info:flags.0?PaymentRequestedInfo shipping:flags.1?ShippingOption currency:string total_amount:long credentials_title:string users:Vector<User> = payments.PaymentReceipt;
//
func (m *Payments_PaymentReceipt) To_PaymentsPaymentReceipt() *TLPaymentsPaymentReceipt {
	return &TLPaymentsPaymentReceipt{
		Data2: m,
	}
}

//  payments.paymentReceipt#500911e1 flags:# date:int bot_id:int invoice:Invoice provider_id:int info:flags.0?PaymentRequestedInfo shipping:flags.1?ShippingOption currency:string total_amount:long credentials_title:string users:Vector<User> = payments.PaymentReceipt;
//
func (m *TLPaymentsPaymentReceipt) To_Payments_PaymentReceipt() *Payments_PaymentReceipt {
	m.Data2.Cmd = Cmd_PaymentsPaymentReceipt
	return m.Data2
}

//  + TL_PaymentsPaymentResult
//  + TL_PaymentsPaymentVerficationNeeded
//  + TL_PaymentsPaymentVerificationNeeded
//

//  payments.paymentResult#4e5f810d updates:Updates = payments.PaymentResult;
//
func (m *Payments_PaymentResult) To_PaymentsPaymentResult() *TLPaymentsPaymentResult {
	return &TLPaymentsPaymentResult{
		Data2: m,
	}
}

//  payments.paymentVerficationNeeded#6b56b921 url:string = payments.PaymentResult;
//
func (m *Payments_PaymentResult) To_PaymentsPaymentVerficationNeeded() *TLPaymentsPaymentVerficationNeeded {
	return &TLPaymentsPaymentVerficationNeeded{
		Data2: m,
	}
}

//  payments.paymentVerificationNeeded#d8411139 url:string = payments.PaymentResult;
//
func (m *Payments_PaymentResult) To_PaymentsPaymentVerificationNeeded() *TLPaymentsPaymentVerificationNeeded {
	return &TLPaymentsPaymentVerificationNeeded{
		Data2: m,
	}
}

//  payments.paymentResult#4e5f810d updates:Updates = payments.PaymentResult;
//
func (m *TLPaymentsPaymentResult) To_Payments_PaymentResult() *Payments_PaymentResult {
	m.Data2.Cmd = Cmd_PaymentsPaymentResult
	return m.Data2
}

//  payments.paymentVerficationNeeded#6b56b921 url:string = payments.PaymentResult;
//
func (m *TLPaymentsPaymentVerficationNeeded) To_Payments_PaymentResult() *Payments_PaymentResult {
	m.Data2.Cmd = Cmd_PaymentsPaymentVerficationNeeded
	return m.Data2
}

//  payments.paymentVerificationNeeded#d8411139 url:string = payments.PaymentResult;
//
func (m *TLPaymentsPaymentVerificationNeeded) To_Payments_PaymentResult() *Payments_PaymentResult {
	m.Data2.Cmd = Cmd_PaymentsPaymentVerificationNeeded
	return m.Data2
}

//  + TL_PaymentsSavedInfo
//

//  payments.savedInfo#fb8fe43c flags:# has_saved_credentials:flags.1?true saved_info:flags.0?PaymentRequestedInfo = payments.SavedInfo;
//
func (m *Payments_SavedInfo) To_PaymentsSavedInfo() *TLPaymentsSavedInfo {
	return &TLPaymentsSavedInfo{
		Data2: m,
	}
}

//  payments.savedInfo#fb8fe43c flags:# has_saved_credentials:flags.1?true saved_info:flags.0?PaymentRequestedInfo = payments.SavedInfo;
//
func (m *TLPaymentsSavedInfo) To_Payments_SavedInfo() *Payments_SavedInfo {
	m.Data2.Cmd = Cmd_PaymentsSavedInfo
	return m.Data2
}

//  + TL_PaymentsValidatedRequestedInfo
//

//  payments.validatedRequestedInfo#d1451883 flags:# id:flags.0?string shipping_options:flags.1?Vector<ShippingOption> = payments.ValidatedRequestedInfo;
//
func (m *Payments_ValidatedRequestedInfo) To_PaymentsValidatedRequestedInfo() *TLPaymentsValidatedRequestedInfo {
	return &TLPaymentsValidatedRequestedInfo{
		Data2: m,
	}
}

//  payments.validatedRequestedInfo#d1451883 flags:# id:flags.0?string shipping_options:flags.1?Vector<ShippingOption> = payments.ValidatedRequestedInfo;
//
func (m *TLPaymentsValidatedRequestedInfo) To_Payments_ValidatedRequestedInfo() *Payments_ValidatedRequestedInfo {
	m.Data2.Cmd = Cmd_PaymentsValidatedRequestedInfo
	return m.Data2
}

//  + TL_PeerUser
//  + TL_PeerChat
//  + TL_PeerChannel
//

//  peerUser#9db1bc6d user_id:int = Peer;
//
func (m *Peer) To_PeerUser() *TLPeerUser {
	return &TLPeerUser{
		Data2: m,
	}
}

//  peerChat#bad0e5bb chat_id:int = Peer;
//
func (m *Peer) To_PeerChat() *TLPeerChat {
	return &TLPeerChat{
		Data2: m,
	}
}

//  peerChannel#bddde532 channel_id:int = Peer;
//
func (m *Peer) To_PeerChannel() *TLPeerChannel {
	return &TLPeerChannel{
		Data2: m,
	}
}

//  peerUser#9db1bc6d user_id:int = Peer;
//
func (m *TLPeerUser) To_Peer() *Peer {
	m.Data2.Cmd = Cmd_PeerUser
	return m.Data2
}

//  peerChat#bad0e5bb chat_id:int = Peer;
//
func (m *TLPeerChat) To_Peer() *Peer {
	m.Data2.Cmd = Cmd_PeerChat
	return m.Data2
}

//  peerChannel#bddde532 channel_id:int = Peer;
//
func (m *TLPeerChannel) To_Peer() *Peer {
	m.Data2.Cmd = Cmd_PeerChannel
	return m.Data2
}

//  + TL_PeerBlocked
//

//  peerBlocked#e8fd8014 peer_id:Peer date:int = PeerBlocked;
//
func (m *PeerBlocked) To_PeerBlocked() *TLPeerBlocked {
	return &TLPeerBlocked{
		Data2: m,
	}
}

//  peerBlocked#e8fd8014 peer_id:Peer date:int = PeerBlocked;
//
func (m *TLPeerBlocked) To_PeerBlocked() *PeerBlocked {
	m.Data2.Cmd = Cmd_PeerBlocked
	return m.Data2
}

//  + TL_PeerLocated
//  + TL_PeerSelfLocated
//

//  peerLocated#ca461b5d peer:Peer expires:int distance:int = PeerLocated;
//
func (m *PeerLocated) To_PeerLocated() *TLPeerLocated {
	return &TLPeerLocated{
		Data2: m,
	}
}

//  peerSelfLocated#f8ec284b expires:int = PeerLocated;
//
func (m *PeerLocated) To_PeerSelfLocated() *TLPeerSelfLocated {
	return &TLPeerSelfLocated{
		Data2: m,
	}
}

//  peerLocated#ca461b5d peer:Peer expires:int distance:int = PeerLocated;
//
func (m *TLPeerLocated) To_PeerLocated() *PeerLocated {
	m.Data2.Cmd = Cmd_PeerLocated
	return m.Data2
}

//  peerSelfLocated#f8ec284b expires:int = PeerLocated;
//
func (m *TLPeerSelfLocated) To_PeerLocated() *PeerLocated {
	m.Data2.Cmd = Cmd_PeerSelfLocated
	return m.Data2
}

//  + TL_PeerNotifyEventsEmpty
//  + TL_PeerNotifyEventsAll
//

//  peerNotifyEventsEmpty#add53cb3 = PeerNotifyEvents;
//
func (m *PeerNotifyEvents) To_PeerNotifyEventsEmpty() *TLPeerNotifyEventsEmpty {
	return &TLPeerNotifyEventsEmpty{
		Data2: m,
	}
}

//  peerNotifyEventsAll#6d1ded88 = PeerNotifyEvents;
//
func (m *PeerNotifyEvents) To_PeerNotifyEventsAll() *TLPeerNotifyEventsAll {
	return &TLPeerNotifyEventsAll{
		Data2: m,
	}
}

//  peerNotifyEventsEmpty#add53cb3 = PeerNotifyEvents;
//
func (m *TLPeerNotifyEventsEmpty) To_PeerNotifyEvents() *PeerNotifyEvents {
	m.Data2.Cmd = Cmd_PeerNotifyEventsEmpty
	return m.Data2
}

//  peerNotifyEventsAll#6d1ded88 = PeerNotifyEvents;
//
func (m *TLPeerNotifyEventsAll) To_PeerNotifyEvents() *PeerNotifyEvents {
	m.Data2.Cmd = Cmd_PeerNotifyEventsAll
	return m.Data2
}

//  + TL_PeerNotifySettingsEmpty
//  + TL_PeerNotifySettings
//

//  peerNotifySettingsEmpty#70a68512 = PeerNotifySettings;
//
func (m *PeerNotifySettings) To_PeerNotifySettingsEmpty() *TLPeerNotifySettingsEmpty {
	return &TLPeerNotifySettingsEmpty{
		Data2: m,
	}
}

//  peerNotifySettings#9acda4c0 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = PeerNotifySettings;
//  peerNotifySettings#af509d20 flags:# show_previews:flags.0?Bool silent:flags.1?Bool mute_until:flags.2?int sound:flags.3?string = PeerNotifySettings;
//
func (m *PeerNotifySettings) To_PeerNotifySettings() *TLPeerNotifySettings {
	return &TLPeerNotifySettings{
		Data2: m,
	}
}

//  peerNotifySettingsEmpty#70a68512 = PeerNotifySettings;
//
func (m *TLPeerNotifySettingsEmpty) To_PeerNotifySettings() *PeerNotifySettings {
	m.Data2.Cmd = Cmd_PeerNotifySettingsEmpty
	return m.Data2
}

//  peerNotifySettings#9acda4c0 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = PeerNotifySettings;
//  peerNotifySettings#af509d20 flags:# show_previews:flags.0?Bool silent:flags.1?Bool mute_until:flags.2?int sound:flags.3?string = PeerNotifySettings;
//
func (m *TLPeerNotifySettings) To_PeerNotifySettings() *PeerNotifySettings {
	m.Data2.Cmd = Cmd_PeerNotifySettings
	return m.Data2
}

//  + TL_PeerSettings
//

//  peerSettings#818426cd flags:# report_spam:flags.0?true = PeerSettings;
//  peerSettings#733f2961 flags:# report_spam:flags.0?true add_contact:flags.1?true block_contact:flags.2?true share_contact:flags.3?true need_contacts_exception:flags.4?true report_geo:flags.5?true autoarchived:flags.7?true geo_distance:flags.6?int = PeerSettings;
//
func (m *PeerSettings) To_PeerSettings() *TLPeerSettings {
	return &TLPeerSettings{
		Data2: m,
	}
}

//  peerSettings#818426cd flags:# report_spam:flags.0?true = PeerSettings;
//  peerSettings#733f2961 flags:# report_spam:flags.0?true add_contact:flags.1?true block_contact:flags.2?true share_contact:flags.3?true need_contacts_exception:flags.4?true report_geo:flags.5?true autoarchived:flags.7?true geo_distance:flags.6?int = PeerSettings;
//
func (m *TLPeerSettings) To_PeerSettings() *PeerSettings {
	m.Data2.Cmd = Cmd_PeerSettings
	return m.Data2
}

//  + TL_PhoneCallEmpty
//  + TL_PhoneCallRequested
//  + TL_PhoneCallAccepted
//  + TL_PhoneCall
//  + TL_PhoneCallWaiting
//  + TL_PhoneCallDiscarded
//

//  phoneCallEmpty#5366c915 id:long = PhoneCall;
//
func (m *PhoneCall) To_PhoneCallEmpty() *TLPhoneCallEmpty {
	return &TLPhoneCallEmpty{
		Data2: m,
	}
}

//  phoneCallRequested#83761ce4 id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;
//  phoneCallRequested#87eabb53 flags:# video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;
//
func (m *PhoneCall) To_PhoneCallRequested() *TLPhoneCallRequested {
	return &TLPhoneCallRequested{
		Data2: m,
	}
}

//  phoneCallAccepted#6d003d3f id:long access_hash:long date:int admin_id:int participant_id:int g_b:bytes protocol:PhoneCallProtocol = PhoneCall;
//  phoneCallAccepted#997c454a flags:# video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int g_b:bytes protocol:PhoneCallProtocol = PhoneCall;
//
func (m *PhoneCall) To_PhoneCallAccepted() *TLPhoneCallAccepted {
	return &TLPhoneCallAccepted{
		Data2: m,
	}
}

//  phoneCall#ffe6ab67 id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;
//  phoneCall#e6f9ddf3 flags:# p2p_allowed:flags.5?true id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;
//  phoneCall#8742ae7f flags:# p2p_allowed:flags.5?true video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connections:Vector<PhoneConnection> start_date:int = PhoneCall;
//
func (m *PhoneCall) To_PhoneCall() *TLPhoneCall {
	return &TLPhoneCall{
		Data2: m,
	}
}

//  phoneCallWaiting#1b8f4ad1 flags:# video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int protocol:PhoneCallProtocol receive_date:flags.0?int = PhoneCall;
//
func (m *PhoneCall) To_PhoneCallWaiting() *TLPhoneCallWaiting {
	return &TLPhoneCallWaiting{
		Data2: m,
	}
}

//  phoneCallDiscarded#50ca4de1 flags:# need_rating:flags.2?true need_debug:flags.3?true video:flags.6?true id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = PhoneCall;
//
func (m *PhoneCall) To_PhoneCallDiscarded() *TLPhoneCallDiscarded {
	return &TLPhoneCallDiscarded{
		Data2: m,
	}
}

//  phoneCallEmpty#5366c915 id:long = PhoneCall;
//
func (m *TLPhoneCallEmpty) To_PhoneCall() *PhoneCall {
	m.Data2.Cmd = Cmd_PhoneCallEmpty
	return m.Data2
}

//  phoneCallRequested#83761ce4 id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;
//  phoneCallRequested#87eabb53 flags:# video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;
//
func (m *TLPhoneCallRequested) To_PhoneCall() *PhoneCall {
	m.Data2.Cmd = Cmd_PhoneCallRequested
	return m.Data2
}

//  phoneCallAccepted#6d003d3f id:long access_hash:long date:int admin_id:int participant_id:int g_b:bytes protocol:PhoneCallProtocol = PhoneCall;
//  phoneCallAccepted#997c454a flags:# video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int g_b:bytes protocol:PhoneCallProtocol = PhoneCall;
//
func (m *TLPhoneCallAccepted) To_PhoneCall() *PhoneCall {
	m.Data2.Cmd = Cmd_PhoneCallAccepted
	return m.Data2
}

//  phoneCall#ffe6ab67 id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;
//  phoneCall#e6f9ddf3 flags:# p2p_allowed:flags.5?true id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;
//  phoneCall#8742ae7f flags:# p2p_allowed:flags.5?true video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connections:Vector<PhoneConnection> start_date:int = PhoneCall;
//
func (m *TLPhoneCall) To_PhoneCall() *PhoneCall {
	m.Data2.Cmd = Cmd_PhoneCall
	return m.Data2
}

//  phoneCallWaiting#1b8f4ad1 flags:# video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int protocol:PhoneCallProtocol receive_date:flags.0?int = PhoneCall;
//
func (m *TLPhoneCallWaiting) To_PhoneCall() *PhoneCall {
	m.Data2.Cmd = Cmd_PhoneCallWaiting
	return m.Data2
}

//  phoneCallDiscarded#50ca4de1 flags:# need_rating:flags.2?true need_debug:flags.3?true video:flags.6?true id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = PhoneCall;
//
func (m *TLPhoneCallDiscarded) To_PhoneCall() *PhoneCall {
	m.Data2.Cmd = Cmd_PhoneCallDiscarded
	return m.Data2
}

//  + TL_PhoneCallDiscardReasonMissed
//  + TL_PhoneCallDiscardReasonDisconnect
//  + TL_PhoneCallDiscardReasonHangup
//  + TL_PhoneCallDiscardReasonBusy
//

//  phoneCallDiscardReasonMissed#85e42301 = PhoneCallDiscardReason;
//
func (m *PhoneCallDiscardReason) To_PhoneCallDiscardReasonMissed() *TLPhoneCallDiscardReasonMissed {
	return &TLPhoneCallDiscardReasonMissed{
		Data2: m,
	}
}

//  phoneCallDiscardReasonDisconnect#e095c1a0 = PhoneCallDiscardReason;
//
func (m *PhoneCallDiscardReason) To_PhoneCallDiscardReasonDisconnect() *TLPhoneCallDiscardReasonDisconnect {
	return &TLPhoneCallDiscardReasonDisconnect{
		Data2: m,
	}
}

//  phoneCallDiscardReasonHangup#57adc690 = PhoneCallDiscardReason;
//
func (m *PhoneCallDiscardReason) To_PhoneCallDiscardReasonHangup() *TLPhoneCallDiscardReasonHangup {
	return &TLPhoneCallDiscardReasonHangup{
		Data2: m,
	}
}

//  phoneCallDiscardReasonBusy#faf7e8c9 = PhoneCallDiscardReason;
//
func (m *PhoneCallDiscardReason) To_PhoneCallDiscardReasonBusy() *TLPhoneCallDiscardReasonBusy {
	return &TLPhoneCallDiscardReasonBusy{
		Data2: m,
	}
}

//  phoneCallDiscardReasonMissed#85e42301 = PhoneCallDiscardReason;
//
func (m *TLPhoneCallDiscardReasonMissed) To_PhoneCallDiscardReason() *PhoneCallDiscardReason {
	m.Data2.Cmd = Cmd_PhoneCallDiscardReasonMissed
	return m.Data2
}

//  phoneCallDiscardReasonDisconnect#e095c1a0 = PhoneCallDiscardReason;
//
func (m *TLPhoneCallDiscardReasonDisconnect) To_PhoneCallDiscardReason() *PhoneCallDiscardReason {
	m.Data2.Cmd = Cmd_PhoneCallDiscardReasonDisconnect
	return m.Data2
}

//  phoneCallDiscardReasonHangup#57adc690 = PhoneCallDiscardReason;
//
func (m *TLPhoneCallDiscardReasonHangup) To_PhoneCallDiscardReason() *PhoneCallDiscardReason {
	m.Data2.Cmd = Cmd_PhoneCallDiscardReasonHangup
	return m.Data2
}

//  phoneCallDiscardReasonBusy#faf7e8c9 = PhoneCallDiscardReason;
//
func (m *TLPhoneCallDiscardReasonBusy) To_PhoneCallDiscardReason() *PhoneCallDiscardReason {
	m.Data2.Cmd = Cmd_PhoneCallDiscardReasonBusy
	return m.Data2
}

//  + TL_PhoneCallProtocol
//

//  phoneCallProtocol#a2bb35cb flags:# udp_p2p:flags.0?true udp_reflector:flags.1?true min_layer:int max_layer:int = PhoneCallProtocol;
//  phoneCallProtocol#fc878fc8 flags:# udp_p2p:flags.0?true udp_reflector:flags.1?true min_layer:int max_layer:int library_versions:Vector<string> = PhoneCallProtocol;
//
func (m *PhoneCallProtocol) To_PhoneCallProtocol() *TLPhoneCallProtocol {
	return &TLPhoneCallProtocol{
		Data2: m,
	}
}

//  phoneCallProtocol#a2bb35cb flags:# udp_p2p:flags.0?true udp_reflector:flags.1?true min_layer:int max_layer:int = PhoneCallProtocol;
//  phoneCallProtocol#fc878fc8 flags:# udp_p2p:flags.0?true udp_reflector:flags.1?true min_layer:int max_layer:int library_versions:Vector<string> = PhoneCallProtocol;
//
func (m *TLPhoneCallProtocol) To_PhoneCallProtocol() *PhoneCallProtocol {
	m.Data2.Cmd = Cmd_PhoneCallProtocol
	return m.Data2
}

//  + TL_PhoneConnection
//  + TL_PhoneConnectionWebrtc
//

//  phoneConnection#9d4c17c0 id:long ip:string ipv6:string port:int peer_tag:bytes = PhoneConnection;
//
func (m *PhoneConnection) To_PhoneConnection() *TLPhoneConnection {
	return &TLPhoneConnection{
		Data2: m,
	}
}

//  phoneConnectionWebrtc#635fe375 flags:# turn:flags.0?true stun:flags.1?true id:long ip:string ipv6:string port:int username:string password:string = PhoneConnection;
//
func (m *PhoneConnection) To_PhoneConnectionWebrtc() *TLPhoneConnectionWebrtc {
	return &TLPhoneConnectionWebrtc{
		Data2: m,
	}
}

//  phoneConnection#9d4c17c0 id:long ip:string ipv6:string port:int peer_tag:bytes = PhoneConnection;
//
func (m *TLPhoneConnection) To_PhoneConnection() *PhoneConnection {
	m.Data2.Cmd = Cmd_PhoneConnection
	return m.Data2
}

//  phoneConnectionWebrtc#635fe375 flags:# turn:flags.0?true stun:flags.1?true id:long ip:string ipv6:string port:int username:string password:string = PhoneConnection;
//
func (m *TLPhoneConnectionWebrtc) To_PhoneConnection() *PhoneConnection {
	m.Data2.Cmd = Cmd_PhoneConnectionWebrtc
	return m.Data2
}

//  + TL_PhoneGroupCall
//

//  phone.groupCall#66ab0bfc call:GroupCall participants:Vector<GroupCallParticipant> participants_next_offset:string users:Vector<User> = phone.GroupCall;
//
func (m *Phone_GroupCall) To_PhoneGroupCall() *TLPhoneGroupCall {
	return &TLPhoneGroupCall{
		Data2: m,
	}
}

//  phone.groupCall#66ab0bfc call:GroupCall participants:Vector<GroupCallParticipant> participants_next_offset:string users:Vector<User> = phone.GroupCall;
//
func (m *TLPhoneGroupCall) To_Phone_GroupCall() *Phone_GroupCall {
	m.Data2.Cmd = Cmd_PhoneGroupCall
	return m.Data2
}

//  + TL_PhoneGroupParticipants
//

//  phone.groupParticipants#9cfeb92d count:int participants:Vector<GroupCallParticipant> next_offset:string users:Vector<User> version:int = phone.GroupParticipants;
//
func (m *Phone_GroupParticipants) To_PhoneGroupParticipants() *TLPhoneGroupParticipants {
	return &TLPhoneGroupParticipants{
		Data2: m,
	}
}

//  phone.groupParticipants#9cfeb92d count:int participants:Vector<GroupCallParticipant> next_offset:string users:Vector<User> version:int = phone.GroupParticipants;
//
func (m *TLPhoneGroupParticipants) To_Phone_GroupParticipants() *Phone_GroupParticipants {
	m.Data2.Cmd = Cmd_PhoneGroupParticipants
	return m.Data2
}

//  + TL_PhonePhoneCall
//

//  phone.phoneCall#ec82e140 phone_call:PhoneCall users:Vector<User> = phone.PhoneCall;
//
func (m *Phone_PhoneCall) To_PhonePhoneCall() *TLPhonePhoneCall {
	return &TLPhonePhoneCall{
		Data2: m,
	}
}

//  phone.phoneCall#ec82e140 phone_call:PhoneCall users:Vector<User> = phone.PhoneCall;
//
func (m *TLPhonePhoneCall) To_Phone_PhoneCall() *Phone_PhoneCall {
	m.Data2.Cmd = Cmd_PhonePhoneCall
	return m.Data2
}

//  + TL_PhotoEmpty
//  + TL_Photo
//

//  photoEmpty#2331b22d id:long = Photo;
//
func (m *Photo) To_PhotoEmpty() *TLPhotoEmpty {
	return &TLPhotoEmpty{
		Data2: m,
	}
}

//  photo#9288dd29 flags:# has_stickers:flags.0?true id:long access_hash:long date:int sizes:Vector<PhotoSize> = Photo;
//  photo#cded42fe id:long access_hash:long date:int sizes:Vector<PhotoSize> = Photo;
//  photo#9c477dd8 flags:# has_stickers:flags.0?true id:long access_hash:long file_reference:bytes date:int sizes:Vector<PhotoSize> = Photo;
//  photo#d07504a5 flags:# has_stickers:flags.0?true id:long access_hash:long file_reference:bytes date:int sizes:Vector<PhotoSize> dc_id:int = Photo;
//  photo#fb197a65 flags:# has_stickers:flags.0?true id:long access_hash:long file_reference:bytes date:int sizes:Vector<PhotoSize> video_sizes:flags.1?Vector<VideoSize> dc_id:int = Photo;
//
func (m *Photo) To_Photo() *TLPhoto {
	return &TLPhoto{
		Data2: m,
	}
}

//  photoEmpty#2331b22d id:long = Photo;
//
func (m *TLPhotoEmpty) To_Photo() *Photo {
	m.Data2.Cmd = Cmd_PhotoEmpty
	return m.Data2
}

//  photo#9288dd29 flags:# has_stickers:flags.0?true id:long access_hash:long date:int sizes:Vector<PhotoSize> = Photo;
//  photo#cded42fe id:long access_hash:long date:int sizes:Vector<PhotoSize> = Photo;
//  photo#9c477dd8 flags:# has_stickers:flags.0?true id:long access_hash:long file_reference:bytes date:int sizes:Vector<PhotoSize> = Photo;
//  photo#d07504a5 flags:# has_stickers:flags.0?true id:long access_hash:long file_reference:bytes date:int sizes:Vector<PhotoSize> dc_id:int = Photo;
//  photo#fb197a65 flags:# has_stickers:flags.0?true id:long access_hash:long file_reference:bytes date:int sizes:Vector<PhotoSize> video_sizes:flags.1?Vector<VideoSize> dc_id:int = Photo;
//
func (m *TLPhoto) To_Photo() *Photo {
	m.Data2.Cmd = Cmd_Photo
	return m.Data2
}

//  + TL_PhotoSizeEmpty
//  + TL_PhotoSize
//  + TL_PhotoCachedSize
//  + TL_PhotoStrippedSize
//  + TL_PhotoSizeProgressive
//  + TL_PhotoPathSize
//

//  photoSizeEmpty#e17e23c type:string = PhotoSize;
//
func (m *PhotoSize) To_PhotoSizeEmpty() *TLPhotoSizeEmpty {
	return &TLPhotoSizeEmpty{
		Data2: m,
	}
}

//  photoSize#77bfb61b type:string location:FileLocation w:int h:int size:int = PhotoSize;
//
func (m *PhotoSize) To_PhotoSize() *TLPhotoSize {
	return &TLPhotoSize{
		Data2: m,
	}
}

//  photoCachedSize#e9a734fa type:string location:FileLocation w:int h:int bytes:bytes = PhotoSize;
//
func (m *PhotoSize) To_PhotoCachedSize() *TLPhotoCachedSize {
	return &TLPhotoCachedSize{
		Data2: m,
	}
}

//  photoStrippedSize#e0b0bc2e type:string bytes:bytes = PhotoSize;
//
func (m *PhotoSize) To_PhotoStrippedSize() *TLPhotoStrippedSize {
	return &TLPhotoStrippedSize{
		Data2: m,
	}
}

//  photoSizeProgressive#5aa86a51 type:string location:FileLocation w:int h:int sizes:Vector<int> = PhotoSize;
//
func (m *PhotoSize) To_PhotoSizeProgressive() *TLPhotoSizeProgressive {
	return &TLPhotoSizeProgressive{
		Data2: m,
	}
}

//  photoPathSize#d8214d41 type:string bytes:bytes = PhotoSize;
//
func (m *PhotoSize) To_PhotoPathSize() *TLPhotoPathSize {
	return &TLPhotoPathSize{
		Data2: m,
	}
}

//  photoSizeEmpty#e17e23c type:string = PhotoSize;
//
func (m *TLPhotoSizeEmpty) To_PhotoSize() *PhotoSize {
	m.Data2.Cmd = Cmd_PhotoSizeEmpty
	return m.Data2
}

//  photoSize#77bfb61b type:string location:FileLocation w:int h:int size:int = PhotoSize;
//
func (m *TLPhotoSize) To_PhotoSize() *PhotoSize {
	m.Data2.Cmd = Cmd_PhotoSize
	return m.Data2
}

//  photoCachedSize#e9a734fa type:string location:FileLocation w:int h:int bytes:bytes = PhotoSize;
//
func (m *TLPhotoCachedSize) To_PhotoSize() *PhotoSize {
	m.Data2.Cmd = Cmd_PhotoCachedSize
	return m.Data2
}

//  photoStrippedSize#e0b0bc2e type:string bytes:bytes = PhotoSize;
//
func (m *TLPhotoStrippedSize) To_PhotoSize() *PhotoSize {
	m.Data2.Cmd = Cmd_PhotoStrippedSize
	return m.Data2
}

//  photoSizeProgressive#5aa86a51 type:string location:FileLocation w:int h:int sizes:Vector<int> = PhotoSize;
//
func (m *TLPhotoSizeProgressive) To_PhotoSize() *PhotoSize {
	m.Data2.Cmd = Cmd_PhotoSizeProgressive
	return m.Data2
}

//  photoPathSize#d8214d41 type:string bytes:bytes = PhotoSize;
//
func (m *TLPhotoPathSize) To_PhotoSize() *PhotoSize {
	m.Data2.Cmd = Cmd_PhotoPathSize
	return m.Data2
}

//  + TL_PhotosPhoto
//

//  photos.photo#20212ca8 photo:Photo users:Vector<User> = photos.Photo;
//
func (m *Photos_Photo) To_PhotosPhoto() *TLPhotosPhoto {
	return &TLPhotosPhoto{
		Data2: m,
	}
}

//  photos.photo#20212ca8 photo:Photo users:Vector<User> = photos.Photo;
//
func (m *TLPhotosPhoto) To_Photos_Photo() *Photos_Photo {
	m.Data2.Cmd = Cmd_PhotosPhoto
	return m.Data2
}

//  + TL_PhotosPhotos
//  + TL_PhotosPhotosSlice
//

//  photos.photos#8dca6aa5 photos:Vector<Photo> users:Vector<User> = photos.Photos;
//
func (m *Photos_Photos) To_PhotosPhotos() *TLPhotosPhotos {
	return &TLPhotosPhotos{
		Data2: m,
	}
}

//  photos.photosSlice#15051f54 count:int photos:Vector<Photo> users:Vector<User> = photos.Photos;
//
func (m *Photos_Photos) To_PhotosPhotosSlice() *TLPhotosPhotosSlice {
	return &TLPhotosPhotosSlice{
		Data2: m,
	}
}

//  photos.photos#8dca6aa5 photos:Vector<Photo> users:Vector<User> = photos.Photos;
//
func (m *TLPhotosPhotos) To_Photos_Photos() *Photos_Photos {
	m.Data2.Cmd = Cmd_PhotosPhotos
	return m.Data2
}

//  photos.photosSlice#15051f54 count:int photos:Vector<Photo> users:Vector<User> = photos.Photos;
//
func (m *TLPhotosPhotosSlice) To_Photos_Photos() *Photos_Photos {
	m.Data2.Cmd = Cmd_PhotosPhotosSlice
	return m.Data2
}

//  + TL_Poll
//

//  poll#d5529d06 id:long flags:# closed:flags.0?true question:string answers:Vector<PollAnswer> = Poll;
//  poll#86e18161 id:long flags:# closed:flags.0?true public_voters:flags.1?true multiple_choice:flags.2?true quiz:flags.3?true question:string answers:Vector<PollAnswer> close_period:flags.4?int close_date:flags.5?int = Poll;
//
func (m *Poll) To_Poll() *TLPoll {
	return &TLPoll{
		Data2: m,
	}
}

//  poll#d5529d06 id:long flags:# closed:flags.0?true question:string answers:Vector<PollAnswer> = Poll;
//  poll#86e18161 id:long flags:# closed:flags.0?true public_voters:flags.1?true multiple_choice:flags.2?true quiz:flags.3?true question:string answers:Vector<PollAnswer> close_period:flags.4?int close_date:flags.5?int = Poll;
//
func (m *TLPoll) To_Poll() *Poll {
	m.Data2.Cmd = Cmd_Poll
	return m.Data2
}

//  + TL_PollAnswer
//

//  pollAnswer#6ca9c2e9 text:string option:bytes = PollAnswer;
//
func (m *PollAnswer) To_PollAnswer() *TLPollAnswer {
	return &TLPollAnswer{
		Data2: m,
	}
}

//  pollAnswer#6ca9c2e9 text:string option:bytes = PollAnswer;
//
func (m *TLPollAnswer) To_PollAnswer() *PollAnswer {
	m.Data2.Cmd = Cmd_PollAnswer
	return m.Data2
}

//  + TL_PollAnswerVoters
//

//  pollAnswerVoters#3b6ddad2 flags:# chosen:flags.0?true option:bytes voters:int = PollAnswerVoters;
//
func (m *PollAnswerVoters) To_PollAnswerVoters() *TLPollAnswerVoters {
	return &TLPollAnswerVoters{
		Data2: m,
	}
}

//  pollAnswerVoters#3b6ddad2 flags:# chosen:flags.0?true option:bytes voters:int = PollAnswerVoters;
//
func (m *TLPollAnswerVoters) To_PollAnswerVoters() *PollAnswerVoters {
	m.Data2.Cmd = Cmd_PollAnswerVoters
	return m.Data2
}

//  + TL_PollResults
//

//  pollResults#5755785a flags:# min:flags.0?true results:flags.1?Vector<PollAnswerVoters> total_voters:flags.2?int = PollResults;
//  pollResults#c87024a2 flags:# min:flags.0?true results:flags.1?Vector<PollAnswerVoters> total_voters:flags.2?int recent_voters:flags.3?Vector<int> = PollResults;
//  pollResults#badcc1a3 flags:# min:flags.0?true results:flags.1?Vector<PollAnswerVoters> total_voters:flags.2?int recent_voters:flags.3?Vector<int> solution:flags.4?string solution_entities:flags.4?Vector<MessageEntity> = PollResults;
//
func (m *PollResults) To_PollResults() *TLPollResults {
	return &TLPollResults{
		Data2: m,
	}
}

//  pollResults#5755785a flags:# min:flags.0?true results:flags.1?Vector<PollAnswerVoters> total_voters:flags.2?int = PollResults;
//  pollResults#c87024a2 flags:# min:flags.0?true results:flags.1?Vector<PollAnswerVoters> total_voters:flags.2?int recent_voters:flags.3?Vector<int> = PollResults;
//  pollResults#badcc1a3 flags:# min:flags.0?true results:flags.1?Vector<PollAnswerVoters> total_voters:flags.2?int recent_voters:flags.3?Vector<int> solution:flags.4?string solution_entities:flags.4?Vector<MessageEntity> = PollResults;
//
func (m *TLPollResults) To_PollResults() *PollResults {
	m.Data2.Cmd = Cmd_PollResults
	return m.Data2
}

//  + TL_Pong
//

//  pong#347773c5 msg_id:long ping_id:long = Pong;
//
func (m *Pong) To_Pong() *TLPong {
	return &TLPong{
		Data2: m,
	}
}

//  pong#347773c5 msg_id:long ping_id:long = Pong;
//
func (m *TLPong) To_Pong() *Pong {
	m.Data2.Cmd = Cmd_Pong
	return m.Data2
}

//  + TL_PopularContact
//

//  popularContact#5ce14175 client_id:long importers:int = PopularContact;
//
func (m *PopularContact) To_PopularContact() *TLPopularContact {
	return &TLPopularContact{
		Data2: m,
	}
}

//  popularContact#5ce14175 client_id:long importers:int = PopularContact;
//
func (m *TLPopularContact) To_PopularContact() *PopularContact {
	m.Data2.Cmd = Cmd_PopularContact
	return m.Data2
}

//  + TL_PostAddress
//

//  postAddress#1e8caaeb street_line1:string street_line2:string city:string state:string country_iso2:string post_code:string = PostAddress;
//
func (m *PostAddress) To_PostAddress() *TLPostAddress {
	return &TLPostAddress{
		Data2: m,
	}
}

//  postAddress#1e8caaeb street_line1:string street_line2:string city:string state:string country_iso2:string post_code:string = PostAddress;
//
func (m *TLPostAddress) To_PostAddress() *PostAddress {
	m.Data2.Cmd = Cmd_PostAddress
	return m.Data2
}

//  + TL_PrivacyKeyStatusTimestamp
//  + TL_PrivacyKeyChatInvite
//  + TL_PrivacyKeyPhoneCall
//  + TL_PrivacyKeyPhoneP2P
//  + TL_PrivacyKeyForwards
//  + TL_PrivacyKeyProfilePhoto
//  + TL_PrivacyKeyPhoneNumber
//  + TL_PrivacyKeyAddedByPhone
//

//  privacyKeyStatusTimestamp#bc2eab30 = PrivacyKey;
//
func (m *PrivacyKey) To_PrivacyKeyStatusTimestamp() *TLPrivacyKeyStatusTimestamp {
	return &TLPrivacyKeyStatusTimestamp{
		Data2: m,
	}
}

//  privacyKeyChatInvite#500e6dfa = PrivacyKey;
//
func (m *PrivacyKey) To_PrivacyKeyChatInvite() *TLPrivacyKeyChatInvite {
	return &TLPrivacyKeyChatInvite{
		Data2: m,
	}
}

//  privacyKeyPhoneCall#3d662b7b = PrivacyKey;
//
func (m *PrivacyKey) To_PrivacyKeyPhoneCall() *TLPrivacyKeyPhoneCall {
	return &TLPrivacyKeyPhoneCall{
		Data2: m,
	}
}

//  privacyKeyPhoneP2P#39491cc8 = PrivacyKey;
//
func (m *PrivacyKey) To_PrivacyKeyPhoneP2P() *TLPrivacyKeyPhoneP2P {
	return &TLPrivacyKeyPhoneP2P{
		Data2: m,
	}
}

//  privacyKeyForwards#69ec56a3 = PrivacyKey;
//
func (m *PrivacyKey) To_PrivacyKeyForwards() *TLPrivacyKeyForwards {
	return &TLPrivacyKeyForwards{
		Data2: m,
	}
}

//  privacyKeyProfilePhoto#96151fed = PrivacyKey;
//
func (m *PrivacyKey) To_PrivacyKeyProfilePhoto() *TLPrivacyKeyProfilePhoto {
	return &TLPrivacyKeyProfilePhoto{
		Data2: m,
	}
}

//  privacyKeyPhoneNumber#d19ae46d = PrivacyKey;
//
func (m *PrivacyKey) To_PrivacyKeyPhoneNumber() *TLPrivacyKeyPhoneNumber {
	return &TLPrivacyKeyPhoneNumber{
		Data2: m,
	}
}

//  privacyKeyAddedByPhone#42ffd42b = PrivacyKey;
//
func (m *PrivacyKey) To_PrivacyKeyAddedByPhone() *TLPrivacyKeyAddedByPhone {
	return &TLPrivacyKeyAddedByPhone{
		Data2: m,
	}
}

//  privacyKeyStatusTimestamp#bc2eab30 = PrivacyKey;
//
func (m *TLPrivacyKeyStatusTimestamp) To_PrivacyKey() *PrivacyKey {
	m.Data2.Cmd = Cmd_PrivacyKeyStatusTimestamp
	return m.Data2
}

//  privacyKeyChatInvite#500e6dfa = PrivacyKey;
//
func (m *TLPrivacyKeyChatInvite) To_PrivacyKey() *PrivacyKey {
	m.Data2.Cmd = Cmd_PrivacyKeyChatInvite
	return m.Data2
}

//  privacyKeyPhoneCall#3d662b7b = PrivacyKey;
//
func (m *TLPrivacyKeyPhoneCall) To_PrivacyKey() *PrivacyKey {
	m.Data2.Cmd = Cmd_PrivacyKeyPhoneCall
	return m.Data2
}

//  privacyKeyPhoneP2P#39491cc8 = PrivacyKey;
//
func (m *TLPrivacyKeyPhoneP2P) To_PrivacyKey() *PrivacyKey {
	m.Data2.Cmd = Cmd_PrivacyKeyPhoneP2P
	return m.Data2
}

//  privacyKeyForwards#69ec56a3 = PrivacyKey;
//
func (m *TLPrivacyKeyForwards) To_PrivacyKey() *PrivacyKey {
	m.Data2.Cmd = Cmd_PrivacyKeyForwards
	return m.Data2
}

//  privacyKeyProfilePhoto#96151fed = PrivacyKey;
//
func (m *TLPrivacyKeyProfilePhoto) To_PrivacyKey() *PrivacyKey {
	m.Data2.Cmd = Cmd_PrivacyKeyProfilePhoto
	return m.Data2
}

//  privacyKeyPhoneNumber#d19ae46d = PrivacyKey;
//
func (m *TLPrivacyKeyPhoneNumber) To_PrivacyKey() *PrivacyKey {
	m.Data2.Cmd = Cmd_PrivacyKeyPhoneNumber
	return m.Data2
}

//  privacyKeyAddedByPhone#42ffd42b = PrivacyKey;
//
func (m *TLPrivacyKeyAddedByPhone) To_PrivacyKey() *PrivacyKey {
	m.Data2.Cmd = Cmd_PrivacyKeyAddedByPhone
	return m.Data2
}

//  + TL_PrivacyValueAllowContacts
//  + TL_PrivacyValueAllowAll
//  + TL_PrivacyValueAllowUsers
//  + TL_PrivacyValueDisallowContacts
//  + TL_PrivacyValueDisallowAll
//  + TL_PrivacyValueDisallowUsers
//  + TL_PrivacyValueAllowChatParticipants
//  + TL_PrivacyValueDisallowChatParticipants
//

//  privacyValueAllowContacts#fffe1bac = PrivacyRule;
//
func (m *PrivacyRule) To_PrivacyValueAllowContacts() *TLPrivacyValueAllowContacts {
	return &TLPrivacyValueAllowContacts{
		Data2: m,
	}
}

//  privacyValueAllowAll#65427b82 = PrivacyRule;
//
func (m *PrivacyRule) To_PrivacyValueAllowAll() *TLPrivacyValueAllowAll {
	return &TLPrivacyValueAllowAll{
		Data2: m,
	}
}

//  privacyValueAllowUsers#4d5bbe0c users:Vector<int> = PrivacyRule;
//
func (m *PrivacyRule) To_PrivacyValueAllowUsers() *TLPrivacyValueAllowUsers {
	return &TLPrivacyValueAllowUsers{
		Data2: m,
	}
}

//  privacyValueDisallowContacts#f888fa1a = PrivacyRule;
//
func (m *PrivacyRule) To_PrivacyValueDisallowContacts() *TLPrivacyValueDisallowContacts {
	return &TLPrivacyValueDisallowContacts{
		Data2: m,
	}
}

//  privacyValueDisallowAll#8b73e763 = PrivacyRule;
//
func (m *PrivacyRule) To_PrivacyValueDisallowAll() *TLPrivacyValueDisallowAll {
	return &TLPrivacyValueDisallowAll{
		Data2: m,
	}
}

//  privacyValueDisallowUsers#c7f49b7 users:Vector<int> = PrivacyRule;
//
func (m *PrivacyRule) To_PrivacyValueDisallowUsers() *TLPrivacyValueDisallowUsers {
	return &TLPrivacyValueDisallowUsers{
		Data2: m,
	}
}

//  privacyValueAllowChatParticipants#18be796b chats:Vector<int> = PrivacyRule;
//
func (m *PrivacyRule) To_PrivacyValueAllowChatParticipants() *TLPrivacyValueAllowChatParticipants {
	return &TLPrivacyValueAllowChatParticipants{
		Data2: m,
	}
}

//  privacyValueDisallowChatParticipants#acae0690 chats:Vector<int> = PrivacyRule;
//
func (m *PrivacyRule) To_PrivacyValueDisallowChatParticipants() *TLPrivacyValueDisallowChatParticipants {
	return &TLPrivacyValueDisallowChatParticipants{
		Data2: m,
	}
}

//  privacyValueAllowContacts#fffe1bac = PrivacyRule;
//
func (m *TLPrivacyValueAllowContacts) To_PrivacyRule() *PrivacyRule {
	m.Data2.Cmd = Cmd_PrivacyValueAllowContacts
	return m.Data2
}

//  privacyValueAllowAll#65427b82 = PrivacyRule;
//
func (m *TLPrivacyValueAllowAll) To_PrivacyRule() *PrivacyRule {
	m.Data2.Cmd = Cmd_PrivacyValueAllowAll
	return m.Data2
}

//  privacyValueAllowUsers#4d5bbe0c users:Vector<int> = PrivacyRule;
//
func (m *TLPrivacyValueAllowUsers) To_PrivacyRule() *PrivacyRule {
	m.Data2.Cmd = Cmd_PrivacyValueAllowUsers
	return m.Data2
}

//  privacyValueDisallowContacts#f888fa1a = PrivacyRule;
//
func (m *TLPrivacyValueDisallowContacts) To_PrivacyRule() *PrivacyRule {
	m.Data2.Cmd = Cmd_PrivacyValueDisallowContacts
	return m.Data2
}

//  privacyValueDisallowAll#8b73e763 = PrivacyRule;
//
func (m *TLPrivacyValueDisallowAll) To_PrivacyRule() *PrivacyRule {
	m.Data2.Cmd = Cmd_PrivacyValueDisallowAll
	return m.Data2
}

//  privacyValueDisallowUsers#c7f49b7 users:Vector<int> = PrivacyRule;
//
func (m *TLPrivacyValueDisallowUsers) To_PrivacyRule() *PrivacyRule {
	m.Data2.Cmd = Cmd_PrivacyValueDisallowUsers
	return m.Data2
}

//  privacyValueAllowChatParticipants#18be796b chats:Vector<int> = PrivacyRule;
//
func (m *TLPrivacyValueAllowChatParticipants) To_PrivacyRule() *PrivacyRule {
	m.Data2.Cmd = Cmd_PrivacyValueAllowChatParticipants
	return m.Data2
}

//  privacyValueDisallowChatParticipants#acae0690 chats:Vector<int> = PrivacyRule;
//
func (m *TLPrivacyValueDisallowChatParticipants) To_PrivacyRule() *PrivacyRule {
	m.Data2.Cmd = Cmd_PrivacyValueDisallowChatParticipants
	return m.Data2
}

//  + TL_ReceivedNotifyMessage
//

//  receivedNotifyMessage#a384b779 id:int flags:int = ReceivedNotifyMessage;
//
func (m *ReceivedNotifyMessage) To_ReceivedNotifyMessage() *TLReceivedNotifyMessage {
	return &TLReceivedNotifyMessage{
		Data2: m,
	}
}

//  receivedNotifyMessage#a384b779 id:int flags:int = ReceivedNotifyMessage;
//
func (m *TLReceivedNotifyMessage) To_ReceivedNotifyMessage() *ReceivedNotifyMessage {
	m.Data2.Cmd = Cmd_ReceivedNotifyMessage
	return m.Data2
}

//  + TL_RecentMeUrlUnknown
//  + TL_RecentMeUrlUser
//  + TL_RecentMeUrlChat
//  + TL_RecentMeUrlChatInvite
//  + TL_RecentMeUrlStickerSet
//

//  recentMeUrlUnknown#46e1d13d url:string = RecentMeUrl;
//
func (m *RecentMeUrl) To_RecentMeUrlUnknown() *TLRecentMeUrlUnknown {
	return &TLRecentMeUrlUnknown{
		Data2: m,
	}
}

//  recentMeUrlUser#8dbc3336 url:string user_id:int = RecentMeUrl;
//
func (m *RecentMeUrl) To_RecentMeUrlUser() *TLRecentMeUrlUser {
	return &TLRecentMeUrlUser{
		Data2: m,
	}
}

//  recentMeUrlChat#a01b22f9 url:string chat_id:int = RecentMeUrl;
//
func (m *RecentMeUrl) To_RecentMeUrlChat() *TLRecentMeUrlChat {
	return &TLRecentMeUrlChat{
		Data2: m,
	}
}

//  recentMeUrlChatInvite#eb49081d url:string chat_invite:ChatInvite = RecentMeUrl;
//
func (m *RecentMeUrl) To_RecentMeUrlChatInvite() *TLRecentMeUrlChatInvite {
	return &TLRecentMeUrlChatInvite{
		Data2: m,
	}
}

//  recentMeUrlStickerSet#bc0a57dc url:string set:StickerSetCovered = RecentMeUrl;
//
func (m *RecentMeUrl) To_RecentMeUrlStickerSet() *TLRecentMeUrlStickerSet {
	return &TLRecentMeUrlStickerSet{
		Data2: m,
	}
}

//  recentMeUrlUnknown#46e1d13d url:string = RecentMeUrl;
//
func (m *TLRecentMeUrlUnknown) To_RecentMeUrl() *RecentMeUrl {
	m.Data2.Cmd = Cmd_RecentMeUrlUnknown
	return m.Data2
}

//  recentMeUrlUser#8dbc3336 url:string user_id:int = RecentMeUrl;
//
func (m *TLRecentMeUrlUser) To_RecentMeUrl() *RecentMeUrl {
	m.Data2.Cmd = Cmd_RecentMeUrlUser
	return m.Data2
}

//  recentMeUrlChat#a01b22f9 url:string chat_id:int = RecentMeUrl;
//
func (m *TLRecentMeUrlChat) To_RecentMeUrl() *RecentMeUrl {
	m.Data2.Cmd = Cmd_RecentMeUrlChat
	return m.Data2
}

//  recentMeUrlChatInvite#eb49081d url:string chat_invite:ChatInvite = RecentMeUrl;
//
func (m *TLRecentMeUrlChatInvite) To_RecentMeUrl() *RecentMeUrl {
	m.Data2.Cmd = Cmd_RecentMeUrlChatInvite
	return m.Data2
}

//  recentMeUrlStickerSet#bc0a57dc url:string set:StickerSetCovered = RecentMeUrl;
//
func (m *TLRecentMeUrlStickerSet) To_RecentMeUrl() *RecentMeUrl {
	m.Data2.Cmd = Cmd_RecentMeUrlStickerSet
	return m.Data2
}

//  + TL_ReplyKeyboardHide
//  + TL_ReplyKeyboardForceReply
//  + TL_ReplyKeyboardMarkup
//  + TL_ReplyInlineMarkup
//

//  replyKeyboardHide#a03e5b85 flags:# selective:flags.2?true = ReplyMarkup;
//
func (m *ReplyMarkup) To_ReplyKeyboardHide() *TLReplyKeyboardHide {
	return &TLReplyKeyboardHide{
		Data2: m,
	}
}

//  replyKeyboardForceReply#f4108aa0 flags:# single_use:flags.1?true selective:flags.2?true = ReplyMarkup;
//
func (m *ReplyMarkup) To_ReplyKeyboardForceReply() *TLReplyKeyboardForceReply {
	return &TLReplyKeyboardForceReply{
		Data2: m,
	}
}

//  replyKeyboardMarkup#3502758c flags:# resize:flags.0?true single_use:flags.1?true selective:flags.2?true rows:Vector<KeyboardButtonRow> = ReplyMarkup;
//
func (m *ReplyMarkup) To_ReplyKeyboardMarkup() *TLReplyKeyboardMarkup {
	return &TLReplyKeyboardMarkup{
		Data2: m,
	}
}

//  replyInlineMarkup#48a30254 rows:Vector<KeyboardButtonRow> = ReplyMarkup;
//
func (m *ReplyMarkup) To_ReplyInlineMarkup() *TLReplyInlineMarkup {
	return &TLReplyInlineMarkup{
		Data2: m,
	}
}

//  replyKeyboardHide#a03e5b85 flags:# selective:flags.2?true = ReplyMarkup;
//
func (m *TLReplyKeyboardHide) To_ReplyMarkup() *ReplyMarkup {
	m.Data2.Cmd = Cmd_ReplyKeyboardHide
	return m.Data2
}

//  replyKeyboardForceReply#f4108aa0 flags:# single_use:flags.1?true selective:flags.2?true = ReplyMarkup;
//
func (m *TLReplyKeyboardForceReply) To_ReplyMarkup() *ReplyMarkup {
	m.Data2.Cmd = Cmd_ReplyKeyboardForceReply
	return m.Data2
}

//  replyKeyboardMarkup#3502758c flags:# resize:flags.0?true single_use:flags.1?true selective:flags.2?true rows:Vector<KeyboardButtonRow> = ReplyMarkup;
//
func (m *TLReplyKeyboardMarkup) To_ReplyMarkup() *ReplyMarkup {
	m.Data2.Cmd = Cmd_ReplyKeyboardMarkup
	return m.Data2
}

//  replyInlineMarkup#48a30254 rows:Vector<KeyboardButtonRow> = ReplyMarkup;
//
func (m *TLReplyInlineMarkup) To_ReplyMarkup() *ReplyMarkup {
	m.Data2.Cmd = Cmd_ReplyInlineMarkup
	return m.Data2
}

//  + TL_InputReportReasonSpam
//  + TL_InputReportReasonViolence
//  + TL_InputReportReasonPornography
//  + TL_InputReportReasonOther
//  + TL_InputReportReasonCopyright
//  + TL_InputReportReasonChildAbuse
//  + TL_InputReportReasonGeoIrrelevant
//

//  inputReportReasonSpam#58dbcab8 = ReportReason;
//
func (m *ReportReason) To_InputReportReasonSpam() *TLInputReportReasonSpam {
	return &TLInputReportReasonSpam{
		Data2: m,
	}
}

//  inputReportReasonViolence#1e22c78d = ReportReason;
//
func (m *ReportReason) To_InputReportReasonViolence() *TLInputReportReasonViolence {
	return &TLInputReportReasonViolence{
		Data2: m,
	}
}

//  inputReportReasonPornography#2e59d922 = ReportReason;
//
func (m *ReportReason) To_InputReportReasonPornography() *TLInputReportReasonPornography {
	return &TLInputReportReasonPornography{
		Data2: m,
	}
}

//  inputReportReasonOther#e1746d0a text:string = ReportReason;
//
func (m *ReportReason) To_InputReportReasonOther() *TLInputReportReasonOther {
	return &TLInputReportReasonOther{
		Data2: m,
	}
}

//  inputReportReasonCopyright#9b89f93a = ReportReason;
//
func (m *ReportReason) To_InputReportReasonCopyright() *TLInputReportReasonCopyright {
	return &TLInputReportReasonCopyright{
		Data2: m,
	}
}

//  inputReportReasonChildAbuse#adf44ee3 = ReportReason;
//
func (m *ReportReason) To_InputReportReasonChildAbuse() *TLInputReportReasonChildAbuse {
	return &TLInputReportReasonChildAbuse{
		Data2: m,
	}
}

//  inputReportReasonGeoIrrelevant#dbd4feed = ReportReason;
//
func (m *ReportReason) To_InputReportReasonGeoIrrelevant() *TLInputReportReasonGeoIrrelevant {
	return &TLInputReportReasonGeoIrrelevant{
		Data2: m,
	}
}

//  inputReportReasonSpam#58dbcab8 = ReportReason;
//
func (m *TLInputReportReasonSpam) To_ReportReason() *ReportReason {
	m.Data2.Cmd = Cmd_InputReportReasonSpam
	return m.Data2
}

//  inputReportReasonViolence#1e22c78d = ReportReason;
//
func (m *TLInputReportReasonViolence) To_ReportReason() *ReportReason {
	m.Data2.Cmd = Cmd_InputReportReasonViolence
	return m.Data2
}

//  inputReportReasonPornography#2e59d922 = ReportReason;
//
func (m *TLInputReportReasonPornography) To_ReportReason() *ReportReason {
	m.Data2.Cmd = Cmd_InputReportReasonPornography
	return m.Data2
}

//  inputReportReasonOther#e1746d0a text:string = ReportReason;
//
func (m *TLInputReportReasonOther) To_ReportReason() *ReportReason {
	m.Data2.Cmd = Cmd_InputReportReasonOther
	return m.Data2
}

//  inputReportReasonCopyright#9b89f93a = ReportReason;
//
func (m *TLInputReportReasonCopyright) To_ReportReason() *ReportReason {
	m.Data2.Cmd = Cmd_InputReportReasonCopyright
	return m.Data2
}

//  inputReportReasonChildAbuse#adf44ee3 = ReportReason;
//
func (m *TLInputReportReasonChildAbuse) To_ReportReason() *ReportReason {
	m.Data2.Cmd = Cmd_InputReportReasonChildAbuse
	return m.Data2
}

//  inputReportReasonGeoIrrelevant#dbd4feed = ReportReason;
//
func (m *TLInputReportReasonGeoIrrelevant) To_ReportReason() *ReportReason {
	m.Data2.Cmd = Cmd_InputReportReasonGeoIrrelevant
	return m.Data2
}

//  + TL_ResPQ
//

//  resPQ#05162463 nonce:int128 server_nonce:int128 pq:string server_public_key_fingerprints:Vector<long> = ResPQ;
//
func (m *ResPQ) To_ResPQ() *TLResPQ {
	return &TLResPQ{
		Data2: m,
	}
}

//  resPQ#05162463 nonce:int128 server_nonce:int128 pq:string server_public_key_fingerprints:Vector<long> = ResPQ;
//
func (m *TLResPQ) To_ResPQ() *ResPQ {
	m.Data2.Cmd = Cmd_ResPQ
	return m.Data2
}

//  + TL_RestrictionReason
//

//  restrictionReason#d072acb4 platform:string reason:string text:string = RestrictionReason;
//
func (m *RestrictionReason) To_RestrictionReason() *TLRestrictionReason {
	return &TLRestrictionReason{
		Data2: m,
	}
}

//  restrictionReason#d072acb4 platform:string reason:string text:string = RestrictionReason;
//
func (m *TLRestrictionReason) To_RestrictionReason() *RestrictionReason {
	m.Data2.Cmd = Cmd_RestrictionReason
	return m.Data2
}

//  + TL_TextEmpty
//  + TL_TextPlain
//  + TL_TextBold
//  + TL_TextItalic
//  + TL_TextUnderline
//  + TL_TextStrike
//  + TL_TextFixed
//  + TL_TextUrl
//  + TL_TextEmail
//  + TL_TextConcat
//  + TL_TextSubscript
//  + TL_TextSuperscript
//  + TL_TextMarked
//  + TL_TextPhone
//  + TL_TextImage
//  + TL_TextAnchor
//

//  textEmpty#dc3d824f = RichText;
//
func (m *RichText) To_TextEmpty() *TLTextEmpty {
	return &TLTextEmpty{
		Data2: m,
	}
}

//  textPlain#744694e0 text:string = RichText;
//
func (m *RichText) To_TextPlain() *TLTextPlain {
	return &TLTextPlain{
		Data2: m,
	}
}

//  textBold#6724abc4 text:RichText = RichText;
//
func (m *RichText) To_TextBold() *TLTextBold {
	return &TLTextBold{
		Data2: m,
	}
}

//  textItalic#d912a59c text:RichText = RichText;
//
func (m *RichText) To_TextItalic() *TLTextItalic {
	return &TLTextItalic{
		Data2: m,
	}
}

//  textUnderline#c12622c4 text:RichText = RichText;
//
func (m *RichText) To_TextUnderline() *TLTextUnderline {
	return &TLTextUnderline{
		Data2: m,
	}
}

//  textStrike#9bf8bb95 text:RichText = RichText;
//
func (m *RichText) To_TextStrike() *TLTextStrike {
	return &TLTextStrike{
		Data2: m,
	}
}

//  textFixed#6c3f19b9 text:RichText = RichText;
//
func (m *RichText) To_TextFixed() *TLTextFixed {
	return &TLTextFixed{
		Data2: m,
	}
}

//  textUrl#3c2884c1 text:RichText url:string webpage_id:long = RichText;
//
func (m *RichText) To_TextUrl() *TLTextUrl {
	return &TLTextUrl{
		Data2: m,
	}
}

//  textEmail#de5a0dd6 text:RichText email:string = RichText;
//
func (m *RichText) To_TextEmail() *TLTextEmail {
	return &TLTextEmail{
		Data2: m,
	}
}

//  textConcat#7e6260d7 texts:Vector<RichText> = RichText;
//
func (m *RichText) To_TextConcat() *TLTextConcat {
	return &TLTextConcat{
		Data2: m,
	}
}

//  textSubscript#ed6a8504 text:RichText = RichText;
//
func (m *RichText) To_TextSubscript() *TLTextSubscript {
	return &TLTextSubscript{
		Data2: m,
	}
}

//  textSuperscript#c7fb5e01 text:RichText = RichText;
//
func (m *RichText) To_TextSuperscript() *TLTextSuperscript {
	return &TLTextSuperscript{
		Data2: m,
	}
}

//  textMarked#34b8621 text:RichText = RichText;
//
func (m *RichText) To_TextMarked() *TLTextMarked {
	return &TLTextMarked{
		Data2: m,
	}
}

//  textPhone#1ccb966a text:RichText phone:string = RichText;
//
func (m *RichText) To_TextPhone() *TLTextPhone {
	return &TLTextPhone{
		Data2: m,
	}
}

//  textImage#81ccf4f document_id:long w:int h:int = RichText;
//
func (m *RichText) To_TextImage() *TLTextImage {
	return &TLTextImage{
		Data2: m,
	}
}

//  textAnchor#35553762 text:RichText name:string = RichText;
//
func (m *RichText) To_TextAnchor() *TLTextAnchor {
	return &TLTextAnchor{
		Data2: m,
	}
}

//  textEmpty#dc3d824f = RichText;
//
func (m *TLTextEmpty) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextEmpty
	return m.Data2
}

//  textPlain#744694e0 text:string = RichText;
//
func (m *TLTextPlain) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextPlain
	return m.Data2
}

//  textBold#6724abc4 text:RichText = RichText;
//
func (m *TLTextBold) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextBold
	return m.Data2
}

//  textItalic#d912a59c text:RichText = RichText;
//
func (m *TLTextItalic) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextItalic
	return m.Data2
}

//  textUnderline#c12622c4 text:RichText = RichText;
//
func (m *TLTextUnderline) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextUnderline
	return m.Data2
}

//  textStrike#9bf8bb95 text:RichText = RichText;
//
func (m *TLTextStrike) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextStrike
	return m.Data2
}

//  textFixed#6c3f19b9 text:RichText = RichText;
//
func (m *TLTextFixed) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextFixed
	return m.Data2
}

//  textUrl#3c2884c1 text:RichText url:string webpage_id:long = RichText;
//
func (m *TLTextUrl) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextUrl
	return m.Data2
}

//  textEmail#de5a0dd6 text:RichText email:string = RichText;
//
func (m *TLTextEmail) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextEmail
	return m.Data2
}

//  textConcat#7e6260d7 texts:Vector<RichText> = RichText;
//
func (m *TLTextConcat) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextConcat
	return m.Data2
}

//  textSubscript#ed6a8504 text:RichText = RichText;
//
func (m *TLTextSubscript) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextSubscript
	return m.Data2
}

//  textSuperscript#c7fb5e01 text:RichText = RichText;
//
func (m *TLTextSuperscript) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextSuperscript
	return m.Data2
}

//  textMarked#34b8621 text:RichText = RichText;
//
func (m *TLTextMarked) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextMarked
	return m.Data2
}

//  textPhone#1ccb966a text:RichText phone:string = RichText;
//
func (m *TLTextPhone) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextPhone
	return m.Data2
}

//  textImage#81ccf4f document_id:long w:int h:int = RichText;
//
func (m *TLTextImage) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextImage
	return m.Data2
}

//  textAnchor#35553762 text:RichText name:string = RichText;
//
func (m *TLTextAnchor) To_RichText() *RichText {
	m.Data2.Cmd = Cmd_TextAnchor
	return m.Data2
}

//  + TL_RpcAnswerUnknown
//  + TL_RpcAnswerDroppedRunning
//  + TL_RpcAnswerDropped
//

//  rpc_answer_unknown#5e2ad36e = RpcDropAnswer;
//
func (m *RpcDropAnswer) To_RpcAnswerUnknown() *TLRpcAnswerUnknown {
	return &TLRpcAnswerUnknown{
		Data2: m,
	}
}

//  rpc_answer_dropped_running#cd78e586 = RpcDropAnswer;
//
func (m *RpcDropAnswer) To_RpcAnswerDroppedRunning() *TLRpcAnswerDroppedRunning {
	return &TLRpcAnswerDroppedRunning{
		Data2: m,
	}
}

//  rpc_answer_dropped#a43ad8b7 msg_id:long seq_no:int bytes:int = RpcDropAnswer;
//
func (m *RpcDropAnswer) To_RpcAnswerDropped() *TLRpcAnswerDropped {
	return &TLRpcAnswerDropped{
		Data2: m,
	}
}

//  rpc_answer_unknown#5e2ad36e = RpcDropAnswer;
//
func (m *TLRpcAnswerUnknown) To_RpcDropAnswer() *RpcDropAnswer {
	m.Data2.Cmd = Cmd_RpcAnswerUnknown
	return m.Data2
}

//  rpc_answer_dropped_running#cd78e586 = RpcDropAnswer;
//
func (m *TLRpcAnswerDroppedRunning) To_RpcDropAnswer() *RpcDropAnswer {
	m.Data2.Cmd = Cmd_RpcAnswerDroppedRunning
	return m.Data2
}

//  rpc_answer_dropped#a43ad8b7 msg_id:long seq_no:int bytes:int = RpcDropAnswer;
//
func (m *TLRpcAnswerDropped) To_RpcDropAnswer() *RpcDropAnswer {
	m.Data2.Cmd = Cmd_RpcAnswerDropped
	return m.Data2
}

//  + TL_RpcError
//

//  rpc_error#2144ca19 error_code:int error_message:string = RpcError;
//
func (m *RpcError) To_RpcError() *TLRpcError {
	return &TLRpcError{
		Data2: m,
	}
}

//  rpc_error#2144ca19 error_code:int error_message:string = RpcError;
//
func (m *TLRpcError) To_RpcError() *RpcError {
	m.Data2.Cmd = Cmd_RpcError
	return m.Data2
}

//  + TL_SavedPhoneContact
//

//  savedPhoneContact#1142bd56 phone:string first_name:string last_name:string date:int = SavedContact;
//
func (m *SavedContact) To_SavedPhoneContact() *TLSavedPhoneContact {
	return &TLSavedPhoneContact{
		Data2: m,
	}
}

//  savedPhoneContact#1142bd56 phone:string first_name:string last_name:string date:int = SavedContact;
//
func (m *TLSavedPhoneContact) To_SavedContact() *SavedContact {
	m.Data2.Cmd = Cmd_SavedPhoneContact
	return m.Data2
}

//  + TL_SecureCredentialsEncrypted
//

//  secureCredentialsEncrypted#33f0ea47 data:bytes hash:bytes secret:bytes = SecureCredentialsEncrypted;
//
func (m *SecureCredentialsEncrypted) To_SecureCredentialsEncrypted() *TLSecureCredentialsEncrypted {
	return &TLSecureCredentialsEncrypted{
		Data2: m,
	}
}

//  secureCredentialsEncrypted#33f0ea47 data:bytes hash:bytes secret:bytes = SecureCredentialsEncrypted;
//
func (m *TLSecureCredentialsEncrypted) To_SecureCredentialsEncrypted() *SecureCredentialsEncrypted {
	m.Data2.Cmd = Cmd_SecureCredentialsEncrypted
	return m.Data2
}

//  + TL_SecureData
//

//  secureData#8aeabec3 data:bytes data_hash:bytes secret:bytes = SecureData;
//
func (m *SecureData) To_SecureData() *TLSecureData {
	return &TLSecureData{
		Data2: m,
	}
}

//  secureData#8aeabec3 data:bytes data_hash:bytes secret:bytes = SecureData;
//
func (m *TLSecureData) To_SecureData() *SecureData {
	m.Data2.Cmd = Cmd_SecureData
	return m.Data2
}

//  + TL_SecureFileEmpty
//  + TL_SecureFile
//

//  secureFileEmpty#64199744 = SecureFile;
//
func (m *SecureFile) To_SecureFileEmpty() *TLSecureFileEmpty {
	return &TLSecureFileEmpty{
		Data2: m,
	}
}

//  secureFile#e0277a62 id:long access_hash:long size:int dc_id:int date:int file_hash:bytes secret:bytes = SecureFile;
//
func (m *SecureFile) To_SecureFile() *TLSecureFile {
	return &TLSecureFile{
		Data2: m,
	}
}

//  secureFileEmpty#64199744 = SecureFile;
//
func (m *TLSecureFileEmpty) To_SecureFile() *SecureFile {
	m.Data2.Cmd = Cmd_SecureFileEmpty
	return m.Data2
}

//  secureFile#e0277a62 id:long access_hash:long size:int dc_id:int date:int file_hash:bytes secret:bytes = SecureFile;
//
func (m *TLSecureFile) To_SecureFile() *SecureFile {
	m.Data2.Cmd = Cmd_SecureFile
	return m.Data2
}

//  + TL_SecurePasswordKdfAlgoUnknown
//  + TL_SecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000
//  + TL_SecurePasswordKdfAlgoSHA512
//

//  securePasswordKdfAlgoUnknown#4a8537 = SecurePasswordKdfAlgo;
//
func (m *SecurePasswordKdfAlgo) To_SecurePasswordKdfAlgoUnknown() *TLSecurePasswordKdfAlgoUnknown {
	return &TLSecurePasswordKdfAlgoUnknown{
		Data2: m,
	}
}

//  securePasswordKdfAlgoPBKDF2HMACSHA512iter100000#bbf2dda0 salt:bytes = SecurePasswordKdfAlgo;
//
func (m *SecurePasswordKdfAlgo) To_SecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000() *TLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000 {
	return &TLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000{
		Data2: m,
	}
}

//  securePasswordKdfAlgoSHA512#86471d92 salt:bytes = SecurePasswordKdfAlgo;
//
func (m *SecurePasswordKdfAlgo) To_SecurePasswordKdfAlgoSHA512() *TLSecurePasswordKdfAlgoSHA512 {
	return &TLSecurePasswordKdfAlgoSHA512{
		Data2: m,
	}
}

//  securePasswordKdfAlgoUnknown#4a8537 = SecurePasswordKdfAlgo;
//
func (m *TLSecurePasswordKdfAlgoUnknown) To_SecurePasswordKdfAlgo() *SecurePasswordKdfAlgo {
	m.Data2.Cmd = Cmd_SecurePasswordKdfAlgoUnknown
	return m.Data2
}

//  securePasswordKdfAlgoPBKDF2HMACSHA512iter100000#bbf2dda0 salt:bytes = SecurePasswordKdfAlgo;
//
func (m *TLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000) To_SecurePasswordKdfAlgo() *SecurePasswordKdfAlgo {
	m.Data2.Cmd = Cmd_SecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000
	return m.Data2
}

//  securePasswordKdfAlgoSHA512#86471d92 salt:bytes = SecurePasswordKdfAlgo;
//
func (m *TLSecurePasswordKdfAlgoSHA512) To_SecurePasswordKdfAlgo() *SecurePasswordKdfAlgo {
	m.Data2.Cmd = Cmd_SecurePasswordKdfAlgoSHA512
	return m.Data2
}

//  + TL_SecurePlainPhone
//  + TL_SecurePlainEmail
//

//  securePlainPhone#7d6099dd phone:string = SecurePlainData;
//
func (m *SecurePlainData) To_SecurePlainPhone() *TLSecurePlainPhone {
	return &TLSecurePlainPhone{
		Data2: m,
	}
}

//  securePlainEmail#21ec5a5f email:string = SecurePlainData;
//
func (m *SecurePlainData) To_SecurePlainEmail() *TLSecurePlainEmail {
	return &TLSecurePlainEmail{
		Data2: m,
	}
}

//  securePlainPhone#7d6099dd phone:string = SecurePlainData;
//
func (m *TLSecurePlainPhone) To_SecurePlainData() *SecurePlainData {
	m.Data2.Cmd = Cmd_SecurePlainPhone
	return m.Data2
}

//  securePlainEmail#21ec5a5f email:string = SecurePlainData;
//
func (m *TLSecurePlainEmail) To_SecurePlainData() *SecurePlainData {
	m.Data2.Cmd = Cmd_SecurePlainEmail
	return m.Data2
}

//  + TL_SecureRequiredType
//  + TL_SecureRequiredTypeOneOf
//

//  secureRequiredType#829d99da flags:# native_names:flags.0?true selfie_required:flags.1?true translation_required:flags.2?true type:SecureValueType = SecureRequiredType;
//
func (m *SecureRequiredType) To_SecureRequiredType() *TLSecureRequiredType {
	return &TLSecureRequiredType{
		Data2: m,
	}
}

//  secureRequiredTypeOneOf#27477b4 types:Vector<SecureRequiredType> = SecureRequiredType;
//
func (m *SecureRequiredType) To_SecureRequiredTypeOneOf() *TLSecureRequiredTypeOneOf {
	return &TLSecureRequiredTypeOneOf{
		Data2: m,
	}
}

//  secureRequiredType#829d99da flags:# native_names:flags.0?true selfie_required:flags.1?true translation_required:flags.2?true type:SecureValueType = SecureRequiredType;
//
func (m *TLSecureRequiredType) To_SecureRequiredType() *SecureRequiredType {
	m.Data2.Cmd = Cmd_SecureRequiredType
	return m.Data2
}

//  secureRequiredTypeOneOf#27477b4 types:Vector<SecureRequiredType> = SecureRequiredType;
//
func (m *TLSecureRequiredTypeOneOf) To_SecureRequiredType() *SecureRequiredType {
	m.Data2.Cmd = Cmd_SecureRequiredTypeOneOf
	return m.Data2
}

//  + TL_SecureSecretSettings
//

//  secureSecretSettings#1527bcac secure_algo:SecurePasswordKdfAlgo secure_secret:bytes secure_secret_id:long = SecureSecretSettings;
//
func (m *SecureSecretSettings) To_SecureSecretSettings() *TLSecureSecretSettings {
	return &TLSecureSecretSettings{
		Data2: m,
	}
}

//  secureSecretSettings#1527bcac secure_algo:SecurePasswordKdfAlgo secure_secret:bytes secure_secret_id:long = SecureSecretSettings;
//
func (m *TLSecureSecretSettings) To_SecureSecretSettings() *SecureSecretSettings {
	m.Data2.Cmd = Cmd_SecureSecretSettings
	return m.Data2
}

//  + TL_SecureValue
//

//  secureValue#187fa0ca flags:# type:SecureValueType data:flags.0?SecureData front_side:flags.1?SecureFile reverse_side:flags.2?SecureFile selfie:flags.3?SecureFile translation:flags.6?Vector<SecureFile> files:flags.4?Vector<SecureFile> plain_data:flags.5?SecurePlainData hash:bytes = SecureValue;
//
func (m *SecureValue) To_SecureValue() *TLSecureValue {
	return &TLSecureValue{
		Data2: m,
	}
}

//  secureValue#187fa0ca flags:# type:SecureValueType data:flags.0?SecureData front_side:flags.1?SecureFile reverse_side:flags.2?SecureFile selfie:flags.3?SecureFile translation:flags.6?Vector<SecureFile> files:flags.4?Vector<SecureFile> plain_data:flags.5?SecurePlainData hash:bytes = SecureValue;
//
func (m *TLSecureValue) To_SecureValue() *SecureValue {
	m.Data2.Cmd = Cmd_SecureValue
	return m.Data2
}

//  + TL_SecureValueErrorData
//  + TL_SecureValueErrorFrontSide
//  + TL_SecureValueErrorReverseSide
//  + TL_SecureValueErrorSelfie
//  + TL_SecureValueErrorFile
//  + TL_SecureValueErrorFiles
//  + TL_SecureValueError
//  + TL_SecureValueErrorTranslationFile
//  + TL_SecureValueErrorTranslationFiles
//

//  secureValueErrorData#e8a40bd9 type:SecureValueType data_hash:bytes field:string text:string = SecureValueError;
//
func (m *SecureValueError) To_SecureValueErrorData() *TLSecureValueErrorData {
	return &TLSecureValueErrorData{
		Data2: m,
	}
}

//  secureValueErrorFrontSide#be3dfa type:SecureValueType file_hash:bytes text:string = SecureValueError;
//
func (m *SecureValueError) To_SecureValueErrorFrontSide() *TLSecureValueErrorFrontSide {
	return &TLSecureValueErrorFrontSide{
		Data2: m,
	}
}

//  secureValueErrorReverseSide#868a2aa5 type:SecureValueType file_hash:bytes text:string = SecureValueError;
//
func (m *SecureValueError) To_SecureValueErrorReverseSide() *TLSecureValueErrorReverseSide {
	return &TLSecureValueErrorReverseSide{
		Data2: m,
	}
}

//  secureValueErrorSelfie#e537ced6 type:SecureValueType file_hash:bytes text:string = SecureValueError;
//
func (m *SecureValueError) To_SecureValueErrorSelfie() *TLSecureValueErrorSelfie {
	return &TLSecureValueErrorSelfie{
		Data2: m,
	}
}

//  secureValueErrorFile#7a700873 type:SecureValueType file_hash:bytes text:string = SecureValueError;
//
func (m *SecureValueError) To_SecureValueErrorFile() *TLSecureValueErrorFile {
	return &TLSecureValueErrorFile{
		Data2: m,
	}
}

//  secureValueErrorFiles#666220e9 type:SecureValueType file_hash:Vector<bytes> text:string = SecureValueError;
//
func (m *SecureValueError) To_SecureValueErrorFiles() *TLSecureValueErrorFiles {
	return &TLSecureValueErrorFiles{
		Data2: m,
	}
}

//  secureValueError#869d758f type:SecureValueType hash:bytes text:string = SecureValueError;
//
func (m *SecureValueError) To_SecureValueError() *TLSecureValueError {
	return &TLSecureValueError{
		Data2: m,
	}
}

//  secureValueErrorTranslationFile#a1144770 type:SecureValueType file_hash:bytes text:string = SecureValueError;
//
func (m *SecureValueError) To_SecureValueErrorTranslationFile() *TLSecureValueErrorTranslationFile {
	return &TLSecureValueErrorTranslationFile{
		Data2: m,
	}
}

//  secureValueErrorTranslationFiles#34636dd8 type:SecureValueType file_hash:Vector<bytes> text:string = SecureValueError;
//
func (m *SecureValueError) To_SecureValueErrorTranslationFiles() *TLSecureValueErrorTranslationFiles {
	return &TLSecureValueErrorTranslationFiles{
		Data2: m,
	}
}

//  secureValueErrorData#e8a40bd9 type:SecureValueType data_hash:bytes field:string text:string = SecureValueError;
//
func (m *TLSecureValueErrorData) To_SecureValueError() *SecureValueError {
	m.Data2.Cmd = Cmd_SecureValueErrorData
	return m.Data2
}

//  secureValueErrorFrontSide#be3dfa type:SecureValueType file_hash:bytes text:string = SecureValueError;
//
func (m *TLSecureValueErrorFrontSide) To_SecureValueError() *SecureValueError {
	m.Data2.Cmd = Cmd_SecureValueErrorFrontSide
	return m.Data2
}

//  secureValueErrorReverseSide#868a2aa5 type:SecureValueType file_hash:bytes text:string = SecureValueError;
//
func (m *TLSecureValueErrorReverseSide) To_SecureValueError() *SecureValueError {
	m.Data2.Cmd = Cmd_SecureValueErrorReverseSide
	return m.Data2
}

//  secureValueErrorSelfie#e537ced6 type:SecureValueType file_hash:bytes text:string = SecureValueError;
//
func (m *TLSecureValueErrorSelfie) To_SecureValueError() *SecureValueError {
	m.Data2.Cmd = Cmd_SecureValueErrorSelfie
	return m.Data2
}

//  secureValueErrorFile#7a700873 type:SecureValueType file_hash:bytes text:string = SecureValueError;
//
func (m *TLSecureValueErrorFile) To_SecureValueError() *SecureValueError {
	m.Data2.Cmd = Cmd_SecureValueErrorFile
	return m.Data2
}

//  secureValueErrorFiles#666220e9 type:SecureValueType file_hash:Vector<bytes> text:string = SecureValueError;
//
func (m *TLSecureValueErrorFiles) To_SecureValueError() *SecureValueError {
	m.Data2.Cmd = Cmd_SecureValueErrorFiles
	return m.Data2
}

//  secureValueError#869d758f type:SecureValueType hash:bytes text:string = SecureValueError;
//
func (m *TLSecureValueError) To_SecureValueError() *SecureValueError {
	m.Data2.Cmd = Cmd_SecureValueError
	return m.Data2
}

//  secureValueErrorTranslationFile#a1144770 type:SecureValueType file_hash:bytes text:string = SecureValueError;
//
func (m *TLSecureValueErrorTranslationFile) To_SecureValueError() *SecureValueError {
	m.Data2.Cmd = Cmd_SecureValueErrorTranslationFile
	return m.Data2
}

//  secureValueErrorTranslationFiles#34636dd8 type:SecureValueType file_hash:Vector<bytes> text:string = SecureValueError;
//
func (m *TLSecureValueErrorTranslationFiles) To_SecureValueError() *SecureValueError {
	m.Data2.Cmd = Cmd_SecureValueErrorTranslationFiles
	return m.Data2
}

//  + TL_SecureValueHash
//

//  secureValueHash#ed1ecdb0 type:SecureValueType hash:bytes = SecureValueHash;
//
func (m *SecureValueHash) To_SecureValueHash() *TLSecureValueHash {
	return &TLSecureValueHash{
		Data2: m,
	}
}

//  secureValueHash#ed1ecdb0 type:SecureValueType hash:bytes = SecureValueHash;
//
func (m *TLSecureValueHash) To_SecureValueHash() *SecureValueHash {
	m.Data2.Cmd = Cmd_SecureValueHash
	return m.Data2
}

//  + TL_SecureValueTypePersonalDetails
//  + TL_SecureValueTypePassport
//  + TL_SecureValueTypeDriverLicense
//  + TL_SecureValueTypeIdentityCard
//  + TL_SecureValueTypeInternalPassport
//  + TL_SecureValueTypeAddress
//  + TL_SecureValueTypeUtilityBill
//  + TL_SecureValueTypeBankStatement
//  + TL_SecureValueTypeRentalAgreement
//  + TL_SecureValueTypePassportRegistration
//  + TL_SecureValueTypeTemporaryRegistration
//  + TL_SecureValueTypePhone
//  + TL_SecureValueTypeEmail
//

//  secureValueTypePersonalDetails#9d2a81e3 = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypePersonalDetails() *TLSecureValueTypePersonalDetails {
	return &TLSecureValueTypePersonalDetails{
		Data2: m,
	}
}

//  secureValueTypePassport#3dac6a00 = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypePassport() *TLSecureValueTypePassport {
	return &TLSecureValueTypePassport{
		Data2: m,
	}
}

//  secureValueTypeDriverLicense#6e425c4 = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypeDriverLicense() *TLSecureValueTypeDriverLicense {
	return &TLSecureValueTypeDriverLicense{
		Data2: m,
	}
}

//  secureValueTypeIdentityCard#a0d0744b = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypeIdentityCard() *TLSecureValueTypeIdentityCard {
	return &TLSecureValueTypeIdentityCard{
		Data2: m,
	}
}

//  secureValueTypeInternalPassport#99a48f23 = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypeInternalPassport() *TLSecureValueTypeInternalPassport {
	return &TLSecureValueTypeInternalPassport{
		Data2: m,
	}
}

//  secureValueTypeAddress#cbe31e26 = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypeAddress() *TLSecureValueTypeAddress {
	return &TLSecureValueTypeAddress{
		Data2: m,
	}
}

//  secureValueTypeUtilityBill#fc36954e = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypeUtilityBill() *TLSecureValueTypeUtilityBill {
	return &TLSecureValueTypeUtilityBill{
		Data2: m,
	}
}

//  secureValueTypeBankStatement#89137c0d = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypeBankStatement() *TLSecureValueTypeBankStatement {
	return &TLSecureValueTypeBankStatement{
		Data2: m,
	}
}

//  secureValueTypeRentalAgreement#8b883488 = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypeRentalAgreement() *TLSecureValueTypeRentalAgreement {
	return &TLSecureValueTypeRentalAgreement{
		Data2: m,
	}
}

//  secureValueTypePassportRegistration#99e3806a = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypePassportRegistration() *TLSecureValueTypePassportRegistration {
	return &TLSecureValueTypePassportRegistration{
		Data2: m,
	}
}

//  secureValueTypeTemporaryRegistration#ea02ec33 = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypeTemporaryRegistration() *TLSecureValueTypeTemporaryRegistration {
	return &TLSecureValueTypeTemporaryRegistration{
		Data2: m,
	}
}

//  secureValueTypePhone#b320aadb = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypePhone() *TLSecureValueTypePhone {
	return &TLSecureValueTypePhone{
		Data2: m,
	}
}

//  secureValueTypeEmail#8e3ca7ee = SecureValueType;
//
func (m *SecureValueType) To_SecureValueTypeEmail() *TLSecureValueTypeEmail {
	return &TLSecureValueTypeEmail{
		Data2: m,
	}
}

//  secureValueTypePersonalDetails#9d2a81e3 = SecureValueType;
//
func (m *TLSecureValueTypePersonalDetails) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypePersonalDetails
	return m.Data2
}

//  secureValueTypePassport#3dac6a00 = SecureValueType;
//
func (m *TLSecureValueTypePassport) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypePassport
	return m.Data2
}

//  secureValueTypeDriverLicense#6e425c4 = SecureValueType;
//
func (m *TLSecureValueTypeDriverLicense) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypeDriverLicense
	return m.Data2
}

//  secureValueTypeIdentityCard#a0d0744b = SecureValueType;
//
func (m *TLSecureValueTypeIdentityCard) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypeIdentityCard
	return m.Data2
}

//  secureValueTypeInternalPassport#99a48f23 = SecureValueType;
//
func (m *TLSecureValueTypeInternalPassport) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypeInternalPassport
	return m.Data2
}

//  secureValueTypeAddress#cbe31e26 = SecureValueType;
//
func (m *TLSecureValueTypeAddress) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypeAddress
	return m.Data2
}

//  secureValueTypeUtilityBill#fc36954e = SecureValueType;
//
func (m *TLSecureValueTypeUtilityBill) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypeUtilityBill
	return m.Data2
}

//  secureValueTypeBankStatement#89137c0d = SecureValueType;
//
func (m *TLSecureValueTypeBankStatement) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypeBankStatement
	return m.Data2
}

//  secureValueTypeRentalAgreement#8b883488 = SecureValueType;
//
func (m *TLSecureValueTypeRentalAgreement) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypeRentalAgreement
	return m.Data2
}

//  secureValueTypePassportRegistration#99e3806a = SecureValueType;
//
func (m *TLSecureValueTypePassportRegistration) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypePassportRegistration
	return m.Data2
}

//  secureValueTypeTemporaryRegistration#ea02ec33 = SecureValueType;
//
func (m *TLSecureValueTypeTemporaryRegistration) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypeTemporaryRegistration
	return m.Data2
}

//  secureValueTypePhone#b320aadb = SecureValueType;
//
func (m *TLSecureValueTypePhone) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypePhone
	return m.Data2
}

//  secureValueTypeEmail#8e3ca7ee = SecureValueType;
//
func (m *TLSecureValueTypeEmail) To_SecureValueType() *SecureValueType {
	m.Data2.Cmd = Cmd_SecureValueTypeEmail
	return m.Data2
}

//  + TL_SendMessageTypingAction
//  + TL_SendMessageCancelAction
//  + TL_SendMessageRecordVideoAction
//  + TL_SendMessageUploadVideoAction
//  + TL_SendMessageRecordAudioAction
//  + TL_SendMessageUploadAudioAction
//  + TL_SendMessageUploadPhotoAction
//  + TL_SendMessageUploadDocumentAction
//  + TL_SendMessageGeoLocationAction
//  + TL_SendMessageChooseContactAction
//  + TL_SendMessageGamePlayAction
//  + TL_SendMessageRecordRoundAction
//  + TL_SendMessageUploadRoundAction
//  + TL_SpeakingInGroupCallAction
//

//  sendMessageTypingAction#16bf744e = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageTypingAction() *TLSendMessageTypingAction {
	return &TLSendMessageTypingAction{
		Data2: m,
	}
}

//  sendMessageCancelAction#fd5ec8f5 = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageCancelAction() *TLSendMessageCancelAction {
	return &TLSendMessageCancelAction{
		Data2: m,
	}
}

//  sendMessageRecordVideoAction#a187d66f = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageRecordVideoAction() *TLSendMessageRecordVideoAction {
	return &TLSendMessageRecordVideoAction{
		Data2: m,
	}
}

//  sendMessageUploadVideoAction#e9763aec progress:int = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageUploadVideoAction() *TLSendMessageUploadVideoAction {
	return &TLSendMessageUploadVideoAction{
		Data2: m,
	}
}

//  sendMessageRecordAudioAction#d52f73f7 = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageRecordAudioAction() *TLSendMessageRecordAudioAction {
	return &TLSendMessageRecordAudioAction{
		Data2: m,
	}
}

//  sendMessageUploadAudioAction#f351d7ab progress:int = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageUploadAudioAction() *TLSendMessageUploadAudioAction {
	return &TLSendMessageUploadAudioAction{
		Data2: m,
	}
}

//  sendMessageUploadPhotoAction#d1d34a26 progress:int = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageUploadPhotoAction() *TLSendMessageUploadPhotoAction {
	return &TLSendMessageUploadPhotoAction{
		Data2: m,
	}
}

//  sendMessageUploadDocumentAction#aa0cd9e4 progress:int = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageUploadDocumentAction() *TLSendMessageUploadDocumentAction {
	return &TLSendMessageUploadDocumentAction{
		Data2: m,
	}
}

//  sendMessageGeoLocationAction#176f8ba1 = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageGeoLocationAction() *TLSendMessageGeoLocationAction {
	return &TLSendMessageGeoLocationAction{
		Data2: m,
	}
}

//  sendMessageChooseContactAction#628cbc6f = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageChooseContactAction() *TLSendMessageChooseContactAction {
	return &TLSendMessageChooseContactAction{
		Data2: m,
	}
}

//  sendMessageGamePlayAction#dd6a8f48 = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageGamePlayAction() *TLSendMessageGamePlayAction {
	return &TLSendMessageGamePlayAction{
		Data2: m,
	}
}

//  sendMessageRecordRoundAction#88f27fbc = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageRecordRoundAction() *TLSendMessageRecordRoundAction {
	return &TLSendMessageRecordRoundAction{
		Data2: m,
	}
}

//  sendMessageUploadRoundAction#243e1c66 progress:int = SendMessageAction;
//
func (m *SendMessageAction) To_SendMessageUploadRoundAction() *TLSendMessageUploadRoundAction {
	return &TLSendMessageUploadRoundAction{
		Data2: m,
	}
}

//  speakingInGroupCallAction#d92c2285 = SendMessageAction;
//
func (m *SendMessageAction) To_SpeakingInGroupCallAction() *TLSpeakingInGroupCallAction {
	return &TLSpeakingInGroupCallAction{
		Data2: m,
	}
}

//  sendMessageTypingAction#16bf744e = SendMessageAction;
//
func (m *TLSendMessageTypingAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageTypingAction
	return m.Data2
}

//  sendMessageCancelAction#fd5ec8f5 = SendMessageAction;
//
func (m *TLSendMessageCancelAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageCancelAction
	return m.Data2
}

//  sendMessageRecordVideoAction#a187d66f = SendMessageAction;
//
func (m *TLSendMessageRecordVideoAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageRecordVideoAction
	return m.Data2
}

//  sendMessageUploadVideoAction#e9763aec progress:int = SendMessageAction;
//
func (m *TLSendMessageUploadVideoAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageUploadVideoAction
	return m.Data2
}

//  sendMessageRecordAudioAction#d52f73f7 = SendMessageAction;
//
func (m *TLSendMessageRecordAudioAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageRecordAudioAction
	return m.Data2
}

//  sendMessageUploadAudioAction#f351d7ab progress:int = SendMessageAction;
//
func (m *TLSendMessageUploadAudioAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageUploadAudioAction
	return m.Data2
}

//  sendMessageUploadPhotoAction#d1d34a26 progress:int = SendMessageAction;
//
func (m *TLSendMessageUploadPhotoAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageUploadPhotoAction
	return m.Data2
}

//  sendMessageUploadDocumentAction#aa0cd9e4 progress:int = SendMessageAction;
//
func (m *TLSendMessageUploadDocumentAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageUploadDocumentAction
	return m.Data2
}

//  sendMessageGeoLocationAction#176f8ba1 = SendMessageAction;
//
func (m *TLSendMessageGeoLocationAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageGeoLocationAction
	return m.Data2
}

//  sendMessageChooseContactAction#628cbc6f = SendMessageAction;
//
func (m *TLSendMessageChooseContactAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageChooseContactAction
	return m.Data2
}

//  sendMessageGamePlayAction#dd6a8f48 = SendMessageAction;
//
func (m *TLSendMessageGamePlayAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageGamePlayAction
	return m.Data2
}

//  sendMessageRecordRoundAction#88f27fbc = SendMessageAction;
//
func (m *TLSendMessageRecordRoundAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageRecordRoundAction
	return m.Data2
}

//  sendMessageUploadRoundAction#243e1c66 progress:int = SendMessageAction;
//
func (m *TLSendMessageUploadRoundAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SendMessageUploadRoundAction
	return m.Data2
}

//  speakingInGroupCallAction#d92c2285 = SendMessageAction;
//
func (m *TLSpeakingInGroupCallAction) To_SendMessageAction() *SendMessageAction {
	m.Data2.Cmd = Cmd_SpeakingInGroupCallAction
	return m.Data2
}

//  + TL_Server_DHInnerData
//

//  server_DH_inner_data#b5890dba nonce:int128 server_nonce:int128 g:int dh_prime:string g_a:string server_time:int = Server_DH_inner_data;
//
func (m *Server_DHInnerData) To_Server_DHInnerData() *TLServer_DHInnerData {
	return &TLServer_DHInnerData{
		Data2: m,
	}
}

//  server_DH_inner_data#b5890dba nonce:int128 server_nonce:int128 g:int dh_prime:string g_a:string server_time:int = Server_DH_inner_data;
//
func (m *TLServer_DHInnerData) To_Server_DHInnerData() *Server_DHInnerData {
	m.Data2.Cmd = Cmd_Server_DHInnerData
	return m.Data2
}

//  + TL_Server_DHParamsFail
//  + TL_Server_DHParamsOk
//

//  server_DH_params_fail#79cb045d nonce:int128 server_nonce:int128 new_nonce_hash:int128 = Server_DH_Params;
//
func (m *Server_DH_Params) To_Server_DHParamsFail() *TLServer_DHParamsFail {
	return &TLServer_DHParamsFail{
		Data2: m,
	}
}

//  server_DH_params_ok#d0e8075c nonce:int128 server_nonce:int128 encrypted_answer:string = Server_DH_Params;
//
func (m *Server_DH_Params) To_Server_DHParamsOk() *TLServer_DHParamsOk {
	return &TLServer_DHParamsOk{
		Data2: m,
	}
}

//  server_DH_params_fail#79cb045d nonce:int128 server_nonce:int128 new_nonce_hash:int128 = Server_DH_Params;
//
func (m *TLServer_DHParamsFail) To_Server_DH_Params() *Server_DH_Params {
	m.Data2.Cmd = Cmd_Server_DHParamsFail
	return m.Data2
}

//  server_DH_params_ok#d0e8075c nonce:int128 server_nonce:int128 encrypted_answer:string = Server_DH_Params;
//
func (m *TLServer_DHParamsOk) To_Server_DH_Params() *Server_DH_Params {
	m.Data2.Cmd = Cmd_Server_DHParamsOk
	return m.Data2
}

//  + TL_DhGenOk
//  + TL_DhGenRetry
//  + TL_DhGenFail
//

//  dh_gen_ok#3bcbf734 nonce:int128 server_nonce:int128 new_nonce_hash1:int128 = Set_client_DH_params_answer;
//
func (m *SetClient_DHParamsAnswer) To_DhGenOk() *TLDhGenOk {
	return &TLDhGenOk{
		Data2: m,
	}
}

//  dh_gen_retry#46dc1fb9 nonce:int128 server_nonce:int128 new_nonce_hash2:int128 = Set_client_DH_params_answer;
//
func (m *SetClient_DHParamsAnswer) To_DhGenRetry() *TLDhGenRetry {
	return &TLDhGenRetry{
		Data2: m,
	}
}

//  dh_gen_fail#a69dae02 nonce:int128 server_nonce:int128 new_nonce_hash3:int128 = Set_client_DH_params_answer;
//
func (m *SetClient_DHParamsAnswer) To_DhGenFail() *TLDhGenFail {
	return &TLDhGenFail{
		Data2: m,
	}
}

//  dh_gen_ok#3bcbf734 nonce:int128 server_nonce:int128 new_nonce_hash1:int128 = Set_client_DH_params_answer;
//
func (m *TLDhGenOk) To_SetClient_DHParamsAnswer() *SetClient_DHParamsAnswer {
	m.Data2.Cmd = Cmd_DhGenOk
	return m.Data2
}

//  dh_gen_retry#46dc1fb9 nonce:int128 server_nonce:int128 new_nonce_hash2:int128 = Set_client_DH_params_answer;
//
func (m *TLDhGenRetry) To_SetClient_DHParamsAnswer() *SetClient_DHParamsAnswer {
	m.Data2.Cmd = Cmd_DhGenRetry
	return m.Data2
}

//  dh_gen_fail#a69dae02 nonce:int128 server_nonce:int128 new_nonce_hash3:int128 = Set_client_DH_params_answer;
//
func (m *TLDhGenFail) To_SetClient_DHParamsAnswer() *SetClient_DHParamsAnswer {
	m.Data2.Cmd = Cmd_DhGenFail
	return m.Data2
}

//  + TL_ShippingOption
//

//  shippingOption#b6213cdf id:string title:string prices:Vector<LabeledPrice> = ShippingOption;
//
func (m *ShippingOption) To_ShippingOption() *TLShippingOption {
	return &TLShippingOption{
		Data2: m,
	}
}

//  shippingOption#b6213cdf id:string title:string prices:Vector<LabeledPrice> = ShippingOption;
//
func (m *TLShippingOption) To_ShippingOption() *ShippingOption {
	m.Data2.Cmd = Cmd_ShippingOption
	return m.Data2
}

//  + TL_StatsAbsValueAndPrev
//

//  statsAbsValueAndPrev#cb43acde current:double previous:double = StatsAbsValueAndPrev;
//
func (m *StatsAbsValueAndPrev) To_StatsAbsValueAndPrev() *TLStatsAbsValueAndPrev {
	return &TLStatsAbsValueAndPrev{
		Data2: m,
	}
}

//  statsAbsValueAndPrev#cb43acde current:double previous:double = StatsAbsValueAndPrev;
//
func (m *TLStatsAbsValueAndPrev) To_StatsAbsValueAndPrev() *StatsAbsValueAndPrev {
	m.Data2.Cmd = Cmd_StatsAbsValueAndPrev
	return m.Data2
}

//  + TL_StatsDateRangeDays
//

//  statsDateRangeDays#b637edaf min_date:int max_date:int = StatsDateRangeDays;
//
func (m *StatsDateRangeDays) To_StatsDateRangeDays() *TLStatsDateRangeDays {
	return &TLStatsDateRangeDays{
		Data2: m,
	}
}

//  statsDateRangeDays#b637edaf min_date:int max_date:int = StatsDateRangeDays;
//
func (m *TLStatsDateRangeDays) To_StatsDateRangeDays() *StatsDateRangeDays {
	m.Data2.Cmd = Cmd_StatsDateRangeDays
	return m.Data2
}

//  + TL_StatsGraphAsync
//  + TL_StatsGraphError
//  + TL_StatsGraph
//

//  statsGraphAsync#4a27eb2d token:string = StatsGraph;
//
func (m *StatsGraph) To_StatsGraphAsync() *TLStatsGraphAsync {
	return &TLStatsGraphAsync{
		Data2: m,
	}
}

//  statsGraphError#bedc9822 error:string = StatsGraph;
//
func (m *StatsGraph) To_StatsGraphError() *TLStatsGraphError {
	return &TLStatsGraphError{
		Data2: m,
	}
}

//  statsGraph#8ea464b6 flags:# json:DataJSON zoom_token:flags.0?string = StatsGraph;
//
func (m *StatsGraph) To_StatsGraph() *TLStatsGraph {
	return &TLStatsGraph{
		Data2: m,
	}
}

//  statsGraphAsync#4a27eb2d token:string = StatsGraph;
//
func (m *TLStatsGraphAsync) To_StatsGraph() *StatsGraph {
	m.Data2.Cmd = Cmd_StatsGraphAsync
	return m.Data2
}

//  statsGraphError#bedc9822 error:string = StatsGraph;
//
func (m *TLStatsGraphError) To_StatsGraph() *StatsGraph {
	m.Data2.Cmd = Cmd_StatsGraphError
	return m.Data2
}

//  statsGraph#8ea464b6 flags:# json:DataJSON zoom_token:flags.0?string = StatsGraph;
//
func (m *TLStatsGraph) To_StatsGraph() *StatsGraph {
	m.Data2.Cmd = Cmd_StatsGraph
	return m.Data2
}

//  + TL_StatsGroupTopAdmin
//

//  statsGroupTopAdmin#6014f412 user_id:int deleted:int kicked:int banned:int = StatsGroupTopAdmin;
//
func (m *StatsGroupTopAdmin) To_StatsGroupTopAdmin() *TLStatsGroupTopAdmin {
	return &TLStatsGroupTopAdmin{
		Data2: m,
	}
}

//  statsGroupTopAdmin#6014f412 user_id:int deleted:int kicked:int banned:int = StatsGroupTopAdmin;
//
func (m *TLStatsGroupTopAdmin) To_StatsGroupTopAdmin() *StatsGroupTopAdmin {
	m.Data2.Cmd = Cmd_StatsGroupTopAdmin
	return m.Data2
}

//  + TL_StatsGroupTopInviter
//

//  statsGroupTopInviter#31962a4c user_id:int invitations:int = StatsGroupTopInviter;
//
func (m *StatsGroupTopInviter) To_StatsGroupTopInviter() *TLStatsGroupTopInviter {
	return &TLStatsGroupTopInviter{
		Data2: m,
	}
}

//  statsGroupTopInviter#31962a4c user_id:int invitations:int = StatsGroupTopInviter;
//
func (m *TLStatsGroupTopInviter) To_StatsGroupTopInviter() *StatsGroupTopInviter {
	m.Data2.Cmd = Cmd_StatsGroupTopInviter
	return m.Data2
}

//  + TL_StatsGroupTopPoster
//

//  statsGroupTopPoster#18f3d0f7 user_id:int messages:int avg_chars:int = StatsGroupTopPoster;
//
func (m *StatsGroupTopPoster) To_StatsGroupTopPoster() *TLStatsGroupTopPoster {
	return &TLStatsGroupTopPoster{
		Data2: m,
	}
}

//  statsGroupTopPoster#18f3d0f7 user_id:int messages:int avg_chars:int = StatsGroupTopPoster;
//
func (m *TLStatsGroupTopPoster) To_StatsGroupTopPoster() *StatsGroupTopPoster {
	m.Data2.Cmd = Cmd_StatsGroupTopPoster
	return m.Data2
}

//  + TL_StatsPercentValue
//

//  statsPercentValue#cbce2fe0 part:double total:double = StatsPercentValue;
//
func (m *StatsPercentValue) To_StatsPercentValue() *TLStatsPercentValue {
	return &TLStatsPercentValue{
		Data2: m,
	}
}

//  statsPercentValue#cbce2fe0 part:double total:double = StatsPercentValue;
//
func (m *TLStatsPercentValue) To_StatsPercentValue() *StatsPercentValue {
	m.Data2.Cmd = Cmd_StatsPercentValue
	return m.Data2
}

//  + TL_StatsURL
//

//  statsURL#47a971e0 url:string = StatsURL;
//
func (m *StatsURL) To_StatsURL() *TLStatsURL {
	return &TLStatsURL{
		Data2: m,
	}
}

//  statsURL#47a971e0 url:string = StatsURL;
//
func (m *TLStatsURL) To_StatsURL() *StatsURL {
	m.Data2.Cmd = Cmd_StatsURL
	return m.Data2
}

//  + TL_StatsBroadcastStats
//

//  stats.broadcastStats#bdf78394 period:StatsDateRangeDays followers:StatsAbsValueAndPrev views_per_post:StatsAbsValueAndPrev shares_per_post:StatsAbsValueAndPrev enabled_notifications:StatsPercentValue growth_graph:StatsGraph followers_graph:StatsGraph mute_graph:StatsGraph top_hours_graph:StatsGraph interactions_graph:StatsGraph iv_interactions_graph:StatsGraph views_by_source_graph:StatsGraph new_followers_by_source_graph:StatsGraph languages_graph:StatsGraph recent_message_interactions:Vector<MessageInteractionCounters> = stats.BroadcastStats;
//
func (m *Stats_BroadcastStats) To_StatsBroadcastStats() *TLStatsBroadcastStats {
	return &TLStatsBroadcastStats{
		Data2: m,
	}
}

//  stats.broadcastStats#bdf78394 period:StatsDateRangeDays followers:StatsAbsValueAndPrev views_per_post:StatsAbsValueAndPrev shares_per_post:StatsAbsValueAndPrev enabled_notifications:StatsPercentValue growth_graph:StatsGraph followers_graph:StatsGraph mute_graph:StatsGraph top_hours_graph:StatsGraph interactions_graph:StatsGraph iv_interactions_graph:StatsGraph views_by_source_graph:StatsGraph new_followers_by_source_graph:StatsGraph languages_graph:StatsGraph recent_message_interactions:Vector<MessageInteractionCounters> = stats.BroadcastStats;
//
func (m *TLStatsBroadcastStats) To_Stats_BroadcastStats() *Stats_BroadcastStats {
	m.Data2.Cmd = Cmd_StatsBroadcastStats
	return m.Data2
}

//  + TL_StatsMegagroupStats
//

//  stats.megagroupStats#ef7ff916 period:StatsDateRangeDays members:StatsAbsValueAndPrev messages:StatsAbsValueAndPrev viewers:StatsAbsValueAndPrev posters:StatsAbsValueAndPrev growth_graph:StatsGraph members_graph:StatsGraph new_members_by_source_graph:StatsGraph languages_graph:StatsGraph messages_graph:StatsGraph actions_graph:StatsGraph top_hours_graph:StatsGraph weekdays_graph:StatsGraph top_posters:Vector<StatsGroupTopPoster> top_admins:Vector<StatsGroupTopAdmin> top_inviters:Vector<StatsGroupTopInviter> users:Vector<User> = stats.MegagroupStats;
//
func (m *Stats_MegagroupStats) To_StatsMegagroupStats() *TLStatsMegagroupStats {
	return &TLStatsMegagroupStats{
		Data2: m,
	}
}

//  stats.megagroupStats#ef7ff916 period:StatsDateRangeDays members:StatsAbsValueAndPrev messages:StatsAbsValueAndPrev viewers:StatsAbsValueAndPrev posters:StatsAbsValueAndPrev growth_graph:StatsGraph members_graph:StatsGraph new_members_by_source_graph:StatsGraph languages_graph:StatsGraph messages_graph:StatsGraph actions_graph:StatsGraph top_hours_graph:StatsGraph weekdays_graph:StatsGraph top_posters:Vector<StatsGroupTopPoster> top_admins:Vector<StatsGroupTopAdmin> top_inviters:Vector<StatsGroupTopInviter> users:Vector<User> = stats.MegagroupStats;
//
func (m *TLStatsMegagroupStats) To_Stats_MegagroupStats() *Stats_MegagroupStats {
	m.Data2.Cmd = Cmd_StatsMegagroupStats
	return m.Data2
}

//  + TL_StatsMessageStats
//

//  stats.messageStats#8999f295 views_graph:StatsGraph = stats.MessageStats;
//
func (m *Stats_MessageStats) To_StatsMessageStats() *TLStatsMessageStats {
	return &TLStatsMessageStats{
		Data2: m,
	}
}

//  stats.messageStats#8999f295 views_graph:StatsGraph = stats.MessageStats;
//
func (m *TLStatsMessageStats) To_Stats_MessageStats() *Stats_MessageStats {
	m.Data2.Cmd = Cmd_StatsMessageStats
	return m.Data2
}

//  + TL_StickerPack
//

//  stickerPack#12b299d4 emoticon:string documents:Vector<long> = StickerPack;
//
func (m *StickerPack) To_StickerPack() *TLStickerPack {
	return &TLStickerPack{
		Data2: m,
	}
}

//  stickerPack#12b299d4 emoticon:string documents:Vector<long> = StickerPack;
//
func (m *TLStickerPack) To_StickerPack() *StickerPack {
	m.Data2.Cmd = Cmd_StickerPack
	return m.Data2
}

//  + TL_StickerSet
//

//  stickerSet#cd303b41 flags:# installed:flags.0?true archived:flags.1?true official:flags.2?true masks:flags.3?true id:long access_hash:long title:string short_name:string count:int hash:int = StickerSet;
//  stickerSet#5585a139 flags:# archived:flags.1?true official:flags.2?true masks:flags.3?true installed_date:flags.0?int id:long access_hash:long title:string short_name:string count:int hash:int = StickerSet;
//  stickerSet#6a90bcb7 flags:# archived:flags.1?true official:flags.2?true masks:flags.3?true installed_date:flags.0?int id:long access_hash:long title:string short_name:string thumb:flags.4?PhotoSize count:int hash:int = StickerSet;
//  stickerSet#eeb46f27 flags:# archived:flags.1?true official:flags.2?true masks:flags.3?true installed_date:flags.0?int id:long access_hash:long title:string short_name:string thumb:flags.4?PhotoSize thumb_dc_id:flags.4?int count:int hash:int = StickerSet;
//  stickerSet#40e237a8 flags:# archived:flags.1?true official:flags.2?true masks:flags.3?true animated:flags.5?true installed_date:flags.0?int id:long access_hash:long title:string short_name:string thumbs:flags.4?Vector<PhotoSize> thumb_dc_id:flags.4?int count:int hash:int = StickerSet;
//
func (m *StickerSet) To_StickerSet() *TLStickerSet {
	return &TLStickerSet{
		Data2: m,
	}
}

//  stickerSet#cd303b41 flags:# installed:flags.0?true archived:flags.1?true official:flags.2?true masks:flags.3?true id:long access_hash:long title:string short_name:string count:int hash:int = StickerSet;
//  stickerSet#5585a139 flags:# archived:flags.1?true official:flags.2?true masks:flags.3?true installed_date:flags.0?int id:long access_hash:long title:string short_name:string count:int hash:int = StickerSet;
//  stickerSet#6a90bcb7 flags:# archived:flags.1?true official:flags.2?true masks:flags.3?true installed_date:flags.0?int id:long access_hash:long title:string short_name:string thumb:flags.4?PhotoSize count:int hash:int = StickerSet;
//  stickerSet#eeb46f27 flags:# archived:flags.1?true official:flags.2?true masks:flags.3?true installed_date:flags.0?int id:long access_hash:long title:string short_name:string thumb:flags.4?PhotoSize thumb_dc_id:flags.4?int count:int hash:int = StickerSet;
//  stickerSet#40e237a8 flags:# archived:flags.1?true official:flags.2?true masks:flags.3?true animated:flags.5?true installed_date:flags.0?int id:long access_hash:long title:string short_name:string thumbs:flags.4?Vector<PhotoSize> thumb_dc_id:flags.4?int count:int hash:int = StickerSet;
//
func (m *TLStickerSet) To_StickerSet() *StickerSet {
	m.Data2.Cmd = Cmd_StickerSet
	return m.Data2
}

//  + TL_StickerSetCovered
//  + TL_StickerSetMultiCovered
//

//  stickerSetCovered#6410a5d2 set:StickerSet cover:Document = StickerSetCovered;
//
func (m *StickerSetCovered) To_StickerSetCovered() *TLStickerSetCovered {
	return &TLStickerSetCovered{
		Data2: m,
	}
}

//  stickerSetMultiCovered#3407e51b set:StickerSet covers:Vector<Document> = StickerSetCovered;
//
func (m *StickerSetCovered) To_StickerSetMultiCovered() *TLStickerSetMultiCovered {
	return &TLStickerSetMultiCovered{
		Data2: m,
	}
}

//  stickerSetCovered#6410a5d2 set:StickerSet cover:Document = StickerSetCovered;
//
func (m *TLStickerSetCovered) To_StickerSetCovered() *StickerSetCovered {
	m.Data2.Cmd = Cmd_StickerSetCovered
	return m.Data2
}

//  stickerSetMultiCovered#3407e51b set:StickerSet covers:Vector<Document> = StickerSetCovered;
//
func (m *TLStickerSetMultiCovered) To_StickerSetCovered() *StickerSetCovered {
	m.Data2.Cmd = Cmd_StickerSetMultiCovered
	return m.Data2
}

//  + TL_StorageFileUnknown
//  + TL_StorageFilePartial
//  + TL_StorageFileJpeg
//  + TL_StorageFileGif
//  + TL_StorageFilePng
//  + TL_StorageFilePdf
//  + TL_StorageFileMp3
//  + TL_StorageFileMov
//  + TL_StorageFileMp4
//  + TL_StorageFileWebp
//

//  storage.fileUnknown#aa963b05 = storage.FileType;
//
func (m *Storage_FileType) To_StorageFileUnknown() *TLStorageFileUnknown {
	return &TLStorageFileUnknown{
		Data2: m,
	}
}

//  storage.filePartial#40bc6f52 = storage.FileType;
//
func (m *Storage_FileType) To_StorageFilePartial() *TLStorageFilePartial {
	return &TLStorageFilePartial{
		Data2: m,
	}
}

//  storage.fileJpeg#7efe0e = storage.FileType;
//
func (m *Storage_FileType) To_StorageFileJpeg() *TLStorageFileJpeg {
	return &TLStorageFileJpeg{
		Data2: m,
	}
}

//  storage.fileGif#cae1aadf = storage.FileType;
//
func (m *Storage_FileType) To_StorageFileGif() *TLStorageFileGif {
	return &TLStorageFileGif{
		Data2: m,
	}
}

//  storage.filePng#a4f63c0 = storage.FileType;
//
func (m *Storage_FileType) To_StorageFilePng() *TLStorageFilePng {
	return &TLStorageFilePng{
		Data2: m,
	}
}

//  storage.filePdf#ae1e508d = storage.FileType;
//
func (m *Storage_FileType) To_StorageFilePdf() *TLStorageFilePdf {
	return &TLStorageFilePdf{
		Data2: m,
	}
}

//  storage.fileMp3#528a0677 = storage.FileType;
//
func (m *Storage_FileType) To_StorageFileMp3() *TLStorageFileMp3 {
	return &TLStorageFileMp3{
		Data2: m,
	}
}

//  storage.fileMov#4b09ebbc = storage.FileType;
//
func (m *Storage_FileType) To_StorageFileMov() *TLStorageFileMov {
	return &TLStorageFileMov{
		Data2: m,
	}
}

//  storage.fileMp4#b3cea0e4 = storage.FileType;
//
func (m *Storage_FileType) To_StorageFileMp4() *TLStorageFileMp4 {
	return &TLStorageFileMp4{
		Data2: m,
	}
}

//  storage.fileWebp#1081464c = storage.FileType;
//
func (m *Storage_FileType) To_StorageFileWebp() *TLStorageFileWebp {
	return &TLStorageFileWebp{
		Data2: m,
	}
}

//  storage.fileUnknown#aa963b05 = storage.FileType;
//
func (m *TLStorageFileUnknown) To_Storage_FileType() *Storage_FileType {
	m.Data2.Cmd = Cmd_StorageFileUnknown
	return m.Data2
}

//  storage.filePartial#40bc6f52 = storage.FileType;
//
func (m *TLStorageFilePartial) To_Storage_FileType() *Storage_FileType {
	m.Data2.Cmd = Cmd_StorageFilePartial
	return m.Data2
}

//  storage.fileJpeg#7efe0e = storage.FileType;
//
func (m *TLStorageFileJpeg) To_Storage_FileType() *Storage_FileType {
	m.Data2.Cmd = Cmd_StorageFileJpeg
	return m.Data2
}

//  storage.fileGif#cae1aadf = storage.FileType;
//
func (m *TLStorageFileGif) To_Storage_FileType() *Storage_FileType {
	m.Data2.Cmd = Cmd_StorageFileGif
	return m.Data2
}

//  storage.filePng#a4f63c0 = storage.FileType;
//
func (m *TLStorageFilePng) To_Storage_FileType() *Storage_FileType {
	m.Data2.Cmd = Cmd_StorageFilePng
	return m.Data2
}

//  storage.filePdf#ae1e508d = storage.FileType;
//
func (m *TLStorageFilePdf) To_Storage_FileType() *Storage_FileType {
	m.Data2.Cmd = Cmd_StorageFilePdf
	return m.Data2
}

//  storage.fileMp3#528a0677 = storage.FileType;
//
func (m *TLStorageFileMp3) To_Storage_FileType() *Storage_FileType {
	m.Data2.Cmd = Cmd_StorageFileMp3
	return m.Data2
}

//  storage.fileMov#4b09ebbc = storage.FileType;
//
func (m *TLStorageFileMov) To_Storage_FileType() *Storage_FileType {
	m.Data2.Cmd = Cmd_StorageFileMov
	return m.Data2
}

//  storage.fileMp4#b3cea0e4 = storage.FileType;
//
func (m *TLStorageFileMp4) To_Storage_FileType() *Storage_FileType {
	m.Data2.Cmd = Cmd_StorageFileMp4
	return m.Data2
}

//  storage.fileWebp#1081464c = storage.FileType;
//
func (m *TLStorageFileWebp) To_Storage_FileType() *Storage_FileType {
	m.Data2.Cmd = Cmd_StorageFileWebp
	return m.Data2
}

//  + TL_ThemeDocumentNotModified
//  + TL_Theme
//

//  themeDocumentNotModified#483d270c = Theme;
//
func (m *Theme) To_ThemeDocumentNotModified() *TLThemeDocumentNotModified {
	return &TLThemeDocumentNotModified{
		Data2: m,
	}
}

//  theme#574e4de9 flags:# creator:flags.0?true default:flags.1?true id:long access_hash:long slug:string title:string document:Document = Theme;
//  theme#f7d90ce0 flags:# creator:flags.0?true default:flags.1?true id:long access_hash:long slug:string title:string document:flags.2?Document installs_count:int = Theme;
//  theme#28f1114 flags:# creator:flags.0?true default:flags.1?true id:long access_hash:long slug:string title:string document:flags.2?Document settings:flags.3?ThemeSettings installs_count:int = Theme;
//
func (m *Theme) To_Theme() *TLTheme {
	return &TLTheme{
		Data2: m,
	}
}

//  themeDocumentNotModified#483d270c = Theme;
//
func (m *TLThemeDocumentNotModified) To_Theme() *Theme {
	m.Data2.Cmd = Cmd_ThemeDocumentNotModified
	return m.Data2
}

//  theme#574e4de9 flags:# creator:flags.0?true default:flags.1?true id:long access_hash:long slug:string title:string document:Document = Theme;
//  theme#f7d90ce0 flags:# creator:flags.0?true default:flags.1?true id:long access_hash:long slug:string title:string document:flags.2?Document installs_count:int = Theme;
//  theme#28f1114 flags:# creator:flags.0?true default:flags.1?true id:long access_hash:long slug:string title:string document:flags.2?Document settings:flags.3?ThemeSettings installs_count:int = Theme;
//
func (m *TLTheme) To_Theme() *Theme {
	m.Data2.Cmd = Cmd_Theme
	return m.Data2
}

//  + TL_ThemeSettings
//

//  themeSettings#9c14984a flags:# base_theme:BaseTheme accent_color:int message_top_color:flags.0?int message_bottom_color:flags.0?int wallpaper:flags.1?WallPaper = ThemeSettings;
//
func (m *ThemeSettings) To_ThemeSettings() *TLThemeSettings {
	return &TLThemeSettings{
		Data2: m,
	}
}

//  themeSettings#9c14984a flags:# base_theme:BaseTheme accent_color:int message_top_color:flags.0?int message_bottom_color:flags.0?int wallpaper:flags.1?WallPaper = ThemeSettings;
//
func (m *TLThemeSettings) To_ThemeSettings() *ThemeSettings {
	m.Data2.Cmd = Cmd_ThemeSettings
	return m.Data2
}

//  + TL_TopPeer
//

//  topPeer#edcdc05b peer:Peer rating:double = TopPeer;
//
func (m *TopPeer) To_TopPeer() *TLTopPeer {
	return &TLTopPeer{
		Data2: m,
	}
}

//  topPeer#edcdc05b peer:Peer rating:double = TopPeer;
//
func (m *TLTopPeer) To_TopPeer() *TopPeer {
	m.Data2.Cmd = Cmd_TopPeer
	return m.Data2
}

//  + TL_TopPeerCategoryBotsPM
//  + TL_TopPeerCategoryBotsInline
//  + TL_TopPeerCategoryCorrespondents
//  + TL_TopPeerCategoryGroups
//  + TL_TopPeerCategoryChannels
//  + TL_TopPeerCategoryPhoneCalls
//  + TL_TopPeerCategoryForwardUsers
//  + TL_TopPeerCategoryForwardChats
//

//  topPeerCategoryBotsPM#ab661b5b = TopPeerCategory;
//
func (m *TopPeerCategory) To_TopPeerCategoryBotsPM() *TLTopPeerCategoryBotsPM {
	return &TLTopPeerCategoryBotsPM{
		Data2: m,
	}
}

//  topPeerCategoryBotsInline#148677e2 = TopPeerCategory;
//
func (m *TopPeerCategory) To_TopPeerCategoryBotsInline() *TLTopPeerCategoryBotsInline {
	return &TLTopPeerCategoryBotsInline{
		Data2: m,
	}
}

//  topPeerCategoryCorrespondents#637b7ed = TopPeerCategory;
//
func (m *TopPeerCategory) To_TopPeerCategoryCorrespondents() *TLTopPeerCategoryCorrespondents {
	return &TLTopPeerCategoryCorrespondents{
		Data2: m,
	}
}

//  topPeerCategoryGroups#bd17a14a = TopPeerCategory;
//
func (m *TopPeerCategory) To_TopPeerCategoryGroups() *TLTopPeerCategoryGroups {
	return &TLTopPeerCategoryGroups{
		Data2: m,
	}
}

//  topPeerCategoryChannels#161d9628 = TopPeerCategory;
//
func (m *TopPeerCategory) To_TopPeerCategoryChannels() *TLTopPeerCategoryChannels {
	return &TLTopPeerCategoryChannels{
		Data2: m,
	}
}

//  topPeerCategoryPhoneCalls#1e76a78c = TopPeerCategory;
//
func (m *TopPeerCategory) To_TopPeerCategoryPhoneCalls() *TLTopPeerCategoryPhoneCalls {
	return &TLTopPeerCategoryPhoneCalls{
		Data2: m,
	}
}

//  topPeerCategoryForwardUsers#a8406ca9 = TopPeerCategory;
//
func (m *TopPeerCategory) To_TopPeerCategoryForwardUsers() *TLTopPeerCategoryForwardUsers {
	return &TLTopPeerCategoryForwardUsers{
		Data2: m,
	}
}

//  topPeerCategoryForwardChats#fbeec0f0 = TopPeerCategory;
//
func (m *TopPeerCategory) To_TopPeerCategoryForwardChats() *TLTopPeerCategoryForwardChats {
	return &TLTopPeerCategoryForwardChats{
		Data2: m,
	}
}

//  topPeerCategoryBotsPM#ab661b5b = TopPeerCategory;
//
func (m *TLTopPeerCategoryBotsPM) To_TopPeerCategory() *TopPeerCategory {
	m.Data2.Cmd = Cmd_TopPeerCategoryBotsPM
	return m.Data2
}

//  topPeerCategoryBotsInline#148677e2 = TopPeerCategory;
//
func (m *TLTopPeerCategoryBotsInline) To_TopPeerCategory() *TopPeerCategory {
	m.Data2.Cmd = Cmd_TopPeerCategoryBotsInline
	return m.Data2
}

//  topPeerCategoryCorrespondents#637b7ed = TopPeerCategory;
//
func (m *TLTopPeerCategoryCorrespondents) To_TopPeerCategory() *TopPeerCategory {
	m.Data2.Cmd = Cmd_TopPeerCategoryCorrespondents
	return m.Data2
}

//  topPeerCategoryGroups#bd17a14a = TopPeerCategory;
//
func (m *TLTopPeerCategoryGroups) To_TopPeerCategory() *TopPeerCategory {
	m.Data2.Cmd = Cmd_TopPeerCategoryGroups
	return m.Data2
}

//  topPeerCategoryChannels#161d9628 = TopPeerCategory;
//
func (m *TLTopPeerCategoryChannels) To_TopPeerCategory() *TopPeerCategory {
	m.Data2.Cmd = Cmd_TopPeerCategoryChannels
	return m.Data2
}

//  topPeerCategoryPhoneCalls#1e76a78c = TopPeerCategory;
//
func (m *TLTopPeerCategoryPhoneCalls) To_TopPeerCategory() *TopPeerCategory {
	m.Data2.Cmd = Cmd_TopPeerCategoryPhoneCalls
	return m.Data2
}

//  topPeerCategoryForwardUsers#a8406ca9 = TopPeerCategory;
//
func (m *TLTopPeerCategoryForwardUsers) To_TopPeerCategory() *TopPeerCategory {
	m.Data2.Cmd = Cmd_TopPeerCategoryForwardUsers
	return m.Data2
}

//  topPeerCategoryForwardChats#fbeec0f0 = TopPeerCategory;
//
func (m *TLTopPeerCategoryForwardChats) To_TopPeerCategory() *TopPeerCategory {
	m.Data2.Cmd = Cmd_TopPeerCategoryForwardChats
	return m.Data2
}

//  + TL_TopPeerCategoryPeers
//

//  topPeerCategoryPeers#fb834291 category:TopPeerCategory count:int peers:Vector<TopPeer> = TopPeerCategoryPeers;
//
func (m *TopPeerCategoryPeers) To_TopPeerCategoryPeers() *TLTopPeerCategoryPeers {
	return &TLTopPeerCategoryPeers{
		Data2: m,
	}
}

//  topPeerCategoryPeers#fb834291 category:TopPeerCategory count:int peers:Vector<TopPeer> = TopPeerCategoryPeers;
//
func (m *TLTopPeerCategoryPeers) To_TopPeerCategoryPeers() *TopPeerCategoryPeers {
	m.Data2.Cmd = Cmd_TopPeerCategoryPeers
	return m.Data2
}

//  + TL_True
//

//  true#3fedd339 = True;
//
func (m *True) To_True() *TLTrue {
	return &TLTrue{
		Data2: m,
	}
}

//  true#3fedd339 = True;
//
func (m *TLTrue) To_True() *True {
	m.Data2.Cmd = Cmd_True
	return m.Data2
}

//  + TL_UpdateNewMessage
//  + TL_UpdateMessageID
//  + TL_UpdateDeleteMessages
//  + TL_UpdateUserTyping
//  + TL_UpdateChatUserTyping
//  + TL_UpdateChatParticipants
//  + TL_UpdateUserStatus
//  + TL_UpdateUserName
//  + TL_UpdateUserPhoto
//  + TL_UpdateContactRegistered
//  + TL_UpdateContactLink
//  + TL_UpdateNewEncryptedMessage
//  + TL_UpdateEncryptedChatTyping
//  + TL_UpdateEncryption
//  + TL_UpdateEncryptedMessagesRead
//  + TL_UpdateChatParticipantAdd
//  + TL_UpdateChatParticipantDelete
//  + TL_UpdateDcOptions
//  + TL_UpdateUserBlocked
//  + TL_UpdateNotifySettings
//  + TL_UpdateServiceNotification
//  + TL_UpdatePrivacy
//  + TL_UpdateUserPhone
//  + TL_UpdateReadHistoryInbox
//  + TL_UpdateReadHistoryOutbox
//  + TL_UpdateWebPage
//  + TL_UpdateReadMessagesContents
//  + TL_UpdateChannelTooLong
//  + TL_UpdateChannel
//  + TL_UpdateNewChannelMessage
//  + TL_UpdateReadChannelInbox
//  + TL_UpdateDeleteChannelMessages
//  + TL_UpdateChannelMessageViews
//  + TL_UpdateChatAdmins
//  + TL_UpdateChatParticipantAdmin
//  + TL_UpdateNewStickerSet
//  + TL_UpdateStickerSetsOrder
//  + TL_UpdateStickerSets
//  + TL_UpdateSavedGifs
//  + TL_UpdateBotInlineQuery
//  + TL_UpdateBotInlineSend
//  + TL_UpdateEditChannelMessage
//  + TL_UpdateChannelPinnedMessage
//  + TL_UpdateBotCallbackQuery
//  + TL_UpdateEditMessage
//  + TL_UpdateInlineBotCallbackQuery
//  + TL_UpdateReadChannelOutbox
//  + TL_UpdateDraftMessage
//  + TL_UpdateReadFeaturedStickers
//  + TL_UpdateRecentStickers
//  + TL_UpdateConfig
//  + TL_UpdatePtsChanged
//  + TL_UpdateChannelWebPage
//  + TL_UpdateDialogPinned
//  + TL_UpdatePinnedDialogs
//  + TL_UpdateBotWebhookJSON
//  + TL_UpdateBotWebhookJSONQuery
//  + TL_UpdateBotShippingQuery
//  + TL_UpdateBotPrecheckoutQuery
//  + TL_UpdatePhoneCall
//  + TL_UpdateLangPackTooLong
//  + TL_UpdateLangPack
//  + TL_UpdateFavedStickers
//  + TL_UpdateChannelReadMessagesContents
//  + TL_UpdateContactsReset
//  + TL_UpdateNewAuthorization
//  + TL_UpdateChannelGroup
//  + TL_UpdateChannelAvailableMessages
//  + TL_UpdateDialogUnreadMark
//  + TL_UpdateUserPinnedMessage
//  + TL_UpdateChatPinnedMessage
//  + TL_UpdateMessagePoll
//  + TL_UpdateChatDefaultBannedRights
//  + TL_UpdateFolderPeers
//  + TL_UpdatePeerSettings
//  + TL_UpdatePeerLocated
//  + TL_UpdateNewScheduledMessage
//  + TL_UpdateDeleteScheduledMessages
//  + TL_UpdateTheme
//  + TL_UpdateGeoLiveViewed
//  + TL_UpdateLoginToken
//  + TL_UpdateMessagePollVote
//  + TL_UpdateDialogFilter
//  + TL_UpdateDialogFilterOrder
//  + TL_UpdateDialogFilters
//  + TL_UpdatePhoneCallSignalingData
//  + TL_UpdateChannelParticipant
//  + TL_UpdateChannelMessageForwards
//  + TL_UpdateReadChannelDiscussionInbox
//  + TL_UpdateReadChannelDiscussionOutbox
//  + TL_UpdatePeerBlocked
//  + TL_UpdateChannelUserTyping
//  + TL_UpdatePinnedMessages
//  + TL_UpdatePinnedChannelMessages
//  + TL_UpdateChat
//  + TL_UpdateGroupCallParticipants
//  + TL_UpdateGroupCall
//

//  updateNewMessage#1f2b0afd message:Message pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateNewMessage() *TLUpdateNewMessage {
	return &TLUpdateNewMessage{
		Data2: m,
	}
}

//  updateMessageID#4e90bfd6 id:int random_id:long = Update;
//
func (m *Update) To_UpdateMessageID() *TLUpdateMessageID {
	return &TLUpdateMessageID{
		Data2: m,
	}
}

//  updateDeleteMessages#a20db0e5 messages:Vector<int> pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateDeleteMessages() *TLUpdateDeleteMessages {
	return &TLUpdateDeleteMessages{
		Data2: m,
	}
}

//  updateUserTyping#5c486927 user_id:int action:SendMessageAction = Update;
//
func (m *Update) To_UpdateUserTyping() *TLUpdateUserTyping {
	return &TLUpdateUserTyping{
		Data2: m,
	}
}

//  updateChatUserTyping#9a65ea1f chat_id:int user_id:int action:SendMessageAction = Update;
//
func (m *Update) To_UpdateChatUserTyping() *TLUpdateChatUserTyping {
	return &TLUpdateChatUserTyping{
		Data2: m,
	}
}

//  updateChatParticipants#7761198 participants:ChatParticipants = Update;
//
func (m *Update) To_UpdateChatParticipants() *TLUpdateChatParticipants {
	return &TLUpdateChatParticipants{
		Data2: m,
	}
}

//  updateUserStatus#1bfbd823 user_id:int status:UserStatus = Update;
//
func (m *Update) To_UpdateUserStatus() *TLUpdateUserStatus {
	return &TLUpdateUserStatus{
		Data2: m,
	}
}

//  updateUserName#a7332b73 user_id:int first_name:string last_name:string username:string = Update;
//
func (m *Update) To_UpdateUserName() *TLUpdateUserName {
	return &TLUpdateUserName{
		Data2: m,
	}
}

//  updateUserPhoto#95313b0c user_id:int date:int photo:UserProfilePhoto previous:Bool = Update;
//
func (m *Update) To_UpdateUserPhoto() *TLUpdateUserPhoto {
	return &TLUpdateUserPhoto{
		Data2: m,
	}
}

//  updateContactRegistered#2575bbb9 user_id:int date:int = Update;
//
func (m *Update) To_UpdateContactRegistered() *TLUpdateContactRegistered {
	return &TLUpdateContactRegistered{
		Data2: m,
	}
}

//  updateContactLink#9d2e67c5 user_id:int my_link:ContactLink foreign_link:ContactLink = Update;
//
func (m *Update) To_UpdateContactLink() *TLUpdateContactLink {
	return &TLUpdateContactLink{
		Data2: m,
	}
}

//  updateNewEncryptedMessage#12bcbd9a message:EncryptedMessage qts:int = Update;
//
func (m *Update) To_UpdateNewEncryptedMessage() *TLUpdateNewEncryptedMessage {
	return &TLUpdateNewEncryptedMessage{
		Data2: m,
	}
}

//  updateEncryptedChatTyping#1710f156 chat_id:int = Update;
//
func (m *Update) To_UpdateEncryptedChatTyping() *TLUpdateEncryptedChatTyping {
	return &TLUpdateEncryptedChatTyping{
		Data2: m,
	}
}

//  updateEncryption#b4a2e88d chat:EncryptedChat date:int = Update;
//
func (m *Update) To_UpdateEncryption() *TLUpdateEncryption {
	return &TLUpdateEncryption{
		Data2: m,
	}
}

//  updateEncryptedMessagesRead#38fe25b7 chat_id:int max_date:int date:int = Update;
//
func (m *Update) To_UpdateEncryptedMessagesRead() *TLUpdateEncryptedMessagesRead {
	return &TLUpdateEncryptedMessagesRead{
		Data2: m,
	}
}

//  updateChatParticipantAdd#ea4b0e5c chat_id:int user_id:int inviter_id:int date:int version:int = Update;
//
func (m *Update) To_UpdateChatParticipantAdd() *TLUpdateChatParticipantAdd {
	return &TLUpdateChatParticipantAdd{
		Data2: m,
	}
}

//  updateChatParticipantDelete#6e5f8c22 chat_id:int user_id:int version:int = Update;
//
func (m *Update) To_UpdateChatParticipantDelete() *TLUpdateChatParticipantDelete {
	return &TLUpdateChatParticipantDelete{
		Data2: m,
	}
}

//  updateDcOptions#8e5e9873 dc_options:Vector<DcOption> = Update;
//
func (m *Update) To_UpdateDcOptions() *TLUpdateDcOptions {
	return &TLUpdateDcOptions{
		Data2: m,
	}
}

//  updateUserBlocked#80ece81a user_id:int blocked:Bool = Update;
//
func (m *Update) To_UpdateUserBlocked() *TLUpdateUserBlocked {
	return &TLUpdateUserBlocked{
		Data2: m,
	}
}

//  updateNotifySettings#bec268ef peer:NotifyPeer notify_settings:PeerNotifySettings = Update;
//
func (m *Update) To_UpdateNotifySettings() *TLUpdateNotifySettings {
	return &TLUpdateNotifySettings{
		Data2: m,
	}
}

//  updateServiceNotification#ebe46819 flags:# popup:flags.0?true inbox_date:flags.1?int type:string message:string media:MessageMedia entities:Vector<MessageEntity> = Update;
//  updateServiceNotification#382dd3e4 type:string message:string media:MessageMedia popup:Bool = Update;
//
func (m *Update) To_UpdateServiceNotification() *TLUpdateServiceNotification {
	return &TLUpdateServiceNotification{
		Data2: m,
	}
}

//  updatePrivacy#ee3b272a key:PrivacyKey rules:Vector<PrivacyRule> = Update;
//
func (m *Update) To_UpdatePrivacy() *TLUpdatePrivacy {
	return &TLUpdatePrivacy{
		Data2: m,
	}
}

//  updateUserPhone#12b9417b user_id:int phone:string = Update;
//
func (m *Update) To_UpdateUserPhone() *TLUpdateUserPhone {
	return &TLUpdateUserPhone{
		Data2: m,
	}
}

//  updateReadHistoryInbox#9961fd5c peer:Peer max_id:int pts:int pts_count:int = Update;
//  updateReadHistoryInbox#9c974fdf flags:# folder_id:flags.0?int peer:Peer max_id:int still_unread_count:int pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateReadHistoryInbox() *TLUpdateReadHistoryInbox {
	return &TLUpdateReadHistoryInbox{
		Data2: m,
	}
}

//  updateReadHistoryOutbox#2f2f21bf peer:Peer max_id:int pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateReadHistoryOutbox() *TLUpdateReadHistoryOutbox {
	return &TLUpdateReadHistoryOutbox{
		Data2: m,
	}
}

//  updateWebPage#7f891213 webpage:WebPage pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateWebPage() *TLUpdateWebPage {
	return &TLUpdateWebPage{
		Data2: m,
	}
}

//  updateReadMessagesContents#68c13933 messages:Vector<int> pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateReadMessagesContents() *TLUpdateReadMessagesContents {
	return &TLUpdateReadMessagesContents{
		Data2: m,
	}
}

//  updateChannelTooLong#eb0467fb flags:# channel_id:int pts:flags.0?int = Update;
//
func (m *Update) To_UpdateChannelTooLong() *TLUpdateChannelTooLong {
	return &TLUpdateChannelTooLong{
		Data2: m,
	}
}

//  updateChannel#b6d45656 channel_id:int = Update;
//
func (m *Update) To_UpdateChannel() *TLUpdateChannel {
	return &TLUpdateChannel{
		Data2: m,
	}
}

//  updateNewChannelMessage#62ba04d9 message:Message pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateNewChannelMessage() *TLUpdateNewChannelMessage {
	return &TLUpdateNewChannelMessage{
		Data2: m,
	}
}

//  updateReadChannelInbox#4214f37f channel_id:int max_id:int = Update;
//  updateReadChannelInbox#330b5424 flags:# folder_id:flags.0?int channel_id:int max_id:int still_unread_count:int pts:int = Update;
//
func (m *Update) To_UpdateReadChannelInbox() *TLUpdateReadChannelInbox {
	return &TLUpdateReadChannelInbox{
		Data2: m,
	}
}

//  updateDeleteChannelMessages#c37521c9 channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateDeleteChannelMessages() *TLUpdateDeleteChannelMessages {
	return &TLUpdateDeleteChannelMessages{
		Data2: m,
	}
}

//  updateChannelMessageViews#98a12b4b channel_id:int id:int views:int = Update;
//
func (m *Update) To_UpdateChannelMessageViews() *TLUpdateChannelMessageViews {
	return &TLUpdateChannelMessageViews{
		Data2: m,
	}
}

//  updateChatAdmins#6e947941 chat_id:int enabled:Bool version:int = Update;
//
func (m *Update) To_UpdateChatAdmins() *TLUpdateChatAdmins {
	return &TLUpdateChatAdmins{
		Data2: m,
	}
}

//  updateChatParticipantAdmin#b6901959 chat_id:int user_id:int is_admin:Bool version:int = Update;
//
func (m *Update) To_UpdateChatParticipantAdmin() *TLUpdateChatParticipantAdmin {
	return &TLUpdateChatParticipantAdmin{
		Data2: m,
	}
}

//  updateNewStickerSet#688a30aa stickerset:messages.StickerSet = Update;
//
func (m *Update) To_UpdateNewStickerSet() *TLUpdateNewStickerSet {
	return &TLUpdateNewStickerSet{
		Data2: m,
	}
}

//  updateStickerSetsOrder#bb2d201 flags:# masks:flags.0?true order:Vector<long> = Update;
//  updateStickerSetsOrder#f0dfb451 order:Vector<long> = Update;
//
func (m *Update) To_UpdateStickerSetsOrder() *TLUpdateStickerSetsOrder {
	return &TLUpdateStickerSetsOrder{
		Data2: m,
	}
}

//  updateStickerSets#43ae3dec = Update;
//
func (m *Update) To_UpdateStickerSets() *TLUpdateStickerSets {
	return &TLUpdateStickerSets{
		Data2: m,
	}
}

//  updateSavedGifs#9375341e = Update;
//
func (m *Update) To_UpdateSavedGifs() *TLUpdateSavedGifs {
	return &TLUpdateSavedGifs{
		Data2: m,
	}
}

//  updateBotInlineQuery#54826690 flags:# query_id:long user_id:int query:string geo:flags.0?GeoPoint offset:string = Update;
//  updateBotInlineQuery#3f2038db flags:# query_id:long user_id:int query:string geo:flags.0?GeoPoint peer_type:flags.1?InlineQueryPeerType offset:string = Update;
//
func (m *Update) To_UpdateBotInlineQuery() *TLUpdateBotInlineQuery {
	return &TLUpdateBotInlineQuery{
		Data2: m,
	}
}

//  updateBotInlineSend#e48f964 flags:# user_id:int query:string geo:flags.0?GeoPoint id:string msg_id:flags.1?InputBotInlineMessageID = Update;
//
func (m *Update) To_UpdateBotInlineSend() *TLUpdateBotInlineSend {
	return &TLUpdateBotInlineSend{
		Data2: m,
	}
}

//  updateEditChannelMessage#1b3f4df7 message:Message pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateEditChannelMessage() *TLUpdateEditChannelMessage {
	return &TLUpdateEditChannelMessage{
		Data2: m,
	}
}

//  updateChannelPinnedMessage#98592475 channel_id:int id:int = Update;
//
func (m *Update) To_UpdateChannelPinnedMessage() *TLUpdateChannelPinnedMessage {
	return &TLUpdateChannelPinnedMessage{
		Data2: m,
	}
}

//  updateBotCallbackQuery#e73547e1 flags:# query_id:long user_id:int peer:Peer msg_id:int chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;
//  updateBotCallbackQuery#a68c688c query_id:long user_id:int peer:Peer msg_id:int data:bytes = Update;
//
func (m *Update) To_UpdateBotCallbackQuery() *TLUpdateBotCallbackQuery {
	return &TLUpdateBotCallbackQuery{
		Data2: m,
	}
}

//  updateEditMessage#e40370a3 message:Message pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateEditMessage() *TLUpdateEditMessage {
	return &TLUpdateEditMessage{
		Data2: m,
	}
}

//  updateInlineBotCallbackQuery#f9d27a5a flags:# query_id:long user_id:int msg_id:InputBotInlineMessageID chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;
//  updateInlineBotCallbackQuery#2cbd95af query_id:long user_id:int msg_id:InputBotInlineMessageID data:bytes = Update;
//
func (m *Update) To_UpdateInlineBotCallbackQuery() *TLUpdateInlineBotCallbackQuery {
	return &TLUpdateInlineBotCallbackQuery{
		Data2: m,
	}
}

//  updateReadChannelOutbox#25d6c9c7 channel_id:int max_id:int = Update;
//
func (m *Update) To_UpdateReadChannelOutbox() *TLUpdateReadChannelOutbox {
	return &TLUpdateReadChannelOutbox{
		Data2: m,
	}
}

//  updateDraftMessage#ee2bb969 peer:Peer draft:DraftMessage = Update;
//
func (m *Update) To_UpdateDraftMessage() *TLUpdateDraftMessage {
	return &TLUpdateDraftMessage{
		Data2: m,
	}
}

//  updateReadFeaturedStickers#571d2742 = Update;
//
func (m *Update) To_UpdateReadFeaturedStickers() *TLUpdateReadFeaturedStickers {
	return &TLUpdateReadFeaturedStickers{
		Data2: m,
	}
}

//  updateRecentStickers#9a422c20 = Update;
//
func (m *Update) To_UpdateRecentStickers() *TLUpdateRecentStickers {
	return &TLUpdateRecentStickers{
		Data2: m,
	}
}

//  updateConfig#a229dd06 = Update;
//
func (m *Update) To_UpdateConfig() *TLUpdateConfig {
	return &TLUpdateConfig{
		Data2: m,
	}
}

//  updatePtsChanged#3354678f = Update;
//
func (m *Update) To_UpdatePtsChanged() *TLUpdatePtsChanged {
	return &TLUpdatePtsChanged{
		Data2: m,
	}
}

//  updateChannelWebPage#40771900 channel_id:int webpage:WebPage pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateChannelWebPage() *TLUpdateChannelWebPage {
	return &TLUpdateChannelWebPage{
		Data2: m,
	}
}

//  updateDialogPinned#d711a2cc flags:# pinned:flags.0?true peer:Peer = Update;
//  updateDialogPinned#19d27f3c flags:# pinned:flags.0?true peer:DialogPeer = Update;
//  updateDialogPinned#6e6fe51c flags:# pinned:flags.0?true folder_id:flags.1?int peer:DialogPeer = Update;
//
func (m *Update) To_UpdateDialogPinned() *TLUpdateDialogPinned {
	return &TLUpdateDialogPinned{
		Data2: m,
	}
}

//  updatePinnedDialogs#d8caf68d flags:# order:flags.0?Vector<Peer> = Update;
//  updatePinnedDialogs#ea4cb65b flags:# order:flags.0?Vector<DialogPeer> = Update;
//  updatePinnedDialogs#fa0f3ca2 flags:# folder_id:flags.1?int order:flags.0?Vector<DialogPeer> = Update;
//
func (m *Update) To_UpdatePinnedDialogs() *TLUpdatePinnedDialogs {
	return &TLUpdatePinnedDialogs{
		Data2: m,
	}
}

//  updateBotWebhookJSON#8317c0c3 data:DataJSON = Update;
//
func (m *Update) To_UpdateBotWebhookJSON() *TLUpdateBotWebhookJSON {
	return &TLUpdateBotWebhookJSON{
		Data2: m,
	}
}

//  updateBotWebhookJSONQuery#9b9240a6 query_id:long data:DataJSON timeout:int = Update;
//
func (m *Update) To_UpdateBotWebhookJSONQuery() *TLUpdateBotWebhookJSONQuery {
	return &TLUpdateBotWebhookJSONQuery{
		Data2: m,
	}
}

//  updateBotShippingQuery#e0cdc940 query_id:long user_id:int payload:bytes shipping_address:PostAddress = Update;
//
func (m *Update) To_UpdateBotShippingQuery() *TLUpdateBotShippingQuery {
	return &TLUpdateBotShippingQuery{
		Data2: m,
	}
}

//  updateBotPrecheckoutQuery#5d2f3aa9 flags:# query_id:long user_id:int payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string currency:string total_amount:long = Update;
//
func (m *Update) To_UpdateBotPrecheckoutQuery() *TLUpdateBotPrecheckoutQuery {
	return &TLUpdateBotPrecheckoutQuery{
		Data2: m,
	}
}

//  updatePhoneCall#ab0f6b1e phone_call:PhoneCall = Update;
//
func (m *Update) To_UpdatePhoneCall() *TLUpdatePhoneCall {
	return &TLUpdatePhoneCall{
		Data2: m,
	}
}

//  updateLangPackTooLong#10c2404b = Update;
//  updateLangPackTooLong#46560264 lang_code:string = Update;
//
func (m *Update) To_UpdateLangPackTooLong() *TLUpdateLangPackTooLong {
	return &TLUpdateLangPackTooLong{
		Data2: m,
	}
}

//  updateLangPack#56022f4d difference:LangPackDifference = Update;
//
func (m *Update) To_UpdateLangPack() *TLUpdateLangPack {
	return &TLUpdateLangPack{
		Data2: m,
	}
}

//  updateFavedStickers#e511996d = Update;
//
func (m *Update) To_UpdateFavedStickers() *TLUpdateFavedStickers {
	return &TLUpdateFavedStickers{
		Data2: m,
	}
}

//  updateChannelReadMessagesContents#89893b45 channel_id:int messages:Vector<int> = Update;
//
func (m *Update) To_UpdateChannelReadMessagesContents() *TLUpdateChannelReadMessagesContents {
	return &TLUpdateChannelReadMessagesContents{
		Data2: m,
	}
}

//  updateContactsReset#7084a7be = Update;
//
func (m *Update) To_UpdateContactsReset() *TLUpdateContactsReset {
	return &TLUpdateContactsReset{
		Data2: m,
	}
}

//  updateNewAuthorization#8f06529a auth_key_id:long date:int device:string location:string = Update;
//
func (m *Update) To_UpdateNewAuthorization() *TLUpdateNewAuthorization {
	return &TLUpdateNewAuthorization{
		Data2: m,
	}
}

//  updateChannelGroup#c36c1e3c channel_id:int group:MessageGroup = Update;
//
func (m *Update) To_UpdateChannelGroup() *TLUpdateChannelGroup {
	return &TLUpdateChannelGroup{
		Data2: m,
	}
}

//  updateChannelAvailableMessages#70db6837 channel_id:int available_min_id:int = Update;
//
func (m *Update) To_UpdateChannelAvailableMessages() *TLUpdateChannelAvailableMessages {
	return &TLUpdateChannelAvailableMessages{
		Data2: m,
	}
}

//  updateDialogUnreadMark#e16459c3 flags:# unread:flags.0?true peer:DialogPeer = Update;
//
func (m *Update) To_UpdateDialogUnreadMark() *TLUpdateDialogUnreadMark {
	return &TLUpdateDialogUnreadMark{
		Data2: m,
	}
}

//  updateUserPinnedMessage#4c43da18 user_id:int id:int = Update;
//
func (m *Update) To_UpdateUserPinnedMessage() *TLUpdateUserPinnedMessage {
	return &TLUpdateUserPinnedMessage{
		Data2: m,
	}
}

//  updateChatPinnedMessage#22893b26 chat_id:int id:int = Update;
//  updateChatPinnedMessage#e10db349 chat_id:int id:int version:int = Update;
//
func (m *Update) To_UpdateChatPinnedMessage() *TLUpdateChatPinnedMessage {
	return &TLUpdateChatPinnedMessage{
		Data2: m,
	}
}

//  updateMessagePoll#aca1657b flags:# poll_id:long poll:flags.0?Poll results:PollResults = Update;
//
func (m *Update) To_UpdateMessagePoll() *TLUpdateMessagePoll {
	return &TLUpdateMessagePoll{
		Data2: m,
	}
}

//  updateChatDefaultBannedRights#54c01850 peer:Peer default_banned_rights:ChatBannedRights version:int = Update;
//
func (m *Update) To_UpdateChatDefaultBannedRights() *TLUpdateChatDefaultBannedRights {
	return &TLUpdateChatDefaultBannedRights{
		Data2: m,
	}
}

//  updateFolderPeers#19360dc0 folder_peers:Vector<FolderPeer> pts:int pts_count:int = Update;
//
func (m *Update) To_UpdateFolderPeers() *TLUpdateFolderPeers {
	return &TLUpdateFolderPeers{
		Data2: m,
	}
}

//  updatePeerSettings#6a7e7366 peer:Peer settings:PeerSettings = Update;
//
func (m *Update) To_UpdatePeerSettings() *TLUpdatePeerSettings {
	return &TLUpdatePeerSettings{
		Data2: m,
	}
}

//  updatePeerLocated#b4afcfb0 peers:Vector<PeerLocated> = Update;
//
func (m *Update) To_UpdatePeerLocated() *TLUpdatePeerLocated {
	return &TLUpdatePeerLocated{
		Data2: m,
	}
}

//  updateNewScheduledMessage#39a51dfb message:Message = Update;
//
func (m *Update) To_UpdateNewScheduledMessage() *TLUpdateNewScheduledMessage {
	return &TLUpdateNewScheduledMessage{
		Data2: m,
	}
}

//  updateDeleteScheduledMessages#90866cee peer:Peer messages:Vector<int> = Update;
//
func (m *Update) To_UpdateDeleteScheduledMessages() *TLUpdateDeleteScheduledMessages {
	return &TLUpdateDeleteScheduledMessages{
		Data2: m,
	}
}

//  updateTheme#8216fba3 theme:Theme = Update;
//
func (m *Update) To_UpdateTheme() *TLUpdateTheme {
	return &TLUpdateTheme{
		Data2: m,
	}
}

//  updateGeoLiveViewed#871fb939 peer:Peer msg_id:int = Update;
//
func (m *Update) To_UpdateGeoLiveViewed() *TLUpdateGeoLiveViewed {
	return &TLUpdateGeoLiveViewed{
		Data2: m,
	}
}

//  updateLoginToken#564fe691 = Update;
//
func (m *Update) To_UpdateLoginToken() *TLUpdateLoginToken {
	return &TLUpdateLoginToken{
		Data2: m,
	}
}

//  updateMessagePollVote#42f88f2c poll_id:long user_id:int options:Vector<bytes> = Update;
//
func (m *Update) To_UpdateMessagePollVote() *TLUpdateMessagePollVote {
	return &TLUpdateMessagePollVote{
		Data2: m,
	}
}

//  updateDialogFilter#26ffde7d flags:# id:int filter:flags.0?DialogFilter = Update;
//
func (m *Update) To_UpdateDialogFilter() *TLUpdateDialogFilter {
	return &TLUpdateDialogFilter{
		Data2: m,
	}
}

//  updateDialogFilterOrder#a5d72105 order:Vector<int> = Update;
//
func (m *Update) To_UpdateDialogFilterOrder() *TLUpdateDialogFilterOrder {
	return &TLUpdateDialogFilterOrder{
		Data2: m,
	}
}

//  updateDialogFilters#3504914f = Update;
//
func (m *Update) To_UpdateDialogFilters() *TLUpdateDialogFilters {
	return &TLUpdateDialogFilters{
		Data2: m,
	}
}

//  updatePhoneCallSignalingData#2661bf09 phone_call_id:long data:bytes = Update;
//
func (m *Update) To_UpdatePhoneCallSignalingData() *TLUpdatePhoneCallSignalingData {
	return &TLUpdatePhoneCallSignalingData{
		Data2: m,
	}
}

//  updateChannelParticipant#65d2b464 flags:# channel_id:int date:int user_id:int prev_participant:flags.0?ChannelParticipant new_participant:flags.1?ChannelParticipant qts:int = Update;
//
func (m *Update) To_UpdateChannelParticipant() *TLUpdateChannelParticipant {
	return &TLUpdateChannelParticipant{
		Data2: m,
	}
}

//  updateChannelMessageForwards#6e8a84df channel_id:int id:int forwards:int = Update;
//
func (m *Update) To_UpdateChannelMessageForwards() *TLUpdateChannelMessageForwards {
	return &TLUpdateChannelMessageForwards{
		Data2: m,
	}
}

//  updateReadChannelDiscussionInbox#1cc7de54 flags:# channel_id:int top_msg_id:int read_max_id:int broadcast_id:flags.0?int broadcast_post:flags.0?int = Update;
//
func (m *Update) To_UpdateReadChannelDiscussionInbox() *TLUpdateReadChannelDiscussionInbox {
	return &TLUpdateReadChannelDiscussionInbox{
		Data2: m,
	}
}

//  updateReadChannelDiscussionOutbox#4638a26c channel_id:int top_msg_id:int read_max_id:int = Update;
//
func (m *Update) To_UpdateReadChannelDiscussionOutbox() *TLUpdateReadChannelDiscussionOutbox {
	return &TLUpdateReadChannelDiscussionOutbox{
		Data2: m,
	}
}

//  updatePeerBlocked#246a4b22 peer_id:Peer blocked:Bool = Update;
//
func (m *Update) To_UpdatePeerBlocked() *TLUpdatePeerBlocked {
	return &TLUpdatePeerBlocked{
		Data2: m,
	}
}

//  updateChannelUserTyping#ff2abe9f flags:# channel_id:int top_msg_id:flags.0?int user_id:int action:SendMessageAction = Update;
//
func (m *Update) To_UpdateChannelUserTyping() *TLUpdateChannelUserTyping {
	return &TLUpdateChannelUserTyping{
		Data2: m,
	}
}

//  updatePinnedMessages#ed85eab5 flags:# pinned:flags.0?true peer:Peer messages:Vector<int> pts:int pts_count:int = Update;
//
func (m *Update) To_UpdatePinnedMessages() *TLUpdatePinnedMessages {
	return &TLUpdatePinnedMessages{
		Data2: m,
	}
}

//  updatePinnedChannelMessages#8588878b flags:# pinned:flags.0?true channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
//
func (m *Update) To_UpdatePinnedChannelMessages() *TLUpdatePinnedChannelMessages {
	return &TLUpdatePinnedChannelMessages{
		Data2: m,
	}
}

//  updateChat#1330a196 chat_id:int = Update;
//
func (m *Update) To_UpdateChat() *TLUpdateChat {
	return &TLUpdateChat{
		Data2: m,
	}
}

//  updateGroupCallParticipants#f2ebdb4e call:InputGroupCall participants:Vector<GroupCallParticipant> version:int = Update;
//
func (m *Update) To_UpdateGroupCallParticipants() *TLUpdateGroupCallParticipants {
	return &TLUpdateGroupCallParticipants{
		Data2: m,
	}
}

//  updateGroupCall#a45eb99b chat_id:int call:GroupCall = Update;
//
func (m *Update) To_UpdateGroupCall() *TLUpdateGroupCall {
	return &TLUpdateGroupCall{
		Data2: m,
	}
}

//  updateNewMessage#1f2b0afd message:Message pts:int pts_count:int = Update;
//
func (m *TLUpdateNewMessage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateNewMessage
	return m.Data2
}

//  updateMessageID#4e90bfd6 id:int random_id:long = Update;
//
func (m *TLUpdateMessageID) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateMessageID
	return m.Data2
}

//  updateDeleteMessages#a20db0e5 messages:Vector<int> pts:int pts_count:int = Update;
//
func (m *TLUpdateDeleteMessages) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateDeleteMessages
	return m.Data2
}

//  updateUserTyping#5c486927 user_id:int action:SendMessageAction = Update;
//
func (m *TLUpdateUserTyping) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateUserTyping
	return m.Data2
}

//  updateChatUserTyping#9a65ea1f chat_id:int user_id:int action:SendMessageAction = Update;
//
func (m *TLUpdateChatUserTyping) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChatUserTyping
	return m.Data2
}

//  updateChatParticipants#7761198 participants:ChatParticipants = Update;
//
func (m *TLUpdateChatParticipants) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChatParticipants
	return m.Data2
}

//  updateUserStatus#1bfbd823 user_id:int status:UserStatus = Update;
//
func (m *TLUpdateUserStatus) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateUserStatus
	return m.Data2
}

//  updateUserName#a7332b73 user_id:int first_name:string last_name:string username:string = Update;
//
func (m *TLUpdateUserName) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateUserName
	return m.Data2
}

//  updateUserPhoto#95313b0c user_id:int date:int photo:UserProfilePhoto previous:Bool = Update;
//
func (m *TLUpdateUserPhoto) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateUserPhoto
	return m.Data2
}

//  updateContactRegistered#2575bbb9 user_id:int date:int = Update;
//
func (m *TLUpdateContactRegistered) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateContactRegistered
	return m.Data2
}

//  updateContactLink#9d2e67c5 user_id:int my_link:ContactLink foreign_link:ContactLink = Update;
//
func (m *TLUpdateContactLink) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateContactLink
	return m.Data2
}

//  updateNewEncryptedMessage#12bcbd9a message:EncryptedMessage qts:int = Update;
//
func (m *TLUpdateNewEncryptedMessage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateNewEncryptedMessage
	return m.Data2
}

//  updateEncryptedChatTyping#1710f156 chat_id:int = Update;
//
func (m *TLUpdateEncryptedChatTyping) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateEncryptedChatTyping
	return m.Data2
}

//  updateEncryption#b4a2e88d chat:EncryptedChat date:int = Update;
//
func (m *TLUpdateEncryption) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateEncryption
	return m.Data2
}

//  updateEncryptedMessagesRead#38fe25b7 chat_id:int max_date:int date:int = Update;
//
func (m *TLUpdateEncryptedMessagesRead) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateEncryptedMessagesRead
	return m.Data2
}

//  updateChatParticipantAdd#ea4b0e5c chat_id:int user_id:int inviter_id:int date:int version:int = Update;
//
func (m *TLUpdateChatParticipantAdd) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChatParticipantAdd
	return m.Data2
}

//  updateChatParticipantDelete#6e5f8c22 chat_id:int user_id:int version:int = Update;
//
func (m *TLUpdateChatParticipantDelete) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChatParticipantDelete
	return m.Data2
}

//  updateDcOptions#8e5e9873 dc_options:Vector<DcOption> = Update;
//
func (m *TLUpdateDcOptions) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateDcOptions
	return m.Data2
}

//  updateUserBlocked#80ece81a user_id:int blocked:Bool = Update;
//
func (m *TLUpdateUserBlocked) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateUserBlocked
	return m.Data2
}

//  updateNotifySettings#bec268ef peer:NotifyPeer notify_settings:PeerNotifySettings = Update;
//
func (m *TLUpdateNotifySettings) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateNotifySettings
	return m.Data2
}

//  updateServiceNotification#ebe46819 flags:# popup:flags.0?true inbox_date:flags.1?int type:string message:string media:MessageMedia entities:Vector<MessageEntity> = Update;
//  updateServiceNotification#382dd3e4 type:string message:string media:MessageMedia popup:Bool = Update;
//
func (m *TLUpdateServiceNotification) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateServiceNotification
	return m.Data2
}

//  updatePrivacy#ee3b272a key:PrivacyKey rules:Vector<PrivacyRule> = Update;
//
func (m *TLUpdatePrivacy) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdatePrivacy
	return m.Data2
}

//  updateUserPhone#12b9417b user_id:int phone:string = Update;
//
func (m *TLUpdateUserPhone) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateUserPhone
	return m.Data2
}

//  updateReadHistoryInbox#9961fd5c peer:Peer max_id:int pts:int pts_count:int = Update;
//  updateReadHistoryInbox#9c974fdf flags:# folder_id:flags.0?int peer:Peer max_id:int still_unread_count:int pts:int pts_count:int = Update;
//
func (m *TLUpdateReadHistoryInbox) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateReadHistoryInbox
	return m.Data2
}

//  updateReadHistoryOutbox#2f2f21bf peer:Peer max_id:int pts:int pts_count:int = Update;
//
func (m *TLUpdateReadHistoryOutbox) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateReadHistoryOutbox
	return m.Data2
}

//  updateWebPage#7f891213 webpage:WebPage pts:int pts_count:int = Update;
//
func (m *TLUpdateWebPage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateWebPage
	return m.Data2
}

//  updateReadMessagesContents#68c13933 messages:Vector<int> pts:int pts_count:int = Update;
//
func (m *TLUpdateReadMessagesContents) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateReadMessagesContents
	return m.Data2
}

//  updateChannelTooLong#eb0467fb flags:# channel_id:int pts:flags.0?int = Update;
//
func (m *TLUpdateChannelTooLong) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChannelTooLong
	return m.Data2
}

//  updateChannel#b6d45656 channel_id:int = Update;
//
func (m *TLUpdateChannel) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChannel
	return m.Data2
}

//  updateNewChannelMessage#62ba04d9 message:Message pts:int pts_count:int = Update;
//
func (m *TLUpdateNewChannelMessage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateNewChannelMessage
	return m.Data2
}

//  updateReadChannelInbox#4214f37f channel_id:int max_id:int = Update;
//  updateReadChannelInbox#330b5424 flags:# folder_id:flags.0?int channel_id:int max_id:int still_unread_count:int pts:int = Update;
//
func (m *TLUpdateReadChannelInbox) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateReadChannelInbox
	return m.Data2
}

//  updateDeleteChannelMessages#c37521c9 channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
//
func (m *TLUpdateDeleteChannelMessages) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateDeleteChannelMessages
	return m.Data2
}

//  updateChannelMessageViews#98a12b4b channel_id:int id:int views:int = Update;
//
func (m *TLUpdateChannelMessageViews) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChannelMessageViews
	return m.Data2
}

//  updateChatAdmins#6e947941 chat_id:int enabled:Bool version:int = Update;
//
func (m *TLUpdateChatAdmins) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChatAdmins
	return m.Data2
}

//  updateChatParticipantAdmin#b6901959 chat_id:int user_id:int is_admin:Bool version:int = Update;
//
func (m *TLUpdateChatParticipantAdmin) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChatParticipantAdmin
	return m.Data2
}

//  updateNewStickerSet#688a30aa stickerset:messages.StickerSet = Update;
//
func (m *TLUpdateNewStickerSet) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateNewStickerSet
	return m.Data2
}

//  updateStickerSetsOrder#bb2d201 flags:# masks:flags.0?true order:Vector<long> = Update;
//  updateStickerSetsOrder#f0dfb451 order:Vector<long> = Update;
//
func (m *TLUpdateStickerSetsOrder) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateStickerSetsOrder
	return m.Data2
}

//  updateStickerSets#43ae3dec = Update;
//
func (m *TLUpdateStickerSets) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateStickerSets
	return m.Data2
}

//  updateSavedGifs#9375341e = Update;
//
func (m *TLUpdateSavedGifs) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateSavedGifs
	return m.Data2
}

//  updateBotInlineQuery#54826690 flags:# query_id:long user_id:int query:string geo:flags.0?GeoPoint offset:string = Update;
//  updateBotInlineQuery#3f2038db flags:# query_id:long user_id:int query:string geo:flags.0?GeoPoint peer_type:flags.1?InlineQueryPeerType offset:string = Update;
//
func (m *TLUpdateBotInlineQuery) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateBotInlineQuery
	return m.Data2
}

//  updateBotInlineSend#e48f964 flags:# user_id:int query:string geo:flags.0?GeoPoint id:string msg_id:flags.1?InputBotInlineMessageID = Update;
//
func (m *TLUpdateBotInlineSend) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateBotInlineSend
	return m.Data2
}

//  updateEditChannelMessage#1b3f4df7 message:Message pts:int pts_count:int = Update;
//
func (m *TLUpdateEditChannelMessage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateEditChannelMessage
	return m.Data2
}

//  updateChannelPinnedMessage#98592475 channel_id:int id:int = Update;
//
func (m *TLUpdateChannelPinnedMessage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChannelPinnedMessage
	return m.Data2
}

//  updateBotCallbackQuery#e73547e1 flags:# query_id:long user_id:int peer:Peer msg_id:int chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;
//  updateBotCallbackQuery#a68c688c query_id:long user_id:int peer:Peer msg_id:int data:bytes = Update;
//
func (m *TLUpdateBotCallbackQuery) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateBotCallbackQuery
	return m.Data2
}

//  updateEditMessage#e40370a3 message:Message pts:int pts_count:int = Update;
//
func (m *TLUpdateEditMessage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateEditMessage
	return m.Data2
}

//  updateInlineBotCallbackQuery#f9d27a5a flags:# query_id:long user_id:int msg_id:InputBotInlineMessageID chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;
//  updateInlineBotCallbackQuery#2cbd95af query_id:long user_id:int msg_id:InputBotInlineMessageID data:bytes = Update;
//
func (m *TLUpdateInlineBotCallbackQuery) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateInlineBotCallbackQuery
	return m.Data2
}

//  updateReadChannelOutbox#25d6c9c7 channel_id:int max_id:int = Update;
//
func (m *TLUpdateReadChannelOutbox) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateReadChannelOutbox
	return m.Data2
}

//  updateDraftMessage#ee2bb969 peer:Peer draft:DraftMessage = Update;
//
func (m *TLUpdateDraftMessage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateDraftMessage
	return m.Data2
}

//  updateReadFeaturedStickers#571d2742 = Update;
//
func (m *TLUpdateReadFeaturedStickers) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateReadFeaturedStickers
	return m.Data2
}

//  updateRecentStickers#9a422c20 = Update;
//
func (m *TLUpdateRecentStickers) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateRecentStickers
	return m.Data2
}

//  updateConfig#a229dd06 = Update;
//
func (m *TLUpdateConfig) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateConfig
	return m.Data2
}

//  updatePtsChanged#3354678f = Update;
//
func (m *TLUpdatePtsChanged) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdatePtsChanged
	return m.Data2
}

//  updateChannelWebPage#40771900 channel_id:int webpage:WebPage pts:int pts_count:int = Update;
//
func (m *TLUpdateChannelWebPage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChannelWebPage
	return m.Data2
}

//  updateDialogPinned#d711a2cc flags:# pinned:flags.0?true peer:Peer = Update;
//  updateDialogPinned#19d27f3c flags:# pinned:flags.0?true peer:DialogPeer = Update;
//  updateDialogPinned#6e6fe51c flags:# pinned:flags.0?true folder_id:flags.1?int peer:DialogPeer = Update;
//
func (m *TLUpdateDialogPinned) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateDialogPinned
	return m.Data2
}

//  updatePinnedDialogs#d8caf68d flags:# order:flags.0?Vector<Peer> = Update;
//  updatePinnedDialogs#ea4cb65b flags:# order:flags.0?Vector<DialogPeer> = Update;
//  updatePinnedDialogs#fa0f3ca2 flags:# folder_id:flags.1?int order:flags.0?Vector<DialogPeer> = Update;
//
func (m *TLUpdatePinnedDialogs) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdatePinnedDialogs
	return m.Data2
}

//  updateBotWebhookJSON#8317c0c3 data:DataJSON = Update;
//
func (m *TLUpdateBotWebhookJSON) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateBotWebhookJSON
	return m.Data2
}

//  updateBotWebhookJSONQuery#9b9240a6 query_id:long data:DataJSON timeout:int = Update;
//
func (m *TLUpdateBotWebhookJSONQuery) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateBotWebhookJSONQuery
	return m.Data2
}

//  updateBotShippingQuery#e0cdc940 query_id:long user_id:int payload:bytes shipping_address:PostAddress = Update;
//
func (m *TLUpdateBotShippingQuery) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateBotShippingQuery
	return m.Data2
}

//  updateBotPrecheckoutQuery#5d2f3aa9 flags:# query_id:long user_id:int payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string currency:string total_amount:long = Update;
//
func (m *TLUpdateBotPrecheckoutQuery) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateBotPrecheckoutQuery
	return m.Data2
}

//  updatePhoneCall#ab0f6b1e phone_call:PhoneCall = Update;
//
func (m *TLUpdatePhoneCall) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdatePhoneCall
	return m.Data2
}

//  updateLangPackTooLong#10c2404b = Update;
//  updateLangPackTooLong#46560264 lang_code:string = Update;
//
func (m *TLUpdateLangPackTooLong) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateLangPackTooLong
	return m.Data2
}

//  updateLangPack#56022f4d difference:LangPackDifference = Update;
//
func (m *TLUpdateLangPack) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateLangPack
	return m.Data2
}

//  updateFavedStickers#e511996d = Update;
//
func (m *TLUpdateFavedStickers) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateFavedStickers
	return m.Data2
}

//  updateChannelReadMessagesContents#89893b45 channel_id:int messages:Vector<int> = Update;
//
func (m *TLUpdateChannelReadMessagesContents) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChannelReadMessagesContents
	return m.Data2
}

//  updateContactsReset#7084a7be = Update;
//
func (m *TLUpdateContactsReset) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateContactsReset
	return m.Data2
}

//  updateNewAuthorization#8f06529a auth_key_id:long date:int device:string location:string = Update;
//
func (m *TLUpdateNewAuthorization) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateNewAuthorization
	return m.Data2
}

//  updateChannelGroup#c36c1e3c channel_id:int group:MessageGroup = Update;
//
func (m *TLUpdateChannelGroup) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChannelGroup
	return m.Data2
}

//  updateChannelAvailableMessages#70db6837 channel_id:int available_min_id:int = Update;
//
func (m *TLUpdateChannelAvailableMessages) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChannelAvailableMessages
	return m.Data2
}

//  updateDialogUnreadMark#e16459c3 flags:# unread:flags.0?true peer:DialogPeer = Update;
//
func (m *TLUpdateDialogUnreadMark) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateDialogUnreadMark
	return m.Data2
}

//  updateUserPinnedMessage#4c43da18 user_id:int id:int = Update;
//
func (m *TLUpdateUserPinnedMessage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateUserPinnedMessage
	return m.Data2
}

//  updateChatPinnedMessage#22893b26 chat_id:int id:int = Update;
//  updateChatPinnedMessage#e10db349 chat_id:int id:int version:int = Update;
//
func (m *TLUpdateChatPinnedMessage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChatPinnedMessage
	return m.Data2
}

//  updateMessagePoll#aca1657b flags:# poll_id:long poll:flags.0?Poll results:PollResults = Update;
//
func (m *TLUpdateMessagePoll) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateMessagePoll
	return m.Data2
}

//  updateChatDefaultBannedRights#54c01850 peer:Peer default_banned_rights:ChatBannedRights version:int = Update;
//
func (m *TLUpdateChatDefaultBannedRights) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChatDefaultBannedRights
	return m.Data2
}

//  updateFolderPeers#19360dc0 folder_peers:Vector<FolderPeer> pts:int pts_count:int = Update;
//
func (m *TLUpdateFolderPeers) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateFolderPeers
	return m.Data2
}

//  updatePeerSettings#6a7e7366 peer:Peer settings:PeerSettings = Update;
//
func (m *TLUpdatePeerSettings) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdatePeerSettings
	return m.Data2
}

//  updatePeerLocated#b4afcfb0 peers:Vector<PeerLocated> = Update;
//
func (m *TLUpdatePeerLocated) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdatePeerLocated
	return m.Data2
}

//  updateNewScheduledMessage#39a51dfb message:Message = Update;
//
func (m *TLUpdateNewScheduledMessage) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateNewScheduledMessage
	return m.Data2
}

//  updateDeleteScheduledMessages#90866cee peer:Peer messages:Vector<int> = Update;
//
func (m *TLUpdateDeleteScheduledMessages) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateDeleteScheduledMessages
	return m.Data2
}

//  updateTheme#8216fba3 theme:Theme = Update;
//
func (m *TLUpdateTheme) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateTheme
	return m.Data2
}

//  updateGeoLiveViewed#871fb939 peer:Peer msg_id:int = Update;
//
func (m *TLUpdateGeoLiveViewed) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateGeoLiveViewed
	return m.Data2
}

//  updateLoginToken#564fe691 = Update;
//
func (m *TLUpdateLoginToken) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateLoginToken
	return m.Data2
}

//  updateMessagePollVote#42f88f2c poll_id:long user_id:int options:Vector<bytes> = Update;
//
func (m *TLUpdateMessagePollVote) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateMessagePollVote
	return m.Data2
}

//  updateDialogFilter#26ffde7d flags:# id:int filter:flags.0?DialogFilter = Update;
//
func (m *TLUpdateDialogFilter) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateDialogFilter
	return m.Data2
}

//  updateDialogFilterOrder#a5d72105 order:Vector<int> = Update;
//
func (m *TLUpdateDialogFilterOrder) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateDialogFilterOrder
	return m.Data2
}

//  updateDialogFilters#3504914f = Update;
//
func (m *TLUpdateDialogFilters) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateDialogFilters
	return m.Data2
}

//  updatePhoneCallSignalingData#2661bf09 phone_call_id:long data:bytes = Update;
//
func (m *TLUpdatePhoneCallSignalingData) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdatePhoneCallSignalingData
	return m.Data2
}

//  updateChannelParticipant#65d2b464 flags:# channel_id:int date:int user_id:int prev_participant:flags.0?ChannelParticipant new_participant:flags.1?ChannelParticipant qts:int = Update;
//
func (m *TLUpdateChannelParticipant) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChannelParticipant
	return m.Data2
}

//  updateChannelMessageForwards#6e8a84df channel_id:int id:int forwards:int = Update;
//
func (m *TLUpdateChannelMessageForwards) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChannelMessageForwards
	return m.Data2
}

//  updateReadChannelDiscussionInbox#1cc7de54 flags:# channel_id:int top_msg_id:int read_max_id:int broadcast_id:flags.0?int broadcast_post:flags.0?int = Update;
//
func (m *TLUpdateReadChannelDiscussionInbox) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateReadChannelDiscussionInbox
	return m.Data2
}

//  updateReadChannelDiscussionOutbox#4638a26c channel_id:int top_msg_id:int read_max_id:int = Update;
//
func (m *TLUpdateReadChannelDiscussionOutbox) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateReadChannelDiscussionOutbox
	return m.Data2
}

//  updatePeerBlocked#246a4b22 peer_id:Peer blocked:Bool = Update;
//
func (m *TLUpdatePeerBlocked) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdatePeerBlocked
	return m.Data2
}

//  updateChannelUserTyping#ff2abe9f flags:# channel_id:int top_msg_id:flags.0?int user_id:int action:SendMessageAction = Update;
//
func (m *TLUpdateChannelUserTyping) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChannelUserTyping
	return m.Data2
}

//  updatePinnedMessages#ed85eab5 flags:# pinned:flags.0?true peer:Peer messages:Vector<int> pts:int pts_count:int = Update;
//
func (m *TLUpdatePinnedMessages) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdatePinnedMessages
	return m.Data2
}

//  updatePinnedChannelMessages#8588878b flags:# pinned:flags.0?true channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
//
func (m *TLUpdatePinnedChannelMessages) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdatePinnedChannelMessages
	return m.Data2
}

//  updateChat#1330a196 chat_id:int = Update;
//
func (m *TLUpdateChat) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateChat
	return m.Data2
}

//  updateGroupCallParticipants#f2ebdb4e call:InputGroupCall participants:Vector<GroupCallParticipant> version:int = Update;
//
func (m *TLUpdateGroupCallParticipants) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateGroupCallParticipants
	return m.Data2
}

//  updateGroupCall#a45eb99b chat_id:int call:GroupCall = Update;
//
func (m *TLUpdateGroupCall) To_Update() *Update {
	m.Data2.Cmd = Cmd_UpdateGroupCall
	return m.Data2
}

//  + TL_UpdatesTooLong
//  + TL_UpdateShortMessage
//  + TL_UpdateShortChatMessage
//  + TL_UpdateShort
//  + TL_UpdatesCombined
//  + TL_Updates
//  + TL_UpdateShortSentMessage
//

//  updatesTooLong#e317af7e = Updates;
//
func (m *Updates) To_UpdatesTooLong() *TLUpdatesTooLong {
	return &TLUpdatesTooLong{
		Data2: m,
	}
}

//  updateShortMessage#914fbf11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
//  updateShortMessage#2296d2c8 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to:flags.3?MessageReplyHeader entities:flags.7?Vector<MessageEntity> = Updates;
//
func (m *Updates) To_UpdateShortMessage() *TLUpdateShortMessage {
	return &TLUpdateShortMessage{
		Data2: m,
	}
}

//  updateShortChatMessage#16812688 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
//  updateShortChatMessage#402d5dbb flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to:flags.3?MessageReplyHeader entities:flags.7?Vector<MessageEntity> = Updates;
//
func (m *Updates) To_UpdateShortChatMessage() *TLUpdateShortChatMessage {
	return &TLUpdateShortChatMessage{
		Data2: m,
	}
}

//  updateShort#78d4dec1 update:Update date:int = Updates;
//
func (m *Updates) To_UpdateShort() *TLUpdateShort {
	return &TLUpdateShort{
		Data2: m,
	}
}

//  updatesCombined#725b04c3 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq_start:int seq:int = Updates;
//
func (m *Updates) To_UpdatesCombined() *TLUpdatesCombined {
	return &TLUpdatesCombined{
		Data2: m,
	}
}

//  updates#74ae4240 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq:int = Updates;
//
func (m *Updates) To_Updates() *TLUpdates {
	return &TLUpdates{
		Data2: m,
	}
}

//  updateShortSentMessage#11f1331c flags:# out:flags.1?true id:int pts:int pts_count:int date:int media:flags.9?MessageMedia entities:flags.7?Vector<MessageEntity> = Updates;
//
func (m *Updates) To_UpdateShortSentMessage() *TLUpdateShortSentMessage {
	return &TLUpdateShortSentMessage{
		Data2: m,
	}
}

//  updatesTooLong#e317af7e = Updates;
//
func (m *TLUpdatesTooLong) To_Updates() *Updates {
	m.Data2.Cmd = Cmd_UpdatesTooLong
	return m.Data2
}

//  updateShortMessage#914fbf11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
//  updateShortMessage#2296d2c8 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to:flags.3?MessageReplyHeader entities:flags.7?Vector<MessageEntity> = Updates;
//
func (m *TLUpdateShortMessage) To_Updates() *Updates {
	m.Data2.Cmd = Cmd_UpdateShortMessage
	return m.Data2
}

//  updateShortChatMessage#16812688 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
//  updateShortChatMessage#402d5dbb flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to:flags.3?MessageReplyHeader entities:flags.7?Vector<MessageEntity> = Updates;
//
func (m *TLUpdateShortChatMessage) To_Updates() *Updates {
	m.Data2.Cmd = Cmd_UpdateShortChatMessage
	return m.Data2
}

//  updateShort#78d4dec1 update:Update date:int = Updates;
//
func (m *TLUpdateShort) To_Updates() *Updates {
	m.Data2.Cmd = Cmd_UpdateShort
	return m.Data2
}

//  updatesCombined#725b04c3 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq_start:int seq:int = Updates;
//
func (m *TLUpdatesCombined) To_Updates() *Updates {
	m.Data2.Cmd = Cmd_UpdatesCombined
	return m.Data2
}

//  updates#74ae4240 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq:int = Updates;
//
func (m *TLUpdates) To_Updates() *Updates {
	m.Data2.Cmd = Cmd_Updates
	return m.Data2
}

//  updateShortSentMessage#11f1331c flags:# out:flags.1?true id:int pts:int pts_count:int date:int media:flags.9?MessageMedia entities:flags.7?Vector<MessageEntity> = Updates;
//
func (m *TLUpdateShortSentMessage) To_Updates() *Updates {
	m.Data2.Cmd = Cmd_UpdateShortSentMessage
	return m.Data2
}

//  + TL_UpdatesChannelDifferenceEmpty
//  + TL_UpdatesChannelDifferenceTooLong
//  + TL_UpdatesChannelDifference
//

//  updates.channelDifferenceEmpty#3e11affb flags:# final:flags.0?true pts:int timeout:flags.1?int = updates.ChannelDifference;
//
func (m *Updates_ChannelDifference) To_UpdatesChannelDifferenceEmpty() *TLUpdatesChannelDifferenceEmpty {
	return &TLUpdatesChannelDifferenceEmpty{
		Data2: m,
	}
}

//  updates.channelDifferenceTooLong#6a9d7b35 flags:# final:flags.0?true pts:int timeout:flags.1?int top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
//  updates.channelDifferenceTooLong#5e167646 flags:# final:flags.0?true pts:int timeout:flags.1?int top_message:int top_important_message:int read_inbox_max_id:int unread_count:int unread_important_count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
//  updates.channelDifferenceTooLong#a4bcc6fe flags:# final:flags.0?true timeout:flags.1?int dialog:Dialog messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
//
func (m *Updates_ChannelDifference) To_UpdatesChannelDifferenceTooLong() *TLUpdatesChannelDifferenceTooLong {
	return &TLUpdatesChannelDifferenceTooLong{
		Data2: m,
	}
}

//  updates.channelDifference#2064674e flags:# final:flags.0?true pts:int timeout:flags.1?int new_messages:Vector<Message> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
//
func (m *Updates_ChannelDifference) To_UpdatesChannelDifference() *TLUpdatesChannelDifference {
	return &TLUpdatesChannelDifference{
		Data2: m,
	}
}

//  updates.channelDifferenceEmpty#3e11affb flags:# final:flags.0?true pts:int timeout:flags.1?int = updates.ChannelDifference;
//
func (m *TLUpdatesChannelDifferenceEmpty) To_Updates_ChannelDifference() *Updates_ChannelDifference {
	m.Data2.Cmd = Cmd_UpdatesChannelDifferenceEmpty
	return m.Data2
}

//  updates.channelDifferenceTooLong#6a9d7b35 flags:# final:flags.0?true pts:int timeout:flags.1?int top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
//  updates.channelDifferenceTooLong#5e167646 flags:# final:flags.0?true pts:int timeout:flags.1?int top_message:int top_important_message:int read_inbox_max_id:int unread_count:int unread_important_count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
//  updates.channelDifferenceTooLong#a4bcc6fe flags:# final:flags.0?true timeout:flags.1?int dialog:Dialog messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
//
func (m *TLUpdatesChannelDifferenceTooLong) To_Updates_ChannelDifference() *Updates_ChannelDifference {
	m.Data2.Cmd = Cmd_UpdatesChannelDifferenceTooLong
	return m.Data2
}

//  updates.channelDifference#2064674e flags:# final:flags.0?true pts:int timeout:flags.1?int new_messages:Vector<Message> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
//
func (m *TLUpdatesChannelDifference) To_Updates_ChannelDifference() *Updates_ChannelDifference {
	m.Data2.Cmd = Cmd_UpdatesChannelDifference
	return m.Data2
}

//  + TL_UpdatesDifferenceEmpty
//  + TL_UpdatesDifference
//  + TL_UpdatesDifferenceSlice
//  + TL_UpdatesDifferenceTooLong
//

//  updates.differenceEmpty#5d75a138 date:int seq:int = updates.Difference;
//
func (m *Updates_Difference) To_UpdatesDifferenceEmpty() *TLUpdatesDifferenceEmpty {
	return &TLUpdatesDifferenceEmpty{
		Data2: m,
	}
}

//  updates.difference#f49ca0 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> state:updates.State = updates.Difference;
//
func (m *Updates_Difference) To_UpdatesDifference() *TLUpdatesDifference {
	return &TLUpdatesDifference{
		Data2: m,
	}
}

//  updates.differenceSlice#a8fb1981 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> intermediate_state:updates.State = updates.Difference;
//
func (m *Updates_Difference) To_UpdatesDifferenceSlice() *TLUpdatesDifferenceSlice {
	return &TLUpdatesDifferenceSlice{
		Data2: m,
	}
}

//  updates.differenceTooLong#4afe8f6d pts:int = updates.Difference;
//
func (m *Updates_Difference) To_UpdatesDifferenceTooLong() *TLUpdatesDifferenceTooLong {
	return &TLUpdatesDifferenceTooLong{
		Data2: m,
	}
}

//  updates.differenceEmpty#5d75a138 date:int seq:int = updates.Difference;
//
func (m *TLUpdatesDifferenceEmpty) To_Updates_Difference() *Updates_Difference {
	m.Data2.Cmd = Cmd_UpdatesDifferenceEmpty
	return m.Data2
}

//  updates.difference#f49ca0 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> state:updates.State = updates.Difference;
//
func (m *TLUpdatesDifference) To_Updates_Difference() *Updates_Difference {
	m.Data2.Cmd = Cmd_UpdatesDifference
	return m.Data2
}

//  updates.differenceSlice#a8fb1981 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> intermediate_state:updates.State = updates.Difference;
//
func (m *TLUpdatesDifferenceSlice) To_Updates_Difference() *Updates_Difference {
	m.Data2.Cmd = Cmd_UpdatesDifferenceSlice
	return m.Data2
}

//  updates.differenceTooLong#4afe8f6d pts:int = updates.Difference;
//
func (m *TLUpdatesDifferenceTooLong) To_Updates_Difference() *Updates_Difference {
	m.Data2.Cmd = Cmd_UpdatesDifferenceTooLong
	return m.Data2
}

//  + TL_UpdatesState
//

//  updates.state#a56c2a3e pts:int qts:int date:int seq:int unread_count:int = updates.State;
//
func (m *Updates_State) To_UpdatesState() *TLUpdatesState {
	return &TLUpdatesState{
		Data2: m,
	}
}

//  updates.state#a56c2a3e pts:int qts:int date:int seq:int unread_count:int = updates.State;
//
func (m *TLUpdatesState) To_Updates_State() *Updates_State {
	m.Data2.Cmd = Cmd_UpdatesState
	return m.Data2
}

//  + TL_UploadCdnFileReuploadNeeded
//  + TL_UploadCdnFile
//

//  upload.cdnFileReuploadNeeded#eea8e46e request_token:bytes = upload.CdnFile;
//
func (m *Upload_CdnFile) To_UploadCdnFileReuploadNeeded() *TLUploadCdnFileReuploadNeeded {
	return &TLUploadCdnFileReuploadNeeded{
		Data2: m,
	}
}

//  upload.cdnFile#a99fca4f bytes:bytes = upload.CdnFile;
//
func (m *Upload_CdnFile) To_UploadCdnFile() *TLUploadCdnFile {
	return &TLUploadCdnFile{
		Data2: m,
	}
}

//  upload.cdnFileReuploadNeeded#eea8e46e request_token:bytes = upload.CdnFile;
//
func (m *TLUploadCdnFileReuploadNeeded) To_Upload_CdnFile() *Upload_CdnFile {
	m.Data2.Cmd = Cmd_UploadCdnFileReuploadNeeded
	return m.Data2
}

//  upload.cdnFile#a99fca4f bytes:bytes = upload.CdnFile;
//
func (m *TLUploadCdnFile) To_Upload_CdnFile() *Upload_CdnFile {
	m.Data2.Cmd = Cmd_UploadCdnFile
	return m.Data2
}

//  + TL_UploadFile
//  + TL_UploadFileCdnRedirect
//

//  upload.file#96a18d5 type:storage.FileType mtime:int bytes:bytes = upload.File;
//
func (m *Upload_File) To_UploadFile() *TLUploadFile {
	return &TLUploadFile{
		Data2: m,
	}
}

//  upload.fileCdnRedirect#ea52fe5a dc_id:int file_token:bytes encryption_key:bytes encryption_iv:bytes cdn_file_hashes:Vector<CdnFileHash> = upload.File;
//  upload.fileCdnRedirect#f18cda44 dc_id:int file_token:bytes encryption_key:bytes encryption_iv:bytes file_hashes:Vector<FileHash> = upload.File;
//
func (m *Upload_File) To_UploadFileCdnRedirect() *TLUploadFileCdnRedirect {
	return &TLUploadFileCdnRedirect{
		Data2: m,
	}
}

//  upload.file#96a18d5 type:storage.FileType mtime:int bytes:bytes = upload.File;
//
func (m *TLUploadFile) To_Upload_File() *Upload_File {
	m.Data2.Cmd = Cmd_UploadFile
	return m.Data2
}

//  upload.fileCdnRedirect#ea52fe5a dc_id:int file_token:bytes encryption_key:bytes encryption_iv:bytes cdn_file_hashes:Vector<CdnFileHash> = upload.File;
//  upload.fileCdnRedirect#f18cda44 dc_id:int file_token:bytes encryption_key:bytes encryption_iv:bytes file_hashes:Vector<FileHash> = upload.File;
//
func (m *TLUploadFileCdnRedirect) To_Upload_File() *Upload_File {
	m.Data2.Cmd = Cmd_UploadFileCdnRedirect
	return m.Data2
}

//  + TL_UploadWebFile
//

//  upload.webFile#21e753bc size:int mime_type:string file_type:storage.FileType mtime:int bytes:bytes = upload.WebFile;
//
func (m *Upload_WebFile) To_UploadWebFile() *TLUploadWebFile {
	return &TLUploadWebFile{
		Data2: m,
	}
}

//  upload.webFile#21e753bc size:int mime_type:string file_type:storage.FileType mtime:int bytes:bytes = upload.WebFile;
//
func (m *TLUploadWebFile) To_Upload_WebFile() *Upload_WebFile {
	m.Data2.Cmd = Cmd_UploadWebFile
	return m.Data2
}

//  + TL_UrlAuthResultRequest
//  + TL_UrlAuthResultAccepted
//  + TL_UrlAuthResultDefault
//

//  urlAuthResultRequest#92d33a0e flags:# request_write_access:flags.0?true bot:User domain:string = UrlAuthResult;
//
func (m *UrlAuthResult) To_UrlAuthResultRequest() *TLUrlAuthResultRequest {
	return &TLUrlAuthResultRequest{
		Data2: m,
	}
}

//  urlAuthResultAccepted#8f8c0e4e url:string = UrlAuthResult;
//
func (m *UrlAuthResult) To_UrlAuthResultAccepted() *TLUrlAuthResultAccepted {
	return &TLUrlAuthResultAccepted{
		Data2: m,
	}
}

//  urlAuthResultDefault#a9d6db1f = UrlAuthResult;
//
func (m *UrlAuthResult) To_UrlAuthResultDefault() *TLUrlAuthResultDefault {
	return &TLUrlAuthResultDefault{
		Data2: m,
	}
}

//  urlAuthResultRequest#92d33a0e flags:# request_write_access:flags.0?true bot:User domain:string = UrlAuthResult;
//
func (m *TLUrlAuthResultRequest) To_UrlAuthResult() *UrlAuthResult {
	m.Data2.Cmd = Cmd_UrlAuthResultRequest
	return m.Data2
}

//  urlAuthResultAccepted#8f8c0e4e url:string = UrlAuthResult;
//
func (m *TLUrlAuthResultAccepted) To_UrlAuthResult() *UrlAuthResult {
	m.Data2.Cmd = Cmd_UrlAuthResultAccepted
	return m.Data2
}

//  urlAuthResultDefault#a9d6db1f = UrlAuthResult;
//
func (m *TLUrlAuthResultDefault) To_UrlAuthResult() *UrlAuthResult {
	m.Data2.Cmd = Cmd_UrlAuthResultDefault
	return m.Data2
}

//  + TL_UserEmpty
//  + TL_User
//

//  userEmpty#200250ba id:int = User;
//
func (m *User) To_UserEmpty() *TLUserEmpty {
	return &TLUserEmpty{
		Data2: m,
	}
}

//  user#2e13f4c3 flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true id:int access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?string bot_inline_placeholder:flags.19?string lang_code:flags.22?string = User;
//  user#d10d979a flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true id:int access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?string bot_inline_placeholder:flags.19?string = User;
//  user#938458c1 flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true support:flags.23?true scam:flags.24?true apply_min_photo:flags.25?true id:int access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?Vector<RestrictionReason> bot_inline_placeholder:flags.19?string lang_code:flags.22?string = User;
//
func (m *User) To_User() *TLUser {
	return &TLUser{
		Data2: m,
	}
}

//  userEmpty#200250ba id:int = User;
//
func (m *TLUserEmpty) To_User() *User {
	m.Data2.Cmd = Cmd_UserEmpty
	return m.Data2
}

//  user#2e13f4c3 flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true id:int access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?string bot_inline_placeholder:flags.19?string lang_code:flags.22?string = User;
//  user#d10d979a flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true id:int access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?string bot_inline_placeholder:flags.19?string = User;
//  user#938458c1 flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true support:flags.23?true scam:flags.24?true apply_min_photo:flags.25?true id:int access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?Vector<RestrictionReason> bot_inline_placeholder:flags.19?string lang_code:flags.22?string = User;
//
func (m *TLUser) To_User() *User {
	m.Data2.Cmd = Cmd_User
	return m.Data2
}

//  + TL_UserFull
//

//  userFull#f220f3f flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo common_chats_count:int = UserFull;
//  userFull#5932fc03 flags:# blocked:flags.0?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo = UserFull;
//  userFull#8ea4a881 flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true can_pin_message:flags.7?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo pinned_msg_id:flags.6?int common_chats_count:int = UserFull;
//  userFull#745559cc flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true can_pin_message:flags.7?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo pinned_msg_id:flags.6?int common_chats_count:int folder_id:flags.11?int = UserFull;
//  userFull#edf17c12 flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true can_pin_message:flags.7?true has_scheduled:flags.12?true video_calls_available:flags.13?true user:User about:flags.1?string settings:PeerSettings profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo pinned_msg_id:flags.6?int common_chats_count:int folder_id:flags.11?int = UserFull;
//
func (m *UserFull) To_UserFull() *TLUserFull {
	return &TLUserFull{
		Data2: m,
	}
}

//  userFull#f220f3f flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo common_chats_count:int = UserFull;
//  userFull#5932fc03 flags:# blocked:flags.0?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo = UserFull;
//  userFull#8ea4a881 flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true can_pin_message:flags.7?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo pinned_msg_id:flags.6?int common_chats_count:int = UserFull;
//  userFull#745559cc flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true can_pin_message:flags.7?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo pinned_msg_id:flags.6?int common_chats_count:int folder_id:flags.11?int = UserFull;
//  userFull#edf17c12 flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true can_pin_message:flags.7?true has_scheduled:flags.12?true video_calls_available:flags.13?true user:User about:flags.1?string settings:PeerSettings profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo pinned_msg_id:flags.6?int common_chats_count:int folder_id:flags.11?int = UserFull;
//
func (m *TLUserFull) To_UserFull() *UserFull {
	m.Data2.Cmd = Cmd_UserFull
	return m.Data2
}

//  + TL_UserProfilePhotoEmpty
//  + TL_UserProfilePhoto
//

//  userProfilePhotoEmpty#4f11bae1 = UserProfilePhoto;
//
func (m *UserProfilePhoto) To_UserProfilePhotoEmpty() *TLUserProfilePhotoEmpty {
	return &TLUserProfilePhotoEmpty{
		Data2: m,
	}
}

//  userProfilePhoto#d559d8c8 photo_id:long photo_small:FileLocation photo_big:FileLocation = UserProfilePhoto;
//  userProfilePhoto#ecd75d8c photo_id:long photo_small:FileLocation photo_big:FileLocation dc_id:int = UserProfilePhoto;
//  userProfilePhoto#69d3ab26 flags:# has_video:flags.0?true photo_id:long photo_small:FileLocation photo_big:FileLocation dc_id:int = UserProfilePhoto;
//
func (m *UserProfilePhoto) To_UserProfilePhoto() *TLUserProfilePhoto {
	return &TLUserProfilePhoto{
		Data2: m,
	}
}

//  userProfilePhotoEmpty#4f11bae1 = UserProfilePhoto;
//
func (m *TLUserProfilePhotoEmpty) To_UserProfilePhoto() *UserProfilePhoto {
	m.Data2.Cmd = Cmd_UserProfilePhotoEmpty
	return m.Data2
}

//  userProfilePhoto#d559d8c8 photo_id:long photo_small:FileLocation photo_big:FileLocation = UserProfilePhoto;
//  userProfilePhoto#ecd75d8c photo_id:long photo_small:FileLocation photo_big:FileLocation dc_id:int = UserProfilePhoto;
//  userProfilePhoto#69d3ab26 flags:# has_video:flags.0?true photo_id:long photo_small:FileLocation photo_big:FileLocation dc_id:int = UserProfilePhoto;
//
func (m *TLUserProfilePhoto) To_UserProfilePhoto() *UserProfilePhoto {
	m.Data2.Cmd = Cmd_UserProfilePhoto
	return m.Data2
}

//  + TL_UserStatusEmpty
//  + TL_UserStatusOnline
//  + TL_UserStatusOffline
//  + TL_UserStatusRecently
//  + TL_UserStatusLastWeek
//  + TL_UserStatusLastMonth
//

//  userStatusEmpty#9d05049 = UserStatus;
//
func (m *UserStatus) To_UserStatusEmpty() *TLUserStatusEmpty {
	return &TLUserStatusEmpty{
		Data2: m,
	}
}

//  userStatusOnline#edb93949 expires:int = UserStatus;
//
func (m *UserStatus) To_UserStatusOnline() *TLUserStatusOnline {
	return &TLUserStatusOnline{
		Data2: m,
	}
}

//  userStatusOffline#8c703f was_online:int = UserStatus;
//
func (m *UserStatus) To_UserStatusOffline() *TLUserStatusOffline {
	return &TLUserStatusOffline{
		Data2: m,
	}
}

//  userStatusRecently#e26f42f1 = UserStatus;
//
func (m *UserStatus) To_UserStatusRecently() *TLUserStatusRecently {
	return &TLUserStatusRecently{
		Data2: m,
	}
}

//  userStatusLastWeek#7bf09fc = UserStatus;
//
func (m *UserStatus) To_UserStatusLastWeek() *TLUserStatusLastWeek {
	return &TLUserStatusLastWeek{
		Data2: m,
	}
}

//  userStatusLastMonth#77ebc742 = UserStatus;
//
func (m *UserStatus) To_UserStatusLastMonth() *TLUserStatusLastMonth {
	return &TLUserStatusLastMonth{
		Data2: m,
	}
}

//  userStatusEmpty#9d05049 = UserStatus;
//
func (m *TLUserStatusEmpty) To_UserStatus() *UserStatus {
	m.Data2.Cmd = Cmd_UserStatusEmpty
	return m.Data2
}

//  userStatusOnline#edb93949 expires:int = UserStatus;
//
func (m *TLUserStatusOnline) To_UserStatus() *UserStatus {
	m.Data2.Cmd = Cmd_UserStatusOnline
	return m.Data2
}

//  userStatusOffline#8c703f was_online:int = UserStatus;
//
func (m *TLUserStatusOffline) To_UserStatus() *UserStatus {
	m.Data2.Cmd = Cmd_UserStatusOffline
	return m.Data2
}

//  userStatusRecently#e26f42f1 = UserStatus;
//
func (m *TLUserStatusRecently) To_UserStatus() *UserStatus {
	m.Data2.Cmd = Cmd_UserStatusRecently
	return m.Data2
}

//  userStatusLastWeek#7bf09fc = UserStatus;
//
func (m *TLUserStatusLastWeek) To_UserStatus() *UserStatus {
	m.Data2.Cmd = Cmd_UserStatusLastWeek
	return m.Data2
}

//  userStatusLastMonth#77ebc742 = UserStatus;
//
func (m *TLUserStatusLastMonth) To_UserStatus() *UserStatus {
	m.Data2.Cmd = Cmd_UserStatusLastMonth
	return m.Data2
}

//  + TL_VideoSize
//

//  videoSize#435bb987 type:string location:FileLocation w:int h:int size:int = VideoSize;
//  videoSize#e831c556 flags:# type:string location:FileLocation w:int h:int size:int video_start_ts:flags.0?double = VideoSize;
//
func (m *VideoSize) To_VideoSize() *TLVideoSize {
	return &TLVideoSize{
		Data2: m,
	}
}

//  videoSize#435bb987 type:string location:FileLocation w:int h:int size:int = VideoSize;
//  videoSize#e831c556 flags:# type:string location:FileLocation w:int h:int size:int video_start_ts:flags.0?double = VideoSize;
//
func (m *TLVideoSize) To_VideoSize() *VideoSize {
	m.Data2.Cmd = Cmd_VideoSize
	return m.Data2
}

//  + TL_WallPaper
//  + TL_WallPaperSolid
//  + TL_WallPaperNoFile
//

//  wallPaper#ccb03657 id:int title:string sizes:Vector<PhotoSize> color:int = WallPaper;
//  wallPaper#f04f91ec id:long flags:# creator:flags.0?true default:flags.1?true access_hash:long slug:string document:Document = WallPaper;
//  wallPaper#a437c3ed id:long flags:# creator:flags.0?true default:flags.1?true pattern:flags.3?true dark:flags.4?true access_hash:long slug:string document:Document settings:flags.2?WallPaperSettings = WallPaper;
//
func (m *WallPaper) To_WallPaper() *TLWallPaper {
	return &TLWallPaper{
		Data2: m,
	}
}

//  wallPaperSolid#63117f24 id:int title:string bg_color:int color:int = WallPaper;
//
func (m *WallPaper) To_WallPaperSolid() *TLWallPaperSolid {
	return &TLWallPaperSolid{
		Data2: m,
	}
}

//  wallPaperNoFile#8af40b25 flags:# default:flags.1?true dark:flags.4?true settings:flags.2?WallPaperSettings = WallPaper;
//
func (m *WallPaper) To_WallPaperNoFile() *TLWallPaperNoFile {
	return &TLWallPaperNoFile{
		Data2: m,
	}
}

//  wallPaper#ccb03657 id:int title:string sizes:Vector<PhotoSize> color:int = WallPaper;
//  wallPaper#f04f91ec id:long flags:# creator:flags.0?true default:flags.1?true access_hash:long slug:string document:Document = WallPaper;
//  wallPaper#a437c3ed id:long flags:# creator:flags.0?true default:flags.1?true pattern:flags.3?true dark:flags.4?true access_hash:long slug:string document:Document settings:flags.2?WallPaperSettings = WallPaper;
//
func (m *TLWallPaper) To_WallPaper() *WallPaper {
	m.Data2.Cmd = Cmd_WallPaper
	return m.Data2
}

//  wallPaperSolid#63117f24 id:int title:string bg_color:int color:int = WallPaper;
//
func (m *TLWallPaperSolid) To_WallPaper() *WallPaper {
	m.Data2.Cmd = Cmd_WallPaperSolid
	return m.Data2
}

//  wallPaperNoFile#8af40b25 flags:# default:flags.1?true dark:flags.4?true settings:flags.2?WallPaperSettings = WallPaper;
//
func (m *TLWallPaperNoFile) To_WallPaper() *WallPaper {
	m.Data2.Cmd = Cmd_WallPaperNoFile
	return m.Data2
}

//  + TL_WallPaperSettings
//

//  wallPaperSettings#a12f40b8 flags:# blur:flags.1?true motion:flags.2?true background_color:flags.0?int intensity:flags.3?int = WallPaperSettings;
//  wallPaperSettings#5086cf8 flags:# blur:flags.1?true motion:flags.2?true background_color:flags.0?int second_background_color:flags.4?int intensity:flags.3?int rotation:flags.4?int = WallPaperSettings;
//
func (m *WallPaperSettings) To_WallPaperSettings() *TLWallPaperSettings {
	return &TLWallPaperSettings{
		Data2: m,
	}
}

//  wallPaperSettings#a12f40b8 flags:# blur:flags.1?true motion:flags.2?true background_color:flags.0?int intensity:flags.3?int = WallPaperSettings;
//  wallPaperSettings#5086cf8 flags:# blur:flags.1?true motion:flags.2?true background_color:flags.0?int second_background_color:flags.4?int intensity:flags.3?int rotation:flags.4?int = WallPaperSettings;
//
func (m *TLWallPaperSettings) To_WallPaperSettings() *WallPaperSettings {
	m.Data2.Cmd = Cmd_WallPaperSettings
	return m.Data2
}

//  + TL_WalletSecretSalt
//

//  wallet.secretSalt#dd484d64 salt:bytes = wallet.KeySecretSalt;
//
func (m *Wallet_KeySecretSalt) To_WalletSecretSalt() *TLWalletSecretSalt {
	return &TLWalletSecretSalt{
		Data2: m,
	}
}

//  wallet.secretSalt#dd484d64 salt:bytes = wallet.KeySecretSalt;
//
func (m *TLWalletSecretSalt) To_Wallet_KeySecretSalt() *Wallet_KeySecretSalt {
	m.Data2.Cmd = Cmd_WalletSecretSalt
	return m.Data2
}

//  + TL_WalletLiteResponse
//

//  wallet.liteResponse#764386d7 response:bytes = wallet.LiteResponse;
//
func (m *Wallet_LiteResponse) To_WalletLiteResponse() *TLWalletLiteResponse {
	return &TLWalletLiteResponse{
		Data2: m,
	}
}

//  wallet.liteResponse#764386d7 response:bytes = wallet.LiteResponse;
//
func (m *TLWalletLiteResponse) To_Wallet_LiteResponse() *Wallet_LiteResponse {
	m.Data2.Cmd = Cmd_WalletLiteResponse
	return m.Data2
}

//  + TL_WebAuthorization
//

//  webAuthorization#cac943f2 hash:long bot_id:int domain:string browser:string platform:string date_created:int date_active:int ip:string region:string = WebAuthorization;
//
func (m *WebAuthorization) To_WebAuthorization() *TLWebAuthorization {
	return &TLWebAuthorization{
		Data2: m,
	}
}

//  webAuthorization#cac943f2 hash:long bot_id:int domain:string browser:string platform:string date_created:int date_active:int ip:string region:string = WebAuthorization;
//
func (m *TLWebAuthorization) To_WebAuthorization() *WebAuthorization {
	m.Data2.Cmd = Cmd_WebAuthorization
	return m.Data2
}

//  + TL_WebDocument
//  + TL_WebDocumentNoProxy
//

//  webDocument#c61acbd8 url:string access_hash:long size:int mime_type:string attributes:Vector<DocumentAttribute> dc_id:int = WebDocument;
//  webDocument#1c570ed1 url:string access_hash:long size:int mime_type:string attributes:Vector<DocumentAttribute> = WebDocument;
//
func (m *WebDocument) To_WebDocument() *TLWebDocument {
	return &TLWebDocument{
		Data2: m,
	}
}

//  webDocumentNoProxy#f9c8bcc6 url:string size:int mime_type:string attributes:Vector<DocumentAttribute> = WebDocument;
//
func (m *WebDocument) To_WebDocumentNoProxy() *TLWebDocumentNoProxy {
	return &TLWebDocumentNoProxy{
		Data2: m,
	}
}

//  webDocument#c61acbd8 url:string access_hash:long size:int mime_type:string attributes:Vector<DocumentAttribute> dc_id:int = WebDocument;
//  webDocument#1c570ed1 url:string access_hash:long size:int mime_type:string attributes:Vector<DocumentAttribute> = WebDocument;
//
func (m *TLWebDocument) To_WebDocument() *WebDocument {
	m.Data2.Cmd = Cmd_WebDocument
	return m.Data2
}

//  webDocumentNoProxy#f9c8bcc6 url:string size:int mime_type:string attributes:Vector<DocumentAttribute> = WebDocument;
//
func (m *TLWebDocumentNoProxy) To_WebDocument() *WebDocument {
	m.Data2.Cmd = Cmd_WebDocumentNoProxy
	return m.Data2
}

//  + TL_WebPageEmpty
//  + TL_WebPagePending
//  + TL_WebPage
//  + TL_WebPageNotModified
//

//  webPageEmpty#eb1477e8 id:long = WebPage;
//
func (m *WebPage) To_WebPageEmpty() *TLWebPageEmpty {
	return &TLWebPageEmpty{
		Data2: m,
	}
}

//  webPagePending#c586da1c id:long date:int = WebPage;
//
func (m *WebPage) To_WebPagePending() *TLWebPagePending {
	return &TLWebPagePending{
		Data2: m,
	}
}

//  webPage#5f07b4bc flags:# id:long url:string display_url:string hash:int type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document cached_page:flags.10?Page = WebPage;
//  webPage#ca820ed7 flags:# id:long url:string display_url:string type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document = WebPage;
//  webPage#fa64e172 flags:# id:long url:string display_url:string hash:int type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document documents:flags.11?Vector<Document> cached_page:flags.10?Page = WebPage;
//  webPage#e89c45b2 flags:# id:long url:string display_url:string hash:int type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document cached_page:flags.10?Page attributes:flags.12?Vector<WebPageAttribute> = WebPage;
//
func (m *WebPage) To_WebPage() *TLWebPage {
	return &TLWebPage{
		Data2: m,
	}
}

//  webPageNotModified#85849473 = WebPage;
//  webPageNotModified#7311ca11 flags:# cached_page_views:flags.0?int = WebPage;
//
func (m *WebPage) To_WebPageNotModified() *TLWebPageNotModified {
	return &TLWebPageNotModified{
		Data2: m,
	}
}

//  webPageEmpty#eb1477e8 id:long = WebPage;
//
func (m *TLWebPageEmpty) To_WebPage() *WebPage {
	m.Data2.Cmd = Cmd_WebPageEmpty
	return m.Data2
}

//  webPagePending#c586da1c id:long date:int = WebPage;
//
func (m *TLWebPagePending) To_WebPage() *WebPage {
	m.Data2.Cmd = Cmd_WebPagePending
	return m.Data2
}

//  webPage#5f07b4bc flags:# id:long url:string display_url:string hash:int type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document cached_page:flags.10?Page = WebPage;
//  webPage#ca820ed7 flags:# id:long url:string display_url:string type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document = WebPage;
//  webPage#fa64e172 flags:# id:long url:string display_url:string hash:int type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document documents:flags.11?Vector<Document> cached_page:flags.10?Page = WebPage;
//  webPage#e89c45b2 flags:# id:long url:string display_url:string hash:int type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document cached_page:flags.10?Page attributes:flags.12?Vector<WebPageAttribute> = WebPage;
//
func (m *TLWebPage) To_WebPage() *WebPage {
	m.Data2.Cmd = Cmd_WebPage
	return m.Data2
}

//  webPageNotModified#85849473 = WebPage;
//  webPageNotModified#7311ca11 flags:# cached_page_views:flags.0?int = WebPage;
//
func (m *TLWebPageNotModified) To_WebPage() *WebPage {
	m.Data2.Cmd = Cmd_WebPageNotModified
	return m.Data2
}

//  + TL_WebPageAttributeTheme
//

//  webPageAttributeTheme#54b56617 flags:# documents:flags.0?Vector<Document> settings:flags.1?ThemeSettings = WebPageAttribute;
//
func (m *WebPageAttribute) To_WebPageAttributeTheme() *TLWebPageAttributeTheme {
	return &TLWebPageAttributeTheme{
		Data2: m,
	}
}

//  webPageAttributeTheme#54b56617 flags:# documents:flags.0?Vector<Document> settings:flags.1?ThemeSettings = WebPageAttribute;
//
func (m *TLWebPageAttributeTheme) To_WebPageAttribute() *WebPageAttribute {
	m.Data2.Cmd = Cmd_WebPageAttributeTheme
	return m.Data2
}
