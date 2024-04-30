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
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
)

func (impl *Impl) ReqChatCommon(ctx context.Context, request *chatService.ChatCommon) (*chatService.ChatList, error) {
	chatCommonList, err := impl.chatManager.GetCommonChat(request.UserId)
	if err != nil {
		log.Errorf("ReqChatCommon error:%s", err.Error())
		return nil, err
	}

	resp := &chatService.ChatList{Values: make([]*chatService.Chat, 0, len(chatCommonList))}

	for _, v := range chatCommonList {
		resp.Values = append(resp.Values, &chatService.Chat{ChatData: v})
	}

	return resp, nil
}
