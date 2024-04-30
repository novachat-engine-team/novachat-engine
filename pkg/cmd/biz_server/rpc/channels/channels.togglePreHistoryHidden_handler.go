/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.togglePreHistoryHidden_handler.go
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

//  channels.togglePreHistoryHidden#eabbb94c channel:InputChannel enabled:Bool = Updates;
//
func (s *ChannelsServiceImpl) ChannelsTogglePreHistoryHidden(ctx context.Context, request *mtproto.TLChannelsTogglePreHistoryHidden) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsTogglePreHistoryHidden %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqHistoryHidden(ctx, &chatService.HistoryHidden{
		ChatId:    constants.PeerTypeFromUserIDType32(request.Channel.ChannelId).ToInt(),
		UserId:    md.UserId,
		AuthKeyId: md.AuthKeyId,
		Enabled:   mtproto.ToBool(request.Enabled),
		Date:      int32(time.Now().Unix()),
	})
	if err != nil {
		log.Warnf("ChannelsTogglePreHistoryHidden %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	log.Infof("ChannelsTogglePreHistoryHidden %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	err = types.UnmarshalAny(resp, &updates)
	return &updates, nil
}
