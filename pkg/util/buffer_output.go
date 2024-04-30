
package util

import (
	"encoding/binary"
	"math"
)

type BufferOutput struct {
	buf []byte
}

func NewBufferOutput(cap int) *BufferOutput {
	return &BufferOutput{make([]byte, 0, cap)}
}

func (e *BufferOutput) Buf() []byte {
	return e.buf
}

func (e *BufferOutput) Len() int {
	return len(e.buf)
}

func (e *BufferOutput) Byte(s byte) {
	e.buf = append(e.buf, s)
}

func (e *BufferOutput) Int16(s int16) {
	e.buf = append(e.buf, 0, 0)
	binary.LittleEndian.PutUint16(e.buf[len(e.buf)-2:], uint16(s))
}

func (e *BufferOutput) UInt16(s uint16) {
	e.buf = append(e.buf, 0, 0)
	binary.LittleEndian.PutUint16(e.buf[len(e.buf)-2:], s)
}

func (e *BufferOutput) Int32(s int32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], uint32(s))
}

func (e *BufferOutput) UInt32(s uint32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], s)
}

func (e *BufferOutput) Int64(s int64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], uint64(s))
}

func (e *BufferOutput) UInt64(s uint64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], s)
}

func (e *BufferOutput) Double(s float64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], math.Float64bits(s))
}

func (e *BufferOutput) Bytes(s []byte) {
	e.buf = append(e.buf, s...)
}

func (e *BufferOutput) ByteSize() int {
	return len(e.buf)
}
