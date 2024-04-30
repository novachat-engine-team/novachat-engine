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
	"errors"
	"fmt"
	"github.com/gogo/protobuf/types"
	jsoniter "github.com/json-iterator/go"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/cmd/chat/service"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
)

func (impl *Impl) ReqReadInMessages(ctx context.Context, request *chatService.ReadInMessages) (*types.Any, error) {
	peerType := constants.PeerTypeFromInt32(request.Value.PeerType)
	switch peerType {
	case constants.PeerTypeChat, constants.PeerTypeChannel:
		break
	default:
		log.Errorf("ReqReadInMessages peerType:%d error", peerType)
		return nil, errorsService.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST, fmt.Errorf("ReqReadInMessages peerType:%d error", peerType).Error())
	}

	chat, err := impl.chatManager.GetChat(request.Value.PeerId)
	if err != nil {
		if errors.Is(err, service.DefaultErrChatIdInvalid) {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
		} else {
			return nil, err
		}
	}
	if !chat.Invalid() || chat.Deleted() {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
	}

	participantOfMe := chat.GetChatInfo().GetParticipants(request.Value.UserId)
	if participantOfMe == nil {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
	}

	var idx int32
	var key string
	var ok bool
	var value []*msgService.ReadHistory
	inboxPiece := impl.getPiece()
	chatReadHistoryMessageList := make([]*msgService.ReadHistory, inboxPiece)
	pieceKeyMap := make(map[string][]*msgService.ReadHistory, chat.GetChatInfo().Count)
	chat.GetChatInfo().Iteration(func(participant *data_chat.ChatParticipant) {
		if participant.UserId != 0 && participant.UserId != request.Value.UserId && participant.State == data_chat.ParticipantState_ParticipantStateNormal {
			key = msgService.GetMsgByUserIdKey(participant.UserId)
			if value, ok = pieceKeyMap[key]; !ok {
				value = make([]*msgService.ReadHistory, 0, 4)
			}
			value = append(value, &msgService.ReadHistory{
				UserId:          participant.UserId,
				PeerId:          request.Value.PeerId,
				PeerType:        request.Value.PeerType,
				GlobalMessageId: request.Value.GlobalMessageId,
			})
			pieceKeyMap[key] = value
		}
	})
	for key, value = range pieceKeyMap {
		idx = 0
		for _, v := range value {
			chatReadHistoryMessageList[idx] = v
			idx++
			if idx >= inboxPiece {
				data, _ := jsoniter.Marshal(chatReadHistoryMessageList)
				_, _, err = impl.chatInProducer.SendMessage(constants.InboxChatMessageReadHistory, data)
				if err != nil {
					log.Warnf("ReqReadInMessages chatId:%d SendMessage error:%s", request.Value.PeerId, err.Error())
					return nil, err
				}
				idx = 0
			}
		}

		if idx > 0 {
			data, _ := jsoniter.Marshal(chatReadHistoryMessageList[:idx])
			_, _, err = impl.chatInProducer.SendMessage(constants.InboxChatMessageReadHistory, data)
			if err != nil {
				log.Warnf("ReqReadInMessages chatId:%d SendMessage error:%s", request.Value.PeerId, err.Error())
				return nil, err
			}
		}
	}
	return types.MarshalAny(mtproto.ToMTBool(true))
}
