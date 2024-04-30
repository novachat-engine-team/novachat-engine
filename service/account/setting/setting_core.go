/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package setting

import (
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/setting"
	data_setting "novachat_engine/service/data/setting"
)

func makePeerId(peerId int64, peerType constants.PeerType) int64 {
	if peerType == constants.PeerTypeChat {
		return -peerId
	}
	return peerId
}

type NotifyPeerSettingType struct {
	UserId         int64
	PeerId         int64
	PeerType       constants.PeerType
	PeerNotifyType constants.PeerNotifyType
}

type Core struct {
	core *setting.Core
}

func (c *Core) PeerSettings(userId int64, peerId int64) *mtproto.PeerSettings {
	return mtproto.NewTLPeerSettings(&mtproto.PeerSettings{
		ReportSpam:            userId != peerId,
		AddContact:            userId != peerId,
		BlockContact:          userId != peerId,
		ShareContact:          userId != peerId,
		NeedContactsException: false,
		ReportGeo:             false,
		Autoarchived:          false,
		GeoDistance:           0,
	}).To_PeerSettings()
}

func (c *Core) GetNotifySetting(notifyPeerSetting *NotifyPeerSettingType) (*data_setting.NotifySetting, error) {
	notifyPeerSetting.PeerId = makePeerId(notifyPeerSetting.PeerId, notifyPeerSetting.PeerType)
	return c.core.GetNotifySetting(notifyPeerSetting.UserId, notifyPeerSetting.PeerId, notifyPeerSetting.PeerNotifyType.ToInt32())
}

func (c *Core) GetNotifySettingList(notifyPeerSettingList []*NotifyPeerSettingType) ([]*data_setting.Setting, error) {
	l := make([]*setting.NotifySettingListType, 0, len(notifyPeerSettingList))
	for _, v := range notifyPeerSettingList {
		v.PeerId = makePeerId(v.PeerId, v.PeerType)
		l = append(l, &setting.NotifySettingListType{
			UserId:         v.UserId,
			PeerId:         v.PeerId,
			PeerType:       v.PeerType.ToInt32(),
			PeerNotifyType: v.PeerNotifyType.ToInt32(),
		})
	}
	settingList, err := c.core.GetNotifySettingList(l)
	if err != nil {
		return nil, err
	}
	for _, v := range settingList {
		v.PeerId = makePeerId(v.PeerId, constants.PeerTypeFromInt32(v.PeerType))
	}
	return settingList, nil
}

func (c *Core) ResetNotifySetting(userId int64) (bool, error) {
	return c.core.ResetNotifySetting(userId, constants.PeerGlobalSetting.ToInt32())
}

func (c *Core) UpdateNotifySetting(notifyPeerSetting *NotifyPeerSettingType, previews bool, silent bool, until int32, sound string) (bool, error) {
	notifyPeerSetting.PeerId = makePeerId(notifyPeerSetting.PeerId, notifyPeerSetting.PeerType)
	return c.core.UpdateNotifySetting(notifyPeerSetting.UserId, notifyPeerSetting.PeerId, notifyPeerSetting.PeerType.ToInt32(), notifyPeerSetting.PeerNotifyType.ToInt32(), previews, silent, until, sound)
}

func NewSettingCore(core *setting.Core) *Core {
	return &Core{
		core: core,
	}
}
