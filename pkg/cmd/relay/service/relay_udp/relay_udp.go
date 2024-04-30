/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package relay_udp

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go.uber.org/atomic"
	"net"
	"novachat_engine/pkg/cmd/relay/conf"
	relayService "novachat_engine/pkg/cmd/relay/rpc_client"
	"novachat_engine/pkg/hack"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/net/netip"
	"novachat_engine/pkg/net/udp"
	"novachat_engine/pkg/util"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

////#define TLID_DECRYPTED_AUDIO_BLOCK 0xDBF948C1
////#define TLID_SIMPLE_AUDIO_BLOCK 0xCC0D0E76
////#define TLID_UDP_REFLECTOR_PEER_INFO 0x27D9371C
////#define TLID_UDP_REFLECTOR_PEER_INFO_IPV6 0x83fc73b1
////#define TLID_UDP_REFLECTOR_SELF_INFO 0xc01572c7
////#define TLID_UDP_REFLECTOR_REQUEST_PACKETS_INFO 0x1a06fc96
////#define TLID_UDP_REFLECTOR_LAST_PACKETS_INFO 0x0e107305
////#define TLID_VECTOR 0x1cb5c415
////#define PAD4(x) (4-(x+(x<=253 ? 1 : 0))%4)

const (
	TLID_DECRYPTED_AUDIO_BLOCK        = 0xDBF948C1
	TLID_SIMPLE_AUDIO_BLOCK           = 0xCC0D0E76
	TLID_UDP_REFLECTOR_PEER_INFO      = 0x27D9371C
	TLID_UDP_REFLECTOR_PEER_INFO_IPV6 = 0x83fc73b1
	TLID_UDP_REFLECTOR_SELF_INFO      = 0xc01572c7
)

const callIdLength = 16

type RelayPeer struct {
	callId     int64
	peerTag    string
	peers      []*net.UDPAddr
	createTime int64
}

func makeCallId(adminId, participantId int64) string {
	sha := sha256.New()
	sha.Write([]byte(fmt.Sprintf("%d_%d_%d", adminId, participantId, time.Now().Unix())))
	return hex.EncodeToString(sha.Sum(nil))[:callIdLength]
}

type none struct{}

func ipEmptyString(ip net.IP) string {
	if len(ip) == 0 {
		return ""
	}
	return ip.String()
}

// Service
// v2.4.4
type Service struct {
	udpConn   *net.UDPConn
	closeChan chan none
	bClose    *atomic.Bool
	locker    sync.RWMutex
	relayMap  map[string]*RelayPeer
	idMap     map[int64]string
	conf      conf.UdpRelay
}

func NewRelayUdpService(conf conf.UdpRelay) (*Service, error) {
	if len(conf.EndPoint) == 0 {
		conf.EndPoint = "0.0.0.0"
	}

	if conf.Port <= 0 {
		log.Warn("relayUdpService is nil")
		return nil, nil
	}

	udpConn, err := udp.ListenUdp(fmt.Sprintf("%s:%d", conf.EndPoint, conf.Port))
	if err != nil {
		return nil, err
	}

	r := Service{
		udpConn:   udpConn,
		closeChan: make(chan none, 1),
		bClose:    atomic.NewBool(false),
		conf:      conf,
	}

	go r.runForever()
	return &r, nil
}

func (r *Service) Close() {
	if r.bClose.Load() == false {
		r.bClose.Store(true)
		r.closeChan <- none{}
	}
}

func (r *Service) runForever() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("runForever r:%v error:%s", r, debug.Stack())
		} else {
			log.Infof("runForever quit!!!!!!!!!!!!!!!")
		}
	}()

	buf := make([]byte, 4096)
	var (
		bufLen int
		remote *net.UDPAddr
		err    error
	)

	for {
		select {
		case <-r.closeChan:
			return
		default:
			bufLen, remote, err = r.udpConn.ReadFromUDP(buf)
			if err != nil {
				if strings.Contains(err.Error(), "use of closed network connection") == true {
					//log.Debugf("readFromUDP udp server info:%v", err)
				} else {
					log.Errorf("readFromUDP error - %v", err)
				}
				break
			} else {
				r.onRelayData(remote, buf[:bufLen])
			}
		}
	}
}

