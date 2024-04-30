/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/5/26 14:42
 * @File : phone_call.go
 */

package phone_call

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cache"
	relayClient "novachat_engine/pkg/cmd/relay/rpc_client"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/account/phone"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"strings"
	"time"
)

const p2p = true
const version = "2.4.4"

var supportLibVersion = []string{"2.4.4", "2.8.8", "2.7.7", "3.0.0"}

func checkVersion(libVersion string) string {
	for _, v := range supportLibVersion {
		if strings.Compare(v, libVersion) == 0 {
			return libVersion
		}
	}

	return version
}

func MakeConnection(id int64, ip string, port int32, ipv6 string, peerTag []byte) *mtproto.PhoneConnection {
	return mtproto.NewTLPhoneConnection(&mtproto.PhoneConnection{
		Id:      id,
		Ip:      ip,
		Ipv6:    ipv6,
		Port:    port,
		PeerTag: peerTag,
	}).To_PhoneConnection()
}

type PhoneCallCore struct {
	useWebRTC  bool
	p2p        bool
	libVersion string
	node       *snowflake.Node
	phoneCall  *phone.PhoneCall
}

func NewPhoneCallCore(useWebrtc bool, p2p bool, libVersion string, serverNode int32, r *cache.RedisClient) *PhoneCallCore {
	impl := &PhoneCallCore{
		useWebRTC:  useWebrtc,
		phoneCall:  phone.NewPhoneCall(r),
		p2p:        p2p,
		libVersion: libVersion,
	}

	impl.libVersion = checkVersion(libVersion)

	var err error
	impl.node, err = snowflake.NewNode(int64(serverNode))
	if err != nil {
		panic(err)
	}
	return impl
}

func (m *PhoneCallCore) CheckComingPhoneCallAndUpdates(userId int64) *mtproto.Updates {
	phoneCallInfo := m.phoneCall.CheckComingPhoneCall(userId)
	if phoneCallInfo != nil && phoneCallInfo.ID != 0 {
		updates := mtproto.NewTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{mtproto.NewTLUpdatePhoneCall(&mtproto.Update{
				PhoneCall: mtproto.NewTLPhoneCallRequested(&mtproto.PhoneCall{
					Video:         phoneCallInfo.Video,
					Id:            phoneCallInfo.ID,
					AccessHash:    phoneCallInfo.AccessHash,
					Date:          phoneCallInfo.Date,
					AdminId:       constants.PeerTypeFromUserIDType(phoneCallInfo.Admin).ToInt32(),
					ParticipantId: constants.PeerTypeFromUserIDType(phoneCallInfo.ParticipantId).ToInt32(),
					GAHash:        phoneCallInfo.ToGAHash(),
					Protocol:      phoneCallInfo.Protocol,
				}).To_PhoneCall(),
			}).To_Update()},
			Users: []*mtproto.User{},
			Date:  int32(time.Now().Unix()),
			Seq:   0,
		})
		return updates.To_Updates()
	}
	return nil
}

func (m *PhoneCallCore) ReceivedPhoneCall(userId int64, phonePeerId int64) (bool, error) {

	info, err := m.phoneCall.TokenPhoneCall(phonePeerId)
	if err != nil {
		log.Errorf("ReceivedPhoneCall TokenPhoneCall userId:%d phoneCallId:%d error:%s", userId, phonePeerId, err.Error())
		return false, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED, err.Error())
	}

	if info.ID == 0 {
		log.Warnf("ReceivedPhoneCall TokenPhoneCall userId:%d phoneCallId:%d not found", userId, phonePeerId)
		return false, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED, "")
	}

	log.Infof("ReceivedPhoneCall userId:%d peerType:%d peerId:%d", userId, phonePeerId)

	//TODO:Coder
	// sync
	//go func() {
	//	updates := mtproto.NewTLUpdates(&mtproto.Updates{
	//		Updates: []*mtproto.Update{mtproto.NewTLUpdatePhoneCall(&mtproto.Update{
	//			PhoneCall: mtproto.NewTLPhoneCallWaiting(&mtproto.PhoneCall{
	//				Id:            info.ID,
	//				AccessHash:    info.AccessHash,
	//				Video:         info.Video,
	//				Date:          int32(time.Now().Unix()),
	//				AdminId:       info.Admin,
	//				ParticipantId: info.ParticipantId,
	//				Protocol:      info.Protocol,
	//				ReceiveDate:   info.Date,
	//			}).To_PhoneCall(),
	//		}).To_Update()},
	//		Users: []*mtproto.User{},
	//		Date:  int32(time.Now().Unix()),
	//		Seq:   0,
	//	})
	//	_, err = syncClient.GetSyncClient().SyncUpdates(context.Background(), &sync_pb.SyncUpdates{
	//		UserId:  info.Admin,
	//		Updates: updates.To_Updates(),
	//	})
	//	if err != nil {
	//		log.Errorf("syncClient.GetSyncClient() error:%s", err.Error())
	//	}
	//}()

	return true, nil
}

