/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : updates.getChannelDifference_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/util"
	"novachat_engine/service/account/setting"
	"novachat_engine/service/chat"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
	data_message "novachat_engine/service/data/messages/message"
	"novachat_engine/service/notify_setting"
	"time"
)

const (
	channelPollTimeout  = int32(30)
	channelMessageLimit = int32(1000)
)

//  updates.getChannelDifference#3173d78 flags:# force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;
//
func (s *UpdatesServiceImpl) UpdatesGetChannelDifference(ctx context.Context, request *mtproto.TLUpdatesGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("UpdatesGetChannelDifference %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	_ = request.Force
	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqGetParticipants(ctx, &chatService.GetParticipants{
		ChatId: chatId,
		UserId: md.UserId,
	})
	if err != nil {
		log.Errorf("UpdatesGetChannelDifference ReqGetParticipants %v error:%s", err.Error())
		return nil, err
	}

	var chatInfo chatService.Chat
	err = types.UnmarshalAny(resp, &chatInfo)
	if err != nil {
		log.Errorf("UpdatesGetChannelDifference ReqGetParticipants resp error:%s", err.Error())
		return nil, err
	}

	var participant *data_chat.ChatParticipant
	if idx := util.Index(chatInfo.ParticipantList, func(idx int) bool {
		return chatInfo.ParticipantList[idx].UserId == md.UserId
	}); idx >= 0 {
		participant = chatInfo.ParticipantList[idx]
	}

	if participant != nil && participant.State != data_chat.ParticipantState_ParticipantStateNormal {
		return mtproto.NewTLUpdatesChannelDifferenceEmpty(nil).To_Updates_ChannelDifference(), nil
	}

	conversation, err := s.messageCore.GetConversation(md.UserId, &data_message.Conversation{
		UserId:   md.UserId,
		PeerId:   chatId,
		PeerType: constants.PeerTypeChannel.ToInt32(),
	})
	if err != nil {
		log.Errorf("UpdatesGetChannelDifference GetConversation %v error:%s", request, err.Error())
		return nil, err
	}
	//
	//if request.GetPts() > conversation.Pts {
	//	return mtproto.NewTLUpdatesChannelDifferenceEmpty(&mtproto.Updates_ChannelDifference{
	//		Final:   false,
	//		Pts:     conversation.Pts,
	//		Timeout: channelPollTimeout,
	//	}).To_Updates_ChannelDifference(), nil
	//}
	if conversation.Pts > request.GetPts()+channelMessageLimit {
		notifySetting, err1 := s.accountSettingCore.GetNotifySetting(&setting.NotifyPeerSettingType{
			UserId:         md.UserId,
			PeerId:         chatId,
			PeerType:       constants.PeerTypeChat,
			PeerNotifyType: constants.PeerGlobalSetting,
		})
		if err1 != nil {
			log.Errorf("UpdatesGetChannelDifference GetNotifySetting: error:%s", err.Error())
			return nil, err
		}

		//  dialog#2c171f72 flags:# pinned:flags.2?true unread_mark:flags.3?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage folder_id:flags.4?int = Dialog;
		//  dialog#e4def5db flags:# pinned:flags.2?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage = Dialog;
		dialog := mtproto.NewTLDialog(&mtproto.Dialog{
			Pinned:              conversation.Pinned,
			UnreadMark:          false,
			Peer:                mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChannelIDType(chatId).ToInt32()}).To_Peer(),
			TopMessage:          conversation.Top,
			ReadInboxMaxId:      conversation.InboxMaxId,
			ReadOutboxMaxId:     conversation.OutboxMaxId,
			UnreadCount:         conversation.UnreadCount,
			UnreadMentionsCount: 0,
			NotifySettings:      notify_setting.ToPeerNotifySettings(notifySetting, md.Layer),
			Pts:                 conversation.Pts,
			FolderId:            conversation.FolderId,
		}).To_Dialog()

		if len(conversation.Draft) > 0 {
			dialog.Draft = mtproto.NewTLDraftMessageEmpty(&mtproto.DraftMessage{
				Date: int32(time.Now().Unix()),
			}).To_DraftMessage()
		} else {
			//  draftMessage#fd8e711f flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int message:string entities:flags.3?Vector<MessageEntity> date:int = DraftMessage;
			dialog.Draft = mtproto.NewTLDraftMessage(&mtproto.DraftMessage{
				Message: conversation.Draft,
				Date:    int32(time.Now().Unix()),
			}).To_DraftMessage()
		}

		//  updates.channelDifferenceTooLong#a4bcc6fe flags:# final:flags.0?true timeout:flags.1?int dialog:Dialog messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
		//  updates.channelDifferenceTooLong#6a9d7b35 flags:# final:flags.0?true pts:int timeout:flags.1?int top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
		//
		updatesChannelDifferenceTooLong := mtproto.NewTLUpdatesChannelDifferenceTooLong(&mtproto.Updates_ChannelDifference{
			Final:               false,
			Timeout:             channelPollTimeout,
			Pts:                 conversation.Pts,
			Dialog:              dialog,
			Messages:            []*mtproto.Message{},
			Chats:               []*mtproto.Chat{chat.ToChat(md.UserId, &chatInfo, md.Layer)},
			Users:               []*mtproto.User{},
			TopMessage:          conversation.Top,
			ReadInboxMaxId:      conversation.InboxMaxId,
			ReadOutboxMaxId:     conversation.OutboxMaxId,
			UnreadCount:         conversation.UnreadCount,
			UnreadMentionsCount: 0,
		})
		return updatesChannelDifferenceTooLong.To_Updates_ChannelDifference(), nil
	}

	updatesGetChannelDifference, err := s.accountUpdateCore.GetUpdateChannelDifference(md.UserId, chatId, request.Force, request.Pts, request.Limit, md.Layer)
	if err != nil {
		log.Errorf("ChannelsGetChannels %v, request: %v GetChats error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	updatesGetChannelDifference.Chats = []*mtproto.Chat{chat.ToChat(md.UserId, &chatInfo, md.Layer)}
	updatesGetChannelDifference.Timeout = channelPollTimeout

	for idx := range updatesGetChannelDifference.NewMessages {
		updatesGetChannelDifference.NewMessages[idx] = compat.MessageCompat(updatesGetChannelDifference.NewMessages[idx], md.Layer)
	}

	for idx := range updatesGetChannelDifference.OtherUpdates {
		updatesGetChannelDifference.OtherUpdates[idx] = compat.UpdateCompat(updatesGetChannelDifference.OtherUpdates[idx], md.Layer)
	}

	for idx := range updatesGetChannelDifference.Users {
		updatesGetChannelDifference.Users[idx] = compat.CompatUser(updatesGetChannelDifference.Users[idx], md.Layer)
	}

	for idx := range updatesGetChannelDifference.Chats {
		updatesGetChannelDifference.Chats[idx] = compat.CompatChat(updatesGetChannelDifference.Chats[idx], md.Layer)
	}

	for idx := range updatesGetChannelDifference.Messages {
		updatesGetChannelDifference.Messages[idx] = compat.MessageCompat(updatesGetChannelDifference.Messages[idx], md.Layer)
	}
	return updatesGetChannelDifference, nil
}