func (r *Service) onRelayData(remote *net.UDPAddr, bytes []byte) {
	if len(bytes) < 32 {
		log.Warnf("onRelayData invalid data len %d, addr %s", len(bytes), remote.String())
		return
	}

	buf := util.NewBufferInput(bytes)
	peerTag := hack.String(buf.Bytes(16))
	packetTypes := make([]uint32, 4)
	packetTypes[0] = buf.UInt32()
	packetTypes[1] = buf.UInt32()
	packetTypes[2] = buf.UInt32()
	packetTypes[3] = buf.UInt32()

	if packetTypes[0] == 0xFFFFFFFF &&
		packetTypes[1] == 0xFFFFFFFF &&
		packetTypes[2] == 0xFFFFFFFF &&
		packetTypes[3] == 0xFFFFFFFF {
		r.onReflectorPeerInfo(remote, peerTag)
	} else if packetTypes[0] == 0xFFFFFFFF &&
		packetTypes[1] == 0xFFFFFFFF &&
		packetTypes[2] == 0xFFFFFFFF &&
		packetTypes[3] == 0xFFFFFFFE {
		r.onReflectorSelfInfo(remote, peerTag, buf.UInt64())
	} else {
		r.onRelayDataAvailable(remote, peerTag, bytes)
	}
}

func (r *Service) onReflectorSelfInfo(cAddr *net.UDPAddr, peerTag string, queryId uint64) {
	log.Debugf("onReflectorSelfInfo addr:%s, peer_tag:%s, query_id: %d", cAddr.String(), peerTag, queryId)
	oBuf := util.NewBufferOutput(1024)
	oBuf.Bytes(hack.Bytes(peerTag))
	oBuf.UInt32(0xFFFFFFFF)
	oBuf.UInt32(0xFFFFFFFF)
	oBuf.UInt32(0xFFFFFFFF)
	oBuf.UInt32(TLID_UDP_REFLECTOR_SELF_INFO)
	oBuf.UInt32(uint32(time.Now().Unix()))
	oBuf.UInt64(queryId)

	ip := hack.Bytes(ipEmptyString(cAddr.IP))
	oBuf.Bytes(ip)
	oBuf.Bytes(make([]byte, 16-len(ip)))
	oBuf.UInt32(uint32(cAddr.Port))

	_, err := r.udpConn.WriteToUDP(oBuf.Buf(), cAddr)
	if err != nil {
		log.Warnf("onReflectorSelfInfo addr:%d error:%s", cAddr.String(), err.Error())
	}
}

//uint32_t myAddr=(uint32_t) in.ReadInt32();
//uint32_t myPort=(uint32_t) in.ReadInt32();
//uint32_t peerAddr=(uint32_t) in.ReadInt32();
//uint32_t peerPort=(uint32_t) in.ReadInt32();
func (r *Service) onReflectorPeerInfo(cAddr *net.UDPAddr, peerTag string) {
	log.Infof("onReflectorPeerInfo peer:%s, peer_tag: %s", cAddr.String(), peerTag)
	oBuf := util.NewBufferOutput(1024)
	oBuf.Bytes(hack.Bytes(peerTag))
	oBuf.UInt32(0xFFFFFFFF)
	oBuf.UInt32(0xFFFFFFFF)
	oBuf.UInt32(0xFFFFFFFF)
	oBuf.UInt32(TLID_UDP_REFLECTOR_PEER_INFO)

	r.locker.RLock()
	v, ok := r.relayMap[peerTag]
	if ok == false {
		r.locker.RUnlock()
		return
	}
	r.locker.RUnlock()
	v.createTime = time.Now().Unix()

	if len(v.peers) < 2 {
		log.Errorf("onReflectorPeerInfo peer:%s, peer_tag:%s len:%d", cAddr.String(), peerTag, len(v.peers))
		return
	}

	oBuf.Int32(int32(netip.IPToUInt32(v.peers[0].IP)))
	oBuf.Int32(int32(v.peers[0].Port))
	oBuf.Int32(int32(netip.IPToUInt32(v.peers[1].IP)))
	oBuf.Int32(int32(v.peers[1].Port))

	log.Debugf("onReflectorPeerInfo peers[0].IP:%s peers[0].Port:%d peers[1].IP:%s peers[1].Port:%d", v.peers[0].IP, v.peers[0].Port, v.peers[1].IP, v.peers[1].Port)

	_, err := r.udpConn.WriteToUDP(oBuf.Buf(), cAddr)
	if err != nil {
		log.Warnf("onReflectorPeerInfo addr:%d error:%s", cAddr.String(), err.Error())
	}
}

