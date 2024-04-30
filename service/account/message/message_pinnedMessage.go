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
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	msgClient "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	data_message "novachat_engine/service/data/messages/message"
	"novachat_engine/service/input"
)

func (c *Core) PinnedMessage(userId int64, authKeyId int64, inputPeer *input.InputPeer, messageId int32, unpin bool, side bool) (*mtproto.Updates, error) {

	conversation, err := c.messageCore.GetConversation(userId, &data_message.Conversation{PeerId: inputPeer.GetPeerId(), PeerType: inputPeer.GetPeerType().ToInt32()})
	if err != nil {
		log.Errorf("PinnedMessage error:%s", err.Error())
		return nil, err
	}

	var resp *types.Any
	pinnedMessage := &msgClient.PinnedMessage{
		UserId:    userId,
		PeerId:    inputPeer.GetPeerId(),
		PeerType:  inputPeer.GetPeerType().ToInt32(),
		MessageId: messageId,
		Unpin:     unpin,
		OneSide:   side,
		AuthKeyId: authKeyId,
	}

	if idx := util.Index(conversation.PinIds, func(idx int) bool {
		return conversation.PinIds[idx].OwnerUserId == userId
	}); idx >= 0 {
		if unpin {
			copy(conversation.PinIds, conversation.PinIds[:idx])
			copy(conversation.PinIds[idx:], conversation.PinIds[idx+1:])
			conversation.PinIds = conversation.PinIds[:len(conversation.PinIds)-1]
			resp, err = msgClient.GetMsgClient().ReqPinnedMessage(context.Background(), pinnedMessage)
		} else {
			if !conversation.PinIds[idx].Full && side == false {
				conversation.PinIds[idx].Full = side == false
				resp, err = msgClient.GetMsgClient().ReqPinnedMessage(context.Background(), pinnedMessage)
			}
		}
	} else {
		if unpin {
			return mtproto.NewTLUpdates(nil).To_Updates(), nil
		}
		resp, err = msgClient.GetMsgClient().ReqPinnedMessage(context.Background(), pinnedMessage)
	}
	if err != nil {
		log.Errorf("PinnedMessage PinnedMessage error:%s", err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	types.UnmarshalAny(resp, &updates)
	return &updates, nil
}
