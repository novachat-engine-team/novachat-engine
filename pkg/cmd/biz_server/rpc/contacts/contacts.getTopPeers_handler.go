/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.getTopPeers_handler.go
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

//  contacts.getTopPeers#d4982db5 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true phone_calls:flags.3?true forward_users:flags.4?true forward_chats:flags.5?true groups:flags.10?true channels:flags.15?true offset:int limit:int hash:int = contacts.TopPeers;
//
func (s *ContactsServiceImpl) ContactsGetTopPeers(ctx context.Context, request *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
    md := metadata.RpcMetaDataFromContext(ctx)
    log.Debugf("ContactsGetTopPeers %v, request: %v", metadata.RpcMetaDataDebug(md), request)

    //TODO:(Coder)
    topPeers := &mtproto.TLContactsTopPeers{Data2: &mtproto.Contacts_TopPeers{
        Categories: []*mtproto.TopPeerCategoryPeers{},
        Chats:      []*mtproto.Chat{},
        Users:      []*mtproto.User{},
    }}

    if request.Correspondents { // most frequently with users

    }

    //request.BotsPm
    //request.BotsInline
    if request.PhoneCalls {

    }

    if request.Groups {

    }
    if request.Channels {

    }

    log.Debugf("ContactsGetTopPeers %v, request: %v ok!!!!!!!!", md, request)
    return topPeers.To_Contacts_TopPeers(), nil
}
