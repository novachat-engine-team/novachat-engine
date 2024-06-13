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
	"context"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
)

func (c *Core) DeleteChannelMessages(chatId int64, userId int64, messageIdList []int32) (int32, error) {
	pts, err := c.messageCore.DeleteChannelMessages(chatId, userId, messageIdList)
	if err != nil {
		log.Errorf("DeleteChannelMessages error:%s", err.Error())
		return 0, err
	}

	err = c.messageCore.ConversationUpdatePts(context.Background(), userId, chatId, constants.PeerTypeChannel, pts)
	if err != nil {
		log.Warnf("DeleteChannelMessages ConversationUpdatePts error:%s", err.Error())
	}
	return pts, nil
}

func (c *Core) DeleteChannelUserMessages(chatId int64, userId int64) (int32, []int32, error) {
	var messageIdList []int32
	messageList, err := c.messageCore.GetChannelMessageListByUserId(chatId, userId)
	if err != nil {
		log.Errorf("DeleteChannelUserMessages GetChannelMessageListByUserId error:%s", err.Error())
		return 0, nil, err
	}
	if len(messageList) == 0 {
		return 0, nil, nil
	}
	messageIdList = make([]int32, 0, len(messageList))
	for _, v := range messageList {
		messageIdList = append(messageIdList, v.Id)
	}

	pts, err := c.messageCore.DeleteChannelMessages(chatId, userId, messageIdList)
	if err != nil {
		log.Errorf("DeleteChannelUserMessages DeleteChannelMessages error:%s", err.Error())
		return 0, nil, err
	}

	err = c.messageCore.ConversationUpdatePts(context.Background(), userId, chatId, constants.PeerTypeChannel, pts)
	if err != nil {
		log.Warnf("DeleteChannelUserMessages ConversationUpdatePts error:%s", err.Error())
	}
	return pts, messageIdList, nil
}

func (c *Core) DeleteChannelHistory(chatId int64, userId int64, maxId int32, date int32) (int32, []int32, error) {
	var messageIdList []int32
	messageList, err := c.messageCore.GetChannelMessageListByMaxId(chatId, maxId)
	if err != nil {
		log.Errorf("DeleteChannelHistory GetChannelMessageListByMaxId error:%s", err.Error())
		return 0, nil, err
	}
	if len(messageList) == 0 {
		return 0, nil, nil
	}
	messageIdList = make([]int32, 0, len(messageList))
	for _, v := range messageList {
		messageIdList = append(messageIdList, v.Id)
	}

	pts, err := c.messageCore.DeleteChannelMessages(chatId, userId, messageIdList)
	if err != nil {
		log.Errorf("DeleteChannelHistory DeleteChannelMessages error:%s", err.Error())
		return 0, nil, err
	}

	err = c.messageCore.ConversationUpdatePts(context.Background(), userId, chatId, constants.PeerTypeChannel, pts)
	if err != nil {
		log.Warnf("DeleteChannelHistory ConversationUpdatePts error:%s", err.Error())
	}
	return pts, messageIdList, nil
}
