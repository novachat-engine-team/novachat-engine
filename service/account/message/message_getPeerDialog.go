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
	"novachat_engine/service/data/messages/message"
	"novachat_engine/service/input"
)

func (c *Core) GetPeerDialog(userId int64, inputPeerList []*mtproto.InputPeer, foldId int32, layer int32) (*mtproto.Messages_PeerDialogs, error) {
	log.Debugf("GetPeerDialog GetPeerDialog userId:%d", userId)

	peerConversationList := make([]*data_message.Conversation, 0, len(inputPeerList))
	for _, v := range inputPeerList {
		inputPeer := input.MakeInputPeer(v)

		peerType := inputPeer.GetPeerType()
		if peerType == constants.PeerTypeEmpty {
			continue
		}

		if peerType == constants.PeerTypeSelf {
			inputPeer = input.MakeInputPeerValue(userId, constants.PeerTypeUser)
		}

		peerConversationList = append(peerConversationList, &data_message.Conversation{
			UserId:   userId,
			PeerId:   inputPeer.GetPeerId(),
			PeerType: inputPeer.GetPeerType().ToInt32(),
			FolderId: foldId,
		})
	}

	conversationList, err := c.messageCore.GetConversationByPeerList(userId, peerConversationList)
	dialog, err := c.conversationToDialog(userId, conversationList, foldId, layer)
	if err != nil {
		log.Debugf("GetPeerDialog GetPeerDialog userId:%d error:%s", userId, err.Error())
		return nil, err
	}

	log.Infof("GetPeerDialog GetPeerDialog userId:%d dialog:%+v", userId, dialog)
	return mtproto.NewTLMessagesPeerDialogs(&mtproto.Messages_PeerDialogs{
		Dialogs:  dialog.Dialogs,
		Messages: dialog.Messages,
		Chats:    dialog.Chats,
		Users:    dialog.Users,
	}).To_Messages_PeerDialogs(), nil
}
