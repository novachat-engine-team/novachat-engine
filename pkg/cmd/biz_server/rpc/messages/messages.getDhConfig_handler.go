/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getDhConfig_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/dh_config"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
)

//  messages.getDhConfig#26cf8950 version:int random_length:int = messages.DhConfig;
//
func (s *MessagesServiceImpl) MessagesGetDhConfig(ctx context.Context, request *mtproto.TLMessagesGetDhConfig) (*mtproto.Messages_DhConfig, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetDhConfig %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	dhConfig := mtproto.NewTLMessagesDhConfig(&mtproto.Messages_DhConfig{
		Random:  crypto.RandomBytes(int(request.RandomLength)),
		G:       int32(dh_config.DH2048G[0]),
		P:       dh_config.DH2048P,
		Version: constants.CurrentDhConfigVersion,
	})

	log.Debugf("MessagesGetDhConfig %v, request: %v reply ok!!!!!!!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request)
	return dhConfig.To_Messages_DhConfig(), nil
}
