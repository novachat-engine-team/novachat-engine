/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/2/21 16:13
 * @File : handshake.go
 */

package handshake

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/gogo/protobuf/types"
	"math/big"
	"math/rand"
	"novachat_engine/mtproto"
	authClient "novachat_engine/pkg/cmd/auth/rpc_client"
	"novachat_engine/pkg/cmd/gate/data"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/dh_config"
	"novachat_engine/pkg/hack"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/net/mtserver/codec"
	"novachat_engine/service/app/dc"
	"novachat_engine/service/constants"
	"time"
)

type HandShake interface {
	OnHandshake(ctx *data.Context, mtpBuf []byte, connId uint64) ([]byte, error)
}

var SaltUtilTime = int32(30 * 60)
var KeyFingerPrintLength = 20

const (
	SHA_DIGEST_LENGTH = 20

	NonceLength = 16 // 16 * 8 = 128
)

var (
	gBigIntDH2048P *big.Int
	gBigIntDH2048G *big.Int
)

func init() {
	rand.NewSource(time.Now().UnixNano())

	gBigIntDH2048P = new(big.Int).SetBytes(dh_config.DH2048P)
	gBigIntDH2048G = new(big.Int).SetBytes(dh_config.DH2048G)
}

type handshake struct {
	rsa            *crypto.RSACryptor
	keyFingerprint int64
	dh2048p        []byte
	dh2048g        []byte
}

//
func NewHandshake(keyFile string, keyFingerprint int64) (*handshake, error) {
	rsa, err := crypto.NewRSACryptor(keyFile)
	if err != nil {
		return nil, err
	}
	return &handshake{
		rsa:            rsa,
		keyFingerprint: keyFingerprint,
		dh2048p:        dh_config.DH2048P,
		dh2048g:        dh_config.DH2048G,
	}, nil
}

func (s *handshake) OnHandshake(ctx *data.Context, mtpBuf []byte, connId uint64) ([]byte, error) {
	var (
		err error
		res mtproto.TLObject
	)

	if len(mtpBuf) < 8 {
		err = fmt.Errorf("invalid data len < 8")
		log.Error(err.Error())
		return nil, err
	}

	decode := mtproto.NewDecodeBuf(mtpBuf[:8])
	authKeyId := decode.Long() // 0-7
	_ = authKeyId              // value = 0

	decode = mtproto.NewDecodeBuf(mtpBuf[8:])
	messageId := decode.Long()    // 8-15
	messageLength := decode.Int() //16-19

	layer := int32(0)
	obj := decode.Object(layer)
	err = decode.GetError()
	if err != nil {
		err = decode.GetError()
		return nil, err
	}

	if decode.GetOffset() != int(messageLength) {

	}

	switch obj.(type) {
	case *mtproto.TLReqPq:
		//message_id = 51e57ac42770964a
		//message_length = 20
		//if ctx.Status() == data.StatusReqPQ {
		res, err = s.onReqPQ(ctx, obj.(*mtproto.TLReqPq), messageId, connId)
		if err == nil {
			ctx.SetStatus(data.StatusResQP)
		}
		//} else {
		//	return nil, fmt.Errorf("type: ReqPq status != StatusReqPQ status:%v", ctx.Status())
		//}
	case *mtproto.TLReqPqMulti:
		// message_id = 51e57ac42770964a
		// message_length = 20
		//if ctx.Status() == data.StatusReqPQ {
		res, err = s.onReqPQMulti(ctx, obj.(*mtproto.TLReqPqMulti), messageId, connId)
		if err == nil {
			ctx.SetStatus(data.StatusResQP)
		}
		//} else {
		//	return nil, fmt.Errorf("type: TLReqPQMulti status != StatusReqPQ status:%v", ctx.Status())
		//}
	case *mtproto.TLReq_DHParams:
		//if ctx.Status() == data.StatusResQP {
		ctx.SetStatus(data.StatusReqDH)
		res, err = s.onReqDHParams(ctx, obj.(*mtproto.TLReq_DHParams), messageId, connId)
		if err == nil {
			ctx.SetStatus(data.StatusResDH)
		}
		//} else {
		//	return nil, fmt.Errorf("type: TLReqPQMulti status != StatusReqPQ status:%v", ctx.Status())
		//}
	case *mtproto.TLMsgsAck:
		res, err = s.onMsgsAck(ctx, obj.(*mtproto.TLMsgsAck))
		if err != nil {

		}
		//return nil, err
	case *mtproto.TLSetClient_DHParams:
		//if ctx.Status() == data.StatusResDH {
		var Answer *mtproto.SetClient_DHParamsAnswer
		Answer, err = s.onSetClient_DHParams(ctx, obj.(*mtproto.TLSetClient_DHParams))
		res = Answer
		if err == nil && Answer.ClassName == mtproto.ClassDhGenOk {
			ctx.SetStatus(data.StatusDHParamOk)
		} else {
			ctx.SetStatus(data.StatusDHParamFail)
		}
		//} else {
		//	return nil, fmt.Errorf("type: TLReqPQMulti status != StatusReqPQ status:%v", ctx.Status())
		//}

	default:
		err = fmt.Errorf("onHandshake error obj:%d", obj)
		log.Errorf(err.Error())
		return nil, err
	}

	if err == nil {
		messageId = codec.GenerateMessageId()
		//messageId = message_id
		buf := res.Encode(0)
		encoder := mtproto.NewEncodeBuf(8 + 8 + 4 + len(buf))
		encoder.Long(0) // auth_key_id
		encoder.Long(messageId)
		encoder.Int(int32(len(buf)))
		encoder.Bytes(buf)
		return encoder.GetBuf(), nil
	}

	return nil, err
}

