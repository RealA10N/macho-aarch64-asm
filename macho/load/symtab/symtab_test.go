package symtab_test

import (
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/macho/load"
	"github.com/RealA10N/macho-aarch64-asm/macho/load/symtab"
	"github.com/stretchr/testify/assert"
)

func TestSymtabHeaderExpectedBinary(t *testing.T) {
	expectedSymtab := []byte{
		0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00,
		0x50, 0x02, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00,
		0xA0, 0x02, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00,
	}

	header := symtab.SymtabHeader{
		CommandType:       load.SymbolTable,
		CommandSize:       uint32(symtab.SymTabHeaderSize),
		SymbolTableOffset: 0x250,
		NumOfSymbols:      5,
		StringTableOffset: 0x2A0,
		StringTableSize:   40,
	}

	got, err := header.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, expectedSymtab, got)
}
