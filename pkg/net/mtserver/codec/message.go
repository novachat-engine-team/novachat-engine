package codec

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/log"
	"reflect"
)

const (
	QUICK_ACKID = iota
	UNENCRYPTED_MESSAGE
	ENCRYPTED_MESSAGE
)

type QuickAckMessage struct {
	ackId int32
}

func (m *QuickAckMessage) MessageType() int {
	return QUICK_ACKID
}

func (m *QuickAckMessage) encode() ([]byte, error) {
	return nil, nil
}

func (m *QuickAckMessage) decode(b []byte) error {
	if len(b) != 4 {
		return fmt.Errorf("message len: %d (need 4)", len(b))
	}
	m.ackId = int32(binary.LittleEndian.Uint32(b))
	return nil
}

type UnencryptedMessage struct {
	NeedAck bool

	// authKeyId int64
	MessageId int64
	// messageDataLength int32
	// messageData []byte

	// classID int32
	Object mtproto.TLObject
}

func (m *UnencryptedMessage) MessageType() int {
	return UNENCRYPTED_MESSAGE
}

func (m *UnencryptedMessage) Encode(layer int32) []byte {
	buf, _ := m.encode(layer)
	return buf
}

func (m *UnencryptedMessage) encode(layer int32) ([]byte, error) {

	x := mtproto.NewEncodeBuf(512)
	x.Long(0)
	m.MessageId = GenerateMessageId()
	x.Long(m.MessageId)

	if m.Object == nil {
		x.Int(0)
	} else {
		b := m.Object.Encode(layer)
		x.Int(int32(len(b)))
		x.Bytes(b)
	}
	return x.GetBuf(), nil
}

func (m *UnencryptedMessage) Decode(b []byte, layer int32) error {
	return m.decode(b, layer)
}

func (m *UnencryptedMessage) decode(b []byte, layer int32) error {
	dbuf := mtproto.NewDecodeBuf(b)
	// m.authKeyId = dbuf.Long()
	m.MessageId = dbuf.Long()

	// log.Info("messageId:", m.messageId)
	// mod := m.messageId & 3
	// if mod != 1 && mod != 3 {
	// 	return fmt.Errorf("Wrong bits of message_id: %d", mod)
	// }

	messageLen := dbuf.Int()
	if messageLen < 4 {
		return fmt.Errorf("message len(%d) < 4", messageLen)
	}
	// log.Info("messageLen:", m.messageId)

	if int(messageLen) != dbuf.GetSize()-12 {
		return fmt.Errorf("message len: %d (need %d)", messageLen, dbuf.GetSize()-12)
	}

	m.Object = dbuf.Object(layer)
	if m.Object == nil {
		return fmt.Errorf("decode object is nil")
	}

	// proto.Message()
	// log.Info("Recved object: ", m.Object.(proto.Message).String())
	return dbuf.GetError()
}

////////////////////////////////////////////////////////////////////////////////////////////
// MsgDetailedInfo
type MsgDetailedInfoContainer struct {
	Message *EncryptedMessage2
}

type EncryptedMessage2 struct {
	authKeyId int64
	NeedAck   bool
	msgKey    []byte
	Salt      int64
	SessionId int64

	MessageId int64
	SeqNo     int32
	Object    mtproto.TLObject
}

func NewEncryptedMessage2(authKeyId int64) *EncryptedMessage2 {
	return &EncryptedMessage2{
		authKeyId: authKeyId,
	}
}

func (m *EncryptedMessage2) String() string {
	return fmt.Sprintf("{auth_key_id: %d. salt: %d, session_id: %d, message_id: %d, seq_no: %d, object: %v}",
		m.authKeyId, m.Salt, m.SessionId, m.MessageId, m.SeqNo, reflect.TypeOf(m.Object))
}

func (m *EncryptedMessage2) MessageType() int {
	return ENCRYPTED_MESSAGE
}

func (m *EncryptedMessage2) Encode(authKeyId int64, authKey []byte) ([]byte, error) {
	buf, err := m.encode(authKeyId, authKey)
	return buf, err
}

