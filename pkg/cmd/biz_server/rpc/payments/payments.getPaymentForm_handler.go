/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : payments.getPaymentForm_handler.go
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

//  payments.getPaymentForm#99f09745 msg_id:int = payments.PaymentForm;
//
func (s *PaymentsServiceImpl) PaymentsGetPaymentForm(ctx context.Context, request *mtproto.TLPaymentsGetPaymentForm) (*mtproto.Payments_PaymentForm, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PaymentsGetPaymentForm %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	return mtproto.NewTLPaymentsPaymentForm(nil).To_Payments_PaymentForm(), nil
}
