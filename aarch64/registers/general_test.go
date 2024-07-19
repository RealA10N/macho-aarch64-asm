package registers_test

import (
	"macho-aarch64-asm/aarch64/registers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAllValid(t *testing.T) {
	for _, extended := range []bool{false, true} {
		var name uint8 = 0
		for ; name < 32; name++ {
			_, err := registers.NewGeneratePurposeRegister(name, extended)
			assert.NoError(t, err)
		}
	}
}

func TestNewRegister32(t *testing.T) {
	for _, extended := range []bool{false, true} {
		_, err := registers.NewGeneratePurposeRegister(32, extended)
		assert.Error(t, err)
	}
}

func TestRegisterMethods(t *testing.T) {
	register, err := registers.NewGeneratePurposeRegister(0, false)
	assert.NoError(t, err)
	assert.Equal(t, "w0", register.String())
	assert.Equal(t, false, register.IsExtended())

	register, err = registers.NewGeneratePurposeRegister(31, true)
	assert.NoError(t, err)
	assert.Equal(t, "x31", register.String())
	assert.Equal(t, true, register.IsExtended())
}
