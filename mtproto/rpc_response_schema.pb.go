/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2023/06/12 18:03
 * @File : codec_schema.service.pb.go
 * @Desc :
 *
 */

package mtproto

import (
	"fmt"
	"novachat_engine/pkg/log"
)

var registerMessageObject = map[string]func() TLObject{
	"mtproto.AccountDaysTTL": func() TLObject {
		return &AccountDaysTTL{}
	},
	"mtproto.Authorization": func() TLObject {
		return &Authorization{}
	},
	"mtproto.Bool": func() TLObject {
		return &Bool{}
	},
	"mtproto.CdnConfig": func() TLObject {
		return &CdnConfig{}
	},
	"mtproto.ChatInvite": func() TLObject {
		return &ChatInvite{}
	},
	"mtproto.ChatOnlines": func() TLObject {
		return &ChatOnlines{}
	},
	"mtproto.Config": func() TLObject {
		return &Config{}
	},
	"mtproto.DataJSON": func() TLObject {
		return &DataJSON{}
	},
	"mtproto.Document": func() TLObject {
		return &Document{}
	},
	"mtproto.EmojiKeywordsDifference": func() TLObject {
		return &EmojiKeywordsDifference{}
	},
	"mtproto.EmojiURL": func() TLObject {
		return &EmojiURL{}
	},
	"mtproto.EncryptedChat": func() TLObject {
		return &EncryptedChat{}
	},
	"mtproto.EncryptedFile": func() TLObject {
		return &EncryptedFile{}
	},
	"mtproto.ExportedChatInvite": func() TLObject {
		return &ExportedChatInvite{}
	},
	"mtproto.ExportedMessageLink": func() TLObject {
		return &ExportedMessageLink{}
	},
	"mtproto.GlobalPrivacySettings": func() TLObject {
		return &GlobalPrivacySettings{}
	},
	"mtproto.JSONValue": func() TLObject {
		return &JSONValue{}
	},
	"mtproto.LangPackDifference": func() TLObject {
		return &LangPackDifference{}
	},
	"mtproto.LangPackLanguage": func() TLObject {
		return &LangPackLanguage{}
	},
	"mtproto.MessageMedia": func() TLObject {
		return &MessageMedia{}
	},
	"mtproto.NearestDc": func() TLObject {
		return &NearestDc{}
	},
	"mtproto.PeerNotifySettings": func() TLObject {
		return &PeerNotifySettings{}
	},
	"mtproto.PeerSettings": func() TLObject {
		return &PeerSettings{}
	},
	"mtproto.SecureValue": func() TLObject {
		return &SecureValue{}
	},
	"mtproto.StatsGraph": func() TLObject {
		return &StatsGraph{}
	},
	"mtproto.StatsURL": func() TLObject {
		return &StatsURL{}
	},
	"mtproto.Theme": func() TLObject {
		return &Theme{}
	},
	"mtproto.Updates": func() TLObject {
		return &Updates{}
	},
	"mtproto.UrlAuthResult": func() TLObject {
		return &UrlAuthResult{}
	},
	"mtproto.User": func() TLObject {
		return &User{}
	},
	"mtproto.UserFull": func() TLObject {
		return &UserFull{}
	},
	"mtproto.UserProfilePhoto": func() TLObject {
		return &Response_PhotosUpdateProfilePhoto{}
	},
	"mtproto.Vector_CdnFileHash": func() TLObject {
		return &Response_UploadReuploadCdnFile{}
	},
	"mtproto.Vector_ContactStatus": func() TLObject {
		return &Vector_ContactStatus{}
	},
	"mtproto.Vector_DialogFilter": func() TLObject {
		return &Vector_DialogFilter{}
	},
	"mtproto.Vector_DialogFilterSuggested": func() TLObject {
		return &Vector_DialogFilterSuggested{}
	},
	"mtproto.Vector_DialogPeer": func() TLObject {
		return &Vector_DialogPeer{}
	},
	"mtproto.Vector_EmojiLanguage": func() TLObject {
		return &Vector_EmojiLanguage{}
	},
	"mtproto.Vector_FileHash": func() TLObject {
		return &Vector_FileHash{}
	},
	"mtproto.Vector_LangPackLanguage": func() TLObject {
		return &Vector_LangPackLanguage{}
	},
	"mtproto.Vector_LangPackString": func() TLObject {
		return &Vector_LangPackString{}
	},
	"mtproto.Vector_MessageRange": func() TLObject {
		return &Vector_MessageRange{}
	},
	"mtproto.Vector_ReceivedNotifyMessage": func() TLObject {
		return &Vector_ReceivedNotifyMessage{}
	},
	"mtproto.Vector_SavedContact": func() TLObject {
		return &Vector_SavedContact{}
	},
	"mtproto.Vector_SecureValue": func() TLObject {
		return &Vector_SecureValue{}
	},
	"mtproto.Vector_StickerSetCovered": func() TLObject {
		return &Vector_StickerSetCovered{}
	},
	"mtproto.Vector_User": func() TLObject {
		return &Vector_User{}
	},
	"mtproto.Vector_WallPaper": func() TLObject {
		return &Response_AccountGetWallPapers{}
	},
	"mtproto.Vector_int": func() TLObject {
		return &Response_MessagesGetMessagesViews{}
	},
	"mtproto.Vector_long": func() TLObject {
		return &VectorLong{}
	},
	"mtproto.Vector_messages_SearchCounter": func() TLObject {
		return &VectorMessages_SearchCounter{}
	},
	"mtproto.WallPaper": func() TLObject {
		return &WallPaper{}
	},
	"mtproto.WebPage": func() TLObject {
		return &WebPage{}
	},
	"mtproto.account_AuthorizationForm": func() TLObject {
		return &Account_AuthorizationForm{}
	},
	"mtproto.account_Authorizations": func() TLObject {
		return &Account_Authorizations{}
	},
	"mtproto.account_AutoDownloadSettings": func() TLObject {
		return &Account_AutoDownloadSettings{}
	},
	"mtproto.account_ContentSettings": func() TLObject {
		return &Account_ContentSettings{}
	},
	"mtproto.account_Password": func() TLObject {
		return &Account_Password{}
	},
	"mtproto.account_PasswordSettings": func() TLObject {
		return &Account_PasswordSettings{}
	},
	"mtproto.account_PrivacyRules": func() TLObject {
		return &Account_PrivacyRules{}
	},
	"mtproto.account_SentEmailCode": func() TLObject {
		return &Account_SentEmailCode{}
	},
	"mtproto.account_Takeout": func() TLObject {
		return &Account_Takeout{}
	},
	"mtproto.account_Themes": func() TLObject {
		return &Account_Themes{}
	},
	"mtproto.account_TmpPassword": func() TLObject {
		return &Account_TmpPassword{}
	},
	"mtproto.account_WebAuthorizations": func() TLObject {
		return &Account_WebAuthorizations{}
	},
	"mtproto.auth_Authorization": func() TLObject {
		return &Auth_Authorization{}
	},
	"mtproto.auth_CheckedPhone": func() TLObject {
		return &Auth_CheckedPhone{}
	},
	"mtproto.auth_ExportedAuthorization": func() TLObject {
		return &Auth_ExportedAuthorization{}
	},
	"mtproto.auth_LoginToken": func() TLObject {
		return &Auth_LoginToken{}
	},
	"mtproto.auth_PasswordRecovery": func() TLObject {
		return &Auth_PasswordRecovery{}
	},
	"mtproto.auth_SentCode": func() TLObject {
		return &Auth_SentCode{}
	},
	"mtproto.channels_AdminLogResults": func() TLObject {
		return &Channels_AdminLogResults{}
	},
	"mtproto.channels_ChannelParticipant": func() TLObject {
		return &Channels_ChannelParticipant{}
	},
	"mtproto.channels_ChannelParticipants": func() TLObject {
		return &Channels_ChannelParticipants{}
	},
	"mtproto.contacts_Blocked": func() TLObject {
		return &Contacts_Blocked{}
	},
	"mtproto.contacts_Contacts": func() TLObject {
		return &Contacts_Contacts{}
	},
	"mtproto.contacts_Found": func() TLObject {
		return &Contacts_Found{}
	},
	"mtproto.contacts_ImportedContacts": func() TLObject {
		return &Contacts_ImportedContacts{}
	},
	"mtproto.contacts_Link": func() TLObject {
		return &Contacts_Link{}
	},
	"mtproto.contacts_ResolvedPeer": func() TLObject {
		return &Contacts_ResolvedPeer{}
	},
	"mtproto.contacts_TopPeers": func() TLObject {
		return &Contacts_TopPeers{}
	},
	"mtproto.help_AppUpdate": func() TLObject {
		return &Help_AppUpdate{}
	},
	"mtproto.help_CountriesList": func() TLObject {
		return &Help_CountriesList{}
	},
	"mtproto.help_DeepLinkInfo": func() TLObject {
		return &Help_DeepLinkInfo{}
	},
	"mtproto.help_InviteText": func() TLObject {
		return &Help_InviteText{}
	},
	"mtproto.help_PassportConfig": func() TLObject {
		return &Help_PassportConfig{}
	},
	"mtproto.help_PromoData": func() TLObject {
		return &Help_PromoData{}
	},
	"mtproto.help_ProxyData": func() TLObject {
		return &Help_ProxyData{}
	},
	"mtproto.help_RecentMeUrls": func() TLObject {
		return &Help_RecentMeUrls{}
	},
	"mtproto.help_Support": func() TLObject {
		return &Help_Support{}
	},
	"mtproto.help_SupportName": func() TLObject {
		return &Help_SupportName{}
	},
	"mtproto.help_TermsOfService": func() TLObject {
		return &Help_TermsOfService{}
	},
	"mtproto.help_TermsOfServiceUpdate": func() TLObject {
		return &Help_TermsOfServiceUpdate{}
	},
	"mtproto.help_UserInfo": func() TLObject {
		return &Help_UserInfo{}
	},
	"mtproto.messages_AffectedHistory": func() TLObject {
		return &Messages_AffectedHistory{}
	},
	"mtproto.messages_AffectedMessages": func() TLObject {
		return &Response_MessagesReadHistory{}
	},
	"mtproto.messages_AllStickers": func() TLObject {
		return &Messages_AllStickers{}
	},
	"mtproto.messages_ArchivedStickers": func() TLObject {
		return &Messages_ArchivedStickers{}
	},
	"mtproto.messages_BotCallbackAnswer": func() TLObject {
		return &Messages_BotCallbackAnswer{}
	},
	"mtproto.messages_BotResults": func() TLObject {
		return &Messages_BotResults{}
	},
	"mtproto.messages_ChatFull": func() TLObject {
		return &Messages_ChatFull{}
	},
	"mtproto.messages_Chats": func() TLObject {
		return &Messages_Chats{}
	},
	"mtproto.messages_DhConfig": func() TLObject {
		return &Messages_DhConfig{}
	},
	"mtproto.messages_Dialogs": func() TLObject {
		return &Messages_Dialogs{}
	},
	"mtproto.messages_DiscussionMessage": func() TLObject {
		return &Messages_DiscussionMessage{}
	},
	"mtproto.messages_FavedStickers": func() TLObject {
		return &Messages_FavedStickers{}
	},
	"mtproto.messages_FeaturedStickers": func() TLObject {
		return &Messages_FeaturedStickers{}
	},
	"mtproto.messages_FoundGifs": func() TLObject {
		return &Messages_FoundGifs{}
	},
	"mtproto.messages_FoundStickerSets": func() TLObject {
		return &Messages_FoundStickerSets{}
	},
	"mtproto.messages_HighScores": func() TLObject {
		return &Messages_HighScores{}
	},
	"mtproto.messages_InactiveChats": func() TLObject {
		return &Messages_InactiveChats{}
	},
	"mtproto.messages_MessageEditData": func() TLObject {
		return &Messages_MessageEditData{}
	},
	"mtproto.messages_Messages": func() TLObject {
		return &Messages_Messages{}
	},
	"mtproto.messages_PeerDialogs": func() TLObject {
		return &Messages_PeerDialogs{}
	},
	"mtproto.messages_RecentStickers": func() TLObject {
		return &Messages_RecentStickers{}
	},
	"mtproto.messages_SavedGifs": func() TLObject {
		return &Messages_SavedGifs{}
	},
	"mtproto.messages_SentEncryptedMessage": func() TLObject {
		return &Messages_SentEncryptedMessage{}
	},
	"mtproto.messages_StickerSet": func() TLObject {
		return &Messages_StickerSet{}
	},
	"mtproto.messages_StickerSetInstallResult": func() TLObject {
		return &Response_MessagesInstallStickerSet{}
	},
	"mtproto.messages_Stickers": func() TLObject {
		return &Messages_Stickers{}
	},
	"mtproto.messages_VotesList": func() TLObject {
		return &Messages_VotesList{}
	},
	"mtproto.payments_BankCardData": func() TLObject {
		return &Payments_BankCardData{}
	},
	"mtproto.payments_PaymentForm": func() TLObject {
		return &Payments_PaymentForm{}
	},
	"mtproto.payments_PaymentReceipt": func() TLObject {
		return &Payments_PaymentReceipt{}
	},
	"mtproto.payments_PaymentResult": func() TLObject {
		return &Payments_PaymentResult{}
	},
	"mtproto.payments_SavedInfo": func() TLObject {
		return &Payments_SavedInfo{}
	},
	"mtproto.payments_ValidatedRequestedInfo": func() TLObject {
		return &Payments_ValidatedRequestedInfo{}
	},
	"mtproto.phone_GroupCall": func() TLObject {
		return &Phone_GroupCall{}
	},
	"mtproto.phone_GroupParticipants": func() TLObject {
		return &Phone_GroupParticipants{}
	},
	"mtproto.phone_PhoneCall": func() TLObject {
		return &Phone_PhoneCall{}
	},
	"mtproto.photos_Photo": func() TLObject {
		return &Photos_Photo{}
	},
	"mtproto.photos_Photos": func() TLObject {
		return &Photos_Photos{}
	},
	"mtproto.stats_BroadcastStats": func() TLObject {
		return &Stats_BroadcastStats{}
	},
	"mtproto.stats_MegagroupStats": func() TLObject {
		return &Stats_MegagroupStats{}
	},
	"mtproto.stats_MessageStats": func() TLObject {
		return &Stats_MessageStats{}
	},
	"mtproto.updates_ChannelDifference": func() TLObject {
		return &Updates_ChannelDifference{}
	},
	"mtproto.updates_Difference": func() TLObject {
		return &Updates_Difference{}
	},
	"mtproto.updates_State": func() TLObject {
		return &Updates_State{}
	},
	"mtproto.upload_CdnFile": func() TLObject {
		return &Upload_CdnFile{}
	},
	"mtproto.upload_File": func() TLObject {
		return &Upload_File{}
	},
	"mtproto.upload_WebFile": func() TLObject {
		return &Upload_WebFile{}
	},
	"mtproto.wallet_KeySecretSalt": func() TLObject {
		return &Wallet_KeySecretSalt{}
	},
	"mtproto.wallet_LiteResponse": func() TLObject {
		return &Wallet_LiteResponse{}
	},
}

func GetClassByMessage(messageName string) TLObject {
	maker, ok := registerMessageObject[messageName]
	if !ok {
		err := fmt.Errorf("GetClassByMessage not found messageName:%s", messageName)
		log.Errorf(err.Error())
		return nil
	}
	return maker()
}
