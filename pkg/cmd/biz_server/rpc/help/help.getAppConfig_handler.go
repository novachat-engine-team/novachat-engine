/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getAppConfig_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/app/appconfig"
)

//  help.getAppConfig#98914110 = JSONValue;
//
func (s *HelpServiceImpl) HelpGetAppConfig(ctx context.Context, request *mtproto.TLHelpGetAppConfig) (*mtproto.JSONValue, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("HelpGetAppConfig %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	appConfig := appconfig.DefaultAppConfig

	conf := mtproto.NewTLJsonObject(nil)
	values := make([]*mtproto.JSONObjectValue, 0, 0)

	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "dialog_filters_enabled",
		Value: mtproto.NewTLJsonBool(&mtproto.JSONValue{
			ValueC7345E6A90: mtproto.ToMTBool(appConfig.DialogFiltersEnabled),
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "dialog_filters_tooltip",
		Value: mtproto.NewTLJsonBool(&mtproto.JSONValue{
			ValueC7345E6A90: mtproto.ToMTBool(appConfig.DialogFiltersTooltip),
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "youtube_pip",
		Value: mtproto.NewTLJsonString(&mtproto.JSONValue{
			ValueB71E767A90: appConfig.YoutubePip,
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "background_connection",
		Value: mtproto.NewTLJsonBool(&mtproto.JSONValue{
			ValueC7345E6A90: mtproto.ToMTBool(appConfig.BackgroundConnection),
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "keep_alive_service",
		Value: mtproto.NewTLJsonBool(&mtproto.JSONValue{
			ValueC7345E6A90: mtproto.ToMTBool(appConfig.KeepAliveService),
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "qr_login_camera",
		Value: mtproto.NewTLJsonBool(&mtproto.JSONValue{
			ValueC7345E6A90: mtproto.ToMTBool(appConfig.QrLoginCamera),
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	qrLoginCode := appConfig.QRLoginCode
	if len(s.conf.QRLoginCodeKey) == 0 {
		qrLoginCode = "disabled"
	}
	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "qr_login_code",
		Value: mtproto.NewTLJsonString(&mtproto.JSONValue{
			ValueB71E767A90: qrLoginCode,
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "save_gifs_with_stickers",
		Value: mtproto.NewTLJsonBool(&mtproto.JSONValue{
			ValueC7345E6A90: mtproto.ToMTBool(appConfig.SaveGifsWithStickers),
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	v := make([]*mtproto.JSONValue, 0, len(appConfig.AutoLoginDomains))
	for idx := range appConfig.AutoLoginDomains {
		v = append(v, mtproto.NewTLJsonString(&mtproto.JSONValue{
			ValueB71E767A90: appConfig.AutoLoginDomains[idx],
		}).To_JSONValue())
	}
	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "autologin_domains",
		Value: mtproto.NewTLJsonArray(&mtproto.JSONValue{
			ValueF744476390: v,
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "autologin_token",
		Value: mtproto.NewTLJsonString(&mtproto.JSONValue{
			ValueB71E767A90: appConfig.AutoLoginToken,
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	v = make([]*mtproto.JSONValue, 0, len(appConfig.EmojiesSendDice))
	for idx := range appConfig.EmojiesSendDice {
		v = append(v, mtproto.NewTLJsonString(&mtproto.JSONValue{
			ValueB71E767A90: appConfig.EmojiesSendDice[idx],
		}).To_JSONValue())
	}
	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "emojies_send_dice",
		Value: mtproto.NewTLJsonArray(&mtproto.JSONValue{
			ValueF744476390: v,
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	v = make([]*mtproto.JSONValue, 0, len(appConfig.GifSearchEmojies))
	for idx := range appConfig.GifSearchEmojies {
		v = append(v, mtproto.NewTLJsonString(&mtproto.JSONValue{
			ValueB71E767A90: appConfig.GifSearchEmojies[idx],
		}).To_JSONValue())
	}
	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "gif_search_emojies",
		Value: mtproto.NewTLJsonArray(&mtproto.JSONValue{
			ValueF744476390: v,
		}).To_JSONValue(),
	}).To_JSONObjectValue())
	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "autoarchive_setting_available",
		Value: mtproto.NewTLJsonBool(&mtproto.JSONValue{
			ValueC7345E6A90: mtproto.ToMTBool(appConfig.AutoArchiveSettingAvailable),
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "stickers_emoji_suggest_only_api",
		Value: mtproto.NewTLJsonBool(&mtproto.JSONValue{
			ValueC7345E6A90: mtproto.ToMTBool(appConfig.StickersEmojiSuggestOnlyApi),
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	v = make([]*mtproto.JSONValue, 0, len(appConfig.ExportRegex))
	for idx := range appConfig.ExportRegex {
		v = append(v, mtproto.NewTLJsonString(&mtproto.JSONValue{
			ValueB71E767A90: appConfig.ExportRegex[idx],
		}).To_JSONValue())
	}
	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "export_regex",
		Value: mtproto.NewTLJsonArray(&mtproto.JSONValue{
			ValueF744476390: v,
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	v = make([]*mtproto.JSONValue, 0, len(appConfig.ExportGroupUrls))
	for idx := range appConfig.ExportGroupUrls {
		v = append(v, mtproto.NewTLJsonString(&mtproto.JSONValue{
			ValueB71E767A90: appConfig.ExportGroupUrls[idx],
		}).To_JSONValue())
	}
	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "export_group_urls",
		Value: mtproto.NewTLJsonArray(&mtproto.JSONValue{
			ValueF744476390: v,
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	v = make([]*mtproto.JSONValue, 0, len(appConfig.ExportPrivateUrls))
	for idx := range appConfig.ExportPrivateUrls {
		v = append(v, mtproto.NewTLJsonString(&mtproto.JSONValue{
			ValueB71E767A90: appConfig.ExportPrivateUrls[idx],
		}).To_JSONValue())
	}
	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "export_private_urls",
		Value: mtproto.NewTLJsonArray(&mtproto.JSONValue{
			ValueF744476390: v,
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	v = make([]*mtproto.JSONValue, 0, len(appConfig.PendingSuggestions))
	for idx := range appConfig.PendingSuggestions {
		v = append(v, mtproto.NewTLJsonString(&mtproto.JSONValue{
			ValueB71E767A90: appConfig.PendingSuggestions[idx],
		}).To_JSONValue())
	}
	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "pending_suggestions",
		Value: mtproto.NewTLJsonArray(&mtproto.JSONValue{
			ValueF744476390: v,
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "chat_read_mark_expire_period",
		Value: mtproto.NewTLJsonNumber(&mtproto.JSONValue{
			Value2BE0DFA490: float64(appConfig.ChatReadMarkExpirePeriod),
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	values = append(values, mtproto.NewTLJsonObjectValue(&mtproto.JSONObjectValue{
		Key: "chat_read_mark_size_threshold",
		Value: mtproto.NewTLJsonNumber(&mtproto.JSONValue{
			Value2BE0DFA490: float64(appConfig.ChatReadMarkSizeThreshold),
		}).To_JSONValue(),
	}).To_JSONObjectValue())

	//TODO:(Coder)
	//EmojiesSendDiceSuccess EmojiesSendDiceSuccessConfig `json:"emojies_send_dice_success"`
	//EmojiesSounds EmojiesSoundsConfig `json:"emojies_sounds"`

	conf.SetValue99C1D49D90(values)
	return conf.To_JSONValue(), nil
}
