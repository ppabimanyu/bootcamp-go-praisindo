/*
 * Copyright (c) 2024. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package phone

import (
	"github.com/nyaruka/phonenumbers"
)

type RegionCode string

const (
	RegionCodeMalaysia  RegionCode = "MY"
	RegionCodeIndonesia RegionCode = "ID"
)

type PhoneNumber struct {
	number *phonenumbers.PhoneNumber
}

func NewPhoneNumber(number string, regionCode RegionCode) (*PhoneNumber, error) {
	parse, err := phonenumbers.Parse(number, string(regionCode))
	if err != nil {
		return nil, err
	}

	return &PhoneNumber{number: parse}, nil
}

func (p *PhoneNumber) String() string {
	return phonenumbers.Format(p.number, phonenumbers.E164)[1:]
}

func (p *PhoneNumber) E164() string {
	return phonenumbers.Format(p.number, phonenumbers.E164)
}

func (p *PhoneNumber) National() string {
	return phonenumbers.Format(p.number, phonenumbers.NATIONAL)
}

func (p *PhoneNumber) International() string {
	return phonenumbers.Format(p.number, phonenumbers.INTERNATIONAL)
}

func (p *PhoneNumber) IsValid() bool {
	return phonenumbers.IsValidNumber(p.number)
}
