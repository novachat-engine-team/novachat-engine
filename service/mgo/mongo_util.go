/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/2 13:39
 * @File : mongo_util.go
 */

package mgo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

func ToBsonM(obj interface{}) bson.M {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	value := bson.M{}
	switch t.Kind() {
	case reflect.Map:
		for _, k := range v.MapKeys() {
			mv := v.MapIndex(k)
			value[fmt.Sprintf("%v", k.Interface())] = mv.Interface()
		}

	default:
		for idx := 0; idx < t.NumField(); idx++ {
			fv := v.Field(idx).Interface()
			tfv := reflect.TypeOf(fv)
			ffv := reflect.ValueOf(fv)
			if tfv.Kind() == reflect.Ptr {
				tfv = tfv.Elem()
				ffv = ffv.Elem()
			}
			switch tfv.Kind() {
			case reflect.Struct, reflect.Map:
				value[t.Field(idx).Name] = ToBsonM(ffv.Interface())
			default:
				value[t.Field(idx).Name] = fv
			}
		}
	}
	return value
}

func ToBsonMList(obj interface{}) []bson.M {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if !(t.Kind() == reflect.Array  || t.Kind() == reflect.Slice) {
		panic(fmt.Sprintf("ToBsonMList type:%s error", t.Kind().String()))
	}

	value := make([]bson.M, 0, 1)
	for i:=0; i< v.Len(); i++ {
		value = append(value, ToBsonM(v.Index(i).Interface()))
	}

	return value
}

func BsonListToInterfaceList(l []bson.M) []interface{} {
	r := make([]interface{}, 0, len(l))
	for _, v := range l {
		r = append(r, v)
	}

	return r
}