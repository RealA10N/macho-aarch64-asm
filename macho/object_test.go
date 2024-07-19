package macho_test

import (
	"strings"
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/aarch64"
	"github.com/RealA10N/macho-aarch64-asm/aarch64/instructions"
	"github.com/RealA10N/macho-aarch64-asm/aarch64/other"
	"github.com/RealA10N/macho-aarch64-asm/aarch64/registers"
	"github.com/RealA10N/macho-aarch64-asm/macho"
	"github.com/RealA10N/macho-aarch64-asm/macho/sections"
	"github.com/stretchr/testify/assert"
)

func TestSimpleFunctionAssemblyGen(t *testing.T) {

	x0, err := registers.NewGeneratePurposeRegister(0, true)
	assert.NoError(t, err)

	x1, err := registers.NewGeneratePurposeRegister(1, true)
	assert.NoError(t, err)

	inst, err := instructions.NewAddRegister(x0, x0, x1)
	assert.NoError(t, err)

	global := sections.GlobalSection{Names: []string{"_add"}}

	text := sections.TextSection{Instructions: []aarch64.Instruction{
		other.Label{Name: "_add"}, inst, instructions.Ret{},
	}}

	object := macho.Object{GlobalSection: global, TextSection: text}

	buffer := new(strings.Builder)
	n, err := object.WriteTo(buffer)

	const expected = `.global _add
_add:
add x0, x0, x1
ret
`
	assert.NoError(t, err)
	assert.Equal(t, expected, buffer.String())
	assert.EqualValues(t, len(expected), n)
}
