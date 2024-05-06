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
	"math"
	"novachat_engine/service/constants"
)

const (
	DBMessage              = "db_messages"
	TableMessage           = "tb_messages"
	TableMessageBox        = "tb_messages_box"
	TableConversation      = "tb_conversations"
	TableChannelMessage    = "tb_channel_messages"
	TableChannelMessageBox = "tb_channel_messages_box"
)

type BoxType int32

const (
	InboxType  BoxType = 0
	OutboxType BoxType = 1
)

const (
	MaxTableLimit = 256
)

var tableLimit = 0

func UpdateTableLimit(v int) {
	if v <= 0 {
		v = 1
	}
	tableLimit = v
}

func GetTableLimit() int {
	if tableLimit <= 0 {
		return 1
	}
	return tableLimit
}

func TableName(userId int64, tableName string) string {
	if userId < 0 {
		userId = -userId
	}

	idx := userId % int64(math.Min(float64(GetTableLimit()), MaxTableLimit))
	if idx == 0 {
		return tableName
	}
	return fmt.Sprintf("%s_%d", tableName, idx)
}

func MakePeerType(peerId int64) (int64, constants.PeerType) {
	if peerId < 0 {
		return -peerId, constants.PeerTypeChannel
	}
	return peerId, constants.PeerTypeUser
}

func MakePeerId(peerId int64, peerType constants.PeerType) int64 {
	if peerType == constants.PeerTypeChat ||
		peerType == constants.PeerTypeChannel {
		return -peerId
	}
	return peerId
}

func MakeMessageUserId(userId int64, boxType BoxType) int64 {
	if boxType == OutboxType {
		return userId
	}
	return -userId
}

func MakeDialogId(userId, peerId int64, peerType constants.PeerType) string {
	return fmt.Sprintf("%d%d", peerId, userId)
}
