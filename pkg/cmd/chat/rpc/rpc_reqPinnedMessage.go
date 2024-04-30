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
	"novachat_engine/service/banned_right"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
)

func (impl *Impl) ReqPinnedMessage(ctx context.Context, request *chatService.ChatPinnedMessage) (*types.Any, error) {
	peerType := constants.PeerTypeFromInt32(request.PinnedMessage.PeerType)
	switch peerType {
	case constants.PeerTypeChat, constants.PeerTypeChannel:
		break
	default:
		log.Errorf("ReqPinnedMessage peerType:%d error", peerType)
		return nil, errorsService.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST, fmt.Errorf("ReqPinnedMessage peerType:%d error", peerType).Error())
	}

	chat, err := impl.chatManager.GetChat(request.PinnedMessage.PeerId)
	if err != nil {
		if errors.Is(err, service.DefaultErrChatIdInvalid) {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
		} else {
			return nil, err
		}
	}
	if !chat.Invalid() || chat.Deleted() {
		log.Warnf("ReqPinnedMessage chat invalid chatId:%d", request.PinnedMessage.PeerId)
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
	}

	chatInfo := chat.GetChatInfo()

	participantOfMe := chatInfo.GetParticipants(request.PinnedMessage.UserId)
	if participantOfMe == nil {
		log.Warnf("ReqPinnedMessage chat not found Participants chatId:%d userId:%d", request.PinnedMessage.PeerId, request.PinnedMessage.UserId)
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
	}
	if participantOfMe.State != data_chat.ParticipantState_ParticipantStateNormal {
		log.Warnf("ReqPinnedMessage chat state:%d Participants chatId:%d userId:%d", participantOfMe.State, request.PinnedMessage.PeerId, request.PinnedMessage.UserId)
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
	}

	if !(chatInfo.ChatData.Creator == request.PinnedMessage.UserId ||
		participantOfMe.Admin) {
		chatBannedRights := banned_right.RightsToChatBannedRight(chat.GetChatInfo().ChatData.BannedRights, chat.GetChatInfo().ChatData.RightsUtilDate)
		if chatBannedRights.PinMessages {
			log.Warnf("ReqPinnedMessage chat cannot PinMessages Participants chatId:%d userId:%d", request.PinnedMessage.PeerId, request.PinnedMessage.UserId)
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PIN_RESTRICTED)
		}
	}

	resp, err := msgService.GetMsgClient().ReqPinnedMessage(context.Background(), request.PinnedMessage)
	if err != nil {
		log.Errorf("ReqPinnedMessage chatId:%d userId:%d error:%s", request.PinnedMessage.PeerId, request.PinnedMessage.UserId, err.Error())
		return nil, err
	}

	return resp, nil
}

func (impl *Impl) ReqInPinnedMessage(ctx context.Context, request *chatService.ChatInPinnedMessage) (*types.Any, error) {
	peerType := constants.PeerTypeFromInt32(request.PinnedMessage.PeerType)
	switch peerType {
	case constants.PeerTypeChat:
		break
	default:
		log.Errorf("ReqPinnedMessage peerType:%d error", peerType)
		return nil, errorsService.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST, fmt.Errorf("ReqPinnedMessage peerType:%d error", peerType).Error())
	}

	chat, err := impl.chatManager.GetChat(request.PinnedMessage.PeerId)
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

	participantOfMe := chat.GetChatInfo().GetParticipants(request.PinnedMessage.UserId)
	if participantOfMe == nil {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
	}

	var idx int32
	var key string
	var ok bool
	var value []*msgService.PinnedMessage
	inboxPiece := impl.getPiece()
	chatPinnedMessageList := make([]*msgService.PinnedMessage, inboxPiece)
	pieceKeyMap := make(map[string][]*msgService.PinnedMessage, chat.GetChatInfo().Count)
	chat.GetChatInfo().Iteration(func(participant *data_chat.ChatParticipant) {
		if participant.UserId != 0 && participant.UserId != request.PinnedMessage.UserId && participant.State == data_chat.ParticipantState_ParticipantStateNormal {
			key = msgService.GetMsgByUserIdKey(participant.UserId)
			if value, ok = pieceKeyMap[key]; !ok {
				value = make([]*msgService.PinnedMessage, 0, 4)
			}
			pinnedMessage := *request.PinnedMessage
			pinnedMessage.UserId = participant.UserId
			value = append(value, &pinnedMessage)
			pieceKeyMap[key] = value
		}
	})
	for key, value = range pieceKeyMap {
		idx = 0
		for _, v := range value {
			chatPinnedMessageList[idx] = v
			idx++
			if idx >= inboxPiece {
				data, _ := jsoniter.Marshal(chatPinnedMessageList)
				_, _, err = impl.chatInProducer.SendMessage(constants.InboxChatMessagePinnedMessages, data)
				if err != nil {
					log.Warnf("ReqPinnedMessage chatId:%d SendMessage error:%s", request.PinnedMessage.PeerId, err.Error())
					return nil, err
				}
				idx = 0
			}
		}

		if idx > 0 {
			data, _ := jsoniter.Marshal(chatPinnedMessageList[:idx])
			_, _, err = impl.chatInProducer.SendMessage(constants.InboxChatMessagePinnedMessages, data)
			if err != nil {
				log.Warnf("ReqPinnedMessage chatId:%d SendMessage error:%s", request.PinnedMessage.PeerId, err.Error())
				return nil, err
			}
		}
	}
	return types.MarshalAny(mtproto.ToMTBool(true))
}
