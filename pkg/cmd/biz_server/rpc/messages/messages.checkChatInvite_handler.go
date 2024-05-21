/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.checkChatInvite_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/chat"
	"novachat_engine/service/common/errors"
)

//  messages.checkChatInvite#3eadb1bb hash:string = ChatInvite;
//
func (s *MessagesServiceImpl) MessagesCheckChatInvite(ctx context.Context, request *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesCheckChatInvite %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.Hash) == 0 {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_INVITE_HASH_EMPTY)
	}

	chatData, err := chatService.GetChatClientByKeyId(md.UserId).ReqChatByName(ctx, &chatService.ChatByName{
		Username: request.Hash,
		UserId:   md.UserId,
	})
	if err != nil {
		log.Errorf("MessagesCheckChatInvite %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	participantUserIdList := make([]int64, 0, len(chatData.ParticipantList))
	found := false
	for _, v := range chatData.ParticipantList {
		if v.UserId == md.UserId {
			found = true
		}
		participantUserIdList = append(participantUserIdList, v.UserId)
	}
	if found {
		return mtproto.NewTLChatInviteAlready(&mtproto.ChatInvite{
			Chat: chat.ToChat(md.UserId, chatData, md.Layer),
		}).To_ChatInvite(), nil
	}
	//TODO:(Coderxw) peeking into the group to read messages without joining it.
	//if s.conf.PeerChat.Enable {
	//	mtproto.NewTLChatInvitePeek(&mtproto.ChatInvite{Chat: chat.ToChat(md.UserId, chatData), Expires: s.conf.PeerChat.Expires})
	//}
	retChat := chat.ToChat(md.UserId, chatData, md.Layer)

	userList, err := s.accountUsersCore.GetUserList(md.UserId, participantUserIdList, md.Layer)
	if err != nil {
		log.Warnf("MessagesCheckChatInvite GetUserList error:%s", err.Error())
	}
	//TODO:(Coderxw) Channel

	return mtproto.NewTLChatInvite(&mtproto.ChatInvite{
		Channel:           false,
		Broadcast:         false,
		Public:            false,
		Megagroup:         false,
		Title:             retChat.Title,
		PhotoDB74F55871:   retChat.Photo,
		ParticipantsCount: retChat.ParticipantsCount,
		Participants:      userList,
		//PhotoDFC2F58E98:   retChat.Photo.PhotoSmall,
	}).To_ChatInvite(), nil
}
