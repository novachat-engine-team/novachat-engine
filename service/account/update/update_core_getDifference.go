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
	"time"
)

func (c *Core) GetUpdateDifference(authKeyId int64, userId int64, pts int32, qts int32, date int32, getDifferenceFirstSync bool, limit int32, layer int32) (*mtproto.Updates_Difference, error) {
	updateList, err := c.updateCore.GetUpdateDifference(userId, pts, qts, date, limit, getDifferenceFirstSync)
	if err != nil {
		log.Errorf("GetDifference authKeyId:%d userId:%d error:%s", authKeyId, userId, err.Error())
		return nil, err
	}

	updateState := mtproto.NewTLUpdatesState(nil).To_Updates_State()
	updateDifference := &mtproto.Updates_Difference{
		Date:                 int32(time.Now().Unix()),
		Seq:                  0,
		NewMessages:          []*mtproto.Message{},
		NewEncryptedMessages: []*mtproto.EncryptedMessage{},
		OtherUpdates:         []*mtproto.Update{},
		Chats:                nil,
		Users:                nil,
		Pts:                  0,
	}
	if len(updateList) == 0 {
		log.Infof("GetDifference authKeyId:%d userId:%d UpdatesDifferenceEmpty", authKeyId, userId)
		return mtproto.NewTLUpdatesDifferenceEmpty(updateDifference).To_Updates_Difference(), nil
	}

	if getDifferenceFirstSync {
		return mtproto.NewTLUpdatesDifferenceTooLong(&mtproto.Updates_Difference{
			Pts: updateList[0].Pts,
		}).To_Updates_Difference(), nil
	}

	updateState.Date = updateList[len(updateList)-1].Date

	userIdList := make([]int64, 0, 24)
	for _, u := range updateList {
		update := UserUpdateToUpdate(u)
		FilterAppendUserId(&userIdList, constants.PeerTypeFromUserIDType32(update.UserId).ToInt())
		FilterAppendPeer(&userIdList, update.PeerId)
		FilterAppendPeer(&userIdList, update.Peer9961FD5C71)
		for _, o := range update.OrderD8CAF68D71 {
			FilterAppendPeer(&userIdList, o)
		}

		switch constants.UpdateTypeFromInt32(u.UpdateType) {
		case constants.UpdateTypeUpdateNewMessage:
			if update.GetPts() > updateState.Pts {
				updateState.Pts = update.GetPts()
			}
			fallthrough
		case constants.UpdateTypeUpdateNewChannelMessage:
			updateDifference.NewMessages = append(updateDifference.NewMessages, update.GetMessage1F2B0AFD71())
			if update.GetMessage1F2B0AFD71() != nil && update.GetMessage1F2B0AFD71().Media != nil && update.GetMessage1F2B0AFD71().Media.ClassName == mtproto.ClassMessageMediaContact {
				FilterAppendUserId(&userIdList, constants.PeerTypeFromUserIDType32(update.GetMessage1F2B0AFD71().Media.UserId).ToInt())
			}
			FilterAppendPeer(&userIdList, update.GetMessage1F2B0AFD71().PeerId)
			FilterAppendPeer(&userIdList, update.GetMessage1F2B0AFD71().ToId)
			FilterAppendPeer(&userIdList, update.GetMessage1F2B0AFD71().FromId286FA604119)
			FilterAppendUserId(&userIdList, constants.PeerTypeFromUserIDType32(update.GetMessage1F2B0AFD71().FromId90DDDC1171).ToInt())
			break
		case constants.UpdateTypeUpdateNewEncryptedMessage:
			updateDifference.NewEncryptedMessages = append(updateDifference.NewEncryptedMessages, update.GetMessage12BCBD9A71())
			if update.GetQts() > updateState.Qts {
				updateState.Qts = update.GetQts()
			}
			break
		case constants.UpdateTypeUpdateEditMessage,
			constants.UpdateTypeUpdateDeleteMessages,
			constants.UpdateTypeUpdateReadHistoryOutbox,
			constants.UpdateTypeUpdateReadHistoryInbox,
			constants.UpdateTypeUpdateReadMessagesContents,
			constants.UpdateTypeUpdateFolderPeers,
			constants.UpdateTypeUpdateRequestContacts:
			updateDifference.OtherUpdates = append(updateDifference.OtherUpdates, update)
			if update.GetPts() > updateState.Pts {
				updateState.Pts = update.GetPts()
			}
			break
		case constants.UpdateTypeUpdateReadChannelInbox,
			constants.UpdateTypeUpdateDeleteChannelMessages,
			constants.UpdateTypeUpdateReadChannelOutbox,
			constants.UpdateTypeUpdateChannelReadMessagesContents,
			constants.UpdateTypeUpdateChannelMessageForwards,
			constants.UpdateTypeUpdateChannelPinnedMessage:
			updateDifference.OtherUpdates = append(updateDifference.OtherUpdates, update)
			break
		default:
			if update.GetPts() > updateState.Pts {
				updateState.Pts = update.GetPts()
			}
			log.Infof("GetDifference updateType:%v", constants.UpdateTypeFromInt32(u.UpdateType))
			break
		}
	}

	userList, err := c.accountUsersCore.GetUserList(userId, userIdList, layer)
	if err != nil {
		log.Warnf("GetDifference GetUserList error:%s", err.Error())
	}
	updateDifference.Users = userList

	updateDifference.Pts = updateState.Pts

	if int32(len(updateList)) >= limit {
		updateDifference.IntermediateState = updateState
		return mtproto.NewTLUpdatesDifferenceSlice(updateDifference).To_Updates_Difference(), nil
	} else {
		updateDifference.State = updateState
		return mtproto.NewTLUpdatesDifference(updateDifference).To_Updates_Difference(), nil
	}
}
