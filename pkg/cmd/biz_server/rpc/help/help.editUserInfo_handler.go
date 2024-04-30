/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.editUserInfo_handler.go
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

//  help.editUserInfo#66b91b70 user_id:InputUser message:string entities:Vector<MessageEntity> = help.UserInfo;
//
func (s *HelpServiceImpl) HelpEditUserInfo(ctx context.Context, request *mtproto.TLHelpEditUserInfo) (*mtproto.Help_UserInfo, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpEditUserInfo %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl HelpEditUserInfo logic

	return nil, fmt.Errorf("%s", "Not impl HelpEditUserInfo")
}