func (m *PhoneCallCore) RequestPhoneCall(authKeyId int64, userId int64, peerUserId int64, video bool, randomId int32, hash []byte, protocol *mtproto.PhoneCallProtocol) (*mtproto.TLPhonePhoneCall, error) {
	phoneCall := NewPhoneCall(userId, peerUserId, video, randomId, protocol)
	reply, err := phoneCall.CreatePhoneCall(m.node)
	if err != nil {
		log.Errorf("RequestPhoneCall userId:%d peerId:%d randomId:%d error:%s", userId, peerUserId, randomId, err.Error())
		return nil, err
	}
	protocol.LibraryVersions = []string{m.libVersion}
	phoneCallReply := reply.GetData2().GetPhoneCall()
	phoneCallInfo := phone.PhoneCallInfo{
		ID:            phoneCallReply.GetId(),
		Status:        constants.PhoneCallStatusRequest,
		Type:          phone.PhoneCallTypeUser,
		Date:          phoneCallReply.GetReceiveDate(),
		Admin:         constants.PeerTypeFromUserIDType(userId).ToInt(),
		ParticipantId: constants.PeerTypeFromUserIDType(peerUserId).ToInt(),
		AccessHash:    phoneCallReply.AccessHash,
		Video:         video,
		RandomId:      randomId,
		Protocol:      protocol,
		P2pAllowed:    m.p2p,
		AuthKeyId:     authKeyId,
	}

	phoneCallInfo.FromGAHash(hash)
	_, err = m.phoneCall.PutPhoneCall(&phoneCallInfo)
	if err != nil {
		log.Errorf("RequestPhoneCall userId:%d peerId:%d randomId:%d PutPhoneCall error:%s", userId, peerUserId, randomId, err.Error())
		return nil, err
	}

	log.Infof("RequestPhoneCall userId:%d peerId:%d phoneCallId:%d randomId:%d", userId, peerUserId, phoneCallReply.Id, randomId)

	//TODO:Coder
	// sync
	//go func() {
	//	// phoneCallRequested#87eabb53 flags:# video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;
	//	// phoneCallRequested#83761ce4 id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;
	//	//userInfo, _ := dbproxyClient.GetDBProxyClient().UserByUserIdList(context.Background(), &dbproxy_pb.UserByUserIdList{
	//	//	OwnerUserId:          inputUser.GetId(),
	//	//	UserIdList:           []int32{userId, inputUser.GetId()},
	//	//  Now: int32(time.Now().Unix()),
	//	//})
	//
	//	updates := mtproto.NewTLUpdates(&mtproto.Updates{
	//		Updates: []*mtproto.Update{mtproto.NewTLUpdatePhoneCall(&mtproto.Update{
	//			PhoneCall: mtproto.NewTLPhoneCallRequested(&mtproto.PhoneCall{
	//				Video:         video,
	//				Id:            phoneCallReply.GetId(),
	//				AccessHash:    phoneCallReply.GetAccessHash(),
	//				Date:          phoneCallReply.GetDate(),
	//				AdminId:       userId,
	//				ParticipantId: inputUser.GetId(),
	//				GAHash:        hash,
	//				Protocol:      protocol,
	//			}).To_PhoneCall(),
	//		}).To_Update()},
	//		Users: []*mtproto.User{},
	//		Date:  int32(time.Now().Unix()),
	//		Seq:   0,
	//	})
	//
	//	_, err = syncClient.GetSyncClient().PushUpdates(context.Background(), &sync_pb.PushUpdates{
	//		UserId:  inputUser.GetId(),
	//		Updates: updates.To_Updates(),
	//	})
	//	if err != nil {
	//		log.Errorf("syncClient.GetSyncClient() error:%s", err.Error())
	//	}
	//}()

	return reply, err
}

