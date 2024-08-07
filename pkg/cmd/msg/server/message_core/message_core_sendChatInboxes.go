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
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/chat"
	"novachat_engine/service/constants"
)

func (m *MessageCoreService) sendInboxChatMessages(userId int64, peerId int64, list []*msgService.SendMessageData, ptsIdList []int64) error {
	var err error
	for _, v := range list {
		v.Message.Out = false
	}
	chatList, err := chatService.GetChatClientByKeyId(peerId).ReqAllChat(context.Background(), &chatService.AllChat{
		Ids:    []int64{peerId},
		UserId: userId,
	})
	if err != nil {
		log.Errorf("sendInboxChatMessages ReqGetParticipants userId:%d peerId:%d error:%s", userId, peerId, err.Error())
		return err
	}
	if len(chatList.Values) == 0 {
		return nil
	}

	var pts int64
	if ptsIdList != nil {
		pts = ptsIdList[0]
	}
	for _, ptsV := range ptsIdList {
		if pts < ptsV {
			pts = ptsV
		}
	}

	err = m.inboxesCore.SendChannelMessages(userId, peerId, constants.PeerTypeChannel, list, int32(pts))
	if err != nil {
		log.Errorf("sendInboxChatMessages userId:%d peerId:%d error:%s", userId, peerId, err.Error())
		return fmt.Errorf("sendInboxChatMessages userId:%d peerId:%d error", userId, peerId)
	}

	messageIdList := make([]int32, 0, len(list))
	updates := &mtproto.Updates{
		Updates: make([]*mtproto.Update, 0, len(list)),
		//TODO:(Coderxw channel)
		Chats: []*mtproto.Chat{chat.ToChat(userId, chatList.Values[0], mtproto.CurrentLayer)},
	}

	for idx := range list {
		pts = 0
		if len(ptsIdList) > idx {
			pts = ptsIdList[idx]
		}
		messageIdList = append(messageIdList, list[idx].Message.Id)

		//  updateNewMessage#1f2b0afd message:Message pts:int pts_count:int = Update;
		updates.Updates = append(updates.Updates, mtproto.NewTLUpdateNewChannelMessage(&mtproto.Update{
			Message1F2B0AFD71: list[idx].Message,
			Pts:               int32(pts),
			PtsCount:          1,
		}).To_Update())
	}
	_, err = syncClient.GetSyncClientById(userId).ReqSyncUpdate(context.TODO(), &syncClient.SyncUpdate{
		UserId:   userId,
		Updates:  mtproto.NewTLUpdates(updates).To_Updates(),
		PeerType: constants.PeerTypeUser.ToInt32(),
	})
	if err != nil {
		log.Warnf("sendInboxChatMessages userId:%d peerId:%d error:%s", userId, peerId, err.Error())
	}

	log.Infof("sendInboxChatMessages userId:%d peerId:%d messageIdList:%v", userId, peerId, messageIdList)
	return nil
}
