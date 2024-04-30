/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getMessageEditData_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
)

//  messages.getMessageEditData#fda68d36 peer:InputPeer id:int = messages.MessageEditData;
//
func (s *MessagesServiceImpl) MessagesGetMessageEditData(ctx context.Context, request *mtproto.TLMessagesGetMessageEditData) (*mtproto.Messages_MessageEditData, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesGetMessageEditData %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	inputPeer := input.MakeInputPeer(request.Peer)
	if inputPeer.GetPeerType() == constants.PeerTypeUser ||
		inputPeer.GetPeerType() == constants.PeerTypeSelf ||
		inputPeer.GetPeerType() == constants.PeerTypeChat || inputPeer.GetPeerType() == constants.PeerTypeChannel {
		l, err := s.accountMessageCore.GetMessages(md.UserId, []int32{request.Id}, false)
		if err != nil {
			log.Infof("MessagesGetMessageEditData - request: %v error:%s", request, err.Error())
			return nil, err
		}
		if len(l) == 0 {
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MESSAGE_ID_INVALID)
		}

		return mtproto.NewTLMessagesMessageEditData(&mtproto.Messages_MessageEditData{
			Caption: l[0].Out,
		}).To_Messages_MessageEditData(), nil
	} else {
		panic(fmt.Sprintf("MessagesGetMessageEditData %+v", inputPeer.GetPeerType()))
	}

	return nil, fmt.Errorf("%s", "Not impl MessagesGetMessageEditData")
}
