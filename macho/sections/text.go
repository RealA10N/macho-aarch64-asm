package sections

import (
	"errors"
	"io"

	"github.com/RealA10N/macho-aarch64-asm/aarch64"
)

type TextSection struct {
	Instructions []aarch64.Instruction
}

func (section TextSection) WriteTo(writer io.Writer) (int64, error) {
	return 0, errors.New("not implemented")
}