//func (m *EncryptedMessage2) EncodeToLayer(authKeyId int64, authKey []byte, layer int) ([]byte, error) {
//	buf, err := m.encodeToLayer(authKeyId, authKey, layer)
//	return buf, err
//}
//
//func (m *EncryptedMessage2) encodeToLayer(authKeyId int64, authKey []byte, layer int) ([]byte, error) {
//	objData := m.Object.EncodeToLayer(layer)
//	var additionalSize = (32 + len(objData)) % 16
//	if additionalSize != 0 {
//		additionalSize = 16 - additionalSize
//	}
//	if MTPROTO_VERSION == 2 && additionalSize < 12 {
//		additionalSize += 16
//	}
//
//	x := NewEncodeBuf(32 + len(objData) + additionalSize)
//	// x.Long(authKeyId)
//	// msgKey := make([]byte, 16)
//	// x.Bytes(msgKey)
//	x.Long(m.Salt)
//	x.Long(m.SessionId)
//	if m.MessageId == 0 {
//		m.MessageId = GenerateMessageId()
//	}
//	x.Long(m.MessageId)
//	x.Int(m.SeqNo)
//	x.Int(int32(len(objData)))
//	x.Bytes(objData)
//	x.Bytes(crypto.GenerateNonce(additionalSize))
//
//	// log.Info("Encode object: ", m.Object)
//
//	encryptedData, _ := m.encrypt(authKey, x.buf, len(objData))
//	x2 := NewEncodeBuf(56 + len(objData) + additionalSize)
//	x2.Long(authKeyId)
//	x2.Bytes(m.msgKey)
//	x2.Bytes(encryptedData)
//
//	// log.Info(x2.buf)
//	return x2.buf, nil
//}

func (m *EncryptedMessage2) encode(authKeyId int64, authKey []byte) ([]byte, error) {
	//objData := m.Object.Encode(0)
	//var additionalSize = (32 + len(objData)) % 16
	//if additionalSize != 0 {
	//	additionalSize = 16 - additionalSize
	//}
	//if MTPROTO_VERSION == 2 && additionalSize < 12 {
	//	additionalSize += 16
	//}
	//
	//x := NewEncodeBuf(32 + len(objData) + additionalSize)
	//// x.Long(authKeyId)
	//// msgKey := make([]byte, 16)
	//// x.Bytes(msgKey)
	//x.Long(m.Salt)
	//x.Long(m.SessionId)
	//if m.MessageId == 0 {
	//	m.MessageId = GenerateMessageId()
	//}
	//x.Long(m.MessageId)
	//x.Int(m.SeqNo)
	//x.Int(int32(len(objData)))
	//x.Bytes(objData)
	//x.Bytes(crypto.GenerateNonce(additionalSize))
	//
	//// log.Infof("Encode object: ", m.Object)
	//
	//encryptedData, _ := m.encrypt(authKey, x.buf, len(objData))
	//x2 := NewEncodeBuf(56 + len(objData) + additionalSize)
	//x2.Long(authKeyId)
	//x2.Bytes(m.msgKey)
	//x2.Bytes(encryptedData)
	//
	//// log.Info(x2.buf)
	//return x2.buf, nil
	return nil, nil
}

func (m *EncryptedMessage2) Decode(authKeyId int64, authKey, b []byte) error {
	_ = authKeyId
	return m.decode(authKey, b)
}

func (m *EncryptedMessage2) decode(authKey []byte, b []byte) error {
	//msgKey := b[:16]
	//// aesKey, aesIV := generateMessageKey(msgKey, authKey, false)
	//// x, err := doAES256IGEdecrypt(b[16:], aesKey, aesIV)
	//
	//x, err := m.decrypt(msgKey, authKey, b[16:])
	//if err != nil {
	//	return err
	//}
	//
	//dbuf := NewDecodeBuf(x)
	//
	//m.Salt = dbuf.Long()      // salt
	//m.SessionId = dbuf.Long() // session_id
	//m.MessageId = dbuf.Long()
	//
	//// mod := m.messageId & 3
	//// if mod != 1 && mod != 3 {
	////	return fmt.Errorf("Wrong bits of message_id: %d", mod)
	//// }
	//
	//m.SeqNo = dbuf.Int()
	//messageLen := dbuf.Int()
	//if int(messageLen) > dbuf.size-32 {
	//	// 	return fmt.Errorf("Message len: %d (need less than %d)", messagxeLen, dbuf.size-32)
	//}
	//
	//m.Object = dbuf.Object()
	//if m.Object == nil {
	//	log.Errorf("salt: %d, sessionId: %d, messageId: %d, seqNo: %d, messageLen: %d", m.Salt, m.SessionId, m.MessageId, m.SeqNo, messageLen)
	//	return fmt.Errorf("decode object is nil")
	//}
	//
	//// log.Info("Recved object: ", m.Object.String())

	return nil
}

