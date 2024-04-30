package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Md5V(salt string, str string) string {
	strFinal := salt + str
	h := md5.New()
	h.Write([]byte(strFinal))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckBit(flags int32, offset int32) bool {
	return flags & (1 << offset) != 0
}

func MysqlRealEscapeString(value string) string {
	var sb strings.Builder
	for i := 0; i < len(value); i++ {
		c := value[i]
		switch c {
		case '\\', 0, '\n', '\r', '\'', '"':
			sb.WriteByte('\\')
			sb.WriteByte(c)
		case '\032':
			sb.WriteByte('\\')
			sb.WriteByte('Z')
		default:
			sb.WriteByte(c)
		}
	}
	return sb.String()
}