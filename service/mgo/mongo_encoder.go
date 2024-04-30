/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/11 13:36
 * @File : mongo_encoder.go
 */

package mgo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"strconv"
	"strings"
)

type Tag string

func (t Tag) String() string {
	return string(t)
}

const (
	DefaultTag Tag = "json"
	BsonTag    Tag = "bson"
)

var (
	DefaultJsonEncoder = NewEncoder(WithEncodeTag(DefaultTag), WithEmitDefaults(true))
	DefaultBsonEncoder = NewEncoder(WithEncodeTag(BsonTag), WithEmitDefaults(true))
	DJE                = DefaultJsonEncoder
	DBE                = DefaultBsonEncoder
)

var _supportTagList = []Tag{DefaultTag, BsonTag}

type Encoder struct {
	encodeTag        Tag
	EmitDefaults     bool
	Enable64ToString bool
}

type Option func(*Encoder)

func WithEncodeTag(t Tag) Option {
	return func(e *Encoder) {
		e.encodeTag = t
	}
}

func WithEmitDefaults(v bool) Option {
	return func(e *Encoder) {
		e.EmitDefaults = v
	}
}

func WithEnable64ToString(v bool) Option {
	return func(e *Encoder) {
		e.Enable64ToString = v
	}
}

func NewEncoder(opts ...Option) *Encoder {
	e := &Encoder{}

	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (m *Encoder) toValueArray(value, valueType reflect.Value, tagName Tag, tList []string) interface{} {
	switch valueType.Kind() {
	case reflect.Uint8:
		return value.Interface().([]byte)
	default:
		v := make([]interface{}, value.Len())
		for j := 0; j < value.Len(); j++ {
			v[j] = m.toValue(value.Index(j), true, tagName, []string{})
			//v[j] = value.Index(j).Interface()
		}
		return v
	}

}

func (m *Encoder) toValue(value reflect.Value, EmitDefaults bool, tagName Tag, tList []string) interface{} {
	switch value.Kind() {
	case reflect.Bool:
		v := value.Bool()
		if v == false && EmitDefaults == false {
			return nil
		}
		return v
	case reflect.Int:
		v := value.Int()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		return int(v)
	case reflect.Int8:
		v := value.Int()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		return int8(v)
	case reflect.Int16:
		v := value.Int()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		return int16(v)
	case reflect.Int32:
		v := value.Int()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		return int32(v)
	case reflect.Int64:
		v := value.Int()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		if m.Enable64ToString == true {
			return strconv.FormatInt(v, 10)
		}
		return v
	case reflect.Uint:
		v := value.Uint()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		return uint(v)
	case reflect.Uint8:
		v := value.Uint()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		return uint8(v)
	case reflect.Uint16:
		v := value.Uint()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		return uint16(v)
	case reflect.Uint32:
		v := value.Uint()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		return uint32(v)
	case reflect.Uint64:
		v := value.Uint()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		if m.Enable64ToString == true {
			return strconv.FormatUint(v, 10)
		}
		return v
	case reflect.Float32:
		v := value.Float()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		return float32(v)
	case reflect.Float64:
		v := value.Float()
		if v == 0 && EmitDefaults == false {
			return nil
		}
		return v
	case reflect.Slice:
		if value.Len() == 0 {
			if EmitDefaults == false {
				return nil
			}
			return nil
		}
		return m.toValueArray(value, value.Index(0), tagName, []string{})
	case reflect.Array:
		if value.Len() == 0 {
			if EmitDefaults == false {
				return nil
			}
			return nil
		}
		return m.toValueArray(value, value.Index(0), tagName, []string{})

	case reflect.String:
		v := value.String()
		if len(v) == 0 && EmitDefaults == false {
			return nil
		}
		return v
	case reflect.Struct, reflect.Ptr:
		v := m.toMap(value.Interface(), tagName, []string{})
		return v
	default:
		panic(fmt.Sprintf("toValue kink:%v value:%s", value.Kind(), value.Type().String()))
	}
	return nil
}

func (m *Encoder) toPtr(sf reflect.StructField, sf1 reflect.Type, value reflect.Value, tagName Tag, tList []string) interface{} {
	switch sf1.Kind() {
	case reflect.Struct:
		m := m.toMap(value.Interface(), tagName, m.subStructList(sf, tList))
		if len(m) > 0 {
			return m
		}
	default:
		v := value.Elem()
		if !v.IsValid() {
			return nil
		}
		return m.toValue(v, m.EmitDefaults, tagName, tList)
	}
	return nil
}

func (m *Encoder) subStructList(sf reflect.StructField, tList []string) []string {
	tmpList := make([]string, 0, len(tList))
	for _, v := range tList {
		if len(v) >= len(sf.Name)+1 && (strings.Compare(sf.Name+".", v[:len(sf.Name)+1]) == 0) {
			n := v[len(sf.Name):]
			if n[1:] == "*" || n[1:] == "" {
				tmpList = tmpList[:]
				break
			}
			tmpList = append(tmpList, n[1:])
		}
	}
	return tmpList
}

func (m *Encoder) toMap(b interface{}, tagName Tag, tList []string) map[string]interface{} {
	bsonMap := make(map[string]interface{})
	if b == nil {
		return bsonMap
	}
	typeOf := reflect.TypeOf(b)
	valueOf := reflect.ValueOf(b)

	if typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
		valueOf = valueOf.Elem()
	}
	if !valueOf.IsValid() {
		return bsonMap
	}

	if typeOf.Kind() != reflect.Struct {
		panic(fmt.Sprintf("typeOf.Kind() != reflect.Struct"))
	}

	for i := 0; i < typeOf.NumField(); i++ {
		sf := typeOf.Field(i)
		if len(tList) > 0 {
			found := true
			for _, vv := range tList {
				if sf.Name == vv || len(vv) >= len(sf.Name)+1 && (strings.Compare(sf.Name+".", vv[:len(sf.Name)+1]) == 0) {
					found = false
					break
				}
			}
			if found {
				continue
			}
		}

		bsonTag := sf.Tag.Get(tagName.String())
		if len(bsonTag) == 0 {
			bsonTag = sf.Tag.Get(DefaultTag.String())
		}
		fieldValue := valueOf.Field(i)
		tagParts := strings.Split(bsonTag, ",")
		bsonTag = tagParts[0]

		switch fieldValue.Kind() {
		case reflect.Struct:
			m1 := m.toMap(fieldValue.Interface(), tagName, m.subStructList(sf, tList))
			if len(m1) > 0 {
				return m1
			}
		case reflect.Ptr:
			m1 := m.toPtr(sf, sf.Type.Elem(), fieldValue, tagName, tList)
			if m1 != nil {
				bsonMap[bsonTag] = m1
			} else {
				if m.EmitDefaults == true {
					bsonMap[bsonTag] = nil
				}
			}
		default:
			m1 := m.toValue(fieldValue, m.EmitDefaults, tagName, m.subStructList(sf, tList))
			if m1 != nil {
				bsonMap[bsonTag] = m1
			} else {
				if m.EmitDefaults == true {
					bsonMap[bsonTag] = nil
				}
			}
		}
	}

	return bsonMap
}

func (m *Encoder) MarshalMap(b interface{}) map[string]interface{} {
	m.encodeTag = DefaultTag
	return m.toMap(b, m.encodeTag, nil)
}

func (m *Encoder) MarshalBsonM(b interface{}) bson.M {
	m.encodeTag = BsonTag
	return m.toMap(b, m.encodeTag, nil)
}

func (m *Encoder) MarshalCustomMap(b interface{}, tag Tag) map[string]interface{} {
	m.encodeTag = tag
	return m.toMap(b, m.encodeTag, nil)
}

func (m *Encoder) MarshalCustomBsonM(b interface{}, tag Tag) bson.M {
	m.encodeTag = tag
	return m.toMap(b, m.encodeTag, nil)
}

func (m *Encoder) MarshalCustomSpecMap(b interface{}, opt ...string) map[string]interface{} {
	return m.toMap(b, m.encodeTag, opt)
}
