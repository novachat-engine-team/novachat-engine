/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.getSaved_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  contacts.getSaved#82f1e39f = Vector<SavedContact>;
//
func (s *ContactsServiceImpl) ContactsGetSaved(ctx context.Context, request *mtproto.TLContactsGetSaved) (*mtproto.Vector_SavedContact, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsGetSaved %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	contactsList, err := s.accountContactCore.GetContacts(md.UserId)
	if err != nil {
		log.Warnf("ContactsGetSaved request: %v error:%s", request, err.Error())
	}
	resp := mtproto.Vector_SavedContact{
		Datas: make([]*mtproto.SavedContact, 0, len(contactsList)),
	}

	for _, v := range contactsList {
		resp.Datas = append(resp.Datas, mtproto.NewTLSavedPhoneContact(&mtproto.SavedContact{
			Phone:     v.Phone,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Date:      v.Date,
		}).To_SavedContact())
	}

	return &resp, nil
}
