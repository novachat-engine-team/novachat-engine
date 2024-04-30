/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/2/21 16:13
 * @File :
 */

package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"novachat_engine/pkg/log"
)

type RSACryptor struct {
	key *rsa.PrivateKey
}

func NewRSACryptor(keyFile string) (*RSACryptor, error) {
	pkcs1PemPrivateKey, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Error("invalid pemsKeyFile: " + keyFile)
		return nil, err
	}
	block, _ := pem.Decode(pkcs1PemPrivateKey)
	if block == nil {
		log.Error("Invalid pemsKeyData: " + string(pkcs1PemPrivateKey))
		return nil, err
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Error("Failed to parse private key: " + err.Error())
		return nil, err
	}

	return &RSACryptor{key}, nil
}

func (m *RSACryptor) Encrypt(b []byte) []byte {
	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(b), big.NewInt(int64(m.key.E)), m.key.N)

	e := c.Bytes()
	var size = len(e)
	if size < 256 {
		size = 256
	}

	res := make([]byte, size)
	copy(res, c.Bytes())

	return res
}

func (m *RSACryptor) Decrypt(b []byte) []byte {
	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(b), m.key.D, m.key.N)
	return c.Bytes()
}
