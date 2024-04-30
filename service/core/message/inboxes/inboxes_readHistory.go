/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package inboxes

import (
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message"
)

func (c *Core) ReadHistoryMessageData(userId int64, peerId int64, peerType constants.PeerType, globalMessageId int64, now int32) (int32, int32, error) {
	messageDataList, err := c.messageCore.GetMessageByGlobalMessageIdList(userId, []int64{globalMessageId}, false)
	if err != nil {
		log.Errorf("ReadHistoryMessageData error:%s", err.Error())
		return 0, 0, err
	}
	if len(messageDataList) == 0 {
		log.Warnf("ReadHistoryMessageData userId:%d peerId:%d peerType:%v globalMessageId messageDataList is empty", userId, peerId, peerType, globalMessageId)
		return 0, 0, nil
	}

	pts, folderId, err := c.messageCore.ReadHistory(userId, peerId, peerType, messageDataList[0].Id, now, message.InboxType)
	if err != nil {
		log.Errorf("ReadHistoryMessageData ReadHistory error:%s", err.Error())
		return 0, 0, err
	}

	log.Debugf("ReadHistoryMessageData userId:%d peerId:%d peerType:%v globalMessageId:%v messageId:%d pts:%d folderId:%d", userId, peerId, peerType, globalMessageId, messageDataList[0].Id, pts, folderId)
	return pts, messageDataList[0].Id, nil
}
