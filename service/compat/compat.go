/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/15 15:01
 * @File : photo_compat.com.go
 */

package compat

import "novachat_engine/mtproto"

func CompatPhoto(photo *mtproto.Photo, layer int32) *mtproto.Photo {
	if photo == nil {
		return photo
	}

	for idx := range photo.Sizes {
		if photo.Sizes[idx] == nil {
			continue
		}

		photo.Sizes[idx].Location = CompatFileLocation(photo.Sizes[idx].Location, layer)
	}

	for idx := range photo.VideoSizes {
		if photo.VideoSizes[idx] == nil {
			continue
		}

		photo.VideoSizes[idx].Location = CompatFileLocation(photo.VideoSizes[idx].Location, layer)
	}
	return photo
}

func CompatFileLocation(fileLocation *mtproto.FileLocation, layer int32) *mtproto.FileLocation {
	if fileLocation == nil {
		return fileLocation
	}

	if fileLocation.ClassName == mtproto.ClassFileLocationUnavailable {
		return fileLocation
	}

	fileLocation = NewFileLocationByLayer(fileLocation.VolumeId, fileLocation.LocalId, layer, fileLocation.DcId)
	return fileLocation
}

func CompatGame(game *mtproto.Game, layer int32) *mtproto.Game {
	if game == nil {
		return game
	}
	game.Photo = CompatPhoto(game.Photo, layer)
	game.Document = CompatDocument(game.Document, layer)

	return game
}

func CompatDocument(document *mtproto.Document, layer int32) *mtproto.Document {
	if document == nil {
		return document
	}

	if document.Thumb != nil {
		document.Thumb.Location = CompatFileLocation(document.Thumb.Location, layer)
	}

	for idx := range document.Thumbs {
		if document.Thumbs[idx] == nil {
			continue
		}

		document.Thumbs[idx].Location = CompatFileLocation(document.Thumbs[idx].Location, layer)
	}

	for idx := range document.VideoThumbs {
		if document.VideoThumbs[idx] == nil {
			continue
		}

		document.VideoThumbs[idx].Location = CompatFileLocation(document.VideoThumbs[idx].Location, layer)
	}
	return document
}

func CompatWebPage(webPage *mtproto.WebPage, layer int32) *mtproto.WebPage {
	if webPage == nil {
		return webPage
	}

	webPage.Photo = CompatPhoto(webPage.Photo, layer)
	webPage.Document = CompatDocument(webPage.Document, layer)
	webPage.CachedPage = CompatPage(webPage.CachedPage, layer)

	for idx := range webPage.Attributes {
		webPage.Attributes[idx] = CompatWebPageAttribute(webPage.Attributes[idx], layer)
	}
	return webPage
}

func CompatPage(page *mtproto.Page, layer int32) *mtproto.Page {
	if page == nil {
		return page
	}

	for idx := range page.Photos {
		page.Photos[idx] = CompatPhoto(page.Photos[idx], layer)
	}

	for idx := range page.Documents {
		page.Documents[idx] = CompatDocument(page.Documents[idx], layer)
	}

	return page
}

func CompatWebPageAttribute(attr *mtproto.WebPageAttribute, layer int32) *mtproto.WebPageAttribute {
	if attr == nil {
		return nil
	}

	for idx := range attr.Documents {
		attr.Documents[idx] = CompatDocument(attr.Documents[idx], layer)
	}

	if attr.Settings != nil && attr.Settings.Wallpaper != nil {
		attr.Settings.Wallpaper.Document = CompatDocument(attr.Settings.Wallpaper.Document, layer)

		for idx := range attr.Settings.Wallpaper.Sizes{
			attr.Settings.Wallpaper.Sizes[idx].Location = CompatFileLocation(attr.Settings.Wallpaper.Sizes[idx].Location, layer)
		}
	}

	return attr
}

func CompatDialog(dialog *mtproto.Dialog, layer int32) *mtproto.Dialog {
	if dialog == nil {
		return dialog
	}
	if dialog.Folder != nil {
		if dialog.Folder.Photo != nil {
			if dialog.Folder.Photo.PhotoBig != nil {
				dialog.Folder.Photo.PhotoBig = CompatFileLocation(dialog.Folder.Photo.PhotoBig, layer)
			}

			if dialog.Folder.Photo.PhotoSmall != nil {
				dialog.Folder.Photo.PhotoSmall = CompatFileLocation(dialog.Folder.Photo.PhotoSmall, layer)
			}
		}
	}
	return dialog
}

func CompatChat(chat *mtproto.Chat, layer int32) *mtproto.Chat {
	if chat == nil {
		return chat
	}
	if chat.Photo != nil {
		if chat.Photo.PhotoBig != nil {
			chat.Photo.PhotoBig = CompatFileLocation(chat.Photo.PhotoBig, layer)
		}

		if chat.Photo.PhotoSmall != nil {
			chat.Photo.PhotoSmall = CompatFileLocation(chat.Photo.PhotoSmall, layer)
		}
	}
	return chat
}

func CompatUser(user *mtproto.User, layer int32) *mtproto.User {
	if user == nil {
		return user
	}
	if user.Photo != nil {
		if user.Photo.PhotoBig != nil {
			user.Photo.PhotoBig = CompatFileLocation(user.Photo.PhotoBig, layer)
		}

		if user.Photo.PhotoSmall != nil {
			user.Photo.PhotoSmall = CompatFileLocation(user.Photo.PhotoSmall, layer)
		}
	}
	return user
}