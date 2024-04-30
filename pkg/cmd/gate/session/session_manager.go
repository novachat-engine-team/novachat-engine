/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/2/21 16:13
 * @File : session_manager.go
 */

package session

import (
	"sync"
)

const constSlot uint64 = 32

func getSlot(authKeyId int64) int64 {
	return int64(uint64(authKeyId) % constSlot)
}

type detail struct {
	sessionId map[int64]uint64
	//connID   uint64
}

type sessionData struct {
	sessionMap map[int64]*detail
}

type SessionManager struct {
	rwLock sync.RWMutex
	data   [constSlot]*sessionData
}

func NewSessionManager() *SessionManager {
	s := &SessionManager{}
	return s
}

// add new connId return old connId
func (s *SessionManager) AddSession(authKeyId, sessionId int64, connId uint64) uint64 {

	slotData := s.data[getSlot(authKeyId)]
	if slotData == nil {
		slotData = &sessionData{}
		s.data[getSlot(authKeyId)] = slotData
	}
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	if slotData.sessionMap == nil {
		slotData.sessionMap = make(map[int64]*detail)
	}
	detailData, ok := slotData.sessionMap[authKeyId]
	if ok == false {
		detailData = &detail{
			sessionId: make(map[int64]uint64),
		}
		slotData.sessionMap[authKeyId] = detailData
	}

	oldConnId, ok := detailData.sessionId[sessionId]
	if ok == true {
		detailData.sessionId[sessionId] = connId
		return oldConnId
	}

	detailData.sessionId[sessionId] = connId
	return 0
}

func (s *SessionManager) CloseSession(authKeyId int64) {
	slotData := s.data[getSlot(authKeyId)]
	if slotData == nil {
		return
	}

	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	delete(slotData.sessionMap, authKeyId)
}

func (s *SessionManager) GetSession(authKeyId, sessionId int64) uint64 {
	slotData := s.data[getSlot(authKeyId)]
	if slotData == nil {
		return 0
	}
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()

	detailData, ok := slotData.sessionMap[authKeyId]
	if ok == false {
		return 0
	}

	connId := detailData.sessionId[sessionId]
	return connId
}
