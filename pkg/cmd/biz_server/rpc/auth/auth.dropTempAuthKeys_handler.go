/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.dropTempAuthKeys_handler.go
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

//  auth.dropTempAuthKeys#8e48a188 except_auth_keys:Vector<long> = Bool;
//
func (s *AuthServiceImpl) AuthDropTempAuthKeys(ctx context.Context, request *mtproto.TLAuthDropTempAuthKeys) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthDropTempAuthKeys %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	panic("AuthDropTempAuthKeys")

	//ok, err := authClient.GetAuthClientByAuthKey(md.AuthKeyId).DropTempAuthKeys(context.Background(), &auth_pb.DropTempAuthKeys{
	//	AuthKeyId: md.AuthKeyId,
	//	UserId:    md.UserId,
	//	AuthKeys:  request.ExceptAuthKeys,
	//})
	//if err != nil {
	//	log.Fatalf("AuthDropTempAuthKeys - request: %v error:%s", request, err)
	//	return nil, err
	//}
	//
	//return ok, nil
}
