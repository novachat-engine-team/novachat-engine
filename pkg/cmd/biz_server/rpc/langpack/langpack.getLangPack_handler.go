/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : langpack.getLangPack_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/app/lang"
	"novachat_engine/service/common/errors"
)

//  langpack.getLangPack#f2f2330a lang_pack:string lang_code:string = LangPackDifference;
//  langpack.getLangPack#9ab5c58e lang_code:string = LangPackDifference;
//
func (s *LangpackServiceImpl) LangpackGetLangPack(ctx context.Context, request *mtproto.TLLangpackGetLangPack) (*mtproto.LangPackDifference, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("LangpackGetLangPack %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	langPacks := lang.GetLangPacks(request.GetLangCode())
	if langPacks == nil {
		log.Errorf("LangpackGetLangPack request:%v Invalid LangCode", request)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_LANG_PACK_INVALID_LANGCODE)
	}

	diff := mtproto.NewTLLangPackDifference(nil)
	diff.SetLangCode(request.LangCode)
	diff.SetVersion(langPacks.Version)
	diff.SetFromVersion(0)

	diffStrings := make([]*mtproto.LangPackString, 0)
	for _, strings := range langPacks.Strings {
		ss := mtproto.NewTLLangPackString(nil)
		langPackString := ss.To_LangPackString()
		langPackString.Key = strings.Key
		langPackString.Value = strings.Value
		langPackString.ZeroValue = strings.ZeroValue
		langPackString.OneValue = strings.OneValue
		langPackString.TwoValue = strings.TwoValue
		langPackString.FewValue = strings.FewValue
		langPackString.ManyValue = strings.ManyValue
		langPackString.OtherValue = strings.OtherValue

		diffStrings = append(diffStrings, langPackString)
	}

	for _, stringPluralizeds := range langPacks.StringPluralizeds {
		ss := mtproto.NewTLLangPackString(nil)
		langPackString := ss.To_LangPackString()
		langPackString.Key = stringPluralizeds.Key
		langPackString.Value = stringPluralizeds.Value
		langPackString.ZeroValue = stringPluralizeds.ZeroValue
		langPackString.OneValue = stringPluralizeds.OneValue
		langPackString.TwoValue = stringPluralizeds.TwoValue
		langPackString.FewValue = stringPluralizeds.FewValue
		langPackString.ManyValue = stringPluralizeds.ManyValue
		langPackString.OtherValue = stringPluralizeds.OtherValue

		diffStrings = append(diffStrings, langPackString)
	}

	diff.SetStrings(diffStrings)
	log.Infof("LangpackGetLangPack %v, request: %v ok!!! reply", md, request)
	return diff.To_LangPackDifference(), nil
}
