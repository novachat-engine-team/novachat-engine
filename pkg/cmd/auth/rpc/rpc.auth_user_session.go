/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package rpc

import (
	"context"
	authService "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/log"
)

func (m *AuthImpl) ReqUserSession(ctx context.Context, request *authService.UserSession) (*authService.UserSessionInfo, error) {
	authList, err := m.accountAuthCore.UserSession(request.UserIdList)
	if err != nil {
		log.Errorf("ReqUserSession device:%+v error:%s", request, err.Error())
		return nil, err
	}

	ret := make([]*authService.UserSessionInfo_SessionInfo, 0, len(authList))
	for _, v := range authList {
		if v.AuthInfo == nil {
			continue
		}

		info := &authService.UserSessionInfo_SessionInfo{
			UserId:    v.AuthInfo.UserId,
			AuthKeyId: v.AuthInfo.AuthKeyId,
			SessionId: v.SessionIdList,
			Device:    nil,
		}

		if v.AuthInfo.Devices != nil {
			info.Device = &authService.Device{
				TokenType:    v.AuthInfo.Devices.TokenType,
				AppSandBox:   v.AuthInfo.Devices.AppSandBox,
				Secret:       v.AuthInfo.Devices.Secret,
				NoMuted:      v.AuthInfo.Devices.NoMuted,
				DeviceModel:  v.AuthInfo.Devices.DeviceModel,
				LangPack:     v.AuthInfo.Devices.LangPack,
				OtherUidList: v.AuthInfo.Devices.OtherUidList,
			}
		}
		ret = append(ret, info)
	}

	log.Debugf("ReqUserSession userList:%+v ret:%+v authList:%+v", request.UserIdList, ret, authList)
	return &authService.UserSessionInfo{
		SessionInfoList: ret,
	}, nil
}
