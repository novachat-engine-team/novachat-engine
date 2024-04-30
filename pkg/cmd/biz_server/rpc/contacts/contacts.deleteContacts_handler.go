/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.deleteContacts_handler.go
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

//  contacts.deleteContacts#96a0e00 id:Vector<InputUser> = Updates;
//  contacts.deleteContacts#59ab389e id:Vector<InputUser> = Bool;
//
func (s *ContactsServiceImpl) ContactsDeleteContacts(ctx context.Context, request *mtproto.TLContactsDeleteContacts) (*mtproto.Response_ContactsDeleteContacts, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsDeleteContacts %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl ContactsDeleteContacts logic
	idList := make([]int64, 0, len(request.Id))
	for _, id := range request.Id {
		if id.UserId == 0 {
			continue
		}
		idList = append(idList, constants.PeerTypeFromUserIDType32(id.UserId).ToInt())
	}

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	})
	if len(idList) > 0 {
		_, err := s.accountContactCore.ContactsRemove(md.UserId, idList)
		if err != nil {
			log.Errorf("ContactsDeleteContacts - request: %v error:%s", request, err.Error())
			return nil, err
		}

		var userInfoList []*data_users.Users
		userCacheList, err := s.accountCore.UserList(idList)
		if err != nil {
			userInfoList, err = s.userCore.UserIdList(idList)
			if err != nil {
				log.Errorf("ContactsDeleteContacts - request: %v error:%s", request, err.Error())
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

		go func() {
			_, err = syncService.GetSyncClientById(md.UserId).ReqSyncUpdate(context.Background(), &syncService.SyncUpdate{
				UserId:          md.UserId,
				IgnoreAuthKeyId: md.AuthKeyId,
				Updates:         updates.To_Updates(),
				Push:            false,
				PeerType:        constants.PeerTypeUser.ToInt32(),
			})
			if err != nil {
				log.Warnf("ContactsDeleteContacts ReqSyncUpdate error:%s", err.Error())
			}
		}()
	}

	ret := &mtproto.Response_ContactsDeleteContacts{
		ContactsDeleteContacts96A0E00:  updates.To_Updates(),
		ContactsDeleteContacts59Ab389E: mtproto.ToMTBool(true),
	}

	if md.Layer >= mtproto.CurrentLayer {
		ret.Cmd = mtproto.Cmd_ContactsDeleteContacts96a0e00.ToUInt32()
	} else {
		ret.Cmd = mtproto.Cmd_ContactsDeleteContacts.ToUInt32()
	}
	return ret, nil
}
