/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/30 17:31
 * @File : rpc.go
 */

package rpc

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ParamError = errors.New("params error")
	//locker           sync.Mutex
	_rpcMethodClient map[string]*ClientImpl
)

type ClientImpl struct {
	ImplValue      reflect.Value
	FullMethod     string
	RetValueFunc   func() interface{}
	RpcMethodName  string
	RpcMethodValue reflect.Value
}

func init() {
	_rpcMethodClient = make(map[string]*ClientImpl)
}

func RegisterRpcImpl(desc map[string][]interface{}, impl interface{}) error {
	if desc == nil || impl == nil {
		return ParamError
	}
	value := reflect.ValueOf(impl)

	//locker.Lock()
	//defer locker.Unlock()
	for k, v := range desc {
		_, ok := _rpcMethodClient[k]
		if ok == true {
			return fmt.Errorf("RegisterRpc MethodName is registed error: %s %s desc:%v", k, v[0], desc)
		}

		_rpcMethodClient[k] = &ClientImpl{
			ImplValue:      value,
			FullMethod:     v[0].(string),
			RetValueFunc:   v[1].(func() interface{}),
			RpcMethodName:  v[2].(string),
			RpcMethodValue: value.MethodByName(v[2].(string)),
		}
	}

	return nil
}

func GetRpcImplValue(methodName string) (*ClientImpl, bool) {
	val, ok := _rpcMethodClient[methodName]
	if !ok {
		return nil, false
	}
	return val, true
}

func GetObjectTypeName(object interface{}) string {
	rt := reflect.TypeOf(object)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	return rt.Name()
}
