/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/25 16:07
 * @File : virtual_phone.go
 */

package phone

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
)

var (
	ErrVirtualPhoneNumberFormat = fmt.Errorf("virtual phone number format error: +20000000000000xx or 20000000000000xx")
	ErrVirtualPhoneNumberLen = fmt.Errorf("virtual phone number limit error: +20000000000000xx or 20000000000000xx")
)

type virtualPhoneNumber struct {
	number string
}

func parse(number string) (*virtualPhoneNumber, error) {
	if len(number) > 0 && number[0] == phonenumbers.PLUS_SIGN {
		number = number[1:]
	}

	for _, v := range number {
		ov := v - '0'
		if ov < 0 || ov > 9 {
			return nil, ErrVirtualPhoneNumberFormat
		}
	}

	if len(number) != virtualPhoneNumberLen {
		return nil, ErrVirtualPhoneNumberLen
	}

	return &virtualPhoneNumber{number: number}, nil
}

func (v *virtualPhoneNumber) Number() string {
	return v.number
}
