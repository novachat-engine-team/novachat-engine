/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getPeerSettings_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/input"
)

//  messages.getPeerSettings#3672e09c peer:InputPeer = PeerSettings;
//
func (s *MessagesServiceImpl) MessagesGetPeerSettings(ctx context.Context, request *mtproto.TLMessagesGetPeerSettings) (*mtproto.PeerSettings, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetPeerSettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	inputPeer := input.MakeInputPeer(request.Peer)
	peerSetting := s.accountSettingCore.PeerSettings(md.UserId, inputPeer.GetPeerId())
	contactData, err := s.accountContactCore.GetContactById(md.UserId, inputPeer.GetPeerId())
	if err != nil {
		log.Warnf("MessagesGetPeerSettings request: %v error:%s", request, err.Error())
	} else {
		if contactData != nil {
			peerSetting.AddContact = contactData.PeerId != inputPeer.GetPeerId()
			peerSetting.ShareContact = contactData.PeerId == inputPeer.GetPeerId()
			peerSetting.BlockContact = contactData.PeerId == inputPeer.GetPeerId()
		}
	}
	return peerSetting, nil
}
