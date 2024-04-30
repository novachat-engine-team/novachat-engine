/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/13 9:47
 * @File : object.util.go
 */

package mtproto

// layer 71 TLPeerNotifySettingsEmpty
// layer 122  TLPeerNotifySettings replace TLPeerNotifySettingsEmpty
func NewPeerNotifySettingEmptyLayer(layer int32) *PeerNotifySettings {
	if layer <= 71 {
		return NewTLPeerNotifySettingsEmpty(nil).To_PeerNotifySettings()
	}

	return NewTLPeerNotifySettings(nil).To_PeerNotifySettings()
}
