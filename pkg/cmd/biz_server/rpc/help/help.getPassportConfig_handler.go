/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getPassportConfig_handler.go
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

//  help.getPassportConfig#c661ad08 hash:int = help.PassportConfig;
//
func (s *HelpServiceImpl) HelpGetPassportConfig(ctx context.Context, request *mtproto.TLHelpGetPassportConfig) (*mtproto.Help_PassportConfig, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetPassportConfig %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl HelpGetPassportConfig logic

	return nil, fmt.Errorf("%s", "Not impl HelpGetPassportConfig")
}
