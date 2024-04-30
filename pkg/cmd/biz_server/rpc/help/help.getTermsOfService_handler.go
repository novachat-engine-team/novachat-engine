/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getTermsOfService_handler.go
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

//  help.getTermsOfService#350170f3 = help.TermsOfService;
//
func (s *HelpServiceImpl) HelpGetTermsOfService(ctx context.Context, request *mtproto.TLHelpGetTermsOfService) (*mtproto.Help_TermsOfService, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetTermsOfService %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	return mtproto.NewTLHelpTermsOfService(nil).To_Help_TermsOfService(), nil
}
