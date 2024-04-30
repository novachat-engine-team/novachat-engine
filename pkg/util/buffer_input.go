package util

import (
	"encoding/binary"
	"encoding/hex"
	"io"
	"math"
	"novachat_engine/pkg/log"
)

type BufferInput struct {
	buf  []byte
	off  int
	size int
	err  error
}

func NewBufferInput(buf []byte) *BufferInput {
	return &BufferInput{buf, 0, len(buf), nil}
}

func (m *BufferInput) Error() error {
	return m.err
}

func (m *BufferInput) Buf() ([]byte, error) {
	if m.err != nil {
		return nil, m.err
	}

	return m.buf[m.off:], nil
}

func (m *BufferInput) Byte() byte {
	if m.err != nil {
		return 0
	}
	if m.off+1 > m.size {
		log.Errorf("decode byte error - {off: %d, size: %d}", m.off, m.size)
		m.err = io.EOF
		return 0
	}
	x := m.buf[m.off]
	m.off += 1
	return x
}

func (m *BufferInput) Int16() int16 {
	if m.err != nil {
		return 0
	}
	if m.off+2 > m.size {
		log.Errorf("decode int16 error - {off: %d, size: %d}", m.off, m.size)
		m.err = io.EOF
		return 0
	}
	x := binary.LittleEndian.Uint16(m.buf[m.off : m.off+2])
	m.off += 2
	return int16(x)
}

func (m *BufferInput) UInt16() uint16 {
	if m.err != nil {
		return 0
	}
	if m.off+2 > m.size {
		log.Errorf("decode uint16 error - {off: %d, size: %d}", m.off, m.size)
		m.err = io.EOF
		return 0
	}
	x := binary.LittleEndian.Uint16(m.buf[m.off : m.off+2])
	m.off += 2
	return x
}

func (m *BufferInput) Int32() int32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		log.Errorf("decode int32 error - {off: %d, size: %d}", m.off, m.size)
		m.err = io.EOF
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return int32(x)
}

func (m *BufferInput) UInt32() uint32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		log.Errorf("decode uint32 error - {off: %d, size: %d}", m.off, m.size)
		m.err = io.EOF
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return x
}

func (m *BufferInput) UInt64() uint64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		log.Errorf("decode uint64 error - {off: %d, size: %d}", m.off, m.size)
		m.err = io.EOF
		return 0
	}
	x := binary.LittleEndian.Uint64(m.buf[m.off : m.off+8])
	m.off += 8
	return x
}

func (m *BufferInput) Int64() int64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		log.Errorf("decode int64 error - {off: %d, size: %d}", m.off, m.size)
		m.err = io.EOF
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *BufferInput) Double() float64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		log.Errorf("decode double error - {off: %d, size: %d}", m.off, m.size)
		m.err = io.EOF
		return 0
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *BufferInput) Bytes(size int) []byte {
	if m.err != nil {
		return nil
	}
	if m.off+size > m.size {
		log.Errorf("decode size_bytes(%d) error - {off: %d, size: %d}", size, m.off, m.size)
		m.err = io.EOF
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size
	return x
}

func (m *BufferInput) DumpSize(size int) string {
	if m.off+size > m.size {
		size = m.size - m.off
	}
	return hex.Dump(m.buf[m.off : m.off+size])
}

func (m *BufferInput) Dump() string {
	return hex.Dump(m.buf[m.off:m.size])
}
