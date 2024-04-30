/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/24 18:03
 * @File : object.codec.go
 */

package mtproto

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"io"
	"novachat_engine/pkg/log"
)

func NewTLMessageTransform() *TLMessageTransform {
	return &TLMessageTransform{}
}

func (m *TLMessageTransform) Encode(layer int32) []byte {
	objectBytes := m.Object.Encode(layer)
	if len(objectBytes) == 0 {
		_ = objectBytes
	}
	objectBytesLen := len(objectBytes)
	m.Bytes = int32(objectBytesLen)

	x := NewEncodeBuf(8 + 4 + 4 + objectBytesLen)
	x.Long(m.MsgId)
	x.Int(m.Seqno)
	x.Int(m.Bytes)
	x.Bytes(objectBytes)

	return x.GetBuf()
}

func (m *TLMessageTransform) Decode(dBuf *DecodeBuf, layer int32) error {
	m.MsgId = dBuf.Long()
	m.Seqno = dBuf.Int()
	m.Bytes = dBuf.Int()

	if err := dBuf.GetError(); err != nil {
		return err
	}
	data := dBuf.Bytes(int(m.Bytes))
	decode := NewDecodeBuf(data)
	m.Object = decode.Object(layer)
	if m.Object == nil {
		err := fmt.Errorf("TLMessageTransform decode error: %s", hex.EncodeToString(data))
		log.Error(err.Error())

		rpcErr := NewTLRpcError(nil)
		rpcErr.Data2.ErrorMessage = err.Error()
		//rpcErr := NewTLRpcError()
		//rpcErr.Data2.ErrorMessage = err.Error()
		m.Object = rpcErr
		return err
	}

	return decode.GetError()
}

func (m *TLMessageTransform) String() string {
	return fmt.Sprintf("TLMessageTransform MsgId:%d Seqno:%d BytesLen:%d", m.MsgId, m.Seqno, m.Bytes)
}

func (m *TLMessageTransform) GoString() string {
	return fmt.Sprintf("TLMessageTransform MsgId:%d Seqno:%d BytesLen:%d", m.MsgId, m.Seqno, m.Bytes)
}

func (m *TLMessageTransform) DebugString() string {
	return fmt.Sprintf("TLMessageTransform MsgId:%d Seqno:%d BytesLen:%d", m.MsgId, m.Seqno, m.Bytes)
}

func NewTLMsgContainer() *TLMsgContainer {
	return &TLMsgContainer{
		Cmd: Cmd_msg_container,
	}
}

func (m *TLMsgContainer) Encode(layer int32) []byte {
	x := NewEncodeBuf(4 + 8 + 128)
	x.UInt(m.Cmd.ToUInt32())

	x.Int(int32(len(m.Messages)))
	for _, d := range m.Messages {
		x.Bytes(d.Encode(layer))
	}

	return x.GetBuf()
}

func (m *TLMsgContainer) Decode(dBuf *DecodeBuf, layer int32) error {
	var err error
	msgLen := dBuf.Int()
	m.Messages = make([]*TLMessageTransform, 0, msgLen)
	for i := int32(0); i < msgLen; i++ {
		message := NewTLMessageTransform()
		err = message.Decode(dBuf, layer)
		if err != nil {
			log.Errorf("TLMsgContainer::Decode error:%v", err.Error())
			return err
		}
		m.Messages = append(m.Messages, message)
	}

	return dBuf.GetError()
}

func (m *TLMsgContainer) GoString() string {
	return m.DebugString()
}

func (m *TLMsgContainer) String() string {
	return m.DebugString()
}

func (m *TLMsgContainer) DebugString() string {
	return fmt.Sprintf("TLMsgContainer  Messages len:%d", len(m.Messages))
}

func (m *TLMsgCopy) Encode(layer int32) []byte {
	x := NewEncodeBuf(4 + 8 + 128)
	x.UInt(m.Cmd.ToUInt32())

	x.Bytes(m.OrigMessage.Encode(layer))

	return x.GetBuf()
}