func (s *handshake) onReqPQ(ctx *data.Context, reqPQ *mtproto.TLReqPq, messageId int64, connId uint64) (*mtproto.TLResPQ, error) {
	if len(reqPQ.GetNonce()) != NonceLength {
		return nil, fmt.Errorf("handshake::onReqPQ nonce length error lenth:%d", len(reqPQ.GetNonce()))
	}

	resPQ := mtproto.NewTLResPQ(nil)
	resPQ.SetNonce(reqPQ.GetNonce())
	resPQ.SetPq(dh_config.PQ) // 算法 p q
	resPQ.SetServerNonce(crypto.GenerateNonce(NonceLength))
	resPQ.SetServerPublicKeyFingerprints([]int64{int64(s.keyFingerprint)}) // 传输公钥 需要配置

	ctx.SetNonce(reqPQ.GetNonce())
	ctx.SetServerNonce(resPQ.GetServerNonce())
	ctx.SetP(dh_config.P)
	ctx.SetQ(dh_config.Q)
	return resPQ, nil
}

func (s *handshake) onReqPQMulti(ctx *data.Context, multi *mtproto.TLReqPqMulti, messageId int64, connId uint64) (*mtproto.TLResPQ, error) {
	if len(multi.GetNonce()) != NonceLength {
		return nil, fmt.Errorf("handshake::onReqPQMulti nonce length error lenth:%d", len(multi.GetNonce()))
	}

	resPQ := mtproto.NewTLResPQ(nil)
	resPQ.SetNonce(multi.GetNonce())
	resPQ.SetPq(dh_config.PQ) // 算法 p q
	resPQ.SetServerNonce(crypto.GenerateNonce(NonceLength))
	resPQ.SetServerPublicKeyFingerprints([]int64{int64(s.keyFingerprint)}) // 传输公钥 需要配置

	ctx.SetNonce(multi.GetNonce())
	ctx.SetServerNonce(resPQ.GetServerNonce())
	ctx.SetP(dh_config.P)
	ctx.SetQ(dh_config.Q)
	return resPQ, nil
}

