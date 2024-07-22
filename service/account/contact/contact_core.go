/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package contact

import (
	"novachat_engine/service/core/account/contact"
	data_contact "novachat_engine/service/data/contact"
)

type Blocked int32

var (
	BlockedNone  Blocked = 0
	BlockedSelf  Blocked = 1
	BlockedOther Blocked = 2
)

type Core struct {
	contactCore *contact.Core
}

func NewContactCore(c *contact.Core) *Core {
	return &Core{
		contactCore: c,
	}
}

func (c *Core) CheckBlock(userId int64, peerId int64, now int32) (Blocked, error) {
	if userId == peerId || peerId == 0 {
		return BlockedNone, nil
	}

	return BlockedNone, nil
}

func (c *Core) Block(userId int64, peerId int64, now int32) error {
	return c.contactCore.ContactsBlock(userId, peerId, now)
}

func (c *Core) GetBlocked(userId int64) ([]*data_contact.Contact, error) {
	return c.contactCore.GetBlocked(userId)
}

func (c *Core) ContactsUnblock(userId int64, peerId int64, now int32) (bool, error) {
	return c.contactCore.ContactsUnblock(userId, peerId, now)
}

func (c *Core) GetContactById(userId int64, peerId int64) (*data_contact.Contact, error) {
	return c.contactCore.GetContactById(userId, peerId)
}

func (c *Core) GetContactByIdList(userId int64, peerIdList []int64) ([]*data_contact.Contact, error) {
	return c.contactCore.GetContactByIdList(userId, peerIdList)
}

func (c *Core) GetContacts(userId int64) ([]*data_contact.Contact, error) {
	return c.contactCore.GetContacts(userId)
}

func (c *Core) AddContact(userId int64, phone string, peerId int64, name string, name2 string, now int32) (data_contact.MutualType, error) {
	return c.contactCore.AddContact(userId, phone, peerId, name, name2, now)
}

func (c *Core) ContactsRemove(userId int64, list []int64) (bool, error) {
	return c.contactCore.ContactsRemove(userId, list)
}

func (c *Core) ModifyContact(userId int64, phone string, peerId int64, firstname string, lastname string, now int32) error {
	return c.contactCore.ModifyContact(userId, phone, peerId, firstname, lastname, now)
}
