/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.acceptContact_handler.go
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
	"novachat_engine/service/account/users"
	"novachat_engine/service/constants"
	"novachat_engine/service/data/contact"
	"novachat_engine/service/data/users"
	"time"
)

//  contacts.acceptContact#f831a20f id:InputUser = Updates;
//
func (s *ContactsServiceImpl) ContactsAcceptContact(ctx context.Context, request *mtproto.TLContactsAcceptContact) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsAcceptContact %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	contact, err1 := s.accountContactCore.GetContactById(md.UserId, constants.PeerTypeFromUserIDType32(request.Id.UserId).ToInt())
	if err1 != nil {
		log.Infof("ContactsAcceptContact - request: %v error:%s", request, err1.Error())
		return nil, err1
	}

	updates := mtproto.NewTLUpdates(nil)
	if constants.PeerTypeFromUserIDType32(request.Id.UserId).ToInt() != md.UserId && (contact == nil || contact.GetContact() <= data_contact.MutualTypeMyContact) {
		contactState, err := s.accountContactCore.AddContact(md.UserId, "", constants.PeerTypeFromUserIDType32(request.Id.UserId).ToInt(), "", "", int32(time.Now().Unix()))
		if err != nil {
			log.Infof("ContactsAcceptContact - request: %v AcceptContact error:%s", request, err.Error())
			return nil, err
		}

		var userInfo *data_users.Users
		userInfo, err = s.accountUserCore.UserData(constants.PeerTypeFromUserIDType32(request.Id.UserId).ToInt())
		if err != nil {
			log.Infof("ContactsAcceptContact - request: %v User error:%s", request, err.Error())
			return nil, err
		}
		user := users.UserCore2Users(userInfo)
		user = users.UserCoreContactUser(user, true, contactState >= data_contact.MutualTypeMutual)
		updates.SetUsers([]*mtproto.User{user})

		go func() {
			_, err = syncService.GetSyncClientById(md.UserId).ReqSyncUpdate(context.Background(), &syncService.SyncUpdate{
				UserId:          md.UserId,
				IgnoreAuthKeyId: md.AuthKeyId,
				Updates:         updates.To_Updates(),
				Push:            false,
				PeerType:        constants.PeerTypeUser.ToInt32(),
			})
			if err != nil {
				log.Warnf("ContactsAcceptContact ReqSyncUpdate error:%s", err.Error())
			}
		}()
	}

	return updates.To_Updates(), nil
}
