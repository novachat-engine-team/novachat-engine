package codec

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"novachat_engine/pkg/log"
	net2 "novachat_engine/pkg/net"
	"time"
)

type MTProtoHttpProxyCodec struct {
	conn    net.Conn
	canSend bool
}

func NewMTProtoHttpProxyCodec(conn net.Conn) *MTProtoHttpProxyCodec {
	// conn.SetReadDeadline(time.Now().Add(time.Second * 60))
	return &MTProtoHttpProxyCodec{
		conn:    conn,
		canSend: false,
	}
}

func (c *MTProtoHttpProxyCodec) Init(rw io.ReadWriter) {
	c.conn = rw.(net.Conn)
}

func (c *MTProtoHttpProxyCodec) Receive() (interface{}, error) {
	c.conn.SetReadDeadline(time.Now().Add(time.Second * 60))
	req, err := http.ReadRequest(c.conn.(*net2.BufferedConn).BufioReader())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if len(body) < 8 {
		err = fmt.Errorf("not enough uint64 len error - %d", len(body))
		log.Error(err.Error())
		return nil, err
	}

	authKeyId := int64(binary.LittleEndian.Uint64(body))
	msg := NewMTPRawMessage(authKeyId, TRANSPORT_HTTP)
	err = msg.Decode(body)
	if err != nil {
		log.Error(err.Error())
		// conn.Close()
		return nil, err
	}

	c.canSend = true
	return msg, nil
}

func (c *MTProtoHttpProxyCodec) Send(msg interface{}) error {
	// SendToHttpReply(msg, w)
	message, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		log.Error(err.Error())
		// conn.Close()
		return err
	}

	b := message.Encode()

	rsp := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 1,
		Request:    &http.Request{Method: "POST"},
		Header: http.Header{
			"Access-Control-Allow-Headers": {"origin, content-type"},
			"Access-Control-Allow-Methods": {"POST, OPTIONS"},
			"Access-Control-Allow-Origin":  {"*"},
			"Access-Control-Max-Age":       {"1728000"},
			"Cache-control":                {"no-store"},
			"Connection":                   {"keep-alive"},
			// "Connection":                {"close"},
			"Content-type":              {"application/octet-stream"},
			"Pragma":                    {"no-cache"},
			"Strict-Transport-Security": {"max-age=15768000"},
		},
		ContentLength: int64(len(b)),
		Body:          ioutil.NopCloser(bytes.NewReader(b)),
		Close:         false,
		// Close: true,
	}

	err := rsp.Write(c.conn)
	if err != nil {
		log.Error(err.Error())
	}

	c.canSend = false
	return err
}

func (c *MTProtoHttpProxyCodec) Close() error {
	c.canSend = false
	return c.conn.Close()
}
