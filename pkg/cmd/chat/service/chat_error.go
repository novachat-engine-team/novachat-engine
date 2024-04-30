/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package service

import (
	"fmt"
)

var DefaultErrChatIdInvalid = &ErrChatIdInvalid{}

type ErrChatIdInvalid struct {
	chatId int64
}

func (e *ErrChatIdInvalid) Error() string {
	return fmt.Sprintf("Invalid chatId = %d", e.chatId)
}

func (e *ErrChatIdInvalid) Is(err error) bool {
	return err == DefaultErrChatIdInvalid
}

var DefaultErrUserIdInvalid = &ErrUserIdInvalid{}

type ErrUserIdInvalid struct {
	userId int64
}

func (e *ErrUserIdInvalid) Error() string {
	return fmt.Sprintf("Invalid userId = %d", e.userId)
}

func (e *ErrUserIdInvalid) Is(err error) bool {
	return err == DefaultErrUserIdInvalid
}
