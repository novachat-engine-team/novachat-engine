/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/20 16:03
 * @File : send_message_util.go
 */

package message

import (
	"github.com/mvdan/xurls"
	"math/rand"
	"novachat_engine/pkg/mention"
	"time"
	"unicode/utf8"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Pos struct {
	Beg  int32
	End  int32
	Text string
}

func SplitMessage(msg string) []int {
	return _split(msg)
}

func _split(msg string) []int {
	idxList := make([]int, 0, len(msg))
	for i := 0; i < len(msg); {
		_, n := utf8.DecodeRuneInString(msg[i:])
		idxList = append(idxList, i)
		i += n
	}

	return idxList
}

func Index(a, b []rune, start int) int {
	if len(a) < len(b)+start {
		return -1
	}

	aLen := len(a)
	var distance int = -1
	for j := range a {
		if j < start {
			continue
		}
		for i := range b {
			if aLen > j+i {
				if a[j+i] == b[i] {
					if distance == -1 {
						distance = j
					}
				} else {
					distance = -1
					break
				}
			} else {
				return -1
			}
		}
		if distance != -1 {
			return distance
		}
	}

	return distance
}

func ParseUrl(text string) ([]Pos, bool) {
	utf16Text := []rune(text)
	subStrList := xurls.Relaxed.FindAllStringIndex(text, -1)
	if len(subStrList) > 0 {
		var lastIndex int32 = 0
		posList := make([]Pos, 0, len(subStrList))
		for _, v := range subStrList {
			utf16SubStr := []rune(text[v[0]:v[1]])
			var pos Pos
			pos.Beg = int32(Index(utf16Text, utf16SubStr, int(lastIndex)))
			if pos.Beg > 0 {
				pos.End = int32(len(utf16SubStr))
				lastIndex = pos.Beg + pos.End + 1
				pos.Text = text[v[0]:v[1]]
				posList = append(posList, pos)
			}
		}

		return posList, true
	} else {
		return nil, false
	}
}

func ParseMention(text string) ([]string, []Pos, bool) {
	tags := mention.GetTags('@', text)
	if len(tags) == 0 {
		return nil, nil, false
	}
	var nameList = make([]string, 0, len(tags))
	posList := make([]Pos, 0, len(tags))
	for _, tag := range tags {
		nameList = append(nameList, tag.Tag)
		posList = append(posList, Pos{
			Beg: int32(tag.Index),
			End: int32(len(tag.Tag)),
		})
	}
	return nameList, posList, false
}
