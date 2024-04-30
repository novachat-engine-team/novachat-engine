/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.deleteChannel_handler.go
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

//  channels.deleteChannel#c0111fe3 channel:InputChannel = Updates;
//
func (s *ChannelsServiceImpl) ChannelsDeleteChannel(ctx context.Context, request *mtproto.TLChannelsDeleteChannel) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsDeleteChannel %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqDeleteChat(ctx, &chatService.DeleteChat{
		ChatId:    chatId,
		UserId:    md.UserId,
		AuthKeyId: md.AuthKeyId,
		Date:      int32(time.Now().Unix()),
	})
	if err != nil {
		log.Warnf("ChannelsDeleteChannel %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	err = types.UnmarshalAny(resp, &updates)
	return &updates, nil
}
