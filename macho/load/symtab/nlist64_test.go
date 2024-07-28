package symtab_test

import (
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/macho/load/symtab"
	"github.com/stretchr/testify/assert"
)

func TestNlist64ExpectedBinary(t *testing.T) {
	expectedNlist := []byte{
		0x01, 0x00, 0x00, 0x00, 0x0F, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	nlist := symtab.Nlist64{
		StringTableOffset: 1,
		SymbolType:        symtab.SectionSymbolType | symtab.ExternalSymbolBit,
		Section:           1,
		Description:       symtab.ReferenceFlagUndefinedNonLazy,
		Value:             0,
	}

	data, err := nlist.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, expectedNlist, data)
}
