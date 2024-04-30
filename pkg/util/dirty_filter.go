package util

import (
	"bufio"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"os"
	"strings"
	"sync"
)

var sensitiveWord = &sync.Map{}

var invalidWord = make(map[rune]interface{})

func ByteListToRuneList(b []byte) []rune {
	ret := make([]rune, 0, len(b)/2)
	for idx := 0; idx < len(b); idx += 2 {
		v1 := int32(b[idx+1])<<8 | int32(b[idx])
		ret = append(ret, rune(v1))
	}
	return ret
}

func RuneListToByteList(b []rune) []byte {
	ret := make([]byte, 0, len(b)*2)
	for idx := 0; idx < len(b); idx += 1 {
		ret = append(ret, byte(b[idx]&0xff))
		ret = append(ret, byte((b[idx]&0xff00)>>8))
	}
	return ret
}

func dirtyFilterSensitive(set map[string]interface{}) {

	for key := range set {

		str := ByteListToRuneList([]byte(key))
		nowMap := sensitiveWord

		for i := 0; i < len(str); i++ {
			if _, ok := nowMap.Load(str[i]); !ok {
				thisMap := &sync.Map{}
				thisMap.Store(rune(0), false)
				nowMap.Store(str[i], thisMap)
				nowMap = thisMap
			} else {
				v, _ := nowMap.Load(str[i])
				nowMap = v.(*sync.Map)
			}

			if i == len(str)-1 {
				nowMap.Store(rune(0), true)
			}
		}
	}

}

func DirtyFilterSensitiveWords(txt string) (string, bool) {

	bsUTF16LE, _, err := transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder(), []byte(txt))
	if err != nil {
		return "", false
	}
	str := ByteListToRuneList(bsUTF16LE)
	is := false
	nowMap := sensitiveWord
	start := -1
	tag := -1
	var thisMap *sync.Map
	var ok bool
	var vv interface{}

	for i := 0; i < len(str); i++ {
		if _, ok := invalidWord[str[i]]; ok {
			continue
		}

		if vv, ok = nowMap.Load(str[i]); ok {
			thisMap = vv.(*sync.Map)
			tag++
			if tag == 0 {
				start = i
			}

			if vv, _ = thisMap.Load(rune(0)); vv != nil {
				isEnd := vv.(bool)
				if isEnd {
					for y := start; y < i+1; y++ {
						str[y] = 42

						if !is {
							is = true
						}
					}
					nowMap = sensitiveWord
					start = -1
					tag = -1

				} else {
					vv, ok = nowMap.Load(str[i])
					if ok == true {
						nowMap = vv.(*sync.Map)
					}
				}
			}
		} else {
			if start != -1 {
				i = start + 1
			}
			nowMap = sensitiveWord
			start = -1
			tag = -1
		}

	}

	aStr := RuneListToByteList(str)
	bsUTF8LE, _, _ := transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder(), aStr)
	return string(bsUTF8LE), is
}

func dirtyFilterOpenFile(filePath string, invalid string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	sensitiveWord = &sync.Map{}
	var setSensitiveWord = make(map[string]interface{})
	var line string
	buff := bufio.NewReader(file)
	for {

		line, err = buff.ReadString('\n')
		if err != nil {
			if io.EOF == err {
				break
			}
			return err
		}

		line = SplitWord(line)
		if len(line) == 0 {
			continue
		}

		setSensitiveWord[line] = nil
	}

	if invalid != "" {
		utf16Words, _, _ := transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder(), []byte(invalid))
		words := ByteListToRuneList(utf16Words)

		for idx := 0; idx < len(words); idx++ {
			invalidWord[words[idx]] = nil
		}
	}

	dirtyFilterSensitive(setSensitiveWord)
	return nil
}

func SplitWord(words string) string {
	words = strings.TrimSpace(words)
	if len(words) == 0 {
		return ""
	}
	utf16Line, _, _ := transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder(), []byte(words))
	return string(utf16Line)
}

func DirtyFilterOpenFile(filePath string, invalid string) error {
	err := dirtyFilterOpenFile(filePath, invalid)
	if err != nil {
		return err
	}

	return nil
}

func FilterDirtyWord(msg string, bReplace bool) (string, bool) {
	return DirtyFilterSensitiveWords(msg)
}

func FilterDirtyPush(msg []string) bool {
	var setSensitiveWord = make(map[string]interface{})
	for _, m := range msg {
		m = SplitWord(m)

		if len(m) == 0 {
			continue
		}

		setSensitiveWord[m] = nil
	}

	if len(setSensitiveWord) > 0 {
		dirtyFilterSensitive(setSensitiveWord)
	}
	return len(setSensitiveWord) > 0
}
