/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.dismissSuggestion_handler.go
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

//  help.dismissSuggestion#77fa99f suggestion:string = Bool;
//
func (s *HelpServiceImpl) HelpDismissSuggestion(ctx context.Context, request *mtproto.TLHelpDismissSuggestion) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpDismissSuggestion %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl HelpDismissSuggestion logic
	log.Warnf("HelpDismissSuggestion userId:%d suggestion:%s", md.UserId, request.Suggestion)

	return mtproto.ToMTBool(true), nil
}
