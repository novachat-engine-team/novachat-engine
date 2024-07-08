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
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
	messageUtl "novachat_engine/service/core/message"
	"novachat_engine/service/data/messages/message"
	"novachat_engine/service/input"
)

func toMessageType(message *mtproto.Message) constants.MessageType {
	switch message.ClassName {
	case mtproto.ClassMessageEmpty:
		return constants.MessageTypeEmpty
	case mtproto.ClassMessage:
		return constants.MessageTypeMessage
	case mtproto.ClassMessageService:
		return constants.MessageTypeMessageService
	default:
		panic(fmt.Sprintf("toMessageType error type:%s", message.ClassName))
	}
}

func dataToMessage(message *data_message.Message, layer int32) *mtproto.Message {
	peerId, peerType := messageUtl.MakePeerType(message.PeerId)
	m := &mtproto.Message{
		Id:                message.Id,
		Out:               message.Out,
		Mentioned:         false,
		MediaUnread:       false,
		Silent:            false,
		Post:              false,
		FromId90DDDC1171:  constants.PeerTypeFromUserIDType(message.FromUserId).ToInt32(),
		ToId:              input.MakeInputPeerValue(peerId, peerType).ToPeer(),
		FwdFrom:           toMessageFwd(message.Fwd),
		ViaBotId:          0,
		ReplyToMsgId:      message.ReplyTo,
		Date:              message.Date,
		Message:           message.Message,
		Media:             toMessageMedia(message.Media, layer),
		ReplyMarkup:       nil,
		Entities:          EntitiesToMessageEntities(message.Entities),
		Views:             0,
		EditDate:          message.EditDate,
		PostAuthor:        message.PostAuthor,
		Action:            toMessageAction(message.Action),
		Unread:            false,
		GroupedId:         message.GroupId,
		FromScheduled:     false,
		Legacy:            false,
		EditHide:          false,
		RestrictionReason: nil,
		FromId286FA604119: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(message.FromUserId).ToInt32()}).To_Peer(),
		PeerId:            input.MakeInputPeerValue(peerId, peerType).ToPeer(),
		ReplyTo:           nil,
		Pinned:            message.Pinned,
		Forwards:          0,
		Replies:           nil,
	}

	//if m.Out == false && peerType == constants.PeerTypeUser {
	//	m.PeerId, m.ToId = m.FromId286FA604119, m.FromId286FA604119
	//}
	return m
}

func ToMessageData(userId int64, fromUserId int64, peerId int64, peerType constants.PeerType, message *mtproto.Message) *data_message.Message {
	return &data_message.Message{
		UserId:     userId,
		FromUserId: fromUserId,
		PeerId:     messageUtl.MakePeerId(peerId, peerType),
		Id:         message.Id,
		Out:        message.Out,
		GroupId:    message.GroupedId,
		Pinned:     message.Pinned,
		Media:      messageMediaUtil(message.Media),
		Action:     messageActonUtil(message.Action),
		Fwd:        messageFwdUtil(message.FwdFrom),
		Entities:   messageEntitiesUtil(message.Entities),
		Message:    message.Message,
		Deleted:    false,
		ReplyTo:    message.ReplyToMsgId,
		Type:       toMessageType(message).ToInt32(),
		Date:       message.Date,
		EditDate:   message.EditDate,
		PostAuthor: message.PostAuthor,
	}
}

func ToMessage(message *data_message.Message, layer int32) *mtproto.Message {
	switch message.Type {
	case constants.MessageTypeEmpty.ToInt32():
		return mtproto.NewTLMessageEmpty(nil).To_Message()
	case constants.MessageTypeMessage.ToInt32():
		return mtproto.NewTLMessage(dataToMessage(message, layer)).To_Message()
	case constants.MessageTypeMessageService.ToInt32():
		return mtproto.NewTLMessageService(dataToMessage(message, layer)).To_Message()
	default:
		panic(fmt.Sprintf("toMessageType error type:%v", message.Type))
	}
}
