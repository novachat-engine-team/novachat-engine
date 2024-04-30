/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : langpack.getLanguage_handler.go
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
)

//  langpack.getLanguage#6a596502 lang_pack:string lang_code:string = LangPackLanguage;
//
func (s *LangpackServiceImpl) LangpackGetLanguage(ctx context.Context, request *mtproto.TLLangpackGetLanguage) (*mtproto.LangPackLanguage, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("LangpackGetLanguage#6a596502 %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	language, err := lang.GetLanguage(request.LangCode)
	log.Infof("LangpackGetLanguage#6a596502 %v, request: %v reply:%s", md, request, language)
	return language, err
}
