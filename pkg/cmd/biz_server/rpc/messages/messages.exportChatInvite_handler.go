/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.exportChatInvite_handler.go
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
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
)

//  messages.exportChatInvite#df7534c peer:InputPeer = ExportedChatInvite;
//  messages.exportChatInvite#7d885289 chat_id:int = ExportedChatInvite;
//
func (s *MessagesServiceImpl) MessagesExportChatInvite(ctx context.Context, request *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesExportChatInvite %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	chatId := constants.PeerTypeFromChatIDType32(request.ChatId).ToInt()
	if chatId == 0 {
		inputPeer := input.MakeInputPeer(request.Peer)
		if inputPeer.GetPeerType() == constants.PeerTypeChat || inputPeer.GetPeerType() == constants.PeerTypeChannel {
			chatId = inputPeer.GetPeerId()
		} else {
			err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
			log.Errorf("MessagesExportChatInvite %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
			return nil, err
		}
	}
	if chatId == 0 {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
		log.Errorf("MessagesExportChatInvite %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	fullChat, err := chatService.GetChatClientByKeyId(chatId).ReqFullChat(ctx, &chatService.FullChat{ChatId: chatId})
	if err != nil {
		log.Errorf("MessagesExportChatInvite %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	if fullChat.ChatData.Deleted {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
	}

	var ret *mtproto.ExportedChatInvite
	if len(fullChat.ChatData.Username) == 0 {
		ret = mtproto.NewTLChatInviteEmpty(nil).To_ExportedChatInvite()
	} else {
		ret = mtproto.NewTLChatInviteExported(&mtproto.ExportedChatInvite{
			Link: fullChat.ChatData.Username,
		}).To_ExportedChatInvite()
	}
	return ret, nil
}
