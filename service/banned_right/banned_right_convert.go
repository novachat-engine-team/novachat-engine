/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/3 18:31
 * @File : banned_right_convert.go
 */

package banned_right

import "novachat_engine/mtproto"

//
const (
	BannedRight_ViewMessages = 0
	BannedRight_SendMessages = 1
	BannedRight_SendMedia    = 2
	BannedRight_SendStickers = 3
	BannedRight_SendGifs     = 4
	BannedRight_SendGames    = 5
	BannedRight_SendInline   = 6
	BannedRight_EmbedLinks   = 7
	BannedRight_SendPolls    = 8
	BannedRight_ChangeInfo   = 10
	BannedRight_InviteUsers  = 15
	BannedRight_PinMessages  = 17
	BannedRight_BAN_WISPER   = 18
)

const (
	AdminRight_ChangeInfo     = 0
	AdminRight_PostMessages   = 1
	AdminRight_EditMessages   = 2
	AdminRight_DeleteMessages = 3
	AdminRight_BanUsers       = 4
	AdminRight_InviteUsers    = 5
	AdminRight_InviteLink     = 6 //ChannelAdminRight
	AdminRight_PinMessages    = 7
	AdminRight_AddAdmins      = 9
	AdminRight_Anonymous      = 10
	AdminRight_ManageCall     = 11
	AdminRight_OTHER          = 12
)

// client:
// open switch  => send false
// close switch => send true
// server:
// 0 => send false
// 1 => send true

func checkRight(r int32, rightType int32) bool {
	return r&(1<<rightType) != 0
}

func setRight(r *int32, flag bool, rightType int32) {
	if flag {
		*r = (*r) | (1 << rightType)
	} else {
		*r = (*r) & ^(1 << rightType)
	}
}

func FullAdminBannedRight() *mtproto.ChatAdminRights {
	right := mtproto.NewTLChatAdminRights(nil)
	right.SetChangeInfo(true)
	right.SetPostMessages(false)
	right.SetEditMessages(false)
	right.SetDeleteMessages(true)
	right.SetBanUsers(true)
	right.SetInviteUsers(true)
	right.SetPinMessages(true)
	right.SetAddAdmins(true)
	right.SetAnonymous(false)
	right.SetManageCall(true)

	return right.To_ChatAdminRights()
}

// AdminBannedRightRemove
// remove admin
func AdminBannedRightRemove(right *mtproto.ChatAdminRights) bool {
	return right.ChangeInfo == false &&
		right.PostMessages == false &&
		right.EditMessages == false &&
		right.DeleteMessages == false &&
		right.BanUsers == false &&
		right.InviteUsers == false &&
		right.PinMessages == false &&
		right.AddAdmins == false &&
		right.Anonymous == false &&
		right.ManageCall == false
}

func ChatBannedRightsUtilDateValid(rights *mtproto.ChatBannedRights, date int32) bool {
	if rights.UntilDate == 0 {
		return true
	}

	return rights.UntilDate >= date
}

func RightsToChatBannedRight(r int32, utilDate int32) *mtproto.ChatBannedRights {
	right := mtproto.NewTLChatBannedRights(nil)
	right.SetViewMessages(checkRight(r, BannedRight_ViewMessages))
	right.SetSendMessages(checkRight(r, BannedRight_SendMessages))
	right.SetSendMedia(checkRight(r, BannedRight_SendMedia))
	right.SetSendStickers(checkRight(r, BannedRight_SendStickers))
	right.SetSendGifs(checkRight(r, BannedRight_SendGifs))
	right.SetSendGames(checkRight(r, BannedRight_SendGames))
	right.SetSendInline(checkRight(r, BannedRight_SendInline))
	right.SetEmbedLinks(checkRight(r, BannedRight_EmbedLinks))
	right.SetSendPolls(checkRight(r, BannedRight_SendPolls))
	right.SetChangeInfo(checkRight(r, BannedRight_ChangeInfo))
	right.SetInviteUsers(checkRight(r, BannedRight_InviteUsers))
	right.SetPinMessages(checkRight(r, BannedRight_PinMessages))
	//right.SetBanWisper(checkRight(r, BannedRight_BAN_WISPER))

	right.SetUntilDate(utilDate)

	return right.To_ChatBannedRights()
}

/*
limit:
 banned_rights: { chatBannedRights
      flags: 165248 [INT],
      view_messages: [ SKIPPED BY BIT 0 IN FIELD flags ],
      send_messages: [ SKIPPED BY BIT 1 IN FIELD flags ],
      send_media: [ SKIPPED BY BIT 2 IN FIELD flags ],
      send_stickers: [ SKIPPED BY BIT 3 IN FIELD flags ],
      send_gifs: [ SKIPPED BY BIT 4 IN FIELD flags ],
      send_games: [ SKIPPED BY BIT 5 IN FIELD flags ],
      send_inline: [ SKIPPED BY BIT 6 IN FIELD flags ],
      embed_links: YES [ BY BIT 7 IN FIELD flags ],
      send_polls: YES [ BY BIT 8 IN FIELD flags ],
      change_info: YES [ BY BIT 10 IN FIELD flags ],
      invite_users: YES [ BY BIT 15 IN FIELD flags ],
      pin_messages: YES [ BY BIT 17 IN FIELD flags ],
      until_date: 0 [INT],
    },
*/

