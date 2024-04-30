/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.sendConfirmPhoneCode_handler.go
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

//  account.sendConfirmPhoneCode#1b3faa88 hash:string settings:CodeSettings = auth.SentCode;
//  account.sendConfirmPhoneCode#1516d7bd flags:# allow_flashcall:flags.0?true hash:string current_number:flags.0?Bool = auth.SentCode;
//
func (s *AccountServiceImpl) AccountSendConfirmPhoneCode(ctx context.Context, request *mtproto.TLAccountSendConfirmPhoneCode) (*mtproto.Auth_SentCode, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountSendConfirmPhoneCode %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl AccountSendConfirmPhoneCode logic

	return nil, fmt.Errorf("%s", "Not impl AccountSendConfirmPhoneCode")
}
