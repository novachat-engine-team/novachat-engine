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
