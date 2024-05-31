/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	jsoniter "github.com/json-iterator/go"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
)

func (impl *Impl) ReqSendInBoxesMessages(ctx context.Context, request *chatService.SendInBoxesMessages) (*types.Any, error) {
	peerType := constants.PeerTypeFromInt32(request.Message.PeerType)
	switch peerType {
	case constants.PeerTypeChat, constants.PeerTypeChannel:
		break
	default:
		log.Errorf("ReqSendInBoxesMessages peerType:%d error", peerType)
		return types.MarshalAny(mtproto.ToMTBool(true))
	}

	inboxPiece := impl.getPiece()

	chat, err := impl.chatManager.GetChat(request.Message.PeerId)
	if err != nil {
		log.Errorf("ReqSendInBoxesMessages GetChat chatId:%d error:%s", request.Message.PeerId, err.Error())
		return nil, err
	}
	if chat.Invalid() || chat.Deleted() {
		log.Warnf("ReqSendInBoxesMessages chatId:%d invalid:%v deleted:%v", request.Message.PeerId, chat.Invalid(), chat.Deleted())
		return types.MarshalAny(mtproto.ToMTBool(true))
	}

	var key string
	var ok bool
	var value []*msgService.SendMessages
	pieceKeyMap := make(map[string][]*msgService.SendMessages, chat.GetChatInfo().Count)
	chat.GetChatInfo().Iteration(func(participant *data_chat.ChatParticipant) {
		if participant.UserId != 0 &&
			participant.UserId != request.Message.FromUserId &&
			participant.State == data_chat.ParticipantState_ParticipantStateNormal {
			key = msgService.GetMsgByUserIdKey(participant.UserId)
			if value, ok = pieceKeyMap[key]; !ok {
				value = make([]*msgService.SendMessages, 0, 4)
			}
			value = append(value, &msgService.SendMessages{
				FromUserId:          participant.UserId,
				PeerId:              request.Message.PeerId,
				PeerType:            request.Message.PeerType,
				DataList:            request.Message.DataList,
				GlobalMessageIdList: request.Message.GlobalMessageIdList,
			})
			pieceKeyMap[key] = value
		}
	})

	var idx int32
	sendMessageList := make([]*msgService.SendMessages, inboxPiece)
	for key, value = range pieceKeyMap {
		idx = 0
		for _, v := range value {
			sendMessageList[idx] = v
			idx++
			if idx >= inboxPiece {
				data, _ := jsoniter.Marshal(sendMessageList)
				_, _, err = impl.chatInProducer.SendMessage(constants.InboxChatMessage, data)
				if err != nil {
					log.Warnf("ReqSendInBoxesMessages chatId:%d SendMessage error:%s", request.Message.PeerId, err.Error())
					return nil, err
				}
				idx = 0
			}
		}

		if idx > 0 {
			data, _ := jsoniter.Marshal(sendMessageList[:idx])
			_, _, err = impl.chatInProducer.SendMessage(constants.InboxChatMessage, data)
			if err != nil {
				log.Warnf("ReqSendInBoxesMessages chatId:%d SendMessage error:%s", request.Message.PeerId, err.Error())
				return nil, err
			}
		}
	}
	return types.MarshalAny(mtproto.ToMTBool(true))
}