func (s *handshake) onReqDHParams(ctx *data.Context, reqDHParams *mtproto.TLReq_DHParams, messageId int64, connId uint64) (*mtproto.Server_DH_Params, error) {
	errNonce := crypto.GenerateNonce(NonceLength)
	if bytes.Equal(ctx.Nonce(), reqDHParams.GetNonce()) == false || bytes.Equal(ctx.ServerNonce(), reqDHParams.GetServerNonce()) == false {
		serverDHParams := mtproto.NewTLServer_DHParamsFail(nil)
		serverDHParams.Data2.Nonce = errNonce
		serverDHParams.Data2.ServerNonce = errNonce
		serverDHParams.Data2.NewNonceHash = calcNewNonceHash(errNonce, errNonce, 0x03)
		log.Errorf("handshake::onReqDHParams ctx:%v reqDHParams:%v error: Nonce or ServerNonce no equal ", ctx, reqDHParams)
		return serverDHParams.To_Server_DH_Params(), nil
	}

	if bytes.Equal(ctx.Q(), []byte(reqDHParams.GetQ())) == false || bytes.Equal(ctx.P(), []byte(reqDHParams.GetP())) == false {
		serverDHParams := mtproto.NewTLServer_DHParamsFail(nil)
		serverDHParams.Data2.Nonce = errNonce
		serverDHParams.Data2.ServerNonce = errNonce
		serverDHParams.Data2.NewNonceHash = calcNewNonceHash(errNonce, errNonce, 0x03)
		log.Errorf("handshake::onReqDHParams ctx:%v reqDHParams:%v error: P or Q no equal ", ctx, reqDHParams)
		return serverDHParams.To_Server_DH_Params(), nil
	}

	if reqDHParams.GetPublicKeyFingerprint() != s.keyFingerprint {
		serverDHParams := mtproto.NewTLServer_DHParamsFail(nil)
		serverDHParams.Data2.Nonce = errNonce
		serverDHParams.Data2.ServerNonce = errNonce
		serverDHParams.Data2.NewNonceHash = calcNewNonceHash(errNonce, errNonce, 0x03)
		log.Errorf("handshake::onReqDHParams ctx:%v reqDHParams:%v error: keyFingerprint ", ctx, reqDHParams)
		return serverDHParams.To_Server_DH_Params(), nil
	}

	// ???
	// SHA1 (data) = DB761C27718A2305044F71F2AD951629D78B2449
	// RSA (data_with_hash, server_public_key) = 7BB0100A523161904D9C69FA04BC60DECFC5DD74B99995C768EB60D8716E2109BAF2D4601DAB6B09610DC11067BB89021E09471FCFA52DBD0F23204AD8CA8B012BF40A112F44695AB6C266955386114EF5211E6372227ADBD34995D3E0E5FF02EC63A43F9926878962F7C570E6A6E78BF8366AF917A5272675C46064BE62E3E202EFA8B1ADFB1C32A898C2987BE27B5F31D57C9BB963ABCB734B16F652CEDB4293CBB7C878A3A3FFAC9DBEA9DF7C67BC9E9508E111C78FC46E057F5C65ADE381D91FEE430A6B576A99BDF8551FDB1BE2B57069B1A45730618F27427E8A04720B4971EF4A9215983D68F2830C3EAA6E40385562F970D38A05C9F1246DC33438E6
	// The length of the final string was 256 bytes.
	// ???
	decryptData := s.rsa.Decrypt([]byte(reqDHParams.EncryptedData))

	if len(decryptData) < 255 {
		serverDHParams := mtproto.NewTLServer_DHParamsFail(nil)
		serverDHParams.Data2.Nonce = errNonce
		serverDHParams.Data2.ServerNonce = errNonce
		serverDHParams.Data2.NewNonceHash = calcNewNonceHash(errNonce, errNonce, 0x03)
		err := fmt.Errorf("handshake::onReqDHParams - len(encryptedPQInnerData) != 255")
		log.Error(err.Error())
		return serverDHParams.To_Server_DH_Params(), nil
	}

	paddingLen := len(decryptData) - SHA_DIGEST_LENGTH
	if len(decryptData) > 255 {
		paddingLen = 0
	}
	if !checkSha1(decryptData, paddingLen) {
		serverDHParams := mtproto.NewTLServer_DHParamsFail(nil)
		serverDHParams.Data2.Nonce = errNonce
		serverDHParams.Data2.ServerNonce = errNonce
		serverDHParams.Data2.NewNonceHash = calcNewNonceHash(errNonce, errNonce, 0x03)
		err := fmt.Errorf("handshake::onReqDHParams - sha1Check error")
		log.Error(err.Error())
		return serverDHParams.To_Server_DH_Params(), nil
	}
	// 2. 反序列化出pqInnerData
	dbuf := mtproto.NewDecodeBuf(decryptData[SHA_DIGEST_LENGTH:])
	o := dbuf.Object(0)
	if dbuf.GetError() != nil {
		err := fmt.Errorf("handshake::onReqDHParams - decode P_Q_inner_data error")
		log.Error(err.Error())
		return nil, err
	}

	var pqInnerData *mtproto.P_QInnerData

	switch o.(type) {
	case *mtproto.TLPQInnerData:
		pqInnerData = o.(*mtproto.TLPQInnerData).To_P_QInnerData()
		pqInnerData.Dc = dc.DefaultDc
		ctx.SetAuthKeyTemp(constants.AuthKeyStatusDC)
		break
	case *mtproto.TLPQInnerDataDc:
		pqInnerData = o.(*mtproto.TLPQInnerDataDc).To_P_QInnerData()
		ctx.SetAuthKeyTemp(constants.AuthKeyStatusDC)
		break
	case *mtproto.TLPQInnerDataTemp:
		pqInnerData = o.(*mtproto.TLPQInnerDataTemp).To_P_QInnerData()
		pqInnerData.Dc = dc.DefaultDc
		ctx.SetAuthKeyTemp(constants.AuthKeyStatusTemp)
		break
	case *mtproto.TLPQInnerDataTempDc:
		pqInnerData = o.(*mtproto.TLPQInnerDataTempDc).To_P_QInnerData()
		ctx.SetAuthKeyTemp(constants.AuthKeyStatusTemp)
		break
	default:
		err := fmt.Errorf("handshake::onReqDHParams - decode P_Q_inner_data params error :%v", o)
		serverDHParams := mtproto.NewTLServer_DHParamsFail(nil)
		serverDHParams.Data2.Nonce = errNonce
		serverDHParams.Data2.ServerNonce = errNonce
		serverDHParams.Data2.NewNonceHash = calcNewNonceHash(errNonce, errNonce, 0x03)
		log.Error(err.Error())
		return serverDHParams.To_Server_DH_Params(), nil
	}

	if bytes.Equal(ctx.Nonce(), pqInnerData.GetNonce()) == false || bytes.Equal(ctx.ServerNonce(), pqInnerData.GetServerNonce()) == false {
		serverDHParams := mtproto.NewTLServer_DHParamsFail(nil)
		serverDHParams.Data2.Nonce = ctx.Nonce()
		serverDHParams.Data2.ServerNonce = ctx.ServerNonce()
		serverDHParams.Data2.NewNonceHash = calcNewNonceHash(ctx.NewNonce(), ctx.ServerNonce(), 0x03)
		log.Errorf("handshake::onReqDHParams ctx:%v pqInnerData:%v error: Nonce or ServerNonce no equal ", ctx, pqInnerData)
		return serverDHParams.To_Server_DH_Params(), nil
	}

	if bytes.Equal(ctx.Q(), []byte(pqInnerData.GetQ())) == false || bytes.Equal(ctx.P(), []byte(reqDHParams.GetP())) == false {
		serverDHParams := mtproto.NewTLServer_DHParamsFail(nil)
		serverDHParams.Data2.Nonce = errNonce
		serverDHParams.Data2.ServerNonce = errNonce
		serverDHParams.Data2.NewNonceHash = calcNewNonceHash(errNonce, errNonce, 0x03)
		log.Errorf("handshake::onReqDHParams ctx:%v pqInnerData:%v error: P or Q no equal ", ctx, pqInnerData)
		return serverDHParams.To_Server_DH_Params(), nil
	}

	ctx.SetDc(pqInnerData.GetDc())
	ctx.SetExpiresIn(pqInnerData.GetExpiresIn())
	ctx.SetNewNonce(pqInnerData.GetNewNonce())

	A := crypto.GenerateNonce(256)
	ctx.SetA(A)
	ctx.SetP(s.dh2048p)

	bigIntA := new(big.Int).SetBytes(A)

	gA := new(big.Int).Exp(gBigIntDH2048G, bigIntA, gBigIntDH2048P)

	serverDHInnerData := mtproto.NewTLServer_DHInnerData(nil)
	serverDHInnerData.SetNonce(ctx.Nonce())
	serverDHInnerData.SetServerNonce(ctx.ServerNonce())
	serverDHInnerData.SetG(int32(s.dh2048g[0]))
	serverDHInnerData.SetGA(string(gA.Bytes()))
	serverDHInnerData.SetDhPrime(string(s.dh2048p))
	serverDHInnerData.SetServerTime(int32(time.Now().Unix()))

	serverDHInnerDataBuf := serverDHInnerData.Encode(0)

	tmpAesKeyAndIV := make([]byte, 64)
	sha1A := sha1.Sum(append(ctx.NewNonce(), ctx.ServerNonce()...))
	sha1B := sha1.Sum(append(ctx.ServerNonce(), ctx.NewNonce()...))
	sha1C := sha1.Sum(append(ctx.NewNonce(), ctx.NewNonce()...))
	copy(tmpAesKeyAndIV, sha1A[:])
	copy(tmpAesKeyAndIV[20:], sha1B[:])
	copy(tmpAesKeyAndIV[40:], sha1C[:])
	copy(tmpAesKeyAndIV[60:], ctx.NewNonce()[:4])

	tmpLen := 20 + len(serverDHInnerDataBuf)
	if tmpLen%16 > 0 {
		tmpLen = (tmpLen/16 + 1) * 16
	} else {
		tmpLen = 20 + len(serverDHInnerDataBuf)
	}

	tmpEncryptedAnswer := make([]byte, tmpLen)
	sha1Tmp := sha1.Sum(serverDHInnerDataBuf)
	copy(tmpEncryptedAnswer, sha1Tmp[:])
	copy(tmpEncryptedAnswer[20:], serverDHInnerDataBuf)

	e := crypto.NewAES256IGECryptor(tmpAesKeyAndIV[:32], tmpAesKeyAndIV[32:64])
	tmpEncryptedAnswer, _ = e.Encrypt(tmpEncryptedAnswer)

	serverDHParams := mtproto.NewTLServer_DHParamsOk(nil)
	serverDHParams.Data2.Nonce = ctx.Nonce()
	serverDHParams.Data2.ServerNonce = ctx.ServerNonce()
	serverDHParams.Data2.EncryptedAnswer = hack.String(tmpEncryptedAnswer)
	return serverDHParams.To_Server_DH_Params(), nil
}

