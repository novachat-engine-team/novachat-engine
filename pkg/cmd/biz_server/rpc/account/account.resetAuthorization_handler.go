/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.resetAuthorization_handler.go
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

//  account.resetAuthorization#df77f3bc hash:long = Bool;
//
func (s *AccountServiceImpl) AccountResetAuthorization(ctx context.Context, request *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AccountResetAuthorization %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//authorizations, err := s.authCore.GetAuthorizations(md.UserId)
	//if err != nil {
	//	log.Errorf("AccountResetAuthorization - request: %v AuthResetAuthorization error:%s", request, err.Error())
	//	return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	//}
	//
	//localAuthorization := func() *auth_pb.AuthKeyInfo {
	//	for _, v := range authorizations {
	//		if v.AuthKeyId == md.AuthKeyId {
	//			return v
	//		}
	//	}
	//	return nil
	//}()
	//if localAuthorization == nil {
	//	log.Errorf("AccountResetAuthorization - request: %v AuthResetAuthorization localAuthorization nil error:", request)
	//	return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	//}
	//
	//now := time.Now().Unix()
	//if int32(now)-localAuthorization.Date < 24*3600 {
	//	return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_NOT_ACCEPTABLE_FRESH_RESET_AUTHORISATION_FORBIDDEN)
	//}
	//
	//requestAuthorization := func() *auth_pb.AuthKeyInfo {
	//	for _, v := range authorizations {
	//		if v.Info != nil && v.Info.Hash == request.Hash {
	//			return v
	//		}
	//	}
	//	return nil
	//}()
	//if requestAuthorization == nil {
	//	return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST)
	//}
	//
	//if localAuthorization.Date < requestAuthorization.Date {
	//	return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST, "can't reset old session")
	//}
	//
	//err = s.authCore.ResetAuthorization([]*auth_pb.AuthKeyInfo{requestAuthorization})
	//if err != nil {
	//	log.Warnf("AccountResetAuthorization - request: %v AuthResetAuthorization authDevice empty ", request)
	//	return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	//}
	//
	//TODO:Coder Updates
	//authKeyIdList = append(authKeyIdList, v.AuthKeyId)
	//	updateList.PushUpdatesList = append(updateList.PushUpdatesList, &sync_pb.PushUpdates{
	//		UserId:               userId,
	//		AuthKeyId:            v.AuthKeyId,
	//		ExcludeAuthKeyIdList: nil,
	//		Updates:              mtproto.NewTLUpdatesTooLong(nil).To_Updates(),
	//	})
	return mtproto.ToMTBool(true), nil
}
