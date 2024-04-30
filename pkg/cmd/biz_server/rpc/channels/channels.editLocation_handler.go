/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.editLocation_handler.go
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
)

//  channels.editLocation#58e63f6d channel:InputChannel geo_point:InputGeoPoint address:string = Bool;
//
func (s *ChannelsServiceImpl) ChannelsEditLocation(ctx context.Context, request *mtproto.TLChannelsEditLocation) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsEditLocation %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	geopoint := chatService.GeoPoint{
		Lat:     0,
		Long:    0,
		Address: request.Address,
	}
	if request.GeoPoint != nil {
		geopoint.Long = request.GeoPoint.Long
		geopoint.Lat = request.GeoPoint.Lat
		_ = request.GeoPoint.AccuracyRadius
	}

	chatId := constants.PeerTypeFromChannelIDType32(request.Channel.ChannelId).ToInt()
	resp, err := chatService.GetChatClientByKeyId(chatId).ReqEditGeoPoint(ctx, &chatService.EditGeoPoint{
		PeerId:    chatId,
		UserId:    md.UserId,
		GeoPoint:  &geopoint,
		AuthKeyId: md.AuthKeyId,
	})
	if err != nil {
		log.Errorf("ChannelsEditLocation %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var ok mtproto.Bool
	err = types.UnmarshalAny(resp, &ok)
	log.Infof("ChannelsEditLocation %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	return &ok, nil
}
