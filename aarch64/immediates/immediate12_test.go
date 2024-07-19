package immediates_test

import (
	"fmt"
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/aarch64/immediates"
	"github.com/stretchr/testify/assert"
)

func TestValidCreation(t *testing.T) {
	valids := []uint16{0, 1, 1337, 4095}

	for _, value := range valids {
		imm, err := immediates.NewImmediate12(value)
		assert.NoError(t, err)

		expected := fmt.Sprintf("#%d", value)
		assert.Equal(t, expected, imm.String())
	}
}

func TestInvalidCreation(t *testing.T) {
	_, err := immediates.NewImmediate12(4096)
	assert.Error(t, err)
}
