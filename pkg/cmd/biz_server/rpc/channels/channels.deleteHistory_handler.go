/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.deleteHistory_handler.go
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
	"time"
)

//  channels.deleteHistory#af369d42 channel:InputChannel max_id:int = Bool;
//
func (s *ChannelsServiceImpl) ChannelsDeleteHistory(ctx context.Context, request *mtproto.TLChannelsDeleteHistory) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsDeleteHistory %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqDeleteHistory(ctx, &chatService.DeleteHistory{
		ChatId:    chatId,
		UserId:    md.UserId,
		AuthKeyId: md.AuthKeyId,
		MaxId:     request.MaxId,
		Date:      int32(time.Now().Unix()),
	})
	if err != nil {
		log.Warnf("ChannelsDeleteHistory %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var ok mtproto.Bool
	err = types.UnmarshalAny(resp, &ok)
	return &ok, nil
}
