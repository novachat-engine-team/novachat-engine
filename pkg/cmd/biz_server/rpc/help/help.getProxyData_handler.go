/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getProxyData_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  help.getProxyData#3d7758e1 = help.ProxyData;
//
func (s *HelpServiceImpl) HelpGetProxyData(ctx context.Context, request *mtproto.TLHelpGetProxyData) (*mtproto.Help_ProxyData, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetProxyData %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	log.Infof("HelpGetProxyData %v, request: %v ok!!! reply", metadata.RpcMetaDataDebug(md), request)
	return nil, fmt.Errorf("%s", "Not impl HelpGetCdnConfig")
}
