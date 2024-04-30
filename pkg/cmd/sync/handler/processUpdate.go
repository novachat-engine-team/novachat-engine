package handler

import (
	"errors"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	accountUpdate "novachat_engine/service/account/update"
	"novachat_engine/service/constants"
	data_update "novachat_engine/service/data/update"
	"novachat_engine/service/update"
)

//  + TL_updateNewMessage
//  + TL_updateMessageID
//  + TL_updateDeleteMessages
//  + TL_updateUserTyping
//  + TL_updateChatUserTyping
//  + TL_updateChatParticipants
//  + TL_updateUserStatus
//  + TL_updateUserName
//  + TL_updateUserPhoto
//  + TL_updateNewEncryptedMessage
//  + TL_updateEncryptedChatTyping
//  + TL_updateEncryption
//  + TL_updateEncryptedMessagesRead
//  + TL_updateChatParticipantAdd
//  + TL_updateChatParticipantDelete
//  + TL_updateDcOptions
//  + TL_updateNotifySettings
//  + TL_updateServiceNotification
//  + TL_updatePrivacy
//  + TL_updateUserPhone
//  + TL_updateReadHistoryInbox
//  + TL_updateReadHistoryOutbox
//  + TL_updateWebPage
//  + TL_updateReadMessagesContents
//  + TL_updateChannelTooLong
//  + TL_updateChannel
//  + TL_updateNewChannelMessage
//  + TL_updateReadChannelInbox
//  + TL_updateDeleteChannelMessages
//  + TL_updateChannelMessageViews
//  + TL_updateChatParticipantAdmin
//  + TL_updateNewStickerSet
//  + TL_updateStickerSetsOrder
//  + TL_updateStickerSets
//  + TL_updateSavedGifs
//  + TL_updateBotInlineQuery
//  + TL_updateBotInlineSend
//  + TL_updateEditChannelMessage
//  + TL_updateBotCallbackQuery
//  + TL_updateEditMessage
//  + TL_updateInlineBotCallbackQuery
//  + TL_updateReadChannelOutbox
//  + TL_updateDraftMessage
//  + TL_updateReadFeaturedStickers
//  + TL_updateRecentStickers
//  + TL_updateConfig
//  + TL_updatePtsChanged
//  + TL_updateChannelWebPage
//  + TL_updateDialogPinned
//  + TL_updatePinnedDialogs
//  + TL_updateBotWebhookJSON
//  + TL_updateBotWebhookJSONQuery
//  + TL_updateBotShippingQuery
//  + TL_updateBotPrecheckoutQuery
//  + TL_updatePhoneCall
//  + TL_updateLangPackTooLong
//  + TL_updateLangPack
//  + TL_updateFavedStickers
//  + TL_updateChannelReadMessagesContents
//  + TL_updateContactsReset
//  + TL_updateChannelAvailableMessages
//  + TL_updateDialogUnreadMark
//  + TL_updateMessagePoll
//  + TL_updateChatDefaultBannedRights
//  + TL_updateFolderPeers
//  + TL_updatePeerSettings
//  + TL_updatePeerLocated
//  + TL_updateNewScheduledMessage
//  + TL_updateDeleteScheduledMessages
//  + TL_updateTheme
//  + TL_updateGeoLiveViewed
//  + TL_updateLoginToken
//  + TL_updateMessagePollVote
//  + TL_updateDialogFilter
//  + TL_updateDialogFilterOrder
//  + TL_updateDialogFilters
//  + TL_updatePhoneCallSignalingData
//  + TL_updateChannelParticipant
//  + TL_updateChannelMessageForwards
//  + TL_updateReadChannelDiscussionInbox
//  + TL_updateReadChannelDiscussionOutbox
//  + TL_updatePeerBlocked
//  + TL_updateChannelUserTyping
//  + TL_updatePinnedMessages
//  + TL_updatePinnedChannelMessages
//  + TL_updateChat
//  + TL_updateGroupCallParticipants
//  + TL_updateGroupCall
//  + TL_updateContactRegistered
//  + TL_updateContactLink
//  + TL_updateUserBlocked
//  + TL_updateChatAdmins
//  + TL_updateChannelPinnedMessage

var nothingDoErr = errors.New("nothing do")

