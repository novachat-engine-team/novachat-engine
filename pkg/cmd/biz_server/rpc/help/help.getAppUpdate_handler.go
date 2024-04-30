/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getAppUpdate_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  help.getAppUpdate#522d5a7d source:string = help.AppUpdate;
//  help.getAppUpdate#ae2de196 = help.AppUpdate;
//
func (s *HelpServiceImpl) HelpGetAppUpdate(ctx context.Context, request *mtproto.TLHelpGetAppUpdate) (*mtproto.Help_AppUpdate, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetAppUpdate %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl HelpGetAppUpdate logic

	return mtproto.NewTLHelpNoAppUpdate(nil).To_Help_AppUpdate(), nil
}