func (m *PhoneCallCore) toConnectPhoneCall(userId int64, info *phone.PhoneCallInfo) *mtproto.TLPhoneCall {
	// phoneCall#8742ae7f flags:# p2p_allowed:flags.5?true video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connections:Vector<PhoneConnection> start_date:int = PhoneCall;
	// phoneCall#ffe6ab67 id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;
	phoneCall := mtproto.NewTLPhoneCall(&mtproto.PhoneCall{
		P2PAllowed:    info.P2pAllowed,
		Video:         info.Video,
		Id:            info.ID,
		AccessHash:    info.AccessHash,
		Date:          info.Date,
		AdminId:       constants.PeerTypeFromUserIDType(info.Admin).ToInt32(),
		ParticipantId: constants.PeerTypeFromUserIDType(info.ParticipantId).ToInt32(),
	})
	if userId == info.Admin {
		//phoneCall.SetGAOrB(info.ToGAHash())
		phoneCall.SetGAOrB(info.ToGB())
	} else {
		//phoneCall.SetGAOrB(info.ToGB())
		phoneCall.SetGAOrB(info.ToGAHash())
	}
	return phoneCall
}

func (m *PhoneCallCore) ConfirmPhoneCall(userId int64, authKeyId int64, phonePeerId int64, ga []byte, fingerprint int64, protocol *mtproto.PhoneCallProtocol) (*mtproto.TLPhonePhoneCall, error) {

	info, err := m.phoneCall.TokenPhoneCall(phonePeerId)
	if err != nil {
		log.Errorf("ConfirmPhoneCall TokenPhoneCall userId:%d phoneCallId:%d error:%s", userId, phonePeerId, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED, err.Error())
	}

	if info.ID == 0 {
		log.Warnf("ConfirmPhoneCall TokenPhoneCall userId:%d phoneCallId:%d error not found", userId, phonePeerId)
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED, "")
	}

	info.FromGAHash(ga)

	protocol.UdpP2P = m.p2p
	if protocol.MinLayer > info.Protocol.MinLayer {
		protocol.MinLayer = info.Protocol.MinLayer
	}
	if protocol.MaxLayer < info.Protocol.MaxLayer {
		protocol.MaxLayer = info.Protocol.MaxLayer
	}
	//protocol.LibraryVersions = append(protocol.LibraryVersions, info.Protocol.LibraryVersions...)
	protocol.LibraryVersions = []string{m.libVersion}

	relayRpc := relayClient.GetRelayClientByKey(fmt.Sprintf("%d", info.Admin))
	if relayRpc == nil {
		s := fmt.Sprintf("ConfirmPhoneCall userId:%d phoneCallId:%d participantId:%d GetRelayClientByKey error!!!!!!!!!!!!", info.Admin, info.ID, info.ParticipantId)
		log.Error(s)
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED, s)
	}

	ok, err := m.phoneCall.ConfirmPhoneCall(info.ID, info.ParticipantId)
	if err != nil || ok == false {
		if err != nil {
			log.Errorf("AcceptPhoneCall UserByUserIdList userId:%d ConfirmPhoneCall phoneCallId:%d error:%s", userId, phonePeerId, err.Error())
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED)
		} else {
			log.Errorf("AcceptPhoneCall UserByUserIdList userId:%d ConfirmPhoneCall phoneCallId:%d timeout:%d!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", userId, phonePeerId, phone.PhoneCallTimeout)
			return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED)
		}
	}

	callConnection, err := relayRpc.ReqCreatePhoneCall(context.Background(), &relayClient.CreatePhoneCall{
		Id:            info.ID,
		AdminId:       info.Admin,
		ParticipantId: info.ParticipantId,
	})
	if err != nil {
		log.Errorf("ConfirmPhoneCall CreatePhoneCall userId:%d phoneCallId:%d participantId:%d error:%s", info.Admin, info.ID, info.ParticipantId, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED)
	}

	log.Infof("ConfirmPhoneCall userId:%d phoneCallId:%d info:%+v", userId, phonePeerId, info)

	now := int32(time.Now().Unix())
	//TODO:Coder
	// sync
	//go func() {
	//	phoneCall := m.toConnectPhoneCall(info.ParticipantId, info)
	//	phoneCall.SetConnection(callConnection.Connection)
	//	phoneCall.SetAlternativeConnections(callConnection.AlternativeConnections)
	//	phoneCall.SetStartDate(now)
	//	phoneCall.SetProtocol(protocol)
	//	phoneCall.SetKeyFingerprint(fingerprint)
	//	phoneCall.SetConnections(callConnection.ConnectionList)
	//
	//	//userInfo, _ := dbproxyClient.GetDBProxyClient().UserByUserIdList(context.Background(), &dbproxy_pb.UserByUserIdList{
	//	//	OwnerUserId:          info.ParticipantId,
	//	//	UserIdList:           []int32{info.ParticipantId, info.Admin},
	//	//  Now: int32(time.Now().Unix()),
	//	//})
	//
	//	Reason := mtproto.NewTLPhoneCallDiscardReasonBusy(nil)
	//	phoneCallDiscarded := mtproto.NewTLPhoneCallDiscarded(&mtproto.PhoneCall{
	//		Id:        info.ID,
	//		NeedDebug: true,
	//		Reason:    Reason.To_PhoneCallDiscardReason(),
	//	})
	//	_ = phoneCallDiscarded
	//
	//	_, err = syncClient.GetSyncClient().SyncUpdatesList(context.Background(), &sync_pb.SyncUpdatesList{
	//		SyncUpdatesList: []*sync_pb.SyncUpdates{
	//			{
	//				UserId:    info.ParticipantId,
	//				AuthKeyId: info.AcceptAuthKeyId,
	//				Updates: mtproto.NewTLUpdates(&mtproto.Updates{
	//					Updates: []*mtproto.Update{mtproto.NewTLUpdatePhoneCall(&mtproto.Update{
	//						PhoneCall: phoneCall.To_PhoneCall(),
	//					}).To_Update(),
	//					},
	//					Users: []*mtproto.User{},
	//					Date:  int32(time.Now().Unix()),
	//					Seq:   0,
	//				}).To_Updates(),
	//			},
	//			{
	//				UserId:               info.ParticipantId,
	//				ExcludeAuthKeyIdList: []int64{info.AcceptAuthKeyId},
	//				Updates: mtproto.NewTLUpdates(&mtproto.Updates{
	//					Updates: []*mtproto.Update{
	//						mtproto.NewTLUpdatePhoneCall(&mtproto.Update{
	//							PhoneCall: phoneCallDiscarded.To_PhoneCall(),
	//						}).To_Update(),
	//					},
	//					Users: []*mtproto.User{},
	//					Date:  int32(time.Now().Unix()),
	//					Seq:   0,
	//				}).To_Updates(),
	//			},
	//		},
	//	})
	//	if err != nil {
	//		log.Errorf("syncClient.GetSyncClient() error:%s", err.Error())
	//	}
	//}()

	//  phoneCall#8742ae7f flags:# p2p_allowed:flags.5?true video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connections:Vector<PhoneConnection> start_date:int = PhoneCall;
	//  phoneCall#ffe6ab67 id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;

	connection := MakeConnection(info.ID, callConnection.Connect.Ip, callConnection.Connect.Port, callConnection.Connect.Ipv6, callConnection.Connect.PeerTag)
	phoneCall := m.toConnectPhoneCall(userId, info)
	phoneCall.SetConnection(connection)
	phoneCall.SetAlternativeConnections(nil)
	phoneCall.SetStartDate(now)
	phoneCall.SetProtocol(protocol)
	phoneCall.SetKeyFingerprint(fingerprint)
	phoneCall.SetConnections([]*mtproto.PhoneConnection{connection})

	phoneCallReply := mtproto.NewTLPhonePhoneCall(&mtproto.Phone_PhoneCall{
		PhoneCall: phoneCall.To_PhoneCall(),
		//Users:     userInfo.UserList,
	})

	log.Infof("ConfirmPhoneCall userId:%d phoneCallId:%d info:%+v", userId, phonePeerId, phoneCall)
	return phoneCallReply, nil
}

