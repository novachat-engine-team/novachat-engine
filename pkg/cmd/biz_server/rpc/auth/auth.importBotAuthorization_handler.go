/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.importBotAuthorization_handler.go
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

//  auth.importBotAuthorization#67a3ff2c flags:int api_id:int api_hash:string bot_auth_token:string = auth.Authorization;
//
func (s *AuthServiceImpl) AuthImportBotAuthorization(ctx context.Context, request *mtproto.TLAuthImportBotAuthorization) (*mtproto.Auth_Authorization, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthImportBotAuthorization %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	authorization := mtproto.NewTLAuthAuthorization(&mtproto.Auth_Authorization{
		User: mtproto.NewTLUserEmpty(nil).To_User(),
	})

	//log.Debugf("AuthImportBotAuthorization - metadata- reply: %s", log.JsonDebugData(authorization))
	return authorization.To_Auth_Authorization(), nil
}
