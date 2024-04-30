/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/22 23:11
 * @File : connection_manager.go
 */

package net

import "sync"

type Manager struct {
	rwLock sync.RWMutex
	connectMap map[uint64]Connection

	close    sync.Once
}

func NewManager() *Manager {
	return &Manager{
		connectMap: make(map[uint64]Connection),
	}
}

func (m *Manager) connection(id uint64) Connection {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	return m.connectMap[id]
}

func (m *Manager) addConnection(conn Connection) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	m.connectMap[conn.GetConnID()] = conn
}

func (m *Manager) delConnection(conn Connection) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	delete(m.connectMap, conn.GetConnID())
}

func (m *Manager) Close() {
	m.close.Do(func() {
		for _,v := range m.connectMap {
			v.Close()
		}
	})
}



