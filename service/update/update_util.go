/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/31 17:48
 * @File : update_util.go
 */

package update

import (
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
)

type UpdateVersion int32

const (
	UpdateVersion0 UpdateVersion = 0
	UpdateVersion1 UpdateVersion = 1
)

func FromInt32UpdateVersion(m int32) UpdateVersion {
	return UpdateVersion(m)
}

func (m UpdateVersion) ToInt32() int32 {
	return int32(m)
}

func UpdatesToUpdateType(u *mtproto.Updates) constants.UpdateType {
	switch u.ClassName {
	case mtproto.ClassUpdates:
		return constants.UpdateTypeUpdates
	case mtproto.ClassUpdatesTooLong:
		return constants.UpdateTypeUpdatesTooLong
	case mtproto.ClassUpdateShortMessage:
		return constants.UpdateTypeUpdateShortMessage
	case mtproto.ClassUpdateShortChatMessage:
		return constants.UpdateTypeUpdateShortChatMessage
	case mtproto.ClassUpdateShort:
		return constants.UpdateTypeUpdateShort
	case mtproto.ClassUpdatesCombined:
		return constants.UpdateTypeUpdatesCombined
	case mtproto.ClassUpdateShortSentMessage:
		return constants.UpdateTypeUpdateShortSentMessage
	default:
		return constants.UpdateTypeNone
	}
}

