/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.updateUsername_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
)

//  channels.updateUsername#3514b3de channel:InputChannel username:string = Bool;
//
func (s *ChannelsServiceImpl) ChannelsUpdateUsername(ctx context.Context, request *mtproto.TLChannelsUpdateUsername) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsUpdateUsername %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqUpdateUsername(ctx, &chatService.UpdateUsername{
		ChatId:    chatId,
		UserId:    md.UserId,
		AuthKeyId: md.AuthKeyId,
		Username:  request.Username,
	})
	if err != nil {
		log.Errorf("ChannelsUpdateUsername %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var ok mtproto.Bool
	err = types.UnmarshalAny(resp, &ok)
	log.Infof("ChannelsUpdateUsername %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	return &ok, nil
}
