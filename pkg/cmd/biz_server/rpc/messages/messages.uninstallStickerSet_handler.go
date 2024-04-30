/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.uninstallStickerSet_handler.go
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

//  messages.uninstallStickerSet#f96e55de stickerset:InputStickerSet = Bool;
//
func (s *MessagesServiceImpl) MessagesUninstallStickerSet(ctx context.Context, request *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesUninstallStickerSet %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var err error
	var ok bool
	stickerSetID := request.Stickerset.Id
	if stickerSetID == 0 {
		ok = false
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_STICKERSET_INVALID)
	} else {
		err = s.accountStickerSetCore.Uninstall(md.UserId, stickerSetID)
		if err != nil {
			log.Errorf("MessagesUninstallStickerSet %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		} else {
			log.Infof("MessagesUninstallStickerSet %v, request: %v ok!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request)
		}
		ok = true
	}

	return mtproto.ToMTBool(ok), err
}
