package util

import (
	"encoding/hex"
	"novachat_engine/pkg/hack"
)

func WriteString(x *BufferOutput, s string) {
	WriteBytes(x, hack.Bytes(s))
}

func WriteBytes(x *BufferOutput, b []byte) {
	x.UInt32(uint32(len(b)))
	x.Bytes(b)
}

func ReadString(dbuf *BufferInput) (s string, err error) {
	var b []byte
	if b, err = ReadBytes(dbuf); err != nil {
		s = hack.String(b)
	}
	return
}

func ReadBytes(dbuf *BufferInput) (b []byte, err error) {
	var n = dbuf.UInt32()
	b = dbuf.Bytes(int(n))
	err = dbuf.Error()
	return
}

func DumpSize(size int, buf []byte) string {
	if size > len(buf) {
		size = len(buf)
	}
	return hex.Dump(buf[:size])
}

func Dump(buf []byte) string {
	return DumpSize(128, buf)
}

func HexDumpSize(size int, buf []byte) string {
	if size > len(buf) {
		size = len(buf)
	}
	return hex.EncodeToString(buf[:size])
}

func HexDump(buf []byte) string {
	return HexDumpSize(128, buf)
}
