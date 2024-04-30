/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getAuthorizations_handler.go
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

//  account.getAuthorizations#e320c158 = account.Authorizations;
//
func (s *AccountServiceImpl) AccountGetAuthorizations(ctx context.Context, request *mtproto.TLAccountGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AccountGetAuthorizations %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//authorizationInfos, err := s.authCore.GetAuthorizations(md.UserId)
	//if err != nil {
	//	log.Errorf("AccountGetAuthorizations - request: %v error:%s", request, err.Error())
	//	return nil, err
	//}
	//
	//authorizations := make([]*mtproto.Authorization, 0, len(authorizationInfos))
	reply := mtproto.NewTLAccountAuthorizations(nil)

	//for _, info := range authorizationInfos {
	//	authorization := mtproto.NewTLAuthorization(&mtproto.Authorization{
	//		Hash:          info.Info.Hash,
	//		DeviceModel:   info.Info.DeviceModel,
	//		Platform:      info.Info.LangPack,
	//		SystemVersion: info.Info.SystemVersion,
	//		ApiId:         info.Info.ApiId,
	//		AppVersion:    info.Info.AppVersion,
	//		Ip:            info.Info.Ip,
	//		Country:       info.Info.LangCode,
	//		Region:        info.Info.SystemLangCode,
	//	}).To_Authorization()
	//
	//	if info.AuthKeyId == md.AuthKeyId {
	//		authorization.Current = true
	//	}
	//
	//	if info.Info.ApiId == constants.OfficialAPPID {
	//		authorization.OfficialApp = true
	//	}
	//
	//	authorizations = append(authorizations, authorization)
	//}
	//
	//reply.SetAuthorizations(authorizations)
	log.Infof("AccountGetAuthorizations - request: %v reply ok!!!!!!!!!!!!!!!!!!!!!!!!!!!", request)
	return reply.To_Account_Authorizations(), nil
}
