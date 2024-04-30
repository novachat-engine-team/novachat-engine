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
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	msgClient "novachat_engine/pkg/cmd/msg/rpc_client"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"time"
)

func (m *MessageCoreService) pinnedMessage(userId int64, peerId int64, peerType constants.PeerType, globalMessageId int64, unpin bool) error {
	var err error
	mData, _ := proto.Marshal(&msgClient.PinnedMessage{
		UserId:          userId,
		PeerId:          peerId,
		PeerType:        peerType.ToInt32(),
		Unpin:           unpin,
		GlobalMessageId: globalMessageId,
	})
	switch peerType {
	case constants.PeerTypeSelf, constants.PeerTypeUser:
		_, _, err = m.producer.SendMessage(constants.OutboxMessagePinnedMessages, mData)
	case constants.PeerTypeChat:
		_, _, err = m.chatProducer.SendMessage(constants.OutboxChatMessagePinnedMessages, mData)
	default:
		break
	}
	if err != nil {
		log.Errorf("pinnedMessage userId:%d peerId:%d peerType:%d globalMessageId:%d unpin:%v", userId, peerId, peerType, globalMessageId, unpin)
	}
	return err
}

func (m *MessageCoreService) PinnedMessage(authKeyId int64, userId int64, peerId int64, peerType constants.PeerType, msgId int32, unpin bool, side bool, globalMessageId int64) (*types.Any, error) {
	log.Debugf("PinnedMessage authKeyId:%d userId:%d peerId:%d peerType:%d msg:%d unpin:%v side:%v", authKeyId, userId, peerId, peerType, msgId, unpin, side)

	var err error
	var pts int32
	if authKeyId == 0 {
		pts = int32(globalMessageId)
	}

	if peerType == constants.PeerTypeSelf || peerType == constants.PeerTypeUser {
		pts, globalMessageId, err = m.outboxesCore.PinnedMessage(userId, peerId, peerType, msgId, unpin, side, globalMessageId)
		if err != nil {
			log.Errorf("PinnedMessage userId:%d msgId:%d error:%s", userId, msgId, err.Error())
			return nil, err
		}
	} else if peerType == constants.PeerTypeChat || peerType == constants.PeerTypeChannel {
		err = m.outboxesCore.PinnedChannelMessage(userId, peerId, peerType, msgId, unpin, side)
		if err != nil {
			log.Errorf("PinnedMessage userId:%d msgId:%d error:%s", userId, msgId, err.Error())
			return nil, err
		}
	}

	if !side && !(peerType == constants.PeerTypeUser && peerId == userId) && globalMessageId != 0 {
		err = m.pinnedMessage(userId, peerId, peerType, globalMessageId, unpin)
		if err != nil {
			log.Fatalf("PinnedMessage userId:%d msgId:%d globalMessageId:%d error:%s", userId, msgId, globalMessageId, err.Error())
		}
	}

	if pts == 0 {
		log.Warnf("PinnedMessage userId:%d msgId:%d pts = 0", userId, msgId)
		return types.MarshalAny(mtproto.NewTLUpdates(nil).To_Updates())
	}

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{mtproto.NewTLUpdatePinnedMessages(&mtproto.Update{
			Pinned:         !unpin,
			Peer9961FD5C71: input.MakeInputPeerValue(peerId, peerType).ToPeer(),
			Messages:       []int32{msgId},
			Pts:            pts,
			PtsCount:       1,
		}).To_Update()},
		Users: []*mtproto.User{},
		Chats: []*mtproto.Chat{},
		Date:  int32(time.Now().Unix()),
		Seq:   0,
	}).To_Updates()

	_, err = syncService.GetSyncClientById(userId).ReqSyncUpdate(context.Background(), &syncService.SyncUpdate{
		UserId:          userId,
		IgnoreAuthKeyId: authKeyId,
		Updates:         updates,
		PeerType:        constants.PeerTypeUser.ToInt32(),
	})
	if err != nil {
		log.Warnf("PinnedMessage ReqSyncUpdate error:%s", err.Error())
	}
	return types.MarshalAny(updates)
}

func (m *MessageCoreService) PinnedMessageDataList(r []*msgClient.PinnedMessage) error {
	var err error
	for _, v := range r {
		err = m.PinnedMessageData(v)
		if err != nil {
			log.Errorf("PinnedMessageDataList pinnedMessage:%+v error:%s", v, err.Error())
			break
		}
	}
	return err
}

func (m *MessageCoreService) PinnedMessageData(pinnedMessage *msgClient.PinnedMessage) error {
	log.Debugf("PinnedMessageData pinnedMessage:%+v", pinnedMessage)

	if constants.PeerTypeFromInt32(pinnedMessage.PeerType) == constants.PeerTypeUser {
		pinnedMessage.UserId, pinnedMessage.PeerId = pinnedMessage.PeerId, pinnedMessage.UserId
	}
	_, err := m.PinnedMessage(0,
		pinnedMessage.UserId,
		pinnedMessage.PeerId,
		constants.PeerTypeFromInt32(pinnedMessage.PeerType),
		0,
		pinnedMessage.Unpin,
		true, pinnedMessage.GlobalMessageId)
	if err != nil {
		log.Errorf("PinnedMessageData pinnedMessage:%+v error:%s", pinnedMessage, err.Error())
		return err
	}
	return err
}