func (m *TLMsgCopy) Decode(dBuf *DecodeBuf, layer int32) error {
	m.OrigMessage = NewTLMessageTransform()
	return m.OrigMessage.Decode(dBuf, layer)
}

func (m *TLMsgCopy) GoString() string {
	return m.DebugString()
}

func (m *TLMsgCopy) String() string {
	return m.DebugString()
}

func (m *TLMsgCopy) DebugString() string {
	return fmt.Sprintf("TLMsgCopy OrigMessage:%+v", m.OrigMessage)
}

func (m *TLGZipPacked) Encode(layer int32) []byte {
	x := NewEncodeBuf(4 + 8 + 128)
	x.UInt(m.Cmd.ToUInt32())

	x.Bytes(m.PackedData)

	return x.GetBuf()
}

func (m *TLGZipPacked) Decode(dBuf *DecodeBuf, layer int32) error {
	data := dBuf.StringBytes()
	if dBuf.err != nil {
		return dBuf.err
	}
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		dBuf.err = err
		return fmt.Errorf("TLGZipPacked gzip read: %v", err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		dBuf.err = err
		return fmt.Errorf("gzip read: %v", err)
	}
	if clErr != nil {
		dBuf.err = clErr
		return clErr
	}

	m.PackedData = buf.Bytes()
	return dBuf.GetError()
}

func (m *TLGZipPacked) GoString() string {
	return m.DebugString()
}

func (m *TLGZipPacked) String() string {
	return m.DebugString()
}

func (m *TLGZipPacked) DebugString() string {
	return fmt.Sprintf("TLGZipPacked PackedData len:%d", len(m.PackedData))
}

func (m *TLInvokeWithLayer) Encode(layer int32) []byte {
	d := NewEncodeBuf(256)
	d.UInt(m.Cmd.ToUInt32())
	d.Int(m.Layer)
	d.Bytes(m.Query.Encode(layer))
	return d.GetBuf()
}

func (m *TLInvokeWithLayer) Decode(dBuf *DecodeBuf, layer int32) error {
	m.Layer = dBuf.Int()
	m.Query = dBuf.Object(layer)
	return dBuf.GetError()
}

func (m *TLInvokeWithLayer) Reset()      { *m = TLInvokeWithLayer{} }
func (*TLInvokeWithLayer) ProtoMessage() {}

func (m *TLInvokeWithLayer) GoString() string {
	return m.DebugString()
}

func (m *TLInvokeWithLayer) String() string {
	return m.DebugString()
}

func (m *TLInvokeWithLayer) DebugString() string {
	return proto.CompactTextString(m)
}

func (m *TLInvokeWithoutUpdates) Encode(layer int32) []byte {
	d := NewEncodeBuf(256)
	d.UInt(m.Cmd.ToUInt32())
	d.Bytes(m.Query.Encode(layer))
	return d.GetBuf()
}

func (m *TLInvokeWithoutUpdates) Decode(dBuf *DecodeBuf, layer int32) error {
	m.Query = dBuf.Object(layer)
	return dBuf.GetError()
}
func (m *TLInvokeWithoutUpdates) Reset()      { *m = TLInvokeWithoutUpdates{} }
func (*TLInvokeWithoutUpdates) ProtoMessage() {}

func (m *TLInvokeWithoutUpdates) GoString() string {
	return m.DebugString()
}

func (m *TLInvokeWithoutUpdates) String() string {
	return m.DebugString()
}

func (m *TLInvokeWithoutUpdates) DebugString() string {
	return proto.CompactTextString(m)
}

func (m *TLInvokeWithMessagesRange) Encode(layer int32) []byte {
	d := NewEncodeBuf(256)
	d.UInt(m.Cmd.ToUInt32())
	d.Bytes(m.MessageRange.Encode(layer))
	d.Bytes(m.Query.Encode(layer))
	return d.GetBuf()
}

func (m *TLInvokeWithMessagesRange) Decode(dBuf *DecodeBuf, layer int32) error {
	m.MessageRange = dBuf.Object(layer).(*TLMessageRange)
	m.Query = dBuf.Object(layer)
	return dBuf.GetError()
}
func (m *TLInvokeWithMessagesRange) Reset()      { *m = TLInvokeWithMessagesRange{} }
func (*TLInvokeWithMessagesRange) ProtoMessage() {}

