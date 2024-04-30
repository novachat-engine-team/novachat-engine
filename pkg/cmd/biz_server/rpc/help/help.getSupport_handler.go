/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getSupport_handler.go
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

//  help.getSupport#9cdf08cd = help.Support;
//
func (s *HelpServiceImpl) HelpGetSupport(ctx context.Context, request *mtproto.TLHelpGetSupport) (*mtproto.Help_Support, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetSupport %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//if s.conf.SupportUserID.SupportUserID != 0 {
	//	supportUserId := s.conf.SupportUserID.SupportUserID
	//	if supportUserId != 0 {
	//		support := mtproto.NewTLHelpSupport(nil)
	//		userInfo, err := s.accountCache.User(supportUserId)
	//		if err != nil {
	//			log.Warnf("HelpGetSupport error:%s", err.Error())
	//		} else {
	//			user, err := account.CacheInfo2User(userInfo)
	//			if err != nil {
	//				log.Warnf("HelpGetSupport CacheInfo2User error:%s", err.Error())
	//			} else {
	//				support.SetUser(usersUtil.UserCoreSelfUsers(usersUtil.UserCore2Users(user)))
	//				support.SetPhoneNumber(user.GetPhone())
	//			}
	//		}
	//		return support.To_Help_Support(), nil
	//	} else {
	//		log.Errorf("HelpGetSupport not found SupportUserID.SupportUserID config:%+v", s.conf.SupportUserID)
	//		return nil, fmt.Errorf("%s", "Not impl HelpGetSupport: not found")
	//	}
	//} else {
	//	log.Errorf("HelpGetSupport not found SupportUserID.SupportUserID")
	//	return nil, fmt.Errorf("%s", "Not impl HelpGetSupport: not found")
	//}
	log.Errorf("HelpGetSupport not found SupportUserID.SupportUserID")
	return nil, fmt.Errorf("%s", "Not impl HelpGetSupport: not found")
}
