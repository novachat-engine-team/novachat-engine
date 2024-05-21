/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.exportLoginToken_handler.go
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

//  auth.exportLoginToken#b1b41517 api_id:int api_hash:string except_ids:Vector<int> = auth.LoginToken;
//
func (s *AuthServiceImpl) AuthExportLoginToken(ctx context.Context, request *mtproto.TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthExportLoginToken %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	token, date, userId, err := s.qrLoginCore.ExportLoginToken(
		md.AuthKeyId,
		md.SessionId,
		md,
		request.ApiId,
		int32(time.Now().Unix()))
	if err != nil {
		log.Errorf("AuthExportLoginToken ExportLoginToken  error:%s", err.Error())
		return nil, err
	}

	if userId == 0 {
		authLoginToken := mtproto.NewTLAuthLoginToken(nil)
		authLoginToken.SetToken(token)
		authLoginToken.SetExpires(date)
		return authLoginToken.To_Auth_LoginToken(), nil
	} else {
		log.Debugf("AuthExportLoginToken ExportLoginToken GetUser userId:%d ", userId)
		user, err1 := s.accountUsersCore.GetUser(userId, userId, md.Layer)
		if err1 != nil {
			log.Errorf("AuthExportLoginToken ExportLoginToken GetUser userId:%d error:%s", userId, err.Error())
			return nil, err
		}

		return mtproto.NewTLAuthLoginTokenSuccess(&mtproto.Auth_LoginToken{
			Authorization: mtproto.NewTLAuthAuthorization(&mtproto.Auth_Authorization{
				User: user,
			}).To_Auth_Authorization(),
		}).To_Auth_LoginToken(), nil
	}
}
