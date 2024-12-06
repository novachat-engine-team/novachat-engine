/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/1 9:46
 * @File : user_updates_util.go
 */

package update

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/util"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message/message"
	data_update "novachat_engine/service/data/update"
	"novachat_engine/service/input"
	"novachat_engine/service/update"
	"time"
)

func UserUpdateToUpdate(userUpdate *data_update.UserUpdate, layer int32) *mtproto.Update {
	up := &mtproto.Update{}
	t := constants.UpdateTypeFromInt32(userUpdate.UpdateType)
	switch t {
	case constants.UpdateTypeUpdateNewMessage, constants.UpdateTypeUpdateEditMessage,
		constants.UpdateTypeUpdateNewChannelMessage, constants.UpdateTypeUpdateEditChannelMessage:
		// updateEditMessage#e40370a3 message:Message pts:int pts_count:int = Update;
		// updateNewMessage#1f2b0afd message:Message pts:int pts_count:int = Update;
		// updateNewChannelMessage#62ba04d9 message:Message pts:int pts_count:int = Update;
		// updateEditChannelMessage#1b3f4df7 message:Message pts:int pts_count:int = Update;

		up.Pts = userUpdate.Pts
		up.PtsCount = userUpdate.PtsCount
		up.Date = userUpdate.Date
		up.RandomId = userUpdate.RandomId

		up.Message1F2B0AFD71 = message.ToMessage(userUpdate.MessageData, layer)
		up.MessageEBE4681971 = up.Message1F2B0AFD71.Message
		up.RandomId = userUpdate.RandomId

		if t == constants.UpdateTypeUpdateNewMessage {
			return mtproto.NewTLUpdateNewMessage(up).To_Update()
		} else if t == constants.UpdateTypeUpdateEditMessage {
			return mtproto.NewTLUpdateEditMessage(up).To_Update()
		} else if t == constants.UpdateTypeUpdateNewChannelMessage {
			return mtproto.NewTLUpdateNewChannelMessage(up).To_Update()
		}
		return mtproto.NewTLUpdateEditChannelMessage(up).To_Update()

	case constants.UpdateTypeUpdatePinnedMessages:
		//  updatePinnedMessages#ed85eab5 flags:# pinned:flags.0?true peer:Peer messages:Vector<int> pts:int pts_count:int = Update;
		up.Pinned = userUpdate.FoldId == 1
		up.Pts = userUpdate.Pts
		up.PtsCount = userUpdate.PtsCount
		up.Messages = userUpdate.MessageBoxIds
		up.Peer9961FD5C71 = input.MakeInputPeerValue(userUpdate.PeerId, constants.PeerTypeFromInt32(userUpdate.PeerType)).ToPeer()
		return mtproto.NewTLUpdatePinnedMessages(up).To_Update()

	case constants.UpdateTypeUpdateMessageID:
		up.Id4E90BFD671 = userUpdate.Id
		up.RandomId = userUpdate.RandomId
		return mtproto.NewTLUpdateMessageID(up).To_Update()
	case constants.UpdateTypeUpdateDeleteMessages:
		// updateDeleteMessages#a20db0e5 messages:Vector<int> pts:int pts_count:int = Update;
		up.Messages = userUpdate.MessageBoxIds
		up.Pts = userUpdate.Pts
		up.PtsCount = userUpdate.PtsCount
		return mtproto.NewTLUpdateDeleteMessages(up).To_Update()
	case constants.UpdateTypeUpdateReadMessagesContents:
		// updateReadMessagesContents#68c13933 messages:Vector<int> pts:int pts_count:int = Update;
		up.Messages = userUpdate.MessageBoxIds
		up.Pts = userUpdate.Pts
		up.PtsCount = userUpdate.PtsCount
		return mtproto.NewTLUpdateReadMessagesContents(up).To_Update()
	case constants.UpdateTypeUpdateReadHistoryOutbox:
		// updateReadHistoryOutbox#2f2f21bf peer:Peer max_id:int pts:int pts_count:int = Update;
		up.Pts = userUpdate.Pts
		up.PtsCount = userUpdate.PtsCount
		up.MaxId = userUpdate.MaxId
		up.Peer9961FD5C71 = input.MakeInputPeerValue(userUpdate.PeerId, constants.PeerType(userUpdate.PeerType)).ToPeer()
		return mtproto.NewTLUpdateReadHistoryOutbox(up).To_Update()
	case constants.UpdateTypeUpdateReadHistoryInbox:
		// updateReadHistoryInbox#9c974fdf flags:# folder_id:flags.0?int peer:Peer max_id:int still_unread_count:int pts:int pts_count:int = Update;
		// updateReadHistoryInbox#9961fd5c peer:Peer max_id:int pts:int pts_count:int = Update;
		up.Pts = userUpdate.Pts
		up.PtsCount = userUpdate.PtsCount
		up.MaxId = userUpdate.MaxId
		up.FolderId = userUpdate.FoldId
		up.StillUnreadCount = userUpdate.StillUnreadCount
		up.Peer9961FD5C71 = input.MakeInputPeerValue(userUpdate.PeerId, constants.PeerType(userUpdate.PeerType)).ToPeer()
		return mtproto.NewTLUpdateReadHistoryInbox(up).To_Update()
	//case constants.UpdateTypeUpdateChannel:
	//	//  updateChannel#b6d45656 channel_id:int = Update;
	//case constants.UpdateTypeUpdateChannelTooLong:
	//  updateChannelTooLong#eb0467fb flags:# channel_id:int pts:flags.0?int = Update;
	case constants.UpdateTypeUpdateReadChannelInbox:
		//  updateReadChannelInbox#330b5424 flags:# folder_id:flags.0?int channel_id:int max_id:int still_unread_count:int pts:int = Update;
		//  updateReadChannelInbox#4214f37f channel_id:int max_id:int = Update;
		up.Pts = userUpdate.Pts
		up.MaxId = userUpdate.MaxId
		up.FolderId = userUpdate.FoldId
		up.ChannelId = constants.PeerTypeFromChannelIDType(userUpdate.ChatId).ToInt32()
		up.StillUnreadCount = userUpdate.StillUnreadCount
		return mtproto.NewTLUpdateReadChannelInbox(up).To_Update()
	case constants.UpdateTypeUpdateDeleteChannelMessages:
		//  updateDeleteChannelMessages#c37521c9 channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
		up.ChannelId = constants.PeerTypeFromChannelIDType(userUpdate.ChatId).ToInt32()
		up.Messages = userUpdate.MessageBoxIds
		up.Pts = userUpdate.Pts
		up.PtsCount = userUpdate.PtsCount
		return mtproto.NewTLUpdateDeleteChannelMessages(up).To_Update()
	//case constants.UpdateTypeUpdateChannelMessageViews:
	//  updateChannelMessageViews#98a12b4b channel_id:int id:int views:int = Update;
	//mtproto.NewTLUpdateChannelMessageViews()
	case constants.UpdateTypeUpdateReadChannelOutbox:
		//  updateReadChannelOutbox#25d6c9c7 channel_id:int max_id:int = Update;
		up.ChannelId = constants.PeerTypeFromChannelIDType(userUpdate.ChatId).ToInt32()
		up.MaxId = userUpdate.MaxId
		return mtproto.NewTLUpdateReadChannelOutbox(up).To_Update()
	//case constants.UpdateTypeUpdateChannelWebPage:
	//	//  updateChannelWebPage#40771900 channel_id:int webpage:WebPage pts:int pts_count:int = Update;
	//	up.Webpage = userUpdate.WebPage
	//	mtproto.NewTLUpdateChannelWebPage()
	case constants.UpdateTypeUpdateChannelReadMessagesContents:
		//  updateChannelReadMessagesContents#89893b45 channel_id:int messages:Vector<int> = Update;
		up.ChannelId = constants.PeerTypeFromChannelIDType(userUpdate.ChatId).ToInt32()
		up.Messages = userUpdate.MessageBoxIds
		return mtproto.NewTLUpdateChannelReadMessagesContents(up).To_Update()
	case constants.UpdateTypeUpdateChannelAvailableMessages:
		//  updateChannelAvailableMessages#70db6837 channel_id:int available_min_id:int = Update;
		up.AvailableMinId = userUpdate.FoldId
		up.ChannelId = constants.PeerTypeFromChannelIDType(userUpdate.ChatId).ToInt32()
		return mtproto.NewTLUpdateChannelAvailableMessages(up).To_Update()
	//case constants.UpdateTypeUpdateChannelParticipant:
	//	//  updateChannelParticipant#65d2b464 flags:# channel_id:int date:int user_id:int prev_participant:flags.0?ChannelParticipant new_participant:flags.1?ChannelParticipant qts:int = Update;
	//	up.ChannelId = userUpdate.ChatId
	//	up.Date = userUpdate.Date
	//	up.UserId = userUpdate.UserId
	//	up.PrevParticipant
	//	up.NewParticipant
	//	up.Qts
	//	return mtproto.NewTLUpdateChannelParticipant(up).To_Update()
	case constants.UpdateTypeUpdateChannelMessageForwards:
		//  updateChannelMessageForwards#6e8a84df channel_id:int id:int forwards:int = Update;
		up.ChannelId = constants.PeerTypeFromChannelIDType(userUpdate.ChatId).ToInt32()
		up.Id4E90BFD671 = userUpdate.MessageBoxIds[0]
		up.Forwards = userUpdate.FoldId
		return mtproto.NewTLUpdateChannelMessageForwards(up).To_Update()
	//case constants.UpdateTypeUpdateReadChannelDiscussionInbox:
	//	//  updateReadChannelDiscussionInbox#1cc7de54 flags:# channel_id:int top_msg_id:int read_max_id:int broadcast_id:flags.0?int broadcast_post:flags.0?int = Update;
	//	up.ChannelId = userUpdate.ChatId
	//	up.TopMsgId
	//	up.ReadMaxId
	//	up.BroadcastId
	//	up.BroadcastPost
	//	return mtproto.NewTLUpdateReadChannelDiscussionInbox(up).To_Update()
	case constants.UpdateTypeUpdateReadChannelDiscussionOutbox:
		//  updateReadChannelDiscussionOutbox#4638a26c channel_id:int top_msg_id:int read_max_id:int = Update;
		up.ChannelId = constants.PeerTypeFromChannelIDType(userUpdate.ChatId).ToInt32()
		up.TopMsgId = userUpdate.FoldId
		up.ReadMaxId = userUpdate.MaxId
		return mtproto.NewTLUpdateReadChannelDiscussionOutbox(up).To_Update()
	case constants.UpdateTypeUpdateChannelUserTyping:
		//  updateChannelUserTyping#ff2abe9f flags:# channel_id:int top_msg_id:flags.0?int user_id:int action:SendMessageAction = Update;
		up.ChannelId = constants.PeerTypeFromChannelIDType(userUpdate.ChatId).ToInt32()
		up.TopMsgId = userUpdate.FoldId
		up.UserId = constants.PeerTypeFromUserIDType(userUpdate.UserId).ToInt32()
		up.Action = &mtproto.SendMessageAction{}
		jsoniter.Unmarshal([]byte(userUpdate.MessageAction), up.Action)
		return mtproto.NewTLUpdateChannelUserTyping(up).To_Update()
	case constants.UpdateTypeUpdatePinnedChannelMessages:
		//  updatePinnedChannelMessages#8588878b flags:# pinned:flags.0?true channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
		up.Pinned = userUpdate.FoldId == 1
		up.ChannelId = constants.PeerTypeFromChannelIDType(userUpdate.ChatId).ToInt32()
		up.Messages = userUpdate.MessageBoxIds
		up.Pts = userUpdate.Pts
		up.PtsCount = userUpdate.PtsCount
		return mtproto.NewTLUpdatePinnedChannelMessages(up).To_Update()
	case constants.UpdateTypeUpdateChannelPinnedMessage:
		//  updateChannelPinnedMessage#98592475 channel_id:int id:int = Update;
		up.ChannelId = constants.PeerTypeFromChannelIDType(userUpdate.ChatId).ToInt32()
		up.Id4E90BFD671 = userUpdate.FoldId
		return mtproto.NewTLUpdateChannelPinnedMessage(up).To_Update()
	case constants.UpdateTypeUpdateChatUserTyping:
		//  updateChatUserTyping#9a65ea1f chat_id:int user_id:int action:SendMessageAction = Update;
		up.ChatId = constants.PeerTypeFromChannelIDType(userUpdate.ChatId).ToInt32()
		up.UserId = constants.PeerTypeFromUserIDType(userUpdate.UserId).ToInt32()
		up.Action = &mtproto.SendMessageAction{}
		jsoniter.Unmarshal([]byte(userUpdate.MessageAction), up.Action)
		return mtproto.NewTLUpdateChatUserTyping(up).To_Update()
	case constants.UpdateTypeUpdateChatParticipants:
		//  chatParticipant#c8d7493e user_id:int inviter_id:int date:int = ChatParticipant;
		up.UserId = constants.PeerTypeFromUserIDType(userUpdate.UserId).ToInt32()
		up.InviterId = userUpdate.FoldId
		up.Date = userUpdate.Date
		return mtproto.NewTLUpdateChatParticipants(up).To_Update()
	//case constants.UpdateTypeUpdateEncryptedChatTyping:
	//	//  updateEncryptedChatTyping#1710f156 chat_id:int = Update;
	//	mtproto.NewTLUpdateEncryptedChatTyping(up).To_Update()
	case constants.UpdateTypeUpdateChatParticipantAdd:
		//  updateChatParticipantAdd#ea4b0e5c chat_id:int user_id:int inviter_id:int date:int version:int = Update;
		up.ChatId = constants.PeerTypeFromChatIDType(userUpdate.ChatId).ToInt32()
		up.Version = userUpdate.MaxId
		up.Date = userUpdate.Date
		up.InviterId = userUpdate.FoldId
		up.UserId = constants.PeerTypeFromUserIDType(userUpdate.UserId).ToInt32()
		return mtproto.NewTLUpdateChatParticipantAdd(up).To_Update()
	//case constants.UpdateTypeUpdateChatParticipantDelete:
	//case constants.UpdateTypeUpdateChatParticipantAdmin:
	//case constants.UpdateTypeUpdateChatDefaultBannedRights:
	//case constants.UpdateTypeUpdateChat:
	//case constants.UpdateTypeUpdateChatAdmins:
	//case constants.UpdateTypeUpdateShortChatMessage:
	//  updateShortChatMessage#16812688 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
	case constants.UpdateTypeUpdateFolderPeers:
		//  updateFolderPeers#19360dc0 folder_peers:Vector<FolderPeer> pts:int pts_count:int = Update;
		up.UserId = constants.PeerTypeFromUserIDType(userUpdate.UserId).ToInt32()
		up.Pts = userUpdate.Pts
		up.PtsCount = userUpdate.PtsCount
		if len(userUpdate.MessageAction) > 0 {
			//TODO:(Coderxw)
			//up.FolderPeers = make([]*mtproto.FolderPeer, 0, 1)
			//jsoniter.Unmarshal([]byte(userUpdate.MessageAction), &up.RequestContacts)
		} else {
			up.FolderPeers = []*mtproto.FolderPeer{}
		}
		return mtproto.NewTLUpdateFolderPeers(up).To_Update()
	default:
		return nil
	}
	panic(fmt.Sprintf("error:%v", t))
}

