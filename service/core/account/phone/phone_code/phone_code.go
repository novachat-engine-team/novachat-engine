/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/25 18:49
 * @File : phone_code.to.go
 */

package phone_code

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"novachat_engine/pkg/cache"
	//cache2 "novachat_engine/service/cache-deprecated"
)

const (
	cacheSMSPhoneCode = "string:phone:number:sms:code:"
)

type PhoneCode struct {
	redisClient *cache.RedisClient
}

var (
	ErrPhoneCodeExpired = errors.New("phone code expired")
)

func NewPhoneCode(r *cache.RedisClient) *PhoneCode {
	return &PhoneCode{redisClient: r}
}

func makePhoneCodeKey(phoneNumber string, t int64) string {
	return fmt.Sprintf("%s%s:t:%d", cacheSMSPhoneCode, phoneNumber, t)
}

func (p *PhoneCode) SetPhoneCode(phoneNumber string, t int64, code string, ttl int32) error {
	_, err := p.redisClient.Do("SETEX", makePhoneCodeKey(phoneNumber, t), ttl, code)
	if err != nil {
		return err
	}

	return nil
}

func (p *PhoneCode) GetPhoneCode(phoneNumber string, t int64) (string, error) {
	v, err := redis.String(p.redisClient.Do("GET", makePhoneCodeKey(phoneNumber, t)))
	if err != nil && err != redis.ErrNil {
		return "", err
	}

	//if err == redis.ErrNil {
	//	return "", ErrPhoneCodeExpired
	//}

	return v, nil
}

func (p *PhoneCode) Clear(phoneNumber string, t int64) error {
	_, err := redis.String(p.redisClient.Do("DEL", makePhoneCodeKey(phoneNumber, t)))
	if err != nil && err != redis.ErrNil {
		return err
	}

	return nil
}

func (p *PhoneCode) TTL(phoneNumber string, t int64) (int32, error) {
	v, err := redis.Int(p.redisClient.Do("TTL", makePhoneCodeKey(phoneNumber, t)))
	if err != nil && err != redis.ErrNil {
		return 0, err
	}

	return int32(v), nil
}
