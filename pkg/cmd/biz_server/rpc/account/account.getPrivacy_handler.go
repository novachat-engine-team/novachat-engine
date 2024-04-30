/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.getPrivacy_handler.go
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
	"novachat_engine/service/constants"
)

//  account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
//
func (s *AccountServiceImpl) AccountGetPrivacy(ctx context.Context, request *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountGetPrivacy %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	key := constants.ToPrivacyKey(request.Key)
	privacyValue, err := s.privacyCore.GetPrivacy(md.UserId, key.Int32())
	if err != nil {
		log.Errorf("AccountGetPrivacy %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())

		return mtproto.NewTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
			Rules: []*mtproto.PrivacyRule{},
			Chats: []*mtproto.Chat{},
			Users: []*mtproto.User{},
		}).To_Account_PrivacyRules(), nil
	} else {
		userIdList := make([]int64, 0, 64)
		chatIdList := make([]int64, 0, 64)
		f := func(u []int64, t int64) []int64 {
			if t == 0 {
				return u
			}
			found := false
			for _, vv := range u {
				found = false
				if vv == t {
					found = true
					break
				}
			}
			if found == false {
				u = append(u, t)
			}
			return u
		}

		fl := func(u []int64, t []int64) []int64 {
			for _, v := range t {
				u = f(u, v)
			}
			return u
		}

		rules := make([]*mtproto.PrivacyRule, 0, len(privacyValue))
		for _, v := range privacyValue {
			switch constants.OptionPrivacy(v.PrivacyKey) {
			case constants.OptionPrivacyAllowAll:
				rules = append(rules, mtproto.NewTLPrivacyValueAllowAll(nil).To_PrivacyRule())
			case constants.OptionPrivacyAllowUsers:
				rules = append(rules, mtproto.NewTLPrivacyValueAllowUsers(&mtproto.PrivacyRule{
					Users: constants.ArrayInt64To32(v.AllowList),
				}).To_PrivacyRule())
				userIdList = fl(userIdList, v.AllowList)
			case constants.OptionPrivacyAllowChatParticipants:
				rules = append(rules, mtproto.NewTLPrivacyValueAllowChatParticipants(&mtproto.PrivacyRule{
					Chats: constants.ArrayInt64To32(v.AllowList),
				}).To_PrivacyRule())
				chatIdList = fl(chatIdList, v.AllowList)
			case constants.OptionPrivacyAllowAllContacts:
				rules = append(rules, mtproto.NewTLPrivacyValueAllowContacts(nil).To_PrivacyRule())
			case constants.OptionPrivacyDisallowAll:
				rules = append(rules, mtproto.NewTLPrivacyValueDisallowAll(nil).To_PrivacyRule())
			case constants.OptionPrivacyDisallowUsers:
				rules = append(rules, mtproto.NewTLPrivacyValueDisallowUsers(&mtproto.PrivacyRule{
					Users: constants.ArrayInt64To32(v.DisallowList),
				}).To_PrivacyRule())
				userIdList = fl(userIdList, v.DisallowList)
			case constants.OptionPrivacyDisallowChatParticipants:
				rules = append(rules, mtproto.NewTLPrivacyValueDisallowChatParticipants(&mtproto.PrivacyRule{
					Chats: constants.ArrayInt64To32(v.DisallowList),
				}).To_PrivacyRule())
				chatIdList = fl(chatIdList, v.DisallowList)
			case constants.OptionPrivacyDisallowContacts:
				rules = append(rules, mtproto.NewTLPrivacyValueDisallowContacts(nil).To_PrivacyRule())
			default:
				fmt.Errorf("OptionPrivacy error:%v", v.PrivacyKey)
			}
		}

		var chatList []*mtproto.Chat
		userList := make([]*mtproto.User, 0, len(userIdList))
		if key == constants.KeyPrivacyUserStatus {

			//TODO:Coder
			//tmpUserList, _ := s.usersCore.UserIdList(md.UserId, userIdList)
			//hideStatus := mtproto.NewTLUserStatusEmpty(nil).To_UserStatus()
			//for _, vv := range tmpUserList {
			//	if !data_privacy.CheckPrivacy(vv.Id, 0, privacyValue, false) {
			//		vv.Status = hideStatus
			//	}
			//
			//	userList = append(userList, vv)
			//}

			//chatList, _ = group.NewChatCore().GetChats(md.UserId, chatIdList, true)
		}

		return mtproto.NewTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
			Rules: rules,
			Chats: chatList,
			Users: userList,
		}).To_Account_PrivacyRules(), nil
	}
}
