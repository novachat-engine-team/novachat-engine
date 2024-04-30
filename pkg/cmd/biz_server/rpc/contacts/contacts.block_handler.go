/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.block_handler.go
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
	errors1 "novachat_engine/service/common/errors"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	data_users "novachat_engine/service/data/users"
	"novachat_engine/service/input"
	"time"
)

//  contacts.block#68cc1411 id:InputPeer = Bool;
//  contacts.block#332b49fc id:InputUser = Bool;
//
func (s *ContactsServiceImpl) ContactsBlock(ctx context.Context, request *mtproto.TLContactsBlock) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsBlock %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if request.Id332B49FC71 != nil {
		request.Id68CC1411119 = mtproto.NewTLInputPeerUser(&mtproto.InputPeer{UserId: request.Id332B49FC71.UserId}).To_InputPeer()
	}

	inputPeer := input.MakeInputPeer(request.Id68CC1411119)
	if inputPeer.GetPeerType() != constants.PeerTypeUser && inputPeer.GetPeerType() != constants.PeerTypeSelf {
		log.Errorf("ContactsBlock - request: %v inputPeer type error:%v", request, inputPeer.GetPeerType())
		return mtproto.ToMTBool(false), nil
	}

	var userInfo *data_users.Users
	userInfo, err := s.accountUserCore.UserData(inputPeer.GetPeerId())
	if err != nil {
		log.Infof("ContactsBlock - request: %v User error:%s", request, err.Error())
		return nil, err
	}

	if userInfo.Id == 0 || userInfo.GetRestricted() || userInfo.GetDeleted() {
		return nil, errors1.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_FORBIDDEN_USER_RESTRICTED)
	}

	err = s.accountContactCore.Block(md.UserId, inputPeer.GetPeerId(), int32(time.Now().Unix()))
	if err != nil {
		log.Errorf("ContactsBlock - request: %v error:%s", request, err.Error())
		return nil, err
	}

	go func() {
		_, err = syncService.GetSyncClientById(md.UserId).ReqSyncUpdate(context.Background(), &syncService.SyncUpdate{
			UserId:          md.UserId,
			IgnoreAuthKeyId: md.AuthKeyId,
			Updates: mtproto.NewTLUpdates(&mtproto.Updates{
				Updates: []*mtproto.Update{compat.NewUserBlock(inputPeer.GetPeerId(), true, md.Layer)},
				Users:   []*mtproto.User{},
				Chats:   []*mtproto.Chat{},
				Date:    int32(time.Now().Unix()),
				Seq:     0,
			}).To_Updates(),
			Push:     false,
			PeerType: constants.PeerTypeUser.ToInt32(),
		})
		if err != nil {
			log.Warnf("ContactsBlock ReqSyncUpdate error:%s", err.Error())
		}
	}()

	return mtproto.ToMTBool(true), nil
}
