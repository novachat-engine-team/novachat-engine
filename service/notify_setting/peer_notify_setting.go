/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/10 16:33
 * @File : peer_notify_setting.go
 */

package notify_setting

import (
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
	data_setting "novachat_engine/service/data/setting"
)

type PeerNotifySettings struct {
	UserId         int32
	PeerId         int32
	PeerNotifyType constants.PeerNotifyType
	Setting        *mtproto.PeerNotifySettings
}

func InputPeerNotifySettingToNotifySetting(setting *mtproto.InputPeerNotifySettings) *mtproto.PeerNotifySettings {
	showPreviews := setting.ShowPreviews38935EB271
	if setting.ShowPreviews9C3D198E85 != nil {
		showPreviews = mtproto.ToBool(setting.ShowPreviews9C3D198E85)
	}
	silent := setting.Silent38935EB271
	if setting.Silent9C3D198E85 != nil {
		silent = mtproto.ToBool(setting.Silent9C3D198E85)
	}
	return mtproto.NewTLPeerNotifySettings(&mtproto.PeerNotifySettings{
		ShowPreviewsAF509D2085: mtproto.ToMTBool(showPreviews),
		SilentAF509D2085:       mtproto.ToMTBool(silent),
		MuteUntil:              setting.MuteUntil,
		Sound:                  setting.Sound,
		ShowPreviews9ACDA4C071: showPreviews,
		Silent9ACDA4C071:       silent,
	}).To_PeerNotifySettings()
}

func ToPeerNotifySettings(peerNotifySetting *data_setting.NotifySetting, layer int32) *mtproto.PeerNotifySettings {
	if peerNotifySetting == nil {
		return mtproto.NewPeerNotifySettingEmptyLayer(layer)
	}

	return mtproto.NewTLPeerNotifySettings(&mtproto.PeerNotifySettings{
		ShowPreviewsAF509D2085: mtproto.ToMTBool(peerNotifySetting.ShowPreviews),
		SilentAF509D2085:       mtproto.ToMTBool(peerNotifySetting.Silent),
		MuteUntil:              peerNotifySetting.MuteUntil,
		Sound:                  peerNotifySetting.Sound,
		ShowPreviews9ACDA4C071: peerNotifySetting.ShowPreviews,
		Silent9ACDA4C071:       peerNotifySetting.Silent,
	}).To_PeerNotifySettings()
}
