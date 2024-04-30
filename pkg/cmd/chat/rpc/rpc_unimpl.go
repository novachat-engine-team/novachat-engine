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
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
)

func (impl *Impl) ReqDeleteChat(ctx context.Context, request *chatService.DeleteChat) (*types.Any, error) {
	panic("ReqDeleteChat")
}

func (impl *Impl) ReqHistoryHidden(ctx context.Context, request *chatService.HistoryHidden) (*types.Any, error) {
	panic("ReqHistoryHidden")
}

func (impl *Impl) ReqUpdateUsername(ctx context.Context, request *chatService.UpdateUsername) (*types.Any, error) {
	panic("ReqUpdateUsername")
}

func (impl *Impl) ReqDeleteHistory(ctx context.Context, request *chatService.DeleteHistory) (*types.Any, error) {
	panic("ReqDeleteHistory")
}

func (impl *Impl) ReqDeleteMessages(ctx context.Context, request *chatService.DeleteMessages) (*types.Any, error) {
	panic("ReqDeleteMessages")
}

func (impl *Impl) ReqDeleteUserHistory(ctx context.Context, request *chatService.DeleteUserHistory) (*types.Any, error) {
	panic("ReqDeleteUserHistory")
}
