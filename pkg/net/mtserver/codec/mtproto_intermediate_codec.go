package codec

import (
	"encoding/binary"
	"fmt"
	"io"
	"novachat_engine/pkg/log"
	net2 "novachat_engine/pkg/net"
)

// https://core.telegram.org/mtproto#tcp-transport
//
// In case 4-byte data alignment is needed,
// an intermediate version of the original protocol may be used:
// if the client sends 0xeeeeeeee as the first int (four bytes),
// then packet length is encoded always by four bytes as in the original version,
// but the sequence number and CRC32 are omitted,
// thus decreasing total packet size by 8 bytes.
//+----+----...----+
//+len.+  payload  +
//+----+----...----+

type MTProtoIntermediateCodec struct {
	conn *net2.BufferedConn
}

func NewMTProtoIntermediateCodec(conn *net2.BufferedConn) *MTProtoIntermediateCodec {
	return &MTProtoIntermediateCodec{
		conn: conn,
	}
}

func (c *MTProtoIntermediateCodec) Init(rw io.ReadWriter) {
	c.conn = rw.(*net2.BufferedConn)
}

func (c *MTProtoIntermediateCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 4)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	size = int(binary.LittleEndian.Uint32(b) << 2)

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(c.conn, buf[size-left:])
		if err != nil {
			log.Errorf("readFull2 error: %v", err)
			return nil, err
		}
		left -= n
	}

	// TODO:process report ack and quickack
	if size == 4 {
		log.Errorf("Server response error: ", int32(binary.LittleEndian.Uint32(buf)))
		// return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!!") //  connId: ", c.stream, ", by client ", m.RemoteAddr())
		return nil, nil
	}

	authKeyId := int64(binary.LittleEndian.Uint64(buf))
	message := NewMTPRawMessage(authKeyId, TRANSPORT_TCP)
	message.Decode(buf)
	return message, nil
}

func (c *MTProtoIntermediateCodec) Send(msg interface{}) error {
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

	//if size < 127 {
	//	sb = []byte{byte(size)}
	//} else {
	binary.LittleEndian.PutUint32(sb, uint32(size))
	//}

	b = append(sb, b...)
	_, err := c.conn.Write(b)

	if err != nil {
		log.Errorf("Send msg error: %s", err)
	}

	return err
}

func (c *MTProtoIntermediateCodec) Close() error {
	return c.conn.Close()
}