func (m *Handler) processUpdates(userId int64, updates *mtproto.Updates) error {
	var (
		err      error
		date     int32
		pts      int32
		ptsCount int32
		seqStart int32
		seq      int32
	)
	var dataUpdateList []*data_update.UserUpdate
	updateType := update.UpdatesToUpdateType(updates)
	switch updateType {
	case constants.UpdateTypeUpdatesTooLong:
		updates.To_UpdatesTooLong()
	case constants.UpdateTypeUpdateShortMessage:
		//  updateShortMessage#914fbf11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
		//  updateShortMessage#2296d2c8 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to:flags.3?MessageReplyHeader entities:flags.7?Vector<MessageEntity> = Updates;
		updateShortMessage := updates.To_UpdateShortMessage()
		date, pts, ptsCount = updateShortMessage.GetDate(), updateShortMessage.GetPts(), updateShortMessage.GetPtsCount()
	case constants.UpdateTypeUpdateShortChatMessage:
		//  updateShortChatMessage#16812688 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
		//  updateShortChatMessage#402d5dbb flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to:flags.3?MessageReplyHeader entities:flags.7?Vector<MessageEntity> = Updates;
		updateShortChatMessage := updates.To_UpdateShortChatMessage()
		date, pts, ptsCount = updateShortChatMessage.GetDate(), updateShortChatMessage.GetPts(), updateShortChatMessage.GetPtsCount()
	case constants.UpdateTypeUpdateShort:
		//  updateShort#78d4dec1 update:Update date:int = Updates;
		updateShort := updates.To_UpdateShort()
		dataUpdateList = m.processUpdateList(userId, []*mtproto.Update{updateShort.GetUpdate()})
		date = updateShort.GetDate()
	case constants.UpdateTypeUpdatesCombined:
		//  updatesCombined#725b04c3 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq_start:int seq:int = Updates;
		updatesCombined := updates.To_UpdatesCombined()
		date, seqStart, seq = updatesCombined.GetDate(), updatesCombined.GetSeqStart(), updatesCombined.GetSeq()
	case constants.UpdateTypeUpdateShortSentMessage:
		//  updateShortSentMessage#11f1331c flags:# out:flags.1?true id:int pts:int pts_count:int date:int media:flags.9?MessageMedia entities:flags.7?Vector<MessageEntity> = Updates;
		updateShortSentMessage := updates.To_UpdateShortSentMessage()
		date, pts, ptsCount = updateShortSentMessage.GetDate(), updateShortSentMessage.GetPts(), updateShortSentMessage.GetPtsCount()
	case constants.UpdateTypeUpdates:
		dataUpdateList = m.processUpdateList(userId, updates.To_Updates().GetUpdates())
		seq = updates.GetSeq()
		date = updates.GetDate()
	default:
		err = fmt.Errorf("processUpdates not updateType:%v", updateType)
		log.Warnf(err.Error())
		return nothingDoErr
	}

	if len(dataUpdateList) > 0 {
		err = m.accountUpdateCore.SaveUpdateDataList(userId, dataUpdateList)
	}
	_ = dataUpdateList
	_, _, _, _, _ = date, pts, ptsCount, seqStart, seq
	if err != nil {
		log.Errorf("processUpdates updateType:%v error:%s", updateType, err.Error())
	}
	return err
}

func (m *Handler) processUpdateList(userId int64, updateList []*mtproto.Update) []*data_update.UserUpdate {

	var peer *mtproto.Peer
	dataUpdateList := make([]*data_update.UserUpdate, 0, len(updateList))
	for _, up := range updateList {
		if up.PtsCount == 0 && up.Pts == 0 {
			if up.ChatId != 0 || up.ChannelId != 0 {
			} else {
				continue
			}
		}

		if up.PeerId != nil {
			peer = up.PeerId
		} else if up.Peer9961FD5C71 != nil {
			peer = up.Peer9961FD5C71
		} else {
			if up.Message1F2B0AFD71 != nil {
				if up.Message1F2B0AFD71.PeerId != nil {
					peer = up.Message1F2B0AFD71.PeerId
				} else if up.Message1F2B0AFD71.ToId != nil {
					peer = up.Message1F2B0AFD71.ToId
				}
			}
		}
		if peer == nil {
			log.Warnf("processUpdateList peerId = 0 update:%s", up.ClassName)
			continue
		}

		syncUpdate := accountUpdate.UpdateToUserUpdate(userId, up)

		switch peer.ClassName {
		case mtproto.ClassPeerUser:
			syncUpdate.PeerId = constants.PeerTypeFromUserIDType32(peer.UserId).ToInt()
			syncUpdate.PeerType = constants.PeerTypeUser.ToInt32()
		case mtproto.ClassPeerChat:
			syncUpdate.PeerId = constants.PeerTypeFromChatIDType32(peer.ChatId).ToInt()
			syncUpdate.PeerType = constants.PeerTypeChat.ToInt32()
		case mtproto.ClassPeerChannel:
			syncUpdate.PeerId = constants.PeerTypeFromChannelIDType32(peer.ChannelId).ToInt()
			syncUpdate.PeerType = constants.PeerTypeChannel.ToInt32()
		default:
			panic(fmt.Sprintf("processUpdateList peerId:%v", peer))
		}

		syncUpdate.UpdateKey = m.updateKeyNode.Generate().String()
		dataUpdateList = append(dataUpdateList, syncUpdate)
	}

	return dataUpdateList
}
