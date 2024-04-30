/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : account.reportPeer_handler.go
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

//  account.reportPeer#ae189d5f peer:InputPeer reason:ReportReason = Bool;
//
func (s *AccountServiceImpl) AccountReportPeer(ctx context.Context, request *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AccountReportPeer %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	//reportType, text := report.ToReportType(request.Reason)
	//inputPeer := input.MakeInputPeer(request.Peer)
	//if inputPeer.GetPeerType() == constants.PeerTypeEmpty {
	//    return mtproto.ToMTBool(true), nil
	//}
	//if s.conf.External.Enable == true {
	//    peerType := int32(0)
	//    if inputPeer.GetPeerType() == constants.PeerTypeUser || inputPeer.GetPeerType() == constants.PeerTypeSelf {
	//        peerType = 1
	//    } else {
	//        peerType = 2
	//    }
	//
	//    feedbackBody := &externalClient.RequestReportFeedback{
	//        UserId: constants.PeerTypeFromUserIDType(md.UserId).ToInt(),
	//        PeerId: inputPeer.GetPeerId(),
	//        PeerType: peerType,
	//        Info: text,
	//        ReportType: reportType.ToInt32(),
	//        Message: "",
	//    }
	//    bodyBuf, _ := jsoniter.Marshal(feedbackBody)
	//    md5Hash := md5.New()
	//    md5Hash.Write(bodyBuf)
	//    feedbackBody.Hash = hex.EncodeToString(md5Hash.Sum(nil))
	//
	//    body, _ := jsoniter.MarshalToString(feedbackBody)
	//    log.Debugf("AccountReportPeer userId:%d req:%v", md.UserId, body)
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
	//        log.Fatalf("AccountReportPeer userId:%d body:`%s` error:%s", md.UserId, body, err.Error())
	//    }
	//}

	return mtproto.ToMTBool(true), nil
}
