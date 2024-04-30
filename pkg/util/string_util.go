package util

import "strconv"

func JoinInt32List(s []int32, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendInt(buf, int64(s[i]), 10)
		// buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

func JoinUint32List(s []uint32, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendUint(buf, uint64(s[i]), 10)
		buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

func JoinInt64List(s []int64, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendInt(buf, s[i], 10)
		buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

func JoinUint64List(s []uint64, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendUint(buf, s[i], 10)
		buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

func StringListToInt64List(s []string) []int64 {
	l := make([]int64, 0, len(s))
	for _, v := range s {
		vv, _ := strconv.ParseInt(v, 10, 64)
		l = append(l, vv)
	}

	return l
}

// IsAlNumString returns true if an alpha numeric string consists of characters a-zA-Z0-9

// https://github.com/nebulaim/telegramd/issues/99
//
// there are some issue in username validation
// 1 - issue : username can be only numeric [must check first character for alphabet]
// 2 - issue :underscore not allowed [must edit IsAlNumString and add underscore support to it.
//   by the way underscore can't repeat without any character between them. for example s_a_b_a. ]
// Accepted characters: A-z (case-insensitive), 0-9 and underscores.
// Length: 5-32 characters

const LimitUserNameMinLength = 5
const LimitUserNameMaxLength = 32

func IsAlNumString(s string) bool {
	return allNumberString(s, LimitUserNameMinLength, true)
}

func IsAllNumStringLimit(s string, limit int32) bool {
	return allNumberString(s, limit, false)
}

func allNumberString(s string, limit int32, b bool) bool {
	if limit <= 0 {
		limit = 1
	}

	if int32(len(s)) < limit || len(s) > LimitUserNameMaxLength {
		return false
	}

	if b == true {
		if '0' <= s[0] && s[0] <= '9' {
			return false
		}
	}

	c := 0
	prevtmp := ' '
	for _, r := range s {
		switch {
		case '0' <= r && r <= '9':
			c++
			break
		case 'a' <= r && r <= 'z':
			c++
			break
		case 'A' <= r && r <= 'Z':
			c++
			break
		case prevtmp != '_' && '_' == r:
			c++
			break
		}
		prevtmp = r
	}
	return len(s) == c
}