/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.getBlocked_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
)

//  contacts.getBlocked#f57c350f offset:int limit:int = contacts.Blocked;
//
func (s *ContactsServiceImpl) ContactsGetBlocked(ctx context.Context, request *mtproto.TLContactsGetBlocked) (*mtproto.Contacts_Blocked, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsGetBlocked %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	contactsList, err := s.accountContactCore.GetBlocked(md.UserId)
	if err != nil {
		log.Errorf("ContactsGetBlocked %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	blockedPeer := make([]*mtproto.PeerBlocked, 0, len(contactsList))
	contactsBlockedList := make([]*mtproto.ContactBlocked, 0, len(contactsList))

	for _, v := range contactsList {
		blockedPeer = append(blockedPeer, mtproto.NewTLPeerBlocked(&mtproto.PeerBlocked{
			PeerId: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(v.PeerId).ToInt32()}).To_Peer(),
			Date:   v.Date,
		}).To_PeerBlocked())

		contactsBlockedList = append(contactsBlockedList, mtproto.NewTLContactBlocked(&mtproto.ContactBlocked{
			UserId: constants.PeerTypeFromUserIDType(v.PeerId).ToInt32(),
			Date:   v.Date,
		}).To_ContactBlocked())
	}

	return mtproto.NewTLContactsBlockedSlice(&mtproto.Contacts_Blocked{
		BlockedADE1591119: blockedPeer,
		Chats:             []*mtproto.Chat{},
		Users:             []*mtproto.User{},
		Count:             int32(len(contactsList)),
		Blocked1C138D1571: contactsBlockedList,
	}).To_Contacts_Blocked(), nil
}
