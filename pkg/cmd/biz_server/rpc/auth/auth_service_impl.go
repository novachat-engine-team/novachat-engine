/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth_service_impl.go
 * @Desc :
 *
 */

package rpc

import (
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/cmd/biz_server/conf"
	"novachat_engine/pkg/log"
	"novachat_engine/service/account/qr_login"
	accountUsers "novachat_engine/service/account/users"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/core/account/banned_phone_number"
	"novachat_engine/service/core/account/banned_phone_number/banned_memory"
	"novachat_engine/service/core/account/users/sql_core"
)

type AuthServiceImpl struct {
	conf             *conf.Config
	banned           banned_phone_number.BannedPhoneNumber
	accountCore      *account.Core
	usersCore        *sql_core.Core
	accountUsersCore *accountUsers.Core
	qrLoginCore      *qr_login.Core
}

func NewAuthServiceImpl(conf *conf.Config) *AuthServiceImpl {
	var err error
	impl := &AuthServiceImpl{
		conf: conf,
	}

	impl.banned, err = banned_phone_number.InstallBannedPhoneNumber(banned_memory.BannedNameMemory)
	if err != nil {
		log.Fatalf("NewAuthServiceImpl InstallBannedPhoneNumber error:%s", err.Error())
	}
	impl.banned.Initialize()
	impl.accountCore = account.NewAccountCore(cache.GetRedisClient())
	impl.usersCore = sql_core.NewUsersCore(&conf.MySQL)
	impl.accountUsersCore = accountUsers.NewUsersCore(impl.accountCore, nil, impl.usersCore)
	impl.qrLoginCore = qr_login.NewQRLoginCore(conf.QRLoginCodeKey, conf.Redis)
	return impl
}

//
//func (s *AuthServiceImpl) authSignInMulti(authKeyId int64, userId int32, data *types.Any) error {
//
//    info := &mtproto.TLInitConnection{}
//    err := types.UnmarshalAny(data, info)
//    if err != nil {
//        log.Errorf("authSignInMulti userId:%d authKeyId:%d", userId, authKeyId)
//        return err
//    }
//
//    authorizations, err := authClient.GetAuthClient().AuthGetAuthorizations(context.Background(), &auth_pb.AuthGetAuthorizations{
//        UserId:               userId,
//        AuthKeyId:            authKeyId,
//    })
//    if err != nil {
//        log.Errorf("authSignInMulti userId:%d authKeyId:%d AuthGetAuthorizations error:%s", userId, authKeyId, err.Error())
//        return err
//    }
//
//    log.Infof("authSignInMulti userId:%d authKeyId:%d AuthGetAuthorizations:%+v", userId, authKeyId, authorizations.Values)
//    kickAuthorizationList := make([]*auth_pb.VersionInfo, 0, 1)
//    platType := strings.ToLower(info.LangPack)
//    for _, authorization := range authorizations.Values {
//        if authorization.AuthKeyId != authKeyId {
//            if ((strings.Compare(platType, constants.DevicePlatformTypeAndroid.ToString()) == 0 || strings.Compare(platType, constants.DevicePlatformTypeIOS.ToString()) == 0) &&
//                (strings.Compare(strings.ToLower(authorization.LangPack), constants.DevicePlatformTypeAndroid.ToString()) == 0 || strings.Compare(strings.ToLower(authorization.LangPack), constants.DevicePlatformTypeIOS.ToString()) == 0)) ||
//                strings.Compare(strings.ToLower(authorization.LangPack), platType) == 0 {
//
//                kickAuthorizationList = append(kickAuthorizationList, authorization)
//            }
//        }
//    }
//
//    return s.accountCore.ResetAuthorization(userId, kickAuthorizationList)
//}
