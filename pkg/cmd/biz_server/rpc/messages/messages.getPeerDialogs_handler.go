/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getPeerDialogs_handler.go
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
	"novachat_engine/service/compat"
)

//  messages.getPeerDialogs#e470bcfd peers:Vector<InputDialogPeer> = messages.PeerDialogs;
//  messages.getPeerDialogs#2d9776b9 peers:Vector<InputPeer> = messages.PeerDialogs;
//
func (s *MessagesServiceImpl) MessagesGetPeerDialogs(ctx context.Context, request *mtproto.TLMessagesGetPeerDialogs) (*mtproto.Messages_PeerDialogs, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetPeerDialogs %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var foldId int32
	var inputPeerList []*mtproto.InputPeer
	if len(request.PeersE470BCFD85) > 0 {
		inputPeerList = make([]*mtproto.InputPeer, 0, len(request.PeersE470BCFD85))
		for _, v := range request.PeersE470BCFD85 {
			inputPeerList = append(inputPeerList, v.Peer)

			if v.FolderId != 0 {
				foldId = v.FolderId
			}
		}
	} else {
		inputPeerList = make([]*mtproto.InputPeer, 0, len(request.Peers2D9776B971))
		for _, p := range request.Peers2D9776B971 {
			inputPeerList = append(inputPeerList, p)
		}
	}

	state := mtproto.NewTLUpdatesState(nil).To_Updates_State()
	messagePeerDialogs, err := s.accountMessageCore.GetPeerDialog(md.UserId, inputPeerList, foldId, md.Layer)
	if err != nil {
		log.Errorf("MessagesGetPeerDialogs %v, request: %v GetDialog error:%s", md, request, err.Error())
		return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
	}

	for idx := range messagePeerDialogs.Dialogs {
		messagePeerDialogs.Dialogs[idx] = compat.CompatDialog(messagePeerDialogs.Dialogs[idx], md.Layer)
	}

	for idx := range messagePeerDialogs.Messages {
		messagePeerDialogs.Messages[idx] = compat.MessageCompat(messagePeerDialogs.Messages[idx], md.Layer)
	}

	for idx := range messagePeerDialogs.Users {
		messagePeerDialogs.Users[idx] = compat.CompatUser(messagePeerDialogs.Users[idx], md.Layer)
	}

	for idx := range messagePeerDialogs.Chats {
		messagePeerDialogs.Chats[idx] = compat.CompatChat(messagePeerDialogs.Chats[idx], md.Layer)
	}

	messagePeerDialogs.State = state
	log.Debugf("MessagesGetPeerDialogs userId:%d, inputDialogPeerList:%v UpdateGetState messagePeerDialogs:(%+v)", md.UserId,
		inputPeerList, messagePeerDialogs)
	return messagePeerDialogs, nil
}
