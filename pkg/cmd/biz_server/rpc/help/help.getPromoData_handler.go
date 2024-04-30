/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getPromoData_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"time"
)

//  help.getPromoData#c0977421 = help.PromoData;
//
func (s *HelpServiceImpl) HelpGetPromoData(ctx context.Context, request *mtproto.TLHelpGetPromoData) (*mtproto.Help_PromoData, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("HelpGetPromoData %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	return mtproto.NewTLHelpPromoDataEmpty(&mtproto.Help_PromoData{
		Expires: int32(time.Now().Unix() + 3600),
	}).To_Help_PromoData(), nil
}
