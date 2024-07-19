package instructions_test

import (
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/aarch64/immediates"
	"github.com/RealA10N/macho-aarch64-asm/aarch64/instructions"
	"github.com/RealA10N/macho-aarch64-asm/aarch64/registers"
	"github.com/stretchr/testify/assert"
)

func TestBasicImmediateAssemblyGen(t *testing.T) {
	x0, err := registers.NewGeneratePurposeRegister(0, true)
	assert.NoError(t, err)

	x1, err := registers.NewGeneratePurposeRegister(1, true)
	assert.NoError(t, err)

	imm, err := immediates.NewImmediate12(1337)
	assert.NoError(t, err)

	inst, err := instructions.NewAddImmediate(x0, x1, imm)
	assert.NoError(t, err)

	assert.Equal(t, "add x0, x1, #1337", inst.String())
}
