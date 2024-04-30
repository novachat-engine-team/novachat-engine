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
	"novachat_engine/pkg/log"
	"novachat_engine/service/common/draft"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message"
)

func (c *Core) SaveDraft(userId int64, peerId int64, peerType constants.PeerType, replyToMsgId int32, text string, entities []*mtproto.MessageEntity, date int32) (bool, error) {
	log.Debugf("SaveDraft userId:%d peerId:%d peerType:%v", userId, peerId, peerType)
	var draftValue *mtproto.DraftMessage
	if len(text) != 0 || len(entities) != 0 {
		draftValue = mtproto.NewTLDraftMessage(&mtproto.DraftMessage{
			ReplyToMsgId: replyToMsgId,
			Message:      text,
			Entities:     entities,
			Date:         date,
		}).To_DraftMessage()
	}

	return true, c.messageCore.ConversationSaveDraft(userId, peerId, peerType, draft.EncodeDraftMessage(draftValue))
}

type Draft struct {
	PeerId   int64
	PeerType constants.PeerType
	Draft    *mtproto.DraftMessage
}

func (c *Core) GetAllDraft(userId int64) ([]*Draft, error) {
	log.Debugf("GetAllDraft userId:%d", userId)
	conversationList, err := c.messageCore.GetConversationList(userId, true, constants.PeerTypeEmpty, 0, 0, 0)
	if err != nil {
		log.Errorf("GetAllDraft userId:%d error:%s", userId, err.Error())
		return nil, err
	}
	draftList := make([]*Draft, 0, len(conversationList))
	for _, v := range conversationList {
		if len(v.Draft) > 0 {
			v.PeerId = message.MakePeerId(v.PeerId, constants.PeerTypeFromInt32(v.PeerType))
			draftList = append(draftList, &Draft{
				PeerId:   v.PeerId,
				PeerType: constants.PeerTypeFromInt32(v.PeerType),
				Draft:    draft.DecodeDraftMessage(v.Draft),
			})
		}
	}
	return draftList, nil
}

func (c *Core) ClearAllDrafts(userId int64) error {
	log.Debugf("ClearAllDrafts userId:%d", userId)
	return c.messageCore.ConversationClearAllDraft(userId)
}
