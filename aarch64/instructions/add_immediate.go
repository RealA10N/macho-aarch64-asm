package instructions

import (
	"fmt"

	"github.com/RealA10N/macho-aarch64-asm/aarch64/immediates"
	"github.com/RealA10N/macho-aarch64-asm/aarch64/registers"
)

type AddImmediate struct {
	destination, source registers.GeneralPurposeRegister
	immediate           immediates.Immediate12
}

func NewAddImmediate(destination, source registers.GeneralPurposeRegister, immediate immediates.Immediate12) (inst AddImmediate, err error) {
	inst = AddImmediate{destination, source, immediate}

	if source.IsExtended() != destination.IsExtended() {
		err = fmt.Errorf("register type %s does not match type %s", source, destination)
		return
	}

	return
}

func (inst AddImmediate) String() string {
	return fmt.Sprintf("add %s, %s, %s", inst.destination, inst.source, inst.immediate)
}
