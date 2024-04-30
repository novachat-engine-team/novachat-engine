/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getCdnConfig_handler.go
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

//  help.getCdnConfig#52029342 = CdnConfig;
//
func (s *HelpServiceImpl) HelpGetCdnConfig(ctx context.Context, request *mtproto.TLHelpGetCdnConfig) (*mtproto.CdnConfig, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetCdnConfig %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl HelpGetCdnConfig logic

	return nil, fmt.Errorf("%s", "Not impl HelpGetCdnConfig")
}
