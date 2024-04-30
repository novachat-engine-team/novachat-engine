/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getThemes_handler.go
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

//  account.getThemes#285946f8 format:string hash:int = account.Themes;
//
func (s *AccountServiceImpl) AccountGetThemes(ctx context.Context, request *mtproto.TLAccountGetThemes) (*mtproto.Account_Themes, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("AccountGetThemes %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//TODO:(Coder) Impl AccountGetThemes logic
	themes := mtproto.NewTLAccountThemes(&mtproto.Account_Themes{
		Hash:   0,
		Themes: []*mtproto.Theme{},
	})
	return themes.To_Account_Themes(), nil
}