func (m *PhoneCallCore) AcceptPhoneCall(userId int64, authKeyId int64, phonePeerId int64, gb []byte, protocol *mtproto.PhoneCallProtocol) (*mtproto.TLPhonePhoneCall, error) {
	ok, err := m.phoneCall.CheckAcceptedPhoneCall(userId, authKeyId, phonePeerId)
	if err != nil {
		log.Errorf("AcceptPhoneCall TokenPhoneCall userId:%d phoneCallId:%d error:%s", userId, phonePeerId, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED, err.Error())
	}

	if ok == false {
		log.Errorf("AcceptPhoneCall TokenPhoneCall userId:%d phoneCallId:%d accepted", userId, phonePeerId)
		return nil, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_ACCEPTED)
	}

	info, err := m.phoneCall.TokenPhoneCall(phonePeerId)
	if err != nil {
		log.Errorf("AcceptPhoneCall TokenPhoneCall userId:%d phoneCallId:%d error:%s", userId, phonePeerId, err.Error())
		return nil, err
	}

	if info.ID == 0 {
		log.Warnf("AcceptPhoneCall TokenPhoneCall userId:%d phoneCallId:%d error not found", userId, phonePeerId)
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED, "")
	}

	protocol.UdpP2P = m.p2p
	if protocol.MinLayer > info.Protocol.MinLayer {
		protocol.MinLayer = info.Protocol.MinLayer
	}
	if protocol.MaxLayer < info.Protocol.MaxLayer {
		protocol.MaxLayer = info.Protocol.MaxLayer
	}
	//protocol.LibraryVersions = append(protocol.LibraryVersions, info.Protocol.LibraryVersions...)
	protocol.LibraryVersions = []string{m.libVersion}
	info.FromGB(gb)
	info.AcceptAuthKeyId = authKeyId
	info.Protocol = protocol // participantId protocol
	_, err = m.phoneCall.PutPhoneCall(info)

	log.Infof("AcceptPhoneCall userId:%d phoneCallId:%d participantId:%d", userId, phonePeerId, info.ParticipantId)

	var phoneCall *mtproto.PhoneCall
	if info.ID != phonePeerId {
		phoneCall = mtproto.NewTLPhoneCallDiscarded(&mtproto.PhoneCall{
			Id:        phonePeerId,
			NeedDebug: true,
			Reason:    mtproto.NewTLPhoneCallDiscardReasonMissed(&mtproto.PhoneCallDiscardReason{}).To_PhoneCallDiscardReason(),
			Duration:  phone.PhoneCallTimeout,
		}).To_PhoneCall()
	} else {
		// phoneCallWaiting#1b8f4ad1 flags:# video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int protocol:PhoneCallProtocol receive_date:flags.0?int = PhoneCall;
		phoneCall = mtproto.NewTLPhoneCallWaiting(&mtproto.PhoneCall{
			Id:            info.ID,
			AccessHash:    info.AccessHash,
			Video:         info.Video,
			Date:          int32(time.Now().Unix()),
			AdminId:       constants.PeerTypeFromUserIDType(info.Admin).ToInt32(),
			ParticipantId: constants.PeerTypeFromUserIDType(info.ParticipantId).ToInt32(),
			Protocol:      info.Protocol,
			ReceiveDate:   info.Date,
		}).To_PhoneCall()
		//go func() {
		//	updatePhoneCall := mtproto.NewTLUpdatePhoneCall(&mtproto.Update{
		//		PhoneCall: mtproto.NewTLPhoneCallAccepted(&mtproto.PhoneCall{
		//			Id:            info.ID,
		//			AccessHash:    info.AccessHash,
		//			Date:          info.Date,
		//			AdminId:       constants.PeerTypeFromUserIDType(info.Admin).ToInt32(),
		//			ParticipantId: constants.PeerTypeFromUserIDType(info.ParticipantId).ToInt32(),
		//			GB:            gb,
		//			Protocol:      protocol,
		//		}).To_PhoneCall(),
		//	})
		//
		//	//userInfo, _ := dbproxyClient.GetDBProxyClient().UserByUserIdList(context.Background(), &dbproxy_pb.UserByUserIdList{
		//	//	OwnerUserId:          userId,
		//	//	UserIdList:           []int32{info.Admin, info.ParticipantId},
		//	//  Now: int32(time.Now().Unix()),
		//	//})
		//	syncClient.GetSyncClient().PushUpdates(context.Background(), &sync_pb.PushUpdates{
		//		UserId:    constants.PeerTypeFromUserIDType(info.Admin).ToInt32(),
		//		AuthKeyId: info.AuthKeyId,
		//		Updates: mtproto.NewTLUpdates(&mtproto.Updates{
		//			Updates: []*mtproto.Update{
		//				updatePhoneCall.To_Update(),
		//			},
		//			Users: []*mtproto.User{},
		//			Date:  int32(time.Now().Unix()),
		//			Seq:   0,
		//		}).To_Updates(),
		//	})
		//}()
	}

	phonePhoneCall := mtproto.NewTLPhonePhoneCall(&mtproto.Phone_PhoneCall{
		PhoneCall: phoneCall,
		//Users:     userInfo.UserList,
	})
	return phonePhoneCall, nil
}

