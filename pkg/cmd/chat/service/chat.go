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
	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/log"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/core/chat"
	chatCore "novachat_engine/service/core/chat"
	data_chat "novachat_engine/service/data/chat"
	data_fs "novachat_engine/service/data/fs"
	"sync"
)

type ChatInfo struct {
	ChatData         *data_chat.Chat
	Count            int32
	AdminCount       int32
	chatParticipants sync.Map
}

func NewChatInfo(chatData *data_chat.Chat) *ChatInfo {
	return &ChatInfo{ChatData: chatData}
}

func (m *ChatInfo) AddParticipants(userId int64, p *data_chat.ChatParticipant) {
	m.chatParticipants.Store(userId, p)
}

func (m *ChatInfo) GetParticipants(userId int64) *data_chat.ChatParticipant {
	v, ok := m.chatParticipants.Load(userId)
	if !ok {
		return nil
	}
	return v.(*data_chat.ChatParticipant)
}

func (m *ChatInfo) Iteration(f func(participant *data_chat.ChatParticipant)) {
	m.chatParticipants.Range(func(key, value interface{}) bool {
		if value != nil {
			p, ok := value.(*data_chat.ChatParticipant)
			if !ok {
				_ = p
			} else {
				f(p)
			}
		}
		return true
	})
}

func (m *ChatInfo) IterationCondition(f func(participant *data_chat.ChatParticipant) bool) {
	m.chatParticipants.Range(func(key, value interface{}) bool {
		if value != nil {
			return f(value.(*data_chat.ChatParticipant))
		}
		return true
	})
}

type Chat struct {
	chatCore    *chat.Core
	chatId      int64
	redisClient *cache.RedisClient
	chatInfo    *ChatInfo
	chatMax     int32
}

func NewChat(chatId int64, chatCore *chatCore.Core, chatMax int32) (*Chat, error) {

	if chatMax <= 0 {
		chatMax = defaultChatMax
	}

	chat := &Chat{
		chatId:      chatId,
		chatCore:    chatCore,
		redisClient: cache.GetRedisClient(),
		//chatParticipant: make(map[int64]*data_chat.ChatParticipant, chatMax),
	}
	chat.chatMax = chatMax

	if err := chat.getChat(); err != nil {
		log.Errorf("NewChat error:%s", err.Error())
		return nil, err
	}

	return chat, nil
}

func (m *Chat) getChat() error {
	log.Debugf("getChat chatId:%d", m.chatId)

	chatInfo := &ChatInfo{
		ChatData:   &data_chat.Chat{},
		Count:      0,
		AdminCount: 0,
	}
	var chatData *data_chat.Chat
	info, err := redis.String(m.redisClient.Do("GET", makeChatKey(m.chatId)))
	if err != nil {
		log.Warnf("getChat error:%s", err.Error())
	} else {
		err = jsoniter.Unmarshal([]byte(info), &chatData)
	}
	if err != nil || len(info) == 0 {
		chatData, err = m.chatCore.
			GetChat(m.chatId)
		if err != nil {
			log.Errorf("getChat chatId:%d error:%s", m.chatId, err.Error())
			return err
		}

		if chatData.ChatId != m.chatId {
			m.chatInfo = chatInfo
			return nil
		}
	}
	chatInfo.ChatData = chatData

	if !chatData.Deleted {
		var chatParticipantList []*data_chat.ChatParticipant
		chatParticipantList, err = m.chatCore.GetChatParticipants(m.chatId, nil)
		if err != nil {
			log.Errorf("getChat chatId:%d Participants error:%s", m.chatId, err.Error())
			return err
		}

		for _, v := range chatParticipantList {
			chatInfo.Count++
			if v.Admin {
				chatInfo.AdminCount++
			}
			chatInfo.AddParticipants(v.UserId, v)
		}
	}

	m.chatInfo = chatInfo
	return nil
}

func (m *Chat) GetChatInfo() *ChatInfo {
	return m.chatInfo
}

func (m *Chat) GetChatMax() int32 {
	return m.chatMax
}

func (m *Chat) Invalid() bool {
	return m.chatInfo == nil || m.chatInfo.ChatData == nil || m.chatInfo.ChatData.ChatId == 0
}

