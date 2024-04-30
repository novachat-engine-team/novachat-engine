/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package message_core

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"novachat_engine/mtproto"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"time"
)

func (m *MessageCoreService) ReadHistoryMessage(readHistory *msgService.ReadHistory) error {
	switch constants.PeerTypeFromInt32(readHistory.PeerType) {
	case constants.PeerTypeSelf, constants.PeerTypeUser:
		mData, _ := proto.Marshal(readHistory)
		_, _, err := m.producer.SendMessage(constants.InboxMessageReadHistory, mData)
		return err
	case constants.PeerTypeChat:
		mData, _ := proto.Marshal(readHistory)
		_, _, err := m.chatProducer.SendMessage(constants.OutboxChatMessageReadHistory, mData)
		return err
	default:
		log.Warnf("ReadHistoryMessage type:%v", readHistory.PeerType)
		return nil
	}
}

func (m *MessageCoreService) ReadHistoryMessageData(readHistory *msgService.ReadHistory) error {

	peerType := constants.PeerTypeFromInt32(readHistory.PeerType)
	peerId, userId := readHistory.PeerId, readHistory.UserId
	switch peerType {
	case constants.PeerTypeChannel:
		fallthrough
	default:
		panic(fmt.Sprintf("RevokeMessageData %v", constants.PeerTypeFromInt32(readHistory.PeerType)))
	case constants.PeerTypeChat:
		//peerId = message.MakePeerId(peerId, peerType)
	case constants.PeerTypeUser:
		userId, peerId = readHistory.PeerId, readHistory.UserId
	}

	now := int32(time.Now().Unix())
	pts, messageId, err := m.inboxesCore.ReadHistoryMessageData(userId, peerId, peerType, readHistory.GlobalMessageId, now)
	if err != nil {
		log.Errorf("ReadHistoryMessageData inboxesCore error:%s", err.Error())
		return err
	}
	if pts <= 0 {
		log.Warnf("ReadHistoryMessageData inboxesCore pts = %d", pts)
		return nil
	}

	updateOutboxReadHistory := mtproto.NewTLUpdateReadHistoryOutbox(nil)
	updateOutboxReadHistory.SetPeer9961FD5C71(input.MakeInputPeerValue(peerId, peerType).ToPeer())
	updateOutboxReadHistory.SetMaxId(messageId)
	updateOutboxReadHistory.SetPts(pts)
	updateOutboxReadHistory.SetPtsCount(1)

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{updateOutboxReadHistory.To_Update()},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    now,
		Seq:     0,
	})

	//TODO:(Coderxw)
	_, err = syncClient.GetSyncClientById(userId).ReqSyncUpdate(context.TODO(), &syncClient.SyncUpdate{
		UserId:   userId,
		Updates:  updates.To_Updates(),
		PeerType: constants.PeerTypeUser.ToInt32(),
	})
	if err != nil {
		log.Warnf("ReadHistoryMessageData userId:%d error:%s", userId, err.Error())
	}
	return nil
}

func (m *MessageCoreService) ReadHistoryMessageDataList(r []*msgService.ReadHistory) error {
	if len(r) == 0 {
		log.Warnf("ReadHistoryMessageDataList len(value) == 0")
		return nil
	}
	switch constants.PeerTypeFromInt32(r[0].PeerType) {
	//case constants.PeerTypeSelf, constants.PeerTypeUser:
	//	return m.sendInboxUserMessages(r.FromUserId, r.PeerId, r.DataList, r.GlobalMessageIdList)
	case constants.PeerTypeChat:
		var err error
		for idx, v := range r {
			err = m.ReadHistoryMessageData(v)
			if err != nil {
				log.Errorf("ReadHistoryMessageDataList idx:%d v:%+v error:%s", idx, v, err.Error())
				return err
			}
		}
		return nil
	//case constants.PeerTypeChannel:
	//	return m.sendInboxChannelMessages(r.FromUserId, r.PeerId, r.DataList, r.GlobalMessageIdList)
	default:
		log.Errorf("ReadHistoryMessageDataList peerType:%+v", r[0].PeerType)
		return nil
	}
}
