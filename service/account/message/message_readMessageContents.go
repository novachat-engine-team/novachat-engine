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
	"novachat_engine/mtproto"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/message"
	"time"
)

func (c *Core) ReadMessageContents(userId int64, idList []int32) (*mtproto.Messages_AffectedMessages, error) {
	messageList, err := c.messageCore.GetMessageList(userId, idList)
	if err != nil {
		log.Errorf("ReadMessageContents userId:%d error:%s", userId, err.Error())
		return nil, err
	}

	if len(messageList) == 0 {
		return mtproto.NewTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
			Pts:      -1,
			PtsCount: 0,
		}).To_Messages_AffectedMessages(), nil
	}

	globalMessageIdList := make([]int64, 0, len(messageList))
	for _, v := range messageList {
		globalMessageIdList = append(globalMessageIdList, v.GlobalMessageId)
	}

	messageList, err = c.messageCore.GetMessageByGlobalMessageIdList(userId, globalMessageIdList, true)
	if err != nil {
		log.Errorf("ReadMessageContents GetMessageByGlobalMessageIdList:%v error:%s", globalMessageIdList, err.Error())
		return nil, err
	}

	userMap := map[int64][]int32{}
	appendUserMap := func(userId int64, id int32) {
		v, ok := userMap[userId]
		if !ok {
			v = []int32{}
		}
		if util.IndexInt32s(&v, id) < 0 {
			v = append(v, id)
		}
	}
	for _, v := range messageList {
		peerId, peerType := message.MakePeerType(v.PeerId)
		switch peerType {
		case constants.PeerTypeUser, constants.PeerTypeSelf:
			if v.Out == true {
				if v.FromUserId == userId {
					continue
				}
				appendUserMap(v.FromUserId, v.Id)
			} else {
				if peerId == userId {
					continue
				}
				appendUserMap(peerId, v.Id)
			}
		case constants.PeerTypeChat:
			if v.FromUserId == userId {
				continue
			}
			appendUserMap(v.FromUserId, v.Id)
		default:
			continue
		}
	}

	for key, values := range userMap {
		if len(values) == 0 {
			continue
		}

		_, err = syncClient.GetSyncClientById(key).ReqSyncUpdate(context.Background(), &syncClient.SyncUpdate{
			UserId: key,
			Updates: mtproto.NewTLUpdates(&mtproto.Updates{
				Updates: []*mtproto.Update{mtproto.NewTLUpdateReadMessagesContents(&mtproto.Update{
					Messages: values,
					Pts:      -1,
					PtsCount: 0,
				}).To_Update(),
				},
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
				Date:  int32(time.Now().Unix()),
				Seq:   0,
			}).To_Updates(),
			Push:     false,
			PeerType: constants.PeerTypeUser.ToInt32(),
		})
		if err != nil {
			log.Warnf("ReadMessageContents userId:%d, values:%v error:%s", key, values, err.Error())
		}
	}

	return mtproto.NewTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      -1,
		PtsCount: 0,
	}).To_Messages_AffectedMessages(), nil
}
