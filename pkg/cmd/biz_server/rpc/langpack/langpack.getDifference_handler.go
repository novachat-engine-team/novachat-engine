/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : langpack.getDifference_handler.go
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

//  langpack.getDifference#cd984aa5 lang_pack:string lang_code:string from_version:int = LangPackDifference;
//  langpack.getDifference#b2e4d7d from_version:int = LangPackDifference;
//
func (s *LangpackServiceImpl) LangpackGetDifference(ctx context.Context, request *mtproto.TLLangpackGetDifference) (*mtproto.LangPackDifference, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("LangpackGetDifference %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	langPacks := lang.GetLangPacks(request.GetLangCode())
	if langPacks == nil {
		log.Errorf("LangpackGetDifference request:%v Invalid LangCode", request)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_LANG_PACK_INVALID_LANGCODE)
	}

	langPackVersion := langPacks.Version
	diff := mtproto.NewTLLangPackDifference(nil)
	diff.SetLangCode(request.LangCode)
	diff.SetVersion(langPackVersion)
	diff.SetFromVersion(request.GetFromVersion())

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

	log.Debugf("langpack.getDifference#b2e4d7d - reply: %s", log.JsonDebugData(diff))
	return diff.To_LangPackDifference(), nil
}
