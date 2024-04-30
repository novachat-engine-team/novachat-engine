/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : phone.toggleGroupCallSettings_handler.go
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

//  phone.toggleGroupCallSettings#74bbb43d flags:# call:InputGroupCall join_muted:flags.0?Bool = Updates;
//
func (s *PhoneServiceImpl) PhoneToggleGroupCallSettings(ctx context.Context, request *mtproto.TLPhoneToggleGroupCallSettings) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("PhoneToggleGroupCallSettings %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl PhoneToggleGroupCallSettings logic

	return nil, fmt.Errorf("%s", "Not impl PhoneToggleGroupCallSettings")
}
