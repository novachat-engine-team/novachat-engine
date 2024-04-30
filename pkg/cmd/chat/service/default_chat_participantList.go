/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package service

import (
	data_chat "novachat_engine/service/data/chat"
	"time"
)

func DefaultChatParticipant() *data_chat.ChatParticipant {
	return &data_chat.ChatParticipant{
		ChatId:      0,
		UserId:      0,
		InviteId:    0,
		Date:        int32(time.Now().Unix()),
		Admin:       false,
		AdminRights: 0,
		State:       data_chat.ParticipantState_ParticipantStateNormal,
	}
}
