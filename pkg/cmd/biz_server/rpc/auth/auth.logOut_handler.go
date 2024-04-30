/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.logOut_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  auth.logOut#5717da40 = Bool;
//
func (s *AuthServiceImpl) AuthLogOut(ctx context.Context, request *mtproto.TLAuthLogOut) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AuthLogOut %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	ok := mtproto.ToMTBool(true)
	resp, err := authClient.GetAuthClientByAuthKey(md.PermAuthKeyId).ReqUnbindAuthKeyUser(context.Background(), &authClient.UnbindAuthKeyUser{
		AuthKeyId:     md.AuthKeyId,
		PermAuthKeyId: md.PermAuthKeyId,
		UserId:        md.UserId,
	})
	if err != nil {
		log.Errorf("AuthLogOut - request: %v error:%s", request, err.Error())
	} else {
		types.UnmarshalAny(resp, ok)
		if mtproto.ToBool(ok) == false {
			log.Fatalf("AuthLogOut - request: %v ReqUnbindAuthKeyUser false", request)
		}
	}
	return ok, nil
}
