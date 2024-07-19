package other_test

import (
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/aarch64/other"
	"github.com/stretchr/testify/assert"
)

func TestBasicLabelAssemblyGen(t *testing.T) {
	label := other.Label{Name: "_main"}
	assert.Equal(t, "_main:", label.String())
}
