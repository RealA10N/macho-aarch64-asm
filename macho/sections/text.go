package sections

import (
	"fmt"
	"io"

	"github.com/RealA10N/macho-aarch64-asm/aarch64"
)

type TextSection struct {
	Instructions []aarch64.Instruction
}

func (section TextSection) WriteTo(writer io.Writer) (int64, error) {
	var n int64 = 0

	for _, inst := range section.Instructions {
		toWrite := fmt.Sprintf("%s\n", inst.String())
		k, err := io.WriteString(writer, toWrite)
		n += int64(k)
		if err != nil {
			return n, err
		}
	}

	return n, nil
}
