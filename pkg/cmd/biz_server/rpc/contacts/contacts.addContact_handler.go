/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.addContact_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	usersUtil "novachat_engine/service/account/users"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/data/contact"
	data_users "novachat_engine/service/data/users"
	"time"
)

//  contacts.addContact#e8f463d0 flags:# add_phone_privacy_exception:flags.0?true id:InputUser first_name:string last_name:string phone:string = Updates;
//
func (s *ContactsServiceImpl) ContactsAddContact(ctx context.Context, request *mtproto.TLContactsAddContact) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ContactsAddContact %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.LastName) == 0 && len(request.FirstName) == 0 {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MSG_ID_INVALID)
		log.Errorf("ContactsAddContact %v, request: %v error:%s", md, request, err.Error())
		return nil, err
	}

	//+ TL_inputUserEmpty
	//+ TL_inputUserSelf
	//+ TL_inputUser
	//+ TL_inputUserFromMessage

	updates := mtproto.NewTLUpdates(nil)
	switch request.Id.ClassName {
	case mtproto.ClassInputUserEmpty:
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CONTACT_ID_INVALID)
		log.Errorf("ContactsAddContact %v, request: %v error:%s", md, request, err.Error())
		return nil, err
	case mtproto.ClassInputUserSelf, mtproto.ClassInputUser:

		var userInfo *data_users.Users
		userInfo, err := s.accountUserCore.UserData(constants.PeerTypeFromUserIDType32(request.Id.UserId).ToInt())
		if err != nil {
			log.Errorf("ContactsAddContact %v, request: %v GetByUserId error:%s", md, request, err.Error())
			return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
		}

		user := usersUtil.UserCore2Users(userInfo)
		if user.Id == 0 || user.Deleted || user.Restricted {
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_FORBIDDEN_USER_RESTRICTED)
		}

		if mtproto.ClassInputUserSelf == request.Id.ClassName {
			user = usersUtil.UserCoreSelfUsers(user)
		} else {
			contact, err := s.accountContactCore.GetContactById(md.UserId, constants.PeerTypeFromUserIDType32(request.Id.UserId).ToInt())
			if err != nil {
				log.Errorf("ContactsAddContact %v, request: %v GetContactsList error:%s", md, request, err.Error())
				return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
			}
			modify := false
			state := contact.GetContact()
			if contact == nil || contact.GetContact() <= data_contact.MutualTypeDefault {
				state, err = s.accountContactCore.AddContact(md.UserId, request.Phone, constants.PeerTypeFromUserIDType32(request.Id.UserId).ToInt(), request.LastName, request.LastName, int32(time.Now().Unix()))
				if err != nil {
					log.Errorf("ContactsAddContact %v, request: %v AddContact error:%s", md, request, err.Error())
					return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
				}
				modify = true
			} else if contact.GetContact() > data_contact.MutualTypeMyContact && !(
				contact.LastName == request.LastName &&
					contact.FirstName == request.FirstName &&
					contact.Phone == request.Phone) {
				err = s.accountContactCore.ModifyContact(md.UserId, request.Phone, constants.PeerTypeFromUserIDType32(request.Id.UserId).ToInt(), request.LastName, request.LastName, int32(time.Now().Unix()))
				if err != nil {
					log.Errorf("ContactsAddContact %v, request: %v AddContact error:%s", md, request, err.Error())
					return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
				}
				modify = true
				state = contact.GetContact()
			}
			if modify {
				//TODO:Coder
				// sync to devices
				user = usersUtil.UserCoreContactUser(user, true, state >= data_contact.MutualTypeMyContact)
				updates.SetUpdates([]*mtproto.Update{
					mtproto.NewTLUpdateUserName(&mtproto.Update{
						UserId:    request.Id.UserId,
						FirstName: request.FirstName,
						LastName:  request.LastName,
						Username:  user.Username,
					}).To_Update(),
					mtproto.NewTLUpdateUserPhone(&mtproto.Update{
						UserId: request.Id.UserId,
						Phone:  request.Phone,
					}).To_Update()})
				updates.SetUsers([]*mtproto.User{user})
			}
		}
	default:
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CONTACT_ID_INVALID)
		log.Errorf("ContactsAddContact %v, request: %v error:%s", md, request, err.Error())
		return nil, err

	}

	log.Debugf("ContactsAddContact %v, request: %v reply ok !!!!!!!!!!!!", md, request)
	return updates.To_Updates(), nil
}
