/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getRecentMeUrls_handler.go
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

//  help.getRecentMeUrls#3dc0f114 referer:string = help.RecentMeUrls;
//
func (s *HelpServiceImpl) HelpGetRecentMeUrls(ctx context.Context, request *mtproto.TLHelpGetRecentMeUrls) (*mtproto.Help_RecentMeUrls, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetRecentMeUrls %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl HelpGetRecentMeUrls logic

	return nil, fmt.Errorf("%s", "Not impl HelpGetRecentMeUrls")
}
