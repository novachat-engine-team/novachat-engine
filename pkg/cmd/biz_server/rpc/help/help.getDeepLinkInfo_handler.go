/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getDeepLinkInfo_handler.go
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

//  help.getDeepLinkInfo#3fedc75f path:string = help.DeepLinkInfo;
//
func (s *HelpServiceImpl) HelpGetDeepLinkInfo(ctx context.Context, request *mtproto.TLHelpGetDeepLinkInfo) (*mtproto.Help_DeepLinkInfo, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetDeepLinkInfo %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl HelpGetDeepLinkInfo logic

	return nil, fmt.Errorf("%s", "Not impl HelpGetDeepLinkInfo")
}
