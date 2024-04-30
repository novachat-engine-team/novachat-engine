/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.sendMedia_handler.go
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
	"novachat_engine/service/constants"
	data_users "novachat_engine/service/data/users"
)

//  messages.sendMedia#3491eba9 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int media:InputMedia message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int = Updates;
//  messages.sendMedia#c8f16791 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int media:InputMedia random_id:long reply_markup:flags.2?ReplyMarkup = Updates;
//
func (s *MessagesServiceImpl) MessagesSendMedia(ctx context.Context, request *mtproto.TLMessagesSendMedia) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesSendMedia %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.Message) >= constants.MessageLongLimit {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MESSAGE_TOO_LONG)
		log.Errorf("MessagesSendMedia %v, request: %v AuthGetAuthKey error:%s", md, request, err.Error())
		return nil, err
	}

	var userInfo *data_users.Users
	userInfo, err := s.accountUsersCore.UserData(md.UserId)
	if err != nil {
		log.Errorf("MessagesSendMedia - request: %v GetByUserId error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}

	if userInfo.GetDeleted() {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_UNREGISTERED)
		log.Errorf("MessagesSendMedia - request: %v error:%s", request, err.Error())
		return nil, err
	}

	//
	//msg, ok = filter.FilterDirtyWord(request.Message, bReplace, s.conf.BizLogic.FilterCloud.FilterCloud && s.conf.BizLogic.FilterCloud.Message)
	//_ = msg
	//if ok == true && bReplace == false {
	//	err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_FORBIDDEN_DIRTY_WORD)
	//	log.Errorf("MessagesSendMedia - request: %v FilterDirtyWord error:%s", request, err.Error())
	//	return nil, err
	//}

	updates, err := s.accountMessageCore.SendMessageMedia(md.UserId, md.AuthKeyId, request, md.Layer)
	if err != nil {
		log.Errorf("MessagesSendMedia - request: %v error:%s", request, err.Error())
		return nil, err
	}

	updates = compat.UpdatesCompat(updates, md.Layer)
	return updates, nil
}
