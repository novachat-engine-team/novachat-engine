/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getAllDrafts_handler.go
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
	"time"
)

//  messages.getAllDrafts#6a3f8d65 = Updates;
//
func (s *MessagesServiceImpl) MessagesGetAllDrafts(ctx context.Context, request *mtproto.TLMessagesGetAllDrafts) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetAllDrafts %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//Impl MessagesGetAllDrafts logic
	draftList, err := s.accountMessageCore.GetAllDraft(md.UserId)
	if err != nil {
		log.Errorf("MessagesGetAllDrafts %v, request: %v ConversationPeerList error:%s", md, request, err.Error())
		return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
	}

	updates := make([]*mtproto.Update, 0, len(draftList))
	for _, c := range draftList {
		updates = append(updates, mtproto.NewTLUpdateDraftMessage(&mtproto.Update{
			Peer9961FD5C71: mtproto.NewTLPeerUser(&mtproto.Peer{
				UserId: constants.PeerTypeFromUserIDType(c.PeerId).ToInt32(),
			}).To_Peer(),
			Draft: c.Draft,
		}).To_Update())
	}

	ret := mtproto.NewTLUpdates(nil)
	ret.SetUpdates(updates)
	ret.SetUsers(nil)
	ret.SetSeq(0)
	ret.SetDate(int32(time.Now().Unix()))
	log.Debugf("MessagesGetAllDrafts %v, request: %v reply ok!!!!!!!!!!!!", md, request)
	return ret.To_Updates(), nil
}
