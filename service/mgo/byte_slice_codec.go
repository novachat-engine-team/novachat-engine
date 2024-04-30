// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package mgo

import (
	"encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsonoptions"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// ByteSliceCodec is the Codec used for []byte values.
type ByteSliceCodec struct {
	EncodeNilAsEmpty bool
}

var tByteSlice = reflect.TypeOf([]byte(nil))

var (
	defaultByteSliceCodec = NewByteSliceCodec()

	_ bsoncodec.ValueCodec = defaultByteSliceCodec
	//_ bsoncodec.typeDecoder = defaultByteSliceCodec
)

type decodeBinaryError struct {
	subtype  byte
	typeName string
}

func (d decodeBinaryError) Error() string {
	return fmt.Sprintf("only binary values with subtype 0x00 or 0x02 can be decoded into %s, but got subtype %v", d.typeName, d.subtype)
}

// NewByteSliceCodec returns a StringCodec with options opts.
func NewByteSliceCodec(opts ...*bsonoptions.ByteSliceCodecOptions) *ByteSliceCodec {
	byteSliceOpt := bsonoptions.MergeByteSliceCodecOptions(opts...)
	codec := ByteSliceCodec{}
	if byteSliceOpt.EncodeNilAsEmpty != nil {
		codec.EncodeNilAsEmpty = *byteSliceOpt.EncodeNilAsEmpty
	}
	return &codec
}

// EncodeValue is the ValueEncoder for []byte.
func (bsc *ByteSliceCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tByteSlice {
		return bsoncodec.ValueEncoderError{Name: "ByteSliceEncodeValue", Types: []reflect.Type{tByteSlice}, Received: val}
	}
	if val.IsNil() && !bsc.EncodeNilAsEmpty {
		return vw.WriteNull()
	}

	encoding := base64.StdEncoding
	size := encoding.EncodedLen(val.Len())
	buf := make([]byte, size)
	encoding.Encode(buf, val.Interface().([]byte))
	return vw.WriteString(string(buf))
	//return vw.WriteBinary(val.Interface().([]byte))
}

func (bsc *ByteSliceCodec) decodeType(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, t reflect.Type) (reflect.Value, error) {
	if t != tByteSlice {
		return emptyValue, bsoncodec.ValueDecoderError{
			Name:     "ByteSliceDecodeValue",
			Types:    []reflect.Type{tByteSlice},
			Received: reflect.Zero(t),
		}
	}

	var data []byte
	var err error
	switch vrType := vr.Type(); vrType {
	case bsontype.String:
		str, err := vr.ReadString()
		if err != nil {
			return emptyValue, err
		}
		data = []byte(str)
		encoding := base64.StdEncoding
		size := encoding.DecodedLen(len(data))
		buf := make([]byte, size)
		base64.StdEncoding.Decode(buf, data)
		data = buf

		//str, err := vr.ReadString()
		//if err != nil {
		//	return emptyValue, err
		//}
		//data = []byte(str)
	case bsontype.Symbol:
		sym, err := vr.ReadSymbol()
		if err != nil {
			return emptyValue, err
		}
		data = []byte(sym)
	case bsontype.Binary:
		var subtype byte
		data, subtype, err = vr.ReadBinary()
		if err != nil {
			return emptyValue, err
		}
		if subtype != bsontype.BinaryGeneric && subtype != bsontype.BinaryBinaryOld {
			return emptyValue, decodeBinaryError{subtype: subtype, typeName: "[]byte"}
		}
	case bsontype.Null:
		err = vr.ReadNull()
	case bsontype.Undefined:
		err = vr.ReadUndefined()
	default:
		return emptyValue, fmt.Errorf("cannot decode %v into a []byte", vrType)
	}
	if err != nil {
		return emptyValue, err
	}

	return reflect.ValueOf(data), nil
}

// DecodeValue is the ValueDecoder for []byte.
func (bsc *ByteSliceCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tByteSlice {
		return bsoncodec.ValueDecoderError{Name: "ByteSliceDecodeValue", Types: []reflect.Type{tByteSlice}, Received: val}
	}

	elem, err := bsc.decodeType(dc, vr, tByteSlice)
	if err != nil {
		return err
	}

	val.Set(elem)
	return nil
}
