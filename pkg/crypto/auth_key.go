/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/2/21 16:13
 * @File :
 */


package crypto

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"novachat_engine/pkg/log"
)

const (
	mtprotoVersion = 2
	serverSide     = 0
	clientSide     = 1
)

type AuthKey struct {
	authKeyId  int64
	authKey    []byte
	side       int // client or server
	mtpVersion int // 2
}

func NewAuthKey(keyId int64, keyData []byte) *AuthKey {
	// check keyId
	return &AuthKey{
		authKeyId:  keyId,
		authKey:    keyData,
		side:       serverSide,
		mtpVersion: mtprotoVersion,
	}
}

func NewClientAuthKey(keyId int64, keyData []byte) *AuthKey {
	// check keyId
	return &AuthKey{
		authKeyId:  keyId,
		authKey:    keyData,
		side:       clientSide,
		mtpVersion: mtprotoVersion,
	}
}

func NewAuthKey1(keyId int64, keyData []byte) *AuthKey {
	// check keyId
	return &AuthKey{
		authKeyId:  keyId,
		authKey:    keyData,
		side:       clientSide,
		mtpVersion: clientSide,
	}
}

func (k *AuthKey) AuthKeyId() int64 {
	return k.authKeyId
}

func (k *AuthKey) AuthKey() []byte {
	return k.authKey
}

func (k *AuthKey) Equals(o *AuthKey) bool {
	return bytes.Equal(k.authKey, o.authKey)
}

func (k *AuthKey) CalcAuthKeyId() int64 {
	sha1 := Sha1Digest(k.authKey)

	// Lower 64 bits = 8 bytes of 20 byte SHA1 hash.
	return int64(binary.LittleEndian.Uint64(sha1[12:]))
}

func (k *AuthKey) calcX(incoming bool) int {
	var x = 0
	if k.side == serverSide {
		if incoming {
			x = 8
		}
	} else {
		if !incoming {
			x = 8
		}
	}
	return x
}

func (k *AuthKey) prepareAES(msgKey []byte, incoming bool) ([]byte, []byte) {
	x := k.calcX(incoming)

	aesKey := make([]byte, 32)
	aesIV := make([]byte, 32)

	switch k.mtpVersion {
	case mtprotoVersion:
		dataA := make([]byte, 16+36)
		copy(dataA, msgKey[:16])
		copy(dataA[16:], k.authKey[x:x+36])
		sha256A := Sha256Digest(dataA)

		dataB := make([]byte, 36+16)
		copy(dataB, k.authKey[40+x:40+x+36])
		copy(dataB[36:], msgKey[:16])
		sha256B := Sha256Digest(dataB)

		copy(aesKey, sha256A[:8])
		copy(aesKey[8:], sha256B[8:8+16])
		copy(aesKey[8+16:], sha256A[24:24+8])
		copy(aesIV, sha256B[:8])
		copy(aesIV[8:], sha256A[8:8+16])
		copy(aesIV[8+16:], sha256B[24:24+8])
	default:
		dataA := make([]byte, 16+32)
		copy(dataA, msgKey[:16])
		copy(dataA[16:], k.authKey[x:x+32])
		sha1A := Sha1Digest(dataA)

		dataB := make([]byte, 16+16+16)
		copy(dataB, k.authKey[32+x:32+x+16])
		copy(dataB[16:], msgKey[:16])
		copy(dataB[32:], k.authKey[48+x:48+x+16])
		sha1B := Sha1Digest(dataB)

		dataC := make([]byte, 32+16)
		copy(dataC, k.authKey[64+x:64+x+32])
		copy(dataC[32:], msgKey[:16])
		sha1C := Sha1Digest(dataC)

		dataD := make([]byte, 16+32)
		copy(dataD, msgKey[:16])
		copy(dataD[16:], k.authKey[96+x:96+x+32])
		sha1D := Sha1Digest(dataD)

		copy(aesKey, sha1A[:8])
		copy(aesKey[8:], sha1B[8:8+12])
		copy(aesKey[8+12:], sha1C[4:4+12])
		copy(aesIV, sha1A[8:8+12])
		copy(aesIV[12:], sha1B[:8])
		copy(aesIV[12+8:], sha1C[16:16+4])
		copy(aesIV[12+8+4:], sha1D[:8])
	}

	return aesKey, aesIV
}

func (k *AuthKey) partForMsgKey(incoming bool) []byte {
	x := k.calcX(incoming)
	return k.authKey[88+x : 88+x+32]
}

