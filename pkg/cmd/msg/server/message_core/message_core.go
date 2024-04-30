/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/11 10:48
 * @File : message_core.go
 */

package message_core

import (
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
)

func (m *MessageCoreService) SendMessageData(r *msgService.SendMessages) error {
	switch constants.PeerTypeFromInt32(r.PeerType) {
	case constants.PeerTypeSelf, constants.PeerTypeUser:
		return m.sendInboxUserMessages(r.FromUserId, r.PeerId, r.DataList, r.GlobalMessageIdList)
	default:
		log.Errorf("SendMessageData peerType:%+v", r.PeerType)
		return nil
	}
}

func (m *MessageCoreService) SendMessageDataList(r []*msgService.SendMessages) error {
	if len(r) == 0 {
		log.Warnf("SendChatMessageData len(value) == 0")
		return nil
	}
	switch constants.PeerTypeFromInt32(r[0].PeerType) {
	case constants.PeerTypeChat, constants.PeerTypeChannel:
		var err error
		for idx, v := range r {
			err = m.sendInboxChatMessages(v.FromUserId, v.PeerId, v.DataList, v.GlobalMessageIdList)
			if err != nil {
				log.Errorf("SendMessageDataList idx:%d v:%+v error:%s", idx, v, err.Error())
				return err
			}
		}
		return nil
	default:
		log.Errorf("SendChatMessageData peerType:%+v", r[0].PeerType)
		return nil
	}
}
