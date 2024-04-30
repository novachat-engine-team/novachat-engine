/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/8/9 18:20
 * @File : mongo_error.go
 */

package mgo

import "go.mongodb.org/mongo-driver/mongo"


const (
	CollesionExists = 48
)

func MongoErrorCode(err error, code int) bool {
	if err == nil {
		return false
	}
	for ; err != nil; err = unwrap(err) {
		if e, ok := err.(mongo.ServerError); ok {
			return e.HasErrorCode(code)
		}
	}

	return false
}

func unwrap(err error) error {
	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return nil
	}
	return u.Unwrap()
}

