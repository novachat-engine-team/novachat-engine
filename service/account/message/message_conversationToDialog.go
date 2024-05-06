/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package message

import (
	"context"
	"math"
	"novachat_engine/mtproto"
	chatClient "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	"novachat_engine/service/account/setting"
	"novachat_engine/service/chat"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message/message"
	data_message "novachat_engine/service/data/messages/message"
	"novachat_engine/service/input"
	"novachat_engine/service/notify_setting"
)

func (c *Core) conversationToDialog(userId int64, conversationList []*data_message.Conversation, foldId int32, layer int32) (*mtproto.Messages_Dialogs, error) {
	var peerType constants.PeerType
	chatIdList := make([]int64, 0, len(conversationList))
	userIdList := make([]int64, 0, len(conversationList))
	channelMsgIdList := make([]*message.ChannelMessageId, 0, len(conversationList))
	messageIdList := make([]int32, 0, len(conversationList))
	dialogs := make([]*mtproto.Dialog, 0, len(conversationList))
	notifySettingList := make([]*setting.NotifyPeerSettingType, 0, len(conversationList))
	for _, conversation := range conversationList {

		d := mtproto.NewTLDialog(nil).To_Dialog()
		inputPeerValue := input.MakeInputPeerValue(conversation.PeerId, peerType)
		d.UnreadMentionsCount = 0
		d.ReadOutboxMaxId = conversation.OutboxMaxId
		d.ReadInboxMaxId = conversation.InboxMaxId
		d.Pinned = conversation.Pinned
		d.UnreadMark = false
		d.Peer = inputPeerValue.ToPeer()
		d.TopMessage = conversation.Top
		d.UnreadCount = conversation.UnreadCount
		d.FolderId = foldId

		if userId == conversation.PeerId && conversation.PeerType == constants.PeerTypeUser.ToInt32() {
			d.ReadInboxMaxId = d.ReadOutboxMaxId
			d.UnreadCount = 0
		}

		if len(conversation.Draft) == 0 {
			d.Draft = mtproto.NewTLDraftMessageEmpty(nil).To_DraftMessage()
		} else {
			d.Draft = mtproto.NewTLDraftMessage(&mtproto.DraftMessage{
				Message: conversation.Draft,
			}).To_DraftMessage()
		}

		peerType = constants.PeerType(conversation.PeerType)
		switch peerType {
		case constants.PeerTypeSelf:
			messageIdList = append(messageIdList, conversation.Top)
		case constants.PeerTypeUser:
			userIdList = append(userIdList, conversation.PeerId)
			messageIdList = append(messageIdList, conversation.Top)
		case constants.PeerTypeChat, constants.PeerTypeChannel:
			chatIdList = append(chatIdList, conversation.PeerId)
			channelMsgIdList = append(channelMsgIdList, &message.ChannelMessageId{ChannelId: conversation.PeerId, IdList: []int32{conversation.Top}})
			d.Pts = conversation.Pts
			if conversation.Deleted {
				d.ReadInboxMaxId = 0
				d.ReadOutboxMaxId = 0
				d.UnreadCount = 0
				d.TopMessage = 0
				d.Draft = mtproto.NewTLDraftMessageEmpty(nil).To_DraftMessage()
			}
		}
		notifySettingList = append(notifySettingList, &setting.NotifyPeerSettingType{
			UserId:         userId,
			PeerId:         conversation.PeerId,
			PeerType:       peerType,
			PeerNotifyType: constants.PeerNotifyInputNotifyPeer,
		})

		if foldId == 0 || layer < 99 {
			dialogs = append(dialogs, d)
		} else {
			d.Folder = mtproto.NewTLFolder(&mtproto.Folder{
				AutofillNewBroadcasts:     false,
				AutofillPublicGroups:      false,
				AutofillNewCorrespondents: false,
				Id:                        foldId,
				Title:                     "",
				Photo:                     nil,
			}).To_Folder()
			d.FolderId = foldId
			dialogs = append(dialogs, d.To_DialogFolder().To_Dialog())
		}
	}

	var messageList []*mtproto.Message
	var err error
	if len(messageIdList) > 0 {
		messageList, err = c.GetMessages(userId, messageIdList, false)
		if err != nil {
			log.Warnf("conversationToDialog userId:%d peerType:%v GetMessageList error:%s", userId, peerType, err.Error())
		}
	}
	if len(notifySettingList) > 0 {
		settingList, err1 := c.settingCore.GetNotifySettingList(notifySettingList)
		if err1 != nil {
			log.Warnf("conversationToDialog userId:%d peerType:%v GetNotifySettingList error:%s", userId, peerType, err.Error())
		} else {
			for _, v := range dialogs {
				inputPeer := input.MakePeer(v.Peer)
				if settingList == nil {
					v.NotifySettings = mtproto.NewPeerNotifySettingEmptyLayer(layer)
					continue
				}

				if idx := util.Index(settingList, func(idx int) bool {
					return settingList[idx].PeerId == inputPeer.GetPeerId() &&
						constants.PeerTypeFromInt32(settingList[idx].PeerType) == inputPeer.GetPeerType()
				}); idx >= 0 {
					v.NotifySettings = notify_setting.ToPeerNotifySettings(settingList[idx].Setting, layer)
				} else {
					v.NotifySettings = mtproto.NewPeerNotifySettingEmptyLayer(layer)
				}
			}
		}
	}
	if len(channelMsgIdList) > 0 {
		for _, v := range channelMsgIdList {
			channelMessageList, err1 := c.GetChannelMessageList(v)
			if err1 != nil {
				log.Warnf("conversationToDialog - GetChannelMessageList error:%s", err.Error())
			} else {
				messageList = append(messageList, channelMessageList...)
			}
		}
	}

	userList, err := c.accountUsersCore.GetUserList(userId, userIdList)
	if err != nil {
		log.Warnf("conversationToDialog - GetUserList error:%s", err.Error())
	}

	var chats []*mtproto.Chat
	if len(chatIdList) > 0 {
		//TODO:(Coderxw)
		chatList, err1 := chatClient.GetChatClientByKeyId(userId).ReqAllChat(context.Background(), &chatClient.AllChat{
			Ids:    chatIdList,
			UserId: userId,
		})
		if err1 != nil {
			log.Warnf("conversationToDialog - ReqAllChat error:%s", err.Error())
		} else {
			var idx int
			for _, chatInfo := range chatList.Values {
				if idx = util.Index(conversationList, func(idx int) bool {
					return conversationList[idx].PeerId == chatInfo.ChatData.ChatId && conversationList[idx].PeerType == constants.PeerTypeChannel.ToInt32()
				}); idx < 0 {
					continue
				}
				if conversationList[idx].Deleted {
					chats = append(chats, mtproto.NewTLChannelForbidden(&mtproto.Chat{
						Broadcast:  chatInfo.ChatData.Broadcast,
						Megagroup:  true,
						Id:         constants.PeerTypeFromChannelIDType(chatInfo.ChatData.ChatId).ToInt32(),
						AccessHash: chatInfo.ChatData.AccessHash,
						Title:      chatInfo.ChatData.Title,
						UntilDate:  math.MaxInt32,
					}).To_Chat())
					continue
				}

				chats = append(chats, chat.ToChat(userId, chatInfo, layer))
			}
		}
	}

	return mtproto.NewTLMessagesDialogs(&mtproto.Messages_Dialogs{
		Dialogs:  dialogs,
		Messages: messageList,
		Chats:    chats,
		Users:    userList,
		Count:    int32(len(dialogs)),
	}).To_Messages_Dialogs(), nil
}
