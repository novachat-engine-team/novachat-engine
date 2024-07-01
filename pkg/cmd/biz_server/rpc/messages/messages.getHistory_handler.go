/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getHistory_handler.go
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
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"sort"
	"time"
)

//  messages.getHistory#dcbb8260 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:int = messages.Messages;
//  messages.getHistory#afa92846 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
//
func (s *MessagesServiceImpl) MessagesGetHistory(ctx context.Context, request *mtproto.TLMessagesGetHistory) (*mtproto.Messages_Messages, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetHistory %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	var (
		err             error
		messageMessages *mtproto.Messages_Messages
	)
	inputPeer := input.MakeInputPeer(request.Peer)
	if inputPeer.GetPeerType() == constants.PeerTypeSelf {
		inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
			UserId:     constants.PeerTypeFromUserIDType(md.UserId).ToInt32(),
			AccessHash: request.Peer.AccessHash,
		}).To_InputPeer())
	}

	switch inputPeer.GetPeerType() {
	case constants.PeerTypeUser, constants.PeerTypeSelf:
	case constants.PeerTypeChat:
	case constants.PeerTypeChannel:
	default:
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_NOT_SUPPORTED)
		log.Debugf("MessagesGetHistory %v, request: %v invalid peer:%v error:%s", md, request, inputPeer.GetPeerType(), err.Error())
		return nil, err
	}

	messageMessages, err = s.accountMessageCore.GetHistory(md.UserId, inputPeer, request.OffsetId, request.OffsetDate, request.AddOffset, request.MinId, request.MaxId, request.Limit, request.Hash, md.Layer)
	if err != nil {
		log.Errorf("MessagesGetHistory %v, request: %v GetHistory error:%s", md, request, err.Error())
		return nil, err
	}

	for idx := range messageMessages.Users {
		if messageMessages.Users[idx] != nil && messageMessages.Users[idx].Photo != nil {
			messageMessages.Users[idx].Photo.PhotoSmall = compat.CompatFileLocation(messageMessages.Users[idx].Photo.PhotoSmall, md.Layer)

			messageMessages.Users[idx].Photo.PhotoBig = compat.CompatFileLocation(messageMessages.Users[idx].Photo.PhotoBig, md.Layer)
		}
	}

	for idx := range messageMessages.Chats {
		if messageMessages.Chats[idx] != nil && messageMessages.Chats[idx].Photo != nil {
			messageMessages.Chats[idx].Photo.PhotoSmall = compat.CompatFileLocation(messageMessages.Chats[idx].Photo.PhotoSmall, md.Layer)
			messageMessages.Chats[idx].Photo.PhotoBig = compat.CompatFileLocation(messageMessages.Chats[idx].Photo.PhotoBig, md.Layer)
		}
	}

	sort.SliceStable(messageMessages.Messages, func(i, j int) bool {
		return messageMessages.Messages[i].Id > messageMessages.Messages[j].Id
	})

	idList := make([]int32, 0, len(messageMessages.Messages))
	for _, v := range messageMessages.Messages {
		idList = append(idList, v.Id)
	}

	if md.LangPack == constants.DevicePlatformTypePC.ToString() {
		if len(messageMessages.Messages) > 0 {
			_, err = s.accountMessageCore.ReadHistory(md.UserId, md.AuthKeyId, request.MaxId, inputPeer, int32(time.Now().Unix()))
			if err != nil {
				log.Warnf("MessagesGetHistory ReadHistory error:%s", err.Error())
			}
		}
	}

	log.Debugf("MessagesGetHistory %v, request: %v idList:%+v ok!!!!!!!!!!!!", md, request, idList)
	return messageMessages, nil
}
