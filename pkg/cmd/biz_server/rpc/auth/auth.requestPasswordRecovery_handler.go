/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.requestPasswordRecovery_handler.go
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

//  auth.requestPasswordRecovery#d897bc66 = auth.PasswordRecovery;
//
func (s *AuthServiceImpl) AuthRequestPasswordRecovery(ctx context.Context, request *mtproto.TLAuthRequestPasswordRecovery) (*mtproto.Auth_PasswordRecovery, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("AuthRequestPasswordRecovery %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl AuthRequestPasswordRecovery logic

    return nil, fmt.Errorf("%s", "Not impl AuthRequestPasswordRecovery")
}