func (m *Chat) Deleted() bool {
	if m.Invalid() {
		return true
	}

	return m.chatInfo != nil && m.chatInfo.ChatData != nil && m.chatInfo.ChatData.Deleted
}

func (m *Chat) EditPhoto(photo *data_fs.PhotoProfile) (*ChatInfo, error) {
	ok, err := m.chatCore.EditProperty(m.chatId, data_chat.Chat{ChatId: m.chatId, Photo: photo}, "Photo")
	if err != nil {
		log.Errorf("EditPhoto chatId:%d error:%s", m.chatId, err.Error())
		return nil, err
	}

	_ = ok
	//if m.chatInfo.ChatData.Photo == nil {
	//	m.chatInfo.ChatData.Photo = &data_chat.ChatPhoto{}
	//}
	m.chatInfo.ChatData.Photo = photo
	return m.chatInfo, nil
}

func (m *Chat) EditTitle(title string) (*ChatInfo, bool, error) {
	if m.chatInfo.ChatData.Title == title {
		return m.chatInfo, false, nil
	}

	ok, err := m.chatCore.EditProperty(m.chatId, data_chat.Chat{ChatId: m.chatId, Title: title}, "Title")
	if err != nil {
		log.Errorf("EditTitle chatId:%d error:%s", m.chatId, err.Error())
		return nil, false, err
	}

	_ = ok
	m.chatInfo.ChatData.Title = title
	return m.chatInfo, true, nil
}

func (m *Chat) EditGeoPoint(point *data_fs.GeoPoint, address string) (*ChatInfo, error) {

	ok, err := m.chatCore.EditProperty(m.chatId, data_chat.Chat{ChatId: m.chatId, GeoPoint: point, Address: address}, "GeoPoint", "Address")
	if err != nil {
		log.Errorf("EditGeoPoint chatId:%d error:%s", m.chatId, err.Error())
		return nil, err
	}

	_ = ok
	m.chatInfo.ChatData.GeoPoint = point
	return m.chatInfo, nil
}

func (m *Chat) EditAbout(about string) (*ChatInfo, bool, error) {
	if m.chatInfo.ChatData.About == about {
		return m.chatInfo, false, nil
	}

	ok, err := m.chatCore.EditProperty(m.chatId, data_chat.Chat{ChatId: m.chatId, About: about}, "About")
	if err != nil {
		log.Errorf("EditAbout chatId:%d error:%s", m.chatId, err.Error())
		return nil, false, err
	}

	_ = ok
	m.chatInfo.ChatData.About = about
	return m.chatInfo, true, nil
}

func (m *Chat) AddUsers(userId int64, userIdList []int64) ([]*data_chat.ChatParticipant, []*data_chat.ChatParticipant, error) {

	participantValue, ok := m.chatInfo.chatParticipants.Load(userId)
	if !ok || participantValue.(*data_chat.ChatParticipant).State == data_chat.ParticipantState_ParticipantStateNormal {
		if participantValue != nil {
			log.Warnf("AddUsers chatId:%d userId:%d State:%v", m.chatId, userId, participantValue.(*data_chat.ChatParticipant).State)
		} else {
			log.Warnf("AddUsers chatId:%d userId:%d not in", m.chatId, userId)
		}
		return nil, nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_WRITE_FORBIDDEN)
	}

	modifyUserList := make([]*data_chat.ChatParticipant, 0, len(userIdList))
	var participant *data_chat.ChatParticipant
	newUserIdList := make([]*data_chat.ChatParticipant, 0, len(userIdList))
	for _, v := range userIdList {
		participantValue, ok = m.chatInfo.chatParticipants.Load(v)
		if !ok {
			participant = DefaultChatParticipant()
			participant.ChatId = m.chatId
			participant.InviteId = userId
			participant.UserId = userId
			newUserIdList = append(newUserIdList, participant)
		} else {
			switch participantValue.(*data_chat.ChatParticipant).State {
			case data_chat.ParticipantState_ParticipantStateLeft:
				modifyUserList = append(modifyUserList, participant)
			case data_chat.ParticipantState_ParticipantStateBan:
				fallthrough
			case data_chat.ParticipantState_ParticipantStateDelete:
				participant = DefaultChatParticipant()
				participant.ChatId = m.chatId
				participant.InviteId = userId
				participant.UserId = userId
				modifyUserList = append(modifyUserList, participant)
			case data_chat.ParticipantState_ParticipantStateNormal:
				fallthrough
			default:
				break
			}
		}
	}

	if len(newUserIdList) == 0 && len(modifyUserList) == 0 {
		return nil, nil, nil
	}

	ok, err := m.chatCore.ModifyParticipant(m.chatId, newUserIdList, modifyUserList)
	if err != nil {
		log.Errorf("AddUsers ModifyParticipant chatId:%d error:%s", m.chatId, err.Error())
		return nil, nil, err
	}

	for _, v := range newUserIdList {
		m.chatInfo.AddParticipants(v.UserId, v)
	}

	for _, v := range modifyUserList {
		m.chatInfo.AddParticipants(v.UserId, v)
	}

	m.chatInfo.Count += int32(len(newUserIdList) + len(modifyUserList))
	return newUserIdList, modifyUserList, nil
}

