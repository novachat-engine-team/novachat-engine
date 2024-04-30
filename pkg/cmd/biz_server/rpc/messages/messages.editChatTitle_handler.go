/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.editChatTitle_handler.go
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

//  messages.editChatTitle#dc452855 chat_id:int title:string = Updates;
//
func (s *MessagesServiceImpl) MessagesEditChatTitle(ctx context.Context, request *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesEditChatTitle %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	title := strings.TrimSpace(request.Title)
	if len(title) == 0 {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_TITLE_EMPTY)
		log.Errorf("MessagesEditChatTitle %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	chatId := constants.PeerTypeFromChatIDType32(request.ChatId).ToInt()
	if chatId == 0 {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
		log.Errorf("MessagesEditChatTitle %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
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
		log.Errorf("MessagesEditChatTitle %v, request: %v ReqEditTitle error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	var updates mtproto.Updates
	types.UnmarshalAny(editTitleResp, &updates)
	return compat.UpdatesCompat(&updates, md.Layer), nil
}
