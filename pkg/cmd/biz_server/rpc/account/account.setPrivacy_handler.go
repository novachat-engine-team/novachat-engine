/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.setPrivacy_handler.go
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
	data_privacy "novachat_engine/service/data/privacy"
)

//  account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
//
func (s *AccountServiceImpl) AccountSetPrivacy(ctx context.Context, request *mtproto.TLAccountSetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountSetPrivacy %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//var err error
	keyPrivacy := constants.ToPrivacyKey(request.GetKey())
	updateRules := make([]*mtproto.PrivacyRule, 0, len(request.GetRules()))
	if keyPrivacy == constants.KeyPrivacyNone {
		return mtproto.NewTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
			Rules: updateRules,
			Chats: []*mtproto.Chat{},
			Users: []*mtproto.User{},
		}).To_Account_PrivacyRules(), nil
	}

	updatePrivacy := make([]*data_privacy.PrivacyOption, 0, 4)
	for _, r := range request.GetRules() {
		value := constants.ToPrivacyValue(r)
		updateRules = append(updateRules, constants.ToPrivacyValueMTValue(md.UserId, r))

		v := constants.ToPrivacyValueMTValue(md.UserId, r)
		if v == nil {
			continue
		}

		switch value {
		case constants.OptionPrivacyAllowAll:
			updatePrivacy = append(updatePrivacy, &data_privacy.PrivacyOption{
				PrivacyKey:   value.Int32(),
				AllowList:    []int64{},
				DisallowList: []int64{},
			})
		case constants.OptionPrivacyAllowUsers:
			updatePrivacy = append(updatePrivacy, &data_privacy.PrivacyOption{
				PrivacyKey:   value.Int32(),
				AllowList:    constants.ArrayInt32To64(v.Users),
				DisallowList: []int64{},
			})
		case constants.OptionPrivacyAllowChatParticipants:
			updatePrivacy = append(updatePrivacy, &data_privacy.PrivacyOption{
				PrivacyKey:   value.Int32(),
				AllowList:    constants.ArrayInt32To64(v.Chats),
				DisallowList: []int64{},
			})
		case constants.OptionPrivacyAllowAllContacts:
			updatePrivacy = append(updatePrivacy, &data_privacy.PrivacyOption{
				PrivacyKey:   value.Int32(),
				AllowList:    []int64{},
				DisallowList: []int64{},
			})
		case constants.OptionPrivacyDisallowAll:
			updatePrivacy = append(updatePrivacy, &data_privacy.PrivacyOption{
				PrivacyKey:   value.Int32(),
				AllowList:    []int64{},
				DisallowList: []int64{},
			})
		case constants.OptionPrivacyDisallowUsers:
			updatePrivacy = append(updatePrivacy, &data_privacy.PrivacyOption{
				PrivacyKey:   value.Int32(),
				AllowList:    []int64{},
				DisallowList: constants.ArrayInt32To64(v.Users),
			})
		case constants.OptionPrivacyDisallowChatParticipants:
			updatePrivacy = append(updatePrivacy, &data_privacy.PrivacyOption{
				PrivacyKey:   value.Int32(),
				AllowList:    []int64{},
				DisallowList: constants.ArrayInt32To64(v.Chats),
			})
		case constants.OptionPrivacyDisallowContacts:
			updatePrivacy = append(updatePrivacy, &data_privacy.PrivacyOption{
				PrivacyKey:   value.Int32(),
				AllowList:    []int64{},
				DisallowList: []int64{},
			})
		default:
			fmt.Errorf("OptionPrivacy error:%v", value)
		}
	}

	var err error
	_, err = s.privacyCore.UpdatePrivacy(md.UserId, keyPrivacy.Int32(), updatePrivacy)
	if err != nil {
		log.Errorf("AccountSetPrivacy - request: %v error:%s", request, err.Error())
		return nil, err
	}

	//TODO:Coder updates
	//go func() {
	//	_, err = syncClient.GetSyncClient().PushUpdates(context.Background(), &sync_pb.PushUpdates{
	//		UserId:               md.UserId,
	//		ExcludeAuthKeyIdList: []int64{md.AuthKeyId},
	//		Updates: mtproto.NewTLUpdates(&mtproto.Updates{
	//			Users: []*mtproto.User{},
	//			Chats: []*mtproto.Chat{},
	//			Updates: []*mtproto.Update{
	//				mtproto.NewTLUpdatePrivacy(&mtproto.Update{
	//					Key:   constants.ToPrivacyKeyMTKey(request.GetKey()),
	//					Rules: updateRules,
	//				}).To_Update()},
	//			Seq:  0,
	//			Date: int32(time.Now().Unix()),
	//		}).To_Updates(),
	//	})
	//	if err != nil {
	//		log.Errorf("AccountSetPrivacy - request: %v error:%s", request, err.Error())
	//	}
	//}()

	return mtproto.NewTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
		Rules: updateRules,
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	}).To_Account_PrivacyRules(), nil
}