func (m *TLInvokeWithMessagesRange) GoString() string {
	return m.DebugString()
}

func (m *TLInvokeWithMessagesRange) String() string {
	return m.DebugString()
}

func (m *TLInvokeWithMessagesRange) DebugString() string {
	return proto.CompactTextString(m)
}

func (m *TLInvokeWithTakeout) Encode(layer int32) []byte {
	d := NewEncodeBuf(256)
	d.UInt(m.Cmd.ToUInt32())
	d.Long(m.TakeoutId)
	d.Bytes(m.Query.Encode(layer))
	return d.GetBuf()
}

func (m *TLInvokeWithTakeout) Decode(dBuf *DecodeBuf, layer int32) error {
	m.TakeoutId = dBuf.Long()
	m.Query = dBuf.Object(layer)
	return dBuf.GetError()
}

func (m *TLInvokeWithTakeout) Reset()      { *m = TLInvokeWithTakeout{} }
func (*TLInvokeWithTakeout) ProtoMessage() {}

func (m *TLInvokeWithTakeout) GoString() string {
	return m.DebugString()
}

func (m *TLInvokeWithTakeout) String() string {
	return m.DebugString()
}

func (m *TLInvokeWithTakeout) DebugString() string {
	return proto.CompactTextString(m)
}

func NewTLInitConnection() *TLInitConnection {
	return &TLInitConnection{
		Cmd:       Cmd_init_connection,
		ClassName: Class_init_connection,
	}
}

func (m *TLInitConnection) Unmarshal(d []byte) error {
	return json.Unmarshal(d, m)
}

func (m *TLInitConnection) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *TLInitConnection) encode_c7481da6(d *EncodeBuf, layer int32) []byte {
	d.Int(m.ApiId)
	d.String(m.DeviceModel)
	d.String(m.SystemVersion)
	d.String(m.AppVersion)
	d.String(m.SystemLangCode)
	d.String(m.LangPack)
	d.String(m.LangCode)
	d.Bytes(m.Query.Encode(layer))
	return d.GetBuf()
}

func (m *TLInitConnection) Encode(layer int32) []byte {
	d := NewEncodeBuf(256)

	switch m.Cmd {
	case Cmd_init_connectionc7481da6:
		return m.encode_c7481da6(d, layer)
	case Cmd_init_connection:
		flag := uint32(0)
		if m.Proxy != nil {
			flag |= 1 << 0
		}
		if m.JSONValue != nil {
			flag |= 1 << 1
		}
		d.UInt(m.Cmd.ToUInt32())
		d.UInt(flag)
		d.Int(m.ApiId)
		d.String(m.DeviceModel)
		d.String(m.SystemVersion)
		d.String(m.AppVersion)
		d.String(m.SystemLangCode)
		d.String(m.LangPack)
		d.String(m.LangCode)
		if m.Proxy != nil {
			d.Bytes(m.Proxy.Encode(layer))
		}
		if m.JSONValue != nil {
			d.Bytes(m.JSONValue.Encode(layer))
		}
		d.Bytes(m.Query.Encode(layer))
		break
	}

	return d.GetBuf()
}

func (m *TLInitConnection) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_init_connectionc7481da6:
		m.ApiId = dBuf.Int()
		m.DeviceModel = dBuf.String()
		m.SystemVersion = dBuf.String()
		m.AppVersion = dBuf.String()
		m.SystemLangCode = dBuf.String()
		m.LangPack = dBuf.String()
		m.LangCode = dBuf.String()
		m.Query = dBuf.Object(layer)
		break
	case Cmd_init_connection:
		flags := dBuf.Int()

		m.Flags = flags
		m.ApiId = dBuf.Int()
		m.DeviceModel = dBuf.String()
		m.SystemVersion = dBuf.String()
		m.AppVersion = dBuf.String()
		m.SystemLangCode = dBuf.String()
		m.LangPack = dBuf.String()
		m.LangCode = dBuf.String()

		if (flags & (1 << 0)) != 0 {
			m.Proxy = dBuf.Object(layer).(*TLInputClientProxy)
		}
		if (flags & (1 << 1)) != 0 {
			m.JSONValue = dBuf.Object(layer).(*TLJsonObject).To_JSONValue()
		}
		m.Query = dBuf.Object(layer)
	}
	return dBuf.GetError()
}
func (m *TLInitConnection) Reset()      { *m = TLInitConnection{} }
func (*TLInitConnection) ProtoMessage() {}

