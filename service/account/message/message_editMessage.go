/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package message

import (
	"context"
	"novachat_engine/mtproto"
	msgClient "novachat_engine/pkg/cmd/msg/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/rpc_context"
)

type Option func(message *msgClient.EditMessage)

func WithEditMessage(s string) Option {
	return func(message *msgClient.EditMessage) { message.Message = s }
}
func WithEditMedia(m *mtproto.MessageMedia) Option {
	return func(message *msgClient.EditMessage) { message.Media = m }
}
func WithEditReplyMarkup(s *mtproto.ReplyMarkup) Option {
	return func(message *msgClient.EditMessage) { message.ReplyMarkup = s }
}
func WithEditEntities(s []*mtproto.MessageEntity) Option {
	return func(message *msgClient.EditMessage) { message.Entities = s }
}
func WithEditScheduleDate(s int32) Option {
	return func(message *msgClient.EditMessage) { message.ScheduleData = s }
}

func (c *Core) EditMessage(authKeyId int64, userId int64, peerId int64, peerType int32, id int32, inputMedia *mtproto.InputMedia, layer int32, opts ...Option) (*mtproto.Updates, error) {
	if inputMedia == nil && len(opts) == 0 {
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MESSAGE_NOT_MODIFIED)
	}

	editMessage := &msgClient.EditMessage{
		AuthKeyId:  authKeyId,
		FromUserId: userId,
		PeerId:     peerId,
		PeerType:   peerType,
		MessageId:  id,
	}
	if inputMedia != nil {
		mediaList, err := c.makeInputMedia(authKeyId, layer, []*mtproto.InputMedia{inputMedia})
		if err != nil {
			log.Infof("EditMessage - EditMessage error:%s", err.Error())
			return nil, err
		}

		if len(mediaList) == 0 {
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MEDIA_INVALID)
		}
		WithEditMedia(mediaList[0])(editMessage)
	}

	for _, opt := range opts {
		opt(editMessage)
	}
	ctx, cancel := rpc_context.Context(
		context.Background(),
		rpc_context.WithTimeout(constants.RpcDefaultTimeout),
	)
	updates, err := msgClient.GetMsgClient().ReqEditMessage(ctx, editMessage)
	cancel()
	if err != nil {
		log.Errorf("EditMessage EditMessage error:%s", err.Error())
		return nil, err
	}

	//TODO:(Coderxw) sync
	return updates, nil
}