func ToUpdateType(u *mtproto.Update) constants.UpdateType {
	value := constants.UpdateTypeNone

	switch u.ClassName {
	case mtproto.ClassUpdateNewMessage:
		value = constants.UpdateTypeUpdateNewMessage
	case mtproto.ClassUpdateMessageID:
		value = constants.UpdateTypeUpdateMessageID
	case mtproto.ClassUpdateDeleteMessages:
		value = constants.UpdateTypeUpdateDeleteMessages
	case mtproto.ClassUpdateUserTyping:
		value = constants.UpdateTypeUpdateUserTyping
	case mtproto.ClassUpdateChatUserTyping:
		value = constants.UpdateTypeUpdateChatUserTyping
	case mtproto.ClassUpdateChatParticipants:
		value = constants.UpdateTypeUpdateChatParticipants
	case mtproto.ClassUpdateUserStatus:
		value = constants.UpdateTypeUpdateUserStatus
	case mtproto.ClassUpdateUserName:
		value = constants.UpdateTypeUpdateUserName
	case mtproto.ClassUpdateUserPhoto:
		value = constants.UpdateTypeUpdateUserPhoto
	case mtproto.ClassUpdateNewEncryptedMessage:
		value = constants.UpdateTypeUpdateNewEncryptedMessage
	case mtproto.ClassUpdateEncryptedChatTyping:
		value = constants.UpdateTypeUpdateEncryptedChatTyping
	case mtproto.ClassUpdateEncryption:
		value = constants.UpdateTypeUpdateEncryption
	case mtproto.ClassUpdateEncryptedMessagesRead:
		value = constants.UpdateTypeUpdateEncryptedMessagesRead
	case mtproto.ClassUpdateChatParticipantAdd:
		value = constants.UpdateTypeUpdateChatParticipantAdd
	case mtproto.ClassUpdateChatParticipantDelete:
		value = constants.UpdateTypeUpdateChatParticipantDelete
	case mtproto.ClassUpdateDcOptions:
		value = constants.UpdateTypeUpdateDcOptions
	case mtproto.ClassUpdateNotifySettings:
		value = constants.UpdateTypeUpdateNotifySettings
	case mtproto.ClassUpdateServiceNotification:
		value = constants.UpdateTypeUpdateServiceNotification
	case mtproto.ClassUpdatePrivacy:
		value = constants.UpdateTypeUpdatePrivacy
	case mtproto.ClassUpdateUserPhone:
		value = constants.UpdateTypeUpdateUserPhone
	case mtproto.ClassUpdateReadHistoryInbox:
		value = constants.UpdateTypeUpdateReadHistoryInbox
	case mtproto.ClassUpdateReadHistoryOutbox:
		value = constants.UpdateTypeUpdateReadHistoryOutbox
	case mtproto.ClassUpdateWebPage:
		value = constants.UpdateTypeUpdateWebPage
	case mtproto.ClassUpdateReadMessagesContents:
		value = constants.UpdateTypeUpdateReadMessagesContents
	case mtproto.ClassUpdateChannelTooLong:
		value = constants.UpdateTypeUpdateChannelTooLong
	case mtproto.ClassUpdateChannel:
		value = constants.UpdateTypeUpdateChannel
	case mtproto.ClassUpdateNewChannelMessage:
		value = constants.UpdateTypeUpdateNewChannelMessage
	case mtproto.ClassUpdateReadChannelInbox:
		value = constants.UpdateTypeUpdateReadChannelInbox
	case mtproto.ClassUpdateDeleteChannelMessages:
		value = constants.UpdateTypeUpdateDeleteChannelMessages
	case mtproto.ClassUpdateChannelMessageViews:
		value = constants.UpdateTypeUpdateChannelMessageViews
	case mtproto.ClassUpdateChatParticipantAdmin:
		value = constants.UpdateTypeUpdateChatParticipantAdmin
	case mtproto.ClassUpdateNewStickerSet:
		value = constants.UpdateTypeUpdateNewStickerSet
	case mtproto.ClassUpdateStickerSetsOrder:
		value = constants.UpdateTypeUpdateStickerSetsOrder
	case mtproto.ClassUpdateStickerSets:
		value = constants.UpdateTypeUpdateStickerSets
	case mtproto.ClassUpdateSavedGifs:
		value = constants.UpdateTypeUpdateSavedGifs
	case mtproto.ClassUpdateBotInlineQuery:
		value = constants.UpdateTypeUpdateBotInlineQuery
	case mtproto.ClassUpdateBotInlineSend:
		value = constants.UpdateTypeUpdateBotInlineSend
	case mtproto.ClassUpdateEditChannelMessage:
		value = constants.UpdateTypeUpdateEditChannelMessage
	case mtproto.ClassUpdateBotCallbackQuery:
		value = constants.UpdateTypeUpdateBotCallbackQuery
	case mtproto.ClassUpdateEditMessage:
		value = constants.UpdateTypeUpdateEditMessage
	case mtproto.ClassUpdateInlineBotCallbackQuery:
		value = constants.UpdateTypeUpdateInlineBotCallbackQuery
	case mtproto.ClassUpdateReadChannelOutbox:
		value = constants.UpdateTypeUpdateReadChannelOutbox
	case mtproto.ClassUpdateDraftMessage:
		value = constants.UpdateTypeUpdateDraftMessage
	case mtproto.ClassUpdateReadFeaturedStickers:
		value = constants.UpdateTypeUpdateReadFeaturedStickers
	case mtproto.ClassUpdateRecentStickers:
		value = constants.UpdateTypeUpdateRecentStickers
	case mtproto.ClassUpdateConfig:
		value = constants.UpdateTypeUpdateConfig
	case mtproto.ClassUpdatePtsChanged:
		value = constants.UpdateTypeUpdatePtsChanged
	case mtproto.ClassUpdateChannelWebPage:
		value = constants.UpdateTypeUpdateChannelWebPage
	case mtproto.ClassUpdateDialogPinned:
		value = constants.UpdateTypeUpdateDialogPinned
	case mtproto.ClassUpdatePinnedDialogs:
		value = constants.UpdateTypeUpdatePinnedDialogs
	case mtproto.ClassUpdateBotWebhookJSON:
		value = constants.UpdateTypeUpdateBotWebhookJSON
	case mtproto.ClassUpdateBotWebhookJSONQuery:
		value = constants.UpdateTypeUpdateBotWebhookJSONQuery
	case mtproto.ClassUpdateBotShippingQuery:
		value = constants.UpdateTypeUpdateBotShippingQuery
	case mtproto.ClassUpdateBotPrecheckoutQuery:
		value = constants.UpdateTypeUpdateBotPrecheckoutQuery
	case mtproto.ClassUpdatePhoneCall:
		value = constants.UpdateTypeUpdatePhoneCall
	case mtproto.ClassUpdateLangPackTooLong:
		value = constants.UpdateTypeUpdateLangPackTooLong
	case mtproto.ClassUpdateLangPack:
		value = constants.UpdateTypeUpdateLangPack
	case mtproto.ClassUpdateFavedStickers:
		value = constants.UpdateTypeUpdateFavedStickers
	case mtproto.ClassUpdateChannelReadMessagesContents:
		value = constants.UpdateTypeUpdateChannelReadMessagesContents
	case mtproto.ClassUpdateContactsReset:
		value = constants.UpdateTypeUpdateContactsReset
	case mtproto.ClassUpdateChannelAvailableMessages:
		value = constants.UpdateTypeUpdateChannelAvailableMessages
	case mtproto.ClassUpdateDialogUnreadMark:
		value = constants.UpdateTypeUpdateDialogUnreadMark
	case mtproto.ClassUpdateMessagePoll:
		value = constants.UpdateTypeUpdateMessagePoll
	case mtproto.ClassUpdateChatDefaultBannedRights:
		value = constants.UpdateTypeUpdateChatDefaultBannedRights
	case mtproto.ClassUpdateFolderPeers:
		value = constants.UpdateTypeUpdateFolderPeers
	case mtproto.ClassUpdatePeerSettings:
		value = constants.UpdateTypeUpdatePeerSettings
	case mtproto.ClassUpdatePeerLocated:
		value = constants.UpdateTypeUpdatePeerLocated
	case mtproto.ClassUpdateNewScheduledMessage:
		value = constants.UpdateTypeUpdateNewScheduledMessage
	case mtproto.ClassUpdateDeleteScheduledMessages:
		value = constants.UpdateTypeUpdateDeleteScheduledMessages
	case mtproto.ClassUpdateTheme:
		value = constants.UpdateTypeUpdateTheme
	case mtproto.ClassUpdateGeoLiveViewed:
		value = constants.UpdateTypeUpdateGeoLiveViewed
	case mtproto.ClassUpdateLoginToken:
		value = constants.UpdateTypeUpdateLoginToken
	case mtproto.ClassUpdateMessagePollVote:
		value = constants.UpdateTypeUpdateMessagePollVote
	case mtproto.ClassUpdateDialogFilter:
		value = constants.UpdateTypeUpdateDialogFilter
	case mtproto.ClassUpdateDialogFilterOrder:
		value = constants.UpdateTypeUpdateDialogFilterOrder
	case mtproto.ClassUpdateDialogFilters:
		value = constants.UpdateTypeUpdateDialogFilters
	case mtproto.ClassUpdatePhoneCallSignalingData:
		value = constants.UpdateTypeUpdatePhoneCallSignalingData
	case mtproto.ClassUpdateChannelParticipant:
		value = constants.UpdateTypeUpdateChannelParticipant
	case mtproto.ClassUpdateChannelMessageForwards:
		value = constants.UpdateTypeUpdateChannelMessageForwards
	case mtproto.ClassUpdateReadChannelDiscussionInbox:
		value = constants.UpdateTypeUpdateReadChannelDiscussionInbox
	case mtproto.ClassUpdateReadChannelDiscussionOutbox:
		value = constants.UpdateTypeUpdateReadChannelDiscussionOutbox
	case mtproto.ClassUpdatePeerBlocked:
		value = constants.UpdateTypeUpdatePeerBlocked
	case mtproto.ClassUpdateChannelUserTyping:
		value = constants.UpdateTypeUpdateChannelUserTyping
	case mtproto.ClassUpdatePinnedMessages:
		value = constants.UpdateTypeUpdatePinnedMessages
	case mtproto.ClassUpdatePinnedChannelMessages:
		value = constants.UpdateTypeUpdatePinnedChannelMessages
	case mtproto.ClassUpdateChat:
		value = constants.UpdateTypeUpdateChat
	case mtproto.ClassUpdateGroupCallParticipants:
		value = constants.UpdateTypeUpdateGroupCallParticipants
	case mtproto.ClassUpdateGroupCall:
		value = constants.UpdateTypeUpdateGroupCall
	case mtproto.ClassUpdateContactRegistered:
		value = constants.UpdateTypeUpdateContactRegistered
	case mtproto.ClassUpdateContactLink:
		value = constants.UpdateTypeUpdateContactLink
	case mtproto.ClassUpdateUserBlocked:
		value = constants.UpdateTypeUpdateUserBlocked
	case mtproto.ClassUpdateChatAdmins:
		value = constants.UpdateTypeUpdateChatAdmins
	case mtproto.ClassUpdateChannelPinnedMessage:
		value = constants.UpdateTypeUpdateChannelPinnedMessage

	default:
	}

	return value
}

