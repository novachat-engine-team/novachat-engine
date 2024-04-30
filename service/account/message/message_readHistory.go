/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package message

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	chatClient "novachat_engine/pkg/cmd/chat/rpc_client"
	msgClient "novachat_engine/pkg/cmd/msg/rpc_client"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message"
	"novachat_engine/service/input"
	"time"
)

func (c *Core) ReadHistory(
	userId int64,
	authKeyId int64,
	maxId int32,
	inputPeer *input.InputPeer,
	date int32) (*mtproto.Messages_AffectedHistory, error) {

	if userId == inputPeer.GetPeerId() && inputPeer.GetPeerType() == constants.PeerTypeUser {
		return mtproto.NewTLMessagesAffectedHistory(nil).To_Messages_AffectedHistory(), nil
	}

	log.Debugf("ReadHistory userId:%d peerId:%d peerType:%v maxId:%d", userId, inputPeer.GetPeerId(), inputPeer.GetPeerType(), maxId)
	pts, foldId, err := c.messageCore.ReadHistory(userId, inputPeer.GetPeerId(), inputPeer.GetPeerType(), maxId, date, message.InboxType)
	if err != nil {
		log.Errorf("ReadHistory ReadHistory userId:%d peerId:%d peerType:%v maxId:%d error:%s", userId, inputPeer.GetPeerId(), inputPeer.GetPeerType(), maxId, err.Error())
		return nil, err
	}

	messageDataList, err := c.messageCore.GetMessageList(userId, []int32{maxId})
	if err != nil {
		log.Errorf("ReadHistory GetMessageList userId:%d peerId:%d peerType:%v maxId:%d error:%s", userId, inputPeer.GetPeerId(), inputPeer.GetPeerType(), maxId, err.Error())
		return nil, err
	}
	if len(messageDataList) == 0 {
		log.Warnf("ReadHistory GetMessageList userId:%d peerId:%d peerType:%v maxId:%d empty message", userId, inputPeer.GetPeerId(), inputPeer.GetPeerType(), maxId)
		return mtproto.NewTLMessagesAffectedHistory(nil).To_Messages_AffectedHistory(), nil
	}
	if pts <= 0 {
		log.Warnf("ReadHistory GetMessageList userId:%d peerId:%d peerType:%v maxId:%d empty message", userId, inputPeer.GetPeerId(), inputPeer.GetPeerType(), maxId)
		return mtproto.NewTLMessagesAffectedHistory(nil).To_Messages_AffectedHistory(), nil
	}

	affectedHistory := mtproto.NewTLMessagesAffectedHistory(nil).To_Messages_AffectedHistory()
	affectedHistory.Pts = pts
	affectedHistory.PtsCount = 1

	switch inputPeer.GetPeerType() {
	case constants.PeerTypeUser:
		_, err = msgClient.GetMsgClient().ReqReadHistory(context.Background(), &msgClient.ReadHistory{
			UserId:          userId,
			PeerId:          inputPeer.GetPeerId(),
			PeerType:        inputPeer.GetPeerType().ToInt32(),
			GlobalMessageId: messageDataList[len(messageDataList)-1].GlobalMessageId,
		})
	case constants.PeerTypeChat, constants.PeerTypeChannel:
		_, err = chatClient.GetChatClientByKeyId(inputPeer.GetPeerId()).ReqReadHistory(context.Background(), &chatClient.ChatReadHistory{
			Value: &msgClient.ReadHistory{
				UserId:          userId,
				PeerId:          inputPeer.GetPeerId(),
				PeerType:        inputPeer.GetPeerType().ToInt32(),
				GlobalMessageId: messageDataList[len(messageDataList)-1].GlobalMessageId,
			}})
	default:
		panic(fmt.Sprintf("ReadHistory known PeerType:%v", inputPeer.GetPeerType()))
	}
	if err != nil {
		log.Errorf("ReadHistory ReqUserReadHistory userId:%d peerId:%d peerType:%v maxId:%d error:%s", userId, inputPeer.GetPeerId(), inputPeer.GetPeerType(), maxId, err.Error())
		return nil, err
	}

	if pts > 0 {
		updateInboxReadHistory := mtproto.NewTLUpdateReadHistoryInbox(nil)
		updateInboxReadHistory.SetPeer9961FD5C71(inputPeer.ToPeer())
		updateInboxReadHistory.SetMaxId(maxId)
		updateInboxReadHistory.SetPts(pts)
		updateInboxReadHistory.SetPtsCount(1)
		updateInboxReadHistory.SetFolderId(foldId)

		updates := mtproto.NewTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{updateInboxReadHistory.To_Update()},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		})

		//TODO:(Coderxw)
		_, err = syncClient.GetSyncClientById(userId).ReqSyncUpdate(context.TODO(), &syncClient.SyncUpdate{
			UserId:          userId,
			IgnoreAuthKeyId: authKeyId,
			Updates:         updates.To_Updates(),
			PeerType:        constants.PeerTypeUser.ToInt32(),
		})
		if err != nil {
			log.Warnf("ReadHistory userId:%d error:%s", userId, err.Error())
		}
	}

	return affectedHistory, nil
}
