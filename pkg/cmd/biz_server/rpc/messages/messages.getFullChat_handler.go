/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getFullChat_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"math"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/util"
	"novachat_engine/service/account/setting"
	"novachat_engine/service/app/dc"
	"novachat_engine/service/banned_right"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
	data_fs "novachat_engine/service/data/fs"
	data_message "novachat_engine/service/data/messages/message"
	"novachat_engine/service/notify_setting"
	"novachat_engine/service/upload/photo"
)

//  messages.getFullChat#3b831c66 chat_id:int = messages.ChatFull;
//
func (s *MessagesServiceImpl) MessagesGetFullChat(ctx context.Context, request *mtproto.TLMessagesGetFullChat) (*mtproto.Messages_ChatFull, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetFullChat %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChatIDType32(request.ChatId).ToInt()
	dataChat, err := chatService.GetChatClientByKeyId(chatId).ReqFullChat(ctx, &chatService.FullChat{ChatId: chatId})
	if err != nil {
		log.Errorf("MessagesGetFullChat %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	settingV, err := s.accountSettingCore.GetNotifySetting(&setting.NotifyPeerSettingType{
		UserId:         md.UserId,
		PeerId:         chatId,
		PeerType:       constants.PeerTypeChannel,
		PeerNotifyType: constants.PeerNotifyInputNotifyPeer,
	})
	if err != nil {
		log.Warnf("MessagesGetFullChat request: %v GetNotifySetting error:%s", request, err.Error())
	}

	chatData := dataChat.GetChatData()

	fullChat := mtproto.NewTLChannelFull(&mtproto.ChatFull{
		CanSetUsername:      false,
		HasScheduled:        false,
		CanSetStickers:      false,
		CanViewParticipants: true,
		CanSetLocation:      false,
		CanViewStats:        false,
		HiddenPrehistory:    chatData.HiddenHistory,
		Id:                  request.ChatId,
		About:               chatData.About,
		Participants:        nil,
		ChatPhoto:           nil,
		NotifySettings:      notify_setting.ToPeerNotifySettings(settingV, md.Layer),
		ExportedInvite:      nil,
		BotInfo:             nil,
		PinnedMsgId:         0,
		FolderId:            0,
		ParticipantsCount:   dataChat.Count,
		AdminsCount:         0,
		Call:                nil,
	}).To_ChatFull()

	var chat *mtproto.Chat
	if chatData.Deleted {
		chat = mtproto.NewTLChannelForbidden(&mtproto.Chat{
			Title:      chatData.Title,
			Id:         request.ChatId,
			Broadcast:  false,
			Megagroup:  true,
			AccessHash: 0,
			UntilDate:  math.MaxInt32,
		}).To_Chat()
	} else {
		conversation, err1 := s.messageCore.GetConversationByPeerList(md.UserId, []*data_message.Conversation{{PeerId: chatId, PeerType: constants.PeerTypeChannel.ToInt32()}})
		if err1 != nil {
			log.Errorf("MessagesGetFullChat %v, request: %v GetConversationByPeerList error:%s", metadata.RpcMetaDataDebug(md), request, err1.Error())
			return nil, err1
		}
		if len(conversation) > 0 {
			fullChat.ReadInboxMaxId = conversation[0].InboxMaxId
			fullChat.ReadOutboxMaxId = conversation[0].OutboxMaxId
			fullChat.UnreadCount = conversation[0].UnreadCount
			fullChat.Pts = conversation[0].Pts
			if len(conversation[0].PinIds) > 0 {
				fullChat.PinnedMsgId = conversation[0].PinIds[0].Id
			}
			fullChat.FolderId = conversation[0].FolderId
		}

		chat = mtproto.NewTLChannel(&mtproto.Chat{
			Creator:      dataChat.ChatData.Creator == md.UserId,
			Kicked:       false,
			Megagroup:    true,
			Title:        dataChat.ChatData.Title,
			Left:         false,
			Deactivated:  false,
			CallActive:   false, //TODO:(Coderxw)
			CallNotEmpty: false, //TODO:(Coderxw)
			Id:           request.ChatId,
			Photo:        nil,
			//ParticipantsCount:     int32(len(dataChat.ParticipantList)),
			Date:                  dataChat.ChatData.Date,
			Version:               0, //TODO:(Coderxw)
			MigratedTo:            nil,
			AdminRights4DF3083493: nil,
			DefaultBannedRights:   nil,
		}).To_Chat()

		if dataChat.ChatData.Creator != md.UserId {
			participantList := dataChat.GetParticipantList()

			var idx int32
			util.Foreach(participantList, func(i interface{}, value interface{}) {
				iValue := i.(int)
				switch participantList[iValue].State {
				case data_chat.ParticipantState_ParticipantStateBan:
					fullChat.BannedCount++
				case data_chat.ParticipantState_ParticipantStateLeft:
					fullChat.KickedCount++
				case data_chat.ParticipantState_ParticipantStateDelete:
					fullChat.KickedCount++
				default:
					if participantList[iValue].Admin {
						fullChat.AdminsCount++
					}
				}

				if participantList[iValue].UserId == md.UserId {
					idx = int32(iValue)
				}
			})
			if idx >= 0 {
				switch participantList[idx].State {
				case data_chat.ParticipantState_ParticipantStateBan:
					fallthrough
				case data_chat.ParticipantState_ParticipantStateLeft:
					fullChat.Blocked = true
					fallthrough
				case data_chat.ParticipantState_ParticipantStateDelete:
					chat.Left = true
				default:
					if participantList[idx].Admin {
						chat.AdminRights4DF3083493 = banned_right.RightsToChatAdminBannedRight(participantList[idx].AdminRights)
						fullChat.CanSetUsername = true
					} else {
						chat.BannedRights4DF3083493 = banned_right.RightsToChatBannedRight(participantList[idx].Rights, 0)
					}
				}
			} else {
				fullChat.CanViewParticipants = false
			}
		} else {
			fullChat.CanSetUsername = true
			fullChat.CanViewStats = true
		}
		chat.DefaultBannedRights = banned_right.RightsToChatBannedRight(chatData.BannedRights, chatData.RightsUtilDate)
		if len(chatData.Username) == 0 {
			fullChat.ExportedInvite = mtproto.NewTLChatInviteEmpty(nil).To_ExportedChatInvite()
		} else {
			fullChat.ExportedInvite = mtproto.NewTLChatInviteExported(&mtproto.ExportedChatInvite{
				Link: chatData.Username,
			}).To_ExportedChatInvite()
			chat.Username = chatData.Username
		}
		if dataChat.GetChatData().GeoPoint != nil {
			chat.HasGeo = true
			fullChat.Location = mtproto.NewTLChannelLocation(&mtproto.ChannelLocation{
				GeoPoint: mtproto.NewTLGeoPoint(&mtproto.GeoPoint{
					Long:           dataChat.GetChatData().GeoPoint.Long,
					Lat:            dataChat.GetChatData().GeoPoint.Lat,
					AccessHash:     dataChat.GetChatData().GeoPoint.AccessHash,
					AccuracyRadius: dataChat.GetChatData().GeoPoint.AccuracyRadius,
				}).To_GeoPoint(),
				Address: dataChat.GetChatData().Address,
			}).To_ChannelLocation()
		} else {
			fullChat.Location = mtproto.NewTLChannelLocationEmpty(nil).To_ChannelLocation()
		}

		if chatData.Photo == nil || chatData.Photo.Photo == nil {
			chat.Photo = mtproto.NewTLChatPhotoEmpty(nil).To_ChatPhoto()
			fullChat.ChatPhoto = mtproto.NewTLPhotoEmpty(&mtproto.Photo{Id: 0}).To_Photo()
		} else {
			chat.Photo = mtproto.NewTLChatPhoto(&mtproto.ChatPhoto{
				PhotoSmall: nil,
				PhotoBig:   nil,
				DcId:       dc.DefaultDc,
				HasVideo:   false,
			}).To_ChatPhoto()

			chat.GetPhoto().PhotoSmall = compat.NewFileLocationByLayer(
				chatData.Photo.Photo.VolumeId,
				chatData.Photo.Photo.LocalId,
				md.Layer,
				dc.DefaultDc)
			if chatData.Photo.Video != nil {
				chat.GetPhoto().PhotoBig = compat.NewFileLocationByLayer(
					chatData.Photo.Video.VolumeId,
					chatData.Photo.Video.LocalId,
					md.Layer,
					dc.DefaultDc)
			}

			fullChat.ChatPhoto = photo.PhotoData2Photo(&data_fs.Photo{
				VolumeId: chatData.Photo.Photo.VolumeId,
				LocalId:  chatData.Photo.Photo.LocalId,
				FilePath: chatData.Photo.Photo.FilePath,
				Detail:   []*data_fs.PhotoDetail{chatData.Photo.Photo},
			})
			chat.GetPhoto().HasVideo = len(fullChat.GetChatPhoto().VideoSizes) > 0
		}
	}

	log.Infof("MessagesGetFullChat %v, request: %v fullChat:%+v chat:%+v", metadata.RpcMetaDataDebug(md), request, fullChat, chat)
	return mtproto.NewTLMessagesChatFull(&mtproto.Messages_ChatFull{
		FullChat: fullChat,
		Chats:    []*mtproto.Chat{chat},
	}).To_Messages_ChatFull(), nil
}
