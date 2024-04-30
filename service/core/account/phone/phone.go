/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/25 16:08
 * @File : phone.go
 */

package phone

import (
	"errors"
	"fmt"
	"github.com/nyaruka/phonenumbers"
)

type PhoneNumber struct {
	pn *phonenumbers.PhoneNumber
	pv *virtualPhoneNumber
}

func IsVirtualPhone(number string) bool {
	p, err := Parse(number)
	if err != nil {
		return false
	}

	return p.IsVirtualPhone()
}

func Parse(number string) (*PhoneNumber, error) {
	var (
		pn *phonenumbers.PhoneNumber
		err error
		pv *virtualPhoneNumber
	)
	if len(number) == 0 {
		return nil, errors.New("invalid phone number")
	}
	if number[0] != phonenumbers.PLUS_SIGN {
		number = fmt.Sprintf("%c%s",phonenumbers.PLUS_SIGN , number)
	}
	pn, err = phonenumbers.Parse(number, "")
	if err != nil {
		pv, err = parse(number)
	} else {
		v := pn.GetNationalNumber()
		if v == 0 {
			pv, err = parse(number)
		}
	}

	if err != nil {
		return nil, err
	}

	return &PhoneNumber{pn: pn, pv: pv}, nil
}

func (m *PhoneNumber) IsVirtualPhone() bool {
	return m.pv != nil
}

func (m *PhoneNumber) NormalizeDigitsOnly() string {
	if m.pv != nil {
		return m.pv.Number()
	}
	return phonenumbers.NormalizeDigitsOnly(phonenumbers.Format(m.pn, phonenumbers.E164))
}

func (m *PhoneNumber) GetRegionCode() string {
	if m.pv != nil {
		return ""
	}
	return phonenumbers.GetRegionCodeForNumber(m.pn)
}