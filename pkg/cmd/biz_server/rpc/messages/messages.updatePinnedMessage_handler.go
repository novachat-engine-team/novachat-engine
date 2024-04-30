/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.updatePinnedMessage_handler.go
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
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
)

//  messages.updatePinnedMessage#d2aaf7ec flags:# silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer:InputPeer id:int = Updates;
//
func (s *MessagesServiceImpl) MessagesUpdatePinnedMessage(ctx context.Context, request *mtproto.TLMessagesUpdatePinnedMessage) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesUpdatePinnedMessage %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	unpin := request.Flags&(1<<1) != 0
	pmOneSide := request.Flags&(1<<2) != 0

	var updates *mtproto.Updates
	var err error
	inputPeer := input.MakeInputPeer(request.Peer)
	switch inputPeer.GetPeerType() {
	case constants.PeerTypeUser:
		updates, err = s.accountMessageCore.PinnedMessage(md.UserId, md.AuthKeyId, inputPeer, request.Id, unpin, pmOneSide)
	case constants.PeerTypeChat, constants.PeerTypeChannel:
		var resp *types.Any
		resp, err = chatService.GetChatClientByKeyId(inputPeer.GetPeerId()).ReqPinnedMessage(context.Background(),
			&chatService.ChatPinnedMessage{
				PinnedMessage: &msgService.PinnedMessage{
					UserId:    md.UserId,
					PeerId:    inputPeer.GetPeerId(),
					PeerType:  inputPeer.GetPeerType().ToInt32(),
					MessageId: request.Id,
					Unpin:     unpin,
					OneSide:   pmOneSide,
					AuthKeyId: md.AuthKeyId,
				},
			})
		if err == nil {
			updates = &mtproto.Updates{}
			err = types.UnmarshalAny(resp, updates)
		}
	default:
		err = errorsService.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST, fmt.Errorf("UpdatePinnedMessage peerType:%d error", inputPeer.GetPeerType()).Error())
	}
	if err != nil {
		log.Errorf("MessagesUpdatePinnedMessage request: %v error:%s", request, err.Error())
		return nil, err
	}

	return updates, nil
}
