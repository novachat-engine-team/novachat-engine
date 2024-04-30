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
	sessionService "novachat_engine/pkg/cmd/session/rpc_client"
	"novachat_engine/pkg/log"
)

func (m *SessionImpl) Close(ctx context.Context, closeEvent *sessionService.CloseEvent) (*types.Any, error) {

	err := m.messageChannel.CloseSession(closeEvent)
	if err != nil {
		log.Errorf("SessionImpl error:%s", err.Error())
		return types.MarshalAny(mtproto.ToMTBool(false))
	}

	return types.MarshalAny(mtproto.ToMTBool(true))
}
