/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : payments.getPaymentReceipt_handler.go
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

//  payments.getPaymentReceipt#a092a980 msg_id:int = payments.PaymentReceipt;
//
func (s *PaymentsServiceImpl) PaymentsGetPaymentReceipt(ctx context.Context, request *mtproto.TLPaymentsGetPaymentReceipt) (*mtproto.Payments_PaymentReceipt, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("PaymentsGetPaymentReceipt %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    // Impl PaymentsGetPaymentReceipt logic

    return nil, fmt.Errorf("%s", "Not impl PaymentsGetPaymentReceipt")
}
