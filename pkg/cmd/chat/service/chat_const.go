/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package service

import "fmt"

const (
	keyChat       string = "string:chat"
	keyCommonChat string = "set:commonChat:list"
)

const defaultChatMax int32 = 200

func makeChatKey(chatId int64) string {
	return fmt.Sprintf("%s%d", keyChat, chatId)
}

func makeCommonChatKey(userId int64) string {
	return fmt.Sprintf("%s%d", keyCommonChat, userId)
}
