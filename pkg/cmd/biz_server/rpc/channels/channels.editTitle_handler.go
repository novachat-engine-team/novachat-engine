/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.editTitle_handler.go
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
	"novachat_engine/service/common/errors"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	"strings"
	"time"
)

//  channels.editTitle#566decd0 channel:InputChannel title:string = Updates;
//
func (s *ChannelsServiceImpl) ChannelsEditTitle(ctx context.Context, request *mtproto.TLChannelsEditTitle) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsEditTitle %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	title := strings.TrimSpace(request.Title)
	if len(title) == 0 {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_TITLE_EMPTY)
		log.Errorf("ChannelsEditTitle %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	chatId := constants.PeerTypeFromChatIDType32(request.Channel.ChannelId).ToInt()
	if chatId == 0 {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
		log.Errorf("ChannelsEditTitle %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	editTitleResp, err := chatService.GetChatClientByKeyId(chatId).ReqEditTitle(ctx, &chatService.EditTitle{
		ChatId:    chatId,
		Title:     request.Title,
		UserId:    md.UserId,
		AuthKeyId: md.AuthKeyId,
		Date:      int32(time.Now().Unix()),
	})
	if err != nil {
		log.Errorf("ChannelsEditTitle %v, request: %v ReqEditTitle error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	var updates mtproto.Updates
	types.UnmarshalAny(editTitleResp, &updates)
	return compat.UpdatesCompat(&updates, md.Layer), nil
}
