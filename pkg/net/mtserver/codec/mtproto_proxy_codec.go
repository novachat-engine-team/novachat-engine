package codec

import (
	"encoding/hex"
	"errors"
	"io"
	"net"
	"novachat_engine/pkg/crypto"
	"novachat_engine/pkg/log"
	pkg_net "novachat_engine/pkg/net"
	// "time"
	"fmt"
	//"novachat_engine/pkg/log"
)

// 客户端or服务端
const (
	CODEC_TYPE_CLIENT = 1 // Client端
	CODEC_TYPE_SERVER = 2 // Server端
)

// Transport类型，不支持UDP
const (
	TRANSPORT_TCP  = 1 // TCP
	TRANSPORT_HTTP = 2 // HTTP
)

const (
	STATE_CONNECTED = iota
)

const (
	// Tcp Transport
	MTPROTO_ABRIDGED_FLAG            = 0xef
	MTPROTO_INTERMEDIATE_FLAG        = 0xeeeeeeee
	MTPROTO_PADDED_INTERMEDIATE_FLAG = 0xdddddddd

	// Http Transport
	HTTP_HEAD_FLAG   = 0x44414548
	HTTP_POST_FLAG   = 0x54534f50
	HTTP_GET_FLAG    = 0x20544547
	HTTP_OPTION_FLAG = 0x4954504f

	VAL2_FLAG = 0x00000000
)

func init() {
	pkg_net.RegisterCodecMaker(pkg_net.MTProtoCodec, NewMTProtoProxyCodec)
}

func NewMTProtoProxyCodec() pkg_net.Codec {
	return &MTProtoProxyCodec{
		codecType: TRANSPORT_TCP,
		State:     STATE_CONNECTED,
	}
}

type MTProtoProxyCodec struct {
	codecType int // codec type
	conn      net.Conn
	State     int
	codec     pkg_net.Codec
	//proto     *MTProtoProxy
}

func (c *MTProtoProxyCodec) Init(rw io.ReadWriter) {
	c.conn = rw.(net.Conn)
}

/////////////////////////////////////////////////////////////////////////////////////
/**
  Android client code:

	RAND_bytes(bytes, 64);
	uint32_t val = (bytes[3] << 24) | (bytes[2] << 16) | (bytes[1] << 8) | (bytes[0]);
	uint32_t val2 = (bytes[7] << 24) | (bytes[6] << 16) | (bytes[5] << 8) | (bytes[4]);
	if (bytes[0] != 0xef &&
		val != 0x44414548 &&
		val != 0x54534f50 &&
		val != 0x20544547 &&
		val != 0x4954504f &&
		val != 0xeeeeeeee &&
		val2 != 0x00000000) {
		bytes[56] = bytes[57] = bytes[58] = bytes[59] = 0xef;
		break;
	}
*/
func (c *MTProtoProxyCodec) peekCodec() error {
	//if c.State == STATE_DATA {
	//	return nil
	//}
	conn, _ := c.conn.(*pkg_net.BufferedConn)
	// var b_0_1 = make([]byte, 1)
	b_0_1, err := conn.Peek(1)

	// _, err := io.ReadFull(c.conn, b_0_1)
	if err != nil {
		log.Warnf("MTProtoProxyCodec - read b_0_1 error: %v", err)
		return err
	}

	if b_0_1[0] == MTPROTO_ABRIDGED_FLAG {
		log.Info("mtproto abridged version.")
		c.codec = NewMTProtoAbridgedCodec(conn)

		conn.Discard(1)
		return nil
	}

	// not abridged version, we'll lookup codec!
	// b_1_3 = make([]byte, 3)
	b_1_3, err := conn.Peek(4)
	if err != nil {
		log.Warnf("MTProtoProxyCodec - read b_1_3 error: %v", err)
		return err
	}

	b_1_3 = b_1_3[1:4]
	// first uint32
	switch (uint32(b_1_3[2]) << 24) | (uint32(b_1_3[1]) << 16) | (uint32(b_1_3[0]) << 8) | (uint32(b_0_1[0])) {
	case HTTP_HEAD_FLAG,
		HTTP_POST_FLAG,
		HTTP_OPTION_FLAG,
		HTTP_GET_FLAG: // http 协议
		log.Info("mtproto http.")
		c.codecType = TRANSPORT_HTTP
		c.codec = NewMTProtoHttpProxyCodec(c.conn)
		return nil
	case MTPROTO_INTERMEDIATE_FLAG: // an intermediate version
		log.Info("mtproto intermediate version.")
		c.codec = NewMTProtoIntermediateCodec(conn)
		conn.Discard(4)
		return nil
	case MTPROTO_PADDED_INTERMEDIATE_FLAG:
		log.Info("mtproto padded intermediate version.")
		c.codec = NewMTProtoPaddedIntermediateCodec(conn)
		conn.Discard(4)
		return nil
	default:
		break
	}

	b_4_60, err := conn.Peek(64)
	if err != nil {
		log.Warnf("MTProtoProxyCodec - read b_4_60 error: %v", err)
		return err
	}

	datacenterId := b_4_60[60] | b_4_60[61]<<8
	log.Debugf("MTProtoProxyCodec datacenterId:%d", datacenterId)

	b_4_60 = b_4_60[4:64]
	val2 := (uint32(b_4_60[3]) << 24) | (uint32(b_4_60[2]) << 16) | (uint32(b_4_60[1]) << 8) | (uint32(b_4_60[0]))
	if val2 == VAL2_FLAG {
		log.Info("mtproto full version.")
		c.codec = NewMTProtoFullCodec(conn)
		return nil
	}

	var tmp [64]byte
	// 生成decrypt_key
	for i := 0; i < 48; i++ {
		tmp[i] = b_4_60[51-i]
	}

	e, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		log.Errorf("NewAesCTR128Encrypt error: %s", err)
		return err
	}

	d, err := crypto.NewAesCTR128Encrypt(b_4_60[4:36], b_4_60[36:52])
	if err != nil {
		log.Warnf("NewAesCTR128Encrypt error: %s", err)
		return err
	}

	d.Encrypt(b_0_1)
	d.Encrypt(b_1_3)
	d.Encrypt(b_4_60)

	if b_4_60[52] != 0xef && b_4_60[53] != 0xef && b_4_60[54] != 0xef && b_4_60[55] != 0xef {
		log.Warnf("MTProtoProxyCodec - first 56~59 byte != 0xef")
		return errors.New("mtproto buf[56:60]'s byte != 0xef")
	}

	log.Debugf("first_bytes_64: %s%s%s", hex.EncodeToString(b_0_1), hex.EncodeToString(b_1_3), hex.EncodeToString(b_4_60))
	c.codec = NewMTProtoAppCodec(conn, d, e)
	conn.Discard(64)

	return nil
}

func (c *MTProtoProxyCodec) Receive() (interface{}, error) {
	if c.codec == nil {
		err := c.peekCodec()
		if err != nil {
			return nil, err
		}
	}
	return c.codec.Receive()
}

func (c *MTProtoProxyCodec) Send(msg interface{}) error {
	if c.codec != nil {
		return c.codec.Send(msg)
	}
	return fmt.Errorf("codec is nil")
}

func (c *MTProtoProxyCodec) Close() error {
	if c.codec != nil {
		return c.codec.Close()
	} else {
		return nil
	}
}
