/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.editAbout_handler.go
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
	"strings"
	"time"
)

//  channels.editAbout#13e27f1e channel:InputChannel about:string = Bool;
//
func (s *ChannelsServiceImpl) ChannelsEditAbout(ctx context.Context, request *mtproto.TLChannelsEditAbout) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsEditAbout %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	about := strings.TrimSpace(request.About)
	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqEditAbout(ctx, &chatService.EditAbout{
		ChatId:    chatId,
		About:     about,
		UserId:    md.UserId,
		AuthKeyId: md.AuthKeyId,
		Date:      int32(time.Now().Unix()),
	})
	if err != nil {
		log.Warnf("ChannelsEditAbout %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	log.Infof("ChannelsEditAbout %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	var ok mtproto.Bool
	err = types.UnmarshalAny(resp, &ok)
	return &ok, err
}
