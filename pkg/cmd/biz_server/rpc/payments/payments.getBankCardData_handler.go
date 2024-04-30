/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : payments.getBankCardData_handler.go
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

//  payments.getBankCardData#2e79d779 number:string = payments.BankCardData;
//
func (s *PaymentsServiceImpl) PaymentsGetBankCardData(ctx context.Context, request *mtproto.TLPaymentsGetBankCardData) (*mtproto.Payments_BankCardData, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PaymentsGetBankCardData %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PaymentsGetBankCardData logic

	return nil, fmt.Errorf("%s", "Not impl PaymentsGetBankCardData")
}
