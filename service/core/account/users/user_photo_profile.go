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

func AddPhotoId(m *data_users.ProfilePhotoIds, id int64, video int64) {
	idList := make([]*data_users.ProfilePhotoData, 0, len(m.IdList))
	idList = append(idList, &data_users.ProfilePhotoData{
		PhotoId: id,
		VideoId: video,
	})
	for _, v := range m.IdList {
		if id != v.PhotoId {
			idList = append(idList, &data_users.ProfilePhotoData{
				PhotoId: v.PhotoId,
				VideoId: v.VideoId,
			})
		}
	}
	m.IdList = idList
	m.Default = id
}

func RemovePhotoId(m *data_users.ProfilePhotoIds, id int64) int64 {
	if len(m.IdList) <= 1 {
		m.IdList = make([]*data_users.ProfilePhotoData, 0, len(m.IdList))
		m.Default = 0
	} else {
		if id == m.Default {
			id = m.IdList[1].PhotoId
			m.IdList = m.IdList[1:]
			if len(m.IdList) > 0 {
				m.Default = m.IdList[0].PhotoId
			} else {
				m.Default = 0
			}
		} else {
			idList := make([]*data_users.ProfilePhotoData, 0, len(m.IdList))
			for _, j := range m.IdList {
				if j.PhotoId != id {
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
