/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : updates.getDifference_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/compat"
	data_users "novachat_engine/service/data/users"
)

//pts	int	PTS, see updates.
//pts_total_limit	flags.0?int	For fast updating: if provided and pts + pts_total_limit < remote pts, updates.differenceTooLong will be returned.
//Simply tells the server to not return the difference if it is bigger than pts_total_limit
//If the remote pts is too big (> ~4000000), this field will default to 1000000
//date	int	date, see updates.
//qts	int	QTS, see updates.

//It is recommended to use limit 10-100 for channels and 1000-10000 otherwise.

//401	AUTH_KEY_PERM_EMPTY	The temporary auth key must be binded to the permanent auth key to use these methods.
//400	CDN_METHOD_INVALID	You can't call this method in a CDN DC
//400	DATE_EMPTY	Date empty
//400	PERSISTENT_TIMESTAMP_EMPTY	Persistent timestamp empty
//400	PERSISTENT_TIMESTAMP_INVALID	Persistent timestamp invalid

//  updates.getDifference#25939651 flags:# pts:int pts_total_limit:flags.0?int date:int qts:int = updates.Difference;
//
func (s *UpdatesServiceImpl) UpdatesGetDifference(ctx context.Context, request *mtproto.TLUpdatesGetDifference) (*mtproto.Updates_Difference, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("UpdatesGetDifference %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	limit := request.PtsTotalLimit
	getDifferenceFirstSync := false
	if request.PtsTotalLimit == 0 {
		limit = 100
	}

	var userInfo *data_users.Users
	userInfo, err := s.accountUsersCore.UserData(md.UserId)
	if err != nil {
		log.Errorf("UpdatesGetDifference - request: %v GetByUserId error:%s", request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}

	if userInfo.GetDeleted() { //|| userInfo.GetRestricted() {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_UNREGISTERED)
		log.Errorf("UpdatesGetDifference - request: %v error:%s", request, err.Error())
		return nil, err
	}

	if request.Date == 0 {
		//err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_DATE_EMPTY)
		//log.Errorf("UpdatesGetDifference %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		//return nil, err
		log.Warnf("UpdatesGetDifference %v, request: %v date == 0", metadata.RpcMetaDataDebug(md), request)
	}
	//
	if request.Qts != 0 {
		//TODO:(Coder)
		//RpcErrorCode_BAD_REQUEST_PERSISTENT_TIMESTAMP_EMPTY;
		//RpcErrorCode_BAD_REQUEST_PERSISTENT_TIMESTAMP_INVALID;
	} else if request.Pts <= 0 && request.Date == 0 {
		getDifferenceFirstSync = true
	}

	updateDifferent, err := s.accountUpdateCore.GetUpdateDifference(md.AuthKeyId, md.UserId, request.Pts, request.Qts, request.Date, getDifferenceFirstSync, limit)
	if err != nil {
		log.Errorf("UpdatesGetDifference %v, request: %v GetDifference error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
	}

	up := s.phoneCallCore.CheckComingPhoneCallAndUpdates(md.UserId)
	if up != nil {
		updateDifferent.OtherUpdates = append(updateDifferent.OtherUpdates, up.Updates...)
	}

	for idx := range updateDifferent.NewMessages {
		updateDifferent.NewMessages[idx] = compat.MessageCompat(updateDifferent.NewMessages[idx], md.Layer)
	}

	for idx := range updateDifferent.OtherUpdates {
		updateDifferent.OtherUpdates[idx] = compat.UpdateCompat(updateDifferent.OtherUpdates[idx], md.Layer)
	}

	for idx := range updateDifferent.Users {
		updateDifferent.Users[idx] = compat.CompatUser(updateDifferent.Users[idx], md.Layer)
	}

	for idx := range updateDifferent.Chats {
		updateDifferent.Chats[idx] = compat.CompatChat(updateDifferent.Chats[idx], md.Layer)
	}
	//updateState.Date = updateDifferent.Date
	//updateDifferent.State = updateState
	//updateDifferent.IntermediateState = updateState

	log.Infof("UpdatesGetDifference %v, request: %v GetDifference:%+v reply ok!!!!!!!!", metadata.RpcMetaDataDebug(md), request, updateDifferent)
	return updateDifferent, nil
}
