/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.checkUsername_handler.go
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
	"novachat_engine/pkg/util"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
)

//  channels.checkUsername#10e6bd2c channel:InputChannel username:string = Bool;
//
func (s *ChannelsServiceImpl) ChannelsCheckUsername(ctx context.Context, request *mtproto.TLChannelsCheckUsername) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)

	var err error
	inputPeer := input.MakeInputPeerValue(constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt(), constants.PeerTypeChannel)
	if request.Username == "1" {
		//TODO:(Coderxw) CHANNELS_ADMIN_PUBLIC_TOO_MUCH
		// check admin public channels
		return mtproto.ToMTBool(true), nil
	}

	if !util.IsAllNumStringLimit(request.Username, util.LimitUserNameMaxLength) {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USERNAME_INVALID)
		log.Errorf("ChannelsCheckUsername %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	resp, err := chatService.GetChatClientByKeyId(inputPeer.GetPeerId()).ReqCheckUsername(ctx, &chatService.CheckUsername{
		PeerId:   inputPeer.GetPeerId(),
		PeerType: inputPeer.GetPeerType().ToInt32(),
		Username: request.Username,
	})
	if err != nil {
		log.Warnf("ChannelsCheckUsername %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var ok mtproto.Bool
	types.UnmarshalAny(resp, &ok)
	return &ok, err
}
