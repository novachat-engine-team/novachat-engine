/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.setGlobalPrivacySettings_handler.go
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
)

// AccountSetGlobalPrivacySettings
// attach appconfig:autoarchive_setting_available
//  account.setGlobalPrivacySettings#1edaaac2 settings:GlobalPrivacySettings = GlobalPrivacySettings;
//
func (s *AccountServiceImpl) AccountSetGlobalPrivacySettings(ctx context.Context, request *mtproto.TLAccountSetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountSetGlobalPrivacySettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	mute := false
	if request.Settings.ArchiveAndMuteNewNoncontactPeers != nil {
		mute = mtproto.ToBool(request.Settings.ArchiveAndMuteNewNoncontactPeers)
	}
	mute, err := s.accountSettingCore.UpdateNotifySetting(&setting.NotifyPeerSettingType{
		UserId:         md.UserId,
		PeerId:         0,
		PeerType:       0,
		PeerNotifyType: constants.PeerGlobalSetting,
	}, false, mute, 0, "")
	if err != nil {
		log.Errorf("AccountSetGlobalPrivacySettings - %v, request: %v", request)
		return nil, err
	}
	return request.Settings, nil
}
