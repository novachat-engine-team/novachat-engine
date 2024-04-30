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
)

func (c *Core) GetPinnedDialog(userId int64, foldId int32, layer int32) (*mtproto.Messages_PeerDialogs, error) {

	log.Debugf("GetPinnedDialog, userId:%d foldId:%d", userId, foldId)

	dialog, err := c.GetDialog(userId, false, foldId, 0, 0, mtproto.NewTLInputPeerEmpty(nil).To_InputPeer(), 0, 0, layer)
	if err != nil {
		log.Errorf("GetPinnedDialog %v, userId:%d foldId:%d error:%s", userId, foldId, err.Error())
		return nil, err
	}

	log.Infof("GetPinnedDialog userId:%d foldId:%d dialog:%+v", userId, foldId, dialog)
	return mtproto.NewTLMessagesPeerDialogs(&mtproto.Messages_PeerDialogs{
		Dialogs:  dialog.Dialogs,
		Messages: dialog.Messages,
		Chats:    dialog.Chats,
		Users:    dialog.Users,
		State:    mtproto.NewTLUpdatesState(nil).To_Updates_State(),
	}).To_Messages_PeerDialogs(), nil
}
