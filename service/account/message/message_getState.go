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
)

// GetState
// Updates_State init pts/qts/seq then more ...
func (c *Core) GetState(userId int64) (*mtproto.Updates_State, error) {
	conversationList, err := c.messageCore.GetConversationList(userId, true, constants.PeerTypeEmpty, 0, 0, 0)
	if err != nil {
		log.Errorf("GetState userId:%d error:%s", userId, err.Error())
		return nil, err
	}

	updatesState := mtproto.Updates_State{
		Pts:         1,
		Qts:         1,
		Date:        0,
		Seq:         0,
		UnreadCount: 0,
	}
	for _, v := range conversationList {
		if v.PeerType == constants.PeerTypeChannel.ToInt32() {
			continue
		}

		if v.FolderId == 0 {
			updatesState.UnreadCount += v.UnreadCount
		}

		if updatesState.Pts < v.Pts {
			updatesState.Pts = v.Pts
			updatesState.Date = v.Date
		}
	}

	return mtproto.NewTLUpdatesState(&updatesState).To_Updates_State(), nil
}
