/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package appconfig

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"novachat_engine/pkg/log"
	"os"
)

type EmojiesSoundsConfig struct {
	Id                  int64
	AccessHash          int64
	FileReferenceBase64 string
}

type EmojiesSendDiceSuccessConfig struct {
}

type AppConfig struct {
	DialogFiltersEnabled bool     `json:"dialog_filters_enabled"`
	DialogFiltersTooltip bool     `json:"dialog_filters_tooltip"`
	YoutubePip           string   `json:"youtube_pip"`
	BackgroundConnection bool     `json:"background_connection"`
	KeepAliveService     bool     `json:"keep_alive_service"`
	QrLoginCamera        bool     `json:"qr_login_camera"`
	QRLoginCode          string   `json:"qr_login_code"` // "disabled", "primary" or "secondary" primary/secondary/disabled
	SaveGifsWithStickers bool     `json:"save_gifs_with_stickers"`
	AutoLoginDomains     []string `json:"autologin_domains"`
	AutoLoginToken       string   `json:"autologin_token"`
	EmojiesSendDice      []string `json:"emojies_send_dice"`
	GifSearchEmojies     []string `json:"gif_search_emojies"`
	//EmojiesSendDiceSuccess EmojiesSendDiceSuccessConfig `json:"emojies_send_dice_success"`
	AutoArchiveSettingAvailable bool     `json:"autoarchive_setting_available"`
	StickersEmojiSuggestOnlyApi bool     `json:"stickers_emoji_suggest_only_api"`
	ExportRegex                 []string `json:"export_regex"`
	ExportGroupUrls             []string `json:"export_group_urls"`
	ExportPrivateUrls           []string `json:"export_private_urls"`
	PendingSuggestions          []string `json:"pending_suggestions"`
	//EmojiesSounds EmojiesSoundsConfig `json:"emojies_sounds"`
	ChatReadMarkExpirePeriod  int32 `json:"chat_read_mark_expire_period"`
	ChatReadMarkSizeThreshold int32 `json:"chat_read_mark_size_threshold"`
}

var DefaultAppConfig = &AppConfig{
	DialogFiltersEnabled:        false,
	DialogFiltersTooltip:        false,
	YoutubePip:                  "",
	BackgroundConnection:        true,
	KeepAliveService:            true,
	QrLoginCamera:               false,
	QRLoginCode:                 "primary",
	SaveGifsWithStickers:        false,
	AutoLoginDomains:            nil,
	AutoLoginToken:              "",
	EmojiesSendDice:             nil,
	GifSearchEmojies:            nil,
	AutoArchiveSettingAvailable: false,
	StickersEmojiSuggestOnlyApi: false,
	ExportRegex:                 nil,
	ExportGroupUrls:             nil,
	ExportPrivateUrls:           nil,
	PendingSuggestions:          nil,
	ChatReadMarkExpirePeriod:    604800,
	ChatReadMarkSizeThreshold:   100,
}

func InstallAppConfig(file string) {
	sourceFileStat, err := os.Stat(file)
	if err != nil {
		log.Warnf("file:%s error:%s", file, err.Error())
		return
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Warnf("%s is not a regular file", file)
		return
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Warnf("%s ReadFile error:%s", file, err.Error())
		return
	}

	ac := &AppConfig{}
	err = jsoniter.Unmarshal(data, ac)
	if err != nil {
		log.Warnf("%s Unmarshal error:%s", file, err.Error())
		return
	}

	*DefaultAppConfig = *ac
	log.Debugf("AppConfig %s load ok", file)
}

// {
//      "emojies_animated_zoom": 1.0,
//"dialog_filters_enabled": false,
//"dialog_filters_tooltip": false,
//"youtube_pip": "",
//"background_connection": true,
//"keep_alive_service": true,
//"qr_login_camera": true,
//"save_gifs_with_stickers": true,
//"autologin_domains": ["11", "222"],
//"autologin_token": "",
//"emojies_send_dice": ["1", "2", "3"],
//"gif_search_emojies": ["1"],
////"emojies_send_dice_success":
//"autoarchive_setting_available": false,
//"stickers_emoji_suggest_only_api": false,
//"export_regex": ["url"],
//"export_group_urls": ["url"],
//"export_private_urls": ["private_url"],
//"pending_suggestions": ["suggestion"],
////emojies_sounds
// }
//

