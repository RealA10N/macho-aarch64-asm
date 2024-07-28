package nlist64_builders_test

import (
	"testing"

	nlist64 "github.com/RealA10N/macho-aarch64-asm/macho/load/nlist64"
	nlist64_builders "github.com/RealA10N/macho-aarch64-asm/macho/load/nlist64/builders"
	"github.com/RealA10N/macho-aarch64-asm/macho/load/symtab/symbol"
	"github.com/stretchr/testify/assert"
)

func TestSectionNlist64BuilderExpectedBinary(t *testing.T) {
	builder := nlist64_builders.SectionNlist64Builder{
		Name:        "_foo",
		Type:        nlist64.ExternalSymbol,
		Section:     1,
		Offset:      2902,
		Description: nlist64.ReferenceFlagUndefinedNonLazy,
	}

	assert.Equal(t, "_foo", builder.GenString())

	ctx := symbol.EntryListContext{StringTableOffset: 1337}
	nlist, err := builder.GenEntryList(&ctx)
	assert.NoError(t, err)

	expectedNlist := nlist64.Nlist64{
		StringTableOffset: 1337,
		SymbolType:        nlist64.ExternalSymbol | nlist64.SectionSymbolType,
		Section:           1,
		Description:       nlist64.ReferenceFlagUndefinedNonLazy,
		Value:             2902,
	}

	assert.Equal(t, expectedNlist, nlist)
}
