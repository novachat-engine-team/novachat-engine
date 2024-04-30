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
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
	"novachat_engine/service/data/messages/message"
	"novachat_engine/service/input"
)

//  + TL_messageEntityUnknown
//  + TL_messageEntityMention
//  + TL_messageEntityHashtag
//  + TL_messageEntityBotCommand
//  + TL_messageEntityUrl
//  + TL_messageEntityEmail
//  + TL_messageEntityBold
//  + TL_messageEntityItalic
//  + TL_messageEntityCode
//  + TL_messageEntityPre
//  + TL_messageEntityTextUrl
//  + TL_messageEntityMentionName
//  + TL_inputMessageEntityMentionName
//  + TL_messageEntityPhone
//  + TL_messageEntityCashtag
//  + TL_messageEntityUnderline
//  + TL_messageEntityStrike
//  + TL_messageEntityBlockquote
//  + TL_messageEntityBankCard

type MessageEntityType int32

const (
	MessageEntityTypeUnknown          MessageEntityType = 0
	MessageEntityTypeMention          MessageEntityType = 1
	MessageEntityTypeHashtag          MessageEntityType = 2
	MessageEntityTypeBotCommand       MessageEntityType = 3
	MessageEntityTypeUrl              MessageEntityType = 4
	MessageEntityTypeEmail            MessageEntityType = 5
	MessageEntityTypeBold             MessageEntityType = 6
	MessageEntityTypeItalic           MessageEntityType = 7
	MessageEntityTypeCode             MessageEntityType = 8
	MessageEntityTypePre              MessageEntityType = 9
	MessageEntityTypeTextUrl          MessageEntityType = 10
	MessageEntityTypeMentionName      MessageEntityType = 11
	InputMessageTypeEntityMentionName MessageEntityType = 12
	MessageEntityTypePhone            MessageEntityType = 13
	MessageEntityTypeCashtag          MessageEntityType = 14
	MessageEntityTypeUnderline        MessageEntityType = 15
	MessageEntityTypeStrike           MessageEntityType = 16
	MessageEntityTypeBlockquote       MessageEntityType = 17
	MessageEntityTypeBankCard         MessageEntityType = 18
)

func (m MessageEntityType) ToInt32() int32 {
	return int32(m)
}

func FromMessageEntityTypeInt32(v int32) MessageEntityType {
	return MessageEntityType(v)
}

func ToMessageEntity(t MessageEntityType, v *mtproto.MessageEntity) *mtproto.MessageEntity {
	if v == nil {
		v = &mtproto.MessageEntity{}
	}

	switch t {
	case MessageEntityTypeUnknown:
		return v.To_MessageEntityUnknown().To_MessageEntity()
	case MessageEntityTypeMention:
		return v.To_MessageEntityMention().To_MessageEntity()
	case MessageEntityTypeHashtag:
		return v.To_MessageEntityHashtag().To_MessageEntity()
	case MessageEntityTypeBotCommand:
		return v.To_MessageEntityBotCommand().To_MessageEntity()
	case MessageEntityTypeUrl:
		return v.To_MessageEntityUrl().To_MessageEntity()
	case MessageEntityTypeEmail:
		return v.To_MessageEntityEmail().To_MessageEntity()
	case MessageEntityTypeBold:
		return v.To_MessageEntityBold().To_MessageEntity()
	case MessageEntityTypeItalic:
		return v.To_MessageEntityItalic().To_MessageEntity()
	case MessageEntityTypeCode:
		return v.To_MessageEntityCode().To_MessageEntity()
	case MessageEntityTypePre:
		return v.To_MessageEntityPre().To_MessageEntity()
	case MessageEntityTypeTextUrl:
		return v.To_MessageEntityTextUrl().To_MessageEntity()
	case MessageEntityTypeMentionName:
		return v.To_MessageEntityMentionName().To_MessageEntity()
	case InputMessageTypeEntityMentionName:
		return v.To_InputMessageEntityMentionName().To_MessageEntity()
	case MessageEntityTypePhone:
		return v.To_MessageEntityPhone().To_MessageEntity()
	case MessageEntityTypeCashtag:
		return v.To_MessageEntityCashtag().To_MessageEntity()
	case MessageEntityTypeUnderline:
		return v.To_MessageEntityUnderline().To_MessageEntity()
	case MessageEntityTypeStrike:
		return v.To_MessageEntityStrike().To_MessageEntity()
	case MessageEntityTypeBlockquote:
		return v.To_MessageEntityBlockquote().To_MessageEntity()
	case MessageEntityTypeBankCard:
		return v.To_MessageEntityBankCard().To_MessageEntity()
	}
	return v.To_MessageEntityUnknown().To_MessageEntity()
}

