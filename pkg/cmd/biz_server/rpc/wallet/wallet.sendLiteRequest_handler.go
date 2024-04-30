/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/09/22 10:25
 * @File : wallet.sendLiteRequest_handler.go
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

//  wallet.sendLiteRequest#e2c9d33e body:bytes = wallet.LiteResponse;
//
func (s *WalletServiceImpl) WalletSendLiteRequest(ctx context.Context, request *mtproto.TLWalletSendLiteRequest) (*mtproto.Wallet_LiteResponse, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("WalletSendLiteRequest %v, request: %v", md, request)

    // Impl WalletSendLiteRequest logic

    return nil, fmt.Errorf("%s", "Not impl WalletSendLiteRequest")
}
