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
	"errors"
	"github.com/gogo/protobuf/proto"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"time"
)

func (m *MessageCoreService) RevokeMessage(revokeMessage *msgService.RevokeMessages) error {
	switch constants.PeerTypeFromInt32(revokeMessage.PeerType) {
	case constants.PeerTypeSelf, constants.PeerTypeUser:
		mData, _ := proto.Marshal(revokeMessage)
		_, _, err := m.producer.SendMessage(constants.InboxMessageRevokeMessages, mData)
		return err
	default:
		log.Warnf("RevokeMessage type:%v", revokeMessage.PeerType)
		return nil
	}
}

func (m *MessageCoreService) RevokeMessageData(revokeMessage *msgService.RevokeMessages) error {

	var update *mtproto.Update
	peerType := constants.PeerTypeFromInt32(revokeMessage.PeerType)
	peerId, userId := revokeMessage.PeerId, revokeMessage.UserId
	switch peerType {
	case constants.PeerTypeChannel, constants.PeerTypeChat:
		bSync, err := m.inboxesCore.RevokeChannelMessageData(userId, peerId, peerType, revokeMessage.ChannelPts)
		if err != nil {
			log.Errorf("RevokeMessageData inboxes RevokeMessageData error:%s", err.Error())
			return err
		}
		if bSync {
			return nil
		}

		messageIdList := make([]int32, 0, len(revokeMessage.GlobalMessageIdList))
		for _, v := range revokeMessage.GlobalMessageIdList {
			messageIdList = append(messageIdList, int32(v))
		}
		update = mtproto.NewTLUpdateDeleteChannelMessages(&mtproto.Update{
			ChannelId: constants.PeerTypeFromChannelIDType(peerId).ToInt32(),
			PeerId:    mtproto.NewTLPeerChannel(&mtproto.Peer{ChannelId: constants.PeerTypeFromChannelIDType(peerId).ToInt32()}).To_Peer(),
			Messages:  messageIdList,
			Pts:       revokeMessage.ChannelPts,
			PtsCount:  1,
		}).To_Update()
		break

	case constants.PeerTypeUser:
		userId, peerId = revokeMessage.PeerId, revokeMessage.UserId
		pts, messageIdList, err := m.inboxesCore.RevokeMessageData(userId, peerId, peerType, revokeMessage.GlobalMessageIdList, revokeMessage.Range)
		if err != nil {
			log.Errorf("RevokeMessageData inboxes RevokeMessageData error:%s", err.Error())
			return err
		}
		if len(messageIdList) == 0 {
			return nil
		}
		update = mtproto.NewTLUpdateDeleteMessages(&mtproto.Update{
			PeerId:   mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(userId).ToInt32()}).To_Peer(),
			Messages: messageIdList,
			Pts:      pts,
			PtsCount: 1,
		}).To_Update()
	}

	//TODO:(Coderxw)
	_, err := syncClient.GetSyncClientById(userId).ReqSyncUpdate(context.TODO(), &syncClient.SyncUpdate{
		UserId: userId,
		Updates: mtproto.NewTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{update},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}).To_Updates(),
		PeerType: constants.PeerTypeUser.ToInt32(),
	})
	if err != nil {
		log.Warnf("RevokeMessageData userId:%d peerId:%d idList:%v error:%s", userId, peerId, err.Error())
	}
	return nil
}

func (m *MessageCoreService) RevokeMessageDataList(r []*msgService.RevokeMessages) error {
	var err error
	for _, v := range r {
		err = m.RevokeMessageData(v)
		if err != nil {
			log.Errorf("RevokeMessageDataList RevokeMessages:%+v error:%s", v, err.Error())
			break
		}
	}
	return err
}

func (m *MessageCoreService) DeleteChannelMessagesData(r *chatService.DeleteMessagesUpdates) error {
	log.Debugf("DeleteChannelMessagesData r:%+v", r)
	_, err := chatService.GetChatClientByKeyId(r.PeerId).ReqDeleteMessagesUpdates(context.Background(), r)
	if errors.Is(err, mtproto.DefaultRpcError) {
		log.Warnf("DeleteChannelMessagesData userId:%d chatId:%d error:%s", r.UserId, r.PeerId)
		err = nil
	}
	if err != nil {
		log.Errorf("DeleteChannelMessagesData userId:%d chatId:%d error:%s", r.UserId, r.PeerId)
	}
	return err
}
