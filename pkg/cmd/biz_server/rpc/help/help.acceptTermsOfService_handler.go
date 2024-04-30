/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.acceptTermsOfService_handler.go
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

//  help.acceptTermsOfService#ee72f79a id:DataJSON = Bool;
//
func (s *HelpServiceImpl) HelpAcceptTermsOfService(ctx context.Context, request *mtproto.TLHelpAcceptTermsOfService) (*mtproto.Bool, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("HelpAcceptTermsOfService %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    log.Infof("HelpAcceptTermsOfService %v, request: %v ok!!! reply", metadata.RpcMetaDataDebug(md), request)
    return mtproto.ToMTBool(true), nil
}
