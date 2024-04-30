/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/25 13:46
 * @File : seq.go
 */

package seq

type SeqNoGen struct {
	value int32
}

func NewSeqNoGen() *SeqNoGen {
	return &SeqNoGen{
		value: 0,
	}
}

func (s *SeqNoGen) Generator(reply bool) int32 {
	v := s.value
	if reply == false {
		return v << 2
	}

	s.value++
	return (v << 2) + 1
}


