/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getTermsOfServiceUpdate_handler.go
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

//  help.getTermsOfServiceUpdate#2ca51fd1 = help.TermsOfServiceUpdate;
//
func (s *HelpServiceImpl) HelpGetTermsOfServiceUpdate(ctx context.Context, request *mtproto.TLHelpGetTermsOfServiceUpdate) (*mtproto.Help_TermsOfServiceUpdate, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetTermsOfServiceUpdate %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	termsOfServiceUpdate := mtproto.NewTLHelpTermsOfServiceUpdateEmpty(&mtproto.Help_TermsOfServiceUpdate{
		Expires: int32(time.Now().Unix() + 3600),
	})
	return termsOfServiceUpdate.To_Help_TermsOfServiceUpdate(), nil
}
