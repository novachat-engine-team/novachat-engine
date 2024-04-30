package codec

import (
	"fmt"
	"novachat_engine/mtproto"
)

func NewMTPRawMessage(authKeyId int64, connType int) *MTPRawMessage {
	return &MTPRawMessage{
		connType:   connType,
		authKeyId:  authKeyId,
	}
}

////////////////////////////////////////////////////////////////////////////
// 代理使用
type MTPRawMessage struct {
	connType   int
	authKeyId  int64 // 由原始数据解压获得

	// 原始数据
	Payload []byte
}

func (m *MTPRawMessage) String() string {
	return fmt.Sprintf("{MTPRawMessage conn_type: %d, auth_key_id: %d, payload_len: %d}",
		m.connType,
		m.authKeyId,
		len(m.Payload))
}

func (m *MTPRawMessage) GoString() string {
	return fmt.Sprintf("{MTPRawMessage conn_type: %d, auth_key_id: %d, payload_len: %d}",
		m.connType,
		m.authKeyId,
		len(m.Payload))
}

func (m *MTPRawMessage) ConnType() int {
	return m.connType
}

func (m *MTPRawMessage) AuthKeyId() int64 {
	return m.authKeyId
}

////////////////////////////////////////////////////////////////////////////
func (m *MTPRawMessage) Encode() []byte {
	return m.Payload
}

func (m *MTPRawMessage) Decode(b []byte) error {
	m.Payload = b
	return nil
}

func (m *MTPRawMessage) DebugString() string {
	return fmt.Sprintf(`{"conn_type":%d,"auth_key_id":%d, "payload_len":%d}`,
		m.connType,
		m.authKeyId,
		len(m.Payload))
}

////////////////////////////////////////////////////////////////////////////
func NewUnencryptedRawMessage() *UnencryptedRawMessage {
	return &UnencryptedRawMessage{
		AuthKeyId: 0,
	}
}

type UnencryptedRawMessage struct {
	AuthKeyId   int64
	MessageId   int64
	MessageData []byte
}

func (m *UnencryptedRawMessage) Encode() []byte {
	x := mtproto.NewEncodeBuf(20 + len(m.MessageData))
	x.Long(0)
	m.MessageId = GenerateMessageId()
	x.Long(m.MessageId)
	x.Int(int32(len(m.MessageData)))
	x.Bytes(m.MessageData)
	return x.GetBuf()
}

func (m *UnencryptedRawMessage) Decode(b []byte) error {
	dbuf := mtproto.NewDecodeBuf(b)
	m.MessageId = dbuf.Long()
	messageLen := dbuf.Int()
	if int(messageLen) != dbuf.GetSize()-12 {
		return fmt.Errorf("invalid UnencryptedRawMessage len: %d (need %d)", messageLen, dbuf.GetSize()-12)
	}
	m.MessageData = dbuf.Bytes(int(messageLen))
	return dbuf.GetError()
}

type EncryptedRawMessage struct {
	AuthKeyId     int64
	MsgKey        []byte
	EncryptedData []byte
}

func NewEncryptedRawMessage(authKeyId int64) *EncryptedRawMessage {
	return &EncryptedRawMessage{
		AuthKeyId: authKeyId,
	}
}

func (m *EncryptedRawMessage) Encode() []byte {
	// 一次性分配
	x := mtproto.NewEncodeBuf(24 + len(m.EncryptedData))
	x.Long(m.AuthKeyId)
	x.Bytes(m.MsgKey)
	x.Bytes(m.EncryptedData)
	return x.GetBuf()
}

func (m *EncryptedRawMessage) Decode(b []byte) error {
	dbuf := mtproto.NewDecodeBuf(b)
	m.MsgKey = dbuf.Bytes(16)
	m.EncryptedData = dbuf.Bytes(len(b) - 16)
	return dbuf.GetError()
}
