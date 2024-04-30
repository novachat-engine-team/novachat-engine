/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.getStatuses_handler.go
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
)

//  contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
//
func (s *ContactsServiceImpl) ContactsGetStatuses(ctx context.Context, request *mtproto.TLContactsGetStatuses) (*mtproto.Vector_ContactStatus, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ContactsGetStatuses %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	contactsList, err := s.accountContactCore.GetContacts(md.UserId)
	if err != nil {
		log.Errorf("ContactsGetStatuses %v, request: %v GetContactsList error:%s", md, request, err.Error())
		return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
	}

	statusList := &mtproto.Vector_ContactStatus{
		Datas: make([]*mtproto.ContactStatus, 0, len(contactsList)),
	}

	userIdList := make([]int64, 0, len(contactsList))
	for k := range contactsList {
		userIdList = append(userIdList, contactsList[k].UserId)
	}

	//TODO:Coder
	// status

	////userList, _ := s.Users.GetByUserIdList(0, userIdList)
	//
	//statusMap, err := s.Status.GetStatusList(md.UserId, userIdList, true)
	//if err != nil {
	//	log.Errorf("ContactsGetStatuses %v, request: %v GetStatusList error:%s", md, request, err.Error())
	//	return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
	//}
	//
	//for uid, m := range statusMap {
	//	statusList.Datas = append(statusList.Datas, mtproto.NewTLContactStatus(
	//		&mtproto.ContactStatus{
	//			UserId: uid,
	//			Status: m,
	//		},
	//	).To_ContactStatus())
	//}

	log.Debugf("ContactsGetStatuses %v, request: %v reply ok!!!!!!!!!!!!", md, request)
	return statusList, nil
}
