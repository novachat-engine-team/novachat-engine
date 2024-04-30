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
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
)

func (c *Core) GetDialog(userId int64,
	excludePinned bool,
	foldId int32,
	offsetId int32,
	offsetDate int32,
	offsetPeer *mtproto.InputPeer,
	limit int32,
	hash int32,
	layer int32) (*mtproto.Messages_Dialogs, error) {
	inputPeer := input.MakeInputPeer(offsetPeer)

	// TODO: (Coderxw) thinking... offsetId(0/-1)
	_ = offsetId

	peerType := inputPeer.GetPeerType()
	switch inputPeer.GetPeerType() {
	case constants.PeerTypeUser, constants.PeerTypeSelf:
		peerType = constants.PeerTypeUser
	case constants.PeerTypeChat, constants.PeerTypeChannel:
	default:
		peerType = constants.PeerTypeEmpty
	}

	conversationList, err := c.messageCore.GetConversationList(userId, excludePinned, peerType, offsetDate, int64(limit), foldId)
	if err != nil {
		log.Warnf("GetDialog userId:%d excludePinned:%v peerType:%v offsetDate:%d limit:%d error:%s", userId, excludePinned, peerType, offsetDate, limit, err.Error())
		return nil, err
	}

	log.Debugf("GetDialog userId:%d excludePinned:%v peerType:%v offsetDate:%d limit:%d len()=%d", userId, excludePinned, peerType, offsetDate, limit, len(conversationList))

	if len(conversationList) == 0 {
		return mtproto.NewTLMessagesDialogsSlice(nil).To_Messages_Dialogs(), nil
		//return mtproto.NewTLMessagesDialogsNotModified(&mtproto.Messages_Dialogs{Count: 0}).To_Messages_Dialogs(), nil
	}

	dialog, err := c.conversationToDialog(userId, conversationList, foldId, layer)
	if err != nil {
		log.Errorf("GetDialog userId:%d excludePinned:%v peerType:%v offsetDate:%d limit:%d GetMessageList error:%s", userId, excludePinned, peerType, offsetDate, limit, err.Error())
		return nil, err
	}

	if int32(len(conversationList)) == limit || limit == 0 {
		return mtproto.NewTLMessagesDialogsSlice(dialog).To_Messages_Dialogs(), nil
	}

	return mtproto.NewTLMessagesDialogs(dialog).To_Messages_Dialogs(), nil
}