func (m *PhoneCallCore) DiscardCall(md *metadata.RpcMetaData, userId int64, authKeyId int64, phonePeerId int64, reason *mtproto.PhoneCallDiscardReason, duration int32, video bool, connectionId int64) (*mtproto.TLUpdates, error) {

	info, err := m.phoneCall.TokenPhoneCall(phonePeerId)
	if err != nil {
		log.Errorf("DiscardCall userId:%d phoneCallId:%d error:%s", userId, phonePeerId, err.Error())
		return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED, err.Error())
	}

	if info.ID == 0 {
		return mtproto.NewTLUpdates(nil), nil
	}

	callId := info.ID
	if info.ID != phonePeerId {
		callId = phonePeerId
	}

	_ = callId
	//var toAuthKeyId int64
	var toId int64
	if userId == info.Admin {
		toId = info.ParticipantId
		//toAuthKeyId = info.AcceptAuthKeyId
	} else {
		toId = info.Admin
		//toAuthKeyId = info.AuthKeyId
	}

	var ok *types.Any
	relayRpc := relayClient.GetRelayClientByKey(fmt.Sprintf("%d", info.Admin))
	if relayRpc != nil {
		ok, err = relayRpc.ReqDiscardPhoneCall(context.Background(), &relayClient.DiscardPhoneCall{
			Id: phonePeerId,
		})
	}

	//if mtproto.ToBool(ok) == true {
	//	_ = ok
	//} else {
	_ = ok
	//}

	//userInfo, err := dbproxyClient.GetDBProxyClient().UserByUserIdList(context.Background(), &dbproxy_pb.UserByUserIdList{
	//	OwnerUserId:          userId,
	//	UserIdList:           []int32{info.Admin, info.ParticipantId},
	//  Now: int32(time.Now().Unix()),
	//})
	//if err != nil {
	//	log.Errorf("DiscardCall userId:%d phoneCallId:%d UserByUserIdList error:%s", userId, peer.GetId(), err.Error())
	//	return nil, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST, err.Error())
	//}

	log.Infof("DiscardCall userId:%d phoneCallId:%d participantId:%d", userId, phonePeerId, info.ParticipantId)
	m.phoneCall.DiscardPhoneCall(info.ID, info.ParticipantId)
	// add message
	if info.Admin > 0 && info.ParticipantId > 0 {
		// messageActionPhoneCall#80e11a7f flags:# video:flags.2?true call_id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = MessageAction;
		//message := mtproto.NewTLMessageService(&mtproto.Message{
		//	Out:               true,
		//	Date:              int32(time.Now().Unix()),
		//	FromId286FA604119: mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(info.Admin).ToInt32()}).To_Peer(),
		//	FromId90DDDC1171:  constants.PeerTypeFromUserIDType(info.Admin).ToInt32(),
		//	ToId:              mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(info.ParticipantId).ToInt32()}).To_Peer(),
		//	PeerId:            mtproto.NewTLPeerUser(&mtproto.Peer{UserId: constants.PeerTypeFromUserIDType(info.ParticipantId).ToInt32()}).To_Peer(),
		//	Action: mtproto.NewTLMessageActionPhoneCall(&mtproto.MessageAction{
		//		Video:    video,
		//		CallId:   callId,
		//		Reason:   reason,
		//		Duration: duration,
		//	}).To_MessageAction(),
		//})

		//TODO: Coder
		// msg
		//md1 := proto.Clone(md).(*mtproto.RpcMetaData)
		//md1.UserId = info.Admin
		//md1.AuthKeyId = info.AuthKeyId
		//ctx := metadata.NewContextWithMetaData(context.Background(), md1)
		//msgClient.GetMsgClient().SendUserMessage(ctx, &msg_pb.SendUserMessage{
		//	Data: &msg_pb.SendMessage{
		//		AuthKeyId:  info.AuthKeyId,
		//		FromUserId: info.Admin,
		//		Peer:       mtproto.NewTLInputPeerUser(&mtproto.InputPeer{UserId: info.ParticipantId}).To_InputPeer(),
		//		Data: &msg_pb.SendMessageData{
		//			RandomId: rand.Int63(),
		//			Message:  message.To_Message(),
		//		},
		//	},
		//})
	}

	if toId > 0 {
		//go func() {
		//	//userInfo, _ := dbproxyClient.GetDBProxyClient().UserByUserIdList(context.Background(), &dbproxy_pb.UserByUserIdList{
		//	//	OwnerUserId:          toId,
		//	//	UserIdList:           []int32{info.Admin, info.ParticipantId},
		//	//  Now: int32(time.Now().Unix()),
		//	//})
		//	syncClient.GetSyncClient().PushUpdates(context.Background(), &sync_pb.PushUpdates{
		//		UserId:    toId,
		//		AuthKeyId: toAuthKeyId,
		//		Updates: mtproto.NewTLUpdates(&mtproto.Updates{
		//			Updates: []*mtproto.Update{
		//				mtproto.NewTLUpdatePhoneCall(&mtproto.Update{
		//					PhoneCall: mtproto.NewTLPhoneCallDiscarded(&mtproto.PhoneCall{
		//						Id:        peer.Id,
		//						NeedDebug: true,
		//						Reason:    reason,
		//						Duration:  duration,
		//					}).To_PhoneCall(),
		//				}).To_Update(),
		//			},
		//			Users: []*mtproto.User{},
		//			Date:  int32(time.Now().Unix()),
		//			Seq:   0,
		//		}).To_Updates(),
		//	})
		//}()
	}

	return mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{
			mtproto.NewTLUpdatePhoneCall(&mtproto.Update{
				PhoneCall: mtproto.NewTLPhoneCallDiscarded(&mtproto.PhoneCall{
					Id:        phonePeerId,
					NeedDebug: true,
					Reason:    reason,
					Duration:  duration,
				}).To_PhoneCall(),
			}).To_Update(),
		},
		Users: []*mtproto.User{},
		Date:  int32(time.Now().Unix()),
		Seq:   0,
	}), nil
}

