package nlist64_test

import (
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/macho/load/nlist64"
	"github.com/stretchr/testify/assert"
)

func TestNlist64ExpectedBinary(t *testing.T) {
	expectedNlist := []byte{
		0x01, 0x00, 0x00, 0x00, 0x0F, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	nlist := nlist64.Nlist64{
		StringTableOffset: 1,
		SymbolType:        nlist64.SectionSymbolType | nlist64.ExternalSymbol,
		Section:           1,
		Description:       nlist64.ReferenceFlagUndefinedNonLazy,
		Value:             0,
	}

	data, err := nlist.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, expectedNlist, data)
}
