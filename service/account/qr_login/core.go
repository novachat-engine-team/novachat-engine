/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package qr_login

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/util"
	"novachat_engine/service/common/errors"
	"time"
)

const expireTime = 30
const qrLoginKey = "qrlogin:"

func makeQRLoginKey(key string) string {
	return fmt.Sprintf("%s%s", qrLoginKey, key)
}

type qrLoginData struct {
	AuthKeyId      int64  `json:"auth_key_id" protobuf:"varint,1,opt,name=aki,proto3"`
	SessionId      int64  `json:"session_id" protobuf:"varint,2,opt,name=si,proto3"`
	Date           int32  `json:"date" protobuf:"varint,3,opt,name=date,proto3"`
	UserId         int64  `json:"user_id" protobuf:"varint,4,opt,name=ui,proto3"`
	Ip             string `json:"ip" protobuf:"bytes,5,opt,name=ip,proto3"`
	ApiId          int32  `json:"api_id" protobuf:"varint,6,opt,name=ai,proto3"`
	DeviceModel    string `json:"device_model" protobuf:"varint,7,opt,name=dm,proto3"`
	SystemVersion  string `json:"system_version" protobuf:"varint,8,opt,name=sv,proto3"`
	SystemLangCode string `json:"system_lang_code" protobuf:"varint,9,opt,name=slc,proto3"`
	AppVersion     string `json:"app_version" protobuf:"varint,10,opt,name=av,proto3"`
	LangPack       string `json:"lang_pack" protobuf:"varint,11,opt,name=lp,proto3"`
	LangCode       string `json:"lang_code" protobuf:"varint,12,opt,name=lc,proto3"`
}

var defaultQrLoginData = qrLoginData{}

func (q *qrLoginData) Reset() {
	*q = qrLoginData{}
}
func (q *qrLoginData) String() string {
	return proto.CompactTextString(q)
}
func (q *qrLoginData) ProtoMessage() {

}

func makeQRLoginData(
	authKeyId int64,
	sessionId int64,
	date int32,
	ip string,
	deviceModel string,
	systemVersion string,
	systemLangCode string,
	appVersion string,
	langPack string,
	langCode string,
	apiId int32,
	userId int64) (string, string) {
	qrLoginDataValue := &qrLoginData{
		AuthKeyId: authKeyId,
		SessionId: sessionId,
		Date:      0,
	}

	encoder := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: false,
		Indent:       "",
		OrigName:     false,
		AnyResolver:  nil,
	}

	vvv, _ := encoder.MarshalToString(qrLoginDataValue)
	_ = vvv
	buffer := bytes.Buffer{}
	_ = encoder.Marshal(&buffer, qrLoginDataValue)
	data := buffer.Bytes()
	hash := sha256.New()
	hash.Write(data)
	key := hex.EncodeToString(hash.Sum(nil))
	qrLoginDataValue.Date = date
	qrLoginDataValue.Ip = ip
	qrLoginDataValue.ApiId = apiId
	qrLoginDataValue.UserId = userId
	qrLoginDataValue.DeviceModel = deviceModel
	qrLoginDataValue.SystemVersion = systemVersion
	qrLoginDataValue.SystemLangCode = systemLangCode
	qrLoginDataValue.AppVersion = appVersion
	qrLoginDataValue.LangPack = langPack
	qrLoginDataValue.LangCode = langCode
	data, _ = jsoniter.Marshal(qrLoginDataValue)
	return key, string(data)
}

func parseQRLoginData(data string) qrLoginData {
	qrLoginDataValue := qrLoginData{}
	if err := jsoniter.UnmarshalFromString(data, &qrLoginDataValue); err != nil {
		log.Errorf("parseQRLoginData data:`%s` error:%s", data, err.Error())
	}

	return qrLoginDataValue
}

type Core struct {
	key string
}

func NewQRLoginCore(key string, conf config.RedisConfig) *Core {
	cache.InstallRedis(conf)
	return &Core{
		key: key,
	}
}

func (c *Core) getToken(key string) (string, error) {
	ret, err := redis.String(cache.GetRedisClient().Do("GET", makeQRLoginKey(key)))
	if err != nil && err != redis.ErrNil {
		log.Errorf("getToken key:%s error:%s", key, err.Error())
		return "", err
	}
	return ret, nil
}

