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
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
)

func (c *Core) GetUpdateChannelDifference(userId, chat int64, force bool, pts int32, limit int32, layer int32) (*mtproto.Updates_ChannelDifference, error) {

	updateChannelDifference := mtproto.NewTLUpdatesChannelDifference(&mtproto.Updates_ChannelDifference{
		Final:        false,
		Pts:          pts,
		NewMessages:  []*mtproto.Message{},
		OtherUpdates: []*mtproto.Update{},
		Chats:        []*mtproto.Chat{},
		Users:        []*mtproto.User{},
	}).To_Updates_ChannelDifference()
	updateList, err := c.updateCore.GetUpdateChannelDifference(userId, chat, pts, limit)
	if err != nil {
		log.Errorf("GetUpdateChannelDifference error:%s", err.Error())
		return nil, err
	}

	pts = 0
	userIdList := make([]int64, 0, 24)
	for _, u := range updateList {
		update := UserUpdateToUpdate(u)
		FilterAppendUserId(&userIdList, constants.PeerTypeFromUserIDType32(update.UserId).ToInt())

		switch constants.UpdateTypeFromInt32(u.UpdateType) {
		case constants.UpdateTypeUpdateNewMessage:
			fallthrough
		case constants.UpdateTypeUpdateNewChannelMessage:
			if update.GetPts() > pts {
				pts = update.GetPts()
			}
			updateChannelDifference.NewMessages = append(updateChannelDifference.NewMessages, update.GetMessage1F2B0AFD71())
			FilterAppendPeer(&userIdList, update.GetMessage1F2B0AFD71().FromId286FA604119)
			FilterAppendUserId(&userIdList, constants.PeerTypeFromUserIDType32(update.GetMessage1F2B0AFD71().FromId90DDDC1171).ToInt())
			break
		case constants.UpdateTypeUpdateEditChannelMessage,
			constants.UpdateTypeUpdateReadChannelInbox,
			constants.UpdateTypeUpdateDeleteChannelMessages,
			constants.UpdateTypeUpdateReadChannelOutbox,
			constants.UpdateTypeUpdateChannelReadMessagesContents,
			constants.UpdateTypeUpdateChannelMessageForwards,
			constants.UpdateTypeUpdateChannelPinnedMessage:
			updateChannelDifference.OtherUpdates = append(updateChannelDifference.OtherUpdates, update)
			if update.GetPts() > pts {
				pts = update.GetPts()
			}
			break
		default:
			if update.GetPts() > pts {
				pts = update.GetPts()
			}
			log.Infof("GetUpdateChannelDifference updateType:%v", constants.UpdateTypeFromInt32(u.UpdateType))
			break
		}
	}
	updateChannelDifference.Pts = pts
	userList, err := c.accountUsersCore.GetUserList(userId, userIdList, layer)
	if err != nil {
		log.Warnf("GetDifference GetUserList error:%s", err.Error())
	}
	updateChannelDifference.Users = userList

	return updateChannelDifference, nil
}