func (m *TLInitConnection) GoString() string {
	return m.DebugString()
}

func (m *TLInitConnection) String() string {
	return m.DebugString()
}

func (m *TLInitConnection) DebugString() string {
	return proto.CompactTextString(m)
}

func (m *TLInvokeAfterMsgs) Encode(layer int32) []byte {
	d := NewEncodeBuf(256)

	switch m.Cmd {
	case Cmd_invokeAfterMsgs:
		d.VectorLong(m.MsgIds)
		d.Bytes(m.Encode(layer))
		break
	default:
		log.Errorf("invalid TLInvokeAfterMsgs encode layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}

	return d.GetBuf()
}

func (m *TLInvokeAfterMsgs) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_invokeAfterMsgs:
		m.MsgIds = dBuf.VectorLong()
		m.Query = dBuf.Object(layer)
		break
	default:
		log.Errorf("TLInvokeAfterMsgs Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}
func (m *TLInvokeAfterMsgs) Reset()      { *m = TLInvokeAfterMsgs{} }
func (*TLInvokeAfterMsgs) ProtoMessage() {}

func (m *TLInvokeAfterMsgs) GoString() string {
	return m.DebugString()
}

func (m *TLInvokeAfterMsgs) String() string {
	return m.DebugString()
}

func (m *TLInvokeAfterMsgs) DebugString() string {
	return proto.CompactTextString(m)
}

func (m *TLInvokeAfterMsg) Encode(layer int32) []byte {
	d := NewEncodeBuf(256)

	switch m.Cmd {
	case Cmd_invokeAfterMsg:
		d.Long(m.MsgId)
		d.Bytes(m.Encode(layer))
		break
	default:
		log.Errorf("invalid TLInvokeAfterMsg encode layer:%d cmd: %x", layer, m.Cmd.ToUInt32())
		return nil
	}

	return d.GetBuf()
}

func (m *TLInvokeAfterMsg) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_invokeAfterMsg:
		m.MsgId = dBuf.Long()
		m.Query = dBuf.Object(layer)
		break
	default:
		log.Errorf("TLInvokeAfterMsg Decode Invalid error cmd:%x", m.Cmd.ToUInt32())
	}
	return dBuf.GetError()
}
func (m *TLInvokeAfterMsg) Reset()      { *m = TLInvokeAfterMsg{} }
func (*TLInvokeAfterMsg) ProtoMessage() {}

func (m *TLInvokeAfterMsg) GoString() string {
	return m.DebugString()
}

func (m *TLInvokeAfterMsg) String() string {
	return m.DebugString()
}

func (m *TLInvokeAfterMsg) DebugString() string {
	return proto.CompactTextString(m)
}

func (m *TLRpcResult) String() string {
	return "{rpc_result#f35c6d01: req_msg_id:" + string(m.ReqMsgId) + "}"
}

func (m *TLRpcResult) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_rpc_result))
	x.Long(m.ReqMsgId)
	x.Bytes(m.Result.Encode(layer))
	return x.buf
}

func (m *TLRpcResult) Decode(dBuf *DecodeBuf, layer int32) error {
	m.ReqMsgId = dBuf.Long()
	m.Result = dBuf.Object(layer)
	return dBuf.GetError()
}

func (m *TLRpcResult) DebugString() string {
	return fmt.Sprintf(`{"rpc_result#f35c6d01": {"req_msg_id": %d}}`, m.ReqMsgId)
}

func (m *TLRpcResult) Reset() { *m = TLRpcResult{} }

func (*TLRpcResult) ProtoMessage() {}

func (m *TLRpcResult) GoString() string {
	return m.DebugString()
}
