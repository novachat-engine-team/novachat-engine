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
// If a payload (packet) needs to be transmitted from server to client or from client to server,
// it is encapsulated as follows:
// 4 length bytes are added at the front
// (to include the length, the sequence number, and CRC32; always divisible by 4)
// and 4 bytes with the packet sequence number within this TCP connection
// (the first packet sent is numbered 0, the next one 1, etc.),
// and 4 CRC32 bytes at the end (length, sequence number, and payload together).
//+----+----+----...----+----+
//|len.|seq.|  payload  |crc.|
//+----+----+----...----+----+

type MTProtoFullCodec struct {
	conn *net2.BufferedConn
}

func NewMTProtoFullCodec(conn *net2.BufferedConn) *MTProtoFullCodec {
	return &MTProtoFullCodec{
		conn: conn,
	}
}

func (c *MTProtoFullCodec) Init(rw io.ReadWriter) {
	c.conn = rw.(*net2.BufferedConn)
}

func (c *MTProtoFullCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 4)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	size = int(binary.LittleEndian.Uint32(b) << 2)
	// Check bufLen
	if size < 12 || size%4 != 0 {
		err = fmt.Errorf("invalid len: %d", size)
		return nil, err
	}

	left := size
	buf := make([]byte, size-4)
	for left > 0 {
		n, err = io.ReadFull(c.conn, buf[size-left:])
		if err != nil {
			log.Errorf("ReadFull2 error: %v", err)
			return nil, err
		}
		left -= n
	}

	seqNum := binary.LittleEndian.Uint32(buf[:4])
	// TODO: check seqNum, save last seq_num
	_ = seqNum

	crc32 := binary.LittleEndian.Uint32(buf[len(buf)-4:])
	// TODO: check crc32
	_ = crc32

	authKeyId := int64(binary.LittleEndian.Uint64(buf[4:]))
	message := NewMTPRawMessage(authKeyId, TRANSPORT_TCP)
	message.Decode(buf)
	return message, nil
}

func (c *MTProtoFullCodec) Send(msg interface{}) error {
	message, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		log.Error(err.Error())
		return err
	}

	b := message.Encode()

	sb := make([]byte, 8)
	// minus padding
	size := len(b) / 4

	binary.LittleEndian.PutUint32(sb, uint32(size))
	// TODO: gen seq_num
	var seqNum uint32 = 0
	binary.LittleEndian.PutUint32(sb[4:], seqNum)
	b = append(sb, b...)
	var crc32Buf = make([]byte, 4)
	var crc32 uint32 = 0
	binary.LittleEndian.PutUint32(crc32Buf, crc32)
	b = append(sb, crc32Buf...)

	_, err := c.conn.Write(b)
	if err != nil {
		log.Errorf("Send msg error: %s", err)
	}

	return err
}

func (c *MTProtoFullCodec) Close() error {
	return c.conn.Close()
}
