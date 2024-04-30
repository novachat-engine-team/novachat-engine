/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/9/2 13:43
 * @File : user_photo_profile.go
 */

package users

import (
	"github.com/json-iterator/go"
	data_users "novachat_engine/service/data/users"
)

func MakeProfilePhotoData(jsonData string) *data_users.ProfilePhotoIds {
	if jsonData == "" {
		return &data_users.ProfilePhotoIds{}
	}
	data2 := &data_users.ProfilePhotoIds{}
	err := jsoniter.Unmarshal([]byte(jsonData), data2)
	if err != nil {
		return &data_users.ProfilePhotoIds{}
	}
	return data2
}

func AddPhotoId(m *data_users.ProfilePhotoIds, id int64) {
	idList := make([]int64, 0, len(m.IdList))
	idList = append(idList, id)
	for _, v := range m.IdList {
		if id != v {
			idList = append(idList, v)
		}
	}
	m.IdList = idList
	m.Default = id
}

func RemovePhotoId(m *data_users.ProfilePhotoIds, id int64) int64 {
	if len(m.IdList) <= 1 {
		m.IdList = []int64{}
		m.Default = 0
	} else {
		if id == m.Default {
			id = m.IdList[1]
			m.IdList = m.IdList[1:]
			if len(m.IdList) > 0 {
				m.Default = m.IdList[0]
			} else {
				m.Default = 0
			}
		} else {

			idList := make([]int64, 0, len(m.IdList))
			for _, j := range m.IdList {
				if j != id {
					idList = append(idList, j)
				}
			}
			m.IdList = idList
		}
	}
	return m.Default
}

func ToProfilePhotoIDsJson(m *data_users.ProfilePhotoIds) string {
	data, err := jsoniter.Marshal(m)
	if err != nil {
		return ""
	}
	return string(data)
}
