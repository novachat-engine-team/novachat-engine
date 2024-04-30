/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.getDialogFilters_handler.go
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

//DialogFilter
//Dialog filter (folders)
//
//Telegram allows placing chats into folders, based on their type, mute status, or other custom criteria, thanks to folder blacklists and whitelists.
//
//Name	Type	Description
//flags	#	Flags, see TL conditional fields
//contacts	flags.0?true	Whether to include all contacts in this folder
//non_contacts	flags.1?true	Whether to include all non-contacts in this folder
//groups	flags.2?true	Whether to include all groups in this folder
//broadcasts	flags.3?true	Whether to include all channels in this folder
//bots	flags.4?true	Whether to include all bots in this folder
//exclude_muted	flags.11?true	Whether to exclude muted chats from this folder
//exclude_read	flags.12?true	Whether to exclude read chats from this folder
//exclude_archived	flags.13?true	Whether to exclude archived chats from this folder
//id	int	Folder ID
//title	string	Folder name
//emoticon	flags.25?string	Folder emoticon
//pinned_peers	Vector<InputPeer>	Pinned chats, folders can have unlimited pinned chats
//include_peers	Vector<InputPeer>	Include the following chats in this folder
//exclude_peers	Vector<InputPeer>	Exclude the following chats from this folder
//Type
//DialogFilter
//
//Related pages
//Folders
//Telegram allows placing chats into folders, based on their type, mute status, or other custom criteria, thanks to folder blacklists and whitelists.

//  messages.getDialogFilters#f19ed96d = Vector<DialogFilter>;
//
func (s *MessagesServiceImpl) MessagesGetDialogFilters(ctx context.Context, request *mtproto.TLMessagesGetDialogFilters) (*mtproto.Vector_DialogFilter, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesGetDialogFilters %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	return &mtproto.Vector_DialogFilter{
		Datas: nil,
	}, nil
}