/*
banned:

banned_rights: { chatBannedRights
  flags: 255 [INT],
  view_messages: YES [ BY BIT 0 IN FIELD flags ],
  send_messages: YES [ BY BIT 1 IN FIELD flags ],
  send_media: YES [ BY BIT 2 IN FIELD flags ],
  send_stickers: YES [ BY BIT 3 IN FIELD flags ],
  send_gifs: YES [ BY BIT 4 IN FIELD flags ],
  send_games: YES [ BY BIT 5 IN FIELD flags ],
  send_inline: YES [ BY BIT 6 IN FIELD flags ],
  embed_links: YES [ BY BIT 7 IN FIELD flags ],
  send_polls: [ SKIPPED BY BIT 8 IN FIELD flags ],
  change_info: [ SKIPPED BY BIT 10 IN FIELD flags ],
  invite_users: [ SKIPPED BY BIT 15 IN FIELD flags ],
  pin_messages: [ SKIPPED BY BIT 17 IN FIELD flags ],
  until_date: 2147483647 [INT],
*/
func BannedRightToBanned(r *mtproto.ChatBannedRights) bool {
	return r.GetViewMessages() == false && // view message is false and other
		(r.GetSendMessages() == true ||
			r.GetSendMedia() == true ||
			r.GetSendStickers() == true ||
			r.GetSendGifs() == true ||
			r.GetSendGames() == true ||
			r.GetSendInline() == true ||
			r.GetEmbedLinks() == true ||
			r.GetSendPolls() == true ||
			r.GetChangeInfo() == true ||
			r.GetInviteUsers() == true ||
			r.GetPinMessages() == true)
}

func BannedRightToDeleted(r *mtproto.ChatBannedRights) bool {
	return r.GetViewMessages() == false && // all false
		r.GetSendMessages() == false &&
		r.GetSendMedia() == false &&
		r.GetSendStickers() == false &&
		r.GetSendGifs() == false &&
		r.GetSendGames() == false &&
		r.GetSendInline() == false &&
		r.GetEmbedLinks() == false &&
		r.GetSendPolls() == false &&
		r.GetChangeInfo() == false &&
		r.GetInviteUsers() == false &&
		r.GetPinMessages() == false
}

//func BannedRightToKick(r *mtproto.ChatBannedRights) bool {
//	return r.GetViewMessages() // all true
//}

func RightsToChatAdminBannedRight(r int32) *mtproto.ChatAdminRights {
	right := mtproto.NewTLChatAdminRights(nil)
	right.SetChangeInfo(checkRight(r, AdminRight_ChangeInfo))
	right.SetPostMessages(checkRight(r, AdminRight_PostMessages))
	right.SetEditMessages(checkRight(r, AdminRight_EditMessages))
	right.SetDeleteMessages(checkRight(r, AdminRight_DeleteMessages))
	right.SetBanUsers(checkRight(r, AdminRight_BanUsers))
	right.SetInviteUsers(checkRight(r, AdminRight_InviteUsers))
	right.SetPinMessages(checkRight(r, AdminRight_PinMessages))
	right.SetAddAdmins(checkRight(r, AdminRight_AddAdmins))
	right.SetAnonymous(checkRight(r, AdminRight_Anonymous))
	right.SetManageCall(checkRight(r, AdminRight_ManageCall))

	return right.To_ChatAdminRights()
}

func ChatBannedRightToRights(r *mtproto.ChatBannedRights) (int32, int32) {
	var right int32
	setRight(&right, r.GetViewMessages(), BannedRight_ViewMessages)
	setRight(&right, r.GetSendMessages(), BannedRight_SendMessages)
	setRight(&right, r.GetSendMedia(), BannedRight_SendMedia)
	setRight(&right, r.GetSendStickers(), BannedRight_SendStickers)
	setRight(&right, r.GetSendGifs(), BannedRight_SendGifs)
	setRight(&right, r.GetSendGames(), BannedRight_SendGames)
	setRight(&right, r.GetSendInline(), BannedRight_SendInline)
	setRight(&right, r.GetEmbedLinks(), BannedRight_EmbedLinks)
	setRight(&right, r.GetSendPolls(), BannedRight_SendPolls)
	setRight(&right, r.GetChangeInfo(), BannedRight_ChangeInfo)
	setRight(&right, r.GetInviteUsers(), BannedRight_InviteUsers)
	setRight(&right, r.GetPinMessages(), BannedRight_PinMessages)
	//setRight(&right, r.GetBanWisper(), BannedRight_BAN_WISPER)

	return right, r.GetUntilDate()
}

func ChatAdminBannedRightToRights(r *mtproto.ChatAdminRights) int32 {
	var right int32
	setRight(&right, r.GetChangeInfo(), AdminRight_ChangeInfo)
	setRight(&right, r.GetPostMessages(), AdminRight_PostMessages)
	setRight(&right, r.GetEditMessages(), AdminRight_EditMessages)
	setRight(&right, r.GetDeleteMessages(), AdminRight_DeleteMessages)
	setRight(&right, r.GetBanUsers(), AdminRight_BanUsers)
	setRight(&right, r.GetInviteUsers(), AdminRight_InviteUsers)
	//setRight(&right, r.GetInviteLink(), AdminRight_InviteLink)
	setRight(&right, r.GetPinMessages(), AdminRight_PinMessages)
	setRight(&right, r.GetAddAdmins(), AdminRight_AddAdmins)
	setRight(&right, r.GetAnonymous(), AdminRight_Anonymous)
	setRight(&right, r.GetManageCall(), AdminRight_ManageCall)

	return right
}
