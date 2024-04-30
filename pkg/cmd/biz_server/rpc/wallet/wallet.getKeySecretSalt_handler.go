/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2022/09/22 10:25
 * @File : wallet.getKeySecretSalt_handler.go
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

//  wallet.getKeySecretSalt#b57f346 revoke:Bool = wallet.KeySecretSalt;
//
func (s *WalletServiceImpl) WalletGetKeySecretSalt(ctx context.Context, request *mtproto.TLWalletGetKeySecretSalt) (*mtproto.Wallet_KeySecretSalt, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("WalletGetKeySecretSalt %v, request: %v", md, request)

	// Impl WalletGetKeySecretSalt logic

	return nil, fmt.Errorf("%s", "Not impl WalletGetKeySecretSalt")
}
