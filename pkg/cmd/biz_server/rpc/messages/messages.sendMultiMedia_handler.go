/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.sendMultiMedia_handler.go
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
	"novachat_engine/service/compat"
	data_users "novachat_engine/service/data/users"
)

//  messages.sendMultiMedia#cc0110cb flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int multi_media:Vector<InputSingleMedia> schedule_date:flags.10?int = Updates;
//
func (s *MessagesServiceImpl) MessagesSendMultiMedia(ctx context.Context, request *mtproto.TLMessagesSendMultiMedia) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesSendMultiMedia %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var userInfo *data_users.Users
	userInfo, err := s.accountUsersCore.UserData(md.UserId)
	if err != nil {
		log.Errorf("MessagesSendMultiMedia - request: %v GetByUserId error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}

	if userInfo.GetDeleted() {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_UNREGISTERED)
		log.Errorf("MessagesSendMultiMedia - request: %v error:%s", request, err.Error())
		return nil, err
	}

	updates, err := s.accountMessageCore.SendMultiMessage(md.UserId, md.AuthKeyId, request, md.Layer)
	if err != nil {
		log.Errorf("MessagesSendMultiMedia - request: %v error:%s", request, err.Error())
		return nil, err
	}

	updates = compat.UpdatesCompat(updates, md.Layer)

	log.Debugf("MessagesSendMultiMedia %v, request: %v reply ok!!!!!!!!!", md, request)
	return updates, nil
}
