/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.search_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/util"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"strings"
)

//  contacts.search#11f812d8 q:string limit:int = contacts.Found;
//
func (s *ContactsServiceImpl) ContactsSearch(ctx context.Context, request *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("ContactsSearch %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	query := request.Q
	if len(query) > 0 && query[0] == '@' {
		query = query[1:]
	}
	query = strings.ReplaceAll(query, "_", "\\_")

	if len(query) == 0 {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_SEARCH_QUERY_EMPTY)
		log.Debugf("ContactsSearch %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	if len(query) <= 1 {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_QUERY_TOO_SHORT)
		log.Debugf("ContactsSearch %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	found := mtproto.NewTLContactsFound(nil)
	if len(query) == 0 || request.Limit <= 0 {
		return found.To_Contacts_Found(), nil
	}

	if util.IsAllNumStringLimit(query, 0) == false {
		return found.To_Contacts_Found(), nil
	}

	userIdMap := map[int64]struct{}{}
	contactList, err := s.accountContactCore.GetContacts(md.UserId)
	if err != nil {
		log.Warnf("ContactsSearch GetContacts error:%s", err.Error())
	} else {
		if len(contactList) > 0 {
			userIdList := make([]int64, 0, len(contactList))
			for _, v := range contactList {
				userIdList = append(userIdList, v.UserId)
				userIdMap[v.UserId] = struct{}{}
			}
			userInfoList, _ := s.accountUserCore.UserDataList(userIdList)
			myResultPeer := make([]*mtproto.Peer, 0, request.Limit)
			for _, v := range userInfoList {
				if !v.Username.Valid && strings.Index(v.Username.String, query) >= 0 && int32(len(myResultPeer)) < request.Limit {
					myResultPeer = append(myResultPeer, mtproto.NewTLPeerUser(&mtproto.Peer{
						UserId: constants.PeerTypeFromUserIDType(v.Id).ToInt32(),
					}).To_Peer())
				}
			}
			found.SetMyResults(myResultPeer)
		}
	}

	if len(query) > 4 {
		userDataList, _ := s.accountUserCore.SearchUserData(query, 0, request.Limit*3)
		resultPeer := make([]*mtproto.Peer, 0, len(userDataList))
		for _, v := range userDataList {
			resultPeer = append(resultPeer, mtproto.NewTLPeerUser(&mtproto.Peer{
				UserId: constants.PeerTypeFromUserIDType(v.Id).ToInt32(),
			}).To_Peer())
			userIdMap[v.Id] = struct{}{}
		}
		found.SetResults(resultPeer)
	}

	if len(userIdMap) > 0 {
		userIdList := make([]int64, 0, len(userIdMap))
		for k := range userIdMap {
			userIdList = append(userIdList, k)
		}
		userList, _ := s.accountUserCore.GetUserList(md.UserId, userIdList, md.Layer)
		found.SetUsers(userList)
	}

	//found.SetChats(chatsList)
	log.Debugf("ContactsSearch %v, request: %v userList found:%+v reply ok!!!!!!", md, request, found)
	return found.To_Contacts_Found(), nil
}
