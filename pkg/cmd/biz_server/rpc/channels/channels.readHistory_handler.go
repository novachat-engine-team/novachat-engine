/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.readHistory_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"time"
)

//  channels.readHistory#cc104937 channel:InputChannel max_id:int = Bool;
//
func (s *ChannelsServiceImpl) ChannelsReadHistory(ctx context.Context, request *mtproto.TLChannelsReadHistory) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsReadHistory %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	inputPeer := input.MakeInputPeerValue(chatId, constants.PeerTypeChannel)
	affectedHistory, err := s.accountMessageCore.ReadHistory(md.UserId, md.AuthKeyId, request.MaxId, inputPeer, int32(time.Now().Unix()))
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
		log.Errorf("ChannelsReadHistory %v, request: %v ReadHistory error:%s", md, request, err.Error())
		return nil, err
	}

	_ = affectedHistory
	return mtproto.ToMTBool(true), nil
}
