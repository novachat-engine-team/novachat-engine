/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getNearestDc_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/app/dc"
	"novachat_engine/service/app/help"
)

//  help.getNearestDc#1fb33026 = NearestDc;
//
func (s *HelpServiceImpl) HelpGetNearestDc(ctx context.Context, request *mtproto.TLHelpGetNearestDc) (*mtproto.NearestDc, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetNearestDc#1fb33026 %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	reply := dc.GetNearestDc(help.GetConfig().ThisDc)

	log.Infof("HelpGetNearestDc#1fb33026 %v, request: %v reply:%v", md, request, reply)
	return reply, nil
}