func ToMessageEntityType(entity *mtproto.MessageEntity) MessageEntityType {
	switch entity.ClassName {
	case mtproto.ClassMessageEntityUnknown:
		return MessageEntityTypeUnknown
	case mtproto.ClassMessageEntityMention:
		return MessageEntityTypeMention
	case mtproto.ClassMessageEntityHashtag:
		return MessageEntityTypeHashtag
	case mtproto.ClassMessageEntityBotCommand:
		return MessageEntityTypeBotCommand
	case mtproto.ClassMessageEntityUrl:
		return MessageEntityTypeUrl
	case mtproto.ClassMessageEntityEmail:
		return MessageEntityTypeEmail
	case mtproto.ClassMessageEntityBold:
		return MessageEntityTypeBold
	case mtproto.ClassMessageEntityItalic:
		return MessageEntityTypeItalic
	case mtproto.ClassMessageEntityCode:
		return MessageEntityTypeCode
	case mtproto.ClassMessageEntityPre:
		return MessageEntityTypePre
	case mtproto.ClassMessageEntityTextUrl:
		return MessageEntityTypeTextUrl
	case mtproto.ClassMessageEntityMentionName:
		return MessageEntityTypeMentionName
	case mtproto.ClassInputMessageEntityMentionName:
		return InputMessageTypeEntityMentionName
	case mtproto.ClassMessageEntityPhone:
		return MessageEntityTypePhone
	case mtproto.ClassMessageEntityCashtag:
		return MessageEntityTypeCashtag
	case mtproto.ClassMessageEntityUnderline:
		return MessageEntityTypeUnderline
	case mtproto.ClassMessageEntityStrike:
		return MessageEntityTypeStrike
	case mtproto.ClassMessageEntityBlockquote:
		return MessageEntityTypeBlockquote
	case mtproto.ClassMessageEntityBankCard:
		return MessageEntityTypeBankCard
	}
	return MessageEntityTypeUnknown
}

func MessageEntityToEntity(entity *mtproto.MessageEntity) *data_message.Entity {
	e := &data_message.Entity{
		Type:       ToMessageEntityType(entity).ToInt32(),
		Offset:     entity.Offset,
		Length:     entity.Length,
		Language:   entity.Language,
		Url:        entity.Url,
		UserId:     constants.PeerTypeFromUserIDType32(entity.UserId352DCA5871).ToInt(),
		PeerType:   0,
		PeerId:     0,
		AccessHash: 0,
	}
	if entity.UserId208E68C971 != nil {
		inputUser := input.MakeInputUser(entity.UserId208E68C971)
		e.PeerId = inputUser.GetId()
		e.PeerType = inputUser.GetInputUserType().ToInt32()
		e.AccessHash = inputUser.GetAccessHash()
	}
	return e
}

func EntityToMessageEntity(entity *data_message.Entity) *mtproto.MessageEntity {
	e := &mtproto.MessageEntity{
		Offset:           entity.Offset,
		Length:           entity.Length,
		Language:         entity.Language,
		Url:              entity.Url,
		UserId352DCA5871: constants.PeerTypeFromUserIDType(entity.UserId).ToInt32(),
	}
	switch constants.InputUserTypeFromInt32(entity.PeerType) {
	case constants.InputUserTypeSelf:
		e.UserId208E68C971 = mtproto.NewTLInputUserSelf(&mtproto.InputUser{
			UserId:     constants.PeerTypeFromUserIDType(entity.PeerId).ToInt32(),
			AccessHash: entity.AccessHash,
		}).To_InputUser()
	case constants.InputUserTypeUser:
		e.UserId208E68C971 = mtproto.NewTLInputUser(&mtproto.InputUser{
			UserId:     constants.PeerTypeFromUserIDType(entity.PeerId).ToInt32(),
			AccessHash: entity.AccessHash,
		}).To_InputUser()
	case constants.InputUserTypeUserFromMessage:
		e.UserId208E68C971 = mtproto.NewTLInputUserFromMessage(&mtproto.InputUser{
			MsgId:      int32(entity.PeerId),
			AccessHash: entity.AccessHash,
		}).To_InputUser()
	default:
	}

	return ToMessageEntity(FromMessageEntityTypeInt32(entity.Type), e)
}

func MessageEntitiesToEntities(entities []*mtproto.MessageEntity) []*data_message.Entity {
	if len(entities) == 0 {
		return nil
	}

	entitiesList := make([]*data_message.Entity, 0, len(entities))

	for _, e := range entities {
		entitiesList = append(entitiesList, MessageEntityToEntity(e))
	}

	return entitiesList
}

func EntitiesToMessageEntities(es []*data_message.Entity) []*mtproto.MessageEntity {
	if len(es) == 0 {
		return []*mtproto.MessageEntity{}
	}

	entitiesList := make([]*mtproto.MessageEntity, 0, len(es))
	for _, e := range es {
		entitiesList = append(entitiesList, EntityToMessageEntity(e))
	}

	return entitiesList
}

func messageEntitiesUtil(entities []*mtproto.MessageEntity) []*data_message.Entity {
	if len(entities) == 0 {
		return nil
	}

	return MessageEntitiesToEntities(entities)
}