func (m *EncryptedMessage2) decrypt(msgKey, authKey, data []byte) ([]byte, error) {
	// dbuf := NewDecodeBuf(data)
	// m.authKeyId = dbuf.Long()
	// msgKey := dbuf.Bytes(16)

	var dataLen = uint32(len(data))
	// 创建aesKey, aesIV
	aesKey, aesIV := GenerateMessageKey(msgKey, authKey, false)
	d := crypto.NewAES256IGECryptor(aesKey, aesIV)

	x, err := d.Decrypt(data)
	if err != nil {
		log.Errorf("descrypted data error: %v", err)
		return nil, err
	}

	//// 校验解密后的数据合法性
	messageLen := binary.LittleEndian.Uint32(x[28:])
	// log.Info("descrypt - messageLen = ", messageLen)
	if messageLen+32 > dataLen {
		// 	return fmt.Errorf("Message len: %d (need less than %d)", messageLen, dbuf.size-32)
		err = fmt.Errorf("descrypted data error: Wrong message length %d", messageLen)
		log.Error(err.Error())
		return nil, err
	}

	messageKey := make([]byte, 96)
	switch MTPROTO_VERSION {
	case 2:
		tmpData := make([]byte, 0, 32+dataLen)
		tmpData = append(tmpData, authKey[88:88+32]...)
		tmpData = append(tmpData, x[:dataLen]...)
		copy(messageKey, crypto.Sha256Digest(tmpData))
	default:
		copy(messageKey[4:], crypto.Sha1Digest(x[:32+messageLen]))
	}

	if !bytes.Equal(messageKey[8:8+16], msgKey[:16]) {
		err = fmt.Errorf("descrypted data error: (data: %s, aesKey: %s, aseIV: %s, authKeyId: %d, authKey: %s), msgKey verify error, sign: %s, msgKey: %s",
			hex.EncodeToString(data[:64]),
			hex.EncodeToString(aesKey),
			hex.EncodeToString(aesIV),
			m.authKeyId,
			hex.EncodeToString(authKey[88:88+32]),
			hex.EncodeToString(messageKey[8:8+16]),
			hex.EncodeToString(msgKey[:16]))
		log.Error(err.Error())
		return nil, err
	}

	return x, nil
}

func (m *EncryptedMessage2) encrypt(authKey []byte, data []byte, messageSize int) ([]byte, error) {
	messageKey := make([]byte, 32)
	switch MTPROTO_VERSION {
	case 2:
		tmpData := make([]byte, 0, 32+len(data))
		tmpData = append(tmpData, authKey[88+8:88+8+32]...)
		tmpData = append(tmpData, data...)
		copy(messageKey, crypto.Sha256Digest(tmpData))
	default:
		copy(messageKey[4:], crypto.Sha1Digest(data[:32+messageSize]))
	}

	// log.Info(data[:messageSize])
	// log.Info(messageKey)
	// copy(message_key[8:], )
	// memcpy(p_data + 8, message_key + 8, 16);

	aesKey, aesIV := GenerateMessageKey(messageKey[8:8+16], authKey, true)
	e := crypto.NewAES256IGECryptor(aesKey, aesIV)

	x, err := e.Encrypt(data)
	if err != nil {
		log.Errorf("Encrypt data error: %v", err)
		return nil, err
	}

	m.msgKey = messageKey[8 : 8+16]
	return x, nil
}