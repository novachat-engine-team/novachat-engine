/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getAppChangelog_handler.go
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

//  help.getAppChangelog#9010ef6f prev_app_version:string = Updates;
//  help.getAppChangelog#5bab7fb2 device_model:string system_version:string app_version:string lang_code:string = help.AppChangelog;
//
func (s *HelpServiceImpl) HelpGetAppChangelog(ctx context.Context, request *mtproto.TLHelpGetAppChangelog) (*mtproto.Response_HelpGetAppChangelog, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetAppChangelog %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//appChangeLog := ""
	//updates := make([]*mtproto.Update, 0, 1)

	//notifications, err := service_notification.GetNotification()
	//if err != nil {
	//    log.Fatalf("HelpGetAppChangelog - GetNotification error:%s", err.Error())
	//}

	//now := int32(time.Now().Unix())
	//for _, v := range notifications {
	//    if now > v.ValidTime && v.ValidTime != 0 {
	//        continue
	//    }
	//
	//    updates = append(updates, service_notification.Notification(v.PopUp,
	//        v.Message,
	//        v.Media,
	//        v.NotifyType,
	//        v.MessageEntity))
	//}
	//
	//ret := &mtproto.Response_HelpGetAppChangelog{
	//	HelpGetAppChangelog9010Ef6F: mtproto.NewTLUpdates(&mtproto.Updates{
	//		Updates: updates,
	//		Users:   []*mtproto.User{},
	//		Chats:   []*mtproto.Chat{},
	//		Seq:     0,
	//		Date:    now,
	//	}).To_Updates(),
	//	HelpGetAppChangelog5Bab7Fb2: mtproto.NewTLHelpAppChangelog(&mtproto.Help_AppChangelog{
	//		Text: appChangeLog,
	//	}).To_Help_AppChangelog(),
	//}
	//
	//if md.Layer <= 51 {
	//	ret.Cmd = mtproto.Cmd_HelpGetAppChangelog5bab7fb2.ToUInt32()
	//} else {
	//	ret.Cmd = mtproto.Cmd_HelpGetAppChangelog.ToUInt32()
	//}
	//return ret, nil

	return nil, fmt.Errorf("%s", "Not impl HelpGetAppChangelog")
}
