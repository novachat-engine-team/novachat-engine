/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/19 19:20
 * @File : lang.go
 */

package lang

import (
	"encoding/json"
	"io/ioutil"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/common/errors"
	"strings"
)

const (
	LANG_PACK_EN_FILE            = "lang_pack_en.json"
	LANG_PACK_RU_FILE            = "lang_pack_ru.json"
	LANG_PACK_CLASSIC_ZH_CN_FILE = "lang_pack_classic-zh-cn.json"
	LANG_PACK_ZH_HANS_RAW_FILE   = "lang_pack_zh-hans-raw.json"
	LANG_PACK_ZH_CN_FILE         = "lang_pack_zh_cn.json"
	LANG_PACK_VI_FILE            = "lang_pack_vi.json"
)

type LangPacks struct {
	LangCode          string                    `json:"lang_code,omitempty"`
	Version           int32                     `json:"version,omitempty"`
	Strings           []*mtproto.LangPackString `json:"strings,omitempty"`
	StringPluralizeds []*mtproto.LangPackString `json:"string_pluralizeds,omitempty"`
	StringDeleteds    []*mtproto.LangPackString `json:"string_deleteds,omitempty"`
}

func (m *LangPacks) GetString(key string) string {
	for _, s := range m.Strings {
		if strings.Compare(s.GetKey(), key) == 0 {
			return s.GetValue()
		}
	}
	return ""
}

var (
	langPacksEn          LangPacks
	langPacksRu          LangPacks
	langPacksClassicZhCN LangPacks // "classic-zh-cn"
	langPacksZhHansRaw   LangPacks // "zh-hans-raw"
	langPacksZhCN        LangPacks // zh_CN
	langPacksVi          LangPacks // Vi
)

func (m *LangPacks) Query(k string) string {
	for _, lp := range m.Strings {
		if lp.GetKey() == k {
			return lp.GetValue()
		}
	}
	return ""
}

func loadLangPacks(filePath string, packs *LangPacks) {
	jData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("loadLangPacks file=[%s] error:%s", filePath, err.Error())
		return
	}

	err = json.Unmarshal([]byte(jData), packs)
	if err != nil {
		log.Error(err.Error())
		return
	}

	for _, s := range packs.Strings {
		s.ClassName = mtproto.ClassLangPackString
		s.Cmd = mtproto.Cmd_LangPackString
	}

	for _, s := range packs.StringDeleteds {
		s.ClassName = mtproto.ClassLangPackStringDeleted
		s.Cmd = mtproto.Cmd_LangPackStringDeleted
	}

	for _, s := range packs.StringPluralizeds {
		s.ClassName = mtproto.ClassLangPackStringPluralized
		s.Cmd = mtproto.Cmd_LangPackStringPluralized
	}
}

func InstallLangClient(pathName string) {
	loadLangPacks(pathName + LANG_PACK_EN_FILE, &langPacksEn)
	loadLangPacks(pathName + LANG_PACK_RU_FILE, &langPacksRu)
	loadLangPacks(pathName + LANG_PACK_CLASSIC_ZH_CN_FILE, &langPacksClassicZhCN)
	loadLangPacks(pathName + LANG_PACK_ZH_HANS_RAW_FILE, &langPacksZhHansRaw)
	loadLangPacks(pathName + LANG_PACK_ZH_CN_FILE, &langPacksZhCN)
	loadLangPacks(pathName + LANG_PACK_VI_FILE, &langPacksVi)
}

func GetLangPacks(langCode string) *LangPacks {
	var langPacks *LangPacks

	switch langCode {
	case "en":
		langPacks = &langPacksEn
	case "ru":
		langPacks = &langPacksRu
	case "classic-zh-cn":
		langPacks = &langPacksClassicZhCN
	case "zh-hans-raw":
		langPacks = &langPacksZhHansRaw
	case "zh_CN":
		langPacks = &langPacksZhCN
	case "vi":
		langPacks = &langPacksVi
	default:
		langPacks = &langPacksZhCN
	}

	return langPacks
}

