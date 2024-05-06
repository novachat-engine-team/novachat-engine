/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.deleteUserHistory_handler.go
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

//  channels.deleteUserHistory#d10dd71b channel:InputChannel user_id:InputUser = messages.AffectedHistory;
//
func (s *ChannelsServiceImpl) ChannelsDeleteUserHistory(ctx context.Context, request *mtproto.TLChannelsDeleteUserHistory) (*mtproto.Messages_AffectedHistory, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsDeleteUserHistory %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	//msgService.GetMsgClient().
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqDeleteUserHistory(ctx, &chatService.DeleteUserHistory{
		ChatId:    chatId,
		UserId:    md.UserId,
		AuthKeyId: md.AuthKeyId,
		PeerId:    constants.PeerTypeFromUserIDType32(request.UserId.UserId).ToInt(),
		Date:      int32(time.Now().Unix()),
	})
	if err != nil {
		log.Warnf("ChannelsDeleteMessages %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var pts mtproto.TLInt32
	err = types.UnmarshalAny(resp, &pts)
	ptsCount := int32(0)
	if pts.Value > 0 {
		ptsCount = 1
	}
	affectedHistory := mtproto.NewTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
		Pts:      pts.Value,
		PtsCount: ptsCount,
		Offset:   0,
	}).To_Messages_AffectedHistory()
	return affectedHistory, nil
}
