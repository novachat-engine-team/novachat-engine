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
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAesCTR128Encrypt(t *testing.T) {
	key := []byte("12345678901234561234567890123456")
	iv := []byte("1234567890123456")

	MyString := "This is my string and I want to protect it with encryption"

	fmt.Printf("We start with a plain text: %s \n", MyString)
	MyStringByte := []byte(MyString)
	encryptor, _ := NewAesCTR128Encrypt(key, iv)
	Encrypted := encryptor.Encrypt(MyStringByte)
	fmt.Printf("We encrypted the string this way: %s \n", hex.EncodeToString(Encrypted))

	decryptor, _ := NewAesCTR128Encrypt(key, iv)
	Decrypted := decryptor.Encrypt(MyStringByte)
	fmt.Printf("Than we have the plain text again: %s \n", string(Decrypted))
}
