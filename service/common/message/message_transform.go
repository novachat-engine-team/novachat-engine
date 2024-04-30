/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/25 14:01
 * @File : message_transform.go
 */

package message

import (
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
)

type Transform struct {
	Salt      int64
	SessionId int64
	Buf       []byte
}

func (t *Transform) Encode(layer int32) []byte {
	encoder := mtproto.NewEncodeBuf(8 + 8 + len(t.Buf))
	encoder.Long(t.Salt)
	encoder.Long(t.SessionId)
	encoder.Bytes(t.Buf)
	return encoder.GetBuf()
}

func (t *Transform) Decode(d []byte, layer int32) error {
	decoder := mtproto.NewDecodeBuf(d)
	t.Salt = decoder.Long()
	err := decoder.GetError()
	if err != nil {
		log.Errorf("Transform Salt error:%s", err.Error())
		return err
	}
	t.SessionId = decoder.Long()
	err = decoder.GetError()
	if err != nil {
		log.Errorf("Transform SessionId error:%s", err.Error())
		return err
	}
	t.Buf = decoder.Bytes(decoder.GetSize() - decoder.GetOffset())
	err = decoder.GetError()
	if err != nil {
		log.Errorf("Transform Buf error:%s", err.Error())
		return err
	}
	return nil
}
