/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getSupportName_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  help.getSupportName#d360e72c = help.SupportName;
//
func (s *HelpServiceImpl) HelpGetSupportName(ctx context.Context, request *mtproto.TLHelpGetSupportName) (*mtproto.Help_SupportName, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetSupportName %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//supportName := mtproto.NewTLHelpSupportName(nil)
	//if s.conf.SupportUserID.SupportUserID != 0 {
	//	supportUserId := s.conf.SupportUserID.SupportUserID
	//
	//	if supportUserId != 0 {
	//		userInfo, err := s.accountCache.User(supportUserId)
	//		if err != nil {
	//			log.Warnf("HelpGetSupportName error:%s", err.Error())
	//		} else {
	//			user, err := account.CacheInfo2User(userInfo)
	//			if err != nil {
	//				log.Warnf("HelpGetSupportName CacheInfo2User error:%s", err.Error())
	//			} else {
	//				supportName.SetName(user.GetUsername())
	//			}
	//		}
	//		return supportName.To_Help_SupportName(), nil
	//	} else {
	//		log.Errorf("HelpGetSupportName not found SupportUserID.SupportUserID config:%+v", s.conf.SupportUserID)
	//		return supportName.To_Help_SupportName(), nil
	//	}
	//} else {
	log.Errorf("HelpGetSupportName not found SupportUserID.SupportUserID")
	return nil, fmt.Errorf("%s", "Not impl HelpServiceImpl")
	//}
}