func (m *Chat) ModifyParticipant(userId int64, userIdList []int64, reason data_chat.ParticipantState, rights int32) ([]*data_chat.ChatParticipant, error) {

	participantValue, ok := m.chatInfo.chatParticipants.Load(userId)
	if !ok || participantValue.(*data_chat.ChatParticipant).State == data_chat.ParticipantState_ParticipantStateNormal {
		if participantValue != nil {
			log.Warnf("ModifyParticipant chatId:%d userId:%d State:%v", m.chatId, userId, participantValue.(*data_chat.ChatParticipant).State)
		} else {
			log.Warnf("ModifyParticipant chatId:%d userId:%d not in", m.chatId, userId)
		}
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_WRITE_FORBIDDEN)
	}

	var removeCount int32
	var adminRemoveCount int32
	modifyUserList := make([]*data_chat.ChatParticipant, 0, len(userIdList))

	var participant *data_chat.ChatParticipant
	for _, v := range userIdList {
		participantValue, ok = m.chatInfo.chatParticipants.Load(v)
		if !ok {
			log.Warnf("ModifyParticipant chatId:%d userId:%d not found", m.chatId, v)
			continue
		} else {
			participant = participantValue.(*data_chat.ChatParticipant)
			switch participant.State {
			case data_chat.ParticipantState_ParticipantStateLeft:
			case data_chat.ParticipantState_ParticipantStateDelete:
			case data_chat.ParticipantState_ParticipantStateBan:
				if reason == data_chat.ParticipantState_ParticipantStateDelete {
					participant.State = reason
					modifyUserList = append(modifyUserList, participant)
				}
			case data_chat.ParticipantState_ParticipantStateNormal:
				fallthrough
			default:
				switch reason {
				case data_chat.ParticipantState_ParticipantStateLeft:
					if participant.Admin {
						adminRemoveCount++
					}
					removeCount++
				case data_chat.ParticipantState_ParticipantStateDelete:
					if participant.Admin {
						adminRemoveCount++
					}
					removeCount++
					fallthrough
				case data_chat.ParticipantState_ParticipantStateBan:
					fallthrough
				case data_chat.ParticipantState_ParticipantStateNormal:
					fallthrough
				default:
					participant.Rights = rights
				}
				participant.State = reason
				modifyUserList = append(modifyUserList, participant)
				break
			}
		}
	}
	ok, err := m.chatCore.ModifyParticipant(m.chatId, nil, modifyUserList)
	if err != nil {
		log.Errorf("ModifyParticipant ModifyParticipant chatId:%d error:%s", m.chatId, err.Error())
		return nil, err
	}
	for _, v := range modifyUserList {
		m.chatInfo.AddParticipants(v.UserId, v)
	}
	m.chatInfo.Count -= removeCount
	if m.chatInfo.Count < 0 {
		log.Fatalf("ModifyParticipant ModifyParticipant chatId:%d Count:%s", m.chatId, m.chatInfo.Count)
		m.chatInfo.Count = 0
	}
	m.chatInfo.AdminCount -= adminRemoveCount
	if m.chatInfo.AdminCount < 0 {
		log.Fatalf("ModifyParticipant ModifyParticipant chatId:%d AdminCount:%s", m.chatId, m.chatInfo.AdminCount)
		m.chatInfo.AdminCount = 0
	}
	return modifyUserList, nil
}

