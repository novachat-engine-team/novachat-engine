/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/2 13:57
 * @File : banned_right.go
 */

package banned_right

import "novachat_engine/mtproto"

func ChannelBannedRightToChatBannedRight(rights *mtproto.ChannelBannedRights) *mtproto.ChatBannedRights {
	return mtproto.NewTLChatBannedRights(&mtproto.ChatBannedRights{
		ViewMessages: rights.ViewMessages,
		SendMessages: rights.SendMessages,
		SendMedia:    rights.SendMedia,
		SendStickers: rights.SendStickers,
		SendGifs:     rights.SendGifs,
		SendGames:    rights.SendGames,
		SendInline:   rights.SendInline,
		EmbedLinks:   rights.EmbedLinks,
		SendPolls:    true,
		ChangeInfo:   true,
		InviteUsers:  true,
		PinMessages:  true,
		UntilDate:    rights.UntilDate,
	}).To_ChatBannedRights()
}

func ChatBannedRightToChannelBannedRight(rights *mtproto.ChatBannedRights) *mtproto.ChannelBannedRights {
	if rights == nil {
		return nil
	}
	return mtproto.NewTLChannelBannedRights(&mtproto.ChannelBannedRights{
		ViewMessages: rights.ViewMessages,
		SendMessages: rights.SendMessages,
		SendMedia:    rights.SendMedia,
		SendStickers: rights.SendStickers,
		SendGifs:     rights.SendGifs,
		SendGames:    rights.SendGames,
		SendInline:   rights.SendInline,
		EmbedLinks:   rights.EmbedLinks,
		UntilDate:    rights.UntilDate,
	}).To_ChannelBannedRights()
}

func ChannelAdminRightToChatAdminRight(rights *mtproto.ChannelAdminRights) *mtproto.ChatAdminRights {
	return mtproto.NewTLChatAdminRights(&mtproto.ChatAdminRights{
		ChangeInfo:     rights.ChangeInfo,
		PostMessages:   rights.PostMessages,
		EditMessages:   rights.EditMessages,
		DeleteMessages: rights.DeleteMessages,
		BanUsers:       rights.BanUsers,
		InviteUsers:    rights.InviteUsers,
		PinMessages:    rights.PinMessages,
		AddAdmins:      rights.AddAdmins,
		Anonymous:      false,
		ManageCall:     false,
	}).To_ChatAdminRights()
}

func ChatAdminRightToChannelAdminRight(rights *mtproto.ChatAdminRights) *mtproto.ChannelAdminRights {
	if rights == nil {
		return nil
	}

	return mtproto.NewTLChannelAdminRights(&mtproto.ChannelAdminRights{
		ChangeInfo:     rights.ChangeInfo,
		PostMessages:   rights.PostMessages,
		EditMessages:   rights.EditMessages,
		DeleteMessages: rights.DeleteMessages,
		BanUsers:       rights.BanUsers,
		InviteUsers:    rights.InviteUsers,
		InviteLink:     true,
		PinMessages:    rights.PinMessages,
		AddAdmins:      rights.AddAdmins,
	}).To_ChannelAdminRights()
}

func ChannelAdminRightToChannelBannedRight(adminRight *mtproto.ChannelAdminRights) *mtproto.ChannelBannedRights {
	channelBannedRight := mtproto.NewTLChannelBannedRights(&mtproto.ChannelBannedRights{
		ViewMessages: false,
		SendMessages: false,
		SendMedia:    false,
		SendStickers: false,
		SendGifs:     false,
		SendGames:    false,
		SendInline:   false,
		EmbedLinks:   false,
		UntilDate:    0,
	}).To_ChannelBannedRights()

	return channelBannedRight
}

func ChatAdminRightToChatBannedRight(adminRight *mtproto.ChatAdminRights) *mtproto.ChatBannedRights {
	chatBannedRight := mtproto.NewTLChatBannedRights(&mtproto.ChatBannedRights{
		ViewMessages: false,
		SendMessages: false,
		SendMedia:    false,
		SendStickers: false,
		SendGifs:     false,
		SendGames:    false,
		SendInline:   false,
		EmbedLinks:   false,
		SendPolls:    false,
		ChangeInfo:   adminRight.ChangeInfo,
		InviteUsers:  adminRight.InviteUsers,
		PinMessages:  adminRight.PinMessages,
		//BanWisper:            adminRight.BanUsers,
		UntilDate: 0,
	}).To_ChatBannedRights()

	return chatBannedRight
}