func (s *handshake) onMsgsAck(ctx *data.Context, ack *mtproto.TLMsgsAck) (*mtproto.TLMsgsAck, error) {
	//TODO:(Coder)
	return ack, nil
}

func (s *handshake) onSetClient_DHParams(ctx *data.Context, setClientDHParams *mtproto.TLSetClient_DHParams) (*mtproto.SetClient_DHParamsAnswer, error) {
	if bytes.Equal(ctx.Nonce(), setClientDHParams.GetNonce()) == false || bytes.Equal(ctx.ServerNonce(), setClientDHParams.GetServerNonce()) == false {
		log.Errorf("handshake::onSetClient_DHParams ctx:%v setClientDHParams:%v  Nonce or ServerNonce error", ctx, setClientDHParams)
		setClientDHParamsAnswer := mtproto.NewTLDhGenFail(nil)
		setClientDHParamsAnswer.SetNonce(setClientDHParams.Nonce)
		setClientDHParamsAnswer.SetServerNonce(setClientDHParams.ServerNonce)
		setClientDHParamsAnswer.SetNewNonceHash3(calcNewNonceHash(ctx.NewNonce(), ctx.ServerNonce(), 0x03))
		return setClientDHParamsAnswer.To_SetClient_DHParamsAnswer(), nil
	}

	encryptedData := []byte(setClientDHParams.EncryptedData)

	// 创建aes和iv key
	tmpAesKeyAndIv := make([]byte, 64)
	sha1A := sha1.Sum(append(ctx.NewNonce(), ctx.ServerNonce()...))
	sha1B := sha1.Sum(append(ctx.ServerNonce(), ctx.NewNonce()...))
	sha1C := sha1.Sum(append(ctx.NewNonce(), ctx.NewNonce()...))
	copy(tmpAesKeyAndIv, sha1A[:])
	copy(tmpAesKeyAndIv[20:], sha1B[:])
	copy(tmpAesKeyAndIv[40:], sha1C[:])
	copy(tmpAesKeyAndIv[60:], ctx.NewNonce()[:4])

	d := crypto.NewAES256IGECryptor(tmpAesKeyAndIv[:32], tmpAesKeyAndIv[32:64])
	decryptedData, err := d.Decrypt(encryptedData)
	if err != nil {
		err := fmt.Errorf("handshake::onSetClient_DHParams - AES256IGECryptor descrypt error")
		log.Error(err.Error())
		setClientDHParamsAnswer := mtproto.NewTLDhGenFail(nil)
		setClientDHParamsAnswer.SetNonce(setClientDHParams.Nonce)
		setClientDHParamsAnswer.SetServerNonce(setClientDHParams.ServerNonce)
		setClientDHParamsAnswer.SetNewNonceHash3(calcNewNonceHash(ctx.NewNonce(), ctx.ServerNonce(), 0x03))
		return setClientDHParamsAnswer.To_SetClient_DHParamsAnswer(), nil
	}

	dbuf := mtproto.NewDecodeBuf(decryptedData[20:])
	obj := dbuf.Object(0)
	if err != nil {
		log.Errorf("handshake::onSetClient_DHParams - Client_DH_Inner_Data decode error: %s", err)
		return nil, err
	}
	clientDHInnerData := obj.(*mtproto.TLClient_DHInnerData)

	if bytes.Equal(clientDHInnerData.GetNonce(), ctx.Nonce()) == false || bytes.Equal(clientDHInnerData.GetServerNonce(), ctx.ServerNonce()) == false {
		err = fmt.Errorf("handshake::onSetClient_DHParams - clientDHInnerData:%v  Nonce or ServerNonce error", clientDHInnerData)
		log.Error(err.Error())
		setClientDHParamsAnswer := mtproto.NewTLDhGenFail(nil)
		setClientDHParamsAnswer.SetNonce(setClientDHParams.Nonce)
		setClientDHParamsAnswer.SetServerNonce(setClientDHParams.ServerNonce)
		setClientDHParamsAnswer.SetNewNonceHash3(calcNewNonceHash(ctx.NewNonce(), ctx.ServerNonce(), 0x03))
		return setClientDHParamsAnswer.To_SetClient_DHParamsAnswer(), nil
	}

	bigIntA := new(big.Int).SetBytes(ctx.A())

	// hash_key
	authKeyNum := new(big.Int)
	authKeyNum.Exp(new(big.Int).SetBytes([]byte(clientDHInnerData.GetGB())), bigIntA, gBigIntDH2048P)

	authKey := make([]byte, 256)

	copy(authKey[256-len(authKeyNum.Bytes()):], authKeyNum.Bytes())

	authKeyAuxHash := make([]byte, len(ctx.NewNonce()))
	copy(authKeyAuxHash, ctx.NewNonce())
	authKeyAuxHash = append(authKeyAuxHash, byte(0x01))
	sha1D := sha1.Sum(authKey)
	authKeyAuxHash = append(authKeyAuxHash, sha1D[:]...)
	sha1E := sha1.Sum(authKeyAuxHash[:len(authKeyAuxHash)-12])
	authKeyAuxHash = append(authKeyAuxHash, sha1E[:]...)

	// 至此key已经创建成功
	authKeyId := int64(binary.LittleEndian.Uint64(authKeyAuxHash[len(ctx.NewNonce())+1+12 : len(ctx.NewNonce())+1+12+8]))
	ctx.SetAuthKeyId(authKeyId)
	ctx.SetAuthKey(authKey)
	log.Debugf("onSetClient_DHParams OK AuthKeyId:`%d` NewNonce:`%s` AuthKey:`%s`", authKeyId, hex.EncodeToString(ctx.NewNonce()), hex.EncodeToString(authKey))

	if !s.saveAuthKeyInfo(ctx) {
		setClientDHParamsAnswer := mtproto.NewTLDhGenRetry(nil)
		setClientDHParamsAnswer.SetNonce(setClientDHParams.Nonce)
		setClientDHParamsAnswer.SetServerNonce(setClientDHParams.ServerNonce)

		setClientDHParamsAnswer.SetNewNonceHash2(calcNewNonceHash(ctx.NewNonce(), authKey, 0x02))
		log.Errorf("handshake::onSetClient_DHParams ctx:%v setClientDHParams:%v setClientDHParamsAnswer:%v ok", ctx, setClientDHParams, setClientDHParamsAnswer)
		return setClientDHParamsAnswer.To_SetClient_DHParamsAnswer(), nil
	}

	setClientDHParamsAnswer := mtproto.NewTLDhGenOk(&mtproto.SetClient_DHParamsAnswer{
		ClassName: mtproto.ClassDhGenOk,
		Cmd:       mtproto.Cmd_DhGenOk,
	})
	setClientDHParamsAnswer.SetNonce(setClientDHParams.Nonce)
	setClientDHParamsAnswer.SetServerNonce(setClientDHParams.ServerNonce)
	setClientDHParamsAnswer.SetNewNonceHash1(calcNewNonceHash(ctx.NewNonce(), authKey, 0x01))
	return setClientDHParamsAnswer.To_SetClient_DHParamsAnswer(), nil
}