func UpdateToUserUpdate(userId int64, up *mtproto.Update) *data_update.UserUpdate {
	updateType := update.ToUpdateType(up)
	userUpdate := &data_update.UserUpdate{
		OwnerUserId:   userId,
		Id:            0,
		UserId:        0,
		PeerId:        0,
		PeerType:      0,
		Pts:           up.Pts,
		PtsCount:      up.PtsCount,
		UpdateType:    updateType.ToInt32(),
		Date:          up.Date,
		Version:       update.UpdateVersion1.ToInt32(),
		Time:          util.FormatYMDHMSNow(time.Now()),
		RandomId:      0,
		MessageBoxIds: nil,
		MessageAction: "",
		ChatId:        0,
		MaxId:         0,
	}
	switch updateType {
	case constants.UpdateTypeUpdateNewMessage, constants.UpdateTypeUpdateEditMessage,
		constants.UpdateTypeUpdateNewChannelMessage, constants.UpdateTypeUpdateEditChannelMessage:
		// updateEditMessage#e40370a3 message:Message pts:int pts_count:int = Update;
		// updateNewMessage#1f2b0afd message:Message pts:int pts_count:int = Update;
		// updateNewChannelMessage#62ba04d9 message:Message pts:int pts_count:int = Update;
		// updateEditChannelMessage#1b3f4df7 message:Message pts:int pts_count:int = Update;
		//
		userUpdate.Date = up.Date

		inputPeer := input.MakePeer(up.Message1F2B0AFD71.PeerId)
		userUpdate.MessageData = message.ToMessageData(0,
			constants.PeerTypeFromUserIDType32(up.Message1F2B0AFD71.FromId90DDDC1171).ToInt(),
			inputPeer.GetPeerId(), inputPeer.GetPeerType(), up.Message1F2B0AFD71)
		userUpdate.RandomId = up.RandomId
		return userUpdate
	case constants.UpdateTypeUpdatePinnedMessages:
		//  updatePinnedMessages#ed85eab5 flags:# pinned:flags.0?true peer:Peer messages:Vector<int> pts:int pts_count:int = Update;
		inputPeer := input.MakePeer(up.Peer9961FD5C71)
		userUpdate.PeerType = inputPeer.GetPeerType().ToInt32()
		userUpdate.PeerId = inputPeer.GetPeerId()
		if up.Pinned {
			userUpdate.FoldId = 1
		}
		userUpdate.Pts = up.Pts
		userUpdate.PtsCount = up.PtsCount
		userUpdate.MessageBoxIds = up.Messages
		return userUpdate

	case constants.UpdateTypeUpdateMessageID:
		userUpdate.Id = up.Id4E90BFD671
		userUpdate.Time = util.FormatYMDHMSNow(time.Now())
		userUpdate.RandomId = up.RandomId
		return userUpdate
	case constants.UpdateTypeUpdateDeleteMessages:
		// updateDeleteMessages#a20db0e5 messages:Vector<int> pts:int pts_count:int = Update;
		userUpdate.MessageBoxIds = up.Messages
		userUpdate.Pts = up.Pts
		userUpdate.PtsCount = up.PtsCount
		return userUpdate
	case constants.UpdateTypeUpdateReadMessagesContents:
		// updateReadMessagesContents#68c13933 messages:Vector<int> pts:int pts_count:int = Update;
		userUpdate.MessageBoxIds = up.Messages
		userUpdate.Pts = up.Pts
		userUpdate.PtsCount = up.PtsCount
		return userUpdate
	case constants.UpdateTypeUpdateReadHistoryOutbox:
		// updateReadHistoryOutbox#2f2f21bf peer:Peer max_id:int pts:int pts_count:int = Update;
		upInputPeer := input.MakePeer(up.Peer9961FD5C71)
		userUpdate.Pts = up.Pts
		userUpdate.PtsCount = up.PtsCount
		userUpdate.MaxId = up.MaxId
		userUpdate.PeerId = upInputPeer.GetPeerId()
		userUpdate.PeerType = upInputPeer.GetPeerType().ToInt32()
		return userUpdate
	case constants.UpdateTypeUpdateReadHistoryInbox:
		// updateReadHistoryInbox#9c974fdf flags:# folder_id:flags.0?int peer:Peer max_id:int still_unread_count:int pts:int pts_count:int = Update;
		// updateReadHistoryInbox#9961fd5c peer:Peer max_id:int pts:int pts_count:int = Update;
		upInputPeer := input.MakePeer(up.Peer9961FD5C71)
		userUpdate.Pts = up.Pts
		userUpdate.PtsCount = up.PtsCount
		userUpdate.MaxId = up.MaxId
		userUpdate.FoldId = up.FolderId
		userUpdate.StillUnreadCount = up.StillUnreadCount
		userUpdate.PeerId = upInputPeer.GetPeerId()
		userUpdate.PeerType = upInputPeer.GetPeerType().ToInt32()
		return userUpdate
	case constants.UpdateTypeUpdateReadChannelInbox:
		//  updateReadChannelInbox#330b5424 flags:# folder_id:flags.0?int channel_id:int max_id:int still_unread_count:int pts:int = Update;
		//  updateReadChannelInbox#4214f37f channel_id:int max_id:int = Update;
		userUpdate.Pts = up.Pts
		userUpdate.MaxId = up.MaxId
		userUpdate.FoldId = up.FolderId
		userUpdate.ChatId = constants.PeerTypeFromChannelIDType32(up.ChannelId).ToInt()
		userUpdate.StillUnreadCount = up.StillUnreadCount
		return userUpdate
	case constants.UpdateTypeUpdateDeleteChannelMessages:
		//  updateDeleteChannelMessages#c37521c9 channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
		userUpdate.ChatId = constants.PeerTypeFromChannelIDType32(up.ChannelId).ToInt()
		userUpdate.MessageBoxIds = up.Messages
		userUpdate.Pts = up.Pts
		userUpdate.PtsCount = up.PtsCount
		return userUpdate
	case constants.UpdateTypeUpdateReadChannelOutbox:
		//  updateReadChannelOutbox#25d6c9c7 channel_id:int max_id:int = Update;
		userUpdate.ChatId = constants.PeerTypeFromChannelIDType32(up.ChannelId).ToInt()
		userUpdate.MaxId = up.MaxId
		return userUpdate
	//case constants.UpdateTypeUpdateChannelWebPage:
	//	//  updateChannelWebPage#40771900 channel_id:int webpage:WebPage pts:int pts_count:int = Update;
	//	userUpdate.Webpage = up.Webpage
	//	return userUpdate
	case constants.UpdateTypeUpdateChannelReadMessagesContents:
		//  updateChannelReadMessagesContents#89893b45 channel_id:int messages:Vector<int> = Update;
		userUpdate.ChatId = constants.PeerTypeFromChannelIDType32(up.ChannelId).ToInt()
		userUpdate.MessageBoxIds = up.Messages
		return userUpdate
	case constants.UpdateTypeUpdateChannelAvailableMessages:
		//  updateChannelAvailableMessages#70db6837 channel_id:int available_min_id:int = Update;
		userUpdate.FoldId = up.AvailableMinId
		userUpdate.ChatId = constants.PeerTypeFromChannelIDType32(up.ChannelId).ToInt()
		return userUpdate
	//case constants.UpdateTypeUpdateChannelParticipant:
	//	//  updateChannelParticipant#65d2b464 flags:# channel_id:int date:int user_id:int prev_participant:flags.0?ChannelParticipant new_participant:flags.1?ChannelParticipant qts:int = Update;
	//	up.ChannelId = userUpdate.ChatId
	//	up.Date = userUpdate.Date
	//	up.UserId = userUpdate.UserId
	//	up.PrevParticipant
	//	up.NewParticipant
	//	up.Qts
	//	return userUpdate
	case constants.UpdateTypeUpdateChannelMessageForwards:
		//  updateChannelMessageForwards#6e8a84df channel_id:int id:int forwards:int = Update;
		userUpdate.ChatId = constants.PeerTypeFromChannelIDType32(up.ChannelId).ToInt()
		userUpdate.MessageBoxIds = []int32{up.Id4E90BFD671}
		userUpdate.FoldId = up.Forwards
		return userUpdate
	//case constants.UpdateTypeUpdateReadChannelDiscussionInbox:
	//	//  updateReadChannelDiscussionInbox#1cc7de54 flags:# channel_id:int top_msg_id:int read_max_id:int broadcast_id:flags.0?int broadcast_post:flags.0?int = Update;
	//	up.ChannelId = userUpdate.ChatId
	//	up.TopMsgId
	//	up.ReadMaxId
	//	up.BroadcastId
	//	up.BroadcastPost
	//	return userUpdate
	case constants.UpdateTypeUpdateReadChannelDiscussionOutbox:
		//  updateReadChannelDiscussionOutbox#4638a26c channel_id:int top_msg_id:int read_max_id:int = Update;
		userUpdate.ChatId = constants.PeerTypeFromChannelIDType32(up.ChannelId).ToInt()
		userUpdate.FoldId = up.TopMsgId
		userUpdate.MaxId = up.ReadMaxId
		return userUpdate
	case constants.UpdateTypeUpdateChannelUserTyping:
		//  updateChannelUserTyping#ff2abe9f flags:# channel_id:int top_msg_id:flags.0?int user_id:int action:SendMessageAction = Update;
		userUpdate.ChatId = constants.PeerTypeFromChannelIDType32(up.ChannelId).ToInt()
		userUpdate.FoldId = up.TopMsgId
		userUpdate.UserId = constants.PeerTypeFromUserIDType32(up.UserId).ToInt()
		up.Action = &mtproto.SendMessageAction{}
		buf, _ := jsoniter.Marshal(up.Action)
		userUpdate.MessageAction = string(buf)
		return userUpdate
	case constants.UpdateTypeUpdatePinnedChannelMessages:
		//  updatePinnedChannelMessages#8588878b flags:# pinned:flags.0?true channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
		userUpdate.FoldId = 0
		if up.Pinned {
			userUpdate.FoldId = 1
		}
		userUpdate.ChatId = constants.PeerTypeFromChannelIDType32(up.ChannelId).ToInt()
		userUpdate.MessageBoxIds = up.Messages
		userUpdate.Pts = up.Pts
		userUpdate.PtsCount = up.PtsCount
		return userUpdate
	case constants.UpdateTypeUpdateChannelPinnedMessage:
		//  updateChannelPinnedMessage#98592475 channel_id:int id:int = Update;
		userUpdate.ChatId = constants.PeerTypeFromChannelIDType32(up.ChannelId).ToInt()
		userUpdate.FoldId = up.Id4E90BFD671
		return userUpdate
	case constants.UpdateTypeUpdateChatUserTyping:
		//  updateChatUserTyping#9a65ea1f chat_id:int user_id:int action:SendMessageAction = Update;
		userUpdate.ChatId = constants.PeerTypeFromChatIDType32(up.ChatId).ToInt()
		userUpdate.UserId = constants.PeerTypeFromUserIDType32(up.UserId).ToInt()
		up.Action = &mtproto.SendMessageAction{}
		buf, _ := jsoniter.Marshal(up.Action)
		userUpdate.MessageAction = string(buf)
		return userUpdate
	case constants.UpdateTypeUpdateChatParticipants:
		//  chatParticipant#c8d7493e user_id:int inviter_id:int date:int = ChatParticipant;
		userUpdate.UserId = constants.PeerTypeFromChatIDType32(up.UserId).ToInt()
		userUpdate.FoldId = up.InviterId
		userUpdate.Date = up.Date
		return userUpdate
	case constants.UpdateTypeUpdateChatParticipantAdd:
		//  updateChatParticipantAdd#ea4b0e5c chat_id:int user_id:int inviter_id:int date:int version:int = Update;
		userUpdate.ChatId = constants.PeerTypeFromChatIDType32(up.ChatId).ToInt()
		userUpdate.MaxId = up.Version
		userUpdate.Date = up.Date
		userUpdate.FoldId = up.InviterId
		userUpdate.UserId = constants.PeerTypeFromUserIDType32(up.UserId).ToInt()
		return userUpdate
	case constants.UpdateTypeUpdateFolderPeers:
		//  updateFolderPeers#19360dc0 folder_peers:Vector<FolderPeer> pts:int pts_count:int = Update;
		userUpdate.UserId = constants.PeerTypeFromUserIDType32(up.UserId).ToInt()
		userUpdate.Pts = up.Pts
		userUpdate.PtsCount = up.PtsCount
		buf, _ := jsoniter.Marshal(up.FolderPeers)
		userUpdate.MessageAction = string(buf)
		return userUpdate
	default:
		panic(fmt.Sprintf("error:%v", updateType))
		return nil
	}
}

func FilterAppendUserId(userIdList *[]int64, id int64) {
	if id == 0 {
		return
	}
	idx := util.IndexInt64s(userIdList, id)
	if idx >= 0 {
		return
	}

	*userIdList = append(*userIdList, id)
}

func FilterAppendPeer(userIdList *[]int64, peer *mtproto.Peer) {
	if peer != nil && peer.ClassName == mtproto.ClassPeerUser {
		FilterAppendUserId(userIdList, constants.PeerTypeFromUserIDType32(peer.UserId).ToInt())
		return
	}
}
