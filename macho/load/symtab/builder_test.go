package symtab_test

import (
	"bytes"
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/macho/builder/context"
	"github.com/RealA10N/macho-aarch64-asm/macho/load/nlist64"
	"github.com/RealA10N/macho-aarch64-asm/macho/load/symtab"
	"github.com/RealA10N/macho-aarch64-asm/macho/load/symtab/symbol"
	"github.com/stretchr/testify/assert"
)

type MySymbolBuilder struct {
	Name  string
	Nlist nlist64.Nlist64
}

func (sym MySymbolBuilder) GenString() string {
	return sym.Name
}

func (sym MySymbolBuilder) GenEntryList(
	ctx *symbol.EntryListContext,
) (nlist64.Nlist64, error) {
	nlist := sym.Nlist
	nlist.StringTableOffset = ctx.StringTableOffset
	return nlist, nil
}

func TestSymtabBuilderExpectedBinary(t *testing.T) {
	expectedHeader := []byte{
		0x02, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00,
		0x56, 0x0B, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
		0x76, 0x0B, 0x00, 0x00, 0x0B, 0x00, 0x00, 0x00,
	}

	expectedData := []byte{
		0x01, 0x00, 0x00, 0x00, 0x0F, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x06, 0x00, 0x00, 0x00, 0x0F, 0x01, 0x00, 0x00,
		0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x5F, 0x61, 0x64, 0x64, 0x00, 0x5F, 0x73,
		0x75, 0x62, 0x00,
	}

	addNlist := nlist64.Nlist64{
		SymbolType:  nlist64.SectionSymbolType | nlist64.ExternalSymbol,
		Section:     1,
		Description: nlist64.ReferenceFlagUndefinedNonLazy,
		Value:       0,
	}

	subNlist := nlist64.Nlist64{
		SymbolType:  nlist64.SectionSymbolType | nlist64.ExternalSymbol,
		Section:     1,
		Description: nlist64.ReferenceFlagUndefinedNonLazy,
		Value:       32,
	}

	builder := symtab.SymtabBuilder{
		Symbols: []symbol.SymbolBuilder{
			MySymbolBuilder{
				Name:  "_add",
				Nlist: addNlist,
			},
			MySymbolBuilder{
				Name:  "_sub",
				Nlist: subNlist,
			},
		},
	}

	// HeaderLen
	assert.EqualValues(t, 24, builder.HeaderLen())

	// DataLen
	assert.EqualValues(t, 43, builder.DataLen())

	{
		// HeaderWriteTo
		buffer := bytes.Buffer{}
		ctx := context.CommandContext{DataOffset: 2902}
		n, err := builder.HeaderWriteTo(&buffer, &ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, len(expectedHeader), n)
		assert.Equal(t, expectedHeader, buffer.Bytes())
	}

	{
		// DataWriteTo
		buffer := bytes.Buffer{}
		n, err := builder.DataWriteTo(&buffer)
		assert.NoError(t, err)
		assert.EqualValues(t, n, len(expectedData))
		assert.Equal(t, expectedData, buffer.Bytes())
	}
}
