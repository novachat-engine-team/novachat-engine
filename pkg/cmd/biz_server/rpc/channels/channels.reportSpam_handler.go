/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : channels.reportSpam_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

const saveMessageLen = 10

//  channels.reportSpam#fe087810 channel:InputChannel user_id:InputUser id:Vector<int> = Bool;
//
func (s *ChannelsServiceImpl) ChannelsReportSpam(ctx context.Context, request *mtproto.TLChannelsReportSpam) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("ChannelsReportSpam %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//reportType, text := report.ToReportType(mtproto.NewTLInputReportReasonSpam(nil).To_ReportReason())
	//var inputPeer *input.InputPeer
	//if request.Channel.ChannelId != 0 {
	//    inputPeer = input.MakeInputPeer(mtproto.NewTLInputPeerChannel(&mtproto.InputPeer{ChannelId: request.Channel.ChannelId}).To_InputPeer())
	//} else {
	//    inputPeer = input.MakeInputPeer(request.Channel.Peer)
	//}
	//if inputPeer.GetPeerType() == constants.PeerTypeEmpty {
	//    return mtproto.ToMTBool(true), nil
	//}
	//messageList := make([]string, 0, len(request.Id))
	//if len(request.Id) > 0 {
	//    messageData := make([]*dbproxy_pb.MessageMessageData, 0, 1)
	//    messageData = append(messageData, &dbproxy_pb.MessageMessageData{
	//        UserId:               0,
	//        MessageIdList:        request.Id,
	//        InputPeer:            inputPeer.ToInputPeer(),
	//    })
	//
	//    messages, _ := dbproxyClient.GetDBProxyClient().MessageGetMessagesByMessageIds(context.Background(), &dbproxy_pb.MessageGetMessagesByMessageIds{
	//        Value:                nil,
	//        UserId:               md.UserId,
	//    })
	//
	//    for _, v := range messages.MessageList {
	//        if len(v.Message) > 0 {
	//            messageList = append(messageList, v.Message)
	//        }
	//    }
	//}
	//
	//if s.conf.External.Enable == true {
	//    peerType := int32(0)
	//    if inputPeer.GetPeerType() == constants.PeerTypeUser || inputPeer.GetPeerType() == constants.PeerTypeSelf {
	//        peerType = 1
	//    } else {
	//        peerType = 2
	//    }
	//
	//    if len(messageList) > saveMessageLen {
	//        messageList = messageList[:saveMessageLen]
	//    }
	//
	//    feedbackBody := &externalClient.RequestReportFeedback{
	//        UserId: constants.PeerTypeFromUserIDType(md.UserId).ToInt(),
	//        PeerId: inputPeer.GetPeerId(),
	//        PeerType: peerType,
	//        Info: text,
	//        ReportType: reportType.ToInt32(),
	//        Message: strings.Join(messageList, "\n"),
	//    }
	//    bodyBuf, _ := jsoniter.Marshal(feedbackBody)
	//    md5Hash := md5.New()
	//    md5Hash.Write(bodyBuf)
	//    feedbackBody.Hash = hex.EncodeToString(md5Hash.Sum(nil))
	//
	//    body, _ := jsoniter.MarshalToString(feedbackBody)
	//    log.Debugf("ChannelsReportSpam userId:%d req:%v", md.UserId, body)
	//    tlInitConnection := mtproto.NewTLInitConnection()
	//    err := types.UnmarshalAny(md.ExtendData, tlInitConnection)
	//    if err != nil {
	//        _ = err
	//    }
	//    header := &external_pb.RequestHeader{
	//        OriginAuthKeyId:      md.OriginAuthKeyId,
	//        AuthKeyId:            md.AuthKeyId,
	//        SessionId:            md.SessionId,
	//        UserId:               constants.PeerTypeFromUserIDType(md.UserId).ToInt(),
	//        ServerPeer:           md.ServerPeer,
	//        Layer:                md.Layer,
	//        Mac:                  md.Mac,
	//        Ip:                   md.Ip,
	//        ApiId:                tlInitConnection.ApiId,
	//        DeviceModel:          tlInitConnection.DeviceModel,
	//        SystemVersion:        tlInitConnection.SystemVersion,
	//        AppVersion:           tlInitConnection.AppVersion,
	//        SystemLangCode:       tlInitConnection.SystemLangCode,
	//        LangPack:             tlInitConnection.LangPack,
	//        LangCode:             tlInitConnection.LangCode,
	//    }
	//    reportNotify := &external_pb.ExternalRequest{
	//        Header: header,
	//        Body: &external_pb.RequestBody{
	//            Method: externalClient.ReportFeedback,
	//            Body:   body,
	//        },
	//    }
	//
	//    body, _ = jsoniter.MarshalToString(reportNotify)
	//    err = externalClient.GetExternalMQClient().SendMessage(reportNotify.Body.Method, body)
	//    if err != nil {
	//        log.Fatalf("ChannelsReportSpam userId:%d body:`%s` error:%s", md.UserId, body, err.Error())
	//    }
	//}

	return mtproto.ToMTBool(true), nil
}
