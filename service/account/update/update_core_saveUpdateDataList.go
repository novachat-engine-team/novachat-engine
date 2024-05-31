/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package update

import (
	"novachat_engine/pkg/log"
	data_update "novachat_engine/service/data/update"
)

func (c *Core) SaveUpdateDataList(userId int64, updateDataList []*data_update.UserUpdate, updateChannelDataList []*data_update.UserUpdate) error {
	var err error
	if len(updateChannelDataList) > 0 {
		err = c.updateCore.SaveUpdateChannelDataList(updateChannelDataList)
		if err != nil {
			log.Errorf("SaveUpdateDataList updateChannelDataList error:%s", err.Error())
			return err
		}
	}

	if len(updateDataList) > 0 {
		err = c.updateCore.SaveUpdateDataList(userId, updateDataList)
		if err != nil {
			log.Errorf("SaveUpdateDataList updateDataList error:%s", err.Error())
			return err
		}
	}

	return nil
}
