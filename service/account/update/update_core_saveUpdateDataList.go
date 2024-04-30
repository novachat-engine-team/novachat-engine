/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package update

import data_update "novachat_engine/service/data/update"

func (c *Core) SaveUpdateDataList(userId int64, updateDataList []*data_update.UserUpdate) error {
	if len(updateDataList) == 0 {
		return nil
	}

	return c.updateCore.SaveUpdateDataList(userId, updateDataList)
}
