/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getGlobalPrivacySettings_handler.go
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

//  account.getGlobalPrivacySettings#eb2b4cf6 = GlobalPrivacySettings;
//
func (s *AccountServiceImpl) AccountGetGlobalPrivacySettings(ctx context.Context, request *mtproto.TLAccountGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountGetGlobalPrivacySettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	mute, err := s.accountSettingCore.GetNotifySetting(
		&setting.NotifyPeerSettingType{
			UserId:         md.UserId,
			PeerId:         0,
			PeerType:       constants.PeerTypeUser,
			PeerNotifyType: constants.PeerGlobalSetting,
		})
	if err != nil {
		log.Errorf("AccountGetGlobalPrivacySettings - request: %v", request)
		return nil, err
	}

	_ = mute
	return mtproto.NewTLGlobalPrivacySettings(&mtproto.GlobalPrivacySettings{
		ArchiveAndMuteNewNoncontactPeers: mtproto.ToMTBool(false),
	}).To_GlobalPrivacySettings(), nil
}