//{ jsonObjectValue
//key: "emojies_send_dice" [STRING],
//value: { jsonArray
//value: [ vector<0x0>
//{ jsonString
//value: "🎲" [STRING],
//},
//{ jsonString
//value: "🎯" [STRING],
//},
//{ jsonString
//value: "🏀" [STRING],
//},
//{ jsonString
//value: "⚽" [STRING],
//},
//{ jsonString
//value: "⚽️" [STRING],
//},
//{ jsonString
//value: "🎰" [STRING],
//},
//{ jsonString
//value: "🎳" [STRING],
//},
//],
//},
//},

//{ jsonObjectValue
//key: "emojies_send_dice_success" [STRING],
//value: { jsonObject
//value: [ vector<0x0>
//{ jsonObjectValue
//key: "🎯" [STRING],
//value: { jsonObject
//value: [ vector<0x0>
//{ jsonObjectValue
//key: "value" [STRING],
//value: { jsonNumber
//value: 6 [DOUBLE],
//},
//},
//{ jsonObjectValue
//key: "frame_start" [STRING],
//value: { jsonNumber
//value: 62 [DOUBLE],
//},
//},
//],
//},
//},
//{ jsonObjectValue
//key: "🏀" [STRING],
//value: { jsonObject
//value: [ vector<0x0>
//{ jsonObjectValue
//key: "value" [STRING],
//value: { jsonNumber
//value: 5 [DOUBLE],
//},
//},
//{ jsonObjectValue
//key: "frame_start" [STRING],
//value: { jsonNumber
//value: 110 [DOUBLE],
//},
//},
//],
//},
//},
//{ jsonObjectValue
//key: "⚽" [STRING],
//value: { jsonObject
//value: [ vector<0x0>
//{ jsonObjectValue
//key: "value" [STRING],
//value: { jsonNumber
//value: 5 [DOUBLE],
//},
//},
//{ jsonObjectValue
//key: "frame_start" [STRING],
//value: { jsonNumber
//value: 110 [DOUBLE],
//},
//},
//],
//},
//},
//{ jsonObjectValue
//key: "⚽️" [STRING],
//value: { jsonObject
//value: [ vector<0x0>
//{ jsonObjectValue
//key: "value" [STRING],
//value: { jsonNumber
//value: 5 [DOUBLE],
//},
//},
//{ jsonObjectValue
//key: "frame_start" [STRING],
//value: { jsonNumber
//value: 110 [DOUBLE],
//},
//},
//],
//},
//},
//{ jsonObjectValue
//key: "🎰" [STRING],
//value: { jsonObject
//value: [ vector<0x0>
//{ jsonObjectValue
//key: "value" [STRING],
//value: { jsonNumber
//value: 64 [DOUBLE],
//},
//},
//{ jsonObjectValue
//key: "frame_start" [STRING],
//value: { jsonNumber
//value: 110 [DOUBLE],
//},
//},
//],
//},
//},
//{ jsonObjectValue
//key: "🎳" [STRING],
//value: { jsonObject
//value: [ vector<0x0>
//{ jsonObjectValue
//key: "value" [STRING],
//value: { jsonNumber
//value: 6 [DOUBLE],
//},
//},
//{ jsonObjectValue
//key: "frame_start" [STRING],
//value: { jsonNumber
//value: 110 [DOUBLE],
//},
//},
//],
//},
//},
//],
//},
//},

//
//{ jsonObjectValue
//key: "gif_search_branding" [STRING],
//value: { jsonString
//value: "tenor" [STRING],
//},
//},
//{ jsonObjectValue
//key: "gif_search_emojies" [STRING],
//value: { jsonArray
//value: [ vector<0x0>
//{ jsonString
//value: "👍" [STRING],
//},
//{ jsonString
//value: "😘" [STRING],
//},
//{ jsonString
//value: "😍" [STRING],
//},
//{ jsonString
//value: "😡" [STRING],
//},
//{ jsonString
//value: "🥳" [STRING],
//},
//{ jsonString
//value: "😂" [STRING],
//},
//{ jsonString
//value: "😮" [STRING],
//},
//{ jsonString
//value: "🙄" [STRING],
//},
//{ jsonString
//value: "😎" [STRING],
//},
//{ jsonString
//value: "👎" [STRING],
//},
//],
//},
//},
