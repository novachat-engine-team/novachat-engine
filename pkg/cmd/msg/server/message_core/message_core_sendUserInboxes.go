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
	"novachat_engine/mtproto"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
)

func (m *MessageCoreService) sendInboxUserMessages(userId int64, peerId int64, list []*msgService.SendMessageData, globalMessageIdList []int64) error {
	for _, v := range list {
		v.Message.Out = false
	}
	ptsList, err := m.inboxesCore.SendMessages(userId, peerId, constants.PeerTypeUser, list, globalMessageIdList)
	if err != nil {
		log.Errorf("sendInboxUserMessages userId:%d peerId:%d error:%s", userId, peerId, err.Error())
		return fmt.Errorf("sendInboxUserMessages userId:%d peerId:%d error", userId, peerId)
	}

	messageIdList := make([]int32, 0, len(list))
	updates := &mtproto.Updates{
		Updates: make([]*mtproto.Update, 0, len(ptsList)),
	}
	for idx := range list {
		if !list[idx].Message.Out {
			list[idx].Message.PeerId = mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(userId).ToInt32()}).To_Peer()
			list[idx].Message.ToId = list[idx].Message.PeerId
		}

		messageIdList = append(messageIdList, list[idx].Message.Id)
		updates.Updates = append(updates.Updates, mtproto.NewTLUpdateNewMessage(&mtproto.Update{
			Message1F2B0AFD71: list[idx].Message,
			MessageEBE4681971: list[idx].Message.Message,
			Pts:               ptsList[idx],
			PtsCount:          1,
			Date:              list[idx].Message.Date,
		}).To_Update())
	}
	_, err = syncClient.GetSyncClientById(peerId).ReqSyncUpdate(context.TODO(), &syncClient.SyncUpdate{
		UserId:   peerId,
		Updates:  mtproto.NewTLUpdates(updates).To_Updates(),
		PeerType: constants.PeerTypeUser.ToInt32(),
		Push:     true,
	})
	if err != nil {
		log.Warnf("sendInboxUserMessages userId:%d peerId:%d error:%s", userId, peerId, err.Error())
	}

	log.Infof("sendInboxUserMessages userId:%d peerId:%d messageIdList:%v", userId, peerId, messageIdList)
	return nil
}
