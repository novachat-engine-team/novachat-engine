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
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonoptions"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"math"
	"reflect"
	"strconv"
)

type decodeAdapter struct {
	bsoncodec.ValueDecoderFunc
}

var emptyValue = reflect.Value{}

var errCannotTruncate = errors.New("float64 can only be truncated to an integer type when truncation is enabled")

type UIntCodec struct {
	EncodeToMinSize bool
}

var (
	defaultUIntCodec = NewUIntCodec()
)

func NewUIntCodec(opts ...*bsonoptions.UIntCodecOptions) *UIntCodec {
	uintOpt := bsonoptions.MergeUIntCodecOptions(opts...)

	codec := UIntCodec{}
	if uintOpt.EncodeToMinSize != nil {
		codec.EncodeToMinSize = *uintOpt.EncodeToMinSize
	}
	return &codec
}

func (uic *UIntCodec) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	switch val.Kind() {
	case reflect.Uint8, reflect.Uint16:
		return vw.WriteInt32(int32(val.Uint()))
	case reflect.Uint, reflect.Uint32, reflect.Uint64:
		u64 := val.Uint()

		return vw.WriteString(fmt.Sprintf("%d", u64))
	}

	return bsoncodec.ValueEncoderError{
		Name:     "UintEncodeValue",
		Kinds:    []reflect.Kind{reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint},
		Received: val,
	}
}

func (uic *UIntCodec) decodeType(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, t reflect.Type) (reflect.Value, error) {
	var i64 int64
	var err error
	switch vrType := vr.Type(); vrType {
	case bsontype.Int32:
		i32, err := vr.ReadInt32()
		if err != nil {
			return emptyValue, err
		}
		i64 = int64(i32)
	case bsontype.Int64:
		i64, err = vr.ReadInt64()
		if err != nil {
			return emptyValue, err
		}
	case bsontype.Double:
		f64, err := vr.ReadDouble()
		if err != nil {
			return emptyValue, err
		}
		if !dc.Truncate && math.Floor(f64) != f64 {
			return emptyValue, errCannotTruncate
		}
		if f64 > float64(math.MaxInt64) {
			return emptyValue, fmt.Errorf("%g overflows int64", f64)
		}
		i64 = int64(f64)
	case bsontype.Boolean:
		b, err := vr.ReadBoolean()
		if err != nil {
			return emptyValue, err
		}
		if b {
			i64 = 1
		}
	case bsontype.Null:
		if err = vr.ReadNull(); err != nil {
			return emptyValue, err
		}
	case bsontype.String:
		s, err := vr.ReadString()
		if err != nil {
			return emptyValue, err
		}
		i64, _ = strconv.ParseInt(s, 10, 64)
	case bsontype.Undefined:
		if err = vr.ReadUndefined(); err != nil {
			return emptyValue, err
		}
	default:
		return emptyValue, fmt.Errorf("cannot decode %v into an integer type", vrType)
	}

	switch t.Kind() {
	case reflect.Uint8:
		if i64 < 0 || i64 > math.MaxUint8 {
			return emptyValue, fmt.Errorf("%d overflows uint8", i64)
		}

		return reflect.ValueOf(uint8(i64)), nil
	case reflect.Uint16:
		if i64 < 0 || i64 > math.MaxUint16 {
			return emptyValue, fmt.Errorf("%d overflows uint16", i64)
		}

		return reflect.ValueOf(uint16(i64)), nil
	case reflect.Uint32:
		if i64 < 0 || i64 > math.MaxUint32 {
			return emptyValue, fmt.Errorf("%d overflows uint32", i64)
		}

		return reflect.ValueOf(uint32(i64)), nil
	case reflect.Uint64:
		if i64 < 0 {
			return emptyValue, fmt.Errorf("%d overflows uint64", i64)
		}

		return reflect.ValueOf(uint64(i64)), nil
	case reflect.Uint:
		if i64 < 0 || int64(uint(i64)) != i64 { // Can we fit this inside of an uint
			return emptyValue, fmt.Errorf("%d overflows uint", i64)
		}

		return reflect.ValueOf(uint(i64)), nil
	default:
		return emptyValue, bsoncodec.ValueDecoderError{
			Name:     "UintDecodeValue",
			Kinds:    []reflect.Kind{reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint},
			Received: reflect.Zero(t),
		}
	}
}

