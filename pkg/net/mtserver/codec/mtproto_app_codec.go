package codec

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/log"
	pkg_net "novachat_engine/pkg/net"
)

type MTProtoAppCodec struct {
	// conn *net.TCPConn
	stream *AesCTR128Stream
}

func NewMTProtoAppCodec(conn *pkg_net.BufferedConn, d *crypto.AesCTR128Encrypt, e *crypto.AesCTR128Encrypt) *MTProtoAppCodec {
	return &MTProtoAppCodec{
		// conn:   conn,
		stream: NewAesCTR128Stream(conn, d, e),
	}
}

func (c *MTProtoAppCodec) Init(rw io.ReadWriter) {
}

func (c *MTProtoAppCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 1)
	n, err = io.ReadFull(c.stream, b)
	if err != nil {
		return nil, err
	}

	needAck := bool(b[0]>>7 == 1)
	_ = needAck

	b[0] = b[0] & 0x7f
	// log.Debugf("first_byte2: %s", hex.EncodeToString(b[:1]))

	if b[0] < 0x7f {
		size = int(b[0]) << 2
		log.Infof("size1: %d", size)
		if size == 0 {
			return nil, nil
		}
	} else {
		log.Debugf("first_byte2: %s", hex.EncodeToString(b[:1]))
		b2 := make([]byte, 3)
		n, err = io.ReadFull(c.stream, b2)
		if err != nil {
			return nil, err
		}
		size = (int(b2[0]) | int(b2[1])<<8 | int(b2[2])<<16) << 2
		log.Debugf("size2:%d ", size)
	}

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(c.stream, buf[size-left:])
		if err != nil {
			log.Errorf("ReadFull2 error: %v", err)
			return nil, err
		}
		left -= n
	}
	if size > 10240 {
		log.Debugf("ReadFull2: %s", hex.EncodeToString(buf[:256]))
	}

	if size == 4 {
		log.Errorf("Server response error: ", int32(binary.LittleEndian.Uint32(buf)))
		// return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!!") //  connId: ", c.stream, ", by client ", m.RemoteAddr())
		return nil, nil
	}

	authKeyId := int64(binary.LittleEndian.Uint64(buf))
	message := NewMTPRawMessage(authKeyId, TRANSPORT_TCP)
	message.Decode(buf)
	return message, nil

	//var message MessageBase
	//if authKeyId == 0 {
	//	message = NewUnencryptedRawMessage()
	//	// message.Decode(buf[8:])
	//} else {
	//	message = NewEncryptedRawMessage(authKeyId)
	//}
	//
	//err = message.Decode(buf[8:])
	//if err != nil {
	//	log.Errorf("decode message error: {%v}", err)
	//	return nil, err
	//}
	//
	// return message, nil
}

func (c *MTProtoAppCodec) Send(msg interface{}) error {
	message, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		log.Error(err.Error())
		return err
	}

	b := message.Encode()

	sb := make([]byte, 4)
	// minus padding
	size := len(b) / 4

	if size < 127 {
		sb = []byte{byte(size)}
	} else {
		binary.LittleEndian.PutUint32(sb, uint32(size<<8|127))
	}

	b = append(sb, b...)
	_, err := c.stream.Write(b)

	if err != nil {
		log.Errorf("Send msg error: %s", err)
	}

	return err
}

func (c *MTProtoAppCodec) Close() error {
	log.Errorf("Close: %+v", c.stream.conn.RemoteAddr().String())
	return c.stream.conn.Close()
}

type AesCTR128Stream struct {
	conn    *pkg_net.BufferedConn
	encrypt *crypto.AesCTR128Encrypt
	decrypt *crypto.AesCTR128Encrypt
}

func NewAesCTR128Stream(conn *pkg_net.BufferedConn, d *crypto.AesCTR128Encrypt, e *crypto.AesCTR128Encrypt) *AesCTR128Stream {
	return &AesCTR128Stream{
		conn:    conn,
		decrypt: d,
		encrypt: e,
	}
}

func (this *AesCTR128Stream) Read(p []byte) (int, error) {
	n, err := this.conn.Read(p)
	if err == nil {
		this.decrypt.Encrypt(p[:n])
		return n, nil
	}
	return n, err
}

func (this *AesCTR128Stream) Write(p []byte) (int, error) {
	this.encrypt.Encrypt(p[:])
	return this.conn.Write(p)
}
