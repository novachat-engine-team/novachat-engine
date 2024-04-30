package codec

import (
	"encoding/binary"
	"fmt"
	"io"
	"novachat_engine/pkg/log"
	net2 "novachat_engine/pkg/net"
)

// https://core.telegram.org/mtproto/mtproto-transports#padded-intermediate
//Padded version of the intermediate protocol, to use with obfuscation enabled to bypass ISP blocks.
//
//Overhead: small-medium
//Minimum envelope length: random
//Maximum envelope length: random
//Before sending anything into the underlying socket (see transports), the client must first send 0xdddddddd as the first int (four bytes, the server will not send 0xdddddddd as the first int in the first reply).
//Then, payloads are wrapped in the following envelope:
//
//+----+----...----+----...----+
//|tlen|  payload  |  padding  |
//+----+----...----+----...----+
//
//Envelope description:
//
//Total length: payload+padding length encoded as 4 length bytes (little endian)
//Payload: the MTProto payload
//Padding: A random padding string of length 0-15

type MTProtoPaddedIntermediateCodec struct {
	conn *net2.BufferedConn
}

func NewMTProtoPaddedIntermediateCodec(conn *net2.BufferedConn) *MTProtoPaddedIntermediateCodec {
	return &MTProtoPaddedIntermediateCodec{
		conn: conn,
	}
}

func (c *MTProtoPaddedIntermediateCodec) Init(rw io.ReadWriter) {
	c.conn = rw.(*net2.BufferedConn)
}

func (c *MTProtoPaddedIntermediateCodec) Receive() (interface{}, error) {
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

	if size == 4 {
		log.Errorf("Server response error: ", int32(binary.LittleEndian.Uint32(buf)))
		// return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!!") //  connId: ", c.stream, ", by client ", m.RemoteAddr())
		return nil, nil
	}

	authKeyId := int64(binary.LittleEndian.Uint64(buf))
	message := NewMTPRawMessage(authKeyId, TRANSPORT_TCP)
	message.Decode(buf[:size-16])
	return message, nil
}

func (c *MTProtoPaddedIntermediateCodec) Send(msg interface{}) error {
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

func (c *MTProtoPaddedIntermediateCodec) Close() error {
	return c.conn.Close()
}