// DecodeValue is the ValueDecoder for uint types.
func (uic *UIntCodec) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() {
		return bsoncodec.ValueDecoderError{
			Name:     "UintDecodeValue",
			Kinds:    []reflect.Kind{reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint},
			Received: val,
		}
	}

	elem, err := uic.decodeType(dc, vr, val.Type())
	if err != nil {
		return err
	}

	val.SetUint(elem.Uint())
	return nil
}

func fitsIn32Bits(i int64) bool {
	return math.MinInt32 <= i && i <= math.MaxInt32
}

func IntEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32:
		return vw.WriteInt32(int32(val.Int()))
	case reflect.Int:
		i64 := val.Int()
		if fitsIn32Bits(i64) {
			return vw.WriteInt32(int32(i64))
		}
		return vw.WriteInt64(i64)
	case reflect.Int64:
		i64 := val.Int()
		return vw.WriteString(fmt.Sprintf("%d", i64))
	}

	return bsoncodec.ValueEncoderError{
		Name:     "IntEncodeValue",
		Kinds:    []reflect.Kind{reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int},
		Received: val,
	}
}

func IntDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() {
		return bsoncodec.ValueDecoderError{
			Name:     "IntDecodeValue",
			Kinds:    []reflect.Kind{reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int},
			Received: val,
		}
	}

	elem, err := IntDecodeType(dc, vr, val.Type())
	if err != nil {
		return err
	}

	val.SetInt(elem.Int())
	return nil
}

func IntDecodeType(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, t reflect.Type) (reflect.Value, error) {
	var i64 int64
	var err error
	switch vrType := vr.Type(); vrType {
	case bsontype.Int32:
		i32, err := vr.ReadInt32()
		if err != nil {
			return emptyValue, err
		}
		i64 = int64(i32)
	case bsontype.Int64:
		i64, err = vr.ReadInt64()
		if err != nil {
			return emptyValue, err
		}
	case bsontype.Double:
		f64, err := vr.ReadDouble()
		if err != nil {
			return emptyValue, err
		}
		if !dc.Truncate && math.Floor(f64) != f64 {
			return emptyValue, errCannotTruncate
		}
		if f64 > float64(math.MaxInt64) {
			return emptyValue, fmt.Errorf("%g overflows int64", f64)
		}
		i64 = int64(f64)
	case bsontype.Boolean:
		b, err := vr.ReadBoolean()
		if err != nil {
			return emptyValue, err
		}
		if b {
			i64 = 1
		}
	case bsontype.Null:
		if err = vr.ReadNull(); err != nil {
			return emptyValue, err
		}
	case bsontype.String:
		s, err := vr.ReadString()
		if err != nil {
			return emptyValue, err
		}
		i64, _ = strconv.ParseInt(s, 10, 64)

	case bsontype.Undefined:
		if err = vr.ReadUndefined(); err != nil {
			return emptyValue, err
		}
	default:
		return emptyValue, fmt.Errorf("cannot decode %v into an integer type", vrType)
	}

	switch t.Kind() {
	case reflect.Int8:
		if i64 < math.MinInt8 || i64 > math.MaxInt8 {
			return emptyValue, fmt.Errorf("%d overflows int8", i64)
		}

		return reflect.ValueOf(int8(i64)), nil
	case reflect.Int16:
		if i64 < math.MinInt16 || i64 > math.MaxInt16 {
			return emptyValue, fmt.Errorf("%d overflows int16", i64)
		}

		return reflect.ValueOf(int16(i64)), nil
	case reflect.Int32:
		if i64 < math.MinInt32 || i64 > math.MaxInt32 {
			return emptyValue, fmt.Errorf("%d overflows int32", i64)
		}

		return reflect.ValueOf(int32(i64)), nil
	case reflect.Int64:
		return reflect.ValueOf(i64), nil
	case reflect.Int:
		if int64(int(i64)) != i64 { // Can we fit this inside of an int
			return emptyValue, fmt.Errorf("%d overflows int", i64)
		}

		return reflect.ValueOf(int(i64)), nil
	default:
		return emptyValue, bsoncodec.ValueDecoderError{
			Name:     "IntDecodeValue",
			Kinds:    []reflect.Kind{reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int},
			Received: reflect.Zero(t),
		}
	}
}
