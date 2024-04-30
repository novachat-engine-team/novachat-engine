/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getInviteText_handler.go
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

//  help.getInviteText#4d392343 = help.InviteText;
//
func (s *HelpServiceImpl) HelpGetInviteText(ctx context.Context, request *mtproto.TLHelpGetInviteText) (*mtproto.Help_InviteText, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("HelpGetInviteText %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	text := mtproto.NewTLHelpInviteText(nil)
	text.SetMessage(s.conf.InviteText)
	return text.To_Help_InviteText(), nil
}
