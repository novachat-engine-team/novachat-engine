/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.readHistory_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"time"
)

//  messages.readHistory#e306d3a peer:InputPeer max_id:int = messages.AffectedMessages;
//  messages.readHistory#b04f2510 peer:InputPeer max_id:int offset:int = messages.AffectedHistory;
//
func (s *MessagesServiceImpl) MessagesReadHistory(ctx context.Context, request *mtproto.TLMessagesReadHistory) (*mtproto.Response_MessagesReadHistory, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesReadHistory %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	var (
		err              error
		affectedMessages *mtproto.Messages_AffectedMessages
		affectedHistory  *mtproto.Messages_AffectedHistory
	)
	inputPeer := input.MakeInputPeer(request.Peer)
	if inputPeer.GetPeerType() == constants.PeerTypeSelf {
		inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
			UserId:     constants.PeerTypeFromUserIDType(md.UserId).ToInt32(),
			AccessHash: request.Peer.AccessHash,
		}).To_InputPeer())
	}

	switch inputPeer.GetPeerType() {
	case constants.PeerTypeUser, constants.PeerTypeSelf, constants.PeerTypeChat, constants.PeerTypeChannel:
		//var ok bool
		affectedHistory, err = s.accountMessageCore.ReadHistory(md.UserId, md.AuthKeyId, request.MaxId, inputPeer, int32(time.Now().Unix()))
		if err != nil {
			err = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
			log.Errorf("MessagesReadHistory %v, request: %v ReadHistory error:%s", md, request, err.Error())
			return nil, err
		}
	default:
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_NOT_SUPPORTED)
		log.Debugf("MessagesReadHistory %v, request: %v invalid peer:%v error:%s", md, request, inputPeer.GetPeerType(), err.Error())
		return nil, err
	}
	affectedMessages = mtproto.NewTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      affectedHistory.Pts,
		PtsCount: affectedHistory.PtsCount,
	}).To_Messages_AffectedMessages()

	resp := mtproto.NewResponse_MessagesReadHistory()
	if md.Layer >= 71 {
		resp.Cmd = mtproto.Cmd_MessagesReadHistory.ToUInt32()
		resp.MessagesReadHistorye306D3A = affectedMessages
	} else {
		resp.Cmd = mtproto.Cmd_MessagesReadHistoryb04f2510.ToUInt32()
		resp.MessagesReadHistoryb04F2510 = affectedHistory
	}
	log.Debugf("MessagesReadHistory %v, request: %v reply ok!!!!!!!!!!!!", md, request)
	return resp, nil
}
