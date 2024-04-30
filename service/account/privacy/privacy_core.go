/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package privacy

import (
	"novachat_engine/service/core/account/privacy"
	"novachat_engine/service/data/privacy"
)

type Core struct {
	pri *privacy.Core
}

func NewPrivacyCore(p *privacy.Core) *Core {

	return &Core{
		pri: p,
	}
}

func (c *Core) GetPrivacy(userId int64, key int32) ([]*data_privacy.PrivacyOption, error) {
	return c.pri.GetPrivacy(userId, key)
}

func (c *Core) GetPrivacyKeys(userId int64, keys []int32) (*data_privacy.Privacy, error) {
	return c.pri.GetPrivacyKeys(userId, keys)
}

func (c *Core) UpdatePrivacy(userId int64, key int32, updatePrivacy []*data_privacy.PrivacyOption) (bool, error) {
	return c.pri.UpdatePrivacy(userId, key, updatePrivacy)
}
