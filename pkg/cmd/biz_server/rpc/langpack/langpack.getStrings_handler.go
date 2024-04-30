/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : langpack.getStrings_handler.go
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

//  langpack.getStrings#efea3803 lang_pack:string lang_code:string keys:Vector<string> = Vector<LangPackString>;
//  langpack.getStrings#2e1ee318 lang_code:string keys:Vector<string> = Vector<LangPackString>;
//
func (s *LangpackServiceImpl) LangpackGetStrings(ctx context.Context, request *mtproto.TLLangpackGetStrings) (*mtproto.Vector_LangPackString, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("LangpackGetStrings %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	langPacks := lang.GetLangPacks(request.GetLangCode())
	if langPacks == nil {
		log.Errorf("LangpackGetStrings request:%v Invalid LangCode", request)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_LANG_PACK_INVALID_LANGCODE)
	}

	reply := make([]*mtproto.LangPackString, 0, len(request.Keys))
	for idx := range request.Keys {
		r := mtproto.NewTLLangPackString(nil)
		r.SetKey(request.Keys[idx])
		r.SetValue(langPacks.GetString(request.Keys[idx]))
		reply = append(reply, r.To_LangPackString())
	}

	log.Infof("LangpackGetStrings %v, request: %v ok!!! reply", md, request)
	return &mtproto.Vector_LangPackString{Datas: reply}, nil
}
