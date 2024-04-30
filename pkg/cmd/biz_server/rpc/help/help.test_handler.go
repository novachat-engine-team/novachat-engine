/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/09/27 18:00
 * @File : help.test_handler.go
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

//  help.test#c0e202f7 = Bool;
//
func (s *HelpServiceImpl) HelpTest(ctx context.Context, request *mtproto.TLHelpTest) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpTest %v, request: %v", md, request)

	// Impl HelpTest logic

	return mtproto.ToMTBool(true), nil
}
