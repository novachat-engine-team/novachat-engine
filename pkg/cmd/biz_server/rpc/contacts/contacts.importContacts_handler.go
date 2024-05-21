/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.importContacts_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/util"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/account/phone"
	"time"
)

//  contacts.importContacts#2c800be5 contacts:Vector<InputContact> = contacts.ImportedContacts;
//
func (s *ContactsServiceImpl) ContactsImportContacts(ctx context.Context, request *mtproto.TLContactsImportContacts) (*mtproto.Contacts_ImportedContacts, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ContactsImportContacts %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//TODO:Coder
	// ContactsImportContacts
	phoneList := make([]string, 0, len(request.Contacts))
	for _, r := range request.Contacts {
		if len(r.GetFirstName()) == 0 && len(r.GetLastName()) == 0 {
			continue
		}

		number, err1 := phone.Parse(r.Phone)
		if err1 != nil {
			continue
		}

		phoneNumber := number.NormalizeDigitsOnly()
		phoneList = append(phoneList, phoneNumber)
	}
	if len(phoneList) == 0 {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_NUMBER_INVALID)
	}

	userList, err := s.accountUserCore.PhoneUserList(phoneList)
	if err != nil {
		log.Errorf("ContactsImportContacts %v, request: %v PhoneUserList error:%s", md, request, err.Error())
		return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
	}

	imported := make([]*mtproto.ImportedContact, 0, len(phoneList))
	if len(userList) == 0 {
		log.Warnf("ContactsImportContacts request: %v PhoneUserList not found user", request)
		return mtproto.NewTLContactsImportedContacts(&mtproto.Contacts_ImportedContacts{
			Imported:       imported,
			PopularInvites: nil,
			RetryContacts:  []int64{},
			Users:          nil,
		}).To_Contacts_ImportedContacts(), nil

	} else {

		userIdList := make([]int64, 0, len(userList))
		for _, v := range userList {
			userIdList = append(userIdList, v.Id)
		}
		contact := request.Contacts[util.Index(request.Contacts, func(idx int) bool {
			number, _ := phone.Parse(request.Contacts[idx].Phone)
			return number.NormalizeDigitsOnly() == phoneList[0]
		})]

		mutualType, err := s.accountContactCore.AddContact(md.UserId,
			"",
			userIdList[0],
			contact.FirstName,
			contact.LastName,
			int32(time.Now().Unix()))
		if err != nil {
			log.Errorf("ContactsImportContacts %v, request: %v GetContactsList error:%s", md, request, err.Error())
			return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
		}
		_ = mutualType
		imported = append(imported, mtproto.NewTLImportedContact(&mtproto.ImportedContact{
			UserId:   constants.PeerTypeFromUserIDType(userIdList[0]).ToInt32(),
			ClientId: contact.ClientId,
		}).To_ImportedContact())

		users, _ := s.accountUserCore.GetUserList(md.UserId, userIdList, md.Layer)
		retryContacts := make([]int64, 0, len(phoneList))
		ret := mtproto.NewTLContactsImportedContacts(&mtproto.Contacts_ImportedContacts{
			Imported:       imported,
			PopularInvites: nil,
			RetryContacts:  retryContacts,
			Users:          users,
		})
		log.Debugf("ContactsImportContacts %v, request: %v reply ok!!!!!!!", md, request)
		return ret.To_Contacts_ImportedContacts(), nil
	}
}