func GetLanguage(langCode string) (*mtproto.LangPackLanguage, error) {
	var language *mtproto.LangPackLanguage
	switch langCode {
	case "en":
		language = mtproto.NewTLLangPackLanguage(&mtproto.LangPackLanguage{
			Official: true,
			Name:       "English",
			NativeName: "English",
			LangCode:   "en",
			PluralCode:      "en",
			StringsCount:    1978,
			TranslatedCount: 1978,
			TranslationsUrl: "https://translations.telegram.org/en/",
		}).To_LangPackLanguage()
	case "ru":
		language = mtproto.NewTLLangPackLanguage(&mtproto.LangPackLanguage{
			Official: true,
			Name:       "Russian",
			NativeName: "Русский",
			// BaseLangCode:
			LangCode:        "ru",
			StringsCount:    1978,
			TranslatedCount: 1976,
			TranslationsUrl: "https://translations.telegram.org/ru/",
		}).To_LangPackLanguage()
	case "classic-zh-cn":
		language = mtproto.NewTLLangPackLanguage(&mtproto.LangPackLanguage{
			Name:            "Chinese (Simplified)",
			NativeName:      "简体中文",
			LangCode:        "classic-zh-cn",
			BaseLangCode:    "zh-hans-raw",
			StringsCount:    2005,
			TranslatedCount: 2002,
			TranslationsUrl: "https://translations.telegram.org/classic-zh-cn/",
		}).To_LangPackLanguage()
	case "zh_CN":
		language = mtproto.NewTLLangPackLanguage(&mtproto.LangPackLanguage{
			Name:            "Chinese (Simplified, @zh_CN)",
			NativeName:      "简体中文 (@zh_CN 版)",
			LangCode:        "zh_CN",
			BaseLangCode:    "zh-hans-raw",
			StringsCount:    2005,
			TranslatedCount: 2002,
			TranslationsUrl: "https://translations.telegram.org/classic-zh-cn/",
		}).To_LangPackLanguage()
	case "vi":
		language = mtproto.NewTLLangPackLanguage(&mtproto.LangPackLanguage{
			Name:       "Vietnamese",
			NativeName: "Thổ nhĩ kỳ",
			LangCode:   "vi",
			StringsCount:    2005,
			TranslatedCount: 2002,
			TranslationsUrl: "https://translations.telegram.org/vi/",
		}).To_LangPackLanguage()
	default:
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_LANG_PACK_INVALID)
		log.Errorf("langpack.getLanguage#6a596502 - error: %v", err)
		return nil, err
	}
	return language, nil
}

func GetLanguages() *mtproto.Vector_LangPackLanguage {
	languages := &mtproto.Vector_LangPackLanguage{}
	language := mtproto.NewTLLangPackLanguage(&mtproto.LangPackLanguage{
		Official: true,
		// Rtl: false,
		// Beta: false,
		Name:       "English",
		NativeName: "English",
		LangCode:   "en",
		// BaseLangCode:
		PluralCode:      "en",
		StringsCount:    1978,
		TranslatedCount: 1978,
		TranslationsUrl: "https://translations.telegram.org/en/",
	})
	languages.Datas = append(languages.Datas, language.To_LangPackLanguage())

	languageRu := mtproto.NewTLLangPackLanguage(&mtproto.LangPackLanguage{
		Official: true,
		// Rtl: false,
		// Beta: false,
		Name:       "Russian",
		NativeName: "Русский",
		// BaseLangCode:
		LangCode:        "ru",
		StringsCount:    1978,
		TranslatedCount: 1976,
		TranslationsUrl: "https://translations.telegram.org/ru/",
	})
	languages.Datas = append(languages.Datas, languageRu.To_LangPackLanguage())

	languageZhCN := mtproto.NewTLLangPackLanguage(&mtproto.LangPackLanguage{
		// Official: false,
		// Rtl:      false,
		// Beta:            false,
		Name:            "Chinese (Simplified, @zh_CN)",
		NativeName:      "简体中文 (@zh_CN 版)",
		LangCode:        "zh_CN",
		BaseLangCode:    "zh-hans-raw",
		StringsCount:    2005,
		TranslatedCount: 2002,
		TranslationsUrl: "https://translations.telegram.org/classic-zh-cn/",
	})
	languages.Datas = append(languages.Datas, languageZhCN.To_LangPackLanguage())

	languageVi := mtproto.NewTLLangPackLanguage(&mtproto.LangPackLanguage{
		// Official: false,
		// Rtl:      false,
		// Beta:            false,
		Name:       "Vietnamese",
		NativeName: "Thổ nhĩ kỳ",
		LangCode:   "vi",
		//BaseLangCode:    &types.StringValue{Value: "zh-hans-raw"},
		StringsCount:    2005,
		TranslatedCount: 2002,
		TranslationsUrl: "https://translations.telegram.org/vi/",
	})
	languages.Datas = append(languages.Datas, languageVi.To_LangPackLanguage())

	return languages
}