func (s *handshake) saveAuthKeyInfo(ctx *data.Context) bool {
	var (
		salt = int64(0)
		now  = int32(time.Now().Unix())
	)
	for a := 7; a >= 0; a-- {
		salt <<= 8
		salt |= int64(ctx.NewNonce()[a] ^ ctx.ServerNonce()[a])
	}
	resp, err := authClient.GetAuthClientByAuthKey(ctx.AuthKeyId()).
		ReqUpdateAuthKey(context.Background(), &authClient.UpdateAuthKey{
			AuthKeyId:   ctx.AuthKeyId(),
			AuthKey:     ctx.AuthKey(),
			TempAuthKey: ctx.AuthKeyTemp().ToInt32(),
			ExpiresIn:   ctx.ExpiresIn(),
			ValidSince:  now,
			ValidUntil:  now + SaltUtilTime,
			Salt:        salt,
		})
	if err != nil {
		log.Errorf("handshake::saveAuthKeyInfo error:%v", err.Error())
		return false
	}

	ok := &mtproto.Bool{}
	err = types.UnmarshalAny(resp, ok)
	if err != nil {
		log.Errorf("handshake::saveAuthKeyInfo resp:%+v error:%v", resp, err.Error())
		return false
	}
	return mtproto.ToBool(ok)
}

