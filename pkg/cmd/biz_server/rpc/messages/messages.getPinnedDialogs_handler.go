/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getPinnedDialogs_handler.go
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
)

//  messages.getPinnedDialogs#d6b94df2 folder_id:int = messages.PeerDialogs;
//  messages.getPinnedDialogs#e254d64e = messages.PeerDialogs;
//
func (s *MessagesServiceImpl) MessagesGetPinnedDialogs(ctx context.Context, request *mtproto.TLMessagesGetPinnedDialogs) (*mtproto.Messages_PeerDialogs, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetPinnedDialogs %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//var state *mtproto.Updates_State

	messagePeerDialogs, err := s.accountMessageCore.GetPinnedDialog(md.UserId, request.FolderId, md.Layer)
	if err != nil {
		log.Errorf("MessagesGetPinnedDialogs %v, request: %v GetDialog error:%s", md, request, err.Error())
		return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
	}

	for idx := range messagePeerDialogs.Dialogs {
		messagePeerDialogs.Dialogs[idx] = compat.CompatDialog(messagePeerDialogs.Dialogs[idx], md.Layer)
	}

	for idx := range messagePeerDialogs.Messages {
		messagePeerDialogs.Messages[idx] = compat.MessageCompat(messagePeerDialogs.Messages[idx], md.Layer)
	}

	for idx := range messagePeerDialogs.Users {
		messagePeerDialogs.Users[idx] = compat.CompatUser(messagePeerDialogs.Users[idx], md.Layer)
	}

	for idx := range messagePeerDialogs.Chats {
		messagePeerDialogs.Chats[idx] = compat.CompatChat(messagePeerDialogs.Chats[idx], md.Layer)
	}
	//messagePeerDialogs.State = state

	log.Infof("MessagesGetPinnedDialogs %v, request: %v users:len()=%d messages:len()=%d reply ok", md, request, len(messagePeerDialogs.Users), len(messagePeerDialogs.Messages))
	return messagePeerDialogs, nil
}
