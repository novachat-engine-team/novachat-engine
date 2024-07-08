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
	"math"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/common/message"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	message2 "novachat_engine/service/core/message"
	messageCore "novachat_engine/service/core/message/message"
	"novachat_engine/service/input"
)

func (c *Core) GetHistory(
	userId int64,
	inputPeer *input.InputPeer,
	offsetId int32,
	offsetDate int32,
	addOffset int32,
	minId int32,
	maxId int32,
	limit int32,
	hash int32,
	layer int32) (*mtproto.Messages_Messages, error) {
	log.Infof("GetHistory userId:%d inputPeer:%v offsetDate:%d limit:%d offsetId:%d addOffset:%d", userId, inputPeer, offsetDate, limit, offsetId, addOffset)

	var err error
	messageList := make([]*mtproto.Message, 0, 0)
	chatList := make([]*mtproto.Chat, 0, 0)
	offsetIdOffset := int32(0)

	historyType := constants.GenerateLoadHistoryType(inputPeer.GetPeerType() == constants.PeerTypeChannel, addOffset, limit, maxId)
	switch historyType {
	case constants.LoadHistoryTypeBackward:
		if addOffset == 0 && offsetId == 0 || addOffset == -1 && offsetId == 0 {
			offsetId = math.MaxInt32
		}
		messageList, offsetIdOffset, err = c.loadBackwardHistoryMessages(userId, inputPeer, offsetId, addOffset+limit, minId, layer)
	case constants.LoadHistoryTypeForward:
		messageList, offsetIdOffset, err = c.loadForwardHistoryMessages(userId, inputPeer, offsetId, -addOffset, minId, layer)
	case constants.LoadHistoryTypeFirstUnRead:
		messageList, offsetIdOffset, err = c.loadBackwardHistoryMessages(userId, inputPeer, offsetId, addOffset+limit, minId, layer)
	case constants.LoadHistoryTypeFirstAroundMessage:
		messageList, offsetIdOffset, err = c.loadForwardHistoryMessages(userId, inputPeer, offsetId, limit, minId, layer)
	case constants.LoadHistoryTypeFirstAroundDate:
		messageList, offsetIdOffset, err = c.loadAroundDataHistoryMessages(userId, inputPeer, offsetId, limit, offsetDate, minId, layer)
	case constants.LoadHistoryTypeMax:
		offsetId = math.MaxInt32
		messageList, offsetIdOffset, err = c.loadBackwardHistoryMessages(userId, inputPeer, offsetId, limit, minId, layer)
	}

	pickPeerIdListType := message.PickMessagePeerIdList(messageList)
	userList, err := c.accountUsersCore.GetUserList(userId, pickPeerIdListType.PeerIdList, layer)
	if err != nil {
		log.Warnf("GetHistory UserByUserIdList error:%s", err.Error())
	}

	var messageMessages *mtproto.Messages_Messages
	if inputPeer.GetPeerType() == constants.PeerTypeChannel {
		//  messages.channelMessages#64479808 flags:# inexact:flags.1?true pts:int count:int offset_id_offset:flags.2?int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
		//  messages.channelMessages#99262e37 flags:# pts:int count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
		messageMessages = mtproto.NewTLMessagesChannelMessages(nil).To_Messages_Messages()
		if int32(len(messageList)) > limit {
			messageMessages.Count = limit
			messageMessages.Messages = messageList[:limit]
		} else {
			messageMessages.Count = int32(len(messageList))
			messageMessages.Messages = messageList
		}

		messageMessages.Chats = chatList
		messageMessages.Users = userList
		messageMessages.OffsetIdOffset = offsetIdOffset
	} else {
		if int32(len(messageList)) > limit {
			//  messages.messagesSlice#3a54685e flags:# inexact:flags.1?true count:int next_rate:flags.0?int offset_id_offset:flags.2?int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
			//  messages.messagesSlice#b446ae3 count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
			messageMessages = mtproto.NewTLMessagesMessagesSlice(nil).To_Messages_Messages()
			messageMessages.Count = limit
			messageMessages.Messages = messageList[:limit]
		} else {
			messageMessages = mtproto.NewTLMessagesMessages(nil).To_Messages_Messages()
			messageMessages.Count = int32(len(messageList))
			messageMessages.Messages = messageList
		}

		messageMessages.Chats = chatList
		messageMessages.Users = userList
		messageMessages.OffsetIdOffset = offsetIdOffset
	}
	if messageMessages != nil {
		for idx, m := range messageMessages.Messages {
			if m.Media != nil {
				m.Media.PhotoB5223B0F71 = compat.CompatPhoto(m.Media.PhotoB5223B0F71, layer)
				m.Media.Document = compat.CompatDocument(m.Media.Document, layer)
				m.Media.Game = compat.CompatGame(m.Media.Game, layer)
				m.Media.Webpage = compat.CompatWebPage(m.Media.Webpage, layer)
			}

			if m.Action != nil && m.Action.Photo != nil {
				m.Action.Photo = compat.CompatPhoto(m.Action.Photo, layer)
			}

			messageMessages.Messages[idx] = m
		}
	}
	return messageMessages, nil
}