func (c *Core) fetchToken(key string, token string, expire int32) (string, error) {
	ret, err := redis.String(cache.GetRedisClient().Do("SET", makeQRLoginKey(key), token, "EX", expire, "NX"))
	if err != nil && err != redis.ErrNil {
		log.Errorf("fetchToken key:%s error:%s", key, err.Error())
		return "", err
	}

	if err == redis.ErrNil || ret == "nil" {
		return c.getToken(key)
	}
	return token, nil
}

func (c *Core) ExportLoginToken(authKeyId int64, sessionId int64, md *metadata.RpcMetaData, apiId int32, date int32) ([]byte, int32, int64, error) {
	key, data := makeQRLoginData(
		authKeyId,
		sessionId,
		date,
		md.Ip,
		md.DeviceModel,
		md.SystemVersion,
		md.SystemLangCode,
		md.AppVersion,
		md.LangPack,
		md.LangCode,
		apiId,
		0)
	token, err := c.fetchToken(key, data, expireTime)
	if err != nil {
		log.Errorf("ExportLoginToken fetchToken authKeyId:%d sessionId:%d date:%d error:%s", authKeyId, sessionId, date, err.Error())
		return nil, 0, 0, err
	}

	if len(token) == 0 {
		token, err = c.fetchToken(key, data, expireTime)
		if err != nil {
			log.Errorf("ExportLoginToken fetchToken authKeyId:%d sessionId:%d date:%d error:%s", authKeyId, sessionId, date, err.Error())
			return nil, 0, 0, err
		}
	}
	value := parseQRLoginData(token)

	token, err = util.AesEncrypt(token, c.key)
	if err != nil {
		log.Fatalf("ExportLoginToken fetchToken conf.QrLoginKey%s error:%s", c.key, err.Error())
		return nil, 0, 0, err
	}
	tokenData, _ := base64.StdEncoding.DecodeString(token)
	return tokenData, value.Date + expireTime, value.UserId, nil
}

func (c *Core) AcceptLoginToken(token []byte, userId int64) (qrLoginData, error) {
	data, err := util.AesDecrypt(base64.StdEncoding.EncodeToString(token), c.key)
	if err != nil {
		log.Errorf("AcceptLoginToken Decrypt token:`%v` error:%s", token, err.Error())
		return defaultQrLoginData, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_AUTH_TOKEN_INVALIDX)
	}
	qrLoginDataValue := parseQRLoginData(data)
	if qrLoginDataValue.AuthKeyId == 0 {
		log.Error("AcceptLoginToken parseQRLoginData authKeyId = 0")
		return defaultQrLoginData, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_AUTH_TOKEN_INVALID)
	}
	key, value := makeQRLoginData(
		qrLoginDataValue.AuthKeyId,
		qrLoginDataValue.SessionId,
		qrLoginDataValue.Date,
		qrLoginDataValue.Ip,
		qrLoginDataValue.DeviceModel,
		qrLoginDataValue.SystemVersion,
		qrLoginDataValue.SystemLangCode,
		qrLoginDataValue.AppVersion,
		qrLoginDataValue.LangPack,
		qrLoginDataValue.LangCode,
		qrLoginDataValue.ApiId,
		userId)
	tokenValue, err := c.getToken(key)
	if err != nil {
		log.Errorf("AcceptLoginToken getToken key:%s error:%s", key, err.Error())
		return defaultQrLoginData, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_AUTH_TOKEN_EXCEPTION)
	}
	if len(tokenValue) == 0 {
		log.Errorf("AcceptLoginToken getToken key:%s error:EXPIRED", key)
		return defaultQrLoginData, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_AUTH_TOKEN_EXPIRED)
	}

	expire := expireTime - (int32(time.Now().Unix()) - qrLoginDataValue.Date)
	if expire < 0 {
		return defaultQrLoginData, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_AUTH_TOKEN_EXPIRED)
	}
	_, err = c.fetchToken(key, value, expire)
	if err != nil {
		log.Errorf("AcceptLoginToken fetchToken token:`%v` error:%s", token, err.Error())
		return defaultQrLoginData, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_AUTH_TOKEN_INVALIDX)
	}

	return qrLoginDataValue, nil
}
