/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/11 13:37
 * @File : mongo_encoder_test.go
 */

package mgo

import (
	"fmt"
	"testing"
)

type BsonTags struct {
	Id int64 `json:"id" bson:"bson_id"`
	Name string `json:"name" bson:"bson_name"`
	BsonTag *BsonTags `json:"json_bson_tag" bson:"bson_tag"`
	IdV *int64 `json:"id_v" bson:"bson_id_v"`
	IdSlice []int64 `json:"id_slice" bson:"bson_id_slice"`
	IdInt int64 `json:"id_int" bson:"bson_id_int" db:"db_id_int"`
}

func TestBsonTag(t *testing.T) {
	b := BsonTags{
		Id:   100,
		Name: "ssssss",
		BsonTag: &BsonTags{
			Id:      222,
			Name:    "2222",
			BsonTag: &BsonTags{
				Id:      333,
				Name:    "3333",
				BsonTag: nil,
			},
		},
	}

	e := Encoder{
		EmitDefaults: false,
	}
	t.Log(fmt.Sprintf("%+v", e.MarshalMap(b)))
	t.Log(fmt.Sprintf("%+v", e.MarshalBsonM(b)))
	t.Log(fmt.Sprintf("%+v", e.MarshalCustomMap(b, "db")))
}