func equalAddr(a1, a2 *net.UDPAddr) bool {
	if a1 == nil || a2 == nil {
		return false
	}

	return bytes.Equal(a1.IP, a2.IP) && a1.Port == a2.Port && a1.Zone == a2.Zone
}

func (r *Service) onRelayDataAvailable(cAddr *net.UDPAddr, peerTag string, buf []byte) {
	var v *RelayPeer
	var ok bool
	r.locker.RLock()
	if v, ok = r.relayMap[peerTag]; !ok {
		r.locker.RUnlock()
		log.Errorf("not found {peer_tag: %s} by {addr: %s}", peerTag, cAddr.String())
		return
	} else {
		r.locker.RUnlock()
	}
	v.createTime = time.Now().Unix()

	found := false
	for _, e := range v.peers {
		if equalAddr(cAddr, e) {
			found = true
			break
		}
	}
	if !found {
		v.peers = append(v.peers, cAddr)
	}

	for _, e := range v.peers {
		if !equalAddr(e, cAddr) {
			r.udpConn.WriteToUDP(buf, e)
		}
	}
}

func (r *Service) Check(callId int64) error {
	r.locker.RLock()
	if peerTag, ok := r.idMap[callId]; ok {
		vv, ok := r.relayMap[peerTag]
		if !ok {
			r.locker.RUnlock()
			r.locker.Lock()
			delete(r.idMap, callId)
			r.locker.Unlock()
			log.Warnf("relayMap not found peerTag:%s callId:%d", peerTag, callId)
			return nil
		}
		if time.Now().Unix()-vv.createTime < 60 {
			r.locker.RUnlock()
			return nil
		} else {
			r.locker.Lock()
			delete(r.idMap, callId)
			delete(r.relayMap, peerTag)
			r.locker.Unlock()
			return fmt.Errorf("caldId:%d timeout", callId)
		}
	}

	r.locker.RUnlock()
	return nil
}

func (r *Service) DiscardPhoneCall(callId int64) error {
	log.Debugf("DiscardPhoneCall  callId:%d", callId)

	r.locker.Lock()
	defer r.locker.Unlock()

	if peerTag, ok := r.idMap[callId]; ok {
		delete(r.relayMap, peerTag)
		delete(r.idMap, callId)
		log.Infof("DiscardPhoneCall peerTag:%s ", peerTag)
	}

	log.Info("DiscardPhoneCall reply - true")
	return nil
}

func (r *Service) CreatePhoneCall(callId int64, adminId int64, participantId int64, version string) (*relayService.CallConnections, error) {
	log.Debugf("CreatePhoneCall  callId:%d adminId:%d participantId:%d version:%s", callId, adminId, participantId, version)

	r.locker.Lock()
	defer r.locker.Unlock()
	var resp *relayService.CallConnections

	if peerTag, ok := r.idMap[callId]; ok {
		resp = &relayService.CallConnections{
			Connect: &relayService.CallConnections_Connect{
				PeerTag: hack.Bytes(peerTag),
				Ip:      r.conf.EndPoint,
				Port:    r.conf.Port,
			},
		}
	} else {
		peerTag = makeCallId(adminId, participantId)
		if vv, ok := r.relayMap[peerTag]; ok == true {
			log.Warnf("CreatePhoneCall callId:%d adminId:%d participantId:%d peerTag:%s", callId, adminId, participantId, peerTag)
			vv.createTime = time.Now().Unix()
			resp = &relayService.CallConnections{
				Connect: &relayService.CallConnections_Connect{
					PeerTag: hack.Bytes(peerTag),
					Ip:      r.conf.EndPoint,
					Port:    r.conf.Port,
				},
			}
		} else {
			r.idMap[callId] = peerTag
			v := &RelayPeer{callId: callId, peerTag: peerTag, peers: []*net.UDPAddr{}, createTime: time.Now().Unix()}
			r.relayMap[peerTag] = v
			resp = &relayService.CallConnections{
				Connect: &relayService.CallConnections_Connect{
					PeerTag: hack.Bytes(peerTag),
					Ip:      r.conf.EndPoint,
					Port:    r.conf.Port,
				},
			}
		}
	}
	resp.EndpointId = callId
	log.Infof("CreatePhoneCall reply:%s", resp)
	return resp, nil
}
