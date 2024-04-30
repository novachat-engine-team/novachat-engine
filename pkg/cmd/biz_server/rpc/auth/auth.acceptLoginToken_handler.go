/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.acceptLoginToken_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	sessionService "novachat_engine/pkg/cmd/session/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"time"
)

//  auth.acceptLoginToken#e894ad4d token:bytes = Authorization;
//
func (s *AuthServiceImpl) AuthAcceptLoginToken(ctx context.Context, request *mtproto.TLAuthAcceptLoginToken) (*mtproto.Authorization, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthAcceptLoginToken %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AuthAcceptLoginToken logic
	qrLoginData, err := s.qrLoginCore.AcceptLoginToken(request.Token, md.UserId)
	if err != nil {
		return nil, err
	}

	updates := mtproto.NewTLUpdateShort(nil)
	updates.SetUpdate(mtproto.NewTLUpdateLoginToken(nil).To_Update())
	any, _ := types.MarshalAny(updates)
	conn := sessionService.GetSessionClientByKey(fmt.Sprintf("%d", qrLoginData.AuthKeyId))
	conn.PushUpdates(ctx, &sessionService.PushUpdatesEvent{
		AuthKeyId:  qrLoginData.AuthKeyId,
		SessionId:  []int64{qrLoginData.SessionId},
		Updates:    any,
		ServerPeer: "",
	})

	return mtproto.NewTLAuthorization(&mtproto.Authorization{
		Hash:            0,
		DeviceModel:     qrLoginData.DeviceModel,
		Platform:        "",
		SystemVersion:   qrLoginData.SystemVersion,
		ApiId:           qrLoginData.ApiId,
		AppName:         "",
		AppVersion:      qrLoginData.AppVersion,
		DateCreated:     qrLoginData.Date,
		DateActive:      int32(time.Now().Unix()),
		Ip:              qrLoginData.Ip,
		Country:         "",
		Region:          "",
		Current:         false,
		OfficialApp:     true,
		PasswordPending: false,
	}).To_Authorization(), nil
}
