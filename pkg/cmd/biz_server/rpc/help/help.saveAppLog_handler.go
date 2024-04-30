/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.saveAppLog_handler.go
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

//  help.saveAppLog#6f02f748 events:Vector<InputAppEvent> = Bool;
//
func (s *HelpServiceImpl) HelpSaveAppLog(ctx context.Context, request *mtproto.TLHelpSaveAppLog) (*mtproto.Bool, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Infof("HelpSaveAppLog %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    log.Warnf("HelpSaveAppLog events:%v", request.Events)

    return mtproto.ToMTBool(true), nil
}
