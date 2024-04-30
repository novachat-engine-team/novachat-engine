/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : updates.getState_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
)

//  updates.getState#edd4882a = updates.State;
//
func (s *UpdatesServiceImpl) UpdatesGetState(ctx context.Context, request *mtproto.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("UpdatesGetState %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	updateState, err := s.accountMessageCore.GetState(md.UserId)
	if err != nil {
		log.Debugf("UpdatesGetState %v, request: %v", metadata.RpcMetaDataDebug(md), request)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST)
	}

	log.Infof("UpdatesGetState %v, request: %v reply:%v", md, request, updateState)
	return updateState, nil
}
