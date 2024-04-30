/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package bot

import (
	"novachat_engine/mtproto"
	"novachat_engine/service/constants"
)

func (c *Core) BotInfo(botId int64) *mtproto.BotInfo {
	botInfo := mtproto.NewTLBotInfo(&mtproto.BotInfo{
		UserId:      constants.PeerTypeFromUserIDType(botId).ToInt32(),
		Description: "bot info",
		Commands:    nil,
	}).To_BotInfo()

	botInfo.Commands = make([]*mtproto.BotCommand, 0, 5)
	botInfo.Commands = append(botInfo.Commands, mtproto.NewTLBotCommand(&mtproto.BotCommand{
		Command:     "help",
		Description: "help",
	}).To_BotCommand())
	return botInfo
}
