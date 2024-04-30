/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.editChatAdmin_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/banned_right"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"time"
)

//  messages.editChatAdmin#a9e69f2e chat_id:int user_id:InputUser is_admin:Bool = Bool;
//
func (s *MessagesServiceImpl) MessagesEditChatAdmin(ctx context.Context, request *mtproto.TLMessagesEditChatAdmin) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesEditChatAdmin %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	chatId := constants.PeerTypeFromChatIDType32(request.ChatId).ToInt()
	rights := int32(0)
	if mtproto.ToBool(request.IsAdmin) {
		rights = banned_right.ChatAdminBannedRightToRights(banned_right.FullAdminBannedRight())
	}
	_, err := chatService.GetChatClientByKeyId(chatId).ReqEditAdmin(ctx, &chatService.EditAdmin{
		ChatId:    chatId,
		PeerId:    input.MakeInputUser(request.UserId).GetId(),
		IsAdmin:   mtproto.ToBool(request.IsAdmin),
		UserId:    md.UserId,
		AuthKeyId: md.AuthKeyId,
		Date:      int32(time.Now().Unix()),
		Right:     rights,
	})
	if err != nil {
		log.Errorf("MessagesEditChatAdmin %v, request: %v EditChatAdmin error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	return mtproto.ToMTBool(true), nil
}
