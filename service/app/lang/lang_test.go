/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/20 18:55
 * @File : lang_test.go
 */

package lang

import (
	"encoding/json"
	"io/ioutil"
	"novachat_engine/mtproto"
	"testing"
)

type StringValue struct {
	Value string `json:"value"`
}

type testLangPackString struct {
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value"`
	ZeroValue            StringValue   `protobuf:"bytes,5,opt,name=zero_value,json=zeroValue,proto3" json:"zero_value"`
	OneValue             StringValue   `protobuf:"bytes,6,opt,name=one_value,json=oneValue,proto3" json:"one_value"`
	TwoValue             StringValue   `protobuf:"bytes,7,opt,name=two_value,json=twoValue,proto3" json:"two_value"`
	FewValue             StringValue   `protobuf:"bytes,8,opt,name=few_value,json=fewValue,proto3" json:"few_value"`
	ManyValue            StringValue   `protobuf:"bytes,9,opt,name=many_value,json=manyValue,proto3" json:"many_value"`
	OtherValue           string   `protobuf:"bytes,10,opt,name=other_value,json=otherValue,proto3" json:"other_value"`
}

type testLangPacks struct {
	LangCode          string                    `json:"lang_code,omitempty"`
	Version           int32                     `json:"version,omitempty"`
	Strings           []*testLangPackString `json:"strings,omitempty"`
	StringPluralizeds []*testLangPackString `json:"string_pluralizeds,omitempty"`
	StringDeleteds    []*testLangPackString `json:"string_deleteds,omitempty"`
}
//
//var (
//	langPacksEn          LangPacks
//	langPacksRu          LangPacks
//	langPacksClassicZhCN LangPacks // "classic-zh-cn"
//	langPacksZhHansRaw   LangPacks // "zh-hans-raw"
//	langPacksZhCN        LangPacks // zh_CN
//	langPacksVi          LangPacks // Vi
//)

func testLoadLangPacks(filePath string, packs *testLangPacks) {
	jData, err := ioutil.ReadFile(filePath)
	if err != nil {
		//log.Errorf("loadLangPacks file=[%s] error:%s", filePath, err.Error())
		return
	}

	err = json.Unmarshal([]byte(jData), packs)
	if err != nil {
		//log.Error(err.Error())
		return
	}
}

func testSaveLangPacks(filePath string, packs *LangPacks) {
	buf, err := json.MarshalIndent(packs, "", "\t")
	if err != nil {
		panic(err)
		return
	}

	ioutil.WriteFile(filePath, buf, 0644)
}


var (
	testLangPacksEn          testLangPacks
	testLangPacksRu          testLangPacks
	testLangPacksClassicZhCN testLangPacks // "classic-zh-cn"
	testLangPacksZhHansRaw   testLangPacks // "zh-hans-raw"
	testLangPacksZhCN        testLangPacks // zh_CN
	testLangPacksVi          testLangPacks // Vi`
)

func testLangPacks2LangPacks(in *testLangPacks, out *LangPacks) {
	out.LangCode = in.LangCode
	out.Version = in.Version

	for idx := range in.StringDeleteds {
		t := in.StringDeleteds[idx]
		out.StringDeleteds = append(out.StringDeleteds, &mtproto.LangPackString{
			Key:                  t.Key,
			Value:                t.Value,
			ZeroValue:            t.ZeroValue.Value,
			OneValue:             t.OneValue.Value,
			TwoValue:             t.TwoValue.Value,
			FewValue:             t.FewValue.Value,
			ManyValue:            t.ManyValue.Value,
			OtherValue:           t.OtherValue,
		})
	}

	for idx := range in.StringPluralizeds {
		t := in.StringPluralizeds[idx]
		out.StringPluralizeds = append(out.StringPluralizeds, &mtproto.LangPackString{
			Key:                  t.Key,
			Value:                t.Value,
			ZeroValue:            t.ZeroValue.Value,
			OneValue:             t.OneValue.Value,
			TwoValue:             t.TwoValue.Value,
			FewValue:             t.FewValue.Value,
			ManyValue:            t.ManyValue.Value,
			OtherValue:           t.OtherValue,
		})
	}

	for idx := range in.Strings {
		t := in.Strings[idx]
		out.Strings = append(out.Strings, &mtproto.LangPackString{
			Key:                  t.Key,
			Value:                t.Value,
			ZeroValue:            t.ZeroValue.Value,
			OneValue:             t.OneValue.Value,
			TwoValue:             t.TwoValue.Value,
			FewValue:             t.FewValue.Value,
			ManyValue:            t.ManyValue.Value,
			OtherValue:           t.OtherValue,
		})
	}
}

func TestInstallLangClient(t *testing.T) {
	pathName := "E:\\test\\novachat_engine\\server\\biz\\"

	testLoadLangPacks(pathName + LANG_PACK_EN_FILE, &testLangPacksEn)
	testLoadLangPacks(pathName + LANG_PACK_RU_FILE, &testLangPacksRu)
	testLoadLangPacks(pathName + LANG_PACK_CLASSIC_ZH_CN_FILE, &testLangPacksClassicZhCN)
	testLoadLangPacks(pathName + LANG_PACK_ZH_HANS_RAW_FILE, &testLangPacksZhHansRaw)
	testLoadLangPacks(pathName + LANG_PACK_ZH_CN_FILE, &testLangPacksZhCN)
	testLoadLangPacks(pathName + LANG_PACK_VI_FILE, &testLangPacksVi)


	testLangPacks2LangPacks(&testLangPacksEn, &langPacksEn)
	testLangPacks2LangPacks(&testLangPacksRu, &langPacksRu)

	testLangPacks2LangPacks(&testLangPacksClassicZhCN, &langPacksClassicZhCN)
	testLangPacks2LangPacks(&testLangPacksZhHansRaw, &langPacksZhHansRaw)
	testLangPacks2LangPacks(&testLangPacksZhCN, &langPacksZhCN)
	testLangPacks2LangPacks(&testLangPacksVi, &langPacksVi)

	testSaveLangPacks(pathName + LANG_PACK_EN_FILE + "1", &langPacksEn)
	testSaveLangPacks(pathName + LANG_PACK_RU_FILE+ "1", &langPacksRu)
	testSaveLangPacks(pathName + LANG_PACK_CLASSIC_ZH_CN_FILE+ "1", &langPacksClassicZhCN)
	testSaveLangPacks(pathName + LANG_PACK_ZH_HANS_RAW_FILE+ "1", &langPacksZhHansRaw)
	testSaveLangPacks(pathName + LANG_PACK_ZH_CN_FILE+ "1", &langPacksZhCN)
	testSaveLangPacks(pathName + LANG_PACK_VI_FILE+ "1", &langPacksVi)
}
