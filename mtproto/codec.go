/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/8 15:21
 * @File : codec.go
 */

package mtproto

import (
	"math"
	"strings"
)

const (
	MtProtoCurrentLayer = 122
)

func SetFlag(flag bool, bitIndex int32) uint32 {
	if flag == true {
		return 1 << bitIndex
	} else {
		return 0
	}
}

func GetFlag(flag uint32, bitIndex int32) bool {
	return flag&(1<<bitIndex) != 0
}

func doubleIsEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}

func (x RpcErrorCode) ToInt32() int32 {
	return int32(x)
}

func RpcErrorCodeFromInt32(v int32) RpcErrorCode {
	return RpcErrorCode(v)
}

func (x RpcErrorCode) GetClientCodePost() int32 {
	return x.ToInt32() % RpcErrorCode_BASE_NUMBER.ToInt32()
}

func (x RpcErrorCode) GetClientCode() int32 {
	return x.ToInt32() / RpcErrorCode_BASE_NUMBER.ToInt32()
}

func (x RpcErrorCode) GetClientErrorString() string {
	baseValue := x.ToInt32() - x.ToInt32()%RpcErrorCode_BASE_NUMBER.ToInt32()
	if x.ToInt32() == baseValue {
		baseString := RpcErrorCode_name[x.ToInt32()]
		idx := strings.Index(baseString, "_")
		if idx == -1 {
			return baseString
		}
		return baseString[idx+1:]
	}

	baseString := RpcErrorCode_name[baseValue]
	codeMsg := RpcErrorCode_name[x.ToInt32()]

	idx := strings.Index(codeMsg, baseString)
	if idx == -1 {
		return RpcErrorCode_name[RpcErrorCode_ErrorFormatUNKNOWN.ToInt32()]
	}

	return codeMsg[idx+len(baseString)+1:]
}

func (m *InitConnection) Encode(layer int32) []byte {
	d := NewEncodeBuf(0)
	return d.GetBuf()
}

func (m *InitConnection) Decode(dBuf *DecodeBuf, layer int32) error {
	return dBuf.GetError()
}

func (m *TLRpcError) Error() string {
	return m.GetErrorMessage()
}

func (m *TLRpcError) Unwrap() error {
	return nil
}

func (m *TLRpcError) Is(e error) bool {
	return m == e
}

func (m *TLRpcError) As(e interface{}) bool {
	_, ok := e.(*TLRpcError)
	return ok
}

//func (m *VectorString) Encode(layer int32) []byte {
//	x := NewEncodeBuf(512)
//	x.VectorLong(m.Datas)
//	return x.buf
//}
//
//func (m *VectorString) Decode(dbuf *DecodeBuf) error {
//	m.Datas = dbuf.VectorString()
//
//	return dbuf.err
//}
