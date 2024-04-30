/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getUserInfo_handler.go
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

//  help.getUserInfo#38a08d3 user_id:InputUser = help.UserInfo;
//
func (s *HelpServiceImpl) HelpGetUserInfo(ctx context.Context, request *mtproto.TLHelpGetUserInfo) (*mtproto.Help_UserInfo, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetUserInfo %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl HelpGetUserInfo logic

	return nil, fmt.Errorf("%s", "Not impl HelpGetUserInfo")
}
