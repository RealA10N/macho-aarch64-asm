package instructions_test

import (
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/aarch64/instructions"
	"github.com/RealA10N/macho-aarch64-asm/aarch64/registers"

	"github.com/stretchr/testify/assert"
)

func TestSimpleWordAddition(t *testing.T) {
	w0, err := registers.NewGeneratePurposeRegister(0, false)
	assert.NoError(t, err)

	w1, err := registers.NewGeneratePurposeRegister(1, false)
	assert.NoError(t, err)

	inst, err := instructions.NewAddRegister(w0, w0, w1)
	assert.NoError(t, err)

	assert.Equal(t, inst.String(), "add w0, w0, w1")
}

func TestSimpleExtendedAddition(t *testing.T) {
	x1, err := registers.NewGeneratePurposeRegister(1, true)
	assert.NoError(t, err)

	x2, err := registers.NewGeneratePurposeRegister(2, true)
	assert.NoError(t, err)

	x3, err := registers.NewGeneratePurposeRegister(3, true)
	assert.NoError(t, err)

	inst, err := instructions.NewAddRegister(x1, x2, x3)
	assert.NoError(t, err)

	assert.Equal(t, inst.String(), "add x1, x2, x3")
}

func TestWrongRegisterTypes(t *testing.T) {
	w0, err := registers.NewGeneratePurposeRegister(0, false)
	assert.NoError(t, err)

	x1, err := registers.NewGeneratePurposeRegister(1, true)
	assert.NoError(t, err)

	_, err = instructions.NewAddRegister(w0, w0, x1)
	assert.Error(t, err)
}
