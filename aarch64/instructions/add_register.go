package instructions

import (
	"fmt"

	"github.com/RealA10N/macho-aarch64-asm/aarch64/registers"
)

type AddRegister struct {
	destination, first, second registers.GeneralPurposeRegister
}

func NewAddRegister(destination, first, second registers.GeneralPurposeRegister) (inst AddRegister, err error) {
	inst = AddRegister{destination, first, second}

	for _, register := range []*registers.GeneralPurposeRegister{&first, &second} {
		if register.IsExtended() != destination.IsExtended() {
			err = fmt.Errorf("register type %s does not match destination type %s", register, destination)
			return
		}
	}

	return
}

func (inst AddRegister) String() string {
	return fmt.Sprintf("ADD %s, %s, %s", inst.destination, inst.first, inst.second)
}
