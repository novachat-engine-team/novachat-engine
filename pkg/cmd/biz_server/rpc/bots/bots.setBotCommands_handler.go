/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : bots.setBotCommands_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  bots.setBotCommands#805d46f6 commands:Vector<BotCommand> = Bool;
//
func (s *BotsServiceImpl) BotsSetBotCommands(ctx context.Context, request *mtproto.TLBotsSetBotCommands) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("BotsSetBotCommands %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl BotsSetBotCommands logic

	return nil, fmt.Errorf("%s", "Not impl BotsSetBotCommands")
}
