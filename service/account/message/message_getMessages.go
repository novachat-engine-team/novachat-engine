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
	message2 "novachat_engine/service/core/message"
	"novachat_engine/service/core/message/message"
)

func (c *Core) GetMessages(userId int64, messageIdList []int32, replyTo bool) ([]*mtproto.Message, error) {
	log.Debugf("GetMessages userId:%d messageIdList:%v", userId, messageIdList)

	messageDataList, err := c.messageCore.GetMessageList(userId, messageIdList)
	if err != nil {
		log.Errorf("GetMessages userId:%d error:%s", userId, err.Error())
		return nil, err
	}

	messageList := make([]*mtproto.Message, 0, len(messageDataList))
	for _, v := range messageDataList {
		_, peerType := message2.MakePeerType(v.PeerId)
		if !v.Out && peerType == constants.PeerTypeUser {
			v.PeerId = v.FromUserId
		}
		messageList = append(messageList, message.ToMessage(v))
	}

	log.Infof("GetMessages userId:%d messageIdList:%v len()=%d", userId, messageIdList, len(messageList))
	return messageList, nil
}