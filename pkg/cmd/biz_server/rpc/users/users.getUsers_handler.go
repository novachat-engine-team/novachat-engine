/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : users.getUsers_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/message"
	"novachat_engine/service/constants"
)

//  users.getUsers#d91a548 id:Vector<InputUser> = Vector<User>;
//
func (s *UsersServiceImpl) UsersGetUsers(ctx context.Context, request *mtproto.TLUsersGetUsers) (*mtproto.Vector_User, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("UsersGetUsers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	userIdList := make([]int64, 0, len(request.Id))
	for _, v := range request.Id {
		if v.ClassName == mtproto.ClassInputUserEmpty {
			userIdList = message.PickUserList(userIdList, md.UserId)
		} else if v.ClassName == mtproto.ClassInputUser {
			userIdList = message.PickUserList(userIdList, constants.PeerTypeFromUserIDType32(v.UserId).ToInt())
		}
	}
	if len(userIdList) == 0 {
		return &mtproto.Vector_User{}, nil
	}

	userCacheList, err := s.accountUserCore.GetUserList(md.UserId, userIdList, md.Layer)
	if err != nil {
		log.Warnf("UsersGetFullUser - request: %v GetByUserId error:%s", request, err.Error())
		return nil, err
	}

	ret := &mtproto.Vector_User{
		Datas: userCacheList,
	}
	log.Infof("UsersGetUsers - request: %v len()=%d ok!!!!!", request, len(ret.Datas))
	return ret, nil
}
