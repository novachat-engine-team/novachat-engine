/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.updateNotifySettings_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/account/setting"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/input"
	"novachat_engine/service/notify_setting"
	"time"
)

//  account.updateNotifySettings#84be5b93 peer:InputNotifyPeer settings:InputPeerNotifySettings = Bool;
//
func (s *AccountServiceImpl) AccountUpdateNotifySettings(ctx context.Context, request *mtproto.TLAccountUpdateNotifySettings) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountUpdateNotifySettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	//panic("AccountUpdateNotifySettings")

	peerNotify := input.MakePeerNotify(request.Peer, constants.PeerTypeFromUserIDType(md.UserId))
	if peerNotify.GetPeerNotifyType() == constants.PeerNotifyInputNotifyAll {
		peerNotify = input.MakePeerNotify(mtproto.NewTLInputNotifyAll(nil).To_InputNotifyPeer(), constants.PeerTypeFromUserIDType(md.UserId))
		return mtproto.ToMTBool(true), nil
	}

	showPreviews := request.Settings.ShowPreviews38935EB271
	if request.Settings.ShowPreviews9C3D198E85 != nil {
		showPreviews = mtproto.ToBool(request.Settings.ShowPreviews9C3D198E85)
	}
	silent := request.Settings.Silent38935EB271
	if request.Settings.Silent9C3D198E85 != nil {
		silent = mtproto.ToBool(request.Settings.Silent9C3D198E85)
	}

	ok, err := s.accountSettingCore.UpdateNotifySetting(
		&setting.NotifyPeerSettingType{
			UserId:         md.UserId,
			PeerId:         peerNotify.GetPeerId(),
			PeerType:       peerNotify.GetPeerType(),
			PeerNotifyType: peerNotify.GetPeerNotifyType(),
		},
		showPreviews,
		silent,
		request.Settings.MuteUntil,
		request.Settings.Sound)
	if err != nil {
		log.Errorf("AccountUpdateNotifySettings - request: %v error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
	}

	if ok == true {
		go func() {
			update := mtproto.NewTLUpdateNotifySettings(&mtproto.Update{
				PeerBEC268EF71: peerNotify.ToNotifyPeer(),
				NotifySettings: notify_setting.InputPeerNotifySettingToNotifySetting(request.Settings),
			})

			_, err = syncService.GetSyncClientById(md.UserId).ReqSyncUpdate(context.Background(), &syncService.SyncUpdate{
				UserId:          md.UserId,
				IgnoreAuthKeyId: md.AuthKeyId,
				Updates: mtproto.NewTLUpdates(&mtproto.Updates{
					Updates: []*mtproto.Update{update.To_Update()},
					Users:   []*mtproto.User{},
					Chats:   []*mtproto.Chat{},
					Date:    int32(time.Now().Unix()),
					Seq:     0,
				}).To_Updates(),
				PeerType: constants.PeerTypeUser.ToInt32(),
			})
			if err != nil {
				log.Errorf("AccountUpdateNotifySettings PushUpdate userId:%d authKeyId:%d error:%s", md.UserId, md.AuthKeyId, err.Error())
			}
		}()
	}

	return mtproto.ToMTBool(true), nil
}
