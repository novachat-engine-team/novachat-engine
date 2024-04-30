/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : bots.answerWebhookJSONQuery_handler.go
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

//  bots.answerWebhookJSONQuery#e6213f4d query_id:long data:DataJSON = Bool;
//
func (s *BotsServiceImpl) BotsAnswerWebhookJSONQuery(ctx context.Context, request *mtproto.TLBotsAnswerWebhookJSONQuery) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("BotsAnswerWebhookJSONQuery %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl BotsAnswerWebhookJSONQuery logic

	return nil, fmt.Errorf("%s", "Not impl BotsAnswerWebhookJSONQuery")
}
