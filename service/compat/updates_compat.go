/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/15 15:46
 * @File : updates_compat.go
 */

package compat

import (
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
)

func MessageCompat(message *mtproto.Message, layer int32) *mtproto.Message {
	if message == nil {
		return message
	}

	if message.Media != nil {
		message.Media.PhotoB5223B0F71 = CompatPhoto(message.Media.PhotoB5223B0F71, layer)
		message.Media.Game = CompatGame(message.Media.Game, layer)
		message.Media.Webpage = CompatWebPage(message.Media.Webpage, layer)
		message.Media.Document = CompatDocument(message.Media.Document, layer)
	}

	if message.Action != nil {
		message.Action.Photo = CompatPhoto(message.Action.Photo, layer)
	}

	return message
}

func UpdateCompat(update *mtproto.Update, layer int32) *mtproto.Update {
	if update == nil {
		return update
	}

	if update.ClassName == mtproto.ClassUpdateUserBlocked {
		update = NewUserBlock(constants.PeerTypeFromUserIDType32(update.UserId).ToInt(), mtproto.ToBool(update.Blocked), layer)
	}
	if update.ClassName == mtproto.ClassUpdatePeerBlocked {
		update = NewUserBlock(constants.PeerTypeFromUserIDType32(update.PeerId.UserId).ToInt(), mtproto.ToBool(update.Blocked), layer)
	}

	if update.Message1F2B0AFD71 != nil {
		if update.Message1F2B0AFD71.Media != nil {
			update.Message1F2B0AFD71.Media.PhotoB5223B0F71 = CompatPhoto(update.Message1F2B0AFD71.Media.PhotoB5223B0F71, layer)
			update.Message1F2B0AFD71.Media.Game = CompatGame(update.Message1F2B0AFD71.Media.Game, layer)
			update.Message1F2B0AFD71.Media.Webpage = CompatWebPage(update.Message1F2B0AFD71.Media.Webpage, layer)
			update.Message1F2B0AFD71.Media.Document = CompatDocument(update.Message1F2B0AFD71.Media.Document, layer)
		}

		if update.Message1F2B0AFD71.Action != nil {
			update.Message1F2B0AFD71.Action.Photo = CompatPhoto(update.Message1F2B0AFD71.Action.Photo, layer)
		}
	}

	if update.Media != nil {
		update.Media.PhotoB5223B0F71 = CompatPhoto(update.Media.PhotoB5223B0F71, layer)
		update.Media.Game = CompatGame(update.Media.Game, layer)
		update.Media.Webpage = CompatWebPage(update.Media.Webpage, layer)
		update.Media.Document = CompatDocument(update.Media.Document, layer)
	}

	if update.Photo != nil {
		update.Photo.PhotoSmall = CompatFileLocation(update.Photo.PhotoSmall, layer)
		update.Photo.PhotoBig = CompatFileLocation(update.Photo.PhotoBig, layer)
	}

	update.Stickerset = CompatMessageStickerset(update.Stickerset, layer)

	return update
}

func CompatMessageStickerset(stickerset *mtproto.Messages_StickerSet, layer int32) *mtproto.Messages_StickerSet {
	if stickerset == nil {
		return nil
	}
	stickerset.Set = CompatStickerset(stickerset.Set, layer)

	for idx := range stickerset.Documents {
		stickerset.Documents[idx] = CompatDocument(stickerset.Documents[idx], layer)
	}
	return stickerset
}

func CompatStickerset(set *mtproto.StickerSet, layer int32) *mtproto.StickerSet {

	for idx := range set.Thumbs {
		if set.Thumbs[idx] == nil {
			continue
		}

		set.Thumbs[idx].Location = CompatFileLocation(set.Thumbs[idx].Location, layer)
	}

	return set
}

func UpdatesCompat(updates *mtproto.Updates, layer int32) *mtproto.Updates {
	if updates == nil {
		return updates
	}

	updates.Update = UpdateCompat(updates.Update, layer)

	for idx := range updates.Updates {
		updates.Updates[idx] = UpdateCompat(updates.Updates[idx], layer)
	}

	if updates.Media != nil {
		updates.Media.PhotoB5223B0F71 = CompatPhoto(updates.Media.PhotoB5223B0F71, layer)
		updates.Media.Game = CompatGame(updates.Media.Game, layer)
		updates.Media.Webpage = CompatWebPage(updates.Media.Webpage, layer)
	}

	return updates
}
