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
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
)

func (impl *Impl) ReqCheckUsername(ctx context.Context, request *chatService.CheckUsername) (*types.Any, error) {
	chat, err := impl.chatManager.GetChatByName(request.GetUsername())
	if err != nil {
		log.Errorf("ReqChatByName error:%s", err.Error())
		return nil, err
	}
	if chat.ChatId == 0 {
		return types.MarshalAny(mtproto.ToMTBool(true))
	}
	return types.MarshalAny(mtproto.ToMTBool(chat.ChatId != request.PeerId))
}
