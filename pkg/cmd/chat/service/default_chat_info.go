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
	data_chat "novachat_engine/service/data/chat"
)

func DefaultChatData(date int32) *data_chat.Chat {
	return &data_chat.Chat{
		Title:             "",
		Photo:             nil,
		Creator:           0,
		Date:              date,
		Verified:          false,
		Scam:              false,
		AccessHash:        0,
		BannedRights:      0,
		GeoPoint:          nil,
		RestrictionReason: nil,
		RightsUtilDate:    0,
		Deleted:           false,
	}
}
