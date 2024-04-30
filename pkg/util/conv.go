package util

import (
	"reflect"
	"strconv"
)

func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func StringToInt32(s string) (int32, error) {
	i, err := strconv.Atoi(s)
	return int32(i), err
}

func StringToUint32(s string) (uint32, error) {
	i, err := strconv.Atoi(s)
	return uint32(i), err
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func BoolToInt8(b bool) int8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func Int8ToBool(b int8) bool {
	if b == 1 {
		return true
	} else {
		return false
	}
}

func BoolToInt32(b bool) int32 {
	if b {
		return 1
	} else {
		return 0
	}
}

func Int32ToBool(b int32) bool {
	if b == 1 {
		return true
	} else {
		return false
	}
}

func IndexInt32s(l *[]int32, dst int32) int {
	if l == nil {
		return -1
	}

	for idx, v := range *l {
		if v == dst {
			return idx
		}
	}
	return -1
}

func IndexInt64s(l *[]int64, dst int64) int {
	if l == nil {
		return -1
	}

	for idx, v := range *l {
		if v == dst {
			return idx
		}
	}
	return -1
}

func IndexStrings(l *[]string, dst string) int {
	if l == nil {
		return -1
	}

	for idx, v := range *l {
		if v == dst {
			return idx
		}
	}
	return -1
}

func Index(x interface{}, f func(idx int) bool) int {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic(&reflect.ValueError{Method: "Index", Kind: v.Kind()})
	}

	for idx := 0; idx < v.Len(); idx++ {
		if f != nil && f(idx) {
			return idx
		}
	}
	return -1
}

func Find(x interface{}, f func(key interface{}) bool) interface{} {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Map {
		panic(&reflect.ValueError{Method: "Find", Kind: v.Kind()})
	}

	mapIter := v.MapRange()
	for mapIter.Next() {
		if f(mapIter.Key().Interface()) {
			return mapIter.Value().Interface()
		}
	}
	return nil
}

func Foreach(x interface{}, f func(first interface{}, second interface{})) {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Map && v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic(&reflect.ValueError{Method: "Foreach", Kind: v.Kind()})
	}

	if v.Kind() == reflect.Map {
		mapIter := v.MapRange()
		for mapIter.Next() {
			f(mapIter.Key().Interface(), mapIter.Value().Interface())
		}
	} else {
		for idx := 0; idx < v.Len(); idx++ {
			f(idx, v.Index(idx))
		}
	}
}
