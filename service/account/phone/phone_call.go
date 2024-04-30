/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/18 15:00
 * @File : phone_call,go.go
 */

package phone

import (
	"encoding/base64"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/json-iterator/go"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/log"
	"novachat_engine/service/constants"
	"strings"
)

const PhoneCallTimeout = 1 * 60          //s
const PhoneCallDiscardTimeout = 10       //s
const PhoneCallPreMaxTime = 24 * 60 * 60 //s

const userPipeLine = true

type PhoneCallType int32

const (
	PhoneCallTypeUser  PhoneCallType = 0
	PhoneCallTypeGroup PhoneCallType = 1
)

const (
	cachePhoneCallUserRequested = "string:phonecall:requested:userId:"
	cachePhoneCallUser          = "string:phonecall:id:"
	cachePhoneCallChat          = "string:phonecall:chat:"
)

func makePhoneCallRequestedKey(userId int64) string {
	return fmt.Sprintf("%s%d", cachePhoneCallUserRequested, userId)
}
func makePhoneCallKey(id int64) string {
	return fmt.Sprintf("%s%d", cachePhoneCallUser, id)
}

type PhoneCallInfo struct {
	ID              int64                      `json:"id" bson:"_id"`
	Status          constants.PhoneCallStatus  `json:"status" bson:"status"`
	Type            PhoneCallType              `json:"type" bson:"type"`
	Date            int32                      `json:"date" bson:"date"`
	Admin           int64                      `json:"admin" bson:"admin"`
	ParticipantId   int64                      `json:"participantId" bson:"participantId"`
	AccessHash      int64                      `json:"access_hash" bson:"access_hash"`
	Video           bool                       `json:"video" bson:"video"`
	GAHash          string                     `json:"gahash" bson:"gahash"`
	GB              string                     `json:"gb" bson:"gb"`
	RandomId        int32                      `json:"random" bson:"random"`
	Protocol        *mtproto.PhoneCallProtocol `json:"protocol" bson:"protocol"`
	P2pAllowed      bool                       `json:"p2p_allowed" bson:"p2p_allowed"`
	AcceptAuthKeyId int64                      `json:"accept_authKey_id" bson:"accept_authKey_id"`
	AuthKeyId       int64                      `json:"authKey_id" bson:"authKey_id"`
}

func (m *PhoneCallInfo) FromGAHash(gaHash []byte) {
	m.GAHash = base64.StdEncoding.EncodeToString(gaHash)
}

func (m *PhoneCallInfo) FromGB(gb []byte) {
	m.GB = base64.StdEncoding.EncodeToString(gb)
}

func (m *PhoneCallInfo) ToGAHash() []byte {
	b, _ := base64.StdEncoding.DecodeString(m.GAHash)
	return b
}

func (m *PhoneCallInfo) ToGB() []byte {
	b, _ := base64.StdEncoding.DecodeString(m.GB)
	return b
}

type PhoneCall struct {
	pool *cache.RedisClient
}

func NewPhoneCall(r *cache.RedisClient) *PhoneCall {
	return &PhoneCall{pool: r}
}

func (m *PhoneCall) PutPhoneCall(info *PhoneCallInfo) (bool, error) {
	var buf, _ = jsoniter.Marshal(info)
	if userPipeLine {
		res, err := m.pool.RedisDoPipeline([]cache.PipeListParam{
			{"SETEX", []interface{}{makePhoneCallKey(info.ID), PhoneCallTimeout, string(buf)}},
			{"SETEX", []interface{}{makePhoneCallRequestedKey(info.ParticipantId), PhoneCallTimeout, info.ID}},
		})
		if err != nil {
			log.Errorf("cache PhoneCall PutPhoneCall:%d error:%s", info.Admin, err.Error())
			return false, err
		}

		log.Debugf("PutPhoneCall RedisDoPipeline res:%+v", res)
		return true, nil

	} else {
		reply, err := redis.String(m.pool.Do("SETEX", makePhoneCallKey(info.ID), PhoneCallTimeout, string(buf)))
		if err != nil {
			log.Errorf("cache PhoneCall PutPhoneCall:%d error:%s", info.Admin, err.Error())
			return false, err
		}

		log.Infof("cache PhoneCall PutPhoneCall:%d reply", info.Admin, reply)

		if strings.Compare(reply, "OK") == 0 {
			return true, nil
		}
	}

	return false, nil
}

