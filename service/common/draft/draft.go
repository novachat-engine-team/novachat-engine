/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/31 14:36
 * @File : draft.go
 */

package draft

import (
	jsoniter "github.com/json-iterator/go"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/core/message/message"
	data_message "novachat_engine/service/data/messages/message"
	"time"
)

type Draft struct {
	ReplyToMsgId int32
	Text         string
	Entities     []*data_message.Entity
	Date         int32
}

func DecodeDraftMessage(draftMessage string) *mtproto.DraftMessage {
	if len(draftMessage) == 0 {
		return mtproto.NewTLDraftMessageEmpty(&mtproto.DraftMessage{
			Date: int32(time.Now().Unix()),
		}).To_DraftMessage()
	}

	var draftValue Draft
	log.Debugf("DecodeDraftMessage %s", draftMessage)
	jsoniter.Unmarshal([]byte(draftMessage), &draftValue)

	draftMessageS := &mtproto.DraftMessage{
		ReplyToMsgId: draftValue.ReplyToMsgId,
		Message:      draftValue.Text,
		Entities:     nil,
		Date:         draftValue.Date,
	}
	for _, v := range draftValue.Entities {
		draftMessageS.Entities = append(draftMessageS.Entities, message.EntityToMessageEntity(v))
	}
	return mtproto.NewTLDraftMessage(draftMessageS).To_DraftMessage()
}

func EncodeDraftMessage(draftMessageS *mtproto.DraftMessage) string {
	if draftMessageS == nil || draftMessageS.ClassName == mtproto.ClassDraftMessageEmpty {
		return ""
	}
	log.Debugf("EncodeDraftMessage %v", draftMessageS)

	draftValue := &Draft{
		ReplyToMsgId: draftMessageS.ReplyToMsgId,
		Text:         draftMessageS.Message,
		Entities:     make([]*data_message.Entity, 0, len(draftMessageS.Entities)),
		Date:         draftMessageS.Date,
	}
	for _, v := range draftMessageS.Entities {
		draftValue.Entities = append(draftValue.Entities, message.MessageEntityToEntity(v))
	}

	s, _ := jsoniter.MarshalToString(draftValue)
	return s
}
