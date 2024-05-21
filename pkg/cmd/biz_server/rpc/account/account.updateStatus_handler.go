/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.updateStatus_handler.go
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
	"novachat_engine/service/constants"
	"strings"
	"time"
)

//  account.updateStatus#6628562c offline:Bool = Bool;
//
func (s *AccountServiceImpl) AccountUpdateStatus(ctx context.Context, request *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountUpdateStatus %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	now := time.Now()
	var status *mtproto.UserStatus
	if mtproto.ToBool(request.Offline) == true {
		status = mtproto.NewTLUserStatusOffline(&mtproto.UserStatus{
			WasOnline: int32(now.Unix()),
		}).To_UserStatus()
	} else {
		status = mtproto.NewTLUserStatusOnline(&mtproto.UserStatus{
			Expires: int32(now.Unix() + 60),
		}).To_UserStatus()
	}

	contactList, err := s.accountContactCore.GetContacts(md.UserId)
	if err != nil {
		log.Warnf("AccountUpdateStatus - request: %v error:%s", request, err.Error())
		return mtproto.ToMTBool(false), nil
	}

	userIdList := make([]int64, 0, len(contactList))
	for _, v := range contactList {
		userIdList = append(userIdList, v.UserId)
	}
	var userList []*mtproto.User
	if len(userIdList) > 0 {
		userList, err = s.accountUsersCore.GetUserList(md.UserId, userIdList, md.Layer)
		if err != nil {
			log.Warnf("AccountUpdateStatus - request: %v GetUserList error:%s", request, err.Error())
			return mtproto.ToMTBool(false), nil
		}
	}

	pushUpdatesList := make([]*syncService.UpdateData, 0, len(contactList))
	updateUserStatus := mtproto.NewTLUpdateUserStatus(&mtproto.Update{
		UserId: constants.PeerTypeFromUserIDType(md.UserId).ToInt32(),
		Status: status,
	}).To_Update()
	var contactUser *mtproto.User
	for _, c := range contactList {
		if c.UserId == 0 {
			continue
		}
		contactUser = nil
		for _, vv := range userList {
			if c.UserId == constants.PeerTypeFromUserIDType32(vv.Id).ToInt() {
				contactUser = vv
				break
			}
		}

		if contactUser == nil || contactUser.Status == nil || strings.Compare(contactUser.Status.ClassName, mtproto.ClassUserStatusOnline) != 0 {
			continue
		}

		pushUpdatesList = append(pushUpdatesList, &syncService.UpdateData{
			UserId: c.UserId,
			Updates: mtproto.NewTLUpdates(&mtproto.Updates{
				Updates: []*mtproto.Update{updateUserStatus},
				Users:   []*mtproto.User{},
				Chats:   []*mtproto.Chat{},
				Date:    int32(time.Now().Unix()),
				Seq:     0,
			}).To_Updates(),
		})
	}

	if len(pushUpdatesList) > 0 {
		go func() {
			_, err = syncService.GetSyncClientById(md.UserId).ReqSyncUpdates(context.Background(),
				&syncService.SyncUpdates{
					UpdateDataList: pushUpdatesList,
					Push:           false,
					PeerType:       constants.PeerTypeUser.ToInt32(),
				})
			if err != nil {
				log.Infof("AccountUpdateStatus - PushUpdatesList error:%s", err.Error())
			}
		}()
	}
	return mtproto.ToMTBool(true), nil
}
