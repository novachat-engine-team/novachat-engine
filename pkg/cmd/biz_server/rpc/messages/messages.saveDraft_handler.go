/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.saveDraft_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"time"
)

//  messages.saveDraft#bc39e14b flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> = Bool;
//
func (s *MessagesServiceImpl) MessagesSaveDraft(ctx context.Context, request *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesSaveDraft %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var err error
	inputPeer := input.MakeInputPeer(request.Peer)
	if inputPeer.GetPeerType() == constants.PeerTypeSelf {
		inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
			UserId:     constants.PeerTypeFromUserIDType(md.UserId).ToInt32(),
			AccessHash: request.Peer.AccessHash,
		}).To_InputPeer())
	}

	now := time.Now().Unix()
	switch inputPeer.GetPeerType() {
	case constants.PeerTypeUser, constants.PeerTypeSelf, constants.PeerTypeChat, constants.PeerTypeChannel:
		_ = request.NoWebpage
		_, err = s.accountMessageCore.SaveDraft(md.UserId, inputPeer.GetPeerId(), inputPeer.GetPeerType(),
			request.ReplyToMsgId,
			request.Message,
			request.Entities,
			int32(now),
		)
		if err != nil {
			log.Errorf("MessagesSaveDraft %v, request: %v ConversationEditDraftMessage error:%s", md, request, err.Error())
		}
	default:
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_NOT_SUPPORTED)
		log.Debugf("MessagesSaveDraft %v, request: %v invalid peer:%v error:%s", md, request, inputPeer.GetPeerType(), err.Error())
		return nil, err
	}

	go func() {
		var draftMessage *mtproto.DraftMessage
		if len(request.Message) == 0 && len(request.Entities) == 0 {
			draftMessage = mtproto.NewTLDraftMessageEmpty(&mtproto.DraftMessage{
				Date: int32(now),
			}).To_DraftMessage()
		} else {
			draftMessage = mtproto.NewTLDraftMessage(&mtproto.DraftMessage{
				Date:         int32(now),
				NoWebpage:    request.NoWebpage,
				ReplyToMsgId: request.ReplyToMsgId,
				Message:      request.Message,
				Entities:     request.Entities,
			}).To_DraftMessage()
		}

		updates := mtproto.NewTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{mtproto.NewTLUpdateDraftMessage(&mtproto.Update{
				Peer9961FD5C71: inputPeer.ToPeer(),
				Draft:          draftMessage,
			}).To_Update()},
			Users: []*mtproto.User{},
		})
		_, err = syncService.GetSyncClientById(md.UserId).ReqSyncUpdate(context.Background(),
			&syncService.SyncUpdate{
				UserId:          md.UserId,
				IgnoreAuthKeyId: md.AuthKeyId,
				Updates:         updates.To_Updates(),
				PeerType:        constants.PeerTypeUser.ToInt32(),
			})
		if err != nil {
			log.Errorf("MessagesSaveDraft %v, request: %v PushUpdates", md, request)
		}
	}()

	log.Debugf("MessagesSaveDraft %v, request: %v reply ok!!!!!!!!!!!!!!!!!!!", md, request)
	return mtproto.ToMTBool(true), nil
}
