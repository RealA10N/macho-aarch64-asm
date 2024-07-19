package immediates

import "fmt"

type Immediate12 struct {
	value uint16
}

func NewImmediate12(value uint16) (imm Immediate12, err error) {
	imm = Immediate12{value}

	if value >= 4096 {
		err = fmt.Errorf("value %d overflows immediate12 (maximum value 4046)", value)
		return
	}

	return
}

func (imm Immediate12) String() string {
	return fmt.Sprintf("#%d", imm.value)
}
