/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.readMentions_handler.go
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
)

//  messages.readMentions#f0189d3 peer:InputPeer = messages.AffectedHistory;
//
func (s *MessagesServiceImpl) MessagesReadMentions(ctx context.Context, request *mtproto.TLMessagesReadMentions) (*mtproto.Messages_AffectedHistory, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesReadMentions %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl MessagesReadMentions logic
	//inputPeer := input.MakeInputPeer(request.Peer)
	//if inputPeer.GetPeerType() == constants.PeerTypeSelf {
	//    inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerUser(&mtproto.InputPeer{
	//        UserId:                 md.UserId,
	//        AccessHash:             request.Peer.AccessHash,
	//    }).To_InputPeer())
	//} else if inputPeer.GetPeerType() == constants.PeerTypeChat {
	//    inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerChannel(&mtproto.InputPeer{
	//        ChannelId:                 request.Peer.ChatId,
	//        AccessHash:             request.Peer.AccessHash,
	//    }).To_InputPeer())
	//}
	//
	//var seqID sequence.SeqID
	//var err error
	//if inputPeer.GetPeerType() == constants.PeerTypeUser {
	//    seqID, err = pts_sequence.GetPtsSequence().CurrentPts(md.UserId, inputPeer.GetPeerId())
	//} else {
	//    seqID, err = chat_pts_sequence.GetPtsSequence().CurrentPts(md.UserId, inputPeer.GetPeerId())
	//}
	//if err != nil {
	//    log.Errorf("MessagesReadMentions %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//return mtproto.NewTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
	//    Pts:                  seqID.ToInt32(),
	//    PtsCount:             0,
	//    Offset:               0,
	//}).To_Messages_AffectedHistory(), nil

	panic(fmt.Sprintf("%+v", request))
}
