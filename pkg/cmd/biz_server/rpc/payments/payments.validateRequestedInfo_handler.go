/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : payments.validateRequestedInfo_handler.go
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

//  payments.validateRequestedInfo#770a8e74 flags:# save:flags.0?true msg_id:int info:PaymentRequestedInfo = payments.ValidatedRequestedInfo;
//
func (s *PaymentsServiceImpl) PaymentsValidateRequestedInfo(ctx context.Context, request *mtproto.TLPaymentsValidateRequestedInfo) (*mtproto.Payments_ValidatedRequestedInfo, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PaymentsValidateRequestedInfo %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PaymentsValidateRequestedInfo logic

	return nil, fmt.Errorf("%s", "Not impl PaymentsValidateRequestedInfo")
}
