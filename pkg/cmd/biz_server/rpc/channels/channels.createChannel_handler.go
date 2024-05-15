/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.createChannel_handler.go
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
	"novachat_engine/service/app/help"
	"novachat_engine/service/common/errors"
	"strings"
	"time"
)

//  channels.createChannel#3d5fb10f flags:# broadcast:flags.0?true megagroup:flags.1?true title:string about:string geo_point:flags.2?InputGeoPoint address:flags.2?string = Updates;
//  channels.createChannel#f4893d7f flags:# broadcast:flags.0?true megagroup:flags.1?true title:string about:string = Updates;
//
func (s *ChannelsServiceImpl) ChannelsCreateChannel(ctx context.Context, request *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ChannelsCreateChannel %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	tlUser, err := s.accountUsersCore.GetUser(md.UserId, md.UserId)
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
		log.Errorf("MessagesCreateChat %v, request: %v GetByUserId error:%s", md, request, err.Error())
		return nil, err
	}
	if tlUser.Restricted || tlUser.Deleted {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_UNREGISTERED)
		return nil, err
	}

	title := strings.TrimSpace(request.Title)
	if len(title) == 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_TITLE_EMPTY)
		log.Errorf("MessagesCreateChat %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var geoPoint *chatService.GeoPoint
	if request.GeoPoint != nil {
		geoPoint = &chatService.GeoPoint{
			Lat:     request.GeoPoint.Lat,
			Long:    request.GeoPoint.Long,
			Address: request.Address,
		}
	}

	createChatResp, err := chatService.GetChatClientByKeyId(md.UserId).ReqCreateChat(context.Background(),
		&chatService.CreateChat{
			UserId:      md.UserId,
			AuthKeyId:   md.AuthKeyId,
			PeerId:      nil,
			Title:       title,
			ChatSizeMax: help.GetConfig().ChatSizeMax,
			Date:        int32(time.Now().Unix()),
			GeoPoint:    geoPoint,
			Layer:       md.Layer,
		})
	if err != nil {
		log.Errorf("MessagesCreateChat %v, request: %v CreateChat error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	types.UnmarshalAny(createChatResp, &updates)
	log.Infof("MessagesCreateChat %v, request: %v OK!!!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request)
	return &updates, nil
}
