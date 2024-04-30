/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2023/06/12 18:03
 * @File : schema.pb.go
 * @Desc :
 *
 */

package mtproto

import (
	"fmt"
	"novachat_engine/pkg/log"
)

var registerObject = map[TLCmd]func() TLObject{
	Cmd_message_transform: func() TLObject {
		return &TLMessageTransform{}
	},
	Cmd_msg_container: func() TLObject {
		return &TLMsgContainer{}
	},
	Cmd_msg_copy: func() TLObject {
		return &TLMsgCopy{}
	},
	Cmd_gzip_packed: func() TLObject {
		return &TLGZipPacked{}
	},
	Cmd_init_connection: func() TLObject {
		return &TLInitConnection{
			Cmd:       Cmd_init_connection,
			ClassName: Class_init_connection,
		}
	},
	Cmd_init_connectionc7481da6: func() TLObject {
		return &TLInitConnection{
			Cmd:       Cmd_init_connectionc7481da6,
			ClassName: Class_init_connection,
		}
	},
	Cmd_init_connectionc785188b8: func() TLObject {
		return &TLInitConnection{
			Cmd:       Cmd_init_connectionc785188b8,
			ClassName: Class_init_connection,
		}
	},
	Cmd_invokeWithLayer: func() TLObject {
		return &TLInvokeWithLayer{
			Cmd:       Cmd_invokeWithLayer,
			ClassName: Class_invokeWithLayer,
		}
	},
	Cmd_invokeWithoutUpdates: func() TLObject {
		return &TLInvokeWithoutUpdates{
			Cmd:       Cmd_invokeWithoutUpdates,
			ClassName: Class_invokeWithoutUpdates,
		}
	},
	Cmd_invokeWithMessagesRange: func() TLObject {
		return &TLInvokeWithMessagesRange{
			Cmd:       Cmd_invokeWithMessagesRange,
			ClassName: Class_invokeWithMessagesRange,
		}
	},
	Cmd_invokeWithTakeout: func() TLObject {
		return &TLInvokeWithTakeout{
			Cmd:       Cmd_invokeWithTakeout,
			ClassName: Class_invokeWithTakeout,
		}
	},
	Cmd_invokeAfterMsg: func() TLObject {
		return &TLInvokeAfterMsg{
			Cmd:       Cmd_invokeAfterMsg,
			ClassName: Class_invokeAfterMsg,
		}
	},
	Cmd_invokeAfterMsgs: func() TLObject {
		return &TLInvokeAfterMsgs{
			Cmd:       Cmd_invokeAfterMsgs,
			ClassName: Class_invokeAfterMsgs,
		}
	},
	//Cmd_rpc_result : func() TLObject {
	//    return &TLRpcResult{}
	//},

	Cmd_DestroyAuthKey: func() TLObject {
		return NewTLDestroyAuthKey()
	},
	Cmd_ReqPq: func() TLObject {
		return NewTLReqPq()
	},
	Cmd_ReqPqMulti: func() TLObject {
		return NewTLReqPqMulti()
	},
	Cmd_Req_DHParams: func() TLObject {
		return NewTLReq_DHParams()
	},
	Cmd_SetClient_DHParams: func() TLObject {
		return NewTLSetClient_DHParams()
	},

	Cmd_BindAuthKeyInner: func() TLObject {
		return NewTLBindAuthKeyInner(nil)
	},
	Cmd_Client_DHInnerData: func() TLObject {
		return NewTLClient_DHInnerData(nil)
	},
	Cmd_DestroyAuthKeyFail: func() TLObject {
		return NewTLDestroyAuthKeyFail(nil)
	},
	Cmd_DestroyAuthKeyNone: func() TLObject {
		return NewTLDestroyAuthKeyNone(nil)
	},
	Cmd_DestroyAuthKeyOk: func() TLObject {
		return NewTLDestroyAuthKeyOk(nil)
	},
	Cmd_DhGenFail: func() TLObject {
		return NewTLDhGenFail(nil)
	},
	Cmd_DhGenOk: func() TLObject {
		return NewTLDhGenOk(nil)
	},
	Cmd_DhGenRetry: func() TLObject {
		return NewTLDhGenRetry(nil)
	},
	Cmd_PQInnerData: func() TLObject {
		return NewTLPQInnerData(nil)
	},
	Cmd_PQInnerDataDc: func() TLObject {
		return NewTLPQInnerDataDc(nil)
	},
	Cmd_PQInnerDataTemp: func() TLObject {
		return NewTLPQInnerDataTemp(nil)
	},
	Cmd_PQInnerDataTempDc: func() TLObject {
		return NewTLPQInnerDataTempDc(nil)
	},
	Cmd_ResPQ: func() TLObject {
		return NewTLResPQ(nil)
	},
	Cmd_Server_DHInnerData: func() TLObject {
		return NewTLServer_DHInnerData(nil)
	},
	Cmd_Server_DHParamsFail: func() TLObject {
		return NewTLServer_DHParamsFail(nil)
	},
	Cmd_Server_DHParamsOk: func() TLObject {
		return NewTLServer_DHParamsOk(nil)
	},
	Cmd_ContestSaveDeveloperInfo: func() TLObject {
		return NewTLContestSaveDeveloperInfo()
	},
	Cmd_DestroySession: func() TLObject {
		return NewTLDestroySession()
	},
	Cmd_GetFutureSalts: func() TLObject {
		return NewTLGetFutureSalts()
	},
	Cmd_Ping: func() TLObject {
		return NewTLPing()
	},
	Cmd_PingDelayDisconnect: func() TLObject {
		return NewTLPingDelayDisconnect()
	},
	Cmd_RpcDropAnswer: func() TLObject {
		return NewTLRpcDropAnswer()
	},

	Cmd_AccessPointRule: func() TLObject {
		return NewTLAccessPointRule(nil)
	},
	Cmd_BadMsgNotification: func() TLObject {
		return NewTLBadMsgNotification(nil)
	},
	Cmd_BadServerSalt: func() TLObject {
		return NewTLBadServerSalt(nil)
	},
	Cmd_DestroySessionNone: func() TLObject {
		return NewTLDestroySessionNone(nil)
	},
	Cmd_DestroySessionOk: func() TLObject {
		return NewTLDestroySessionOk(nil)
	},
	Cmd_FutureSalt: func() TLObject {
		return NewTLFutureSalt(nil)
	},
	Cmd_FutureSalts: func() TLObject {
		return NewTLFutureSalts(nil)
	},
	Cmd_HelpConfigSimple: func() TLObject {
		return NewTLHelpConfigSimple(nil)
	},
	Cmd_HelpConfigSimple5a592a6c: func() TLObject {
		return NewTLHelpConfigSimple(nil)
	},
	Cmd_HttpWait: func() TLObject {
		return NewTLHttpWait(nil)
	},
	Cmd_IpPort: func() TLObject {
		return NewTLIpPort(nil)
	},
	Cmd_IpPortSecret: func() TLObject {
		return NewTLIpPortSecret(nil)
	},
	Cmd_MsgDetailedInfo: func() TLObject {
		return NewTLMsgDetailedInfo(nil)
	},
	Cmd_MsgNewDetailedInfo: func() TLObject {
		return NewTLMsgNewDetailedInfo(nil)
	},
	Cmd_MsgResendReq: func() TLObject {
		return NewTLMsgResendReq(nil)
	},
	Cmd_MsgsAck: func() TLObject {
		return NewTLMsgsAck(nil)
	},
	Cmd_MsgsAllInfo: func() TLObject {
		return NewTLMsgsAllInfo(nil)
	},
	Cmd_MsgsStateInfo: func() TLObject {
		return NewTLMsgsStateInfo(nil)
	},
	Cmd_MsgsStateReq: func() TLObject {
		return NewTLMsgsStateReq(nil)
	},
	Cmd_NewSessionCreated: func() TLObject {
		return NewTLNewSessionCreated(nil)
	},
	Cmd_Pong: func() TLObject {
		return NewTLPong(nil)
	},
	Cmd_RpcAnswerDropped: func() TLObject {
		return NewTLRpcAnswerDropped(nil)
	},
	Cmd_RpcAnswerDroppedRunning: func() TLObject {
		return NewTLRpcAnswerDroppedRunning(nil)
	},
	Cmd_RpcAnswerUnknown: func() TLObject {
		return NewTLRpcAnswerUnknown(nil)
	},
	Cmd_RpcError: func() TLObject {
		return NewTLRpcError(nil)
	},
	Cmd_AccountAcceptAuthorization: func() TLObject {
		return NewTLAccountAcceptAuthorization()
	},
	Cmd_AccountCancelPasswordEmail: func() TLObject {
		return NewTLAccountCancelPasswordEmail()
	},
	Cmd_AccountChangePhone: func() TLObject {
		return NewTLAccountChangePhone()
	},
	Cmd_AccountCheckUsername: func() TLObject {
		return NewTLAccountCheckUsername()
	},
	Cmd_AccountConfirmPasswordEmail: func() TLObject {
		return NewTLAccountConfirmPasswordEmail()
	},
	Cmd_AccountConfirmPhone: func() TLObject {
		return NewTLAccountConfirmPhone()
	},
	Cmd_AccountCreateTheme: func() TLObject {
		return NewTLAccountCreateTheme()
	},
	Cmd_AccountCreateTheme8432c21f: func() TLObject {
		return NewTLAccountCreateTheme8432c21f()
	},
	Cmd_AccountDeleteAccount: func() TLObject {
		return NewTLAccountDeleteAccount()
	},
	Cmd_AccountDeleteSecureValue: func() TLObject {
		return NewTLAccountDeleteSecureValue()
	},
	Cmd_AccountFinishTakeoutSession: func() TLObject {
		return NewTLAccountFinishTakeoutSession()
	},
	Cmd_AccountGetAccountTTL: func() TLObject {
		return NewTLAccountGetAccountTTL()
	},
	Cmd_AccountGetAllSecureValues: func() TLObject {
		return NewTLAccountGetAllSecureValues()
	},
	Cmd_AccountGetAuthorizationForm: func() TLObject {
		return NewTLAccountGetAuthorizationForm()
	},
	Cmd_AccountGetAuthorizations: func() TLObject {
		return NewTLAccountGetAuthorizations()
	},
	Cmd_AccountGetAutoDownloadSettings: func() TLObject {
		return NewTLAccountGetAutoDownloadSettings()
	},
	Cmd_AccountGetContactSignUpNotification: func() TLObject {
		return NewTLAccountGetContactSignUpNotification()
	},
	Cmd_AccountGetContentSettings: func() TLObject {
		return NewTLAccountGetContentSettings()
	},
	Cmd_AccountGetGlobalPrivacySettings: func() TLObject {
		return NewTLAccountGetGlobalPrivacySettings()
	},
	Cmd_AccountGetMultiWallPapers: func() TLObject {
		return NewTLAccountGetMultiWallPapers()
	},
	Cmd_AccountGetNotifyExceptions: func() TLObject {
		return NewTLAccountGetNotifyExceptions()
	},
	Cmd_AccountGetNotifySettings: func() TLObject {
		return NewTLAccountGetNotifySettings()
	},
	Cmd_AccountGetPassword: func() TLObject {
		return NewTLAccountGetPassword()
	},
	Cmd_AccountGetPasswordSettings: func() TLObject {
		return NewTLAccountGetPasswordSettings()
	},
	Cmd_AccountGetPasswordSettings9cd4eaf9: func() TLObject {
		return NewTLAccountGetPasswordSettings9cd4eaf9()
	},
	Cmd_AccountGetPrivacy: func() TLObject {
		return NewTLAccountGetPrivacy()
	},
	Cmd_AccountGetSecureValue: func() TLObject {
		return NewTLAccountGetSecureValue()
	},
	Cmd_AccountGetTheme: func() TLObject {
		return NewTLAccountGetTheme()
	},
	Cmd_AccountGetThemes: func() TLObject {
		return NewTLAccountGetThemes()
	},
	Cmd_AccountGetTmpPassword: func() TLObject {
		return NewTLAccountGetTmpPassword()
	},
	Cmd_AccountGetTmpPassword449e0b51: func() TLObject {
		return NewTLAccountGetTmpPassword449e0b51()
	},
	Cmd_AccountGetWallPaper: func() TLObject {
		return NewTLAccountGetWallPaper()
	},
	Cmd_AccountGetWallPapers: func() TLObject {
		return NewTLAccountGetWallPapers()
	},
	Cmd_AccountGetWallPapersaabb1763: func() TLObject {
		return NewTLAccountGetWallPapersaabb1763()
	},
	Cmd_AccountGetWebAuthorizations: func() TLObject {
		return NewTLAccountGetWebAuthorizations()
	},
	Cmd_AccountInitTakeoutSession: func() TLObject {
		return NewTLAccountInitTakeoutSession()
	},
	Cmd_AccountInstallTheme: func() TLObject {
		return NewTLAccountInstallTheme()
	},
	Cmd_AccountInstallTheme7ae43737: func() TLObject {
		return NewTLAccountInstallTheme7ae43737()
	},
	Cmd_AccountInstallWallPaper: func() TLObject {
		return NewTLAccountInstallWallPaper()
	},
	Cmd_AccountInstallWallPaperfeed5769: func() TLObject {
		return NewTLAccountInstallWallPaperfeed5769()
	},
	Cmd_AccountRegisterDevice: func() TLObject {
		return NewTLAccountRegisterDevice()
	},
	Cmd_AccountRegisterDevice446c712c: func() TLObject {
		return NewTLAccountRegisterDevice446c712c()
	},
	Cmd_AccountRegisterDevice5cbea590: func() TLObject {
		return NewTLAccountRegisterDevice5cbea590()
	},
	Cmd_AccountRegisterDevice68976c6f: func() TLObject {
		return NewTLAccountRegisterDevice68976c6f()
	},
	Cmd_AccountReportPeer: func() TLObject {
		return NewTLAccountReportPeer()
	},
	Cmd_AccountResendPasswordEmail: func() TLObject {
		return NewTLAccountResendPasswordEmail()
	},
	Cmd_AccountResetAuthorization: func() TLObject {
		return NewTLAccountResetAuthorization()
	},
	Cmd_AccountResetNotifySettings: func() TLObject {
		return NewTLAccountResetNotifySettings()
	},
	Cmd_AccountResetWallPapers: func() TLObject {
		return NewTLAccountResetWallPapers()
	},
	Cmd_AccountResetWebAuthorization: func() TLObject {
		return NewTLAccountResetWebAuthorization()
	},
	Cmd_AccountResetWebAuthorizations: func() TLObject {
		return NewTLAccountResetWebAuthorizations()
	},
	Cmd_AccountSaveAutoDownloadSettings: func() TLObject {
		return NewTLAccountSaveAutoDownloadSettings()
	},
	Cmd_AccountSaveSecureValue: func() TLObject {
		return NewTLAccountSaveSecureValue()
	},
	Cmd_AccountSaveTheme: func() TLObject {
		return NewTLAccountSaveTheme()
	},
	Cmd_AccountSaveWallPaper: func() TLObject {
		return NewTLAccountSaveWallPaper()
	},
	Cmd_AccountSaveWallPaper6c5a5b37: func() TLObject {
		return NewTLAccountSaveWallPaper6c5a5b37()
	},
	Cmd_AccountSendChangePhoneCode: func() TLObject {
		return NewTLAccountSendChangePhoneCode()
	},
	Cmd_AccountSendChangePhoneCode82574ae5: func() TLObject {
		return NewTLAccountSendChangePhoneCode82574ae5()
	},
	Cmd_AccountSendConfirmPhoneCode: func() TLObject {
		return NewTLAccountSendConfirmPhoneCode()
	},
	Cmd_AccountSendConfirmPhoneCode1b3faa88: func() TLObject {
		return NewTLAccountSendConfirmPhoneCode1b3faa88()
	},
	Cmd_AccountSendVerifyEmailCode: func() TLObject {
		return NewTLAccountSendVerifyEmailCode()
	},
	Cmd_AccountSendVerifyPhoneCode: func() TLObject {
		return NewTLAccountSendVerifyPhoneCode()
	},
	Cmd_AccountSendVerifyPhoneCodea5a356f9: func() TLObject {
		return NewTLAccountSendVerifyPhoneCodea5a356f9()
	},
	Cmd_AccountSetAccountTTL: func() TLObject {
		return NewTLAccountSetAccountTTL()
	},
	Cmd_AccountSetContactSignUpNotification: func() TLObject {
		return NewTLAccountSetContactSignUpNotification()
	},
	Cmd_AccountSetContentSettings: func() TLObject {
		return NewTLAccountSetContentSettings()
	},
	Cmd_AccountSetGlobalPrivacySettings: func() TLObject {
		return NewTLAccountSetGlobalPrivacySettings()
	},
	Cmd_AccountSetPrivacy: func() TLObject {
		return NewTLAccountSetPrivacy()
	},
	Cmd_AccountUnregisterDevice: func() TLObject {
		return NewTLAccountUnregisterDevice()
	},
	Cmd_AccountUnregisterDevice3076c4bf: func() TLObject {
		return NewTLAccountUnregisterDevice3076c4bf()
	},
	Cmd_AccountUpdateDeviceLocked: func() TLObject {
		return NewTLAccountUpdateDeviceLocked()
	},
	Cmd_AccountUpdateNotifySettings: func() TLObject {
		return NewTLAccountUpdateNotifySettings()
	},
	Cmd_AccountUpdatePasswordSettings: func() TLObject {
		return NewTLAccountUpdatePasswordSettings()
	},
	Cmd_AccountUpdatePasswordSettingsa59b102f: func() TLObject {
		return NewTLAccountUpdatePasswordSettingsa59b102f()
	},
	Cmd_AccountUpdateProfile: func() TLObject {
		return NewTLAccountUpdateProfile()
	},
	Cmd_AccountUpdateStatus: func() TLObject {
		return NewTLAccountUpdateStatus()
	},
	Cmd_AccountUpdateTheme: func() TLObject {
		return NewTLAccountUpdateTheme()
	},
	Cmd_AccountUpdateTheme3b8ea202: func() TLObject {
		return NewTLAccountUpdateTheme3b8ea202()
	},
	Cmd_AccountUpdateTheme5cb367d5: func() TLObject {
		return NewTLAccountUpdateTheme5cb367d5()
	},
	Cmd_AccountUpdateUsername: func() TLObject {
		return NewTLAccountUpdateUsername()
	},
	Cmd_AccountUploadTheme: func() TLObject {
		return NewTLAccountUploadTheme()
	},
	Cmd_AccountUploadWallPaper: func() TLObject {
		return NewTLAccountUploadWallPaper()
	},
	Cmd_AccountUploadWallPaperdd853661: func() TLObject {
		return NewTLAccountUploadWallPaperdd853661()
	},
	Cmd_AccountVerifyEmail: func() TLObject {
		return NewTLAccountVerifyEmail()
	},
	Cmd_AccountVerifyPhone: func() TLObject {
		return NewTLAccountVerifyPhone()
	},
	Cmd_AuthAcceptLoginToken: func() TLObject {
		return NewTLAuthAcceptLoginToken()
	},
	Cmd_AuthBindTempAuthKey: func() TLObject {
		return NewTLAuthBindTempAuthKey()
	},
	Cmd_AuthCancelCode: func() TLObject {
		return NewTLAuthCancelCode()
	},
	Cmd_AuthCheckPassword: func() TLObject {
		return NewTLAuthCheckPassword()
	},
	Cmd_AuthCheckPasswordd18b4d16: func() TLObject {
		return NewTLAuthCheckPasswordd18b4d16()
	},
	Cmd_AuthCheckPhone: func() TLObject {
		return NewTLAuthCheckPhone()
	},
	Cmd_AuthDropTempAuthKeys: func() TLObject {
		return NewTLAuthDropTempAuthKeys()
	},
	Cmd_AuthExportAuthorization: func() TLObject {
		return NewTLAuthExportAuthorization()
	},
	Cmd_AuthExportLoginToken: func() TLObject {
		return NewTLAuthExportLoginToken()
	},
	Cmd_AuthImportAuthorization: func() TLObject {
		return NewTLAuthImportAuthorization()
	},
	Cmd_AuthImportBotAuthorization: func() TLObject {
		return NewTLAuthImportBotAuthorization()
	},
	Cmd_AuthImportLoginToken: func() TLObject {
		return NewTLAuthImportLoginToken()
	},
	Cmd_AuthLogOut: func() TLObject {
		return NewTLAuthLogOut()
	},
	Cmd_AuthRecoverPassword: func() TLObject {
		return NewTLAuthRecoverPassword()
	},
	Cmd_AuthRequestPasswordRecovery: func() TLObject {
		return NewTLAuthRequestPasswordRecovery()
	},
	Cmd_AuthResendCode: func() TLObject {
		return NewTLAuthResendCode()
	},
	Cmd_AuthResetAuthorizations: func() TLObject {
		return NewTLAuthResetAuthorizations()
	},
	Cmd_AuthSendCode: func() TLObject {
		return NewTLAuthSendCode()
	},
	Cmd_AuthSendCodea677244f: func() TLObject {
		return NewTLAuthSendCodea677244f()
	},
	Cmd_AuthSendCodeccfd70cf: func() TLObject {
		return NewTLAuthSendCodeccfd70cf()
	},
	Cmd_AuthSendInvites: func() TLObject {
		return NewTLAuthSendInvites()
	},
	Cmd_AuthSignIn: func() TLObject {
		return NewTLAuthSignIn()
	},
	Cmd_AuthSignUp: func() TLObject {
		return NewTLAuthSignUp()
	},
	Cmd_AuthSignUp80eee427: func() TLObject {
		return NewTLAuthSignUp80eee427()
	},
	Cmd_BotsAnswerWebhookJSONQuery: func() TLObject {
		return NewTLBotsAnswerWebhookJSONQuery()
	},
	Cmd_BotsSendCustomRequest: func() TLObject {
		return NewTLBotsSendCustomRequest()
	},
	Cmd_BotsSetBotCommands: func() TLObject {
		return NewTLBotsSetBotCommands()
	},
	Cmd_ChannelsCheckUsername: func() TLObject {
		return NewTLChannelsCheckUsername()
	},
	Cmd_ChannelsCreateChannel: func() TLObject {
		return NewTLChannelsCreateChannel()
	},
	Cmd_ChannelsCreateChannel3d5fb10f: func() TLObject {
		return NewTLChannelsCreateChannel3d5fb10f()
	},
	Cmd_ChannelsDeleteChannel: func() TLObject {
		return NewTLChannelsDeleteChannel()
	},
	Cmd_ChannelsDeleteHistory: func() TLObject {
		return NewTLChannelsDeleteHistory()
	},
	Cmd_ChannelsDeleteMessages: func() TLObject {
		return NewTLChannelsDeleteMessages()
	},
	Cmd_ChannelsDeleteUserHistory: func() TLObject {
		return NewTLChannelsDeleteUserHistory()
	},
	Cmd_ChannelsEditAbout: func() TLObject {
		return NewTLChannelsEditAbout()
	},
	Cmd_ChannelsEditAdmin: func() TLObject {
		return NewTLChannelsEditAdmin()
	},
	Cmd_ChannelsEditAdmin70f893ba: func() TLObject {
		return NewTLChannelsEditAdmin70f893ba()
	},
	Cmd_ChannelsEditAdmind33c8902: func() TLObject {
		return NewTLChannelsEditAdmind33c8902()
	},
	Cmd_ChannelsEditAdmineb7611d0: func() TLObject {
		return NewTLChannelsEditAdmineb7611d0()
	},
	Cmd_ChannelsEditBanned: func() TLObject {
		return NewTLChannelsEditBanned()
	},
	Cmd_ChannelsEditBanned72796912: func() TLObject {
		return NewTLChannelsEditBanned72796912()
	},
	Cmd_ChannelsEditCreator: func() TLObject {
		return NewTLChannelsEditCreator()
	},
	Cmd_ChannelsEditLocation: func() TLObject {
		return NewTLChannelsEditLocation()
	},
	Cmd_ChannelsEditPhoto: func() TLObject {
		return NewTLChannelsEditPhoto()
	},
	Cmd_ChannelsEditTitle: func() TLObject {
		return NewTLChannelsEditTitle()
	},
	Cmd_ChannelsExportInvite: func() TLObject {
		return NewTLChannelsExportInvite()
	},
	Cmd_ChannelsExportMessageLink: func() TLObject {
		return NewTLChannelsExportMessageLink()
	},
	Cmd_ChannelsExportMessageLinkceb77163: func() TLObject {
		return NewTLChannelsExportMessageLinkceb77163()
	},
	Cmd_ChannelsExportMessageLinke63fadeb: func() TLObject {
		return NewTLChannelsExportMessageLinke63fadeb()
	},
	Cmd_ChannelsGetAdminLog: func() TLObject {
		return NewTLChannelsGetAdminLog()
	},
	Cmd_ChannelsGetAdminedPublicChannels: func() TLObject {
		return NewTLChannelsGetAdminedPublicChannels()
	},
	Cmd_ChannelsGetAdminedPublicChannelsf8b036af: func() TLObject {
		return NewTLChannelsGetAdminedPublicChannelsf8b036af()
	},
	Cmd_ChannelsGetBroadcastsForDiscussion: func() TLObject {
		return NewTLChannelsGetBroadcastsForDiscussion()
	},
	Cmd_ChannelsGetChannels: func() TLObject {
		return NewTLChannelsGetChannels()
	},
	Cmd_ChannelsGetDialogs: func() TLObject {
		return NewTLChannelsGetDialogs()
	},
	Cmd_ChannelsGetFullChannel: func() TLObject {
		return NewTLChannelsGetFullChannel()
	},
	Cmd_ChannelsGetGroupsForDiscussion: func() TLObject {
		return NewTLChannelsGetGroupsForDiscussion()
	},
	Cmd_ChannelsGetImportantHistory: func() TLObject {
		return NewTLChannelsGetImportantHistory()
	},
	Cmd_ChannelsGetInactiveChannels: func() TLObject {
		return NewTLChannelsGetInactiveChannels()
	},
	Cmd_ChannelsGetLeftChannels: func() TLObject {
		return NewTLChannelsGetLeftChannels()
	},
	Cmd_ChannelsGetMessages: func() TLObject {
		return NewTLChannelsGetMessages()
	},
	Cmd_ChannelsGetMessagesad8c9a23: func() TLObject {
		return NewTLChannelsGetMessagesad8c9a23()
	},
	Cmd_ChannelsGetParticipant: func() TLObject {
		return NewTLChannelsGetParticipant()
	},
	Cmd_ChannelsGetParticipants: func() TLObject {
		return NewTLChannelsGetParticipants()
	},
	Cmd_ChannelsGetParticipants123e05e9: func() TLObject {
		return NewTLChannelsGetParticipants123e05e9()
	},
	Cmd_ChannelsInviteToChannel: func() TLObject {
		return NewTLChannelsInviteToChannel()
	},
	Cmd_ChannelsJoinChannel: func() TLObject {
		return NewTLChannelsJoinChannel()
	},
	Cmd_ChannelsKickFromChannel: func() TLObject {
		return NewTLChannelsKickFromChannel()
	},
	Cmd_ChannelsLeaveChannel: func() TLObject {
		return NewTLChannelsLeaveChannel()
	},
	Cmd_ChannelsReadHistory: func() TLObject {
		return NewTLChannelsReadHistory()
	},
	Cmd_ChannelsReadMessageContents: func() TLObject {
		return NewTLChannelsReadMessageContents()
	},
	Cmd_ChannelsReportSpam: func() TLObject {
		return NewTLChannelsReportSpam()
	},
	Cmd_ChannelsSetDiscussionGroup: func() TLObject {
		return NewTLChannelsSetDiscussionGroup()
	},
	Cmd_ChannelsSetStickers: func() TLObject {
		return NewTLChannelsSetStickers()
	},
	Cmd_ChannelsToggleComments: func() TLObject {
		return NewTLChannelsToggleComments()
	},
	Cmd_ChannelsToggleInvites: func() TLObject {
		return NewTLChannelsToggleInvites()
	},
	Cmd_ChannelsTogglePreHistoryHidden: func() TLObject {
		return NewTLChannelsTogglePreHistoryHidden()
	},
	Cmd_ChannelsToggleSignatures: func() TLObject {
		return NewTLChannelsToggleSignatures()
	},
	Cmd_ChannelsToggleSlowMode: func() TLObject {
		return NewTLChannelsToggleSlowMode()
	},
	Cmd_ChannelsUpdatePinnedMessage: func() TLObject {
		return NewTLChannelsUpdatePinnedMessage()
	},
	Cmd_ChannelsUpdateUsername: func() TLObject {
		return NewTLChannelsUpdateUsername()
	},
	Cmd_ContactsAcceptContact: func() TLObject {
		return NewTLContactsAcceptContact()
	},
	Cmd_ContactsAddContact: func() TLObject {
		return NewTLContactsAddContact()
	},
	Cmd_ContactsBlock: func() TLObject {
		return NewTLContactsBlock()
	},
	Cmd_ContactsBlock68cc1411: func() TLObject {
		return NewTLContactsBlock68cc1411()
	},
	Cmd_ContactsBlockFromReplies: func() TLObject {
		return NewTLContactsBlockFromReplies()
	},
	Cmd_ContactsDeleteByPhones: func() TLObject {
		return NewTLContactsDeleteByPhones()
	},
	Cmd_ContactsDeleteContact: func() TLObject {
		return NewTLContactsDeleteContact()
	},
	Cmd_ContactsDeleteContacts: func() TLObject {
		return NewTLContactsDeleteContacts()
	},
	Cmd_ContactsDeleteContacts96a0e00: func() TLObject {
		return NewTLContactsDeleteContacts96a0e00()
	},
	Cmd_ContactsExportCard: func() TLObject {
		return NewTLContactsExportCard()
	},
	Cmd_ContactsGetBlocked: func() TLObject {
		return NewTLContactsGetBlocked()
	},
	Cmd_ContactsGetContactIDs: func() TLObject {
		return NewTLContactsGetContactIDs()
	},
	Cmd_ContactsGetContacts: func() TLObject {
		return NewTLContactsGetContacts()
	},
	Cmd_ContactsGetContacts22c6aa08: func() TLObject {
		return NewTLContactsGetContacts22c6aa08()
	},
	Cmd_ContactsGetLocated: func() TLObject {
		return NewTLContactsGetLocated()
	},
	Cmd_ContactsGetLocatedd348bc44: func() TLObject {
		return NewTLContactsGetLocatedd348bc44()
	},
	Cmd_ContactsGetSaved: func() TLObject {
		return NewTLContactsGetSaved()
	},
	Cmd_ContactsGetStatuses: func() TLObject {
		return NewTLContactsGetStatuses()
	},
	Cmd_ContactsGetTopPeers: func() TLObject {
		return NewTLContactsGetTopPeers()
	},
	Cmd_ContactsImportCard: func() TLObject {
		return NewTLContactsImportCard()
	},
	Cmd_ContactsImportContacts: func() TLObject {
		return NewTLContactsImportContacts()
	},
	Cmd_ContactsImportContactsda30b32d: func() TLObject {
		return NewTLContactsImportContactsda30b32d()
	},
	Cmd_ContactsResetSaved: func() TLObject {
		return NewTLContactsResetSaved()
	},
	Cmd_ContactsResetTopPeerRating: func() TLObject {
		return NewTLContactsResetTopPeerRating()
	},
	Cmd_ContactsResolveUsername: func() TLObject {
		return NewTLContactsResolveUsername()
	},
	Cmd_ContactsSearch: func() TLObject {
		return NewTLContactsSearch()
	},
	Cmd_ContactsToggleTopPeers: func() TLObject {
		return NewTLContactsToggleTopPeers()
	},
	Cmd_ContactsUnblock: func() TLObject {
		return NewTLContactsUnblock()
	},
	Cmd_ContactsUnblockbea65d50: func() TLObject {
		return NewTLContactsUnblockbea65d50()
	},
	Cmd_FoldersDeleteFolder: func() TLObject {
		return NewTLFoldersDeleteFolder()
	},
	Cmd_FoldersEditPeerFolders: func() TLObject {
		return NewTLFoldersEditPeerFolders()
	},
	Cmd_HelpAcceptTermsOfService: func() TLObject {
		return NewTLHelpAcceptTermsOfService()
	},
	Cmd_HelpDismissSuggestion: func() TLObject {
		return NewTLHelpDismissSuggestion()
	},
	Cmd_HelpEditUserInfo: func() TLObject {
		return NewTLHelpEditUserInfo()
	},
	Cmd_HelpGetAppChangelog: func() TLObject {
		return NewTLHelpGetAppChangelog()
	},
	Cmd_HelpGetAppChangelog5bab7fb2: func() TLObject {
		return NewTLHelpGetAppChangelog5bab7fb2()
	},
	Cmd_HelpGetAppConfig: func() TLObject {
		return NewTLHelpGetAppConfig()
	},
	Cmd_HelpGetAppUpdate: func() TLObject {
		return NewTLHelpGetAppUpdate()
	},
	Cmd_HelpGetAppUpdate522d5a7d: func() TLObject {
		return NewTLHelpGetAppUpdate522d5a7d()
	},
	Cmd_HelpGetAppUpdatec812ac7e: func() TLObject {
		return NewTLHelpGetAppUpdatec812ac7e()
	},
	Cmd_HelpGetCdnConfig: func() TLObject {
		return NewTLHelpGetCdnConfig()
	},
	Cmd_HelpGetConfig: func() TLObject {
		return NewTLHelpGetConfig()
	},
	Cmd_HelpGetCountriesList: func() TLObject {
		return NewTLHelpGetCountriesList()
	},
	Cmd_HelpGetDeepLinkInfo: func() TLObject {
		return NewTLHelpGetDeepLinkInfo()
	},
	Cmd_HelpGetInviteText: func() TLObject {
		return NewTLHelpGetInviteText()
	},
	Cmd_HelpGetInviteTexta4a95186: func() TLObject {
		return NewTLHelpGetInviteTexta4a95186()
	},
	Cmd_HelpGetNearestDc: func() TLObject {
		return NewTLHelpGetNearestDc()
	},
	Cmd_HelpGetPassportConfig: func() TLObject {
		return NewTLHelpGetPassportConfig()
	},
	Cmd_HelpGetPromoData: func() TLObject {
		return NewTLHelpGetPromoData()
	},
	Cmd_HelpGetProxyData: func() TLObject {
		return NewTLHelpGetProxyData()
	},
	Cmd_HelpGetRecentMeUrls: func() TLObject {
		return NewTLHelpGetRecentMeUrls()
	},
	Cmd_HelpGetSupport: func() TLObject {
		return NewTLHelpGetSupport()
	},
	Cmd_HelpGetSupportName: func() TLObject {
		return NewTLHelpGetSupportName()
	},
	Cmd_HelpGetTermsOfService: func() TLObject {
		return NewTLHelpGetTermsOfService()
	},
	Cmd_HelpGetTermsOfService37d78f83: func() TLObject {
		return NewTLHelpGetTermsOfService37d78f83()
	},
	Cmd_HelpGetTermsOfServiceUpdate: func() TLObject {
		return NewTLHelpGetTermsOfServiceUpdate()
	},
	Cmd_HelpGetUserInfo: func() TLObject {
		return NewTLHelpGetUserInfo()
	},
	Cmd_HelpHidePromoData: func() TLObject {
		return NewTLHelpHidePromoData()
	},
	Cmd_HelpSaveAppLog: func() TLObject {
		return NewTLHelpSaveAppLog()
	},
	Cmd_HelpSetBotUpdatesStatus: func() TLObject {
		return NewTLHelpSetBotUpdatesStatus()
	},
	Cmd_HelpTest: func() TLObject {
		return NewTLHelpTest()
	},
	Cmd_LangpackGetDifference: func() TLObject {
		return NewTLLangpackGetDifference()
	},
	Cmd_LangpackGetDifference9d51e814: func() TLObject {
		return NewTLLangpackGetDifference9d51e814()
	},
	Cmd_LangpackGetDifferencecd984aa5: func() TLObject {
		return NewTLLangpackGetDifferencecd984aa5()
	},
	Cmd_LangpackGetLangPack: func() TLObject {
		return NewTLLangpackGetLangPack()
	},
	Cmd_LangpackGetLangPackf2f2330a: func() TLObject {
		return NewTLLangpackGetLangPackf2f2330a()
	},
	Cmd_LangpackGetLanguage: func() TLObject {
		return NewTLLangpackGetLanguage()
	},
	Cmd_LangpackGetLanguages: func() TLObject {
		return NewTLLangpackGetLanguages()
	},
	Cmd_LangpackGetLanguages42c6978f: func() TLObject {
		return NewTLLangpackGetLanguages42c6978f()
	},
	Cmd_LangpackGetStrings: func() TLObject {
		return NewTLLangpackGetStrings()
	},
	Cmd_LangpackGetStringsefea3803: func() TLObject {
		return NewTLLangpackGetStringsefea3803()
	},
	Cmd_MessagesAcceptEncryption: func() TLObject {
		return NewTLMessagesAcceptEncryption()
	},
	Cmd_MessagesAcceptUrlAuth: func() TLObject {
		return NewTLMessagesAcceptUrlAuth()
	},
	Cmd_MessagesAddChatUser: func() TLObject {
		return NewTLMessagesAddChatUser()
	},
	Cmd_MessagesCheckChatInvite: func() TLObject {
		return NewTLMessagesCheckChatInvite()
	},
	Cmd_MessagesClearAllDrafts: func() TLObject {
		return NewTLMessagesClearAllDrafts()
	},
	Cmd_MessagesClearRecentStickers: func() TLObject {
		return NewTLMessagesClearRecentStickers()
	},
	Cmd_MessagesCreateChat: func() TLObject {
		return NewTLMessagesCreateChat()
	},
	Cmd_MessagesDeleteChatUser: func() TLObject {
		return NewTLMessagesDeleteChatUser()
	},
	Cmd_MessagesDeleteHistory: func() TLObject {
		return NewTLMessagesDeleteHistory()
	},
	Cmd_MessagesDeleteHistory1c015b09: func() TLObject {
		return NewTLMessagesDeleteHistory1c015b09()
	},
	Cmd_MessagesDeleteMessages: func() TLObject {
		return NewTLMessagesDeleteMessages()
	},
	Cmd_MessagesDeleteMessagesa5f18925: func() TLObject {
		return NewTLMessagesDeleteMessagesa5f18925()
	},
	Cmd_MessagesDeleteScheduledMessages: func() TLObject {
		return NewTLMessagesDeleteScheduledMessages()
	},
	Cmd_MessagesDiscardEncryption: func() TLObject {
		return NewTLMessagesDiscardEncryption()
	},
	Cmd_MessagesEditChatAbout: func() TLObject {
		return NewTLMessagesEditChatAbout()
	},
	Cmd_MessagesEditChatAdmin: func() TLObject {
		return NewTLMessagesEditChatAdmin()
	},
	Cmd_MessagesEditChatDefaultBannedRights: func() TLObject {
		return NewTLMessagesEditChatDefaultBannedRights()
	},
	Cmd_MessagesEditChatPhoto: func() TLObject {
		return NewTLMessagesEditChatPhoto()
	},
	Cmd_MessagesEditChatTitle: func() TLObject {
		return NewTLMessagesEditChatTitle()
	},
	Cmd_MessagesEditInlineBotMessage: func() TLObject {
		return NewTLMessagesEditInlineBotMessage()
	},
	Cmd_MessagesEditInlineBotMessage83557dba: func() TLObject {
		return NewTLMessagesEditInlineBotMessage83557dba()
	},
	Cmd_MessagesEditInlineBotMessageadc3e828: func() TLObject {
		return NewTLMessagesEditInlineBotMessageadc3e828()
	},
	Cmd_MessagesEditMessage: func() TLObject {
		return NewTLMessagesEditMessage()
	},
	Cmd_MessagesEditMessage48f71778: func() TLObject {
		return NewTLMessagesEditMessage48f71778()
	},
	Cmd_MessagesEditMessagec000e4c8: func() TLObject {
		return NewTLMessagesEditMessagec000e4c8()
	},
	Cmd_MessagesEditMessaged116f31e: func() TLObject {
		return NewTLMessagesEditMessaged116f31e()
	},
	Cmd_MessagesExportChatInvite: func() TLObject {
		return NewTLMessagesExportChatInvite()
	},
	Cmd_MessagesExportChatInvitedf7534c: func() TLObject {
		return NewTLMessagesExportChatInvitedf7534c()
	},
	Cmd_MessagesFaveSticker: func() TLObject {
		return NewTLMessagesFaveSticker()
	},
	Cmd_MessagesForwardMessage: func() TLObject {
		return NewTLMessagesForwardMessage()
	},
	Cmd_MessagesForwardMessages: func() TLObject {
		return NewTLMessagesForwardMessages()
	},
	Cmd_MessagesForwardMessagesd9fee60e: func() TLObject {
		return NewTLMessagesForwardMessagesd9fee60e()
	},
	Cmd_MessagesGetAllChats: func() TLObject {
		return NewTLMessagesGetAllChats()
	},
	Cmd_MessagesGetAllDrafts: func() TLObject {
		return NewTLMessagesGetAllDrafts()
	},
	Cmd_MessagesGetAllStickers: func() TLObject {
		return NewTLMessagesGetAllStickers()
	},
	Cmd_MessagesGetArchivedStickers: func() TLObject {
		return NewTLMessagesGetArchivedStickers()
	},
	Cmd_MessagesGetAttachedStickers: func() TLObject {
		return NewTLMessagesGetAttachedStickers()
	},
	Cmd_MessagesGetBotCallbackAnswer: func() TLObject {
		return NewTLMessagesGetBotCallbackAnswer()
	},
	Cmd_MessagesGetBotCallbackAnswer9342ca07: func() TLObject {
		return NewTLMessagesGetBotCallbackAnswer9342ca07()
	},
	Cmd_MessagesGetBotCallbackAnswera6e94f04: func() TLObject {
		return NewTLMessagesGetBotCallbackAnswera6e94f04()
	},
	Cmd_MessagesGetChats: func() TLObject {
		return NewTLMessagesGetChats()
	},
	Cmd_MessagesGetCommonChats: func() TLObject {
		return NewTLMessagesGetCommonChats()
	},
	Cmd_MessagesGetDhConfig: func() TLObject {
		return NewTLMessagesGetDhConfig()
	},
	Cmd_MessagesGetDialogFilters: func() TLObject {
		return NewTLMessagesGetDialogFilters()
	},
	Cmd_MessagesGetDialogUnreadMarks: func() TLObject {
		return NewTLMessagesGetDialogUnreadMarks()
	},
	Cmd_MessagesGetDialogs: func() TLObject {
		return NewTLMessagesGetDialogs()
	},
	Cmd_MessagesGetDialogs6b47f94d: func() TLObject {
		return NewTLMessagesGetDialogs6b47f94d()
	},
	Cmd_MessagesGetDialogsa0ee3b73: func() TLObject {
		return NewTLMessagesGetDialogsa0ee3b73()
	},
	Cmd_MessagesGetDialogsb098aee6: func() TLObject {
		return NewTLMessagesGetDialogsb098aee6()
	},
	Cmd_MessagesGetDiscussionMessage: func() TLObject {
		return NewTLMessagesGetDiscussionMessage()
	},
	Cmd_MessagesGetDocumentByHash: func() TLObject {
		return NewTLMessagesGetDocumentByHash()
	},
	Cmd_MessagesGetEmojiKeywords: func() TLObject {
		return NewTLMessagesGetEmojiKeywords()
	},
	Cmd_MessagesGetEmojiKeywordsDifference: func() TLObject {
		return NewTLMessagesGetEmojiKeywordsDifference()
	},
	Cmd_MessagesGetEmojiKeywordsLanguages: func() TLObject {
		return NewTLMessagesGetEmojiKeywordsLanguages()
	},
	Cmd_MessagesGetEmojiURL: func() TLObject {
		return NewTLMessagesGetEmojiURL()
	},
	Cmd_MessagesGetFavedStickers: func() TLObject {
		return NewTLMessagesGetFavedStickers()
	},
	Cmd_MessagesGetFeaturedStickers: func() TLObject {
		return NewTLMessagesGetFeaturedStickers()
	},
	Cmd_MessagesGetFullChat: func() TLObject {
		return NewTLMessagesGetFullChat()
	},
	Cmd_MessagesGetGameHighScores: func() TLObject {
		return NewTLMessagesGetGameHighScores()
	},
	Cmd_MessagesGetHistory: func() TLObject {
		return NewTLMessagesGetHistory()
	},
	Cmd_MessagesGetHistorydcbb8260: func() TLObject {
		return NewTLMessagesGetHistorydcbb8260()
	},
	Cmd_MessagesGetInlineBotResults: func() TLObject {
		return NewTLMessagesGetInlineBotResults()
	},
	Cmd_MessagesGetInlineGameHighScores: func() TLObject {
		return NewTLMessagesGetInlineGameHighScores()
	},
	Cmd_MessagesGetMaskStickers: func() TLObject {
		return NewTLMessagesGetMaskStickers()
	},
	Cmd_MessagesGetMessageEditData: func() TLObject {
		return NewTLMessagesGetMessageEditData()
	},
	Cmd_MessagesGetMessages: func() TLObject {
		return NewTLMessagesGetMessages()
	},
	Cmd_MessagesGetMessages63c66506: func() TLObject {
		return NewTLMessagesGetMessages63c66506()
	},
	Cmd_MessagesGetMessagesViews: func() TLObject {
		return NewTLMessagesGetMessagesViews()
	},
	Cmd_MessagesGetMessagesViews5784d3e1: func() TLObject {
		return NewTLMessagesGetMessagesViews5784d3e1()
	},
	Cmd_MessagesGetOldFeaturedStickers: func() TLObject {
		return NewTLMessagesGetOldFeaturedStickers()
	},
	Cmd_MessagesGetOnlines: func() TLObject {
		return NewTLMessagesGetOnlines()
	},
	Cmd_MessagesGetPeerDialogs: func() TLObject {
		return NewTLMessagesGetPeerDialogs()
	},
	Cmd_MessagesGetPeerDialogse470bcfd: func() TLObject {
		return NewTLMessagesGetPeerDialogse470bcfd()
	},
	Cmd_MessagesGetPeerSettings: func() TLObject {
		return NewTLMessagesGetPeerSettings()
	},
	Cmd_MessagesGetPinnedDialogs: func() TLObject {
		return NewTLMessagesGetPinnedDialogs()
	},
	Cmd_MessagesGetPinnedDialogsd6b94df2: func() TLObject {
		return NewTLMessagesGetPinnedDialogsd6b94df2()
	},
	Cmd_MessagesGetPollResults: func() TLObject {
		return NewTLMessagesGetPollResults()
	},
	Cmd_MessagesGetPollVotes: func() TLObject {
		return NewTLMessagesGetPollVotes()
	},
	Cmd_MessagesGetRecentLocations: func() TLObject {
		return NewTLMessagesGetRecentLocations()
	},
	Cmd_MessagesGetRecentLocationsbbc45b09: func() TLObject {
		return NewTLMessagesGetRecentLocationsbbc45b09()
	},
	Cmd_MessagesGetRecentStickers: func() TLObject {
		return NewTLMessagesGetRecentStickers()
	},
	Cmd_MessagesGetReplies: func() TLObject {
		return NewTLMessagesGetReplies()
	},
	Cmd_MessagesGetSavedGifs: func() TLObject {
		return NewTLMessagesGetSavedGifs()
	},
	Cmd_MessagesGetScheduledHistory: func() TLObject {
		return NewTLMessagesGetScheduledHistory()
	},
	Cmd_MessagesGetScheduledMessages: func() TLObject {
		return NewTLMessagesGetScheduledMessages()
	},
	Cmd_MessagesGetSearchCounters: func() TLObject {
		return NewTLMessagesGetSearchCounters()
	},
	Cmd_MessagesGetSplitRanges: func() TLObject {
		return NewTLMessagesGetSplitRanges()
	},
	Cmd_MessagesGetStatsURL: func() TLObject {
		return NewTLMessagesGetStatsURL()
	},
	Cmd_MessagesGetStatsURL812c2ae6: func() TLObject {
		return NewTLMessagesGetStatsURL812c2ae6()
	},
	Cmd_MessagesGetStickerSet: func() TLObject {
		return NewTLMessagesGetStickerSet()
	},
	Cmd_MessagesGetStickers: func() TLObject {
		return NewTLMessagesGetStickers()
	},
	Cmd_MessagesGetStickers43d4f2c: func() TLObject {
		return NewTLMessagesGetStickers43d4f2c()
	},
	Cmd_MessagesGetSuggestedDialogFilters: func() TLObject {
		return NewTLMessagesGetSuggestedDialogFilters()
	},
	Cmd_MessagesGetUnreadMentions: func() TLObject {
		return NewTLMessagesGetUnreadMentions()
	},
	Cmd_MessagesGetWebPage: func() TLObject {
		return NewTLMessagesGetWebPage()
	},
	Cmd_MessagesGetWebPagePreview: func() TLObject {
		return NewTLMessagesGetWebPagePreview()
	},
	Cmd_MessagesGetWebPagePreview8b68b0cc: func() TLObject {
		return NewTLMessagesGetWebPagePreview8b68b0cc()
	},
	Cmd_MessagesHidePeerSettingsBar: func() TLObject {
		return NewTLMessagesHidePeerSettingsBar()
	},
	Cmd_MessagesHideReportSpam: func() TLObject {
		return NewTLMessagesHideReportSpam()
	},
	Cmd_MessagesImportChatInvite: func() TLObject {
		return NewTLMessagesImportChatInvite()
	},
	Cmd_MessagesInstallStickerSet: func() TLObject {
		return NewTLMessagesInstallStickerSet()
	},
	Cmd_MessagesInstallStickerSet7b30c3a6: func() TLObject {
		return NewTLMessagesInstallStickerSet7b30c3a6()
	},
	Cmd_MessagesMarkDialogUnread: func() TLObject {
		return NewTLMessagesMarkDialogUnread()
	},
	Cmd_MessagesMigrateChat: func() TLObject {
		return NewTLMessagesMigrateChat()
	},
	Cmd_MessagesReadDiscussion: func() TLObject {
		return NewTLMessagesReadDiscussion()
	},
	Cmd_MessagesReadEncryptedHistory: func() TLObject {
		return NewTLMessagesReadEncryptedHistory()
	},
	Cmd_MessagesReadFeaturedStickers: func() TLObject {
		return NewTLMessagesReadFeaturedStickers()
	},
	Cmd_MessagesReadHistory: func() TLObject {
		return NewTLMessagesReadHistory()
	},
	Cmd_MessagesReadHistoryb04f2510: func() TLObject {
		return NewTLMessagesReadHistoryb04f2510()
	},
	Cmd_MessagesReadMentions: func() TLObject {
		return NewTLMessagesReadMentions()
	},
	Cmd_MessagesReadMessageContents: func() TLObject {
		return NewTLMessagesReadMessageContents()
	},
	Cmd_MessagesReceivedMessages: func() TLObject {
		return NewTLMessagesReceivedMessages()
	},
	Cmd_MessagesReceivedQueue: func() TLObject {
		return NewTLMessagesReceivedQueue()
	},
	Cmd_MessagesReorderPinnedDialogs: func() TLObject {
		return NewTLMessagesReorderPinnedDialogs()
	},
	Cmd_MessagesReorderPinnedDialogs3b1adf37: func() TLObject {
		return NewTLMessagesReorderPinnedDialogs3b1adf37()
	},
	Cmd_MessagesReorderPinnedDialogs5b51d63f: func() TLObject {
		return NewTLMessagesReorderPinnedDialogs5b51d63f()
	},
	Cmd_MessagesReorderStickerSets: func() TLObject {
		return NewTLMessagesReorderStickerSets()
	},
	Cmd_MessagesReorderStickerSets9fcfbc30: func() TLObject {
		return NewTLMessagesReorderStickerSets9fcfbc30()
	},
	Cmd_MessagesReport: func() TLObject {
		return NewTLMessagesReport()
	},
	Cmd_MessagesReportEncryptedSpam: func() TLObject {
		return NewTLMessagesReportEncryptedSpam()
	},
	Cmd_MessagesReportSpam: func() TLObject {
		return NewTLMessagesReportSpam()
	},
	Cmd_MessagesRequestEncryption: func() TLObject {
		return NewTLMessagesRequestEncryption()
	},
	Cmd_MessagesRequestUrlAuth: func() TLObject {
		return NewTLMessagesRequestUrlAuth()
	},
	Cmd_MessagesSaveDraft: func() TLObject {
		return NewTLMessagesSaveDraft()
	},
	Cmd_MessagesSaveGif: func() TLObject {
		return NewTLMessagesSaveGif()
	},
	Cmd_MessagesSaveRecentSticker: func() TLObject {
		return NewTLMessagesSaveRecentSticker()
	},
	Cmd_MessagesSaveRecentSticker348e39bf: func() TLObject {
		return NewTLMessagesSaveRecentSticker348e39bf()
	},
	Cmd_MessagesSearch: func() TLObject {
		return NewTLMessagesSearch()
	},
	Cmd_MessagesSearch4e17810b: func() TLObject {
		return NewTLMessagesSearch4e17810b()
	},
	Cmd_MessagesSearch8614ef68: func() TLObject {
		return NewTLMessagesSearch8614ef68()
	},
	Cmd_MessagesSearchGifs: func() TLObject {
		return NewTLMessagesSearchGifs()
	},
	Cmd_MessagesSearchGlobal: func() TLObject {
		return NewTLMessagesSearchGlobal()
	},
	Cmd_MessagesSearchGlobal4bc6589a: func() TLObject {
		return NewTLMessagesSearchGlobal4bc6589a()
	},
	Cmd_MessagesSearchGlobalbf7225a4: func() TLObject {
		return NewTLMessagesSearchGlobalbf7225a4()
	},
	Cmd_MessagesSearchGlobalf79c611: func() TLObject {
		return NewTLMessagesSearchGlobalf79c611()
	},
	Cmd_MessagesSearchStickerSets: func() TLObject {
		return NewTLMessagesSearchStickerSets()
	},
	Cmd_MessagesSearchc352eec: func() TLObject {
		return NewTLMessagesSearchc352eec()
	},
	Cmd_MessagesSearchd4569248: func() TLObject {
		return NewTLMessagesSearchd4569248()
	},
	Cmd_MessagesSearchf288a275: func() TLObject {
		return NewTLMessagesSearchf288a275()
	},
	Cmd_MessagesSendBroadcast: func() TLObject {
		return NewTLMessagesSendBroadcast()
	},
	Cmd_MessagesSendEncrypted: func() TLObject {
		return NewTLMessagesSendEncrypted()
	},
	Cmd_MessagesSendEncrypted44fa7a15: func() TLObject {
		return NewTLMessagesSendEncrypted44fa7a15()
	},
	Cmd_MessagesSendEncryptedFile: func() TLObject {
		return NewTLMessagesSendEncryptedFile()
	},
	Cmd_MessagesSendEncryptedFile5559481d: func() TLObject {
		return NewTLMessagesSendEncryptedFile5559481d()
	},
	Cmd_MessagesSendEncryptedService: func() TLObject {
		return NewTLMessagesSendEncryptedService()
	},
	Cmd_MessagesSendInlineBotResult: func() TLObject {
		return NewTLMessagesSendInlineBotResult()
	},
	Cmd_MessagesSendInlineBotResult220815b0: func() TLObject {
		return NewTLMessagesSendInlineBotResult220815b0()
	},
	Cmd_MessagesSendMedia: func() TLObject {
		return NewTLMessagesSendMedia()
	},
	Cmd_MessagesSendMedia3491eba9: func() TLObject {
		return NewTLMessagesSendMedia3491eba9()
	},
	Cmd_MessagesSendMediab8d1262b: func() TLObject {
		return NewTLMessagesSendMediab8d1262b()
	},
	Cmd_MessagesSendMessage: func() TLObject {
		return NewTLMessagesSendMessage()
	},
	Cmd_MessagesSendMessage520c3870: func() TLObject {
		return NewTLMessagesSendMessage520c3870()
	},
	Cmd_MessagesSendMultiMedia: func() TLObject {
		return NewTLMessagesSendMultiMedia()
	},
	Cmd_MessagesSendMultiMediacc0110cb: func() TLObject {
		return NewTLMessagesSendMultiMediacc0110cb()
	},
	Cmd_MessagesSendScheduledMessages: func() TLObject {
		return NewTLMessagesSendScheduledMessages()
	},
	Cmd_MessagesSendScreenshotNotification: func() TLObject {
		return NewTLMessagesSendScreenshotNotification()
	},
	Cmd_MessagesSendVote: func() TLObject {
		return NewTLMessagesSendVote()
	},
	Cmd_MessagesSetBotCallbackAnswer: func() TLObject {
		return NewTLMessagesSetBotCallbackAnswer()
	},
	Cmd_MessagesSetBotCallbackAnswer481c591a: func() TLObject {
		return NewTLMessagesSetBotCallbackAnswer481c591a()
	},
	Cmd_MessagesSetBotPrecheckoutResults: func() TLObject {
		return NewTLMessagesSetBotPrecheckoutResults()
	},
	Cmd_MessagesSetBotShippingResults: func() TLObject {
		return NewTLMessagesSetBotShippingResults()
	},
	Cmd_MessagesSetEncryptedTyping: func() TLObject {
		return NewTLMessagesSetEncryptedTyping()
	},
	Cmd_MessagesSetGameScore: func() TLObject {
		return NewTLMessagesSetGameScore()
	},
	Cmd_MessagesSetInlineBotResults: func() TLObject {
		return NewTLMessagesSetInlineBotResults()
	},
	Cmd_MessagesSetInlineGameScore: func() TLObject {
		return NewTLMessagesSetInlineGameScore()
	},
	Cmd_MessagesSetTyping: func() TLObject {
		return NewTLMessagesSetTyping()
	},
	Cmd_MessagesSetTyping58943ee2: func() TLObject {
		return NewTLMessagesSetTyping58943ee2()
	},
	Cmd_MessagesStartBot: func() TLObject {
		return NewTLMessagesStartBot()
	},
	Cmd_MessagesToggleChatAdmins: func() TLObject {
		return NewTLMessagesToggleChatAdmins()
	},
	Cmd_MessagesToggleDialogPin: func() TLObject {
		return NewTLMessagesToggleDialogPin()
	},
	Cmd_MessagesToggleDialogPina731e257: func() TLObject {
		return NewTLMessagesToggleDialogPina731e257()
	},
	Cmd_MessagesToggleStickerSets: func() TLObject {
		return NewTLMessagesToggleStickerSets()
	},
	Cmd_MessagesUninstallStickerSet: func() TLObject {
		return NewTLMessagesUninstallStickerSet()
	},
	Cmd_MessagesUnpinAllMessages: func() TLObject {
		return NewTLMessagesUnpinAllMessages()
	},
	Cmd_MessagesUpdateDialogFilter: func() TLObject {
		return NewTLMessagesUpdateDialogFilter()
	},
	Cmd_MessagesUpdateDialogFiltersOrder: func() TLObject {
		return NewTLMessagesUpdateDialogFiltersOrder()
	},
	Cmd_MessagesUpdatePinnedMessage: func() TLObject {
		return NewTLMessagesUpdatePinnedMessage()
	},
	Cmd_MessagesUploadEncryptedFile: func() TLObject {
		return NewTLMessagesUploadEncryptedFile()
	},
	Cmd_MessagesUploadMedia: func() TLObject {
		return NewTLMessagesUploadMedia()
	},
	Cmd_PaymentsClearSavedInfo: func() TLObject {
		return NewTLPaymentsClearSavedInfo()
	},
	Cmd_PaymentsGetBankCardData: func() TLObject {
		return NewTLPaymentsGetBankCardData()
	},
	Cmd_PaymentsGetPaymentForm: func() TLObject {
		return NewTLPaymentsGetPaymentForm()
	},
	Cmd_PaymentsGetPaymentReceipt: func() TLObject {
		return NewTLPaymentsGetPaymentReceipt()
	},
	Cmd_PaymentsGetSavedInfo: func() TLObject {
		return NewTLPaymentsGetSavedInfo()
	},
	Cmd_PaymentsSendPaymentForm: func() TLObject {
		return NewTLPaymentsSendPaymentForm()
	},
	Cmd_PaymentsValidateRequestedInfo: func() TLObject {
		return NewTLPaymentsValidateRequestedInfo()
	},
	Cmd_PhoneAcceptCall: func() TLObject {
		return NewTLPhoneAcceptCall()
	},
	Cmd_PhoneCheckGroupCall: func() TLObject {
		return NewTLPhoneCheckGroupCall()
	},
	Cmd_PhoneConfirmCall: func() TLObject {
		return NewTLPhoneConfirmCall()
	},
	Cmd_PhoneCreateGroupCall: func() TLObject {
		return NewTLPhoneCreateGroupCall()
	},
	Cmd_PhoneDiscardCall: func() TLObject {
		return NewTLPhoneDiscardCall()
	},
	Cmd_PhoneDiscardCallb2cbc1c0: func() TLObject {
		return NewTLPhoneDiscardCallb2cbc1c0()
	},
	Cmd_PhoneDiscardGroupCall: func() TLObject {
		return NewTLPhoneDiscardGroupCall()
	},
	Cmd_PhoneEditGroupCallMember: func() TLObject {
		return NewTLPhoneEditGroupCallMember()
	},
	Cmd_PhoneGetCallConfig: func() TLObject {
		return NewTLPhoneGetCallConfig()
	},
	Cmd_PhoneGetGroupCall: func() TLObject {
		return NewTLPhoneGetGroupCall()
	},
	Cmd_PhoneGetGroupParticipants: func() TLObject {
		return NewTLPhoneGetGroupParticipants()
	},
	Cmd_PhoneInviteToGroupCall: func() TLObject {
		return NewTLPhoneInviteToGroupCall()
	},
	Cmd_PhoneJoinGroupCall: func() TLObject {
		return NewTLPhoneJoinGroupCall()
	},
	Cmd_PhoneLeaveGroupCall: func() TLObject {
		return NewTLPhoneLeaveGroupCall()
	},
	Cmd_PhoneReceivedCall: func() TLObject {
		return NewTLPhoneReceivedCall()
	},
	Cmd_PhoneRequestCall: func() TLObject {
		return NewTLPhoneRequestCall()
	},
	Cmd_PhoneRequestCall42ff96ed: func() TLObject {
		return NewTLPhoneRequestCall42ff96ed()
	},
	Cmd_PhoneSaveCallDebug: func() TLObject {
		return NewTLPhoneSaveCallDebug()
	},
	Cmd_PhoneSendSignalingData: func() TLObject {
		return NewTLPhoneSendSignalingData()
	},
	Cmd_PhoneSetCallRating: func() TLObject {
		return NewTLPhoneSetCallRating()
	},
	Cmd_PhoneSetCallRating59ead627: func() TLObject {
		return NewTLPhoneSetCallRating59ead627()
	},
	Cmd_PhoneToggleGroupCallSettings: func() TLObject {
		return NewTLPhoneToggleGroupCallSettings()
	},
	Cmd_PhotosDeletePhotos: func() TLObject {
		return NewTLPhotosDeletePhotos()
	},
	Cmd_PhotosGetUserPhotos: func() TLObject {
		return NewTLPhotosGetUserPhotos()
	},
	Cmd_PhotosUpdateProfilePhoto: func() TLObject {
		return NewTLPhotosUpdateProfilePhoto()
	},
	Cmd_PhotosUpdateProfilePhoto72d4742c: func() TLObject {
		return NewTLPhotosUpdateProfilePhoto72d4742c()
	},
	Cmd_PhotosUpdateProfilePhotoeef579a0: func() TLObject {
		return NewTLPhotosUpdateProfilePhotoeef579a0()
	},
	Cmd_PhotosUploadProfilePhoto: func() TLObject {
		return NewTLPhotosUploadProfilePhoto()
	},
	Cmd_PhotosUploadProfilePhoto89f30f69: func() TLObject {
		return NewTLPhotosUploadProfilePhoto89f30f69()
	},
	Cmd_PhotosUploadProfilePhotod50f9c88: func() TLObject {
		return NewTLPhotosUploadProfilePhotod50f9c88()
	},
	Cmd_StatsGetBroadcastStats: func() TLObject {
		return NewTLStatsGetBroadcastStats()
	},
	Cmd_StatsGetBroadcastStatse6300dba: func() TLObject {
		return NewTLStatsGetBroadcastStatse6300dba()
	},
	Cmd_StatsGetMegagroupStats: func() TLObject {
		return NewTLStatsGetMegagroupStats()
	},
	Cmd_StatsGetMessagePublicForwards: func() TLObject {
		return NewTLStatsGetMessagePublicForwards()
	},
	Cmd_StatsGetMessageStats: func() TLObject {
		return NewTLStatsGetMessageStats()
	},
	Cmd_StatsLoadAsyncGraph: func() TLObject {
		return NewTLStatsLoadAsyncGraph()
	},
	Cmd_StickersAddStickerToSet: func() TLObject {
		return NewTLStickersAddStickerToSet()
	},
	Cmd_StickersChangeStickerPosition: func() TLObject {
		return NewTLStickersChangeStickerPosition()
	},
	Cmd_StickersCreateStickerSet: func() TLObject {
		return NewTLStickersCreateStickerSet()
	},
	Cmd_StickersCreateStickerSetf1036780: func() TLObject {
		return NewTLStickersCreateStickerSetf1036780()
	},
	Cmd_StickersRemoveStickerFromSet: func() TLObject {
		return NewTLStickersRemoveStickerFromSet()
	},
	Cmd_StickersSetStickerSetThumb: func() TLObject {
		return NewTLStickersSetStickerSetThumb()
	},
	Cmd_UpdatesGetChannelDifference: func() TLObject {
		return NewTLUpdatesGetChannelDifference()
	},
	Cmd_UpdatesGetChannelDifferencebb32d7c0: func() TLObject {
		return NewTLUpdatesGetChannelDifferencebb32d7c0()
	},
	Cmd_UpdatesGetDifference: func() TLObject {
		return NewTLUpdatesGetDifference()
	},
	Cmd_UpdatesGetDifferencea041495: func() TLObject {
		return NewTLUpdatesGetDifferencea041495()
	},
	Cmd_UpdatesGetState: func() TLObject {
		return NewTLUpdatesGetState()
	},
	Cmd_UploadGetCdnFile: func() TLObject {
		return NewTLUploadGetCdnFile()
	},
	Cmd_UploadGetCdnFileHashes: func() TLObject {
		return NewTLUploadGetCdnFileHashes()
	},
	Cmd_UploadGetCdnFileHashes4da54231: func() TLObject {
		return NewTLUploadGetCdnFileHashes4da54231()
	},
	Cmd_UploadGetFile: func() TLObject {
		return NewTLUploadGetFile()
	},
	Cmd_UploadGetFileHashes: func() TLObject {
		return NewTLUploadGetFileHashes()
	},
	Cmd_UploadGetFileb15a9afc: func() TLObject {
		return NewTLUploadGetFileb15a9afc()
	},
	Cmd_UploadGetWebFile: func() TLObject {
		return NewTLUploadGetWebFile()
	},
	Cmd_UploadReuploadCdnFile: func() TLObject {
		return NewTLUploadReuploadCdnFile()
	},
	Cmd_UploadReuploadCdnFile9b2754a8: func() TLObject {
		return NewTLUploadReuploadCdnFile9b2754a8()
	},
	Cmd_UploadSaveBigFilePart: func() TLObject {
		return NewTLUploadSaveBigFilePart()
	},
	Cmd_UploadSaveFilePart: func() TLObject {
		return NewTLUploadSaveFilePart()
	},
	Cmd_UsersGetFullUser: func() TLObject {
		return NewTLUsersGetFullUser()
	},
	Cmd_UsersGetUsers: func() TLObject {
		return NewTLUsersGetUsers()
	},
	Cmd_UsersSetSecureValueErrors: func() TLObject {
		return NewTLUsersSetSecureValueErrors()
	},
	Cmd_WalletGetKeySecretSalt: func() TLObject {
		return NewTLWalletGetKeySecretSalt()
	},
	Cmd_WalletSendLiteRequest: func() TLObject {
		return NewTLWalletSendLiteRequest()
	},

	Cmd_AccountAuthorizationForm: func() TLObject {
		return NewTLAccountAuthorizationForm(nil)
	},
	Cmd_AccountAuthorizations: func() TLObject {
		return NewTLAccountAuthorizations(nil)
	},
	Cmd_AccountAutoDownloadSettings: func() TLObject {
		return NewTLAccountAutoDownloadSettings(nil)
	},
	Cmd_AccountContentSettings: func() TLObject {
		return NewTLAccountContentSettings(nil)
	},
	Cmd_AccountDaysTTL: func() TLObject {
		return NewTLAccountDaysTTL(nil)
	},
	Cmd_AccountNoPassword: func() TLObject {
		return NewTLAccountNoPassword(nil)
	},
	Cmd_AccountPassword: func() TLObject {
		return NewTLAccountPassword(nil)
	},
	Cmd_AccountPasswordInputSettings: func() TLObject {
		return NewTLAccountPasswordInputSettings(nil)
	},
	Cmd_AccountPasswordInputSettingsc23727c9: func() TLObject {
		return NewTLAccountPasswordInputSettings(nil)
	},
	Cmd_AccountPasswordSettings: func() TLObject {
		return NewTLAccountPasswordSettings(nil)
	},
	Cmd_AccountPasswordSettings9a5c33e5: func() TLObject {
		return NewTLAccountPasswordSettings(nil)
	},
	Cmd_AccountPasswordad2641f8: func() TLObject {
		return NewTLAccountPassword(nil)
	},
	Cmd_AccountPrivacyRules: func() TLObject {
		return NewTLAccountPrivacyRules(nil)
	},
	Cmd_AccountPrivacyRules50a04e45: func() TLObject {
		return NewTLAccountPrivacyRules(nil)
	},
	Cmd_AccountSentEmailCode: func() TLObject {
		return NewTLAccountSentEmailCode(nil)
	},
	Cmd_AccountTakeout: func() TLObject {
		return NewTLAccountTakeout(nil)
	},
	Cmd_AccountThemes: func() TLObject {
		return NewTLAccountThemes(nil)
	},
	Cmd_AccountThemesNotModified: func() TLObject {
		return NewTLAccountThemesNotModified(nil)
	},
	Cmd_AccountTmpPassword: func() TLObject {
		return NewTLAccountTmpPassword(nil)
	},
	Cmd_AccountWallPapers: func() TLObject {
		return NewTLAccountWallPapers(nil)
	},
	Cmd_AccountWallPapersNotModified: func() TLObject {
		return NewTLAccountWallPapersNotModified(nil)
	},
	Cmd_AccountWebAuthorizations: func() TLObject {
		return NewTLAccountWebAuthorizations(nil)
	},
	Cmd_AuthAuthorization: func() TLObject {
		return NewTLAuthAuthorization(nil)
	},
	Cmd_AuthAuthorizationSignUpRequired: func() TLObject {
		return NewTLAuthAuthorizationSignUpRequired(nil)
	},
	Cmd_AuthAuthorizationff036af1: func() TLObject {
		return NewTLAuthAuthorization(nil)
	},
	Cmd_AuthCheckedPhone: func() TLObject {
		return NewTLAuthCheckedPhone(nil)
	},
	Cmd_AuthCodeTypeCall: func() TLObject {
		return NewTLAuthCodeTypeCall(nil)
	},
	Cmd_AuthCodeTypeFlashCall: func() TLObject {
		return NewTLAuthCodeTypeFlashCall(nil)
	},
	Cmd_AuthCodeTypeSms: func() TLObject {
		return NewTLAuthCodeTypeSms(nil)
	},
	Cmd_AuthExportedAuthorization: func() TLObject {
		return NewTLAuthExportedAuthorization(nil)
	},
	Cmd_AuthLoginToken: func() TLObject {
		return NewTLAuthLoginToken(nil)
	},
	Cmd_AuthLoginTokenMigrateTo: func() TLObject {
		return NewTLAuthLoginTokenMigrateTo(nil)
	},
	Cmd_AuthLoginTokenSuccess: func() TLObject {
		return NewTLAuthLoginTokenSuccess(nil)
	},
	Cmd_AuthPasswordRecovery: func() TLObject {
		return NewTLAuthPasswordRecovery(nil)
	},
	Cmd_AuthSentCode: func() TLObject {
		return NewTLAuthSentCode(nil)
	},
	Cmd_AuthSentCode38faab5f: func() TLObject {
		return NewTLAuthSentCode(nil)
	},
	Cmd_AuthSentCodeTypeApp: func() TLObject {
		return NewTLAuthSentCodeTypeApp(nil)
	},
	Cmd_AuthSentCodeTypeCall: func() TLObject {
		return NewTLAuthSentCodeTypeCall(nil)
	},
	Cmd_AuthSentCodeTypeFlashCall: func() TLObject {
		return NewTLAuthSentCodeTypeFlashCall(nil)
	},
	Cmd_AuthSentCodeTypeSms: func() TLObject {
		return NewTLAuthSentCodeTypeSms(nil)
	},
	Cmd_Authorization: func() TLObject {
		return NewTLAuthorization(nil)
	},
	Cmd_Authorizationad01d61d: func() TLObject {
		return NewTLAuthorization(nil)
	},
	Cmd_AutoDownloadSettings: func() TLObject {
		return NewTLAutoDownloadSettings(nil)
	},
	Cmd_AutoDownloadSettingse04232f3: func() TLObject {
		return NewTLAutoDownloadSettings(nil)
	},
	Cmd_BankCardOpenUrl: func() TLObject {
		return NewTLBankCardOpenUrl(nil)
	},
	Cmd_BaseThemeArctic: func() TLObject {
		return NewTLBaseThemeArctic(nil)
	},
	Cmd_BaseThemeClassic: func() TLObject {
		return NewTLBaseThemeClassic(nil)
	},
	Cmd_BaseThemeDay: func() TLObject {
		return NewTLBaseThemeDay(nil)
	},
	Cmd_BaseThemeNight: func() TLObject {
		return NewTLBaseThemeNight(nil)
	},
	Cmd_BaseThemeTinted: func() TLObject {
		return NewTLBaseThemeTinted(nil)
	},
	Cmd_BoolFalse: func() TLObject {
		return NewTLBoolFalse(nil)
	},
	Cmd_BoolTrue: func() TLObject {
		return NewTLBoolTrue(nil)
	},
	Cmd_BotCommand: func() TLObject {
		return NewTLBotCommand(nil)
	},
	Cmd_BotInfo: func() TLObject {
		return NewTLBotInfo(nil)
	},
	Cmd_BotInlineMediaResult: func() TLObject {
		return NewTLBotInlineMediaResult(nil)
	},
	Cmd_BotInlineMessageMediaAuto: func() TLObject {
		return NewTLBotInlineMessageMediaAuto(nil)
	},
	Cmd_BotInlineMessageMediaAuto764cf810: func() TLObject {
		return NewTLBotInlineMessageMediaAuto(nil)
	},
	Cmd_BotInlineMessageMediaContact: func() TLObject {
		return NewTLBotInlineMessageMediaContact(nil)
	},
	Cmd_BotInlineMessageMediaContact18d1cdc2: func() TLObject {
		return NewTLBotInlineMessageMediaContact(nil)
	},
	Cmd_BotInlineMessageMediaGeo: func() TLObject {
		return NewTLBotInlineMessageMediaGeo(nil)
	},
	Cmd_BotInlineMessageMediaGeo51846fd: func() TLObject {
		return NewTLBotInlineMessageMediaGeo(nil)
	},
	Cmd_BotInlineMessageMediaGeob722de65: func() TLObject {
		return NewTLBotInlineMessageMediaGeo(nil)
	},
	Cmd_BotInlineMessageMediaVenue: func() TLObject {
		return NewTLBotInlineMessageMediaVenue(nil)
	},
	Cmd_BotInlineMessageMediaVenue8a86659c: func() TLObject {
		return NewTLBotInlineMessageMediaVenue(nil)
	},
	Cmd_BotInlineMessageText: func() TLObject {
		return NewTLBotInlineMessageText(nil)
	},
	Cmd_BotInlineResult: func() TLObject {
		return NewTLBotInlineResult(nil)
	},
	Cmd_BotInlineResult11965f3a: func() TLObject {
		return NewTLBotInlineResult(nil)
	},
	Cmd_CdnConfig: func() TLObject {
		return NewTLCdnConfig(nil)
	},
	Cmd_CdnFileHash: func() TLObject {
		return NewTLCdnFileHash(nil)
	},
	Cmd_CdnPublicKey: func() TLObject {
		return NewTLCdnPublicKey(nil)
	},
	Cmd_Channel: func() TLObject {
		return NewTLChannel(nil)
	},
	Cmd_Channel4df30834: func() TLObject {
		return NewTLChannel(nil)
	},
	Cmd_ChannelAdminLogEvent: func() TLObject {
		return NewTLChannelAdminLogEvent(nil)
	},
	Cmd_ChannelAdminLogEventActionChangeAbout: func() TLObject {
		return NewTLChannelAdminLogEventActionChangeAbout(nil)
	},
	Cmd_ChannelAdminLogEventActionChangeLinkedChat: func() TLObject {
		return NewTLChannelAdminLogEventActionChangeLinkedChat(nil)
	},
	Cmd_ChannelAdminLogEventActionChangeLocation: func() TLObject {
		return NewTLChannelAdminLogEventActionChangeLocation(nil)
	},
	Cmd_ChannelAdminLogEventActionChangePhoto: func() TLObject {
		return NewTLChannelAdminLogEventActionChangePhoto(nil)
	},
	Cmd_ChannelAdminLogEventActionChangePhoto434bd2af: func() TLObject {
		return NewTLChannelAdminLogEventActionChangePhoto(nil)
	},
	Cmd_ChannelAdminLogEventActionChangeStickerSet: func() TLObject {
		return NewTLChannelAdminLogEventActionChangeStickerSet(nil)
	},
	Cmd_ChannelAdminLogEventActionChangeTitle: func() TLObject {
		return NewTLChannelAdminLogEventActionChangeTitle(nil)
	},
	Cmd_ChannelAdminLogEventActionChangeUsername: func() TLObject {
		return NewTLChannelAdminLogEventActionChangeUsername(nil)
	},
	Cmd_ChannelAdminLogEventActionDefaultBannedRights: func() TLObject {
		return NewTLChannelAdminLogEventActionDefaultBannedRights(nil)
	},
	Cmd_ChannelAdminLogEventActionDeleteMessage: func() TLObject {
		return NewTLChannelAdminLogEventActionDeleteMessage(nil)
	},
	Cmd_ChannelAdminLogEventActionDiscardGroupCall: func() TLObject {
		return NewTLChannelAdminLogEventActionDiscardGroupCall(nil)
	},
	Cmd_ChannelAdminLogEventActionEditMessage: func() TLObject {
		return NewTLChannelAdminLogEventActionEditMessage(nil)
	},
	Cmd_ChannelAdminLogEventActionParticipantInvite: func() TLObject {
		return NewTLChannelAdminLogEventActionParticipantInvite(nil)
	},
	Cmd_ChannelAdminLogEventActionParticipantJoin: func() TLObject {
		return NewTLChannelAdminLogEventActionParticipantJoin(nil)
	},
	Cmd_ChannelAdminLogEventActionParticipantLeave: func() TLObject {
		return NewTLChannelAdminLogEventActionParticipantLeave(nil)
	},
	Cmd_ChannelAdminLogEventActionParticipantMute: func() TLObject {
		return NewTLChannelAdminLogEventActionParticipantMute(nil)
	},
	Cmd_ChannelAdminLogEventActionParticipantToggleAdmin: func() TLObject {
		return NewTLChannelAdminLogEventActionParticipantToggleAdmin(nil)
	},
	Cmd_ChannelAdminLogEventActionParticipantToggleBan: func() TLObject {
		return NewTLChannelAdminLogEventActionParticipantToggleBan(nil)
	},
	Cmd_ChannelAdminLogEventActionParticipantUnmute: func() TLObject {
		return NewTLChannelAdminLogEventActionParticipantUnmute(nil)
	},
	Cmd_ChannelAdminLogEventActionStartGroupCall: func() TLObject {
		return NewTLChannelAdminLogEventActionStartGroupCall(nil)
	},
	Cmd_ChannelAdminLogEventActionStopPoll: func() TLObject {
		return NewTLChannelAdminLogEventActionStopPoll(nil)
	},
	Cmd_ChannelAdminLogEventActionToggleGroupCallSetting: func() TLObject {
		return NewTLChannelAdminLogEventActionToggleGroupCallSetting(nil)
	},
	Cmd_ChannelAdminLogEventActionToggleInvites: func() TLObject {
		return NewTLChannelAdminLogEventActionToggleInvites(nil)
	},
	Cmd_ChannelAdminLogEventActionTogglePreHistoryHidden: func() TLObject {
		return NewTLChannelAdminLogEventActionTogglePreHistoryHidden(nil)
	},
	Cmd_ChannelAdminLogEventActionToggleSignatures: func() TLObject {
		return NewTLChannelAdminLogEventActionToggleSignatures(nil)
	},
	Cmd_ChannelAdminLogEventActionToggleSlowMode: func() TLObject {
		return NewTLChannelAdminLogEventActionToggleSlowMode(nil)
	},
	Cmd_ChannelAdminLogEventActionUpdatePinned: func() TLObject {
		return NewTLChannelAdminLogEventActionUpdatePinned(nil)
	},
	Cmd_ChannelAdminLogEventsFilter: func() TLObject {
		return NewTLChannelAdminLogEventsFilter(nil)
	},
	Cmd_ChannelAdminRights: func() TLObject {
		return NewTLChannelAdminRights(nil)
	},
	Cmd_ChannelBannedRights: func() TLObject {
		return NewTLChannelBannedRights(nil)
	},
	Cmd_ChannelForbidden: func() TLObject {
		return NewTLChannelForbidden(nil)
	},
	Cmd_ChannelForbidden2d85832c: func() TLObject {
		return NewTLChannelForbidden(nil)
	},
	Cmd_ChannelFull: func() TLObject {
		return NewTLChannelFull(nil)
	},
	Cmd_ChannelFull10916653: func() TLObject {
		return NewTLChannelFull(nil)
	},
	Cmd_ChannelFull1c87a71a: func() TLObject {
		return NewTLChannelFull(nil)
	},
	Cmd_ChannelFull2d895c74: func() TLObject {
		return NewTLChannelFull(nil)
	},
	Cmd_ChannelFull3648977: func() TLObject {
		return NewTLChannelFull(nil)
	},
	Cmd_ChannelFull76af5481: func() TLObject {
		return NewTLChannelFull(nil)
	},
	Cmd_ChannelFull97bee562: func() TLObject {
		return NewTLChannelFull(nil)
	},
	Cmd_ChannelFull9882e516: func() TLObject {
		return NewTLChannelFull(nil)
	},
	Cmd_ChannelFullef3a6acd: func() TLObject {
		return NewTLChannelFull(nil)
	},
	Cmd_ChannelFullf0e6672a: func() TLObject {
		return NewTLChannelFull(nil)
	},
	Cmd_ChannelLocation: func() TLObject {
		return NewTLChannelLocation(nil)
	},
	Cmd_ChannelLocationEmpty: func() TLObject {
		return NewTLChannelLocationEmpty(nil)
	},
	Cmd_ChannelMessagesFilter: func() TLObject {
		return NewTLChannelMessagesFilter(nil)
	},
	Cmd_ChannelMessagesFilterCollapsed: func() TLObject {
		return NewTLChannelMessagesFilterCollapsed(nil)
	},
	Cmd_ChannelMessagesFilterEmpty: func() TLObject {
		return NewTLChannelMessagesFilterEmpty(nil)
	},
	Cmd_ChannelParticipant: func() TLObject {
		return NewTLChannelParticipant(nil)
	},
	Cmd_ChannelParticipantAdmin: func() TLObject {
		return NewTLChannelParticipantAdmin(nil)
	},
	Cmd_ChannelParticipantAdmin5daa6e23: func() TLObject {
		return NewTLChannelParticipantAdmin(nil)
	},
	Cmd_ChannelParticipantAdminccbebbaf: func() TLObject {
		return NewTLChannelParticipantAdmin(nil)
	},
	Cmd_ChannelParticipantBanned: func() TLObject {
		return NewTLChannelParticipantBanned(nil)
	},
	Cmd_ChannelParticipantBanned1c0facaf: func() TLObject {
		return NewTLChannelParticipantBanned(nil)
	},
	Cmd_ChannelParticipantCreator: func() TLObject {
		return NewTLChannelParticipantCreator(nil)
	},
	Cmd_ChannelParticipantCreator447dca4b: func() TLObject {
		return NewTLChannelParticipantCreator(nil)
	},
	Cmd_ChannelParticipantCreator808d15a4: func() TLObject {
		return NewTLChannelParticipantCreator(nil)
	},
	Cmd_ChannelParticipantEditor: func() TLObject {
		return NewTLChannelParticipantEditor(nil)
	},
	Cmd_ChannelParticipantKicked: func() TLObject {
		return NewTLChannelParticipantKicked(nil)
	},
	Cmd_ChannelParticipantLeft: func() TLObject {
		return NewTLChannelParticipantLeft(nil)
	},
	Cmd_ChannelParticipantModerator: func() TLObject {
		return NewTLChannelParticipantModerator(nil)
	},
	Cmd_ChannelParticipantSelf: func() TLObject {
		return NewTLChannelParticipantSelf(nil)
	},
	Cmd_ChannelParticipantsAdmins: func() TLObject {
		return NewTLChannelParticipantsAdmins(nil)
	},
	Cmd_ChannelParticipantsBanned: func() TLObject {
		return NewTLChannelParticipantsBanned(nil)
	},
	Cmd_ChannelParticipantsBots: func() TLObject {
		return NewTLChannelParticipantsBots(nil)
	},
	Cmd_ChannelParticipantsContacts: func() TLObject {
		return NewTLChannelParticipantsContacts(nil)
	},
	Cmd_ChannelParticipantsKicked: func() TLObject {
		return NewTLChannelParticipantsKicked(nil)
	},
	Cmd_ChannelParticipantsKicked3c37bb7a: func() TLObject {
		return NewTLChannelParticipantsKicked(nil)
	},
	Cmd_ChannelParticipantsMentions: func() TLObject {
		return NewTLChannelParticipantsMentions(nil)
	},
	Cmd_ChannelParticipantsRecent: func() TLObject {
		return NewTLChannelParticipantsRecent(nil)
	},
	Cmd_ChannelParticipantsSearch: func() TLObject {
		return NewTLChannelParticipantsSearch(nil)
	},
	Cmd_ChannelRoleEditor: func() TLObject {
		return NewTLChannelRoleEditor(nil)
	},
	Cmd_ChannelRoleEmpty: func() TLObject {
		return NewTLChannelRoleEmpty(nil)
	},
	Cmd_ChannelRoleModerator: func() TLObject {
		return NewTLChannelRoleModerator(nil)
	},
	Cmd_Channela14dca52: func() TLObject {
		return NewTLChannel(nil)
	},
	Cmd_Channelc88974ac: func() TLObject {
		return NewTLChannel(nil)
	},
	Cmd_Channeld31a961e: func() TLObject {
		return NewTLChannel(nil)
	},
	Cmd_ChannelsAdminLogResults: func() TLObject {
		return NewTLChannelsAdminLogResults(nil)
	},
	Cmd_ChannelsChannelParticipant: func() TLObject {
		return NewTLChannelsChannelParticipant(nil)
	},
	Cmd_ChannelsChannelParticipants: func() TLObject {
		return NewTLChannelsChannelParticipants(nil)
	},
	Cmd_ChannelsChannelParticipantsNotModified: func() TLObject {
		return NewTLChannelsChannelParticipantsNotModified(nil)
	},
	Cmd_Chat: func() TLObject {
		return NewTLChat(nil)
	},
	Cmd_Chat3bda1bde: func() TLObject {
		return NewTLChat(nil)
	},
	Cmd_ChatAdminRights: func() TLObject {
		return NewTLChatAdminRights(nil)
	},
	Cmd_ChatBannedRights: func() TLObject {
		return NewTLChatBannedRights(nil)
	},
	Cmd_ChatEmpty: func() TLObject {
		return NewTLChatEmpty(nil)
	},
	Cmd_ChatForbidden: func() TLObject {
		return NewTLChatForbidden(nil)
	},
	Cmd_ChatFull: func() TLObject {
		return NewTLChatFull(nil)
	},
	Cmd_ChatFull1b7c9db3: func() TLObject {
		return NewTLChatFull(nil)
	},
	Cmd_ChatFull22a235da: func() TLObject {
		return NewTLChatFull(nil)
	},
	Cmd_ChatFulldc8c181: func() TLObject {
		return NewTLChatFull(nil)
	},
	Cmd_ChatFulledd2a791: func() TLObject {
		return NewTLChatFull(nil)
	},
	Cmd_ChatInvite: func() TLObject {
		return NewTLChatInvite(nil)
	},
	Cmd_ChatInvite93e99b60: func() TLObject {
		return NewTLChatInvite(nil)
	},
	Cmd_ChatInviteAlready: func() TLObject {
		return NewTLChatInviteAlready(nil)
	},
	Cmd_ChatInviteEmpty: func() TLObject {
		return NewTLChatInviteEmpty(nil)
	},
	Cmd_ChatInviteExported: func() TLObject {
		return NewTLChatInviteExported(nil)
	},
	Cmd_ChatInvitePeek: func() TLObject {
		return NewTLChatInvitePeek(nil)
	},
	Cmd_ChatInvitedfc2f58e: func() TLObject {
		return NewTLChatInvite(nil)
	},
	Cmd_ChatOnlines: func() TLObject {
		return NewTLChatOnlines(nil)
	},
	Cmd_ChatParticipant: func() TLObject {
		return NewTLChatParticipant(nil)
	},
	Cmd_ChatParticipantAdmin: func() TLObject {
		return NewTLChatParticipantAdmin(nil)
	},
	Cmd_ChatParticipantCreator: func() TLObject {
		return NewTLChatParticipantCreator(nil)
	},
	Cmd_ChatParticipants: func() TLObject {
		return NewTLChatParticipants(nil)
	},
	Cmd_ChatParticipantsForbidden: func() TLObject {
		return NewTLChatParticipantsForbidden(nil)
	},
	Cmd_ChatPhoto: func() TLObject {
		return NewTLChatPhoto(nil)
	},
	Cmd_ChatPhoto475cdbd5: func() TLObject {
		return NewTLChatPhoto(nil)
	},
	Cmd_ChatPhotoEmpty: func() TLObject {
		return NewTLChatPhotoEmpty(nil)
	},
	Cmd_ChatPhotod20b9f3c: func() TLObject {
		return NewTLChatPhoto(nil)
	},
	Cmd_CodeSettings: func() TLObject {
		return NewTLCodeSettings(nil)
	},
	Cmd_CodeSettingsdebebe83: func() TLObject {
		return NewTLCodeSettings(nil)
	},
	Cmd_Config: func() TLObject {
		return NewTLConfig(nil)
	},
	Cmd_Config317ceef4: func() TLObject {
		return NewTLConfig(nil)
	},
	Cmd_Config3213dbba: func() TLObject {
		return NewTLConfig(nil)
	},
	Cmd_Config330b4067: func() TLObject {
		return NewTLConfig(nil)
	},
	Cmd_Confige6ca25f6: func() TLObject {
		return NewTLConfig(nil)
	},
	Cmd_Contact: func() TLObject {
		return NewTLContact(nil)
	},
	Cmd_ContactBlocked: func() TLObject {
		return NewTLContactBlocked(nil)
	},
	Cmd_ContactLinkContact: func() TLObject {
		return NewTLContactLinkContact(nil)
	},
	Cmd_ContactLinkHasPhone: func() TLObject {
		return NewTLContactLinkHasPhone(nil)
	},
	Cmd_ContactLinkNone: func() TLObject {
		return NewTLContactLinkNone(nil)
	},
	Cmd_ContactLinkUnknown: func() TLObject {
		return NewTLContactLinkUnknown(nil)
	},
	Cmd_ContactStatus: func() TLObject {
		return NewTLContactStatus(nil)
	},
	Cmd_ContactsBlocked: func() TLObject {
		return NewTLContactsBlocked(nil)
	},
	Cmd_ContactsBlockedSlice: func() TLObject {
		return NewTLContactsBlockedSlice(nil)
	},
	Cmd_ContactsBlockedSlicee1664194: func() TLObject {
		return NewTLContactsBlockedSlice(nil)
	},
	Cmd_ContactsBlockedade1591: func() TLObject {
		return NewTLContactsBlocked(nil)
	},
	Cmd_ContactsContacts: func() TLObject {
		return NewTLContactsContacts(nil)
	},
	Cmd_ContactsContacts6f8b8cb2: func() TLObject {
		return NewTLContactsContacts(nil)
	},
	Cmd_ContactsContactsNotModified: func() TLObject {
		return NewTLContactsContactsNotModified(nil)
	},
	Cmd_ContactsFound: func() TLObject {
		return NewTLContactsFound(nil)
	},
	Cmd_ContactsFoundb3134d9d: func() TLObject {
		return NewTLContactsFound(nil)
	},
	Cmd_ContactsImportedContacts: func() TLObject {
		return NewTLContactsImportedContacts(nil)
	},
	Cmd_ContactsImportedContactsad524315: func() TLObject {
		return NewTLContactsImportedContacts(nil)
	},
	Cmd_ContactsLink: func() TLObject {
		return NewTLContactsLink(nil)
	},
	Cmd_ContactsResolvedPeer: func() TLObject {
		return NewTLContactsResolvedPeer(nil)
	},
	Cmd_ContactsTopPeers: func() TLObject {
		return NewTLContactsTopPeers(nil)
	},
	Cmd_ContactsTopPeersDisabled: func() TLObject {
		return NewTLContactsTopPeersDisabled(nil)
	},
	Cmd_ContactsTopPeersNotModified: func() TLObject {
		return NewTLContactsTopPeersNotModified(nil)
	},
	Cmd_DataJSON: func() TLObject {
		return NewTLDataJSON(nil)
	},
	Cmd_DcOption: func() TLObject {
		return NewTLDcOption(nil)
	},
	Cmd_DcOption18b7a10d: func() TLObject {
		return NewTLDcOption(nil)
	},
	Cmd_Dialog: func() TLObject {
		return NewTLDialog(nil)
	},
	Cmd_Dialog2c171f72: func() TLObject {
		return NewTLDialog(nil)
	},
	Cmd_DialogChannel: func() TLObject {
		return NewTLDialogChannel(nil)
	},
	Cmd_DialogFilter: func() TLObject {
		return NewTLDialogFilter(nil)
	},
	Cmd_DialogFilterSuggested: func() TLObject {
		return NewTLDialogFilterSuggested(nil)
	},
	Cmd_DialogFolder: func() TLObject {
		return NewTLDialogFolder(nil)
	},
	Cmd_DialogPeer: func() TLObject {
		return NewTLDialogPeer(nil)
	},
	Cmd_DialogPeerFolder: func() TLObject {
		return NewTLDialogPeerFolder(nil)
	},
	Cmd_Dialogc1dd804a: func() TLObject {
		return NewTLDialog(nil)
	},
	Cmd_DisabledFeature: func() TLObject {
		return NewTLDisabledFeature(nil)
	},
	Cmd_Document: func() TLObject {
		return NewTLDocument(nil)
	},
	Cmd_Document1e87342b: func() TLObject {
		return NewTLDocument(nil)
	},
	Cmd_Document59534e4c: func() TLObject {
		return NewTLDocument(nil)
	},
	Cmd_Document9ba29cc1: func() TLObject {
		return NewTLDocument(nil)
	},
	Cmd_DocumentAttributeAnimated: func() TLObject {
		return NewTLDocumentAttributeAnimated(nil)
	},
	Cmd_DocumentAttributeAudio: func() TLObject {
		return NewTLDocumentAttributeAudio(nil)
	},
	Cmd_DocumentAttributeFilename: func() TLObject {
		return NewTLDocumentAttributeFilename(nil)
	},
	Cmd_DocumentAttributeHasStickers: func() TLObject {
		return NewTLDocumentAttributeHasStickers(nil)
	},
	Cmd_DocumentAttributeImageSize: func() TLObject {
		return NewTLDocumentAttributeImageSize(nil)
	},
	Cmd_DocumentAttributeSticker: func() TLObject {
		return NewTLDocumentAttributeSticker(nil)
	},
	Cmd_DocumentAttributeSticker3a556302: func() TLObject {
		return NewTLDocumentAttributeSticker(nil)
	},
	Cmd_DocumentAttributeVideo: func() TLObject {
		return NewTLDocumentAttributeVideo(nil)
	},
	Cmd_DocumentAttributeVideo5910cccb: func() TLObject {
		return NewTLDocumentAttributeVideo(nil)
	},
	Cmd_DocumentEmpty: func() TLObject {
		return NewTLDocumentEmpty(nil)
	},
	Cmd_Documentf9a39f4f: func() TLObject {
		return NewTLDocument(nil)
	},
	Cmd_DraftMessage: func() TLObject {
		return NewTLDraftMessage(nil)
	},
	Cmd_DraftMessageEmpty: func() TLObject {
		return NewTLDraftMessageEmpty(nil)
	},
	Cmd_DraftMessageEmpty1b0c841a: func() TLObject {
		return NewTLDraftMessageEmpty(nil)
	},
	Cmd_EmojiKeyword: func() TLObject {
		return NewTLEmojiKeyword(nil)
	},
	Cmd_EmojiKeywordDeleted: func() TLObject {
		return NewTLEmojiKeywordDeleted(nil)
	},
	Cmd_EmojiKeywordsDifference: func() TLObject {
		return NewTLEmojiKeywordsDifference(nil)
	},
	Cmd_EmojiLanguage: func() TLObject {
		return NewTLEmojiLanguage(nil)
	},
	Cmd_EmojiURL: func() TLObject {
		return NewTLEmojiURL(nil)
	},
	Cmd_EncryptedChat: func() TLObject {
		return NewTLEncryptedChat(nil)
	},
	Cmd_EncryptedChatDiscarded: func() TLObject {
		return NewTLEncryptedChatDiscarded(nil)
	},
	Cmd_EncryptedChatEmpty: func() TLObject {
		return NewTLEncryptedChatEmpty(nil)
	},
	Cmd_EncryptedChatRequested: func() TLObject {
		return NewTLEncryptedChatRequested(nil)
	},
	Cmd_EncryptedChatRequested62718a82: func() TLObject {
		return NewTLEncryptedChatRequested(nil)
	},
	Cmd_EncryptedChatWaiting: func() TLObject {
		return NewTLEncryptedChatWaiting(nil)
	},
	Cmd_EncryptedFile: func() TLObject {
		return NewTLEncryptedFile(nil)
	},
	Cmd_EncryptedFileEmpty: func() TLObject {
		return NewTLEncryptedFileEmpty(nil)
	},
	Cmd_EncryptedMessage: func() TLObject {
		return NewTLEncryptedMessage(nil)
	},
	Cmd_EncryptedMessageService: func() TLObject {
		return NewTLEncryptedMessageService(nil)
	},
	Cmd_Error: func() TLObject {
		return NewTLError(nil)
	},
	Cmd_ExportedMessageLink: func() TLObject {
		return NewTLExportedMessageLink(nil)
	},
	Cmd_ExportedMessageLink5dab1af4: func() TLObject {
		return NewTLExportedMessageLink(nil)
	},
	Cmd_FileHash: func() TLObject {
		return NewTLFileHash(nil)
	},
	Cmd_FileLocation: func() TLObject {
		return NewTLFileLocation(nil)
	},
	Cmd_FileLocation91d11eb: func() TLObject {
		return NewTLFileLocation(nil)
	},
	Cmd_FileLocationToBeDeprecated: func() TLObject {
		return NewTLFileLocationToBeDeprecated(nil)
	},
	Cmd_FileLocationUnavailable: func() TLObject {
		return NewTLFileLocationUnavailable(nil)
	},
	Cmd_Folder: func() TLObject {
		return NewTLFolder(nil)
	},
	Cmd_FolderPeer: func() TLObject {
		return NewTLFolderPeer(nil)
	},
	Cmd_FoundGif: func() TLObject {
		return NewTLFoundGif(nil)
	},
	Cmd_FoundGifCached: func() TLObject {
		return NewTLFoundGifCached(nil)
	},
	Cmd_Game: func() TLObject {
		return NewTLGame(nil)
	},
	Cmd_GeoPoint: func() TLObject {
		return NewTLGeoPoint(nil)
	},
	Cmd_GeoPoint296f104: func() TLObject {
		return NewTLGeoPoint(nil)
	},
	Cmd_GeoPointEmpty: func() TLObject {
		return NewTLGeoPointEmpty(nil)
	},
	Cmd_GeoPointb2a2f663: func() TLObject {
		return NewTLGeoPoint(nil)
	},
	Cmd_GlobalPrivacySettings: func() TLObject {
		return NewTLGlobalPrivacySettings(nil)
	},
	Cmd_GroupCall: func() TLObject {
		return NewTLGroupCall(nil)
	},
	Cmd_GroupCallDiscarded: func() TLObject {
		return NewTLGroupCallDiscarded(nil)
	},
	Cmd_GroupCallParticipant: func() TLObject {
		return NewTLGroupCallParticipant(nil)
	},
	Cmd_HelpAppChangelog: func() TLObject {
		return NewTLHelpAppChangelog(nil)
	},
	Cmd_HelpAppChangelogEmpty: func() TLObject {
		return NewTLHelpAppChangelogEmpty(nil)
	},
	Cmd_HelpAppUpdate: func() TLObject {
		return NewTLHelpAppUpdate(nil)
	},
	Cmd_HelpAppUpdate1da7158f: func() TLObject {
		return NewTLHelpAppUpdate(nil)
	},
	Cmd_HelpCountriesList: func() TLObject {
		return NewTLHelpCountriesList(nil)
	},
	Cmd_HelpCountriesListNotModified: func() TLObject {
		return NewTLHelpCountriesListNotModified(nil)
	},
	Cmd_HelpCountry: func() TLObject {
		return NewTLHelpCountry(nil)
	},
	Cmd_HelpCountryCode: func() TLObject {
		return NewTLHelpCountryCode(nil)
	},
	Cmd_HelpDeepLinkInfo: func() TLObject {
		return NewTLHelpDeepLinkInfo(nil)
	},
	Cmd_HelpDeepLinkInfoEmpty: func() TLObject {
		return NewTLHelpDeepLinkInfoEmpty(nil)
	},
	Cmd_HelpInviteText: func() TLObject {
		return NewTLHelpInviteText(nil)
	},
	Cmd_HelpNoAppUpdate: func() TLObject {
		return NewTLHelpNoAppUpdate(nil)
	},
	Cmd_HelpPassportConfig: func() TLObject {
		return NewTLHelpPassportConfig(nil)
	},
	Cmd_HelpPassportConfigNotModified: func() TLObject {
		return NewTLHelpPassportConfigNotModified(nil)
	},
	Cmd_HelpPromoData: func() TLObject {
		return NewTLHelpPromoData(nil)
	},
	Cmd_HelpPromoData8c39793f: func() TLObject {
		return NewTLHelpPromoData(nil)
	},
	Cmd_HelpPromoDataEmpty: func() TLObject {
		return NewTLHelpPromoDataEmpty(nil)
	},
	Cmd_HelpProxyDataEmpty: func() TLObject {
		return NewTLHelpProxyDataEmpty(nil)
	},
	Cmd_HelpProxyDataPromo: func() TLObject {
		return NewTLHelpProxyDataPromo(nil)
	},
	Cmd_HelpRecentMeUrls: func() TLObject {
		return NewTLHelpRecentMeUrls(nil)
	},
	Cmd_HelpSupport: func() TLObject {
		return NewTLHelpSupport(nil)
	},
	Cmd_HelpSupportName: func() TLObject {
		return NewTLHelpSupportName(nil)
	},
	Cmd_HelpTermsOfService: func() TLObject {
		return NewTLHelpTermsOfService(nil)
	},
	Cmd_HelpTermsOfService780a0310: func() TLObject {
		return NewTLHelpTermsOfService(nil)
	},
	Cmd_HelpTermsOfServiceUpdate: func() TLObject {
		return NewTLHelpTermsOfServiceUpdate(nil)
	},
	Cmd_HelpTermsOfServiceUpdateEmpty: func() TLObject {
		return NewTLHelpTermsOfServiceUpdateEmpty(nil)
	},
	Cmd_HelpUserInfo: func() TLObject {
		return NewTLHelpUserInfo(nil)
	},
	Cmd_HelpUserInfoEmpty: func() TLObject {
		return NewTLHelpUserInfoEmpty(nil)
	},
	Cmd_HighScore: func() TLObject {
		return NewTLHighScore(nil)
	},
	Cmd_ImportedContact: func() TLObject {
		return NewTLImportedContact(nil)
	},
	Cmd_InlineBotSwitchPM: func() TLObject {
		return NewTLInlineBotSwitchPM(nil)
	},
	Cmd_InlineQueryPeerTypeBroadcast: func() TLObject {
		return NewTLInlineQueryPeerTypeBroadcast(nil)
	},
	Cmd_InlineQueryPeerTypeChat: func() TLObject {
		return NewTLInlineQueryPeerTypeChat(nil)
	},
	Cmd_InlineQueryPeerTypeMegagroup: func() TLObject {
		return NewTLInlineQueryPeerTypeMegagroup(nil)
	},
	Cmd_InlineQueryPeerTypePM: func() TLObject {
		return NewTLInlineQueryPeerTypePM(nil)
	},
	Cmd_InlineQueryPeerTypeSameBotPM: func() TLObject {
		return NewTLInlineQueryPeerTypeSameBotPM(nil)
	},
	Cmd_InputAppEvent: func() TLObject {
		return NewTLInputAppEvent(nil)
	},
	Cmd_InputAppEvent1d1b1245: func() TLObject {
		return NewTLInputAppEvent(nil)
	},
	Cmd_InputBotInlineMessageGame: func() TLObject {
		return NewTLInputBotInlineMessageGame(nil)
	},
	Cmd_InputBotInlineMessageID: func() TLObject {
		return NewTLInputBotInlineMessageID(nil)
	},
	Cmd_InputBotInlineMessageMediaAuto: func() TLObject {
		return NewTLInputBotInlineMessageMediaAuto(nil)
	},
	Cmd_InputBotInlineMessageMediaAuto3380c786: func() TLObject {
		return NewTLInputBotInlineMessageMediaAuto(nil)
	},
	Cmd_InputBotInlineMessageMediaContact: func() TLObject {
		return NewTLInputBotInlineMessageMediaContact(nil)
	},
	Cmd_InputBotInlineMessageMediaContacta6edbffd: func() TLObject {
		return NewTLInputBotInlineMessageMediaContact(nil)
	},
	Cmd_InputBotInlineMessageMediaGeo: func() TLObject {
		return NewTLInputBotInlineMessageMediaGeo(nil)
	},
	Cmd_InputBotInlineMessageMediaGeo96929a85: func() TLObject {
		return NewTLInputBotInlineMessageMediaGeo(nil)
	},
	Cmd_InputBotInlineMessageMediaGeoc1b15d65: func() TLObject {
		return NewTLInputBotInlineMessageMediaGeo(nil)
	},
	Cmd_InputBotInlineMessageMediaVenue: func() TLObject {
		return NewTLInputBotInlineMessageMediaVenue(nil)
	},
	Cmd_InputBotInlineMessageMediaVenue417bbf11: func() TLObject {
		return NewTLInputBotInlineMessageMediaVenue(nil)
	},
	Cmd_InputBotInlineMessageText: func() TLObject {
		return NewTLInputBotInlineMessageText(nil)
	},
	Cmd_InputBotInlineResult: func() TLObject {
		return NewTLInputBotInlineResult(nil)
	},
	Cmd_InputBotInlineResult88bf9319: func() TLObject {
		return NewTLInputBotInlineResult(nil)
	},
	Cmd_InputBotInlineResultDocument: func() TLObject {
		return NewTLInputBotInlineResultDocument(nil)
	},
	Cmd_InputBotInlineResultGame: func() TLObject {
		return NewTLInputBotInlineResultGame(nil)
	},
	Cmd_InputBotInlineResultPhoto: func() TLObject {
		return NewTLInputBotInlineResultPhoto(nil)
	},
	Cmd_InputChannel: func() TLObject {
		return NewTLInputChannel(nil)
	},
	Cmd_InputChannelEmpty: func() TLObject {
		return NewTLInputChannelEmpty(nil)
	},
	Cmd_InputChannelFromMessage: func() TLObject {
		return NewTLInputChannelFromMessage(nil)
	},
	Cmd_InputChatPhoto: func() TLObject {
		return NewTLInputChatPhoto(nil)
	},
	Cmd_InputChatPhotoEmpty: func() TLObject {
		return NewTLInputChatPhotoEmpty(nil)
	},
	Cmd_InputChatPhotob2e1bf08: func() TLObject {
		return NewTLInputChatPhoto(nil)
	},
	Cmd_InputChatUploadedPhoto: func() TLObject {
		return NewTLInputChatUploadedPhoto(nil)
	},
	Cmd_InputChatUploadedPhoto94254732: func() TLObject {
		return NewTLInputChatUploadedPhoto(nil)
	},
	Cmd_InputChatUploadedPhotoc642724e: func() TLObject {
		return NewTLInputChatUploadedPhoto(nil)
	},
	Cmd_InputCheckPasswordEmpty: func() TLObject {
		return NewTLInputCheckPasswordEmpty(nil)
	},
	Cmd_InputCheckPasswordSRP: func() TLObject {
		return NewTLInputCheckPasswordSRP(nil)
	},
	Cmd_InputClientProxy: func() TLObject {
		return NewTLInputClientProxy(nil)
	},
	Cmd_InputDialogPeer: func() TLObject {
		return NewTLInputDialogPeer(nil)
	},
	Cmd_InputDialogPeerFolder: func() TLObject {
		return NewTLInputDialogPeerFolder(nil)
	},
	Cmd_InputDocument: func() TLObject {
		return NewTLInputDocument(nil)
	},
	Cmd_InputDocument1abfb575: func() TLObject {
		return NewTLInputDocument(nil)
	},
	Cmd_InputDocumentEmpty: func() TLObject {
		return NewTLInputDocumentEmpty(nil)
	},
	Cmd_InputDocumentFileLocation: func() TLObject {
		return NewTLInputDocumentFileLocation(nil)
	},
	Cmd_InputDocumentFileLocation196683d9: func() TLObject {
		return NewTLInputDocumentFileLocation(nil)
	},
	Cmd_InputDocumentFileLocation4e45abe9: func() TLObject {
		return NewTLInputDocumentFileLocation(nil)
	},
	Cmd_InputDocumentFileLocationbad07584: func() TLObject {
		return NewTLInputDocumentFileLocation(nil)
	},
	Cmd_InputEncryptedChat: func() TLObject {
		return NewTLInputEncryptedChat(nil)
	},
	Cmd_InputEncryptedFile: func() TLObject {
		return NewTLInputEncryptedFile(nil)
	},
	Cmd_InputEncryptedFileBigUploaded: func() TLObject {
		return NewTLInputEncryptedFileBigUploaded(nil)
	},
	Cmd_InputEncryptedFileEmpty: func() TLObject {
		return NewTLInputEncryptedFileEmpty(nil)
	},
	Cmd_InputEncryptedFileLocation: func() TLObject {
		return NewTLInputEncryptedFileLocation(nil)
	},
	Cmd_InputEncryptedFileUploaded: func() TLObject {
		return NewTLInputEncryptedFileUploaded(nil)
	},
	Cmd_InputFile: func() TLObject {
		return NewTLInputFile(nil)
	},
	Cmd_InputFileBig: func() TLObject {
		return NewTLInputFileBig(nil)
	},
	Cmd_InputFileLocation: func() TLObject {
		return NewTLInputFileLocation(nil)
	},
	Cmd_InputFileLocationdfdaabe1: func() TLObject {
		return NewTLInputFileLocation(nil)
	},
	Cmd_InputFolderPeer: func() TLObject {
		return NewTLInputFolderPeer(nil)
	},
	Cmd_InputGameID: func() TLObject {
		return NewTLInputGameID(nil)
	},
	Cmd_InputGameShortName: func() TLObject {
		return NewTLInputGameShortName(nil)
	},
	Cmd_InputGeoPoint: func() TLObject {
		return NewTLInputGeoPoint(nil)
	},
	Cmd_InputGeoPoint48222faf: func() TLObject {
		return NewTLInputGeoPoint(nil)
	},
	Cmd_InputGeoPointEmpty: func() TLObject {
		return NewTLInputGeoPointEmpty(nil)
	},
	Cmd_InputGroupCall: func() TLObject {
		return NewTLInputGroupCall(nil)
	},
	Cmd_InputKeyboardButtonUrlAuth: func() TLObject {
		return NewTLInputKeyboardButtonUrlAuth(nil)
	},
	Cmd_InputMediaContact: func() TLObject {
		return NewTLInputMediaContact(nil)
	},
	Cmd_InputMediaContactf8ab7dfb: func() TLObject {
		return NewTLInputMediaContact(nil)
	},
	Cmd_InputMediaDice: func() TLObject {
		return NewTLInputMediaDice(nil)
	},
	Cmd_InputMediaDicee66fbf7b: func() TLObject {
		return NewTLInputMediaDice(nil)
	},
	Cmd_InputMediaDocument: func() TLObject {
		return NewTLInputMediaDocument(nil)
	},
	Cmd_InputMediaDocument1a77f29c: func() TLObject {
		return NewTLInputMediaDocument(nil)
	},
	Cmd_InputMediaDocument23ab23d2: func() TLObject {
		return NewTLInputMediaDocument(nil)
	},
	Cmd_InputMediaDocument33473058: func() TLObject {
		return NewTLInputMediaDocument(nil)
	},
	Cmd_InputMediaDocumentExternal: func() TLObject {
		return NewTLInputMediaDocumentExternal(nil)
	},
	Cmd_InputMediaDocumentExternalfb52dc99: func() TLObject {
		return NewTLInputMediaDocumentExternal(nil)
	},
	Cmd_InputMediaEmpty: func() TLObject {
		return NewTLInputMediaEmpty(nil)
	},
	Cmd_InputMediaGame: func() TLObject {
		return NewTLInputMediaGame(nil)
	},
	Cmd_InputMediaGeoLive: func() TLObject {
		return NewTLInputMediaGeoLive(nil)
	},
	Cmd_InputMediaGeoLive971fa843: func() TLObject {
		return NewTLInputMediaGeoLive(nil)
	},
	Cmd_InputMediaGeoLivece4e82fd: func() TLObject {
		return NewTLInputMediaGeoLive(nil)
	},
	Cmd_InputMediaGeoPoint: func() TLObject {
		return NewTLInputMediaGeoPoint(nil)
	},
	Cmd_InputMediaGifExternal: func() TLObject {
		return NewTLInputMediaGifExternal(nil)
	},
	Cmd_InputMediaInvoice: func() TLObject {
		return NewTLInputMediaInvoice(nil)
	},
	Cmd_InputMediaInvoicef4e096c3: func() TLObject {
		return NewTLInputMediaInvoice(nil)
	},
	Cmd_InputMediaPhoto: func() TLObject {
		return NewTLInputMediaPhoto(nil)
	},
	Cmd_InputMediaPhotoExternal: func() TLObject {
		return NewTLInputMediaPhotoExternal(nil)
	},
	Cmd_InputMediaPhotoExternale5bbfe1a: func() TLObject {
		return NewTLInputMediaPhotoExternal(nil)
	},
	Cmd_InputMediaPhotob3ba0635: func() TLObject {
		return NewTLInputMediaPhoto(nil)
	},
	Cmd_InputMediaPhotoe9bfb4f3: func() TLObject {
		return NewTLInputMediaPhoto(nil)
	},
	Cmd_InputMediaPoll: func() TLObject {
		return NewTLInputMediaPoll(nil)
	},
	Cmd_InputMediaPollabe9ca25: func() TLObject {
		return NewTLInputMediaPoll(nil)
	},
	Cmd_InputMediaPollf94e5f1: func() TLObject {
		return NewTLInputMediaPoll(nil)
	},
	Cmd_InputMediaUploadedDocument: func() TLObject {
		return NewTLInputMediaUploadedDocument(nil)
	},
	Cmd_InputMediaUploadedDocument1d89306d: func() TLObject {
		return NewTLInputMediaUploadedDocument(nil)
	},
	Cmd_InputMediaUploadedDocument5b38c6c1: func() TLObject {
		return NewTLInputMediaUploadedDocument(nil)
	},
	Cmd_InputMediaUploadedPhoto: func() TLObject {
		return NewTLInputMediaUploadedPhoto(nil)
	},
	Cmd_InputMediaUploadedPhoto1e287d04: func() TLObject {
		return NewTLInputMediaUploadedPhoto(nil)
	},
	Cmd_InputMediaUploadedPhotof7aff1c0: func() TLObject {
		return NewTLInputMediaUploadedPhoto(nil)
	},
	Cmd_InputMediaUploadedThumbDocument: func() TLObject {
		return NewTLInputMediaUploadedThumbDocument(nil)
	},
	Cmd_InputMediaVenue: func() TLObject {
		return NewTLInputMediaVenue(nil)
	},
	Cmd_InputMediaVenuec13d1c11: func() TLObject {
		return NewTLInputMediaVenue(nil)
	},
	Cmd_InputMessageCallbackQuery: func() TLObject {
		return NewTLInputMessageCallbackQuery(nil)
	},
	Cmd_InputMessageEntityMentionName: func() TLObject {
		return NewTLInputMessageEntityMentionName(nil)
	},
	Cmd_InputMessageID: func() TLObject {
		return NewTLInputMessageID(nil)
	},
	Cmd_InputMessagePinned: func() TLObject {
		return NewTLInputMessagePinned(nil)
	},
	Cmd_InputMessageReplyTo: func() TLObject {
		return NewTLInputMessageReplyTo(nil)
	},
	Cmd_InputMessagesFilterChatPhotos: func() TLObject {
		return NewTLInputMessagesFilterChatPhotos(nil)
	},
	Cmd_InputMessagesFilterContacts: func() TLObject {
		return NewTLInputMessagesFilterContacts(nil)
	},
	Cmd_InputMessagesFilterDocument: func() TLObject {
		return NewTLInputMessagesFilterDocument(nil)
	},
	Cmd_InputMessagesFilterEmpty: func() TLObject {
		return NewTLInputMessagesFilterEmpty(nil)
	},
	Cmd_InputMessagesFilterGeo: func() TLObject {
		return NewTLInputMessagesFilterGeo(nil)
	},
	Cmd_InputMessagesFilterGif: func() TLObject {
		return NewTLInputMessagesFilterGif(nil)
	},
	Cmd_InputMessagesFilterMusic: func() TLObject {
		return NewTLInputMessagesFilterMusic(nil)
	},
	Cmd_InputMessagesFilterMyMentions: func() TLObject {
		return NewTLInputMessagesFilterMyMentions(nil)
	},
	Cmd_InputMessagesFilterPhoneCalls: func() TLObject {
		return NewTLInputMessagesFilterPhoneCalls(nil)
	},
	Cmd_InputMessagesFilterPhotoVideo: func() TLObject {
		return NewTLInputMessagesFilterPhotoVideo(nil)
	},
	Cmd_InputMessagesFilterPhotoVideoDocuments: func() TLObject {
		return NewTLInputMessagesFilterPhotoVideoDocuments(nil)
	},
	Cmd_InputMessagesFilterPhotos: func() TLObject {
		return NewTLInputMessagesFilterPhotos(nil)
	},
	Cmd_InputMessagesFilterPinned: func() TLObject {
		return NewTLInputMessagesFilterPinned(nil)
	},
	Cmd_InputMessagesFilterRoundVideo: func() TLObject {
		return NewTLInputMessagesFilterRoundVideo(nil)
	},
	Cmd_InputMessagesFilterRoundVoice: func() TLObject {
		return NewTLInputMessagesFilterRoundVoice(nil)
	},
	Cmd_InputMessagesFilterUrl: func() TLObject {
		return NewTLInputMessagesFilterUrl(nil)
	},
	Cmd_InputMessagesFilterVideo: func() TLObject {
		return NewTLInputMessagesFilterVideo(nil)
	},
	Cmd_InputMessagesFilterVoice: func() TLObject {
		return NewTLInputMessagesFilterVoice(nil)
	},
	Cmd_InputNotifyAll: func() TLObject {
		return NewTLInputNotifyAll(nil)
	},
	Cmd_InputNotifyBroadcasts: func() TLObject {
		return NewTLInputNotifyBroadcasts(nil)
	},
	Cmd_InputNotifyChats: func() TLObject {
		return NewTLInputNotifyChats(nil)
	},
	Cmd_InputNotifyPeer: func() TLObject {
		return NewTLInputNotifyPeer(nil)
	},
	Cmd_InputNotifyUsers: func() TLObject {
		return NewTLInputNotifyUsers(nil)
	},
	Cmd_InputPaymentCredentials: func() TLObject {
		return NewTLInputPaymentCredentials(nil)
	},
	Cmd_InputPaymentCredentialsAndroidPay: func() TLObject {
		return NewTLInputPaymentCredentialsAndroidPay(nil)
	},
	Cmd_InputPaymentCredentialsApplePay: func() TLObject {
		return NewTLInputPaymentCredentialsApplePay(nil)
	},
	Cmd_InputPaymentCredentialsSaved: func() TLObject {
		return NewTLInputPaymentCredentialsSaved(nil)
	},
	Cmd_InputPeerChannel: func() TLObject {
		return NewTLInputPeerChannel(nil)
	},
	Cmd_InputPeerChannelFromMessage: func() TLObject {
		return NewTLInputPeerChannelFromMessage(nil)
	},
	Cmd_InputPeerChat: func() TLObject {
		return NewTLInputPeerChat(nil)
	},
	Cmd_InputPeerEmpty: func() TLObject {
		return NewTLInputPeerEmpty(nil)
	},
	Cmd_InputPeerNotifyEventsAll: func() TLObject {
		return NewTLInputPeerNotifyEventsAll(nil)
	},
	Cmd_InputPeerNotifyEventsEmpty: func() TLObject {
		return NewTLInputPeerNotifyEventsEmpty(nil)
	},
	Cmd_InputPeerNotifySettings: func() TLObject {
		return NewTLInputPeerNotifySettings(nil)
	},
	Cmd_InputPeerNotifySettings9c3d198e: func() TLObject {
		return NewTLInputPeerNotifySettings(nil)
	},
	Cmd_InputPeerPhotoFileLocation: func() TLObject {
		return NewTLInputPeerPhotoFileLocation(nil)
	},
	Cmd_InputPeerSelf: func() TLObject {
		return NewTLInputPeerSelf(nil)
	},
	Cmd_InputPeerUser: func() TLObject {
		return NewTLInputPeerUser(nil)
	},
	Cmd_InputPeerUserFromMessage: func() TLObject {
		return NewTLInputPeerUserFromMessage(nil)
	},
	Cmd_InputPhoneCall: func() TLObject {
		return NewTLInputPhoneCall(nil)
	},
	Cmd_InputPhoneContact: func() TLObject {
		return NewTLInputPhoneContact(nil)
	},
	Cmd_InputPhoto: func() TLObject {
		return NewTLInputPhoto(nil)
	},
	Cmd_InputPhoto3bb3b94a: func() TLObject {
		return NewTLInputPhoto(nil)
	},
	Cmd_InputPhotoCrop: func() TLObject {
		return NewTLInputPhotoCrop(nil)
	},
	Cmd_InputPhotoCropAuto: func() TLObject {
		return NewTLInputPhotoCropAuto(nil)
	},
	Cmd_InputPhotoEmpty: func() TLObject {
		return NewTLInputPhotoEmpty(nil)
	},
	Cmd_InputPhotoFileLocation: func() TLObject {
		return NewTLInputPhotoFileLocation(nil)
	},
	Cmd_InputPhotoLegacyFileLocation: func() TLObject {
		return NewTLInputPhotoLegacyFileLocation(nil)
	},
	Cmd_InputPrivacyKeyAddedByPhone: func() TLObject {
		return NewTLInputPrivacyKeyAddedByPhone(nil)
	},
	Cmd_InputPrivacyKeyChatInvite: func() TLObject {
		return NewTLInputPrivacyKeyChatInvite(nil)
	},
	Cmd_InputPrivacyKeyForwards: func() TLObject {
		return NewTLInputPrivacyKeyForwards(nil)
	},
	Cmd_InputPrivacyKeyPhoneCall: func() TLObject {
		return NewTLInputPrivacyKeyPhoneCall(nil)
	},
	Cmd_InputPrivacyKeyPhoneNumber: func() TLObject {
		return NewTLInputPrivacyKeyPhoneNumber(nil)
	},
	Cmd_InputPrivacyKeyPhoneP2P: func() TLObject {
		return NewTLInputPrivacyKeyPhoneP2P(nil)
	},
	Cmd_InputPrivacyKeyProfilePhoto: func() TLObject {
		return NewTLInputPrivacyKeyProfilePhoto(nil)
	},
	Cmd_InputPrivacyKeyStatusTimestamp: func() TLObject {
		return NewTLInputPrivacyKeyStatusTimestamp(nil)
	},
	Cmd_InputPrivacyValueAllowAll: func() TLObject {
		return NewTLInputPrivacyValueAllowAll(nil)
	},
	Cmd_InputPrivacyValueAllowChatParticipants: func() TLObject {
		return NewTLInputPrivacyValueAllowChatParticipants(nil)
	},
	Cmd_InputPrivacyValueAllowContacts: func() TLObject {
		return NewTLInputPrivacyValueAllowContacts(nil)
	},
	Cmd_InputPrivacyValueAllowUsers: func() TLObject {
		return NewTLInputPrivacyValueAllowUsers(nil)
	},
	Cmd_InputPrivacyValueDisallowAll: func() TLObject {
		return NewTLInputPrivacyValueDisallowAll(nil)
	},
	Cmd_InputPrivacyValueDisallowChatParticipants: func() TLObject {
		return NewTLInputPrivacyValueDisallowChatParticipants(nil)
	},
	Cmd_InputPrivacyValueDisallowContacts: func() TLObject {
		return NewTLInputPrivacyValueDisallowContacts(nil)
	},
	Cmd_InputPrivacyValueDisallowUsers: func() TLObject {
		return NewTLInputPrivacyValueDisallowUsers(nil)
	},
	Cmd_InputReportReasonChildAbuse: func() TLObject {
		return NewTLInputReportReasonChildAbuse(nil)
	},
	Cmd_InputReportReasonCopyright: func() TLObject {
		return NewTLInputReportReasonCopyright(nil)
	},
	Cmd_InputReportReasonGeoIrrelevant: func() TLObject {
		return NewTLInputReportReasonGeoIrrelevant(nil)
	},
	Cmd_InputReportReasonOther: func() TLObject {
		return NewTLInputReportReasonOther(nil)
	},
	Cmd_InputReportReasonPornography: func() TLObject {
		return NewTLInputReportReasonPornography(nil)
	},
	Cmd_InputReportReasonSpam: func() TLObject {
		return NewTLInputReportReasonSpam(nil)
	},
	Cmd_InputReportReasonViolence: func() TLObject {
		return NewTLInputReportReasonViolence(nil)
	},
	Cmd_InputSecureFile: func() TLObject {
		return NewTLInputSecureFile(nil)
	},
	Cmd_InputSecureFileLocation: func() TLObject {
		return NewTLInputSecureFileLocation(nil)
	},
	Cmd_InputSecureFileUploaded: func() TLObject {
		return NewTLInputSecureFileUploaded(nil)
	},
	Cmd_InputSecureValue: func() TLObject {
		return NewTLInputSecureValue(nil)
	},
	Cmd_InputSingleMedia: func() TLObject {
		return NewTLInputSingleMedia(nil)
	},
	Cmd_InputStickerSetAnimatedEmoji: func() TLObject {
		return NewTLInputStickerSetAnimatedEmoji(nil)
	},
	Cmd_InputStickerSetDice: func() TLObject {
		return NewTLInputStickerSetDice(nil)
	},
	Cmd_InputStickerSetDicee67f520e: func() TLObject {
		return NewTLInputStickerSetDice(nil)
	},
	Cmd_InputStickerSetEmpty: func() TLObject {
		return NewTLInputStickerSetEmpty(nil)
	},
	Cmd_InputStickerSetID: func() TLObject {
		return NewTLInputStickerSetID(nil)
	},
	Cmd_InputStickerSetItem: func() TLObject {
		return NewTLInputStickerSetItem(nil)
	},
	Cmd_InputStickerSetShortName: func() TLObject {
		return NewTLInputStickerSetShortName(nil)
	},
	Cmd_InputStickerSetThumb: func() TLObject {
		return NewTLInputStickerSetThumb(nil)
	},
	Cmd_InputStickeredMediaDocument: func() TLObject {
		return NewTLInputStickeredMediaDocument(nil)
	},
	Cmd_InputStickeredMediaPhoto: func() TLObject {
		return NewTLInputStickeredMediaPhoto(nil)
	},
	Cmd_InputTakeoutFileLocation: func() TLObject {
		return NewTLInputTakeoutFileLocation(nil)
	},
	Cmd_InputTheme: func() TLObject {
		return NewTLInputTheme(nil)
	},
	Cmd_InputThemeSettings: func() TLObject {
		return NewTLInputThemeSettings(nil)
	},
	Cmd_InputThemeSlug: func() TLObject {
		return NewTLInputThemeSlug(nil)
	},
	Cmd_InputUser: func() TLObject {
		return NewTLInputUser(nil)
	},
	Cmd_InputUserEmpty: func() TLObject {
		return NewTLInputUserEmpty(nil)
	},
	Cmd_InputUserFromMessage: func() TLObject {
		return NewTLInputUserFromMessage(nil)
	},
	Cmd_InputUserSelf: func() TLObject {
		return NewTLInputUserSelf(nil)
	},
	Cmd_InputWallPaper: func() TLObject {
		return NewTLInputWallPaper(nil)
	},
	Cmd_InputWallPaperNoFile: func() TLObject {
		return NewTLInputWallPaperNoFile(nil)
	},
	Cmd_InputWallPaperSlug: func() TLObject {
		return NewTLInputWallPaperSlug(nil)
	},
	Cmd_InputWebDocument: func() TLObject {
		return NewTLInputWebDocument(nil)
	},
	Cmd_InputWebFileGeoPointLocation: func() TLObject {
		return NewTLInputWebFileGeoPointLocation(nil)
	},
	Cmd_InputWebFileLocation: func() TLObject {
		return NewTLInputWebFileLocation(nil)
	},
	Cmd_Invoice: func() TLObject {
		return NewTLInvoice(nil)
	},
	Cmd_JsonArray: func() TLObject {
		return NewTLJsonArray(nil)
	},
	Cmd_JsonBool: func() TLObject {
		return NewTLJsonBool(nil)
	},
	Cmd_JsonNull: func() TLObject {
		return NewTLJsonNull(nil)
	},
	Cmd_JsonNumber: func() TLObject {
		return NewTLJsonNumber(nil)
	},
	Cmd_JsonObject: func() TLObject {
		return NewTLJsonObject(nil)
	},
	Cmd_JsonObjectValue: func() TLObject {
		return NewTLJsonObjectValue(nil)
	},
	Cmd_JsonString: func() TLObject {
		return NewTLJsonString(nil)
	},
	Cmd_KeyboardButton: func() TLObject {
		return NewTLKeyboardButton(nil)
	},
	Cmd_KeyboardButtonBuy: func() TLObject {
		return NewTLKeyboardButtonBuy(nil)
	},
	Cmd_KeyboardButtonCallback: func() TLObject {
		return NewTLKeyboardButtonCallback(nil)
	},
	Cmd_KeyboardButtonCallback35bbdb6b: func() TLObject {
		return NewTLKeyboardButtonCallback(nil)
	},
	Cmd_KeyboardButtonGame: func() TLObject {
		return NewTLKeyboardButtonGame(nil)
	},
	Cmd_KeyboardButtonRequestGeoLocation: func() TLObject {
		return NewTLKeyboardButtonRequestGeoLocation(nil)
	},
	Cmd_KeyboardButtonRequestPhone: func() TLObject {
		return NewTLKeyboardButtonRequestPhone(nil)
	},
	Cmd_KeyboardButtonRequestPoll: func() TLObject {
		return NewTLKeyboardButtonRequestPoll(nil)
	},
	Cmd_KeyboardButtonRow: func() TLObject {
		return NewTLKeyboardButtonRow(nil)
	},
	Cmd_KeyboardButtonSwitchInline: func() TLObject {
		return NewTLKeyboardButtonSwitchInline(nil)
	},
	Cmd_KeyboardButtonSwitchInlineea1b7a14: func() TLObject {
		return NewTLKeyboardButtonSwitchInline(nil)
	},
	Cmd_KeyboardButtonUrl: func() TLObject {
		return NewTLKeyboardButtonUrl(nil)
	},
	Cmd_KeyboardButtonUrlAuth: func() TLObject {
		return NewTLKeyboardButtonUrlAuth(nil)
	},
	Cmd_LabeledPrice: func() TLObject {
		return NewTLLabeledPrice(nil)
	},
	Cmd_LangPackDifference: func() TLObject {
		return NewTLLangPackDifference(nil)
	},
	Cmd_LangPackLanguage: func() TLObject {
		return NewTLLangPackLanguage(nil)
	},
	Cmd_LangPackLanguage651b98d: func() TLObject {
		return NewTLLangPackLanguage(nil)
	},
	Cmd_LangPackLanguageeeca5ce3: func() TLObject {
		return NewTLLangPackLanguage(nil)
	},
	Cmd_LangPackString: func() TLObject {
		return NewTLLangPackString(nil)
	},
	Cmd_LangPackStringDeleted: func() TLObject {
		return NewTLLangPackStringDeleted(nil)
	},
	Cmd_LangPackStringPluralized: func() TLObject {
		return NewTLLangPackStringPluralized(nil)
	},
	Cmd_MaskCoords: func() TLObject {
		return NewTLMaskCoords(nil)
	},
	Cmd_Message: func() TLObject {
		return NewTLMessage(nil)
	},
	Cmd_Message44f9b43d: func() TLObject {
		return NewTLMessage(nil)
	},
	Cmd_Message452c0e65: func() TLObject {
		return NewTLMessage(nil)
	},
	Cmd_Message58ae39c9: func() TLObject {
		return NewTLMessage(nil)
	},
	Cmd_MessageActionBotAllowed: func() TLObject {
		return NewTLMessageActionBotAllowed(nil)
	},
	Cmd_MessageActionChannelCreate: func() TLObject {
		return NewTLMessageActionChannelCreate(nil)
	},
	Cmd_MessageActionChannelMigrateFrom: func() TLObject {
		return NewTLMessageActionChannelMigrateFrom(nil)
	},
	Cmd_MessageActionChatAddUser: func() TLObject {
		return NewTLMessageActionChatAddUser(nil)
	},
	Cmd_MessageActionChatCreate: func() TLObject {
		return NewTLMessageActionChatCreate(nil)
	},
	Cmd_MessageActionChatDeletePhoto: func() TLObject {
		return NewTLMessageActionChatDeletePhoto(nil)
	},
	Cmd_MessageActionChatDeleteUser: func() TLObject {
		return NewTLMessageActionChatDeleteUser(nil)
	},
	Cmd_MessageActionChatEditPhoto: func() TLObject {
		return NewTLMessageActionChatEditPhoto(nil)
	},
	Cmd_MessageActionChatEditTitle: func() TLObject {
		return NewTLMessageActionChatEditTitle(nil)
	},
	Cmd_MessageActionChatJoinedByLink: func() TLObject {
		return NewTLMessageActionChatJoinedByLink(nil)
	},
	Cmd_MessageActionChatMigrateTo: func() TLObject {
		return NewTLMessageActionChatMigrateTo(nil)
	},
	Cmd_MessageActionContactSignUp: func() TLObject {
		return NewTLMessageActionContactSignUp(nil)
	},
	Cmd_MessageActionContactSignUpf3f25f76: func() TLObject {
		return NewTLMessageActionContactSignUp(nil)
	},
	Cmd_MessageActionCustomAction: func() TLObject {
		return NewTLMessageActionCustomAction(nil)
	},
	Cmd_MessageActionEmpty: func() TLObject {
		return NewTLMessageActionEmpty(nil)
	},
	Cmd_MessageActionGameScore: func() TLObject {
		return NewTLMessageActionGameScore(nil)
	},
	Cmd_MessageActionGeoProximityReached: func() TLObject {
		return NewTLMessageActionGeoProximityReached(nil)
	},
	Cmd_MessageActionGroupCall: func() TLObject {
		return NewTLMessageActionGroupCall(nil)
	},
	Cmd_MessageActionHistoryClear: func() TLObject {
		return NewTLMessageActionHistoryClear(nil)
	},
	Cmd_MessageActionInviteToGroupCall: func() TLObject {
		return NewTLMessageActionInviteToGroupCall(nil)
	},
	Cmd_MessageActionPaymentSent: func() TLObject {
		return NewTLMessageActionPaymentSent(nil)
	},
	Cmd_MessageActionPaymentSentMe: func() TLObject {
		return NewTLMessageActionPaymentSentMe(nil)
	},
	Cmd_MessageActionPhoneCall: func() TLObject {
		return NewTLMessageActionPhoneCall(nil)
	},
	Cmd_MessageActionPinMessage: func() TLObject {
		return NewTLMessageActionPinMessage(nil)
	},
	Cmd_MessageActionScreenshotTaken: func() TLObject {
		return NewTLMessageActionScreenshotTaken(nil)
	},
	Cmd_MessageActionSecureValuesSent: func() TLObject {
		return NewTLMessageActionSecureValuesSent(nil)
	},
	Cmd_MessageActionSecureValuesSentMe: func() TLObject {
		return NewTLMessageActionSecureValuesSentMe(nil)
	},
	Cmd_MessageEmpty: func() TLObject {
		return NewTLMessageEmpty(nil)
	},
	Cmd_MessageEntityBankCard: func() TLObject {
		return NewTLMessageEntityBankCard(nil)
	},
	Cmd_MessageEntityBlockquote: func() TLObject {
		return NewTLMessageEntityBlockquote(nil)
	},
	Cmd_MessageEntityBold: func() TLObject {
		return NewTLMessageEntityBold(nil)
	},
	Cmd_MessageEntityBotCommand: func() TLObject {
		return NewTLMessageEntityBotCommand(nil)
	},
	Cmd_MessageEntityCashtag: func() TLObject {
		return NewTLMessageEntityCashtag(nil)
	},
	Cmd_MessageEntityCode: func() TLObject {
		return NewTLMessageEntityCode(nil)
	},
	Cmd_MessageEntityEmail: func() TLObject {
		return NewTLMessageEntityEmail(nil)
	},
	Cmd_MessageEntityHashtag: func() TLObject {
		return NewTLMessageEntityHashtag(nil)
	},
	Cmd_MessageEntityItalic: func() TLObject {
		return NewTLMessageEntityItalic(nil)
	},
	Cmd_MessageEntityMention: func() TLObject {
		return NewTLMessageEntityMention(nil)
	},
	Cmd_MessageEntityMentionName: func() TLObject {
		return NewTLMessageEntityMentionName(nil)
	},
	Cmd_MessageEntityPhone: func() TLObject {
		return NewTLMessageEntityPhone(nil)
	},
	Cmd_MessageEntityPre: func() TLObject {
		return NewTLMessageEntityPre(nil)
	},
	Cmd_MessageEntityStrike: func() TLObject {
		return NewTLMessageEntityStrike(nil)
	},
	Cmd_MessageEntityTextUrl: func() TLObject {
		return NewTLMessageEntityTextUrl(nil)
	},
	Cmd_MessageEntityUnderline: func() TLObject {
		return NewTLMessageEntityUnderline(nil)
	},
	Cmd_MessageEntityUnknown: func() TLObject {
		return NewTLMessageEntityUnknown(nil)
	},
	Cmd_MessageEntityUrl: func() TLObject {
		return NewTLMessageEntityUrl(nil)
	},
	Cmd_MessageFwdHeader: func() TLObject {
		return NewTLMessageFwdHeader(nil)
	},
	Cmd_MessageFwdHeader353a686b: func() TLObject {
		return NewTLMessageFwdHeader(nil)
	},
	Cmd_MessageFwdHeader559ebe6d: func() TLObject {
		return NewTLMessageFwdHeader(nil)
	},
	Cmd_MessageFwdHeader5f777dce: func() TLObject {
		return NewTLMessageFwdHeader(nil)
	},
	Cmd_MessageFwdHeaderc786ddcb: func() TLObject {
		return NewTLMessageFwdHeader(nil)
	},
	Cmd_MessageFwdHeaderec338270: func() TLObject {
		return NewTLMessageFwdHeader(nil)
	},
	Cmd_MessageGroup: func() TLObject {
		return NewTLMessageGroup(nil)
	},
	Cmd_MessageInteractionCounters: func() TLObject {
		return NewTLMessageInteractionCounters(nil)
	},
	Cmd_MessageMediaContact: func() TLObject {
		return NewTLMessageMediaContact(nil)
	},
	Cmd_MessageMediaContactcbf24940: func() TLObject {
		return NewTLMessageMediaContact(nil)
	},
	Cmd_MessageMediaDice: func() TLObject {
		return NewTLMessageMediaDice(nil)
	},
	Cmd_MessageMediaDice3f7ee58b: func() TLObject {
		return NewTLMessageMediaDice(nil)
	},
	Cmd_MessageMediaDocument: func() TLObject {
		return NewTLMessageMediaDocument(nil)
	},
	Cmd_MessageMediaDocument9cb070d7: func() TLObject {
		return NewTLMessageMediaDocument(nil)
	},
	Cmd_MessageMediaDocumentf3e02ea8: func() TLObject {
		return NewTLMessageMediaDocument(nil)
	},
	Cmd_MessageMediaEmpty: func() TLObject {
		return NewTLMessageMediaEmpty(nil)
	},
	Cmd_MessageMediaGame: func() TLObject {
		return NewTLMessageMediaGame(nil)
	},
	Cmd_MessageMediaGeo: func() TLObject {
		return NewTLMessageMediaGeo(nil)
	},
	Cmd_MessageMediaGeoLive: func() TLObject {
		return NewTLMessageMediaGeoLive(nil)
	},
	Cmd_MessageMediaGeoLiveb940c666: func() TLObject {
		return NewTLMessageMediaGeoLive(nil)
	},
	Cmd_MessageMediaInvoice: func() TLObject {
		return NewTLMessageMediaInvoice(nil)
	},
	Cmd_MessageMediaPhoto: func() TLObject {
		return NewTLMessageMediaPhoto(nil)
	},
	Cmd_MessageMediaPhoto3d8ce53d: func() TLObject {
		return NewTLMessageMediaPhoto(nil)
	},
	Cmd_MessageMediaPhoto695150d7: func() TLObject {
		return NewTLMessageMediaPhoto(nil)
	},
	Cmd_MessageMediaPoll: func() TLObject {
		return NewTLMessageMediaPoll(nil)
	},
	Cmd_MessageMediaUnsupported: func() TLObject {
		return NewTLMessageMediaUnsupported(nil)
	},
	Cmd_MessageMediaVenue: func() TLObject {
		return NewTLMessageMediaVenue(nil)
	},
	Cmd_MessageMediaVenue2ec0533f: func() TLObject {
		return NewTLMessageMediaVenue(nil)
	},
	Cmd_MessageMediaWebPage: func() TLObject {
		return NewTLMessageMediaWebPage(nil)
	},
	Cmd_MessageRange: func() TLObject {
		return NewTLMessageRange(nil)
	},
	Cmd_MessageReplies: func() TLObject {
		return NewTLMessageReplies(nil)
	},
	Cmd_MessageReplyHeader: func() TLObject {
		return NewTLMessageReplyHeader(nil)
	},
	Cmd_MessageService: func() TLObject {
		return NewTLMessageService(nil)
	},
	Cmd_MessageService286fa604: func() TLObject {
		return NewTLMessageService(nil)
	},
	Cmd_MessageUserVote: func() TLObject {
		return NewTLMessageUserVote(nil)
	},
	Cmd_MessageUserVoteInputOption: func() TLObject {
		return NewTLMessageUserVoteInputOption(nil)
	},
	Cmd_MessageUserVoteMultiple: func() TLObject {
		return NewTLMessageUserVoteMultiple(nil)
	},
	Cmd_MessageUserVotea28e5559: func() TLObject {
		return NewTLMessageUserVote(nil)
	},
	Cmd_MessageViews: func() TLObject {
		return NewTLMessageViews(nil)
	},
	Cmd_Messagec09be45f: func() TLObject {
		return NewTLMessage(nil)
	},
	Cmd_MessagesAffectedHistory: func() TLObject {
		return NewTLMessagesAffectedHistory(nil)
	},
	Cmd_MessagesAffectedMessages: func() TLObject {
		return NewTLMessagesAffectedMessages(nil)
	},
	Cmd_MessagesAllStickers: func() TLObject {
		return NewTLMessagesAllStickers(nil)
	},
	Cmd_MessagesAllStickersNotModified: func() TLObject {
		return NewTLMessagesAllStickersNotModified(nil)
	},
	Cmd_MessagesArchivedStickers: func() TLObject {
		return NewTLMessagesArchivedStickers(nil)
	},
	Cmd_MessagesBotCallbackAnswer: func() TLObject {
		return NewTLMessagesBotCallbackAnswer(nil)
	},
	Cmd_MessagesBotCallbackAnswer1264f1c6: func() TLObject {
		return NewTLMessagesBotCallbackAnswer(nil)
	},
	Cmd_MessagesBotResults: func() TLObject {
		return NewTLMessagesBotResults(nil)
	},
	Cmd_MessagesBotResults256709a6: func() TLObject {
		return NewTLMessagesBotResults(nil)
	},
	Cmd_MessagesBotResults947ca848: func() TLObject {
		return NewTLMessagesBotResults(nil)
	},
	Cmd_MessagesChannelMessages: func() TLObject {
		return NewTLMessagesChannelMessages(nil)
	},
	Cmd_MessagesChannelMessages64479808: func() TLObject {
		return NewTLMessagesChannelMessages(nil)
	},
	Cmd_MessagesChannelMessagesbc0f17bc: func() TLObject {
		return NewTLMessagesChannelMessages(nil)
	},
	Cmd_MessagesChatFull: func() TLObject {
		return NewTLMessagesChatFull(nil)
	},
	Cmd_MessagesChats: func() TLObject {
		return NewTLMessagesChats(nil)
	},
	Cmd_MessagesChatsSlice: func() TLObject {
		return NewTLMessagesChatsSlice(nil)
	},
	Cmd_MessagesDhConfig: func() TLObject {
		return NewTLMessagesDhConfig(nil)
	},
	Cmd_MessagesDhConfigNotModified: func() TLObject {
		return NewTLMessagesDhConfigNotModified(nil)
	},
	Cmd_MessagesDialogs: func() TLObject {
		return NewTLMessagesDialogs(nil)
	},
	Cmd_MessagesDialogsNotModified: func() TLObject {
		return NewTLMessagesDialogsNotModified(nil)
	},
	Cmd_MessagesDialogsSlice: func() TLObject {
		return NewTLMessagesDialogsSlice(nil)
	},
	Cmd_MessagesDiscussionMessage: func() TLObject {
		return NewTLMessagesDiscussionMessage(nil)
	},
	Cmd_MessagesFavedStickers: func() TLObject {
		return NewTLMessagesFavedStickers(nil)
	},
	Cmd_MessagesFavedStickersNotModified: func() TLObject {
		return NewTLMessagesFavedStickersNotModified(nil)
	},
	Cmd_MessagesFeaturedStickers: func() TLObject {
		return NewTLMessagesFeaturedStickers(nil)
	},
	Cmd_MessagesFeaturedStickersNotModified: func() TLObject {
		return NewTLMessagesFeaturedStickersNotModified(nil)
	},
	Cmd_MessagesFeaturedStickersNotModifiedc6dc0c66: func() TLObject {
		return NewTLMessagesFeaturedStickersNotModified(nil)
	},
	Cmd_MessagesFeaturedStickersb6abc341: func() TLObject {
		return NewTLMessagesFeaturedStickers(nil)
	},
	Cmd_MessagesFoundGifs: func() TLObject {
		return NewTLMessagesFoundGifs(nil)
	},
	Cmd_MessagesFoundStickerSets: func() TLObject {
		return NewTLMessagesFoundStickerSets(nil)
	},
	Cmd_MessagesFoundStickerSetsNotModified: func() TLObject {
		return NewTLMessagesFoundStickerSetsNotModified(nil)
	},
	Cmd_MessagesHighScores: func() TLObject {
		return NewTLMessagesHighScores(nil)
	},
	Cmd_MessagesInactiveChats: func() TLObject {
		return NewTLMessagesInactiveChats(nil)
	},
	Cmd_MessagesMessageEditData: func() TLObject {
		return NewTLMessagesMessageEditData(nil)
	},
	Cmd_MessagesMessageViews: func() TLObject {
		return NewTLMessagesMessageViews(nil)
	},
	Cmd_MessagesMessages: func() TLObject {
		return NewTLMessagesMessages(nil)
	},
	Cmd_MessagesMessagesNotModified: func() TLObject {
		return NewTLMessagesMessagesNotModified(nil)
	},
	Cmd_MessagesMessagesSlice: func() TLObject {
		return NewTLMessagesMessagesSlice(nil)
	},
	Cmd_MessagesMessagesSlice3a54685e: func() TLObject {
		return NewTLMessagesMessagesSlice(nil)
	},
	Cmd_MessagesMessagesSlicea6c47aaa: func() TLObject {
		return NewTLMessagesMessagesSlice(nil)
	},
	Cmd_MessagesMessagesSlicec8edce1e: func() TLObject {
		return NewTLMessagesMessagesSlice(nil)
	},
	Cmd_MessagesPeerDialogs: func() TLObject {
		return NewTLMessagesPeerDialogs(nil)
	},
	Cmd_MessagesRecentStickers: func() TLObject {
		return NewTLMessagesRecentStickers(nil)
	},
	Cmd_MessagesRecentStickers22f3afb3: func() TLObject {
		return NewTLMessagesRecentStickers(nil)
	},
	Cmd_MessagesRecentStickersNotModified: func() TLObject {
		return NewTLMessagesRecentStickersNotModified(nil)
	},
	Cmd_MessagesSavedGifs: func() TLObject {
		return NewTLMessagesSavedGifs(nil)
	},
	Cmd_MessagesSavedGifsNotModified: func() TLObject {
		return NewTLMessagesSavedGifsNotModified(nil)
	},
	Cmd_MessagesSearchCounter: func() TLObject {
		return NewTLMessagesSearchCounter(nil)
	},
	Cmd_MessagesSentEncryptedFile: func() TLObject {
		return NewTLMessagesSentEncryptedFile(nil)
	},
	Cmd_MessagesSentEncryptedMessage: func() TLObject {
		return NewTLMessagesSentEncryptedMessage(nil)
	},
	Cmd_MessagesStickerSet: func() TLObject {
		return NewTLMessagesStickerSet(nil)
	},
	Cmd_MessagesStickerSetInstallResultArchive: func() TLObject {
		return NewTLMessagesStickerSetInstallResultArchive(nil)
	},
	Cmd_MessagesStickerSetInstallResultSuccess: func() TLObject {
		return NewTLMessagesStickerSetInstallResultSuccess(nil)
	},
	Cmd_MessagesStickers: func() TLObject {
		return NewTLMessagesStickers(nil)
	},
	Cmd_MessagesStickersNotModified: func() TLObject {
		return NewTLMessagesStickersNotModified(nil)
	},
	Cmd_MessagesStickerse4599bbd: func() TLObject {
		return NewTLMessagesStickers(nil)
	},
	Cmd_MessagesVotesList: func() TLObject {
		return NewTLMessagesVotesList(nil)
	},
	Cmd_NearestDc: func() TLObject {
		return NewTLNearestDc(nil)
	},
	Cmd_NotifyAll: func() TLObject {
		return NewTLNotifyAll(nil)
	},
	Cmd_NotifyBroadcasts: func() TLObject {
		return NewTLNotifyBroadcasts(nil)
	},
	Cmd_NotifyChats: func() TLObject {
		return NewTLNotifyChats(nil)
	},
	Cmd_NotifyPeer: func() TLObject {
		return NewTLNotifyPeer(nil)
	},
	Cmd_NotifyUsers: func() TLObject {
		return NewTLNotifyUsers(nil)
	},
	Cmd_Null: func() TLObject {
		return NewTLNull(nil)
	},
	Cmd_Page: func() TLObject {
		return NewTLPage(nil)
	},
	Cmd_Page98657f0d: func() TLObject {
		return NewTLPage(nil)
	},
	Cmd_PageBlockAnchor: func() TLObject {
		return NewTLPageBlockAnchor(nil)
	},
	Cmd_PageBlockAudio: func() TLObject {
		return NewTLPageBlockAudio(nil)
	},
	Cmd_PageBlockAudio804361ea: func() TLObject {
		return NewTLPageBlockAudio(nil)
	},
	Cmd_PageBlockAuthorDate: func() TLObject {
		return NewTLPageBlockAuthorDate(nil)
	},
	Cmd_PageBlockBlockquote: func() TLObject {
		return NewTLPageBlockBlockquote(nil)
	},
	Cmd_PageBlockChannel: func() TLObject {
		return NewTLPageBlockChannel(nil)
	},
	Cmd_PageBlockCollage: func() TLObject {
		return NewTLPageBlockCollage(nil)
	},
	Cmd_PageBlockCollage65a0fa4d: func() TLObject {
		return NewTLPageBlockCollage(nil)
	},
	Cmd_PageBlockCover: func() TLObject {
		return NewTLPageBlockCover(nil)
	},
	Cmd_PageBlockDetails: func() TLObject {
		return NewTLPageBlockDetails(nil)
	},
	Cmd_PageBlockDivider: func() TLObject {
		return NewTLPageBlockDivider(nil)
	},
	Cmd_PageBlockEmbed: func() TLObject {
		return NewTLPageBlockEmbed(nil)
	},
	Cmd_PageBlockEmbedPost: func() TLObject {
		return NewTLPageBlockEmbedPost(nil)
	},
	Cmd_PageBlockEmbedPostf259a80b: func() TLObject {
		return NewTLPageBlockEmbedPost(nil)
	},
	Cmd_PageBlockEmbeda8718dc5: func() TLObject {
		return NewTLPageBlockEmbed(nil)
	},
	Cmd_PageBlockFooter: func() TLObject {
		return NewTLPageBlockFooter(nil)
	},
	Cmd_PageBlockHeader: func() TLObject {
		return NewTLPageBlockHeader(nil)
	},
	Cmd_PageBlockKicker: func() TLObject {
		return NewTLPageBlockKicker(nil)
	},
	Cmd_PageBlockList: func() TLObject {
		return NewTLPageBlockList(nil)
	},
	Cmd_PageBlockListe4e88011: func() TLObject {
		return NewTLPageBlockList(nil)
	},
	Cmd_PageBlockMap: func() TLObject {
		return NewTLPageBlockMap(nil)
	},
	Cmd_PageBlockOrderedList: func() TLObject {
		return NewTLPageBlockOrderedList(nil)
	},
	Cmd_PageBlockParagraph: func() TLObject {
		return NewTLPageBlockParagraph(nil)
	},
	Cmd_PageBlockPhoto: func() TLObject {
		return NewTLPageBlockPhoto(nil)
	},
	Cmd_PageBlockPhoto1759c560: func() TLObject {
		return NewTLPageBlockPhoto(nil)
	},
	Cmd_PageBlockPreformatted: func() TLObject {
		return NewTLPageBlockPreformatted(nil)
	},
	Cmd_PageBlockPullquote: func() TLObject {
		return NewTLPageBlockPullquote(nil)
	},
	Cmd_PageBlockRelatedArticles: func() TLObject {
		return NewTLPageBlockRelatedArticles(nil)
	},
	Cmd_PageBlockSlideshow: func() TLObject {
		return NewTLPageBlockSlideshow(nil)
	},
	Cmd_PageBlockSlideshow31f9590: func() TLObject {
		return NewTLPageBlockSlideshow(nil)
	},
	Cmd_PageBlockSubheader: func() TLObject {
		return NewTLPageBlockSubheader(nil)
	},
	Cmd_PageBlockSubtitle: func() TLObject {
		return NewTLPageBlockSubtitle(nil)
	},
	Cmd_PageBlockTable: func() TLObject {
		return NewTLPageBlockTable(nil)
	},
	Cmd_PageBlockTitle: func() TLObject {
		return NewTLPageBlockTitle(nil)
	},
	Cmd_PageBlockUnsupported: func() TLObject {
		return NewTLPageBlockUnsupported(nil)
	},
	Cmd_PageBlockVideo: func() TLObject {
		return NewTLPageBlockVideo(nil)
	},
	Cmd_PageBlockVideo7c8fe7b6: func() TLObject {
		return NewTLPageBlockVideo(nil)
	},
	Cmd_PageCaption: func() TLObject {
		return NewTLPageCaption(nil)
	},
	Cmd_PageFull: func() TLObject {
		return NewTLPageFull(nil)
	},
	Cmd_PageListItemBlocks: func() TLObject {
		return NewTLPageListItemBlocks(nil)
	},
	Cmd_PageListItemText: func() TLObject {
		return NewTLPageListItemText(nil)
	},
	Cmd_PageListOrderedItemBlocks: func() TLObject {
		return NewTLPageListOrderedItemBlocks(nil)
	},
	Cmd_PageListOrderedItemText: func() TLObject {
		return NewTLPageListOrderedItemText(nil)
	},
	Cmd_PagePart: func() TLObject {
		return NewTLPagePart(nil)
	},
	Cmd_PageRelatedArticle: func() TLObject {
		return NewTLPageRelatedArticle(nil)
	},
	Cmd_PageRelatedArticleb390dc08: func() TLObject {
		return NewTLPageRelatedArticle(nil)
	},
	Cmd_PageTableCell: func() TLObject {
		return NewTLPageTableCell(nil)
	},
	Cmd_PageTableRow: func() TLObject {
		return NewTLPageTableRow(nil)
	},
	Cmd_Pageae891bec: func() TLObject {
		return NewTLPage(nil)
	},
	Cmd_PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow: func() TLObject {
		return NewTLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow(nil)
	},
	Cmd_PasswordKdfAlgoUnknown: func() TLObject {
		return NewTLPasswordKdfAlgoUnknown(nil)
	},
	Cmd_PaymentCharge: func() TLObject {
		return NewTLPaymentCharge(nil)
	},
	Cmd_PaymentRequestedInfo: func() TLObject {
		return NewTLPaymentRequestedInfo(nil)
	},
	Cmd_PaymentSavedCredentialsCard: func() TLObject {
		return NewTLPaymentSavedCredentialsCard(nil)
	},
	Cmd_PaymentsBankCardData: func() TLObject {
		return NewTLPaymentsBankCardData(nil)
	},
	Cmd_PaymentsPaymentForm: func() TLObject {
		return NewTLPaymentsPaymentForm(nil)
	},
	Cmd_PaymentsPaymentReceipt: func() TLObject {
		return NewTLPaymentsPaymentReceipt(nil)
	},
	Cmd_PaymentsPaymentResult: func() TLObject {
		return NewTLPaymentsPaymentResult(nil)
	},
	Cmd_PaymentsPaymentVerficationNeeded: func() TLObject {
		return NewTLPaymentsPaymentVerficationNeeded(nil)
	},
	Cmd_PaymentsPaymentVerificationNeeded: func() TLObject {
		return NewTLPaymentsPaymentVerificationNeeded(nil)
	},
	Cmd_PaymentsSavedInfo: func() TLObject {
		return NewTLPaymentsSavedInfo(nil)
	},
	Cmd_PaymentsValidatedRequestedInfo: func() TLObject {
		return NewTLPaymentsValidatedRequestedInfo(nil)
	},
	Cmd_PeerBlocked: func() TLObject {
		return NewTLPeerBlocked(nil)
	},
	Cmd_PeerChannel: func() TLObject {
		return NewTLPeerChannel(nil)
	},
	Cmd_PeerChat: func() TLObject {
		return NewTLPeerChat(nil)
	},
	Cmd_PeerLocated: func() TLObject {
		return NewTLPeerLocated(nil)
	},
	Cmd_PeerNotifyEventsAll: func() TLObject {
		return NewTLPeerNotifyEventsAll(nil)
	},
	Cmd_PeerNotifyEventsEmpty: func() TLObject {
		return NewTLPeerNotifyEventsEmpty(nil)
	},
	Cmd_PeerNotifySettings: func() TLObject {
		return NewTLPeerNotifySettings(nil)
	},
	Cmd_PeerNotifySettingsEmpty: func() TLObject {
		return NewTLPeerNotifySettingsEmpty(nil)
	},
	Cmd_PeerNotifySettingsaf509d20: func() TLObject {
		return NewTLPeerNotifySettings(nil)
	},
	Cmd_PeerSelfLocated: func() TLObject {
		return NewTLPeerSelfLocated(nil)
	},
	Cmd_PeerSettings: func() TLObject {
		return NewTLPeerSettings(nil)
	},
	Cmd_PeerSettings733f2961: func() TLObject {
		return NewTLPeerSettings(nil)
	},
	Cmd_PeerUser: func() TLObject {
		return NewTLPeerUser(nil)
	},
	Cmd_PhoneCall: func() TLObject {
		return NewTLPhoneCall(nil)
	},
	Cmd_PhoneCall8742ae7f: func() TLObject {
		return NewTLPhoneCall(nil)
	},
	Cmd_PhoneCallAccepted: func() TLObject {
		return NewTLPhoneCallAccepted(nil)
	},
	Cmd_PhoneCallAccepted997c454a: func() TLObject {
		return NewTLPhoneCallAccepted(nil)
	},
	Cmd_PhoneCallDiscardReasonBusy: func() TLObject {
		return NewTLPhoneCallDiscardReasonBusy(nil)
	},
	Cmd_PhoneCallDiscardReasonDisconnect: func() TLObject {
		return NewTLPhoneCallDiscardReasonDisconnect(nil)
	},
	Cmd_PhoneCallDiscardReasonHangup: func() TLObject {
		return NewTLPhoneCallDiscardReasonHangup(nil)
	},
	Cmd_PhoneCallDiscardReasonMissed: func() TLObject {
		return NewTLPhoneCallDiscardReasonMissed(nil)
	},
	Cmd_PhoneCallDiscarded: func() TLObject {
		return NewTLPhoneCallDiscarded(nil)
	},
	Cmd_PhoneCallEmpty: func() TLObject {
		return NewTLPhoneCallEmpty(nil)
	},
	Cmd_PhoneCallProtocol: func() TLObject {
		return NewTLPhoneCallProtocol(nil)
	},
	Cmd_PhoneCallProtocolfc878fc8: func() TLObject {
		return NewTLPhoneCallProtocol(nil)
	},
	Cmd_PhoneCallRequested: func() TLObject {
		return NewTLPhoneCallRequested(nil)
	},
	Cmd_PhoneCallRequested87eabb53: func() TLObject {
		return NewTLPhoneCallRequested(nil)
	},
	Cmd_PhoneCallWaiting: func() TLObject {
		return NewTLPhoneCallWaiting(nil)
	},
	Cmd_PhoneCalle6f9ddf3: func() TLObject {
		return NewTLPhoneCall(nil)
	},
	Cmd_PhoneConnection: func() TLObject {
		return NewTLPhoneConnection(nil)
	},
	Cmd_PhoneConnectionWebrtc: func() TLObject {
		return NewTLPhoneConnectionWebrtc(nil)
	},
	Cmd_PhoneGroupCall: func() TLObject {
		return NewTLPhoneGroupCall(nil)
	},
	Cmd_PhoneGroupParticipants: func() TLObject {
		return NewTLPhoneGroupParticipants(nil)
	},
	Cmd_PhonePhoneCall: func() TLObject {
		return NewTLPhonePhoneCall(nil)
	},
	Cmd_Photo: func() TLObject {
		return NewTLPhoto(nil)
	},
	Cmd_Photo9c477dd8: func() TLObject {
		return NewTLPhoto(nil)
	},
	Cmd_PhotoCachedSize: func() TLObject {
		return NewTLPhotoCachedSize(nil)
	},
	Cmd_PhotoEmpty: func() TLObject {
		return NewTLPhotoEmpty(nil)
	},
	Cmd_PhotoPathSize: func() TLObject {
		return NewTLPhotoPathSize(nil)
	},
	Cmd_PhotoSize: func() TLObject {
		return NewTLPhotoSize(nil)
	},
	Cmd_PhotoSizeEmpty: func() TLObject {
		return NewTLPhotoSizeEmpty(nil)
	},
	Cmd_PhotoSizeProgressive: func() TLObject {
		return NewTLPhotoSizeProgressive(nil)
	},
	Cmd_PhotoStrippedSize: func() TLObject {
		return NewTLPhotoStrippedSize(nil)
	},
	Cmd_Photocded42fe: func() TLObject {
		return NewTLPhoto(nil)
	},
	Cmd_Photod07504a5: func() TLObject {
		return NewTLPhoto(nil)
	},
	Cmd_Photofb197a65: func() TLObject {
		return NewTLPhoto(nil)
	},
	Cmd_PhotosPhoto: func() TLObject {
		return NewTLPhotosPhoto(nil)
	},
	Cmd_PhotosPhotos: func() TLObject {
		return NewTLPhotosPhotos(nil)
	},
	Cmd_PhotosPhotosSlice: func() TLObject {
		return NewTLPhotosPhotosSlice(nil)
	},
	Cmd_Poll: func() TLObject {
		return NewTLPoll(nil)
	},
	Cmd_Poll86e18161: func() TLObject {
		return NewTLPoll(nil)
	},
	Cmd_PollAnswer: func() TLObject {
		return NewTLPollAnswer(nil)
	},
	Cmd_PollAnswerVoters: func() TLObject {
		return NewTLPollAnswerVoters(nil)
	},
	Cmd_PollResults: func() TLObject {
		return NewTLPollResults(nil)
	},
	Cmd_PollResultsbadcc1a3: func() TLObject {
		return NewTLPollResults(nil)
	},
	Cmd_PollResultsc87024a2: func() TLObject {
		return NewTLPollResults(nil)
	},
	Cmd_PopularContact: func() TLObject {
		return NewTLPopularContact(nil)
	},
	Cmd_PostAddress: func() TLObject {
		return NewTLPostAddress(nil)
	},
	Cmd_PrivacyKeyAddedByPhone: func() TLObject {
		return NewTLPrivacyKeyAddedByPhone(nil)
	},
	Cmd_PrivacyKeyChatInvite: func() TLObject {
		return NewTLPrivacyKeyChatInvite(nil)
	},
	Cmd_PrivacyKeyForwards: func() TLObject {
		return NewTLPrivacyKeyForwards(nil)
	},
	Cmd_PrivacyKeyPhoneCall: func() TLObject {
		return NewTLPrivacyKeyPhoneCall(nil)
	},
	Cmd_PrivacyKeyPhoneNumber: func() TLObject {
		return NewTLPrivacyKeyPhoneNumber(nil)
	},
	Cmd_PrivacyKeyPhoneP2P: func() TLObject {
		return NewTLPrivacyKeyPhoneP2P(nil)
	},
	Cmd_PrivacyKeyProfilePhoto: func() TLObject {
		return NewTLPrivacyKeyProfilePhoto(nil)
	},
	Cmd_PrivacyKeyStatusTimestamp: func() TLObject {
		return NewTLPrivacyKeyStatusTimestamp(nil)
	},
	Cmd_PrivacyValueAllowAll: func() TLObject {
		return NewTLPrivacyValueAllowAll(nil)
	},
	Cmd_PrivacyValueAllowChatParticipants: func() TLObject {
		return NewTLPrivacyValueAllowChatParticipants(nil)
	},
	Cmd_PrivacyValueAllowContacts: func() TLObject {
		return NewTLPrivacyValueAllowContacts(nil)
	},
	Cmd_PrivacyValueAllowUsers: func() TLObject {
		return NewTLPrivacyValueAllowUsers(nil)
	},
	Cmd_PrivacyValueDisallowAll: func() TLObject {
		return NewTLPrivacyValueDisallowAll(nil)
	},
	Cmd_PrivacyValueDisallowChatParticipants: func() TLObject {
		return NewTLPrivacyValueDisallowChatParticipants(nil)
	},
	Cmd_PrivacyValueDisallowContacts: func() TLObject {
		return NewTLPrivacyValueDisallowContacts(nil)
	},
	Cmd_PrivacyValueDisallowUsers: func() TLObject {
		return NewTLPrivacyValueDisallowUsers(nil)
	},
	Cmd_ReceivedNotifyMessage: func() TLObject {
		return NewTLReceivedNotifyMessage(nil)
	},
	Cmd_RecentMeUrlChat: func() TLObject {
		return NewTLRecentMeUrlChat(nil)
	},
	Cmd_RecentMeUrlChatInvite: func() TLObject {
		return NewTLRecentMeUrlChatInvite(nil)
	},
	Cmd_RecentMeUrlStickerSet: func() TLObject {
		return NewTLRecentMeUrlStickerSet(nil)
	},
	Cmd_RecentMeUrlUnknown: func() TLObject {
		return NewTLRecentMeUrlUnknown(nil)
	},
	Cmd_RecentMeUrlUser: func() TLObject {
		return NewTLRecentMeUrlUser(nil)
	},
	Cmd_ReplyInlineMarkup: func() TLObject {
		return NewTLReplyInlineMarkup(nil)
	},
	Cmd_ReplyKeyboardForceReply: func() TLObject {
		return NewTLReplyKeyboardForceReply(nil)
	},
	Cmd_ReplyKeyboardHide: func() TLObject {
		return NewTLReplyKeyboardHide(nil)
	},
	Cmd_ReplyKeyboardMarkup: func() TLObject {
		return NewTLReplyKeyboardMarkup(nil)
	},
	Cmd_RestrictionReason: func() TLObject {
		return NewTLRestrictionReason(nil)
	},
	Cmd_SavedPhoneContact: func() TLObject {
		return NewTLSavedPhoneContact(nil)
	},
	Cmd_SecureCredentialsEncrypted: func() TLObject {
		return NewTLSecureCredentialsEncrypted(nil)
	},
	Cmd_SecureData: func() TLObject {
		return NewTLSecureData(nil)
	},
	Cmd_SecureFile: func() TLObject {
		return NewTLSecureFile(nil)
	},
	Cmd_SecureFileEmpty: func() TLObject {
		return NewTLSecureFileEmpty(nil)
	},
	Cmd_SecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000: func() TLObject {
		return NewTLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000(nil)
	},
	Cmd_SecurePasswordKdfAlgoSHA512: func() TLObject {
		return NewTLSecurePasswordKdfAlgoSHA512(nil)
	},
	Cmd_SecurePasswordKdfAlgoUnknown: func() TLObject {
		return NewTLSecurePasswordKdfAlgoUnknown(nil)
	},
	Cmd_SecurePlainEmail: func() TLObject {
		return NewTLSecurePlainEmail(nil)
	},
	Cmd_SecurePlainPhone: func() TLObject {
		return NewTLSecurePlainPhone(nil)
	},
	Cmd_SecureRequiredType: func() TLObject {
		return NewTLSecureRequiredType(nil)
	},
	Cmd_SecureRequiredTypeOneOf: func() TLObject {
		return NewTLSecureRequiredTypeOneOf(nil)
	},
	Cmd_SecureSecretSettings: func() TLObject {
		return NewTLSecureSecretSettings(nil)
	},
	Cmd_SecureValue: func() TLObject {
		return NewTLSecureValue(nil)
	},
	Cmd_SecureValueError: func() TLObject {
		return NewTLSecureValueError(nil)
	},
	Cmd_SecureValueErrorData: func() TLObject {
		return NewTLSecureValueErrorData(nil)
	},
	Cmd_SecureValueErrorFile: func() TLObject {
		return NewTLSecureValueErrorFile(nil)
	},
	Cmd_SecureValueErrorFiles: func() TLObject {
		return NewTLSecureValueErrorFiles(nil)
	},
	Cmd_SecureValueErrorFrontSide: func() TLObject {
		return NewTLSecureValueErrorFrontSide(nil)
	},
	Cmd_SecureValueErrorReverseSide: func() TLObject {
		return NewTLSecureValueErrorReverseSide(nil)
	},
	Cmd_SecureValueErrorSelfie: func() TLObject {
		return NewTLSecureValueErrorSelfie(nil)
	},
	Cmd_SecureValueErrorTranslationFile: func() TLObject {
		return NewTLSecureValueErrorTranslationFile(nil)
	},
	Cmd_SecureValueErrorTranslationFiles: func() TLObject {
		return NewTLSecureValueErrorTranslationFiles(nil)
	},
	Cmd_SecureValueHash: func() TLObject {
		return NewTLSecureValueHash(nil)
	},
	Cmd_SecureValueTypeAddress: func() TLObject {
		return NewTLSecureValueTypeAddress(nil)
	},
	Cmd_SecureValueTypeBankStatement: func() TLObject {
		return NewTLSecureValueTypeBankStatement(nil)
	},
	Cmd_SecureValueTypeDriverLicense: func() TLObject {
		return NewTLSecureValueTypeDriverLicense(nil)
	},
	Cmd_SecureValueTypeEmail: func() TLObject {
		return NewTLSecureValueTypeEmail(nil)
	},
	Cmd_SecureValueTypeIdentityCard: func() TLObject {
		return NewTLSecureValueTypeIdentityCard(nil)
	},
	Cmd_SecureValueTypeInternalPassport: func() TLObject {
		return NewTLSecureValueTypeInternalPassport(nil)
	},
	Cmd_SecureValueTypePassport: func() TLObject {
		return NewTLSecureValueTypePassport(nil)
	},
	Cmd_SecureValueTypePassportRegistration: func() TLObject {
		return NewTLSecureValueTypePassportRegistration(nil)
	},
	Cmd_SecureValueTypePersonalDetails: func() TLObject {
		return NewTLSecureValueTypePersonalDetails(nil)
	},
	Cmd_SecureValueTypePhone: func() TLObject {
		return NewTLSecureValueTypePhone(nil)
	},
	Cmd_SecureValueTypeRentalAgreement: func() TLObject {
		return NewTLSecureValueTypeRentalAgreement(nil)
	},
	Cmd_SecureValueTypeTemporaryRegistration: func() TLObject {
		return NewTLSecureValueTypeTemporaryRegistration(nil)
	},
	Cmd_SecureValueTypeUtilityBill: func() TLObject {
		return NewTLSecureValueTypeUtilityBill(nil)
	},
	Cmd_SendMessageCancelAction: func() TLObject {
		return NewTLSendMessageCancelAction(nil)
	},
	Cmd_SendMessageChooseContactAction: func() TLObject {
		return NewTLSendMessageChooseContactAction(nil)
	},
	Cmd_SendMessageGamePlayAction: func() TLObject {
		return NewTLSendMessageGamePlayAction(nil)
	},
	Cmd_SendMessageGeoLocationAction: func() TLObject {
		return NewTLSendMessageGeoLocationAction(nil)
	},
	Cmd_SendMessageRecordAudioAction: func() TLObject {
		return NewTLSendMessageRecordAudioAction(nil)
	},
	Cmd_SendMessageRecordRoundAction: func() TLObject {
		return NewTLSendMessageRecordRoundAction(nil)
	},
	Cmd_SendMessageRecordVideoAction: func() TLObject {
		return NewTLSendMessageRecordVideoAction(nil)
	},
	Cmd_SendMessageTypingAction: func() TLObject {
		return NewTLSendMessageTypingAction(nil)
	},
	Cmd_SendMessageUploadAudioAction: func() TLObject {
		return NewTLSendMessageUploadAudioAction(nil)
	},
	Cmd_SendMessageUploadDocumentAction: func() TLObject {
		return NewTLSendMessageUploadDocumentAction(nil)
	},
	Cmd_SendMessageUploadPhotoAction: func() TLObject {
		return NewTLSendMessageUploadPhotoAction(nil)
	},
	Cmd_SendMessageUploadRoundAction: func() TLObject {
		return NewTLSendMessageUploadRoundAction(nil)
	},
	Cmd_SendMessageUploadVideoAction: func() TLObject {
		return NewTLSendMessageUploadVideoAction(nil)
	},
	Cmd_ShippingOption: func() TLObject {
		return NewTLShippingOption(nil)
	},
	Cmd_SpeakingInGroupCallAction: func() TLObject {
		return NewTLSpeakingInGroupCallAction(nil)
	},
	Cmd_StatsAbsValueAndPrev: func() TLObject {
		return NewTLStatsAbsValueAndPrev(nil)
	},
	Cmd_StatsBroadcastStats: func() TLObject {
		return NewTLStatsBroadcastStats(nil)
	},
	Cmd_StatsDateRangeDays: func() TLObject {
		return NewTLStatsDateRangeDays(nil)
	},
	Cmd_StatsGraph: func() TLObject {
		return NewTLStatsGraph(nil)
	},
	Cmd_StatsGraphAsync: func() TLObject {
		return NewTLStatsGraphAsync(nil)
	},
	Cmd_StatsGraphError: func() TLObject {
		return NewTLStatsGraphError(nil)
	},
	Cmd_StatsGroupTopAdmin: func() TLObject {
		return NewTLStatsGroupTopAdmin(nil)
	},
	Cmd_StatsGroupTopInviter: func() TLObject {
		return NewTLStatsGroupTopInviter(nil)
	},
	Cmd_StatsGroupTopPoster: func() TLObject {
		return NewTLStatsGroupTopPoster(nil)
	},
	Cmd_StatsMegagroupStats: func() TLObject {
		return NewTLStatsMegagroupStats(nil)
	},
	Cmd_StatsMessageStats: func() TLObject {
		return NewTLStatsMessageStats(nil)
	},
	Cmd_StatsPercentValue: func() TLObject {
		return NewTLStatsPercentValue(nil)
	},
	Cmd_StatsURL: func() TLObject {
		return NewTLStatsURL(nil)
	},
	Cmd_StickerPack: func() TLObject {
		return NewTLStickerPack(nil)
	},
	Cmd_StickerSet: func() TLObject {
		return NewTLStickerSet(nil)
	},
	Cmd_StickerSet40e237a8: func() TLObject {
		return NewTLStickerSet(nil)
	},
	Cmd_StickerSet5585a139: func() TLObject {
		return NewTLStickerSet(nil)
	},
	Cmd_StickerSet6a90bcb7: func() TLObject {
		return NewTLStickerSet(nil)
	},
	Cmd_StickerSetCovered: func() TLObject {
		return NewTLStickerSetCovered(nil)
	},
	Cmd_StickerSetMultiCovered: func() TLObject {
		return NewTLStickerSetMultiCovered(nil)
	},
	Cmd_StickerSeteeb46f27: func() TLObject {
		return NewTLStickerSet(nil)
	},
	Cmd_StorageFileGif: func() TLObject {
		return NewTLStorageFileGif(nil)
	},
	Cmd_StorageFileJpeg: func() TLObject {
		return NewTLStorageFileJpeg(nil)
	},
	Cmd_StorageFileMov: func() TLObject {
		return NewTLStorageFileMov(nil)
	},
	Cmd_StorageFileMp3: func() TLObject {
		return NewTLStorageFileMp3(nil)
	},
	Cmd_StorageFileMp4: func() TLObject {
		return NewTLStorageFileMp4(nil)
	},
	Cmd_StorageFilePartial: func() TLObject {
		return NewTLStorageFilePartial(nil)
	},
	Cmd_StorageFilePdf: func() TLObject {
		return NewTLStorageFilePdf(nil)
	},
	Cmd_StorageFilePng: func() TLObject {
		return NewTLStorageFilePng(nil)
	},
	Cmd_StorageFileUnknown: func() TLObject {
		return NewTLStorageFileUnknown(nil)
	},
	Cmd_StorageFileWebp: func() TLObject {
		return NewTLStorageFileWebp(nil)
	},
	Cmd_TextAnchor: func() TLObject {
		return NewTLTextAnchor(nil)
	},
	Cmd_TextBold: func() TLObject {
		return NewTLTextBold(nil)
	},
	Cmd_TextConcat: func() TLObject {
		return NewTLTextConcat(nil)
	},
	Cmd_TextEmail: func() TLObject {
		return NewTLTextEmail(nil)
	},
	Cmd_TextEmpty: func() TLObject {
		return NewTLTextEmpty(nil)
	},
	Cmd_TextFixed: func() TLObject {
		return NewTLTextFixed(nil)
	},
	Cmd_TextImage: func() TLObject {
		return NewTLTextImage(nil)
	},
	Cmd_TextItalic: func() TLObject {
		return NewTLTextItalic(nil)
	},
	Cmd_TextMarked: func() TLObject {
		return NewTLTextMarked(nil)
	},
	Cmd_TextPhone: func() TLObject {
		return NewTLTextPhone(nil)
	},
	Cmd_TextPlain: func() TLObject {
		return NewTLTextPlain(nil)
	},
	Cmd_TextStrike: func() TLObject {
		return NewTLTextStrike(nil)
	},
	Cmd_TextSubscript: func() TLObject {
		return NewTLTextSubscript(nil)
	},
	Cmd_TextSuperscript: func() TLObject {
		return NewTLTextSuperscript(nil)
	},
	Cmd_TextUnderline: func() TLObject {
		return NewTLTextUnderline(nil)
	},
	Cmd_TextUrl: func() TLObject {
		return NewTLTextUrl(nil)
	},
	Cmd_Theme: func() TLObject {
		return NewTLTheme(nil)
	},
	Cmd_Theme28f1114: func() TLObject {
		return NewTLTheme(nil)
	},
	Cmd_ThemeDocumentNotModified: func() TLObject {
		return NewTLThemeDocumentNotModified(nil)
	},
	Cmd_ThemeSettings: func() TLObject {
		return NewTLThemeSettings(nil)
	},
	Cmd_Themef7d90ce0: func() TLObject {
		return NewTLTheme(nil)
	},
	Cmd_TopPeer: func() TLObject {
		return NewTLTopPeer(nil)
	},
	Cmd_TopPeerCategoryBotsInline: func() TLObject {
		return NewTLTopPeerCategoryBotsInline(nil)
	},
	Cmd_TopPeerCategoryBotsPM: func() TLObject {
		return NewTLTopPeerCategoryBotsPM(nil)
	},
	Cmd_TopPeerCategoryChannels: func() TLObject {
		return NewTLTopPeerCategoryChannels(nil)
	},
	Cmd_TopPeerCategoryCorrespondents: func() TLObject {
		return NewTLTopPeerCategoryCorrespondents(nil)
	},
	Cmd_TopPeerCategoryForwardChats: func() TLObject {
		return NewTLTopPeerCategoryForwardChats(nil)
	},
	Cmd_TopPeerCategoryForwardUsers: func() TLObject {
		return NewTLTopPeerCategoryForwardUsers(nil)
	},
	Cmd_TopPeerCategoryGroups: func() TLObject {
		return NewTLTopPeerCategoryGroups(nil)
	},
	Cmd_TopPeerCategoryPeers: func() TLObject {
		return NewTLTopPeerCategoryPeers(nil)
	},
	Cmd_TopPeerCategoryPhoneCalls: func() TLObject {
		return NewTLTopPeerCategoryPhoneCalls(nil)
	},
	Cmd_True: func() TLObject {
		return NewTLTrue(nil)
	},
	Cmd_UpdateBotCallbackQuery: func() TLObject {
		return NewTLUpdateBotCallbackQuery(nil)
	},
	Cmd_UpdateBotCallbackQuerya68c688c: func() TLObject {
		return NewTLUpdateBotCallbackQuery(nil)
	},
	Cmd_UpdateBotInlineQuery: func() TLObject {
		return NewTLUpdateBotInlineQuery(nil)
	},
	Cmd_UpdateBotInlineQuery3f2038db: func() TLObject {
		return NewTLUpdateBotInlineQuery(nil)
	},
	Cmd_UpdateBotInlineSend: func() TLObject {
		return NewTLUpdateBotInlineSend(nil)
	},
	Cmd_UpdateBotPrecheckoutQuery: func() TLObject {
		return NewTLUpdateBotPrecheckoutQuery(nil)
	},
	Cmd_UpdateBotShippingQuery: func() TLObject {
		return NewTLUpdateBotShippingQuery(nil)
	},
	Cmd_UpdateBotWebhookJSON: func() TLObject {
		return NewTLUpdateBotWebhookJSON(nil)
	},
	Cmd_UpdateBotWebhookJSONQuery: func() TLObject {
		return NewTLUpdateBotWebhookJSONQuery(nil)
	},
	Cmd_UpdateChannel: func() TLObject {
		return NewTLUpdateChannel(nil)
	},
	Cmd_UpdateChannelAvailableMessages: func() TLObject {
		return NewTLUpdateChannelAvailableMessages(nil)
	},
	Cmd_UpdateChannelGroup: func() TLObject {
		return NewTLUpdateChannelGroup(nil)
	},
	Cmd_UpdateChannelMessageForwards: func() TLObject {
		return NewTLUpdateChannelMessageForwards(nil)
	},
	Cmd_UpdateChannelMessageViews: func() TLObject {
		return NewTLUpdateChannelMessageViews(nil)
	},
	Cmd_UpdateChannelParticipant: func() TLObject {
		return NewTLUpdateChannelParticipant(nil)
	},
	Cmd_UpdateChannelPinnedMessage: func() TLObject {
		return NewTLUpdateChannelPinnedMessage(nil)
	},
	Cmd_UpdateChannelReadMessagesContents: func() TLObject {
		return NewTLUpdateChannelReadMessagesContents(nil)
	},
	Cmd_UpdateChannelTooLong: func() TLObject {
		return NewTLUpdateChannelTooLong(nil)
	},
	Cmd_UpdateChannelUserTyping: func() TLObject {
		return NewTLUpdateChannelUserTyping(nil)
	},
	Cmd_UpdateChannelWebPage: func() TLObject {
		return NewTLUpdateChannelWebPage(nil)
	},
	Cmd_UpdateChat: func() TLObject {
		return NewTLUpdateChat(nil)
	},
	Cmd_UpdateChatAdmins: func() TLObject {
		return NewTLUpdateChatAdmins(nil)
	},
	Cmd_UpdateChatDefaultBannedRights: func() TLObject {
		return NewTLUpdateChatDefaultBannedRights(nil)
	},
	Cmd_UpdateChatParticipantAdd: func() TLObject {
		return NewTLUpdateChatParticipantAdd(nil)
	},
	Cmd_UpdateChatParticipantAdmin: func() TLObject {
		return NewTLUpdateChatParticipantAdmin(nil)
	},
	Cmd_UpdateChatParticipantDelete: func() TLObject {
		return NewTLUpdateChatParticipantDelete(nil)
	},
	Cmd_UpdateChatParticipants: func() TLObject {
		return NewTLUpdateChatParticipants(nil)
	},
	Cmd_UpdateChatPinnedMessage: func() TLObject {
		return NewTLUpdateChatPinnedMessage(nil)
	},
	Cmd_UpdateChatPinnedMessagee10db349: func() TLObject {
		return NewTLUpdateChatPinnedMessage(nil)
	},
	Cmd_UpdateChatUserTyping: func() TLObject {
		return NewTLUpdateChatUserTyping(nil)
	},
	Cmd_UpdateConfig: func() TLObject {
		return NewTLUpdateConfig(nil)
	},
	Cmd_UpdateContactLink: func() TLObject {
		return NewTLUpdateContactLink(nil)
	},
	Cmd_UpdateContactRegistered: func() TLObject {
		return NewTLUpdateContactRegistered(nil)
	},
	Cmd_UpdateContactsReset: func() TLObject {
		return NewTLUpdateContactsReset(nil)
	},
	Cmd_UpdateDcOptions: func() TLObject {
		return NewTLUpdateDcOptions(nil)
	},
	Cmd_UpdateDeleteChannelMessages: func() TLObject {
		return NewTLUpdateDeleteChannelMessages(nil)
	},
	Cmd_UpdateDeleteMessages: func() TLObject {
		return NewTLUpdateDeleteMessages(nil)
	},
	Cmd_UpdateDeleteScheduledMessages: func() TLObject {
		return NewTLUpdateDeleteScheduledMessages(nil)
	},
	Cmd_UpdateDialogFilter: func() TLObject {
		return NewTLUpdateDialogFilter(nil)
	},
	Cmd_UpdateDialogFilterOrder: func() TLObject {
		return NewTLUpdateDialogFilterOrder(nil)
	},
	Cmd_UpdateDialogFilters: func() TLObject {
		return NewTLUpdateDialogFilters(nil)
	},
	Cmd_UpdateDialogPinned: func() TLObject {
		return NewTLUpdateDialogPinned(nil)
	},
	Cmd_UpdateDialogPinned19d27f3c: func() TLObject {
		return NewTLUpdateDialogPinned(nil)
	},
	Cmd_UpdateDialogPinned6e6fe51c: func() TLObject {
		return NewTLUpdateDialogPinned(nil)
	},
	Cmd_UpdateDialogUnreadMark: func() TLObject {
		return NewTLUpdateDialogUnreadMark(nil)
	},
	Cmd_UpdateDraftMessage: func() TLObject {
		return NewTLUpdateDraftMessage(nil)
	},
	Cmd_UpdateEditChannelMessage: func() TLObject {
		return NewTLUpdateEditChannelMessage(nil)
	},
	Cmd_UpdateEditMessage: func() TLObject {
		return NewTLUpdateEditMessage(nil)
	},
	Cmd_UpdateEncryptedChatTyping: func() TLObject {
		return NewTLUpdateEncryptedChatTyping(nil)
	},
	Cmd_UpdateEncryptedMessagesRead: func() TLObject {
		return NewTLUpdateEncryptedMessagesRead(nil)
	},
	Cmd_UpdateEncryption: func() TLObject {
		return NewTLUpdateEncryption(nil)
	},
	Cmd_UpdateFavedStickers: func() TLObject {
		return NewTLUpdateFavedStickers(nil)
	},
	Cmd_UpdateFolderPeers: func() TLObject {
		return NewTLUpdateFolderPeers(nil)
	},
	Cmd_UpdateGeoLiveViewed: func() TLObject {
		return NewTLUpdateGeoLiveViewed(nil)
	},
	Cmd_UpdateGroupCall: func() TLObject {
		return NewTLUpdateGroupCall(nil)
	},
	Cmd_UpdateGroupCallParticipants: func() TLObject {
		return NewTLUpdateGroupCallParticipants(nil)
	},
	Cmd_UpdateInlineBotCallbackQuery: func() TLObject {
		return NewTLUpdateInlineBotCallbackQuery(nil)
	},
	Cmd_UpdateInlineBotCallbackQuery2cbd95af: func() TLObject {
		return NewTLUpdateInlineBotCallbackQuery(nil)
	},
	Cmd_UpdateLangPack: func() TLObject {
		return NewTLUpdateLangPack(nil)
	},
	Cmd_UpdateLangPackTooLong: func() TLObject {
		return NewTLUpdateLangPackTooLong(nil)
	},
	Cmd_UpdateLangPackTooLong46560264: func() TLObject {
		return NewTLUpdateLangPackTooLong(nil)
	},
	Cmd_UpdateLoginToken: func() TLObject {
		return NewTLUpdateLoginToken(nil)
	},
	Cmd_UpdateMessageID: func() TLObject {
		return NewTLUpdateMessageID(nil)
	},
	Cmd_UpdateMessagePoll: func() TLObject {
		return NewTLUpdateMessagePoll(nil)
	},
	Cmd_UpdateMessagePollVote: func() TLObject {
		return NewTLUpdateMessagePollVote(nil)
	},
	Cmd_UpdateNewAuthorization: func() TLObject {
		return NewTLUpdateNewAuthorization(nil)
	},
	Cmd_UpdateNewChannelMessage: func() TLObject {
		return NewTLUpdateNewChannelMessage(nil)
	},
	Cmd_UpdateNewEncryptedMessage: func() TLObject {
		return NewTLUpdateNewEncryptedMessage(nil)
	},
	Cmd_UpdateNewMessage: func() TLObject {
		return NewTLUpdateNewMessage(nil)
	},
	Cmd_UpdateNewScheduledMessage: func() TLObject {
		return NewTLUpdateNewScheduledMessage(nil)
	},
	Cmd_UpdateNewStickerSet: func() TLObject {
		return NewTLUpdateNewStickerSet(nil)
	},
	Cmd_UpdateNotifySettings: func() TLObject {
		return NewTLUpdateNotifySettings(nil)
	},
	Cmd_UpdatePeerBlocked: func() TLObject {
		return NewTLUpdatePeerBlocked(nil)
	},
	Cmd_UpdatePeerLocated: func() TLObject {
		return NewTLUpdatePeerLocated(nil)
	},
	Cmd_UpdatePeerSettings: func() TLObject {
		return NewTLUpdatePeerSettings(nil)
	},
	Cmd_UpdatePhoneCall: func() TLObject {
		return NewTLUpdatePhoneCall(nil)
	},
	Cmd_UpdatePhoneCallSignalingData: func() TLObject {
		return NewTLUpdatePhoneCallSignalingData(nil)
	},
	Cmd_UpdatePinnedChannelMessages: func() TLObject {
		return NewTLUpdatePinnedChannelMessages(nil)
	},
	Cmd_UpdatePinnedDialogs: func() TLObject {
		return NewTLUpdatePinnedDialogs(nil)
	},
	Cmd_UpdatePinnedDialogsea4cb65b: func() TLObject {
		return NewTLUpdatePinnedDialogs(nil)
	},
	Cmd_UpdatePinnedDialogsfa0f3ca2: func() TLObject {
		return NewTLUpdatePinnedDialogs(nil)
	},
	Cmd_UpdatePinnedMessages: func() TLObject {
		return NewTLUpdatePinnedMessages(nil)
	},
	Cmd_UpdatePrivacy: func() TLObject {
		return NewTLUpdatePrivacy(nil)
	},
	Cmd_UpdatePtsChanged: func() TLObject {
		return NewTLUpdatePtsChanged(nil)
	},
	Cmd_UpdateReadChannelDiscussionInbox: func() TLObject {
		return NewTLUpdateReadChannelDiscussionInbox(nil)
	},
	Cmd_UpdateReadChannelDiscussionOutbox: func() TLObject {
		return NewTLUpdateReadChannelDiscussionOutbox(nil)
	},
	Cmd_UpdateReadChannelInbox: func() TLObject {
		return NewTLUpdateReadChannelInbox(nil)
	},
	Cmd_UpdateReadChannelInbox330b5424: func() TLObject {
		return NewTLUpdateReadChannelInbox(nil)
	},
	Cmd_UpdateReadChannelOutbox: func() TLObject {
		return NewTLUpdateReadChannelOutbox(nil)
	},
	Cmd_UpdateReadFeaturedStickers: func() TLObject {
		return NewTLUpdateReadFeaturedStickers(nil)
	},
	Cmd_UpdateReadHistoryInbox: func() TLObject {
		return NewTLUpdateReadHistoryInbox(nil)
	},
	Cmd_UpdateReadHistoryInbox9c974fdf: func() TLObject {
		return NewTLUpdateReadHistoryInbox(nil)
	},
	Cmd_UpdateReadHistoryOutbox: func() TLObject {
		return NewTLUpdateReadHistoryOutbox(nil)
	},
	Cmd_UpdateReadMessagesContents: func() TLObject {
		return NewTLUpdateReadMessagesContents(nil)
	},
	Cmd_UpdateRecentStickers: func() TLObject {
		return NewTLUpdateRecentStickers(nil)
	},
	Cmd_UpdateSavedGifs: func() TLObject {
		return NewTLUpdateSavedGifs(nil)
	},
	Cmd_UpdateServiceNotification: func() TLObject {
		return NewTLUpdateServiceNotification(nil)
	},
	Cmd_UpdateServiceNotification382dd3e4: func() TLObject {
		return NewTLUpdateServiceNotification(nil)
	},
	Cmd_UpdateShort: func() TLObject {
		return NewTLUpdateShort(nil)
	},
	Cmd_UpdateShortChatMessage: func() TLObject {
		return NewTLUpdateShortChatMessage(nil)
	},
	Cmd_UpdateShortChatMessage402d5dbb: func() TLObject {
		return NewTLUpdateShortChatMessage(nil)
	},
	Cmd_UpdateShortMessage: func() TLObject {
		return NewTLUpdateShortMessage(nil)
	},
	Cmd_UpdateShortMessage2296d2c8: func() TLObject {
		return NewTLUpdateShortMessage(nil)
	},
	Cmd_UpdateShortSentMessage: func() TLObject {
		return NewTLUpdateShortSentMessage(nil)
	},
	Cmd_UpdateStickerSets: func() TLObject {
		return NewTLUpdateStickerSets(nil)
	},
	Cmd_UpdateStickerSetsOrder: func() TLObject {
		return NewTLUpdateStickerSetsOrder(nil)
	},
	Cmd_UpdateStickerSetsOrderf0dfb451: func() TLObject {
		return NewTLUpdateStickerSetsOrder(nil)
	},
	Cmd_UpdateTheme: func() TLObject {
		return NewTLUpdateTheme(nil)
	},
	Cmd_UpdateUserBlocked: func() TLObject {
		return NewTLUpdateUserBlocked(nil)
	},
	Cmd_UpdateUserName: func() TLObject {
		return NewTLUpdateUserName(nil)
	},
	Cmd_UpdateUserPhone: func() TLObject {
		return NewTLUpdateUserPhone(nil)
	},
	Cmd_UpdateUserPhoto: func() TLObject {
		return NewTLUpdateUserPhoto(nil)
	},
	Cmd_UpdateUserPinnedMessage: func() TLObject {
		return NewTLUpdateUserPinnedMessage(nil)
	},
	Cmd_UpdateUserStatus: func() TLObject {
		return NewTLUpdateUserStatus(nil)
	},
	Cmd_UpdateUserTyping: func() TLObject {
		return NewTLUpdateUserTyping(nil)
	},
	Cmd_UpdateWebPage: func() TLObject {
		return NewTLUpdateWebPage(nil)
	},
	Cmd_Updates: func() TLObject {
		return NewTLUpdates(nil)
	},
	Cmd_UpdatesChannelDifference: func() TLObject {
		return NewTLUpdatesChannelDifference(nil)
	},
	Cmd_UpdatesChannelDifferenceEmpty: func() TLObject {
		return NewTLUpdatesChannelDifferenceEmpty(nil)
	},
	Cmd_UpdatesChannelDifferenceTooLong: func() TLObject {
		return NewTLUpdatesChannelDifferenceTooLong(nil)
	},
	Cmd_UpdatesChannelDifferenceTooLong5e167646: func() TLObject {
		return NewTLUpdatesChannelDifferenceTooLong(nil)
	},
	Cmd_UpdatesChannelDifferenceTooLonga4bcc6fe: func() TLObject {
		return NewTLUpdatesChannelDifferenceTooLong(nil)
	},
	Cmd_UpdatesCombined: func() TLObject {
		return NewTLUpdatesCombined(nil)
	},
	Cmd_UpdatesDifference: func() TLObject {
		return NewTLUpdatesDifference(nil)
	},
	Cmd_UpdatesDifferenceEmpty: func() TLObject {
		return NewTLUpdatesDifferenceEmpty(nil)
	},
	Cmd_UpdatesDifferenceSlice: func() TLObject {
		return NewTLUpdatesDifferenceSlice(nil)
	},
	Cmd_UpdatesDifferenceTooLong: func() TLObject {
		return NewTLUpdatesDifferenceTooLong(nil)
	},
	Cmd_UpdatesState: func() TLObject {
		return NewTLUpdatesState(nil)
	},
	Cmd_UpdatesTooLong: func() TLObject {
		return NewTLUpdatesTooLong(nil)
	},
	Cmd_UploadCdnFile: func() TLObject {
		return NewTLUploadCdnFile(nil)
	},
	Cmd_UploadCdnFileReuploadNeeded: func() TLObject {
		return NewTLUploadCdnFileReuploadNeeded(nil)
	},
	Cmd_UploadFile: func() TLObject {
		return NewTLUploadFile(nil)
	},
	Cmd_UploadFileCdnRedirect: func() TLObject {
		return NewTLUploadFileCdnRedirect(nil)
	},
	Cmd_UploadFileCdnRedirectf18cda44: func() TLObject {
		return NewTLUploadFileCdnRedirect(nil)
	},
	Cmd_UploadWebFile: func() TLObject {
		return NewTLUploadWebFile(nil)
	},
	Cmd_UrlAuthResultAccepted: func() TLObject {
		return NewTLUrlAuthResultAccepted(nil)
	},
	Cmd_UrlAuthResultDefault: func() TLObject {
		return NewTLUrlAuthResultDefault(nil)
	},
	Cmd_UrlAuthResultRequest: func() TLObject {
		return NewTLUrlAuthResultRequest(nil)
	},
	Cmd_User: func() TLObject {
		return NewTLUser(nil)
	},
	Cmd_User938458c1: func() TLObject {
		return NewTLUser(nil)
	},
	Cmd_UserEmpty: func() TLObject {
		return NewTLUserEmpty(nil)
	},
	Cmd_UserFull: func() TLObject {
		return NewTLUserFull(nil)
	},
	Cmd_UserFull5932fc03: func() TLObject {
		return NewTLUserFull(nil)
	},
	Cmd_UserFull745559cc: func() TLObject {
		return NewTLUserFull(nil)
	},
	Cmd_UserFull8ea4a881: func() TLObject {
		return NewTLUserFull(nil)
	},
	Cmd_UserFulledf17c12: func() TLObject {
		return NewTLUserFull(nil)
	},
	Cmd_UserProfilePhoto: func() TLObject {
		return NewTLUserProfilePhoto(nil)
	},
	Cmd_UserProfilePhoto69d3ab26: func() TLObject {
		return NewTLUserProfilePhoto(nil)
	},
	Cmd_UserProfilePhotoEmpty: func() TLObject {
		return NewTLUserProfilePhotoEmpty(nil)
	},
	Cmd_UserProfilePhotoecd75d8c: func() TLObject {
		return NewTLUserProfilePhoto(nil)
	},
	Cmd_UserStatusEmpty: func() TLObject {
		return NewTLUserStatusEmpty(nil)
	},
	Cmd_UserStatusLastMonth: func() TLObject {
		return NewTLUserStatusLastMonth(nil)
	},
	Cmd_UserStatusLastWeek: func() TLObject {
		return NewTLUserStatusLastWeek(nil)
	},
	Cmd_UserStatusOffline: func() TLObject {
		return NewTLUserStatusOffline(nil)
	},
	Cmd_UserStatusOnline: func() TLObject {
		return NewTLUserStatusOnline(nil)
	},
	Cmd_UserStatusRecently: func() TLObject {
		return NewTLUserStatusRecently(nil)
	},
	Cmd_Userd10d979a: func() TLObject {
		return NewTLUser(nil)
	},
	Cmd_VideoSize: func() TLObject {
		return NewTLVideoSize(nil)
	},
	Cmd_VideoSizee831c556: func() TLObject {
		return NewTLVideoSize(nil)
	},
	Cmd_WallPaper: func() TLObject {
		return NewTLWallPaper(nil)
	},
	Cmd_WallPaperNoFile: func() TLObject {
		return NewTLWallPaperNoFile(nil)
	},
	Cmd_WallPaperSettings: func() TLObject {
		return NewTLWallPaperSettings(nil)
	},
	Cmd_WallPaperSettings5086cf8: func() TLObject {
		return NewTLWallPaperSettings(nil)
	},
	Cmd_WallPaperSolid: func() TLObject {
		return NewTLWallPaperSolid(nil)
	},
	Cmd_WallPapera437c3ed: func() TLObject {
		return NewTLWallPaper(nil)
	},
	Cmd_WallPaperf04f91ec: func() TLObject {
		return NewTLWallPaper(nil)
	},
	Cmd_WalletLiteResponse: func() TLObject {
		return NewTLWalletLiteResponse(nil)
	},
	Cmd_WalletSecretSalt: func() TLObject {
		return NewTLWalletSecretSalt(nil)
	},
	Cmd_WebAuthorization: func() TLObject {
		return NewTLWebAuthorization(nil)
	},
	Cmd_WebDocument: func() TLObject {
		return NewTLWebDocument(nil)
	},
	Cmd_WebDocument1c570ed1: func() TLObject {
		return NewTLWebDocument(nil)
	},
	Cmd_WebDocumentNoProxy: func() TLObject {
		return NewTLWebDocumentNoProxy(nil)
	},
	Cmd_WebPage: func() TLObject {
		return NewTLWebPage(nil)
	},
	Cmd_WebPageAttributeTheme: func() TLObject {
		return NewTLWebPageAttributeTheme(nil)
	},
	Cmd_WebPageEmpty: func() TLObject {
		return NewTLWebPageEmpty(nil)
	},
	Cmd_WebPageNotModified: func() TLObject {
		return NewTLWebPageNotModified(nil)
	},
	Cmd_WebPageNotModified7311ca11: func() TLObject {
		return NewTLWebPageNotModified(nil)
	},
	Cmd_WebPagePending: func() TLObject {
		return NewTLWebPagePending(nil)
	},
	Cmd_WebPageca820ed7: func() TLObject {
		return NewTLWebPage(nil)
	},
	Cmd_WebPagee89c45b2: func() TLObject {
		return NewTLWebPage(nil)
	},
	Cmd_WebPagefa64e172: func() TLObject {
		return NewTLWebPage(nil)
	},
}

func GetClassById(cmdId TLCmd) TLObject {
	maker, ok := registerObject[cmdId]
	if !ok {
		err := fmt.Errorf("GetClassById not found cmdId:%x", cmdId)
		log.Errorf(err.Error())
		return nil
	}
	return maker()
}
