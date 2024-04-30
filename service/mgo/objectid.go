/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/2 13:39
 * @File :
 *
 */

package mgo

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var ErrObjectIDLimitCounter = errors.New("math.MaxUint16 limit")
var ErrObjectIDTime = errors.New("time less then last")
var objectIDCounter = readRandomUint32()
var machineCode = uint16(0)
var epoch = time.Time{}.UnixMilli()
var t = int64(0)
var mu = sync.Mutex{}

func MachineCode(m uint16) {
	machineCode = m
}

func Epoch(m int64) {
	epoch = m
}

func readRandomUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot initialize objectid package with crypto.rand.Reader: %v", err))
	}

	return (uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
}

// NewObjectID generates a new ObjectID.
func NewObjectID() (primitive.ObjectID, error) {
	return NewObjectIDFromTimestamp(time.Now())
}

// NewObjectIDFromTimestamp
// mongo driver:
//    binary.BigEndian.PutUint32(b[0:4], uint32(timestamp.Unix()))
//    copy(b[4:9], processUnique[:])
//    putUint24(b[9:12], atomic.AddUint32(&objectIDCounter, 1))
func NewObjectIDFromTimestamp(timestamp time.Time) (primitive.ObjectID, error) {

	var err error
	var b [12]byte

	mu.Lock()
	idCounter := atomic.AddUint32(&objectIDCounter, 1)
	tm := timestamp.UnixMilli() - epoch

	if t == tm && idCounter == 0 {
		mu.Unlock()
		return b, ErrObjectIDLimitCounter
	}
	if tm < t {
		log.Printf("NewObjectIDFromTimestamp panic last time:%d tm:%d\n", tm+epoch, t+epoch)
		err = ErrObjectIDTime
	} else {
		t = tm
	}

	putBytes3(b[0:6], uint64(tm))
	binary.BigEndian.PutUint16(b[6:8], machineCode)
	binary.BigEndian.PutUint32(b[8:12], idCounter)

	mu.Unlock()
	return b, err
}

// ObjectIDTimestamp extracts the time part of the ObjectId.
func ObjectIDTimestamp(objectId primitive.ObjectID) time.Time {
	unixSecs := bytes3(objectId[0:6])
	return time.Unix(int64(unixSecs), 0).UTC()
}

func putBytes3(b []byte, v uint64) {
	b[0] = byte(v >> 40)
	b[1] = byte(v >> 32)
	b[2] = byte(v >> 24)
	b[3] = byte(v >> 16)
	b[4] = byte(v >> 8)
	b[5] = byte(v)
}

func bytes3(b []byte) uint64 {
	return uint64(b[5]) | uint64(b[4])<<8 | uint64(b[3])<<16 | uint64(b[2])<<24 |
		uint64(b[1])<<32 | uint64(b[0])<<40
}
