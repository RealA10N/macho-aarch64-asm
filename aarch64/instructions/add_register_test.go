package instructions_test

import (
	"macho-aarch64-asm/aarch64/instructions"
	"macho-aarch64-asm/aarch64/registers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleWordAddition(t *testing.T) {
	w0, err := registers.NewGeneratePurposeRegister(0, false)
	assert.NoError(t, err, "expected valid register")

	w1, err := registers.NewGeneratePurposeRegister(1, false)
	assert.NoError(t, err, "expected valid register")

	inst, err := instructions.NewAddRegister(w0, w0, w1)
	assert.NoError(t, err, "expected instruction")

	assert.Equal(t, inst.String(), "ADD w0, w0, w1")
}

func TestSimpleExtendedAddition(t *testing.T) {
	x1, err := registers.NewGeneratePurposeRegister(1, true)
	assert.NoError(t, err, "expected valid register")

	x2, err := registers.NewGeneratePurposeRegister(2, true)
	assert.NoError(t, err, "expected valid register")

	x3, err := registers.NewGeneratePurposeRegister(3, true)
	assert.NoError(t, err, "expected valid register")

	inst, err := instructions.NewAddRegister(x1, x2, x3)
	assert.NoError(t, err, "expected instruction")

	assert.Equal(t, inst.String(), "ADD x1, x2, x3")
}

func TestWrongRegisterTypes(t *testing.T) {
	w0, err := registers.NewGeneratePurposeRegister(0, false)
	assert.NoError(t, err, "expected valid register")

	x1, err := registers.NewGeneratePurposeRegister(1, true)
	assert.NoError(t, err, "expected valid register")

	_, err = instructions.NewAddRegister(w0, w0, x1)
	assert.Error(t, err, "expected invalid instruction")
}
