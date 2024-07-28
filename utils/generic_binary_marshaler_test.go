package utils_test

import (
	"math"
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/utils"
	"github.com/stretchr/testify/assert"
)

type MyStruct struct {
	X int16
	Y uint64
	Z float64
}

func TestExpected(t *testing.T) {
	s := MyStruct{X: 1337, Y: 18446744073709551615, Z: math.Inf(1)}
	got, err := utils.GenericMarshalBinary(s)
	expected := []byte{
		0x39, 0x05,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x7F,
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}
