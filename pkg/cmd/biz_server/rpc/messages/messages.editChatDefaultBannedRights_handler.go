/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.editChatDefaultBannedRights_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/banned_right"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"time"
)

//  messages.editChatDefaultBannedRights#a5866b41 peer:InputPeer banned_rights:ChatBannedRights = Updates;
//
func (s *MessagesServiceImpl) MessagesEditChatDefaultBannedRights(ctx context.Context, request *mtproto.TLMessagesEditChatDefaultBannedRights) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesEditChatDefaultBannedRights %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	rights, utilDate := banned_right.ChatBannedRightToRights(request.BannedRights)
	var err error
	var bannedRightsResp *types.Any
	inputPeer := input.MakeInputPeer(request.Peer)
	switch inputPeer.GetPeerType() {
	case constants.PeerTypeChannel, constants.PeerTypeChat:
		bannedRightsResp, err = chatService.GetChatClientByKeyId(inputPeer.GetPeerId()).ReqBannedRights(ctx, &chatService.BannedRights{
			ChatId:    inputPeer.GetPeerId(),
			Rights:    rights,
			UtilDate:  utilDate,
			UserId:    md.UserId,
			AuthKeyId: md.AuthKeyId,
			Date:      int32(time.Now().Unix()),
		})
	default:
		panic(fmt.Sprintf("MessagesEditChatAbout %v, request: %v peerType:%v", metadata.RpcMetaDataDebug(md), request, inputPeer.GetPeerType()))
	}

	if err != nil {
		log.Errorf("MessagesEditChatAbout %v, request: %v ReqBannedRights error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	var updates mtproto.Updates
	types.UnmarshalAny(bannedRightsResp, &updates)
	return compat.UpdatesCompat(&updates, md.Layer), nil
}
