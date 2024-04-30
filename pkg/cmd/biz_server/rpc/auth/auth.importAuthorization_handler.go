/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.importAuthorization_handler.go
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

//  auth.importAuthorization#e3ef9613 id:int bytes:bytes = auth.Authorization;
//
func (s *AuthServiceImpl) AuthImportAuthorization(ctx context.Context, request *mtproto.TLAuthImportAuthorization) (*mtproto.Auth_Authorization, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthImportAuthorization %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	authorization := mtproto.NewTLAuthAuthorization(&mtproto.Auth_Authorization{
		TmpSessions: request.GetId(),
		User:        mtproto.NewTLUserEmpty(nil).To_User(),
	})

	//log.Debugf("AuthImportAuthorization - metadata- reply: %s", log.JsonDebugData(authorization))
	return authorization.To_Auth_Authorization(), nil
}
