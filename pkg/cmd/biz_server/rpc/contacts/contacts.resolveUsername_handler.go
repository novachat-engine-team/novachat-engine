/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.resolveUsername_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/account/users"
	usersUtil "novachat_engine/service/account/users"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/data/contact"
)

//  contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
//
func (s *ContactsServiceImpl) ContactsResolveUsername(ctx context.Context, request *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsResolveUsername %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	contactsResolvedPeer := mtproto.NewTLContactsResolvedPeer(&mtproto.Contacts_ResolvedPeer{
		Peer:  nil,
		Chats: nil,
		Users: nil,
	})
	var user *mtproto.User
	if len(request.Username) > 0 {
		userInfo, err := s.accountUserCore.UserDataUsername(request.Username)
		if err != nil {
			log.Errorf("ContactsResolveUsername - request: %v error:%s", request, err.Error())
			return contactsResolvedPeer.To_Contacts_ResolvedPeer(), nil
		}
		user = users.UserCore2Users(userInfo)
		if constants.PeerTypeFromUserIDType32(user.Id).ToInt() == md.UserId {
			user = users.UserCoreSelfUsers(user)
		} else {
			contact, err := s.accountContactCore.GetContactById(md.UserId, constants.PeerTypeFromUserIDType32(user.Id).ToInt())
			if err != nil {
				log.Errorf("ContactsResolveUsername - request: %v GetContactById error:%s", request, err.Error())
				return contactsResolvedPeer.To_Contacts_ResolvedPeer(), nil
			}

			if contact != nil {
				user = usersUtil.UserCoreContactUser(usersUtil.UserCore2Users(userInfo), !contact.Deleted, !contact.Deleted && contact.GetContact() > data_contact.MutualTypeMyContact, contact)
			} else {
				user = usersUtil.UserCoreContactUser(usersUtil.UserCore2Users(userInfo), false, false, nil)
			}
		}
	}

	if user != nil {
		if user.Id == 0 {
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USERNAME_INVALID)
			//			contactsResolvedPeer.SetPeer(mtproto.NewTLPeerUser(&mtproto.Peer{UserId: user.Id}).To_Peer())
		} else {
			contactsResolvedPeer.SetPeer(mtproto.NewTLPeerUser(&mtproto.Peer{UserId: user.Id}).To_Peer())
			contactsResolvedPeer.SetUsers([]*mtproto.User{user})
		}
	} else {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USERNAME_INVALID)
	}
	return contactsResolvedPeer.To_Contacts_ResolvedPeer(), nil
}
