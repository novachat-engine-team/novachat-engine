/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.blockFromReplies_handler.go
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

//  contacts.blockFromReplies#29a8962c flags:# delete_message:flags.0?true delete_history:flags.1?true report_spam:flags.2?true msg_id:int = Updates;
//
func (s *ContactsServiceImpl) ContactsBlockFromReplies(ctx context.Context, request *mtproto.TLContactsBlockFromReplies) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsBlockFromReplies %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	// Impl ContactsBlockFromReplies logic

	return nil, fmt.Errorf("%s", "Not impl ContactsBlockFromReplies")
}