/*
	| salt <br> int64	| `session_id` <br> int64 | `message_id` <br> int64 | `seq_no` <br> int32 |`message_data_length` <br> int32	| `message_data` <br> bytes | padding12..1024 <br> bytes|
	|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
*/
func (k *AuthKey) AesIgeEncrypt(rawData []byte) ([]byte, []byte, error) {
	var additionalSize = len(rawData) % 16
	if additionalSize != 0 {
		additionalSize = 16 - additionalSize
	}

	if k.mtpVersion == mtprotoVersion && additionalSize < 12 {
		additionalSize += 16
	}

	// var tmpData []byte
	// if additionalSize >
	tmpData := make([]byte, 0, len(rawData)+additionalSize)
	tmpData = append(tmpData, rawData...)
	if additionalSize > 0 {
		tmpData = append(tmpData, GenerateNonce(additionalSize)...)
	}

	// calc msg_key
	msgKey := make([]byte, 32)
	switch k.mtpVersion {
	case mtprotoVersion:
		sha256Dig := sha256.New()
		sha256Dig.Write(k.partForMsgKey(true))
		sha256Dig.Write(tmpData)
		copy(msgKey, sha256Dig.Sum(nil))
	default:
		copy(msgKey[4:], Sha1Digest(rawData))
	}

	aesKey, aesIV := k.prepareAES(msgKey[8:8+16], true)
	e := NewAES256IGECryptor(aesKey, aesIV)

	x, err := e.Encrypt(tmpData)
	if err != nil {
		log.Errorf("aesIgeEncrypt data error: %v", err)
		return nil, nil, err
	}

	return msgKey[8 : 8+16], x, nil
}

func (k *AuthKey) AesIgeDecrypt(msgKey, rawData []byte) ([]byte, error) {
	aesKey, aesIV := k.prepareAES(msgKey, false)
	d := NewAES256IGECryptor(aesKey, aesIV)
	x, err := d.Decrypt(rawData)
	if err != nil {
		log.Errorf("aesIgeDecrypt data error: %v", err)
		return nil, err
	}

	//// 校验解密后的数据合法性
	var dataLen = uint32(len(rawData))
	messageLen := binary.LittleEndian.Uint32(x[28:])
	// log.Info("descrypt - messageLen = ", messageLen)
	if messageLen+32 > dataLen {
		// 	return fmt.Errorf("Message len: %d (need less than %d)", messageLen, dbuf.size-32)
		err = fmt.Errorf("aesIgeDecrypt data error - Wrong message length %d", messageLen)
		log.Error(err.Error())
		return nil, err
	}

	calcMsgKey := make([]byte, 96)
	switch k.mtpVersion {
	case 2:
		sha256Dig := sha256.New()
		sha256Dig.Write(k.partForMsgKey(false))
		sha256Dig.Write(x[:dataLen])
		copy(calcMsgKey, sha256Dig.Sum(nil))
	default:
		copy(calcMsgKey[4:], Sha1Digest(x[:32+messageLen]))
	}

	if !bytes.Equal(calcMsgKey[8:8+16], msgKey[:16]) {
		err = fmt.Errorf("aesIgeDecrypt data error - (data: %s, aesKey: %s, aseIV: %s, authKeyId: %d, authKey: %s), msgKey verify error, sign: %s, msgKey: %s",
			hex.EncodeToString(rawData[:64]),
			hex.EncodeToString(aesKey),
			hex.EncodeToString(aesIV),
			k.authKeyId,
			hex.EncodeToString(k.authKey[88:88+32]),
			hex.EncodeToString(calcMsgKey[8:8+16]),
			hex.EncodeToString(msgKey[:16]))
		log.Error(err.Error())
		return nil, err
	}

	return x, nil
}

func (k *AuthKey) AesIgeDecrypt1(msgKey, rawData []byte) ([]byte, error) {
	aesKey, aesIV := k.prepareAES(msgKey, true)
	d := NewAES256IGECryptor(aesKey, aesIV)
	x, err := d.Decrypt(rawData)
	if err != nil {
		log.Errorf("aesIgeDecrypt data error: %v", err)
		return nil, err
	}

	//// 校验解密后的数据合法性
	var dataLen = uint32(len(rawData))
	messageLen := binary.LittleEndian.Uint32(x[28:])
	// log.Info("descrypt - messageLen = ", messageLen)
	if messageLen+32 > dataLen {
		// 	return fmt.Errorf("Message len: %d (need less than %d)", messageLen, dbuf.size-32)
		err = fmt.Errorf("aesIgeDecrypt data error - Wrong message length %d", messageLen)
		log.Error(err.Error())
		return nil, err
	}

	calcMsgKey := make([]byte, 96)
	switch k.mtpVersion {
	case 2:
		sha256Dig := sha256.New()
		sha256Dig.Write(k.partForMsgKey(false))
		sha256Dig.Write(x[:dataLen])
		copy(calcMsgKey, sha256Dig.Sum(nil))
	default:
		copy(calcMsgKey[4:], Sha1Digest(x[:32+messageLen]))
	}

	if !bytes.Equal(calcMsgKey[8:8+16], msgKey[:16]) {
		err = fmt.Errorf("aesIgeDecrypt data error - (data: %s, aesKey: %s, aseIV: %s, authKeyId: %d, authKey: %s), msgKey verify error, sign: %s, msgKey: %s",
			hex.EncodeToString(rawData[:64]),
			hex.EncodeToString(aesKey),
			hex.EncodeToString(aesIV),
			k.authKeyId,
			hex.EncodeToString(k.authKey[88:88+32]),
			hex.EncodeToString(calcMsgKey[8:8+16]),
			hex.EncodeToString(msgKey[:16]))
		log.Error(err.Error())
		return nil, err
	}

	return x, nil
}