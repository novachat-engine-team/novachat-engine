/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.deleteContact_handler.go
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
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/data/users"
	"time"
)

//  contacts.deleteContact#8e953744 id:InputUser = contacts.Link;
//
func (s *ContactsServiceImpl) ContactsDeleteContact(ctx context.Context, request *mtproto.TLContactsDeleteContact) (*mtproto.Contacts_Link, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsDeleteContact %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	idList := []int64{constants.PeerTypeFromUserIDType32(request.Id.UserId).ToInt()}
	_, err := s.accountContactCore.ContactsRemove(md.UserId, idList)
	if err != nil {
		log.Errorf("ContactsDeleteContact - request: %v error:%s", request, err.Error())
		return nil, err
	}

	var userInfoList []*data_users.Users
	userCacheList, err := s.accountCore.UserList(idList)
	if err != nil {
		userInfoList, err = s.userCore.UserIdList(idList)
		if err != nil {
			log.Errorf("ContactsDeleteContact - request: %v error:%s", request, err.Error())
			return nil, err
		}
	} else {
		var userInfo *data_users.Users
		for _, v := range userCacheList {
			userInfo, _ = account.CacheInfo2User(v)
			if userInfo == nil {
				continue
			}

			userInfoList = append(userInfoList, userInfo)
		}
	}

	updates := mtproto.NewTLUpdates(nil)
	var user *mtproto.User
	userList := make([]*mtproto.User, 0, len(userInfoList))
	for _, v := range userInfoList {
		user = users.UserCore2Users(v)
		if v.Id == md.UserId {
			users.UserCoreSelfUsers(user)
		} else {
			user = users.UserCoreContactUser(user, false, false)
		}
		userList = append(userList, user)
	}
	updates.SetUsers(userList)
	updates.SetDate(int32(time.Now().Unix()))

	go func() {
		_, err = syncService.GetSyncClientById(md.UserId).ReqSyncUpdate(context.Background(), &syncService.SyncUpdate{
			UserId:          md.UserId,
			IgnoreAuthKeyId: md.AuthKeyId,
			Updates:         updates.To_Updates(),
			Push:            false,
			PeerType:        constants.PeerTypeUser.ToInt32(),
		})
		if err != nil {
			log.Warnf("ContactsDeleteContact ReqSyncUpdate error:%s", err.Error())
		}
	}()

	link := mtproto.NewTLContactsLink(nil).To_Contacts_Link()
	link.User = user
	link.MyLink = mtproto.NewTLContactLinkHasPhone(nil).To_ContactLink()
	link.ForeignLink = mtproto.NewTLContactLinkUnknown(nil).To_ContactLink()

	return link, nil
}
