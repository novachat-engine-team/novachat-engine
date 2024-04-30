/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/6 17:53
 * @File : msg.go
 */

package manual_msgs

import (
	"fmt"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
)

func JsonValue(v *mtproto.JSONValue) string {

	var printJson func(val *mtproto.JSONValue) string
	printJson = func(val *mtproto.JSONValue) string {
		switch val.ClassName {
		case mtproto.ClassJsonBool:
			return fmt.Sprintf("%v", mtproto.ToBool(val.ValueC7345E6A90))
		case mtproto.ClassJsonString:
			return fmt.Sprintf("%v", val.ValueB71E767A90)
		case mtproto.ClassJsonNumber:
			return fmt.Sprintf("%v", val.Value2BE0DFA490)
		case mtproto.ClassJsonNull:
			return "NULL"
		case mtproto.ClassJsonArray:
			ret := "["
			for _, i := range val.ValueF744476390 {
				ret += printJson(i) + ","
			}
			ret += "]"
			return ret
		case mtproto.ClassJsonObject:
			ret := "{"
			for _, i := range val.Value99C1D49D90 {

				ret += fmt.Sprintf("key:%v v:%v, ", i.Key, printJson(i.Value))
			}
			ret += "}"
			return ret
		default:
			return "unknown"
		}
	}

	ret := printJson(v)

	log.Infof("*mtproto.TLJsonObject value:%v", ret)
	return ret
}

func ManualTLInvokeWithLayer(msgId int64, seqNo int32, msg *mtproto.TLInvokeWithLayer) ([]*mtproto.MessageTransform, *mtproto.TLInitConnection, error) {
	msgs := make([]*mtproto.MessageTransform, 0, 1)

	if msg.Query == nil {
		log.Errorf("ManualTLInvokeWithLayer Query == nil msgId:%d", msgId)
		return msgs, nil, nil
	}

	initConnection, ok := msg.Query.(*mtproto.TLInitConnection)
	if ok == false {
		log.Errorf("ManualTLInvokeWithLayer Query not *mtproto.TLInitConnection, object:%v", msg.Query)
		return msgs, nil, nil
	}

	if initConnection.JSONValue != nil {
		JsonValue(initConnection.JSONValue)
	}

	msgs = append(msgs, &mtproto.MessageTransform{
		Message: mtproto.TLMessageTransform{
			MsgId:  msgId,
			Seqno:  seqNo,
			Object: initConnection.Query,
		},
	})

	initConnection.Query = nil
	return msgs, initConnection, nil
}
