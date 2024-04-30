/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/24 18:03
 * @File : object.codec.go
 */

package mtproto

import "fmt"

const CurrentLayer = int32(122)

var DefaultRpcError = &TLRpcError{}

type TLObject interface {
	Encode(layer int32) []byte
	Decode(dBuf *DecodeBuf, layer int32) error
	GoString() string
	String() string
}

var boolTrue = &Bool{
	ClassName: ClassBoolTrue,
	Cmd:       Cmd_BoolTrue,
}

var boolFalse = &Bool{
	ClassName: ClassBoolFalse,
	Cmd:       Cmd_BoolFalse,
}

func ToMTBool(b bool) *Bool {
	if b {
		return boolTrue
	} else {
		return boolFalse
	}
}

func ToBool(b *Bool) bool {
	return ClassBoolTrue == b.GetClassName()
}

func (m TLCmd) ToUInt32() uint32 {
	return uint32(m)
}

func CmdFromUInt32(m uint32) TLCmd {
	return TLCmd(m)
}

type MessageTransform struct {
	Message TLMessageTransform
}

//message msg_id:long seqno:int bytes:int body:Object = Message; // parsed manually
type TLMessageTransform struct {
	MsgId  int64
	Seqno  int32
	Bytes  int32
	Object TLObject
}

//msg_container#73f1f8dc messages:vector<message> = MessageContainer; // parsed manually
type TLMsgContainer struct {
	Cmd       TLCmd
	ClassName string
	Messages  []*TLMessageTransform
}

//msg_copy#e06046b2 orig_message:Message = MessageCopy; // parsed manually, not used - use msg_container
type TLMsgCopy struct {
	Cmd         TLCmd
	ClassName   string
	OrigMessage *TLMessageTransform
}

//gzip_packed#3072cfa1 packed_data:string = Object; // parsed manually
type TLGZipPacked struct {
	Cmd        TLCmd
	ClassName  string
	PackedData []byte
}

//invokeAfterMsg#cb9f372d {X:Type} msg_id:long query:!X = X;
type TLInvokeAfterMsg struct {
	Cmd       TLCmd
	ClassName string
	MsgId     int64
	Query     TLObject
}

//invokeAfterMsgs#3dc4b4f0 {X:Type} msg_ids:Vector<long> query:!X = X;
type TLInvokeAfterMsgs struct {
	Cmd       TLCmd
	ClassName string
	MsgIds    []int64
	Query     TLObject
}

//initConnection#c1cd5ea9 {X:Type} flags:# api_id:int device_model:string system_version:string app_version:string system_lang_code:string lang_pack:string lang_code:string
//proxy:flags.0?InputClientProxy params:flags.1?JSONValue query:!X = X;
type TLInitConnection struct {
	Cmd            TLCmd
	ClassName      string
	ApiId          int32
	DeviceModel    string
	SystemVersion  string
	AppVersion     string
	SystemLangCode string
	LangPack       string
	LangCode       string
	Proxy          *TLInputClientProxy
	JSONValue      *JSONValue
	Query          TLObject
	Flags          int32
}

//invokeWithLayer#da9b0d0d {X:Type} layer:int query:!X = X;
type TLInvokeWithLayer struct {
	Cmd       TLCmd
	ClassName string
	Layer     int32
	Query     TLObject
}

//invokeWithoutUpdates#bf9459b7 {X:Type} query:!X = X;
type TLInvokeWithoutUpdates struct {
	Cmd       TLCmd
	ClassName string
	Query     TLObject
}

//invokeWithMessagesRange#365275f2 {X:Type} range:MessageRange query:!X = X;
type TLInvokeWithMessagesRange struct {
	Cmd          TLCmd
	ClassName    string
	MessageRange *TLMessageRange
	Query        TLObject
}

//invokeWithTakeout#aca9fd2e {X:Type} takeout_id:long query:!X = X;
type TLInvokeWithTakeout struct {
	Cmd       TLCmd
	ClassName string
	TakeoutId int64
	Query     TLObject
}

//rpc_result#f35c6d01 req_msg_id:long result:Object = RpcResult; // parsed manually
type TLRpcResult struct {
	ReqMsgId int64
	Result   TLObject
}

type Transform struct {
	SessionId int64  `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Salt      int64  `protobuf:"varint,3,opt,name=salt,proto3" json:"salt,omitempty"`
	Buf       []byte `protobuf:"bytes,4,opt,name=buf,proto3" json:"buf,omitempty"`
}

func (m *Transform) Encode(layer int32) []byte {
	encoder := NewEncodeBuf(8 + 8 + len(m.Buf))
	encoder.Long(m.Salt)
	encoder.Long(m.SessionId)
	encoder.Bytes(m.Buf)
	return encoder.GetBuf()
}

func (m *Transform) Decode(d []byte, layer int32) error {
	decoder := NewDecodeBuf(d)
	m.Salt = decoder.Long()
	err := decoder.GetError()
	if err != nil {
		fmt.Errorf(" Transform Salt error:%s", err.Error())
		return err
	}
	m.SessionId = decoder.Long()
	err = decoder.GetError()
	if err != nil {
		fmt.Errorf(" Transform SessionId error:%s", err.Error())
		return err
	}
	m.Buf = decoder.Bytes(decoder.GetSize() - decoder.GetOffset())
	err = decoder.GetError()
	if err != nil {
		fmt.Errorf(" Transform Buf error:%s", err.Error())
		return err
	}
	return nil
}
