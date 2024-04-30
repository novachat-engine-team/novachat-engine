/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.importChatInvite_handler.go
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
	"novachat_engine/service/compat"
	"time"
)

//  messages.importChatInvite#6c50051c hash:string = Updates;
//
func (s *MessagesServiceImpl) MessagesImportChatInvite(ctx context.Context, request *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesImportChatInvite %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.Hash) == 0 {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_INVITE_HASH_EMPTY)
	}

	chatData, err := chatService.GetChatClientByKeyId(md.UserId).ReqChatByName(ctx, &chatService.ChatByName{
		Username: request.Hash,
		UserId:   md.UserId,
	})
	if err != nil {
		log.Errorf("MessagesImportChatInvite %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	if chatData.ChatData.Deleted {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_RESTRICTED)
		log.Warnf("MessagesImportChatInvite %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	if util.Index(chatData.ParticipantList, func(idx int) bool {
		return chatData.ParticipantList[idx].UserId == md.UserId
	}) >= 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USER_ALREADY_PARTICIPANT)
		log.Warnf("MessagesImportChatInvite %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	addUserResp, err := chatService.GetChatClientByKeyId(chatData.ChatData.ChatId).ReqAddUser(ctx, &chatService.AddUser{
		UserId:    md.UserId,
		ChatId:    chatData.ChatData.ChatId,
		PeerId:    md.UserId,
		FwdLimit:  0,
		AuthKeyId: md.AuthKeyId,
		Date:      int32(time.Now().Unix()),
	})

	var updates mtproto.Updates
	types.UnmarshalAny(addUserResp, &updates)
	log.Infof("MessagesImportChatInvite %v, request: %v ok!!!!!!!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request)
	return compat.UpdatesCompat(&updates, md.Layer), nil
}
