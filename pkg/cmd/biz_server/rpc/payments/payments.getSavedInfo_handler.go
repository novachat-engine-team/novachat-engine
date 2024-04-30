/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : payments.getSavedInfo_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  payments.getSavedInfo#227d824b = payments.SavedInfo;
//
func (s *PaymentsServiceImpl) PaymentsGetSavedInfo(ctx context.Context, request *mtproto.TLPaymentsGetSavedInfo) (*mtproto.Payments_SavedInfo, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PaymentsGetSavedInfo %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PaymentsGetSavedInfo logic

	return nil, fmt.Errorf("%s", "Not impl PaymentsGetSavedInfo")
}