func (m *Chat) EditAdmin(userId int64, peerId int64, isAdmin bool) (*data_chat.ChatParticipant, error) {
	participantValue, ok := m.chatInfo.chatParticipants.Load(userId)
	if !ok || participantValue.(*data_chat.ChatParticipant).State == data_chat.ParticipantState_ParticipantStateNormal {
		if participantValue != nil {
			log.Warnf("EditAdmin chatId:%d userId:%d State:%v", m.chatId, userId, participantValue.(*data_chat.ChatParticipant).State)
		} else {
			log.Warnf("EditAdmin chatId:%d userId:%d not in", m.chatId, userId)
		}
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_WRITE_FORBIDDEN)
	}

	participant := participantValue.(*data_chat.ChatParticipant)
	if !participant.Admin {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
	}

	participantValue, ok = m.chatInfo.chatParticipants.Load(peerId)
	if !ok {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
	}

	participant = &data_chat.ChatParticipant{}
	*participant = *participantValue.(*data_chat.ChatParticipant)

	participant.Admin = isAdmin
	ok, err := m.chatCore.ModifyParticipant(m.chatId, nil, []*data_chat.ChatParticipant{participant})
	if err != nil {
		log.Errorf("EditAdmin ModifyParticipant chatId:%d error:%s", m.chatId, err.Error())
		return nil, err
	}

	if participantValue.(*data_chat.ChatParticipant).Admin != isAdmin {
		if isAdmin {
			m.chatInfo.AdminCount++
		} else {
			m.chatInfo.AdminCount--
		}
	}
	m.chatInfo.AddParticipants(peerId, participant)
	return participant, nil
}

func (m *Chat) EditBannedRights(rights int32) (*ChatInfo, error) {

	ok, err := m.chatCore.EditProperty(m.chatId, data_chat.Chat{
		ChatId:       m.chatId,
		BannedRights: rights,
	}, "BannedRights")
	if err != nil {
		log.Errorf("EditBannedRights error:%s", err.Error())
		return nil, err
	}

	_ = ok
	m.chatInfo.ChatData.BannedRights = rights
	return m.chatInfo, nil
}

func (m *Chat) EditUsername(username string) (*ChatInfo, bool, error) {
	if m.chatInfo.ChatData.Username != username {
		return m.chatInfo, false, nil
	}

	ok, err := m.chatCore.EditProperty(m.chatId, data_chat.Chat{ChatId: m.chatId, Username: username}, "Username")
	if err != nil {
		log.Errorf("EditUsername chatId:%d error:%s", m.chatId, err.Error())
		return nil, false, err
	}

	_ = ok
	m.chatInfo.ChatData.Username = username
	return m.chatInfo, true, nil
}

func (m *Chat) EditHiddenHistory(hiddenHistory bool) (*ChatInfo, bool, error) {
	ok, err := m.chatCore.EditProperty(m.chatId, data_chat.Chat{ChatId: m.chatId, HiddenHistory: hiddenHistory}, "HiddenHistory")
	if err != nil {
		log.Errorf("EditHiddenHistory chatId:%d error:%s", m.chatId, err.Error())
		return nil, false, err
	}

	_ = ok
	m.chatInfo.ChatData.HiddenHistory = hiddenHistory
	return m.chatInfo, true, nil
}

func (m *Chat) EditDeleted() (*ChatInfo, bool, error) {
	ok, err := m.chatCore.EditProperty(m.chatId, data_chat.Chat{ChatId: m.chatId, Deleted: true}, "Deleted")
	if err != nil {
		log.Errorf("EditHiddenHistory chatId:%d error:%s", m.chatId, err.Error())
		return nil, false, err
	}

	_ = ok
	m.chatInfo.ChatData.Deleted = true
	return m.chatInfo, true, nil
}
