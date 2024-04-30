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
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/util"
	"novachat_engine/service/common/errors"
	data_chat "novachat_engine/service/data/chat"
)

func (impl *Impl) ReqChatByName(ctx context.Context, request *chatService.ChatByName) (*chatService.Chat, error) {
	chat, err := impl.chatManager.GetChatByName(request.GetUsername())
	if err != nil {
		log.Errorf("ReqChatByName error:%s", err.Error())
		return nil, err
	}
	if chat.ChatId == 0 {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
	}

	chatResp, err := chatService.GetChatClientByKeyId(chat.ChatId).ReqFullChat(ctx, &chatService.FullChat{ChatId: chat.ChatId})
	if err != nil {
		log.Errorf("ReqChatByName ReqFullChat chatId:%d error:%s", chat.ChatId, err.Error())
		return nil, err
	}

	if idx := util.Index(chatResp.ParticipantList, func(idx int) bool {
		return chatResp.ParticipantList[idx].UserId == request.UserId
	}); idx >= 0 {
		if chatResp.ParticipantList[idx].State == data_chat.ParticipantState_ParticipantStateBan {
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_INVITE_HASH_INVALID)
		}
	}

	log.Debugf("ReqChatByName username:%d chatResp:%+v", request.Username, chatResp)
	return chatResp, nil
}