//
//func ToUpdateMessage(up *mtproto.Update) (*mtproto.UpdateMessage, bool) {
//	updateMessage := &mtproto.UpdateMessage{}
//	updateMessage.UpdateType = ToUpdateType(up).ToInt32()
//	updateMessage.Pts = up.Pts
//	updateMessage.PtsCount = up.PtsCount
//	updateMessage.Date = up.Date
//	pass := true
//	switch up.ClassName {
//	case mtproto.ClassUpdateNewMessage:
//		updateNewMessage := up.To_UpdateNewMessage()
//		buf, _ := jsoniter.Marshal(updateNewMessage.GetMessage1F2B0AFD122())
//		updateMessage.MessageId = []int32{updateNewMessage.GetMessage1F2B0AFD122().GetId()}
//		updateMessage.Message = string(buf)
//		updateMessage.Version = UpdateVersion1.ToInt32()
//	case mtproto.ClassUpdateDeleteMessages:
//		updateDeleteMessages := up.To_UpdateDeleteMessages()
//		updateMessage.MessageId = updateDeleteMessages.GetMessages()
//	case mtproto.ClassUpdateReadHistoryInbox,
//		mtproto.ClassUpdateReadHistoryOutbox,
//		mtproto.ClassUpdateReadMessagesContents,
//		mtproto.ClassUpdateEditMessage:
//		buf, _ := jsoniter.Marshal(up)
//		updateMessage.Message = string(buf)
//		updateMessage.Version = UpdateVersion1.ToInt32()
//
//	default:
//		pass = false
//	}
//
//	return updateMessage, pass
//}
//
//func ToUpdate(u *mtproto.UpdateMessage) *mtproto.Update {
//	value := &mtproto.Update{
//		Pts: u.Pts,
//		PtsCount: u.PtsCount,
//		Date: u.Date,
//		MessageEBE46819122: u.Message,
//		Messages: u.MessageId,
//	}
//
//	switch constants.UpdateTypeFromInt32(u.UpdateType) {
//	case constants.UpdateTypeUpdateNewMessage:
//		value = mtproto.NewTLUpdateNewMessage(value).To_Update()
//		if len(u.Message) == 0 {
//			value.Message1F2B0AFD122 = mtproto.NewTLMessageEmpty(nil).To_Message()
//		} else {
//			if u.Version == UpdateVersion1.ToInt32() {
//				value.Message1F2B0AFD122 = &mtproto.Message{}
//				err := jsoniter.Unmarshal([]byte(u.Message), value.Message1F2B0AFD122)
//				if err != nil {
//					log.Errorf("ToUpdate :%s", u.Message)
//				}
//			} else {
//				value.MessageEBE46819122 = u.Message
//				value.Message1F2B0AFD122 = mtproto.NewTLMessage(&mtproto.Message{
//					Message: u.Message,
//				}).To_Message()
//			}
//		}
//
//	case constants.UpdateTypeUpdateMessageID:
//		value = mtproto.NewTLUpdateMessageID(value).To_Update()
//	case constants.UpdateTypeUpdateDeleteMessages:
//		value = mtproto.NewTLUpdateDeleteMessages(value).To_Update()
//	case constants.UpdateTypeUpdateUserTyping:
//		value = mtproto.NewTLUpdateUserTyping(value).To_Update()
//	case constants.UpdateTypeUpdateChatUserTyping:
//		value = mtproto.NewTLUpdateChatUserTyping(value).To_Update()
//	case constants.UpdateTypeUpdateChatParticipants:
//		value = mtproto.NewTLUpdateChatParticipants(value).To_Update()
//	case constants.UpdateTypeUpdateUserStatus:
//		value = mtproto.NewTLUpdateUserStatus(value).To_Update()
//	case constants.UpdateTypeUpdateUserName:
//		value = mtproto.NewTLUpdateUserName(value).To_Update()
//	case constants.UpdateTypeUpdateUserPhoto:
//		value = mtproto.NewTLUpdateUserPhoto(value).To_Update()
//	case constants.UpdateTypeUpdateNewEncryptedMessage:
//		value = mtproto.NewTLUpdateNewEncryptedMessage(value).To_Update()
//	case constants.UpdateTypeUpdateEncryptedChatTyping:
//		value = mtproto.NewTLUpdateEncryptedChatTyping(value).To_Update()
//	case constants.UpdateTypeUpdateEncryption:
//		value = mtproto.NewTLUpdateEncryption(value).To_Update()
//	case constants.UpdateTypeUpdateEncryptedMessagesRead:
//		value = mtproto.NewTLUpdateEncryptedMessagesRead(value).To_Update()
//	case constants.UpdateTypeUpdateChatParticipantAdd:
//		value = mtproto.NewTLUpdateChatParticipantAdd(value).To_Update()
//	case constants.UpdateTypeUpdateChatParticipantDelete:
//		value = mtproto.NewTLUpdateChatParticipantDelete(value).To_Update()
//	case constants.UpdateTypeUpdateDcOptions:
//		value = mtproto.NewTLUpdateDcOptions(value).To_Update()
//	case constants.UpdateTypeUpdateNotifySettings:
//		value = mtproto.NewTLUpdateNotifySettings(value).To_Update()
//
//	case constants.UpdateTypeUpdateEditMessage:
//		value = mtproto.NewTLUpdateEditMessage(value).To_Update()
//		err := jsoniter.Unmarshal([]byte(u.Message), value)
//		if err != nil {
//			log.Errorf("UpdateTypeUpdateEditMessage ToUpdate :%s", u.Message)
//		}
//	case constants.UpdateTypeUpdateReadMessagesContents:
//		value = mtproto.NewTLUpdateReadMessagesContents(value).To_Update()
//		err := jsoniter.Unmarshal([]byte(u.Message), value)
//		if err != nil {
//			log.Errorf("UpdateTypeUpdateReadMessagesContents ToUpdate :%s", u.Message)
//		}
//	case constants.UpdateTypeUpdateReadHistoryOutbox:
//		value = mtproto.NewTLUpdateReadHistoryOutbox(value).To_Update()
//		err := jsoniter.Unmarshal([]byte(u.Message), value)
//		if err != nil {
//			log.Errorf("UpdateTypeUpdateReadHistoryOutbox ToUpdate :%s", u.Message)
//		}
//	case constants.UpdateTypeUpdateReadHistoryInbox:
//		value = mtproto.NewTLUpdateReadHistoryInbox(value).To_Update()
//		err := jsoniter.Unmarshal([]byte(u.Message), value)
//		if err != nil {
//			log.Errorf("UpdateTypeUpdateReadHistoryInbox ToUpdate :%s", u.Message)
//		}
//	//case mtproto.ClassUpdateServiceNotification:
//	//	value = constants.UpdateTypeUpdateServiceNotification
//	//case mtproto.ClassUpdatePrivacy:
//	//	value = constants.UpdateTypeUpdatePrivacy
//	//case mtproto.ClassUpdateUserPhone:
//	//	value = constants.UpdateTypeUpdateUserPhone
//
//	//case mtproto.ClassUpdateWebPage:
//	//	value = constants.UpdateTypeUpdateWebPage
//	//case mtproto.ClassUpdateChannelTooLong:
//	//	value = constants.UpdateTypeUpdateChannelTooLong
//	//case mtproto.ClassUpdateChannel:
//	//	value = constants.UpdateTypeUpdateChannel
//	//case mtproto.ClassUpdateNewChannelMessage:
//	//	value = constants.UpdateTypeUpdateNewChannelMessage
//	//case mtproto.ClassUpdateReadChannelInbox:
//	//	value = constants.UpdateTypeUpdateReadChannelInbox
//	//case mtproto.ClassUpdateDeleteChannelMessages:
//	//	value = constants.UpdateTypeUpdateDeleteChannelMessages
//	//case mtproto.ClassUpdateChannelMessageViews:
//	//	value = constants.UpdateTypeUpdateChannelMessageViews
//	//case mtproto.ClassUpdateChatParticipantAdmin:
//	//	value = constants.UpdateTypeUpdateChatParticipantAdmin
//	//case mtproto.ClassUpdateNewStickerSet:
//	//	value = constants.UpdateTypeUpdateNewStickerSet
//	//case mtproto.ClassUpdateStickerSetsOrder:
//	//	value = constants.UpdateTypeUpdateStickerSetsOrder
//	//case mtproto.ClassUpdateStickerSets:
//	//	value = constants.UpdateTypeUpdateStickerSets
//	//case mtproto.ClassUpdateSavedGifs:
//	//	value = constants.UpdateTypeUpdateSavedGifs
//	//case mtproto.ClassUpdateBotInlineQuery:
//	//	value = constants.UpdateTypeUpdateBotInlineQuery
//	//case mtproto.ClassUpdateBotInlineSend:
//	//	value = constants.UpdateTypeUpdateBotInlineSend
//	//case mtproto.ClassUpdateEditChannelMessage:
//	//	value = constants.UpdateTypeUpdateEditChannelMessage
//	//case mtproto.ClassUpdateBotCallbackQuery:
//	//	value = constants.UpdateTypeUpdateBotCallbackQuery
//	//case mtproto.ClassUpdateInlineBotCallbackQuery:
//	//	value = constants.UpdateTypeUpdateInlineBotCallbackQuery
//	//case mtproto.ClassUpdateReadChannelOutbox:
//	//	value = constants.UpdateTypeUpdateReadChannelOutbox
//	//case mtproto.ClassUpdateDraftMessage:
//	//	value = constants.UpdateTypeUpdateDraftMessage
//	//case mtproto.ClassUpdateReadFeaturedStickers:
//	//	value = constants.UpdateTypeUpdateReadFeaturedStickers
//	//case mtproto.ClassUpdateRecentStickers:
//	//	value = constants.UpdateTypeUpdateRecentStickers
//	//case mtproto.ClassUpdateConfig:
//	//	value = constants.UpdateTypeUpdateConfig
//	//case mtproto.ClassUpdatePtsChanged:
//	//	value = constants.UpdateTypeUpdatePtsChanged
//	//case mtproto.ClassUpdateChannelWebPage:
//	//	value = constants.UpdateTypeUpdateChannelWebPage
//	//case mtproto.ClassUpdateDialogPinned:
//	//	value = constants.UpdateTypeUpdateDialogPinned
//	//case mtproto.ClassUpdatePinnedDialogs:
//	//	value = constants.UpdateTypeUpdatePinnedDialogs
//	//case mtproto.ClassUpdateBotWebhookJSON:
//	//	value = constants.UpdateTypeUpdateBotWebhookJSON
//	//case mtproto.ClassUpdateBotWebhookJSONQuery:
//	//	value = constants.UpdateTypeUpdateBotWebhookJSONQuery
//	//case mtproto.ClassUpdateBotShippingQuery:
//	//	value = constants.UpdateTypeUpdateBotShippingQuery
//	//case mtproto.ClassUpdateBotPrecheckoutQuery:
//	//	value = constants.UpdateTypeUpdateBotPrecheckoutQuery
//	//case mtproto.ClassUpdatePhoneCall:
//	//	value = constants.UpdateTypeUpdatePhoneCall
//	//case mtproto.ClassUpdateLangPackTooLong:
//	//	value = constants.UpdateTypeUpdateLangPackTooLong
//	//case mtproto.ClassUpdateLangPack:
//	//	value = constants.UpdateTypeUpdateLangPack
//	//case mtproto.ClassUpdateFavedStickers:
//	//	value = constants.UpdateTypeUpdateFavedStickers
//	//case mtproto.ClassUpdateChannelReadMessagesContents:
//	//	value = constants.UpdateTypeUpdateChannelReadMessagesContents
//	//case mtproto.ClassUpdateContactsReset:
//	//	value = constants.UpdateTypeUpdateContactsReset
//	//case mtproto.ClassUpdateChannelAvailableMessages:
//	//	value = constants.UpdateTypeUpdateChannelAvailableMessages
//	//case mtproto.ClassUpdateDialogUnreadMark:
//	//	value = constants.UpdateTypeUpdateDialogUnreadMark
//	//case mtproto.ClassUpdateMessagePoll:
//	//	value = constants.UpdateTypeUpdateMessagePoll
//	//case mtproto.ClassUpdateChatDefaultBannedRights:
//	//	value = constants.UpdateTypeUpdateChatDefaultBannedRights
//	//case mtproto.ClassUpdateFolderPeers:
//	//	value = constants.UpdateTypeUpdateFolderPeers
//	//case mtproto.ClassUpdatePeerSettings:
//	//	value = constants.UpdateTypeUpdatePeerSettings
//	//case mtproto.ClassUpdatePeerLocated:
//	//	value = constants.UpdateTypeUpdatePeerLocated
//	//case mtproto.ClassUpdateNewScheduledMessage:
//	//	value = constants.UpdateTypeUpdateNewScheduledMessage
//	//case mtproto.ClassUpdateDeleteScheduledMessages:
//	//	value = constants.UpdateTypeUpdateDeleteScheduledMessages
//	//case mtproto.ClassUpdateTheme:
//	//	value = constants.UpdateTypeUpdateTheme
//	//case mtproto.ClassUpdateGeoLiveViewed:
//	//	value = constants.UpdateTypeUpdateGeoLiveViewed
//	//case mtproto.ClassUpdateLoginToken:
//	//	value = constants.UpdateTypeUpdateLoginToken
//	//case mtproto.ClassUpdateMessagePollVote:
//	//	value = constants.UpdateTypeUpdateMessagePollVote
//	//case mtproto.ClassUpdateDialogFilter:
//	//	value = constants.UpdateTypeUpdateDialogFilter
//	//case mtproto.ClassUpdateDialogFilterOrder:
//	//	value = constants.UpdateTypeUpdateDialogFilterOrder
//	//case mtproto.ClassUpdateDialogFilters:
//	//	value = constants.UpdateTypeUpdateDialogFilters
//	//case mtproto.ClassUpdatePhoneCallSignalingData:
//	//	value = constants.UpdateTypeUpdatePhoneCallSignalingData
//	//case mtproto.ClassUpdateChannelParticipant:
//	//	value = constants.UpdateTypeUpdateChannelParticipant
//	//case mtproto.ClassUpdateChannelMessageForwards:
//	//	value = constants.UpdateTypeUpdateChannelMessageForwards
//	//case mtproto.ClassUpdateReadChannelDiscussionInbox:
//	//	value = constants.UpdateTypeUpdateReadChannelDiscussionInbox
//	//case mtproto.ClassUpdateReadChannelDiscussionOutbox:
//	//	value = constants.UpdateTypeUpdateReadChannelDiscussionOutbox
//	//case mtproto.ClassUpdatePeerBlocked:
//	//	value = constants.UpdateTypeUpdatePeerBlocked
//	//case mtproto.ClassUpdateChannelUserTyping:
//	//	value = constants.UpdateTypeUpdateChannelUserTyping
//	//case mtproto.ClassUpdatePinnedMessages:
//	//	value = constants.UpdateTypeUpdatePinnedMessages
//	//case mtproto.ClassUpdatePinnedChannelMessages:
//	//	value = constants.UpdateTypeUpdatePinnedChannelMessages
//	//case mtproto.ClassUpdateChat:
//	//	value = constants.UpdateTypeUpdateChat
//	//case mtproto.ClassUpdateGroupCallParticipants:
//	//	value = constants.UpdateTypeUpdateGroupCallParticipants
//	//case mtproto.ClassUpdateGroupCall:
//	//	value = constants.UpdateTypeUpdateGroupCall
//	//case mtproto.ClassUpdateContactRegistered:
//	//	value = constants.UpdateTypeUpdateContactRegistered
//	//case mtproto.ClassUpdateContactLink:
//	//	value = constants.UpdateTypeUpdateContactLink
//	//case mtproto.ClassUpdateUserBlocked:
//	//	value = constants.UpdateTypeUpdateUserBlocked
//	//case mtproto.ClassUpdateChatAdmins:
//	//	value = constants.UpdateTypeUpdateChatAdmins
//	//case mtproto.ClassUpdateChannelPinnedMessage:
//	//	value = constants.UpdateTypeUpdateChannelPinnedMessage
//
//	default:
//	}
//
//	return value
//}
