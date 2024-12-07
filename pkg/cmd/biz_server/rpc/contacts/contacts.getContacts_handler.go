/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.getContacts_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	usersUtil "novachat_engine/service/account/users"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/common/hash"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/account/account"
	data_contact "novachat_engine/service/data/contact"
	data_users "novachat_engine/service/data/users"
)

/*
private long getContactsHash(ArrayList<TLRPC.TL_contact> contacts) {
        long acc = 0;
        contacts = new ArrayList<>(contacts);
        Collections.sort(contacts, (tl_contact, tl_contact2) -> {
            if (tl_contact.user_id > tl_contact2.user_id) {
                return 1;
            } else if (tl_contact.user_id < tl_contact2.user_id) {
                return -1;
            }
            return 0;
        });
        int count = contacts.size();
        for (int a = -1; a < count; a++) {
            if (a == -1) {
                acc = MediaDataController.calcHash(acc, getUserConfig().contactsSavedCount);
            } else {
                TLRPC.TL_contact set = contacts.get(a);
                acc = MediaDataController.calcHash(acc, set.user_id);
            }
        }
        return acc;
    }
*/

//  contacts.getContacts#c023849f hash:int = contacts.Contacts;
//
func (s *ContactsServiceImpl) ContactsGetContacts(ctx context.Context, request *mtproto.TLContactsGetContacts) (*mtproto.Contacts_Contacts, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ContactsGetContacts %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	contacts, err := s.accountContactCore.GetContacts(md.UserId)
	if err != nil {
		log.Errorf("ContactsGetContacts - request: %v error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
	}

	contactsMap := make(map[int64]*mtproto.Contact, len(contacts))
	contactsCacheMap := make(map[int64]*data_contact.Contact, len(contacts))
	userInterList := make([]string, 0, len(contacts))
	userIdList := make([]int64, 0, len(contacts)+1)
	for _, v := range contacts {
		c := mtproto.NewTLContact(nil)
		c.SetUserId(constants.PeerTypeFromUserIDType(v.UserId).ToInt32())
		c.SetMutual(mtproto.ToMTBool(v.GetContact() >= data_contact.MutualTypeMutual))
		userIdList = append(userIdList, v.UserId)
		userInterList = append(userInterList, fmt.Sprintf("%d", v.UserId))
		contactsCacheMap[v.UserId] = v
		contactsMap[v.UserId] = c.To_Contact()
	}
	if len(contacts) == 0 {
		return mtproto.NewTLContactsContacts(nil).To_Contacts_Contacts(), nil
	}

	if hash.HashIdList(constants.ArrayInt64To32(userIdList), true, false) == request.HashC023849F71 {
		log.Debugf("ContactsGetContacts - request: %v reply ok!!!!!!!!!", request)
		return mtproto.NewTLContactsContactsNotModified(nil).To_Contacts_Contacts(), nil
	}

	userIdList = append(userIdList, md.UserId)
	ret := mtproto.NewTLContactsContacts(nil)
	ret.SetSavedCount(int32(len(contacts)))

	data, err := s.accountCore.UserList(userIdList)
	if err != nil {
		log.Errorf("ContactsGetContacts - request: %v UserByUserIdList error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
	}

	contactsList := make([]*mtproto.Contact, 0, len(contacts))
	userList := make([]*mtproto.User, 0, len(data))
	var user *mtproto.User
	var userInfo *data_users.Users
	var vv *data_contact.Contact
	for _, v := range data {
		if len(v) == 0 {
			continue
		}

		userInfo, err = account.CacheInfo2User(v)
		if err != nil {
			log.Warnf("ContactsGetContacts - request: %v CacheInfo2User error:%s", request, err.Error())
			continue
		}

		vv, _ = contactsCacheMap[userInfo.Id]
		if vv == nil || vv.Deleted {
			log.Warnf("ContactsGetContacts - request: %v contactsCacheMap error:%s", request, err.Error())
			continue
		}
		if userInfo.Id == md.UserId {
			user = usersUtil.UserCoreSelfUsers(usersUtil.UserCore2Users(userInfo))
		} else {
			user = usersUtil.UserCoreContactUser(usersUtil.UserCore2Users(userInfo), !vv.Deleted, !vv.Deleted && vv.GetContact() > data_contact.MutualTypeMyContact, vv)
		}

		userList = append(userList, user)
		cc, ok1 := contactsMap[userInfo.Id]
		if ok1 == true {
			contactsList = append(contactsList, cc)
		}
	}

	ret.SetUsers(userList)
	ret.SetContacts(contactsList)
	ret.SetSavedCount(int32(len(userList)))

	log.Debugf("ContactsGetContacts - request: %v reply ok!!!!!!!!!", request)
	return ret.To_Contacts_Contacts(), err
}
