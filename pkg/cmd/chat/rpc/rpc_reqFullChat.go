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
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/cmd/chat/service"
	errorsService "novachat_engine/service/common/errors"
	data_chat "novachat_engine/service/data/chat"
)

func (impl *Impl) ReqFullChat(ctx context.Context, request *chatService.FullChat) (*chatService.Chat, error) {
	chat, err := impl.chatManager.GetChat(request.ChatId)
	if err != nil {
		if errors.Is(err, service.DefaultErrChatIdInvalid) {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
		} else {
			return nil, err
		}
	}

	if chat.Invalid() || chat.Deleted() {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
	}

	participantList := make([]*data_chat.ChatParticipant, 0, chat.GetChatInfo().Count)
	chat.GetChatInfo().Iteration(func(participant *data_chat.ChatParticipant) {
		participantList = append(participantList, participant)
	})
	return &chatService.Chat{
		ChatData:        chat.GetChatInfo().ChatData,
		ParticipantList: participantList,
	}, nil
}
