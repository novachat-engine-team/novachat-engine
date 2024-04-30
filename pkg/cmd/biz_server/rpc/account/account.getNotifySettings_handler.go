/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getNotifySettings_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/account/setting"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
)

//  + TL_inputNotifyPeer
//  + TL_inputNotifyUsers
//  + TL_inputNotifyChats
//  + TL_inputNotifyBroadcasts
//  + TL_inputNotifyAll

//  setting
//  account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
//
func (s *AccountServiceImpl) AccountGetNotifySettings(ctx context.Context, request *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AccountGetNotifySettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	peerNotify := input.MakePeerNotify(request.Peer, constants.PeerTypeFromUserIDType(md.UserId))
	if peerNotify.GetPeerNotifyType() == constants.PeerNotifyInputNotifyAll {
		return mtproto.NewPeerNotifySettingEmptyLayer(md.Layer), nil
	}

	// InputNotifyPeer <--
	//  + TL_inputNotifyPeer
	//  + TL_inputNotifyUsers
	//  + TL_inputNotifyChats
	//  + TL_inputNotifyAll
	//  + TL_inputNotifyBroadcasts
	switch request.Peer.ClassName {
	case mtproto.ClassInputNotifyChats:
		fallthrough
	case mtproto.ClassInputNotifyUsers:
		fallthrough
	case mtproto.ClassInputNotifyBroadcasts:
		fallthrough
	case mtproto.ClassInputNotifyAll:
		return mtproto.NewTLPeerNotifySettings(&mtproto.PeerNotifySettings{
			ShowPreviewsAF509D2085: mtproto.ToMTBool(true),
			SilentAF509D2085:       mtproto.ToMTBool(true),
			MuteUntil:              0,
			Sound:                  "",
			ShowPreviews9ACDA4C071: true,
			Silent9ACDA4C071:       true,
		}).To_PeerNotifySettings(), nil
	case mtproto.ClassInputNotifyPeer:
		fallthrough
	default:
		inputPeer := input.MakeInputPeer(request.Peer.Peer)
		settingV, err := s.accountSettingCore.GetNotifySetting(
			&setting.NotifyPeerSettingType{
				UserId:         md.UserId,
				PeerId:         peerNotify.GetPeerId(),
				PeerType:       inputPeer.GetPeerType(),
				PeerNotifyType: peerNotify.GetPeerNotifyType(),
			})
		if err != nil {
			log.Errorf("AccountGetNotifySettings request:%v error:%s", request, err.Error())
			return nil, err
		}

		if settingV == nil {
			return mtproto.NewPeerNotifySettingEmptyLayer(md.Layer), nil
		}

		return mtproto.NewTLPeerNotifySettings(&mtproto.PeerNotifySettings{
			ShowPreviewsAF509D2085: mtproto.ToMTBool(settingV.ShowPreviews),
			SilentAF509D2085:       mtproto.ToMTBool(settingV.Silent),
			MuteUntil:              settingV.MuteUntil,
			Sound:                  settingV.Sound,
			ShowPreviews9ACDA4C071: settingV.ShowPreviews,
			Silent9ACDA4C071:       settingV.Silent,
		}).To_PeerNotifySettings(), nil
	}
}
