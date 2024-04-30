/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.editChatAbout_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"strings"
	"time"
)

//  messages.editChatAbout#def60797 peer:InputPeer about:string = Bool;
//
func (s *MessagesServiceImpl) MessagesEditChatAbout(ctx context.Context, request *mtproto.TLMessagesEditChatAbout) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesEditChatAbout %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	about := strings.TrimSpace(request.About)
	//if len(about) == 0 {
	//	err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_TITLE_EMPTY)
	//	log.Errorf("MessagesEditChatAbout %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//	return nil, err
	//}

	var err error
	inputPeer := input.MakeInputPeer(request.Peer)
	switch inputPeer.GetPeerType() {
	case constants.PeerTypeChannel, constants.PeerTypeChat:
		_, err = chatService.GetChatClientByKeyId(inputPeer.GetPeerId()).ReqEditAbout(ctx, &chatService.EditAbout{
			ChatId:    inputPeer.GetPeerId(),
			About:     about,
			UserId:    md.UserId,
			AuthKeyId: md.AuthKeyId,
			Date:      int32(time.Now().Unix()),
		})
	default:
		panic(fmt.Sprintf("MessagesEditChatAbout %v, request: %v peerType:%v", metadata.RpcMetaDataDebug(md), request, inputPeer.GetPeerType()))
	}

	if err != nil {
		log.Errorf("MessagesEditChatAbout %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	}
	return mtproto.ToMTBool(true), err
}
