/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package handler

import (
	"novachat_engine/pkg/cmd/session/conf"
	sessionService "novachat_engine/pkg/cmd/session/rpc_client"
)

const defaultChannelCount = 128

type MessageChannel struct {
	channels [defaultChannelCount]*Session
}

func NewMessageChannel(h *SessionHandler, conf *conf.Conf) *MessageChannel {
	q := &MessageChannel{}
	for idx := range q.channels {
		q.channels[idx] = NewSession(h, conf)
	}
	h.SetMessageChannel(q)
	return q
}

func (q *MessageChannel) GetChannel(authKeyId int64) *Session {
	return q.channels[int64(uint64(authKeyId)%defaultChannelCount)]
}

func (q *MessageChannel) CloseSession(e *sessionService.CloseEvent) error {
	return q.GetChannel(e.AuthKeyId).CloseSession(e)
}
