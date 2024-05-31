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
	"fmt"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	chatCore "novachat_engine/service/core/chat"
	data_chat "novachat_engine/service/data/chat"
	data_fs "novachat_engine/service/data/fs"
	"time"

	"github.com/dgraph-io/ristretto"
)

const ttl = 5
const commonChatTTL = 10

var emptyChatCore = &chatCore.Core{}

type ChatManager struct {
	chatCache    *ristretto.Cache
	chatCore     *chatCore.Core
	commonChat   *ChatCommon
	conf         *config.MongodbConfig
	chatMaxLimit int32
}

func NewChatManager(conf *config.MongodbConfig) *ChatManager {
	var err error
	mgr := &ChatManager{
		conf: conf,
	}
	mgr.chatCache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 10,
		BufferItems: 64,
	})
	if err != nil {
		panic(fmt.Sprintf("NewChatManager cache %s", err.Error()))
	}

	mgr.chatCore = chatCore.NewChatCore(conf)
	commonChat, err := NewChatCommon(chatCore.NewChatCommonCore(conf), mgr.chatCore)
	if err != nil {
		panic(fmt.Sprintf("NewChatCommon %s", err.Error()))
	}

	mgr.commonChat = commonChat
	mgr.chatMaxLimit = 200
	return mgr
}

func (m *ChatManager) GetChat(chatId int64) (*Chat, error) {
	if chatId <= 0 {
		return nil, &ErrChatIdInvalid{chatId: chatId}
	}

	var ok bool
	var err error
	key := makeChatKey(chatId)
	var chat *Chat
	chatCache, ok := m.chatCache.Get(key)
	if !ok || chatCache == nil {
		chat, err = NewChat(chatId, m.chatCore, m.chatMaxLimit)
		if err != nil {
			return nil, err
		}
		if chat.Invalid() {
			ok = m.chatCache.SetWithTTL(key, chat, 1, ttl*time.Minute)
			if !ok {
				log.Warnf("GetChat Update Cache chatId:%d false", chatId)
			} else {
				log.Debugf("GetChat loadChat chatId:%d", chatId)
			}
		}
	} else {
		chat = chatCache.(*Chat)
	}

	return chat, nil
}

func (m *ChatManager) GetCommonChat(userId int64) ([]*data_chat.Chat, error) {
	if userId <= 0 {
		return nil, &ErrUserIdInvalid{userId: userId}
	}

	return m.commonChat.getChatCommon(userId)
}

func (m *ChatManager) CreateChat(creatorUserId int64, userIdList []int64, title string, date int32, point *data_fs.GeoPoint, address string) (*Chat, error) {

	var chatParticipantList []*data_chat.ChatParticipant
	chatData := DefaultChatData(date)
	chatData.Creator = creatorUserId
	chatData.Title = title
	chatInfo := NewChatInfo(chatData)
	chatInfo.Count = 0
	chatInfo.ChatData.Address = address
	chatInfo.ChatData.GeoPoint = point

	for _, v := range userIdList {
		cp := DefaultChatParticipant()
		cp.UserId = v
		cp.Date = chatData.Date
		cp.InviteId = creatorUserId
		cp.Admin = v == chatData.Creator

		chatInfo.Count += 1
		chatParticipantList = append(chatParticipantList, cp)
	}

	chatData, err := m.chatCore.CreateChat(chatData, chatParticipantList)
	if err != nil {
		log.Errorf("CreateChat error:%s", err.Error())
		return nil, err
	}

	log.Debugf("CreateChat creatorUserId:%d chatId:%d", creatorUserId, chatData.ChatId)
	return &Chat{chatInfo: chatInfo}, nil
}

func (m *ChatManager) GetChatByName(username string) (*data_chat.Chat, error) {
	return m.chatCore.GetChatByName(username)
}

func (m *ChatManager) GetChats(idList, exceptIds []int64) ([]*Chat, error) {
	coreChatList, err := m.chatCore.GetChatList(chatCore.GetChatListWithIds(idList), chatCore.GetChatListWithExceptIds(exceptIds))
	if err != nil {
		return nil, err
	}
	chatList := make([]*Chat, 0, len(coreChatList))
	for _, v := range coreChatList {
		chat, err := NewChat(v.ChatId, emptyChatCore, 0)
		if err != nil {
			log.Errorf("GetChats chatId:%d NewChat error:%s", v.ChatId, err.Error())
			return nil, err
		}

		chatParticipantList, err := m.chatCore.GetChatParticipants(v.ChatId, nil)
		if err != nil {
			log.Errorf("GetChats chatId:%d Participants error:%s", v.ChatId, err.Error())
			return nil, err
		}

		for _, vv := range chatParticipantList {
			chat.chatInfo.AddParticipants(vv.UserId, vv)
		}
		chatList = append(chatList, chat)
	}
	return chatList, nil
}
