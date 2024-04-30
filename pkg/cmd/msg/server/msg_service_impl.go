/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/9 18:36
 * @File : msg_service_impl.go
 */

package server

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	msgService "novachat_engine/pkg/cmd/msg/rpc_client"
	handler2 "novachat_engine/pkg/cmd/msg/server/handler"
	"novachat_engine/pkg/log"
)

type MsgServiceImpl struct {
	handler *handler2.MsgHandler
}

func NewMsgServiceImpl(h *handler2.MsgHandler) *MsgServiceImpl {
	impl := new(MsgServiceImpl)
	impl.handler = h

	return impl
}

func (m *MsgServiceImpl) ReqSendMessages(ctx context.Context, sendMessages *msgService.SendMessages) (*mtproto.Updates, error) {
	_ = ctx
	log.Debugf("ReqSendMessages sendMessages fromUserId:%d peerId:%d peerType:%d", sendMessages.FromUserId, sendMessages.PeerId, sendMessages.PeerType)
	return m.handler.SendMessageOutBoxes(sendMessages)
}

func (m *MsgServiceImpl) ReqEditMessage(ctx context.Context, edit *msgService.EditMessage) (*mtproto.Updates, error) {
	_ = ctx
	log.Debugf("ReqEditMessage editMessages fromUserId:%d peerId:%d peerType:%d", edit.FromUserId, edit.PeerId, edit.PeerType)
	return m.handler.SendEditMessageOutBoxes(edit)
}

func (m *MsgServiceImpl) ReqRevokeMessages(ctx context.Context, revokeMessage *msgService.RevokeMessages) (*types.Any, error) {
	_ = ctx
	log.Debugf("ReqRevokeMessages ReadHistory:%+v", revokeMessage)
	return m.handler.RevokeMessage(revokeMessage)
}

func (m *MsgServiceImpl) ReqReadHistory(ctx context.Context, readHistory *msgService.ReadHistory) (*types.Any, error) {
	_ = ctx
	log.Debugf("ReqReadHistory ReadHistory:%+v", readHistory)
	return m.handler.ReadHistory(readHistory)
}

func (m *MsgServiceImpl) ReqPinnedMessage(ctx context.Context, pinnedMessage *msgService.PinnedMessage) (*types.Any, error) {
	_ = ctx
	log.Debugf("ReqPinnedMessage pinnedMessage:%+v", pinnedMessage)
	return m.handler.PinnedMessage(pinnedMessage)
}
