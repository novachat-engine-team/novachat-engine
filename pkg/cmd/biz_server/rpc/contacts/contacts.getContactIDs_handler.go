/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : contacts.getContactIDs_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
)

//  contacts.getContactIDs#2caa4a42 hash:int = Vector<int>;
//
func (s *ContactsServiceImpl) ContactsGetContactIDs(ctx context.Context, request *mtproto.TLContactsGetContactIDs) (*mtproto.VectorInt, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ContactsGetContactIDs %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	v, err := s.accountContactCore.GetContacts(md.UserId)
	if err != nil {
		log.Errorf("ContactsGetContactIDs - request: %v error:%s", request, err.Error())
		return nil, err
	}

	ret := &mtproto.VectorInt{
		Datas: make([]int32, 0, len(v)),
	}
	for k := range v {
		ret.Datas = append(ret.Datas, constants.PeerTypeFromUserIDType(v[k].UserId).ToInt32())
	}

	return ret, nil
}
