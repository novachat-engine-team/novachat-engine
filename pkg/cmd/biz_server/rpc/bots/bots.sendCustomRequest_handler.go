/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : bots.sendCustomRequest_handler.go
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

//  bots.sendCustomRequest#aa2769ed custom_method:string params:DataJSON = DataJSON;
//
func (s *BotsServiceImpl) BotsSendCustomRequest(ctx context.Context, request *mtproto.TLBotsSendCustomRequest) (*mtproto.DataJSON, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("BotsSendCustomRequest %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl BotsSendCustomRequest logic

	return nil, fmt.Errorf("%s", "Not impl BotsSendCustomRequest")
}