func (m *PhoneCallCore) SendSignalingData(md *metadata.RpcMetaData, userId int64, peer *mtproto.InputPhoneCall, signalingData []byte) (bool, error) {
	info, err := m.phoneCall.TokenPhoneCall(peer.Id)
	if err != nil {
		log.Errorf("SendSignalingData userId:%d phoneCallId:%d error:%s", userId, peer.GetId(), err.Error())
		return false, errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_BAD_REQUEST_CALL_ALREADY_DECLINED, err.Error())
	}

	if info.ID == 0 {
		return false, nil
	}

	//phoneCallId := info.ID
	//if info.ID != peer.Id {
	//	phoneCallId = peer.Id
	//}

	//toAuthKeyId := info.AuthKeyId
	//toUserId := info.Admin
	//if userId == info.Admin {
	//	toUserId = info.ParticipantId
	//	toAuthKeyId = info.AcceptAuthKeyId
	//}
	go func() {
		//userInfo, _ := dbproxyClient.GetDBProxyClient().UserByUserIdList(context.Background(), &dbproxy_pb.UserByUserIdList{
		//	OwnerUserId:          toUserId,
		//	UserIdList:           []int32{info.Admin, info.ParticipantId},
		//  Now: int32(time.Now().Unix()),
		//})

		// updatePhoneCallSignalingData#2661bf09 phone_call_id:long data:bytes = Update;
		//updates := mtproto.NewTLUpdates(&mtproto.Updates{
		//	Updates: []*mtproto.Update{
		//		mtproto.NewTLUpdatePhoneCallSignalingData(&mtproto.Update{
		//			PhoneCallId:    phoneCallId,
		//			DataE73547E171: signalingData,
		//		}).To_Update()},
		//	Users: []*mtproto.User{},
		//	Date:  int32(time.Now().Unix()),
		//	Seq:   0,
		//})
		//_, err = syncClient.GetSyncClient().SyncUpdates(context.Background(), &sync_pb.SyncUpdates{
		//	UserId:    toUserId,
		//	AuthKeyId: toAuthKeyId,
		//	Updates:   updates.To_Updates(),
		//})
		//if err != nil {
		//	log.Errorf("syncClient.GetSyncClient() error:%s", err.Error())
		//}
	}()

	return true, nil
}

