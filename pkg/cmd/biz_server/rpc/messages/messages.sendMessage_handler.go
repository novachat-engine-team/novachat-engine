/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.sendMessage_handler.go
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

//  messages.sendMessage#520c3870 flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int = Updates;
//  messages.sendMessage#fa88427a flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Updates;
//
func (s *MessagesServiceImpl) MessagesSendMessage(ctx context.Context, request *mtproto.TLMessagesSendMessage) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesSendMessage %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var err error

	if len(request.Message) == 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MESSAGE_EMPTY)
		log.Errorf("MessagesSendMessage - request: %v AuthGetAuthKey error:%s", request, err.Error())
		return nil, err
	}

	if len(request.Message) >= constants.MessageLongLimit {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MESSAGE_TOO_LONG)
		log.Errorf("MessagesSendMessage - request: %v AuthGetAuthKey error:%s", request, err.Error())
		return nil, err
	}

	var userInfo *data_users.Users
	userInfo, err = s.accountUsersCore.UserData(md.UserId)
	if err != nil {
		log.Errorf("MessagesSendMessage - request: %v GetByUserId error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}

	if userInfo.GetDeleted() {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_UNREGISTERED)
		log.Errorf("MessagesSendMessage - request: %v error:%s", request, err.Error())
		return nil, err
	}

	//msg, ok := filter.FilterDirtyWord(request.Message, false, s.conf.BizLogic.FilterCloud.FilterCloud && s.conf.BizLogic.FilterCloud.Message)
	//_ = msg
	//if ok == true {
	//	err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_FORBIDDEN_DIRTY_WORD)
	//	log.Errorf("MessagesSendMessage - request: %v FilterDirtyWord error:%s", request, err.Error())
	//	return nil, err
	//}
	updates, err := s.accountMessageCore.SendMessage(md.UserId, md.AuthKeyId, request, md.Layer)
	if err != nil {
		log.Errorf("MessagesSendMessage - request: %v error:%s", request, err.Error())
		return nil, err
	}

	updates = compat.UpdatesCompat(updates, md.Layer)

	log.Debugf("MessagesSendMessage %v, request: %v reply ok!!!!!!!!!", md, request)
	return updates, nil
}
