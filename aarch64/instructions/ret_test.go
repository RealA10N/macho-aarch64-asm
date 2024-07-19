package instructions_test

import (
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/aarch64/instructions"
	"github.com/stretchr/testify/assert"
)

func TestRetString(t *testing.T) {
	inst := instructions.Ret{}
	assert.Equal(t, "ret", inst.String())
}
