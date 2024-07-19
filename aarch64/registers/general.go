package registers

import (
	"errors"
	"fmt"
)

type GeneralPurposeRegister struct {
	name     uint8
	extended bool
}

func NewGeneratePurposeRegister(name uint8, extended bool) (register GeneralPurposeRegister, err error) {
	register = GeneralPurposeRegister{name: name, extended: extended}

	if name >= 32 {
		err = errors.New("invalid register name")
	}

	return
}

func (register GeneralPurposeRegister) String() string {
	var prefix rune
	if register.extended {
		prefix = 'x'
	} else {
		prefix = 'w'
	}

	return fmt.Sprintf("%c%d", prefix, register.name)
}

func (register GeneralPurposeRegister) IsExtended() bool {
	return register.extended
}