func checkSha1(data []byte, maxPaddingLen int) bool {
	var (
		dataLen  = len(data)
		sha1Data = data[:SHA_DIGEST_LENGTH]
	)

	if maxPaddingLen > 0 {
		for paddingLen := 0; paddingLen < maxPaddingLen; paddingLen++ {
			sha1Check := sha1.Sum(data[SHA_DIGEST_LENGTH : dataLen-paddingLen])
			if bytes.Equal(sha1Check[:], sha1Data) {
				return true
			}
		}
	} else {
		sha1Check := sha1.Sum(data[SHA_DIGEST_LENGTH:dataLen])
		if bytes.Equal(sha1Check[:], sha1Data) {
			return true
		}
	}
	return false
}

func calcNewNonceHash(newNonce, authKey []byte, b byte) []byte {
	authKeyAuxHash := make([]byte, len(newNonce))
	copy(authKeyAuxHash, newNonce)
	authKeyAuxHash = append(authKeyAuxHash, b)
	sha1D := sha1.Sum(authKey)
	authKeyAuxHash = append(authKeyAuxHash, sha1D[:]...)
	sha1E := sha1.Sum(authKeyAuxHash[:len(authKeyAuxHash)-12])
	authKeyAuxHash = append(authKeyAuxHash, sha1E[:]...)
	return authKeyAuxHash[len(authKeyAuxHash)-16:]
}
