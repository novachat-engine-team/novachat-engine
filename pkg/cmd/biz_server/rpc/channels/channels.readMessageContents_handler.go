/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.readMessageContents_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
)

//  channels.ChannelsReadMessageContents#eab5dc38 channel:InputChannel id:Vector<int> = Bool;
//
func (s *ChannelsServiceImpl) ChannelsReadMessageContents(ctx context.Context, request *mtproto.TLChannelsReadMessageContents) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsReadMessageContents %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.Id) == 0 {
		return mtproto.ToMTBool(true), nil
	}

	_, err := s.accountMessageCore.ReadChannelMessageContents(
		md.AuthKeyId,
		md.UserId,
		constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt(),
		request.GetId())
	if err != nil {
		log.Errorf("ChannelsReadMessageContents request: %v error:%s", request, err.Error())
		return nil, err
	}
	return mtproto.ToMTBool(true), nil
}