func (m *PhoneCall) TokenPhoneCall(phoneCallId int64) (*PhoneCallInfo, error) {
	reply, err := redis.String(m.pool.Do("GET", makePhoneCallKey(phoneCallId)))
	if err != nil && err != redis.ErrNil {
		log.Errorf("cache PhoneCall phoneCallId:%d error:%s", phoneCallId, err.Error())
		return nil, err
	}

	if redis.ErrNil == err {
		return &PhoneCallInfo{}, nil
	}

	log.Infof("cache PhoneCall PutPhoneCall:%d reply:%v", phoneCallId, reply)

	info := &PhoneCallInfo{}
	err = jsoniter.Unmarshal([]byte(reply), info)
	if err != nil {
		log.Errorf("cache PhoneCall phoneCallId:%d error:%s", phoneCallId, err.Error())
		return nil, err
	}

	return info, nil
}

func (m *PhoneCall) CheckComingPhoneCall(userId int64) *PhoneCallInfo {
	if userPipeLine {
		phoneCallId, err := redis.Int64(m.pool.Do("GET", makePhoneCallRequestedKey(userId)))
		if err != nil && err != redis.ErrNil {
			log.Warnf("CheckComingPhoneCall userId:%d error:%s", userId, err.Error())
			return nil
		}

		if err == redis.ErrNil || phoneCallId == 0 {
			return nil
		}

		phoneCallInfo, err := m.TokenPhoneCall(phoneCallId)
		if err != nil {
			log.Errorf("CheckComingPhoneCall phoneCallId:%d error:%s", phoneCallId, err.Error())
			return nil
		}
		return phoneCallInfo
	}
	return nil
}

func (m *PhoneCall) CheckAcceptedPhoneCall(userId int64, authKeyId int64, phoneCallId int64) (bool, error) {
	if userPipeLine {
		id, err := redis.Int64(m.pool.Do("GET", makePhoneCallRequestedKey(userId)))
		if err != nil {
			log.Errorf("cache CheckAcceptedPhoneCall phoneCallId:%d error:%s", phoneCallId, err.Error())
			return false, err
		}
		return id == phoneCallId, nil
	}
	return true, nil
}

func (m *PhoneCall) ConfirmPhoneCall(phoneCallId int64, participantId int64) (bool, error) {

	if userPipeLine {
		res, err := m.pool.RedisDoPipeline([]cache.PipeListParam{
			{"EXPIRE", []interface{}{makePhoneCallKey(phoneCallId), PhoneCallPreMaxTime}},
			{"EXPIRE", []interface{}{makePhoneCallRequestedKey(participantId), -1}},
		})
		if err != nil {
			log.Errorf("ConfirmPhoneCall phoneCallId:%d participantId:%d error:%s", phoneCallId, participantId, err.Error())
			return false, err
		}

		log.Debugf("ConfirmPhoneCall RedisDoPipeline res:%+v", res)
		return true, nil

	} else {
		reply, err := redis.Int(m.pool.Do("EXPIRE", makePhoneCallKey(phoneCallId), PhoneCallPreMaxTime))
		if err != nil && err != redis.ErrNil {
			log.Errorf("ConfirmPhoneCall phoneCallId:%d error:%s", phoneCallId, err.Error())
			return false, err
		}

		log.Infof("cache PhoneCall ConfirmPhoneCall phoneCallId:%d reply:%v", phoneCallId, reply)

		if err == redis.ErrNil {
			return false, nil
		}

	}
	return true, nil
}

func (m *PhoneCall) DiscardPhoneCall(phoneCallId int64, participantId int64) (bool, error) {
	if userPipeLine {
		res, err := m.pool.RedisDoPipeline([]cache.PipeListParam{
			{"EXPIRE", []interface{}{makePhoneCallKey(phoneCallId), -1}},
			{"EXPIRE", []interface{}{makePhoneCallRequestedKey(participantId), -1}},
		})
		if err != nil {
			log.Errorf("DiscardPhoneCall phoneCallId:%d participantId:%d error:%s", phoneCallId, participantId, err.Error())
			return false, err
		}

		log.Debugf("DiscardPhoneCall RedisDoPipeline res:%+v", res)
		return true, nil

	} else {
		reply, err := redis.Int(m.pool.Do("EXPIRE", makePhoneCallKey(phoneCallId), PhoneCallDiscardTimeout))
		if err != nil && err != redis.ErrNil {
			log.Errorf("DiscardPhoneCall phoneCallId:%d error:%s", phoneCallId, err.Error())
			return false, err
		}
		log.Infof("cache PhoneCall DiscardPhoneCall phoneCallId:%d reply:%v", phoneCallId, reply)
	}

	return true, nil
}
