/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package data

import (
	"fmt"
	"novachat_engine/pkg/crypto"
	"novachat_engine/service/constants"
)

type Status int32

const (
	StatusReady       = 0
	StatusReqPQ       = StatusReady
	StatusResQP       = 1
	StatusReqDH       = 2
	StatusResDH       = 3
	StatusDHParamFail = 4
	StatusDHParamOk   = 5
	StatusAuthKey     = 6
	StatusBindAuthKey = 7
)

type Context struct {
	status        Status
	cryptoAuthKey *crypto.AuthKey
	nonce         []byte
	serverNonce   []byte
	p             []byte
	q             []byte
	newNonce      []byte
	a             []byte
	dc            int32 // config.json dc  默认值为2

	layer       int32
	validSince  int32
	validUntil  int32
	expiresIn   int32 // 有效时间
	authKeyTemp constants.AuthKeyStatus

	authKeyId int64
	authKey   []byte

	sessionIdMap map[int64]int
}

func (c *Context) AuthKeyTemp() constants.AuthKeyStatus {
	return c.authKeyTemp
}

func (c *Context) SetAuthKeyTemp(authKeyTemp constants.AuthKeyStatus) {
	c.authKeyTemp = authKeyTemp
}

func (c *Context) ExistsSessionId(sessionId int64) bool {
	if c.sessionIdMap == nil {
		c.sessionIdMap = make(map[int64]int, 1)
	}

	_, ok := c.sessionIdMap[sessionId]

	return ok
}

func (c *Context) SessionIdList() []int64 {
	ret := make([]int64, 0, len(c.sessionIdMap))

	for id := range c.sessionIdMap {
		ret = append(ret, id)
	}

	return ret
}

func (c *Context) SetSessionId(sessionId int64) {
	c.sessionIdMap[sessionId] = 1
}

func (c *Context) CryptoAuthKey() *crypto.AuthKey {
	return c.cryptoAuthKey
}

func (c *Context) SetCryptoAuthKey(cryptoAuthKey *crypto.AuthKey) {
	c.cryptoAuthKey = cryptoAuthKey
}

func (c *Context) AuthKeyId() int64 {
	return c.authKeyId
}

func (c *Context) SetAuthKeyId(authKeyId int64) {
	c.authKeyId = authKeyId
}

func (c *Context) ExpiresIn() int32 {
	return c.expiresIn
}

func (c *Context) SetExpiresIn(expiresIn int32) {
	c.expiresIn = expiresIn
}

func (c *Context) SetValidSince(since int32) {
	c.validSince = since
}

func (c *Context) SetValidUntil(until int32) {
	c.validUntil = until
}

func (c *Context) ValidSince() int32 {
	return c.validSince
}

func (c *Context) ValidUntil() int32 {
	return c.validUntil
}

func (c *Context) SetLayer(layer int32) {
	c.layer = layer
}

func (c *Context) Layer() int32 {
	return c.layer
}

func (c *Context) Dc() int32 {
	return c.dc
}

func (c *Context) SetDc(dc int32) {
	c.dc = dc
}

func (c *Context) A() []byte {
	return c.a
}

func (c *Context) SetA(a []byte) {
	c.a = a
}

func (c *Context) NewNonce() []byte {
	return c.newNonce
}

func (c *Context) SetNewNonce(newNonce []byte) {
	c.newNonce = newNonce
}

func (c *Context) P() []byte {
	return c.p
}

func (c *Context) SetP(p []byte) {
	c.p = p
}

func (c *Context) Q() []byte {
	return c.q
}

func (c *Context) SetQ(q []byte) {
	c.q = q
}

func (c *Context) ServerNonce() []byte {
	return c.serverNonce
}

func (c *Context) SetServerNonce(serverNonce []byte) {
	c.serverNonce = serverNonce
}

func (c *Context) Nonce() []byte {
	return c.nonce
}

func (c *Context) SetNonce(nonce []byte) {
	c.nonce = nonce
}

func (c *Context) AuthKey() []byte {
	return c.authKey
}

func (c *Context) SetAuthKey(authKey []byte) {
	c.authKey = authKey
}

func (c *Context) Status() Status {
	return c.status
}

func (c *Context) SetStatus(status Status) {
	c.status = status
}

func (c *Context) GoString() string {
	return fmt.Sprintf("Context status:%v authKeyId:%x authKeyId:%d", c.status, uint64(c.authKeyId), c.authKeyId)
}

func (c *Context) String() string {
	return fmt.Sprintf("Context status:%v authKeyId:%x authKeyId:%d", c.status, uint64(c.authKeyId), c.authKeyId)
}

func (c *Context) Reset() {
	c.status = StatusReady
}
