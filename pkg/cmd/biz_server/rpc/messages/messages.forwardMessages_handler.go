/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.forwardMessages_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/account/contact"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"time"
)

//  messages.forwardMessages#d9fee60e flags:# silent:flags.5?true background:flags.6?true with_my_score:flags.8?true from_peer:InputPeer id:Vector<int> random_id:Vector<long> to_peer:InputPeer schedule_date:flags.10?int = Updates;
//  messages.forwardMessages#708e0195 flags:# silent:flags.5?true background:flags.6?true with_my_score:flags.8?true from_peer:InputPeer id:Vector<int> random_id:Vector<long> to_peer:InputPeer = Updates;
//
func (s *MessagesServiceImpl) MessagesForwardMessages(ctx context.Context, request *mtproto.TLMessagesForwardMessages) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesForwardMessages %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	fromPeer := input.MakeInputPeer(request.FromPeer)
	if fromPeer.GetPeerType() == constants.PeerTypeSelf || fromPeer.GetPeerType() == constants.PeerTypeEmpty {
		fromPeer = input.MakeInputPeer(mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
			UserId: constants.PeerTypeFromUserIDType(md.UserId).ToInt32(),
			//AccessHash: request.FromPeer.AccessHash,
		}).To_InputPeer())
	}

	toPeer := input.MakeInputPeer(request.ToPeer)
	if toPeer.GetPeerType() == constants.PeerTypeSelf {
		toPeer = input.MakeInputPeer(mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
			UserId:     constants.PeerTypeFromUserIDType(md.UserId).ToInt32(),
			AccessHash: request.ToPeer.AccessHash,
		}).To_InputPeer())
	} else if toPeer.GetPeerType() == constants.PeerTypeChat {
		toPeer = input.MakeInputPeer(mtproto.NewTLInputPeerChannel(&mtproto.InputPeer{
			ChannelId:  request.ToPeer.ChatId,
			AccessHash: request.ToPeer.AccessHash,
		}).To_InputPeer())
	}

	if toPeer.GetPeerType() == constants.PeerTypeUser && toPeer.GetPeerId() != md.UserId {
		blocked, err := s.accountContactCore.CheckBlock(md.UserId, toPeer.GetPeerId(), int32(time.Now().Unix()))
		if err != nil {
			log.Errorf("ForwardMessages userId:%d request.Peer:%+v error:%s", md.UserId, toPeer.ToPeer(), err.Error())
			return nil, err
		}

		if blocked > contact.BlockedNone {
			err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_USER_IS_BLOCKED)
			return nil, err
		}
	} else if toPeer.GetPeerType() == constants.PeerTypeChat {
		toPeer = input.MakeInputPeer(mtproto.NewTLInputPeerChannel(&mtproto.InputPeer{
			ChannelId:  request.ToPeer.ChatId,
			AccessHash: request.ToPeer.AccessHash,
		}).To_InputPeer())
	}

	updates, err := s.accountMessageCore.ForwardMessages(md.UserId, md.AuthKeyId, fromPeer, toPeer, request.Id, request.RandomId, request.Silent)
	if err != nil {
		log.Errorf("MessagesForwardMessages %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	return updates, nil
}
