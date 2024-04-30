/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.leaveChannel_handler.go
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

//  channels.leaveChannel#f836aa95 channel:InputChannel = Updates;
//
func (s *ChannelsServiceImpl) ChannelsLeaveChannel(ctx context.Context, request *mtproto.TLChannelsLeaveChannel) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsLeaveChannel %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqBannedUser(ctx, &chatService.BannedUser{
		ChatId:    chatId,
		PeerId:    md.UserId,
		UserId:    md.UserId,
		IsLeft:    true,
		Date:      int32(time.Now().Unix()),
		AuthKeyId: md.AuthKeyId,
	})
	if err != nil {
		log.Errorf("ChannelsLeaveChannel %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	err = types.UnmarshalAny(resp, &updates)
	log.Infof("ChannelsLeaveChannel %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	return &updates, nil
}
