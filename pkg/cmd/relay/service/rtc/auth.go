/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package rtc

import (
	"github.com/pion/turn/v2"
	"net"
	"novachat_engine/pkg/cmd/relay/conf"
	"novachat_engine/pkg/log"
	"sync"
)

var userPermMutex sync.RWMutex
var userPerm map[string][]byte

func init() {
	userPerm = map[string][]byte{}
}

func (s *Service) RegisterAuth(conf *conf.RTC) turn.AuthHandler {
	userPermMutex.Lock()
	defer userPermMutex.Unlock()

	for _, v := range conf.Realm {
		if len(v) < 2 || len(v[0]) == 0 || len(v[1]) == 0 {
			log.Warnf("conf.Realm:%s", v)
		} else {
			userPerm[v[0]] = turn.GenerateAuthKey(v[0], conf.RealmTag, v[1])
		}
	}
	return s.relayAuth
}

func (s *Service) relayAuth(username, realm string, srcAddr net.Addr) ([]byte, bool) {
	userPermMutex.RLock()
	defer userPermMutex.RUnlock()

	if key, ok := userPerm[username]; ok {
		return key, ok
	}

	return nil, false
}
