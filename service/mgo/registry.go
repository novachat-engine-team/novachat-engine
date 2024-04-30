/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/2 13:39
 * @File :
 *
 */

package mgo

import (
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonoptions"
	"reflect"
)

var int64ToStringRegistry *bsoncodec.Registry

func init() {

	rb := bsoncodec.NewRegistryBuilder()

	dve := bsoncodec.DefaultValueEncoders{}
	dve.RegisterDefaultEncoders(rb)
	rb.RegisterDefaultEncoder(reflect.Int64, bsoncodec.ValueEncoderFunc(IntEncodeValue))
	rb.RegisterDefaultEncoder(reflect.Struct, newDefaultStructCodec())
	rb.RegisterDefaultEncoder(reflect.Uint64, defaultUIntCodec)
	rb.RegisterDefaultEncoder(reflect.Slice, NewByteSliceCodec(bsonoptions.ByteSliceCodec().SetEncodeNilAsEmpty(true)))

	dvd := bsoncodec.DefaultValueDecoders{}
	dvd.RegisterDefaultDecoders(rb)
	intDecoder := decodeAdapter{IntDecodeValue}
	rb.RegisterDefaultDecoder(reflect.Int64, intDecoder)
	rb.RegisterDefaultDecoder(reflect.Uint64, defaultUIntCodec)
	rb.RegisterDefaultDecoder(reflect.Struct, newDefaultStructCodec())
	rb.RegisterDefaultDecoder(reflect.Slice, NewByteSliceCodec(bsonoptions.ByteSliceCodec().SetEncodeNilAsEmpty(false)))
	rb.RegisterTypeDecoder(tByteSlice, NewByteSliceCodec(bsonoptions.ByteSliceCodec().SetEncodeNilAsEmpty(false)))
	int64ToStringRegistry = rb.Build()
}

func Registry() *bsoncodec.Registry {
	return int64ToStringRegistry
}
