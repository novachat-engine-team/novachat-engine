/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package chat

type GetChatListOptions func(*getChatListOptions)

type getChatListOptions struct {
	Ids       []int64
	ExceptIds []int64
}

func GetChatListWithIds(ids []int64) GetChatListOptions {
	return func(opt *getChatListOptions) {
		opt.Ids = ids
	}
}

func GetChatListWithExceptIds(ids []int64) GetChatListOptions {
	return func(opt *getChatListOptions) {
		opt.ExceptIds = ids
	}
}
