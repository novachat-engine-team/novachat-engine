/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.editMessage_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/util"
	"novachat_engine/service/account/message"
	"novachat_engine/service/app/help"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"time"
)

//  messages.editMessage#48f71778 flags:# no_webpage:flags.1?true peer:InputPeer id:int message:flags.11?string media:flags.14?InputMedia reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.15?int = Updates;
//  messages.editMessage#ce91e4ca flags:# no_webpage:flags.1?true peer:InputPeer id:int message:flags.11?string reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Updates;
//
func (s *MessagesServiceImpl) MessagesEditMessage(ctx context.Context, request *mtproto.TLMessagesEditMessage) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesEditMessage %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	inputPeer := input.MakeInputPeer(request.Peer)
	if inputPeer.GetPeerType() == constants.PeerTypeUser ||
		inputPeer.GetPeerType() == constants.PeerTypeSelf ||
		inputPeer.GetPeerType() == constants.PeerTypeChat ||
		inputPeer.GetPeerType() == constants.PeerTypeChannel {
		l, err := s.accountMessageCore.GetMessages(md.UserId, []int32{request.Id}, false, md.Layer)
		if err != nil {
			log.Infof("MessagesEditMessage - request: %v error:%s", request, err.Error())
			return nil, err
		}
		if len(l) == 0 {
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MESSAGE_ID_INVALID)
		}

		if !l[0].Out {
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MESSAGE_AUTHOR_REQUIRED)
		}

		if help.GetConfig().EditTimeLimit > 0 && int32(time.Now().Unix())-l[0].Date < help.GetConfig().EditTimeLimit {
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MESSAGE_EDIT_TIME_EXPIRED)
		}

		opts := make([]message.Option, 0, 1)
		if util.CheckBit(request.Flags, 11) {
			if len(request.Message) == 0 {
				return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_MESSAGE_EMPTY)
			}
			opts = append(opts, message.WithEditMessage(request.Message))
		}

		if util.CheckBit(request.Flags, 2) {
			opts = append(opts, message.WithEditReplyMarkup(request.ReplyMarkup))
		}

		if util.CheckBit(request.Flags, 3) {
			opts = append(opts, message.WithEditEntities(request.Entities))
		}

		if util.CheckBit(request.Flags, 15) {
			opts = append(opts, message.WithEditScheduleDate(request.ScheduleDate))
		}

		updates, err := s.accountMessageCore.EditMessage(md.AuthKeyId, md.UserId, inputPeer.GetPeerId(), inputPeer.GetPeerType().ToInt32(), request.Id, request.Media, md.Layer, opts...)
		if err != nil {
			log.Infof("MessagesEditMessage - request: %v EditMessage error:%s", request, err.Error())
			return nil, err
		}
		return updates, nil
	} else {
		panic(fmt.Sprintf("MessagesGetMessageEditData %+v", inputPeer.GetPeerType()))
		return nil, fmt.Errorf("%s", "Not impl MessagesGetMessageEditData")
	}
}
