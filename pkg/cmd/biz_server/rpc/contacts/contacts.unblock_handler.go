/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.unblock_handler.go
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
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"time"
)

//  contacts.unblock#bea65d50 id:InputPeer = Bool;
//  contacts.unblock#e54100bd id:InputUser = Bool;
//
func (s *ContactsServiceImpl) ContactsUnblock(ctx context.Context, request *mtproto.TLContactsUnblock) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsUnblock %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if request.IdE54100BD71 != nil {
		request.IdBEA65D50119 = mtproto.NewTLInputPeerUser(&mtproto.InputPeer{UserId: request.IdE54100BD71.UserId}).To_InputPeer()
	}

	inputPeer := input.MakeInputPeer(request.IdBEA65D50119)
	if inputPeer.GetPeerType() != constants.PeerTypeUser && inputPeer.GetPeerType() != constants.PeerTypeSelf {
		log.Errorf("ContactsUnblock - request: %v inputPeer type error:%v", request, inputPeer.GetPeerType())
		return mtproto.ToMTBool(false), nil
	}

	ok, err := s.accountContactCore.ContactsUnblock(md.UserId, inputPeer.GetPeerId(), int32(time.Now().Unix()))
	if err != nil {
		log.Errorf("ContactsUnblock - request: %v error:%s", request, err.Error())
		return nil, err
	}

	go func() {
		_, err = syncService.GetSyncClientById(md.UserId).ReqSyncUpdate(context.Background(),
			&syncService.SyncUpdate{
				UserId:          md.UserId,
				IgnoreAuthKeyId: md.AuthKeyId,
				Updates: mtproto.NewTLUpdates(&mtproto.Updates{
					Updates: []*mtproto.Update{compat.NewUserBlock(inputPeer.GetPeerId(), false, md.Layer)},
					Users:   []*mtproto.User{},
					Chats:   []*mtproto.Chat{},
					Date:    int32(time.Now().Unix()),
					Seq:     0,
				}).To_Updates(),
				Push:     false,
				PeerType: constants.PeerTypeUser.ToInt32(),
			})
		if err != nil {
			log.Warnf("ContactsUnblock SyncUpdates error:%s", err.Error())
		}
	}()
	_ = ok
	return mtproto.ToMTBool(true), nil
}