func (m *PhoneCallCore) SaveCallDebug(userId int64, peer *mtproto.InputPhoneCall, debug *mtproto.DataJSON) {
	log.Warnf("SaveCallDebug userId:%d phoneCallId:%d debug:%s", userId, peer, debug)
}

type PhoneCall struct {
	FromUserId int64
	PeerType   int32
	PeerId     int64
	RandomId   int64
	Video      bool
	Protocol   *mtproto.PhoneCallProtocol
	AccessHash int64
}

func NewPhoneCall(userId int64, peerId int64, video bool, randomId int32, protocol *mtproto.PhoneCallProtocol) *PhoneCall {
	return &PhoneCall{
		FromUserId: userId,
		PeerId:     peerId,
		Video:      video,
		RandomId:   int64(randomId),
		Protocol:   protocol,
		AccessHash: crypto.GenerateAccessHash(),
	}
}

func (m *PhoneCall) CreatePhoneCall(node *snowflake.Node) (*mtproto.TLPhonePhoneCall, error) {

	//userInfo, err := dbproxyClient.GetDBProxyClient().UserByUserIdList(context.Background(), &dbproxy_pb.UserByUserIdList{
	//	OwnerUserId: m.FromUserId,
	//	UserIdList:  []int32{m.FromUserId, m.PeerId},
	//	Now:         int32(time.Now().Unix()),
	//})
	//if err != nil {
	//	log.Errorf("CreatePhoneCall userId:%d peerId:%d error:%s", m.FromUserId, m.PeerId, err.Error())
	//	return nil, err
	//}

	id := node.Generate().Int64()
	// phoneCallWaiting#1b8f4ad1 flags:# video:flags.6?true id:long access_hash:long date:int admin_id:int participant_id:int protocol:PhoneCallProtocol receive_date:flags.0?int = PhoneCall;
	phoneCallWaiting := mtproto.NewTLPhoneCallWaiting(&mtproto.PhoneCall{
		Id:            id,
		AccessHash:    m.AccessHash,
		Video:         m.Video,
		Date:          int32(time.Now().Unix()),
		AdminId:       constants.PeerTypeFromUserIDType(m.FromUserId).ToInt32(),
		ParticipantId: constants.PeerTypeFromUserIDType(m.PeerId).ToInt32(),
		Protocol:      m.Protocol,
		ReceiveDate:   int32(time.Now().Unix()),
	})

	return mtproto.NewTLPhonePhoneCall(&mtproto.Phone_PhoneCall{
		PhoneCall: phoneCallWaiting.To_PhoneCall(),
		//Users:     userInfo.UserList,
	}), nil
}