func (c *Core) loadForwardHistoryMessages(
	userId int64,
	inputPeer *input.InputPeer,
	messageId int32,
	limit int32,
	minId int32,
	layer int32) ([]*mtproto.Message, int32, error) {

	log.Debugf("loadForwardHistoryMessages userId:%d peerId:%v, messageId:%d, limit:%d", userId, inputPeer, messageId, limit)

	messageDataList, err := c.messageCore.GetHistoryMessages(userId, inputPeer.GetPeerId(), inputPeer.GetPeerType(), messageId, limit, minId, 0, false)
	if err != nil {
		log.Errorf("loadForwardHistoryMessages userId:%d peerId:%v, messageId:%d, limit:%d error:%s", userId, inputPeer, messageId, limit, err.Error())
		return nil, 0, err
	}

	offset := int32(0)
	messageList := make([]*mtproto.Message, 0, len(messageDataList))
	for idx, v := range messageDataList {
		if messageId == v.Id {
			offset = int32(idx)
		}
		_, peerType := message2.MakePeerType(v.PeerId)
		if peerType == constants.PeerTypeUser {
			if !v.Out {
				v.PeerId = v.FromUserId
			}
		} else {
			v.Out = userId == v.FromUserId
		}
		messageList = append(messageList, messageCore.ToMessage(v, layer))
	}
	return messageList, offset, nil
}

func (c *Core) loadBackwardHistoryMessages(
	userId int64,
	inputPeer *input.InputPeer,
	messageId int32,
	limit int32,
	minId int32,
	layer int32) ([]*mtproto.Message, int32, error) {
	log.Debugf("loadBackwardHistoryMessages userId:%d peerId:%v, messageId:%d, limit:%d", userId, inputPeer, messageId, limit)

	messageDataList, err := c.messageCore.GetHistoryMessages(userId, inputPeer.GetPeerId(), inputPeer.GetPeerType(), messageId, -limit, minId, 0, false)
	if err != nil {
		log.Errorf("loadBackwardHistoryMessages userId:%d peerId:%v, messageId:%d, limit:%d error:%s", userId, inputPeer, messageId, limit, err.Error())
		return nil, 0, err
	}

	maxMessageId := int32(0)
	offset := int32(0)
	messageList := make([]*mtproto.Message, 0, len(messageDataList))
	for idx, v := range messageDataList {
		if messageId == v.Id {
			offset = int32(idx)
		}
		if v.Id >= maxMessageId {
			maxMessageId = v.Id
		}
		_, peerType := message2.MakePeerType(v.PeerId)
		if peerType == constants.PeerTypeUser {
			if !v.Out {
				v.PeerId = v.FromUserId
			}
		} else {
			v.Out = userId == v.FromUserId
		}
		messageList = append(messageList, messageCore.ToMessage(v, layer))
	}

	return messageList, offset, nil
}

func (c *Core) loadAroundDataHistoryMessages(
	userId int64,
	inputPeer *input.InputPeer,
	messageId int32,
	limit int32,
	date int32,
	minId int32,
	layer int32) ([]*mtproto.Message, int32, error) {

	messageDataList, err := c.messageCore.GetHistoryMessages(userId, inputPeer.GetPeerId(), inputPeer.GetPeerType(), messageId, limit, minId, date, false)
	if err != nil {
		log.Errorf("loadBackwardHistoryMessages userId:%d peerId:%v, messageId:%d, limit:%d error:%s", userId, inputPeer, messageId, limit, err.Error())
		return nil, 0, err
	}

	offset := int32(0)
	messageList := make([]*mtproto.Message, 0, len(messageDataList))
	for idx, v := range messageDataList {
		if messageId == v.Id {
			offset = int32(idx)
		}
		_, peerType := message2.MakePeerType(v.PeerId)
		if peerType == constants.PeerTypeUser {
			if !v.Out {
				v.PeerId = v.FromUserId
			}
		} else {
			v.Out = userId == v.FromUserId
		}
		messageList = append(messageList, messageCore.ToMessage(v, layer))
	}
	return messageList, offset, nil
}
