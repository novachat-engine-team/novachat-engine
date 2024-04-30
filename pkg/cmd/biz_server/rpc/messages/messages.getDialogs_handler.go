/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getDialogs_handler.go
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

//req.offset_peer.access_hash = ((long) dialogsLoadOffset[UserConfig.i_dialogsLoadOffsetAccess_1]) | ((long) dialogsLoadOffset[UserConfig.i_dialogsLoadOffsetAccess_1] << 32);

//  messages.getDialogs#a0ee3b73 flags:# exclude_pinned:flags.0?true folder_id:flags.1?int offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:int = messages.Dialogs;
//  messages.getDialogs#191ba9c5 flags:# exclude_pinned:flags.0?true offset_date:int offset_id:int offset_peer:InputPeer limit:int = messages.Dialogs;
//
func (s *MessagesServiceImpl) MessagesGetDialogs(ctx context.Context, request *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetDialogs %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if request.OffsetPeer != nil {
		log.Debugf("MessagesGetDialogs %v, request: %v access_hash:%v", md, request, request.OffsetPeer.AccessHash)
	}

	dialogs, err := s.accountMessageCore.GetDialog(md.UserId, request.ExcludePinned, request.FolderId, request.OffsetId, request.OffsetDate, request.OffsetPeer, request.Limit, request.Hash, md.Layer)
	if err != nil {
		log.Errorf("MessagesGetDialogs%v, error:%s", md, err.Error())
		return nil, errors.NewRpcErrorWithCodeString(mtproto.RpcErrorCode_INTERNAL.ToInt32(), err.Error())
	}

	for idx := range dialogs.Dialogs {
		dialogs.Dialogs[idx] = compat.CompatDialog(dialogs.Dialogs[idx], md.Layer)
	}

	for idx := range dialogs.Messages {
		dialogs.Messages[idx] = compat.MessageCompat(dialogs.Messages[idx], md.Layer)
	}

	for idx := range dialogs.Users {
		dialogs.Users[idx] = compat.CompatUser(dialogs.Users[idx], md.Layer)
	}

	for idx := range dialogs.Chats {
		dialogs.Chats[idx] = compat.CompatChat(dialogs.Chats[idx], md.Layer)
	}

	log.Debugf("MessagesGetDialogs %v, request: %v reply:%v", md, request, dialogs)
	log.Infof("MessagesGetDialogs %v, request: %v reply count:%d ok!!!!!!!!!", md, request, dialogs.Count)
	return dialogs, nil
}
