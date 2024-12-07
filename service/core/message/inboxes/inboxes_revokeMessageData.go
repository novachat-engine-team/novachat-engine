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
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	data_message "novachat_engine/service/data/messages/message"
	"time"
)

func (c *Core) RevokeMessageData(userId int64, peerId int64, peerType constants.PeerType, globalMessageIdList []int64, rRange bool, all bool) (int32, []int32, error) {
	conversation, err := c.messageCore.GetConversation(userId, &data_message.Conversation{
		UserId:   userId,
		PeerId:   peerId,
		PeerType: peerType.ToInt32(),
	})
	if err != nil {
		log.Errorf("RevokeMessageData Conversation error:%s", err.Error())
		return 0, nil, err
	}

	if rRange == true {
		var messageMaxId int32
		if len(globalMessageIdList) == 0 {
			messageMaxId = conversation.Top
		} else {
			var messageDataList []*data_message.Message
			messageDataList, err = c.messageCore.GetMessageByGlobalMessageIdList(userId, globalMessageIdList, false)
			if err != nil {
				log.Errorf("RevokeMessageData GetMessageByGlobalMessageIdList error:%s", err.Error())
				return 0, nil, err
			}
			if len(messageDataList) == 0 {
				messageMaxId = messageDataList[0].Id
			} else {
				messageMaxId = conversation.Top
			}
		}
		log.Debugf("RevokeMessageData userId:%d peerId:%d peerType:%v globalMessageIdList:%v messageMaxId:%d", userId, peerId, peerType, globalMessageIdList, messageMaxId)
		return c.messageCore.DeleteMessagesByMaxId(userId, peerId, peerType, messageMaxId, int32(time.Now().Unix()), 0)
	} else {
		var pts int32
		var messageIdList []int32
		if !all {
			var messageDataList []*data_message.Message
			if len(globalMessageIdList) == 0 {
				log.Warnf("RevokeMessageData userId:%d peerId:%d peerType:%v globalMessageId is empty", userId, peerId, peerType)
				return 0, nil, nil
			}
			log.Debugf("RevokeMessageData  DeleteMessages peerId :%d peerType:%v globalMessageIdList:%v", peerId, peerType, globalMessageIdList)
			messageDataList, err = c.messageCore.GetMessageByGlobalMessageIdList(userId, globalMessageIdList, false)
			if err != nil {
				log.Errorf("RevokeMessageData GetMessageByGlobalMessageIdList error:%s", err.Error())
				return 0, nil, err
			}
			if len(messageDataList) == 0 {
				log.Warnf("RevokeMessageData GetMessageByGlobalMessageIdList userId:%d globalMessageIdList:%v message is empty", userId, globalMessageIdList)
				return 0, nil, nil
			}
			messageIdList = make([]int32, 0, len(messageDataList))
			for _, v := range messageDataList {
				messageIdList = append(messageIdList, v.Id)
			}
			pts, err = c.messageCore.DeleteMessages(userId, messageIdList)
			if err != nil {
				err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
				log.Errorf("RevokeMessageData DeleteMessages error:%s need retry", err.Error())
				return 0, nil, err
			}
		} else {
			pts, messageIdList, err = c.messageCore.DeleteMessagesByMaxId(userId, peerId, peerType, conversation.Top, 0, 0)
			if err != nil {
				log.Errorf("RevokeMessageData DeleteMessagesByMaxId error:%s", err.Error())
				return 0, nil, err
			}
		}

		log.Debugf("RevokeMessageData DeleteMessages userId:%d peerId:%d peerType:%v globalMessageIdList:%v messageIdList:%v all:%+v", userId, peerId, peerType, globalMessageIdList, messageIdList, all)
		return pts, messageIdList, nil
	}
}

func (c *Core) RevokeChannelMessageData(userId int64, peerId int64, peerType constants.PeerType, pts int32) (bool, error) {
	conversation, err := c.messageCore.GetConversation(userId, &data_message.Conversation{
		UserId:   userId,
		PeerId:   peerId,
		PeerType: peerType.ToInt32(),
	})
	if err != nil {
		log.Errorf("RevokeChannelMessageData Conversation error:%s", err.Error())
		return false, err
	}
	if conversation.Pts >= pts {
		return true, nil
	}

	err = c.messageCore.ConversationUpdatePts(context.Background(), userId, peerId, peerType, pts)
	if err != nil {
		log.Errorf("RevokeChannelMessageData Conversation error:%s", err.Error())
		return false, err
	}

	return false, nil
}
