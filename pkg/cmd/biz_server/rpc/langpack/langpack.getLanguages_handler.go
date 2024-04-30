/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : langpack.getLanguages_handler.go
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

//  langpack.getLanguages#42c6978f lang_pack:string = Vector<LangPackLanguage>;
//  langpack.getLanguages#800fd57d = Vector<LangPackLanguage>;
//
func (s *LangpackServiceImpl) LangpackGetLanguages(ctx context.Context, request *mtproto.TLLangpackGetLanguages) (*mtproto.Vector_LangPackLanguage, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("LangpackGetLanguages#800fd57d %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	reply := lang.GetLanguages()

	log.Infof("LangpackGetLanguages#800fd57d %v, request: %v ok!!!", md, request)
	return reply, nil
}
