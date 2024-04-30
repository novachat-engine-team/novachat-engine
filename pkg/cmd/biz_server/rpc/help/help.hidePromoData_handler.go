/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.hidePromoData_handler.go
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

//  help.hidePromoData#1e251c95 peer:InputPeer = Bool;
//
func (s *HelpServiceImpl) HelpHidePromoData(ctx context.Context, request *mtproto.TLHelpHidePromoData) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpHidePromoData %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl HelpHidePromoData logic

	return nil, fmt.Errorf("%s", "Not impl HelpHidePromoData")
}
