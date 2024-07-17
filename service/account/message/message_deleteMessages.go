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
	"github.com/go-redsync/redsync"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cache"
	msgClient "novachat_engine/pkg/cmd/msg/rpc_client"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/common/hash"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message"
	data_message "novachat_engine/service/data/messages/message"
	"novachat_engine/service/input"
	"time"
)

func (c *Core) DeleteMessages(
	userId int64,
	authKeyId int64,
	revoke bool,
	idList []int32) (*mtproto.Messages_AffectedMessages, error) {

	mutex := cache.GetRedSyncClient().NewMutex(
		makeMutexDeleteMessageName(userId, hash.HashIdList(idList, true, true)),
		redsync.SetExpiry(3*time.Second),
		redsync.SetTries(3))
	err := mutex.Lock()
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
		log.Errorf("DeleteMessages NewMutex error:%s need retry", err.Error())
		return nil, err
	}
	defer mutex.Unlock()

	var globalMessageIdList []int64
	var messageList []*data_message.Message
	messageList, err = c.messageCore.GetMessageList(userId, idList)
	if err != nil {
		log.Errorf("DeleteMessages GetMessageList error:%s", err.Error())
		return nil, err
	}

	var peerId int64
	var peerType constants.PeerType
	tempIdList := make([]int32, 0, len(messageList))
	if len(messageList) > 0 {
		globalMessageIdList = make([]int64, 0, len(messageList))
		for _, v := range messageList {
			if !v.Deleted {
				tempIdList = append(tempIdList, v.Id)
			}

			peerId, peerType = message.MakePeerType(v.PeerId)
			globalMessageIdList = append(globalMessageIdList, v.GlobalMessageId)
		}
	} else {
		log.Warnf("DeleteMessages GetMessageList not found userId:%d messageIdList:%v", userId, idList)
		return mtproto.NewTLMessagesAffectedMessages(nil).To_Messages_AffectedMessages(), nil
	}

	affectedMessage := &mtproto.Messages_AffectedMessages{}
	var pts int32
	if len(tempIdList) > 0 {
		pts, err = c.messageCore.DeleteMessages(userId, tempIdList)
		if err != nil {
			err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
			log.Errorf("DeleteMessages DeleteMessages error:%s need retry", err.Error())
			return nil, err
		}
		affectedMessage.Pts = pts
		affectedMessage.PtsCount = 1
	}

	if revoke && !(peerType == constants.PeerTypeUser && peerId == userId) {
		_, err = msgClient.GetMsgClient().ReqRevokeMessages(context.Background(), &msgClient.RevokeMessages{
			UserId:              userId,
			PeerId:              peerId,
			PeerType:            peerType.ToInt32(),
			GlobalMessageIdList: globalMessageIdList,
		})
		if err != nil {
			log.Errorf("DeleteMessages GetMessageList ReqRevokeMessages userId:%d messageIdList:%v error:%s", userId, idList, err.Error())
			return nil, err
		}
	}

	if len(idList) > 0 && pts > 0 {
		//TODO:(Coderxw)
		_, err = syncClient.GetSyncClientById(userId).ReqSyncUpdate(context.TODO(), &syncClient.SyncUpdate{
			UserId:          userId,
			IgnoreAuthKeyId: authKeyId,
			Updates: mtproto.NewTLUpdates(&mtproto.Updates{
				Updates: []*mtproto.Update{mtproto.NewTLUpdateDeleteMessages(&mtproto.Update{
					PeerId:   mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(userId).ToInt32()}).To_Peer(),
					Messages: idList,
					Pts:      pts,
					PtsCount: 1,
				}).To_Update()},
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
				Date:  int32(time.Now().Unix()),
				Seq:   0,
			}).To_Updates(),
			PeerType: constants.PeerTypeUser.ToInt32(),
		})
		if err != nil {
			log.Warnf("DeleteMessages userId:%d peerId:%d idList:%v error:%s", userId, peerId, idList, err.Error())
		}
	}

	return mtproto.NewTLMessagesAffectedMessages(affectedMessage).To_Messages_AffectedMessages(), nil
}

func (c *Core) DeleteHistory(
	userId int64,
	authKeyId int64,
	justClear bool,
	revoke bool,
	inputPeer *input.InputPeer,
	maxId int32,
	now int32) (*mtproto.Messages_AffectedHistory, error) {

	var globalMessageIdList []int64
	var messageList []*data_message.Message
	messageList, err := c.messageCore.GetHistoryMessages(
		userId,
		inputPeer.GetPeerId(),
		inputPeer.GetPeerType(),
		0,
		-1,
		maxId,
		0,
		true)
	if err != nil {
		log.Errorf("DeleteHistory GetHistoryMessages error:%s", err.Error())
		return nil, err
	}

	var messageIdList []int32
	peerId, peerType := inputPeer.GetPeerId(), inputPeer.GetPeerType()
	tempIdList := make([]int32, 0, len(messageList))
	if len(messageList) > 0 {
		for _, v := range messageList {
			if !v.Deleted {
				tempIdList = append(tempIdList, v.Id)
			}
			if peerType == constants.PeerTypeChannel ||
				peerType == constants.PeerTypeChat {
				globalMessageIdList = append(globalMessageIdList, int64(v.Id))
			} else {
				globalMessageIdList = append(globalMessageIdList, v.GlobalMessageId)
			}
		}

	} else {
		log.Warnf("DeleteHistory GetMessageList not found userId:%d messageIdList:%v", userId, maxId)
	}

	affectedHistory := &mtproto.Messages_AffectedHistory{}
	var pts int32
	if len(tempIdList) > 0 {
		pts, messageIdList, err = c.messageCore.DeleteMessagesByMaxId(userId, peerId, peerType, maxId, now, 0)
		if err != nil {
			err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
			log.Errorf("DeleteHistory DeleteMessagesByMaxId error:%s need retry", err.Error())
			return nil, err
		}

		affectedHistory.Pts = pts
		affectedHistory.PtsCount = 1
	}

	if justClear && revoke {
		_, err = msgClient.GetMsgClient().ReqRevokeMessages(context.Background(), &msgClient.RevokeMessages{
			UserId:              userId,
			PeerId:              peerId,
			PeerType:            peerType.ToInt32(),
			GlobalMessageIdList: globalMessageIdList,
			Range:               false,
		})
		if err != nil {
			log.Errorf("DeleteMessages GetMessageList ReqRevokeMessages userId:%d messageIdList:%v error:%s", userId, messageIdList, err.Error())
			return nil, err
		}
	}

	if len(messageIdList) > 0 {
		//TODO:(Coderxw)
		_, err = syncClient.GetSyncClientById(userId).ReqSyncUpdate(context.TODO(), &syncClient.SyncUpdate{
			UserId:          userId,
			IgnoreAuthKeyId: authKeyId,
			Updates: mtproto.NewTLUpdates(&mtproto.Updates{
				Updates: []*mtproto.Update{mtproto.NewTLUpdateDeleteMessages(&mtproto.Update{
					PeerId:   mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(userId).ToInt32()}).To_Peer(),
					Messages: messageIdList,
					Pts:      pts,
					PtsCount: 1,
				}).To_Update()},
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
				Date:  int32(time.Now().Unix()),
				Seq:   0,
			}).To_Updates(),
			PeerType: constants.PeerTypeUser.ToInt32(),
		})
		if err != nil {
			log.Warnf("DeleteMessages userId:%d peerId:%d idList:%v error:%s", userId, peerId, messageIdList, err.Error())
		}
	}

	return mtproto.NewTLMessagesAffectedHistory(affectedHistory).To_Messages_AffectedHistory(), nil
}
